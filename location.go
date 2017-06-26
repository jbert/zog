package zog

import (
	"fmt"
	"strings"
)

type Loc8 interface {
	Read8(z *Zog) (byte, error)
	Write8(z *Zog, n byte) error
	String() string
}

type Loc16 interface {
	Read16(z *Zog) (uint16, error)
	Write16(z *Zog, nn uint16) error
	String() string
}

// (HL), (BC), (SP) could refer to a byte addr or a word addr
type Loc interface {
	Read8(z *Zog) (byte, error)
	Write8(z *Zog, n byte) error
	Read16(z *Zog) (uint16, error)
	Write16(z *Zog, nn uint16) error
	String() string
}

type Registers struct {
	A, F byte
	B, C byte
	D, E byte
	H, L byte

	IXH, IXL byte
	IYH, IYL byte

	I, R byte

	// Alternate register set
	A_PRIME, F_PRIME byte
	B_PRIME, C_PRIME byte
	D_PRIME, E_PRIME byte
	H_PRIME, L_PRIME byte

	SP uint16
	PC uint16
}

func (r Registers) String() string {
	return fmt.Sprintf("AF %04X BC %04X DE %04X HL %04X AF' %04X BC' %04X DE' %04X HL' %04X SP %04X PC %04X IX %04X IY %04X",
		r.Read16(AF),
		r.Read16(BC),
		r.Read16(DE),
		r.Read16(HL),
		r.Read16(AF_PRIME),
		r.Read16(BC_PRIME),
		r.Read16(DE_PRIME),
		r.Read16(HL_PRIME),
		r.Read16(SP),
		r.Read16(PC),
		r.Read16(IX),
		r.Read16(IY))
}

type R8 int

const (
	A R8 = iota
	F
	B
	C
	D
	E
	H
	L
	IXL
	IXH
	IYL
	IYH
	I
	R
)

func (r *Registers) Read8(l R8) byte {
	switch l {
	case A:
		return r.A
	case F:
		return r.F
	case B:
		return r.B
	case C:
		return r.C
	case D:
		return r.D
	case E:
		return r.E
	case H:
		return r.H
	case L:
		return r.L
	case IXL:
		return r.IXL
	case IXH:
		return r.IXH
	case IYL:
		return r.IYL
	case IYH:
		return r.IYH
	case I:
		return r.I
	case R:
		return r.R
	default:
		panic(fmt.Sprintf("Unknown R8: %d", l))
	}
}

func (r *Registers) Write8(l R8, n byte) {
	switch l {
	case A:
		r.A = n
	case F:
		r.F = n
	case B:
		r.B = n
	case C:
		r.C = n
	case D:
		r.D = n
	case E:
		r.E = n
	case H:
		r.H = n
	case L:
		r.L = n
	case IXL:
		r.IXL = n
	case IXH:
		r.IXH = n
	case IYL:
		r.IYL = n
	case IYH:
		r.IYH = n
	case I:
		r.I = n
	case R:
		r.R = n
	default:
		panic(fmt.Sprintf("Unknown R8: %d", l))
	}
}

type r8name struct {
	r    R8
	name string
}

var R8Names []r8name = []r8name{
	{A, "A"},
	{F, "F"},

	{B, "B"},
	{C, "C"},

	{D, "D"},
	{E, "E"},

	{H, "H"},
	{L, "L"},

	{IXL, "IXL"},
	{IXH, "IXH"},
	{IYL, "IYL"},
	{IYH, "IYH"},

	{I, "I"},
	{R, "R"},
}

func LookupR8Name(name string) R8 {
	name = strings.ToUpper(name)
	for _, r8name := range R8Names {
		if r8name.name == name {
			return r8name.r
		}
	}
	panic(fmt.Errorf("Unrecognised R8 name : %s", name))
}

func (r R8) String() string {
	for _, r8name := range R8Names {
		if r8name.r == r {
			return r8name.name
		}
	}
	panic(fmt.Errorf("Unrecognised R8 : %d", int(r)))
}

