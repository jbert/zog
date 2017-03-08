package zog

import (
	"fmt"
)

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
	flags := z.Read8(F)
	if new {
		flags = flags | mask
	} else {
		mask = ^mask
		flags = flags & mask
	}
	z.Write8(F, flags)
}
func (z *Zog) GetFlag(f flag) bool {
	mask := byte(1) << uint(f)
	flags := z.Read8(F)
	flag := flags & mask
	return flag != 0
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
	HL_CONTENTS
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
	case HL_CONTENTS:
		addr := z.Read16(HL)
		n, err := z.Peek(addr)
		if err != nil {
			panic(fmt.Sprintf("Failed to fetch (HL) [%04X]: %s", addr, err))
		}
		return n

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
	case HL_CONTENTS:
		addr := z.Read16(HL)
		err := z.Poke(addr, n)
		if err != nil {
			panic(fmt.Errorf("Can't write to addr [%04x]: %s", addr, err))
		}

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
	case HL_CONTENTS:
		return "(HL)"

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
	Encode() []byte
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
func (ld *ILD8) Encode() []byte {
	hi3 := byte(ld.dst)
	// Can read F, but not write to it
	if ld.dst == H {
		hi3 = byte(4)
	}
	lo3 := byte(ld.src)
	top2 := byte(1)
	return []byte{top2<<6 | hi3<<3 | lo3}
}

func decodeLD8(hi3, lo3 byte) (*ILD8, error) {
	src := R8Loc(lo3)
	dst := R8Loc(hi3)
	if hi3 == 4 {
		// Can read F, but not write to it
		dst = H
	}
	return &ILD8{src: src, dst: dst}, nil
}

type ILD8Immediate struct {
	dst R8Loc
	n   byte
}

func (ld *ILD8Immediate) String() string {
	return fmt.Sprintf("LD %s, 0x%X", ld.dst, ld.n)
}
func (ld *ILD8Immediate) Execute(z *Zog) error {
	z.Write8(ld.dst, ld.n)
	return nil
}
func (ld *ILD8Immediate) Encode() []byte {
	hi3 := byte(ld.dst)
	lo3 := byte(6)
	top2 := byte(0)
	return []byte{top2<<6 | hi3<<3 | lo3, ld.n}
}

func decodeLD8Immediate(hi3 byte, getNext func() (byte, error)) (Instruction, error) {
	dst := R8Loc(hi3)
	n, err := getNext()
	if err != nil {
		return nil, err
	}
	return &ILD8Immediate{dst: dst, n: n}, nil
}

type IAccumOp struct {
	src  R8Loc
	op   func(z *Zog, a, n byte) error
	name string
}

func (i *IAccumOp) String() string {
	return fmt.Sprintf("%s A, %s", i.name, i.src)
}
func (i *IAccumOp) Execute(z *Zog) error {
	a := z.Read8(A)
	n := z.Read8(i.src)
	i.op(z, a, n)
	return nil
}
func (i *IAccumOp) Encode() []byte {
	top2 := byte(2)
	lo3 := byte(i.src)
	var hi3 byte
	// TODO - drive this from the same table used in the decode logic
	switch i.name {
	case "ADD":
		hi3 = 0
	case "ADC":
		hi3 = 1
	case "SUB":
		hi3 = 2
	case "SBC":
		hi3 = 3
	case "AND":
		hi3 = 4
	case "XOR":
		hi3 = 5
	case "OR":
		hi3 = 6
	case "CP":
		hi3 = 7
	default:
		panic("ack")
	}
	return []byte{top2<<6 | hi3<<3 | lo3}
}

func AccumAdd(z *Zog, a, n byte) error {
	ret := int(a) + int(n)
	// TODO: consider other flags, add helpers
	z.SetFlag(F_C, ret > 0xff)
	z.Write8(A, byte(ret))
	return nil
}

func AccumAdc(z *Zog, a, n byte) error {
	ret := int(a) + int(n)
	carry := z.GetFlag(F_C)
	if carry {
		ret++
	}
	// TODO: consider other flags, add helpers
	z.SetFlag(F_C, ret > 0xff)
	z.Write8(A, byte(ret))
	return nil
}

func AccumSub(z *Zog, a, n byte) error {
	ret := int(a) - int(n)
	// TODO: consider other flags, add helpers
	z.SetFlag(F_C, ret < 0x00)
	z.Write8(A, byte(ret))
	return nil
}

func AccumSbc(z *Zog, a, n byte) error {
	ret := int(a) - int(n)
	carry := z.GetFlag(F_C)
	if carry {
		ret--
	}
	// TODO: consider other flags, add helpers
	z.SetFlag(F_C, ret < 0x00)
	z.Write8(A, byte(ret))
	return nil
}

func AccumAnd(z *Zog, a, n byte) error {
	ret := a & n
	z.SetFlag(F_C, false)
	z.Write8(A, ret)
	return nil
}

func AccumXor(z *Zog, a, n byte) error {
	ret := a ^ n
	z.SetFlag(F_C, false)
	z.Write8(A, ret)
	return nil
}

func AccumOr(z *Zog, a, n byte) error {
	ret := a | n
	z.SetFlag(F_C, false)
	z.Write8(A, ret)
	return nil
}

func AccumCp(z *Zog, a, n byte) error {
	ret := int(a) - int(n)
	// TODO: consider other flags, add helpers
	z.SetFlag(F_C, ret < 0x00)
	return nil
}

func decodeAccumOp(hi3, lo3 byte) (Instruction, error) {
	// Arithmetic and logical with accumulator.
	ops := []struct {
		name string
		op   func(z *Zog, a, n byte) error
	}{
		{"ADD", AccumAdd},
		{"ADC", AccumAdc},
		{"SUB", AccumSub},
		{"SBC", AccumSbc},
		{"AND", AccumAnd},
		{"XOR", AccumXor},
		{"OR", AccumOr},
		{"CP", AccumCp},
	}

	src := R8Loc(lo3)
	return &IAccumOp{src: src, name: ops[hi3].name, op: ops[hi3].op}, nil
}

// Gaps in this table are handled programatically
type simpleSingle struct {
	encoding byte
	i        Instruction
	name     string
}

var simpleSingles []simpleSingle = []simpleSingle{
	{0x37, I_SCF, "SCF"},
	{0x3f, I_CCF, "CCF"},

	{0x76, I_HALT, "HALT"},
}

func findSimpleInstructionByEncoding(n byte) (Instruction, bool) {
	for _, ss := range simpleSingles {
		if ss.encoding == n {
			return ss.i, true
		}
	}
	return nil, false
}
func findSimpleByInstruction(i ISimple) (simpleSingle, bool) {
	for _, ss := range simpleSingles {
		if ss.i == i {
			return ss, true
		}
	}
	return simpleSingle{}, false
}

type ISimple int

const (
	I_HALT ISimple = iota
	I_SCF
	I_CCF
)

func (i ISimple) String() string {
	ss, ok := findSimpleByInstruction(i)
	if !ok {
		panic(fmt.Sprintf("Unrecognised simple instruction: %v", i))
	}

	return ss.name
}

func (i ISimple) Execute(z *Zog) error {
	switch i {
	case I_HALT:
		panic("Attempt to execute HALT")

	case I_SCF:
		z.SetFlag(F_C, true)
	case I_CCF:
		f := z.GetFlag(F_C)
		z.SetFlag(F_C, !f)

	default:
		panic("Unrecognised various instruction")
	}
	return nil
}
func (i ISimple) Encode() []byte {
	ss, ok := findSimpleByInstruction(i)
	if !ok {
		panic(fmt.Sprintf("Unrecognised simple instruction: %v", i))
	}

	return []byte{ss.encoding}
}

func Decode(getNext func() (byte, error)) (Instruction, error) {
	var n byte
	var err error
	for {
		n, err = getNext()
		if err != nil {
			return nil, err
		}

		// Table lookup has precedence
		i, ok := findSimpleInstructionByEncoding(n)
		if ok {
			return i, nil
		}

		lo3 := n & 0x07
		hi3 := (n & 0x38) >> 3
		top2 := (n & 0xc0) >> 6

		// fmt.Printf("top2 %x, hi3 %x, lo3 %x\n", top2, hi3, lo3)

		switch top2 {

		case 0:
			switch lo3 {
			case 6:
				return decodeLD8Immediate(hi3, getNext)
			default:
				panic(fmt.Sprintf("Failed to decode top0 instruction: 0x%02X", n))
			}

		case 1:
			// Main part of 8bit load group

			// In place of LD (HL), (HL) we have HALT
			if n == 0x76 {
				return I_HALT, nil
			}
			return decodeLD8(hi3, lo3)

		case 2:
			return decodeAccumOp(hi3, lo3)

		default:
			panic(fmt.Sprintf("Failed to decode instruction: 0x%02X", n))
		}
	}
	return nil, fmt.Errorf("Failed to decode: %02x", n)
}

func (z *Zog) Run() (byte, error) {
	getNext := func() (byte, error) {
		n, err := z.Peek(z.reg.PC)
		//		fmt.Printf("PC: %04X %02X\n", z.reg.PC, n)
		z.reg.PC++
		return n, err
	}

	for {
		// TODO - decoder with state for multiple-byte instructions
		i, err := Decode(getNext)
		if err != nil {
			return 0, err
		}
		fmt.Printf("I: %s\n", i)
		if v, ok := i.(ISimple); ok && v == I_HALT {
			// HALT instruction - return A reg to caller
			return z.reg.A, nil
		}
		i.Execute(z)
	}
}

func (z *Zog) Encode(addr uint16, instructions []Instruction) error {
	// The i here is the instruction, not the loop var
	for _, i := range instructions {
		buf := i.Encode()
		for _, b := range buf {
			z.Poke(addr, b)
			addr++
		}
	}
	return nil
}

func (z *Zog) Execute(addr uint16, program string) (byte, error) {
	assembler := NewAssembler()

	instructions, err := assembler.Assemble(program)
	if err != nil {
		return 0, fmt.Errorf("Failed to assemble: %s", err)
	}
	err = z.Encode(addr, instructions)
	if err != nil {
		return 0, fmt.Errorf("Failed to encode: %s", err)
	}
	a, err := z.Run()
	if err != nil {
		return 0, fmt.Errorf("Failed to run: %s", err)
	}
	return a, nil
}
