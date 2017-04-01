package zog

import "fmt"

type Table struct {
	inCh   chan byte
	wantIX bool
	wantIY bool
}

// These tables participate in IX/IY replacement
var baseTableR []Loc8 = []Loc8{B, C, D, E, H, L, Contents{HL}, A}
var baseTableRP []Loc16 = []Loc16{BC, DE, HL, SP}
var baseTableRP2 []Loc16 = []Loc16{BC, DE, HL, AF}

var tableCC []Conditional = []Conditional{Not{FT_Z}, FT_Z, Not{FT_C}, FT_C, FT_PO, FT_PE, FT_P, FT_M}
var tableBLI [][]Instruction = [][]Instruction{
	[]Instruction{LDI, CPI, INI, OUTI},
	[]Instruction{LDD, CPD, IND, OUTD},
	[]Instruction{LDIR, CPIR, INIR, OTIR},
	[]Instruction{LDDR, CPDR, INDR, OTDR},
}

func NewTable(inCh chan byte) *Table {
	return &Table{inCh: inCh}
}

func (t *Table) ResetPrefix(n byte) {
	t.wantIX = false
	t.wantIY = false
	switch n {
	case 0xDD:
		t.wantIX = true
	case 0xFD:
		t.wantIY = true
	}
}

func (t *Table) LookupR(i byte) Loc8 {
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
		d, err := getImmd(t.inCh)
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

func (t *Table) LookupRP(i byte) Loc16 {
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

func (t *Table) LookupRP2(i byte) Loc16 {
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

func (t *Table) LookupBLI(a, b byte) Instruction {
	return tableBLI[a][b]
}

type AccumInfo struct {
	name string
	//	f    AccumFunc
}

var tableALU []AccumInfo = []AccumInfo{
	{"ADD"},
	{"ADC"},
	{"SUB"},
	{"SBC"},
	{"AND"},
	{"XOR"},
	{"OR"},
	{"CP"},
}

type RotInfo struct {
	name string
	//	f    AccumFunc
}

var tableROT []RotInfo = []RotInfo{
	{"RLC"},
	{"RRC"},
	{"RL"},
	{"RR"},
	{"SLA"},
	{"SRA"},
	{"SLL"},
	{"SRL"},
}
