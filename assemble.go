package zog

import "fmt"

func Assemble(s string) (*Assembly, error) {

	s += "\n"

	assembler := &PegAssembler{Buffer: s}
	assembler.Init()
	assembler.Current.Init()
	err := assembler.Parse()
	if err != nil {
		//assembler.Print()
		assembler.PrintSyntaxTree()
		return nil, fmt.Errorf("Can't parse: %s", err)
	}
	//assembler.Print()
	//assembler.PrintSyntaxTree()
	assembler.Execute()
	assembly := assembler.GetAssembly()

	err = assembly.ResolveAddresses()
	if err != nil {
		return nil, fmt.Errorf("Can't resolve addresses: %s", err)
	}

	return assembly, nil
}

type LabelledInstruction struct {
	Label string
	Inst  Instruction
	Addr  uint16
}

type Assembly struct {
	BaseAddr uint16
	Linsts   []LabelledInstruction
	Labels   map[string]int
	resolved bool
}

func (a *Assembly) Encode() ([]byte, error) {
	err := a.ResolveAddresses()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 0)
	for _, linst := range a.Linsts {
		instBuf := linst.Inst.Encode()
		buf = append(buf, instBuf...)
	}
	return buf, nil
}

func (a *Assembly) ResolveAddresses() error {
	if a.resolved {
		return nil
	}

	// One pass to work out addresses for instructions
	addr := a.BaseAddr
	for i := range a.Linsts {
		a.Linsts[i].Addr = addr
		buf := a.Linsts[i].Inst.Encode()
		addr += uint16(len(buf))
	}
	// One pass to find and resolve labels
	for i := range a.Linsts {
		err := a.Linsts[i].Inst.Resolve(a)
		if err != nil {
			return fmt.Errorf("Failed to resolve label for [%s]: %s", a.Linsts[i].Inst, err)
		}
		fmt.Printf("%04X: %s\n", a.Linsts[i].Addr, a.Linsts[i].Inst)
	}
	a.resolved = true
	return nil
}

func (a *Assembly) ResolveLoc8(l Loc8) error {
	contents, ok := l.(Contents)
	if !ok {
		return nil
	}
	return a.resolveContents(contents)
}

func (a *Assembly) ResolveLoc16(l Loc16) error {
	label, ok := l.(*Label)
	if ok {
		addr, err := a.FindLabelAddr(label.name)
		if err != nil {
			return err
		}
		label.Imm16 = Imm16(addr)
		return nil
	}

	contents, ok := l.(Contents)
	if !ok {
		return nil
	}
	return a.resolveContents(contents)
}

func (a *Assembly) resolveContents(c Contents) error {
	label, ok := c.addr.(*Label)
	if !ok {
		return nil
	}

	addr, err := a.FindLabelAddr(label.name)
	if err != nil {
		return err
	}
	label.Imm16 = Imm16(addr)
	return nil
}

func (a *Assembly) FindLabelAddr(name string) (uint16, error) {
	index, ok := a.Labels[name]
	if !ok {
		return 0, fmt.Errorf("Can't find label: %s", name)
	}

	if index < 0 || index >= len(a.Linsts) {
		return 0, fmt.Errorf("Label index out of range (%d > %d)", index, len(a.Linsts))
	}
	return a.Linsts[index].Addr, nil
}

func (a *Assembly) Init() {
	a.Labels = make(map[string]int)
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
