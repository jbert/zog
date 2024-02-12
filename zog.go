package zog // import "github.com/jbert/zog"

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Zog struct {
	Mem *Memory
	reg Registers

	/* In the Z80 CPU, there is
	   an interrupt enable flip-flop (IFF) that is set or reset by the programmer using the Enable
	   Interrupt (EI) and Disable Interrupt (DI) instructions. When the IFF is reset, an interrupt
	   cannot be accepted by the CPU.
	*/
	is InterruptState

	interruptCh chan byte

	outputHandler func(uint16, byte)
	inputHandler  func(uint16) byte

	traces Regions

	eTrace executeTrace

	instIters int

	numRecentTraces   int
	indexRecentTraces int
	recentTraces      []executeTrace
}

type InterruptState struct {
	IFF1 bool
	IFF2 bool
	Mode byte
}

type locWatch struct {
	new byte
	old byte
}

type executeTrace struct {
	ops     uint64
	pc      uint16
	inst    Instruction
	reg     Registers
	watches map[uint16]locWatch
}

func (et *executeTrace) String() string {
	s := fmt.Sprintf("%d: %04X %s %s", et.ops, et.pc, et.reg.Summary(), et.inst)
	for addr, lw := range et.watches {
		s += fmt.Sprintf(" W:%04X [%02X->%02X]", addr, lw.old, lw.new)
	}
	return s
}

func New(memSize uint16) *Zog {
	is := InterruptState{
		Mode: 1,
	}
	z := &Zog{
		Mem:         NewMemory(memSize),
		interruptCh: make(chan byte),
		is:          is,
	}
	z.Clear()
	return z
}

type Region struct {
	start uint16
	end   uint16
}
type Regions []Region

func NewRegion(start, end uint16) Region {
	return Region{start: start, end: end}
}

func ParseRegions(s string) (Regions, error) {
	var regions Regions
	s = strings.TrimSpace(s)
	if s == "" {
		return regions, nil
	}
	startEnds := strings.Split(s, ",")
	for _, startEnd := range startEnds {
		bits := strings.Split(startEnd, "-")
		start, err := strconv.ParseUint(bits[0], 16, 16)
		if err != nil {
			return regions, err
		}
		end, err := strconv.ParseUint(bits[1], 16, 16)
		if err != nil {
			return regions, err
		}
		region := NewRegion(uint16(start), uint16(end))
		regions.add(Regions{region})
	}
	return regions, nil
}

func (rs *Regions) contains(addr uint16) bool {
	for _, r := range *rs {
		if r.contains(addr) {
			return true
		}
	}
	return false
}

func (rs *Regions) add(regions Regions) {
	*rs = append(*rs, regions...)
}

func (rs Regions) String() string {
	var strs []string
	for _, r := range rs {
		strs = append(strs, r.String())
	}
	return strings.Join(strs, ",")
}

func (r *Region) contains(addr uint16) bool {
	return r.start <= addr && addr < r.end
}

func (r Region) String() string {
	return fmt.Sprintf("%04X-%04X", r.start, r.end)
}

func (z *Zog) TraceRegions(regions Regions) error {
	z.traces.add(regions)
	return nil
}

func (z *Zog) WatchRegions(regions Regions) error {
	z.Mem.SetWatchFunc(z.memWatchSeen)
	z.Mem.watches.add(regions)
	return nil
}

func (z *Zog) TraceOnHalt(numHaltTraces int) {
	z.numRecentTraces = numHaltTraces
}

func (z *Zog) State() string {
	return fmt.Sprintf("%s AF %04X BC %04X DE %04X HL %04X SP %04X IX %04X IY %04X",
		z.FlagString(),
		z.reg.Read16(AF),
		z.reg.Read16(BC),
		z.reg.Read16(DE),
		z.reg.Read16(HL),
		z.reg.Read16(SP),
		z.reg.Read16(IX),
		z.reg.Read16(IY))
}

func (z *Zog) RunAssembly(a *Assembly) error {
	err := z.Load(a)
	if err != nil {
		return err
	}
	return z.execute(a.BaseAddr)
}

func (z *Zog) RunBytes(loadAddr uint16, buf []byte, runAddr uint16) error {
	err := z.LoadBytes(loadAddr, buf)
	if err != nil {
		return nil
	}
	return z.execute(runAddr)
}

func (z *Zog) LoadROMFile(addr uint16, fname string) error {
	return z.LoadFile(addr, fname, true)
}

func (z *Zog) LoadFile(addr uint16, fname string, readonly bool) error {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return fmt.Errorf("Can't load file [%s]: %s", fname, err)
	}
	if readonly {
		roRegion := NewRegion(addr, uint16(len(buf)))
		fmt.Printf("Add RO region %s\n", roRegion)
		z.Mem.readonly = append(z.Mem.readonly, roRegion)
	}
	return z.LoadBytes(addr, buf)
}

func (z *Zog) LoadRegisters(reg Registers) {
	z.reg = reg
}

func (z *Zog) GetRegisters() Registers {
	return z.reg
}

func (z *Zog) LoadInterruptState(is InterruptState) {
	z.is = is
}

