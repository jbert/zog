package zog

import (
	"fmt"
	"testing"
)

// ZX spectrum manual has/had list of opcodes:
// http://www.worldofspectrum.org/ZXBasicManual/zxmanappa.html

func TestDecodeOddities(t *testing.T) {
	testCases := []struct {
		expected string
		buf      []byte
	}{
		{"SET 0,b", []byte{0xcb, 0xc0}},
		{"SET 0,(IX+10),b", []byte{0xdd, 0xcb, 0x0a, 0xc0}},

		{"rlc (iy+10),b", []byte{0xfd, 0xcb, 0x0a, 0x00}},
		{"rlc (iy+10)", []byte{0xfd, 0xcb, 0x0a, 0x06}},

		{"rlc (ix+10),b", []byte{0xdd, 0xcb, 0x0a, 0x00}},
		{"rlc (ix+10)", []byte{0xdd, 0xcb, 0x0a, 0x06}},

		{"INC IX", []byte{0xDD, 0x23}},
		{"LD A, (IX+1)", []byte{0xFD, 0xDD, 0x7e, 0x01}},
		{"EX DE, HL", []byte{0xDD, 0xeb}},
		{"LD H, (IX+1)", []byte{0xDD, 0x66, 0x01}},
		{"ADD IX, BC", []byte{0xDD, 0x09}},
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
		t.Fatalf("Wrong decode for [%s] got [%s] expected [%s]", hexBuf, insts[0].String(), expected)
	}
	fmt.Printf("Decoded [%s] to [%s]\n", hexBuf, insts[0].String())
}
