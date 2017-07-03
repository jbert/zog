package zog

import (
	"errors"
	"fmt"
)

type Zog struct {
	mem *Memory
	reg Registers
}

func New(memSize uint16) *Zog {
	z := &Zog{
		mem: NewMemory(memSize),
	}
	z.Clear()
	return z
}

func (z *Zog) Clear() {
	z.mem.Clear()
	z.reg = Registers{}
}

func (z *Zog) Load(a *Assembly) error {
	buf, err := a.Encode()
	if err != nil {
		return err
	}
	err = z.mem.Copy(a.BaseAddr, buf)
	if err != nil {
		return err
	}
	return nil
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

func (z *Zog) Run(a *Assembly) error {
	err := z.Load(a)
	if err != nil {
		return err
	}
	return z.Execute(a.BaseAddr)
}

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

func (z *Zog) Execute(addr uint16) error {

	var err error
	var inst Instruction
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
