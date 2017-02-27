package zog

import "fmt"

type Registers struct {
	A, F byte
	B, C byte
	D, E byte
	H, L byte

	SP uint16
	PC uint16

	IX uint16
	IY uint16
}

type Zog struct {
	mem []byte
	reg Registers
}

func New(memSize uint16) *Zog {
	return &Zog{
		mem: make([]byte, memSize),
	}
}

func (z *Zog) Peek(addr uint16) (byte, error) {
	if int(addr) >= len(z.mem) {
		return 0, fmt.Errorf("Out of bounds memory read: %d", addr)
	}
	return z.mem[addr], nil
}

func (z *Zog) Poke(addr uint16, n byte) error {
	if int(addr) >= len(z.mem) {
		return fmt.Errorf("Out of bounds memory write: %d (%d)", addr, n)
	}
	z.mem[addr] = n
	return nil
}

func (z *Zog) Read8(l Location8) byte {
	return l.Read8(z)
}

func (z *Zog) Write8(l Location8, n byte) {
	l.Write8(z, n)
	return
}

func (z *Zog) Read16(l Location16) uint16 {
	return l.Read16(z)
}

func (z *Zog) Write16(l Location16, nn uint16) {
	l.Write16(z, nn)
	return
}

type Location8 interface {
	Read8(z *Zog) byte
	Write8(z *Zog, n byte)
	String() string
}

type R8Loc int

const (
	B R8Loc = iota
	C
	D
	E
	A
	F
	H
	L
)

func (l R8Loc) Read8(z *Zog) byte {
	switch l {
	case B:
		return z.reg.B
	case C:
		return z.reg.C

	case D:
		return z.reg.D
	case E:
		return z.reg.E

	case A:
		return z.reg.A
	case F:
		return z.reg.F

	case H:
		return z.reg.H
	case L:
		return z.reg.L

	default:
		panic(fmt.Errorf("Unrecognised R8 Location: %d", int(l)))
	}
}

func (l R8Loc) Write8(z *Zog, n byte) {
	switch l {
	case B:
		z.reg.B = n
	case C:
		z.reg.C = n

	case D:
		z.reg.D = n
	case E:
		z.reg.E = n

	case A:
		z.reg.A = n
	case F:
		z.reg.F = n

	case H:
		z.reg.H = n
	case L:
		z.reg.L = n

	default:
		panic(fmt.Errorf("Unrecognised R8 Location: %d", int(l)))
	}
}

func (l R8Loc) String() string {
	switch l {
	case B:
		return "B"
	case C:
		return "C"

	case D:
		return "D"
	case E:
		return "E"

	case A:
		return "A"
	case F:
		return "F"

	case H:
		return "H"
	case L:
		return "L"

	default:
		panic(fmt.Errorf("Unrecognised R8 Location: %d", int(l)))
	}
}

type R16Loc int

const (
	AF R16Loc = iota
	BC
	DE
	HL
	SP
	IX
	IY
)

type Location16 interface {
	Read16(z *Zog) uint16
	Write16(z *Zog, nn uint16)
	String() string
}

func (l R16Loc) Read16(z *Zog) uint16 {
	combine := func(h, l R8Loc) uint16 {
		hi := z.Read8(h)
		lo := z.Read8(l)
		return uint16(hi)<<8 | uint16(lo)
	}
	switch l {
	case BC:
		return combine(B, C)
	case DE:
		return combine(D, E)
	case AF:
		return combine(A, F)
	case HL:
		return combine(H, L)

	case SP:
		return z.reg.SP
	case IX:
		return z.reg.IX
	case IY:
		return z.reg.IY

	default:
		panic(fmt.Errorf("Unrecognised R16 Location: %d", int(l)))
	}
}

func (l R16Loc) Write16(z *Zog, nn uint16) {
	split := func(h, l R8Loc, nn uint16) {
		hi := byte(nn >> 8)
		h.Write8(z, hi)
		lo := byte(nn)
		l.Write8(z, lo)
	}
	switch l {
	case BC:
		split(B, C, nn)
		return
	case DE:
		split(D, E, nn)
		return
	case AF:
		split(A, F, nn)
		return
	case HL:
		split(H, L, nn)
		return

	case SP:
		z.reg.SP = nn
	case IX:
		z.reg.IX = nn
	case IY:
		z.reg.IY = nn

	default:
		panic(fmt.Errorf("Unrecognised R16 Location: %d", int(l)))
	}
}

func (l R16Loc) String() string {
	switch l {
	case BC:
		return "BC"
	case DE:
		return "DE"
	case AF:
		return "AF"
	case HL:
		return "HL"

	case SP:
		return "SP"
	case IX:
		return "IX"
	case IY:
		return "IY"

	default:
		panic(fmt.Errorf("Unrecognised R16 Location: %d", int(l)))
	}
}

type Instruction interface {
	Execute(z *Zog) error
}

func Decode(n byte) (Instruction, error) {
	return nil, fmt.Errorf("Failed to decode: %d", n)
}

func (z *Zog) Run() error {
	for {
		n, err := z.Peek(z.reg.PC)
		if err != nil {
			return err
		}

		// TODO - decoder with state for multiple-byte instructions
		i, err := Decode(n)
		if err != nil {
			return err
		}
		fmt.Printf("I: %s\n", i)
		i.Execute(z)
		z.reg.PC++
	}
}
