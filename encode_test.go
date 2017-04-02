package zog

import (
	"fmt"
	"testing"
)

func TestEncodeBasic(t *testing.T) {
	testCases := []struct {
		buf      []byte
		expected string
	}{
		{[]byte{0x00}, "NOP"},
		{[]byte{0xf3}, "DI"},
		//		{[]byte{0x77}, "LD (HL), A"},
		//		{[]byte{0xdd, 0x77, 0x10}, "LD (IX+10h), A"},
		//		{[]byte{0xfd, 0x77, 0x10}, "LD (IY+10h), A"},
	}

	for _, tc := range testCases {
		testEncodeOne(t, tc.buf, tc.expected)
	}
}

/*
func TestEncodeAll(t *testing.T) {
	testUtilRunAll(t, func(t *testing.T, byteForm []byte, stringForm string) {
		testEncodeOne(t, byteForm, stringForm)
	})
}
*/

func testEncodeOne(t *testing.T, byteForm []byte, stringForm string) {
	hexBuf := bufToHex(byteForm)

	fmt.Printf("== Encode buf [%s] -> string [%s]\n", hexBuf, stringForm)
	insts, err := Assemble(stringForm)
	if err != nil {
		t.Fatalf("Error for byte [%s]: %s (%s)", hexBuf, err, stringForm)
	}

	if len(insts) == 0 {
		t.Fatalf("No instructions for byte [%s] (%s)", hexBuf, stringForm)
	}
	if len(insts) != 1 {
		t.Fatalf("More than one instruction (%d) for byte [%s]: %v", len(insts), hexBuf, insts)
	}

	encodedBuf := insts[0].Encode()

	if bufToHex(encodedBuf) != bufToHex(byteForm) {
		t.Fatalf("Wrong encode for [%s] [%s] != [%s]", stringForm, bufToHex(encodedBuf), bufToHex(byteForm))
	}
	fmt.Printf("Encoded [%s] to [%s]\n", stringForm, bufToHex(encodedBuf))
}
