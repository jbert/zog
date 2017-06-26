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
		{"LD A,10h", []executeAssertion{{A, 0x10}}},
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
