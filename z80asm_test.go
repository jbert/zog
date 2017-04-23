package zog

import (
	"fmt"
	"testing"
)

func TestZ80AsmBasic(t *testing.T) {
	testCases := []string {
		"LD (IX+10h), A",
	}

	for _, s := range testCases {
		testZ80AsmOne(t, nil, s)
	}
}

/*
func TestZ80AsmAll(t *testing.T) {
	testUtilRunAll(t, func(t *testing.T, byteForm []byte, stringForm string) {
		testZ80AsmOne(t, byteForm, stringForm)
	})
}
*/

func testZ80AsmOne(t *testing.T, byteForm []byte, stringForm string) {
	z80Buf := z80asmAssemble(stringForm)
	insts, err := Assemble(stringForm)
	if err != nil {
		t.Fatalf("Failed to assemble [%s]: %s", stringForm, err)
	}
	if len(insts) != 1 {
		t.Fatalf("Got more or less instructions than 1")
	}
	fmt.Printf("Assembled [%s]\n", stringForm)
	ourBuf := insts[0].Encode()

	if bufToHex(z80Buf) == bufToHex(ourBuf) {
		fmt.Printf("Successfully compared [%s]\n", stringForm)
	} else {
		t.Fatalf("Failed z80 assembly for [%s]: got [%s] z80asm [%s]", stringForm, bufToHex(ourBuf), bufToHex(z80Buf))
	}
}
