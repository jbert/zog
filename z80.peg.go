package zog

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint16

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
	ruleEDSimple
	ruleNeg
	ruleRetn
	ruleReti
	ruleRrd
	ruleRld
	ruleIm0
	ruleIm1
	ruleIm2
	ruleBlit
	ruleBlitIO
	ruleLdi
	ruleCpi
	ruleIni
	ruleOuti
	ruleLdd
	ruleCpd
	ruleInd
	ruleOutd
	ruleLdir
	ruleCpir
	ruleInir
	ruleOtir
	ruleLddr
	ruleCpdr
	ruleIndr
	ruleOtdr
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
	ruleAction84
	ruleAction85
	ruleAction86
	ruleAction87
	ruleAction88
	ruleAction89
	ruleAction90
	ruleAction91
	ruleAction92
	ruleAction93
	ruleAction94
	ruleAction95
	ruleAction96
	ruleAction97
	ruleAction98
	ruleAction99
	ruleAction100
	ruleAction101
	ruleAction102
	ruleAction103
	ruleAction104
	ruleAction105
	ruleAction106
	ruleAction107
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
	"EDSimple",
	"Neg",
	"Retn",
	"Reti",
	"Rrd",
	"Rld",
	"Im0",
	"Im1",
	"Im2",
	"Blit",
	"BlitIO",
	"Ldi",
	"Cpi",
	"Ini",
	"Outi",
	"Ldd",
	"Cpd",
	"Ind",
	"Outd",
	"Ldir",
	"Cpir",
	"Inir",
	"Otir",
	"Lddr",
	"Cpdr",
	"Indr",
	"Otdr",
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
	"Action84",
	"Action85",
	"Action86",
	"Action87",
	"Action88",
	"Action89",
	"Action90",
	"Action91",
	"Action92",
	"Action93",
	"Action94",
	"Action95",
	"Action96",
	"Action97",
	"Action98",
	"Action99",
	"Action100",
	"Action101",
	"Action102",
	"Action103",
	"Action104",
	"Action105",
	"Action106",
	"Action107",
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
	rules  [266]func() bool
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
			p.EDSimple(buffer[begin:end])
		case ruleAction62:
			p.EDSimple(buffer[begin:end])
		case ruleAction63:
			p.EDSimple(buffer[begin:end])
		case ruleAction64:
			p.EDSimple(buffer[begin:end])
		case ruleAction65:
			p.EDSimple(buffer[begin:end])
		case ruleAction66:
			p.EDSimple(buffer[begin:end])
		case ruleAction67:
			p.EDSimple(buffer[begin:end])
		case ruleAction68:
			p.EDSimple(buffer[begin:end])
		case ruleAction69:
			p.EDSimple(buffer[begin:end])
		case ruleAction70:
			p.EDSimple(buffer[begin:end])
		case ruleAction71:
			p.EDSimple(buffer[begin:end])
		case ruleAction72:
			p.EDSimple(buffer[begin:end])
		case ruleAction73:
			p.EDSimple(buffer[begin:end])
		case ruleAction74:
			p.EDSimple(buffer[begin:end])
		case ruleAction75:
			p.EDSimple(buffer[begin:end])
		case ruleAction76:
			p.EDSimple(buffer[begin:end])
		case ruleAction77:
			p.EDSimple(buffer[begin:end])
		case ruleAction78:
			p.EDSimple(buffer[begin:end])
		case ruleAction79:
			p.EDSimple(buffer[begin:end])
		case ruleAction80:
			p.EDSimple(buffer[begin:end])
		case ruleAction81:
			p.EDSimple(buffer[begin:end])
		case ruleAction82:
			p.EDSimple(buffer[begin:end])
		case ruleAction83:
			p.EDSimple(buffer[begin:end])
		case ruleAction84:
			p.EDSimple(buffer[begin:end])
		case ruleAction85:
			p.Rst()
		case ruleAction86:
			p.Call()
		case ruleAction87:
			p.Ret()
		case ruleAction88:
			p.Jp()
		case ruleAction89:
			p.Jr()
		case ruleAction90:
			p.Djnz()
		case ruleAction91:
			p.In()
		case ruleAction92:
			p.Out()
		case ruleAction93:
			p.Nhex(buffer[begin:end])
		case ruleAction94:
			p.Nhex(buffer[begin:end])
		case ruleAction95:
			p.Ndec(buffer[begin:end])
		case ruleAction96:
			p.NNhex(buffer[begin:end])
		case ruleAction97:
			p.NNhex(buffer[begin:end])
		case ruleAction98:
			p.ODigit(buffer[begin:end])
		case ruleAction99:
			p.SignedDecimalByte(buffer[begin:end])
		case ruleAction100:
			p.Conditional(Not{FT_Z})
		case ruleAction101:
			p.Conditional(FT_Z)
		case ruleAction102:
			p.Conditional(Not{FT_C})
		case ruleAction103:
			p.Conditional(FT_C)
		case ruleAction104:
			p.Conditional(FT_PO)
		case ruleAction105:
			p.Conditional(FT_PE)
		case ruleAction106:
			p.Conditional(FT_P)
		case ruleAction107:
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
													if buffer[position] != rune('e') {
														goto l330
													}
													position++
													goto l329
												l330:
													position, tokenIndex = position329, tokenIndex329
													if buffer[position] != rune('E') {
														goto l324
													}
													position++
												}
											l329:
												{
													position331, tokenIndex331 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l332
													}
													position++
													goto l331
												l332:
													position, tokenIndex = position331, tokenIndex331
													if buffer[position] != rune('T') {
														goto l324
													}
													position++
												}
											l331:
												{
													position333, tokenIndex333 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l334
													}
													position++
													goto l333
												l334:
													position, tokenIndex = position333, tokenIndex333
													if buffer[position] != rune('N') {
														goto l324
													}
													position++
												}
											l333:
												add(rulePegText, position326)
											}
											{
												add(ruleAction62, position)
											}
											add(ruleRetn, position325)
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
													if buffer[position] != rune('e') {
														goto l342
													}
													position++
													goto l341
												l342:
													position, tokenIndex = position341, tokenIndex341
													if buffer[position] != rune('E') {
														goto l336
													}
													position++
												}
											l341:
												{
													position343, tokenIndex343 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l344
													}
													position++
													goto l343
												l344:
													position, tokenIndex = position343, tokenIndex343
													if buffer[position] != rune('T') {
														goto l336
													}
													position++
												}
											l343:
												{
													position345, tokenIndex345 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l346
													}
													position++
													goto l345
												l346:
													position, tokenIndex = position345, tokenIndex345
													if buffer[position] != rune('I') {
														goto l336
													}
													position++
												}
											l345:
												add(rulePegText, position338)
											}
											{
												add(ruleAction63, position)
											}
											add(ruleReti, position337)
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
													if buffer[position] != rune('r') {
														goto l354
													}
													position++
													goto l353
												l354:
													position, tokenIndex = position353, tokenIndex353
													if buffer[position] != rune('R') {
														goto l348
													}
													position++
												}
											l353:
												{
													position355, tokenIndex355 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l356
													}
													position++
													goto l355
												l356:
													position, tokenIndex = position355, tokenIndex355
													if buffer[position] != rune('D') {
														goto l348
													}
													position++
												}
											l355:
												add(rulePegText, position350)
											}
											{
												add(ruleAction64, position)
											}
											add(ruleRrd, position349)
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
													if buffer[position] != rune('i') {
														goto l362
													}
													position++
													goto l361
												l362:
													position, tokenIndex = position361, tokenIndex361
													if buffer[position] != rune('I') {
														goto l358
													}
													position++
												}
											l361:
												{
													position363, tokenIndex363 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l364
													}
													position++
													goto l363
												l364:
													position, tokenIndex = position363, tokenIndex363
													if buffer[position] != rune('M') {
														goto l358
													}
													position++
												}
											l363:
												if buffer[position] != rune(' ') {
													goto l358
												}
												position++
												if buffer[position] != rune('0') {
													goto l358
												}
												position++
												add(rulePegText, position360)
											}
											{
												add(ruleAction66, position)
											}
											add(ruleIm0, position359)
										}
										goto l323
									l358:
										position, tokenIndex = position323, tokenIndex323
										{
											position367 := position
											{
												position368 := position
												{
													position369, tokenIndex369 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l370
													}
													position++
													goto l369
												l370:
													position, tokenIndex = position369, tokenIndex369
													if buffer[position] != rune('I') {
														goto l366
													}
													position++
												}
											l369:
												{
													position371, tokenIndex371 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l372
													}
													position++
													goto l371
												l372:
													position, tokenIndex = position371, tokenIndex371
													if buffer[position] != rune('M') {
														goto l366
													}
													position++
												}
											l371:
												if buffer[position] != rune(' ') {
													goto l366
												}
												position++
												if buffer[position] != rune('1') {
													goto l366
												}
												position++
												add(rulePegText, position368)
											}
											{
												add(ruleAction67, position)
											}
											add(ruleIm1, position367)
										}
										goto l323
									l366:
										position, tokenIndex = position323, tokenIndex323
										{
											position375 := position
											{
												position376 := position
												{
													position377, tokenIndex377 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l378
													}
													position++
													goto l377
												l378:
													position, tokenIndex = position377, tokenIndex377
													if buffer[position] != rune('I') {
														goto l374
													}
													position++
												}
											l377:
												{
													position379, tokenIndex379 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l380
													}
													position++
													goto l379
												l380:
													position, tokenIndex = position379, tokenIndex379
													if buffer[position] != rune('M') {
														goto l374
													}
													position++
												}
											l379:
												if buffer[position] != rune(' ') {
													goto l374
												}
												position++
												if buffer[position] != rune('2') {
													goto l374
												}
												position++
												add(rulePegText, position376)
											}
											{
												add(ruleAction68, position)
											}
											add(ruleIm2, position375)
										}
										goto l323
									l374:
										position, tokenIndex = position323, tokenIndex323
										{
											switch buffer[position] {
											case 'I', 'O', 'i', 'o':
												{
													position383 := position
													{
														position384, tokenIndex384 := position, tokenIndex
														{
															position386 := position
															{
																position387 := position
																{
																	position388, tokenIndex388 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l389
																	}
																	position++
																	goto l388
																l389:
																	position, tokenIndex = position388, tokenIndex388
																	if buffer[position] != rune('I') {
																		goto l385
																	}
																	position++
																}
															l388:
																{
																	position390, tokenIndex390 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l391
																	}
																	position++
																	goto l390
																l391:
																	position, tokenIndex = position390, tokenIndex390
																	if buffer[position] != rune('N') {
																		goto l385
																	}
																	position++
																}
															l390:
																{
																	position392, tokenIndex392 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l393
																	}
																	position++
																	goto l392
																l393:
																	position, tokenIndex = position392, tokenIndex392
																	if buffer[position] != rune('I') {
																		goto l385
																	}
																	position++
																}
															l392:
																{
																	position394, tokenIndex394 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l395
																	}
																	position++
																	goto l394
																l395:
																	position, tokenIndex = position394, tokenIndex394
																	if buffer[position] != rune('R') {
																		goto l385
																	}
																	position++
																}
															l394:
																add(rulePegText, position387)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleInir, position386)
														}
														goto l384
													l385:
														position, tokenIndex = position384, tokenIndex384
														{
															position398 := position
															{
																position399 := position
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
																		goto l397
																	}
																	position++
																}
															l400:
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
																		goto l397
																	}
																	position++
																}
															l402:
																{
																	position404, tokenIndex404 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l405
																	}
																	position++
																	goto l404
																l405:
																	position, tokenIndex = position404, tokenIndex404
																	if buffer[position] != rune('I') {
																		goto l397
																	}
																	position++
																}
															l404:
																add(rulePegText, position399)
															}
															{
																add(ruleAction71, position)
															}
															add(ruleIni, position398)
														}
														goto l384
													l397:
														position, tokenIndex = position384, tokenIndex384
														{
															position408 := position
															{
																position409 := position
																{
																	position410, tokenIndex410 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l411
																	}
																	position++
																	goto l410
																l411:
																	position, tokenIndex = position410, tokenIndex410
																	if buffer[position] != rune('O') {
																		goto l407
																	}
																	position++
																}
															l410:
																{
																	position412, tokenIndex412 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l413
																	}
																	position++
																	goto l412
																l413:
																	position, tokenIndex = position412, tokenIndex412
																	if buffer[position] != rune('T') {
																		goto l407
																	}
																	position++
																}
															l412:
																{
																	position414, tokenIndex414 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l415
																	}
																	position++
																	goto l414
																l415:
																	position, tokenIndex = position414, tokenIndex414
																	if buffer[position] != rune('I') {
																		goto l407
																	}
																	position++
																}
															l414:
																{
																	position416, tokenIndex416 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l417
																	}
																	position++
																	goto l416
																l417:
																	position, tokenIndex = position416, tokenIndex416
																	if buffer[position] != rune('R') {
																		goto l407
																	}
																	position++
																}
															l416:
																add(rulePegText, position409)
															}
															{
																add(ruleAction80, position)
															}
															add(ruleOtir, position408)
														}
														goto l384
													l407:
														position, tokenIndex = position384, tokenIndex384
														{
															position420 := position
															{
																position421 := position
																{
																	position422, tokenIndex422 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l423
																	}
																	position++
																	goto l422
																l423:
																	position, tokenIndex = position422, tokenIndex422
																	if buffer[position] != rune('O') {
																		goto l419
																	}
																	position++
																}
															l422:
																{
																	position424, tokenIndex424 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l425
																	}
																	position++
																	goto l424
																l425:
																	position, tokenIndex = position424, tokenIndex424
																	if buffer[position] != rune('U') {
																		goto l419
																	}
																	position++
																}
															l424:
																{
																	position426, tokenIndex426 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l427
																	}
																	position++
																	goto l426
																l427:
																	position, tokenIndex = position426, tokenIndex426
																	if buffer[position] != rune('T') {
																		goto l419
																	}
																	position++
																}
															l426:
																{
																	position428, tokenIndex428 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l429
																	}
																	position++
																	goto l428
																l429:
																	position, tokenIndex = position428, tokenIndex428
																	if buffer[position] != rune('I') {
																		goto l419
																	}
																	position++
																}
															l428:
																add(rulePegText, position421)
															}
															{
																add(ruleAction72, position)
															}
															add(ruleOuti, position420)
														}
														goto l384
													l419:
														position, tokenIndex = position384, tokenIndex384
														{
															position432 := position
															{
																position433 := position
																{
																	position434, tokenIndex434 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l435
																	}
																	position++
																	goto l434
																l435:
																	position, tokenIndex = position434, tokenIndex434
																	if buffer[position] != rune('I') {
																		goto l431
																	}
																	position++
																}
															l434:
																{
																	position436, tokenIndex436 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l437
																	}
																	position++
																	goto l436
																l437:
																	position, tokenIndex = position436, tokenIndex436
																	if buffer[position] != rune('N') {
																		goto l431
																	}
																	position++
																}
															l436:
																{
																	position438, tokenIndex438 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l439
																	}
																	position++
																	goto l438
																l439:
																	position, tokenIndex = position438, tokenIndex438
																	if buffer[position] != rune('D') {
																		goto l431
																	}
																	position++
																}
															l438:
																{
																	position440, tokenIndex440 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l441
																	}
																	position++
																	goto l440
																l441:
																	position, tokenIndex = position440, tokenIndex440
																	if buffer[position] != rune('R') {
																		goto l431
																	}
																	position++
																}
															l440:
																add(rulePegText, position433)
															}
															{
																add(ruleAction83, position)
															}
															add(ruleIndr, position432)
														}
														goto l384
													l431:
														position, tokenIndex = position384, tokenIndex384
														{
															position444 := position
															{
																position445 := position
																{
																	position446, tokenIndex446 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l447
																	}
																	position++
																	goto l446
																l447:
																	position, tokenIndex = position446, tokenIndex446
																	if buffer[position] != rune('I') {
																		goto l443
																	}
																	position++
																}
															l446:
																{
																	position448, tokenIndex448 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l449
																	}
																	position++
																	goto l448
																l449:
																	position, tokenIndex = position448, tokenIndex448
																	if buffer[position] != rune('N') {
																		goto l443
																	}
																	position++
																}
															l448:
																{
																	position450, tokenIndex450 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l451
																	}
																	position++
																	goto l450
																l451:
																	position, tokenIndex = position450, tokenIndex450
																	if buffer[position] != rune('D') {
																		goto l443
																	}
																	position++
																}
															l450:
																add(rulePegText, position445)
															}
															{
																add(ruleAction75, position)
															}
															add(ruleInd, position444)
														}
														goto l384
													l443:
														position, tokenIndex = position384, tokenIndex384
														{
															position454 := position
															{
																position455 := position
																{
																	position456, tokenIndex456 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l457
																	}
																	position++
																	goto l456
																l457:
																	position, tokenIndex = position456, tokenIndex456
																	if buffer[position] != rune('O') {
																		goto l453
																	}
																	position++
																}
															l456:
																{
																	position458, tokenIndex458 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l459
																	}
																	position++
																	goto l458
																l459:
																	position, tokenIndex = position458, tokenIndex458
																	if buffer[position] != rune('T') {
																		goto l453
																	}
																	position++
																}
															l458:
																{
																	position460, tokenIndex460 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l461
																	}
																	position++
																	goto l460
																l461:
																	position, tokenIndex = position460, tokenIndex460
																	if buffer[position] != rune('D') {
																		goto l453
																	}
																	position++
																}
															l460:
																{
																	position462, tokenIndex462 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l463
																	}
																	position++
																	goto l462
																l463:
																	position, tokenIndex = position462, tokenIndex462
																	if buffer[position] != rune('R') {
																		goto l453
																	}
																	position++
																}
															l462:
																add(rulePegText, position455)
															}
															{
																add(ruleAction84, position)
															}
															add(ruleOtdr, position454)
														}
														goto l384
													l453:
														position, tokenIndex = position384, tokenIndex384
														{
															position465 := position
															{
																position466 := position
																{
																	position467, tokenIndex467 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l468
																	}
																	position++
																	goto l467
																l468:
																	position, tokenIndex = position467, tokenIndex467
																	if buffer[position] != rune('O') {
																		goto l321
																	}
																	position++
																}
															l467:
																{
																	position469, tokenIndex469 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l470
																	}
																	position++
																	goto l469
																l470:
																	position, tokenIndex = position469, tokenIndex469
																	if buffer[position] != rune('U') {
																		goto l321
																	}
																	position++
																}
															l469:
																{
																	position471, tokenIndex471 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l472
																	}
																	position++
																	goto l471
																l472:
																	position, tokenIndex = position471, tokenIndex471
																	if buffer[position] != rune('T') {
																		goto l321
																	}
																	position++
																}
															l471:
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
																		goto l321
																	}
																	position++
																}
															l473:
																add(rulePegText, position466)
															}
															{
																add(ruleAction76, position)
															}
															add(ruleOutd, position465)
														}
													}
												l384:
													add(ruleBlitIO, position383)
												}
												break
											case 'R', 'r':
												{
													position476 := position
													{
														position477 := position
														{
															position478, tokenIndex478 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l479
															}
															position++
															goto l478
														l479:
															position, tokenIndex = position478, tokenIndex478
															if buffer[position] != rune('R') {
																goto l321
															}
															position++
														}
													l478:
														{
															position480, tokenIndex480 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l481
															}
															position++
															goto l480
														l481:
															position, tokenIndex = position480, tokenIndex480
															if buffer[position] != rune('L') {
																goto l321
															}
															position++
														}
													l480:
														{
															position482, tokenIndex482 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l483
															}
															position++
															goto l482
														l483:
															position, tokenIndex = position482, tokenIndex482
															if buffer[position] != rune('D') {
																goto l321
															}
															position++
														}
													l482:
														add(rulePegText, position477)
													}
													{
														add(ruleAction65, position)
													}
													add(ruleRld, position476)
												}
												break
											case 'N', 'n':
												{
													position485 := position
													{
														position486 := position
														{
															position487, tokenIndex487 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l488
															}
															position++
															goto l487
														l488:
															position, tokenIndex = position487, tokenIndex487
															if buffer[position] != rune('N') {
																goto l321
															}
															position++
														}
													l487:
														{
															position489, tokenIndex489 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l490
															}
															position++
															goto l489
														l490:
															position, tokenIndex = position489, tokenIndex489
															if buffer[position] != rune('E') {
																goto l321
															}
															position++
														}
													l489:
														{
															position491, tokenIndex491 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l492
															}
															position++
															goto l491
														l492:
															position, tokenIndex = position491, tokenIndex491
															if buffer[position] != rune('G') {
																goto l321
															}
															position++
														}
													l491:
														add(rulePegText, position486)
													}
													{
														add(ruleAction61, position)
													}
													add(ruleNeg, position485)
												}
												break
											default:
												{
													position494 := position
													{
														position495, tokenIndex495 := position, tokenIndex
														{
															position497 := position
															{
																position498 := position
																{
																	position499, tokenIndex499 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l500
																	}
																	position++
																	goto l499
																l500:
																	position, tokenIndex = position499, tokenIndex499
																	if buffer[position] != rune('L') {
																		goto l496
																	}
																	position++
																}
															l499:
																{
																	position501, tokenIndex501 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l502
																	}
																	position++
																	goto l501
																l502:
																	position, tokenIndex = position501, tokenIndex501
																	if buffer[position] != rune('D') {
																		goto l496
																	}
																	position++
																}
															l501:
																{
																	position503, tokenIndex503 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l504
																	}
																	position++
																	goto l503
																l504:
																	position, tokenIndex = position503, tokenIndex503
																	if buffer[position] != rune('I') {
																		goto l496
																	}
																	position++
																}
															l503:
																{
																	position505, tokenIndex505 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l506
																	}
																	position++
																	goto l505
																l506:
																	position, tokenIndex = position505, tokenIndex505
																	if buffer[position] != rune('R') {
																		goto l496
																	}
																	position++
																}
															l505:
																add(rulePegText, position498)
															}
															{
																add(ruleAction77, position)
															}
															add(ruleLdir, position497)
														}
														goto l495
													l496:
														position, tokenIndex = position495, tokenIndex495
														{
															position509 := position
															{
																position510 := position
																{
																	position511, tokenIndex511 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l512
																	}
																	position++
																	goto l511
																l512:
																	position, tokenIndex = position511, tokenIndex511
																	if buffer[position] != rune('L') {
																		goto l508
																	}
																	position++
																}
															l511:
																{
																	position513, tokenIndex513 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l514
																	}
																	position++
																	goto l513
																l514:
																	position, tokenIndex = position513, tokenIndex513
																	if buffer[position] != rune('D') {
																		goto l508
																	}
																	position++
																}
															l513:
																{
																	position515, tokenIndex515 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l516
																	}
																	position++
																	goto l515
																l516:
																	position, tokenIndex = position515, tokenIndex515
																	if buffer[position] != rune('I') {
																		goto l508
																	}
																	position++
																}
															l515:
																add(rulePegText, position510)
															}
															{
																add(ruleAction69, position)
															}
															add(ruleLdi, position509)
														}
														goto l495
													l508:
														position, tokenIndex = position495, tokenIndex495
														{
															position519 := position
															{
																position520 := position
																{
																	position521, tokenIndex521 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l522
																	}
																	position++
																	goto l521
																l522:
																	position, tokenIndex = position521, tokenIndex521
																	if buffer[position] != rune('C') {
																		goto l518
																	}
																	position++
																}
															l521:
																{
																	position523, tokenIndex523 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l524
																	}
																	position++
																	goto l523
																l524:
																	position, tokenIndex = position523, tokenIndex523
																	if buffer[position] != rune('P') {
																		goto l518
																	}
																	position++
																}
															l523:
																{
																	position525, tokenIndex525 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l526
																	}
																	position++
																	goto l525
																l526:
																	position, tokenIndex = position525, tokenIndex525
																	if buffer[position] != rune('I') {
																		goto l518
																	}
																	position++
																}
															l525:
																{
																	position527, tokenIndex527 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l528
																	}
																	position++
																	goto l527
																l528:
																	position, tokenIndex = position527, tokenIndex527
																	if buffer[position] != rune('R') {
																		goto l518
																	}
																	position++
																}
															l527:
																add(rulePegText, position520)
															}
															{
																add(ruleAction78, position)
															}
															add(ruleCpir, position519)
														}
														goto l495
													l518:
														position, tokenIndex = position495, tokenIndex495
														{
															position531 := position
															{
																position532 := position
																{
																	position533, tokenIndex533 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l534
																	}
																	position++
																	goto l533
																l534:
																	position, tokenIndex = position533, tokenIndex533
																	if buffer[position] != rune('C') {
																		goto l530
																	}
																	position++
																}
															l533:
																{
																	position535, tokenIndex535 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l536
																	}
																	position++
																	goto l535
																l536:
																	position, tokenIndex = position535, tokenIndex535
																	if buffer[position] != rune('P') {
																		goto l530
																	}
																	position++
																}
															l535:
																{
																	position537, tokenIndex537 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l538
																	}
																	position++
																	goto l537
																l538:
																	position, tokenIndex = position537, tokenIndex537
																	if buffer[position] != rune('I') {
																		goto l530
																	}
																	position++
																}
															l537:
																add(rulePegText, position532)
															}
															{
																add(ruleAction70, position)
															}
															add(ruleCpi, position531)
														}
														goto l495
													l530:
														position, tokenIndex = position495, tokenIndex495
														{
															position541 := position
															{
																position542 := position
																{
																	position543, tokenIndex543 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l544
																	}
																	position++
																	goto l543
																l544:
																	position, tokenIndex = position543, tokenIndex543
																	if buffer[position] != rune('L') {
																		goto l540
																	}
																	position++
																}
															l543:
																{
																	position545, tokenIndex545 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l546
																	}
																	position++
																	goto l545
																l546:
																	position, tokenIndex = position545, tokenIndex545
																	if buffer[position] != rune('D') {
																		goto l540
																	}
																	position++
																}
															l545:
																{
																	position547, tokenIndex547 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l548
																	}
																	position++
																	goto l547
																l548:
																	position, tokenIndex = position547, tokenIndex547
																	if buffer[position] != rune('D') {
																		goto l540
																	}
																	position++
																}
															l547:
																{
																	position549, tokenIndex549 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l550
																	}
																	position++
																	goto l549
																l550:
																	position, tokenIndex = position549, tokenIndex549
																	if buffer[position] != rune('R') {
																		goto l540
																	}
																	position++
																}
															l549:
																add(rulePegText, position542)
															}
															{
																add(ruleAction81, position)
															}
															add(ruleLddr, position541)
														}
														goto l495
													l540:
														position, tokenIndex = position495, tokenIndex495
														{
															position553 := position
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
																		goto l552
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
																		goto l552
																	}
																	position++
																}
															l557:
																{
																	position559, tokenIndex559 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l560
																	}
																	position++
																	goto l559
																l560:
																	position, tokenIndex = position559, tokenIndex559
																	if buffer[position] != rune('D') {
																		goto l552
																	}
																	position++
																}
															l559:
																add(rulePegText, position554)
															}
															{
																add(ruleAction73, position)
															}
															add(ruleLdd, position553)
														}
														goto l495
													l552:
														position, tokenIndex = position495, tokenIndex495
														{
															position563 := position
															{
																position564 := position
																{
																	position565, tokenIndex565 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l566
																	}
																	position++
																	goto l565
																l566:
																	position, tokenIndex = position565, tokenIndex565
																	if buffer[position] != rune('C') {
																		goto l562
																	}
																	position++
																}
															l565:
																{
																	position567, tokenIndex567 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l568
																	}
																	position++
																	goto l567
																l568:
																	position, tokenIndex = position567, tokenIndex567
																	if buffer[position] != rune('P') {
																		goto l562
																	}
																	position++
																}
															l567:
																{
																	position569, tokenIndex569 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l570
																	}
																	position++
																	goto l569
																l570:
																	position, tokenIndex = position569, tokenIndex569
																	if buffer[position] != rune('D') {
																		goto l562
																	}
																	position++
																}
															l569:
																{
																	position571, tokenIndex571 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l572
																	}
																	position++
																	goto l571
																l572:
																	position, tokenIndex = position571, tokenIndex571
																	if buffer[position] != rune('R') {
																		goto l562
																	}
																	position++
																}
															l571:
																add(rulePegText, position564)
															}
															{
																add(ruleAction82, position)
															}
															add(ruleCpdr, position563)
														}
														goto l495
													l562:
														position, tokenIndex = position495, tokenIndex495
														{
															position574 := position
															{
																position575 := position
																{
																	position576, tokenIndex576 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l577
																	}
																	position++
																	goto l576
																l577:
																	position, tokenIndex = position576, tokenIndex576
																	if buffer[position] != rune('C') {
																		goto l321
																	}
																	position++
																}
															l576:
																{
																	position578, tokenIndex578 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l579
																	}
																	position++
																	goto l578
																l579:
																	position, tokenIndex = position578, tokenIndex578
																	if buffer[position] != rune('P') {
																		goto l321
																	}
																	position++
																}
															l578:
																{
																	position580, tokenIndex580 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l581
																	}
																	position++
																	goto l580
																l581:
																	position, tokenIndex = position580, tokenIndex580
																	if buffer[position] != rune('D') {
																		goto l321
																	}
																	position++
																}
															l580:
																add(rulePegText, position575)
															}
															{
																add(ruleAction74, position)
															}
															add(ruleCpd, position574)
														}
													}
												l495:
													add(ruleBlit, position494)
												}
												break
											}
										}

									}
								l323:
									add(ruleEDSimple, position322)
								}
								goto l13
							l321:
								position, tokenIndex = position13, tokenIndex13
								{
									position584 := position
									{
										position585, tokenIndex585 := position, tokenIndex
										{
											position587 := position
											{
												position588 := position
												{
													position589, tokenIndex589 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l590
													}
													position++
													goto l589
												l590:
													position, tokenIndex = position589, tokenIndex589
													if buffer[position] != rune('R') {
														goto l586
													}
													position++
												}
											l589:
												{
													position591, tokenIndex591 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l592
													}
													position++
													goto l591
												l592:
													position, tokenIndex = position591, tokenIndex591
													if buffer[position] != rune('L') {
														goto l586
													}
													position++
												}
											l591:
												{
													position593, tokenIndex593 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l594
													}
													position++
													goto l593
												l594:
													position, tokenIndex = position593, tokenIndex593
													if buffer[position] != rune('C') {
														goto l586
													}
													position++
												}
											l593:
												{
													position595, tokenIndex595 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l596
													}
													position++
													goto l595
												l596:
													position, tokenIndex = position595, tokenIndex595
													if buffer[position] != rune('A') {
														goto l586
													}
													position++
												}
											l595:
												add(rulePegText, position588)
											}
											{
												add(ruleAction50, position)
											}
											add(ruleRlca, position587)
										}
										goto l585
									l586:
										position, tokenIndex = position585, tokenIndex585
										{
											position599 := position
											{
												position600 := position
												{
													position601, tokenIndex601 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l602
													}
													position++
													goto l601
												l602:
													position, tokenIndex = position601, tokenIndex601
													if buffer[position] != rune('R') {
														goto l598
													}
													position++
												}
											l601:
												{
													position603, tokenIndex603 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l604
													}
													position++
													goto l603
												l604:
													position, tokenIndex = position603, tokenIndex603
													if buffer[position] != rune('R') {
														goto l598
													}
													position++
												}
											l603:
												{
													position605, tokenIndex605 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l606
													}
													position++
													goto l605
												l606:
													position, tokenIndex = position605, tokenIndex605
													if buffer[position] != rune('C') {
														goto l598
													}
													position++
												}
											l605:
												{
													position607, tokenIndex607 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l608
													}
													position++
													goto l607
												l608:
													position, tokenIndex = position607, tokenIndex607
													if buffer[position] != rune('A') {
														goto l598
													}
													position++
												}
											l607:
												add(rulePegText, position600)
											}
											{
												add(ruleAction51, position)
											}
											add(ruleRrca, position599)
										}
										goto l585
									l598:
										position, tokenIndex = position585, tokenIndex585
										{
											position611 := position
											{
												position612 := position
												{
													position613, tokenIndex613 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l614
													}
													position++
													goto l613
												l614:
													position, tokenIndex = position613, tokenIndex613
													if buffer[position] != rune('R') {
														goto l610
													}
													position++
												}
											l613:
												{
													position615, tokenIndex615 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l616
													}
													position++
													goto l615
												l616:
													position, tokenIndex = position615, tokenIndex615
													if buffer[position] != rune('L') {
														goto l610
													}
													position++
												}
											l615:
												{
													position617, tokenIndex617 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l618
													}
													position++
													goto l617
												l618:
													position, tokenIndex = position617, tokenIndex617
													if buffer[position] != rune('A') {
														goto l610
													}
													position++
												}
											l617:
												add(rulePegText, position612)
											}
											{
												add(ruleAction52, position)
											}
											add(ruleRla, position611)
										}
										goto l585
									l610:
										position, tokenIndex = position585, tokenIndex585
										{
											position621 := position
											{
												position622 := position
												{
													position623, tokenIndex623 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l624
													}
													position++
													goto l623
												l624:
													position, tokenIndex = position623, tokenIndex623
													if buffer[position] != rune('D') {
														goto l620
													}
													position++
												}
											l623:
												{
													position625, tokenIndex625 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l626
													}
													position++
													goto l625
												l626:
													position, tokenIndex = position625, tokenIndex625
													if buffer[position] != rune('A') {
														goto l620
													}
													position++
												}
											l625:
												{
													position627, tokenIndex627 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l628
													}
													position++
													goto l627
												l628:
													position, tokenIndex = position627, tokenIndex627
													if buffer[position] != rune('A') {
														goto l620
													}
													position++
												}
											l627:
												add(rulePegText, position622)
											}
											{
												add(ruleAction54, position)
											}
											add(ruleDaa, position621)
										}
										goto l585
									l620:
										position, tokenIndex = position585, tokenIndex585
										{
											position631 := position
											{
												position632 := position
												{
													position633, tokenIndex633 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l634
													}
													position++
													goto l633
												l634:
													position, tokenIndex = position633, tokenIndex633
													if buffer[position] != rune('C') {
														goto l630
													}
													position++
												}
											l633:
												{
													position635, tokenIndex635 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l636
													}
													position++
													goto l635
												l636:
													position, tokenIndex = position635, tokenIndex635
													if buffer[position] != rune('P') {
														goto l630
													}
													position++
												}
											l635:
												{
													position637, tokenIndex637 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l638
													}
													position++
													goto l637
												l638:
													position, tokenIndex = position637, tokenIndex637
													if buffer[position] != rune('L') {
														goto l630
													}
													position++
												}
											l637:
												add(rulePegText, position632)
											}
											{
												add(ruleAction55, position)
											}
											add(ruleCpl, position631)
										}
										goto l585
									l630:
										position, tokenIndex = position585, tokenIndex585
										{
											position641 := position
											{
												position642 := position
												{
													position643, tokenIndex643 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l644
													}
													position++
													goto l643
												l644:
													position, tokenIndex = position643, tokenIndex643
													if buffer[position] != rune('E') {
														goto l640
													}
													position++
												}
											l643:
												{
													position645, tokenIndex645 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l646
													}
													position++
													goto l645
												l646:
													position, tokenIndex = position645, tokenIndex645
													if buffer[position] != rune('X') {
														goto l640
													}
													position++
												}
											l645:
												{
													position647, tokenIndex647 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l648
													}
													position++
													goto l647
												l648:
													position, tokenIndex = position647, tokenIndex647
													if buffer[position] != rune('X') {
														goto l640
													}
													position++
												}
											l647:
												add(rulePegText, position642)
											}
											{
												add(ruleAction58, position)
											}
											add(ruleExx, position641)
										}
										goto l585
									l640:
										position, tokenIndex = position585, tokenIndex585
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position651 := position
													{
														position652 := position
														{
															position653, tokenIndex653 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l654
															}
															position++
															goto l653
														l654:
															position, tokenIndex = position653, tokenIndex653
															if buffer[position] != rune('E') {
																goto l583
															}
															position++
														}
													l653:
														{
															position655, tokenIndex655 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l656
															}
															position++
															goto l655
														l656:
															position, tokenIndex = position655, tokenIndex655
															if buffer[position] != rune('I') {
																goto l583
															}
															position++
														}
													l655:
														add(rulePegText, position652)
													}
													{
														add(ruleAction60, position)
													}
													add(ruleEi, position651)
												}
												break
											case 'D', 'd':
												{
													position658 := position
													{
														position659 := position
														{
															position660, tokenIndex660 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l661
															}
															position++
															goto l660
														l661:
															position, tokenIndex = position660, tokenIndex660
															if buffer[position] != rune('D') {
																goto l583
															}
															position++
														}
													l660:
														{
															position662, tokenIndex662 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l663
															}
															position++
															goto l662
														l663:
															position, tokenIndex = position662, tokenIndex662
															if buffer[position] != rune('I') {
																goto l583
															}
															position++
														}
													l662:
														add(rulePegText, position659)
													}
													{
														add(ruleAction59, position)
													}
													add(ruleDi, position658)
												}
												break
											case 'C', 'c':
												{
													position665 := position
													{
														position666 := position
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
																goto l583
															}
															position++
														}
													l667:
														{
															position669, tokenIndex669 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l670
															}
															position++
															goto l669
														l670:
															position, tokenIndex = position669, tokenIndex669
															if buffer[position] != rune('C') {
																goto l583
															}
															position++
														}
													l669:
														{
															position671, tokenIndex671 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l672
															}
															position++
															goto l671
														l672:
															position, tokenIndex = position671, tokenIndex671
															if buffer[position] != rune('F') {
																goto l583
															}
															position++
														}
													l671:
														add(rulePegText, position666)
													}
													{
														add(ruleAction57, position)
													}
													add(ruleCcf, position665)
												}
												break
											case 'S', 's':
												{
													position674 := position
													{
														position675 := position
														{
															position676, tokenIndex676 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l677
															}
															position++
															goto l676
														l677:
															position, tokenIndex = position676, tokenIndex676
															if buffer[position] != rune('S') {
																goto l583
															}
															position++
														}
													l676:
														{
															position678, tokenIndex678 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l679
															}
															position++
															goto l678
														l679:
															position, tokenIndex = position678, tokenIndex678
															if buffer[position] != rune('C') {
																goto l583
															}
															position++
														}
													l678:
														{
															position680, tokenIndex680 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l681
															}
															position++
															goto l680
														l681:
															position, tokenIndex = position680, tokenIndex680
															if buffer[position] != rune('F') {
																goto l583
															}
															position++
														}
													l680:
														add(rulePegText, position675)
													}
													{
														add(ruleAction56, position)
													}
													add(ruleScf, position674)
												}
												break
											case 'R', 'r':
												{
													position683 := position
													{
														position684 := position
														{
															position685, tokenIndex685 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l686
															}
															position++
															goto l685
														l686:
															position, tokenIndex = position685, tokenIndex685
															if buffer[position] != rune('R') {
																goto l583
															}
															position++
														}
													l685:
														{
															position687, tokenIndex687 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l688
															}
															position++
															goto l687
														l688:
															position, tokenIndex = position687, tokenIndex687
															if buffer[position] != rune('R') {
																goto l583
															}
															position++
														}
													l687:
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
																goto l583
															}
															position++
														}
													l689:
														add(rulePegText, position684)
													}
													{
														add(ruleAction53, position)
													}
													add(ruleRra, position683)
												}
												break
											case 'H', 'h':
												{
													position692 := position
													{
														position693 := position
														{
															position694, tokenIndex694 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l695
															}
															position++
															goto l694
														l695:
															position, tokenIndex = position694, tokenIndex694
															if buffer[position] != rune('H') {
																goto l583
															}
															position++
														}
													l694:
														{
															position696, tokenIndex696 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l697
															}
															position++
															goto l696
														l697:
															position, tokenIndex = position696, tokenIndex696
															if buffer[position] != rune('A') {
																goto l583
															}
															position++
														}
													l696:
														{
															position698, tokenIndex698 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l699
															}
															position++
															goto l698
														l699:
															position, tokenIndex = position698, tokenIndex698
															if buffer[position] != rune('L') {
																goto l583
															}
															position++
														}
													l698:
														{
															position700, tokenIndex700 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l701
															}
															position++
															goto l700
														l701:
															position, tokenIndex = position700, tokenIndex700
															if buffer[position] != rune('T') {
																goto l583
															}
															position++
														}
													l700:
														add(rulePegText, position693)
													}
													{
														add(ruleAction49, position)
													}
													add(ruleHalt, position692)
												}
												break
											default:
												{
													position703 := position
													{
														position704 := position
														{
															position705, tokenIndex705 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l706
															}
															position++
															goto l705
														l706:
															position, tokenIndex = position705, tokenIndex705
															if buffer[position] != rune('N') {
																goto l583
															}
															position++
														}
													l705:
														{
															position707, tokenIndex707 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l708
															}
															position++
															goto l707
														l708:
															position, tokenIndex = position707, tokenIndex707
															if buffer[position] != rune('O') {
																goto l583
															}
															position++
														}
													l707:
														{
															position709, tokenIndex709 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l710
															}
															position++
															goto l709
														l710:
															position, tokenIndex = position709, tokenIndex709
															if buffer[position] != rune('P') {
																goto l583
															}
															position++
														}
													l709:
														add(rulePegText, position704)
													}
													{
														add(ruleAction48, position)
													}
													add(ruleNop, position703)
												}
												break
											}
										}

									}
								l585:
									add(ruleSimple, position584)
								}
								goto l13
							l583:
								position, tokenIndex = position13, tokenIndex13
								{
									position713 := position
									{
										position714, tokenIndex714 := position, tokenIndex
										{
											position716 := position
											{
												position717, tokenIndex717 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l718
												}
												position++
												goto l717
											l718:
												position, tokenIndex = position717, tokenIndex717
												if buffer[position] != rune('R') {
													goto l715
												}
												position++
											}
										l717:
											{
												position719, tokenIndex719 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l720
												}
												position++
												goto l719
											l720:
												position, tokenIndex = position719, tokenIndex719
												if buffer[position] != rune('S') {
													goto l715
												}
												position++
											}
										l719:
											{
												position721, tokenIndex721 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l722
												}
												position++
												goto l721
											l722:
												position, tokenIndex = position721, tokenIndex721
												if buffer[position] != rune('T') {
													goto l715
												}
												position++
											}
										l721:
											if !_rules[rulews]() {
												goto l715
											}
											if !_rules[rulen]() {
												goto l715
											}
											{
												add(ruleAction85, position)
											}
											add(ruleRst, position716)
										}
										goto l714
									l715:
										position, tokenIndex = position714, tokenIndex714
										{
											position725 := position
											{
												position726, tokenIndex726 := position, tokenIndex
												if buffer[position] != rune('j') {
													goto l727
												}
												position++
												goto l726
											l727:
												position, tokenIndex = position726, tokenIndex726
												if buffer[position] != rune('J') {
													goto l724
												}
												position++
											}
										l726:
											{
												position728, tokenIndex728 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l729
												}
												position++
												goto l728
											l729:
												position, tokenIndex = position728, tokenIndex728
												if buffer[position] != rune('P') {
													goto l724
												}
												position++
											}
										l728:
											if !_rules[rulews]() {
												goto l724
											}
											{
												position730, tokenIndex730 := position, tokenIndex
												if !_rules[rulecc]() {
													goto l730
												}
												if !_rules[rulesep]() {
													goto l730
												}
												goto l731
											l730:
												position, tokenIndex = position730, tokenIndex730
											}
										l731:
											if !_rules[ruleSrc16]() {
												goto l724
											}
											{
												add(ruleAction88, position)
											}
											add(ruleJp, position725)
										}
										goto l714
									l724:
										position, tokenIndex = position714, tokenIndex714
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position734 := position
													{
														position735, tokenIndex735 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l736
														}
														position++
														goto l735
													l736:
														position, tokenIndex = position735, tokenIndex735
														if buffer[position] != rune('D') {
															goto l712
														}
														position++
													}
												l735:
													{
														position737, tokenIndex737 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l738
														}
														position++
														goto l737
													l738:
														position, tokenIndex = position737, tokenIndex737
														if buffer[position] != rune('J') {
															goto l712
														}
														position++
													}
												l737:
													{
														position739, tokenIndex739 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l740
														}
														position++
														goto l739
													l740:
														position, tokenIndex = position739, tokenIndex739
														if buffer[position] != rune('N') {
															goto l712
														}
														position++
													}
												l739:
													{
														position741, tokenIndex741 := position, tokenIndex
														if buffer[position] != rune('z') {
															goto l742
														}
														position++
														goto l741
													l742:
														position, tokenIndex = position741, tokenIndex741
														if buffer[position] != rune('Z') {
															goto l712
														}
														position++
													}
												l741:
													if !_rules[rulews]() {
														goto l712
													}
													if !_rules[ruledisp]() {
														goto l712
													}
													{
														add(ruleAction90, position)
													}
													add(ruleDjnz, position734)
												}
												break
											case 'J', 'j':
												{
													position744 := position
													{
														position745, tokenIndex745 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l746
														}
														position++
														goto l745
													l746:
														position, tokenIndex = position745, tokenIndex745
														if buffer[position] != rune('J') {
															goto l712
														}
														position++
													}
												l745:
													{
														position747, tokenIndex747 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l748
														}
														position++
														goto l747
													l748:
														position, tokenIndex = position747, tokenIndex747
														if buffer[position] != rune('R') {
															goto l712
														}
														position++
													}
												l747:
													if !_rules[rulews]() {
														goto l712
													}
													{
														position749, tokenIndex749 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l749
														}
														if !_rules[rulesep]() {
															goto l749
														}
														goto l750
													l749:
														position, tokenIndex = position749, tokenIndex749
													}
												l750:
													if !_rules[ruledisp]() {
														goto l712
													}
													{
														add(ruleAction89, position)
													}
													add(ruleJr, position744)
												}
												break
											case 'R', 'r':
												{
													position752 := position
													{
														position753, tokenIndex753 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l754
														}
														position++
														goto l753
													l754:
														position, tokenIndex = position753, tokenIndex753
														if buffer[position] != rune('R') {
															goto l712
														}
														position++
													}
												l753:
													{
														position755, tokenIndex755 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l756
														}
														position++
														goto l755
													l756:
														position, tokenIndex = position755, tokenIndex755
														if buffer[position] != rune('E') {
															goto l712
														}
														position++
													}
												l755:
													{
														position757, tokenIndex757 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l758
														}
														position++
														goto l757
													l758:
														position, tokenIndex = position757, tokenIndex757
														if buffer[position] != rune('T') {
															goto l712
														}
														position++
													}
												l757:
													{
														position759, tokenIndex759 := position, tokenIndex
														if !_rules[rulews]() {
															goto l759
														}
														if !_rules[rulecc]() {
															goto l759
														}
														goto l760
													l759:
														position, tokenIndex = position759, tokenIndex759
													}
												l760:
													{
														add(ruleAction87, position)
													}
													add(ruleRet, position752)
												}
												break
											default:
												{
													position762 := position
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
															goto l712
														}
														position++
													}
												l763:
													{
														position765, tokenIndex765 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l766
														}
														position++
														goto l765
													l766:
														position, tokenIndex = position765, tokenIndex765
														if buffer[position] != rune('A') {
															goto l712
														}
														position++
													}
												l765:
													{
														position767, tokenIndex767 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l768
														}
														position++
														goto l767
													l768:
														position, tokenIndex = position767, tokenIndex767
														if buffer[position] != rune('L') {
															goto l712
														}
														position++
													}
												l767:
													{
														position769, tokenIndex769 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l770
														}
														position++
														goto l769
													l770:
														position, tokenIndex = position769, tokenIndex769
														if buffer[position] != rune('L') {
															goto l712
														}
														position++
													}
												l769:
													if !_rules[rulews]() {
														goto l712
													}
													{
														position771, tokenIndex771 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l771
														}
														if !_rules[rulesep]() {
															goto l771
														}
														goto l772
													l771:
														position, tokenIndex = position771, tokenIndex771
													}
												l772:
													if !_rules[ruleSrc16]() {
														goto l712
													}
													{
														add(ruleAction86, position)
													}
													add(ruleCall, position762)
												}
												break
											}
										}

									}
								l714:
									add(ruleJump, position713)
								}
								goto l13
							l712:
								position, tokenIndex = position13, tokenIndex13
								{
									position774 := position
									{
										position775, tokenIndex775 := position, tokenIndex
										{
											position777 := position
											{
												position778, tokenIndex778 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l779
												}
												position++
												goto l778
											l779:
												position, tokenIndex = position778, tokenIndex778
												if buffer[position] != rune('I') {
													goto l776
												}
												position++
											}
										l778:
											{
												position780, tokenIndex780 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l781
												}
												position++
												goto l780
											l781:
												position, tokenIndex = position780, tokenIndex780
												if buffer[position] != rune('N') {
													goto l776
												}
												position++
											}
										l780:
											if !_rules[rulews]() {
												goto l776
											}
											if !_rules[ruleReg8]() {
												goto l776
											}
											if !_rules[rulesep]() {
												goto l776
											}
											if !_rules[rulePort]() {
												goto l776
											}
											{
												add(ruleAction91, position)
											}
											add(ruleIN, position777)
										}
										goto l775
									l776:
										position, tokenIndex = position775, tokenIndex775
										{
											position783 := position
											{
												position784, tokenIndex784 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l785
												}
												position++
												goto l784
											l785:
												position, tokenIndex = position784, tokenIndex784
												if buffer[position] != rune('O') {
													goto l0
												}
												position++
											}
										l784:
											{
												position786, tokenIndex786 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l787
												}
												position++
												goto l786
											l787:
												position, tokenIndex = position786, tokenIndex786
												if buffer[position] != rune('U') {
													goto l0
												}
												position++
											}
										l786:
											{
												position788, tokenIndex788 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l789
												}
												position++
												goto l788
											l789:
												position, tokenIndex = position788, tokenIndex788
												if buffer[position] != rune('T') {
													goto l0
												}
												position++
											}
										l788:
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
												add(ruleAction92, position)
											}
											add(ruleOUT, position783)
										}
									}
								l775:
									add(ruleIO, position774)
								}
							}
						l13:
							add(ruleInstruction, position10)
						}
						{
							position791, tokenIndex791 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l791
							}
							position++
							goto l792
						l791:
							position, tokenIndex = position791, tokenIndex791
						}
					l792:
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
						position794, tokenIndex794 := position, tokenIndex
						{
							position796 := position
						l797:
							{
								position798, tokenIndex798 := position, tokenIndex
								if !_rules[rulews]() {
									goto l798
								}
								goto l797
							l798:
								position, tokenIndex = position798, tokenIndex798
							}
							if buffer[position] != rune('\n') {
								goto l795
							}
							position++
							add(ruleBlankLine, position796)
						}
						goto l794
					l795:
						position, tokenIndex = position794, tokenIndex794
						{
							position799 := position
							{
								position800 := position
							l801:
								{
									position802, tokenIndex802 := position, tokenIndex
									if !_rules[rulews]() {
										goto l802
									}
									goto l801
								l802:
									position, tokenIndex = position802, tokenIndex802
								}
								{
									position803, tokenIndex803 := position, tokenIndex
									{
										position805 := position
										{
											position806, tokenIndex806 := position, tokenIndex
											{
												position808 := position
												{
													position809, tokenIndex809 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l810
													}
													position++
													goto l809
												l810:
													position, tokenIndex = position809, tokenIndex809
													if buffer[position] != rune('P') {
														goto l807
													}
													position++
												}
											l809:
												{
													position811, tokenIndex811 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l812
													}
													position++
													goto l811
												l812:
													position, tokenIndex = position811, tokenIndex811
													if buffer[position] != rune('U') {
														goto l807
													}
													position++
												}
											l811:
												{
													position813, tokenIndex813 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l814
													}
													position++
													goto l813
												l814:
													position, tokenIndex = position813, tokenIndex813
													if buffer[position] != rune('S') {
														goto l807
													}
													position++
												}
											l813:
												{
													position815, tokenIndex815 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l816
													}
													position++
													goto l815
												l816:
													position, tokenIndex = position815, tokenIndex815
													if buffer[position] != rune('H') {
														goto l807
													}
													position++
												}
											l815:
												if !_rules[rulews]() {
													goto l807
												}
												if !_rules[ruleSrc16]() {
													goto l807
												}
												{
													add(ruleAction3, position)
												}
												add(rulePush, position808)
											}
											goto l806
										l807:
											position, tokenIndex = position806, tokenIndex806
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position819 := position
														{
															position820, tokenIndex820 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l821
															}
															position++
															goto l820
														l821:
															position, tokenIndex = position820, tokenIndex820
															if buffer[position] != rune('E') {
																goto l804
															}
															position++
														}
													l820:
														{
															position822, tokenIndex822 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l823
															}
															position++
															goto l822
														l823:
															position, tokenIndex = position822, tokenIndex822
															if buffer[position] != rune('X') {
																goto l804
															}
															position++
														}
													l822:
														if !_rules[rulews]() {
															goto l804
														}
														if !_rules[ruleDst16]() {
															goto l804
														}
														if !_rules[rulesep]() {
															goto l804
														}
														if !_rules[ruleSrc16]() {
															goto l804
														}
														{
															add(ruleAction5, position)
														}
														add(ruleEx, position819)
													}
													break
												case 'P', 'p':
													{
														position825 := position
														{
															position826, tokenIndex826 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l827
															}
															position++
															goto l826
														l827:
															position, tokenIndex = position826, tokenIndex826
															if buffer[position] != rune('P') {
																goto l804
															}
															position++
														}
													l826:
														{
															position828, tokenIndex828 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l829
															}
															position++
															goto l828
														l829:
															position, tokenIndex = position828, tokenIndex828
															if buffer[position] != rune('O') {
																goto l804
															}
															position++
														}
													l828:
														{
															position830, tokenIndex830 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l831
															}
															position++
															goto l830
														l831:
															position, tokenIndex = position830, tokenIndex830
															if buffer[position] != rune('P') {
																goto l804
															}
															position++
														}
													l830:
														if !_rules[rulews]() {
															goto l804
														}
														if !_rules[ruleDst16]() {
															goto l804
														}
														{
															add(ruleAction4, position)
														}
														add(rulePop, position825)
													}
													break
												default:
													{
														position833 := position
														{
															position834, tokenIndex834 := position, tokenIndex
															{
																position836 := position
																{
																	position837, tokenIndex837 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l838
																	}
																	position++
																	goto l837
																l838:
																	position, tokenIndex = position837, tokenIndex837
																	if buffer[position] != rune('L') {
																		goto l835
																	}
																	position++
																}
															l837:
																{
																	position839, tokenIndex839 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l840
																	}
																	position++
																	goto l839
																l840:
																	position, tokenIndex = position839, tokenIndex839
																	if buffer[position] != rune('D') {
																		goto l835
																	}
																	position++
																}
															l839:
																if !_rules[rulews]() {
																	goto l835
																}
																if !_rules[ruleDst16]() {
																	goto l835
																}
																if !_rules[rulesep]() {
																	goto l835
																}
																if !_rules[ruleSrc16]() {
																	goto l835
																}
																{
																	add(ruleAction2, position)
																}
																add(ruleLoad16, position836)
															}
															goto l834
														l835:
															position, tokenIndex = position834, tokenIndex834
															{
																position842 := position
																{
																	position843, tokenIndex843 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l844
																	}
																	position++
																	goto l843
																l844:
																	position, tokenIndex = position843, tokenIndex843
																	if buffer[position] != rune('L') {
																		goto l804
																	}
																	position++
																}
															l843:
																{
																	position845, tokenIndex845 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l846
																	}
																	position++
																	goto l845
																l846:
																	position, tokenIndex = position845, tokenIndex845
																	if buffer[position] != rune('D') {
																		goto l804
																	}
																	position++
																}
															l845:
																if !_rules[rulews]() {
																	goto l804
																}
																{
																	position847 := position
																	{
																		position848, tokenIndex848 := position, tokenIndex
																		if !_rules[ruleReg8]() {
																			goto l849
																		}
																		goto l848
																	l849:
																		position, tokenIndex = position848, tokenIndex848
																		if !_rules[ruleReg16Contents]() {
																			goto l850
																		}
																		goto l848
																	l850:
																		position, tokenIndex = position848, tokenIndex848
																		if !_rules[rulenn_contents]() {
																			goto l804
																		}
																	}
																l848:
																	{
																		add(ruleAction15, position)
																	}
																	add(ruleDst8, position847)
																}
																if !_rules[rulesep]() {
																	goto l804
																}
																if !_rules[ruleSrc8]() {
																	goto l804
																}
																{
																	add(ruleAction1, position)
																}
																add(ruleLoad8, position842)
															}
														}
													l834:
														add(ruleLoad, position833)
													}
													break
												}
											}

										}
									l806:
										add(ruleAssignment, position805)
									}
									goto l803
								l804:
									position, tokenIndex = position803, tokenIndex803
									{
										position854 := position
										{
											position855, tokenIndex855 := position, tokenIndex
											{
												position857 := position
												{
													position858, tokenIndex858 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l859
													}
													position++
													goto l858
												l859:
													position, tokenIndex = position858, tokenIndex858
													if buffer[position] != rune('I') {
														goto l856
													}
													position++
												}
											l858:
												{
													position860, tokenIndex860 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l861
													}
													position++
													goto l860
												l861:
													position, tokenIndex = position860, tokenIndex860
													if buffer[position] != rune('N') {
														goto l856
													}
													position++
												}
											l860:
												{
													position862, tokenIndex862 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l863
													}
													position++
													goto l862
												l863:
													position, tokenIndex = position862, tokenIndex862
													if buffer[position] != rune('C') {
														goto l856
													}
													position++
												}
											l862:
												if !_rules[rulews]() {
													goto l856
												}
												if !_rules[ruleILoc8]() {
													goto l856
												}
												{
													add(ruleAction6, position)
												}
												add(ruleInc16Indexed8, position857)
											}
											goto l855
										l856:
											position, tokenIndex = position855, tokenIndex855
											{
												position866 := position
												{
													position867, tokenIndex867 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l868
													}
													position++
													goto l867
												l868:
													position, tokenIndex = position867, tokenIndex867
													if buffer[position] != rune('I') {
														goto l865
													}
													position++
												}
											l867:
												{
													position869, tokenIndex869 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l870
													}
													position++
													goto l869
												l870:
													position, tokenIndex = position869, tokenIndex869
													if buffer[position] != rune('N') {
														goto l865
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
														goto l865
													}
													position++
												}
											l871:
												if !_rules[rulews]() {
													goto l865
												}
												if !_rules[ruleLoc16]() {
													goto l865
												}
												{
													add(ruleAction8, position)
												}
												add(ruleInc16, position866)
											}
											goto l855
										l865:
											position, tokenIndex = position855, tokenIndex855
											{
												position874 := position
												{
													position875, tokenIndex875 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l876
													}
													position++
													goto l875
												l876:
													position, tokenIndex = position875, tokenIndex875
													if buffer[position] != rune('I') {
														goto l853
													}
													position++
												}
											l875:
												{
													position877, tokenIndex877 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l878
													}
													position++
													goto l877
												l878:
													position, tokenIndex = position877, tokenIndex877
													if buffer[position] != rune('N') {
														goto l853
													}
													position++
												}
											l877:
												{
													position879, tokenIndex879 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l880
													}
													position++
													goto l879
												l880:
													position, tokenIndex = position879, tokenIndex879
													if buffer[position] != rune('C') {
														goto l853
													}
													position++
												}
											l879:
												if !_rules[rulews]() {
													goto l853
												}
												if !_rules[ruleLoc8]() {
													goto l853
												}
												{
													add(ruleAction7, position)
												}
												add(ruleInc8, position874)
											}
										}
									l855:
										add(ruleInc, position854)
									}
									goto l803
								l853:
									position, tokenIndex = position803, tokenIndex803
									{
										position883 := position
										{
											position884, tokenIndex884 := position, tokenIndex
											{
												position886 := position
												{
													position887, tokenIndex887 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l888
													}
													position++
													goto l887
												l888:
													position, tokenIndex = position887, tokenIndex887
													if buffer[position] != rune('D') {
														goto l885
													}
													position++
												}
											l887:
												{
													position889, tokenIndex889 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l890
													}
													position++
													goto l889
												l890:
													position, tokenIndex = position889, tokenIndex889
													if buffer[position] != rune('E') {
														goto l885
													}
													position++
												}
											l889:
												{
													position891, tokenIndex891 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l892
													}
													position++
													goto l891
												l892:
													position, tokenIndex = position891, tokenIndex891
													if buffer[position] != rune('C') {
														goto l885
													}
													position++
												}
											l891:
												if !_rules[rulews]() {
													goto l885
												}
												if !_rules[ruleILoc8]() {
													goto l885
												}
												{
													add(ruleAction9, position)
												}
												add(ruleDec16Indexed8, position886)
											}
											goto l884
										l885:
											position, tokenIndex = position884, tokenIndex884
											{
												position895 := position
												{
													position896, tokenIndex896 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l897
													}
													position++
													goto l896
												l897:
													position, tokenIndex = position896, tokenIndex896
													if buffer[position] != rune('D') {
														goto l894
													}
													position++
												}
											l896:
												{
													position898, tokenIndex898 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l899
													}
													position++
													goto l898
												l899:
													position, tokenIndex = position898, tokenIndex898
													if buffer[position] != rune('E') {
														goto l894
													}
													position++
												}
											l898:
												{
													position900, tokenIndex900 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l901
													}
													position++
													goto l900
												l901:
													position, tokenIndex = position900, tokenIndex900
													if buffer[position] != rune('C') {
														goto l894
													}
													position++
												}
											l900:
												if !_rules[rulews]() {
													goto l894
												}
												if !_rules[ruleLoc16]() {
													goto l894
												}
												{
													add(ruleAction11, position)
												}
												add(ruleDec16, position895)
											}
											goto l884
										l894:
											position, tokenIndex = position884, tokenIndex884
											{
												position903 := position
												{
													position904, tokenIndex904 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l905
													}
													position++
													goto l904
												l905:
													position, tokenIndex = position904, tokenIndex904
													if buffer[position] != rune('D') {
														goto l882
													}
													position++
												}
											l904:
												{
													position906, tokenIndex906 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l907
													}
													position++
													goto l906
												l907:
													position, tokenIndex = position906, tokenIndex906
													if buffer[position] != rune('E') {
														goto l882
													}
													position++
												}
											l906:
												{
													position908, tokenIndex908 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l909
													}
													position++
													goto l908
												l909:
													position, tokenIndex = position908, tokenIndex908
													if buffer[position] != rune('C') {
														goto l882
													}
													position++
												}
											l908:
												if !_rules[rulews]() {
													goto l882
												}
												if !_rules[ruleLoc8]() {
													goto l882
												}
												{
													add(ruleAction10, position)
												}
												add(ruleDec8, position903)
											}
										}
									l884:
										add(ruleDec, position883)
									}
									goto l803
								l882:
									position, tokenIndex = position803, tokenIndex803
									{
										position912 := position
										{
											position913, tokenIndex913 := position, tokenIndex
											{
												position915 := position
												{
													position916, tokenIndex916 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l917
													}
													position++
													goto l916
												l917:
													position, tokenIndex = position916, tokenIndex916
													if buffer[position] != rune('A') {
														goto l914
													}
													position++
												}
											l916:
												{
													position918, tokenIndex918 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l919
													}
													position++
													goto l918
												l919:
													position, tokenIndex = position918, tokenIndex918
													if buffer[position] != rune('D') {
														goto l914
													}
													position++
												}
											l918:
												{
													position920, tokenIndex920 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l921
													}
													position++
													goto l920
												l921:
													position, tokenIndex = position920, tokenIndex920
													if buffer[position] != rune('D') {
														goto l914
													}
													position++
												}
											l920:
												if !_rules[rulews]() {
													goto l914
												}
												if !_rules[ruleDst16]() {
													goto l914
												}
												if !_rules[rulesep]() {
													goto l914
												}
												if !_rules[ruleSrc16]() {
													goto l914
												}
												{
													add(ruleAction12, position)
												}
												add(ruleAdd16, position915)
											}
											goto l913
										l914:
											position, tokenIndex = position913, tokenIndex913
											{
												position924 := position
												{
													position925, tokenIndex925 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l926
													}
													position++
													goto l925
												l926:
													position, tokenIndex = position925, tokenIndex925
													if buffer[position] != rune('A') {
														goto l923
													}
													position++
												}
											l925:
												{
													position927, tokenIndex927 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l928
													}
													position++
													goto l927
												l928:
													position, tokenIndex = position927, tokenIndex927
													if buffer[position] != rune('D') {
														goto l923
													}
													position++
												}
											l927:
												{
													position929, tokenIndex929 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l930
													}
													position++
													goto l929
												l930:
													position, tokenIndex = position929, tokenIndex929
													if buffer[position] != rune('C') {
														goto l923
													}
													position++
												}
											l929:
												if !_rules[rulews]() {
													goto l923
												}
												if !_rules[ruleDst16]() {
													goto l923
												}
												if !_rules[rulesep]() {
													goto l923
												}
												if !_rules[ruleSrc16]() {
													goto l923
												}
												{
													add(ruleAction13, position)
												}
												add(ruleAdc16, position924)
											}
											goto l913
										l923:
											position, tokenIndex = position913, tokenIndex913
											{
												position932 := position
												{
													position933, tokenIndex933 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l934
													}
													position++
													goto l933
												l934:
													position, tokenIndex = position933, tokenIndex933
													if buffer[position] != rune('S') {
														goto l911
													}
													position++
												}
											l933:
												{
													position935, tokenIndex935 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l936
													}
													position++
													goto l935
												l936:
													position, tokenIndex = position935, tokenIndex935
													if buffer[position] != rune('B') {
														goto l911
													}
													position++
												}
											l935:
												{
													position937, tokenIndex937 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l938
													}
													position++
													goto l937
												l938:
													position, tokenIndex = position937, tokenIndex937
													if buffer[position] != rune('C') {
														goto l911
													}
													position++
												}
											l937:
												if !_rules[rulews]() {
													goto l911
												}
												if !_rules[ruleDst16]() {
													goto l911
												}
												if !_rules[rulesep]() {
													goto l911
												}
												if !_rules[ruleSrc16]() {
													goto l911
												}
												{
													add(ruleAction14, position)
												}
												add(ruleSbc16, position932)
											}
										}
									l913:
										add(ruleAlu16, position912)
									}
									goto l803
								l911:
									position, tokenIndex = position803, tokenIndex803
									{
										position941 := position
										{
											position942, tokenIndex942 := position, tokenIndex
											{
												position944 := position
												{
													position945, tokenIndex945 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l946
													}
													position++
													goto l945
												l946:
													position, tokenIndex = position945, tokenIndex945
													if buffer[position] != rune('A') {
														goto l943
													}
													position++
												}
											l945:
												{
													position947, tokenIndex947 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l948
													}
													position++
													goto l947
												l948:
													position, tokenIndex = position947, tokenIndex947
													if buffer[position] != rune('D') {
														goto l943
													}
													position++
												}
											l947:
												{
													position949, tokenIndex949 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l950
													}
													position++
													goto l949
												l950:
													position, tokenIndex = position949, tokenIndex949
													if buffer[position] != rune('D') {
														goto l943
													}
													position++
												}
											l949:
												if !_rules[rulews]() {
													goto l943
												}
												{
													position951, tokenIndex951 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l952
													}
													position++
													goto l951
												l952:
													position, tokenIndex = position951, tokenIndex951
													if buffer[position] != rune('A') {
														goto l943
													}
													position++
												}
											l951:
												if !_rules[rulesep]() {
													goto l943
												}
												if !_rules[ruleSrc8]() {
													goto l943
												}
												{
													add(ruleAction29, position)
												}
												add(ruleAdd, position944)
											}
											goto l942
										l943:
											position, tokenIndex = position942, tokenIndex942
											{
												position955 := position
												{
													position956, tokenIndex956 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l957
													}
													position++
													goto l956
												l957:
													position, tokenIndex = position956, tokenIndex956
													if buffer[position] != rune('A') {
														goto l954
													}
													position++
												}
											l956:
												{
													position958, tokenIndex958 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l959
													}
													position++
													goto l958
												l959:
													position, tokenIndex = position958, tokenIndex958
													if buffer[position] != rune('D') {
														goto l954
													}
													position++
												}
											l958:
												{
													position960, tokenIndex960 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l961
													}
													position++
													goto l960
												l961:
													position, tokenIndex = position960, tokenIndex960
													if buffer[position] != rune('C') {
														goto l954
													}
													position++
												}
											l960:
												if !_rules[rulews]() {
													goto l954
												}
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
														goto l954
													}
													position++
												}
											l962:
												if !_rules[rulesep]() {
													goto l954
												}
												if !_rules[ruleSrc8]() {
													goto l954
												}
												{
													add(ruleAction30, position)
												}
												add(ruleAdc, position955)
											}
											goto l942
										l954:
											position, tokenIndex = position942, tokenIndex942
											{
												position966 := position
												{
													position967, tokenIndex967 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l968
													}
													position++
													goto l967
												l968:
													position, tokenIndex = position967, tokenIndex967
													if buffer[position] != rune('S') {
														goto l965
													}
													position++
												}
											l967:
												{
													position969, tokenIndex969 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l970
													}
													position++
													goto l969
												l970:
													position, tokenIndex = position969, tokenIndex969
													if buffer[position] != rune('U') {
														goto l965
													}
													position++
												}
											l969:
												{
													position971, tokenIndex971 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l972
													}
													position++
													goto l971
												l972:
													position, tokenIndex = position971, tokenIndex971
													if buffer[position] != rune('B') {
														goto l965
													}
													position++
												}
											l971:
												if !_rules[rulews]() {
													goto l965
												}
												if !_rules[ruleSrc8]() {
													goto l965
												}
												{
													add(ruleAction31, position)
												}
												add(ruleSub, position966)
											}
											goto l942
										l965:
											position, tokenIndex = position942, tokenIndex942
											{
												switch buffer[position] {
												case 'C', 'c':
													{
														position975 := position
														{
															position976, tokenIndex976 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l977
															}
															position++
															goto l976
														l977:
															position, tokenIndex = position976, tokenIndex976
															if buffer[position] != rune('C') {
																goto l940
															}
															position++
														}
													l976:
														{
															position978, tokenIndex978 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l979
															}
															position++
															goto l978
														l979:
															position, tokenIndex = position978, tokenIndex978
															if buffer[position] != rune('P') {
																goto l940
															}
															position++
														}
													l978:
														if !_rules[rulews]() {
															goto l940
														}
														if !_rules[ruleSrc8]() {
															goto l940
														}
														{
															add(ruleAction36, position)
														}
														add(ruleCp, position975)
													}
													break
												case 'O', 'o':
													{
														position981 := position
														{
															position982, tokenIndex982 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l983
															}
															position++
															goto l982
														l983:
															position, tokenIndex = position982, tokenIndex982
															if buffer[position] != rune('O') {
																goto l940
															}
															position++
														}
													l982:
														{
															position984, tokenIndex984 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l985
															}
															position++
															goto l984
														l985:
															position, tokenIndex = position984, tokenIndex984
															if buffer[position] != rune('R') {
																goto l940
															}
															position++
														}
													l984:
														if !_rules[rulews]() {
															goto l940
														}
														if !_rules[ruleSrc8]() {
															goto l940
														}
														{
															add(ruleAction35, position)
														}
														add(ruleOr, position981)
													}
													break
												case 'X', 'x':
													{
														position987 := position
														{
															position988, tokenIndex988 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l989
															}
															position++
															goto l988
														l989:
															position, tokenIndex = position988, tokenIndex988
															if buffer[position] != rune('X') {
																goto l940
															}
															position++
														}
													l988:
														{
															position990, tokenIndex990 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l991
															}
															position++
															goto l990
														l991:
															position, tokenIndex = position990, tokenIndex990
															if buffer[position] != rune('O') {
																goto l940
															}
															position++
														}
													l990:
														{
															position992, tokenIndex992 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l993
															}
															position++
															goto l992
														l993:
															position, tokenIndex = position992, tokenIndex992
															if buffer[position] != rune('R') {
																goto l940
															}
															position++
														}
													l992:
														if !_rules[rulews]() {
															goto l940
														}
														if !_rules[ruleSrc8]() {
															goto l940
														}
														{
															add(ruleAction34, position)
														}
														add(ruleXor, position987)
													}
													break
												case 'A', 'a':
													{
														position995 := position
														{
															position996, tokenIndex996 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l997
															}
															position++
															goto l996
														l997:
															position, tokenIndex = position996, tokenIndex996
															if buffer[position] != rune('A') {
																goto l940
															}
															position++
														}
													l996:
														{
															position998, tokenIndex998 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l999
															}
															position++
															goto l998
														l999:
															position, tokenIndex = position998, tokenIndex998
															if buffer[position] != rune('N') {
																goto l940
															}
															position++
														}
													l998:
														{
															position1000, tokenIndex1000 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1001
															}
															position++
															goto l1000
														l1001:
															position, tokenIndex = position1000, tokenIndex1000
															if buffer[position] != rune('D') {
																goto l940
															}
															position++
														}
													l1000:
														if !_rules[rulews]() {
															goto l940
														}
														if !_rules[ruleSrc8]() {
															goto l940
														}
														{
															add(ruleAction33, position)
														}
														add(ruleAnd, position995)
													}
													break
												default:
													{
														position1003 := position
														{
															position1004, tokenIndex1004 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1005
															}
															position++
															goto l1004
														l1005:
															position, tokenIndex = position1004, tokenIndex1004
															if buffer[position] != rune('S') {
																goto l940
															}
															position++
														}
													l1004:
														{
															position1006, tokenIndex1006 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1007
															}
															position++
															goto l1006
														l1007:
															position, tokenIndex = position1006, tokenIndex1006
															if buffer[position] != rune('B') {
																goto l940
															}
															position++
														}
													l1006:
														{
															position1008, tokenIndex1008 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1009
															}
															position++
															goto l1008
														l1009:
															position, tokenIndex = position1008, tokenIndex1008
															if buffer[position] != rune('C') {
																goto l940
															}
															position++
														}
													l1008:
														if !_rules[rulews]() {
															goto l940
														}
														{
															position1010, tokenIndex1010 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1011
															}
															position++
															goto l1010
														l1011:
															position, tokenIndex = position1010, tokenIndex1010
															if buffer[position] != rune('A') {
																goto l940
															}
															position++
														}
													l1010:
														if !_rules[rulesep]() {
															goto l940
														}
														if !_rules[ruleSrc8]() {
															goto l940
														}
														{
															add(ruleAction32, position)
														}
														add(ruleSbc, position1003)
													}
													break
												}
											}

										}
									l942:
										add(ruleAlu, position941)
									}
									goto l803
								l940:
									position, tokenIndex = position803, tokenIndex803
									{
										position1014 := position
										{
											position1015, tokenIndex1015 := position, tokenIndex
											{
												position1017 := position
												{
													position1018, tokenIndex1018 := position, tokenIndex
													{
														position1020 := position
														{
															position1021, tokenIndex1021 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1022
															}
															position++
															goto l1021
														l1022:
															position, tokenIndex = position1021, tokenIndex1021
															if buffer[position] != rune('R') {
																goto l1019
															}
															position++
														}
													l1021:
														{
															position1023, tokenIndex1023 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1024
															}
															position++
															goto l1023
														l1024:
															position, tokenIndex = position1023, tokenIndex1023
															if buffer[position] != rune('L') {
																goto l1019
															}
															position++
														}
													l1023:
														{
															position1025, tokenIndex1025 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1026
															}
															position++
															goto l1025
														l1026:
															position, tokenIndex = position1025, tokenIndex1025
															if buffer[position] != rune('C') {
																goto l1019
															}
															position++
														}
													l1025:
														if !_rules[rulews]() {
															goto l1019
														}
														if !_rules[ruleLoc8]() {
															goto l1019
														}
														{
															add(ruleAction37, position)
														}
														add(ruleRlc, position1020)
													}
													goto l1018
												l1019:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1029 := position
														{
															position1030, tokenIndex1030 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1031
															}
															position++
															goto l1030
														l1031:
															position, tokenIndex = position1030, tokenIndex1030
															if buffer[position] != rune('R') {
																goto l1028
															}
															position++
														}
													l1030:
														{
															position1032, tokenIndex1032 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1033
															}
															position++
															goto l1032
														l1033:
															position, tokenIndex = position1032, tokenIndex1032
															if buffer[position] != rune('R') {
																goto l1028
															}
															position++
														}
													l1032:
														{
															position1034, tokenIndex1034 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1035
															}
															position++
															goto l1034
														l1035:
															position, tokenIndex = position1034, tokenIndex1034
															if buffer[position] != rune('C') {
																goto l1028
															}
															position++
														}
													l1034:
														if !_rules[rulews]() {
															goto l1028
														}
														if !_rules[ruleLoc8]() {
															goto l1028
														}
														{
															add(ruleAction38, position)
														}
														add(ruleRrc, position1029)
													}
													goto l1018
												l1028:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1038 := position
														{
															position1039, tokenIndex1039 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1040
															}
															position++
															goto l1039
														l1040:
															position, tokenIndex = position1039, tokenIndex1039
															if buffer[position] != rune('R') {
																goto l1037
															}
															position++
														}
													l1039:
														{
															position1041, tokenIndex1041 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1042
															}
															position++
															goto l1041
														l1042:
															position, tokenIndex = position1041, tokenIndex1041
															if buffer[position] != rune('L') {
																goto l1037
															}
															position++
														}
													l1041:
														if !_rules[rulews]() {
															goto l1037
														}
														if !_rules[ruleLoc8]() {
															goto l1037
														}
														{
															add(ruleAction39, position)
														}
														add(ruleRl, position1038)
													}
													goto l1018
												l1037:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1045 := position
														{
															position1046, tokenIndex1046 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1047
															}
															position++
															goto l1046
														l1047:
															position, tokenIndex = position1046, tokenIndex1046
															if buffer[position] != rune('R') {
																goto l1044
															}
															position++
														}
													l1046:
														{
															position1048, tokenIndex1048 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1049
															}
															position++
															goto l1048
														l1049:
															position, tokenIndex = position1048, tokenIndex1048
															if buffer[position] != rune('R') {
																goto l1044
															}
															position++
														}
													l1048:
														if !_rules[rulews]() {
															goto l1044
														}
														if !_rules[ruleLoc8]() {
															goto l1044
														}
														{
															add(ruleAction40, position)
														}
														add(ruleRr, position1045)
													}
													goto l1018
												l1044:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1052 := position
														{
															position1053, tokenIndex1053 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1054
															}
															position++
															goto l1053
														l1054:
															position, tokenIndex = position1053, tokenIndex1053
															if buffer[position] != rune('S') {
																goto l1051
															}
															position++
														}
													l1053:
														{
															position1055, tokenIndex1055 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1056
															}
															position++
															goto l1055
														l1056:
															position, tokenIndex = position1055, tokenIndex1055
															if buffer[position] != rune('L') {
																goto l1051
															}
															position++
														}
													l1055:
														{
															position1057, tokenIndex1057 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1058
															}
															position++
															goto l1057
														l1058:
															position, tokenIndex = position1057, tokenIndex1057
															if buffer[position] != rune('A') {
																goto l1051
															}
															position++
														}
													l1057:
														if !_rules[rulews]() {
															goto l1051
														}
														if !_rules[ruleLoc8]() {
															goto l1051
														}
														{
															add(ruleAction41, position)
														}
														add(ruleSla, position1052)
													}
													goto l1018
												l1051:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1061 := position
														{
															position1062, tokenIndex1062 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1063
															}
															position++
															goto l1062
														l1063:
															position, tokenIndex = position1062, tokenIndex1062
															if buffer[position] != rune('S') {
																goto l1060
															}
															position++
														}
													l1062:
														{
															position1064, tokenIndex1064 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1065
															}
															position++
															goto l1064
														l1065:
															position, tokenIndex = position1064, tokenIndex1064
															if buffer[position] != rune('R') {
																goto l1060
															}
															position++
														}
													l1064:
														{
															position1066, tokenIndex1066 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1067
															}
															position++
															goto l1066
														l1067:
															position, tokenIndex = position1066, tokenIndex1066
															if buffer[position] != rune('A') {
																goto l1060
															}
															position++
														}
													l1066:
														if !_rules[rulews]() {
															goto l1060
														}
														if !_rules[ruleLoc8]() {
															goto l1060
														}
														{
															add(ruleAction42, position)
														}
														add(ruleSra, position1061)
													}
													goto l1018
												l1060:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1070 := position
														{
															position1071, tokenIndex1071 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1072
															}
															position++
															goto l1071
														l1072:
															position, tokenIndex = position1071, tokenIndex1071
															if buffer[position] != rune('S') {
																goto l1069
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
																goto l1069
															}
															position++
														}
													l1073:
														{
															position1075, tokenIndex1075 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1076
															}
															position++
															goto l1075
														l1076:
															position, tokenIndex = position1075, tokenIndex1075
															if buffer[position] != rune('L') {
																goto l1069
															}
															position++
														}
													l1075:
														if !_rules[rulews]() {
															goto l1069
														}
														if !_rules[ruleLoc8]() {
															goto l1069
														}
														{
															add(ruleAction43, position)
														}
														add(ruleSll, position1070)
													}
													goto l1018
												l1069:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1078 := position
														{
															position1079, tokenIndex1079 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1080
															}
															position++
															goto l1079
														l1080:
															position, tokenIndex = position1079, tokenIndex1079
															if buffer[position] != rune('S') {
																goto l1016
															}
															position++
														}
													l1079:
														{
															position1081, tokenIndex1081 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1082
															}
															position++
															goto l1081
														l1082:
															position, tokenIndex = position1081, tokenIndex1081
															if buffer[position] != rune('R') {
																goto l1016
															}
															position++
														}
													l1081:
														{
															position1083, tokenIndex1083 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1084
															}
															position++
															goto l1083
														l1084:
															position, tokenIndex = position1083, tokenIndex1083
															if buffer[position] != rune('L') {
																goto l1016
															}
															position++
														}
													l1083:
														if !_rules[rulews]() {
															goto l1016
														}
														if !_rules[ruleLoc8]() {
															goto l1016
														}
														{
															add(ruleAction44, position)
														}
														add(ruleSrl, position1078)
													}
												}
											l1018:
												add(ruleRot, position1017)
											}
											goto l1015
										l1016:
											position, tokenIndex = position1015, tokenIndex1015
											{
												switch buffer[position] {
												case 'S', 's':
													{
														position1087 := position
														{
															position1088, tokenIndex1088 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1089
															}
															position++
															goto l1088
														l1089:
															position, tokenIndex = position1088, tokenIndex1088
															if buffer[position] != rune('S') {
																goto l1013
															}
															position++
														}
													l1088:
														{
															position1090, tokenIndex1090 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1091
															}
															position++
															goto l1090
														l1091:
															position, tokenIndex = position1090, tokenIndex1090
															if buffer[position] != rune('E') {
																goto l1013
															}
															position++
														}
													l1090:
														{
															position1092, tokenIndex1092 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1093
															}
															position++
															goto l1092
														l1093:
															position, tokenIndex = position1092, tokenIndex1092
															if buffer[position] != rune('T') {
																goto l1013
															}
															position++
														}
													l1092:
														if !_rules[rulews]() {
															goto l1013
														}
														if !_rules[ruleoctaldigit]() {
															goto l1013
														}
														if !_rules[rulesep]() {
															goto l1013
														}
														if !_rules[ruleLoc8]() {
															goto l1013
														}
														{
															add(ruleAction47, position)
														}
														add(ruleSet, position1087)
													}
													break
												case 'R', 'r':
													{
														position1095 := position
														{
															position1096, tokenIndex1096 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1097
															}
															position++
															goto l1096
														l1097:
															position, tokenIndex = position1096, tokenIndex1096
															if buffer[position] != rune('R') {
																goto l1013
															}
															position++
														}
													l1096:
														{
															position1098, tokenIndex1098 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1099
															}
															position++
															goto l1098
														l1099:
															position, tokenIndex = position1098, tokenIndex1098
															if buffer[position] != rune('E') {
																goto l1013
															}
															position++
														}
													l1098:
														{
															position1100, tokenIndex1100 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1101
															}
															position++
															goto l1100
														l1101:
															position, tokenIndex = position1100, tokenIndex1100
															if buffer[position] != rune('S') {
																goto l1013
															}
															position++
														}
													l1100:
														if !_rules[rulews]() {
															goto l1013
														}
														if !_rules[ruleoctaldigit]() {
															goto l1013
														}
														if !_rules[rulesep]() {
															goto l1013
														}
														if !_rules[ruleLoc8]() {
															goto l1013
														}
														{
															add(ruleAction46, position)
														}
														add(ruleRes, position1095)
													}
													break
												default:
													{
														position1103 := position
														{
															position1104, tokenIndex1104 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1105
															}
															position++
															goto l1104
														l1105:
															position, tokenIndex = position1104, tokenIndex1104
															if buffer[position] != rune('B') {
																goto l1013
															}
															position++
														}
													l1104:
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
																goto l1013
															}
															position++
														}
													l1106:
														{
															position1108, tokenIndex1108 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1109
															}
															position++
															goto l1108
														l1109:
															position, tokenIndex = position1108, tokenIndex1108
															if buffer[position] != rune('T') {
																goto l1013
															}
															position++
														}
													l1108:
														if !_rules[rulews]() {
															goto l1013
														}
														if !_rules[ruleoctaldigit]() {
															goto l1013
														}
														if !_rules[rulesep]() {
															goto l1013
														}
														if !_rules[ruleLoc8]() {
															goto l1013
														}
														{
															add(ruleAction45, position)
														}
														add(ruleBit, position1103)
													}
													break
												}
											}

										}
									l1015:
										add(ruleBitOp, position1014)
									}
									goto l803
								l1013:
									position, tokenIndex = position803, tokenIndex803
									{
										position1112 := position
										{
											position1113, tokenIndex1113 := position, tokenIndex
											{
												position1115 := position
												{
													position1116 := position
													{
														position1117, tokenIndex1117 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1118
														}
														position++
														goto l1117
													l1118:
														position, tokenIndex = position1117, tokenIndex1117
														if buffer[position] != rune('R') {
															goto l1114
														}
														position++
													}
												l1117:
													{
														position1119, tokenIndex1119 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1120
														}
														position++
														goto l1119
													l1120:
														position, tokenIndex = position1119, tokenIndex1119
														if buffer[position] != rune('E') {
															goto l1114
														}
														position++
													}
												l1119:
													{
														position1121, tokenIndex1121 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1122
														}
														position++
														goto l1121
													l1122:
														position, tokenIndex = position1121, tokenIndex1121
														if buffer[position] != rune('T') {
															goto l1114
														}
														position++
													}
												l1121:
													{
														position1123, tokenIndex1123 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1124
														}
														position++
														goto l1123
													l1124:
														position, tokenIndex = position1123, tokenIndex1123
														if buffer[position] != rune('N') {
															goto l1114
														}
														position++
													}
												l1123:
													add(rulePegText, position1116)
												}
												{
													add(ruleAction62, position)
												}
												add(ruleRetn, position1115)
											}
											goto l1113
										l1114:
											position, tokenIndex = position1113, tokenIndex1113
											{
												position1127 := position
												{
													position1128 := position
													{
														position1129, tokenIndex1129 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1130
														}
														position++
														goto l1129
													l1130:
														position, tokenIndex = position1129, tokenIndex1129
														if buffer[position] != rune('R') {
															goto l1126
														}
														position++
													}
												l1129:
													{
														position1131, tokenIndex1131 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1132
														}
														position++
														goto l1131
													l1132:
														position, tokenIndex = position1131, tokenIndex1131
														if buffer[position] != rune('E') {
															goto l1126
														}
														position++
													}
												l1131:
													{
														position1133, tokenIndex1133 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1134
														}
														position++
														goto l1133
													l1134:
														position, tokenIndex = position1133, tokenIndex1133
														if buffer[position] != rune('T') {
															goto l1126
														}
														position++
													}
												l1133:
													{
														position1135, tokenIndex1135 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1136
														}
														position++
														goto l1135
													l1136:
														position, tokenIndex = position1135, tokenIndex1135
														if buffer[position] != rune('I') {
															goto l1126
														}
														position++
													}
												l1135:
													add(rulePegText, position1128)
												}
												{
													add(ruleAction63, position)
												}
												add(ruleReti, position1127)
											}
											goto l1113
										l1126:
											position, tokenIndex = position1113, tokenIndex1113
											{
												position1139 := position
												{
													position1140 := position
													{
														position1141, tokenIndex1141 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1142
														}
														position++
														goto l1141
													l1142:
														position, tokenIndex = position1141, tokenIndex1141
														if buffer[position] != rune('R') {
															goto l1138
														}
														position++
													}
												l1141:
													{
														position1143, tokenIndex1143 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1144
														}
														position++
														goto l1143
													l1144:
														position, tokenIndex = position1143, tokenIndex1143
														if buffer[position] != rune('R') {
															goto l1138
														}
														position++
													}
												l1143:
													{
														position1145, tokenIndex1145 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1146
														}
														position++
														goto l1145
													l1146:
														position, tokenIndex = position1145, tokenIndex1145
														if buffer[position] != rune('D') {
															goto l1138
														}
														position++
													}
												l1145:
													add(rulePegText, position1140)
												}
												{
													add(ruleAction64, position)
												}
												add(ruleRrd, position1139)
											}
											goto l1113
										l1138:
											position, tokenIndex = position1113, tokenIndex1113
											{
												position1149 := position
												{
													position1150 := position
													{
														position1151, tokenIndex1151 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1152
														}
														position++
														goto l1151
													l1152:
														position, tokenIndex = position1151, tokenIndex1151
														if buffer[position] != rune('I') {
															goto l1148
														}
														position++
													}
												l1151:
													{
														position1153, tokenIndex1153 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1154
														}
														position++
														goto l1153
													l1154:
														position, tokenIndex = position1153, tokenIndex1153
														if buffer[position] != rune('M') {
															goto l1148
														}
														position++
													}
												l1153:
													if buffer[position] != rune(' ') {
														goto l1148
													}
													position++
													if buffer[position] != rune('0') {
														goto l1148
													}
													position++
													add(rulePegText, position1150)
												}
												{
													add(ruleAction66, position)
												}
												add(ruleIm0, position1149)
											}
											goto l1113
										l1148:
											position, tokenIndex = position1113, tokenIndex1113
											{
												position1157 := position
												{
													position1158 := position
													{
														position1159, tokenIndex1159 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1160
														}
														position++
														goto l1159
													l1160:
														position, tokenIndex = position1159, tokenIndex1159
														if buffer[position] != rune('I') {
															goto l1156
														}
														position++
													}
												l1159:
													{
														position1161, tokenIndex1161 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1162
														}
														position++
														goto l1161
													l1162:
														position, tokenIndex = position1161, tokenIndex1161
														if buffer[position] != rune('M') {
															goto l1156
														}
														position++
													}
												l1161:
													if buffer[position] != rune(' ') {
														goto l1156
													}
													position++
													if buffer[position] != rune('1') {
														goto l1156
													}
													position++
													add(rulePegText, position1158)
												}
												{
													add(ruleAction67, position)
												}
												add(ruleIm1, position1157)
											}
											goto l1113
										l1156:
											position, tokenIndex = position1113, tokenIndex1113
											{
												position1165 := position
												{
													position1166 := position
													{
														position1167, tokenIndex1167 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1168
														}
														position++
														goto l1167
													l1168:
														position, tokenIndex = position1167, tokenIndex1167
														if buffer[position] != rune('I') {
															goto l1164
														}
														position++
													}
												l1167:
													{
														position1169, tokenIndex1169 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1170
														}
														position++
														goto l1169
													l1170:
														position, tokenIndex = position1169, tokenIndex1169
														if buffer[position] != rune('M') {
															goto l1164
														}
														position++
													}
												l1169:
													if buffer[position] != rune(' ') {
														goto l1164
													}
													position++
													if buffer[position] != rune('2') {
														goto l1164
													}
													position++
													add(rulePegText, position1166)
												}
												{
													add(ruleAction68, position)
												}
												add(ruleIm2, position1165)
											}
											goto l1113
										l1164:
											position, tokenIndex = position1113, tokenIndex1113
											{
												switch buffer[position] {
												case 'I', 'O', 'i', 'o':
													{
														position1173 := position
														{
															position1174, tokenIndex1174 := position, tokenIndex
															{
																position1176 := position
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
																			goto l1175
																		}
																		position++
																	}
																l1178:
																	{
																		position1180, tokenIndex1180 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1181
																		}
																		position++
																		goto l1180
																	l1181:
																		position, tokenIndex = position1180, tokenIndex1180
																		if buffer[position] != rune('N') {
																			goto l1175
																		}
																		position++
																	}
																l1180:
																	{
																		position1182, tokenIndex1182 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1183
																		}
																		position++
																		goto l1182
																	l1183:
																		position, tokenIndex = position1182, tokenIndex1182
																		if buffer[position] != rune('I') {
																			goto l1175
																		}
																		position++
																	}
																l1182:
																	{
																		position1184, tokenIndex1184 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1185
																		}
																		position++
																		goto l1184
																	l1185:
																		position, tokenIndex = position1184, tokenIndex1184
																		if buffer[position] != rune('R') {
																			goto l1175
																		}
																		position++
																	}
																l1184:
																	add(rulePegText, position1177)
																}
																{
																	add(ruleAction79, position)
																}
																add(ruleInir, position1176)
															}
															goto l1174
														l1175:
															position, tokenIndex = position1174, tokenIndex1174
															{
																position1188 := position
																{
																	position1189 := position
																	{
																		position1190, tokenIndex1190 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1191
																		}
																		position++
																		goto l1190
																	l1191:
																		position, tokenIndex = position1190, tokenIndex1190
																		if buffer[position] != rune('I') {
																			goto l1187
																		}
																		position++
																	}
																l1190:
																	{
																		position1192, tokenIndex1192 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1193
																		}
																		position++
																		goto l1192
																	l1193:
																		position, tokenIndex = position1192, tokenIndex1192
																		if buffer[position] != rune('N') {
																			goto l1187
																		}
																		position++
																	}
																l1192:
																	{
																		position1194, tokenIndex1194 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1195
																		}
																		position++
																		goto l1194
																	l1195:
																		position, tokenIndex = position1194, tokenIndex1194
																		if buffer[position] != rune('I') {
																			goto l1187
																		}
																		position++
																	}
																l1194:
																	add(rulePegText, position1189)
																}
																{
																	add(ruleAction71, position)
																}
																add(ruleIni, position1188)
															}
															goto l1174
														l1187:
															position, tokenIndex = position1174, tokenIndex1174
															{
																position1198 := position
																{
																	position1199 := position
																	{
																		position1200, tokenIndex1200 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1201
																		}
																		position++
																		goto l1200
																	l1201:
																		position, tokenIndex = position1200, tokenIndex1200
																		if buffer[position] != rune('O') {
																			goto l1197
																		}
																		position++
																	}
																l1200:
																	{
																		position1202, tokenIndex1202 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1203
																		}
																		position++
																		goto l1202
																	l1203:
																		position, tokenIndex = position1202, tokenIndex1202
																		if buffer[position] != rune('T') {
																			goto l1197
																		}
																		position++
																	}
																l1202:
																	{
																		position1204, tokenIndex1204 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1205
																		}
																		position++
																		goto l1204
																	l1205:
																		position, tokenIndex = position1204, tokenIndex1204
																		if buffer[position] != rune('I') {
																			goto l1197
																		}
																		position++
																	}
																l1204:
																	{
																		position1206, tokenIndex1206 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1207
																		}
																		position++
																		goto l1206
																	l1207:
																		position, tokenIndex = position1206, tokenIndex1206
																		if buffer[position] != rune('R') {
																			goto l1197
																		}
																		position++
																	}
																l1206:
																	add(rulePegText, position1199)
																}
																{
																	add(ruleAction80, position)
																}
																add(ruleOtir, position1198)
															}
															goto l1174
														l1197:
															position, tokenIndex = position1174, tokenIndex1174
															{
																position1210 := position
																{
																	position1211 := position
																	{
																		position1212, tokenIndex1212 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1213
																		}
																		position++
																		goto l1212
																	l1213:
																		position, tokenIndex = position1212, tokenIndex1212
																		if buffer[position] != rune('O') {
																			goto l1209
																		}
																		position++
																	}
																l1212:
																	{
																		position1214, tokenIndex1214 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1215
																		}
																		position++
																		goto l1214
																	l1215:
																		position, tokenIndex = position1214, tokenIndex1214
																		if buffer[position] != rune('U') {
																			goto l1209
																		}
																		position++
																	}
																l1214:
																	{
																		position1216, tokenIndex1216 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1217
																		}
																		position++
																		goto l1216
																	l1217:
																		position, tokenIndex = position1216, tokenIndex1216
																		if buffer[position] != rune('T') {
																			goto l1209
																		}
																		position++
																	}
																l1216:
																	{
																		position1218, tokenIndex1218 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1219
																		}
																		position++
																		goto l1218
																	l1219:
																		position, tokenIndex = position1218, tokenIndex1218
																		if buffer[position] != rune('I') {
																			goto l1209
																		}
																		position++
																	}
																l1218:
																	add(rulePegText, position1211)
																}
																{
																	add(ruleAction72, position)
																}
																add(ruleOuti, position1210)
															}
															goto l1174
														l1209:
															position, tokenIndex = position1174, tokenIndex1174
															{
																position1222 := position
																{
																	position1223 := position
																	{
																		position1224, tokenIndex1224 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1225
																		}
																		position++
																		goto l1224
																	l1225:
																		position, tokenIndex = position1224, tokenIndex1224
																		if buffer[position] != rune('I') {
																			goto l1221
																		}
																		position++
																	}
																l1224:
																	{
																		position1226, tokenIndex1226 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1227
																		}
																		position++
																		goto l1226
																	l1227:
																		position, tokenIndex = position1226, tokenIndex1226
																		if buffer[position] != rune('N') {
																			goto l1221
																		}
																		position++
																	}
																l1226:
																	{
																		position1228, tokenIndex1228 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1229
																		}
																		position++
																		goto l1228
																	l1229:
																		position, tokenIndex = position1228, tokenIndex1228
																		if buffer[position] != rune('D') {
																			goto l1221
																		}
																		position++
																	}
																l1228:
																	{
																		position1230, tokenIndex1230 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1231
																		}
																		position++
																		goto l1230
																	l1231:
																		position, tokenIndex = position1230, tokenIndex1230
																		if buffer[position] != rune('R') {
																			goto l1221
																		}
																		position++
																	}
																l1230:
																	add(rulePegText, position1223)
																}
																{
																	add(ruleAction83, position)
																}
																add(ruleIndr, position1222)
															}
															goto l1174
														l1221:
															position, tokenIndex = position1174, tokenIndex1174
															{
																position1234 := position
																{
																	position1235 := position
																	{
																		position1236, tokenIndex1236 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1237
																		}
																		position++
																		goto l1236
																	l1237:
																		position, tokenIndex = position1236, tokenIndex1236
																		if buffer[position] != rune('I') {
																			goto l1233
																		}
																		position++
																	}
																l1236:
																	{
																		position1238, tokenIndex1238 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1239
																		}
																		position++
																		goto l1238
																	l1239:
																		position, tokenIndex = position1238, tokenIndex1238
																		if buffer[position] != rune('N') {
																			goto l1233
																		}
																		position++
																	}
																l1238:
																	{
																		position1240, tokenIndex1240 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1241
																		}
																		position++
																		goto l1240
																	l1241:
																		position, tokenIndex = position1240, tokenIndex1240
																		if buffer[position] != rune('D') {
																			goto l1233
																		}
																		position++
																	}
																l1240:
																	add(rulePegText, position1235)
																}
																{
																	add(ruleAction75, position)
																}
																add(ruleInd, position1234)
															}
															goto l1174
														l1233:
															position, tokenIndex = position1174, tokenIndex1174
															{
																position1244 := position
																{
																	position1245 := position
																	{
																		position1246, tokenIndex1246 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1247
																		}
																		position++
																		goto l1246
																	l1247:
																		position, tokenIndex = position1246, tokenIndex1246
																		if buffer[position] != rune('O') {
																			goto l1243
																		}
																		position++
																	}
																l1246:
																	{
																		position1248, tokenIndex1248 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1249
																		}
																		position++
																		goto l1248
																	l1249:
																		position, tokenIndex = position1248, tokenIndex1248
																		if buffer[position] != rune('T') {
																			goto l1243
																		}
																		position++
																	}
																l1248:
																	{
																		position1250, tokenIndex1250 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1251
																		}
																		position++
																		goto l1250
																	l1251:
																		position, tokenIndex = position1250, tokenIndex1250
																		if buffer[position] != rune('D') {
																			goto l1243
																		}
																		position++
																	}
																l1250:
																	{
																		position1252, tokenIndex1252 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1253
																		}
																		position++
																		goto l1252
																	l1253:
																		position, tokenIndex = position1252, tokenIndex1252
																		if buffer[position] != rune('R') {
																			goto l1243
																		}
																		position++
																	}
																l1252:
																	add(rulePegText, position1245)
																}
																{
																	add(ruleAction84, position)
																}
																add(ruleOtdr, position1244)
															}
															goto l1174
														l1243:
															position, tokenIndex = position1174, tokenIndex1174
															{
																position1255 := position
																{
																	position1256 := position
																	{
																		position1257, tokenIndex1257 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1258
																		}
																		position++
																		goto l1257
																	l1258:
																		position, tokenIndex = position1257, tokenIndex1257
																		if buffer[position] != rune('O') {
																			goto l1111
																		}
																		position++
																	}
																l1257:
																	{
																		position1259, tokenIndex1259 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1260
																		}
																		position++
																		goto l1259
																	l1260:
																		position, tokenIndex = position1259, tokenIndex1259
																		if buffer[position] != rune('U') {
																			goto l1111
																		}
																		position++
																	}
																l1259:
																	{
																		position1261, tokenIndex1261 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1262
																		}
																		position++
																		goto l1261
																	l1262:
																		position, tokenIndex = position1261, tokenIndex1261
																		if buffer[position] != rune('T') {
																			goto l1111
																		}
																		position++
																	}
																l1261:
																	{
																		position1263, tokenIndex1263 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1264
																		}
																		position++
																		goto l1263
																	l1264:
																		position, tokenIndex = position1263, tokenIndex1263
																		if buffer[position] != rune('D') {
																			goto l1111
																		}
																		position++
																	}
																l1263:
																	add(rulePegText, position1256)
																}
																{
																	add(ruleAction76, position)
																}
																add(ruleOutd, position1255)
															}
														}
													l1174:
														add(ruleBlitIO, position1173)
													}
													break
												case 'R', 'r':
													{
														position1266 := position
														{
															position1267 := position
															{
																position1268, tokenIndex1268 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1269
																}
																position++
																goto l1268
															l1269:
																position, tokenIndex = position1268, tokenIndex1268
																if buffer[position] != rune('R') {
																	goto l1111
																}
																position++
															}
														l1268:
															{
																position1270, tokenIndex1270 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1271
																}
																position++
																goto l1270
															l1271:
																position, tokenIndex = position1270, tokenIndex1270
																if buffer[position] != rune('L') {
																	goto l1111
																}
																position++
															}
														l1270:
															{
																position1272, tokenIndex1272 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1273
																}
																position++
																goto l1272
															l1273:
																position, tokenIndex = position1272, tokenIndex1272
																if buffer[position] != rune('D') {
																	goto l1111
																}
																position++
															}
														l1272:
															add(rulePegText, position1267)
														}
														{
															add(ruleAction65, position)
														}
														add(ruleRld, position1266)
													}
													break
												case 'N', 'n':
													{
														position1275 := position
														{
															position1276 := position
															{
																position1277, tokenIndex1277 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1278
																}
																position++
																goto l1277
															l1278:
																position, tokenIndex = position1277, tokenIndex1277
																if buffer[position] != rune('N') {
																	goto l1111
																}
																position++
															}
														l1277:
															{
																position1279, tokenIndex1279 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1280
																}
																position++
																goto l1279
															l1280:
																position, tokenIndex = position1279, tokenIndex1279
																if buffer[position] != rune('E') {
																	goto l1111
																}
																position++
															}
														l1279:
															{
																position1281, tokenIndex1281 := position, tokenIndex
																if buffer[position] != rune('g') {
																	goto l1282
																}
																position++
																goto l1281
															l1282:
																position, tokenIndex = position1281, tokenIndex1281
																if buffer[position] != rune('G') {
																	goto l1111
																}
																position++
															}
														l1281:
															add(rulePegText, position1276)
														}
														{
															add(ruleAction61, position)
														}
														add(ruleNeg, position1275)
													}
													break
												default:
													{
														position1284 := position
														{
															position1285, tokenIndex1285 := position, tokenIndex
															{
																position1287 := position
																{
																	position1288 := position
																	{
																		position1289, tokenIndex1289 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1290
																		}
																		position++
																		goto l1289
																	l1290:
																		position, tokenIndex = position1289, tokenIndex1289
																		if buffer[position] != rune('L') {
																			goto l1286
																		}
																		position++
																	}
																l1289:
																	{
																		position1291, tokenIndex1291 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1292
																		}
																		position++
																		goto l1291
																	l1292:
																		position, tokenIndex = position1291, tokenIndex1291
																		if buffer[position] != rune('D') {
																			goto l1286
																		}
																		position++
																	}
																l1291:
																	{
																		position1293, tokenIndex1293 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1294
																		}
																		position++
																		goto l1293
																	l1294:
																		position, tokenIndex = position1293, tokenIndex1293
																		if buffer[position] != rune('I') {
																			goto l1286
																		}
																		position++
																	}
																l1293:
																	{
																		position1295, tokenIndex1295 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1296
																		}
																		position++
																		goto l1295
																	l1296:
																		position, tokenIndex = position1295, tokenIndex1295
																		if buffer[position] != rune('R') {
																			goto l1286
																		}
																		position++
																	}
																l1295:
																	add(rulePegText, position1288)
																}
																{
																	add(ruleAction77, position)
																}
																add(ruleLdir, position1287)
															}
															goto l1285
														l1286:
															position, tokenIndex = position1285, tokenIndex1285
															{
																position1299 := position
																{
																	position1300 := position
																	{
																		position1301, tokenIndex1301 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1302
																		}
																		position++
																		goto l1301
																	l1302:
																		position, tokenIndex = position1301, tokenIndex1301
																		if buffer[position] != rune('L') {
																			goto l1298
																		}
																		position++
																	}
																l1301:
																	{
																		position1303, tokenIndex1303 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1304
																		}
																		position++
																		goto l1303
																	l1304:
																		position, tokenIndex = position1303, tokenIndex1303
																		if buffer[position] != rune('D') {
																			goto l1298
																		}
																		position++
																	}
																l1303:
																	{
																		position1305, tokenIndex1305 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1306
																		}
																		position++
																		goto l1305
																	l1306:
																		position, tokenIndex = position1305, tokenIndex1305
																		if buffer[position] != rune('I') {
																			goto l1298
																		}
																		position++
																	}
																l1305:
																	add(rulePegText, position1300)
																}
																{
																	add(ruleAction69, position)
																}
																add(ruleLdi, position1299)
															}
															goto l1285
														l1298:
															position, tokenIndex = position1285, tokenIndex1285
															{
																position1309 := position
																{
																	position1310 := position
																	{
																		position1311, tokenIndex1311 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1312
																		}
																		position++
																		goto l1311
																	l1312:
																		position, tokenIndex = position1311, tokenIndex1311
																		if buffer[position] != rune('C') {
																			goto l1308
																		}
																		position++
																	}
																l1311:
																	{
																		position1313, tokenIndex1313 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1314
																		}
																		position++
																		goto l1313
																	l1314:
																		position, tokenIndex = position1313, tokenIndex1313
																		if buffer[position] != rune('P') {
																			goto l1308
																		}
																		position++
																	}
																l1313:
																	{
																		position1315, tokenIndex1315 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1316
																		}
																		position++
																		goto l1315
																	l1316:
																		position, tokenIndex = position1315, tokenIndex1315
																		if buffer[position] != rune('I') {
																			goto l1308
																		}
																		position++
																	}
																l1315:
																	{
																		position1317, tokenIndex1317 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1318
																		}
																		position++
																		goto l1317
																	l1318:
																		position, tokenIndex = position1317, tokenIndex1317
																		if buffer[position] != rune('R') {
																			goto l1308
																		}
																		position++
																	}
																l1317:
																	add(rulePegText, position1310)
																}
																{
																	add(ruleAction78, position)
																}
																add(ruleCpir, position1309)
															}
															goto l1285
														l1308:
															position, tokenIndex = position1285, tokenIndex1285
															{
																position1321 := position
																{
																	position1322 := position
																	{
																		position1323, tokenIndex1323 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1324
																		}
																		position++
																		goto l1323
																	l1324:
																		position, tokenIndex = position1323, tokenIndex1323
																		if buffer[position] != rune('C') {
																			goto l1320
																		}
																		position++
																	}
																l1323:
																	{
																		position1325, tokenIndex1325 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1326
																		}
																		position++
																		goto l1325
																	l1326:
																		position, tokenIndex = position1325, tokenIndex1325
																		if buffer[position] != rune('P') {
																			goto l1320
																		}
																		position++
																	}
																l1325:
																	{
																		position1327, tokenIndex1327 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1328
																		}
																		position++
																		goto l1327
																	l1328:
																		position, tokenIndex = position1327, tokenIndex1327
																		if buffer[position] != rune('I') {
																			goto l1320
																		}
																		position++
																	}
																l1327:
																	add(rulePegText, position1322)
																}
																{
																	add(ruleAction70, position)
																}
																add(ruleCpi, position1321)
															}
															goto l1285
														l1320:
															position, tokenIndex = position1285, tokenIndex1285
															{
																position1331 := position
																{
																	position1332 := position
																	{
																		position1333, tokenIndex1333 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1334
																		}
																		position++
																		goto l1333
																	l1334:
																		position, tokenIndex = position1333, tokenIndex1333
																		if buffer[position] != rune('L') {
																			goto l1330
																		}
																		position++
																	}
																l1333:
																	{
																		position1335, tokenIndex1335 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1336
																		}
																		position++
																		goto l1335
																	l1336:
																		position, tokenIndex = position1335, tokenIndex1335
																		if buffer[position] != rune('D') {
																			goto l1330
																		}
																		position++
																	}
																l1335:
																	{
																		position1337, tokenIndex1337 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1338
																		}
																		position++
																		goto l1337
																	l1338:
																		position, tokenIndex = position1337, tokenIndex1337
																		if buffer[position] != rune('D') {
																			goto l1330
																		}
																		position++
																	}
																l1337:
																	{
																		position1339, tokenIndex1339 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1340
																		}
																		position++
																		goto l1339
																	l1340:
																		position, tokenIndex = position1339, tokenIndex1339
																		if buffer[position] != rune('R') {
																			goto l1330
																		}
																		position++
																	}
																l1339:
																	add(rulePegText, position1332)
																}
																{
																	add(ruleAction81, position)
																}
																add(ruleLddr, position1331)
															}
															goto l1285
														l1330:
															position, tokenIndex = position1285, tokenIndex1285
															{
																position1343 := position
																{
																	position1344 := position
																	{
																		position1345, tokenIndex1345 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1346
																		}
																		position++
																		goto l1345
																	l1346:
																		position, tokenIndex = position1345, tokenIndex1345
																		if buffer[position] != rune('L') {
																			goto l1342
																		}
																		position++
																	}
																l1345:
																	{
																		position1347, tokenIndex1347 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1348
																		}
																		position++
																		goto l1347
																	l1348:
																		position, tokenIndex = position1347, tokenIndex1347
																		if buffer[position] != rune('D') {
																			goto l1342
																		}
																		position++
																	}
																l1347:
																	{
																		position1349, tokenIndex1349 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1350
																		}
																		position++
																		goto l1349
																	l1350:
																		position, tokenIndex = position1349, tokenIndex1349
																		if buffer[position] != rune('D') {
																			goto l1342
																		}
																		position++
																	}
																l1349:
																	add(rulePegText, position1344)
																}
																{
																	add(ruleAction73, position)
																}
																add(ruleLdd, position1343)
															}
															goto l1285
														l1342:
															position, tokenIndex = position1285, tokenIndex1285
															{
																position1353 := position
																{
																	position1354 := position
																	{
																		position1355, tokenIndex1355 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1356
																		}
																		position++
																		goto l1355
																	l1356:
																		position, tokenIndex = position1355, tokenIndex1355
																		if buffer[position] != rune('C') {
																			goto l1352
																		}
																		position++
																	}
																l1355:
																	{
																		position1357, tokenIndex1357 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1358
																		}
																		position++
																		goto l1357
																	l1358:
																		position, tokenIndex = position1357, tokenIndex1357
																		if buffer[position] != rune('P') {
																			goto l1352
																		}
																		position++
																	}
																l1357:
																	{
																		position1359, tokenIndex1359 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1360
																		}
																		position++
																		goto l1359
																	l1360:
																		position, tokenIndex = position1359, tokenIndex1359
																		if buffer[position] != rune('D') {
																			goto l1352
																		}
																		position++
																	}
																l1359:
																	{
																		position1361, tokenIndex1361 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1362
																		}
																		position++
																		goto l1361
																	l1362:
																		position, tokenIndex = position1361, tokenIndex1361
																		if buffer[position] != rune('R') {
																			goto l1352
																		}
																		position++
																	}
																l1361:
																	add(rulePegText, position1354)
																}
																{
																	add(ruleAction82, position)
																}
																add(ruleCpdr, position1353)
															}
															goto l1285
														l1352:
															position, tokenIndex = position1285, tokenIndex1285
															{
																position1364 := position
																{
																	position1365 := position
																	{
																		position1366, tokenIndex1366 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1367
																		}
																		position++
																		goto l1366
																	l1367:
																		position, tokenIndex = position1366, tokenIndex1366
																		if buffer[position] != rune('C') {
																			goto l1111
																		}
																		position++
																	}
																l1366:
																	{
																		position1368, tokenIndex1368 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1369
																		}
																		position++
																		goto l1368
																	l1369:
																		position, tokenIndex = position1368, tokenIndex1368
																		if buffer[position] != rune('P') {
																			goto l1111
																		}
																		position++
																	}
																l1368:
																	{
																		position1370, tokenIndex1370 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1371
																		}
																		position++
																		goto l1370
																	l1371:
																		position, tokenIndex = position1370, tokenIndex1370
																		if buffer[position] != rune('D') {
																			goto l1111
																		}
																		position++
																	}
																l1370:
																	add(rulePegText, position1365)
																}
																{
																	add(ruleAction74, position)
																}
																add(ruleCpd, position1364)
															}
														}
													l1285:
														add(ruleBlit, position1284)
													}
													break
												}
											}

										}
									l1113:
										add(ruleEDSimple, position1112)
									}
									goto l803
								l1111:
									position, tokenIndex = position803, tokenIndex803
									{
										position1374 := position
										{
											position1375, tokenIndex1375 := position, tokenIndex
											{
												position1377 := position
												{
													position1378 := position
													{
														position1379, tokenIndex1379 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1380
														}
														position++
														goto l1379
													l1380:
														position, tokenIndex = position1379, tokenIndex1379
														if buffer[position] != rune('R') {
															goto l1376
														}
														position++
													}
												l1379:
													{
														position1381, tokenIndex1381 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1382
														}
														position++
														goto l1381
													l1382:
														position, tokenIndex = position1381, tokenIndex1381
														if buffer[position] != rune('L') {
															goto l1376
														}
														position++
													}
												l1381:
													{
														position1383, tokenIndex1383 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1384
														}
														position++
														goto l1383
													l1384:
														position, tokenIndex = position1383, tokenIndex1383
														if buffer[position] != rune('C') {
															goto l1376
														}
														position++
													}
												l1383:
													{
														position1385, tokenIndex1385 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1386
														}
														position++
														goto l1385
													l1386:
														position, tokenIndex = position1385, tokenIndex1385
														if buffer[position] != rune('A') {
															goto l1376
														}
														position++
													}
												l1385:
													add(rulePegText, position1378)
												}
												{
													add(ruleAction50, position)
												}
												add(ruleRlca, position1377)
											}
											goto l1375
										l1376:
											position, tokenIndex = position1375, tokenIndex1375
											{
												position1389 := position
												{
													position1390 := position
													{
														position1391, tokenIndex1391 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1392
														}
														position++
														goto l1391
													l1392:
														position, tokenIndex = position1391, tokenIndex1391
														if buffer[position] != rune('R') {
															goto l1388
														}
														position++
													}
												l1391:
													{
														position1393, tokenIndex1393 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1394
														}
														position++
														goto l1393
													l1394:
														position, tokenIndex = position1393, tokenIndex1393
														if buffer[position] != rune('R') {
															goto l1388
														}
														position++
													}
												l1393:
													{
														position1395, tokenIndex1395 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1396
														}
														position++
														goto l1395
													l1396:
														position, tokenIndex = position1395, tokenIndex1395
														if buffer[position] != rune('C') {
															goto l1388
														}
														position++
													}
												l1395:
													{
														position1397, tokenIndex1397 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1398
														}
														position++
														goto l1397
													l1398:
														position, tokenIndex = position1397, tokenIndex1397
														if buffer[position] != rune('A') {
															goto l1388
														}
														position++
													}
												l1397:
													add(rulePegText, position1390)
												}
												{
													add(ruleAction51, position)
												}
												add(ruleRrca, position1389)
											}
											goto l1375
										l1388:
											position, tokenIndex = position1375, tokenIndex1375
											{
												position1401 := position
												{
													position1402 := position
													{
														position1403, tokenIndex1403 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1404
														}
														position++
														goto l1403
													l1404:
														position, tokenIndex = position1403, tokenIndex1403
														if buffer[position] != rune('R') {
															goto l1400
														}
														position++
													}
												l1403:
													{
														position1405, tokenIndex1405 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1406
														}
														position++
														goto l1405
													l1406:
														position, tokenIndex = position1405, tokenIndex1405
														if buffer[position] != rune('L') {
															goto l1400
														}
														position++
													}
												l1405:
													{
														position1407, tokenIndex1407 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1408
														}
														position++
														goto l1407
													l1408:
														position, tokenIndex = position1407, tokenIndex1407
														if buffer[position] != rune('A') {
															goto l1400
														}
														position++
													}
												l1407:
													add(rulePegText, position1402)
												}
												{
													add(ruleAction52, position)
												}
												add(ruleRla, position1401)
											}
											goto l1375
										l1400:
											position, tokenIndex = position1375, tokenIndex1375
											{
												position1411 := position
												{
													position1412 := position
													{
														position1413, tokenIndex1413 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1414
														}
														position++
														goto l1413
													l1414:
														position, tokenIndex = position1413, tokenIndex1413
														if buffer[position] != rune('D') {
															goto l1410
														}
														position++
													}
												l1413:
													{
														position1415, tokenIndex1415 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1416
														}
														position++
														goto l1415
													l1416:
														position, tokenIndex = position1415, tokenIndex1415
														if buffer[position] != rune('A') {
															goto l1410
														}
														position++
													}
												l1415:
													{
														position1417, tokenIndex1417 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1418
														}
														position++
														goto l1417
													l1418:
														position, tokenIndex = position1417, tokenIndex1417
														if buffer[position] != rune('A') {
															goto l1410
														}
														position++
													}
												l1417:
													add(rulePegText, position1412)
												}
												{
													add(ruleAction54, position)
												}
												add(ruleDaa, position1411)
											}
											goto l1375
										l1410:
											position, tokenIndex = position1375, tokenIndex1375
											{
												position1421 := position
												{
													position1422 := position
													{
														position1423, tokenIndex1423 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1424
														}
														position++
														goto l1423
													l1424:
														position, tokenIndex = position1423, tokenIndex1423
														if buffer[position] != rune('C') {
															goto l1420
														}
														position++
													}
												l1423:
													{
														position1425, tokenIndex1425 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1426
														}
														position++
														goto l1425
													l1426:
														position, tokenIndex = position1425, tokenIndex1425
														if buffer[position] != rune('P') {
															goto l1420
														}
														position++
													}
												l1425:
													{
														position1427, tokenIndex1427 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1428
														}
														position++
														goto l1427
													l1428:
														position, tokenIndex = position1427, tokenIndex1427
														if buffer[position] != rune('L') {
															goto l1420
														}
														position++
													}
												l1427:
													add(rulePegText, position1422)
												}
												{
													add(ruleAction55, position)
												}
												add(ruleCpl, position1421)
											}
											goto l1375
										l1420:
											position, tokenIndex = position1375, tokenIndex1375
											{
												position1431 := position
												{
													position1432 := position
													{
														position1433, tokenIndex1433 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1434
														}
														position++
														goto l1433
													l1434:
														position, tokenIndex = position1433, tokenIndex1433
														if buffer[position] != rune('E') {
															goto l1430
														}
														position++
													}
												l1433:
													{
														position1435, tokenIndex1435 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1436
														}
														position++
														goto l1435
													l1436:
														position, tokenIndex = position1435, tokenIndex1435
														if buffer[position] != rune('X') {
															goto l1430
														}
														position++
													}
												l1435:
													{
														position1437, tokenIndex1437 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1438
														}
														position++
														goto l1437
													l1438:
														position, tokenIndex = position1437, tokenIndex1437
														if buffer[position] != rune('X') {
															goto l1430
														}
														position++
													}
												l1437:
													add(rulePegText, position1432)
												}
												{
													add(ruleAction58, position)
												}
												add(ruleExx, position1431)
											}
											goto l1375
										l1430:
											position, tokenIndex = position1375, tokenIndex1375
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position1441 := position
														{
															position1442 := position
															{
																position1443, tokenIndex1443 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1444
																}
																position++
																goto l1443
															l1444:
																position, tokenIndex = position1443, tokenIndex1443
																if buffer[position] != rune('E') {
																	goto l1373
																}
																position++
															}
														l1443:
															{
																position1445, tokenIndex1445 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1446
																}
																position++
																goto l1445
															l1446:
																position, tokenIndex = position1445, tokenIndex1445
																if buffer[position] != rune('I') {
																	goto l1373
																}
																position++
															}
														l1445:
															add(rulePegText, position1442)
														}
														{
															add(ruleAction60, position)
														}
														add(ruleEi, position1441)
													}
													break
												case 'D', 'd':
													{
														position1448 := position
														{
															position1449 := position
															{
																position1450, tokenIndex1450 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1451
																}
																position++
																goto l1450
															l1451:
																position, tokenIndex = position1450, tokenIndex1450
																if buffer[position] != rune('D') {
																	goto l1373
																}
																position++
															}
														l1450:
															{
																position1452, tokenIndex1452 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1453
																}
																position++
																goto l1452
															l1453:
																position, tokenIndex = position1452, tokenIndex1452
																if buffer[position] != rune('I') {
																	goto l1373
																}
																position++
															}
														l1452:
															add(rulePegText, position1449)
														}
														{
															add(ruleAction59, position)
														}
														add(ruleDi, position1448)
													}
													break
												case 'C', 'c':
													{
														position1455 := position
														{
															position1456 := position
															{
																position1457, tokenIndex1457 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1458
																}
																position++
																goto l1457
															l1458:
																position, tokenIndex = position1457, tokenIndex1457
																if buffer[position] != rune('C') {
																	goto l1373
																}
																position++
															}
														l1457:
															{
																position1459, tokenIndex1459 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1460
																}
																position++
																goto l1459
															l1460:
																position, tokenIndex = position1459, tokenIndex1459
																if buffer[position] != rune('C') {
																	goto l1373
																}
																position++
															}
														l1459:
															{
																position1461, tokenIndex1461 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1462
																}
																position++
																goto l1461
															l1462:
																position, tokenIndex = position1461, tokenIndex1461
																if buffer[position] != rune('F') {
																	goto l1373
																}
																position++
															}
														l1461:
															add(rulePegText, position1456)
														}
														{
															add(ruleAction57, position)
														}
														add(ruleCcf, position1455)
													}
													break
												case 'S', 's':
													{
														position1464 := position
														{
															position1465 := position
															{
																position1466, tokenIndex1466 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1467
																}
																position++
																goto l1466
															l1467:
																position, tokenIndex = position1466, tokenIndex1466
																if buffer[position] != rune('S') {
																	goto l1373
																}
																position++
															}
														l1466:
															{
																position1468, tokenIndex1468 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1469
																}
																position++
																goto l1468
															l1469:
																position, tokenIndex = position1468, tokenIndex1468
																if buffer[position] != rune('C') {
																	goto l1373
																}
																position++
															}
														l1468:
															{
																position1470, tokenIndex1470 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1471
																}
																position++
																goto l1470
															l1471:
																position, tokenIndex = position1470, tokenIndex1470
																if buffer[position] != rune('F') {
																	goto l1373
																}
																position++
															}
														l1470:
															add(rulePegText, position1465)
														}
														{
															add(ruleAction56, position)
														}
														add(ruleScf, position1464)
													}
													break
												case 'R', 'r':
													{
														position1473 := position
														{
															position1474 := position
															{
																position1475, tokenIndex1475 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1476
																}
																position++
																goto l1475
															l1476:
																position, tokenIndex = position1475, tokenIndex1475
																if buffer[position] != rune('R') {
																	goto l1373
																}
																position++
															}
														l1475:
															{
																position1477, tokenIndex1477 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1478
																}
																position++
																goto l1477
															l1478:
																position, tokenIndex = position1477, tokenIndex1477
																if buffer[position] != rune('R') {
																	goto l1373
																}
																position++
															}
														l1477:
															{
																position1479, tokenIndex1479 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1480
																}
																position++
																goto l1479
															l1480:
																position, tokenIndex = position1479, tokenIndex1479
																if buffer[position] != rune('A') {
																	goto l1373
																}
																position++
															}
														l1479:
															add(rulePegText, position1474)
														}
														{
															add(ruleAction53, position)
														}
														add(ruleRra, position1473)
													}
													break
												case 'H', 'h':
													{
														position1482 := position
														{
															position1483 := position
															{
																position1484, tokenIndex1484 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l1485
																}
																position++
																goto l1484
															l1485:
																position, tokenIndex = position1484, tokenIndex1484
																if buffer[position] != rune('H') {
																	goto l1373
																}
																position++
															}
														l1484:
															{
																position1486, tokenIndex1486 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1487
																}
																position++
																goto l1486
															l1487:
																position, tokenIndex = position1486, tokenIndex1486
																if buffer[position] != rune('A') {
																	goto l1373
																}
																position++
															}
														l1486:
															{
																position1488, tokenIndex1488 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1489
																}
																position++
																goto l1488
															l1489:
																position, tokenIndex = position1488, tokenIndex1488
																if buffer[position] != rune('L') {
																	goto l1373
																}
																position++
															}
														l1488:
															{
																position1490, tokenIndex1490 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1491
																}
																position++
																goto l1490
															l1491:
																position, tokenIndex = position1490, tokenIndex1490
																if buffer[position] != rune('T') {
																	goto l1373
																}
																position++
															}
														l1490:
															add(rulePegText, position1483)
														}
														{
															add(ruleAction49, position)
														}
														add(ruleHalt, position1482)
													}
													break
												default:
													{
														position1493 := position
														{
															position1494 := position
															{
																position1495, tokenIndex1495 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1496
																}
																position++
																goto l1495
															l1496:
																position, tokenIndex = position1495, tokenIndex1495
																if buffer[position] != rune('N') {
																	goto l1373
																}
																position++
															}
														l1495:
															{
																position1497, tokenIndex1497 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l1498
																}
																position++
																goto l1497
															l1498:
																position, tokenIndex = position1497, tokenIndex1497
																if buffer[position] != rune('O') {
																	goto l1373
																}
																position++
															}
														l1497:
															{
																position1499, tokenIndex1499 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1500
																}
																position++
																goto l1499
															l1500:
																position, tokenIndex = position1499, tokenIndex1499
																if buffer[position] != rune('P') {
																	goto l1373
																}
																position++
															}
														l1499:
															add(rulePegText, position1494)
														}
														{
															add(ruleAction48, position)
														}
														add(ruleNop, position1493)
													}
													break
												}
											}

										}
									l1375:
										add(ruleSimple, position1374)
									}
									goto l803
								l1373:
									position, tokenIndex = position803, tokenIndex803
									{
										position1503 := position
										{
											position1504, tokenIndex1504 := position, tokenIndex
											{
												position1506 := position
												{
													position1507, tokenIndex1507 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1508
													}
													position++
													goto l1507
												l1508:
													position, tokenIndex = position1507, tokenIndex1507
													if buffer[position] != rune('R') {
														goto l1505
													}
													position++
												}
											l1507:
												{
													position1509, tokenIndex1509 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1510
													}
													position++
													goto l1509
												l1510:
													position, tokenIndex = position1509, tokenIndex1509
													if buffer[position] != rune('S') {
														goto l1505
													}
													position++
												}
											l1509:
												{
													position1511, tokenIndex1511 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1512
													}
													position++
													goto l1511
												l1512:
													position, tokenIndex = position1511, tokenIndex1511
													if buffer[position] != rune('T') {
														goto l1505
													}
													position++
												}
											l1511:
												if !_rules[rulews]() {
													goto l1505
												}
												if !_rules[rulen]() {
													goto l1505
												}
												{
													add(ruleAction85, position)
												}
												add(ruleRst, position1506)
											}
											goto l1504
										l1505:
											position, tokenIndex = position1504, tokenIndex1504
											{
												position1515 := position
												{
													position1516, tokenIndex1516 := position, tokenIndex
													if buffer[position] != rune('j') {
														goto l1517
													}
													position++
													goto l1516
												l1517:
													position, tokenIndex = position1516, tokenIndex1516
													if buffer[position] != rune('J') {
														goto l1514
													}
													position++
												}
											l1516:
												{
													position1518, tokenIndex1518 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l1519
													}
													position++
													goto l1518
												l1519:
													position, tokenIndex = position1518, tokenIndex1518
													if buffer[position] != rune('P') {
														goto l1514
													}
													position++
												}
											l1518:
												if !_rules[rulews]() {
													goto l1514
												}
												{
													position1520, tokenIndex1520 := position, tokenIndex
													if !_rules[rulecc]() {
														goto l1520
													}
													if !_rules[rulesep]() {
														goto l1520
													}
													goto l1521
												l1520:
													position, tokenIndex = position1520, tokenIndex1520
												}
											l1521:
												if !_rules[ruleSrc16]() {
													goto l1514
												}
												{
													add(ruleAction88, position)
												}
												add(ruleJp, position1515)
											}
											goto l1504
										l1514:
											position, tokenIndex = position1504, tokenIndex1504
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position1524 := position
														{
															position1525, tokenIndex1525 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1526
															}
															position++
															goto l1525
														l1526:
															position, tokenIndex = position1525, tokenIndex1525
															if buffer[position] != rune('D') {
																goto l1502
															}
															position++
														}
													l1525:
														{
															position1527, tokenIndex1527 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1528
															}
															position++
															goto l1527
														l1528:
															position, tokenIndex = position1527, tokenIndex1527
															if buffer[position] != rune('J') {
																goto l1502
															}
															position++
														}
													l1527:
														{
															position1529, tokenIndex1529 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1530
															}
															position++
															goto l1529
														l1530:
															position, tokenIndex = position1529, tokenIndex1529
															if buffer[position] != rune('N') {
																goto l1502
															}
															position++
														}
													l1529:
														{
															position1531, tokenIndex1531 := position, tokenIndex
															if buffer[position] != rune('z') {
																goto l1532
															}
															position++
															goto l1531
														l1532:
															position, tokenIndex = position1531, tokenIndex1531
															if buffer[position] != rune('Z') {
																goto l1502
															}
															position++
														}
													l1531:
														if !_rules[rulews]() {
															goto l1502
														}
														if !_rules[ruledisp]() {
															goto l1502
														}
														{
															add(ruleAction90, position)
														}
														add(ruleDjnz, position1524)
													}
													break
												case 'J', 'j':
													{
														position1534 := position
														{
															position1535, tokenIndex1535 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1536
															}
															position++
															goto l1535
														l1536:
															position, tokenIndex = position1535, tokenIndex1535
															if buffer[position] != rune('J') {
																goto l1502
															}
															position++
														}
													l1535:
														{
															position1537, tokenIndex1537 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1538
															}
															position++
															goto l1537
														l1538:
															position, tokenIndex = position1537, tokenIndex1537
															if buffer[position] != rune('R') {
																goto l1502
															}
															position++
														}
													l1537:
														if !_rules[rulews]() {
															goto l1502
														}
														{
															position1539, tokenIndex1539 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1539
															}
															if !_rules[rulesep]() {
																goto l1539
															}
															goto l1540
														l1539:
															position, tokenIndex = position1539, tokenIndex1539
														}
													l1540:
														if !_rules[ruledisp]() {
															goto l1502
														}
														{
															add(ruleAction89, position)
														}
														add(ruleJr, position1534)
													}
													break
												case 'R', 'r':
													{
														position1542 := position
														{
															position1543, tokenIndex1543 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1544
															}
															position++
															goto l1543
														l1544:
															position, tokenIndex = position1543, tokenIndex1543
															if buffer[position] != rune('R') {
																goto l1502
															}
															position++
														}
													l1543:
														{
															position1545, tokenIndex1545 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1546
															}
															position++
															goto l1545
														l1546:
															position, tokenIndex = position1545, tokenIndex1545
															if buffer[position] != rune('E') {
																goto l1502
															}
															position++
														}
													l1545:
														{
															position1547, tokenIndex1547 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1548
															}
															position++
															goto l1547
														l1548:
															position, tokenIndex = position1547, tokenIndex1547
															if buffer[position] != rune('T') {
																goto l1502
															}
															position++
														}
													l1547:
														{
															position1549, tokenIndex1549 := position, tokenIndex
															if !_rules[rulews]() {
																goto l1549
															}
															if !_rules[rulecc]() {
																goto l1549
															}
															goto l1550
														l1549:
															position, tokenIndex = position1549, tokenIndex1549
														}
													l1550:
														{
															add(ruleAction87, position)
														}
														add(ruleRet, position1542)
													}
													break
												default:
													{
														position1552 := position
														{
															position1553, tokenIndex1553 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1554
															}
															position++
															goto l1553
														l1554:
															position, tokenIndex = position1553, tokenIndex1553
															if buffer[position] != rune('C') {
																goto l1502
															}
															position++
														}
													l1553:
														{
															position1555, tokenIndex1555 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1556
															}
															position++
															goto l1555
														l1556:
															position, tokenIndex = position1555, tokenIndex1555
															if buffer[position] != rune('A') {
																goto l1502
															}
															position++
														}
													l1555:
														{
															position1557, tokenIndex1557 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1558
															}
															position++
															goto l1557
														l1558:
															position, tokenIndex = position1557, tokenIndex1557
															if buffer[position] != rune('L') {
																goto l1502
															}
															position++
														}
													l1557:
														{
															position1559, tokenIndex1559 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1560
															}
															position++
															goto l1559
														l1560:
															position, tokenIndex = position1559, tokenIndex1559
															if buffer[position] != rune('L') {
																goto l1502
															}
															position++
														}
													l1559:
														if !_rules[rulews]() {
															goto l1502
														}
														{
															position1561, tokenIndex1561 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1561
															}
															if !_rules[rulesep]() {
																goto l1561
															}
															goto l1562
														l1561:
															position, tokenIndex = position1561, tokenIndex1561
														}
													l1562:
														if !_rules[ruleSrc16]() {
															goto l1502
														}
														{
															add(ruleAction86, position)
														}
														add(ruleCall, position1552)
													}
													break
												}
											}

										}
									l1504:
										add(ruleJump, position1503)
									}
									goto l803
								l1502:
									position, tokenIndex = position803, tokenIndex803
									{
										position1564 := position
										{
											position1565, tokenIndex1565 := position, tokenIndex
											{
												position1567 := position
												{
													position1568, tokenIndex1568 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1569
													}
													position++
													goto l1568
												l1569:
													position, tokenIndex = position1568, tokenIndex1568
													if buffer[position] != rune('I') {
														goto l1566
													}
													position++
												}
											l1568:
												{
													position1570, tokenIndex1570 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1571
													}
													position++
													goto l1570
												l1571:
													position, tokenIndex = position1570, tokenIndex1570
													if buffer[position] != rune('N') {
														goto l1566
													}
													position++
												}
											l1570:
												if !_rules[rulews]() {
													goto l1566
												}
												if !_rules[ruleReg8]() {
													goto l1566
												}
												if !_rules[rulesep]() {
													goto l1566
												}
												if !_rules[rulePort]() {
													goto l1566
												}
												{
													add(ruleAction91, position)
												}
												add(ruleIN, position1567)
											}
											goto l1565
										l1566:
											position, tokenIndex = position1565, tokenIndex1565
											{
												position1573 := position
												{
													position1574, tokenIndex1574 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1575
													}
													position++
													goto l1574
												l1575:
													position, tokenIndex = position1574, tokenIndex1574
													if buffer[position] != rune('O') {
														goto l3
													}
													position++
												}
											l1574:
												{
													position1576, tokenIndex1576 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1577
													}
													position++
													goto l1576
												l1577:
													position, tokenIndex = position1576, tokenIndex1576
													if buffer[position] != rune('U') {
														goto l3
													}
													position++
												}
											l1576:
												{
													position1578, tokenIndex1578 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1579
													}
													position++
													goto l1578
												l1579:
													position, tokenIndex = position1578, tokenIndex1578
													if buffer[position] != rune('T') {
														goto l3
													}
													position++
												}
											l1578:
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
													add(ruleAction92, position)
												}
												add(ruleOUT, position1573)
											}
										}
									l1565:
										add(ruleIO, position1564)
									}
								}
							l803:
								add(ruleInstruction, position800)
							}
							{
								position1581, tokenIndex1581 := position, tokenIndex
								if buffer[position] != rune('\n') {
									goto l1581
								}
								position++
								goto l1582
							l1581:
								position, tokenIndex = position1581, tokenIndex1581
							}
						l1582:
							{
								add(ruleAction0, position)
							}
							add(ruleLine, position799)
						}
					}
				l794:
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
		/* 3 Instruction <- <(ws* (Assignment / Inc / Dec / Alu16 / Alu / BitOp / EDSimple / Simple / Jump / IO))> */
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
			position1607, tokenIndex1607 := position, tokenIndex
			{
				position1608 := position
				{
					position1609, tokenIndex1609 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1610
					}
					goto l1609
				l1610:
					position, tokenIndex = position1609, tokenIndex1609
					if !_rules[ruleReg8]() {
						goto l1611
					}
					goto l1609
				l1611:
					position, tokenIndex = position1609, tokenIndex1609
					if !_rules[ruleReg16Contents]() {
						goto l1612
					}
					goto l1609
				l1612:
					position, tokenIndex = position1609, tokenIndex1609
					if !_rules[rulenn_contents]() {
						goto l1607
					}
				}
			l1609:
				{
					add(ruleAction16, position)
				}
				add(ruleSrc8, position1608)
			}
			return true
		l1607:
			position, tokenIndex = position1607, tokenIndex1607
			return false
		},
		/* 25 Loc8 <- <((Reg8 / Reg16Contents) Action17)> */
		func() bool {
			position1614, tokenIndex1614 := position, tokenIndex
			{
				position1615 := position
				{
					position1616, tokenIndex1616 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1617
					}
					goto l1616
				l1617:
					position, tokenIndex = position1616, tokenIndex1616
					if !_rules[ruleReg16Contents]() {
						goto l1614
					}
				}
			l1616:
				{
					add(ruleAction17, position)
				}
				add(ruleLoc8, position1615)
			}
			return true
		l1614:
			position, tokenIndex = position1614, tokenIndex1614
			return false
		},
		/* 26 ILoc8 <- <(IReg8 Action18)> */
		func() bool {
			position1619, tokenIndex1619 := position, tokenIndex
			{
				position1620 := position
				if !_rules[ruleIReg8]() {
					goto l1619
				}
				{
					add(ruleAction18, position)
				}
				add(ruleILoc8, position1620)
			}
			return true
		l1619:
			position, tokenIndex = position1619, tokenIndex1619
			return false
		},
		/* 27 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action19)> */
		func() bool {
			position1622, tokenIndex1622 := position, tokenIndex
			{
				position1623 := position
				{
					position1624 := position
					{
						position1625, tokenIndex1625 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1626
						}
						goto l1625
					l1626:
						position, tokenIndex = position1625, tokenIndex1625
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1628 := position
									{
										position1629, tokenIndex1629 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1630
										}
										position++
										goto l1629
									l1630:
										position, tokenIndex = position1629, tokenIndex1629
										if buffer[position] != rune('R') {
											goto l1622
										}
										position++
									}
								l1629:
									add(ruleR, position1628)
								}
								break
							case 'I', 'i':
								{
									position1631 := position
									{
										position1632, tokenIndex1632 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1633
										}
										position++
										goto l1632
									l1633:
										position, tokenIndex = position1632, tokenIndex1632
										if buffer[position] != rune('I') {
											goto l1622
										}
										position++
									}
								l1632:
									add(ruleI, position1631)
								}
								break
							case 'L', 'l':
								{
									position1634 := position
									{
										position1635, tokenIndex1635 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1636
										}
										position++
										goto l1635
									l1636:
										position, tokenIndex = position1635, tokenIndex1635
										if buffer[position] != rune('L') {
											goto l1622
										}
										position++
									}
								l1635:
									add(ruleL, position1634)
								}
								break
							case 'H', 'h':
								{
									position1637 := position
									{
										position1638, tokenIndex1638 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1639
										}
										position++
										goto l1638
									l1639:
										position, tokenIndex = position1638, tokenIndex1638
										if buffer[position] != rune('H') {
											goto l1622
										}
										position++
									}
								l1638:
									add(ruleH, position1637)
								}
								break
							case 'E', 'e':
								{
									position1640 := position
									{
										position1641, tokenIndex1641 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1642
										}
										position++
										goto l1641
									l1642:
										position, tokenIndex = position1641, tokenIndex1641
										if buffer[position] != rune('E') {
											goto l1622
										}
										position++
									}
								l1641:
									add(ruleE, position1640)
								}
								break
							case 'D', 'd':
								{
									position1643 := position
									{
										position1644, tokenIndex1644 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1645
										}
										position++
										goto l1644
									l1645:
										position, tokenIndex = position1644, tokenIndex1644
										if buffer[position] != rune('D') {
											goto l1622
										}
										position++
									}
								l1644:
									add(ruleD, position1643)
								}
								break
							case 'C', 'c':
								{
									position1646 := position
									{
										position1647, tokenIndex1647 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1648
										}
										position++
										goto l1647
									l1648:
										position, tokenIndex = position1647, tokenIndex1647
										if buffer[position] != rune('C') {
											goto l1622
										}
										position++
									}
								l1647:
									add(ruleC, position1646)
								}
								break
							case 'B', 'b':
								{
									position1649 := position
									{
										position1650, tokenIndex1650 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1651
										}
										position++
										goto l1650
									l1651:
										position, tokenIndex = position1650, tokenIndex1650
										if buffer[position] != rune('B') {
											goto l1622
										}
										position++
									}
								l1650:
									add(ruleB, position1649)
								}
								break
							case 'F', 'f':
								{
									position1652 := position
									{
										position1653, tokenIndex1653 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1654
										}
										position++
										goto l1653
									l1654:
										position, tokenIndex = position1653, tokenIndex1653
										if buffer[position] != rune('F') {
											goto l1622
										}
										position++
									}
								l1653:
									add(ruleF, position1652)
								}
								break
							default:
								{
									position1655 := position
									{
										position1656, tokenIndex1656 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1657
										}
										position++
										goto l1656
									l1657:
										position, tokenIndex = position1656, tokenIndex1656
										if buffer[position] != rune('A') {
											goto l1622
										}
										position++
									}
								l1656:
									add(ruleA, position1655)
								}
								break
							}
						}

					}
				l1625:
					add(rulePegText, position1624)
				}
				{
					add(ruleAction19, position)
				}
				add(ruleReg8, position1623)
			}
			return true
		l1622:
			position, tokenIndex = position1622, tokenIndex1622
			return false
		},
		/* 28 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action20)> */
		func() bool {
			position1659, tokenIndex1659 := position, tokenIndex
			{
				position1660 := position
				{
					position1661 := position
					{
						position1662, tokenIndex1662 := position, tokenIndex
						{
							position1664 := position
							{
								position1665, tokenIndex1665 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1666
								}
								position++
								goto l1665
							l1666:
								position, tokenIndex = position1665, tokenIndex1665
								if buffer[position] != rune('I') {
									goto l1663
								}
								position++
							}
						l1665:
							{
								position1667, tokenIndex1667 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1668
								}
								position++
								goto l1667
							l1668:
								position, tokenIndex = position1667, tokenIndex1667
								if buffer[position] != rune('X') {
									goto l1663
								}
								position++
							}
						l1667:
							{
								position1669, tokenIndex1669 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1670
								}
								position++
								goto l1669
							l1670:
								position, tokenIndex = position1669, tokenIndex1669
								if buffer[position] != rune('H') {
									goto l1663
								}
								position++
							}
						l1669:
							add(ruleIXH, position1664)
						}
						goto l1662
					l1663:
						position, tokenIndex = position1662, tokenIndex1662
						{
							position1672 := position
							{
								position1673, tokenIndex1673 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1674
								}
								position++
								goto l1673
							l1674:
								position, tokenIndex = position1673, tokenIndex1673
								if buffer[position] != rune('I') {
									goto l1671
								}
								position++
							}
						l1673:
							{
								position1675, tokenIndex1675 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1676
								}
								position++
								goto l1675
							l1676:
								position, tokenIndex = position1675, tokenIndex1675
								if buffer[position] != rune('X') {
									goto l1671
								}
								position++
							}
						l1675:
							{
								position1677, tokenIndex1677 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1678
								}
								position++
								goto l1677
							l1678:
								position, tokenIndex = position1677, tokenIndex1677
								if buffer[position] != rune('L') {
									goto l1671
								}
								position++
							}
						l1677:
							add(ruleIXL, position1672)
						}
						goto l1662
					l1671:
						position, tokenIndex = position1662, tokenIndex1662
						{
							position1680 := position
							{
								position1681, tokenIndex1681 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1682
								}
								position++
								goto l1681
							l1682:
								position, tokenIndex = position1681, tokenIndex1681
								if buffer[position] != rune('I') {
									goto l1679
								}
								position++
							}
						l1681:
							{
								position1683, tokenIndex1683 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1684
								}
								position++
								goto l1683
							l1684:
								position, tokenIndex = position1683, tokenIndex1683
								if buffer[position] != rune('Y') {
									goto l1679
								}
								position++
							}
						l1683:
							{
								position1685, tokenIndex1685 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1686
								}
								position++
								goto l1685
							l1686:
								position, tokenIndex = position1685, tokenIndex1685
								if buffer[position] != rune('H') {
									goto l1679
								}
								position++
							}
						l1685:
							add(ruleIYH, position1680)
						}
						goto l1662
					l1679:
						position, tokenIndex = position1662, tokenIndex1662
						{
							position1687 := position
							{
								position1688, tokenIndex1688 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1689
								}
								position++
								goto l1688
							l1689:
								position, tokenIndex = position1688, tokenIndex1688
								if buffer[position] != rune('I') {
									goto l1659
								}
								position++
							}
						l1688:
							{
								position1690, tokenIndex1690 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1691
								}
								position++
								goto l1690
							l1691:
								position, tokenIndex = position1690, tokenIndex1690
								if buffer[position] != rune('Y') {
									goto l1659
								}
								position++
							}
						l1690:
							{
								position1692, tokenIndex1692 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1693
								}
								position++
								goto l1692
							l1693:
								position, tokenIndex = position1692, tokenIndex1692
								if buffer[position] != rune('L') {
									goto l1659
								}
								position++
							}
						l1692:
							add(ruleIYL, position1687)
						}
					}
				l1662:
					add(rulePegText, position1661)
				}
				{
					add(ruleAction20, position)
				}
				add(ruleIReg8, position1660)
			}
			return true
		l1659:
			position, tokenIndex = position1659, tokenIndex1659
			return false
		},
		/* 29 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action21)> */
		func() bool {
			position1695, tokenIndex1695 := position, tokenIndex
			{
				position1696 := position
				{
					position1697, tokenIndex1697 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1698
					}
					goto l1697
				l1698:
					position, tokenIndex = position1697, tokenIndex1697
					if !_rules[rulenn_contents]() {
						goto l1699
					}
					goto l1697
				l1699:
					position, tokenIndex = position1697, tokenIndex1697
					if !_rules[ruleReg16Contents]() {
						goto l1695
					}
				}
			l1697:
				{
					add(ruleAction21, position)
				}
				add(ruleDst16, position1696)
			}
			return true
		l1695:
			position, tokenIndex = position1695, tokenIndex1695
			return false
		},
		/* 30 Src16 <- <((Reg16 / nn / nn_contents) Action22)> */
		func() bool {
			position1701, tokenIndex1701 := position, tokenIndex
			{
				position1702 := position
				{
					position1703, tokenIndex1703 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1704
					}
					goto l1703
				l1704:
					position, tokenIndex = position1703, tokenIndex1703
					if !_rules[rulenn]() {
						goto l1705
					}
					goto l1703
				l1705:
					position, tokenIndex = position1703, tokenIndex1703
					if !_rules[rulenn_contents]() {
						goto l1701
					}
				}
			l1703:
				{
					add(ruleAction22, position)
				}
				add(ruleSrc16, position1702)
			}
			return true
		l1701:
			position, tokenIndex = position1701, tokenIndex1701
			return false
		},
		/* 31 Loc16 <- <(Reg16 Action23)> */
		func() bool {
			position1707, tokenIndex1707 := position, tokenIndex
			{
				position1708 := position
				if !_rules[ruleReg16]() {
					goto l1707
				}
				{
					add(ruleAction23, position)
				}
				add(ruleLoc16, position1708)
			}
			return true
		l1707:
			position, tokenIndex = position1707, tokenIndex1707
			return false
		},
		/* 32 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action24)> */
		func() bool {
			position1710, tokenIndex1710 := position, tokenIndex
			{
				position1711 := position
				{
					position1712 := position
					{
						position1713, tokenIndex1713 := position, tokenIndex
						{
							position1715 := position
							{
								position1716, tokenIndex1716 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1717
								}
								position++
								goto l1716
							l1717:
								position, tokenIndex = position1716, tokenIndex1716
								if buffer[position] != rune('A') {
									goto l1714
								}
								position++
							}
						l1716:
							{
								position1718, tokenIndex1718 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1719
								}
								position++
								goto l1718
							l1719:
								position, tokenIndex = position1718, tokenIndex1718
								if buffer[position] != rune('F') {
									goto l1714
								}
								position++
							}
						l1718:
							if buffer[position] != rune('\'') {
								goto l1714
							}
							position++
							add(ruleAF_PRIME, position1715)
						}
						goto l1713
					l1714:
						position, tokenIndex = position1713, tokenIndex1713
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1710
								}
								break
							case 'S', 's':
								{
									position1721 := position
									{
										position1722, tokenIndex1722 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1723
										}
										position++
										goto l1722
									l1723:
										position, tokenIndex = position1722, tokenIndex1722
										if buffer[position] != rune('S') {
											goto l1710
										}
										position++
									}
								l1722:
									{
										position1724, tokenIndex1724 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1725
										}
										position++
										goto l1724
									l1725:
										position, tokenIndex = position1724, tokenIndex1724
										if buffer[position] != rune('P') {
											goto l1710
										}
										position++
									}
								l1724:
									add(ruleSP, position1721)
								}
								break
							case 'H', 'h':
								{
									position1726 := position
									{
										position1727, tokenIndex1727 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1728
										}
										position++
										goto l1727
									l1728:
										position, tokenIndex = position1727, tokenIndex1727
										if buffer[position] != rune('H') {
											goto l1710
										}
										position++
									}
								l1727:
									{
										position1729, tokenIndex1729 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1730
										}
										position++
										goto l1729
									l1730:
										position, tokenIndex = position1729, tokenIndex1729
										if buffer[position] != rune('L') {
											goto l1710
										}
										position++
									}
								l1729:
									add(ruleHL, position1726)
								}
								break
							case 'D', 'd':
								{
									position1731 := position
									{
										position1732, tokenIndex1732 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1733
										}
										position++
										goto l1732
									l1733:
										position, tokenIndex = position1732, tokenIndex1732
										if buffer[position] != rune('D') {
											goto l1710
										}
										position++
									}
								l1732:
									{
										position1734, tokenIndex1734 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1735
										}
										position++
										goto l1734
									l1735:
										position, tokenIndex = position1734, tokenIndex1734
										if buffer[position] != rune('E') {
											goto l1710
										}
										position++
									}
								l1734:
									add(ruleDE, position1731)
								}
								break
							case 'B', 'b':
								{
									position1736 := position
									{
										position1737, tokenIndex1737 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1738
										}
										position++
										goto l1737
									l1738:
										position, tokenIndex = position1737, tokenIndex1737
										if buffer[position] != rune('B') {
											goto l1710
										}
										position++
									}
								l1737:
									{
										position1739, tokenIndex1739 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1740
										}
										position++
										goto l1739
									l1740:
										position, tokenIndex = position1739, tokenIndex1739
										if buffer[position] != rune('C') {
											goto l1710
										}
										position++
									}
								l1739:
									add(ruleBC, position1736)
								}
								break
							default:
								{
									position1741 := position
									{
										position1742, tokenIndex1742 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1743
										}
										position++
										goto l1742
									l1743:
										position, tokenIndex = position1742, tokenIndex1742
										if buffer[position] != rune('A') {
											goto l1710
										}
										position++
									}
								l1742:
									{
										position1744, tokenIndex1744 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1745
										}
										position++
										goto l1744
									l1745:
										position, tokenIndex = position1744, tokenIndex1744
										if buffer[position] != rune('F') {
											goto l1710
										}
										position++
									}
								l1744:
									add(ruleAF, position1741)
								}
								break
							}
						}

					}
				l1713:
					add(rulePegText, position1712)
				}
				{
					add(ruleAction24, position)
				}
				add(ruleReg16, position1711)
			}
			return true
		l1710:
			position, tokenIndex = position1710, tokenIndex1710
			return false
		},
		/* 33 IReg16 <- <(<(IX / IY)> Action25)> */
		func() bool {
			position1747, tokenIndex1747 := position, tokenIndex
			{
				position1748 := position
				{
					position1749 := position
					{
						position1750, tokenIndex1750 := position, tokenIndex
						{
							position1752 := position
							{
								position1753, tokenIndex1753 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1754
								}
								position++
								goto l1753
							l1754:
								position, tokenIndex = position1753, tokenIndex1753
								if buffer[position] != rune('I') {
									goto l1751
								}
								position++
							}
						l1753:
							{
								position1755, tokenIndex1755 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1756
								}
								position++
								goto l1755
							l1756:
								position, tokenIndex = position1755, tokenIndex1755
								if buffer[position] != rune('X') {
									goto l1751
								}
								position++
							}
						l1755:
							add(ruleIX, position1752)
						}
						goto l1750
					l1751:
						position, tokenIndex = position1750, tokenIndex1750
						{
							position1757 := position
							{
								position1758, tokenIndex1758 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1759
								}
								position++
								goto l1758
							l1759:
								position, tokenIndex = position1758, tokenIndex1758
								if buffer[position] != rune('I') {
									goto l1747
								}
								position++
							}
						l1758:
							{
								position1760, tokenIndex1760 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1761
								}
								position++
								goto l1760
							l1761:
								position, tokenIndex = position1760, tokenIndex1760
								if buffer[position] != rune('Y') {
									goto l1747
								}
								position++
							}
						l1760:
							add(ruleIY, position1757)
						}
					}
				l1750:
					add(rulePegText, position1749)
				}
				{
					add(ruleAction25, position)
				}
				add(ruleIReg16, position1748)
			}
			return true
		l1747:
			position, tokenIndex = position1747, tokenIndex1747
			return false
		},
		/* 34 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1763, tokenIndex1763 := position, tokenIndex
			{
				position1764 := position
				{
					position1765, tokenIndex1765 := position, tokenIndex
					{
						position1767 := position
						if buffer[position] != rune('(') {
							goto l1766
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1766
						}
						{
							position1768, tokenIndex1768 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1768
							}
							goto l1769
						l1768:
							position, tokenIndex = position1768, tokenIndex1768
						}
					l1769:
						if !_rules[rulesignedDecimalByte]() {
							goto l1766
						}
						{
							position1770, tokenIndex1770 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1770
							}
							goto l1771
						l1770:
							position, tokenIndex = position1770, tokenIndex1770
						}
					l1771:
						if buffer[position] != rune(')') {
							goto l1766
						}
						position++
						{
							add(ruleAction27, position)
						}
						add(ruleIndexedR16C, position1767)
					}
					goto l1765
				l1766:
					position, tokenIndex = position1765, tokenIndex1765
					{
						position1773 := position
						if buffer[position] != rune('(') {
							goto l1763
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1763
						}
						if buffer[position] != rune(')') {
							goto l1763
						}
						position++
						{
							add(ruleAction26, position)
						}
						add(rulePlainR16C, position1773)
					}
				}
			l1765:
				add(ruleReg16Contents, position1764)
			}
			return true
		l1763:
			position, tokenIndex = position1763, tokenIndex1763
			return false
		},
		/* 35 PlainR16C <- <('(' Reg16 ')' Action26)> */
		nil,
		/* 36 IndexedR16C <- <('(' IReg16 ws? signedDecimalByte ws? ')' Action27)> */
		nil,
		/* 37 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1777, tokenIndex1777 := position, tokenIndex
			{
				position1778 := position
				{
					position1779, tokenIndex1779 := position, tokenIndex
					{
						position1781 := position
						{
							position1782 := position
							if !_rules[rulehexdigit]() {
								goto l1780
							}
							if !_rules[rulehexdigit]() {
								goto l1780
							}
							add(rulePegText, position1782)
						}
						{
							position1783, tokenIndex1783 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1784
							}
							position++
							goto l1783
						l1784:
							position, tokenIndex = position1783, tokenIndex1783
							if buffer[position] != rune('H') {
								goto l1780
							}
							position++
						}
					l1783:
						{
							add(ruleAction93, position)
						}
						add(rulehexByteH, position1781)
					}
					goto l1779
				l1780:
					position, tokenIndex = position1779, tokenIndex1779
					{
						position1787 := position
						if buffer[position] != rune('0') {
							goto l1786
						}
						position++
						{
							position1788, tokenIndex1788 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1789
							}
							position++
							goto l1788
						l1789:
							position, tokenIndex = position1788, tokenIndex1788
							if buffer[position] != rune('X') {
								goto l1786
							}
							position++
						}
					l1788:
						{
							position1790 := position
							if !_rules[rulehexdigit]() {
								goto l1786
							}
							if !_rules[rulehexdigit]() {
								goto l1786
							}
							add(rulePegText, position1790)
						}
						{
							add(ruleAction94, position)
						}
						add(rulehexByte0x, position1787)
					}
					goto l1779
				l1786:
					position, tokenIndex = position1779, tokenIndex1779
					{
						position1792 := position
						{
							position1793 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1777
							}
							position++
						l1794:
							{
								position1795, tokenIndex1795 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1795
								}
								position++
								goto l1794
							l1795:
								position, tokenIndex = position1795, tokenIndex1795
							}
							add(rulePegText, position1793)
						}
						{
							add(ruleAction95, position)
						}
						add(ruledecimalByte, position1792)
					}
				}
			l1779:
				add(rulen, position1778)
			}
			return true
		l1777:
			position, tokenIndex = position1777, tokenIndex1777
			return false
		},
		/* 38 nn <- <(hexWordH / hexWord0x)> */
		func() bool {
			position1797, tokenIndex1797 := position, tokenIndex
			{
				position1798 := position
				{
					position1799, tokenIndex1799 := position, tokenIndex
					{
						position1801 := position
						{
							position1802 := position
							if !_rules[rulehexdigit]() {
								goto l1800
							}
							if !_rules[rulehexdigit]() {
								goto l1800
							}
							if !_rules[rulehexdigit]() {
								goto l1800
							}
							if !_rules[rulehexdigit]() {
								goto l1800
							}
							add(rulePegText, position1802)
						}
						{
							position1803, tokenIndex1803 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1804
							}
							position++
							goto l1803
						l1804:
							position, tokenIndex = position1803, tokenIndex1803
							if buffer[position] != rune('H') {
								goto l1800
							}
							position++
						}
					l1803:
						{
							add(ruleAction96, position)
						}
						add(rulehexWordH, position1801)
					}
					goto l1799
				l1800:
					position, tokenIndex = position1799, tokenIndex1799
					{
						position1806 := position
						if buffer[position] != rune('0') {
							goto l1797
						}
						position++
						{
							position1807, tokenIndex1807 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1808
							}
							position++
							goto l1807
						l1808:
							position, tokenIndex = position1807, tokenIndex1807
							if buffer[position] != rune('X') {
								goto l1797
							}
							position++
						}
					l1807:
						{
							position1809 := position
							if !_rules[rulehexdigit]() {
								goto l1797
							}
							if !_rules[rulehexdigit]() {
								goto l1797
							}
							if !_rules[rulehexdigit]() {
								goto l1797
							}
							if !_rules[rulehexdigit]() {
								goto l1797
							}
							add(rulePegText, position1809)
						}
						{
							add(ruleAction97, position)
						}
						add(rulehexWord0x, position1806)
					}
				}
			l1799:
				add(rulenn, position1798)
			}
			return true
		l1797:
			position, tokenIndex = position1797, tokenIndex1797
			return false
		},
		/* 39 nn_contents <- <('(' nn ')' Action28)> */
		func() bool {
			position1811, tokenIndex1811 := position, tokenIndex
			{
				position1812 := position
				if buffer[position] != rune('(') {
					goto l1811
				}
				position++
				if !_rules[rulenn]() {
					goto l1811
				}
				if buffer[position] != rune(')') {
					goto l1811
				}
				position++
				{
					add(ruleAction28, position)
				}
				add(rulenn_contents, position1812)
			}
			return true
		l1811:
			position, tokenIndex = position1811, tokenIndex1811
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
		/* 76 EDSimple <- <(Retn / Reti / Rrd / Im0 / Im1 / Im2 / ((&('I' | 'O' | 'i' | 'o') BlitIO) | (&('R' | 'r') Rld) | (&('N' | 'n') Neg) | (&('C' | 'L' | 'c' | 'l') Blit)))> */
		nil,
		/* 77 Neg <- <(<(('n' / 'N') ('e' / 'E') ('g' / 'G'))> Action61)> */
		nil,
		/* 78 Retn <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N'))> Action62)> */
		nil,
		/* 79 Reti <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('i' / 'I'))> Action63)> */
		nil,
		/* 80 Rrd <- <(<(('r' / 'R') ('r' / 'R') ('d' / 'D'))> Action64)> */
		nil,
		/* 81 Rld <- <(<(('r' / 'R') ('l' / 'L') ('d' / 'D'))> Action65)> */
		nil,
		/* 82 Im0 <- <(<(('i' / 'I') ('m' / 'M') ' ' '0')> Action66)> */
		nil,
		/* 83 Im1 <- <(<(('i' / 'I') ('m' / 'M') ' ' '1')> Action67)> */
		nil,
		/* 84 Im2 <- <(<(('i' / 'I') ('m' / 'M') ' ' '2')> Action68)> */
		nil,
		/* 85 Blit <- <(Ldir / Ldi / Cpir / Cpi / Lddr / Ldd / Cpdr / Cpd)> */
		nil,
		/* 86 BlitIO <- <(Inir / Ini / Otir / Outi / Indr / Ind / Otdr / Outd)> */
		nil,
		/* 87 Ldi <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I'))> Action69)> */
		nil,
		/* 88 Cpi <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I'))> Action70)> */
		nil,
		/* 89 Ini <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I'))> Action71)> */
		nil,
		/* 90 Outi <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('i' / 'I'))> Action72)> */
		nil,
		/* 91 Ldd <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D'))> Action73)> */
		nil,
		/* 92 Cpd <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D'))> Action74)> */
		nil,
		/* 93 Ind <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D'))> Action75)> */
		nil,
		/* 94 Outd <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('d' / 'D'))> Action76)> */
		nil,
		/* 95 Ldir <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I') ('r' / 'R'))> Action77)> */
		nil,
		/* 96 Cpir <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I') ('r' / 'R'))> Action78)> */
		nil,
		/* 97 Inir <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I') ('r' / 'R'))> Action79)> */
		nil,
		/* 98 Otir <- <(<(('o' / 'O') ('t' / 'T') ('i' / 'I') ('r' / 'R'))> Action80)> */
		nil,
		/* 99 Lddr <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D') ('r' / 'R'))> Action81)> */
		nil,
		/* 100 Cpdr <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D') ('r' / 'R'))> Action82)> */
		nil,
		/* 101 Indr <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D') ('r' / 'R'))> Action83)> */
		nil,
		/* 102 Otdr <- <(<(('o' / 'O') ('t' / 'T') ('d' / 'D') ('r' / 'R'))> Action84)> */
		nil,
		/* 103 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 104 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action85)> */
		nil,
		/* 105 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action86)> */
		nil,
		/* 106 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action87)> */
		nil,
		/* 107 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action88)> */
		nil,
		/* 108 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action89)> */
		nil,
		/* 109 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action90)> */
		nil,
		/* 110 disp <- <signedDecimalByte> */
		func() bool {
			position1884, tokenIndex1884 := position, tokenIndex
			{
				position1885 := position
				if !_rules[rulesignedDecimalByte]() {
					goto l1884
				}
				add(ruledisp, position1885)
			}
			return true
		l1884:
			position, tokenIndex = position1884, tokenIndex1884
			return false
		},
		/* 111 IO <- <(IN / OUT)> */
		nil,
		/* 112 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action91)> */
		nil,
		/* 113 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action92)> */
		nil,
		/* 114 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position1889, tokenIndex1889 := position, tokenIndex
			{
				position1890 := position
				{
					position1891, tokenIndex1891 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l1892
					}
					position++
					{
						position1893, tokenIndex1893 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l1894
						}
						position++
						goto l1893
					l1894:
						position, tokenIndex = position1893, tokenIndex1893
						if buffer[position] != rune('C') {
							goto l1892
						}
						position++
					}
				l1893:
					if buffer[position] != rune(')') {
						goto l1892
					}
					position++
					goto l1891
				l1892:
					position, tokenIndex = position1891, tokenIndex1891
					if buffer[position] != rune('(') {
						goto l1889
					}
					position++
					if !_rules[rulen]() {
						goto l1889
					}
					if buffer[position] != rune(')') {
						goto l1889
					}
					position++
				}
			l1891:
				add(rulePort, position1890)
			}
			return true
		l1889:
			position, tokenIndex = position1889, tokenIndex1889
			return false
		},
		/* 115 sep <- <(ws? ',' ws?)> */
		func() bool {
			position1895, tokenIndex1895 := position, tokenIndex
			{
				position1896 := position
				{
					position1897, tokenIndex1897 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1897
					}
					goto l1898
				l1897:
					position, tokenIndex = position1897, tokenIndex1897
				}
			l1898:
				if buffer[position] != rune(',') {
					goto l1895
				}
				position++
				{
					position1899, tokenIndex1899 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1899
					}
					goto l1900
				l1899:
					position, tokenIndex = position1899, tokenIndex1899
				}
			l1900:
				add(rulesep, position1896)
			}
			return true
		l1895:
			position, tokenIndex = position1895, tokenIndex1895
			return false
		},
		/* 116 ws <- <' '+> */
		func() bool {
			position1901, tokenIndex1901 := position, tokenIndex
			{
				position1902 := position
				if buffer[position] != rune(' ') {
					goto l1901
				}
				position++
			l1903:
				{
					position1904, tokenIndex1904 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l1904
					}
					position++
					goto l1903
				l1904:
					position, tokenIndex = position1904, tokenIndex1904
				}
				add(rulews, position1902)
			}
			return true
		l1901:
			position, tokenIndex = position1901, tokenIndex1901
			return false
		},
		/* 117 A <- <('a' / 'A')> */
		nil,
		/* 118 F <- <('f' / 'F')> */
		nil,
		/* 119 B <- <('b' / 'B')> */
		nil,
		/* 120 C <- <('c' / 'C')> */
		nil,
		/* 121 D <- <('d' / 'D')> */
		nil,
		/* 122 E <- <('e' / 'E')> */
		nil,
		/* 123 H <- <('h' / 'H')> */
		nil,
		/* 124 L <- <('l' / 'L')> */
		nil,
		/* 125 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 126 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 127 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 128 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 129 I <- <('i' / 'I')> */
		nil,
		/* 130 R <- <('r' / 'R')> */
		nil,
		/* 131 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 132 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 133 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 134 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 135 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 136 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 137 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 138 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 139 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action93)> */
		nil,
		/* 140 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action94)> */
		nil,
		/* 141 decimalByte <- <(<[0-9]+> Action95)> */
		nil,
		/* 142 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action96)> */
		nil,
		/* 143 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action97)> */
		nil,
		/* 144 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position1932, tokenIndex1932 := position, tokenIndex
			{
				position1933 := position
				{
					position1934, tokenIndex1934 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1935
					}
					position++
					goto l1934
				l1935:
					position, tokenIndex = position1934, tokenIndex1934
					{
						position1936, tokenIndex1936 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l1937
						}
						position++
						goto l1936
					l1937:
						position, tokenIndex = position1936, tokenIndex1936
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l1932
						}
						position++
					}
				l1936:
				}
			l1934:
				add(rulehexdigit, position1933)
			}
			return true
		l1932:
			position, tokenIndex = position1932, tokenIndex1932
			return false
		},
		/* 145 octaldigit <- <(<[0-7]> Action98)> */
		func() bool {
			position1938, tokenIndex1938 := position, tokenIndex
			{
				position1939 := position
				{
					position1940 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l1938
					}
					position++
					add(rulePegText, position1940)
				}
				{
					add(ruleAction98, position)
				}
				add(ruleoctaldigit, position1939)
			}
			return true
		l1938:
			position, tokenIndex = position1938, tokenIndex1938
			return false
		},
		/* 146 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action99)> */
		func() bool {
			position1942, tokenIndex1942 := position, tokenIndex
			{
				position1943 := position
				{
					position1944 := position
					{
						position1945, tokenIndex1945 := position, tokenIndex
						{
							position1947, tokenIndex1947 := position, tokenIndex
							if buffer[position] != rune('-') {
								goto l1948
							}
							position++
							goto l1947
						l1948:
							position, tokenIndex = position1947, tokenIndex1947
							if buffer[position] != rune('+') {
								goto l1945
							}
							position++
						}
					l1947:
						goto l1946
					l1945:
						position, tokenIndex = position1945, tokenIndex1945
					}
				l1946:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1942
					}
					position++
				l1949:
					{
						position1950, tokenIndex1950 := position, tokenIndex
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1950
						}
						position++
						goto l1949
					l1950:
						position, tokenIndex = position1950, tokenIndex1950
					}
					add(rulePegText, position1944)
				}
				{
					add(ruleAction99, position)
				}
				add(rulesignedDecimalByte, position1943)
			}
			return true
		l1942:
			position, tokenIndex = position1942, tokenIndex1942
			return false
		},
		/* 147 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position1952, tokenIndex1952 := position, tokenIndex
			{
				position1953 := position
				{
					position1954, tokenIndex1954 := position, tokenIndex
					{
						position1956 := position
						{
							position1957, tokenIndex1957 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l1958
							}
							position++
							goto l1957
						l1958:
							position, tokenIndex = position1957, tokenIndex1957
							if buffer[position] != rune('N') {
								goto l1955
							}
							position++
						}
					l1957:
						{
							position1959, tokenIndex1959 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l1960
							}
							position++
							goto l1959
						l1960:
							position, tokenIndex = position1959, tokenIndex1959
							if buffer[position] != rune('Z') {
								goto l1955
							}
							position++
						}
					l1959:
						{
							add(ruleAction100, position)
						}
						add(ruleFT_NZ, position1956)
					}
					goto l1954
				l1955:
					position, tokenIndex = position1954, tokenIndex1954
					{
						position1963 := position
						{
							position1964, tokenIndex1964 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1965
							}
							position++
							goto l1964
						l1965:
							position, tokenIndex = position1964, tokenIndex1964
							if buffer[position] != rune('P') {
								goto l1962
							}
							position++
						}
					l1964:
						{
							position1966, tokenIndex1966 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l1967
							}
							position++
							goto l1966
						l1967:
							position, tokenIndex = position1966, tokenIndex1966
							if buffer[position] != rune('O') {
								goto l1962
							}
							position++
						}
					l1966:
						{
							add(ruleAction104, position)
						}
						add(ruleFT_PO, position1963)
					}
					goto l1954
				l1962:
					position, tokenIndex = position1954, tokenIndex1954
					{
						position1970 := position
						{
							position1971, tokenIndex1971 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l1972
							}
							position++
							goto l1971
						l1972:
							position, tokenIndex = position1971, tokenIndex1971
							if buffer[position] != rune('P') {
								goto l1969
							}
							position++
						}
					l1971:
						{
							position1973, tokenIndex1973 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l1974
							}
							position++
							goto l1973
						l1974:
							position, tokenIndex = position1973, tokenIndex1973
							if buffer[position] != rune('E') {
								goto l1969
							}
							position++
						}
					l1973:
						{
							add(ruleAction105, position)
						}
						add(ruleFT_PE, position1970)
					}
					goto l1954
				l1969:
					position, tokenIndex = position1954, tokenIndex1954
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position1977 := position
								{
									position1978, tokenIndex1978 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l1979
									}
									position++
									goto l1978
								l1979:
									position, tokenIndex = position1978, tokenIndex1978
									if buffer[position] != rune('M') {
										goto l1952
									}
									position++
								}
							l1978:
								{
									add(ruleAction107, position)
								}
								add(ruleFT_M, position1977)
							}
							break
						case 'P', 'p':
							{
								position1981 := position
								{
									position1982, tokenIndex1982 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l1983
									}
									position++
									goto l1982
								l1983:
									position, tokenIndex = position1982, tokenIndex1982
									if buffer[position] != rune('P') {
										goto l1952
									}
									position++
								}
							l1982:
								{
									add(ruleAction106, position)
								}
								add(ruleFT_P, position1981)
							}
							break
						case 'C', 'c':
							{
								position1985 := position
								{
									position1986, tokenIndex1986 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1987
									}
									position++
									goto l1986
								l1987:
									position, tokenIndex = position1986, tokenIndex1986
									if buffer[position] != rune('C') {
										goto l1952
									}
									position++
								}
							l1986:
								{
									add(ruleAction103, position)
								}
								add(ruleFT_C, position1985)
							}
							break
						case 'N', 'n':
							{
								position1989 := position
								{
									position1990, tokenIndex1990 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l1991
									}
									position++
									goto l1990
								l1991:
									position, tokenIndex = position1990, tokenIndex1990
									if buffer[position] != rune('N') {
										goto l1952
									}
									position++
								}
							l1990:
								{
									position1992, tokenIndex1992 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1993
									}
									position++
									goto l1992
								l1993:
									position, tokenIndex = position1992, tokenIndex1992
									if buffer[position] != rune('C') {
										goto l1952
									}
									position++
								}
							l1992:
								{
									add(ruleAction102, position)
								}
								add(ruleFT_NC, position1989)
							}
							break
						default:
							{
								position1995 := position
								{
									position1996, tokenIndex1996 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l1997
									}
									position++
									goto l1996
								l1997:
									position, tokenIndex = position1996, tokenIndex1996
									if buffer[position] != rune('Z') {
										goto l1952
									}
									position++
								}
							l1996:
								{
									add(ruleAction101, position)
								}
								add(ruleFT_Z, position1995)
							}
							break
						}
					}

				}
			l1954:
				add(rulecc, position1953)
			}
			return true
		l1952:
			position, tokenIndex = position1952, tokenIndex1952
			return false
		},
		/* 148 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action100)> */
		nil,
		/* 149 FT_Z <- <(('z' / 'Z') Action101)> */
		nil,
		/* 150 FT_NC <- <(('n' / 'N') ('c' / 'C') Action102)> */
		nil,
		/* 151 FT_C <- <(('c' / 'C') Action103)> */
		nil,
		/* 152 FT_PO <- <(('p' / 'P') ('o' / 'O') Action104)> */
		nil,
		/* 153 FT_PE <- <(('p' / 'P') ('e' / 'E') Action105)> */
		nil,
		/* 154 FT_P <- <(('p' / 'P') Action106)> */
		nil,
		/* 155 FT_M <- <(('m' / 'M') Action107)> */
		nil,
		/* 157 Action0 <- <{ p.Emit() }> */
		nil,
		/* 158 Action1 <- <{ p.LD8() }> */
		nil,
		/* 159 Action2 <- <{ p.LD16() }> */
		nil,
		/* 160 Action3 <- <{ p.Push() }> */
		nil,
		/* 161 Action4 <- <{ p.Pop() }> */
		nil,
		/* 162 Action5 <- <{ p.Ex() }> */
		nil,
		/* 163 Action6 <- <{ p.Inc8() }> */
		nil,
		/* 164 Action7 <- <{ p.Inc8() }> */
		nil,
		/* 165 Action8 <- <{ p.Inc16() }> */
		nil,
		/* 166 Action9 <- <{ p.Dec8() }> */
		nil,
		/* 167 Action10 <- <{ p.Dec8() }> */
		nil,
		/* 168 Action11 <- <{ p.Dec16() }> */
		nil,
		/* 169 Action12 <- <{ p.Add16() }> */
		nil,
		/* 170 Action13 <- <{ p.Adc16() }> */
		nil,
		/* 171 Action14 <- <{ p.Sbc16() }> */
		nil,
		/* 172 Action15 <- <{ p.Dst8() }> */
		nil,
		/* 173 Action16 <- <{ p.Src8() }> */
		nil,
		/* 174 Action17 <- <{ p.Loc8() }> */
		nil,
		/* 175 Action18 <- <{ p.Loc8() }> */
		nil,
		nil,
		/* 177 Action19 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 178 Action20 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 179 Action21 <- <{ p.Dst16() }> */
		nil,
		/* 180 Action22 <- <{ p.Src16() }> */
		nil,
		/* 181 Action23 <- <{ p.Loc16() }> */
		nil,
		/* 182 Action24 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 183 Action25 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 184 Action26 <- <{ p.R16Contents() }> */
		nil,
		/* 185 Action27 <- <{ p.IR16Contents() }> */
		nil,
		/* 186 Action28 <- <{ p.NNContents() }> */
		nil,
		/* 187 Action29 <- <{ p.Accum("ADD") }> */
		nil,
		/* 188 Action30 <- <{ p.Accum("ADC") }> */
		nil,
		/* 189 Action31 <- <{ p.Accum("SUB") }> */
		nil,
		/* 190 Action32 <- <{ p.Accum("SBC") }> */
		nil,
		/* 191 Action33 <- <{ p.Accum("AND") }> */
		nil,
		/* 192 Action34 <- <{ p.Accum("XOR") }> */
		nil,
		/* 193 Action35 <- <{ p.Accum("OR") }> */
		nil,
		/* 194 Action36 <- <{ p.Accum("CP") }> */
		nil,
		/* 195 Action37 <- <{ p.Rot("RLC") }> */
		nil,
		/* 196 Action38 <- <{ p.Rot("RRC") }> */
		nil,
		/* 197 Action39 <- <{ p.Rot("RL") }> */
		nil,
		/* 198 Action40 <- <{ p.Rot("RR") }> */
		nil,
		/* 199 Action41 <- <{ p.Rot("SLA") }> */
		nil,
		/* 200 Action42 <- <{ p.Rot("SRA") }> */
		nil,
		/* 201 Action43 <- <{ p.Rot("SLL") }> */
		nil,
		/* 202 Action44 <- <{ p.Rot("SRL") }> */
		nil,
		/* 203 Action45 <- <{ p.Bit() }> */
		nil,
		/* 204 Action46 <- <{ p.Res() }> */
		nil,
		/* 205 Action47 <- <{ p.Set() }> */
		nil,
		/* 206 Action48 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 207 Action49 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 208 Action50 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 209 Action51 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 210 Action52 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 211 Action53 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 212 Action54 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 213 Action55 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 214 Action56 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 215 Action57 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 216 Action58 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 217 Action59 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 218 Action60 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 219 Action61 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 220 Action62 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 221 Action63 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 222 Action64 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 223 Action65 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 224 Action66 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 225 Action67 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 226 Action68 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 227 Action69 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 228 Action70 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 229 Action71 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 230 Action72 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 231 Action73 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 232 Action74 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 233 Action75 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 234 Action76 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 235 Action77 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 236 Action78 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 237 Action79 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 238 Action80 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 239 Action81 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 240 Action82 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 241 Action83 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 242 Action84 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 243 Action85 <- <{ p.Rst() }> */
		nil,
		/* 244 Action86 <- <{ p.Call() }> */
		nil,
		/* 245 Action87 <- <{ p.Ret() }> */
		nil,
		/* 246 Action88 <- <{ p.Jp() }> */
		nil,
		/* 247 Action89 <- <{ p.Jr() }> */
		nil,
		/* 248 Action90 <- <{ p.Djnz() }> */
		nil,
		/* 249 Action91 <- <{ p.In() }> */
		nil,
		/* 250 Action92 <- <{ p.Out() }> */
		nil,
		/* 251 Action93 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 252 Action94 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 253 Action95 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 254 Action96 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 255 Action97 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 256 Action98 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 257 Action99 <- <{ p.SignedDecimalByte(buffer[begin:end]) }> */
		nil,
		/* 258 Action100 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 259 Action101 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 260 Action102 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 261 Action103 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 262 Action104 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 263 Action105 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 264 Action106 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 265 Action107 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
