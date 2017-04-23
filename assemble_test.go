package zog

import "testing"

func TestAssembleAll(t *testing.T) {
	testUtilRunAll(t, func(t *testing.T, byteForm []byte, stringForm string) {
		testAssembleOne(t, stringForm)
	})
}

func TestAssembleBasic(t *testing.T) {
	testCases := []string{
		"RLC (IX+1), B",
		"RLC (IX+1)",

		"SRL (IX+1), B",
		"SRL (IX+1)",

		"SET 7, (IX+1), B",
		"SET 7, (IX+1)",

		"RES 7, (IX+1), B",
		"RES 7, (IX+1)",

		"BIT 7, (IX+1)",

		"LD A, (IX+10)",
		"LD A, (IX-10)",

		// TODO: test hex parses
		//		"LD A, (IX+0x0a)",
		//		"LD A, (IX-0x0a)",
		//		"LD A, (IX+0ah)",
		//		"LD A, (IX-0ah)",

		"OUT (0xff), A",
		"IN A, (0xff)",
		"OUT (c), A",
		"IN A, (c)",

		"EX (SP), HL",

		"LD (0x1234), A",

		"inc iy",
		"inc iyh",

		"add iy, bc",

		"INC B",
		"DEC B",

		"LD A, B",
		"LD A, 0x10",
		"LD A, 0x10",

		"INC DE",
		"ADD DE, HL",
		"EX AF,AF'",
		"RET C",
		"CALL DE",

		"RET C",
		"RST 8",
		"RST 16",
		"DJNZ -10",
		"CALL Z, DE",

		"RL A",
		"SET 4, A",
		"SLA F",

		"LD DE, 0x1234",
		"LD DE, (0x1234)",
		"LD (0x1234), HL",

		"LD (0x1234), H",

		"LD A, (HL)",
		"LD (HL), A",
	}

	for _, s := range testCases {
		testAssembleOne(t, s)
	}
}

func testAssembleOne(t *testing.T, s string) {
	insts, err := Assemble(s)
	if err != nil {
		t.Fatalf("Failed to assemble [%s]: %s", s, err)
	}

	assembledStr := ""
	for _, inst := range insts {
		assembledStr += inst.String() + "\n"
	}
	if !compareAssembly(assembledStr, s) {
		t.Fatalf("Assembled str not equal [%s] != [%s]", assembledStr, s)
	}
}