func (z *Zog) LoadBytes(addr uint16, buf []byte) error {
	err := z.Mem.Copy(addr, buf)
	if err != nil {
		return err
	}
	//	fmt.Printf("L: %04X [%04X]\n", addr, len(buf))
	return nil
}

func (z *Zog) Load(a *Assembly) error {
	buf, err := a.Encode()
	if err != nil {
		return err
	}
	return z.LoadBytes(a.BaseAddr, buf)
}

func (z *Zog) Clear() {
	z.Mem.Clear()
	z.reg = Registers{}
	// 64KB will give zero here, correctly
	z.reg.SP = uint16(z.Mem.Len())
	z.is.IFF1 = false
	z.is.IFF2 = false
}

func (z *Zog) RegisterOutputHandler(handler func(uint16, byte)) error {
	if z.outputHandler != nil {
		return errors.New("Output handler already registered")
	}
	z.outputHandler = handler
	return nil
}

func (z *Zog) RegisterInputHandler(handler func(uint16) byte) error {
	if z.inputHandler != nil {
		return errors.New("Input handler already registered")
	}
	z.inputHandler = handler
	return nil
}

// F flag register:
// S Z X H  X P/V N C
type flag int

const (
	F_C flag = iota
	F_N
	F_PV
	F_3
	F_H
	F_5
	F_Z
	F_S
)

func (f flag) String() string {
	switch f {
	case F_C:
		return "C"
	case F_N:
		return "N"
	case F_PV:
		return "PV"
	case F_3:
		return "X1"
	case F_H:
		return "H"
	case F_5:
		return "X2"
	case F_Z:
		return "Z"
	case F_S:
		return "S"
	default:
		panic(fmt.Sprintf("Unknown flag: %d", f))
	}
}

func (z *Zog) SetFlag(f flag, new bool) {
	mask := byte(1) << uint(f)
	flags, err := F.Read8(z)
	if err != nil {
		panic(fmt.Sprintf("Error reading flags: %s", err))
	}
	if new {
		flags = flags | mask
	} else {
		mask = ^mask
		flags = flags & mask
	}
	err = F.Write8(z, flags)
	if err != nil {
		panic(fmt.Sprintf("Error writing flags: %s", err))
	}
}
func (z *Zog) GetFlag(f flag) bool {
	mask := byte(1) << uint(f)
	flags, err := F.Read8(z)
	if err != nil {
		panic(fmt.Sprintf("Error reading flags: %s", err))
	}
	flag := flags & mask
	return flag != 0
}

func (z *Zog) FlagString() string {
	s := ""
	for i := 7; i >= 0; i-- {
		f := flag(i)
		if f == F_3 || f == F_5 {
			continue
		}
		v := 0
		if z.GetFlag(f) {
			v = 1
		}
		s += fmt.Sprintf("%s%d", f, v)
	}
	return s
}

var ErrHalted = errors.New("HALT called")

// Implement io.Reader
func (z *Zog) Read(buf []byte) (int, error) {
	if len(buf) != 1 {
		panic("Non-byte read")
	}
	n, err := z.Mem.Peek(z.reg.PC)
	if err != nil {
		return 0, fmt.Errorf("Error reading: %s", err)
	}
	z.reg.PC++
	buf[0] = n
	return 1, nil
}

func (z *Zog) jp(addr uint16) {
	z.reg.PC = addr
	//	fmt.Printf("JP: %04X\n", z.reg.PC)
}

func (z *Zog) jr(d int8) {
	z.reg.PC += uint16(d) // Wrapping works out
	//	fmt.Printf("JR: %04X [%d]\n", z.reg.PC, d)
}

func (z *Zog) di() error {
	z.is.IFF1 = false
	z.is.IFF2 = false
	return nil
}

func (z *Zog) ei() error {
	z.is.IFF1 = true
	z.is.IFF2 = true
	return nil
}

func (z *Zog) InterruptEnabled() bool {
	return z.is.IFF1
}

func (z *Zog) im(mode byte) error {
	if mode != 0 && mode != 1 && mode != 2 {
		panic(fmt.Sprintf("Invalid interrupt mode: %d", mode))
	}
	z.is.Mode = mode
	return nil
}

func (z *Zog) push(nn uint16) {
	z.reg.SP--
	z.reg.SP--
	err := z.Mem.Poke16(z.reg.SP, nn)
	if err != nil {
		panic(fmt.Sprintf("Can't write to SP [%04X]: %s", z.reg.SP, err))
	}
}

func (z *Zog) pop() uint16 {
	nn, err := z.Mem.Peek16(z.reg.SP)
	if err != nil {
		panic(fmt.Sprintf("Can't write to SP [%04X]: %s", z.reg.SP, err))
	}
	z.reg.SP++
	z.reg.SP++
	return nn
}

func (z *Zog) out(port uint16, n byte) {
	if z.outputHandler == nil {
		fmt.Fprintf(os.Stderr, "Output handler not yet registered, %02X not send to %04X\n", n, port)
	}
	z.outputHandler(port, n)
}

