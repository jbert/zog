package zog

import "fmt"

func Assemble(s string) ([]Instruction, error) {

	assembler := &PegAssembler{Buffer: s}
	assembler.Init()
	err := assembler.Parse()
	if err != nil {
		return nil, fmt.Errorf("Can't parse: %s", err)
	}
	// assembler.Print()
	//assembler.PrintSyntaxTree()
	assembler.Execute()

	insts := assembler.GetInstructions()

	return insts, nil
}
