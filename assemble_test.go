package zog

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestAssembleBasic(t *testing.T) {
	testCases := []string{

		"LD A, B",
		"LD A, 0x10",
		"LD A, 10h",

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
		"LD A, (HL)",
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

func normaliseWhiteSpace(s string) string {
	collapseSpaces := regexp.MustCompile(" +")
	s = collapseSpaces.ReplaceAllString(s, " ")
	commaSpace := regexp.MustCompile(", ")
	s = commaSpace.ReplaceAllString(s, ",")
	return s
}
func normaliseHex(s string) string {
	re := regexp.MustCompile("0x([[:xdigit:]]{1,4})")
	s = re.ReplaceAllString(s, "${1}h")
	return s
}

func normaliseAssembly(s string) string {
	s = strings.TrimSpace(s)
	s = normaliseHex(s)
	s = normaliseWhiteSpace(s)
	return s
}

func compareAssembly(a, b string) bool {
	a = normaliseAssembly(a)
	b = normaliseAssembly(b)
	fmt.Printf("a [%s] b [%s]\n", a, b)
	return a == b
}
