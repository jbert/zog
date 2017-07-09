package zog

import (
	"errors"
	"fmt"
)

type Zog struct {
	mem *Memory
	reg Registers

	/* In the Z80 CPU, there is
	   an interrupt enable flip-flop (IFF) that is set or reset by the programmer using the Enable
	   Interrupt (EI) and Disable Interrupt (DI) instructions. When the IFF is reset, an interrupt
	   cannot be accepted by the CPU.
	*/
	iff1          bool
	iff2          bool
	interruptMode int
}

func New(memSize uint16) *Zog {
	z := &Zog{
		mem: NewMemory(memSize),
	}
	z.Clear()
	return z
}

func (z *Zog) Run(a *Assembly) error {
	err := z.load(a)
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

func (z *Zog) LoadBytes(addr uint16, buf []byte) error {
	err := z.mem.Copy(addr, buf)
	if err != nil {
		return err
	}
	return nil
}

func (z *Zog) Clear() {
	z.mem.Clear()
	z.reg = Registers{}
	// 64KB will give zero here, correctly
	z.reg.SP = uint16(z.mem.Len())
	z.iff1 = false
	z.iff2 = false
}

func (z *Zog) load(a *Assembly) error {
	buf, err := a.Encode()
	if err != nil {
		return err
	}
	return z.LoadBytes(a.BaseAddr, buf)
}

// F flag register:
// S Z X H  X P/V N C
type flag int

const (
	F_C flag = iota
	F_N
	F_PV
	F_X1
	F_H
	F_X2
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
	case F_X1:
		return "X1"
	case F_H:
		return "H"
	case F_X2:
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

var ErrHalted = errors.New("HALT called")

// Implement io.Reader
func (z *Zog) Read(buf []byte) (int, error) {
	if len(buf) != 1 {
		panic("Non-byte read")
	}
	n, err := z.mem.Peek(z.reg.PC)
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
	z.iff1 = false
	z.iff2 = false
	return nil
}

func (z *Zog) ei() error {
	z.iff1 = true
	z.iff2 = true
	return nil
}

func (z *Zog) im(mode int) error {
	if mode != 0 && mode != 1 && mode != 2 {
		panic(fmt.Sprintf("Invalid interrupt mode: %d", mode))
	}
	z.interruptMode = mode
	return nil
}

func (z *Zog) push(nn uint16) {
	z.reg.SP--
	z.reg.SP--
	err := z.mem.Poke16(z.reg.SP, nn)
	if err != nil {
		panic(fmt.Sprintf("Can't write to SP [%04X]: %s", z.reg.SP, err))
	}
}

func (z *Zog) pop() uint16 {
	nn, err := z.mem.Peek16(z.reg.SP)
	if err != nil {
		panic(fmt.Sprintf("Can't write to SP [%04X]: %s", z.reg.SP, err))
	}
	z.reg.SP++
	z.reg.SP++
	return nn
}

func (z *Zog) out(port uint16, n byte) {
	fmt.Printf("OUT: [%04X] %02X\n", port, n)
}

func (z *Zog) in(port uint16) byte {
	n := byte(0)
	fmt.Printf("IN: [%04X] %02X\n", port, n)
	return n
}

func (z *Zog) execute(addr uint16) error {

	var err error
	var inst Instruction

	z.reg.PC = addr

EXECUTING:
	for {
		inst, err = DecodeOne(z)
		if err != nil {
			fmt.Printf("Error decoding: %s\n", err)
			break EXECUTING
		}
		fmt.Printf("I: %s\n", inst)
		err = inst.Execute(z)
		if err != nil {
			// Error handling after the loop
			break EXECUTING
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
