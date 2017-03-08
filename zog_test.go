package zog

import (
	"fmt"
	"testing"
)

// EF == expected flag
type ef struct {
	f     flag
	value bool
}
type testCase struct {
	assembly  string
	expectedA byte
	flags     []ef
}

func TestLD8(t *testing.T) {
	testCases := []testCase{
		{"LD A, 10", 10, nil},
		{"LD A, 0x10", 0x10, nil},
		{"LD A, 0xff", 0xff, nil},

		{"LD B, 0x34; LD A, B", 0x34, nil},
		{"LD C, 0x34; LD A, C", 0x34, nil},
		{"LD D, 0x34; LD A, D", 0x34, nil},
		{"LD E, 0x34; LD A, E", 0x34, nil},
	}

	zogTest(t, testCases)
}

//func TestLD(t *testing.T) {
//	testCases := []testCase{
//		{"LD HL, 0x100; LD (HL), 0x34; LD A, (HL)", 0x34, nil},
//	}
//
//	zogTest(t, testCases)
//}

func TestAccum(t *testing.T) {
	testCases := []testCase{
		{"LD A, 10",
			10, nil},
		{"LD A, 0x10; LD B, 0x20; ADD A, B",
			0x30, []ef{{F_C, false}}},
		{"LD A, 0xFF; LD B, 0x02; ADD A, B",
			0x01, []ef{{F_C, true}}},
	}

	zogTest(t, testCases)
}

func zogTest(t *testing.T, testCases []testCase) {
	memSize := uint16(1024)
	for _, tc := range testCases {
		fmt.Printf("A: %s\n", tc.assembly)
		z := New(memSize)
		a, err := z.Execute(0, tc.assembly+"; HALT")
		if err != nil {
			t.Fatalf("Can't execute test prog [%s]: %s", tc.assembly, err)
		}
		if a != tc.expectedA {
			t.Fatalf("Wrong accum for [%s]: 0x%02X != 0x%02X", tc.assembly, a, tc.expectedA)
		}
		for _, ef := range tc.flags {
			fv := z.GetFlag(ef.f)
			if fv != ef.value {
				t.Fatalf("Wrong flag for [%s]: flag %s != %v", tc.assembly, ef.f, ef.value)
			}
		}
		fmt.Printf("\n")
	}
}
