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
	z.Poke(0, 0x41)
	err := z.Run()
	if err != nil {
		fmt.Printf("Terminated: %s\n", err)
	}
}
