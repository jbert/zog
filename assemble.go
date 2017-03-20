package zog

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Assembler struct {
	lookup map[string]func([]string) (Instruction, error)
}

func NewAssembler() *Assembler {
	a := Assembler{}

	a.lookup = map[string]func([]string) (Instruction, error){
		"LD": ParseLD,

		"ADD": MakeParseAccum(0),
		"ADC": MakeParseAccum(1),
		"SUB": MakeParseAccum(2),
		"SBC": MakeParseAccum(3),
		"AND": MakeParseAccum(4),
		"XOR": MakeParseAccum(5),
		"OR":  MakeParseAccum(6),
		"CP":  MakeParseAccum(7),

		"PUSH": ParsePush,
		"POP":  ParsePop,
	}

	for _, info := range decoder.SimpleInfo() {
		a.lookup[info.name] = MakeNoArgs(info.i)
	}

	return &a
}

func MakeNoArgs(i Instruction) func([]string) (Instruction, error) {
	return func(s []string) (Instruction, error) {
		return i, nil
	}
}

var Reg2R8Loc map[string]R8Loc = map[string]R8Loc{
	"B":    B,
	"C":    C,
	"D":    D,
	"E":    E,
	"F":    F,
	"L":    L,
	"(HL)": HL_CONTENTS,
	"A":    A,
	"H":    H,
}

var Reg2R16Loc map[string]R16Loc = map[string]R16Loc{
	"AF": AF,
	"BC": BC,
	"DE": DE,
	"HL": HL,
	"SP": SP,
	"IX": IX,
	"IY": IY,
}

func MakeParseAccum(hi3 byte) func(tokens []string) (Instruction, error) {
	return func(tokens []string) (Instruction, error) {
		// We permit (but do not require) a leading "A, "
		if len(tokens) == 2 {
			if tokens[0] == "A" {
				tokens = tokens[1:]
			} else {
				return nil, fmt.Errorf("Must have one token (or leading A,): [%v]", tokens)
			}
		}

		src, ok := Reg2R8Loc[tokens[0]]
		if !ok {
			return nil, fmt.Errorf("Can't parse [%s] as src R8Loc", tokens[0])
		}

		return NewAccumOp(hi3, byte(src)), nil
	}
}

func ParseLD16(tokens []string) (Instruction, error) {
	var dst16 Location16
	var ok bool
	dstStr := tokens[0]
	if dstStr[0] == '(' && dstStr[len(dstStr)-1] == ')' {
		dstAddr, err := strconv.ParseUint(dstStr[1:len(dstStr)-2], 0, 16)
		if err != nil {
			return nil, fmt.Errorf("Can't parse [%s]:  %s", dstStr, err)
		}
		dst16 = Addr16(dstAddr)
	} else {
		dst16, ok = Reg2R16Loc[dstStr]
		if !ok {
			return nil, fmt.Errorf("Can't parse [%s] as dst R8Loc or R16Loc or immediate addr", dstStr)
		}
	}

	srcStr := tokens[1]
	src16, ok := Reg2R16Loc[srcStr]
	if ok {
		return &ILD16{src: src16, dst: dst16}, nil
	}

	r16loc, ok := dst16.(R16Loc)
	if !ok {
		return nil, fmt.Errorf("Can't have immediate destination and source in LD16")
	}

	// Maybe 16 bit immediate?
	num, err := strconv.ParseUint(srcStr, 0, 16)
	if err != nil {
		return nil, fmt.Errorf("Can't parse [%s]:  %s", srcStr, err)
	}
	return &ILD16Immediate{dst: r16loc, nn: uint16(num)}, nil
}

func ParsePush(tokens []string) (Instruction, error) {
	if len(tokens) != 1 {
		return nil, fmt.Errorf("Must have one token for Push: [%v]", tokens)
	}
	src16, ok := Reg2R16Loc[tokens[0]]
	if !ok {
		return nil, fmt.Errorf("Can't parse push r16 [%s]", tokens[0])
	}
	return &IPush{src: src16}, nil
}

func ParsePop(tokens []string) (Instruction, error) {
	if len(tokens) != 1 {
		return nil, fmt.Errorf("Must have one token for Pop: [%v]", tokens)
	}
	dst16, ok := Reg2R16Loc[tokens[0]]
	if !ok {
		return nil, fmt.Errorf("Can't parse pop r16 [%s]", tokens[0])
	}
	return &IPop{dst: dst16}, nil
}

func ParseLD(tokens []string) (Instruction, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("Must have two tokens for LD: [%v]", tokens)
	}

	dst8, ok := Reg2R8Loc[tokens[0]]
	if !ok {
		return ParseLD16(tokens)
	}
	if dst8 == H {
		return nil, errors.New("Can't LD to H")
	}

	src, ok := Reg2R8Loc[tokens[1]]
	if ok {
		return &ILD8{src: src, dst: dst8}, nil
	}

	// Maybe 8 bit immediate?
	num, err := strconv.ParseUint(tokens[1], 0, 8)
	if err != nil {
		return nil, fmt.Errorf("Can't parse [%s]:  %s", tokens[1], err)
	}
	return &ILD8Immediate{dst: dst8, n: byte(num)}, nil
}

func (a *Assembler) AssembleOne(s string) (Instruction, error) {

	tokens := strings.Split(s, " ")
	if len(tokens) < 1 {
		return nil, fmt.Errorf("Blank string? [%s]", s)
	}
	iStr := tokens[0]

	// Drop trailing commas
	for i := range tokens {
		tokens[i] = strings.Trim(tokens[i], ",")
	}

	/*
		str := func() string {
			for {
				tokens = tokens[1:] // skip last
				if len(tokens) < 1 {
					panic("Insufficient tokens")
				}
				tok := tokens[0]
				return tok
			}
		}
	*/

	parser, ok := a.lookup[strings.ToUpper(iStr)]
	if !ok {
		return nil, fmt.Errorf("Can't find parser for [%s]", iStr)
	}

	return parser(tokens[1:])
}

func (a *Assembler) Assemble(s string) ([]Instruction, error) {
	// Support / for single-line assembly
	s = strings.Replace(s, "/", "\n", -1)
	strs := strings.Split(s, "\n")
	return a.AssembleLines(strs)
}

func (a *Assembler) AssembleLines(strs []string) ([]Instruction, error) {
	var instructions []Instruction
	for lineNumber, s := range strs {
		comment := strings.Index(s, ";")
		if comment > 0 {
			s = s[:comment]
		}
		ts := strings.Trim(s, " \t")
		if ts == "" {
			continue
		}
		i, err := a.AssembleOne(ts)
		if err != nil {
			return nil, fmt.Errorf("line %d: %s", lineNumber, err)
		}
		instructions = append(instructions, i)
	}

	return instructions, nil
}
