package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jbert/zog"
)

func main() {
	if len(os.Args) < 2 {
		usage("Missing filename")
	}
	fname := os.Args[1]

	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Failed to open file [%s] : %s\n", fname, err)
	}

	z := zog.New(0)
	err = loadPseudoCPM(z)
	if err != nil {
		log.Fatalf("Failed to load pCPM: %s", err)
	}
	loadAddr := uint16(0x0100)
	runAddr := uint16(0x0100)
	err = z.RunBytes(loadAddr, buf, runAddr)
	if err != nil {
		log.Fatalf("RunBytes returned error: %s", err)
	}
}

func printByte(n byte) {
	fmt.Fprintf(os.Stderr, "%c", n)
}

func loadPseudoCPM(z *zog.Zog) error {
	z.RegisterOutputHandler(0xffff, printByte)
	assembly, err := zog.Assemble(`
	ORG 0000h
	HALT
	NOP
	NOP
	NOP
	NOP
	; One entry point at 0005h.
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
			`)

	if err != nil {
		return fmt.Errorf("Failed to assemble prelude: %s", err)
	}
	return z.Load(assembly)
}

func usage(reason string) {
	fmt.Printf(`%s

%s <filename>

`, reason, os.Args[0])
	os.Exit(1)
}
