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
	ruleComment
	ruleLineEnd
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
	ruleCopy8
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
	ruledisp
	rulesignedDecimalByte
	rulesignedHexByteH
	rulesignedHexByte0x
	rulehexByteH
	rulehexByte0x
	ruledecimalByte
	rulehexWordH
	rulehexWord0x
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
	rulehexdigit
	ruleoctaldigit
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
	ruleAction19
	rulePegText
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
	ruleAction108
	ruleAction109
	ruleAction110
)

var rul3s = [...]string{
	"Unknown",
	"Program",
	"BlankLine",
	"Line",
	"Comment",
	"LineEnd",
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
	"Copy8",
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
	"disp",
	"signedDecimalByte",
	"signedHexByteH",
	"signedHexByte0x",
	"hexByteH",
	"hexByte0x",
	"decimalByte",
	"hexWordH",
	"hexWord0x",
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
	"hexdigit",
	"octaldigit",
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
	"Action19",
	"PegText",
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
	"Action108",
	"Action109",
	"Action110",
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
	rules  [274]func() bool
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
			p.Copy8()
		case ruleAction19:
			p.Loc8()
		case ruleAction20:
			p.R8(buffer[begin:end])
		case ruleAction21:
			p.R8(buffer[begin:end])
		case ruleAction22:
			p.Dst16()
		case ruleAction23:
			p.Src16()
		case ruleAction24:
			p.Loc16()
		case ruleAction25:
			p.R16(buffer[begin:end])
		case ruleAction26:
			p.R16(buffer[begin:end])
		case ruleAction27:
			p.R16Contents()
		case ruleAction28:
			p.IR16Contents()
		case ruleAction29:
			p.DispDecimal(buffer[begin:end])
		case ruleAction30:
			p.DispHex(buffer[begin:end])
		case ruleAction31:
			p.Disp0xHex(buffer[begin:end])
		case ruleAction32:
			p.Nhex(buffer[begin:end])
		case ruleAction33:
			p.Nhex(buffer[begin:end])
		case ruleAction34:
			p.Ndec(buffer[begin:end])
		case ruleAction35:
			p.NNhex(buffer[begin:end])
		case ruleAction36:
			p.NNhex(buffer[begin:end])
		case ruleAction37:
			p.NNContents()
		case ruleAction38:
			p.Accum("ADD")
		case ruleAction39:
			p.Accum("ADC")
		case ruleAction40:
			p.Accum("SUB")
		case ruleAction41:
			p.Accum("SBC")
		case ruleAction42:
			p.Accum("AND")
		case ruleAction43:
			p.Accum("XOR")
		case ruleAction44:
			p.Accum("OR")
		case ruleAction45:
			p.Accum("CP")
		case ruleAction46:
			p.Rot("RLC")
		case ruleAction47:
			p.Rot("RRC")
		case ruleAction48:
			p.Rot("RL")
		case ruleAction49:
			p.Rot("RR")
		case ruleAction50:
			p.Rot("SLA")
		case ruleAction51:
			p.Rot("SRA")
		case ruleAction52:
			p.Rot("SLL")
		case ruleAction53:
			p.Rot("SRL")
		case ruleAction54:
			p.Bit()
		case ruleAction55:
			p.Res()
		case ruleAction56:
			p.Set()
		case ruleAction57:
			p.Simple(buffer[begin:end])
		case ruleAction58:
			p.Simple(buffer[begin:end])
		case ruleAction59:
			p.Simple(buffer[begin:end])
		case ruleAction60:
			p.Simple(buffer[begin:end])
		case ruleAction61:
			p.Simple(buffer[begin:end])
		case ruleAction62:
			p.Simple(buffer[begin:end])
		case ruleAction63:
			p.Simple(buffer[begin:end])
		case ruleAction64:
			p.Simple(buffer[begin:end])
		case ruleAction65:
			p.Simple(buffer[begin:end])
		case ruleAction66:
			p.Simple(buffer[begin:end])
		case ruleAction67:
			p.Simple(buffer[begin:end])
		case ruleAction68:
			p.Simple(buffer[begin:end])
		case ruleAction69:
			p.Simple(buffer[begin:end])
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
			p.EDSimple(buffer[begin:end])
		case ruleAction86:
			p.EDSimple(buffer[begin:end])
		case ruleAction87:
			p.EDSimple(buffer[begin:end])
		case ruleAction88:
			p.EDSimple(buffer[begin:end])
		case ruleAction89:
			p.EDSimple(buffer[begin:end])
		case ruleAction90:
			p.EDSimple(buffer[begin:end])
		case ruleAction91:
			p.EDSimple(buffer[begin:end])
		case ruleAction92:
			p.EDSimple(buffer[begin:end])
		case ruleAction93:
			p.EDSimple(buffer[begin:end])
		case ruleAction94:
			p.Rst()
		case ruleAction95:
			p.Call()
		case ruleAction96:
			p.Ret()
		case ruleAction97:
			p.Jp()
		case ruleAction98:
			p.Jr()
		case ruleAction99:
			p.Djnz()
		case ruleAction100:
			p.In()
		case ruleAction101:
			p.Out()
		case ruleAction102:
			p.ODigit(buffer[begin:end])
		case ruleAction103:
			p.Conditional(Not{FT_Z})
		case ruleAction104:
			p.Conditional(FT_Z)
		case ruleAction105:
			p.Conditional(Not{FT_C})
		case ruleAction106:
			p.Conditional(FT_C)
		case ruleAction107:
			p.Conditional(FT_PO)
		case ruleAction108:
			p.Conditional(FT_PE)
		case ruleAction109:
			p.Conditional(FT_P)
		case ruleAction110:
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

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
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
						if !_rules[ruleLineEnd]() {
							goto l5
						}
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
												add(ruleAction38, position)
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
												add(ruleAction39, position)
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
												add(ruleAction40, position)
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
														add(ruleAction45, position)
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
														add(ruleAction44, position)
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
														add(ruleAction43, position)
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
														add(ruleAction42, position)
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
														add(ruleAction41, position)
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
														position237, tokenIndex237 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l237
														}
														if !_rules[ruleCopy8]() {
															goto l237
														}
														goto l238
													l237:
														position, tokenIndex = position237, tokenIndex237
													}
												l238:
													{
														add(ruleAction46, position)
													}
													add(ruleRlc, position230)
												}
												goto l228
											l229:
												position, tokenIndex = position228, tokenIndex228
												{
													position241 := position
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
															goto l240
														}
														position++
													}
												l242:
													{
														position244, tokenIndex244 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l245
														}
														position++
														goto l244
													l245:
														position, tokenIndex = position244, tokenIndex244
														if buffer[position] != rune('R') {
															goto l240
														}
														position++
													}
												l244:
													{
														position246, tokenIndex246 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l247
														}
														position++
														goto l246
													l247:
														position, tokenIndex = position246, tokenIndex246
														if buffer[position] != rune('C') {
															goto l240
														}
														position++
													}
												l246:
													if !_rules[rulews]() {
														goto l240
													}
													if !_rules[ruleLoc8]() {
														goto l240
													}
													{
														position248, tokenIndex248 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l248
														}
														if !_rules[ruleCopy8]() {
															goto l248
														}
														goto l249
													l248:
														position, tokenIndex = position248, tokenIndex248
													}
												l249:
													{
														add(ruleAction47, position)
													}
													add(ruleRrc, position241)
												}
												goto l228
											l240:
												position, tokenIndex = position228, tokenIndex228
												{
													position252 := position
													{
														position253, tokenIndex253 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l254
														}
														position++
														goto l253
													l254:
														position, tokenIndex = position253, tokenIndex253
														if buffer[position] != rune('R') {
															goto l251
														}
														position++
													}
												l253:
													{
														position255, tokenIndex255 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l256
														}
														position++
														goto l255
													l256:
														position, tokenIndex = position255, tokenIndex255
														if buffer[position] != rune('L') {
															goto l251
														}
														position++
													}
												l255:
													if !_rules[rulews]() {
														goto l251
													}
													if !_rules[ruleLoc8]() {
														goto l251
													}
													{
														position257, tokenIndex257 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l257
														}
														if !_rules[ruleCopy8]() {
															goto l257
														}
														goto l258
													l257:
														position, tokenIndex = position257, tokenIndex257
													}
												l258:
													{
														add(ruleAction48, position)
													}
													add(ruleRl, position252)
												}
												goto l228
											l251:
												position, tokenIndex = position228, tokenIndex228
												{
													position261 := position
													{
														position262, tokenIndex262 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l263
														}
														position++
														goto l262
													l263:
														position, tokenIndex = position262, tokenIndex262
														if buffer[position] != rune('R') {
															goto l260
														}
														position++
													}
												l262:
													{
														position264, tokenIndex264 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l265
														}
														position++
														goto l264
													l265:
														position, tokenIndex = position264, tokenIndex264
														if buffer[position] != rune('R') {
															goto l260
														}
														position++
													}
												l264:
													if !_rules[rulews]() {
														goto l260
													}
													if !_rules[ruleLoc8]() {
														goto l260
													}
													{
														position266, tokenIndex266 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l266
														}
														if !_rules[ruleCopy8]() {
															goto l266
														}
														goto l267
													l266:
														position, tokenIndex = position266, tokenIndex266
													}
												l267:
													{
														add(ruleAction49, position)
													}
													add(ruleRr, position261)
												}
												goto l228
											l260:
												position, tokenIndex = position228, tokenIndex228
												{
													position270 := position
													{
														position271, tokenIndex271 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l272
														}
														position++
														goto l271
													l272:
														position, tokenIndex = position271, tokenIndex271
														if buffer[position] != rune('S') {
															goto l269
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
															goto l269
														}
														position++
													}
												l273:
													{
														position275, tokenIndex275 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l276
														}
														position++
														goto l275
													l276:
														position, tokenIndex = position275, tokenIndex275
														if buffer[position] != rune('A') {
															goto l269
														}
														position++
													}
												l275:
													if !_rules[rulews]() {
														goto l269
													}
													if !_rules[ruleLoc8]() {
														goto l269
													}
													{
														position277, tokenIndex277 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l277
														}
														if !_rules[ruleCopy8]() {
															goto l277
														}
														goto l278
													l277:
														position, tokenIndex = position277, tokenIndex277
													}
												l278:
													{
														add(ruleAction50, position)
													}
													add(ruleSla, position270)
												}
												goto l228
											l269:
												position, tokenIndex = position228, tokenIndex228
												{
													position281 := position
													{
														position282, tokenIndex282 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l283
														}
														position++
														goto l282
													l283:
														position, tokenIndex = position282, tokenIndex282
														if buffer[position] != rune('S') {
															goto l280
														}
														position++
													}
												l282:
													{
														position284, tokenIndex284 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l285
														}
														position++
														goto l284
													l285:
														position, tokenIndex = position284, tokenIndex284
														if buffer[position] != rune('R') {
															goto l280
														}
														position++
													}
												l284:
													{
														position286, tokenIndex286 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l287
														}
														position++
														goto l286
													l287:
														position, tokenIndex = position286, tokenIndex286
														if buffer[position] != rune('A') {
															goto l280
														}
														position++
													}
												l286:
													if !_rules[rulews]() {
														goto l280
													}
													if !_rules[ruleLoc8]() {
														goto l280
													}
													{
														position288, tokenIndex288 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l288
														}
														if !_rules[ruleCopy8]() {
															goto l288
														}
														goto l289
													l288:
														position, tokenIndex = position288, tokenIndex288
													}
												l289:
													{
														add(ruleAction51, position)
													}
													add(ruleSra, position281)
												}
												goto l228
											l280:
												position, tokenIndex = position228, tokenIndex228
												{
													position292 := position
													{
														position293, tokenIndex293 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l294
														}
														position++
														goto l293
													l294:
														position, tokenIndex = position293, tokenIndex293
														if buffer[position] != rune('S') {
															goto l291
														}
														position++
													}
												l293:
													{
														position295, tokenIndex295 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l296
														}
														position++
														goto l295
													l296:
														position, tokenIndex = position295, tokenIndex295
														if buffer[position] != rune('L') {
															goto l291
														}
														position++
													}
												l295:
													{
														position297, tokenIndex297 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l298
														}
														position++
														goto l297
													l298:
														position, tokenIndex = position297, tokenIndex297
														if buffer[position] != rune('L') {
															goto l291
														}
														position++
													}
												l297:
													if !_rules[rulews]() {
														goto l291
													}
													if !_rules[ruleLoc8]() {
														goto l291
													}
													{
														position299, tokenIndex299 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l299
														}
														if !_rules[ruleCopy8]() {
															goto l299
														}
														goto l300
													l299:
														position, tokenIndex = position299, tokenIndex299
													}
												l300:
													{
														add(ruleAction52, position)
													}
													add(ruleSll, position292)
												}
												goto l228
											l291:
												position, tokenIndex = position228, tokenIndex228
												{
													position302 := position
													{
														position303, tokenIndex303 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l304
														}
														position++
														goto l303
													l304:
														position, tokenIndex = position303, tokenIndex303
														if buffer[position] != rune('S') {
															goto l226
														}
														position++
													}
												l303:
													{
														position305, tokenIndex305 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l306
														}
														position++
														goto l305
													l306:
														position, tokenIndex = position305, tokenIndex305
														if buffer[position] != rune('R') {
															goto l226
														}
														position++
													}
												l305:
													{
														position307, tokenIndex307 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l308
														}
														position++
														goto l307
													l308:
														position, tokenIndex = position307, tokenIndex307
														if buffer[position] != rune('L') {
															goto l226
														}
														position++
													}
												l307:
													if !_rules[rulews]() {
														goto l226
													}
													if !_rules[ruleLoc8]() {
														goto l226
													}
													{
														position309, tokenIndex309 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l309
														}
														if !_rules[ruleCopy8]() {
															goto l309
														}
														goto l310
													l309:
														position, tokenIndex = position309, tokenIndex309
													}
												l310:
													{
														add(ruleAction53, position)
													}
													add(ruleSrl, position302)
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
													position313 := position
													{
														position314, tokenIndex314 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l315
														}
														position++
														goto l314
													l315:
														position, tokenIndex = position314, tokenIndex314
														if buffer[position] != rune('S') {
															goto l223
														}
														position++
													}
												l314:
													{
														position316, tokenIndex316 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l317
														}
														position++
														goto l316
													l317:
														position, tokenIndex = position316, tokenIndex316
														if buffer[position] != rune('E') {
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
														position320, tokenIndex320 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l320
														}
														if !_rules[ruleCopy8]() {
															goto l320
														}
														goto l321
													l320:
														position, tokenIndex = position320, tokenIndex320
													}
												l321:
													{
														add(ruleAction56, position)
													}
													add(ruleSet, position313)
												}
												break
											case 'R', 'r':
												{
													position323 := position
													{
														position324, tokenIndex324 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l325
														}
														position++
														goto l324
													l325:
														position, tokenIndex = position324, tokenIndex324
														if buffer[position] != rune('R') {
															goto l223
														}
														position++
													}
												l324:
													{
														position326, tokenIndex326 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l327
														}
														position++
														goto l326
													l327:
														position, tokenIndex = position326, tokenIndex326
														if buffer[position] != rune('E') {
															goto l223
														}
														position++
													}
												l326:
													{
														position328, tokenIndex328 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l329
														}
														position++
														goto l328
													l329:
														position, tokenIndex = position328, tokenIndex328
														if buffer[position] != rune('S') {
															goto l223
														}
														position++
													}
												l328:
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
														position330, tokenIndex330 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l330
														}
														if !_rules[ruleCopy8]() {
															goto l330
														}
														goto l331
													l330:
														position, tokenIndex = position330, tokenIndex330
													}
												l331:
													{
														add(ruleAction55, position)
													}
													add(ruleRes, position323)
												}
												break
											default:
												{
													position333 := position
													{
														position334, tokenIndex334 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l335
														}
														position++
														goto l334
													l335:
														position, tokenIndex = position334, tokenIndex334
														if buffer[position] != rune('B') {
															goto l223
														}
														position++
													}
												l334:
													{
														position336, tokenIndex336 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l337
														}
														position++
														goto l336
													l337:
														position, tokenIndex = position336, tokenIndex336
														if buffer[position] != rune('I') {
															goto l223
														}
														position++
													}
												l336:
													{
														position338, tokenIndex338 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l339
														}
														position++
														goto l338
													l339:
														position, tokenIndex = position338, tokenIndex338
														if buffer[position] != rune('T') {
															goto l223
														}
														position++
													}
												l338:
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
														add(ruleAction54, position)
													}
													add(ruleBit, position333)
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
									position342 := position
									{
										position343, tokenIndex343 := position, tokenIndex
										{
											position345 := position
											{
												position346 := position
												{
													position347, tokenIndex347 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l348
													}
													position++
													goto l347
												l348:
													position, tokenIndex = position347, tokenIndex347
													if buffer[position] != rune('R') {
														goto l344
													}
													position++
												}
											l347:
												{
													position349, tokenIndex349 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l350
													}
													position++
													goto l349
												l350:
													position, tokenIndex = position349, tokenIndex349
													if buffer[position] != rune('E') {
														goto l344
													}
													position++
												}
											l349:
												{
													position351, tokenIndex351 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l352
													}
													position++
													goto l351
												l352:
													position, tokenIndex = position351, tokenIndex351
													if buffer[position] != rune('T') {
														goto l344
													}
													position++
												}
											l351:
												{
													position353, tokenIndex353 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l354
													}
													position++
													goto l353
												l354:
													position, tokenIndex = position353, tokenIndex353
													if buffer[position] != rune('N') {
														goto l344
													}
													position++
												}
											l353:
												add(rulePegText, position346)
											}
											{
												add(ruleAction71, position)
											}
											add(ruleRetn, position345)
										}
										goto l343
									l344:
										position, tokenIndex = position343, tokenIndex343
										{
											position357 := position
											{
												position358 := position
												{
													position359, tokenIndex359 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l360
													}
													position++
													goto l359
												l360:
													position, tokenIndex = position359, tokenIndex359
													if buffer[position] != rune('R') {
														goto l356
													}
													position++
												}
											l359:
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
														goto l356
													}
													position++
												}
											l361:
												{
													position363, tokenIndex363 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l364
													}
													position++
													goto l363
												l364:
													position, tokenIndex = position363, tokenIndex363
													if buffer[position] != rune('T') {
														goto l356
													}
													position++
												}
											l363:
												{
													position365, tokenIndex365 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l366
													}
													position++
													goto l365
												l366:
													position, tokenIndex = position365, tokenIndex365
													if buffer[position] != rune('I') {
														goto l356
													}
													position++
												}
											l365:
												add(rulePegText, position358)
											}
											{
												add(ruleAction72, position)
											}
											add(ruleReti, position357)
										}
										goto l343
									l356:
										position, tokenIndex = position343, tokenIndex343
										{
											position369 := position
											{
												position370 := position
												{
													position371, tokenIndex371 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l372
													}
													position++
													goto l371
												l372:
													position, tokenIndex = position371, tokenIndex371
													if buffer[position] != rune('R') {
														goto l368
													}
													position++
												}
											l371:
												{
													position373, tokenIndex373 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l374
													}
													position++
													goto l373
												l374:
													position, tokenIndex = position373, tokenIndex373
													if buffer[position] != rune('R') {
														goto l368
													}
													position++
												}
											l373:
												{
													position375, tokenIndex375 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l376
													}
													position++
													goto l375
												l376:
													position, tokenIndex = position375, tokenIndex375
													if buffer[position] != rune('D') {
														goto l368
													}
													position++
												}
											l375:
												add(rulePegText, position370)
											}
											{
												add(ruleAction73, position)
											}
											add(ruleRrd, position369)
										}
										goto l343
									l368:
										position, tokenIndex = position343, tokenIndex343
										{
											position379 := position
											{
												position380 := position
												{
													position381, tokenIndex381 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l382
													}
													position++
													goto l381
												l382:
													position, tokenIndex = position381, tokenIndex381
													if buffer[position] != rune('I') {
														goto l378
													}
													position++
												}
											l381:
												{
													position383, tokenIndex383 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l384
													}
													position++
													goto l383
												l384:
													position, tokenIndex = position383, tokenIndex383
													if buffer[position] != rune('M') {
														goto l378
													}
													position++
												}
											l383:
												if buffer[position] != rune(' ') {
													goto l378
												}
												position++
												if buffer[position] != rune('0') {
													goto l378
												}
												position++
												add(rulePegText, position380)
											}
											{
												add(ruleAction75, position)
											}
											add(ruleIm0, position379)
										}
										goto l343
									l378:
										position, tokenIndex = position343, tokenIndex343
										{
											position387 := position
											{
												position388 := position
												{
													position389, tokenIndex389 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l390
													}
													position++
													goto l389
												l390:
													position, tokenIndex = position389, tokenIndex389
													if buffer[position] != rune('I') {
														goto l386
													}
													position++
												}
											l389:
												{
													position391, tokenIndex391 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l392
													}
													position++
													goto l391
												l392:
													position, tokenIndex = position391, tokenIndex391
													if buffer[position] != rune('M') {
														goto l386
													}
													position++
												}
											l391:
												if buffer[position] != rune(' ') {
													goto l386
												}
												position++
												if buffer[position] != rune('1') {
													goto l386
												}
												position++
												add(rulePegText, position388)
											}
											{
												add(ruleAction76, position)
											}
											add(ruleIm1, position387)
										}
										goto l343
									l386:
										position, tokenIndex = position343, tokenIndex343
										{
											position395 := position
											{
												position396 := position
												{
													position397, tokenIndex397 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l398
													}
													position++
													goto l397
												l398:
													position, tokenIndex = position397, tokenIndex397
													if buffer[position] != rune('I') {
														goto l394
													}
													position++
												}
											l397:
												{
													position399, tokenIndex399 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l400
													}
													position++
													goto l399
												l400:
													position, tokenIndex = position399, tokenIndex399
													if buffer[position] != rune('M') {
														goto l394
													}
													position++
												}
											l399:
												if buffer[position] != rune(' ') {
													goto l394
												}
												position++
												if buffer[position] != rune('2') {
													goto l394
												}
												position++
												add(rulePegText, position396)
											}
											{
												add(ruleAction77, position)
											}
											add(ruleIm2, position395)
										}
										goto l343
									l394:
										position, tokenIndex = position343, tokenIndex343
										{
											switch buffer[position] {
											case 'I', 'O', 'i', 'o':
												{
													position403 := position
													{
														position404, tokenIndex404 := position, tokenIndex
														{
															position406 := position
															{
																position407 := position
																{
																	position408, tokenIndex408 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l409
																	}
																	position++
																	goto l408
																l409:
																	position, tokenIndex = position408, tokenIndex408
																	if buffer[position] != rune('I') {
																		goto l405
																	}
																	position++
																}
															l408:
																{
																	position410, tokenIndex410 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l411
																	}
																	position++
																	goto l410
																l411:
																	position, tokenIndex = position410, tokenIndex410
																	if buffer[position] != rune('N') {
																		goto l405
																	}
																	position++
																}
															l410:
																{
																	position412, tokenIndex412 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l413
																	}
																	position++
																	goto l412
																l413:
																	position, tokenIndex = position412, tokenIndex412
																	if buffer[position] != rune('I') {
																		goto l405
																	}
																	position++
																}
															l412:
																{
																	position414, tokenIndex414 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l415
																	}
																	position++
																	goto l414
																l415:
																	position, tokenIndex = position414, tokenIndex414
																	if buffer[position] != rune('R') {
																		goto l405
																	}
																	position++
																}
															l414:
																add(rulePegText, position407)
															}
															{
																add(ruleAction88, position)
															}
															add(ruleInir, position406)
														}
														goto l404
													l405:
														position, tokenIndex = position404, tokenIndex404
														{
															position418 := position
															{
																position419 := position
																{
																	position420, tokenIndex420 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l421
																	}
																	position++
																	goto l420
																l421:
																	position, tokenIndex = position420, tokenIndex420
																	if buffer[position] != rune('I') {
																		goto l417
																	}
																	position++
																}
															l420:
																{
																	position422, tokenIndex422 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l423
																	}
																	position++
																	goto l422
																l423:
																	position, tokenIndex = position422, tokenIndex422
																	if buffer[position] != rune('N') {
																		goto l417
																	}
																	position++
																}
															l422:
																{
																	position424, tokenIndex424 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l425
																	}
																	position++
																	goto l424
																l425:
																	position, tokenIndex = position424, tokenIndex424
																	if buffer[position] != rune('I') {
																		goto l417
																	}
																	position++
																}
															l424:
																add(rulePegText, position419)
															}
															{
																add(ruleAction80, position)
															}
															add(ruleIni, position418)
														}
														goto l404
													l417:
														position, tokenIndex = position404, tokenIndex404
														{
															position428 := position
															{
																position429 := position
																{
																	position430, tokenIndex430 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l431
																	}
																	position++
																	goto l430
																l431:
																	position, tokenIndex = position430, tokenIndex430
																	if buffer[position] != rune('O') {
																		goto l427
																	}
																	position++
																}
															l430:
																{
																	position432, tokenIndex432 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l433
																	}
																	position++
																	goto l432
																l433:
																	position, tokenIndex = position432, tokenIndex432
																	if buffer[position] != rune('T') {
																		goto l427
																	}
																	position++
																}
															l432:
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
																		goto l427
																	}
																	position++
																}
															l434:
																{
																	position436, tokenIndex436 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l437
																	}
																	position++
																	goto l436
																l437:
																	position, tokenIndex = position436, tokenIndex436
																	if buffer[position] != rune('R') {
																		goto l427
																	}
																	position++
																}
															l436:
																add(rulePegText, position429)
															}
															{
																add(ruleAction89, position)
															}
															add(ruleOtir, position428)
														}
														goto l404
													l427:
														position, tokenIndex = position404, tokenIndex404
														{
															position440 := position
															{
																position441 := position
																{
																	position442, tokenIndex442 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l443
																	}
																	position++
																	goto l442
																l443:
																	position, tokenIndex = position442, tokenIndex442
																	if buffer[position] != rune('O') {
																		goto l439
																	}
																	position++
																}
															l442:
																{
																	position444, tokenIndex444 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l445
																	}
																	position++
																	goto l444
																l445:
																	position, tokenIndex = position444, tokenIndex444
																	if buffer[position] != rune('U') {
																		goto l439
																	}
																	position++
																}
															l444:
																{
																	position446, tokenIndex446 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l447
																	}
																	position++
																	goto l446
																l447:
																	position, tokenIndex = position446, tokenIndex446
																	if buffer[position] != rune('T') {
																		goto l439
																	}
																	position++
																}
															l446:
																{
																	position448, tokenIndex448 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l449
																	}
																	position++
																	goto l448
																l449:
																	position, tokenIndex = position448, tokenIndex448
																	if buffer[position] != rune('I') {
																		goto l439
																	}
																	position++
																}
															l448:
																add(rulePegText, position441)
															}
															{
																add(ruleAction81, position)
															}
															add(ruleOuti, position440)
														}
														goto l404
													l439:
														position, tokenIndex = position404, tokenIndex404
														{
															position452 := position
															{
																position453 := position
																{
																	position454, tokenIndex454 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l455
																	}
																	position++
																	goto l454
																l455:
																	position, tokenIndex = position454, tokenIndex454
																	if buffer[position] != rune('I') {
																		goto l451
																	}
																	position++
																}
															l454:
																{
																	position456, tokenIndex456 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l457
																	}
																	position++
																	goto l456
																l457:
																	position, tokenIndex = position456, tokenIndex456
																	if buffer[position] != rune('N') {
																		goto l451
																	}
																	position++
																}
															l456:
																{
																	position458, tokenIndex458 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l459
																	}
																	position++
																	goto l458
																l459:
																	position, tokenIndex = position458, tokenIndex458
																	if buffer[position] != rune('D') {
																		goto l451
																	}
																	position++
																}
															l458:
																{
																	position460, tokenIndex460 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l461
																	}
																	position++
																	goto l460
																l461:
																	position, tokenIndex = position460, tokenIndex460
																	if buffer[position] != rune('R') {
																		goto l451
																	}
																	position++
																}
															l460:
																add(rulePegText, position453)
															}
															{
																add(ruleAction92, position)
															}
															add(ruleIndr, position452)
														}
														goto l404
													l451:
														position, tokenIndex = position404, tokenIndex404
														{
															position464 := position
															{
																position465 := position
																{
																	position466, tokenIndex466 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l467
																	}
																	position++
																	goto l466
																l467:
																	position, tokenIndex = position466, tokenIndex466
																	if buffer[position] != rune('I') {
																		goto l463
																	}
																	position++
																}
															l466:
																{
																	position468, tokenIndex468 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l469
																	}
																	position++
																	goto l468
																l469:
																	position, tokenIndex = position468, tokenIndex468
																	if buffer[position] != rune('N') {
																		goto l463
																	}
																	position++
																}
															l468:
																{
																	position470, tokenIndex470 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l471
																	}
																	position++
																	goto l470
																l471:
																	position, tokenIndex = position470, tokenIndex470
																	if buffer[position] != rune('D') {
																		goto l463
																	}
																	position++
																}
															l470:
																add(rulePegText, position465)
															}
															{
																add(ruleAction84, position)
															}
															add(ruleInd, position464)
														}
														goto l404
													l463:
														position, tokenIndex = position404, tokenIndex404
														{
															position474 := position
															{
																position475 := position
																{
																	position476, tokenIndex476 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l477
																	}
																	position++
																	goto l476
																l477:
																	position, tokenIndex = position476, tokenIndex476
																	if buffer[position] != rune('O') {
																		goto l473
																	}
																	position++
																}
															l476:
																{
																	position478, tokenIndex478 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l479
																	}
																	position++
																	goto l478
																l479:
																	position, tokenIndex = position478, tokenIndex478
																	if buffer[position] != rune('T') {
																		goto l473
																	}
																	position++
																}
															l478:
																{
																	position480, tokenIndex480 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l481
																	}
																	position++
																	goto l480
																l481:
																	position, tokenIndex = position480, tokenIndex480
																	if buffer[position] != rune('D') {
																		goto l473
																	}
																	position++
																}
															l480:
																{
																	position482, tokenIndex482 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l483
																	}
																	position++
																	goto l482
																l483:
																	position, tokenIndex = position482, tokenIndex482
																	if buffer[position] != rune('R') {
																		goto l473
																	}
																	position++
																}
															l482:
																add(rulePegText, position475)
															}
															{
																add(ruleAction93, position)
															}
															add(ruleOtdr, position474)
														}
														goto l404
													l473:
														position, tokenIndex = position404, tokenIndex404
														{
															position485 := position
															{
																position486 := position
																{
																	position487, tokenIndex487 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l488
																	}
																	position++
																	goto l487
																l488:
																	position, tokenIndex = position487, tokenIndex487
																	if buffer[position] != rune('O') {
																		goto l341
																	}
																	position++
																}
															l487:
																{
																	position489, tokenIndex489 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l490
																	}
																	position++
																	goto l489
																l490:
																	position, tokenIndex = position489, tokenIndex489
																	if buffer[position] != rune('U') {
																		goto l341
																	}
																	position++
																}
															l489:
																{
																	position491, tokenIndex491 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l492
																	}
																	position++
																	goto l491
																l492:
																	position, tokenIndex = position491, tokenIndex491
																	if buffer[position] != rune('T') {
																		goto l341
																	}
																	position++
																}
															l491:
																{
																	position493, tokenIndex493 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l494
																	}
																	position++
																	goto l493
																l494:
																	position, tokenIndex = position493, tokenIndex493
																	if buffer[position] != rune('D') {
																		goto l341
																	}
																	position++
																}
															l493:
																add(rulePegText, position486)
															}
															{
																add(ruleAction85, position)
															}
															add(ruleOutd, position485)
														}
													}
												l404:
													add(ruleBlitIO, position403)
												}
												break
											case 'R', 'r':
												{
													position496 := position
													{
														position497 := position
														{
															position498, tokenIndex498 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l499
															}
															position++
															goto l498
														l499:
															position, tokenIndex = position498, tokenIndex498
															if buffer[position] != rune('R') {
																goto l341
															}
															position++
														}
													l498:
														{
															position500, tokenIndex500 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l501
															}
															position++
															goto l500
														l501:
															position, tokenIndex = position500, tokenIndex500
															if buffer[position] != rune('L') {
																goto l341
															}
															position++
														}
													l500:
														{
															position502, tokenIndex502 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l503
															}
															position++
															goto l502
														l503:
															position, tokenIndex = position502, tokenIndex502
															if buffer[position] != rune('D') {
																goto l341
															}
															position++
														}
													l502:
														add(rulePegText, position497)
													}
													{
														add(ruleAction74, position)
													}
													add(ruleRld, position496)
												}
												break
											case 'N', 'n':
												{
													position505 := position
													{
														position506 := position
														{
															position507, tokenIndex507 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l508
															}
															position++
															goto l507
														l508:
															position, tokenIndex = position507, tokenIndex507
															if buffer[position] != rune('N') {
																goto l341
															}
															position++
														}
													l507:
														{
															position509, tokenIndex509 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l510
															}
															position++
															goto l509
														l510:
															position, tokenIndex = position509, tokenIndex509
															if buffer[position] != rune('E') {
																goto l341
															}
															position++
														}
													l509:
														{
															position511, tokenIndex511 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l512
															}
															position++
															goto l511
														l512:
															position, tokenIndex = position511, tokenIndex511
															if buffer[position] != rune('G') {
																goto l341
															}
															position++
														}
													l511:
														add(rulePegText, position506)
													}
													{
														add(ruleAction70, position)
													}
													add(ruleNeg, position505)
												}
												break
											default:
												{
													position514 := position
													{
														position515, tokenIndex515 := position, tokenIndex
														{
															position517 := position
															{
																position518 := position
																{
																	position519, tokenIndex519 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l520
																	}
																	position++
																	goto l519
																l520:
																	position, tokenIndex = position519, tokenIndex519
																	if buffer[position] != rune('L') {
																		goto l516
																	}
																	position++
																}
															l519:
																{
																	position521, tokenIndex521 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l522
																	}
																	position++
																	goto l521
																l522:
																	position, tokenIndex = position521, tokenIndex521
																	if buffer[position] != rune('D') {
																		goto l516
																	}
																	position++
																}
															l521:
																{
																	position523, tokenIndex523 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l524
																	}
																	position++
																	goto l523
																l524:
																	position, tokenIndex = position523, tokenIndex523
																	if buffer[position] != rune('I') {
																		goto l516
																	}
																	position++
																}
															l523:
																{
																	position525, tokenIndex525 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l526
																	}
																	position++
																	goto l525
																l526:
																	position, tokenIndex = position525, tokenIndex525
																	if buffer[position] != rune('R') {
																		goto l516
																	}
																	position++
																}
															l525:
																add(rulePegText, position518)
															}
															{
																add(ruleAction86, position)
															}
															add(ruleLdir, position517)
														}
														goto l515
													l516:
														position, tokenIndex = position515, tokenIndex515
														{
															position529 := position
															{
																position530 := position
																{
																	position531, tokenIndex531 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l532
																	}
																	position++
																	goto l531
																l532:
																	position, tokenIndex = position531, tokenIndex531
																	if buffer[position] != rune('L') {
																		goto l528
																	}
																	position++
																}
															l531:
																{
																	position533, tokenIndex533 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l534
																	}
																	position++
																	goto l533
																l534:
																	position, tokenIndex = position533, tokenIndex533
																	if buffer[position] != rune('D') {
																		goto l528
																	}
																	position++
																}
															l533:
																{
																	position535, tokenIndex535 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l536
																	}
																	position++
																	goto l535
																l536:
																	position, tokenIndex = position535, tokenIndex535
																	if buffer[position] != rune('I') {
																		goto l528
																	}
																	position++
																}
															l535:
																add(rulePegText, position530)
															}
															{
																add(ruleAction78, position)
															}
															add(ruleLdi, position529)
														}
														goto l515
													l528:
														position, tokenIndex = position515, tokenIndex515
														{
															position539 := position
															{
																position540 := position
																{
																	position541, tokenIndex541 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l542
																	}
																	position++
																	goto l541
																l542:
																	position, tokenIndex = position541, tokenIndex541
																	if buffer[position] != rune('C') {
																		goto l538
																	}
																	position++
																}
															l541:
																{
																	position543, tokenIndex543 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l544
																	}
																	position++
																	goto l543
																l544:
																	position, tokenIndex = position543, tokenIndex543
																	if buffer[position] != rune('P') {
																		goto l538
																	}
																	position++
																}
															l543:
																{
																	position545, tokenIndex545 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l546
																	}
																	position++
																	goto l545
																l546:
																	position, tokenIndex = position545, tokenIndex545
																	if buffer[position] != rune('I') {
																		goto l538
																	}
																	position++
																}
															l545:
																{
																	position547, tokenIndex547 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l548
																	}
																	position++
																	goto l547
																l548:
																	position, tokenIndex = position547, tokenIndex547
																	if buffer[position] != rune('R') {
																		goto l538
																	}
																	position++
																}
															l547:
																add(rulePegText, position540)
															}
															{
																add(ruleAction87, position)
															}
															add(ruleCpir, position539)
														}
														goto l515
													l538:
														position, tokenIndex = position515, tokenIndex515
														{
															position551 := position
															{
																position552 := position
																{
																	position553, tokenIndex553 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l554
																	}
																	position++
																	goto l553
																l554:
																	position, tokenIndex = position553, tokenIndex553
																	if buffer[position] != rune('C') {
																		goto l550
																	}
																	position++
																}
															l553:
																{
																	position555, tokenIndex555 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l556
																	}
																	position++
																	goto l555
																l556:
																	position, tokenIndex = position555, tokenIndex555
																	if buffer[position] != rune('P') {
																		goto l550
																	}
																	position++
																}
															l555:
																{
																	position557, tokenIndex557 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l558
																	}
																	position++
																	goto l557
																l558:
																	position, tokenIndex = position557, tokenIndex557
																	if buffer[position] != rune('I') {
																		goto l550
																	}
																	position++
																}
															l557:
																add(rulePegText, position552)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleCpi, position551)
														}
														goto l515
													l550:
														position, tokenIndex = position515, tokenIndex515
														{
															position561 := position
															{
																position562 := position
																{
																	position563, tokenIndex563 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l564
																	}
																	position++
																	goto l563
																l564:
																	position, tokenIndex = position563, tokenIndex563
																	if buffer[position] != rune('L') {
																		goto l560
																	}
																	position++
																}
															l563:
																{
																	position565, tokenIndex565 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l566
																	}
																	position++
																	goto l565
																l566:
																	position, tokenIndex = position565, tokenIndex565
																	if buffer[position] != rune('D') {
																		goto l560
																	}
																	position++
																}
															l565:
																{
																	position567, tokenIndex567 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l568
																	}
																	position++
																	goto l567
																l568:
																	position, tokenIndex = position567, tokenIndex567
																	if buffer[position] != rune('D') {
																		goto l560
																	}
																	position++
																}
															l567:
																{
																	position569, tokenIndex569 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l570
																	}
																	position++
																	goto l569
																l570:
																	position, tokenIndex = position569, tokenIndex569
																	if buffer[position] != rune('R') {
																		goto l560
																	}
																	position++
																}
															l569:
																add(rulePegText, position562)
															}
															{
																add(ruleAction90, position)
															}
															add(ruleLddr, position561)
														}
														goto l515
													l560:
														position, tokenIndex = position515, tokenIndex515
														{
															position573 := position
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
																		goto l572
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
																		goto l572
																	}
																	position++
																}
															l577:
																{
																	position579, tokenIndex579 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l580
																	}
																	position++
																	goto l579
																l580:
																	position, tokenIndex = position579, tokenIndex579
																	if buffer[position] != rune('D') {
																		goto l572
																	}
																	position++
																}
															l579:
																add(rulePegText, position574)
															}
															{
																add(ruleAction82, position)
															}
															add(ruleLdd, position573)
														}
														goto l515
													l572:
														position, tokenIndex = position515, tokenIndex515
														{
															position583 := position
															{
																position584 := position
																{
																	position585, tokenIndex585 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l586
																	}
																	position++
																	goto l585
																l586:
																	position, tokenIndex = position585, tokenIndex585
																	if buffer[position] != rune('C') {
																		goto l582
																	}
																	position++
																}
															l585:
																{
																	position587, tokenIndex587 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l588
																	}
																	position++
																	goto l587
																l588:
																	position, tokenIndex = position587, tokenIndex587
																	if buffer[position] != rune('P') {
																		goto l582
																	}
																	position++
																}
															l587:
																{
																	position589, tokenIndex589 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l590
																	}
																	position++
																	goto l589
																l590:
																	position, tokenIndex = position589, tokenIndex589
																	if buffer[position] != rune('D') {
																		goto l582
																	}
																	position++
																}
															l589:
																{
																	position591, tokenIndex591 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l592
																	}
																	position++
																	goto l591
																l592:
																	position, tokenIndex = position591, tokenIndex591
																	if buffer[position] != rune('R') {
																		goto l582
																	}
																	position++
																}
															l591:
																add(rulePegText, position584)
															}
															{
																add(ruleAction91, position)
															}
															add(ruleCpdr, position583)
														}
														goto l515
													l582:
														position, tokenIndex = position515, tokenIndex515
														{
															position594 := position
															{
																position595 := position
																{
																	position596, tokenIndex596 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l597
																	}
																	position++
																	goto l596
																l597:
																	position, tokenIndex = position596, tokenIndex596
																	if buffer[position] != rune('C') {
																		goto l341
																	}
																	position++
																}
															l596:
																{
																	position598, tokenIndex598 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l599
																	}
																	position++
																	goto l598
																l599:
																	position, tokenIndex = position598, tokenIndex598
																	if buffer[position] != rune('P') {
																		goto l341
																	}
																	position++
																}
															l598:
																{
																	position600, tokenIndex600 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l601
																	}
																	position++
																	goto l600
																l601:
																	position, tokenIndex = position600, tokenIndex600
																	if buffer[position] != rune('D') {
																		goto l341
																	}
																	position++
																}
															l600:
																add(rulePegText, position595)
															}
															{
																add(ruleAction83, position)
															}
															add(ruleCpd, position594)
														}
													}
												l515:
													add(ruleBlit, position514)
												}
												break
											}
										}

									}
								l343:
									add(ruleEDSimple, position342)
								}
								goto l13
							l341:
								position, tokenIndex = position13, tokenIndex13
								{
									position604 := position
									{
										position605, tokenIndex605 := position, tokenIndex
										{
											position607 := position
											{
												position608 := position
												{
													position609, tokenIndex609 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l610
													}
													position++
													goto l609
												l610:
													position, tokenIndex = position609, tokenIndex609
													if buffer[position] != rune('R') {
														goto l606
													}
													position++
												}
											l609:
												{
													position611, tokenIndex611 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l612
													}
													position++
													goto l611
												l612:
													position, tokenIndex = position611, tokenIndex611
													if buffer[position] != rune('L') {
														goto l606
													}
													position++
												}
											l611:
												{
													position613, tokenIndex613 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l614
													}
													position++
													goto l613
												l614:
													position, tokenIndex = position613, tokenIndex613
													if buffer[position] != rune('C') {
														goto l606
													}
													position++
												}
											l613:
												{
													position615, tokenIndex615 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l616
													}
													position++
													goto l615
												l616:
													position, tokenIndex = position615, tokenIndex615
													if buffer[position] != rune('A') {
														goto l606
													}
													position++
												}
											l615:
												add(rulePegText, position608)
											}
											{
												add(ruleAction59, position)
											}
											add(ruleRlca, position607)
										}
										goto l605
									l606:
										position, tokenIndex = position605, tokenIndex605
										{
											position619 := position
											{
												position620 := position
												{
													position621, tokenIndex621 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l622
													}
													position++
													goto l621
												l622:
													position, tokenIndex = position621, tokenIndex621
													if buffer[position] != rune('R') {
														goto l618
													}
													position++
												}
											l621:
												{
													position623, tokenIndex623 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l624
													}
													position++
													goto l623
												l624:
													position, tokenIndex = position623, tokenIndex623
													if buffer[position] != rune('R') {
														goto l618
													}
													position++
												}
											l623:
												{
													position625, tokenIndex625 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l626
													}
													position++
													goto l625
												l626:
													position, tokenIndex = position625, tokenIndex625
													if buffer[position] != rune('C') {
														goto l618
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
														goto l618
													}
													position++
												}
											l627:
												add(rulePegText, position620)
											}
											{
												add(ruleAction60, position)
											}
											add(ruleRrca, position619)
										}
										goto l605
									l618:
										position, tokenIndex = position605, tokenIndex605
										{
											position631 := position
											{
												position632 := position
												{
													position633, tokenIndex633 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l634
													}
													position++
													goto l633
												l634:
													position, tokenIndex = position633, tokenIndex633
													if buffer[position] != rune('R') {
														goto l630
													}
													position++
												}
											l633:
												{
													position635, tokenIndex635 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l636
													}
													position++
													goto l635
												l636:
													position, tokenIndex = position635, tokenIndex635
													if buffer[position] != rune('L') {
														goto l630
													}
													position++
												}
											l635:
												{
													position637, tokenIndex637 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l638
													}
													position++
													goto l637
												l638:
													position, tokenIndex = position637, tokenIndex637
													if buffer[position] != rune('A') {
														goto l630
													}
													position++
												}
											l637:
												add(rulePegText, position632)
											}
											{
												add(ruleAction61, position)
											}
											add(ruleRla, position631)
										}
										goto l605
									l630:
										position, tokenIndex = position605, tokenIndex605
										{
											position641 := position
											{
												position642 := position
												{
													position643, tokenIndex643 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l644
													}
													position++
													goto l643
												l644:
													position, tokenIndex = position643, tokenIndex643
													if buffer[position] != rune('D') {
														goto l640
													}
													position++
												}
											l643:
												{
													position645, tokenIndex645 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l646
													}
													position++
													goto l645
												l646:
													position, tokenIndex = position645, tokenIndex645
													if buffer[position] != rune('A') {
														goto l640
													}
													position++
												}
											l645:
												{
													position647, tokenIndex647 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l648
													}
													position++
													goto l647
												l648:
													position, tokenIndex = position647, tokenIndex647
													if buffer[position] != rune('A') {
														goto l640
													}
													position++
												}
											l647:
												add(rulePegText, position642)
											}
											{
												add(ruleAction63, position)
											}
											add(ruleDaa, position641)
										}
										goto l605
									l640:
										position, tokenIndex = position605, tokenIndex605
										{
											position651 := position
											{
												position652 := position
												{
													position653, tokenIndex653 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l654
													}
													position++
													goto l653
												l654:
													position, tokenIndex = position653, tokenIndex653
													if buffer[position] != rune('C') {
														goto l650
													}
													position++
												}
											l653:
												{
													position655, tokenIndex655 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l656
													}
													position++
													goto l655
												l656:
													position, tokenIndex = position655, tokenIndex655
													if buffer[position] != rune('P') {
														goto l650
													}
													position++
												}
											l655:
												{
													position657, tokenIndex657 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l658
													}
													position++
													goto l657
												l658:
													position, tokenIndex = position657, tokenIndex657
													if buffer[position] != rune('L') {
														goto l650
													}
													position++
												}
											l657:
												add(rulePegText, position652)
											}
											{
												add(ruleAction64, position)
											}
											add(ruleCpl, position651)
										}
										goto l605
									l650:
										position, tokenIndex = position605, tokenIndex605
										{
											position661 := position
											{
												position662 := position
												{
													position663, tokenIndex663 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l664
													}
													position++
													goto l663
												l664:
													position, tokenIndex = position663, tokenIndex663
													if buffer[position] != rune('E') {
														goto l660
													}
													position++
												}
											l663:
												{
													position665, tokenIndex665 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l666
													}
													position++
													goto l665
												l666:
													position, tokenIndex = position665, tokenIndex665
													if buffer[position] != rune('X') {
														goto l660
													}
													position++
												}
											l665:
												{
													position667, tokenIndex667 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l668
													}
													position++
													goto l667
												l668:
													position, tokenIndex = position667, tokenIndex667
													if buffer[position] != rune('X') {
														goto l660
													}
													position++
												}
											l667:
												add(rulePegText, position662)
											}
											{
												add(ruleAction67, position)
											}
											add(ruleExx, position661)
										}
										goto l605
									l660:
										position, tokenIndex = position605, tokenIndex605
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position671 := position
													{
														position672 := position
														{
															position673, tokenIndex673 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l674
															}
															position++
															goto l673
														l674:
															position, tokenIndex = position673, tokenIndex673
															if buffer[position] != rune('E') {
																goto l603
															}
															position++
														}
													l673:
														{
															position675, tokenIndex675 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l676
															}
															position++
															goto l675
														l676:
															position, tokenIndex = position675, tokenIndex675
															if buffer[position] != rune('I') {
																goto l603
															}
															position++
														}
													l675:
														add(rulePegText, position672)
													}
													{
														add(ruleAction69, position)
													}
													add(ruleEi, position671)
												}
												break
											case 'D', 'd':
												{
													position678 := position
													{
														position679 := position
														{
															position680, tokenIndex680 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l681
															}
															position++
															goto l680
														l681:
															position, tokenIndex = position680, tokenIndex680
															if buffer[position] != rune('D') {
																goto l603
															}
															position++
														}
													l680:
														{
															position682, tokenIndex682 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l683
															}
															position++
															goto l682
														l683:
															position, tokenIndex = position682, tokenIndex682
															if buffer[position] != rune('I') {
																goto l603
															}
															position++
														}
													l682:
														add(rulePegText, position679)
													}
													{
														add(ruleAction68, position)
													}
													add(ruleDi, position678)
												}
												break
											case 'C', 'c':
												{
													position685 := position
													{
														position686 := position
														{
															position687, tokenIndex687 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l688
															}
															position++
															goto l687
														l688:
															position, tokenIndex = position687, tokenIndex687
															if buffer[position] != rune('C') {
																goto l603
															}
															position++
														}
													l687:
														{
															position689, tokenIndex689 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l690
															}
															position++
															goto l689
														l690:
															position, tokenIndex = position689, tokenIndex689
															if buffer[position] != rune('C') {
																goto l603
															}
															position++
														}
													l689:
														{
															position691, tokenIndex691 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l692
															}
															position++
															goto l691
														l692:
															position, tokenIndex = position691, tokenIndex691
															if buffer[position] != rune('F') {
																goto l603
															}
															position++
														}
													l691:
														add(rulePegText, position686)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleCcf, position685)
												}
												break
											case 'S', 's':
												{
													position694 := position
													{
														position695 := position
														{
															position696, tokenIndex696 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l697
															}
															position++
															goto l696
														l697:
															position, tokenIndex = position696, tokenIndex696
															if buffer[position] != rune('S') {
																goto l603
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
																goto l603
															}
															position++
														}
													l698:
														{
															position700, tokenIndex700 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l701
															}
															position++
															goto l700
														l701:
															position, tokenIndex = position700, tokenIndex700
															if buffer[position] != rune('F') {
																goto l603
															}
															position++
														}
													l700:
														add(rulePegText, position695)
													}
													{
														add(ruleAction65, position)
													}
													add(ruleScf, position694)
												}
												break
											case 'R', 'r':
												{
													position703 := position
													{
														position704 := position
														{
															position705, tokenIndex705 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l706
															}
															position++
															goto l705
														l706:
															position, tokenIndex = position705, tokenIndex705
															if buffer[position] != rune('R') {
																goto l603
															}
															position++
														}
													l705:
														{
															position707, tokenIndex707 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l708
															}
															position++
															goto l707
														l708:
															position, tokenIndex = position707, tokenIndex707
															if buffer[position] != rune('R') {
																goto l603
															}
															position++
														}
													l707:
														{
															position709, tokenIndex709 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l710
															}
															position++
															goto l709
														l710:
															position, tokenIndex = position709, tokenIndex709
															if buffer[position] != rune('A') {
																goto l603
															}
															position++
														}
													l709:
														add(rulePegText, position704)
													}
													{
														add(ruleAction62, position)
													}
													add(ruleRra, position703)
												}
												break
											case 'H', 'h':
												{
													position712 := position
													{
														position713 := position
														{
															position714, tokenIndex714 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l715
															}
															position++
															goto l714
														l715:
															position, tokenIndex = position714, tokenIndex714
															if buffer[position] != rune('H') {
																goto l603
															}
															position++
														}
													l714:
														{
															position716, tokenIndex716 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l717
															}
															position++
															goto l716
														l717:
															position, tokenIndex = position716, tokenIndex716
															if buffer[position] != rune('A') {
																goto l603
															}
															position++
														}
													l716:
														{
															position718, tokenIndex718 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l719
															}
															position++
															goto l718
														l719:
															position, tokenIndex = position718, tokenIndex718
															if buffer[position] != rune('L') {
																goto l603
															}
															position++
														}
													l718:
														{
															position720, tokenIndex720 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l721
															}
															position++
															goto l720
														l721:
															position, tokenIndex = position720, tokenIndex720
															if buffer[position] != rune('T') {
																goto l603
															}
															position++
														}
													l720:
														add(rulePegText, position713)
													}
													{
														add(ruleAction58, position)
													}
													add(ruleHalt, position712)
												}
												break
											default:
												{
													position723 := position
													{
														position724 := position
														{
															position725, tokenIndex725 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l726
															}
															position++
															goto l725
														l726:
															position, tokenIndex = position725, tokenIndex725
															if buffer[position] != rune('N') {
																goto l603
															}
															position++
														}
													l725:
														{
															position727, tokenIndex727 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l728
															}
															position++
															goto l727
														l728:
															position, tokenIndex = position727, tokenIndex727
															if buffer[position] != rune('O') {
																goto l603
															}
															position++
														}
													l727:
														{
															position729, tokenIndex729 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l730
															}
															position++
															goto l729
														l730:
															position, tokenIndex = position729, tokenIndex729
															if buffer[position] != rune('P') {
																goto l603
															}
															position++
														}
													l729:
														add(rulePegText, position724)
													}
													{
														add(ruleAction57, position)
													}
													add(ruleNop, position723)
												}
												break
											}
										}

									}
								l605:
									add(ruleSimple, position604)
								}
								goto l13
							l603:
								position, tokenIndex = position13, tokenIndex13
								{
									position733 := position
									{
										position734, tokenIndex734 := position, tokenIndex
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
												if buffer[position] != rune('s') {
													goto l740
												}
												position++
												goto l739
											l740:
												position, tokenIndex = position739, tokenIndex739
												if buffer[position] != rune('S') {
													goto l735
												}
												position++
											}
										l739:
											{
												position741, tokenIndex741 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l742
												}
												position++
												goto l741
											l742:
												position, tokenIndex = position741, tokenIndex741
												if buffer[position] != rune('T') {
													goto l735
												}
												position++
											}
										l741:
											if !_rules[rulews]() {
												goto l735
											}
											if !_rules[rulen]() {
												goto l735
											}
											{
												add(ruleAction94, position)
											}
											add(ruleRst, position736)
										}
										goto l734
									l735:
										position, tokenIndex = position734, tokenIndex734
										{
											position745 := position
											{
												position746, tokenIndex746 := position, tokenIndex
												if buffer[position] != rune('j') {
													goto l747
												}
												position++
												goto l746
											l747:
												position, tokenIndex = position746, tokenIndex746
												if buffer[position] != rune('J') {
													goto l744
												}
												position++
											}
										l746:
											{
												position748, tokenIndex748 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l749
												}
												position++
												goto l748
											l749:
												position, tokenIndex = position748, tokenIndex748
												if buffer[position] != rune('P') {
													goto l744
												}
												position++
											}
										l748:
											if !_rules[rulews]() {
												goto l744
											}
											{
												position750, tokenIndex750 := position, tokenIndex
												if !_rules[rulecc]() {
													goto l750
												}
												if !_rules[rulesep]() {
													goto l750
												}
												goto l751
											l750:
												position, tokenIndex = position750, tokenIndex750
											}
										l751:
											if !_rules[ruleSrc16]() {
												goto l744
											}
											{
												add(ruleAction97, position)
											}
											add(ruleJp, position745)
										}
										goto l734
									l744:
										position, tokenIndex = position734, tokenIndex734
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position754 := position
													{
														position755, tokenIndex755 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l756
														}
														position++
														goto l755
													l756:
														position, tokenIndex = position755, tokenIndex755
														if buffer[position] != rune('D') {
															goto l732
														}
														position++
													}
												l755:
													{
														position757, tokenIndex757 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l758
														}
														position++
														goto l757
													l758:
														position, tokenIndex = position757, tokenIndex757
														if buffer[position] != rune('J') {
															goto l732
														}
														position++
													}
												l757:
													{
														position759, tokenIndex759 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l760
														}
														position++
														goto l759
													l760:
														position, tokenIndex = position759, tokenIndex759
														if buffer[position] != rune('N') {
															goto l732
														}
														position++
													}
												l759:
													{
														position761, tokenIndex761 := position, tokenIndex
														if buffer[position] != rune('z') {
															goto l762
														}
														position++
														goto l761
													l762:
														position, tokenIndex = position761, tokenIndex761
														if buffer[position] != rune('Z') {
															goto l732
														}
														position++
													}
												l761:
													if !_rules[rulews]() {
														goto l732
													}
													if !_rules[ruledisp]() {
														goto l732
													}
													{
														add(ruleAction99, position)
													}
													add(ruleDjnz, position754)
												}
												break
											case 'J', 'j':
												{
													position764 := position
													{
														position765, tokenIndex765 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l766
														}
														position++
														goto l765
													l766:
														position, tokenIndex = position765, tokenIndex765
														if buffer[position] != rune('J') {
															goto l732
														}
														position++
													}
												l765:
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
															goto l732
														}
														position++
													}
												l767:
													if !_rules[rulews]() {
														goto l732
													}
													{
														position769, tokenIndex769 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l769
														}
														if !_rules[rulesep]() {
															goto l769
														}
														goto l770
													l769:
														position, tokenIndex = position769, tokenIndex769
													}
												l770:
													if !_rules[ruledisp]() {
														goto l732
													}
													{
														add(ruleAction98, position)
													}
													add(ruleJr, position764)
												}
												break
											case 'R', 'r':
												{
													position772 := position
													{
														position773, tokenIndex773 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l774
														}
														position++
														goto l773
													l774:
														position, tokenIndex = position773, tokenIndex773
														if buffer[position] != rune('R') {
															goto l732
														}
														position++
													}
												l773:
													{
														position775, tokenIndex775 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l776
														}
														position++
														goto l775
													l776:
														position, tokenIndex = position775, tokenIndex775
														if buffer[position] != rune('E') {
															goto l732
														}
														position++
													}
												l775:
													{
														position777, tokenIndex777 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l778
														}
														position++
														goto l777
													l778:
														position, tokenIndex = position777, tokenIndex777
														if buffer[position] != rune('T') {
															goto l732
														}
														position++
													}
												l777:
													{
														position779, tokenIndex779 := position, tokenIndex
														if !_rules[rulews]() {
															goto l779
														}
														if !_rules[rulecc]() {
															goto l779
														}
														goto l780
													l779:
														position, tokenIndex = position779, tokenIndex779
													}
												l780:
													{
														add(ruleAction96, position)
													}
													add(ruleRet, position772)
												}
												break
											default:
												{
													position782 := position
													{
														position783, tokenIndex783 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l784
														}
														position++
														goto l783
													l784:
														position, tokenIndex = position783, tokenIndex783
														if buffer[position] != rune('C') {
															goto l732
														}
														position++
													}
												l783:
													{
														position785, tokenIndex785 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l786
														}
														position++
														goto l785
													l786:
														position, tokenIndex = position785, tokenIndex785
														if buffer[position] != rune('A') {
															goto l732
														}
														position++
													}
												l785:
													{
														position787, tokenIndex787 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l788
														}
														position++
														goto l787
													l788:
														position, tokenIndex = position787, tokenIndex787
														if buffer[position] != rune('L') {
															goto l732
														}
														position++
													}
												l787:
													{
														position789, tokenIndex789 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l790
														}
														position++
														goto l789
													l790:
														position, tokenIndex = position789, tokenIndex789
														if buffer[position] != rune('L') {
															goto l732
														}
														position++
													}
												l789:
													if !_rules[rulews]() {
														goto l732
													}
													{
														position791, tokenIndex791 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l791
														}
														if !_rules[rulesep]() {
															goto l791
														}
														goto l792
													l791:
														position, tokenIndex = position791, tokenIndex791
													}
												l792:
													if !_rules[ruleSrc16]() {
														goto l732
													}
													{
														add(ruleAction95, position)
													}
													add(ruleCall, position782)
												}
												break
											}
										}

									}
								l734:
									add(ruleJump, position733)
								}
								goto l13
							l732:
								position, tokenIndex = position13, tokenIndex13
								{
									position794 := position
									{
										position795, tokenIndex795 := position, tokenIndex
										{
											position797 := position
											{
												position798, tokenIndex798 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l799
												}
												position++
												goto l798
											l799:
												position, tokenIndex = position798, tokenIndex798
												if buffer[position] != rune('I') {
													goto l796
												}
												position++
											}
										l798:
											{
												position800, tokenIndex800 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l801
												}
												position++
												goto l800
											l801:
												position, tokenIndex = position800, tokenIndex800
												if buffer[position] != rune('N') {
													goto l796
												}
												position++
											}
										l800:
											if !_rules[rulews]() {
												goto l796
											}
											if !_rules[ruleReg8]() {
												goto l796
											}
											if !_rules[rulesep]() {
												goto l796
											}
											if !_rules[rulePort]() {
												goto l796
											}
											{
												add(ruleAction100, position)
											}
											add(ruleIN, position797)
										}
										goto l795
									l796:
										position, tokenIndex = position795, tokenIndex795
										{
											position803 := position
											{
												position804, tokenIndex804 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l805
												}
												position++
												goto l804
											l805:
												position, tokenIndex = position804, tokenIndex804
												if buffer[position] != rune('O') {
													goto l0
												}
												position++
											}
										l804:
											{
												position806, tokenIndex806 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l807
												}
												position++
												goto l806
											l807:
												position, tokenIndex = position806, tokenIndex806
												if buffer[position] != rune('U') {
													goto l0
												}
												position++
											}
										l806:
											{
												position808, tokenIndex808 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l809
												}
												position++
												goto l808
											l809:
												position, tokenIndex = position808, tokenIndex808
												if buffer[position] != rune('T') {
													goto l0
												}
												position++
											}
										l808:
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
												add(ruleAction101, position)
											}
											add(ruleOUT, position803)
										}
									}
								l795:
									add(ruleIO, position794)
								}
							}
						l13:
							add(ruleInstruction, position10)
						}
						{
							position811, tokenIndex811 := position, tokenIndex
							if !_rules[ruleLineEnd]() {
								goto l811
							}
							goto l812
						l811:
							position, tokenIndex = position811, tokenIndex811
						}
					l812:
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
						position814, tokenIndex814 := position, tokenIndex
						{
							position816 := position
						l817:
							{
								position818, tokenIndex818 := position, tokenIndex
								if !_rules[rulews]() {
									goto l818
								}
								goto l817
							l818:
								position, tokenIndex = position818, tokenIndex818
							}
							if !_rules[ruleLineEnd]() {
								goto l815
							}
							add(ruleBlankLine, position816)
						}
						goto l814
					l815:
						position, tokenIndex = position814, tokenIndex814
						{
							position819 := position
							{
								position820 := position
							l821:
								{
									position822, tokenIndex822 := position, tokenIndex
									if !_rules[rulews]() {
										goto l822
									}
									goto l821
								l822:
									position, tokenIndex = position822, tokenIndex822
								}
								{
									position823, tokenIndex823 := position, tokenIndex
									{
										position825 := position
										{
											position826, tokenIndex826 := position, tokenIndex
											{
												position828 := position
												{
													position829, tokenIndex829 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l830
													}
													position++
													goto l829
												l830:
													position, tokenIndex = position829, tokenIndex829
													if buffer[position] != rune('P') {
														goto l827
													}
													position++
												}
											l829:
												{
													position831, tokenIndex831 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l832
													}
													position++
													goto l831
												l832:
													position, tokenIndex = position831, tokenIndex831
													if buffer[position] != rune('U') {
														goto l827
													}
													position++
												}
											l831:
												{
													position833, tokenIndex833 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l834
													}
													position++
													goto l833
												l834:
													position, tokenIndex = position833, tokenIndex833
													if buffer[position] != rune('S') {
														goto l827
													}
													position++
												}
											l833:
												{
													position835, tokenIndex835 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l836
													}
													position++
													goto l835
												l836:
													position, tokenIndex = position835, tokenIndex835
													if buffer[position] != rune('H') {
														goto l827
													}
													position++
												}
											l835:
												if !_rules[rulews]() {
													goto l827
												}
												if !_rules[ruleSrc16]() {
													goto l827
												}
												{
													add(ruleAction3, position)
												}
												add(rulePush, position828)
											}
											goto l826
										l827:
											position, tokenIndex = position826, tokenIndex826
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position839 := position
														{
															position840, tokenIndex840 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l841
															}
															position++
															goto l840
														l841:
															position, tokenIndex = position840, tokenIndex840
															if buffer[position] != rune('E') {
																goto l824
															}
															position++
														}
													l840:
														{
															position842, tokenIndex842 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l843
															}
															position++
															goto l842
														l843:
															position, tokenIndex = position842, tokenIndex842
															if buffer[position] != rune('X') {
																goto l824
															}
															position++
														}
													l842:
														if !_rules[rulews]() {
															goto l824
														}
														if !_rules[ruleDst16]() {
															goto l824
														}
														if !_rules[rulesep]() {
															goto l824
														}
														if !_rules[ruleSrc16]() {
															goto l824
														}
														{
															add(ruleAction5, position)
														}
														add(ruleEx, position839)
													}
													break
												case 'P', 'p':
													{
														position845 := position
														{
															position846, tokenIndex846 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l847
															}
															position++
															goto l846
														l847:
															position, tokenIndex = position846, tokenIndex846
															if buffer[position] != rune('P') {
																goto l824
															}
															position++
														}
													l846:
														{
															position848, tokenIndex848 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l849
															}
															position++
															goto l848
														l849:
															position, tokenIndex = position848, tokenIndex848
															if buffer[position] != rune('O') {
																goto l824
															}
															position++
														}
													l848:
														{
															position850, tokenIndex850 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l851
															}
															position++
															goto l850
														l851:
															position, tokenIndex = position850, tokenIndex850
															if buffer[position] != rune('P') {
																goto l824
															}
															position++
														}
													l850:
														if !_rules[rulews]() {
															goto l824
														}
														if !_rules[ruleDst16]() {
															goto l824
														}
														{
															add(ruleAction4, position)
														}
														add(rulePop, position845)
													}
													break
												default:
													{
														position853 := position
														{
															position854, tokenIndex854 := position, tokenIndex
															{
																position856 := position
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
																		goto l855
																	}
																	position++
																}
															l857:
																{
																	position859, tokenIndex859 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l860
																	}
																	position++
																	goto l859
																l860:
																	position, tokenIndex = position859, tokenIndex859
																	if buffer[position] != rune('D') {
																		goto l855
																	}
																	position++
																}
															l859:
																if !_rules[rulews]() {
																	goto l855
																}
																if !_rules[ruleDst16]() {
																	goto l855
																}
																if !_rules[rulesep]() {
																	goto l855
																}
																if !_rules[ruleSrc16]() {
																	goto l855
																}
																{
																	add(ruleAction2, position)
																}
																add(ruleLoad16, position856)
															}
															goto l854
														l855:
															position, tokenIndex = position854, tokenIndex854
															{
																position862 := position
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
																		goto l824
																	}
																	position++
																}
															l863:
																{
																	position865, tokenIndex865 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l866
																	}
																	position++
																	goto l865
																l866:
																	position, tokenIndex = position865, tokenIndex865
																	if buffer[position] != rune('D') {
																		goto l824
																	}
																	position++
																}
															l865:
																if !_rules[rulews]() {
																	goto l824
																}
																{
																	position867 := position
																	{
																		position868, tokenIndex868 := position, tokenIndex
																		if !_rules[ruleReg8]() {
																			goto l869
																		}
																		goto l868
																	l869:
																		position, tokenIndex = position868, tokenIndex868
																		if !_rules[ruleReg16Contents]() {
																			goto l870
																		}
																		goto l868
																	l870:
																		position, tokenIndex = position868, tokenIndex868
																		if !_rules[rulenn_contents]() {
																			goto l824
																		}
																	}
																l868:
																	{
																		add(ruleAction15, position)
																	}
																	add(ruleDst8, position867)
																}
																if !_rules[rulesep]() {
																	goto l824
																}
																if !_rules[ruleSrc8]() {
																	goto l824
																}
																{
																	add(ruleAction1, position)
																}
																add(ruleLoad8, position862)
															}
														}
													l854:
														add(ruleLoad, position853)
													}
													break
												}
											}

										}
									l826:
										add(ruleAssignment, position825)
									}
									goto l823
								l824:
									position, tokenIndex = position823, tokenIndex823
									{
										position874 := position
										{
											position875, tokenIndex875 := position, tokenIndex
											{
												position877 := position
												{
													position878, tokenIndex878 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l879
													}
													position++
													goto l878
												l879:
													position, tokenIndex = position878, tokenIndex878
													if buffer[position] != rune('I') {
														goto l876
													}
													position++
												}
											l878:
												{
													position880, tokenIndex880 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l881
													}
													position++
													goto l880
												l881:
													position, tokenIndex = position880, tokenIndex880
													if buffer[position] != rune('N') {
														goto l876
													}
													position++
												}
											l880:
												{
													position882, tokenIndex882 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l883
													}
													position++
													goto l882
												l883:
													position, tokenIndex = position882, tokenIndex882
													if buffer[position] != rune('C') {
														goto l876
													}
													position++
												}
											l882:
												if !_rules[rulews]() {
													goto l876
												}
												if !_rules[ruleILoc8]() {
													goto l876
												}
												{
													add(ruleAction6, position)
												}
												add(ruleInc16Indexed8, position877)
											}
											goto l875
										l876:
											position, tokenIndex = position875, tokenIndex875
											{
												position886 := position
												{
													position887, tokenIndex887 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l888
													}
													position++
													goto l887
												l888:
													position, tokenIndex = position887, tokenIndex887
													if buffer[position] != rune('I') {
														goto l885
													}
													position++
												}
											l887:
												{
													position889, tokenIndex889 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l890
													}
													position++
													goto l889
												l890:
													position, tokenIndex = position889, tokenIndex889
													if buffer[position] != rune('N') {
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
												if !_rules[ruleLoc16]() {
													goto l885
												}
												{
													add(ruleAction8, position)
												}
												add(ruleInc16, position886)
											}
											goto l875
										l885:
											position, tokenIndex = position875, tokenIndex875
											{
												position894 := position
												{
													position895, tokenIndex895 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l896
													}
													position++
													goto l895
												l896:
													position, tokenIndex = position895, tokenIndex895
													if buffer[position] != rune('I') {
														goto l873
													}
													position++
												}
											l895:
												{
													position897, tokenIndex897 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l898
													}
													position++
													goto l897
												l898:
													position, tokenIndex = position897, tokenIndex897
													if buffer[position] != rune('N') {
														goto l873
													}
													position++
												}
											l897:
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
														goto l873
													}
													position++
												}
											l899:
												if !_rules[rulews]() {
													goto l873
												}
												if !_rules[ruleLoc8]() {
													goto l873
												}
												{
													add(ruleAction7, position)
												}
												add(ruleInc8, position894)
											}
										}
									l875:
										add(ruleInc, position874)
									}
									goto l823
								l873:
									position, tokenIndex = position823, tokenIndex823
									{
										position903 := position
										{
											position904, tokenIndex904 := position, tokenIndex
											{
												position906 := position
												{
													position907, tokenIndex907 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l908
													}
													position++
													goto l907
												l908:
													position, tokenIndex = position907, tokenIndex907
													if buffer[position] != rune('D') {
														goto l905
													}
													position++
												}
											l907:
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
														goto l905
													}
													position++
												}
											l909:
												{
													position911, tokenIndex911 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l912
													}
													position++
													goto l911
												l912:
													position, tokenIndex = position911, tokenIndex911
													if buffer[position] != rune('C') {
														goto l905
													}
													position++
												}
											l911:
												if !_rules[rulews]() {
													goto l905
												}
												if !_rules[ruleILoc8]() {
													goto l905
												}
												{
													add(ruleAction9, position)
												}
												add(ruleDec16Indexed8, position906)
											}
											goto l904
										l905:
											position, tokenIndex = position904, tokenIndex904
											{
												position915 := position
												{
													position916, tokenIndex916 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l917
													}
													position++
													goto l916
												l917:
													position, tokenIndex = position916, tokenIndex916
													if buffer[position] != rune('D') {
														goto l914
													}
													position++
												}
											l916:
												{
													position918, tokenIndex918 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l919
													}
													position++
													goto l918
												l919:
													position, tokenIndex = position918, tokenIndex918
													if buffer[position] != rune('E') {
														goto l914
													}
													position++
												}
											l918:
												{
													position920, tokenIndex920 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l921
													}
													position++
													goto l920
												l921:
													position, tokenIndex = position920, tokenIndex920
													if buffer[position] != rune('C') {
														goto l914
													}
													position++
												}
											l920:
												if !_rules[rulews]() {
													goto l914
												}
												if !_rules[ruleLoc16]() {
													goto l914
												}
												{
													add(ruleAction11, position)
												}
												add(ruleDec16, position915)
											}
											goto l904
										l914:
											position, tokenIndex = position904, tokenIndex904
											{
												position923 := position
												{
													position924, tokenIndex924 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l925
													}
													position++
													goto l924
												l925:
													position, tokenIndex = position924, tokenIndex924
													if buffer[position] != rune('D') {
														goto l902
													}
													position++
												}
											l924:
												{
													position926, tokenIndex926 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l927
													}
													position++
													goto l926
												l927:
													position, tokenIndex = position926, tokenIndex926
													if buffer[position] != rune('E') {
														goto l902
													}
													position++
												}
											l926:
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
														goto l902
													}
													position++
												}
											l928:
												if !_rules[rulews]() {
													goto l902
												}
												if !_rules[ruleLoc8]() {
													goto l902
												}
												{
													add(ruleAction10, position)
												}
												add(ruleDec8, position923)
											}
										}
									l904:
										add(ruleDec, position903)
									}
									goto l823
								l902:
									position, tokenIndex = position823, tokenIndex823
									{
										position932 := position
										{
											position933, tokenIndex933 := position, tokenIndex
											{
												position935 := position
												{
													position936, tokenIndex936 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l937
													}
													position++
													goto l936
												l937:
													position, tokenIndex = position936, tokenIndex936
													if buffer[position] != rune('A') {
														goto l934
													}
													position++
												}
											l936:
												{
													position938, tokenIndex938 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l939
													}
													position++
													goto l938
												l939:
													position, tokenIndex = position938, tokenIndex938
													if buffer[position] != rune('D') {
														goto l934
													}
													position++
												}
											l938:
												{
													position940, tokenIndex940 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l941
													}
													position++
													goto l940
												l941:
													position, tokenIndex = position940, tokenIndex940
													if buffer[position] != rune('D') {
														goto l934
													}
													position++
												}
											l940:
												if !_rules[rulews]() {
													goto l934
												}
												if !_rules[ruleDst16]() {
													goto l934
												}
												if !_rules[rulesep]() {
													goto l934
												}
												if !_rules[ruleSrc16]() {
													goto l934
												}
												{
													add(ruleAction12, position)
												}
												add(ruleAdd16, position935)
											}
											goto l933
										l934:
											position, tokenIndex = position933, tokenIndex933
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
													if buffer[position] != rune('c') {
														goto l950
													}
													position++
													goto l949
												l950:
													position, tokenIndex = position949, tokenIndex949
													if buffer[position] != rune('C') {
														goto l943
													}
													position++
												}
											l949:
												if !_rules[rulews]() {
													goto l943
												}
												if !_rules[ruleDst16]() {
													goto l943
												}
												if !_rules[rulesep]() {
													goto l943
												}
												if !_rules[ruleSrc16]() {
													goto l943
												}
												{
													add(ruleAction13, position)
												}
												add(ruleAdc16, position944)
											}
											goto l933
										l943:
											position, tokenIndex = position933, tokenIndex933
											{
												position952 := position
												{
													position953, tokenIndex953 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l954
													}
													position++
													goto l953
												l954:
													position, tokenIndex = position953, tokenIndex953
													if buffer[position] != rune('S') {
														goto l931
													}
													position++
												}
											l953:
												{
													position955, tokenIndex955 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l956
													}
													position++
													goto l955
												l956:
													position, tokenIndex = position955, tokenIndex955
													if buffer[position] != rune('B') {
														goto l931
													}
													position++
												}
											l955:
												{
													position957, tokenIndex957 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l958
													}
													position++
													goto l957
												l958:
													position, tokenIndex = position957, tokenIndex957
													if buffer[position] != rune('C') {
														goto l931
													}
													position++
												}
											l957:
												if !_rules[rulews]() {
													goto l931
												}
												if !_rules[ruleDst16]() {
													goto l931
												}
												if !_rules[rulesep]() {
													goto l931
												}
												if !_rules[ruleSrc16]() {
													goto l931
												}
												{
													add(ruleAction14, position)
												}
												add(ruleSbc16, position952)
											}
										}
									l933:
										add(ruleAlu16, position932)
									}
									goto l823
								l931:
									position, tokenIndex = position823, tokenIndex823
									{
										position961 := position
										{
											position962, tokenIndex962 := position, tokenIndex
											{
												position964 := position
												{
													position965, tokenIndex965 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l966
													}
													position++
													goto l965
												l966:
													position, tokenIndex = position965, tokenIndex965
													if buffer[position] != rune('A') {
														goto l963
													}
													position++
												}
											l965:
												{
													position967, tokenIndex967 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l968
													}
													position++
													goto l967
												l968:
													position, tokenIndex = position967, tokenIndex967
													if buffer[position] != rune('D') {
														goto l963
													}
													position++
												}
											l967:
												{
													position969, tokenIndex969 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l970
													}
													position++
													goto l969
												l970:
													position, tokenIndex = position969, tokenIndex969
													if buffer[position] != rune('D') {
														goto l963
													}
													position++
												}
											l969:
												if !_rules[rulews]() {
													goto l963
												}
												{
													position971, tokenIndex971 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l972
													}
													position++
													goto l971
												l972:
													position, tokenIndex = position971, tokenIndex971
													if buffer[position] != rune('A') {
														goto l963
													}
													position++
												}
											l971:
												if !_rules[rulesep]() {
													goto l963
												}
												if !_rules[ruleSrc8]() {
													goto l963
												}
												{
													add(ruleAction38, position)
												}
												add(ruleAdd, position964)
											}
											goto l962
										l963:
											position, tokenIndex = position962, tokenIndex962
											{
												position975 := position
												{
													position976, tokenIndex976 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l977
													}
													position++
													goto l976
												l977:
													position, tokenIndex = position976, tokenIndex976
													if buffer[position] != rune('A') {
														goto l974
													}
													position++
												}
											l976:
												{
													position978, tokenIndex978 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l979
													}
													position++
													goto l978
												l979:
													position, tokenIndex = position978, tokenIndex978
													if buffer[position] != rune('D') {
														goto l974
													}
													position++
												}
											l978:
												{
													position980, tokenIndex980 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l981
													}
													position++
													goto l980
												l981:
													position, tokenIndex = position980, tokenIndex980
													if buffer[position] != rune('C') {
														goto l974
													}
													position++
												}
											l980:
												if !_rules[rulews]() {
													goto l974
												}
												{
													position982, tokenIndex982 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l983
													}
													position++
													goto l982
												l983:
													position, tokenIndex = position982, tokenIndex982
													if buffer[position] != rune('A') {
														goto l974
													}
													position++
												}
											l982:
												if !_rules[rulesep]() {
													goto l974
												}
												if !_rules[ruleSrc8]() {
													goto l974
												}
												{
													add(ruleAction39, position)
												}
												add(ruleAdc, position975)
											}
											goto l962
										l974:
											position, tokenIndex = position962, tokenIndex962
											{
												position986 := position
												{
													position987, tokenIndex987 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l988
													}
													position++
													goto l987
												l988:
													position, tokenIndex = position987, tokenIndex987
													if buffer[position] != rune('S') {
														goto l985
													}
													position++
												}
											l987:
												{
													position989, tokenIndex989 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l990
													}
													position++
													goto l989
												l990:
													position, tokenIndex = position989, tokenIndex989
													if buffer[position] != rune('U') {
														goto l985
													}
													position++
												}
											l989:
												{
													position991, tokenIndex991 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l992
													}
													position++
													goto l991
												l992:
													position, tokenIndex = position991, tokenIndex991
													if buffer[position] != rune('B') {
														goto l985
													}
													position++
												}
											l991:
												if !_rules[rulews]() {
													goto l985
												}
												if !_rules[ruleSrc8]() {
													goto l985
												}
												{
													add(ruleAction40, position)
												}
												add(ruleSub, position986)
											}
											goto l962
										l985:
											position, tokenIndex = position962, tokenIndex962
											{
												switch buffer[position] {
												case 'C', 'c':
													{
														position995 := position
														{
															position996, tokenIndex996 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l997
															}
															position++
															goto l996
														l997:
															position, tokenIndex = position996, tokenIndex996
															if buffer[position] != rune('C') {
																goto l960
															}
															position++
														}
													l996:
														{
															position998, tokenIndex998 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l999
															}
															position++
															goto l998
														l999:
															position, tokenIndex = position998, tokenIndex998
															if buffer[position] != rune('P') {
																goto l960
															}
															position++
														}
													l998:
														if !_rules[rulews]() {
															goto l960
														}
														if !_rules[ruleSrc8]() {
															goto l960
														}
														{
															add(ruleAction45, position)
														}
														add(ruleCp, position995)
													}
													break
												case 'O', 'o':
													{
														position1001 := position
														{
															position1002, tokenIndex1002 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1003
															}
															position++
															goto l1002
														l1003:
															position, tokenIndex = position1002, tokenIndex1002
															if buffer[position] != rune('O') {
																goto l960
															}
															position++
														}
													l1002:
														{
															position1004, tokenIndex1004 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1005
															}
															position++
															goto l1004
														l1005:
															position, tokenIndex = position1004, tokenIndex1004
															if buffer[position] != rune('R') {
																goto l960
															}
															position++
														}
													l1004:
														if !_rules[rulews]() {
															goto l960
														}
														if !_rules[ruleSrc8]() {
															goto l960
														}
														{
															add(ruleAction44, position)
														}
														add(ruleOr, position1001)
													}
													break
												case 'X', 'x':
													{
														position1007 := position
														{
															position1008, tokenIndex1008 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l1009
															}
															position++
															goto l1008
														l1009:
															position, tokenIndex = position1008, tokenIndex1008
															if buffer[position] != rune('X') {
																goto l960
															}
															position++
														}
													l1008:
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
																goto l960
															}
															position++
														}
													l1010:
														{
															position1012, tokenIndex1012 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1013
															}
															position++
															goto l1012
														l1013:
															position, tokenIndex = position1012, tokenIndex1012
															if buffer[position] != rune('R') {
																goto l960
															}
															position++
														}
													l1012:
														if !_rules[rulews]() {
															goto l960
														}
														if !_rules[ruleSrc8]() {
															goto l960
														}
														{
															add(ruleAction43, position)
														}
														add(ruleXor, position1007)
													}
													break
												case 'A', 'a':
													{
														position1015 := position
														{
															position1016, tokenIndex1016 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1017
															}
															position++
															goto l1016
														l1017:
															position, tokenIndex = position1016, tokenIndex1016
															if buffer[position] != rune('A') {
																goto l960
															}
															position++
														}
													l1016:
														{
															position1018, tokenIndex1018 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1019
															}
															position++
															goto l1018
														l1019:
															position, tokenIndex = position1018, tokenIndex1018
															if buffer[position] != rune('N') {
																goto l960
															}
															position++
														}
													l1018:
														{
															position1020, tokenIndex1020 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1021
															}
															position++
															goto l1020
														l1021:
															position, tokenIndex = position1020, tokenIndex1020
															if buffer[position] != rune('D') {
																goto l960
															}
															position++
														}
													l1020:
														if !_rules[rulews]() {
															goto l960
														}
														if !_rules[ruleSrc8]() {
															goto l960
														}
														{
															add(ruleAction42, position)
														}
														add(ruleAnd, position1015)
													}
													break
												default:
													{
														position1023 := position
														{
															position1024, tokenIndex1024 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1025
															}
															position++
															goto l1024
														l1025:
															position, tokenIndex = position1024, tokenIndex1024
															if buffer[position] != rune('S') {
																goto l960
															}
															position++
														}
													l1024:
														{
															position1026, tokenIndex1026 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1027
															}
															position++
															goto l1026
														l1027:
															position, tokenIndex = position1026, tokenIndex1026
															if buffer[position] != rune('B') {
																goto l960
															}
															position++
														}
													l1026:
														{
															position1028, tokenIndex1028 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1029
															}
															position++
															goto l1028
														l1029:
															position, tokenIndex = position1028, tokenIndex1028
															if buffer[position] != rune('C') {
																goto l960
															}
															position++
														}
													l1028:
														if !_rules[rulews]() {
															goto l960
														}
														{
															position1030, tokenIndex1030 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1031
															}
															position++
															goto l1030
														l1031:
															position, tokenIndex = position1030, tokenIndex1030
															if buffer[position] != rune('A') {
																goto l960
															}
															position++
														}
													l1030:
														if !_rules[rulesep]() {
															goto l960
														}
														if !_rules[ruleSrc8]() {
															goto l960
														}
														{
															add(ruleAction41, position)
														}
														add(ruleSbc, position1023)
													}
													break
												}
											}

										}
									l962:
										add(ruleAlu, position961)
									}
									goto l823
								l960:
									position, tokenIndex = position823, tokenIndex823
									{
										position1034 := position
										{
											position1035, tokenIndex1035 := position, tokenIndex
											{
												position1037 := position
												{
													position1038, tokenIndex1038 := position, tokenIndex
													{
														position1040 := position
														{
															position1041, tokenIndex1041 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1042
															}
															position++
															goto l1041
														l1042:
															position, tokenIndex = position1041, tokenIndex1041
															if buffer[position] != rune('R') {
																goto l1039
															}
															position++
														}
													l1041:
														{
															position1043, tokenIndex1043 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1044
															}
															position++
															goto l1043
														l1044:
															position, tokenIndex = position1043, tokenIndex1043
															if buffer[position] != rune('L') {
																goto l1039
															}
															position++
														}
													l1043:
														{
															position1045, tokenIndex1045 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1046
															}
															position++
															goto l1045
														l1046:
															position, tokenIndex = position1045, tokenIndex1045
															if buffer[position] != rune('C') {
																goto l1039
															}
															position++
														}
													l1045:
														if !_rules[rulews]() {
															goto l1039
														}
														if !_rules[ruleLoc8]() {
															goto l1039
														}
														{
															position1047, tokenIndex1047 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1047
															}
															if !_rules[ruleCopy8]() {
																goto l1047
															}
															goto l1048
														l1047:
															position, tokenIndex = position1047, tokenIndex1047
														}
													l1048:
														{
															add(ruleAction46, position)
														}
														add(ruleRlc, position1040)
													}
													goto l1038
												l1039:
													position, tokenIndex = position1038, tokenIndex1038
													{
														position1051 := position
														{
															position1052, tokenIndex1052 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1053
															}
															position++
															goto l1052
														l1053:
															position, tokenIndex = position1052, tokenIndex1052
															if buffer[position] != rune('R') {
																goto l1050
															}
															position++
														}
													l1052:
														{
															position1054, tokenIndex1054 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1055
															}
															position++
															goto l1054
														l1055:
															position, tokenIndex = position1054, tokenIndex1054
															if buffer[position] != rune('R') {
																goto l1050
															}
															position++
														}
													l1054:
														{
															position1056, tokenIndex1056 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1057
															}
															position++
															goto l1056
														l1057:
															position, tokenIndex = position1056, tokenIndex1056
															if buffer[position] != rune('C') {
																goto l1050
															}
															position++
														}
													l1056:
														if !_rules[rulews]() {
															goto l1050
														}
														if !_rules[ruleLoc8]() {
															goto l1050
														}
														{
															position1058, tokenIndex1058 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1058
															}
															if !_rules[ruleCopy8]() {
																goto l1058
															}
															goto l1059
														l1058:
															position, tokenIndex = position1058, tokenIndex1058
														}
													l1059:
														{
															add(ruleAction47, position)
														}
														add(ruleRrc, position1051)
													}
													goto l1038
												l1050:
													position, tokenIndex = position1038, tokenIndex1038
													{
														position1062 := position
														{
															position1063, tokenIndex1063 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1064
															}
															position++
															goto l1063
														l1064:
															position, tokenIndex = position1063, tokenIndex1063
															if buffer[position] != rune('R') {
																goto l1061
															}
															position++
														}
													l1063:
														{
															position1065, tokenIndex1065 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1066
															}
															position++
															goto l1065
														l1066:
															position, tokenIndex = position1065, tokenIndex1065
															if buffer[position] != rune('L') {
																goto l1061
															}
															position++
														}
													l1065:
														if !_rules[rulews]() {
															goto l1061
														}
														if !_rules[ruleLoc8]() {
															goto l1061
														}
														{
															position1067, tokenIndex1067 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1067
															}
															if !_rules[ruleCopy8]() {
																goto l1067
															}
															goto l1068
														l1067:
															position, tokenIndex = position1067, tokenIndex1067
														}
													l1068:
														{
															add(ruleAction48, position)
														}
														add(ruleRl, position1062)
													}
													goto l1038
												l1061:
													position, tokenIndex = position1038, tokenIndex1038
													{
														position1071 := position
														{
															position1072, tokenIndex1072 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1073
															}
															position++
															goto l1072
														l1073:
															position, tokenIndex = position1072, tokenIndex1072
															if buffer[position] != rune('R') {
																goto l1070
															}
															position++
														}
													l1072:
														{
															position1074, tokenIndex1074 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1075
															}
															position++
															goto l1074
														l1075:
															position, tokenIndex = position1074, tokenIndex1074
															if buffer[position] != rune('R') {
																goto l1070
															}
															position++
														}
													l1074:
														if !_rules[rulews]() {
															goto l1070
														}
														if !_rules[ruleLoc8]() {
															goto l1070
														}
														{
															position1076, tokenIndex1076 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1076
															}
															if !_rules[ruleCopy8]() {
																goto l1076
															}
															goto l1077
														l1076:
															position, tokenIndex = position1076, tokenIndex1076
														}
													l1077:
														{
															add(ruleAction49, position)
														}
														add(ruleRr, position1071)
													}
													goto l1038
												l1070:
													position, tokenIndex = position1038, tokenIndex1038
													{
														position1080 := position
														{
															position1081, tokenIndex1081 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1082
															}
															position++
															goto l1081
														l1082:
															position, tokenIndex = position1081, tokenIndex1081
															if buffer[position] != rune('S') {
																goto l1079
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
																goto l1079
															}
															position++
														}
													l1083:
														{
															position1085, tokenIndex1085 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1086
															}
															position++
															goto l1085
														l1086:
															position, tokenIndex = position1085, tokenIndex1085
															if buffer[position] != rune('A') {
																goto l1079
															}
															position++
														}
													l1085:
														if !_rules[rulews]() {
															goto l1079
														}
														if !_rules[ruleLoc8]() {
															goto l1079
														}
														{
															position1087, tokenIndex1087 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1087
															}
															if !_rules[ruleCopy8]() {
																goto l1087
															}
															goto l1088
														l1087:
															position, tokenIndex = position1087, tokenIndex1087
														}
													l1088:
														{
															add(ruleAction50, position)
														}
														add(ruleSla, position1080)
													}
													goto l1038
												l1079:
													position, tokenIndex = position1038, tokenIndex1038
													{
														position1091 := position
														{
															position1092, tokenIndex1092 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1093
															}
															position++
															goto l1092
														l1093:
															position, tokenIndex = position1092, tokenIndex1092
															if buffer[position] != rune('S') {
																goto l1090
															}
															position++
														}
													l1092:
														{
															position1094, tokenIndex1094 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1095
															}
															position++
															goto l1094
														l1095:
															position, tokenIndex = position1094, tokenIndex1094
															if buffer[position] != rune('R') {
																goto l1090
															}
															position++
														}
													l1094:
														{
															position1096, tokenIndex1096 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1097
															}
															position++
															goto l1096
														l1097:
															position, tokenIndex = position1096, tokenIndex1096
															if buffer[position] != rune('A') {
																goto l1090
															}
															position++
														}
													l1096:
														if !_rules[rulews]() {
															goto l1090
														}
														if !_rules[ruleLoc8]() {
															goto l1090
														}
														{
															position1098, tokenIndex1098 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1098
															}
															if !_rules[ruleCopy8]() {
																goto l1098
															}
															goto l1099
														l1098:
															position, tokenIndex = position1098, tokenIndex1098
														}
													l1099:
														{
															add(ruleAction51, position)
														}
														add(ruleSra, position1091)
													}
													goto l1038
												l1090:
													position, tokenIndex = position1038, tokenIndex1038
													{
														position1102 := position
														{
															position1103, tokenIndex1103 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1104
															}
															position++
															goto l1103
														l1104:
															position, tokenIndex = position1103, tokenIndex1103
															if buffer[position] != rune('S') {
																goto l1101
															}
															position++
														}
													l1103:
														{
															position1105, tokenIndex1105 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1106
															}
															position++
															goto l1105
														l1106:
															position, tokenIndex = position1105, tokenIndex1105
															if buffer[position] != rune('L') {
																goto l1101
															}
															position++
														}
													l1105:
														{
															position1107, tokenIndex1107 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1108
															}
															position++
															goto l1107
														l1108:
															position, tokenIndex = position1107, tokenIndex1107
															if buffer[position] != rune('L') {
																goto l1101
															}
															position++
														}
													l1107:
														if !_rules[rulews]() {
															goto l1101
														}
														if !_rules[ruleLoc8]() {
															goto l1101
														}
														{
															position1109, tokenIndex1109 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1109
															}
															if !_rules[ruleCopy8]() {
																goto l1109
															}
															goto l1110
														l1109:
															position, tokenIndex = position1109, tokenIndex1109
														}
													l1110:
														{
															add(ruleAction52, position)
														}
														add(ruleSll, position1102)
													}
													goto l1038
												l1101:
													position, tokenIndex = position1038, tokenIndex1038
													{
														position1112 := position
														{
															position1113, tokenIndex1113 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1114
															}
															position++
															goto l1113
														l1114:
															position, tokenIndex = position1113, tokenIndex1113
															if buffer[position] != rune('S') {
																goto l1036
															}
															position++
														}
													l1113:
														{
															position1115, tokenIndex1115 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1116
															}
															position++
															goto l1115
														l1116:
															position, tokenIndex = position1115, tokenIndex1115
															if buffer[position] != rune('R') {
																goto l1036
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
																goto l1036
															}
															position++
														}
													l1117:
														if !_rules[rulews]() {
															goto l1036
														}
														if !_rules[ruleLoc8]() {
															goto l1036
														}
														{
															position1119, tokenIndex1119 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1119
															}
															if !_rules[ruleCopy8]() {
																goto l1119
															}
															goto l1120
														l1119:
															position, tokenIndex = position1119, tokenIndex1119
														}
													l1120:
														{
															add(ruleAction53, position)
														}
														add(ruleSrl, position1112)
													}
												}
											l1038:
												add(ruleRot, position1037)
											}
											goto l1035
										l1036:
											position, tokenIndex = position1035, tokenIndex1035
											{
												switch buffer[position] {
												case 'S', 's':
													{
														position1123 := position
														{
															position1124, tokenIndex1124 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1125
															}
															position++
															goto l1124
														l1125:
															position, tokenIndex = position1124, tokenIndex1124
															if buffer[position] != rune('S') {
																goto l1033
															}
															position++
														}
													l1124:
														{
															position1126, tokenIndex1126 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1127
															}
															position++
															goto l1126
														l1127:
															position, tokenIndex = position1126, tokenIndex1126
															if buffer[position] != rune('E') {
																goto l1033
															}
															position++
														}
													l1126:
														{
															position1128, tokenIndex1128 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1129
															}
															position++
															goto l1128
														l1129:
															position, tokenIndex = position1128, tokenIndex1128
															if buffer[position] != rune('T') {
																goto l1033
															}
															position++
														}
													l1128:
														if !_rules[rulews]() {
															goto l1033
														}
														if !_rules[ruleoctaldigit]() {
															goto l1033
														}
														if !_rules[rulesep]() {
															goto l1033
														}
														if !_rules[ruleLoc8]() {
															goto l1033
														}
														{
															position1130, tokenIndex1130 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1130
															}
															if !_rules[ruleCopy8]() {
																goto l1130
															}
															goto l1131
														l1130:
															position, tokenIndex = position1130, tokenIndex1130
														}
													l1131:
														{
															add(ruleAction56, position)
														}
														add(ruleSet, position1123)
													}
													break
												case 'R', 'r':
													{
														position1133 := position
														{
															position1134, tokenIndex1134 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1135
															}
															position++
															goto l1134
														l1135:
															position, tokenIndex = position1134, tokenIndex1134
															if buffer[position] != rune('R') {
																goto l1033
															}
															position++
														}
													l1134:
														{
															position1136, tokenIndex1136 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1137
															}
															position++
															goto l1136
														l1137:
															position, tokenIndex = position1136, tokenIndex1136
															if buffer[position] != rune('E') {
																goto l1033
															}
															position++
														}
													l1136:
														{
															position1138, tokenIndex1138 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1139
															}
															position++
															goto l1138
														l1139:
															position, tokenIndex = position1138, tokenIndex1138
															if buffer[position] != rune('S') {
																goto l1033
															}
															position++
														}
													l1138:
														if !_rules[rulews]() {
															goto l1033
														}
														if !_rules[ruleoctaldigit]() {
															goto l1033
														}
														if !_rules[rulesep]() {
															goto l1033
														}
														if !_rules[ruleLoc8]() {
															goto l1033
														}
														{
															position1140, tokenIndex1140 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1140
															}
															if !_rules[ruleCopy8]() {
																goto l1140
															}
															goto l1141
														l1140:
															position, tokenIndex = position1140, tokenIndex1140
														}
													l1141:
														{
															add(ruleAction55, position)
														}
														add(ruleRes, position1133)
													}
													break
												default:
													{
														position1143 := position
														{
															position1144, tokenIndex1144 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1145
															}
															position++
															goto l1144
														l1145:
															position, tokenIndex = position1144, tokenIndex1144
															if buffer[position] != rune('B') {
																goto l1033
															}
															position++
														}
													l1144:
														{
															position1146, tokenIndex1146 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1147
															}
															position++
															goto l1146
														l1147:
															position, tokenIndex = position1146, tokenIndex1146
															if buffer[position] != rune('I') {
																goto l1033
															}
															position++
														}
													l1146:
														{
															position1148, tokenIndex1148 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1149
															}
															position++
															goto l1148
														l1149:
															position, tokenIndex = position1148, tokenIndex1148
															if buffer[position] != rune('T') {
																goto l1033
															}
															position++
														}
													l1148:
														if !_rules[rulews]() {
															goto l1033
														}
														if !_rules[ruleoctaldigit]() {
															goto l1033
														}
														if !_rules[rulesep]() {
															goto l1033
														}
														if !_rules[ruleLoc8]() {
															goto l1033
														}
														{
															add(ruleAction54, position)
														}
														add(ruleBit, position1143)
													}
													break
												}
											}

										}
									l1035:
										add(ruleBitOp, position1034)
									}
									goto l823
								l1033:
									position, tokenIndex = position823, tokenIndex823
									{
										position1152 := position
										{
											position1153, tokenIndex1153 := position, tokenIndex
											{
												position1155 := position
												{
													position1156 := position
													{
														position1157, tokenIndex1157 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1158
														}
														position++
														goto l1157
													l1158:
														position, tokenIndex = position1157, tokenIndex1157
														if buffer[position] != rune('R') {
															goto l1154
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
															goto l1154
														}
														position++
													}
												l1159:
													{
														position1161, tokenIndex1161 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1162
														}
														position++
														goto l1161
													l1162:
														position, tokenIndex = position1161, tokenIndex1161
														if buffer[position] != rune('T') {
															goto l1154
														}
														position++
													}
												l1161:
													{
														position1163, tokenIndex1163 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1164
														}
														position++
														goto l1163
													l1164:
														position, tokenIndex = position1163, tokenIndex1163
														if buffer[position] != rune('N') {
															goto l1154
														}
														position++
													}
												l1163:
													add(rulePegText, position1156)
												}
												{
													add(ruleAction71, position)
												}
												add(ruleRetn, position1155)
											}
											goto l1153
										l1154:
											position, tokenIndex = position1153, tokenIndex1153
											{
												position1167 := position
												{
													position1168 := position
													{
														position1169, tokenIndex1169 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1170
														}
														position++
														goto l1169
													l1170:
														position, tokenIndex = position1169, tokenIndex1169
														if buffer[position] != rune('R') {
															goto l1166
														}
														position++
													}
												l1169:
													{
														position1171, tokenIndex1171 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1172
														}
														position++
														goto l1171
													l1172:
														position, tokenIndex = position1171, tokenIndex1171
														if buffer[position] != rune('E') {
															goto l1166
														}
														position++
													}
												l1171:
													{
														position1173, tokenIndex1173 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1174
														}
														position++
														goto l1173
													l1174:
														position, tokenIndex = position1173, tokenIndex1173
														if buffer[position] != rune('T') {
															goto l1166
														}
														position++
													}
												l1173:
													{
														position1175, tokenIndex1175 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1176
														}
														position++
														goto l1175
													l1176:
														position, tokenIndex = position1175, tokenIndex1175
														if buffer[position] != rune('I') {
															goto l1166
														}
														position++
													}
												l1175:
													add(rulePegText, position1168)
												}
												{
													add(ruleAction72, position)
												}
												add(ruleReti, position1167)
											}
											goto l1153
										l1166:
											position, tokenIndex = position1153, tokenIndex1153
											{
												position1179 := position
												{
													position1180 := position
													{
														position1181, tokenIndex1181 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1182
														}
														position++
														goto l1181
													l1182:
														position, tokenIndex = position1181, tokenIndex1181
														if buffer[position] != rune('R') {
															goto l1178
														}
														position++
													}
												l1181:
													{
														position1183, tokenIndex1183 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1184
														}
														position++
														goto l1183
													l1184:
														position, tokenIndex = position1183, tokenIndex1183
														if buffer[position] != rune('R') {
															goto l1178
														}
														position++
													}
												l1183:
													{
														position1185, tokenIndex1185 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1186
														}
														position++
														goto l1185
													l1186:
														position, tokenIndex = position1185, tokenIndex1185
														if buffer[position] != rune('D') {
															goto l1178
														}
														position++
													}
												l1185:
													add(rulePegText, position1180)
												}
												{
													add(ruleAction73, position)
												}
												add(ruleRrd, position1179)
											}
											goto l1153
										l1178:
											position, tokenIndex = position1153, tokenIndex1153
											{
												position1189 := position
												{
													position1190 := position
													{
														position1191, tokenIndex1191 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1192
														}
														position++
														goto l1191
													l1192:
														position, tokenIndex = position1191, tokenIndex1191
														if buffer[position] != rune('I') {
															goto l1188
														}
														position++
													}
												l1191:
													{
														position1193, tokenIndex1193 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1194
														}
														position++
														goto l1193
													l1194:
														position, tokenIndex = position1193, tokenIndex1193
														if buffer[position] != rune('M') {
															goto l1188
														}
														position++
													}
												l1193:
													if buffer[position] != rune(' ') {
														goto l1188
													}
													position++
													if buffer[position] != rune('0') {
														goto l1188
													}
													position++
													add(rulePegText, position1190)
												}
												{
													add(ruleAction75, position)
												}
												add(ruleIm0, position1189)
											}
											goto l1153
										l1188:
											position, tokenIndex = position1153, tokenIndex1153
											{
												position1197 := position
												{
													position1198 := position
													{
														position1199, tokenIndex1199 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1200
														}
														position++
														goto l1199
													l1200:
														position, tokenIndex = position1199, tokenIndex1199
														if buffer[position] != rune('I') {
															goto l1196
														}
														position++
													}
												l1199:
													{
														position1201, tokenIndex1201 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1202
														}
														position++
														goto l1201
													l1202:
														position, tokenIndex = position1201, tokenIndex1201
														if buffer[position] != rune('M') {
															goto l1196
														}
														position++
													}
												l1201:
													if buffer[position] != rune(' ') {
														goto l1196
													}
													position++
													if buffer[position] != rune('1') {
														goto l1196
													}
													position++
													add(rulePegText, position1198)
												}
												{
													add(ruleAction76, position)
												}
												add(ruleIm1, position1197)
											}
											goto l1153
										l1196:
											position, tokenIndex = position1153, tokenIndex1153
											{
												position1205 := position
												{
													position1206 := position
													{
														position1207, tokenIndex1207 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1208
														}
														position++
														goto l1207
													l1208:
														position, tokenIndex = position1207, tokenIndex1207
														if buffer[position] != rune('I') {
															goto l1204
														}
														position++
													}
												l1207:
													{
														position1209, tokenIndex1209 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1210
														}
														position++
														goto l1209
													l1210:
														position, tokenIndex = position1209, tokenIndex1209
														if buffer[position] != rune('M') {
															goto l1204
														}
														position++
													}
												l1209:
													if buffer[position] != rune(' ') {
														goto l1204
													}
													position++
													if buffer[position] != rune('2') {
														goto l1204
													}
													position++
													add(rulePegText, position1206)
												}
												{
													add(ruleAction77, position)
												}
												add(ruleIm2, position1205)
											}
											goto l1153
										l1204:
											position, tokenIndex = position1153, tokenIndex1153
											{
												switch buffer[position] {
												case 'I', 'O', 'i', 'o':
													{
														position1213 := position
														{
															position1214, tokenIndex1214 := position, tokenIndex
															{
																position1216 := position
																{
																	position1217 := position
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
																			goto l1215
																		}
																		position++
																	}
																l1218:
																	{
																		position1220, tokenIndex1220 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1221
																		}
																		position++
																		goto l1220
																	l1221:
																		position, tokenIndex = position1220, tokenIndex1220
																		if buffer[position] != rune('N') {
																			goto l1215
																		}
																		position++
																	}
																l1220:
																	{
																		position1222, tokenIndex1222 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1223
																		}
																		position++
																		goto l1222
																	l1223:
																		position, tokenIndex = position1222, tokenIndex1222
																		if buffer[position] != rune('I') {
																			goto l1215
																		}
																		position++
																	}
																l1222:
																	{
																		position1224, tokenIndex1224 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1225
																		}
																		position++
																		goto l1224
																	l1225:
																		position, tokenIndex = position1224, tokenIndex1224
																		if buffer[position] != rune('R') {
																			goto l1215
																		}
																		position++
																	}
																l1224:
																	add(rulePegText, position1217)
																}
																{
																	add(ruleAction88, position)
																}
																add(ruleInir, position1216)
															}
															goto l1214
														l1215:
															position, tokenIndex = position1214, tokenIndex1214
															{
																position1228 := position
																{
																	position1229 := position
																	{
																		position1230, tokenIndex1230 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1231
																		}
																		position++
																		goto l1230
																	l1231:
																		position, tokenIndex = position1230, tokenIndex1230
																		if buffer[position] != rune('I') {
																			goto l1227
																		}
																		position++
																	}
																l1230:
																	{
																		position1232, tokenIndex1232 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1233
																		}
																		position++
																		goto l1232
																	l1233:
																		position, tokenIndex = position1232, tokenIndex1232
																		if buffer[position] != rune('N') {
																			goto l1227
																		}
																		position++
																	}
																l1232:
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
																			goto l1227
																		}
																		position++
																	}
																l1234:
																	add(rulePegText, position1229)
																}
																{
																	add(ruleAction80, position)
																}
																add(ruleIni, position1228)
															}
															goto l1214
														l1227:
															position, tokenIndex = position1214, tokenIndex1214
															{
																position1238 := position
																{
																	position1239 := position
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
																			goto l1237
																		}
																		position++
																	}
																l1240:
																	{
																		position1242, tokenIndex1242 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1243
																		}
																		position++
																		goto l1242
																	l1243:
																		position, tokenIndex = position1242, tokenIndex1242
																		if buffer[position] != rune('T') {
																			goto l1237
																		}
																		position++
																	}
																l1242:
																	{
																		position1244, tokenIndex1244 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1245
																		}
																		position++
																		goto l1244
																	l1245:
																		position, tokenIndex = position1244, tokenIndex1244
																		if buffer[position] != rune('I') {
																			goto l1237
																		}
																		position++
																	}
																l1244:
																	{
																		position1246, tokenIndex1246 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1247
																		}
																		position++
																		goto l1246
																	l1247:
																		position, tokenIndex = position1246, tokenIndex1246
																		if buffer[position] != rune('R') {
																			goto l1237
																		}
																		position++
																	}
																l1246:
																	add(rulePegText, position1239)
																}
																{
																	add(ruleAction89, position)
																}
																add(ruleOtir, position1238)
															}
															goto l1214
														l1237:
															position, tokenIndex = position1214, tokenIndex1214
															{
																position1250 := position
																{
																	position1251 := position
																	{
																		position1252, tokenIndex1252 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1253
																		}
																		position++
																		goto l1252
																	l1253:
																		position, tokenIndex = position1252, tokenIndex1252
																		if buffer[position] != rune('O') {
																			goto l1249
																		}
																		position++
																	}
																l1252:
																	{
																		position1254, tokenIndex1254 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1255
																		}
																		position++
																		goto l1254
																	l1255:
																		position, tokenIndex = position1254, tokenIndex1254
																		if buffer[position] != rune('U') {
																			goto l1249
																		}
																		position++
																	}
																l1254:
																	{
																		position1256, tokenIndex1256 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1257
																		}
																		position++
																		goto l1256
																	l1257:
																		position, tokenIndex = position1256, tokenIndex1256
																		if buffer[position] != rune('T') {
																			goto l1249
																		}
																		position++
																	}
																l1256:
																	{
																		position1258, tokenIndex1258 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1259
																		}
																		position++
																		goto l1258
																	l1259:
																		position, tokenIndex = position1258, tokenIndex1258
																		if buffer[position] != rune('I') {
																			goto l1249
																		}
																		position++
																	}
																l1258:
																	add(rulePegText, position1251)
																}
																{
																	add(ruleAction81, position)
																}
																add(ruleOuti, position1250)
															}
															goto l1214
														l1249:
															position, tokenIndex = position1214, tokenIndex1214
															{
																position1262 := position
																{
																	position1263 := position
																	{
																		position1264, tokenIndex1264 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1265
																		}
																		position++
																		goto l1264
																	l1265:
																		position, tokenIndex = position1264, tokenIndex1264
																		if buffer[position] != rune('I') {
																			goto l1261
																		}
																		position++
																	}
																l1264:
																	{
																		position1266, tokenIndex1266 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1267
																		}
																		position++
																		goto l1266
																	l1267:
																		position, tokenIndex = position1266, tokenIndex1266
																		if buffer[position] != rune('N') {
																			goto l1261
																		}
																		position++
																	}
																l1266:
																	{
																		position1268, tokenIndex1268 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1269
																		}
																		position++
																		goto l1268
																	l1269:
																		position, tokenIndex = position1268, tokenIndex1268
																		if buffer[position] != rune('D') {
																			goto l1261
																		}
																		position++
																	}
																l1268:
																	{
																		position1270, tokenIndex1270 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1271
																		}
																		position++
																		goto l1270
																	l1271:
																		position, tokenIndex = position1270, tokenIndex1270
																		if buffer[position] != rune('R') {
																			goto l1261
																		}
																		position++
																	}
																l1270:
																	add(rulePegText, position1263)
																}
																{
																	add(ruleAction92, position)
																}
																add(ruleIndr, position1262)
															}
															goto l1214
														l1261:
															position, tokenIndex = position1214, tokenIndex1214
															{
																position1274 := position
																{
																	position1275 := position
																	{
																		position1276, tokenIndex1276 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1277
																		}
																		position++
																		goto l1276
																	l1277:
																		position, tokenIndex = position1276, tokenIndex1276
																		if buffer[position] != rune('I') {
																			goto l1273
																		}
																		position++
																	}
																l1276:
																	{
																		position1278, tokenIndex1278 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1279
																		}
																		position++
																		goto l1278
																	l1279:
																		position, tokenIndex = position1278, tokenIndex1278
																		if buffer[position] != rune('N') {
																			goto l1273
																		}
																		position++
																	}
																l1278:
																	{
																		position1280, tokenIndex1280 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1281
																		}
																		position++
																		goto l1280
																	l1281:
																		position, tokenIndex = position1280, tokenIndex1280
																		if buffer[position] != rune('D') {
																			goto l1273
																		}
																		position++
																	}
																l1280:
																	add(rulePegText, position1275)
																}
																{
																	add(ruleAction84, position)
																}
																add(ruleInd, position1274)
															}
															goto l1214
														l1273:
															position, tokenIndex = position1214, tokenIndex1214
															{
																position1284 := position
																{
																	position1285 := position
																	{
																		position1286, tokenIndex1286 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1287
																		}
																		position++
																		goto l1286
																	l1287:
																		position, tokenIndex = position1286, tokenIndex1286
																		if buffer[position] != rune('O') {
																			goto l1283
																		}
																		position++
																	}
																l1286:
																	{
																		position1288, tokenIndex1288 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1289
																		}
																		position++
																		goto l1288
																	l1289:
																		position, tokenIndex = position1288, tokenIndex1288
																		if buffer[position] != rune('T') {
																			goto l1283
																		}
																		position++
																	}
																l1288:
																	{
																		position1290, tokenIndex1290 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1291
																		}
																		position++
																		goto l1290
																	l1291:
																		position, tokenIndex = position1290, tokenIndex1290
																		if buffer[position] != rune('D') {
																			goto l1283
																		}
																		position++
																	}
																l1290:
																	{
																		position1292, tokenIndex1292 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1293
																		}
																		position++
																		goto l1292
																	l1293:
																		position, tokenIndex = position1292, tokenIndex1292
																		if buffer[position] != rune('R') {
																			goto l1283
																		}
																		position++
																	}
																l1292:
																	add(rulePegText, position1285)
																}
																{
																	add(ruleAction93, position)
																}
																add(ruleOtdr, position1284)
															}
															goto l1214
														l1283:
															position, tokenIndex = position1214, tokenIndex1214
															{
																position1295 := position
																{
																	position1296 := position
																	{
																		position1297, tokenIndex1297 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1298
																		}
																		position++
																		goto l1297
																	l1298:
																		position, tokenIndex = position1297, tokenIndex1297
																		if buffer[position] != rune('O') {
																			goto l1151
																		}
																		position++
																	}
																l1297:
																	{
																		position1299, tokenIndex1299 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1300
																		}
																		position++
																		goto l1299
																	l1300:
																		position, tokenIndex = position1299, tokenIndex1299
																		if buffer[position] != rune('U') {
																			goto l1151
																		}
																		position++
																	}
																l1299:
																	{
																		position1301, tokenIndex1301 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1302
																		}
																		position++
																		goto l1301
																	l1302:
																		position, tokenIndex = position1301, tokenIndex1301
																		if buffer[position] != rune('T') {
																			goto l1151
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
																			goto l1151
																		}
																		position++
																	}
																l1303:
																	add(rulePegText, position1296)
																}
																{
																	add(ruleAction85, position)
																}
																add(ruleOutd, position1295)
															}
														}
													l1214:
														add(ruleBlitIO, position1213)
													}
													break
												case 'R', 'r':
													{
														position1306 := position
														{
															position1307 := position
															{
																position1308, tokenIndex1308 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1309
																}
																position++
																goto l1308
															l1309:
																position, tokenIndex = position1308, tokenIndex1308
																if buffer[position] != rune('R') {
																	goto l1151
																}
																position++
															}
														l1308:
															{
																position1310, tokenIndex1310 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1311
																}
																position++
																goto l1310
															l1311:
																position, tokenIndex = position1310, tokenIndex1310
																if buffer[position] != rune('L') {
																	goto l1151
																}
																position++
															}
														l1310:
															{
																position1312, tokenIndex1312 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1313
																}
																position++
																goto l1312
															l1313:
																position, tokenIndex = position1312, tokenIndex1312
																if buffer[position] != rune('D') {
																	goto l1151
																}
																position++
															}
														l1312:
															add(rulePegText, position1307)
														}
														{
															add(ruleAction74, position)
														}
														add(ruleRld, position1306)
													}
													break
												case 'N', 'n':
													{
														position1315 := position
														{
															position1316 := position
															{
																position1317, tokenIndex1317 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1318
																}
																position++
																goto l1317
															l1318:
																position, tokenIndex = position1317, tokenIndex1317
																if buffer[position] != rune('N') {
																	goto l1151
																}
																position++
															}
														l1317:
															{
																position1319, tokenIndex1319 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1320
																}
																position++
																goto l1319
															l1320:
																position, tokenIndex = position1319, tokenIndex1319
																if buffer[position] != rune('E') {
																	goto l1151
																}
																position++
															}
														l1319:
															{
																position1321, tokenIndex1321 := position, tokenIndex
																if buffer[position] != rune('g') {
																	goto l1322
																}
																position++
																goto l1321
															l1322:
																position, tokenIndex = position1321, tokenIndex1321
																if buffer[position] != rune('G') {
																	goto l1151
																}
																position++
															}
														l1321:
															add(rulePegText, position1316)
														}
														{
															add(ruleAction70, position)
														}
														add(ruleNeg, position1315)
													}
													break
												default:
													{
														position1324 := position
														{
															position1325, tokenIndex1325 := position, tokenIndex
															{
																position1327 := position
																{
																	position1328 := position
																	{
																		position1329, tokenIndex1329 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1330
																		}
																		position++
																		goto l1329
																	l1330:
																		position, tokenIndex = position1329, tokenIndex1329
																		if buffer[position] != rune('L') {
																			goto l1326
																		}
																		position++
																	}
																l1329:
																	{
																		position1331, tokenIndex1331 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1332
																		}
																		position++
																		goto l1331
																	l1332:
																		position, tokenIndex = position1331, tokenIndex1331
																		if buffer[position] != rune('D') {
																			goto l1326
																		}
																		position++
																	}
																l1331:
																	{
																		position1333, tokenIndex1333 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1334
																		}
																		position++
																		goto l1333
																	l1334:
																		position, tokenIndex = position1333, tokenIndex1333
																		if buffer[position] != rune('I') {
																			goto l1326
																		}
																		position++
																	}
																l1333:
																	{
																		position1335, tokenIndex1335 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1336
																		}
																		position++
																		goto l1335
																	l1336:
																		position, tokenIndex = position1335, tokenIndex1335
																		if buffer[position] != rune('R') {
																			goto l1326
																		}
																		position++
																	}
																l1335:
																	add(rulePegText, position1328)
																}
																{
																	add(ruleAction86, position)
																}
																add(ruleLdir, position1327)
															}
															goto l1325
														l1326:
															position, tokenIndex = position1325, tokenIndex1325
															{
																position1339 := position
																{
																	position1340 := position
																	{
																		position1341, tokenIndex1341 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1342
																		}
																		position++
																		goto l1341
																	l1342:
																		position, tokenIndex = position1341, tokenIndex1341
																		if buffer[position] != rune('L') {
																			goto l1338
																		}
																		position++
																	}
																l1341:
																	{
																		position1343, tokenIndex1343 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1344
																		}
																		position++
																		goto l1343
																	l1344:
																		position, tokenIndex = position1343, tokenIndex1343
																		if buffer[position] != rune('D') {
																			goto l1338
																		}
																		position++
																	}
																l1343:
																	{
																		position1345, tokenIndex1345 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1346
																		}
																		position++
																		goto l1345
																	l1346:
																		position, tokenIndex = position1345, tokenIndex1345
																		if buffer[position] != rune('I') {
																			goto l1338
																		}
																		position++
																	}
																l1345:
																	add(rulePegText, position1340)
																}
																{
																	add(ruleAction78, position)
																}
																add(ruleLdi, position1339)
															}
															goto l1325
														l1338:
															position, tokenIndex = position1325, tokenIndex1325
															{
																position1349 := position
																{
																	position1350 := position
																	{
																		position1351, tokenIndex1351 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1352
																		}
																		position++
																		goto l1351
																	l1352:
																		position, tokenIndex = position1351, tokenIndex1351
																		if buffer[position] != rune('C') {
																			goto l1348
																		}
																		position++
																	}
																l1351:
																	{
																		position1353, tokenIndex1353 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1354
																		}
																		position++
																		goto l1353
																	l1354:
																		position, tokenIndex = position1353, tokenIndex1353
																		if buffer[position] != rune('P') {
																			goto l1348
																		}
																		position++
																	}
																l1353:
																	{
																		position1355, tokenIndex1355 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1356
																		}
																		position++
																		goto l1355
																	l1356:
																		position, tokenIndex = position1355, tokenIndex1355
																		if buffer[position] != rune('I') {
																			goto l1348
																		}
																		position++
																	}
																l1355:
																	{
																		position1357, tokenIndex1357 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1358
																		}
																		position++
																		goto l1357
																	l1358:
																		position, tokenIndex = position1357, tokenIndex1357
																		if buffer[position] != rune('R') {
																			goto l1348
																		}
																		position++
																	}
																l1357:
																	add(rulePegText, position1350)
																}
																{
																	add(ruleAction87, position)
																}
																add(ruleCpir, position1349)
															}
															goto l1325
														l1348:
															position, tokenIndex = position1325, tokenIndex1325
															{
																position1361 := position
																{
																	position1362 := position
																	{
																		position1363, tokenIndex1363 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1364
																		}
																		position++
																		goto l1363
																	l1364:
																		position, tokenIndex = position1363, tokenIndex1363
																		if buffer[position] != rune('C') {
																			goto l1360
																		}
																		position++
																	}
																l1363:
																	{
																		position1365, tokenIndex1365 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1366
																		}
																		position++
																		goto l1365
																	l1366:
																		position, tokenIndex = position1365, tokenIndex1365
																		if buffer[position] != rune('P') {
																			goto l1360
																		}
																		position++
																	}
																l1365:
																	{
																		position1367, tokenIndex1367 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1368
																		}
																		position++
																		goto l1367
																	l1368:
																		position, tokenIndex = position1367, tokenIndex1367
																		if buffer[position] != rune('I') {
																			goto l1360
																		}
																		position++
																	}
																l1367:
																	add(rulePegText, position1362)
																}
																{
																	add(ruleAction79, position)
																}
																add(ruleCpi, position1361)
															}
															goto l1325
														l1360:
															position, tokenIndex = position1325, tokenIndex1325
															{
																position1371 := position
																{
																	position1372 := position
																	{
																		position1373, tokenIndex1373 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1374
																		}
																		position++
																		goto l1373
																	l1374:
																		position, tokenIndex = position1373, tokenIndex1373
																		if buffer[position] != rune('L') {
																			goto l1370
																		}
																		position++
																	}
																l1373:
																	{
																		position1375, tokenIndex1375 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1376
																		}
																		position++
																		goto l1375
																	l1376:
																		position, tokenIndex = position1375, tokenIndex1375
																		if buffer[position] != rune('D') {
																			goto l1370
																		}
																		position++
																	}
																l1375:
																	{
																		position1377, tokenIndex1377 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1378
																		}
																		position++
																		goto l1377
																	l1378:
																		position, tokenIndex = position1377, tokenIndex1377
																		if buffer[position] != rune('D') {
																			goto l1370
																		}
																		position++
																	}
																l1377:
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
																			goto l1370
																		}
																		position++
																	}
																l1379:
																	add(rulePegText, position1372)
																}
																{
																	add(ruleAction90, position)
																}
																add(ruleLddr, position1371)
															}
															goto l1325
														l1370:
															position, tokenIndex = position1325, tokenIndex1325
															{
																position1383 := position
																{
																	position1384 := position
																	{
																		position1385, tokenIndex1385 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1386
																		}
																		position++
																		goto l1385
																	l1386:
																		position, tokenIndex = position1385, tokenIndex1385
																		if buffer[position] != rune('L') {
																			goto l1382
																		}
																		position++
																	}
																l1385:
																	{
																		position1387, tokenIndex1387 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1388
																		}
																		position++
																		goto l1387
																	l1388:
																		position, tokenIndex = position1387, tokenIndex1387
																		if buffer[position] != rune('D') {
																			goto l1382
																		}
																		position++
																	}
																l1387:
																	{
																		position1389, tokenIndex1389 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1390
																		}
																		position++
																		goto l1389
																	l1390:
																		position, tokenIndex = position1389, tokenIndex1389
																		if buffer[position] != rune('D') {
																			goto l1382
																		}
																		position++
																	}
																l1389:
																	add(rulePegText, position1384)
																}
																{
																	add(ruleAction82, position)
																}
																add(ruleLdd, position1383)
															}
															goto l1325
														l1382:
															position, tokenIndex = position1325, tokenIndex1325
															{
																position1393 := position
																{
																	position1394 := position
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
																			goto l1392
																		}
																		position++
																	}
																l1395:
																	{
																		position1397, tokenIndex1397 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1398
																		}
																		position++
																		goto l1397
																	l1398:
																		position, tokenIndex = position1397, tokenIndex1397
																		if buffer[position] != rune('P') {
																			goto l1392
																		}
																		position++
																	}
																l1397:
																	{
																		position1399, tokenIndex1399 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1400
																		}
																		position++
																		goto l1399
																	l1400:
																		position, tokenIndex = position1399, tokenIndex1399
																		if buffer[position] != rune('D') {
																			goto l1392
																		}
																		position++
																	}
																l1399:
																	{
																		position1401, tokenIndex1401 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1402
																		}
																		position++
																		goto l1401
																	l1402:
																		position, tokenIndex = position1401, tokenIndex1401
																		if buffer[position] != rune('R') {
																			goto l1392
																		}
																		position++
																	}
																l1401:
																	add(rulePegText, position1394)
																}
																{
																	add(ruleAction91, position)
																}
																add(ruleCpdr, position1393)
															}
															goto l1325
														l1392:
															position, tokenIndex = position1325, tokenIndex1325
															{
																position1404 := position
																{
																	position1405 := position
																	{
																		position1406, tokenIndex1406 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1407
																		}
																		position++
																		goto l1406
																	l1407:
																		position, tokenIndex = position1406, tokenIndex1406
																		if buffer[position] != rune('C') {
																			goto l1151
																		}
																		position++
																	}
																l1406:
																	{
																		position1408, tokenIndex1408 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1409
																		}
																		position++
																		goto l1408
																	l1409:
																		position, tokenIndex = position1408, tokenIndex1408
																		if buffer[position] != rune('P') {
																			goto l1151
																		}
																		position++
																	}
																l1408:
																	{
																		position1410, tokenIndex1410 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1411
																		}
																		position++
																		goto l1410
																	l1411:
																		position, tokenIndex = position1410, tokenIndex1410
																		if buffer[position] != rune('D') {
																			goto l1151
																		}
																		position++
																	}
																l1410:
																	add(rulePegText, position1405)
																}
																{
																	add(ruleAction83, position)
																}
																add(ruleCpd, position1404)
															}
														}
													l1325:
														add(ruleBlit, position1324)
													}
													break
												}
											}

										}
									l1153:
										add(ruleEDSimple, position1152)
									}
									goto l823
								l1151:
									position, tokenIndex = position823, tokenIndex823
									{
										position1414 := position
										{
											position1415, tokenIndex1415 := position, tokenIndex
											{
												position1417 := position
												{
													position1418 := position
													{
														position1419, tokenIndex1419 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1420
														}
														position++
														goto l1419
													l1420:
														position, tokenIndex = position1419, tokenIndex1419
														if buffer[position] != rune('R') {
															goto l1416
														}
														position++
													}
												l1419:
													{
														position1421, tokenIndex1421 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1422
														}
														position++
														goto l1421
													l1422:
														position, tokenIndex = position1421, tokenIndex1421
														if buffer[position] != rune('L') {
															goto l1416
														}
														position++
													}
												l1421:
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
															goto l1416
														}
														position++
													}
												l1423:
													{
														position1425, tokenIndex1425 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1426
														}
														position++
														goto l1425
													l1426:
														position, tokenIndex = position1425, tokenIndex1425
														if buffer[position] != rune('A') {
															goto l1416
														}
														position++
													}
												l1425:
													add(rulePegText, position1418)
												}
												{
													add(ruleAction59, position)
												}
												add(ruleRlca, position1417)
											}
											goto l1415
										l1416:
											position, tokenIndex = position1415, tokenIndex1415
											{
												position1429 := position
												{
													position1430 := position
													{
														position1431, tokenIndex1431 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1432
														}
														position++
														goto l1431
													l1432:
														position, tokenIndex = position1431, tokenIndex1431
														if buffer[position] != rune('R') {
															goto l1428
														}
														position++
													}
												l1431:
													{
														position1433, tokenIndex1433 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1434
														}
														position++
														goto l1433
													l1434:
														position, tokenIndex = position1433, tokenIndex1433
														if buffer[position] != rune('R') {
															goto l1428
														}
														position++
													}
												l1433:
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
															goto l1428
														}
														position++
													}
												l1435:
													{
														position1437, tokenIndex1437 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1438
														}
														position++
														goto l1437
													l1438:
														position, tokenIndex = position1437, tokenIndex1437
														if buffer[position] != rune('A') {
															goto l1428
														}
														position++
													}
												l1437:
													add(rulePegText, position1430)
												}
												{
													add(ruleAction60, position)
												}
												add(ruleRrca, position1429)
											}
											goto l1415
										l1428:
											position, tokenIndex = position1415, tokenIndex1415
											{
												position1441 := position
												{
													position1442 := position
													{
														position1443, tokenIndex1443 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1444
														}
														position++
														goto l1443
													l1444:
														position, tokenIndex = position1443, tokenIndex1443
														if buffer[position] != rune('R') {
															goto l1440
														}
														position++
													}
												l1443:
													{
														position1445, tokenIndex1445 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1446
														}
														position++
														goto l1445
													l1446:
														position, tokenIndex = position1445, tokenIndex1445
														if buffer[position] != rune('L') {
															goto l1440
														}
														position++
													}
												l1445:
													{
														position1447, tokenIndex1447 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1448
														}
														position++
														goto l1447
													l1448:
														position, tokenIndex = position1447, tokenIndex1447
														if buffer[position] != rune('A') {
															goto l1440
														}
														position++
													}
												l1447:
													add(rulePegText, position1442)
												}
												{
													add(ruleAction61, position)
												}
												add(ruleRla, position1441)
											}
											goto l1415
										l1440:
											position, tokenIndex = position1415, tokenIndex1415
											{
												position1451 := position
												{
													position1452 := position
													{
														position1453, tokenIndex1453 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1454
														}
														position++
														goto l1453
													l1454:
														position, tokenIndex = position1453, tokenIndex1453
														if buffer[position] != rune('D') {
															goto l1450
														}
														position++
													}
												l1453:
													{
														position1455, tokenIndex1455 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1456
														}
														position++
														goto l1455
													l1456:
														position, tokenIndex = position1455, tokenIndex1455
														if buffer[position] != rune('A') {
															goto l1450
														}
														position++
													}
												l1455:
													{
														position1457, tokenIndex1457 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1458
														}
														position++
														goto l1457
													l1458:
														position, tokenIndex = position1457, tokenIndex1457
														if buffer[position] != rune('A') {
															goto l1450
														}
														position++
													}
												l1457:
													add(rulePegText, position1452)
												}
												{
													add(ruleAction63, position)
												}
												add(ruleDaa, position1451)
											}
											goto l1415
										l1450:
											position, tokenIndex = position1415, tokenIndex1415
											{
												position1461 := position
												{
													position1462 := position
													{
														position1463, tokenIndex1463 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1464
														}
														position++
														goto l1463
													l1464:
														position, tokenIndex = position1463, tokenIndex1463
														if buffer[position] != rune('C') {
															goto l1460
														}
														position++
													}
												l1463:
													{
														position1465, tokenIndex1465 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1466
														}
														position++
														goto l1465
													l1466:
														position, tokenIndex = position1465, tokenIndex1465
														if buffer[position] != rune('P') {
															goto l1460
														}
														position++
													}
												l1465:
													{
														position1467, tokenIndex1467 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1468
														}
														position++
														goto l1467
													l1468:
														position, tokenIndex = position1467, tokenIndex1467
														if buffer[position] != rune('L') {
															goto l1460
														}
														position++
													}
												l1467:
													add(rulePegText, position1462)
												}
												{
													add(ruleAction64, position)
												}
												add(ruleCpl, position1461)
											}
											goto l1415
										l1460:
											position, tokenIndex = position1415, tokenIndex1415
											{
												position1471 := position
												{
													position1472 := position
													{
														position1473, tokenIndex1473 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1474
														}
														position++
														goto l1473
													l1474:
														position, tokenIndex = position1473, tokenIndex1473
														if buffer[position] != rune('E') {
															goto l1470
														}
														position++
													}
												l1473:
													{
														position1475, tokenIndex1475 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1476
														}
														position++
														goto l1475
													l1476:
														position, tokenIndex = position1475, tokenIndex1475
														if buffer[position] != rune('X') {
															goto l1470
														}
														position++
													}
												l1475:
													{
														position1477, tokenIndex1477 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1478
														}
														position++
														goto l1477
													l1478:
														position, tokenIndex = position1477, tokenIndex1477
														if buffer[position] != rune('X') {
															goto l1470
														}
														position++
													}
												l1477:
													add(rulePegText, position1472)
												}
												{
													add(ruleAction67, position)
												}
												add(ruleExx, position1471)
											}
											goto l1415
										l1470:
											position, tokenIndex = position1415, tokenIndex1415
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position1481 := position
														{
															position1482 := position
															{
																position1483, tokenIndex1483 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1484
																}
																position++
																goto l1483
															l1484:
																position, tokenIndex = position1483, tokenIndex1483
																if buffer[position] != rune('E') {
																	goto l1413
																}
																position++
															}
														l1483:
															{
																position1485, tokenIndex1485 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1486
																}
																position++
																goto l1485
															l1486:
																position, tokenIndex = position1485, tokenIndex1485
																if buffer[position] != rune('I') {
																	goto l1413
																}
																position++
															}
														l1485:
															add(rulePegText, position1482)
														}
														{
															add(ruleAction69, position)
														}
														add(ruleEi, position1481)
													}
													break
												case 'D', 'd':
													{
														position1488 := position
														{
															position1489 := position
															{
																position1490, tokenIndex1490 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1491
																}
																position++
																goto l1490
															l1491:
																position, tokenIndex = position1490, tokenIndex1490
																if buffer[position] != rune('D') {
																	goto l1413
																}
																position++
															}
														l1490:
															{
																position1492, tokenIndex1492 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1493
																}
																position++
																goto l1492
															l1493:
																position, tokenIndex = position1492, tokenIndex1492
																if buffer[position] != rune('I') {
																	goto l1413
																}
																position++
															}
														l1492:
															add(rulePegText, position1489)
														}
														{
															add(ruleAction68, position)
														}
														add(ruleDi, position1488)
													}
													break
												case 'C', 'c':
													{
														position1495 := position
														{
															position1496 := position
															{
																position1497, tokenIndex1497 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1498
																}
																position++
																goto l1497
															l1498:
																position, tokenIndex = position1497, tokenIndex1497
																if buffer[position] != rune('C') {
																	goto l1413
																}
																position++
															}
														l1497:
															{
																position1499, tokenIndex1499 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1500
																}
																position++
																goto l1499
															l1500:
																position, tokenIndex = position1499, tokenIndex1499
																if buffer[position] != rune('C') {
																	goto l1413
																}
																position++
															}
														l1499:
															{
																position1501, tokenIndex1501 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1502
																}
																position++
																goto l1501
															l1502:
																position, tokenIndex = position1501, tokenIndex1501
																if buffer[position] != rune('F') {
																	goto l1413
																}
																position++
															}
														l1501:
															add(rulePegText, position1496)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleCcf, position1495)
													}
													break
												case 'S', 's':
													{
														position1504 := position
														{
															position1505 := position
															{
																position1506, tokenIndex1506 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1507
																}
																position++
																goto l1506
															l1507:
																position, tokenIndex = position1506, tokenIndex1506
																if buffer[position] != rune('S') {
																	goto l1413
																}
																position++
															}
														l1506:
															{
																position1508, tokenIndex1508 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1509
																}
																position++
																goto l1508
															l1509:
																position, tokenIndex = position1508, tokenIndex1508
																if buffer[position] != rune('C') {
																	goto l1413
																}
																position++
															}
														l1508:
															{
																position1510, tokenIndex1510 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1511
																}
																position++
																goto l1510
															l1511:
																position, tokenIndex = position1510, tokenIndex1510
																if buffer[position] != rune('F') {
																	goto l1413
																}
																position++
															}
														l1510:
															add(rulePegText, position1505)
														}
														{
															add(ruleAction65, position)
														}
														add(ruleScf, position1504)
													}
													break
												case 'R', 'r':
													{
														position1513 := position
														{
															position1514 := position
															{
																position1515, tokenIndex1515 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1516
																}
																position++
																goto l1515
															l1516:
																position, tokenIndex = position1515, tokenIndex1515
																if buffer[position] != rune('R') {
																	goto l1413
																}
																position++
															}
														l1515:
															{
																position1517, tokenIndex1517 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1518
																}
																position++
																goto l1517
															l1518:
																position, tokenIndex = position1517, tokenIndex1517
																if buffer[position] != rune('R') {
																	goto l1413
																}
																position++
															}
														l1517:
															{
																position1519, tokenIndex1519 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1520
																}
																position++
																goto l1519
															l1520:
																position, tokenIndex = position1519, tokenIndex1519
																if buffer[position] != rune('A') {
																	goto l1413
																}
																position++
															}
														l1519:
															add(rulePegText, position1514)
														}
														{
															add(ruleAction62, position)
														}
														add(ruleRra, position1513)
													}
													break
												case 'H', 'h':
													{
														position1522 := position
														{
															position1523 := position
															{
																position1524, tokenIndex1524 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l1525
																}
																position++
																goto l1524
															l1525:
																position, tokenIndex = position1524, tokenIndex1524
																if buffer[position] != rune('H') {
																	goto l1413
																}
																position++
															}
														l1524:
															{
																position1526, tokenIndex1526 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1527
																}
																position++
																goto l1526
															l1527:
																position, tokenIndex = position1526, tokenIndex1526
																if buffer[position] != rune('A') {
																	goto l1413
																}
																position++
															}
														l1526:
															{
																position1528, tokenIndex1528 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1529
																}
																position++
																goto l1528
															l1529:
																position, tokenIndex = position1528, tokenIndex1528
																if buffer[position] != rune('L') {
																	goto l1413
																}
																position++
															}
														l1528:
															{
																position1530, tokenIndex1530 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1531
																}
																position++
																goto l1530
															l1531:
																position, tokenIndex = position1530, tokenIndex1530
																if buffer[position] != rune('T') {
																	goto l1413
																}
																position++
															}
														l1530:
															add(rulePegText, position1523)
														}
														{
															add(ruleAction58, position)
														}
														add(ruleHalt, position1522)
													}
													break
												default:
													{
														position1533 := position
														{
															position1534 := position
															{
																position1535, tokenIndex1535 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1536
																}
																position++
																goto l1535
															l1536:
																position, tokenIndex = position1535, tokenIndex1535
																if buffer[position] != rune('N') {
																	goto l1413
																}
																position++
															}
														l1535:
															{
																position1537, tokenIndex1537 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l1538
																}
																position++
																goto l1537
															l1538:
																position, tokenIndex = position1537, tokenIndex1537
																if buffer[position] != rune('O') {
																	goto l1413
																}
																position++
															}
														l1537:
															{
																position1539, tokenIndex1539 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1540
																}
																position++
																goto l1539
															l1540:
																position, tokenIndex = position1539, tokenIndex1539
																if buffer[position] != rune('P') {
																	goto l1413
																}
																position++
															}
														l1539:
															add(rulePegText, position1534)
														}
														{
															add(ruleAction57, position)
														}
														add(ruleNop, position1533)
													}
													break
												}
											}

										}
									l1415:
										add(ruleSimple, position1414)
									}
									goto l823
								l1413:
									position, tokenIndex = position823, tokenIndex823
									{
										position1543 := position
										{
											position1544, tokenIndex1544 := position, tokenIndex
											{
												position1546 := position
												{
													position1547, tokenIndex1547 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1548
													}
													position++
													goto l1547
												l1548:
													position, tokenIndex = position1547, tokenIndex1547
													if buffer[position] != rune('R') {
														goto l1545
													}
													position++
												}
											l1547:
												{
													position1549, tokenIndex1549 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1550
													}
													position++
													goto l1549
												l1550:
													position, tokenIndex = position1549, tokenIndex1549
													if buffer[position] != rune('S') {
														goto l1545
													}
													position++
												}
											l1549:
												{
													position1551, tokenIndex1551 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1552
													}
													position++
													goto l1551
												l1552:
													position, tokenIndex = position1551, tokenIndex1551
													if buffer[position] != rune('T') {
														goto l1545
													}
													position++
												}
											l1551:
												if !_rules[rulews]() {
													goto l1545
												}
												if !_rules[rulen]() {
													goto l1545
												}
												{
													add(ruleAction94, position)
												}
												add(ruleRst, position1546)
											}
											goto l1544
										l1545:
											position, tokenIndex = position1544, tokenIndex1544
											{
												position1555 := position
												{
													position1556, tokenIndex1556 := position, tokenIndex
													if buffer[position] != rune('j') {
														goto l1557
													}
													position++
													goto l1556
												l1557:
													position, tokenIndex = position1556, tokenIndex1556
													if buffer[position] != rune('J') {
														goto l1554
													}
													position++
												}
											l1556:
												{
													position1558, tokenIndex1558 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l1559
													}
													position++
													goto l1558
												l1559:
													position, tokenIndex = position1558, tokenIndex1558
													if buffer[position] != rune('P') {
														goto l1554
													}
													position++
												}
											l1558:
												if !_rules[rulews]() {
													goto l1554
												}
												{
													position1560, tokenIndex1560 := position, tokenIndex
													if !_rules[rulecc]() {
														goto l1560
													}
													if !_rules[rulesep]() {
														goto l1560
													}
													goto l1561
												l1560:
													position, tokenIndex = position1560, tokenIndex1560
												}
											l1561:
												if !_rules[ruleSrc16]() {
													goto l1554
												}
												{
													add(ruleAction97, position)
												}
												add(ruleJp, position1555)
											}
											goto l1544
										l1554:
											position, tokenIndex = position1544, tokenIndex1544
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position1564 := position
														{
															position1565, tokenIndex1565 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1566
															}
															position++
															goto l1565
														l1566:
															position, tokenIndex = position1565, tokenIndex1565
															if buffer[position] != rune('D') {
																goto l1542
															}
															position++
														}
													l1565:
														{
															position1567, tokenIndex1567 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1568
															}
															position++
															goto l1567
														l1568:
															position, tokenIndex = position1567, tokenIndex1567
															if buffer[position] != rune('J') {
																goto l1542
															}
															position++
														}
													l1567:
														{
															position1569, tokenIndex1569 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1570
															}
															position++
															goto l1569
														l1570:
															position, tokenIndex = position1569, tokenIndex1569
															if buffer[position] != rune('N') {
																goto l1542
															}
															position++
														}
													l1569:
														{
															position1571, tokenIndex1571 := position, tokenIndex
															if buffer[position] != rune('z') {
																goto l1572
															}
															position++
															goto l1571
														l1572:
															position, tokenIndex = position1571, tokenIndex1571
															if buffer[position] != rune('Z') {
																goto l1542
															}
															position++
														}
													l1571:
														if !_rules[rulews]() {
															goto l1542
														}
														if !_rules[ruledisp]() {
															goto l1542
														}
														{
															add(ruleAction99, position)
														}
														add(ruleDjnz, position1564)
													}
													break
												case 'J', 'j':
													{
														position1574 := position
														{
															position1575, tokenIndex1575 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1576
															}
															position++
															goto l1575
														l1576:
															position, tokenIndex = position1575, tokenIndex1575
															if buffer[position] != rune('J') {
																goto l1542
															}
															position++
														}
													l1575:
														{
															position1577, tokenIndex1577 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1578
															}
															position++
															goto l1577
														l1578:
															position, tokenIndex = position1577, tokenIndex1577
															if buffer[position] != rune('R') {
																goto l1542
															}
															position++
														}
													l1577:
														if !_rules[rulews]() {
															goto l1542
														}
														{
															position1579, tokenIndex1579 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1579
															}
															if !_rules[rulesep]() {
																goto l1579
															}
															goto l1580
														l1579:
															position, tokenIndex = position1579, tokenIndex1579
														}
													l1580:
														if !_rules[ruledisp]() {
															goto l1542
														}
														{
															add(ruleAction98, position)
														}
														add(ruleJr, position1574)
													}
													break
												case 'R', 'r':
													{
														position1582 := position
														{
															position1583, tokenIndex1583 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1584
															}
															position++
															goto l1583
														l1584:
															position, tokenIndex = position1583, tokenIndex1583
															if buffer[position] != rune('R') {
																goto l1542
															}
															position++
														}
													l1583:
														{
															position1585, tokenIndex1585 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1586
															}
															position++
															goto l1585
														l1586:
															position, tokenIndex = position1585, tokenIndex1585
															if buffer[position] != rune('E') {
																goto l1542
															}
															position++
														}
													l1585:
														{
															position1587, tokenIndex1587 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1588
															}
															position++
															goto l1587
														l1588:
															position, tokenIndex = position1587, tokenIndex1587
															if buffer[position] != rune('T') {
																goto l1542
															}
															position++
														}
													l1587:
														{
															position1589, tokenIndex1589 := position, tokenIndex
															if !_rules[rulews]() {
																goto l1589
															}
															if !_rules[rulecc]() {
																goto l1589
															}
															goto l1590
														l1589:
															position, tokenIndex = position1589, tokenIndex1589
														}
													l1590:
														{
															add(ruleAction96, position)
														}
														add(ruleRet, position1582)
													}
													break
												default:
													{
														position1592 := position
														{
															position1593, tokenIndex1593 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1594
															}
															position++
															goto l1593
														l1594:
															position, tokenIndex = position1593, tokenIndex1593
															if buffer[position] != rune('C') {
																goto l1542
															}
															position++
														}
													l1593:
														{
															position1595, tokenIndex1595 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1596
															}
															position++
															goto l1595
														l1596:
															position, tokenIndex = position1595, tokenIndex1595
															if buffer[position] != rune('A') {
																goto l1542
															}
															position++
														}
													l1595:
														{
															position1597, tokenIndex1597 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1598
															}
															position++
															goto l1597
														l1598:
															position, tokenIndex = position1597, tokenIndex1597
															if buffer[position] != rune('L') {
																goto l1542
															}
															position++
														}
													l1597:
														{
															position1599, tokenIndex1599 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1600
															}
															position++
															goto l1599
														l1600:
															position, tokenIndex = position1599, tokenIndex1599
															if buffer[position] != rune('L') {
																goto l1542
															}
															position++
														}
													l1599:
														if !_rules[rulews]() {
															goto l1542
														}
														{
															position1601, tokenIndex1601 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1601
															}
															if !_rules[rulesep]() {
																goto l1601
															}
															goto l1602
														l1601:
															position, tokenIndex = position1601, tokenIndex1601
														}
													l1602:
														if !_rules[ruleSrc16]() {
															goto l1542
														}
														{
															add(ruleAction95, position)
														}
														add(ruleCall, position1592)
													}
													break
												}
											}

										}
									l1544:
										add(ruleJump, position1543)
									}
									goto l823
								l1542:
									position, tokenIndex = position823, tokenIndex823
									{
										position1604 := position
										{
											position1605, tokenIndex1605 := position, tokenIndex
											{
												position1607 := position
												{
													position1608, tokenIndex1608 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1609
													}
													position++
													goto l1608
												l1609:
													position, tokenIndex = position1608, tokenIndex1608
													if buffer[position] != rune('I') {
														goto l1606
													}
													position++
												}
											l1608:
												{
													position1610, tokenIndex1610 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1611
													}
													position++
													goto l1610
												l1611:
													position, tokenIndex = position1610, tokenIndex1610
													if buffer[position] != rune('N') {
														goto l1606
													}
													position++
												}
											l1610:
												if !_rules[rulews]() {
													goto l1606
												}
												if !_rules[ruleReg8]() {
													goto l1606
												}
												if !_rules[rulesep]() {
													goto l1606
												}
												if !_rules[rulePort]() {
													goto l1606
												}
												{
													add(ruleAction100, position)
												}
												add(ruleIN, position1607)
											}
											goto l1605
										l1606:
											position, tokenIndex = position1605, tokenIndex1605
											{
												position1613 := position
												{
													position1614, tokenIndex1614 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1615
													}
													position++
													goto l1614
												l1615:
													position, tokenIndex = position1614, tokenIndex1614
													if buffer[position] != rune('O') {
														goto l3
													}
													position++
												}
											l1614:
												{
													position1616, tokenIndex1616 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1617
													}
													position++
													goto l1616
												l1617:
													position, tokenIndex = position1616, tokenIndex1616
													if buffer[position] != rune('U') {
														goto l3
													}
													position++
												}
											l1616:
												{
													position1618, tokenIndex1618 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1619
													}
													position++
													goto l1618
												l1619:
													position, tokenIndex = position1618, tokenIndex1618
													if buffer[position] != rune('T') {
														goto l3
													}
													position++
												}
											l1618:
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
													add(ruleAction101, position)
												}
												add(ruleOUT, position1613)
											}
										}
									l1605:
										add(ruleIO, position1604)
									}
								}
							l823:
								add(ruleInstruction, position820)
							}
							{
								position1621, tokenIndex1621 := position, tokenIndex
								if !_rules[ruleLineEnd]() {
									goto l1621
								}
								goto l1622
							l1621:
								position, tokenIndex = position1621, tokenIndex1621
							}
						l1622:
							{
								add(ruleAction0, position)
							}
							add(ruleLine, position819)
						}
					}
				l814:
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
		/* 1 BlankLine <- <(ws* LineEnd)> */
		nil,
		/* 2 Line <- <(Instruction LineEnd? Action0)> */
		nil,
		/* 3 Comment <- <((';' / '#') .*)> */
		nil,
		/* 4 LineEnd <- <(Comment? ('\n' / ':'))> */
		func() bool {
			position1627, tokenIndex1627 := position, tokenIndex
			{
				position1628 := position
				{
					position1629, tokenIndex1629 := position, tokenIndex
					{
						position1631 := position
						{
							position1632, tokenIndex1632 := position, tokenIndex
							if buffer[position] != rune(';') {
								goto l1633
							}
							position++
							goto l1632
						l1633:
							position, tokenIndex = position1632, tokenIndex1632
							if buffer[position] != rune('#') {
								goto l1629
							}
							position++
						}
					l1632:
					l1634:
						{
							position1635, tokenIndex1635 := position, tokenIndex
							if !matchDot() {
								goto l1635
							}
							goto l1634
						l1635:
							position, tokenIndex = position1635, tokenIndex1635
						}
						add(ruleComment, position1631)
					}
					goto l1630
				l1629:
					position, tokenIndex = position1629, tokenIndex1629
				}
			l1630:
				{
					position1636, tokenIndex1636 := position, tokenIndex
					if buffer[position] != rune('\n') {
						goto l1637
					}
					position++
					goto l1636
				l1637:
					position, tokenIndex = position1636, tokenIndex1636
					if buffer[position] != rune(':') {
						goto l1627
					}
					position++
				}
			l1636:
				add(ruleLineEnd, position1628)
			}
			return true
		l1627:
			position, tokenIndex = position1627, tokenIndex1627
			return false
		},
		/* 5 Instruction <- <(ws* (Assignment / Inc / Dec / Alu16 / Alu / BitOp / EDSimple / Simple / Jump / IO))> */
		nil,
		/* 6 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 7 Load <- <(Load16 / Load8)> */
		nil,
		/* 8 Load8 <- <(('l' / 'L') ('d' / 'D') ws Dst8 sep Src8 Action1)> */
		nil,
		/* 9 Load16 <- <(('l' / 'L') ('d' / 'D') ws Dst16 sep Src16 Action2)> */
		nil,
		/* 10 Push <- <(('p' / 'P') ('u' / 'U') ('s' / 'S') ('h' / 'H') ws Src16 Action3)> */
		nil,
		/* 11 Pop <- <(('p' / 'P') ('o' / 'O') ('p' / 'P') ws Dst16 Action4)> */
		nil,
		/* 12 Ex <- <(('e' / 'E') ('x' / 'X') ws Dst16 sep Src16 Action5)> */
		nil,
		/* 13 Inc <- <(Inc16Indexed8 / Inc16 / Inc8)> */
		nil,
		/* 14 Inc16Indexed8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws ILoc8 Action6)> */
		nil,
		/* 15 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action7)> */
		nil,
		/* 16 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action8)> */
		nil,
		/* 17 Dec <- <(Dec16Indexed8 / Dec16 / Dec8)> */
		nil,
		/* 18 Dec16Indexed8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws ILoc8 Action9)> */
		nil,
		/* 19 Dec8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc8 Action10)> */
		nil,
		/* 20 Dec16 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc16 Action11)> */
		nil,
		/* 21 Alu16 <- <(Add16 / Adc16 / Sbc16)> */
		nil,
		/* 22 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action12)> */
		nil,
		/* 23 Adc16 <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws Dst16 sep Src16 Action13)> */
		nil,
		/* 24 Sbc16 <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws Dst16 sep Src16 Action14)> */
		nil,
		/* 25 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action15)> */
		nil,
		/* 26 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action16)> */
		func() bool {
			position1659, tokenIndex1659 := position, tokenIndex
			{
				position1660 := position
				{
					position1661, tokenIndex1661 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1662
					}
					goto l1661
				l1662:
					position, tokenIndex = position1661, tokenIndex1661
					if !_rules[ruleReg8]() {
						goto l1663
					}
					goto l1661
				l1663:
					position, tokenIndex = position1661, tokenIndex1661
					if !_rules[ruleReg16Contents]() {
						goto l1664
					}
					goto l1661
				l1664:
					position, tokenIndex = position1661, tokenIndex1661
					if !_rules[rulenn_contents]() {
						goto l1659
					}
				}
			l1661:
				{
					add(ruleAction16, position)
				}
				add(ruleSrc8, position1660)
			}
			return true
		l1659:
			position, tokenIndex = position1659, tokenIndex1659
			return false
		},
		/* 27 Loc8 <- <((Reg8 / Reg16Contents) Action17)> */
		func() bool {
			position1666, tokenIndex1666 := position, tokenIndex
			{
				position1667 := position
				{
					position1668, tokenIndex1668 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1669
					}
					goto l1668
				l1669:
					position, tokenIndex = position1668, tokenIndex1668
					if !_rules[ruleReg16Contents]() {
						goto l1666
					}
				}
			l1668:
				{
					add(ruleAction17, position)
				}
				add(ruleLoc8, position1667)
			}
			return true
		l1666:
			position, tokenIndex = position1666, tokenIndex1666
			return false
		},
		/* 28 Copy8 <- <(Reg8 Action18)> */
		func() bool {
			position1671, tokenIndex1671 := position, tokenIndex
			{
				position1672 := position
				if !_rules[ruleReg8]() {
					goto l1671
				}
				{
					add(ruleAction18, position)
				}
				add(ruleCopy8, position1672)
			}
			return true
		l1671:
			position, tokenIndex = position1671, tokenIndex1671
			return false
		},
		/* 29 ILoc8 <- <(IReg8 Action19)> */
		func() bool {
			position1674, tokenIndex1674 := position, tokenIndex
			{
				position1675 := position
				if !_rules[ruleIReg8]() {
					goto l1674
				}
				{
					add(ruleAction19, position)
				}
				add(ruleILoc8, position1675)
			}
			return true
		l1674:
			position, tokenIndex = position1674, tokenIndex1674
			return false
		},
		/* 30 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action20)> */
		func() bool {
			position1677, tokenIndex1677 := position, tokenIndex
			{
				position1678 := position
				{
					position1679 := position
					{
						position1680, tokenIndex1680 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1681
						}
						goto l1680
					l1681:
						position, tokenIndex = position1680, tokenIndex1680
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1683 := position
									{
										position1684, tokenIndex1684 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1685
										}
										position++
										goto l1684
									l1685:
										position, tokenIndex = position1684, tokenIndex1684
										if buffer[position] != rune('R') {
											goto l1677
										}
										position++
									}
								l1684:
									add(ruleR, position1683)
								}
								break
							case 'I', 'i':
								{
									position1686 := position
									{
										position1687, tokenIndex1687 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1688
										}
										position++
										goto l1687
									l1688:
										position, tokenIndex = position1687, tokenIndex1687
										if buffer[position] != rune('I') {
											goto l1677
										}
										position++
									}
								l1687:
									add(ruleI, position1686)
								}
								break
							case 'L', 'l':
								{
									position1689 := position
									{
										position1690, tokenIndex1690 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1691
										}
										position++
										goto l1690
									l1691:
										position, tokenIndex = position1690, tokenIndex1690
										if buffer[position] != rune('L') {
											goto l1677
										}
										position++
									}
								l1690:
									add(ruleL, position1689)
								}
								break
							case 'H', 'h':
								{
									position1692 := position
									{
										position1693, tokenIndex1693 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1694
										}
										position++
										goto l1693
									l1694:
										position, tokenIndex = position1693, tokenIndex1693
										if buffer[position] != rune('H') {
											goto l1677
										}
										position++
									}
								l1693:
									add(ruleH, position1692)
								}
								break
							case 'E', 'e':
								{
									position1695 := position
									{
										position1696, tokenIndex1696 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1697
										}
										position++
										goto l1696
									l1697:
										position, tokenIndex = position1696, tokenIndex1696
										if buffer[position] != rune('E') {
											goto l1677
										}
										position++
									}
								l1696:
									add(ruleE, position1695)
								}
								break
							case 'D', 'd':
								{
									position1698 := position
									{
										position1699, tokenIndex1699 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1700
										}
										position++
										goto l1699
									l1700:
										position, tokenIndex = position1699, tokenIndex1699
										if buffer[position] != rune('D') {
											goto l1677
										}
										position++
									}
								l1699:
									add(ruleD, position1698)
								}
								break
							case 'C', 'c':
								{
									position1701 := position
									{
										position1702, tokenIndex1702 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1703
										}
										position++
										goto l1702
									l1703:
										position, tokenIndex = position1702, tokenIndex1702
										if buffer[position] != rune('C') {
											goto l1677
										}
										position++
									}
								l1702:
									add(ruleC, position1701)
								}
								break
							case 'B', 'b':
								{
									position1704 := position
									{
										position1705, tokenIndex1705 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1706
										}
										position++
										goto l1705
									l1706:
										position, tokenIndex = position1705, tokenIndex1705
										if buffer[position] != rune('B') {
											goto l1677
										}
										position++
									}
								l1705:
									add(ruleB, position1704)
								}
								break
							case 'F', 'f':
								{
									position1707 := position
									{
										position1708, tokenIndex1708 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1709
										}
										position++
										goto l1708
									l1709:
										position, tokenIndex = position1708, tokenIndex1708
										if buffer[position] != rune('F') {
											goto l1677
										}
										position++
									}
								l1708:
									add(ruleF, position1707)
								}
								break
							default:
								{
									position1710 := position
									{
										position1711, tokenIndex1711 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1712
										}
										position++
										goto l1711
									l1712:
										position, tokenIndex = position1711, tokenIndex1711
										if buffer[position] != rune('A') {
											goto l1677
										}
										position++
									}
								l1711:
									add(ruleA, position1710)
								}
								break
							}
						}

					}
				l1680:
					add(rulePegText, position1679)
				}
				{
					add(ruleAction20, position)
				}
				add(ruleReg8, position1678)
			}
			return true
		l1677:
			position, tokenIndex = position1677, tokenIndex1677
			return false
		},
		/* 31 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action21)> */
		func() bool {
			position1714, tokenIndex1714 := position, tokenIndex
			{
				position1715 := position
				{
					position1716 := position
					{
						position1717, tokenIndex1717 := position, tokenIndex
						{
							position1719 := position
							{
								position1720, tokenIndex1720 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1721
								}
								position++
								goto l1720
							l1721:
								position, tokenIndex = position1720, tokenIndex1720
								if buffer[position] != rune('I') {
									goto l1718
								}
								position++
							}
						l1720:
							{
								position1722, tokenIndex1722 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1723
								}
								position++
								goto l1722
							l1723:
								position, tokenIndex = position1722, tokenIndex1722
								if buffer[position] != rune('X') {
									goto l1718
								}
								position++
							}
						l1722:
							{
								position1724, tokenIndex1724 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1725
								}
								position++
								goto l1724
							l1725:
								position, tokenIndex = position1724, tokenIndex1724
								if buffer[position] != rune('H') {
									goto l1718
								}
								position++
							}
						l1724:
							add(ruleIXH, position1719)
						}
						goto l1717
					l1718:
						position, tokenIndex = position1717, tokenIndex1717
						{
							position1727 := position
							{
								position1728, tokenIndex1728 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1729
								}
								position++
								goto l1728
							l1729:
								position, tokenIndex = position1728, tokenIndex1728
								if buffer[position] != rune('I') {
									goto l1726
								}
								position++
							}
						l1728:
							{
								position1730, tokenIndex1730 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1731
								}
								position++
								goto l1730
							l1731:
								position, tokenIndex = position1730, tokenIndex1730
								if buffer[position] != rune('X') {
									goto l1726
								}
								position++
							}
						l1730:
							{
								position1732, tokenIndex1732 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1733
								}
								position++
								goto l1732
							l1733:
								position, tokenIndex = position1732, tokenIndex1732
								if buffer[position] != rune('L') {
									goto l1726
								}
								position++
							}
						l1732:
							add(ruleIXL, position1727)
						}
						goto l1717
					l1726:
						position, tokenIndex = position1717, tokenIndex1717
						{
							position1735 := position
							{
								position1736, tokenIndex1736 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1737
								}
								position++
								goto l1736
							l1737:
								position, tokenIndex = position1736, tokenIndex1736
								if buffer[position] != rune('I') {
									goto l1734
								}
								position++
							}
						l1736:
							{
								position1738, tokenIndex1738 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1739
								}
								position++
								goto l1738
							l1739:
								position, tokenIndex = position1738, tokenIndex1738
								if buffer[position] != rune('Y') {
									goto l1734
								}
								position++
							}
						l1738:
							{
								position1740, tokenIndex1740 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1741
								}
								position++
								goto l1740
							l1741:
								position, tokenIndex = position1740, tokenIndex1740
								if buffer[position] != rune('H') {
									goto l1734
								}
								position++
							}
						l1740:
							add(ruleIYH, position1735)
						}
						goto l1717
					l1734:
						position, tokenIndex = position1717, tokenIndex1717
						{
							position1742 := position
							{
								position1743, tokenIndex1743 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1744
								}
								position++
								goto l1743
							l1744:
								position, tokenIndex = position1743, tokenIndex1743
								if buffer[position] != rune('I') {
									goto l1714
								}
								position++
							}
						l1743:
							{
								position1745, tokenIndex1745 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1746
								}
								position++
								goto l1745
							l1746:
								position, tokenIndex = position1745, tokenIndex1745
								if buffer[position] != rune('Y') {
									goto l1714
								}
								position++
							}
						l1745:
							{
								position1747, tokenIndex1747 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1748
								}
								position++
								goto l1747
							l1748:
								position, tokenIndex = position1747, tokenIndex1747
								if buffer[position] != rune('L') {
									goto l1714
								}
								position++
							}
						l1747:
							add(ruleIYL, position1742)
						}
					}
				l1717:
					add(rulePegText, position1716)
				}
				{
					add(ruleAction21, position)
				}
				add(ruleIReg8, position1715)
			}
			return true
		l1714:
			position, tokenIndex = position1714, tokenIndex1714
			return false
		},
		/* 32 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action22)> */
		func() bool {
			position1750, tokenIndex1750 := position, tokenIndex
			{
				position1751 := position
				{
					position1752, tokenIndex1752 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1753
					}
					goto l1752
				l1753:
					position, tokenIndex = position1752, tokenIndex1752
					if !_rules[rulenn_contents]() {
						goto l1754
					}
					goto l1752
				l1754:
					position, tokenIndex = position1752, tokenIndex1752
					if !_rules[ruleReg16Contents]() {
						goto l1750
					}
				}
			l1752:
				{
					add(ruleAction22, position)
				}
				add(ruleDst16, position1751)
			}
			return true
		l1750:
			position, tokenIndex = position1750, tokenIndex1750
			return false
		},
		/* 33 Src16 <- <((Reg16 / nn / nn_contents) Action23)> */
		func() bool {
			position1756, tokenIndex1756 := position, tokenIndex
			{
				position1757 := position
				{
					position1758, tokenIndex1758 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1759
					}
					goto l1758
				l1759:
					position, tokenIndex = position1758, tokenIndex1758
					if !_rules[rulenn]() {
						goto l1760
					}
					goto l1758
				l1760:
					position, tokenIndex = position1758, tokenIndex1758
					if !_rules[rulenn_contents]() {
						goto l1756
					}
				}
			l1758:
				{
					add(ruleAction23, position)
				}
				add(ruleSrc16, position1757)
			}
			return true
		l1756:
			position, tokenIndex = position1756, tokenIndex1756
			return false
		},
		/* 34 Loc16 <- <(Reg16 Action24)> */
		func() bool {
			position1762, tokenIndex1762 := position, tokenIndex
			{
				position1763 := position
				if !_rules[ruleReg16]() {
					goto l1762
				}
				{
					add(ruleAction24, position)
				}
				add(ruleLoc16, position1763)
			}
			return true
		l1762:
			position, tokenIndex = position1762, tokenIndex1762
			return false
		},
		/* 35 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action25)> */
		func() bool {
			position1765, tokenIndex1765 := position, tokenIndex
			{
				position1766 := position
				{
					position1767 := position
					{
						position1768, tokenIndex1768 := position, tokenIndex
						{
							position1770 := position
							{
								position1771, tokenIndex1771 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1772
								}
								position++
								goto l1771
							l1772:
								position, tokenIndex = position1771, tokenIndex1771
								if buffer[position] != rune('A') {
									goto l1769
								}
								position++
							}
						l1771:
							{
								position1773, tokenIndex1773 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1774
								}
								position++
								goto l1773
							l1774:
								position, tokenIndex = position1773, tokenIndex1773
								if buffer[position] != rune('F') {
									goto l1769
								}
								position++
							}
						l1773:
							if buffer[position] != rune('\'') {
								goto l1769
							}
							position++
							add(ruleAF_PRIME, position1770)
						}
						goto l1768
					l1769:
						position, tokenIndex = position1768, tokenIndex1768
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1765
								}
								break
							case 'S', 's':
								{
									position1776 := position
									{
										position1777, tokenIndex1777 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1778
										}
										position++
										goto l1777
									l1778:
										position, tokenIndex = position1777, tokenIndex1777
										if buffer[position] != rune('S') {
											goto l1765
										}
										position++
									}
								l1777:
									{
										position1779, tokenIndex1779 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1780
										}
										position++
										goto l1779
									l1780:
										position, tokenIndex = position1779, tokenIndex1779
										if buffer[position] != rune('P') {
											goto l1765
										}
										position++
									}
								l1779:
									add(ruleSP, position1776)
								}
								break
							case 'H', 'h':
								{
									position1781 := position
									{
										position1782, tokenIndex1782 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1783
										}
										position++
										goto l1782
									l1783:
										position, tokenIndex = position1782, tokenIndex1782
										if buffer[position] != rune('H') {
											goto l1765
										}
										position++
									}
								l1782:
									{
										position1784, tokenIndex1784 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1785
										}
										position++
										goto l1784
									l1785:
										position, tokenIndex = position1784, tokenIndex1784
										if buffer[position] != rune('L') {
											goto l1765
										}
										position++
									}
								l1784:
									add(ruleHL, position1781)
								}
								break
							case 'D', 'd':
								{
									position1786 := position
									{
										position1787, tokenIndex1787 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1788
										}
										position++
										goto l1787
									l1788:
										position, tokenIndex = position1787, tokenIndex1787
										if buffer[position] != rune('D') {
											goto l1765
										}
										position++
									}
								l1787:
									{
										position1789, tokenIndex1789 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1790
										}
										position++
										goto l1789
									l1790:
										position, tokenIndex = position1789, tokenIndex1789
										if buffer[position] != rune('E') {
											goto l1765
										}
										position++
									}
								l1789:
									add(ruleDE, position1786)
								}
								break
							case 'B', 'b':
								{
									position1791 := position
									{
										position1792, tokenIndex1792 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1793
										}
										position++
										goto l1792
									l1793:
										position, tokenIndex = position1792, tokenIndex1792
										if buffer[position] != rune('B') {
											goto l1765
										}
										position++
									}
								l1792:
									{
										position1794, tokenIndex1794 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1795
										}
										position++
										goto l1794
									l1795:
										position, tokenIndex = position1794, tokenIndex1794
										if buffer[position] != rune('C') {
											goto l1765
										}
										position++
									}
								l1794:
									add(ruleBC, position1791)
								}
								break
							default:
								{
									position1796 := position
									{
										position1797, tokenIndex1797 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1798
										}
										position++
										goto l1797
									l1798:
										position, tokenIndex = position1797, tokenIndex1797
										if buffer[position] != rune('A') {
											goto l1765
										}
										position++
									}
								l1797:
									{
										position1799, tokenIndex1799 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1800
										}
										position++
										goto l1799
									l1800:
										position, tokenIndex = position1799, tokenIndex1799
										if buffer[position] != rune('F') {
											goto l1765
										}
										position++
									}
								l1799:
									add(ruleAF, position1796)
								}
								break
							}
						}

					}
				l1768:
					add(rulePegText, position1767)
				}
				{
					add(ruleAction25, position)
				}
				add(ruleReg16, position1766)
			}
			return true
		l1765:
			position, tokenIndex = position1765, tokenIndex1765
			return false
		},
		/* 36 IReg16 <- <(<(IX / IY)> Action26)> */
		func() bool {
			position1802, tokenIndex1802 := position, tokenIndex
			{
				position1803 := position
				{
					position1804 := position
					{
						position1805, tokenIndex1805 := position, tokenIndex
						{
							position1807 := position
							{
								position1808, tokenIndex1808 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1809
								}
								position++
								goto l1808
							l1809:
								position, tokenIndex = position1808, tokenIndex1808
								if buffer[position] != rune('I') {
									goto l1806
								}
								position++
							}
						l1808:
							{
								position1810, tokenIndex1810 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1811
								}
								position++
								goto l1810
							l1811:
								position, tokenIndex = position1810, tokenIndex1810
								if buffer[position] != rune('X') {
									goto l1806
								}
								position++
							}
						l1810:
							add(ruleIX, position1807)
						}
						goto l1805
					l1806:
						position, tokenIndex = position1805, tokenIndex1805
						{
							position1812 := position
							{
								position1813, tokenIndex1813 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1814
								}
								position++
								goto l1813
							l1814:
								position, tokenIndex = position1813, tokenIndex1813
								if buffer[position] != rune('I') {
									goto l1802
								}
								position++
							}
						l1813:
							{
								position1815, tokenIndex1815 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1816
								}
								position++
								goto l1815
							l1816:
								position, tokenIndex = position1815, tokenIndex1815
								if buffer[position] != rune('Y') {
									goto l1802
								}
								position++
							}
						l1815:
							add(ruleIY, position1812)
						}
					}
				l1805:
					add(rulePegText, position1804)
				}
				{
					add(ruleAction26, position)
				}
				add(ruleIReg16, position1803)
			}
			return true
		l1802:
			position, tokenIndex = position1802, tokenIndex1802
			return false
		},
		/* 37 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1818, tokenIndex1818 := position, tokenIndex
			{
				position1819 := position
				{
					position1820, tokenIndex1820 := position, tokenIndex
					{
						position1822 := position
						if buffer[position] != rune('(') {
							goto l1821
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1821
						}
						{
							position1823, tokenIndex1823 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1823
							}
							goto l1824
						l1823:
							position, tokenIndex = position1823, tokenIndex1823
						}
					l1824:
						if !_rules[ruledisp]() {
							goto l1821
						}
						{
							position1825, tokenIndex1825 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1825
							}
							goto l1826
						l1825:
							position, tokenIndex = position1825, tokenIndex1825
						}
					l1826:
						if buffer[position] != rune(')') {
							goto l1821
						}
						position++
						{
							add(ruleAction28, position)
						}
						add(ruleIndexedR16C, position1822)
					}
					goto l1820
				l1821:
					position, tokenIndex = position1820, tokenIndex1820
					{
						position1828 := position
						if buffer[position] != rune('(') {
							goto l1818
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1818
						}
						if buffer[position] != rune(')') {
							goto l1818
						}
						position++
						{
							add(ruleAction27, position)
						}
						add(rulePlainR16C, position1828)
					}
				}
			l1820:
				add(ruleReg16Contents, position1819)
			}
			return true
		l1818:
			position, tokenIndex = position1818, tokenIndex1818
			return false
		},
		/* 38 PlainR16C <- <('(' Reg16 ')' Action27)> */
		nil,
		/* 39 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action28)> */
		nil,
		/* 40 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1832, tokenIndex1832 := position, tokenIndex
			{
				position1833 := position
				{
					position1834, tokenIndex1834 := position, tokenIndex
					{
						position1836 := position
						{
							position1837 := position
							if !_rules[rulehexdigit]() {
								goto l1835
							}
							if !_rules[rulehexdigit]() {
								goto l1835
							}
							add(rulePegText, position1837)
						}
						{
							position1838, tokenIndex1838 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1839
							}
							position++
							goto l1838
						l1839:
							position, tokenIndex = position1838, tokenIndex1838
							if buffer[position] != rune('H') {
								goto l1835
							}
							position++
						}
					l1838:
						{
							add(ruleAction32, position)
						}
						add(rulehexByteH, position1836)
					}
					goto l1834
				l1835:
					position, tokenIndex = position1834, tokenIndex1834
					{
						position1842 := position
						if buffer[position] != rune('0') {
							goto l1841
						}
						position++
						{
							position1843, tokenIndex1843 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1844
							}
							position++
							goto l1843
						l1844:
							position, tokenIndex = position1843, tokenIndex1843
							if buffer[position] != rune('X') {
								goto l1841
							}
							position++
						}
					l1843:
						{
							position1845 := position
							if !_rules[rulehexdigit]() {
								goto l1841
							}
							if !_rules[rulehexdigit]() {
								goto l1841
							}
							add(rulePegText, position1845)
						}
						{
							add(ruleAction33, position)
						}
						add(rulehexByte0x, position1842)
					}
					goto l1834
				l1841:
					position, tokenIndex = position1834, tokenIndex1834
					{
						position1847 := position
						{
							position1848 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1832
							}
							position++
						l1849:
							{
								position1850, tokenIndex1850 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1850
								}
								position++
								goto l1849
							l1850:
								position, tokenIndex = position1850, tokenIndex1850
							}
							add(rulePegText, position1848)
						}
						{
							add(ruleAction34, position)
						}
						add(ruledecimalByte, position1847)
					}
				}
			l1834:
				add(rulen, position1833)
			}
			return true
		l1832:
			position, tokenIndex = position1832, tokenIndex1832
			return false
		},
		/* 41 nn <- <(hexWordH / hexWord0x)> */
		func() bool {
			position1852, tokenIndex1852 := position, tokenIndex
			{
				position1853 := position
				{
					position1854, tokenIndex1854 := position, tokenIndex
					{
						position1856 := position
						{
							position1857 := position
							if !_rules[rulehexdigit]() {
								goto l1855
							}
							if !_rules[rulehexdigit]() {
								goto l1855
							}
							if !_rules[rulehexdigit]() {
								goto l1855
							}
							if !_rules[rulehexdigit]() {
								goto l1855
							}
							add(rulePegText, position1857)
						}
						{
							position1858, tokenIndex1858 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1859
							}
							position++
							goto l1858
						l1859:
							position, tokenIndex = position1858, tokenIndex1858
							if buffer[position] != rune('H') {
								goto l1855
							}
							position++
						}
					l1858:
						{
							add(ruleAction35, position)
						}
						add(rulehexWordH, position1856)
					}
					goto l1854
				l1855:
					position, tokenIndex = position1854, tokenIndex1854
					{
						position1861 := position
						if buffer[position] != rune('0') {
							goto l1852
						}
						position++
						{
							position1862, tokenIndex1862 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1863
							}
							position++
							goto l1862
						l1863:
							position, tokenIndex = position1862, tokenIndex1862
							if buffer[position] != rune('X') {
								goto l1852
							}
							position++
						}
					l1862:
						{
							position1864 := position
							if !_rules[rulehexdigit]() {
								goto l1852
							}
							if !_rules[rulehexdigit]() {
								goto l1852
							}
							if !_rules[rulehexdigit]() {
								goto l1852
							}
							if !_rules[rulehexdigit]() {
								goto l1852
							}
							add(rulePegText, position1864)
						}
						{
							add(ruleAction36, position)
						}
						add(rulehexWord0x, position1861)
					}
				}
			l1854:
				add(rulenn, position1853)
			}
			return true
		l1852:
			position, tokenIndex = position1852, tokenIndex1852
			return false
		},
		/* 42 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position1866, tokenIndex1866 := position, tokenIndex
			{
				position1867 := position
				{
					position1868, tokenIndex1868 := position, tokenIndex
					{
						position1870 := position
						{
							position1871 := position
							{
								position1872, tokenIndex1872 := position, tokenIndex
								{
									position1874, tokenIndex1874 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1875
									}
									position++
									goto l1874
								l1875:
									position, tokenIndex = position1874, tokenIndex1874
									if buffer[position] != rune('+') {
										goto l1872
									}
									position++
								}
							l1874:
								goto l1873
							l1872:
								position, tokenIndex = position1872, tokenIndex1872
							}
						l1873:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1869
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1869
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1869
									}
									position++
									break
								}
							}

						l1876:
							{
								position1877, tokenIndex1877 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1877
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1877
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1877
										}
										position++
										break
									}
								}

								goto l1876
							l1877:
								position, tokenIndex = position1877, tokenIndex1877
							}
							add(rulePegText, position1871)
						}
						{
							position1880, tokenIndex1880 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1881
							}
							position++
							goto l1880
						l1881:
							position, tokenIndex = position1880, tokenIndex1880
							if buffer[position] != rune('H') {
								goto l1869
							}
							position++
						}
					l1880:
						{
							add(ruleAction30, position)
						}
						add(rulesignedHexByteH, position1870)
					}
					goto l1868
				l1869:
					position, tokenIndex = position1868, tokenIndex1868
					{
						position1884 := position
						{
							position1885 := position
							{
								position1886, tokenIndex1886 := position, tokenIndex
								{
									position1888, tokenIndex1888 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1889
									}
									position++
									goto l1888
								l1889:
									position, tokenIndex = position1888, tokenIndex1888
									if buffer[position] != rune('+') {
										goto l1886
									}
									position++
								}
							l1888:
								goto l1887
							l1886:
								position, tokenIndex = position1886, tokenIndex1886
							}
						l1887:
							if buffer[position] != rune('0') {
								goto l1883
							}
							position++
							{
								position1890, tokenIndex1890 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1891
								}
								position++
								goto l1890
							l1891:
								position, tokenIndex = position1890, tokenIndex1890
								if buffer[position] != rune('X') {
									goto l1883
								}
								position++
							}
						l1890:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1883
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1883
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1883
									}
									position++
									break
								}
							}

						l1892:
							{
								position1893, tokenIndex1893 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1893
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1893
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1893
										}
										position++
										break
									}
								}

								goto l1892
							l1893:
								position, tokenIndex = position1893, tokenIndex1893
							}
							add(rulePegText, position1885)
						}
						{
							add(ruleAction31, position)
						}
						add(rulesignedHexByte0x, position1884)
					}
					goto l1868
				l1883:
					position, tokenIndex = position1868, tokenIndex1868
					{
						position1897 := position
						{
							position1898 := position
							{
								position1899, tokenIndex1899 := position, tokenIndex
								{
									position1901, tokenIndex1901 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1902
									}
									position++
									goto l1901
								l1902:
									position, tokenIndex = position1901, tokenIndex1901
									if buffer[position] != rune('+') {
										goto l1899
									}
									position++
								}
							l1901:
								goto l1900
							l1899:
								position, tokenIndex = position1899, tokenIndex1899
							}
						l1900:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1866
							}
							position++
						l1903:
							{
								position1904, tokenIndex1904 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1904
								}
								position++
								goto l1903
							l1904:
								position, tokenIndex = position1904, tokenIndex1904
							}
							add(rulePegText, position1898)
						}
						{
							add(ruleAction29, position)
						}
						add(rulesignedDecimalByte, position1897)
					}
				}
			l1868:
				add(ruledisp, position1867)
			}
			return true
		l1866:
			position, tokenIndex = position1866, tokenIndex1866
			return false
		},
		/* 43 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action29)> */
		nil,
		/* 44 signedHexByteH <- <(<(('-' / '+')? ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> ('h' / 'H') Action30)> */
		nil,
		/* 45 signedHexByte0x <- <(<(('-' / '+')? ('0' ('x' / 'X')) ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> Action31)> */
		nil,
		/* 46 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action32)> */
		nil,
		/* 47 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action33)> */
		nil,
		/* 48 decimalByte <- <(<[0-9]+> Action34)> */
		nil,
		/* 49 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action35)> */
		nil,
		/* 50 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action36)> */
		nil,
		/* 51 nn_contents <- <('(' nn ')' Action37)> */
		func() bool {
			position1914, tokenIndex1914 := position, tokenIndex
			{
				position1915 := position
				if buffer[position] != rune('(') {
					goto l1914
				}
				position++
				if !_rules[rulenn]() {
					goto l1914
				}
				if buffer[position] != rune(')') {
					goto l1914
				}
				position++
				{
					add(ruleAction37, position)
				}
				add(rulenn_contents, position1915)
			}
			return true
		l1914:
			position, tokenIndex = position1914, tokenIndex1914
			return false
		},
		/* 52 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 53 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action38)> */
		nil,
		/* 54 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action39)> */
		nil,
		/* 55 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action40)> */
		nil,
		/* 56 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action41)> */
		nil,
		/* 57 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action42)> */
		nil,
		/* 58 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action43)> */
		nil,
		/* 59 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action44)> */
		nil,
		/* 60 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action45)> */
		nil,
		/* 61 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 62 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 63 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 (sep Copy8)? Action46)> */
		nil,
		/* 64 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 (sep Copy8)? Action47)> */
		nil,
		/* 65 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action48)> */
		nil,
		/* 66 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 (sep Copy8)? Action49)> */
		nil,
		/* 67 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 (sep Copy8)? Action50)> */
		nil,
		/* 68 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 (sep Copy8)? Action51)> */
		nil,
		/* 69 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 (sep Copy8)? Action52)> */
		nil,
		/* 70 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action53)> */
		nil,
		/* 71 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action54)> */
		nil,
		/* 72 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 (sep Copy8)? Action55)> */
		nil,
		/* 73 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 (sep Copy8)? Action56)> */
		nil,
		/* 74 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 75 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action57)> */
		nil,
		/* 76 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action58)> */
		nil,
		/* 77 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action59)> */
		nil,
		/* 78 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action60)> */
		nil,
		/* 79 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action61)> */
		nil,
		/* 80 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action62)> */
		nil,
		/* 81 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action63)> */
		nil,
		/* 82 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action64)> */
		nil,
		/* 83 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action65)> */
		nil,
		/* 84 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action66)> */
		nil,
		/* 85 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action67)> */
		nil,
		/* 86 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action68)> */
		nil,
		/* 87 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action69)> */
		nil,
		/* 88 EDSimple <- <(Retn / Reti / Rrd / Im0 / Im1 / Im2 / ((&('I' | 'O' | 'i' | 'o') BlitIO) | (&('R' | 'r') Rld) | (&('N' | 'n') Neg) | (&('C' | 'L' | 'c' | 'l') Blit)))> */
		nil,
		/* 89 Neg <- <(<(('n' / 'N') ('e' / 'E') ('g' / 'G'))> Action70)> */
		nil,
		/* 90 Retn <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N'))> Action71)> */
		nil,
		/* 91 Reti <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('i' / 'I'))> Action72)> */
		nil,
		/* 92 Rrd <- <(<(('r' / 'R') ('r' / 'R') ('d' / 'D'))> Action73)> */
		nil,
		/* 93 Rld <- <(<(('r' / 'R') ('l' / 'L') ('d' / 'D'))> Action74)> */
		nil,
		/* 94 Im0 <- <(<(('i' / 'I') ('m' / 'M') ' ' '0')> Action75)> */
		nil,
		/* 95 Im1 <- <(<(('i' / 'I') ('m' / 'M') ' ' '1')> Action76)> */
		nil,
		/* 96 Im2 <- <(<(('i' / 'I') ('m' / 'M') ' ' '2')> Action77)> */
		nil,
		/* 97 Blit <- <(Ldir / Ldi / Cpir / Cpi / Lddr / Ldd / Cpdr / Cpd)> */
		nil,
		/* 98 BlitIO <- <(Inir / Ini / Otir / Outi / Indr / Ind / Otdr / Outd)> */
		nil,
		/* 99 Ldi <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I'))> Action78)> */
		nil,
		/* 100 Cpi <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I'))> Action79)> */
		nil,
		/* 101 Ini <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I'))> Action80)> */
		nil,
		/* 102 Outi <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('i' / 'I'))> Action81)> */
		nil,
		/* 103 Ldd <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D'))> Action82)> */
		nil,
		/* 104 Cpd <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D'))> Action83)> */
		nil,
		/* 105 Ind <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D'))> Action84)> */
		nil,
		/* 106 Outd <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('d' / 'D'))> Action85)> */
		nil,
		/* 107 Ldir <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I') ('r' / 'R'))> Action86)> */
		nil,
		/* 108 Cpir <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I') ('r' / 'R'))> Action87)> */
		nil,
		/* 109 Inir <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I') ('r' / 'R'))> Action88)> */
		nil,
		/* 110 Otir <- <(<(('o' / 'O') ('t' / 'T') ('i' / 'I') ('r' / 'R'))> Action89)> */
		nil,
		/* 111 Lddr <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D') ('r' / 'R'))> Action90)> */
		nil,
		/* 112 Cpdr <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D') ('r' / 'R'))> Action91)> */
		nil,
		/* 113 Indr <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D') ('r' / 'R'))> Action92)> */
		nil,
		/* 114 Otdr <- <(<(('o' / 'O') ('t' / 'T') ('d' / 'D') ('r' / 'R'))> Action93)> */
		nil,
		/* 115 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 116 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action94)> */
		nil,
		/* 117 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action95)> */
		nil,
		/* 118 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action96)> */
		nil,
		/* 119 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action97)> */
		nil,
		/* 120 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action98)> */
		nil,
		/* 121 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action99)> */
		nil,
		/* 122 IO <- <(IN / OUT)> */
		nil,
		/* 123 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action100)> */
		nil,
		/* 124 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action101)> */
		nil,
		/* 125 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position1990, tokenIndex1990 := position, tokenIndex
			{
				position1991 := position
				{
					position1992, tokenIndex1992 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l1993
					}
					position++
					{
						position1994, tokenIndex1994 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l1995
						}
						position++
						goto l1994
					l1995:
						position, tokenIndex = position1994, tokenIndex1994
						if buffer[position] != rune('C') {
							goto l1993
						}
						position++
					}
				l1994:
					if buffer[position] != rune(')') {
						goto l1993
					}
					position++
					goto l1992
				l1993:
					position, tokenIndex = position1992, tokenIndex1992
					if buffer[position] != rune('(') {
						goto l1990
					}
					position++
					if !_rules[rulen]() {
						goto l1990
					}
					if buffer[position] != rune(')') {
						goto l1990
					}
					position++
				}
			l1992:
				add(rulePort, position1991)
			}
			return true
		l1990:
			position, tokenIndex = position1990, tokenIndex1990
			return false
		},
		/* 126 sep <- <(ws? ',' ws?)> */
		func() bool {
			position1996, tokenIndex1996 := position, tokenIndex
			{
				position1997 := position
				{
					position1998, tokenIndex1998 := position, tokenIndex
					if !_rules[rulews]() {
						goto l1998
					}
					goto l1999
				l1998:
					position, tokenIndex = position1998, tokenIndex1998
				}
			l1999:
				if buffer[position] != rune(',') {
					goto l1996
				}
				position++
				{
					position2000, tokenIndex2000 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2000
					}
					goto l2001
				l2000:
					position, tokenIndex = position2000, tokenIndex2000
				}
			l2001:
				add(rulesep, position1997)
			}
			return true
		l1996:
			position, tokenIndex = position1996, tokenIndex1996
			return false
		},
		/* 127 ws <- <' '+> */
		func() bool {
			position2002, tokenIndex2002 := position, tokenIndex
			{
				position2003 := position
				if buffer[position] != rune(' ') {
					goto l2002
				}
				position++
			l2004:
				{
					position2005, tokenIndex2005 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2005
					}
					position++
					goto l2004
				l2005:
					position, tokenIndex = position2005, tokenIndex2005
				}
				add(rulews, position2003)
			}
			return true
		l2002:
			position, tokenIndex = position2002, tokenIndex2002
			return false
		},
		/* 128 A <- <('a' / 'A')> */
		nil,
		/* 129 F <- <('f' / 'F')> */
		nil,
		/* 130 B <- <('b' / 'B')> */
		nil,
		/* 131 C <- <('c' / 'C')> */
		nil,
		/* 132 D <- <('d' / 'D')> */
		nil,
		/* 133 E <- <('e' / 'E')> */
		nil,
		/* 134 H <- <('h' / 'H')> */
		nil,
		/* 135 L <- <('l' / 'L')> */
		nil,
		/* 136 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 137 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 138 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 139 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 140 I <- <('i' / 'I')> */
		nil,
		/* 141 R <- <('r' / 'R')> */
		nil,
		/* 142 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 143 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 144 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 145 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 146 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 147 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 148 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 149 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 150 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position2028, tokenIndex2028 := position, tokenIndex
			{
				position2029 := position
				{
					position2030, tokenIndex2030 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2031
					}
					position++
					goto l2030
				l2031:
					position, tokenIndex = position2030, tokenIndex2030
					{
						position2032, tokenIndex2032 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2033
						}
						position++
						goto l2032
					l2033:
						position, tokenIndex = position2032, tokenIndex2032
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2028
						}
						position++
					}
				l2032:
				}
			l2030:
				add(rulehexdigit, position2029)
			}
			return true
		l2028:
			position, tokenIndex = position2028, tokenIndex2028
			return false
		},
		/* 151 octaldigit <- <(<[0-7]> Action102)> */
		func() bool {
			position2034, tokenIndex2034 := position, tokenIndex
			{
				position2035 := position
				{
					position2036 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2034
					}
					position++
					add(rulePegText, position2036)
				}
				{
					add(ruleAction102, position)
				}
				add(ruleoctaldigit, position2035)
			}
			return true
		l2034:
			position, tokenIndex = position2034, tokenIndex2034
			return false
		},
		/* 152 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2038, tokenIndex2038 := position, tokenIndex
			{
				position2039 := position
				{
					position2040, tokenIndex2040 := position, tokenIndex
					{
						position2042 := position
						{
							position2043, tokenIndex2043 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2044
							}
							position++
							goto l2043
						l2044:
							position, tokenIndex = position2043, tokenIndex2043
							if buffer[position] != rune('N') {
								goto l2041
							}
							position++
						}
					l2043:
						{
							position2045, tokenIndex2045 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2046
							}
							position++
							goto l2045
						l2046:
							position, tokenIndex = position2045, tokenIndex2045
							if buffer[position] != rune('Z') {
								goto l2041
							}
							position++
						}
					l2045:
						{
							add(ruleAction103, position)
						}
						add(ruleFT_NZ, position2042)
					}
					goto l2040
				l2041:
					position, tokenIndex = position2040, tokenIndex2040
					{
						position2049 := position
						{
							position2050, tokenIndex2050 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2051
							}
							position++
							goto l2050
						l2051:
							position, tokenIndex = position2050, tokenIndex2050
							if buffer[position] != rune('P') {
								goto l2048
							}
							position++
						}
					l2050:
						{
							position2052, tokenIndex2052 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2053
							}
							position++
							goto l2052
						l2053:
							position, tokenIndex = position2052, tokenIndex2052
							if buffer[position] != rune('O') {
								goto l2048
							}
							position++
						}
					l2052:
						{
							add(ruleAction107, position)
						}
						add(ruleFT_PO, position2049)
					}
					goto l2040
				l2048:
					position, tokenIndex = position2040, tokenIndex2040
					{
						position2056 := position
						{
							position2057, tokenIndex2057 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2058
							}
							position++
							goto l2057
						l2058:
							position, tokenIndex = position2057, tokenIndex2057
							if buffer[position] != rune('P') {
								goto l2055
							}
							position++
						}
					l2057:
						{
							position2059, tokenIndex2059 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2060
							}
							position++
							goto l2059
						l2060:
							position, tokenIndex = position2059, tokenIndex2059
							if buffer[position] != rune('E') {
								goto l2055
							}
							position++
						}
					l2059:
						{
							add(ruleAction108, position)
						}
						add(ruleFT_PE, position2056)
					}
					goto l2040
				l2055:
					position, tokenIndex = position2040, tokenIndex2040
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2063 := position
								{
									position2064, tokenIndex2064 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2065
									}
									position++
									goto l2064
								l2065:
									position, tokenIndex = position2064, tokenIndex2064
									if buffer[position] != rune('M') {
										goto l2038
									}
									position++
								}
							l2064:
								{
									add(ruleAction110, position)
								}
								add(ruleFT_M, position2063)
							}
							break
						case 'P', 'p':
							{
								position2067 := position
								{
									position2068, tokenIndex2068 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2069
									}
									position++
									goto l2068
								l2069:
									position, tokenIndex = position2068, tokenIndex2068
									if buffer[position] != rune('P') {
										goto l2038
									}
									position++
								}
							l2068:
								{
									add(ruleAction109, position)
								}
								add(ruleFT_P, position2067)
							}
							break
						case 'C', 'c':
							{
								position2071 := position
								{
									position2072, tokenIndex2072 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2073
									}
									position++
									goto l2072
								l2073:
									position, tokenIndex = position2072, tokenIndex2072
									if buffer[position] != rune('C') {
										goto l2038
									}
									position++
								}
							l2072:
								{
									add(ruleAction106, position)
								}
								add(ruleFT_C, position2071)
							}
							break
						case 'N', 'n':
							{
								position2075 := position
								{
									position2076, tokenIndex2076 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2077
									}
									position++
									goto l2076
								l2077:
									position, tokenIndex = position2076, tokenIndex2076
									if buffer[position] != rune('N') {
										goto l2038
									}
									position++
								}
							l2076:
								{
									position2078, tokenIndex2078 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2079
									}
									position++
									goto l2078
								l2079:
									position, tokenIndex = position2078, tokenIndex2078
									if buffer[position] != rune('C') {
										goto l2038
									}
									position++
								}
							l2078:
								{
									add(ruleAction105, position)
								}
								add(ruleFT_NC, position2075)
							}
							break
						default:
							{
								position2081 := position
								{
									position2082, tokenIndex2082 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2083
									}
									position++
									goto l2082
								l2083:
									position, tokenIndex = position2082, tokenIndex2082
									if buffer[position] != rune('Z') {
										goto l2038
									}
									position++
								}
							l2082:
								{
									add(ruleAction104, position)
								}
								add(ruleFT_Z, position2081)
							}
							break
						}
					}

				}
			l2040:
				add(rulecc, position2039)
			}
			return true
		l2038:
			position, tokenIndex = position2038, tokenIndex2038
			return false
		},
		/* 153 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action103)> */
		nil,
		/* 154 FT_Z <- <(('z' / 'Z') Action104)> */
		nil,
		/* 155 FT_NC <- <(('n' / 'N') ('c' / 'C') Action105)> */
		nil,
		/* 156 FT_C <- <(('c' / 'C') Action106)> */
		nil,
		/* 157 FT_PO <- <(('p' / 'P') ('o' / 'O') Action107)> */
		nil,
		/* 158 FT_PE <- <(('p' / 'P') ('e' / 'E') Action108)> */
		nil,
		/* 159 FT_P <- <(('p' / 'P') Action109)> */
		nil,
		/* 160 FT_M <- <(('m' / 'M') Action110)> */
		nil,
		/* 162 Action0 <- <{ p.Emit() }> */
		nil,
		/* 163 Action1 <- <{ p.LD8() }> */
		nil,
		/* 164 Action2 <- <{ p.LD16() }> */
		nil,
		/* 165 Action3 <- <{ p.Push() }> */
		nil,
		/* 166 Action4 <- <{ p.Pop() }> */
		nil,
		/* 167 Action5 <- <{ p.Ex() }> */
		nil,
		/* 168 Action6 <- <{ p.Inc8() }> */
		nil,
		/* 169 Action7 <- <{ p.Inc8() }> */
		nil,
		/* 170 Action8 <- <{ p.Inc16() }> */
		nil,
		/* 171 Action9 <- <{ p.Dec8() }> */
		nil,
		/* 172 Action10 <- <{ p.Dec8() }> */
		nil,
		/* 173 Action11 <- <{ p.Dec16() }> */
		nil,
		/* 174 Action12 <- <{ p.Add16() }> */
		nil,
		/* 175 Action13 <- <{ p.Adc16() }> */
		nil,
		/* 176 Action14 <- <{ p.Sbc16() }> */
		nil,
		/* 177 Action15 <- <{ p.Dst8() }> */
		nil,
		/* 178 Action16 <- <{ p.Src8() }> */
		nil,
		/* 179 Action17 <- <{ p.Loc8() }> */
		nil,
		/* 180 Action18 <- <{ p.Copy8() }> */
		nil,
		/* 181 Action19 <- <{ p.Loc8() }> */
		nil,
		nil,
		/* 183 Action20 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 184 Action21 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 185 Action22 <- <{ p.Dst16() }> */
		nil,
		/* 186 Action23 <- <{ p.Src16() }> */
		nil,
		/* 187 Action24 <- <{ p.Loc16() }> */
		nil,
		/* 188 Action25 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 189 Action26 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 190 Action27 <- <{ p.R16Contents() }> */
		nil,
		/* 191 Action28 <- <{ p.IR16Contents() }> */
		nil,
		/* 192 Action29 <- <{ p.DispDecimal(buffer[begin:end]) }> */
		nil,
		/* 193 Action30 <- <{ p.DispHex(buffer[begin:end]) }> */
		nil,
		/* 194 Action31 <- <{ p.Disp0xHex(buffer[begin:end]) }> */
		nil,
		/* 195 Action32 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 196 Action33 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 197 Action34 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 198 Action35 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 199 Action36 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 200 Action37 <- <{ p.NNContents() }> */
		nil,
		/* 201 Action38 <- <{ p.Accum("ADD") }> */
		nil,
		/* 202 Action39 <- <{ p.Accum("ADC") }> */
		nil,
		/* 203 Action40 <- <{ p.Accum("SUB") }> */
		nil,
		/* 204 Action41 <- <{ p.Accum("SBC") }> */
		nil,
		/* 205 Action42 <- <{ p.Accum("AND") }> */
		nil,
		/* 206 Action43 <- <{ p.Accum("XOR") }> */
		nil,
		/* 207 Action44 <- <{ p.Accum("OR") }> */
		nil,
		/* 208 Action45 <- <{ p.Accum("CP") }> */
		nil,
		/* 209 Action46 <- <{ p.Rot("RLC") }> */
		nil,
		/* 210 Action47 <- <{ p.Rot("RRC") }> */
		nil,
		/* 211 Action48 <- <{ p.Rot("RL") }> */
		nil,
		/* 212 Action49 <- <{ p.Rot("RR") }> */
		nil,
		/* 213 Action50 <- <{ p.Rot("SLA") }> */
		nil,
		/* 214 Action51 <- <{ p.Rot("SRA") }> */
		nil,
		/* 215 Action52 <- <{ p.Rot("SLL") }> */
		nil,
		/* 216 Action53 <- <{ p.Rot("SRL") }> */
		nil,
		/* 217 Action54 <- <{ p.Bit() }> */
		nil,
		/* 218 Action55 <- <{ p.Res() }> */
		nil,
		/* 219 Action56 <- <{ p.Set() }> */
		nil,
		/* 220 Action57 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 221 Action58 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 222 Action59 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 223 Action60 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 224 Action61 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 225 Action62 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 226 Action63 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 227 Action64 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 228 Action65 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 229 Action66 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 230 Action67 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 231 Action68 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 232 Action69 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 233 Action70 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 234 Action71 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 235 Action72 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 236 Action73 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 237 Action74 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 238 Action75 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 239 Action76 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 240 Action77 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 241 Action78 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 242 Action79 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 243 Action80 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 244 Action81 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 245 Action82 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 246 Action83 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 247 Action84 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 248 Action85 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 249 Action86 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 250 Action87 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 251 Action88 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 252 Action89 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 253 Action90 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 254 Action91 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 255 Action92 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 256 Action93 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 257 Action94 <- <{ p.Rst() }> */
		nil,
		/* 258 Action95 <- <{ p.Call() }> */
		nil,
		/* 259 Action96 <- <{ p.Ret() }> */
		nil,
		/* 260 Action97 <- <{ p.Jp() }> */
		nil,
		/* 261 Action98 <- <{ p.Jr() }> */
		nil,
		/* 262 Action99 <- <{ p.Djnz() }> */
		nil,
		/* 263 Action100 <- <{ p.In() }> */
		nil,
		/* 264 Action101 <- <{ p.Out() }> */
		nil,
		/* 265 Action102 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 266 Action103 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 267 Action104 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 268 Action105 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 269 Action106 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 270 Action107 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 271 Action108 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 272 Action109 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 273 Action110 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
