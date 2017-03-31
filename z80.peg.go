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
	ruleInc16Indexed8
	ruleInc8
	ruleInc16
	ruleDec
	ruleDec16Indexed8
	ruleDec8
	ruleDec16
	ruleAdd16
	ruleDst8
	ruleSrc8
	ruleLoc8
	ruleILoc8
	ruleReg8
	ruleIReg8
	ruleDst16
	ruleSrc16
	ruleLoc16
	ruleReg16
	ruleIReg16
	ruleReg16Contents
	rulePlainR16C
	ruleIndexedR16C
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
	ruleIO
	ruleIN
	ruleOUT
	rulePort
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
	ruleAction14
	ruleAction15
	ruleAction16
	rulePegText
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
	ruleAction73
	ruleAction74
	ruleAction75
	ruleAction76
	ruleAction77
	ruleAction78
	ruleAction79
	ruleAction80
	ruleAction81
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
	"Inc16Indexed8",
	"Inc8",
	"Inc16",
	"Dec",
	"Dec16Indexed8",
	"Dec8",
	"Dec16",
	"Add16",
	"Dst8",
	"Src8",
	"Loc8",
	"ILoc8",
	"Reg8",
	"IReg8",
	"Dst16",
	"Src16",
	"Loc16",
	"Reg16",
	"IReg16",
	"Reg16Contents",
	"PlainR16C",
	"IndexedR16C",
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
	"IO",
	"IN",
	"OUT",
	"Port",
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
	"Action14",
	"Action15",
	"Action16",
	"PegText",
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
	"Action73",
	"Action74",
	"Action75",
	"Action76",
	"Action77",
	"Action78",
	"Action79",
	"Action80",
	"Action81",
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
	rules  [208]func() bool
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
			p.Inc8()
		case ruleAction8:
			p.Inc16()
		case ruleAction9:
			p.Dec8()
		case ruleAction10:
			p.Dec8()
		case ruleAction11:
			p.Dec16()
		case ruleAction12:
			p.Add16()
		case ruleAction13:
			p.Dst8()
		case ruleAction14:
			p.Src8()
		case ruleAction15:
			p.Loc8()
		case ruleAction16:
			p.Loc8()
		case ruleAction17:
			p.R8(buffer[begin:end])
		case ruleAction18:
			p.R8(buffer[begin:end])
		case ruleAction19:
			p.Dst16()
		case ruleAction20:
			p.Src16()
		case ruleAction21:
			p.Loc16()
		case ruleAction22:
			p.R16(buffer[begin:end])
		case ruleAction23:
			p.R16(buffer[begin:end])
		case ruleAction24:
			p.R16Contents()
		case ruleAction25:
			p.IR16Contents()
		case ruleAction26:
			p.NNContents()
		case ruleAction27:
			p.Accum("ADD")
		case ruleAction28:
			p.Accum("ADC")
		case ruleAction29:
			p.Accum("SUB")
		case ruleAction30:
			p.Accum("SBC")
		case ruleAction31:
			p.Accum("AND")
		case ruleAction32:
			p.Accum("XOR")
		case ruleAction33:
			p.Accum("OR")
		case ruleAction34:
			p.Accum("CP")
		case ruleAction35:
			p.Rot("RLC")
		case ruleAction36:
			p.Rot("RRC")
		case ruleAction37:
			p.Rot("RL")
		case ruleAction38:
			p.Rot("RR")
		case ruleAction39:
			p.Rot("SLA")
		case ruleAction40:
			p.Rot("SRA")
		case ruleAction41:
			p.Rot("SLL")
		case ruleAction42:
			p.Rot("SRL")
		case ruleAction43:
			p.Bit()
		case ruleAction44:
			p.Res()
		case ruleAction45:
			p.Set()
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
			p.Simple(buffer[begin:end])
		case ruleAction53:
			p.Simple(buffer[begin:end])
		case ruleAction54:
			p.Simple(buffer[begin:end])
		case ruleAction55:
			p.Simple(buffer[begin:end])
		case ruleAction56:
			p.Simple(buffer[begin:end])
		case ruleAction57:
			p.Simple(buffer[begin:end])
		case ruleAction58:
			p.Simple(buffer[begin:end])
		case ruleAction59:
			p.Rst()
		case ruleAction60:
			p.Call()
		case ruleAction61:
			p.Ret()
		case ruleAction62:
			p.Jp()
		case ruleAction63:
			p.Jr()
		case ruleAction64:
			p.Djnz()
		case ruleAction65:
			p.In()
		case ruleAction66:
			p.Out()
		case ruleAction67:
			p.Nhex(buffer[begin:end])
		case ruleAction68:
			p.Nhex(buffer[begin:end])
		case ruleAction69:
			p.Ndec(buffer[begin:end])
		case ruleAction70:
			p.NNhex(buffer[begin:end])
		case ruleAction71:
			p.NNhex(buffer[begin:end])
		case ruleAction72:
			p.ODigit(buffer[begin:end])
		case ruleAction73:
			p.SignedDecimalByte(buffer[begin:end])
		case ruleAction74:
			p.Conditional(Not{FT_Z})
		case ruleAction75:
			p.Conditional(FT_Z)
		case ruleAction76:
			p.Conditional(Not{FT_C})
		case ruleAction77:
			p.Conditional(FT_C)
		case ruleAction78:
			p.Conditional(FT_PO)
		case ruleAction79:
			p.Conditional(FT_PE)
		case ruleAction80:
			p.Conditional(FT_P)
		case ruleAction81:
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
															if !_rules[ruleDst16]() {
																goto l45
															}
															if !_rules[rulesep]() {
																goto l45
															}
															if !_rules[ruleSrc16]() {
																goto l45
															}
															{
																add(ruleAction2, position)
															}
															add(ruleLoad16, position46)
														}
														goto l44
													l45:
														position, tokenIndex = position44, tokenIndex44
														{
															position52 := position
															{
																position53, tokenIndex53 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l54
																}
																position++
																goto l53
															l54:
																position, tokenIndex = position53, tokenIndex53
																if buffer[position] != rune('L') {
																	goto l14
																}
																position++
															}
														l53:
															{
																position55, tokenIndex55 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l56
																}
																position++
																goto l55
															l56:
																position, tokenIndex = position55, tokenIndex55
																if buffer[position] != rune('D') {
																	goto l14
																}
																position++
															}
														l55:
															if !_rules[rulews]() {
																goto l14
															}
															{
																position57 := position
																{
																	position58, tokenIndex58 := position, tokenIndex
																	if !_rules[ruleReg8]() {
																		goto l59
																	}
																	goto l58
																l59:
																	position, tokenIndex = position58, tokenIndex58
																	if !_rules[ruleReg16Contents]() {
																		goto l60
																	}
																	goto l58
																l60:
																	position, tokenIndex = position58, tokenIndex58
																	if !_rules[rulenn_contents]() {
																		goto l14
																	}
																}
															l58:
																{
																	add(ruleAction13, position)
																}
																add(ruleDst8, position57)
															}
															if !_rules[rulesep]() {
																goto l14
															}
															if !_rules[ruleSrc8]() {
																goto l14
															}
															{
																add(ruleAction1, position)
															}
															add(ruleLoad8, position52)
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
									position64 := position
									{
										position65, tokenIndex65 := position, tokenIndex
										{
											position67 := position
											{
												position68, tokenIndex68 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l69
												}
												position++
												goto l68
											l69:
												position, tokenIndex = position68, tokenIndex68
												if buffer[position] != rune('I') {
													goto l66
												}
												position++
											}
										l68:
											{
												position70, tokenIndex70 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l71
												}
												position++
												goto l70
											l71:
												position, tokenIndex = position70, tokenIndex70
												if buffer[position] != rune('N') {
													goto l66
												}
												position++
											}
										l70:
											{
												position72, tokenIndex72 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l73
												}
												position++
												goto l72
											l73:
												position, tokenIndex = position72, tokenIndex72
												if buffer[position] != rune('C') {
													goto l66
												}
												position++
											}
										l72:
											if !_rules[rulews]() {
												goto l66
											}
											if !_rules[ruleILoc8]() {
												goto l66
											}
											{
												add(ruleAction6, position)
											}
											add(ruleInc16Indexed8, position67)
										}
										goto l65
									l66:
										position, tokenIndex = position65, tokenIndex65
										{
											position76 := position
											{
												position77, tokenIndex77 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l78
												}
												position++
												goto l77
											l78:
												position, tokenIndex = position77, tokenIndex77
												if buffer[position] != rune('I') {
													goto l75
												}
												position++
											}
										l77:
											{
												position79, tokenIndex79 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l80
												}
												position++
												goto l79
											l80:
												position, tokenIndex = position79, tokenIndex79
												if buffer[position] != rune('N') {
													goto l75
												}
												position++
											}
										l79:
											{
												position81, tokenIndex81 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l82
												}
												position++
												goto l81
											l82:
												position, tokenIndex = position81, tokenIndex81
												if buffer[position] != rune('C') {
													goto l75
												}
												position++
											}
										l81:
											if !_rules[rulews]() {
												goto l75
											}
											if !_rules[ruleLoc16]() {
												goto l75
											}
											{
												add(ruleAction8, position)
											}
											add(ruleInc16, position76)
										}
										goto l65
									l75:
										position, tokenIndex = position65, tokenIndex65
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
													goto l63
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
													goto l63
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
													goto l63
												}
												position++
											}
										l89:
											if !_rules[rulews]() {
												goto l63
											}
											if !_rules[ruleLoc8]() {
												goto l63
											}
											{
												add(ruleAction7, position)
											}
											add(ruleInc8, position84)
										}
									}
								l65:
									add(ruleInc, position64)
								}
								goto l13
							l63:
								position, tokenIndex = position13, tokenIndex13
								{
									position93 := position
									{
										position94, tokenIndex94 := position, tokenIndex
										{
											position96 := position
											{
												position97, tokenIndex97 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l98
												}
												position++
												goto l97
											l98:
												position, tokenIndex = position97, tokenIndex97
												if buffer[position] != rune('D') {
													goto l95
												}
												position++
											}
										l97:
											{
												position99, tokenIndex99 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l100
												}
												position++
												goto l99
											l100:
												position, tokenIndex = position99, tokenIndex99
												if buffer[position] != rune('E') {
													goto l95
												}
												position++
											}
										l99:
											{
												position101, tokenIndex101 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l102
												}
												position++
												goto l101
											l102:
												position, tokenIndex = position101, tokenIndex101
												if buffer[position] != rune('C') {
													goto l95
												}
												position++
											}
										l101:
											if !_rules[rulews]() {
												goto l95
											}
											if !_rules[ruleILoc8]() {
												goto l95
											}
											{
												add(ruleAction9, position)
											}
											add(ruleDec16Indexed8, position96)
										}
										goto l94
									l95:
										position, tokenIndex = position94, tokenIndex94
										{
											position105 := position
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
													goto l104
												}
												position++
											}
										l106:
											{
												position108, tokenIndex108 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l109
												}
												position++
												goto l108
											l109:
												position, tokenIndex = position108, tokenIndex108
												if buffer[position] != rune('E') {
													goto l104
												}
												position++
											}
										l108:
											{
												position110, tokenIndex110 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l111
												}
												position++
												goto l110
											l111:
												position, tokenIndex = position110, tokenIndex110
												if buffer[position] != rune('C') {
													goto l104
												}
												position++
											}
										l110:
											if !_rules[rulews]() {
												goto l104
											}
											if !_rules[ruleLoc16]() {
												goto l104
											}
											{
												add(ruleAction11, position)
											}
											add(ruleDec16, position105)
										}
										goto l94
									l104:
										position, tokenIndex = position94, tokenIndex94
										{
											position113 := position
											{
												position114, tokenIndex114 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l115
												}
												position++
												goto l114
											l115:
												position, tokenIndex = position114, tokenIndex114
												if buffer[position] != rune('D') {
													goto l92
												}
												position++
											}
										l114:
											{
												position116, tokenIndex116 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l117
												}
												position++
												goto l116
											l117:
												position, tokenIndex = position116, tokenIndex116
												if buffer[position] != rune('E') {
													goto l92
												}
												position++
											}
										l116:
											{
												position118, tokenIndex118 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l119
												}
												position++
												goto l118
											l119:
												position, tokenIndex = position118, tokenIndex118
												if buffer[position] != rune('C') {
													goto l92
												}
												position++
											}
										l118:
											if !_rules[rulews]() {
												goto l92
											}
											if !_rules[ruleLoc8]() {
												goto l92
											}
											{
												add(ruleAction10, position)
											}
											add(ruleDec8, position113)
										}
									}
								l94:
									add(ruleDec, position93)
								}
								goto l13
							l92:
								position, tokenIndex = position13, tokenIndex13
								{
									position122 := position
									{
										position123, tokenIndex123 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l124
										}
										position++
										goto l123
									l124:
										position, tokenIndex = position123, tokenIndex123
										if buffer[position] != rune('A') {
											goto l121
										}
										position++
									}
								l123:
									{
										position125, tokenIndex125 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l126
										}
										position++
										goto l125
									l126:
										position, tokenIndex = position125, tokenIndex125
										if buffer[position] != rune('D') {
											goto l121
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
											goto l121
										}
										position++
									}
								l127:
									if !_rules[rulews]() {
										goto l121
									}
									if !_rules[ruleDst16]() {
										goto l121
									}
									if !_rules[rulesep]() {
										goto l121
									}
									if !_rules[ruleSrc16]() {
										goto l121
									}
									{
										add(ruleAction12, position)
									}
									add(ruleAdd16, position122)
								}
								goto l13
							l121:
								position, tokenIndex = position13, tokenIndex13
								{
									position131 := position
									{
										position132, tokenIndex132 := position, tokenIndex
										{
											position134 := position
											{
												position135, tokenIndex135 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l136
												}
												position++
												goto l135
											l136:
												position, tokenIndex = position135, tokenIndex135
												if buffer[position] != rune('A') {
													goto l133
												}
												position++
											}
										l135:
											{
												position137, tokenIndex137 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l138
												}
												position++
												goto l137
											l138:
												position, tokenIndex = position137, tokenIndex137
												if buffer[position] != rune('D') {
													goto l133
												}
												position++
											}
										l137:
											{
												position139, tokenIndex139 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l140
												}
												position++
												goto l139
											l140:
												position, tokenIndex = position139, tokenIndex139
												if buffer[position] != rune('D') {
													goto l133
												}
												position++
											}
										l139:
											if !_rules[rulews]() {
												goto l133
											}
											{
												position141, tokenIndex141 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l142
												}
												position++
												goto l141
											l142:
												position, tokenIndex = position141, tokenIndex141
												if buffer[position] != rune('A') {
													goto l133
												}
												position++
											}
										l141:
											if !_rules[rulesep]() {
												goto l133
											}
											if !_rules[ruleSrc8]() {
												goto l133
											}
											{
												add(ruleAction27, position)
											}
											add(ruleAdd, position134)
										}
										goto l132
									l133:
										position, tokenIndex = position132, tokenIndex132
										{
											position145 := position
											{
												position146, tokenIndex146 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l147
												}
												position++
												goto l146
											l147:
												position, tokenIndex = position146, tokenIndex146
												if buffer[position] != rune('A') {
													goto l144
												}
												position++
											}
										l146:
											{
												position148, tokenIndex148 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l149
												}
												position++
												goto l148
											l149:
												position, tokenIndex = position148, tokenIndex148
												if buffer[position] != rune('D') {
													goto l144
												}
												position++
											}
										l148:
											{
												position150, tokenIndex150 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l151
												}
												position++
												goto l150
											l151:
												position, tokenIndex = position150, tokenIndex150
												if buffer[position] != rune('C') {
													goto l144
												}
												position++
											}
										l150:
											if !_rules[rulews]() {
												goto l144
											}
											{
												position152, tokenIndex152 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l153
												}
												position++
												goto l152
											l153:
												position, tokenIndex = position152, tokenIndex152
												if buffer[position] != rune('A') {
													goto l144
												}
												position++
											}
										l152:
											if !_rules[rulesep]() {
												goto l144
											}
											if !_rules[ruleSrc8]() {
												goto l144
											}
											{
												add(ruleAction28, position)
											}
											add(ruleAdc, position145)
										}
										goto l132
									l144:
										position, tokenIndex = position132, tokenIndex132
										{
											position156 := position
											{
												position157, tokenIndex157 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l158
												}
												position++
												goto l157
											l158:
												position, tokenIndex = position157, tokenIndex157
												if buffer[position] != rune('S') {
													goto l155
												}
												position++
											}
										l157:
											{
												position159, tokenIndex159 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l160
												}
												position++
												goto l159
											l160:
												position, tokenIndex = position159, tokenIndex159
												if buffer[position] != rune('U') {
													goto l155
												}
												position++
											}
										l159:
											{
												position161, tokenIndex161 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l162
												}
												position++
												goto l161
											l162:
												position, tokenIndex = position161, tokenIndex161
												if buffer[position] != rune('B') {
													goto l155
												}
												position++
											}
										l161:
											if !_rules[rulews]() {
												goto l155
											}
											if !_rules[ruleSrc8]() {
												goto l155
											}
											{
												add(ruleAction29, position)
											}
											add(ruleSub, position156)
										}
										goto l132
									l155:
										position, tokenIndex = position132, tokenIndex132
										{
											switch buffer[position] {
											case 'C', 'c':
												{
													position165 := position
													{
														position166, tokenIndex166 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l167
														}
														position++
														goto l166
													l167:
														position, tokenIndex = position166, tokenIndex166
														if buffer[position] != rune('C') {
															goto l130
														}
														position++
													}
												l166:
													{
														position168, tokenIndex168 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l169
														}
														position++
														goto l168
													l169:
														position, tokenIndex = position168, tokenIndex168
														if buffer[position] != rune('P') {
															goto l130
														}
														position++
													}
												l168:
													if !_rules[rulews]() {
														goto l130
													}
													if !_rules[ruleSrc8]() {
														goto l130
													}
													{
														add(ruleAction34, position)
													}
													add(ruleCp, position165)
												}
												break
											case 'O', 'o':
												{
													position171 := position
													{
														position172, tokenIndex172 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l173
														}
														position++
														goto l172
													l173:
														position, tokenIndex = position172, tokenIndex172
														if buffer[position] != rune('O') {
															goto l130
														}
														position++
													}
												l172:
													{
														position174, tokenIndex174 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l175
														}
														position++
														goto l174
													l175:
														position, tokenIndex = position174, tokenIndex174
														if buffer[position] != rune('R') {
															goto l130
														}
														position++
													}
												l174:
													if !_rules[rulews]() {
														goto l130
													}
													if !_rules[ruleSrc8]() {
														goto l130
													}
													{
														add(ruleAction33, position)
													}
													add(ruleOr, position171)
												}
												break
											case 'X', 'x':
												{
													position177 := position
													{
														position178, tokenIndex178 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l179
														}
														position++
														goto l178
													l179:
														position, tokenIndex = position178, tokenIndex178
														if buffer[position] != rune('X') {
															goto l130
														}
														position++
													}
												l178:
													{
														position180, tokenIndex180 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l181
														}
														position++
														goto l180
													l181:
														position, tokenIndex = position180, tokenIndex180
														if buffer[position] != rune('O') {
															goto l130
														}
														position++
													}
												l180:
													{
														position182, tokenIndex182 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l183
														}
														position++
														goto l182
													l183:
														position, tokenIndex = position182, tokenIndex182
														if buffer[position] != rune('R') {
															goto l130
														}
														position++
													}
												l182:
													if !_rules[rulews]() {
														goto l130
													}
													if !_rules[ruleSrc8]() {
														goto l130
													}
													{
														add(ruleAction32, position)
													}
													add(ruleXor, position177)
												}
												break
											case 'A', 'a':
												{
													position185 := position
													{
														position186, tokenIndex186 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l187
														}
														position++
														goto l186
													l187:
														position, tokenIndex = position186, tokenIndex186
														if buffer[position] != rune('A') {
															goto l130
														}
														position++
													}
												l186:
													{
														position188, tokenIndex188 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l189
														}
														position++
														goto l188
													l189:
														position, tokenIndex = position188, tokenIndex188
														if buffer[position] != rune('N') {
															goto l130
														}
														position++
													}
												l188:
													{
														position190, tokenIndex190 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l191
														}
														position++
														goto l190
													l191:
														position, tokenIndex = position190, tokenIndex190
														if buffer[position] != rune('D') {
															goto l130
														}
														position++
													}
												l190:
													if !_rules[rulews]() {
														goto l130
													}
													if !_rules[ruleSrc8]() {
														goto l130
													}
													{
														add(ruleAction31, position)
													}
													add(ruleAnd, position185)
												}
												break
											default:
												{
													position193 := position
													{
														position194, tokenIndex194 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l195
														}
														position++
														goto l194
													l195:
														position, tokenIndex = position194, tokenIndex194
														if buffer[position] != rune('S') {
															goto l130
														}
														position++
													}
												l194:
													{
														position196, tokenIndex196 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l197
														}
														position++
														goto l196
													l197:
														position, tokenIndex = position196, tokenIndex196
														if buffer[position] != rune('B') {
															goto l130
														}
														position++
													}
												l196:
													{
														position198, tokenIndex198 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l199
														}
														position++
														goto l198
													l199:
														position, tokenIndex = position198, tokenIndex198
														if buffer[position] != rune('C') {
															goto l130
														}
														position++
													}
												l198:
													if !_rules[rulews]() {
														goto l130
													}
													{
														position200, tokenIndex200 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l201
														}
														position++
														goto l200
													l201:
														position, tokenIndex = position200, tokenIndex200
														if buffer[position] != rune('A') {
															goto l130
														}
														position++
													}
												l200:
													if !_rules[rulesep]() {
														goto l130
													}
													if !_rules[ruleSrc8]() {
														goto l130
													}
													{
														add(ruleAction30, position)
													}
													add(ruleSbc, position193)
												}
												break
											}
										}

									}
								l132:
									add(ruleAlu, position131)
								}
								goto l13
							l130:
								position, tokenIndex = position13, tokenIndex13
								{
									position204 := position
									{
										position205, tokenIndex205 := position, tokenIndex
										{
											position207 := position
											{
												position208, tokenIndex208 := position, tokenIndex
												{
													position210 := position
													{
														position211, tokenIndex211 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l212
														}
														position++
														goto l211
													l212:
														position, tokenIndex = position211, tokenIndex211
														if buffer[position] != rune('R') {
															goto l209
														}
														position++
													}
												l211:
													{
														position213, tokenIndex213 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l214
														}
														position++
														goto l213
													l214:
														position, tokenIndex = position213, tokenIndex213
														if buffer[position] != rune('L') {
															goto l209
														}
														position++
													}
												l213:
													{
														position215, tokenIndex215 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l216
														}
														position++
														goto l215
													l216:
														position, tokenIndex = position215, tokenIndex215
														if buffer[position] != rune('C') {
															goto l209
														}
														position++
													}
												l215:
													if !_rules[rulews]() {
														goto l209
													}
													if !_rules[ruleLoc8]() {
														goto l209
													}
													{
														add(ruleAction35, position)
													}
													add(ruleRlc, position210)
												}
												goto l208
											l209:
												position, tokenIndex = position208, tokenIndex208
												{
													position219 := position
													{
														position220, tokenIndex220 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l221
														}
														position++
														goto l220
													l221:
														position, tokenIndex = position220, tokenIndex220
														if buffer[position] != rune('R') {
															goto l218
														}
														position++
													}
												l220:
													{
														position222, tokenIndex222 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l223
														}
														position++
														goto l222
													l223:
														position, tokenIndex = position222, tokenIndex222
														if buffer[position] != rune('R') {
															goto l218
														}
														position++
													}
												l222:
													{
														position224, tokenIndex224 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l225
														}
														position++
														goto l224
													l225:
														position, tokenIndex = position224, tokenIndex224
														if buffer[position] != rune('C') {
															goto l218
														}
														position++
													}
												l224:
													if !_rules[rulews]() {
														goto l218
													}
													if !_rules[ruleLoc8]() {
														goto l218
													}
													{
														add(ruleAction36, position)
													}
													add(ruleRrc, position219)
												}
												goto l208
											l218:
												position, tokenIndex = position208, tokenIndex208
												{
													position228 := position
													{
														position229, tokenIndex229 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l230
														}
														position++
														goto l229
													l230:
														position, tokenIndex = position229, tokenIndex229
														if buffer[position] != rune('R') {
															goto l227
														}
														position++
													}
												l229:
													{
														position231, tokenIndex231 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l232
														}
														position++
														goto l231
													l232:
														position, tokenIndex = position231, tokenIndex231
														if buffer[position] != rune('L') {
															goto l227
														}
														position++
													}
												l231:
													if !_rules[rulews]() {
														goto l227
													}
													if !_rules[ruleLoc8]() {
														goto l227
													}
													{
														add(ruleAction37, position)
													}
													add(ruleRl, position228)
												}
												goto l208
											l227:
												position, tokenIndex = position208, tokenIndex208
												{
													position235 := position
													{
														position236, tokenIndex236 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l237
														}
														position++
														goto l236
													l237:
														position, tokenIndex = position236, tokenIndex236
														if buffer[position] != rune('R') {
															goto l234
														}
														position++
													}
												l236:
													{
														position238, tokenIndex238 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l239
														}
														position++
														goto l238
													l239:
														position, tokenIndex = position238, tokenIndex238
														if buffer[position] != rune('R') {
															goto l234
														}
														position++
													}
												l238:
													if !_rules[rulews]() {
														goto l234
													}
													if !_rules[ruleLoc8]() {
														goto l234
													}
													{
														add(ruleAction38, position)
													}
													add(ruleRr, position235)
												}
												goto l208
											l234:
												position, tokenIndex = position208, tokenIndex208
												{
													position242 := position
													{
														position243, tokenIndex243 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l244
														}
														position++
														goto l243
													l244:
														position, tokenIndex = position243, tokenIndex243
														if buffer[position] != rune('S') {
															goto l241
														}
														position++
													}
												l243:
													{
														position245, tokenIndex245 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l246
														}
														position++
														goto l245
													l246:
														position, tokenIndex = position245, tokenIndex245
														if buffer[position] != rune('L') {
															goto l241
														}
														position++
													}
												l245:
													{
														position247, tokenIndex247 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l248
														}
														position++
														goto l247
													l248:
														position, tokenIndex = position247, tokenIndex247
														if buffer[position] != rune('A') {
															goto l241
														}
														position++
													}
												l247:
													if !_rules[rulews]() {
														goto l241
													}
													if !_rules[ruleLoc8]() {
														goto l241
													}
													{
														add(ruleAction39, position)
													}
													add(ruleSla, position242)
												}
												goto l208
											l241:
												position, tokenIndex = position208, tokenIndex208
												{
													position251 := position
													{
														position252, tokenIndex252 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l253
														}
														position++
														goto l252
													l253:
														position, tokenIndex = position252, tokenIndex252
														if buffer[position] != rune('S') {
															goto l250
														}
														position++
													}
												l252:
													{
														position254, tokenIndex254 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l255
														}
														position++
														goto l254
													l255:
														position, tokenIndex = position254, tokenIndex254
														if buffer[position] != rune('R') {
															goto l250
														}
														position++
													}
												l254:
													{
														position256, tokenIndex256 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l257
														}
														position++
														goto l256
													l257:
														position, tokenIndex = position256, tokenIndex256
														if buffer[position] != rune('A') {
															goto l250
														}
														position++
													}
												l256:
													if !_rules[rulews]() {
														goto l250
													}
													if !_rules[ruleLoc8]() {
														goto l250
													}
													{
														add(ruleAction40, position)
													}
													add(ruleSra, position251)
												}
												goto l208
											l250:
												position, tokenIndex = position208, tokenIndex208
												{
													position260 := position
													{
														position261, tokenIndex261 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l262
														}
														position++
														goto l261
													l262:
														position, tokenIndex = position261, tokenIndex261
														if buffer[position] != rune('S') {
															goto l259
														}
														position++
													}
												l261:
													{
														position263, tokenIndex263 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l264
														}
														position++
														goto l263
													l264:
														position, tokenIndex = position263, tokenIndex263
														if buffer[position] != rune('L') {
															goto l259
														}
														position++
													}
												l263:
													{
														position265, tokenIndex265 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l266
														}
														position++
														goto l265
													l266:
														position, tokenIndex = position265, tokenIndex265
														if buffer[position] != rune('L') {
															goto l259
														}
														position++
													}
												l265:
													if !_rules[rulews]() {
														goto l259
													}
													if !_rules[ruleLoc8]() {
														goto l259
													}
													{
														add(ruleAction41, position)
													}
													add(ruleSll, position260)
												}
												goto l208
											l259:
												position, tokenIndex = position208, tokenIndex208
												{
													position268 := position
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
															goto l206
														}
														position++
													}
												l269:
													{
														position271, tokenIndex271 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l272
														}
														position++
														goto l271
													l272:
														position, tokenIndex = position271, tokenIndex271
														if buffer[position] != rune('R') {
															goto l206
														}
														position++
													}
												l271:
													{
														position273, tokenIndex273 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l274
														}
														position++
														goto l273
													l274:
														position, tokenIndex = position273, tokenIndex273
														if buffer[position] != rune('L') {
															goto l206
														}
														position++
													}
												l273:
													if !_rules[rulews]() {
														goto l206
													}
													if !_rules[ruleLoc8]() {
														goto l206
													}
													{
														add(ruleAction42, position)
													}
													add(ruleSrl, position268)
												}
											}
										l208:
											add(ruleRot, position207)
										}
										goto l205
									l206:
										position, tokenIndex = position205, tokenIndex205
										{
											switch buffer[position] {
											case 'S', 's':
												{
													position277 := position
													{
														position278, tokenIndex278 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l279
														}
														position++
														goto l278
													l279:
														position, tokenIndex = position278, tokenIndex278
														if buffer[position] != rune('S') {
															goto l203
														}
														position++
													}
												l278:
													{
														position280, tokenIndex280 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l281
														}
														position++
														goto l280
													l281:
														position, tokenIndex = position280, tokenIndex280
														if buffer[position] != rune('E') {
															goto l203
														}
														position++
													}
												l280:
													{
														position282, tokenIndex282 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l283
														}
														position++
														goto l282
													l283:
														position, tokenIndex = position282, tokenIndex282
														if buffer[position] != rune('T') {
															goto l203
														}
														position++
													}
												l282:
													if !_rules[rulews]() {
														goto l203
													}
													if !_rules[ruleoctaldigit]() {
														goto l203
													}
													if !_rules[rulesep]() {
														goto l203
													}
													if !_rules[ruleLoc8]() {
														goto l203
													}
													{
														add(ruleAction45, position)
													}
													add(ruleSet, position277)
												}
												break
											case 'R', 'r':
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
															goto l203
														}
														position++
													}
												l286:
													{
														position288, tokenIndex288 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l289
														}
														position++
														goto l288
													l289:
														position, tokenIndex = position288, tokenIndex288
														if buffer[position] != rune('E') {
															goto l203
														}
														position++
													}
												l288:
													{
														position290, tokenIndex290 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l291
														}
														position++
														goto l290
													l291:
														position, tokenIndex = position290, tokenIndex290
														if buffer[position] != rune('S') {
															goto l203
														}
														position++
													}
												l290:
													if !_rules[rulews]() {
														goto l203
													}
													if !_rules[ruleoctaldigit]() {
														goto l203
													}
													if !_rules[rulesep]() {
														goto l203
													}
													if !_rules[ruleLoc8]() {
														goto l203
													}
													{
														add(ruleAction44, position)
													}
													add(ruleRes, position285)
												}
												break
											default:
												{
													position293 := position
													{
														position294, tokenIndex294 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l295
														}
														position++
														goto l294
													l295:
														position, tokenIndex = position294, tokenIndex294
														if buffer[position] != rune('B') {
															goto l203
														}
														position++
													}
												l294:
													{
														position296, tokenIndex296 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l297
														}
														position++
														goto l296
													l297:
														position, tokenIndex = position296, tokenIndex296
														if buffer[position] != rune('I') {
															goto l203
														}
														position++
													}
												l296:
													{
														position298, tokenIndex298 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l299
														}
														position++
														goto l298
													l299:
														position, tokenIndex = position298, tokenIndex298
														if buffer[position] != rune('T') {
															goto l203
														}
														position++
													}
												l298:
													if !_rules[rulews]() {
														goto l203
													}
													if !_rules[ruleoctaldigit]() {
														goto l203
													}
													if !_rules[rulesep]() {
														goto l203
													}
													if !_rules[ruleLoc8]() {
														goto l203
													}
													{
														add(ruleAction43, position)
													}
													add(ruleBit, position293)
												}
												break
											}
										}

									}
								l205:
									add(ruleBitOp, position204)
								}
								goto l13
							l203:
								position, tokenIndex = position13, tokenIndex13
								{
									position302 := position
									{
										position303, tokenIndex303 := position, tokenIndex
										{
											position305 := position
											{
												position306 := position
												{
													position307, tokenIndex307 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l308
													}
													position++
													goto l307
												l308:
													position, tokenIndex = position307, tokenIndex307
													if buffer[position] != rune('R') {
														goto l304
													}
													position++
												}
											l307:
												{
													position309, tokenIndex309 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l310
													}
													position++
													goto l309
												l310:
													position, tokenIndex = position309, tokenIndex309
													if buffer[position] != rune('L') {
														goto l304
													}
													position++
												}
											l309:
												{
													position311, tokenIndex311 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l312
													}
													position++
													goto l311
												l312:
													position, tokenIndex = position311, tokenIndex311
													if buffer[position] != rune('C') {
														goto l304
													}
													position++
												}
											l311:
												{
													position313, tokenIndex313 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l314
													}
													position++
													goto l313
												l314:
													position, tokenIndex = position313, tokenIndex313
													if buffer[position] != rune('A') {
														goto l304
													}
													position++
												}
											l313:
												add(rulePegText, position306)
											}
											{
												add(ruleAction48, position)
											}
											add(ruleRlca, position305)
										}
										goto l303
									l304:
										position, tokenIndex = position303, tokenIndex303
										{
											position317 := position
											{
												position318 := position
												{
													position319, tokenIndex319 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l320
													}
													position++
													goto l319
												l320:
													position, tokenIndex = position319, tokenIndex319
													if buffer[position] != rune('R') {
														goto l316
													}
													position++
												}
											l319:
												{
													position321, tokenIndex321 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l322
													}
													position++
													goto l321
												l322:
													position, tokenIndex = position321, tokenIndex321
													if buffer[position] != rune('R') {
														goto l316
													}
													position++
												}
											l321:
												{
													position323, tokenIndex323 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l324
													}
													position++
													goto l323
												l324:
													position, tokenIndex = position323, tokenIndex323
													if buffer[position] != rune('C') {
														goto l316
													}
													position++
												}
											l323:
												{
													position325, tokenIndex325 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l326
													}
													position++
													goto l325
												l326:
													position, tokenIndex = position325, tokenIndex325
													if buffer[position] != rune('A') {
														goto l316
													}
													position++
												}
											l325:
												add(rulePegText, position318)
											}
											{
												add(ruleAction49, position)
											}
											add(ruleRrca, position317)
										}
										goto l303
									l316:
										position, tokenIndex = position303, tokenIndex303
										{
											position329 := position
											{
												position330 := position
												{
													position331, tokenIndex331 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l332
													}
													position++
													goto l331
												l332:
													position, tokenIndex = position331, tokenIndex331
													if buffer[position] != rune('R') {
														goto l328
													}
													position++
												}
											l331:
												{
													position333, tokenIndex333 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l334
													}
													position++
													goto l333
												l334:
													position, tokenIndex = position333, tokenIndex333
													if buffer[position] != rune('L') {
														goto l328
													}
													position++
												}
											l333:
												{
													position335, tokenIndex335 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l336
													}
													position++
													goto l335
												l336:
													position, tokenIndex = position335, tokenIndex335
													if buffer[position] != rune('A') {
														goto l328
													}
													position++
												}
											l335:
												add(rulePegText, position330)
											}
											{
												add(ruleAction50, position)
											}
											add(ruleRla, position329)
										}
										goto l303
									l328:
										position, tokenIndex = position303, tokenIndex303
										{
											position339 := position
											{
												position340 := position
												{
													position341, tokenIndex341 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l342
													}
													position++
													goto l341
												l342:
													position, tokenIndex = position341, tokenIndex341
													if buffer[position] != rune('D') {
														goto l338
													}
													position++
												}
											l341:
												{
													position343, tokenIndex343 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l344
													}
													position++
													goto l343
												l344:
													position, tokenIndex = position343, tokenIndex343
													if buffer[position] != rune('A') {
														goto l338
													}
													position++
												}
											l343:
												{
													position345, tokenIndex345 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l346
													}
													position++
													goto l345
												l346:
													position, tokenIndex = position345, tokenIndex345
													if buffer[position] != rune('A') {
														goto l338
													}
													position++
												}
											l345:
												add(rulePegText, position340)
											}
											{
												add(ruleAction52, position)
											}
											add(ruleDaa, position339)
										}
										goto l303
									l338:
										position, tokenIndex = position303, tokenIndex303
										{
											position349 := position
											{
												position350 := position
												{
													position351, tokenIndex351 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l352
													}
													position++
													goto l351
												l352:
													position, tokenIndex = position351, tokenIndex351
													if buffer[position] != rune('C') {
														goto l348
													}
													position++
												}
											l351:
												{
													position353, tokenIndex353 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l354
													}
													position++
													goto l353
												l354:
													position, tokenIndex = position353, tokenIndex353
													if buffer[position] != rune('P') {
														goto l348
													}
													position++
												}
											l353:
												{
													position355, tokenIndex355 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l356
													}
													position++
													goto l355
												l356:
													position, tokenIndex = position355, tokenIndex355
													if buffer[position] != rune('L') {
														goto l348
													}
													position++
												}
											l355:
												add(rulePegText, position350)
											}
											{
												add(ruleAction53, position)
											}
											add(ruleCpl, position349)
										}
										goto l303
									l348:
										position, tokenIndex = position303, tokenIndex303
										{
											position359 := position
											{
												position360 := position
												{
													position361, tokenIndex361 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l362
													}
													position++
													goto l361
												l362:
													position, tokenIndex = position361, tokenIndex361
													if buffer[position] != rune('E') {
														goto l358
													}
													position++
												}
											l361:
												{
													position363, tokenIndex363 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l364
													}
													position++
													goto l363
												l364:
													position, tokenIndex = position363, tokenIndex363
													if buffer[position] != rune('X') {
														goto l358
													}
													position++
												}
											l363:
												{
													position365, tokenIndex365 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l366
													}
													position++
													goto l365
												l366:
													position, tokenIndex = position365, tokenIndex365
													if buffer[position] != rune('X') {
														goto l358
													}
													position++
												}
											l365:
												add(rulePegText, position360)
											}
											{
												add(ruleAction56, position)
											}
											add(ruleExx, position359)
										}
										goto l303
									l358:
										position, tokenIndex = position303, tokenIndex303
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position369 := position
													{
														position370 := position
														{
															position371, tokenIndex371 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l372
															}
															position++
															goto l371
														l372:
															position, tokenIndex = position371, tokenIndex371
															if buffer[position] != rune('E') {
																goto l301
															}
															position++
														}
													l371:
														{
															position373, tokenIndex373 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l374
															}
															position++
															goto l373
														l374:
															position, tokenIndex = position373, tokenIndex373
															if buffer[position] != rune('I') {
																goto l301
															}
															position++
														}
													l373:
														add(rulePegText, position370)
													}
													{
														add(ruleAction58, position)
													}
													add(ruleEi, position369)
												}
												break
											case 'D', 'd':
												{
													position376 := position
													{
														position377 := position
														{
															position378, tokenIndex378 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l379
															}
															position++
															goto l378
														l379:
															position, tokenIndex = position378, tokenIndex378
															if buffer[position] != rune('D') {
																goto l301
															}
															position++
														}
													l378:
														{
															position380, tokenIndex380 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l381
															}
															position++
															goto l380
														l381:
															position, tokenIndex = position380, tokenIndex380
															if buffer[position] != rune('I') {
																goto l301
															}
															position++
														}
													l380:
														add(rulePegText, position377)
													}
													{
														add(ruleAction57, position)
													}
													add(ruleDi, position376)
												}
												break
											case 'C', 'c':
												{
													position383 := position
													{
														position384 := position
														{
															position385, tokenIndex385 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l386
															}
															position++
															goto l385
														l386:
															position, tokenIndex = position385, tokenIndex385
															if buffer[position] != rune('C') {
																goto l301
															}
															position++
														}
													l385:
														{
															position387, tokenIndex387 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l388
															}
															position++
															goto l387
														l388:
															position, tokenIndex = position387, tokenIndex387
															if buffer[position] != rune('C') {
																goto l301
															}
															position++
														}
													l387:
														{
															position389, tokenIndex389 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l390
															}
															position++
															goto l389
														l390:
															position, tokenIndex = position389, tokenIndex389
															if buffer[position] != rune('F') {
																goto l301
															}
															position++
														}
													l389:
														add(rulePegText, position384)
													}
													{
														add(ruleAction55, position)
													}
													add(ruleCcf, position383)
												}
												break
											case 'S', 's':
												{
													position392 := position
													{
														position393 := position
														{
															position394, tokenIndex394 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l395
															}
															position++
															goto l394
														l395:
															position, tokenIndex = position394, tokenIndex394
															if buffer[position] != rune('S') {
																goto l301
															}
															position++
														}
													l394:
														{
															position396, tokenIndex396 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l397
															}
															position++
															goto l396
														l397:
															position, tokenIndex = position396, tokenIndex396
															if buffer[position] != rune('C') {
																goto l301
															}
															position++
														}
													l396:
														{
															position398, tokenIndex398 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l399
															}
															position++
															goto l398
														l399:
															position, tokenIndex = position398, tokenIndex398
															if buffer[position] != rune('F') {
																goto l301
															}
															position++
														}
													l398:
														add(rulePegText, position393)
													}
													{
														add(ruleAction54, position)
													}
													add(ruleScf, position392)
												}
												break
											case 'R', 'r':
												{
													position401 := position
													{
														position402 := position
														{
															position403, tokenIndex403 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l404
															}
															position++
															goto l403
														l404:
															position, tokenIndex = position403, tokenIndex403
															if buffer[position] != rune('R') {
																goto l301
															}
															position++
														}
													l403:
														{
															position405, tokenIndex405 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l406
															}
															position++
															goto l405
														l406:
															position, tokenIndex = position405, tokenIndex405
															if buffer[position] != rune('R') {
																goto l301
															}
															position++
														}
													l405:
														{
															position407, tokenIndex407 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l408
															}
															position++
															goto l407
														l408:
															position, tokenIndex = position407, tokenIndex407
															if buffer[position] != rune('A') {
																goto l301
															}
															position++
														}
													l407:
														add(rulePegText, position402)
													}
													{
														add(ruleAction51, position)
													}
													add(ruleRra, position401)
												}
												break
											case 'H', 'h':
												{
													position410 := position
													{
														position411 := position
														{
															position412, tokenIndex412 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l413
															}
															position++
															goto l412
														l413:
															position, tokenIndex = position412, tokenIndex412
															if buffer[position] != rune('H') {
																goto l301
															}
															position++
														}
													l412:
														{
															position414, tokenIndex414 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l415
															}
															position++
															goto l414
														l415:
															position, tokenIndex = position414, tokenIndex414
															if buffer[position] != rune('A') {
																goto l301
															}
															position++
														}
													l414:
														{
															position416, tokenIndex416 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l417
															}
															position++
															goto l416
														l417:
															position, tokenIndex = position416, tokenIndex416
															if buffer[position] != rune('L') {
																goto l301
															}
															position++
														}
													l416:
														{
															position418, tokenIndex418 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l419
															}
															position++
															goto l418
														l419:
															position, tokenIndex = position418, tokenIndex418
															if buffer[position] != rune('T') {
																goto l301
															}
															position++
														}
													l418:
														add(rulePegText, position411)
													}
													{
														add(ruleAction47, position)
													}
													add(ruleHalt, position410)
												}
												break
											default:
												{
													position421 := position
													{
														position422 := position
														{
															position423, tokenIndex423 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l424
															}
															position++
															goto l423
														l424:
															position, tokenIndex = position423, tokenIndex423
															if buffer[position] != rune('N') {
																goto l301
															}
															position++
														}
													l423:
														{
															position425, tokenIndex425 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l426
															}
															position++
															goto l425
														l426:
															position, tokenIndex = position425, tokenIndex425
															if buffer[position] != rune('O') {
																goto l301
															}
															position++
														}
													l425:
														{
															position427, tokenIndex427 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l428
															}
															position++
															goto l427
														l428:
															position, tokenIndex = position427, tokenIndex427
															if buffer[position] != rune('P') {
																goto l301
															}
															position++
														}
													l427:
														add(rulePegText, position422)
													}
													{
														add(ruleAction46, position)
													}
													add(ruleNop, position421)
												}
												break
											}
										}

									}
								l303:
									add(ruleSimple, position302)
								}
								goto l13
							l301:
								position, tokenIndex = position13, tokenIndex13
								{
									position431 := position
									{
										position432, tokenIndex432 := position, tokenIndex
										{
											position434 := position
											{
												position435, tokenIndex435 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l436
												}
												position++
												goto l435
											l436:
												position, tokenIndex = position435, tokenIndex435
												if buffer[position] != rune('R') {
													goto l433
												}
												position++
											}
										l435:
											{
												position437, tokenIndex437 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l438
												}
												position++
												goto l437
											l438:
												position, tokenIndex = position437, tokenIndex437
												if buffer[position] != rune('S') {
													goto l433
												}
												position++
											}
										l437:
											{
												position439, tokenIndex439 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l440
												}
												position++
												goto l439
											l440:
												position, tokenIndex = position439, tokenIndex439
												if buffer[position] != rune('T') {
													goto l433
												}
												position++
											}
										l439:
											if !_rules[rulews]() {
												goto l433
											}
											if !_rules[rulen]() {
												goto l433
											}
											{
												add(ruleAction59, position)
											}
											add(ruleRst, position434)
										}
										goto l432
									l433:
										position, tokenIndex = position432, tokenIndex432
										{
											position443 := position
											{
												position444, tokenIndex444 := position, tokenIndex
												if buffer[position] != rune('j') {
													goto l445
												}
												position++
												goto l444
											l445:
												position, tokenIndex = position444, tokenIndex444
												if buffer[position] != rune('J') {
													goto l442
												}
												position++
											}
										l444:
											{
												position446, tokenIndex446 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l447
												}
												position++
												goto l446
											l447:
												position, tokenIndex = position446, tokenIndex446
												if buffer[position] != rune('P') {
													goto l442
												}
												position++
											}
										l446:
											if !_rules[rulews]() {
												goto l442
											}
											{
												position448, tokenIndex448 := position, tokenIndex
												if !_rules[rulecc]() {
													goto l448
												}
												if !_rules[rulesep]() {
													goto l448
												}
												goto l449
											l448:
												position, tokenIndex = position448, tokenIndex448
											}
										l449:
											if !_rules[ruleSrc16]() {
												goto l442
											}
											{
												add(ruleAction62, position)
											}
											add(ruleJp, position443)
										}
										goto l432
									l442:
										position, tokenIndex = position432, tokenIndex432
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position452 := position
													{
														position453, tokenIndex453 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l454
														}
														position++
														goto l453
													l454:
														position, tokenIndex = position453, tokenIndex453
														if buffer[position] != rune('D') {
															goto l430
														}
														position++
													}
												l453:
													{
														position455, tokenIndex455 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l456
														}
														position++
														goto l455
													l456:
														position, tokenIndex = position455, tokenIndex455
														if buffer[position] != rune('J') {
															goto l430
														}
														position++
													}
												l455:
													{
														position457, tokenIndex457 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l458
														}
														position++
														goto l457
													l458:
														position, tokenIndex = position457, tokenIndex457
														if buffer[position] != rune('N') {
															goto l430
														}
														position++
													}
												l457:
													{
														position459, tokenIndex459 := position, tokenIndex
														if buffer[position] != rune('z') {
															goto l460
														}
														position++
														goto l459
													l460:
														position, tokenIndex = position459, tokenIndex459
														if buffer[position] != rune('Z') {
															goto l430
														}
														position++
													}
												l459:
													if !_rules[rulews]() {
														goto l430
													}
													if !_rules[ruledisp]() {
														goto l430
													}
													{
														add(ruleAction64, position)
													}
													add(ruleDjnz, position452)
												}
												break
											case 'J', 'j':
												{
													position462 := position
													{
														position463, tokenIndex463 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l464
														}
														position++
														goto l463
													l464:
														position, tokenIndex = position463, tokenIndex463
														if buffer[position] != rune('J') {
															goto l430
														}
														position++
													}
												l463:
													{
														position465, tokenIndex465 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l466
														}
														position++
														goto l465
													l466:
														position, tokenIndex = position465, tokenIndex465
														if buffer[position] != rune('R') {
															goto l430
														}
														position++
													}
												l465:
													if !_rules[rulews]() {
														goto l430
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
													if !_rules[ruledisp]() {
														goto l430
													}
													{
														add(ruleAction63, position)
													}
													add(ruleJr, position462)
												}
												break
											case 'R', 'r':
												{
													position470 := position
													{
														position471, tokenIndex471 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l472
														}
														position++
														goto l471
													l472:
														position, tokenIndex = position471, tokenIndex471
														if buffer[position] != rune('R') {
															goto l430
														}
														position++
													}
												l471:
													{
														position473, tokenIndex473 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l474
														}
														position++
														goto l473
													l474:
														position, tokenIndex = position473, tokenIndex473
														if buffer[position] != rune('E') {
															goto l430
														}
														position++
													}
												l473:
													{
														position475, tokenIndex475 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l476
														}
														position++
														goto l475
													l476:
														position, tokenIndex = position475, tokenIndex475
														if buffer[position] != rune('T') {
															goto l430
														}
														position++
													}
												l475:
													{
														position477, tokenIndex477 := position, tokenIndex
														if !_rules[rulews]() {
															goto l477
														}
														if !_rules[rulecc]() {
															goto l477
														}
														goto l478
													l477:
														position, tokenIndex = position477, tokenIndex477
													}
												l478:
													{
														add(ruleAction61, position)
													}
													add(ruleRet, position470)
												}
												break
											default:
												{
													position480 := position
													{
														position481, tokenIndex481 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l482
														}
														position++
														goto l481
													l482:
														position, tokenIndex = position481, tokenIndex481
														if buffer[position] != rune('C') {
															goto l430
														}
														position++
													}
												l481:
													{
														position483, tokenIndex483 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l484
														}
														position++
														goto l483
													l484:
														position, tokenIndex = position483, tokenIndex483
														if buffer[position] != rune('A') {
															goto l430
														}
														position++
													}
												l483:
													{
														position485, tokenIndex485 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l486
														}
														position++
														goto l485
													l486:
														position, tokenIndex = position485, tokenIndex485
														if buffer[position] != rune('L') {
															goto l430
														}
														position++
													}
												l485:
													{
														position487, tokenIndex487 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l488
														}
														position++
														goto l487
													l488:
														position, tokenIndex = position487, tokenIndex487
														if buffer[position] != rune('L') {
															goto l430
														}
														position++
													}
												l487:
													if !_rules[rulews]() {
														goto l430
													}
													{
														position489, tokenIndex489 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l489
														}
														if !_rules[rulesep]() {
															goto l489
														}
														goto l490
													l489:
														position, tokenIndex = position489, tokenIndex489
													}
												l490:
													if !_rules[ruleSrc16]() {
														goto l430
													}
													{
														add(ruleAction60, position)
													}
													add(ruleCall, position480)
												}
												break
											}
										}

									}
								l432:
									add(ruleJump, position431)
								}
								goto l13
							l430:
								position, tokenIndex = position13, tokenIndex13
								{
									position492 := position
									{
										position493, tokenIndex493 := position, tokenIndex
										{
											position495 := position
											{
												position496, tokenIndex496 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l497
												}
												position++
												goto l496
											l497:
												position, tokenIndex = position496, tokenIndex496
												if buffer[position] != rune('I') {
													goto l494
												}
												position++
											}
										l496:
											{
												position498, tokenIndex498 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l499
												}
												position++
												goto l498
											l499:
												position, tokenIndex = position498, tokenIndex498
												if buffer[position] != rune('N') {
													goto l494
												}
												position++
											}
										l498:
											if !_rules[rulews]() {
												goto l494
											}
											if !_rules[ruleReg8]() {
												goto l494
											}
											if !_rules[rulesep]() {
												goto l494
											}
											if !_rules[rulePort]() {
												goto l494
											}
											{
												add(ruleAction65, position)
											}
											add(ruleIN, position495)
										}
										goto l493
									l494:
										position, tokenIndex = position493, tokenIndex493
										{
											position501 := position
											{
												position502, tokenIndex502 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l503
												}
												position++
												goto l502
											l503:
												position, tokenIndex = position502, tokenIndex502
												if buffer[position] != rune('O') {
													goto l0
												}
												position++
											}
										l502:
											{
												position504, tokenIndex504 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l505
												}
												position++
												goto l504
											l505:
												position, tokenIndex = position504, tokenIndex504
												if buffer[position] != rune('U') {
													goto l0
												}
												position++
											}
										l504:
											{
												position506, tokenIndex506 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l507
												}
												position++
												goto l506
											l507:
												position, tokenIndex = position506, tokenIndex506
												if buffer[position] != rune('T') {
													goto l0
												}
												position++
											}
										l506:
											if !_rules[rulews]() {
												goto l0
											}
											if !_rules[rulePort]() {
												goto l0
											}
											if !_rules[rulesep]() {
												goto l0
											}
											if !_rules[ruleReg8]() {
												goto l0
											}
											{
												add(ruleAction66, position)
											}
											add(ruleOUT, position501)
										}
									}
								l493:
									add(ruleIO, position492)
								}
							}
						l13:
							add(ruleInstruction, position10)
						}
						{
							position509, tokenIndex509 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l509
							}
							position++
							goto l510
						l509:
							position, tokenIndex = position509, tokenIndex509
						}
					l510:
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
						position512, tokenIndex512 := position, tokenIndex
						{
							position514 := position
						l515:
							{
								position516, tokenIndex516 := position, tokenIndex
								if !_rules[rulews]() {
									goto l516
								}
								goto l515
							l516:
								position, tokenIndex = position516, tokenIndex516
							}
							if buffer[position] != rune('\n') {
								goto l513
							}
							position++
							add(ruleBlankLine, position514)
						}
						goto l512
					l513:
						position, tokenIndex = position512, tokenIndex512
						{
							position517 := position
							{
								position518 := position
							l519:
								{
									position520, tokenIndex520 := position, tokenIndex
									if !_rules[rulews]() {
										goto l520
									}
									goto l519
								l520:
									position, tokenIndex = position520, tokenIndex520
								}
								{
									position521, tokenIndex521 := position, tokenIndex
									{
										position523 := position
										{
											position524, tokenIndex524 := position, tokenIndex
											{
												position526 := position
												{
													position527, tokenIndex527 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l528
													}
													position++
													goto l527
												l528:
													position, tokenIndex = position527, tokenIndex527
													if buffer[position] != rune('P') {
														goto l525
													}
													position++
												}
											l527:
												{
													position529, tokenIndex529 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l530
													}
													position++
													goto l529
												l530:
													position, tokenIndex = position529, tokenIndex529
													if buffer[position] != rune('U') {
														goto l525
													}
													position++
												}
											l529:
												{
													position531, tokenIndex531 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l532
													}
													position++
													goto l531
												l532:
													position, tokenIndex = position531, tokenIndex531
													if buffer[position] != rune('S') {
														goto l525
													}
													position++
												}
											l531:
												{
													position533, tokenIndex533 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l534
													}
													position++
													goto l533
												l534:
													position, tokenIndex = position533, tokenIndex533
													if buffer[position] != rune('H') {
														goto l525
													}
													position++
												}
											l533:
												if !_rules[rulews]() {
													goto l525
												}
												if !_rules[ruleSrc16]() {
													goto l525
												}
												{
													add(ruleAction3, position)
												}
												add(rulePush, position526)
											}
											goto l524
										l525:
											position, tokenIndex = position524, tokenIndex524
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position537 := position
														{
															position538, tokenIndex538 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l539
															}
															position++
															goto l538
														l539:
															position, tokenIndex = position538, tokenIndex538
															if buffer[position] != rune('E') {
																goto l522
															}
															position++
														}
													l538:
														{
															position540, tokenIndex540 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l541
															}
															position++
															goto l540
														l541:
															position, tokenIndex = position540, tokenIndex540
															if buffer[position] != rune('X') {
																goto l522
															}
															position++
														}
													l540:
														if !_rules[rulews]() {
															goto l522
														}
														if !_rules[ruleDst16]() {
															goto l522
														}
														if !_rules[rulesep]() {
															goto l522
														}
														if !_rules[ruleSrc16]() {
															goto l522
														}
														{
															add(ruleAction5, position)
														}
														add(ruleEx, position537)
													}
													break
												case 'P', 'p':
													{
														position543 := position
														{
															position544, tokenIndex544 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l545
															}
															position++
															goto l544
														l545:
															position, tokenIndex = position544, tokenIndex544
															if buffer[position] != rune('P') {
																goto l522
															}
															position++
														}
													l544:
														{
															position546, tokenIndex546 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l547
															}
															position++
															goto l546
														l547:
															position, tokenIndex = position546, tokenIndex546
															if buffer[position] != rune('O') {
																goto l522
															}
															position++
														}
													l546:
														{
															position548, tokenIndex548 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l549
															}
															position++
															goto l548
														l549:
															position, tokenIndex = position548, tokenIndex548
															if buffer[position] != rune('P') {
																goto l522
															}
															position++
														}
													l548:
														if !_rules[rulews]() {
															goto l522
														}
														if !_rules[ruleDst16]() {
															goto l522
														}
														{
															add(ruleAction4, position)
														}
														add(rulePop, position543)
													}
													break
												default:
													{
														position551 := position
														{
															position552, tokenIndex552 := position, tokenIndex
															{
																position554 := position
																{
																	position555, tokenIndex555 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l556
																	}
																	position++
																	goto l555
																l556:
																	position, tokenIndex = position555, tokenIndex555
																	if buffer[position] != rune('L') {
																		goto l553
																	}
																	position++
																}
															l555:
																{
																	position557, tokenIndex557 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l558
																	}
																	position++
																	goto l557
																l558:
																	position, tokenIndex = position557, tokenIndex557
																	if buffer[position] != rune('D') {
																		goto l553
																	}
																	position++
																}
															l557:
																if !_rules[rulews]() {
																	goto l553
																}
																if !_rules[ruleDst16]() {
																	goto l553
																}
																if !_rules[rulesep]() {
																	goto l553
																}
																if !_rules[ruleSrc16]() {
																	goto l553
																}
																{
																	add(ruleAction2, position)
																}
																add(ruleLoad16, position554)
															}
															goto l552
														l553:
															position, tokenIndex = position552, tokenIndex552
															{
																position560 := position
																{
																	position561, tokenIndex561 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l562
																	}
																	position++
																	goto l561
																l562:
																	position, tokenIndex = position561, tokenIndex561
																	if buffer[position] != rune('L') {
																		goto l522
																	}
																	position++
																}
															l561:
																{
																	position563, tokenIndex563 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l564
																	}
																	position++
																	goto l563
																l564:
																	position, tokenIndex = position563, tokenIndex563
																	if buffer[position] != rune('D') {
																		goto l522
																	}
																	position++
																}
															l563:
																if !_rules[rulews]() {
																	goto l522
																}
																{
																	position565 := position
																	{
																		position566, tokenIndex566 := position, tokenIndex
																		if !_rules[ruleReg8]() {
																			goto l567
																		}
																		goto l566
																	l567:
																		position, tokenIndex = position566, tokenIndex566
																		if !_rules[ruleReg16Contents]() {
																			goto l568
																		}
																		goto l566
																	l568:
																		position, tokenIndex = position566, tokenIndex566
																		if !_rules[rulenn_contents]() {
																			goto l522
																		}
																	}
																l566:
																	{
																		add(ruleAction13, position)
																	}
																	add(ruleDst8, position565)
																}
																if !_rules[rulesep]() {
																	goto l522
																}
																if !_rules[ruleSrc8]() {
																	goto l522
																}
																{
																	add(ruleAction1, position)
																}
																add(ruleLoad8, position560)
															}
														}
													l552:
														add(ruleLoad, position551)
													}
													break
												}
											}

										}
									l524:
										add(ruleAssignment, position523)
									}
									goto l521
								l522:
									position, tokenIndex = position521, tokenIndex521
									{
										position572 := position
										{
											position573, tokenIndex573 := position, tokenIndex
											{
												position575 := position
												{
													position576, tokenIndex576 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l577
													}
													position++
													goto l576
												l577:
													position, tokenIndex = position576, tokenIndex576
													if buffer[position] != rune('I') {
														goto l574
													}
													position++
												}
											l576:
												{
													position578, tokenIndex578 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l579
													}
													position++
													goto l578
												l579:
													position, tokenIndex = position578, tokenIndex578
													if buffer[position] != rune('N') {
														goto l574
													}
													position++
												}
											l578:
												{
													position580, tokenIndex580 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l581
													}
													position++
													goto l580
												l581:
													position, tokenIndex = position580, tokenIndex580
													if buffer[position] != rune('C') {
														goto l574
													}
													position++
												}
											l580:
												if !_rules[rulews]() {
													goto l574
												}
												if !_rules[ruleILoc8]() {
													goto l574
												}
												{
													add(ruleAction6, position)
												}
												add(ruleInc16Indexed8, position575)
											}
											goto l573
										l574:
											position, tokenIndex = position573, tokenIndex573
											{
												position584 := position
												{
													position585, tokenIndex585 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l586
													}
													position++
													goto l585
												l586:
													position, tokenIndex = position585, tokenIndex585
													if buffer[position] != rune('I') {
														goto l583
													}
													position++
												}
											l585:
												{
													position587, tokenIndex587 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l588
													}
													position++
													goto l587
												l588:
													position, tokenIndex = position587, tokenIndex587
													if buffer[position] != rune('N') {
														goto l583
													}
													position++
												}
											l587:
												{
													position589, tokenIndex589 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l590
													}
													position++
													goto l589
												l590:
													position, tokenIndex = position589, tokenIndex589
													if buffer[position] != rune('C') {
														goto l583
													}
													position++
												}
											l589:
												if !_rules[rulews]() {
													goto l583
												}
												if !_rules[ruleLoc16]() {
													goto l583
												}
												{
													add(ruleAction8, position)
												}
												add(ruleInc16, position584)
											}
											goto l573
										l583:
											position, tokenIndex = position573, tokenIndex573
											{
												position592 := position
												{
													position593, tokenIndex593 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l594
													}
													position++
													goto l593
												l594:
													position, tokenIndex = position593, tokenIndex593
													if buffer[position] != rune('I') {
														goto l571
													}
													position++
												}
											l593:
												{
													position595, tokenIndex595 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l596
													}
													position++
													goto l595
												l596:
													position, tokenIndex = position595, tokenIndex595
													if buffer[position] != rune('N') {
														goto l571
													}
													position++
												}
											l595:
												{
													position597, tokenIndex597 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l598
													}
													position++
													goto l597
												l598:
													position, tokenIndex = position597, tokenIndex597
													if buffer[position] != rune('C') {
														goto l571
													}
													position++
												}
											l597:
												if !_rules[rulews]() {
													goto l571
												}
												if !_rules[ruleLoc8]() {
													goto l571
												}
												{
													add(ruleAction7, position)
												}
												add(ruleInc8, position592)
											}
										}
									l573:
										add(ruleInc, position572)
									}
									goto l521
								l571:
									position, tokenIndex = position521, tokenIndex521
									{
										position601 := position
										{
											position602, tokenIndex602 := position, tokenIndex
											{
												position604 := position
												{
													position605, tokenIndex605 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l606
													}
													position++
													goto l605
												l606:
													position, tokenIndex = position605, tokenIndex605
													if buffer[position] != rune('D') {
														goto l603
													}
													position++
												}
											l605:
												{
													position607, tokenIndex607 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l608
													}
													position++
													goto l607
												l608:
													position, tokenIndex = position607, tokenIndex607
													if buffer[position] != rune('E') {
														goto l603
													}
													position++
												}
											l607:
												{
													position609, tokenIndex609 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l610
													}
													position++
													goto l609
												l610:
													position, tokenIndex = position609, tokenIndex609
													if buffer[position] != rune('C') {
														goto l603
													}
													position++
												}
											l609:
												if !_rules[rulews]() {
													goto l603
												}
												if !_rules[ruleILoc8]() {
													goto l603
												}
												{
													add(ruleAction9, position)
												}
												add(ruleDec16Indexed8, position604)
											}
											goto l602
										l603:
											position, tokenIndex = position602, tokenIndex602
											{
												position613 := position
												{
													position614, tokenIndex614 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l615
													}
													position++
													goto l614
												l615:
													position, tokenIndex = position614, tokenIndex614
													if buffer[position] != rune('D') {
														goto l612
													}
													position++
												}
											l614:
												{
													position616, tokenIndex616 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l617
													}
													position++
													goto l616
												l617:
													position, tokenIndex = position616, tokenIndex616
													if buffer[position] != rune('E') {
														goto l612
													}
													position++
												}
											l616:
												{
													position618, tokenIndex618 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l619
													}
													position++
													goto l618
												l619:
													position, tokenIndex = position618, tokenIndex618
													if buffer[position] != rune('C') {
														goto l612
													}
													position++
												}
											l618:
												if !_rules[rulews]() {
													goto l612
												}
												if !_rules[ruleLoc16]() {
													goto l612
												}
												{
													add(ruleAction11, position)
												}
												add(ruleDec16, position613)
											}
											goto l602
										l612:
											position, tokenIndex = position602, tokenIndex602
											{
												position621 := position
												{
													position622, tokenIndex622 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l623
													}
													position++
													goto l622
												l623:
													position, tokenIndex = position622, tokenIndex622
													if buffer[position] != rune('D') {
														goto l600
													}
													position++
												}
											l622:
												{
													position624, tokenIndex624 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l625
													}
													position++
													goto l624
												l625:
													position, tokenIndex = position624, tokenIndex624
													if buffer[position] != rune('E') {
														goto l600
													}
													position++
												}
											l624:
												{
													position626, tokenIndex626 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l627
													}
													position++
													goto l626
												l627:
													position, tokenIndex = position626, tokenIndex626
													if buffer[position] != rune('C') {
														goto l600
													}
													position++
												}
											l626:
												if !_rules[rulews]() {
													goto l600
												}
												if !_rules[ruleLoc8]() {
													goto l600
												}
												{
													add(ruleAction10, position)
												}
												add(ruleDec8, position621)
											}
										}
									l602:
										add(ruleDec, position601)
									}
									goto l521
								l600:
									position, tokenIndex = position521, tokenIndex521
									{
										position630 := position
										{
											position631, tokenIndex631 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l632
											}
											position++
											goto l631
										l632:
											position, tokenIndex = position631, tokenIndex631
											if buffer[position] != rune('A') {
												goto l629
											}
											position++
										}
									l631:
										{
											position633, tokenIndex633 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l634
											}
											position++
											goto l633
										l634:
											position, tokenIndex = position633, tokenIndex633
											if buffer[position] != rune('D') {
												goto l629
											}
											position++
										}
									l633:
										{
											position635, tokenIndex635 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l636
											}
											position++
											goto l635
										l636:
											position, tokenIndex = position635, tokenIndex635
											if buffer[position] != rune('D') {
												goto l629
											}
											position++
										}
									l635:
										if !_rules[rulews]() {
											goto l629
										}
										if !_rules[ruleDst16]() {
											goto l629
										}
										if !_rules[rulesep]() {
											goto l629
										}
										if !_rules[ruleSrc16]() {
											goto l629
										}
										{
											add(ruleAction12, position)
										}
										add(ruleAdd16, position630)
									}
									goto l521
								l629:
									position, tokenIndex = position521, tokenIndex521
									{
										position639 := position
										{
											position640, tokenIndex640 := position, tokenIndex
											{
												position642 := position
												{
													position643, tokenIndex643 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l644
													}
													position++
													goto l643
												l644:
													position, tokenIndex = position643, tokenIndex643
													if buffer[position] != rune('A') {
														goto l641
													}
													position++
												}
											l643:
												{
													position645, tokenIndex645 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l646
													}
													position++
													goto l645
												l646:
													position, tokenIndex = position645, tokenIndex645
													if buffer[position] != rune('D') {
														goto l641
													}
													position++
												}
											l645:
												{
													position647, tokenIndex647 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l648
													}
													position++
													goto l647
												l648:
													position, tokenIndex = position647, tokenIndex647
													if buffer[position] != rune('D') {
														goto l641
													}
													position++
												}
											l647:
												if !_rules[rulews]() {
													goto l641
												}
												{
													position649, tokenIndex649 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l650
													}
													position++
													goto l649
												l650:
													position, tokenIndex = position649, tokenIndex649
													if buffer[position] != rune('A') {
														goto l641
													}
													position++
												}
											l649:
												if !_rules[rulesep]() {
													goto l641
												}
												if !_rules[ruleSrc8]() {
													goto l641
												}
												{
													add(ruleAction27, position)
												}
												add(ruleAdd, position642)
											}
											goto l640
										l641:
											position, tokenIndex = position640, tokenIndex640
											{
												position653 := position
												{
													position654, tokenIndex654 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l655
													}
													position++
													goto l654
												l655:
													position, tokenIndex = position654, tokenIndex654
													if buffer[position] != rune('A') {
														goto l652
													}
													position++
												}
											l654:
												{
													position656, tokenIndex656 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l657
													}
													position++
													goto l656
												l657:
													position, tokenIndex = position656, tokenIndex656
													if buffer[position] != rune('D') {
														goto l652
													}
													position++
												}
											l656:
												{
													position658, tokenIndex658 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l659
													}
													position++
													goto l658
												l659:
													position, tokenIndex = position658, tokenIndex658
													if buffer[position] != rune('C') {
														goto l652
													}
													position++
												}
											l658:
												if !_rules[rulews]() {
													goto l652
												}
												{
													position660, tokenIndex660 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l661
													}
													position++
													goto l660
												l661:
													position, tokenIndex = position660, tokenIndex660
													if buffer[position] != rune('A') {
														goto l652
													}
													position++
												}
											l660:
												if !_rules[rulesep]() {
													goto l652
												}
												if !_rules[ruleSrc8]() {
													goto l652
												}
												{
													add(ruleAction28, position)
												}
												add(ruleAdc, position653)
											}
											goto l640
										l652:
											position, tokenIndex = position640, tokenIndex640
											{
												position664 := position
												{
													position665, tokenIndex665 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l666
													}
													position++
													goto l665
												l666:
													position, tokenIndex = position665, tokenIndex665
													if buffer[position] != rune('S') {
														goto l663
													}
													position++
												}
											l665:
												{
													position667, tokenIndex667 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l668
													}
													position++
													goto l667
												l668:
													position, tokenIndex = position667, tokenIndex667
													if buffer[position] != rune('U') {
														goto l663
													}
													position++
												}
											l667:
												{
													position669, tokenIndex669 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l670
													}
													position++
													goto l669
												l670:
													position, tokenIndex = position669, tokenIndex669
													if buffer[position] != rune('B') {
														goto l663
													}
													position++
												}
											l669:
												if !_rules[rulews]() {
													goto l663
												}
												if !_rules[ruleSrc8]() {
													goto l663
												}
												{
													add(ruleAction29, position)
												}
												add(ruleSub, position664)
											}
											goto l640
										l663:
											position, tokenIndex = position640, tokenIndex640
											{
												switch buffer[position] {
												case 'C', 'c':
													{
														position673 := position
														{
															position674, tokenIndex674 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l675
															}
															position++
															goto l674
														l675:
															position, tokenIndex = position674, tokenIndex674
															if buffer[position] != rune('C') {
																goto l638
															}
															position++
														}
													l674:
														{
															position676, tokenIndex676 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l677
															}
															position++
															goto l676
														l677:
															position, tokenIndex = position676, tokenIndex676
															if buffer[position] != rune('P') {
																goto l638
															}
															position++
														}
													l676:
														if !_rules[rulews]() {
															goto l638
														}
														if !_rules[ruleSrc8]() {
															goto l638
														}
														{
															add(ruleAction34, position)
														}
														add(ruleCp, position673)
													}
													break
												case 'O', 'o':
													{
														position679 := position
														{
															position680, tokenIndex680 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l681
															}
															position++
															goto l680
														l681:
															position, tokenIndex = position680, tokenIndex680
															if buffer[position] != rune('O') {
																goto l638
															}
															position++
														}
													l680:
														{
															position682, tokenIndex682 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l683
															}
															position++
															goto l682
														l683:
															position, tokenIndex = position682, tokenIndex682
															if buffer[position] != rune('R') {
																goto l638
															}
															position++
														}
													l682:
														if !_rules[rulews]() {
															goto l638
														}
														if !_rules[ruleSrc8]() {
															goto l638
														}
														{
															add(ruleAction33, position)
														}
														add(ruleOr, position679)
													}
													break
												case 'X', 'x':
													{
														position685 := position
														{
															position686, tokenIndex686 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l687
															}
															position++
															goto l686
														l687:
															position, tokenIndex = position686, tokenIndex686
															if buffer[position] != rune('X') {
																goto l638
															}
															position++
														}
													l686:
														{
															position688, tokenIndex688 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l689
															}
															position++
															goto l688
														l689:
															position, tokenIndex = position688, tokenIndex688
															if buffer[position] != rune('O') {
																goto l638
															}
															position++
														}
													l688:
														{
															position690, tokenIndex690 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l691
															}
															position++
															goto l690
														l691:
															position, tokenIndex = position690, tokenIndex690
															if buffer[position] != rune('R') {
																goto l638
															}
															position++
														}
													l690:
														if !_rules[rulews]() {
															goto l638
														}
														if !_rules[ruleSrc8]() {
															goto l638
														}
														{
															add(ruleAction32, position)
														}
														add(ruleXor, position685)
													}
													break
												case 'A', 'a':
													{
														position693 := position
														{
															position694, tokenIndex694 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l695
															}
															position++
															goto l694
														l695:
															position, tokenIndex = position694, tokenIndex694
															if buffer[position] != rune('A') {
																goto l638
															}
															position++
														}
													l694:
														{
															position696, tokenIndex696 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l697
															}
															position++
															goto l696
														l697:
															position, tokenIndex = position696, tokenIndex696
															if buffer[position] != rune('N') {
																goto l638
															}
															position++
														}
													l696:
														{
															position698, tokenIndex698 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l699
															}
															position++
															goto l698
														l699:
															position, tokenIndex = position698, tokenIndex698
															if buffer[position] != rune('D') {
																goto l638
															}
															position++
														}
													l698:
														if !_rules[rulews]() {
															goto l638
														}
														if !_rules[ruleSrc8]() {
															goto l638
														}
														{
															add(ruleAction31, position)
														}
														add(ruleAnd, position693)
													}
													break
												default:
													{
														position701 := position
														{
															position702, tokenIndex702 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l703
															}
															position++
															goto l702
														l703:
															position, tokenIndex = position702, tokenIndex702
															if buffer[position] != rune('S') {
																goto l638
															}
															position++
														}
													l702:
														{
															position704, tokenIndex704 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l705
															}
															position++
															goto l704
														l705:
															position, tokenIndex = position704, tokenIndex704
															if buffer[position] != rune('B') {
																goto l638
															}
															position++
														}
													l704:
														{
															position706, tokenIndex706 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l707
															}
															position++
															goto l706
														l707:
															position, tokenIndex = position706, tokenIndex706
															if buffer[position] != rune('C') {
																goto l638
															}
															position++
														}
													l706:
														if !_rules[rulews]() {
															goto l638
														}
														{
															position708, tokenIndex708 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l709
															}
															position++
															goto l708
														l709:
															position, tokenIndex = position708, tokenIndex708
															if buffer[position] != rune('A') {
																goto l638
															}
															position++
														}
													l708:
														if !_rules[rulesep]() {
															goto l638
														}
														if !_rules[ruleSrc8]() {
															goto l638
														}
														{
															add(ruleAction30, position)
														}
														add(ruleSbc, position701)
													}
													break
												}
											}

										}
									l640:
										add(ruleAlu, position639)
									}
									goto l521
								l638:
									position, tokenIndex = position521, tokenIndex521
									{
										position712 := position
										{
											position713, tokenIndex713 := position, tokenIndex
											{
												position715 := position
												{
													position716, tokenIndex716 := position, tokenIndex
													{
														position718 := position
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
																goto l717
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
																goto l717
															}
															position++
														}
													l721:
														{
															position723, tokenIndex723 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l724
															}
															position++
															goto l723
														l724:
															position, tokenIndex = position723, tokenIndex723
															if buffer[position] != rune('C') {
																goto l717
															}
															position++
														}
													l723:
														if !_rules[rulews]() {
															goto l717
														}
														if !_rules[ruleLoc8]() {
															goto l717
														}
														{
															add(ruleAction35, position)
														}
														add(ruleRlc, position718)
													}
													goto l716
												l717:
													position, tokenIndex = position716, tokenIndex716
													{
														position727 := position
														{
															position728, tokenIndex728 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l729
															}
															position++
															goto l728
														l729:
															position, tokenIndex = position728, tokenIndex728
															if buffer[position] != rune('R') {
																goto l726
															}
															position++
														}
													l728:
														{
															position730, tokenIndex730 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l731
															}
															position++
															goto l730
														l731:
															position, tokenIndex = position730, tokenIndex730
															if buffer[position] != rune('R') {
																goto l726
															}
															position++
														}
													l730:
														{
															position732, tokenIndex732 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l733
															}
															position++
															goto l732
														l733:
															position, tokenIndex = position732, tokenIndex732
															if buffer[position] != rune('C') {
																goto l726
															}
															position++
														}
													l732:
														if !_rules[rulews]() {
															goto l726
														}
														if !_rules[ruleLoc8]() {
															goto l726
														}
														{
															add(ruleAction36, position)
														}
														add(ruleRrc, position727)
													}
													goto l716
												l726:
													position, tokenIndex = position716, tokenIndex716
													{
														position736 := position
														{
															position737, tokenIndex737 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l738
															}
															position++
															goto l737
														l738:
															position, tokenIndex = position737, tokenIndex737
															if buffer[position] != rune('R') {
																goto l735
															}
															position++
														}
													l737:
														{
															position739, tokenIndex739 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l740
															}
															position++
															goto l739
														l740:
															position, tokenIndex = position739, tokenIndex739
															if buffer[position] != rune('L') {
																goto l735
															}
															position++
														}
													l739:
														if !_rules[rulews]() {
															goto l735
														}
														if !_rules[ruleLoc8]() {
															goto l735
														}
														{
															add(ruleAction37, position)
														}
														add(ruleRl, position736)
													}
													goto l716
												l735:
													position, tokenIndex = position716, tokenIndex716
													{
														position743 := position
														{
															position744, tokenIndex744 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l745
															}
															position++
															goto l744
														l745:
															position, tokenIndex = position744, tokenIndex744
															if buffer[position] != rune('R') {
																goto l742
															}
															position++
														}
													l744:
														{
															position746, tokenIndex746 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l747
															}
															position++
															goto l746
														l747:
															position, tokenIndex = position746, tokenIndex746
															if buffer[position] != rune('R') {
																goto l742
															}
															position++
														}
													l746:
														if !_rules[rulews]() {
															goto l742
														}
														if !_rules[ruleLoc8]() {
															goto l742
														}
														{
															add(ruleAction38, position)
														}
														add(ruleRr, position743)
													}
													goto l716
												l742:
													position, tokenIndex = position716, tokenIndex716
													{
														position750 := position
														{
															position751, tokenIndex751 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l752
															}
															position++
															goto l751
														l752:
															position, tokenIndex = position751, tokenIndex751
															if buffer[position] != rune('S') {
																goto l749
															}
															position++
														}
													l751:
														{
															position753, tokenIndex753 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l754
															}
															position++
															goto l753
														l754:
															position, tokenIndex = position753, tokenIndex753
															if buffer[position] != rune('L') {
																goto l749
															}
															position++
														}
													l753:
														{
															position755, tokenIndex755 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l756
															}
															position++
															goto l755
														l756:
															position, tokenIndex = position755, tokenIndex755
															if buffer[position] != rune('A') {
																goto l749
															}
															position++
														}
													l755:
														if !_rules[rulews]() {
															goto l749
														}
														if !_rules[ruleLoc8]() {
															goto l749
														}
														{
															add(ruleAction39, position)
														}
														add(ruleSla, position750)
													}
													goto l716
												l749:
													position, tokenIndex = position716, tokenIndex716
													{
														position759 := position
														{
															position760, tokenIndex760 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l761
															}
															position++
															goto l760
														l761:
															position, tokenIndex = position760, tokenIndex760
															if buffer[position] != rune('S') {
																goto l758
															}
															position++
														}
													l760:
														{
															position762, tokenIndex762 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l763
															}
															position++
															goto l762
														l763:
															position, tokenIndex = position762, tokenIndex762
															if buffer[position] != rune('R') {
																goto l758
															}
															position++
														}
													l762:
														{
															position764, tokenIndex764 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l765
															}
															position++
															goto l764
														l765:
															position, tokenIndex = position764, tokenIndex764
															if buffer[position] != rune('A') {
																goto l758
															}
															position++
														}
													l764:
														if !_rules[rulews]() {
															goto l758
														}
														if !_rules[ruleLoc8]() {
															goto l758
														}
														{
															add(ruleAction40, position)
														}
														add(ruleSra, position759)
													}
													goto l716
												l758:
													position, tokenIndex = position716, tokenIndex716
													{
														position768 := position
														{
															position769, tokenIndex769 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l770
															}
															position++
															goto l769
														l770:
															position, tokenIndex = position769, tokenIndex769
															if buffer[position] != rune('S') {
																goto l767
															}
															position++
														}
													l769:
														{
															position771, tokenIndex771 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l772
															}
															position++
															goto l771
														l772:
															position, tokenIndex = position771, tokenIndex771
															if buffer[position] != rune('L') {
																goto l767
															}
															position++
														}
													l771:
														{
															position773, tokenIndex773 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l774
															}
															position++
															goto l773
														l774:
															position, tokenIndex = position773, tokenIndex773
															if buffer[position] != rune('L') {
																goto l767
															}
															position++
														}
													l773:
														if !_rules[rulews]() {
															goto l767
														}
														if !_rules[ruleLoc8]() {
															goto l767
														}
														{
															add(ruleAction41, position)
														}
														add(ruleSll, position768)
													}
													goto l716
												l767:
													position, tokenIndex = position716, tokenIndex716
													{
														position776 := position
														{
															position777, tokenIndex777 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l778
															}
															position++
															goto l777
														l778:
															position, tokenIndex = position777, tokenIndex777
															if buffer[position] != rune('S') {
																goto l714
															}
															position++
														}
													l777:
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
																goto l714
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
																goto l714
															}
															position++
														}
													l781:
														if !_rules[rulews]() {
															goto l714
														}
														if !_rules[ruleLoc8]() {
															goto l714
														}
														{
															add(ruleAction42, position)
														}
														add(ruleSrl, position776)
													}
												}
											l716:
												add(ruleRot, position715)
											}
											goto l713
										l714:
											position, tokenIndex = position713, tokenIndex713
											{
												switch buffer[position] {
												case 'S', 's':
													{
														position785 := position
														{
															position786, tokenIndex786 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l787
															}
															position++
															goto l786
														l787:
															position, tokenIndex = position786, tokenIndex786
															if buffer[position] != rune('S') {
																goto l711
															}
															position++
														}
													l786:
														{
															position788, tokenIndex788 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l789
															}
															position++
															goto l788
														l789:
															position, tokenIndex = position788, tokenIndex788
															if buffer[position] != rune('E') {
																goto l711
															}
															position++
														}
													l788:
														{
															position790, tokenIndex790 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l791
															}
															position++
															goto l790
														l791:
															position, tokenIndex = position790, tokenIndex790
															if buffer[position] != rune('T') {
																goto l711
															}
															position++
														}
													l790:
														if !_rules[rulews]() {
															goto l711
														}
														if !_rules[ruleoctaldigit]() {
															goto l711
														}
														if !_rules[rulesep]() {
															goto l711
														}
														if !_rules[ruleLoc8]() {
															goto l711
														}
														{
															add(ruleAction45, position)
														}
														add(ruleSet, position785)
													}
													break
												case 'R', 'r':
													{
														position793 := position
														{
															position794, tokenIndex794 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l795
															}
															position++
															goto l794
														l795:
															position, tokenIndex = position794, tokenIndex794
															if buffer[position] != rune('R') {
																goto l711
															}
															position++
														}
													l794:
														{
															position796, tokenIndex796 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l797
															}
															position++
															goto l796
														l797:
															position, tokenIndex = position796, tokenIndex796
															if buffer[position] != rune('E') {
																goto l711
															}
															position++
														}
													l796:
														{
															position798, tokenIndex798 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l799
															}
															position++
															goto l798
														l799:
															position, tokenIndex = position798, tokenIndex798
															if buffer[position] != rune('S') {
																goto l711
															}
															position++
														}
													l798:
														if !_rules[rulews]() {
															goto l711
														}
														if !_rules[ruleoctaldigit]() {
															goto l711
														}
														if !_rules[rulesep]() {
															goto l711
														}
														if !_rules[ruleLoc8]() {
															goto l711
														}
														{
															add(ruleAction44, position)
														}
														add(ruleRes, position793)
													}
													break
												default:
													{
														position801 := position
														{
															position802, tokenIndex802 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l803
															}
															position++
															goto l802
														l803:
															position, tokenIndex = position802, tokenIndex802
															if buffer[position] != rune('B') {
																goto l711
															}
															position++
														}
													l802:
														{
															position804, tokenIndex804 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l805
															}
															position++
															goto l804
														l805:
															position, tokenIndex = position804, tokenIndex804
															if buffer[position] != rune('I') {
																goto l711
															}
															position++
														}
													l804:
														{
															position806, tokenIndex806 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l807
															}
															position++
															goto l806
														l807:
															position, tokenIndex = position806, tokenIndex806
															if buffer[position] != rune('T') {
																goto l711
															}
															position++
														}
													l806:
														if !_rules[rulews]() {
															goto l711
														}
														if !_rules[ruleoctaldigit]() {
															goto l711
														}
														if !_rules[rulesep]() {
															goto l711
														}
														if !_rules[ruleLoc8]() {
															goto l711
														}
														{
															add(ruleAction43, position)
														}
														add(ruleBit, position801)
													}
													break
												}
											}

										}
									l713:
										add(ruleBitOp, position712)
									}
									goto l521
								l711:
									position, tokenIndex = position521, tokenIndex521
									{
										position810 := position
										{
											position811, tokenIndex811 := position, tokenIndex
											{
												position813 := position
												{
													position814 := position
													{
														position815, tokenIndex815 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l816
														}
														position++
														goto l815
													l816:
														position, tokenIndex = position815, tokenIndex815
														if buffer[position] != rune('R') {
															goto l812
														}
														position++
													}
												l815:
													{
														position817, tokenIndex817 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l818
														}
														position++
														goto l817
													l818:
														position, tokenIndex = position817, tokenIndex817
														if buffer[position] != rune('L') {
															goto l812
														}
														position++
													}
												l817:
													{
														position819, tokenIndex819 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l820
														}
														position++
														goto l819
													l820:
														position, tokenIndex = position819, tokenIndex819
														if buffer[position] != rune('C') {
															goto l812
														}
														position++
													}
												l819:
													{
														position821, tokenIndex821 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l822
														}
														position++
														goto l821
													l822:
														position, tokenIndex = position821, tokenIndex821
														if buffer[position] != rune('A') {
															goto l812
														}
														position++
													}
												l821:
													add(rulePegText, position814)
												}
												{
													add(ruleAction48, position)
												}
												add(ruleRlca, position813)
											}
											goto l811
										l812:
											position, tokenIndex = position811, tokenIndex811
											{
												position825 := position
												{
													position826 := position
													{
														position827, tokenIndex827 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l828
														}
														position++
														goto l827
													l828:
														position, tokenIndex = position827, tokenIndex827
														if buffer[position] != rune('R') {
															goto l824
														}
														position++
													}
												l827:
													{
														position829, tokenIndex829 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l830
														}
														position++
														goto l829
													l830:
														position, tokenIndex = position829, tokenIndex829
														if buffer[position] != rune('R') {
															goto l824
														}
														position++
													}
												l829:
													{
														position831, tokenIndex831 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l832
														}
														position++
														goto l831
													l832:
														position, tokenIndex = position831, tokenIndex831
														if buffer[position] != rune('C') {
															goto l824
														}
														position++
													}
												l831:
													{
														position833, tokenIndex833 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l834
														}
														position++
														goto l833
													l834:
														position, tokenIndex = position833, tokenIndex833
														if buffer[position] != rune('A') {
															goto l824
														}
														position++
													}
												l833:
													add(rulePegText, position826)
												}
												{
													add(ruleAction49, position)
												}
												add(ruleRrca, position825)
											}
											goto l811
										l824:
											position, tokenIndex = position811, tokenIndex811
											{
												position837 := position
												{
													position838 := position
													{
														position839, tokenIndex839 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l840
														}
														position++
														goto l839
													l840:
														position, tokenIndex = position839, tokenIndex839
														if buffer[position] != rune('R') {
															goto l836
														}
														position++
													}
												l839:
													{
														position841, tokenIndex841 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l842
														}
														position++
														goto l841
													l842:
														position, tokenIndex = position841, tokenIndex841
														if buffer[position] != rune('L') {
															goto l836
														}
														position++
													}
												l841:
													{
														position843, tokenIndex843 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l844
														}
														position++
														goto l843
													l844:
														position, tokenIndex = position843, tokenIndex843
														if buffer[position] != rune('A') {
															goto l836
														}
														position++
													}
												l843:
													add(rulePegText, position838)
												}
												{
													add(ruleAction50, position)
												}
												add(ruleRla, position837)
											}
											goto l811
										l836:
											position, tokenIndex = position811, tokenIndex811
											{
												position847 := position
												{
													position848 := position
													{
														position849, tokenIndex849 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l850
														}
														position++
														goto l849
													l850:
														position, tokenIndex = position849, tokenIndex849
														if buffer[position] != rune('D') {
															goto l846
														}
														position++
													}
												l849:
													{
														position851, tokenIndex851 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l852
														}
														position++
														goto l851
													l852:
														position, tokenIndex = position851, tokenIndex851
														if buffer[position] != rune('A') {
															goto l846
														}
														position++
													}
												l851:
													{
														position853, tokenIndex853 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l854
														}
														position++
														goto l853
													l854:
														position, tokenIndex = position853, tokenIndex853
														if buffer[position] != rune('A') {
															goto l846
														}
														position++
													}
												l853:
													add(rulePegText, position848)
												}
												{
													add(ruleAction52, position)
												}
												add(ruleDaa, position847)
											}
											goto l811
										l846:
											position, tokenIndex = position811, tokenIndex811
											{
												position857 := position
												{
													position858 := position
													{
														position859, tokenIndex859 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l860
														}
														position++
														goto l859
													l860:
														position, tokenIndex = position859, tokenIndex859
														if buffer[position] != rune('C') {
															goto l856
														}
														position++
													}
												l859:
													{
														position861, tokenIndex861 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l862
														}
														position++
														goto l861
													l862:
														position, tokenIndex = position861, tokenIndex861
														if buffer[position] != rune('P') {
															goto l856
														}
														position++
													}
												l861:
													{
														position863, tokenIndex863 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l864
														}
														position++
														goto l863
													l864:
														position, tokenIndex = position863, tokenIndex863
														if buffer[position] != rune('L') {
															goto l856
														}
														position++
													}
												l863:
													add(rulePegText, position858)
												}
												{
													add(ruleAction53, position)
												}
												add(ruleCpl, position857)
											}
											goto l811
										l856:
											position, tokenIndex = position811, tokenIndex811
											{
												position867 := position
												{
													position868 := position
													{
														position869, tokenIndex869 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l870
														}
														position++
														goto l869
													l870:
														position, tokenIndex = position869, tokenIndex869
														if buffer[position] != rune('E') {
															goto l866
														}
														position++
													}
												l869:
													{
														position871, tokenIndex871 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l872
														}
														position++
														goto l871
													l872:
														position, tokenIndex = position871, tokenIndex871
														if buffer[position] != rune('X') {
															goto l866
														}
														position++
													}
												l871:
													{
														position873, tokenIndex873 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l874
														}
														position++
														goto l873
													l874:
														position, tokenIndex = position873, tokenIndex873
														if buffer[position] != rune('X') {
															goto l866
														}
														position++
													}
												l873:
													add(rulePegText, position868)
												}
												{
													add(ruleAction56, position)
												}
												add(ruleExx, position867)
											}
											goto l811
										l866:
											position, tokenIndex = position811, tokenIndex811
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position877 := position
														{
															position878 := position
															{
																position879, tokenIndex879 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l880
																}
																position++
																goto l879
															l880:
																position, tokenIndex = position879, tokenIndex879
																if buffer[position] != rune('E') {
																	goto l809
																}
																position++
															}
														l879:
															{
																position881, tokenIndex881 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l882
																}
																position++
																goto l881
															l882:
																position, tokenIndex = position881, tokenIndex881
																if buffer[position] != rune('I') {
																	goto l809
																}
																position++
															}
														l881:
															add(rulePegText, position878)
														}
														{
															add(ruleAction58, position)
														}
														add(ruleEi, position877)
													}
													break
												case 'D', 'd':
													{
														position884 := position
														{
															position885 := position
															{
																position886, tokenIndex886 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l887
																}
																position++
																goto l886
															l887:
																position, tokenIndex = position886, tokenIndex886
																if buffer[position] != rune('D') {
																	goto l809
																}
																position++
															}
														l886:
															{
																position888, tokenIndex888 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l889
																}
																position++
																goto l888
															l889:
																position, tokenIndex = position888, tokenIndex888
																if buffer[position] != rune('I') {
																	goto l809
																}
																position++
															}
														l888:
															add(rulePegText, position885)
														}
														{
															add(ruleAction57, position)
														}
														add(ruleDi, position884)
													}
													break
												case 'C', 'c':
													{
														position891 := position
														{
															position892 := position
															{
																position893, tokenIndex893 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l894
																}
																position++
																goto l893
															l894:
																position, tokenIndex = position893, tokenIndex893
																if buffer[position] != rune('C') {
																	goto l809
																}
																position++
															}
														l893:
															{
																position895, tokenIndex895 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l896
																}
																position++
																goto l895
															l896:
																position, tokenIndex = position895, tokenIndex895
																if buffer[position] != rune('C') {
																	goto l809
																}
																position++
															}
														l895:
															{
																position897, tokenIndex897 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l898
																}
																position++
																goto l897
															l898:
																position, tokenIndex = position897, tokenIndex897
																if buffer[position] != rune('F') {
																	goto l809
																}
																position++
															}
														l897:
															add(rulePegText, position892)
														}
														{
															add(ruleAction55, position)
														}
														add(ruleCcf, position891)
													}
													break
												case 'S', 's':
													{
														position900 := position
														{
															position901 := position
															{
																position902, tokenIndex902 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l903
																}
																position++
																goto l902
															l903:
																position, tokenIndex = position902, tokenIndex902
																if buffer[position] != rune('S') {
																	goto l809
																}
																position++
															}
														l902:
															{
																position904, tokenIndex904 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l905
																}
																position++
																goto l904
															l905:
																position, tokenIndex = position904, tokenIndex904
																if buffer[position] != rune('C') {
																	goto l809
																}
																position++
															}
														l904:
															{
																position906, tokenIndex906 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l907
																}
																position++
																goto l906
															l907:
																position, tokenIndex = position906, tokenIndex906
																if buffer[position] != rune('F') {
																	goto l809
																}
																position++
															}
														l906:
															add(rulePegText, position901)
														}
														{
															add(ruleAction54, position)
														}
														add(ruleScf, position900)
													}
													break
												case 'R', 'r':
													{
														position909 := position
														{
															position910 := position
															{
																position911, tokenIndex911 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l912
																}
																position++
																goto l911
															l912:
																position, tokenIndex = position911, tokenIndex911
																if buffer[position] != rune('R') {
																	goto l809
																}
																position++
															}
														l911:
															{
																position913, tokenIndex913 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l914
																}
																position++
																goto l913
															l914:
																position, tokenIndex = position913, tokenIndex913
																if buffer[position] != rune('R') {
																	goto l809
																}
																position++
															}
														l913:
															{
																position915, tokenIndex915 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l916
																}
																position++
																goto l915
															l916:
																position, tokenIndex = position915, tokenIndex915
																if buffer[position] != rune('A') {
																	goto l809
																}
																position++
															}
														l915:
															add(rulePegText, position910)
														}
														{
															add(ruleAction51, position)
														}
														add(ruleRra, position909)
													}
													break
												case 'H', 'h':
													{
														position918 := position
														{
															position919 := position
															{
																position920, tokenIndex920 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l921
																}
																position++
																goto l920
															l921:
																position, tokenIndex = position920, tokenIndex920
																if buffer[position] != rune('H') {
																	goto l809
																}
																position++
															}
														l920:
															{
																position922, tokenIndex922 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l923
																}
																position++
																goto l922
															l923:
																position, tokenIndex = position922, tokenIndex922
																if buffer[position] != rune('A') {
																	goto l809
																}
																position++
															}
														l922:
															{
																position924, tokenIndex924 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l925
																}
																position++
																goto l924
															l925:
																position, tokenIndex = position924, tokenIndex924
																if buffer[position] != rune('L') {
																	goto l809
																}
																position++
															}
														l924:
															{
																position926, tokenIndex926 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l927
																}
																position++
																goto l926
															l927:
																position, tokenIndex = position926, tokenIndex926
																if buffer[position] != rune('T') {
																	goto l809
																}
																position++
															}
														l926:
															add(rulePegText, position919)
														}
														{
															add(ruleAction47, position)
														}
														add(ruleHalt, position918)
													}
													break
												default:
													{
														position929 := position
														{
															position930 := position
															{
																position931, tokenIndex931 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l932
																}
																position++
																goto l931
															l932:
																position, tokenIndex = position931, tokenIndex931
																if buffer[position] != rune('N') {
																	goto l809
																}
																position++
															}
														l931:
															{
																position933, tokenIndex933 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l934
																}
																position++
																goto l933
															l934:
																position, tokenIndex = position933, tokenIndex933
																if buffer[position] != rune('O') {
																	goto l809
																}
																position++
															}
														l933:
															{
																position935, tokenIndex935 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l936
																}
																position++
																goto l935
															l936:
																position, tokenIndex = position935, tokenIndex935
																if buffer[position] != rune('P') {
																	goto l809
																}
																position++
															}
														l935:
															add(rulePegText, position930)
														}
														{
															add(ruleAction46, position)
														}
														add(ruleNop, position929)
													}
													break
												}
											}

										}
									l811:
										add(ruleSimple, position810)
									}
									goto l521
								l809:
									position, tokenIndex = position521, tokenIndex521
									{
										position939 := position
										{
											position940, tokenIndex940 := position, tokenIndex
											{
												position942 := position
												{
													position943, tokenIndex943 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l944
													}
													position++
													goto l943
												l944:
													position, tokenIndex = position943, tokenIndex943
													if buffer[position] != rune('R') {
														goto l941
													}
													position++
												}
											l943:
												{
													position945, tokenIndex945 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l946
													}
													position++
													goto l945
												l946:
													position, tokenIndex = position945, tokenIndex945
													if buffer[position] != rune('S') {
														goto l941
													}
													position++
												}
											l945:
												{
													position947, tokenIndex947 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l948
													}
													position++
													goto l947
												l948:
													position, tokenIndex = position947, tokenIndex947
													if buffer[position] != rune('T') {
														goto l941
													}
													position++
												}
											l947:
												if !_rules[rulews]() {
													goto l941
												}
												if !_rules[rulen]() {
													goto l941
												}
												{
													add(ruleAction59, position)
												}
												add(ruleRst, position942)
											}
											goto l940
										l941:
											position, tokenIndex = position940, tokenIndex940
											{
												position951 := position
												{
													position952, tokenIndex952 := position, tokenIndex
													if buffer[position] != rune('j') {
														goto l953
													}
													position++
													goto l952
												l953:
													position, tokenIndex = position952, tokenIndex952
													if buffer[position] != rune('J') {
														goto l950
													}
													position++
												}
											l952:
												{
													position954, tokenIndex954 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l955
													}
													position++
													goto l954
												l955:
													position, tokenIndex = position954, tokenIndex954
													if buffer[position] != rune('P') {
														goto l950
													}
													position++
												}
											l954:
												if !_rules[rulews]() {
													goto l950
												}
												{
													position956, tokenIndex956 := position, tokenIndex
													if !_rules[rulecc]() {
														goto l956
													}
													if !_rules[rulesep]() {
														goto l956
													}
													goto l957
												l956:
													position, tokenIndex = position956, tokenIndex956
												}
											l957:
												if !_rules[ruleSrc16]() {
													goto l950
												}
												{
													add(ruleAction62, position)
												}
												add(ruleJp, position951)
											}
											goto l940
										l950:
											position, tokenIndex = position940, tokenIndex940
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position960 := position
														{
															position961, tokenIndex961 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l962
															}
															position++
															goto l961
														l962:
															position, tokenIndex = position961, tokenIndex961
															if buffer[position] != rune('D') {
																goto l938
															}
															position++
														}
													l961:
														{
															position963, tokenIndex963 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l964
															}
															position++
															goto l963
														l964:
															position, tokenIndex = position963, tokenIndex963
															if buffer[position] != rune('J') {
																goto l938
															}
															position++
														}
													l963:
														{
															position965, tokenIndex965 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l966
															}
															position++
															goto l965
														l966:
															position, tokenIndex = position965, tokenIndex965
															if buffer[position] != rune('N') {
																goto l938
															}
															position++
														}
													l965:
														{
															position967, tokenIndex967 := position, tokenIndex
															if buffer[position] != rune('z') {
																goto l968
															}
															position++
															goto l967
														l968:
															position, tokenIndex = position967, tokenIndex967
															if buffer[position] != rune('Z') {
																goto l938
															}
															position++
														}
													l967:
														if !_rules[rulews]() {
															goto l938
														}
														if !_rules[ruledisp]() {
															goto l938
														}
														{
															add(ruleAction64, position)
														}
														add(ruleDjnz, position960)
													}
													break
												case 'J', 'j':
													{
														position970 := position
														{
															position971, tokenIndex971 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l972
															}
															position++
															goto l971
														l972:
															position, tokenIndex = position971, tokenIndex971
															if buffer[position] != rune('J') {
																goto l938
															}
															position++
														}
													l971:
														{
															position973, tokenIndex973 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l974
															}
															position++
															goto l973
														l974:
															position, tokenIndex = position973, tokenIndex973
															if buffer[position] != rune('R') {
																goto l938
															}
															position++
														}
													l973:
														if !_rules[rulews]() {
															goto l938
														}
														{
															position975, tokenIndex975 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l975
															}
															if !_rules[rulesep]() {
																goto l975
															}
															goto l976
														l975:
															position, tokenIndex = position975, tokenIndex975
														}
													l976:
														if !_rules[ruledisp]() {
															goto l938
														}
														{
															add(ruleAction63, position)
														}
														add(ruleJr, position970)
													}
													break
												case 'R', 'r':
													{
														position978 := position
														{
															position979, tokenIndex979 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l980
															}
															position++
															goto l979
														l980:
															position, tokenIndex = position979, tokenIndex979
															if buffer[position] != rune('R') {
																goto l938
															}
															position++
														}
													l979:
														{
															position981, tokenIndex981 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l982
															}
															position++
															goto l981
														l982:
															position, tokenIndex = position981, tokenIndex981
															if buffer[position] != rune('E') {
																goto l938
															}
															position++
														}
													l981:
														{
															position983, tokenIndex983 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l984
															}
															position++
															goto l983
														l984:
															position, tokenIndex = position983, tokenIndex983
															if buffer[position] != rune('T') {
																goto l938
															}
															position++
														}
													l983:
														{
															position985, tokenIndex985 := position, tokenIndex
															if !_rules[rulews]() {
																goto l985
															}
															if !_rules[rulecc]() {
																goto l985
															}
															goto l986
														l985:
															position, tokenIndex = position985, tokenIndex985
														}
													l986:
														{
															add(ruleAction61, position)
														}
														add(ruleRet, position978)
													}
													break
												default:
													{
														position988 := position
														{
															position989, tokenIndex989 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l990
															}
															position++
															goto l989
														l990:
															position, tokenIndex = position989, tokenIndex989
															if buffer[position] != rune('C') {
																goto l938
															}
															position++
														}
													l989:
														{
															position991, tokenIndex991 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l992
															}
															position++
															goto l991
														l992:
															position, tokenIndex = position991, tokenIndex991
															if buffer[position] != rune('A') {
																goto l938
															}
															position++
														}
													l991:
														{
															position993, tokenIndex993 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l994
															}
															position++
															goto l993
														l994:
															position, tokenIndex = position993, tokenIndex993
															if buffer[position] != rune('L') {
																goto l938
															}
															position++
														}
													l993:
														{
															position995, tokenIndex995 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l996
															}
															position++
															goto l995
														l996:
															position, tokenIndex = position995, tokenIndex995
															if buffer[position] != rune('L') {
																goto l938
															}
															position++
														}
													l995:
														if !_rules[rulews]() {
															goto l938
														}
														{
															position997, tokenIndex997 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l997
															}
															if !_rules[rulesep]() {
																goto l997
															}
															goto l998
														l997:
															position, tokenIndex = position997, tokenIndex997
														}
													l998:
														if !_rules[ruleSrc16]() {
															goto l938
														}
														{
															add(ruleAction60, position)
														}
														add(ruleCall, position988)
													}
													break
												}
											}

										}
									l940:
										add(ruleJump, position939)
									}
									goto l521
								l938:
									position, tokenIndex = position521, tokenIndex521
									{
										position1000 := position
										{
											position1001, tokenIndex1001 := position, tokenIndex
											{
												position1003 := position
												{
													position1004, tokenIndex1004 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1005
													}
													position++
													goto l1004
												l1005:
													position, tokenIndex = position1004, tokenIndex1004
													if buffer[position] != rune('I') {
														goto l1002
													}
													position++
												}
											l1004:
												{
													position1006, tokenIndex1006 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1007
													}
													position++
													goto l1006
												l1007:
													position, tokenIndex = position1006, tokenIndex1006
													if buffer[position] != rune('N') {
														goto l1002
													}
													position++
												}
											l1006:
												if !_rules[rulews]() {
													goto l1002
												}
												if !_rules[ruleReg8]() {
													goto l1002
												}
												if !_rules[rulesep]() {
													goto l1002
												}
												if !_rules[rulePort]() {
													goto l1002
												}
												{
													add(ruleAction65, position)
												}
												add(ruleIN, position1003)
											}
											goto l1001
										l1002:
											position, tokenIndex = position1001, tokenIndex1001
											{
												position1009 := position
												{
													position1010, tokenIndex1010 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1011
													}
													position++
													goto l1010
												l1011:
													position, tokenIndex = position1010, tokenIndex1010
													if buffer[position] != rune('O') {
														goto l3
													}
													position++
												}
											l1010:
												{
													position1012, tokenIndex1012 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1013
													}
													position++
													goto l1012
												l1013:
													position, tokenIndex = position1012, tokenIndex1012
													if buffer[position] != rune('U') {
														goto l3
													}
													position++
												}
											l1012:
												{
													position1014, tokenIndex1014 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1015
													}
													position++
													goto l1014
												l1015:
													position, tokenIndex = position1014, tokenIndex1014
													if buffer[position] != rune('T') {
														goto l3
													}
													position++
												}
											l1014:
												if !_rules[rulews]() {
													goto l3
												}
												if !_rules[rulePort]() {
													goto l3
												}
												if !_rules[rulesep]() {
													goto l3
												}
												if !_rules[ruleReg8]() {
													goto l3
												}
												{
													add(ruleAction66, position)
												}
												add(ruleOUT, position1009)
											}
										}
									l1001:
										add(ruleIO, position1000)
									}
								}
							l521:
								add(ruleInstruction, position518)
							}
							{
								position1017, tokenIndex1017 := position, tokenIndex
								if buffer[position] != rune('\n') {
									goto l1017
								}
								position++
								goto l1018
							l1017:
								position, tokenIndex = position1017, tokenIndex1017
							}
						l1018:
							{
								add(ruleAction0, position)
							}
							add(ruleLine, position517)
						}
					}
				l512:
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
		/* 3 Instruction <- <(ws* (Assignment / Inc / Dec / Add16 / Alu / BitOp / Simple / Jump / IO))> */
		nil,
		/* 4 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 5 Load <- <(Load16 / Load8)> */
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
		/* 11 Inc <- <(Inc16Indexed8 / Inc16 / Inc8)> */
		nil,
		/* 12 Inc16Indexed8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws ILoc8 Action6)> */
		nil,
		/* 13 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action7)> */
		nil,
		/* 14 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action8)> */
		nil,
		/* 15 Dec <- <(Dec16Indexed8 / Dec16 / Dec8)> */
		nil,
		/* 16 Dec16Indexed8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws ILoc8 Action9)> */
		nil,
		/* 17 Dec8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc8 Action10)> */
		nil,
		/* 18 Dec16 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc16 Action11)> */
		nil,
		/* 19 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action12)> */
		nil,
		/* 20 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action13)> */
		nil,
		/* 21 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action14)> */
		func() bool {
			position1040, tokenIndex1040 := position, tokenIndex
			{
				position1041 := position
				{
					position1042, tokenIndex1042 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1043
					}
					goto l1042
				l1043:
					position, tokenIndex = position1042, tokenIndex1042
					if !_rules[ruleReg8]() {
						goto l1044
					}
					goto l1042
				l1044:
					position, tokenIndex = position1042, tokenIndex1042
					if !_rules[ruleReg16Contents]() {
						goto l1045
					}
					goto l1042
				l1045:
					position, tokenIndex = position1042, tokenIndex1042
					if !_rules[rulenn_contents]() {
						goto l1040
					}
				}
			l1042:
				{
					add(ruleAction14, position)
				}
				add(ruleSrc8, position1041)
			}
			return true
		l1040:
			position, tokenIndex = position1040, tokenIndex1040
			return false
		},
		/* 22 Loc8 <- <((Reg8 / Reg16Contents) Action15)> */
		func() bool {
			position1047, tokenIndex1047 := position, tokenIndex
			{
				position1048 := position
				{
					position1049, tokenIndex1049 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1050
					}
					goto l1049
				l1050:
					position, tokenIndex = position1049, tokenIndex1049
					if !_rules[ruleReg16Contents]() {
						goto l1047
					}
				}
			l1049:
				{
					add(ruleAction15, position)
				}
				add(ruleLoc8, position1048)
			}
			return true
		l1047:
			position, tokenIndex = position1047, tokenIndex1047
			return false
		},
		/* 23 ILoc8 <- <(IReg8 Action16)> */
		func() bool {
			position1052, tokenIndex1052 := position, tokenIndex
			{
				position1053 := position
				if !_rules[ruleIReg8]() {
					goto l1052
				}
				{
					add(ruleAction16, position)
				}
				add(ruleILoc8, position1053)
			}
			return true
		l1052:
			position, tokenIndex = position1052, tokenIndex1052
			return false
		},
		/* 24 Reg8 <- <(<((&('I' | 'i') IReg8) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A))> Action17)> */
		func() bool {
			position1055, tokenIndex1055 := position, tokenIndex
			{
				position1056 := position
				{
					position1057 := position
					{
						switch buffer[position] {
						case 'I', 'i':
							if !_rules[ruleIReg8]() {
								goto l1055
							}
							break
						case 'L', 'l':
							{
								position1059 := position
								{
									position1060, tokenIndex1060 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l1061
									}
									position++
									goto l1060
								l1061:
									position, tokenIndex = position1060, tokenIndex1060
									if buffer[position] != rune('L') {
										goto l1055
									}
									position++
								}
							l1060:
								add(ruleL, position1059)
							}
							break
						case 'H', 'h':
							{
								position1062 := position
								{
									position1063, tokenIndex1063 := position, tokenIndex
									if buffer[position] != rune('h') {
										goto l1064
									}
									position++
									goto l1063
								l1064:
									position, tokenIndex = position1063, tokenIndex1063
									if buffer[position] != rune('H') {
										goto l1055
									}
									position++
								}
							l1063:
								add(ruleH, position1062)
							}
							break
						case 'E', 'e':
							{
								position1065 := position
								{
									position1066, tokenIndex1066 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l1067
									}
									position++
									goto l1066
								l1067:
									position, tokenIndex = position1066, tokenIndex1066
									if buffer[position] != rune('E') {
										goto l1055
									}
									position++
								}
							l1066:
								add(ruleE, position1065)
							}
							break
						case 'D', 'd':
							{
								position1068 := position
								{
									position1069, tokenIndex1069 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l1070
									}
									position++
									goto l1069
								l1070:
									position, tokenIndex = position1069, tokenIndex1069
									if buffer[position] != rune('D') {
										goto l1055
									}
									position++
								}
							l1069:
								add(ruleD, position1068)
							}
							break
						case 'C', 'c':
							{
								position1071 := position
								{
									position1072, tokenIndex1072 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1073
									}
									position++
									goto l1072
								l1073:
									position, tokenIndex = position1072, tokenIndex1072
									if buffer[position] != rune('C') {
										goto l1055
									}
									position++
								}
							l1072:
								add(ruleC, position1071)
							}
							break
						case 'B', 'b':
							{
								position1074 := position
								{
									position1075, tokenIndex1075 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l1076
									}
									position++
									goto l1075
								l1076:
									position, tokenIndex = position1075, tokenIndex1075
									if buffer[position] != rune('B') {
										goto l1055
									}
									position++
								}
							l1075:
								add(ruleB, position1074)
							}
							break
						case 'F', 'f':
							{
								position1077 := position
								{
									position1078, tokenIndex1078 := position, tokenIndex
									if buffer[position] != rune('f') {
										goto l1079
									}
									position++
									goto l1078
								l1079:
									position, tokenIndex = position1078, tokenIndex1078
									if buffer[position] != rune('F') {
										goto l1055
									}
									position++
								}
							l1078:
								add(ruleF, position1077)
							}
							break
						default:
							{
								position1080 := position
								{
									position1081, tokenIndex1081 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l1082
									}
									position++
									goto l1081
								l1082:
									position, tokenIndex = position1081, tokenIndex1081
									if buffer[position] != rune('A') {
										goto l1055
									}
									position++
								}
							l1081:
								add(ruleA, position1080)
							}
							break
						}
					}

					add(rulePegText, position1057)
				}
				{
					add(ruleAction17, position)
				}
				add(ruleReg8, position1056)
			}
			return true
		l1055:
			position, tokenIndex = position1055, tokenIndex1055
			return false
		},
		/* 25 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action18)> */
		func() bool {
			position1084, tokenIndex1084 := position, tokenIndex
			{
				position1085 := position
				{
					position1086 := position
					{
						position1087, tokenIndex1087 := position, tokenIndex
						{
							position1089 := position
							{
								position1090, tokenIndex1090 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1091
								}
								position++
								goto l1090
							l1091:
								position, tokenIndex = position1090, tokenIndex1090
								if buffer[position] != rune('I') {
									goto l1088
								}
								position++
							}
						l1090:
							{
								position1092, tokenIndex1092 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1093
								}
								position++
								goto l1092
							l1093:
								position, tokenIndex = position1092, tokenIndex1092
								if buffer[position] != rune('X') {
									goto l1088
								}
								position++
							}
						l1092:
							{
								position1094, tokenIndex1094 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1095
								}
								position++
								goto l1094
							l1095:
								position, tokenIndex = position1094, tokenIndex1094
								if buffer[position] != rune('H') {
									goto l1088
								}
								position++
							}
						l1094:
							add(ruleIXH, position1089)
						}
						goto l1087
					l1088:
						position, tokenIndex = position1087, tokenIndex1087
						{
							position1097 := position
							{
								position1098, tokenIndex1098 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1099
								}
								position++
								goto l1098
							l1099:
								position, tokenIndex = position1098, tokenIndex1098
								if buffer[position] != rune('I') {
									goto l1096
								}
								position++
							}
						l1098:
							{
								position1100, tokenIndex1100 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1101
								}
								position++
								goto l1100
							l1101:
								position, tokenIndex = position1100, tokenIndex1100
								if buffer[position] != rune('X') {
									goto l1096
								}
								position++
							}
						l1100:
							{
								position1102, tokenIndex1102 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1103
								}
								position++
								goto l1102
							l1103:
								position, tokenIndex = position1102, tokenIndex1102
								if buffer[position] != rune('L') {
									goto l1096
								}
								position++
							}
						l1102:
							add(ruleIXL, position1097)
						}
						goto l1087
					l1096:
						position, tokenIndex = position1087, tokenIndex1087
						{
							position1105 := position
							{
								position1106, tokenIndex1106 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1107
								}
								position++
								goto l1106
							l1107:
								position, tokenIndex = position1106, tokenIndex1106
								if buffer[position] != rune('I') {
									goto l1104
								}
								position++
							}
						l1106:
							{
								position1108, tokenIndex1108 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1109
								}
								position++
								goto l1108
							l1109:
								position, tokenIndex = position1108, tokenIndex1108
								if buffer[position] != rune('Y') {
									goto l1104
								}
								position++
							}
						l1108:
							{
								position1110, tokenIndex1110 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1111
								}
								position++
								goto l1110
							l1111:
								position, tokenIndex = position1110, tokenIndex1110
								if buffer[position] != rune('H') {
									goto l1104
								}
								position++
							}
						l1110:
							add(ruleIYH, position1105)
						}
						goto l1087
					l1104:
						position, tokenIndex = position1087, tokenIndex1087
						{
							position1112 := position
							{
								position1113, tokenIndex1113 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1114
								}
								position++
								goto l1113
							l1114:
								position, tokenIndex = position1113, tokenIndex1113
								if buffer[position] != rune('I') {
									goto l1084
								}
								position++
							}
						l1113:
							{
								position1115, tokenIndex1115 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1116
								}
								position++
								goto l1115
							l1116:
								position, tokenIndex = position1115, tokenIndex1115
								if buffer[position] != rune('Y') {
									goto l1084
								}
								position++
							}
						l1115:
							{
								position1117, tokenIndex1117 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1118
								}
								position++
								goto l1117
							l1118:
								position, tokenIndex = position1117, tokenIndex1117
								if buffer[position] != rune('L') {
									goto l1084
								}
								position++
							}
						l1117:
							add(ruleIYL, position1112)
						}
					}
				l1087:
					add(rulePegText, position1086)
				}
				{
					add(ruleAction18, position)
				}
				add(ruleIReg8, position1085)
			}
			return true
		l1084:
			position, tokenIndex = position1084, tokenIndex1084
			return false
		},
		/* 26 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action19)> */
		func() bool {
			position1120, tokenIndex1120 := position, tokenIndex
			{
				position1121 := position
				{
					position1122, tokenIndex1122 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1123
					}
					goto l1122
				l1123:
					position, tokenIndex = position1122, tokenIndex1122
					if !_rules[rulenn_contents]() {
						goto l1124
					}
					goto l1122
				l1124:
					position, tokenIndex = position1122, tokenIndex1122
					if !_rules[ruleReg16Contents]() {
						goto l1120
					}
				}
			l1122:
				{
					add(ruleAction19, position)
				}
				add(ruleDst16, position1121)
			}
			return true
		l1120:
			position, tokenIndex = position1120, tokenIndex1120
			return false
		},
		/* 27 Src16 <- <((Reg16 / nn / nn_contents) Action20)> */
		func() bool {
			position1126, tokenIndex1126 := position, tokenIndex
			{
				position1127 := position
				{
					position1128, tokenIndex1128 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1129
					}
					goto l1128
				l1129:
					position, tokenIndex = position1128, tokenIndex1128
					if !_rules[rulenn]() {
						goto l1130
					}
					goto l1128
				l1130:
					position, tokenIndex = position1128, tokenIndex1128
					if !_rules[rulenn_contents]() {
						goto l1126
					}
				}
			l1128:
				{
					add(ruleAction20, position)
				}
				add(ruleSrc16, position1127)
			}
			return true
		l1126:
			position, tokenIndex = position1126, tokenIndex1126
			return false
		},
		/* 28 Loc16 <- <(Reg16 Action21)> */
		func() bool {
			position1132, tokenIndex1132 := position, tokenIndex
			{
				position1133 := position
				if !_rules[ruleReg16]() {
					goto l1132
				}
				{
					add(ruleAction21, position)
				}
				add(ruleLoc16, position1133)
			}
			return true
		l1132:
			position, tokenIndex = position1132, tokenIndex1132
			return false
		},
		/* 29 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action22)> */
		func() bool {
			position1135, tokenIndex1135 := position, tokenIndex
			{
				position1136 := position
				{
					position1137 := position
					{
						position1138, tokenIndex1138 := position, tokenIndex
						{
							position1140 := position
							{
								position1141, tokenIndex1141 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1142
								}
								position++
								goto l1141
							l1142:
								position, tokenIndex = position1141, tokenIndex1141
								if buffer[position] != rune('A') {
									goto l1139
								}
								position++
							}
						l1141:
							{
								position1143, tokenIndex1143 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1144
								}
								position++
								goto l1143
							l1144:
								position, tokenIndex = position1143, tokenIndex1143
								if buffer[position] != rune('F') {
									goto l1139
								}
								position++
							}
						l1143:
							if buffer[position] != rune('\'') {
								goto l1139
							}
							position++
							add(ruleAF_PRIME, position1140)
						}
						goto l1138
					l1139:
						position, tokenIndex = position1138, tokenIndex1138
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1135
								}
								break
							case 'S', 's':
								{
									position1146 := position
									{
										position1147, tokenIndex1147 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1148
										}
										position++
										goto l1147
									l1148:
										position, tokenIndex = position1147, tokenIndex1147
										if buffer[position] != rune('S') {
											goto l1135
										}
										position++
									}
								l1147:
									{
										position1149, tokenIndex1149 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1150
										}
										position++
										goto l1149
									l1150:
										position, tokenIndex = position1149, tokenIndex1149
										if buffer[position] != rune('P') {
											goto l1135
										}
										position++
									}
								l1149:
									add(ruleSP, position1146)
								}
								break
							case 'H', 'h':
								{
									position1151 := position
									{
										position1152, tokenIndex1152 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1153
										}
										position++
										goto l1152
									l1153:
										position, tokenIndex = position1152, tokenIndex1152
										if buffer[position] != rune('H') {
											goto l1135
										}
										position++
									}
								l1152:
									{
										position1154, tokenIndex1154 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1155
										}
										position++
										goto l1154
									l1155:
										position, tokenIndex = position1154, tokenIndex1154
										if buffer[position] != rune('L') {
											goto l1135
										}
										position++
									}
								l1154:
									add(ruleHL, position1151)
								}
								break
							case 'D', 'd':
								{
									position1156 := position
									{
										position1157, tokenIndex1157 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1158
										}
										position++
										goto l1157
									l1158:
										position, tokenIndex = position1157, tokenIndex1157
										if buffer[position] != rune('D') {
											goto l1135
										}
										position++
									}
								l1157:
									{
										position1159, tokenIndex1159 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1160
										}
										position++
										goto l1159
									l1160:
										position, tokenIndex = position1159, tokenIndex1159
										if buffer[position] != rune('E') {
											goto l1135
										}
										position++
									}
								l1159:
									add(ruleDE, position1156)
								}
								break
							case 'B', 'b':
								{
									position1161 := position
									{
										position1162, tokenIndex1162 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1163
										}
										position++
										goto l1162
									l1163:
										position, tokenIndex = position1162, tokenIndex1162
										if buffer[position] != rune('B') {
											goto l1135
										}
										position++
									}
								l1162:
									{
										position1164, tokenIndex1164 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1165
										}
										position++
										goto l1164
									l1165:
										position, tokenIndex = position1164, tokenIndex1164
										if buffer[position] != rune('C') {
											goto l1135
										}
										position++
									}
								l1164:
									add(ruleBC, position1161)
								}
								break
							default:
								{
									position1166 := position
									{
										position1167, tokenIndex1167 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1168
										}
										position++
										goto l1167
									l1168:
										position, tokenIndex = position1167, tokenIndex1167
										if buffer[position] != rune('A') {
											goto l1135
										}
										position++
									}
								l1167:
									{
										position1169, tokenIndex1169 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1170
										}
										position++
										goto l1169
									l1170:
										position, tokenIndex = position1169, tokenIndex1169
										if buffer[position] != rune('F') {
											goto l1135
										}
										position++
									}
								l1169:
									add(ruleAF, position1166)
								}
								break
							}
						}

					}
				l1138:
					add(rulePegText, position1137)
				}
				{
					add(ruleAction22, position)
				}
				add(ruleReg16, position1136)
			}
			return true
		l1135:
			position, tokenIndex = position1135, tokenIndex1135
			return false
		},
		/* 30 IReg16 <- <(<(IX / IY)> Action23)> */
		func() bool {
			position1172, tokenIndex1172 := position, tokenIndex
			{
				position1173 := position
				{
					position1174 := position
					{
						position1175, tokenIndex1175 := position, tokenIndex
						{
							position1177 := position
							{
								position1178, tokenIndex1178 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1179
								}
								position++
								goto l1178
							l1179:
								position, tokenIndex = position1178, tokenIndex1178
								if buffer[position] != rune('I') {
									goto l1176
								}
								position++
							}
						l1178:
							{
								position1180, tokenIndex1180 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1181
								}
								position++
								goto l1180
							l1181:
								position, tokenIndex = position1180, tokenIndex1180
								if buffer[position] != rune('X') {
									goto l1176
								}
								position++
							}
						l1180:
							add(ruleIX, position1177)
						}
						goto l1175
					l1176:
						position, tokenIndex = position1175, tokenIndex1175
						{
							position1182 := position
							{
								position1183, tokenIndex1183 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1184
								}
								position++
								goto l1183
							l1184:
								position, tokenIndex = position1183, tokenIndex1183
								if buffer[position] != rune('I') {
									goto l1172
								}
								position++
							}
						l1183:
							{
								position1185, tokenIndex1185 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1186
								}
								position++
								goto l1185
							l1186:
								position, tokenIndex = position1185, tokenIndex1185
								if buffer[position] != rune('Y') {
									goto l1172
								}
								position++
							}
						l1185:
							add(ruleIY, position1182)
						}
					}
				l1175:
					add(rulePegText, position1174)
				}
				{
					add(ruleAction23, position)
				}
				add(ruleIReg16, position1173)
			}
			return true
		l1172:
			position, tokenIndex = position1172, tokenIndex1172
			return false
		},
		/* 31 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1188, tokenIndex1188 := position, tokenIndex
			{
				position1189 := position
				{
					position1190, tokenIndex1190 := position, tokenIndex
					{
						position1192 := position
						if buffer[position] != rune('(') {
							goto l1191
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1191
						}
						{
							position1193, tokenIndex1193 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1193
							}
							goto l1194
						l1193:
							position, tokenIndex = position1193, tokenIndex1193
						}
					l1194:
						if !_rules[rulesignedDecimalByte]() {
							goto l1191
						}
						{
							position1195, tokenIndex1195 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1195
							}
							goto l1196
						l1195:
							position, tokenIndex = position1195, tokenIndex1195
						}
					l1196:
						if buffer[position] != rune(')') {
							goto l1191
						}
						position++
						{
							add(ruleAction25, position)
						}
						add(ruleIndexedR16C, position1192)
					}
					goto l1190
				l1191:
					position, tokenIndex = position1190, tokenIndex1190
					{
						position1198 := position
						if buffer[position] != rune('(') {
							goto l1188
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1188
						}
						if buffer[position] != rune(')') {
							goto l1188
						}
						position++
						{
							add(ruleAction24, position)
						}
						add(rulePlainR16C, position1198)
					}
				}
			l1190:
				add(ruleReg16Contents, position1189)
			}
			return true
		l1188:
			position, tokenIndex = position1188, tokenIndex1188
			return false
		},
		/* 32 PlainR16C <- <('(' Reg16 ')' Action24)> */
		nil,
		/* 33 IndexedR16C <- <('(' IReg16 ws? signedDecimalByte ws? ')' Action25)> */
		nil,
		/* 34 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1202, tokenIndex1202 := position, tokenIndex
			{
				position1203 := position
				{
					position1204, tokenIndex1204 := position, tokenIndex
					{
						position1206 := position
						{
							position1207 := position
							if !_rules[rulehexdigit]() {
								goto l1205
							}
							if !_rules[rulehexdigit]() {
								goto l1205
							}
							add(rulePegText, position1207)
						}
						{
							position1208, tokenIndex1208 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1209
							}
							position++
							goto l1208
						l1209:
							position, tokenIndex = position1208, tokenIndex1208
							if buffer[position] != rune('H') {
								goto l1205
							}
							position++
						}
					l1208:
						{
							add(ruleAction67, position)
						}
						add(rulehexByteH, position1206)
					}
					goto l1204
				l1205:
					position, tokenIndex = position1204, tokenIndex1204
					{
						position1212 := position
						if buffer[position] != rune('0') {
							goto l1211
						}
						position++
						{
							position1213, tokenIndex1213 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1214
							}
							position++
							goto l1213
						l1214:
							position, tokenIndex = position1213, tokenIndex1213
							if buffer[position] != rune('X') {
								goto l1211
							}
							position++
						}
					l1213:
						{
							position1215 := position
							if !_rules[rulehexdigit]() {
								goto l1211
							}
							if !_rules[rulehexdigit]() {
								goto l1211
							}
							add(rulePegText, position1215)
						}
						{
							add(ruleAction68, position)
						}
						add(rulehexByte0x, position1212)
					}
					goto l1204
				l1211:
					position, tokenIndex = position1204, tokenIndex1204
					{
						position1217 := position
						{
							position1218 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1202
							}
							position++
						l1219:
							{
								position1220, tokenIndex1220 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1220
								}
								position++
								goto l1219
							l1220:
								position, tokenIndex = position1220, tokenIndex1220
							}
							add(rulePegText, position1218)
						}
						{
							add(ruleAction69, position)
						}
						add(ruledecimalByte, position1217)
					}
				}
			l1204:
				add(rulen, position1203)
			}
			return true
		l1202:
			position, tokenIndex = position1202, tokenIndex1202
			return false
		},
		/* 35 nn <- <(hexWordH / hexWord0x)> */
		func() bool {
			position1222, tokenIndex1222 := position, tokenIndex
			{
				position1223 := position
				{
					position1224, tokenIndex1224 := position, tokenIndex
					{
						position1226 := position
						{
							position1227 := position
							if !_rules[rulehexdigit]() {
								goto l1225
							}
							if !_rules[rulehexdigit]() {
								goto l1225
							}
							if !_rules[rulehexdigit]() {
								goto l1225
							}
							if !_rules[rulehexdigit]() {
								goto l1225
							}
							add(rulePegText, position1227)
						}
						{
							position1228, tokenIndex1228 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1229
							}
							position++
							goto l1228
						l1229:
							position, tokenIndex = position1228, tokenIndex1228
							if buffer[position] != rune('H') {
								goto l1225
							}
							position++
						}
					l1228:
						{
							add(ruleAction70, position)
						}
						add(rulehexWordH, position1226)
					}
					goto l1224
				l1225:
					position, tokenIndex = position1224, tokenIndex1224
					{
						position1231 := position
						if buffer[position] != rune('0') {
							goto l1222
						}
						position++
						{
							position1232, tokenIndex1232 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1233
							}
							position++
							goto l1232
						l1233:
							position, tokenIndex = position1232, tokenIndex1232
							if buffer[position] != rune('X') {
								goto l1222
							}
							position++
						}
					l1232:
						{
							position1234 := position
							if !_rules[rulehexdigit]() {
								goto l1222
							}
							if !_rules[rulehexdigit]() {
								goto l1222
							}
							if !_rules[rulehexdigit]() {
								goto l1222
							}
							if !_rules[rulehexdigit]() {
								goto l1222
							}
							add(rulePegText, position1234)
						}
						{
							add(ruleAction71, position)
						}
						add(rulehexWord0x, position1231)
					}
				}
			l1224:
				add(rulenn, position1223)
			}
			return true
		l1222:
			position, tokenIndex = position1222, tokenIndex1222
			return false
		},
		/* 36 nn_contents <- <('(' nn ')' Action26)> */
		func() bool {
			position1236, tokenIndex1236 := position, tokenIndex
			{
				position1237 := position
				if buffer[position] != rune('(') {
					goto l1236
				}
				position++
				if !_rules[rulenn]() {
					goto l1236
				}
				if buffer[position] != rune(')') {
					goto l1236
				}
				position++
				{
					add(ruleAction26, position)
				}
				add(rulenn_contents, position1237)
			}
			return true
		l1236:
			position, tokenIndex = position1236, tokenIndex1236
			return false
		},
		/* 37 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 38 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action27)> */
		nil,
		/* 39 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action28)> */
		nil,
		/* 40 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action29)> */
		nil,
		/* 41 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action30)> */
		nil,
		/* 42 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action31)> */
		nil,
		/* 43 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action32)> */
		nil,
		/* 44 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action33)> */
		nil,
		/* 45 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action34)> */
		nil,
		/* 46 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 47 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 48 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 Action35)> */
		nil,
		/* 49 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 Action36)> */
		nil,
		/* 50 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 Action37)> */
		nil,
		/* 51 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 Action38)> */
		nil,
		/* 52 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 Action39)> */
		nil,
		/* 53 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 Action40)> */
		nil,
		/* 54 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 Action41)> */
		nil,
		/* 55 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 Action42)> */
		nil,
		/* 56 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action43)> */
		nil,
		/* 57 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 Action44)> */
		nil,
		/* 58 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 Action45)> */
		nil,
		/* 59 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 60 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action46)> */
		nil,
		/* 61 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action47)> */
		nil,
		/* 62 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action48)> */
		nil,
		/* 63 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action49)> */
		nil,
		/* 64 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action50)> */
		nil,
		/* 65 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action51)> */
		nil,
		/* 66 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action52)> */
		nil,
		/* 67 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action53)> */
		nil,
		/* 68 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action54)> */
		nil,
		/* 69 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action55)> */
		nil,
		/* 70 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action56)> */
		nil,
		/* 71 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action57)> */
		nil,
		/* 72 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action58)> */
		nil,
		/* 73 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 74 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action59)> */
		nil,
		/* 75 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action60)> */
		nil,
		/* 76 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action61)> */
		nil,
		/* 77 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action62)> */
		nil,
		/* 78 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action63)> */
		nil,
		/* 79 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action64)> */
		nil,
		/* 80 disp <- <signedDecimalByte> */
		func() bool {
			position1282, tokenIndex1282 := position, tokenIndex
			{
				position1283 := position
				if !_rules[rulesignedDecimalByte]() {
					goto l1282
				}
				add(ruledisp, position1283)
			}
			return true
		l1282:
			position, tokenIndex = position1282, tokenIndex1282
			return false
		},
		/* 81 IO <- <(IN / OUT)> */
		nil,
		/* 82 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action65)> */
		nil,
		/* 83 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action66)> */
		nil,
		/* 84 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position1287, tokenIndex1287 := position, tokenIndex
			{
				position1288 := position
				{
					position1289, tokenIndex1289 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l1290
					}
					position++
					{
						position1291, tokenIndex1291 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l1292
						}
						position++
						goto l1291
					l1292:
						position, tokenIndex = position1291, tokenIndex1291
						if buffer[position] != rune('C') {
							goto l1290
						}
						position++
					}
				l1291:
					if buffer[position] != rune(')') {
						goto l1290
					}
					position++
					goto l1289
				l1290:
					position, tokenIndex = position1289, tokenIndex1289
					if buffer[position] != rune('(') {
						goto l1287
					}
					position++
					if !_rules[rulen]() {
						goto l1287
					}
					if buffer[position] != rune(')') {
						goto l1287
					}
					position++
				}
			l1289:
				add(rulePort, position1288)
			}
			return true
		l1287:
			position, tokenIndex = position1287, tokenIndex1287
			return false
		},
		/* 85 sep <- <(ws? ',' ws?)> */
		func() bool {
			position1293, tokenIndex1293 := position, tokenIndex
			{
				position1294 := position
				{
					position1295, tokenIndex1295 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1295
					}
					goto l1296
				l1295:
					position, tokenIndex = position1295, tokenIndex1295
				}
			l1296:
				if buffer[position] != rune(',') {
					goto l1293
				}
				position++
				{
					position1297, tokenIndex1297 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1297
					}
					goto l1298
				l1297:
					position, tokenIndex = position1297, tokenIndex1297
				}
			l1298:
				add(rulesep, position1294)
			}
			return true
		l1293:
			position, tokenIndex = position1293, tokenIndex1293
			return false
		},
		/* 86 ws <- <' '+> */
		func() bool {
			position1299, tokenIndex1299 := position, tokenIndex
			{
				position1300 := position
				if buffer[position] != rune(' ') {
					goto l1299
				}
				position++
			l1301:
				{
					position1302, tokenIndex1302 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l1302
					}
					position++
					goto l1301
				l1302:
					position, tokenIndex = position1302, tokenIndex1302
				}
				add(rulews, position1300)
			}
			return true
		l1299:
			position, tokenIndex = position1299, tokenIndex1299
			return false
		},
		/* 87 A <- <('a' / 'A')> */
		nil,
		/* 88 F <- <('f' / 'F')> */
		nil,
		/* 89 B <- <('b' / 'B')> */
		nil,
		/* 90 C <- <('c' / 'C')> */
		nil,
		/* 91 D <- <('d' / 'D')> */
		nil,
		/* 92 E <- <('e' / 'E')> */
		nil,
		/* 93 H <- <('h' / 'H')> */
		nil,
		/* 94 L <- <('l' / 'L')> */
		nil,
		/* 95 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 96 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 97 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 98 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 99 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 100 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 101 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 102 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 103 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 104 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 105 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 106 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 107 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action67)> */
		nil,
		/* 108 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action68)> */
		nil,
		/* 109 decimalByte <- <(<[0-9]+> Action69)> */
		nil,
		/* 110 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action70)> */
		nil,
		/* 111 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action71)> */
		nil,
		/* 112 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position1328, tokenIndex1328 := position, tokenIndex
			{
				position1329 := position
				{
					position1330, tokenIndex1330 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1331
					}
					position++
					goto l1330
				l1331:
					position, tokenIndex = position1330, tokenIndex1330
					{
						position1332, tokenIndex1332 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l1333
						}
						position++
						goto l1332
					l1333:
						position, tokenIndex = position1332, tokenIndex1332
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l1328
						}
						position++
					}
				l1332:
				}
			l1330:
				add(rulehexdigit, position1329)
			}
			return true
		l1328:
			position, tokenIndex = position1328, tokenIndex1328
			return false
		},
		/* 113 octaldigit <- <(<[0-7]> Action72)> */
		func() bool {
			position1334, tokenIndex1334 := position, tokenIndex
			{
				position1335 := position
				{
					position1336 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l1334
					}
					position++
					add(rulePegText, position1336)
				}
				{
					add(ruleAction72, position)
				}
				add(ruleoctaldigit, position1335)
			}
			return true
		l1334:
			position, tokenIndex = position1334, tokenIndex1334
			return false
		},
		/* 114 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action73)> */
		func() bool {
			position1338, tokenIndex1338 := position, tokenIndex
			{
				position1339 := position
				{
					position1340 := position
					{
						position1341, tokenIndex1341 := position, tokenIndex
						{
							position1343, tokenIndex1343 := position, tokenIndex
							if buffer[position] != rune('-') {
								goto l1344
							}
							position++
							goto l1343
						l1344:
							position, tokenIndex = position1343, tokenIndex1343
							if buffer[position] != rune('+') {
								goto l1341
							}
							position++
						}
					l1343:
						goto l1342
					l1341:
						position, tokenIndex = position1341, tokenIndex1341
					}
				l1342:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1338
					}
					position++
				l1345:
					{
						position1346, tokenIndex1346 := position, tokenIndex
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1346
						}
						position++
						goto l1345
					l1346:
						position, tokenIndex = position1346, tokenIndex1346
					}
					add(rulePegText, position1340)
				}
				{
					add(ruleAction73, position)
				}
				add(rulesignedDecimalByte, position1339)
			}
			return true
		l1338:
			position, tokenIndex = position1338, tokenIndex1338
			return false
		},
		/* 115 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position1348, tokenIndex1348 := position, tokenIndex
			{
				position1349 := position
				{
					position1350, tokenIndex1350 := position, tokenIndex
					{
						position1352 := position
						{
							position1353, tokenIndex1353 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l1354
							}
							position++
							goto l1353
						l1354:
							position, tokenIndex = position1353, tokenIndex1353
							if buffer[position] != rune('N') {
								goto l1351
							}
							position++
						}
					l1353:
						{
							position1355, tokenIndex1355 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l1356
							}
							position++
							goto l1355
						l1356:
							position, tokenIndex = position1355, tokenIndex1355
							if buffer[position] != rune('Z') {
								goto l1351
							}
							position++
						}
					l1355:
						{
							add(ruleAction74, position)
						}
						add(ruleFT_NZ, position1352)
					}
					goto l1350
				l1351:
					position, tokenIndex = position1350, tokenIndex1350
					{
						position1359 := position
						{
							position1360, tokenIndex1360 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1361
							}
							position++
							goto l1360
						l1361:
							position, tokenIndex = position1360, tokenIndex1360
							if buffer[position] != rune('P') {
								goto l1358
							}
							position++
						}
					l1360:
						{
							position1362, tokenIndex1362 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l1363
							}
							position++
							goto l1362
						l1363:
							position, tokenIndex = position1362, tokenIndex1362
							if buffer[position] != rune('O') {
								goto l1358
							}
							position++
						}
					l1362:
						{
							add(ruleAction78, position)
						}
						add(ruleFT_PO, position1359)
					}
					goto l1350
				l1358:
					position, tokenIndex = position1350, tokenIndex1350
					{
						position1366 := position
						{
							position1367, tokenIndex1367 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1368
							}
							position++
							goto l1367
						l1368:
							position, tokenIndex = position1367, tokenIndex1367
							if buffer[position] != rune('P') {
								goto l1365
							}
							position++
						}
					l1367:
						{
							position1369, tokenIndex1369 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l1370
							}
							position++
							goto l1369
						l1370:
							position, tokenIndex = position1369, tokenIndex1369
							if buffer[position] != rune('E') {
								goto l1365
							}
							position++
						}
					l1369:
						{
							add(ruleAction79, position)
						}
						add(ruleFT_PE, position1366)
					}
					goto l1350
				l1365:
					position, tokenIndex = position1350, tokenIndex1350
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position1373 := position
								{
									position1374, tokenIndex1374 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l1375
									}
									position++
									goto l1374
								l1375:
									position, tokenIndex = position1374, tokenIndex1374
									if buffer[position] != rune('M') {
										goto l1348
									}
									position++
								}
							l1374:
								{
									add(ruleAction81, position)
								}
								add(ruleFT_M, position1373)
							}
							break
						case 'P', 'p':
							{
								position1377 := position
								{
									position1378, tokenIndex1378 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l1379
									}
									position++
									goto l1378
								l1379:
									position, tokenIndex = position1378, tokenIndex1378
									if buffer[position] != rune('P') {
										goto l1348
									}
									position++
								}
							l1378:
								{
									add(ruleAction80, position)
								}
								add(ruleFT_P, position1377)
							}
							break
						case 'C', 'c':
							{
								position1381 := position
								{
									position1382, tokenIndex1382 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1383
									}
									position++
									goto l1382
								l1383:
									position, tokenIndex = position1382, tokenIndex1382
									if buffer[position] != rune('C') {
										goto l1348
									}
									position++
								}
							l1382:
								{
									add(ruleAction77, position)
								}
								add(ruleFT_C, position1381)
							}
							break
						case 'N', 'n':
							{
								position1385 := position
								{
									position1386, tokenIndex1386 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l1387
									}
									position++
									goto l1386
								l1387:
									position, tokenIndex = position1386, tokenIndex1386
									if buffer[position] != rune('N') {
										goto l1348
									}
									position++
								}
							l1386:
								{
									position1388, tokenIndex1388 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1389
									}
									position++
									goto l1388
								l1389:
									position, tokenIndex = position1388, tokenIndex1388
									if buffer[position] != rune('C') {
										goto l1348
									}
									position++
								}
							l1388:
								{
									add(ruleAction76, position)
								}
								add(ruleFT_NC, position1385)
							}
							break
						default:
							{
								position1391 := position
								{
									position1392, tokenIndex1392 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l1393
									}
									position++
									goto l1392
								l1393:
									position, tokenIndex = position1392, tokenIndex1392
									if buffer[position] != rune('Z') {
										goto l1348
									}
									position++
								}
							l1392:
								{
									add(ruleAction75, position)
								}
								add(ruleFT_Z, position1391)
							}
							break
						}
					}

				}
			l1350:
				add(rulecc, position1349)
			}
			return true
		l1348:
			position, tokenIndex = position1348, tokenIndex1348
			return false
		},
		/* 116 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action74)> */
		nil,
		/* 117 FT_Z <- <(('z' / 'Z') Action75)> */
		nil,
		/* 118 FT_NC <- <(('n' / 'N') ('c' / 'C') Action76)> */
		nil,
		/* 119 FT_C <- <(('c' / 'C') Action77)> */
		nil,
		/* 120 FT_PO <- <(('p' / 'P') ('o' / 'O') Action78)> */
		nil,
		/* 121 FT_PE <- <(('p' / 'P') ('e' / 'E') Action79)> */
		nil,
		/* 122 FT_P <- <(('p' / 'P') Action80)> */
		nil,
		/* 123 FT_M <- <(('m' / 'M') Action81)> */
		nil,
		/* 125 Action0 <- <{ p.Emit() }> */
		nil,
		/* 126 Action1 <- <{ p.LD8() }> */
		nil,
		/* 127 Action2 <- <{ p.LD16() }> */
		nil,
		/* 128 Action3 <- <{ p.Push() }> */
		nil,
		/* 129 Action4 <- <{ p.Pop() }> */
		nil,
		/* 130 Action5 <- <{ p.Ex() }> */
		nil,
		/* 131 Action6 <- <{ p.Inc8() }> */
		nil,
		/* 132 Action7 <- <{ p.Inc8() }> */
		nil,
		/* 133 Action8 <- <{ p.Inc16() }> */
		nil,
		/* 134 Action9 <- <{ p.Dec8() }> */
		nil,
		/* 135 Action10 <- <{ p.Dec8() }> */
		nil,
		/* 136 Action11 <- <{ p.Dec16() }> */
		nil,
		/* 137 Action12 <- <{ p.Add16() }> */
		nil,
		/* 138 Action13 <- <{ p.Dst8() }> */
		nil,
		/* 139 Action14 <- <{ p.Src8() }> */
		nil,
		/* 140 Action15 <- <{ p.Loc8() }> */
		nil,
		/* 141 Action16 <- <{ p.Loc8() }> */
		nil,
		nil,
		/* 143 Action17 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 144 Action18 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 145 Action19 <- <{ p.Dst16() }> */
		nil,
		/* 146 Action20 <- <{ p.Src16() }> */
		nil,
		/* 147 Action21 <- <{ p.Loc16() }> */
		nil,
		/* 148 Action22 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 149 Action23 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 150 Action24 <- <{ p.R16Contents() }> */
		nil,
		/* 151 Action25 <- <{ p.IR16Contents() }> */
		nil,
		/* 152 Action26 <- <{ p.NNContents() }> */
		nil,
		/* 153 Action27 <- <{ p.Accum("ADD") }> */
		nil,
		/* 154 Action28 <- <{ p.Accum("ADC") }> */
		nil,
		/* 155 Action29 <- <{ p.Accum("SUB") }> */
		nil,
		/* 156 Action30 <- <{ p.Accum("SBC") }> */
		nil,
		/* 157 Action31 <- <{ p.Accum("AND") }> */
		nil,
		/* 158 Action32 <- <{ p.Accum("XOR") }> */
		nil,
		/* 159 Action33 <- <{ p.Accum("OR") }> */
		nil,
		/* 160 Action34 <- <{ p.Accum("CP") }> */
		nil,
		/* 161 Action35 <- <{ p.Rot("RLC") }> */
		nil,
		/* 162 Action36 <- <{ p.Rot("RRC") }> */
		nil,
		/* 163 Action37 <- <{ p.Rot("RL") }> */
		nil,
		/* 164 Action38 <- <{ p.Rot("RR") }> */
		nil,
		/* 165 Action39 <- <{ p.Rot("SLA") }> */
		nil,
		/* 166 Action40 <- <{ p.Rot("SRA") }> */
		nil,
		/* 167 Action41 <- <{ p.Rot("SLL") }> */
		nil,
		/* 168 Action42 <- <{ p.Rot("SRL") }> */
		nil,
		/* 169 Action43 <- <{ p.Bit() }> */
		nil,
		/* 170 Action44 <- <{ p.Res() }> */
		nil,
		/* 171 Action45 <- <{ p.Set() }> */
		nil,
		/* 172 Action46 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 173 Action47 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 174 Action48 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 175 Action49 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 176 Action50 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 177 Action51 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 178 Action52 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 179 Action53 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 180 Action54 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 181 Action55 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 182 Action56 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 183 Action57 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 184 Action58 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 185 Action59 <- <{ p.Rst() }> */
		nil,
		/* 186 Action60 <- <{ p.Call() }> */
		nil,
		/* 187 Action61 <- <{ p.Ret() }> */
		nil,
		/* 188 Action62 <- <{ p.Jp() }> */
		nil,
		/* 189 Action63 <- <{ p.Jr() }> */
		nil,
		/* 190 Action64 <- <{ p.Djnz() }> */
		nil,
		/* 191 Action65 <- <{ p.In() }> */
		nil,
		/* 192 Action66 <- <{ p.Out() }> */
		nil,
		/* 193 Action67 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 194 Action68 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 195 Action69 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 196 Action70 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 197 Action71 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 198 Action72 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 199 Action73 <- <{ p.SignedDecimalByte(buffer[begin:end]) }> */
		nil,
		/* 200 Action74 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 201 Action75 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 202 Action76 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 203 Action77 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 204 Action78 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 205 Action79 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 206 Action80 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 207 Action81 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
