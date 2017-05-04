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
	ruleLine
	ruleLabel
	rulealphanum
	rulealpha
	rulenum
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
	rulePegText
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
	ruleAction111
)

var rul3s = [...]string{
	"Unknown",
	"Program",
	"Line",
	"Label",
	"alphanum",
	"alpha",
	"num",
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
	"PegText",
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
	"Action111",
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
	rules  [278]func() bool
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
			p.Label(buffer[begin:end])
		case ruleAction2:
			p.LD8()
		case ruleAction3:
			p.LD16()
		case ruleAction4:
			p.Push()
		case ruleAction5:
			p.Pop()
		case ruleAction6:
			p.Ex()
		case ruleAction7:
			p.Inc8()
		case ruleAction8:
			p.Inc8()
		case ruleAction9:
			p.Inc16()
		case ruleAction10:
			p.Dec8()
		case ruleAction11:
			p.Dec8()
		case ruleAction12:
			p.Dec16()
		case ruleAction13:
			p.Add16()
		case ruleAction14:
			p.Adc16()
		case ruleAction15:
			p.Sbc16()
		case ruleAction16:
			p.Dst8()
		case ruleAction17:
			p.Src8()
		case ruleAction18:
			p.Loc8()
		case ruleAction19:
			p.Copy8()
		case ruleAction20:
			p.Loc8()
		case ruleAction21:
			p.R8(buffer[begin:end])
		case ruleAction22:
			p.R8(buffer[begin:end])
		case ruleAction23:
			p.Dst16()
		case ruleAction24:
			p.Src16()
		case ruleAction25:
			p.Loc16()
		case ruleAction26:
			p.R16(buffer[begin:end])
		case ruleAction27:
			p.R16(buffer[begin:end])
		case ruleAction28:
			p.R16Contents()
		case ruleAction29:
			p.IR16Contents()
		case ruleAction30:
			p.DispDecimal(buffer[begin:end])
		case ruleAction31:
			p.DispHex(buffer[begin:end])
		case ruleAction32:
			p.Disp0xHex(buffer[begin:end])
		case ruleAction33:
			p.Nhex(buffer[begin:end])
		case ruleAction34:
			p.Nhex(buffer[begin:end])
		case ruleAction35:
			p.Ndec(buffer[begin:end])
		case ruleAction36:
			p.NNhex(buffer[begin:end])
		case ruleAction37:
			p.NNhex(buffer[begin:end])
		case ruleAction38:
			p.NNContents()
		case ruleAction39:
			p.Accum("ADD")
		case ruleAction40:
			p.Accum("ADC")
		case ruleAction41:
			p.Accum("SUB")
		case ruleAction42:
			p.Accum("SBC")
		case ruleAction43:
			p.Accum("AND")
		case ruleAction44:
			p.Accum("XOR")
		case ruleAction45:
			p.Accum("OR")
		case ruleAction46:
			p.Accum("CP")
		case ruleAction47:
			p.Rot("RLC")
		case ruleAction48:
			p.Rot("RRC")
		case ruleAction49:
			p.Rot("RL")
		case ruleAction50:
			p.Rot("RR")
		case ruleAction51:
			p.Rot("SLA")
		case ruleAction52:
			p.Rot("SRA")
		case ruleAction53:
			p.Rot("SLL")
		case ruleAction54:
			p.Rot("SRL")
		case ruleAction55:
			p.Bit()
		case ruleAction56:
			p.Res()
		case ruleAction57:
			p.Set()
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
			p.Simple(buffer[begin:end])
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
			p.EDSimple(buffer[begin:end])
		case ruleAction95:
			p.Rst()
		case ruleAction96:
			p.Call()
		case ruleAction97:
			p.Ret()
		case ruleAction98:
			p.Jp()
		case ruleAction99:
			p.Jr()
		case ruleAction100:
			p.Djnz()
		case ruleAction101:
			p.In()
		case ruleAction102:
			p.Out()
		case ruleAction103:
			p.ODigit(buffer[begin:end])
		case ruleAction104:
			p.Conditional(Not{FT_Z})
		case ruleAction105:
			p.Conditional(FT_Z)
		case ruleAction106:
			p.Conditional(Not{FT_C})
		case ruleAction107:
			p.Conditional(FT_C)
		case ruleAction108:
			p.Conditional(FT_PO)
		case ruleAction109:
			p.Conditional(FT_PE)
		case ruleAction110:
			p.Conditional(FT_P)
		case ruleAction111:
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
		/* 0 Program <- <Line+> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position4 := position
					{
						position5, tokenIndex5 := position, tokenIndex
						{
							position7 := position
							{
								position8 := position
								if !_rules[rulealpha]() {
									goto l5
								}
							l9:
								{
									position10, tokenIndex10 := position, tokenIndex
									{
										position11 := position
										{
											position12, tokenIndex12 := position, tokenIndex
											if !_rules[rulealpha]() {
												goto l13
											}
											goto l12
										l13:
											position, tokenIndex = position12, tokenIndex12
											{
												position14 := position
												if c := buffer[position]; c < rune('0') || c > rune('9') {
													goto l10
												}
												position++
												add(rulenum, position14)
											}
										}
									l12:
										add(rulealphanum, position11)
									}
									goto l9
								l10:
									position, tokenIndex = position10, tokenIndex10
								}
								add(rulePegText, position8)
							}
							if buffer[position] != rune(':') {
								goto l5
							}
							position++
							{
								add(ruleAction1, position)
							}
							add(ruleLabel, position7)
						}
						goto l6
					l5:
						position, tokenIndex = position5, tokenIndex5
					}
				l6:
				l16:
					{
						position17, tokenIndex17 := position, tokenIndex
						if !_rules[rulews]() {
							goto l17
						}
						goto l16
					l17:
						position, tokenIndex = position17, tokenIndex17
					}
					{
						position18, tokenIndex18 := position, tokenIndex
						{
							position20 := position
							{
								position21, tokenIndex21 := position, tokenIndex
								{
									position23 := position
									{
										position24, tokenIndex24 := position, tokenIndex
										{
											position26 := position
											{
												position27, tokenIndex27 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l28
												}
												position++
												goto l27
											l28:
												position, tokenIndex = position27, tokenIndex27
												if buffer[position] != rune('P') {
													goto l25
												}
												position++
											}
										l27:
											{
												position29, tokenIndex29 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l30
												}
												position++
												goto l29
											l30:
												position, tokenIndex = position29, tokenIndex29
												if buffer[position] != rune('U') {
													goto l25
												}
												position++
											}
										l29:
											{
												position31, tokenIndex31 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l32
												}
												position++
												goto l31
											l32:
												position, tokenIndex = position31, tokenIndex31
												if buffer[position] != rune('S') {
													goto l25
												}
												position++
											}
										l31:
											{
												position33, tokenIndex33 := position, tokenIndex
												if buffer[position] != rune('h') {
													goto l34
												}
												position++
												goto l33
											l34:
												position, tokenIndex = position33, tokenIndex33
												if buffer[position] != rune('H') {
													goto l25
												}
												position++
											}
										l33:
											if !_rules[rulews]() {
												goto l25
											}
											if !_rules[ruleSrc16]() {
												goto l25
											}
											{
												add(ruleAction4, position)
											}
											add(rulePush, position26)
										}
										goto l24
									l25:
										position, tokenIndex = position24, tokenIndex24
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position37 := position
													{
														position38, tokenIndex38 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l39
														}
														position++
														goto l38
													l39:
														position, tokenIndex = position38, tokenIndex38
														if buffer[position] != rune('E') {
															goto l22
														}
														position++
													}
												l38:
													{
														position40, tokenIndex40 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l41
														}
														position++
														goto l40
													l41:
														position, tokenIndex = position40, tokenIndex40
														if buffer[position] != rune('X') {
															goto l22
														}
														position++
													}
												l40:
													if !_rules[rulews]() {
														goto l22
													}
													if !_rules[ruleDst16]() {
														goto l22
													}
													if !_rules[rulesep]() {
														goto l22
													}
													if !_rules[ruleSrc16]() {
														goto l22
													}
													{
														add(ruleAction6, position)
													}
													add(ruleEx, position37)
												}
												break
											case 'P', 'p':
												{
													position43 := position
													{
														position44, tokenIndex44 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l45
														}
														position++
														goto l44
													l45:
														position, tokenIndex = position44, tokenIndex44
														if buffer[position] != rune('P') {
															goto l22
														}
														position++
													}
												l44:
													{
														position46, tokenIndex46 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l47
														}
														position++
														goto l46
													l47:
														position, tokenIndex = position46, tokenIndex46
														if buffer[position] != rune('O') {
															goto l22
														}
														position++
													}
												l46:
													{
														position48, tokenIndex48 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l49
														}
														position++
														goto l48
													l49:
														position, tokenIndex = position48, tokenIndex48
														if buffer[position] != rune('P') {
															goto l22
														}
														position++
													}
												l48:
													if !_rules[rulews]() {
														goto l22
													}
													if !_rules[ruleDst16]() {
														goto l22
													}
													{
														add(ruleAction5, position)
													}
													add(rulePop, position43)
												}
												break
											default:
												{
													position51 := position
													{
														position52, tokenIndex52 := position, tokenIndex
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
																	goto l53
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
																	goto l53
																}
																position++
															}
														l57:
															if !_rules[rulews]() {
																goto l53
															}
															if !_rules[ruleDst16]() {
																goto l53
															}
															if !_rules[rulesep]() {
																goto l53
															}
															if !_rules[ruleSrc16]() {
																goto l53
															}
															{
																add(ruleAction3, position)
															}
															add(ruleLoad16, position54)
														}
														goto l52
													l53:
														position, tokenIndex = position52, tokenIndex52
														{
															position60 := position
															{
																position61, tokenIndex61 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l62
																}
																position++
																goto l61
															l62:
																position, tokenIndex = position61, tokenIndex61
																if buffer[position] != rune('L') {
																	goto l22
																}
																position++
															}
														l61:
															{
																position63, tokenIndex63 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l64
																}
																position++
																goto l63
															l64:
																position, tokenIndex = position63, tokenIndex63
																if buffer[position] != rune('D') {
																	goto l22
																}
																position++
															}
														l63:
															if !_rules[rulews]() {
																goto l22
															}
															{
																position65 := position
																{
																	position66, tokenIndex66 := position, tokenIndex
																	if !_rules[ruleReg8]() {
																		goto l67
																	}
																	goto l66
																l67:
																	position, tokenIndex = position66, tokenIndex66
																	if !_rules[ruleReg16Contents]() {
																		goto l68
																	}
																	goto l66
																l68:
																	position, tokenIndex = position66, tokenIndex66
																	if !_rules[rulenn_contents]() {
																		goto l22
																	}
																}
															l66:
																{
																	add(ruleAction16, position)
																}
																add(ruleDst8, position65)
															}
															if !_rules[rulesep]() {
																goto l22
															}
															if !_rules[ruleSrc8]() {
																goto l22
															}
															{
																add(ruleAction2, position)
															}
															add(ruleLoad8, position60)
														}
													}
												l52:
													add(ruleLoad, position51)
												}
												break
											}
										}

									}
								l24:
									add(ruleAssignment, position23)
								}
								goto l21
							l22:
								position, tokenIndex = position21, tokenIndex21
								{
									position72 := position
									{
										position73, tokenIndex73 := position, tokenIndex
										{
											position75 := position
											{
												position76, tokenIndex76 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l77
												}
												position++
												goto l76
											l77:
												position, tokenIndex = position76, tokenIndex76
												if buffer[position] != rune('I') {
													goto l74
												}
												position++
											}
										l76:
											{
												position78, tokenIndex78 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l79
												}
												position++
												goto l78
											l79:
												position, tokenIndex = position78, tokenIndex78
												if buffer[position] != rune('N') {
													goto l74
												}
												position++
											}
										l78:
											{
												position80, tokenIndex80 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l81
												}
												position++
												goto l80
											l81:
												position, tokenIndex = position80, tokenIndex80
												if buffer[position] != rune('C') {
													goto l74
												}
												position++
											}
										l80:
											if !_rules[rulews]() {
												goto l74
											}
											if !_rules[ruleILoc8]() {
												goto l74
											}
											{
												add(ruleAction7, position)
											}
											add(ruleInc16Indexed8, position75)
										}
										goto l73
									l74:
										position, tokenIndex = position73, tokenIndex73
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
											add(ruleInc16, position84)
										}
										goto l73
									l83:
										position, tokenIndex = position73, tokenIndex73
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
													goto l71
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
													goto l71
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
													goto l71
												}
												position++
											}
										l97:
											if !_rules[rulews]() {
												goto l71
											}
											if !_rules[ruleLoc8]() {
												goto l71
											}
											{
												add(ruleAction8, position)
											}
											add(ruleInc8, position92)
										}
									}
								l73:
									add(ruleInc, position72)
								}
								goto l21
							l71:
								position, tokenIndex = position21, tokenIndex21
								{
									position101 := position
									{
										position102, tokenIndex102 := position, tokenIndex
										{
											position104 := position
											{
												position105, tokenIndex105 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l106
												}
												position++
												goto l105
											l106:
												position, tokenIndex = position105, tokenIndex105
												if buffer[position] != rune('D') {
													goto l103
												}
												position++
											}
										l105:
											{
												position107, tokenIndex107 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l108
												}
												position++
												goto l107
											l108:
												position, tokenIndex = position107, tokenIndex107
												if buffer[position] != rune('E') {
													goto l103
												}
												position++
											}
										l107:
											{
												position109, tokenIndex109 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l110
												}
												position++
												goto l109
											l110:
												position, tokenIndex = position109, tokenIndex109
												if buffer[position] != rune('C') {
													goto l103
												}
												position++
											}
										l109:
											if !_rules[rulews]() {
												goto l103
											}
											if !_rules[ruleILoc8]() {
												goto l103
											}
											{
												add(ruleAction10, position)
											}
											add(ruleDec16Indexed8, position104)
										}
										goto l102
									l103:
										position, tokenIndex = position102, tokenIndex102
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
													goto l112
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
													goto l112
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
													goto l112
												}
												position++
											}
										l118:
											if !_rules[rulews]() {
												goto l112
											}
											if !_rules[ruleLoc16]() {
												goto l112
											}
											{
												add(ruleAction12, position)
											}
											add(ruleDec16, position113)
										}
										goto l102
									l112:
										position, tokenIndex = position102, tokenIndex102
										{
											position121 := position
											{
												position122, tokenIndex122 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l123
												}
												position++
												goto l122
											l123:
												position, tokenIndex = position122, tokenIndex122
												if buffer[position] != rune('D') {
													goto l100
												}
												position++
											}
										l122:
											{
												position124, tokenIndex124 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l125
												}
												position++
												goto l124
											l125:
												position, tokenIndex = position124, tokenIndex124
												if buffer[position] != rune('E') {
													goto l100
												}
												position++
											}
										l124:
											{
												position126, tokenIndex126 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l127
												}
												position++
												goto l126
											l127:
												position, tokenIndex = position126, tokenIndex126
												if buffer[position] != rune('C') {
													goto l100
												}
												position++
											}
										l126:
											if !_rules[rulews]() {
												goto l100
											}
											if !_rules[ruleLoc8]() {
												goto l100
											}
											{
												add(ruleAction11, position)
											}
											add(ruleDec8, position121)
										}
									}
								l102:
									add(ruleDec, position101)
								}
								goto l21
							l100:
								position, tokenIndex = position21, tokenIndex21
								{
									position130 := position
									{
										position131, tokenIndex131 := position, tokenIndex
										{
											position133 := position
											{
												position134, tokenIndex134 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l135
												}
												position++
												goto l134
											l135:
												position, tokenIndex = position134, tokenIndex134
												if buffer[position] != rune('A') {
													goto l132
												}
												position++
											}
										l134:
											{
												position136, tokenIndex136 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l137
												}
												position++
												goto l136
											l137:
												position, tokenIndex = position136, tokenIndex136
												if buffer[position] != rune('D') {
													goto l132
												}
												position++
											}
										l136:
											{
												position138, tokenIndex138 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l139
												}
												position++
												goto l138
											l139:
												position, tokenIndex = position138, tokenIndex138
												if buffer[position] != rune('D') {
													goto l132
												}
												position++
											}
										l138:
											if !_rules[rulews]() {
												goto l132
											}
											if !_rules[ruleDst16]() {
												goto l132
											}
											if !_rules[rulesep]() {
												goto l132
											}
											if !_rules[ruleSrc16]() {
												goto l132
											}
											{
												add(ruleAction13, position)
											}
											add(ruleAdd16, position133)
										}
										goto l131
									l132:
										position, tokenIndex = position131, tokenIndex131
										{
											position142 := position
											{
												position143, tokenIndex143 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l144
												}
												position++
												goto l143
											l144:
												position, tokenIndex = position143, tokenIndex143
												if buffer[position] != rune('A') {
													goto l141
												}
												position++
											}
										l143:
											{
												position145, tokenIndex145 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l146
												}
												position++
												goto l145
											l146:
												position, tokenIndex = position145, tokenIndex145
												if buffer[position] != rune('D') {
													goto l141
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
													goto l141
												}
												position++
											}
										l147:
											if !_rules[rulews]() {
												goto l141
											}
											if !_rules[ruleDst16]() {
												goto l141
											}
											if !_rules[rulesep]() {
												goto l141
											}
											if !_rules[ruleSrc16]() {
												goto l141
											}
											{
												add(ruleAction14, position)
											}
											add(ruleAdc16, position142)
										}
										goto l131
									l141:
										position, tokenIndex = position131, tokenIndex131
										{
											position150 := position
											{
												position151, tokenIndex151 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l152
												}
												position++
												goto l151
											l152:
												position, tokenIndex = position151, tokenIndex151
												if buffer[position] != rune('S') {
													goto l129
												}
												position++
											}
										l151:
											{
												position153, tokenIndex153 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l154
												}
												position++
												goto l153
											l154:
												position, tokenIndex = position153, tokenIndex153
												if buffer[position] != rune('B') {
													goto l129
												}
												position++
											}
										l153:
											{
												position155, tokenIndex155 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l156
												}
												position++
												goto l155
											l156:
												position, tokenIndex = position155, tokenIndex155
												if buffer[position] != rune('C') {
													goto l129
												}
												position++
											}
										l155:
											if !_rules[rulews]() {
												goto l129
											}
											if !_rules[ruleDst16]() {
												goto l129
											}
											if !_rules[rulesep]() {
												goto l129
											}
											if !_rules[ruleSrc16]() {
												goto l129
											}
											{
												add(ruleAction15, position)
											}
											add(ruleSbc16, position150)
										}
									}
								l131:
									add(ruleAlu16, position130)
								}
								goto l21
							l129:
								position, tokenIndex = position21, tokenIndex21
								{
									position159 := position
									{
										position160, tokenIndex160 := position, tokenIndex
										{
											position162 := position
											{
												position163, tokenIndex163 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l164
												}
												position++
												goto l163
											l164:
												position, tokenIndex = position163, tokenIndex163
												if buffer[position] != rune('A') {
													goto l161
												}
												position++
											}
										l163:
											{
												position165, tokenIndex165 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l166
												}
												position++
												goto l165
											l166:
												position, tokenIndex = position165, tokenIndex165
												if buffer[position] != rune('D') {
													goto l161
												}
												position++
											}
										l165:
											{
												position167, tokenIndex167 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l168
												}
												position++
												goto l167
											l168:
												position, tokenIndex = position167, tokenIndex167
												if buffer[position] != rune('D') {
													goto l161
												}
												position++
											}
										l167:
											if !_rules[rulews]() {
												goto l161
											}
											{
												position169, tokenIndex169 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l170
												}
												position++
												goto l169
											l170:
												position, tokenIndex = position169, tokenIndex169
												if buffer[position] != rune('A') {
													goto l161
												}
												position++
											}
										l169:
											if !_rules[rulesep]() {
												goto l161
											}
											if !_rules[ruleSrc8]() {
												goto l161
											}
											{
												add(ruleAction39, position)
											}
											add(ruleAdd, position162)
										}
										goto l160
									l161:
										position, tokenIndex = position160, tokenIndex160
										{
											position173 := position
											{
												position174, tokenIndex174 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l175
												}
												position++
												goto l174
											l175:
												position, tokenIndex = position174, tokenIndex174
												if buffer[position] != rune('A') {
													goto l172
												}
												position++
											}
										l174:
											{
												position176, tokenIndex176 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l177
												}
												position++
												goto l176
											l177:
												position, tokenIndex = position176, tokenIndex176
												if buffer[position] != rune('D') {
													goto l172
												}
												position++
											}
										l176:
											{
												position178, tokenIndex178 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l179
												}
												position++
												goto l178
											l179:
												position, tokenIndex = position178, tokenIndex178
												if buffer[position] != rune('C') {
													goto l172
												}
												position++
											}
										l178:
											if !_rules[rulews]() {
												goto l172
											}
											{
												position180, tokenIndex180 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l181
												}
												position++
												goto l180
											l181:
												position, tokenIndex = position180, tokenIndex180
												if buffer[position] != rune('A') {
													goto l172
												}
												position++
											}
										l180:
											if !_rules[rulesep]() {
												goto l172
											}
											if !_rules[ruleSrc8]() {
												goto l172
											}
											{
												add(ruleAction40, position)
											}
											add(ruleAdc, position173)
										}
										goto l160
									l172:
										position, tokenIndex = position160, tokenIndex160
										{
											position184 := position
											{
												position185, tokenIndex185 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l186
												}
												position++
												goto l185
											l186:
												position, tokenIndex = position185, tokenIndex185
												if buffer[position] != rune('S') {
													goto l183
												}
												position++
											}
										l185:
											{
												position187, tokenIndex187 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l188
												}
												position++
												goto l187
											l188:
												position, tokenIndex = position187, tokenIndex187
												if buffer[position] != rune('U') {
													goto l183
												}
												position++
											}
										l187:
											{
												position189, tokenIndex189 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l190
												}
												position++
												goto l189
											l190:
												position, tokenIndex = position189, tokenIndex189
												if buffer[position] != rune('B') {
													goto l183
												}
												position++
											}
										l189:
											if !_rules[rulews]() {
												goto l183
											}
											if !_rules[ruleSrc8]() {
												goto l183
											}
											{
												add(ruleAction41, position)
											}
											add(ruleSub, position184)
										}
										goto l160
									l183:
										position, tokenIndex = position160, tokenIndex160
										{
											switch buffer[position] {
											case 'C', 'c':
												{
													position193 := position
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
															goto l158
														}
														position++
													}
												l194:
													{
														position196, tokenIndex196 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l197
														}
														position++
														goto l196
													l197:
														position, tokenIndex = position196, tokenIndex196
														if buffer[position] != rune('P') {
															goto l158
														}
														position++
													}
												l196:
													if !_rules[rulews]() {
														goto l158
													}
													if !_rules[ruleSrc8]() {
														goto l158
													}
													{
														add(ruleAction46, position)
													}
													add(ruleCp, position193)
												}
												break
											case 'O', 'o':
												{
													position199 := position
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
															goto l158
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
															goto l158
														}
														position++
													}
												l202:
													if !_rules[rulews]() {
														goto l158
													}
													if !_rules[ruleSrc8]() {
														goto l158
													}
													{
														add(ruleAction45, position)
													}
													add(ruleOr, position199)
												}
												break
											case 'X', 'x':
												{
													position205 := position
													{
														position206, tokenIndex206 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l207
														}
														position++
														goto l206
													l207:
														position, tokenIndex = position206, tokenIndex206
														if buffer[position] != rune('X') {
															goto l158
														}
														position++
													}
												l206:
													{
														position208, tokenIndex208 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l209
														}
														position++
														goto l208
													l209:
														position, tokenIndex = position208, tokenIndex208
														if buffer[position] != rune('O') {
															goto l158
														}
														position++
													}
												l208:
													{
														position210, tokenIndex210 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l211
														}
														position++
														goto l210
													l211:
														position, tokenIndex = position210, tokenIndex210
														if buffer[position] != rune('R') {
															goto l158
														}
														position++
													}
												l210:
													if !_rules[rulews]() {
														goto l158
													}
													if !_rules[ruleSrc8]() {
														goto l158
													}
													{
														add(ruleAction44, position)
													}
													add(ruleXor, position205)
												}
												break
											case 'A', 'a':
												{
													position213 := position
													{
														position214, tokenIndex214 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l215
														}
														position++
														goto l214
													l215:
														position, tokenIndex = position214, tokenIndex214
														if buffer[position] != rune('A') {
															goto l158
														}
														position++
													}
												l214:
													{
														position216, tokenIndex216 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l217
														}
														position++
														goto l216
													l217:
														position, tokenIndex = position216, tokenIndex216
														if buffer[position] != rune('N') {
															goto l158
														}
														position++
													}
												l216:
													{
														position218, tokenIndex218 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l219
														}
														position++
														goto l218
													l219:
														position, tokenIndex = position218, tokenIndex218
														if buffer[position] != rune('D') {
															goto l158
														}
														position++
													}
												l218:
													if !_rules[rulews]() {
														goto l158
													}
													if !_rules[ruleSrc8]() {
														goto l158
													}
													{
														add(ruleAction43, position)
													}
													add(ruleAnd, position213)
												}
												break
											default:
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
															goto l158
														}
														position++
													}
												l222:
													{
														position224, tokenIndex224 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l225
														}
														position++
														goto l224
													l225:
														position, tokenIndex = position224, tokenIndex224
														if buffer[position] != rune('B') {
															goto l158
														}
														position++
													}
												l224:
													{
														position226, tokenIndex226 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l227
														}
														position++
														goto l226
													l227:
														position, tokenIndex = position226, tokenIndex226
														if buffer[position] != rune('C') {
															goto l158
														}
														position++
													}
												l226:
													if !_rules[rulews]() {
														goto l158
													}
													{
														position228, tokenIndex228 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l229
														}
														position++
														goto l228
													l229:
														position, tokenIndex = position228, tokenIndex228
														if buffer[position] != rune('A') {
															goto l158
														}
														position++
													}
												l228:
													if !_rules[rulesep]() {
														goto l158
													}
													if !_rules[ruleSrc8]() {
														goto l158
													}
													{
														add(ruleAction42, position)
													}
													add(ruleSbc, position221)
												}
												break
											}
										}

									}
								l160:
									add(ruleAlu, position159)
								}
								goto l21
							l158:
								position, tokenIndex = position21, tokenIndex21
								{
									position232 := position
									{
										position233, tokenIndex233 := position, tokenIndex
										{
											position235 := position
											{
												position236, tokenIndex236 := position, tokenIndex
												{
													position238 := position
													{
														position239, tokenIndex239 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l240
														}
														position++
														goto l239
													l240:
														position, tokenIndex = position239, tokenIndex239
														if buffer[position] != rune('R') {
															goto l237
														}
														position++
													}
												l239:
													{
														position241, tokenIndex241 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l242
														}
														position++
														goto l241
													l242:
														position, tokenIndex = position241, tokenIndex241
														if buffer[position] != rune('L') {
															goto l237
														}
														position++
													}
												l241:
													{
														position243, tokenIndex243 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l244
														}
														position++
														goto l243
													l244:
														position, tokenIndex = position243, tokenIndex243
														if buffer[position] != rune('C') {
															goto l237
														}
														position++
													}
												l243:
													if !_rules[rulews]() {
														goto l237
													}
													if !_rules[ruleLoc8]() {
														goto l237
													}
													{
														position245, tokenIndex245 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l245
														}
														if !_rules[ruleCopy8]() {
															goto l245
														}
														goto l246
													l245:
														position, tokenIndex = position245, tokenIndex245
													}
												l246:
													{
														add(ruleAction47, position)
													}
													add(ruleRlc, position238)
												}
												goto l236
											l237:
												position, tokenIndex = position236, tokenIndex236
												{
													position249 := position
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
															goto l248
														}
														position++
													}
												l250:
													{
														position252, tokenIndex252 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l253
														}
														position++
														goto l252
													l253:
														position, tokenIndex = position252, tokenIndex252
														if buffer[position] != rune('R') {
															goto l248
														}
														position++
													}
												l252:
													{
														position254, tokenIndex254 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l255
														}
														position++
														goto l254
													l255:
														position, tokenIndex = position254, tokenIndex254
														if buffer[position] != rune('C') {
															goto l248
														}
														position++
													}
												l254:
													if !_rules[rulews]() {
														goto l248
													}
													if !_rules[ruleLoc8]() {
														goto l248
													}
													{
														position256, tokenIndex256 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l256
														}
														if !_rules[ruleCopy8]() {
															goto l256
														}
														goto l257
													l256:
														position, tokenIndex = position256, tokenIndex256
													}
												l257:
													{
														add(ruleAction48, position)
													}
													add(ruleRrc, position249)
												}
												goto l236
											l248:
												position, tokenIndex = position236, tokenIndex236
												{
													position260 := position
													{
														position261, tokenIndex261 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l262
														}
														position++
														goto l261
													l262:
														position, tokenIndex = position261, tokenIndex261
														if buffer[position] != rune('R') {
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
													if !_rules[rulews]() {
														goto l259
													}
													if !_rules[ruleLoc8]() {
														goto l259
													}
													{
														position265, tokenIndex265 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l265
														}
														if !_rules[ruleCopy8]() {
															goto l265
														}
														goto l266
													l265:
														position, tokenIndex = position265, tokenIndex265
													}
												l266:
													{
														add(ruleAction49, position)
													}
													add(ruleRl, position260)
												}
												goto l236
											l259:
												position, tokenIndex = position236, tokenIndex236
												{
													position269 := position
													{
														position270, tokenIndex270 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l271
														}
														position++
														goto l270
													l271:
														position, tokenIndex = position270, tokenIndex270
														if buffer[position] != rune('R') {
															goto l268
														}
														position++
													}
												l270:
													{
														position272, tokenIndex272 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l273
														}
														position++
														goto l272
													l273:
														position, tokenIndex = position272, tokenIndex272
														if buffer[position] != rune('R') {
															goto l268
														}
														position++
													}
												l272:
													if !_rules[rulews]() {
														goto l268
													}
													if !_rules[ruleLoc8]() {
														goto l268
													}
													{
														position274, tokenIndex274 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l274
														}
														if !_rules[ruleCopy8]() {
															goto l274
														}
														goto l275
													l274:
														position, tokenIndex = position274, tokenIndex274
													}
												l275:
													{
														add(ruleAction50, position)
													}
													add(ruleRr, position269)
												}
												goto l236
											l268:
												position, tokenIndex = position236, tokenIndex236
												{
													position278 := position
													{
														position279, tokenIndex279 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l280
														}
														position++
														goto l279
													l280:
														position, tokenIndex = position279, tokenIndex279
														if buffer[position] != rune('S') {
															goto l277
														}
														position++
													}
												l279:
													{
														position281, tokenIndex281 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l282
														}
														position++
														goto l281
													l282:
														position, tokenIndex = position281, tokenIndex281
														if buffer[position] != rune('L') {
															goto l277
														}
														position++
													}
												l281:
													{
														position283, tokenIndex283 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l284
														}
														position++
														goto l283
													l284:
														position, tokenIndex = position283, tokenIndex283
														if buffer[position] != rune('A') {
															goto l277
														}
														position++
													}
												l283:
													if !_rules[rulews]() {
														goto l277
													}
													if !_rules[ruleLoc8]() {
														goto l277
													}
													{
														position285, tokenIndex285 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l285
														}
														if !_rules[ruleCopy8]() {
															goto l285
														}
														goto l286
													l285:
														position, tokenIndex = position285, tokenIndex285
													}
												l286:
													{
														add(ruleAction51, position)
													}
													add(ruleSla, position278)
												}
												goto l236
											l277:
												position, tokenIndex = position236, tokenIndex236
												{
													position289 := position
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
															goto l288
														}
														position++
													}
												l290:
													{
														position292, tokenIndex292 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l293
														}
														position++
														goto l292
													l293:
														position, tokenIndex = position292, tokenIndex292
														if buffer[position] != rune('R') {
															goto l288
														}
														position++
													}
												l292:
													{
														position294, tokenIndex294 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l295
														}
														position++
														goto l294
													l295:
														position, tokenIndex = position294, tokenIndex294
														if buffer[position] != rune('A') {
															goto l288
														}
														position++
													}
												l294:
													if !_rules[rulews]() {
														goto l288
													}
													if !_rules[ruleLoc8]() {
														goto l288
													}
													{
														position296, tokenIndex296 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l296
														}
														if !_rules[ruleCopy8]() {
															goto l296
														}
														goto l297
													l296:
														position, tokenIndex = position296, tokenIndex296
													}
												l297:
													{
														add(ruleAction52, position)
													}
													add(ruleSra, position289)
												}
												goto l236
											l288:
												position, tokenIndex = position236, tokenIndex236
												{
													position300 := position
													{
														position301, tokenIndex301 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l302
														}
														position++
														goto l301
													l302:
														position, tokenIndex = position301, tokenIndex301
														if buffer[position] != rune('S') {
															goto l299
														}
														position++
													}
												l301:
													{
														position303, tokenIndex303 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l304
														}
														position++
														goto l303
													l304:
														position, tokenIndex = position303, tokenIndex303
														if buffer[position] != rune('L') {
															goto l299
														}
														position++
													}
												l303:
													{
														position305, tokenIndex305 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l306
														}
														position++
														goto l305
													l306:
														position, tokenIndex = position305, tokenIndex305
														if buffer[position] != rune('L') {
															goto l299
														}
														position++
													}
												l305:
													if !_rules[rulews]() {
														goto l299
													}
													if !_rules[ruleLoc8]() {
														goto l299
													}
													{
														position307, tokenIndex307 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l307
														}
														if !_rules[ruleCopy8]() {
															goto l307
														}
														goto l308
													l307:
														position, tokenIndex = position307, tokenIndex307
													}
												l308:
													{
														add(ruleAction53, position)
													}
													add(ruleSll, position300)
												}
												goto l236
											l299:
												position, tokenIndex = position236, tokenIndex236
												{
													position310 := position
													{
														position311, tokenIndex311 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l312
														}
														position++
														goto l311
													l312:
														position, tokenIndex = position311, tokenIndex311
														if buffer[position] != rune('S') {
															goto l234
														}
														position++
													}
												l311:
													{
														position313, tokenIndex313 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l314
														}
														position++
														goto l313
													l314:
														position, tokenIndex = position313, tokenIndex313
														if buffer[position] != rune('R') {
															goto l234
														}
														position++
													}
												l313:
													{
														position315, tokenIndex315 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l316
														}
														position++
														goto l315
													l316:
														position, tokenIndex = position315, tokenIndex315
														if buffer[position] != rune('L') {
															goto l234
														}
														position++
													}
												l315:
													if !_rules[rulews]() {
														goto l234
													}
													if !_rules[ruleLoc8]() {
														goto l234
													}
													{
														position317, tokenIndex317 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l317
														}
														if !_rules[ruleCopy8]() {
															goto l317
														}
														goto l318
													l317:
														position, tokenIndex = position317, tokenIndex317
													}
												l318:
													{
														add(ruleAction54, position)
													}
													add(ruleSrl, position310)
												}
											}
										l236:
											add(ruleRot, position235)
										}
										goto l233
									l234:
										position, tokenIndex = position233, tokenIndex233
										{
											switch buffer[position] {
											case 'S', 's':
												{
													position321 := position
													{
														position322, tokenIndex322 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l323
														}
														position++
														goto l322
													l323:
														position, tokenIndex = position322, tokenIndex322
														if buffer[position] != rune('S') {
															goto l231
														}
														position++
													}
												l322:
													{
														position324, tokenIndex324 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l325
														}
														position++
														goto l324
													l325:
														position, tokenIndex = position324, tokenIndex324
														if buffer[position] != rune('E') {
															goto l231
														}
														position++
													}
												l324:
													{
														position326, tokenIndex326 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l327
														}
														position++
														goto l326
													l327:
														position, tokenIndex = position326, tokenIndex326
														if buffer[position] != rune('T') {
															goto l231
														}
														position++
													}
												l326:
													if !_rules[rulews]() {
														goto l231
													}
													if !_rules[ruleoctaldigit]() {
														goto l231
													}
													if !_rules[rulesep]() {
														goto l231
													}
													if !_rules[ruleLoc8]() {
														goto l231
													}
													{
														position328, tokenIndex328 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l328
														}
														if !_rules[ruleCopy8]() {
															goto l328
														}
														goto l329
													l328:
														position, tokenIndex = position328, tokenIndex328
													}
												l329:
													{
														add(ruleAction57, position)
													}
													add(ruleSet, position321)
												}
												break
											case 'R', 'r':
												{
													position331 := position
													{
														position332, tokenIndex332 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l333
														}
														position++
														goto l332
													l333:
														position, tokenIndex = position332, tokenIndex332
														if buffer[position] != rune('R') {
															goto l231
														}
														position++
													}
												l332:
													{
														position334, tokenIndex334 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l335
														}
														position++
														goto l334
													l335:
														position, tokenIndex = position334, tokenIndex334
														if buffer[position] != rune('E') {
															goto l231
														}
														position++
													}
												l334:
													{
														position336, tokenIndex336 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l337
														}
														position++
														goto l336
													l337:
														position, tokenIndex = position336, tokenIndex336
														if buffer[position] != rune('S') {
															goto l231
														}
														position++
													}
												l336:
													if !_rules[rulews]() {
														goto l231
													}
													if !_rules[ruleoctaldigit]() {
														goto l231
													}
													if !_rules[rulesep]() {
														goto l231
													}
													if !_rules[ruleLoc8]() {
														goto l231
													}
													{
														position338, tokenIndex338 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l338
														}
														if !_rules[ruleCopy8]() {
															goto l338
														}
														goto l339
													l338:
														position, tokenIndex = position338, tokenIndex338
													}
												l339:
													{
														add(ruleAction56, position)
													}
													add(ruleRes, position331)
												}
												break
											default:
												{
													position341 := position
													{
														position342, tokenIndex342 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l343
														}
														position++
														goto l342
													l343:
														position, tokenIndex = position342, tokenIndex342
														if buffer[position] != rune('B') {
															goto l231
														}
														position++
													}
												l342:
													{
														position344, tokenIndex344 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l345
														}
														position++
														goto l344
													l345:
														position, tokenIndex = position344, tokenIndex344
														if buffer[position] != rune('I') {
															goto l231
														}
														position++
													}
												l344:
													{
														position346, tokenIndex346 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l347
														}
														position++
														goto l346
													l347:
														position, tokenIndex = position346, tokenIndex346
														if buffer[position] != rune('T') {
															goto l231
														}
														position++
													}
												l346:
													if !_rules[rulews]() {
														goto l231
													}
													if !_rules[ruleoctaldigit]() {
														goto l231
													}
													if !_rules[rulesep]() {
														goto l231
													}
													if !_rules[ruleLoc8]() {
														goto l231
													}
													{
														add(ruleAction55, position)
													}
													add(ruleBit, position341)
												}
												break
											}
										}

									}
								l233:
									add(ruleBitOp, position232)
								}
								goto l21
							l231:
								position, tokenIndex = position21, tokenIndex21
								{
									position350 := position
									{
										position351, tokenIndex351 := position, tokenIndex
										{
											position353 := position
											{
												position354 := position
												{
													position355, tokenIndex355 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l356
													}
													position++
													goto l355
												l356:
													position, tokenIndex = position355, tokenIndex355
													if buffer[position] != rune('R') {
														goto l352
													}
													position++
												}
											l355:
												{
													position357, tokenIndex357 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l358
													}
													position++
													goto l357
												l358:
													position, tokenIndex = position357, tokenIndex357
													if buffer[position] != rune('E') {
														goto l352
													}
													position++
												}
											l357:
												{
													position359, tokenIndex359 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l360
													}
													position++
													goto l359
												l360:
													position, tokenIndex = position359, tokenIndex359
													if buffer[position] != rune('T') {
														goto l352
													}
													position++
												}
											l359:
												{
													position361, tokenIndex361 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l362
													}
													position++
													goto l361
												l362:
													position, tokenIndex = position361, tokenIndex361
													if buffer[position] != rune('N') {
														goto l352
													}
													position++
												}
											l361:
												add(rulePegText, position354)
											}
											{
												add(ruleAction72, position)
											}
											add(ruleRetn, position353)
										}
										goto l351
									l352:
										position, tokenIndex = position351, tokenIndex351
										{
											position365 := position
											{
												position366 := position
												{
													position367, tokenIndex367 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l368
													}
													position++
													goto l367
												l368:
													position, tokenIndex = position367, tokenIndex367
													if buffer[position] != rune('R') {
														goto l364
													}
													position++
												}
											l367:
												{
													position369, tokenIndex369 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l370
													}
													position++
													goto l369
												l370:
													position, tokenIndex = position369, tokenIndex369
													if buffer[position] != rune('E') {
														goto l364
													}
													position++
												}
											l369:
												{
													position371, tokenIndex371 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l372
													}
													position++
													goto l371
												l372:
													position, tokenIndex = position371, tokenIndex371
													if buffer[position] != rune('T') {
														goto l364
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
														goto l364
													}
													position++
												}
											l373:
												add(rulePegText, position366)
											}
											{
												add(ruleAction73, position)
											}
											add(ruleReti, position365)
										}
										goto l351
									l364:
										position, tokenIndex = position351, tokenIndex351
										{
											position377 := position
											{
												position378 := position
												{
													position379, tokenIndex379 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l380
													}
													position++
													goto l379
												l380:
													position, tokenIndex = position379, tokenIndex379
													if buffer[position] != rune('R') {
														goto l376
													}
													position++
												}
											l379:
												{
													position381, tokenIndex381 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l382
													}
													position++
													goto l381
												l382:
													position, tokenIndex = position381, tokenIndex381
													if buffer[position] != rune('R') {
														goto l376
													}
													position++
												}
											l381:
												{
													position383, tokenIndex383 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l384
													}
													position++
													goto l383
												l384:
													position, tokenIndex = position383, tokenIndex383
													if buffer[position] != rune('D') {
														goto l376
													}
													position++
												}
											l383:
												add(rulePegText, position378)
											}
											{
												add(ruleAction74, position)
											}
											add(ruleRrd, position377)
										}
										goto l351
									l376:
										position, tokenIndex = position351, tokenIndex351
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
												if buffer[position] != rune('0') {
													goto l386
												}
												position++
												add(rulePegText, position388)
											}
											{
												add(ruleAction76, position)
											}
											add(ruleIm0, position387)
										}
										goto l351
									l386:
										position, tokenIndex = position351, tokenIndex351
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
												if buffer[position] != rune('1') {
													goto l394
												}
												position++
												add(rulePegText, position396)
											}
											{
												add(ruleAction77, position)
											}
											add(ruleIm1, position395)
										}
										goto l351
									l394:
										position, tokenIndex = position351, tokenIndex351
										{
											position403 := position
											{
												position404 := position
												{
													position405, tokenIndex405 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l406
													}
													position++
													goto l405
												l406:
													position, tokenIndex = position405, tokenIndex405
													if buffer[position] != rune('I') {
														goto l402
													}
													position++
												}
											l405:
												{
													position407, tokenIndex407 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l408
													}
													position++
													goto l407
												l408:
													position, tokenIndex = position407, tokenIndex407
													if buffer[position] != rune('M') {
														goto l402
													}
													position++
												}
											l407:
												if buffer[position] != rune(' ') {
													goto l402
												}
												position++
												if buffer[position] != rune('2') {
													goto l402
												}
												position++
												add(rulePegText, position404)
											}
											{
												add(ruleAction78, position)
											}
											add(ruleIm2, position403)
										}
										goto l351
									l402:
										position, tokenIndex = position351, tokenIndex351
										{
											switch buffer[position] {
											case 'I', 'O', 'i', 'o':
												{
													position411 := position
													{
														position412, tokenIndex412 := position, tokenIndex
														{
															position414 := position
															{
																position415 := position
																{
																	position416, tokenIndex416 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l417
																	}
																	position++
																	goto l416
																l417:
																	position, tokenIndex = position416, tokenIndex416
																	if buffer[position] != rune('I') {
																		goto l413
																	}
																	position++
																}
															l416:
																{
																	position418, tokenIndex418 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l419
																	}
																	position++
																	goto l418
																l419:
																	position, tokenIndex = position418, tokenIndex418
																	if buffer[position] != rune('N') {
																		goto l413
																	}
																	position++
																}
															l418:
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
																		goto l413
																	}
																	position++
																}
															l420:
																{
																	position422, tokenIndex422 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l423
																	}
																	position++
																	goto l422
																l423:
																	position, tokenIndex = position422, tokenIndex422
																	if buffer[position] != rune('R') {
																		goto l413
																	}
																	position++
																}
															l422:
																add(rulePegText, position415)
															}
															{
																add(ruleAction89, position)
															}
															add(ruleInir, position414)
														}
														goto l412
													l413:
														position, tokenIndex = position412, tokenIndex412
														{
															position426 := position
															{
																position427 := position
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
																		goto l425
																	}
																	position++
																}
															l428:
																{
																	position430, tokenIndex430 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l431
																	}
																	position++
																	goto l430
																l431:
																	position, tokenIndex = position430, tokenIndex430
																	if buffer[position] != rune('N') {
																		goto l425
																	}
																	position++
																}
															l430:
																{
																	position432, tokenIndex432 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l433
																	}
																	position++
																	goto l432
																l433:
																	position, tokenIndex = position432, tokenIndex432
																	if buffer[position] != rune('I') {
																		goto l425
																	}
																	position++
																}
															l432:
																add(rulePegText, position427)
															}
															{
																add(ruleAction81, position)
															}
															add(ruleIni, position426)
														}
														goto l412
													l425:
														position, tokenIndex = position412, tokenIndex412
														{
															position436 := position
															{
																position437 := position
																{
																	position438, tokenIndex438 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l439
																	}
																	position++
																	goto l438
																l439:
																	position, tokenIndex = position438, tokenIndex438
																	if buffer[position] != rune('O') {
																		goto l435
																	}
																	position++
																}
															l438:
																{
																	position440, tokenIndex440 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l441
																	}
																	position++
																	goto l440
																l441:
																	position, tokenIndex = position440, tokenIndex440
																	if buffer[position] != rune('T') {
																		goto l435
																	}
																	position++
																}
															l440:
																{
																	position442, tokenIndex442 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l443
																	}
																	position++
																	goto l442
																l443:
																	position, tokenIndex = position442, tokenIndex442
																	if buffer[position] != rune('I') {
																		goto l435
																	}
																	position++
																}
															l442:
																{
																	position444, tokenIndex444 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l445
																	}
																	position++
																	goto l444
																l445:
																	position, tokenIndex = position444, tokenIndex444
																	if buffer[position] != rune('R') {
																		goto l435
																	}
																	position++
																}
															l444:
																add(rulePegText, position437)
															}
															{
																add(ruleAction90, position)
															}
															add(ruleOtir, position436)
														}
														goto l412
													l435:
														position, tokenIndex = position412, tokenIndex412
														{
															position448 := position
															{
																position449 := position
																{
																	position450, tokenIndex450 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l451
																	}
																	position++
																	goto l450
																l451:
																	position, tokenIndex = position450, tokenIndex450
																	if buffer[position] != rune('O') {
																		goto l447
																	}
																	position++
																}
															l450:
																{
																	position452, tokenIndex452 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l453
																	}
																	position++
																	goto l452
																l453:
																	position, tokenIndex = position452, tokenIndex452
																	if buffer[position] != rune('U') {
																		goto l447
																	}
																	position++
																}
															l452:
																{
																	position454, tokenIndex454 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l455
																	}
																	position++
																	goto l454
																l455:
																	position, tokenIndex = position454, tokenIndex454
																	if buffer[position] != rune('T') {
																		goto l447
																	}
																	position++
																}
															l454:
																{
																	position456, tokenIndex456 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l457
																	}
																	position++
																	goto l456
																l457:
																	position, tokenIndex = position456, tokenIndex456
																	if buffer[position] != rune('I') {
																		goto l447
																	}
																	position++
																}
															l456:
																add(rulePegText, position449)
															}
															{
																add(ruleAction82, position)
															}
															add(ruleOuti, position448)
														}
														goto l412
													l447:
														position, tokenIndex = position412, tokenIndex412
														{
															position460 := position
															{
																position461 := position
																{
																	position462, tokenIndex462 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l463
																	}
																	position++
																	goto l462
																l463:
																	position, tokenIndex = position462, tokenIndex462
																	if buffer[position] != rune('I') {
																		goto l459
																	}
																	position++
																}
															l462:
																{
																	position464, tokenIndex464 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l465
																	}
																	position++
																	goto l464
																l465:
																	position, tokenIndex = position464, tokenIndex464
																	if buffer[position] != rune('N') {
																		goto l459
																	}
																	position++
																}
															l464:
																{
																	position466, tokenIndex466 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l467
																	}
																	position++
																	goto l466
																l467:
																	position, tokenIndex = position466, tokenIndex466
																	if buffer[position] != rune('D') {
																		goto l459
																	}
																	position++
																}
															l466:
																{
																	position468, tokenIndex468 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l469
																	}
																	position++
																	goto l468
																l469:
																	position, tokenIndex = position468, tokenIndex468
																	if buffer[position] != rune('R') {
																		goto l459
																	}
																	position++
																}
															l468:
																add(rulePegText, position461)
															}
															{
																add(ruleAction93, position)
															}
															add(ruleIndr, position460)
														}
														goto l412
													l459:
														position, tokenIndex = position412, tokenIndex412
														{
															position472 := position
															{
																position473 := position
																{
																	position474, tokenIndex474 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l475
																	}
																	position++
																	goto l474
																l475:
																	position, tokenIndex = position474, tokenIndex474
																	if buffer[position] != rune('I') {
																		goto l471
																	}
																	position++
																}
															l474:
																{
																	position476, tokenIndex476 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l477
																	}
																	position++
																	goto l476
																l477:
																	position, tokenIndex = position476, tokenIndex476
																	if buffer[position] != rune('N') {
																		goto l471
																	}
																	position++
																}
															l476:
																{
																	position478, tokenIndex478 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l479
																	}
																	position++
																	goto l478
																l479:
																	position, tokenIndex = position478, tokenIndex478
																	if buffer[position] != rune('D') {
																		goto l471
																	}
																	position++
																}
															l478:
																add(rulePegText, position473)
															}
															{
																add(ruleAction85, position)
															}
															add(ruleInd, position472)
														}
														goto l412
													l471:
														position, tokenIndex = position412, tokenIndex412
														{
															position482 := position
															{
																position483 := position
																{
																	position484, tokenIndex484 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l485
																	}
																	position++
																	goto l484
																l485:
																	position, tokenIndex = position484, tokenIndex484
																	if buffer[position] != rune('O') {
																		goto l481
																	}
																	position++
																}
															l484:
																{
																	position486, tokenIndex486 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l487
																	}
																	position++
																	goto l486
																l487:
																	position, tokenIndex = position486, tokenIndex486
																	if buffer[position] != rune('T') {
																		goto l481
																	}
																	position++
																}
															l486:
																{
																	position488, tokenIndex488 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l489
																	}
																	position++
																	goto l488
																l489:
																	position, tokenIndex = position488, tokenIndex488
																	if buffer[position] != rune('D') {
																		goto l481
																	}
																	position++
																}
															l488:
																{
																	position490, tokenIndex490 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l491
																	}
																	position++
																	goto l490
																l491:
																	position, tokenIndex = position490, tokenIndex490
																	if buffer[position] != rune('R') {
																		goto l481
																	}
																	position++
																}
															l490:
																add(rulePegText, position483)
															}
															{
																add(ruleAction94, position)
															}
															add(ruleOtdr, position482)
														}
														goto l412
													l481:
														position, tokenIndex = position412, tokenIndex412
														{
															position493 := position
															{
																position494 := position
																{
																	position495, tokenIndex495 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l496
																	}
																	position++
																	goto l495
																l496:
																	position, tokenIndex = position495, tokenIndex495
																	if buffer[position] != rune('O') {
																		goto l349
																	}
																	position++
																}
															l495:
																{
																	position497, tokenIndex497 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l498
																	}
																	position++
																	goto l497
																l498:
																	position, tokenIndex = position497, tokenIndex497
																	if buffer[position] != rune('U') {
																		goto l349
																	}
																	position++
																}
															l497:
																{
																	position499, tokenIndex499 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l500
																	}
																	position++
																	goto l499
																l500:
																	position, tokenIndex = position499, tokenIndex499
																	if buffer[position] != rune('T') {
																		goto l349
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
																		goto l349
																	}
																	position++
																}
															l501:
																add(rulePegText, position494)
															}
															{
																add(ruleAction86, position)
															}
															add(ruleOutd, position493)
														}
													}
												l412:
													add(ruleBlitIO, position411)
												}
												break
											case 'R', 'r':
												{
													position504 := position
													{
														position505 := position
														{
															position506, tokenIndex506 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l507
															}
															position++
															goto l506
														l507:
															position, tokenIndex = position506, tokenIndex506
															if buffer[position] != rune('R') {
																goto l349
															}
															position++
														}
													l506:
														{
															position508, tokenIndex508 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l509
															}
															position++
															goto l508
														l509:
															position, tokenIndex = position508, tokenIndex508
															if buffer[position] != rune('L') {
																goto l349
															}
															position++
														}
													l508:
														{
															position510, tokenIndex510 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l511
															}
															position++
															goto l510
														l511:
															position, tokenIndex = position510, tokenIndex510
															if buffer[position] != rune('D') {
																goto l349
															}
															position++
														}
													l510:
														add(rulePegText, position505)
													}
													{
														add(ruleAction75, position)
													}
													add(ruleRld, position504)
												}
												break
											case 'N', 'n':
												{
													position513 := position
													{
														position514 := position
														{
															position515, tokenIndex515 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l516
															}
															position++
															goto l515
														l516:
															position, tokenIndex = position515, tokenIndex515
															if buffer[position] != rune('N') {
																goto l349
															}
															position++
														}
													l515:
														{
															position517, tokenIndex517 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l518
															}
															position++
															goto l517
														l518:
															position, tokenIndex = position517, tokenIndex517
															if buffer[position] != rune('E') {
																goto l349
															}
															position++
														}
													l517:
														{
															position519, tokenIndex519 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l520
															}
															position++
															goto l519
														l520:
															position, tokenIndex = position519, tokenIndex519
															if buffer[position] != rune('G') {
																goto l349
															}
															position++
														}
													l519:
														add(rulePegText, position514)
													}
													{
														add(ruleAction71, position)
													}
													add(ruleNeg, position513)
												}
												break
											default:
												{
													position522 := position
													{
														position523, tokenIndex523 := position, tokenIndex
														{
															position525 := position
															{
																position526 := position
																{
																	position527, tokenIndex527 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l528
																	}
																	position++
																	goto l527
																l528:
																	position, tokenIndex = position527, tokenIndex527
																	if buffer[position] != rune('L') {
																		goto l524
																	}
																	position++
																}
															l527:
																{
																	position529, tokenIndex529 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l530
																	}
																	position++
																	goto l529
																l530:
																	position, tokenIndex = position529, tokenIndex529
																	if buffer[position] != rune('D') {
																		goto l524
																	}
																	position++
																}
															l529:
																{
																	position531, tokenIndex531 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l532
																	}
																	position++
																	goto l531
																l532:
																	position, tokenIndex = position531, tokenIndex531
																	if buffer[position] != rune('I') {
																		goto l524
																	}
																	position++
																}
															l531:
																{
																	position533, tokenIndex533 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l534
																	}
																	position++
																	goto l533
																l534:
																	position, tokenIndex = position533, tokenIndex533
																	if buffer[position] != rune('R') {
																		goto l524
																	}
																	position++
																}
															l533:
																add(rulePegText, position526)
															}
															{
																add(ruleAction87, position)
															}
															add(ruleLdir, position525)
														}
														goto l523
													l524:
														position, tokenIndex = position523, tokenIndex523
														{
															position537 := position
															{
																position538 := position
																{
																	position539, tokenIndex539 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l540
																	}
																	position++
																	goto l539
																l540:
																	position, tokenIndex = position539, tokenIndex539
																	if buffer[position] != rune('L') {
																		goto l536
																	}
																	position++
																}
															l539:
																{
																	position541, tokenIndex541 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l542
																	}
																	position++
																	goto l541
																l542:
																	position, tokenIndex = position541, tokenIndex541
																	if buffer[position] != rune('D') {
																		goto l536
																	}
																	position++
																}
															l541:
																{
																	position543, tokenIndex543 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l544
																	}
																	position++
																	goto l543
																l544:
																	position, tokenIndex = position543, tokenIndex543
																	if buffer[position] != rune('I') {
																		goto l536
																	}
																	position++
																}
															l543:
																add(rulePegText, position538)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleLdi, position537)
														}
														goto l523
													l536:
														position, tokenIndex = position523, tokenIndex523
														{
															position547 := position
															{
																position548 := position
																{
																	position549, tokenIndex549 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l550
																	}
																	position++
																	goto l549
																l550:
																	position, tokenIndex = position549, tokenIndex549
																	if buffer[position] != rune('C') {
																		goto l546
																	}
																	position++
																}
															l549:
																{
																	position551, tokenIndex551 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l552
																	}
																	position++
																	goto l551
																l552:
																	position, tokenIndex = position551, tokenIndex551
																	if buffer[position] != rune('P') {
																		goto l546
																	}
																	position++
																}
															l551:
																{
																	position553, tokenIndex553 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l554
																	}
																	position++
																	goto l553
																l554:
																	position, tokenIndex = position553, tokenIndex553
																	if buffer[position] != rune('I') {
																		goto l546
																	}
																	position++
																}
															l553:
																{
																	position555, tokenIndex555 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l556
																	}
																	position++
																	goto l555
																l556:
																	position, tokenIndex = position555, tokenIndex555
																	if buffer[position] != rune('R') {
																		goto l546
																	}
																	position++
																}
															l555:
																add(rulePegText, position548)
															}
															{
																add(ruleAction88, position)
															}
															add(ruleCpir, position547)
														}
														goto l523
													l546:
														position, tokenIndex = position523, tokenIndex523
														{
															position559 := position
															{
																position560 := position
																{
																	position561, tokenIndex561 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l562
																	}
																	position++
																	goto l561
																l562:
																	position, tokenIndex = position561, tokenIndex561
																	if buffer[position] != rune('C') {
																		goto l558
																	}
																	position++
																}
															l561:
																{
																	position563, tokenIndex563 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l564
																	}
																	position++
																	goto l563
																l564:
																	position, tokenIndex = position563, tokenIndex563
																	if buffer[position] != rune('P') {
																		goto l558
																	}
																	position++
																}
															l563:
																{
																	position565, tokenIndex565 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l566
																	}
																	position++
																	goto l565
																l566:
																	position, tokenIndex = position565, tokenIndex565
																	if buffer[position] != rune('I') {
																		goto l558
																	}
																	position++
																}
															l565:
																add(rulePegText, position560)
															}
															{
																add(ruleAction80, position)
															}
															add(ruleCpi, position559)
														}
														goto l523
													l558:
														position, tokenIndex = position523, tokenIndex523
														{
															position569 := position
															{
																position570 := position
																{
																	position571, tokenIndex571 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l572
																	}
																	position++
																	goto l571
																l572:
																	position, tokenIndex = position571, tokenIndex571
																	if buffer[position] != rune('L') {
																		goto l568
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
																		goto l568
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
																		goto l568
																	}
																	position++
																}
															l575:
																{
																	position577, tokenIndex577 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l578
																	}
																	position++
																	goto l577
																l578:
																	position, tokenIndex = position577, tokenIndex577
																	if buffer[position] != rune('R') {
																		goto l568
																	}
																	position++
																}
															l577:
																add(rulePegText, position570)
															}
															{
																add(ruleAction91, position)
															}
															add(ruleLddr, position569)
														}
														goto l523
													l568:
														position, tokenIndex = position523, tokenIndex523
														{
															position581 := position
															{
																position582 := position
																{
																	position583, tokenIndex583 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l584
																	}
																	position++
																	goto l583
																l584:
																	position, tokenIndex = position583, tokenIndex583
																	if buffer[position] != rune('L') {
																		goto l580
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
																		goto l580
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
																		goto l580
																	}
																	position++
																}
															l587:
																add(rulePegText, position582)
															}
															{
																add(ruleAction83, position)
															}
															add(ruleLdd, position581)
														}
														goto l523
													l580:
														position, tokenIndex = position523, tokenIndex523
														{
															position591 := position
															{
																position592 := position
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
																		goto l590
																	}
																	position++
																}
															l593:
																{
																	position595, tokenIndex595 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l596
																	}
																	position++
																	goto l595
																l596:
																	position, tokenIndex = position595, tokenIndex595
																	if buffer[position] != rune('P') {
																		goto l590
																	}
																	position++
																}
															l595:
																{
																	position597, tokenIndex597 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l598
																	}
																	position++
																	goto l597
																l598:
																	position, tokenIndex = position597, tokenIndex597
																	if buffer[position] != rune('D') {
																		goto l590
																	}
																	position++
																}
															l597:
																{
																	position599, tokenIndex599 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l600
																	}
																	position++
																	goto l599
																l600:
																	position, tokenIndex = position599, tokenIndex599
																	if buffer[position] != rune('R') {
																		goto l590
																	}
																	position++
																}
															l599:
																add(rulePegText, position592)
															}
															{
																add(ruleAction92, position)
															}
															add(ruleCpdr, position591)
														}
														goto l523
													l590:
														position, tokenIndex = position523, tokenIndex523
														{
															position602 := position
															{
																position603 := position
																{
																	position604, tokenIndex604 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l605
																	}
																	position++
																	goto l604
																l605:
																	position, tokenIndex = position604, tokenIndex604
																	if buffer[position] != rune('C') {
																		goto l349
																	}
																	position++
																}
															l604:
																{
																	position606, tokenIndex606 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l607
																	}
																	position++
																	goto l606
																l607:
																	position, tokenIndex = position606, tokenIndex606
																	if buffer[position] != rune('P') {
																		goto l349
																	}
																	position++
																}
															l606:
																{
																	position608, tokenIndex608 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l609
																	}
																	position++
																	goto l608
																l609:
																	position, tokenIndex = position608, tokenIndex608
																	if buffer[position] != rune('D') {
																		goto l349
																	}
																	position++
																}
															l608:
																add(rulePegText, position603)
															}
															{
																add(ruleAction84, position)
															}
															add(ruleCpd, position602)
														}
													}
												l523:
													add(ruleBlit, position522)
												}
												break
											}
										}

									}
								l351:
									add(ruleEDSimple, position350)
								}
								goto l21
							l349:
								position, tokenIndex = position21, tokenIndex21
								{
									position612 := position
									{
										position613, tokenIndex613 := position, tokenIndex
										{
											position615 := position
											{
												position616 := position
												{
													position617, tokenIndex617 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l618
													}
													position++
													goto l617
												l618:
													position, tokenIndex = position617, tokenIndex617
													if buffer[position] != rune('R') {
														goto l614
													}
													position++
												}
											l617:
												{
													position619, tokenIndex619 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l620
													}
													position++
													goto l619
												l620:
													position, tokenIndex = position619, tokenIndex619
													if buffer[position] != rune('L') {
														goto l614
													}
													position++
												}
											l619:
												{
													position621, tokenIndex621 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l622
													}
													position++
													goto l621
												l622:
													position, tokenIndex = position621, tokenIndex621
													if buffer[position] != rune('C') {
														goto l614
													}
													position++
												}
											l621:
												{
													position623, tokenIndex623 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l624
													}
													position++
													goto l623
												l624:
													position, tokenIndex = position623, tokenIndex623
													if buffer[position] != rune('A') {
														goto l614
													}
													position++
												}
											l623:
												add(rulePegText, position616)
											}
											{
												add(ruleAction60, position)
											}
											add(ruleRlca, position615)
										}
										goto l613
									l614:
										position, tokenIndex = position613, tokenIndex613
										{
											position627 := position
											{
												position628 := position
												{
													position629, tokenIndex629 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l630
													}
													position++
													goto l629
												l630:
													position, tokenIndex = position629, tokenIndex629
													if buffer[position] != rune('R') {
														goto l626
													}
													position++
												}
											l629:
												{
													position631, tokenIndex631 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l632
													}
													position++
													goto l631
												l632:
													position, tokenIndex = position631, tokenIndex631
													if buffer[position] != rune('R') {
														goto l626
													}
													position++
												}
											l631:
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
														goto l626
													}
													position++
												}
											l633:
												{
													position635, tokenIndex635 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l636
													}
													position++
													goto l635
												l636:
													position, tokenIndex = position635, tokenIndex635
													if buffer[position] != rune('A') {
														goto l626
													}
													position++
												}
											l635:
												add(rulePegText, position628)
											}
											{
												add(ruleAction61, position)
											}
											add(ruleRrca, position627)
										}
										goto l613
									l626:
										position, tokenIndex = position613, tokenIndex613
										{
											position639 := position
											{
												position640 := position
												{
													position641, tokenIndex641 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l642
													}
													position++
													goto l641
												l642:
													position, tokenIndex = position641, tokenIndex641
													if buffer[position] != rune('R') {
														goto l638
													}
													position++
												}
											l641:
												{
													position643, tokenIndex643 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l644
													}
													position++
													goto l643
												l644:
													position, tokenIndex = position643, tokenIndex643
													if buffer[position] != rune('L') {
														goto l638
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
														goto l638
													}
													position++
												}
											l645:
												add(rulePegText, position640)
											}
											{
												add(ruleAction62, position)
											}
											add(ruleRla, position639)
										}
										goto l613
									l638:
										position, tokenIndex = position613, tokenIndex613
										{
											position649 := position
											{
												position650 := position
												{
													position651, tokenIndex651 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l652
													}
													position++
													goto l651
												l652:
													position, tokenIndex = position651, tokenIndex651
													if buffer[position] != rune('D') {
														goto l648
													}
													position++
												}
											l651:
												{
													position653, tokenIndex653 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l654
													}
													position++
													goto l653
												l654:
													position, tokenIndex = position653, tokenIndex653
													if buffer[position] != rune('A') {
														goto l648
													}
													position++
												}
											l653:
												{
													position655, tokenIndex655 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l656
													}
													position++
													goto l655
												l656:
													position, tokenIndex = position655, tokenIndex655
													if buffer[position] != rune('A') {
														goto l648
													}
													position++
												}
											l655:
												add(rulePegText, position650)
											}
											{
												add(ruleAction64, position)
											}
											add(ruleDaa, position649)
										}
										goto l613
									l648:
										position, tokenIndex = position613, tokenIndex613
										{
											position659 := position
											{
												position660 := position
												{
													position661, tokenIndex661 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l662
													}
													position++
													goto l661
												l662:
													position, tokenIndex = position661, tokenIndex661
													if buffer[position] != rune('C') {
														goto l658
													}
													position++
												}
											l661:
												{
													position663, tokenIndex663 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l664
													}
													position++
													goto l663
												l664:
													position, tokenIndex = position663, tokenIndex663
													if buffer[position] != rune('P') {
														goto l658
													}
													position++
												}
											l663:
												{
													position665, tokenIndex665 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l666
													}
													position++
													goto l665
												l666:
													position, tokenIndex = position665, tokenIndex665
													if buffer[position] != rune('L') {
														goto l658
													}
													position++
												}
											l665:
												add(rulePegText, position660)
											}
											{
												add(ruleAction65, position)
											}
											add(ruleCpl, position659)
										}
										goto l613
									l658:
										position, tokenIndex = position613, tokenIndex613
										{
											position669 := position
											{
												position670 := position
												{
													position671, tokenIndex671 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l672
													}
													position++
													goto l671
												l672:
													position, tokenIndex = position671, tokenIndex671
													if buffer[position] != rune('E') {
														goto l668
													}
													position++
												}
											l671:
												{
													position673, tokenIndex673 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l674
													}
													position++
													goto l673
												l674:
													position, tokenIndex = position673, tokenIndex673
													if buffer[position] != rune('X') {
														goto l668
													}
													position++
												}
											l673:
												{
													position675, tokenIndex675 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l676
													}
													position++
													goto l675
												l676:
													position, tokenIndex = position675, tokenIndex675
													if buffer[position] != rune('X') {
														goto l668
													}
													position++
												}
											l675:
												add(rulePegText, position670)
											}
											{
												add(ruleAction68, position)
											}
											add(ruleExx, position669)
										}
										goto l613
									l668:
										position, tokenIndex = position613, tokenIndex613
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position679 := position
													{
														position680 := position
														{
															position681, tokenIndex681 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l682
															}
															position++
															goto l681
														l682:
															position, tokenIndex = position681, tokenIndex681
															if buffer[position] != rune('E') {
																goto l611
															}
															position++
														}
													l681:
														{
															position683, tokenIndex683 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l684
															}
															position++
															goto l683
														l684:
															position, tokenIndex = position683, tokenIndex683
															if buffer[position] != rune('I') {
																goto l611
															}
															position++
														}
													l683:
														add(rulePegText, position680)
													}
													{
														add(ruleAction70, position)
													}
													add(ruleEi, position679)
												}
												break
											case 'D', 'd':
												{
													position686 := position
													{
														position687 := position
														{
															position688, tokenIndex688 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l689
															}
															position++
															goto l688
														l689:
															position, tokenIndex = position688, tokenIndex688
															if buffer[position] != rune('D') {
																goto l611
															}
															position++
														}
													l688:
														{
															position690, tokenIndex690 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l691
															}
															position++
															goto l690
														l691:
															position, tokenIndex = position690, tokenIndex690
															if buffer[position] != rune('I') {
																goto l611
															}
															position++
														}
													l690:
														add(rulePegText, position687)
													}
													{
														add(ruleAction69, position)
													}
													add(ruleDi, position686)
												}
												break
											case 'C', 'c':
												{
													position693 := position
													{
														position694 := position
														{
															position695, tokenIndex695 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l696
															}
															position++
															goto l695
														l696:
															position, tokenIndex = position695, tokenIndex695
															if buffer[position] != rune('C') {
																goto l611
															}
															position++
														}
													l695:
														{
															position697, tokenIndex697 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l698
															}
															position++
															goto l697
														l698:
															position, tokenIndex = position697, tokenIndex697
															if buffer[position] != rune('C') {
																goto l611
															}
															position++
														}
													l697:
														{
															position699, tokenIndex699 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l700
															}
															position++
															goto l699
														l700:
															position, tokenIndex = position699, tokenIndex699
															if buffer[position] != rune('F') {
																goto l611
															}
															position++
														}
													l699:
														add(rulePegText, position694)
													}
													{
														add(ruleAction67, position)
													}
													add(ruleCcf, position693)
												}
												break
											case 'S', 's':
												{
													position702 := position
													{
														position703 := position
														{
															position704, tokenIndex704 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l705
															}
															position++
															goto l704
														l705:
															position, tokenIndex = position704, tokenIndex704
															if buffer[position] != rune('S') {
																goto l611
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
																goto l611
															}
															position++
														}
													l706:
														{
															position708, tokenIndex708 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l709
															}
															position++
															goto l708
														l709:
															position, tokenIndex = position708, tokenIndex708
															if buffer[position] != rune('F') {
																goto l611
															}
															position++
														}
													l708:
														add(rulePegText, position703)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleScf, position702)
												}
												break
											case 'R', 'r':
												{
													position711 := position
													{
														position712 := position
														{
															position713, tokenIndex713 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l714
															}
															position++
															goto l713
														l714:
															position, tokenIndex = position713, tokenIndex713
															if buffer[position] != rune('R') {
																goto l611
															}
															position++
														}
													l713:
														{
															position715, tokenIndex715 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l716
															}
															position++
															goto l715
														l716:
															position, tokenIndex = position715, tokenIndex715
															if buffer[position] != rune('R') {
																goto l611
															}
															position++
														}
													l715:
														{
															position717, tokenIndex717 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l718
															}
															position++
															goto l717
														l718:
															position, tokenIndex = position717, tokenIndex717
															if buffer[position] != rune('A') {
																goto l611
															}
															position++
														}
													l717:
														add(rulePegText, position712)
													}
													{
														add(ruleAction63, position)
													}
													add(ruleRra, position711)
												}
												break
											case 'H', 'h':
												{
													position720 := position
													{
														position721 := position
														{
															position722, tokenIndex722 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l723
															}
															position++
															goto l722
														l723:
															position, tokenIndex = position722, tokenIndex722
															if buffer[position] != rune('H') {
																goto l611
															}
															position++
														}
													l722:
														{
															position724, tokenIndex724 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l725
															}
															position++
															goto l724
														l725:
															position, tokenIndex = position724, tokenIndex724
															if buffer[position] != rune('A') {
																goto l611
															}
															position++
														}
													l724:
														{
															position726, tokenIndex726 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l727
															}
															position++
															goto l726
														l727:
															position, tokenIndex = position726, tokenIndex726
															if buffer[position] != rune('L') {
																goto l611
															}
															position++
														}
													l726:
														{
															position728, tokenIndex728 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l729
															}
															position++
															goto l728
														l729:
															position, tokenIndex = position728, tokenIndex728
															if buffer[position] != rune('T') {
																goto l611
															}
															position++
														}
													l728:
														add(rulePegText, position721)
													}
													{
														add(ruleAction59, position)
													}
													add(ruleHalt, position720)
												}
												break
											default:
												{
													position731 := position
													{
														position732 := position
														{
															position733, tokenIndex733 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l734
															}
															position++
															goto l733
														l734:
															position, tokenIndex = position733, tokenIndex733
															if buffer[position] != rune('N') {
																goto l611
															}
															position++
														}
													l733:
														{
															position735, tokenIndex735 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l736
															}
															position++
															goto l735
														l736:
															position, tokenIndex = position735, tokenIndex735
															if buffer[position] != rune('O') {
																goto l611
															}
															position++
														}
													l735:
														{
															position737, tokenIndex737 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l738
															}
															position++
															goto l737
														l738:
															position, tokenIndex = position737, tokenIndex737
															if buffer[position] != rune('P') {
																goto l611
															}
															position++
														}
													l737:
														add(rulePegText, position732)
													}
													{
														add(ruleAction58, position)
													}
													add(ruleNop, position731)
												}
												break
											}
										}

									}
								l613:
									add(ruleSimple, position612)
								}
								goto l21
							l611:
								position, tokenIndex = position21, tokenIndex21
								{
									position741 := position
									{
										position742, tokenIndex742 := position, tokenIndex
										{
											position744 := position
											{
												position745, tokenIndex745 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l746
												}
												position++
												goto l745
											l746:
												position, tokenIndex = position745, tokenIndex745
												if buffer[position] != rune('R') {
													goto l743
												}
												position++
											}
										l745:
											{
												position747, tokenIndex747 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l748
												}
												position++
												goto l747
											l748:
												position, tokenIndex = position747, tokenIndex747
												if buffer[position] != rune('S') {
													goto l743
												}
												position++
											}
										l747:
											{
												position749, tokenIndex749 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l750
												}
												position++
												goto l749
											l750:
												position, tokenIndex = position749, tokenIndex749
												if buffer[position] != rune('T') {
													goto l743
												}
												position++
											}
										l749:
											if !_rules[rulews]() {
												goto l743
											}
											if !_rules[rulen]() {
												goto l743
											}
											{
												add(ruleAction95, position)
											}
											add(ruleRst, position744)
										}
										goto l742
									l743:
										position, tokenIndex = position742, tokenIndex742
										{
											position753 := position
											{
												position754, tokenIndex754 := position, tokenIndex
												if buffer[position] != rune('j') {
													goto l755
												}
												position++
												goto l754
											l755:
												position, tokenIndex = position754, tokenIndex754
												if buffer[position] != rune('J') {
													goto l752
												}
												position++
											}
										l754:
											{
												position756, tokenIndex756 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l757
												}
												position++
												goto l756
											l757:
												position, tokenIndex = position756, tokenIndex756
												if buffer[position] != rune('P') {
													goto l752
												}
												position++
											}
										l756:
											if !_rules[rulews]() {
												goto l752
											}
											{
												position758, tokenIndex758 := position, tokenIndex
												if !_rules[rulecc]() {
													goto l758
												}
												if !_rules[rulesep]() {
													goto l758
												}
												goto l759
											l758:
												position, tokenIndex = position758, tokenIndex758
											}
										l759:
											if !_rules[ruleSrc16]() {
												goto l752
											}
											{
												add(ruleAction98, position)
											}
											add(ruleJp, position753)
										}
										goto l742
									l752:
										position, tokenIndex = position742, tokenIndex742
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position762 := position
													{
														position763, tokenIndex763 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l764
														}
														position++
														goto l763
													l764:
														position, tokenIndex = position763, tokenIndex763
														if buffer[position] != rune('D') {
															goto l740
														}
														position++
													}
												l763:
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
															goto l740
														}
														position++
													}
												l765:
													{
														position767, tokenIndex767 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l768
														}
														position++
														goto l767
													l768:
														position, tokenIndex = position767, tokenIndex767
														if buffer[position] != rune('N') {
															goto l740
														}
														position++
													}
												l767:
													{
														position769, tokenIndex769 := position, tokenIndex
														if buffer[position] != rune('z') {
															goto l770
														}
														position++
														goto l769
													l770:
														position, tokenIndex = position769, tokenIndex769
														if buffer[position] != rune('Z') {
															goto l740
														}
														position++
													}
												l769:
													if !_rules[rulews]() {
														goto l740
													}
													if !_rules[ruledisp]() {
														goto l740
													}
													{
														add(ruleAction100, position)
													}
													add(ruleDjnz, position762)
												}
												break
											case 'J', 'j':
												{
													position772 := position
													{
														position773, tokenIndex773 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l774
														}
														position++
														goto l773
													l774:
														position, tokenIndex = position773, tokenIndex773
														if buffer[position] != rune('J') {
															goto l740
														}
														position++
													}
												l773:
													{
														position775, tokenIndex775 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l776
														}
														position++
														goto l775
													l776:
														position, tokenIndex = position775, tokenIndex775
														if buffer[position] != rune('R') {
															goto l740
														}
														position++
													}
												l775:
													if !_rules[rulews]() {
														goto l740
													}
													{
														position777, tokenIndex777 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l777
														}
														if !_rules[rulesep]() {
															goto l777
														}
														goto l778
													l777:
														position, tokenIndex = position777, tokenIndex777
													}
												l778:
													if !_rules[ruledisp]() {
														goto l740
													}
													{
														add(ruleAction99, position)
													}
													add(ruleJr, position772)
												}
												break
											case 'R', 'r':
												{
													position780 := position
													{
														position781, tokenIndex781 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l782
														}
														position++
														goto l781
													l782:
														position, tokenIndex = position781, tokenIndex781
														if buffer[position] != rune('R') {
															goto l740
														}
														position++
													}
												l781:
													{
														position783, tokenIndex783 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l784
														}
														position++
														goto l783
													l784:
														position, tokenIndex = position783, tokenIndex783
														if buffer[position] != rune('E') {
															goto l740
														}
														position++
													}
												l783:
													{
														position785, tokenIndex785 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l786
														}
														position++
														goto l785
													l786:
														position, tokenIndex = position785, tokenIndex785
														if buffer[position] != rune('T') {
															goto l740
														}
														position++
													}
												l785:
													{
														position787, tokenIndex787 := position, tokenIndex
														if !_rules[rulews]() {
															goto l787
														}
														if !_rules[rulecc]() {
															goto l787
														}
														goto l788
													l787:
														position, tokenIndex = position787, tokenIndex787
													}
												l788:
													{
														add(ruleAction97, position)
													}
													add(ruleRet, position780)
												}
												break
											default:
												{
													position790 := position
													{
														position791, tokenIndex791 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l792
														}
														position++
														goto l791
													l792:
														position, tokenIndex = position791, tokenIndex791
														if buffer[position] != rune('C') {
															goto l740
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
															goto l740
														}
														position++
													}
												l793:
													{
														position795, tokenIndex795 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l796
														}
														position++
														goto l795
													l796:
														position, tokenIndex = position795, tokenIndex795
														if buffer[position] != rune('L') {
															goto l740
														}
														position++
													}
												l795:
													{
														position797, tokenIndex797 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l798
														}
														position++
														goto l797
													l798:
														position, tokenIndex = position797, tokenIndex797
														if buffer[position] != rune('L') {
															goto l740
														}
														position++
													}
												l797:
													if !_rules[rulews]() {
														goto l740
													}
													{
														position799, tokenIndex799 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l799
														}
														if !_rules[rulesep]() {
															goto l799
														}
														goto l800
													l799:
														position, tokenIndex = position799, tokenIndex799
													}
												l800:
													if !_rules[ruleSrc16]() {
														goto l740
													}
													{
														add(ruleAction96, position)
													}
													add(ruleCall, position790)
												}
												break
											}
										}

									}
								l742:
									add(ruleJump, position741)
								}
								goto l21
							l740:
								position, tokenIndex = position21, tokenIndex21
								{
									position802 := position
									{
										position803, tokenIndex803 := position, tokenIndex
										{
											position805 := position
											{
												position806, tokenIndex806 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l807
												}
												position++
												goto l806
											l807:
												position, tokenIndex = position806, tokenIndex806
												if buffer[position] != rune('I') {
													goto l804
												}
												position++
											}
										l806:
											{
												position808, tokenIndex808 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l809
												}
												position++
												goto l808
											l809:
												position, tokenIndex = position808, tokenIndex808
												if buffer[position] != rune('N') {
													goto l804
												}
												position++
											}
										l808:
											if !_rules[rulews]() {
												goto l804
											}
											if !_rules[ruleReg8]() {
												goto l804
											}
											if !_rules[rulesep]() {
												goto l804
											}
											if !_rules[rulePort]() {
												goto l804
											}
											{
												add(ruleAction101, position)
											}
											add(ruleIN, position805)
										}
										goto l803
									l804:
										position, tokenIndex = position803, tokenIndex803
										{
											position811 := position
											{
												position812, tokenIndex812 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l813
												}
												position++
												goto l812
											l813:
												position, tokenIndex = position812, tokenIndex812
												if buffer[position] != rune('O') {
													goto l18
												}
												position++
											}
										l812:
											{
												position814, tokenIndex814 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l815
												}
												position++
												goto l814
											l815:
												position, tokenIndex = position814, tokenIndex814
												if buffer[position] != rune('U') {
													goto l18
												}
												position++
											}
										l814:
											{
												position816, tokenIndex816 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l817
												}
												position++
												goto l816
											l817:
												position, tokenIndex = position816, tokenIndex816
												if buffer[position] != rune('T') {
													goto l18
												}
												position++
											}
										l816:
											if !_rules[rulews]() {
												goto l18
											}
											if !_rules[rulePort]() {
												goto l18
											}
											if !_rules[rulesep]() {
												goto l18
											}
											if !_rules[ruleReg8]() {
												goto l18
											}
											{
												add(ruleAction102, position)
											}
											add(ruleOUT, position811)
										}
									}
								l803:
									add(ruleIO, position802)
								}
							}
						l21:
							add(ruleInstruction, position20)
						}
						goto l19
					l18:
						position, tokenIndex = position18, tokenIndex18
					}
				l19:
					{
						position819 := position
						{
							position820, tokenIndex820 := position, tokenIndex
							if !_rules[rulews]() {
								goto l820
							}
							goto l821
						l820:
							position, tokenIndex = position820, tokenIndex820
						}
					l821:
						{
							position822, tokenIndex822 := position, tokenIndex
							{
								position824 := position
								{
									position825, tokenIndex825 := position, tokenIndex
									if buffer[position] != rune(';') {
										goto l826
									}
									position++
									goto l825
								l826:
									position, tokenIndex = position825, tokenIndex825
									if buffer[position] != rune('#') {
										goto l822
									}
									position++
								}
							l825:
							l827:
								{
									position828, tokenIndex828 := position, tokenIndex
									{
										position829, tokenIndex829 := position, tokenIndex
										if buffer[position] != rune('\n') {
											goto l829
										}
										position++
										goto l828
									l829:
										position, tokenIndex = position829, tokenIndex829
									}
									if !matchDot() {
										goto l828
									}
									goto l827
								l828:
									position, tokenIndex = position828, tokenIndex828
								}
								add(ruleComment, position824)
							}
							goto l823
						l822:
							position, tokenIndex = position822, tokenIndex822
						}
					l823:
						{
							position830, tokenIndex830 := position, tokenIndex
							if !_rules[rulews]() {
								goto l830
							}
							goto l831
						l830:
							position, tokenIndex = position830, tokenIndex830
						}
					l831:
						{
							position832, tokenIndex832 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l833
							}
							position++
							goto l832
						l833:
							position, tokenIndex = position832, tokenIndex832
							if buffer[position] != rune(':') {
								goto l0
							}
							position++
						}
					l832:
						add(ruleLineEnd, position819)
					}
					{
						add(ruleAction0, position)
					}
					add(ruleLine, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position835 := position
						{
							position836, tokenIndex836 := position, tokenIndex
							{
								position838 := position
								{
									position839 := position
									if !_rules[rulealpha]() {
										goto l836
									}
								l840:
									{
										position841, tokenIndex841 := position, tokenIndex
										{
											position842 := position
											{
												position843, tokenIndex843 := position, tokenIndex
												if !_rules[rulealpha]() {
													goto l844
												}
												goto l843
											l844:
												position, tokenIndex = position843, tokenIndex843
												{
													position845 := position
													if c := buffer[position]; c < rune('0') || c > rune('9') {
														goto l841
													}
													position++
													add(rulenum, position845)
												}
											}
										l843:
											add(rulealphanum, position842)
										}
										goto l840
									l841:
										position, tokenIndex = position841, tokenIndex841
									}
									add(rulePegText, position839)
								}
								if buffer[position] != rune(':') {
									goto l836
								}
								position++
								{
									add(ruleAction1, position)
								}
								add(ruleLabel, position838)
							}
							goto l837
						l836:
							position, tokenIndex = position836, tokenIndex836
						}
					l837:
					l847:
						{
							position848, tokenIndex848 := position, tokenIndex
							if !_rules[rulews]() {
								goto l848
							}
							goto l847
						l848:
							position, tokenIndex = position848, tokenIndex848
						}
						{
							position849, tokenIndex849 := position, tokenIndex
							{
								position851 := position
								{
									position852, tokenIndex852 := position, tokenIndex
									{
										position854 := position
										{
											position855, tokenIndex855 := position, tokenIndex
											{
												position857 := position
												{
													position858, tokenIndex858 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l859
													}
													position++
													goto l858
												l859:
													position, tokenIndex = position858, tokenIndex858
													if buffer[position] != rune('P') {
														goto l856
													}
													position++
												}
											l858:
												{
													position860, tokenIndex860 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l861
													}
													position++
													goto l860
												l861:
													position, tokenIndex = position860, tokenIndex860
													if buffer[position] != rune('U') {
														goto l856
													}
													position++
												}
											l860:
												{
													position862, tokenIndex862 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l863
													}
													position++
													goto l862
												l863:
													position, tokenIndex = position862, tokenIndex862
													if buffer[position] != rune('S') {
														goto l856
													}
													position++
												}
											l862:
												{
													position864, tokenIndex864 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l865
													}
													position++
													goto l864
												l865:
													position, tokenIndex = position864, tokenIndex864
													if buffer[position] != rune('H') {
														goto l856
													}
													position++
												}
											l864:
												if !_rules[rulews]() {
													goto l856
												}
												if !_rules[ruleSrc16]() {
													goto l856
												}
												{
													add(ruleAction4, position)
												}
												add(rulePush, position857)
											}
											goto l855
										l856:
											position, tokenIndex = position855, tokenIndex855
											{
												switch buffer[position] {
												case 'E', 'e':
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
																goto l853
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
																goto l853
															}
															position++
														}
													l871:
														if !_rules[rulews]() {
															goto l853
														}
														if !_rules[ruleDst16]() {
															goto l853
														}
														if !_rules[rulesep]() {
															goto l853
														}
														if !_rules[ruleSrc16]() {
															goto l853
														}
														{
															add(ruleAction6, position)
														}
														add(ruleEx, position868)
													}
													break
												case 'P', 'p':
													{
														position874 := position
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
																goto l853
															}
															position++
														}
													l875:
														{
															position877, tokenIndex877 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l878
															}
															position++
															goto l877
														l878:
															position, tokenIndex = position877, tokenIndex877
															if buffer[position] != rune('O') {
																goto l853
															}
															position++
														}
													l877:
														{
															position879, tokenIndex879 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l880
															}
															position++
															goto l879
														l880:
															position, tokenIndex = position879, tokenIndex879
															if buffer[position] != rune('P') {
																goto l853
															}
															position++
														}
													l879:
														if !_rules[rulews]() {
															goto l853
														}
														if !_rules[ruleDst16]() {
															goto l853
														}
														{
															add(ruleAction5, position)
														}
														add(rulePop, position874)
													}
													break
												default:
													{
														position882 := position
														{
															position883, tokenIndex883 := position, tokenIndex
															{
																position885 := position
																{
																	position886, tokenIndex886 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l887
																	}
																	position++
																	goto l886
																l887:
																	position, tokenIndex = position886, tokenIndex886
																	if buffer[position] != rune('L') {
																		goto l884
																	}
																	position++
																}
															l886:
																{
																	position888, tokenIndex888 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l889
																	}
																	position++
																	goto l888
																l889:
																	position, tokenIndex = position888, tokenIndex888
																	if buffer[position] != rune('D') {
																		goto l884
																	}
																	position++
																}
															l888:
																if !_rules[rulews]() {
																	goto l884
																}
																if !_rules[ruleDst16]() {
																	goto l884
																}
																if !_rules[rulesep]() {
																	goto l884
																}
																if !_rules[ruleSrc16]() {
																	goto l884
																}
																{
																	add(ruleAction3, position)
																}
																add(ruleLoad16, position885)
															}
															goto l883
														l884:
															position, tokenIndex = position883, tokenIndex883
															{
																position891 := position
																{
																	position892, tokenIndex892 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l893
																	}
																	position++
																	goto l892
																l893:
																	position, tokenIndex = position892, tokenIndex892
																	if buffer[position] != rune('L') {
																		goto l853
																	}
																	position++
																}
															l892:
																{
																	position894, tokenIndex894 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l895
																	}
																	position++
																	goto l894
																l895:
																	position, tokenIndex = position894, tokenIndex894
																	if buffer[position] != rune('D') {
																		goto l853
																	}
																	position++
																}
															l894:
																if !_rules[rulews]() {
																	goto l853
																}
																{
																	position896 := position
																	{
																		position897, tokenIndex897 := position, tokenIndex
																		if !_rules[ruleReg8]() {
																			goto l898
																		}
																		goto l897
																	l898:
																		position, tokenIndex = position897, tokenIndex897
																		if !_rules[ruleReg16Contents]() {
																			goto l899
																		}
																		goto l897
																	l899:
																		position, tokenIndex = position897, tokenIndex897
																		if !_rules[rulenn_contents]() {
																			goto l853
																		}
																	}
																l897:
																	{
																		add(ruleAction16, position)
																	}
																	add(ruleDst8, position896)
																}
																if !_rules[rulesep]() {
																	goto l853
																}
																if !_rules[ruleSrc8]() {
																	goto l853
																}
																{
																	add(ruleAction2, position)
																}
																add(ruleLoad8, position891)
															}
														}
													l883:
														add(ruleLoad, position882)
													}
													break
												}
											}

										}
									l855:
										add(ruleAssignment, position854)
									}
									goto l852
								l853:
									position, tokenIndex = position852, tokenIndex852
									{
										position903 := position
										{
											position904, tokenIndex904 := position, tokenIndex
											{
												position906 := position
												{
													position907, tokenIndex907 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l908
													}
													position++
													goto l907
												l908:
													position, tokenIndex = position907, tokenIndex907
													if buffer[position] != rune('I') {
														goto l905
													}
													position++
												}
											l907:
												{
													position909, tokenIndex909 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l910
													}
													position++
													goto l909
												l910:
													position, tokenIndex = position909, tokenIndex909
													if buffer[position] != rune('N') {
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
													add(ruleAction7, position)
												}
												add(ruleInc16Indexed8, position906)
											}
											goto l904
										l905:
											position, tokenIndex = position904, tokenIndex904
											{
												position915 := position
												{
													position916, tokenIndex916 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l917
													}
													position++
													goto l916
												l917:
													position, tokenIndex = position916, tokenIndex916
													if buffer[position] != rune('I') {
														goto l914
													}
													position++
												}
											l916:
												{
													position918, tokenIndex918 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l919
													}
													position++
													goto l918
												l919:
													position, tokenIndex = position918, tokenIndex918
													if buffer[position] != rune('N') {
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
													add(ruleAction9, position)
												}
												add(ruleInc16, position915)
											}
											goto l904
										l914:
											position, tokenIndex = position904, tokenIndex904
											{
												position923 := position
												{
													position924, tokenIndex924 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l925
													}
													position++
													goto l924
												l925:
													position, tokenIndex = position924, tokenIndex924
													if buffer[position] != rune('I') {
														goto l902
													}
													position++
												}
											l924:
												{
													position926, tokenIndex926 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l927
													}
													position++
													goto l926
												l927:
													position, tokenIndex = position926, tokenIndex926
													if buffer[position] != rune('N') {
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
													add(ruleAction8, position)
												}
												add(ruleInc8, position923)
											}
										}
									l904:
										add(ruleInc, position903)
									}
									goto l852
								l902:
									position, tokenIndex = position852, tokenIndex852
									{
										position932 := position
										{
											position933, tokenIndex933 := position, tokenIndex
											{
												position935 := position
												{
													position936, tokenIndex936 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l937
													}
													position++
													goto l936
												l937:
													position, tokenIndex = position936, tokenIndex936
													if buffer[position] != rune('D') {
														goto l934
													}
													position++
												}
											l936:
												{
													position938, tokenIndex938 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l939
													}
													position++
													goto l938
												l939:
													position, tokenIndex = position938, tokenIndex938
													if buffer[position] != rune('E') {
														goto l934
													}
													position++
												}
											l938:
												{
													position940, tokenIndex940 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l941
													}
													position++
													goto l940
												l941:
													position, tokenIndex = position940, tokenIndex940
													if buffer[position] != rune('C') {
														goto l934
													}
													position++
												}
											l940:
												if !_rules[rulews]() {
													goto l934
												}
												if !_rules[ruleILoc8]() {
													goto l934
												}
												{
													add(ruleAction10, position)
												}
												add(ruleDec16Indexed8, position935)
											}
											goto l933
										l934:
											position, tokenIndex = position933, tokenIndex933
											{
												position944 := position
												{
													position945, tokenIndex945 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l946
													}
													position++
													goto l945
												l946:
													position, tokenIndex = position945, tokenIndex945
													if buffer[position] != rune('D') {
														goto l943
													}
													position++
												}
											l945:
												{
													position947, tokenIndex947 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l948
													}
													position++
													goto l947
												l948:
													position, tokenIndex = position947, tokenIndex947
													if buffer[position] != rune('E') {
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
												if !_rules[ruleLoc16]() {
													goto l943
												}
												{
													add(ruleAction12, position)
												}
												add(ruleDec16, position944)
											}
											goto l933
										l943:
											position, tokenIndex = position933, tokenIndex933
											{
												position952 := position
												{
													position953, tokenIndex953 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l954
													}
													position++
													goto l953
												l954:
													position, tokenIndex = position953, tokenIndex953
													if buffer[position] != rune('D') {
														goto l931
													}
													position++
												}
											l953:
												{
													position955, tokenIndex955 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l956
													}
													position++
													goto l955
												l956:
													position, tokenIndex = position955, tokenIndex955
													if buffer[position] != rune('E') {
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
												if !_rules[ruleLoc8]() {
													goto l931
												}
												{
													add(ruleAction11, position)
												}
												add(ruleDec8, position952)
											}
										}
									l933:
										add(ruleDec, position932)
									}
									goto l852
								l931:
									position, tokenIndex = position852, tokenIndex852
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
												if !_rules[ruleDst16]() {
													goto l963
												}
												if !_rules[rulesep]() {
													goto l963
												}
												if !_rules[ruleSrc16]() {
													goto l963
												}
												{
													add(ruleAction13, position)
												}
												add(ruleAdd16, position964)
											}
											goto l962
										l963:
											position, tokenIndex = position962, tokenIndex962
											{
												position973 := position
												{
													position974, tokenIndex974 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l975
													}
													position++
													goto l974
												l975:
													position, tokenIndex = position974, tokenIndex974
													if buffer[position] != rune('A') {
														goto l972
													}
													position++
												}
											l974:
												{
													position976, tokenIndex976 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l977
													}
													position++
													goto l976
												l977:
													position, tokenIndex = position976, tokenIndex976
													if buffer[position] != rune('D') {
														goto l972
													}
													position++
												}
											l976:
												{
													position978, tokenIndex978 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l979
													}
													position++
													goto l978
												l979:
													position, tokenIndex = position978, tokenIndex978
													if buffer[position] != rune('C') {
														goto l972
													}
													position++
												}
											l978:
												if !_rules[rulews]() {
													goto l972
												}
												if !_rules[ruleDst16]() {
													goto l972
												}
												if !_rules[rulesep]() {
													goto l972
												}
												if !_rules[ruleSrc16]() {
													goto l972
												}
												{
													add(ruleAction14, position)
												}
												add(ruleAdc16, position973)
											}
											goto l962
										l972:
											position, tokenIndex = position962, tokenIndex962
											{
												position981 := position
												{
													position982, tokenIndex982 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l983
													}
													position++
													goto l982
												l983:
													position, tokenIndex = position982, tokenIndex982
													if buffer[position] != rune('S') {
														goto l960
													}
													position++
												}
											l982:
												{
													position984, tokenIndex984 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l985
													}
													position++
													goto l984
												l985:
													position, tokenIndex = position984, tokenIndex984
													if buffer[position] != rune('B') {
														goto l960
													}
													position++
												}
											l984:
												{
													position986, tokenIndex986 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l987
													}
													position++
													goto l986
												l987:
													position, tokenIndex = position986, tokenIndex986
													if buffer[position] != rune('C') {
														goto l960
													}
													position++
												}
											l986:
												if !_rules[rulews]() {
													goto l960
												}
												if !_rules[ruleDst16]() {
													goto l960
												}
												if !_rules[rulesep]() {
													goto l960
												}
												if !_rules[ruleSrc16]() {
													goto l960
												}
												{
													add(ruleAction15, position)
												}
												add(ruleSbc16, position981)
											}
										}
									l962:
										add(ruleAlu16, position961)
									}
									goto l852
								l960:
									position, tokenIndex = position852, tokenIndex852
									{
										position990 := position
										{
											position991, tokenIndex991 := position, tokenIndex
											{
												position993 := position
												{
													position994, tokenIndex994 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l995
													}
													position++
													goto l994
												l995:
													position, tokenIndex = position994, tokenIndex994
													if buffer[position] != rune('A') {
														goto l992
													}
													position++
												}
											l994:
												{
													position996, tokenIndex996 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l997
													}
													position++
													goto l996
												l997:
													position, tokenIndex = position996, tokenIndex996
													if buffer[position] != rune('D') {
														goto l992
													}
													position++
												}
											l996:
												{
													position998, tokenIndex998 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l999
													}
													position++
													goto l998
												l999:
													position, tokenIndex = position998, tokenIndex998
													if buffer[position] != rune('D') {
														goto l992
													}
													position++
												}
											l998:
												if !_rules[rulews]() {
													goto l992
												}
												{
													position1000, tokenIndex1000 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l1001
													}
													position++
													goto l1000
												l1001:
													position, tokenIndex = position1000, tokenIndex1000
													if buffer[position] != rune('A') {
														goto l992
													}
													position++
												}
											l1000:
												if !_rules[rulesep]() {
													goto l992
												}
												if !_rules[ruleSrc8]() {
													goto l992
												}
												{
													add(ruleAction39, position)
												}
												add(ruleAdd, position993)
											}
											goto l991
										l992:
											position, tokenIndex = position991, tokenIndex991
											{
												position1004 := position
												{
													position1005, tokenIndex1005 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l1006
													}
													position++
													goto l1005
												l1006:
													position, tokenIndex = position1005, tokenIndex1005
													if buffer[position] != rune('A') {
														goto l1003
													}
													position++
												}
											l1005:
												{
													position1007, tokenIndex1007 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l1008
													}
													position++
													goto l1007
												l1008:
													position, tokenIndex = position1007, tokenIndex1007
													if buffer[position] != rune('D') {
														goto l1003
													}
													position++
												}
											l1007:
												{
													position1009, tokenIndex1009 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l1010
													}
													position++
													goto l1009
												l1010:
													position, tokenIndex = position1009, tokenIndex1009
													if buffer[position] != rune('C') {
														goto l1003
													}
													position++
												}
											l1009:
												if !_rules[rulews]() {
													goto l1003
												}
												{
													position1011, tokenIndex1011 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l1012
													}
													position++
													goto l1011
												l1012:
													position, tokenIndex = position1011, tokenIndex1011
													if buffer[position] != rune('A') {
														goto l1003
													}
													position++
												}
											l1011:
												if !_rules[rulesep]() {
													goto l1003
												}
												if !_rules[ruleSrc8]() {
													goto l1003
												}
												{
													add(ruleAction40, position)
												}
												add(ruleAdc, position1004)
											}
											goto l991
										l1003:
											position, tokenIndex = position991, tokenIndex991
											{
												position1015 := position
												{
													position1016, tokenIndex1016 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1017
													}
													position++
													goto l1016
												l1017:
													position, tokenIndex = position1016, tokenIndex1016
													if buffer[position] != rune('S') {
														goto l1014
													}
													position++
												}
											l1016:
												{
													position1018, tokenIndex1018 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1019
													}
													position++
													goto l1018
												l1019:
													position, tokenIndex = position1018, tokenIndex1018
													if buffer[position] != rune('U') {
														goto l1014
													}
													position++
												}
											l1018:
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
														goto l1014
													}
													position++
												}
											l1020:
												if !_rules[rulews]() {
													goto l1014
												}
												if !_rules[ruleSrc8]() {
													goto l1014
												}
												{
													add(ruleAction41, position)
												}
												add(ruleSub, position1015)
											}
											goto l991
										l1014:
											position, tokenIndex = position991, tokenIndex991
											{
												switch buffer[position] {
												case 'C', 'c':
													{
														position1024 := position
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
																goto l989
															}
															position++
														}
													l1025:
														{
															position1027, tokenIndex1027 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l1028
															}
															position++
															goto l1027
														l1028:
															position, tokenIndex = position1027, tokenIndex1027
															if buffer[position] != rune('P') {
																goto l989
															}
															position++
														}
													l1027:
														if !_rules[rulews]() {
															goto l989
														}
														if !_rules[ruleSrc8]() {
															goto l989
														}
														{
															add(ruleAction46, position)
														}
														add(ruleCp, position1024)
													}
													break
												case 'O', 'o':
													{
														position1030 := position
														{
															position1031, tokenIndex1031 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1032
															}
															position++
															goto l1031
														l1032:
															position, tokenIndex = position1031, tokenIndex1031
															if buffer[position] != rune('O') {
																goto l989
															}
															position++
														}
													l1031:
														{
															position1033, tokenIndex1033 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1034
															}
															position++
															goto l1033
														l1034:
															position, tokenIndex = position1033, tokenIndex1033
															if buffer[position] != rune('R') {
																goto l989
															}
															position++
														}
													l1033:
														if !_rules[rulews]() {
															goto l989
														}
														if !_rules[ruleSrc8]() {
															goto l989
														}
														{
															add(ruleAction45, position)
														}
														add(ruleOr, position1030)
													}
													break
												case 'X', 'x':
													{
														position1036 := position
														{
															position1037, tokenIndex1037 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l1038
															}
															position++
															goto l1037
														l1038:
															position, tokenIndex = position1037, tokenIndex1037
															if buffer[position] != rune('X') {
																goto l989
															}
															position++
														}
													l1037:
														{
															position1039, tokenIndex1039 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1040
															}
															position++
															goto l1039
														l1040:
															position, tokenIndex = position1039, tokenIndex1039
															if buffer[position] != rune('O') {
																goto l989
															}
															position++
														}
													l1039:
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
																goto l989
															}
															position++
														}
													l1041:
														if !_rules[rulews]() {
															goto l989
														}
														if !_rules[ruleSrc8]() {
															goto l989
														}
														{
															add(ruleAction44, position)
														}
														add(ruleXor, position1036)
													}
													break
												case 'A', 'a':
													{
														position1044 := position
														{
															position1045, tokenIndex1045 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1046
															}
															position++
															goto l1045
														l1046:
															position, tokenIndex = position1045, tokenIndex1045
															if buffer[position] != rune('A') {
																goto l989
															}
															position++
														}
													l1045:
														{
															position1047, tokenIndex1047 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1048
															}
															position++
															goto l1047
														l1048:
															position, tokenIndex = position1047, tokenIndex1047
															if buffer[position] != rune('N') {
																goto l989
															}
															position++
														}
													l1047:
														{
															position1049, tokenIndex1049 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1050
															}
															position++
															goto l1049
														l1050:
															position, tokenIndex = position1049, tokenIndex1049
															if buffer[position] != rune('D') {
																goto l989
															}
															position++
														}
													l1049:
														if !_rules[rulews]() {
															goto l989
														}
														if !_rules[ruleSrc8]() {
															goto l989
														}
														{
															add(ruleAction43, position)
														}
														add(ruleAnd, position1044)
													}
													break
												default:
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
																goto l989
															}
															position++
														}
													l1053:
														{
															position1055, tokenIndex1055 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1056
															}
															position++
															goto l1055
														l1056:
															position, tokenIndex = position1055, tokenIndex1055
															if buffer[position] != rune('B') {
																goto l989
															}
															position++
														}
													l1055:
														{
															position1057, tokenIndex1057 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1058
															}
															position++
															goto l1057
														l1058:
															position, tokenIndex = position1057, tokenIndex1057
															if buffer[position] != rune('C') {
																goto l989
															}
															position++
														}
													l1057:
														if !_rules[rulews]() {
															goto l989
														}
														{
															position1059, tokenIndex1059 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1060
															}
															position++
															goto l1059
														l1060:
															position, tokenIndex = position1059, tokenIndex1059
															if buffer[position] != rune('A') {
																goto l989
															}
															position++
														}
													l1059:
														if !_rules[rulesep]() {
															goto l989
														}
														if !_rules[ruleSrc8]() {
															goto l989
														}
														{
															add(ruleAction42, position)
														}
														add(ruleSbc, position1052)
													}
													break
												}
											}

										}
									l991:
										add(ruleAlu, position990)
									}
									goto l852
								l989:
									position, tokenIndex = position852, tokenIndex852
									{
										position1063 := position
										{
											position1064, tokenIndex1064 := position, tokenIndex
											{
												position1066 := position
												{
													position1067, tokenIndex1067 := position, tokenIndex
													{
														position1069 := position
														{
															position1070, tokenIndex1070 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1071
															}
															position++
															goto l1070
														l1071:
															position, tokenIndex = position1070, tokenIndex1070
															if buffer[position] != rune('R') {
																goto l1068
															}
															position++
														}
													l1070:
														{
															position1072, tokenIndex1072 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1073
															}
															position++
															goto l1072
														l1073:
															position, tokenIndex = position1072, tokenIndex1072
															if buffer[position] != rune('L') {
																goto l1068
															}
															position++
														}
													l1072:
														{
															position1074, tokenIndex1074 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1075
															}
															position++
															goto l1074
														l1075:
															position, tokenIndex = position1074, tokenIndex1074
															if buffer[position] != rune('C') {
																goto l1068
															}
															position++
														}
													l1074:
														if !_rules[rulews]() {
															goto l1068
														}
														if !_rules[ruleLoc8]() {
															goto l1068
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
															add(ruleAction47, position)
														}
														add(ruleRlc, position1069)
													}
													goto l1067
												l1068:
													position, tokenIndex = position1067, tokenIndex1067
													{
														position1080 := position
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
																goto l1079
															}
															position++
														}
													l1081:
														{
															position1083, tokenIndex1083 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1084
															}
															position++
															goto l1083
														l1084:
															position, tokenIndex = position1083, tokenIndex1083
															if buffer[position] != rune('R') {
																goto l1079
															}
															position++
														}
													l1083:
														{
															position1085, tokenIndex1085 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1086
															}
															position++
															goto l1085
														l1086:
															position, tokenIndex = position1085, tokenIndex1085
															if buffer[position] != rune('C') {
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
															add(ruleAction48, position)
														}
														add(ruleRrc, position1080)
													}
													goto l1067
												l1079:
													position, tokenIndex = position1067, tokenIndex1067
													{
														position1091 := position
														{
															position1092, tokenIndex1092 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1093
															}
															position++
															goto l1092
														l1093:
															position, tokenIndex = position1092, tokenIndex1092
															if buffer[position] != rune('R') {
																goto l1090
															}
															position++
														}
													l1092:
														{
															position1094, tokenIndex1094 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1095
															}
															position++
															goto l1094
														l1095:
															position, tokenIndex = position1094, tokenIndex1094
															if buffer[position] != rune('L') {
																goto l1090
															}
															position++
														}
													l1094:
														if !_rules[rulews]() {
															goto l1090
														}
														if !_rules[ruleLoc8]() {
															goto l1090
														}
														{
															position1096, tokenIndex1096 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1096
															}
															if !_rules[ruleCopy8]() {
																goto l1096
															}
															goto l1097
														l1096:
															position, tokenIndex = position1096, tokenIndex1096
														}
													l1097:
														{
															add(ruleAction49, position)
														}
														add(ruleRl, position1091)
													}
													goto l1067
												l1090:
													position, tokenIndex = position1067, tokenIndex1067
													{
														position1100 := position
														{
															position1101, tokenIndex1101 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1102
															}
															position++
															goto l1101
														l1102:
															position, tokenIndex = position1101, tokenIndex1101
															if buffer[position] != rune('R') {
																goto l1099
															}
															position++
														}
													l1101:
														{
															position1103, tokenIndex1103 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1104
															}
															position++
															goto l1103
														l1104:
															position, tokenIndex = position1103, tokenIndex1103
															if buffer[position] != rune('R') {
																goto l1099
															}
															position++
														}
													l1103:
														if !_rules[rulews]() {
															goto l1099
														}
														if !_rules[ruleLoc8]() {
															goto l1099
														}
														{
															position1105, tokenIndex1105 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1105
															}
															if !_rules[ruleCopy8]() {
																goto l1105
															}
															goto l1106
														l1105:
															position, tokenIndex = position1105, tokenIndex1105
														}
													l1106:
														{
															add(ruleAction50, position)
														}
														add(ruleRr, position1100)
													}
													goto l1067
												l1099:
													position, tokenIndex = position1067, tokenIndex1067
													{
														position1109 := position
														{
															position1110, tokenIndex1110 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1111
															}
															position++
															goto l1110
														l1111:
															position, tokenIndex = position1110, tokenIndex1110
															if buffer[position] != rune('S') {
																goto l1108
															}
															position++
														}
													l1110:
														{
															position1112, tokenIndex1112 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1113
															}
															position++
															goto l1112
														l1113:
															position, tokenIndex = position1112, tokenIndex1112
															if buffer[position] != rune('L') {
																goto l1108
															}
															position++
														}
													l1112:
														{
															position1114, tokenIndex1114 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1115
															}
															position++
															goto l1114
														l1115:
															position, tokenIndex = position1114, tokenIndex1114
															if buffer[position] != rune('A') {
																goto l1108
															}
															position++
														}
													l1114:
														if !_rules[rulews]() {
															goto l1108
														}
														if !_rules[ruleLoc8]() {
															goto l1108
														}
														{
															position1116, tokenIndex1116 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1116
															}
															if !_rules[ruleCopy8]() {
																goto l1116
															}
															goto l1117
														l1116:
															position, tokenIndex = position1116, tokenIndex1116
														}
													l1117:
														{
															add(ruleAction51, position)
														}
														add(ruleSla, position1109)
													}
													goto l1067
												l1108:
													position, tokenIndex = position1067, tokenIndex1067
													{
														position1120 := position
														{
															position1121, tokenIndex1121 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1122
															}
															position++
															goto l1121
														l1122:
															position, tokenIndex = position1121, tokenIndex1121
															if buffer[position] != rune('S') {
																goto l1119
															}
															position++
														}
													l1121:
														{
															position1123, tokenIndex1123 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1124
															}
															position++
															goto l1123
														l1124:
															position, tokenIndex = position1123, tokenIndex1123
															if buffer[position] != rune('R') {
																goto l1119
															}
															position++
														}
													l1123:
														{
															position1125, tokenIndex1125 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1126
															}
															position++
															goto l1125
														l1126:
															position, tokenIndex = position1125, tokenIndex1125
															if buffer[position] != rune('A') {
																goto l1119
															}
															position++
														}
													l1125:
														if !_rules[rulews]() {
															goto l1119
														}
														if !_rules[ruleLoc8]() {
															goto l1119
														}
														{
															position1127, tokenIndex1127 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1127
															}
															if !_rules[ruleCopy8]() {
																goto l1127
															}
															goto l1128
														l1127:
															position, tokenIndex = position1127, tokenIndex1127
														}
													l1128:
														{
															add(ruleAction52, position)
														}
														add(ruleSra, position1120)
													}
													goto l1067
												l1119:
													position, tokenIndex = position1067, tokenIndex1067
													{
														position1131 := position
														{
															position1132, tokenIndex1132 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1133
															}
															position++
															goto l1132
														l1133:
															position, tokenIndex = position1132, tokenIndex1132
															if buffer[position] != rune('S') {
																goto l1130
															}
															position++
														}
													l1132:
														{
															position1134, tokenIndex1134 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1135
															}
															position++
															goto l1134
														l1135:
															position, tokenIndex = position1134, tokenIndex1134
															if buffer[position] != rune('L') {
																goto l1130
															}
															position++
														}
													l1134:
														{
															position1136, tokenIndex1136 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1137
															}
															position++
															goto l1136
														l1137:
															position, tokenIndex = position1136, tokenIndex1136
															if buffer[position] != rune('L') {
																goto l1130
															}
															position++
														}
													l1136:
														if !_rules[rulews]() {
															goto l1130
														}
														if !_rules[ruleLoc8]() {
															goto l1130
														}
														{
															position1138, tokenIndex1138 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1138
															}
															if !_rules[ruleCopy8]() {
																goto l1138
															}
															goto l1139
														l1138:
															position, tokenIndex = position1138, tokenIndex1138
														}
													l1139:
														{
															add(ruleAction53, position)
														}
														add(ruleSll, position1131)
													}
													goto l1067
												l1130:
													position, tokenIndex = position1067, tokenIndex1067
													{
														position1141 := position
														{
															position1142, tokenIndex1142 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1143
															}
															position++
															goto l1142
														l1143:
															position, tokenIndex = position1142, tokenIndex1142
															if buffer[position] != rune('S') {
																goto l1065
															}
															position++
														}
													l1142:
														{
															position1144, tokenIndex1144 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1145
															}
															position++
															goto l1144
														l1145:
															position, tokenIndex = position1144, tokenIndex1144
															if buffer[position] != rune('R') {
																goto l1065
															}
															position++
														}
													l1144:
														{
															position1146, tokenIndex1146 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1147
															}
															position++
															goto l1146
														l1147:
															position, tokenIndex = position1146, tokenIndex1146
															if buffer[position] != rune('L') {
																goto l1065
															}
															position++
														}
													l1146:
														if !_rules[rulews]() {
															goto l1065
														}
														if !_rules[ruleLoc8]() {
															goto l1065
														}
														{
															position1148, tokenIndex1148 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1148
															}
															if !_rules[ruleCopy8]() {
																goto l1148
															}
															goto l1149
														l1148:
															position, tokenIndex = position1148, tokenIndex1148
														}
													l1149:
														{
															add(ruleAction54, position)
														}
														add(ruleSrl, position1141)
													}
												}
											l1067:
												add(ruleRot, position1066)
											}
											goto l1064
										l1065:
											position, tokenIndex = position1064, tokenIndex1064
											{
												switch buffer[position] {
												case 'S', 's':
													{
														position1152 := position
														{
															position1153, tokenIndex1153 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1154
															}
															position++
															goto l1153
														l1154:
															position, tokenIndex = position1153, tokenIndex1153
															if buffer[position] != rune('S') {
																goto l1062
															}
															position++
														}
													l1153:
														{
															position1155, tokenIndex1155 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1156
															}
															position++
															goto l1155
														l1156:
															position, tokenIndex = position1155, tokenIndex1155
															if buffer[position] != rune('E') {
																goto l1062
															}
															position++
														}
													l1155:
														{
															position1157, tokenIndex1157 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1158
															}
															position++
															goto l1157
														l1158:
															position, tokenIndex = position1157, tokenIndex1157
															if buffer[position] != rune('T') {
																goto l1062
															}
															position++
														}
													l1157:
														if !_rules[rulews]() {
															goto l1062
														}
														if !_rules[ruleoctaldigit]() {
															goto l1062
														}
														if !_rules[rulesep]() {
															goto l1062
														}
														if !_rules[ruleLoc8]() {
															goto l1062
														}
														{
															position1159, tokenIndex1159 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1159
															}
															if !_rules[ruleCopy8]() {
																goto l1159
															}
															goto l1160
														l1159:
															position, tokenIndex = position1159, tokenIndex1159
														}
													l1160:
														{
															add(ruleAction57, position)
														}
														add(ruleSet, position1152)
													}
													break
												case 'R', 'r':
													{
														position1162 := position
														{
															position1163, tokenIndex1163 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1164
															}
															position++
															goto l1163
														l1164:
															position, tokenIndex = position1163, tokenIndex1163
															if buffer[position] != rune('R') {
																goto l1062
															}
															position++
														}
													l1163:
														{
															position1165, tokenIndex1165 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1166
															}
															position++
															goto l1165
														l1166:
															position, tokenIndex = position1165, tokenIndex1165
															if buffer[position] != rune('E') {
																goto l1062
															}
															position++
														}
													l1165:
														{
															position1167, tokenIndex1167 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1168
															}
															position++
															goto l1167
														l1168:
															position, tokenIndex = position1167, tokenIndex1167
															if buffer[position] != rune('S') {
																goto l1062
															}
															position++
														}
													l1167:
														if !_rules[rulews]() {
															goto l1062
														}
														if !_rules[ruleoctaldigit]() {
															goto l1062
														}
														if !_rules[rulesep]() {
															goto l1062
														}
														if !_rules[ruleLoc8]() {
															goto l1062
														}
														{
															position1169, tokenIndex1169 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1169
															}
															if !_rules[ruleCopy8]() {
																goto l1169
															}
															goto l1170
														l1169:
															position, tokenIndex = position1169, tokenIndex1169
														}
													l1170:
														{
															add(ruleAction56, position)
														}
														add(ruleRes, position1162)
													}
													break
												default:
													{
														position1172 := position
														{
															position1173, tokenIndex1173 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1174
															}
															position++
															goto l1173
														l1174:
															position, tokenIndex = position1173, tokenIndex1173
															if buffer[position] != rune('B') {
																goto l1062
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
																goto l1062
															}
															position++
														}
													l1175:
														{
															position1177, tokenIndex1177 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1178
															}
															position++
															goto l1177
														l1178:
															position, tokenIndex = position1177, tokenIndex1177
															if buffer[position] != rune('T') {
																goto l1062
															}
															position++
														}
													l1177:
														if !_rules[rulews]() {
															goto l1062
														}
														if !_rules[ruleoctaldigit]() {
															goto l1062
														}
														if !_rules[rulesep]() {
															goto l1062
														}
														if !_rules[ruleLoc8]() {
															goto l1062
														}
														{
															add(ruleAction55, position)
														}
														add(ruleBit, position1172)
													}
													break
												}
											}

										}
									l1064:
										add(ruleBitOp, position1063)
									}
									goto l852
								l1062:
									position, tokenIndex = position852, tokenIndex852
									{
										position1181 := position
										{
											position1182, tokenIndex1182 := position, tokenIndex
											{
												position1184 := position
												{
													position1185 := position
													{
														position1186, tokenIndex1186 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1187
														}
														position++
														goto l1186
													l1187:
														position, tokenIndex = position1186, tokenIndex1186
														if buffer[position] != rune('R') {
															goto l1183
														}
														position++
													}
												l1186:
													{
														position1188, tokenIndex1188 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1189
														}
														position++
														goto l1188
													l1189:
														position, tokenIndex = position1188, tokenIndex1188
														if buffer[position] != rune('E') {
															goto l1183
														}
														position++
													}
												l1188:
													{
														position1190, tokenIndex1190 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1191
														}
														position++
														goto l1190
													l1191:
														position, tokenIndex = position1190, tokenIndex1190
														if buffer[position] != rune('T') {
															goto l1183
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
															goto l1183
														}
														position++
													}
												l1192:
													add(rulePegText, position1185)
												}
												{
													add(ruleAction72, position)
												}
												add(ruleRetn, position1184)
											}
											goto l1182
										l1183:
											position, tokenIndex = position1182, tokenIndex1182
											{
												position1196 := position
												{
													position1197 := position
													{
														position1198, tokenIndex1198 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1199
														}
														position++
														goto l1198
													l1199:
														position, tokenIndex = position1198, tokenIndex1198
														if buffer[position] != rune('R') {
															goto l1195
														}
														position++
													}
												l1198:
													{
														position1200, tokenIndex1200 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1201
														}
														position++
														goto l1200
													l1201:
														position, tokenIndex = position1200, tokenIndex1200
														if buffer[position] != rune('E') {
															goto l1195
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
															goto l1195
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
															goto l1195
														}
														position++
													}
												l1204:
													add(rulePegText, position1197)
												}
												{
													add(ruleAction73, position)
												}
												add(ruleReti, position1196)
											}
											goto l1182
										l1195:
											position, tokenIndex = position1182, tokenIndex1182
											{
												position1208 := position
												{
													position1209 := position
													{
														position1210, tokenIndex1210 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1211
														}
														position++
														goto l1210
													l1211:
														position, tokenIndex = position1210, tokenIndex1210
														if buffer[position] != rune('R') {
															goto l1207
														}
														position++
													}
												l1210:
													{
														position1212, tokenIndex1212 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1213
														}
														position++
														goto l1212
													l1213:
														position, tokenIndex = position1212, tokenIndex1212
														if buffer[position] != rune('R') {
															goto l1207
														}
														position++
													}
												l1212:
													{
														position1214, tokenIndex1214 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1215
														}
														position++
														goto l1214
													l1215:
														position, tokenIndex = position1214, tokenIndex1214
														if buffer[position] != rune('D') {
															goto l1207
														}
														position++
													}
												l1214:
													add(rulePegText, position1209)
												}
												{
													add(ruleAction74, position)
												}
												add(ruleRrd, position1208)
											}
											goto l1182
										l1207:
											position, tokenIndex = position1182, tokenIndex1182
											{
												position1218 := position
												{
													position1219 := position
													{
														position1220, tokenIndex1220 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1221
														}
														position++
														goto l1220
													l1221:
														position, tokenIndex = position1220, tokenIndex1220
														if buffer[position] != rune('I') {
															goto l1217
														}
														position++
													}
												l1220:
													{
														position1222, tokenIndex1222 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1223
														}
														position++
														goto l1222
													l1223:
														position, tokenIndex = position1222, tokenIndex1222
														if buffer[position] != rune('M') {
															goto l1217
														}
														position++
													}
												l1222:
													if buffer[position] != rune(' ') {
														goto l1217
													}
													position++
													if buffer[position] != rune('0') {
														goto l1217
													}
													position++
													add(rulePegText, position1219)
												}
												{
													add(ruleAction76, position)
												}
												add(ruleIm0, position1218)
											}
											goto l1182
										l1217:
											position, tokenIndex = position1182, tokenIndex1182
											{
												position1226 := position
												{
													position1227 := position
													{
														position1228, tokenIndex1228 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1229
														}
														position++
														goto l1228
													l1229:
														position, tokenIndex = position1228, tokenIndex1228
														if buffer[position] != rune('I') {
															goto l1225
														}
														position++
													}
												l1228:
													{
														position1230, tokenIndex1230 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1231
														}
														position++
														goto l1230
													l1231:
														position, tokenIndex = position1230, tokenIndex1230
														if buffer[position] != rune('M') {
															goto l1225
														}
														position++
													}
												l1230:
													if buffer[position] != rune(' ') {
														goto l1225
													}
													position++
													if buffer[position] != rune('1') {
														goto l1225
													}
													position++
													add(rulePegText, position1227)
												}
												{
													add(ruleAction77, position)
												}
												add(ruleIm1, position1226)
											}
											goto l1182
										l1225:
											position, tokenIndex = position1182, tokenIndex1182
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
														if buffer[position] != rune('m') {
															goto l1239
														}
														position++
														goto l1238
													l1239:
														position, tokenIndex = position1238, tokenIndex1238
														if buffer[position] != rune('M') {
															goto l1233
														}
														position++
													}
												l1238:
													if buffer[position] != rune(' ') {
														goto l1233
													}
													position++
													if buffer[position] != rune('2') {
														goto l1233
													}
													position++
													add(rulePegText, position1235)
												}
												{
													add(ruleAction78, position)
												}
												add(ruleIm2, position1234)
											}
											goto l1182
										l1233:
											position, tokenIndex = position1182, tokenIndex1182
											{
												switch buffer[position] {
												case 'I', 'O', 'i', 'o':
													{
														position1242 := position
														{
															position1243, tokenIndex1243 := position, tokenIndex
															{
																position1245 := position
																{
																	position1246 := position
																	{
																		position1247, tokenIndex1247 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1248
																		}
																		position++
																		goto l1247
																	l1248:
																		position, tokenIndex = position1247, tokenIndex1247
																		if buffer[position] != rune('I') {
																			goto l1244
																		}
																		position++
																	}
																l1247:
																	{
																		position1249, tokenIndex1249 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1250
																		}
																		position++
																		goto l1249
																	l1250:
																		position, tokenIndex = position1249, tokenIndex1249
																		if buffer[position] != rune('N') {
																			goto l1244
																		}
																		position++
																	}
																l1249:
																	{
																		position1251, tokenIndex1251 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1252
																		}
																		position++
																		goto l1251
																	l1252:
																		position, tokenIndex = position1251, tokenIndex1251
																		if buffer[position] != rune('I') {
																			goto l1244
																		}
																		position++
																	}
																l1251:
																	{
																		position1253, tokenIndex1253 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1254
																		}
																		position++
																		goto l1253
																	l1254:
																		position, tokenIndex = position1253, tokenIndex1253
																		if buffer[position] != rune('R') {
																			goto l1244
																		}
																		position++
																	}
																l1253:
																	add(rulePegText, position1246)
																}
																{
																	add(ruleAction89, position)
																}
																add(ruleInir, position1245)
															}
															goto l1243
														l1244:
															position, tokenIndex = position1243, tokenIndex1243
															{
																position1257 := position
																{
																	position1258 := position
																	{
																		position1259, tokenIndex1259 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1260
																		}
																		position++
																		goto l1259
																	l1260:
																		position, tokenIndex = position1259, tokenIndex1259
																		if buffer[position] != rune('I') {
																			goto l1256
																		}
																		position++
																	}
																l1259:
																	{
																		position1261, tokenIndex1261 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1262
																		}
																		position++
																		goto l1261
																	l1262:
																		position, tokenIndex = position1261, tokenIndex1261
																		if buffer[position] != rune('N') {
																			goto l1256
																		}
																		position++
																	}
																l1261:
																	{
																		position1263, tokenIndex1263 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1264
																		}
																		position++
																		goto l1263
																	l1264:
																		position, tokenIndex = position1263, tokenIndex1263
																		if buffer[position] != rune('I') {
																			goto l1256
																		}
																		position++
																	}
																l1263:
																	add(rulePegText, position1258)
																}
																{
																	add(ruleAction81, position)
																}
																add(ruleIni, position1257)
															}
															goto l1243
														l1256:
															position, tokenIndex = position1243, tokenIndex1243
															{
																position1267 := position
																{
																	position1268 := position
																	{
																		position1269, tokenIndex1269 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1270
																		}
																		position++
																		goto l1269
																	l1270:
																		position, tokenIndex = position1269, tokenIndex1269
																		if buffer[position] != rune('O') {
																			goto l1266
																		}
																		position++
																	}
																l1269:
																	{
																		position1271, tokenIndex1271 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1272
																		}
																		position++
																		goto l1271
																	l1272:
																		position, tokenIndex = position1271, tokenIndex1271
																		if buffer[position] != rune('T') {
																			goto l1266
																		}
																		position++
																	}
																l1271:
																	{
																		position1273, tokenIndex1273 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1274
																		}
																		position++
																		goto l1273
																	l1274:
																		position, tokenIndex = position1273, tokenIndex1273
																		if buffer[position] != rune('I') {
																			goto l1266
																		}
																		position++
																	}
																l1273:
																	{
																		position1275, tokenIndex1275 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1276
																		}
																		position++
																		goto l1275
																	l1276:
																		position, tokenIndex = position1275, tokenIndex1275
																		if buffer[position] != rune('R') {
																			goto l1266
																		}
																		position++
																	}
																l1275:
																	add(rulePegText, position1268)
																}
																{
																	add(ruleAction90, position)
																}
																add(ruleOtir, position1267)
															}
															goto l1243
														l1266:
															position, tokenIndex = position1243, tokenIndex1243
															{
																position1279 := position
																{
																	position1280 := position
																	{
																		position1281, tokenIndex1281 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1282
																		}
																		position++
																		goto l1281
																	l1282:
																		position, tokenIndex = position1281, tokenIndex1281
																		if buffer[position] != rune('O') {
																			goto l1278
																		}
																		position++
																	}
																l1281:
																	{
																		position1283, tokenIndex1283 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1284
																		}
																		position++
																		goto l1283
																	l1284:
																		position, tokenIndex = position1283, tokenIndex1283
																		if buffer[position] != rune('U') {
																			goto l1278
																		}
																		position++
																	}
																l1283:
																	{
																		position1285, tokenIndex1285 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1286
																		}
																		position++
																		goto l1285
																	l1286:
																		position, tokenIndex = position1285, tokenIndex1285
																		if buffer[position] != rune('T') {
																			goto l1278
																		}
																		position++
																	}
																l1285:
																	{
																		position1287, tokenIndex1287 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1288
																		}
																		position++
																		goto l1287
																	l1288:
																		position, tokenIndex = position1287, tokenIndex1287
																		if buffer[position] != rune('I') {
																			goto l1278
																		}
																		position++
																	}
																l1287:
																	add(rulePegText, position1280)
																}
																{
																	add(ruleAction82, position)
																}
																add(ruleOuti, position1279)
															}
															goto l1243
														l1278:
															position, tokenIndex = position1243, tokenIndex1243
															{
																position1291 := position
																{
																	position1292 := position
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
																			goto l1290
																		}
																		position++
																	}
																l1293:
																	{
																		position1295, tokenIndex1295 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1296
																		}
																		position++
																		goto l1295
																	l1296:
																		position, tokenIndex = position1295, tokenIndex1295
																		if buffer[position] != rune('N') {
																			goto l1290
																		}
																		position++
																	}
																l1295:
																	{
																		position1297, tokenIndex1297 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1298
																		}
																		position++
																		goto l1297
																	l1298:
																		position, tokenIndex = position1297, tokenIndex1297
																		if buffer[position] != rune('D') {
																			goto l1290
																		}
																		position++
																	}
																l1297:
																	{
																		position1299, tokenIndex1299 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1300
																		}
																		position++
																		goto l1299
																	l1300:
																		position, tokenIndex = position1299, tokenIndex1299
																		if buffer[position] != rune('R') {
																			goto l1290
																		}
																		position++
																	}
																l1299:
																	add(rulePegText, position1292)
																}
																{
																	add(ruleAction93, position)
																}
																add(ruleIndr, position1291)
															}
															goto l1243
														l1290:
															position, tokenIndex = position1243, tokenIndex1243
															{
																position1303 := position
																{
																	position1304 := position
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
																			goto l1302
																		}
																		position++
																	}
																l1305:
																	{
																		position1307, tokenIndex1307 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1308
																		}
																		position++
																		goto l1307
																	l1308:
																		position, tokenIndex = position1307, tokenIndex1307
																		if buffer[position] != rune('N') {
																			goto l1302
																		}
																		position++
																	}
																l1307:
																	{
																		position1309, tokenIndex1309 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1310
																		}
																		position++
																		goto l1309
																	l1310:
																		position, tokenIndex = position1309, tokenIndex1309
																		if buffer[position] != rune('D') {
																			goto l1302
																		}
																		position++
																	}
																l1309:
																	add(rulePegText, position1304)
																}
																{
																	add(ruleAction85, position)
																}
																add(ruleInd, position1303)
															}
															goto l1243
														l1302:
															position, tokenIndex = position1243, tokenIndex1243
															{
																position1313 := position
																{
																	position1314 := position
																	{
																		position1315, tokenIndex1315 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1316
																		}
																		position++
																		goto l1315
																	l1316:
																		position, tokenIndex = position1315, tokenIndex1315
																		if buffer[position] != rune('O') {
																			goto l1312
																		}
																		position++
																	}
																l1315:
																	{
																		position1317, tokenIndex1317 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1318
																		}
																		position++
																		goto l1317
																	l1318:
																		position, tokenIndex = position1317, tokenIndex1317
																		if buffer[position] != rune('T') {
																			goto l1312
																		}
																		position++
																	}
																l1317:
																	{
																		position1319, tokenIndex1319 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1320
																		}
																		position++
																		goto l1319
																	l1320:
																		position, tokenIndex = position1319, tokenIndex1319
																		if buffer[position] != rune('D') {
																			goto l1312
																		}
																		position++
																	}
																l1319:
																	{
																		position1321, tokenIndex1321 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1322
																		}
																		position++
																		goto l1321
																	l1322:
																		position, tokenIndex = position1321, tokenIndex1321
																		if buffer[position] != rune('R') {
																			goto l1312
																		}
																		position++
																	}
																l1321:
																	add(rulePegText, position1314)
																}
																{
																	add(ruleAction94, position)
																}
																add(ruleOtdr, position1313)
															}
															goto l1243
														l1312:
															position, tokenIndex = position1243, tokenIndex1243
															{
																position1324 := position
																{
																	position1325 := position
																	{
																		position1326, tokenIndex1326 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1327
																		}
																		position++
																		goto l1326
																	l1327:
																		position, tokenIndex = position1326, tokenIndex1326
																		if buffer[position] != rune('O') {
																			goto l1180
																		}
																		position++
																	}
																l1326:
																	{
																		position1328, tokenIndex1328 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1329
																		}
																		position++
																		goto l1328
																	l1329:
																		position, tokenIndex = position1328, tokenIndex1328
																		if buffer[position] != rune('U') {
																			goto l1180
																		}
																		position++
																	}
																l1328:
																	{
																		position1330, tokenIndex1330 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1331
																		}
																		position++
																		goto l1330
																	l1331:
																		position, tokenIndex = position1330, tokenIndex1330
																		if buffer[position] != rune('T') {
																			goto l1180
																		}
																		position++
																	}
																l1330:
																	{
																		position1332, tokenIndex1332 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1333
																		}
																		position++
																		goto l1332
																	l1333:
																		position, tokenIndex = position1332, tokenIndex1332
																		if buffer[position] != rune('D') {
																			goto l1180
																		}
																		position++
																	}
																l1332:
																	add(rulePegText, position1325)
																}
																{
																	add(ruleAction86, position)
																}
																add(ruleOutd, position1324)
															}
														}
													l1243:
														add(ruleBlitIO, position1242)
													}
													break
												case 'R', 'r':
													{
														position1335 := position
														{
															position1336 := position
															{
																position1337, tokenIndex1337 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1338
																}
																position++
																goto l1337
															l1338:
																position, tokenIndex = position1337, tokenIndex1337
																if buffer[position] != rune('R') {
																	goto l1180
																}
																position++
															}
														l1337:
															{
																position1339, tokenIndex1339 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1340
																}
																position++
																goto l1339
															l1340:
																position, tokenIndex = position1339, tokenIndex1339
																if buffer[position] != rune('L') {
																	goto l1180
																}
																position++
															}
														l1339:
															{
																position1341, tokenIndex1341 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1342
																}
																position++
																goto l1341
															l1342:
																position, tokenIndex = position1341, tokenIndex1341
																if buffer[position] != rune('D') {
																	goto l1180
																}
																position++
															}
														l1341:
															add(rulePegText, position1336)
														}
														{
															add(ruleAction75, position)
														}
														add(ruleRld, position1335)
													}
													break
												case 'N', 'n':
													{
														position1344 := position
														{
															position1345 := position
															{
																position1346, tokenIndex1346 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1347
																}
																position++
																goto l1346
															l1347:
																position, tokenIndex = position1346, tokenIndex1346
																if buffer[position] != rune('N') {
																	goto l1180
																}
																position++
															}
														l1346:
															{
																position1348, tokenIndex1348 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1349
																}
																position++
																goto l1348
															l1349:
																position, tokenIndex = position1348, tokenIndex1348
																if buffer[position] != rune('E') {
																	goto l1180
																}
																position++
															}
														l1348:
															{
																position1350, tokenIndex1350 := position, tokenIndex
																if buffer[position] != rune('g') {
																	goto l1351
																}
																position++
																goto l1350
															l1351:
																position, tokenIndex = position1350, tokenIndex1350
																if buffer[position] != rune('G') {
																	goto l1180
																}
																position++
															}
														l1350:
															add(rulePegText, position1345)
														}
														{
															add(ruleAction71, position)
														}
														add(ruleNeg, position1344)
													}
													break
												default:
													{
														position1353 := position
														{
															position1354, tokenIndex1354 := position, tokenIndex
															{
																position1356 := position
																{
																	position1357 := position
																	{
																		position1358, tokenIndex1358 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1359
																		}
																		position++
																		goto l1358
																	l1359:
																		position, tokenIndex = position1358, tokenIndex1358
																		if buffer[position] != rune('L') {
																			goto l1355
																		}
																		position++
																	}
																l1358:
																	{
																		position1360, tokenIndex1360 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1361
																		}
																		position++
																		goto l1360
																	l1361:
																		position, tokenIndex = position1360, tokenIndex1360
																		if buffer[position] != rune('D') {
																			goto l1355
																		}
																		position++
																	}
																l1360:
																	{
																		position1362, tokenIndex1362 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1363
																		}
																		position++
																		goto l1362
																	l1363:
																		position, tokenIndex = position1362, tokenIndex1362
																		if buffer[position] != rune('I') {
																			goto l1355
																		}
																		position++
																	}
																l1362:
																	{
																		position1364, tokenIndex1364 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1365
																		}
																		position++
																		goto l1364
																	l1365:
																		position, tokenIndex = position1364, tokenIndex1364
																		if buffer[position] != rune('R') {
																			goto l1355
																		}
																		position++
																	}
																l1364:
																	add(rulePegText, position1357)
																}
																{
																	add(ruleAction87, position)
																}
																add(ruleLdir, position1356)
															}
															goto l1354
														l1355:
															position, tokenIndex = position1354, tokenIndex1354
															{
																position1368 := position
																{
																	position1369 := position
																	{
																		position1370, tokenIndex1370 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1371
																		}
																		position++
																		goto l1370
																	l1371:
																		position, tokenIndex = position1370, tokenIndex1370
																		if buffer[position] != rune('L') {
																			goto l1367
																		}
																		position++
																	}
																l1370:
																	{
																		position1372, tokenIndex1372 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1373
																		}
																		position++
																		goto l1372
																	l1373:
																		position, tokenIndex = position1372, tokenIndex1372
																		if buffer[position] != rune('D') {
																			goto l1367
																		}
																		position++
																	}
																l1372:
																	{
																		position1374, tokenIndex1374 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1375
																		}
																		position++
																		goto l1374
																	l1375:
																		position, tokenIndex = position1374, tokenIndex1374
																		if buffer[position] != rune('I') {
																			goto l1367
																		}
																		position++
																	}
																l1374:
																	add(rulePegText, position1369)
																}
																{
																	add(ruleAction79, position)
																}
																add(ruleLdi, position1368)
															}
															goto l1354
														l1367:
															position, tokenIndex = position1354, tokenIndex1354
															{
																position1378 := position
																{
																	position1379 := position
																	{
																		position1380, tokenIndex1380 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1381
																		}
																		position++
																		goto l1380
																	l1381:
																		position, tokenIndex = position1380, tokenIndex1380
																		if buffer[position] != rune('C') {
																			goto l1377
																		}
																		position++
																	}
																l1380:
																	{
																		position1382, tokenIndex1382 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1383
																		}
																		position++
																		goto l1382
																	l1383:
																		position, tokenIndex = position1382, tokenIndex1382
																		if buffer[position] != rune('P') {
																			goto l1377
																		}
																		position++
																	}
																l1382:
																	{
																		position1384, tokenIndex1384 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1385
																		}
																		position++
																		goto l1384
																	l1385:
																		position, tokenIndex = position1384, tokenIndex1384
																		if buffer[position] != rune('I') {
																			goto l1377
																		}
																		position++
																	}
																l1384:
																	{
																		position1386, tokenIndex1386 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1387
																		}
																		position++
																		goto l1386
																	l1387:
																		position, tokenIndex = position1386, tokenIndex1386
																		if buffer[position] != rune('R') {
																			goto l1377
																		}
																		position++
																	}
																l1386:
																	add(rulePegText, position1379)
																}
																{
																	add(ruleAction88, position)
																}
																add(ruleCpir, position1378)
															}
															goto l1354
														l1377:
															position, tokenIndex = position1354, tokenIndex1354
															{
																position1390 := position
																{
																	position1391 := position
																	{
																		position1392, tokenIndex1392 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1393
																		}
																		position++
																		goto l1392
																	l1393:
																		position, tokenIndex = position1392, tokenIndex1392
																		if buffer[position] != rune('C') {
																			goto l1389
																		}
																		position++
																	}
																l1392:
																	{
																		position1394, tokenIndex1394 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1395
																		}
																		position++
																		goto l1394
																	l1395:
																		position, tokenIndex = position1394, tokenIndex1394
																		if buffer[position] != rune('P') {
																			goto l1389
																		}
																		position++
																	}
																l1394:
																	{
																		position1396, tokenIndex1396 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1397
																		}
																		position++
																		goto l1396
																	l1397:
																		position, tokenIndex = position1396, tokenIndex1396
																		if buffer[position] != rune('I') {
																			goto l1389
																		}
																		position++
																	}
																l1396:
																	add(rulePegText, position1391)
																}
																{
																	add(ruleAction80, position)
																}
																add(ruleCpi, position1390)
															}
															goto l1354
														l1389:
															position, tokenIndex = position1354, tokenIndex1354
															{
																position1400 := position
																{
																	position1401 := position
																	{
																		position1402, tokenIndex1402 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1403
																		}
																		position++
																		goto l1402
																	l1403:
																		position, tokenIndex = position1402, tokenIndex1402
																		if buffer[position] != rune('L') {
																			goto l1399
																		}
																		position++
																	}
																l1402:
																	{
																		position1404, tokenIndex1404 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1405
																		}
																		position++
																		goto l1404
																	l1405:
																		position, tokenIndex = position1404, tokenIndex1404
																		if buffer[position] != rune('D') {
																			goto l1399
																		}
																		position++
																	}
																l1404:
																	{
																		position1406, tokenIndex1406 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1407
																		}
																		position++
																		goto l1406
																	l1407:
																		position, tokenIndex = position1406, tokenIndex1406
																		if buffer[position] != rune('D') {
																			goto l1399
																		}
																		position++
																	}
																l1406:
																	{
																		position1408, tokenIndex1408 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1409
																		}
																		position++
																		goto l1408
																	l1409:
																		position, tokenIndex = position1408, tokenIndex1408
																		if buffer[position] != rune('R') {
																			goto l1399
																		}
																		position++
																	}
																l1408:
																	add(rulePegText, position1401)
																}
																{
																	add(ruleAction91, position)
																}
																add(ruleLddr, position1400)
															}
															goto l1354
														l1399:
															position, tokenIndex = position1354, tokenIndex1354
															{
																position1412 := position
																{
																	position1413 := position
																	{
																		position1414, tokenIndex1414 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1415
																		}
																		position++
																		goto l1414
																	l1415:
																		position, tokenIndex = position1414, tokenIndex1414
																		if buffer[position] != rune('L') {
																			goto l1411
																		}
																		position++
																	}
																l1414:
																	{
																		position1416, tokenIndex1416 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1417
																		}
																		position++
																		goto l1416
																	l1417:
																		position, tokenIndex = position1416, tokenIndex1416
																		if buffer[position] != rune('D') {
																			goto l1411
																		}
																		position++
																	}
																l1416:
																	{
																		position1418, tokenIndex1418 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1419
																		}
																		position++
																		goto l1418
																	l1419:
																		position, tokenIndex = position1418, tokenIndex1418
																		if buffer[position] != rune('D') {
																			goto l1411
																		}
																		position++
																	}
																l1418:
																	add(rulePegText, position1413)
																}
																{
																	add(ruleAction83, position)
																}
																add(ruleLdd, position1412)
															}
															goto l1354
														l1411:
															position, tokenIndex = position1354, tokenIndex1354
															{
																position1422 := position
																{
																	position1423 := position
																	{
																		position1424, tokenIndex1424 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1425
																		}
																		position++
																		goto l1424
																	l1425:
																		position, tokenIndex = position1424, tokenIndex1424
																		if buffer[position] != rune('C') {
																			goto l1421
																		}
																		position++
																	}
																l1424:
																	{
																		position1426, tokenIndex1426 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1427
																		}
																		position++
																		goto l1426
																	l1427:
																		position, tokenIndex = position1426, tokenIndex1426
																		if buffer[position] != rune('P') {
																			goto l1421
																		}
																		position++
																	}
																l1426:
																	{
																		position1428, tokenIndex1428 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1429
																		}
																		position++
																		goto l1428
																	l1429:
																		position, tokenIndex = position1428, tokenIndex1428
																		if buffer[position] != rune('D') {
																			goto l1421
																		}
																		position++
																	}
																l1428:
																	{
																		position1430, tokenIndex1430 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1431
																		}
																		position++
																		goto l1430
																	l1431:
																		position, tokenIndex = position1430, tokenIndex1430
																		if buffer[position] != rune('R') {
																			goto l1421
																		}
																		position++
																	}
																l1430:
																	add(rulePegText, position1423)
																}
																{
																	add(ruleAction92, position)
																}
																add(ruleCpdr, position1422)
															}
															goto l1354
														l1421:
															position, tokenIndex = position1354, tokenIndex1354
															{
																position1433 := position
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
																			goto l1180
																		}
																		position++
																	}
																l1435:
																	{
																		position1437, tokenIndex1437 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1438
																		}
																		position++
																		goto l1437
																	l1438:
																		position, tokenIndex = position1437, tokenIndex1437
																		if buffer[position] != rune('P') {
																			goto l1180
																		}
																		position++
																	}
																l1437:
																	{
																		position1439, tokenIndex1439 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1440
																		}
																		position++
																		goto l1439
																	l1440:
																		position, tokenIndex = position1439, tokenIndex1439
																		if buffer[position] != rune('D') {
																			goto l1180
																		}
																		position++
																	}
																l1439:
																	add(rulePegText, position1434)
																}
																{
																	add(ruleAction84, position)
																}
																add(ruleCpd, position1433)
															}
														}
													l1354:
														add(ruleBlit, position1353)
													}
													break
												}
											}

										}
									l1182:
										add(ruleEDSimple, position1181)
									}
									goto l852
								l1180:
									position, tokenIndex = position852, tokenIndex852
									{
										position1443 := position
										{
											position1444, tokenIndex1444 := position, tokenIndex
											{
												position1446 := position
												{
													position1447 := position
													{
														position1448, tokenIndex1448 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1449
														}
														position++
														goto l1448
													l1449:
														position, tokenIndex = position1448, tokenIndex1448
														if buffer[position] != rune('R') {
															goto l1445
														}
														position++
													}
												l1448:
													{
														position1450, tokenIndex1450 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1451
														}
														position++
														goto l1450
													l1451:
														position, tokenIndex = position1450, tokenIndex1450
														if buffer[position] != rune('L') {
															goto l1445
														}
														position++
													}
												l1450:
													{
														position1452, tokenIndex1452 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1453
														}
														position++
														goto l1452
													l1453:
														position, tokenIndex = position1452, tokenIndex1452
														if buffer[position] != rune('C') {
															goto l1445
														}
														position++
													}
												l1452:
													{
														position1454, tokenIndex1454 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1455
														}
														position++
														goto l1454
													l1455:
														position, tokenIndex = position1454, tokenIndex1454
														if buffer[position] != rune('A') {
															goto l1445
														}
														position++
													}
												l1454:
													add(rulePegText, position1447)
												}
												{
													add(ruleAction60, position)
												}
												add(ruleRlca, position1446)
											}
											goto l1444
										l1445:
											position, tokenIndex = position1444, tokenIndex1444
											{
												position1458 := position
												{
													position1459 := position
													{
														position1460, tokenIndex1460 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1461
														}
														position++
														goto l1460
													l1461:
														position, tokenIndex = position1460, tokenIndex1460
														if buffer[position] != rune('R') {
															goto l1457
														}
														position++
													}
												l1460:
													{
														position1462, tokenIndex1462 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1463
														}
														position++
														goto l1462
													l1463:
														position, tokenIndex = position1462, tokenIndex1462
														if buffer[position] != rune('R') {
															goto l1457
														}
														position++
													}
												l1462:
													{
														position1464, tokenIndex1464 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1465
														}
														position++
														goto l1464
													l1465:
														position, tokenIndex = position1464, tokenIndex1464
														if buffer[position] != rune('C') {
															goto l1457
														}
														position++
													}
												l1464:
													{
														position1466, tokenIndex1466 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1467
														}
														position++
														goto l1466
													l1467:
														position, tokenIndex = position1466, tokenIndex1466
														if buffer[position] != rune('A') {
															goto l1457
														}
														position++
													}
												l1466:
													add(rulePegText, position1459)
												}
												{
													add(ruleAction61, position)
												}
												add(ruleRrca, position1458)
											}
											goto l1444
										l1457:
											position, tokenIndex = position1444, tokenIndex1444
											{
												position1470 := position
												{
													position1471 := position
													{
														position1472, tokenIndex1472 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1473
														}
														position++
														goto l1472
													l1473:
														position, tokenIndex = position1472, tokenIndex1472
														if buffer[position] != rune('R') {
															goto l1469
														}
														position++
													}
												l1472:
													{
														position1474, tokenIndex1474 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1475
														}
														position++
														goto l1474
													l1475:
														position, tokenIndex = position1474, tokenIndex1474
														if buffer[position] != rune('L') {
															goto l1469
														}
														position++
													}
												l1474:
													{
														position1476, tokenIndex1476 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1477
														}
														position++
														goto l1476
													l1477:
														position, tokenIndex = position1476, tokenIndex1476
														if buffer[position] != rune('A') {
															goto l1469
														}
														position++
													}
												l1476:
													add(rulePegText, position1471)
												}
												{
													add(ruleAction62, position)
												}
												add(ruleRla, position1470)
											}
											goto l1444
										l1469:
											position, tokenIndex = position1444, tokenIndex1444
											{
												position1480 := position
												{
													position1481 := position
													{
														position1482, tokenIndex1482 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1483
														}
														position++
														goto l1482
													l1483:
														position, tokenIndex = position1482, tokenIndex1482
														if buffer[position] != rune('D') {
															goto l1479
														}
														position++
													}
												l1482:
													{
														position1484, tokenIndex1484 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1485
														}
														position++
														goto l1484
													l1485:
														position, tokenIndex = position1484, tokenIndex1484
														if buffer[position] != rune('A') {
															goto l1479
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
															goto l1479
														}
														position++
													}
												l1486:
													add(rulePegText, position1481)
												}
												{
													add(ruleAction64, position)
												}
												add(ruleDaa, position1480)
											}
											goto l1444
										l1479:
											position, tokenIndex = position1444, tokenIndex1444
											{
												position1490 := position
												{
													position1491 := position
													{
														position1492, tokenIndex1492 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1493
														}
														position++
														goto l1492
													l1493:
														position, tokenIndex = position1492, tokenIndex1492
														if buffer[position] != rune('C') {
															goto l1489
														}
														position++
													}
												l1492:
													{
														position1494, tokenIndex1494 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1495
														}
														position++
														goto l1494
													l1495:
														position, tokenIndex = position1494, tokenIndex1494
														if buffer[position] != rune('P') {
															goto l1489
														}
														position++
													}
												l1494:
													{
														position1496, tokenIndex1496 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1497
														}
														position++
														goto l1496
													l1497:
														position, tokenIndex = position1496, tokenIndex1496
														if buffer[position] != rune('L') {
															goto l1489
														}
														position++
													}
												l1496:
													add(rulePegText, position1491)
												}
												{
													add(ruleAction65, position)
												}
												add(ruleCpl, position1490)
											}
											goto l1444
										l1489:
											position, tokenIndex = position1444, tokenIndex1444
											{
												position1500 := position
												{
													position1501 := position
													{
														position1502, tokenIndex1502 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1503
														}
														position++
														goto l1502
													l1503:
														position, tokenIndex = position1502, tokenIndex1502
														if buffer[position] != rune('E') {
															goto l1499
														}
														position++
													}
												l1502:
													{
														position1504, tokenIndex1504 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1505
														}
														position++
														goto l1504
													l1505:
														position, tokenIndex = position1504, tokenIndex1504
														if buffer[position] != rune('X') {
															goto l1499
														}
														position++
													}
												l1504:
													{
														position1506, tokenIndex1506 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1507
														}
														position++
														goto l1506
													l1507:
														position, tokenIndex = position1506, tokenIndex1506
														if buffer[position] != rune('X') {
															goto l1499
														}
														position++
													}
												l1506:
													add(rulePegText, position1501)
												}
												{
													add(ruleAction68, position)
												}
												add(ruleExx, position1500)
											}
											goto l1444
										l1499:
											position, tokenIndex = position1444, tokenIndex1444
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position1510 := position
														{
															position1511 := position
															{
																position1512, tokenIndex1512 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1513
																}
																position++
																goto l1512
															l1513:
																position, tokenIndex = position1512, tokenIndex1512
																if buffer[position] != rune('E') {
																	goto l1442
																}
																position++
															}
														l1512:
															{
																position1514, tokenIndex1514 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1515
																}
																position++
																goto l1514
															l1515:
																position, tokenIndex = position1514, tokenIndex1514
																if buffer[position] != rune('I') {
																	goto l1442
																}
																position++
															}
														l1514:
															add(rulePegText, position1511)
														}
														{
															add(ruleAction70, position)
														}
														add(ruleEi, position1510)
													}
													break
												case 'D', 'd':
													{
														position1517 := position
														{
															position1518 := position
															{
																position1519, tokenIndex1519 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1520
																}
																position++
																goto l1519
															l1520:
																position, tokenIndex = position1519, tokenIndex1519
																if buffer[position] != rune('D') {
																	goto l1442
																}
																position++
															}
														l1519:
															{
																position1521, tokenIndex1521 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1522
																}
																position++
																goto l1521
															l1522:
																position, tokenIndex = position1521, tokenIndex1521
																if buffer[position] != rune('I') {
																	goto l1442
																}
																position++
															}
														l1521:
															add(rulePegText, position1518)
														}
														{
															add(ruleAction69, position)
														}
														add(ruleDi, position1517)
													}
													break
												case 'C', 'c':
													{
														position1524 := position
														{
															position1525 := position
															{
																position1526, tokenIndex1526 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1527
																}
																position++
																goto l1526
															l1527:
																position, tokenIndex = position1526, tokenIndex1526
																if buffer[position] != rune('C') {
																	goto l1442
																}
																position++
															}
														l1526:
															{
																position1528, tokenIndex1528 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1529
																}
																position++
																goto l1528
															l1529:
																position, tokenIndex = position1528, tokenIndex1528
																if buffer[position] != rune('C') {
																	goto l1442
																}
																position++
															}
														l1528:
															{
																position1530, tokenIndex1530 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1531
																}
																position++
																goto l1530
															l1531:
																position, tokenIndex = position1530, tokenIndex1530
																if buffer[position] != rune('F') {
																	goto l1442
																}
																position++
															}
														l1530:
															add(rulePegText, position1525)
														}
														{
															add(ruleAction67, position)
														}
														add(ruleCcf, position1524)
													}
													break
												case 'S', 's':
													{
														position1533 := position
														{
															position1534 := position
															{
																position1535, tokenIndex1535 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1536
																}
																position++
																goto l1535
															l1536:
																position, tokenIndex = position1535, tokenIndex1535
																if buffer[position] != rune('S') {
																	goto l1442
																}
																position++
															}
														l1535:
															{
																position1537, tokenIndex1537 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1538
																}
																position++
																goto l1537
															l1538:
																position, tokenIndex = position1537, tokenIndex1537
																if buffer[position] != rune('C') {
																	goto l1442
																}
																position++
															}
														l1537:
															{
																position1539, tokenIndex1539 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1540
																}
																position++
																goto l1539
															l1540:
																position, tokenIndex = position1539, tokenIndex1539
																if buffer[position] != rune('F') {
																	goto l1442
																}
																position++
															}
														l1539:
															add(rulePegText, position1534)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleScf, position1533)
													}
													break
												case 'R', 'r':
													{
														position1542 := position
														{
															position1543 := position
															{
																position1544, tokenIndex1544 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1545
																}
																position++
																goto l1544
															l1545:
																position, tokenIndex = position1544, tokenIndex1544
																if buffer[position] != rune('R') {
																	goto l1442
																}
																position++
															}
														l1544:
															{
																position1546, tokenIndex1546 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1547
																}
																position++
																goto l1546
															l1547:
																position, tokenIndex = position1546, tokenIndex1546
																if buffer[position] != rune('R') {
																	goto l1442
																}
																position++
															}
														l1546:
															{
																position1548, tokenIndex1548 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1549
																}
																position++
																goto l1548
															l1549:
																position, tokenIndex = position1548, tokenIndex1548
																if buffer[position] != rune('A') {
																	goto l1442
																}
																position++
															}
														l1548:
															add(rulePegText, position1543)
														}
														{
															add(ruleAction63, position)
														}
														add(ruleRra, position1542)
													}
													break
												case 'H', 'h':
													{
														position1551 := position
														{
															position1552 := position
															{
																position1553, tokenIndex1553 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l1554
																}
																position++
																goto l1553
															l1554:
																position, tokenIndex = position1553, tokenIndex1553
																if buffer[position] != rune('H') {
																	goto l1442
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
																	goto l1442
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
																	goto l1442
																}
																position++
															}
														l1557:
															{
																position1559, tokenIndex1559 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1560
																}
																position++
																goto l1559
															l1560:
																position, tokenIndex = position1559, tokenIndex1559
																if buffer[position] != rune('T') {
																	goto l1442
																}
																position++
															}
														l1559:
															add(rulePegText, position1552)
														}
														{
															add(ruleAction59, position)
														}
														add(ruleHalt, position1551)
													}
													break
												default:
													{
														position1562 := position
														{
															position1563 := position
															{
																position1564, tokenIndex1564 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1565
																}
																position++
																goto l1564
															l1565:
																position, tokenIndex = position1564, tokenIndex1564
																if buffer[position] != rune('N') {
																	goto l1442
																}
																position++
															}
														l1564:
															{
																position1566, tokenIndex1566 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l1567
																}
																position++
																goto l1566
															l1567:
																position, tokenIndex = position1566, tokenIndex1566
																if buffer[position] != rune('O') {
																	goto l1442
																}
																position++
															}
														l1566:
															{
																position1568, tokenIndex1568 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1569
																}
																position++
																goto l1568
															l1569:
																position, tokenIndex = position1568, tokenIndex1568
																if buffer[position] != rune('P') {
																	goto l1442
																}
																position++
															}
														l1568:
															add(rulePegText, position1563)
														}
														{
															add(ruleAction58, position)
														}
														add(ruleNop, position1562)
													}
													break
												}
											}

										}
									l1444:
										add(ruleSimple, position1443)
									}
									goto l852
								l1442:
									position, tokenIndex = position852, tokenIndex852
									{
										position1572 := position
										{
											position1573, tokenIndex1573 := position, tokenIndex
											{
												position1575 := position
												{
													position1576, tokenIndex1576 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1577
													}
													position++
													goto l1576
												l1577:
													position, tokenIndex = position1576, tokenIndex1576
													if buffer[position] != rune('R') {
														goto l1574
													}
													position++
												}
											l1576:
												{
													position1578, tokenIndex1578 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1579
													}
													position++
													goto l1578
												l1579:
													position, tokenIndex = position1578, tokenIndex1578
													if buffer[position] != rune('S') {
														goto l1574
													}
													position++
												}
											l1578:
												{
													position1580, tokenIndex1580 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1581
													}
													position++
													goto l1580
												l1581:
													position, tokenIndex = position1580, tokenIndex1580
													if buffer[position] != rune('T') {
														goto l1574
													}
													position++
												}
											l1580:
												if !_rules[rulews]() {
													goto l1574
												}
												if !_rules[rulen]() {
													goto l1574
												}
												{
													add(ruleAction95, position)
												}
												add(ruleRst, position1575)
											}
											goto l1573
										l1574:
											position, tokenIndex = position1573, tokenIndex1573
											{
												position1584 := position
												{
													position1585, tokenIndex1585 := position, tokenIndex
													if buffer[position] != rune('j') {
														goto l1586
													}
													position++
													goto l1585
												l1586:
													position, tokenIndex = position1585, tokenIndex1585
													if buffer[position] != rune('J') {
														goto l1583
													}
													position++
												}
											l1585:
												{
													position1587, tokenIndex1587 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l1588
													}
													position++
													goto l1587
												l1588:
													position, tokenIndex = position1587, tokenIndex1587
													if buffer[position] != rune('P') {
														goto l1583
													}
													position++
												}
											l1587:
												if !_rules[rulews]() {
													goto l1583
												}
												{
													position1589, tokenIndex1589 := position, tokenIndex
													if !_rules[rulecc]() {
														goto l1589
													}
													if !_rules[rulesep]() {
														goto l1589
													}
													goto l1590
												l1589:
													position, tokenIndex = position1589, tokenIndex1589
												}
											l1590:
												if !_rules[ruleSrc16]() {
													goto l1583
												}
												{
													add(ruleAction98, position)
												}
												add(ruleJp, position1584)
											}
											goto l1573
										l1583:
											position, tokenIndex = position1573, tokenIndex1573
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position1593 := position
														{
															position1594, tokenIndex1594 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1595
															}
															position++
															goto l1594
														l1595:
															position, tokenIndex = position1594, tokenIndex1594
															if buffer[position] != rune('D') {
																goto l1571
															}
															position++
														}
													l1594:
														{
															position1596, tokenIndex1596 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1597
															}
															position++
															goto l1596
														l1597:
															position, tokenIndex = position1596, tokenIndex1596
															if buffer[position] != rune('J') {
																goto l1571
															}
															position++
														}
													l1596:
														{
															position1598, tokenIndex1598 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1599
															}
															position++
															goto l1598
														l1599:
															position, tokenIndex = position1598, tokenIndex1598
															if buffer[position] != rune('N') {
																goto l1571
															}
															position++
														}
													l1598:
														{
															position1600, tokenIndex1600 := position, tokenIndex
															if buffer[position] != rune('z') {
																goto l1601
															}
															position++
															goto l1600
														l1601:
															position, tokenIndex = position1600, tokenIndex1600
															if buffer[position] != rune('Z') {
																goto l1571
															}
															position++
														}
													l1600:
														if !_rules[rulews]() {
															goto l1571
														}
														if !_rules[ruledisp]() {
															goto l1571
														}
														{
															add(ruleAction100, position)
														}
														add(ruleDjnz, position1593)
													}
													break
												case 'J', 'j':
													{
														position1603 := position
														{
															position1604, tokenIndex1604 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1605
															}
															position++
															goto l1604
														l1605:
															position, tokenIndex = position1604, tokenIndex1604
															if buffer[position] != rune('J') {
																goto l1571
															}
															position++
														}
													l1604:
														{
															position1606, tokenIndex1606 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1607
															}
															position++
															goto l1606
														l1607:
															position, tokenIndex = position1606, tokenIndex1606
															if buffer[position] != rune('R') {
																goto l1571
															}
															position++
														}
													l1606:
														if !_rules[rulews]() {
															goto l1571
														}
														{
															position1608, tokenIndex1608 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1608
															}
															if !_rules[rulesep]() {
																goto l1608
															}
															goto l1609
														l1608:
															position, tokenIndex = position1608, tokenIndex1608
														}
													l1609:
														if !_rules[ruledisp]() {
															goto l1571
														}
														{
															add(ruleAction99, position)
														}
														add(ruleJr, position1603)
													}
													break
												case 'R', 'r':
													{
														position1611 := position
														{
															position1612, tokenIndex1612 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1613
															}
															position++
															goto l1612
														l1613:
															position, tokenIndex = position1612, tokenIndex1612
															if buffer[position] != rune('R') {
																goto l1571
															}
															position++
														}
													l1612:
														{
															position1614, tokenIndex1614 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1615
															}
															position++
															goto l1614
														l1615:
															position, tokenIndex = position1614, tokenIndex1614
															if buffer[position] != rune('E') {
																goto l1571
															}
															position++
														}
													l1614:
														{
															position1616, tokenIndex1616 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1617
															}
															position++
															goto l1616
														l1617:
															position, tokenIndex = position1616, tokenIndex1616
															if buffer[position] != rune('T') {
																goto l1571
															}
															position++
														}
													l1616:
														{
															position1618, tokenIndex1618 := position, tokenIndex
															if !_rules[rulews]() {
																goto l1618
															}
															if !_rules[rulecc]() {
																goto l1618
															}
															goto l1619
														l1618:
															position, tokenIndex = position1618, tokenIndex1618
														}
													l1619:
														{
															add(ruleAction97, position)
														}
														add(ruleRet, position1611)
													}
													break
												default:
													{
														position1621 := position
														{
															position1622, tokenIndex1622 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1623
															}
															position++
															goto l1622
														l1623:
															position, tokenIndex = position1622, tokenIndex1622
															if buffer[position] != rune('C') {
																goto l1571
															}
															position++
														}
													l1622:
														{
															position1624, tokenIndex1624 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1625
															}
															position++
															goto l1624
														l1625:
															position, tokenIndex = position1624, tokenIndex1624
															if buffer[position] != rune('A') {
																goto l1571
															}
															position++
														}
													l1624:
														{
															position1626, tokenIndex1626 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1627
															}
															position++
															goto l1626
														l1627:
															position, tokenIndex = position1626, tokenIndex1626
															if buffer[position] != rune('L') {
																goto l1571
															}
															position++
														}
													l1626:
														{
															position1628, tokenIndex1628 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1629
															}
															position++
															goto l1628
														l1629:
															position, tokenIndex = position1628, tokenIndex1628
															if buffer[position] != rune('L') {
																goto l1571
															}
															position++
														}
													l1628:
														if !_rules[rulews]() {
															goto l1571
														}
														{
															position1630, tokenIndex1630 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1630
															}
															if !_rules[rulesep]() {
																goto l1630
															}
															goto l1631
														l1630:
															position, tokenIndex = position1630, tokenIndex1630
														}
													l1631:
														if !_rules[ruleSrc16]() {
															goto l1571
														}
														{
															add(ruleAction96, position)
														}
														add(ruleCall, position1621)
													}
													break
												}
											}

										}
									l1573:
										add(ruleJump, position1572)
									}
									goto l852
								l1571:
									position, tokenIndex = position852, tokenIndex852
									{
										position1633 := position
										{
											position1634, tokenIndex1634 := position, tokenIndex
											{
												position1636 := position
												{
													position1637, tokenIndex1637 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1638
													}
													position++
													goto l1637
												l1638:
													position, tokenIndex = position1637, tokenIndex1637
													if buffer[position] != rune('I') {
														goto l1635
													}
													position++
												}
											l1637:
												{
													position1639, tokenIndex1639 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1640
													}
													position++
													goto l1639
												l1640:
													position, tokenIndex = position1639, tokenIndex1639
													if buffer[position] != rune('N') {
														goto l1635
													}
													position++
												}
											l1639:
												if !_rules[rulews]() {
													goto l1635
												}
												if !_rules[ruleReg8]() {
													goto l1635
												}
												if !_rules[rulesep]() {
													goto l1635
												}
												if !_rules[rulePort]() {
													goto l1635
												}
												{
													add(ruleAction101, position)
												}
												add(ruleIN, position1636)
											}
											goto l1634
										l1635:
											position, tokenIndex = position1634, tokenIndex1634
											{
												position1642 := position
												{
													position1643, tokenIndex1643 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1644
													}
													position++
													goto l1643
												l1644:
													position, tokenIndex = position1643, tokenIndex1643
													if buffer[position] != rune('O') {
														goto l849
													}
													position++
												}
											l1643:
												{
													position1645, tokenIndex1645 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1646
													}
													position++
													goto l1645
												l1646:
													position, tokenIndex = position1645, tokenIndex1645
													if buffer[position] != rune('U') {
														goto l849
													}
													position++
												}
											l1645:
												{
													position1647, tokenIndex1647 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1648
													}
													position++
													goto l1647
												l1648:
													position, tokenIndex = position1647, tokenIndex1647
													if buffer[position] != rune('T') {
														goto l849
													}
													position++
												}
											l1647:
												if !_rules[rulews]() {
													goto l849
												}
												if !_rules[rulePort]() {
													goto l849
												}
												if !_rules[rulesep]() {
													goto l849
												}
												if !_rules[ruleReg8]() {
													goto l849
												}
												{
													add(ruleAction102, position)
												}
												add(ruleOUT, position1642)
											}
										}
									l1634:
										add(ruleIO, position1633)
									}
								}
							l852:
								add(ruleInstruction, position851)
							}
							goto l850
						l849:
							position, tokenIndex = position849, tokenIndex849
						}
					l850:
						{
							position1650 := position
							{
								position1651, tokenIndex1651 := position, tokenIndex
								if !_rules[rulews]() {
									goto l1651
								}
								goto l1652
							l1651:
								position, tokenIndex = position1651, tokenIndex1651
							}
						l1652:
							{
								position1653, tokenIndex1653 := position, tokenIndex
								{
									position1655 := position
									{
										position1656, tokenIndex1656 := position, tokenIndex
										if buffer[position] != rune(';') {
											goto l1657
										}
										position++
										goto l1656
									l1657:
										position, tokenIndex = position1656, tokenIndex1656
										if buffer[position] != rune('#') {
											goto l1653
										}
										position++
									}
								l1656:
								l1658:
									{
										position1659, tokenIndex1659 := position, tokenIndex
										{
											position1660, tokenIndex1660 := position, tokenIndex
											if buffer[position] != rune('\n') {
												goto l1660
											}
											position++
											goto l1659
										l1660:
											position, tokenIndex = position1660, tokenIndex1660
										}
										if !matchDot() {
											goto l1659
										}
										goto l1658
									l1659:
										position, tokenIndex = position1659, tokenIndex1659
									}
									add(ruleComment, position1655)
								}
								goto l1654
							l1653:
								position, tokenIndex = position1653, tokenIndex1653
							}
						l1654:
							{
								position1661, tokenIndex1661 := position, tokenIndex
								if !_rules[rulews]() {
									goto l1661
								}
								goto l1662
							l1661:
								position, tokenIndex = position1661, tokenIndex1661
							}
						l1662:
							{
								position1663, tokenIndex1663 := position, tokenIndex
								if buffer[position] != rune('\n') {
									goto l1664
								}
								position++
								goto l1663
							l1664:
								position, tokenIndex = position1663, tokenIndex1663
								if buffer[position] != rune(':') {
									goto l3
								}
								position++
							}
						l1663:
							add(ruleLineEnd, position1650)
						}
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position835)
					}
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
		/* 1 Line <- <(Label? ws* Instruction? LineEnd Action0)> */
		nil,
		/* 2 Label <- <(<(alpha alphanum*)> ':' Action1)> */
		nil,
		/* 3 alphanum <- <(alpha / num)> */
		nil,
		/* 4 alpha <- <([a-z] / [A-Z])> */
		func() bool {
			position1669, tokenIndex1669 := position, tokenIndex
			{
				position1670 := position
				{
					position1671, tokenIndex1671 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1672
					}
					position++
					goto l1671
				l1672:
					position, tokenIndex = position1671, tokenIndex1671
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1669
					}
					position++
				}
			l1671:
				add(rulealpha, position1670)
			}
			return true
		l1669:
			position, tokenIndex = position1669, tokenIndex1669
			return false
		},
		/* 5 num <- <[0-9]> */
		nil,
		/* 6 Comment <- <((';' / '#') (!'\n' .)*)> */
		nil,
		/* 7 LineEnd <- <(ws? Comment? ws? ('\n' / ':'))> */
		nil,
		/* 8 Instruction <- <(Assignment / Inc / Dec / Alu16 / Alu / BitOp / EDSimple / Simple / Jump / IO)> */
		nil,
		/* 9 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 10 Load <- <(Load16 / Load8)> */
		nil,
		/* 11 Load8 <- <(('l' / 'L') ('d' / 'D') ws Dst8 sep Src8 Action2)> */
		nil,
		/* 12 Load16 <- <(('l' / 'L') ('d' / 'D') ws Dst16 sep Src16 Action3)> */
		nil,
		/* 13 Push <- <(('p' / 'P') ('u' / 'U') ('s' / 'S') ('h' / 'H') ws Src16 Action4)> */
		nil,
		/* 14 Pop <- <(('p' / 'P') ('o' / 'O') ('p' / 'P') ws Dst16 Action5)> */
		nil,
		/* 15 Ex <- <(('e' / 'E') ('x' / 'X') ws Dst16 sep Src16 Action6)> */
		nil,
		/* 16 Inc <- <(Inc16Indexed8 / Inc16 / Inc8)> */
		nil,
		/* 17 Inc16Indexed8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws ILoc8 Action7)> */
		nil,
		/* 18 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action8)> */
		nil,
		/* 19 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action9)> */
		nil,
		/* 20 Dec <- <(Dec16Indexed8 / Dec16 / Dec8)> */
		nil,
		/* 21 Dec16Indexed8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws ILoc8 Action10)> */
		nil,
		/* 22 Dec8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc8 Action11)> */
		nil,
		/* 23 Dec16 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc16 Action12)> */
		nil,
		/* 24 Alu16 <- <(Add16 / Adc16 / Sbc16)> */
		nil,
		/* 25 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action13)> */
		nil,
		/* 26 Adc16 <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws Dst16 sep Src16 Action14)> */
		nil,
		/* 27 Sbc16 <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws Dst16 sep Src16 Action15)> */
		nil,
		/* 28 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action16)> */
		nil,
		/* 29 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action17)> */
		func() bool {
			position1697, tokenIndex1697 := position, tokenIndex
			{
				position1698 := position
				{
					position1699, tokenIndex1699 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1700
					}
					goto l1699
				l1700:
					position, tokenIndex = position1699, tokenIndex1699
					if !_rules[ruleReg8]() {
						goto l1701
					}
					goto l1699
				l1701:
					position, tokenIndex = position1699, tokenIndex1699
					if !_rules[ruleReg16Contents]() {
						goto l1702
					}
					goto l1699
				l1702:
					position, tokenIndex = position1699, tokenIndex1699
					if !_rules[rulenn_contents]() {
						goto l1697
					}
				}
			l1699:
				{
					add(ruleAction17, position)
				}
				add(ruleSrc8, position1698)
			}
			return true
		l1697:
			position, tokenIndex = position1697, tokenIndex1697
			return false
		},
		/* 30 Loc8 <- <((Reg8 / Reg16Contents) Action18)> */
		func() bool {
			position1704, tokenIndex1704 := position, tokenIndex
			{
				position1705 := position
				{
					position1706, tokenIndex1706 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1707
					}
					goto l1706
				l1707:
					position, tokenIndex = position1706, tokenIndex1706
					if !_rules[ruleReg16Contents]() {
						goto l1704
					}
				}
			l1706:
				{
					add(ruleAction18, position)
				}
				add(ruleLoc8, position1705)
			}
			return true
		l1704:
			position, tokenIndex = position1704, tokenIndex1704
			return false
		},
		/* 31 Copy8 <- <(Reg8 Action19)> */
		func() bool {
			position1709, tokenIndex1709 := position, tokenIndex
			{
				position1710 := position
				if !_rules[ruleReg8]() {
					goto l1709
				}
				{
					add(ruleAction19, position)
				}
				add(ruleCopy8, position1710)
			}
			return true
		l1709:
			position, tokenIndex = position1709, tokenIndex1709
			return false
		},
		/* 32 ILoc8 <- <(IReg8 Action20)> */
		func() bool {
			position1712, tokenIndex1712 := position, tokenIndex
			{
				position1713 := position
				if !_rules[ruleIReg8]() {
					goto l1712
				}
				{
					add(ruleAction20, position)
				}
				add(ruleILoc8, position1713)
			}
			return true
		l1712:
			position, tokenIndex = position1712, tokenIndex1712
			return false
		},
		/* 33 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action21)> */
		func() bool {
			position1715, tokenIndex1715 := position, tokenIndex
			{
				position1716 := position
				{
					position1717 := position
					{
						position1718, tokenIndex1718 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1719
						}
						goto l1718
					l1719:
						position, tokenIndex = position1718, tokenIndex1718
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1721 := position
									{
										position1722, tokenIndex1722 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1723
										}
										position++
										goto l1722
									l1723:
										position, tokenIndex = position1722, tokenIndex1722
										if buffer[position] != rune('R') {
											goto l1715
										}
										position++
									}
								l1722:
									add(ruleR, position1721)
								}
								break
							case 'I', 'i':
								{
									position1724 := position
									{
										position1725, tokenIndex1725 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1726
										}
										position++
										goto l1725
									l1726:
										position, tokenIndex = position1725, tokenIndex1725
										if buffer[position] != rune('I') {
											goto l1715
										}
										position++
									}
								l1725:
									add(ruleI, position1724)
								}
								break
							case 'L', 'l':
								{
									position1727 := position
									{
										position1728, tokenIndex1728 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1729
										}
										position++
										goto l1728
									l1729:
										position, tokenIndex = position1728, tokenIndex1728
										if buffer[position] != rune('L') {
											goto l1715
										}
										position++
									}
								l1728:
									add(ruleL, position1727)
								}
								break
							case 'H', 'h':
								{
									position1730 := position
									{
										position1731, tokenIndex1731 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1732
										}
										position++
										goto l1731
									l1732:
										position, tokenIndex = position1731, tokenIndex1731
										if buffer[position] != rune('H') {
											goto l1715
										}
										position++
									}
								l1731:
									add(ruleH, position1730)
								}
								break
							case 'E', 'e':
								{
									position1733 := position
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
											goto l1715
										}
										position++
									}
								l1734:
									add(ruleE, position1733)
								}
								break
							case 'D', 'd':
								{
									position1736 := position
									{
										position1737, tokenIndex1737 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1738
										}
										position++
										goto l1737
									l1738:
										position, tokenIndex = position1737, tokenIndex1737
										if buffer[position] != rune('D') {
											goto l1715
										}
										position++
									}
								l1737:
									add(ruleD, position1736)
								}
								break
							case 'C', 'c':
								{
									position1739 := position
									{
										position1740, tokenIndex1740 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1741
										}
										position++
										goto l1740
									l1741:
										position, tokenIndex = position1740, tokenIndex1740
										if buffer[position] != rune('C') {
											goto l1715
										}
										position++
									}
								l1740:
									add(ruleC, position1739)
								}
								break
							case 'B', 'b':
								{
									position1742 := position
									{
										position1743, tokenIndex1743 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1744
										}
										position++
										goto l1743
									l1744:
										position, tokenIndex = position1743, tokenIndex1743
										if buffer[position] != rune('B') {
											goto l1715
										}
										position++
									}
								l1743:
									add(ruleB, position1742)
								}
								break
							case 'F', 'f':
								{
									position1745 := position
									{
										position1746, tokenIndex1746 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1747
										}
										position++
										goto l1746
									l1747:
										position, tokenIndex = position1746, tokenIndex1746
										if buffer[position] != rune('F') {
											goto l1715
										}
										position++
									}
								l1746:
									add(ruleF, position1745)
								}
								break
							default:
								{
									position1748 := position
									{
										position1749, tokenIndex1749 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1750
										}
										position++
										goto l1749
									l1750:
										position, tokenIndex = position1749, tokenIndex1749
										if buffer[position] != rune('A') {
											goto l1715
										}
										position++
									}
								l1749:
									add(ruleA, position1748)
								}
								break
							}
						}

					}
				l1718:
					add(rulePegText, position1717)
				}
				{
					add(ruleAction21, position)
				}
				add(ruleReg8, position1716)
			}
			return true
		l1715:
			position, tokenIndex = position1715, tokenIndex1715
			return false
		},
		/* 34 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action22)> */
		func() bool {
			position1752, tokenIndex1752 := position, tokenIndex
			{
				position1753 := position
				{
					position1754 := position
					{
						position1755, tokenIndex1755 := position, tokenIndex
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
									goto l1756
								}
								position++
							}
						l1758:
							{
								position1760, tokenIndex1760 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1761
								}
								position++
								goto l1760
							l1761:
								position, tokenIndex = position1760, tokenIndex1760
								if buffer[position] != rune('X') {
									goto l1756
								}
								position++
							}
						l1760:
							{
								position1762, tokenIndex1762 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1763
								}
								position++
								goto l1762
							l1763:
								position, tokenIndex = position1762, tokenIndex1762
								if buffer[position] != rune('H') {
									goto l1756
								}
								position++
							}
						l1762:
							add(ruleIXH, position1757)
						}
						goto l1755
					l1756:
						position, tokenIndex = position1755, tokenIndex1755
						{
							position1765 := position
							{
								position1766, tokenIndex1766 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1767
								}
								position++
								goto l1766
							l1767:
								position, tokenIndex = position1766, tokenIndex1766
								if buffer[position] != rune('I') {
									goto l1764
								}
								position++
							}
						l1766:
							{
								position1768, tokenIndex1768 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1769
								}
								position++
								goto l1768
							l1769:
								position, tokenIndex = position1768, tokenIndex1768
								if buffer[position] != rune('X') {
									goto l1764
								}
								position++
							}
						l1768:
							{
								position1770, tokenIndex1770 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1771
								}
								position++
								goto l1770
							l1771:
								position, tokenIndex = position1770, tokenIndex1770
								if buffer[position] != rune('L') {
									goto l1764
								}
								position++
							}
						l1770:
							add(ruleIXL, position1765)
						}
						goto l1755
					l1764:
						position, tokenIndex = position1755, tokenIndex1755
						{
							position1773 := position
							{
								position1774, tokenIndex1774 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1775
								}
								position++
								goto l1774
							l1775:
								position, tokenIndex = position1774, tokenIndex1774
								if buffer[position] != rune('I') {
									goto l1772
								}
								position++
							}
						l1774:
							{
								position1776, tokenIndex1776 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1777
								}
								position++
								goto l1776
							l1777:
								position, tokenIndex = position1776, tokenIndex1776
								if buffer[position] != rune('Y') {
									goto l1772
								}
								position++
							}
						l1776:
							{
								position1778, tokenIndex1778 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1779
								}
								position++
								goto l1778
							l1779:
								position, tokenIndex = position1778, tokenIndex1778
								if buffer[position] != rune('H') {
									goto l1772
								}
								position++
							}
						l1778:
							add(ruleIYH, position1773)
						}
						goto l1755
					l1772:
						position, tokenIndex = position1755, tokenIndex1755
						{
							position1780 := position
							{
								position1781, tokenIndex1781 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1782
								}
								position++
								goto l1781
							l1782:
								position, tokenIndex = position1781, tokenIndex1781
								if buffer[position] != rune('I') {
									goto l1752
								}
								position++
							}
						l1781:
							{
								position1783, tokenIndex1783 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1784
								}
								position++
								goto l1783
							l1784:
								position, tokenIndex = position1783, tokenIndex1783
								if buffer[position] != rune('Y') {
									goto l1752
								}
								position++
							}
						l1783:
							{
								position1785, tokenIndex1785 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1786
								}
								position++
								goto l1785
							l1786:
								position, tokenIndex = position1785, tokenIndex1785
								if buffer[position] != rune('L') {
									goto l1752
								}
								position++
							}
						l1785:
							add(ruleIYL, position1780)
						}
					}
				l1755:
					add(rulePegText, position1754)
				}
				{
					add(ruleAction22, position)
				}
				add(ruleIReg8, position1753)
			}
			return true
		l1752:
			position, tokenIndex = position1752, tokenIndex1752
			return false
		},
		/* 35 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action23)> */
		func() bool {
			position1788, tokenIndex1788 := position, tokenIndex
			{
				position1789 := position
				{
					position1790, tokenIndex1790 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1791
					}
					goto l1790
				l1791:
					position, tokenIndex = position1790, tokenIndex1790
					if !_rules[rulenn_contents]() {
						goto l1792
					}
					goto l1790
				l1792:
					position, tokenIndex = position1790, tokenIndex1790
					if !_rules[ruleReg16Contents]() {
						goto l1788
					}
				}
			l1790:
				{
					add(ruleAction23, position)
				}
				add(ruleDst16, position1789)
			}
			return true
		l1788:
			position, tokenIndex = position1788, tokenIndex1788
			return false
		},
		/* 36 Src16 <- <((Reg16 / nn / nn_contents) Action24)> */
		func() bool {
			position1794, tokenIndex1794 := position, tokenIndex
			{
				position1795 := position
				{
					position1796, tokenIndex1796 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1797
					}
					goto l1796
				l1797:
					position, tokenIndex = position1796, tokenIndex1796
					if !_rules[rulenn]() {
						goto l1798
					}
					goto l1796
				l1798:
					position, tokenIndex = position1796, tokenIndex1796
					if !_rules[rulenn_contents]() {
						goto l1794
					}
				}
			l1796:
				{
					add(ruleAction24, position)
				}
				add(ruleSrc16, position1795)
			}
			return true
		l1794:
			position, tokenIndex = position1794, tokenIndex1794
			return false
		},
		/* 37 Loc16 <- <(Reg16 Action25)> */
		func() bool {
			position1800, tokenIndex1800 := position, tokenIndex
			{
				position1801 := position
				if !_rules[ruleReg16]() {
					goto l1800
				}
				{
					add(ruleAction25, position)
				}
				add(ruleLoc16, position1801)
			}
			return true
		l1800:
			position, tokenIndex = position1800, tokenIndex1800
			return false
		},
		/* 38 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action26)> */
		func() bool {
			position1803, tokenIndex1803 := position, tokenIndex
			{
				position1804 := position
				{
					position1805 := position
					{
						position1806, tokenIndex1806 := position, tokenIndex
						{
							position1808 := position
							{
								position1809, tokenIndex1809 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1810
								}
								position++
								goto l1809
							l1810:
								position, tokenIndex = position1809, tokenIndex1809
								if buffer[position] != rune('A') {
									goto l1807
								}
								position++
							}
						l1809:
							{
								position1811, tokenIndex1811 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1812
								}
								position++
								goto l1811
							l1812:
								position, tokenIndex = position1811, tokenIndex1811
								if buffer[position] != rune('F') {
									goto l1807
								}
								position++
							}
						l1811:
							if buffer[position] != rune('\'') {
								goto l1807
							}
							position++
							add(ruleAF_PRIME, position1808)
						}
						goto l1806
					l1807:
						position, tokenIndex = position1806, tokenIndex1806
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1803
								}
								break
							case 'S', 's':
								{
									position1814 := position
									{
										position1815, tokenIndex1815 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1816
										}
										position++
										goto l1815
									l1816:
										position, tokenIndex = position1815, tokenIndex1815
										if buffer[position] != rune('S') {
											goto l1803
										}
										position++
									}
								l1815:
									{
										position1817, tokenIndex1817 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1818
										}
										position++
										goto l1817
									l1818:
										position, tokenIndex = position1817, tokenIndex1817
										if buffer[position] != rune('P') {
											goto l1803
										}
										position++
									}
								l1817:
									add(ruleSP, position1814)
								}
								break
							case 'H', 'h':
								{
									position1819 := position
									{
										position1820, tokenIndex1820 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1821
										}
										position++
										goto l1820
									l1821:
										position, tokenIndex = position1820, tokenIndex1820
										if buffer[position] != rune('H') {
											goto l1803
										}
										position++
									}
								l1820:
									{
										position1822, tokenIndex1822 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1823
										}
										position++
										goto l1822
									l1823:
										position, tokenIndex = position1822, tokenIndex1822
										if buffer[position] != rune('L') {
											goto l1803
										}
										position++
									}
								l1822:
									add(ruleHL, position1819)
								}
								break
							case 'D', 'd':
								{
									position1824 := position
									{
										position1825, tokenIndex1825 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1826
										}
										position++
										goto l1825
									l1826:
										position, tokenIndex = position1825, tokenIndex1825
										if buffer[position] != rune('D') {
											goto l1803
										}
										position++
									}
								l1825:
									{
										position1827, tokenIndex1827 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1828
										}
										position++
										goto l1827
									l1828:
										position, tokenIndex = position1827, tokenIndex1827
										if buffer[position] != rune('E') {
											goto l1803
										}
										position++
									}
								l1827:
									add(ruleDE, position1824)
								}
								break
							case 'B', 'b':
								{
									position1829 := position
									{
										position1830, tokenIndex1830 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1831
										}
										position++
										goto l1830
									l1831:
										position, tokenIndex = position1830, tokenIndex1830
										if buffer[position] != rune('B') {
											goto l1803
										}
										position++
									}
								l1830:
									{
										position1832, tokenIndex1832 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1833
										}
										position++
										goto l1832
									l1833:
										position, tokenIndex = position1832, tokenIndex1832
										if buffer[position] != rune('C') {
											goto l1803
										}
										position++
									}
								l1832:
									add(ruleBC, position1829)
								}
								break
							default:
								{
									position1834 := position
									{
										position1835, tokenIndex1835 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1836
										}
										position++
										goto l1835
									l1836:
										position, tokenIndex = position1835, tokenIndex1835
										if buffer[position] != rune('A') {
											goto l1803
										}
										position++
									}
								l1835:
									{
										position1837, tokenIndex1837 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1838
										}
										position++
										goto l1837
									l1838:
										position, tokenIndex = position1837, tokenIndex1837
										if buffer[position] != rune('F') {
											goto l1803
										}
										position++
									}
								l1837:
									add(ruleAF, position1834)
								}
								break
							}
						}

					}
				l1806:
					add(rulePegText, position1805)
				}
				{
					add(ruleAction26, position)
				}
				add(ruleReg16, position1804)
			}
			return true
		l1803:
			position, tokenIndex = position1803, tokenIndex1803
			return false
		},
		/* 39 IReg16 <- <(<(IX / IY)> Action27)> */
		func() bool {
			position1840, tokenIndex1840 := position, tokenIndex
			{
				position1841 := position
				{
					position1842 := position
					{
						position1843, tokenIndex1843 := position, tokenIndex
						{
							position1845 := position
							{
								position1846, tokenIndex1846 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1847
								}
								position++
								goto l1846
							l1847:
								position, tokenIndex = position1846, tokenIndex1846
								if buffer[position] != rune('I') {
									goto l1844
								}
								position++
							}
						l1846:
							{
								position1848, tokenIndex1848 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1849
								}
								position++
								goto l1848
							l1849:
								position, tokenIndex = position1848, tokenIndex1848
								if buffer[position] != rune('X') {
									goto l1844
								}
								position++
							}
						l1848:
							add(ruleIX, position1845)
						}
						goto l1843
					l1844:
						position, tokenIndex = position1843, tokenIndex1843
						{
							position1850 := position
							{
								position1851, tokenIndex1851 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1852
								}
								position++
								goto l1851
							l1852:
								position, tokenIndex = position1851, tokenIndex1851
								if buffer[position] != rune('I') {
									goto l1840
								}
								position++
							}
						l1851:
							{
								position1853, tokenIndex1853 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1854
								}
								position++
								goto l1853
							l1854:
								position, tokenIndex = position1853, tokenIndex1853
								if buffer[position] != rune('Y') {
									goto l1840
								}
								position++
							}
						l1853:
							add(ruleIY, position1850)
						}
					}
				l1843:
					add(rulePegText, position1842)
				}
				{
					add(ruleAction27, position)
				}
				add(ruleIReg16, position1841)
			}
			return true
		l1840:
			position, tokenIndex = position1840, tokenIndex1840
			return false
		},
		/* 40 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1856, tokenIndex1856 := position, tokenIndex
			{
				position1857 := position
				{
					position1858, tokenIndex1858 := position, tokenIndex
					{
						position1860 := position
						if buffer[position] != rune('(') {
							goto l1859
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1859
						}
						{
							position1861, tokenIndex1861 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1861
							}
							goto l1862
						l1861:
							position, tokenIndex = position1861, tokenIndex1861
						}
					l1862:
						if !_rules[ruledisp]() {
							goto l1859
						}
						{
							position1863, tokenIndex1863 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1863
							}
							goto l1864
						l1863:
							position, tokenIndex = position1863, tokenIndex1863
						}
					l1864:
						if buffer[position] != rune(')') {
							goto l1859
						}
						position++
						{
							add(ruleAction29, position)
						}
						add(ruleIndexedR16C, position1860)
					}
					goto l1858
				l1859:
					position, tokenIndex = position1858, tokenIndex1858
					{
						position1866 := position
						if buffer[position] != rune('(') {
							goto l1856
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1856
						}
						if buffer[position] != rune(')') {
							goto l1856
						}
						position++
						{
							add(ruleAction28, position)
						}
						add(rulePlainR16C, position1866)
					}
				}
			l1858:
				add(ruleReg16Contents, position1857)
			}
			return true
		l1856:
			position, tokenIndex = position1856, tokenIndex1856
			return false
		},
		/* 41 PlainR16C <- <('(' Reg16 ')' Action28)> */
		nil,
		/* 42 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action29)> */
		nil,
		/* 43 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1870, tokenIndex1870 := position, tokenIndex
			{
				position1871 := position
				{
					position1872, tokenIndex1872 := position, tokenIndex
					{
						position1874 := position
						{
							position1875 := position
							if !_rules[rulehexdigit]() {
								goto l1873
							}
							if !_rules[rulehexdigit]() {
								goto l1873
							}
							add(rulePegText, position1875)
						}
						{
							position1876, tokenIndex1876 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1877
							}
							position++
							goto l1876
						l1877:
							position, tokenIndex = position1876, tokenIndex1876
							if buffer[position] != rune('H') {
								goto l1873
							}
							position++
						}
					l1876:
						{
							add(ruleAction33, position)
						}
						add(rulehexByteH, position1874)
					}
					goto l1872
				l1873:
					position, tokenIndex = position1872, tokenIndex1872
					{
						position1880 := position
						if buffer[position] != rune('0') {
							goto l1879
						}
						position++
						{
							position1881, tokenIndex1881 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1882
							}
							position++
							goto l1881
						l1882:
							position, tokenIndex = position1881, tokenIndex1881
							if buffer[position] != rune('X') {
								goto l1879
							}
							position++
						}
					l1881:
						{
							position1883 := position
							if !_rules[rulehexdigit]() {
								goto l1879
							}
							if !_rules[rulehexdigit]() {
								goto l1879
							}
							add(rulePegText, position1883)
						}
						{
							add(ruleAction34, position)
						}
						add(rulehexByte0x, position1880)
					}
					goto l1872
				l1879:
					position, tokenIndex = position1872, tokenIndex1872
					{
						position1885 := position
						{
							position1886 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1870
							}
							position++
						l1887:
							{
								position1888, tokenIndex1888 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1888
								}
								position++
								goto l1887
							l1888:
								position, tokenIndex = position1888, tokenIndex1888
							}
							add(rulePegText, position1886)
						}
						{
							add(ruleAction35, position)
						}
						add(ruledecimalByte, position1885)
					}
				}
			l1872:
				add(rulen, position1871)
			}
			return true
		l1870:
			position, tokenIndex = position1870, tokenIndex1870
			return false
		},
		/* 44 nn <- <(hexWordH / hexWord0x)> */
		func() bool {
			position1890, tokenIndex1890 := position, tokenIndex
			{
				position1891 := position
				{
					position1892, tokenIndex1892 := position, tokenIndex
					{
						position1894 := position
						{
							position1895 := position
							if !_rules[rulehexdigit]() {
								goto l1893
							}
							if !_rules[rulehexdigit]() {
								goto l1893
							}
							if !_rules[rulehexdigit]() {
								goto l1893
							}
							if !_rules[rulehexdigit]() {
								goto l1893
							}
							add(rulePegText, position1895)
						}
						{
							position1896, tokenIndex1896 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1897
							}
							position++
							goto l1896
						l1897:
							position, tokenIndex = position1896, tokenIndex1896
							if buffer[position] != rune('H') {
								goto l1893
							}
							position++
						}
					l1896:
						{
							add(ruleAction36, position)
						}
						add(rulehexWordH, position1894)
					}
					goto l1892
				l1893:
					position, tokenIndex = position1892, tokenIndex1892
					{
						position1899 := position
						if buffer[position] != rune('0') {
							goto l1890
						}
						position++
						{
							position1900, tokenIndex1900 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1901
							}
							position++
							goto l1900
						l1901:
							position, tokenIndex = position1900, tokenIndex1900
							if buffer[position] != rune('X') {
								goto l1890
							}
							position++
						}
					l1900:
						{
							position1902 := position
							if !_rules[rulehexdigit]() {
								goto l1890
							}
							if !_rules[rulehexdigit]() {
								goto l1890
							}
							if !_rules[rulehexdigit]() {
								goto l1890
							}
							if !_rules[rulehexdigit]() {
								goto l1890
							}
							add(rulePegText, position1902)
						}
						{
							add(ruleAction37, position)
						}
						add(rulehexWord0x, position1899)
					}
				}
			l1892:
				add(rulenn, position1891)
			}
			return true
		l1890:
			position, tokenIndex = position1890, tokenIndex1890
			return false
		},
		/* 45 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position1904, tokenIndex1904 := position, tokenIndex
			{
				position1905 := position
				{
					position1906, tokenIndex1906 := position, tokenIndex
					{
						position1908 := position
						{
							position1909 := position
							{
								position1910, tokenIndex1910 := position, tokenIndex
								{
									position1912, tokenIndex1912 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1913
									}
									position++
									goto l1912
								l1913:
									position, tokenIndex = position1912, tokenIndex1912
									if buffer[position] != rune('+') {
										goto l1910
									}
									position++
								}
							l1912:
								goto l1911
							l1910:
								position, tokenIndex = position1910, tokenIndex1910
							}
						l1911:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1907
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1907
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1907
									}
									position++
									break
								}
							}

						l1914:
							{
								position1915, tokenIndex1915 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1915
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1915
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1915
										}
										position++
										break
									}
								}

								goto l1914
							l1915:
								position, tokenIndex = position1915, tokenIndex1915
							}
							add(rulePegText, position1909)
						}
						{
							position1918, tokenIndex1918 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1919
							}
							position++
							goto l1918
						l1919:
							position, tokenIndex = position1918, tokenIndex1918
							if buffer[position] != rune('H') {
								goto l1907
							}
							position++
						}
					l1918:
						{
							add(ruleAction31, position)
						}
						add(rulesignedHexByteH, position1908)
					}
					goto l1906
				l1907:
					position, tokenIndex = position1906, tokenIndex1906
					{
						position1922 := position
						{
							position1923 := position
							{
								position1924, tokenIndex1924 := position, tokenIndex
								{
									position1926, tokenIndex1926 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1927
									}
									position++
									goto l1926
								l1927:
									position, tokenIndex = position1926, tokenIndex1926
									if buffer[position] != rune('+') {
										goto l1924
									}
									position++
								}
							l1926:
								goto l1925
							l1924:
								position, tokenIndex = position1924, tokenIndex1924
							}
						l1925:
							if buffer[position] != rune('0') {
								goto l1921
							}
							position++
							{
								position1928, tokenIndex1928 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1929
								}
								position++
								goto l1928
							l1929:
								position, tokenIndex = position1928, tokenIndex1928
								if buffer[position] != rune('X') {
									goto l1921
								}
								position++
							}
						l1928:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1921
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1921
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1921
									}
									position++
									break
								}
							}

						l1930:
							{
								position1931, tokenIndex1931 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1931
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1931
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1931
										}
										position++
										break
									}
								}

								goto l1930
							l1931:
								position, tokenIndex = position1931, tokenIndex1931
							}
							add(rulePegText, position1923)
						}
						{
							add(ruleAction32, position)
						}
						add(rulesignedHexByte0x, position1922)
					}
					goto l1906
				l1921:
					position, tokenIndex = position1906, tokenIndex1906
					{
						position1935 := position
						{
							position1936 := position
							{
								position1937, tokenIndex1937 := position, tokenIndex
								{
									position1939, tokenIndex1939 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1940
									}
									position++
									goto l1939
								l1940:
									position, tokenIndex = position1939, tokenIndex1939
									if buffer[position] != rune('+') {
										goto l1937
									}
									position++
								}
							l1939:
								goto l1938
							l1937:
								position, tokenIndex = position1937, tokenIndex1937
							}
						l1938:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1904
							}
							position++
						l1941:
							{
								position1942, tokenIndex1942 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1942
								}
								position++
								goto l1941
							l1942:
								position, tokenIndex = position1942, tokenIndex1942
							}
							add(rulePegText, position1936)
						}
						{
							add(ruleAction30, position)
						}
						add(rulesignedDecimalByte, position1935)
					}
				}
			l1906:
				add(ruledisp, position1905)
			}
			return true
		l1904:
			position, tokenIndex = position1904, tokenIndex1904
			return false
		},
		/* 46 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action30)> */
		nil,
		/* 47 signedHexByteH <- <(<(('-' / '+')? ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> ('h' / 'H') Action31)> */
		nil,
		/* 48 signedHexByte0x <- <(<(('-' / '+')? ('0' ('x' / 'X')) ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> Action32)> */
		nil,
		/* 49 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action33)> */
		nil,
		/* 50 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action34)> */
		nil,
		/* 51 decimalByte <- <(<[0-9]+> Action35)> */
		nil,
		/* 52 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action36)> */
		nil,
		/* 53 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action37)> */
		nil,
		/* 54 nn_contents <- <('(' nn ')' Action38)> */
		func() bool {
			position1952, tokenIndex1952 := position, tokenIndex
			{
				position1953 := position
				if buffer[position] != rune('(') {
					goto l1952
				}
				position++
				if !_rules[rulenn]() {
					goto l1952
				}
				if buffer[position] != rune(')') {
					goto l1952
				}
				position++
				{
					add(ruleAction38, position)
				}
				add(rulenn_contents, position1953)
			}
			return true
		l1952:
			position, tokenIndex = position1952, tokenIndex1952
			return false
		},
		/* 55 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 56 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action39)> */
		nil,
		/* 57 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action40)> */
		nil,
		/* 58 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action41)> */
		nil,
		/* 59 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action42)> */
		nil,
		/* 60 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action43)> */
		nil,
		/* 61 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action44)> */
		nil,
		/* 62 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action45)> */
		nil,
		/* 63 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action46)> */
		nil,
		/* 64 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 65 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 66 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 (sep Copy8)? Action47)> */
		nil,
		/* 67 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 (sep Copy8)? Action48)> */
		nil,
		/* 68 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action49)> */
		nil,
		/* 69 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 (sep Copy8)? Action50)> */
		nil,
		/* 70 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 (sep Copy8)? Action51)> */
		nil,
		/* 71 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 (sep Copy8)? Action52)> */
		nil,
		/* 72 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 (sep Copy8)? Action53)> */
		nil,
		/* 73 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action54)> */
		nil,
		/* 74 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action55)> */
		nil,
		/* 75 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 (sep Copy8)? Action56)> */
		nil,
		/* 76 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 (sep Copy8)? Action57)> */
		nil,
		/* 77 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 78 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action58)> */
		nil,
		/* 79 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action59)> */
		nil,
		/* 80 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action60)> */
		nil,
		/* 81 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action61)> */
		nil,
		/* 82 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action62)> */
		nil,
		/* 83 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action63)> */
		nil,
		/* 84 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action64)> */
		nil,
		/* 85 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action65)> */
		nil,
		/* 86 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action66)> */
		nil,
		/* 87 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action67)> */
		nil,
		/* 88 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action68)> */
		nil,
		/* 89 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action69)> */
		nil,
		/* 90 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action70)> */
		nil,
		/* 91 EDSimple <- <(Retn / Reti / Rrd / Im0 / Im1 / Im2 / ((&('I' | 'O' | 'i' | 'o') BlitIO) | (&('R' | 'r') Rld) | (&('N' | 'n') Neg) | (&('C' | 'L' | 'c' | 'l') Blit)))> */
		nil,
		/* 92 Neg <- <(<(('n' / 'N') ('e' / 'E') ('g' / 'G'))> Action71)> */
		nil,
		/* 93 Retn <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N'))> Action72)> */
		nil,
		/* 94 Reti <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('i' / 'I'))> Action73)> */
		nil,
		/* 95 Rrd <- <(<(('r' / 'R') ('r' / 'R') ('d' / 'D'))> Action74)> */
		nil,
		/* 96 Rld <- <(<(('r' / 'R') ('l' / 'L') ('d' / 'D'))> Action75)> */
		nil,
		/* 97 Im0 <- <(<(('i' / 'I') ('m' / 'M') ' ' '0')> Action76)> */
		nil,
		/* 98 Im1 <- <(<(('i' / 'I') ('m' / 'M') ' ' '1')> Action77)> */
		nil,
		/* 99 Im2 <- <(<(('i' / 'I') ('m' / 'M') ' ' '2')> Action78)> */
		nil,
		/* 100 Blit <- <(Ldir / Ldi / Cpir / Cpi / Lddr / Ldd / Cpdr / Cpd)> */
		nil,
		/* 101 BlitIO <- <(Inir / Ini / Otir / Outi / Indr / Ind / Otdr / Outd)> */
		nil,
		/* 102 Ldi <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I'))> Action79)> */
		nil,
		/* 103 Cpi <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I'))> Action80)> */
		nil,
		/* 104 Ini <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I'))> Action81)> */
		nil,
		/* 105 Outi <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('i' / 'I'))> Action82)> */
		nil,
		/* 106 Ldd <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D'))> Action83)> */
		nil,
		/* 107 Cpd <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D'))> Action84)> */
		nil,
		/* 108 Ind <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D'))> Action85)> */
		nil,
		/* 109 Outd <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('d' / 'D'))> Action86)> */
		nil,
		/* 110 Ldir <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I') ('r' / 'R'))> Action87)> */
		nil,
		/* 111 Cpir <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I') ('r' / 'R'))> Action88)> */
		nil,
		/* 112 Inir <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I') ('r' / 'R'))> Action89)> */
		nil,
		/* 113 Otir <- <(<(('o' / 'O') ('t' / 'T') ('i' / 'I') ('r' / 'R'))> Action90)> */
		nil,
		/* 114 Lddr <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D') ('r' / 'R'))> Action91)> */
		nil,
		/* 115 Cpdr <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D') ('r' / 'R'))> Action92)> */
		nil,
		/* 116 Indr <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D') ('r' / 'R'))> Action93)> */
		nil,
		/* 117 Otdr <- <(<(('o' / 'O') ('t' / 'T') ('d' / 'D') ('r' / 'R'))> Action94)> */
		nil,
		/* 118 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 119 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action95)> */
		nil,
		/* 120 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action96)> */
		nil,
		/* 121 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action97)> */
		nil,
		/* 122 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action98)> */
		nil,
		/* 123 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action99)> */
		nil,
		/* 124 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action100)> */
		nil,
		/* 125 IO <- <(IN / OUT)> */
		nil,
		/* 126 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action101)> */
		nil,
		/* 127 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action102)> */
		nil,
		/* 128 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position2028, tokenIndex2028 := position, tokenIndex
			{
				position2029 := position
				{
					position2030, tokenIndex2030 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2031
					}
					position++
					{
						position2032, tokenIndex2032 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l2033
						}
						position++
						goto l2032
					l2033:
						position, tokenIndex = position2032, tokenIndex2032
						if buffer[position] != rune('C') {
							goto l2031
						}
						position++
					}
				l2032:
					if buffer[position] != rune(')') {
						goto l2031
					}
					position++
					goto l2030
				l2031:
					position, tokenIndex = position2030, tokenIndex2030
					if buffer[position] != rune('(') {
						goto l2028
					}
					position++
					if !_rules[rulen]() {
						goto l2028
					}
					if buffer[position] != rune(')') {
						goto l2028
					}
					position++
				}
			l2030:
				add(rulePort, position2029)
			}
			return true
		l2028:
			position, tokenIndex = position2028, tokenIndex2028
			return false
		},
		/* 129 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2034, tokenIndex2034 := position, tokenIndex
			{
				position2035 := position
				{
					position2036, tokenIndex2036 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2036
					}
					goto l2037
				l2036:
					position, tokenIndex = position2036, tokenIndex2036
				}
			l2037:
				if buffer[position] != rune(',') {
					goto l2034
				}
				position++
				{
					position2038, tokenIndex2038 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2038
					}
					goto l2039
				l2038:
					position, tokenIndex = position2038, tokenIndex2038
				}
			l2039:
				add(rulesep, position2035)
			}
			return true
		l2034:
			position, tokenIndex = position2034, tokenIndex2034
			return false
		},
		/* 130 ws <- <' '+> */
		func() bool {
			position2040, tokenIndex2040 := position, tokenIndex
			{
				position2041 := position
				if buffer[position] != rune(' ') {
					goto l2040
				}
				position++
			l2042:
				{
					position2043, tokenIndex2043 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2043
					}
					position++
					goto l2042
				l2043:
					position, tokenIndex = position2043, tokenIndex2043
				}
				add(rulews, position2041)
			}
			return true
		l2040:
			position, tokenIndex = position2040, tokenIndex2040
			return false
		},
		/* 131 A <- <('a' / 'A')> */
		nil,
		/* 132 F <- <('f' / 'F')> */
		nil,
		/* 133 B <- <('b' / 'B')> */
		nil,
		/* 134 C <- <('c' / 'C')> */
		nil,
		/* 135 D <- <('d' / 'D')> */
		nil,
		/* 136 E <- <('e' / 'E')> */
		nil,
		/* 137 H <- <('h' / 'H')> */
		nil,
		/* 138 L <- <('l' / 'L')> */
		nil,
		/* 139 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 140 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 141 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 142 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 143 I <- <('i' / 'I')> */
		nil,
		/* 144 R <- <('r' / 'R')> */
		nil,
		/* 145 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 146 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 147 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 148 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 149 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 150 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 151 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 152 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 153 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position2066, tokenIndex2066 := position, tokenIndex
			{
				position2067 := position
				{
					position2068, tokenIndex2068 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2069
					}
					position++
					goto l2068
				l2069:
					position, tokenIndex = position2068, tokenIndex2068
					{
						position2070, tokenIndex2070 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2071
						}
						position++
						goto l2070
					l2071:
						position, tokenIndex = position2070, tokenIndex2070
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2066
						}
						position++
					}
				l2070:
				}
			l2068:
				add(rulehexdigit, position2067)
			}
			return true
		l2066:
			position, tokenIndex = position2066, tokenIndex2066
			return false
		},
		/* 154 octaldigit <- <(<[0-7]> Action103)> */
		func() bool {
			position2072, tokenIndex2072 := position, tokenIndex
			{
				position2073 := position
				{
					position2074 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2072
					}
					position++
					add(rulePegText, position2074)
				}
				{
					add(ruleAction103, position)
				}
				add(ruleoctaldigit, position2073)
			}
			return true
		l2072:
			position, tokenIndex = position2072, tokenIndex2072
			return false
		},
		/* 155 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2076, tokenIndex2076 := position, tokenIndex
			{
				position2077 := position
				{
					position2078, tokenIndex2078 := position, tokenIndex
					{
						position2080 := position
						{
							position2081, tokenIndex2081 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2082
							}
							position++
							goto l2081
						l2082:
							position, tokenIndex = position2081, tokenIndex2081
							if buffer[position] != rune('N') {
								goto l2079
							}
							position++
						}
					l2081:
						{
							position2083, tokenIndex2083 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2084
							}
							position++
							goto l2083
						l2084:
							position, tokenIndex = position2083, tokenIndex2083
							if buffer[position] != rune('Z') {
								goto l2079
							}
							position++
						}
					l2083:
						{
							add(ruleAction104, position)
						}
						add(ruleFT_NZ, position2080)
					}
					goto l2078
				l2079:
					position, tokenIndex = position2078, tokenIndex2078
					{
						position2087 := position
						{
							position2088, tokenIndex2088 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2089
							}
							position++
							goto l2088
						l2089:
							position, tokenIndex = position2088, tokenIndex2088
							if buffer[position] != rune('P') {
								goto l2086
							}
							position++
						}
					l2088:
						{
							position2090, tokenIndex2090 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2091
							}
							position++
							goto l2090
						l2091:
							position, tokenIndex = position2090, tokenIndex2090
							if buffer[position] != rune('O') {
								goto l2086
							}
							position++
						}
					l2090:
						{
							add(ruleAction108, position)
						}
						add(ruleFT_PO, position2087)
					}
					goto l2078
				l2086:
					position, tokenIndex = position2078, tokenIndex2078
					{
						position2094 := position
						{
							position2095, tokenIndex2095 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2096
							}
							position++
							goto l2095
						l2096:
							position, tokenIndex = position2095, tokenIndex2095
							if buffer[position] != rune('P') {
								goto l2093
							}
							position++
						}
					l2095:
						{
							position2097, tokenIndex2097 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2098
							}
							position++
							goto l2097
						l2098:
							position, tokenIndex = position2097, tokenIndex2097
							if buffer[position] != rune('E') {
								goto l2093
							}
							position++
						}
					l2097:
						{
							add(ruleAction109, position)
						}
						add(ruleFT_PE, position2094)
					}
					goto l2078
				l2093:
					position, tokenIndex = position2078, tokenIndex2078
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2101 := position
								{
									position2102, tokenIndex2102 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2103
									}
									position++
									goto l2102
								l2103:
									position, tokenIndex = position2102, tokenIndex2102
									if buffer[position] != rune('M') {
										goto l2076
									}
									position++
								}
							l2102:
								{
									add(ruleAction111, position)
								}
								add(ruleFT_M, position2101)
							}
							break
						case 'P', 'p':
							{
								position2105 := position
								{
									position2106, tokenIndex2106 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2107
									}
									position++
									goto l2106
								l2107:
									position, tokenIndex = position2106, tokenIndex2106
									if buffer[position] != rune('P') {
										goto l2076
									}
									position++
								}
							l2106:
								{
									add(ruleAction110, position)
								}
								add(ruleFT_P, position2105)
							}
							break
						case 'C', 'c':
							{
								position2109 := position
								{
									position2110, tokenIndex2110 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2111
									}
									position++
									goto l2110
								l2111:
									position, tokenIndex = position2110, tokenIndex2110
									if buffer[position] != rune('C') {
										goto l2076
									}
									position++
								}
							l2110:
								{
									add(ruleAction107, position)
								}
								add(ruleFT_C, position2109)
							}
							break
						case 'N', 'n':
							{
								position2113 := position
								{
									position2114, tokenIndex2114 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2115
									}
									position++
									goto l2114
								l2115:
									position, tokenIndex = position2114, tokenIndex2114
									if buffer[position] != rune('N') {
										goto l2076
									}
									position++
								}
							l2114:
								{
									position2116, tokenIndex2116 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2117
									}
									position++
									goto l2116
								l2117:
									position, tokenIndex = position2116, tokenIndex2116
									if buffer[position] != rune('C') {
										goto l2076
									}
									position++
								}
							l2116:
								{
									add(ruleAction106, position)
								}
								add(ruleFT_NC, position2113)
							}
							break
						default:
							{
								position2119 := position
								{
									position2120, tokenIndex2120 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2121
									}
									position++
									goto l2120
								l2121:
									position, tokenIndex = position2120, tokenIndex2120
									if buffer[position] != rune('Z') {
										goto l2076
									}
									position++
								}
							l2120:
								{
									add(ruleAction105, position)
								}
								add(ruleFT_Z, position2119)
							}
							break
						}
					}

				}
			l2078:
				add(rulecc, position2077)
			}
			return true
		l2076:
			position, tokenIndex = position2076, tokenIndex2076
			return false
		},
		/* 156 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action104)> */
		nil,
		/* 157 FT_Z <- <(('z' / 'Z') Action105)> */
		nil,
		/* 158 FT_NC <- <(('n' / 'N') ('c' / 'C') Action106)> */
		nil,
		/* 159 FT_C <- <(('c' / 'C') Action107)> */
		nil,
		/* 160 FT_PO <- <(('p' / 'P') ('o' / 'O') Action108)> */
		nil,
		/* 161 FT_PE <- <(('p' / 'P') ('e' / 'E') Action109)> */
		nil,
		/* 162 FT_P <- <(('p' / 'P') Action110)> */
		nil,
		/* 163 FT_M <- <(('m' / 'M') Action111)> */
		nil,
		/* 165 Action0 <- <{ p.Emit() }> */
		nil,
		nil,
		/* 167 Action1 <- <{ p.Label(buffer[begin:end])}> */
		nil,
		/* 168 Action2 <- <{ p.LD8() }> */
		nil,
		/* 169 Action3 <- <{ p.LD16() }> */
		nil,
		/* 170 Action4 <- <{ p.Push() }> */
		nil,
		/* 171 Action5 <- <{ p.Pop() }> */
		nil,
		/* 172 Action6 <- <{ p.Ex() }> */
		nil,
		/* 173 Action7 <- <{ p.Inc8() }> */
		nil,
		/* 174 Action8 <- <{ p.Inc8() }> */
		nil,
		/* 175 Action9 <- <{ p.Inc16() }> */
		nil,
		/* 176 Action10 <- <{ p.Dec8() }> */
		nil,
		/* 177 Action11 <- <{ p.Dec8() }> */
		nil,
		/* 178 Action12 <- <{ p.Dec16() }> */
		nil,
		/* 179 Action13 <- <{ p.Add16() }> */
		nil,
		/* 180 Action14 <- <{ p.Adc16() }> */
		nil,
		/* 181 Action15 <- <{ p.Sbc16() }> */
		nil,
		/* 182 Action16 <- <{ p.Dst8() }> */
		nil,
		/* 183 Action17 <- <{ p.Src8() }> */
		nil,
		/* 184 Action18 <- <{ p.Loc8() }> */
		nil,
		/* 185 Action19 <- <{ p.Copy8() }> */
		nil,
		/* 186 Action20 <- <{ p.Loc8() }> */
		nil,
		/* 187 Action21 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 188 Action22 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 189 Action23 <- <{ p.Dst16() }> */
		nil,
		/* 190 Action24 <- <{ p.Src16() }> */
		nil,
		/* 191 Action25 <- <{ p.Loc16() }> */
		nil,
		/* 192 Action26 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 193 Action27 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 194 Action28 <- <{ p.R16Contents() }> */
		nil,
		/* 195 Action29 <- <{ p.IR16Contents() }> */
		nil,
		/* 196 Action30 <- <{ p.DispDecimal(buffer[begin:end]) }> */
		nil,
		/* 197 Action31 <- <{ p.DispHex(buffer[begin:end]) }> */
		nil,
		/* 198 Action32 <- <{ p.Disp0xHex(buffer[begin:end]) }> */
		nil,
		/* 199 Action33 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 200 Action34 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 201 Action35 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 202 Action36 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 203 Action37 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 204 Action38 <- <{ p.NNContents() }> */
		nil,
		/* 205 Action39 <- <{ p.Accum("ADD") }> */
		nil,
		/* 206 Action40 <- <{ p.Accum("ADC") }> */
		nil,
		/* 207 Action41 <- <{ p.Accum("SUB") }> */
		nil,
		/* 208 Action42 <- <{ p.Accum("SBC") }> */
		nil,
		/* 209 Action43 <- <{ p.Accum("AND") }> */
		nil,
		/* 210 Action44 <- <{ p.Accum("XOR") }> */
		nil,
		/* 211 Action45 <- <{ p.Accum("OR") }> */
		nil,
		/* 212 Action46 <- <{ p.Accum("CP") }> */
		nil,
		/* 213 Action47 <- <{ p.Rot("RLC") }> */
		nil,
		/* 214 Action48 <- <{ p.Rot("RRC") }> */
		nil,
		/* 215 Action49 <- <{ p.Rot("RL") }> */
		nil,
		/* 216 Action50 <- <{ p.Rot("RR") }> */
		nil,
		/* 217 Action51 <- <{ p.Rot("SLA") }> */
		nil,
		/* 218 Action52 <- <{ p.Rot("SRA") }> */
		nil,
		/* 219 Action53 <- <{ p.Rot("SLL") }> */
		nil,
		/* 220 Action54 <- <{ p.Rot("SRL") }> */
		nil,
		/* 221 Action55 <- <{ p.Bit() }> */
		nil,
		/* 222 Action56 <- <{ p.Res() }> */
		nil,
		/* 223 Action57 <- <{ p.Set() }> */
		nil,
		/* 224 Action58 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 225 Action59 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 226 Action60 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 227 Action61 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 228 Action62 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 229 Action63 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 230 Action64 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 231 Action65 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 232 Action66 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 233 Action67 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 234 Action68 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 235 Action69 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 236 Action70 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 237 Action71 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 238 Action72 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 239 Action73 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 240 Action74 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 241 Action75 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 242 Action76 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 243 Action77 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 244 Action78 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 245 Action79 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 246 Action80 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 247 Action81 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 248 Action82 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 249 Action83 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 250 Action84 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 251 Action85 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 252 Action86 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 253 Action87 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 254 Action88 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 255 Action89 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 256 Action90 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 257 Action91 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 258 Action92 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 259 Action93 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 260 Action94 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 261 Action95 <- <{ p.Rst() }> */
		nil,
		/* 262 Action96 <- <{ p.Call() }> */
		nil,
		/* 263 Action97 <- <{ p.Ret() }> */
		nil,
		/* 264 Action98 <- <{ p.Jp() }> */
		nil,
		/* 265 Action99 <- <{ p.Jr() }> */
		nil,
		/* 266 Action100 <- <{ p.Djnz() }> */
		nil,
		/* 267 Action101 <- <{ p.In() }> */
		nil,
		/* 268 Action102 <- <{ p.Out() }> */
		nil,
		/* 269 Action103 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 270 Action104 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 271 Action105 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 272 Action106 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 273 Action107 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 274 Action108 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 275 Action109 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 276 Action110 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 277 Action111 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