func (r R8) Read8(z *Zog) (byte, error) {
	n := z.reg.Read8(r)
	fmt.Printf("Z: %02X <- %s\n", n, r)
	return n, nil
}
func (r R8) Write8(z *Zog, n byte) error {
	z.reg.Write8(r, n)
	fmt.Printf("Z: %s <- %02X\n", r, n)
	return nil
}

type R16 int

const (
	AF R16 = iota
	BC
	DE
	HL
	IX
	IY
	SP
	PC
	AF_PRIME
	BC_PRIME
	DE_PRIME
	HL_PRIME
)

func (r *Registers) Read16(l R16) uint16 {
	var lo, hi byte
	switch l {
	case AF:
		hi = r.A
		lo = r.F
	case BC:
		hi = r.B
		lo = r.C
	case DE:
		hi = r.D
		lo = r.E
	case HL:
		hi = r.H
		lo = r.L
	case IX:
		hi = r.IXH
		lo = r.IXL
	case IY:
		hi = r.IYH
		lo = r.IYL
	case SP:
		return r.SP
	case AF_PRIME:
		hi = r.A_PRIME
		lo = r.F_PRIME
	default:
		panic(fmt.Sprintf("Unknown R16: %d", l))
	}
	return (uint16(hi) << 8) | uint16(lo)
}
func (r *Registers) Write16(l R16, nn uint16) {
	hi := byte(nn >> 8)
	lo := byte(nn)
	switch l {
	case AF:
		r.A = hi
		r.F = lo
	case BC:
		r.B = hi
		r.C = lo
	case DE:
		r.D = hi
		r.E = lo
	case HL:
		r.H = hi
		r.L = lo
	case IX:
		r.IXH = hi
		r.IXL = lo
	case IY:
		r.IYH = hi
		r.IYL = lo
	case SP:
		r.SP = nn
	case AF_PRIME:
		r.A_PRIME = hi
		r.F_PRIME = lo
	default:
		panic(fmt.Sprintf("Unknown R16: %d", l))
	}
}

type r16name struct {
	r    R16
	name string
}

var R16Names []r16name = []r16name{
	{AF, "AF"},
	{AF_PRIME, "AF'"},
	{BC, "BC"},
	{DE, "DE"},
	{HL, "HL"},
	{IX, "IX"},
	{IY, "IY"},
	{SP, "SP"},
}

func LookupR16Name(name string) R16 {
	name = strings.ToUpper(name)
	for _, r16name := range R16Names {
		if r16name.name == name {
			return r16name.r
		}
	}
	panic(fmt.Errorf("Unrecognised R16 name : %s", name))
}

func (r R16) String() string {
	for _, r16name := range R16Names {
		if r16name.r == r {
			return r16name.name
		}
	}
	panic(fmt.Errorf("Unrecognised R16 : %d", int(r)))
}

func (r R16) Read16(z *Zog) (uint16, error) {
	nn := z.reg.Read16(r)
	fmt.Printf("Z: %04X <- %s\n", nn, r)
	return nn, nil
}
func (r R16) Write16(z *Zog, nn uint16) error {
	z.reg.Write16(r, nn)
	fmt.Printf("Z: %s <- %04X\n", r, nn)
	return nil
}

type Contents struct {
	addr Loc16
}

func (c Contents) String() string {
	return fmt.Sprintf("(%s)", c.addr)
}
func (c Contents) Read8(z *Zog) (byte, error) {
	addr, err := c.addr.Read16(z)
	if err != nil {
		return 0, fmt.Errorf("Can't get contents of [%s]: %s", c.addr, err)
	}
	n, err := z.mem.Peek(addr)
	if err != nil {
		return 0, fmt.Errorf("Can't read contents of [%s]: %s", c, err)
	}
	fmt.Printf("Z: %02X <- %s\n", n, c)
	return n, nil
}
func (c Contents) Write8(z *Zog, n byte) error {
	addr, err := c.addr.Read16(z)
	if err != nil {
		return fmt.Errorf("Can't get contents of [%s]: %s", c.addr, err)
	}
	err = z.mem.Poke(addr, n)
	if err != nil {
		return fmt.Errorf("Can't write contents of [%s]: %s", c, err)
	}
	fmt.Printf("Z: %s <- %02X\n", c, n)
	return nil
}

