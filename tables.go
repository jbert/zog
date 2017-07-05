package zog

import (
	"fmt"
	"io"
)

type DecodeTable struct {
	r      io.Reader
	wantIX bool
	wantIY bool
}

// These tables participate in IX/IY replacement
var baseTableR []Loc8 = []Loc8{B, C, D, E, H, L, Contents{HL}, A}
var baseTableRP []Loc16 = []Loc16{BC, DE, HL, SP}

const HL_RP_INDEX = 2

var baseTableRP2 []Loc16 = []Loc16{BC, DE, HL, AF}

var tableCC []Conditional = []Conditional{Not{FT_Z}, FT_Z, Not{FT_C}, FT_C, FT_PO, FT_PE, FT_P, FT_M}
var tableBLI [][]Instruction = [][]Instruction{
	[]Instruction{LDI, CPI, INI, OUTI},
	[]Instruction{LDD, CPD, IND, OUTD},
	[]Instruction{LDIR, CPIR, INIR, OTIR},
	[]Instruction{LDDR, CPDR, INDR, OTDR},
}

func NewDecodeTable(r io.Reader) *DecodeTable {
	return &DecodeTable{r: r}
}

func (t *DecodeTable) ResetPrefix(n byte) {
	t.wantIX = false
	t.wantIY = false
	switch n {
	case 0xDD:
		t.wantIX = true
	case 0xFD:
		t.wantIY = true
	}
}

func (t *DecodeTable) LookupR(i byte) Loc8 {
	l := baseTableR[i]
	if !t.wantIX && !t.wantIY {
		return l
	}

	switch i {
	case 4: // H
		l = IXH
		if t.wantIY {
			l = IYH
		}
	case 5: // L
		l = IXL
		if t.wantIY {
			l = IYL
		}
	case 6: // (HL)
		d, err := getImmd(t.r)
		// TODO: panic is messy here- opens us up to panic on decode
		if err != nil {
			panic(fmt.Errorf("Can't get index displacemnt: %s", err))
		}
		l = IndexedContents{IX, d}
		if t.wantIY {
			l = IndexedContents{IY, d}
		}
	}

	return l
}

func (t *DecodeTable) LookupRP(i byte) Loc16 {
	l := baseTableRP[i]
	if l == HL {
		if t.wantIX {
			l = IX
		}
		if t.wantIY {
			l = IY
		}
	}
	return l
}

func (t *DecodeTable) LookupRP2(i byte) Loc16 {
	l := baseTableRP2[i]
	if l == HL {
		if t.wantIX {
			l = IX
		}
		if t.wantIY {
			l = IY
		}
	}
	return l
}

func (t *DecodeTable) LookupBLI(a, b byte) Instruction {
	return tableBLI[a][b]
}

type AccumInfo struct {
	name string
	f    accumFunc
}

var tableALU []AccumInfo = []AccumInfo{
	{"ADD", aluAdd},
	{"ADC", aluAdc},
	{"SUB", aluSub},
	{"SBC", aluSbc},
	{"AND", aluAnd},
	{"XOR", aluXor},
	{"OR", aluOr},
	{"CP", aluCp},
}

func findFuncInTableALU(name string) accumFunc {
	for _, info := range tableALU {
		if info.name == name {
			return info.f
		}
	}
	panic(fmt.Sprintf("Not found in tableALU: %s", name))
}

type RotInfo struct {
	name string
	f    rotFunc
}

var tableROT []RotInfo = []RotInfo{
	{"RLC", rotRlc},
	{"RRC", rotRrc},
	{"RL", rotRl},
	{"RR", rotRr},
	{"SLA", rotSla},
	{"SRA", rotSra},
	{"SLL", rotSll},
	{"SRL", rotSrl},
}

func findInTableR(l Loc8) byte {
	for i := range baseTableR {
		// String compare to get (HL) to work
		if baseTableR[i].String() == l.String() {
			return byte(i)
		}
	}
	panic(fmt.Sprintf("Not found in table R: %s", l))
}

func findInTableRP(l Loc16) byte {
	for i := range baseTableRP {
		if baseTableRP[i].String() == l.String() {
			return byte(i)
		}
	}
	panic(fmt.Sprintf("Not found in table RP: %s", l))
}

func findInTableRP2(l Loc16) byte {
	for i := range baseTableRP2 {
		if baseTableRP2[i].String() == l.String() {
			return byte(i)
		}
	}
	panic(fmt.Sprintf("Not found in table RP2: %s", l))
}

func findInTableALU(name string) byte {
	for i, info := range tableALU {
		if info.name == name {
			return byte(i)
		}
	}
	panic(fmt.Sprintf("Not found in tableALU: %s", name))
}

func findInTableCC(c Conditional) byte {
	for i := range tableCC {
		// String compare
		if tableCC[i].String() == c.String() {
			return byte(i)
		}
	}
	panic(fmt.Sprintf("Not found in tableCC: %s", c))
}

func findInTableROT(name string) byte {
	for i, info := range tableROT {
		if info.name == name {
			return byte(i)
		}
	}
	panic(fmt.Sprintf("Not found in tableROT: %s", name))
}

func findFuncInTableROT(name string) rotFunc {
	for _, info := range tableROT {
		if info.name == name {
			return info.f
		}
	}
	panic(fmt.Sprintf("Not found in tableROT: %s", name))
}
