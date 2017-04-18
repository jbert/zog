package zog

import (
	"fmt"
	"strconv"
	"strings"
)

type Current struct {
	loc8 Loc8
	dst8 Loc8
	src8 Loc8

	loc16 Loc16
	dst16 Loc16
	src16 Loc16

	r8          R8
	r16         R16
	r16Contents Loc
	odigit      byte
	disp        Disp
	cc          Conditional
	nn          Loc16
	n           Loc8

	inst Instruction

	insts []Instruction
}

func (c *Current) GetInstructions() []Instruction {
	return c.insts
}

func (c *Current) LD8() {
	c.inst = NewLD8(c.dst8, c.src8)
}

func (c *Current) LD16() {
	c.inst = NewLD16(c.dst16, c.src16)
}
func (c *Current) Push() {
	c.inst = NewPUSH(c.src16)
}
func (c *Current) Pop() {
	c.inst = NewPOP(c.dst16)
}
func (c *Current) Ex() {
	c.inst = NewEX(c.dst16, c.src16)
}

func (c *Current) Inc8() {
	c.inst = NewINC8(c.loc8)
}
func (c *Current) Inc16() {
	c.inst = NewINC16(c.loc16)
}

func (c *Current) Dec8() {
	c.inst = NewDEC8(c.loc8)
}
func (c *Current) Dec16() {
	c.inst = NewDEC16(c.loc16)
}

func (c *Current) Add16() {
	c.inst = NewADD16(c.dst16, c.src16)
}

func (c *Current) Adc16() {
	c.inst = &ADC16{c.dst16, c.src16}
}

func (c *Current) Sbc16() {
	c.inst = &SBC16{c.dst16, c.src16}
}

func (c *Current) Accum(name string) {
	c.inst = NewAccum(name, c.src8)
}

func (c *Current) Rot(name string) {
	c.inst = NewRot(name, c.loc8)
}

func (c *Current) Bit() {
	c.inst = NewBIT(c.odigit, c.loc8)
}

func (c *Current) Res() {
	c.inst = NewRES(c.odigit, c.loc8)
}

func (c *Current) Set() {
	c.inst = NewSET(c.odigit, c.loc8)
}

func (c *Current) Simple(name string) {
	c.inst = LookupSimpleName(name)
}

func (c *Current) EDSimple(name string) {
	c.inst = LookupEDSimpleName(name)
}

func (c *Current) Rst() {
	// n is in  must be a byte
	imm8, ok := c.n.(Imm8)
	if !ok {
		panic("RST without immediate byte")
	}
	c.inst = &RST{byte(imm8)}
}

func (c *Current) Call() {
	c.inst = NewCALL(c.cc, c.src16)
}

func (c *Current) Ret() {
	c.inst = &RET{c.cc}
}

func (c *Current) Jp() {
	c.inst = NewJP(c.cc, c.src16)
}

func (c *Current) Jr() {
	c.inst = &JR{c.cc, c.disp}
}

func (c *Current) Djnz() {
	c.inst = &DJNZ{c.disp}
}

func (c *Current) In() {
	port := c.n
	if port == nil {
		port = C
	}
	c.inst = &IN{dst: c.r8, port: port}
}

func (c *Current) Out() {
	port := c.n
	if port == nil {
		port = C
	}
	c.inst = &OUT{port: port, value: c.r8}
}

func (c *Current) Nhex(s string) {
	n, err := strconv.ParseUint(s, 16, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid byte: %s", s))
	}
	c.n = Imm8(n)
}

func (c *Current) NNhex(s string) {
	nn, err := strconv.ParseUint(s, 16, 16)
	if err != nil {
		panic(fmt.Errorf("Invalid byte: %s", s))
	}
	c.nn = Imm16(nn)
}

func (c *Current) NNContents() {
	c.nn = Contents{c.nn}
}

func (c *Current) Ndec(s string) {
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid byte: %s", s))
	}
	c.n = Imm8(n)
}

func (c *Current) ODigit(s string) {
	n, err := strconv.ParseUint(s, 8, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid octal digit: %s", s))
	}
	c.odigit = byte(n)
}

func (c *Current) Disp0xHex(s string) {
	s = strings.Replace(s, "0x", "", 1)
	n, err := strconv.ParseInt(s, 16, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid signed decimal byte: %s", s))
	}
	c.disp = Disp(int8(n))
}

func (c *Current) DispHex(s string) {
	n, err := strconv.ParseInt(s, 16, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid signed decimal byte: %s", s))
	}
	c.disp = Disp(int8(n))
}

func (c *Current) DispDecimal(s string) {
	n, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid signed decimal byte: %s", s))
	}
	c.disp = Disp(int8(n))
}

func (c *Current) Conditional(cc Conditional) {
	c.cc = cc
}

func (c *Current) Loc8() {
	if c.loc8 != nil {
		return
	}
	if c.r16Contents != nil {
		c.loc8 = c.r16Contents
		c.r16Contents = nil
		return
	}
	c.loc8 = c.r8
}

func (c *Current) Dst8() {
	if c.dst8 != nil {
		return
	}
	if c.nn != nil {
		nn_contents, ok := c.nn.(Loc8)
		if !ok {
			panic("Dst8 set to NN but not contents")
		}
		c.dst8 = nn_contents
		c.nn = nil
		return
	}
	if c.r16Contents != nil {
		c.dst8 = c.r16Contents
		c.r16Contents = nil
		return
	}
	c.dst8 = c.r8
}

func (c *Current) Src8() {
	if c.src8 != nil {
		return
	}
	if c.n != nil {
		c.src8 = c.n
		c.n = nil
		return
	}
	if c.nn != nil {
		nn_contents, ok := c.nn.(Loc8)
		if !ok {
			panic("Dst16 set to NN but not contents")
		}
		c.src8 = nn_contents
		c.nn = nil
		return
	}
	if c.r16Contents != nil {
		c.src8 = c.r16Contents
		c.r16Contents = nil
		return
	}
	c.src8 = c.r8
}

func (c *Current) Loc16() {
	if c.loc16 != nil {
		return
	}
	c.loc16 = c.r16
}

func (c *Current) Dst16() {
	if c.nn != nil {
		nn_contents, ok := c.nn.(Loc16)
		if !ok {
			panic("Dst16 set to NN but not contents")
		}
		c.dst16 = nn_contents
		c.nn = nil
		return
	}
	if c.r16Contents != nil {
		c.dst16 = c.r16Contents
		c.r16Contents = nil
		return
	}
	c.dst16 = c.r16
}

func (c *Current) Src16() {
	if c.nn != nil {
		c.src16 = c.nn
		c.nn = nil
		return
	}
	if c.r16Contents != nil {
		c.src16 = c.r16Contents
		c.r16Contents = nil
		return
	}
	c.src16 = c.r16
}

func (c *Current) R16Contents() {
	c.r16Contents = Contents{c.r16}
}

func (c *Current) IR16Contents() {
	c.r16Contents = IndexedContents{c.r16, c.disp}
}

func (c *Current) R8(s string) {
	c.r8 = LookupR8Name(s)
}

func (c *Current) R16(s string) {
	c.r16 = LookupR16Name(s)
}

func (c *Current) Emit() {
	c.insts = append(c.insts, c.inst)
	c.clean()
}

func (c *Current) clean() {
	*c = Current{insts: c.insts, cc: True}
}
