package zog

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleProgram
	ruleBlankLine
	ruleLine
	ruleInstruction
	ruleAssignment
	ruleLoad
	ruleLoad8
	ruleLoad16
	rulePush
	rulePop
	ruleEx
	ruleInc
	ruleInc8
	ruleInc16
	ruleDec
	ruleDec8
	ruleDec16
	ruleAdd16
	ruleDst8
	ruleSrc8
	ruleLoc8
	ruleReg8
	ruleDst16
	ruleSrc16
	ruleLoc16
	ruleReg16
	rulen
	rulenn
	rulenn_contents
	ruleAlu
	ruleAdd
	ruleAdc
	ruleSub
	ruleSbc
	ruleAnd
	ruleXor
	ruleOr
	ruleCp
	ruleBitOp
	ruleRot
	ruleRlc
	ruleRrc
	ruleRl
	ruleRr
	ruleSla
	ruleSra
	ruleSll
	ruleSrl
	ruleBit
	ruleRes
	ruleSet
	ruleSimple
	ruleNop
	ruleHalt
	ruleRlca
	ruleRrca
	ruleRla
	ruleRra
	ruleDaa
	ruleCpl
	ruleScf
	ruleCcf
	ruleExx
	ruleDi
	ruleEi
	ruleJump
	ruleRst
	ruleCall
	ruleRet
	ruleJp
	ruleJr
	ruleDjnz
	ruledisp
	rulesep
	rulews
	ruleA
	ruleF
	ruleB
	ruleC
	ruleD
	ruleE
	ruleH
	ruleL
	ruleIXH
	ruleIXL
	ruleIYH
	ruleIYL
	ruleAF
	ruleAF_PRIME
	ruleBC
	ruleDE
	ruleHL
	ruleIX
	ruleIY
	ruleSP
	rulehexByteH
	rulehexByte0x
	ruledecimalByte
	rulehexWordH
	rulehexWord0x
	rulehexdigit
	ruleoctaldigit
	rulesignedDecimalByte
	rulecc
	ruleFT_NZ
	ruleFT_Z
	ruleFT_NC
	ruleFT_C
	ruleFT_PO
	ruleFT_PE
	ruleFT_P
	ruleFT_M
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
	ruleAction8
	ruleAction9
	ruleAction10
	ruleAction11
	ruleAction12
	ruleAction13
	rulePegText
	ruleAction14
	ruleAction15
	ruleAction16
	ruleAction17
	ruleAction18
	ruleAction19
	ruleAction20
	ruleAction21
	ruleAction22
	ruleAction23
	ruleAction24
	ruleAction25
	ruleAction26
	ruleAction27
	ruleAction28
	ruleAction29
	ruleAction30
	ruleAction31
	ruleAction32
	ruleAction33
	ruleAction34
	ruleAction35
	ruleAction36
	ruleAction37
	ruleAction38
	ruleAction39
	ruleAction40
	ruleAction41
	ruleAction42
	ruleAction43
	ruleAction44
	ruleAction45
	ruleAction46
	ruleAction47
	ruleAction48
	ruleAction49
	ruleAction50
	ruleAction51
	ruleAction52
	ruleAction53
	ruleAction54
	ruleAction55
	ruleAction56
	ruleAction57
	ruleAction58
	ruleAction59
	ruleAction60
	ruleAction61
	ruleAction62
	ruleAction63
	ruleAction64
	ruleAction65
	ruleAction66
	ruleAction67
	ruleAction68
	ruleAction69
	ruleAction70
	ruleAction71
	ruleAction72
)

