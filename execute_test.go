package zog

import (
	"fmt"
	"testing"
)

type executeAssertion struct {
	loc Loc8
	val byte
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
		assembly, err := Assemble(tc.prog)
		if err != nil {
			t.Fatalf("Failed to assemble [%s]: %s", tc.prog, err)
		}

		z := New(0x1000)
		err = z.Run(assembly)
		if err != nil {
			t.Fatalf("Failed to execute [%s]: %s", tc.prog, err)
		}
	}
}
