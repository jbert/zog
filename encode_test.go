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
		{[]byte{0xed, 0x51}, "OUT (C), D"},
		{[]byte{0xed, 0x50}, "IN D, (C)"},

		{[]byte{0xcb, 0x20}, "SLA B"},
		{[]byte{0xcb, 0x26}, "SLA (HL)"},
		{[]byte{0xdd, 0xcb, 0x26, 0x10}, "SLA (IX+10h)"},

		{[]byte{0xe5}, "PUSH HL"},
		{[]byte{0xdd, 0xe5}, "PUSH IX"},
		{[]byte{0xe1}, "POP HL"},
		{[]byte{0xdd, 0xe1}, "POP IX"},

		{[]byte{0x27}, "DAA"},

		{[]byte{0x2b}, "DEC HL"},
		{[]byte{0xdd, 0x2b}, "DEC IX"},

		{[]byte{0x22, 0x34, 0x12}, "LD (1234h), HL"},
		{[]byte{0xdd, 0x22, 0x34, 0x12}, "LD (1234h), IX"},
		{[]byte{0x2a, 0x34, 0x12}, "LD HL, (1234h)"},
		{[]byte{0xdd, 0x2a, 0x34, 0x12}, "LD IX, (1234h)"},

		{[]byte{0x3a, 0x34, 0x12}, "LD A, (1234h)"},
		{[]byte{0x32, 0x34, 0x12}, "LD (1234h), A"},

		{[]byte{0xcf}, "RST 8"},
		{[]byte{0xc7}, "RST 0"},

		{[]byte{0xc6, 0x10}, "ADD A, 10h"},
		{[]byte{0xe6, 0x10}, "AND 10h"},

		{[]byte{0xcd, 0x34, 0x12}, "CALL 1234h"},
		{[]byte{0xf5}, "PUSH AF"},

		{[]byte{0xcc, 0x34, 0x12}, "CALL Z, 0x1234"},

		{[]byte{0xdb, 0x20}, "IN A, (20h)"},
		{[]byte{0xd3, 0x20}, "OUT (20h), A"},

		{[]byte{0xeb}, "EX DE, HL"},
		{[]byte{0xe3}, "EX (SP), HL"},
		{[]byte{0xdd, 0xe3}, "EX (SP), IX"},

		{[]byte{0xc3, 0x34, 0x12}, "JP 0x1234"},
		{[]byte{0xc2, 0x34, 0x12}, "JP NZ, 0x1234"},

		{[]byte{0xe9}, "JP HL"},
		{[]byte{0xf9}, "LD SP, HL"},
		{[]byte{0xdd, 0xe9}, "JP IX"},
		{[]byte{0xdd, 0xf9}, "LD SP, IX"},

		{[]byte{0xd9}, "EXX"},

		{[]byte{0xc9}, "RET"},
		{[]byte{0xc8}, "RET Z"},

		{[]byte{0x07}, "RLCA"},

		{[]byte{0x0b}, "DEC BC"},

		{[]byte{0x09}, "ADD HL, BC"},
		{[]byte{0xdd, 0x09}, "ADD IX, BC"},

		{[]byte{0x03}, "INC BC"},
		{[]byte{0x23}, "INC HL"},
		{[]byte{0xdd, 0x23}, "INC IX"},
		{[]byte{0xfd, 0x23}, "INC IY"},

		{[]byte{0x3e, 0xab}, "LD A, ABh"},

		{[]byte{0x01, 0x34, 0x12}, "LD BC, 0x1234"},
		{[]byte{0x21, 0x34, 0x12}, "LD HL, 0x1234"},
		{[]byte{0xdd, 0x21, 0x34, 0x12}, "LD IX, 0x1234"},
		{[]byte{0xfd, 0x21, 0x34, 0x12}, "LD IY, 0x1234"},

		{[]byte{0x18, 0xf0}, "JR -10h"},
		{[]byte{0x20, 0xf0}, "JR NZ, -10h"},
		{[]byte{0x28, 0xf0}, "JR Z, -10h"},
		{[]byte{0x08}, "EX AF, AF'"},

		{[]byte{0x10, 0xf0}, "DJNZ -10h"},

		{[]byte{0xcb, 0xd0}, "SET 2, B"},
		{[]byte{0xcb, 0xd6}, "SET 2, (HL)"},
		{[]byte{0xdd, 0xcb, 0xd6, 0x10}, "SET 2, (IX+0x10)"},
		{[]byte{0xfd, 0xcb, 0xd6, 0x10}, "SET 2, (IY+0x10)"},

		{[]byte{0xcb, 0x90}, "RES 2, B"},
		{[]byte{0xcb, 0x96}, "RES 2, (HL)"},
		{[]byte{0xdd, 0xcb, 0x96, 0x10}, "RES 2, (IX+0x10)"},
		{[]byte{0xfd, 0xcb, 0x96, 0x10}, "RES 2, (IY+0x10)"},

		{[]byte{0xcb, 0x50}, "BIT 2, B"},
		{[]byte{0xcb, 0x56}, "BIT 2, (HL)"},
		{[]byte{0xdd, 0xcb, 0x56, 0x10}, "BIT 2, (IX+0x10)"},
		{[]byte{0xfd, 0xcb, 0x56, 0x10}, "BIT 2, (IY+0x10)"},

		{[]byte{0x80}, "ADD A, B"},
		{[]byte{0x86}, "ADD A, (HL)"},
		{[]byte{0xdd, 0x86, 0x10}, "ADD A, (IX+0x10)"},
		{[]byte{0xfd, 0x86, 0x10}, "ADD A, (IY+0x10)"},

		{[]byte{0xa0}, "AND B"},
		{[]byte{0xa6}, "AND (HL)"},
		{[]byte{0xdd, 0xa6, 0x10}, "AND (IX+0x10)"},
		{[]byte{0xfd, 0xa6, 0x10}, "AND (IY+0x10)"},

		{[]byte{0x05}, "DEC B"},
		{[]byte{0x0D}, "DEC C"},
		{[]byte{0x35}, "DEC (HL)"},
		{[]byte{0xdd, 0x35, 0x10}, "DEC (IX+0x10)"},
		{[]byte{0xfd, 0x35, 0x10}, "DEC (IY+0x10)"},

		{[]byte{0x00}, "NOP"},
		{[]byte{0xf3}, "DI"},
		{[]byte{0x76}, "HALT"},

		{[]byte{0x04}, "INC B"},
		{[]byte{0x0C}, "INC C"},
		{[]byte{0x34}, "INC (HL)"},
		{[]byte{0xdd, 0x34, 0x10}, "INC (IX+0x10)"},
		{[]byte{0xfd, 0x34, 0x10}, "INC (IY+0x10)"},

		{[]byte{0x41}, "LD B, C"},
		{[]byte{0x51}, "LD D, C"},

		{[]byte{0x46}, "LD B, (HL)"},
		{[]byte{0xdd, 0x46, 0x10}, "LD B, (IX+10h)"},
		{[]byte{0xfd, 0x46, 0x10}, "LD B, (IY+10h)"},

		{[]byte{0x70}, "LD (HL), B"},
		{[]byte{0xdd, 0x70, 0x10}, "LD (IX+10h), B"},
		{[]byte{0xfd, 0x70, 0x10}, "LD (IY+10h), B"},

		{[]byte{0x7e}, "LD A, (HL)"},
		{[]byte{0xdd, 0x7e, 0x10}, "LD A, (IX+10h)"},
		{[]byte{0xfd, 0x7e, 0x10}, "LD A, (IY+10h)"},

		{[]byte{0x77}, "LD (HL), A"},
		{[]byte{0xdd, 0x77, 0x10}, "LD (IX+10h), A"},
		{[]byte{0xfd, 0x77, 0x10}, "LD (IY+10h), A"},
	}

	for _, tc := range testCases {
		testEncodeOne(t, tc.buf, tc.expected)
	}
}

func TestEncodeAll(t *testing.T) {
	testUtilRunAll(t, func(t *testing.T, byteForm []byte, stringForm string) {
		testEncodeOne(t, byteForm, stringForm)
	})
}

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
		t.Fatalf("Wrong encode for [%s] got [%s] expected [%s]", stringForm, bufToHex(encodedBuf), bufToHex(byteForm))
	}
	fmt.Printf("Encoded [%s] to [%s]\n", stringForm, bufToHex(encodedBuf))
}
