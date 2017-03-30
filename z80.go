package zog

import (
	"fmt"
	"strconv"
)

type Current struct {
	dst8 Dst8
	src8 Src8
	r8   R8
	r16  R16

	dst16 Dst16
	src16 Src16
	inst  Instruction

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

func (c *Current) N(s string) {
	// TODO: suffix h -> prefix 0x
	n, err := strconv.ParseInt(s, 16, 8)
	if err != nil {
		panic(fmt.Errorf("Invalid byte: %s", s))
	}
	c.src8 = Imm8(n)
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
	*c = Current{insts: c.insts}
}
