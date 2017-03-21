package zog

type Table struct {
}

// These tables participate in IX/IY replacement
var baseTableR []Loc8 = []Loc8{B, C, D, E, H, L, Contents{HL}, A}
var baseTableRP []Loc16 = []Loc16{BC, DE, HL, SP}
var baseTableRP2 []Loc16 = []Loc16{BC, DE, HL, AF}

var tableCC []Conditional = []Conditional{Not{FT_Z}, FT_Z, Not{FT_C}, FT_C, FT_PO, FT_PE, FT_P, FT_M}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) LookupR(i byte) Loc8 {
	return baseTableR[i]
}

func (t *Table) LookupRP(i byte) Loc16 {
	return baseTableRP[i]
}

func (t *Table) LookupRP2(i byte) Loc16 {
	return baseTableRP2[i]
}

type AccumInfo struct {
	name string
	//	f    AccumFunc
}

var tableALU []AccumInfo = []AccumInfo{
	/*
		{"ADD", AccumADD8},
		{"ADC", AccumADC8},
		{"SUB", AccumSUB8},
		{"SBC", AccumSBC8},
		{"AND", AccumAND8},
		{"XOR", AccumXOR8},
		{"OR", AccumOR8},
		{"CP", AccumCP8},
	*/
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
