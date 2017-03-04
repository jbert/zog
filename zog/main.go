package main

import (
	"fmt"

	"github.com/jbert/zog"
)

func main() {
	memSize := uint16(16 * 1024)
	z := zog.New(memSize)
	/*
			z.Assemble(0, `
			LD A, 10
			HALT
		`)
	*/
	z.Poke(0, 0x0E) // LD C, imm
	z.Poke(1, 0x11) // 0x11
	z.Poke(2, 0x3E) // LD A, imm
	z.Poke(3, 0x22) // 0x22
	z.Poke(4, 0x81) // ADD A, C
	z.Poke(5, 0x76) // HALT
	a, err := z.Run()
	if err != nil {
		fmt.Printf("Terminated: %s\n", err)
	}
	fmt.Printf("A is 0x%02X\n", a)
}
