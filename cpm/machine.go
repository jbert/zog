package cpm

import (
	"fmt"
	"os"

	"github.com/jbert/zog"
)

type Machine struct {
	z *zog.Zog
}

func NewMachine(z *zog.Zog) *Machine {
	return &Machine{z: z}
}

func (m Machine) LoadAddr() uint16 {
	return 0x0100
}

func (m Machine) RunAddr() uint16 {
	return 0x0100
}

func (m Machine) Name() string {
	return "cpm"
}

var printAssembly = `
	; Function to call is in C
	; Func 2 => Print ASCII code of reg E to console
	; Func 9 => Print ASCII string starting at DE until $ to console
	LD A, 2
	CP C
	JP NZ, next1
	CALL printchar
	RET
next1:
	LD A, 9
	CP C
	JP NZ, next2
	CALL printstr
	RET
next2:
	HALT
; Print char in E to console
printchar:
	PUSH BC
	LD BC, 0ffffh
	OUT (C), E
	POP BC
	RET
; Print $-terminated string at DE to console
printstr:
	PUSH HL
	PUSH DE
  POP HL
	LD A, 24h		; '$'
printstr_nextchar:
	CP (HL)
	JP Z, printstr_end
	LD E, (HL)
	CALL printchar
	INC HL
	JP printstr_nextchar
printstr_end:
	POP HL
	RET
`

func (m *Machine) Stop() {
}

func (m *Machine) Start() error {
	m.z.RegisterOutputHandler(0xffff, printByte)
	zeroPageAssembly, err := zog.Assemble(`
	ORG 0000h
	HALT
	NOP			; would be addr of warm start (with JP inst at 0000)
	NOP
	NOP			; The 'intel standard iobyte'? http://www.gaby.de/cpm/manuals/archive/cpm22htm/ch6.htm#Section_6.9
	NOP
	; One entry point at 0005h
	; but this is also "the lowest address used by CP/M"
	; and used to the set the SP (by zexall)
	JP 0xf000
`)
	if err != nil {
		return fmt.Errorf("Failed to assemble prelude: %s", err)
	}
	err = m.z.Load(zeroPageAssembly)
	if err != nil {
		return fmt.Errorf("Load zero page assembly: %s", err)
	}

	highAssembly, err := zog.Assemble("ORG 0xf000\n" + printAssembly)

	if err != nil {
		return fmt.Errorf("Failed to assemble prelude: %s", err)
	}
	return m.z.Load(highAssembly)
}

func printByte(n byte) {
	fmt.Fprintf(os.Stderr, "%c", n)
}