func (c Contents) Read16(z *Zog) (uint16, error) {
	addr, err := c.addr.Read16(z)
	if err != nil {
		return 0, fmt.Errorf("Can't get contents of [%s]: %s", c.addr, err)
	}
	nn, err := z.mem.Peek16(addr)
	if err != nil {
		return 0, fmt.Errorf("Can't read contents of [%s]: %s", c, err)
	}
	fmt.Printf("Z: %04X <- %s\n", nn, c)
	return nn, nil
}
func (c Contents) Write16(z *Zog, nn uint16) error {
	addr, err := c.addr.Read16(z)
	if err != nil {
		return fmt.Errorf("Can't get contents of [%s]: %s", c.addr, err)
	}
	err = z.mem.Poke16(addr, nn)
	if err != nil {
		return fmt.Errorf("Can't write contents of [%s]: %s", c, err)
	}
	fmt.Printf("Z: %s <- %04X\n", c, nn)
	return nil
}

type IndexedContents struct {
	addr Loc16
	d    Disp
}

func (ic IndexedContents) String() string {
	return fmt.Sprintf("(%s%+d)", ic.addr, int8(ic.d))
}
func (ic IndexedContents) Read8(z *Zog) (byte, error) {
	// TODO: debug
	var n byte
	fmt.Printf("Z: %02X <- %s\n", n, ic.addr)
	return n, nil
}
func (ic IndexedContents) Write8(z *Zog, n byte) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", ic.addr, n)
	return nil
}

func (ic IndexedContents) Read16(z *Zog) (uint16, error) {
	// TODO: debug
	var nn uint16
	fmt.Printf("Z: %04X <- %s\n", nn, ic.addr)
	return nn, nil
}
func (ic IndexedContents) Write16(z *Zog, nn uint16) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %04X\n", ic.addr, nn)
	return nil
}

type Label struct {
	// Default to the zero value because we need to encode
	// this location *before* it is resolved with an address
	// so that we can calculate instruction length.
	// See assemble/ResolveAddr
	Imm16
	name string
}

func (l Label) String() string {
	addrStr := l.Imm16.String()
	return fmt.Sprintf("%s (%s)", l.name, addrStr)
}

type Imm16 uint16

func (nn Imm16) String() string {
	return fmt.Sprintf("0x%04X", uint16(nn))
}
func (nn Imm16) Read16(z *Zog) (uint16, error) {
	return uint16(nn), nil
}
func (nn Imm16) Write16(z *Zog, n uint16) error {
	return fmt.Errorf("Attempt to write [%02X] to immediate 16bit value [%04X]", n, nn)
}

type Imm8 byte

func (n Imm8) String() string {
	return fmt.Sprintf("0x%02X", byte(n))
}
func (n Imm8) Read8(z *Zog) (byte, error) {
	return byte(n), nil
}
func (n Imm8) Write8(z *Zog, tmp byte) error {
	return fmt.Errorf("Attempt to write [%02X] to immediate 8bit value [%04X]", n, tmp)
}

type Conditional interface {
	String() string
}

type FlagTest int

const (
	FT_Z FlagTest = iota
	FT_C
	FT_PO
	FT_PE
	FT_P
	FT_M
)

func (ft FlagTest) String() string {
	switch ft {
	case FT_Z:
		return "Z"
	case FT_C:
		return "C"
	case FT_PO:
		return "PO"
	case FT_PE:
		return "PE"
	case FT_P:
		return "P"
	case FT_M:
		return "M"
	default:
		panic(fmt.Sprintf("Unknown flag test [%d]", int(ft)))
	}
}

type LogicConstant struct{}

var True LogicConstant

func (l LogicConstant) String() string {
	panic("Attempt to render 'true' as string")
}

type Not struct {
	ft FlagTest
}

func (n Not) String() string {
	return fmt.Sprintf("N%s", n.ft)
}

type Disp int8

func (d Disp) String() string {
	return fmt.Sprintf("%d", d)
}