var rul3s = [...]string{
	"Unknown",
	"Program",
	"BlankLine",
	"Line",
	"Instruction",
	"Assignment",
	"Load",
	"Load8",
	"Load16",
	"Push",
	"Pop",
	"Ex",
	"Inc",
	"Inc8",
	"Inc16",
	"Dec",
	"Dec8",
	"Dec16",
	"Add16",
	"Dst8",
	"Src8",
	"Loc8",
	"Reg8",
	"Dst16",
	"Src16",
	"Loc16",
	"Reg16",
	"n",
	"nn",
	"nn_contents",
	"Alu",
	"Add",
	"Adc",
	"Sub",
	"Sbc",
	"And",
	"Xor",
	"Or",
	"Cp",
	"BitOp",
	"Rot",
	"Rlc",
	"Rrc",
	"Rl",
	"Rr",
	"Sla",
	"Sra",
	"Sll",
	"Srl",
	"Bit",
	"Res",
	"Set",
	"Simple",
	"Nop",
	"Halt",
	"Rlca",
	"Rrca",
	"Rla",
	"Rra",
	"Daa",
	"Cpl",
	"Scf",
	"Ccf",
	"Exx",
	"Di",
	"Ei",
	"Jump",
	"Rst",
	"Call",
	"Ret",
	"Jp",
	"Jr",
	"Djnz",
	"disp",
	"sep",
	"ws",
	"A",
	"F",
	"B",
	"C",
	"D",
	"E",
	"H",
	"L",
	"IXH",
	"IXL",
	"IYH",
	"IYL",
	"AF",
	"AF_PRIME",
	"BC",
	"DE",
	"HL",
	"IX",
	"IY",
	"SP",
	"hexByteH",
	"hexByte0x",
	"decimalByte",
	"hexWordH",
	"hexWord0x",
	"hexdigit",
	"octaldigit",
	"signedDecimalByte",
	"cc",
	"FT_NZ",
	"FT_Z",
	"FT_NC",
	"FT_C",
	"FT_PO",
	"FT_PE",
	"FT_P",
	"FT_M",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
	"Action8",
	"Action9",
	"Action10",
	"Action11",
	"Action12",
	"Action13",
	"PegText",
	"Action14",
	"Action15",
	"Action16",
	"Action17",
	"Action18",
	"Action19",
	"Action20",
	"Action21",
	"Action22",
	"Action23",
	"Action24",
	"Action25",
	"Action26",
	"Action27",
	"Action28",
	"Action29",
	"Action30",
	"Action31",
	"Action32",
	"Action33",
	"Action34",
	"Action35",
	"Action36",
	"Action37",
	"Action38",
	"Action39",
	"Action40",
	"Action41",
	"Action42",
	"Action43",
	"Action44",
	"Action45",
	"Action46",
	"Action47",
	"Action48",
	"Action49",
	"Action50",
	"Action51",
	"Action52",
	"Action53",
	"Action54",
	"Action55",
	"Action56",
	"Action57",
	"Action58",
	"Action59",
	"Action60",
	"Action61",
	"Action62",
	"Action63",
	"Action64",
	"Action65",
	"Action66",
	"Action67",
	"Action68",
	"Action69",
	"Action70",
	"Action71",
	"Action72",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Printf(" ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Printf("%v %v\n", rule, quote)
			} else {
				fmt.Printf("\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(buffer string) {
	node.print(false, buffer)
}

func (node *node32) PrettyPrint(buffer string) {
	node.print(true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type PegAssembler struct {
	Current

	Buffer string
	buffer []rune
	rules  [187]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *PegAssembler) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *PegAssembler) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *PegAssembler
	max token32
}

func (e *parseError) Error() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}

func (p *PegAssembler) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *PegAssembler) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:
			p.Emit()
		case ruleAction1:
			p.LD8()
		case ruleAction2:
			p.LD16()
		case ruleAction3:
			p.Push()
		case ruleAction4:
			p.Pop()
		case ruleAction5:
			p.Ex()
		case ruleAction6:
			p.Inc8()
		case ruleAction7:
			p.Inc16()
		case ruleAction8:
			p.Dec8()
		case ruleAction9:
			p.Dec16()
		case ruleAction10:
			p.Add16()
		case ruleAction11:
			p.Dst8()
		case ruleAction12:
			p.Src8()
		case ruleAction13:
			p.Loc8()
		case ruleAction14:
			p.R8(buffer[begin:end])
		case ruleAction15:
			p.Dst16()
		case ruleAction16:
			p.Src16()
		case ruleAction17:
			p.Loc16()
		case ruleAction18:
			p.R16(buffer[begin:end])
		case ruleAction19:
			p.NNContents()
		case ruleAction20:
			p.Accum("ADD")
		case ruleAction21:
			p.Accum("ADC")
		case ruleAction22:
			p.Accum("SUB")
		case ruleAction23:
			p.Accum("SBC")
		case ruleAction24:
			p.Accum("AND")
		case ruleAction25:
			p.Accum("XOR")
		case ruleAction26:
			p.Accum("OR")
		case ruleAction27:
			p.Accum("CP")
		case ruleAction28:
			p.Rot("RLC")
		case ruleAction29:
			p.Rot("RRC")
		case ruleAction30:
			p.Rot("RL")
		case ruleAction31:
			p.Rot("RR")
		case ruleAction32:
			p.Rot("SLA")
		case ruleAction33:
			p.Rot("SRA")
		case ruleAction34:
			p.Rot("SLL")
		case ruleAction35:
			p.Rot("SRL")
		case ruleAction36:
			p.Bit()
		case ruleAction37:
			p.Res()
		case ruleAction38:
			p.Set()
		case ruleAction39:
			p.Simple(buffer[begin:end])
		case ruleAction40:
			p.Simple(buffer[begin:end])
		case ruleAction41:
			p.Simple(buffer[begin:end])
		case ruleAction42:
			p.Simple(buffer[begin:end])
		case ruleAction43:
			p.Simple(buffer[begin:end])
		case ruleAction44:
			p.Simple(buffer[begin:end])
		case ruleAction45:
			p.Simple(buffer[begin:end])
		case ruleAction46:
			p.Simple(buffer[begin:end])
		case ruleAction47:
			p.Simple(buffer[begin:end])
		case ruleAction48:
			p.Simple(buffer[begin:end])
		case ruleAction49:
			p.Simple(buffer[begin:end])
		case ruleAction50:
			p.Simple(buffer[begin:end])
		case ruleAction51:
			p.Simple(buffer[begin:end])
		case ruleAction52:
			p.Rst()
		case ruleAction53:
			p.Call()
		case ruleAction54:
			p.Ret()
		case ruleAction55:
			p.Jp()
		case ruleAction56:
			p.Jr()
		case ruleAction57:
			p.Djnz()
		case ruleAction58:
			p.Nhex(buffer[begin:end])
		case ruleAction59:
			p.Nhex(buffer[begin:end])
		case ruleAction60:
			p.Ndec(buffer[begin:end])
		case ruleAction61:
			p.NNhex(buffer[begin:end])
		case ruleAction62:
			p.NNhex(buffer[begin:end])
		case ruleAction63:
			p.ODigit(buffer[begin:end])
		case ruleAction64:
			p.SignedDecimalByte(buffer[begin:end])
		case ruleAction65:
			p.Conditional(Not{FT_Z})
		case ruleAction66:
			p.Conditional(FT_Z)
		case ruleAction67:
			p.Conditional(Not{FT_C})
		case ruleAction68:
			p.Conditional(FT_C)
		case ruleAction69:
			p.Conditional(FT_PO)
		case ruleAction70:
			p.Conditional(FT_PE)
		case ruleAction71:
			p.Conditional(FT_P)
		case ruleAction72:
			p.Conditional(FT_M)

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func (p *PegAssembler) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := tokens32{tree: make([]token32, math.MaxInt16)}
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 Program <- <(BlankLine / Line)+> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position4, tokenIndex4 := position, tokenIndex
					{
						position6 := position
					l7:
						{
							position8, tokenIndex8 := position, tokenIndex
							if !_rules[rulews]() {
								goto l8
							}
							goto l7
						l8:
							position, tokenIndex = position8, tokenIndex8
						}
						if buffer[position] != rune('\n') {
							goto l5
						}
						position++
						add(ruleBlankLine, position6)
					}
					goto l4
				l5:
					position, tokenIndex = position4, tokenIndex4
					{
						position9 := position
						{
							position10 := position
						l11:
							{
								position12, tokenIndex12 := position, tokenIndex
								if !_rules[rulews]() {
									goto l12
								}
								goto l11
							l12:
								position, tokenIndex = position12, tokenIndex12
							}
							{
								position13, tokenIndex13 := position, tokenIndex
								{
									position15 := position
									{
										position16, tokenIndex16 := position, tokenIndex
										{
											position18 := position
											{
												position19, tokenIndex19 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l20
												}
												position++
												goto l19
											l20:
												position, tokenIndex = position19, tokenIndex19
												if buffer[position] != rune('P') {
													goto l17
												}
												position++
											}
										l19:
											{
												position21, tokenIndex21 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l22
												}
												position++
												goto l21
											l22:
												position, tokenIndex = position21, tokenIndex21
												if buffer[position] != rune('U') {
													goto l17
												}
												position++
											}
										l21:
											{
												position23, tokenIndex23 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l24
												}
												position++
												goto l23
											l24:
												position, tokenIndex = position23, tokenIndex23
												if buffer[position] != rune('S') {
													goto l17
												}
												position++
											}
										l23:
											{
												position25, tokenIndex25 := position, tokenIndex
												if buffer[position] != rune('h') {
													goto l26
												}
												position++
												goto l25
											l26:
												position, tokenIndex = position25, tokenIndex25
												if buffer[position] != rune('H') {
													goto l17
												}
												position++
											}
										l25:
											if !_rules[rulews]() {
												goto l17
											}
											if !_rules[ruleSrc16]() {
												goto l17
											}
											{
												add(ruleAction3, position)
											}
											add(rulePush, position18)
										}
										goto l16
									l17:
										position, tokenIndex = position16, tokenIndex16
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position29 := position
													{
														position30, tokenIndex30 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l31
														}
														position++
														goto l30
													l31:
														position, tokenIndex = position30, tokenIndex30
														if buffer[position] != rune('E') {
															goto l14
														}
														position++
													}
												l30:
													{
														position32, tokenIndex32 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l33
														}
														position++
														goto l32
													l33:
														position, tokenIndex = position32, tokenIndex32
														if buffer[position] != rune('X') {
															goto l14
														}
														position++
													}
												l32:
													if !_rules[rulews]() {
														goto l14
													}
													if !_rules[ruleDst16]() {
														goto l14
													}
													if !_rules[rulesep]() {
														goto l14
													}
													if !_rules[ruleSrc16]() {
														goto l14
													}
													{
														add(ruleAction5, position)
													}
													add(ruleEx, position29)
												}
												break
											case 'P', 'p':
												{
													position35 := position
													{
														position36, tokenIndex36 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l37
														}
														position++
														goto l36
													l37:
														position, tokenIndex = position36, tokenIndex36
														if buffer[position] != rune('P') {
															goto l14
														}
														position++
													}
												l36:
													{
														position38, tokenIndex38 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l39
														}
														position++
														goto l38
													l39:
														position, tokenIndex = position38, tokenIndex38
														if buffer[position] != rune('O') {
															goto l14
														}
														position++
													}
												l38:
													{
														position40, tokenIndex40 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l41
														}
														position++
														goto l40
													l41:
														position, tokenIndex = position40, tokenIndex40
														if buffer[position] != rune('P') {
															goto l14
														}
														position++
													}
												l40:
													if !_rules[rulews]() {
														goto l14
													}
													if !_rules[ruleDst16]() {
														goto l14
													}
													{
														add(ruleAction4, position)
													}
													add(rulePop, position35)
												}
												break
											default:
												{
													position43 := position
													{
														position44, tokenIndex44 := position, tokenIndex
														{
															position46 := position
															{
																position47, tokenIndex47 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l48
																}
																position++
																goto l47
															l48:
																position, tokenIndex = position47, tokenIndex47
																if buffer[position] != rune('L') {
																	goto l45
																}
																position++
															}
														l47:
															{
																position49, tokenIndex49 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l50
																}
																position++
																goto l49
															l50:
																position, tokenIndex = position49, tokenIndex49
																if buffer[position] != rune('D') {
																	goto l45
																}
																position++
															}
														l49:
															if !_rules[rulews]() {
																goto l45
															}
															{
																position51 := position
																if !_rules[ruleReg8]() {
																	goto l45
																}
																{
																	add(ruleAction11, position)
																}
																add(ruleDst8, position51)
															}
															if !_rules[rulesep]() {
																goto l45
															}
															if !_rules[ruleSrc8]() {
																goto l45
															}
															{
																add(ruleAction1, position)
															}
															add(ruleLoad8, position46)
														}
														goto l44
													l45:
														position, tokenIndex = position44, tokenIndex44
														{
															position54 := position
															{
																position55, tokenIndex55 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l56
																}
																position++
																goto l55
															l56:
																position, tokenIndex = position55, tokenIndex55
																if buffer[position] != rune('L') {
																	goto l14
																}
																position++
															}
														l55:
															{
																position57, tokenIndex57 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l58
																}
																position++
																goto l57
															l58:
																position, tokenIndex = position57, tokenIndex57
																if buffer[position] != rune('D') {
																	goto l14
																}
																position++
															}
														l57:
															if !_rules[rulews]() {
																goto l14
															}
															if !_rules[ruleDst16]() {
																goto l14
															}
															if !_rules[rulesep]() {
																goto l14
															}
															if !_rules[ruleSrc16]() {
																goto l14
															}
															{
																add(ruleAction2, position)
															}
															add(ruleLoad16, position54)
														}
													}
												l44:
													add(ruleLoad, position43)
												}
												break
											}
										}

									}
								l16:
									add(ruleAssignment, position15)
								}
								goto l13
							l14:
								position, tokenIndex = position13, tokenIndex13
								{
									position61 := position
									{
										position62, tokenIndex62 := position, tokenIndex
										{
											position64 := position
											{
												position65, tokenIndex65 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l66
												}
												position++
												goto l65
											l66:
												position, tokenIndex = position65, tokenIndex65
												if buffer[position] != rune('I') {
													goto l63
												}
												position++
											}
										l65:
											{
												position67, tokenIndex67 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l68
												}
												position++
												goto l67
											l68:
												position, tokenIndex = position67, tokenIndex67
												if buffer[position] != rune('N') {
													goto l63
												}
												position++
											}
										l67:
											{
												position69, tokenIndex69 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l70
												}
												position++
												goto l69
											l70:
												position, tokenIndex = position69, tokenIndex69
												if buffer[position] != rune('C') {
													goto l63
												}
												position++
											}
										l69:
											if !_rules[rulews]() {
												goto l63
											}
											if !_rules[ruleLoc16]() {
												goto l63
											}
											{
												add(ruleAction7, position)
											}
											add(ruleInc16, position64)
										}
										goto l62
									l63:
										position, tokenIndex = position62, tokenIndex62
										{
											position72 := position
											{
												position73, tokenIndex73 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l74
												}
												position++
												goto l73
											l74:
												position, tokenIndex = position73, tokenIndex73
												if buffer[position] != rune('I') {
													goto l60
												}
												position++
											}
										l73:
											{
												position75, tokenIndex75 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l76
												}
												position++
												goto l75
											l76:
												position, tokenIndex = position75, tokenIndex75
												if buffer[position] != rune('N') {
													goto l60
												}
												position++
											}
										l75:
											{
												position77, tokenIndex77 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l78
												}
												position++
												goto l77
											l78:
												position, tokenIndex = position77, tokenIndex77
												if buffer[position] != rune('C') {
													goto l60
												}
												position++
											}
										l77:
											if !_rules[rulews]() {
												goto l60
											}
											if !_rules[ruleLoc8]() {
												goto l60
											}
											{
												add(ruleAction6, position)
											}
											add(ruleInc8, position72)
										}
									}
								l62:
									add(ruleInc, position61)
								}
								goto l13
							l60:
								position, tokenIndex = position13, tokenIndex13
								{
									position81 := position
									{
										position82, tokenIndex82 := position, tokenIndex
										{
											position84 := position
											{
												position85, tokenIndex85 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l86
												}
												position++
												goto l85
											l86:
												position, tokenIndex = position85, tokenIndex85
												if buffer[position] != rune('I') {
													goto l83
												}
												position++
											}
										l85:
											{
												position87, tokenIndex87 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l88
												}
												position++
												goto l87
											l88:
												position, tokenIndex = position87, tokenIndex87
												if buffer[position] != rune('N') {
													goto l83
												}
												position++
											}
										l87:
											{
												position89, tokenIndex89 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l90
												}
												position++
												goto l89
											l90:
												position, tokenIndex = position89, tokenIndex89
												if buffer[position] != rune('C') {
													goto l83
												}
												position++
											}
										l89:
											if !_rules[rulews]() {
												goto l83
											}
											if !_rules[ruleLoc16]() {
												goto l83
											}
											{
												add(ruleAction9, position)
											}
											add(ruleDec16, position84)
										}
										goto l82
									l83:
										position, tokenIndex = position82, tokenIndex82
										{
											position92 := position
											{
												position93, tokenIndex93 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l94
												}
												position++
												goto l93
											l94:
												position, tokenIndex = position93, tokenIndex93
												if buffer[position] != rune('I') {
													goto l80
												}
												position++
											}
										l93:
											{
												position95, tokenIndex95 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l96
												}
												position++
												goto l95
											l96:
												position, tokenIndex = position95, tokenIndex95
												if buffer[position] != rune('N') {
													goto l80
												}
												position++
											}
										l95:
											{
												position97, tokenIndex97 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l98
												}
												position++
												goto l97
											l98:
												position, tokenIndex = position97, tokenIndex97
												if buffer[position] != rune('C') {
													goto l80
												}
												position++
											}
										l97:
											if !_rules[rulews]() {
												goto l80
											}
											if !_rules[ruleLoc8]() {
												goto l80
											}
											{
												add(ruleAction8, position)
											}
											add(ruleDec8, position92)
										}
									}
								l82:
									add(ruleDec, position81)
								}
								goto l13
							l80:
								position, tokenIndex = position13, tokenIndex13
								{
									position101 := position
									{
										position102, tokenIndex102 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l103
										}
										position++
										goto l102
									l103:
										position, tokenIndex = position102, tokenIndex102
										if buffer[position] != rune('A') {
											goto l100
										}
										position++
									}
								l102:
									{
										position104, tokenIndex104 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l105
										}
										position++
										goto l104
									l105:
										position, tokenIndex = position104, tokenIndex104
										if buffer[position] != rune('D') {
											goto l100
										}
										position++
									}
								l104:
									{
										position106, tokenIndex106 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l107
										}
										position++
										goto l106
									l107:
										position, tokenIndex = position106, tokenIndex106
										if buffer[position] != rune('D') {
											goto l100
										}
										position++
									}
								l106:
									if !_rules[rulews]() {
										goto l100
									}
									if !_rules[ruleDst16]() {
										goto l100
									}
									if !_rules[rulesep]() {
										goto l100
									}
									if !_rules[ruleSrc16]() {
										goto l100
									}
									{
										add(ruleAction10, position)
									}
									add(ruleAdd16, position101)
								}
								goto l13
							l100:
								position, tokenIndex = position13, tokenIndex13
								{
									position110 := position
									{
										position111, tokenIndex111 := position, tokenIndex
										{
											position113 := position
											{
												position114, tokenIndex114 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l115
												}
												position++
												goto l114
											l115:
												position, tokenIndex = position114, tokenIndex114
												if buffer[position] != rune('A') {
													goto l112
												}
												position++
											}
										l114:
											{
												position116, tokenIndex116 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l117
												}
												position++
												goto l116
											l117:
												position, tokenIndex = position116, tokenIndex116
												if buffer[position] != rune('D') {
													goto l112
												}
												position++
											}
										l116:
											{
												position118, tokenIndex118 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l119
												}
												position++
												goto l118
											l119:
												position, tokenIndex = position118, tokenIndex118
												if buffer[position] != rune('D') {
													goto l112
												}
												position++
											}
										l118:
											if !_rules[rulews]() {
												goto l112
											}
											{
												position120, tokenIndex120 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l121
												}
												position++
												goto l120
											l121:
												position, tokenIndex = position120, tokenIndex120
												if buffer[position] != rune('A') {
													goto l112
												}
												position++
											}
										l120:
											if !_rules[rulesep]() {
												goto l112
											}
											if !_rules[ruleSrc8]() {
												goto l112
											}
											{
												add(ruleAction20, position)
											}
											add(ruleAdd, position113)
										}
										goto l111
									l112:
										position, tokenIndex = position111, tokenIndex111
										{
											position124 := position
											{
												position125, tokenIndex125 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l126
												}
												position++
												goto l125
											l126:
												position, tokenIndex = position125, tokenIndex125
												if buffer[position] != rune('A') {
													goto l123
												}
												position++
											}
										l125:
											{
												position127, tokenIndex127 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l128
												}
												position++
												goto l127
											l128:
												position, tokenIndex = position127, tokenIndex127
												if buffer[position] != rune('D') {
													goto l123
												}
												position++
											}
										l127:
											{
												position129, tokenIndex129 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l130
												}
												position++
												goto l129
											l130:
												position, tokenIndex = position129, tokenIndex129
												if buffer[position] != rune('C') {
													goto l123
												}
												position++
											}
										l129:
											if !_rules[rulews]() {
												goto l123
											}
											{
												position131, tokenIndex131 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l132
												}
												position++
												goto l131
											l132:
												position, tokenIndex = position131, tokenIndex131
												if buffer[position] != rune('A') {
													goto l123
												}
												position++
											}
										l131:
											if !_rules[rulesep]() {
												goto l123
											}
											if !_rules[ruleSrc8]() {
												goto l123
											}
											{
												add(ruleAction21, position)
											}
											add(ruleAdc, position124)
										}
										goto l111
									l123:
										position, tokenIndex = position111, tokenIndex111
										{
											position135 := position
											{
												position136, tokenIndex136 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l137
												}
												position++
												goto l136
											l137:
												position, tokenIndex = position136, tokenIndex136
												if buffer[position] != rune('S') {
													goto l134
												}
												position++
											}
										l136:
											{
												position138, tokenIndex138 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l139
												}
												position++
												goto l138
											l139:
												position, tokenIndex = position138, tokenIndex138
												if buffer[position] != rune('U') {
													goto l134
												}
												position++
											}
										l138:
											{
												position140, tokenIndex140 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l141
												}
												position++
												goto l140
											l141:
												position, tokenIndex = position140, tokenIndex140
												if buffer[position] != rune('B') {
													goto l134
												}
												position++
											}
										l140:
											if !_rules[rulews]() {
												goto l134
											}
											if !_rules[ruleSrc8]() {
												goto l134
											}
											{
												add(ruleAction22, position)
											}
											add(ruleSub, position135)
										}
										goto l111
									l134:
										position, tokenIndex = position111, tokenIndex111
										{
											switch buffer[position] {
											case 'C', 'c':
												{
													position144 := position
													{
														position145, tokenIndex145 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l146
														}
														position++
														goto l145
													l146:
														position, tokenIndex = position145, tokenIndex145
														if buffer[position] != rune('C') {
															goto l109
														}
														position++
													}
												l145:
													{
														position147, tokenIndex147 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l148
														}
														position++
														goto l147
													l148:
														position, tokenIndex = position147, tokenIndex147
														if buffer[position] != rune('P') {
															goto l109
														}
														position++
													}
												l147:
													if !_rules[rulews]() {
														goto l109
													}
													if !_rules[ruleSrc8]() {
														goto l109
													}
													{
														add(ruleAction27, position)
													}
													add(ruleCp, position144)
												}
												break
											case 'O', 'o':
												{
													position150 := position
													{
														position151, tokenIndex151 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l152
														}
														position++
														goto l151
													l152:
														position, tokenIndex = position151, tokenIndex151
														if buffer[position] != rune('O') {
															goto l109
														}
														position++
													}
												l151:
													{
														position153, tokenIndex153 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l154
														}
														position++
														goto l153
													l154:
														position, tokenIndex = position153, tokenIndex153
														if buffer[position] != rune('R') {
															goto l109
														}
														position++
													}
												l153:
													if !_rules[rulews]() {
														goto l109
													}
													if !_rules[ruleSrc8]() {
														goto l109
													}
													{
														add(ruleAction26, position)
													}
													add(ruleOr, position150)
												}
												break
											case 'X', 'x':
												{
													position156 := position
													{
														position157, tokenIndex157 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l158
														}
														position++
														goto l157
													l158:
														position, tokenIndex = position157, tokenIndex157
														if buffer[position] != rune('X') {
															goto l109
														}
														position++
													}
												l157:
													{
														position159, tokenIndex159 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l160
														}
														position++
														goto l159
													l160:
														position, tokenIndex = position159, tokenIndex159
														if buffer[position] != rune('O') {
															goto l109
														}
														position++
													}
												l159:
													{
														position161, tokenIndex161 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l162
														}
														position++
														goto l161
													l162:
														position, tokenIndex = position161, tokenIndex161
														if buffer[position] != rune('R') {
															goto l109
														}
														position++
													}
												l161:
													if !_rules[rulews]() {
														goto l109
													}
													if !_rules[ruleSrc8]() {
														goto l109
													}
													{
														add(ruleAction25, position)
													}
													add(ruleXor, position156)
												}
												break
											case 'A', 'a':
												{
													position164 := position
													{
														position165, tokenIndex165 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l166
														}
														position++
														goto l165
													l166:
														position, tokenIndex = position165, tokenIndex165
														if buffer[position] != rune('A') {
															goto l109
														}
														position++
													}
												l165:
													{
														position167, tokenIndex167 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l168
														}
														position++
														goto l167
													l168:
														position, tokenIndex = position167, tokenIndex167
														if buffer[position] != rune('N') {
															goto l109
														}
														position++
													}
												l167:
													{
														position169, tokenIndex169 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l170
														}
														position++
														goto l169
													l170:
														position, tokenIndex = position169, tokenIndex169
														if buffer[position] != rune('D') {
															goto l109
														}
														position++
													}
												l169:
													if !_rules[rulews]() {
														goto l109
													}
													if !_rules[ruleSrc8]() {
														goto l109
													}
													{
														add(ruleAction24, position)
													}
													add(ruleAnd, position164)
												}
												break
											default:
												{
													position172 := position
													{
														position173, tokenIndex173 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l174
														}
														position++
														goto l173
													l174:
														position, tokenIndex = position173, tokenIndex173
														if buffer[position] != rune('S') {
															goto l109
														}
														position++
													}
												l173:
													{
														position175, tokenIndex175 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l176
														}
														position++
														goto l175
													l176:
														position, tokenIndex = position175, tokenIndex175
														if buffer[position] != rune('B') {
															goto l109
														}
														position++
													}
												l175:
													{
														position177, tokenIndex177 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l178
														}
														position++
														goto l177
													l178:
														position, tokenIndex = position177, tokenIndex177
														if buffer[position] != rune('C') {
															goto l109
														}
														position++
													}
												l177:
													if !_rules[rulews]() {
														goto l109
													}
													{
														position179, tokenIndex179 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l180
														}
														position++
														goto l179
													l180:
														position, tokenIndex = position179, tokenIndex179
														if buffer[position] != rune('A') {
															goto l109
														}
														position++
													}
												l179:
													if !_rules[rulesep]() {
														goto l109
													}
													if !_rules[ruleSrc8]() {
														goto l109
													}
													{
														add(ruleAction23, position)
													}
													add(ruleSbc, position172)
												}
												break
											}
										}

									}
								l111:
									add(ruleAlu, position110)
								}
								goto l13
							l109:
								position, tokenIndex = position13, tokenIndex13
								{
									position183 := position
									{
										position184, tokenIndex184 := position, tokenIndex
										{
											position186 := position
											{
												position187, tokenIndex187 := position, tokenIndex
												{
													position189 := position
													{
														position190, tokenIndex190 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l191
														}
														position++
														goto l190
													l191:
														position, tokenIndex = position190, tokenIndex190
														if buffer[position] != rune('R') {
															goto l188
														}
														position++
													}
												l190:
													{
														position192, tokenIndex192 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l193
														}
														position++
														goto l192
													l193:
														position, tokenIndex = position192, tokenIndex192
														if buffer[position] != rune('L') {
															goto l188
														}
														position++
													}
												l192:
													{
														position194, tokenIndex194 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l195
														}
														position++
														goto l194
													l195:
														position, tokenIndex = position194, tokenIndex194
														if buffer[position] != rune('C') {
															goto l188
														}
														position++
													}
												l194:
													if !_rules[rulews]() {
														goto l188
													}
													if !_rules[ruleLoc8]() {
														goto l188
													}
													{
														add(ruleAction28, position)
													}
													add(ruleRlc, position189)
												}
												goto l187
											l188:
												position, tokenIndex = position187, tokenIndex187
												{
													position198 := position
													{
														position199, tokenIndex199 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l200
														}
														position++
														goto l199
													l200:
														position, tokenIndex = position199, tokenIndex199
														if buffer[position] != rune('R') {
															goto l197
														}
														position++
													}
												l199:
													{
														position201, tokenIndex201 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l202
														}
														position++
														goto l201
													l202:
														position, tokenIndex = position201, tokenIndex201
														if buffer[position] != rune('R') {
															goto l197
														}
														position++
													}
												l201:
													{
														position203, tokenIndex203 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l204
														}
														position++
														goto l203
													l204:
														position, tokenIndex = position203, tokenIndex203
														if buffer[position] != rune('C') {
															goto l197
														}
														position++
													}
												l203:
													if !_rules[rulews]() {
														goto l197
													}
													if !_rules[ruleLoc8]() {
														goto l197
													}
													{
														add(ruleAction29, position)
													}
													add(ruleRrc, position198)
												}
												goto l187
											l197:
												position, tokenIndex = position187, tokenIndex187
												{
													position207 := position
													{
														position208, tokenIndex208 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l209
														}
														position++
														goto l208
													l209:
														position, tokenIndex = position208, tokenIndex208
														if buffer[position] != rune('R') {
															goto l206
														}
														position++
													}
												l208:
													{
														position210, tokenIndex210 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l211
														}
														position++
														goto l210
													l211:
														position, tokenIndex = position210, tokenIndex210
														if buffer[position] != rune('L') {
															goto l206
														}
														position++
													}
												l210:
													if !_rules[rulews]() {
														goto l206
													}
													if !_rules[ruleLoc8]() {
														goto l206
													}
													{
														add(ruleAction30, position)
													}
													add(ruleRl, position207)
												}
												goto l187
											l206:
												position, tokenIndex = position187, tokenIndex187
												{
													position214 := position
													{
														position215, tokenIndex215 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l216
														}
														position++
														goto l215
													l216:
														position, tokenIndex = position215, tokenIndex215
														if buffer[position] != rune('R') {
															goto l213
														}
														position++
													}
												l215:
													{
														position217, tokenIndex217 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l218
														}
														position++
														goto l217
													l218:
														position, tokenIndex = position217, tokenIndex217
														if buffer[position] != rune('R') {
															goto l213
														}
														position++
													}
												l217:
													if !_rules[rulews]() {
														goto l213
													}
													if !_rules[ruleLoc8]() {
														goto l213
													}
													{
														add(ruleAction31, position)
													}
													add(ruleRr, position214)
												}
												goto l187
											l213:
												position, tokenIndex = position187, tokenIndex187
												{
													position221 := position
													{
														position222, tokenIndex222 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l223
														}
														position++
														goto l222
													l223:
														position, tokenIndex = position222, tokenIndex222
														if buffer[position] != rune('S') {
															goto l220
														}
														position++
													}
												l222:
													{
														position224, tokenIndex224 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l225
														}
														position++
														goto l224
													l225:
														position, tokenIndex = position224, tokenIndex224
														if buffer[position] != rune('L') {
															goto l220
														}
														position++
													}
												l224:
													{
														position226, tokenIndex226 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l227
														}
														position++
														goto l226
													l227:
														position, tokenIndex = position226, tokenIndex226
														if buffer[position] != rune('A') {
															goto l220
														}
														position++
													}
												l226:
													if !_rules[rulews]() {
														goto l220
													}
													if !_rules[ruleLoc8]() {
														goto l220
													}
													{
														add(ruleAction32, position)
													}
													add(ruleSla, position221)
												}
												goto l187
											l220:
												position, tokenIndex = position187, tokenIndex187
												{
													position230 := position
													{
														position231, tokenIndex231 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l232
														}
														position++
														goto l231
													l232:
														position, tokenIndex = position231, tokenIndex231
														if buffer[position] != rune('S') {
															goto l229
														}
														position++
													}
												l231:
													{
														position233, tokenIndex233 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l234
														}
														position++
														goto l233
													l234:
														position, tokenIndex = position233, tokenIndex233
														if buffer[position] != rune('R') {
															goto l229
														}
														position++
													}
												l233:
													{
														position235, tokenIndex235 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l236
														}
														position++
														goto l235
													l236:
														position, tokenIndex = position235, tokenIndex235
														if buffer[position] != rune('A') {
															goto l229
														}
														position++
													}
												l235:
													if !_rules[rulews]() {
														goto l229
													}
													if !_rules[ruleLoc8]() {
														goto l229
													}
													{
														add(ruleAction33, position)
													}
													add(ruleSra, position230)
												}
												goto l187
											l229:
												position, tokenIndex = position187, tokenIndex187
												{
													position239 := position
													{
														position240, tokenIndex240 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l241
														}
														position++
														goto l240
													l241:
														position, tokenIndex = position240, tokenIndex240
														if buffer[position] != rune('S') {
															goto l238
														}
														position++
													}
												l240:
													{
														position242, tokenIndex242 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l243
														}
														position++
														goto l242
													l243:
														position, tokenIndex = position242, tokenIndex242
														if buffer[position] != rune('L') {
															goto l238
														}
														position++
													}
												l242:
													{
														position244, tokenIndex244 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l245
														}
														position++
														goto l244
													l245:
														position, tokenIndex = position244, tokenIndex244
														if buffer[position] != rune('L') {
															goto l238
														}
														position++
													}
												l244:
													if !_rules[rulews]() {
														goto l238
													}
													if !_rules[ruleLoc8]() {
														goto l238
													}
													{
														add(ruleAction34, position)
													}
													add(ruleSll, position239)
												}
												goto l187
											l238:
												position, tokenIndex = position187, tokenIndex187
												{
													position247 := position
													{
														position248, tokenIndex248 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l249
														}
														position++
														goto l248
													l249:
														position, tokenIndex = position248, tokenIndex248
														if buffer[position] != rune('S') {
															goto l185
														}
														position++
													}
												l248:
													{
														position250, tokenIndex250 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l251
														}
														position++
														goto l250
													l251:
														position, tokenIndex = position250, tokenIndex250
														if buffer[position] != rune('R') {
															goto l185
														}
														position++
													}
												l250:
													{
														position252, tokenIndex252 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l253
														}
														position++
														goto l252
													l253:
														position, tokenIndex = position252, tokenIndex252
														if buffer[position] != rune('L') {
															goto l185
														}
														position++
													}
												l252:
													if !_rules[rulews]() {
														goto l185
													}
													if !_rules[ruleLoc8]() {
														goto l185
													}
													{
														add(ruleAction35, position)
													}
													add(ruleSrl, position247)
												}
											}
										l187:
											add(ruleRot, position186)
										}
										goto l184
									l185:
										position, tokenIndex = position184, tokenIndex184
										{
											switch buffer[position] {
											case 'S', 's':
												{
													position256 := position
													{
														position257, tokenIndex257 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l258
														}
														position++
														goto l257
													l258:
														position, tokenIndex = position257, tokenIndex257
														if buffer[position] != rune('S') {
															goto l182
														}
														position++
													}
												l257:
													{
														position259, tokenIndex259 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l260
														}
														position++
														goto l259
													l260:
														position, tokenIndex = position259, tokenIndex259
														if buffer[position] != rune('E') {
															goto l182
														}
														position++
													}
												l259:
													{
														position261, tokenIndex261 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l262
														}
														position++
														goto l261
													l262:
														position, tokenIndex = position261, tokenIndex261
														if buffer[position] != rune('T') {
															goto l182
														}
														position++
													}
												l261:
													if !_rules[rulews]() {
														goto l182
													}
													if !_rules[ruleoctaldigit]() {
														goto l182
													}
													if !_rules[rulesep]() {
														goto l182
													}
													if !_rules[ruleLoc8]() {
														goto l182
													}
													{
														add(ruleAction38, position)
													}
													add(ruleSet, position256)
												}
												break
											case 'R', 'r':
												{
													position264 := position
													{
														position265, tokenIndex265 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l266
														}
														position++
														goto l265
													l266:
														position, tokenIndex = position265, tokenIndex265
														if buffer[position] != rune('R') {
															goto l182
														}
														position++
													}
												l265:
													{
														position267, tokenIndex267 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l268
														}
														position++
														goto l267
													l268:
														position, tokenIndex = position267, tokenIndex267
														if buffer[position] != rune('E') {
															goto l182
														}
														position++
													}
												l267:
													{
														position269, tokenIndex269 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l270
														}
														position++
														goto l269
													l270:
														position, tokenIndex = position269, tokenIndex269
														if buffer[position] != rune('S') {
															goto l182
														}
														position++
													}
												l269:
													if !_rules[rulews]() {
														goto l182
													}
													if !_rules[ruleoctaldigit]() {
														goto l182
													}
													if !_rules[rulesep]() {
														goto l182
													}
													if !_rules[ruleLoc8]() {
														goto l182
													}
													{
														add(ruleAction37, position)
													}
													add(ruleRes, position264)
												}
												break
											default:
												{
													position272 := position
													{
														position273, tokenIndex273 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l274
														}
														position++
														goto l273
													l274:
														position, tokenIndex = position273, tokenIndex273
														if buffer[position] != rune('B') {
															goto l182
														}
														position++
													}
												l273:
													{
														position275, tokenIndex275 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l276
														}
														position++
														goto l275
													l276:
														position, tokenIndex = position275, tokenIndex275
														if buffer[position] != rune('I') {
															goto l182
														}
														position++
													}
												l275:
													{
														position277, tokenIndex277 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l278
														}
														position++
														goto l277
													l278:
														position, tokenIndex = position277, tokenIndex277
														if buffer[position] != rune('T') {
															goto l182
														}
														position++
													}
												l277:
													if !_rules[rulews]() {
														goto l182
													}
													if !_rules[ruleoctaldigit]() {
														goto l182
													}
													if !_rules[rulesep]() {
														goto l182
													}
													if !_rules[ruleLoc8]() {
														goto l182
													}
													{
														add(ruleAction36, position)
													}
													add(ruleBit, position272)
												}
												break
											}
										}

									}
								l184:
									add(ruleBitOp, position183)
								}
								goto l13
							l182:
								position, tokenIndex = position13, tokenIndex13
								{
									position281 := position
									{
										position282, tokenIndex282 := position, tokenIndex
										{
											position284 := position
											{
												position285 := position
												{
													position286, tokenIndex286 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l287
													}
													position++
													goto l286
												l287:
													position, tokenIndex = position286, tokenIndex286
													if buffer[position] != rune('R') {
														goto l283
													}
													position++
												}
											l286:
												{
													position288, tokenIndex288 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l289
													}
													position++
													goto l288
												l289:
													position, tokenIndex = position288, tokenIndex288
													if buffer[position] != rune('L') {
														goto l283
													}
													position++
												}
											l288:
												{
													position290, tokenIndex290 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l291
													}
													position++
													goto l290
												l291:
													position, tokenIndex = position290, tokenIndex290
													if buffer[position] != rune('C') {
														goto l283
													}
													position++
												}
											l290:
												{
													position292, tokenIndex292 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l293
													}
													position++
													goto l292
												l293:
													position, tokenIndex = position292, tokenIndex292
													if buffer[position] != rune('A') {
														goto l283
													}
													position++
												}
											l292:
												add(rulePegText, position285)
											}
											{
												add(ruleAction41, position)
											}
											add(ruleRlca, position284)
										}
										goto l282
									l283:
										position, tokenIndex = position282, tokenIndex282
										{
											position296 := position
											{
												position297 := position
												{
													position298, tokenIndex298 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l299
													}
													position++
													goto l298
												l299:
													position, tokenIndex = position298, tokenIndex298
													if buffer[position] != rune('R') {
														goto l295
													}
													position++
												}
											l298:
												{
													position300, tokenIndex300 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l301
													}
													position++
													goto l300
												l301:
													position, tokenIndex = position300, tokenIndex300
													if buffer[position] != rune('R') {
														goto l295
													}
													position++
												}
											l300:
												{
													position302, tokenIndex302 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l303
													}
													position++
													goto l302
												l303:
													position, tokenIndex = position302, tokenIndex302
													if buffer[position] != rune('C') {
														goto l295
													}
													position++
												}
											l302:
												{
													position304, tokenIndex304 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l305
													}
													position++
													goto l304
												l305:
													position, tokenIndex = position304, tokenIndex304
													if buffer[position] != rune('A') {
														goto l295
													}
													position++
												}
											l304:
												add(rulePegText, position297)
											}
											{
												add(ruleAction42, position)
											}
											add(ruleRrca, position296)
										}
										goto l282
									l295:
										position, tokenIndex = position282, tokenIndex282
										{
											position308 := position
											{
												position309 := position
												{
													position310, tokenIndex310 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l311
													}
													position++
													goto l310
												l311:
													position, tokenIndex = position310, tokenIndex310
													if buffer[position] != rune('R') {
														goto l307
													}
													position++
												}
											l310:
												{
													position312, tokenIndex312 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l313
													}
													position++
													goto l312
												l313:
													position, tokenIndex = position312, tokenIndex312
													if buffer[position] != rune('L') {
														goto l307
													}
													position++
												}
											l312:
												{
													position314, tokenIndex314 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l315
													}
													position++
													goto l314
												l315:
													position, tokenIndex = position314, tokenIndex314
													if buffer[position] != rune('A') {
														goto l307
													}
													position++
												}
											l314:
												add(rulePegText, position309)
											}
											{
												add(ruleAction43, position)
											}
											add(ruleRla, position308)
										}
										goto l282
									l307:
										position, tokenIndex = position282, tokenIndex282
										{
											position318 := position
											{
												position319 := position
												{
													position320, tokenIndex320 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l321
													}
													position++
													goto l320
												l321:
													position, tokenIndex = position320, tokenIndex320
													if buffer[position] != rune('D') {
														goto l317
													}
													position++
												}
											l320:
												{
													position322, tokenIndex322 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l323
													}
													position++
													goto l322
												l323:
													position, tokenIndex = position322, tokenIndex322
													if buffer[position] != rune('A') {
														goto l317
													}
													position++
												}
											l322:
												{
													position324, tokenIndex324 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l325
													}
													position++
													goto l324
												l325:
													position, tokenIndex = position324, tokenIndex324
													if buffer[position] != rune('A') {
														goto l317
													}
													position++
												}
											l324:
												add(rulePegText, position319)
											}
											{
												add(ruleAction45, position)
											}
											add(ruleDaa, position318)
										}
										goto l282
									l317:
										position, tokenIndex = position282, tokenIndex282
										{
											position328 := position
											{
												position329 := position
												{
													position330, tokenIndex330 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l331
													}
													position++
													goto l330
												l331:
													position, tokenIndex = position330, tokenIndex330
													if buffer[position] != rune('C') {
														goto l327
													}
													position++
												}
											l330:
												{
													position332, tokenIndex332 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l333
													}
													position++
													goto l332
												l333:
													position, tokenIndex = position332, tokenIndex332
													if buffer[position] != rune('P') {
														goto l327
													}
													position++
												}
											l332:
												{
													position334, tokenIndex334 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l335
													}
													position++
													goto l334
												l335:
													position, tokenIndex = position334, tokenIndex334
													if buffer[position] != rune('L') {
														goto l327
													}
													position++
												}
											l334:
												add(rulePegText, position329)
											}
											{
												add(ruleAction46, position)
											}
											add(ruleCpl, position328)
										}
										goto l282
									l327:
										position, tokenIndex = position282, tokenIndex282
										{
											position338 := position
											{
												position339 := position
												{
													position340, tokenIndex340 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l341
													}
													position++
													goto l340
												l341:
													position, tokenIndex = position340, tokenIndex340
													if buffer[position] != rune('E') {
														goto l337
													}
													position++
												}
											l340:
												{
													position342, tokenIndex342 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l343
													}
													position++
													goto l342
												l343:
													position, tokenIndex = position342, tokenIndex342
													if buffer[position] != rune('X') {
														goto l337
													}
													position++
												}
											l342:
												{
													position344, tokenIndex344 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l345
													}
													position++
													goto l344
												l345:
													position, tokenIndex = position344, tokenIndex344
													if buffer[position] != rune('X') {
														goto l337
													}
													position++
												}
											l344:
												add(rulePegText, position339)
											}
											{
												add(ruleAction49, position)
											}
											add(ruleExx, position338)
										}
										goto l282
									l337:
										position, tokenIndex = position282, tokenIndex282
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position348 := position
													{
														position349 := position
														{
															position350, tokenIndex350 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l351
															}
															position++
															goto l350
														l351:
															position, tokenIndex = position350, tokenIndex350
															if buffer[position] != rune('E') {
																goto l280
															}
															position++
														}
													l350:
														{
															position352, tokenIndex352 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l353
															}
															position++
															goto l352
														l353:
															position, tokenIndex = position352, tokenIndex352
															if buffer[position] != rune('I') {
																goto l280
															}
															position++
														}
													l352:
														add(rulePegText, position349)
													}
													{
														add(ruleAction51, position)
													}
													add(ruleEi, position348)
												}
												break
											case 'D', 'd':
												{
													position355 := position
													{
														position356 := position
														{
															position357, tokenIndex357 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l358
															}
															position++
															goto l357
														l358:
															position, tokenIndex = position357, tokenIndex357
															if buffer[position] != rune('D') {
																goto l280
															}
															position++
														}
													l357:
														{
															position359, tokenIndex359 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l360
															}
															position++
															goto l359
														l360:
															position, tokenIndex = position359, tokenIndex359
															if buffer[position] != rune('I') {
																goto l280
															}
															position++
														}
													l359:
														add(rulePegText, position356)
													}
													{
														add(ruleAction50, position)
													}
													add(ruleDi, position355)
												}
												break
											case 'C', 'c':
												{
													position362 := position
													{
														position363 := position
														{
															position364, tokenIndex364 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l365
															}
															position++
															goto l364
														l365:
															position, tokenIndex = position364, tokenIndex364
															if buffer[position] != rune('C') {
																goto l280
															}
															position++
														}
													l364:
														{
															position366, tokenIndex366 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l367
															}
															position++
															goto l366
														l367:
															position, tokenIndex = position366, tokenIndex366
															if buffer[position] != rune('C') {
																goto l280
															}
															position++
														}
													l366:
														{
															position368, tokenIndex368 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l369
															}
															position++
															goto l368
														l369:
															position, tokenIndex = position368, tokenIndex368
															if buffer[position] != rune('F') {
																goto l280
															}
															position++
														}
													l368:
														add(rulePegText, position363)
													}
													{
														add(ruleAction48, position)
													}
													add(ruleCcf, position362)
												}
												break
											case 'S', 's':
												{
													position371 := position
													{
														position372 := position
														{
															position373, tokenIndex373 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l374
															}
															position++
															goto l373
														l374:
															position, tokenIndex = position373, tokenIndex373
															if buffer[position] != rune('S') {
																goto l280
															}
															position++
														}
													l373:
														{
															position375, tokenIndex375 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l376
															}
															position++
															goto l375
														l376:
															position, tokenIndex = position375, tokenIndex375
															if buffer[position] != rune('C') {
																goto l280
															}
															position++
														}
													l375:
														{
															position377, tokenIndex377 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l378
															}
															position++
															goto l377
														l378:
															position, tokenIndex = position377, tokenIndex377
															if buffer[position] != rune('F') {
																goto l280
															}
															position++
														}
													l377:
														add(rulePegText, position372)
													}
													{
														add(ruleAction47, position)
													}
													add(ruleScf, position371)
												}
												break
											case 'R', 'r':
												{
													position380 := position
													{
														position381 := position
														{
															position382, tokenIndex382 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l383
															}
															position++
															goto l382
														l383:
															position, tokenIndex = position382, tokenIndex382
															if buffer[position] != rune('R') {
																goto l280
															}
															position++
														}
													l382:
														{
															position384, tokenIndex384 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l385
															}
															position++
															goto l384
														l385:
															position, tokenIndex = position384, tokenIndex384
															if buffer[position] != rune('R') {
																goto l280
															}
															position++
														}
													l384:
														{
															position386, tokenIndex386 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l387
															}
															position++
															goto l386
														l387:
															position, tokenIndex = position386, tokenIndex386
															if buffer[position] != rune('A') {
																goto l280
															}
															position++
														}
													l386:
														add(rulePegText, position381)
													}
													{
														add(ruleAction44, position)
													}
													add(ruleRra, position380)
												}
												break
											case 'H', 'h':
												{
													position389 := position
													{
														position390 := position
														{
															position391, tokenIndex391 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l392
															}
															position++
															goto l391
														l392:
															position, tokenIndex = position391, tokenIndex391
															if buffer[position] != rune('H') {
																goto l280
															}
															position++
														}
													l391:
														{
															position393, tokenIndex393 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l394
															}
															position++
															goto l393
														l394:
															position, tokenIndex = position393, tokenIndex393
															if buffer[position] != rune('A') {
																goto l280
															}
															position++
														}
													l393:
														{
															position395, tokenIndex395 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l396
															}
															position++
															goto l395
														l396:
															position, tokenIndex = position395, tokenIndex395
															if buffer[position] != rune('L') {
																goto l280
															}
															position++
														}
													l395:
														{
															position397, tokenIndex397 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l398
															}
															position++
															goto l397
														l398:
															position, tokenIndex = position397, tokenIndex397
															if buffer[position] != rune('T') {
																goto l280
															}
															position++
														}
													l397:
														add(rulePegText, position390)
													}
													{
														add(ruleAction40, position)
													}
													add(ruleHalt, position389)
												}
												break
											default:
												{
													position400 := position
													{
														position401 := position
														{
															position402, tokenIndex402 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l403
															}
															position++
															goto l402
														l403:
															position, tokenIndex = position402, tokenIndex402
															if buffer[position] != rune('N') {
																goto l280
															}
															position++
														}
													l402:
														{
															position404, tokenIndex404 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l405
															}
															position++
															goto l404
														l405:
															position, tokenIndex = position404, tokenIndex404
															if buffer[position] != rune('O') {
																goto l280
															}
															position++
														}
													l404:
														{
															position406, tokenIndex406 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l407
															}
															position++
															goto l406
														l407:
															position, tokenIndex = position406, tokenIndex406
															if buffer[position] != rune('P') {
																goto l280
															}
															position++
														}
													l406:
														add(rulePegText, position401)
													}
													{
														add(ruleAction39, position)
													}
													add(ruleNop, position400)
												}
												break
											}
										}

									}
								l282:
									add(ruleSimple, position281)
								}
								goto l13
							l280:
								position, tokenIndex = position13, tokenIndex13
								{
									position409 := position
									{
										position410, tokenIndex410 := position, tokenIndex
										{
											position412 := position
											{
												position413, tokenIndex413 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l414
												}
												position++
												goto l413
											l414:
												position, tokenIndex = position413, tokenIndex413
												if buffer[position] != rune('R') {
													goto l411
												}
												position++
											}
										l413:
											{
												position415, tokenIndex415 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l416
												}
												position++
												goto l415
											l416:
												position, tokenIndex = position415, tokenIndex415
												if buffer[position] != rune('S') {
													goto l411
												}
												position++
											}
										l415:
											{
												position417, tokenIndex417 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l418
												}
												position++
												goto l417
											l418:
												position, tokenIndex = position417, tokenIndex417
												if buffer[position] != rune('T') {
													goto l411
												}
												position++
											}
										l417:
											if !_rules[rulews]() {
												goto l411
											}
											if !_rules[rulen]() {
												goto l411
											}
											{
												add(ruleAction52, position)
											}
											add(ruleRst, position412)
										}
										goto l410
									l411:
										position, tokenIndex = position410, tokenIndex410
										{
											position421 := position
											{
												position422, tokenIndex422 := position, tokenIndex
												if buffer[position] != rune('j') {
													goto l423
												}
												position++
												goto l422
											l423:
												position, tokenIndex = position422, tokenIndex422
												if buffer[position] != rune('J') {
													goto l420
												}
												position++
											}
										l422:
											{
												position424, tokenIndex424 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l425
												}
												position++
												goto l424
											l425:
												position, tokenIndex = position424, tokenIndex424
												if buffer[position] != rune('P') {
													goto l420
												}
												position++
											}
										l424:
											if !_rules[rulews]() {
												goto l420
											}
											{
												position426, tokenIndex426 := position, tokenIndex
												if !_rules[rulecc]() {
													goto l426
												}
												if !_rules[rulesep]() {
													goto l426
												}
												goto l427
											l426:
												position, tokenIndex = position426, tokenIndex426
											}
										l427:
											if !_rules[ruleSrc16]() {
												goto l420
											}
											{
												add(ruleAction55, position)
											}
											add(ruleJp, position421)
										}
										goto l410
									l420:
										position, tokenIndex = position410, tokenIndex410
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position430 := position
													{
														position431, tokenIndex431 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l432
														}
														position++
														goto l431
													l432:
														position, tokenIndex = position431, tokenIndex431
														if buffer[position] != rune('D') {
															goto l0
														}
														position++
													}
												l431:
													{
														position433, tokenIndex433 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l434
														}
														position++
														goto l433
													l434:
														position, tokenIndex = position433, tokenIndex433
														if buffer[position] != rune('J') {
															goto l0
														}
														position++
													}
												l433:
													{
														position435, tokenIndex435 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l436
														}
														position++
														goto l435
													l436:
														position, tokenIndex = position435, tokenIndex435
														if buffer[position] != rune('N') {
															goto l0
														}
														position++
													}
												l435:
													{
														position437, tokenIndex437 := position, tokenIndex
														if buffer[position] != rune('z') {
															goto l438
														}
														position++
														goto l437
													l438:
														position, tokenIndex = position437, tokenIndex437
														if buffer[position] != rune('Z') {
															goto l0
														}
														position++
													}
												l437:
													if !_rules[rulews]() {
														goto l0
													}
													if !_rules[ruledisp]() {
														goto l0
													}
													{
														add(ruleAction57, position)
													}
													add(ruleDjnz, position430)
												}
												break
											case 'J', 'j':
												{
													position440 := position
													{
														position441, tokenIndex441 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l442
														}
														position++
														goto l441
													l442:
														position, tokenIndex = position441, tokenIndex441
														if buffer[position] != rune('J') {
															goto l0
														}
														position++
													}
												l441:
													{
														position443, tokenIndex443 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l444
														}
														position++
														goto l443
													l444:
														position, tokenIndex = position443, tokenIndex443
														if buffer[position] != rune('R') {
															goto l0
														}
														position++
													}
												l443:
													if !_rules[rulews]() {
														goto l0
													}
													{
														position445, tokenIndex445 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l445
														}
														if !_rules[rulesep]() {
															goto l445
														}
														goto l446
													l445:
														position, tokenIndex = position445, tokenIndex445
													}
												l446:
													if !_rules[ruledisp]() {
														goto l0
													}
													{
														add(ruleAction56, position)
													}
													add(ruleJr, position440)
												}
												break
											case 'R', 'r':
												{
													position448 := position
													{
														position449, tokenIndex449 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l450
														}
														position++
														goto l449
													l450:
														position, tokenIndex = position449, tokenIndex449
														if buffer[position] != rune('R') {
															goto l0
														}
														position++
													}
												l449:
													{
														position451, tokenIndex451 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l452
														}
														position++
														goto l451
													l452:
														position, tokenIndex = position451, tokenIndex451
														if buffer[position] != rune('E') {
															goto l0
														}
														position++
													}
												l451:
													{
														position453, tokenIndex453 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l454
														}
														position++
														goto l453
													l454:
														position, tokenIndex = position453, tokenIndex453
														if buffer[position] != rune('T') {
															goto l0
														}
														position++
													}
												l453:
													{
														position455, tokenIndex455 := position, tokenIndex
														if !_rules[rulews]() {
															goto l455
														}
														if !_rules[rulecc]() {
															goto l455
														}
														goto l456
													l455:
														position, tokenIndex = position455, tokenIndex455
													}
												l456:
													{
														add(ruleAction54, position)
													}
													add(ruleRet, position448)
												}
												break
											default:
												{
													position458 := position
													{
														position459, tokenIndex459 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l460
														}
														position++
														goto l459
													l460:
														position, tokenIndex = position459, tokenIndex459
														if buffer[position] != rune('C') {
															goto l0
														}
														position++
													}
												l459:
													{
														position461, tokenIndex461 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l462
														}
														position++
														goto l461
													l462:
														position, tokenIndex = position461, tokenIndex461
														if buffer[position] != rune('A') {
															goto l0
														}
														position++
													}
												l461:
													{
														position463, tokenIndex463 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l464
														}
														position++
														goto l463
													l464:
														position, tokenIndex = position463, tokenIndex463
														if buffer[position] != rune('L') {
															goto l0
														}
														position++
													}
												l463:
													{
														position465, tokenIndex465 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l466
														}
														position++
														goto l465
													l466:
														position, tokenIndex = position465, tokenIndex465
														if buffer[position] != rune('L') {
															goto l0
														}
														position++
													}
												l465:
													if !_rules[rulews]() {
														goto l0
													}
													{
														position467, tokenIndex467 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l467
														}
														if !_rules[rulesep]() {
															goto l467
														}
														goto l468
													l467:
														position, tokenIndex = position467, tokenIndex467
													}
												l468:
													if !_rules[ruleSrc16]() {
														goto l0
													}
													{
														add(ruleAction53, position)
													}
													add(ruleCall, position458)
												}
												break
											}
										}

									}
								l410:
									add(ruleJump, position409)
								}
							}
						l13:
							add(ruleInstruction, position10)
						}
						{
							position470, tokenIndex470 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l470
							}
							position++
							goto l471
						l470:
							position, tokenIndex = position470, tokenIndex470
						}
					l471:
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position9)
					}
				}
			l4:
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position473, tokenIndex473 := position, tokenIndex
						{
							position475 := position
						l476:
							{
								position477, tokenIndex477 := position, tokenIndex
								if !_rules[rulews]() {
									goto l477
								}
								goto l476
							l477:
								position, tokenIndex = position477, tokenIndex477
							}
							if buffer[position] != rune('\n') {
								goto l474
							}
							position++
							add(ruleBlankLine, position475)
						}
						goto l473
					l474:
						position, tokenIndex = position473, tokenIndex473
						{
							position478 := position
							{
								position479 := position
							l480:
								{
									position481, tokenIndex481 := position, tokenIndex
									if !_rules[rulews]() {
										goto l481
									}
									goto l480
								l481:
									position, tokenIndex = position481, tokenIndex481
								}
								{
									position482, tokenIndex482 := position, tokenIndex
									{
										position484 := position
										{
											position485, tokenIndex485 := position, tokenIndex
											{
												position487 := position
												{
													position488, tokenIndex488 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l489
													}
													position++
													goto l488
												l489:
													position, tokenIndex = position488, tokenIndex488
													if buffer[position] != rune('P') {
														goto l486
													}
													position++
												}
											l488:
												{
													position490, tokenIndex490 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l491
													}
													position++
													goto l490
												l491:
													position, tokenIndex = position490, tokenIndex490
													if buffer[position] != rune('U') {
														goto l486
													}
													position++
												}
											l490:
												{
													position492, tokenIndex492 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l493
													}
													position++
													goto l492
												l493:
													position, tokenIndex = position492, tokenIndex492
													if buffer[position] != rune('S') {
														goto l486
													}
													position++
												}
											l492:
												{
													position494, tokenIndex494 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l495
													}
													position++
													goto l494
												l495:
													position, tokenIndex = position494, tokenIndex494
													if buffer[position] != rune('H') {
														goto l486
													}
													position++
												}
											l494:
												if !_rules[rulews]() {
													goto l486
												}
												if !_rules[ruleSrc16]() {
													goto l486
												}
												{
													add(ruleAction3, position)
												}
												add(rulePush, position487)
											}
											goto l485
										l486:
											position, tokenIndex = position485, tokenIndex485
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position498 := position
														{
															position499, tokenIndex499 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l500
															}
															position++
															goto l499
														l500:
															position, tokenIndex = position499, tokenIndex499
															if buffer[position] != rune('E') {
																goto l483
															}
															position++
														}
													l499:
														{
															position501, tokenIndex501 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l502
															}
															position++
															goto l501
														l502:
															position, tokenIndex = position501, tokenIndex501
															if buffer[position] != rune('X') {
																goto l483
															}
															position++
														}
													l501:
														if !_rules[rulews]() {
															goto l483
														}
														if !_rules[ruleDst16]() {
															goto l483
														}
														if !_rules[rulesep]() {
															goto l483
														}
														if !_rules[ruleSrc16]() {
															goto l483
														}
														{
															add(ruleAction5, position)
														}
														add(ruleEx, position498)
													}
													break
												case 'P', 'p':
													{
														position504 := position
														{
															position505, tokenIndex505 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l506
															}
															position++
															goto l505
														l506:
															position, tokenIndex = position505, tokenIndex505
															if buffer[position] != rune('P') {
																goto l483
															}
															position++
														}
													l505:
														{
															position507, tokenIndex507 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l508
															}
															position++
															goto l507
														l508:
															position, tokenIndex = position507, tokenIndex507
															if buffer[position] != rune('O') {
																goto l483
															}
															position++
														}
													l507:
														{
															position509, tokenIndex509 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l510
															}
															position++
															goto l509
														l510:
															position, tokenIndex = position509, tokenIndex509
															if buffer[position] != rune('P') {
																goto l483
															}
															position++
														}
													l509:
														if !_rules[rulews]() {
															goto l483
														}
														if !_rules[ruleDst16]() {
															goto l483
														}
														{
															add(ruleAction4, position)
														}
														add(rulePop, position504)
													}
													break
												default:
													{
														position512 := position
														{
															position513, tokenIndex513 := position, tokenIndex
															{
																position515 := position
																{
																	position516, tokenIndex516 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l517
																	}
																	position++
																	goto l516
																l517:
																	position, tokenIndex = position516, tokenIndex516
																	if buffer[position] != rune('L') {
																		goto l514
																	}
																	position++
																}
															l516:
																{
																	position518, tokenIndex518 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l519
																	}
																	position++
																	goto l518
																l519:
																	position, tokenIndex = position518, tokenIndex518
																	if buffer[position] != rune('D') {
																		goto l514
																	}
																	position++
																}
															l518:
																if !_rules[rulews]() {
																	goto l514
																}
																{
																	position520 := position
																	if !_rules[ruleReg8]() {
																		goto l514
																	}
																	{
																		add(ruleAction11, position)
																	}
																	add(ruleDst8, position520)
																}
																if !_rules[rulesep]() {
																	goto l514
																}
																if !_rules[ruleSrc8]() {
																	goto l514
																}
																{
																	add(ruleAction1, position)
																}
																add(ruleLoad8, position515)
															}
															goto l513
														l514:
															position, tokenIndex = position513, tokenIndex513
															{
																position523 := position
																{
																	position524, tokenIndex524 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l525
																	}
																	position++
																	goto l524
																l525:
																	position, tokenIndex = position524, tokenIndex524
																	if buffer[position] != rune('L') {
																		goto l483
																	}
																	position++
																}
															l524:
																{
																	position526, tokenIndex526 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l527
																	}
																	position++
																	goto l526
																l527:
																	position, tokenIndex = position526, tokenIndex526
																	if buffer[position] != rune('D') {
																		goto l483
																	}
																	position++
																}
															l526:
																if !_rules[rulews]() {
																	goto l483
																}
																if !_rules[ruleDst16]() {
																	goto l483
																}
																if !_rules[rulesep]() {
																	goto l483
																}
																if !_rules[ruleSrc16]() {
																	goto l483
																}
																{
																	add(ruleAction2, position)
																}
																add(ruleLoad16, position523)
															}
														}
													l513:
														add(ruleLoad, position512)
													}
													break
												}
											}

										}
									l485:
										add(ruleAssignment, position484)
									}
									goto l482
								l483:
									position, tokenIndex = position482, tokenIndex482
									{
										position530 := position
										{
											position531, tokenIndex531 := position, tokenIndex
											{
												position533 := position
												{
													position534, tokenIndex534 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l535
													}
													position++
													goto l534
												l535:
													position, tokenIndex = position534, tokenIndex534
													if buffer[position] != rune('I') {
														goto l532
													}
													position++
												}
											l534:
												{
													position536, tokenIndex536 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l537
													}
													position++
													goto l536
												l537:
													position, tokenIndex = position536, tokenIndex536
													if buffer[position] != rune('N') {
														goto l532
													}
													position++
												}
											l536:
												{
													position538, tokenIndex538 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l539
													}
													position++
													goto l538
												l539:
													position, tokenIndex = position538, tokenIndex538
													if buffer[position] != rune('C') {
														goto l532
													}
													position++
												}
											l538:
												if !_rules[rulews]() {
													goto l532
												}
												if !_rules[ruleLoc16]() {
													goto l532
												}
												{
													add(ruleAction7, position)
												}
												add(ruleInc16, position533)
											}
											goto l531
										l532:
											position, tokenIndex = position531, tokenIndex531
											{
												position541 := position
												{
													position542, tokenIndex542 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l543
													}
													position++
													goto l542
												l543:
													position, tokenIndex = position542, tokenIndex542
													if buffer[position] != rune('I') {
														goto l529
													}
													position++
												}
											l542:
												{
													position544, tokenIndex544 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l545
													}
													position++
													goto l544
												l545:
													position, tokenIndex = position544, tokenIndex544
													if buffer[position] != rune('N') {
														goto l529
													}
													position++
												}
											l544:
												{
													position546, tokenIndex546 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l547
													}
													position++
													goto l546
												l547:
													position, tokenIndex = position546, tokenIndex546
													if buffer[position] != rune('C') {
														goto l529
													}
													position++
												}
											l546:
												if !_rules[rulews]() {
													goto l529
												}
												if !_rules[ruleLoc8]() {
													goto l529
												}
												{
													add(ruleAction6, position)
												}
												add(ruleInc8, position541)
											}
										}
									l531:
										add(ruleInc, position530)
									}
									goto l482
								l529:
									position, tokenIndex = position482, tokenIndex482
									{
										position550 := position
										{
											position551, tokenIndex551 := position, tokenIndex
											{
												position553 := position
												{
													position554, tokenIndex554 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l555
													}
													position++
													goto l554
												l555:
													position, tokenIndex = position554, tokenIndex554
													if buffer[position] != rune('I') {
														goto l552
													}
													position++
												}
											l554:
												{
													position556, tokenIndex556 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l557
													}
													position++
													goto l556
												l557:
													position, tokenIndex = position556, tokenIndex556
													if buffer[position] != rune('N') {
														goto l552
													}
													position++
												}
											l556:
												{
													position558, tokenIndex558 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l559
													}
													position++
													goto l558
												l559:
													position, tokenIndex = position558, tokenIndex558
													if buffer[position] != rune('C') {
														goto l552
													}
													position++
												}
											l558:
												if !_rules[rulews]() {
													goto l552
												}
												if !_rules[ruleLoc16]() {
													goto l552
												}
												{
													add(ruleAction9, position)
												}
												add(ruleDec16, position553)
											}
											goto l551
										l552:
											position, tokenIndex = position551, tokenIndex551
											{
												position561 := position
												{
													position562, tokenIndex562 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l563
													}
													position++
													goto l562
												l563:
													position, tokenIndex = position562, tokenIndex562
													if buffer[position] != rune('I') {
														goto l549
													}
													position++
												}
											l562:
												{
													position564, tokenIndex564 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l565
													}
													position++
													goto l564
												l565:
													position, tokenIndex = position564, tokenIndex564
													if buffer[position] != rune('N') {
														goto l549
													}
													position++
												}
											l564:
												{
													position566, tokenIndex566 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l567
													}
													position++
													goto l566
												l567:
													position, tokenIndex = position566, tokenIndex566
													if buffer[position] != rune('C') {
														goto l549
													}
													position++
												}
											l566:
												if !_rules[rulews]() {
													goto l549
												}
												if !_rules[ruleLoc8]() {
													goto l549
												}
												{
													add(ruleAction8, position)
												}
												add(ruleDec8, position561)
											}
										}
									l551:
										add(ruleDec, position550)
									}
									goto l482
								l549:
									position, tokenIndex = position482, tokenIndex482
									{
										position570 := position
										{
											position571, tokenIndex571 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l572
											}
											position++
											goto l571
										l572:
											position, tokenIndex = position571, tokenIndex571
											if buffer[position] != rune('A') {
												goto l569
											}
											position++
										}
									l571:
										{
											position573, tokenIndex573 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l574
											}
											position++
											goto l573
										l574:
											position, tokenIndex = position573, tokenIndex573
											if buffer[position] != rune('D') {
												goto l569
											}
											position++
										}
									l573:
										{
											position575, tokenIndex575 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l576
											}
											position++
											goto l575
										l576:
											position, tokenIndex = position575, tokenIndex575
											if buffer[position] != rune('D') {
												goto l569
											}
											position++
										}
									l575:
										if !_rules[rulews]() {
											goto l569
										}
										if !_rules[ruleDst16]() {
											goto l569
										}
										if !_rules[rulesep]() {
											goto l569
										}
										if !_rules[ruleSrc16]() {
											goto l569
										}
										{
											add(ruleAction10, position)
										}
										add(ruleAdd16, position570)
									}
									goto l482
								l569:
									position, tokenIndex = position482, tokenIndex482
									{
										position579 := position
										{
											position580, tokenIndex580 := position, tokenIndex
											{
												position582 := position
												{
													position583, tokenIndex583 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l584
													}
													position++
													goto l583
												l584:
													position, tokenIndex = position583, tokenIndex583
													if buffer[position] != rune('A') {
														goto l581
													}
													position++
												}
											l583:
												{
													position585, tokenIndex585 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l586
													}
													position++
													goto l585
												l586:
													position, tokenIndex = position585, tokenIndex585
													if buffer[position] != rune('D') {
														goto l581
													}
													position++
												}
											l585:
												{
													position587, tokenIndex587 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l588
													}
													position++
													goto l587
												l588:
													position, tokenIndex = position587, tokenIndex587
													if buffer[position] != rune('D') {
														goto l581
													}
													position++
												}
											l587:
												if !_rules[rulews]() {
													goto l581
												}
												{
													position589, tokenIndex589 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l590
													}
													position++
													goto l589
												l590:
													position, tokenIndex = position589, tokenIndex589
													if buffer[position] != rune('A') {
														goto l581
													}
													position++
												}
											l589:
												if !_rules[rulesep]() {
													goto l581
												}
												if !_rules[ruleSrc8]() {
													goto l581
												}
												{
													add(ruleAction20, position)
												}
												add(ruleAdd, position582)
											}
											goto l580
										l581:
											position, tokenIndex = position580, tokenIndex580
											{
												position593 := position
												{
													position594, tokenIndex594 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l595
													}
													position++
													goto l594
												l595:
													position, tokenIndex = position594, tokenIndex594
													if buffer[position] != rune('A') {
														goto l592
													}
													position++
												}
											l594:
												{
													position596, tokenIndex596 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l597
													}
													position++
													goto l596
												l597:
													position, tokenIndex = position596, tokenIndex596
													if buffer[position] != rune('D') {
														goto l592
													}
													position++
												}
											l596:
												{
													position598, tokenIndex598 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l599
													}
													position++
													goto l598
												l599:
													position, tokenIndex = position598, tokenIndex598
													if buffer[position] != rune('C') {
														goto l592
													}
													position++
												}
											l598:
												if !_rules[rulews]() {
													goto l592
												}
												{
													position600, tokenIndex600 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l601
													}
													position++
													goto l600
												l601:
													position, tokenIndex = position600, tokenIndex600
													if buffer[position] != rune('A') {
														goto l592
													}
													position++
												}
											l600:
												if !_rules[rulesep]() {
													goto l592
												}
												if !_rules[ruleSrc8]() {
													goto l592
												}
												{
													add(ruleAction21, position)
												}
												add(ruleAdc, position593)
											}
											goto l580
										l592:
											position, tokenIndex = position580, tokenIndex580
											{
												position604 := position
												{
													position605, tokenIndex605 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l606
													}
													position++
													goto l605
												l606:
													position, tokenIndex = position605, tokenIndex605
													if buffer[position] != rune('S') {
														goto l603
													}
													position++
												}
											l605:
												{
													position607, tokenIndex607 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l608
													}
													position++
													goto l607
												l608:
													position, tokenIndex = position607, tokenIndex607
													if buffer[position] != rune('U') {
														goto l603
													}
													position++
												}
											l607:
												{
													position609, tokenIndex609 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l610
													}
													position++
													goto l609
												l610:
													position, tokenIndex = position609, tokenIndex609
													if buffer[position] != rune('B') {
														goto l603
													}
													position++
												}
											l609:
												if !_rules[rulews]() {
													goto l603
												}
												if !_rules[ruleSrc8]() {
													goto l603
												}
												{
													add(ruleAction22, position)
												}
												add(ruleSub, position604)
											}
											goto l580
										l603:
											position, tokenIndex = position580, tokenIndex580
											{
												switch buffer[position] {
												case 'C', 'c':
													{
														position613 := position
														{
															position614, tokenIndex614 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l615
															}
															position++
															goto l614
														l615:
															position, tokenIndex = position614, tokenIndex614
															if buffer[position] != rune('C') {
																goto l578
															}
															position++
														}
													l614:
														{
															position616, tokenIndex616 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l617
															}
															position++
															goto l616
														l617:
															position, tokenIndex = position616, tokenIndex616
															if buffer[position] != rune('P') {
																goto l578
															}
															position++
														}
													l616:
														if !_rules[rulews]() {
															goto l578
														}
														if !_rules[ruleSrc8]() {
															goto l578
														}
														{
															add(ruleAction27, position)
														}
														add(ruleCp, position613)
													}
													break
												case 'O', 'o':
													{
														position619 := position
														{
															position620, tokenIndex620 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l621
															}
															position++
															goto l620
														l621:
															position, tokenIndex = position620, tokenIndex620
															if buffer[position] != rune('O') {
																goto l578
															}
															position++
														}
													l620:
														{
															position622, tokenIndex622 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l623
															}
															position++
															goto l622
														l623:
															position, tokenIndex = position622, tokenIndex622
															if buffer[position] != rune('R') {
																goto l578
															}
															position++
														}
													l622:
														if !_rules[rulews]() {
															goto l578
														}
														if !_rules[ruleSrc8]() {
															goto l578
														}
														{
															add(ruleAction26, position)
														}
														add(ruleOr, position619)
													}
													break
												case 'X', 'x':
													{
														position625 := position
														{
															position626, tokenIndex626 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l627
															}
															position++
															goto l626
														l627:
															position, tokenIndex = position626, tokenIndex626
															if buffer[position] != rune('X') {
																goto l578
															}
															position++
														}
													l626:
														{
															position628, tokenIndex628 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l629
															}
															position++
															goto l628
														l629:
															position, tokenIndex = position628, tokenIndex628
															if buffer[position] != rune('O') {
																goto l578
															}
															position++
														}
													l628:
														{
															position630, tokenIndex630 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l631
															}
															position++
															goto l630
														l631:
															position, tokenIndex = position630, tokenIndex630
															if buffer[position] != rune('R') {
																goto l578
															}
															position++
														}
													l630:
														if !_rules[rulews]() {
															goto l578
														}
														if !_rules[ruleSrc8]() {
															goto l578
														}
														{
															add(ruleAction25, position)
														}
														add(ruleXor, position625)
													}
													break
												case 'A', 'a':
													{
														position633 := position
														{
															position634, tokenIndex634 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l635
															}
															position++
															goto l634
														l635:
															position, tokenIndex = position634, tokenIndex634
															if buffer[position] != rune('A') {
																goto l578
															}
															position++
														}
													l634:
														{
															position636, tokenIndex636 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l637
															}
															position++
															goto l636
														l637:
															position, tokenIndex = position636, tokenIndex636
															if buffer[position] != rune('N') {
																goto l578
															}
															position++
														}
													l636:
														{
															position638, tokenIndex638 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l639
															}
															position++
															goto l638
														l639:
															position, tokenIndex = position638, tokenIndex638
															if buffer[position] != rune('D') {
																goto l578
															}
															position++
														}
													l638:
														if !_rules[rulews]() {
															goto l578
														}
														if !_rules[ruleSrc8]() {
															goto l578
														}
														{
															add(ruleAction24, position)
														}
														add(ruleAnd, position633)
													}
													break
												default:
													{
														position641 := position
														{
															position642, tokenIndex642 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l643
															}
															position++
															goto l642
														l643:
															position, tokenIndex = position642, tokenIndex642
															if buffer[position] != rune('S') {
																goto l578
															}
															position++
														}
													l642:
														{
															position644, tokenIndex644 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l645
															}
															position++
															goto l644
														l645:
															position, tokenIndex = position644, tokenIndex644
															if buffer[position] != rune('B') {
																goto l578
															}
															position++
														}
													l644:
														{
															position646, tokenIndex646 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l647
															}
															position++
															goto l646
														l647:
															position, tokenIndex = position646, tokenIndex646
															if buffer[position] != rune('C') {
																goto l578
															}
															position++
														}
													l646:
														if !_rules[rulews]() {
															goto l578
														}
														{
															position648, tokenIndex648 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l649
															}
															position++
															goto l648
														l649:
															position, tokenIndex = position648, tokenIndex648
															if buffer[position] != rune('A') {
																goto l578
															}
															position++
														}
													l648:
														if !_rules[rulesep]() {
															goto l578
														}
														if !_rules[ruleSrc8]() {
															goto l578
														}
														{
															add(ruleAction23, position)
														}
														add(ruleSbc, position641)
													}
													break
												}
											}

										}
									l580:
										add(ruleAlu, position579)
									}
									goto l482
								l578:
									position, tokenIndex = position482, tokenIndex482
									{
										position652 := position
										{
											position653, tokenIndex653 := position, tokenIndex
											{
												position655 := position
												{
													position656, tokenIndex656 := position, tokenIndex
													{
														position658 := position
														{
															position659, tokenIndex659 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l660
															}
															position++
															goto l659
														l660:
															position, tokenIndex = position659, tokenIndex659
															if buffer[position] != rune('R') {
																goto l657
															}
															position++
														}
													l659:
														{
															position661, tokenIndex661 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l662
															}
															position++
															goto l661
														l662:
															position, tokenIndex = position661, tokenIndex661
															if buffer[position] != rune('L') {
																goto l657
															}
															position++
														}
													l661:
														{
															position663, tokenIndex663 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l664
															}
															position++
															goto l663
														l664:
															position, tokenIndex = position663, tokenIndex663
															if buffer[position] != rune('C') {
																goto l657
															}
															position++
														}
													l663:
														if !_rules[rulews]() {
															goto l657
														}
														if !_rules[ruleLoc8]() {
															goto l657
														}
														{
															add(ruleAction28, position)
														}
														add(ruleRlc, position658)
													}
													goto l656
												l657:
													position, tokenIndex = position656, tokenIndex656
													{
														position667 := position
														{
															position668, tokenIndex668 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l669
															}
															position++
															goto l668
														l669:
															position, tokenIndex = position668, tokenIndex668
															if buffer[position] != rune('R') {
																goto l666
															}
															position++
														}
													l668:
														{
															position670, tokenIndex670 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l671
															}
															position++
															goto l670
														l671:
															position, tokenIndex = position670, tokenIndex670
															if buffer[position] != rune('R') {
																goto l666
															}
															position++
														}
													l670:
														{
															position672, tokenIndex672 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l673
															}
															position++
															goto l672
														l673:
															position, tokenIndex = position672, tokenIndex672
															if buffer[position] != rune('C') {
																goto l666
															}
															position++
														}
													l672:
														if !_rules[rulews]() {
															goto l666
														}
														if !_rules[ruleLoc8]() {
															goto l666
														}
														{
															add(ruleAction29, position)
														}
														add(ruleRrc, position667)
													}
													goto l656
												l666:
													position, tokenIndex = position656, tokenIndex656
													{
														position676 := position
														{
															position677, tokenIndex677 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l678
															}
															position++
															goto l677
														l678:
															position, tokenIndex = position677, tokenIndex677
															if buffer[position] != rune('R') {
																goto l675
															}
															position++
														}
													l677:
														{
															position679, tokenIndex679 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l680
															}
															position++
															goto l679
														l680:
															position, tokenIndex = position679, tokenIndex679
															if buffer[position] != rune('L') {
																goto l675
															}
															position++
														}
													l679:
														if !_rules[rulews]() {
															goto l675
														}
														if !_rules[ruleLoc8]() {
															goto l675
														}
														{
															add(ruleAction30, position)
														}
														add(ruleRl, position676)
													}
													goto l656
												l675:
													position, tokenIndex = position656, tokenIndex656
													{
														position683 := position
														{
															position684, tokenIndex684 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l685
															}
															position++
															goto l684
														l685:
															position, tokenIndex = position684, tokenIndex684
															if buffer[position] != rune('R') {
																goto l682
															}
															position++
														}
													l684:
														{
															position686, tokenIndex686 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l687
															}
															position++
															goto l686
														l687:
															position, tokenIndex = position686, tokenIndex686
															if buffer[position] != rune('R') {
																goto l682
															}
															position++
														}
													l686:
														if !_rules[rulews]() {
															goto l682
														}
														if !_rules[ruleLoc8]() {
															goto l682
														}
														{
															add(ruleAction31, position)
														}
														add(ruleRr, position683)
													}
													goto l656
												l682:
													position, tokenIndex = position656, tokenIndex656
													{
														position690 := position
														{
															position691, tokenIndex691 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l692
															}
															position++
															goto l691
														l692:
															position, tokenIndex = position691, tokenIndex691
															if buffer[position] != rune('S') {
																goto l689
															}
															position++
														}
													l691:
														{
															position693, tokenIndex693 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l694
															}
															position++
															goto l693
														l694:
															position, tokenIndex = position693, tokenIndex693
															if buffer[position] != rune('L') {
																goto l689
															}
															position++
														}
													l693:
														{
															position695, tokenIndex695 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l696
															}
															position++
															goto l695
														l696:
															position, tokenIndex = position695, tokenIndex695
															if buffer[position] != rune('A') {
																goto l689
															}
															position++
														}
													l695:
														if !_rules[rulews]() {
															goto l689
														}
														if !_rules[ruleLoc8]() {
															goto l689
														}
														{
															add(ruleAction32, position)
														}
														add(ruleSla, position690)
													}
													goto l656
												l689:
													position, tokenIndex = position656, tokenIndex656
													{
														position699 := position
														{
															position700, tokenIndex700 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l701
															}
															position++
															goto l700
														l701:
															position, tokenIndex = position700, tokenIndex700
															if buffer[position] != rune('S') {
																goto l698
															}
															position++
														}
													l700:
														{
															position702, tokenIndex702 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l703
															}
															position++
															goto l702
														l703:
															position, tokenIndex = position702, tokenIndex702
															if buffer[position] != rune('R') {
																goto l698
															}
															position++
														}
													l702:
														{
															position704, tokenIndex704 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l705
															}
															position++
															goto l704
														l705:
															position, tokenIndex = position704, tokenIndex704
															if buffer[position] != rune('A') {
																goto l698
															}
															position++
														}
													l704:
														if !_rules[rulews]() {
															goto l698
														}
														if !_rules[ruleLoc8]() {
															goto l698
														}
														{
															add(ruleAction33, position)
														}
														add(ruleSra, position699)
													}
													goto l656
												l698:
													position, tokenIndex = position656, tokenIndex656
													{
														position708 := position
														{
															position709, tokenIndex709 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l710
															}
															position++
															goto l709
														l710:
															position, tokenIndex = position709, tokenIndex709
															if buffer[position] != rune('S') {
																goto l707
															}
															position++
														}
													l709:
														{
															position711, tokenIndex711 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l712
															}
															position++
															goto l711
														l712:
															position, tokenIndex = position711, tokenIndex711
															if buffer[position] != rune('L') {
																goto l707
															}
															position++
														}
													l711:
														{
															position713, tokenIndex713 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l714
															}
															position++
															goto l713
														l714:
															position, tokenIndex = position713, tokenIndex713
															if buffer[position] != rune('L') {
																goto l707
															}
															position++
														}
													l713:
														if !_rules[rulews]() {
															goto l707
														}
														if !_rules[ruleLoc8]() {
															goto l707
														}
														{
															add(ruleAction34, position)
														}
														add(ruleSll, position708)
													}
													goto l656
												l707:
													position, tokenIndex = position656, tokenIndex656
													{
														position716 := position
														{
															position717, tokenIndex717 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l718
															}
															position++
															goto l717
														l718:
															position, tokenIndex = position717, tokenIndex717
															if buffer[position] != rune('S') {
																goto l654
															}
															position++
														}
													l717:
														{
															position719, tokenIndex719 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l720
															}
															position++
															goto l719
														l720:
															position, tokenIndex = position719, tokenIndex719
															if buffer[position] != rune('R') {
																goto l654
															}
															position++
														}
													l719:
														{
															position721, tokenIndex721 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l722
															}
															position++
															goto l721
														l722:
															position, tokenIndex = position721, tokenIndex721
															if buffer[position] != rune('L') {
																goto l654
															}
															position++
														}
													l721:
														if !_rules[rulews]() {
															goto l654
														}
														if !_rules[ruleLoc8]() {
															goto l654
														}
														{
															add(ruleAction35, position)
														}
														add(ruleSrl, position716)
													}
												}
											l656:
												add(ruleRot, position655)
											}
											goto l653
										l654:
											position, tokenIndex = position653, tokenIndex653
											{
												switch buffer[position] {
												case 'S', 's':
													{
														position725 := position
														{
															position726, tokenIndex726 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l727
															}
															position++
															goto l726
														l727:
															position, tokenIndex = position726, tokenIndex726
															if buffer[position] != rune('S') {
																goto l651
															}
															position++
														}
													l726:
														{
															position728, tokenIndex728 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l729
															}
															position++
															goto l728
														l729:
															position, tokenIndex = position728, tokenIndex728
															if buffer[position] != rune('E') {
																goto l651
															}
															position++
														}
													l728:
														{
															position730, tokenIndex730 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l731
															}
															position++
															goto l730
														l731:
															position, tokenIndex = position730, tokenIndex730
															if buffer[position] != rune('T') {
																goto l651
															}
															position++
														}
													l730:
														if !_rules[rulews]() {
															goto l651
														}
														if !_rules[ruleoctaldigit]() {
															goto l651
														}
														if !_rules[rulesep]() {
															goto l651
														}
														if !_rules[ruleLoc8]() {
															goto l651
														}
														{
															add(ruleAction38, position)
														}
														add(ruleSet, position725)
													}
													break
												case 'R', 'r':
													{
														position733 := position
														{
															position734, tokenIndex734 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l735
															}
															position++
															goto l734
														l735:
															position, tokenIndex = position734, tokenIndex734
															if buffer[position] != rune('R') {
																goto l651
															}
															position++
														}
													l734:
														{
															position736, tokenIndex736 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l737
															}
															position++
															goto l736
														l737:
															position, tokenIndex = position736, tokenIndex736
															if buffer[position] != rune('E') {
																goto l651
															}
															position++
														}
													l736:
														{
															position738, tokenIndex738 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l739
															}
															position++
															goto l738
														l739:
															position, tokenIndex = position738, tokenIndex738
															if buffer[position] != rune('S') {
																goto l651
															}
															position++
														}
													l738:
														if !_rules[rulews]() {
															goto l651
														}
														if !_rules[ruleoctaldigit]() {
															goto l651
														}
														if !_rules[rulesep]() {
															goto l651
														}
														if !_rules[ruleLoc8]() {
															goto l651
														}
														{
															add(ruleAction37, position)
														}
														add(ruleRes, position733)
													}
													break
												default:
													{
														position741 := position
														{
															position742, tokenIndex742 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l743
															}
															position++
															goto l742
														l743:
															position, tokenIndex = position742, tokenIndex742
															if buffer[position] != rune('B') {
																goto l651
															}
															position++
														}
													l742:
														{
															position744, tokenIndex744 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l745
															}
															position++
															goto l744
														l745:
															position, tokenIndex = position744, tokenIndex744
															if buffer[position] != rune('I') {
																goto l651
															}
															position++
														}
													l744:
														{
															position746, tokenIndex746 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l747
															}
															position++
															goto l746
														l747:
															position, tokenIndex = position746, tokenIndex746
															if buffer[position] != rune('T') {
																goto l651
															}
															position++
														}
													l746:
														if !_rules[rulews]() {
															goto l651
														}
														if !_rules[ruleoctaldigit]() {
															goto l651
														}
														if !_rules[rulesep]() {
															goto l651
														}
														if !_rules[ruleLoc8]() {
															goto l651
														}
														{
															add(ruleAction36, position)
														}
														add(ruleBit, position741)
													}
													break
												}
											}

										}
									l653:
										add(ruleBitOp, position652)
									}
									goto l482
								l651:
									position, tokenIndex = position482, tokenIndex482
									{
										position750 := position
										{
											position751, tokenIndex751 := position, tokenIndex
											{
												position753 := position
												{
													position754 := position
													{
														position755, tokenIndex755 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l756
														}
														position++
														goto l755
													l756:
														position, tokenIndex = position755, tokenIndex755
														if buffer[position] != rune('R') {
															goto l752
														}
														position++
													}
												l755:
													{
														position757, tokenIndex757 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l758
														}
														position++
														goto l757
													l758:
														position, tokenIndex = position757, tokenIndex757
														if buffer[position] != rune('L') {
															goto l752
														}
														position++
													}
												l757:
													{
														position759, tokenIndex759 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l760
														}
														position++
														goto l759
													l760:
														position, tokenIndex = position759, tokenIndex759
														if buffer[position] != rune('C') {
															goto l752
														}
														position++
													}
												l759:
													{
														position761, tokenIndex761 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l762
														}
														position++
														goto l761
													l762:
														position, tokenIndex = position761, tokenIndex761
														if buffer[position] != rune('A') {
															goto l752
														}
														position++
													}
												l761:
													add(rulePegText, position754)
												}
												{
													add(ruleAction41, position)
												}
												add(ruleRlca, position753)
											}
											goto l751
										l752:
											position, tokenIndex = position751, tokenIndex751
											{
												position765 := position
												{
													position766 := position
													{
														position767, tokenIndex767 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l768
														}
														position++
														goto l767
													l768:
														position, tokenIndex = position767, tokenIndex767
														if buffer[position] != rune('R') {
															goto l764
														}
														position++
													}
												l767:
													{
														position769, tokenIndex769 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l770
														}
														position++
														goto l769
													l770:
														position, tokenIndex = position769, tokenIndex769
														if buffer[position] != rune('R') {
															goto l764
														}
														position++
													}
												l769:
													{
														position771, tokenIndex771 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l772
														}
														position++
														goto l771
													l772:
														position, tokenIndex = position771, tokenIndex771
														if buffer[position] != rune('C') {
															goto l764
														}
														position++
													}
												l771:
													{
														position773, tokenIndex773 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l774
														}
														position++
														goto l773
													l774:
														position, tokenIndex = position773, tokenIndex773
														if buffer[position] != rune('A') {
															goto l764
														}
														position++
													}
												l773:
													add(rulePegText, position766)
												}
												{
													add(ruleAction42, position)
												}
												add(ruleRrca, position765)
											}
											goto l751
										l764:
											position, tokenIndex = position751, tokenIndex751
											{
												position777 := position
												{
													position778 := position
													{
														position779, tokenIndex779 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l780
														}
														position++
														goto l779
													l780:
														position, tokenIndex = position779, tokenIndex779
														if buffer[position] != rune('R') {
															goto l776
														}
														position++
													}
												l779:
													{
														position781, tokenIndex781 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l782
														}
														position++
														goto l781
													l782:
														position, tokenIndex = position781, tokenIndex781
														if buffer[position] != rune('L') {
															goto l776
														}
														position++
													}
												l781:
													{
														position783, tokenIndex783 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l784
														}
														position++
														goto l783
													l784:
														position, tokenIndex = position783, tokenIndex783
														if buffer[position] != rune('A') {
															goto l776
														}
														position++
													}
												l783:
													add(rulePegText, position778)
												}
												{
													add(ruleAction43, position)
												}
												add(ruleRla, position777)
											}
											goto l751
										l776:
											position, tokenIndex = position751, tokenIndex751
											{
												position787 := position
												{
													position788 := position
													{
														position789, tokenIndex789 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l790
														}
														position++
														goto l789
													l790:
														position, tokenIndex = position789, tokenIndex789
														if buffer[position] != rune('D') {
															goto l786
														}
														position++
													}
												l789:
													{
														position791, tokenIndex791 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l792
														}
														position++
														goto l791
													l792:
														position, tokenIndex = position791, tokenIndex791
														if buffer[position] != rune('A') {
															goto l786
														}
														position++
													}
												l791:
													{
														position793, tokenIndex793 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l794
														}
														position++
														goto l793
													l794:
														position, tokenIndex = position793, tokenIndex793
														if buffer[position] != rune('A') {
															goto l786
														}
														position++
													}
												l793:
													add(rulePegText, position788)
												}
												{
													add(ruleAction45, position)
												}
												add(ruleDaa, position787)
											}
											goto l751
										l786:
											position, tokenIndex = position751, tokenIndex751
											{
												position797 := position
												{
													position798 := position
													{
														position799, tokenIndex799 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l800
														}
														position++
														goto l799
													l800:
														position, tokenIndex = position799, tokenIndex799
														if buffer[position] != rune('C') {
															goto l796
														}
														position++
													}
												l799:
													{
														position801, tokenIndex801 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l802
														}
														position++
														goto l801
													l802:
														position, tokenIndex = position801, tokenIndex801
														if buffer[position] != rune('P') {
															goto l796
														}
														position++
													}
												l801:
													{
														position803, tokenIndex803 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l804
														}
														position++
														goto l803
													l804:
														position, tokenIndex = position803, tokenIndex803
														if buffer[position] != rune('L') {
															goto l796
														}
														position++
													}
												l803:
													add(rulePegText, position798)
												}
												{
													add(ruleAction46, position)
												}
												add(ruleCpl, position797)
											}
											goto l751
										l796:
											position, tokenIndex = position751, tokenIndex751
											{
												position807 := position
												{
													position808 := position
													{
														position809, tokenIndex809 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l810
														}
														position++
														goto l809
													l810:
														position, tokenIndex = position809, tokenIndex809
														if buffer[position] != rune('E') {
															goto l806
														}
														position++
													}
												l809:
													{
														position811, tokenIndex811 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l812
														}
														position++
														goto l811
													l812:
														position, tokenIndex = position811, tokenIndex811
														if buffer[position] != rune('X') {
															goto l806
														}
														position++
													}
												l811:
													{
														position813, tokenIndex813 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l814
														}
														position++
														goto l813
													l814:
														position, tokenIndex = position813, tokenIndex813
														if buffer[position] != rune('X') {
															goto l806
														}
														position++
													}
												l813:
													add(rulePegText, position808)
												}
												{
													add(ruleAction49, position)
												}
												add(ruleExx, position807)
											}
											goto l751
										l806:
											position, tokenIndex = position751, tokenIndex751
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position817 := position
														{
															position818 := position
															{
																position819, tokenIndex819 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l820
																}
																position++
																goto l819
															l820:
																position, tokenIndex = position819, tokenIndex819
																if buffer[position] != rune('E') {
																	goto l749
																}
																position++
															}
														l819:
															{
																position821, tokenIndex821 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l822
																}
																position++
																goto l821
															l822:
																position, tokenIndex = position821, tokenIndex821
																if buffer[position] != rune('I') {
																	goto l749
																}
																position++
															}
														l821:
															add(rulePegText, position818)
														}
														{
															add(ruleAction51, position)
														}
														add(ruleEi, position817)
													}
													break
												case 'D', 'd':
													{
														position824 := position
														{
															position825 := position
															{
																position826, tokenIndex826 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l827
																}
																position++
																goto l826
															l827:
																position, tokenIndex = position826, tokenIndex826
																if buffer[position] != rune('D') {
																	goto l749
																}
																position++
															}
														l826:
															{
																position828, tokenIndex828 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l829
																}
																position++
																goto l828
															l829:
																position, tokenIndex = position828, tokenIndex828
																if buffer[position] != rune('I') {
																	goto l749
																}
																position++
															}
														l828:
															add(rulePegText, position825)
														}
														{
															add(ruleAction50, position)
														}
														add(ruleDi, position824)
													}
													break
												case 'C', 'c':
													{
														position831 := position
														{
															position832 := position
															{
																position833, tokenIndex833 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l834
																}
																position++
																goto l833
															l834:
																position, tokenIndex = position833, tokenIndex833
																if buffer[position] != rune('C') {
																	goto l749
																}
																position++
															}
														l833:
															{
																position835, tokenIndex835 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l836
																}
																position++
																goto l835
															l836:
																position, tokenIndex = position835, tokenIndex835
																if buffer[position] != rune('C') {
																	goto l749
																}
																position++
															}
														l835:
															{
																position837, tokenIndex837 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l838
																}
																position++
																goto l837
															l838:
																position, tokenIndex = position837, tokenIndex837
																if buffer[position] != rune('F') {
																	goto l749
																}
																position++
															}
														l837:
															add(rulePegText, position832)
														}
														{
															add(ruleAction48, position)
														}
														add(ruleCcf, position831)
													}
													break
												case 'S', 's':
													{
														position840 := position
														{
															position841 := position
															{
																position842, tokenIndex842 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l843
																}
																position++
																goto l842
															l843:
																position, tokenIndex = position842, tokenIndex842
																if buffer[position] != rune('S') {
																	goto l749
																}
																position++
															}
														l842:
															{
																position844, tokenIndex844 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l845
																}
																position++
																goto l844
															l845:
																position, tokenIndex = position844, tokenIndex844
																if buffer[position] != rune('C') {
																	goto l749
																}
																position++
															}
														l844:
															{
																position846, tokenIndex846 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l847
																}
																position++
																goto l846
															l847:
																position, tokenIndex = position846, tokenIndex846
																if buffer[position] != rune('F') {
																	goto l749
																}
																position++
															}
														l846:
															add(rulePegText, position841)
														}
														{
															add(ruleAction47, position)
														}
														add(ruleScf, position840)
													}
													break
												case 'R', 'r':
													{
														position849 := position
														{
															position850 := position
															{
																position851, tokenIndex851 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l852
																}
																position++
																goto l851
															l852:
																position, tokenIndex = position851, tokenIndex851
																if buffer[position] != rune('R') {
																	goto l749
																}
																position++
															}
														l851:
															{
																position853, tokenIndex853 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l854
																}
																position++
																goto l853
															l854:
																position, tokenIndex = position853, tokenIndex853
																if buffer[position] != rune('R') {
																	goto l749
																}
																position++
															}
														l853:
															{
																position855, tokenIndex855 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l856
																}
																position++
																goto l855
															l856:
																position, tokenIndex = position855, tokenIndex855
																if buffer[position] != rune('A') {
																	goto l749
																}
																position++
															}
														l855:
															add(rulePegText, position850)
														}
														{
															add(ruleAction44, position)
														}
														add(ruleRra, position849)
													}
													break
												case 'H', 'h':
													{
														position858 := position
														{
															position859 := position
															{
																position860, tokenIndex860 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l861
																}
																position++
																goto l860
															l861:
																position, tokenIndex = position860, tokenIndex860
																if buffer[position] != rune('H') {
																	goto l749
																}
																position++
															}
														l860:
															{
																position862, tokenIndex862 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l863
																}
																position++
																goto l862
															l863:
																position, tokenIndex = position862, tokenIndex862
																if buffer[position] != rune('A') {
																	goto l749
																}
																position++
															}
														l862:
															{
																position864, tokenIndex864 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l865
																}
																position++
																goto l864
															l865:
																position, tokenIndex = position864, tokenIndex864
																if buffer[position] != rune('L') {
																	goto l749
																}
																position++
															}
														l864:
															{
																position866, tokenIndex866 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l867
																}
																position++
																goto l866
															l867:
																position, tokenIndex = position866, tokenIndex866
																if buffer[position] != rune('T') {
																	goto l749
																}
																position++
															}
														l866:
															add(rulePegText, position859)
														}
														{
															add(ruleAction40, position)
														}
														add(ruleHalt, position858)
													}
													break
												default:
													{
														position869 := position
														{
															position870 := position
															{
																position871, tokenIndex871 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l872
																}
																position++
																goto l871
															l872:
																position, tokenIndex = position871, tokenIndex871
																if buffer[position] != rune('N') {
																	goto l749
																}
																position++
															}
														l871:
															{
																position873, tokenIndex873 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l874
																}
																position++
																goto l873
															l874:
																position, tokenIndex = position873, tokenIndex873
																if buffer[position] != rune('O') {
																	goto l749
																}
																position++
															}
														l873:
															{
																position875, tokenIndex875 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l876
																}
																position++
																goto l875
															l876:
																position, tokenIndex = position875, tokenIndex875
																if buffer[position] != rune('P') {
																	goto l749
																}
																position++
															}
														l875:
															add(rulePegText, position870)
														}
														{
															add(ruleAction39, position)
														}
														add(ruleNop, position869)
													}
													break
												}
											}

										}
									l751:
										add(ruleSimple, position750)
									}
									goto l482
								l749:
									position, tokenIndex = position482, tokenIndex482
									{
										position878 := position
										{
											position879, tokenIndex879 := position, tokenIndex
											{
												position881 := position
												{
													position882, tokenIndex882 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l883
													}
													position++
													goto l882
												l883:
													position, tokenIndex = position882, tokenIndex882
													if buffer[position] != rune('R') {
														goto l880
													}
													position++
												}
											l882:
												{
													position884, tokenIndex884 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l885
													}
													position++
													goto l884
												l885:
													position, tokenIndex = position884, tokenIndex884
													if buffer[position] != rune('S') {
														goto l880
													}
													position++
												}
											l884:
												{
													position886, tokenIndex886 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l887
													}
													position++
													goto l886
												l887:
													position, tokenIndex = position886, tokenIndex886
													if buffer[position] != rune('T') {
														goto l880
													}
													position++
												}
											l886:
												if !_rules[rulews]() {
													goto l880
												}
												if !_rules[rulen]() {
													goto l880
												}
												{
													add(ruleAction52, position)
												}
												add(ruleRst, position881)
											}
											goto l879
										l880:
											position, tokenIndex = position879, tokenIndex879
											{
												position890 := position
												{
													position891, tokenIndex891 := position, tokenIndex
													if buffer[position] != rune('j') {
														goto l892
													}
													position++
													goto l891
												l892:
													position, tokenIndex = position891, tokenIndex891
													if buffer[position] != rune('J') {
														goto l889
													}
													position++
												}
											l891:
												{
													position893, tokenIndex893 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l894
													}
													position++
													goto l893
												l894:
													position, tokenIndex = position893, tokenIndex893
													if buffer[position] != rune('P') {
														goto l889
													}
													position++
												}
											l893:
												if !_rules[rulews]() {
													goto l889
												}
												{
													position895, tokenIndex895 := position, tokenIndex
													if !_rules[rulecc]() {
														goto l895
													}
													if !_rules[rulesep]() {
														goto l895
													}
													goto l896
												l895:
													position, tokenIndex = position895, tokenIndex895
												}
											l896:
												if !_rules[ruleSrc16]() {
													goto l889
												}
												{
													add(ruleAction55, position)
												}
												add(ruleJp, position890)
											}
											goto l879
										l889:
											position, tokenIndex = position879, tokenIndex879
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position899 := position
														{
															position900, tokenIndex900 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l901
															}
															position++
															goto l900
														l901:
															position, tokenIndex = position900, tokenIndex900
															if buffer[position] != rune('D') {
																goto l3
															}
															position++
														}
													l900:
														{
															position902, tokenIndex902 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l903
															}
															position++
															goto l902
														l903:
															position, tokenIndex = position902, tokenIndex902
															if buffer[position] != rune('J') {
																goto l3
															}
															position++
														}
													l902:
														{
															position904, tokenIndex904 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l905
															}
															position++
															goto l904
														l905:
															position, tokenIndex = position904, tokenIndex904
															if buffer[position] != rune('N') {
																goto l3
															}
															position++
														}
													l904:
														{
															position906, tokenIndex906 := position, tokenIndex
															if buffer[position] != rune('z') {
																goto l907
															}
															position++
															goto l906
														l907:
															position, tokenIndex = position906, tokenIndex906
															if buffer[position] != rune('Z') {
																goto l3
															}
															position++
														}
													l906:
														if !_rules[rulews]() {
															goto l3
														}
														if !_rules[ruledisp]() {
															goto l3
														}
														{
															add(ruleAction57, position)
														}
														add(ruleDjnz, position899)
													}
													break
												case 'J', 'j':
													{
														position909 := position
														{
															position910, tokenIndex910 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l911
															}
															position++
															goto l910
														l911:
															position, tokenIndex = position910, tokenIndex910
															if buffer[position] != rune('J') {
																goto l3
															}
															position++
														}
													l910:
														{
															position912, tokenIndex912 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l913
															}
															position++
															goto l912
														l913:
															position, tokenIndex = position912, tokenIndex912
															if buffer[position] != rune('R') {
																goto l3
															}
															position++
														}
													l912:
														if !_rules[rulews]() {
															goto l3
														}
														{
															position914, tokenIndex914 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l914
															}
															if !_rules[rulesep]() {
																goto l914
															}
															goto l915
														l914:
															position, tokenIndex = position914, tokenIndex914
														}
													l915:
														if !_rules[ruledisp]() {
															goto l3
														}
														{
															add(ruleAction56, position)
														}
														add(ruleJr, position909)
													}
													break
												case 'R', 'r':
													{
														position917 := position
														{
															position918, tokenIndex918 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l919
															}
															position++
															goto l918
														l919:
															position, tokenIndex = position918, tokenIndex918
															if buffer[position] != rune('R') {
																goto l3
															}
															position++
														}
													l918:
														{
															position920, tokenIndex920 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l921
															}
															position++
															goto l920
														l921:
															position, tokenIndex = position920, tokenIndex920
															if buffer[position] != rune('E') {
																goto l3
															}
															position++
														}
													l920:
														{
															position922, tokenIndex922 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l923
															}
															position++
															goto l922
														l923:
															position, tokenIndex = position922, tokenIndex922
															if buffer[position] != rune('T') {
																goto l3
															}
															position++
														}
													l922:
														{
															position924, tokenIndex924 := position, tokenIndex
															if !_rules[rulews]() {
																goto l924
															}
															if !_rules[rulecc]() {
																goto l924
															}
															goto l925
														l924:
															position, tokenIndex = position924, tokenIndex924
														}
													l925:
														{
															add(ruleAction54, position)
														}
														add(ruleRet, position917)
													}
													break
												default:
													{
														position927 := position
														{
															position928, tokenIndex928 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l929
															}
															position++
															goto l928
														l929:
															position, tokenIndex = position928, tokenIndex928
															if buffer[position] != rune('C') {
																goto l3
															}
															position++
														}
													l928:
														{
															position930, tokenIndex930 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l931
															}
															position++
															goto l930
														l931:
															position, tokenIndex = position930, tokenIndex930
															if buffer[position] != rune('A') {
																goto l3
															}
															position++
														}
													l930:
														{
															position932, tokenIndex932 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l933
															}
															position++
															goto l932
														l933:
															position, tokenIndex = position932, tokenIndex932
															if buffer[position] != rune('L') {
																goto l3
															}
															position++
														}
													l932:
														{
															position934, tokenIndex934 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l935
															}
															position++
															goto l934
														l935:
															position, tokenIndex = position934, tokenIndex934
															if buffer[position] != rune('L') {
																goto l3
															}
															position++
														}
													l934:
														if !_rules[rulews]() {
															goto l3
														}
														{
															position936, tokenIndex936 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l936
															}
															if !_rules[rulesep]() {
																goto l936
															}
															goto l937
														l936:
															position, tokenIndex = position936, tokenIndex936
														}
													l937:
														if !_rules[ruleSrc16]() {
															goto l3
														}
														{
															add(ruleAction53, position)
														}
														add(ruleCall, position927)
													}
													break
												}
											}

										}
									l879:
										add(ruleJump, position878)
									}
								}
							l482:
								add(ruleInstruction, position479)
							}
							{
								position939, tokenIndex939 := position, tokenIndex
								if buffer[position] != rune('\n') {
									goto l939
								}
								position++
								goto l940
							l939:
								position, tokenIndex = position939, tokenIndex939
							}
						l940:
							{
								add(ruleAction0, position)
							}
							add(ruleLine, position478)
						}
					}
				l473:
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				add(ruleProgram, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 BlankLine <- <(ws* '\n')> */
		nil,
		/* 2 Line <- <(Instruction '\n'? Action0)> */
		nil,
		/* 3 Instruction <- <(ws* (Assignment / Inc / Dec / Add16 / Alu / BitOp / Simple / Jump))> */
		nil,
		/* 4 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 5 Load <- <(Load8 / Load16)> */
		nil,
		/* 6 Load8 <- <(('l' / 'L') ('d' / 'D') ws Dst8 sep Src8 Action1)> */
		nil,
		/* 7 Load16 <- <(('l' / 'L') ('d' / 'D') ws Dst16 sep Src16 Action2)> */
		nil,
		/* 8 Push <- <(('p' / 'P') ('u' / 'U') ('s' / 'S') ('h' / 'H') ws Src16 Action3)> */
		nil,
		/* 9 Pop <- <(('p' / 'P') ('o' / 'O') ('p' / 'P') ws Dst16 Action4)> */
		nil,
		/* 10 Ex <- <(('e' / 'E') ('x' / 'X') ws Dst16 sep Src16 Action5)> */
		nil,
		/* 11 Inc <- <(Inc16 / Inc8)> */
		nil,
		/* 12 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action6)> */
		nil,
		/* 13 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action7)> */
		nil,
		/* 14 Dec <- <(Dec16 / Dec8)> */
		nil,
		/* 15 Dec8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action8)> */
		nil,
		/* 16 Dec16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action9)> */
		nil,
		/* 17 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action10)> */
		nil,
		/* 18 Dst8 <- <(Reg8 Action11)> */
		nil,
		/* 19 Src8 <- <((n / Reg8) Action12)> */
		func() bool {
			position960, tokenIndex960 := position, tokenIndex
			{
				position961 := position
				{
					position962, tokenIndex962 := position, tokenIndex
					if !_rules[rulen]() {
						goto l963
					}
					goto l962
				l963:
					position, tokenIndex = position962, tokenIndex962
					if !_rules[ruleReg8]() {
						goto l960
					}
				}
			l962:
				{
					add(ruleAction12, position)
				}
				add(ruleSrc8, position961)
			}
			return true
		l960:
			position, tokenIndex = position960, tokenIndex960
			return false
		},
		/* 20 Loc8 <- <(Reg8 Action13)> */
		func() bool {
			position965, tokenIndex965 := position, tokenIndex
			{
				position966 := position
				if !_rules[ruleReg8]() {
					goto l965
				}
				{
					add(ruleAction13, position)
				}
				add(ruleLoc8, position966)
			}
			return true
		l965:
			position, tokenIndex = position965, tokenIndex965
			return false
		},
		/* 21 Reg8 <- <(<(IXH / IXL / IYH / ((&('I' | 'i') IYL) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action14)> */
		func() bool {
			position968, tokenIndex968 := position, tokenIndex
			{
				position969 := position
				{
					position970 := position
					{
						position971, tokenIndex971 := position, tokenIndex
						{
							position973 := position
							{
								position974, tokenIndex974 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l975
								}
								position++
								goto l974
							l975:
								position, tokenIndex = position974, tokenIndex974
								if buffer[position] != rune('I') {
									goto l972
								}
								position++
							}
						l974:
							{
								position976, tokenIndex976 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l977
								}
								position++
								goto l976
							l977:
								position, tokenIndex = position976, tokenIndex976
								if buffer[position] != rune('X') {
									goto l972
								}
								position++
							}
						l976:
							{
								position978, tokenIndex978 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l979
								}
								position++
								goto l978
							l979:
								position, tokenIndex = position978, tokenIndex978
								if buffer[position] != rune('H') {
									goto l972
								}
								position++
							}
						l978:
							add(ruleIXH, position973)
						}
						goto l971
					l972:
						position, tokenIndex = position971, tokenIndex971
						{
							position981 := position
							{
								position982, tokenIndex982 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l983
								}
								position++
								goto l982
							l983:
								position, tokenIndex = position982, tokenIndex982
								if buffer[position] != rune('I') {
									goto l980
								}
								position++
							}
						l982:
							{
								position984, tokenIndex984 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l985
								}
								position++
								goto l984
							l985:
								position, tokenIndex = position984, tokenIndex984
								if buffer[position] != rune('X') {
									goto l980
								}
								position++
							}
						l984:
							{
								position986, tokenIndex986 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l987
								}
								position++
								goto l986
							l987:
								position, tokenIndex = position986, tokenIndex986
								if buffer[position] != rune('L') {
									goto l980
								}
								position++
							}
						l986:
							add(ruleIXL, position981)
						}
						goto l971
					l980:
						position, tokenIndex = position971, tokenIndex971
						{
							position989 := position
							{
								position990, tokenIndex990 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l991
								}
								position++
								goto l990
							l991:
								position, tokenIndex = position990, tokenIndex990
								if buffer[position] != rune('I') {
									goto l988
								}
								position++
							}
						l990:
							{
								position992, tokenIndex992 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l993
								}
								position++
								goto l992
							l993:
								position, tokenIndex = position992, tokenIndex992
								if buffer[position] != rune('Y') {
									goto l988
								}
								position++
							}
						l992:
							{
								position994, tokenIndex994 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l995
								}
								position++
								goto l994
							l995:
								position, tokenIndex = position994, tokenIndex994
								if buffer[position] != rune('H') {
									goto l988
								}
								position++
							}
						l994:
							add(ruleIYH, position989)
						}
						goto l971
					l988:
						position, tokenIndex = position971, tokenIndex971
						{
							switch buffer[position] {
							case 'I', 'i':
								{
									position997 := position
									{
										position998, tokenIndex998 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l999
										}
										position++
										goto l998
									l999:
										position, tokenIndex = position998, tokenIndex998
										if buffer[position] != rune('I') {
											goto l968
										}
										position++
									}
								l998:
									{
										position1000, tokenIndex1000 := position, tokenIndex
										if buffer[position] != rune('y') {
											goto l1001
										}
										position++
										goto l1000
									l1001:
										position, tokenIndex = position1000, tokenIndex1000
										if buffer[position] != rune('Y') {
											goto l968
										}
										position++
									}
								l1000:
									{
										position1002, tokenIndex1002 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1003
										}
										position++
										goto l1002
									l1003:
										position, tokenIndex = position1002, tokenIndex1002
										if buffer[position] != rune('L') {
											goto l968
										}
										position++
									}
								l1002:
									add(ruleIYL, position997)
								}
								break
							case 'L', 'l':
								{
									position1004 := position
									{
										position1005, tokenIndex1005 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1006
										}
										position++
										goto l1005
									l1006:
										position, tokenIndex = position1005, tokenIndex1005
										if buffer[position] != rune('L') {
											goto l968
										}
										position++
									}
								l1005:
									add(ruleL, position1004)
								}
								break
							case 'H', 'h':
								{
									position1007 := position
									{
										position1008, tokenIndex1008 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1009
										}
										position++
										goto l1008
									l1009:
										position, tokenIndex = position1008, tokenIndex1008
										if buffer[position] != rune('H') {
											goto l968
										}
										position++
									}
								l1008:
									add(ruleH, position1007)
								}
								break
							case 'E', 'e':
								{
									position1010 := position
									{
										position1011, tokenIndex1011 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1012
										}
										position++
										goto l1011
									l1012:
										position, tokenIndex = position1011, tokenIndex1011
										if buffer[position] != rune('E') {
											goto l968
										}
										position++
									}
								l1011:
									add(ruleE, position1010)
								}
								break
							case 'D', 'd':
								{
									position1013 := position
									{
										position1014, tokenIndex1014 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1015
										}
										position++
										goto l1014
									l1015:
										position, tokenIndex = position1014, tokenIndex1014
										if buffer[position] != rune('D') {
											goto l968
										}
										position++
									}
								l1014:
									add(ruleD, position1013)
								}
								break
							case 'C', 'c':
								{
									position1016 := position
									{
										position1017, tokenIndex1017 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1018
										}
										position++
										goto l1017
									l1018:
										position, tokenIndex = position1017, tokenIndex1017
										if buffer[position] != rune('C') {
											goto l968
										}
										position++
									}
								l1017:
									add(ruleC, position1016)
								}
								break
							case 'B', 'b':
								{
									position1019 := position
									{
										position1020, tokenIndex1020 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1021
										}
										position++
										goto l1020
									l1021:
										position, tokenIndex = position1020, tokenIndex1020
										if buffer[position] != rune('B') {
											goto l968
										}
										position++
									}
								l1020:
									add(ruleB, position1019)
								}
								break
							case 'F', 'f':
								{
									position1022 := position
									{
										position1023, tokenIndex1023 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1024
										}
										position++
										goto l1023
									l1024:
										position, tokenIndex = position1023, tokenIndex1023
										if buffer[position] != rune('F') {
											goto l968
										}
										position++
									}
								l1023:
									add(ruleF, position1022)
								}
								break
							default:
								{
									position1025 := position
									{
										position1026, tokenIndex1026 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1027
										}
										position++
										goto l1026
									l1027:
										position, tokenIndex = position1026, tokenIndex1026
										if buffer[position] != rune('A') {
											goto l968
										}
										position++
									}
								l1026:
									add(ruleA, position1025)
								}
								break
							}
						}

					}
				l971:
					add(rulePegText, position970)
				}
				{
					add(ruleAction14, position)
				}
				add(ruleReg8, position969)
			}
			return true
		l968:
			position, tokenIndex = position968, tokenIndex968
			return false
		},
		/* 22 Dst16 <- <((Reg16 / nn_contents) Action15)> */
		func() bool {
			position1029, tokenIndex1029 := position, tokenIndex
			{
				position1030 := position
				{
					position1031, tokenIndex1031 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1032
					}
					goto l1031
				l1032:
					position, tokenIndex = position1031, tokenIndex1031
					if !_rules[rulenn_contents]() {
						goto l1029
					}
				}
			l1031:
				{
					add(ruleAction15, position)
				}
				add(ruleDst16, position1030)
			}
			return true
		l1029:
			position, tokenIndex = position1029, tokenIndex1029
			return false
		},
		/* 23 Src16 <- <((Reg16 / nn / nn_contents) Action16)> */
		func() bool {
			position1034, tokenIndex1034 := position, tokenIndex
			{
				position1035 := position
				{
					position1036, tokenIndex1036 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1037
					}
					goto l1036
				l1037:
					position, tokenIndex = position1036, tokenIndex1036
					if !_rules[rulenn]() {
						goto l1038
					}
					goto l1036
				l1038:
					position, tokenIndex = position1036, tokenIndex1036
					if !_rules[rulenn_contents]() {
						goto l1034
					}
				}
			l1036:
				{
					add(ruleAction16, position)
				}
				add(ruleSrc16, position1035)
			}
			return true
		l1034:
			position, tokenIndex = position1034, tokenIndex1034
			return false
		},
		/* 24 Loc16 <- <(Reg16 Action17)> */
		func() bool {
			position1040, tokenIndex1040 := position, tokenIndex
			{
				position1041 := position
				if !_rules[ruleReg16]() {
					goto l1040
				}
				{
					add(ruleAction17, position)
				}
				add(ruleLoc16, position1041)
			}
			return true
		l1040:
			position, tokenIndex = position1040, tokenIndex1040
			return false
		},
		/* 25 Reg16 <- <(<(AF_PRIME / IX / ((&('S' | 's') SP) | (&('I' | 'i') IY) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action18)> */
		func() bool {
			position1043, tokenIndex1043 := position, tokenIndex
			{
				position1044 := position
				{
					position1045 := position
					{
						position1046, tokenIndex1046 := position, tokenIndex
						{
							position1048 := position
							{
								position1049, tokenIndex1049 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1050
								}
								position++
								goto l1049
							l1050:
								position, tokenIndex = position1049, tokenIndex1049
								if buffer[position] != rune('A') {
									goto l1047
								}
								position++
							}
						l1049:
							{
								position1051, tokenIndex1051 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1052
								}
								position++
								goto l1051
							l1052:
								position, tokenIndex = position1051, tokenIndex1051
								if buffer[position] != rune('F') {
									goto l1047
								}
								position++
							}
						l1051:
							if buffer[position] != rune('\'') {
								goto l1047
							}
							position++
							add(ruleAF_PRIME, position1048)
						}
						goto l1046
					l1047:
						position, tokenIndex = position1046, tokenIndex1046
						{
							position1054 := position
							{
								position1055, tokenIndex1055 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1056
								}
								position++
								goto l1055
							l1056:
								position, tokenIndex = position1055, tokenIndex1055
								if buffer[position] != rune('I') {
									goto l1053
								}
								position++
							}
						l1055:
							{
								position1057, tokenIndex1057 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1058
								}
								position++
								goto l1057
							l1058:
								position, tokenIndex = position1057, tokenIndex1057
								if buffer[position] != rune('X') {
									goto l1053
								}
								position++
							}
						l1057:
							add(ruleIX, position1054)
						}
						goto l1046
					l1053:
						position, tokenIndex = position1046, tokenIndex1046
						{
							switch buffer[position] {
							case 'S', 's':
								{
									position1060 := position
									{
										position1061, tokenIndex1061 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1062
										}
										position++
										goto l1061
									l1062:
										position, tokenIndex = position1061, tokenIndex1061
										if buffer[position] != rune('S') {
											goto l1043
										}
										position++
									}
								l1061:
									{
										position1063, tokenIndex1063 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1064
										}
										position++
										goto l1063
									l1064:
										position, tokenIndex = position1063, tokenIndex1063
										if buffer[position] != rune('P') {
											goto l1043
										}
										position++
									}
								l1063:
									add(ruleSP, position1060)
								}
								break
							case 'I', 'i':
								{
									position1065 := position
									{
										position1066, tokenIndex1066 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1067
										}
										position++
										goto l1066
									l1067:
										position, tokenIndex = position1066, tokenIndex1066
										if buffer[position] != rune('I') {
											goto l1043
										}
										position++
									}
								l1066:
									{
										position1068, tokenIndex1068 := position, tokenIndex
										if buffer[position] != rune('y') {
											goto l1069
										}
										position++
										goto l1068
									l1069:
										position, tokenIndex = position1068, tokenIndex1068
										if buffer[position] != rune('Y') {
											goto l1043
										}
										position++
									}
								l1068:
									add(ruleIY, position1065)
								}
								break
							case 'H', 'h':
								{
									position1070 := position
									{
										position1071, tokenIndex1071 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1072
										}
										position++
										goto l1071
									l1072:
										position, tokenIndex = position1071, tokenIndex1071
										if buffer[position] != rune('H') {
											goto l1043
										}
										position++
									}
								l1071:
									{
										position1073, tokenIndex1073 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1074
										}
										position++
										goto l1073
									l1074:
										position, tokenIndex = position1073, tokenIndex1073
										if buffer[position] != rune('L') {
											goto l1043
										}
										position++
									}
								l1073:
									add(ruleHL, position1070)
								}
								break
							case 'D', 'd':
								{
									position1075 := position
									{
										position1076, tokenIndex1076 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1077
										}
										position++
										goto l1076
									l1077:
										position, tokenIndex = position1076, tokenIndex1076
										if buffer[position] != rune('D') {
											goto l1043
										}
										position++
									}
								l1076:
									{
										position1078, tokenIndex1078 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1079
										}
										position++
										goto l1078
									l1079:
										position, tokenIndex = position1078, tokenIndex1078
										if buffer[position] != rune('E') {
											goto l1043
										}
										position++
									}
								l1078:
									add(ruleDE, position1075)
								}
								break
							case 'B', 'b':
								{
									position1080 := position
									{
										position1081, tokenIndex1081 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1082
										}
										position++
										goto l1081
									l1082:
										position, tokenIndex = position1081, tokenIndex1081
										if buffer[position] != rune('B') {
											goto l1043
										}
										position++
									}
								l1081:
									{
										position1083, tokenIndex1083 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1084
										}
										position++
										goto l1083
									l1084:
										position, tokenIndex = position1083, tokenIndex1083
										if buffer[position] != rune('C') {
											goto l1043
										}
										position++
									}
								l1083:
									add(ruleBC, position1080)
								}
								break
							default:
								{
									position1085 := position
									{
										position1086, tokenIndex1086 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1087
										}
										position++
										goto l1086
									l1087:
										position, tokenIndex = position1086, tokenIndex1086
										if buffer[position] != rune('A') {
											goto l1043
										}
										position++
									}
								l1086:
									{
										position1088, tokenIndex1088 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1089
										}
										position++
										goto l1088
									l1089:
										position, tokenIndex = position1088, tokenIndex1088
										if buffer[position] != rune('F') {
											goto l1043
										}
										position++
									}
								l1088:
									add(ruleAF, position1085)
								}
								break
							}
						}

					}
				l1046:
					add(rulePegText, position1045)
				}
				{
					add(ruleAction18, position)
				}
				add(ruleReg16, position1044)
			}
			return true
		l1043:
			position, tokenIndex = position1043, tokenIndex1043
			return false
		},
		/* 26 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1091, tokenIndex1091 := position, tokenIndex
			{
				position1092 := position
				{
					position1093, tokenIndex1093 := position, tokenIndex
					{
						position1095 := position
						{
							position1096 := position
							if !_rules[rulehexdigit]() {
								goto l1094
							}
							if !_rules[rulehexdigit]() {
								goto l1094
							}
							add(rulePegText, position1096)
						}
						{
							position1097, tokenIndex1097 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1098
							}
							position++
							goto l1097
						l1098:
							position, tokenIndex = position1097, tokenIndex1097
							if buffer[position] != rune('H') {
								goto l1094
							}
							position++
						}
					l1097:
						{
							add(ruleAction58, position)
						}
						add(rulehexByteH, position1095)
					}
					goto l1093
				l1094:
					position, tokenIndex = position1093, tokenIndex1093
					{
						position1101 := position
						if buffer[position] != rune('0') {
							goto l1100
						}
						position++
						{
							position1102, tokenIndex1102 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1103
							}
							position++
							goto l1102
						l1103:
							position, tokenIndex = position1102, tokenIndex1102
							if buffer[position] != rune('X') {
								goto l1100
							}
							position++
						}
					l1102:
						{
							position1104 := position
							if !_rules[rulehexdigit]() {
								goto l1100
							}
							if !_rules[rulehexdigit]() {
								goto l1100
							}
							add(rulePegText, position1104)
						}
						{
							add(ruleAction59, position)
						}
						add(rulehexByte0x, position1101)
					}
					goto l1093
				l1100:
					position, tokenIndex = position1093, tokenIndex1093
					{
						position1106 := position
						{
							position1107 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1091
							}
							position++
						l1108:
							{
								position1109, tokenIndex1109 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1109
								}
								position++
								goto l1108
							l1109:
								position, tokenIndex = position1109, tokenIndex1109
							}
							add(rulePegText, position1107)
						}
						{
							add(ruleAction60, position)
						}
						add(ruledecimalByte, position1106)
					}
				}
			l1093:
				add(rulen, position1092)
			}
			return true
		l1091:
			position, tokenIndex = position1091, tokenIndex1091
			return false
		},
		/* 27 nn <- <(hexWordH / hexWord0x)> */
		func() bool {
			position1111, tokenIndex1111 := position, tokenIndex
			{
				position1112 := position
				{
					position1113, tokenIndex1113 := position, tokenIndex
					{
						position1115 := position
						{
							position1116 := position
							if !_rules[rulehexdigit]() {
								goto l1114
							}
							if !_rules[rulehexdigit]() {
								goto l1114
							}
							if !_rules[rulehexdigit]() {
								goto l1114
							}
							if !_rules[rulehexdigit]() {
								goto l1114
							}
							add(rulePegText, position1116)
						}
						{
							position1117, tokenIndex1117 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1118
							}
							position++
							goto l1117
						l1118:
							position, tokenIndex = position1117, tokenIndex1117
							if buffer[position] != rune('H') {
								goto l1114
							}
							position++
						}
					l1117:
						{
							add(ruleAction61, position)
						}
						add(rulehexWordH, position1115)
					}
					goto l1113
				l1114:
					position, tokenIndex = position1113, tokenIndex1113
					{
						position1120 := position
						if buffer[position] != rune('0') {
							goto l1111
						}
						position++
						{
							position1121, tokenIndex1121 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1122
							}
							position++
							goto l1121
						l1122:
							position, tokenIndex = position1121, tokenIndex1121
							if buffer[position] != rune('X') {
								goto l1111
							}
							position++
						}
					l1121:
						{
							position1123 := position
							if !_rules[rulehexdigit]() {
								goto l1111
							}
							if !_rules[rulehexdigit]() {
								goto l1111
							}
							if !_rules[rulehexdigit]() {
								goto l1111
							}
							if !_rules[rulehexdigit]() {
								goto l1111
							}
							add(rulePegText, position1123)
						}
						{
							add(ruleAction62, position)
						}
						add(rulehexWord0x, position1120)
					}
				}
			l1113:
				add(rulenn, position1112)
			}
			return true
		l1111:
			position, tokenIndex = position1111, tokenIndex1111
			return false
		},
		/* 28 nn_contents <- <('(' nn ')' Action19)> */
		func() bool {
			position1125, tokenIndex1125 := position, tokenIndex
			{
				position1126 := position
				if buffer[position] != rune('(') {
					goto l1125
				}
				position++
				if !_rules[rulenn]() {
					goto l1125
				}
				if buffer[position] != rune(')') {
					goto l1125
				}
				position++
				{
					add(ruleAction19, position)
				}
				add(rulenn_contents, position1126)
			}
			return true
		l1125:
			position, tokenIndex = position1125, tokenIndex1125
			return false
		},
		/* 29 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 30 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action20)> */
		nil,
		/* 31 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action21)> */
		nil,
		/* 32 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action22)> */
		nil,
		/* 33 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action23)> */
		nil,
		/* 34 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action24)> */
		nil,
		/* 35 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action25)> */
		nil,
		/* 36 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action26)> */
		nil,
		/* 37 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action27)> */
		nil,
		/* 38 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 39 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 40 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 Action28)> */
		nil,
		/* 41 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 Action29)> */
		nil,
		/* 42 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 Action30)> */
		nil,
		/* 43 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 Action31)> */
		nil,
		/* 44 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 Action32)> */
		nil,
		/* 45 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 Action33)> */
		nil,
		/* 46 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 Action34)> */
		nil,
		/* 47 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 Action35)> */
		nil,
		/* 48 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action36)> */
		nil,
		/* 49 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 Action37)> */
		nil,
		/* 50 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 Action38)> */
		nil,
		/* 51 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 52 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action39)> */
		nil,
		/* 53 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action40)> */
		nil,
		/* 54 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action41)> */
		nil,
		/* 55 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action42)> */
		nil,
		/* 56 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action43)> */
		nil,
		/* 57 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action44)> */
		nil,
		/* 58 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action45)> */
		nil,
		/* 59 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action46)> */
		nil,
		/* 60 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action47)> */
		nil,
		/* 61 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action48)> */
		nil,
		/* 62 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action49)> */
		nil,
		/* 63 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action50)> */
		nil,
		/* 64 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action51)> */
		nil,
		/* 65 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 66 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action52)> */
		nil,
		/* 67 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action53)> */
		nil,
		/* 68 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action54)> */
		nil,
		/* 69 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action55)> */
		nil,
		/* 70 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action56)> */
		nil,
		/* 71 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action57)> */
		nil,
		/* 72 disp <- <signedDecimalByte> */
		func() bool {
			position1171, tokenIndex1171 := position, tokenIndex
			{
				position1172 := position
				{
					position1173 := position
					{
						position1174 := position
						{
							position1175, tokenIndex1175 := position, tokenIndex
							if buffer[position] != rune('-') {
								goto l1175
							}
							position++
							goto l1176
						l1175:
							position, tokenIndex = position1175, tokenIndex1175
						}
					l1176:
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1171
						}
						position++
					l1177:
						{
							position1178, tokenIndex1178 := position, tokenIndex
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1178
							}
							position++
							goto l1177
						l1178:
							position, tokenIndex = position1178, tokenIndex1178
						}
						add(rulePegText, position1174)
					}
					{
						add(ruleAction64, position)
					}
					add(rulesignedDecimalByte, position1173)
				}
				add(ruledisp, position1172)
			}
			return true
		l1171:
			position, tokenIndex = position1171, tokenIndex1171
			return false
		},
		/* 73 sep <- <(ws? ',' ws?)> */
		func() bool {
			position1180, tokenIndex1180 := position, tokenIndex
			{
				position1181 := position
				{
					position1182, tokenIndex1182 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1182
					}
					goto l1183
				l1182:
					position, tokenIndex = position1182, tokenIndex1182
				}
			l1183:
				if buffer[position] != rune(',') {
					goto l1180
				}
				position++
				{
					position1184, tokenIndex1184 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1184
					}
					goto l1185
				l1184:
					position, tokenIndex = position1184, tokenIndex1184
				}
			l1185:
				add(rulesep, position1181)
			}
			return true
		l1180:
			position, tokenIndex = position1180, tokenIndex1180
			return false
		},
		/* 74 ws <- <' '+> */
		func() bool {
			position1186, tokenIndex1186 := position, tokenIndex
			{
				position1187 := position
				if buffer[position] != rune(' ') {
					goto l1186
				}
				position++
			l1188:
				{
					position1189, tokenIndex1189 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l1189
					}
					position++
					goto l1188
				l1189:
					position, tokenIndex = position1189, tokenIndex1189
				}
				add(rulews, position1187)
			}
			return true
		l1186:
			position, tokenIndex = position1186, tokenIndex1186
			return false
		},
		/* 75 A <- <('a' / 'A')> */
		nil,
		/* 76 F <- <('f' / 'F')> */
		nil,
		/* 77 B <- <('b' / 'B')> */
		nil,
		/* 78 C <- <('c' / 'C')> */
		nil,
		/* 79 D <- <('d' / 'D')> */
		nil,
		/* 80 E <- <('e' / 'E')> */
		nil,
		/* 81 H <- <('h' / 'H')> */
		nil,
		/* 82 L <- <('l' / 'L')> */
		nil,
		/* 83 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 84 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 85 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 86 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 87 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 88 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 89 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 90 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 91 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 92 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 93 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 94 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 95 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action58)> */
		nil,
		/* 96 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action59)> */
		nil,
		/* 97 decimalByte <- <(<[0-9]+> Action60)> */
		nil,
		/* 98 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action61)> */
		nil,
		/* 99 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action62)> */
		nil,
		/* 100 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position1215, tokenIndex1215 := position, tokenIndex
			{
				position1216 := position
				{
					position1217, tokenIndex1217 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1218
					}
					position++
					goto l1217
				l1218:
					position, tokenIndex = position1217, tokenIndex1217
					{
						position1219, tokenIndex1219 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l1220
						}
						position++
						goto l1219
					l1220:
						position, tokenIndex = position1219, tokenIndex1219
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l1215
						}
						position++
					}
				l1219:
				}
			l1217:
				add(rulehexdigit, position1216)
			}
			return true
		l1215:
			position, tokenIndex = position1215, tokenIndex1215
			return false
		},
		/* 101 octaldigit <- <(<[0-7]> Action63)> */
		func() bool {
			position1221, tokenIndex1221 := position, tokenIndex
			{
				position1222 := position
				{
					position1223 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l1221
					}
					position++
					add(rulePegText, position1223)
				}
				{
					add(ruleAction63, position)
				}
				add(ruleoctaldigit, position1222)
			}
			return true
		l1221:
			position, tokenIndex = position1221, tokenIndex1221
			return false
		},
		/* 102 signedDecimalByte <- <(<('-'? [0-9]+)> Action64)> */
		nil,
		/* 103 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position1226, tokenIndex1226 := position, tokenIndex
			{
				position1227 := position
				{
					position1228, tokenIndex1228 := position, tokenIndex
					{
						position1230 := position
						{
							position1231, tokenIndex1231 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l1232
							}
							position++
							goto l1231
						l1232:
							position, tokenIndex = position1231, tokenIndex1231
							if buffer[position] != rune('N') {
								goto l1229
							}
							position++
						}
					l1231:
						{
							position1233, tokenIndex1233 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l1234
							}
							position++
							goto l1233
						l1234:
							position, tokenIndex = position1233, tokenIndex1233
							if buffer[position] != rune('Z') {
								goto l1229
							}
							position++
						}
					l1233:
						{
							add(ruleAction65, position)
						}
						add(ruleFT_NZ, position1230)
					}
					goto l1228
				l1229:
					position, tokenIndex = position1228, tokenIndex1228
					{
						position1237 := position
						{
							position1238, tokenIndex1238 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1239
							}
							position++
							goto l1238
						l1239:
							position, tokenIndex = position1238, tokenIndex1238
							if buffer[position] != rune('P') {
								goto l1236
							}
							position++
						}
					l1238:
						{
							position1240, tokenIndex1240 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l1241
							}
							position++
							goto l1240
						l1241:
							position, tokenIndex = position1240, tokenIndex1240
							if buffer[position] != rune('O') {
								goto l1236
							}
							position++
						}
					l1240:
						{
							add(ruleAction69, position)
						}
						add(ruleFT_PO, position1237)
					}
					goto l1228
				l1236:
					position, tokenIndex = position1228, tokenIndex1228
					{
						position1244 := position
						{
							position1245, tokenIndex1245 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1246
							}
							position++
							goto l1245
						l1246:
							position, tokenIndex = position1245, tokenIndex1245
							if buffer[position] != rune('P') {
								goto l1243
							}
							position++
						}
					l1245:
						{
							position1247, tokenIndex1247 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l1248
							}
							position++
							goto l1247
						l1248:
							position, tokenIndex = position1247, tokenIndex1247
							if buffer[position] != rune('E') {
								goto l1243
							}
							position++
						}
					l1247:
						{
							add(ruleAction70, position)
						}
						add(ruleFT_PE, position1244)
					}
					goto l1228
				l1243:
					position, tokenIndex = position1228, tokenIndex1228
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position1251 := position
								{
									position1252, tokenIndex1252 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l1253
									}
									position++
									goto l1252
								l1253:
									position, tokenIndex = position1252, tokenIndex1252
									if buffer[position] != rune('M') {
										goto l1226
									}
									position++
								}
							l1252:
								{
									add(ruleAction72, position)
								}
								add(ruleFT_M, position1251)
							}
							break
						case 'P', 'p':
							{
								position1255 := position
								{
									position1256, tokenIndex1256 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l1257
									}
									position++
									goto l1256
								l1257:
									position, tokenIndex = position1256, tokenIndex1256
									if buffer[position] != rune('P') {
										goto l1226
									}
									position++
								}
							l1256:
								{
									add(ruleAction71, position)
								}
								add(ruleFT_P, position1255)
							}
							break
						case 'C', 'c':
							{
								position1259 := position
								{
									position1260, tokenIndex1260 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1261
									}
									position++
									goto l1260
								l1261:
									position, tokenIndex = position1260, tokenIndex1260
									if buffer[position] != rune('C') {
										goto l1226
									}
									position++
								}
							l1260:
								{
									add(ruleAction68, position)
								}
								add(ruleFT_C, position1259)
							}
							break
						case 'N', 'n':
							{
								position1263 := position
								{
									position1264, tokenIndex1264 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l1265
									}
									position++
									goto l1264
								l1265:
									position, tokenIndex = position1264, tokenIndex1264
									if buffer[position] != rune('N') {
										goto l1226
									}
									position++
								}
							l1264:
								{
									position1266, tokenIndex1266 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1267
									}
									position++
									goto l1266
								l1267:
									position, tokenIndex = position1266, tokenIndex1266
									if buffer[position] != rune('C') {
										goto l1226
									}
									position++
								}
							l1266:
								{
									add(ruleAction67, position)
								}
								add(ruleFT_NC, position1263)
							}
							break
						default:
							{
								position1269 := position
								{
									position1270, tokenIndex1270 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l1271
									}
									position++
									goto l1270
								l1271:
									position, tokenIndex = position1270, tokenIndex1270
									if buffer[position] != rune('Z') {
										goto l1226
									}
									position++
								}
							l1270:
								{
									add(ruleAction66, position)
								}
								add(ruleFT_Z, position1269)
							}
							break
						}
					}

				}
			l1228:
				add(rulecc, position1227)
			}
			return true
		l1226:
			position, tokenIndex = position1226, tokenIndex1226
			return false
		},
		/* 104 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action65)> */
		nil,
		/* 105 FT_Z <- <(('z' / 'Z') Action66)> */
		nil,
		/* 106 FT_NC <- <(('n' / 'N') ('c' / 'C') Action67)> */
		nil,
		/* 107 FT_C <- <(('c' / 'C') Action68)> */
		nil,
		/* 108 FT_PO <- <(('p' / 'P') ('o' / 'O') Action69)> */
		nil,
		/* 109 FT_PE <- <(('p' / 'P') ('e' / 'E') Action70)> */
		nil,
		/* 110 FT_P <- <(('p' / 'P') Action71)> */
		nil,
		/* 111 FT_M <- <(('m' / 'M') Action72)> */
		nil,
		/* 113 Action0 <- <{ p.Emit() }> */
		nil,
		/* 114 Action1 <- <{ p.LD8() }> */
		nil,
		/* 115 Action2 <- <{ p.LD16() }> */
		nil,
		/* 116 Action3 <- <{ p.Push() }> */
		nil,
		/* 117 Action4 <- <{ p.Pop() }> */
		nil,
		/* 118 Action5 <- <{ p.Ex() }> */
		nil,
		/* 119 Action6 <- <{ p.Inc8() }> */
		nil,
		/* 120 Action7 <- <{ p.Inc16() }> */
		nil,
		/* 121 Action8 <- <{ p.Dec8() }> */
		nil,
		/* 122 Action9 <- <{ p.Dec16() }> */
		nil,
		/* 123 Action10 <- <{ p.Add16() }> */
		nil,
		/* 124 Action11 <- <{ p.Dst8() }> */
		nil,
		/* 125 Action12 <- <{ p.Src8() }> */
		nil,
		/* 126 Action13 <- <{ p.Loc8() }> */
		nil,
		nil,
		/* 128 Action14 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 129 Action15 <- <{ p.Dst16() }> */
		nil,
		/* 130 Action16 <- <{ p.Src16() }> */
		nil,
		/* 131 Action17 <- <{ p.Loc16() }> */
		nil,
		/* 132 Action18 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 133 Action19 <- <{ p.NNContents() }> */
		nil,
		/* 134 Action20 <- <{ p.Accum("ADD") }> */
		nil,
		/* 135 Action21 <- <{ p.Accum("ADC") }> */
		nil,
		/* 136 Action22 <- <{ p.Accum("SUB") }> */
		nil,
		/* 137 Action23 <- <{ p.Accum("SBC") }> */
		nil,
		/* 138 Action24 <- <{ p.Accum("AND") }> */
		nil,
		/* 139 Action25 <- <{ p.Accum("XOR") }> */
		nil,
		/* 140 Action26 <- <{ p.Accum("OR") }> */
		nil,
		/* 141 Action27 <- <{ p.Accum("CP") }> */
		nil,
		/* 142 Action28 <- <{ p.Rot("RLC") }> */
		nil,
		/* 143 Action29 <- <{ p.Rot("RRC") }> */
		nil,
		/* 144 Action30 <- <{ p.Rot("RL") }> */
		nil,
		/* 145 Action31 <- <{ p.Rot("RR") }> */
		nil,
		/* 146 Action32 <- <{ p.Rot("SLA") }> */
		nil,
		/* 147 Action33 <- <{ p.Rot("SRA") }> */
		nil,
		/* 148 Action34 <- <{ p.Rot("SLL") }> */
		nil,
		/* 149 Action35 <- <{ p.Rot("SRL") }> */
		nil,
		/* 150 Action36 <- <{ p.Bit() }> */
		nil,
		/* 151 Action37 <- <{ p.Res() }> */
		nil,
		/* 152 Action38 <- <{ p.Set() }> */
		nil,
		/* 153 Action39 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 154 Action40 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 155 Action41 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 156 Action42 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 157 Action43 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 158 Action44 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 159 Action45 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 160 Action46 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 161 Action47 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 162 Action48 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 163 Action49 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 164 Action50 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 165 Action51 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 166 Action52 <- <{ p.Rst() }> */
		nil,
		/* 167 Action53 <- <{ p.Call() }> */
		nil,
		/* 168 Action54 <- <{ p.Ret() }> */
		nil,
		/* 169 Action55 <- <{ p.Jp() }> */
		nil,
		/* 170 Action56 <- <{ p.Jr() }> */
		nil,
		/* 171 Action57 <- <{ p.Djnz() }> */
		nil,
		/* 172 Action58 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 173 Action59 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 174 Action60 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 175 Action61 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 176 Action62 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 177 Action63 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 178 Action64 <- <{ p.SignedDecimalByte(buffer[begin:end]) }> */
		nil,
		/* 179 Action65 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 180 Action66 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 181 Action67 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 182 Action68 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 183 Action69 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 184 Action70 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 185 Action71 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 186 Action72 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
