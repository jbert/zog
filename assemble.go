package zog

import (
	"fmt"
	"strconv"
	"strings"
)

func Single(i Instruction) func([]string) (Instruction, error) {
	return func([]string) (Instruction, error) {
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

var AssemblyParser map[string]func([]string) (Instruction, error) = map[string]func([]string) (Instruction, error){
	"LD":   ParseLD,
	"ADD":  ParseADD,
	"HALT": Single(&IHalt{}),
}

func ParseADD(tokens []string) (Instruction, error) {
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

	//	return &IAccumOp{src: src, name: ops[hi3].name, op: ops[hi3].op}, nil
	hi3 := byte(0)
	return decodeAccumOp(hi3, byte(src))
}

func ParseLD(tokens []string) (Instruction, error) {
	if len(tokens) != 2 {
		return nil, fmt.Errorf("Must have two tokens for LD: [%v]", tokens)
	}

	dst, ok := Reg2R8Loc[tokens[0]]
	if !ok {
		return nil, fmt.Errorf("Can't parse [%s] as dst R8Loc", tokens[0])
	}

	src, ok := Reg2R8Loc[tokens[1]]
	if ok {
		return &ILD8{src: src, dst: dst}, nil
	}

	// Maybe 8 bit immediate?
	num, err := strconv.ParseInt(tokens[1], 0, 8)
	if err != nil {
		return nil, fmt.Errorf("Can't parse [%s]", tokens[1])
	}
	return &ILD8Immediate{dst: dst, n: byte(num)}, nil
}

func AssembleOne(s string) (Instruction, error) {

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

	parser, ok := AssemblyParser[strings.ToUpper(iStr)]
	if !ok {
		return nil, fmt.Errorf("Can't find parser for [%s]", iStr)
	}

	return parser(tokens[1:])
}

func Assemble(s string) ([]Instruction, error) {
	strs := strings.Split(s, "\n")
	var tStrs []string
	for _, s := range strs {
		ts := strings.Trim(s, " \t")
		if ts == "" {
			continue
		}
		tStrs = append(tStrs, ts)
	}
	return AssembleStrings(tStrs)
}

func AssembleStrings(strs []string) ([]Instruction, error) {
	var instructions []Instruction
	for lineNumber, s := range strs {
		i, err := AssembleOne(s)
		if err != nil {
			return nil, fmt.Errorf("line %d: %s", lineNumber, err)
		}
		instructions = append(instructions, i)
	}

	return instructions, nil
}
