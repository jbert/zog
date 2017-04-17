package zog

import (
	"fmt"
	"strings"
)

type Src16 interface {
	Read16(z *Zog) (uint16, error)
	String() string
}
type Dst16 interface {
	Write16(z *Zog, nn uint16) error
	String() string
}

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
	// TODO: debug
	var n byte
	fmt.Printf("Z: %02X <- %s\n", n, r)
	return n, nil
}
func (r R8) Write8(z *Zog, n byte) error {
	// TODO: debug
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
	AF_PRIME
)

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
	// TODO: debug
	var nn uint16
	fmt.Printf("Z: %02X <- %s\n", nn, r)
	return nn, nil
}
func (r R16) Write16(z *Zog, nn uint16) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", r, nn)
	return nil
}

type Contents struct {
	addr Src16
}

func (c Contents) String() string {
	return fmt.Sprintf("(%s)", c.addr)
}
func (c Contents) Read8(z *Zog) (byte, error) {
	// TODO: debug
	var n byte
	fmt.Printf("Z: %02X <- %s\n", n, c.addr)
	return n, nil
}
func (c Contents) Write8(z *Zog, n byte) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", c.addr, n)
	return nil
}

func (c Contents) Read16(z *Zog) (uint16, error) {
	// TODO: debug
	var nn uint16
	fmt.Printf("Z: %04X <- %s\n", nn, c.addr)
	return nn, nil
}
func (c Contents) Write16(z *Zog, nn uint16) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %04X\n", c.addr, nn)
	return nil
}

type IndexedContents struct {
	addr Src16
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

type Imm16 uint16

func (nn Imm16) String() string {
	return fmt.Sprintf("0x%04X", uint16(nn))
}
func (nn Imm16) Read16(z *Zog) (uint16, error) {
	return uint16(nn), nil
}

type Imm8 byte

func (n Imm8) String() string {
	return fmt.Sprintf("0x%02X", byte(n))
}
func (n Imm8) Read8(z *Zog) (byte, error) {
	return byte(n), nil
}
func (n Imm8) Write8(z *Zog, tmp byte) error {
	panic("Attempt to write to immediate 8bit value")
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
