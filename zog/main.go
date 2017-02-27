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
	b := z.Read8(zog.B)
	fmt.Printf("%s holds %v\n", zog.B, b)
	z.Write8(zog.B, 0x10)
	b = z.Read8(zog.B)
	fmt.Printf("%s holds %v\n", zog.B, b)
}
