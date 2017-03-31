package zog

import (
	"fmt"
	"strconv"
)

type Current struct {
	loc8 Loc8
	dst8 Dst8
	src8 Src8

	loc16 Loc16
	dst16 Dst16
	src16 Src16

	r8         R8
	r16        R16
	odigit     byte
	signedByte int8
	cc         Conditional

	inst Instruction

	insts []Instruction
}

func (c *Current) GetInstructions() []Instruction {
	return c.insts
}

func (c *Current) LD8() {
	c.inst = &LD8{c.dst8, c.src8}
}

func (c *Current) LD16() {
	c.inst = &LD16{c.dst16, c.src16}
}
func (c *Current) Push() {
	c.inst = &PUSH{c.src16}
}
func (c *Current) Pop() {
	c.inst = &POP{c.dst16}
}
func (c *Current) Ex() {
	c.inst = &EX{c.dst16, c.src16}
}

func (c *Current) Inc8() {
	c.inst = &INC8{c.loc8}
}
func (c *Current) Inc16() {
	c.inst = &INC16{c.loc16}
}

func (c *Current) Dec8() {
	c.inst = &DEC8{c.loc8}
}
func (c *Current) Dec16() {
	c.inst = &DEC16{c.loc16}
}

func (c *Current) Add16() {
	c.inst = &ADD16{c.dst16, c.src16}
}

func (c *Current) Accum(name string) {
	c.inst = NewAccum(name, c.src8)
}

func (c *Current) Rot(name string) {
	c.inst = NewRot(name, c.loc8)
}

func (c *Current) Bit() {
	c.inst = &BIT{num: c.odigit, r: c.loc8}
}

func (c *Current) Res() {
	c.inst = &RES{num: c.odigit, r: c.loc8}
}

func (c *Current) Set() {
	c.inst = &SET{num: c.odigit, r: c.loc8}
}

func (c *Current) Simple(name string) {
	c.inst = LookupSimpleName(name)
}

func (c *Current) Rst() {
	// n is in src8 and must be a byte
	imm8, ok := c.src8.(Imm8)
	if !ok {
		panic("RST without immediate byte")
	}
	c.inst = &RST{byte(imm8)}
}

func (c *Current) Call() {
	c.inst = &CALL{c.cc, c.src16}
}

func (c *Current) Ret() {
	c.inst = &RET{c.cc}
}

func (c *Current) Jp() {
	c.inst = &JP{c.cc, c.src16}
}

func (c *Current) Jr() {
	c.inst = &JR{c.cc, Disp(c.signedByte)}
}

func (c *Current) Djnz() {
	c.inst = &DJNZ{Disp(c.signedByte)}
}

func (c *Current) In() {
	c.inst = &IN{dst: c.dst8, port: c.src8}
}

func (c *Current) OutC() {
	c.inst = &OUT{port: C, value: c.src8}
}

func (c *Current) OutN() {
	c.inst = &OUT{port: C, value: c.src8}
}

func (c *Current) Nhex(s string) {
	n, err := strconv.ParseUint(s, 16, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid byte: %s", s))
	}
	c.src8 = Imm8(n)
}

func (c *Current) NNhex(s string) {
	nn, err := strconv.ParseUint(s, 16, 16)
	if err != nil {
		panic(fmt.Errorf("Invalid byte: %s", s))
	}
	c.src16 = Imm16(nn)
}

func (c *Current) NNContents() {
	c.src16 = Contents{c.src16}
}

func (c *Current) Ndec(s string) {
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid byte: %s", s))
	}
	c.src8 = Imm8(n)
}

func (c *Current) ODigit(s string) {
	n, err := strconv.ParseUint(s, 8, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid octal digit: %s", s))
	}
	c.odigit = byte(n)
}

func (c *Current) SignedDecimalByte(s string) {
	n, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid signed decimal byte: %s", s))
	}
	c.signedByte = int8(n)
}

func (c *Current) Conditional(cc Conditional) {
	c.cc = cc
}

func (c *Current) Loc8() {
	if c.loc8 != nil {
		return
	}
	c.loc8 = c.r8
}

func (c *Current) Dst8() {
	if c.dst8 != nil {
		return
	}
	c.dst8 = c.r8
}

func (c *Current) Src8() {
	if c.src8 != nil {
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
	if c.dst16 != nil {
		return
	}
	c.dst16 = c.r16
}

func (c *Current) Src16() {
	if c.src16 != nil {
		return
	}
	c.src16 = c.r16
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
