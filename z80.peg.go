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
	ruleAlu16
	ruleAdd16
	ruleAdc16
	ruleSbc16
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
	ruleI
	ruleR
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
	ruleAction17
	ruleAction18
	rulePegText
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
	ruleAction82
	ruleAction83
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
	"Alu16",
	"Add16",
	"Adc16",
	"Sbc16",
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
	"I",
	"R",
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
	"Action17",
	"Action18",
	"PegText",
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
	"Action82",
	"Action83",
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
	rules  [215]func() bool
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
			p.Adc16()
		case ruleAction14:
			p.Sbc16()
		case ruleAction15:
			p.Dst8()
		case ruleAction16:
			p.Src8()
		case ruleAction17:
			p.Loc8()
		case ruleAction18:
			p.Loc8()
		case ruleAction19:
			p.R8(buffer[begin:end])
		case ruleAction20:
			p.R8(buffer[begin:end])
		case ruleAction21:
			p.Dst16()
		case ruleAction22:
			p.Src16()
		case ruleAction23:
			p.Loc16()
		case ruleAction24:
			p.R16(buffer[begin:end])
		case ruleAction25:
			p.R16(buffer[begin:end])
		case ruleAction26:
			p.R16Contents()
		case ruleAction27:
			p.IR16Contents()
		case ruleAction28:
			p.NNContents()
		case ruleAction29:
			p.Accum("ADD")
		case ruleAction30:
			p.Accum("ADC")
		case ruleAction31:
			p.Accum("SUB")
		case ruleAction32:
			p.Accum("SBC")
		case ruleAction33:
			p.Accum("AND")
		case ruleAction34:
			p.Accum("XOR")
		case ruleAction35:
			p.Accum("OR")
		case ruleAction36:
			p.Accum("CP")
		case ruleAction37:
			p.Rot("RLC")
		case ruleAction38:
			p.Rot("RRC")
		case ruleAction39:
			p.Rot("RL")
		case ruleAction40:
			p.Rot("RR")
		case ruleAction41:
			p.Rot("SLA")
		case ruleAction42:
			p.Rot("SRA")
		case ruleAction43:
			p.Rot("SLL")
		case ruleAction44:
			p.Rot("SRL")
		case ruleAction45:
			p.Bit()
		case ruleAction46:
			p.Res()
		case ruleAction47:
			p.Set()
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
			p.Simple(buffer[begin:end])
		case ruleAction60:
			p.Simple(buffer[begin:end])
		case ruleAction61:
			p.Rst()
		case ruleAction62:
			p.Call()
		case ruleAction63:
			p.Ret()
		case ruleAction64:
			p.Jp()
		case ruleAction65:
			p.Jr()
		case ruleAction66:
			p.Djnz()
		case ruleAction67:
			p.In()
		case ruleAction68:
			p.Out()
		case ruleAction69:
			p.Nhex(buffer[begin:end])
		case ruleAction70:
			p.Nhex(buffer[begin:end])
		case ruleAction71:
			p.Ndec(buffer[begin:end])
		case ruleAction72:
			p.NNhex(buffer[begin:end])
		case ruleAction73:
			p.NNhex(buffer[begin:end])
		case ruleAction74:
			p.ODigit(buffer[begin:end])
		case ruleAction75:
			p.SignedDecimalByte(buffer[begin:end])
		case ruleAction76:
			p.Conditional(Not{FT_Z})
		case ruleAction77:
			p.Conditional(FT_Z)
		case ruleAction78:
			p.Conditional(Not{FT_C})
		case ruleAction79:
			p.Conditional(FT_C)
		case ruleAction80:
			p.Conditional(FT_PO)
		case ruleAction81:
			p.Conditional(FT_PE)
		case ruleAction82:
			p.Conditional(FT_P)
		case ruleAction83:
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
																	add(ruleAction15, position)
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
										{
											position125 := position
											{
												position126, tokenIndex126 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l127
												}
												position++
												goto l126
											l127:
												position, tokenIndex = position126, tokenIndex126
												if buffer[position] != rune('A') {
													goto l124
												}
												position++
											}
										l126:
											{
												position128, tokenIndex128 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l129
												}
												position++
												goto l128
											l129:
												position, tokenIndex = position128, tokenIndex128
												if buffer[position] != rune('D') {
													goto l124
												}
												position++
											}
										l128:
											{
												position130, tokenIndex130 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l131
												}
												position++
												goto l130
											l131:
												position, tokenIndex = position130, tokenIndex130
												if buffer[position] != rune('D') {
													goto l124
												}
												position++
											}
										l130:
											if !_rules[rulews]() {
												goto l124
											}
											if !_rules[ruleDst16]() {
												goto l124
											}
											if !_rules[rulesep]() {
												goto l124
											}
											if !_rules[ruleSrc16]() {
												goto l124
											}
											{
												add(ruleAction12, position)
											}
											add(ruleAdd16, position125)
										}
										goto l123
									l124:
										position, tokenIndex = position123, tokenIndex123
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
												if buffer[position] != rune('c') {
													goto l140
												}
												position++
												goto l139
											l140:
												position, tokenIndex = position139, tokenIndex139
												if buffer[position] != rune('C') {
													goto l133
												}
												position++
											}
										l139:
											if !_rules[rulews]() {
												goto l133
											}
											if !_rules[ruleDst16]() {
												goto l133
											}
											if !_rules[rulesep]() {
												goto l133
											}
											if !_rules[ruleSrc16]() {
												goto l133
											}
											{
												add(ruleAction13, position)
											}
											add(ruleAdc16, position134)
										}
										goto l123
									l133:
										position, tokenIndex = position123, tokenIndex123
										{
											position142 := position
											{
												position143, tokenIndex143 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l144
												}
												position++
												goto l143
											l144:
												position, tokenIndex = position143, tokenIndex143
												if buffer[position] != rune('S') {
													goto l121
												}
												position++
											}
										l143:
											{
												position145, tokenIndex145 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l146
												}
												position++
												goto l145
											l146:
												position, tokenIndex = position145, tokenIndex145
												if buffer[position] != rune('B') {
													goto l121
												}
												position++
											}
										l145:
											{
												position147, tokenIndex147 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l148
												}
												position++
												goto l147
											l148:
												position, tokenIndex = position147, tokenIndex147
												if buffer[position] != rune('C') {
													goto l121
												}
												position++
											}
										l147:
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
												add(ruleAction14, position)
											}
											add(ruleSbc16, position142)
										}
									}
								l123:
									add(ruleAlu16, position122)
								}
								goto l13
							l121:
								position, tokenIndex = position13, tokenIndex13
								{
									position151 := position
									{
										position152, tokenIndex152 := position, tokenIndex
										{
											position154 := position
											{
												position155, tokenIndex155 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l156
												}
												position++
												goto l155
											l156:
												position, tokenIndex = position155, tokenIndex155
												if buffer[position] != rune('A') {
													goto l153
												}
												position++
											}
										l155:
											{
												position157, tokenIndex157 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l158
												}
												position++
												goto l157
											l158:
												position, tokenIndex = position157, tokenIndex157
												if buffer[position] != rune('D') {
													goto l153
												}
												position++
											}
										l157:
											{
												position159, tokenIndex159 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l160
												}
												position++
												goto l159
											l160:
												position, tokenIndex = position159, tokenIndex159
												if buffer[position] != rune('D') {
													goto l153
												}
												position++
											}
										l159:
											if !_rules[rulews]() {
												goto l153
											}
											{
												position161, tokenIndex161 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l162
												}
												position++
												goto l161
											l162:
												position, tokenIndex = position161, tokenIndex161
												if buffer[position] != rune('A') {
													goto l153
												}
												position++
											}
										l161:
											if !_rules[rulesep]() {
												goto l153
											}
											if !_rules[ruleSrc8]() {
												goto l153
											}
											{
												add(ruleAction29, position)
											}
											add(ruleAdd, position154)
										}
										goto l152
									l153:
										position, tokenIndex = position152, tokenIndex152
										{
											position165 := position
											{
												position166, tokenIndex166 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l167
												}
												position++
												goto l166
											l167:
												position, tokenIndex = position166, tokenIndex166
												if buffer[position] != rune('A') {
													goto l164
												}
												position++
											}
										l166:
											{
												position168, tokenIndex168 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l169
												}
												position++
												goto l168
											l169:
												position, tokenIndex = position168, tokenIndex168
												if buffer[position] != rune('D') {
													goto l164
												}
												position++
											}
										l168:
											{
												position170, tokenIndex170 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l171
												}
												position++
												goto l170
											l171:
												position, tokenIndex = position170, tokenIndex170
												if buffer[position] != rune('C') {
													goto l164
												}
												position++
											}
										l170:
											if !_rules[rulews]() {
												goto l164
											}
											{
												position172, tokenIndex172 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l173
												}
												position++
												goto l172
											l173:
												position, tokenIndex = position172, tokenIndex172
												if buffer[position] != rune('A') {
													goto l164
												}
												position++
											}
										l172:
											if !_rules[rulesep]() {
												goto l164
											}
											if !_rules[ruleSrc8]() {
												goto l164
											}
											{
												add(ruleAction30, position)
											}
											add(ruleAdc, position165)
										}
										goto l152
									l164:
										position, tokenIndex = position152, tokenIndex152
										{
											position176 := position
											{
												position177, tokenIndex177 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l178
												}
												position++
												goto l177
											l178:
												position, tokenIndex = position177, tokenIndex177
												if buffer[position] != rune('S') {
													goto l175
												}
												position++
											}
										l177:
											{
												position179, tokenIndex179 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l180
												}
												position++
												goto l179
											l180:
												position, tokenIndex = position179, tokenIndex179
												if buffer[position] != rune('U') {
													goto l175
												}
												position++
											}
										l179:
											{
												position181, tokenIndex181 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l182
												}
												position++
												goto l181
											l182:
												position, tokenIndex = position181, tokenIndex181
												if buffer[position] != rune('B') {
													goto l175
												}
												position++
											}
										l181:
											if !_rules[rulews]() {
												goto l175
											}
											if !_rules[ruleSrc8]() {
												goto l175
											}
											{
												add(ruleAction31, position)
											}
											add(ruleSub, position176)
										}
										goto l152
									l175:
										position, tokenIndex = position152, tokenIndex152
										{
											switch buffer[position] {
											case 'C', 'c':
												{
													position185 := position
													{
														position186, tokenIndex186 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l187
														}
														position++
														goto l186
													l187:
														position, tokenIndex = position186, tokenIndex186
														if buffer[position] != rune('C') {
															goto l150
														}
														position++
													}
												l186:
													{
														position188, tokenIndex188 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l189
														}
														position++
														goto l188
													l189:
														position, tokenIndex = position188, tokenIndex188
														if buffer[position] != rune('P') {
															goto l150
														}
														position++
													}
												l188:
													if !_rules[rulews]() {
														goto l150
													}
													if !_rules[ruleSrc8]() {
														goto l150
													}
													{
														add(ruleAction36, position)
													}
													add(ruleCp, position185)
												}
												break
											case 'O', 'o':
												{
													position191 := position
													{
														position192, tokenIndex192 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l193
														}
														position++
														goto l192
													l193:
														position, tokenIndex = position192, tokenIndex192
														if buffer[position] != rune('O') {
															goto l150
														}
														position++
													}
												l192:
													{
														position194, tokenIndex194 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l195
														}
														position++
														goto l194
													l195:
														position, tokenIndex = position194, tokenIndex194
														if buffer[position] != rune('R') {
															goto l150
														}
														position++
													}
												l194:
													if !_rules[rulews]() {
														goto l150
													}
													if !_rules[ruleSrc8]() {
														goto l150
													}
													{
														add(ruleAction35, position)
													}
													add(ruleOr, position191)
												}
												break
											case 'X', 'x':
												{
													position197 := position
													{
														position198, tokenIndex198 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l199
														}
														position++
														goto l198
													l199:
														position, tokenIndex = position198, tokenIndex198
														if buffer[position] != rune('X') {
															goto l150
														}
														position++
													}
												l198:
													{
														position200, tokenIndex200 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l201
														}
														position++
														goto l200
													l201:
														position, tokenIndex = position200, tokenIndex200
														if buffer[position] != rune('O') {
															goto l150
														}
														position++
													}
												l200:
													{
														position202, tokenIndex202 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l203
														}
														position++
														goto l202
													l203:
														position, tokenIndex = position202, tokenIndex202
														if buffer[position] != rune('R') {
															goto l150
														}
														position++
													}
												l202:
													if !_rules[rulews]() {
														goto l150
													}
													if !_rules[ruleSrc8]() {
														goto l150
													}
													{
														add(ruleAction34, position)
													}
													add(ruleXor, position197)
												}
												break
											case 'A', 'a':
												{
													position205 := position
													{
														position206, tokenIndex206 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l207
														}
														position++
														goto l206
													l207:
														position, tokenIndex = position206, tokenIndex206
														if buffer[position] != rune('A') {
															goto l150
														}
														position++
													}
												l206:
													{
														position208, tokenIndex208 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l209
														}
														position++
														goto l208
													l209:
														position, tokenIndex = position208, tokenIndex208
														if buffer[position] != rune('N') {
															goto l150
														}
														position++
													}
												l208:
													{
														position210, tokenIndex210 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l211
														}
														position++
														goto l210
													l211:
														position, tokenIndex = position210, tokenIndex210
														if buffer[position] != rune('D') {
															goto l150
														}
														position++
													}
												l210:
													if !_rules[rulews]() {
														goto l150
													}
													if !_rules[ruleSrc8]() {
														goto l150
													}
													{
														add(ruleAction33, position)
													}
													add(ruleAnd, position205)
												}
												break
											default:
												{
													position213 := position
													{
														position214, tokenIndex214 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l215
														}
														position++
														goto l214
													l215:
														position, tokenIndex = position214, tokenIndex214
														if buffer[position] != rune('S') {
															goto l150
														}
														position++
													}
												l214:
													{
														position216, tokenIndex216 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l217
														}
														position++
														goto l216
													l217:
														position, tokenIndex = position216, tokenIndex216
														if buffer[position] != rune('B') {
															goto l150
														}
														position++
													}
												l216:
													{
														position218, tokenIndex218 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l219
														}
														position++
														goto l218
													l219:
														position, tokenIndex = position218, tokenIndex218
														if buffer[position] != rune('C') {
															goto l150
														}
														position++
													}
												l218:
													if !_rules[rulews]() {
														goto l150
													}
													{
														position220, tokenIndex220 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l221
														}
														position++
														goto l220
													l221:
														position, tokenIndex = position220, tokenIndex220
														if buffer[position] != rune('A') {
															goto l150
														}
														position++
													}
												l220:
													if !_rules[rulesep]() {
														goto l150
													}
													if !_rules[ruleSrc8]() {
														goto l150
													}
													{
														add(ruleAction32, position)
													}
													add(ruleSbc, position213)
												}
												break
											}
										}

									}
								l152:
									add(ruleAlu, position151)
								}
								goto l13
							l150:
								position, tokenIndex = position13, tokenIndex13
								{
									position224 := position
									{
										position225, tokenIndex225 := position, tokenIndex
										{
											position227 := position
											{
												position228, tokenIndex228 := position, tokenIndex
												{
													position230 := position
													{
														position231, tokenIndex231 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l232
														}
														position++
														goto l231
													l232:
														position, tokenIndex = position231, tokenIndex231
														if buffer[position] != rune('R') {
															goto l229
														}
														position++
													}
												l231:
													{
														position233, tokenIndex233 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l234
														}
														position++
														goto l233
													l234:
														position, tokenIndex = position233, tokenIndex233
														if buffer[position] != rune('L') {
															goto l229
														}
														position++
													}
												l233:
													{
														position235, tokenIndex235 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l236
														}
														position++
														goto l235
													l236:
														position, tokenIndex = position235, tokenIndex235
														if buffer[position] != rune('C') {
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
														add(ruleAction37, position)
													}
													add(ruleRlc, position230)
												}
												goto l228
											l229:
												position, tokenIndex = position228, tokenIndex228
												{
													position239 := position
													{
														position240, tokenIndex240 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l241
														}
														position++
														goto l240
													l241:
														position, tokenIndex = position240, tokenIndex240
														if buffer[position] != rune('R') {
															goto l238
														}
														position++
													}
												l240:
													{
														position242, tokenIndex242 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l243
														}
														position++
														goto l242
													l243:
														position, tokenIndex = position242, tokenIndex242
														if buffer[position] != rune('R') {
															goto l238
														}
														position++
													}
												l242:
													{
														position244, tokenIndex244 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l245
														}
														position++
														goto l244
													l245:
														position, tokenIndex = position244, tokenIndex244
														if buffer[position] != rune('C') {
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
														add(ruleAction38, position)
													}
													add(ruleRrc, position239)
												}
												goto l228
											l238:
												position, tokenIndex = position228, tokenIndex228
												{
													position248 := position
													{
														position249, tokenIndex249 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l250
														}
														position++
														goto l249
													l250:
														position, tokenIndex = position249, tokenIndex249
														if buffer[position] != rune('R') {
															goto l247
														}
														position++
													}
												l249:
													{
														position251, tokenIndex251 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l252
														}
														position++
														goto l251
													l252:
														position, tokenIndex = position251, tokenIndex251
														if buffer[position] != rune('L') {
															goto l247
														}
														position++
													}
												l251:
													if !_rules[rulews]() {
														goto l247
													}
													if !_rules[ruleLoc8]() {
														goto l247
													}
													{
														add(ruleAction39, position)
													}
													add(ruleRl, position248)
												}
												goto l228
											l247:
												position, tokenIndex = position228, tokenIndex228
												{
													position255 := position
													{
														position256, tokenIndex256 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l257
														}
														position++
														goto l256
													l257:
														position, tokenIndex = position256, tokenIndex256
														if buffer[position] != rune('R') {
															goto l254
														}
														position++
													}
												l256:
													{
														position258, tokenIndex258 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l259
														}
														position++
														goto l258
													l259:
														position, tokenIndex = position258, tokenIndex258
														if buffer[position] != rune('R') {
															goto l254
														}
														position++
													}
												l258:
													if !_rules[rulews]() {
														goto l254
													}
													if !_rules[ruleLoc8]() {
														goto l254
													}
													{
														add(ruleAction40, position)
													}
													add(ruleRr, position255)
												}
												goto l228
											l254:
												position, tokenIndex = position228, tokenIndex228
												{
													position262 := position
													{
														position263, tokenIndex263 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l264
														}
														position++
														goto l263
													l264:
														position, tokenIndex = position263, tokenIndex263
														if buffer[position] != rune('S') {
															goto l261
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
															goto l261
														}
														position++
													}
												l265:
													{
														position267, tokenIndex267 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l268
														}
														position++
														goto l267
													l268:
														position, tokenIndex = position267, tokenIndex267
														if buffer[position] != rune('A') {
															goto l261
														}
														position++
													}
												l267:
													if !_rules[rulews]() {
														goto l261
													}
													if !_rules[ruleLoc8]() {
														goto l261
													}
													{
														add(ruleAction41, position)
													}
													add(ruleSla, position262)
												}
												goto l228
											l261:
												position, tokenIndex = position228, tokenIndex228
												{
													position271 := position
													{
														position272, tokenIndex272 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l273
														}
														position++
														goto l272
													l273:
														position, tokenIndex = position272, tokenIndex272
														if buffer[position] != rune('S') {
															goto l270
														}
														position++
													}
												l272:
													{
														position274, tokenIndex274 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l275
														}
														position++
														goto l274
													l275:
														position, tokenIndex = position274, tokenIndex274
														if buffer[position] != rune('R') {
															goto l270
														}
														position++
													}
												l274:
													{
														position276, tokenIndex276 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l277
														}
														position++
														goto l276
													l277:
														position, tokenIndex = position276, tokenIndex276
														if buffer[position] != rune('A') {
															goto l270
														}
														position++
													}
												l276:
													if !_rules[rulews]() {
														goto l270
													}
													if !_rules[ruleLoc8]() {
														goto l270
													}
													{
														add(ruleAction42, position)
													}
													add(ruleSra, position271)
												}
												goto l228
											l270:
												position, tokenIndex = position228, tokenIndex228
												{
													position280 := position
													{
														position281, tokenIndex281 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l282
														}
														position++
														goto l281
													l282:
														position, tokenIndex = position281, tokenIndex281
														if buffer[position] != rune('S') {
															goto l279
														}
														position++
													}
												l281:
													{
														position283, tokenIndex283 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l284
														}
														position++
														goto l283
													l284:
														position, tokenIndex = position283, tokenIndex283
														if buffer[position] != rune('L') {
															goto l279
														}
														position++
													}
												l283:
													{
														position285, tokenIndex285 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l286
														}
														position++
														goto l285
													l286:
														position, tokenIndex = position285, tokenIndex285
														if buffer[position] != rune('L') {
															goto l279
														}
														position++
													}
												l285:
													if !_rules[rulews]() {
														goto l279
													}
													if !_rules[ruleLoc8]() {
														goto l279
													}
													{
														add(ruleAction43, position)
													}
													add(ruleSll, position280)
												}
												goto l228
											l279:
												position, tokenIndex = position228, tokenIndex228
												{
													position288 := position
													{
														position289, tokenIndex289 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l290
														}
														position++
														goto l289
													l290:
														position, tokenIndex = position289, tokenIndex289
														if buffer[position] != rune('S') {
															goto l226
														}
														position++
													}
												l289:
													{
														position291, tokenIndex291 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l292
														}
														position++
														goto l291
													l292:
														position, tokenIndex = position291, tokenIndex291
														if buffer[position] != rune('R') {
															goto l226
														}
														position++
													}
												l291:
													{
														position293, tokenIndex293 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l294
														}
														position++
														goto l293
													l294:
														position, tokenIndex = position293, tokenIndex293
														if buffer[position] != rune('L') {
															goto l226
														}
														position++
													}
												l293:
													if !_rules[rulews]() {
														goto l226
													}
													if !_rules[ruleLoc8]() {
														goto l226
													}
													{
														add(ruleAction44, position)
													}
													add(ruleSrl, position288)
												}
											}
										l228:
											add(ruleRot, position227)
										}
										goto l225
									l226:
										position, tokenIndex = position225, tokenIndex225
										{
											switch buffer[position] {
											case 'S', 's':
												{
													position297 := position
													{
														position298, tokenIndex298 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l299
														}
														position++
														goto l298
													l299:
														position, tokenIndex = position298, tokenIndex298
														if buffer[position] != rune('S') {
															goto l223
														}
														position++
													}
												l298:
													{
														position300, tokenIndex300 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l301
														}
														position++
														goto l300
													l301:
														position, tokenIndex = position300, tokenIndex300
														if buffer[position] != rune('E') {
															goto l223
														}
														position++
													}
												l300:
													{
														position302, tokenIndex302 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l303
														}
														position++
														goto l302
													l303:
														position, tokenIndex = position302, tokenIndex302
														if buffer[position] != rune('T') {
															goto l223
														}
														position++
													}
												l302:
													if !_rules[rulews]() {
														goto l223
													}
													if !_rules[ruleoctaldigit]() {
														goto l223
													}
													if !_rules[rulesep]() {
														goto l223
													}
													if !_rules[ruleLoc8]() {
														goto l223
													}
													{
														add(ruleAction47, position)
													}
													add(ruleSet, position297)
												}
												break
											case 'R', 'r':
												{
													position305 := position
													{
														position306, tokenIndex306 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l307
														}
														position++
														goto l306
													l307:
														position, tokenIndex = position306, tokenIndex306
														if buffer[position] != rune('R') {
															goto l223
														}
														position++
													}
												l306:
													{
														position308, tokenIndex308 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l309
														}
														position++
														goto l308
													l309:
														position, tokenIndex = position308, tokenIndex308
														if buffer[position] != rune('E') {
															goto l223
														}
														position++
													}
												l308:
													{
														position310, tokenIndex310 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l311
														}
														position++
														goto l310
													l311:
														position, tokenIndex = position310, tokenIndex310
														if buffer[position] != rune('S') {
															goto l223
														}
														position++
													}
												l310:
													if !_rules[rulews]() {
														goto l223
													}
													if !_rules[ruleoctaldigit]() {
														goto l223
													}
													if !_rules[rulesep]() {
														goto l223
													}
													if !_rules[ruleLoc8]() {
														goto l223
													}
													{
														add(ruleAction46, position)
													}
													add(ruleRes, position305)
												}
												break
											default:
												{
													position313 := position
													{
														position314, tokenIndex314 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l315
														}
														position++
														goto l314
													l315:
														position, tokenIndex = position314, tokenIndex314
														if buffer[position] != rune('B') {
															goto l223
														}
														position++
													}
												l314:
													{
														position316, tokenIndex316 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l317
														}
														position++
														goto l316
													l317:
														position, tokenIndex = position316, tokenIndex316
														if buffer[position] != rune('I') {
															goto l223
														}
														position++
													}
												l316:
													{
														position318, tokenIndex318 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l319
														}
														position++
														goto l318
													l319:
														position, tokenIndex = position318, tokenIndex318
														if buffer[position] != rune('T') {
															goto l223
														}
														position++
													}
												l318:
													if !_rules[rulews]() {
														goto l223
													}
													if !_rules[ruleoctaldigit]() {
														goto l223
													}
													if !_rules[rulesep]() {
														goto l223
													}
													if !_rules[ruleLoc8]() {
														goto l223
													}
													{
														add(ruleAction45, position)
													}
													add(ruleBit, position313)
												}
												break
											}
										}

									}
								l225:
									add(ruleBitOp, position224)
								}
								goto l13
							l223:
								position, tokenIndex = position13, tokenIndex13
								{
									position322 := position
									{
										position323, tokenIndex323 := position, tokenIndex
										{
											position325 := position
											{
												position326 := position
												{
													position327, tokenIndex327 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l328
													}
													position++
													goto l327
												l328:
													position, tokenIndex = position327, tokenIndex327
													if buffer[position] != rune('R') {
														goto l324
													}
													position++
												}
											l327:
												{
													position329, tokenIndex329 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l330
													}
													position++
													goto l329
												l330:
													position, tokenIndex = position329, tokenIndex329
													if buffer[position] != rune('L') {
														goto l324
													}
													position++
												}
											l329:
												{
													position331, tokenIndex331 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l332
													}
													position++
													goto l331
												l332:
													position, tokenIndex = position331, tokenIndex331
													if buffer[position] != rune('C') {
														goto l324
													}
													position++
												}
											l331:
												{
													position333, tokenIndex333 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l334
													}
													position++
													goto l333
												l334:
													position, tokenIndex = position333, tokenIndex333
													if buffer[position] != rune('A') {
														goto l324
													}
													position++
												}
											l333:
												add(rulePegText, position326)
											}
											{
												add(ruleAction50, position)
											}
											add(ruleRlca, position325)
										}
										goto l323
									l324:
										position, tokenIndex = position323, tokenIndex323
										{
											position337 := position
											{
												position338 := position
												{
													position339, tokenIndex339 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l340
													}
													position++
													goto l339
												l340:
													position, tokenIndex = position339, tokenIndex339
													if buffer[position] != rune('R') {
														goto l336
													}
													position++
												}
											l339:
												{
													position341, tokenIndex341 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l342
													}
													position++
													goto l341
												l342:
													position, tokenIndex = position341, tokenIndex341
													if buffer[position] != rune('R') {
														goto l336
													}
													position++
												}
											l341:
												{
													position343, tokenIndex343 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l344
													}
													position++
													goto l343
												l344:
													position, tokenIndex = position343, tokenIndex343
													if buffer[position] != rune('C') {
														goto l336
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
														goto l336
													}
													position++
												}
											l345:
												add(rulePegText, position338)
											}
											{
												add(ruleAction51, position)
											}
											add(ruleRrca, position337)
										}
										goto l323
									l336:
										position, tokenIndex = position323, tokenIndex323
										{
											position349 := position
											{
												position350 := position
												{
													position351, tokenIndex351 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l352
													}
													position++
													goto l351
												l352:
													position, tokenIndex = position351, tokenIndex351
													if buffer[position] != rune('R') {
														goto l348
													}
													position++
												}
											l351:
												{
													position353, tokenIndex353 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l354
													}
													position++
													goto l353
												l354:
													position, tokenIndex = position353, tokenIndex353
													if buffer[position] != rune('L') {
														goto l348
													}
													position++
												}
											l353:
												{
													position355, tokenIndex355 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l356
													}
													position++
													goto l355
												l356:
													position, tokenIndex = position355, tokenIndex355
													if buffer[position] != rune('A') {
														goto l348
													}
													position++
												}
											l355:
												add(rulePegText, position350)
											}
											{
												add(ruleAction52, position)
											}
											add(ruleRla, position349)
										}
										goto l323
									l348:
										position, tokenIndex = position323, tokenIndex323
										{
											position359 := position
											{
												position360 := position
												{
													position361, tokenIndex361 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l362
													}
													position++
													goto l361
												l362:
													position, tokenIndex = position361, tokenIndex361
													if buffer[position] != rune('D') {
														goto l358
													}
													position++
												}
											l361:
												{
													position363, tokenIndex363 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l364
													}
													position++
													goto l363
												l364:
													position, tokenIndex = position363, tokenIndex363
													if buffer[position] != rune('A') {
														goto l358
													}
													position++
												}
											l363:
												{
													position365, tokenIndex365 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l366
													}
													position++
													goto l365
												l366:
													position, tokenIndex = position365, tokenIndex365
													if buffer[position] != rune('A') {
														goto l358
													}
													position++
												}
											l365:
												add(rulePegText, position360)
											}
											{
												add(ruleAction54, position)
											}
											add(ruleDaa, position359)
										}
										goto l323
									l358:
										position, tokenIndex = position323, tokenIndex323
										{
											position369 := position
											{
												position370 := position
												{
													position371, tokenIndex371 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l372
													}
													position++
													goto l371
												l372:
													position, tokenIndex = position371, tokenIndex371
													if buffer[position] != rune('C') {
														goto l368
													}
													position++
												}
											l371:
												{
													position373, tokenIndex373 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l374
													}
													position++
													goto l373
												l374:
													position, tokenIndex = position373, tokenIndex373
													if buffer[position] != rune('P') {
														goto l368
													}
													position++
												}
											l373:
												{
													position375, tokenIndex375 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l376
													}
													position++
													goto l375
												l376:
													position, tokenIndex = position375, tokenIndex375
													if buffer[position] != rune('L') {
														goto l368
													}
													position++
												}
											l375:
												add(rulePegText, position370)
											}
											{
												add(ruleAction55, position)
											}
											add(ruleCpl, position369)
										}
										goto l323
									l368:
										position, tokenIndex = position323, tokenIndex323
										{
											position379 := position
											{
												position380 := position
												{
													position381, tokenIndex381 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l382
													}
													position++
													goto l381
												l382:
													position, tokenIndex = position381, tokenIndex381
													if buffer[position] != rune('E') {
														goto l378
													}
													position++
												}
											l381:
												{
													position383, tokenIndex383 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l384
													}
													position++
													goto l383
												l384:
													position, tokenIndex = position383, tokenIndex383
													if buffer[position] != rune('X') {
														goto l378
													}
													position++
												}
											l383:
												{
													position385, tokenIndex385 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l386
													}
													position++
													goto l385
												l386:
													position, tokenIndex = position385, tokenIndex385
													if buffer[position] != rune('X') {
														goto l378
													}
													position++
												}
											l385:
												add(rulePegText, position380)
											}
											{
												add(ruleAction58, position)
											}
											add(ruleExx, position379)
										}
										goto l323
									l378:
										position, tokenIndex = position323, tokenIndex323
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position389 := position
													{
														position390 := position
														{
															position391, tokenIndex391 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l392
															}
															position++
															goto l391
														l392:
															position, tokenIndex = position391, tokenIndex391
															if buffer[position] != rune('E') {
																goto l321
															}
															position++
														}
													l391:
														{
															position393, tokenIndex393 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l394
															}
															position++
															goto l393
														l394:
															position, tokenIndex = position393, tokenIndex393
															if buffer[position] != rune('I') {
																goto l321
															}
															position++
														}
													l393:
														add(rulePegText, position390)
													}
													{
														add(ruleAction60, position)
													}
													add(ruleEi, position389)
												}
												break
											case 'D', 'd':
												{
													position396 := position
													{
														position397 := position
														{
															position398, tokenIndex398 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l399
															}
															position++
															goto l398
														l399:
															position, tokenIndex = position398, tokenIndex398
															if buffer[position] != rune('D') {
																goto l321
															}
															position++
														}
													l398:
														{
															position400, tokenIndex400 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l401
															}
															position++
															goto l400
														l401:
															position, tokenIndex = position400, tokenIndex400
															if buffer[position] != rune('I') {
																goto l321
															}
															position++
														}
													l400:
														add(rulePegText, position397)
													}
													{
														add(ruleAction59, position)
													}
													add(ruleDi, position396)
												}
												break
											case 'C', 'c':
												{
													position403 := position
													{
														position404 := position
														{
															position405, tokenIndex405 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l406
															}
															position++
															goto l405
														l406:
															position, tokenIndex = position405, tokenIndex405
															if buffer[position] != rune('C') {
																goto l321
															}
															position++
														}
													l405:
														{
															position407, tokenIndex407 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l408
															}
															position++
															goto l407
														l408:
															position, tokenIndex = position407, tokenIndex407
															if buffer[position] != rune('C') {
																goto l321
															}
															position++
														}
													l407:
														{
															position409, tokenIndex409 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l410
															}
															position++
															goto l409
														l410:
															position, tokenIndex = position409, tokenIndex409
															if buffer[position] != rune('F') {
																goto l321
															}
															position++
														}
													l409:
														add(rulePegText, position404)
													}
													{
														add(ruleAction57, position)
													}
													add(ruleCcf, position403)
												}
												break
											case 'S', 's':
												{
													position412 := position
													{
														position413 := position
														{
															position414, tokenIndex414 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l415
															}
															position++
															goto l414
														l415:
															position, tokenIndex = position414, tokenIndex414
															if buffer[position] != rune('S') {
																goto l321
															}
															position++
														}
													l414:
														{
															position416, tokenIndex416 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l417
															}
															position++
															goto l416
														l417:
															position, tokenIndex = position416, tokenIndex416
															if buffer[position] != rune('C') {
																goto l321
															}
															position++
														}
													l416:
														{
															position418, tokenIndex418 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l419
															}
															position++
															goto l418
														l419:
															position, tokenIndex = position418, tokenIndex418
															if buffer[position] != rune('F') {
																goto l321
															}
															position++
														}
													l418:
														add(rulePegText, position413)
													}
													{
														add(ruleAction56, position)
													}
													add(ruleScf, position412)
												}
												break
											case 'R', 'r':
												{
													position421 := position
													{
														position422 := position
														{
															position423, tokenIndex423 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l424
															}
															position++
															goto l423
														l424:
															position, tokenIndex = position423, tokenIndex423
															if buffer[position] != rune('R') {
																goto l321
															}
															position++
														}
													l423:
														{
															position425, tokenIndex425 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l426
															}
															position++
															goto l425
														l426:
															position, tokenIndex = position425, tokenIndex425
															if buffer[position] != rune('R') {
																goto l321
															}
															position++
														}
													l425:
														{
															position427, tokenIndex427 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l428
															}
															position++
															goto l427
														l428:
															position, tokenIndex = position427, tokenIndex427
															if buffer[position] != rune('A') {
																goto l321
															}
															position++
														}
													l427:
														add(rulePegText, position422)
													}
													{
														add(ruleAction53, position)
													}
													add(ruleRra, position421)
												}
												break
											case 'H', 'h':
												{
													position430 := position
													{
														position431 := position
														{
															position432, tokenIndex432 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l433
															}
															position++
															goto l432
														l433:
															position, tokenIndex = position432, tokenIndex432
															if buffer[position] != rune('H') {
																goto l321
															}
															position++
														}
													l432:
														{
															position434, tokenIndex434 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l435
															}
															position++
															goto l434
														l435:
															position, tokenIndex = position434, tokenIndex434
															if buffer[position] != rune('A') {
																goto l321
															}
															position++
														}
													l434:
														{
															position436, tokenIndex436 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l437
															}
															position++
															goto l436
														l437:
															position, tokenIndex = position436, tokenIndex436
															if buffer[position] != rune('L') {
																goto l321
															}
															position++
														}
													l436:
														{
															position438, tokenIndex438 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l439
															}
															position++
															goto l438
														l439:
															position, tokenIndex = position438, tokenIndex438
															if buffer[position] != rune('T') {
																goto l321
															}
															position++
														}
													l438:
														add(rulePegText, position431)
													}
													{
														add(ruleAction49, position)
													}
													add(ruleHalt, position430)
												}
												break
											default:
												{
													position441 := position
													{
														position442 := position
														{
															position443, tokenIndex443 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l444
															}
															position++
															goto l443
														l444:
															position, tokenIndex = position443, tokenIndex443
															if buffer[position] != rune('N') {
																goto l321
															}
															position++
														}
													l443:
														{
															position445, tokenIndex445 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l446
															}
															position++
															goto l445
														l446:
															position, tokenIndex = position445, tokenIndex445
															if buffer[position] != rune('O') {
																goto l321
															}
															position++
														}
													l445:
														{
															position447, tokenIndex447 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l448
															}
															position++
															goto l447
														l448:
															position, tokenIndex = position447, tokenIndex447
															if buffer[position] != rune('P') {
																goto l321
															}
															position++
														}
													l447:
														add(rulePegText, position442)
													}
													{
														add(ruleAction48, position)
													}
													add(ruleNop, position441)
												}
												break
											}
										}

									}
								l323:
									add(ruleSimple, position322)
								}
								goto l13
							l321:
								position, tokenIndex = position13, tokenIndex13
								{
									position451 := position
									{
										position452, tokenIndex452 := position, tokenIndex
										{
											position454 := position
											{
												position455, tokenIndex455 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l456
												}
												position++
												goto l455
											l456:
												position, tokenIndex = position455, tokenIndex455
												if buffer[position] != rune('R') {
													goto l453
												}
												position++
											}
										l455:
											{
												position457, tokenIndex457 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l458
												}
												position++
												goto l457
											l458:
												position, tokenIndex = position457, tokenIndex457
												if buffer[position] != rune('S') {
													goto l453
												}
												position++
											}
										l457:
											{
												position459, tokenIndex459 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l460
												}
												position++
												goto l459
											l460:
												position, tokenIndex = position459, tokenIndex459
												if buffer[position] != rune('T') {
													goto l453
												}
												position++
											}
										l459:
											if !_rules[rulews]() {
												goto l453
											}
											if !_rules[rulen]() {
												goto l453
											}
											{
												add(ruleAction61, position)
											}
											add(ruleRst, position454)
										}
										goto l452
									l453:
										position, tokenIndex = position452, tokenIndex452
										{
											position463 := position
											{
												position464, tokenIndex464 := position, tokenIndex
												if buffer[position] != rune('j') {
													goto l465
												}
												position++
												goto l464
											l465:
												position, tokenIndex = position464, tokenIndex464
												if buffer[position] != rune('J') {
													goto l462
												}
												position++
											}
										l464:
											{
												position466, tokenIndex466 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l467
												}
												position++
												goto l466
											l467:
												position, tokenIndex = position466, tokenIndex466
												if buffer[position] != rune('P') {
													goto l462
												}
												position++
											}
										l466:
											if !_rules[rulews]() {
												goto l462
											}
											{
												position468, tokenIndex468 := position, tokenIndex
												if !_rules[rulecc]() {
													goto l468
												}
												if !_rules[rulesep]() {
													goto l468
												}
												goto l469
											l468:
												position, tokenIndex = position468, tokenIndex468
											}
										l469:
											if !_rules[ruleSrc16]() {
												goto l462
											}
											{
												add(ruleAction64, position)
											}
											add(ruleJp, position463)
										}
										goto l452
									l462:
										position, tokenIndex = position452, tokenIndex452
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position472 := position
													{
														position473, tokenIndex473 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l474
														}
														position++
														goto l473
													l474:
														position, tokenIndex = position473, tokenIndex473
														if buffer[position] != rune('D') {
															goto l450
														}
														position++
													}
												l473:
													{
														position475, tokenIndex475 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l476
														}
														position++
														goto l475
													l476:
														position, tokenIndex = position475, tokenIndex475
														if buffer[position] != rune('J') {
															goto l450
														}
														position++
													}
												l475:
													{
														position477, tokenIndex477 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l478
														}
														position++
														goto l477
													l478:
														position, tokenIndex = position477, tokenIndex477
														if buffer[position] != rune('N') {
															goto l450
														}
														position++
													}
												l477:
													{
														position479, tokenIndex479 := position, tokenIndex
														if buffer[position] != rune('z') {
															goto l480
														}
														position++
														goto l479
													l480:
														position, tokenIndex = position479, tokenIndex479
														if buffer[position] != rune('Z') {
															goto l450
														}
														position++
													}
												l479:
													if !_rules[rulews]() {
														goto l450
													}
													if !_rules[ruledisp]() {
														goto l450
													}
													{
														add(ruleAction66, position)
													}
													add(ruleDjnz, position472)
												}
												break
											case 'J', 'j':
												{
													position482 := position
													{
														position483, tokenIndex483 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l484
														}
														position++
														goto l483
													l484:
														position, tokenIndex = position483, tokenIndex483
														if buffer[position] != rune('J') {
															goto l450
														}
														position++
													}
												l483:
													{
														position485, tokenIndex485 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l486
														}
														position++
														goto l485
													l486:
														position, tokenIndex = position485, tokenIndex485
														if buffer[position] != rune('R') {
															goto l450
														}
														position++
													}
												l485:
													if !_rules[rulews]() {
														goto l450
													}
													{
														position487, tokenIndex487 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l487
														}
														if !_rules[rulesep]() {
															goto l487
														}
														goto l488
													l487:
														position, tokenIndex = position487, tokenIndex487
													}
												l488:
													if !_rules[ruledisp]() {
														goto l450
													}
													{
														add(ruleAction65, position)
													}
													add(ruleJr, position482)
												}
												break
											case 'R', 'r':
												{
													position490 := position
													{
														position491, tokenIndex491 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l492
														}
														position++
														goto l491
													l492:
														position, tokenIndex = position491, tokenIndex491
														if buffer[position] != rune('R') {
															goto l450
														}
														position++
													}
												l491:
													{
														position493, tokenIndex493 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l494
														}
														position++
														goto l493
													l494:
														position, tokenIndex = position493, tokenIndex493
														if buffer[position] != rune('E') {
															goto l450
														}
														position++
													}
												l493:
													{
														position495, tokenIndex495 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l496
														}
														position++
														goto l495
													l496:
														position, tokenIndex = position495, tokenIndex495
														if buffer[position] != rune('T') {
															goto l450
														}
														position++
													}
												l495:
													{
														position497, tokenIndex497 := position, tokenIndex
														if !_rules[rulews]() {
															goto l497
														}
														if !_rules[rulecc]() {
															goto l497
														}
														goto l498
													l497:
														position, tokenIndex = position497, tokenIndex497
													}
												l498:
													{
														add(ruleAction63, position)
													}
													add(ruleRet, position490)
												}
												break
											default:
												{
													position500 := position
													{
														position501, tokenIndex501 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l502
														}
														position++
														goto l501
													l502:
														position, tokenIndex = position501, tokenIndex501
														if buffer[position] != rune('C') {
															goto l450
														}
														position++
													}
												l501:
													{
														position503, tokenIndex503 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l504
														}
														position++
														goto l503
													l504:
														position, tokenIndex = position503, tokenIndex503
														if buffer[position] != rune('A') {
															goto l450
														}
														position++
													}
												l503:
													{
														position505, tokenIndex505 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l506
														}
														position++
														goto l505
													l506:
														position, tokenIndex = position505, tokenIndex505
														if buffer[position] != rune('L') {
															goto l450
														}
														position++
													}
												l505:
													{
														position507, tokenIndex507 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l508
														}
														position++
														goto l507
													l508:
														position, tokenIndex = position507, tokenIndex507
														if buffer[position] != rune('L') {
															goto l450
														}
														position++
													}
												l507:
													if !_rules[rulews]() {
														goto l450
													}
													{
														position509, tokenIndex509 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l509
														}
														if !_rules[rulesep]() {
															goto l509
														}
														goto l510
													l509:
														position, tokenIndex = position509, tokenIndex509
													}
												l510:
													if !_rules[ruleSrc16]() {
														goto l450
													}
													{
														add(ruleAction62, position)
													}
													add(ruleCall, position500)
												}
												break
											}
										}

									}
								l452:
									add(ruleJump, position451)
								}
								goto l13
							l450:
								position, tokenIndex = position13, tokenIndex13
								{
									position512 := position
									{
										position513, tokenIndex513 := position, tokenIndex
										{
											position515 := position
											{
												position516, tokenIndex516 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l517
												}
												position++
												goto l516
											l517:
												position, tokenIndex = position516, tokenIndex516
												if buffer[position] != rune('I') {
													goto l514
												}
												position++
											}
										l516:
											{
												position518, tokenIndex518 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l519
												}
												position++
												goto l518
											l519:
												position, tokenIndex = position518, tokenIndex518
												if buffer[position] != rune('N') {
													goto l514
												}
												position++
											}
										l518:
											if !_rules[rulews]() {
												goto l514
											}
											if !_rules[ruleReg8]() {
												goto l514
											}
											if !_rules[rulesep]() {
												goto l514
											}
											if !_rules[rulePort]() {
												goto l514
											}
											{
												add(ruleAction67, position)
											}
											add(ruleIN, position515)
										}
										goto l513
									l514:
										position, tokenIndex = position513, tokenIndex513
										{
											position521 := position
											{
												position522, tokenIndex522 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l523
												}
												position++
												goto l522
											l523:
												position, tokenIndex = position522, tokenIndex522
												if buffer[position] != rune('O') {
													goto l0
												}
												position++
											}
										l522:
											{
												position524, tokenIndex524 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l525
												}
												position++
												goto l524
											l525:
												position, tokenIndex = position524, tokenIndex524
												if buffer[position] != rune('U') {
													goto l0
												}
												position++
											}
										l524:
											{
												position526, tokenIndex526 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l527
												}
												position++
												goto l526
											l527:
												position, tokenIndex = position526, tokenIndex526
												if buffer[position] != rune('T') {
													goto l0
												}
												position++
											}
										l526:
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
												add(ruleAction68, position)
											}
											add(ruleOUT, position521)
										}
									}
								l513:
									add(ruleIO, position512)
								}
							}
						l13:
							add(ruleInstruction, position10)
						}
						{
							position529, tokenIndex529 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l529
							}
							position++
							goto l530
						l529:
							position, tokenIndex = position529, tokenIndex529
						}
					l530:
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
						position532, tokenIndex532 := position, tokenIndex
						{
							position534 := position
						l535:
							{
								position536, tokenIndex536 := position, tokenIndex
								if !_rules[rulews]() {
									goto l536
								}
								goto l535
							l536:
								position, tokenIndex = position536, tokenIndex536
							}
							if buffer[position] != rune('\n') {
								goto l533
							}
							position++
							add(ruleBlankLine, position534)
						}
						goto l532
					l533:
						position, tokenIndex = position532, tokenIndex532
						{
							position537 := position
							{
								position538 := position
							l539:
								{
									position540, tokenIndex540 := position, tokenIndex
									if !_rules[rulews]() {
										goto l540
									}
									goto l539
								l540:
									position, tokenIndex = position540, tokenIndex540
								}
								{
									position541, tokenIndex541 := position, tokenIndex
									{
										position543 := position
										{
											position544, tokenIndex544 := position, tokenIndex
											{
												position546 := position
												{
													position547, tokenIndex547 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l548
													}
													position++
													goto l547
												l548:
													position, tokenIndex = position547, tokenIndex547
													if buffer[position] != rune('P') {
														goto l545
													}
													position++
												}
											l547:
												{
													position549, tokenIndex549 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l550
													}
													position++
													goto l549
												l550:
													position, tokenIndex = position549, tokenIndex549
													if buffer[position] != rune('U') {
														goto l545
													}
													position++
												}
											l549:
												{
													position551, tokenIndex551 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l552
													}
													position++
													goto l551
												l552:
													position, tokenIndex = position551, tokenIndex551
													if buffer[position] != rune('S') {
														goto l545
													}
													position++
												}
											l551:
												{
													position553, tokenIndex553 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l554
													}
													position++
													goto l553
												l554:
													position, tokenIndex = position553, tokenIndex553
													if buffer[position] != rune('H') {
														goto l545
													}
													position++
												}
											l553:
												if !_rules[rulews]() {
													goto l545
												}
												if !_rules[ruleSrc16]() {
													goto l545
												}
												{
													add(ruleAction3, position)
												}
												add(rulePush, position546)
											}
											goto l544
										l545:
											position, tokenIndex = position544, tokenIndex544
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position557 := position
														{
															position558, tokenIndex558 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l559
															}
															position++
															goto l558
														l559:
															position, tokenIndex = position558, tokenIndex558
															if buffer[position] != rune('E') {
																goto l542
															}
															position++
														}
													l558:
														{
															position560, tokenIndex560 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l561
															}
															position++
															goto l560
														l561:
															position, tokenIndex = position560, tokenIndex560
															if buffer[position] != rune('X') {
																goto l542
															}
															position++
														}
													l560:
														if !_rules[rulews]() {
															goto l542
														}
														if !_rules[ruleDst16]() {
															goto l542
														}
														if !_rules[rulesep]() {
															goto l542
														}
														if !_rules[ruleSrc16]() {
															goto l542
														}
														{
															add(ruleAction5, position)
														}
														add(ruleEx, position557)
													}
													break
												case 'P', 'p':
													{
														position563 := position
														{
															position564, tokenIndex564 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l565
															}
															position++
															goto l564
														l565:
															position, tokenIndex = position564, tokenIndex564
															if buffer[position] != rune('P') {
																goto l542
															}
															position++
														}
													l564:
														{
															position566, tokenIndex566 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l567
															}
															position++
															goto l566
														l567:
															position, tokenIndex = position566, tokenIndex566
															if buffer[position] != rune('O') {
																goto l542
															}
															position++
														}
													l566:
														{
															position568, tokenIndex568 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l569
															}
															position++
															goto l568
														l569:
															position, tokenIndex = position568, tokenIndex568
															if buffer[position] != rune('P') {
																goto l542
															}
															position++
														}
													l568:
														if !_rules[rulews]() {
															goto l542
														}
														if !_rules[ruleDst16]() {
															goto l542
														}
														{
															add(ruleAction4, position)
														}
														add(rulePop, position563)
													}
													break
												default:
													{
														position571 := position
														{
															position572, tokenIndex572 := position, tokenIndex
															{
																position574 := position
																{
																	position575, tokenIndex575 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l576
																	}
																	position++
																	goto l575
																l576:
																	position, tokenIndex = position575, tokenIndex575
																	if buffer[position] != rune('L') {
																		goto l573
																	}
																	position++
																}
															l575:
																{
																	position577, tokenIndex577 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l578
																	}
																	position++
																	goto l577
																l578:
																	position, tokenIndex = position577, tokenIndex577
																	if buffer[position] != rune('D') {
																		goto l573
																	}
																	position++
																}
															l577:
																if !_rules[rulews]() {
																	goto l573
																}
																if !_rules[ruleDst16]() {
																	goto l573
																}
																if !_rules[rulesep]() {
																	goto l573
																}
																if !_rules[ruleSrc16]() {
																	goto l573
																}
																{
																	add(ruleAction2, position)
																}
																add(ruleLoad16, position574)
															}
															goto l572
														l573:
															position, tokenIndex = position572, tokenIndex572
															{
																position580 := position
																{
																	position581, tokenIndex581 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l582
																	}
																	position++
																	goto l581
																l582:
																	position, tokenIndex = position581, tokenIndex581
																	if buffer[position] != rune('L') {
																		goto l542
																	}
																	position++
																}
															l581:
																{
																	position583, tokenIndex583 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l584
																	}
																	position++
																	goto l583
																l584:
																	position, tokenIndex = position583, tokenIndex583
																	if buffer[position] != rune('D') {
																		goto l542
																	}
																	position++
																}
															l583:
																if !_rules[rulews]() {
																	goto l542
																}
																{
																	position585 := position
																	{
																		position586, tokenIndex586 := position, tokenIndex
																		if !_rules[ruleReg8]() {
																			goto l587
																		}
																		goto l586
																	l587:
																		position, tokenIndex = position586, tokenIndex586
																		if !_rules[ruleReg16Contents]() {
																			goto l588
																		}
																		goto l586
																	l588:
																		position, tokenIndex = position586, tokenIndex586
																		if !_rules[rulenn_contents]() {
																			goto l542
																		}
																	}
																l586:
																	{
																		add(ruleAction15, position)
																	}
																	add(ruleDst8, position585)
																}
																if !_rules[rulesep]() {
																	goto l542
																}
																if !_rules[ruleSrc8]() {
																	goto l542
																}
																{
																	add(ruleAction1, position)
																}
																add(ruleLoad8, position580)
															}
														}
													l572:
														add(ruleLoad, position571)
													}
													break
												}
											}

										}
									l544:
										add(ruleAssignment, position543)
									}
									goto l541
								l542:
									position, tokenIndex = position541, tokenIndex541
									{
										position592 := position
										{
											position593, tokenIndex593 := position, tokenIndex
											{
												position595 := position
												{
													position596, tokenIndex596 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l597
													}
													position++
													goto l596
												l597:
													position, tokenIndex = position596, tokenIndex596
													if buffer[position] != rune('I') {
														goto l594
													}
													position++
												}
											l596:
												{
													position598, tokenIndex598 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l599
													}
													position++
													goto l598
												l599:
													position, tokenIndex = position598, tokenIndex598
													if buffer[position] != rune('N') {
														goto l594
													}
													position++
												}
											l598:
												{
													position600, tokenIndex600 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l601
													}
													position++
													goto l600
												l601:
													position, tokenIndex = position600, tokenIndex600
													if buffer[position] != rune('C') {
														goto l594
													}
													position++
												}
											l600:
												if !_rules[rulews]() {
													goto l594
												}
												if !_rules[ruleILoc8]() {
													goto l594
												}
												{
													add(ruleAction6, position)
												}
												add(ruleInc16Indexed8, position595)
											}
											goto l593
										l594:
											position, tokenIndex = position593, tokenIndex593
											{
												position604 := position
												{
													position605, tokenIndex605 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l606
													}
													position++
													goto l605
												l606:
													position, tokenIndex = position605, tokenIndex605
													if buffer[position] != rune('I') {
														goto l603
													}
													position++
												}
											l605:
												{
													position607, tokenIndex607 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l608
													}
													position++
													goto l607
												l608:
													position, tokenIndex = position607, tokenIndex607
													if buffer[position] != rune('N') {
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
												if !_rules[ruleLoc16]() {
													goto l603
												}
												{
													add(ruleAction8, position)
												}
												add(ruleInc16, position604)
											}
											goto l593
										l603:
											position, tokenIndex = position593, tokenIndex593
											{
												position612 := position
												{
													position613, tokenIndex613 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l614
													}
													position++
													goto l613
												l614:
													position, tokenIndex = position613, tokenIndex613
													if buffer[position] != rune('I') {
														goto l591
													}
													position++
												}
											l613:
												{
													position615, tokenIndex615 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l616
													}
													position++
													goto l615
												l616:
													position, tokenIndex = position615, tokenIndex615
													if buffer[position] != rune('N') {
														goto l591
													}
													position++
												}
											l615:
												{
													position617, tokenIndex617 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l618
													}
													position++
													goto l617
												l618:
													position, tokenIndex = position617, tokenIndex617
													if buffer[position] != rune('C') {
														goto l591
													}
													position++
												}
											l617:
												if !_rules[rulews]() {
													goto l591
												}
												if !_rules[ruleLoc8]() {
													goto l591
												}
												{
													add(ruleAction7, position)
												}
												add(ruleInc8, position612)
											}
										}
									l593:
										add(ruleInc, position592)
									}
									goto l541
								l591:
									position, tokenIndex = position541, tokenIndex541
									{
										position621 := position
										{
											position622, tokenIndex622 := position, tokenIndex
											{
												position624 := position
												{
													position625, tokenIndex625 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l626
													}
													position++
													goto l625
												l626:
													position, tokenIndex = position625, tokenIndex625
													if buffer[position] != rune('D') {
														goto l623
													}
													position++
												}
											l625:
												{
													position627, tokenIndex627 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l628
													}
													position++
													goto l627
												l628:
													position, tokenIndex = position627, tokenIndex627
													if buffer[position] != rune('E') {
														goto l623
													}
													position++
												}
											l627:
												{
													position629, tokenIndex629 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l630
													}
													position++
													goto l629
												l630:
													position, tokenIndex = position629, tokenIndex629
													if buffer[position] != rune('C') {
														goto l623
													}
													position++
												}
											l629:
												if !_rules[rulews]() {
													goto l623
												}
												if !_rules[ruleILoc8]() {
													goto l623
												}
												{
													add(ruleAction9, position)
												}
												add(ruleDec16Indexed8, position624)
											}
											goto l622
										l623:
											position, tokenIndex = position622, tokenIndex622
											{
												position633 := position
												{
													position634, tokenIndex634 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l635
													}
													position++
													goto l634
												l635:
													position, tokenIndex = position634, tokenIndex634
													if buffer[position] != rune('D') {
														goto l632
													}
													position++
												}
											l634:
												{
													position636, tokenIndex636 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l637
													}
													position++
													goto l636
												l637:
													position, tokenIndex = position636, tokenIndex636
													if buffer[position] != rune('E') {
														goto l632
													}
													position++
												}
											l636:
												{
													position638, tokenIndex638 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l639
													}
													position++
													goto l638
												l639:
													position, tokenIndex = position638, tokenIndex638
													if buffer[position] != rune('C') {
														goto l632
													}
													position++
												}
											l638:
												if !_rules[rulews]() {
													goto l632
												}
												if !_rules[ruleLoc16]() {
													goto l632
												}
												{
													add(ruleAction11, position)
												}
												add(ruleDec16, position633)
											}
											goto l622
										l632:
											position, tokenIndex = position622, tokenIndex622
											{
												position641 := position
												{
													position642, tokenIndex642 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l643
													}
													position++
													goto l642
												l643:
													position, tokenIndex = position642, tokenIndex642
													if buffer[position] != rune('D') {
														goto l620
													}
													position++
												}
											l642:
												{
													position644, tokenIndex644 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l645
													}
													position++
													goto l644
												l645:
													position, tokenIndex = position644, tokenIndex644
													if buffer[position] != rune('E') {
														goto l620
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
														goto l620
													}
													position++
												}
											l646:
												if !_rules[rulews]() {
													goto l620
												}
												if !_rules[ruleLoc8]() {
													goto l620
												}
												{
													add(ruleAction10, position)
												}
												add(ruleDec8, position641)
											}
										}
									l622:
										add(ruleDec, position621)
									}
									goto l541
								l620:
									position, tokenIndex = position541, tokenIndex541
									{
										position650 := position
										{
											position651, tokenIndex651 := position, tokenIndex
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
													if buffer[position] != rune('d') {
														goto l659
													}
													position++
													goto l658
												l659:
													position, tokenIndex = position658, tokenIndex658
													if buffer[position] != rune('D') {
														goto l652
													}
													position++
												}
											l658:
												if !_rules[rulews]() {
													goto l652
												}
												if !_rules[ruleDst16]() {
													goto l652
												}
												if !_rules[rulesep]() {
													goto l652
												}
												if !_rules[ruleSrc16]() {
													goto l652
												}
												{
													add(ruleAction12, position)
												}
												add(ruleAdd16, position653)
											}
											goto l651
										l652:
											position, tokenIndex = position651, tokenIndex651
											{
												position662 := position
												{
													position663, tokenIndex663 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l664
													}
													position++
													goto l663
												l664:
													position, tokenIndex = position663, tokenIndex663
													if buffer[position] != rune('A') {
														goto l661
													}
													position++
												}
											l663:
												{
													position665, tokenIndex665 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l666
													}
													position++
													goto l665
												l666:
													position, tokenIndex = position665, tokenIndex665
													if buffer[position] != rune('D') {
														goto l661
													}
													position++
												}
											l665:
												{
													position667, tokenIndex667 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l668
													}
													position++
													goto l667
												l668:
													position, tokenIndex = position667, tokenIndex667
													if buffer[position] != rune('C') {
														goto l661
													}
													position++
												}
											l667:
												if !_rules[rulews]() {
													goto l661
												}
												if !_rules[ruleDst16]() {
													goto l661
												}
												if !_rules[rulesep]() {
													goto l661
												}
												if !_rules[ruleSrc16]() {
													goto l661
												}
												{
													add(ruleAction13, position)
												}
												add(ruleAdc16, position662)
											}
											goto l651
										l661:
											position, tokenIndex = position651, tokenIndex651
											{
												position670 := position
												{
													position671, tokenIndex671 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l672
													}
													position++
													goto l671
												l672:
													position, tokenIndex = position671, tokenIndex671
													if buffer[position] != rune('S') {
														goto l649
													}
													position++
												}
											l671:
												{
													position673, tokenIndex673 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l674
													}
													position++
													goto l673
												l674:
													position, tokenIndex = position673, tokenIndex673
													if buffer[position] != rune('B') {
														goto l649
													}
													position++
												}
											l673:
												{
													position675, tokenIndex675 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l676
													}
													position++
													goto l675
												l676:
													position, tokenIndex = position675, tokenIndex675
													if buffer[position] != rune('C') {
														goto l649
													}
													position++
												}
											l675:
												if !_rules[rulews]() {
													goto l649
												}
												if !_rules[ruleDst16]() {
													goto l649
												}
												if !_rules[rulesep]() {
													goto l649
												}
												if !_rules[ruleSrc16]() {
													goto l649
												}
												{
													add(ruleAction14, position)
												}
												add(ruleSbc16, position670)
											}
										}
									l651:
										add(ruleAlu16, position650)
									}
									goto l541
								l649:
									position, tokenIndex = position541, tokenIndex541
									{
										position679 := position
										{
											position680, tokenIndex680 := position, tokenIndex
											{
												position682 := position
												{
													position683, tokenIndex683 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l684
													}
													position++
													goto l683
												l684:
													position, tokenIndex = position683, tokenIndex683
													if buffer[position] != rune('A') {
														goto l681
													}
													position++
												}
											l683:
												{
													position685, tokenIndex685 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l686
													}
													position++
													goto l685
												l686:
													position, tokenIndex = position685, tokenIndex685
													if buffer[position] != rune('D') {
														goto l681
													}
													position++
												}
											l685:
												{
													position687, tokenIndex687 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l688
													}
													position++
													goto l687
												l688:
													position, tokenIndex = position687, tokenIndex687
													if buffer[position] != rune('D') {
														goto l681
													}
													position++
												}
											l687:
												if !_rules[rulews]() {
													goto l681
												}
												{
													position689, tokenIndex689 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l690
													}
													position++
													goto l689
												l690:
													position, tokenIndex = position689, tokenIndex689
													if buffer[position] != rune('A') {
														goto l681
													}
													position++
												}
											l689:
												if !_rules[rulesep]() {
													goto l681
												}
												if !_rules[ruleSrc8]() {
													goto l681
												}
												{
													add(ruleAction29, position)
												}
												add(ruleAdd, position682)
											}
											goto l680
										l681:
											position, tokenIndex = position680, tokenIndex680
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
														goto l692
													}
													position++
												}
											l694:
												{
													position696, tokenIndex696 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l697
													}
													position++
													goto l696
												l697:
													position, tokenIndex = position696, tokenIndex696
													if buffer[position] != rune('D') {
														goto l692
													}
													position++
												}
											l696:
												{
													position698, tokenIndex698 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l699
													}
													position++
													goto l698
												l699:
													position, tokenIndex = position698, tokenIndex698
													if buffer[position] != rune('C') {
														goto l692
													}
													position++
												}
											l698:
												if !_rules[rulews]() {
													goto l692
												}
												{
													position700, tokenIndex700 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l701
													}
													position++
													goto l700
												l701:
													position, tokenIndex = position700, tokenIndex700
													if buffer[position] != rune('A') {
														goto l692
													}
													position++
												}
											l700:
												if !_rules[rulesep]() {
													goto l692
												}
												if !_rules[ruleSrc8]() {
													goto l692
												}
												{
													add(ruleAction30, position)
												}
												add(ruleAdc, position693)
											}
											goto l680
										l692:
											position, tokenIndex = position680, tokenIndex680
											{
												position704 := position
												{
													position705, tokenIndex705 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l706
													}
													position++
													goto l705
												l706:
													position, tokenIndex = position705, tokenIndex705
													if buffer[position] != rune('S') {
														goto l703
													}
													position++
												}
											l705:
												{
													position707, tokenIndex707 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l708
													}
													position++
													goto l707
												l708:
													position, tokenIndex = position707, tokenIndex707
													if buffer[position] != rune('U') {
														goto l703
													}
													position++
												}
											l707:
												{
													position709, tokenIndex709 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l710
													}
													position++
													goto l709
												l710:
													position, tokenIndex = position709, tokenIndex709
													if buffer[position] != rune('B') {
														goto l703
													}
													position++
												}
											l709:
												if !_rules[rulews]() {
													goto l703
												}
												if !_rules[ruleSrc8]() {
													goto l703
												}
												{
													add(ruleAction31, position)
												}
												add(ruleSub, position704)
											}
											goto l680
										l703:
											position, tokenIndex = position680, tokenIndex680
											{
												switch buffer[position] {
												case 'C', 'c':
													{
														position713 := position
														{
															position714, tokenIndex714 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l715
															}
															position++
															goto l714
														l715:
															position, tokenIndex = position714, tokenIndex714
															if buffer[position] != rune('C') {
																goto l678
															}
															position++
														}
													l714:
														{
															position716, tokenIndex716 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l717
															}
															position++
															goto l716
														l717:
															position, tokenIndex = position716, tokenIndex716
															if buffer[position] != rune('P') {
																goto l678
															}
															position++
														}
													l716:
														if !_rules[rulews]() {
															goto l678
														}
														if !_rules[ruleSrc8]() {
															goto l678
														}
														{
															add(ruleAction36, position)
														}
														add(ruleCp, position713)
													}
													break
												case 'O', 'o':
													{
														position719 := position
														{
															position720, tokenIndex720 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l721
															}
															position++
															goto l720
														l721:
															position, tokenIndex = position720, tokenIndex720
															if buffer[position] != rune('O') {
																goto l678
															}
															position++
														}
													l720:
														{
															position722, tokenIndex722 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l723
															}
															position++
															goto l722
														l723:
															position, tokenIndex = position722, tokenIndex722
															if buffer[position] != rune('R') {
																goto l678
															}
															position++
														}
													l722:
														if !_rules[rulews]() {
															goto l678
														}
														if !_rules[ruleSrc8]() {
															goto l678
														}
														{
															add(ruleAction35, position)
														}
														add(ruleOr, position719)
													}
													break
												case 'X', 'x':
													{
														position725 := position
														{
															position726, tokenIndex726 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l727
															}
															position++
															goto l726
														l727:
															position, tokenIndex = position726, tokenIndex726
															if buffer[position] != rune('X') {
																goto l678
															}
															position++
														}
													l726:
														{
															position728, tokenIndex728 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l729
															}
															position++
															goto l728
														l729:
															position, tokenIndex = position728, tokenIndex728
															if buffer[position] != rune('O') {
																goto l678
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
																goto l678
															}
															position++
														}
													l730:
														if !_rules[rulews]() {
															goto l678
														}
														if !_rules[ruleSrc8]() {
															goto l678
														}
														{
															add(ruleAction34, position)
														}
														add(ruleXor, position725)
													}
													break
												case 'A', 'a':
													{
														position733 := position
														{
															position734, tokenIndex734 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l735
															}
															position++
															goto l734
														l735:
															position, tokenIndex = position734, tokenIndex734
															if buffer[position] != rune('A') {
																goto l678
															}
															position++
														}
													l734:
														{
															position736, tokenIndex736 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l737
															}
															position++
															goto l736
														l737:
															position, tokenIndex = position736, tokenIndex736
															if buffer[position] != rune('N') {
																goto l678
															}
															position++
														}
													l736:
														{
															position738, tokenIndex738 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l739
															}
															position++
															goto l738
														l739:
															position, tokenIndex = position738, tokenIndex738
															if buffer[position] != rune('D') {
																goto l678
															}
															position++
														}
													l738:
														if !_rules[rulews]() {
															goto l678
														}
														if !_rules[ruleSrc8]() {
															goto l678
														}
														{
															add(ruleAction33, position)
														}
														add(ruleAnd, position733)
													}
													break
												default:
													{
														position741 := position
														{
															position742, tokenIndex742 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l743
															}
															position++
															goto l742
														l743:
															position, tokenIndex = position742, tokenIndex742
															if buffer[position] != rune('S') {
																goto l678
															}
															position++
														}
													l742:
														{
															position744, tokenIndex744 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l745
															}
															position++
															goto l744
														l745:
															position, tokenIndex = position744, tokenIndex744
															if buffer[position] != rune('B') {
																goto l678
															}
															position++
														}
													l744:
														{
															position746, tokenIndex746 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l747
															}
															position++
															goto l746
														l747:
															position, tokenIndex = position746, tokenIndex746
															if buffer[position] != rune('C') {
																goto l678
															}
															position++
														}
													l746:
														if !_rules[rulews]() {
															goto l678
														}
														{
															position748, tokenIndex748 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l749
															}
															position++
															goto l748
														l749:
															position, tokenIndex = position748, tokenIndex748
															if buffer[position] != rune('A') {
																goto l678
															}
															position++
														}
													l748:
														if !_rules[rulesep]() {
															goto l678
														}
														if !_rules[ruleSrc8]() {
															goto l678
														}
														{
															add(ruleAction32, position)
														}
														add(ruleSbc, position741)
													}
													break
												}
											}

										}
									l680:
										add(ruleAlu, position679)
									}
									goto l541
								l678:
									position, tokenIndex = position541, tokenIndex541
									{
										position752 := position
										{
											position753, tokenIndex753 := position, tokenIndex
											{
												position755 := position
												{
													position756, tokenIndex756 := position, tokenIndex
													{
														position758 := position
														{
															position759, tokenIndex759 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l760
															}
															position++
															goto l759
														l760:
															position, tokenIndex = position759, tokenIndex759
															if buffer[position] != rune('R') {
																goto l757
															}
															position++
														}
													l759:
														{
															position761, tokenIndex761 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l762
															}
															position++
															goto l761
														l762:
															position, tokenIndex = position761, tokenIndex761
															if buffer[position] != rune('L') {
																goto l757
															}
															position++
														}
													l761:
														{
															position763, tokenIndex763 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l764
															}
															position++
															goto l763
														l764:
															position, tokenIndex = position763, tokenIndex763
															if buffer[position] != rune('C') {
																goto l757
															}
															position++
														}
													l763:
														if !_rules[rulews]() {
															goto l757
														}
														if !_rules[ruleLoc8]() {
															goto l757
														}
														{
															add(ruleAction37, position)
														}
														add(ruleRlc, position758)
													}
													goto l756
												l757:
													position, tokenIndex = position756, tokenIndex756
													{
														position767 := position
														{
															position768, tokenIndex768 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l769
															}
															position++
															goto l768
														l769:
															position, tokenIndex = position768, tokenIndex768
															if buffer[position] != rune('R') {
																goto l766
															}
															position++
														}
													l768:
														{
															position770, tokenIndex770 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l771
															}
															position++
															goto l770
														l771:
															position, tokenIndex = position770, tokenIndex770
															if buffer[position] != rune('R') {
																goto l766
															}
															position++
														}
													l770:
														{
															position772, tokenIndex772 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l773
															}
															position++
															goto l772
														l773:
															position, tokenIndex = position772, tokenIndex772
															if buffer[position] != rune('C') {
																goto l766
															}
															position++
														}
													l772:
														if !_rules[rulews]() {
															goto l766
														}
														if !_rules[ruleLoc8]() {
															goto l766
														}
														{
															add(ruleAction38, position)
														}
														add(ruleRrc, position767)
													}
													goto l756
												l766:
													position, tokenIndex = position756, tokenIndex756
													{
														position776 := position
														{
															position777, tokenIndex777 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l778
															}
															position++
															goto l777
														l778:
															position, tokenIndex = position777, tokenIndex777
															if buffer[position] != rune('R') {
																goto l775
															}
															position++
														}
													l777:
														{
															position779, tokenIndex779 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l780
															}
															position++
															goto l779
														l780:
															position, tokenIndex = position779, tokenIndex779
															if buffer[position] != rune('L') {
																goto l775
															}
															position++
														}
													l779:
														if !_rules[rulews]() {
															goto l775
														}
														if !_rules[ruleLoc8]() {
															goto l775
														}
														{
															add(ruleAction39, position)
														}
														add(ruleRl, position776)
													}
													goto l756
												l775:
													position, tokenIndex = position756, tokenIndex756
													{
														position783 := position
														{
															position784, tokenIndex784 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l785
															}
															position++
															goto l784
														l785:
															position, tokenIndex = position784, tokenIndex784
															if buffer[position] != rune('R') {
																goto l782
															}
															position++
														}
													l784:
														{
															position786, tokenIndex786 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l787
															}
															position++
															goto l786
														l787:
															position, tokenIndex = position786, tokenIndex786
															if buffer[position] != rune('R') {
																goto l782
															}
															position++
														}
													l786:
														if !_rules[rulews]() {
															goto l782
														}
														if !_rules[ruleLoc8]() {
															goto l782
														}
														{
															add(ruleAction40, position)
														}
														add(ruleRr, position783)
													}
													goto l756
												l782:
													position, tokenIndex = position756, tokenIndex756
													{
														position790 := position
														{
															position791, tokenIndex791 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l792
															}
															position++
															goto l791
														l792:
															position, tokenIndex = position791, tokenIndex791
															if buffer[position] != rune('S') {
																goto l789
															}
															position++
														}
													l791:
														{
															position793, tokenIndex793 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l794
															}
															position++
															goto l793
														l794:
															position, tokenIndex = position793, tokenIndex793
															if buffer[position] != rune('L') {
																goto l789
															}
															position++
														}
													l793:
														{
															position795, tokenIndex795 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l796
															}
															position++
															goto l795
														l796:
															position, tokenIndex = position795, tokenIndex795
															if buffer[position] != rune('A') {
																goto l789
															}
															position++
														}
													l795:
														if !_rules[rulews]() {
															goto l789
														}
														if !_rules[ruleLoc8]() {
															goto l789
														}
														{
															add(ruleAction41, position)
														}
														add(ruleSla, position790)
													}
													goto l756
												l789:
													position, tokenIndex = position756, tokenIndex756
													{
														position799 := position
														{
															position800, tokenIndex800 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l801
															}
															position++
															goto l800
														l801:
															position, tokenIndex = position800, tokenIndex800
															if buffer[position] != rune('S') {
																goto l798
															}
															position++
														}
													l800:
														{
															position802, tokenIndex802 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l803
															}
															position++
															goto l802
														l803:
															position, tokenIndex = position802, tokenIndex802
															if buffer[position] != rune('R') {
																goto l798
															}
															position++
														}
													l802:
														{
															position804, tokenIndex804 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l805
															}
															position++
															goto l804
														l805:
															position, tokenIndex = position804, tokenIndex804
															if buffer[position] != rune('A') {
																goto l798
															}
															position++
														}
													l804:
														if !_rules[rulews]() {
															goto l798
														}
														if !_rules[ruleLoc8]() {
															goto l798
														}
														{
															add(ruleAction42, position)
														}
														add(ruleSra, position799)
													}
													goto l756
												l798:
													position, tokenIndex = position756, tokenIndex756
													{
														position808 := position
														{
															position809, tokenIndex809 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l810
															}
															position++
															goto l809
														l810:
															position, tokenIndex = position809, tokenIndex809
															if buffer[position] != rune('S') {
																goto l807
															}
															position++
														}
													l809:
														{
															position811, tokenIndex811 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l812
															}
															position++
															goto l811
														l812:
															position, tokenIndex = position811, tokenIndex811
															if buffer[position] != rune('L') {
																goto l807
															}
															position++
														}
													l811:
														{
															position813, tokenIndex813 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l814
															}
															position++
															goto l813
														l814:
															position, tokenIndex = position813, tokenIndex813
															if buffer[position] != rune('L') {
																goto l807
															}
															position++
														}
													l813:
														if !_rules[rulews]() {
															goto l807
														}
														if !_rules[ruleLoc8]() {
															goto l807
														}
														{
															add(ruleAction43, position)
														}
														add(ruleSll, position808)
													}
													goto l756
												l807:
													position, tokenIndex = position756, tokenIndex756
													{
														position816 := position
														{
															position817, tokenIndex817 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l818
															}
															position++
															goto l817
														l818:
															position, tokenIndex = position817, tokenIndex817
															if buffer[position] != rune('S') {
																goto l754
															}
															position++
														}
													l817:
														{
															position819, tokenIndex819 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l820
															}
															position++
															goto l819
														l820:
															position, tokenIndex = position819, tokenIndex819
															if buffer[position] != rune('R') {
																goto l754
															}
															position++
														}
													l819:
														{
															position821, tokenIndex821 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l822
															}
															position++
															goto l821
														l822:
															position, tokenIndex = position821, tokenIndex821
															if buffer[position] != rune('L') {
																goto l754
															}
															position++
														}
													l821:
														if !_rules[rulews]() {
															goto l754
														}
														if !_rules[ruleLoc8]() {
															goto l754
														}
														{
															add(ruleAction44, position)
														}
														add(ruleSrl, position816)
													}
												}
											l756:
												add(ruleRot, position755)
											}
											goto l753
										l754:
											position, tokenIndex = position753, tokenIndex753
											{
												switch buffer[position] {
												case 'S', 's':
													{
														position825 := position
														{
															position826, tokenIndex826 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l827
															}
															position++
															goto l826
														l827:
															position, tokenIndex = position826, tokenIndex826
															if buffer[position] != rune('S') {
																goto l751
															}
															position++
														}
													l826:
														{
															position828, tokenIndex828 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l829
															}
															position++
															goto l828
														l829:
															position, tokenIndex = position828, tokenIndex828
															if buffer[position] != rune('E') {
																goto l751
															}
															position++
														}
													l828:
														{
															position830, tokenIndex830 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l831
															}
															position++
															goto l830
														l831:
															position, tokenIndex = position830, tokenIndex830
															if buffer[position] != rune('T') {
																goto l751
															}
															position++
														}
													l830:
														if !_rules[rulews]() {
															goto l751
														}
														if !_rules[ruleoctaldigit]() {
															goto l751
														}
														if !_rules[rulesep]() {
															goto l751
														}
														if !_rules[ruleLoc8]() {
															goto l751
														}
														{
															add(ruleAction47, position)
														}
														add(ruleSet, position825)
													}
													break
												case 'R', 'r':
													{
														position833 := position
														{
															position834, tokenIndex834 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l835
															}
															position++
															goto l834
														l835:
															position, tokenIndex = position834, tokenIndex834
															if buffer[position] != rune('R') {
																goto l751
															}
															position++
														}
													l834:
														{
															position836, tokenIndex836 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l837
															}
															position++
															goto l836
														l837:
															position, tokenIndex = position836, tokenIndex836
															if buffer[position] != rune('E') {
																goto l751
															}
															position++
														}
													l836:
														{
															position838, tokenIndex838 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l839
															}
															position++
															goto l838
														l839:
															position, tokenIndex = position838, tokenIndex838
															if buffer[position] != rune('S') {
																goto l751
															}
															position++
														}
													l838:
														if !_rules[rulews]() {
															goto l751
														}
														if !_rules[ruleoctaldigit]() {
															goto l751
														}
														if !_rules[rulesep]() {
															goto l751
														}
														if !_rules[ruleLoc8]() {
															goto l751
														}
														{
															add(ruleAction46, position)
														}
														add(ruleRes, position833)
													}
													break
												default:
													{
														position841 := position
														{
															position842, tokenIndex842 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l843
															}
															position++
															goto l842
														l843:
															position, tokenIndex = position842, tokenIndex842
															if buffer[position] != rune('B') {
																goto l751
															}
															position++
														}
													l842:
														{
															position844, tokenIndex844 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l845
															}
															position++
															goto l844
														l845:
															position, tokenIndex = position844, tokenIndex844
															if buffer[position] != rune('I') {
																goto l751
															}
															position++
														}
													l844:
														{
															position846, tokenIndex846 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l847
															}
															position++
															goto l846
														l847:
															position, tokenIndex = position846, tokenIndex846
															if buffer[position] != rune('T') {
																goto l751
															}
															position++
														}
													l846:
														if !_rules[rulews]() {
															goto l751
														}
														if !_rules[ruleoctaldigit]() {
															goto l751
														}
														if !_rules[rulesep]() {
															goto l751
														}
														if !_rules[ruleLoc8]() {
															goto l751
														}
														{
															add(ruleAction45, position)
														}
														add(ruleBit, position841)
													}
													break
												}
											}

										}
									l753:
										add(ruleBitOp, position752)
									}
									goto l541
								l751:
									position, tokenIndex = position541, tokenIndex541
									{
										position850 := position
										{
											position851, tokenIndex851 := position, tokenIndex
											{
												position853 := position
												{
													position854 := position
													{
														position855, tokenIndex855 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l856
														}
														position++
														goto l855
													l856:
														position, tokenIndex = position855, tokenIndex855
														if buffer[position] != rune('R') {
															goto l852
														}
														position++
													}
												l855:
													{
														position857, tokenIndex857 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l858
														}
														position++
														goto l857
													l858:
														position, tokenIndex = position857, tokenIndex857
														if buffer[position] != rune('L') {
															goto l852
														}
														position++
													}
												l857:
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
															goto l852
														}
														position++
													}
												l859:
													{
														position861, tokenIndex861 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l862
														}
														position++
														goto l861
													l862:
														position, tokenIndex = position861, tokenIndex861
														if buffer[position] != rune('A') {
															goto l852
														}
														position++
													}
												l861:
													add(rulePegText, position854)
												}
												{
													add(ruleAction50, position)
												}
												add(ruleRlca, position853)
											}
											goto l851
										l852:
											position, tokenIndex = position851, tokenIndex851
											{
												position865 := position
												{
													position866 := position
													{
														position867, tokenIndex867 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l868
														}
														position++
														goto l867
													l868:
														position, tokenIndex = position867, tokenIndex867
														if buffer[position] != rune('R') {
															goto l864
														}
														position++
													}
												l867:
													{
														position869, tokenIndex869 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l870
														}
														position++
														goto l869
													l870:
														position, tokenIndex = position869, tokenIndex869
														if buffer[position] != rune('R') {
															goto l864
														}
														position++
													}
												l869:
													{
														position871, tokenIndex871 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l872
														}
														position++
														goto l871
													l872:
														position, tokenIndex = position871, tokenIndex871
														if buffer[position] != rune('C') {
															goto l864
														}
														position++
													}
												l871:
													{
														position873, tokenIndex873 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l874
														}
														position++
														goto l873
													l874:
														position, tokenIndex = position873, tokenIndex873
														if buffer[position] != rune('A') {
															goto l864
														}
														position++
													}
												l873:
													add(rulePegText, position866)
												}
												{
													add(ruleAction51, position)
												}
												add(ruleRrca, position865)
											}
											goto l851
										l864:
											position, tokenIndex = position851, tokenIndex851
											{
												position877 := position
												{
													position878 := position
													{
														position879, tokenIndex879 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l880
														}
														position++
														goto l879
													l880:
														position, tokenIndex = position879, tokenIndex879
														if buffer[position] != rune('R') {
															goto l876
														}
														position++
													}
												l879:
													{
														position881, tokenIndex881 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l882
														}
														position++
														goto l881
													l882:
														position, tokenIndex = position881, tokenIndex881
														if buffer[position] != rune('L') {
															goto l876
														}
														position++
													}
												l881:
													{
														position883, tokenIndex883 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l884
														}
														position++
														goto l883
													l884:
														position, tokenIndex = position883, tokenIndex883
														if buffer[position] != rune('A') {
															goto l876
														}
														position++
													}
												l883:
													add(rulePegText, position878)
												}
												{
													add(ruleAction52, position)
												}
												add(ruleRla, position877)
											}
											goto l851
										l876:
											position, tokenIndex = position851, tokenIndex851
											{
												position887 := position
												{
													position888 := position
													{
														position889, tokenIndex889 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l890
														}
														position++
														goto l889
													l890:
														position, tokenIndex = position889, tokenIndex889
														if buffer[position] != rune('D') {
															goto l886
														}
														position++
													}
												l889:
													{
														position891, tokenIndex891 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l892
														}
														position++
														goto l891
													l892:
														position, tokenIndex = position891, tokenIndex891
														if buffer[position] != rune('A') {
															goto l886
														}
														position++
													}
												l891:
													{
														position893, tokenIndex893 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l894
														}
														position++
														goto l893
													l894:
														position, tokenIndex = position893, tokenIndex893
														if buffer[position] != rune('A') {
															goto l886
														}
														position++
													}
												l893:
													add(rulePegText, position888)
												}
												{
													add(ruleAction54, position)
												}
												add(ruleDaa, position887)
											}
											goto l851
										l886:
											position, tokenIndex = position851, tokenIndex851
											{
												position897 := position
												{
													position898 := position
													{
														position899, tokenIndex899 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l900
														}
														position++
														goto l899
													l900:
														position, tokenIndex = position899, tokenIndex899
														if buffer[position] != rune('C') {
															goto l896
														}
														position++
													}
												l899:
													{
														position901, tokenIndex901 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l902
														}
														position++
														goto l901
													l902:
														position, tokenIndex = position901, tokenIndex901
														if buffer[position] != rune('P') {
															goto l896
														}
														position++
													}
												l901:
													{
														position903, tokenIndex903 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l904
														}
														position++
														goto l903
													l904:
														position, tokenIndex = position903, tokenIndex903
														if buffer[position] != rune('L') {
															goto l896
														}
														position++
													}
												l903:
													add(rulePegText, position898)
												}
												{
													add(ruleAction55, position)
												}
												add(ruleCpl, position897)
											}
											goto l851
										l896:
											position, tokenIndex = position851, tokenIndex851
											{
												position907 := position
												{
													position908 := position
													{
														position909, tokenIndex909 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l910
														}
														position++
														goto l909
													l910:
														position, tokenIndex = position909, tokenIndex909
														if buffer[position] != rune('E') {
															goto l906
														}
														position++
													}
												l909:
													{
														position911, tokenIndex911 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l912
														}
														position++
														goto l911
													l912:
														position, tokenIndex = position911, tokenIndex911
														if buffer[position] != rune('X') {
															goto l906
														}
														position++
													}
												l911:
													{
														position913, tokenIndex913 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l914
														}
														position++
														goto l913
													l914:
														position, tokenIndex = position913, tokenIndex913
														if buffer[position] != rune('X') {
															goto l906
														}
														position++
													}
												l913:
													add(rulePegText, position908)
												}
												{
													add(ruleAction58, position)
												}
												add(ruleExx, position907)
											}
											goto l851
										l906:
											position, tokenIndex = position851, tokenIndex851
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position917 := position
														{
															position918 := position
															{
																position919, tokenIndex919 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l920
																}
																position++
																goto l919
															l920:
																position, tokenIndex = position919, tokenIndex919
																if buffer[position] != rune('E') {
																	goto l849
																}
																position++
															}
														l919:
															{
																position921, tokenIndex921 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l922
																}
																position++
																goto l921
															l922:
																position, tokenIndex = position921, tokenIndex921
																if buffer[position] != rune('I') {
																	goto l849
																}
																position++
															}
														l921:
															add(rulePegText, position918)
														}
														{
															add(ruleAction60, position)
														}
														add(ruleEi, position917)
													}
													break
												case 'D', 'd':
													{
														position924 := position
														{
															position925 := position
															{
																position926, tokenIndex926 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l927
																}
																position++
																goto l926
															l927:
																position, tokenIndex = position926, tokenIndex926
																if buffer[position] != rune('D') {
																	goto l849
																}
																position++
															}
														l926:
															{
																position928, tokenIndex928 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l929
																}
																position++
																goto l928
															l929:
																position, tokenIndex = position928, tokenIndex928
																if buffer[position] != rune('I') {
																	goto l849
																}
																position++
															}
														l928:
															add(rulePegText, position925)
														}
														{
															add(ruleAction59, position)
														}
														add(ruleDi, position924)
													}
													break
												case 'C', 'c':
													{
														position931 := position
														{
															position932 := position
															{
																position933, tokenIndex933 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l934
																}
																position++
																goto l933
															l934:
																position, tokenIndex = position933, tokenIndex933
																if buffer[position] != rune('C') {
																	goto l849
																}
																position++
															}
														l933:
															{
																position935, tokenIndex935 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l936
																}
																position++
																goto l935
															l936:
																position, tokenIndex = position935, tokenIndex935
																if buffer[position] != rune('C') {
																	goto l849
																}
																position++
															}
														l935:
															{
																position937, tokenIndex937 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l938
																}
																position++
																goto l937
															l938:
																position, tokenIndex = position937, tokenIndex937
																if buffer[position] != rune('F') {
																	goto l849
																}
																position++
															}
														l937:
															add(rulePegText, position932)
														}
														{
															add(ruleAction57, position)
														}
														add(ruleCcf, position931)
													}
													break
												case 'S', 's':
													{
														position940 := position
														{
															position941 := position
															{
																position942, tokenIndex942 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l943
																}
																position++
																goto l942
															l943:
																position, tokenIndex = position942, tokenIndex942
																if buffer[position] != rune('S') {
																	goto l849
																}
																position++
															}
														l942:
															{
																position944, tokenIndex944 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l945
																}
																position++
																goto l944
															l945:
																position, tokenIndex = position944, tokenIndex944
																if buffer[position] != rune('C') {
																	goto l849
																}
																position++
															}
														l944:
															{
																position946, tokenIndex946 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l947
																}
																position++
																goto l946
															l947:
																position, tokenIndex = position946, tokenIndex946
																if buffer[position] != rune('F') {
																	goto l849
																}
																position++
															}
														l946:
															add(rulePegText, position941)
														}
														{
															add(ruleAction56, position)
														}
														add(ruleScf, position940)
													}
													break
												case 'R', 'r':
													{
														position949 := position
														{
															position950 := position
															{
																position951, tokenIndex951 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l952
																}
																position++
																goto l951
															l952:
																position, tokenIndex = position951, tokenIndex951
																if buffer[position] != rune('R') {
																	goto l849
																}
																position++
															}
														l951:
															{
																position953, tokenIndex953 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l954
																}
																position++
																goto l953
															l954:
																position, tokenIndex = position953, tokenIndex953
																if buffer[position] != rune('R') {
																	goto l849
																}
																position++
															}
														l953:
															{
																position955, tokenIndex955 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l956
																}
																position++
																goto l955
															l956:
																position, tokenIndex = position955, tokenIndex955
																if buffer[position] != rune('A') {
																	goto l849
																}
																position++
															}
														l955:
															add(rulePegText, position950)
														}
														{
															add(ruleAction53, position)
														}
														add(ruleRra, position949)
													}
													break
												case 'H', 'h':
													{
														position958 := position
														{
															position959 := position
															{
																position960, tokenIndex960 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l961
																}
																position++
																goto l960
															l961:
																position, tokenIndex = position960, tokenIndex960
																if buffer[position] != rune('H') {
																	goto l849
																}
																position++
															}
														l960:
															{
																position962, tokenIndex962 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l963
																}
																position++
																goto l962
															l963:
																position, tokenIndex = position962, tokenIndex962
																if buffer[position] != rune('A') {
																	goto l849
																}
																position++
															}
														l962:
															{
																position964, tokenIndex964 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l965
																}
																position++
																goto l964
															l965:
																position, tokenIndex = position964, tokenIndex964
																if buffer[position] != rune('L') {
																	goto l849
																}
																position++
															}
														l964:
															{
																position966, tokenIndex966 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l967
																}
																position++
																goto l966
															l967:
																position, tokenIndex = position966, tokenIndex966
																if buffer[position] != rune('T') {
																	goto l849
																}
																position++
															}
														l966:
															add(rulePegText, position959)
														}
														{
															add(ruleAction49, position)
														}
														add(ruleHalt, position958)
													}
													break
												default:
													{
														position969 := position
														{
															position970 := position
															{
																position971, tokenIndex971 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l972
																}
																position++
																goto l971
															l972:
																position, tokenIndex = position971, tokenIndex971
																if buffer[position] != rune('N') {
																	goto l849
																}
																position++
															}
														l971:
															{
																position973, tokenIndex973 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l974
																}
																position++
																goto l973
															l974:
																position, tokenIndex = position973, tokenIndex973
																if buffer[position] != rune('O') {
																	goto l849
																}
																position++
															}
														l973:
															{
																position975, tokenIndex975 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l976
																}
																position++
																goto l975
															l976:
																position, tokenIndex = position975, tokenIndex975
																if buffer[position] != rune('P') {
																	goto l849
																}
																position++
															}
														l975:
															add(rulePegText, position970)
														}
														{
															add(ruleAction48, position)
														}
														add(ruleNop, position969)
													}
													break
												}
											}

										}
									l851:
										add(ruleSimple, position850)
									}
									goto l541
								l849:
									position, tokenIndex = position541, tokenIndex541
									{
										position979 := position
										{
											position980, tokenIndex980 := position, tokenIndex
											{
												position982 := position
												{
													position983, tokenIndex983 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l984
													}
													position++
													goto l983
												l984:
													position, tokenIndex = position983, tokenIndex983
													if buffer[position] != rune('R') {
														goto l981
													}
													position++
												}
											l983:
												{
													position985, tokenIndex985 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l986
													}
													position++
													goto l985
												l986:
													position, tokenIndex = position985, tokenIndex985
													if buffer[position] != rune('S') {
														goto l981
													}
													position++
												}
											l985:
												{
													position987, tokenIndex987 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l988
													}
													position++
													goto l987
												l988:
													position, tokenIndex = position987, tokenIndex987
													if buffer[position] != rune('T') {
														goto l981
													}
													position++
												}
											l987:
												if !_rules[rulews]() {
													goto l981
												}
												if !_rules[rulen]() {
													goto l981
												}
												{
													add(ruleAction61, position)
												}
												add(ruleRst, position982)
											}
											goto l980
										l981:
											position, tokenIndex = position980, tokenIndex980
											{
												position991 := position
												{
													position992, tokenIndex992 := position, tokenIndex
													if buffer[position] != rune('j') {
														goto l993
													}
													position++
													goto l992
												l993:
													position, tokenIndex = position992, tokenIndex992
													if buffer[position] != rune('J') {
														goto l990
													}
													position++
												}
											l992:
												{
													position994, tokenIndex994 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l995
													}
													position++
													goto l994
												l995:
													position, tokenIndex = position994, tokenIndex994
													if buffer[position] != rune('P') {
														goto l990
													}
													position++
												}
											l994:
												if !_rules[rulews]() {
													goto l990
												}
												{
													position996, tokenIndex996 := position, tokenIndex
													if !_rules[rulecc]() {
														goto l996
													}
													if !_rules[rulesep]() {
														goto l996
													}
													goto l997
												l996:
													position, tokenIndex = position996, tokenIndex996
												}
											l997:
												if !_rules[ruleSrc16]() {
													goto l990
												}
												{
													add(ruleAction64, position)
												}
												add(ruleJp, position991)
											}
											goto l980
										l990:
											position, tokenIndex = position980, tokenIndex980
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position1000 := position
														{
															position1001, tokenIndex1001 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1002
															}
															position++
															goto l1001
														l1002:
															position, tokenIndex = position1001, tokenIndex1001
															if buffer[position] != rune('D') {
																goto l978
															}
															position++
														}
													l1001:
														{
															position1003, tokenIndex1003 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1004
															}
															position++
															goto l1003
														l1004:
															position, tokenIndex = position1003, tokenIndex1003
															if buffer[position] != rune('J') {
																goto l978
															}
															position++
														}
													l1003:
														{
															position1005, tokenIndex1005 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1006
															}
															position++
															goto l1005
														l1006:
															position, tokenIndex = position1005, tokenIndex1005
															if buffer[position] != rune('N') {
																goto l978
															}
															position++
														}
													l1005:
														{
															position1007, tokenIndex1007 := position, tokenIndex
															if buffer[position] != rune('z') {
																goto l1008
															}
															position++
															goto l1007
														l1008:
															position, tokenIndex = position1007, tokenIndex1007
															if buffer[position] != rune('Z') {
																goto l978
															}
															position++
														}
													l1007:
														if !_rules[rulews]() {
															goto l978
														}
														if !_rules[ruledisp]() {
															goto l978
														}
														{
															add(ruleAction66, position)
														}
														add(ruleDjnz, position1000)
													}
													break
												case 'J', 'j':
													{
														position1010 := position
														{
															position1011, tokenIndex1011 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1012
															}
															position++
															goto l1011
														l1012:
															position, tokenIndex = position1011, tokenIndex1011
															if buffer[position] != rune('J') {
																goto l978
															}
															position++
														}
													l1011:
														{
															position1013, tokenIndex1013 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1014
															}
															position++
															goto l1013
														l1014:
															position, tokenIndex = position1013, tokenIndex1013
															if buffer[position] != rune('R') {
																goto l978
															}
															position++
														}
													l1013:
														if !_rules[rulews]() {
															goto l978
														}
														{
															position1015, tokenIndex1015 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1015
															}
															if !_rules[rulesep]() {
																goto l1015
															}
															goto l1016
														l1015:
															position, tokenIndex = position1015, tokenIndex1015
														}
													l1016:
														if !_rules[ruledisp]() {
															goto l978
														}
														{
															add(ruleAction65, position)
														}
														add(ruleJr, position1010)
													}
													break
												case 'R', 'r':
													{
														position1018 := position
														{
															position1019, tokenIndex1019 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1020
															}
															position++
															goto l1019
														l1020:
															position, tokenIndex = position1019, tokenIndex1019
															if buffer[position] != rune('R') {
																goto l978
															}
															position++
														}
													l1019:
														{
															position1021, tokenIndex1021 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1022
															}
															position++
															goto l1021
														l1022:
															position, tokenIndex = position1021, tokenIndex1021
															if buffer[position] != rune('E') {
																goto l978
															}
															position++
														}
													l1021:
														{
															position1023, tokenIndex1023 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1024
															}
															position++
															goto l1023
														l1024:
															position, tokenIndex = position1023, tokenIndex1023
															if buffer[position] != rune('T') {
																goto l978
															}
															position++
														}
													l1023:
														{
															position1025, tokenIndex1025 := position, tokenIndex
															if !_rules[rulews]() {
																goto l1025
															}
															if !_rules[rulecc]() {
																goto l1025
															}
															goto l1026
														l1025:
															position, tokenIndex = position1025, tokenIndex1025
														}
													l1026:
														{
															add(ruleAction63, position)
														}
														add(ruleRet, position1018)
													}
													break
												default:
													{
														position1028 := position
														{
															position1029, tokenIndex1029 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1030
															}
															position++
															goto l1029
														l1030:
															position, tokenIndex = position1029, tokenIndex1029
															if buffer[position] != rune('C') {
																goto l978
															}
															position++
														}
													l1029:
														{
															position1031, tokenIndex1031 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1032
															}
															position++
															goto l1031
														l1032:
															position, tokenIndex = position1031, tokenIndex1031
															if buffer[position] != rune('A') {
																goto l978
															}
															position++
														}
													l1031:
														{
															position1033, tokenIndex1033 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1034
															}
															position++
															goto l1033
														l1034:
															position, tokenIndex = position1033, tokenIndex1033
															if buffer[position] != rune('L') {
																goto l978
															}
															position++
														}
													l1033:
														{
															position1035, tokenIndex1035 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1036
															}
															position++
															goto l1035
														l1036:
															position, tokenIndex = position1035, tokenIndex1035
															if buffer[position] != rune('L') {
																goto l978
															}
															position++
														}
													l1035:
														if !_rules[rulews]() {
															goto l978
														}
														{
															position1037, tokenIndex1037 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1037
															}
															if !_rules[rulesep]() {
																goto l1037
															}
															goto l1038
														l1037:
															position, tokenIndex = position1037, tokenIndex1037
														}
													l1038:
														if !_rules[ruleSrc16]() {
															goto l978
														}
														{
															add(ruleAction62, position)
														}
														add(ruleCall, position1028)
													}
													break
												}
											}

										}
									l980:
										add(ruleJump, position979)
									}
									goto l541
								l978:
									position, tokenIndex = position541, tokenIndex541
									{
										position1040 := position
										{
											position1041, tokenIndex1041 := position, tokenIndex
											{
												position1043 := position
												{
													position1044, tokenIndex1044 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1045
													}
													position++
													goto l1044
												l1045:
													position, tokenIndex = position1044, tokenIndex1044
													if buffer[position] != rune('I') {
														goto l1042
													}
													position++
												}
											l1044:
												{
													position1046, tokenIndex1046 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1047
													}
													position++
													goto l1046
												l1047:
													position, tokenIndex = position1046, tokenIndex1046
													if buffer[position] != rune('N') {
														goto l1042
													}
													position++
												}
											l1046:
												if !_rules[rulews]() {
													goto l1042
												}
												if !_rules[ruleReg8]() {
													goto l1042
												}
												if !_rules[rulesep]() {
													goto l1042
												}
												if !_rules[rulePort]() {
													goto l1042
												}
												{
													add(ruleAction67, position)
												}
												add(ruleIN, position1043)
											}
											goto l1041
										l1042:
											position, tokenIndex = position1041, tokenIndex1041
											{
												position1049 := position
												{
													position1050, tokenIndex1050 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1051
													}
													position++
													goto l1050
												l1051:
													position, tokenIndex = position1050, tokenIndex1050
													if buffer[position] != rune('O') {
														goto l3
													}
													position++
												}
											l1050:
												{
													position1052, tokenIndex1052 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1053
													}
													position++
													goto l1052
												l1053:
													position, tokenIndex = position1052, tokenIndex1052
													if buffer[position] != rune('U') {
														goto l3
													}
													position++
												}
											l1052:
												{
													position1054, tokenIndex1054 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1055
													}
													position++
													goto l1054
												l1055:
													position, tokenIndex = position1054, tokenIndex1054
													if buffer[position] != rune('T') {
														goto l3
													}
													position++
												}
											l1054:
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
													add(ruleAction68, position)
												}
												add(ruleOUT, position1049)
											}
										}
									l1041:
										add(ruleIO, position1040)
									}
								}
							l541:
								add(ruleInstruction, position538)
							}
							{
								position1057, tokenIndex1057 := position, tokenIndex
								if buffer[position] != rune('\n') {
									goto l1057
								}
								position++
								goto l1058
							l1057:
								position, tokenIndex = position1057, tokenIndex1057
							}
						l1058:
							{
								add(ruleAction0, position)
							}
							add(ruleLine, position537)
						}
					}
				l532:
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
		/* 3 Instruction <- <(ws* (Assignment / Inc / Dec / Alu16 / Alu / BitOp / Simple / Jump / IO))> */
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
		/* 19 Alu16 <- <(Add16 / Adc16 / Sbc16)> */
		nil,
		/* 20 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action12)> */
		nil,
		/* 21 Adc16 <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws Dst16 sep Src16 Action13)> */
		nil,
		/* 22 Sbc16 <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws Dst16 sep Src16 Action14)> */
		nil,
		/* 23 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action15)> */
		nil,
		/* 24 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action16)> */
		func() bool {
			position1083, tokenIndex1083 := position, tokenIndex
			{
				position1084 := position
				{
					position1085, tokenIndex1085 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1086
					}
					goto l1085
				l1086:
					position, tokenIndex = position1085, tokenIndex1085
					if !_rules[ruleReg8]() {
						goto l1087
					}
					goto l1085
				l1087:
					position, tokenIndex = position1085, tokenIndex1085
					if !_rules[ruleReg16Contents]() {
						goto l1088
					}
					goto l1085
				l1088:
					position, tokenIndex = position1085, tokenIndex1085
					if !_rules[rulenn_contents]() {
						goto l1083
					}
				}
			l1085:
				{
					add(ruleAction16, position)
				}
				add(ruleSrc8, position1084)
			}
			return true
		l1083:
			position, tokenIndex = position1083, tokenIndex1083
			return false
		},
		/* 25 Loc8 <- <((Reg8 / Reg16Contents) Action17)> */
		func() bool {
			position1090, tokenIndex1090 := position, tokenIndex
			{
				position1091 := position
				{
					position1092, tokenIndex1092 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1093
					}
					goto l1092
				l1093:
					position, tokenIndex = position1092, tokenIndex1092
					if !_rules[ruleReg16Contents]() {
						goto l1090
					}
				}
			l1092:
				{
					add(ruleAction17, position)
				}
				add(ruleLoc8, position1091)
			}
			return true
		l1090:
			position, tokenIndex = position1090, tokenIndex1090
			return false
		},
		/* 26 ILoc8 <- <(IReg8 Action18)> */
		func() bool {
			position1095, tokenIndex1095 := position, tokenIndex
			{
				position1096 := position
				if !_rules[ruleIReg8]() {
					goto l1095
				}
				{
					add(ruleAction18, position)
				}
				add(ruleILoc8, position1096)
			}
			return true
		l1095:
			position, tokenIndex = position1095, tokenIndex1095
			return false
		},
		/* 27 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action19)> */
		func() bool {
			position1098, tokenIndex1098 := position, tokenIndex
			{
				position1099 := position
				{
					position1100 := position
					{
						position1101, tokenIndex1101 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1102
						}
						goto l1101
					l1102:
						position, tokenIndex = position1101, tokenIndex1101
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1104 := position
									{
										position1105, tokenIndex1105 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1106
										}
										position++
										goto l1105
									l1106:
										position, tokenIndex = position1105, tokenIndex1105
										if buffer[position] != rune('R') {
											goto l1098
										}
										position++
									}
								l1105:
									add(ruleR, position1104)
								}
								break
							case 'I', 'i':
								{
									position1107 := position
									{
										position1108, tokenIndex1108 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1109
										}
										position++
										goto l1108
									l1109:
										position, tokenIndex = position1108, tokenIndex1108
										if buffer[position] != rune('I') {
											goto l1098
										}
										position++
									}
								l1108:
									add(ruleI, position1107)
								}
								break
							case 'L', 'l':
								{
									position1110 := position
									{
										position1111, tokenIndex1111 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1112
										}
										position++
										goto l1111
									l1112:
										position, tokenIndex = position1111, tokenIndex1111
										if buffer[position] != rune('L') {
											goto l1098
										}
										position++
									}
								l1111:
									add(ruleL, position1110)
								}
								break
							case 'H', 'h':
								{
									position1113 := position
									{
										position1114, tokenIndex1114 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1115
										}
										position++
										goto l1114
									l1115:
										position, tokenIndex = position1114, tokenIndex1114
										if buffer[position] != rune('H') {
											goto l1098
										}
										position++
									}
								l1114:
									add(ruleH, position1113)
								}
								break
							case 'E', 'e':
								{
									position1116 := position
									{
										position1117, tokenIndex1117 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1118
										}
										position++
										goto l1117
									l1118:
										position, tokenIndex = position1117, tokenIndex1117
										if buffer[position] != rune('E') {
											goto l1098
										}
										position++
									}
								l1117:
									add(ruleE, position1116)
								}
								break
							case 'D', 'd':
								{
									position1119 := position
									{
										position1120, tokenIndex1120 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1121
										}
										position++
										goto l1120
									l1121:
										position, tokenIndex = position1120, tokenIndex1120
										if buffer[position] != rune('D') {
											goto l1098
										}
										position++
									}
								l1120:
									add(ruleD, position1119)
								}
								break
							case 'C', 'c':
								{
									position1122 := position
									{
										position1123, tokenIndex1123 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1124
										}
										position++
										goto l1123
									l1124:
										position, tokenIndex = position1123, tokenIndex1123
										if buffer[position] != rune('C') {
											goto l1098
										}
										position++
									}
								l1123:
									add(ruleC, position1122)
								}
								break
							case 'B', 'b':
								{
									position1125 := position
									{
										position1126, tokenIndex1126 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1127
										}
										position++
										goto l1126
									l1127:
										position, tokenIndex = position1126, tokenIndex1126
										if buffer[position] != rune('B') {
											goto l1098
										}
										position++
									}
								l1126:
									add(ruleB, position1125)
								}
								break
							case 'F', 'f':
								{
									position1128 := position
									{
										position1129, tokenIndex1129 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1130
										}
										position++
										goto l1129
									l1130:
										position, tokenIndex = position1129, tokenIndex1129
										if buffer[position] != rune('F') {
											goto l1098
										}
										position++
									}
								l1129:
									add(ruleF, position1128)
								}
								break
							default:
								{
									position1131 := position
									{
										position1132, tokenIndex1132 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1133
										}
										position++
										goto l1132
									l1133:
										position, tokenIndex = position1132, tokenIndex1132
										if buffer[position] != rune('A') {
											goto l1098
										}
										position++
									}
								l1132:
									add(ruleA, position1131)
								}
								break
							}
						}

					}
				l1101:
					add(rulePegText, position1100)
				}
				{
					add(ruleAction19, position)
				}
				add(ruleReg8, position1099)
			}
			return true
		l1098:
			position, tokenIndex = position1098, tokenIndex1098
			return false
		},
		/* 28 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action20)> */
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
								if buffer[position] != rune('i') {
									goto l1142
								}
								position++
								goto l1141
							l1142:
								position, tokenIndex = position1141, tokenIndex1141
								if buffer[position] != rune('I') {
									goto l1139
								}
								position++
							}
						l1141:
							{
								position1143, tokenIndex1143 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1144
								}
								position++
								goto l1143
							l1144:
								position, tokenIndex = position1143, tokenIndex1143
								if buffer[position] != rune('X') {
									goto l1139
								}
								position++
							}
						l1143:
							{
								position1145, tokenIndex1145 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1146
								}
								position++
								goto l1145
							l1146:
								position, tokenIndex = position1145, tokenIndex1145
								if buffer[position] != rune('H') {
									goto l1139
								}
								position++
							}
						l1145:
							add(ruleIXH, position1140)
						}
						goto l1138
					l1139:
						position, tokenIndex = position1138, tokenIndex1138
						{
							position1148 := position
							{
								position1149, tokenIndex1149 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1150
								}
								position++
								goto l1149
							l1150:
								position, tokenIndex = position1149, tokenIndex1149
								if buffer[position] != rune('I') {
									goto l1147
								}
								position++
							}
						l1149:
							{
								position1151, tokenIndex1151 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1152
								}
								position++
								goto l1151
							l1152:
								position, tokenIndex = position1151, tokenIndex1151
								if buffer[position] != rune('X') {
									goto l1147
								}
								position++
							}
						l1151:
							{
								position1153, tokenIndex1153 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1154
								}
								position++
								goto l1153
							l1154:
								position, tokenIndex = position1153, tokenIndex1153
								if buffer[position] != rune('L') {
									goto l1147
								}
								position++
							}
						l1153:
							add(ruleIXL, position1148)
						}
						goto l1138
					l1147:
						position, tokenIndex = position1138, tokenIndex1138
						{
							position1156 := position
							{
								position1157, tokenIndex1157 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1158
								}
								position++
								goto l1157
							l1158:
								position, tokenIndex = position1157, tokenIndex1157
								if buffer[position] != rune('I') {
									goto l1155
								}
								position++
							}
						l1157:
							{
								position1159, tokenIndex1159 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1160
								}
								position++
								goto l1159
							l1160:
								position, tokenIndex = position1159, tokenIndex1159
								if buffer[position] != rune('Y') {
									goto l1155
								}
								position++
							}
						l1159:
							{
								position1161, tokenIndex1161 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1162
								}
								position++
								goto l1161
							l1162:
								position, tokenIndex = position1161, tokenIndex1161
								if buffer[position] != rune('H') {
									goto l1155
								}
								position++
							}
						l1161:
							add(ruleIYH, position1156)
						}
						goto l1138
					l1155:
						position, tokenIndex = position1138, tokenIndex1138
						{
							position1163 := position
							{
								position1164, tokenIndex1164 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1165
								}
								position++
								goto l1164
							l1165:
								position, tokenIndex = position1164, tokenIndex1164
								if buffer[position] != rune('I') {
									goto l1135
								}
								position++
							}
						l1164:
							{
								position1166, tokenIndex1166 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1167
								}
								position++
								goto l1166
							l1167:
								position, tokenIndex = position1166, tokenIndex1166
								if buffer[position] != rune('Y') {
									goto l1135
								}
								position++
							}
						l1166:
							{
								position1168, tokenIndex1168 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1169
								}
								position++
								goto l1168
							l1169:
								position, tokenIndex = position1168, tokenIndex1168
								if buffer[position] != rune('L') {
									goto l1135
								}
								position++
							}
						l1168:
							add(ruleIYL, position1163)
						}
					}
				l1138:
					add(rulePegText, position1137)
				}
				{
					add(ruleAction20, position)
				}
				add(ruleIReg8, position1136)
			}
			return true
		l1135:
			position, tokenIndex = position1135, tokenIndex1135
			return false
		},
		/* 29 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action21)> */
		func() bool {
			position1171, tokenIndex1171 := position, tokenIndex
			{
				position1172 := position
				{
					position1173, tokenIndex1173 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1174
					}
					goto l1173
				l1174:
					position, tokenIndex = position1173, tokenIndex1173
					if !_rules[rulenn_contents]() {
						goto l1175
					}
					goto l1173
				l1175:
					position, tokenIndex = position1173, tokenIndex1173
					if !_rules[ruleReg16Contents]() {
						goto l1171
					}
				}
			l1173:
				{
					add(ruleAction21, position)
				}
				add(ruleDst16, position1172)
			}
			return true
		l1171:
			position, tokenIndex = position1171, tokenIndex1171
			return false
		},
		/* 30 Src16 <- <((Reg16 / nn / nn_contents) Action22)> */
		func() bool {
			position1177, tokenIndex1177 := position, tokenIndex
			{
				position1178 := position
				{
					position1179, tokenIndex1179 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1180
					}
					goto l1179
				l1180:
					position, tokenIndex = position1179, tokenIndex1179
					if !_rules[rulenn]() {
						goto l1181
					}
					goto l1179
				l1181:
					position, tokenIndex = position1179, tokenIndex1179
					if !_rules[rulenn_contents]() {
						goto l1177
					}
				}
			l1179:
				{
					add(ruleAction22, position)
				}
				add(ruleSrc16, position1178)
			}
			return true
		l1177:
			position, tokenIndex = position1177, tokenIndex1177
			return false
		},
		/* 31 Loc16 <- <(Reg16 Action23)> */
		func() bool {
			position1183, tokenIndex1183 := position, tokenIndex
			{
				position1184 := position
				if !_rules[ruleReg16]() {
					goto l1183
				}
				{
					add(ruleAction23, position)
				}
				add(ruleLoc16, position1184)
			}
			return true
		l1183:
			position, tokenIndex = position1183, tokenIndex1183
			return false
		},
		/* 32 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action24)> */
		func() bool {
			position1186, tokenIndex1186 := position, tokenIndex
			{
				position1187 := position
				{
					position1188 := position
					{
						position1189, tokenIndex1189 := position, tokenIndex
						{
							position1191 := position
							{
								position1192, tokenIndex1192 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1193
								}
								position++
								goto l1192
							l1193:
								position, tokenIndex = position1192, tokenIndex1192
								if buffer[position] != rune('A') {
									goto l1190
								}
								position++
							}
						l1192:
							{
								position1194, tokenIndex1194 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1195
								}
								position++
								goto l1194
							l1195:
								position, tokenIndex = position1194, tokenIndex1194
								if buffer[position] != rune('F') {
									goto l1190
								}
								position++
							}
						l1194:
							if buffer[position] != rune('\'') {
								goto l1190
							}
							position++
							add(ruleAF_PRIME, position1191)
						}
						goto l1189
					l1190:
						position, tokenIndex = position1189, tokenIndex1189
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1186
								}
								break
							case 'S', 's':
								{
									position1197 := position
									{
										position1198, tokenIndex1198 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1199
										}
										position++
										goto l1198
									l1199:
										position, tokenIndex = position1198, tokenIndex1198
										if buffer[position] != rune('S') {
											goto l1186
										}
										position++
									}
								l1198:
									{
										position1200, tokenIndex1200 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1201
										}
										position++
										goto l1200
									l1201:
										position, tokenIndex = position1200, tokenIndex1200
										if buffer[position] != rune('P') {
											goto l1186
										}
										position++
									}
								l1200:
									add(ruleSP, position1197)
								}
								break
							case 'H', 'h':
								{
									position1202 := position
									{
										position1203, tokenIndex1203 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1204
										}
										position++
										goto l1203
									l1204:
										position, tokenIndex = position1203, tokenIndex1203
										if buffer[position] != rune('H') {
											goto l1186
										}
										position++
									}
								l1203:
									{
										position1205, tokenIndex1205 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1206
										}
										position++
										goto l1205
									l1206:
										position, tokenIndex = position1205, tokenIndex1205
										if buffer[position] != rune('L') {
											goto l1186
										}
										position++
									}
								l1205:
									add(ruleHL, position1202)
								}
								break
							case 'D', 'd':
								{
									position1207 := position
									{
										position1208, tokenIndex1208 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1209
										}
										position++
										goto l1208
									l1209:
										position, tokenIndex = position1208, tokenIndex1208
										if buffer[position] != rune('D') {
											goto l1186
										}
										position++
									}
								l1208:
									{
										position1210, tokenIndex1210 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1211
										}
										position++
										goto l1210
									l1211:
										position, tokenIndex = position1210, tokenIndex1210
										if buffer[position] != rune('E') {
											goto l1186
										}
										position++
									}
								l1210:
									add(ruleDE, position1207)
								}
								break
							case 'B', 'b':
								{
									position1212 := position
									{
										position1213, tokenIndex1213 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1214
										}
										position++
										goto l1213
									l1214:
										position, tokenIndex = position1213, tokenIndex1213
										if buffer[position] != rune('B') {
											goto l1186
										}
										position++
									}
								l1213:
									{
										position1215, tokenIndex1215 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1216
										}
										position++
										goto l1215
									l1216:
										position, tokenIndex = position1215, tokenIndex1215
										if buffer[position] != rune('C') {
											goto l1186
										}
										position++
									}
								l1215:
									add(ruleBC, position1212)
								}
								break
							default:
								{
									position1217 := position
									{
										position1218, tokenIndex1218 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1219
										}
										position++
										goto l1218
									l1219:
										position, tokenIndex = position1218, tokenIndex1218
										if buffer[position] != rune('A') {
											goto l1186
										}
										position++
									}
								l1218:
									{
										position1220, tokenIndex1220 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1221
										}
										position++
										goto l1220
									l1221:
										position, tokenIndex = position1220, tokenIndex1220
										if buffer[position] != rune('F') {
											goto l1186
										}
										position++
									}
								l1220:
									add(ruleAF, position1217)
								}
								break
							}
						}

					}
				l1189:
					add(rulePegText, position1188)
				}
				{
					add(ruleAction24, position)
				}
				add(ruleReg16, position1187)
			}
			return true
		l1186:
			position, tokenIndex = position1186, tokenIndex1186
			return false
		},
		/* 33 IReg16 <- <(<(IX / IY)> Action25)> */
		func() bool {
			position1223, tokenIndex1223 := position, tokenIndex
			{
				position1224 := position
				{
					position1225 := position
					{
						position1226, tokenIndex1226 := position, tokenIndex
						{
							position1228 := position
							{
								position1229, tokenIndex1229 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1230
								}
								position++
								goto l1229
							l1230:
								position, tokenIndex = position1229, tokenIndex1229
								if buffer[position] != rune('I') {
									goto l1227
								}
								position++
							}
						l1229:
							{
								position1231, tokenIndex1231 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1232
								}
								position++
								goto l1231
							l1232:
								position, tokenIndex = position1231, tokenIndex1231
								if buffer[position] != rune('X') {
									goto l1227
								}
								position++
							}
						l1231:
							add(ruleIX, position1228)
						}
						goto l1226
					l1227:
						position, tokenIndex = position1226, tokenIndex1226
						{
							position1233 := position
							{
								position1234, tokenIndex1234 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1235
								}
								position++
								goto l1234
							l1235:
								position, tokenIndex = position1234, tokenIndex1234
								if buffer[position] != rune('I') {
									goto l1223
								}
								position++
							}
						l1234:
							{
								position1236, tokenIndex1236 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1237
								}
								position++
								goto l1236
							l1237:
								position, tokenIndex = position1236, tokenIndex1236
								if buffer[position] != rune('Y') {
									goto l1223
								}
								position++
							}
						l1236:
							add(ruleIY, position1233)
						}
					}
				l1226:
					add(rulePegText, position1225)
				}
				{
					add(ruleAction25, position)
				}
				add(ruleIReg16, position1224)
			}
			return true
		l1223:
			position, tokenIndex = position1223, tokenIndex1223
			return false
		},
		/* 34 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1239, tokenIndex1239 := position, tokenIndex
			{
				position1240 := position
				{
					position1241, tokenIndex1241 := position, tokenIndex
					{
						position1243 := position
						if buffer[position] != rune('(') {
							goto l1242
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1242
						}
						{
							position1244, tokenIndex1244 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1244
							}
							goto l1245
						l1244:
							position, tokenIndex = position1244, tokenIndex1244
						}
					l1245:
						if !_rules[rulesignedDecimalByte]() {
							goto l1242
						}
						{
							position1246, tokenIndex1246 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1246
							}
							goto l1247
						l1246:
							position, tokenIndex = position1246, tokenIndex1246
						}
					l1247:
						if buffer[position] != rune(')') {
							goto l1242
						}
						position++
						{
							add(ruleAction27, position)
						}
						add(ruleIndexedR16C, position1243)
					}
					goto l1241
				l1242:
					position, tokenIndex = position1241, tokenIndex1241
					{
						position1249 := position
						if buffer[position] != rune('(') {
							goto l1239
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1239
						}
						if buffer[position] != rune(')') {
							goto l1239
						}
						position++
						{
							add(ruleAction26, position)
						}
						add(rulePlainR16C, position1249)
					}
				}
			l1241:
				add(ruleReg16Contents, position1240)
			}
			return true
		l1239:
			position, tokenIndex = position1239, tokenIndex1239
			return false
		},
		/* 35 PlainR16C <- <('(' Reg16 ')' Action26)> */
		nil,
		/* 36 IndexedR16C <- <('(' IReg16 ws? signedDecimalByte ws? ')' Action27)> */
		nil,
		/* 37 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1253, tokenIndex1253 := position, tokenIndex
			{
				position1254 := position
				{
					position1255, tokenIndex1255 := position, tokenIndex
					{
						position1257 := position
						{
							position1258 := position
							if !_rules[rulehexdigit]() {
								goto l1256
							}
							if !_rules[rulehexdigit]() {
								goto l1256
							}
							add(rulePegText, position1258)
						}
						{
							position1259, tokenIndex1259 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1260
							}
							position++
							goto l1259
						l1260:
							position, tokenIndex = position1259, tokenIndex1259
							if buffer[position] != rune('H') {
								goto l1256
							}
							position++
						}
					l1259:
						{
							add(ruleAction69, position)
						}
						add(rulehexByteH, position1257)
					}
					goto l1255
				l1256:
					position, tokenIndex = position1255, tokenIndex1255
					{
						position1263 := position
						if buffer[position] != rune('0') {
							goto l1262
						}
						position++
						{
							position1264, tokenIndex1264 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1265
							}
							position++
							goto l1264
						l1265:
							position, tokenIndex = position1264, tokenIndex1264
							if buffer[position] != rune('X') {
								goto l1262
							}
							position++
						}
					l1264:
						{
							position1266 := position
							if !_rules[rulehexdigit]() {
								goto l1262
							}
							if !_rules[rulehexdigit]() {
								goto l1262
							}
							add(rulePegText, position1266)
						}
						{
							add(ruleAction70, position)
						}
						add(rulehexByte0x, position1263)
					}
					goto l1255
				l1262:
					position, tokenIndex = position1255, tokenIndex1255
					{
						position1268 := position
						{
							position1269 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1253
							}
							position++
						l1270:
							{
								position1271, tokenIndex1271 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1271
								}
								position++
								goto l1270
							l1271:
								position, tokenIndex = position1271, tokenIndex1271
							}
							add(rulePegText, position1269)
						}
						{
							add(ruleAction71, position)
						}
						add(ruledecimalByte, position1268)
					}
				}
			l1255:
				add(rulen, position1254)
			}
			return true
		l1253:
			position, tokenIndex = position1253, tokenIndex1253
			return false
		},
		/* 38 nn <- <(hexWordH / hexWord0x)> */
		func() bool {
			position1273, tokenIndex1273 := position, tokenIndex
			{
				position1274 := position
				{
					position1275, tokenIndex1275 := position, tokenIndex
					{
						position1277 := position
						{
							position1278 := position
							if !_rules[rulehexdigit]() {
								goto l1276
							}
							if !_rules[rulehexdigit]() {
								goto l1276
							}
							if !_rules[rulehexdigit]() {
								goto l1276
							}
							if !_rules[rulehexdigit]() {
								goto l1276
							}
							add(rulePegText, position1278)
						}
						{
							position1279, tokenIndex1279 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1280
							}
							position++
							goto l1279
						l1280:
							position, tokenIndex = position1279, tokenIndex1279
							if buffer[position] != rune('H') {
								goto l1276
							}
							position++
						}
					l1279:
						{
							add(ruleAction72, position)
						}
						add(rulehexWordH, position1277)
					}
					goto l1275
				l1276:
					position, tokenIndex = position1275, tokenIndex1275
					{
						position1282 := position
						if buffer[position] != rune('0') {
							goto l1273
						}
						position++
						{
							position1283, tokenIndex1283 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1284
							}
							position++
							goto l1283
						l1284:
							position, tokenIndex = position1283, tokenIndex1283
							if buffer[position] != rune('X') {
								goto l1273
							}
							position++
						}
					l1283:
						{
							position1285 := position
							if !_rules[rulehexdigit]() {
								goto l1273
							}
							if !_rules[rulehexdigit]() {
								goto l1273
							}
							if !_rules[rulehexdigit]() {
								goto l1273
							}
							if !_rules[rulehexdigit]() {
								goto l1273
							}
							add(rulePegText, position1285)
						}
						{
							add(ruleAction73, position)
						}
						add(rulehexWord0x, position1282)
					}
				}
			l1275:
				add(rulenn, position1274)
			}
			return true
		l1273:
			position, tokenIndex = position1273, tokenIndex1273
			return false
		},
		/* 39 nn_contents <- <('(' nn ')' Action28)> */
		func() bool {
			position1287, tokenIndex1287 := position, tokenIndex
			{
				position1288 := position
				if buffer[position] != rune('(') {
					goto l1287
				}
				position++
				if !_rules[rulenn]() {
					goto l1287
				}
				if buffer[position] != rune(')') {
					goto l1287
				}
				position++
				{
					add(ruleAction28, position)
				}
				add(rulenn_contents, position1288)
			}
			return true
		l1287:
			position, tokenIndex = position1287, tokenIndex1287
			return false
		},
		/* 40 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 41 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action29)> */
		nil,
		/* 42 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action30)> */
		nil,
		/* 43 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action31)> */
		nil,
		/* 44 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action32)> */
		nil,
		/* 45 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action33)> */
		nil,
		/* 46 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action34)> */
		nil,
		/* 47 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action35)> */
		nil,
		/* 48 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action36)> */
		nil,
		/* 49 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 50 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 51 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 Action37)> */
		nil,
		/* 52 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 Action38)> */
		nil,
		/* 53 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 Action39)> */
		nil,
		/* 54 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 Action40)> */
		nil,
		/* 55 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 Action41)> */
		nil,
		/* 56 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 Action42)> */
		nil,
		/* 57 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 Action43)> */
		nil,
		/* 58 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 Action44)> */
		nil,
		/* 59 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action45)> */
		nil,
		/* 60 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 Action46)> */
		nil,
		/* 61 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 Action47)> */
		nil,
		/* 62 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 63 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action48)> */
		nil,
		/* 64 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action49)> */
		nil,
		/* 65 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action50)> */
		nil,
		/* 66 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action51)> */
		nil,
		/* 67 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action52)> */
		nil,
		/* 68 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action53)> */
		nil,
		/* 69 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action54)> */
		nil,
		/* 70 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action55)> */
		nil,
		/* 71 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action56)> */
		nil,
		/* 72 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action57)> */
		nil,
		/* 73 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action58)> */
		nil,
		/* 74 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action59)> */
		nil,
		/* 75 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action60)> */
		nil,
		/* 76 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 77 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action61)> */
		nil,
		/* 78 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action62)> */
		nil,
		/* 79 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action63)> */
		nil,
		/* 80 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action64)> */
		nil,
		/* 81 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action65)> */
		nil,
		/* 82 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action66)> */
		nil,
		/* 83 disp <- <signedDecimalByte> */
		func() bool {
			position1333, tokenIndex1333 := position, tokenIndex
			{
				position1334 := position
				if !_rules[rulesignedDecimalByte]() {
					goto l1333
				}
				add(ruledisp, position1334)
			}
			return true
		l1333:
			position, tokenIndex = position1333, tokenIndex1333
			return false
		},
		/* 84 IO <- <(IN / OUT)> */
		nil,
		/* 85 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action67)> */
		nil,
		/* 86 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action68)> */
		nil,
		/* 87 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position1338, tokenIndex1338 := position, tokenIndex
			{
				position1339 := position
				{
					position1340, tokenIndex1340 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l1341
					}
					position++
					{
						position1342, tokenIndex1342 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l1343
						}
						position++
						goto l1342
					l1343:
						position, tokenIndex = position1342, tokenIndex1342
						if buffer[position] != rune('C') {
							goto l1341
						}
						position++
					}
				l1342:
					if buffer[position] != rune(')') {
						goto l1341
					}
					position++
					goto l1340
				l1341:
					position, tokenIndex = position1340, tokenIndex1340
					if buffer[position] != rune('(') {
						goto l1338
					}
					position++
					if !_rules[rulen]() {
						goto l1338
					}
					if buffer[position] != rune(')') {
						goto l1338
					}
					position++
				}
			l1340:
				add(rulePort, position1339)
			}
			return true
		l1338:
			position, tokenIndex = position1338, tokenIndex1338
			return false
		},
		/* 88 sep <- <(ws? ',' ws?)> */
		func() bool {
			position1344, tokenIndex1344 := position, tokenIndex
			{
				position1345 := position
				{
					position1346, tokenIndex1346 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1346
					}
					goto l1347
				l1346:
					position, tokenIndex = position1346, tokenIndex1346
				}
			l1347:
				if buffer[position] != rune(',') {
					goto l1344
				}
				position++
				{
					position1348, tokenIndex1348 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1348
					}
					goto l1349
				l1348:
					position, tokenIndex = position1348, tokenIndex1348
				}
			l1349:
				add(rulesep, position1345)
			}
			return true
		l1344:
			position, tokenIndex = position1344, tokenIndex1344
			return false
		},
		/* 89 ws <- <' '+> */
		func() bool {
			position1350, tokenIndex1350 := position, tokenIndex
			{
				position1351 := position
				if buffer[position] != rune(' ') {
					goto l1350
				}
				position++
			l1352:
				{
					position1353, tokenIndex1353 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l1353
					}
					position++
					goto l1352
				l1353:
					position, tokenIndex = position1353, tokenIndex1353
				}
				add(rulews, position1351)
			}
			return true
		l1350:
			position, tokenIndex = position1350, tokenIndex1350
			return false
		},
		/* 90 A <- <('a' / 'A')> */
		nil,
		/* 91 F <- <('f' / 'F')> */
		nil,
		/* 92 B <- <('b' / 'B')> */
		nil,
		/* 93 C <- <('c' / 'C')> */
		nil,
		/* 94 D <- <('d' / 'D')> */
		nil,
		/* 95 E <- <('e' / 'E')> */
		nil,
		/* 96 H <- <('h' / 'H')> */
		nil,
		/* 97 L <- <('l' / 'L')> */
		nil,
		/* 98 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 99 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 100 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 101 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 102 I <- <('i' / 'I')> */
		nil,
		/* 103 R <- <('r' / 'R')> */
		nil,
		/* 104 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 105 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 106 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 107 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 108 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 109 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 110 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 111 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 112 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action69)> */
		nil,
		/* 113 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action70)> */
		nil,
		/* 114 decimalByte <- <(<[0-9]+> Action71)> */
		nil,
		/* 115 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action72)> */
		nil,
		/* 116 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action73)> */
		nil,
		/* 117 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position1381, tokenIndex1381 := position, tokenIndex
			{
				position1382 := position
				{
					position1383, tokenIndex1383 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1384
					}
					position++
					goto l1383
				l1384:
					position, tokenIndex = position1383, tokenIndex1383
					{
						position1385, tokenIndex1385 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l1386
						}
						position++
						goto l1385
					l1386:
						position, tokenIndex = position1385, tokenIndex1385
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l1381
						}
						position++
					}
				l1385:
				}
			l1383:
				add(rulehexdigit, position1382)
			}
			return true
		l1381:
			position, tokenIndex = position1381, tokenIndex1381
			return false
		},
		/* 118 octaldigit <- <(<[0-7]> Action74)> */
		func() bool {
			position1387, tokenIndex1387 := position, tokenIndex
			{
				position1388 := position
				{
					position1389 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l1387
					}
					position++
					add(rulePegText, position1389)
				}
				{
					add(ruleAction74, position)
				}
				add(ruleoctaldigit, position1388)
			}
			return true
		l1387:
			position, tokenIndex = position1387, tokenIndex1387
			return false
		},
		/* 119 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action75)> */
		func() bool {
			position1391, tokenIndex1391 := position, tokenIndex
			{
				position1392 := position
				{
					position1393 := position
					{
						position1394, tokenIndex1394 := position, tokenIndex
						{
							position1396, tokenIndex1396 := position, tokenIndex
							if buffer[position] != rune('-') {
								goto l1397
							}
							position++
							goto l1396
						l1397:
							position, tokenIndex = position1396, tokenIndex1396
							if buffer[position] != rune('+') {
								goto l1394
							}
							position++
						}
					l1396:
						goto l1395
					l1394:
						position, tokenIndex = position1394, tokenIndex1394
					}
				l1395:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1391
					}
					position++
				l1398:
					{
						position1399, tokenIndex1399 := position, tokenIndex
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1399
						}
						position++
						goto l1398
					l1399:
						position, tokenIndex = position1399, tokenIndex1399
					}
					add(rulePegText, position1393)
				}
				{
					add(ruleAction75, position)
				}
				add(rulesignedDecimalByte, position1392)
			}
			return true
		l1391:
			position, tokenIndex = position1391, tokenIndex1391
			return false
		},
		/* 120 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position1401, tokenIndex1401 := position, tokenIndex
			{
				position1402 := position
				{
					position1403, tokenIndex1403 := position, tokenIndex
					{
						position1405 := position
						{
							position1406, tokenIndex1406 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l1407
							}
							position++
							goto l1406
						l1407:
							position, tokenIndex = position1406, tokenIndex1406
							if buffer[position] != rune('N') {
								goto l1404
							}
							position++
						}
					l1406:
						{
							position1408, tokenIndex1408 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l1409
							}
							position++
							goto l1408
						l1409:
							position, tokenIndex = position1408, tokenIndex1408
							if buffer[position] != rune('Z') {
								goto l1404
							}
							position++
						}
					l1408:
						{
							add(ruleAction76, position)
						}
						add(ruleFT_NZ, position1405)
					}
					goto l1403
				l1404:
					position, tokenIndex = position1403, tokenIndex1403
					{
						position1412 := position
						{
							position1413, tokenIndex1413 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1414
							}
							position++
							goto l1413
						l1414:
							position, tokenIndex = position1413, tokenIndex1413
							if buffer[position] != rune('P') {
								goto l1411
							}
							position++
						}
					l1413:
						{
							position1415, tokenIndex1415 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l1416
							}
							position++
							goto l1415
						l1416:
							position, tokenIndex = position1415, tokenIndex1415
							if buffer[position] != rune('O') {
								goto l1411
							}
							position++
						}
					l1415:
						{
							add(ruleAction80, position)
						}
						add(ruleFT_PO, position1412)
					}
					goto l1403
				l1411:
					position, tokenIndex = position1403, tokenIndex1403
					{
						position1419 := position
						{
							position1420, tokenIndex1420 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1421
							}
							position++
							goto l1420
						l1421:
							position, tokenIndex = position1420, tokenIndex1420
							if buffer[position] != rune('P') {
								goto l1418
							}
							position++
						}
					l1420:
						{
							position1422, tokenIndex1422 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l1423
							}
							position++
							goto l1422
						l1423:
							position, tokenIndex = position1422, tokenIndex1422
							if buffer[position] != rune('E') {
								goto l1418
							}
							position++
						}
					l1422:
						{
							add(ruleAction81, position)
						}
						add(ruleFT_PE, position1419)
					}
					goto l1403
				l1418:
					position, tokenIndex = position1403, tokenIndex1403
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position1426 := position
								{
									position1427, tokenIndex1427 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l1428
									}
									position++
									goto l1427
								l1428:
									position, tokenIndex = position1427, tokenIndex1427
									if buffer[position] != rune('M') {
										goto l1401
									}
									position++
								}
							l1427:
								{
									add(ruleAction83, position)
								}
								add(ruleFT_M, position1426)
							}
							break
						case 'P', 'p':
							{
								position1430 := position
								{
									position1431, tokenIndex1431 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l1432
									}
									position++
									goto l1431
								l1432:
									position, tokenIndex = position1431, tokenIndex1431
									if buffer[position] != rune('P') {
										goto l1401
									}
									position++
								}
							l1431:
								{
									add(ruleAction82, position)
								}
								add(ruleFT_P, position1430)
							}
							break
						case 'C', 'c':
							{
								position1434 := position
								{
									position1435, tokenIndex1435 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1436
									}
									position++
									goto l1435
								l1436:
									position, tokenIndex = position1435, tokenIndex1435
									if buffer[position] != rune('C') {
										goto l1401
									}
									position++
								}
							l1435:
								{
									add(ruleAction79, position)
								}
								add(ruleFT_C, position1434)
							}
							break
						case 'N', 'n':
							{
								position1438 := position
								{
									position1439, tokenIndex1439 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l1440
									}
									position++
									goto l1439
								l1440:
									position, tokenIndex = position1439, tokenIndex1439
									if buffer[position] != rune('N') {
										goto l1401
									}
									position++
								}
							l1439:
								{
									position1441, tokenIndex1441 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1442
									}
									position++
									goto l1441
								l1442:
									position, tokenIndex = position1441, tokenIndex1441
									if buffer[position] != rune('C') {
										goto l1401
									}
									position++
								}
							l1441:
								{
									add(ruleAction78, position)
								}
								add(ruleFT_NC, position1438)
							}
							break
						default:
							{
								position1444 := position
								{
									position1445, tokenIndex1445 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l1446
									}
									position++
									goto l1445
								l1446:
									position, tokenIndex = position1445, tokenIndex1445
									if buffer[position] != rune('Z') {
										goto l1401
									}
									position++
								}
							l1445:
								{
									add(ruleAction77, position)
								}
								add(ruleFT_Z, position1444)
							}
							break
						}
					}

				}
			l1403:
				add(rulecc, position1402)
			}
			return true
		l1401:
			position, tokenIndex = position1401, tokenIndex1401
			return false
		},
		/* 121 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action76)> */
		nil,
		/* 122 FT_Z <- <(('z' / 'Z') Action77)> */
		nil,
		/* 123 FT_NC <- <(('n' / 'N') ('c' / 'C') Action78)> */
		nil,
		/* 124 FT_C <- <(('c' / 'C') Action79)> */
		nil,
		/* 125 FT_PO <- <(('p' / 'P') ('o' / 'O') Action80)> */
		nil,
		/* 126 FT_PE <- <(('p' / 'P') ('e' / 'E') Action81)> */
		nil,
		/* 127 FT_P <- <(('p' / 'P') Action82)> */
		nil,
		/* 128 FT_M <- <(('m' / 'M') Action83)> */
		nil,
		/* 130 Action0 <- <{ p.Emit() }> */
		nil,
		/* 131 Action1 <- <{ p.LD8() }> */
		nil,
		/* 132 Action2 <- <{ p.LD16() }> */
		nil,
		/* 133 Action3 <- <{ p.Push() }> */
		nil,
		/* 134 Action4 <- <{ p.Pop() }> */
		nil,
		/* 135 Action5 <- <{ p.Ex() }> */
		nil,
		/* 136 Action6 <- <{ p.Inc8() }> */
		nil,
		/* 137 Action7 <- <{ p.Inc8() }> */
		nil,
		/* 138 Action8 <- <{ p.Inc16() }> */
		nil,
		/* 139 Action9 <- <{ p.Dec8() }> */
		nil,
		/* 140 Action10 <- <{ p.Dec8() }> */
		nil,
		/* 141 Action11 <- <{ p.Dec16() }> */
		nil,
		/* 142 Action12 <- <{ p.Add16() }> */
		nil,
		/* 143 Action13 <- <{ p.Adc16() }> */
		nil,
		/* 144 Action14 <- <{ p.Sbc16() }> */
		nil,
		/* 145 Action15 <- <{ p.Dst8() }> */
		nil,
		/* 146 Action16 <- <{ p.Src8() }> */
		nil,
		/* 147 Action17 <- <{ p.Loc8() }> */
		nil,
		/* 148 Action18 <- <{ p.Loc8() }> */
		nil,
		nil,
		/* 150 Action19 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 151 Action20 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 152 Action21 <- <{ p.Dst16() }> */
		nil,
		/* 153 Action22 <- <{ p.Src16() }> */
		nil,
		/* 154 Action23 <- <{ p.Loc16() }> */
		nil,
		/* 155 Action24 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 156 Action25 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 157 Action26 <- <{ p.R16Contents() }> */
		nil,
		/* 158 Action27 <- <{ p.IR16Contents() }> */
		nil,
		/* 159 Action28 <- <{ p.NNContents() }> */
		nil,
		/* 160 Action29 <- <{ p.Accum("ADD") }> */
		nil,
		/* 161 Action30 <- <{ p.Accum("ADC") }> */
		nil,
		/* 162 Action31 <- <{ p.Accum("SUB") }> */
		nil,
		/* 163 Action32 <- <{ p.Accum("SBC") }> */
		nil,
		/* 164 Action33 <- <{ p.Accum("AND") }> */
		nil,
		/* 165 Action34 <- <{ p.Accum("XOR") }> */
		nil,
		/* 166 Action35 <- <{ p.Accum("OR") }> */
		nil,
		/* 167 Action36 <- <{ p.Accum("CP") }> */
		nil,
		/* 168 Action37 <- <{ p.Rot("RLC") }> */
		nil,
		/* 169 Action38 <- <{ p.Rot("RRC") }> */
		nil,
		/* 170 Action39 <- <{ p.Rot("RL") }> */
		nil,
		/* 171 Action40 <- <{ p.Rot("RR") }> */
		nil,
		/* 172 Action41 <- <{ p.Rot("SLA") }> */
		nil,
		/* 173 Action42 <- <{ p.Rot("SRA") }> */
		nil,
		/* 174 Action43 <- <{ p.Rot("SLL") }> */
		nil,
		/* 175 Action44 <- <{ p.Rot("SRL") }> */
		nil,
		/* 176 Action45 <- <{ p.Bit() }> */
		nil,
		/* 177 Action46 <- <{ p.Res() }> */
		nil,
		/* 178 Action47 <- <{ p.Set() }> */
		nil,
		/* 179 Action48 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 180 Action49 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 181 Action50 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 182 Action51 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 183 Action52 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 184 Action53 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 185 Action54 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 186 Action55 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 187 Action56 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 188 Action57 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 189 Action58 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 190 Action59 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 191 Action60 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 192 Action61 <- <{ p.Rst() }> */
		nil,
		/* 193 Action62 <- <{ p.Call() }> */
		nil,
		/* 194 Action63 <- <{ p.Ret() }> */
		nil,
		/* 195 Action64 <- <{ p.Jp() }> */
		nil,
		/* 196 Action65 <- <{ p.Jr() }> */
		nil,
		/* 197 Action66 <- <{ p.Djnz() }> */
		nil,
		/* 198 Action67 <- <{ p.In() }> */
		nil,
		/* 199 Action68 <- <{ p.Out() }> */
		nil,
		/* 200 Action69 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 201 Action70 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 202 Action71 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 203 Action72 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 204 Action73 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 205 Action74 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 206 Action75 <- <{ p.SignedDecimalByte(buffer[begin:end]) }> */
		nil,
		/* 207 Action76 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 208 Action77 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 209 Action78 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 210 Action79 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 211 Action80 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 212 Action81 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 213 Action82 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 214 Action83 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
