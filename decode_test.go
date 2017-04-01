package zog

import (
	"fmt"
	"testing"
)

// ZX spectrum manual has/had list of opcodes:
// http://www.worldofspectrum.org/ZXBasicManual/zxmanappa.html

func TestDecodeOddities(t *testing.T) {
	testCases := []struct {
		reason   string
		buf      []byte
		expected string
	}{
		{"Multiple prefixes, last one wins",
			[]byte{0xFD, 0xDD, 0x7e, 0x01}, "LD A, (IX+1)"},
		{"EX DE, HL is an exception to index prefix",
			[]byte{0xDD, 0xeb}, "EX DE, HL"},
		{"If we index (HL), we don't index H or L",
			[]byte{0xDD, 0x66, 0x01}, "LD H, (IX+1)"},
		{"Indxed 16bit add",
			[]byte{0xDD, 0x09}, "ADD IX, BC"},
	}

	for _, tc := range testCases {
		testDecodeOne(t, tc.buf, tc.expected)
	}
}

func TestDecodeAll(t *testing.T) {
	testUtilRunAll(t, func(t *testing.T, byteForm []byte, stringForm string) {
		testDecodeOne(t, byteForm, stringForm)
	})
}

func testDecodeOne(t *testing.T, byteForm []byte, expected string) {
	hexBuf := bufToHex(byteForm)

	fmt.Printf("== Decode: buf [%s] -> string [%s]\n", hexBuf, expected)
	insts, err := DecodeBytes(byteForm)
	if err != nil {
		t.Fatalf("Error for byte [%s]: %s (%s)", hexBuf, err, expected)
	}
	if len(insts) == 0 {
		t.Fatalf("No instructions for byte [%s] (%s)", hexBuf, expected)
	}
	if len(insts) != 1 {
		t.Fatalf("More than one instruction (%d) for byte [%s]: %v", len(insts), hexBuf, insts)
	}
	if !compareAssembly(insts[0].String(), expected) {
		t.Fatalf("Wrong decode for [%s] [%s] != [%s]", hexBuf, insts[0].String(), expected)
	}
	fmt.Printf("Decoded [%s] to [%s]\n", hexBuf, insts[0].String())
}
