package zog

import "fmt"

func Assemble(s string) (*Assembly, error) {

	assembler := &PegAssembler{Buffer: s}
	assembler.Init()
	err := assembler.Parse()
	if err != nil {
		return nil, fmt.Errorf("Can't parse: %s", err)
	}
	// assembler.Print()
	//assembler.PrintSyntaxTree()
	assembler.Execute()

	assembly := assembler.GetAssembly()

	return assembly, nil
}

type LabelledInstruction struct {
	Label string
	Inst  Instruction
}

type Assembly struct {
	BaseAddr uint16
	Linsts   []LabelledInstruction
	Labels   map[string]int
}

func (a *Assembly) Instructions() []Instruction {
	insts := make([]Instruction, 0)
	for _, linst := range a.Linsts {
		insts = append(insts, linst.Inst)
	}
	return insts
}

func (a *Assembly) String() string {
	colWidth := 20
	str := ""

	printPrefix := func(prefix string, iStr string) string {
		if prefix != "" {
			prefix += ":"
		}

		s := ""
		s += fmt.Sprintf(prefix)
		width := colWidth - len(prefix)
		for i := 0; i < width; i++ {
			s += fmt.Sprintf(" ")
		}
		s += iStr
		s += "\n"
		return s
	}
	str += printPrefix("", fmt.Sprintf("org %04x", a.BaseAddr))
	str += printPrefix("", "")
	for _, linst := range a.Linsts {
		str += printPrefix(linst.Label, linst.Inst.String())
	}
	str += printPrefix("", "")

	return str
}
