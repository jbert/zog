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

		{"LD B, 0x34 / LD A, B", 0x34, nil},
		{"LD C, 0x34 / LD A, C", 0x34, nil},
		{"LD D, 0x34 / LD A, D", 0x34, nil},
		{"LD E, 0x34 / LD A, E", 0x34, nil},
	}

	zogTest(t, testCases)
}

//func TestLD(t *testing.T) {
//	testCases := []testCase{
//		{"LD HL, 0x100 / LD (HL), 0x34 / LD A, (HL)", 0x34, nil},
//	}
//
//	zogTest(t, testCases)
//}

func TestFlag(t *testing.T) {
	testCases := []testCase{
		{"",
			0x00, []ef{{F_C, false}}},
		{"SCF",
			0x00, []ef{{F_C, true}}},
		{"SCF / CCF",
			0x00, []ef{{F_C, false}}},
		{"SCF / CCF / CCF",
			0x00, []ef{{F_C, true}}},
		{"SCF / OR A",
			0x00, []ef{{F_C, false}}},
		{"SCF / AND A",
			0x00, []ef{{F_C, false}}},
		{"SCF / XOR A",
			0x00, []ef{{F_C, false}}},
	}

	zogTest(t, testCases)
}

func TestAccum(t *testing.T) {
	testCases := []testCase{
		{"LD A, 0x10 / LD B, 0x20 / ADD A, B",
			0x30, []ef{{F_C, false}, {F_Z, false}}},
		{"LD A, 0xFF / LD B, 0x00 / ADD A, B",
			0xff, []ef{{F_C, false}, {F_Z, false}}},
		{"LD A, 0xFF / LD B, 0x01 / ADD A, B",
			0x00, []ef{{F_C, true}, {F_Z, true}}},
		{"LD A, 0xFF / LD B, 0x02 / ADD A, B",
			0x01, []ef{{F_C, true}, {F_Z, false}}},

		{"LD A, 0x10 / LD B, 0x20 / ADC A, B",
			0x30, []ef{{F_C, false}, {F_Z, false}}},
		{"LD A, 0xFF / LD B, 0x02 / ADC A, B",
			0x01, []ef{{F_C, true}, {F_Z, false}}},
		{"SCF / LD A, 0x10 / LD B, 0x20 / ADC A, B",
			0x31, []ef{{F_C, false}, {F_Z, false}}},
		{"SCF / LD A, 0xFF / LD B, 0x02 / ADC A, B",
			0x02, []ef{{F_C, true}, {F_Z, false}}},

		{"LD A, 0x20 / LD B, 0x10 / SUB B",
			0x10, []ef{{F_C, false}, {F_Z, false}}},
		{"LD A, 0x20 / LD B, 0x30 / SUB B",
			0xf0, []ef{{F_C, true}, {F_Z, false}}},
		{"LD A, 0x20 / LD B, 0x20 / SUB B",
			0x00, []ef{{F_C, false}, {F_Z, true}}},

		{"SCF / LD A, 0x20 / LD B, 0x10 / SBC B",
			0x0f, []ef{{F_C, false}, {F_Z, false}}},
		{"SCF / LD A, 0x20 / LD B, 0x30 / SBC B",
			0xef, []ef{{F_C, true}, {F_Z, false}}},
		{"LD A, 0x20 / LD B, 0x10 / SBC B",
			0x10, []ef{{F_C, false}, {F_Z, false}}},
		{"LD A, 0x20 / LD B, 0x30 / SBC B",
			0xf0, []ef{{F_C, true}, {F_Z, false}}},

		{"SCF / LD A, 0x30 / LD B, 0x10 / AND B",
			0x10, []ef{{F_C, false}, {F_Z, false}}},

		{"SCF / LD A, 0x30 / LD B, 0x10 / XOR B",
			0x20, []ef{{F_C, false}, {F_Z, false}}},

		{"SCF / LD A, 0x30 / LD B, 0x10 / OR B",
			0x30, []ef{{F_C, false}, {F_Z, false}}},

		{"SCF / LD A, 0x30 / LD B, 0x10 / CP B",
			0x30, []ef{{F_C, false}, {F_Z, false}}},
		{"SCF / LD A, 0x30 / LD B, 0x30 / CP B",
			0x30, []ef{{F_C, false}, {F_Z, true}}},
	}

	zogTest(t, testCases)
}

func zogTest(t *testing.T, testCases []testCase) {
	memSize := uint16(1024)
	for _, tc := range testCases {
		fmt.Printf("A: %s\n", tc.assembly)
		z := New(memSize)
		a, err := z.Execute(0, tc.assembly+" / HALT")
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