func (z *Zog) in(port uint16) byte {
	if z.inputHandler != nil {
		return z.inputHandler(port)
	}
	fmt.Fprintf(os.Stderr, "Input handler not yet registered, %04X not read\n", port)
	return 0xff
}

func (z *Zog) addRecentTrace(et executeTrace) {
	if z.numRecentTraces <= 0 {
		panic("wtf")
	}
	if len(z.recentTraces) < z.numRecentTraces {
		z.recentTraces = append(z.recentTraces, et)
		return
	}
	z.recentTraces[z.indexRecentTraces] = et
	z.indexRecentTraces++
	z.indexRecentTraces = z.indexRecentTraces % z.numRecentTraces
	return
}

func (z *Zog) DoInterrupt() {
	if !z.InterruptEnabled() {
		return
	}
	z.interruptCh <- z.is.Mode
}

func (z *Zog) getInstruction(halted bool) (Instruction, error) {
	// Check for interrupt
	imMode := byte(0)
	select {
	case imMode = <-z.interruptCh:
		return z.processInterrupt(imMode)
	default:
	}

	if !halted {
		return DecodeOne(z)
	}
	imMode = <-z.interruptCh
	return z.processInterrupt(imMode)
}

func (z *Zog) processInterrupt(imMode byte) (Instruction, error) {
	z.di()
	switch imMode {
	case 0:
		return nil, fmt.Errorf("TODO: interrupt in mode 0")
	case 1:
		// We need to do RST 38h
		return &RST{0x38}, nil
	case 2:
		// We get LSB from data bus (we choose 0) and MSG from I reg
		addr := uint16(z.reg.I) << 8
		return NewCALL(True, Imm16(addr)), nil
	default:
		return nil, fmt.Errorf("Unknown interrupt mode: %d", imMode)
	}
}

func (z *Zog) execute(addr uint16) (errRet error) {
	z.reg.PC = addr
	return z.Run()
}

const ClockHz = 3500000 / 2

var TStateDuration = time.Second / time.Duration(ClockHz)

func (z *Zog) Run() (errRet error) {
	ops := uint64(0)
	TStates := uint64(0)
	lastOps := uint64(0)
	lastTStates := uint64(0)
	statsEvery := uint64(1000000)
	startTime := time.Now()
	lastEmit := startTime

	var err error
	var inst Instruction

	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				errRet = v
			default:
				errRet = fmt.Errorf("PANIC: %v", v)
			}
		}
		if z.numRecentTraces > 0 {
			for i := range z.recentTraces {
				fmt.Printf("%s\n", z.recentTraces[(i+z.indexRecentTraces)%z.numRecentTraces].String())
			}
		}
	}()

	halted := false

EXECUTING:
	for {

		before := time.Now()

		lastPC := z.reg.PC
		// May be from PC, or may be interrupt
		inst, err = z.getInstruction(halted)
		halted = false
		if err != nil {
			fmt.Printf("Error decoding: %s\n", err)
			break EXECUTING
		}

		//		fmt.Printf("I: %04X %s\n", lastPC, inst)
		z.instIters = 0
		instErr := inst.Execute(z)
		waitTStates := inst.TStates(z)

		z.eTrace = executeTrace{ops: ops, pc: lastPC, reg: z.reg, inst: inst, watches: make(map[uint16]locWatch)}
		if z.traces.contains(lastPC) {
			println(z.eTrace.String())
		}
		if z.numRecentTraces > 0 {
			z.addRecentTrace(z.eTrace)
		}
		if instErr != nil {
			if instErr == ErrHalted && z.InterruptEnabled() {
				halted = true
				continue EXECUTING
			}

			// Error handling after the loop
			err = instErr
			break EXECUTING
		}
		ops++
		TStates += uint64(waitTStates)
		if ops%statsEvery == 0 {
			now := time.Now()
			dur := now.Sub(lastEmit)
			secs := dur.Seconds()
			opsPerSec := float64(ops-lastOps) / secs
			TStatesPerSec := float64(TStates-lastTStates) / secs
			fmt.Printf("%5.2f: PC [%04X] %d ops %d t : %f ops/sec %f t/sec\n", now.Sub(startTime).Seconds(), lastPC, ops, TStates, opsPerSec, TStatesPerSec)
			lastEmit = now
			lastOps = ops
			lastTStates = TStates
		}

		//if waitTStates > 30 {
		//	println(waitTStates, " ", z.eTrace.String())
		//}
		waitDuration := time.Duration(waitTStates) * TStateDuration
		// Busy wait - can't get time.Sleep or syscall.Nanosleep to give me good enough granularity
		waitUntil := before.Add(waitDuration)
		for time.Now().Before(waitUntil) {
		}
	}

	// The only return should be on HALT. nil is bad here.
	if err == ErrHalted {
		return nil
	}
	if err == nil {
		return errors.New("Execute returned nil error")
	}
	return fmt.Errorf("Failed to execute: %s", err)
}

func (z *Zog) memWatchSeen(addr uint16, old byte, new byte) {
	z.eTrace.watches[addr] = locWatch{old: old, new: new}
}
