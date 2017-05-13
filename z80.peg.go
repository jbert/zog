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
	ruleDefb
	ruleDefs
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
	ruleAction3
	ruleAction4
	rulePegText
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
	ruleAction114
	ruleAction115
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
	"Defb",
	"Defs",
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
	"Action3",
	"Action4",
	"PegText",
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
	"Action114",
	"Action115",
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
	rules  [290]func() bool
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
			p.DefByte()
		case ruleAction3:
			p.DefSpace()
		case ruleAction4:
			p.LabelDefn(buffer[begin:end])
		case ruleAction5:
			p.LD8()
		case ruleAction6:
			p.LD16()
		case ruleAction7:
			p.Push()
		case ruleAction8:
			p.Pop()
		case ruleAction9:
			p.Ex()
		case ruleAction10:
			p.Inc8()
		case ruleAction11:
			p.Inc8()
		case ruleAction12:
			p.Inc16()
		case ruleAction13:
			p.Dec8()
		case ruleAction14:
			p.Dec8()
		case ruleAction15:
			p.Dec16()
		case ruleAction16:
			p.Add16()
		case ruleAction17:
			p.Adc16()
		case ruleAction18:
			p.Sbc16()
		case ruleAction19:
			p.Dst8()
		case ruleAction20:
			p.Src8()
		case ruleAction21:
			p.Loc8()
		case ruleAction22:
			p.Copy8()
		case ruleAction23:
			p.Loc8()
		case ruleAction24:
			p.R8(buffer[begin:end])
		case ruleAction25:
			p.R8(buffer[begin:end])
		case ruleAction26:
			p.Dst16()
		case ruleAction27:
			p.Src16()
		case ruleAction28:
			p.Loc16()
		case ruleAction29:
			p.R16(buffer[begin:end])
		case ruleAction30:
			p.R16(buffer[begin:end])
		case ruleAction31:
			p.R16Contents()
		case ruleAction32:
			p.IR16Contents()
		case ruleAction33:
			p.DispDecimal(buffer[begin:end])
		case ruleAction34:
			p.DispHex(buffer[begin:end])
		case ruleAction35:
			p.Disp0xHex(buffer[begin:end])
		case ruleAction36:
			p.Nhex(buffer[begin:end])
		case ruleAction37:
			p.Nhex(buffer[begin:end])
		case ruleAction38:
			p.Ndec(buffer[begin:end])
		case ruleAction39:
			p.NNLabel(buffer[begin:end])
		case ruleAction40:
			p.NNhex(buffer[begin:end])
		case ruleAction41:
			p.NNhex(buffer[begin:end])
		case ruleAction42:
			p.NNContents()
		case ruleAction43:
			p.Accum("ADD")
		case ruleAction44:
			p.Accum("ADC")
		case ruleAction45:
			p.Accum("SUB")
		case ruleAction46:
			p.Accum("SBC")
		case ruleAction47:
			p.Accum("AND")
		case ruleAction48:
			p.Accum("XOR")
		case ruleAction49:
			p.Accum("OR")
		case ruleAction50:
			p.Accum("CP")
		case ruleAction51:
			p.Rot("RLC")
		case ruleAction52:
			p.Rot("RRC")
		case ruleAction53:
			p.Rot("RL")
		case ruleAction54:
			p.Rot("RR")
		case ruleAction55:
			p.Rot("SLA")
		case ruleAction56:
			p.Rot("SRA")
		case ruleAction57:
			p.Rot("SLL")
		case ruleAction58:
			p.Rot("SRL")
		case ruleAction59:
			p.Bit()
		case ruleAction60:
			p.Res()
		case ruleAction61:
			p.Set()
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
			p.Simple(buffer[begin:end])
		case ruleAction74:
			p.Simple(buffer[begin:end])
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
			p.EDSimple(buffer[begin:end])
		case ruleAction98:
			p.EDSimple(buffer[begin:end])
		case ruleAction99:
			p.Rst()
		case ruleAction100:
			p.Call()
		case ruleAction101:
			p.Ret()
		case ruleAction102:
			p.Jp()
		case ruleAction103:
			p.Jr()
		case ruleAction104:
			p.Djnz()
		case ruleAction105:
			p.In()
		case ruleAction106:
			p.Out()
		case ruleAction107:
			p.ODigit(buffer[begin:end])
		case ruleAction108:
			p.Conditional(Not{FT_Z})
		case ruleAction109:
			p.Conditional(FT_Z)
		case ruleAction110:
			p.Conditional(Not{FT_C})
		case ruleAction111:
			p.Conditional(FT_C)
		case ruleAction112:
			p.Conditional(FT_PO)
		case ruleAction113:
			p.Conditional(FT_PE)
		case ruleAction114:
			p.Conditional(FT_P)
		case ruleAction115:
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
								add(ruleAction4, position)
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
										position19, tokenIndex19 := position, tokenIndex
										{
											position21 := position
											{
												position22, tokenIndex22 := position, tokenIndex
												{
													position24, tokenIndex24 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l25
													}
													position++
													goto l24
												l25:
													position, tokenIndex = position24, tokenIndex24
													if buffer[position] != rune('D') {
														goto l23
													}
													position++
												}
											l24:
												{
													position26, tokenIndex26 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l27
													}
													position++
													goto l26
												l27:
													position, tokenIndex = position26, tokenIndex26
													if buffer[position] != rune('E') {
														goto l23
													}
													position++
												}
											l26:
												{
													position28, tokenIndex28 := position, tokenIndex
													if buffer[position] != rune('f') {
														goto l29
													}
													position++
													goto l28
												l29:
													position, tokenIndex = position28, tokenIndex28
													if buffer[position] != rune('F') {
														goto l23
													}
													position++
												}
											l28:
												{
													position30, tokenIndex30 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l31
													}
													position++
													goto l30
												l31:
													position, tokenIndex = position30, tokenIndex30
													if buffer[position] != rune('B') {
														goto l23
													}
													position++
												}
											l30:
												goto l22
											l23:
												position, tokenIndex = position22, tokenIndex22
												{
													position33, tokenIndex33 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l34
													}
													position++
													goto l33
												l34:
													position, tokenIndex = position33, tokenIndex33
													if buffer[position] != rune('D') {
														goto l32
													}
													position++
												}
											l33:
												{
													position35, tokenIndex35 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l36
													}
													position++
													goto l35
												l36:
													position, tokenIndex = position35, tokenIndex35
													if buffer[position] != rune('B') {
														goto l32
													}
													position++
												}
											l35:
												goto l22
											l32:
												position, tokenIndex = position22, tokenIndex22
												{
													position37, tokenIndex37 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l38
													}
													position++
													goto l37
												l38:
													position, tokenIndex = position37, tokenIndex37
													if buffer[position] != rune('E') {
														goto l20
													}
													position++
												}
											l37:
												{
													position39, tokenIndex39 := position, tokenIndex
													if buffer[position] != rune('q') {
														goto l40
													}
													position++
													goto l39
												l40:
													position, tokenIndex = position39, tokenIndex39
													if buffer[position] != rune('Q') {
														goto l20
													}
													position++
												}
											l39:
												{
													position41, tokenIndex41 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l42
													}
													position++
													goto l41
												l42:
													position, tokenIndex = position41, tokenIndex41
													if buffer[position] != rune('U') {
														goto l20
													}
													position++
												}
											l41:
											}
										l22:
											if !_rules[rulews]() {
												goto l20
											}
											if !_rules[rulen]() {
												goto l20
											}
											{
												add(ruleAction2, position)
											}
											add(ruleDefb, position21)
										}
										goto l19
									l20:
										position, tokenIndex = position19, tokenIndex19
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position45 := position
													{
														position46, tokenIndex46 := position, tokenIndex
														{
															position48, tokenIndex48 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l49
															}
															position++
															goto l48
														l49:
															position, tokenIndex = position48, tokenIndex48
															if buffer[position] != rune('D') {
																goto l47
															}
															position++
														}
													l48:
														{
															position50, tokenIndex50 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l51
															}
															position++
															goto l50
														l51:
															position, tokenIndex = position50, tokenIndex50
															if buffer[position] != rune('E') {
																goto l47
															}
															position++
														}
													l50:
														{
															position52, tokenIndex52 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l53
															}
															position++
															goto l52
														l53:
															position, tokenIndex = position52, tokenIndex52
															if buffer[position] != rune('F') {
																goto l47
															}
															position++
														}
													l52:
														{
															position54, tokenIndex54 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l55
															}
															position++
															goto l54
														l55:
															position, tokenIndex = position54, tokenIndex54
															if buffer[position] != rune('S') {
																goto l47
															}
															position++
														}
													l54:
														goto l46
													l47:
														position, tokenIndex = position46, tokenIndex46
														{
															position56, tokenIndex56 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l57
															}
															position++
															goto l56
														l57:
															position, tokenIndex = position56, tokenIndex56
															if buffer[position] != rune('D') {
																goto l17
															}
															position++
														}
													l56:
														{
															position58, tokenIndex58 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l59
															}
															position++
															goto l58
														l59:
															position, tokenIndex = position58, tokenIndex58
															if buffer[position] != rune('S') {
																goto l17
															}
															position++
														}
													l58:
													}
												l46:
													if !_rules[rulews]() {
														goto l17
													}
													if !_rules[rulen]() {
														goto l17
													}
													{
														add(ruleAction3, position)
													}
													add(ruleDefs, position45)
												}
												break
											case 'O', 'o':
												{
													position61 := position
													{
														position62, tokenIndex62 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l63
														}
														position++
														goto l62
													l63:
														position, tokenIndex = position62, tokenIndex62
														if buffer[position] != rune('O') {
															goto l17
														}
														position++
													}
												l62:
													{
														position64, tokenIndex64 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l65
														}
														position++
														goto l64
													l65:
														position, tokenIndex = position64, tokenIndex64
														if buffer[position] != rune('R') {
															goto l17
														}
														position++
													}
												l64:
													{
														position66, tokenIndex66 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l67
														}
														position++
														goto l66
													l67:
														position, tokenIndex = position66, tokenIndex66
														if buffer[position] != rune('G') {
															goto l17
														}
														position++
													}
												l66:
													if !_rules[rulews]() {
														goto l17
													}
													if !_rules[rulenn]() {
														goto l17
													}
													{
														add(ruleAction1, position)
													}
													add(ruleOrg, position61)
												}
												break
											case 'a':
												{
													position69 := position
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
													add(ruleAseg, position69)
												}
												break
											default:
												{
													position70 := position
													{
														position71, tokenIndex71 := position, tokenIndex
														if buffer[position] != rune('.') {
															goto l71
														}
														position++
														goto l72
													l71:
														position, tokenIndex = position71, tokenIndex71
													}
												l72:
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
												l73:
													{
														position74, tokenIndex74 := position, tokenIndex
														{
															position75, tokenIndex75 := position, tokenIndex
															if buffer[position] != rune('\'') {
																goto l75
															}
															position++
															goto l74
														l75:
															position, tokenIndex = position75, tokenIndex75
														}
														if !matchDot() {
															goto l74
														}
														goto l73
													l74:
														position, tokenIndex = position74, tokenIndex74
													}
													if buffer[position] != rune('\'') {
														goto l17
													}
													position++
													add(ruleTitle, position70)
												}
												break
											}
										}

									}
								l19:
									add(ruleDirective, position18)
								}
								goto l16
							l17:
								position, tokenIndex = position16, tokenIndex16
								{
									position76 := position
									{
										position77, tokenIndex77 := position, tokenIndex
										{
											position79 := position
											{
												position80, tokenIndex80 := position, tokenIndex
												{
													position82 := position
													{
														position83, tokenIndex83 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l84
														}
														position++
														goto l83
													l84:
														position, tokenIndex = position83, tokenIndex83
														if buffer[position] != rune('P') {
															goto l81
														}
														position++
													}
												l83:
													{
														position85, tokenIndex85 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l86
														}
														position++
														goto l85
													l86:
														position, tokenIndex = position85, tokenIndex85
														if buffer[position] != rune('U') {
															goto l81
														}
														position++
													}
												l85:
													{
														position87, tokenIndex87 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l88
														}
														position++
														goto l87
													l88:
														position, tokenIndex = position87, tokenIndex87
														if buffer[position] != rune('S') {
															goto l81
														}
														position++
													}
												l87:
													{
														position89, tokenIndex89 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l90
														}
														position++
														goto l89
													l90:
														position, tokenIndex = position89, tokenIndex89
														if buffer[position] != rune('H') {
															goto l81
														}
														position++
													}
												l89:
													if !_rules[rulews]() {
														goto l81
													}
													if !_rules[ruleSrc16]() {
														goto l81
													}
													{
														add(ruleAction7, position)
													}
													add(rulePush, position82)
												}
												goto l80
											l81:
												position, tokenIndex = position80, tokenIndex80
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position93 := position
															{
																position94, tokenIndex94 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l95
																}
																position++
																goto l94
															l95:
																position, tokenIndex = position94, tokenIndex94
																if buffer[position] != rune('E') {
																	goto l78
																}
																position++
															}
														l94:
															{
																position96, tokenIndex96 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l97
																}
																position++
																goto l96
															l97:
																position, tokenIndex = position96, tokenIndex96
																if buffer[position] != rune('X') {
																	goto l78
																}
																position++
															}
														l96:
															if !_rules[rulews]() {
																goto l78
															}
															if !_rules[ruleDst16]() {
																goto l78
															}
															if !_rules[rulesep]() {
																goto l78
															}
															if !_rules[ruleSrc16]() {
																goto l78
															}
															{
																add(ruleAction9, position)
															}
															add(ruleEx, position93)
														}
														break
													case 'P', 'p':
														{
															position99 := position
															{
																position100, tokenIndex100 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l101
																}
																position++
																goto l100
															l101:
																position, tokenIndex = position100, tokenIndex100
																if buffer[position] != rune('P') {
																	goto l78
																}
																position++
															}
														l100:
															{
																position102, tokenIndex102 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l103
																}
																position++
																goto l102
															l103:
																position, tokenIndex = position102, tokenIndex102
																if buffer[position] != rune('O') {
																	goto l78
																}
																position++
															}
														l102:
															{
																position104, tokenIndex104 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l105
																}
																position++
																goto l104
															l105:
																position, tokenIndex = position104, tokenIndex104
																if buffer[position] != rune('P') {
																	goto l78
																}
																position++
															}
														l104:
															if !_rules[rulews]() {
																goto l78
															}
															if !_rules[ruleDst16]() {
																goto l78
															}
															{
																add(ruleAction8, position)
															}
															add(rulePop, position99)
														}
														break
													default:
														{
															position107 := position
															{
																position108, tokenIndex108 := position, tokenIndex
																{
																	position110 := position
																	{
																		position111, tokenIndex111 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l112
																		}
																		position++
																		goto l111
																	l112:
																		position, tokenIndex = position111, tokenIndex111
																		if buffer[position] != rune('L') {
																			goto l109
																		}
																		position++
																	}
																l111:
																	{
																		position113, tokenIndex113 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l114
																		}
																		position++
																		goto l113
																	l114:
																		position, tokenIndex = position113, tokenIndex113
																		if buffer[position] != rune('D') {
																			goto l109
																		}
																		position++
																	}
																l113:
																	if !_rules[rulews]() {
																		goto l109
																	}
																	if !_rules[ruleDst16]() {
																		goto l109
																	}
																	if !_rules[rulesep]() {
																		goto l109
																	}
																	if !_rules[ruleSrc16]() {
																		goto l109
																	}
																	{
																		add(ruleAction6, position)
																	}
																	add(ruleLoad16, position110)
																}
																goto l108
															l109:
																position, tokenIndex = position108, tokenIndex108
																{
																	position116 := position
																	{
																		position117, tokenIndex117 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l118
																		}
																		position++
																		goto l117
																	l118:
																		position, tokenIndex = position117, tokenIndex117
																		if buffer[position] != rune('L') {
																			goto l78
																		}
																		position++
																	}
																l117:
																	{
																		position119, tokenIndex119 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l120
																		}
																		position++
																		goto l119
																	l120:
																		position, tokenIndex = position119, tokenIndex119
																		if buffer[position] != rune('D') {
																			goto l78
																		}
																		position++
																	}
																l119:
																	if !_rules[rulews]() {
																		goto l78
																	}
																	{
																		position121 := position
																		{
																			position122, tokenIndex122 := position, tokenIndex
																			if !_rules[ruleReg8]() {
																				goto l123
																			}
																			goto l122
																		l123:
																			position, tokenIndex = position122, tokenIndex122
																			if !_rules[ruleReg16Contents]() {
																				goto l124
																			}
																			goto l122
																		l124:
																			position, tokenIndex = position122, tokenIndex122
																			if !_rules[rulenn_contents]() {
																				goto l78
																			}
																		}
																	l122:
																		{
																			add(ruleAction19, position)
																		}
																		add(ruleDst8, position121)
																	}
																	if !_rules[rulesep]() {
																		goto l78
																	}
																	if !_rules[ruleSrc8]() {
																		goto l78
																	}
																	{
																		add(ruleAction5, position)
																	}
																	add(ruleLoad8, position116)
																}
															}
														l108:
															add(ruleLoad, position107)
														}
														break
													}
												}

											}
										l80:
											add(ruleAssignment, position79)
										}
										goto l77
									l78:
										position, tokenIndex = position77, tokenIndex77
										{
											position128 := position
											{
												position129, tokenIndex129 := position, tokenIndex
												{
													position131 := position
													{
														position132, tokenIndex132 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l133
														}
														position++
														goto l132
													l133:
														position, tokenIndex = position132, tokenIndex132
														if buffer[position] != rune('I') {
															goto l130
														}
														position++
													}
												l132:
													{
														position134, tokenIndex134 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l135
														}
														position++
														goto l134
													l135:
														position, tokenIndex = position134, tokenIndex134
														if buffer[position] != rune('N') {
															goto l130
														}
														position++
													}
												l134:
													{
														position136, tokenIndex136 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l137
														}
														position++
														goto l136
													l137:
														position, tokenIndex = position136, tokenIndex136
														if buffer[position] != rune('C') {
															goto l130
														}
														position++
													}
												l136:
													if !_rules[rulews]() {
														goto l130
													}
													if !_rules[ruleILoc8]() {
														goto l130
													}
													{
														add(ruleAction10, position)
													}
													add(ruleInc16Indexed8, position131)
												}
												goto l129
											l130:
												position, tokenIndex = position129, tokenIndex129
												{
													position140 := position
													{
														position141, tokenIndex141 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l142
														}
														position++
														goto l141
													l142:
														position, tokenIndex = position141, tokenIndex141
														if buffer[position] != rune('I') {
															goto l139
														}
														position++
													}
												l141:
													{
														position143, tokenIndex143 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l144
														}
														position++
														goto l143
													l144:
														position, tokenIndex = position143, tokenIndex143
														if buffer[position] != rune('N') {
															goto l139
														}
														position++
													}
												l143:
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
															goto l139
														}
														position++
													}
												l145:
													if !_rules[rulews]() {
														goto l139
													}
													if !_rules[ruleLoc16]() {
														goto l139
													}
													{
														add(ruleAction12, position)
													}
													add(ruleInc16, position140)
												}
												goto l129
											l139:
												position, tokenIndex = position129, tokenIndex129
												{
													position148 := position
													{
														position149, tokenIndex149 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l150
														}
														position++
														goto l149
													l150:
														position, tokenIndex = position149, tokenIndex149
														if buffer[position] != rune('I') {
															goto l127
														}
														position++
													}
												l149:
													{
														position151, tokenIndex151 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l152
														}
														position++
														goto l151
													l152:
														position, tokenIndex = position151, tokenIndex151
														if buffer[position] != rune('N') {
															goto l127
														}
														position++
													}
												l151:
													{
														position153, tokenIndex153 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l154
														}
														position++
														goto l153
													l154:
														position, tokenIndex = position153, tokenIndex153
														if buffer[position] != rune('C') {
															goto l127
														}
														position++
													}
												l153:
													if !_rules[rulews]() {
														goto l127
													}
													if !_rules[ruleLoc8]() {
														goto l127
													}
													{
														add(ruleAction11, position)
													}
													add(ruleInc8, position148)
												}
											}
										l129:
											add(ruleInc, position128)
										}
										goto l77
									l127:
										position, tokenIndex = position77, tokenIndex77
										{
											position157 := position
											{
												position158, tokenIndex158 := position, tokenIndex
												{
													position160 := position
													{
														position161, tokenIndex161 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l162
														}
														position++
														goto l161
													l162:
														position, tokenIndex = position161, tokenIndex161
														if buffer[position] != rune('D') {
															goto l159
														}
														position++
													}
												l161:
													{
														position163, tokenIndex163 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l164
														}
														position++
														goto l163
													l164:
														position, tokenIndex = position163, tokenIndex163
														if buffer[position] != rune('E') {
															goto l159
														}
														position++
													}
												l163:
													{
														position165, tokenIndex165 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l166
														}
														position++
														goto l165
													l166:
														position, tokenIndex = position165, tokenIndex165
														if buffer[position] != rune('C') {
															goto l159
														}
														position++
													}
												l165:
													if !_rules[rulews]() {
														goto l159
													}
													if !_rules[ruleILoc8]() {
														goto l159
													}
													{
														add(ruleAction13, position)
													}
													add(ruleDec16Indexed8, position160)
												}
												goto l158
											l159:
												position, tokenIndex = position158, tokenIndex158
												{
													position169 := position
													{
														position170, tokenIndex170 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l171
														}
														position++
														goto l170
													l171:
														position, tokenIndex = position170, tokenIndex170
														if buffer[position] != rune('D') {
															goto l168
														}
														position++
													}
												l170:
													{
														position172, tokenIndex172 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l173
														}
														position++
														goto l172
													l173:
														position, tokenIndex = position172, tokenIndex172
														if buffer[position] != rune('E') {
															goto l168
														}
														position++
													}
												l172:
													{
														position174, tokenIndex174 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l175
														}
														position++
														goto l174
													l175:
														position, tokenIndex = position174, tokenIndex174
														if buffer[position] != rune('C') {
															goto l168
														}
														position++
													}
												l174:
													if !_rules[rulews]() {
														goto l168
													}
													if !_rules[ruleLoc16]() {
														goto l168
													}
													{
														add(ruleAction15, position)
													}
													add(ruleDec16, position169)
												}
												goto l158
											l168:
												position, tokenIndex = position158, tokenIndex158
												{
													position177 := position
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
															goto l156
														}
														position++
													}
												l178:
													{
														position180, tokenIndex180 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l181
														}
														position++
														goto l180
													l181:
														position, tokenIndex = position180, tokenIndex180
														if buffer[position] != rune('E') {
															goto l156
														}
														position++
													}
												l180:
													{
														position182, tokenIndex182 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l183
														}
														position++
														goto l182
													l183:
														position, tokenIndex = position182, tokenIndex182
														if buffer[position] != rune('C') {
															goto l156
														}
														position++
													}
												l182:
													if !_rules[rulews]() {
														goto l156
													}
													if !_rules[ruleLoc8]() {
														goto l156
													}
													{
														add(ruleAction14, position)
													}
													add(ruleDec8, position177)
												}
											}
										l158:
											add(ruleDec, position157)
										}
										goto l77
									l156:
										position, tokenIndex = position77, tokenIndex77
										{
											position186 := position
											{
												position187, tokenIndex187 := position, tokenIndex
												{
													position189 := position
													{
														position190, tokenIndex190 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l191
														}
														position++
														goto l190
													l191:
														position, tokenIndex = position190, tokenIndex190
														if buffer[position] != rune('A') {
															goto l188
														}
														position++
													}
												l190:
													{
														position192, tokenIndex192 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l193
														}
														position++
														goto l192
													l193:
														position, tokenIndex = position192, tokenIndex192
														if buffer[position] != rune('D') {
															goto l188
														}
														position++
													}
												l192:
													{
														position194, tokenIndex194 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l195
														}
														position++
														goto l194
													l195:
														position, tokenIndex = position194, tokenIndex194
														if buffer[position] != rune('D') {
															goto l188
														}
														position++
													}
												l194:
													if !_rules[rulews]() {
														goto l188
													}
													if !_rules[ruleDst16]() {
														goto l188
													}
													if !_rules[rulesep]() {
														goto l188
													}
													if !_rules[ruleSrc16]() {
														goto l188
													}
													{
														add(ruleAction16, position)
													}
													add(ruleAdd16, position189)
												}
												goto l187
											l188:
												position, tokenIndex = position187, tokenIndex187
												{
													position198 := position
													{
														position199, tokenIndex199 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l200
														}
														position++
														goto l199
													l200:
														position, tokenIndex = position199, tokenIndex199
														if buffer[position] != rune('A') {
															goto l197
														}
														position++
													}
												l199:
													{
														position201, tokenIndex201 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l202
														}
														position++
														goto l201
													l202:
														position, tokenIndex = position201, tokenIndex201
														if buffer[position] != rune('D') {
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
													if !_rules[ruleDst16]() {
														goto l197
													}
													if !_rules[rulesep]() {
														goto l197
													}
													if !_rules[ruleSrc16]() {
														goto l197
													}
													{
														add(ruleAction17, position)
													}
													add(ruleAdc16, position198)
												}
												goto l187
											l197:
												position, tokenIndex = position187, tokenIndex187
												{
													position206 := position
													{
														position207, tokenIndex207 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l208
														}
														position++
														goto l207
													l208:
														position, tokenIndex = position207, tokenIndex207
														if buffer[position] != rune('S') {
															goto l185
														}
														position++
													}
												l207:
													{
														position209, tokenIndex209 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l210
														}
														position++
														goto l209
													l210:
														position, tokenIndex = position209, tokenIndex209
														if buffer[position] != rune('B') {
															goto l185
														}
														position++
													}
												l209:
													{
														position211, tokenIndex211 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l212
														}
														position++
														goto l211
													l212:
														position, tokenIndex = position211, tokenIndex211
														if buffer[position] != rune('C') {
															goto l185
														}
														position++
													}
												l211:
													if !_rules[rulews]() {
														goto l185
													}
													if !_rules[ruleDst16]() {
														goto l185
													}
													if !_rules[rulesep]() {
														goto l185
													}
													if !_rules[ruleSrc16]() {
														goto l185
													}
													{
														add(ruleAction18, position)
													}
													add(ruleSbc16, position206)
												}
											}
										l187:
											add(ruleAlu16, position186)
										}
										goto l77
									l185:
										position, tokenIndex = position77, tokenIndex77
										{
											position215 := position
											{
												position216, tokenIndex216 := position, tokenIndex
												{
													position218 := position
													{
														position219, tokenIndex219 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l220
														}
														position++
														goto l219
													l220:
														position, tokenIndex = position219, tokenIndex219
														if buffer[position] != rune('A') {
															goto l217
														}
														position++
													}
												l219:
													{
														position221, tokenIndex221 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l222
														}
														position++
														goto l221
													l222:
														position, tokenIndex = position221, tokenIndex221
														if buffer[position] != rune('D') {
															goto l217
														}
														position++
													}
												l221:
													{
														position223, tokenIndex223 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l224
														}
														position++
														goto l223
													l224:
														position, tokenIndex = position223, tokenIndex223
														if buffer[position] != rune('D') {
															goto l217
														}
														position++
													}
												l223:
													if !_rules[rulews]() {
														goto l217
													}
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
															goto l217
														}
														position++
													}
												l225:
													if !_rules[rulesep]() {
														goto l217
													}
													if !_rules[ruleSrc8]() {
														goto l217
													}
													{
														add(ruleAction43, position)
													}
													add(ruleAdd, position218)
												}
												goto l216
											l217:
												position, tokenIndex = position216, tokenIndex216
												{
													position229 := position
													{
														position230, tokenIndex230 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l231
														}
														position++
														goto l230
													l231:
														position, tokenIndex = position230, tokenIndex230
														if buffer[position] != rune('A') {
															goto l228
														}
														position++
													}
												l230:
													{
														position232, tokenIndex232 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l233
														}
														position++
														goto l232
													l233:
														position, tokenIndex = position232, tokenIndex232
														if buffer[position] != rune('D') {
															goto l228
														}
														position++
													}
												l232:
													{
														position234, tokenIndex234 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l235
														}
														position++
														goto l234
													l235:
														position, tokenIndex = position234, tokenIndex234
														if buffer[position] != rune('C') {
															goto l228
														}
														position++
													}
												l234:
													if !_rules[rulews]() {
														goto l228
													}
													{
														position236, tokenIndex236 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l237
														}
														position++
														goto l236
													l237:
														position, tokenIndex = position236, tokenIndex236
														if buffer[position] != rune('A') {
															goto l228
														}
														position++
													}
												l236:
													if !_rules[rulesep]() {
														goto l228
													}
													if !_rules[ruleSrc8]() {
														goto l228
													}
													{
														add(ruleAction44, position)
													}
													add(ruleAdc, position229)
												}
												goto l216
											l228:
												position, tokenIndex = position216, tokenIndex216
												{
													position240 := position
													{
														position241, tokenIndex241 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l242
														}
														position++
														goto l241
													l242:
														position, tokenIndex = position241, tokenIndex241
														if buffer[position] != rune('S') {
															goto l239
														}
														position++
													}
												l241:
													{
														position243, tokenIndex243 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l244
														}
														position++
														goto l243
													l244:
														position, tokenIndex = position243, tokenIndex243
														if buffer[position] != rune('U') {
															goto l239
														}
														position++
													}
												l243:
													{
														position245, tokenIndex245 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l246
														}
														position++
														goto l245
													l246:
														position, tokenIndex = position245, tokenIndex245
														if buffer[position] != rune('B') {
															goto l239
														}
														position++
													}
												l245:
													if !_rules[rulews]() {
														goto l239
													}
													if !_rules[ruleSrc8]() {
														goto l239
													}
													{
														add(ruleAction45, position)
													}
													add(ruleSub, position240)
												}
												goto l216
											l239:
												position, tokenIndex = position216, tokenIndex216
												{
													switch buffer[position] {
													case 'C', 'c':
														{
															position249 := position
															{
																position250, tokenIndex250 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l251
																}
																position++
																goto l250
															l251:
																position, tokenIndex = position250, tokenIndex250
																if buffer[position] != rune('C') {
																	goto l214
																}
																position++
															}
														l250:
															{
																position252, tokenIndex252 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l253
																}
																position++
																goto l252
															l253:
																position, tokenIndex = position252, tokenIndex252
																if buffer[position] != rune('P') {
																	goto l214
																}
																position++
															}
														l252:
															if !_rules[rulews]() {
																goto l214
															}
															if !_rules[ruleSrc8]() {
																goto l214
															}
															{
																add(ruleAction50, position)
															}
															add(ruleCp, position249)
														}
														break
													case 'O', 'o':
														{
															position255 := position
															{
																position256, tokenIndex256 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l257
																}
																position++
																goto l256
															l257:
																position, tokenIndex = position256, tokenIndex256
																if buffer[position] != rune('O') {
																	goto l214
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
																	goto l214
																}
																position++
															}
														l258:
															if !_rules[rulews]() {
																goto l214
															}
															if !_rules[ruleSrc8]() {
																goto l214
															}
															{
																add(ruleAction49, position)
															}
															add(ruleOr, position255)
														}
														break
													case 'X', 'x':
														{
															position261 := position
															{
																position262, tokenIndex262 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l263
																}
																position++
																goto l262
															l263:
																position, tokenIndex = position262, tokenIndex262
																if buffer[position] != rune('X') {
																	goto l214
																}
																position++
															}
														l262:
															{
																position264, tokenIndex264 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l265
																}
																position++
																goto l264
															l265:
																position, tokenIndex = position264, tokenIndex264
																if buffer[position] != rune('O') {
																	goto l214
																}
																position++
															}
														l264:
															{
																position266, tokenIndex266 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l267
																}
																position++
																goto l266
															l267:
																position, tokenIndex = position266, tokenIndex266
																if buffer[position] != rune('R') {
																	goto l214
																}
																position++
															}
														l266:
															if !_rules[rulews]() {
																goto l214
															}
															if !_rules[ruleSrc8]() {
																goto l214
															}
															{
																add(ruleAction48, position)
															}
															add(ruleXor, position261)
														}
														break
													case 'A', 'a':
														{
															position269 := position
															{
																position270, tokenIndex270 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l271
																}
																position++
																goto l270
															l271:
																position, tokenIndex = position270, tokenIndex270
																if buffer[position] != rune('A') {
																	goto l214
																}
																position++
															}
														l270:
															{
																position272, tokenIndex272 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l273
																}
																position++
																goto l272
															l273:
																position, tokenIndex = position272, tokenIndex272
																if buffer[position] != rune('N') {
																	goto l214
																}
																position++
															}
														l272:
															{
																position274, tokenIndex274 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l275
																}
																position++
																goto l274
															l275:
																position, tokenIndex = position274, tokenIndex274
																if buffer[position] != rune('D') {
																	goto l214
																}
																position++
															}
														l274:
															if !_rules[rulews]() {
																goto l214
															}
															if !_rules[ruleSrc8]() {
																goto l214
															}
															{
																add(ruleAction47, position)
															}
															add(ruleAnd, position269)
														}
														break
													default:
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
																	goto l214
																}
																position++
															}
														l278:
															{
																position280, tokenIndex280 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l281
																}
																position++
																goto l280
															l281:
																position, tokenIndex = position280, tokenIndex280
																if buffer[position] != rune('B') {
																	goto l214
																}
																position++
															}
														l280:
															{
																position282, tokenIndex282 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l283
																}
																position++
																goto l282
															l283:
																position, tokenIndex = position282, tokenIndex282
																if buffer[position] != rune('C') {
																	goto l214
																}
																position++
															}
														l282:
															if !_rules[rulews]() {
																goto l214
															}
															{
																position284, tokenIndex284 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l285
																}
																position++
																goto l284
															l285:
																position, tokenIndex = position284, tokenIndex284
																if buffer[position] != rune('A') {
																	goto l214
																}
																position++
															}
														l284:
															if !_rules[rulesep]() {
																goto l214
															}
															if !_rules[ruleSrc8]() {
																goto l214
															}
															{
																add(ruleAction46, position)
															}
															add(ruleSbc, position277)
														}
														break
													}
												}

											}
										l216:
											add(ruleAlu, position215)
										}
										goto l77
									l214:
										position, tokenIndex = position77, tokenIndex77
										{
											position288 := position
											{
												position289, tokenIndex289 := position, tokenIndex
												{
													position291 := position
													{
														position292, tokenIndex292 := position, tokenIndex
														{
															position294 := position
															{
																position295, tokenIndex295 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l296
																}
																position++
																goto l295
															l296:
																position, tokenIndex = position295, tokenIndex295
																if buffer[position] != rune('R') {
																	goto l293
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
																	goto l293
																}
																position++
															}
														l297:
															{
																position299, tokenIndex299 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l300
																}
																position++
																goto l299
															l300:
																position, tokenIndex = position299, tokenIndex299
																if buffer[position] != rune('C') {
																	goto l293
																}
																position++
															}
														l299:
															if !_rules[rulews]() {
																goto l293
															}
															if !_rules[ruleLoc8]() {
																goto l293
															}
															{
																position301, tokenIndex301 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l301
																}
																if !_rules[ruleCopy8]() {
																	goto l301
																}
																goto l302
															l301:
																position, tokenIndex = position301, tokenIndex301
															}
														l302:
															{
																add(ruleAction51, position)
															}
															add(ruleRlc, position294)
														}
														goto l292
													l293:
														position, tokenIndex = position292, tokenIndex292
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
																	goto l304
																}
																position++
															}
														l306:
															{
																position308, tokenIndex308 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l309
																}
																position++
																goto l308
															l309:
																position, tokenIndex = position308, tokenIndex308
																if buffer[position] != rune('R') {
																	goto l304
																}
																position++
															}
														l308:
															{
																position310, tokenIndex310 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l311
																}
																position++
																goto l310
															l311:
																position, tokenIndex = position310, tokenIndex310
																if buffer[position] != rune('C') {
																	goto l304
																}
																position++
															}
														l310:
															if !_rules[rulews]() {
																goto l304
															}
															if !_rules[ruleLoc8]() {
																goto l304
															}
															{
																position312, tokenIndex312 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l312
																}
																if !_rules[ruleCopy8]() {
																	goto l312
																}
																goto l313
															l312:
																position, tokenIndex = position312, tokenIndex312
															}
														l313:
															{
																add(ruleAction52, position)
															}
															add(ruleRrc, position305)
														}
														goto l292
													l304:
														position, tokenIndex = position292, tokenIndex292
														{
															position316 := position
															{
																position317, tokenIndex317 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l318
																}
																position++
																goto l317
															l318:
																position, tokenIndex = position317, tokenIndex317
																if buffer[position] != rune('R') {
																	goto l315
																}
																position++
															}
														l317:
															{
																position319, tokenIndex319 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l320
																}
																position++
																goto l319
															l320:
																position, tokenIndex = position319, tokenIndex319
																if buffer[position] != rune('L') {
																	goto l315
																}
																position++
															}
														l319:
															if !_rules[rulews]() {
																goto l315
															}
															if !_rules[ruleLoc8]() {
																goto l315
															}
															{
																position321, tokenIndex321 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l321
																}
																if !_rules[ruleCopy8]() {
																	goto l321
																}
																goto l322
															l321:
																position, tokenIndex = position321, tokenIndex321
															}
														l322:
															{
																add(ruleAction53, position)
															}
															add(ruleRl, position316)
														}
														goto l292
													l315:
														position, tokenIndex = position292, tokenIndex292
														{
															position325 := position
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
																	goto l324
																}
																position++
															}
														l326:
															{
																position328, tokenIndex328 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l329
																}
																position++
																goto l328
															l329:
																position, tokenIndex = position328, tokenIndex328
																if buffer[position] != rune('R') {
																	goto l324
																}
																position++
															}
														l328:
															if !_rules[rulews]() {
																goto l324
															}
															if !_rules[ruleLoc8]() {
																goto l324
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
																add(ruleAction54, position)
															}
															add(ruleRr, position325)
														}
														goto l292
													l324:
														position, tokenIndex = position292, tokenIndex292
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
																	goto l333
																}
																position++
															}
														l335:
															{
																position337, tokenIndex337 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l338
																}
																position++
																goto l337
															l338:
																position, tokenIndex = position337, tokenIndex337
																if buffer[position] != rune('L') {
																	goto l333
																}
																position++
															}
														l337:
															{
																position339, tokenIndex339 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l340
																}
																position++
																goto l339
															l340:
																position, tokenIndex = position339, tokenIndex339
																if buffer[position] != rune('A') {
																	goto l333
																}
																position++
															}
														l339:
															if !_rules[rulews]() {
																goto l333
															}
															if !_rules[ruleLoc8]() {
																goto l333
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
																add(ruleAction55, position)
															}
															add(ruleSla, position334)
														}
														goto l292
													l333:
														position, tokenIndex = position292, tokenIndex292
														{
															position345 := position
															{
																position346, tokenIndex346 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l347
																}
																position++
																goto l346
															l347:
																position, tokenIndex = position346, tokenIndex346
																if buffer[position] != rune('S') {
																	goto l344
																}
																position++
															}
														l346:
															{
																position348, tokenIndex348 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l349
																}
																position++
																goto l348
															l349:
																position, tokenIndex = position348, tokenIndex348
																if buffer[position] != rune('R') {
																	goto l344
																}
																position++
															}
														l348:
															{
																position350, tokenIndex350 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l351
																}
																position++
																goto l350
															l351:
																position, tokenIndex = position350, tokenIndex350
																if buffer[position] != rune('A') {
																	goto l344
																}
																position++
															}
														l350:
															if !_rules[rulews]() {
																goto l344
															}
															if !_rules[ruleLoc8]() {
																goto l344
															}
															{
																position352, tokenIndex352 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l352
																}
																if !_rules[ruleCopy8]() {
																	goto l352
																}
																goto l353
															l352:
																position, tokenIndex = position352, tokenIndex352
															}
														l353:
															{
																add(ruleAction56, position)
															}
															add(ruleSra, position345)
														}
														goto l292
													l344:
														position, tokenIndex = position292, tokenIndex292
														{
															position356 := position
															{
																position357, tokenIndex357 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l358
																}
																position++
																goto l357
															l358:
																position, tokenIndex = position357, tokenIndex357
																if buffer[position] != rune('S') {
																	goto l355
																}
																position++
															}
														l357:
															{
																position359, tokenIndex359 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l360
																}
																position++
																goto l359
															l360:
																position, tokenIndex = position359, tokenIndex359
																if buffer[position] != rune('L') {
																	goto l355
																}
																position++
															}
														l359:
															{
																position361, tokenIndex361 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l362
																}
																position++
																goto l361
															l362:
																position, tokenIndex = position361, tokenIndex361
																if buffer[position] != rune('L') {
																	goto l355
																}
																position++
															}
														l361:
															if !_rules[rulews]() {
																goto l355
															}
															if !_rules[ruleLoc8]() {
																goto l355
															}
															{
																position363, tokenIndex363 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l363
																}
																if !_rules[ruleCopy8]() {
																	goto l363
																}
																goto l364
															l363:
																position, tokenIndex = position363, tokenIndex363
															}
														l364:
															{
																add(ruleAction57, position)
															}
															add(ruleSll, position356)
														}
														goto l292
													l355:
														position, tokenIndex = position292, tokenIndex292
														{
															position366 := position
															{
																position367, tokenIndex367 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l368
																}
																position++
																goto l367
															l368:
																position, tokenIndex = position367, tokenIndex367
																if buffer[position] != rune('S') {
																	goto l290
																}
																position++
															}
														l367:
															{
																position369, tokenIndex369 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l370
																}
																position++
																goto l369
															l370:
																position, tokenIndex = position369, tokenIndex369
																if buffer[position] != rune('R') {
																	goto l290
																}
																position++
															}
														l369:
															{
																position371, tokenIndex371 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l372
																}
																position++
																goto l371
															l372:
																position, tokenIndex = position371, tokenIndex371
																if buffer[position] != rune('L') {
																	goto l290
																}
																position++
															}
														l371:
															if !_rules[rulews]() {
																goto l290
															}
															if !_rules[ruleLoc8]() {
																goto l290
															}
															{
																position373, tokenIndex373 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l373
																}
																if !_rules[ruleCopy8]() {
																	goto l373
																}
																goto l374
															l373:
																position, tokenIndex = position373, tokenIndex373
															}
														l374:
															{
																add(ruleAction58, position)
															}
															add(ruleSrl, position366)
														}
													}
												l292:
													add(ruleRot, position291)
												}
												goto l289
											l290:
												position, tokenIndex = position289, tokenIndex289
												{
													switch buffer[position] {
													case 'S', 's':
														{
															position377 := position
															{
																position378, tokenIndex378 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l379
																}
																position++
																goto l378
															l379:
																position, tokenIndex = position378, tokenIndex378
																if buffer[position] != rune('S') {
																	goto l287
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
																	goto l287
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
																	goto l287
																}
																position++
															}
														l382:
															if !_rules[rulews]() {
																goto l287
															}
															if !_rules[ruleoctaldigit]() {
																goto l287
															}
															if !_rules[rulesep]() {
																goto l287
															}
															if !_rules[ruleLoc8]() {
																goto l287
															}
															{
																position384, tokenIndex384 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l384
																}
																if !_rules[ruleCopy8]() {
																	goto l384
																}
																goto l385
															l384:
																position, tokenIndex = position384, tokenIndex384
															}
														l385:
															{
																add(ruleAction61, position)
															}
															add(ruleSet, position377)
														}
														break
													case 'R', 'r':
														{
															position387 := position
															{
																position388, tokenIndex388 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l389
																}
																position++
																goto l388
															l389:
																position, tokenIndex = position388, tokenIndex388
																if buffer[position] != rune('R') {
																	goto l287
																}
																position++
															}
														l388:
															{
																position390, tokenIndex390 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l391
																}
																position++
																goto l390
															l391:
																position, tokenIndex = position390, tokenIndex390
																if buffer[position] != rune('E') {
																	goto l287
																}
																position++
															}
														l390:
															{
																position392, tokenIndex392 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l393
																}
																position++
																goto l392
															l393:
																position, tokenIndex = position392, tokenIndex392
																if buffer[position] != rune('S') {
																	goto l287
																}
																position++
															}
														l392:
															if !_rules[rulews]() {
																goto l287
															}
															if !_rules[ruleoctaldigit]() {
																goto l287
															}
															if !_rules[rulesep]() {
																goto l287
															}
															if !_rules[ruleLoc8]() {
																goto l287
															}
															{
																position394, tokenIndex394 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l394
																}
																if !_rules[ruleCopy8]() {
																	goto l394
																}
																goto l395
															l394:
																position, tokenIndex = position394, tokenIndex394
															}
														l395:
															{
																add(ruleAction60, position)
															}
															add(ruleRes, position387)
														}
														break
													default:
														{
															position397 := position
															{
																position398, tokenIndex398 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l399
																}
																position++
																goto l398
															l399:
																position, tokenIndex = position398, tokenIndex398
																if buffer[position] != rune('B') {
																	goto l287
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
																	goto l287
																}
																position++
															}
														l400:
															{
																position402, tokenIndex402 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l403
																}
																position++
																goto l402
															l403:
																position, tokenIndex = position402, tokenIndex402
																if buffer[position] != rune('T') {
																	goto l287
																}
																position++
															}
														l402:
															if !_rules[rulews]() {
																goto l287
															}
															if !_rules[ruleoctaldigit]() {
																goto l287
															}
															if !_rules[rulesep]() {
																goto l287
															}
															if !_rules[ruleLoc8]() {
																goto l287
															}
															{
																add(ruleAction59, position)
															}
															add(ruleBit, position397)
														}
														break
													}
												}

											}
										l289:
											add(ruleBitOp, position288)
										}
										goto l77
									l287:
										position, tokenIndex = position77, tokenIndex77
										{
											position406 := position
											{
												position407, tokenIndex407 := position, tokenIndex
												{
													position409 := position
													{
														position410 := position
														{
															position411, tokenIndex411 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l412
															}
															position++
															goto l411
														l412:
															position, tokenIndex = position411, tokenIndex411
															if buffer[position] != rune('R') {
																goto l408
															}
															position++
														}
													l411:
														{
															position413, tokenIndex413 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l414
															}
															position++
															goto l413
														l414:
															position, tokenIndex = position413, tokenIndex413
															if buffer[position] != rune('E') {
																goto l408
															}
															position++
														}
													l413:
														{
															position415, tokenIndex415 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l416
															}
															position++
															goto l415
														l416:
															position, tokenIndex = position415, tokenIndex415
															if buffer[position] != rune('T') {
																goto l408
															}
															position++
														}
													l415:
														{
															position417, tokenIndex417 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l418
															}
															position++
															goto l417
														l418:
															position, tokenIndex = position417, tokenIndex417
															if buffer[position] != rune('N') {
																goto l408
															}
															position++
														}
													l417:
														add(rulePegText, position410)
													}
													{
														add(ruleAction76, position)
													}
													add(ruleRetn, position409)
												}
												goto l407
											l408:
												position, tokenIndex = position407, tokenIndex407
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
																goto l420
															}
															position++
														}
													l423:
														{
															position425, tokenIndex425 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l426
															}
															position++
															goto l425
														l426:
															position, tokenIndex = position425, tokenIndex425
															if buffer[position] != rune('E') {
																goto l420
															}
															position++
														}
													l425:
														{
															position427, tokenIndex427 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l428
															}
															position++
															goto l427
														l428:
															position, tokenIndex = position427, tokenIndex427
															if buffer[position] != rune('T') {
																goto l420
															}
															position++
														}
													l427:
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
																goto l420
															}
															position++
														}
													l429:
														add(rulePegText, position422)
													}
													{
														add(ruleAction77, position)
													}
													add(ruleReti, position421)
												}
												goto l407
											l420:
												position, tokenIndex = position407, tokenIndex407
												{
													position433 := position
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
																goto l432
															}
															position++
														}
													l435:
														{
															position437, tokenIndex437 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l438
															}
															position++
															goto l437
														l438:
															position, tokenIndex = position437, tokenIndex437
															if buffer[position] != rune('R') {
																goto l432
															}
															position++
														}
													l437:
														{
															position439, tokenIndex439 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l440
															}
															position++
															goto l439
														l440:
															position, tokenIndex = position439, tokenIndex439
															if buffer[position] != rune('D') {
																goto l432
															}
															position++
														}
													l439:
														add(rulePegText, position434)
													}
													{
														add(ruleAction78, position)
													}
													add(ruleRrd, position433)
												}
												goto l407
											l432:
												position, tokenIndex = position407, tokenIndex407
												{
													position443 := position
													{
														position444 := position
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
																goto l442
															}
															position++
														}
													l445:
														{
															position447, tokenIndex447 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l448
															}
															position++
															goto l447
														l448:
															position, tokenIndex = position447, tokenIndex447
															if buffer[position] != rune('M') {
																goto l442
															}
															position++
														}
													l447:
														if buffer[position] != rune(' ') {
															goto l442
														}
														position++
														if buffer[position] != rune('0') {
															goto l442
														}
														position++
														add(rulePegText, position444)
													}
													{
														add(ruleAction80, position)
													}
													add(ruleIm0, position443)
												}
												goto l407
											l442:
												position, tokenIndex = position407, tokenIndex407
												{
													position451 := position
													{
														position452 := position
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
																goto l450
															}
															position++
														}
													l453:
														{
															position455, tokenIndex455 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l456
															}
															position++
															goto l455
														l456:
															position, tokenIndex = position455, tokenIndex455
															if buffer[position] != rune('M') {
																goto l450
															}
															position++
														}
													l455:
														if buffer[position] != rune(' ') {
															goto l450
														}
														position++
														if buffer[position] != rune('1') {
															goto l450
														}
														position++
														add(rulePegText, position452)
													}
													{
														add(ruleAction81, position)
													}
													add(ruleIm1, position451)
												}
												goto l407
											l450:
												position, tokenIndex = position407, tokenIndex407
												{
													position459 := position
													{
														position460 := position
														{
															position461, tokenIndex461 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l462
															}
															position++
															goto l461
														l462:
															position, tokenIndex = position461, tokenIndex461
															if buffer[position] != rune('I') {
																goto l458
															}
															position++
														}
													l461:
														{
															position463, tokenIndex463 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l464
															}
															position++
															goto l463
														l464:
															position, tokenIndex = position463, tokenIndex463
															if buffer[position] != rune('M') {
																goto l458
															}
															position++
														}
													l463:
														if buffer[position] != rune(' ') {
															goto l458
														}
														position++
														if buffer[position] != rune('2') {
															goto l458
														}
														position++
														add(rulePegText, position460)
													}
													{
														add(ruleAction82, position)
													}
													add(ruleIm2, position459)
												}
												goto l407
											l458:
												position, tokenIndex = position407, tokenIndex407
												{
													switch buffer[position] {
													case 'I', 'O', 'i', 'o':
														{
															position467 := position
															{
																position468, tokenIndex468 := position, tokenIndex
																{
																	position470 := position
																	{
																		position471 := position
																		{
																			position472, tokenIndex472 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l473
																			}
																			position++
																			goto l472
																		l473:
																			position, tokenIndex = position472, tokenIndex472
																			if buffer[position] != rune('I') {
																				goto l469
																			}
																			position++
																		}
																	l472:
																		{
																			position474, tokenIndex474 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l475
																			}
																			position++
																			goto l474
																		l475:
																			position, tokenIndex = position474, tokenIndex474
																			if buffer[position] != rune('N') {
																				goto l469
																			}
																			position++
																		}
																	l474:
																		{
																			position476, tokenIndex476 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l477
																			}
																			position++
																			goto l476
																		l477:
																			position, tokenIndex = position476, tokenIndex476
																			if buffer[position] != rune('I') {
																				goto l469
																			}
																			position++
																		}
																	l476:
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
																				goto l469
																			}
																			position++
																		}
																	l478:
																		add(rulePegText, position471)
																	}
																	{
																		add(ruleAction93, position)
																	}
																	add(ruleInir, position470)
																}
																goto l468
															l469:
																position, tokenIndex = position468, tokenIndex468
																{
																	position482 := position
																	{
																		position483 := position
																		{
																			position484, tokenIndex484 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l485
																			}
																			position++
																			goto l484
																		l485:
																			position, tokenIndex = position484, tokenIndex484
																			if buffer[position] != rune('I') {
																				goto l481
																			}
																			position++
																		}
																	l484:
																		{
																			position486, tokenIndex486 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l487
																			}
																			position++
																			goto l486
																		l487:
																			position, tokenIndex = position486, tokenIndex486
																			if buffer[position] != rune('N') {
																				goto l481
																			}
																			position++
																		}
																	l486:
																		{
																			position488, tokenIndex488 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l489
																			}
																			position++
																			goto l488
																		l489:
																			position, tokenIndex = position488, tokenIndex488
																			if buffer[position] != rune('I') {
																				goto l481
																			}
																			position++
																		}
																	l488:
																		add(rulePegText, position483)
																	}
																	{
																		add(ruleAction85, position)
																	}
																	add(ruleIni, position482)
																}
																goto l468
															l481:
																position, tokenIndex = position468, tokenIndex468
																{
																	position492 := position
																	{
																		position493 := position
																		{
																			position494, tokenIndex494 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l495
																			}
																			position++
																			goto l494
																		l495:
																			position, tokenIndex = position494, tokenIndex494
																			if buffer[position] != rune('O') {
																				goto l491
																			}
																			position++
																		}
																	l494:
																		{
																			position496, tokenIndex496 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l497
																			}
																			position++
																			goto l496
																		l497:
																			position, tokenIndex = position496, tokenIndex496
																			if buffer[position] != rune('T') {
																				goto l491
																			}
																			position++
																		}
																	l496:
																		{
																			position498, tokenIndex498 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l499
																			}
																			position++
																			goto l498
																		l499:
																			position, tokenIndex = position498, tokenIndex498
																			if buffer[position] != rune('I') {
																				goto l491
																			}
																			position++
																		}
																	l498:
																		{
																			position500, tokenIndex500 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l501
																			}
																			position++
																			goto l500
																		l501:
																			position, tokenIndex = position500, tokenIndex500
																			if buffer[position] != rune('R') {
																				goto l491
																			}
																			position++
																		}
																	l500:
																		add(rulePegText, position493)
																	}
																	{
																		add(ruleAction94, position)
																	}
																	add(ruleOtir, position492)
																}
																goto l468
															l491:
																position, tokenIndex = position468, tokenIndex468
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
																				goto l503
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
																				goto l503
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
																				goto l503
																			}
																			position++
																		}
																	l510:
																		{
																			position512, tokenIndex512 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l513
																			}
																			position++
																			goto l512
																		l513:
																			position, tokenIndex = position512, tokenIndex512
																			if buffer[position] != rune('I') {
																				goto l503
																			}
																			position++
																		}
																	l512:
																		add(rulePegText, position505)
																	}
																	{
																		add(ruleAction86, position)
																	}
																	add(ruleOuti, position504)
																}
																goto l468
															l503:
																position, tokenIndex = position468, tokenIndex468
																{
																	position516 := position
																	{
																		position517 := position
																		{
																			position518, tokenIndex518 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l519
																			}
																			position++
																			goto l518
																		l519:
																			position, tokenIndex = position518, tokenIndex518
																			if buffer[position] != rune('I') {
																				goto l515
																			}
																			position++
																		}
																	l518:
																		{
																			position520, tokenIndex520 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l521
																			}
																			position++
																			goto l520
																		l521:
																			position, tokenIndex = position520, tokenIndex520
																			if buffer[position] != rune('N') {
																				goto l515
																			}
																			position++
																		}
																	l520:
																		{
																			position522, tokenIndex522 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l523
																			}
																			position++
																			goto l522
																		l523:
																			position, tokenIndex = position522, tokenIndex522
																			if buffer[position] != rune('D') {
																				goto l515
																			}
																			position++
																		}
																	l522:
																		{
																			position524, tokenIndex524 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l525
																			}
																			position++
																			goto l524
																		l525:
																			position, tokenIndex = position524, tokenIndex524
																			if buffer[position] != rune('R') {
																				goto l515
																			}
																			position++
																		}
																	l524:
																		add(rulePegText, position517)
																	}
																	{
																		add(ruleAction97, position)
																	}
																	add(ruleIndr, position516)
																}
																goto l468
															l515:
																position, tokenIndex = position468, tokenIndex468
																{
																	position528 := position
																	{
																		position529 := position
																		{
																			position530, tokenIndex530 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l531
																			}
																			position++
																			goto l530
																		l531:
																			position, tokenIndex = position530, tokenIndex530
																			if buffer[position] != rune('I') {
																				goto l527
																			}
																			position++
																		}
																	l530:
																		{
																			position532, tokenIndex532 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l533
																			}
																			position++
																			goto l532
																		l533:
																			position, tokenIndex = position532, tokenIndex532
																			if buffer[position] != rune('N') {
																				goto l527
																			}
																			position++
																		}
																	l532:
																		{
																			position534, tokenIndex534 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l535
																			}
																			position++
																			goto l534
																		l535:
																			position, tokenIndex = position534, tokenIndex534
																			if buffer[position] != rune('D') {
																				goto l527
																			}
																			position++
																		}
																	l534:
																		add(rulePegText, position529)
																	}
																	{
																		add(ruleAction89, position)
																	}
																	add(ruleInd, position528)
																}
																goto l468
															l527:
																position, tokenIndex = position468, tokenIndex468
																{
																	position538 := position
																	{
																		position539 := position
																		{
																			position540, tokenIndex540 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l541
																			}
																			position++
																			goto l540
																		l541:
																			position, tokenIndex = position540, tokenIndex540
																			if buffer[position] != rune('O') {
																				goto l537
																			}
																			position++
																		}
																	l540:
																		{
																			position542, tokenIndex542 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l543
																			}
																			position++
																			goto l542
																		l543:
																			position, tokenIndex = position542, tokenIndex542
																			if buffer[position] != rune('T') {
																				goto l537
																			}
																			position++
																		}
																	l542:
																		{
																			position544, tokenIndex544 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l545
																			}
																			position++
																			goto l544
																		l545:
																			position, tokenIndex = position544, tokenIndex544
																			if buffer[position] != rune('D') {
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
																		add(ruleAction98, position)
																	}
																	add(ruleOtdr, position538)
																}
																goto l468
															l537:
																position, tokenIndex = position468, tokenIndex468
																{
																	position549 := position
																	{
																		position550 := position
																		{
																			position551, tokenIndex551 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l552
																			}
																			position++
																			goto l551
																		l552:
																			position, tokenIndex = position551, tokenIndex551
																			if buffer[position] != rune('O') {
																				goto l405
																			}
																			position++
																		}
																	l551:
																		{
																			position553, tokenIndex553 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l554
																			}
																			position++
																			goto l553
																		l554:
																			position, tokenIndex = position553, tokenIndex553
																			if buffer[position] != rune('U') {
																				goto l405
																			}
																			position++
																		}
																	l553:
																		{
																			position555, tokenIndex555 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l556
																			}
																			position++
																			goto l555
																		l556:
																			position, tokenIndex = position555, tokenIndex555
																			if buffer[position] != rune('T') {
																				goto l405
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
																				goto l405
																			}
																			position++
																		}
																	l557:
																		add(rulePegText, position550)
																	}
																	{
																		add(ruleAction90, position)
																	}
																	add(ruleOutd, position549)
																}
															}
														l468:
															add(ruleBlitIO, position467)
														}
														break
													case 'R', 'r':
														{
															position560 := position
															{
																position561 := position
																{
																	position562, tokenIndex562 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l563
																	}
																	position++
																	goto l562
																l563:
																	position, tokenIndex = position562, tokenIndex562
																	if buffer[position] != rune('R') {
																		goto l405
																	}
																	position++
																}
															l562:
																{
																	position564, tokenIndex564 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l565
																	}
																	position++
																	goto l564
																l565:
																	position, tokenIndex = position564, tokenIndex564
																	if buffer[position] != rune('L') {
																		goto l405
																	}
																	position++
																}
															l564:
																{
																	position566, tokenIndex566 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l567
																	}
																	position++
																	goto l566
																l567:
																	position, tokenIndex = position566, tokenIndex566
																	if buffer[position] != rune('D') {
																		goto l405
																	}
																	position++
																}
															l566:
																add(rulePegText, position561)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleRld, position560)
														}
														break
													case 'N', 'n':
														{
															position569 := position
															{
																position570 := position
																{
																	position571, tokenIndex571 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l572
																	}
																	position++
																	goto l571
																l572:
																	position, tokenIndex = position571, tokenIndex571
																	if buffer[position] != rune('N') {
																		goto l405
																	}
																	position++
																}
															l571:
																{
																	position573, tokenIndex573 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l574
																	}
																	position++
																	goto l573
																l574:
																	position, tokenIndex = position573, tokenIndex573
																	if buffer[position] != rune('E') {
																		goto l405
																	}
																	position++
																}
															l573:
																{
																	position575, tokenIndex575 := position, tokenIndex
																	if buffer[position] != rune('g') {
																		goto l576
																	}
																	position++
																	goto l575
																l576:
																	position, tokenIndex = position575, tokenIndex575
																	if buffer[position] != rune('G') {
																		goto l405
																	}
																	position++
																}
															l575:
																add(rulePegText, position570)
															}
															{
																add(ruleAction75, position)
															}
															add(ruleNeg, position569)
														}
														break
													default:
														{
															position578 := position
															{
																position579, tokenIndex579 := position, tokenIndex
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
																			if buffer[position] != rune('i') {
																				goto l588
																			}
																			position++
																			goto l587
																		l588:
																			position, tokenIndex = position587, tokenIndex587
																			if buffer[position] != rune('I') {
																				goto l580
																			}
																			position++
																		}
																	l587:
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
																				goto l580
																			}
																			position++
																		}
																	l589:
																		add(rulePegText, position582)
																	}
																	{
																		add(ruleAction91, position)
																	}
																	add(ruleLdir, position581)
																}
																goto l579
															l580:
																position, tokenIndex = position579, tokenIndex579
																{
																	position593 := position
																	{
																		position594 := position
																		{
																			position595, tokenIndex595 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l596
																			}
																			position++
																			goto l595
																		l596:
																			position, tokenIndex = position595, tokenIndex595
																			if buffer[position] != rune('L') {
																				goto l592
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
																				goto l592
																			}
																			position++
																		}
																	l597:
																		{
																			position599, tokenIndex599 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l600
																			}
																			position++
																			goto l599
																		l600:
																			position, tokenIndex = position599, tokenIndex599
																			if buffer[position] != rune('I') {
																				goto l592
																			}
																			position++
																		}
																	l599:
																		add(rulePegText, position594)
																	}
																	{
																		add(ruleAction83, position)
																	}
																	add(ruleLdi, position593)
																}
																goto l579
															l592:
																position, tokenIndex = position579, tokenIndex579
																{
																	position603 := position
																	{
																		position604 := position
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
																				goto l602
																			}
																			position++
																		}
																	l605:
																		{
																			position607, tokenIndex607 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l608
																			}
																			position++
																			goto l607
																		l608:
																			position, tokenIndex = position607, tokenIndex607
																			if buffer[position] != rune('P') {
																				goto l602
																			}
																			position++
																		}
																	l607:
																		{
																			position609, tokenIndex609 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l610
																			}
																			position++
																			goto l609
																		l610:
																			position, tokenIndex = position609, tokenIndex609
																			if buffer[position] != rune('I') {
																				goto l602
																			}
																			position++
																		}
																	l609:
																		{
																			position611, tokenIndex611 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l612
																			}
																			position++
																			goto l611
																		l612:
																			position, tokenIndex = position611, tokenIndex611
																			if buffer[position] != rune('R') {
																				goto l602
																			}
																			position++
																		}
																	l611:
																		add(rulePegText, position604)
																	}
																	{
																		add(ruleAction92, position)
																	}
																	add(ruleCpir, position603)
																}
																goto l579
															l602:
																position, tokenIndex = position579, tokenIndex579
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
																				goto l614
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
																				goto l614
																			}
																			position++
																		}
																	l619:
																		{
																			position621, tokenIndex621 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l622
																			}
																			position++
																			goto l621
																		l622:
																			position, tokenIndex = position621, tokenIndex621
																			if buffer[position] != rune('I') {
																				goto l614
																			}
																			position++
																		}
																	l621:
																		add(rulePegText, position616)
																	}
																	{
																		add(ruleAction84, position)
																	}
																	add(ruleCpi, position615)
																}
																goto l579
															l614:
																position, tokenIndex = position579, tokenIndex579
																{
																	position625 := position
																	{
																		position626 := position
																		{
																			position627, tokenIndex627 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l628
																			}
																			position++
																			goto l627
																		l628:
																			position, tokenIndex = position627, tokenIndex627
																			if buffer[position] != rune('L') {
																				goto l624
																			}
																			position++
																		}
																	l627:
																		{
																			position629, tokenIndex629 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l630
																			}
																			position++
																			goto l629
																		l630:
																			position, tokenIndex = position629, tokenIndex629
																			if buffer[position] != rune('D') {
																				goto l624
																			}
																			position++
																		}
																	l629:
																		{
																			position631, tokenIndex631 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l632
																			}
																			position++
																			goto l631
																		l632:
																			position, tokenIndex = position631, tokenIndex631
																			if buffer[position] != rune('D') {
																				goto l624
																			}
																			position++
																		}
																	l631:
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
																				goto l624
																			}
																			position++
																		}
																	l633:
																		add(rulePegText, position626)
																	}
																	{
																		add(ruleAction95, position)
																	}
																	add(ruleLddr, position625)
																}
																goto l579
															l624:
																position, tokenIndex = position579, tokenIndex579
																{
																	position637 := position
																	{
																		position638 := position
																		{
																			position639, tokenIndex639 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l640
																			}
																			position++
																			goto l639
																		l640:
																			position, tokenIndex = position639, tokenIndex639
																			if buffer[position] != rune('L') {
																				goto l636
																			}
																			position++
																		}
																	l639:
																		{
																			position641, tokenIndex641 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l642
																			}
																			position++
																			goto l641
																		l642:
																			position, tokenIndex = position641, tokenIndex641
																			if buffer[position] != rune('D') {
																				goto l636
																			}
																			position++
																		}
																	l641:
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
																				goto l636
																			}
																			position++
																		}
																	l643:
																		add(rulePegText, position638)
																	}
																	{
																		add(ruleAction87, position)
																	}
																	add(ruleLdd, position637)
																}
																goto l579
															l636:
																position, tokenIndex = position579, tokenIndex579
																{
																	position647 := position
																	{
																		position648 := position
																		{
																			position649, tokenIndex649 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l650
																			}
																			position++
																			goto l649
																		l650:
																			position, tokenIndex = position649, tokenIndex649
																			if buffer[position] != rune('C') {
																				goto l646
																			}
																			position++
																		}
																	l649:
																		{
																			position651, tokenIndex651 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l652
																			}
																			position++
																			goto l651
																		l652:
																			position, tokenIndex = position651, tokenIndex651
																			if buffer[position] != rune('P') {
																				goto l646
																			}
																			position++
																		}
																	l651:
																		{
																			position653, tokenIndex653 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l654
																			}
																			position++
																			goto l653
																		l654:
																			position, tokenIndex = position653, tokenIndex653
																			if buffer[position] != rune('D') {
																				goto l646
																			}
																			position++
																		}
																	l653:
																		{
																			position655, tokenIndex655 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l656
																			}
																			position++
																			goto l655
																		l656:
																			position, tokenIndex = position655, tokenIndex655
																			if buffer[position] != rune('R') {
																				goto l646
																			}
																			position++
																		}
																	l655:
																		add(rulePegText, position648)
																	}
																	{
																		add(ruleAction96, position)
																	}
																	add(ruleCpdr, position647)
																}
																goto l579
															l646:
																position, tokenIndex = position579, tokenIndex579
																{
																	position658 := position
																	{
																		position659 := position
																		{
																			position660, tokenIndex660 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l661
																			}
																			position++
																			goto l660
																		l661:
																			position, tokenIndex = position660, tokenIndex660
																			if buffer[position] != rune('C') {
																				goto l405
																			}
																			position++
																		}
																	l660:
																		{
																			position662, tokenIndex662 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l663
																			}
																			position++
																			goto l662
																		l663:
																			position, tokenIndex = position662, tokenIndex662
																			if buffer[position] != rune('P') {
																				goto l405
																			}
																			position++
																		}
																	l662:
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
																				goto l405
																			}
																			position++
																		}
																	l664:
																		add(rulePegText, position659)
																	}
																	{
																		add(ruleAction88, position)
																	}
																	add(ruleCpd, position658)
																}
															}
														l579:
															add(ruleBlit, position578)
														}
														break
													}
												}

											}
										l407:
											add(ruleEDSimple, position406)
										}
										goto l77
									l405:
										position, tokenIndex = position77, tokenIndex77
										{
											position668 := position
											{
												position669, tokenIndex669 := position, tokenIndex
												{
													position671 := position
													{
														position672 := position
														{
															position673, tokenIndex673 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l674
															}
															position++
															goto l673
														l674:
															position, tokenIndex = position673, tokenIndex673
															if buffer[position] != rune('R') {
																goto l670
															}
															position++
														}
													l673:
														{
															position675, tokenIndex675 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l676
															}
															position++
															goto l675
														l676:
															position, tokenIndex = position675, tokenIndex675
															if buffer[position] != rune('L') {
																goto l670
															}
															position++
														}
													l675:
														{
															position677, tokenIndex677 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l678
															}
															position++
															goto l677
														l678:
															position, tokenIndex = position677, tokenIndex677
															if buffer[position] != rune('C') {
																goto l670
															}
															position++
														}
													l677:
														{
															position679, tokenIndex679 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l680
															}
															position++
															goto l679
														l680:
															position, tokenIndex = position679, tokenIndex679
															if buffer[position] != rune('A') {
																goto l670
															}
															position++
														}
													l679:
														add(rulePegText, position672)
													}
													{
														add(ruleAction64, position)
													}
													add(ruleRlca, position671)
												}
												goto l669
											l670:
												position, tokenIndex = position669, tokenIndex669
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
																goto l682
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
																goto l682
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
																goto l682
															}
															position++
														}
													l689:
														{
															position691, tokenIndex691 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l692
															}
															position++
															goto l691
														l692:
															position, tokenIndex = position691, tokenIndex691
															if buffer[position] != rune('A') {
																goto l682
															}
															position++
														}
													l691:
														add(rulePegText, position684)
													}
													{
														add(ruleAction65, position)
													}
													add(ruleRrca, position683)
												}
												goto l669
											l682:
												position, tokenIndex = position669, tokenIndex669
												{
													position695 := position
													{
														position696 := position
														{
															position697, tokenIndex697 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l698
															}
															position++
															goto l697
														l698:
															position, tokenIndex = position697, tokenIndex697
															if buffer[position] != rune('R') {
																goto l694
															}
															position++
														}
													l697:
														{
															position699, tokenIndex699 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l700
															}
															position++
															goto l699
														l700:
															position, tokenIndex = position699, tokenIndex699
															if buffer[position] != rune('L') {
																goto l694
															}
															position++
														}
													l699:
														{
															position701, tokenIndex701 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l702
															}
															position++
															goto l701
														l702:
															position, tokenIndex = position701, tokenIndex701
															if buffer[position] != rune('A') {
																goto l694
															}
															position++
														}
													l701:
														add(rulePegText, position696)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleRla, position695)
												}
												goto l669
											l694:
												position, tokenIndex = position669, tokenIndex669
												{
													position705 := position
													{
														position706 := position
														{
															position707, tokenIndex707 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l708
															}
															position++
															goto l707
														l708:
															position, tokenIndex = position707, tokenIndex707
															if buffer[position] != rune('D') {
																goto l704
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
																goto l704
															}
															position++
														}
													l709:
														{
															position711, tokenIndex711 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l712
															}
															position++
															goto l711
														l712:
															position, tokenIndex = position711, tokenIndex711
															if buffer[position] != rune('A') {
																goto l704
															}
															position++
														}
													l711:
														add(rulePegText, position706)
													}
													{
														add(ruleAction68, position)
													}
													add(ruleDaa, position705)
												}
												goto l669
											l704:
												position, tokenIndex = position669, tokenIndex669
												{
													position715 := position
													{
														position716 := position
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
																goto l714
															}
															position++
														}
													l717:
														{
															position719, tokenIndex719 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l720
															}
															position++
															goto l719
														l720:
															position, tokenIndex = position719, tokenIndex719
															if buffer[position] != rune('P') {
																goto l714
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
																goto l714
															}
															position++
														}
													l721:
														add(rulePegText, position716)
													}
													{
														add(ruleAction69, position)
													}
													add(ruleCpl, position715)
												}
												goto l669
											l714:
												position, tokenIndex = position669, tokenIndex669
												{
													position725 := position
													{
														position726 := position
														{
															position727, tokenIndex727 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l728
															}
															position++
															goto l727
														l728:
															position, tokenIndex = position727, tokenIndex727
															if buffer[position] != rune('E') {
																goto l724
															}
															position++
														}
													l727:
														{
															position729, tokenIndex729 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l730
															}
															position++
															goto l729
														l730:
															position, tokenIndex = position729, tokenIndex729
															if buffer[position] != rune('X') {
																goto l724
															}
															position++
														}
													l729:
														{
															position731, tokenIndex731 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l732
															}
															position++
															goto l731
														l732:
															position, tokenIndex = position731, tokenIndex731
															if buffer[position] != rune('X') {
																goto l724
															}
															position++
														}
													l731:
														add(rulePegText, position726)
													}
													{
														add(ruleAction72, position)
													}
													add(ruleExx, position725)
												}
												goto l669
											l724:
												position, tokenIndex = position669, tokenIndex669
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position735 := position
															{
																position736 := position
																{
																	position737, tokenIndex737 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l738
																	}
																	position++
																	goto l737
																l738:
																	position, tokenIndex = position737, tokenIndex737
																	if buffer[position] != rune('E') {
																		goto l667
																	}
																	position++
																}
															l737:
																{
																	position739, tokenIndex739 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l740
																	}
																	position++
																	goto l739
																l740:
																	position, tokenIndex = position739, tokenIndex739
																	if buffer[position] != rune('I') {
																		goto l667
																	}
																	position++
																}
															l739:
																add(rulePegText, position736)
															}
															{
																add(ruleAction74, position)
															}
															add(ruleEi, position735)
														}
														break
													case 'D', 'd':
														{
															position742 := position
															{
																position743 := position
																{
																	position744, tokenIndex744 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l745
																	}
																	position++
																	goto l744
																l745:
																	position, tokenIndex = position744, tokenIndex744
																	if buffer[position] != rune('D') {
																		goto l667
																	}
																	position++
																}
															l744:
																{
																	position746, tokenIndex746 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l747
																	}
																	position++
																	goto l746
																l747:
																	position, tokenIndex = position746, tokenIndex746
																	if buffer[position] != rune('I') {
																		goto l667
																	}
																	position++
																}
															l746:
																add(rulePegText, position743)
															}
															{
																add(ruleAction73, position)
															}
															add(ruleDi, position742)
														}
														break
													case 'C', 'c':
														{
															position749 := position
															{
																position750 := position
																{
																	position751, tokenIndex751 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l752
																	}
																	position++
																	goto l751
																l752:
																	position, tokenIndex = position751, tokenIndex751
																	if buffer[position] != rune('C') {
																		goto l667
																	}
																	position++
																}
															l751:
																{
																	position753, tokenIndex753 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l754
																	}
																	position++
																	goto l753
																l754:
																	position, tokenIndex = position753, tokenIndex753
																	if buffer[position] != rune('C') {
																		goto l667
																	}
																	position++
																}
															l753:
																{
																	position755, tokenIndex755 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l756
																	}
																	position++
																	goto l755
																l756:
																	position, tokenIndex = position755, tokenIndex755
																	if buffer[position] != rune('F') {
																		goto l667
																	}
																	position++
																}
															l755:
																add(rulePegText, position750)
															}
															{
																add(ruleAction71, position)
															}
															add(ruleCcf, position749)
														}
														break
													case 'S', 's':
														{
															position758 := position
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
																		goto l667
																	}
																	position++
																}
															l760:
																{
																	position762, tokenIndex762 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l763
																	}
																	position++
																	goto l762
																l763:
																	position, tokenIndex = position762, tokenIndex762
																	if buffer[position] != rune('C') {
																		goto l667
																	}
																	position++
																}
															l762:
																{
																	position764, tokenIndex764 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l765
																	}
																	position++
																	goto l764
																l765:
																	position, tokenIndex = position764, tokenIndex764
																	if buffer[position] != rune('F') {
																		goto l667
																	}
																	position++
																}
															l764:
																add(rulePegText, position759)
															}
															{
																add(ruleAction70, position)
															}
															add(ruleScf, position758)
														}
														break
													case 'R', 'r':
														{
															position767 := position
															{
																position768 := position
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
																		goto l667
																	}
																	position++
																}
															l769:
																{
																	position771, tokenIndex771 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l772
																	}
																	position++
																	goto l771
																l772:
																	position, tokenIndex = position771, tokenIndex771
																	if buffer[position] != rune('R') {
																		goto l667
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
																		goto l667
																	}
																	position++
																}
															l773:
																add(rulePegText, position768)
															}
															{
																add(ruleAction67, position)
															}
															add(ruleRra, position767)
														}
														break
													case 'H', 'h':
														{
															position776 := position
															{
																position777 := position
																{
																	position778, tokenIndex778 := position, tokenIndex
																	if buffer[position] != rune('h') {
																		goto l779
																	}
																	position++
																	goto l778
																l779:
																	position, tokenIndex = position778, tokenIndex778
																	if buffer[position] != rune('H') {
																		goto l667
																	}
																	position++
																}
															l778:
																{
																	position780, tokenIndex780 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l781
																	}
																	position++
																	goto l780
																l781:
																	position, tokenIndex = position780, tokenIndex780
																	if buffer[position] != rune('A') {
																		goto l667
																	}
																	position++
																}
															l780:
																{
																	position782, tokenIndex782 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l783
																	}
																	position++
																	goto l782
																l783:
																	position, tokenIndex = position782, tokenIndex782
																	if buffer[position] != rune('L') {
																		goto l667
																	}
																	position++
																}
															l782:
																{
																	position784, tokenIndex784 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l785
																	}
																	position++
																	goto l784
																l785:
																	position, tokenIndex = position784, tokenIndex784
																	if buffer[position] != rune('T') {
																		goto l667
																	}
																	position++
																}
															l784:
																add(rulePegText, position777)
															}
															{
																add(ruleAction63, position)
															}
															add(ruleHalt, position776)
														}
														break
													default:
														{
															position787 := position
															{
																position788 := position
																{
																	position789, tokenIndex789 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l790
																	}
																	position++
																	goto l789
																l790:
																	position, tokenIndex = position789, tokenIndex789
																	if buffer[position] != rune('N') {
																		goto l667
																	}
																	position++
																}
															l789:
																{
																	position791, tokenIndex791 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l792
																	}
																	position++
																	goto l791
																l792:
																	position, tokenIndex = position791, tokenIndex791
																	if buffer[position] != rune('O') {
																		goto l667
																	}
																	position++
																}
															l791:
																{
																	position793, tokenIndex793 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l794
																	}
																	position++
																	goto l793
																l794:
																	position, tokenIndex = position793, tokenIndex793
																	if buffer[position] != rune('P') {
																		goto l667
																	}
																	position++
																}
															l793:
																add(rulePegText, position788)
															}
															{
																add(ruleAction62, position)
															}
															add(ruleNop, position787)
														}
														break
													}
												}

											}
										l669:
											add(ruleSimple, position668)
										}
										goto l77
									l667:
										position, tokenIndex = position77, tokenIndex77
										{
											position797 := position
											{
												position798, tokenIndex798 := position, tokenIndex
												{
													position800 := position
													{
														position801, tokenIndex801 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l802
														}
														position++
														goto l801
													l802:
														position, tokenIndex = position801, tokenIndex801
														if buffer[position] != rune('R') {
															goto l799
														}
														position++
													}
												l801:
													{
														position803, tokenIndex803 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l804
														}
														position++
														goto l803
													l804:
														position, tokenIndex = position803, tokenIndex803
														if buffer[position] != rune('S') {
															goto l799
														}
														position++
													}
												l803:
													{
														position805, tokenIndex805 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l806
														}
														position++
														goto l805
													l806:
														position, tokenIndex = position805, tokenIndex805
														if buffer[position] != rune('T') {
															goto l799
														}
														position++
													}
												l805:
													if !_rules[rulews]() {
														goto l799
													}
													if !_rules[rulen]() {
														goto l799
													}
													{
														add(ruleAction99, position)
													}
													add(ruleRst, position800)
												}
												goto l798
											l799:
												position, tokenIndex = position798, tokenIndex798
												{
													position809 := position
													{
														position810, tokenIndex810 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l811
														}
														position++
														goto l810
													l811:
														position, tokenIndex = position810, tokenIndex810
														if buffer[position] != rune('J') {
															goto l808
														}
														position++
													}
												l810:
													{
														position812, tokenIndex812 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l813
														}
														position++
														goto l812
													l813:
														position, tokenIndex = position812, tokenIndex812
														if buffer[position] != rune('P') {
															goto l808
														}
														position++
													}
												l812:
													if !_rules[rulews]() {
														goto l808
													}
													{
														position814, tokenIndex814 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l814
														}
														if !_rules[rulesep]() {
															goto l814
														}
														goto l815
													l814:
														position, tokenIndex = position814, tokenIndex814
													}
												l815:
													if !_rules[ruleSrc16]() {
														goto l808
													}
													{
														add(ruleAction102, position)
													}
													add(ruleJp, position809)
												}
												goto l798
											l808:
												position, tokenIndex = position798, tokenIndex798
												{
													switch buffer[position] {
													case 'D', 'd':
														{
															position818 := position
															{
																position819, tokenIndex819 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l820
																}
																position++
																goto l819
															l820:
																position, tokenIndex = position819, tokenIndex819
																if buffer[position] != rune('D') {
																	goto l796
																}
																position++
															}
														l819:
															{
																position821, tokenIndex821 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l822
																}
																position++
																goto l821
															l822:
																position, tokenIndex = position821, tokenIndex821
																if buffer[position] != rune('J') {
																	goto l796
																}
																position++
															}
														l821:
															{
																position823, tokenIndex823 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l824
																}
																position++
																goto l823
															l824:
																position, tokenIndex = position823, tokenIndex823
																if buffer[position] != rune('N') {
																	goto l796
																}
																position++
															}
														l823:
															{
																position825, tokenIndex825 := position, tokenIndex
																if buffer[position] != rune('z') {
																	goto l826
																}
																position++
																goto l825
															l826:
																position, tokenIndex = position825, tokenIndex825
																if buffer[position] != rune('Z') {
																	goto l796
																}
																position++
															}
														l825:
															if !_rules[rulews]() {
																goto l796
															}
															if !_rules[ruledisp]() {
																goto l796
															}
															{
																add(ruleAction104, position)
															}
															add(ruleDjnz, position818)
														}
														break
													case 'J', 'j':
														{
															position828 := position
															{
																position829, tokenIndex829 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l830
																}
																position++
																goto l829
															l830:
																position, tokenIndex = position829, tokenIndex829
																if buffer[position] != rune('J') {
																	goto l796
																}
																position++
															}
														l829:
															{
																position831, tokenIndex831 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l832
																}
																position++
																goto l831
															l832:
																position, tokenIndex = position831, tokenIndex831
																if buffer[position] != rune('R') {
																	goto l796
																}
																position++
															}
														l831:
															if !_rules[rulews]() {
																goto l796
															}
															{
																position833, tokenIndex833 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l833
																}
																if !_rules[rulesep]() {
																	goto l833
																}
																goto l834
															l833:
																position, tokenIndex = position833, tokenIndex833
															}
														l834:
															if !_rules[ruledisp]() {
																goto l796
															}
															{
																add(ruleAction103, position)
															}
															add(ruleJr, position828)
														}
														break
													case 'R', 'r':
														{
															position836 := position
															{
																position837, tokenIndex837 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l838
																}
																position++
																goto l837
															l838:
																position, tokenIndex = position837, tokenIndex837
																if buffer[position] != rune('R') {
																	goto l796
																}
																position++
															}
														l837:
															{
																position839, tokenIndex839 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l840
																}
																position++
																goto l839
															l840:
																position, tokenIndex = position839, tokenIndex839
																if buffer[position] != rune('E') {
																	goto l796
																}
																position++
															}
														l839:
															{
																position841, tokenIndex841 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l842
																}
																position++
																goto l841
															l842:
																position, tokenIndex = position841, tokenIndex841
																if buffer[position] != rune('T') {
																	goto l796
																}
																position++
															}
														l841:
															{
																position843, tokenIndex843 := position, tokenIndex
																if !_rules[rulews]() {
																	goto l843
																}
																if !_rules[rulecc]() {
																	goto l843
																}
																goto l844
															l843:
																position, tokenIndex = position843, tokenIndex843
															}
														l844:
															{
																add(ruleAction101, position)
															}
															add(ruleRet, position836)
														}
														break
													default:
														{
															position846 := position
															{
																position847, tokenIndex847 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l848
																}
																position++
																goto l847
															l848:
																position, tokenIndex = position847, tokenIndex847
																if buffer[position] != rune('C') {
																	goto l796
																}
																position++
															}
														l847:
															{
																position849, tokenIndex849 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l850
																}
																position++
																goto l849
															l850:
																position, tokenIndex = position849, tokenIndex849
																if buffer[position] != rune('A') {
																	goto l796
																}
																position++
															}
														l849:
															{
																position851, tokenIndex851 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l852
																}
																position++
																goto l851
															l852:
																position, tokenIndex = position851, tokenIndex851
																if buffer[position] != rune('L') {
																	goto l796
																}
																position++
															}
														l851:
															{
																position853, tokenIndex853 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l854
																}
																position++
																goto l853
															l854:
																position, tokenIndex = position853, tokenIndex853
																if buffer[position] != rune('L') {
																	goto l796
																}
																position++
															}
														l853:
															if !_rules[rulews]() {
																goto l796
															}
															{
																position855, tokenIndex855 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l855
																}
																if !_rules[rulesep]() {
																	goto l855
																}
																goto l856
															l855:
																position, tokenIndex = position855, tokenIndex855
															}
														l856:
															if !_rules[ruleSrc16]() {
																goto l796
															}
															{
																add(ruleAction100, position)
															}
															add(ruleCall, position846)
														}
														break
													}
												}

											}
										l798:
											add(ruleJump, position797)
										}
										goto l77
									l796:
										position, tokenIndex = position77, tokenIndex77
										{
											position858 := position
											{
												position859, tokenIndex859 := position, tokenIndex
												{
													position861 := position
													{
														position862, tokenIndex862 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l863
														}
														position++
														goto l862
													l863:
														position, tokenIndex = position862, tokenIndex862
														if buffer[position] != rune('I') {
															goto l860
														}
														position++
													}
												l862:
													{
														position864, tokenIndex864 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l865
														}
														position++
														goto l864
													l865:
														position, tokenIndex = position864, tokenIndex864
														if buffer[position] != rune('N') {
															goto l860
														}
														position++
													}
												l864:
													if !_rules[rulews]() {
														goto l860
													}
													if !_rules[ruleReg8]() {
														goto l860
													}
													if !_rules[rulesep]() {
														goto l860
													}
													if !_rules[rulePort]() {
														goto l860
													}
													{
														add(ruleAction105, position)
													}
													add(ruleIN, position861)
												}
												goto l859
											l860:
												position, tokenIndex = position859, tokenIndex859
												{
													position867 := position
													{
														position868, tokenIndex868 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l869
														}
														position++
														goto l868
													l869:
														position, tokenIndex = position868, tokenIndex868
														if buffer[position] != rune('O') {
															goto l13
														}
														position++
													}
												l868:
													{
														position870, tokenIndex870 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l871
														}
														position++
														goto l870
													l871:
														position, tokenIndex = position870, tokenIndex870
														if buffer[position] != rune('U') {
															goto l13
														}
														position++
													}
												l870:
													{
														position872, tokenIndex872 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l873
														}
														position++
														goto l872
													l873:
														position, tokenIndex = position872, tokenIndex872
														if buffer[position] != rune('T') {
															goto l13
														}
														position++
													}
												l872:
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
														add(ruleAction106, position)
													}
													add(ruleOUT, position867)
												}
											}
										l859:
											add(ruleIO, position858)
										}
									}
								l77:
									add(ruleInstruction, position76)
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
						position875, tokenIndex875 := position, tokenIndex
						if !_rules[rulews]() {
							goto l875
						}
						goto l876
					l875:
						position, tokenIndex = position875, tokenIndex875
					}
				l876:
					{
						position877, tokenIndex877 := position, tokenIndex
						{
							position879 := position
							{
								position880, tokenIndex880 := position, tokenIndex
								if buffer[position] != rune(';') {
									goto l881
								}
								position++
								goto l880
							l881:
								position, tokenIndex = position880, tokenIndex880
								if buffer[position] != rune('#') {
									goto l877
								}
								position++
							}
						l880:
						l882:
							{
								position883, tokenIndex883 := position, tokenIndex
								{
									position884, tokenIndex884 := position, tokenIndex
									if buffer[position] != rune('\n') {
										goto l884
									}
									position++
									goto l883
								l884:
									position, tokenIndex = position884, tokenIndex884
								}
								if !matchDot() {
									goto l883
								}
								goto l882
							l883:
								position, tokenIndex = position883, tokenIndex883
							}
							add(ruleComment, position879)
						}
						goto l878
					l877:
						position, tokenIndex = position877, tokenIndex877
					}
				l878:
					{
						position885, tokenIndex885 := position, tokenIndex
						if !_rules[rulews]() {
							goto l885
						}
						goto l886
					l885:
						position, tokenIndex = position885, tokenIndex885
					}
				l886:
					{
						position887, tokenIndex887 := position, tokenIndex
						{
							position889, tokenIndex889 := position, tokenIndex
							if buffer[position] != rune('\r') {
								goto l889
							}
							position++
							goto l890
						l889:
							position, tokenIndex = position889, tokenIndex889
						}
					l890:
						if buffer[position] != rune('\n') {
							goto l888
						}
						position++
						goto l887
					l888:
						position, tokenIndex = position887, tokenIndex887
						if buffer[position] != rune(':') {
							goto l0
						}
						position++
					}
				l887:
					{
						add(ruleAction0, position)
					}
					add(ruleLine, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position892 := position
					l893:
						{
							position894, tokenIndex894 := position, tokenIndex
							if !_rules[rulews]() {
								goto l894
							}
							goto l893
						l894:
							position, tokenIndex = position894, tokenIndex894
						}
						{
							position895, tokenIndex895 := position, tokenIndex
							{
								position897 := position
								if !_rules[ruleLabelText]() {
									goto l895
								}
								if buffer[position] != rune(':') {
									goto l895
								}
								position++
								if !_rules[rulews]() {
									goto l895
								}
								{
									add(ruleAction4, position)
								}
								add(ruleLabelDefn, position897)
							}
							goto l896
						l895:
							position, tokenIndex = position895, tokenIndex895
						}
					l896:
					l899:
						{
							position900, tokenIndex900 := position, tokenIndex
							if !_rules[rulews]() {
								goto l900
							}
							goto l899
						l900:
							position, tokenIndex = position900, tokenIndex900
						}
						{
							position901, tokenIndex901 := position, tokenIndex
							{
								position903 := position
								{
									position904, tokenIndex904 := position, tokenIndex
									{
										position906 := position
										{
											position907, tokenIndex907 := position, tokenIndex
											{
												position909 := position
												{
													position910, tokenIndex910 := position, tokenIndex
													{
														position912, tokenIndex912 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l913
														}
														position++
														goto l912
													l913:
														position, tokenIndex = position912, tokenIndex912
														if buffer[position] != rune('D') {
															goto l911
														}
														position++
													}
												l912:
													{
														position914, tokenIndex914 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l915
														}
														position++
														goto l914
													l915:
														position, tokenIndex = position914, tokenIndex914
														if buffer[position] != rune('E') {
															goto l911
														}
														position++
													}
												l914:
													{
														position916, tokenIndex916 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l917
														}
														position++
														goto l916
													l917:
														position, tokenIndex = position916, tokenIndex916
														if buffer[position] != rune('F') {
															goto l911
														}
														position++
													}
												l916:
													{
														position918, tokenIndex918 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l919
														}
														position++
														goto l918
													l919:
														position, tokenIndex = position918, tokenIndex918
														if buffer[position] != rune('B') {
															goto l911
														}
														position++
													}
												l918:
													goto l910
												l911:
													position, tokenIndex = position910, tokenIndex910
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
															goto l920
														}
														position++
													}
												l921:
													{
														position923, tokenIndex923 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l924
														}
														position++
														goto l923
													l924:
														position, tokenIndex = position923, tokenIndex923
														if buffer[position] != rune('B') {
															goto l920
														}
														position++
													}
												l923:
													goto l910
												l920:
													position, tokenIndex = position910, tokenIndex910
													{
														position925, tokenIndex925 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l926
														}
														position++
														goto l925
													l926:
														position, tokenIndex = position925, tokenIndex925
														if buffer[position] != rune('E') {
															goto l908
														}
														position++
													}
												l925:
													{
														position927, tokenIndex927 := position, tokenIndex
														if buffer[position] != rune('q') {
															goto l928
														}
														position++
														goto l927
													l928:
														position, tokenIndex = position927, tokenIndex927
														if buffer[position] != rune('Q') {
															goto l908
														}
														position++
													}
												l927:
													{
														position929, tokenIndex929 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l930
														}
														position++
														goto l929
													l930:
														position, tokenIndex = position929, tokenIndex929
														if buffer[position] != rune('U') {
															goto l908
														}
														position++
													}
												l929:
												}
											l910:
												if !_rules[rulews]() {
													goto l908
												}
												if !_rules[rulen]() {
													goto l908
												}
												{
													add(ruleAction2, position)
												}
												add(ruleDefb, position909)
											}
											goto l907
										l908:
											position, tokenIndex = position907, tokenIndex907
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position933 := position
														{
															position934, tokenIndex934 := position, tokenIndex
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
																	goto l935
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
																	goto l935
																}
																position++
															}
														l938:
															{
																position940, tokenIndex940 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l941
																}
																position++
																goto l940
															l941:
																position, tokenIndex = position940, tokenIndex940
																if buffer[position] != rune('F') {
																	goto l935
																}
																position++
															}
														l940:
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
																	goto l935
																}
																position++
															}
														l942:
															goto l934
														l935:
															position, tokenIndex = position934, tokenIndex934
															{
																position944, tokenIndex944 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l945
																}
																position++
																goto l944
															l945:
																position, tokenIndex = position944, tokenIndex944
																if buffer[position] != rune('D') {
																	goto l905
																}
																position++
															}
														l944:
															{
																position946, tokenIndex946 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l947
																}
																position++
																goto l946
															l947:
																position, tokenIndex = position946, tokenIndex946
																if buffer[position] != rune('S') {
																	goto l905
																}
																position++
															}
														l946:
														}
													l934:
														if !_rules[rulews]() {
															goto l905
														}
														if !_rules[rulen]() {
															goto l905
														}
														{
															add(ruleAction3, position)
														}
														add(ruleDefs, position933)
													}
													break
												case 'O', 'o':
													{
														position949 := position
														{
															position950, tokenIndex950 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l951
															}
															position++
															goto l950
														l951:
															position, tokenIndex = position950, tokenIndex950
															if buffer[position] != rune('O') {
																goto l905
															}
															position++
														}
													l950:
														{
															position952, tokenIndex952 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l953
															}
															position++
															goto l952
														l953:
															position, tokenIndex = position952, tokenIndex952
															if buffer[position] != rune('R') {
																goto l905
															}
															position++
														}
													l952:
														{
															position954, tokenIndex954 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l955
															}
															position++
															goto l954
														l955:
															position, tokenIndex = position954, tokenIndex954
															if buffer[position] != rune('G') {
																goto l905
															}
															position++
														}
													l954:
														if !_rules[rulews]() {
															goto l905
														}
														if !_rules[rulenn]() {
															goto l905
														}
														{
															add(ruleAction1, position)
														}
														add(ruleOrg, position949)
													}
													break
												case 'a':
													{
														position957 := position
														if buffer[position] != rune('a') {
															goto l905
														}
														position++
														if buffer[position] != rune('s') {
															goto l905
														}
														position++
														if buffer[position] != rune('e') {
															goto l905
														}
														position++
														if buffer[position] != rune('g') {
															goto l905
														}
														position++
														add(ruleAseg, position957)
													}
													break
												default:
													{
														position958 := position
														{
															position959, tokenIndex959 := position, tokenIndex
															if buffer[position] != rune('.') {
																goto l959
															}
															position++
															goto l960
														l959:
															position, tokenIndex = position959, tokenIndex959
														}
													l960:
														if buffer[position] != rune('t') {
															goto l905
														}
														position++
														if buffer[position] != rune('i') {
															goto l905
														}
														position++
														if buffer[position] != rune('t') {
															goto l905
														}
														position++
														if buffer[position] != rune('l') {
															goto l905
														}
														position++
														if buffer[position] != rune('e') {
															goto l905
														}
														position++
														if !_rules[rulews]() {
															goto l905
														}
														if buffer[position] != rune('\'') {
															goto l905
														}
														position++
													l961:
														{
															position962, tokenIndex962 := position, tokenIndex
															{
																position963, tokenIndex963 := position, tokenIndex
																if buffer[position] != rune('\'') {
																	goto l963
																}
																position++
																goto l962
															l963:
																position, tokenIndex = position963, tokenIndex963
															}
															if !matchDot() {
																goto l962
															}
															goto l961
														l962:
															position, tokenIndex = position962, tokenIndex962
														}
														if buffer[position] != rune('\'') {
															goto l905
														}
														position++
														add(ruleTitle, position958)
													}
													break
												}
											}

										}
									l907:
										add(ruleDirective, position906)
									}
									goto l904
								l905:
									position, tokenIndex = position904, tokenIndex904
									{
										position964 := position
										{
											position965, tokenIndex965 := position, tokenIndex
											{
												position967 := position
												{
													position968, tokenIndex968 := position, tokenIndex
													{
														position970 := position
														{
															position971, tokenIndex971 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l972
															}
															position++
															goto l971
														l972:
															position, tokenIndex = position971, tokenIndex971
															if buffer[position] != rune('P') {
																goto l969
															}
															position++
														}
													l971:
														{
															position973, tokenIndex973 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l974
															}
															position++
															goto l973
														l974:
															position, tokenIndex = position973, tokenIndex973
															if buffer[position] != rune('U') {
																goto l969
															}
															position++
														}
													l973:
														{
															position975, tokenIndex975 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l976
															}
															position++
															goto l975
														l976:
															position, tokenIndex = position975, tokenIndex975
															if buffer[position] != rune('S') {
																goto l969
															}
															position++
														}
													l975:
														{
															position977, tokenIndex977 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l978
															}
															position++
															goto l977
														l978:
															position, tokenIndex = position977, tokenIndex977
															if buffer[position] != rune('H') {
																goto l969
															}
															position++
														}
													l977:
														if !_rules[rulews]() {
															goto l969
														}
														if !_rules[ruleSrc16]() {
															goto l969
														}
														{
															add(ruleAction7, position)
														}
														add(rulePush, position970)
													}
													goto l968
												l969:
													position, tokenIndex = position968, tokenIndex968
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position981 := position
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
																		goto l966
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
																		goto l966
																	}
																	position++
																}
															l984:
																if !_rules[rulews]() {
																	goto l966
																}
																if !_rules[ruleDst16]() {
																	goto l966
																}
																if !_rules[rulesep]() {
																	goto l966
																}
																if !_rules[ruleSrc16]() {
																	goto l966
																}
																{
																	add(ruleAction9, position)
																}
																add(ruleEx, position981)
															}
															break
														case 'P', 'p':
															{
																position987 := position
																{
																	position988, tokenIndex988 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l989
																	}
																	position++
																	goto l988
																l989:
																	position, tokenIndex = position988, tokenIndex988
																	if buffer[position] != rune('P') {
																		goto l966
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
																		goto l966
																	}
																	position++
																}
															l990:
																{
																	position992, tokenIndex992 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l993
																	}
																	position++
																	goto l992
																l993:
																	position, tokenIndex = position992, tokenIndex992
																	if buffer[position] != rune('P') {
																		goto l966
																	}
																	position++
																}
															l992:
																if !_rules[rulews]() {
																	goto l966
																}
																if !_rules[ruleDst16]() {
																	goto l966
																}
																{
																	add(ruleAction8, position)
																}
																add(rulePop, position987)
															}
															break
														default:
															{
																position995 := position
																{
																	position996, tokenIndex996 := position, tokenIndex
																	{
																		position998 := position
																		{
																			position999, tokenIndex999 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l1000
																			}
																			position++
																			goto l999
																		l1000:
																			position, tokenIndex = position999, tokenIndex999
																			if buffer[position] != rune('L') {
																				goto l997
																			}
																			position++
																		}
																	l999:
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
																				goto l997
																			}
																			position++
																		}
																	l1001:
																		if !_rules[rulews]() {
																			goto l997
																		}
																		if !_rules[ruleDst16]() {
																			goto l997
																		}
																		if !_rules[rulesep]() {
																			goto l997
																		}
																		if !_rules[ruleSrc16]() {
																			goto l997
																		}
																		{
																			add(ruleAction6, position)
																		}
																		add(ruleLoad16, position998)
																	}
																	goto l996
																l997:
																	position, tokenIndex = position996, tokenIndex996
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
																				goto l966
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
																				goto l966
																			}
																			position++
																		}
																	l1007:
																		if !_rules[rulews]() {
																			goto l966
																		}
																		{
																			position1009 := position
																			{
																				position1010, tokenIndex1010 := position, tokenIndex
																				if !_rules[ruleReg8]() {
																					goto l1011
																				}
																				goto l1010
																			l1011:
																				position, tokenIndex = position1010, tokenIndex1010
																				if !_rules[ruleReg16Contents]() {
																					goto l1012
																				}
																				goto l1010
																			l1012:
																				position, tokenIndex = position1010, tokenIndex1010
																				if !_rules[rulenn_contents]() {
																					goto l966
																				}
																			}
																		l1010:
																			{
																				add(ruleAction19, position)
																			}
																			add(ruleDst8, position1009)
																		}
																		if !_rules[rulesep]() {
																			goto l966
																		}
																		if !_rules[ruleSrc8]() {
																			goto l966
																		}
																		{
																			add(ruleAction5, position)
																		}
																		add(ruleLoad8, position1004)
																	}
																}
															l996:
																add(ruleLoad, position995)
															}
															break
														}
													}

												}
											l968:
												add(ruleAssignment, position967)
											}
											goto l965
										l966:
											position, tokenIndex = position965, tokenIndex965
											{
												position1016 := position
												{
													position1017, tokenIndex1017 := position, tokenIndex
													{
														position1019 := position
														{
															position1020, tokenIndex1020 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1021
															}
															position++
															goto l1020
														l1021:
															position, tokenIndex = position1020, tokenIndex1020
															if buffer[position] != rune('I') {
																goto l1018
															}
															position++
														}
													l1020:
														{
															position1022, tokenIndex1022 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1023
															}
															position++
															goto l1022
														l1023:
															position, tokenIndex = position1022, tokenIndex1022
															if buffer[position] != rune('N') {
																goto l1018
															}
															position++
														}
													l1022:
														{
															position1024, tokenIndex1024 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1025
															}
															position++
															goto l1024
														l1025:
															position, tokenIndex = position1024, tokenIndex1024
															if buffer[position] != rune('C') {
																goto l1018
															}
															position++
														}
													l1024:
														if !_rules[rulews]() {
															goto l1018
														}
														if !_rules[ruleILoc8]() {
															goto l1018
														}
														{
															add(ruleAction10, position)
														}
														add(ruleInc16Indexed8, position1019)
													}
													goto l1017
												l1018:
													position, tokenIndex = position1017, tokenIndex1017
													{
														position1028 := position
														{
															position1029, tokenIndex1029 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1030
															}
															position++
															goto l1029
														l1030:
															position, tokenIndex = position1029, tokenIndex1029
															if buffer[position] != rune('I') {
																goto l1027
															}
															position++
														}
													l1029:
														{
															position1031, tokenIndex1031 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1032
															}
															position++
															goto l1031
														l1032:
															position, tokenIndex = position1031, tokenIndex1031
															if buffer[position] != rune('N') {
																goto l1027
															}
															position++
														}
													l1031:
														{
															position1033, tokenIndex1033 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1034
															}
															position++
															goto l1033
														l1034:
															position, tokenIndex = position1033, tokenIndex1033
															if buffer[position] != rune('C') {
																goto l1027
															}
															position++
														}
													l1033:
														if !_rules[rulews]() {
															goto l1027
														}
														if !_rules[ruleLoc16]() {
															goto l1027
														}
														{
															add(ruleAction12, position)
														}
														add(ruleInc16, position1028)
													}
													goto l1017
												l1027:
													position, tokenIndex = position1017, tokenIndex1017
													{
														position1036 := position
														{
															position1037, tokenIndex1037 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1038
															}
															position++
															goto l1037
														l1038:
															position, tokenIndex = position1037, tokenIndex1037
															if buffer[position] != rune('I') {
																goto l1015
															}
															position++
														}
													l1037:
														{
															position1039, tokenIndex1039 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1040
															}
															position++
															goto l1039
														l1040:
															position, tokenIndex = position1039, tokenIndex1039
															if buffer[position] != rune('N') {
																goto l1015
															}
															position++
														}
													l1039:
														{
															position1041, tokenIndex1041 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1042
															}
															position++
															goto l1041
														l1042:
															position, tokenIndex = position1041, tokenIndex1041
															if buffer[position] != rune('C') {
																goto l1015
															}
															position++
														}
													l1041:
														if !_rules[rulews]() {
															goto l1015
														}
														if !_rules[ruleLoc8]() {
															goto l1015
														}
														{
															add(ruleAction11, position)
														}
														add(ruleInc8, position1036)
													}
												}
											l1017:
												add(ruleInc, position1016)
											}
											goto l965
										l1015:
											position, tokenIndex = position965, tokenIndex965
											{
												position1045 := position
												{
													position1046, tokenIndex1046 := position, tokenIndex
													{
														position1048 := position
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
																goto l1047
															}
															position++
														}
													l1049:
														{
															position1051, tokenIndex1051 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1052
															}
															position++
															goto l1051
														l1052:
															position, tokenIndex = position1051, tokenIndex1051
															if buffer[position] != rune('E') {
																goto l1047
															}
															position++
														}
													l1051:
														{
															position1053, tokenIndex1053 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1054
															}
															position++
															goto l1053
														l1054:
															position, tokenIndex = position1053, tokenIndex1053
															if buffer[position] != rune('C') {
																goto l1047
															}
															position++
														}
													l1053:
														if !_rules[rulews]() {
															goto l1047
														}
														if !_rules[ruleILoc8]() {
															goto l1047
														}
														{
															add(ruleAction13, position)
														}
														add(ruleDec16Indexed8, position1048)
													}
													goto l1046
												l1047:
													position, tokenIndex = position1046, tokenIndex1046
													{
														position1057 := position
														{
															position1058, tokenIndex1058 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1059
															}
															position++
															goto l1058
														l1059:
															position, tokenIndex = position1058, tokenIndex1058
															if buffer[position] != rune('D') {
																goto l1056
															}
															position++
														}
													l1058:
														{
															position1060, tokenIndex1060 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1061
															}
															position++
															goto l1060
														l1061:
															position, tokenIndex = position1060, tokenIndex1060
															if buffer[position] != rune('E') {
																goto l1056
															}
															position++
														}
													l1060:
														{
															position1062, tokenIndex1062 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1063
															}
															position++
															goto l1062
														l1063:
															position, tokenIndex = position1062, tokenIndex1062
															if buffer[position] != rune('C') {
																goto l1056
															}
															position++
														}
													l1062:
														if !_rules[rulews]() {
															goto l1056
														}
														if !_rules[ruleLoc16]() {
															goto l1056
														}
														{
															add(ruleAction15, position)
														}
														add(ruleDec16, position1057)
													}
													goto l1046
												l1056:
													position, tokenIndex = position1046, tokenIndex1046
													{
														position1065 := position
														{
															position1066, tokenIndex1066 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1067
															}
															position++
															goto l1066
														l1067:
															position, tokenIndex = position1066, tokenIndex1066
															if buffer[position] != rune('D') {
																goto l1044
															}
															position++
														}
													l1066:
														{
															position1068, tokenIndex1068 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1069
															}
															position++
															goto l1068
														l1069:
															position, tokenIndex = position1068, tokenIndex1068
															if buffer[position] != rune('E') {
																goto l1044
															}
															position++
														}
													l1068:
														{
															position1070, tokenIndex1070 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1071
															}
															position++
															goto l1070
														l1071:
															position, tokenIndex = position1070, tokenIndex1070
															if buffer[position] != rune('C') {
																goto l1044
															}
															position++
														}
													l1070:
														if !_rules[rulews]() {
															goto l1044
														}
														if !_rules[ruleLoc8]() {
															goto l1044
														}
														{
															add(ruleAction14, position)
														}
														add(ruleDec8, position1065)
													}
												}
											l1046:
												add(ruleDec, position1045)
											}
											goto l965
										l1044:
											position, tokenIndex = position965, tokenIndex965
											{
												position1074 := position
												{
													position1075, tokenIndex1075 := position, tokenIndex
													{
														position1077 := position
														{
															position1078, tokenIndex1078 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1079
															}
															position++
															goto l1078
														l1079:
															position, tokenIndex = position1078, tokenIndex1078
															if buffer[position] != rune('A') {
																goto l1076
															}
															position++
														}
													l1078:
														{
															position1080, tokenIndex1080 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1081
															}
															position++
															goto l1080
														l1081:
															position, tokenIndex = position1080, tokenIndex1080
															if buffer[position] != rune('D') {
																goto l1076
															}
															position++
														}
													l1080:
														{
															position1082, tokenIndex1082 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1083
															}
															position++
															goto l1082
														l1083:
															position, tokenIndex = position1082, tokenIndex1082
															if buffer[position] != rune('D') {
																goto l1076
															}
															position++
														}
													l1082:
														if !_rules[rulews]() {
															goto l1076
														}
														if !_rules[ruleDst16]() {
															goto l1076
														}
														if !_rules[rulesep]() {
															goto l1076
														}
														if !_rules[ruleSrc16]() {
															goto l1076
														}
														{
															add(ruleAction16, position)
														}
														add(ruleAdd16, position1077)
													}
													goto l1075
												l1076:
													position, tokenIndex = position1075, tokenIndex1075
													{
														position1086 := position
														{
															position1087, tokenIndex1087 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1088
															}
															position++
															goto l1087
														l1088:
															position, tokenIndex = position1087, tokenIndex1087
															if buffer[position] != rune('A') {
																goto l1085
															}
															position++
														}
													l1087:
														{
															position1089, tokenIndex1089 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1090
															}
															position++
															goto l1089
														l1090:
															position, tokenIndex = position1089, tokenIndex1089
															if buffer[position] != rune('D') {
																goto l1085
															}
															position++
														}
													l1089:
														{
															position1091, tokenIndex1091 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1092
															}
															position++
															goto l1091
														l1092:
															position, tokenIndex = position1091, tokenIndex1091
															if buffer[position] != rune('C') {
																goto l1085
															}
															position++
														}
													l1091:
														if !_rules[rulews]() {
															goto l1085
														}
														if !_rules[ruleDst16]() {
															goto l1085
														}
														if !_rules[rulesep]() {
															goto l1085
														}
														if !_rules[ruleSrc16]() {
															goto l1085
														}
														{
															add(ruleAction17, position)
														}
														add(ruleAdc16, position1086)
													}
													goto l1075
												l1085:
													position, tokenIndex = position1075, tokenIndex1075
													{
														position1094 := position
														{
															position1095, tokenIndex1095 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1096
															}
															position++
															goto l1095
														l1096:
															position, tokenIndex = position1095, tokenIndex1095
															if buffer[position] != rune('S') {
																goto l1073
															}
															position++
														}
													l1095:
														{
															position1097, tokenIndex1097 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1098
															}
															position++
															goto l1097
														l1098:
															position, tokenIndex = position1097, tokenIndex1097
															if buffer[position] != rune('B') {
																goto l1073
															}
															position++
														}
													l1097:
														{
															position1099, tokenIndex1099 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1100
															}
															position++
															goto l1099
														l1100:
															position, tokenIndex = position1099, tokenIndex1099
															if buffer[position] != rune('C') {
																goto l1073
															}
															position++
														}
													l1099:
														if !_rules[rulews]() {
															goto l1073
														}
														if !_rules[ruleDst16]() {
															goto l1073
														}
														if !_rules[rulesep]() {
															goto l1073
														}
														if !_rules[ruleSrc16]() {
															goto l1073
														}
														{
															add(ruleAction18, position)
														}
														add(ruleSbc16, position1094)
													}
												}
											l1075:
												add(ruleAlu16, position1074)
											}
											goto l965
										l1073:
											position, tokenIndex = position965, tokenIndex965
											{
												position1103 := position
												{
													position1104, tokenIndex1104 := position, tokenIndex
													{
														position1106 := position
														{
															position1107, tokenIndex1107 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1108
															}
															position++
															goto l1107
														l1108:
															position, tokenIndex = position1107, tokenIndex1107
															if buffer[position] != rune('A') {
																goto l1105
															}
															position++
														}
													l1107:
														{
															position1109, tokenIndex1109 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1110
															}
															position++
															goto l1109
														l1110:
															position, tokenIndex = position1109, tokenIndex1109
															if buffer[position] != rune('D') {
																goto l1105
															}
															position++
														}
													l1109:
														{
															position1111, tokenIndex1111 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1112
															}
															position++
															goto l1111
														l1112:
															position, tokenIndex = position1111, tokenIndex1111
															if buffer[position] != rune('D') {
																goto l1105
															}
															position++
														}
													l1111:
														if !_rules[rulews]() {
															goto l1105
														}
														{
															position1113, tokenIndex1113 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1114
															}
															position++
															goto l1113
														l1114:
															position, tokenIndex = position1113, tokenIndex1113
															if buffer[position] != rune('A') {
																goto l1105
															}
															position++
														}
													l1113:
														if !_rules[rulesep]() {
															goto l1105
														}
														if !_rules[ruleSrc8]() {
															goto l1105
														}
														{
															add(ruleAction43, position)
														}
														add(ruleAdd, position1106)
													}
													goto l1104
												l1105:
													position, tokenIndex = position1104, tokenIndex1104
													{
														position1117 := position
														{
															position1118, tokenIndex1118 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1119
															}
															position++
															goto l1118
														l1119:
															position, tokenIndex = position1118, tokenIndex1118
															if buffer[position] != rune('A') {
																goto l1116
															}
															position++
														}
													l1118:
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
																goto l1116
															}
															position++
														}
													l1120:
														{
															position1122, tokenIndex1122 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1123
															}
															position++
															goto l1122
														l1123:
															position, tokenIndex = position1122, tokenIndex1122
															if buffer[position] != rune('C') {
																goto l1116
															}
															position++
														}
													l1122:
														if !_rules[rulews]() {
															goto l1116
														}
														{
															position1124, tokenIndex1124 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1125
															}
															position++
															goto l1124
														l1125:
															position, tokenIndex = position1124, tokenIndex1124
															if buffer[position] != rune('A') {
																goto l1116
															}
															position++
														}
													l1124:
														if !_rules[rulesep]() {
															goto l1116
														}
														if !_rules[ruleSrc8]() {
															goto l1116
														}
														{
															add(ruleAction44, position)
														}
														add(ruleAdc, position1117)
													}
													goto l1104
												l1116:
													position, tokenIndex = position1104, tokenIndex1104
													{
														position1128 := position
														{
															position1129, tokenIndex1129 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1130
															}
															position++
															goto l1129
														l1130:
															position, tokenIndex = position1129, tokenIndex1129
															if buffer[position] != rune('S') {
																goto l1127
															}
															position++
														}
													l1129:
														{
															position1131, tokenIndex1131 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1132
															}
															position++
															goto l1131
														l1132:
															position, tokenIndex = position1131, tokenIndex1131
															if buffer[position] != rune('U') {
																goto l1127
															}
															position++
														}
													l1131:
														{
															position1133, tokenIndex1133 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1134
															}
															position++
															goto l1133
														l1134:
															position, tokenIndex = position1133, tokenIndex1133
															if buffer[position] != rune('B') {
																goto l1127
															}
															position++
														}
													l1133:
														if !_rules[rulews]() {
															goto l1127
														}
														if !_rules[ruleSrc8]() {
															goto l1127
														}
														{
															add(ruleAction45, position)
														}
														add(ruleSub, position1128)
													}
													goto l1104
												l1127:
													position, tokenIndex = position1104, tokenIndex1104
													{
														switch buffer[position] {
														case 'C', 'c':
															{
																position1137 := position
																{
																	position1138, tokenIndex1138 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1139
																	}
																	position++
																	goto l1138
																l1139:
																	position, tokenIndex = position1138, tokenIndex1138
																	if buffer[position] != rune('C') {
																		goto l1102
																	}
																	position++
																}
															l1138:
																{
																	position1140, tokenIndex1140 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1141
																	}
																	position++
																	goto l1140
																l1141:
																	position, tokenIndex = position1140, tokenIndex1140
																	if buffer[position] != rune('P') {
																		goto l1102
																	}
																	position++
																}
															l1140:
																if !_rules[rulews]() {
																	goto l1102
																}
																if !_rules[ruleSrc8]() {
																	goto l1102
																}
																{
																	add(ruleAction50, position)
																}
																add(ruleCp, position1137)
															}
															break
														case 'O', 'o':
															{
																position1143 := position
																{
																	position1144, tokenIndex1144 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1145
																	}
																	position++
																	goto l1144
																l1145:
																	position, tokenIndex = position1144, tokenIndex1144
																	if buffer[position] != rune('O') {
																		goto l1102
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
																		goto l1102
																	}
																	position++
																}
															l1146:
																if !_rules[rulews]() {
																	goto l1102
																}
																if !_rules[ruleSrc8]() {
																	goto l1102
																}
																{
																	add(ruleAction49, position)
																}
																add(ruleOr, position1143)
															}
															break
														case 'X', 'x':
															{
																position1149 := position
																{
																	position1150, tokenIndex1150 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1151
																	}
																	position++
																	goto l1150
																l1151:
																	position, tokenIndex = position1150, tokenIndex1150
																	if buffer[position] != rune('X') {
																		goto l1102
																	}
																	position++
																}
															l1150:
																{
																	position1152, tokenIndex1152 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1153
																	}
																	position++
																	goto l1152
																l1153:
																	position, tokenIndex = position1152, tokenIndex1152
																	if buffer[position] != rune('O') {
																		goto l1102
																	}
																	position++
																}
															l1152:
																{
																	position1154, tokenIndex1154 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1155
																	}
																	position++
																	goto l1154
																l1155:
																	position, tokenIndex = position1154, tokenIndex1154
																	if buffer[position] != rune('R') {
																		goto l1102
																	}
																	position++
																}
															l1154:
																if !_rules[rulews]() {
																	goto l1102
																}
																if !_rules[ruleSrc8]() {
																	goto l1102
																}
																{
																	add(ruleAction48, position)
																}
																add(ruleXor, position1149)
															}
															break
														case 'A', 'a':
															{
																position1157 := position
																{
																	position1158, tokenIndex1158 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1159
																	}
																	position++
																	goto l1158
																l1159:
																	position, tokenIndex = position1158, tokenIndex1158
																	if buffer[position] != rune('A') {
																		goto l1102
																	}
																	position++
																}
															l1158:
																{
																	position1160, tokenIndex1160 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1161
																	}
																	position++
																	goto l1160
																l1161:
																	position, tokenIndex = position1160, tokenIndex1160
																	if buffer[position] != rune('N') {
																		goto l1102
																	}
																	position++
																}
															l1160:
																{
																	position1162, tokenIndex1162 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1163
																	}
																	position++
																	goto l1162
																l1163:
																	position, tokenIndex = position1162, tokenIndex1162
																	if buffer[position] != rune('D') {
																		goto l1102
																	}
																	position++
																}
															l1162:
																if !_rules[rulews]() {
																	goto l1102
																}
																if !_rules[ruleSrc8]() {
																	goto l1102
																}
																{
																	add(ruleAction47, position)
																}
																add(ruleAnd, position1157)
															}
															break
														default:
															{
																position1165 := position
																{
																	position1166, tokenIndex1166 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1167
																	}
																	position++
																	goto l1166
																l1167:
																	position, tokenIndex = position1166, tokenIndex1166
																	if buffer[position] != rune('S') {
																		goto l1102
																	}
																	position++
																}
															l1166:
																{
																	position1168, tokenIndex1168 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1169
																	}
																	position++
																	goto l1168
																l1169:
																	position, tokenIndex = position1168, tokenIndex1168
																	if buffer[position] != rune('B') {
																		goto l1102
																	}
																	position++
																}
															l1168:
																{
																	position1170, tokenIndex1170 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1171
																	}
																	position++
																	goto l1170
																l1171:
																	position, tokenIndex = position1170, tokenIndex1170
																	if buffer[position] != rune('C') {
																		goto l1102
																	}
																	position++
																}
															l1170:
																if !_rules[rulews]() {
																	goto l1102
																}
																{
																	position1172, tokenIndex1172 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1173
																	}
																	position++
																	goto l1172
																l1173:
																	position, tokenIndex = position1172, tokenIndex1172
																	if buffer[position] != rune('A') {
																		goto l1102
																	}
																	position++
																}
															l1172:
																if !_rules[rulesep]() {
																	goto l1102
																}
																if !_rules[ruleSrc8]() {
																	goto l1102
																}
																{
																	add(ruleAction46, position)
																}
																add(ruleSbc, position1165)
															}
															break
														}
													}

												}
											l1104:
												add(ruleAlu, position1103)
											}
											goto l965
										l1102:
											position, tokenIndex = position965, tokenIndex965
											{
												position1176 := position
												{
													position1177, tokenIndex1177 := position, tokenIndex
													{
														position1179 := position
														{
															position1180, tokenIndex1180 := position, tokenIndex
															{
																position1182 := position
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
																		goto l1181
																	}
																	position++
																}
															l1183:
																{
																	position1185, tokenIndex1185 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1186
																	}
																	position++
																	goto l1185
																l1186:
																	position, tokenIndex = position1185, tokenIndex1185
																	if buffer[position] != rune('L') {
																		goto l1181
																	}
																	position++
																}
															l1185:
																{
																	position1187, tokenIndex1187 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1188
																	}
																	position++
																	goto l1187
																l1188:
																	position, tokenIndex = position1187, tokenIndex1187
																	if buffer[position] != rune('C') {
																		goto l1181
																	}
																	position++
																}
															l1187:
																if !_rules[rulews]() {
																	goto l1181
																}
																if !_rules[ruleLoc8]() {
																	goto l1181
																}
																{
																	position1189, tokenIndex1189 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1189
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1189
																	}
																	goto l1190
																l1189:
																	position, tokenIndex = position1189, tokenIndex1189
																}
															l1190:
																{
																	add(ruleAction51, position)
																}
																add(ruleRlc, position1182)
															}
															goto l1180
														l1181:
															position, tokenIndex = position1180, tokenIndex1180
															{
																position1193 := position
																{
																	position1194, tokenIndex1194 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1195
																	}
																	position++
																	goto l1194
																l1195:
																	position, tokenIndex = position1194, tokenIndex1194
																	if buffer[position] != rune('R') {
																		goto l1192
																	}
																	position++
																}
															l1194:
																{
																	position1196, tokenIndex1196 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1197
																	}
																	position++
																	goto l1196
																l1197:
																	position, tokenIndex = position1196, tokenIndex1196
																	if buffer[position] != rune('R') {
																		goto l1192
																	}
																	position++
																}
															l1196:
																{
																	position1198, tokenIndex1198 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1199
																	}
																	position++
																	goto l1198
																l1199:
																	position, tokenIndex = position1198, tokenIndex1198
																	if buffer[position] != rune('C') {
																		goto l1192
																	}
																	position++
																}
															l1198:
																if !_rules[rulews]() {
																	goto l1192
																}
																if !_rules[ruleLoc8]() {
																	goto l1192
																}
																{
																	position1200, tokenIndex1200 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1200
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1200
																	}
																	goto l1201
																l1200:
																	position, tokenIndex = position1200, tokenIndex1200
																}
															l1201:
																{
																	add(ruleAction52, position)
																}
																add(ruleRrc, position1193)
															}
															goto l1180
														l1192:
															position, tokenIndex = position1180, tokenIndex1180
															{
																position1204 := position
																{
																	position1205, tokenIndex1205 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1206
																	}
																	position++
																	goto l1205
																l1206:
																	position, tokenIndex = position1205, tokenIndex1205
																	if buffer[position] != rune('R') {
																		goto l1203
																	}
																	position++
																}
															l1205:
																{
																	position1207, tokenIndex1207 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1208
																	}
																	position++
																	goto l1207
																l1208:
																	position, tokenIndex = position1207, tokenIndex1207
																	if buffer[position] != rune('L') {
																		goto l1203
																	}
																	position++
																}
															l1207:
																if !_rules[rulews]() {
																	goto l1203
																}
																if !_rules[ruleLoc8]() {
																	goto l1203
																}
																{
																	position1209, tokenIndex1209 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1209
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1209
																	}
																	goto l1210
																l1209:
																	position, tokenIndex = position1209, tokenIndex1209
																}
															l1210:
																{
																	add(ruleAction53, position)
																}
																add(ruleRl, position1204)
															}
															goto l1180
														l1203:
															position, tokenIndex = position1180, tokenIndex1180
															{
																position1213 := position
																{
																	position1214, tokenIndex1214 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1215
																	}
																	position++
																	goto l1214
																l1215:
																	position, tokenIndex = position1214, tokenIndex1214
																	if buffer[position] != rune('R') {
																		goto l1212
																	}
																	position++
																}
															l1214:
																{
																	position1216, tokenIndex1216 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1217
																	}
																	position++
																	goto l1216
																l1217:
																	position, tokenIndex = position1216, tokenIndex1216
																	if buffer[position] != rune('R') {
																		goto l1212
																	}
																	position++
																}
															l1216:
																if !_rules[rulews]() {
																	goto l1212
																}
																if !_rules[ruleLoc8]() {
																	goto l1212
																}
																{
																	position1218, tokenIndex1218 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1218
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1218
																	}
																	goto l1219
																l1218:
																	position, tokenIndex = position1218, tokenIndex1218
																}
															l1219:
																{
																	add(ruleAction54, position)
																}
																add(ruleRr, position1213)
															}
															goto l1180
														l1212:
															position, tokenIndex = position1180, tokenIndex1180
															{
																position1222 := position
																{
																	position1223, tokenIndex1223 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1224
																	}
																	position++
																	goto l1223
																l1224:
																	position, tokenIndex = position1223, tokenIndex1223
																	if buffer[position] != rune('S') {
																		goto l1221
																	}
																	position++
																}
															l1223:
																{
																	position1225, tokenIndex1225 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1226
																	}
																	position++
																	goto l1225
																l1226:
																	position, tokenIndex = position1225, tokenIndex1225
																	if buffer[position] != rune('L') {
																		goto l1221
																	}
																	position++
																}
															l1225:
																{
																	position1227, tokenIndex1227 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1228
																	}
																	position++
																	goto l1227
																l1228:
																	position, tokenIndex = position1227, tokenIndex1227
																	if buffer[position] != rune('A') {
																		goto l1221
																	}
																	position++
																}
															l1227:
																if !_rules[rulews]() {
																	goto l1221
																}
																if !_rules[ruleLoc8]() {
																	goto l1221
																}
																{
																	position1229, tokenIndex1229 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1229
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1229
																	}
																	goto l1230
																l1229:
																	position, tokenIndex = position1229, tokenIndex1229
																}
															l1230:
																{
																	add(ruleAction55, position)
																}
																add(ruleSla, position1222)
															}
															goto l1180
														l1221:
															position, tokenIndex = position1180, tokenIndex1180
															{
																position1233 := position
																{
																	position1234, tokenIndex1234 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1235
																	}
																	position++
																	goto l1234
																l1235:
																	position, tokenIndex = position1234, tokenIndex1234
																	if buffer[position] != rune('S') {
																		goto l1232
																	}
																	position++
																}
															l1234:
																{
																	position1236, tokenIndex1236 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1237
																	}
																	position++
																	goto l1236
																l1237:
																	position, tokenIndex = position1236, tokenIndex1236
																	if buffer[position] != rune('R') {
																		goto l1232
																	}
																	position++
																}
															l1236:
																{
																	position1238, tokenIndex1238 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1239
																	}
																	position++
																	goto l1238
																l1239:
																	position, tokenIndex = position1238, tokenIndex1238
																	if buffer[position] != rune('A') {
																		goto l1232
																	}
																	position++
																}
															l1238:
																if !_rules[rulews]() {
																	goto l1232
																}
																if !_rules[ruleLoc8]() {
																	goto l1232
																}
																{
																	position1240, tokenIndex1240 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1240
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1240
																	}
																	goto l1241
																l1240:
																	position, tokenIndex = position1240, tokenIndex1240
																}
															l1241:
																{
																	add(ruleAction56, position)
																}
																add(ruleSra, position1233)
															}
															goto l1180
														l1232:
															position, tokenIndex = position1180, tokenIndex1180
															{
																position1244 := position
																{
																	position1245, tokenIndex1245 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1246
																	}
																	position++
																	goto l1245
																l1246:
																	position, tokenIndex = position1245, tokenIndex1245
																	if buffer[position] != rune('S') {
																		goto l1243
																	}
																	position++
																}
															l1245:
																{
																	position1247, tokenIndex1247 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1248
																	}
																	position++
																	goto l1247
																l1248:
																	position, tokenIndex = position1247, tokenIndex1247
																	if buffer[position] != rune('L') {
																		goto l1243
																	}
																	position++
																}
															l1247:
																{
																	position1249, tokenIndex1249 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1250
																	}
																	position++
																	goto l1249
																l1250:
																	position, tokenIndex = position1249, tokenIndex1249
																	if buffer[position] != rune('L') {
																		goto l1243
																	}
																	position++
																}
															l1249:
																if !_rules[rulews]() {
																	goto l1243
																}
																if !_rules[ruleLoc8]() {
																	goto l1243
																}
																{
																	position1251, tokenIndex1251 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1251
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1251
																	}
																	goto l1252
																l1251:
																	position, tokenIndex = position1251, tokenIndex1251
																}
															l1252:
																{
																	add(ruleAction57, position)
																}
																add(ruleSll, position1244)
															}
															goto l1180
														l1243:
															position, tokenIndex = position1180, tokenIndex1180
															{
																position1254 := position
																{
																	position1255, tokenIndex1255 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1256
																	}
																	position++
																	goto l1255
																l1256:
																	position, tokenIndex = position1255, tokenIndex1255
																	if buffer[position] != rune('S') {
																		goto l1178
																	}
																	position++
																}
															l1255:
																{
																	position1257, tokenIndex1257 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1258
																	}
																	position++
																	goto l1257
																l1258:
																	position, tokenIndex = position1257, tokenIndex1257
																	if buffer[position] != rune('R') {
																		goto l1178
																	}
																	position++
																}
															l1257:
																{
																	position1259, tokenIndex1259 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1260
																	}
																	position++
																	goto l1259
																l1260:
																	position, tokenIndex = position1259, tokenIndex1259
																	if buffer[position] != rune('L') {
																		goto l1178
																	}
																	position++
																}
															l1259:
																if !_rules[rulews]() {
																	goto l1178
																}
																if !_rules[ruleLoc8]() {
																	goto l1178
																}
																{
																	position1261, tokenIndex1261 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1261
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1261
																	}
																	goto l1262
																l1261:
																	position, tokenIndex = position1261, tokenIndex1261
																}
															l1262:
																{
																	add(ruleAction58, position)
																}
																add(ruleSrl, position1254)
															}
														}
													l1180:
														add(ruleRot, position1179)
													}
													goto l1177
												l1178:
													position, tokenIndex = position1177, tokenIndex1177
													{
														switch buffer[position] {
														case 'S', 's':
															{
																position1265 := position
																{
																	position1266, tokenIndex1266 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1267
																	}
																	position++
																	goto l1266
																l1267:
																	position, tokenIndex = position1266, tokenIndex1266
																	if buffer[position] != rune('S') {
																		goto l1175
																	}
																	position++
																}
															l1266:
																{
																	position1268, tokenIndex1268 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1269
																	}
																	position++
																	goto l1268
																l1269:
																	position, tokenIndex = position1268, tokenIndex1268
																	if buffer[position] != rune('E') {
																		goto l1175
																	}
																	position++
																}
															l1268:
																{
																	position1270, tokenIndex1270 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1271
																	}
																	position++
																	goto l1270
																l1271:
																	position, tokenIndex = position1270, tokenIndex1270
																	if buffer[position] != rune('T') {
																		goto l1175
																	}
																	position++
																}
															l1270:
																if !_rules[rulews]() {
																	goto l1175
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1175
																}
																if !_rules[rulesep]() {
																	goto l1175
																}
																if !_rules[ruleLoc8]() {
																	goto l1175
																}
																{
																	position1272, tokenIndex1272 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1272
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1272
																	}
																	goto l1273
																l1272:
																	position, tokenIndex = position1272, tokenIndex1272
																}
															l1273:
																{
																	add(ruleAction61, position)
																}
																add(ruleSet, position1265)
															}
															break
														case 'R', 'r':
															{
																position1275 := position
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
																		goto l1175
																	}
																	position++
																}
															l1276:
																{
																	position1278, tokenIndex1278 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1279
																	}
																	position++
																	goto l1278
																l1279:
																	position, tokenIndex = position1278, tokenIndex1278
																	if buffer[position] != rune('E') {
																		goto l1175
																	}
																	position++
																}
															l1278:
																{
																	position1280, tokenIndex1280 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1281
																	}
																	position++
																	goto l1280
																l1281:
																	position, tokenIndex = position1280, tokenIndex1280
																	if buffer[position] != rune('S') {
																		goto l1175
																	}
																	position++
																}
															l1280:
																if !_rules[rulews]() {
																	goto l1175
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1175
																}
																if !_rules[rulesep]() {
																	goto l1175
																}
																if !_rules[ruleLoc8]() {
																	goto l1175
																}
																{
																	position1282, tokenIndex1282 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1282
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1282
																	}
																	goto l1283
																l1282:
																	position, tokenIndex = position1282, tokenIndex1282
																}
															l1283:
																{
																	add(ruleAction60, position)
																}
																add(ruleRes, position1275)
															}
															break
														default:
															{
																position1285 := position
																{
																	position1286, tokenIndex1286 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1287
																	}
																	position++
																	goto l1286
																l1287:
																	position, tokenIndex = position1286, tokenIndex1286
																	if buffer[position] != rune('B') {
																		goto l1175
																	}
																	position++
																}
															l1286:
																{
																	position1288, tokenIndex1288 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l1289
																	}
																	position++
																	goto l1288
																l1289:
																	position, tokenIndex = position1288, tokenIndex1288
																	if buffer[position] != rune('I') {
																		goto l1175
																	}
																	position++
																}
															l1288:
																{
																	position1290, tokenIndex1290 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1291
																	}
																	position++
																	goto l1290
																l1291:
																	position, tokenIndex = position1290, tokenIndex1290
																	if buffer[position] != rune('T') {
																		goto l1175
																	}
																	position++
																}
															l1290:
																if !_rules[rulews]() {
																	goto l1175
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1175
																}
																if !_rules[rulesep]() {
																	goto l1175
																}
																if !_rules[ruleLoc8]() {
																	goto l1175
																}
																{
																	add(ruleAction59, position)
																}
																add(ruleBit, position1285)
															}
															break
														}
													}

												}
											l1177:
												add(ruleBitOp, position1176)
											}
											goto l965
										l1175:
											position, tokenIndex = position965, tokenIndex965
											{
												position1294 := position
												{
													position1295, tokenIndex1295 := position, tokenIndex
													{
														position1297 := position
														{
															position1298 := position
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
																	goto l1296
																}
																position++
															}
														l1299:
															{
																position1301, tokenIndex1301 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1302
																}
																position++
																goto l1301
															l1302:
																position, tokenIndex = position1301, tokenIndex1301
																if buffer[position] != rune('E') {
																	goto l1296
																}
																position++
															}
														l1301:
															{
																position1303, tokenIndex1303 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1304
																}
																position++
																goto l1303
															l1304:
																position, tokenIndex = position1303, tokenIndex1303
																if buffer[position] != rune('T') {
																	goto l1296
																}
																position++
															}
														l1303:
															{
																position1305, tokenIndex1305 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1306
																}
																position++
																goto l1305
															l1306:
																position, tokenIndex = position1305, tokenIndex1305
																if buffer[position] != rune('N') {
																	goto l1296
																}
																position++
															}
														l1305:
															add(rulePegText, position1298)
														}
														{
															add(ruleAction76, position)
														}
														add(ruleRetn, position1297)
													}
													goto l1295
												l1296:
													position, tokenIndex = position1295, tokenIndex1295
													{
														position1309 := position
														{
															position1310 := position
															{
																position1311, tokenIndex1311 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1312
																}
																position++
																goto l1311
															l1312:
																position, tokenIndex = position1311, tokenIndex1311
																if buffer[position] != rune('R') {
																	goto l1308
																}
																position++
															}
														l1311:
															{
																position1313, tokenIndex1313 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1314
																}
																position++
																goto l1313
															l1314:
																position, tokenIndex = position1313, tokenIndex1313
																if buffer[position] != rune('E') {
																	goto l1308
																}
																position++
															}
														l1313:
															{
																position1315, tokenIndex1315 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1316
																}
																position++
																goto l1315
															l1316:
																position, tokenIndex = position1315, tokenIndex1315
																if buffer[position] != rune('T') {
																	goto l1308
																}
																position++
															}
														l1315:
															{
																position1317, tokenIndex1317 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1318
																}
																position++
																goto l1317
															l1318:
																position, tokenIndex = position1317, tokenIndex1317
																if buffer[position] != rune('I') {
																	goto l1308
																}
																position++
															}
														l1317:
															add(rulePegText, position1310)
														}
														{
															add(ruleAction77, position)
														}
														add(ruleReti, position1309)
													}
													goto l1295
												l1308:
													position, tokenIndex = position1295, tokenIndex1295
													{
														position1321 := position
														{
															position1322 := position
															{
																position1323, tokenIndex1323 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1324
																}
																position++
																goto l1323
															l1324:
																position, tokenIndex = position1323, tokenIndex1323
																if buffer[position] != rune('R') {
																	goto l1320
																}
																position++
															}
														l1323:
															{
																position1325, tokenIndex1325 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1326
																}
																position++
																goto l1325
															l1326:
																position, tokenIndex = position1325, tokenIndex1325
																if buffer[position] != rune('R') {
																	goto l1320
																}
																position++
															}
														l1325:
															{
																position1327, tokenIndex1327 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1328
																}
																position++
																goto l1327
															l1328:
																position, tokenIndex = position1327, tokenIndex1327
																if buffer[position] != rune('D') {
																	goto l1320
																}
																position++
															}
														l1327:
															add(rulePegText, position1322)
														}
														{
															add(ruleAction78, position)
														}
														add(ruleRrd, position1321)
													}
													goto l1295
												l1320:
													position, tokenIndex = position1295, tokenIndex1295
													{
														position1331 := position
														{
															position1332 := position
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
																	goto l1330
																}
																position++
															}
														l1333:
															{
																position1335, tokenIndex1335 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1336
																}
																position++
																goto l1335
															l1336:
																position, tokenIndex = position1335, tokenIndex1335
																if buffer[position] != rune('M') {
																	goto l1330
																}
																position++
															}
														l1335:
															if buffer[position] != rune(' ') {
																goto l1330
															}
															position++
															if buffer[position] != rune('0') {
																goto l1330
															}
															position++
															add(rulePegText, position1332)
														}
														{
															add(ruleAction80, position)
														}
														add(ruleIm0, position1331)
													}
													goto l1295
												l1330:
													position, tokenIndex = position1295, tokenIndex1295
													{
														position1339 := position
														{
															position1340 := position
															{
																position1341, tokenIndex1341 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1342
																}
																position++
																goto l1341
															l1342:
																position, tokenIndex = position1341, tokenIndex1341
																if buffer[position] != rune('I') {
																	goto l1338
																}
																position++
															}
														l1341:
															{
																position1343, tokenIndex1343 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1344
																}
																position++
																goto l1343
															l1344:
																position, tokenIndex = position1343, tokenIndex1343
																if buffer[position] != rune('M') {
																	goto l1338
																}
																position++
															}
														l1343:
															if buffer[position] != rune(' ') {
																goto l1338
															}
															position++
															if buffer[position] != rune('1') {
																goto l1338
															}
															position++
															add(rulePegText, position1340)
														}
														{
															add(ruleAction81, position)
														}
														add(ruleIm1, position1339)
													}
													goto l1295
												l1338:
													position, tokenIndex = position1295, tokenIndex1295
													{
														position1347 := position
														{
															position1348 := position
															{
																position1349, tokenIndex1349 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1350
																}
																position++
																goto l1349
															l1350:
																position, tokenIndex = position1349, tokenIndex1349
																if buffer[position] != rune('I') {
																	goto l1346
																}
																position++
															}
														l1349:
															{
																position1351, tokenIndex1351 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1352
																}
																position++
																goto l1351
															l1352:
																position, tokenIndex = position1351, tokenIndex1351
																if buffer[position] != rune('M') {
																	goto l1346
																}
																position++
															}
														l1351:
															if buffer[position] != rune(' ') {
																goto l1346
															}
															position++
															if buffer[position] != rune('2') {
																goto l1346
															}
															position++
															add(rulePegText, position1348)
														}
														{
															add(ruleAction82, position)
														}
														add(ruleIm2, position1347)
													}
													goto l1295
												l1346:
													position, tokenIndex = position1295, tokenIndex1295
													{
														switch buffer[position] {
														case 'I', 'O', 'i', 'o':
															{
																position1355 := position
																{
																	position1356, tokenIndex1356 := position, tokenIndex
																	{
																		position1358 := position
																		{
																			position1359 := position
																			{
																				position1360, tokenIndex1360 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1361
																				}
																				position++
																				goto l1360
																			l1361:
																				position, tokenIndex = position1360, tokenIndex1360
																				if buffer[position] != rune('I') {
																					goto l1357
																				}
																				position++
																			}
																		l1360:
																			{
																				position1362, tokenIndex1362 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1363
																				}
																				position++
																				goto l1362
																			l1363:
																				position, tokenIndex = position1362, tokenIndex1362
																				if buffer[position] != rune('N') {
																					goto l1357
																				}
																				position++
																			}
																		l1362:
																			{
																				position1364, tokenIndex1364 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1365
																				}
																				position++
																				goto l1364
																			l1365:
																				position, tokenIndex = position1364, tokenIndex1364
																				if buffer[position] != rune('I') {
																					goto l1357
																				}
																				position++
																			}
																		l1364:
																			{
																				position1366, tokenIndex1366 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1367
																				}
																				position++
																				goto l1366
																			l1367:
																				position, tokenIndex = position1366, tokenIndex1366
																				if buffer[position] != rune('R') {
																					goto l1357
																				}
																				position++
																			}
																		l1366:
																			add(rulePegText, position1359)
																		}
																		{
																			add(ruleAction93, position)
																		}
																		add(ruleInir, position1358)
																	}
																	goto l1356
																l1357:
																	position, tokenIndex = position1356, tokenIndex1356
																	{
																		position1370 := position
																		{
																			position1371 := position
																			{
																				position1372, tokenIndex1372 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1373
																				}
																				position++
																				goto l1372
																			l1373:
																				position, tokenIndex = position1372, tokenIndex1372
																				if buffer[position] != rune('I') {
																					goto l1369
																				}
																				position++
																			}
																		l1372:
																			{
																				position1374, tokenIndex1374 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1375
																				}
																				position++
																				goto l1374
																			l1375:
																				position, tokenIndex = position1374, tokenIndex1374
																				if buffer[position] != rune('N') {
																					goto l1369
																				}
																				position++
																			}
																		l1374:
																			{
																				position1376, tokenIndex1376 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1377
																				}
																				position++
																				goto l1376
																			l1377:
																				position, tokenIndex = position1376, tokenIndex1376
																				if buffer[position] != rune('I') {
																					goto l1369
																				}
																				position++
																			}
																		l1376:
																			add(rulePegText, position1371)
																		}
																		{
																			add(ruleAction85, position)
																		}
																		add(ruleIni, position1370)
																	}
																	goto l1356
																l1369:
																	position, tokenIndex = position1356, tokenIndex1356
																	{
																		position1380 := position
																		{
																			position1381 := position
																			{
																				position1382, tokenIndex1382 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1383
																				}
																				position++
																				goto l1382
																			l1383:
																				position, tokenIndex = position1382, tokenIndex1382
																				if buffer[position] != rune('O') {
																					goto l1379
																				}
																				position++
																			}
																		l1382:
																			{
																				position1384, tokenIndex1384 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1385
																				}
																				position++
																				goto l1384
																			l1385:
																				position, tokenIndex = position1384, tokenIndex1384
																				if buffer[position] != rune('T') {
																					goto l1379
																				}
																				position++
																			}
																		l1384:
																			{
																				position1386, tokenIndex1386 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1387
																				}
																				position++
																				goto l1386
																			l1387:
																				position, tokenIndex = position1386, tokenIndex1386
																				if buffer[position] != rune('I') {
																					goto l1379
																				}
																				position++
																			}
																		l1386:
																			{
																				position1388, tokenIndex1388 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1389
																				}
																				position++
																				goto l1388
																			l1389:
																				position, tokenIndex = position1388, tokenIndex1388
																				if buffer[position] != rune('R') {
																					goto l1379
																				}
																				position++
																			}
																		l1388:
																			add(rulePegText, position1381)
																		}
																		{
																			add(ruleAction94, position)
																		}
																		add(ruleOtir, position1380)
																	}
																	goto l1356
																l1379:
																	position, tokenIndex = position1356, tokenIndex1356
																	{
																		position1392 := position
																		{
																			position1393 := position
																			{
																				position1394, tokenIndex1394 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1395
																				}
																				position++
																				goto l1394
																			l1395:
																				position, tokenIndex = position1394, tokenIndex1394
																				if buffer[position] != rune('O') {
																					goto l1391
																				}
																				position++
																			}
																		l1394:
																			{
																				position1396, tokenIndex1396 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1397
																				}
																				position++
																				goto l1396
																			l1397:
																				position, tokenIndex = position1396, tokenIndex1396
																				if buffer[position] != rune('U') {
																					goto l1391
																				}
																				position++
																			}
																		l1396:
																			{
																				position1398, tokenIndex1398 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1399
																				}
																				position++
																				goto l1398
																			l1399:
																				position, tokenIndex = position1398, tokenIndex1398
																				if buffer[position] != rune('T') {
																					goto l1391
																				}
																				position++
																			}
																		l1398:
																			{
																				position1400, tokenIndex1400 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1401
																				}
																				position++
																				goto l1400
																			l1401:
																				position, tokenIndex = position1400, tokenIndex1400
																				if buffer[position] != rune('I') {
																					goto l1391
																				}
																				position++
																			}
																		l1400:
																			add(rulePegText, position1393)
																		}
																		{
																			add(ruleAction86, position)
																		}
																		add(ruleOuti, position1392)
																	}
																	goto l1356
																l1391:
																	position, tokenIndex = position1356, tokenIndex1356
																	{
																		position1404 := position
																		{
																			position1405 := position
																			{
																				position1406, tokenIndex1406 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1407
																				}
																				position++
																				goto l1406
																			l1407:
																				position, tokenIndex = position1406, tokenIndex1406
																				if buffer[position] != rune('I') {
																					goto l1403
																				}
																				position++
																			}
																		l1406:
																			{
																				position1408, tokenIndex1408 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1409
																				}
																				position++
																				goto l1408
																			l1409:
																				position, tokenIndex = position1408, tokenIndex1408
																				if buffer[position] != rune('N') {
																					goto l1403
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
																					goto l1403
																				}
																				position++
																			}
																		l1410:
																			{
																				position1412, tokenIndex1412 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1413
																				}
																				position++
																				goto l1412
																			l1413:
																				position, tokenIndex = position1412, tokenIndex1412
																				if buffer[position] != rune('R') {
																					goto l1403
																				}
																				position++
																			}
																		l1412:
																			add(rulePegText, position1405)
																		}
																		{
																			add(ruleAction97, position)
																		}
																		add(ruleIndr, position1404)
																	}
																	goto l1356
																l1403:
																	position, tokenIndex = position1356, tokenIndex1356
																	{
																		position1416 := position
																		{
																			position1417 := position
																			{
																				position1418, tokenIndex1418 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1419
																				}
																				position++
																				goto l1418
																			l1419:
																				position, tokenIndex = position1418, tokenIndex1418
																				if buffer[position] != rune('I') {
																					goto l1415
																				}
																				position++
																			}
																		l1418:
																			{
																				position1420, tokenIndex1420 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1421
																				}
																				position++
																				goto l1420
																			l1421:
																				position, tokenIndex = position1420, tokenIndex1420
																				if buffer[position] != rune('N') {
																					goto l1415
																				}
																				position++
																			}
																		l1420:
																			{
																				position1422, tokenIndex1422 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1423
																				}
																				position++
																				goto l1422
																			l1423:
																				position, tokenIndex = position1422, tokenIndex1422
																				if buffer[position] != rune('D') {
																					goto l1415
																				}
																				position++
																			}
																		l1422:
																			add(rulePegText, position1417)
																		}
																		{
																			add(ruleAction89, position)
																		}
																		add(ruleInd, position1416)
																	}
																	goto l1356
																l1415:
																	position, tokenIndex = position1356, tokenIndex1356
																	{
																		position1426 := position
																		{
																			position1427 := position
																			{
																				position1428, tokenIndex1428 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1429
																				}
																				position++
																				goto l1428
																			l1429:
																				position, tokenIndex = position1428, tokenIndex1428
																				if buffer[position] != rune('O') {
																					goto l1425
																				}
																				position++
																			}
																		l1428:
																			{
																				position1430, tokenIndex1430 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1431
																				}
																				position++
																				goto l1430
																			l1431:
																				position, tokenIndex = position1430, tokenIndex1430
																				if buffer[position] != rune('T') {
																					goto l1425
																				}
																				position++
																			}
																		l1430:
																			{
																				position1432, tokenIndex1432 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1433
																				}
																				position++
																				goto l1432
																			l1433:
																				position, tokenIndex = position1432, tokenIndex1432
																				if buffer[position] != rune('D') {
																					goto l1425
																				}
																				position++
																			}
																		l1432:
																			{
																				position1434, tokenIndex1434 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1435
																				}
																				position++
																				goto l1434
																			l1435:
																				position, tokenIndex = position1434, tokenIndex1434
																				if buffer[position] != rune('R') {
																					goto l1425
																				}
																				position++
																			}
																		l1434:
																			add(rulePegText, position1427)
																		}
																		{
																			add(ruleAction98, position)
																		}
																		add(ruleOtdr, position1426)
																	}
																	goto l1356
																l1425:
																	position, tokenIndex = position1356, tokenIndex1356
																	{
																		position1437 := position
																		{
																			position1438 := position
																			{
																				position1439, tokenIndex1439 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1440
																				}
																				position++
																				goto l1439
																			l1440:
																				position, tokenIndex = position1439, tokenIndex1439
																				if buffer[position] != rune('O') {
																					goto l1293
																				}
																				position++
																			}
																		l1439:
																			{
																				position1441, tokenIndex1441 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1442
																				}
																				position++
																				goto l1441
																			l1442:
																				position, tokenIndex = position1441, tokenIndex1441
																				if buffer[position] != rune('U') {
																					goto l1293
																				}
																				position++
																			}
																		l1441:
																			{
																				position1443, tokenIndex1443 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1444
																				}
																				position++
																				goto l1443
																			l1444:
																				position, tokenIndex = position1443, tokenIndex1443
																				if buffer[position] != rune('T') {
																					goto l1293
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
																					goto l1293
																				}
																				position++
																			}
																		l1445:
																			add(rulePegText, position1438)
																		}
																		{
																			add(ruleAction90, position)
																		}
																		add(ruleOutd, position1437)
																	}
																}
															l1356:
																add(ruleBlitIO, position1355)
															}
															break
														case 'R', 'r':
															{
																position1448 := position
																{
																	position1449 := position
																	{
																		position1450, tokenIndex1450 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1451
																		}
																		position++
																		goto l1450
																	l1451:
																		position, tokenIndex = position1450, tokenIndex1450
																		if buffer[position] != rune('R') {
																			goto l1293
																		}
																		position++
																	}
																l1450:
																	{
																		position1452, tokenIndex1452 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1453
																		}
																		position++
																		goto l1452
																	l1453:
																		position, tokenIndex = position1452, tokenIndex1452
																		if buffer[position] != rune('L') {
																			goto l1293
																		}
																		position++
																	}
																l1452:
																	{
																		position1454, tokenIndex1454 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1455
																		}
																		position++
																		goto l1454
																	l1455:
																		position, tokenIndex = position1454, tokenIndex1454
																		if buffer[position] != rune('D') {
																			goto l1293
																		}
																		position++
																	}
																l1454:
																	add(rulePegText, position1449)
																}
																{
																	add(ruleAction79, position)
																}
																add(ruleRld, position1448)
															}
															break
														case 'N', 'n':
															{
																position1457 := position
																{
																	position1458 := position
																	{
																		position1459, tokenIndex1459 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1460
																		}
																		position++
																		goto l1459
																	l1460:
																		position, tokenIndex = position1459, tokenIndex1459
																		if buffer[position] != rune('N') {
																			goto l1293
																		}
																		position++
																	}
																l1459:
																	{
																		position1461, tokenIndex1461 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1462
																		}
																		position++
																		goto l1461
																	l1462:
																		position, tokenIndex = position1461, tokenIndex1461
																		if buffer[position] != rune('E') {
																			goto l1293
																		}
																		position++
																	}
																l1461:
																	{
																		position1463, tokenIndex1463 := position, tokenIndex
																		if buffer[position] != rune('g') {
																			goto l1464
																		}
																		position++
																		goto l1463
																	l1464:
																		position, tokenIndex = position1463, tokenIndex1463
																		if buffer[position] != rune('G') {
																			goto l1293
																		}
																		position++
																	}
																l1463:
																	add(rulePegText, position1458)
																}
																{
																	add(ruleAction75, position)
																}
																add(ruleNeg, position1457)
															}
															break
														default:
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
																				if buffer[position] != rune('l') {
																					goto l1472
																				}
																				position++
																				goto l1471
																			l1472:
																				position, tokenIndex = position1471, tokenIndex1471
																				if buffer[position] != rune('L') {
																					goto l1468
																				}
																				position++
																			}
																		l1471:
																			{
																				position1473, tokenIndex1473 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1474
																				}
																				position++
																				goto l1473
																			l1474:
																				position, tokenIndex = position1473, tokenIndex1473
																				if buffer[position] != rune('D') {
																					goto l1468
																				}
																				position++
																			}
																		l1473:
																			{
																				position1475, tokenIndex1475 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1476
																				}
																				position++
																				goto l1475
																			l1476:
																				position, tokenIndex = position1475, tokenIndex1475
																				if buffer[position] != rune('I') {
																					goto l1468
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
																					goto l1468
																				}
																				position++
																			}
																		l1477:
																			add(rulePegText, position1470)
																		}
																		{
																			add(ruleAction91, position)
																		}
																		add(ruleLdir, position1469)
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
																				if buffer[position] != rune('l') {
																					goto l1484
																				}
																				position++
																				goto l1483
																			l1484:
																				position, tokenIndex = position1483, tokenIndex1483
																				if buffer[position] != rune('L') {
																					goto l1480
																				}
																				position++
																			}
																		l1483:
																			{
																				position1485, tokenIndex1485 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1486
																				}
																				position++
																				goto l1485
																			l1486:
																				position, tokenIndex = position1485, tokenIndex1485
																				if buffer[position] != rune('D') {
																					goto l1480
																				}
																				position++
																			}
																		l1485:
																			{
																				position1487, tokenIndex1487 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1488
																				}
																				position++
																				goto l1487
																			l1488:
																				position, tokenIndex = position1487, tokenIndex1487
																				if buffer[position] != rune('I') {
																					goto l1480
																				}
																				position++
																			}
																		l1487:
																			add(rulePegText, position1482)
																		}
																		{
																			add(ruleAction83, position)
																		}
																		add(ruleLdi, position1481)
																	}
																	goto l1467
																l1480:
																	position, tokenIndex = position1467, tokenIndex1467
																	{
																		position1491 := position
																		{
																			position1492 := position
																			{
																				position1493, tokenIndex1493 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1494
																				}
																				position++
																				goto l1493
																			l1494:
																				position, tokenIndex = position1493, tokenIndex1493
																				if buffer[position] != rune('C') {
																					goto l1490
																				}
																				position++
																			}
																		l1493:
																			{
																				position1495, tokenIndex1495 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1496
																				}
																				position++
																				goto l1495
																			l1496:
																				position, tokenIndex = position1495, tokenIndex1495
																				if buffer[position] != rune('P') {
																					goto l1490
																				}
																				position++
																			}
																		l1495:
																			{
																				position1497, tokenIndex1497 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1498
																				}
																				position++
																				goto l1497
																			l1498:
																				position, tokenIndex = position1497, tokenIndex1497
																				if buffer[position] != rune('I') {
																					goto l1490
																				}
																				position++
																			}
																		l1497:
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
																					goto l1490
																				}
																				position++
																			}
																		l1499:
																			add(rulePegText, position1492)
																		}
																		{
																			add(ruleAction92, position)
																		}
																		add(ruleCpir, position1491)
																	}
																	goto l1467
																l1490:
																	position, tokenIndex = position1467, tokenIndex1467
																	{
																		position1503 := position
																		{
																			position1504 := position
																			{
																				position1505, tokenIndex1505 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1506
																				}
																				position++
																				goto l1505
																			l1506:
																				position, tokenIndex = position1505, tokenIndex1505
																				if buffer[position] != rune('C') {
																					goto l1502
																				}
																				position++
																			}
																		l1505:
																			{
																				position1507, tokenIndex1507 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1508
																				}
																				position++
																				goto l1507
																			l1508:
																				position, tokenIndex = position1507, tokenIndex1507
																				if buffer[position] != rune('P') {
																					goto l1502
																				}
																				position++
																			}
																		l1507:
																			{
																				position1509, tokenIndex1509 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1510
																				}
																				position++
																				goto l1509
																			l1510:
																				position, tokenIndex = position1509, tokenIndex1509
																				if buffer[position] != rune('I') {
																					goto l1502
																				}
																				position++
																			}
																		l1509:
																			add(rulePegText, position1504)
																		}
																		{
																			add(ruleAction84, position)
																		}
																		add(ruleCpi, position1503)
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
																				if buffer[position] != rune('l') {
																					goto l1516
																				}
																				position++
																				goto l1515
																			l1516:
																				position, tokenIndex = position1515, tokenIndex1515
																				if buffer[position] != rune('L') {
																					goto l1512
																				}
																				position++
																			}
																		l1515:
																			{
																				position1517, tokenIndex1517 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1518
																				}
																				position++
																				goto l1517
																			l1518:
																				position, tokenIndex = position1517, tokenIndex1517
																				if buffer[position] != rune('D') {
																					goto l1512
																				}
																				position++
																			}
																		l1517:
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
																					goto l1512
																				}
																				position++
																			}
																		l1519:
																			{
																				position1521, tokenIndex1521 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1522
																				}
																				position++
																				goto l1521
																			l1522:
																				position, tokenIndex = position1521, tokenIndex1521
																				if buffer[position] != rune('R') {
																					goto l1512
																				}
																				position++
																			}
																		l1521:
																			add(rulePegText, position1514)
																		}
																		{
																			add(ruleAction95, position)
																		}
																		add(ruleLddr, position1513)
																	}
																	goto l1467
																l1512:
																	position, tokenIndex = position1467, tokenIndex1467
																	{
																		position1525 := position
																		{
																			position1526 := position
																			{
																				position1527, tokenIndex1527 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1528
																				}
																				position++
																				goto l1527
																			l1528:
																				position, tokenIndex = position1527, tokenIndex1527
																				if buffer[position] != rune('L') {
																					goto l1524
																				}
																				position++
																			}
																		l1527:
																			{
																				position1529, tokenIndex1529 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1530
																				}
																				position++
																				goto l1529
																			l1530:
																				position, tokenIndex = position1529, tokenIndex1529
																				if buffer[position] != rune('D') {
																					goto l1524
																				}
																				position++
																			}
																		l1529:
																			{
																				position1531, tokenIndex1531 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1532
																				}
																				position++
																				goto l1531
																			l1532:
																				position, tokenIndex = position1531, tokenIndex1531
																				if buffer[position] != rune('D') {
																					goto l1524
																				}
																				position++
																			}
																		l1531:
																			add(rulePegText, position1526)
																		}
																		{
																			add(ruleAction87, position)
																		}
																		add(ruleLdd, position1525)
																	}
																	goto l1467
																l1524:
																	position, tokenIndex = position1467, tokenIndex1467
																	{
																		position1535 := position
																		{
																			position1536 := position
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
																					goto l1534
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
																					goto l1534
																				}
																				position++
																			}
																		l1539:
																			{
																				position1541, tokenIndex1541 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1542
																				}
																				position++
																				goto l1541
																			l1542:
																				position, tokenIndex = position1541, tokenIndex1541
																				if buffer[position] != rune('D') {
																					goto l1534
																				}
																				position++
																			}
																		l1541:
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
																					goto l1534
																				}
																				position++
																			}
																		l1543:
																			add(rulePegText, position1536)
																		}
																		{
																			add(ruleAction96, position)
																		}
																		add(ruleCpdr, position1535)
																	}
																	goto l1467
																l1534:
																	position, tokenIndex = position1467, tokenIndex1467
																	{
																		position1546 := position
																		{
																			position1547 := position
																			{
																				position1548, tokenIndex1548 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1549
																				}
																				position++
																				goto l1548
																			l1549:
																				position, tokenIndex = position1548, tokenIndex1548
																				if buffer[position] != rune('C') {
																					goto l1293
																				}
																				position++
																			}
																		l1548:
																			{
																				position1550, tokenIndex1550 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1551
																				}
																				position++
																				goto l1550
																			l1551:
																				position, tokenIndex = position1550, tokenIndex1550
																				if buffer[position] != rune('P') {
																					goto l1293
																				}
																				position++
																			}
																		l1550:
																			{
																				position1552, tokenIndex1552 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1553
																				}
																				position++
																				goto l1552
																			l1553:
																				position, tokenIndex = position1552, tokenIndex1552
																				if buffer[position] != rune('D') {
																					goto l1293
																				}
																				position++
																			}
																		l1552:
																			add(rulePegText, position1547)
																		}
																		{
																			add(ruleAction88, position)
																		}
																		add(ruleCpd, position1546)
																	}
																}
															l1467:
																add(ruleBlit, position1466)
															}
															break
														}
													}

												}
											l1295:
												add(ruleEDSimple, position1294)
											}
											goto l965
										l1293:
											position, tokenIndex = position965, tokenIndex965
											{
												position1556 := position
												{
													position1557, tokenIndex1557 := position, tokenIndex
													{
														position1559 := position
														{
															position1560 := position
															{
																position1561, tokenIndex1561 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1562
																}
																position++
																goto l1561
															l1562:
																position, tokenIndex = position1561, tokenIndex1561
																if buffer[position] != rune('R') {
																	goto l1558
																}
																position++
															}
														l1561:
															{
																position1563, tokenIndex1563 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1564
																}
																position++
																goto l1563
															l1564:
																position, tokenIndex = position1563, tokenIndex1563
																if buffer[position] != rune('L') {
																	goto l1558
																}
																position++
															}
														l1563:
															{
																position1565, tokenIndex1565 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1566
																}
																position++
																goto l1565
															l1566:
																position, tokenIndex = position1565, tokenIndex1565
																if buffer[position] != rune('C') {
																	goto l1558
																}
																position++
															}
														l1565:
															{
																position1567, tokenIndex1567 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1568
																}
																position++
																goto l1567
															l1568:
																position, tokenIndex = position1567, tokenIndex1567
																if buffer[position] != rune('A') {
																	goto l1558
																}
																position++
															}
														l1567:
															add(rulePegText, position1560)
														}
														{
															add(ruleAction64, position)
														}
														add(ruleRlca, position1559)
													}
													goto l1557
												l1558:
													position, tokenIndex = position1557, tokenIndex1557
													{
														position1571 := position
														{
															position1572 := position
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
																	goto l1570
																}
																position++
															}
														l1573:
															{
																position1575, tokenIndex1575 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1576
																}
																position++
																goto l1575
															l1576:
																position, tokenIndex = position1575, tokenIndex1575
																if buffer[position] != rune('R') {
																	goto l1570
																}
																position++
															}
														l1575:
															{
																position1577, tokenIndex1577 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1578
																}
																position++
																goto l1577
															l1578:
																position, tokenIndex = position1577, tokenIndex1577
																if buffer[position] != rune('C') {
																	goto l1570
																}
																position++
															}
														l1577:
															{
																position1579, tokenIndex1579 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1580
																}
																position++
																goto l1579
															l1580:
																position, tokenIndex = position1579, tokenIndex1579
																if buffer[position] != rune('A') {
																	goto l1570
																}
																position++
															}
														l1579:
															add(rulePegText, position1572)
														}
														{
															add(ruleAction65, position)
														}
														add(ruleRrca, position1571)
													}
													goto l1557
												l1570:
													position, tokenIndex = position1557, tokenIndex1557
													{
														position1583 := position
														{
															position1584 := position
															{
																position1585, tokenIndex1585 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1586
																}
																position++
																goto l1585
															l1586:
																position, tokenIndex = position1585, tokenIndex1585
																if buffer[position] != rune('R') {
																	goto l1582
																}
																position++
															}
														l1585:
															{
																position1587, tokenIndex1587 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1588
																}
																position++
																goto l1587
															l1588:
																position, tokenIndex = position1587, tokenIndex1587
																if buffer[position] != rune('L') {
																	goto l1582
																}
																position++
															}
														l1587:
															{
																position1589, tokenIndex1589 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1590
																}
																position++
																goto l1589
															l1590:
																position, tokenIndex = position1589, tokenIndex1589
																if buffer[position] != rune('A') {
																	goto l1582
																}
																position++
															}
														l1589:
															add(rulePegText, position1584)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleRla, position1583)
													}
													goto l1557
												l1582:
													position, tokenIndex = position1557, tokenIndex1557
													{
														position1593 := position
														{
															position1594 := position
															{
																position1595, tokenIndex1595 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1596
																}
																position++
																goto l1595
															l1596:
																position, tokenIndex = position1595, tokenIndex1595
																if buffer[position] != rune('D') {
																	goto l1592
																}
																position++
															}
														l1595:
															{
																position1597, tokenIndex1597 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1598
																}
																position++
																goto l1597
															l1598:
																position, tokenIndex = position1597, tokenIndex1597
																if buffer[position] != rune('A') {
																	goto l1592
																}
																position++
															}
														l1597:
															{
																position1599, tokenIndex1599 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1600
																}
																position++
																goto l1599
															l1600:
																position, tokenIndex = position1599, tokenIndex1599
																if buffer[position] != rune('A') {
																	goto l1592
																}
																position++
															}
														l1599:
															add(rulePegText, position1594)
														}
														{
															add(ruleAction68, position)
														}
														add(ruleDaa, position1593)
													}
													goto l1557
												l1592:
													position, tokenIndex = position1557, tokenIndex1557
													{
														position1603 := position
														{
															position1604 := position
															{
																position1605, tokenIndex1605 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1606
																}
																position++
																goto l1605
															l1606:
																position, tokenIndex = position1605, tokenIndex1605
																if buffer[position] != rune('C') {
																	goto l1602
																}
																position++
															}
														l1605:
															{
																position1607, tokenIndex1607 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1608
																}
																position++
																goto l1607
															l1608:
																position, tokenIndex = position1607, tokenIndex1607
																if buffer[position] != rune('P') {
																	goto l1602
																}
																position++
															}
														l1607:
															{
																position1609, tokenIndex1609 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1610
																}
																position++
																goto l1609
															l1610:
																position, tokenIndex = position1609, tokenIndex1609
																if buffer[position] != rune('L') {
																	goto l1602
																}
																position++
															}
														l1609:
															add(rulePegText, position1604)
														}
														{
															add(ruleAction69, position)
														}
														add(ruleCpl, position1603)
													}
													goto l1557
												l1602:
													position, tokenIndex = position1557, tokenIndex1557
													{
														position1613 := position
														{
															position1614 := position
															{
																position1615, tokenIndex1615 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1616
																}
																position++
																goto l1615
															l1616:
																position, tokenIndex = position1615, tokenIndex1615
																if buffer[position] != rune('E') {
																	goto l1612
																}
																position++
															}
														l1615:
															{
																position1617, tokenIndex1617 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1618
																}
																position++
																goto l1617
															l1618:
																position, tokenIndex = position1617, tokenIndex1617
																if buffer[position] != rune('X') {
																	goto l1612
																}
																position++
															}
														l1617:
															{
																position1619, tokenIndex1619 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1620
																}
																position++
																goto l1619
															l1620:
																position, tokenIndex = position1619, tokenIndex1619
																if buffer[position] != rune('X') {
																	goto l1612
																}
																position++
															}
														l1619:
															add(rulePegText, position1614)
														}
														{
															add(ruleAction72, position)
														}
														add(ruleExx, position1613)
													}
													goto l1557
												l1612:
													position, tokenIndex = position1557, tokenIndex1557
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1623 := position
																{
																	position1624 := position
																	{
																		position1625, tokenIndex1625 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1626
																		}
																		position++
																		goto l1625
																	l1626:
																		position, tokenIndex = position1625, tokenIndex1625
																		if buffer[position] != rune('E') {
																			goto l1555
																		}
																		position++
																	}
																l1625:
																	{
																		position1627, tokenIndex1627 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1628
																		}
																		position++
																		goto l1627
																	l1628:
																		position, tokenIndex = position1627, tokenIndex1627
																		if buffer[position] != rune('I') {
																			goto l1555
																		}
																		position++
																	}
																l1627:
																	add(rulePegText, position1624)
																}
																{
																	add(ruleAction74, position)
																}
																add(ruleEi, position1623)
															}
															break
														case 'D', 'd':
															{
																position1630 := position
																{
																	position1631 := position
																	{
																		position1632, tokenIndex1632 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1633
																		}
																		position++
																		goto l1632
																	l1633:
																		position, tokenIndex = position1632, tokenIndex1632
																		if buffer[position] != rune('D') {
																			goto l1555
																		}
																		position++
																	}
																l1632:
																	{
																		position1634, tokenIndex1634 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1635
																		}
																		position++
																		goto l1634
																	l1635:
																		position, tokenIndex = position1634, tokenIndex1634
																		if buffer[position] != rune('I') {
																			goto l1555
																		}
																		position++
																	}
																l1634:
																	add(rulePegText, position1631)
																}
																{
																	add(ruleAction73, position)
																}
																add(ruleDi, position1630)
															}
															break
														case 'C', 'c':
															{
																position1637 := position
																{
																	position1638 := position
																	{
																		position1639, tokenIndex1639 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1640
																		}
																		position++
																		goto l1639
																	l1640:
																		position, tokenIndex = position1639, tokenIndex1639
																		if buffer[position] != rune('C') {
																			goto l1555
																		}
																		position++
																	}
																l1639:
																	{
																		position1641, tokenIndex1641 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1642
																		}
																		position++
																		goto l1641
																	l1642:
																		position, tokenIndex = position1641, tokenIndex1641
																		if buffer[position] != rune('C') {
																			goto l1555
																		}
																		position++
																	}
																l1641:
																	{
																		position1643, tokenIndex1643 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1644
																		}
																		position++
																		goto l1643
																	l1644:
																		position, tokenIndex = position1643, tokenIndex1643
																		if buffer[position] != rune('F') {
																			goto l1555
																		}
																		position++
																	}
																l1643:
																	add(rulePegText, position1638)
																}
																{
																	add(ruleAction71, position)
																}
																add(ruleCcf, position1637)
															}
															break
														case 'S', 's':
															{
																position1646 := position
																{
																	position1647 := position
																	{
																		position1648, tokenIndex1648 := position, tokenIndex
																		if buffer[position] != rune('s') {
																			goto l1649
																		}
																		position++
																		goto l1648
																	l1649:
																		position, tokenIndex = position1648, tokenIndex1648
																		if buffer[position] != rune('S') {
																			goto l1555
																		}
																		position++
																	}
																l1648:
																	{
																		position1650, tokenIndex1650 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1651
																		}
																		position++
																		goto l1650
																	l1651:
																		position, tokenIndex = position1650, tokenIndex1650
																		if buffer[position] != rune('C') {
																			goto l1555
																		}
																		position++
																	}
																l1650:
																	{
																		position1652, tokenIndex1652 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1653
																		}
																		position++
																		goto l1652
																	l1653:
																		position, tokenIndex = position1652, tokenIndex1652
																		if buffer[position] != rune('F') {
																			goto l1555
																		}
																		position++
																	}
																l1652:
																	add(rulePegText, position1647)
																}
																{
																	add(ruleAction70, position)
																}
																add(ruleScf, position1646)
															}
															break
														case 'R', 'r':
															{
																position1655 := position
																{
																	position1656 := position
																	{
																		position1657, tokenIndex1657 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1658
																		}
																		position++
																		goto l1657
																	l1658:
																		position, tokenIndex = position1657, tokenIndex1657
																		if buffer[position] != rune('R') {
																			goto l1555
																		}
																		position++
																	}
																l1657:
																	{
																		position1659, tokenIndex1659 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1660
																		}
																		position++
																		goto l1659
																	l1660:
																		position, tokenIndex = position1659, tokenIndex1659
																		if buffer[position] != rune('R') {
																			goto l1555
																		}
																		position++
																	}
																l1659:
																	{
																		position1661, tokenIndex1661 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1662
																		}
																		position++
																		goto l1661
																	l1662:
																		position, tokenIndex = position1661, tokenIndex1661
																		if buffer[position] != rune('A') {
																			goto l1555
																		}
																		position++
																	}
																l1661:
																	add(rulePegText, position1656)
																}
																{
																	add(ruleAction67, position)
																}
																add(ruleRra, position1655)
															}
															break
														case 'H', 'h':
															{
																position1664 := position
																{
																	position1665 := position
																	{
																		position1666, tokenIndex1666 := position, tokenIndex
																		if buffer[position] != rune('h') {
																			goto l1667
																		}
																		position++
																		goto l1666
																	l1667:
																		position, tokenIndex = position1666, tokenIndex1666
																		if buffer[position] != rune('H') {
																			goto l1555
																		}
																		position++
																	}
																l1666:
																	{
																		position1668, tokenIndex1668 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1669
																		}
																		position++
																		goto l1668
																	l1669:
																		position, tokenIndex = position1668, tokenIndex1668
																		if buffer[position] != rune('A') {
																			goto l1555
																		}
																		position++
																	}
																l1668:
																	{
																		position1670, tokenIndex1670 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1671
																		}
																		position++
																		goto l1670
																	l1671:
																		position, tokenIndex = position1670, tokenIndex1670
																		if buffer[position] != rune('L') {
																			goto l1555
																		}
																		position++
																	}
																l1670:
																	{
																		position1672, tokenIndex1672 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1673
																		}
																		position++
																		goto l1672
																	l1673:
																		position, tokenIndex = position1672, tokenIndex1672
																		if buffer[position] != rune('T') {
																			goto l1555
																		}
																		position++
																	}
																l1672:
																	add(rulePegText, position1665)
																}
																{
																	add(ruleAction63, position)
																}
																add(ruleHalt, position1664)
															}
															break
														default:
															{
																position1675 := position
																{
																	position1676 := position
																	{
																		position1677, tokenIndex1677 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1678
																		}
																		position++
																		goto l1677
																	l1678:
																		position, tokenIndex = position1677, tokenIndex1677
																		if buffer[position] != rune('N') {
																			goto l1555
																		}
																		position++
																	}
																l1677:
																	{
																		position1679, tokenIndex1679 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1680
																		}
																		position++
																		goto l1679
																	l1680:
																		position, tokenIndex = position1679, tokenIndex1679
																		if buffer[position] != rune('O') {
																			goto l1555
																		}
																		position++
																	}
																l1679:
																	{
																		position1681, tokenIndex1681 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1682
																		}
																		position++
																		goto l1681
																	l1682:
																		position, tokenIndex = position1681, tokenIndex1681
																		if buffer[position] != rune('P') {
																			goto l1555
																		}
																		position++
																	}
																l1681:
																	add(rulePegText, position1676)
																}
																{
																	add(ruleAction62, position)
																}
																add(ruleNop, position1675)
															}
															break
														}
													}

												}
											l1557:
												add(ruleSimple, position1556)
											}
											goto l965
										l1555:
											position, tokenIndex = position965, tokenIndex965
											{
												position1685 := position
												{
													position1686, tokenIndex1686 := position, tokenIndex
													{
														position1688 := position
														{
															position1689, tokenIndex1689 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1690
															}
															position++
															goto l1689
														l1690:
															position, tokenIndex = position1689, tokenIndex1689
															if buffer[position] != rune('R') {
																goto l1687
															}
															position++
														}
													l1689:
														{
															position1691, tokenIndex1691 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1692
															}
															position++
															goto l1691
														l1692:
															position, tokenIndex = position1691, tokenIndex1691
															if buffer[position] != rune('S') {
																goto l1687
															}
															position++
														}
													l1691:
														{
															position1693, tokenIndex1693 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1694
															}
															position++
															goto l1693
														l1694:
															position, tokenIndex = position1693, tokenIndex1693
															if buffer[position] != rune('T') {
																goto l1687
															}
															position++
														}
													l1693:
														if !_rules[rulews]() {
															goto l1687
														}
														if !_rules[rulen]() {
															goto l1687
														}
														{
															add(ruleAction99, position)
														}
														add(ruleRst, position1688)
													}
													goto l1686
												l1687:
													position, tokenIndex = position1686, tokenIndex1686
													{
														position1697 := position
														{
															position1698, tokenIndex1698 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1699
															}
															position++
															goto l1698
														l1699:
															position, tokenIndex = position1698, tokenIndex1698
															if buffer[position] != rune('J') {
																goto l1696
															}
															position++
														}
													l1698:
														{
															position1700, tokenIndex1700 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l1701
															}
															position++
															goto l1700
														l1701:
															position, tokenIndex = position1700, tokenIndex1700
															if buffer[position] != rune('P') {
																goto l1696
															}
															position++
														}
													l1700:
														if !_rules[rulews]() {
															goto l1696
														}
														{
															position1702, tokenIndex1702 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1702
															}
															if !_rules[rulesep]() {
																goto l1702
															}
															goto l1703
														l1702:
															position, tokenIndex = position1702, tokenIndex1702
														}
													l1703:
														if !_rules[ruleSrc16]() {
															goto l1696
														}
														{
															add(ruleAction102, position)
														}
														add(ruleJp, position1697)
													}
													goto l1686
												l1696:
													position, tokenIndex = position1686, tokenIndex1686
													{
														switch buffer[position] {
														case 'D', 'd':
															{
																position1706 := position
																{
																	position1707, tokenIndex1707 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1708
																	}
																	position++
																	goto l1707
																l1708:
																	position, tokenIndex = position1707, tokenIndex1707
																	if buffer[position] != rune('D') {
																		goto l1684
																	}
																	position++
																}
															l1707:
																{
																	position1709, tokenIndex1709 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1710
																	}
																	position++
																	goto l1709
																l1710:
																	position, tokenIndex = position1709, tokenIndex1709
																	if buffer[position] != rune('J') {
																		goto l1684
																	}
																	position++
																}
															l1709:
																{
																	position1711, tokenIndex1711 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1712
																	}
																	position++
																	goto l1711
																l1712:
																	position, tokenIndex = position1711, tokenIndex1711
																	if buffer[position] != rune('N') {
																		goto l1684
																	}
																	position++
																}
															l1711:
																{
																	position1713, tokenIndex1713 := position, tokenIndex
																	if buffer[position] != rune('z') {
																		goto l1714
																	}
																	position++
																	goto l1713
																l1714:
																	position, tokenIndex = position1713, tokenIndex1713
																	if buffer[position] != rune('Z') {
																		goto l1684
																	}
																	position++
																}
															l1713:
																if !_rules[rulews]() {
																	goto l1684
																}
																if !_rules[ruledisp]() {
																	goto l1684
																}
																{
																	add(ruleAction104, position)
																}
																add(ruleDjnz, position1706)
															}
															break
														case 'J', 'j':
															{
																position1716 := position
																{
																	position1717, tokenIndex1717 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1718
																	}
																	position++
																	goto l1717
																l1718:
																	position, tokenIndex = position1717, tokenIndex1717
																	if buffer[position] != rune('J') {
																		goto l1684
																	}
																	position++
																}
															l1717:
																{
																	position1719, tokenIndex1719 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1720
																	}
																	position++
																	goto l1719
																l1720:
																	position, tokenIndex = position1719, tokenIndex1719
																	if buffer[position] != rune('R') {
																		goto l1684
																	}
																	position++
																}
															l1719:
																if !_rules[rulews]() {
																	goto l1684
																}
																{
																	position1721, tokenIndex1721 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1721
																	}
																	if !_rules[rulesep]() {
																		goto l1721
																	}
																	goto l1722
																l1721:
																	position, tokenIndex = position1721, tokenIndex1721
																}
															l1722:
																if !_rules[ruledisp]() {
																	goto l1684
																}
																{
																	add(ruleAction103, position)
																}
																add(ruleJr, position1716)
															}
															break
														case 'R', 'r':
															{
																position1724 := position
																{
																	position1725, tokenIndex1725 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1726
																	}
																	position++
																	goto l1725
																l1726:
																	position, tokenIndex = position1725, tokenIndex1725
																	if buffer[position] != rune('R') {
																		goto l1684
																	}
																	position++
																}
															l1725:
																{
																	position1727, tokenIndex1727 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1728
																	}
																	position++
																	goto l1727
																l1728:
																	position, tokenIndex = position1727, tokenIndex1727
																	if buffer[position] != rune('E') {
																		goto l1684
																	}
																	position++
																}
															l1727:
																{
																	position1729, tokenIndex1729 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1730
																	}
																	position++
																	goto l1729
																l1730:
																	position, tokenIndex = position1729, tokenIndex1729
																	if buffer[position] != rune('T') {
																		goto l1684
																	}
																	position++
																}
															l1729:
																{
																	position1731, tokenIndex1731 := position, tokenIndex
																	if !_rules[rulews]() {
																		goto l1731
																	}
																	if !_rules[rulecc]() {
																		goto l1731
																	}
																	goto l1732
																l1731:
																	position, tokenIndex = position1731, tokenIndex1731
																}
															l1732:
																{
																	add(ruleAction101, position)
																}
																add(ruleRet, position1724)
															}
															break
														default:
															{
																position1734 := position
																{
																	position1735, tokenIndex1735 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1736
																	}
																	position++
																	goto l1735
																l1736:
																	position, tokenIndex = position1735, tokenIndex1735
																	if buffer[position] != rune('C') {
																		goto l1684
																	}
																	position++
																}
															l1735:
																{
																	position1737, tokenIndex1737 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1738
																	}
																	position++
																	goto l1737
																l1738:
																	position, tokenIndex = position1737, tokenIndex1737
																	if buffer[position] != rune('A') {
																		goto l1684
																	}
																	position++
																}
															l1737:
																{
																	position1739, tokenIndex1739 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1740
																	}
																	position++
																	goto l1739
																l1740:
																	position, tokenIndex = position1739, tokenIndex1739
																	if buffer[position] != rune('L') {
																		goto l1684
																	}
																	position++
																}
															l1739:
																{
																	position1741, tokenIndex1741 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1742
																	}
																	position++
																	goto l1741
																l1742:
																	position, tokenIndex = position1741, tokenIndex1741
																	if buffer[position] != rune('L') {
																		goto l1684
																	}
																	position++
																}
															l1741:
																if !_rules[rulews]() {
																	goto l1684
																}
																{
																	position1743, tokenIndex1743 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1743
																	}
																	if !_rules[rulesep]() {
																		goto l1743
																	}
																	goto l1744
																l1743:
																	position, tokenIndex = position1743, tokenIndex1743
																}
															l1744:
																if !_rules[ruleSrc16]() {
																	goto l1684
																}
																{
																	add(ruleAction100, position)
																}
																add(ruleCall, position1734)
															}
															break
														}
													}

												}
											l1686:
												add(ruleJump, position1685)
											}
											goto l965
										l1684:
											position, tokenIndex = position965, tokenIndex965
											{
												position1746 := position
												{
													position1747, tokenIndex1747 := position, tokenIndex
													{
														position1749 := position
														{
															position1750, tokenIndex1750 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1751
															}
															position++
															goto l1750
														l1751:
															position, tokenIndex = position1750, tokenIndex1750
															if buffer[position] != rune('I') {
																goto l1748
															}
															position++
														}
													l1750:
														{
															position1752, tokenIndex1752 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1753
															}
															position++
															goto l1752
														l1753:
															position, tokenIndex = position1752, tokenIndex1752
															if buffer[position] != rune('N') {
																goto l1748
															}
															position++
														}
													l1752:
														if !_rules[rulews]() {
															goto l1748
														}
														if !_rules[ruleReg8]() {
															goto l1748
														}
														if !_rules[rulesep]() {
															goto l1748
														}
														if !_rules[rulePort]() {
															goto l1748
														}
														{
															add(ruleAction105, position)
														}
														add(ruleIN, position1749)
													}
													goto l1747
												l1748:
													position, tokenIndex = position1747, tokenIndex1747
													{
														position1755 := position
														{
															position1756, tokenIndex1756 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1757
															}
															position++
															goto l1756
														l1757:
															position, tokenIndex = position1756, tokenIndex1756
															if buffer[position] != rune('O') {
																goto l901
															}
															position++
														}
													l1756:
														{
															position1758, tokenIndex1758 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1759
															}
															position++
															goto l1758
														l1759:
															position, tokenIndex = position1758, tokenIndex1758
															if buffer[position] != rune('U') {
																goto l901
															}
															position++
														}
													l1758:
														{
															position1760, tokenIndex1760 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1761
															}
															position++
															goto l1760
														l1761:
															position, tokenIndex = position1760, tokenIndex1760
															if buffer[position] != rune('T') {
																goto l901
															}
															position++
														}
													l1760:
														if !_rules[rulews]() {
															goto l901
														}
														if !_rules[rulePort]() {
															goto l901
														}
														if !_rules[rulesep]() {
															goto l901
														}
														if !_rules[ruleReg8]() {
															goto l901
														}
														{
															add(ruleAction106, position)
														}
														add(ruleOUT, position1755)
													}
												}
											l1747:
												add(ruleIO, position1746)
											}
										}
									l965:
										add(ruleInstruction, position964)
									}
								}
							l904:
								add(ruleStatement, position903)
							}
							goto l902
						l901:
							position, tokenIndex = position901, tokenIndex901
						}
					l902:
						{
							position1763, tokenIndex1763 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1763
							}
							goto l1764
						l1763:
							position, tokenIndex = position1763, tokenIndex1763
						}
					l1764:
						{
							position1765, tokenIndex1765 := position, tokenIndex
							{
								position1767 := position
								{
									position1768, tokenIndex1768 := position, tokenIndex
									if buffer[position] != rune(';') {
										goto l1769
									}
									position++
									goto l1768
								l1769:
									position, tokenIndex = position1768, tokenIndex1768
									if buffer[position] != rune('#') {
										goto l1765
									}
									position++
								}
							l1768:
							l1770:
								{
									position1771, tokenIndex1771 := position, tokenIndex
									{
										position1772, tokenIndex1772 := position, tokenIndex
										if buffer[position] != rune('\n') {
											goto l1772
										}
										position++
										goto l1771
									l1772:
										position, tokenIndex = position1772, tokenIndex1772
									}
									if !matchDot() {
										goto l1771
									}
									goto l1770
								l1771:
									position, tokenIndex = position1771, tokenIndex1771
								}
								add(ruleComment, position1767)
							}
							goto l1766
						l1765:
							position, tokenIndex = position1765, tokenIndex1765
						}
					l1766:
						{
							position1773, tokenIndex1773 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1773
							}
							goto l1774
						l1773:
							position, tokenIndex = position1773, tokenIndex1773
						}
					l1774:
						{
							position1775, tokenIndex1775 := position, tokenIndex
							{
								position1777, tokenIndex1777 := position, tokenIndex
								if buffer[position] != rune('\r') {
									goto l1777
								}
								position++
								goto l1778
							l1777:
								position, tokenIndex = position1777, tokenIndex1777
							}
						l1778:
							if buffer[position] != rune('\n') {
								goto l1776
							}
							position++
							goto l1775
						l1776:
							position, tokenIndex = position1775, tokenIndex1775
							if buffer[position] != rune(':') {
								goto l3
							}
							position++
						}
					l1775:
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position892)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position1780, tokenIndex1780 := position, tokenIndex
					if !matchDot() {
						goto l1780
					}
					goto l0
				l1780:
					position, tokenIndex = position1780, tokenIndex1780
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
		/* 3 Directive <- <(Defb / ((&('D' | 'd') Defs) | (&('O' | 'o') Org) | (&('a') Aseg) | (&('.' | 't') Title)))> */
		nil,
		/* 4 Title <- <('.'? ('t' 'i' 't' 'l' 'e') ws '\'' (!'\'' .)* '\'')> */
		nil,
		/* 5 Aseg <- <('a' 's' 'e' 'g')> */
		nil,
		/* 6 Org <- <(('o' / 'O') ('r' / 'R') ('g' / 'G') ws nn Action1)> */
		nil,
		/* 7 Defb <- <(((('d' / 'D') ('e' / 'E') ('f' / 'F') ('b' / 'B')) / (('d' / 'D') ('b' / 'B')) / (('e' / 'E') ('q' / 'Q') ('u' / 'U'))) ws n Action2)> */
		nil,
		/* 8 Defs <- <(((('d' / 'D') ('e' / 'E') ('f' / 'F') ('s' / 'S')) / (('d' / 'D') ('s' / 'S'))) ws n Action3)> */
		nil,
		/* 9 LabelDefn <- <(LabelText ':' ws Action4)> */
		nil,
		/* 10 LabelText <- <<(alpha alphanum*)>> */
		func() bool {
			position1790, tokenIndex1790 := position, tokenIndex
			{
				position1791 := position
				{
					position1792 := position
					if !_rules[rulealpha]() {
						goto l1790
					}
				l1793:
					{
						position1794, tokenIndex1794 := position, tokenIndex
						{
							position1795 := position
							{
								position1796, tokenIndex1796 := position, tokenIndex
								if !_rules[rulealpha]() {
									goto l1797
								}
								goto l1796
							l1797:
								position, tokenIndex = position1796, tokenIndex1796
								{
									position1798 := position
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1794
									}
									position++
									add(rulenum, position1798)
								}
							}
						l1796:
							add(rulealphanum, position1795)
						}
						goto l1793
					l1794:
						position, tokenIndex = position1794, tokenIndex1794
					}
					add(rulePegText, position1792)
				}
				add(ruleLabelText, position1791)
			}
			return true
		l1790:
			position, tokenIndex = position1790, tokenIndex1790
			return false
		},
		/* 11 alphanum <- <(alpha / num)> */
		nil,
		/* 12 alpha <- <([a-z] / [A-Z])> */
		func() bool {
			position1800, tokenIndex1800 := position, tokenIndex
			{
				position1801 := position
				{
					position1802, tokenIndex1802 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1803
					}
					position++
					goto l1802
				l1803:
					position, tokenIndex = position1802, tokenIndex1802
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1800
					}
					position++
				}
			l1802:
				add(rulealpha, position1801)
			}
			return true
		l1800:
			position, tokenIndex = position1800, tokenIndex1800
			return false
		},
		/* 13 num <- <[0-9]> */
		nil,
		/* 14 Comment <- <((';' / '#') (!'\n' .)*)> */
		nil,
		/* 15 Instruction <- <(Assignment / Inc / Dec / Alu16 / Alu / BitOp / EDSimple / Simple / Jump / IO)> */
		nil,
		/* 16 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 17 Load <- <(Load16 / Load8)> */
		nil,
		/* 18 Load8 <- <(('l' / 'L') ('d' / 'D') ws Dst8 sep Src8 Action5)> */
		nil,
		/* 19 Load16 <- <(('l' / 'L') ('d' / 'D') ws Dst16 sep Src16 Action6)> */
		nil,
		/* 20 Push <- <(('p' / 'P') ('u' / 'U') ('s' / 'S') ('h' / 'H') ws Src16 Action7)> */
		nil,
		/* 21 Pop <- <(('p' / 'P') ('o' / 'O') ('p' / 'P') ws Dst16 Action8)> */
		nil,
		/* 22 Ex <- <(('e' / 'E') ('x' / 'X') ws Dst16 sep Src16 Action9)> */
		nil,
		/* 23 Inc <- <(Inc16Indexed8 / Inc16 / Inc8)> */
		nil,
		/* 24 Inc16Indexed8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws ILoc8 Action10)> */
		nil,
		/* 25 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action11)> */
		nil,
		/* 26 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action12)> */
		nil,
		/* 27 Dec <- <(Dec16Indexed8 / Dec16 / Dec8)> */
		nil,
		/* 28 Dec16Indexed8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws ILoc8 Action13)> */
		nil,
		/* 29 Dec8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc8 Action14)> */
		nil,
		/* 30 Dec16 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc16 Action15)> */
		nil,
		/* 31 Alu16 <- <(Add16 / Adc16 / Sbc16)> */
		nil,
		/* 32 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action16)> */
		nil,
		/* 33 Adc16 <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws Dst16 sep Src16 Action17)> */
		nil,
		/* 34 Sbc16 <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws Dst16 sep Src16 Action18)> */
		nil,
		/* 35 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action19)> */
		nil,
		/* 36 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action20)> */
		func() bool {
			position1827, tokenIndex1827 := position, tokenIndex
			{
				position1828 := position
				{
					position1829, tokenIndex1829 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1830
					}
					goto l1829
				l1830:
					position, tokenIndex = position1829, tokenIndex1829
					if !_rules[ruleReg8]() {
						goto l1831
					}
					goto l1829
				l1831:
					position, tokenIndex = position1829, tokenIndex1829
					if !_rules[ruleReg16Contents]() {
						goto l1832
					}
					goto l1829
				l1832:
					position, tokenIndex = position1829, tokenIndex1829
					if !_rules[rulenn_contents]() {
						goto l1827
					}
				}
			l1829:
				{
					add(ruleAction20, position)
				}
				add(ruleSrc8, position1828)
			}
			return true
		l1827:
			position, tokenIndex = position1827, tokenIndex1827
			return false
		},
		/* 37 Loc8 <- <((Reg8 / Reg16Contents) Action21)> */
		func() bool {
			position1834, tokenIndex1834 := position, tokenIndex
			{
				position1835 := position
				{
					position1836, tokenIndex1836 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1837
					}
					goto l1836
				l1837:
					position, tokenIndex = position1836, tokenIndex1836
					if !_rules[ruleReg16Contents]() {
						goto l1834
					}
				}
			l1836:
				{
					add(ruleAction21, position)
				}
				add(ruleLoc8, position1835)
			}
			return true
		l1834:
			position, tokenIndex = position1834, tokenIndex1834
			return false
		},
		/* 38 Copy8 <- <(Reg8 Action22)> */
		func() bool {
			position1839, tokenIndex1839 := position, tokenIndex
			{
				position1840 := position
				if !_rules[ruleReg8]() {
					goto l1839
				}
				{
					add(ruleAction22, position)
				}
				add(ruleCopy8, position1840)
			}
			return true
		l1839:
			position, tokenIndex = position1839, tokenIndex1839
			return false
		},
		/* 39 ILoc8 <- <(IReg8 Action23)> */
		func() bool {
			position1842, tokenIndex1842 := position, tokenIndex
			{
				position1843 := position
				if !_rules[ruleIReg8]() {
					goto l1842
				}
				{
					add(ruleAction23, position)
				}
				add(ruleILoc8, position1843)
			}
			return true
		l1842:
			position, tokenIndex = position1842, tokenIndex1842
			return false
		},
		/* 40 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action24)> */
		func() bool {
			position1845, tokenIndex1845 := position, tokenIndex
			{
				position1846 := position
				{
					position1847 := position
					{
						position1848, tokenIndex1848 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1849
						}
						goto l1848
					l1849:
						position, tokenIndex = position1848, tokenIndex1848
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1851 := position
									{
										position1852, tokenIndex1852 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1853
										}
										position++
										goto l1852
									l1853:
										position, tokenIndex = position1852, tokenIndex1852
										if buffer[position] != rune('R') {
											goto l1845
										}
										position++
									}
								l1852:
									add(ruleR, position1851)
								}
								break
							case 'I', 'i':
								{
									position1854 := position
									{
										position1855, tokenIndex1855 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1856
										}
										position++
										goto l1855
									l1856:
										position, tokenIndex = position1855, tokenIndex1855
										if buffer[position] != rune('I') {
											goto l1845
										}
										position++
									}
								l1855:
									add(ruleI, position1854)
								}
								break
							case 'L', 'l':
								{
									position1857 := position
									{
										position1858, tokenIndex1858 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1859
										}
										position++
										goto l1858
									l1859:
										position, tokenIndex = position1858, tokenIndex1858
										if buffer[position] != rune('L') {
											goto l1845
										}
										position++
									}
								l1858:
									add(ruleL, position1857)
								}
								break
							case 'H', 'h':
								{
									position1860 := position
									{
										position1861, tokenIndex1861 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1862
										}
										position++
										goto l1861
									l1862:
										position, tokenIndex = position1861, tokenIndex1861
										if buffer[position] != rune('H') {
											goto l1845
										}
										position++
									}
								l1861:
									add(ruleH, position1860)
								}
								break
							case 'E', 'e':
								{
									position1863 := position
									{
										position1864, tokenIndex1864 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1865
										}
										position++
										goto l1864
									l1865:
										position, tokenIndex = position1864, tokenIndex1864
										if buffer[position] != rune('E') {
											goto l1845
										}
										position++
									}
								l1864:
									add(ruleE, position1863)
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
									add(ruleD, position1866)
								}
								break
							case 'C', 'c':
								{
									position1869 := position
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
											goto l1845
										}
										position++
									}
								l1870:
									add(ruleC, position1869)
								}
								break
							case 'B', 'b':
								{
									position1872 := position
									{
										position1873, tokenIndex1873 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1874
										}
										position++
										goto l1873
									l1874:
										position, tokenIndex = position1873, tokenIndex1873
										if buffer[position] != rune('B') {
											goto l1845
										}
										position++
									}
								l1873:
									add(ruleB, position1872)
								}
								break
							case 'F', 'f':
								{
									position1875 := position
									{
										position1876, tokenIndex1876 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1877
										}
										position++
										goto l1876
									l1877:
										position, tokenIndex = position1876, tokenIndex1876
										if buffer[position] != rune('F') {
											goto l1845
										}
										position++
									}
								l1876:
									add(ruleF, position1875)
								}
								break
							default:
								{
									position1878 := position
									{
										position1879, tokenIndex1879 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1880
										}
										position++
										goto l1879
									l1880:
										position, tokenIndex = position1879, tokenIndex1879
										if buffer[position] != rune('A') {
											goto l1845
										}
										position++
									}
								l1879:
									add(ruleA, position1878)
								}
								break
							}
						}

					}
				l1848:
					add(rulePegText, position1847)
				}
				{
					add(ruleAction24, position)
				}
				add(ruleReg8, position1846)
			}
			return true
		l1845:
			position, tokenIndex = position1845, tokenIndex1845
			return false
		},
		/* 41 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action25)> */
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
							{
								position1892, tokenIndex1892 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1893
								}
								position++
								goto l1892
							l1893:
								position, tokenIndex = position1892, tokenIndex1892
								if buffer[position] != rune('H') {
									goto l1886
								}
								position++
							}
						l1892:
							add(ruleIXH, position1887)
						}
						goto l1885
					l1886:
						position, tokenIndex = position1885, tokenIndex1885
						{
							position1895 := position
							{
								position1896, tokenIndex1896 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1897
								}
								position++
								goto l1896
							l1897:
								position, tokenIndex = position1896, tokenIndex1896
								if buffer[position] != rune('I') {
									goto l1894
								}
								position++
							}
						l1896:
							{
								position1898, tokenIndex1898 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1899
								}
								position++
								goto l1898
							l1899:
								position, tokenIndex = position1898, tokenIndex1898
								if buffer[position] != rune('X') {
									goto l1894
								}
								position++
							}
						l1898:
							{
								position1900, tokenIndex1900 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1901
								}
								position++
								goto l1900
							l1901:
								position, tokenIndex = position1900, tokenIndex1900
								if buffer[position] != rune('L') {
									goto l1894
								}
								position++
							}
						l1900:
							add(ruleIXL, position1895)
						}
						goto l1885
					l1894:
						position, tokenIndex = position1885, tokenIndex1885
						{
							position1903 := position
							{
								position1904, tokenIndex1904 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1905
								}
								position++
								goto l1904
							l1905:
								position, tokenIndex = position1904, tokenIndex1904
								if buffer[position] != rune('I') {
									goto l1902
								}
								position++
							}
						l1904:
							{
								position1906, tokenIndex1906 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1907
								}
								position++
								goto l1906
							l1907:
								position, tokenIndex = position1906, tokenIndex1906
								if buffer[position] != rune('Y') {
									goto l1902
								}
								position++
							}
						l1906:
							{
								position1908, tokenIndex1908 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1909
								}
								position++
								goto l1908
							l1909:
								position, tokenIndex = position1908, tokenIndex1908
								if buffer[position] != rune('H') {
									goto l1902
								}
								position++
							}
						l1908:
							add(ruleIYH, position1903)
						}
						goto l1885
					l1902:
						position, tokenIndex = position1885, tokenIndex1885
						{
							position1910 := position
							{
								position1911, tokenIndex1911 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1912
								}
								position++
								goto l1911
							l1912:
								position, tokenIndex = position1911, tokenIndex1911
								if buffer[position] != rune('I') {
									goto l1882
								}
								position++
							}
						l1911:
							{
								position1913, tokenIndex1913 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1914
								}
								position++
								goto l1913
							l1914:
								position, tokenIndex = position1913, tokenIndex1913
								if buffer[position] != rune('Y') {
									goto l1882
								}
								position++
							}
						l1913:
							{
								position1915, tokenIndex1915 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1916
								}
								position++
								goto l1915
							l1916:
								position, tokenIndex = position1915, tokenIndex1915
								if buffer[position] != rune('L') {
									goto l1882
								}
								position++
							}
						l1915:
							add(ruleIYL, position1910)
						}
					}
				l1885:
					add(rulePegText, position1884)
				}
				{
					add(ruleAction25, position)
				}
				add(ruleIReg8, position1883)
			}
			return true
		l1882:
			position, tokenIndex = position1882, tokenIndex1882
			return false
		},
		/* 42 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action26)> */
		func() bool {
			position1918, tokenIndex1918 := position, tokenIndex
			{
				position1919 := position
				{
					position1920, tokenIndex1920 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1921
					}
					goto l1920
				l1921:
					position, tokenIndex = position1920, tokenIndex1920
					if !_rules[rulenn_contents]() {
						goto l1922
					}
					goto l1920
				l1922:
					position, tokenIndex = position1920, tokenIndex1920
					if !_rules[ruleReg16Contents]() {
						goto l1918
					}
				}
			l1920:
				{
					add(ruleAction26, position)
				}
				add(ruleDst16, position1919)
			}
			return true
		l1918:
			position, tokenIndex = position1918, tokenIndex1918
			return false
		},
		/* 43 Src16 <- <((Reg16 / nn / nn_contents) Action27)> */
		func() bool {
			position1924, tokenIndex1924 := position, tokenIndex
			{
				position1925 := position
				{
					position1926, tokenIndex1926 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1927
					}
					goto l1926
				l1927:
					position, tokenIndex = position1926, tokenIndex1926
					if !_rules[rulenn]() {
						goto l1928
					}
					goto l1926
				l1928:
					position, tokenIndex = position1926, tokenIndex1926
					if !_rules[rulenn_contents]() {
						goto l1924
					}
				}
			l1926:
				{
					add(ruleAction27, position)
				}
				add(ruleSrc16, position1925)
			}
			return true
		l1924:
			position, tokenIndex = position1924, tokenIndex1924
			return false
		},
		/* 44 Loc16 <- <(Reg16 Action28)> */
		func() bool {
			position1930, tokenIndex1930 := position, tokenIndex
			{
				position1931 := position
				if !_rules[ruleReg16]() {
					goto l1930
				}
				{
					add(ruleAction28, position)
				}
				add(ruleLoc16, position1931)
			}
			return true
		l1930:
			position, tokenIndex = position1930, tokenIndex1930
			return false
		},
		/* 45 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action29)> */
		func() bool {
			position1933, tokenIndex1933 := position, tokenIndex
			{
				position1934 := position
				{
					position1935 := position
					{
						position1936, tokenIndex1936 := position, tokenIndex
						{
							position1938 := position
							{
								position1939, tokenIndex1939 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1940
								}
								position++
								goto l1939
							l1940:
								position, tokenIndex = position1939, tokenIndex1939
								if buffer[position] != rune('A') {
									goto l1937
								}
								position++
							}
						l1939:
							{
								position1941, tokenIndex1941 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1942
								}
								position++
								goto l1941
							l1942:
								position, tokenIndex = position1941, tokenIndex1941
								if buffer[position] != rune('F') {
									goto l1937
								}
								position++
							}
						l1941:
							if buffer[position] != rune('\'') {
								goto l1937
							}
							position++
							add(ruleAF_PRIME, position1938)
						}
						goto l1936
					l1937:
						position, tokenIndex = position1936, tokenIndex1936
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1933
								}
								break
							case 'S', 's':
								{
									position1944 := position
									{
										position1945, tokenIndex1945 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1946
										}
										position++
										goto l1945
									l1946:
										position, tokenIndex = position1945, tokenIndex1945
										if buffer[position] != rune('S') {
											goto l1933
										}
										position++
									}
								l1945:
									{
										position1947, tokenIndex1947 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1948
										}
										position++
										goto l1947
									l1948:
										position, tokenIndex = position1947, tokenIndex1947
										if buffer[position] != rune('P') {
											goto l1933
										}
										position++
									}
								l1947:
									add(ruleSP, position1944)
								}
								break
							case 'H', 'h':
								{
									position1949 := position
									{
										position1950, tokenIndex1950 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1951
										}
										position++
										goto l1950
									l1951:
										position, tokenIndex = position1950, tokenIndex1950
										if buffer[position] != rune('H') {
											goto l1933
										}
										position++
									}
								l1950:
									{
										position1952, tokenIndex1952 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1953
										}
										position++
										goto l1952
									l1953:
										position, tokenIndex = position1952, tokenIndex1952
										if buffer[position] != rune('L') {
											goto l1933
										}
										position++
									}
								l1952:
									add(ruleHL, position1949)
								}
								break
							case 'D', 'd':
								{
									position1954 := position
									{
										position1955, tokenIndex1955 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1956
										}
										position++
										goto l1955
									l1956:
										position, tokenIndex = position1955, tokenIndex1955
										if buffer[position] != rune('D') {
											goto l1933
										}
										position++
									}
								l1955:
									{
										position1957, tokenIndex1957 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1958
										}
										position++
										goto l1957
									l1958:
										position, tokenIndex = position1957, tokenIndex1957
										if buffer[position] != rune('E') {
											goto l1933
										}
										position++
									}
								l1957:
									add(ruleDE, position1954)
								}
								break
							case 'B', 'b':
								{
									position1959 := position
									{
										position1960, tokenIndex1960 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1961
										}
										position++
										goto l1960
									l1961:
										position, tokenIndex = position1960, tokenIndex1960
										if buffer[position] != rune('B') {
											goto l1933
										}
										position++
									}
								l1960:
									{
										position1962, tokenIndex1962 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1963
										}
										position++
										goto l1962
									l1963:
										position, tokenIndex = position1962, tokenIndex1962
										if buffer[position] != rune('C') {
											goto l1933
										}
										position++
									}
								l1962:
									add(ruleBC, position1959)
								}
								break
							default:
								{
									position1964 := position
									{
										position1965, tokenIndex1965 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1966
										}
										position++
										goto l1965
									l1966:
										position, tokenIndex = position1965, tokenIndex1965
										if buffer[position] != rune('A') {
											goto l1933
										}
										position++
									}
								l1965:
									{
										position1967, tokenIndex1967 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1968
										}
										position++
										goto l1967
									l1968:
										position, tokenIndex = position1967, tokenIndex1967
										if buffer[position] != rune('F') {
											goto l1933
										}
										position++
									}
								l1967:
									add(ruleAF, position1964)
								}
								break
							}
						}

					}
				l1936:
					add(rulePegText, position1935)
				}
				{
					add(ruleAction29, position)
				}
				add(ruleReg16, position1934)
			}
			return true
		l1933:
			position, tokenIndex = position1933, tokenIndex1933
			return false
		},
		/* 46 IReg16 <- <(<(IX / IY)> Action30)> */
		func() bool {
			position1970, tokenIndex1970 := position, tokenIndex
			{
				position1971 := position
				{
					position1972 := position
					{
						position1973, tokenIndex1973 := position, tokenIndex
						{
							position1975 := position
							{
								position1976, tokenIndex1976 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1977
								}
								position++
								goto l1976
							l1977:
								position, tokenIndex = position1976, tokenIndex1976
								if buffer[position] != rune('I') {
									goto l1974
								}
								position++
							}
						l1976:
							{
								position1978, tokenIndex1978 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1979
								}
								position++
								goto l1978
							l1979:
								position, tokenIndex = position1978, tokenIndex1978
								if buffer[position] != rune('X') {
									goto l1974
								}
								position++
							}
						l1978:
							add(ruleIX, position1975)
						}
						goto l1973
					l1974:
						position, tokenIndex = position1973, tokenIndex1973
						{
							position1980 := position
							{
								position1981, tokenIndex1981 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1982
								}
								position++
								goto l1981
							l1982:
								position, tokenIndex = position1981, tokenIndex1981
								if buffer[position] != rune('I') {
									goto l1970
								}
								position++
							}
						l1981:
							{
								position1983, tokenIndex1983 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1984
								}
								position++
								goto l1983
							l1984:
								position, tokenIndex = position1983, tokenIndex1983
								if buffer[position] != rune('Y') {
									goto l1970
								}
								position++
							}
						l1983:
							add(ruleIY, position1980)
						}
					}
				l1973:
					add(rulePegText, position1972)
				}
				{
					add(ruleAction30, position)
				}
				add(ruleIReg16, position1971)
			}
			return true
		l1970:
			position, tokenIndex = position1970, tokenIndex1970
			return false
		},
		/* 47 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1986, tokenIndex1986 := position, tokenIndex
			{
				position1987 := position
				{
					position1988, tokenIndex1988 := position, tokenIndex
					{
						position1990 := position
						if buffer[position] != rune('(') {
							goto l1989
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1989
						}
						{
							position1991, tokenIndex1991 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1991
							}
							goto l1992
						l1991:
							position, tokenIndex = position1991, tokenIndex1991
						}
					l1992:
						if !_rules[ruledisp]() {
							goto l1989
						}
						{
							position1993, tokenIndex1993 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1993
							}
							goto l1994
						l1993:
							position, tokenIndex = position1993, tokenIndex1993
						}
					l1994:
						if buffer[position] != rune(')') {
							goto l1989
						}
						position++
						{
							add(ruleAction32, position)
						}
						add(ruleIndexedR16C, position1990)
					}
					goto l1988
				l1989:
					position, tokenIndex = position1988, tokenIndex1988
					{
						position1996 := position
						if buffer[position] != rune('(') {
							goto l1986
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1986
						}
						if buffer[position] != rune(')') {
							goto l1986
						}
						position++
						{
							add(ruleAction31, position)
						}
						add(rulePlainR16C, position1996)
					}
				}
			l1988:
				add(ruleReg16Contents, position1987)
			}
			return true
		l1986:
			position, tokenIndex = position1986, tokenIndex1986
			return false
		},
		/* 48 PlainR16C <- <('(' Reg16 ')' Action31)> */
		nil,
		/* 49 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action32)> */
		nil,
		/* 50 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position2000, tokenIndex2000 := position, tokenIndex
			{
				position2001 := position
				{
					position2002, tokenIndex2002 := position, tokenIndex
					{
						position2004 := position
						{
							position2005 := position
							if !_rules[rulehexdigit]() {
								goto l2003
							}
							if !_rules[rulehexdigit]() {
								goto l2003
							}
							add(rulePegText, position2005)
						}
						{
							position2006, tokenIndex2006 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2007
							}
							position++
							goto l2006
						l2007:
							position, tokenIndex = position2006, tokenIndex2006
							if buffer[position] != rune('H') {
								goto l2003
							}
							position++
						}
					l2006:
						{
							add(ruleAction36, position)
						}
						add(rulehexByteH, position2004)
					}
					goto l2002
				l2003:
					position, tokenIndex = position2002, tokenIndex2002
					{
						position2010 := position
						if buffer[position] != rune('0') {
							goto l2009
						}
						position++
						{
							position2011, tokenIndex2011 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2012
							}
							position++
							goto l2011
						l2012:
							position, tokenIndex = position2011, tokenIndex2011
							if buffer[position] != rune('X') {
								goto l2009
							}
							position++
						}
					l2011:
						{
							position2013 := position
							if !_rules[rulehexdigit]() {
								goto l2009
							}
							if !_rules[rulehexdigit]() {
								goto l2009
							}
							add(rulePegText, position2013)
						}
						{
							add(ruleAction37, position)
						}
						add(rulehexByte0x, position2010)
					}
					goto l2002
				l2009:
					position, tokenIndex = position2002, tokenIndex2002
					{
						position2015 := position
						{
							position2016 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2000
							}
							position++
						l2017:
							{
								position2018, tokenIndex2018 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2018
								}
								position++
								goto l2017
							l2018:
								position, tokenIndex = position2018, tokenIndex2018
							}
							add(rulePegText, position2016)
						}
						{
							add(ruleAction38, position)
						}
						add(ruledecimalByte, position2015)
					}
				}
			l2002:
				add(rulen, position2001)
			}
			return true
		l2000:
			position, tokenIndex = position2000, tokenIndex2000
			return false
		},
		/* 51 nn <- <(LabelNN / hexWordH / hexWord0x)> */
		func() bool {
			position2020, tokenIndex2020 := position, tokenIndex
			{
				position2021 := position
				{
					position2022, tokenIndex2022 := position, tokenIndex
					{
						position2024 := position
						{
							position2025 := position
							if !_rules[ruleLabelText]() {
								goto l2023
							}
							add(rulePegText, position2025)
						}
						{
							add(ruleAction39, position)
						}
						add(ruleLabelNN, position2024)
					}
					goto l2022
				l2023:
					position, tokenIndex = position2022, tokenIndex2022
					{
						position2028 := position
						{
							position2029 := position
							if !_rules[rulehexdigit]() {
								goto l2027
							}
							if !_rules[rulehexdigit]() {
								goto l2027
							}
							if !_rules[rulehexdigit]() {
								goto l2027
							}
							if !_rules[rulehexdigit]() {
								goto l2027
							}
							add(rulePegText, position2029)
						}
						{
							position2030, tokenIndex2030 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2031
							}
							position++
							goto l2030
						l2031:
							position, tokenIndex = position2030, tokenIndex2030
							if buffer[position] != rune('H') {
								goto l2027
							}
							position++
						}
					l2030:
						{
							add(ruleAction40, position)
						}
						add(rulehexWordH, position2028)
					}
					goto l2022
				l2027:
					position, tokenIndex = position2022, tokenIndex2022
					{
						position2033 := position
						if buffer[position] != rune('0') {
							goto l2020
						}
						position++
						{
							position2034, tokenIndex2034 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2035
							}
							position++
							goto l2034
						l2035:
							position, tokenIndex = position2034, tokenIndex2034
							if buffer[position] != rune('X') {
								goto l2020
							}
							position++
						}
					l2034:
						{
							position2036 := position
							if !_rules[rulehexdigit]() {
								goto l2020
							}
							if !_rules[rulehexdigit]() {
								goto l2020
							}
							if !_rules[rulehexdigit]() {
								goto l2020
							}
							if !_rules[rulehexdigit]() {
								goto l2020
							}
							add(rulePegText, position2036)
						}
						{
							add(ruleAction41, position)
						}
						add(rulehexWord0x, position2033)
					}
				}
			l2022:
				add(rulenn, position2021)
			}
			return true
		l2020:
			position, tokenIndex = position2020, tokenIndex2020
			return false
		},
		/* 52 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position2038, tokenIndex2038 := position, tokenIndex
			{
				position2039 := position
				{
					position2040, tokenIndex2040 := position, tokenIndex
					{
						position2042 := position
						{
							position2043 := position
							{
								position2044, tokenIndex2044 := position, tokenIndex
								{
									position2046, tokenIndex2046 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2047
									}
									position++
									goto l2046
								l2047:
									position, tokenIndex = position2046, tokenIndex2046
									if buffer[position] != rune('+') {
										goto l2044
									}
									position++
								}
							l2046:
								goto l2045
							l2044:
								position, tokenIndex = position2044, tokenIndex2044
							}
						l2045:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2041
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2041
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2041
									}
									position++
									break
								}
							}

						l2048:
							{
								position2049, tokenIndex2049 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2049
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2049
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2049
										}
										position++
										break
									}
								}

								goto l2048
							l2049:
								position, tokenIndex = position2049, tokenIndex2049
							}
							add(rulePegText, position2043)
						}
						{
							position2052, tokenIndex2052 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2053
							}
							position++
							goto l2052
						l2053:
							position, tokenIndex = position2052, tokenIndex2052
							if buffer[position] != rune('H') {
								goto l2041
							}
							position++
						}
					l2052:
						{
							add(ruleAction34, position)
						}
						add(rulesignedHexByteH, position2042)
					}
					goto l2040
				l2041:
					position, tokenIndex = position2040, tokenIndex2040
					{
						position2056 := position
						{
							position2057 := position
							{
								position2058, tokenIndex2058 := position, tokenIndex
								{
									position2060, tokenIndex2060 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2061
									}
									position++
									goto l2060
								l2061:
									position, tokenIndex = position2060, tokenIndex2060
									if buffer[position] != rune('+') {
										goto l2058
									}
									position++
								}
							l2060:
								goto l2059
							l2058:
								position, tokenIndex = position2058, tokenIndex2058
							}
						l2059:
							if buffer[position] != rune('0') {
								goto l2055
							}
							position++
							{
								position2062, tokenIndex2062 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l2063
								}
								position++
								goto l2062
							l2063:
								position, tokenIndex = position2062, tokenIndex2062
								if buffer[position] != rune('X') {
									goto l2055
								}
								position++
							}
						l2062:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2055
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2055
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2055
									}
									position++
									break
								}
							}

						l2064:
							{
								position2065, tokenIndex2065 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2065
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2065
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2065
										}
										position++
										break
									}
								}

								goto l2064
							l2065:
								position, tokenIndex = position2065, tokenIndex2065
							}
							add(rulePegText, position2057)
						}
						{
							add(ruleAction35, position)
						}
						add(rulesignedHexByte0x, position2056)
					}
					goto l2040
				l2055:
					position, tokenIndex = position2040, tokenIndex2040
					{
						position2069 := position
						{
							position2070 := position
							{
								position2071, tokenIndex2071 := position, tokenIndex
								{
									position2073, tokenIndex2073 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2074
									}
									position++
									goto l2073
								l2074:
									position, tokenIndex = position2073, tokenIndex2073
									if buffer[position] != rune('+') {
										goto l2071
									}
									position++
								}
							l2073:
								goto l2072
							l2071:
								position, tokenIndex = position2071, tokenIndex2071
							}
						l2072:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2038
							}
							position++
						l2075:
							{
								position2076, tokenIndex2076 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2076
								}
								position++
								goto l2075
							l2076:
								position, tokenIndex = position2076, tokenIndex2076
							}
							add(rulePegText, position2070)
						}
						{
							add(ruleAction33, position)
						}
						add(rulesignedDecimalByte, position2069)
					}
				}
			l2040:
				add(ruledisp, position2039)
			}
			return true
		l2038:
			position, tokenIndex = position2038, tokenIndex2038
			return false
		},
		/* 53 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action33)> */
		nil,
		/* 54 signedHexByteH <- <(<(('-' / '+')? ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> ('h' / 'H') Action34)> */
		nil,
		/* 55 signedHexByte0x <- <(<(('-' / '+')? ('0' ('x' / 'X')) ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> Action35)> */
		nil,
		/* 56 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action36)> */
		nil,
		/* 57 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action37)> */
		nil,
		/* 58 decimalByte <- <(<[0-9]+> Action38)> */
		nil,
		/* 59 LabelNN <- <(<LabelText> Action39)> */
		nil,
		/* 60 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action40)> */
		nil,
		/* 61 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action41)> */
		nil,
		/* 62 nn_contents <- <('(' nn ')' Action42)> */
		func() bool {
			position2087, tokenIndex2087 := position, tokenIndex
			{
				position2088 := position
				if buffer[position] != rune('(') {
					goto l2087
				}
				position++
				if !_rules[rulenn]() {
					goto l2087
				}
				if buffer[position] != rune(')') {
					goto l2087
				}
				position++
				{
					add(ruleAction42, position)
				}
				add(rulenn_contents, position2088)
			}
			return true
		l2087:
			position, tokenIndex = position2087, tokenIndex2087
			return false
		},
		/* 63 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 64 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action43)> */
		nil,
		/* 65 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action44)> */
		nil,
		/* 66 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action45)> */
		nil,
		/* 67 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action46)> */
		nil,
		/* 68 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action47)> */
		nil,
		/* 69 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action48)> */
		nil,
		/* 70 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action49)> */
		nil,
		/* 71 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action50)> */
		nil,
		/* 72 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 73 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 74 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 (sep Copy8)? Action51)> */
		nil,
		/* 75 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 (sep Copy8)? Action52)> */
		nil,
		/* 76 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action53)> */
		nil,
		/* 77 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 (sep Copy8)? Action54)> */
		nil,
		/* 78 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 (sep Copy8)? Action55)> */
		nil,
		/* 79 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 (sep Copy8)? Action56)> */
		nil,
		/* 80 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 (sep Copy8)? Action57)> */
		nil,
		/* 81 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action58)> */
		nil,
		/* 82 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action59)> */
		nil,
		/* 83 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 (sep Copy8)? Action60)> */
		nil,
		/* 84 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 (sep Copy8)? Action61)> */
		nil,
		/* 85 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 86 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action62)> */
		nil,
		/* 87 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action63)> */
		nil,
		/* 88 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action64)> */
		nil,
		/* 89 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action65)> */
		nil,
		/* 90 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action66)> */
		nil,
		/* 91 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action67)> */
		nil,
		/* 92 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action68)> */
		nil,
		/* 93 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action69)> */
		nil,
		/* 94 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action70)> */
		nil,
		/* 95 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action71)> */
		nil,
		/* 96 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action72)> */
		nil,
		/* 97 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action73)> */
		nil,
		/* 98 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action74)> */
		nil,
		/* 99 EDSimple <- <(Retn / Reti / Rrd / Im0 / Im1 / Im2 / ((&('I' | 'O' | 'i' | 'o') BlitIO) | (&('R' | 'r') Rld) | (&('N' | 'n') Neg) | (&('C' | 'L' | 'c' | 'l') Blit)))> */
		nil,
		/* 100 Neg <- <(<(('n' / 'N') ('e' / 'E') ('g' / 'G'))> Action75)> */
		nil,
		/* 101 Retn <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N'))> Action76)> */
		nil,
		/* 102 Reti <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('i' / 'I'))> Action77)> */
		nil,
		/* 103 Rrd <- <(<(('r' / 'R') ('r' / 'R') ('d' / 'D'))> Action78)> */
		nil,
		/* 104 Rld <- <(<(('r' / 'R') ('l' / 'L') ('d' / 'D'))> Action79)> */
		nil,
		/* 105 Im0 <- <(<(('i' / 'I') ('m' / 'M') ' ' '0')> Action80)> */
		nil,
		/* 106 Im1 <- <(<(('i' / 'I') ('m' / 'M') ' ' '1')> Action81)> */
		nil,
		/* 107 Im2 <- <(<(('i' / 'I') ('m' / 'M') ' ' '2')> Action82)> */
		nil,
		/* 108 Blit <- <(Ldir / Ldi / Cpir / Cpi / Lddr / Ldd / Cpdr / Cpd)> */
		nil,
		/* 109 BlitIO <- <(Inir / Ini / Otir / Outi / Indr / Ind / Otdr / Outd)> */
		nil,
		/* 110 Ldi <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I'))> Action83)> */
		nil,
		/* 111 Cpi <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I'))> Action84)> */
		nil,
		/* 112 Ini <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I'))> Action85)> */
		nil,
		/* 113 Outi <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('i' / 'I'))> Action86)> */
		nil,
		/* 114 Ldd <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D'))> Action87)> */
		nil,
		/* 115 Cpd <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D'))> Action88)> */
		nil,
		/* 116 Ind <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D'))> Action89)> */
		nil,
		/* 117 Outd <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('d' / 'D'))> Action90)> */
		nil,
		/* 118 Ldir <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I') ('r' / 'R'))> Action91)> */
		nil,
		/* 119 Cpir <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I') ('r' / 'R'))> Action92)> */
		nil,
		/* 120 Inir <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I') ('r' / 'R'))> Action93)> */
		nil,
		/* 121 Otir <- <(<(('o' / 'O') ('t' / 'T') ('i' / 'I') ('r' / 'R'))> Action94)> */
		nil,
		/* 122 Lddr <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D') ('r' / 'R'))> Action95)> */
		nil,
		/* 123 Cpdr <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D') ('r' / 'R'))> Action96)> */
		nil,
		/* 124 Indr <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D') ('r' / 'R'))> Action97)> */
		nil,
		/* 125 Otdr <- <(<(('o' / 'O') ('t' / 'T') ('d' / 'D') ('r' / 'R'))> Action98)> */
		nil,
		/* 126 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 127 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action99)> */
		nil,
		/* 128 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action100)> */
		nil,
		/* 129 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action101)> */
		nil,
		/* 130 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action102)> */
		nil,
		/* 131 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action103)> */
		nil,
		/* 132 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action104)> */
		nil,
		/* 133 IO <- <(IN / OUT)> */
		nil,
		/* 134 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action105)> */
		nil,
		/* 135 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action106)> */
		nil,
		/* 136 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position2163, tokenIndex2163 := position, tokenIndex
			{
				position2164 := position
				{
					position2165, tokenIndex2165 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2166
					}
					position++
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
							goto l2166
						}
						position++
					}
				l2167:
					if buffer[position] != rune(')') {
						goto l2166
					}
					position++
					goto l2165
				l2166:
					position, tokenIndex = position2165, tokenIndex2165
					if buffer[position] != rune('(') {
						goto l2163
					}
					position++
					if !_rules[rulen]() {
						goto l2163
					}
					if buffer[position] != rune(')') {
						goto l2163
					}
					position++
				}
			l2165:
				add(rulePort, position2164)
			}
			return true
		l2163:
			position, tokenIndex = position2163, tokenIndex2163
			return false
		},
		/* 137 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2169, tokenIndex2169 := position, tokenIndex
			{
				position2170 := position
				{
					position2171, tokenIndex2171 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2171
					}
					goto l2172
				l2171:
					position, tokenIndex = position2171, tokenIndex2171
				}
			l2172:
				if buffer[position] != rune(',') {
					goto l2169
				}
				position++
				{
					position2173, tokenIndex2173 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2173
					}
					goto l2174
				l2173:
					position, tokenIndex = position2173, tokenIndex2173
				}
			l2174:
				add(rulesep, position2170)
			}
			return true
		l2169:
			position, tokenIndex = position2169, tokenIndex2169
			return false
		},
		/* 138 ws <- <(' ' / '\t')+> */
		func() bool {
			position2175, tokenIndex2175 := position, tokenIndex
			{
				position2176 := position
				{
					position2179, tokenIndex2179 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2180
					}
					position++
					goto l2179
				l2180:
					position, tokenIndex = position2179, tokenIndex2179
					if buffer[position] != rune('\t') {
						goto l2175
					}
					position++
				}
			l2179:
			l2177:
				{
					position2178, tokenIndex2178 := position, tokenIndex
					{
						position2181, tokenIndex2181 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l2182
						}
						position++
						goto l2181
					l2182:
						position, tokenIndex = position2181, tokenIndex2181
						if buffer[position] != rune('\t') {
							goto l2178
						}
						position++
					}
				l2181:
					goto l2177
				l2178:
					position, tokenIndex = position2178, tokenIndex2178
				}
				add(rulews, position2176)
			}
			return true
		l2175:
			position, tokenIndex = position2175, tokenIndex2175
			return false
		},
		/* 139 A <- <('a' / 'A')> */
		nil,
		/* 140 F <- <('f' / 'F')> */
		nil,
		/* 141 B <- <('b' / 'B')> */
		nil,
		/* 142 C <- <('c' / 'C')> */
		nil,
		/* 143 D <- <('d' / 'D')> */
		nil,
		/* 144 E <- <('e' / 'E')> */
		nil,
		/* 145 H <- <('h' / 'H')> */
		nil,
		/* 146 L <- <('l' / 'L')> */
		nil,
		/* 147 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 148 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 149 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 150 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 151 I <- <('i' / 'I')> */
		nil,
		/* 152 R <- <('r' / 'R')> */
		nil,
		/* 153 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 154 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 155 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 156 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 157 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 158 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 159 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 160 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 161 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position2205, tokenIndex2205 := position, tokenIndex
			{
				position2206 := position
				{
					position2207, tokenIndex2207 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2208
					}
					position++
					goto l2207
				l2208:
					position, tokenIndex = position2207, tokenIndex2207
					{
						position2209, tokenIndex2209 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2210
						}
						position++
						goto l2209
					l2210:
						position, tokenIndex = position2209, tokenIndex2209
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2205
						}
						position++
					}
				l2209:
				}
			l2207:
				add(rulehexdigit, position2206)
			}
			return true
		l2205:
			position, tokenIndex = position2205, tokenIndex2205
			return false
		},
		/* 162 octaldigit <- <(<[0-7]> Action107)> */
		func() bool {
			position2211, tokenIndex2211 := position, tokenIndex
			{
				position2212 := position
				{
					position2213 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2211
					}
					position++
					add(rulePegText, position2213)
				}
				{
					add(ruleAction107, position)
				}
				add(ruleoctaldigit, position2212)
			}
			return true
		l2211:
			position, tokenIndex = position2211, tokenIndex2211
			return false
		},
		/* 163 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2215, tokenIndex2215 := position, tokenIndex
			{
				position2216 := position
				{
					position2217, tokenIndex2217 := position, tokenIndex
					{
						position2219 := position
						{
							position2220, tokenIndex2220 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2221
							}
							position++
							goto l2220
						l2221:
							position, tokenIndex = position2220, tokenIndex2220
							if buffer[position] != rune('N') {
								goto l2218
							}
							position++
						}
					l2220:
						{
							position2222, tokenIndex2222 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2223
							}
							position++
							goto l2222
						l2223:
							position, tokenIndex = position2222, tokenIndex2222
							if buffer[position] != rune('Z') {
								goto l2218
							}
							position++
						}
					l2222:
						{
							add(ruleAction108, position)
						}
						add(ruleFT_NZ, position2219)
					}
					goto l2217
				l2218:
					position, tokenIndex = position2217, tokenIndex2217
					{
						position2226 := position
						{
							position2227, tokenIndex2227 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2228
							}
							position++
							goto l2227
						l2228:
							position, tokenIndex = position2227, tokenIndex2227
							if buffer[position] != rune('P') {
								goto l2225
							}
							position++
						}
					l2227:
						{
							position2229, tokenIndex2229 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2230
							}
							position++
							goto l2229
						l2230:
							position, tokenIndex = position2229, tokenIndex2229
							if buffer[position] != rune('O') {
								goto l2225
							}
							position++
						}
					l2229:
						{
							add(ruleAction112, position)
						}
						add(ruleFT_PO, position2226)
					}
					goto l2217
				l2225:
					position, tokenIndex = position2217, tokenIndex2217
					{
						position2233 := position
						{
							position2234, tokenIndex2234 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2235
							}
							position++
							goto l2234
						l2235:
							position, tokenIndex = position2234, tokenIndex2234
							if buffer[position] != rune('P') {
								goto l2232
							}
							position++
						}
					l2234:
						{
							position2236, tokenIndex2236 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2237
							}
							position++
							goto l2236
						l2237:
							position, tokenIndex = position2236, tokenIndex2236
							if buffer[position] != rune('E') {
								goto l2232
							}
							position++
						}
					l2236:
						{
							add(ruleAction113, position)
						}
						add(ruleFT_PE, position2233)
					}
					goto l2217
				l2232:
					position, tokenIndex = position2217, tokenIndex2217
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2240 := position
								{
									position2241, tokenIndex2241 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2242
									}
									position++
									goto l2241
								l2242:
									position, tokenIndex = position2241, tokenIndex2241
									if buffer[position] != rune('M') {
										goto l2215
									}
									position++
								}
							l2241:
								{
									add(ruleAction115, position)
								}
								add(ruleFT_M, position2240)
							}
							break
						case 'P', 'p':
							{
								position2244 := position
								{
									position2245, tokenIndex2245 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2246
									}
									position++
									goto l2245
								l2246:
									position, tokenIndex = position2245, tokenIndex2245
									if buffer[position] != rune('P') {
										goto l2215
									}
									position++
								}
							l2245:
								{
									add(ruleAction114, position)
								}
								add(ruleFT_P, position2244)
							}
							break
						case 'C', 'c':
							{
								position2248 := position
								{
									position2249, tokenIndex2249 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2250
									}
									position++
									goto l2249
								l2250:
									position, tokenIndex = position2249, tokenIndex2249
									if buffer[position] != rune('C') {
										goto l2215
									}
									position++
								}
							l2249:
								{
									add(ruleAction111, position)
								}
								add(ruleFT_C, position2248)
							}
							break
						case 'N', 'n':
							{
								position2252 := position
								{
									position2253, tokenIndex2253 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2254
									}
									position++
									goto l2253
								l2254:
									position, tokenIndex = position2253, tokenIndex2253
									if buffer[position] != rune('N') {
										goto l2215
									}
									position++
								}
							l2253:
								{
									position2255, tokenIndex2255 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2256
									}
									position++
									goto l2255
								l2256:
									position, tokenIndex = position2255, tokenIndex2255
									if buffer[position] != rune('C') {
										goto l2215
									}
									position++
								}
							l2255:
								{
									add(ruleAction110, position)
								}
								add(ruleFT_NC, position2252)
							}
							break
						default:
							{
								position2258 := position
								{
									position2259, tokenIndex2259 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2260
									}
									position++
									goto l2259
								l2260:
									position, tokenIndex = position2259, tokenIndex2259
									if buffer[position] != rune('Z') {
										goto l2215
									}
									position++
								}
							l2259:
								{
									add(ruleAction109, position)
								}
								add(ruleFT_Z, position2258)
							}
							break
						}
					}

				}
			l2217:
				add(rulecc, position2216)
			}
			return true
		l2215:
			position, tokenIndex = position2215, tokenIndex2215
			return false
		},
		/* 164 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action108)> */
		nil,
		/* 165 FT_Z <- <(('z' / 'Z') Action109)> */
		nil,
		/* 166 FT_NC <- <(('n' / 'N') ('c' / 'C') Action110)> */
		nil,
		/* 167 FT_C <- <(('c' / 'C') Action111)> */
		nil,
		/* 168 FT_PO <- <(('p' / 'P') ('o' / 'O') Action112)> */
		nil,
		/* 169 FT_PE <- <(('p' / 'P') ('e' / 'E') Action113)> */
		nil,
		/* 170 FT_P <- <(('p' / 'P') Action114)> */
		nil,
		/* 171 FT_M <- <(('m' / 'M') Action115)> */
		nil,
		/* 173 Action0 <- <{ p.Emit() }> */
		nil,
		/* 174 Action1 <- <{ p.Org() }> */
		nil,
		/* 175 Action2 <- <{ p.DefByte() }> */
		nil,
		/* 176 Action3 <- <{ p.DefSpace() }> */
		nil,
		/* 177 Action4 <- <{ p.LabelDefn(buffer[begin:end])}> */
		nil,
		nil,
		/* 179 Action5 <- <{ p.LD8() }> */
		nil,
		/* 180 Action6 <- <{ p.LD16() }> */
		nil,
		/* 181 Action7 <- <{ p.Push() }> */
		nil,
		/* 182 Action8 <- <{ p.Pop() }> */
		nil,
		/* 183 Action9 <- <{ p.Ex() }> */
		nil,
		/* 184 Action10 <- <{ p.Inc8() }> */
		nil,
		/* 185 Action11 <- <{ p.Inc8() }> */
		nil,
		/* 186 Action12 <- <{ p.Inc16() }> */
		nil,
		/* 187 Action13 <- <{ p.Dec8() }> */
		nil,
		/* 188 Action14 <- <{ p.Dec8() }> */
		nil,
		/* 189 Action15 <- <{ p.Dec16() }> */
		nil,
		/* 190 Action16 <- <{ p.Add16() }> */
		nil,
		/* 191 Action17 <- <{ p.Adc16() }> */
		nil,
		/* 192 Action18 <- <{ p.Sbc16() }> */
		nil,
		/* 193 Action19 <- <{ p.Dst8() }> */
		nil,
		/* 194 Action20 <- <{ p.Src8() }> */
		nil,
		/* 195 Action21 <- <{ p.Loc8() }> */
		nil,
		/* 196 Action22 <- <{ p.Copy8() }> */
		nil,
		/* 197 Action23 <- <{ p.Loc8() }> */
		nil,
		/* 198 Action24 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 199 Action25 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 200 Action26 <- <{ p.Dst16() }> */
		nil,
		/* 201 Action27 <- <{ p.Src16() }> */
		nil,
		/* 202 Action28 <- <{ p.Loc16() }> */
		nil,
		/* 203 Action29 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 204 Action30 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 205 Action31 <- <{ p.R16Contents() }> */
		nil,
		/* 206 Action32 <- <{ p.IR16Contents() }> */
		nil,
		/* 207 Action33 <- <{ p.DispDecimal(buffer[begin:end]) }> */
		nil,
		/* 208 Action34 <- <{ p.DispHex(buffer[begin:end]) }> */
		nil,
		/* 209 Action35 <- <{ p.Disp0xHex(buffer[begin:end]) }> */
		nil,
		/* 210 Action36 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 211 Action37 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 212 Action38 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 213 Action39 <- <{ p.NNLabel(buffer[begin:end]) }> */
		nil,
		/* 214 Action40 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 215 Action41 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 216 Action42 <- <{ p.NNContents() }> */
		nil,
		/* 217 Action43 <- <{ p.Accum("ADD") }> */
		nil,
		/* 218 Action44 <- <{ p.Accum("ADC") }> */
		nil,
		/* 219 Action45 <- <{ p.Accum("SUB") }> */
		nil,
		/* 220 Action46 <- <{ p.Accum("SBC") }> */
		nil,
		/* 221 Action47 <- <{ p.Accum("AND") }> */
		nil,
		/* 222 Action48 <- <{ p.Accum("XOR") }> */
		nil,
		/* 223 Action49 <- <{ p.Accum("OR") }> */
		nil,
		/* 224 Action50 <- <{ p.Accum("CP") }> */
		nil,
		/* 225 Action51 <- <{ p.Rot("RLC") }> */
		nil,
		/* 226 Action52 <- <{ p.Rot("RRC") }> */
		nil,
		/* 227 Action53 <- <{ p.Rot("RL") }> */
		nil,
		/* 228 Action54 <- <{ p.Rot("RR") }> */
		nil,
		/* 229 Action55 <- <{ p.Rot("SLA") }> */
		nil,
		/* 230 Action56 <- <{ p.Rot("SRA") }> */
		nil,
		/* 231 Action57 <- <{ p.Rot("SLL") }> */
		nil,
		/* 232 Action58 <- <{ p.Rot("SRL") }> */
		nil,
		/* 233 Action59 <- <{ p.Bit() }> */
		nil,
		/* 234 Action60 <- <{ p.Res() }> */
		nil,
		/* 235 Action61 <- <{ p.Set() }> */
		nil,
		/* 236 Action62 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 237 Action63 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 238 Action64 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 239 Action65 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 240 Action66 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 241 Action67 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 242 Action68 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 243 Action69 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 244 Action70 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 245 Action71 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 246 Action72 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 247 Action73 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 248 Action74 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 249 Action75 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 250 Action76 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 251 Action77 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 252 Action78 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 253 Action79 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 254 Action80 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 255 Action81 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 256 Action82 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 257 Action83 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 258 Action84 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 259 Action85 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 260 Action86 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 261 Action87 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 262 Action88 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 263 Action89 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 264 Action90 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 265 Action91 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 266 Action92 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 267 Action93 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 268 Action94 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 269 Action95 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 270 Action96 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 271 Action97 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 272 Action98 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 273 Action99 <- <{ p.Rst() }> */
		nil,
		/* 274 Action100 <- <{ p.Call() }> */
		nil,
		/* 275 Action101 <- <{ p.Ret() }> */
		nil,
		/* 276 Action102 <- <{ p.Jp() }> */
		nil,
		/* 277 Action103 <- <{ p.Jr() }> */
		nil,
		/* 278 Action104 <- <{ p.Djnz() }> */
		nil,
		/* 279 Action105 <- <{ p.In() }> */
		nil,
		/* 280 Action106 <- <{ p.Out() }> */
		nil,
		/* 281 Action107 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 282 Action108 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 283 Action109 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 284 Action110 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 285 Action111 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 286 Action112 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 287 Action113 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 288 Action114 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 289 Action115 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
