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
	ruleStatement
	ruleDirective
	ruleTitle
	ruleAseg
	ruleOrg
	ruleLabelDefn
	ruleLabelText
	rulealphanum
	rulealpha
	rulenum
	ruleComment
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
	ruleLabelNN
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
	rulePegText
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
	ruleAction112
	ruleAction113
)

var rul3s = [...]string{
	"Unknown",
	"Program",
	"Line",
	"Statement",
	"Directive",
	"Title",
	"Aseg",
	"Org",
	"LabelDefn",
	"LabelText",
	"alphanum",
	"alpha",
	"num",
	"Comment",
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
	"LabelNN",
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
	"PegText",
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
	"Action112",
	"Action113",
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
	rules  [286]func() bool
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
			p.Org()
		case ruleAction2:
			p.LabelDefn(buffer[begin:end])
		case ruleAction3:
			p.LD8()
		case ruleAction4:
			p.LD16()
		case ruleAction5:
			p.Push()
		case ruleAction6:
			p.Pop()
		case ruleAction7:
			p.Ex()
		case ruleAction8:
			p.Inc8()
		case ruleAction9:
			p.Inc8()
		case ruleAction10:
			p.Inc16()
		case ruleAction11:
			p.Dec8()
		case ruleAction12:
			p.Dec8()
		case ruleAction13:
			p.Dec16()
		case ruleAction14:
			p.Add16()
		case ruleAction15:
			p.Adc16()
		case ruleAction16:
			p.Sbc16()
		case ruleAction17:
			p.Dst8()
		case ruleAction18:
			p.Src8()
		case ruleAction19:
			p.Loc8()
		case ruleAction20:
			p.Copy8()
		case ruleAction21:
			p.Loc8()
		case ruleAction22:
			p.R8(buffer[begin:end])
		case ruleAction23:
			p.R8(buffer[begin:end])
		case ruleAction24:
			p.Dst16()
		case ruleAction25:
			p.Src16()
		case ruleAction26:
			p.Loc16()
		case ruleAction27:
			p.R16(buffer[begin:end])
		case ruleAction28:
			p.R16(buffer[begin:end])
		case ruleAction29:
			p.R16Contents()
		case ruleAction30:
			p.IR16Contents()
		case ruleAction31:
			p.DispDecimal(buffer[begin:end])
		case ruleAction32:
			p.DispHex(buffer[begin:end])
		case ruleAction33:
			p.Disp0xHex(buffer[begin:end])
		case ruleAction34:
			p.Nhex(buffer[begin:end])
		case ruleAction35:
			p.Nhex(buffer[begin:end])
		case ruleAction36:
			p.Ndec(buffer[begin:end])
		case ruleAction37:
			p.NNLabel(buffer[begin:end])
		case ruleAction38:
			p.NNhex(buffer[begin:end])
		case ruleAction39:
			p.NNhex(buffer[begin:end])
		case ruleAction40:
			p.NNContents()
		case ruleAction41:
			p.Accum("ADD")
		case ruleAction42:
			p.Accum("ADC")
		case ruleAction43:
			p.Accum("SUB")
		case ruleAction44:
			p.Accum("SBC")
		case ruleAction45:
			p.Accum("AND")
		case ruleAction46:
			p.Accum("XOR")
		case ruleAction47:
			p.Accum("OR")
		case ruleAction48:
			p.Accum("CP")
		case ruleAction49:
			p.Rot("RLC")
		case ruleAction50:
			p.Rot("RRC")
		case ruleAction51:
			p.Rot("RL")
		case ruleAction52:
			p.Rot("RR")
		case ruleAction53:
			p.Rot("SLA")
		case ruleAction54:
			p.Rot("SRA")
		case ruleAction55:
			p.Rot("SLL")
		case ruleAction56:
			p.Rot("SRL")
		case ruleAction57:
			p.Bit()
		case ruleAction58:
			p.Res()
		case ruleAction59:
			p.Set()
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
			p.Simple(buffer[begin:end])
		case ruleAction72:
			p.Simple(buffer[begin:end])
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
			p.EDSimple(buffer[begin:end])
		case ruleAction96:
			p.EDSimple(buffer[begin:end])
		case ruleAction97:
			p.Rst()
		case ruleAction98:
			p.Call()
		case ruleAction99:
			p.Ret()
		case ruleAction100:
			p.Jp()
		case ruleAction101:
			p.Jr()
		case ruleAction102:
			p.Djnz()
		case ruleAction103:
			p.In()
		case ruleAction104:
			p.Out()
		case ruleAction105:
			p.ODigit(buffer[begin:end])
		case ruleAction106:
			p.Conditional(Not{FT_Z})
		case ruleAction107:
			p.Conditional(FT_Z)
		case ruleAction108:
			p.Conditional(Not{FT_C})
		case ruleAction109:
			p.Conditional(FT_C)
		case ruleAction110:
			p.Conditional(FT_PO)
		case ruleAction111:
			p.Conditional(FT_PE)
		case ruleAction112:
			p.Conditional(FT_P)
		case ruleAction113:
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
		/* 0 Program <- <(Line+ !.)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position4 := position
				l5:
					{
						position6, tokenIndex6 := position, tokenIndex
						if !_rules[rulews]() {
							goto l6
						}
						goto l5
					l6:
						position, tokenIndex = position6, tokenIndex6
					}
					{
						position7, tokenIndex7 := position, tokenIndex
						{
							position9 := position
							if !_rules[ruleLabelText]() {
								goto l7
							}
							if buffer[position] != rune(':') {
								goto l7
							}
							position++
							if !_rules[rulews]() {
								goto l7
							}
							{
								add(ruleAction2, position)
							}
							add(ruleLabelDefn, position9)
						}
						goto l8
					l7:
						position, tokenIndex = position7, tokenIndex7
					}
				l8:
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
										switch buffer[position] {
										case 'a':
											{
												position20 := position
												if buffer[position] != rune('a') {
													goto l17
												}
												position++
												if buffer[position] != rune('s') {
													goto l17
												}
												position++
												if buffer[position] != rune('e') {
													goto l17
												}
												position++
												if buffer[position] != rune('g') {
													goto l17
												}
												position++
												add(ruleAseg, position20)
											}
											break
										case '.':
											{
												position21 := position
												if buffer[position] != rune('.') {
													goto l17
												}
												position++
												if buffer[position] != rune('t') {
													goto l17
												}
												position++
												if buffer[position] != rune('i') {
													goto l17
												}
												position++
												if buffer[position] != rune('t') {
													goto l17
												}
												position++
												if buffer[position] != rune('l') {
													goto l17
												}
												position++
												if buffer[position] != rune('e') {
													goto l17
												}
												position++
												if !_rules[rulews]() {
													goto l17
												}
												if buffer[position] != rune('\'') {
													goto l17
												}
												position++
											l22:
												{
													position23, tokenIndex23 := position, tokenIndex
													{
														position24, tokenIndex24 := position, tokenIndex
														if buffer[position] != rune('\'') {
															goto l24
														}
														position++
														goto l23
													l24:
														position, tokenIndex = position24, tokenIndex24
													}
													if !matchDot() {
														goto l23
													}
													goto l22
												l23:
													position, tokenIndex = position23, tokenIndex23
												}
												if buffer[position] != rune('\'') {
													goto l17
												}
												position++
												add(ruleTitle, position21)
											}
											break
										default:
											{
												position25 := position
												{
													position26, tokenIndex26 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l27
													}
													position++
													goto l26
												l27:
													position, tokenIndex = position26, tokenIndex26
													if buffer[position] != rune('O') {
														goto l17
													}
													position++
												}
											l26:
												{
													position28, tokenIndex28 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l29
													}
													position++
													goto l28
												l29:
													position, tokenIndex = position28, tokenIndex28
													if buffer[position] != rune('R') {
														goto l17
													}
													position++
												}
											l28:
												{
													position30, tokenIndex30 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l31
													}
													position++
													goto l30
												l31:
													position, tokenIndex = position30, tokenIndex30
													if buffer[position] != rune('G') {
														goto l17
													}
													position++
												}
											l30:
												if !_rules[rulews]() {
													goto l17
												}
												if !_rules[rulenn]() {
													goto l17
												}
												{
													add(ruleAction1, position)
												}
												add(ruleOrg, position25)
											}
											break
										}
									}

									add(ruleDirective, position18)
								}
								goto l16
							l17:
								position, tokenIndex = position16, tokenIndex16
								{
									position33 := position
									{
										position34, tokenIndex34 := position, tokenIndex
										{
											position36 := position
											{
												position37, tokenIndex37 := position, tokenIndex
												{
													position39 := position
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
															goto l38
														}
														position++
													}
												l40:
													{
														position42, tokenIndex42 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l43
														}
														position++
														goto l42
													l43:
														position, tokenIndex = position42, tokenIndex42
														if buffer[position] != rune('U') {
															goto l38
														}
														position++
													}
												l42:
													{
														position44, tokenIndex44 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l45
														}
														position++
														goto l44
													l45:
														position, tokenIndex = position44, tokenIndex44
														if buffer[position] != rune('S') {
															goto l38
														}
														position++
													}
												l44:
													{
														position46, tokenIndex46 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l47
														}
														position++
														goto l46
													l47:
														position, tokenIndex = position46, tokenIndex46
														if buffer[position] != rune('H') {
															goto l38
														}
														position++
													}
												l46:
													if !_rules[rulews]() {
														goto l38
													}
													if !_rules[ruleSrc16]() {
														goto l38
													}
													{
														add(ruleAction5, position)
													}
													add(rulePush, position39)
												}
												goto l37
											l38:
												position, tokenIndex = position37, tokenIndex37
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position50 := position
															{
																position51, tokenIndex51 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l52
																}
																position++
																goto l51
															l52:
																position, tokenIndex = position51, tokenIndex51
																if buffer[position] != rune('E') {
																	goto l35
																}
																position++
															}
														l51:
															{
																position53, tokenIndex53 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l54
																}
																position++
																goto l53
															l54:
																position, tokenIndex = position53, tokenIndex53
																if buffer[position] != rune('X') {
																	goto l35
																}
																position++
															}
														l53:
															if !_rules[rulews]() {
																goto l35
															}
															if !_rules[ruleDst16]() {
																goto l35
															}
															if !_rules[rulesep]() {
																goto l35
															}
															if !_rules[ruleSrc16]() {
																goto l35
															}
															{
																add(ruleAction7, position)
															}
															add(ruleEx, position50)
														}
														break
													case 'P', 'p':
														{
															position56 := position
															{
																position57, tokenIndex57 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l58
																}
																position++
																goto l57
															l58:
																position, tokenIndex = position57, tokenIndex57
																if buffer[position] != rune('P') {
																	goto l35
																}
																position++
															}
														l57:
															{
																position59, tokenIndex59 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l60
																}
																position++
																goto l59
															l60:
																position, tokenIndex = position59, tokenIndex59
																if buffer[position] != rune('O') {
																	goto l35
																}
																position++
															}
														l59:
															{
																position61, tokenIndex61 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l62
																}
																position++
																goto l61
															l62:
																position, tokenIndex = position61, tokenIndex61
																if buffer[position] != rune('P') {
																	goto l35
																}
																position++
															}
														l61:
															if !_rules[rulews]() {
																goto l35
															}
															if !_rules[ruleDst16]() {
																goto l35
															}
															{
																add(ruleAction6, position)
															}
															add(rulePop, position56)
														}
														break
													default:
														{
															position64 := position
															{
																position65, tokenIndex65 := position, tokenIndex
																{
																	position67 := position
																	{
																		position68, tokenIndex68 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l69
																		}
																		position++
																		goto l68
																	l69:
																		position, tokenIndex = position68, tokenIndex68
																		if buffer[position] != rune('L') {
																			goto l66
																		}
																		position++
																	}
																l68:
																	{
																		position70, tokenIndex70 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l71
																		}
																		position++
																		goto l70
																	l71:
																		position, tokenIndex = position70, tokenIndex70
																		if buffer[position] != rune('D') {
																			goto l66
																		}
																		position++
																	}
																l70:
																	if !_rules[rulews]() {
																		goto l66
																	}
																	if !_rules[ruleDst16]() {
																		goto l66
																	}
																	if !_rules[rulesep]() {
																		goto l66
																	}
																	if !_rules[ruleSrc16]() {
																		goto l66
																	}
																	{
																		add(ruleAction4, position)
																	}
																	add(ruleLoad16, position67)
																}
																goto l65
															l66:
																position, tokenIndex = position65, tokenIndex65
																{
																	position73 := position
																	{
																		position74, tokenIndex74 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l75
																		}
																		position++
																		goto l74
																	l75:
																		position, tokenIndex = position74, tokenIndex74
																		if buffer[position] != rune('L') {
																			goto l35
																		}
																		position++
																	}
																l74:
																	{
																		position76, tokenIndex76 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l77
																		}
																		position++
																		goto l76
																	l77:
																		position, tokenIndex = position76, tokenIndex76
																		if buffer[position] != rune('D') {
																			goto l35
																		}
																		position++
																	}
																l76:
																	if !_rules[rulews]() {
																		goto l35
																	}
																	{
																		position78 := position
																		{
																			position79, tokenIndex79 := position, tokenIndex
																			if !_rules[ruleReg8]() {
																				goto l80
																			}
																			goto l79
																		l80:
																			position, tokenIndex = position79, tokenIndex79
																			if !_rules[ruleReg16Contents]() {
																				goto l81
																			}
																			goto l79
																		l81:
																			position, tokenIndex = position79, tokenIndex79
																			if !_rules[rulenn_contents]() {
																				goto l35
																			}
																		}
																	l79:
																		{
																			add(ruleAction17, position)
																		}
																		add(ruleDst8, position78)
																	}
																	if !_rules[rulesep]() {
																		goto l35
																	}
																	if !_rules[ruleSrc8]() {
																		goto l35
																	}
																	{
																		add(ruleAction3, position)
																	}
																	add(ruleLoad8, position73)
																}
															}
														l65:
															add(ruleLoad, position64)
														}
														break
													}
												}

											}
										l37:
											add(ruleAssignment, position36)
										}
										goto l34
									l35:
										position, tokenIndex = position34, tokenIndex34
										{
											position85 := position
											{
												position86, tokenIndex86 := position, tokenIndex
												{
													position88 := position
													{
														position89, tokenIndex89 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l90
														}
														position++
														goto l89
													l90:
														position, tokenIndex = position89, tokenIndex89
														if buffer[position] != rune('I') {
															goto l87
														}
														position++
													}
												l89:
													{
														position91, tokenIndex91 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l92
														}
														position++
														goto l91
													l92:
														position, tokenIndex = position91, tokenIndex91
														if buffer[position] != rune('N') {
															goto l87
														}
														position++
													}
												l91:
													{
														position93, tokenIndex93 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l94
														}
														position++
														goto l93
													l94:
														position, tokenIndex = position93, tokenIndex93
														if buffer[position] != rune('C') {
															goto l87
														}
														position++
													}
												l93:
													if !_rules[rulews]() {
														goto l87
													}
													if !_rules[ruleILoc8]() {
														goto l87
													}
													{
														add(ruleAction8, position)
													}
													add(ruleInc16Indexed8, position88)
												}
												goto l86
											l87:
												position, tokenIndex = position86, tokenIndex86
												{
													position97 := position
													{
														position98, tokenIndex98 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l99
														}
														position++
														goto l98
													l99:
														position, tokenIndex = position98, tokenIndex98
														if buffer[position] != rune('I') {
															goto l96
														}
														position++
													}
												l98:
													{
														position100, tokenIndex100 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l101
														}
														position++
														goto l100
													l101:
														position, tokenIndex = position100, tokenIndex100
														if buffer[position] != rune('N') {
															goto l96
														}
														position++
													}
												l100:
													{
														position102, tokenIndex102 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l103
														}
														position++
														goto l102
													l103:
														position, tokenIndex = position102, tokenIndex102
														if buffer[position] != rune('C') {
															goto l96
														}
														position++
													}
												l102:
													if !_rules[rulews]() {
														goto l96
													}
													if !_rules[ruleLoc16]() {
														goto l96
													}
													{
														add(ruleAction10, position)
													}
													add(ruleInc16, position97)
												}
												goto l86
											l96:
												position, tokenIndex = position86, tokenIndex86
												{
													position105 := position
													{
														position106, tokenIndex106 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l107
														}
														position++
														goto l106
													l107:
														position, tokenIndex = position106, tokenIndex106
														if buffer[position] != rune('I') {
															goto l84
														}
														position++
													}
												l106:
													{
														position108, tokenIndex108 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l109
														}
														position++
														goto l108
													l109:
														position, tokenIndex = position108, tokenIndex108
														if buffer[position] != rune('N') {
															goto l84
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
															goto l84
														}
														position++
													}
												l110:
													if !_rules[rulews]() {
														goto l84
													}
													if !_rules[ruleLoc8]() {
														goto l84
													}
													{
														add(ruleAction9, position)
													}
													add(ruleInc8, position105)
												}
											}
										l86:
											add(ruleInc, position85)
										}
										goto l34
									l84:
										position, tokenIndex = position34, tokenIndex34
										{
											position114 := position
											{
												position115, tokenIndex115 := position, tokenIndex
												{
													position117 := position
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
															goto l116
														}
														position++
													}
												l118:
													{
														position120, tokenIndex120 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l121
														}
														position++
														goto l120
													l121:
														position, tokenIndex = position120, tokenIndex120
														if buffer[position] != rune('E') {
															goto l116
														}
														position++
													}
												l120:
													{
														position122, tokenIndex122 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l123
														}
														position++
														goto l122
													l123:
														position, tokenIndex = position122, tokenIndex122
														if buffer[position] != rune('C') {
															goto l116
														}
														position++
													}
												l122:
													if !_rules[rulews]() {
														goto l116
													}
													if !_rules[ruleILoc8]() {
														goto l116
													}
													{
														add(ruleAction11, position)
													}
													add(ruleDec16Indexed8, position117)
												}
												goto l115
											l116:
												position, tokenIndex = position115, tokenIndex115
												{
													position126 := position
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
															goto l125
														}
														position++
													}
												l127:
													{
														position129, tokenIndex129 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l130
														}
														position++
														goto l129
													l130:
														position, tokenIndex = position129, tokenIndex129
														if buffer[position] != rune('E') {
															goto l125
														}
														position++
													}
												l129:
													{
														position131, tokenIndex131 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l132
														}
														position++
														goto l131
													l132:
														position, tokenIndex = position131, tokenIndex131
														if buffer[position] != rune('C') {
															goto l125
														}
														position++
													}
												l131:
													if !_rules[rulews]() {
														goto l125
													}
													if !_rules[ruleLoc16]() {
														goto l125
													}
													{
														add(ruleAction13, position)
													}
													add(ruleDec16, position126)
												}
												goto l115
											l125:
												position, tokenIndex = position115, tokenIndex115
												{
													position134 := position
													{
														position135, tokenIndex135 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l136
														}
														position++
														goto l135
													l136:
														position, tokenIndex = position135, tokenIndex135
														if buffer[position] != rune('D') {
															goto l113
														}
														position++
													}
												l135:
													{
														position137, tokenIndex137 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l138
														}
														position++
														goto l137
													l138:
														position, tokenIndex = position137, tokenIndex137
														if buffer[position] != rune('E') {
															goto l113
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
															goto l113
														}
														position++
													}
												l139:
													if !_rules[rulews]() {
														goto l113
													}
													if !_rules[ruleLoc8]() {
														goto l113
													}
													{
														add(ruleAction12, position)
													}
													add(ruleDec8, position134)
												}
											}
										l115:
											add(ruleDec, position114)
										}
										goto l34
									l113:
										position, tokenIndex = position34, tokenIndex34
										{
											position143 := position
											{
												position144, tokenIndex144 := position, tokenIndex
												{
													position146 := position
													{
														position147, tokenIndex147 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l148
														}
														position++
														goto l147
													l148:
														position, tokenIndex = position147, tokenIndex147
														if buffer[position] != rune('A') {
															goto l145
														}
														position++
													}
												l147:
													{
														position149, tokenIndex149 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l150
														}
														position++
														goto l149
													l150:
														position, tokenIndex = position149, tokenIndex149
														if buffer[position] != rune('D') {
															goto l145
														}
														position++
													}
												l149:
													{
														position151, tokenIndex151 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l152
														}
														position++
														goto l151
													l152:
														position, tokenIndex = position151, tokenIndex151
														if buffer[position] != rune('D') {
															goto l145
														}
														position++
													}
												l151:
													if !_rules[rulews]() {
														goto l145
													}
													if !_rules[ruleDst16]() {
														goto l145
													}
													if !_rules[rulesep]() {
														goto l145
													}
													if !_rules[ruleSrc16]() {
														goto l145
													}
													{
														add(ruleAction14, position)
													}
													add(ruleAdd16, position146)
												}
												goto l144
											l145:
												position, tokenIndex = position144, tokenIndex144
												{
													position155 := position
													{
														position156, tokenIndex156 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l157
														}
														position++
														goto l156
													l157:
														position, tokenIndex = position156, tokenIndex156
														if buffer[position] != rune('A') {
															goto l154
														}
														position++
													}
												l156:
													{
														position158, tokenIndex158 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l159
														}
														position++
														goto l158
													l159:
														position, tokenIndex = position158, tokenIndex158
														if buffer[position] != rune('D') {
															goto l154
														}
														position++
													}
												l158:
													{
														position160, tokenIndex160 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l161
														}
														position++
														goto l160
													l161:
														position, tokenIndex = position160, tokenIndex160
														if buffer[position] != rune('C') {
															goto l154
														}
														position++
													}
												l160:
													if !_rules[rulews]() {
														goto l154
													}
													if !_rules[ruleDst16]() {
														goto l154
													}
													if !_rules[rulesep]() {
														goto l154
													}
													if !_rules[ruleSrc16]() {
														goto l154
													}
													{
														add(ruleAction15, position)
													}
													add(ruleAdc16, position155)
												}
												goto l144
											l154:
												position, tokenIndex = position144, tokenIndex144
												{
													position163 := position
													{
														position164, tokenIndex164 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l165
														}
														position++
														goto l164
													l165:
														position, tokenIndex = position164, tokenIndex164
														if buffer[position] != rune('S') {
															goto l142
														}
														position++
													}
												l164:
													{
														position166, tokenIndex166 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l167
														}
														position++
														goto l166
													l167:
														position, tokenIndex = position166, tokenIndex166
														if buffer[position] != rune('B') {
															goto l142
														}
														position++
													}
												l166:
													{
														position168, tokenIndex168 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l169
														}
														position++
														goto l168
													l169:
														position, tokenIndex = position168, tokenIndex168
														if buffer[position] != rune('C') {
															goto l142
														}
														position++
													}
												l168:
													if !_rules[rulews]() {
														goto l142
													}
													if !_rules[ruleDst16]() {
														goto l142
													}
													if !_rules[rulesep]() {
														goto l142
													}
													if !_rules[ruleSrc16]() {
														goto l142
													}
													{
														add(ruleAction16, position)
													}
													add(ruleSbc16, position163)
												}
											}
										l144:
											add(ruleAlu16, position143)
										}
										goto l34
									l142:
										position, tokenIndex = position34, tokenIndex34
										{
											position172 := position
											{
												position173, tokenIndex173 := position, tokenIndex
												{
													position175 := position
													{
														position176, tokenIndex176 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l177
														}
														position++
														goto l176
													l177:
														position, tokenIndex = position176, tokenIndex176
														if buffer[position] != rune('A') {
															goto l174
														}
														position++
													}
												l176:
													{
														position178, tokenIndex178 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l179
														}
														position++
														goto l178
													l179:
														position, tokenIndex = position178, tokenIndex178
														if buffer[position] != rune('D') {
															goto l174
														}
														position++
													}
												l178:
													{
														position180, tokenIndex180 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l181
														}
														position++
														goto l180
													l181:
														position, tokenIndex = position180, tokenIndex180
														if buffer[position] != rune('D') {
															goto l174
														}
														position++
													}
												l180:
													if !_rules[rulews]() {
														goto l174
													}
													{
														position182, tokenIndex182 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l183
														}
														position++
														goto l182
													l183:
														position, tokenIndex = position182, tokenIndex182
														if buffer[position] != rune('A') {
															goto l174
														}
														position++
													}
												l182:
													if !_rules[rulesep]() {
														goto l174
													}
													if !_rules[ruleSrc8]() {
														goto l174
													}
													{
														add(ruleAction41, position)
													}
													add(ruleAdd, position175)
												}
												goto l173
											l174:
												position, tokenIndex = position173, tokenIndex173
												{
													position186 := position
													{
														position187, tokenIndex187 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l188
														}
														position++
														goto l187
													l188:
														position, tokenIndex = position187, tokenIndex187
														if buffer[position] != rune('A') {
															goto l185
														}
														position++
													}
												l187:
													{
														position189, tokenIndex189 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l190
														}
														position++
														goto l189
													l190:
														position, tokenIndex = position189, tokenIndex189
														if buffer[position] != rune('D') {
															goto l185
														}
														position++
													}
												l189:
													{
														position191, tokenIndex191 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l192
														}
														position++
														goto l191
													l192:
														position, tokenIndex = position191, tokenIndex191
														if buffer[position] != rune('C') {
															goto l185
														}
														position++
													}
												l191:
													if !_rules[rulews]() {
														goto l185
													}
													{
														position193, tokenIndex193 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l194
														}
														position++
														goto l193
													l194:
														position, tokenIndex = position193, tokenIndex193
														if buffer[position] != rune('A') {
															goto l185
														}
														position++
													}
												l193:
													if !_rules[rulesep]() {
														goto l185
													}
													if !_rules[ruleSrc8]() {
														goto l185
													}
													{
														add(ruleAction42, position)
													}
													add(ruleAdc, position186)
												}
												goto l173
											l185:
												position, tokenIndex = position173, tokenIndex173
												{
													position197 := position
													{
														position198, tokenIndex198 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l199
														}
														position++
														goto l198
													l199:
														position, tokenIndex = position198, tokenIndex198
														if buffer[position] != rune('S') {
															goto l196
														}
														position++
													}
												l198:
													{
														position200, tokenIndex200 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l201
														}
														position++
														goto l200
													l201:
														position, tokenIndex = position200, tokenIndex200
														if buffer[position] != rune('U') {
															goto l196
														}
														position++
													}
												l200:
													{
														position202, tokenIndex202 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l203
														}
														position++
														goto l202
													l203:
														position, tokenIndex = position202, tokenIndex202
														if buffer[position] != rune('B') {
															goto l196
														}
														position++
													}
												l202:
													if !_rules[rulews]() {
														goto l196
													}
													if !_rules[ruleSrc8]() {
														goto l196
													}
													{
														add(ruleAction43, position)
													}
													add(ruleSub, position197)
												}
												goto l173
											l196:
												position, tokenIndex = position173, tokenIndex173
												{
													switch buffer[position] {
													case 'C', 'c':
														{
															position206 := position
															{
																position207, tokenIndex207 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l208
																}
																position++
																goto l207
															l208:
																position, tokenIndex = position207, tokenIndex207
																if buffer[position] != rune('C') {
																	goto l171
																}
																position++
															}
														l207:
															{
																position209, tokenIndex209 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l210
																}
																position++
																goto l209
															l210:
																position, tokenIndex = position209, tokenIndex209
																if buffer[position] != rune('P') {
																	goto l171
																}
																position++
															}
														l209:
															if !_rules[rulews]() {
																goto l171
															}
															if !_rules[ruleSrc8]() {
																goto l171
															}
															{
																add(ruleAction48, position)
															}
															add(ruleCp, position206)
														}
														break
													case 'O', 'o':
														{
															position212 := position
															{
																position213, tokenIndex213 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l214
																}
																position++
																goto l213
															l214:
																position, tokenIndex = position213, tokenIndex213
																if buffer[position] != rune('O') {
																	goto l171
																}
																position++
															}
														l213:
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
																	goto l171
																}
																position++
															}
														l215:
															if !_rules[rulews]() {
																goto l171
															}
															if !_rules[ruleSrc8]() {
																goto l171
															}
															{
																add(ruleAction47, position)
															}
															add(ruleOr, position212)
														}
														break
													case 'X', 'x':
														{
															position218 := position
															{
																position219, tokenIndex219 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l220
																}
																position++
																goto l219
															l220:
																position, tokenIndex = position219, tokenIndex219
																if buffer[position] != rune('X') {
																	goto l171
																}
																position++
															}
														l219:
															{
																position221, tokenIndex221 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l222
																}
																position++
																goto l221
															l222:
																position, tokenIndex = position221, tokenIndex221
																if buffer[position] != rune('O') {
																	goto l171
																}
																position++
															}
														l221:
															{
																position223, tokenIndex223 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l224
																}
																position++
																goto l223
															l224:
																position, tokenIndex = position223, tokenIndex223
																if buffer[position] != rune('R') {
																	goto l171
																}
																position++
															}
														l223:
															if !_rules[rulews]() {
																goto l171
															}
															if !_rules[ruleSrc8]() {
																goto l171
															}
															{
																add(ruleAction46, position)
															}
															add(ruleXor, position218)
														}
														break
													case 'A', 'a':
														{
															position226 := position
															{
																position227, tokenIndex227 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l228
																}
																position++
																goto l227
															l228:
																position, tokenIndex = position227, tokenIndex227
																if buffer[position] != rune('A') {
																	goto l171
																}
																position++
															}
														l227:
															{
																position229, tokenIndex229 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l230
																}
																position++
																goto l229
															l230:
																position, tokenIndex = position229, tokenIndex229
																if buffer[position] != rune('N') {
																	goto l171
																}
																position++
															}
														l229:
															{
																position231, tokenIndex231 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l232
																}
																position++
																goto l231
															l232:
																position, tokenIndex = position231, tokenIndex231
																if buffer[position] != rune('D') {
																	goto l171
																}
																position++
															}
														l231:
															if !_rules[rulews]() {
																goto l171
															}
															if !_rules[ruleSrc8]() {
																goto l171
															}
															{
																add(ruleAction45, position)
															}
															add(ruleAnd, position226)
														}
														break
													default:
														{
															position234 := position
															{
																position235, tokenIndex235 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l236
																}
																position++
																goto l235
															l236:
																position, tokenIndex = position235, tokenIndex235
																if buffer[position] != rune('S') {
																	goto l171
																}
																position++
															}
														l235:
															{
																position237, tokenIndex237 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l238
																}
																position++
																goto l237
															l238:
																position, tokenIndex = position237, tokenIndex237
																if buffer[position] != rune('B') {
																	goto l171
																}
																position++
															}
														l237:
															{
																position239, tokenIndex239 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l240
																}
																position++
																goto l239
															l240:
																position, tokenIndex = position239, tokenIndex239
																if buffer[position] != rune('C') {
																	goto l171
																}
																position++
															}
														l239:
															if !_rules[rulews]() {
																goto l171
															}
															{
																position241, tokenIndex241 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l242
																}
																position++
																goto l241
															l242:
																position, tokenIndex = position241, tokenIndex241
																if buffer[position] != rune('A') {
																	goto l171
																}
																position++
															}
														l241:
															if !_rules[rulesep]() {
																goto l171
															}
															if !_rules[ruleSrc8]() {
																goto l171
															}
															{
																add(ruleAction44, position)
															}
															add(ruleSbc, position234)
														}
														break
													}
												}

											}
										l173:
											add(ruleAlu, position172)
										}
										goto l34
									l171:
										position, tokenIndex = position34, tokenIndex34
										{
											position245 := position
											{
												position246, tokenIndex246 := position, tokenIndex
												{
													position248 := position
													{
														position249, tokenIndex249 := position, tokenIndex
														{
															position251 := position
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
																	goto l250
																}
																position++
															}
														l252:
															{
																position254, tokenIndex254 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l255
																}
																position++
																goto l254
															l255:
																position, tokenIndex = position254, tokenIndex254
																if buffer[position] != rune('L') {
																	goto l250
																}
																position++
															}
														l254:
															{
																position256, tokenIndex256 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l257
																}
																position++
																goto l256
															l257:
																position, tokenIndex = position256, tokenIndex256
																if buffer[position] != rune('C') {
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
																position258, tokenIndex258 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l258
																}
																if !_rules[ruleCopy8]() {
																	goto l258
																}
																goto l259
															l258:
																position, tokenIndex = position258, tokenIndex258
															}
														l259:
															{
																add(ruleAction49, position)
															}
															add(ruleRlc, position251)
														}
														goto l249
													l250:
														position, tokenIndex = position249, tokenIndex249
														{
															position262 := position
															{
																position263, tokenIndex263 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l264
																}
																position++
																goto l263
															l264:
																position, tokenIndex = position263, tokenIndex263
																if buffer[position] != rune('R') {
																	goto l261
																}
																position++
															}
														l263:
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
																	goto l261
																}
																position++
															}
														l265:
															{
																position267, tokenIndex267 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l268
																}
																position++
																goto l267
															l268:
																position, tokenIndex = position267, tokenIndex267
																if buffer[position] != rune('C') {
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
																position269, tokenIndex269 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l269
																}
																if !_rules[ruleCopy8]() {
																	goto l269
																}
																goto l270
															l269:
																position, tokenIndex = position269, tokenIndex269
															}
														l270:
															{
																add(ruleAction50, position)
															}
															add(ruleRrc, position262)
														}
														goto l249
													l261:
														position, tokenIndex = position249, tokenIndex249
														{
															position273 := position
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
																	goto l272
																}
																position++
															}
														l274:
															{
																position276, tokenIndex276 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l277
																}
																position++
																goto l276
															l277:
																position, tokenIndex = position276, tokenIndex276
																if buffer[position] != rune('L') {
																	goto l272
																}
																position++
															}
														l276:
															if !_rules[rulews]() {
																goto l272
															}
															if !_rules[ruleLoc8]() {
																goto l272
															}
															{
																position278, tokenIndex278 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l278
																}
																if !_rules[ruleCopy8]() {
																	goto l278
																}
																goto l279
															l278:
																position, tokenIndex = position278, tokenIndex278
															}
														l279:
															{
																add(ruleAction51, position)
															}
															add(ruleRl, position273)
														}
														goto l249
													l272:
														position, tokenIndex = position249, tokenIndex249
														{
															position282 := position
															{
																position283, tokenIndex283 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l284
																}
																position++
																goto l283
															l284:
																position, tokenIndex = position283, tokenIndex283
																if buffer[position] != rune('R') {
																	goto l281
																}
																position++
															}
														l283:
															{
																position285, tokenIndex285 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l286
																}
																position++
																goto l285
															l286:
																position, tokenIndex = position285, tokenIndex285
																if buffer[position] != rune('R') {
																	goto l281
																}
																position++
															}
														l285:
															if !_rules[rulews]() {
																goto l281
															}
															if !_rules[ruleLoc8]() {
																goto l281
															}
															{
																position287, tokenIndex287 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l287
																}
																if !_rules[ruleCopy8]() {
																	goto l287
																}
																goto l288
															l287:
																position, tokenIndex = position287, tokenIndex287
															}
														l288:
															{
																add(ruleAction52, position)
															}
															add(ruleRr, position282)
														}
														goto l249
													l281:
														position, tokenIndex = position249, tokenIndex249
														{
															position291 := position
															{
																position292, tokenIndex292 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l293
																}
																position++
																goto l292
															l293:
																position, tokenIndex = position292, tokenIndex292
																if buffer[position] != rune('S') {
																	goto l290
																}
																position++
															}
														l292:
															{
																position294, tokenIndex294 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l295
																}
																position++
																goto l294
															l295:
																position, tokenIndex = position294, tokenIndex294
																if buffer[position] != rune('L') {
																	goto l290
																}
																position++
															}
														l294:
															{
																position296, tokenIndex296 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l297
																}
																position++
																goto l296
															l297:
																position, tokenIndex = position296, tokenIndex296
																if buffer[position] != rune('A') {
																	goto l290
																}
																position++
															}
														l296:
															if !_rules[rulews]() {
																goto l290
															}
															if !_rules[ruleLoc8]() {
																goto l290
															}
															{
																position298, tokenIndex298 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l298
																}
																if !_rules[ruleCopy8]() {
																	goto l298
																}
																goto l299
															l298:
																position, tokenIndex = position298, tokenIndex298
															}
														l299:
															{
																add(ruleAction53, position)
															}
															add(ruleSla, position291)
														}
														goto l249
													l290:
														position, tokenIndex = position249, tokenIndex249
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
																	goto l301
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
																	goto l301
																}
																position++
															}
														l305:
															{
																position307, tokenIndex307 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l308
																}
																position++
																goto l307
															l308:
																position, tokenIndex = position307, tokenIndex307
																if buffer[position] != rune('A') {
																	goto l301
																}
																position++
															}
														l307:
															if !_rules[rulews]() {
																goto l301
															}
															if !_rules[ruleLoc8]() {
																goto l301
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
																add(ruleAction54, position)
															}
															add(ruleSra, position302)
														}
														goto l249
													l301:
														position, tokenIndex = position249, tokenIndex249
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
																	goto l312
																}
																position++
															}
														l314:
															{
																position316, tokenIndex316 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l317
																}
																position++
																goto l316
															l317:
																position, tokenIndex = position316, tokenIndex316
																if buffer[position] != rune('L') {
																	goto l312
																}
																position++
															}
														l316:
															{
																position318, tokenIndex318 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l319
																}
																position++
																goto l318
															l319:
																position, tokenIndex = position318, tokenIndex318
																if buffer[position] != rune('L') {
																	goto l312
																}
																position++
															}
														l318:
															if !_rules[rulews]() {
																goto l312
															}
															if !_rules[ruleLoc8]() {
																goto l312
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
																add(ruleAction55, position)
															}
															add(ruleSll, position313)
														}
														goto l249
													l312:
														position, tokenIndex = position249, tokenIndex249
														{
															position323 := position
															{
																position324, tokenIndex324 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l325
																}
																position++
																goto l324
															l325:
																position, tokenIndex = position324, tokenIndex324
																if buffer[position] != rune('S') {
																	goto l247
																}
																position++
															}
														l324:
															{
																position326, tokenIndex326 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l327
																}
																position++
																goto l326
															l327:
																position, tokenIndex = position326, tokenIndex326
																if buffer[position] != rune('R') {
																	goto l247
																}
																position++
															}
														l326:
															{
																position328, tokenIndex328 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l329
																}
																position++
																goto l328
															l329:
																position, tokenIndex = position328, tokenIndex328
																if buffer[position] != rune('L') {
																	goto l247
																}
																position++
															}
														l328:
															if !_rules[rulews]() {
																goto l247
															}
															if !_rules[ruleLoc8]() {
																goto l247
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
																add(ruleAction56, position)
															}
															add(ruleSrl, position323)
														}
													}
												l249:
													add(ruleRot, position248)
												}
												goto l246
											l247:
												position, tokenIndex = position246, tokenIndex246
												{
													switch buffer[position] {
													case 'S', 's':
														{
															position334 := position
															{
																position335, tokenIndex335 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l336
																}
																position++
																goto l335
															l336:
																position, tokenIndex = position335, tokenIndex335
																if buffer[position] != rune('S') {
																	goto l244
																}
																position++
															}
														l335:
															{
																position337, tokenIndex337 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l338
																}
																position++
																goto l337
															l338:
																position, tokenIndex = position337, tokenIndex337
																if buffer[position] != rune('E') {
																	goto l244
																}
																position++
															}
														l337:
															{
																position339, tokenIndex339 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l340
																}
																position++
																goto l339
															l340:
																position, tokenIndex = position339, tokenIndex339
																if buffer[position] != rune('T') {
																	goto l244
																}
																position++
															}
														l339:
															if !_rules[rulews]() {
																goto l244
															}
															if !_rules[ruleoctaldigit]() {
																goto l244
															}
															if !_rules[rulesep]() {
																goto l244
															}
															if !_rules[ruleLoc8]() {
																goto l244
															}
															{
																position341, tokenIndex341 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l341
																}
																if !_rules[ruleCopy8]() {
																	goto l341
																}
																goto l342
															l341:
																position, tokenIndex = position341, tokenIndex341
															}
														l342:
															{
																add(ruleAction59, position)
															}
															add(ruleSet, position334)
														}
														break
													case 'R', 'r':
														{
															position344 := position
															{
																position345, tokenIndex345 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l346
																}
																position++
																goto l345
															l346:
																position, tokenIndex = position345, tokenIndex345
																if buffer[position] != rune('R') {
																	goto l244
																}
																position++
															}
														l345:
															{
																position347, tokenIndex347 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l348
																}
																position++
																goto l347
															l348:
																position, tokenIndex = position347, tokenIndex347
																if buffer[position] != rune('E') {
																	goto l244
																}
																position++
															}
														l347:
															{
																position349, tokenIndex349 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l350
																}
																position++
																goto l349
															l350:
																position, tokenIndex = position349, tokenIndex349
																if buffer[position] != rune('S') {
																	goto l244
																}
																position++
															}
														l349:
															if !_rules[rulews]() {
																goto l244
															}
															if !_rules[ruleoctaldigit]() {
																goto l244
															}
															if !_rules[rulesep]() {
																goto l244
															}
															if !_rules[ruleLoc8]() {
																goto l244
															}
															{
																position351, tokenIndex351 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l351
																}
																if !_rules[ruleCopy8]() {
																	goto l351
																}
																goto l352
															l351:
																position, tokenIndex = position351, tokenIndex351
															}
														l352:
															{
																add(ruleAction58, position)
															}
															add(ruleRes, position344)
														}
														break
													default:
														{
															position354 := position
															{
																position355, tokenIndex355 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l356
																}
																position++
																goto l355
															l356:
																position, tokenIndex = position355, tokenIndex355
																if buffer[position] != rune('B') {
																	goto l244
																}
																position++
															}
														l355:
															{
																position357, tokenIndex357 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l358
																}
																position++
																goto l357
															l358:
																position, tokenIndex = position357, tokenIndex357
																if buffer[position] != rune('I') {
																	goto l244
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
																	goto l244
																}
																position++
															}
														l359:
															if !_rules[rulews]() {
																goto l244
															}
															if !_rules[ruleoctaldigit]() {
																goto l244
															}
															if !_rules[rulesep]() {
																goto l244
															}
															if !_rules[ruleLoc8]() {
																goto l244
															}
															{
																add(ruleAction57, position)
															}
															add(ruleBit, position354)
														}
														break
													}
												}

											}
										l246:
											add(ruleBitOp, position245)
										}
										goto l34
									l244:
										position, tokenIndex = position34, tokenIndex34
										{
											position363 := position
											{
												position364, tokenIndex364 := position, tokenIndex
												{
													position366 := position
													{
														position367 := position
														{
															position368, tokenIndex368 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l369
															}
															position++
															goto l368
														l369:
															position, tokenIndex = position368, tokenIndex368
															if buffer[position] != rune('R') {
																goto l365
															}
															position++
														}
													l368:
														{
															position370, tokenIndex370 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l371
															}
															position++
															goto l370
														l371:
															position, tokenIndex = position370, tokenIndex370
															if buffer[position] != rune('E') {
																goto l365
															}
															position++
														}
													l370:
														{
															position372, tokenIndex372 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l373
															}
															position++
															goto l372
														l373:
															position, tokenIndex = position372, tokenIndex372
															if buffer[position] != rune('T') {
																goto l365
															}
															position++
														}
													l372:
														{
															position374, tokenIndex374 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l375
															}
															position++
															goto l374
														l375:
															position, tokenIndex = position374, tokenIndex374
															if buffer[position] != rune('N') {
																goto l365
															}
															position++
														}
													l374:
														add(rulePegText, position367)
													}
													{
														add(ruleAction74, position)
													}
													add(ruleRetn, position366)
												}
												goto l364
											l365:
												position, tokenIndex = position364, tokenIndex364
												{
													position378 := position
													{
														position379 := position
														{
															position380, tokenIndex380 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l381
															}
															position++
															goto l380
														l381:
															position, tokenIndex = position380, tokenIndex380
															if buffer[position] != rune('R') {
																goto l377
															}
															position++
														}
													l380:
														{
															position382, tokenIndex382 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l383
															}
															position++
															goto l382
														l383:
															position, tokenIndex = position382, tokenIndex382
															if buffer[position] != rune('E') {
																goto l377
															}
															position++
														}
													l382:
														{
															position384, tokenIndex384 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l385
															}
															position++
															goto l384
														l385:
															position, tokenIndex = position384, tokenIndex384
															if buffer[position] != rune('T') {
																goto l377
															}
															position++
														}
													l384:
														{
															position386, tokenIndex386 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l387
															}
															position++
															goto l386
														l387:
															position, tokenIndex = position386, tokenIndex386
															if buffer[position] != rune('I') {
																goto l377
															}
															position++
														}
													l386:
														add(rulePegText, position379)
													}
													{
														add(ruleAction75, position)
													}
													add(ruleReti, position378)
												}
												goto l364
											l377:
												position, tokenIndex = position364, tokenIndex364
												{
													position390 := position
													{
														position391 := position
														{
															position392, tokenIndex392 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l393
															}
															position++
															goto l392
														l393:
															position, tokenIndex = position392, tokenIndex392
															if buffer[position] != rune('R') {
																goto l389
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
																goto l389
															}
															position++
														}
													l394:
														{
															position396, tokenIndex396 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l397
															}
															position++
															goto l396
														l397:
															position, tokenIndex = position396, tokenIndex396
															if buffer[position] != rune('D') {
																goto l389
															}
															position++
														}
													l396:
														add(rulePegText, position391)
													}
													{
														add(ruleAction76, position)
													}
													add(ruleRrd, position390)
												}
												goto l364
											l389:
												position, tokenIndex = position364, tokenIndex364
												{
													position400 := position
													{
														position401 := position
														{
															position402, tokenIndex402 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l403
															}
															position++
															goto l402
														l403:
															position, tokenIndex = position402, tokenIndex402
															if buffer[position] != rune('I') {
																goto l399
															}
															position++
														}
													l402:
														{
															position404, tokenIndex404 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l405
															}
															position++
															goto l404
														l405:
															position, tokenIndex = position404, tokenIndex404
															if buffer[position] != rune('M') {
																goto l399
															}
															position++
														}
													l404:
														if buffer[position] != rune(' ') {
															goto l399
														}
														position++
														if buffer[position] != rune('0') {
															goto l399
														}
														position++
														add(rulePegText, position401)
													}
													{
														add(ruleAction78, position)
													}
													add(ruleIm0, position400)
												}
												goto l364
											l399:
												position, tokenIndex = position364, tokenIndex364
												{
													position408 := position
													{
														position409 := position
														{
															position410, tokenIndex410 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l411
															}
															position++
															goto l410
														l411:
															position, tokenIndex = position410, tokenIndex410
															if buffer[position] != rune('I') {
																goto l407
															}
															position++
														}
													l410:
														{
															position412, tokenIndex412 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l413
															}
															position++
															goto l412
														l413:
															position, tokenIndex = position412, tokenIndex412
															if buffer[position] != rune('M') {
																goto l407
															}
															position++
														}
													l412:
														if buffer[position] != rune(' ') {
															goto l407
														}
														position++
														if buffer[position] != rune('1') {
															goto l407
														}
														position++
														add(rulePegText, position409)
													}
													{
														add(ruleAction79, position)
													}
													add(ruleIm1, position408)
												}
												goto l364
											l407:
												position, tokenIndex = position364, tokenIndex364
												{
													position416 := position
													{
														position417 := position
														{
															position418, tokenIndex418 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l419
															}
															position++
															goto l418
														l419:
															position, tokenIndex = position418, tokenIndex418
															if buffer[position] != rune('I') {
																goto l415
															}
															position++
														}
													l418:
														{
															position420, tokenIndex420 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l421
															}
															position++
															goto l420
														l421:
															position, tokenIndex = position420, tokenIndex420
															if buffer[position] != rune('M') {
																goto l415
															}
															position++
														}
													l420:
														if buffer[position] != rune(' ') {
															goto l415
														}
														position++
														if buffer[position] != rune('2') {
															goto l415
														}
														position++
														add(rulePegText, position417)
													}
													{
														add(ruleAction80, position)
													}
													add(ruleIm2, position416)
												}
												goto l364
											l415:
												position, tokenIndex = position364, tokenIndex364
												{
													switch buffer[position] {
													case 'I', 'O', 'i', 'o':
														{
															position424 := position
															{
																position425, tokenIndex425 := position, tokenIndex
																{
																	position427 := position
																	{
																		position428 := position
																		{
																			position429, tokenIndex429 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l430
																			}
																			position++
																			goto l429
																		l430:
																			position, tokenIndex = position429, tokenIndex429
																			if buffer[position] != rune('I') {
																				goto l426
																			}
																			position++
																		}
																	l429:
																		{
																			position431, tokenIndex431 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l432
																			}
																			position++
																			goto l431
																		l432:
																			position, tokenIndex = position431, tokenIndex431
																			if buffer[position] != rune('N') {
																				goto l426
																			}
																			position++
																		}
																	l431:
																		{
																			position433, tokenIndex433 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l434
																			}
																			position++
																			goto l433
																		l434:
																			position, tokenIndex = position433, tokenIndex433
																			if buffer[position] != rune('I') {
																				goto l426
																			}
																			position++
																		}
																	l433:
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
																				goto l426
																			}
																			position++
																		}
																	l435:
																		add(rulePegText, position428)
																	}
																	{
																		add(ruleAction91, position)
																	}
																	add(ruleInir, position427)
																}
																goto l425
															l426:
																position, tokenIndex = position425, tokenIndex425
																{
																	position439 := position
																	{
																		position440 := position
																		{
																			position441, tokenIndex441 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l442
																			}
																			position++
																			goto l441
																		l442:
																			position, tokenIndex = position441, tokenIndex441
																			if buffer[position] != rune('I') {
																				goto l438
																			}
																			position++
																		}
																	l441:
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
																				goto l438
																			}
																			position++
																		}
																	l443:
																		{
																			position445, tokenIndex445 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l446
																			}
																			position++
																			goto l445
																		l446:
																			position, tokenIndex = position445, tokenIndex445
																			if buffer[position] != rune('I') {
																				goto l438
																			}
																			position++
																		}
																	l445:
																		add(rulePegText, position440)
																	}
																	{
																		add(ruleAction83, position)
																	}
																	add(ruleIni, position439)
																}
																goto l425
															l438:
																position, tokenIndex = position425, tokenIndex425
																{
																	position449 := position
																	{
																		position450 := position
																		{
																			position451, tokenIndex451 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l452
																			}
																			position++
																			goto l451
																		l452:
																			position, tokenIndex = position451, tokenIndex451
																			if buffer[position] != rune('O') {
																				goto l448
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
																				goto l448
																			}
																			position++
																		}
																	l453:
																		{
																			position455, tokenIndex455 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l456
																			}
																			position++
																			goto l455
																		l456:
																			position, tokenIndex = position455, tokenIndex455
																			if buffer[position] != rune('I') {
																				goto l448
																			}
																			position++
																		}
																	l455:
																		{
																			position457, tokenIndex457 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l458
																			}
																			position++
																			goto l457
																		l458:
																			position, tokenIndex = position457, tokenIndex457
																			if buffer[position] != rune('R') {
																				goto l448
																			}
																			position++
																		}
																	l457:
																		add(rulePegText, position450)
																	}
																	{
																		add(ruleAction92, position)
																	}
																	add(ruleOtir, position449)
																}
																goto l425
															l448:
																position, tokenIndex = position425, tokenIndex425
																{
																	position461 := position
																	{
																		position462 := position
																		{
																			position463, tokenIndex463 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l464
																			}
																			position++
																			goto l463
																		l464:
																			position, tokenIndex = position463, tokenIndex463
																			if buffer[position] != rune('O') {
																				goto l460
																			}
																			position++
																		}
																	l463:
																		{
																			position465, tokenIndex465 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l466
																			}
																			position++
																			goto l465
																		l466:
																			position, tokenIndex = position465, tokenIndex465
																			if buffer[position] != rune('U') {
																				goto l460
																			}
																			position++
																		}
																	l465:
																		{
																			position467, tokenIndex467 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l468
																			}
																			position++
																			goto l467
																		l468:
																			position, tokenIndex = position467, tokenIndex467
																			if buffer[position] != rune('T') {
																				goto l460
																			}
																			position++
																		}
																	l467:
																		{
																			position469, tokenIndex469 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l470
																			}
																			position++
																			goto l469
																		l470:
																			position, tokenIndex = position469, tokenIndex469
																			if buffer[position] != rune('I') {
																				goto l460
																			}
																			position++
																		}
																	l469:
																		add(rulePegText, position462)
																	}
																	{
																		add(ruleAction84, position)
																	}
																	add(ruleOuti, position461)
																}
																goto l425
															l460:
																position, tokenIndex = position425, tokenIndex425
																{
																	position473 := position
																	{
																		position474 := position
																		{
																			position475, tokenIndex475 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l476
																			}
																			position++
																			goto l475
																		l476:
																			position, tokenIndex = position475, tokenIndex475
																			if buffer[position] != rune('I') {
																				goto l472
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
																				goto l472
																			}
																			position++
																		}
																	l477:
																		{
																			position479, tokenIndex479 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l480
																			}
																			position++
																			goto l479
																		l480:
																			position, tokenIndex = position479, tokenIndex479
																			if buffer[position] != rune('D') {
																				goto l472
																			}
																			position++
																		}
																	l479:
																		{
																			position481, tokenIndex481 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l482
																			}
																			position++
																			goto l481
																		l482:
																			position, tokenIndex = position481, tokenIndex481
																			if buffer[position] != rune('R') {
																				goto l472
																			}
																			position++
																		}
																	l481:
																		add(rulePegText, position474)
																	}
																	{
																		add(ruleAction95, position)
																	}
																	add(ruleIndr, position473)
																}
																goto l425
															l472:
																position, tokenIndex = position425, tokenIndex425
																{
																	position485 := position
																	{
																		position486 := position
																		{
																			position487, tokenIndex487 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l488
																			}
																			position++
																			goto l487
																		l488:
																			position, tokenIndex = position487, tokenIndex487
																			if buffer[position] != rune('I') {
																				goto l484
																			}
																			position++
																		}
																	l487:
																		{
																			position489, tokenIndex489 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l490
																			}
																			position++
																			goto l489
																		l490:
																			position, tokenIndex = position489, tokenIndex489
																			if buffer[position] != rune('N') {
																				goto l484
																			}
																			position++
																		}
																	l489:
																		{
																			position491, tokenIndex491 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l492
																			}
																			position++
																			goto l491
																		l492:
																			position, tokenIndex = position491, tokenIndex491
																			if buffer[position] != rune('D') {
																				goto l484
																			}
																			position++
																		}
																	l491:
																		add(rulePegText, position486)
																	}
																	{
																		add(ruleAction87, position)
																	}
																	add(ruleInd, position485)
																}
																goto l425
															l484:
																position, tokenIndex = position425, tokenIndex425
																{
																	position495 := position
																	{
																		position496 := position
																		{
																			position497, tokenIndex497 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l498
																			}
																			position++
																			goto l497
																		l498:
																			position, tokenIndex = position497, tokenIndex497
																			if buffer[position] != rune('O') {
																				goto l494
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
																				goto l494
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
																				goto l494
																			}
																			position++
																		}
																	l501:
																		{
																			position503, tokenIndex503 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l504
																			}
																			position++
																			goto l503
																		l504:
																			position, tokenIndex = position503, tokenIndex503
																			if buffer[position] != rune('R') {
																				goto l494
																			}
																			position++
																		}
																	l503:
																		add(rulePegText, position496)
																	}
																	{
																		add(ruleAction96, position)
																	}
																	add(ruleOtdr, position495)
																}
																goto l425
															l494:
																position, tokenIndex = position425, tokenIndex425
																{
																	position506 := position
																	{
																		position507 := position
																		{
																			position508, tokenIndex508 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l509
																			}
																			position++
																			goto l508
																		l509:
																			position, tokenIndex = position508, tokenIndex508
																			if buffer[position] != rune('O') {
																				goto l362
																			}
																			position++
																		}
																	l508:
																		{
																			position510, tokenIndex510 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l511
																			}
																			position++
																			goto l510
																		l511:
																			position, tokenIndex = position510, tokenIndex510
																			if buffer[position] != rune('U') {
																				goto l362
																			}
																			position++
																		}
																	l510:
																		{
																			position512, tokenIndex512 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l513
																			}
																			position++
																			goto l512
																		l513:
																			position, tokenIndex = position512, tokenIndex512
																			if buffer[position] != rune('T') {
																				goto l362
																			}
																			position++
																		}
																	l512:
																		{
																			position514, tokenIndex514 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l515
																			}
																			position++
																			goto l514
																		l515:
																			position, tokenIndex = position514, tokenIndex514
																			if buffer[position] != rune('D') {
																				goto l362
																			}
																			position++
																		}
																	l514:
																		add(rulePegText, position507)
																	}
																	{
																		add(ruleAction88, position)
																	}
																	add(ruleOutd, position506)
																}
															}
														l425:
															add(ruleBlitIO, position424)
														}
														break
													case 'R', 'r':
														{
															position517 := position
															{
																position518 := position
																{
																	position519, tokenIndex519 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l520
																	}
																	position++
																	goto l519
																l520:
																	position, tokenIndex = position519, tokenIndex519
																	if buffer[position] != rune('R') {
																		goto l362
																	}
																	position++
																}
															l519:
																{
																	position521, tokenIndex521 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l522
																	}
																	position++
																	goto l521
																l522:
																	position, tokenIndex = position521, tokenIndex521
																	if buffer[position] != rune('L') {
																		goto l362
																	}
																	position++
																}
															l521:
																{
																	position523, tokenIndex523 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l524
																	}
																	position++
																	goto l523
																l524:
																	position, tokenIndex = position523, tokenIndex523
																	if buffer[position] != rune('D') {
																		goto l362
																	}
																	position++
																}
															l523:
																add(rulePegText, position518)
															}
															{
																add(ruleAction77, position)
															}
															add(ruleRld, position517)
														}
														break
													case 'N', 'n':
														{
															position526 := position
															{
																position527 := position
																{
																	position528, tokenIndex528 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l529
																	}
																	position++
																	goto l528
																l529:
																	position, tokenIndex = position528, tokenIndex528
																	if buffer[position] != rune('N') {
																		goto l362
																	}
																	position++
																}
															l528:
																{
																	position530, tokenIndex530 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l531
																	}
																	position++
																	goto l530
																l531:
																	position, tokenIndex = position530, tokenIndex530
																	if buffer[position] != rune('E') {
																		goto l362
																	}
																	position++
																}
															l530:
																{
																	position532, tokenIndex532 := position, tokenIndex
																	if buffer[position] != rune('g') {
																		goto l533
																	}
																	position++
																	goto l532
																l533:
																	position, tokenIndex = position532, tokenIndex532
																	if buffer[position] != rune('G') {
																		goto l362
																	}
																	position++
																}
															l532:
																add(rulePegText, position527)
															}
															{
																add(ruleAction73, position)
															}
															add(ruleNeg, position526)
														}
														break
													default:
														{
															position535 := position
															{
																position536, tokenIndex536 := position, tokenIndex
																{
																	position538 := position
																	{
																		position539 := position
																		{
																			position540, tokenIndex540 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l541
																			}
																			position++
																			goto l540
																		l541:
																			position, tokenIndex = position540, tokenIndex540
																			if buffer[position] != rune('L') {
																				goto l537
																			}
																			position++
																		}
																	l540:
																		{
																			position542, tokenIndex542 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l543
																			}
																			position++
																			goto l542
																		l543:
																			position, tokenIndex = position542, tokenIndex542
																			if buffer[position] != rune('D') {
																				goto l537
																			}
																			position++
																		}
																	l542:
																		{
																			position544, tokenIndex544 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l545
																			}
																			position++
																			goto l544
																		l545:
																			position, tokenIndex = position544, tokenIndex544
																			if buffer[position] != rune('I') {
																				goto l537
																			}
																			position++
																		}
																	l544:
																		{
																			position546, tokenIndex546 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l547
																			}
																			position++
																			goto l546
																		l547:
																			position, tokenIndex = position546, tokenIndex546
																			if buffer[position] != rune('R') {
																				goto l537
																			}
																			position++
																		}
																	l546:
																		add(rulePegText, position539)
																	}
																	{
																		add(ruleAction89, position)
																	}
																	add(ruleLdir, position538)
																}
																goto l536
															l537:
																position, tokenIndex = position536, tokenIndex536
																{
																	position550 := position
																	{
																		position551 := position
																		{
																			position552, tokenIndex552 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l553
																			}
																			position++
																			goto l552
																		l553:
																			position, tokenIndex = position552, tokenIndex552
																			if buffer[position] != rune('L') {
																				goto l549
																			}
																			position++
																		}
																	l552:
																		{
																			position554, tokenIndex554 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l555
																			}
																			position++
																			goto l554
																		l555:
																			position, tokenIndex = position554, tokenIndex554
																			if buffer[position] != rune('D') {
																				goto l549
																			}
																			position++
																		}
																	l554:
																		{
																			position556, tokenIndex556 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l557
																			}
																			position++
																			goto l556
																		l557:
																			position, tokenIndex = position556, tokenIndex556
																			if buffer[position] != rune('I') {
																				goto l549
																			}
																			position++
																		}
																	l556:
																		add(rulePegText, position551)
																	}
																	{
																		add(ruleAction81, position)
																	}
																	add(ruleLdi, position550)
																}
																goto l536
															l549:
																position, tokenIndex = position536, tokenIndex536
																{
																	position560 := position
																	{
																		position561 := position
																		{
																			position562, tokenIndex562 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l563
																			}
																			position++
																			goto l562
																		l563:
																			position, tokenIndex = position562, tokenIndex562
																			if buffer[position] != rune('C') {
																				goto l559
																			}
																			position++
																		}
																	l562:
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
																				goto l559
																			}
																			position++
																		}
																	l564:
																		{
																			position566, tokenIndex566 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l567
																			}
																			position++
																			goto l566
																		l567:
																			position, tokenIndex = position566, tokenIndex566
																			if buffer[position] != rune('I') {
																				goto l559
																			}
																			position++
																		}
																	l566:
																		{
																			position568, tokenIndex568 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l569
																			}
																			position++
																			goto l568
																		l569:
																			position, tokenIndex = position568, tokenIndex568
																			if buffer[position] != rune('R') {
																				goto l559
																			}
																			position++
																		}
																	l568:
																		add(rulePegText, position561)
																	}
																	{
																		add(ruleAction90, position)
																	}
																	add(ruleCpir, position560)
																}
																goto l536
															l559:
																position, tokenIndex = position536, tokenIndex536
																{
																	position572 := position
																	{
																		position573 := position
																		{
																			position574, tokenIndex574 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l575
																			}
																			position++
																			goto l574
																		l575:
																			position, tokenIndex = position574, tokenIndex574
																			if buffer[position] != rune('C') {
																				goto l571
																			}
																			position++
																		}
																	l574:
																		{
																			position576, tokenIndex576 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l577
																			}
																			position++
																			goto l576
																		l577:
																			position, tokenIndex = position576, tokenIndex576
																			if buffer[position] != rune('P') {
																				goto l571
																			}
																			position++
																		}
																	l576:
																		{
																			position578, tokenIndex578 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l579
																			}
																			position++
																			goto l578
																		l579:
																			position, tokenIndex = position578, tokenIndex578
																			if buffer[position] != rune('I') {
																				goto l571
																			}
																			position++
																		}
																	l578:
																		add(rulePegText, position573)
																	}
																	{
																		add(ruleAction82, position)
																	}
																	add(ruleCpi, position572)
																}
																goto l536
															l571:
																position, tokenIndex = position536, tokenIndex536
																{
																	position582 := position
																	{
																		position583 := position
																		{
																			position584, tokenIndex584 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l585
																			}
																			position++
																			goto l584
																		l585:
																			position, tokenIndex = position584, tokenIndex584
																			if buffer[position] != rune('L') {
																				goto l581
																			}
																			position++
																		}
																	l584:
																		{
																			position586, tokenIndex586 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l587
																			}
																			position++
																			goto l586
																		l587:
																			position, tokenIndex = position586, tokenIndex586
																			if buffer[position] != rune('D') {
																				goto l581
																			}
																			position++
																		}
																	l586:
																		{
																			position588, tokenIndex588 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l589
																			}
																			position++
																			goto l588
																		l589:
																			position, tokenIndex = position588, tokenIndex588
																			if buffer[position] != rune('D') {
																				goto l581
																			}
																			position++
																		}
																	l588:
																		{
																			position590, tokenIndex590 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l591
																			}
																			position++
																			goto l590
																		l591:
																			position, tokenIndex = position590, tokenIndex590
																			if buffer[position] != rune('R') {
																				goto l581
																			}
																			position++
																		}
																	l590:
																		add(rulePegText, position583)
																	}
																	{
																		add(ruleAction93, position)
																	}
																	add(ruleLddr, position582)
																}
																goto l536
															l581:
																position, tokenIndex = position536, tokenIndex536
																{
																	position594 := position
																	{
																		position595 := position
																		{
																			position596, tokenIndex596 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l597
																			}
																			position++
																			goto l596
																		l597:
																			position, tokenIndex = position596, tokenIndex596
																			if buffer[position] != rune('L') {
																				goto l593
																			}
																			position++
																		}
																	l596:
																		{
																			position598, tokenIndex598 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l599
																			}
																			position++
																			goto l598
																		l599:
																			position, tokenIndex = position598, tokenIndex598
																			if buffer[position] != rune('D') {
																				goto l593
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
																				goto l593
																			}
																			position++
																		}
																	l600:
																		add(rulePegText, position595)
																	}
																	{
																		add(ruleAction85, position)
																	}
																	add(ruleLdd, position594)
																}
																goto l536
															l593:
																position, tokenIndex = position536, tokenIndex536
																{
																	position604 := position
																	{
																		position605 := position
																		{
																			position606, tokenIndex606 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l607
																			}
																			position++
																			goto l606
																		l607:
																			position, tokenIndex = position606, tokenIndex606
																			if buffer[position] != rune('C') {
																				goto l603
																			}
																			position++
																		}
																	l606:
																		{
																			position608, tokenIndex608 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l609
																			}
																			position++
																			goto l608
																		l609:
																			position, tokenIndex = position608, tokenIndex608
																			if buffer[position] != rune('P') {
																				goto l603
																			}
																			position++
																		}
																	l608:
																		{
																			position610, tokenIndex610 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l611
																			}
																			position++
																			goto l610
																		l611:
																			position, tokenIndex = position610, tokenIndex610
																			if buffer[position] != rune('D') {
																				goto l603
																			}
																			position++
																		}
																	l610:
																		{
																			position612, tokenIndex612 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l613
																			}
																			position++
																			goto l612
																		l613:
																			position, tokenIndex = position612, tokenIndex612
																			if buffer[position] != rune('R') {
																				goto l603
																			}
																			position++
																		}
																	l612:
																		add(rulePegText, position605)
																	}
																	{
																		add(ruleAction94, position)
																	}
																	add(ruleCpdr, position604)
																}
																goto l536
															l603:
																position, tokenIndex = position536, tokenIndex536
																{
																	position615 := position
																	{
																		position616 := position
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
																				goto l362
																			}
																			position++
																		}
																	l617:
																		{
																			position619, tokenIndex619 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l620
																			}
																			position++
																			goto l619
																		l620:
																			position, tokenIndex = position619, tokenIndex619
																			if buffer[position] != rune('P') {
																				goto l362
																			}
																			position++
																		}
																	l619:
																		{
																			position621, tokenIndex621 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l622
																			}
																			position++
																			goto l621
																		l622:
																			position, tokenIndex = position621, tokenIndex621
																			if buffer[position] != rune('D') {
																				goto l362
																			}
																			position++
																		}
																	l621:
																		add(rulePegText, position616)
																	}
																	{
																		add(ruleAction86, position)
																	}
																	add(ruleCpd, position615)
																}
															}
														l536:
															add(ruleBlit, position535)
														}
														break
													}
												}

											}
										l364:
											add(ruleEDSimple, position363)
										}
										goto l34
									l362:
										position, tokenIndex = position34, tokenIndex34
										{
											position625 := position
											{
												position626, tokenIndex626 := position, tokenIndex
												{
													position628 := position
													{
														position629 := position
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
																goto l627
															}
															position++
														}
													l630:
														{
															position632, tokenIndex632 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l633
															}
															position++
															goto l632
														l633:
															position, tokenIndex = position632, tokenIndex632
															if buffer[position] != rune('L') {
																goto l627
															}
															position++
														}
													l632:
														{
															position634, tokenIndex634 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l635
															}
															position++
															goto l634
														l635:
															position, tokenIndex = position634, tokenIndex634
															if buffer[position] != rune('C') {
																goto l627
															}
															position++
														}
													l634:
														{
															position636, tokenIndex636 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l637
															}
															position++
															goto l636
														l637:
															position, tokenIndex = position636, tokenIndex636
															if buffer[position] != rune('A') {
																goto l627
															}
															position++
														}
													l636:
														add(rulePegText, position629)
													}
													{
														add(ruleAction62, position)
													}
													add(ruleRlca, position628)
												}
												goto l626
											l627:
												position, tokenIndex = position626, tokenIndex626
												{
													position640 := position
													{
														position641 := position
														{
															position642, tokenIndex642 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l643
															}
															position++
															goto l642
														l643:
															position, tokenIndex = position642, tokenIndex642
															if buffer[position] != rune('R') {
																goto l639
															}
															position++
														}
													l642:
														{
															position644, tokenIndex644 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l645
															}
															position++
															goto l644
														l645:
															position, tokenIndex = position644, tokenIndex644
															if buffer[position] != rune('R') {
																goto l639
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
																goto l639
															}
															position++
														}
													l646:
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
																goto l639
															}
															position++
														}
													l648:
														add(rulePegText, position641)
													}
													{
														add(ruleAction63, position)
													}
													add(ruleRrca, position640)
												}
												goto l626
											l639:
												position, tokenIndex = position626, tokenIndex626
												{
													position652 := position
													{
														position653 := position
														{
															position654, tokenIndex654 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l655
															}
															position++
															goto l654
														l655:
															position, tokenIndex = position654, tokenIndex654
															if buffer[position] != rune('R') {
																goto l651
															}
															position++
														}
													l654:
														{
															position656, tokenIndex656 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l657
															}
															position++
															goto l656
														l657:
															position, tokenIndex = position656, tokenIndex656
															if buffer[position] != rune('L') {
																goto l651
															}
															position++
														}
													l656:
														{
															position658, tokenIndex658 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l659
															}
															position++
															goto l658
														l659:
															position, tokenIndex = position658, tokenIndex658
															if buffer[position] != rune('A') {
																goto l651
															}
															position++
														}
													l658:
														add(rulePegText, position653)
													}
													{
														add(ruleAction64, position)
													}
													add(ruleRla, position652)
												}
												goto l626
											l651:
												position, tokenIndex = position626, tokenIndex626
												{
													position662 := position
													{
														position663 := position
														{
															position664, tokenIndex664 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l665
															}
															position++
															goto l664
														l665:
															position, tokenIndex = position664, tokenIndex664
															if buffer[position] != rune('D') {
																goto l661
															}
															position++
														}
													l664:
														{
															position666, tokenIndex666 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l667
															}
															position++
															goto l666
														l667:
															position, tokenIndex = position666, tokenIndex666
															if buffer[position] != rune('A') {
																goto l661
															}
															position++
														}
													l666:
														{
															position668, tokenIndex668 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l669
															}
															position++
															goto l668
														l669:
															position, tokenIndex = position668, tokenIndex668
															if buffer[position] != rune('A') {
																goto l661
															}
															position++
														}
													l668:
														add(rulePegText, position663)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleDaa, position662)
												}
												goto l626
											l661:
												position, tokenIndex = position626, tokenIndex626
												{
													position672 := position
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
																goto l671
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
																goto l671
															}
															position++
														}
													l676:
														{
															position678, tokenIndex678 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l679
															}
															position++
															goto l678
														l679:
															position, tokenIndex = position678, tokenIndex678
															if buffer[position] != rune('L') {
																goto l671
															}
															position++
														}
													l678:
														add(rulePegText, position673)
													}
													{
														add(ruleAction67, position)
													}
													add(ruleCpl, position672)
												}
												goto l626
											l671:
												position, tokenIndex = position626, tokenIndex626
												{
													position682 := position
													{
														position683 := position
														{
															position684, tokenIndex684 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l685
															}
															position++
															goto l684
														l685:
															position, tokenIndex = position684, tokenIndex684
															if buffer[position] != rune('E') {
																goto l681
															}
															position++
														}
													l684:
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
																goto l681
															}
															position++
														}
													l686:
														{
															position688, tokenIndex688 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l689
															}
															position++
															goto l688
														l689:
															position, tokenIndex = position688, tokenIndex688
															if buffer[position] != rune('X') {
																goto l681
															}
															position++
														}
													l688:
														add(rulePegText, position683)
													}
													{
														add(ruleAction70, position)
													}
													add(ruleExx, position682)
												}
												goto l626
											l681:
												position, tokenIndex = position626, tokenIndex626
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position692 := position
															{
																position693 := position
																{
																	position694, tokenIndex694 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l695
																	}
																	position++
																	goto l694
																l695:
																	position, tokenIndex = position694, tokenIndex694
																	if buffer[position] != rune('E') {
																		goto l624
																	}
																	position++
																}
															l694:
																{
																	position696, tokenIndex696 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l697
																	}
																	position++
																	goto l696
																l697:
																	position, tokenIndex = position696, tokenIndex696
																	if buffer[position] != rune('I') {
																		goto l624
																	}
																	position++
																}
															l696:
																add(rulePegText, position693)
															}
															{
																add(ruleAction72, position)
															}
															add(ruleEi, position692)
														}
														break
													case 'D', 'd':
														{
															position699 := position
															{
																position700 := position
																{
																	position701, tokenIndex701 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l702
																	}
																	position++
																	goto l701
																l702:
																	position, tokenIndex = position701, tokenIndex701
																	if buffer[position] != rune('D') {
																		goto l624
																	}
																	position++
																}
															l701:
																{
																	position703, tokenIndex703 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l704
																	}
																	position++
																	goto l703
																l704:
																	position, tokenIndex = position703, tokenIndex703
																	if buffer[position] != rune('I') {
																		goto l624
																	}
																	position++
																}
															l703:
																add(rulePegText, position700)
															}
															{
																add(ruleAction71, position)
															}
															add(ruleDi, position699)
														}
														break
													case 'C', 'c':
														{
															position706 := position
															{
																position707 := position
																{
																	position708, tokenIndex708 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l709
																	}
																	position++
																	goto l708
																l709:
																	position, tokenIndex = position708, tokenIndex708
																	if buffer[position] != rune('C') {
																		goto l624
																	}
																	position++
																}
															l708:
																{
																	position710, tokenIndex710 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l711
																	}
																	position++
																	goto l710
																l711:
																	position, tokenIndex = position710, tokenIndex710
																	if buffer[position] != rune('C') {
																		goto l624
																	}
																	position++
																}
															l710:
																{
																	position712, tokenIndex712 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l713
																	}
																	position++
																	goto l712
																l713:
																	position, tokenIndex = position712, tokenIndex712
																	if buffer[position] != rune('F') {
																		goto l624
																	}
																	position++
																}
															l712:
																add(rulePegText, position707)
															}
															{
																add(ruleAction69, position)
															}
															add(ruleCcf, position706)
														}
														break
													case 'S', 's':
														{
															position715 := position
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
																		goto l624
																	}
																	position++
																}
															l717:
																{
																	position719, tokenIndex719 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l720
																	}
																	position++
																	goto l719
																l720:
																	position, tokenIndex = position719, tokenIndex719
																	if buffer[position] != rune('C') {
																		goto l624
																	}
																	position++
																}
															l719:
																{
																	position721, tokenIndex721 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l722
																	}
																	position++
																	goto l721
																l722:
																	position, tokenIndex = position721, tokenIndex721
																	if buffer[position] != rune('F') {
																		goto l624
																	}
																	position++
																}
															l721:
																add(rulePegText, position716)
															}
															{
																add(ruleAction68, position)
															}
															add(ruleScf, position715)
														}
														break
													case 'R', 'r':
														{
															position724 := position
															{
																position725 := position
																{
																	position726, tokenIndex726 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l727
																	}
																	position++
																	goto l726
																l727:
																	position, tokenIndex = position726, tokenIndex726
																	if buffer[position] != rune('R') {
																		goto l624
																	}
																	position++
																}
															l726:
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
																		goto l624
																	}
																	position++
																}
															l728:
																{
																	position730, tokenIndex730 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l731
																	}
																	position++
																	goto l730
																l731:
																	position, tokenIndex = position730, tokenIndex730
																	if buffer[position] != rune('A') {
																		goto l624
																	}
																	position++
																}
															l730:
																add(rulePegText, position725)
															}
															{
																add(ruleAction65, position)
															}
															add(ruleRra, position724)
														}
														break
													case 'H', 'h':
														{
															position733 := position
															{
																position734 := position
																{
																	position735, tokenIndex735 := position, tokenIndex
																	if buffer[position] != rune('h') {
																		goto l736
																	}
																	position++
																	goto l735
																l736:
																	position, tokenIndex = position735, tokenIndex735
																	if buffer[position] != rune('H') {
																		goto l624
																	}
																	position++
																}
															l735:
																{
																	position737, tokenIndex737 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l738
																	}
																	position++
																	goto l737
																l738:
																	position, tokenIndex = position737, tokenIndex737
																	if buffer[position] != rune('A') {
																		goto l624
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
																		goto l624
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
																		goto l624
																	}
																	position++
																}
															l741:
																add(rulePegText, position734)
															}
															{
																add(ruleAction61, position)
															}
															add(ruleHalt, position733)
														}
														break
													default:
														{
															position744 := position
															{
																position745 := position
																{
																	position746, tokenIndex746 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l747
																	}
																	position++
																	goto l746
																l747:
																	position, tokenIndex = position746, tokenIndex746
																	if buffer[position] != rune('N') {
																		goto l624
																	}
																	position++
																}
															l746:
																{
																	position748, tokenIndex748 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l749
																	}
																	position++
																	goto l748
																l749:
																	position, tokenIndex = position748, tokenIndex748
																	if buffer[position] != rune('O') {
																		goto l624
																	}
																	position++
																}
															l748:
																{
																	position750, tokenIndex750 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l751
																	}
																	position++
																	goto l750
																l751:
																	position, tokenIndex = position750, tokenIndex750
																	if buffer[position] != rune('P') {
																		goto l624
																	}
																	position++
																}
															l750:
																add(rulePegText, position745)
															}
															{
																add(ruleAction60, position)
															}
															add(ruleNop, position744)
														}
														break
													}
												}

											}
										l626:
											add(ruleSimple, position625)
										}
										goto l34
									l624:
										position, tokenIndex = position34, tokenIndex34
										{
											position754 := position
											{
												position755, tokenIndex755 := position, tokenIndex
												{
													position757 := position
													{
														position758, tokenIndex758 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l759
														}
														position++
														goto l758
													l759:
														position, tokenIndex = position758, tokenIndex758
														if buffer[position] != rune('R') {
															goto l756
														}
														position++
													}
												l758:
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
															goto l756
														}
														position++
													}
												l760:
													{
														position762, tokenIndex762 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l763
														}
														position++
														goto l762
													l763:
														position, tokenIndex = position762, tokenIndex762
														if buffer[position] != rune('T') {
															goto l756
														}
														position++
													}
												l762:
													if !_rules[rulews]() {
														goto l756
													}
													if !_rules[rulen]() {
														goto l756
													}
													{
														add(ruleAction97, position)
													}
													add(ruleRst, position757)
												}
												goto l755
											l756:
												position, tokenIndex = position755, tokenIndex755
												{
													position766 := position
													{
														position767, tokenIndex767 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l768
														}
														position++
														goto l767
													l768:
														position, tokenIndex = position767, tokenIndex767
														if buffer[position] != rune('J') {
															goto l765
														}
														position++
													}
												l767:
													{
														position769, tokenIndex769 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l770
														}
														position++
														goto l769
													l770:
														position, tokenIndex = position769, tokenIndex769
														if buffer[position] != rune('P') {
															goto l765
														}
														position++
													}
												l769:
													if !_rules[rulews]() {
														goto l765
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
														goto l765
													}
													{
														add(ruleAction100, position)
													}
													add(ruleJp, position766)
												}
												goto l755
											l765:
												position, tokenIndex = position755, tokenIndex755
												{
													switch buffer[position] {
													case 'D', 'd':
														{
															position775 := position
															{
																position776, tokenIndex776 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l777
																}
																position++
																goto l776
															l777:
																position, tokenIndex = position776, tokenIndex776
																if buffer[position] != rune('D') {
																	goto l753
																}
																position++
															}
														l776:
															{
																position778, tokenIndex778 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l779
																}
																position++
																goto l778
															l779:
																position, tokenIndex = position778, tokenIndex778
																if buffer[position] != rune('J') {
																	goto l753
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
																	goto l753
																}
																position++
															}
														l780:
															{
																position782, tokenIndex782 := position, tokenIndex
																if buffer[position] != rune('z') {
																	goto l783
																}
																position++
																goto l782
															l783:
																position, tokenIndex = position782, tokenIndex782
																if buffer[position] != rune('Z') {
																	goto l753
																}
																position++
															}
														l782:
															if !_rules[rulews]() {
																goto l753
															}
															if !_rules[ruledisp]() {
																goto l753
															}
															{
																add(ruleAction102, position)
															}
															add(ruleDjnz, position775)
														}
														break
													case 'J', 'j':
														{
															position785 := position
															{
																position786, tokenIndex786 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l787
																}
																position++
																goto l786
															l787:
																position, tokenIndex = position786, tokenIndex786
																if buffer[position] != rune('J') {
																	goto l753
																}
																position++
															}
														l786:
															{
																position788, tokenIndex788 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l789
																}
																position++
																goto l788
															l789:
																position, tokenIndex = position788, tokenIndex788
																if buffer[position] != rune('R') {
																	goto l753
																}
																position++
															}
														l788:
															if !_rules[rulews]() {
																goto l753
															}
															{
																position790, tokenIndex790 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l790
																}
																if !_rules[rulesep]() {
																	goto l790
																}
																goto l791
															l790:
																position, tokenIndex = position790, tokenIndex790
															}
														l791:
															if !_rules[ruledisp]() {
																goto l753
															}
															{
																add(ruleAction101, position)
															}
															add(ruleJr, position785)
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
																	goto l753
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
																	goto l753
																}
																position++
															}
														l796:
															{
																position798, tokenIndex798 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l799
																}
																position++
																goto l798
															l799:
																position, tokenIndex = position798, tokenIndex798
																if buffer[position] != rune('T') {
																	goto l753
																}
																position++
															}
														l798:
															{
																position800, tokenIndex800 := position, tokenIndex
																if !_rules[rulews]() {
																	goto l800
																}
																if !_rules[rulecc]() {
																	goto l800
																}
																goto l801
															l800:
																position, tokenIndex = position800, tokenIndex800
															}
														l801:
															{
																add(ruleAction99, position)
															}
															add(ruleRet, position793)
														}
														break
													default:
														{
															position803 := position
															{
																position804, tokenIndex804 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l805
																}
																position++
																goto l804
															l805:
																position, tokenIndex = position804, tokenIndex804
																if buffer[position] != rune('C') {
																	goto l753
																}
																position++
															}
														l804:
															{
																position806, tokenIndex806 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l807
																}
																position++
																goto l806
															l807:
																position, tokenIndex = position806, tokenIndex806
																if buffer[position] != rune('A') {
																	goto l753
																}
																position++
															}
														l806:
															{
																position808, tokenIndex808 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l809
																}
																position++
																goto l808
															l809:
																position, tokenIndex = position808, tokenIndex808
																if buffer[position] != rune('L') {
																	goto l753
																}
																position++
															}
														l808:
															{
																position810, tokenIndex810 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l811
																}
																position++
																goto l810
															l811:
																position, tokenIndex = position810, tokenIndex810
																if buffer[position] != rune('L') {
																	goto l753
																}
																position++
															}
														l810:
															if !_rules[rulews]() {
																goto l753
															}
															{
																position812, tokenIndex812 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l812
																}
																if !_rules[rulesep]() {
																	goto l812
																}
																goto l813
															l812:
																position, tokenIndex = position812, tokenIndex812
															}
														l813:
															if !_rules[ruleSrc16]() {
																goto l753
															}
															{
																add(ruleAction98, position)
															}
															add(ruleCall, position803)
														}
														break
													}
												}

											}
										l755:
											add(ruleJump, position754)
										}
										goto l34
									l753:
										position, tokenIndex = position34, tokenIndex34
										{
											position815 := position
											{
												position816, tokenIndex816 := position, tokenIndex
												{
													position818 := position
													{
														position819, tokenIndex819 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l820
														}
														position++
														goto l819
													l820:
														position, tokenIndex = position819, tokenIndex819
														if buffer[position] != rune('I') {
															goto l817
														}
														position++
													}
												l819:
													{
														position821, tokenIndex821 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l822
														}
														position++
														goto l821
													l822:
														position, tokenIndex = position821, tokenIndex821
														if buffer[position] != rune('N') {
															goto l817
														}
														position++
													}
												l821:
													if !_rules[rulews]() {
														goto l817
													}
													if !_rules[ruleReg8]() {
														goto l817
													}
													if !_rules[rulesep]() {
														goto l817
													}
													if !_rules[rulePort]() {
														goto l817
													}
													{
														add(ruleAction103, position)
													}
													add(ruleIN, position818)
												}
												goto l816
											l817:
												position, tokenIndex = position816, tokenIndex816
												{
													position824 := position
													{
														position825, tokenIndex825 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l826
														}
														position++
														goto l825
													l826:
														position, tokenIndex = position825, tokenIndex825
														if buffer[position] != rune('O') {
															goto l13
														}
														position++
													}
												l825:
													{
														position827, tokenIndex827 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l828
														}
														position++
														goto l827
													l828:
														position, tokenIndex = position827, tokenIndex827
														if buffer[position] != rune('U') {
															goto l13
														}
														position++
													}
												l827:
													{
														position829, tokenIndex829 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l830
														}
														position++
														goto l829
													l830:
														position, tokenIndex = position829, tokenIndex829
														if buffer[position] != rune('T') {
															goto l13
														}
														position++
													}
												l829:
													if !_rules[rulews]() {
														goto l13
													}
													if !_rules[rulePort]() {
														goto l13
													}
													if !_rules[rulesep]() {
														goto l13
													}
													if !_rules[ruleReg8]() {
														goto l13
													}
													{
														add(ruleAction104, position)
													}
													add(ruleOUT, position824)
												}
											}
										l816:
											add(ruleIO, position815)
										}
									}
								l34:
									add(ruleInstruction, position33)
								}
							}
						l16:
							add(ruleStatement, position15)
						}
						goto l14
					l13:
						position, tokenIndex = position13, tokenIndex13
					}
				l14:
					{
						position832, tokenIndex832 := position, tokenIndex
						if !_rules[rulews]() {
							goto l832
						}
						goto l833
					l832:
						position, tokenIndex = position832, tokenIndex832
					}
				l833:
					{
						position834, tokenIndex834 := position, tokenIndex
						{
							position836 := position
							{
								position837, tokenIndex837 := position, tokenIndex
								if buffer[position] != rune(';') {
									goto l838
								}
								position++
								goto l837
							l838:
								position, tokenIndex = position837, tokenIndex837
								if buffer[position] != rune('#') {
									goto l834
								}
								position++
							}
						l837:
						l839:
							{
								position840, tokenIndex840 := position, tokenIndex
								{
									position841, tokenIndex841 := position, tokenIndex
									if buffer[position] != rune('\n') {
										goto l841
									}
									position++
									goto l840
								l841:
									position, tokenIndex = position841, tokenIndex841
								}
								if !matchDot() {
									goto l840
								}
								goto l839
							l840:
								position, tokenIndex = position840, tokenIndex840
							}
							add(ruleComment, position836)
						}
						goto l835
					l834:
						position, tokenIndex = position834, tokenIndex834
					}
				l835:
					{
						position842, tokenIndex842 := position, tokenIndex
						if !_rules[rulews]() {
							goto l842
						}
						goto l843
					l842:
						position, tokenIndex = position842, tokenIndex842
					}
				l843:
					{
						position844, tokenIndex844 := position, tokenIndex
						{
							position846, tokenIndex846 := position, tokenIndex
							if buffer[position] != rune('\r') {
								goto l846
							}
							position++
							goto l847
						l846:
							position, tokenIndex = position846, tokenIndex846
						}
					l847:
						if buffer[position] != rune('\n') {
							goto l845
						}
						position++
						goto l844
					l845:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune(':') {
							goto l0
						}
						position++
					}
				l844:
					{
						add(ruleAction0, position)
					}
					add(ruleLine, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position849 := position
					l850:
						{
							position851, tokenIndex851 := position, tokenIndex
							if !_rules[rulews]() {
								goto l851
							}
							goto l850
						l851:
							position, tokenIndex = position851, tokenIndex851
						}
						{
							position852, tokenIndex852 := position, tokenIndex
							{
								position854 := position
								if !_rules[ruleLabelText]() {
									goto l852
								}
								if buffer[position] != rune(':') {
									goto l852
								}
								position++
								if !_rules[rulews]() {
									goto l852
								}
								{
									add(ruleAction2, position)
								}
								add(ruleLabelDefn, position854)
							}
							goto l853
						l852:
							position, tokenIndex = position852, tokenIndex852
						}
					l853:
					l856:
						{
							position857, tokenIndex857 := position, tokenIndex
							if !_rules[rulews]() {
								goto l857
							}
							goto l856
						l857:
							position, tokenIndex = position857, tokenIndex857
						}
						{
							position858, tokenIndex858 := position, tokenIndex
							{
								position860 := position
								{
									position861, tokenIndex861 := position, tokenIndex
									{
										position863 := position
										{
											switch buffer[position] {
											case 'a':
												{
													position865 := position
													if buffer[position] != rune('a') {
														goto l862
													}
													position++
													if buffer[position] != rune('s') {
														goto l862
													}
													position++
													if buffer[position] != rune('e') {
														goto l862
													}
													position++
													if buffer[position] != rune('g') {
														goto l862
													}
													position++
													add(ruleAseg, position865)
												}
												break
											case '.':
												{
													position866 := position
													if buffer[position] != rune('.') {
														goto l862
													}
													position++
													if buffer[position] != rune('t') {
														goto l862
													}
													position++
													if buffer[position] != rune('i') {
														goto l862
													}
													position++
													if buffer[position] != rune('t') {
														goto l862
													}
													position++
													if buffer[position] != rune('l') {
														goto l862
													}
													position++
													if buffer[position] != rune('e') {
														goto l862
													}
													position++
													if !_rules[rulews]() {
														goto l862
													}
													if buffer[position] != rune('\'') {
														goto l862
													}
													position++
												l867:
													{
														position868, tokenIndex868 := position, tokenIndex
														{
															position869, tokenIndex869 := position, tokenIndex
															if buffer[position] != rune('\'') {
																goto l869
															}
															position++
															goto l868
														l869:
															position, tokenIndex = position869, tokenIndex869
														}
														if !matchDot() {
															goto l868
														}
														goto l867
													l868:
														position, tokenIndex = position868, tokenIndex868
													}
													if buffer[position] != rune('\'') {
														goto l862
													}
													position++
													add(ruleTitle, position866)
												}
												break
											default:
												{
													position870 := position
													{
														position871, tokenIndex871 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l872
														}
														position++
														goto l871
													l872:
														position, tokenIndex = position871, tokenIndex871
														if buffer[position] != rune('O') {
															goto l862
														}
														position++
													}
												l871:
													{
														position873, tokenIndex873 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l874
														}
														position++
														goto l873
													l874:
														position, tokenIndex = position873, tokenIndex873
														if buffer[position] != rune('R') {
															goto l862
														}
														position++
													}
												l873:
													{
														position875, tokenIndex875 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l876
														}
														position++
														goto l875
													l876:
														position, tokenIndex = position875, tokenIndex875
														if buffer[position] != rune('G') {
															goto l862
														}
														position++
													}
												l875:
													if !_rules[rulews]() {
														goto l862
													}
													if !_rules[rulenn]() {
														goto l862
													}
													{
														add(ruleAction1, position)
													}
													add(ruleOrg, position870)
												}
												break
											}
										}

										add(ruleDirective, position863)
									}
									goto l861
								l862:
									position, tokenIndex = position861, tokenIndex861
									{
										position878 := position
										{
											position879, tokenIndex879 := position, tokenIndex
											{
												position881 := position
												{
													position882, tokenIndex882 := position, tokenIndex
													{
														position884 := position
														{
															position885, tokenIndex885 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l886
															}
															position++
															goto l885
														l886:
															position, tokenIndex = position885, tokenIndex885
															if buffer[position] != rune('P') {
																goto l883
															}
															position++
														}
													l885:
														{
															position887, tokenIndex887 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l888
															}
															position++
															goto l887
														l888:
															position, tokenIndex = position887, tokenIndex887
															if buffer[position] != rune('U') {
																goto l883
															}
															position++
														}
													l887:
														{
															position889, tokenIndex889 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l890
															}
															position++
															goto l889
														l890:
															position, tokenIndex = position889, tokenIndex889
															if buffer[position] != rune('S') {
																goto l883
															}
															position++
														}
													l889:
														{
															position891, tokenIndex891 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l892
															}
															position++
															goto l891
														l892:
															position, tokenIndex = position891, tokenIndex891
															if buffer[position] != rune('H') {
																goto l883
															}
															position++
														}
													l891:
														if !_rules[rulews]() {
															goto l883
														}
														if !_rules[ruleSrc16]() {
															goto l883
														}
														{
															add(ruleAction5, position)
														}
														add(rulePush, position884)
													}
													goto l882
												l883:
													position, tokenIndex = position882, tokenIndex882
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position895 := position
																{
																	position896, tokenIndex896 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l897
																	}
																	position++
																	goto l896
																l897:
																	position, tokenIndex = position896, tokenIndex896
																	if buffer[position] != rune('E') {
																		goto l880
																	}
																	position++
																}
															l896:
																{
																	position898, tokenIndex898 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l899
																	}
																	position++
																	goto l898
																l899:
																	position, tokenIndex = position898, tokenIndex898
																	if buffer[position] != rune('X') {
																		goto l880
																	}
																	position++
																}
															l898:
																if !_rules[rulews]() {
																	goto l880
																}
																if !_rules[ruleDst16]() {
																	goto l880
																}
																if !_rules[rulesep]() {
																	goto l880
																}
																if !_rules[ruleSrc16]() {
																	goto l880
																}
																{
																	add(ruleAction7, position)
																}
																add(ruleEx, position895)
															}
															break
														case 'P', 'p':
															{
																position901 := position
																{
																	position902, tokenIndex902 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l903
																	}
																	position++
																	goto l902
																l903:
																	position, tokenIndex = position902, tokenIndex902
																	if buffer[position] != rune('P') {
																		goto l880
																	}
																	position++
																}
															l902:
																{
																	position904, tokenIndex904 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l905
																	}
																	position++
																	goto l904
																l905:
																	position, tokenIndex = position904, tokenIndex904
																	if buffer[position] != rune('O') {
																		goto l880
																	}
																	position++
																}
															l904:
																{
																	position906, tokenIndex906 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l907
																	}
																	position++
																	goto l906
																l907:
																	position, tokenIndex = position906, tokenIndex906
																	if buffer[position] != rune('P') {
																		goto l880
																	}
																	position++
																}
															l906:
																if !_rules[rulews]() {
																	goto l880
																}
																if !_rules[ruleDst16]() {
																	goto l880
																}
																{
																	add(ruleAction6, position)
																}
																add(rulePop, position901)
															}
															break
														default:
															{
																position909 := position
																{
																	position910, tokenIndex910 := position, tokenIndex
																	{
																		position912 := position
																		{
																			position913, tokenIndex913 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l914
																			}
																			position++
																			goto l913
																		l914:
																			position, tokenIndex = position913, tokenIndex913
																			if buffer[position] != rune('L') {
																				goto l911
																			}
																			position++
																		}
																	l913:
																		{
																			position915, tokenIndex915 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l916
																			}
																			position++
																			goto l915
																		l916:
																			position, tokenIndex = position915, tokenIndex915
																			if buffer[position] != rune('D') {
																				goto l911
																			}
																			position++
																		}
																	l915:
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
																			add(ruleAction4, position)
																		}
																		add(ruleLoad16, position912)
																	}
																	goto l910
																l911:
																	position, tokenIndex = position910, tokenIndex910
																	{
																		position918 := position
																		{
																			position919, tokenIndex919 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l920
																			}
																			position++
																			goto l919
																		l920:
																			position, tokenIndex = position919, tokenIndex919
																			if buffer[position] != rune('L') {
																				goto l880
																			}
																			position++
																		}
																	l919:
																		{
																			position921, tokenIndex921 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l922
																			}
																			position++
																			goto l921
																		l922:
																			position, tokenIndex = position921, tokenIndex921
																			if buffer[position] != rune('D') {
																				goto l880
																			}
																			position++
																		}
																	l921:
																		if !_rules[rulews]() {
																			goto l880
																		}
																		{
																			position923 := position
																			{
																				position924, tokenIndex924 := position, tokenIndex
																				if !_rules[ruleReg8]() {
																					goto l925
																				}
																				goto l924
																			l925:
																				position, tokenIndex = position924, tokenIndex924
																				if !_rules[ruleReg16Contents]() {
																					goto l926
																				}
																				goto l924
																			l926:
																				position, tokenIndex = position924, tokenIndex924
																				if !_rules[rulenn_contents]() {
																					goto l880
																				}
																			}
																		l924:
																			{
																				add(ruleAction17, position)
																			}
																			add(ruleDst8, position923)
																		}
																		if !_rules[rulesep]() {
																			goto l880
																		}
																		if !_rules[ruleSrc8]() {
																			goto l880
																		}
																		{
																			add(ruleAction3, position)
																		}
																		add(ruleLoad8, position918)
																	}
																}
															l910:
																add(ruleLoad, position909)
															}
															break
														}
													}

												}
											l882:
												add(ruleAssignment, position881)
											}
											goto l879
										l880:
											position, tokenIndex = position879, tokenIndex879
											{
												position930 := position
												{
													position931, tokenIndex931 := position, tokenIndex
													{
														position933 := position
														{
															position934, tokenIndex934 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l935
															}
															position++
															goto l934
														l935:
															position, tokenIndex = position934, tokenIndex934
															if buffer[position] != rune('I') {
																goto l932
															}
															position++
														}
													l934:
														{
															position936, tokenIndex936 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l937
															}
															position++
															goto l936
														l937:
															position, tokenIndex = position936, tokenIndex936
															if buffer[position] != rune('N') {
																goto l932
															}
															position++
														}
													l936:
														{
															position938, tokenIndex938 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l939
															}
															position++
															goto l938
														l939:
															position, tokenIndex = position938, tokenIndex938
															if buffer[position] != rune('C') {
																goto l932
															}
															position++
														}
													l938:
														if !_rules[rulews]() {
															goto l932
														}
														if !_rules[ruleILoc8]() {
															goto l932
														}
														{
															add(ruleAction8, position)
														}
														add(ruleInc16Indexed8, position933)
													}
													goto l931
												l932:
													position, tokenIndex = position931, tokenIndex931
													{
														position942 := position
														{
															position943, tokenIndex943 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l944
															}
															position++
															goto l943
														l944:
															position, tokenIndex = position943, tokenIndex943
															if buffer[position] != rune('I') {
																goto l941
															}
															position++
														}
													l943:
														{
															position945, tokenIndex945 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l946
															}
															position++
															goto l945
														l946:
															position, tokenIndex = position945, tokenIndex945
															if buffer[position] != rune('N') {
																goto l941
															}
															position++
														}
													l945:
														{
															position947, tokenIndex947 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l948
															}
															position++
															goto l947
														l948:
															position, tokenIndex = position947, tokenIndex947
															if buffer[position] != rune('C') {
																goto l941
															}
															position++
														}
													l947:
														if !_rules[rulews]() {
															goto l941
														}
														if !_rules[ruleLoc16]() {
															goto l941
														}
														{
															add(ruleAction10, position)
														}
														add(ruleInc16, position942)
													}
													goto l931
												l941:
													position, tokenIndex = position931, tokenIndex931
													{
														position950 := position
														{
															position951, tokenIndex951 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l952
															}
															position++
															goto l951
														l952:
															position, tokenIndex = position951, tokenIndex951
															if buffer[position] != rune('I') {
																goto l929
															}
															position++
														}
													l951:
														{
															position953, tokenIndex953 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l954
															}
															position++
															goto l953
														l954:
															position, tokenIndex = position953, tokenIndex953
															if buffer[position] != rune('N') {
																goto l929
															}
															position++
														}
													l953:
														{
															position955, tokenIndex955 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l956
															}
															position++
															goto l955
														l956:
															position, tokenIndex = position955, tokenIndex955
															if buffer[position] != rune('C') {
																goto l929
															}
															position++
														}
													l955:
														if !_rules[rulews]() {
															goto l929
														}
														if !_rules[ruleLoc8]() {
															goto l929
														}
														{
															add(ruleAction9, position)
														}
														add(ruleInc8, position950)
													}
												}
											l931:
												add(ruleInc, position930)
											}
											goto l879
										l929:
											position, tokenIndex = position879, tokenIndex879
											{
												position959 := position
												{
													position960, tokenIndex960 := position, tokenIndex
													{
														position962 := position
														{
															position963, tokenIndex963 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l964
															}
															position++
															goto l963
														l964:
															position, tokenIndex = position963, tokenIndex963
															if buffer[position] != rune('D') {
																goto l961
															}
															position++
														}
													l963:
														{
															position965, tokenIndex965 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l966
															}
															position++
															goto l965
														l966:
															position, tokenIndex = position965, tokenIndex965
															if buffer[position] != rune('E') {
																goto l961
															}
															position++
														}
													l965:
														{
															position967, tokenIndex967 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l968
															}
															position++
															goto l967
														l968:
															position, tokenIndex = position967, tokenIndex967
															if buffer[position] != rune('C') {
																goto l961
															}
															position++
														}
													l967:
														if !_rules[rulews]() {
															goto l961
														}
														if !_rules[ruleILoc8]() {
															goto l961
														}
														{
															add(ruleAction11, position)
														}
														add(ruleDec16Indexed8, position962)
													}
													goto l960
												l961:
													position, tokenIndex = position960, tokenIndex960
													{
														position971 := position
														{
															position972, tokenIndex972 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l973
															}
															position++
															goto l972
														l973:
															position, tokenIndex = position972, tokenIndex972
															if buffer[position] != rune('D') {
																goto l970
															}
															position++
														}
													l972:
														{
															position974, tokenIndex974 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l975
															}
															position++
															goto l974
														l975:
															position, tokenIndex = position974, tokenIndex974
															if buffer[position] != rune('E') {
																goto l970
															}
															position++
														}
													l974:
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
																goto l970
															}
															position++
														}
													l976:
														if !_rules[rulews]() {
															goto l970
														}
														if !_rules[ruleLoc16]() {
															goto l970
														}
														{
															add(ruleAction13, position)
														}
														add(ruleDec16, position971)
													}
													goto l960
												l970:
													position, tokenIndex = position960, tokenIndex960
													{
														position979 := position
														{
															position980, tokenIndex980 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l981
															}
															position++
															goto l980
														l981:
															position, tokenIndex = position980, tokenIndex980
															if buffer[position] != rune('D') {
																goto l958
															}
															position++
														}
													l980:
														{
															position982, tokenIndex982 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l983
															}
															position++
															goto l982
														l983:
															position, tokenIndex = position982, tokenIndex982
															if buffer[position] != rune('E') {
																goto l958
															}
															position++
														}
													l982:
														{
															position984, tokenIndex984 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l985
															}
															position++
															goto l984
														l985:
															position, tokenIndex = position984, tokenIndex984
															if buffer[position] != rune('C') {
																goto l958
															}
															position++
														}
													l984:
														if !_rules[rulews]() {
															goto l958
														}
														if !_rules[ruleLoc8]() {
															goto l958
														}
														{
															add(ruleAction12, position)
														}
														add(ruleDec8, position979)
													}
												}
											l960:
												add(ruleDec, position959)
											}
											goto l879
										l958:
											position, tokenIndex = position879, tokenIndex879
											{
												position988 := position
												{
													position989, tokenIndex989 := position, tokenIndex
													{
														position991 := position
														{
															position992, tokenIndex992 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l993
															}
															position++
															goto l992
														l993:
															position, tokenIndex = position992, tokenIndex992
															if buffer[position] != rune('A') {
																goto l990
															}
															position++
														}
													l992:
														{
															position994, tokenIndex994 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l995
															}
															position++
															goto l994
														l995:
															position, tokenIndex = position994, tokenIndex994
															if buffer[position] != rune('D') {
																goto l990
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
																goto l990
															}
															position++
														}
													l996:
														if !_rules[rulews]() {
															goto l990
														}
														if !_rules[ruleDst16]() {
															goto l990
														}
														if !_rules[rulesep]() {
															goto l990
														}
														if !_rules[ruleSrc16]() {
															goto l990
														}
														{
															add(ruleAction14, position)
														}
														add(ruleAdd16, position991)
													}
													goto l989
												l990:
													position, tokenIndex = position989, tokenIndex989
													{
														position1000 := position
														{
															position1001, tokenIndex1001 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1002
															}
															position++
															goto l1001
														l1002:
															position, tokenIndex = position1001, tokenIndex1001
															if buffer[position] != rune('A') {
																goto l999
															}
															position++
														}
													l1001:
														{
															position1003, tokenIndex1003 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1004
															}
															position++
															goto l1003
														l1004:
															position, tokenIndex = position1003, tokenIndex1003
															if buffer[position] != rune('D') {
																goto l999
															}
															position++
														}
													l1003:
														{
															position1005, tokenIndex1005 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1006
															}
															position++
															goto l1005
														l1006:
															position, tokenIndex = position1005, tokenIndex1005
															if buffer[position] != rune('C') {
																goto l999
															}
															position++
														}
													l1005:
														if !_rules[rulews]() {
															goto l999
														}
														if !_rules[ruleDst16]() {
															goto l999
														}
														if !_rules[rulesep]() {
															goto l999
														}
														if !_rules[ruleSrc16]() {
															goto l999
														}
														{
															add(ruleAction15, position)
														}
														add(ruleAdc16, position1000)
													}
													goto l989
												l999:
													position, tokenIndex = position989, tokenIndex989
													{
														position1008 := position
														{
															position1009, tokenIndex1009 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1010
															}
															position++
															goto l1009
														l1010:
															position, tokenIndex = position1009, tokenIndex1009
															if buffer[position] != rune('S') {
																goto l987
															}
															position++
														}
													l1009:
														{
															position1011, tokenIndex1011 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1012
															}
															position++
															goto l1011
														l1012:
															position, tokenIndex = position1011, tokenIndex1011
															if buffer[position] != rune('B') {
																goto l987
															}
															position++
														}
													l1011:
														{
															position1013, tokenIndex1013 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1014
															}
															position++
															goto l1013
														l1014:
															position, tokenIndex = position1013, tokenIndex1013
															if buffer[position] != rune('C') {
																goto l987
															}
															position++
														}
													l1013:
														if !_rules[rulews]() {
															goto l987
														}
														if !_rules[ruleDst16]() {
															goto l987
														}
														if !_rules[rulesep]() {
															goto l987
														}
														if !_rules[ruleSrc16]() {
															goto l987
														}
														{
															add(ruleAction16, position)
														}
														add(ruleSbc16, position1008)
													}
												}
											l989:
												add(ruleAlu16, position988)
											}
											goto l879
										l987:
											position, tokenIndex = position879, tokenIndex879
											{
												position1017 := position
												{
													position1018, tokenIndex1018 := position, tokenIndex
													{
														position1020 := position
														{
															position1021, tokenIndex1021 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1022
															}
															position++
															goto l1021
														l1022:
															position, tokenIndex = position1021, tokenIndex1021
															if buffer[position] != rune('A') {
																goto l1019
															}
															position++
														}
													l1021:
														{
															position1023, tokenIndex1023 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1024
															}
															position++
															goto l1023
														l1024:
															position, tokenIndex = position1023, tokenIndex1023
															if buffer[position] != rune('D') {
																goto l1019
															}
															position++
														}
													l1023:
														{
															position1025, tokenIndex1025 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1026
															}
															position++
															goto l1025
														l1026:
															position, tokenIndex = position1025, tokenIndex1025
															if buffer[position] != rune('D') {
																goto l1019
															}
															position++
														}
													l1025:
														if !_rules[rulews]() {
															goto l1019
														}
														{
															position1027, tokenIndex1027 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1028
															}
															position++
															goto l1027
														l1028:
															position, tokenIndex = position1027, tokenIndex1027
															if buffer[position] != rune('A') {
																goto l1019
															}
															position++
														}
													l1027:
														if !_rules[rulesep]() {
															goto l1019
														}
														if !_rules[ruleSrc8]() {
															goto l1019
														}
														{
															add(ruleAction41, position)
														}
														add(ruleAdd, position1020)
													}
													goto l1018
												l1019:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1031 := position
														{
															position1032, tokenIndex1032 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1033
															}
															position++
															goto l1032
														l1033:
															position, tokenIndex = position1032, tokenIndex1032
															if buffer[position] != rune('A') {
																goto l1030
															}
															position++
														}
													l1032:
														{
															position1034, tokenIndex1034 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1035
															}
															position++
															goto l1034
														l1035:
															position, tokenIndex = position1034, tokenIndex1034
															if buffer[position] != rune('D') {
																goto l1030
															}
															position++
														}
													l1034:
														{
															position1036, tokenIndex1036 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1037
															}
															position++
															goto l1036
														l1037:
															position, tokenIndex = position1036, tokenIndex1036
															if buffer[position] != rune('C') {
																goto l1030
															}
															position++
														}
													l1036:
														if !_rules[rulews]() {
															goto l1030
														}
														{
															position1038, tokenIndex1038 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1039
															}
															position++
															goto l1038
														l1039:
															position, tokenIndex = position1038, tokenIndex1038
															if buffer[position] != rune('A') {
																goto l1030
															}
															position++
														}
													l1038:
														if !_rules[rulesep]() {
															goto l1030
														}
														if !_rules[ruleSrc8]() {
															goto l1030
														}
														{
															add(ruleAction42, position)
														}
														add(ruleAdc, position1031)
													}
													goto l1018
												l1030:
													position, tokenIndex = position1018, tokenIndex1018
													{
														position1042 := position
														{
															position1043, tokenIndex1043 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1044
															}
															position++
															goto l1043
														l1044:
															position, tokenIndex = position1043, tokenIndex1043
															if buffer[position] != rune('S') {
																goto l1041
															}
															position++
														}
													l1043:
														{
															position1045, tokenIndex1045 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1046
															}
															position++
															goto l1045
														l1046:
															position, tokenIndex = position1045, tokenIndex1045
															if buffer[position] != rune('U') {
																goto l1041
															}
															position++
														}
													l1045:
														{
															position1047, tokenIndex1047 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1048
															}
															position++
															goto l1047
														l1048:
															position, tokenIndex = position1047, tokenIndex1047
															if buffer[position] != rune('B') {
																goto l1041
															}
															position++
														}
													l1047:
														if !_rules[rulews]() {
															goto l1041
														}
														if !_rules[ruleSrc8]() {
															goto l1041
														}
														{
															add(ruleAction43, position)
														}
														add(ruleSub, position1042)
													}
													goto l1018
												l1041:
													position, tokenIndex = position1018, tokenIndex1018
													{
														switch buffer[position] {
														case 'C', 'c':
															{
																position1051 := position
																{
																	position1052, tokenIndex1052 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1053
																	}
																	position++
																	goto l1052
																l1053:
																	position, tokenIndex = position1052, tokenIndex1052
																	if buffer[position] != rune('C') {
																		goto l1016
																	}
																	position++
																}
															l1052:
																{
																	position1054, tokenIndex1054 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1055
																	}
																	position++
																	goto l1054
																l1055:
																	position, tokenIndex = position1054, tokenIndex1054
																	if buffer[position] != rune('P') {
																		goto l1016
																	}
																	position++
																}
															l1054:
																if !_rules[rulews]() {
																	goto l1016
																}
																if !_rules[ruleSrc8]() {
																	goto l1016
																}
																{
																	add(ruleAction48, position)
																}
																add(ruleCp, position1051)
															}
															break
														case 'O', 'o':
															{
																position1057 := position
																{
																	position1058, tokenIndex1058 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1059
																	}
																	position++
																	goto l1058
																l1059:
																	position, tokenIndex = position1058, tokenIndex1058
																	if buffer[position] != rune('O') {
																		goto l1016
																	}
																	position++
																}
															l1058:
																{
																	position1060, tokenIndex1060 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1061
																	}
																	position++
																	goto l1060
																l1061:
																	position, tokenIndex = position1060, tokenIndex1060
																	if buffer[position] != rune('R') {
																		goto l1016
																	}
																	position++
																}
															l1060:
																if !_rules[rulews]() {
																	goto l1016
																}
																if !_rules[ruleSrc8]() {
																	goto l1016
																}
																{
																	add(ruleAction47, position)
																}
																add(ruleOr, position1057)
															}
															break
														case 'X', 'x':
															{
																position1063 := position
																{
																	position1064, tokenIndex1064 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1065
																	}
																	position++
																	goto l1064
																l1065:
																	position, tokenIndex = position1064, tokenIndex1064
																	if buffer[position] != rune('X') {
																		goto l1016
																	}
																	position++
																}
															l1064:
																{
																	position1066, tokenIndex1066 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1067
																	}
																	position++
																	goto l1066
																l1067:
																	position, tokenIndex = position1066, tokenIndex1066
																	if buffer[position] != rune('O') {
																		goto l1016
																	}
																	position++
																}
															l1066:
																{
																	position1068, tokenIndex1068 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1069
																	}
																	position++
																	goto l1068
																l1069:
																	position, tokenIndex = position1068, tokenIndex1068
																	if buffer[position] != rune('R') {
																		goto l1016
																	}
																	position++
																}
															l1068:
																if !_rules[rulews]() {
																	goto l1016
																}
																if !_rules[ruleSrc8]() {
																	goto l1016
																}
																{
																	add(ruleAction46, position)
																}
																add(ruleXor, position1063)
															}
															break
														case 'A', 'a':
															{
																position1071 := position
																{
																	position1072, tokenIndex1072 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1073
																	}
																	position++
																	goto l1072
																l1073:
																	position, tokenIndex = position1072, tokenIndex1072
																	if buffer[position] != rune('A') {
																		goto l1016
																	}
																	position++
																}
															l1072:
																{
																	position1074, tokenIndex1074 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1075
																	}
																	position++
																	goto l1074
																l1075:
																	position, tokenIndex = position1074, tokenIndex1074
																	if buffer[position] != rune('N') {
																		goto l1016
																	}
																	position++
																}
															l1074:
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
																		goto l1016
																	}
																	position++
																}
															l1076:
																if !_rules[rulews]() {
																	goto l1016
																}
																if !_rules[ruleSrc8]() {
																	goto l1016
																}
																{
																	add(ruleAction45, position)
																}
																add(ruleAnd, position1071)
															}
															break
														default:
															{
																position1079 := position
																{
																	position1080, tokenIndex1080 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1081
																	}
																	position++
																	goto l1080
																l1081:
																	position, tokenIndex = position1080, tokenIndex1080
																	if buffer[position] != rune('S') {
																		goto l1016
																	}
																	position++
																}
															l1080:
																{
																	position1082, tokenIndex1082 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1083
																	}
																	position++
																	goto l1082
																l1083:
																	position, tokenIndex = position1082, tokenIndex1082
																	if buffer[position] != rune('B') {
																		goto l1016
																	}
																	position++
																}
															l1082:
																{
																	position1084, tokenIndex1084 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1085
																	}
																	position++
																	goto l1084
																l1085:
																	position, tokenIndex = position1084, tokenIndex1084
																	if buffer[position] != rune('C') {
																		goto l1016
																	}
																	position++
																}
															l1084:
																if !_rules[rulews]() {
																	goto l1016
																}
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
																		goto l1016
																	}
																	position++
																}
															l1086:
																if !_rules[rulesep]() {
																	goto l1016
																}
																if !_rules[ruleSrc8]() {
																	goto l1016
																}
																{
																	add(ruleAction44, position)
																}
																add(ruleSbc, position1079)
															}
															break
														}
													}

												}
											l1018:
												add(ruleAlu, position1017)
											}
											goto l879
										l1016:
											position, tokenIndex = position879, tokenIndex879
											{
												position1090 := position
												{
													position1091, tokenIndex1091 := position, tokenIndex
													{
														position1093 := position
														{
															position1094, tokenIndex1094 := position, tokenIndex
															{
																position1096 := position
																{
																	position1097, tokenIndex1097 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1098
																	}
																	position++
																	goto l1097
																l1098:
																	position, tokenIndex = position1097, tokenIndex1097
																	if buffer[position] != rune('R') {
																		goto l1095
																	}
																	position++
																}
															l1097:
																{
																	position1099, tokenIndex1099 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1100
																	}
																	position++
																	goto l1099
																l1100:
																	position, tokenIndex = position1099, tokenIndex1099
																	if buffer[position] != rune('L') {
																		goto l1095
																	}
																	position++
																}
															l1099:
																{
																	position1101, tokenIndex1101 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1102
																	}
																	position++
																	goto l1101
																l1102:
																	position, tokenIndex = position1101, tokenIndex1101
																	if buffer[position] != rune('C') {
																		goto l1095
																	}
																	position++
																}
															l1101:
																if !_rules[rulews]() {
																	goto l1095
																}
																if !_rules[ruleLoc8]() {
																	goto l1095
																}
																{
																	position1103, tokenIndex1103 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1103
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1103
																	}
																	goto l1104
																l1103:
																	position, tokenIndex = position1103, tokenIndex1103
																}
															l1104:
																{
																	add(ruleAction49, position)
																}
																add(ruleRlc, position1096)
															}
															goto l1094
														l1095:
															position, tokenIndex = position1094, tokenIndex1094
															{
																position1107 := position
																{
																	position1108, tokenIndex1108 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1109
																	}
																	position++
																	goto l1108
																l1109:
																	position, tokenIndex = position1108, tokenIndex1108
																	if buffer[position] != rune('R') {
																		goto l1106
																	}
																	position++
																}
															l1108:
																{
																	position1110, tokenIndex1110 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1111
																	}
																	position++
																	goto l1110
																l1111:
																	position, tokenIndex = position1110, tokenIndex1110
																	if buffer[position] != rune('R') {
																		goto l1106
																	}
																	position++
																}
															l1110:
																{
																	position1112, tokenIndex1112 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1113
																	}
																	position++
																	goto l1112
																l1113:
																	position, tokenIndex = position1112, tokenIndex1112
																	if buffer[position] != rune('C') {
																		goto l1106
																	}
																	position++
																}
															l1112:
																if !_rules[rulews]() {
																	goto l1106
																}
																if !_rules[ruleLoc8]() {
																	goto l1106
																}
																{
																	position1114, tokenIndex1114 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1114
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1114
																	}
																	goto l1115
																l1114:
																	position, tokenIndex = position1114, tokenIndex1114
																}
															l1115:
																{
																	add(ruleAction50, position)
																}
																add(ruleRrc, position1107)
															}
															goto l1094
														l1106:
															position, tokenIndex = position1094, tokenIndex1094
															{
																position1118 := position
																{
																	position1119, tokenIndex1119 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1120
																	}
																	position++
																	goto l1119
																l1120:
																	position, tokenIndex = position1119, tokenIndex1119
																	if buffer[position] != rune('R') {
																		goto l1117
																	}
																	position++
																}
															l1119:
																{
																	position1121, tokenIndex1121 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1122
																	}
																	position++
																	goto l1121
																l1122:
																	position, tokenIndex = position1121, tokenIndex1121
																	if buffer[position] != rune('L') {
																		goto l1117
																	}
																	position++
																}
															l1121:
																if !_rules[rulews]() {
																	goto l1117
																}
																if !_rules[ruleLoc8]() {
																	goto l1117
																}
																{
																	position1123, tokenIndex1123 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1123
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1123
																	}
																	goto l1124
																l1123:
																	position, tokenIndex = position1123, tokenIndex1123
																}
															l1124:
																{
																	add(ruleAction51, position)
																}
																add(ruleRl, position1118)
															}
															goto l1094
														l1117:
															position, tokenIndex = position1094, tokenIndex1094
															{
																position1127 := position
																{
																	position1128, tokenIndex1128 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1129
																	}
																	position++
																	goto l1128
																l1129:
																	position, tokenIndex = position1128, tokenIndex1128
																	if buffer[position] != rune('R') {
																		goto l1126
																	}
																	position++
																}
															l1128:
																{
																	position1130, tokenIndex1130 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1131
																	}
																	position++
																	goto l1130
																l1131:
																	position, tokenIndex = position1130, tokenIndex1130
																	if buffer[position] != rune('R') {
																		goto l1126
																	}
																	position++
																}
															l1130:
																if !_rules[rulews]() {
																	goto l1126
																}
																if !_rules[ruleLoc8]() {
																	goto l1126
																}
																{
																	position1132, tokenIndex1132 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1132
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1132
																	}
																	goto l1133
																l1132:
																	position, tokenIndex = position1132, tokenIndex1132
																}
															l1133:
																{
																	add(ruleAction52, position)
																}
																add(ruleRr, position1127)
															}
															goto l1094
														l1126:
															position, tokenIndex = position1094, tokenIndex1094
															{
																position1136 := position
																{
																	position1137, tokenIndex1137 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1138
																	}
																	position++
																	goto l1137
																l1138:
																	position, tokenIndex = position1137, tokenIndex1137
																	if buffer[position] != rune('S') {
																		goto l1135
																	}
																	position++
																}
															l1137:
																{
																	position1139, tokenIndex1139 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1140
																	}
																	position++
																	goto l1139
																l1140:
																	position, tokenIndex = position1139, tokenIndex1139
																	if buffer[position] != rune('L') {
																		goto l1135
																	}
																	position++
																}
															l1139:
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
																		goto l1135
																	}
																	position++
																}
															l1141:
																if !_rules[rulews]() {
																	goto l1135
																}
																if !_rules[ruleLoc8]() {
																	goto l1135
																}
																{
																	position1143, tokenIndex1143 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1143
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1143
																	}
																	goto l1144
																l1143:
																	position, tokenIndex = position1143, tokenIndex1143
																}
															l1144:
																{
																	add(ruleAction53, position)
																}
																add(ruleSla, position1136)
															}
															goto l1094
														l1135:
															position, tokenIndex = position1094, tokenIndex1094
															{
																position1147 := position
																{
																	position1148, tokenIndex1148 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1149
																	}
																	position++
																	goto l1148
																l1149:
																	position, tokenIndex = position1148, tokenIndex1148
																	if buffer[position] != rune('S') {
																		goto l1146
																	}
																	position++
																}
															l1148:
																{
																	position1150, tokenIndex1150 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1151
																	}
																	position++
																	goto l1150
																l1151:
																	position, tokenIndex = position1150, tokenIndex1150
																	if buffer[position] != rune('R') {
																		goto l1146
																	}
																	position++
																}
															l1150:
																{
																	position1152, tokenIndex1152 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1153
																	}
																	position++
																	goto l1152
																l1153:
																	position, tokenIndex = position1152, tokenIndex1152
																	if buffer[position] != rune('A') {
																		goto l1146
																	}
																	position++
																}
															l1152:
																if !_rules[rulews]() {
																	goto l1146
																}
																if !_rules[ruleLoc8]() {
																	goto l1146
																}
																{
																	position1154, tokenIndex1154 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1154
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1154
																	}
																	goto l1155
																l1154:
																	position, tokenIndex = position1154, tokenIndex1154
																}
															l1155:
																{
																	add(ruleAction54, position)
																}
																add(ruleSra, position1147)
															}
															goto l1094
														l1146:
															position, tokenIndex = position1094, tokenIndex1094
															{
																position1158 := position
																{
																	position1159, tokenIndex1159 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1160
																	}
																	position++
																	goto l1159
																l1160:
																	position, tokenIndex = position1159, tokenIndex1159
																	if buffer[position] != rune('S') {
																		goto l1157
																	}
																	position++
																}
															l1159:
																{
																	position1161, tokenIndex1161 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1162
																	}
																	position++
																	goto l1161
																l1162:
																	position, tokenIndex = position1161, tokenIndex1161
																	if buffer[position] != rune('L') {
																		goto l1157
																	}
																	position++
																}
															l1161:
																{
																	position1163, tokenIndex1163 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1164
																	}
																	position++
																	goto l1163
																l1164:
																	position, tokenIndex = position1163, tokenIndex1163
																	if buffer[position] != rune('L') {
																		goto l1157
																	}
																	position++
																}
															l1163:
																if !_rules[rulews]() {
																	goto l1157
																}
																if !_rules[ruleLoc8]() {
																	goto l1157
																}
																{
																	position1165, tokenIndex1165 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1165
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1165
																	}
																	goto l1166
																l1165:
																	position, tokenIndex = position1165, tokenIndex1165
																}
															l1166:
																{
																	add(ruleAction55, position)
																}
																add(ruleSll, position1158)
															}
															goto l1094
														l1157:
															position, tokenIndex = position1094, tokenIndex1094
															{
																position1168 := position
																{
																	position1169, tokenIndex1169 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1170
																	}
																	position++
																	goto l1169
																l1170:
																	position, tokenIndex = position1169, tokenIndex1169
																	if buffer[position] != rune('S') {
																		goto l1092
																	}
																	position++
																}
															l1169:
																{
																	position1171, tokenIndex1171 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1172
																	}
																	position++
																	goto l1171
																l1172:
																	position, tokenIndex = position1171, tokenIndex1171
																	if buffer[position] != rune('R') {
																		goto l1092
																	}
																	position++
																}
															l1171:
																{
																	position1173, tokenIndex1173 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1174
																	}
																	position++
																	goto l1173
																l1174:
																	position, tokenIndex = position1173, tokenIndex1173
																	if buffer[position] != rune('L') {
																		goto l1092
																	}
																	position++
																}
															l1173:
																if !_rules[rulews]() {
																	goto l1092
																}
																if !_rules[ruleLoc8]() {
																	goto l1092
																}
																{
																	position1175, tokenIndex1175 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1175
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1175
																	}
																	goto l1176
																l1175:
																	position, tokenIndex = position1175, tokenIndex1175
																}
															l1176:
																{
																	add(ruleAction56, position)
																}
																add(ruleSrl, position1168)
															}
														}
													l1094:
														add(ruleRot, position1093)
													}
													goto l1091
												l1092:
													position, tokenIndex = position1091, tokenIndex1091
													{
														switch buffer[position] {
														case 'S', 's':
															{
																position1179 := position
																{
																	position1180, tokenIndex1180 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1181
																	}
																	position++
																	goto l1180
																l1181:
																	position, tokenIndex = position1180, tokenIndex1180
																	if buffer[position] != rune('S') {
																		goto l1089
																	}
																	position++
																}
															l1180:
																{
																	position1182, tokenIndex1182 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1183
																	}
																	position++
																	goto l1182
																l1183:
																	position, tokenIndex = position1182, tokenIndex1182
																	if buffer[position] != rune('E') {
																		goto l1089
																	}
																	position++
																}
															l1182:
																{
																	position1184, tokenIndex1184 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1185
																	}
																	position++
																	goto l1184
																l1185:
																	position, tokenIndex = position1184, tokenIndex1184
																	if buffer[position] != rune('T') {
																		goto l1089
																	}
																	position++
																}
															l1184:
																if !_rules[rulews]() {
																	goto l1089
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1089
																}
																if !_rules[rulesep]() {
																	goto l1089
																}
																if !_rules[ruleLoc8]() {
																	goto l1089
																}
																{
																	position1186, tokenIndex1186 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1186
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1186
																	}
																	goto l1187
																l1186:
																	position, tokenIndex = position1186, tokenIndex1186
																}
															l1187:
																{
																	add(ruleAction59, position)
																}
																add(ruleSet, position1179)
															}
															break
														case 'R', 'r':
															{
																position1189 := position
																{
																	position1190, tokenIndex1190 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1191
																	}
																	position++
																	goto l1190
																l1191:
																	position, tokenIndex = position1190, tokenIndex1190
																	if buffer[position] != rune('R') {
																		goto l1089
																	}
																	position++
																}
															l1190:
																{
																	position1192, tokenIndex1192 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1193
																	}
																	position++
																	goto l1192
																l1193:
																	position, tokenIndex = position1192, tokenIndex1192
																	if buffer[position] != rune('E') {
																		goto l1089
																	}
																	position++
																}
															l1192:
																{
																	position1194, tokenIndex1194 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1195
																	}
																	position++
																	goto l1194
																l1195:
																	position, tokenIndex = position1194, tokenIndex1194
																	if buffer[position] != rune('S') {
																		goto l1089
																	}
																	position++
																}
															l1194:
																if !_rules[rulews]() {
																	goto l1089
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1089
																}
																if !_rules[rulesep]() {
																	goto l1089
																}
																if !_rules[ruleLoc8]() {
																	goto l1089
																}
																{
																	position1196, tokenIndex1196 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1196
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1196
																	}
																	goto l1197
																l1196:
																	position, tokenIndex = position1196, tokenIndex1196
																}
															l1197:
																{
																	add(ruleAction58, position)
																}
																add(ruleRes, position1189)
															}
															break
														default:
															{
																position1199 := position
																{
																	position1200, tokenIndex1200 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1201
																	}
																	position++
																	goto l1200
																l1201:
																	position, tokenIndex = position1200, tokenIndex1200
																	if buffer[position] != rune('B') {
																		goto l1089
																	}
																	position++
																}
															l1200:
																{
																	position1202, tokenIndex1202 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l1203
																	}
																	position++
																	goto l1202
																l1203:
																	position, tokenIndex = position1202, tokenIndex1202
																	if buffer[position] != rune('I') {
																		goto l1089
																	}
																	position++
																}
															l1202:
																{
																	position1204, tokenIndex1204 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1205
																	}
																	position++
																	goto l1204
																l1205:
																	position, tokenIndex = position1204, tokenIndex1204
																	if buffer[position] != rune('T') {
																		goto l1089
																	}
																	position++
																}
															l1204:
																if !_rules[rulews]() {
																	goto l1089
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1089
																}
																if !_rules[rulesep]() {
																	goto l1089
																}
																if !_rules[ruleLoc8]() {
																	goto l1089
																}
																{
																	add(ruleAction57, position)
																}
																add(ruleBit, position1199)
															}
															break
														}
													}

												}
											l1091:
												add(ruleBitOp, position1090)
											}
											goto l879
										l1089:
											position, tokenIndex = position879, tokenIndex879
											{
												position1208 := position
												{
													position1209, tokenIndex1209 := position, tokenIndex
													{
														position1211 := position
														{
															position1212 := position
															{
																position1213, tokenIndex1213 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1214
																}
																position++
																goto l1213
															l1214:
																position, tokenIndex = position1213, tokenIndex1213
																if buffer[position] != rune('R') {
																	goto l1210
																}
																position++
															}
														l1213:
															{
																position1215, tokenIndex1215 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1216
																}
																position++
																goto l1215
															l1216:
																position, tokenIndex = position1215, tokenIndex1215
																if buffer[position] != rune('E') {
																	goto l1210
																}
																position++
															}
														l1215:
															{
																position1217, tokenIndex1217 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1218
																}
																position++
																goto l1217
															l1218:
																position, tokenIndex = position1217, tokenIndex1217
																if buffer[position] != rune('T') {
																	goto l1210
																}
																position++
															}
														l1217:
															{
																position1219, tokenIndex1219 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1220
																}
																position++
																goto l1219
															l1220:
																position, tokenIndex = position1219, tokenIndex1219
																if buffer[position] != rune('N') {
																	goto l1210
																}
																position++
															}
														l1219:
															add(rulePegText, position1212)
														}
														{
															add(ruleAction74, position)
														}
														add(ruleRetn, position1211)
													}
													goto l1209
												l1210:
													position, tokenIndex = position1209, tokenIndex1209
													{
														position1223 := position
														{
															position1224 := position
															{
																position1225, tokenIndex1225 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1226
																}
																position++
																goto l1225
															l1226:
																position, tokenIndex = position1225, tokenIndex1225
																if buffer[position] != rune('R') {
																	goto l1222
																}
																position++
															}
														l1225:
															{
																position1227, tokenIndex1227 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1228
																}
																position++
																goto l1227
															l1228:
																position, tokenIndex = position1227, tokenIndex1227
																if buffer[position] != rune('E') {
																	goto l1222
																}
																position++
															}
														l1227:
															{
																position1229, tokenIndex1229 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1230
																}
																position++
																goto l1229
															l1230:
																position, tokenIndex = position1229, tokenIndex1229
																if buffer[position] != rune('T') {
																	goto l1222
																}
																position++
															}
														l1229:
															{
																position1231, tokenIndex1231 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1232
																}
																position++
																goto l1231
															l1232:
																position, tokenIndex = position1231, tokenIndex1231
																if buffer[position] != rune('I') {
																	goto l1222
																}
																position++
															}
														l1231:
															add(rulePegText, position1224)
														}
														{
															add(ruleAction75, position)
														}
														add(ruleReti, position1223)
													}
													goto l1209
												l1222:
													position, tokenIndex = position1209, tokenIndex1209
													{
														position1235 := position
														{
															position1236 := position
															{
																position1237, tokenIndex1237 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1238
																}
																position++
																goto l1237
															l1238:
																position, tokenIndex = position1237, tokenIndex1237
																if buffer[position] != rune('R') {
																	goto l1234
																}
																position++
															}
														l1237:
															{
																position1239, tokenIndex1239 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1240
																}
																position++
																goto l1239
															l1240:
																position, tokenIndex = position1239, tokenIndex1239
																if buffer[position] != rune('R') {
																	goto l1234
																}
																position++
															}
														l1239:
															{
																position1241, tokenIndex1241 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1242
																}
																position++
																goto l1241
															l1242:
																position, tokenIndex = position1241, tokenIndex1241
																if buffer[position] != rune('D') {
																	goto l1234
																}
																position++
															}
														l1241:
															add(rulePegText, position1236)
														}
														{
															add(ruleAction76, position)
														}
														add(ruleRrd, position1235)
													}
													goto l1209
												l1234:
													position, tokenIndex = position1209, tokenIndex1209
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
																if buffer[position] != rune('m') {
																	goto l1250
																}
																position++
																goto l1249
															l1250:
																position, tokenIndex = position1249, tokenIndex1249
																if buffer[position] != rune('M') {
																	goto l1244
																}
																position++
															}
														l1249:
															if buffer[position] != rune(' ') {
																goto l1244
															}
															position++
															if buffer[position] != rune('0') {
																goto l1244
															}
															position++
															add(rulePegText, position1246)
														}
														{
															add(ruleAction78, position)
														}
														add(ruleIm0, position1245)
													}
													goto l1209
												l1244:
													position, tokenIndex = position1209, tokenIndex1209
													{
														position1253 := position
														{
															position1254 := position
															{
																position1255, tokenIndex1255 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1256
																}
																position++
																goto l1255
															l1256:
																position, tokenIndex = position1255, tokenIndex1255
																if buffer[position] != rune('I') {
																	goto l1252
																}
																position++
															}
														l1255:
															{
																position1257, tokenIndex1257 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1258
																}
																position++
																goto l1257
															l1258:
																position, tokenIndex = position1257, tokenIndex1257
																if buffer[position] != rune('M') {
																	goto l1252
																}
																position++
															}
														l1257:
															if buffer[position] != rune(' ') {
																goto l1252
															}
															position++
															if buffer[position] != rune('1') {
																goto l1252
															}
															position++
															add(rulePegText, position1254)
														}
														{
															add(ruleAction79, position)
														}
														add(ruleIm1, position1253)
													}
													goto l1209
												l1252:
													position, tokenIndex = position1209, tokenIndex1209
													{
														position1261 := position
														{
															position1262 := position
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
																	goto l1260
																}
																position++
															}
														l1263:
															{
																position1265, tokenIndex1265 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1266
																}
																position++
																goto l1265
															l1266:
																position, tokenIndex = position1265, tokenIndex1265
																if buffer[position] != rune('M') {
																	goto l1260
																}
																position++
															}
														l1265:
															if buffer[position] != rune(' ') {
																goto l1260
															}
															position++
															if buffer[position] != rune('2') {
																goto l1260
															}
															position++
															add(rulePegText, position1262)
														}
														{
															add(ruleAction80, position)
														}
														add(ruleIm2, position1261)
													}
													goto l1209
												l1260:
													position, tokenIndex = position1209, tokenIndex1209
													{
														switch buffer[position] {
														case 'I', 'O', 'i', 'o':
															{
																position1269 := position
																{
																	position1270, tokenIndex1270 := position, tokenIndex
																	{
																		position1272 := position
																		{
																			position1273 := position
																			{
																				position1274, tokenIndex1274 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1275
																				}
																				position++
																				goto l1274
																			l1275:
																				position, tokenIndex = position1274, tokenIndex1274
																				if buffer[position] != rune('I') {
																					goto l1271
																				}
																				position++
																			}
																		l1274:
																			{
																				position1276, tokenIndex1276 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1277
																				}
																				position++
																				goto l1276
																			l1277:
																				position, tokenIndex = position1276, tokenIndex1276
																				if buffer[position] != rune('N') {
																					goto l1271
																				}
																				position++
																			}
																		l1276:
																			{
																				position1278, tokenIndex1278 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1279
																				}
																				position++
																				goto l1278
																			l1279:
																				position, tokenIndex = position1278, tokenIndex1278
																				if buffer[position] != rune('I') {
																					goto l1271
																				}
																				position++
																			}
																		l1278:
																			{
																				position1280, tokenIndex1280 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1281
																				}
																				position++
																				goto l1280
																			l1281:
																				position, tokenIndex = position1280, tokenIndex1280
																				if buffer[position] != rune('R') {
																					goto l1271
																				}
																				position++
																			}
																		l1280:
																			add(rulePegText, position1273)
																		}
																		{
																			add(ruleAction91, position)
																		}
																		add(ruleInir, position1272)
																	}
																	goto l1270
																l1271:
																	position, tokenIndex = position1270, tokenIndex1270
																	{
																		position1284 := position
																		{
																			position1285 := position
																			{
																				position1286, tokenIndex1286 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1287
																				}
																				position++
																				goto l1286
																			l1287:
																				position, tokenIndex = position1286, tokenIndex1286
																				if buffer[position] != rune('I') {
																					goto l1283
																				}
																				position++
																			}
																		l1286:
																			{
																				position1288, tokenIndex1288 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1289
																				}
																				position++
																				goto l1288
																			l1289:
																				position, tokenIndex = position1288, tokenIndex1288
																				if buffer[position] != rune('N') {
																					goto l1283
																				}
																				position++
																			}
																		l1288:
																			{
																				position1290, tokenIndex1290 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1291
																				}
																				position++
																				goto l1290
																			l1291:
																				position, tokenIndex = position1290, tokenIndex1290
																				if buffer[position] != rune('I') {
																					goto l1283
																				}
																				position++
																			}
																		l1290:
																			add(rulePegText, position1285)
																		}
																		{
																			add(ruleAction83, position)
																		}
																		add(ruleIni, position1284)
																	}
																	goto l1270
																l1283:
																	position, tokenIndex = position1270, tokenIndex1270
																	{
																		position1294 := position
																		{
																			position1295 := position
																			{
																				position1296, tokenIndex1296 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1297
																				}
																				position++
																				goto l1296
																			l1297:
																				position, tokenIndex = position1296, tokenIndex1296
																				if buffer[position] != rune('O') {
																					goto l1293
																				}
																				position++
																			}
																		l1296:
																			{
																				position1298, tokenIndex1298 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1299
																				}
																				position++
																				goto l1298
																			l1299:
																				position, tokenIndex = position1298, tokenIndex1298
																				if buffer[position] != rune('T') {
																					goto l1293
																				}
																				position++
																			}
																		l1298:
																			{
																				position1300, tokenIndex1300 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1301
																				}
																				position++
																				goto l1300
																			l1301:
																				position, tokenIndex = position1300, tokenIndex1300
																				if buffer[position] != rune('I') {
																					goto l1293
																				}
																				position++
																			}
																		l1300:
																			{
																				position1302, tokenIndex1302 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1303
																				}
																				position++
																				goto l1302
																			l1303:
																				position, tokenIndex = position1302, tokenIndex1302
																				if buffer[position] != rune('R') {
																					goto l1293
																				}
																				position++
																			}
																		l1302:
																			add(rulePegText, position1295)
																		}
																		{
																			add(ruleAction92, position)
																		}
																		add(ruleOtir, position1294)
																	}
																	goto l1270
																l1293:
																	position, tokenIndex = position1270, tokenIndex1270
																	{
																		position1306 := position
																		{
																			position1307 := position
																			{
																				position1308, tokenIndex1308 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1309
																				}
																				position++
																				goto l1308
																			l1309:
																				position, tokenIndex = position1308, tokenIndex1308
																				if buffer[position] != rune('O') {
																					goto l1305
																				}
																				position++
																			}
																		l1308:
																			{
																				position1310, tokenIndex1310 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1311
																				}
																				position++
																				goto l1310
																			l1311:
																				position, tokenIndex = position1310, tokenIndex1310
																				if buffer[position] != rune('U') {
																					goto l1305
																				}
																				position++
																			}
																		l1310:
																			{
																				position1312, tokenIndex1312 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1313
																				}
																				position++
																				goto l1312
																			l1313:
																				position, tokenIndex = position1312, tokenIndex1312
																				if buffer[position] != rune('T') {
																					goto l1305
																				}
																				position++
																			}
																		l1312:
																			{
																				position1314, tokenIndex1314 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1315
																				}
																				position++
																				goto l1314
																			l1315:
																				position, tokenIndex = position1314, tokenIndex1314
																				if buffer[position] != rune('I') {
																					goto l1305
																				}
																				position++
																			}
																		l1314:
																			add(rulePegText, position1307)
																		}
																		{
																			add(ruleAction84, position)
																		}
																		add(ruleOuti, position1306)
																	}
																	goto l1270
																l1305:
																	position, tokenIndex = position1270, tokenIndex1270
																	{
																		position1318 := position
																		{
																			position1319 := position
																			{
																				position1320, tokenIndex1320 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1321
																				}
																				position++
																				goto l1320
																			l1321:
																				position, tokenIndex = position1320, tokenIndex1320
																				if buffer[position] != rune('I') {
																					goto l1317
																				}
																				position++
																			}
																		l1320:
																			{
																				position1322, tokenIndex1322 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1323
																				}
																				position++
																				goto l1322
																			l1323:
																				position, tokenIndex = position1322, tokenIndex1322
																				if buffer[position] != rune('N') {
																					goto l1317
																				}
																				position++
																			}
																		l1322:
																			{
																				position1324, tokenIndex1324 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1325
																				}
																				position++
																				goto l1324
																			l1325:
																				position, tokenIndex = position1324, tokenIndex1324
																				if buffer[position] != rune('D') {
																					goto l1317
																				}
																				position++
																			}
																		l1324:
																			{
																				position1326, tokenIndex1326 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1327
																				}
																				position++
																				goto l1326
																			l1327:
																				position, tokenIndex = position1326, tokenIndex1326
																				if buffer[position] != rune('R') {
																					goto l1317
																				}
																				position++
																			}
																		l1326:
																			add(rulePegText, position1319)
																		}
																		{
																			add(ruleAction95, position)
																		}
																		add(ruleIndr, position1318)
																	}
																	goto l1270
																l1317:
																	position, tokenIndex = position1270, tokenIndex1270
																	{
																		position1330 := position
																		{
																			position1331 := position
																			{
																				position1332, tokenIndex1332 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1333
																				}
																				position++
																				goto l1332
																			l1333:
																				position, tokenIndex = position1332, tokenIndex1332
																				if buffer[position] != rune('I') {
																					goto l1329
																				}
																				position++
																			}
																		l1332:
																			{
																				position1334, tokenIndex1334 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1335
																				}
																				position++
																				goto l1334
																			l1335:
																				position, tokenIndex = position1334, tokenIndex1334
																				if buffer[position] != rune('N') {
																					goto l1329
																				}
																				position++
																			}
																		l1334:
																			{
																				position1336, tokenIndex1336 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1337
																				}
																				position++
																				goto l1336
																			l1337:
																				position, tokenIndex = position1336, tokenIndex1336
																				if buffer[position] != rune('D') {
																					goto l1329
																				}
																				position++
																			}
																		l1336:
																			add(rulePegText, position1331)
																		}
																		{
																			add(ruleAction87, position)
																		}
																		add(ruleInd, position1330)
																	}
																	goto l1270
																l1329:
																	position, tokenIndex = position1270, tokenIndex1270
																	{
																		position1340 := position
																		{
																			position1341 := position
																			{
																				position1342, tokenIndex1342 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1343
																				}
																				position++
																				goto l1342
																			l1343:
																				position, tokenIndex = position1342, tokenIndex1342
																				if buffer[position] != rune('O') {
																					goto l1339
																				}
																				position++
																			}
																		l1342:
																			{
																				position1344, tokenIndex1344 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1345
																				}
																				position++
																				goto l1344
																			l1345:
																				position, tokenIndex = position1344, tokenIndex1344
																				if buffer[position] != rune('T') {
																					goto l1339
																				}
																				position++
																			}
																		l1344:
																			{
																				position1346, tokenIndex1346 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1347
																				}
																				position++
																				goto l1346
																			l1347:
																				position, tokenIndex = position1346, tokenIndex1346
																				if buffer[position] != rune('D') {
																					goto l1339
																				}
																				position++
																			}
																		l1346:
																			{
																				position1348, tokenIndex1348 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1349
																				}
																				position++
																				goto l1348
																			l1349:
																				position, tokenIndex = position1348, tokenIndex1348
																				if buffer[position] != rune('R') {
																					goto l1339
																				}
																				position++
																			}
																		l1348:
																			add(rulePegText, position1341)
																		}
																		{
																			add(ruleAction96, position)
																		}
																		add(ruleOtdr, position1340)
																	}
																	goto l1270
																l1339:
																	position, tokenIndex = position1270, tokenIndex1270
																	{
																		position1351 := position
																		{
																			position1352 := position
																			{
																				position1353, tokenIndex1353 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1354
																				}
																				position++
																				goto l1353
																			l1354:
																				position, tokenIndex = position1353, tokenIndex1353
																				if buffer[position] != rune('O') {
																					goto l1207
																				}
																				position++
																			}
																		l1353:
																			{
																				position1355, tokenIndex1355 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1356
																				}
																				position++
																				goto l1355
																			l1356:
																				position, tokenIndex = position1355, tokenIndex1355
																				if buffer[position] != rune('U') {
																					goto l1207
																				}
																				position++
																			}
																		l1355:
																			{
																				position1357, tokenIndex1357 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1358
																				}
																				position++
																				goto l1357
																			l1358:
																				position, tokenIndex = position1357, tokenIndex1357
																				if buffer[position] != rune('T') {
																					goto l1207
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
																					goto l1207
																				}
																				position++
																			}
																		l1359:
																			add(rulePegText, position1352)
																		}
																		{
																			add(ruleAction88, position)
																		}
																		add(ruleOutd, position1351)
																	}
																}
															l1270:
																add(ruleBlitIO, position1269)
															}
															break
														case 'R', 'r':
															{
																position1362 := position
																{
																	position1363 := position
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
																			goto l1207
																		}
																		position++
																	}
																l1364:
																	{
																		position1366, tokenIndex1366 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1367
																		}
																		position++
																		goto l1366
																	l1367:
																		position, tokenIndex = position1366, tokenIndex1366
																		if buffer[position] != rune('L') {
																			goto l1207
																		}
																		position++
																	}
																l1366:
																	{
																		position1368, tokenIndex1368 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1369
																		}
																		position++
																		goto l1368
																	l1369:
																		position, tokenIndex = position1368, tokenIndex1368
																		if buffer[position] != rune('D') {
																			goto l1207
																		}
																		position++
																	}
																l1368:
																	add(rulePegText, position1363)
																}
																{
																	add(ruleAction77, position)
																}
																add(ruleRld, position1362)
															}
															break
														case 'N', 'n':
															{
																position1371 := position
																{
																	position1372 := position
																	{
																		position1373, tokenIndex1373 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1374
																		}
																		position++
																		goto l1373
																	l1374:
																		position, tokenIndex = position1373, tokenIndex1373
																		if buffer[position] != rune('N') {
																			goto l1207
																		}
																		position++
																	}
																l1373:
																	{
																		position1375, tokenIndex1375 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1376
																		}
																		position++
																		goto l1375
																	l1376:
																		position, tokenIndex = position1375, tokenIndex1375
																		if buffer[position] != rune('E') {
																			goto l1207
																		}
																		position++
																	}
																l1375:
																	{
																		position1377, tokenIndex1377 := position, tokenIndex
																		if buffer[position] != rune('g') {
																			goto l1378
																		}
																		position++
																		goto l1377
																	l1378:
																		position, tokenIndex = position1377, tokenIndex1377
																		if buffer[position] != rune('G') {
																			goto l1207
																		}
																		position++
																	}
																l1377:
																	add(rulePegText, position1372)
																}
																{
																	add(ruleAction73, position)
																}
																add(ruleNeg, position1371)
															}
															break
														default:
															{
																position1380 := position
																{
																	position1381, tokenIndex1381 := position, tokenIndex
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
																				if buffer[position] != rune('i') {
																					goto l1390
																				}
																				position++
																				goto l1389
																			l1390:
																				position, tokenIndex = position1389, tokenIndex1389
																				if buffer[position] != rune('I') {
																					goto l1382
																				}
																				position++
																			}
																		l1389:
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
																					goto l1382
																				}
																				position++
																			}
																		l1391:
																			add(rulePegText, position1384)
																		}
																		{
																			add(ruleAction89, position)
																		}
																		add(ruleLdir, position1383)
																	}
																	goto l1381
																l1382:
																	position, tokenIndex = position1381, tokenIndex1381
																	{
																		position1395 := position
																		{
																			position1396 := position
																			{
																				position1397, tokenIndex1397 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1398
																				}
																				position++
																				goto l1397
																			l1398:
																				position, tokenIndex = position1397, tokenIndex1397
																				if buffer[position] != rune('L') {
																					goto l1394
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
																					goto l1394
																				}
																				position++
																			}
																		l1399:
																			{
																				position1401, tokenIndex1401 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1402
																				}
																				position++
																				goto l1401
																			l1402:
																				position, tokenIndex = position1401, tokenIndex1401
																				if buffer[position] != rune('I') {
																					goto l1394
																				}
																				position++
																			}
																		l1401:
																			add(rulePegText, position1396)
																		}
																		{
																			add(ruleAction81, position)
																		}
																		add(ruleLdi, position1395)
																	}
																	goto l1381
																l1394:
																	position, tokenIndex = position1381, tokenIndex1381
																	{
																		position1405 := position
																		{
																			position1406 := position
																			{
																				position1407, tokenIndex1407 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1408
																				}
																				position++
																				goto l1407
																			l1408:
																				position, tokenIndex = position1407, tokenIndex1407
																				if buffer[position] != rune('C') {
																					goto l1404
																				}
																				position++
																			}
																		l1407:
																			{
																				position1409, tokenIndex1409 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1410
																				}
																				position++
																				goto l1409
																			l1410:
																				position, tokenIndex = position1409, tokenIndex1409
																				if buffer[position] != rune('P') {
																					goto l1404
																				}
																				position++
																			}
																		l1409:
																			{
																				position1411, tokenIndex1411 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1412
																				}
																				position++
																				goto l1411
																			l1412:
																				position, tokenIndex = position1411, tokenIndex1411
																				if buffer[position] != rune('I') {
																					goto l1404
																				}
																				position++
																			}
																		l1411:
																			{
																				position1413, tokenIndex1413 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1414
																				}
																				position++
																				goto l1413
																			l1414:
																				position, tokenIndex = position1413, tokenIndex1413
																				if buffer[position] != rune('R') {
																					goto l1404
																				}
																				position++
																			}
																		l1413:
																			add(rulePegText, position1406)
																		}
																		{
																			add(ruleAction90, position)
																		}
																		add(ruleCpir, position1405)
																	}
																	goto l1381
																l1404:
																	position, tokenIndex = position1381, tokenIndex1381
																	{
																		position1417 := position
																		{
																			position1418 := position
																			{
																				position1419, tokenIndex1419 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1420
																				}
																				position++
																				goto l1419
																			l1420:
																				position, tokenIndex = position1419, tokenIndex1419
																				if buffer[position] != rune('C') {
																					goto l1416
																				}
																				position++
																			}
																		l1419:
																			{
																				position1421, tokenIndex1421 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1422
																				}
																				position++
																				goto l1421
																			l1422:
																				position, tokenIndex = position1421, tokenIndex1421
																				if buffer[position] != rune('P') {
																					goto l1416
																				}
																				position++
																			}
																		l1421:
																			{
																				position1423, tokenIndex1423 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1424
																				}
																				position++
																				goto l1423
																			l1424:
																				position, tokenIndex = position1423, tokenIndex1423
																				if buffer[position] != rune('I') {
																					goto l1416
																				}
																				position++
																			}
																		l1423:
																			add(rulePegText, position1418)
																		}
																		{
																			add(ruleAction82, position)
																		}
																		add(ruleCpi, position1417)
																	}
																	goto l1381
																l1416:
																	position, tokenIndex = position1381, tokenIndex1381
																	{
																		position1427 := position
																		{
																			position1428 := position
																			{
																				position1429, tokenIndex1429 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1430
																				}
																				position++
																				goto l1429
																			l1430:
																				position, tokenIndex = position1429, tokenIndex1429
																				if buffer[position] != rune('L') {
																					goto l1426
																				}
																				position++
																			}
																		l1429:
																			{
																				position1431, tokenIndex1431 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1432
																				}
																				position++
																				goto l1431
																			l1432:
																				position, tokenIndex = position1431, tokenIndex1431
																				if buffer[position] != rune('D') {
																					goto l1426
																				}
																				position++
																			}
																		l1431:
																			{
																				position1433, tokenIndex1433 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1434
																				}
																				position++
																				goto l1433
																			l1434:
																				position, tokenIndex = position1433, tokenIndex1433
																				if buffer[position] != rune('D') {
																					goto l1426
																				}
																				position++
																			}
																		l1433:
																			{
																				position1435, tokenIndex1435 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1436
																				}
																				position++
																				goto l1435
																			l1436:
																				position, tokenIndex = position1435, tokenIndex1435
																				if buffer[position] != rune('R') {
																					goto l1426
																				}
																				position++
																			}
																		l1435:
																			add(rulePegText, position1428)
																		}
																		{
																			add(ruleAction93, position)
																		}
																		add(ruleLddr, position1427)
																	}
																	goto l1381
																l1426:
																	position, tokenIndex = position1381, tokenIndex1381
																	{
																		position1439 := position
																		{
																			position1440 := position
																			{
																				position1441, tokenIndex1441 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1442
																				}
																				position++
																				goto l1441
																			l1442:
																				position, tokenIndex = position1441, tokenIndex1441
																				if buffer[position] != rune('L') {
																					goto l1438
																				}
																				position++
																			}
																		l1441:
																			{
																				position1443, tokenIndex1443 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1444
																				}
																				position++
																				goto l1443
																			l1444:
																				position, tokenIndex = position1443, tokenIndex1443
																				if buffer[position] != rune('D') {
																					goto l1438
																				}
																				position++
																			}
																		l1443:
																			{
																				position1445, tokenIndex1445 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1446
																				}
																				position++
																				goto l1445
																			l1446:
																				position, tokenIndex = position1445, tokenIndex1445
																				if buffer[position] != rune('D') {
																					goto l1438
																				}
																				position++
																			}
																		l1445:
																			add(rulePegText, position1440)
																		}
																		{
																			add(ruleAction85, position)
																		}
																		add(ruleLdd, position1439)
																	}
																	goto l1381
																l1438:
																	position, tokenIndex = position1381, tokenIndex1381
																	{
																		position1449 := position
																		{
																			position1450 := position
																			{
																				position1451, tokenIndex1451 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1452
																				}
																				position++
																				goto l1451
																			l1452:
																				position, tokenIndex = position1451, tokenIndex1451
																				if buffer[position] != rune('C') {
																					goto l1448
																				}
																				position++
																			}
																		l1451:
																			{
																				position1453, tokenIndex1453 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1454
																				}
																				position++
																				goto l1453
																			l1454:
																				position, tokenIndex = position1453, tokenIndex1453
																				if buffer[position] != rune('P') {
																					goto l1448
																				}
																				position++
																			}
																		l1453:
																			{
																				position1455, tokenIndex1455 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1456
																				}
																				position++
																				goto l1455
																			l1456:
																				position, tokenIndex = position1455, tokenIndex1455
																				if buffer[position] != rune('D') {
																					goto l1448
																				}
																				position++
																			}
																		l1455:
																			{
																				position1457, tokenIndex1457 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1458
																				}
																				position++
																				goto l1457
																			l1458:
																				position, tokenIndex = position1457, tokenIndex1457
																				if buffer[position] != rune('R') {
																					goto l1448
																				}
																				position++
																			}
																		l1457:
																			add(rulePegText, position1450)
																		}
																		{
																			add(ruleAction94, position)
																		}
																		add(ruleCpdr, position1449)
																	}
																	goto l1381
																l1448:
																	position, tokenIndex = position1381, tokenIndex1381
																	{
																		position1460 := position
																		{
																			position1461 := position
																			{
																				position1462, tokenIndex1462 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1463
																				}
																				position++
																				goto l1462
																			l1463:
																				position, tokenIndex = position1462, tokenIndex1462
																				if buffer[position] != rune('C') {
																					goto l1207
																				}
																				position++
																			}
																		l1462:
																			{
																				position1464, tokenIndex1464 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1465
																				}
																				position++
																				goto l1464
																			l1465:
																				position, tokenIndex = position1464, tokenIndex1464
																				if buffer[position] != rune('P') {
																					goto l1207
																				}
																				position++
																			}
																		l1464:
																			{
																				position1466, tokenIndex1466 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1467
																				}
																				position++
																				goto l1466
																			l1467:
																				position, tokenIndex = position1466, tokenIndex1466
																				if buffer[position] != rune('D') {
																					goto l1207
																				}
																				position++
																			}
																		l1466:
																			add(rulePegText, position1461)
																		}
																		{
																			add(ruleAction86, position)
																		}
																		add(ruleCpd, position1460)
																	}
																}
															l1381:
																add(ruleBlit, position1380)
															}
															break
														}
													}

												}
											l1209:
												add(ruleEDSimple, position1208)
											}
											goto l879
										l1207:
											position, tokenIndex = position879, tokenIndex879
											{
												position1470 := position
												{
													position1471, tokenIndex1471 := position, tokenIndex
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
																	goto l1472
																}
																position++
															}
														l1475:
															{
																position1477, tokenIndex1477 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1478
																}
																position++
																goto l1477
															l1478:
																position, tokenIndex = position1477, tokenIndex1477
																if buffer[position] != rune('L') {
																	goto l1472
																}
																position++
															}
														l1477:
															{
																position1479, tokenIndex1479 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1480
																}
																position++
																goto l1479
															l1480:
																position, tokenIndex = position1479, tokenIndex1479
																if buffer[position] != rune('C') {
																	goto l1472
																}
																position++
															}
														l1479:
															{
																position1481, tokenIndex1481 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1482
																}
																position++
																goto l1481
															l1482:
																position, tokenIndex = position1481, tokenIndex1481
																if buffer[position] != rune('A') {
																	goto l1472
																}
																position++
															}
														l1481:
															add(rulePegText, position1474)
														}
														{
															add(ruleAction62, position)
														}
														add(ruleRlca, position1473)
													}
													goto l1471
												l1472:
													position, tokenIndex = position1471, tokenIndex1471
													{
														position1485 := position
														{
															position1486 := position
															{
																position1487, tokenIndex1487 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1488
																}
																position++
																goto l1487
															l1488:
																position, tokenIndex = position1487, tokenIndex1487
																if buffer[position] != rune('R') {
																	goto l1484
																}
																position++
															}
														l1487:
															{
																position1489, tokenIndex1489 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1490
																}
																position++
																goto l1489
															l1490:
																position, tokenIndex = position1489, tokenIndex1489
																if buffer[position] != rune('R') {
																	goto l1484
																}
																position++
															}
														l1489:
															{
																position1491, tokenIndex1491 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1492
																}
																position++
																goto l1491
															l1492:
																position, tokenIndex = position1491, tokenIndex1491
																if buffer[position] != rune('C') {
																	goto l1484
																}
																position++
															}
														l1491:
															{
																position1493, tokenIndex1493 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1494
																}
																position++
																goto l1493
															l1494:
																position, tokenIndex = position1493, tokenIndex1493
																if buffer[position] != rune('A') {
																	goto l1484
																}
																position++
															}
														l1493:
															add(rulePegText, position1486)
														}
														{
															add(ruleAction63, position)
														}
														add(ruleRrca, position1485)
													}
													goto l1471
												l1484:
													position, tokenIndex = position1471, tokenIndex1471
													{
														position1497 := position
														{
															position1498 := position
															{
																position1499, tokenIndex1499 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1500
																}
																position++
																goto l1499
															l1500:
																position, tokenIndex = position1499, tokenIndex1499
																if buffer[position] != rune('R') {
																	goto l1496
																}
																position++
															}
														l1499:
															{
																position1501, tokenIndex1501 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1502
																}
																position++
																goto l1501
															l1502:
																position, tokenIndex = position1501, tokenIndex1501
																if buffer[position] != rune('L') {
																	goto l1496
																}
																position++
															}
														l1501:
															{
																position1503, tokenIndex1503 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1504
																}
																position++
																goto l1503
															l1504:
																position, tokenIndex = position1503, tokenIndex1503
																if buffer[position] != rune('A') {
																	goto l1496
																}
																position++
															}
														l1503:
															add(rulePegText, position1498)
														}
														{
															add(ruleAction64, position)
														}
														add(ruleRla, position1497)
													}
													goto l1471
												l1496:
													position, tokenIndex = position1471, tokenIndex1471
													{
														position1507 := position
														{
															position1508 := position
															{
																position1509, tokenIndex1509 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1510
																}
																position++
																goto l1509
															l1510:
																position, tokenIndex = position1509, tokenIndex1509
																if buffer[position] != rune('D') {
																	goto l1506
																}
																position++
															}
														l1509:
															{
																position1511, tokenIndex1511 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1512
																}
																position++
																goto l1511
															l1512:
																position, tokenIndex = position1511, tokenIndex1511
																if buffer[position] != rune('A') {
																	goto l1506
																}
																position++
															}
														l1511:
															{
																position1513, tokenIndex1513 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1514
																}
																position++
																goto l1513
															l1514:
																position, tokenIndex = position1513, tokenIndex1513
																if buffer[position] != rune('A') {
																	goto l1506
																}
																position++
															}
														l1513:
															add(rulePegText, position1508)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleDaa, position1507)
													}
													goto l1471
												l1506:
													position, tokenIndex = position1471, tokenIndex1471
													{
														position1517 := position
														{
															position1518 := position
															{
																position1519, tokenIndex1519 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1520
																}
																position++
																goto l1519
															l1520:
																position, tokenIndex = position1519, tokenIndex1519
																if buffer[position] != rune('C') {
																	goto l1516
																}
																position++
															}
														l1519:
															{
																position1521, tokenIndex1521 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1522
																}
																position++
																goto l1521
															l1522:
																position, tokenIndex = position1521, tokenIndex1521
																if buffer[position] != rune('P') {
																	goto l1516
																}
																position++
															}
														l1521:
															{
																position1523, tokenIndex1523 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1524
																}
																position++
																goto l1523
															l1524:
																position, tokenIndex = position1523, tokenIndex1523
																if buffer[position] != rune('L') {
																	goto l1516
																}
																position++
															}
														l1523:
															add(rulePegText, position1518)
														}
														{
															add(ruleAction67, position)
														}
														add(ruleCpl, position1517)
													}
													goto l1471
												l1516:
													position, tokenIndex = position1471, tokenIndex1471
													{
														position1527 := position
														{
															position1528 := position
															{
																position1529, tokenIndex1529 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1530
																}
																position++
																goto l1529
															l1530:
																position, tokenIndex = position1529, tokenIndex1529
																if buffer[position] != rune('E') {
																	goto l1526
																}
																position++
															}
														l1529:
															{
																position1531, tokenIndex1531 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1532
																}
																position++
																goto l1531
															l1532:
																position, tokenIndex = position1531, tokenIndex1531
																if buffer[position] != rune('X') {
																	goto l1526
																}
																position++
															}
														l1531:
															{
																position1533, tokenIndex1533 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1534
																}
																position++
																goto l1533
															l1534:
																position, tokenIndex = position1533, tokenIndex1533
																if buffer[position] != rune('X') {
																	goto l1526
																}
																position++
															}
														l1533:
															add(rulePegText, position1528)
														}
														{
															add(ruleAction70, position)
														}
														add(ruleExx, position1527)
													}
													goto l1471
												l1526:
													position, tokenIndex = position1471, tokenIndex1471
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1537 := position
																{
																	position1538 := position
																	{
																		position1539, tokenIndex1539 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1540
																		}
																		position++
																		goto l1539
																	l1540:
																		position, tokenIndex = position1539, tokenIndex1539
																		if buffer[position] != rune('E') {
																			goto l1469
																		}
																		position++
																	}
																l1539:
																	{
																		position1541, tokenIndex1541 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1542
																		}
																		position++
																		goto l1541
																	l1542:
																		position, tokenIndex = position1541, tokenIndex1541
																		if buffer[position] != rune('I') {
																			goto l1469
																		}
																		position++
																	}
																l1541:
																	add(rulePegText, position1538)
																}
																{
																	add(ruleAction72, position)
																}
																add(ruleEi, position1537)
															}
															break
														case 'D', 'd':
															{
																position1544 := position
																{
																	position1545 := position
																	{
																		position1546, tokenIndex1546 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1547
																		}
																		position++
																		goto l1546
																	l1547:
																		position, tokenIndex = position1546, tokenIndex1546
																		if buffer[position] != rune('D') {
																			goto l1469
																		}
																		position++
																	}
																l1546:
																	{
																		position1548, tokenIndex1548 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1549
																		}
																		position++
																		goto l1548
																	l1549:
																		position, tokenIndex = position1548, tokenIndex1548
																		if buffer[position] != rune('I') {
																			goto l1469
																		}
																		position++
																	}
																l1548:
																	add(rulePegText, position1545)
																}
																{
																	add(ruleAction71, position)
																}
																add(ruleDi, position1544)
															}
															break
														case 'C', 'c':
															{
																position1551 := position
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
																			goto l1469
																		}
																		position++
																	}
																l1553:
																	{
																		position1555, tokenIndex1555 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1556
																		}
																		position++
																		goto l1555
																	l1556:
																		position, tokenIndex = position1555, tokenIndex1555
																		if buffer[position] != rune('C') {
																			goto l1469
																		}
																		position++
																	}
																l1555:
																	{
																		position1557, tokenIndex1557 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1558
																		}
																		position++
																		goto l1557
																	l1558:
																		position, tokenIndex = position1557, tokenIndex1557
																		if buffer[position] != rune('F') {
																			goto l1469
																		}
																		position++
																	}
																l1557:
																	add(rulePegText, position1552)
																}
																{
																	add(ruleAction69, position)
																}
																add(ruleCcf, position1551)
															}
															break
														case 'S', 's':
															{
																position1560 := position
																{
																	position1561 := position
																	{
																		position1562, tokenIndex1562 := position, tokenIndex
																		if buffer[position] != rune('s') {
																			goto l1563
																		}
																		position++
																		goto l1562
																	l1563:
																		position, tokenIndex = position1562, tokenIndex1562
																		if buffer[position] != rune('S') {
																			goto l1469
																		}
																		position++
																	}
																l1562:
																	{
																		position1564, tokenIndex1564 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1565
																		}
																		position++
																		goto l1564
																	l1565:
																		position, tokenIndex = position1564, tokenIndex1564
																		if buffer[position] != rune('C') {
																			goto l1469
																		}
																		position++
																	}
																l1564:
																	{
																		position1566, tokenIndex1566 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1567
																		}
																		position++
																		goto l1566
																	l1567:
																		position, tokenIndex = position1566, tokenIndex1566
																		if buffer[position] != rune('F') {
																			goto l1469
																		}
																		position++
																	}
																l1566:
																	add(rulePegText, position1561)
																}
																{
																	add(ruleAction68, position)
																}
																add(ruleScf, position1560)
															}
															break
														case 'R', 'r':
															{
																position1569 := position
																{
																	position1570 := position
																	{
																		position1571, tokenIndex1571 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1572
																		}
																		position++
																		goto l1571
																	l1572:
																		position, tokenIndex = position1571, tokenIndex1571
																		if buffer[position] != rune('R') {
																			goto l1469
																		}
																		position++
																	}
																l1571:
																	{
																		position1573, tokenIndex1573 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1574
																		}
																		position++
																		goto l1573
																	l1574:
																		position, tokenIndex = position1573, tokenIndex1573
																		if buffer[position] != rune('R') {
																			goto l1469
																		}
																		position++
																	}
																l1573:
																	{
																		position1575, tokenIndex1575 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1576
																		}
																		position++
																		goto l1575
																	l1576:
																		position, tokenIndex = position1575, tokenIndex1575
																		if buffer[position] != rune('A') {
																			goto l1469
																		}
																		position++
																	}
																l1575:
																	add(rulePegText, position1570)
																}
																{
																	add(ruleAction65, position)
																}
																add(ruleRra, position1569)
															}
															break
														case 'H', 'h':
															{
																position1578 := position
																{
																	position1579 := position
																	{
																		position1580, tokenIndex1580 := position, tokenIndex
																		if buffer[position] != rune('h') {
																			goto l1581
																		}
																		position++
																		goto l1580
																	l1581:
																		position, tokenIndex = position1580, tokenIndex1580
																		if buffer[position] != rune('H') {
																			goto l1469
																		}
																		position++
																	}
																l1580:
																	{
																		position1582, tokenIndex1582 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1583
																		}
																		position++
																		goto l1582
																	l1583:
																		position, tokenIndex = position1582, tokenIndex1582
																		if buffer[position] != rune('A') {
																			goto l1469
																		}
																		position++
																	}
																l1582:
																	{
																		position1584, tokenIndex1584 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1585
																		}
																		position++
																		goto l1584
																	l1585:
																		position, tokenIndex = position1584, tokenIndex1584
																		if buffer[position] != rune('L') {
																			goto l1469
																		}
																		position++
																	}
																l1584:
																	{
																		position1586, tokenIndex1586 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1587
																		}
																		position++
																		goto l1586
																	l1587:
																		position, tokenIndex = position1586, tokenIndex1586
																		if buffer[position] != rune('T') {
																			goto l1469
																		}
																		position++
																	}
																l1586:
																	add(rulePegText, position1579)
																}
																{
																	add(ruleAction61, position)
																}
																add(ruleHalt, position1578)
															}
															break
														default:
															{
																position1589 := position
																{
																	position1590 := position
																	{
																		position1591, tokenIndex1591 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1592
																		}
																		position++
																		goto l1591
																	l1592:
																		position, tokenIndex = position1591, tokenIndex1591
																		if buffer[position] != rune('N') {
																			goto l1469
																		}
																		position++
																	}
																l1591:
																	{
																		position1593, tokenIndex1593 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1594
																		}
																		position++
																		goto l1593
																	l1594:
																		position, tokenIndex = position1593, tokenIndex1593
																		if buffer[position] != rune('O') {
																			goto l1469
																		}
																		position++
																	}
																l1593:
																	{
																		position1595, tokenIndex1595 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1596
																		}
																		position++
																		goto l1595
																	l1596:
																		position, tokenIndex = position1595, tokenIndex1595
																		if buffer[position] != rune('P') {
																			goto l1469
																		}
																		position++
																	}
																l1595:
																	add(rulePegText, position1590)
																}
																{
																	add(ruleAction60, position)
																}
																add(ruleNop, position1589)
															}
															break
														}
													}

												}
											l1471:
												add(ruleSimple, position1470)
											}
											goto l879
										l1469:
											position, tokenIndex = position879, tokenIndex879
											{
												position1599 := position
												{
													position1600, tokenIndex1600 := position, tokenIndex
													{
														position1602 := position
														{
															position1603, tokenIndex1603 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1604
															}
															position++
															goto l1603
														l1604:
															position, tokenIndex = position1603, tokenIndex1603
															if buffer[position] != rune('R') {
																goto l1601
															}
															position++
														}
													l1603:
														{
															position1605, tokenIndex1605 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1606
															}
															position++
															goto l1605
														l1606:
															position, tokenIndex = position1605, tokenIndex1605
															if buffer[position] != rune('S') {
																goto l1601
															}
															position++
														}
													l1605:
														{
															position1607, tokenIndex1607 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1608
															}
															position++
															goto l1607
														l1608:
															position, tokenIndex = position1607, tokenIndex1607
															if buffer[position] != rune('T') {
																goto l1601
															}
															position++
														}
													l1607:
														if !_rules[rulews]() {
															goto l1601
														}
														if !_rules[rulen]() {
															goto l1601
														}
														{
															add(ruleAction97, position)
														}
														add(ruleRst, position1602)
													}
													goto l1600
												l1601:
													position, tokenIndex = position1600, tokenIndex1600
													{
														position1611 := position
														{
															position1612, tokenIndex1612 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1613
															}
															position++
															goto l1612
														l1613:
															position, tokenIndex = position1612, tokenIndex1612
															if buffer[position] != rune('J') {
																goto l1610
															}
															position++
														}
													l1612:
														{
															position1614, tokenIndex1614 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l1615
															}
															position++
															goto l1614
														l1615:
															position, tokenIndex = position1614, tokenIndex1614
															if buffer[position] != rune('P') {
																goto l1610
															}
															position++
														}
													l1614:
														if !_rules[rulews]() {
															goto l1610
														}
														{
															position1616, tokenIndex1616 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1616
															}
															if !_rules[rulesep]() {
																goto l1616
															}
															goto l1617
														l1616:
															position, tokenIndex = position1616, tokenIndex1616
														}
													l1617:
														if !_rules[ruleSrc16]() {
															goto l1610
														}
														{
															add(ruleAction100, position)
														}
														add(ruleJp, position1611)
													}
													goto l1600
												l1610:
													position, tokenIndex = position1600, tokenIndex1600
													{
														switch buffer[position] {
														case 'D', 'd':
															{
																position1620 := position
																{
																	position1621, tokenIndex1621 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1622
																	}
																	position++
																	goto l1621
																l1622:
																	position, tokenIndex = position1621, tokenIndex1621
																	if buffer[position] != rune('D') {
																		goto l1598
																	}
																	position++
																}
															l1621:
																{
																	position1623, tokenIndex1623 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1624
																	}
																	position++
																	goto l1623
																l1624:
																	position, tokenIndex = position1623, tokenIndex1623
																	if buffer[position] != rune('J') {
																		goto l1598
																	}
																	position++
																}
															l1623:
																{
																	position1625, tokenIndex1625 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1626
																	}
																	position++
																	goto l1625
																l1626:
																	position, tokenIndex = position1625, tokenIndex1625
																	if buffer[position] != rune('N') {
																		goto l1598
																	}
																	position++
																}
															l1625:
																{
																	position1627, tokenIndex1627 := position, tokenIndex
																	if buffer[position] != rune('z') {
																		goto l1628
																	}
																	position++
																	goto l1627
																l1628:
																	position, tokenIndex = position1627, tokenIndex1627
																	if buffer[position] != rune('Z') {
																		goto l1598
																	}
																	position++
																}
															l1627:
																if !_rules[rulews]() {
																	goto l1598
																}
																if !_rules[ruledisp]() {
																	goto l1598
																}
																{
																	add(ruleAction102, position)
																}
																add(ruleDjnz, position1620)
															}
															break
														case 'J', 'j':
															{
																position1630 := position
																{
																	position1631, tokenIndex1631 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1632
																	}
																	position++
																	goto l1631
																l1632:
																	position, tokenIndex = position1631, tokenIndex1631
																	if buffer[position] != rune('J') {
																		goto l1598
																	}
																	position++
																}
															l1631:
																{
																	position1633, tokenIndex1633 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1634
																	}
																	position++
																	goto l1633
																l1634:
																	position, tokenIndex = position1633, tokenIndex1633
																	if buffer[position] != rune('R') {
																		goto l1598
																	}
																	position++
																}
															l1633:
																if !_rules[rulews]() {
																	goto l1598
																}
																{
																	position1635, tokenIndex1635 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1635
																	}
																	if !_rules[rulesep]() {
																		goto l1635
																	}
																	goto l1636
																l1635:
																	position, tokenIndex = position1635, tokenIndex1635
																}
															l1636:
																if !_rules[ruledisp]() {
																	goto l1598
																}
																{
																	add(ruleAction101, position)
																}
																add(ruleJr, position1630)
															}
															break
														case 'R', 'r':
															{
																position1638 := position
																{
																	position1639, tokenIndex1639 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1640
																	}
																	position++
																	goto l1639
																l1640:
																	position, tokenIndex = position1639, tokenIndex1639
																	if buffer[position] != rune('R') {
																		goto l1598
																	}
																	position++
																}
															l1639:
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
																		goto l1598
																	}
																	position++
																}
															l1641:
																{
																	position1643, tokenIndex1643 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1644
																	}
																	position++
																	goto l1643
																l1644:
																	position, tokenIndex = position1643, tokenIndex1643
																	if buffer[position] != rune('T') {
																		goto l1598
																	}
																	position++
																}
															l1643:
																{
																	position1645, tokenIndex1645 := position, tokenIndex
																	if !_rules[rulews]() {
																		goto l1645
																	}
																	if !_rules[rulecc]() {
																		goto l1645
																	}
																	goto l1646
																l1645:
																	position, tokenIndex = position1645, tokenIndex1645
																}
															l1646:
																{
																	add(ruleAction99, position)
																}
																add(ruleRet, position1638)
															}
															break
														default:
															{
																position1648 := position
																{
																	position1649, tokenIndex1649 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1650
																	}
																	position++
																	goto l1649
																l1650:
																	position, tokenIndex = position1649, tokenIndex1649
																	if buffer[position] != rune('C') {
																		goto l1598
																	}
																	position++
																}
															l1649:
																{
																	position1651, tokenIndex1651 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1652
																	}
																	position++
																	goto l1651
																l1652:
																	position, tokenIndex = position1651, tokenIndex1651
																	if buffer[position] != rune('A') {
																		goto l1598
																	}
																	position++
																}
															l1651:
																{
																	position1653, tokenIndex1653 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1654
																	}
																	position++
																	goto l1653
																l1654:
																	position, tokenIndex = position1653, tokenIndex1653
																	if buffer[position] != rune('L') {
																		goto l1598
																	}
																	position++
																}
															l1653:
																{
																	position1655, tokenIndex1655 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1656
																	}
																	position++
																	goto l1655
																l1656:
																	position, tokenIndex = position1655, tokenIndex1655
																	if buffer[position] != rune('L') {
																		goto l1598
																	}
																	position++
																}
															l1655:
																if !_rules[rulews]() {
																	goto l1598
																}
																{
																	position1657, tokenIndex1657 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1657
																	}
																	if !_rules[rulesep]() {
																		goto l1657
																	}
																	goto l1658
																l1657:
																	position, tokenIndex = position1657, tokenIndex1657
																}
															l1658:
																if !_rules[ruleSrc16]() {
																	goto l1598
																}
																{
																	add(ruleAction98, position)
																}
																add(ruleCall, position1648)
															}
															break
														}
													}

												}
											l1600:
												add(ruleJump, position1599)
											}
											goto l879
										l1598:
											position, tokenIndex = position879, tokenIndex879
											{
												position1660 := position
												{
													position1661, tokenIndex1661 := position, tokenIndex
													{
														position1663 := position
														{
															position1664, tokenIndex1664 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1665
															}
															position++
															goto l1664
														l1665:
															position, tokenIndex = position1664, tokenIndex1664
															if buffer[position] != rune('I') {
																goto l1662
															}
															position++
														}
													l1664:
														{
															position1666, tokenIndex1666 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1667
															}
															position++
															goto l1666
														l1667:
															position, tokenIndex = position1666, tokenIndex1666
															if buffer[position] != rune('N') {
																goto l1662
															}
															position++
														}
													l1666:
														if !_rules[rulews]() {
															goto l1662
														}
														if !_rules[ruleReg8]() {
															goto l1662
														}
														if !_rules[rulesep]() {
															goto l1662
														}
														if !_rules[rulePort]() {
															goto l1662
														}
														{
															add(ruleAction103, position)
														}
														add(ruleIN, position1663)
													}
													goto l1661
												l1662:
													position, tokenIndex = position1661, tokenIndex1661
													{
														position1669 := position
														{
															position1670, tokenIndex1670 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1671
															}
															position++
															goto l1670
														l1671:
															position, tokenIndex = position1670, tokenIndex1670
															if buffer[position] != rune('O') {
																goto l858
															}
															position++
														}
													l1670:
														{
															position1672, tokenIndex1672 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1673
															}
															position++
															goto l1672
														l1673:
															position, tokenIndex = position1672, tokenIndex1672
															if buffer[position] != rune('U') {
																goto l858
															}
															position++
														}
													l1672:
														{
															position1674, tokenIndex1674 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1675
															}
															position++
															goto l1674
														l1675:
															position, tokenIndex = position1674, tokenIndex1674
															if buffer[position] != rune('T') {
																goto l858
															}
															position++
														}
													l1674:
														if !_rules[rulews]() {
															goto l858
														}
														if !_rules[rulePort]() {
															goto l858
														}
														if !_rules[rulesep]() {
															goto l858
														}
														if !_rules[ruleReg8]() {
															goto l858
														}
														{
															add(ruleAction104, position)
														}
														add(ruleOUT, position1669)
													}
												}
											l1661:
												add(ruleIO, position1660)
											}
										}
									l879:
										add(ruleInstruction, position878)
									}
								}
							l861:
								add(ruleStatement, position860)
							}
							goto l859
						l858:
							position, tokenIndex = position858, tokenIndex858
						}
					l859:
						{
							position1677, tokenIndex1677 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1677
							}
							goto l1678
						l1677:
							position, tokenIndex = position1677, tokenIndex1677
						}
					l1678:
						{
							position1679, tokenIndex1679 := position, tokenIndex
							{
								position1681 := position
								{
									position1682, tokenIndex1682 := position, tokenIndex
									if buffer[position] != rune(';') {
										goto l1683
									}
									position++
									goto l1682
								l1683:
									position, tokenIndex = position1682, tokenIndex1682
									if buffer[position] != rune('#') {
										goto l1679
									}
									position++
								}
							l1682:
							l1684:
								{
									position1685, tokenIndex1685 := position, tokenIndex
									{
										position1686, tokenIndex1686 := position, tokenIndex
										if buffer[position] != rune('\n') {
											goto l1686
										}
										position++
										goto l1685
									l1686:
										position, tokenIndex = position1686, tokenIndex1686
									}
									if !matchDot() {
										goto l1685
									}
									goto l1684
								l1685:
									position, tokenIndex = position1685, tokenIndex1685
								}
								add(ruleComment, position1681)
							}
							goto l1680
						l1679:
							position, tokenIndex = position1679, tokenIndex1679
						}
					l1680:
						{
							position1687, tokenIndex1687 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1687
							}
							goto l1688
						l1687:
							position, tokenIndex = position1687, tokenIndex1687
						}
					l1688:
						{
							position1689, tokenIndex1689 := position, tokenIndex
							{
								position1691, tokenIndex1691 := position, tokenIndex
								if buffer[position] != rune('\r') {
									goto l1691
								}
								position++
								goto l1692
							l1691:
								position, tokenIndex = position1691, tokenIndex1691
							}
						l1692:
							if buffer[position] != rune('\n') {
								goto l1690
							}
							position++
							goto l1689
						l1690:
							position, tokenIndex = position1689, tokenIndex1689
							if buffer[position] != rune(':') {
								goto l3
							}
							position++
						}
					l1689:
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position849)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position1694, tokenIndex1694 := position, tokenIndex
					if !matchDot() {
						goto l1694
					}
					goto l0
				l1694:
					position, tokenIndex = position1694, tokenIndex1694
				}
				add(ruleProgram, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 Line <- <(ws* LabelDefn? ws* Statement? ws? Comment? ws? (('\r'? '\n') / ':') Action0)> */
		nil,
		/* 2 Statement <- <(Directive / Instruction)> */
		nil,
		/* 3 Directive <- <((&('a') Aseg) | (&('.') Title) | (&('O' | 'o') Org))> */
		nil,
		/* 4 Title <- <('.' 't' 'i' 't' 'l' 'e' ws '\'' (!'\'' .)* '\'')> */
		nil,
		/* 5 Aseg <- <('a' 's' 'e' 'g')> */
		nil,
		/* 6 Org <- <(('o' / 'O') ('r' / 'R') ('g' / 'G') ws nn Action1)> */
		nil,
		/* 7 LabelDefn <- <(LabelText ':' ws Action2)> */
		nil,
		/* 8 LabelText <- <<(alpha alphanum*)>> */
		func() bool {
			position1702, tokenIndex1702 := position, tokenIndex
			{
				position1703 := position
				{
					position1704 := position
					if !_rules[rulealpha]() {
						goto l1702
					}
				l1705:
					{
						position1706, tokenIndex1706 := position, tokenIndex
						{
							position1707 := position
							{
								position1708, tokenIndex1708 := position, tokenIndex
								if !_rules[rulealpha]() {
									goto l1709
								}
								goto l1708
							l1709:
								position, tokenIndex = position1708, tokenIndex1708
								{
									position1710 := position
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1706
									}
									position++
									add(rulenum, position1710)
								}
							}
						l1708:
							add(rulealphanum, position1707)
						}
						goto l1705
					l1706:
						position, tokenIndex = position1706, tokenIndex1706
					}
					add(rulePegText, position1704)
				}
				add(ruleLabelText, position1703)
			}
			return true
		l1702:
			position, tokenIndex = position1702, tokenIndex1702
			return false
		},
		/* 9 alphanum <- <(alpha / num)> */
		nil,
		/* 10 alpha <- <([a-z] / [A-Z])> */
		func() bool {
			position1712, tokenIndex1712 := position, tokenIndex
			{
				position1713 := position
				{
					position1714, tokenIndex1714 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1715
					}
					position++
					goto l1714
				l1715:
					position, tokenIndex = position1714, tokenIndex1714
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1712
					}
					position++
				}
			l1714:
				add(rulealpha, position1713)
			}
			return true
		l1712:
			position, tokenIndex = position1712, tokenIndex1712
			return false
		},
		/* 11 num <- <[0-9]> */
		nil,
		/* 12 Comment <- <((';' / '#') (!'\n' .)*)> */
		nil,
		/* 13 Instruction <- <(Assignment / Inc / Dec / Alu16 / Alu / BitOp / EDSimple / Simple / Jump / IO)> */
		nil,
		/* 14 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 15 Load <- <(Load16 / Load8)> */
		nil,
		/* 16 Load8 <- <(('l' / 'L') ('d' / 'D') ws Dst8 sep Src8 Action3)> */
		nil,
		/* 17 Load16 <- <(('l' / 'L') ('d' / 'D') ws Dst16 sep Src16 Action4)> */
		nil,
		/* 18 Push <- <(('p' / 'P') ('u' / 'U') ('s' / 'S') ('h' / 'H') ws Src16 Action5)> */
		nil,
		/* 19 Pop <- <(('p' / 'P') ('o' / 'O') ('p' / 'P') ws Dst16 Action6)> */
		nil,
		/* 20 Ex <- <(('e' / 'E') ('x' / 'X') ws Dst16 sep Src16 Action7)> */
		nil,
		/* 21 Inc <- <(Inc16Indexed8 / Inc16 / Inc8)> */
		nil,
		/* 22 Inc16Indexed8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws ILoc8 Action8)> */
		nil,
		/* 23 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action9)> */
		nil,
		/* 24 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action10)> */
		nil,
		/* 25 Dec <- <(Dec16Indexed8 / Dec16 / Dec8)> */
		nil,
		/* 26 Dec16Indexed8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws ILoc8 Action11)> */
		nil,
		/* 27 Dec8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc8 Action12)> */
		nil,
		/* 28 Dec16 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc16 Action13)> */
		nil,
		/* 29 Alu16 <- <(Add16 / Adc16 / Sbc16)> */
		nil,
		/* 30 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action14)> */
		nil,
		/* 31 Adc16 <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws Dst16 sep Src16 Action15)> */
		nil,
		/* 32 Sbc16 <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws Dst16 sep Src16 Action16)> */
		nil,
		/* 33 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action17)> */
		nil,
		/* 34 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action18)> */
		func() bool {
			position1739, tokenIndex1739 := position, tokenIndex
			{
				position1740 := position
				{
					position1741, tokenIndex1741 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1742
					}
					goto l1741
				l1742:
					position, tokenIndex = position1741, tokenIndex1741
					if !_rules[ruleReg8]() {
						goto l1743
					}
					goto l1741
				l1743:
					position, tokenIndex = position1741, tokenIndex1741
					if !_rules[ruleReg16Contents]() {
						goto l1744
					}
					goto l1741
				l1744:
					position, tokenIndex = position1741, tokenIndex1741
					if !_rules[rulenn_contents]() {
						goto l1739
					}
				}
			l1741:
				{
					add(ruleAction18, position)
				}
				add(ruleSrc8, position1740)
			}
			return true
		l1739:
			position, tokenIndex = position1739, tokenIndex1739
			return false
		},
		/* 35 Loc8 <- <((Reg8 / Reg16Contents) Action19)> */
		func() bool {
			position1746, tokenIndex1746 := position, tokenIndex
			{
				position1747 := position
				{
					position1748, tokenIndex1748 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1749
					}
					goto l1748
				l1749:
					position, tokenIndex = position1748, tokenIndex1748
					if !_rules[ruleReg16Contents]() {
						goto l1746
					}
				}
			l1748:
				{
					add(ruleAction19, position)
				}
				add(ruleLoc8, position1747)
			}
			return true
		l1746:
			position, tokenIndex = position1746, tokenIndex1746
			return false
		},
		/* 36 Copy8 <- <(Reg8 Action20)> */
		func() bool {
			position1751, tokenIndex1751 := position, tokenIndex
			{
				position1752 := position
				if !_rules[ruleReg8]() {
					goto l1751
				}
				{
					add(ruleAction20, position)
				}
				add(ruleCopy8, position1752)
			}
			return true
		l1751:
			position, tokenIndex = position1751, tokenIndex1751
			return false
		},
		/* 37 ILoc8 <- <(IReg8 Action21)> */
		func() bool {
			position1754, tokenIndex1754 := position, tokenIndex
			{
				position1755 := position
				if !_rules[ruleIReg8]() {
					goto l1754
				}
				{
					add(ruleAction21, position)
				}
				add(ruleILoc8, position1755)
			}
			return true
		l1754:
			position, tokenIndex = position1754, tokenIndex1754
			return false
		},
		/* 38 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action22)> */
		func() bool {
			position1757, tokenIndex1757 := position, tokenIndex
			{
				position1758 := position
				{
					position1759 := position
					{
						position1760, tokenIndex1760 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1761
						}
						goto l1760
					l1761:
						position, tokenIndex = position1760, tokenIndex1760
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1763 := position
									{
										position1764, tokenIndex1764 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1765
										}
										position++
										goto l1764
									l1765:
										position, tokenIndex = position1764, tokenIndex1764
										if buffer[position] != rune('R') {
											goto l1757
										}
										position++
									}
								l1764:
									add(ruleR, position1763)
								}
								break
							case 'I', 'i':
								{
									position1766 := position
									{
										position1767, tokenIndex1767 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1768
										}
										position++
										goto l1767
									l1768:
										position, tokenIndex = position1767, tokenIndex1767
										if buffer[position] != rune('I') {
											goto l1757
										}
										position++
									}
								l1767:
									add(ruleI, position1766)
								}
								break
							case 'L', 'l':
								{
									position1769 := position
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
											goto l1757
										}
										position++
									}
								l1770:
									add(ruleL, position1769)
								}
								break
							case 'H', 'h':
								{
									position1772 := position
									{
										position1773, tokenIndex1773 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1774
										}
										position++
										goto l1773
									l1774:
										position, tokenIndex = position1773, tokenIndex1773
										if buffer[position] != rune('H') {
											goto l1757
										}
										position++
									}
								l1773:
									add(ruleH, position1772)
								}
								break
							case 'E', 'e':
								{
									position1775 := position
									{
										position1776, tokenIndex1776 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1777
										}
										position++
										goto l1776
									l1777:
										position, tokenIndex = position1776, tokenIndex1776
										if buffer[position] != rune('E') {
											goto l1757
										}
										position++
									}
								l1776:
									add(ruleE, position1775)
								}
								break
							case 'D', 'd':
								{
									position1778 := position
									{
										position1779, tokenIndex1779 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1780
										}
										position++
										goto l1779
									l1780:
										position, tokenIndex = position1779, tokenIndex1779
										if buffer[position] != rune('D') {
											goto l1757
										}
										position++
									}
								l1779:
									add(ruleD, position1778)
								}
								break
							case 'C', 'c':
								{
									position1781 := position
									{
										position1782, tokenIndex1782 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1783
										}
										position++
										goto l1782
									l1783:
										position, tokenIndex = position1782, tokenIndex1782
										if buffer[position] != rune('C') {
											goto l1757
										}
										position++
									}
								l1782:
									add(ruleC, position1781)
								}
								break
							case 'B', 'b':
								{
									position1784 := position
									{
										position1785, tokenIndex1785 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1786
										}
										position++
										goto l1785
									l1786:
										position, tokenIndex = position1785, tokenIndex1785
										if buffer[position] != rune('B') {
											goto l1757
										}
										position++
									}
								l1785:
									add(ruleB, position1784)
								}
								break
							case 'F', 'f':
								{
									position1787 := position
									{
										position1788, tokenIndex1788 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1789
										}
										position++
										goto l1788
									l1789:
										position, tokenIndex = position1788, tokenIndex1788
										if buffer[position] != rune('F') {
											goto l1757
										}
										position++
									}
								l1788:
									add(ruleF, position1787)
								}
								break
							default:
								{
									position1790 := position
									{
										position1791, tokenIndex1791 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1792
										}
										position++
										goto l1791
									l1792:
										position, tokenIndex = position1791, tokenIndex1791
										if buffer[position] != rune('A') {
											goto l1757
										}
										position++
									}
								l1791:
									add(ruleA, position1790)
								}
								break
							}
						}

					}
				l1760:
					add(rulePegText, position1759)
				}
				{
					add(ruleAction22, position)
				}
				add(ruleReg8, position1758)
			}
			return true
		l1757:
			position, tokenIndex = position1757, tokenIndex1757
			return false
		},
		/* 39 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action23)> */
		func() bool {
			position1794, tokenIndex1794 := position, tokenIndex
			{
				position1795 := position
				{
					position1796 := position
					{
						position1797, tokenIndex1797 := position, tokenIndex
						{
							position1799 := position
							{
								position1800, tokenIndex1800 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1801
								}
								position++
								goto l1800
							l1801:
								position, tokenIndex = position1800, tokenIndex1800
								if buffer[position] != rune('I') {
									goto l1798
								}
								position++
							}
						l1800:
							{
								position1802, tokenIndex1802 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1803
								}
								position++
								goto l1802
							l1803:
								position, tokenIndex = position1802, tokenIndex1802
								if buffer[position] != rune('X') {
									goto l1798
								}
								position++
							}
						l1802:
							{
								position1804, tokenIndex1804 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1805
								}
								position++
								goto l1804
							l1805:
								position, tokenIndex = position1804, tokenIndex1804
								if buffer[position] != rune('H') {
									goto l1798
								}
								position++
							}
						l1804:
							add(ruleIXH, position1799)
						}
						goto l1797
					l1798:
						position, tokenIndex = position1797, tokenIndex1797
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
							{
								position1812, tokenIndex1812 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1813
								}
								position++
								goto l1812
							l1813:
								position, tokenIndex = position1812, tokenIndex1812
								if buffer[position] != rune('L') {
									goto l1806
								}
								position++
							}
						l1812:
							add(ruleIXL, position1807)
						}
						goto l1797
					l1806:
						position, tokenIndex = position1797, tokenIndex1797
						{
							position1815 := position
							{
								position1816, tokenIndex1816 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1817
								}
								position++
								goto l1816
							l1817:
								position, tokenIndex = position1816, tokenIndex1816
								if buffer[position] != rune('I') {
									goto l1814
								}
								position++
							}
						l1816:
							{
								position1818, tokenIndex1818 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1819
								}
								position++
								goto l1818
							l1819:
								position, tokenIndex = position1818, tokenIndex1818
								if buffer[position] != rune('Y') {
									goto l1814
								}
								position++
							}
						l1818:
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
									goto l1814
								}
								position++
							}
						l1820:
							add(ruleIYH, position1815)
						}
						goto l1797
					l1814:
						position, tokenIndex = position1797, tokenIndex1797
						{
							position1822 := position
							{
								position1823, tokenIndex1823 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1824
								}
								position++
								goto l1823
							l1824:
								position, tokenIndex = position1823, tokenIndex1823
								if buffer[position] != rune('I') {
									goto l1794
								}
								position++
							}
						l1823:
							{
								position1825, tokenIndex1825 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1826
								}
								position++
								goto l1825
							l1826:
								position, tokenIndex = position1825, tokenIndex1825
								if buffer[position] != rune('Y') {
									goto l1794
								}
								position++
							}
						l1825:
							{
								position1827, tokenIndex1827 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1828
								}
								position++
								goto l1827
							l1828:
								position, tokenIndex = position1827, tokenIndex1827
								if buffer[position] != rune('L') {
									goto l1794
								}
								position++
							}
						l1827:
							add(ruleIYL, position1822)
						}
					}
				l1797:
					add(rulePegText, position1796)
				}
				{
					add(ruleAction23, position)
				}
				add(ruleIReg8, position1795)
			}
			return true
		l1794:
			position, tokenIndex = position1794, tokenIndex1794
			return false
		},
		/* 40 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action24)> */
		func() bool {
			position1830, tokenIndex1830 := position, tokenIndex
			{
				position1831 := position
				{
					position1832, tokenIndex1832 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1833
					}
					goto l1832
				l1833:
					position, tokenIndex = position1832, tokenIndex1832
					if !_rules[rulenn_contents]() {
						goto l1834
					}
					goto l1832
				l1834:
					position, tokenIndex = position1832, tokenIndex1832
					if !_rules[ruleReg16Contents]() {
						goto l1830
					}
				}
			l1832:
				{
					add(ruleAction24, position)
				}
				add(ruleDst16, position1831)
			}
			return true
		l1830:
			position, tokenIndex = position1830, tokenIndex1830
			return false
		},
		/* 41 Src16 <- <((Reg16 / nn / nn_contents) Action25)> */
		func() bool {
			position1836, tokenIndex1836 := position, tokenIndex
			{
				position1837 := position
				{
					position1838, tokenIndex1838 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1839
					}
					goto l1838
				l1839:
					position, tokenIndex = position1838, tokenIndex1838
					if !_rules[rulenn]() {
						goto l1840
					}
					goto l1838
				l1840:
					position, tokenIndex = position1838, tokenIndex1838
					if !_rules[rulenn_contents]() {
						goto l1836
					}
				}
			l1838:
				{
					add(ruleAction25, position)
				}
				add(ruleSrc16, position1837)
			}
			return true
		l1836:
			position, tokenIndex = position1836, tokenIndex1836
			return false
		},
		/* 42 Loc16 <- <(Reg16 Action26)> */
		func() bool {
			position1842, tokenIndex1842 := position, tokenIndex
			{
				position1843 := position
				if !_rules[ruleReg16]() {
					goto l1842
				}
				{
					add(ruleAction26, position)
				}
				add(ruleLoc16, position1843)
			}
			return true
		l1842:
			position, tokenIndex = position1842, tokenIndex1842
			return false
		},
		/* 43 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action27)> */
		func() bool {
			position1845, tokenIndex1845 := position, tokenIndex
			{
				position1846 := position
				{
					position1847 := position
					{
						position1848, tokenIndex1848 := position, tokenIndex
						{
							position1850 := position
							{
								position1851, tokenIndex1851 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1852
								}
								position++
								goto l1851
							l1852:
								position, tokenIndex = position1851, tokenIndex1851
								if buffer[position] != rune('A') {
									goto l1849
								}
								position++
							}
						l1851:
							{
								position1853, tokenIndex1853 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1854
								}
								position++
								goto l1853
							l1854:
								position, tokenIndex = position1853, tokenIndex1853
								if buffer[position] != rune('F') {
									goto l1849
								}
								position++
							}
						l1853:
							if buffer[position] != rune('\'') {
								goto l1849
							}
							position++
							add(ruleAF_PRIME, position1850)
						}
						goto l1848
					l1849:
						position, tokenIndex = position1848, tokenIndex1848
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1845
								}
								break
							case 'S', 's':
								{
									position1856 := position
									{
										position1857, tokenIndex1857 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1858
										}
										position++
										goto l1857
									l1858:
										position, tokenIndex = position1857, tokenIndex1857
										if buffer[position] != rune('S') {
											goto l1845
										}
										position++
									}
								l1857:
									{
										position1859, tokenIndex1859 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1860
										}
										position++
										goto l1859
									l1860:
										position, tokenIndex = position1859, tokenIndex1859
										if buffer[position] != rune('P') {
											goto l1845
										}
										position++
									}
								l1859:
									add(ruleSP, position1856)
								}
								break
							case 'H', 'h':
								{
									position1861 := position
									{
										position1862, tokenIndex1862 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1863
										}
										position++
										goto l1862
									l1863:
										position, tokenIndex = position1862, tokenIndex1862
										if buffer[position] != rune('H') {
											goto l1845
										}
										position++
									}
								l1862:
									{
										position1864, tokenIndex1864 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1865
										}
										position++
										goto l1864
									l1865:
										position, tokenIndex = position1864, tokenIndex1864
										if buffer[position] != rune('L') {
											goto l1845
										}
										position++
									}
								l1864:
									add(ruleHL, position1861)
								}
								break
							case 'D', 'd':
								{
									position1866 := position
									{
										position1867, tokenIndex1867 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1868
										}
										position++
										goto l1867
									l1868:
										position, tokenIndex = position1867, tokenIndex1867
										if buffer[position] != rune('D') {
											goto l1845
										}
										position++
									}
								l1867:
									{
										position1869, tokenIndex1869 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1870
										}
										position++
										goto l1869
									l1870:
										position, tokenIndex = position1869, tokenIndex1869
										if buffer[position] != rune('E') {
											goto l1845
										}
										position++
									}
								l1869:
									add(ruleDE, position1866)
								}
								break
							case 'B', 'b':
								{
									position1871 := position
									{
										position1872, tokenIndex1872 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1873
										}
										position++
										goto l1872
									l1873:
										position, tokenIndex = position1872, tokenIndex1872
										if buffer[position] != rune('B') {
											goto l1845
										}
										position++
									}
								l1872:
									{
										position1874, tokenIndex1874 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1875
										}
										position++
										goto l1874
									l1875:
										position, tokenIndex = position1874, tokenIndex1874
										if buffer[position] != rune('C') {
											goto l1845
										}
										position++
									}
								l1874:
									add(ruleBC, position1871)
								}
								break
							default:
								{
									position1876 := position
									{
										position1877, tokenIndex1877 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1878
										}
										position++
										goto l1877
									l1878:
										position, tokenIndex = position1877, tokenIndex1877
										if buffer[position] != rune('A') {
											goto l1845
										}
										position++
									}
								l1877:
									{
										position1879, tokenIndex1879 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1880
										}
										position++
										goto l1879
									l1880:
										position, tokenIndex = position1879, tokenIndex1879
										if buffer[position] != rune('F') {
											goto l1845
										}
										position++
									}
								l1879:
									add(ruleAF, position1876)
								}
								break
							}
						}

					}
				l1848:
					add(rulePegText, position1847)
				}
				{
					add(ruleAction27, position)
				}
				add(ruleReg16, position1846)
			}
			return true
		l1845:
			position, tokenIndex = position1845, tokenIndex1845
			return false
		},
		/* 44 IReg16 <- <(<(IX / IY)> Action28)> */
		func() bool {
			position1882, tokenIndex1882 := position, tokenIndex
			{
				position1883 := position
				{
					position1884 := position
					{
						position1885, tokenIndex1885 := position, tokenIndex
						{
							position1887 := position
							{
								position1888, tokenIndex1888 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1889
								}
								position++
								goto l1888
							l1889:
								position, tokenIndex = position1888, tokenIndex1888
								if buffer[position] != rune('I') {
									goto l1886
								}
								position++
							}
						l1888:
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
									goto l1886
								}
								position++
							}
						l1890:
							add(ruleIX, position1887)
						}
						goto l1885
					l1886:
						position, tokenIndex = position1885, tokenIndex1885
						{
							position1892 := position
							{
								position1893, tokenIndex1893 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1894
								}
								position++
								goto l1893
							l1894:
								position, tokenIndex = position1893, tokenIndex1893
								if buffer[position] != rune('I') {
									goto l1882
								}
								position++
							}
						l1893:
							{
								position1895, tokenIndex1895 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1896
								}
								position++
								goto l1895
							l1896:
								position, tokenIndex = position1895, tokenIndex1895
								if buffer[position] != rune('Y') {
									goto l1882
								}
								position++
							}
						l1895:
							add(ruleIY, position1892)
						}
					}
				l1885:
					add(rulePegText, position1884)
				}
				{
					add(ruleAction28, position)
				}
				add(ruleIReg16, position1883)
			}
			return true
		l1882:
			position, tokenIndex = position1882, tokenIndex1882
			return false
		},
		/* 45 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1898, tokenIndex1898 := position, tokenIndex
			{
				position1899 := position
				{
					position1900, tokenIndex1900 := position, tokenIndex
					{
						position1902 := position
						if buffer[position] != rune('(') {
							goto l1901
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1901
						}
						{
							position1903, tokenIndex1903 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1903
							}
							goto l1904
						l1903:
							position, tokenIndex = position1903, tokenIndex1903
						}
					l1904:
						if !_rules[ruledisp]() {
							goto l1901
						}
						{
							position1905, tokenIndex1905 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1905
							}
							goto l1906
						l1905:
							position, tokenIndex = position1905, tokenIndex1905
						}
					l1906:
						if buffer[position] != rune(')') {
							goto l1901
						}
						position++
						{
							add(ruleAction30, position)
						}
						add(ruleIndexedR16C, position1902)
					}
					goto l1900
				l1901:
					position, tokenIndex = position1900, tokenIndex1900
					{
						position1908 := position
						if buffer[position] != rune('(') {
							goto l1898
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1898
						}
						if buffer[position] != rune(')') {
							goto l1898
						}
						position++
						{
							add(ruleAction29, position)
						}
						add(rulePlainR16C, position1908)
					}
				}
			l1900:
				add(ruleReg16Contents, position1899)
			}
			return true
		l1898:
			position, tokenIndex = position1898, tokenIndex1898
			return false
		},
		/* 46 PlainR16C <- <('(' Reg16 ')' Action29)> */
		nil,
		/* 47 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action30)> */
		nil,
		/* 48 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1912, tokenIndex1912 := position, tokenIndex
			{
				position1913 := position
				{
					position1914, tokenIndex1914 := position, tokenIndex
					{
						position1916 := position
						{
							position1917 := position
							if !_rules[rulehexdigit]() {
								goto l1915
							}
							if !_rules[rulehexdigit]() {
								goto l1915
							}
							add(rulePegText, position1917)
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
								goto l1915
							}
							position++
						}
					l1918:
						{
							add(ruleAction34, position)
						}
						add(rulehexByteH, position1916)
					}
					goto l1914
				l1915:
					position, tokenIndex = position1914, tokenIndex1914
					{
						position1922 := position
						if buffer[position] != rune('0') {
							goto l1921
						}
						position++
						{
							position1923, tokenIndex1923 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1924
							}
							position++
							goto l1923
						l1924:
							position, tokenIndex = position1923, tokenIndex1923
							if buffer[position] != rune('X') {
								goto l1921
							}
							position++
						}
					l1923:
						{
							position1925 := position
							if !_rules[rulehexdigit]() {
								goto l1921
							}
							if !_rules[rulehexdigit]() {
								goto l1921
							}
							add(rulePegText, position1925)
						}
						{
							add(ruleAction35, position)
						}
						add(rulehexByte0x, position1922)
					}
					goto l1914
				l1921:
					position, tokenIndex = position1914, tokenIndex1914
					{
						position1927 := position
						{
							position1928 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1912
							}
							position++
						l1929:
							{
								position1930, tokenIndex1930 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1930
								}
								position++
								goto l1929
							l1930:
								position, tokenIndex = position1930, tokenIndex1930
							}
							add(rulePegText, position1928)
						}
						{
							add(ruleAction36, position)
						}
						add(ruledecimalByte, position1927)
					}
				}
			l1914:
				add(rulen, position1913)
			}
			return true
		l1912:
			position, tokenIndex = position1912, tokenIndex1912
			return false
		},
		/* 49 nn <- <(LabelNN / hexWordH / hexWord0x)> */
		func() bool {
			position1932, tokenIndex1932 := position, tokenIndex
			{
				position1933 := position
				{
					position1934, tokenIndex1934 := position, tokenIndex
					{
						position1936 := position
						{
							position1937 := position
							if !_rules[ruleLabelText]() {
								goto l1935
							}
							add(rulePegText, position1937)
						}
						{
							add(ruleAction37, position)
						}
						add(ruleLabelNN, position1936)
					}
					goto l1934
				l1935:
					position, tokenIndex = position1934, tokenIndex1934
					{
						position1940 := position
						{
							position1941 := position
							if !_rules[rulehexdigit]() {
								goto l1939
							}
							if !_rules[rulehexdigit]() {
								goto l1939
							}
							if !_rules[rulehexdigit]() {
								goto l1939
							}
							if !_rules[rulehexdigit]() {
								goto l1939
							}
							add(rulePegText, position1941)
						}
						{
							position1942, tokenIndex1942 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1943
							}
							position++
							goto l1942
						l1943:
							position, tokenIndex = position1942, tokenIndex1942
							if buffer[position] != rune('H') {
								goto l1939
							}
							position++
						}
					l1942:
						{
							add(ruleAction38, position)
						}
						add(rulehexWordH, position1940)
					}
					goto l1934
				l1939:
					position, tokenIndex = position1934, tokenIndex1934
					{
						position1945 := position
						if buffer[position] != rune('0') {
							goto l1932
						}
						position++
						{
							position1946, tokenIndex1946 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1947
							}
							position++
							goto l1946
						l1947:
							position, tokenIndex = position1946, tokenIndex1946
							if buffer[position] != rune('X') {
								goto l1932
							}
							position++
						}
					l1946:
						{
							position1948 := position
							if !_rules[rulehexdigit]() {
								goto l1932
							}
							if !_rules[rulehexdigit]() {
								goto l1932
							}
							if !_rules[rulehexdigit]() {
								goto l1932
							}
							if !_rules[rulehexdigit]() {
								goto l1932
							}
							add(rulePegText, position1948)
						}
						{
							add(ruleAction39, position)
						}
						add(rulehexWord0x, position1945)
					}
				}
			l1934:
				add(rulenn, position1933)
			}
			return true
		l1932:
			position, tokenIndex = position1932, tokenIndex1932
			return false
		},
		/* 50 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position1950, tokenIndex1950 := position, tokenIndex
			{
				position1951 := position
				{
					position1952, tokenIndex1952 := position, tokenIndex
					{
						position1954 := position
						{
							position1955 := position
							{
								position1956, tokenIndex1956 := position, tokenIndex
								{
									position1958, tokenIndex1958 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1959
									}
									position++
									goto l1958
								l1959:
									position, tokenIndex = position1958, tokenIndex1958
									if buffer[position] != rune('+') {
										goto l1956
									}
									position++
								}
							l1958:
								goto l1957
							l1956:
								position, tokenIndex = position1956, tokenIndex1956
							}
						l1957:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1953
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1953
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1953
									}
									position++
									break
								}
							}

						l1960:
							{
								position1961, tokenIndex1961 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1961
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1961
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1961
										}
										position++
										break
									}
								}

								goto l1960
							l1961:
								position, tokenIndex = position1961, tokenIndex1961
							}
							add(rulePegText, position1955)
						}
						{
							position1964, tokenIndex1964 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1965
							}
							position++
							goto l1964
						l1965:
							position, tokenIndex = position1964, tokenIndex1964
							if buffer[position] != rune('H') {
								goto l1953
							}
							position++
						}
					l1964:
						{
							add(ruleAction32, position)
						}
						add(rulesignedHexByteH, position1954)
					}
					goto l1952
				l1953:
					position, tokenIndex = position1952, tokenIndex1952
					{
						position1968 := position
						{
							position1969 := position
							{
								position1970, tokenIndex1970 := position, tokenIndex
								{
									position1972, tokenIndex1972 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1973
									}
									position++
									goto l1972
								l1973:
									position, tokenIndex = position1972, tokenIndex1972
									if buffer[position] != rune('+') {
										goto l1970
									}
									position++
								}
							l1972:
								goto l1971
							l1970:
								position, tokenIndex = position1970, tokenIndex1970
							}
						l1971:
							if buffer[position] != rune('0') {
								goto l1967
							}
							position++
							{
								position1974, tokenIndex1974 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1975
								}
								position++
								goto l1974
							l1975:
								position, tokenIndex = position1974, tokenIndex1974
								if buffer[position] != rune('X') {
									goto l1967
								}
								position++
							}
						l1974:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1967
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1967
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1967
									}
									position++
									break
								}
							}

						l1976:
							{
								position1977, tokenIndex1977 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1977
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1977
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1977
										}
										position++
										break
									}
								}

								goto l1976
							l1977:
								position, tokenIndex = position1977, tokenIndex1977
							}
							add(rulePegText, position1969)
						}
						{
							add(ruleAction33, position)
						}
						add(rulesignedHexByte0x, position1968)
					}
					goto l1952
				l1967:
					position, tokenIndex = position1952, tokenIndex1952
					{
						position1981 := position
						{
							position1982 := position
							{
								position1983, tokenIndex1983 := position, tokenIndex
								{
									position1985, tokenIndex1985 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1986
									}
									position++
									goto l1985
								l1986:
									position, tokenIndex = position1985, tokenIndex1985
									if buffer[position] != rune('+') {
										goto l1983
									}
									position++
								}
							l1985:
								goto l1984
							l1983:
								position, tokenIndex = position1983, tokenIndex1983
							}
						l1984:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1950
							}
							position++
						l1987:
							{
								position1988, tokenIndex1988 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1988
								}
								position++
								goto l1987
							l1988:
								position, tokenIndex = position1988, tokenIndex1988
							}
							add(rulePegText, position1982)
						}
						{
							add(ruleAction31, position)
						}
						add(rulesignedDecimalByte, position1981)
					}
				}
			l1952:
				add(ruledisp, position1951)
			}
			return true
		l1950:
			position, tokenIndex = position1950, tokenIndex1950
			return false
		},
		/* 51 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action31)> */
		nil,
		/* 52 signedHexByteH <- <(<(('-' / '+')? ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> ('h' / 'H') Action32)> */
		nil,
		/* 53 signedHexByte0x <- <(<(('-' / '+')? ('0' ('x' / 'X')) ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> Action33)> */
		nil,
		/* 54 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action34)> */
		nil,
		/* 55 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action35)> */
		nil,
		/* 56 decimalByte <- <(<[0-9]+> Action36)> */
		nil,
		/* 57 LabelNN <- <(<LabelText> Action37)> */
		nil,
		/* 58 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action38)> */
		nil,
		/* 59 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action39)> */
		nil,
		/* 60 nn_contents <- <('(' nn ')' Action40)> */
		func() bool {
			position1999, tokenIndex1999 := position, tokenIndex
			{
				position2000 := position
				if buffer[position] != rune('(') {
					goto l1999
				}
				position++
				if !_rules[rulenn]() {
					goto l1999
				}
				if buffer[position] != rune(')') {
					goto l1999
				}
				position++
				{
					add(ruleAction40, position)
				}
				add(rulenn_contents, position2000)
			}
			return true
		l1999:
			position, tokenIndex = position1999, tokenIndex1999
			return false
		},
		/* 61 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 62 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action41)> */
		nil,
		/* 63 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action42)> */
		nil,
		/* 64 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action43)> */
		nil,
		/* 65 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action44)> */
		nil,
		/* 66 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action45)> */
		nil,
		/* 67 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action46)> */
		nil,
		/* 68 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action47)> */
		nil,
		/* 69 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action48)> */
		nil,
		/* 70 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 71 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 72 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 (sep Copy8)? Action49)> */
		nil,
		/* 73 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 (sep Copy8)? Action50)> */
		nil,
		/* 74 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action51)> */
		nil,
		/* 75 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 (sep Copy8)? Action52)> */
		nil,
		/* 76 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 (sep Copy8)? Action53)> */
		nil,
		/* 77 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 (sep Copy8)? Action54)> */
		nil,
		/* 78 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 (sep Copy8)? Action55)> */
		nil,
		/* 79 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action56)> */
		nil,
		/* 80 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action57)> */
		nil,
		/* 81 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 (sep Copy8)? Action58)> */
		nil,
		/* 82 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 (sep Copy8)? Action59)> */
		nil,
		/* 83 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 84 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action60)> */
		nil,
		/* 85 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action61)> */
		nil,
		/* 86 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action62)> */
		nil,
		/* 87 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action63)> */
		nil,
		/* 88 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action64)> */
		nil,
		/* 89 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action65)> */
		nil,
		/* 90 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action66)> */
		nil,
		/* 91 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action67)> */
		nil,
		/* 92 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action68)> */
		nil,
		/* 93 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action69)> */
		nil,
		/* 94 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action70)> */
		nil,
		/* 95 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action71)> */
		nil,
		/* 96 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action72)> */
		nil,
		/* 97 EDSimple <- <(Retn / Reti / Rrd / Im0 / Im1 / Im2 / ((&('I' | 'O' | 'i' | 'o') BlitIO) | (&('R' | 'r') Rld) | (&('N' | 'n') Neg) | (&('C' | 'L' | 'c' | 'l') Blit)))> */
		nil,
		/* 98 Neg <- <(<(('n' / 'N') ('e' / 'E') ('g' / 'G'))> Action73)> */
		nil,
		/* 99 Retn <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N'))> Action74)> */
		nil,
		/* 100 Reti <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('i' / 'I'))> Action75)> */
		nil,
		/* 101 Rrd <- <(<(('r' / 'R') ('r' / 'R') ('d' / 'D'))> Action76)> */
		nil,
		/* 102 Rld <- <(<(('r' / 'R') ('l' / 'L') ('d' / 'D'))> Action77)> */
		nil,
		/* 103 Im0 <- <(<(('i' / 'I') ('m' / 'M') ' ' '0')> Action78)> */
		nil,
		/* 104 Im1 <- <(<(('i' / 'I') ('m' / 'M') ' ' '1')> Action79)> */
		nil,
		/* 105 Im2 <- <(<(('i' / 'I') ('m' / 'M') ' ' '2')> Action80)> */
		nil,
		/* 106 Blit <- <(Ldir / Ldi / Cpir / Cpi / Lddr / Ldd / Cpdr / Cpd)> */
		nil,
		/* 107 BlitIO <- <(Inir / Ini / Otir / Outi / Indr / Ind / Otdr / Outd)> */
		nil,
		/* 108 Ldi <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I'))> Action81)> */
		nil,
		/* 109 Cpi <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I'))> Action82)> */
		nil,
		/* 110 Ini <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I'))> Action83)> */
		nil,
		/* 111 Outi <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('i' / 'I'))> Action84)> */
		nil,
		/* 112 Ldd <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D'))> Action85)> */
		nil,
		/* 113 Cpd <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D'))> Action86)> */
		nil,
		/* 114 Ind <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D'))> Action87)> */
		nil,
		/* 115 Outd <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('d' / 'D'))> Action88)> */
		nil,
		/* 116 Ldir <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I') ('r' / 'R'))> Action89)> */
		nil,
		/* 117 Cpir <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I') ('r' / 'R'))> Action90)> */
		nil,
		/* 118 Inir <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I') ('r' / 'R'))> Action91)> */
		nil,
		/* 119 Otir <- <(<(('o' / 'O') ('t' / 'T') ('i' / 'I') ('r' / 'R'))> Action92)> */
		nil,
		/* 120 Lddr <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D') ('r' / 'R'))> Action93)> */
		nil,
		/* 121 Cpdr <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D') ('r' / 'R'))> Action94)> */
		nil,
		/* 122 Indr <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D') ('r' / 'R'))> Action95)> */
		nil,
		/* 123 Otdr <- <(<(('o' / 'O') ('t' / 'T') ('d' / 'D') ('r' / 'R'))> Action96)> */
		nil,
		/* 124 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 125 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action97)> */
		nil,
		/* 126 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action98)> */
		nil,
		/* 127 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action99)> */
		nil,
		/* 128 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action100)> */
		nil,
		/* 129 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action101)> */
		nil,
		/* 130 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action102)> */
		nil,
		/* 131 IO <- <(IN / OUT)> */
		nil,
		/* 132 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action103)> */
		nil,
		/* 133 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action104)> */
		nil,
		/* 134 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position2075, tokenIndex2075 := position, tokenIndex
			{
				position2076 := position
				{
					position2077, tokenIndex2077 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2078
					}
					position++
					{
						position2079, tokenIndex2079 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l2080
						}
						position++
						goto l2079
					l2080:
						position, tokenIndex = position2079, tokenIndex2079
						if buffer[position] != rune('C') {
							goto l2078
						}
						position++
					}
				l2079:
					if buffer[position] != rune(')') {
						goto l2078
					}
					position++
					goto l2077
				l2078:
					position, tokenIndex = position2077, tokenIndex2077
					if buffer[position] != rune('(') {
						goto l2075
					}
					position++
					if !_rules[rulen]() {
						goto l2075
					}
					if buffer[position] != rune(')') {
						goto l2075
					}
					position++
				}
			l2077:
				add(rulePort, position2076)
			}
			return true
		l2075:
			position, tokenIndex = position2075, tokenIndex2075
			return false
		},
		/* 135 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2081, tokenIndex2081 := position, tokenIndex
			{
				position2082 := position
				{
					position2083, tokenIndex2083 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2083
					}
					goto l2084
				l2083:
					position, tokenIndex = position2083, tokenIndex2083
				}
			l2084:
				if buffer[position] != rune(',') {
					goto l2081
				}
				position++
				{
					position2085, tokenIndex2085 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2085
					}
					goto l2086
				l2085:
					position, tokenIndex = position2085, tokenIndex2085
				}
			l2086:
				add(rulesep, position2082)
			}
			return true
		l2081:
			position, tokenIndex = position2081, tokenIndex2081
			return false
		},
		/* 136 ws <- <(' ' / '\t')+> */
		func() bool {
			position2087, tokenIndex2087 := position, tokenIndex
			{
				position2088 := position
				{
					position2091, tokenIndex2091 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2092
					}
					position++
					goto l2091
				l2092:
					position, tokenIndex = position2091, tokenIndex2091
					if buffer[position] != rune('\t') {
						goto l2087
					}
					position++
				}
			l2091:
			l2089:
				{
					position2090, tokenIndex2090 := position, tokenIndex
					{
						position2093, tokenIndex2093 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l2094
						}
						position++
						goto l2093
					l2094:
						position, tokenIndex = position2093, tokenIndex2093
						if buffer[position] != rune('\t') {
							goto l2090
						}
						position++
					}
				l2093:
					goto l2089
				l2090:
					position, tokenIndex = position2090, tokenIndex2090
				}
				add(rulews, position2088)
			}
			return true
		l2087:
			position, tokenIndex = position2087, tokenIndex2087
			return false
		},
		/* 137 A <- <('a' / 'A')> */
		nil,
		/* 138 F <- <('f' / 'F')> */
		nil,
		/* 139 B <- <('b' / 'B')> */
		nil,
		/* 140 C <- <('c' / 'C')> */
		nil,
		/* 141 D <- <('d' / 'D')> */
		nil,
		/* 142 E <- <('e' / 'E')> */
		nil,
		/* 143 H <- <('h' / 'H')> */
		nil,
		/* 144 L <- <('l' / 'L')> */
		nil,
		/* 145 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 146 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 147 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 148 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 149 I <- <('i' / 'I')> */
		nil,
		/* 150 R <- <('r' / 'R')> */
		nil,
		/* 151 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 152 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 153 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 154 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 155 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 156 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 157 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 158 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 159 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position2117, tokenIndex2117 := position, tokenIndex
			{
				position2118 := position
				{
					position2119, tokenIndex2119 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2120
					}
					position++
					goto l2119
				l2120:
					position, tokenIndex = position2119, tokenIndex2119
					{
						position2121, tokenIndex2121 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2122
						}
						position++
						goto l2121
					l2122:
						position, tokenIndex = position2121, tokenIndex2121
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2117
						}
						position++
					}
				l2121:
				}
			l2119:
				add(rulehexdigit, position2118)
			}
			return true
		l2117:
			position, tokenIndex = position2117, tokenIndex2117
			return false
		},
		/* 160 octaldigit <- <(<[0-7]> Action105)> */
		func() bool {
			position2123, tokenIndex2123 := position, tokenIndex
			{
				position2124 := position
				{
					position2125 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2123
					}
					position++
					add(rulePegText, position2125)
				}
				{
					add(ruleAction105, position)
				}
				add(ruleoctaldigit, position2124)
			}
			return true
		l2123:
			position, tokenIndex = position2123, tokenIndex2123
			return false
		},
		/* 161 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2127, tokenIndex2127 := position, tokenIndex
			{
				position2128 := position
				{
					position2129, tokenIndex2129 := position, tokenIndex
					{
						position2131 := position
						{
							position2132, tokenIndex2132 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2133
							}
							position++
							goto l2132
						l2133:
							position, tokenIndex = position2132, tokenIndex2132
							if buffer[position] != rune('N') {
								goto l2130
							}
							position++
						}
					l2132:
						{
							position2134, tokenIndex2134 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2135
							}
							position++
							goto l2134
						l2135:
							position, tokenIndex = position2134, tokenIndex2134
							if buffer[position] != rune('Z') {
								goto l2130
							}
							position++
						}
					l2134:
						{
							add(ruleAction106, position)
						}
						add(ruleFT_NZ, position2131)
					}
					goto l2129
				l2130:
					position, tokenIndex = position2129, tokenIndex2129
					{
						position2138 := position
						{
							position2139, tokenIndex2139 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2140
							}
							position++
							goto l2139
						l2140:
							position, tokenIndex = position2139, tokenIndex2139
							if buffer[position] != rune('P') {
								goto l2137
							}
							position++
						}
					l2139:
						{
							position2141, tokenIndex2141 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2142
							}
							position++
							goto l2141
						l2142:
							position, tokenIndex = position2141, tokenIndex2141
							if buffer[position] != rune('O') {
								goto l2137
							}
							position++
						}
					l2141:
						{
							add(ruleAction110, position)
						}
						add(ruleFT_PO, position2138)
					}
					goto l2129
				l2137:
					position, tokenIndex = position2129, tokenIndex2129
					{
						position2145 := position
						{
							position2146, tokenIndex2146 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2147
							}
							position++
							goto l2146
						l2147:
							position, tokenIndex = position2146, tokenIndex2146
							if buffer[position] != rune('P') {
								goto l2144
							}
							position++
						}
					l2146:
						{
							position2148, tokenIndex2148 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2149
							}
							position++
							goto l2148
						l2149:
							position, tokenIndex = position2148, tokenIndex2148
							if buffer[position] != rune('E') {
								goto l2144
							}
							position++
						}
					l2148:
						{
							add(ruleAction111, position)
						}
						add(ruleFT_PE, position2145)
					}
					goto l2129
				l2144:
					position, tokenIndex = position2129, tokenIndex2129
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2152 := position
								{
									position2153, tokenIndex2153 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2154
									}
									position++
									goto l2153
								l2154:
									position, tokenIndex = position2153, tokenIndex2153
									if buffer[position] != rune('M') {
										goto l2127
									}
									position++
								}
							l2153:
								{
									add(ruleAction113, position)
								}
								add(ruleFT_M, position2152)
							}
							break
						case 'P', 'p':
							{
								position2156 := position
								{
									position2157, tokenIndex2157 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2158
									}
									position++
									goto l2157
								l2158:
									position, tokenIndex = position2157, tokenIndex2157
									if buffer[position] != rune('P') {
										goto l2127
									}
									position++
								}
							l2157:
								{
									add(ruleAction112, position)
								}
								add(ruleFT_P, position2156)
							}
							break
						case 'C', 'c':
							{
								position2160 := position
								{
									position2161, tokenIndex2161 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2162
									}
									position++
									goto l2161
								l2162:
									position, tokenIndex = position2161, tokenIndex2161
									if buffer[position] != rune('C') {
										goto l2127
									}
									position++
								}
							l2161:
								{
									add(ruleAction109, position)
								}
								add(ruleFT_C, position2160)
							}
							break
						case 'N', 'n':
							{
								position2164 := position
								{
									position2165, tokenIndex2165 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2166
									}
									position++
									goto l2165
								l2166:
									position, tokenIndex = position2165, tokenIndex2165
									if buffer[position] != rune('N') {
										goto l2127
									}
									position++
								}
							l2165:
								{
									position2167, tokenIndex2167 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2168
									}
									position++
									goto l2167
								l2168:
									position, tokenIndex = position2167, tokenIndex2167
									if buffer[position] != rune('C') {
										goto l2127
									}
									position++
								}
							l2167:
								{
									add(ruleAction108, position)
								}
								add(ruleFT_NC, position2164)
							}
							break
						default:
							{
								position2170 := position
								{
									position2171, tokenIndex2171 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2172
									}
									position++
									goto l2171
								l2172:
									position, tokenIndex = position2171, tokenIndex2171
									if buffer[position] != rune('Z') {
										goto l2127
									}
									position++
								}
							l2171:
								{
									add(ruleAction107, position)
								}
								add(ruleFT_Z, position2170)
							}
							break
						}
					}

				}
			l2129:
				add(rulecc, position2128)
			}
			return true
		l2127:
			position, tokenIndex = position2127, tokenIndex2127
			return false
		},
		/* 162 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action106)> */
		nil,
		/* 163 FT_Z <- <(('z' / 'Z') Action107)> */
		nil,
		/* 164 FT_NC <- <(('n' / 'N') ('c' / 'C') Action108)> */
		nil,
		/* 165 FT_C <- <(('c' / 'C') Action109)> */
		nil,
		/* 166 FT_PO <- <(('p' / 'P') ('o' / 'O') Action110)> */
		nil,
		/* 167 FT_PE <- <(('p' / 'P') ('e' / 'E') Action111)> */
		nil,
		/* 168 FT_P <- <(('p' / 'P') Action112)> */
		nil,
		/* 169 FT_M <- <(('m' / 'M') Action113)> */
		nil,
		/* 171 Action0 <- <{ p.Emit() }> */
		nil,
		/* 172 Action1 <- <{ p.Org() }> */
		nil,
		/* 173 Action2 <- <{ p.LabelDefn(buffer[begin:end])}> */
		nil,
		nil,
		/* 175 Action3 <- <{ p.LD8() }> */
		nil,
		/* 176 Action4 <- <{ p.LD16() }> */
		nil,
		/* 177 Action5 <- <{ p.Push() }> */
		nil,
		/* 178 Action6 <- <{ p.Pop() }> */
		nil,
		/* 179 Action7 <- <{ p.Ex() }> */
		nil,
		/* 180 Action8 <- <{ p.Inc8() }> */
		nil,
		/* 181 Action9 <- <{ p.Inc8() }> */
		nil,
		/* 182 Action10 <- <{ p.Inc16() }> */
		nil,
		/* 183 Action11 <- <{ p.Dec8() }> */
		nil,
		/* 184 Action12 <- <{ p.Dec8() }> */
		nil,
		/* 185 Action13 <- <{ p.Dec16() }> */
		nil,
		/* 186 Action14 <- <{ p.Add16() }> */
		nil,
		/* 187 Action15 <- <{ p.Adc16() }> */
		nil,
		/* 188 Action16 <- <{ p.Sbc16() }> */
		nil,
		/* 189 Action17 <- <{ p.Dst8() }> */
		nil,
		/* 190 Action18 <- <{ p.Src8() }> */
		nil,
		/* 191 Action19 <- <{ p.Loc8() }> */
		nil,
		/* 192 Action20 <- <{ p.Copy8() }> */
		nil,
		/* 193 Action21 <- <{ p.Loc8() }> */
		nil,
		/* 194 Action22 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 195 Action23 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 196 Action24 <- <{ p.Dst16() }> */
		nil,
		/* 197 Action25 <- <{ p.Src16() }> */
		nil,
		/* 198 Action26 <- <{ p.Loc16() }> */
		nil,
		/* 199 Action27 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 200 Action28 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 201 Action29 <- <{ p.R16Contents() }> */
		nil,
		/* 202 Action30 <- <{ p.IR16Contents() }> */
		nil,
		/* 203 Action31 <- <{ p.DispDecimal(buffer[begin:end]) }> */
		nil,
		/* 204 Action32 <- <{ p.DispHex(buffer[begin:end]) }> */
		nil,
		/* 205 Action33 <- <{ p.Disp0xHex(buffer[begin:end]) }> */
		nil,
		/* 206 Action34 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 207 Action35 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 208 Action36 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 209 Action37 <- <{ p.NNLabel(buffer[begin:end]) }> */
		nil,
		/* 210 Action38 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 211 Action39 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 212 Action40 <- <{ p.NNContents() }> */
		nil,
		/* 213 Action41 <- <{ p.Accum("ADD") }> */
		nil,
		/* 214 Action42 <- <{ p.Accum("ADC") }> */
		nil,
		/* 215 Action43 <- <{ p.Accum("SUB") }> */
		nil,
		/* 216 Action44 <- <{ p.Accum("SBC") }> */
		nil,
		/* 217 Action45 <- <{ p.Accum("AND") }> */
		nil,
		/* 218 Action46 <- <{ p.Accum("XOR") }> */
		nil,
		/* 219 Action47 <- <{ p.Accum("OR") }> */
		nil,
		/* 220 Action48 <- <{ p.Accum("CP") }> */
		nil,
		/* 221 Action49 <- <{ p.Rot("RLC") }> */
		nil,
		/* 222 Action50 <- <{ p.Rot("RRC") }> */
		nil,
		/* 223 Action51 <- <{ p.Rot("RL") }> */
		nil,
		/* 224 Action52 <- <{ p.Rot("RR") }> */
		nil,
		/* 225 Action53 <- <{ p.Rot("SLA") }> */
		nil,
		/* 226 Action54 <- <{ p.Rot("SRA") }> */
		nil,
		/* 227 Action55 <- <{ p.Rot("SLL") }> */
		nil,
		/* 228 Action56 <- <{ p.Rot("SRL") }> */
		nil,
		/* 229 Action57 <- <{ p.Bit() }> */
		nil,
		/* 230 Action58 <- <{ p.Res() }> */
		nil,
		/* 231 Action59 <- <{ p.Set() }> */
		nil,
		/* 232 Action60 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 233 Action61 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 234 Action62 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 235 Action63 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 236 Action64 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 237 Action65 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 238 Action66 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 239 Action67 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 240 Action68 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 241 Action69 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 242 Action70 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 243 Action71 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 244 Action72 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 245 Action73 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 246 Action74 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 247 Action75 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 248 Action76 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 249 Action77 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 250 Action78 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 251 Action79 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 252 Action80 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 253 Action81 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 254 Action82 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 255 Action83 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 256 Action84 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 257 Action85 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 258 Action86 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 259 Action87 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 260 Action88 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 261 Action89 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 262 Action90 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 263 Action91 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 264 Action92 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 265 Action93 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 266 Action94 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 267 Action95 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 268 Action96 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 269 Action97 <- <{ p.Rst() }> */
		nil,
		/* 270 Action98 <- <{ p.Call() }> */
		nil,
		/* 271 Action99 <- <{ p.Ret() }> */
		nil,
		/* 272 Action100 <- <{ p.Jp() }> */
		nil,
		/* 273 Action101 <- <{ p.Jr() }> */
		nil,
		/* 274 Action102 <- <{ p.Djnz() }> */
		nil,
		/* 275 Action103 <- <{ p.In() }> */
		nil,
		/* 276 Action104 <- <{ p.Out() }> */
		nil,
		/* 277 Action105 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 278 Action106 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 279 Action107 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 280 Action108 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 281 Action109 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 282 Action110 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 283 Action111 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 284 Action112 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 285 Action113 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
