package zog

import (
	"fmt"
	"strings"
	"testing"
)

// ZX spectrum manual has/had list of opcodes:
// http://www.worldofspectrum.org/ZXBasicManual/zxmanappa.html

type testInstruction struct {
	n             byte
	inst          string
	inst_after_cb string
	inst_after_ed string
}

func (tc *testInstruction) getExpected(indexPrefix byte, opPrefix byte, buf []byte) ([]byte, string) {
	var expected string
	switch opPrefix {
	case 0: // No prefix
		expected = tc.inst
	case 0xcb:
		expected = tc.inst_after_cb
	case 0xed:
		expected = tc.inst_after_ed
	default:
		panic(fmt.Sprintf("Unrecognised op prefix %02X", opPrefix))
	}

	switch indexPrefix {
	case 0xdd:
		return indexRegisterMunge("IX", buf, expected)
	case 0xfd:
		return indexRegisterMunge("IY", buf, expected)
	default:
		return buf, expected
	}
}

func indexRegisterMunge(indexRegister string, buf []byte, expected string) ([]byte, string) {
	d := int8(-20)

	// Must match location.go:func (ic IndexedContents) String() format
	hlReplace := fmt.Sprintf("(%s%+d)", indexRegister, d)
	hReplace := indexRegister + "h"
	lReplace := indexRegister + "l"

	if strings.Contains(expected, "(hl)") {
		expected = strings.Replace(expected, "(hl)", hlReplace, -1)
		buf = append(buf, byte(d))
	} else {
		// Exception
		if expected != "ex de,hl" {
			expected = strings.Replace(expected, "hl", indexRegister, -1)
			expected = strings.Replace(expected, "h,", hReplace+",", -1)
			expected = strings.Replace(expected, ",h", ","+hReplace, -1)
			expected = strings.Replace(expected, " h", " "+hReplace, -1)
			expected = strings.Replace(expected, "l,", lReplace+",", -1)
			expected = strings.Replace(expected, ",l", ","+lReplace, -1)
			expected = strings.Replace(expected, " l", " "+lReplace, -1)
		}
	}

	return buf, expected
}

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

func bufToHex(buf []byte) string {
	s := ""
	for _, b := range buf {
		s += fmt.Sprintf("%02X", b)
	}
	return s
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

func expandImmediateData(buf []byte, template string) ([]byte, string) {
	var s string = template
	if strings.Contains(s, "NN") {
		buf = append(buf, 0x34)
		buf = append(buf, 0x12)
		s = strings.Replace(s, "NN", "0x1234", 1)
	}
	if strings.Contains(s, "N") {
		buf = append(buf, 0xab)
		s = strings.Replace(s, "N", "0xab", 1)
	}
	if strings.Contains(s, "DIS") {
		buf = append(buf, 0xf0)
		s = strings.Replace(s, "DIS", "-16", 1)
	}
	return buf, s
}
