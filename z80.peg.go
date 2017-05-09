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
					{
						position5, tokenIndex5 := position, tokenIndex
						{
							position7 := position
							if !_rules[ruleLabelText]() {
								goto l5
							}
							if buffer[position] != rune(':') {
								goto l5
							}
							position++
							if !_rules[rulews]() {
								goto l5
							}
							{
								add(ruleAction2, position)
							}
							add(ruleLabelDefn, position7)
						}
						goto l6
					l5:
						position, tokenIndex = position5, tokenIndex5
					}
				l6:
				l9:
					{
						position10, tokenIndex10 := position, tokenIndex
						if !_rules[rulews]() {
							goto l10
						}
						goto l9
					l10:
						position, tokenIndex = position10, tokenIndex10
					}
					{
						position11, tokenIndex11 := position, tokenIndex
						{
							position13 := position
							{
								position14, tokenIndex14 := position, tokenIndex
								{
									position16 := position
									{
										switch buffer[position] {
										case 'a':
											{
												position18 := position
												if buffer[position] != rune('a') {
													goto l15
												}
												position++
												if buffer[position] != rune('s') {
													goto l15
												}
												position++
												if buffer[position] != rune('e') {
													goto l15
												}
												position++
												if buffer[position] != rune('g') {
													goto l15
												}
												position++
												add(ruleAseg, position18)
											}
											break
										case '.':
											{
												position19 := position
												if buffer[position] != rune('.') {
													goto l15
												}
												position++
												if buffer[position] != rune('t') {
													goto l15
												}
												position++
												if buffer[position] != rune('i') {
													goto l15
												}
												position++
												if buffer[position] != rune('t') {
													goto l15
												}
												position++
												if buffer[position] != rune('l') {
													goto l15
												}
												position++
												if buffer[position] != rune('e') {
													goto l15
												}
												position++
												if !_rules[rulews]() {
													goto l15
												}
												if buffer[position] != rune('\'') {
													goto l15
												}
												position++
											l20:
												{
													position21, tokenIndex21 := position, tokenIndex
													{
														position22, tokenIndex22 := position, tokenIndex
														if buffer[position] != rune('\'') {
															goto l22
														}
														position++
														goto l21
													l22:
														position, tokenIndex = position22, tokenIndex22
													}
													if !matchDot() {
														goto l21
													}
													goto l20
												l21:
													position, tokenIndex = position21, tokenIndex21
												}
												if buffer[position] != rune('\'') {
													goto l15
												}
												position++
												add(ruleTitle, position19)
											}
											break
										default:
											{
												position23 := position
												{
													position24, tokenIndex24 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l25
													}
													position++
													goto l24
												l25:
													position, tokenIndex = position24, tokenIndex24
													if buffer[position] != rune('O') {
														goto l15
													}
													position++
												}
											l24:
												{
													position26, tokenIndex26 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l27
													}
													position++
													goto l26
												l27:
													position, tokenIndex = position26, tokenIndex26
													if buffer[position] != rune('R') {
														goto l15
													}
													position++
												}
											l26:
												{
													position28, tokenIndex28 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l29
													}
													position++
													goto l28
												l29:
													position, tokenIndex = position28, tokenIndex28
													if buffer[position] != rune('G') {
														goto l15
													}
													position++
												}
											l28:
												if !_rules[rulews]() {
													goto l15
												}
												if !_rules[rulenn]() {
													goto l15
												}
												{
													add(ruleAction1, position)
												}
												add(ruleOrg, position23)
											}
											break
										}
									}

									add(ruleDirective, position16)
								}
								goto l14
							l15:
								position, tokenIndex = position14, tokenIndex14
								{
									position31 := position
									{
										position32, tokenIndex32 := position, tokenIndex
										{
											position34 := position
											{
												position35, tokenIndex35 := position, tokenIndex
												{
													position37 := position
													{
														position38, tokenIndex38 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l39
														}
														position++
														goto l38
													l39:
														position, tokenIndex = position38, tokenIndex38
														if buffer[position] != rune('P') {
															goto l36
														}
														position++
													}
												l38:
													{
														position40, tokenIndex40 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l41
														}
														position++
														goto l40
													l41:
														position, tokenIndex = position40, tokenIndex40
														if buffer[position] != rune('U') {
															goto l36
														}
														position++
													}
												l40:
													{
														position42, tokenIndex42 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l43
														}
														position++
														goto l42
													l43:
														position, tokenIndex = position42, tokenIndex42
														if buffer[position] != rune('S') {
															goto l36
														}
														position++
													}
												l42:
													{
														position44, tokenIndex44 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l45
														}
														position++
														goto l44
													l45:
														position, tokenIndex = position44, tokenIndex44
														if buffer[position] != rune('H') {
															goto l36
														}
														position++
													}
												l44:
													if !_rules[rulews]() {
														goto l36
													}
													if !_rules[ruleSrc16]() {
														goto l36
													}
													{
														add(ruleAction5, position)
													}
													add(rulePush, position37)
												}
												goto l35
											l36:
												position, tokenIndex = position35, tokenIndex35
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position48 := position
															{
																position49, tokenIndex49 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l50
																}
																position++
																goto l49
															l50:
																position, tokenIndex = position49, tokenIndex49
																if buffer[position] != rune('E') {
																	goto l33
																}
																position++
															}
														l49:
															{
																position51, tokenIndex51 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l52
																}
																position++
																goto l51
															l52:
																position, tokenIndex = position51, tokenIndex51
																if buffer[position] != rune('X') {
																	goto l33
																}
																position++
															}
														l51:
															if !_rules[rulews]() {
																goto l33
															}
															if !_rules[ruleDst16]() {
																goto l33
															}
															if !_rules[rulesep]() {
																goto l33
															}
															if !_rules[ruleSrc16]() {
																goto l33
															}
															{
																add(ruleAction7, position)
															}
															add(ruleEx, position48)
														}
														break
													case 'P', 'p':
														{
															position54 := position
															{
																position55, tokenIndex55 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l56
																}
																position++
																goto l55
															l56:
																position, tokenIndex = position55, tokenIndex55
																if buffer[position] != rune('P') {
																	goto l33
																}
																position++
															}
														l55:
															{
																position57, tokenIndex57 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l58
																}
																position++
																goto l57
															l58:
																position, tokenIndex = position57, tokenIndex57
																if buffer[position] != rune('O') {
																	goto l33
																}
																position++
															}
														l57:
															{
																position59, tokenIndex59 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l60
																}
																position++
																goto l59
															l60:
																position, tokenIndex = position59, tokenIndex59
																if buffer[position] != rune('P') {
																	goto l33
																}
																position++
															}
														l59:
															if !_rules[rulews]() {
																goto l33
															}
															if !_rules[ruleDst16]() {
																goto l33
															}
															{
																add(ruleAction6, position)
															}
															add(rulePop, position54)
														}
														break
													default:
														{
															position62 := position
															{
																position63, tokenIndex63 := position, tokenIndex
																{
																	position65 := position
																	{
																		position66, tokenIndex66 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l67
																		}
																		position++
																		goto l66
																	l67:
																		position, tokenIndex = position66, tokenIndex66
																		if buffer[position] != rune('L') {
																			goto l64
																		}
																		position++
																	}
																l66:
																	{
																		position68, tokenIndex68 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l69
																		}
																		position++
																		goto l68
																	l69:
																		position, tokenIndex = position68, tokenIndex68
																		if buffer[position] != rune('D') {
																			goto l64
																		}
																		position++
																	}
																l68:
																	if !_rules[rulews]() {
																		goto l64
																	}
																	if !_rules[ruleDst16]() {
																		goto l64
																	}
																	if !_rules[rulesep]() {
																		goto l64
																	}
																	if !_rules[ruleSrc16]() {
																		goto l64
																	}
																	{
																		add(ruleAction4, position)
																	}
																	add(ruleLoad16, position65)
																}
																goto l63
															l64:
																position, tokenIndex = position63, tokenIndex63
																{
																	position71 := position
																	{
																		position72, tokenIndex72 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l73
																		}
																		position++
																		goto l72
																	l73:
																		position, tokenIndex = position72, tokenIndex72
																		if buffer[position] != rune('L') {
																			goto l33
																		}
																		position++
																	}
																l72:
																	{
																		position74, tokenIndex74 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l75
																		}
																		position++
																		goto l74
																	l75:
																		position, tokenIndex = position74, tokenIndex74
																		if buffer[position] != rune('D') {
																			goto l33
																		}
																		position++
																	}
																l74:
																	if !_rules[rulews]() {
																		goto l33
																	}
																	{
																		position76 := position
																		{
																			position77, tokenIndex77 := position, tokenIndex
																			if !_rules[ruleReg8]() {
																				goto l78
																			}
																			goto l77
																		l78:
																			position, tokenIndex = position77, tokenIndex77
																			if !_rules[ruleReg16Contents]() {
																				goto l79
																			}
																			goto l77
																		l79:
																			position, tokenIndex = position77, tokenIndex77
																			if !_rules[rulenn_contents]() {
																				goto l33
																			}
																		}
																	l77:
																		{
																			add(ruleAction17, position)
																		}
																		add(ruleDst8, position76)
																	}
																	if !_rules[rulesep]() {
																		goto l33
																	}
																	if !_rules[ruleSrc8]() {
																		goto l33
																	}
																	{
																		add(ruleAction3, position)
																	}
																	add(ruleLoad8, position71)
																}
															}
														l63:
															add(ruleLoad, position62)
														}
														break
													}
												}

											}
										l35:
											add(ruleAssignment, position34)
										}
										goto l32
									l33:
										position, tokenIndex = position32, tokenIndex32
										{
											position83 := position
											{
												position84, tokenIndex84 := position, tokenIndex
												{
													position86 := position
													{
														position87, tokenIndex87 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l88
														}
														position++
														goto l87
													l88:
														position, tokenIndex = position87, tokenIndex87
														if buffer[position] != rune('I') {
															goto l85
														}
														position++
													}
												l87:
													{
														position89, tokenIndex89 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l90
														}
														position++
														goto l89
													l90:
														position, tokenIndex = position89, tokenIndex89
														if buffer[position] != rune('N') {
															goto l85
														}
														position++
													}
												l89:
													{
														position91, tokenIndex91 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l92
														}
														position++
														goto l91
													l92:
														position, tokenIndex = position91, tokenIndex91
														if buffer[position] != rune('C') {
															goto l85
														}
														position++
													}
												l91:
													if !_rules[rulews]() {
														goto l85
													}
													if !_rules[ruleILoc8]() {
														goto l85
													}
													{
														add(ruleAction8, position)
													}
													add(ruleInc16Indexed8, position86)
												}
												goto l84
											l85:
												position, tokenIndex = position84, tokenIndex84
												{
													position95 := position
													{
														position96, tokenIndex96 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l97
														}
														position++
														goto l96
													l97:
														position, tokenIndex = position96, tokenIndex96
														if buffer[position] != rune('I') {
															goto l94
														}
														position++
													}
												l96:
													{
														position98, tokenIndex98 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l99
														}
														position++
														goto l98
													l99:
														position, tokenIndex = position98, tokenIndex98
														if buffer[position] != rune('N') {
															goto l94
														}
														position++
													}
												l98:
													{
														position100, tokenIndex100 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l101
														}
														position++
														goto l100
													l101:
														position, tokenIndex = position100, tokenIndex100
														if buffer[position] != rune('C') {
															goto l94
														}
														position++
													}
												l100:
													if !_rules[rulews]() {
														goto l94
													}
													if !_rules[ruleLoc16]() {
														goto l94
													}
													{
														add(ruleAction10, position)
													}
													add(ruleInc16, position95)
												}
												goto l84
											l94:
												position, tokenIndex = position84, tokenIndex84
												{
													position103 := position
													{
														position104, tokenIndex104 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l105
														}
														position++
														goto l104
													l105:
														position, tokenIndex = position104, tokenIndex104
														if buffer[position] != rune('I') {
															goto l82
														}
														position++
													}
												l104:
													{
														position106, tokenIndex106 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l107
														}
														position++
														goto l106
													l107:
														position, tokenIndex = position106, tokenIndex106
														if buffer[position] != rune('N') {
															goto l82
														}
														position++
													}
												l106:
													{
														position108, tokenIndex108 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l109
														}
														position++
														goto l108
													l109:
														position, tokenIndex = position108, tokenIndex108
														if buffer[position] != rune('C') {
															goto l82
														}
														position++
													}
												l108:
													if !_rules[rulews]() {
														goto l82
													}
													if !_rules[ruleLoc8]() {
														goto l82
													}
													{
														add(ruleAction9, position)
													}
													add(ruleInc8, position103)
												}
											}
										l84:
											add(ruleInc, position83)
										}
										goto l32
									l82:
										position, tokenIndex = position32, tokenIndex32
										{
											position112 := position
											{
												position113, tokenIndex113 := position, tokenIndex
												{
													position115 := position
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
															goto l114
														}
														position++
													}
												l116:
													{
														position118, tokenIndex118 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l119
														}
														position++
														goto l118
													l119:
														position, tokenIndex = position118, tokenIndex118
														if buffer[position] != rune('E') {
															goto l114
														}
														position++
													}
												l118:
													{
														position120, tokenIndex120 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l121
														}
														position++
														goto l120
													l121:
														position, tokenIndex = position120, tokenIndex120
														if buffer[position] != rune('C') {
															goto l114
														}
														position++
													}
												l120:
													if !_rules[rulews]() {
														goto l114
													}
													if !_rules[ruleILoc8]() {
														goto l114
													}
													{
														add(ruleAction11, position)
													}
													add(ruleDec16Indexed8, position115)
												}
												goto l113
											l114:
												position, tokenIndex = position113, tokenIndex113
												{
													position124 := position
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
															goto l123
														}
														position++
													}
												l125:
													{
														position127, tokenIndex127 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l128
														}
														position++
														goto l127
													l128:
														position, tokenIndex = position127, tokenIndex127
														if buffer[position] != rune('E') {
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
													if !_rules[ruleLoc16]() {
														goto l123
													}
													{
														add(ruleAction13, position)
													}
													add(ruleDec16, position124)
												}
												goto l113
											l123:
												position, tokenIndex = position113, tokenIndex113
												{
													position132 := position
													{
														position133, tokenIndex133 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l134
														}
														position++
														goto l133
													l134:
														position, tokenIndex = position133, tokenIndex133
														if buffer[position] != rune('D') {
															goto l111
														}
														position++
													}
												l133:
													{
														position135, tokenIndex135 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l136
														}
														position++
														goto l135
													l136:
														position, tokenIndex = position135, tokenIndex135
														if buffer[position] != rune('E') {
															goto l111
														}
														position++
													}
												l135:
													{
														position137, tokenIndex137 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l138
														}
														position++
														goto l137
													l138:
														position, tokenIndex = position137, tokenIndex137
														if buffer[position] != rune('C') {
															goto l111
														}
														position++
													}
												l137:
													if !_rules[rulews]() {
														goto l111
													}
													if !_rules[ruleLoc8]() {
														goto l111
													}
													{
														add(ruleAction12, position)
													}
													add(ruleDec8, position132)
												}
											}
										l113:
											add(ruleDec, position112)
										}
										goto l32
									l111:
										position, tokenIndex = position32, tokenIndex32
										{
											position141 := position
											{
												position142, tokenIndex142 := position, tokenIndex
												{
													position144 := position
													{
														position145, tokenIndex145 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l146
														}
														position++
														goto l145
													l146:
														position, tokenIndex = position145, tokenIndex145
														if buffer[position] != rune('A') {
															goto l143
														}
														position++
													}
												l145:
													{
														position147, tokenIndex147 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l148
														}
														position++
														goto l147
													l148:
														position, tokenIndex = position147, tokenIndex147
														if buffer[position] != rune('D') {
															goto l143
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
															goto l143
														}
														position++
													}
												l149:
													if !_rules[rulews]() {
														goto l143
													}
													if !_rules[ruleDst16]() {
														goto l143
													}
													if !_rules[rulesep]() {
														goto l143
													}
													if !_rules[ruleSrc16]() {
														goto l143
													}
													{
														add(ruleAction14, position)
													}
													add(ruleAdd16, position144)
												}
												goto l142
											l143:
												position, tokenIndex = position142, tokenIndex142
												{
													position153 := position
													{
														position154, tokenIndex154 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l155
														}
														position++
														goto l154
													l155:
														position, tokenIndex = position154, tokenIndex154
														if buffer[position] != rune('A') {
															goto l152
														}
														position++
													}
												l154:
													{
														position156, tokenIndex156 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l157
														}
														position++
														goto l156
													l157:
														position, tokenIndex = position156, tokenIndex156
														if buffer[position] != rune('D') {
															goto l152
														}
														position++
													}
												l156:
													{
														position158, tokenIndex158 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l159
														}
														position++
														goto l158
													l159:
														position, tokenIndex = position158, tokenIndex158
														if buffer[position] != rune('C') {
															goto l152
														}
														position++
													}
												l158:
													if !_rules[rulews]() {
														goto l152
													}
													if !_rules[ruleDst16]() {
														goto l152
													}
													if !_rules[rulesep]() {
														goto l152
													}
													if !_rules[ruleSrc16]() {
														goto l152
													}
													{
														add(ruleAction15, position)
													}
													add(ruleAdc16, position153)
												}
												goto l142
											l152:
												position, tokenIndex = position142, tokenIndex142
												{
													position161 := position
													{
														position162, tokenIndex162 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l163
														}
														position++
														goto l162
													l163:
														position, tokenIndex = position162, tokenIndex162
														if buffer[position] != rune('S') {
															goto l140
														}
														position++
													}
												l162:
													{
														position164, tokenIndex164 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l165
														}
														position++
														goto l164
													l165:
														position, tokenIndex = position164, tokenIndex164
														if buffer[position] != rune('B') {
															goto l140
														}
														position++
													}
												l164:
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
															goto l140
														}
														position++
													}
												l166:
													if !_rules[rulews]() {
														goto l140
													}
													if !_rules[ruleDst16]() {
														goto l140
													}
													if !_rules[rulesep]() {
														goto l140
													}
													if !_rules[ruleSrc16]() {
														goto l140
													}
													{
														add(ruleAction16, position)
													}
													add(ruleSbc16, position161)
												}
											}
										l142:
											add(ruleAlu16, position141)
										}
										goto l32
									l140:
										position, tokenIndex = position32, tokenIndex32
										{
											position170 := position
											{
												position171, tokenIndex171 := position, tokenIndex
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
														if buffer[position] != rune('d') {
															goto l179
														}
														position++
														goto l178
													l179:
														position, tokenIndex = position178, tokenIndex178
														if buffer[position] != rune('D') {
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
														add(ruleAction41, position)
													}
													add(ruleAdd, position173)
												}
												goto l171
											l172:
												position, tokenIndex = position171, tokenIndex171
												{
													position184 := position
													{
														position185, tokenIndex185 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l186
														}
														position++
														goto l185
													l186:
														position, tokenIndex = position185, tokenIndex185
														if buffer[position] != rune('A') {
															goto l183
														}
														position++
													}
												l185:
													{
														position187, tokenIndex187 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l188
														}
														position++
														goto l187
													l188:
														position, tokenIndex = position187, tokenIndex187
														if buffer[position] != rune('D') {
															goto l183
														}
														position++
													}
												l187:
													{
														position189, tokenIndex189 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l190
														}
														position++
														goto l189
													l190:
														position, tokenIndex = position189, tokenIndex189
														if buffer[position] != rune('C') {
															goto l183
														}
														position++
													}
												l189:
													if !_rules[rulews]() {
														goto l183
													}
													{
														position191, tokenIndex191 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l192
														}
														position++
														goto l191
													l192:
														position, tokenIndex = position191, tokenIndex191
														if buffer[position] != rune('A') {
															goto l183
														}
														position++
													}
												l191:
													if !_rules[rulesep]() {
														goto l183
													}
													if !_rules[ruleSrc8]() {
														goto l183
													}
													{
														add(ruleAction42, position)
													}
													add(ruleAdc, position184)
												}
												goto l171
											l183:
												position, tokenIndex = position171, tokenIndex171
												{
													position195 := position
													{
														position196, tokenIndex196 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l197
														}
														position++
														goto l196
													l197:
														position, tokenIndex = position196, tokenIndex196
														if buffer[position] != rune('S') {
															goto l194
														}
														position++
													}
												l196:
													{
														position198, tokenIndex198 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l199
														}
														position++
														goto l198
													l199:
														position, tokenIndex = position198, tokenIndex198
														if buffer[position] != rune('U') {
															goto l194
														}
														position++
													}
												l198:
													{
														position200, tokenIndex200 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l201
														}
														position++
														goto l200
													l201:
														position, tokenIndex = position200, tokenIndex200
														if buffer[position] != rune('B') {
															goto l194
														}
														position++
													}
												l200:
													if !_rules[rulews]() {
														goto l194
													}
													if !_rules[ruleSrc8]() {
														goto l194
													}
													{
														add(ruleAction43, position)
													}
													add(ruleSub, position195)
												}
												goto l171
											l194:
												position, tokenIndex = position171, tokenIndex171
												{
													switch buffer[position] {
													case 'C', 'c':
														{
															position204 := position
															{
																position205, tokenIndex205 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l206
																}
																position++
																goto l205
															l206:
																position, tokenIndex = position205, tokenIndex205
																if buffer[position] != rune('C') {
																	goto l169
																}
																position++
															}
														l205:
															{
																position207, tokenIndex207 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l208
																}
																position++
																goto l207
															l208:
																position, tokenIndex = position207, tokenIndex207
																if buffer[position] != rune('P') {
																	goto l169
																}
																position++
															}
														l207:
															if !_rules[rulews]() {
																goto l169
															}
															if !_rules[ruleSrc8]() {
																goto l169
															}
															{
																add(ruleAction48, position)
															}
															add(ruleCp, position204)
														}
														break
													case 'O', 'o':
														{
															position210 := position
															{
																position211, tokenIndex211 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l212
																}
																position++
																goto l211
															l212:
																position, tokenIndex = position211, tokenIndex211
																if buffer[position] != rune('O') {
																	goto l169
																}
																position++
															}
														l211:
															{
																position213, tokenIndex213 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l214
																}
																position++
																goto l213
															l214:
																position, tokenIndex = position213, tokenIndex213
																if buffer[position] != rune('R') {
																	goto l169
																}
																position++
															}
														l213:
															if !_rules[rulews]() {
																goto l169
															}
															if !_rules[ruleSrc8]() {
																goto l169
															}
															{
																add(ruleAction47, position)
															}
															add(ruleOr, position210)
														}
														break
													case 'X', 'x':
														{
															position216 := position
															{
																position217, tokenIndex217 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l218
																}
																position++
																goto l217
															l218:
																position, tokenIndex = position217, tokenIndex217
																if buffer[position] != rune('X') {
																	goto l169
																}
																position++
															}
														l217:
															{
																position219, tokenIndex219 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l220
																}
																position++
																goto l219
															l220:
																position, tokenIndex = position219, tokenIndex219
																if buffer[position] != rune('O') {
																	goto l169
																}
																position++
															}
														l219:
															{
																position221, tokenIndex221 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l222
																}
																position++
																goto l221
															l222:
																position, tokenIndex = position221, tokenIndex221
																if buffer[position] != rune('R') {
																	goto l169
																}
																position++
															}
														l221:
															if !_rules[rulews]() {
																goto l169
															}
															if !_rules[ruleSrc8]() {
																goto l169
															}
															{
																add(ruleAction46, position)
															}
															add(ruleXor, position216)
														}
														break
													case 'A', 'a':
														{
															position224 := position
															{
																position225, tokenIndex225 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l226
																}
																position++
																goto l225
															l226:
																position, tokenIndex = position225, tokenIndex225
																if buffer[position] != rune('A') {
																	goto l169
																}
																position++
															}
														l225:
															{
																position227, tokenIndex227 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l228
																}
																position++
																goto l227
															l228:
																position, tokenIndex = position227, tokenIndex227
																if buffer[position] != rune('N') {
																	goto l169
																}
																position++
															}
														l227:
															{
																position229, tokenIndex229 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l230
																}
																position++
																goto l229
															l230:
																position, tokenIndex = position229, tokenIndex229
																if buffer[position] != rune('D') {
																	goto l169
																}
																position++
															}
														l229:
															if !_rules[rulews]() {
																goto l169
															}
															if !_rules[ruleSrc8]() {
																goto l169
															}
															{
																add(ruleAction45, position)
															}
															add(ruleAnd, position224)
														}
														break
													default:
														{
															position232 := position
															{
																position233, tokenIndex233 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l234
																}
																position++
																goto l233
															l234:
																position, tokenIndex = position233, tokenIndex233
																if buffer[position] != rune('S') {
																	goto l169
																}
																position++
															}
														l233:
															{
																position235, tokenIndex235 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l236
																}
																position++
																goto l235
															l236:
																position, tokenIndex = position235, tokenIndex235
																if buffer[position] != rune('B') {
																	goto l169
																}
																position++
															}
														l235:
															{
																position237, tokenIndex237 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l238
																}
																position++
																goto l237
															l238:
																position, tokenIndex = position237, tokenIndex237
																if buffer[position] != rune('C') {
																	goto l169
																}
																position++
															}
														l237:
															if !_rules[rulews]() {
																goto l169
															}
															{
																position239, tokenIndex239 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l240
																}
																position++
																goto l239
															l240:
																position, tokenIndex = position239, tokenIndex239
																if buffer[position] != rune('A') {
																	goto l169
																}
																position++
															}
														l239:
															if !_rules[rulesep]() {
																goto l169
															}
															if !_rules[ruleSrc8]() {
																goto l169
															}
															{
																add(ruleAction44, position)
															}
															add(ruleSbc, position232)
														}
														break
													}
												}

											}
										l171:
											add(ruleAlu, position170)
										}
										goto l32
									l169:
										position, tokenIndex = position32, tokenIndex32
										{
											position243 := position
											{
												position244, tokenIndex244 := position, tokenIndex
												{
													position246 := position
													{
														position247, tokenIndex247 := position, tokenIndex
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
																if buffer[position] != rune('l') {
																	goto l253
																}
																position++
																goto l252
															l253:
																position, tokenIndex = position252, tokenIndex252
																if buffer[position] != rune('L') {
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
																add(ruleAction49, position)
															}
															add(ruleRlc, position249)
														}
														goto l247
													l248:
														position, tokenIndex = position247, tokenIndex247
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
																if buffer[position] != rune('r') {
																	goto l264
																}
																position++
																goto l263
															l264:
																position, tokenIndex = position263, tokenIndex263
																if buffer[position] != rune('R') {
																	goto l259
																}
																position++
															}
														l263:
															{
																position265, tokenIndex265 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l266
																}
																position++
																goto l265
															l266:
																position, tokenIndex = position265, tokenIndex265
																if buffer[position] != rune('C') {
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
																position267, tokenIndex267 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l267
																}
																if !_rules[ruleCopy8]() {
																	goto l267
																}
																goto l268
															l267:
																position, tokenIndex = position267, tokenIndex267
															}
														l268:
															{
																add(ruleAction50, position)
															}
															add(ruleRrc, position260)
														}
														goto l247
													l259:
														position, tokenIndex = position247, tokenIndex247
														{
															position271 := position
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
																	goto l270
																}
																position++
															}
														l272:
															{
																position274, tokenIndex274 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l275
																}
																position++
																goto l274
															l275:
																position, tokenIndex = position274, tokenIndex274
																if buffer[position] != rune('L') {
																	goto l270
																}
																position++
															}
														l274:
															if !_rules[rulews]() {
																goto l270
															}
															if !_rules[ruleLoc8]() {
																goto l270
															}
															{
																position276, tokenIndex276 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l276
																}
																if !_rules[ruleCopy8]() {
																	goto l276
																}
																goto l277
															l276:
																position, tokenIndex = position276, tokenIndex276
															}
														l277:
															{
																add(ruleAction51, position)
															}
															add(ruleRl, position271)
														}
														goto l247
													l270:
														position, tokenIndex = position247, tokenIndex247
														{
															position280 := position
															{
																position281, tokenIndex281 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l282
																}
																position++
																goto l281
															l282:
																position, tokenIndex = position281, tokenIndex281
																if buffer[position] != rune('R') {
																	goto l279
																}
																position++
															}
														l281:
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
																	goto l279
																}
																position++
															}
														l283:
															if !_rules[rulews]() {
																goto l279
															}
															if !_rules[ruleLoc8]() {
																goto l279
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
																add(ruleAction52, position)
															}
															add(ruleRr, position280)
														}
														goto l247
													l279:
														position, tokenIndex = position247, tokenIndex247
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
																if buffer[position] != rune('l') {
																	goto l293
																}
																position++
																goto l292
															l293:
																position, tokenIndex = position292, tokenIndex292
																if buffer[position] != rune('L') {
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
																add(ruleAction53, position)
															}
															add(ruleSla, position289)
														}
														goto l247
													l288:
														position, tokenIndex = position247, tokenIndex247
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
																if buffer[position] != rune('r') {
																	goto l304
																}
																position++
																goto l303
															l304:
																position, tokenIndex = position303, tokenIndex303
																if buffer[position] != rune('R') {
																	goto l299
																}
																position++
															}
														l303:
															{
																position305, tokenIndex305 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l306
																}
																position++
																goto l305
															l306:
																position, tokenIndex = position305, tokenIndex305
																if buffer[position] != rune('A') {
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
																add(ruleAction54, position)
															}
															add(ruleSra, position300)
														}
														goto l247
													l299:
														position, tokenIndex = position247, tokenIndex247
														{
															position311 := position
															{
																position312, tokenIndex312 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l313
																}
																position++
																goto l312
															l313:
																position, tokenIndex = position312, tokenIndex312
																if buffer[position] != rune('S') {
																	goto l310
																}
																position++
															}
														l312:
															{
																position314, tokenIndex314 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l315
																}
																position++
																goto l314
															l315:
																position, tokenIndex = position314, tokenIndex314
																if buffer[position] != rune('L') {
																	goto l310
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
																	goto l310
																}
																position++
															}
														l316:
															if !_rules[rulews]() {
																goto l310
															}
															if !_rules[ruleLoc8]() {
																goto l310
															}
															{
																position318, tokenIndex318 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l318
																}
																if !_rules[ruleCopy8]() {
																	goto l318
																}
																goto l319
															l318:
																position, tokenIndex = position318, tokenIndex318
															}
														l319:
															{
																add(ruleAction55, position)
															}
															add(ruleSll, position311)
														}
														goto l247
													l310:
														position, tokenIndex = position247, tokenIndex247
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
																	goto l245
																}
																position++
															}
														l322:
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
																	goto l245
																}
																position++
															}
														l324:
															{
																position326, tokenIndex326 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l327
																}
																position++
																goto l326
															l327:
																position, tokenIndex = position326, tokenIndex326
																if buffer[position] != rune('L') {
																	goto l245
																}
																position++
															}
														l326:
															if !_rules[rulews]() {
																goto l245
															}
															if !_rules[ruleLoc8]() {
																goto l245
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
																add(ruleAction56, position)
															}
															add(ruleSrl, position321)
														}
													}
												l247:
													add(ruleRot, position246)
												}
												goto l244
											l245:
												position, tokenIndex = position244, tokenIndex244
												{
													switch buffer[position] {
													case 'S', 's':
														{
															position332 := position
															{
																position333, tokenIndex333 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l334
																}
																position++
																goto l333
															l334:
																position, tokenIndex = position333, tokenIndex333
																if buffer[position] != rune('S') {
																	goto l242
																}
																position++
															}
														l333:
															{
																position335, tokenIndex335 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l336
																}
																position++
																goto l335
															l336:
																position, tokenIndex = position335, tokenIndex335
																if buffer[position] != rune('E') {
																	goto l242
																}
																position++
															}
														l335:
															{
																position337, tokenIndex337 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l338
																}
																position++
																goto l337
															l338:
																position, tokenIndex = position337, tokenIndex337
																if buffer[position] != rune('T') {
																	goto l242
																}
																position++
															}
														l337:
															if !_rules[rulews]() {
																goto l242
															}
															if !_rules[ruleoctaldigit]() {
																goto l242
															}
															if !_rules[rulesep]() {
																goto l242
															}
															if !_rules[ruleLoc8]() {
																goto l242
															}
															{
																position339, tokenIndex339 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l339
																}
																if !_rules[ruleCopy8]() {
																	goto l339
																}
																goto l340
															l339:
																position, tokenIndex = position339, tokenIndex339
															}
														l340:
															{
																add(ruleAction59, position)
															}
															add(ruleSet, position332)
														}
														break
													case 'R', 'r':
														{
															position342 := position
															{
																position343, tokenIndex343 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l344
																}
																position++
																goto l343
															l344:
																position, tokenIndex = position343, tokenIndex343
																if buffer[position] != rune('R') {
																	goto l242
																}
																position++
															}
														l343:
															{
																position345, tokenIndex345 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l346
																}
																position++
																goto l345
															l346:
																position, tokenIndex = position345, tokenIndex345
																if buffer[position] != rune('E') {
																	goto l242
																}
																position++
															}
														l345:
															{
																position347, tokenIndex347 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l348
																}
																position++
																goto l347
															l348:
																position, tokenIndex = position347, tokenIndex347
																if buffer[position] != rune('S') {
																	goto l242
																}
																position++
															}
														l347:
															if !_rules[rulews]() {
																goto l242
															}
															if !_rules[ruleoctaldigit]() {
																goto l242
															}
															if !_rules[rulesep]() {
																goto l242
															}
															if !_rules[ruleLoc8]() {
																goto l242
															}
															{
																position349, tokenIndex349 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l349
																}
																if !_rules[ruleCopy8]() {
																	goto l349
																}
																goto l350
															l349:
																position, tokenIndex = position349, tokenIndex349
															}
														l350:
															{
																add(ruleAction58, position)
															}
															add(ruleRes, position342)
														}
														break
													default:
														{
															position352 := position
															{
																position353, tokenIndex353 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l354
																}
																position++
																goto l353
															l354:
																position, tokenIndex = position353, tokenIndex353
																if buffer[position] != rune('B') {
																	goto l242
																}
																position++
															}
														l353:
															{
																position355, tokenIndex355 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l356
																}
																position++
																goto l355
															l356:
																position, tokenIndex = position355, tokenIndex355
																if buffer[position] != rune('I') {
																	goto l242
																}
																position++
															}
														l355:
															{
																position357, tokenIndex357 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l358
																}
																position++
																goto l357
															l358:
																position, tokenIndex = position357, tokenIndex357
																if buffer[position] != rune('T') {
																	goto l242
																}
																position++
															}
														l357:
															if !_rules[rulews]() {
																goto l242
															}
															if !_rules[ruleoctaldigit]() {
																goto l242
															}
															if !_rules[rulesep]() {
																goto l242
															}
															if !_rules[ruleLoc8]() {
																goto l242
															}
															{
																add(ruleAction57, position)
															}
															add(ruleBit, position352)
														}
														break
													}
												}

											}
										l244:
											add(ruleBitOp, position243)
										}
										goto l32
									l242:
										position, tokenIndex = position32, tokenIndex32
										{
											position361 := position
											{
												position362, tokenIndex362 := position, tokenIndex
												{
													position364 := position
													{
														position365 := position
														{
															position366, tokenIndex366 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l367
															}
															position++
															goto l366
														l367:
															position, tokenIndex = position366, tokenIndex366
															if buffer[position] != rune('R') {
																goto l363
															}
															position++
														}
													l366:
														{
															position368, tokenIndex368 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l369
															}
															position++
															goto l368
														l369:
															position, tokenIndex = position368, tokenIndex368
															if buffer[position] != rune('E') {
																goto l363
															}
															position++
														}
													l368:
														{
															position370, tokenIndex370 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l371
															}
															position++
															goto l370
														l371:
															position, tokenIndex = position370, tokenIndex370
															if buffer[position] != rune('T') {
																goto l363
															}
															position++
														}
													l370:
														{
															position372, tokenIndex372 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l373
															}
															position++
															goto l372
														l373:
															position, tokenIndex = position372, tokenIndex372
															if buffer[position] != rune('N') {
																goto l363
															}
															position++
														}
													l372:
														add(rulePegText, position365)
													}
													{
														add(ruleAction74, position)
													}
													add(ruleRetn, position364)
												}
												goto l362
											l363:
												position, tokenIndex = position362, tokenIndex362
												{
													position376 := position
													{
														position377 := position
														{
															position378, tokenIndex378 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l379
															}
															position++
															goto l378
														l379:
															position, tokenIndex = position378, tokenIndex378
															if buffer[position] != rune('R') {
																goto l375
															}
															position++
														}
													l378:
														{
															position380, tokenIndex380 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l381
															}
															position++
															goto l380
														l381:
															position, tokenIndex = position380, tokenIndex380
															if buffer[position] != rune('E') {
																goto l375
															}
															position++
														}
													l380:
														{
															position382, tokenIndex382 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l383
															}
															position++
															goto l382
														l383:
															position, tokenIndex = position382, tokenIndex382
															if buffer[position] != rune('T') {
																goto l375
															}
															position++
														}
													l382:
														{
															position384, tokenIndex384 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l385
															}
															position++
															goto l384
														l385:
															position, tokenIndex = position384, tokenIndex384
															if buffer[position] != rune('I') {
																goto l375
															}
															position++
														}
													l384:
														add(rulePegText, position377)
													}
													{
														add(ruleAction75, position)
													}
													add(ruleReti, position376)
												}
												goto l362
											l375:
												position, tokenIndex = position362, tokenIndex362
												{
													position388 := position
													{
														position389 := position
														{
															position390, tokenIndex390 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l391
															}
															position++
															goto l390
														l391:
															position, tokenIndex = position390, tokenIndex390
															if buffer[position] != rune('R') {
																goto l387
															}
															position++
														}
													l390:
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
																goto l387
															}
															position++
														}
													l392:
														{
															position394, tokenIndex394 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l395
															}
															position++
															goto l394
														l395:
															position, tokenIndex = position394, tokenIndex394
															if buffer[position] != rune('D') {
																goto l387
															}
															position++
														}
													l394:
														add(rulePegText, position389)
													}
													{
														add(ruleAction76, position)
													}
													add(ruleRrd, position388)
												}
												goto l362
											l387:
												position, tokenIndex = position362, tokenIndex362
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
															if buffer[position] != rune('m') {
																goto l403
															}
															position++
															goto l402
														l403:
															position, tokenIndex = position402, tokenIndex402
															if buffer[position] != rune('M') {
																goto l397
															}
															position++
														}
													l402:
														if buffer[position] != rune(' ') {
															goto l397
														}
														position++
														if buffer[position] != rune('0') {
															goto l397
														}
														position++
														add(rulePegText, position399)
													}
													{
														add(ruleAction78, position)
													}
													add(ruleIm0, position398)
												}
												goto l362
											l397:
												position, tokenIndex = position362, tokenIndex362
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
															if buffer[position] != rune('m') {
																goto l411
															}
															position++
															goto l410
														l411:
															position, tokenIndex = position410, tokenIndex410
															if buffer[position] != rune('M') {
																goto l405
															}
															position++
														}
													l410:
														if buffer[position] != rune(' ') {
															goto l405
														}
														position++
														if buffer[position] != rune('1') {
															goto l405
														}
														position++
														add(rulePegText, position407)
													}
													{
														add(ruleAction79, position)
													}
													add(ruleIm1, position406)
												}
												goto l362
											l405:
												position, tokenIndex = position362, tokenIndex362
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
															if buffer[position] != rune('m') {
																goto l419
															}
															position++
															goto l418
														l419:
															position, tokenIndex = position418, tokenIndex418
															if buffer[position] != rune('M') {
																goto l413
															}
															position++
														}
													l418:
														if buffer[position] != rune(' ') {
															goto l413
														}
														position++
														if buffer[position] != rune('2') {
															goto l413
														}
														position++
														add(rulePegText, position415)
													}
													{
														add(ruleAction80, position)
													}
													add(ruleIm2, position414)
												}
												goto l362
											l413:
												position, tokenIndex = position362, tokenIndex362
												{
													switch buffer[position] {
													case 'I', 'O', 'i', 'o':
														{
															position422 := position
															{
																position423, tokenIndex423 := position, tokenIndex
																{
																	position425 := position
																	{
																		position426 := position
																		{
																			position427, tokenIndex427 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l428
																			}
																			position++
																			goto l427
																		l428:
																			position, tokenIndex = position427, tokenIndex427
																			if buffer[position] != rune('I') {
																				goto l424
																			}
																			position++
																		}
																	l427:
																		{
																			position429, tokenIndex429 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l430
																			}
																			position++
																			goto l429
																		l430:
																			position, tokenIndex = position429, tokenIndex429
																			if buffer[position] != rune('N') {
																				goto l424
																			}
																			position++
																		}
																	l429:
																		{
																			position431, tokenIndex431 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l432
																			}
																			position++
																			goto l431
																		l432:
																			position, tokenIndex = position431, tokenIndex431
																			if buffer[position] != rune('I') {
																				goto l424
																			}
																			position++
																		}
																	l431:
																		{
																			position433, tokenIndex433 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l434
																			}
																			position++
																			goto l433
																		l434:
																			position, tokenIndex = position433, tokenIndex433
																			if buffer[position] != rune('R') {
																				goto l424
																			}
																			position++
																		}
																	l433:
																		add(rulePegText, position426)
																	}
																	{
																		add(ruleAction91, position)
																	}
																	add(ruleInir, position425)
																}
																goto l423
															l424:
																position, tokenIndex = position423, tokenIndex423
																{
																	position437 := position
																	{
																		position438 := position
																		{
																			position439, tokenIndex439 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l440
																			}
																			position++
																			goto l439
																		l440:
																			position, tokenIndex = position439, tokenIndex439
																			if buffer[position] != rune('I') {
																				goto l436
																			}
																			position++
																		}
																	l439:
																		{
																			position441, tokenIndex441 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l442
																			}
																			position++
																			goto l441
																		l442:
																			position, tokenIndex = position441, tokenIndex441
																			if buffer[position] != rune('N') {
																				goto l436
																			}
																			position++
																		}
																	l441:
																		{
																			position443, tokenIndex443 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l444
																			}
																			position++
																			goto l443
																		l444:
																			position, tokenIndex = position443, tokenIndex443
																			if buffer[position] != rune('I') {
																				goto l436
																			}
																			position++
																		}
																	l443:
																		add(rulePegText, position438)
																	}
																	{
																		add(ruleAction83, position)
																	}
																	add(ruleIni, position437)
																}
																goto l423
															l436:
																position, tokenIndex = position423, tokenIndex423
																{
																	position447 := position
																	{
																		position448 := position
																		{
																			position449, tokenIndex449 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l450
																			}
																			position++
																			goto l449
																		l450:
																			position, tokenIndex = position449, tokenIndex449
																			if buffer[position] != rune('O') {
																				goto l446
																			}
																			position++
																		}
																	l449:
																		{
																			position451, tokenIndex451 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l452
																			}
																			position++
																			goto l451
																		l452:
																			position, tokenIndex = position451, tokenIndex451
																			if buffer[position] != rune('T') {
																				goto l446
																			}
																			position++
																		}
																	l451:
																		{
																			position453, tokenIndex453 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l454
																			}
																			position++
																			goto l453
																		l454:
																			position, tokenIndex = position453, tokenIndex453
																			if buffer[position] != rune('I') {
																				goto l446
																			}
																			position++
																		}
																	l453:
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
																				goto l446
																			}
																			position++
																		}
																	l455:
																		add(rulePegText, position448)
																	}
																	{
																		add(ruleAction92, position)
																	}
																	add(ruleOtir, position447)
																}
																goto l423
															l446:
																position, tokenIndex = position423, tokenIndex423
																{
																	position459 := position
																	{
																		position460 := position
																		{
																			position461, tokenIndex461 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l462
																			}
																			position++
																			goto l461
																		l462:
																			position, tokenIndex = position461, tokenIndex461
																			if buffer[position] != rune('O') {
																				goto l458
																			}
																			position++
																		}
																	l461:
																		{
																			position463, tokenIndex463 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l464
																			}
																			position++
																			goto l463
																		l464:
																			position, tokenIndex = position463, tokenIndex463
																			if buffer[position] != rune('U') {
																				goto l458
																			}
																			position++
																		}
																	l463:
																		{
																			position465, tokenIndex465 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l466
																			}
																			position++
																			goto l465
																		l466:
																			position, tokenIndex = position465, tokenIndex465
																			if buffer[position] != rune('T') {
																				goto l458
																			}
																			position++
																		}
																	l465:
																		{
																			position467, tokenIndex467 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l468
																			}
																			position++
																			goto l467
																		l468:
																			position, tokenIndex = position467, tokenIndex467
																			if buffer[position] != rune('I') {
																				goto l458
																			}
																			position++
																		}
																	l467:
																		add(rulePegText, position460)
																	}
																	{
																		add(ruleAction84, position)
																	}
																	add(ruleOuti, position459)
																}
																goto l423
															l458:
																position, tokenIndex = position423, tokenIndex423
																{
																	position471 := position
																	{
																		position472 := position
																		{
																			position473, tokenIndex473 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l474
																			}
																			position++
																			goto l473
																		l474:
																			position, tokenIndex = position473, tokenIndex473
																			if buffer[position] != rune('I') {
																				goto l470
																			}
																			position++
																		}
																	l473:
																		{
																			position475, tokenIndex475 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l476
																			}
																			position++
																			goto l475
																		l476:
																			position, tokenIndex = position475, tokenIndex475
																			if buffer[position] != rune('N') {
																				goto l470
																			}
																			position++
																		}
																	l475:
																		{
																			position477, tokenIndex477 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l478
																			}
																			position++
																			goto l477
																		l478:
																			position, tokenIndex = position477, tokenIndex477
																			if buffer[position] != rune('D') {
																				goto l470
																			}
																			position++
																		}
																	l477:
																		{
																			position479, tokenIndex479 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l480
																			}
																			position++
																			goto l479
																		l480:
																			position, tokenIndex = position479, tokenIndex479
																			if buffer[position] != rune('R') {
																				goto l470
																			}
																			position++
																		}
																	l479:
																		add(rulePegText, position472)
																	}
																	{
																		add(ruleAction95, position)
																	}
																	add(ruleIndr, position471)
																}
																goto l423
															l470:
																position, tokenIndex = position423, tokenIndex423
																{
																	position483 := position
																	{
																		position484 := position
																		{
																			position485, tokenIndex485 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l486
																			}
																			position++
																			goto l485
																		l486:
																			position, tokenIndex = position485, tokenIndex485
																			if buffer[position] != rune('I') {
																				goto l482
																			}
																			position++
																		}
																	l485:
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
																				goto l482
																			}
																			position++
																		}
																	l487:
																		{
																			position489, tokenIndex489 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l490
																			}
																			position++
																			goto l489
																		l490:
																			position, tokenIndex = position489, tokenIndex489
																			if buffer[position] != rune('D') {
																				goto l482
																			}
																			position++
																		}
																	l489:
																		add(rulePegText, position484)
																	}
																	{
																		add(ruleAction87, position)
																	}
																	add(ruleInd, position483)
																}
																goto l423
															l482:
																position, tokenIndex = position423, tokenIndex423
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
																				goto l492
																			}
																			position++
																		}
																	l495:
																		{
																			position497, tokenIndex497 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l498
																			}
																			position++
																			goto l497
																		l498:
																			position, tokenIndex = position497, tokenIndex497
																			if buffer[position] != rune('T') {
																				goto l492
																			}
																			position++
																		}
																	l497:
																		{
																			position499, tokenIndex499 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l500
																			}
																			position++
																			goto l499
																		l500:
																			position, tokenIndex = position499, tokenIndex499
																			if buffer[position] != rune('D') {
																				goto l492
																			}
																			position++
																		}
																	l499:
																		{
																			position501, tokenIndex501 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l502
																			}
																			position++
																			goto l501
																		l502:
																			position, tokenIndex = position501, tokenIndex501
																			if buffer[position] != rune('R') {
																				goto l492
																			}
																			position++
																		}
																	l501:
																		add(rulePegText, position494)
																	}
																	{
																		add(ruleAction96, position)
																	}
																	add(ruleOtdr, position493)
																}
																goto l423
															l492:
																position, tokenIndex = position423, tokenIndex423
																{
																	position504 := position
																	{
																		position505 := position
																		{
																			position506, tokenIndex506 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l507
																			}
																			position++
																			goto l506
																		l507:
																			position, tokenIndex = position506, tokenIndex506
																			if buffer[position] != rune('O') {
																				goto l360
																			}
																			position++
																		}
																	l506:
																		{
																			position508, tokenIndex508 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l509
																			}
																			position++
																			goto l508
																		l509:
																			position, tokenIndex = position508, tokenIndex508
																			if buffer[position] != rune('U') {
																				goto l360
																			}
																			position++
																		}
																	l508:
																		{
																			position510, tokenIndex510 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l511
																			}
																			position++
																			goto l510
																		l511:
																			position, tokenIndex = position510, tokenIndex510
																			if buffer[position] != rune('T') {
																				goto l360
																			}
																			position++
																		}
																	l510:
																		{
																			position512, tokenIndex512 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l513
																			}
																			position++
																			goto l512
																		l513:
																			position, tokenIndex = position512, tokenIndex512
																			if buffer[position] != rune('D') {
																				goto l360
																			}
																			position++
																		}
																	l512:
																		add(rulePegText, position505)
																	}
																	{
																		add(ruleAction88, position)
																	}
																	add(ruleOutd, position504)
																}
															}
														l423:
															add(ruleBlitIO, position422)
														}
														break
													case 'R', 'r':
														{
															position515 := position
															{
																position516 := position
																{
																	position517, tokenIndex517 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l518
																	}
																	position++
																	goto l517
																l518:
																	position, tokenIndex = position517, tokenIndex517
																	if buffer[position] != rune('R') {
																		goto l360
																	}
																	position++
																}
															l517:
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
																		goto l360
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
																		goto l360
																	}
																	position++
																}
															l521:
																add(rulePegText, position516)
															}
															{
																add(ruleAction77, position)
															}
															add(ruleRld, position515)
														}
														break
													case 'N', 'n':
														{
															position524 := position
															{
																position525 := position
																{
																	position526, tokenIndex526 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l527
																	}
																	position++
																	goto l526
																l527:
																	position, tokenIndex = position526, tokenIndex526
																	if buffer[position] != rune('N') {
																		goto l360
																	}
																	position++
																}
															l526:
																{
																	position528, tokenIndex528 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l529
																	}
																	position++
																	goto l528
																l529:
																	position, tokenIndex = position528, tokenIndex528
																	if buffer[position] != rune('E') {
																		goto l360
																	}
																	position++
																}
															l528:
																{
																	position530, tokenIndex530 := position, tokenIndex
																	if buffer[position] != rune('g') {
																		goto l531
																	}
																	position++
																	goto l530
																l531:
																	position, tokenIndex = position530, tokenIndex530
																	if buffer[position] != rune('G') {
																		goto l360
																	}
																	position++
																}
															l530:
																add(rulePegText, position525)
															}
															{
																add(ruleAction73, position)
															}
															add(ruleNeg, position524)
														}
														break
													default:
														{
															position533 := position
															{
																position534, tokenIndex534 := position, tokenIndex
																{
																	position536 := position
																	{
																		position537 := position
																		{
																			position538, tokenIndex538 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l539
																			}
																			position++
																			goto l538
																		l539:
																			position, tokenIndex = position538, tokenIndex538
																			if buffer[position] != rune('L') {
																				goto l535
																			}
																			position++
																		}
																	l538:
																		{
																			position540, tokenIndex540 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l541
																			}
																			position++
																			goto l540
																		l541:
																			position, tokenIndex = position540, tokenIndex540
																			if buffer[position] != rune('D') {
																				goto l535
																			}
																			position++
																		}
																	l540:
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
																				goto l535
																			}
																			position++
																		}
																	l542:
																		{
																			position544, tokenIndex544 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l545
																			}
																			position++
																			goto l544
																		l545:
																			position, tokenIndex = position544, tokenIndex544
																			if buffer[position] != rune('R') {
																				goto l535
																			}
																			position++
																		}
																	l544:
																		add(rulePegText, position537)
																	}
																	{
																		add(ruleAction89, position)
																	}
																	add(ruleLdir, position536)
																}
																goto l534
															l535:
																position, tokenIndex = position534, tokenIndex534
																{
																	position548 := position
																	{
																		position549 := position
																		{
																			position550, tokenIndex550 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l551
																			}
																			position++
																			goto l550
																		l551:
																			position, tokenIndex = position550, tokenIndex550
																			if buffer[position] != rune('L') {
																				goto l547
																			}
																			position++
																		}
																	l550:
																		{
																			position552, tokenIndex552 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l553
																			}
																			position++
																			goto l552
																		l553:
																			position, tokenIndex = position552, tokenIndex552
																			if buffer[position] != rune('D') {
																				goto l547
																			}
																			position++
																		}
																	l552:
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
																				goto l547
																			}
																			position++
																		}
																	l554:
																		add(rulePegText, position549)
																	}
																	{
																		add(ruleAction81, position)
																	}
																	add(ruleLdi, position548)
																}
																goto l534
															l547:
																position, tokenIndex = position534, tokenIndex534
																{
																	position558 := position
																	{
																		position559 := position
																		{
																			position560, tokenIndex560 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l561
																			}
																			position++
																			goto l560
																		l561:
																			position, tokenIndex = position560, tokenIndex560
																			if buffer[position] != rune('C') {
																				goto l557
																			}
																			position++
																		}
																	l560:
																		{
																			position562, tokenIndex562 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l563
																			}
																			position++
																			goto l562
																		l563:
																			position, tokenIndex = position562, tokenIndex562
																			if buffer[position] != rune('P') {
																				goto l557
																			}
																			position++
																		}
																	l562:
																		{
																			position564, tokenIndex564 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l565
																			}
																			position++
																			goto l564
																		l565:
																			position, tokenIndex = position564, tokenIndex564
																			if buffer[position] != rune('I') {
																				goto l557
																			}
																			position++
																		}
																	l564:
																		{
																			position566, tokenIndex566 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l567
																			}
																			position++
																			goto l566
																		l567:
																			position, tokenIndex = position566, tokenIndex566
																			if buffer[position] != rune('R') {
																				goto l557
																			}
																			position++
																		}
																	l566:
																		add(rulePegText, position559)
																	}
																	{
																		add(ruleAction90, position)
																	}
																	add(ruleCpir, position558)
																}
																goto l534
															l557:
																position, tokenIndex = position534, tokenIndex534
																{
																	position570 := position
																	{
																		position571 := position
																		{
																			position572, tokenIndex572 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l573
																			}
																			position++
																			goto l572
																		l573:
																			position, tokenIndex = position572, tokenIndex572
																			if buffer[position] != rune('C') {
																				goto l569
																			}
																			position++
																		}
																	l572:
																		{
																			position574, tokenIndex574 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l575
																			}
																			position++
																			goto l574
																		l575:
																			position, tokenIndex = position574, tokenIndex574
																			if buffer[position] != rune('P') {
																				goto l569
																			}
																			position++
																		}
																	l574:
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
																				goto l569
																			}
																			position++
																		}
																	l576:
																		add(rulePegText, position571)
																	}
																	{
																		add(ruleAction82, position)
																	}
																	add(ruleCpi, position570)
																}
																goto l534
															l569:
																position, tokenIndex = position534, tokenIndex534
																{
																	position580 := position
																	{
																		position581 := position
																		{
																			position582, tokenIndex582 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l583
																			}
																			position++
																			goto l582
																		l583:
																			position, tokenIndex = position582, tokenIndex582
																			if buffer[position] != rune('L') {
																				goto l579
																			}
																			position++
																		}
																	l582:
																		{
																			position584, tokenIndex584 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l585
																			}
																			position++
																			goto l584
																		l585:
																			position, tokenIndex = position584, tokenIndex584
																			if buffer[position] != rune('D') {
																				goto l579
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
																				goto l579
																			}
																			position++
																		}
																	l586:
																		{
																			position588, tokenIndex588 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l589
																			}
																			position++
																			goto l588
																		l589:
																			position, tokenIndex = position588, tokenIndex588
																			if buffer[position] != rune('R') {
																				goto l579
																			}
																			position++
																		}
																	l588:
																		add(rulePegText, position581)
																	}
																	{
																		add(ruleAction93, position)
																	}
																	add(ruleLddr, position580)
																}
																goto l534
															l579:
																position, tokenIndex = position534, tokenIndex534
																{
																	position592 := position
																	{
																		position593 := position
																		{
																			position594, tokenIndex594 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l595
																			}
																			position++
																			goto l594
																		l595:
																			position, tokenIndex = position594, tokenIndex594
																			if buffer[position] != rune('L') {
																				goto l591
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
																				goto l591
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
																				goto l591
																			}
																			position++
																		}
																	l598:
																		add(rulePegText, position593)
																	}
																	{
																		add(ruleAction85, position)
																	}
																	add(ruleLdd, position592)
																}
																goto l534
															l591:
																position, tokenIndex = position534, tokenIndex534
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
																				goto l601
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
																				goto l601
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
																				goto l601
																			}
																			position++
																		}
																	l608:
																		{
																			position610, tokenIndex610 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l611
																			}
																			position++
																			goto l610
																		l611:
																			position, tokenIndex = position610, tokenIndex610
																			if buffer[position] != rune('R') {
																				goto l601
																			}
																			position++
																		}
																	l610:
																		add(rulePegText, position603)
																	}
																	{
																		add(ruleAction94, position)
																	}
																	add(ruleCpdr, position602)
																}
																goto l534
															l601:
																position, tokenIndex = position534, tokenIndex534
																{
																	position613 := position
																	{
																		position614 := position
																		{
																			position615, tokenIndex615 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l616
																			}
																			position++
																			goto l615
																		l616:
																			position, tokenIndex = position615, tokenIndex615
																			if buffer[position] != rune('C') {
																				goto l360
																			}
																			position++
																		}
																	l615:
																		{
																			position617, tokenIndex617 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l618
																			}
																			position++
																			goto l617
																		l618:
																			position, tokenIndex = position617, tokenIndex617
																			if buffer[position] != rune('P') {
																				goto l360
																			}
																			position++
																		}
																	l617:
																		{
																			position619, tokenIndex619 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l620
																			}
																			position++
																			goto l619
																		l620:
																			position, tokenIndex = position619, tokenIndex619
																			if buffer[position] != rune('D') {
																				goto l360
																			}
																			position++
																		}
																	l619:
																		add(rulePegText, position614)
																	}
																	{
																		add(ruleAction86, position)
																	}
																	add(ruleCpd, position613)
																}
															}
														l534:
															add(ruleBlit, position533)
														}
														break
													}
												}

											}
										l362:
											add(ruleEDSimple, position361)
										}
										goto l32
									l360:
										position, tokenIndex = position32, tokenIndex32
										{
											position623 := position
											{
												position624, tokenIndex624 := position, tokenIndex
												{
													position626 := position
													{
														position627 := position
														{
															position628, tokenIndex628 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l629
															}
															position++
															goto l628
														l629:
															position, tokenIndex = position628, tokenIndex628
															if buffer[position] != rune('R') {
																goto l625
															}
															position++
														}
													l628:
														{
															position630, tokenIndex630 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l631
															}
															position++
															goto l630
														l631:
															position, tokenIndex = position630, tokenIndex630
															if buffer[position] != rune('L') {
																goto l625
															}
															position++
														}
													l630:
														{
															position632, tokenIndex632 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l633
															}
															position++
															goto l632
														l633:
															position, tokenIndex = position632, tokenIndex632
															if buffer[position] != rune('C') {
																goto l625
															}
															position++
														}
													l632:
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
																goto l625
															}
															position++
														}
													l634:
														add(rulePegText, position627)
													}
													{
														add(ruleAction62, position)
													}
													add(ruleRlca, position626)
												}
												goto l624
											l625:
												position, tokenIndex = position624, tokenIndex624
												{
													position638 := position
													{
														position639 := position
														{
															position640, tokenIndex640 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l641
															}
															position++
															goto l640
														l641:
															position, tokenIndex = position640, tokenIndex640
															if buffer[position] != rune('R') {
																goto l637
															}
															position++
														}
													l640:
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
																goto l637
															}
															position++
														}
													l642:
														{
															position644, tokenIndex644 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l645
															}
															position++
															goto l644
														l645:
															position, tokenIndex = position644, tokenIndex644
															if buffer[position] != rune('C') {
																goto l637
															}
															position++
														}
													l644:
														{
															position646, tokenIndex646 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l647
															}
															position++
															goto l646
														l647:
															position, tokenIndex = position646, tokenIndex646
															if buffer[position] != rune('A') {
																goto l637
															}
															position++
														}
													l646:
														add(rulePegText, position639)
													}
													{
														add(ruleAction63, position)
													}
													add(ruleRrca, position638)
												}
												goto l624
											l637:
												position, tokenIndex = position624, tokenIndex624
												{
													position650 := position
													{
														position651 := position
														{
															position652, tokenIndex652 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l653
															}
															position++
															goto l652
														l653:
															position, tokenIndex = position652, tokenIndex652
															if buffer[position] != rune('R') {
																goto l649
															}
															position++
														}
													l652:
														{
															position654, tokenIndex654 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l655
															}
															position++
															goto l654
														l655:
															position, tokenIndex = position654, tokenIndex654
															if buffer[position] != rune('L') {
																goto l649
															}
															position++
														}
													l654:
														{
															position656, tokenIndex656 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l657
															}
															position++
															goto l656
														l657:
															position, tokenIndex = position656, tokenIndex656
															if buffer[position] != rune('A') {
																goto l649
															}
															position++
														}
													l656:
														add(rulePegText, position651)
													}
													{
														add(ruleAction64, position)
													}
													add(ruleRla, position650)
												}
												goto l624
											l649:
												position, tokenIndex = position624, tokenIndex624
												{
													position660 := position
													{
														position661 := position
														{
															position662, tokenIndex662 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l663
															}
															position++
															goto l662
														l663:
															position, tokenIndex = position662, tokenIndex662
															if buffer[position] != rune('D') {
																goto l659
															}
															position++
														}
													l662:
														{
															position664, tokenIndex664 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l665
															}
															position++
															goto l664
														l665:
															position, tokenIndex = position664, tokenIndex664
															if buffer[position] != rune('A') {
																goto l659
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
																goto l659
															}
															position++
														}
													l666:
														add(rulePegText, position661)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleDaa, position660)
												}
												goto l624
											l659:
												position, tokenIndex = position624, tokenIndex624
												{
													position670 := position
													{
														position671 := position
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
																goto l669
															}
															position++
														}
													l672:
														{
															position674, tokenIndex674 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l675
															}
															position++
															goto l674
														l675:
															position, tokenIndex = position674, tokenIndex674
															if buffer[position] != rune('P') {
																goto l669
															}
															position++
														}
													l674:
														{
															position676, tokenIndex676 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l677
															}
															position++
															goto l676
														l677:
															position, tokenIndex = position676, tokenIndex676
															if buffer[position] != rune('L') {
																goto l669
															}
															position++
														}
													l676:
														add(rulePegText, position671)
													}
													{
														add(ruleAction67, position)
													}
													add(ruleCpl, position670)
												}
												goto l624
											l669:
												position, tokenIndex = position624, tokenIndex624
												{
													position680 := position
													{
														position681 := position
														{
															position682, tokenIndex682 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l683
															}
															position++
															goto l682
														l683:
															position, tokenIndex = position682, tokenIndex682
															if buffer[position] != rune('E') {
																goto l679
															}
															position++
														}
													l682:
														{
															position684, tokenIndex684 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l685
															}
															position++
															goto l684
														l685:
															position, tokenIndex = position684, tokenIndex684
															if buffer[position] != rune('X') {
																goto l679
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
																goto l679
															}
															position++
														}
													l686:
														add(rulePegText, position681)
													}
													{
														add(ruleAction70, position)
													}
													add(ruleExx, position680)
												}
												goto l624
											l679:
												position, tokenIndex = position624, tokenIndex624
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position690 := position
															{
																position691 := position
																{
																	position692, tokenIndex692 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l693
																	}
																	position++
																	goto l692
																l693:
																	position, tokenIndex = position692, tokenIndex692
																	if buffer[position] != rune('E') {
																		goto l622
																	}
																	position++
																}
															l692:
																{
																	position694, tokenIndex694 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l695
																	}
																	position++
																	goto l694
																l695:
																	position, tokenIndex = position694, tokenIndex694
																	if buffer[position] != rune('I') {
																		goto l622
																	}
																	position++
																}
															l694:
																add(rulePegText, position691)
															}
															{
																add(ruleAction72, position)
															}
															add(ruleEi, position690)
														}
														break
													case 'D', 'd':
														{
															position697 := position
															{
																position698 := position
																{
																	position699, tokenIndex699 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l700
																	}
																	position++
																	goto l699
																l700:
																	position, tokenIndex = position699, tokenIndex699
																	if buffer[position] != rune('D') {
																		goto l622
																	}
																	position++
																}
															l699:
																{
																	position701, tokenIndex701 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l702
																	}
																	position++
																	goto l701
																l702:
																	position, tokenIndex = position701, tokenIndex701
																	if buffer[position] != rune('I') {
																		goto l622
																	}
																	position++
																}
															l701:
																add(rulePegText, position698)
															}
															{
																add(ruleAction71, position)
															}
															add(ruleDi, position697)
														}
														break
													case 'C', 'c':
														{
															position704 := position
															{
																position705 := position
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
																		goto l622
																	}
																	position++
																}
															l706:
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
																		goto l622
																	}
																	position++
																}
															l708:
																{
																	position710, tokenIndex710 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l711
																	}
																	position++
																	goto l710
																l711:
																	position, tokenIndex = position710, tokenIndex710
																	if buffer[position] != rune('F') {
																		goto l622
																	}
																	position++
																}
															l710:
																add(rulePegText, position705)
															}
															{
																add(ruleAction69, position)
															}
															add(ruleCcf, position704)
														}
														break
													case 'S', 's':
														{
															position713 := position
															{
																position714 := position
																{
																	position715, tokenIndex715 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l716
																	}
																	position++
																	goto l715
																l716:
																	position, tokenIndex = position715, tokenIndex715
																	if buffer[position] != rune('S') {
																		goto l622
																	}
																	position++
																}
															l715:
																{
																	position717, tokenIndex717 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l718
																	}
																	position++
																	goto l717
																l718:
																	position, tokenIndex = position717, tokenIndex717
																	if buffer[position] != rune('C') {
																		goto l622
																	}
																	position++
																}
															l717:
																{
																	position719, tokenIndex719 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l720
																	}
																	position++
																	goto l719
																l720:
																	position, tokenIndex = position719, tokenIndex719
																	if buffer[position] != rune('F') {
																		goto l622
																	}
																	position++
																}
															l719:
																add(rulePegText, position714)
															}
															{
																add(ruleAction68, position)
															}
															add(ruleScf, position713)
														}
														break
													case 'R', 'r':
														{
															position722 := position
															{
																position723 := position
																{
																	position724, tokenIndex724 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l725
																	}
																	position++
																	goto l724
																l725:
																	position, tokenIndex = position724, tokenIndex724
																	if buffer[position] != rune('R') {
																		goto l622
																	}
																	position++
																}
															l724:
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
																		goto l622
																	}
																	position++
																}
															l726:
																{
																	position728, tokenIndex728 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l729
																	}
																	position++
																	goto l728
																l729:
																	position, tokenIndex = position728, tokenIndex728
																	if buffer[position] != rune('A') {
																		goto l622
																	}
																	position++
																}
															l728:
																add(rulePegText, position723)
															}
															{
																add(ruleAction65, position)
															}
															add(ruleRra, position722)
														}
														break
													case 'H', 'h':
														{
															position731 := position
															{
																position732 := position
																{
																	position733, tokenIndex733 := position, tokenIndex
																	if buffer[position] != rune('h') {
																		goto l734
																	}
																	position++
																	goto l733
																l734:
																	position, tokenIndex = position733, tokenIndex733
																	if buffer[position] != rune('H') {
																		goto l622
																	}
																	position++
																}
															l733:
																{
																	position735, tokenIndex735 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l736
																	}
																	position++
																	goto l735
																l736:
																	position, tokenIndex = position735, tokenIndex735
																	if buffer[position] != rune('A') {
																		goto l622
																	}
																	position++
																}
															l735:
																{
																	position737, tokenIndex737 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l738
																	}
																	position++
																	goto l737
																l738:
																	position, tokenIndex = position737, tokenIndex737
																	if buffer[position] != rune('L') {
																		goto l622
																	}
																	position++
																}
															l737:
																{
																	position739, tokenIndex739 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l740
																	}
																	position++
																	goto l739
																l740:
																	position, tokenIndex = position739, tokenIndex739
																	if buffer[position] != rune('T') {
																		goto l622
																	}
																	position++
																}
															l739:
																add(rulePegText, position732)
															}
															{
																add(ruleAction61, position)
															}
															add(ruleHalt, position731)
														}
														break
													default:
														{
															position742 := position
															{
																position743 := position
																{
																	position744, tokenIndex744 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l745
																	}
																	position++
																	goto l744
																l745:
																	position, tokenIndex = position744, tokenIndex744
																	if buffer[position] != rune('N') {
																		goto l622
																	}
																	position++
																}
															l744:
																{
																	position746, tokenIndex746 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l747
																	}
																	position++
																	goto l746
																l747:
																	position, tokenIndex = position746, tokenIndex746
																	if buffer[position] != rune('O') {
																		goto l622
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
																		goto l622
																	}
																	position++
																}
															l748:
																add(rulePegText, position743)
															}
															{
																add(ruleAction60, position)
															}
															add(ruleNop, position742)
														}
														break
													}
												}

											}
										l624:
											add(ruleSimple, position623)
										}
										goto l32
									l622:
										position, tokenIndex = position32, tokenIndex32
										{
											position752 := position
											{
												position753, tokenIndex753 := position, tokenIndex
												{
													position755 := position
													{
														position756, tokenIndex756 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l757
														}
														position++
														goto l756
													l757:
														position, tokenIndex = position756, tokenIndex756
														if buffer[position] != rune('R') {
															goto l754
														}
														position++
													}
												l756:
													{
														position758, tokenIndex758 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l759
														}
														position++
														goto l758
													l759:
														position, tokenIndex = position758, tokenIndex758
														if buffer[position] != rune('S') {
															goto l754
														}
														position++
													}
												l758:
													{
														position760, tokenIndex760 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l761
														}
														position++
														goto l760
													l761:
														position, tokenIndex = position760, tokenIndex760
														if buffer[position] != rune('T') {
															goto l754
														}
														position++
													}
												l760:
													if !_rules[rulews]() {
														goto l754
													}
													if !_rules[rulen]() {
														goto l754
													}
													{
														add(ruleAction97, position)
													}
													add(ruleRst, position755)
												}
												goto l753
											l754:
												position, tokenIndex = position753, tokenIndex753
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
															goto l763
														}
														position++
													}
												l765:
													{
														position767, tokenIndex767 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l768
														}
														position++
														goto l767
													l768:
														position, tokenIndex = position767, tokenIndex767
														if buffer[position] != rune('P') {
															goto l763
														}
														position++
													}
												l767:
													if !_rules[rulews]() {
														goto l763
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
													if !_rules[ruleSrc16]() {
														goto l763
													}
													{
														add(ruleAction100, position)
													}
													add(ruleJp, position764)
												}
												goto l753
											l763:
												position, tokenIndex = position753, tokenIndex753
												{
													switch buffer[position] {
													case 'D', 'd':
														{
															position773 := position
															{
																position774, tokenIndex774 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l775
																}
																position++
																goto l774
															l775:
																position, tokenIndex = position774, tokenIndex774
																if buffer[position] != rune('D') {
																	goto l751
																}
																position++
															}
														l774:
															{
																position776, tokenIndex776 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l777
																}
																position++
																goto l776
															l777:
																position, tokenIndex = position776, tokenIndex776
																if buffer[position] != rune('J') {
																	goto l751
																}
																position++
															}
														l776:
															{
																position778, tokenIndex778 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l779
																}
																position++
																goto l778
															l779:
																position, tokenIndex = position778, tokenIndex778
																if buffer[position] != rune('N') {
																	goto l751
																}
																position++
															}
														l778:
															{
																position780, tokenIndex780 := position, tokenIndex
																if buffer[position] != rune('z') {
																	goto l781
																}
																position++
																goto l780
															l781:
																position, tokenIndex = position780, tokenIndex780
																if buffer[position] != rune('Z') {
																	goto l751
																}
																position++
															}
														l780:
															if !_rules[rulews]() {
																goto l751
															}
															if !_rules[ruledisp]() {
																goto l751
															}
															{
																add(ruleAction102, position)
															}
															add(ruleDjnz, position773)
														}
														break
													case 'J', 'j':
														{
															position783 := position
															{
																position784, tokenIndex784 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l785
																}
																position++
																goto l784
															l785:
																position, tokenIndex = position784, tokenIndex784
																if buffer[position] != rune('J') {
																	goto l751
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
																	goto l751
																}
																position++
															}
														l786:
															if !_rules[rulews]() {
																goto l751
															}
															{
																position788, tokenIndex788 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l788
																}
																if !_rules[rulesep]() {
																	goto l788
																}
																goto l789
															l788:
																position, tokenIndex = position788, tokenIndex788
															}
														l789:
															if !_rules[ruledisp]() {
																goto l751
															}
															{
																add(ruleAction101, position)
															}
															add(ruleJr, position783)
														}
														break
													case 'R', 'r':
														{
															position791 := position
															{
																position792, tokenIndex792 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l793
																}
																position++
																goto l792
															l793:
																position, tokenIndex = position792, tokenIndex792
																if buffer[position] != rune('R') {
																	goto l751
																}
																position++
															}
														l792:
															{
																position794, tokenIndex794 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l795
																}
																position++
																goto l794
															l795:
																position, tokenIndex = position794, tokenIndex794
																if buffer[position] != rune('E') {
																	goto l751
																}
																position++
															}
														l794:
															{
																position796, tokenIndex796 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l797
																}
																position++
																goto l796
															l797:
																position, tokenIndex = position796, tokenIndex796
																if buffer[position] != rune('T') {
																	goto l751
																}
																position++
															}
														l796:
															{
																position798, tokenIndex798 := position, tokenIndex
																if !_rules[rulews]() {
																	goto l798
																}
																if !_rules[rulecc]() {
																	goto l798
																}
																goto l799
															l798:
																position, tokenIndex = position798, tokenIndex798
															}
														l799:
															{
																add(ruleAction99, position)
															}
															add(ruleRet, position791)
														}
														break
													default:
														{
															position801 := position
															{
																position802, tokenIndex802 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l803
																}
																position++
																goto l802
															l803:
																position, tokenIndex = position802, tokenIndex802
																if buffer[position] != rune('C') {
																	goto l751
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
																	goto l751
																}
																position++
															}
														l804:
															{
																position806, tokenIndex806 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l807
																}
																position++
																goto l806
															l807:
																position, tokenIndex = position806, tokenIndex806
																if buffer[position] != rune('L') {
																	goto l751
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
																	goto l751
																}
																position++
															}
														l808:
															if !_rules[rulews]() {
																goto l751
															}
															{
																position810, tokenIndex810 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l810
																}
																if !_rules[rulesep]() {
																	goto l810
																}
																goto l811
															l810:
																position, tokenIndex = position810, tokenIndex810
															}
														l811:
															if !_rules[ruleSrc16]() {
																goto l751
															}
															{
																add(ruleAction98, position)
															}
															add(ruleCall, position801)
														}
														break
													}
												}

											}
										l753:
											add(ruleJump, position752)
										}
										goto l32
									l751:
										position, tokenIndex = position32, tokenIndex32
										{
											position813 := position
											{
												position814, tokenIndex814 := position, tokenIndex
												{
													position816 := position
													{
														position817, tokenIndex817 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l818
														}
														position++
														goto l817
													l818:
														position, tokenIndex = position817, tokenIndex817
														if buffer[position] != rune('I') {
															goto l815
														}
														position++
													}
												l817:
													{
														position819, tokenIndex819 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l820
														}
														position++
														goto l819
													l820:
														position, tokenIndex = position819, tokenIndex819
														if buffer[position] != rune('N') {
															goto l815
														}
														position++
													}
												l819:
													if !_rules[rulews]() {
														goto l815
													}
													if !_rules[ruleReg8]() {
														goto l815
													}
													if !_rules[rulesep]() {
														goto l815
													}
													if !_rules[rulePort]() {
														goto l815
													}
													{
														add(ruleAction103, position)
													}
													add(ruleIN, position816)
												}
												goto l814
											l815:
												position, tokenIndex = position814, tokenIndex814
												{
													position822 := position
													{
														position823, tokenIndex823 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l824
														}
														position++
														goto l823
													l824:
														position, tokenIndex = position823, tokenIndex823
														if buffer[position] != rune('O') {
															goto l11
														}
														position++
													}
												l823:
													{
														position825, tokenIndex825 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l826
														}
														position++
														goto l825
													l826:
														position, tokenIndex = position825, tokenIndex825
														if buffer[position] != rune('U') {
															goto l11
														}
														position++
													}
												l825:
													{
														position827, tokenIndex827 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l828
														}
														position++
														goto l827
													l828:
														position, tokenIndex = position827, tokenIndex827
														if buffer[position] != rune('T') {
															goto l11
														}
														position++
													}
												l827:
													if !_rules[rulews]() {
														goto l11
													}
													if !_rules[rulePort]() {
														goto l11
													}
													if !_rules[rulesep]() {
														goto l11
													}
													if !_rules[ruleReg8]() {
														goto l11
													}
													{
														add(ruleAction104, position)
													}
													add(ruleOUT, position822)
												}
											}
										l814:
											add(ruleIO, position813)
										}
									}
								l32:
									add(ruleInstruction, position31)
								}
							}
						l14:
							add(ruleStatement, position13)
						}
						goto l12
					l11:
						position, tokenIndex = position11, tokenIndex11
					}
				l12:
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
						{
							position834 := position
							{
								position835, tokenIndex835 := position, tokenIndex
								if buffer[position] != rune(';') {
									goto l836
								}
								position++
								goto l835
							l836:
								position, tokenIndex = position835, tokenIndex835
								if buffer[position] != rune('#') {
									goto l832
								}
								position++
							}
						l835:
						l837:
							{
								position838, tokenIndex838 := position, tokenIndex
								{
									position839, tokenIndex839 := position, tokenIndex
									if buffer[position] != rune('\n') {
										goto l839
									}
									position++
									goto l838
								l839:
									position, tokenIndex = position839, tokenIndex839
								}
								if !matchDot() {
									goto l838
								}
								goto l837
							l838:
								position, tokenIndex = position838, tokenIndex838
							}
							add(ruleComment, position834)
						}
						goto l833
					l832:
						position, tokenIndex = position832, tokenIndex832
					}
				l833:
					{
						position840, tokenIndex840 := position, tokenIndex
						if !_rules[rulews]() {
							goto l840
						}
						goto l841
					l840:
						position, tokenIndex = position840, tokenIndex840
					}
				l841:
					{
						position842, tokenIndex842 := position, tokenIndex
						{
							position844, tokenIndex844 := position, tokenIndex
							if buffer[position] != rune('\r') {
								goto l844
							}
							position++
							goto l845
						l844:
							position, tokenIndex = position844, tokenIndex844
						}
					l845:
						if buffer[position] != rune('\n') {
							goto l843
						}
						position++
						goto l842
					l843:
						position, tokenIndex = position842, tokenIndex842
						if buffer[position] != rune(':') {
							goto l0
						}
						position++
					}
				l842:
					{
						add(ruleAction0, position)
					}
					add(ruleLine, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position847 := position
						{
							position848, tokenIndex848 := position, tokenIndex
							{
								position850 := position
								if !_rules[ruleLabelText]() {
									goto l848
								}
								if buffer[position] != rune(':') {
									goto l848
								}
								position++
								if !_rules[rulews]() {
									goto l848
								}
								{
									add(ruleAction2, position)
								}
								add(ruleLabelDefn, position850)
							}
							goto l849
						l848:
							position, tokenIndex = position848, tokenIndex848
						}
					l849:
					l852:
						{
							position853, tokenIndex853 := position, tokenIndex
							if !_rules[rulews]() {
								goto l853
							}
							goto l852
						l853:
							position, tokenIndex = position853, tokenIndex853
						}
						{
							position854, tokenIndex854 := position, tokenIndex
							{
								position856 := position
								{
									position857, tokenIndex857 := position, tokenIndex
									{
										position859 := position
										{
											switch buffer[position] {
											case 'a':
												{
													position861 := position
													if buffer[position] != rune('a') {
														goto l858
													}
													position++
													if buffer[position] != rune('s') {
														goto l858
													}
													position++
													if buffer[position] != rune('e') {
														goto l858
													}
													position++
													if buffer[position] != rune('g') {
														goto l858
													}
													position++
													add(ruleAseg, position861)
												}
												break
											case '.':
												{
													position862 := position
													if buffer[position] != rune('.') {
														goto l858
													}
													position++
													if buffer[position] != rune('t') {
														goto l858
													}
													position++
													if buffer[position] != rune('i') {
														goto l858
													}
													position++
													if buffer[position] != rune('t') {
														goto l858
													}
													position++
													if buffer[position] != rune('l') {
														goto l858
													}
													position++
													if buffer[position] != rune('e') {
														goto l858
													}
													position++
													if !_rules[rulews]() {
														goto l858
													}
													if buffer[position] != rune('\'') {
														goto l858
													}
													position++
												l863:
													{
														position864, tokenIndex864 := position, tokenIndex
														{
															position865, tokenIndex865 := position, tokenIndex
															if buffer[position] != rune('\'') {
																goto l865
															}
															position++
															goto l864
														l865:
															position, tokenIndex = position865, tokenIndex865
														}
														if !matchDot() {
															goto l864
														}
														goto l863
													l864:
														position, tokenIndex = position864, tokenIndex864
													}
													if buffer[position] != rune('\'') {
														goto l858
													}
													position++
													add(ruleTitle, position862)
												}
												break
											default:
												{
													position866 := position
													{
														position867, tokenIndex867 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l868
														}
														position++
														goto l867
													l868:
														position, tokenIndex = position867, tokenIndex867
														if buffer[position] != rune('O') {
															goto l858
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
															goto l858
														}
														position++
													}
												l869:
													{
														position871, tokenIndex871 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l872
														}
														position++
														goto l871
													l872:
														position, tokenIndex = position871, tokenIndex871
														if buffer[position] != rune('G') {
															goto l858
														}
														position++
													}
												l871:
													if !_rules[rulews]() {
														goto l858
													}
													if !_rules[rulenn]() {
														goto l858
													}
													{
														add(ruleAction1, position)
													}
													add(ruleOrg, position866)
												}
												break
											}
										}

										add(ruleDirective, position859)
									}
									goto l857
								l858:
									position, tokenIndex = position857, tokenIndex857
									{
										position874 := position
										{
											position875, tokenIndex875 := position, tokenIndex
											{
												position877 := position
												{
													position878, tokenIndex878 := position, tokenIndex
													{
														position880 := position
														{
															position881, tokenIndex881 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l882
															}
															position++
															goto l881
														l882:
															position, tokenIndex = position881, tokenIndex881
															if buffer[position] != rune('P') {
																goto l879
															}
															position++
														}
													l881:
														{
															position883, tokenIndex883 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l884
															}
															position++
															goto l883
														l884:
															position, tokenIndex = position883, tokenIndex883
															if buffer[position] != rune('U') {
																goto l879
															}
															position++
														}
													l883:
														{
															position885, tokenIndex885 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l886
															}
															position++
															goto l885
														l886:
															position, tokenIndex = position885, tokenIndex885
															if buffer[position] != rune('S') {
																goto l879
															}
															position++
														}
													l885:
														{
															position887, tokenIndex887 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l888
															}
															position++
															goto l887
														l888:
															position, tokenIndex = position887, tokenIndex887
															if buffer[position] != rune('H') {
																goto l879
															}
															position++
														}
													l887:
														if !_rules[rulews]() {
															goto l879
														}
														if !_rules[ruleSrc16]() {
															goto l879
														}
														{
															add(ruleAction5, position)
														}
														add(rulePush, position880)
													}
													goto l878
												l879:
													position, tokenIndex = position878, tokenIndex878
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position891 := position
																{
																	position892, tokenIndex892 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l893
																	}
																	position++
																	goto l892
																l893:
																	position, tokenIndex = position892, tokenIndex892
																	if buffer[position] != rune('E') {
																		goto l876
																	}
																	position++
																}
															l892:
																{
																	position894, tokenIndex894 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l895
																	}
																	position++
																	goto l894
																l895:
																	position, tokenIndex = position894, tokenIndex894
																	if buffer[position] != rune('X') {
																		goto l876
																	}
																	position++
																}
															l894:
																if !_rules[rulews]() {
																	goto l876
																}
																if !_rules[ruleDst16]() {
																	goto l876
																}
																if !_rules[rulesep]() {
																	goto l876
																}
																if !_rules[ruleSrc16]() {
																	goto l876
																}
																{
																	add(ruleAction7, position)
																}
																add(ruleEx, position891)
															}
															break
														case 'P', 'p':
															{
																position897 := position
																{
																	position898, tokenIndex898 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l899
																	}
																	position++
																	goto l898
																l899:
																	position, tokenIndex = position898, tokenIndex898
																	if buffer[position] != rune('P') {
																		goto l876
																	}
																	position++
																}
															l898:
																{
																	position900, tokenIndex900 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l901
																	}
																	position++
																	goto l900
																l901:
																	position, tokenIndex = position900, tokenIndex900
																	if buffer[position] != rune('O') {
																		goto l876
																	}
																	position++
																}
															l900:
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
																		goto l876
																	}
																	position++
																}
															l902:
																if !_rules[rulews]() {
																	goto l876
																}
																if !_rules[ruleDst16]() {
																	goto l876
																}
																{
																	add(ruleAction6, position)
																}
																add(rulePop, position897)
															}
															break
														default:
															{
																position905 := position
																{
																	position906, tokenIndex906 := position, tokenIndex
																	{
																		position908 := position
																		{
																			position909, tokenIndex909 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l910
																			}
																			position++
																			goto l909
																		l910:
																			position, tokenIndex = position909, tokenIndex909
																			if buffer[position] != rune('L') {
																				goto l907
																			}
																			position++
																		}
																	l909:
																		{
																			position911, tokenIndex911 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l912
																			}
																			position++
																			goto l911
																		l912:
																			position, tokenIndex = position911, tokenIndex911
																			if buffer[position] != rune('D') {
																				goto l907
																			}
																			position++
																		}
																	l911:
																		if !_rules[rulews]() {
																			goto l907
																		}
																		if !_rules[ruleDst16]() {
																			goto l907
																		}
																		if !_rules[rulesep]() {
																			goto l907
																		}
																		if !_rules[ruleSrc16]() {
																			goto l907
																		}
																		{
																			add(ruleAction4, position)
																		}
																		add(ruleLoad16, position908)
																	}
																	goto l906
																l907:
																	position, tokenIndex = position906, tokenIndex906
																	{
																		position914 := position
																		{
																			position915, tokenIndex915 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l916
																			}
																			position++
																			goto l915
																		l916:
																			position, tokenIndex = position915, tokenIndex915
																			if buffer[position] != rune('L') {
																				goto l876
																			}
																			position++
																		}
																	l915:
																		{
																			position917, tokenIndex917 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l918
																			}
																			position++
																			goto l917
																		l918:
																			position, tokenIndex = position917, tokenIndex917
																			if buffer[position] != rune('D') {
																				goto l876
																			}
																			position++
																		}
																	l917:
																		if !_rules[rulews]() {
																			goto l876
																		}
																		{
																			position919 := position
																			{
																				position920, tokenIndex920 := position, tokenIndex
																				if !_rules[ruleReg8]() {
																					goto l921
																				}
																				goto l920
																			l921:
																				position, tokenIndex = position920, tokenIndex920
																				if !_rules[ruleReg16Contents]() {
																					goto l922
																				}
																				goto l920
																			l922:
																				position, tokenIndex = position920, tokenIndex920
																				if !_rules[rulenn_contents]() {
																					goto l876
																				}
																			}
																		l920:
																			{
																				add(ruleAction17, position)
																			}
																			add(ruleDst8, position919)
																		}
																		if !_rules[rulesep]() {
																			goto l876
																		}
																		if !_rules[ruleSrc8]() {
																			goto l876
																		}
																		{
																			add(ruleAction3, position)
																		}
																		add(ruleLoad8, position914)
																	}
																}
															l906:
																add(ruleLoad, position905)
															}
															break
														}
													}

												}
											l878:
												add(ruleAssignment, position877)
											}
											goto l875
										l876:
											position, tokenIndex = position875, tokenIndex875
											{
												position926 := position
												{
													position927, tokenIndex927 := position, tokenIndex
													{
														position929 := position
														{
															position930, tokenIndex930 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l931
															}
															position++
															goto l930
														l931:
															position, tokenIndex = position930, tokenIndex930
															if buffer[position] != rune('I') {
																goto l928
															}
															position++
														}
													l930:
														{
															position932, tokenIndex932 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l933
															}
															position++
															goto l932
														l933:
															position, tokenIndex = position932, tokenIndex932
															if buffer[position] != rune('N') {
																goto l928
															}
															position++
														}
													l932:
														{
															position934, tokenIndex934 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l935
															}
															position++
															goto l934
														l935:
															position, tokenIndex = position934, tokenIndex934
															if buffer[position] != rune('C') {
																goto l928
															}
															position++
														}
													l934:
														if !_rules[rulews]() {
															goto l928
														}
														if !_rules[ruleILoc8]() {
															goto l928
														}
														{
															add(ruleAction8, position)
														}
														add(ruleInc16Indexed8, position929)
													}
													goto l927
												l928:
													position, tokenIndex = position927, tokenIndex927
													{
														position938 := position
														{
															position939, tokenIndex939 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l940
															}
															position++
															goto l939
														l940:
															position, tokenIndex = position939, tokenIndex939
															if buffer[position] != rune('I') {
																goto l937
															}
															position++
														}
													l939:
														{
															position941, tokenIndex941 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l942
															}
															position++
															goto l941
														l942:
															position, tokenIndex = position941, tokenIndex941
															if buffer[position] != rune('N') {
																goto l937
															}
															position++
														}
													l941:
														{
															position943, tokenIndex943 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l944
															}
															position++
															goto l943
														l944:
															position, tokenIndex = position943, tokenIndex943
															if buffer[position] != rune('C') {
																goto l937
															}
															position++
														}
													l943:
														if !_rules[rulews]() {
															goto l937
														}
														if !_rules[ruleLoc16]() {
															goto l937
														}
														{
															add(ruleAction10, position)
														}
														add(ruleInc16, position938)
													}
													goto l927
												l937:
													position, tokenIndex = position927, tokenIndex927
													{
														position946 := position
														{
															position947, tokenIndex947 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l948
															}
															position++
															goto l947
														l948:
															position, tokenIndex = position947, tokenIndex947
															if buffer[position] != rune('I') {
																goto l925
															}
															position++
														}
													l947:
														{
															position949, tokenIndex949 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l950
															}
															position++
															goto l949
														l950:
															position, tokenIndex = position949, tokenIndex949
															if buffer[position] != rune('N') {
																goto l925
															}
															position++
														}
													l949:
														{
															position951, tokenIndex951 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l952
															}
															position++
															goto l951
														l952:
															position, tokenIndex = position951, tokenIndex951
															if buffer[position] != rune('C') {
																goto l925
															}
															position++
														}
													l951:
														if !_rules[rulews]() {
															goto l925
														}
														if !_rules[ruleLoc8]() {
															goto l925
														}
														{
															add(ruleAction9, position)
														}
														add(ruleInc8, position946)
													}
												}
											l927:
												add(ruleInc, position926)
											}
											goto l875
										l925:
											position, tokenIndex = position875, tokenIndex875
											{
												position955 := position
												{
													position956, tokenIndex956 := position, tokenIndex
													{
														position958 := position
														{
															position959, tokenIndex959 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l960
															}
															position++
															goto l959
														l960:
															position, tokenIndex = position959, tokenIndex959
															if buffer[position] != rune('D') {
																goto l957
															}
															position++
														}
													l959:
														{
															position961, tokenIndex961 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l962
															}
															position++
															goto l961
														l962:
															position, tokenIndex = position961, tokenIndex961
															if buffer[position] != rune('E') {
																goto l957
															}
															position++
														}
													l961:
														{
															position963, tokenIndex963 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l964
															}
															position++
															goto l963
														l964:
															position, tokenIndex = position963, tokenIndex963
															if buffer[position] != rune('C') {
																goto l957
															}
															position++
														}
													l963:
														if !_rules[rulews]() {
															goto l957
														}
														if !_rules[ruleILoc8]() {
															goto l957
														}
														{
															add(ruleAction11, position)
														}
														add(ruleDec16Indexed8, position958)
													}
													goto l956
												l957:
													position, tokenIndex = position956, tokenIndex956
													{
														position967 := position
														{
															position968, tokenIndex968 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l969
															}
															position++
															goto l968
														l969:
															position, tokenIndex = position968, tokenIndex968
															if buffer[position] != rune('D') {
																goto l966
															}
															position++
														}
													l968:
														{
															position970, tokenIndex970 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l971
															}
															position++
															goto l970
														l971:
															position, tokenIndex = position970, tokenIndex970
															if buffer[position] != rune('E') {
																goto l966
															}
															position++
														}
													l970:
														{
															position972, tokenIndex972 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l973
															}
															position++
															goto l972
														l973:
															position, tokenIndex = position972, tokenIndex972
															if buffer[position] != rune('C') {
																goto l966
															}
															position++
														}
													l972:
														if !_rules[rulews]() {
															goto l966
														}
														if !_rules[ruleLoc16]() {
															goto l966
														}
														{
															add(ruleAction13, position)
														}
														add(ruleDec16, position967)
													}
													goto l956
												l966:
													position, tokenIndex = position956, tokenIndex956
													{
														position975 := position
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
																goto l954
															}
															position++
														}
													l976:
														{
															position978, tokenIndex978 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l979
															}
															position++
															goto l978
														l979:
															position, tokenIndex = position978, tokenIndex978
															if buffer[position] != rune('E') {
																goto l954
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
																goto l954
															}
															position++
														}
													l980:
														if !_rules[rulews]() {
															goto l954
														}
														if !_rules[ruleLoc8]() {
															goto l954
														}
														{
															add(ruleAction12, position)
														}
														add(ruleDec8, position975)
													}
												}
											l956:
												add(ruleDec, position955)
											}
											goto l875
										l954:
											position, tokenIndex = position875, tokenIndex875
											{
												position984 := position
												{
													position985, tokenIndex985 := position, tokenIndex
													{
														position987 := position
														{
															position988, tokenIndex988 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l989
															}
															position++
															goto l988
														l989:
															position, tokenIndex = position988, tokenIndex988
															if buffer[position] != rune('A') {
																goto l986
															}
															position++
														}
													l988:
														{
															position990, tokenIndex990 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l991
															}
															position++
															goto l990
														l991:
															position, tokenIndex = position990, tokenIndex990
															if buffer[position] != rune('D') {
																goto l986
															}
															position++
														}
													l990:
														{
															position992, tokenIndex992 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l993
															}
															position++
															goto l992
														l993:
															position, tokenIndex = position992, tokenIndex992
															if buffer[position] != rune('D') {
																goto l986
															}
															position++
														}
													l992:
														if !_rules[rulews]() {
															goto l986
														}
														if !_rules[ruleDst16]() {
															goto l986
														}
														if !_rules[rulesep]() {
															goto l986
														}
														if !_rules[ruleSrc16]() {
															goto l986
														}
														{
															add(ruleAction14, position)
														}
														add(ruleAdd16, position987)
													}
													goto l985
												l986:
													position, tokenIndex = position985, tokenIndex985
													{
														position996 := position
														{
															position997, tokenIndex997 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l998
															}
															position++
															goto l997
														l998:
															position, tokenIndex = position997, tokenIndex997
															if buffer[position] != rune('A') {
																goto l995
															}
															position++
														}
													l997:
														{
															position999, tokenIndex999 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1000
															}
															position++
															goto l999
														l1000:
															position, tokenIndex = position999, tokenIndex999
															if buffer[position] != rune('D') {
																goto l995
															}
															position++
														}
													l999:
														{
															position1001, tokenIndex1001 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1002
															}
															position++
															goto l1001
														l1002:
															position, tokenIndex = position1001, tokenIndex1001
															if buffer[position] != rune('C') {
																goto l995
															}
															position++
														}
													l1001:
														if !_rules[rulews]() {
															goto l995
														}
														if !_rules[ruleDst16]() {
															goto l995
														}
														if !_rules[rulesep]() {
															goto l995
														}
														if !_rules[ruleSrc16]() {
															goto l995
														}
														{
															add(ruleAction15, position)
														}
														add(ruleAdc16, position996)
													}
													goto l985
												l995:
													position, tokenIndex = position985, tokenIndex985
													{
														position1004 := position
														{
															position1005, tokenIndex1005 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1006
															}
															position++
															goto l1005
														l1006:
															position, tokenIndex = position1005, tokenIndex1005
															if buffer[position] != rune('S') {
																goto l983
															}
															position++
														}
													l1005:
														{
															position1007, tokenIndex1007 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1008
															}
															position++
															goto l1007
														l1008:
															position, tokenIndex = position1007, tokenIndex1007
															if buffer[position] != rune('B') {
																goto l983
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
																goto l983
															}
															position++
														}
													l1009:
														if !_rules[rulews]() {
															goto l983
														}
														if !_rules[ruleDst16]() {
															goto l983
														}
														if !_rules[rulesep]() {
															goto l983
														}
														if !_rules[ruleSrc16]() {
															goto l983
														}
														{
															add(ruleAction16, position)
														}
														add(ruleSbc16, position1004)
													}
												}
											l985:
												add(ruleAlu16, position984)
											}
											goto l875
										l983:
											position, tokenIndex = position875, tokenIndex875
											{
												position1013 := position
												{
													position1014, tokenIndex1014 := position, tokenIndex
													{
														position1016 := position
														{
															position1017, tokenIndex1017 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1018
															}
															position++
															goto l1017
														l1018:
															position, tokenIndex = position1017, tokenIndex1017
															if buffer[position] != rune('A') {
																goto l1015
															}
															position++
														}
													l1017:
														{
															position1019, tokenIndex1019 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1020
															}
															position++
															goto l1019
														l1020:
															position, tokenIndex = position1019, tokenIndex1019
															if buffer[position] != rune('D') {
																goto l1015
															}
															position++
														}
													l1019:
														{
															position1021, tokenIndex1021 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1022
															}
															position++
															goto l1021
														l1022:
															position, tokenIndex = position1021, tokenIndex1021
															if buffer[position] != rune('D') {
																goto l1015
															}
															position++
														}
													l1021:
														if !_rules[rulews]() {
															goto l1015
														}
														{
															position1023, tokenIndex1023 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1024
															}
															position++
															goto l1023
														l1024:
															position, tokenIndex = position1023, tokenIndex1023
															if buffer[position] != rune('A') {
																goto l1015
															}
															position++
														}
													l1023:
														if !_rules[rulesep]() {
															goto l1015
														}
														if !_rules[ruleSrc8]() {
															goto l1015
														}
														{
															add(ruleAction41, position)
														}
														add(ruleAdd, position1016)
													}
													goto l1014
												l1015:
													position, tokenIndex = position1014, tokenIndex1014
													{
														position1027 := position
														{
															position1028, tokenIndex1028 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1029
															}
															position++
															goto l1028
														l1029:
															position, tokenIndex = position1028, tokenIndex1028
															if buffer[position] != rune('A') {
																goto l1026
															}
															position++
														}
													l1028:
														{
															position1030, tokenIndex1030 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1031
															}
															position++
															goto l1030
														l1031:
															position, tokenIndex = position1030, tokenIndex1030
															if buffer[position] != rune('D') {
																goto l1026
															}
															position++
														}
													l1030:
														{
															position1032, tokenIndex1032 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1033
															}
															position++
															goto l1032
														l1033:
															position, tokenIndex = position1032, tokenIndex1032
															if buffer[position] != rune('C') {
																goto l1026
															}
															position++
														}
													l1032:
														if !_rules[rulews]() {
															goto l1026
														}
														{
															position1034, tokenIndex1034 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1035
															}
															position++
															goto l1034
														l1035:
															position, tokenIndex = position1034, tokenIndex1034
															if buffer[position] != rune('A') {
																goto l1026
															}
															position++
														}
													l1034:
														if !_rules[rulesep]() {
															goto l1026
														}
														if !_rules[ruleSrc8]() {
															goto l1026
														}
														{
															add(ruleAction42, position)
														}
														add(ruleAdc, position1027)
													}
													goto l1014
												l1026:
													position, tokenIndex = position1014, tokenIndex1014
													{
														position1038 := position
														{
															position1039, tokenIndex1039 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1040
															}
															position++
															goto l1039
														l1040:
															position, tokenIndex = position1039, tokenIndex1039
															if buffer[position] != rune('S') {
																goto l1037
															}
															position++
														}
													l1039:
														{
															position1041, tokenIndex1041 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1042
															}
															position++
															goto l1041
														l1042:
															position, tokenIndex = position1041, tokenIndex1041
															if buffer[position] != rune('U') {
																goto l1037
															}
															position++
														}
													l1041:
														{
															position1043, tokenIndex1043 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1044
															}
															position++
															goto l1043
														l1044:
															position, tokenIndex = position1043, tokenIndex1043
															if buffer[position] != rune('B') {
																goto l1037
															}
															position++
														}
													l1043:
														if !_rules[rulews]() {
															goto l1037
														}
														if !_rules[ruleSrc8]() {
															goto l1037
														}
														{
															add(ruleAction43, position)
														}
														add(ruleSub, position1038)
													}
													goto l1014
												l1037:
													position, tokenIndex = position1014, tokenIndex1014
													{
														switch buffer[position] {
														case 'C', 'c':
															{
																position1047 := position
																{
																	position1048, tokenIndex1048 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1049
																	}
																	position++
																	goto l1048
																l1049:
																	position, tokenIndex = position1048, tokenIndex1048
																	if buffer[position] != rune('C') {
																		goto l1012
																	}
																	position++
																}
															l1048:
																{
																	position1050, tokenIndex1050 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1051
																	}
																	position++
																	goto l1050
																l1051:
																	position, tokenIndex = position1050, tokenIndex1050
																	if buffer[position] != rune('P') {
																		goto l1012
																	}
																	position++
																}
															l1050:
																if !_rules[rulews]() {
																	goto l1012
																}
																if !_rules[ruleSrc8]() {
																	goto l1012
																}
																{
																	add(ruleAction48, position)
																}
																add(ruleCp, position1047)
															}
															break
														case 'O', 'o':
															{
																position1053 := position
																{
																	position1054, tokenIndex1054 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1055
																	}
																	position++
																	goto l1054
																l1055:
																	position, tokenIndex = position1054, tokenIndex1054
																	if buffer[position] != rune('O') {
																		goto l1012
																	}
																	position++
																}
															l1054:
																{
																	position1056, tokenIndex1056 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1057
																	}
																	position++
																	goto l1056
																l1057:
																	position, tokenIndex = position1056, tokenIndex1056
																	if buffer[position] != rune('R') {
																		goto l1012
																	}
																	position++
																}
															l1056:
																if !_rules[rulews]() {
																	goto l1012
																}
																if !_rules[ruleSrc8]() {
																	goto l1012
																}
																{
																	add(ruleAction47, position)
																}
																add(ruleOr, position1053)
															}
															break
														case 'X', 'x':
															{
																position1059 := position
																{
																	position1060, tokenIndex1060 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1061
																	}
																	position++
																	goto l1060
																l1061:
																	position, tokenIndex = position1060, tokenIndex1060
																	if buffer[position] != rune('X') {
																		goto l1012
																	}
																	position++
																}
															l1060:
																{
																	position1062, tokenIndex1062 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1063
																	}
																	position++
																	goto l1062
																l1063:
																	position, tokenIndex = position1062, tokenIndex1062
																	if buffer[position] != rune('O') {
																		goto l1012
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
																		goto l1012
																	}
																	position++
																}
															l1064:
																if !_rules[rulews]() {
																	goto l1012
																}
																if !_rules[ruleSrc8]() {
																	goto l1012
																}
																{
																	add(ruleAction46, position)
																}
																add(ruleXor, position1059)
															}
															break
														case 'A', 'a':
															{
																position1067 := position
																{
																	position1068, tokenIndex1068 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1069
																	}
																	position++
																	goto l1068
																l1069:
																	position, tokenIndex = position1068, tokenIndex1068
																	if buffer[position] != rune('A') {
																		goto l1012
																	}
																	position++
																}
															l1068:
																{
																	position1070, tokenIndex1070 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1071
																	}
																	position++
																	goto l1070
																l1071:
																	position, tokenIndex = position1070, tokenIndex1070
																	if buffer[position] != rune('N') {
																		goto l1012
																	}
																	position++
																}
															l1070:
																{
																	position1072, tokenIndex1072 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1073
																	}
																	position++
																	goto l1072
																l1073:
																	position, tokenIndex = position1072, tokenIndex1072
																	if buffer[position] != rune('D') {
																		goto l1012
																	}
																	position++
																}
															l1072:
																if !_rules[rulews]() {
																	goto l1012
																}
																if !_rules[ruleSrc8]() {
																	goto l1012
																}
																{
																	add(ruleAction45, position)
																}
																add(ruleAnd, position1067)
															}
															break
														default:
															{
																position1075 := position
																{
																	position1076, tokenIndex1076 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1077
																	}
																	position++
																	goto l1076
																l1077:
																	position, tokenIndex = position1076, tokenIndex1076
																	if buffer[position] != rune('S') {
																		goto l1012
																	}
																	position++
																}
															l1076:
																{
																	position1078, tokenIndex1078 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1079
																	}
																	position++
																	goto l1078
																l1079:
																	position, tokenIndex = position1078, tokenIndex1078
																	if buffer[position] != rune('B') {
																		goto l1012
																	}
																	position++
																}
															l1078:
																{
																	position1080, tokenIndex1080 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1081
																	}
																	position++
																	goto l1080
																l1081:
																	position, tokenIndex = position1080, tokenIndex1080
																	if buffer[position] != rune('C') {
																		goto l1012
																	}
																	position++
																}
															l1080:
																if !_rules[rulews]() {
																	goto l1012
																}
																{
																	position1082, tokenIndex1082 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1083
																	}
																	position++
																	goto l1082
																l1083:
																	position, tokenIndex = position1082, tokenIndex1082
																	if buffer[position] != rune('A') {
																		goto l1012
																	}
																	position++
																}
															l1082:
																if !_rules[rulesep]() {
																	goto l1012
																}
																if !_rules[ruleSrc8]() {
																	goto l1012
																}
																{
																	add(ruleAction44, position)
																}
																add(ruleSbc, position1075)
															}
															break
														}
													}

												}
											l1014:
												add(ruleAlu, position1013)
											}
											goto l875
										l1012:
											position, tokenIndex = position875, tokenIndex875
											{
												position1086 := position
												{
													position1087, tokenIndex1087 := position, tokenIndex
													{
														position1089 := position
														{
															position1090, tokenIndex1090 := position, tokenIndex
															{
																position1092 := position
																{
																	position1093, tokenIndex1093 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1094
																	}
																	position++
																	goto l1093
																l1094:
																	position, tokenIndex = position1093, tokenIndex1093
																	if buffer[position] != rune('R') {
																		goto l1091
																	}
																	position++
																}
															l1093:
																{
																	position1095, tokenIndex1095 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1096
																	}
																	position++
																	goto l1095
																l1096:
																	position, tokenIndex = position1095, tokenIndex1095
																	if buffer[position] != rune('L') {
																		goto l1091
																	}
																	position++
																}
															l1095:
																{
																	position1097, tokenIndex1097 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1098
																	}
																	position++
																	goto l1097
																l1098:
																	position, tokenIndex = position1097, tokenIndex1097
																	if buffer[position] != rune('C') {
																		goto l1091
																	}
																	position++
																}
															l1097:
																if !_rules[rulews]() {
																	goto l1091
																}
																if !_rules[ruleLoc8]() {
																	goto l1091
																}
																{
																	position1099, tokenIndex1099 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1099
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1099
																	}
																	goto l1100
																l1099:
																	position, tokenIndex = position1099, tokenIndex1099
																}
															l1100:
																{
																	add(ruleAction49, position)
																}
																add(ruleRlc, position1092)
															}
															goto l1090
														l1091:
															position, tokenIndex = position1090, tokenIndex1090
															{
																position1103 := position
																{
																	position1104, tokenIndex1104 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1105
																	}
																	position++
																	goto l1104
																l1105:
																	position, tokenIndex = position1104, tokenIndex1104
																	if buffer[position] != rune('R') {
																		goto l1102
																	}
																	position++
																}
															l1104:
																{
																	position1106, tokenIndex1106 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1107
																	}
																	position++
																	goto l1106
																l1107:
																	position, tokenIndex = position1106, tokenIndex1106
																	if buffer[position] != rune('R') {
																		goto l1102
																	}
																	position++
																}
															l1106:
																{
																	position1108, tokenIndex1108 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1109
																	}
																	position++
																	goto l1108
																l1109:
																	position, tokenIndex = position1108, tokenIndex1108
																	if buffer[position] != rune('C') {
																		goto l1102
																	}
																	position++
																}
															l1108:
																if !_rules[rulews]() {
																	goto l1102
																}
																if !_rules[ruleLoc8]() {
																	goto l1102
																}
																{
																	position1110, tokenIndex1110 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1110
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1110
																	}
																	goto l1111
																l1110:
																	position, tokenIndex = position1110, tokenIndex1110
																}
															l1111:
																{
																	add(ruleAction50, position)
																}
																add(ruleRrc, position1103)
															}
															goto l1090
														l1102:
															position, tokenIndex = position1090, tokenIndex1090
															{
																position1114 := position
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
																		goto l1113
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
																		goto l1113
																	}
																	position++
																}
															l1117:
																if !_rules[rulews]() {
																	goto l1113
																}
																if !_rules[ruleLoc8]() {
																	goto l1113
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
																	add(ruleAction51, position)
																}
																add(ruleRl, position1114)
															}
															goto l1090
														l1113:
															position, tokenIndex = position1090, tokenIndex1090
															{
																position1123 := position
																{
																	position1124, tokenIndex1124 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1125
																	}
																	position++
																	goto l1124
																l1125:
																	position, tokenIndex = position1124, tokenIndex1124
																	if buffer[position] != rune('R') {
																		goto l1122
																	}
																	position++
																}
															l1124:
																{
																	position1126, tokenIndex1126 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1127
																	}
																	position++
																	goto l1126
																l1127:
																	position, tokenIndex = position1126, tokenIndex1126
																	if buffer[position] != rune('R') {
																		goto l1122
																	}
																	position++
																}
															l1126:
																if !_rules[rulews]() {
																	goto l1122
																}
																if !_rules[ruleLoc8]() {
																	goto l1122
																}
																{
																	position1128, tokenIndex1128 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1128
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1128
																	}
																	goto l1129
																l1128:
																	position, tokenIndex = position1128, tokenIndex1128
																}
															l1129:
																{
																	add(ruleAction52, position)
																}
																add(ruleRr, position1123)
															}
															goto l1090
														l1122:
															position, tokenIndex = position1090, tokenIndex1090
															{
																position1132 := position
																{
																	position1133, tokenIndex1133 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1134
																	}
																	position++
																	goto l1133
																l1134:
																	position, tokenIndex = position1133, tokenIndex1133
																	if buffer[position] != rune('S') {
																		goto l1131
																	}
																	position++
																}
															l1133:
																{
																	position1135, tokenIndex1135 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1136
																	}
																	position++
																	goto l1135
																l1136:
																	position, tokenIndex = position1135, tokenIndex1135
																	if buffer[position] != rune('L') {
																		goto l1131
																	}
																	position++
																}
															l1135:
																{
																	position1137, tokenIndex1137 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1138
																	}
																	position++
																	goto l1137
																l1138:
																	position, tokenIndex = position1137, tokenIndex1137
																	if buffer[position] != rune('A') {
																		goto l1131
																	}
																	position++
																}
															l1137:
																if !_rules[rulews]() {
																	goto l1131
																}
																if !_rules[ruleLoc8]() {
																	goto l1131
																}
																{
																	position1139, tokenIndex1139 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1139
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1139
																	}
																	goto l1140
																l1139:
																	position, tokenIndex = position1139, tokenIndex1139
																}
															l1140:
																{
																	add(ruleAction53, position)
																}
																add(ruleSla, position1132)
															}
															goto l1090
														l1131:
															position, tokenIndex = position1090, tokenIndex1090
															{
																position1143 := position
																{
																	position1144, tokenIndex1144 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1145
																	}
																	position++
																	goto l1144
																l1145:
																	position, tokenIndex = position1144, tokenIndex1144
																	if buffer[position] != rune('S') {
																		goto l1142
																	}
																	position++
																}
															l1144:
																{
																	position1146, tokenIndex1146 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1147
																	}
																	position++
																	goto l1146
																l1147:
																	position, tokenIndex = position1146, tokenIndex1146
																	if buffer[position] != rune('R') {
																		goto l1142
																	}
																	position++
																}
															l1146:
																{
																	position1148, tokenIndex1148 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1149
																	}
																	position++
																	goto l1148
																l1149:
																	position, tokenIndex = position1148, tokenIndex1148
																	if buffer[position] != rune('A') {
																		goto l1142
																	}
																	position++
																}
															l1148:
																if !_rules[rulews]() {
																	goto l1142
																}
																if !_rules[ruleLoc8]() {
																	goto l1142
																}
																{
																	position1150, tokenIndex1150 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1150
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1150
																	}
																	goto l1151
																l1150:
																	position, tokenIndex = position1150, tokenIndex1150
																}
															l1151:
																{
																	add(ruleAction54, position)
																}
																add(ruleSra, position1143)
															}
															goto l1090
														l1142:
															position, tokenIndex = position1090, tokenIndex1090
															{
																position1154 := position
																{
																	position1155, tokenIndex1155 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1156
																	}
																	position++
																	goto l1155
																l1156:
																	position, tokenIndex = position1155, tokenIndex1155
																	if buffer[position] != rune('S') {
																		goto l1153
																	}
																	position++
																}
															l1155:
																{
																	position1157, tokenIndex1157 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1158
																	}
																	position++
																	goto l1157
																l1158:
																	position, tokenIndex = position1157, tokenIndex1157
																	if buffer[position] != rune('L') {
																		goto l1153
																	}
																	position++
																}
															l1157:
																{
																	position1159, tokenIndex1159 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1160
																	}
																	position++
																	goto l1159
																l1160:
																	position, tokenIndex = position1159, tokenIndex1159
																	if buffer[position] != rune('L') {
																		goto l1153
																	}
																	position++
																}
															l1159:
																if !_rules[rulews]() {
																	goto l1153
																}
																if !_rules[ruleLoc8]() {
																	goto l1153
																}
																{
																	position1161, tokenIndex1161 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1161
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1161
																	}
																	goto l1162
																l1161:
																	position, tokenIndex = position1161, tokenIndex1161
																}
															l1162:
																{
																	add(ruleAction55, position)
																}
																add(ruleSll, position1154)
															}
															goto l1090
														l1153:
															position, tokenIndex = position1090, tokenIndex1090
															{
																position1164 := position
																{
																	position1165, tokenIndex1165 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1166
																	}
																	position++
																	goto l1165
																l1166:
																	position, tokenIndex = position1165, tokenIndex1165
																	if buffer[position] != rune('S') {
																		goto l1088
																	}
																	position++
																}
															l1165:
																{
																	position1167, tokenIndex1167 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1168
																	}
																	position++
																	goto l1167
																l1168:
																	position, tokenIndex = position1167, tokenIndex1167
																	if buffer[position] != rune('R') {
																		goto l1088
																	}
																	position++
																}
															l1167:
																{
																	position1169, tokenIndex1169 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1170
																	}
																	position++
																	goto l1169
																l1170:
																	position, tokenIndex = position1169, tokenIndex1169
																	if buffer[position] != rune('L') {
																		goto l1088
																	}
																	position++
																}
															l1169:
																if !_rules[rulews]() {
																	goto l1088
																}
																if !_rules[ruleLoc8]() {
																	goto l1088
																}
																{
																	position1171, tokenIndex1171 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1171
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1171
																	}
																	goto l1172
																l1171:
																	position, tokenIndex = position1171, tokenIndex1171
																}
															l1172:
																{
																	add(ruleAction56, position)
																}
																add(ruleSrl, position1164)
															}
														}
													l1090:
														add(ruleRot, position1089)
													}
													goto l1087
												l1088:
													position, tokenIndex = position1087, tokenIndex1087
													{
														switch buffer[position] {
														case 'S', 's':
															{
																position1175 := position
																{
																	position1176, tokenIndex1176 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1177
																	}
																	position++
																	goto l1176
																l1177:
																	position, tokenIndex = position1176, tokenIndex1176
																	if buffer[position] != rune('S') {
																		goto l1085
																	}
																	position++
																}
															l1176:
																{
																	position1178, tokenIndex1178 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1179
																	}
																	position++
																	goto l1178
																l1179:
																	position, tokenIndex = position1178, tokenIndex1178
																	if buffer[position] != rune('E') {
																		goto l1085
																	}
																	position++
																}
															l1178:
																{
																	position1180, tokenIndex1180 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1181
																	}
																	position++
																	goto l1180
																l1181:
																	position, tokenIndex = position1180, tokenIndex1180
																	if buffer[position] != rune('T') {
																		goto l1085
																	}
																	position++
																}
															l1180:
																if !_rules[rulews]() {
																	goto l1085
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1085
																}
																if !_rules[rulesep]() {
																	goto l1085
																}
																if !_rules[ruleLoc8]() {
																	goto l1085
																}
																{
																	position1182, tokenIndex1182 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1182
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1182
																	}
																	goto l1183
																l1182:
																	position, tokenIndex = position1182, tokenIndex1182
																}
															l1183:
																{
																	add(ruleAction59, position)
																}
																add(ruleSet, position1175)
															}
															break
														case 'R', 'r':
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
																		goto l1085
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
																		goto l1085
																	}
																	position++
																}
															l1188:
																{
																	position1190, tokenIndex1190 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1191
																	}
																	position++
																	goto l1190
																l1191:
																	position, tokenIndex = position1190, tokenIndex1190
																	if buffer[position] != rune('S') {
																		goto l1085
																	}
																	position++
																}
															l1190:
																if !_rules[rulews]() {
																	goto l1085
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1085
																}
																if !_rules[rulesep]() {
																	goto l1085
																}
																if !_rules[ruleLoc8]() {
																	goto l1085
																}
																{
																	position1192, tokenIndex1192 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1192
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1192
																	}
																	goto l1193
																l1192:
																	position, tokenIndex = position1192, tokenIndex1192
																}
															l1193:
																{
																	add(ruleAction58, position)
																}
																add(ruleRes, position1185)
															}
															break
														default:
															{
																position1195 := position
																{
																	position1196, tokenIndex1196 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1197
																	}
																	position++
																	goto l1196
																l1197:
																	position, tokenIndex = position1196, tokenIndex1196
																	if buffer[position] != rune('B') {
																		goto l1085
																	}
																	position++
																}
															l1196:
																{
																	position1198, tokenIndex1198 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l1199
																	}
																	position++
																	goto l1198
																l1199:
																	position, tokenIndex = position1198, tokenIndex1198
																	if buffer[position] != rune('I') {
																		goto l1085
																	}
																	position++
																}
															l1198:
																{
																	position1200, tokenIndex1200 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1201
																	}
																	position++
																	goto l1200
																l1201:
																	position, tokenIndex = position1200, tokenIndex1200
																	if buffer[position] != rune('T') {
																		goto l1085
																	}
																	position++
																}
															l1200:
																if !_rules[rulews]() {
																	goto l1085
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1085
																}
																if !_rules[rulesep]() {
																	goto l1085
																}
																if !_rules[ruleLoc8]() {
																	goto l1085
																}
																{
																	add(ruleAction57, position)
																}
																add(ruleBit, position1195)
															}
															break
														}
													}

												}
											l1087:
												add(ruleBitOp, position1086)
											}
											goto l875
										l1085:
											position, tokenIndex = position875, tokenIndex875
											{
												position1204 := position
												{
													position1205, tokenIndex1205 := position, tokenIndex
													{
														position1207 := position
														{
															position1208 := position
															{
																position1209, tokenIndex1209 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1210
																}
																position++
																goto l1209
															l1210:
																position, tokenIndex = position1209, tokenIndex1209
																if buffer[position] != rune('R') {
																	goto l1206
																}
																position++
															}
														l1209:
															{
																position1211, tokenIndex1211 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1212
																}
																position++
																goto l1211
															l1212:
																position, tokenIndex = position1211, tokenIndex1211
																if buffer[position] != rune('E') {
																	goto l1206
																}
																position++
															}
														l1211:
															{
																position1213, tokenIndex1213 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1214
																}
																position++
																goto l1213
															l1214:
																position, tokenIndex = position1213, tokenIndex1213
																if buffer[position] != rune('T') {
																	goto l1206
																}
																position++
															}
														l1213:
															{
																position1215, tokenIndex1215 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1216
																}
																position++
																goto l1215
															l1216:
																position, tokenIndex = position1215, tokenIndex1215
																if buffer[position] != rune('N') {
																	goto l1206
																}
																position++
															}
														l1215:
															add(rulePegText, position1208)
														}
														{
															add(ruleAction74, position)
														}
														add(ruleRetn, position1207)
													}
													goto l1205
												l1206:
													position, tokenIndex = position1205, tokenIndex1205
													{
														position1219 := position
														{
															position1220 := position
															{
																position1221, tokenIndex1221 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1222
																}
																position++
																goto l1221
															l1222:
																position, tokenIndex = position1221, tokenIndex1221
																if buffer[position] != rune('R') {
																	goto l1218
																}
																position++
															}
														l1221:
															{
																position1223, tokenIndex1223 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1224
																}
																position++
																goto l1223
															l1224:
																position, tokenIndex = position1223, tokenIndex1223
																if buffer[position] != rune('E') {
																	goto l1218
																}
																position++
															}
														l1223:
															{
																position1225, tokenIndex1225 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1226
																}
																position++
																goto l1225
															l1226:
																position, tokenIndex = position1225, tokenIndex1225
																if buffer[position] != rune('T') {
																	goto l1218
																}
																position++
															}
														l1225:
															{
																position1227, tokenIndex1227 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1228
																}
																position++
																goto l1227
															l1228:
																position, tokenIndex = position1227, tokenIndex1227
																if buffer[position] != rune('I') {
																	goto l1218
																}
																position++
															}
														l1227:
															add(rulePegText, position1220)
														}
														{
															add(ruleAction75, position)
														}
														add(ruleReti, position1219)
													}
													goto l1205
												l1218:
													position, tokenIndex = position1205, tokenIndex1205
													{
														position1231 := position
														{
															position1232 := position
															{
																position1233, tokenIndex1233 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1234
																}
																position++
																goto l1233
															l1234:
																position, tokenIndex = position1233, tokenIndex1233
																if buffer[position] != rune('R') {
																	goto l1230
																}
																position++
															}
														l1233:
															{
																position1235, tokenIndex1235 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1236
																}
																position++
																goto l1235
															l1236:
																position, tokenIndex = position1235, tokenIndex1235
																if buffer[position] != rune('R') {
																	goto l1230
																}
																position++
															}
														l1235:
															{
																position1237, tokenIndex1237 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1238
																}
																position++
																goto l1237
															l1238:
																position, tokenIndex = position1237, tokenIndex1237
																if buffer[position] != rune('D') {
																	goto l1230
																}
																position++
															}
														l1237:
															add(rulePegText, position1232)
														}
														{
															add(ruleAction76, position)
														}
														add(ruleRrd, position1231)
													}
													goto l1205
												l1230:
													position, tokenIndex = position1205, tokenIndex1205
													{
														position1241 := position
														{
															position1242 := position
															{
																position1243, tokenIndex1243 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1244
																}
																position++
																goto l1243
															l1244:
																position, tokenIndex = position1243, tokenIndex1243
																if buffer[position] != rune('I') {
																	goto l1240
																}
																position++
															}
														l1243:
															{
																position1245, tokenIndex1245 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1246
																}
																position++
																goto l1245
															l1246:
																position, tokenIndex = position1245, tokenIndex1245
																if buffer[position] != rune('M') {
																	goto l1240
																}
																position++
															}
														l1245:
															if buffer[position] != rune(' ') {
																goto l1240
															}
															position++
															if buffer[position] != rune('0') {
																goto l1240
															}
															position++
															add(rulePegText, position1242)
														}
														{
															add(ruleAction78, position)
														}
														add(ruleIm0, position1241)
													}
													goto l1205
												l1240:
													position, tokenIndex = position1205, tokenIndex1205
													{
														position1249 := position
														{
															position1250 := position
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
																	goto l1248
																}
																position++
															}
														l1251:
															{
																position1253, tokenIndex1253 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1254
																}
																position++
																goto l1253
															l1254:
																position, tokenIndex = position1253, tokenIndex1253
																if buffer[position] != rune('M') {
																	goto l1248
																}
																position++
															}
														l1253:
															if buffer[position] != rune(' ') {
																goto l1248
															}
															position++
															if buffer[position] != rune('1') {
																goto l1248
															}
															position++
															add(rulePegText, position1250)
														}
														{
															add(ruleAction79, position)
														}
														add(ruleIm1, position1249)
													}
													goto l1205
												l1248:
													position, tokenIndex = position1205, tokenIndex1205
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
																if buffer[position] != rune('m') {
																	goto l1262
																}
																position++
																goto l1261
															l1262:
																position, tokenIndex = position1261, tokenIndex1261
																if buffer[position] != rune('M') {
																	goto l1256
																}
																position++
															}
														l1261:
															if buffer[position] != rune(' ') {
																goto l1256
															}
															position++
															if buffer[position] != rune('2') {
																goto l1256
															}
															position++
															add(rulePegText, position1258)
														}
														{
															add(ruleAction80, position)
														}
														add(ruleIm2, position1257)
													}
													goto l1205
												l1256:
													position, tokenIndex = position1205, tokenIndex1205
													{
														switch buffer[position] {
														case 'I', 'O', 'i', 'o':
															{
																position1265 := position
																{
																	position1266, tokenIndex1266 := position, tokenIndex
																	{
																		position1268 := position
																		{
																			position1269 := position
																			{
																				position1270, tokenIndex1270 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1271
																				}
																				position++
																				goto l1270
																			l1271:
																				position, tokenIndex = position1270, tokenIndex1270
																				if buffer[position] != rune('I') {
																					goto l1267
																				}
																				position++
																			}
																		l1270:
																			{
																				position1272, tokenIndex1272 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1273
																				}
																				position++
																				goto l1272
																			l1273:
																				position, tokenIndex = position1272, tokenIndex1272
																				if buffer[position] != rune('N') {
																					goto l1267
																				}
																				position++
																			}
																		l1272:
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
																					goto l1267
																				}
																				position++
																			}
																		l1274:
																			{
																				position1276, tokenIndex1276 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1277
																				}
																				position++
																				goto l1276
																			l1277:
																				position, tokenIndex = position1276, tokenIndex1276
																				if buffer[position] != rune('R') {
																					goto l1267
																				}
																				position++
																			}
																		l1276:
																			add(rulePegText, position1269)
																		}
																		{
																			add(ruleAction91, position)
																		}
																		add(ruleInir, position1268)
																	}
																	goto l1266
																l1267:
																	position, tokenIndex = position1266, tokenIndex1266
																	{
																		position1280 := position
																		{
																			position1281 := position
																			{
																				position1282, tokenIndex1282 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1283
																				}
																				position++
																				goto l1282
																			l1283:
																				position, tokenIndex = position1282, tokenIndex1282
																				if buffer[position] != rune('I') {
																					goto l1279
																				}
																				position++
																			}
																		l1282:
																			{
																				position1284, tokenIndex1284 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1285
																				}
																				position++
																				goto l1284
																			l1285:
																				position, tokenIndex = position1284, tokenIndex1284
																				if buffer[position] != rune('N') {
																					goto l1279
																				}
																				position++
																			}
																		l1284:
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
																					goto l1279
																				}
																				position++
																			}
																		l1286:
																			add(rulePegText, position1281)
																		}
																		{
																			add(ruleAction83, position)
																		}
																		add(ruleIni, position1280)
																	}
																	goto l1266
																l1279:
																	position, tokenIndex = position1266, tokenIndex1266
																	{
																		position1290 := position
																		{
																			position1291 := position
																			{
																				position1292, tokenIndex1292 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1293
																				}
																				position++
																				goto l1292
																			l1293:
																				position, tokenIndex = position1292, tokenIndex1292
																				if buffer[position] != rune('O') {
																					goto l1289
																				}
																				position++
																			}
																		l1292:
																			{
																				position1294, tokenIndex1294 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1295
																				}
																				position++
																				goto l1294
																			l1295:
																				position, tokenIndex = position1294, tokenIndex1294
																				if buffer[position] != rune('T') {
																					goto l1289
																				}
																				position++
																			}
																		l1294:
																			{
																				position1296, tokenIndex1296 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1297
																				}
																				position++
																				goto l1296
																			l1297:
																				position, tokenIndex = position1296, tokenIndex1296
																				if buffer[position] != rune('I') {
																					goto l1289
																				}
																				position++
																			}
																		l1296:
																			{
																				position1298, tokenIndex1298 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1299
																				}
																				position++
																				goto l1298
																			l1299:
																				position, tokenIndex = position1298, tokenIndex1298
																				if buffer[position] != rune('R') {
																					goto l1289
																				}
																				position++
																			}
																		l1298:
																			add(rulePegText, position1291)
																		}
																		{
																			add(ruleAction92, position)
																		}
																		add(ruleOtir, position1290)
																	}
																	goto l1266
																l1289:
																	position, tokenIndex = position1266, tokenIndex1266
																	{
																		position1302 := position
																		{
																			position1303 := position
																			{
																				position1304, tokenIndex1304 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1305
																				}
																				position++
																				goto l1304
																			l1305:
																				position, tokenIndex = position1304, tokenIndex1304
																				if buffer[position] != rune('O') {
																					goto l1301
																				}
																				position++
																			}
																		l1304:
																			{
																				position1306, tokenIndex1306 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1307
																				}
																				position++
																				goto l1306
																			l1307:
																				position, tokenIndex = position1306, tokenIndex1306
																				if buffer[position] != rune('U') {
																					goto l1301
																				}
																				position++
																			}
																		l1306:
																			{
																				position1308, tokenIndex1308 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1309
																				}
																				position++
																				goto l1308
																			l1309:
																				position, tokenIndex = position1308, tokenIndex1308
																				if buffer[position] != rune('T') {
																					goto l1301
																				}
																				position++
																			}
																		l1308:
																			{
																				position1310, tokenIndex1310 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1311
																				}
																				position++
																				goto l1310
																			l1311:
																				position, tokenIndex = position1310, tokenIndex1310
																				if buffer[position] != rune('I') {
																					goto l1301
																				}
																				position++
																			}
																		l1310:
																			add(rulePegText, position1303)
																		}
																		{
																			add(ruleAction84, position)
																		}
																		add(ruleOuti, position1302)
																	}
																	goto l1266
																l1301:
																	position, tokenIndex = position1266, tokenIndex1266
																	{
																		position1314 := position
																		{
																			position1315 := position
																			{
																				position1316, tokenIndex1316 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1317
																				}
																				position++
																				goto l1316
																			l1317:
																				position, tokenIndex = position1316, tokenIndex1316
																				if buffer[position] != rune('I') {
																					goto l1313
																				}
																				position++
																			}
																		l1316:
																			{
																				position1318, tokenIndex1318 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1319
																				}
																				position++
																				goto l1318
																			l1319:
																				position, tokenIndex = position1318, tokenIndex1318
																				if buffer[position] != rune('N') {
																					goto l1313
																				}
																				position++
																			}
																		l1318:
																			{
																				position1320, tokenIndex1320 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1321
																				}
																				position++
																				goto l1320
																			l1321:
																				position, tokenIndex = position1320, tokenIndex1320
																				if buffer[position] != rune('D') {
																					goto l1313
																				}
																				position++
																			}
																		l1320:
																			{
																				position1322, tokenIndex1322 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1323
																				}
																				position++
																				goto l1322
																			l1323:
																				position, tokenIndex = position1322, tokenIndex1322
																				if buffer[position] != rune('R') {
																					goto l1313
																				}
																				position++
																			}
																		l1322:
																			add(rulePegText, position1315)
																		}
																		{
																			add(ruleAction95, position)
																		}
																		add(ruleIndr, position1314)
																	}
																	goto l1266
																l1313:
																	position, tokenIndex = position1266, tokenIndex1266
																	{
																		position1326 := position
																		{
																			position1327 := position
																			{
																				position1328, tokenIndex1328 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1329
																				}
																				position++
																				goto l1328
																			l1329:
																				position, tokenIndex = position1328, tokenIndex1328
																				if buffer[position] != rune('I') {
																					goto l1325
																				}
																				position++
																			}
																		l1328:
																			{
																				position1330, tokenIndex1330 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1331
																				}
																				position++
																				goto l1330
																			l1331:
																				position, tokenIndex = position1330, tokenIndex1330
																				if buffer[position] != rune('N') {
																					goto l1325
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
																					goto l1325
																				}
																				position++
																			}
																		l1332:
																			add(rulePegText, position1327)
																		}
																		{
																			add(ruleAction87, position)
																		}
																		add(ruleInd, position1326)
																	}
																	goto l1266
																l1325:
																	position, tokenIndex = position1266, tokenIndex1266
																	{
																		position1336 := position
																		{
																			position1337 := position
																			{
																				position1338, tokenIndex1338 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1339
																				}
																				position++
																				goto l1338
																			l1339:
																				position, tokenIndex = position1338, tokenIndex1338
																				if buffer[position] != rune('O') {
																					goto l1335
																				}
																				position++
																			}
																		l1338:
																			{
																				position1340, tokenIndex1340 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1341
																				}
																				position++
																				goto l1340
																			l1341:
																				position, tokenIndex = position1340, tokenIndex1340
																				if buffer[position] != rune('T') {
																					goto l1335
																				}
																				position++
																			}
																		l1340:
																			{
																				position1342, tokenIndex1342 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1343
																				}
																				position++
																				goto l1342
																			l1343:
																				position, tokenIndex = position1342, tokenIndex1342
																				if buffer[position] != rune('D') {
																					goto l1335
																				}
																				position++
																			}
																		l1342:
																			{
																				position1344, tokenIndex1344 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1345
																				}
																				position++
																				goto l1344
																			l1345:
																				position, tokenIndex = position1344, tokenIndex1344
																				if buffer[position] != rune('R') {
																					goto l1335
																				}
																				position++
																			}
																		l1344:
																			add(rulePegText, position1337)
																		}
																		{
																			add(ruleAction96, position)
																		}
																		add(ruleOtdr, position1336)
																	}
																	goto l1266
																l1335:
																	position, tokenIndex = position1266, tokenIndex1266
																	{
																		position1347 := position
																		{
																			position1348 := position
																			{
																				position1349, tokenIndex1349 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1350
																				}
																				position++
																				goto l1349
																			l1350:
																				position, tokenIndex = position1349, tokenIndex1349
																				if buffer[position] != rune('O') {
																					goto l1203
																				}
																				position++
																			}
																		l1349:
																			{
																				position1351, tokenIndex1351 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1352
																				}
																				position++
																				goto l1351
																			l1352:
																				position, tokenIndex = position1351, tokenIndex1351
																				if buffer[position] != rune('U') {
																					goto l1203
																				}
																				position++
																			}
																		l1351:
																			{
																				position1353, tokenIndex1353 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1354
																				}
																				position++
																				goto l1353
																			l1354:
																				position, tokenIndex = position1353, tokenIndex1353
																				if buffer[position] != rune('T') {
																					goto l1203
																				}
																				position++
																			}
																		l1353:
																			{
																				position1355, tokenIndex1355 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1356
																				}
																				position++
																				goto l1355
																			l1356:
																				position, tokenIndex = position1355, tokenIndex1355
																				if buffer[position] != rune('D') {
																					goto l1203
																				}
																				position++
																			}
																		l1355:
																			add(rulePegText, position1348)
																		}
																		{
																			add(ruleAction88, position)
																		}
																		add(ruleOutd, position1347)
																	}
																}
															l1266:
																add(ruleBlitIO, position1265)
															}
															break
														case 'R', 'r':
															{
																position1358 := position
																{
																	position1359 := position
																	{
																		position1360, tokenIndex1360 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1361
																		}
																		position++
																		goto l1360
																	l1361:
																		position, tokenIndex = position1360, tokenIndex1360
																		if buffer[position] != rune('R') {
																			goto l1203
																		}
																		position++
																	}
																l1360:
																	{
																		position1362, tokenIndex1362 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1363
																		}
																		position++
																		goto l1362
																	l1363:
																		position, tokenIndex = position1362, tokenIndex1362
																		if buffer[position] != rune('L') {
																			goto l1203
																		}
																		position++
																	}
																l1362:
																	{
																		position1364, tokenIndex1364 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1365
																		}
																		position++
																		goto l1364
																	l1365:
																		position, tokenIndex = position1364, tokenIndex1364
																		if buffer[position] != rune('D') {
																			goto l1203
																		}
																		position++
																	}
																l1364:
																	add(rulePegText, position1359)
																}
																{
																	add(ruleAction77, position)
																}
																add(ruleRld, position1358)
															}
															break
														case 'N', 'n':
															{
																position1367 := position
																{
																	position1368 := position
																	{
																		position1369, tokenIndex1369 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1370
																		}
																		position++
																		goto l1369
																	l1370:
																		position, tokenIndex = position1369, tokenIndex1369
																		if buffer[position] != rune('N') {
																			goto l1203
																		}
																		position++
																	}
																l1369:
																	{
																		position1371, tokenIndex1371 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1372
																		}
																		position++
																		goto l1371
																	l1372:
																		position, tokenIndex = position1371, tokenIndex1371
																		if buffer[position] != rune('E') {
																			goto l1203
																		}
																		position++
																	}
																l1371:
																	{
																		position1373, tokenIndex1373 := position, tokenIndex
																		if buffer[position] != rune('g') {
																			goto l1374
																		}
																		position++
																		goto l1373
																	l1374:
																		position, tokenIndex = position1373, tokenIndex1373
																		if buffer[position] != rune('G') {
																			goto l1203
																		}
																		position++
																	}
																l1373:
																	add(rulePegText, position1368)
																}
																{
																	add(ruleAction73, position)
																}
																add(ruleNeg, position1367)
															}
															break
														default:
															{
																position1376 := position
																{
																	position1377, tokenIndex1377 := position, tokenIndex
																	{
																		position1379 := position
																		{
																			position1380 := position
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
																					goto l1378
																				}
																				position++
																			}
																		l1381:
																			{
																				position1383, tokenIndex1383 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1384
																				}
																				position++
																				goto l1383
																			l1384:
																				position, tokenIndex = position1383, tokenIndex1383
																				if buffer[position] != rune('D') {
																					goto l1378
																				}
																				position++
																			}
																		l1383:
																			{
																				position1385, tokenIndex1385 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1386
																				}
																				position++
																				goto l1385
																			l1386:
																				position, tokenIndex = position1385, tokenIndex1385
																				if buffer[position] != rune('I') {
																					goto l1378
																				}
																				position++
																			}
																		l1385:
																			{
																				position1387, tokenIndex1387 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1388
																				}
																				position++
																				goto l1387
																			l1388:
																				position, tokenIndex = position1387, tokenIndex1387
																				if buffer[position] != rune('R') {
																					goto l1378
																				}
																				position++
																			}
																		l1387:
																			add(rulePegText, position1380)
																		}
																		{
																			add(ruleAction89, position)
																		}
																		add(ruleLdir, position1379)
																	}
																	goto l1377
																l1378:
																	position, tokenIndex = position1377, tokenIndex1377
																	{
																		position1391 := position
																		{
																			position1392 := position
																			{
																				position1393, tokenIndex1393 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1394
																				}
																				position++
																				goto l1393
																			l1394:
																				position, tokenIndex = position1393, tokenIndex1393
																				if buffer[position] != rune('L') {
																					goto l1390
																				}
																				position++
																			}
																		l1393:
																			{
																				position1395, tokenIndex1395 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1396
																				}
																				position++
																				goto l1395
																			l1396:
																				position, tokenIndex = position1395, tokenIndex1395
																				if buffer[position] != rune('D') {
																					goto l1390
																				}
																				position++
																			}
																		l1395:
																			{
																				position1397, tokenIndex1397 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1398
																				}
																				position++
																				goto l1397
																			l1398:
																				position, tokenIndex = position1397, tokenIndex1397
																				if buffer[position] != rune('I') {
																					goto l1390
																				}
																				position++
																			}
																		l1397:
																			add(rulePegText, position1392)
																		}
																		{
																			add(ruleAction81, position)
																		}
																		add(ruleLdi, position1391)
																	}
																	goto l1377
																l1390:
																	position, tokenIndex = position1377, tokenIndex1377
																	{
																		position1401 := position
																		{
																			position1402 := position
																			{
																				position1403, tokenIndex1403 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1404
																				}
																				position++
																				goto l1403
																			l1404:
																				position, tokenIndex = position1403, tokenIndex1403
																				if buffer[position] != rune('C') {
																					goto l1400
																				}
																				position++
																			}
																		l1403:
																			{
																				position1405, tokenIndex1405 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1406
																				}
																				position++
																				goto l1405
																			l1406:
																				position, tokenIndex = position1405, tokenIndex1405
																				if buffer[position] != rune('P') {
																					goto l1400
																				}
																				position++
																			}
																		l1405:
																			{
																				position1407, tokenIndex1407 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1408
																				}
																				position++
																				goto l1407
																			l1408:
																				position, tokenIndex = position1407, tokenIndex1407
																				if buffer[position] != rune('I') {
																					goto l1400
																				}
																				position++
																			}
																		l1407:
																			{
																				position1409, tokenIndex1409 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1410
																				}
																				position++
																				goto l1409
																			l1410:
																				position, tokenIndex = position1409, tokenIndex1409
																				if buffer[position] != rune('R') {
																					goto l1400
																				}
																				position++
																			}
																		l1409:
																			add(rulePegText, position1402)
																		}
																		{
																			add(ruleAction90, position)
																		}
																		add(ruleCpir, position1401)
																	}
																	goto l1377
																l1400:
																	position, tokenIndex = position1377, tokenIndex1377
																	{
																		position1413 := position
																		{
																			position1414 := position
																			{
																				position1415, tokenIndex1415 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1416
																				}
																				position++
																				goto l1415
																			l1416:
																				position, tokenIndex = position1415, tokenIndex1415
																				if buffer[position] != rune('C') {
																					goto l1412
																				}
																				position++
																			}
																		l1415:
																			{
																				position1417, tokenIndex1417 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1418
																				}
																				position++
																				goto l1417
																			l1418:
																				position, tokenIndex = position1417, tokenIndex1417
																				if buffer[position] != rune('P') {
																					goto l1412
																				}
																				position++
																			}
																		l1417:
																			{
																				position1419, tokenIndex1419 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1420
																				}
																				position++
																				goto l1419
																			l1420:
																				position, tokenIndex = position1419, tokenIndex1419
																				if buffer[position] != rune('I') {
																					goto l1412
																				}
																				position++
																			}
																		l1419:
																			add(rulePegText, position1414)
																		}
																		{
																			add(ruleAction82, position)
																		}
																		add(ruleCpi, position1413)
																	}
																	goto l1377
																l1412:
																	position, tokenIndex = position1377, tokenIndex1377
																	{
																		position1423 := position
																		{
																			position1424 := position
																			{
																				position1425, tokenIndex1425 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1426
																				}
																				position++
																				goto l1425
																			l1426:
																				position, tokenIndex = position1425, tokenIndex1425
																				if buffer[position] != rune('L') {
																					goto l1422
																				}
																				position++
																			}
																		l1425:
																			{
																				position1427, tokenIndex1427 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1428
																				}
																				position++
																				goto l1427
																			l1428:
																				position, tokenIndex = position1427, tokenIndex1427
																				if buffer[position] != rune('D') {
																					goto l1422
																				}
																				position++
																			}
																		l1427:
																			{
																				position1429, tokenIndex1429 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1430
																				}
																				position++
																				goto l1429
																			l1430:
																				position, tokenIndex = position1429, tokenIndex1429
																				if buffer[position] != rune('D') {
																					goto l1422
																				}
																				position++
																			}
																		l1429:
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
																					goto l1422
																				}
																				position++
																			}
																		l1431:
																			add(rulePegText, position1424)
																		}
																		{
																			add(ruleAction93, position)
																		}
																		add(ruleLddr, position1423)
																	}
																	goto l1377
																l1422:
																	position, tokenIndex = position1377, tokenIndex1377
																	{
																		position1435 := position
																		{
																			position1436 := position
																			{
																				position1437, tokenIndex1437 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1438
																				}
																				position++
																				goto l1437
																			l1438:
																				position, tokenIndex = position1437, tokenIndex1437
																				if buffer[position] != rune('L') {
																					goto l1434
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
																					goto l1434
																				}
																				position++
																			}
																		l1439:
																			{
																				position1441, tokenIndex1441 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1442
																				}
																				position++
																				goto l1441
																			l1442:
																				position, tokenIndex = position1441, tokenIndex1441
																				if buffer[position] != rune('D') {
																					goto l1434
																				}
																				position++
																			}
																		l1441:
																			add(rulePegText, position1436)
																		}
																		{
																			add(ruleAction85, position)
																		}
																		add(ruleLdd, position1435)
																	}
																	goto l1377
																l1434:
																	position, tokenIndex = position1377, tokenIndex1377
																	{
																		position1445 := position
																		{
																			position1446 := position
																			{
																				position1447, tokenIndex1447 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1448
																				}
																				position++
																				goto l1447
																			l1448:
																				position, tokenIndex = position1447, tokenIndex1447
																				if buffer[position] != rune('C') {
																					goto l1444
																				}
																				position++
																			}
																		l1447:
																			{
																				position1449, tokenIndex1449 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1450
																				}
																				position++
																				goto l1449
																			l1450:
																				position, tokenIndex = position1449, tokenIndex1449
																				if buffer[position] != rune('P') {
																					goto l1444
																				}
																				position++
																			}
																		l1449:
																			{
																				position1451, tokenIndex1451 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1452
																				}
																				position++
																				goto l1451
																			l1452:
																				position, tokenIndex = position1451, tokenIndex1451
																				if buffer[position] != rune('D') {
																					goto l1444
																				}
																				position++
																			}
																		l1451:
																			{
																				position1453, tokenIndex1453 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1454
																				}
																				position++
																				goto l1453
																			l1454:
																				position, tokenIndex = position1453, tokenIndex1453
																				if buffer[position] != rune('R') {
																					goto l1444
																				}
																				position++
																			}
																		l1453:
																			add(rulePegText, position1446)
																		}
																		{
																			add(ruleAction94, position)
																		}
																		add(ruleCpdr, position1445)
																	}
																	goto l1377
																l1444:
																	position, tokenIndex = position1377, tokenIndex1377
																	{
																		position1456 := position
																		{
																			position1457 := position
																			{
																				position1458, tokenIndex1458 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1459
																				}
																				position++
																				goto l1458
																			l1459:
																				position, tokenIndex = position1458, tokenIndex1458
																				if buffer[position] != rune('C') {
																					goto l1203
																				}
																				position++
																			}
																		l1458:
																			{
																				position1460, tokenIndex1460 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1461
																				}
																				position++
																				goto l1460
																			l1461:
																				position, tokenIndex = position1460, tokenIndex1460
																				if buffer[position] != rune('P') {
																					goto l1203
																				}
																				position++
																			}
																		l1460:
																			{
																				position1462, tokenIndex1462 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1463
																				}
																				position++
																				goto l1462
																			l1463:
																				position, tokenIndex = position1462, tokenIndex1462
																				if buffer[position] != rune('D') {
																					goto l1203
																				}
																				position++
																			}
																		l1462:
																			add(rulePegText, position1457)
																		}
																		{
																			add(ruleAction86, position)
																		}
																		add(ruleCpd, position1456)
																	}
																}
															l1377:
																add(ruleBlit, position1376)
															}
															break
														}
													}

												}
											l1205:
												add(ruleEDSimple, position1204)
											}
											goto l875
										l1203:
											position, tokenIndex = position875, tokenIndex875
											{
												position1466 := position
												{
													position1467, tokenIndex1467 := position, tokenIndex
													{
														position1469 := position
														{
															position1470 := position
															{
																position1471, tokenIndex1471 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1472
																}
																position++
																goto l1471
															l1472:
																position, tokenIndex = position1471, tokenIndex1471
																if buffer[position] != rune('R') {
																	goto l1468
																}
																position++
															}
														l1471:
															{
																position1473, tokenIndex1473 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1474
																}
																position++
																goto l1473
															l1474:
																position, tokenIndex = position1473, tokenIndex1473
																if buffer[position] != rune('L') {
																	goto l1468
																}
																position++
															}
														l1473:
															{
																position1475, tokenIndex1475 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1476
																}
																position++
																goto l1475
															l1476:
																position, tokenIndex = position1475, tokenIndex1475
																if buffer[position] != rune('C') {
																	goto l1468
																}
																position++
															}
														l1475:
															{
																position1477, tokenIndex1477 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1478
																}
																position++
																goto l1477
															l1478:
																position, tokenIndex = position1477, tokenIndex1477
																if buffer[position] != rune('A') {
																	goto l1468
																}
																position++
															}
														l1477:
															add(rulePegText, position1470)
														}
														{
															add(ruleAction62, position)
														}
														add(ruleRlca, position1469)
													}
													goto l1467
												l1468:
													position, tokenIndex = position1467, tokenIndex1467
													{
														position1481 := position
														{
															position1482 := position
															{
																position1483, tokenIndex1483 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1484
																}
																position++
																goto l1483
															l1484:
																position, tokenIndex = position1483, tokenIndex1483
																if buffer[position] != rune('R') {
																	goto l1480
																}
																position++
															}
														l1483:
															{
																position1485, tokenIndex1485 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1486
																}
																position++
																goto l1485
															l1486:
																position, tokenIndex = position1485, tokenIndex1485
																if buffer[position] != rune('R') {
																	goto l1480
																}
																position++
															}
														l1485:
															{
																position1487, tokenIndex1487 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1488
																}
																position++
																goto l1487
															l1488:
																position, tokenIndex = position1487, tokenIndex1487
																if buffer[position] != rune('C') {
																	goto l1480
																}
																position++
															}
														l1487:
															{
																position1489, tokenIndex1489 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1490
																}
																position++
																goto l1489
															l1490:
																position, tokenIndex = position1489, tokenIndex1489
																if buffer[position] != rune('A') {
																	goto l1480
																}
																position++
															}
														l1489:
															add(rulePegText, position1482)
														}
														{
															add(ruleAction63, position)
														}
														add(ruleRrca, position1481)
													}
													goto l1467
												l1480:
													position, tokenIndex = position1467, tokenIndex1467
													{
														position1493 := position
														{
															position1494 := position
															{
																position1495, tokenIndex1495 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1496
																}
																position++
																goto l1495
															l1496:
																position, tokenIndex = position1495, tokenIndex1495
																if buffer[position] != rune('R') {
																	goto l1492
																}
																position++
															}
														l1495:
															{
																position1497, tokenIndex1497 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1498
																}
																position++
																goto l1497
															l1498:
																position, tokenIndex = position1497, tokenIndex1497
																if buffer[position] != rune('L') {
																	goto l1492
																}
																position++
															}
														l1497:
															{
																position1499, tokenIndex1499 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1500
																}
																position++
																goto l1499
															l1500:
																position, tokenIndex = position1499, tokenIndex1499
																if buffer[position] != rune('A') {
																	goto l1492
																}
																position++
															}
														l1499:
															add(rulePegText, position1494)
														}
														{
															add(ruleAction64, position)
														}
														add(ruleRla, position1493)
													}
													goto l1467
												l1492:
													position, tokenIndex = position1467, tokenIndex1467
													{
														position1503 := position
														{
															position1504 := position
															{
																position1505, tokenIndex1505 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1506
																}
																position++
																goto l1505
															l1506:
																position, tokenIndex = position1505, tokenIndex1505
																if buffer[position] != rune('D') {
																	goto l1502
																}
																position++
															}
														l1505:
															{
																position1507, tokenIndex1507 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1508
																}
																position++
																goto l1507
															l1508:
																position, tokenIndex = position1507, tokenIndex1507
																if buffer[position] != rune('A') {
																	goto l1502
																}
																position++
															}
														l1507:
															{
																position1509, tokenIndex1509 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1510
																}
																position++
																goto l1509
															l1510:
																position, tokenIndex = position1509, tokenIndex1509
																if buffer[position] != rune('A') {
																	goto l1502
																}
																position++
															}
														l1509:
															add(rulePegText, position1504)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleDaa, position1503)
													}
													goto l1467
												l1502:
													position, tokenIndex = position1467, tokenIndex1467
													{
														position1513 := position
														{
															position1514 := position
															{
																position1515, tokenIndex1515 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1516
																}
																position++
																goto l1515
															l1516:
																position, tokenIndex = position1515, tokenIndex1515
																if buffer[position] != rune('C') {
																	goto l1512
																}
																position++
															}
														l1515:
															{
																position1517, tokenIndex1517 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1518
																}
																position++
																goto l1517
															l1518:
																position, tokenIndex = position1517, tokenIndex1517
																if buffer[position] != rune('P') {
																	goto l1512
																}
																position++
															}
														l1517:
															{
																position1519, tokenIndex1519 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1520
																}
																position++
																goto l1519
															l1520:
																position, tokenIndex = position1519, tokenIndex1519
																if buffer[position] != rune('L') {
																	goto l1512
																}
																position++
															}
														l1519:
															add(rulePegText, position1514)
														}
														{
															add(ruleAction67, position)
														}
														add(ruleCpl, position1513)
													}
													goto l1467
												l1512:
													position, tokenIndex = position1467, tokenIndex1467
													{
														position1523 := position
														{
															position1524 := position
															{
																position1525, tokenIndex1525 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1526
																}
																position++
																goto l1525
															l1526:
																position, tokenIndex = position1525, tokenIndex1525
																if buffer[position] != rune('E') {
																	goto l1522
																}
																position++
															}
														l1525:
															{
																position1527, tokenIndex1527 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1528
																}
																position++
																goto l1527
															l1528:
																position, tokenIndex = position1527, tokenIndex1527
																if buffer[position] != rune('X') {
																	goto l1522
																}
																position++
															}
														l1527:
															{
																position1529, tokenIndex1529 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1530
																}
																position++
																goto l1529
															l1530:
																position, tokenIndex = position1529, tokenIndex1529
																if buffer[position] != rune('X') {
																	goto l1522
																}
																position++
															}
														l1529:
															add(rulePegText, position1524)
														}
														{
															add(ruleAction70, position)
														}
														add(ruleExx, position1523)
													}
													goto l1467
												l1522:
													position, tokenIndex = position1467, tokenIndex1467
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1533 := position
																{
																	position1534 := position
																	{
																		position1535, tokenIndex1535 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1536
																		}
																		position++
																		goto l1535
																	l1536:
																		position, tokenIndex = position1535, tokenIndex1535
																		if buffer[position] != rune('E') {
																			goto l1465
																		}
																		position++
																	}
																l1535:
																	{
																		position1537, tokenIndex1537 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1538
																		}
																		position++
																		goto l1537
																	l1538:
																		position, tokenIndex = position1537, tokenIndex1537
																		if buffer[position] != rune('I') {
																			goto l1465
																		}
																		position++
																	}
																l1537:
																	add(rulePegText, position1534)
																}
																{
																	add(ruleAction72, position)
																}
																add(ruleEi, position1533)
															}
															break
														case 'D', 'd':
															{
																position1540 := position
																{
																	position1541 := position
																	{
																		position1542, tokenIndex1542 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1543
																		}
																		position++
																		goto l1542
																	l1543:
																		position, tokenIndex = position1542, tokenIndex1542
																		if buffer[position] != rune('D') {
																			goto l1465
																		}
																		position++
																	}
																l1542:
																	{
																		position1544, tokenIndex1544 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1545
																		}
																		position++
																		goto l1544
																	l1545:
																		position, tokenIndex = position1544, tokenIndex1544
																		if buffer[position] != rune('I') {
																			goto l1465
																		}
																		position++
																	}
																l1544:
																	add(rulePegText, position1541)
																}
																{
																	add(ruleAction71, position)
																}
																add(ruleDi, position1540)
															}
															break
														case 'C', 'c':
															{
																position1547 := position
																{
																	position1548 := position
																	{
																		position1549, tokenIndex1549 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1550
																		}
																		position++
																		goto l1549
																	l1550:
																		position, tokenIndex = position1549, tokenIndex1549
																		if buffer[position] != rune('C') {
																			goto l1465
																		}
																		position++
																	}
																l1549:
																	{
																		position1551, tokenIndex1551 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1552
																		}
																		position++
																		goto l1551
																	l1552:
																		position, tokenIndex = position1551, tokenIndex1551
																		if buffer[position] != rune('C') {
																			goto l1465
																		}
																		position++
																	}
																l1551:
																	{
																		position1553, tokenIndex1553 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1554
																		}
																		position++
																		goto l1553
																	l1554:
																		position, tokenIndex = position1553, tokenIndex1553
																		if buffer[position] != rune('F') {
																			goto l1465
																		}
																		position++
																	}
																l1553:
																	add(rulePegText, position1548)
																}
																{
																	add(ruleAction69, position)
																}
																add(ruleCcf, position1547)
															}
															break
														case 'S', 's':
															{
																position1556 := position
																{
																	position1557 := position
																	{
																		position1558, tokenIndex1558 := position, tokenIndex
																		if buffer[position] != rune('s') {
																			goto l1559
																		}
																		position++
																		goto l1558
																	l1559:
																		position, tokenIndex = position1558, tokenIndex1558
																		if buffer[position] != rune('S') {
																			goto l1465
																		}
																		position++
																	}
																l1558:
																	{
																		position1560, tokenIndex1560 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1561
																		}
																		position++
																		goto l1560
																	l1561:
																		position, tokenIndex = position1560, tokenIndex1560
																		if buffer[position] != rune('C') {
																			goto l1465
																		}
																		position++
																	}
																l1560:
																	{
																		position1562, tokenIndex1562 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1563
																		}
																		position++
																		goto l1562
																	l1563:
																		position, tokenIndex = position1562, tokenIndex1562
																		if buffer[position] != rune('F') {
																			goto l1465
																		}
																		position++
																	}
																l1562:
																	add(rulePegText, position1557)
																}
																{
																	add(ruleAction68, position)
																}
																add(ruleScf, position1556)
															}
															break
														case 'R', 'r':
															{
																position1565 := position
																{
																	position1566 := position
																	{
																		position1567, tokenIndex1567 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1568
																		}
																		position++
																		goto l1567
																	l1568:
																		position, tokenIndex = position1567, tokenIndex1567
																		if buffer[position] != rune('R') {
																			goto l1465
																		}
																		position++
																	}
																l1567:
																	{
																		position1569, tokenIndex1569 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1570
																		}
																		position++
																		goto l1569
																	l1570:
																		position, tokenIndex = position1569, tokenIndex1569
																		if buffer[position] != rune('R') {
																			goto l1465
																		}
																		position++
																	}
																l1569:
																	{
																		position1571, tokenIndex1571 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1572
																		}
																		position++
																		goto l1571
																	l1572:
																		position, tokenIndex = position1571, tokenIndex1571
																		if buffer[position] != rune('A') {
																			goto l1465
																		}
																		position++
																	}
																l1571:
																	add(rulePegText, position1566)
																}
																{
																	add(ruleAction65, position)
																}
																add(ruleRra, position1565)
															}
															break
														case 'H', 'h':
															{
																position1574 := position
																{
																	position1575 := position
																	{
																		position1576, tokenIndex1576 := position, tokenIndex
																		if buffer[position] != rune('h') {
																			goto l1577
																		}
																		position++
																		goto l1576
																	l1577:
																		position, tokenIndex = position1576, tokenIndex1576
																		if buffer[position] != rune('H') {
																			goto l1465
																		}
																		position++
																	}
																l1576:
																	{
																		position1578, tokenIndex1578 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1579
																		}
																		position++
																		goto l1578
																	l1579:
																		position, tokenIndex = position1578, tokenIndex1578
																		if buffer[position] != rune('A') {
																			goto l1465
																		}
																		position++
																	}
																l1578:
																	{
																		position1580, tokenIndex1580 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1581
																		}
																		position++
																		goto l1580
																	l1581:
																		position, tokenIndex = position1580, tokenIndex1580
																		if buffer[position] != rune('L') {
																			goto l1465
																		}
																		position++
																	}
																l1580:
																	{
																		position1582, tokenIndex1582 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1583
																		}
																		position++
																		goto l1582
																	l1583:
																		position, tokenIndex = position1582, tokenIndex1582
																		if buffer[position] != rune('T') {
																			goto l1465
																		}
																		position++
																	}
																l1582:
																	add(rulePegText, position1575)
																}
																{
																	add(ruleAction61, position)
																}
																add(ruleHalt, position1574)
															}
															break
														default:
															{
																position1585 := position
																{
																	position1586 := position
																	{
																		position1587, tokenIndex1587 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1588
																		}
																		position++
																		goto l1587
																	l1588:
																		position, tokenIndex = position1587, tokenIndex1587
																		if buffer[position] != rune('N') {
																			goto l1465
																		}
																		position++
																	}
																l1587:
																	{
																		position1589, tokenIndex1589 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1590
																		}
																		position++
																		goto l1589
																	l1590:
																		position, tokenIndex = position1589, tokenIndex1589
																		if buffer[position] != rune('O') {
																			goto l1465
																		}
																		position++
																	}
																l1589:
																	{
																		position1591, tokenIndex1591 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1592
																		}
																		position++
																		goto l1591
																	l1592:
																		position, tokenIndex = position1591, tokenIndex1591
																		if buffer[position] != rune('P') {
																			goto l1465
																		}
																		position++
																	}
																l1591:
																	add(rulePegText, position1586)
																}
																{
																	add(ruleAction60, position)
																}
																add(ruleNop, position1585)
															}
															break
														}
													}

												}
											l1467:
												add(ruleSimple, position1466)
											}
											goto l875
										l1465:
											position, tokenIndex = position875, tokenIndex875
											{
												position1595 := position
												{
													position1596, tokenIndex1596 := position, tokenIndex
													{
														position1598 := position
														{
															position1599, tokenIndex1599 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1600
															}
															position++
															goto l1599
														l1600:
															position, tokenIndex = position1599, tokenIndex1599
															if buffer[position] != rune('R') {
																goto l1597
															}
															position++
														}
													l1599:
														{
															position1601, tokenIndex1601 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1602
															}
															position++
															goto l1601
														l1602:
															position, tokenIndex = position1601, tokenIndex1601
															if buffer[position] != rune('S') {
																goto l1597
															}
															position++
														}
													l1601:
														{
															position1603, tokenIndex1603 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1604
															}
															position++
															goto l1603
														l1604:
															position, tokenIndex = position1603, tokenIndex1603
															if buffer[position] != rune('T') {
																goto l1597
															}
															position++
														}
													l1603:
														if !_rules[rulews]() {
															goto l1597
														}
														if !_rules[rulen]() {
															goto l1597
														}
														{
															add(ruleAction97, position)
														}
														add(ruleRst, position1598)
													}
													goto l1596
												l1597:
													position, tokenIndex = position1596, tokenIndex1596
													{
														position1607 := position
														{
															position1608, tokenIndex1608 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1609
															}
															position++
															goto l1608
														l1609:
															position, tokenIndex = position1608, tokenIndex1608
															if buffer[position] != rune('J') {
																goto l1606
															}
															position++
														}
													l1608:
														{
															position1610, tokenIndex1610 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l1611
															}
															position++
															goto l1610
														l1611:
															position, tokenIndex = position1610, tokenIndex1610
															if buffer[position] != rune('P') {
																goto l1606
															}
															position++
														}
													l1610:
														if !_rules[rulews]() {
															goto l1606
														}
														{
															position1612, tokenIndex1612 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1612
															}
															if !_rules[rulesep]() {
																goto l1612
															}
															goto l1613
														l1612:
															position, tokenIndex = position1612, tokenIndex1612
														}
													l1613:
														if !_rules[ruleSrc16]() {
															goto l1606
														}
														{
															add(ruleAction100, position)
														}
														add(ruleJp, position1607)
													}
													goto l1596
												l1606:
													position, tokenIndex = position1596, tokenIndex1596
													{
														switch buffer[position] {
														case 'D', 'd':
															{
																position1616 := position
																{
																	position1617, tokenIndex1617 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1618
																	}
																	position++
																	goto l1617
																l1618:
																	position, tokenIndex = position1617, tokenIndex1617
																	if buffer[position] != rune('D') {
																		goto l1594
																	}
																	position++
																}
															l1617:
																{
																	position1619, tokenIndex1619 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1620
																	}
																	position++
																	goto l1619
																l1620:
																	position, tokenIndex = position1619, tokenIndex1619
																	if buffer[position] != rune('J') {
																		goto l1594
																	}
																	position++
																}
															l1619:
																{
																	position1621, tokenIndex1621 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1622
																	}
																	position++
																	goto l1621
																l1622:
																	position, tokenIndex = position1621, tokenIndex1621
																	if buffer[position] != rune('N') {
																		goto l1594
																	}
																	position++
																}
															l1621:
																{
																	position1623, tokenIndex1623 := position, tokenIndex
																	if buffer[position] != rune('z') {
																		goto l1624
																	}
																	position++
																	goto l1623
																l1624:
																	position, tokenIndex = position1623, tokenIndex1623
																	if buffer[position] != rune('Z') {
																		goto l1594
																	}
																	position++
																}
															l1623:
																if !_rules[rulews]() {
																	goto l1594
																}
																if !_rules[ruledisp]() {
																	goto l1594
																}
																{
																	add(ruleAction102, position)
																}
																add(ruleDjnz, position1616)
															}
															break
														case 'J', 'j':
															{
																position1626 := position
																{
																	position1627, tokenIndex1627 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1628
																	}
																	position++
																	goto l1627
																l1628:
																	position, tokenIndex = position1627, tokenIndex1627
																	if buffer[position] != rune('J') {
																		goto l1594
																	}
																	position++
																}
															l1627:
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
																		goto l1594
																	}
																	position++
																}
															l1629:
																if !_rules[rulews]() {
																	goto l1594
																}
																{
																	position1631, tokenIndex1631 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1631
																	}
																	if !_rules[rulesep]() {
																		goto l1631
																	}
																	goto l1632
																l1631:
																	position, tokenIndex = position1631, tokenIndex1631
																}
															l1632:
																if !_rules[ruledisp]() {
																	goto l1594
																}
																{
																	add(ruleAction101, position)
																}
																add(ruleJr, position1626)
															}
															break
														case 'R', 'r':
															{
																position1634 := position
																{
																	position1635, tokenIndex1635 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1636
																	}
																	position++
																	goto l1635
																l1636:
																	position, tokenIndex = position1635, tokenIndex1635
																	if buffer[position] != rune('R') {
																		goto l1594
																	}
																	position++
																}
															l1635:
																{
																	position1637, tokenIndex1637 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1638
																	}
																	position++
																	goto l1637
																l1638:
																	position, tokenIndex = position1637, tokenIndex1637
																	if buffer[position] != rune('E') {
																		goto l1594
																	}
																	position++
																}
															l1637:
																{
																	position1639, tokenIndex1639 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1640
																	}
																	position++
																	goto l1639
																l1640:
																	position, tokenIndex = position1639, tokenIndex1639
																	if buffer[position] != rune('T') {
																		goto l1594
																	}
																	position++
																}
															l1639:
																{
																	position1641, tokenIndex1641 := position, tokenIndex
																	if !_rules[rulews]() {
																		goto l1641
																	}
																	if !_rules[rulecc]() {
																		goto l1641
																	}
																	goto l1642
																l1641:
																	position, tokenIndex = position1641, tokenIndex1641
																}
															l1642:
																{
																	add(ruleAction99, position)
																}
																add(ruleRet, position1634)
															}
															break
														default:
															{
																position1644 := position
																{
																	position1645, tokenIndex1645 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1646
																	}
																	position++
																	goto l1645
																l1646:
																	position, tokenIndex = position1645, tokenIndex1645
																	if buffer[position] != rune('C') {
																		goto l1594
																	}
																	position++
																}
															l1645:
																{
																	position1647, tokenIndex1647 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1648
																	}
																	position++
																	goto l1647
																l1648:
																	position, tokenIndex = position1647, tokenIndex1647
																	if buffer[position] != rune('A') {
																		goto l1594
																	}
																	position++
																}
															l1647:
																{
																	position1649, tokenIndex1649 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1650
																	}
																	position++
																	goto l1649
																l1650:
																	position, tokenIndex = position1649, tokenIndex1649
																	if buffer[position] != rune('L') {
																		goto l1594
																	}
																	position++
																}
															l1649:
																{
																	position1651, tokenIndex1651 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1652
																	}
																	position++
																	goto l1651
																l1652:
																	position, tokenIndex = position1651, tokenIndex1651
																	if buffer[position] != rune('L') {
																		goto l1594
																	}
																	position++
																}
															l1651:
																if !_rules[rulews]() {
																	goto l1594
																}
																{
																	position1653, tokenIndex1653 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1653
																	}
																	if !_rules[rulesep]() {
																		goto l1653
																	}
																	goto l1654
																l1653:
																	position, tokenIndex = position1653, tokenIndex1653
																}
															l1654:
																if !_rules[ruleSrc16]() {
																	goto l1594
																}
																{
																	add(ruleAction98, position)
																}
																add(ruleCall, position1644)
															}
															break
														}
													}

												}
											l1596:
												add(ruleJump, position1595)
											}
											goto l875
										l1594:
											position, tokenIndex = position875, tokenIndex875
											{
												position1656 := position
												{
													position1657, tokenIndex1657 := position, tokenIndex
													{
														position1659 := position
														{
															position1660, tokenIndex1660 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1661
															}
															position++
															goto l1660
														l1661:
															position, tokenIndex = position1660, tokenIndex1660
															if buffer[position] != rune('I') {
																goto l1658
															}
															position++
														}
													l1660:
														{
															position1662, tokenIndex1662 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1663
															}
															position++
															goto l1662
														l1663:
															position, tokenIndex = position1662, tokenIndex1662
															if buffer[position] != rune('N') {
																goto l1658
															}
															position++
														}
													l1662:
														if !_rules[rulews]() {
															goto l1658
														}
														if !_rules[ruleReg8]() {
															goto l1658
														}
														if !_rules[rulesep]() {
															goto l1658
														}
														if !_rules[rulePort]() {
															goto l1658
														}
														{
															add(ruleAction103, position)
														}
														add(ruleIN, position1659)
													}
													goto l1657
												l1658:
													position, tokenIndex = position1657, tokenIndex1657
													{
														position1665 := position
														{
															position1666, tokenIndex1666 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1667
															}
															position++
															goto l1666
														l1667:
															position, tokenIndex = position1666, tokenIndex1666
															if buffer[position] != rune('O') {
																goto l854
															}
															position++
														}
													l1666:
														{
															position1668, tokenIndex1668 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1669
															}
															position++
															goto l1668
														l1669:
															position, tokenIndex = position1668, tokenIndex1668
															if buffer[position] != rune('U') {
																goto l854
															}
															position++
														}
													l1668:
														{
															position1670, tokenIndex1670 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1671
															}
															position++
															goto l1670
														l1671:
															position, tokenIndex = position1670, tokenIndex1670
															if buffer[position] != rune('T') {
																goto l854
															}
															position++
														}
													l1670:
														if !_rules[rulews]() {
															goto l854
														}
														if !_rules[rulePort]() {
															goto l854
														}
														if !_rules[rulesep]() {
															goto l854
														}
														if !_rules[ruleReg8]() {
															goto l854
														}
														{
															add(ruleAction104, position)
														}
														add(ruleOUT, position1665)
													}
												}
											l1657:
												add(ruleIO, position1656)
											}
										}
									l875:
										add(ruleInstruction, position874)
									}
								}
							l857:
								add(ruleStatement, position856)
							}
							goto l855
						l854:
							position, tokenIndex = position854, tokenIndex854
						}
					l855:
						{
							position1673, tokenIndex1673 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1673
							}
							goto l1674
						l1673:
							position, tokenIndex = position1673, tokenIndex1673
						}
					l1674:
						{
							position1675, tokenIndex1675 := position, tokenIndex
							{
								position1677 := position
								{
									position1678, tokenIndex1678 := position, tokenIndex
									if buffer[position] != rune(';') {
										goto l1679
									}
									position++
									goto l1678
								l1679:
									position, tokenIndex = position1678, tokenIndex1678
									if buffer[position] != rune('#') {
										goto l1675
									}
									position++
								}
							l1678:
							l1680:
								{
									position1681, tokenIndex1681 := position, tokenIndex
									{
										position1682, tokenIndex1682 := position, tokenIndex
										if buffer[position] != rune('\n') {
											goto l1682
										}
										position++
										goto l1681
									l1682:
										position, tokenIndex = position1682, tokenIndex1682
									}
									if !matchDot() {
										goto l1681
									}
									goto l1680
								l1681:
									position, tokenIndex = position1681, tokenIndex1681
								}
								add(ruleComment, position1677)
							}
							goto l1676
						l1675:
							position, tokenIndex = position1675, tokenIndex1675
						}
					l1676:
						{
							position1683, tokenIndex1683 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1683
							}
							goto l1684
						l1683:
							position, tokenIndex = position1683, tokenIndex1683
						}
					l1684:
						{
							position1685, tokenIndex1685 := position, tokenIndex
							{
								position1687, tokenIndex1687 := position, tokenIndex
								if buffer[position] != rune('\r') {
									goto l1687
								}
								position++
								goto l1688
							l1687:
								position, tokenIndex = position1687, tokenIndex1687
							}
						l1688:
							if buffer[position] != rune('\n') {
								goto l1686
							}
							position++
							goto l1685
						l1686:
							position, tokenIndex = position1685, tokenIndex1685
							if buffer[position] != rune(':') {
								goto l3
							}
							position++
						}
					l1685:
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position847)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position1690, tokenIndex1690 := position, tokenIndex
					if !matchDot() {
						goto l1690
					}
					goto l0
				l1690:
					position, tokenIndex = position1690, tokenIndex1690
				}
				add(ruleProgram, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 Line <- <(LabelDefn? ws* Statement? ws? Comment? ws? (('\r'? '\n') / ':') Action0)> */
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
			position1698, tokenIndex1698 := position, tokenIndex
			{
				position1699 := position
				{
					position1700 := position
					if !_rules[rulealpha]() {
						goto l1698
					}
				l1701:
					{
						position1702, tokenIndex1702 := position, tokenIndex
						{
							position1703 := position
							{
								position1704, tokenIndex1704 := position, tokenIndex
								if !_rules[rulealpha]() {
									goto l1705
								}
								goto l1704
							l1705:
								position, tokenIndex = position1704, tokenIndex1704
								{
									position1706 := position
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1702
									}
									position++
									add(rulenum, position1706)
								}
							}
						l1704:
							add(rulealphanum, position1703)
						}
						goto l1701
					l1702:
						position, tokenIndex = position1702, tokenIndex1702
					}
					add(rulePegText, position1700)
				}
				add(ruleLabelText, position1699)
			}
			return true
		l1698:
			position, tokenIndex = position1698, tokenIndex1698
			return false
		},
		/* 9 alphanum <- <(alpha / num)> */
		nil,
		/* 10 alpha <- <([a-z] / [A-Z])> */
		func() bool {
			position1708, tokenIndex1708 := position, tokenIndex
			{
				position1709 := position
				{
					position1710, tokenIndex1710 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1711
					}
					position++
					goto l1710
				l1711:
					position, tokenIndex = position1710, tokenIndex1710
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1708
					}
					position++
				}
			l1710:
				add(rulealpha, position1709)
			}
			return true
		l1708:
			position, tokenIndex = position1708, tokenIndex1708
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
			position1735, tokenIndex1735 := position, tokenIndex
			{
				position1736 := position
				{
					position1737, tokenIndex1737 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1738
					}
					goto l1737
				l1738:
					position, tokenIndex = position1737, tokenIndex1737
					if !_rules[ruleReg8]() {
						goto l1739
					}
					goto l1737
				l1739:
					position, tokenIndex = position1737, tokenIndex1737
					if !_rules[ruleReg16Contents]() {
						goto l1740
					}
					goto l1737
				l1740:
					position, tokenIndex = position1737, tokenIndex1737
					if !_rules[rulenn_contents]() {
						goto l1735
					}
				}
			l1737:
				{
					add(ruleAction18, position)
				}
				add(ruleSrc8, position1736)
			}
			return true
		l1735:
			position, tokenIndex = position1735, tokenIndex1735
			return false
		},
		/* 35 Loc8 <- <((Reg8 / Reg16Contents) Action19)> */
		func() bool {
			position1742, tokenIndex1742 := position, tokenIndex
			{
				position1743 := position
				{
					position1744, tokenIndex1744 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1745
					}
					goto l1744
				l1745:
					position, tokenIndex = position1744, tokenIndex1744
					if !_rules[ruleReg16Contents]() {
						goto l1742
					}
				}
			l1744:
				{
					add(ruleAction19, position)
				}
				add(ruleLoc8, position1743)
			}
			return true
		l1742:
			position, tokenIndex = position1742, tokenIndex1742
			return false
		},
		/* 36 Copy8 <- <(Reg8 Action20)> */
		func() bool {
			position1747, tokenIndex1747 := position, tokenIndex
			{
				position1748 := position
				if !_rules[ruleReg8]() {
					goto l1747
				}
				{
					add(ruleAction20, position)
				}
				add(ruleCopy8, position1748)
			}
			return true
		l1747:
			position, tokenIndex = position1747, tokenIndex1747
			return false
		},
		/* 37 ILoc8 <- <(IReg8 Action21)> */
		func() bool {
			position1750, tokenIndex1750 := position, tokenIndex
			{
				position1751 := position
				if !_rules[ruleIReg8]() {
					goto l1750
				}
				{
					add(ruleAction21, position)
				}
				add(ruleILoc8, position1751)
			}
			return true
		l1750:
			position, tokenIndex = position1750, tokenIndex1750
			return false
		},
		/* 38 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action22)> */
		func() bool {
			position1753, tokenIndex1753 := position, tokenIndex
			{
				position1754 := position
				{
					position1755 := position
					{
						position1756, tokenIndex1756 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1757
						}
						goto l1756
					l1757:
						position, tokenIndex = position1756, tokenIndex1756
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1759 := position
									{
										position1760, tokenIndex1760 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1761
										}
										position++
										goto l1760
									l1761:
										position, tokenIndex = position1760, tokenIndex1760
										if buffer[position] != rune('R') {
											goto l1753
										}
										position++
									}
								l1760:
									add(ruleR, position1759)
								}
								break
							case 'I', 'i':
								{
									position1762 := position
									{
										position1763, tokenIndex1763 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1764
										}
										position++
										goto l1763
									l1764:
										position, tokenIndex = position1763, tokenIndex1763
										if buffer[position] != rune('I') {
											goto l1753
										}
										position++
									}
								l1763:
									add(ruleI, position1762)
								}
								break
							case 'L', 'l':
								{
									position1765 := position
									{
										position1766, tokenIndex1766 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1767
										}
										position++
										goto l1766
									l1767:
										position, tokenIndex = position1766, tokenIndex1766
										if buffer[position] != rune('L') {
											goto l1753
										}
										position++
									}
								l1766:
									add(ruleL, position1765)
								}
								break
							case 'H', 'h':
								{
									position1768 := position
									{
										position1769, tokenIndex1769 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1770
										}
										position++
										goto l1769
									l1770:
										position, tokenIndex = position1769, tokenIndex1769
										if buffer[position] != rune('H') {
											goto l1753
										}
										position++
									}
								l1769:
									add(ruleH, position1768)
								}
								break
							case 'E', 'e':
								{
									position1771 := position
									{
										position1772, tokenIndex1772 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1773
										}
										position++
										goto l1772
									l1773:
										position, tokenIndex = position1772, tokenIndex1772
										if buffer[position] != rune('E') {
											goto l1753
										}
										position++
									}
								l1772:
									add(ruleE, position1771)
								}
								break
							case 'D', 'd':
								{
									position1774 := position
									{
										position1775, tokenIndex1775 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1776
										}
										position++
										goto l1775
									l1776:
										position, tokenIndex = position1775, tokenIndex1775
										if buffer[position] != rune('D') {
											goto l1753
										}
										position++
									}
								l1775:
									add(ruleD, position1774)
								}
								break
							case 'C', 'c':
								{
									position1777 := position
									{
										position1778, tokenIndex1778 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1779
										}
										position++
										goto l1778
									l1779:
										position, tokenIndex = position1778, tokenIndex1778
										if buffer[position] != rune('C') {
											goto l1753
										}
										position++
									}
								l1778:
									add(ruleC, position1777)
								}
								break
							case 'B', 'b':
								{
									position1780 := position
									{
										position1781, tokenIndex1781 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1782
										}
										position++
										goto l1781
									l1782:
										position, tokenIndex = position1781, tokenIndex1781
										if buffer[position] != rune('B') {
											goto l1753
										}
										position++
									}
								l1781:
									add(ruleB, position1780)
								}
								break
							case 'F', 'f':
								{
									position1783 := position
									{
										position1784, tokenIndex1784 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1785
										}
										position++
										goto l1784
									l1785:
										position, tokenIndex = position1784, tokenIndex1784
										if buffer[position] != rune('F') {
											goto l1753
										}
										position++
									}
								l1784:
									add(ruleF, position1783)
								}
								break
							default:
								{
									position1786 := position
									{
										position1787, tokenIndex1787 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1788
										}
										position++
										goto l1787
									l1788:
										position, tokenIndex = position1787, tokenIndex1787
										if buffer[position] != rune('A') {
											goto l1753
										}
										position++
									}
								l1787:
									add(ruleA, position1786)
								}
								break
							}
						}

					}
				l1756:
					add(rulePegText, position1755)
				}
				{
					add(ruleAction22, position)
				}
				add(ruleReg8, position1754)
			}
			return true
		l1753:
			position, tokenIndex = position1753, tokenIndex1753
			return false
		},
		/* 39 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action23)> */
		func() bool {
			position1790, tokenIndex1790 := position, tokenIndex
			{
				position1791 := position
				{
					position1792 := position
					{
						position1793, tokenIndex1793 := position, tokenIndex
						{
							position1795 := position
							{
								position1796, tokenIndex1796 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1797
								}
								position++
								goto l1796
							l1797:
								position, tokenIndex = position1796, tokenIndex1796
								if buffer[position] != rune('I') {
									goto l1794
								}
								position++
							}
						l1796:
							{
								position1798, tokenIndex1798 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1799
								}
								position++
								goto l1798
							l1799:
								position, tokenIndex = position1798, tokenIndex1798
								if buffer[position] != rune('X') {
									goto l1794
								}
								position++
							}
						l1798:
							{
								position1800, tokenIndex1800 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1801
								}
								position++
								goto l1800
							l1801:
								position, tokenIndex = position1800, tokenIndex1800
								if buffer[position] != rune('H') {
									goto l1794
								}
								position++
							}
						l1800:
							add(ruleIXH, position1795)
						}
						goto l1793
					l1794:
						position, tokenIndex = position1793, tokenIndex1793
						{
							position1803 := position
							{
								position1804, tokenIndex1804 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1805
								}
								position++
								goto l1804
							l1805:
								position, tokenIndex = position1804, tokenIndex1804
								if buffer[position] != rune('I') {
									goto l1802
								}
								position++
							}
						l1804:
							{
								position1806, tokenIndex1806 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1807
								}
								position++
								goto l1806
							l1807:
								position, tokenIndex = position1806, tokenIndex1806
								if buffer[position] != rune('X') {
									goto l1802
								}
								position++
							}
						l1806:
							{
								position1808, tokenIndex1808 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1809
								}
								position++
								goto l1808
							l1809:
								position, tokenIndex = position1808, tokenIndex1808
								if buffer[position] != rune('L') {
									goto l1802
								}
								position++
							}
						l1808:
							add(ruleIXL, position1803)
						}
						goto l1793
					l1802:
						position, tokenIndex = position1793, tokenIndex1793
						{
							position1811 := position
							{
								position1812, tokenIndex1812 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1813
								}
								position++
								goto l1812
							l1813:
								position, tokenIndex = position1812, tokenIndex1812
								if buffer[position] != rune('I') {
									goto l1810
								}
								position++
							}
						l1812:
							{
								position1814, tokenIndex1814 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1815
								}
								position++
								goto l1814
							l1815:
								position, tokenIndex = position1814, tokenIndex1814
								if buffer[position] != rune('Y') {
									goto l1810
								}
								position++
							}
						l1814:
							{
								position1816, tokenIndex1816 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1817
								}
								position++
								goto l1816
							l1817:
								position, tokenIndex = position1816, tokenIndex1816
								if buffer[position] != rune('H') {
									goto l1810
								}
								position++
							}
						l1816:
							add(ruleIYH, position1811)
						}
						goto l1793
					l1810:
						position, tokenIndex = position1793, tokenIndex1793
						{
							position1818 := position
							{
								position1819, tokenIndex1819 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1820
								}
								position++
								goto l1819
							l1820:
								position, tokenIndex = position1819, tokenIndex1819
								if buffer[position] != rune('I') {
									goto l1790
								}
								position++
							}
						l1819:
							{
								position1821, tokenIndex1821 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1822
								}
								position++
								goto l1821
							l1822:
								position, tokenIndex = position1821, tokenIndex1821
								if buffer[position] != rune('Y') {
									goto l1790
								}
								position++
							}
						l1821:
							{
								position1823, tokenIndex1823 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1824
								}
								position++
								goto l1823
							l1824:
								position, tokenIndex = position1823, tokenIndex1823
								if buffer[position] != rune('L') {
									goto l1790
								}
								position++
							}
						l1823:
							add(ruleIYL, position1818)
						}
					}
				l1793:
					add(rulePegText, position1792)
				}
				{
					add(ruleAction23, position)
				}
				add(ruleIReg8, position1791)
			}
			return true
		l1790:
			position, tokenIndex = position1790, tokenIndex1790
			return false
		},
		/* 40 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action24)> */
		func() bool {
			position1826, tokenIndex1826 := position, tokenIndex
			{
				position1827 := position
				{
					position1828, tokenIndex1828 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1829
					}
					goto l1828
				l1829:
					position, tokenIndex = position1828, tokenIndex1828
					if !_rules[rulenn_contents]() {
						goto l1830
					}
					goto l1828
				l1830:
					position, tokenIndex = position1828, tokenIndex1828
					if !_rules[ruleReg16Contents]() {
						goto l1826
					}
				}
			l1828:
				{
					add(ruleAction24, position)
				}
				add(ruleDst16, position1827)
			}
			return true
		l1826:
			position, tokenIndex = position1826, tokenIndex1826
			return false
		},
		/* 41 Src16 <- <((Reg16 / nn / nn_contents) Action25)> */
		func() bool {
			position1832, tokenIndex1832 := position, tokenIndex
			{
				position1833 := position
				{
					position1834, tokenIndex1834 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1835
					}
					goto l1834
				l1835:
					position, tokenIndex = position1834, tokenIndex1834
					if !_rules[rulenn]() {
						goto l1836
					}
					goto l1834
				l1836:
					position, tokenIndex = position1834, tokenIndex1834
					if !_rules[rulenn_contents]() {
						goto l1832
					}
				}
			l1834:
				{
					add(ruleAction25, position)
				}
				add(ruleSrc16, position1833)
			}
			return true
		l1832:
			position, tokenIndex = position1832, tokenIndex1832
			return false
		},
		/* 42 Loc16 <- <(Reg16 Action26)> */
		func() bool {
			position1838, tokenIndex1838 := position, tokenIndex
			{
				position1839 := position
				if !_rules[ruleReg16]() {
					goto l1838
				}
				{
					add(ruleAction26, position)
				}
				add(ruleLoc16, position1839)
			}
			return true
		l1838:
			position, tokenIndex = position1838, tokenIndex1838
			return false
		},
		/* 43 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action27)> */
		func() bool {
			position1841, tokenIndex1841 := position, tokenIndex
			{
				position1842 := position
				{
					position1843 := position
					{
						position1844, tokenIndex1844 := position, tokenIndex
						{
							position1846 := position
							{
								position1847, tokenIndex1847 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1848
								}
								position++
								goto l1847
							l1848:
								position, tokenIndex = position1847, tokenIndex1847
								if buffer[position] != rune('A') {
									goto l1845
								}
								position++
							}
						l1847:
							{
								position1849, tokenIndex1849 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1850
								}
								position++
								goto l1849
							l1850:
								position, tokenIndex = position1849, tokenIndex1849
								if buffer[position] != rune('F') {
									goto l1845
								}
								position++
							}
						l1849:
							if buffer[position] != rune('\'') {
								goto l1845
							}
							position++
							add(ruleAF_PRIME, position1846)
						}
						goto l1844
					l1845:
						position, tokenIndex = position1844, tokenIndex1844
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1841
								}
								break
							case 'S', 's':
								{
									position1852 := position
									{
										position1853, tokenIndex1853 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1854
										}
										position++
										goto l1853
									l1854:
										position, tokenIndex = position1853, tokenIndex1853
										if buffer[position] != rune('S') {
											goto l1841
										}
										position++
									}
								l1853:
									{
										position1855, tokenIndex1855 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1856
										}
										position++
										goto l1855
									l1856:
										position, tokenIndex = position1855, tokenIndex1855
										if buffer[position] != rune('P') {
											goto l1841
										}
										position++
									}
								l1855:
									add(ruleSP, position1852)
								}
								break
							case 'H', 'h':
								{
									position1857 := position
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
											goto l1841
										}
										position++
									}
								l1858:
									{
										position1860, tokenIndex1860 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1861
										}
										position++
										goto l1860
									l1861:
										position, tokenIndex = position1860, tokenIndex1860
										if buffer[position] != rune('L') {
											goto l1841
										}
										position++
									}
								l1860:
									add(ruleHL, position1857)
								}
								break
							case 'D', 'd':
								{
									position1862 := position
									{
										position1863, tokenIndex1863 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1864
										}
										position++
										goto l1863
									l1864:
										position, tokenIndex = position1863, tokenIndex1863
										if buffer[position] != rune('D') {
											goto l1841
										}
										position++
									}
								l1863:
									{
										position1865, tokenIndex1865 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1866
										}
										position++
										goto l1865
									l1866:
										position, tokenIndex = position1865, tokenIndex1865
										if buffer[position] != rune('E') {
											goto l1841
										}
										position++
									}
								l1865:
									add(ruleDE, position1862)
								}
								break
							case 'B', 'b':
								{
									position1867 := position
									{
										position1868, tokenIndex1868 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1869
										}
										position++
										goto l1868
									l1869:
										position, tokenIndex = position1868, tokenIndex1868
										if buffer[position] != rune('B') {
											goto l1841
										}
										position++
									}
								l1868:
									{
										position1870, tokenIndex1870 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1871
										}
										position++
										goto l1870
									l1871:
										position, tokenIndex = position1870, tokenIndex1870
										if buffer[position] != rune('C') {
											goto l1841
										}
										position++
									}
								l1870:
									add(ruleBC, position1867)
								}
								break
							default:
								{
									position1872 := position
									{
										position1873, tokenIndex1873 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1874
										}
										position++
										goto l1873
									l1874:
										position, tokenIndex = position1873, tokenIndex1873
										if buffer[position] != rune('A') {
											goto l1841
										}
										position++
									}
								l1873:
									{
										position1875, tokenIndex1875 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1876
										}
										position++
										goto l1875
									l1876:
										position, tokenIndex = position1875, tokenIndex1875
										if buffer[position] != rune('F') {
											goto l1841
										}
										position++
									}
								l1875:
									add(ruleAF, position1872)
								}
								break
							}
						}

					}
				l1844:
					add(rulePegText, position1843)
				}
				{
					add(ruleAction27, position)
				}
				add(ruleReg16, position1842)
			}
			return true
		l1841:
			position, tokenIndex = position1841, tokenIndex1841
			return false
		},
		/* 44 IReg16 <- <(<(IX / IY)> Action28)> */
		func() bool {
			position1878, tokenIndex1878 := position, tokenIndex
			{
				position1879 := position
				{
					position1880 := position
					{
						position1881, tokenIndex1881 := position, tokenIndex
						{
							position1883 := position
							{
								position1884, tokenIndex1884 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1885
								}
								position++
								goto l1884
							l1885:
								position, tokenIndex = position1884, tokenIndex1884
								if buffer[position] != rune('I') {
									goto l1882
								}
								position++
							}
						l1884:
							{
								position1886, tokenIndex1886 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1887
								}
								position++
								goto l1886
							l1887:
								position, tokenIndex = position1886, tokenIndex1886
								if buffer[position] != rune('X') {
									goto l1882
								}
								position++
							}
						l1886:
							add(ruleIX, position1883)
						}
						goto l1881
					l1882:
						position, tokenIndex = position1881, tokenIndex1881
						{
							position1888 := position
							{
								position1889, tokenIndex1889 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1890
								}
								position++
								goto l1889
							l1890:
								position, tokenIndex = position1889, tokenIndex1889
								if buffer[position] != rune('I') {
									goto l1878
								}
								position++
							}
						l1889:
							{
								position1891, tokenIndex1891 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1892
								}
								position++
								goto l1891
							l1892:
								position, tokenIndex = position1891, tokenIndex1891
								if buffer[position] != rune('Y') {
									goto l1878
								}
								position++
							}
						l1891:
							add(ruleIY, position1888)
						}
					}
				l1881:
					add(rulePegText, position1880)
				}
				{
					add(ruleAction28, position)
				}
				add(ruleIReg16, position1879)
			}
			return true
		l1878:
			position, tokenIndex = position1878, tokenIndex1878
			return false
		},
		/* 45 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1894, tokenIndex1894 := position, tokenIndex
			{
				position1895 := position
				{
					position1896, tokenIndex1896 := position, tokenIndex
					{
						position1898 := position
						if buffer[position] != rune('(') {
							goto l1897
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1897
						}
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
						if !_rules[ruledisp]() {
							goto l1897
						}
						{
							position1901, tokenIndex1901 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1901
							}
							goto l1902
						l1901:
							position, tokenIndex = position1901, tokenIndex1901
						}
					l1902:
						if buffer[position] != rune(')') {
							goto l1897
						}
						position++
						{
							add(ruleAction30, position)
						}
						add(ruleIndexedR16C, position1898)
					}
					goto l1896
				l1897:
					position, tokenIndex = position1896, tokenIndex1896
					{
						position1904 := position
						if buffer[position] != rune('(') {
							goto l1894
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1894
						}
						if buffer[position] != rune(')') {
							goto l1894
						}
						position++
						{
							add(ruleAction29, position)
						}
						add(rulePlainR16C, position1904)
					}
				}
			l1896:
				add(ruleReg16Contents, position1895)
			}
			return true
		l1894:
			position, tokenIndex = position1894, tokenIndex1894
			return false
		},
		/* 46 PlainR16C <- <('(' Reg16 ')' Action29)> */
		nil,
		/* 47 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action30)> */
		nil,
		/* 48 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1908, tokenIndex1908 := position, tokenIndex
			{
				position1909 := position
				{
					position1910, tokenIndex1910 := position, tokenIndex
					{
						position1912 := position
						{
							position1913 := position
							if !_rules[rulehexdigit]() {
								goto l1911
							}
							if !_rules[rulehexdigit]() {
								goto l1911
							}
							add(rulePegText, position1913)
						}
						{
							position1914, tokenIndex1914 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1915
							}
							position++
							goto l1914
						l1915:
							position, tokenIndex = position1914, tokenIndex1914
							if buffer[position] != rune('H') {
								goto l1911
							}
							position++
						}
					l1914:
						{
							add(ruleAction34, position)
						}
						add(rulehexByteH, position1912)
					}
					goto l1910
				l1911:
					position, tokenIndex = position1910, tokenIndex1910
					{
						position1918 := position
						if buffer[position] != rune('0') {
							goto l1917
						}
						position++
						{
							position1919, tokenIndex1919 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1920
							}
							position++
							goto l1919
						l1920:
							position, tokenIndex = position1919, tokenIndex1919
							if buffer[position] != rune('X') {
								goto l1917
							}
							position++
						}
					l1919:
						{
							position1921 := position
							if !_rules[rulehexdigit]() {
								goto l1917
							}
							if !_rules[rulehexdigit]() {
								goto l1917
							}
							add(rulePegText, position1921)
						}
						{
							add(ruleAction35, position)
						}
						add(rulehexByte0x, position1918)
					}
					goto l1910
				l1917:
					position, tokenIndex = position1910, tokenIndex1910
					{
						position1923 := position
						{
							position1924 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1908
							}
							position++
						l1925:
							{
								position1926, tokenIndex1926 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1926
								}
								position++
								goto l1925
							l1926:
								position, tokenIndex = position1926, tokenIndex1926
							}
							add(rulePegText, position1924)
						}
						{
							add(ruleAction36, position)
						}
						add(ruledecimalByte, position1923)
					}
				}
			l1910:
				add(rulen, position1909)
			}
			return true
		l1908:
			position, tokenIndex = position1908, tokenIndex1908
			return false
		},
		/* 49 nn <- <(LabelNN / hexWordH / hexWord0x)> */
		func() bool {
			position1928, tokenIndex1928 := position, tokenIndex
			{
				position1929 := position
				{
					position1930, tokenIndex1930 := position, tokenIndex
					{
						position1932 := position
						{
							position1933 := position
							if !_rules[ruleLabelText]() {
								goto l1931
							}
							add(rulePegText, position1933)
						}
						{
							add(ruleAction37, position)
						}
						add(ruleLabelNN, position1932)
					}
					goto l1930
				l1931:
					position, tokenIndex = position1930, tokenIndex1930
					{
						position1936 := position
						{
							position1937 := position
							if !_rules[rulehexdigit]() {
								goto l1935
							}
							if !_rules[rulehexdigit]() {
								goto l1935
							}
							if !_rules[rulehexdigit]() {
								goto l1935
							}
							if !_rules[rulehexdigit]() {
								goto l1935
							}
							add(rulePegText, position1937)
						}
						{
							position1938, tokenIndex1938 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1939
							}
							position++
							goto l1938
						l1939:
							position, tokenIndex = position1938, tokenIndex1938
							if buffer[position] != rune('H') {
								goto l1935
							}
							position++
						}
					l1938:
						{
							add(ruleAction38, position)
						}
						add(rulehexWordH, position1936)
					}
					goto l1930
				l1935:
					position, tokenIndex = position1930, tokenIndex1930
					{
						position1941 := position
						if buffer[position] != rune('0') {
							goto l1928
						}
						position++
						{
							position1942, tokenIndex1942 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1943
							}
							position++
							goto l1942
						l1943:
							position, tokenIndex = position1942, tokenIndex1942
							if buffer[position] != rune('X') {
								goto l1928
							}
							position++
						}
					l1942:
						{
							position1944 := position
							if !_rules[rulehexdigit]() {
								goto l1928
							}
							if !_rules[rulehexdigit]() {
								goto l1928
							}
							if !_rules[rulehexdigit]() {
								goto l1928
							}
							if !_rules[rulehexdigit]() {
								goto l1928
							}
							add(rulePegText, position1944)
						}
						{
							add(ruleAction39, position)
						}
						add(rulehexWord0x, position1941)
					}
				}
			l1930:
				add(rulenn, position1929)
			}
			return true
		l1928:
			position, tokenIndex = position1928, tokenIndex1928
			return false
		},
		/* 50 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position1946, tokenIndex1946 := position, tokenIndex
			{
				position1947 := position
				{
					position1948, tokenIndex1948 := position, tokenIndex
					{
						position1950 := position
						{
							position1951 := position
							{
								position1952, tokenIndex1952 := position, tokenIndex
								{
									position1954, tokenIndex1954 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1955
									}
									position++
									goto l1954
								l1955:
									position, tokenIndex = position1954, tokenIndex1954
									if buffer[position] != rune('+') {
										goto l1952
									}
									position++
								}
							l1954:
								goto l1953
							l1952:
								position, tokenIndex = position1952, tokenIndex1952
							}
						l1953:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1949
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1949
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1949
									}
									position++
									break
								}
							}

						l1956:
							{
								position1957, tokenIndex1957 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1957
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1957
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1957
										}
										position++
										break
									}
								}

								goto l1956
							l1957:
								position, tokenIndex = position1957, tokenIndex1957
							}
							add(rulePegText, position1951)
						}
						{
							position1960, tokenIndex1960 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1961
							}
							position++
							goto l1960
						l1961:
							position, tokenIndex = position1960, tokenIndex1960
							if buffer[position] != rune('H') {
								goto l1949
							}
							position++
						}
					l1960:
						{
							add(ruleAction32, position)
						}
						add(rulesignedHexByteH, position1950)
					}
					goto l1948
				l1949:
					position, tokenIndex = position1948, tokenIndex1948
					{
						position1964 := position
						{
							position1965 := position
							{
								position1966, tokenIndex1966 := position, tokenIndex
								{
									position1968, tokenIndex1968 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1969
									}
									position++
									goto l1968
								l1969:
									position, tokenIndex = position1968, tokenIndex1968
									if buffer[position] != rune('+') {
										goto l1966
									}
									position++
								}
							l1968:
								goto l1967
							l1966:
								position, tokenIndex = position1966, tokenIndex1966
							}
						l1967:
							if buffer[position] != rune('0') {
								goto l1963
							}
							position++
							{
								position1970, tokenIndex1970 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1971
								}
								position++
								goto l1970
							l1971:
								position, tokenIndex = position1970, tokenIndex1970
								if buffer[position] != rune('X') {
									goto l1963
								}
								position++
							}
						l1970:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1963
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1963
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1963
									}
									position++
									break
								}
							}

						l1972:
							{
								position1973, tokenIndex1973 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1973
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1973
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1973
										}
										position++
										break
									}
								}

								goto l1972
							l1973:
								position, tokenIndex = position1973, tokenIndex1973
							}
							add(rulePegText, position1965)
						}
						{
							add(ruleAction33, position)
						}
						add(rulesignedHexByte0x, position1964)
					}
					goto l1948
				l1963:
					position, tokenIndex = position1948, tokenIndex1948
					{
						position1977 := position
						{
							position1978 := position
							{
								position1979, tokenIndex1979 := position, tokenIndex
								{
									position1981, tokenIndex1981 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1982
									}
									position++
									goto l1981
								l1982:
									position, tokenIndex = position1981, tokenIndex1981
									if buffer[position] != rune('+') {
										goto l1979
									}
									position++
								}
							l1981:
								goto l1980
							l1979:
								position, tokenIndex = position1979, tokenIndex1979
							}
						l1980:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1946
							}
							position++
						l1983:
							{
								position1984, tokenIndex1984 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1984
								}
								position++
								goto l1983
							l1984:
								position, tokenIndex = position1984, tokenIndex1984
							}
							add(rulePegText, position1978)
						}
						{
							add(ruleAction31, position)
						}
						add(rulesignedDecimalByte, position1977)
					}
				}
			l1948:
				add(ruledisp, position1947)
			}
			return true
		l1946:
			position, tokenIndex = position1946, tokenIndex1946
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
			position1995, tokenIndex1995 := position, tokenIndex
			{
				position1996 := position
				if buffer[position] != rune('(') {
					goto l1995
				}
				position++
				if !_rules[rulenn]() {
					goto l1995
				}
				if buffer[position] != rune(')') {
					goto l1995
				}
				position++
				{
					add(ruleAction40, position)
				}
				add(rulenn_contents, position1996)
			}
			return true
		l1995:
			position, tokenIndex = position1995, tokenIndex1995
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
			position2071, tokenIndex2071 := position, tokenIndex
			{
				position2072 := position
				{
					position2073, tokenIndex2073 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2074
					}
					position++
					{
						position2075, tokenIndex2075 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l2076
						}
						position++
						goto l2075
					l2076:
						position, tokenIndex = position2075, tokenIndex2075
						if buffer[position] != rune('C') {
							goto l2074
						}
						position++
					}
				l2075:
					if buffer[position] != rune(')') {
						goto l2074
					}
					position++
					goto l2073
				l2074:
					position, tokenIndex = position2073, tokenIndex2073
					if buffer[position] != rune('(') {
						goto l2071
					}
					position++
					if !_rules[rulen]() {
						goto l2071
					}
					if buffer[position] != rune(')') {
						goto l2071
					}
					position++
				}
			l2073:
				add(rulePort, position2072)
			}
			return true
		l2071:
			position, tokenIndex = position2071, tokenIndex2071
			return false
		},
		/* 135 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2077, tokenIndex2077 := position, tokenIndex
			{
				position2078 := position
				{
					position2079, tokenIndex2079 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2079
					}
					goto l2080
				l2079:
					position, tokenIndex = position2079, tokenIndex2079
				}
			l2080:
				if buffer[position] != rune(',') {
					goto l2077
				}
				position++
				{
					position2081, tokenIndex2081 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2081
					}
					goto l2082
				l2081:
					position, tokenIndex = position2081, tokenIndex2081
				}
			l2082:
				add(rulesep, position2078)
			}
			return true
		l2077:
			position, tokenIndex = position2077, tokenIndex2077
			return false
		},
		/* 136 ws <- <(' ' / '\t')+> */
		func() bool {
			position2083, tokenIndex2083 := position, tokenIndex
			{
				position2084 := position
				{
					position2087, tokenIndex2087 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2088
					}
					position++
					goto l2087
				l2088:
					position, tokenIndex = position2087, tokenIndex2087
					if buffer[position] != rune('\t') {
						goto l2083
					}
					position++
				}
			l2087:
			l2085:
				{
					position2086, tokenIndex2086 := position, tokenIndex
					{
						position2089, tokenIndex2089 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l2090
						}
						position++
						goto l2089
					l2090:
						position, tokenIndex = position2089, tokenIndex2089
						if buffer[position] != rune('\t') {
							goto l2086
						}
						position++
					}
				l2089:
					goto l2085
				l2086:
					position, tokenIndex = position2086, tokenIndex2086
				}
				add(rulews, position2084)
			}
			return true
		l2083:
			position, tokenIndex = position2083, tokenIndex2083
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
			position2113, tokenIndex2113 := position, tokenIndex
			{
				position2114 := position
				{
					position2115, tokenIndex2115 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2116
					}
					position++
					goto l2115
				l2116:
					position, tokenIndex = position2115, tokenIndex2115
					{
						position2117, tokenIndex2117 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2118
						}
						position++
						goto l2117
					l2118:
						position, tokenIndex = position2117, tokenIndex2117
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2113
						}
						position++
					}
				l2117:
				}
			l2115:
				add(rulehexdigit, position2114)
			}
			return true
		l2113:
			position, tokenIndex = position2113, tokenIndex2113
			return false
		},
		/* 160 octaldigit <- <(<[0-7]> Action105)> */
		func() bool {
			position2119, tokenIndex2119 := position, tokenIndex
			{
				position2120 := position
				{
					position2121 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2119
					}
					position++
					add(rulePegText, position2121)
				}
				{
					add(ruleAction105, position)
				}
				add(ruleoctaldigit, position2120)
			}
			return true
		l2119:
			position, tokenIndex = position2119, tokenIndex2119
			return false
		},
		/* 161 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2123, tokenIndex2123 := position, tokenIndex
			{
				position2124 := position
				{
					position2125, tokenIndex2125 := position, tokenIndex
					{
						position2127 := position
						{
							position2128, tokenIndex2128 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2129
							}
							position++
							goto l2128
						l2129:
							position, tokenIndex = position2128, tokenIndex2128
							if buffer[position] != rune('N') {
								goto l2126
							}
							position++
						}
					l2128:
						{
							position2130, tokenIndex2130 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2131
							}
							position++
							goto l2130
						l2131:
							position, tokenIndex = position2130, tokenIndex2130
							if buffer[position] != rune('Z') {
								goto l2126
							}
							position++
						}
					l2130:
						{
							add(ruleAction106, position)
						}
						add(ruleFT_NZ, position2127)
					}
					goto l2125
				l2126:
					position, tokenIndex = position2125, tokenIndex2125
					{
						position2134 := position
						{
							position2135, tokenIndex2135 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2136
							}
							position++
							goto l2135
						l2136:
							position, tokenIndex = position2135, tokenIndex2135
							if buffer[position] != rune('P') {
								goto l2133
							}
							position++
						}
					l2135:
						{
							position2137, tokenIndex2137 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2138
							}
							position++
							goto l2137
						l2138:
							position, tokenIndex = position2137, tokenIndex2137
							if buffer[position] != rune('O') {
								goto l2133
							}
							position++
						}
					l2137:
						{
							add(ruleAction110, position)
						}
						add(ruleFT_PO, position2134)
					}
					goto l2125
				l2133:
					position, tokenIndex = position2125, tokenIndex2125
					{
						position2141 := position
						{
							position2142, tokenIndex2142 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2143
							}
							position++
							goto l2142
						l2143:
							position, tokenIndex = position2142, tokenIndex2142
							if buffer[position] != rune('P') {
								goto l2140
							}
							position++
						}
					l2142:
						{
							position2144, tokenIndex2144 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2145
							}
							position++
							goto l2144
						l2145:
							position, tokenIndex = position2144, tokenIndex2144
							if buffer[position] != rune('E') {
								goto l2140
							}
							position++
						}
					l2144:
						{
							add(ruleAction111, position)
						}
						add(ruleFT_PE, position2141)
					}
					goto l2125
				l2140:
					position, tokenIndex = position2125, tokenIndex2125
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2148 := position
								{
									position2149, tokenIndex2149 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2150
									}
									position++
									goto l2149
								l2150:
									position, tokenIndex = position2149, tokenIndex2149
									if buffer[position] != rune('M') {
										goto l2123
									}
									position++
								}
							l2149:
								{
									add(ruleAction113, position)
								}
								add(ruleFT_M, position2148)
							}
							break
						case 'P', 'p':
							{
								position2152 := position
								{
									position2153, tokenIndex2153 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2154
									}
									position++
									goto l2153
								l2154:
									position, tokenIndex = position2153, tokenIndex2153
									if buffer[position] != rune('P') {
										goto l2123
									}
									position++
								}
							l2153:
								{
									add(ruleAction112, position)
								}
								add(ruleFT_P, position2152)
							}
							break
						case 'C', 'c':
							{
								position2156 := position
								{
									position2157, tokenIndex2157 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2158
									}
									position++
									goto l2157
								l2158:
									position, tokenIndex = position2157, tokenIndex2157
									if buffer[position] != rune('C') {
										goto l2123
									}
									position++
								}
							l2157:
								{
									add(ruleAction109, position)
								}
								add(ruleFT_C, position2156)
							}
							break
						case 'N', 'n':
							{
								position2160 := position
								{
									position2161, tokenIndex2161 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2162
									}
									position++
									goto l2161
								l2162:
									position, tokenIndex = position2161, tokenIndex2161
									if buffer[position] != rune('N') {
										goto l2123
									}
									position++
								}
							l2161:
								{
									position2163, tokenIndex2163 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2164
									}
									position++
									goto l2163
								l2164:
									position, tokenIndex = position2163, tokenIndex2163
									if buffer[position] != rune('C') {
										goto l2123
									}
									position++
								}
							l2163:
								{
									add(ruleAction108, position)
								}
								add(ruleFT_NC, position2160)
							}
							break
						default:
							{
								position2166 := position
								{
									position2167, tokenIndex2167 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2168
									}
									position++
									goto l2167
								l2168:
									position, tokenIndex = position2167, tokenIndex2167
									if buffer[position] != rune('Z') {
										goto l2123
									}
									position++
								}
							l2167:
								{
									add(ruleAction107, position)
								}
								add(ruleFT_Z, position2166)
							}
							break
						}
					}

				}
			l2125:
				add(rulecc, position2124)
			}
			return true
		l2123:
			position, tokenIndex = position2123, tokenIndex2123
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
