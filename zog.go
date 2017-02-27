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

// Order and numbering is important here, e.g. see decodeLD8

// LD src		B C D E F L (HL) A
// LD dst		B C D E H L (HL) A
// ADD src	B C D E F L (HL) A
const (
	B R8Loc = iota
	C
	D
	E
	F
	L
	HL_PLACE_HOLDER
	A
	H
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

	case F:
		return z.reg.F
	case H:
		return z.reg.H
	case L:
		return z.reg.L

	case A:
		return z.reg.A

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

	case F:
		z.reg.F = n
	case H:
		z.reg.H = n
	case L:
		z.reg.L = n

	case A:
		z.reg.A = n

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

	case F:
		return "F"
	case H:
		return "H"
	case L:
		return "L"

	case A:
		return "A"

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
		hi := z.reg.H
		lo := z.reg.L
		return uint16(hi)<<8 | uint16(lo)

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
		hi := byte(nn >> 8)
		z.reg.H = hi
		lo := byte(nn)
		z.reg.L = lo
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
	String() string
}

type ILD8 struct {
	src, dst R8Loc
}

func (ld *ILD8) String() string {
	return fmt.Sprintf("LD %s, %s", ld.dst, ld.src)
}
func (ld *ILD8) Execute(z *Zog) error {
	n := z.Read8(ld.src)
	z.Write8(ld.dst, n)
	return nil
}

func decodeLD8(n byte) (*ILD8, error) {
	low3bits := n & 0x07
	if low3bits == 6 {
		panic("(HL) not yet implemented")
	}
	src := R8Loc(low3bits)

	next3 := (n & 0x38) >> 3
	if next3 == 6 {
		panic("(HL) not yet implemented")
	}
	dst := R8Loc(next3)
	if next3 == 4 {
		// Can read F, but not write to it
		dst = H
	}
	return &ILD8{src: src, dst: dst}, nil
}

func Decode(n byte) (Instruction, error) {
	if n >= 0x40 && n <= 0x7f {
		// Main part of 8bit load group
		return decodeLD8(n)
	}
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
