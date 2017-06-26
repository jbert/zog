package zog

import (
	"fmt"
	"testing"
)

type executeAssertion struct {
	loc      Loc8
	expected byte
}

func (ea executeAssertion) check(z *Zog) error {
	actual, err := ea.loc.Read8(z)
	if err != nil {
		return fmt.Errorf("assert failed: failed to read location [%s]: %s", ea.loc, err)
	}
	if actual != ea.expected {
		return fmt.Errorf("assert failed: loc [%s] actual %02X expected %02X", ea.loc, actual, ea.expected)
	}
	return nil
}

type executeTestCase struct {
	prog       string
	assertions []executeAssertion
}

func TestExecuteBasic(t *testing.T) {
	testCases := []executeTestCase{
		{"LD A,10h", []executeAssertion{{A, 0x10}, {B, 0x00}}},
		{"LD B,10h", []executeAssertion{{B, 0x10}, {A, 0x00}}},
		{"LD C,10h", []executeAssertion{{C, 0x10}, {A, 0x00}}},
		{"LD D,10h", []executeAssertion{{D, 0x10}, {A, 0x00}}},
		{"LD E,10h", []executeAssertion{{E, 0x10}, {A, 0x00}}},
		{"LD H,10h", []executeAssertion{{H, 0x10}, {A, 0x00}}},
		{"LD L,10h", []executeAssertion{{L, 0x10}, {A, 0x00}}},

		{"LD BC,1234h", []executeAssertion{{B, 0x12}, {C, 0x34}, {A, 0x00}}},
		{"LD DE,1234h", []executeAssertion{{D, 0x12}, {E, 0x34}, {A, 0x00}}},
		{"LD HL,1234h", []executeAssertion{{H, 0x12}, {L, 0x34}, {A, 0x00}}},
		{"LD IX,1234h", []executeAssertion{{IXH, 0x12}, {IXL, 0x34}, {A, 0x00}}},
		{"LD IY,1234h", []executeAssertion{{IYH, 0x12}, {IYL, 0x34}, {A, 0x00}}},
	}
	for _, tc := range testCases {
		fmt.Printf("Assemble: %s\n", tc.prog)
		prog := tc.prog + ": HALT"
		assembly, err := Assemble(prog)
		if err != nil {
			t.Fatalf("Failed to assemble [%s]: %s", prog, err)
		}

		z := New(0x1000)
		err = z.Run(assembly)
		if err != nil {
			t.Fatalf("Failed to execute [%s]: %s", prog, err)
		}

		for _, assertion := range tc.assertions {
			err := assertion.check(z)
			if err != nil {
				t.Error(err)
			}
		}
	}
}
