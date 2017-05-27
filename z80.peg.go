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
	ruleDefw
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
	rulehexWord
	rulezeroHexWord
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
	rulePegText
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
	"Defw",
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
	"hexWord",
	"zeroHexWord",
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
	"PegText",
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
	rules  [293]func() bool
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
			p.DefWord()
		case ruleAction4:
			p.DefSpace()
		case ruleAction5:
			p.LabelDefn(buffer[begin:end])
		case ruleAction6:
			p.LD8()
		case ruleAction7:
			p.LD16()
		case ruleAction8:
			p.Push()
		case ruleAction9:
			p.Pop()
		case ruleAction10:
			p.Ex()
		case ruleAction11:
			p.Inc8()
		case ruleAction12:
			p.Inc8()
		case ruleAction13:
			p.Inc16()
		case ruleAction14:
			p.Dec8()
		case ruleAction15:
			p.Dec8()
		case ruleAction16:
			p.Dec16()
		case ruleAction17:
			p.Add16()
		case ruleAction18:
			p.Adc16()
		case ruleAction19:
			p.Sbc16()
		case ruleAction20:
			p.Dst8()
		case ruleAction21:
			p.Src8()
		case ruleAction22:
			p.Loc8()
		case ruleAction23:
			p.Copy8()
		case ruleAction24:
			p.Loc8()
		case ruleAction25:
			p.R8(buffer[begin:end])
		case ruleAction26:
			p.R8(buffer[begin:end])
		case ruleAction27:
			p.Dst16()
		case ruleAction28:
			p.Src16()
		case ruleAction29:
			p.Loc16()
		case ruleAction30:
			p.R16(buffer[begin:end])
		case ruleAction31:
			p.R16(buffer[begin:end])
		case ruleAction32:
			p.R16Contents()
		case ruleAction33:
			p.IR16Contents()
		case ruleAction34:
			p.DispDecimal(buffer[begin:end])
		case ruleAction35:
			p.DispHex(buffer[begin:end])
		case ruleAction36:
			p.Disp0xHex(buffer[begin:end])
		case ruleAction37:
			p.Nhex(buffer[begin:end])
		case ruleAction38:
			p.Nhex(buffer[begin:end])
		case ruleAction39:
			p.Ndec(buffer[begin:end])
		case ruleAction40:
			p.NNLabel(buffer[begin:end])
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
								add(ruleAction5, position)
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
													position32, tokenIndex32 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l33
													}
													position++
													goto l32
												l33:
													position, tokenIndex = position32, tokenIndex32
													if buffer[position] != rune('D') {
														goto l20
													}
													position++
												}
											l32:
												{
													position34, tokenIndex34 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l35
													}
													position++
													goto l34
												l35:
													position, tokenIndex = position34, tokenIndex34
													if buffer[position] != rune('B') {
														goto l20
													}
													position++
												}
											l34:
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
											position38 := position
											{
												position39, tokenIndex39 := position, tokenIndex
												{
													position41, tokenIndex41 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l42
													}
													position++
													goto l41
												l42:
													position, tokenIndex = position41, tokenIndex41
													if buffer[position] != rune('D') {
														goto l40
													}
													position++
												}
											l41:
												{
													position43, tokenIndex43 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l44
													}
													position++
													goto l43
												l44:
													position, tokenIndex = position43, tokenIndex43
													if buffer[position] != rune('E') {
														goto l40
													}
													position++
												}
											l43:
												{
													position45, tokenIndex45 := position, tokenIndex
													if buffer[position] != rune('f') {
														goto l46
													}
													position++
													goto l45
												l46:
													position, tokenIndex = position45, tokenIndex45
													if buffer[position] != rune('F') {
														goto l40
													}
													position++
												}
											l45:
												{
													position47, tokenIndex47 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l48
													}
													position++
													goto l47
												l48:
													position, tokenIndex = position47, tokenIndex47
													if buffer[position] != rune('S') {
														goto l40
													}
													position++
												}
											l47:
												goto l39
											l40:
												position, tokenIndex = position39, tokenIndex39
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
														goto l37
													}
													position++
												}
											l49:
												{
													position51, tokenIndex51 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l52
													}
													position++
													goto l51
												l52:
													position, tokenIndex = position51, tokenIndex51
													if buffer[position] != rune('S') {
														goto l37
													}
													position++
												}
											l51:
											}
										l39:
											if !_rules[rulews]() {
												goto l37
											}
											if !_rules[rulen]() {
												goto l37
											}
											{
												add(ruleAction4, position)
											}
											add(ruleDefs, position38)
										}
										goto l19
									l37:
										position, tokenIndex = position19, tokenIndex19
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position55 := position
													{
														position56, tokenIndex56 := position, tokenIndex
														{
															position58, tokenIndex58 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l59
															}
															position++
															goto l58
														l59:
															position, tokenIndex = position58, tokenIndex58
															if buffer[position] != rune('D') {
																goto l57
															}
															position++
														}
													l58:
														{
															position60, tokenIndex60 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l61
															}
															position++
															goto l60
														l61:
															position, tokenIndex = position60, tokenIndex60
															if buffer[position] != rune('E') {
																goto l57
															}
															position++
														}
													l60:
														{
															position62, tokenIndex62 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l63
															}
															position++
															goto l62
														l63:
															position, tokenIndex = position62, tokenIndex62
															if buffer[position] != rune('F') {
																goto l57
															}
															position++
														}
													l62:
														{
															position64, tokenIndex64 := position, tokenIndex
															if buffer[position] != rune('w') {
																goto l65
															}
															position++
															goto l64
														l65:
															position, tokenIndex = position64, tokenIndex64
															if buffer[position] != rune('W') {
																goto l57
															}
															position++
														}
													l64:
														goto l56
													l57:
														position, tokenIndex = position56, tokenIndex56
														{
															position66, tokenIndex66 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l67
															}
															position++
															goto l66
														l67:
															position, tokenIndex = position66, tokenIndex66
															if buffer[position] != rune('D') {
																goto l17
															}
															position++
														}
													l66:
														{
															position68, tokenIndex68 := position, tokenIndex
															if buffer[position] != rune('w') {
																goto l69
															}
															position++
															goto l68
														l69:
															position, tokenIndex = position68, tokenIndex68
															if buffer[position] != rune('W') {
																goto l17
															}
															position++
														}
													l68:
													}
												l56:
													if !_rules[rulews]() {
														goto l17
													}
													if !_rules[rulenn]() {
														goto l17
													}
													{
														add(ruleAction3, position)
													}
													add(ruleDefw, position55)
												}
												break
											case 'O', 'o':
												{
													position71 := position
													{
														position72, tokenIndex72 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l73
														}
														position++
														goto l72
													l73:
														position, tokenIndex = position72, tokenIndex72
														if buffer[position] != rune('O') {
															goto l17
														}
														position++
													}
												l72:
													{
														position74, tokenIndex74 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l75
														}
														position++
														goto l74
													l75:
														position, tokenIndex = position74, tokenIndex74
														if buffer[position] != rune('R') {
															goto l17
														}
														position++
													}
												l74:
													{
														position76, tokenIndex76 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l77
														}
														position++
														goto l76
													l77:
														position, tokenIndex = position76, tokenIndex76
														if buffer[position] != rune('G') {
															goto l17
														}
														position++
													}
												l76:
													if !_rules[rulews]() {
														goto l17
													}
													if !_rules[rulenn]() {
														goto l17
													}
													{
														add(ruleAction1, position)
													}
													add(ruleOrg, position71)
												}
												break
											case 'a':
												{
													position79 := position
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
													add(ruleAseg, position79)
												}
												break
											default:
												{
													position80 := position
													{
														position81, tokenIndex81 := position, tokenIndex
														if buffer[position] != rune('.') {
															goto l81
														}
														position++
														goto l82
													l81:
														position, tokenIndex = position81, tokenIndex81
													}
												l82:
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
												l83:
													{
														position84, tokenIndex84 := position, tokenIndex
														{
															position85, tokenIndex85 := position, tokenIndex
															if buffer[position] != rune('\'') {
																goto l85
															}
															position++
															goto l84
														l85:
															position, tokenIndex = position85, tokenIndex85
														}
														if !matchDot() {
															goto l84
														}
														goto l83
													l84:
														position, tokenIndex = position84, tokenIndex84
													}
													if buffer[position] != rune('\'') {
														goto l17
													}
													position++
													add(ruleTitle, position80)
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
									position86 := position
									{
										position87, tokenIndex87 := position, tokenIndex
										{
											position89 := position
											{
												position90, tokenIndex90 := position, tokenIndex
												{
													position92 := position
													{
														position93, tokenIndex93 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l94
														}
														position++
														goto l93
													l94:
														position, tokenIndex = position93, tokenIndex93
														if buffer[position] != rune('P') {
															goto l91
														}
														position++
													}
												l93:
													{
														position95, tokenIndex95 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l96
														}
														position++
														goto l95
													l96:
														position, tokenIndex = position95, tokenIndex95
														if buffer[position] != rune('U') {
															goto l91
														}
														position++
													}
												l95:
													{
														position97, tokenIndex97 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l98
														}
														position++
														goto l97
													l98:
														position, tokenIndex = position97, tokenIndex97
														if buffer[position] != rune('S') {
															goto l91
														}
														position++
													}
												l97:
													{
														position99, tokenIndex99 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l100
														}
														position++
														goto l99
													l100:
														position, tokenIndex = position99, tokenIndex99
														if buffer[position] != rune('H') {
															goto l91
														}
														position++
													}
												l99:
													if !_rules[rulews]() {
														goto l91
													}
													if !_rules[ruleSrc16]() {
														goto l91
													}
													{
														add(ruleAction8, position)
													}
													add(rulePush, position92)
												}
												goto l90
											l91:
												position, tokenIndex = position90, tokenIndex90
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position103 := position
															{
																position104, tokenIndex104 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l105
																}
																position++
																goto l104
															l105:
																position, tokenIndex = position104, tokenIndex104
																if buffer[position] != rune('E') {
																	goto l88
																}
																position++
															}
														l104:
															{
																position106, tokenIndex106 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l107
																}
																position++
																goto l106
															l107:
																position, tokenIndex = position106, tokenIndex106
																if buffer[position] != rune('X') {
																	goto l88
																}
																position++
															}
														l106:
															if !_rules[rulews]() {
																goto l88
															}
															if !_rules[ruleDst16]() {
																goto l88
															}
															if !_rules[rulesep]() {
																goto l88
															}
															if !_rules[ruleSrc16]() {
																goto l88
															}
															{
																add(ruleAction10, position)
															}
															add(ruleEx, position103)
														}
														break
													case 'P', 'p':
														{
															position109 := position
															{
																position110, tokenIndex110 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l111
																}
																position++
																goto l110
															l111:
																position, tokenIndex = position110, tokenIndex110
																if buffer[position] != rune('P') {
																	goto l88
																}
																position++
															}
														l110:
															{
																position112, tokenIndex112 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l113
																}
																position++
																goto l112
															l113:
																position, tokenIndex = position112, tokenIndex112
																if buffer[position] != rune('O') {
																	goto l88
																}
																position++
															}
														l112:
															{
																position114, tokenIndex114 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l115
																}
																position++
																goto l114
															l115:
																position, tokenIndex = position114, tokenIndex114
																if buffer[position] != rune('P') {
																	goto l88
																}
																position++
															}
														l114:
															if !_rules[rulews]() {
																goto l88
															}
															if !_rules[ruleDst16]() {
																goto l88
															}
															{
																add(ruleAction9, position)
															}
															add(rulePop, position109)
														}
														break
													default:
														{
															position117 := position
															{
																position118, tokenIndex118 := position, tokenIndex
																{
																	position120 := position
																	{
																		position121, tokenIndex121 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l122
																		}
																		position++
																		goto l121
																	l122:
																		position, tokenIndex = position121, tokenIndex121
																		if buffer[position] != rune('L') {
																			goto l119
																		}
																		position++
																	}
																l121:
																	{
																		position123, tokenIndex123 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l124
																		}
																		position++
																		goto l123
																	l124:
																		position, tokenIndex = position123, tokenIndex123
																		if buffer[position] != rune('D') {
																			goto l119
																		}
																		position++
																	}
																l123:
																	if !_rules[rulews]() {
																		goto l119
																	}
																	if !_rules[ruleDst16]() {
																		goto l119
																	}
																	if !_rules[rulesep]() {
																		goto l119
																	}
																	if !_rules[ruleSrc16]() {
																		goto l119
																	}
																	{
																		add(ruleAction7, position)
																	}
																	add(ruleLoad16, position120)
																}
																goto l118
															l119:
																position, tokenIndex = position118, tokenIndex118
																{
																	position126 := position
																	{
																		position127, tokenIndex127 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l128
																		}
																		position++
																		goto l127
																	l128:
																		position, tokenIndex = position127, tokenIndex127
																		if buffer[position] != rune('L') {
																			goto l88
																		}
																		position++
																	}
																l127:
																	{
																		position129, tokenIndex129 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l130
																		}
																		position++
																		goto l129
																	l130:
																		position, tokenIndex = position129, tokenIndex129
																		if buffer[position] != rune('D') {
																			goto l88
																		}
																		position++
																	}
																l129:
																	if !_rules[rulews]() {
																		goto l88
																	}
																	{
																		position131 := position
																		{
																			position132, tokenIndex132 := position, tokenIndex
																			if !_rules[ruleReg8]() {
																				goto l133
																			}
																			goto l132
																		l133:
																			position, tokenIndex = position132, tokenIndex132
																			if !_rules[ruleReg16Contents]() {
																				goto l134
																			}
																			goto l132
																		l134:
																			position, tokenIndex = position132, tokenIndex132
																			if !_rules[rulenn_contents]() {
																				goto l88
																			}
																		}
																	l132:
																		{
																			add(ruleAction20, position)
																		}
																		add(ruleDst8, position131)
																	}
																	if !_rules[rulesep]() {
																		goto l88
																	}
																	if !_rules[ruleSrc8]() {
																		goto l88
																	}
																	{
																		add(ruleAction6, position)
																	}
																	add(ruleLoad8, position126)
																}
															}
														l118:
															add(ruleLoad, position117)
														}
														break
													}
												}

											}
										l90:
											add(ruleAssignment, position89)
										}
										goto l87
									l88:
										position, tokenIndex = position87, tokenIndex87
										{
											position138 := position
											{
												position139, tokenIndex139 := position, tokenIndex
												{
													position141 := position
													{
														position142, tokenIndex142 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l143
														}
														position++
														goto l142
													l143:
														position, tokenIndex = position142, tokenIndex142
														if buffer[position] != rune('I') {
															goto l140
														}
														position++
													}
												l142:
													{
														position144, tokenIndex144 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l145
														}
														position++
														goto l144
													l145:
														position, tokenIndex = position144, tokenIndex144
														if buffer[position] != rune('N') {
															goto l140
														}
														position++
													}
												l144:
													{
														position146, tokenIndex146 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l147
														}
														position++
														goto l146
													l147:
														position, tokenIndex = position146, tokenIndex146
														if buffer[position] != rune('C') {
															goto l140
														}
														position++
													}
												l146:
													if !_rules[rulews]() {
														goto l140
													}
													if !_rules[ruleILoc8]() {
														goto l140
													}
													{
														add(ruleAction11, position)
													}
													add(ruleInc16Indexed8, position141)
												}
												goto l139
											l140:
												position, tokenIndex = position139, tokenIndex139
												{
													position150 := position
													{
														position151, tokenIndex151 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l152
														}
														position++
														goto l151
													l152:
														position, tokenIndex = position151, tokenIndex151
														if buffer[position] != rune('I') {
															goto l149
														}
														position++
													}
												l151:
													{
														position153, tokenIndex153 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l154
														}
														position++
														goto l153
													l154:
														position, tokenIndex = position153, tokenIndex153
														if buffer[position] != rune('N') {
															goto l149
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
															goto l149
														}
														position++
													}
												l155:
													if !_rules[rulews]() {
														goto l149
													}
													if !_rules[ruleLoc16]() {
														goto l149
													}
													{
														add(ruleAction13, position)
													}
													add(ruleInc16, position150)
												}
												goto l139
											l149:
												position, tokenIndex = position139, tokenIndex139
												{
													position158 := position
													{
														position159, tokenIndex159 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l160
														}
														position++
														goto l159
													l160:
														position, tokenIndex = position159, tokenIndex159
														if buffer[position] != rune('I') {
															goto l137
														}
														position++
													}
												l159:
													{
														position161, tokenIndex161 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l162
														}
														position++
														goto l161
													l162:
														position, tokenIndex = position161, tokenIndex161
														if buffer[position] != rune('N') {
															goto l137
														}
														position++
													}
												l161:
													{
														position163, tokenIndex163 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l164
														}
														position++
														goto l163
													l164:
														position, tokenIndex = position163, tokenIndex163
														if buffer[position] != rune('C') {
															goto l137
														}
														position++
													}
												l163:
													if !_rules[rulews]() {
														goto l137
													}
													if !_rules[ruleLoc8]() {
														goto l137
													}
													{
														add(ruleAction12, position)
													}
													add(ruleInc8, position158)
												}
											}
										l139:
											add(ruleInc, position138)
										}
										goto l87
									l137:
										position, tokenIndex = position87, tokenIndex87
										{
											position167 := position
											{
												position168, tokenIndex168 := position, tokenIndex
												{
													position170 := position
													{
														position171, tokenIndex171 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l172
														}
														position++
														goto l171
													l172:
														position, tokenIndex = position171, tokenIndex171
														if buffer[position] != rune('D') {
															goto l169
														}
														position++
													}
												l171:
													{
														position173, tokenIndex173 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l174
														}
														position++
														goto l173
													l174:
														position, tokenIndex = position173, tokenIndex173
														if buffer[position] != rune('E') {
															goto l169
														}
														position++
													}
												l173:
													{
														position175, tokenIndex175 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l176
														}
														position++
														goto l175
													l176:
														position, tokenIndex = position175, tokenIndex175
														if buffer[position] != rune('C') {
															goto l169
														}
														position++
													}
												l175:
													if !_rules[rulews]() {
														goto l169
													}
													if !_rules[ruleILoc8]() {
														goto l169
													}
													{
														add(ruleAction14, position)
													}
													add(ruleDec16Indexed8, position170)
												}
												goto l168
											l169:
												position, tokenIndex = position168, tokenIndex168
												{
													position179 := position
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
															goto l178
														}
														position++
													}
												l180:
													{
														position182, tokenIndex182 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l183
														}
														position++
														goto l182
													l183:
														position, tokenIndex = position182, tokenIndex182
														if buffer[position] != rune('E') {
															goto l178
														}
														position++
													}
												l182:
													{
														position184, tokenIndex184 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l185
														}
														position++
														goto l184
													l185:
														position, tokenIndex = position184, tokenIndex184
														if buffer[position] != rune('C') {
															goto l178
														}
														position++
													}
												l184:
													if !_rules[rulews]() {
														goto l178
													}
													if !_rules[ruleLoc16]() {
														goto l178
													}
													{
														add(ruleAction16, position)
													}
													add(ruleDec16, position179)
												}
												goto l168
											l178:
												position, tokenIndex = position168, tokenIndex168
												{
													position187 := position
													{
														position188, tokenIndex188 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l189
														}
														position++
														goto l188
													l189:
														position, tokenIndex = position188, tokenIndex188
														if buffer[position] != rune('D') {
															goto l166
														}
														position++
													}
												l188:
													{
														position190, tokenIndex190 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l191
														}
														position++
														goto l190
													l191:
														position, tokenIndex = position190, tokenIndex190
														if buffer[position] != rune('E') {
															goto l166
														}
														position++
													}
												l190:
													{
														position192, tokenIndex192 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l193
														}
														position++
														goto l192
													l193:
														position, tokenIndex = position192, tokenIndex192
														if buffer[position] != rune('C') {
															goto l166
														}
														position++
													}
												l192:
													if !_rules[rulews]() {
														goto l166
													}
													if !_rules[ruleLoc8]() {
														goto l166
													}
													{
														add(ruleAction15, position)
													}
													add(ruleDec8, position187)
												}
											}
										l168:
											add(ruleDec, position167)
										}
										goto l87
									l166:
										position, tokenIndex = position87, tokenIndex87
										{
											position196 := position
											{
												position197, tokenIndex197 := position, tokenIndex
												{
													position199 := position
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
															goto l198
														}
														position++
													}
												l200:
													{
														position202, tokenIndex202 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l203
														}
														position++
														goto l202
													l203:
														position, tokenIndex = position202, tokenIndex202
														if buffer[position] != rune('D') {
															goto l198
														}
														position++
													}
												l202:
													{
														position204, tokenIndex204 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l205
														}
														position++
														goto l204
													l205:
														position, tokenIndex = position204, tokenIndex204
														if buffer[position] != rune('D') {
															goto l198
														}
														position++
													}
												l204:
													if !_rules[rulews]() {
														goto l198
													}
													if !_rules[ruleDst16]() {
														goto l198
													}
													if !_rules[rulesep]() {
														goto l198
													}
													if !_rules[ruleSrc16]() {
														goto l198
													}
													{
														add(ruleAction17, position)
													}
													add(ruleAdd16, position199)
												}
												goto l197
											l198:
												position, tokenIndex = position197, tokenIndex197
												{
													position208 := position
													{
														position209, tokenIndex209 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l210
														}
														position++
														goto l209
													l210:
														position, tokenIndex = position209, tokenIndex209
														if buffer[position] != rune('A') {
															goto l207
														}
														position++
													}
												l209:
													{
														position211, tokenIndex211 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l212
														}
														position++
														goto l211
													l212:
														position, tokenIndex = position211, tokenIndex211
														if buffer[position] != rune('D') {
															goto l207
														}
														position++
													}
												l211:
													{
														position213, tokenIndex213 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l214
														}
														position++
														goto l213
													l214:
														position, tokenIndex = position213, tokenIndex213
														if buffer[position] != rune('C') {
															goto l207
														}
														position++
													}
												l213:
													if !_rules[rulews]() {
														goto l207
													}
													if !_rules[ruleDst16]() {
														goto l207
													}
													if !_rules[rulesep]() {
														goto l207
													}
													if !_rules[ruleSrc16]() {
														goto l207
													}
													{
														add(ruleAction18, position)
													}
													add(ruleAdc16, position208)
												}
												goto l197
											l207:
												position, tokenIndex = position197, tokenIndex197
												{
													position216 := position
													{
														position217, tokenIndex217 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l218
														}
														position++
														goto l217
													l218:
														position, tokenIndex = position217, tokenIndex217
														if buffer[position] != rune('S') {
															goto l195
														}
														position++
													}
												l217:
													{
														position219, tokenIndex219 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l220
														}
														position++
														goto l219
													l220:
														position, tokenIndex = position219, tokenIndex219
														if buffer[position] != rune('B') {
															goto l195
														}
														position++
													}
												l219:
													{
														position221, tokenIndex221 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l222
														}
														position++
														goto l221
													l222:
														position, tokenIndex = position221, tokenIndex221
														if buffer[position] != rune('C') {
															goto l195
														}
														position++
													}
												l221:
													if !_rules[rulews]() {
														goto l195
													}
													if !_rules[ruleDst16]() {
														goto l195
													}
													if !_rules[rulesep]() {
														goto l195
													}
													if !_rules[ruleSrc16]() {
														goto l195
													}
													{
														add(ruleAction19, position)
													}
													add(ruleSbc16, position216)
												}
											}
										l197:
											add(ruleAlu16, position196)
										}
										goto l87
									l195:
										position, tokenIndex = position87, tokenIndex87
										{
											position225 := position
											{
												position226, tokenIndex226 := position, tokenIndex
												{
													position228 := position
													{
														position229, tokenIndex229 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l230
														}
														position++
														goto l229
													l230:
														position, tokenIndex = position229, tokenIndex229
														if buffer[position] != rune('A') {
															goto l227
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
															goto l227
														}
														position++
													}
												l231:
													{
														position233, tokenIndex233 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l234
														}
														position++
														goto l233
													l234:
														position, tokenIndex = position233, tokenIndex233
														if buffer[position] != rune('D') {
															goto l227
														}
														position++
													}
												l233:
													if !_rules[rulews]() {
														goto l227
													}
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
															goto l227
														}
														position++
													}
												l235:
													if !_rules[rulesep]() {
														goto l227
													}
													if !_rules[ruleSrc8]() {
														goto l227
													}
													{
														add(ruleAction43, position)
													}
													add(ruleAdd, position228)
												}
												goto l226
											l227:
												position, tokenIndex = position226, tokenIndex226
												{
													position239 := position
													{
														position240, tokenIndex240 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l241
														}
														position++
														goto l240
													l241:
														position, tokenIndex = position240, tokenIndex240
														if buffer[position] != rune('A') {
															goto l238
														}
														position++
													}
												l240:
													{
														position242, tokenIndex242 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l243
														}
														position++
														goto l242
													l243:
														position, tokenIndex = position242, tokenIndex242
														if buffer[position] != rune('D') {
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
													{
														position246, tokenIndex246 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l247
														}
														position++
														goto l246
													l247:
														position, tokenIndex = position246, tokenIndex246
														if buffer[position] != rune('A') {
															goto l238
														}
														position++
													}
												l246:
													if !_rules[rulesep]() {
														goto l238
													}
													if !_rules[ruleSrc8]() {
														goto l238
													}
													{
														add(ruleAction44, position)
													}
													add(ruleAdc, position239)
												}
												goto l226
											l238:
												position, tokenIndex = position226, tokenIndex226
												{
													position250 := position
													{
														position251, tokenIndex251 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l252
														}
														position++
														goto l251
													l252:
														position, tokenIndex = position251, tokenIndex251
														if buffer[position] != rune('S') {
															goto l249
														}
														position++
													}
												l251:
													{
														position253, tokenIndex253 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l254
														}
														position++
														goto l253
													l254:
														position, tokenIndex = position253, tokenIndex253
														if buffer[position] != rune('U') {
															goto l249
														}
														position++
													}
												l253:
													{
														position255, tokenIndex255 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l256
														}
														position++
														goto l255
													l256:
														position, tokenIndex = position255, tokenIndex255
														if buffer[position] != rune('B') {
															goto l249
														}
														position++
													}
												l255:
													if !_rules[rulews]() {
														goto l249
													}
													if !_rules[ruleSrc8]() {
														goto l249
													}
													{
														add(ruleAction45, position)
													}
													add(ruleSub, position250)
												}
												goto l226
											l249:
												position, tokenIndex = position226, tokenIndex226
												{
													switch buffer[position] {
													case 'C', 'c':
														{
															position259 := position
															{
																position260, tokenIndex260 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l261
																}
																position++
																goto l260
															l261:
																position, tokenIndex = position260, tokenIndex260
																if buffer[position] != rune('C') {
																	goto l224
																}
																position++
															}
														l260:
															{
																position262, tokenIndex262 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l263
																}
																position++
																goto l262
															l263:
																position, tokenIndex = position262, tokenIndex262
																if buffer[position] != rune('P') {
																	goto l224
																}
																position++
															}
														l262:
															if !_rules[rulews]() {
																goto l224
															}
															if !_rules[ruleSrc8]() {
																goto l224
															}
															{
																add(ruleAction50, position)
															}
															add(ruleCp, position259)
														}
														break
													case 'O', 'o':
														{
															position265 := position
															{
																position266, tokenIndex266 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l267
																}
																position++
																goto l266
															l267:
																position, tokenIndex = position266, tokenIndex266
																if buffer[position] != rune('O') {
																	goto l224
																}
																position++
															}
														l266:
															{
																position268, tokenIndex268 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l269
																}
																position++
																goto l268
															l269:
																position, tokenIndex = position268, tokenIndex268
																if buffer[position] != rune('R') {
																	goto l224
																}
																position++
															}
														l268:
															if !_rules[rulews]() {
																goto l224
															}
															if !_rules[ruleSrc8]() {
																goto l224
															}
															{
																add(ruleAction49, position)
															}
															add(ruleOr, position265)
														}
														break
													case 'X', 'x':
														{
															position271 := position
															{
																position272, tokenIndex272 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l273
																}
																position++
																goto l272
															l273:
																position, tokenIndex = position272, tokenIndex272
																if buffer[position] != rune('X') {
																	goto l224
																}
																position++
															}
														l272:
															{
																position274, tokenIndex274 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l275
																}
																position++
																goto l274
															l275:
																position, tokenIndex = position274, tokenIndex274
																if buffer[position] != rune('O') {
																	goto l224
																}
																position++
															}
														l274:
															{
																position276, tokenIndex276 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l277
																}
																position++
																goto l276
															l277:
																position, tokenIndex = position276, tokenIndex276
																if buffer[position] != rune('R') {
																	goto l224
																}
																position++
															}
														l276:
															if !_rules[rulews]() {
																goto l224
															}
															if !_rules[ruleSrc8]() {
																goto l224
															}
															{
																add(ruleAction48, position)
															}
															add(ruleXor, position271)
														}
														break
													case 'A', 'a':
														{
															position279 := position
															{
																position280, tokenIndex280 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l281
																}
																position++
																goto l280
															l281:
																position, tokenIndex = position280, tokenIndex280
																if buffer[position] != rune('A') {
																	goto l224
																}
																position++
															}
														l280:
															{
																position282, tokenIndex282 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l283
																}
																position++
																goto l282
															l283:
																position, tokenIndex = position282, tokenIndex282
																if buffer[position] != rune('N') {
																	goto l224
																}
																position++
															}
														l282:
															{
																position284, tokenIndex284 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l285
																}
																position++
																goto l284
															l285:
																position, tokenIndex = position284, tokenIndex284
																if buffer[position] != rune('D') {
																	goto l224
																}
																position++
															}
														l284:
															if !_rules[rulews]() {
																goto l224
															}
															if !_rules[ruleSrc8]() {
																goto l224
															}
															{
																add(ruleAction47, position)
															}
															add(ruleAnd, position279)
														}
														break
													default:
														{
															position287 := position
															{
																position288, tokenIndex288 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l289
																}
																position++
																goto l288
															l289:
																position, tokenIndex = position288, tokenIndex288
																if buffer[position] != rune('S') {
																	goto l224
																}
																position++
															}
														l288:
															{
																position290, tokenIndex290 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l291
																}
																position++
																goto l290
															l291:
																position, tokenIndex = position290, tokenIndex290
																if buffer[position] != rune('B') {
																	goto l224
																}
																position++
															}
														l290:
															{
																position292, tokenIndex292 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l293
																}
																position++
																goto l292
															l293:
																position, tokenIndex = position292, tokenIndex292
																if buffer[position] != rune('C') {
																	goto l224
																}
																position++
															}
														l292:
															if !_rules[rulews]() {
																goto l224
															}
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
																	goto l224
																}
																position++
															}
														l294:
															if !_rules[rulesep]() {
																goto l224
															}
															if !_rules[ruleSrc8]() {
																goto l224
															}
															{
																add(ruleAction46, position)
															}
															add(ruleSbc, position287)
														}
														break
													}
												}

											}
										l226:
											add(ruleAlu, position225)
										}
										goto l87
									l224:
										position, tokenIndex = position87, tokenIndex87
										{
											position298 := position
											{
												position299, tokenIndex299 := position, tokenIndex
												{
													position301 := position
													{
														position302, tokenIndex302 := position, tokenIndex
														{
															position304 := position
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
																	goto l303
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
																	goto l303
																}
																position++
															}
														l307:
															{
																position309, tokenIndex309 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l310
																}
																position++
																goto l309
															l310:
																position, tokenIndex = position309, tokenIndex309
																if buffer[position] != rune('C') {
																	goto l303
																}
																position++
															}
														l309:
															if !_rules[rulews]() {
																goto l303
															}
															if !_rules[ruleLoc8]() {
																goto l303
															}
															{
																position311, tokenIndex311 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l311
																}
																if !_rules[ruleCopy8]() {
																	goto l311
																}
																goto l312
															l311:
																position, tokenIndex = position311, tokenIndex311
															}
														l312:
															{
																add(ruleAction51, position)
															}
															add(ruleRlc, position304)
														}
														goto l302
													l303:
														position, tokenIndex = position302, tokenIndex302
														{
															position315 := position
															{
																position316, tokenIndex316 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l317
																}
																position++
																goto l316
															l317:
																position, tokenIndex = position316, tokenIndex316
																if buffer[position] != rune('R') {
																	goto l314
																}
																position++
															}
														l316:
															{
																position318, tokenIndex318 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l319
																}
																position++
																goto l318
															l319:
																position, tokenIndex = position318, tokenIndex318
																if buffer[position] != rune('R') {
																	goto l314
																}
																position++
															}
														l318:
															{
																position320, tokenIndex320 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l321
																}
																position++
																goto l320
															l321:
																position, tokenIndex = position320, tokenIndex320
																if buffer[position] != rune('C') {
																	goto l314
																}
																position++
															}
														l320:
															if !_rules[rulews]() {
																goto l314
															}
															if !_rules[ruleLoc8]() {
																goto l314
															}
															{
																position322, tokenIndex322 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l322
																}
																if !_rules[ruleCopy8]() {
																	goto l322
																}
																goto l323
															l322:
																position, tokenIndex = position322, tokenIndex322
															}
														l323:
															{
																add(ruleAction52, position)
															}
															add(ruleRrc, position315)
														}
														goto l302
													l314:
														position, tokenIndex = position302, tokenIndex302
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
																	goto l325
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
																	goto l325
																}
																position++
															}
														l329:
															if !_rules[rulews]() {
																goto l325
															}
															if !_rules[ruleLoc8]() {
																goto l325
															}
															{
																position331, tokenIndex331 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l331
																}
																if !_rules[ruleCopy8]() {
																	goto l331
																}
																goto l332
															l331:
																position, tokenIndex = position331, tokenIndex331
															}
														l332:
															{
																add(ruleAction53, position)
															}
															add(ruleRl, position326)
														}
														goto l302
													l325:
														position, tokenIndex = position302, tokenIndex302
														{
															position335 := position
															{
																position336, tokenIndex336 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l337
																}
																position++
																goto l336
															l337:
																position, tokenIndex = position336, tokenIndex336
																if buffer[position] != rune('R') {
																	goto l334
																}
																position++
															}
														l336:
															{
																position338, tokenIndex338 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l339
																}
																position++
																goto l338
															l339:
																position, tokenIndex = position338, tokenIndex338
																if buffer[position] != rune('R') {
																	goto l334
																}
																position++
															}
														l338:
															if !_rules[rulews]() {
																goto l334
															}
															if !_rules[ruleLoc8]() {
																goto l334
															}
															{
																position340, tokenIndex340 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l340
																}
																if !_rules[ruleCopy8]() {
																	goto l340
																}
																goto l341
															l340:
																position, tokenIndex = position340, tokenIndex340
															}
														l341:
															{
																add(ruleAction54, position)
															}
															add(ruleRr, position335)
														}
														goto l302
													l334:
														position, tokenIndex = position302, tokenIndex302
														{
															position344 := position
															{
																position345, tokenIndex345 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l346
																}
																position++
																goto l345
															l346:
																position, tokenIndex = position345, tokenIndex345
																if buffer[position] != rune('S') {
																	goto l343
																}
																position++
															}
														l345:
															{
																position347, tokenIndex347 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l348
																}
																position++
																goto l347
															l348:
																position, tokenIndex = position347, tokenIndex347
																if buffer[position] != rune('L') {
																	goto l343
																}
																position++
															}
														l347:
															{
																position349, tokenIndex349 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l350
																}
																position++
																goto l349
															l350:
																position, tokenIndex = position349, tokenIndex349
																if buffer[position] != rune('A') {
																	goto l343
																}
																position++
															}
														l349:
															if !_rules[rulews]() {
																goto l343
															}
															if !_rules[ruleLoc8]() {
																goto l343
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
																add(ruleAction55, position)
															}
															add(ruleSla, position344)
														}
														goto l302
													l343:
														position, tokenIndex = position302, tokenIndex302
														{
															position355 := position
															{
																position356, tokenIndex356 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l357
																}
																position++
																goto l356
															l357:
																position, tokenIndex = position356, tokenIndex356
																if buffer[position] != rune('S') {
																	goto l354
																}
																position++
															}
														l356:
															{
																position358, tokenIndex358 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l359
																}
																position++
																goto l358
															l359:
																position, tokenIndex = position358, tokenIndex358
																if buffer[position] != rune('R') {
																	goto l354
																}
																position++
															}
														l358:
															{
																position360, tokenIndex360 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l361
																}
																position++
																goto l360
															l361:
																position, tokenIndex = position360, tokenIndex360
																if buffer[position] != rune('A') {
																	goto l354
																}
																position++
															}
														l360:
															if !_rules[rulews]() {
																goto l354
															}
															if !_rules[ruleLoc8]() {
																goto l354
															}
															{
																position362, tokenIndex362 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l362
																}
																if !_rules[ruleCopy8]() {
																	goto l362
																}
																goto l363
															l362:
																position, tokenIndex = position362, tokenIndex362
															}
														l363:
															{
																add(ruleAction56, position)
															}
															add(ruleSra, position355)
														}
														goto l302
													l354:
														position, tokenIndex = position302, tokenIndex302
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
																	goto l365
																}
																position++
															}
														l367:
															{
																position369, tokenIndex369 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l370
																}
																position++
																goto l369
															l370:
																position, tokenIndex = position369, tokenIndex369
																if buffer[position] != rune('L') {
																	goto l365
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
																	goto l365
																}
																position++
															}
														l371:
															if !_rules[rulews]() {
																goto l365
															}
															if !_rules[ruleLoc8]() {
																goto l365
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
																add(ruleAction57, position)
															}
															add(ruleSll, position366)
														}
														goto l302
													l365:
														position, tokenIndex = position302, tokenIndex302
														{
															position376 := position
															{
																position377, tokenIndex377 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l378
																}
																position++
																goto l377
															l378:
																position, tokenIndex = position377, tokenIndex377
																if buffer[position] != rune('S') {
																	goto l300
																}
																position++
															}
														l377:
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
																	goto l300
																}
																position++
															}
														l379:
															{
																position381, tokenIndex381 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l382
																}
																position++
																goto l381
															l382:
																position, tokenIndex = position381, tokenIndex381
																if buffer[position] != rune('L') {
																	goto l300
																}
																position++
															}
														l381:
															if !_rules[rulews]() {
																goto l300
															}
															if !_rules[ruleLoc8]() {
																goto l300
															}
															{
																position383, tokenIndex383 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l383
																}
																if !_rules[ruleCopy8]() {
																	goto l383
																}
																goto l384
															l383:
																position, tokenIndex = position383, tokenIndex383
															}
														l384:
															{
																add(ruleAction58, position)
															}
															add(ruleSrl, position376)
														}
													}
												l302:
													add(ruleRot, position301)
												}
												goto l299
											l300:
												position, tokenIndex = position299, tokenIndex299
												{
													switch buffer[position] {
													case 'S', 's':
														{
															position387 := position
															{
																position388, tokenIndex388 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l389
																}
																position++
																goto l388
															l389:
																position, tokenIndex = position388, tokenIndex388
																if buffer[position] != rune('S') {
																	goto l297
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
																	goto l297
																}
																position++
															}
														l390:
															{
																position392, tokenIndex392 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l393
																}
																position++
																goto l392
															l393:
																position, tokenIndex = position392, tokenIndex392
																if buffer[position] != rune('T') {
																	goto l297
																}
																position++
															}
														l392:
															if !_rules[rulews]() {
																goto l297
															}
															if !_rules[ruleoctaldigit]() {
																goto l297
															}
															if !_rules[rulesep]() {
																goto l297
															}
															if !_rules[ruleLoc8]() {
																goto l297
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
																add(ruleAction61, position)
															}
															add(ruleSet, position387)
														}
														break
													case 'R', 'r':
														{
															position397 := position
															{
																position398, tokenIndex398 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l399
																}
																position++
																goto l398
															l399:
																position, tokenIndex = position398, tokenIndex398
																if buffer[position] != rune('R') {
																	goto l297
																}
																position++
															}
														l398:
															{
																position400, tokenIndex400 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l401
																}
																position++
																goto l400
															l401:
																position, tokenIndex = position400, tokenIndex400
																if buffer[position] != rune('E') {
																	goto l297
																}
																position++
															}
														l400:
															{
																position402, tokenIndex402 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l403
																}
																position++
																goto l402
															l403:
																position, tokenIndex = position402, tokenIndex402
																if buffer[position] != rune('S') {
																	goto l297
																}
																position++
															}
														l402:
															if !_rules[rulews]() {
																goto l297
															}
															if !_rules[ruleoctaldigit]() {
																goto l297
															}
															if !_rules[rulesep]() {
																goto l297
															}
															if !_rules[ruleLoc8]() {
																goto l297
															}
															{
																position404, tokenIndex404 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l404
																}
																if !_rules[ruleCopy8]() {
																	goto l404
																}
																goto l405
															l404:
																position, tokenIndex = position404, tokenIndex404
															}
														l405:
															{
																add(ruleAction60, position)
															}
															add(ruleRes, position397)
														}
														break
													default:
														{
															position407 := position
															{
																position408, tokenIndex408 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l409
																}
																position++
																goto l408
															l409:
																position, tokenIndex = position408, tokenIndex408
																if buffer[position] != rune('B') {
																	goto l297
																}
																position++
															}
														l408:
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
																	goto l297
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
																	goto l297
																}
																position++
															}
														l412:
															if !_rules[rulews]() {
																goto l297
															}
															if !_rules[ruleoctaldigit]() {
																goto l297
															}
															if !_rules[rulesep]() {
																goto l297
															}
															if !_rules[ruleLoc8]() {
																goto l297
															}
															{
																add(ruleAction59, position)
															}
															add(ruleBit, position407)
														}
														break
													}
												}

											}
										l299:
											add(ruleBitOp, position298)
										}
										goto l87
									l297:
										position, tokenIndex = position87, tokenIndex87
										{
											position416 := position
											{
												position417, tokenIndex417 := position, tokenIndex
												{
													position419 := position
													{
														position420 := position
														{
															position421, tokenIndex421 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l422
															}
															position++
															goto l421
														l422:
															position, tokenIndex = position421, tokenIndex421
															if buffer[position] != rune('R') {
																goto l418
															}
															position++
														}
													l421:
														{
															position423, tokenIndex423 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l424
															}
															position++
															goto l423
														l424:
															position, tokenIndex = position423, tokenIndex423
															if buffer[position] != rune('E') {
																goto l418
															}
															position++
														}
													l423:
														{
															position425, tokenIndex425 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l426
															}
															position++
															goto l425
														l426:
															position, tokenIndex = position425, tokenIndex425
															if buffer[position] != rune('T') {
																goto l418
															}
															position++
														}
													l425:
														{
															position427, tokenIndex427 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l428
															}
															position++
															goto l427
														l428:
															position, tokenIndex = position427, tokenIndex427
															if buffer[position] != rune('N') {
																goto l418
															}
															position++
														}
													l427:
														add(rulePegText, position420)
													}
													{
														add(ruleAction76, position)
													}
													add(ruleRetn, position419)
												}
												goto l417
											l418:
												position, tokenIndex = position417, tokenIndex417
												{
													position431 := position
													{
														position432 := position
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
																goto l430
															}
															position++
														}
													l433:
														{
															position435, tokenIndex435 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l436
															}
															position++
															goto l435
														l436:
															position, tokenIndex = position435, tokenIndex435
															if buffer[position] != rune('E') {
																goto l430
															}
															position++
														}
													l435:
														{
															position437, tokenIndex437 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l438
															}
															position++
															goto l437
														l438:
															position, tokenIndex = position437, tokenIndex437
															if buffer[position] != rune('T') {
																goto l430
															}
															position++
														}
													l437:
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
																goto l430
															}
															position++
														}
													l439:
														add(rulePegText, position432)
													}
													{
														add(ruleAction77, position)
													}
													add(ruleReti, position431)
												}
												goto l417
											l430:
												position, tokenIndex = position417, tokenIndex417
												{
													position443 := position
													{
														position444 := position
														{
															position445, tokenIndex445 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l446
															}
															position++
															goto l445
														l446:
															position, tokenIndex = position445, tokenIndex445
															if buffer[position] != rune('R') {
																goto l442
															}
															position++
														}
													l445:
														{
															position447, tokenIndex447 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l448
															}
															position++
															goto l447
														l448:
															position, tokenIndex = position447, tokenIndex447
															if buffer[position] != rune('R') {
																goto l442
															}
															position++
														}
													l447:
														{
															position449, tokenIndex449 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l450
															}
															position++
															goto l449
														l450:
															position, tokenIndex = position449, tokenIndex449
															if buffer[position] != rune('D') {
																goto l442
															}
															position++
														}
													l449:
														add(rulePegText, position444)
													}
													{
														add(ruleAction78, position)
													}
													add(ruleRrd, position443)
												}
												goto l417
											l442:
												position, tokenIndex = position417, tokenIndex417
												{
													position453 := position
													{
														position454 := position
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
																goto l452
															}
															position++
														}
													l455:
														{
															position457, tokenIndex457 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l458
															}
															position++
															goto l457
														l458:
															position, tokenIndex = position457, tokenIndex457
															if buffer[position] != rune('M') {
																goto l452
															}
															position++
														}
													l457:
														if buffer[position] != rune(' ') {
															goto l452
														}
														position++
														if buffer[position] != rune('0') {
															goto l452
														}
														position++
														add(rulePegText, position454)
													}
													{
														add(ruleAction80, position)
													}
													add(ruleIm0, position453)
												}
												goto l417
											l452:
												position, tokenIndex = position417, tokenIndex417
												{
													position461 := position
													{
														position462 := position
														{
															position463, tokenIndex463 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l464
															}
															position++
															goto l463
														l464:
															position, tokenIndex = position463, tokenIndex463
															if buffer[position] != rune('I') {
																goto l460
															}
															position++
														}
													l463:
														{
															position465, tokenIndex465 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l466
															}
															position++
															goto l465
														l466:
															position, tokenIndex = position465, tokenIndex465
															if buffer[position] != rune('M') {
																goto l460
															}
															position++
														}
													l465:
														if buffer[position] != rune(' ') {
															goto l460
														}
														position++
														if buffer[position] != rune('1') {
															goto l460
														}
														position++
														add(rulePegText, position462)
													}
													{
														add(ruleAction81, position)
													}
													add(ruleIm1, position461)
												}
												goto l417
											l460:
												position, tokenIndex = position417, tokenIndex417
												{
													position469 := position
													{
														position470 := position
														{
															position471, tokenIndex471 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l472
															}
															position++
															goto l471
														l472:
															position, tokenIndex = position471, tokenIndex471
															if buffer[position] != rune('I') {
																goto l468
															}
															position++
														}
													l471:
														{
															position473, tokenIndex473 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l474
															}
															position++
															goto l473
														l474:
															position, tokenIndex = position473, tokenIndex473
															if buffer[position] != rune('M') {
																goto l468
															}
															position++
														}
													l473:
														if buffer[position] != rune(' ') {
															goto l468
														}
														position++
														if buffer[position] != rune('2') {
															goto l468
														}
														position++
														add(rulePegText, position470)
													}
													{
														add(ruleAction82, position)
													}
													add(ruleIm2, position469)
												}
												goto l417
											l468:
												position, tokenIndex = position417, tokenIndex417
												{
													switch buffer[position] {
													case 'I', 'O', 'i', 'o':
														{
															position477 := position
															{
																position478, tokenIndex478 := position, tokenIndex
																{
																	position480 := position
																	{
																		position481 := position
																		{
																			position482, tokenIndex482 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l483
																			}
																			position++
																			goto l482
																		l483:
																			position, tokenIndex = position482, tokenIndex482
																			if buffer[position] != rune('I') {
																				goto l479
																			}
																			position++
																		}
																	l482:
																		{
																			position484, tokenIndex484 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l485
																			}
																			position++
																			goto l484
																		l485:
																			position, tokenIndex = position484, tokenIndex484
																			if buffer[position] != rune('N') {
																				goto l479
																			}
																			position++
																		}
																	l484:
																		{
																			position486, tokenIndex486 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l487
																			}
																			position++
																			goto l486
																		l487:
																			position, tokenIndex = position486, tokenIndex486
																			if buffer[position] != rune('I') {
																				goto l479
																			}
																			position++
																		}
																	l486:
																		{
																			position488, tokenIndex488 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l489
																			}
																			position++
																			goto l488
																		l489:
																			position, tokenIndex = position488, tokenIndex488
																			if buffer[position] != rune('R') {
																				goto l479
																			}
																			position++
																		}
																	l488:
																		add(rulePegText, position481)
																	}
																	{
																		add(ruleAction93, position)
																	}
																	add(ruleInir, position480)
																}
																goto l478
															l479:
																position, tokenIndex = position478, tokenIndex478
																{
																	position492 := position
																	{
																		position493 := position
																		{
																			position494, tokenIndex494 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l495
																			}
																			position++
																			goto l494
																		l495:
																			position, tokenIndex = position494, tokenIndex494
																			if buffer[position] != rune('I') {
																				goto l491
																			}
																			position++
																		}
																	l494:
																		{
																			position496, tokenIndex496 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l497
																			}
																			position++
																			goto l496
																		l497:
																			position, tokenIndex = position496, tokenIndex496
																			if buffer[position] != rune('N') {
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
																		add(rulePegText, position493)
																	}
																	{
																		add(ruleAction85, position)
																	}
																	add(ruleIni, position492)
																}
																goto l478
															l491:
																position, tokenIndex = position478, tokenIndex478
																{
																	position502 := position
																	{
																		position503 := position
																		{
																			position504, tokenIndex504 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l505
																			}
																			position++
																			goto l504
																		l505:
																			position, tokenIndex = position504, tokenIndex504
																			if buffer[position] != rune('O') {
																				goto l501
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
																				goto l501
																			}
																			position++
																		}
																	l506:
																		{
																			position508, tokenIndex508 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l509
																			}
																			position++
																			goto l508
																		l509:
																			position, tokenIndex = position508, tokenIndex508
																			if buffer[position] != rune('I') {
																				goto l501
																			}
																			position++
																		}
																	l508:
																		{
																			position510, tokenIndex510 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l511
																			}
																			position++
																			goto l510
																		l511:
																			position, tokenIndex = position510, tokenIndex510
																			if buffer[position] != rune('R') {
																				goto l501
																			}
																			position++
																		}
																	l510:
																		add(rulePegText, position503)
																	}
																	{
																		add(ruleAction94, position)
																	}
																	add(ruleOtir, position502)
																}
																goto l478
															l501:
																position, tokenIndex = position478, tokenIndex478
																{
																	position514 := position
																	{
																		position515 := position
																		{
																			position516, tokenIndex516 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l517
																			}
																			position++
																			goto l516
																		l517:
																			position, tokenIndex = position516, tokenIndex516
																			if buffer[position] != rune('O') {
																				goto l513
																			}
																			position++
																		}
																	l516:
																		{
																			position518, tokenIndex518 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l519
																			}
																			position++
																			goto l518
																		l519:
																			position, tokenIndex = position518, tokenIndex518
																			if buffer[position] != rune('U') {
																				goto l513
																			}
																			position++
																		}
																	l518:
																		{
																			position520, tokenIndex520 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l521
																			}
																			position++
																			goto l520
																		l521:
																			position, tokenIndex = position520, tokenIndex520
																			if buffer[position] != rune('T') {
																				goto l513
																			}
																			position++
																		}
																	l520:
																		{
																			position522, tokenIndex522 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l523
																			}
																			position++
																			goto l522
																		l523:
																			position, tokenIndex = position522, tokenIndex522
																			if buffer[position] != rune('I') {
																				goto l513
																			}
																			position++
																		}
																	l522:
																		add(rulePegText, position515)
																	}
																	{
																		add(ruleAction86, position)
																	}
																	add(ruleOuti, position514)
																}
																goto l478
															l513:
																position, tokenIndex = position478, tokenIndex478
																{
																	position526 := position
																	{
																		position527 := position
																		{
																			position528, tokenIndex528 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l529
																			}
																			position++
																			goto l528
																		l529:
																			position, tokenIndex = position528, tokenIndex528
																			if buffer[position] != rune('I') {
																				goto l525
																			}
																			position++
																		}
																	l528:
																		{
																			position530, tokenIndex530 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l531
																			}
																			position++
																			goto l530
																		l531:
																			position, tokenIndex = position530, tokenIndex530
																			if buffer[position] != rune('N') {
																				goto l525
																			}
																			position++
																		}
																	l530:
																		{
																			position532, tokenIndex532 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l533
																			}
																			position++
																			goto l532
																		l533:
																			position, tokenIndex = position532, tokenIndex532
																			if buffer[position] != rune('D') {
																				goto l525
																			}
																			position++
																		}
																	l532:
																		{
																			position534, tokenIndex534 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l535
																			}
																			position++
																			goto l534
																		l535:
																			position, tokenIndex = position534, tokenIndex534
																			if buffer[position] != rune('R') {
																				goto l525
																			}
																			position++
																		}
																	l534:
																		add(rulePegText, position527)
																	}
																	{
																		add(ruleAction97, position)
																	}
																	add(ruleIndr, position526)
																}
																goto l478
															l525:
																position, tokenIndex = position478, tokenIndex478
																{
																	position538 := position
																	{
																		position539 := position
																		{
																			position540, tokenIndex540 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l541
																			}
																			position++
																			goto l540
																		l541:
																			position, tokenIndex = position540, tokenIndex540
																			if buffer[position] != rune('I') {
																				goto l537
																			}
																			position++
																		}
																	l540:
																		{
																			position542, tokenIndex542 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l543
																			}
																			position++
																			goto l542
																		l543:
																			position, tokenIndex = position542, tokenIndex542
																			if buffer[position] != rune('N') {
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
																		add(rulePegText, position539)
																	}
																	{
																		add(ruleAction89, position)
																	}
																	add(ruleInd, position538)
																}
																goto l478
															l537:
																position, tokenIndex = position478, tokenIndex478
																{
																	position548 := position
																	{
																		position549 := position
																		{
																			position550, tokenIndex550 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l551
																			}
																			position++
																			goto l550
																		l551:
																			position, tokenIndex = position550, tokenIndex550
																			if buffer[position] != rune('O') {
																				goto l547
																			}
																			position++
																		}
																	l550:
																		{
																			position552, tokenIndex552 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l553
																			}
																			position++
																			goto l552
																		l553:
																			position, tokenIndex = position552, tokenIndex552
																			if buffer[position] != rune('T') {
																				goto l547
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
																				goto l547
																			}
																			position++
																		}
																	l554:
																		{
																			position556, tokenIndex556 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l557
																			}
																			position++
																			goto l556
																		l557:
																			position, tokenIndex = position556, tokenIndex556
																			if buffer[position] != rune('R') {
																				goto l547
																			}
																			position++
																		}
																	l556:
																		add(rulePegText, position549)
																	}
																	{
																		add(ruleAction98, position)
																	}
																	add(ruleOtdr, position548)
																}
																goto l478
															l547:
																position, tokenIndex = position478, tokenIndex478
																{
																	position559 := position
																	{
																		position560 := position
																		{
																			position561, tokenIndex561 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l562
																			}
																			position++
																			goto l561
																		l562:
																			position, tokenIndex = position561, tokenIndex561
																			if buffer[position] != rune('O') {
																				goto l415
																			}
																			position++
																		}
																	l561:
																		{
																			position563, tokenIndex563 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l564
																			}
																			position++
																			goto l563
																		l564:
																			position, tokenIndex = position563, tokenIndex563
																			if buffer[position] != rune('U') {
																				goto l415
																			}
																			position++
																		}
																	l563:
																		{
																			position565, tokenIndex565 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l566
																			}
																			position++
																			goto l565
																		l566:
																			position, tokenIndex = position565, tokenIndex565
																			if buffer[position] != rune('T') {
																				goto l415
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
																				goto l415
																			}
																			position++
																		}
																	l567:
																		add(rulePegText, position560)
																	}
																	{
																		add(ruleAction90, position)
																	}
																	add(ruleOutd, position559)
																}
															}
														l478:
															add(ruleBlitIO, position477)
														}
														break
													case 'R', 'r':
														{
															position570 := position
															{
																position571 := position
																{
																	position572, tokenIndex572 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l573
																	}
																	position++
																	goto l572
																l573:
																	position, tokenIndex = position572, tokenIndex572
																	if buffer[position] != rune('R') {
																		goto l415
																	}
																	position++
																}
															l572:
																{
																	position574, tokenIndex574 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l575
																	}
																	position++
																	goto l574
																l575:
																	position, tokenIndex = position574, tokenIndex574
																	if buffer[position] != rune('L') {
																		goto l415
																	}
																	position++
																}
															l574:
																{
																	position576, tokenIndex576 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l577
																	}
																	position++
																	goto l576
																l577:
																	position, tokenIndex = position576, tokenIndex576
																	if buffer[position] != rune('D') {
																		goto l415
																	}
																	position++
																}
															l576:
																add(rulePegText, position571)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleRld, position570)
														}
														break
													case 'N', 'n':
														{
															position579 := position
															{
																position580 := position
																{
																	position581, tokenIndex581 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l582
																	}
																	position++
																	goto l581
																l582:
																	position, tokenIndex = position581, tokenIndex581
																	if buffer[position] != rune('N') {
																		goto l415
																	}
																	position++
																}
															l581:
																{
																	position583, tokenIndex583 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l584
																	}
																	position++
																	goto l583
																l584:
																	position, tokenIndex = position583, tokenIndex583
																	if buffer[position] != rune('E') {
																		goto l415
																	}
																	position++
																}
															l583:
																{
																	position585, tokenIndex585 := position, tokenIndex
																	if buffer[position] != rune('g') {
																		goto l586
																	}
																	position++
																	goto l585
																l586:
																	position, tokenIndex = position585, tokenIndex585
																	if buffer[position] != rune('G') {
																		goto l415
																	}
																	position++
																}
															l585:
																add(rulePegText, position580)
															}
															{
																add(ruleAction75, position)
															}
															add(ruleNeg, position579)
														}
														break
													default:
														{
															position588 := position
															{
																position589, tokenIndex589 := position, tokenIndex
																{
																	position591 := position
																	{
																		position592 := position
																		{
																			position593, tokenIndex593 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l594
																			}
																			position++
																			goto l593
																		l594:
																			position, tokenIndex = position593, tokenIndex593
																			if buffer[position] != rune('L') {
																				goto l590
																			}
																			position++
																		}
																	l593:
																		{
																			position595, tokenIndex595 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l596
																			}
																			position++
																			goto l595
																		l596:
																			position, tokenIndex = position595, tokenIndex595
																			if buffer[position] != rune('D') {
																				goto l590
																			}
																			position++
																		}
																	l595:
																		{
																			position597, tokenIndex597 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l598
																			}
																			position++
																			goto l597
																		l598:
																			position, tokenIndex = position597, tokenIndex597
																			if buffer[position] != rune('I') {
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
																		add(ruleAction91, position)
																	}
																	add(ruleLdir, position591)
																}
																goto l589
															l590:
																position, tokenIndex = position589, tokenIndex589
																{
																	position603 := position
																	{
																		position604 := position
																		{
																			position605, tokenIndex605 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l606
																			}
																			position++
																			goto l605
																		l606:
																			position, tokenIndex = position605, tokenIndex605
																			if buffer[position] != rune('L') {
																				goto l602
																			}
																			position++
																		}
																	l605:
																		{
																			position607, tokenIndex607 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l608
																			}
																			position++
																			goto l607
																		l608:
																			position, tokenIndex = position607, tokenIndex607
																			if buffer[position] != rune('D') {
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
																		add(rulePegText, position604)
																	}
																	{
																		add(ruleAction83, position)
																	}
																	add(ruleLdi, position603)
																}
																goto l589
															l602:
																position, tokenIndex = position589, tokenIndex589
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
																				goto l612
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
																				goto l612
																			}
																			position++
																		}
																	l617:
																		{
																			position619, tokenIndex619 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l620
																			}
																			position++
																			goto l619
																		l620:
																			position, tokenIndex = position619, tokenIndex619
																			if buffer[position] != rune('I') {
																				goto l612
																			}
																			position++
																		}
																	l619:
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
																				goto l612
																			}
																			position++
																		}
																	l621:
																		add(rulePegText, position614)
																	}
																	{
																		add(ruleAction92, position)
																	}
																	add(ruleCpir, position613)
																}
																goto l589
															l612:
																position, tokenIndex = position589, tokenIndex589
																{
																	position625 := position
																	{
																		position626 := position
																		{
																			position627, tokenIndex627 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l628
																			}
																			position++
																			goto l627
																		l628:
																			position, tokenIndex = position627, tokenIndex627
																			if buffer[position] != rune('C') {
																				goto l624
																			}
																			position++
																		}
																	l627:
																		{
																			position629, tokenIndex629 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l630
																			}
																			position++
																			goto l629
																		l630:
																			position, tokenIndex = position629, tokenIndex629
																			if buffer[position] != rune('P') {
																				goto l624
																			}
																			position++
																		}
																	l629:
																		{
																			position631, tokenIndex631 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l632
																			}
																			position++
																			goto l631
																		l632:
																			position, tokenIndex = position631, tokenIndex631
																			if buffer[position] != rune('I') {
																				goto l624
																			}
																			position++
																		}
																	l631:
																		add(rulePegText, position626)
																	}
																	{
																		add(ruleAction84, position)
																	}
																	add(ruleCpi, position625)
																}
																goto l589
															l624:
																position, tokenIndex = position589, tokenIndex589
																{
																	position635 := position
																	{
																		position636 := position
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
																				goto l634
																			}
																			position++
																		}
																	l637:
																		{
																			position639, tokenIndex639 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l640
																			}
																			position++
																			goto l639
																		l640:
																			position, tokenIndex = position639, tokenIndex639
																			if buffer[position] != rune('D') {
																				goto l634
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
																				goto l634
																			}
																			position++
																		}
																	l641:
																		{
																			position643, tokenIndex643 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l644
																			}
																			position++
																			goto l643
																		l644:
																			position, tokenIndex = position643, tokenIndex643
																			if buffer[position] != rune('R') {
																				goto l634
																			}
																			position++
																		}
																	l643:
																		add(rulePegText, position636)
																	}
																	{
																		add(ruleAction95, position)
																	}
																	add(ruleLddr, position635)
																}
																goto l589
															l634:
																position, tokenIndex = position589, tokenIndex589
																{
																	position647 := position
																	{
																		position648 := position
																		{
																			position649, tokenIndex649 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l650
																			}
																			position++
																			goto l649
																		l650:
																			position, tokenIndex = position649, tokenIndex649
																			if buffer[position] != rune('L') {
																				goto l646
																			}
																			position++
																		}
																	l649:
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
																		add(rulePegText, position648)
																	}
																	{
																		add(ruleAction87, position)
																	}
																	add(ruleLdd, position647)
																}
																goto l589
															l646:
																position, tokenIndex = position589, tokenIndex589
																{
																	position657 := position
																	{
																		position658 := position
																		{
																			position659, tokenIndex659 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l660
																			}
																			position++
																			goto l659
																		l660:
																			position, tokenIndex = position659, tokenIndex659
																			if buffer[position] != rune('C') {
																				goto l656
																			}
																			position++
																		}
																	l659:
																		{
																			position661, tokenIndex661 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l662
																			}
																			position++
																			goto l661
																		l662:
																			position, tokenIndex = position661, tokenIndex661
																			if buffer[position] != rune('P') {
																				goto l656
																			}
																			position++
																		}
																	l661:
																		{
																			position663, tokenIndex663 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l664
																			}
																			position++
																			goto l663
																		l664:
																			position, tokenIndex = position663, tokenIndex663
																			if buffer[position] != rune('D') {
																				goto l656
																			}
																			position++
																		}
																	l663:
																		{
																			position665, tokenIndex665 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l666
																			}
																			position++
																			goto l665
																		l666:
																			position, tokenIndex = position665, tokenIndex665
																			if buffer[position] != rune('R') {
																				goto l656
																			}
																			position++
																		}
																	l665:
																		add(rulePegText, position658)
																	}
																	{
																		add(ruleAction96, position)
																	}
																	add(ruleCpdr, position657)
																}
																goto l589
															l656:
																position, tokenIndex = position589, tokenIndex589
																{
																	position668 := position
																	{
																		position669 := position
																		{
																			position670, tokenIndex670 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l671
																			}
																			position++
																			goto l670
																		l671:
																			position, tokenIndex = position670, tokenIndex670
																			if buffer[position] != rune('C') {
																				goto l415
																			}
																			position++
																		}
																	l670:
																		{
																			position672, tokenIndex672 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l673
																			}
																			position++
																			goto l672
																		l673:
																			position, tokenIndex = position672, tokenIndex672
																			if buffer[position] != rune('P') {
																				goto l415
																			}
																			position++
																		}
																	l672:
																		{
																			position674, tokenIndex674 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l675
																			}
																			position++
																			goto l674
																		l675:
																			position, tokenIndex = position674, tokenIndex674
																			if buffer[position] != rune('D') {
																				goto l415
																			}
																			position++
																		}
																	l674:
																		add(rulePegText, position669)
																	}
																	{
																		add(ruleAction88, position)
																	}
																	add(ruleCpd, position668)
																}
															}
														l589:
															add(ruleBlit, position588)
														}
														break
													}
												}

											}
										l417:
											add(ruleEDSimple, position416)
										}
										goto l87
									l415:
										position, tokenIndex = position87, tokenIndex87
										{
											position678 := position
											{
												position679, tokenIndex679 := position, tokenIndex
												{
													position681 := position
													{
														position682 := position
														{
															position683, tokenIndex683 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l684
															}
															position++
															goto l683
														l684:
															position, tokenIndex = position683, tokenIndex683
															if buffer[position] != rune('R') {
																goto l680
															}
															position++
														}
													l683:
														{
															position685, tokenIndex685 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l686
															}
															position++
															goto l685
														l686:
															position, tokenIndex = position685, tokenIndex685
															if buffer[position] != rune('L') {
																goto l680
															}
															position++
														}
													l685:
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
																goto l680
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
																goto l680
															}
															position++
														}
													l689:
														add(rulePegText, position682)
													}
													{
														add(ruleAction64, position)
													}
													add(ruleRlca, position681)
												}
												goto l679
											l680:
												position, tokenIndex = position679, tokenIndex679
												{
													position693 := position
													{
														position694 := position
														{
															position695, tokenIndex695 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l696
															}
															position++
															goto l695
														l696:
															position, tokenIndex = position695, tokenIndex695
															if buffer[position] != rune('R') {
																goto l692
															}
															position++
														}
													l695:
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
																goto l692
															}
															position++
														}
													l697:
														{
															position699, tokenIndex699 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l700
															}
															position++
															goto l699
														l700:
															position, tokenIndex = position699, tokenIndex699
															if buffer[position] != rune('C') {
																goto l692
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
																goto l692
															}
															position++
														}
													l701:
														add(rulePegText, position694)
													}
													{
														add(ruleAction65, position)
													}
													add(ruleRrca, position693)
												}
												goto l679
											l692:
												position, tokenIndex = position679, tokenIndex679
												{
													position705 := position
													{
														position706 := position
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
																goto l704
															}
															position++
														}
													l707:
														{
															position709, tokenIndex709 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l710
															}
															position++
															goto l709
														l710:
															position, tokenIndex = position709, tokenIndex709
															if buffer[position] != rune('L') {
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
														add(ruleAction66, position)
													}
													add(ruleRla, position705)
												}
												goto l679
											l704:
												position, tokenIndex = position679, tokenIndex679
												{
													position715 := position
													{
														position716 := position
														{
															position717, tokenIndex717 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l718
															}
															position++
															goto l717
														l718:
															position, tokenIndex = position717, tokenIndex717
															if buffer[position] != rune('D') {
																goto l714
															}
															position++
														}
													l717:
														{
															position719, tokenIndex719 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l720
															}
															position++
															goto l719
														l720:
															position, tokenIndex = position719, tokenIndex719
															if buffer[position] != rune('A') {
																goto l714
															}
															position++
														}
													l719:
														{
															position721, tokenIndex721 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l722
															}
															position++
															goto l721
														l722:
															position, tokenIndex = position721, tokenIndex721
															if buffer[position] != rune('A') {
																goto l714
															}
															position++
														}
													l721:
														add(rulePegText, position716)
													}
													{
														add(ruleAction68, position)
													}
													add(ruleDaa, position715)
												}
												goto l679
											l714:
												position, tokenIndex = position679, tokenIndex679
												{
													position725 := position
													{
														position726 := position
														{
															position727, tokenIndex727 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l728
															}
															position++
															goto l727
														l728:
															position, tokenIndex = position727, tokenIndex727
															if buffer[position] != rune('C') {
																goto l724
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
																goto l724
															}
															position++
														}
													l729:
														{
															position731, tokenIndex731 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l732
															}
															position++
															goto l731
														l732:
															position, tokenIndex = position731, tokenIndex731
															if buffer[position] != rune('L') {
																goto l724
															}
															position++
														}
													l731:
														add(rulePegText, position726)
													}
													{
														add(ruleAction69, position)
													}
													add(ruleCpl, position725)
												}
												goto l679
											l724:
												position, tokenIndex = position679, tokenIndex679
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
																goto l734
															}
															position++
														}
													l737:
														{
															position739, tokenIndex739 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l740
															}
															position++
															goto l739
														l740:
															position, tokenIndex = position739, tokenIndex739
															if buffer[position] != rune('X') {
																goto l734
															}
															position++
														}
													l739:
														{
															position741, tokenIndex741 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l742
															}
															position++
															goto l741
														l742:
															position, tokenIndex = position741, tokenIndex741
															if buffer[position] != rune('X') {
																goto l734
															}
															position++
														}
													l741:
														add(rulePegText, position736)
													}
													{
														add(ruleAction72, position)
													}
													add(ruleExx, position735)
												}
												goto l679
											l734:
												position, tokenIndex = position679, tokenIndex679
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position745 := position
															{
																position746 := position
																{
																	position747, tokenIndex747 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l748
																	}
																	position++
																	goto l747
																l748:
																	position, tokenIndex = position747, tokenIndex747
																	if buffer[position] != rune('E') {
																		goto l677
																	}
																	position++
																}
															l747:
																{
																	position749, tokenIndex749 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l750
																	}
																	position++
																	goto l749
																l750:
																	position, tokenIndex = position749, tokenIndex749
																	if buffer[position] != rune('I') {
																		goto l677
																	}
																	position++
																}
															l749:
																add(rulePegText, position746)
															}
															{
																add(ruleAction74, position)
															}
															add(ruleEi, position745)
														}
														break
													case 'D', 'd':
														{
															position752 := position
															{
																position753 := position
																{
																	position754, tokenIndex754 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l755
																	}
																	position++
																	goto l754
																l755:
																	position, tokenIndex = position754, tokenIndex754
																	if buffer[position] != rune('D') {
																		goto l677
																	}
																	position++
																}
															l754:
																{
																	position756, tokenIndex756 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l757
																	}
																	position++
																	goto l756
																l757:
																	position, tokenIndex = position756, tokenIndex756
																	if buffer[position] != rune('I') {
																		goto l677
																	}
																	position++
																}
															l756:
																add(rulePegText, position753)
															}
															{
																add(ruleAction73, position)
															}
															add(ruleDi, position752)
														}
														break
													case 'C', 'c':
														{
															position759 := position
															{
																position760 := position
																{
																	position761, tokenIndex761 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l762
																	}
																	position++
																	goto l761
																l762:
																	position, tokenIndex = position761, tokenIndex761
																	if buffer[position] != rune('C') {
																		goto l677
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
																		goto l677
																	}
																	position++
																}
															l763:
																{
																	position765, tokenIndex765 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l766
																	}
																	position++
																	goto l765
																l766:
																	position, tokenIndex = position765, tokenIndex765
																	if buffer[position] != rune('F') {
																		goto l677
																	}
																	position++
																}
															l765:
																add(rulePegText, position760)
															}
															{
																add(ruleAction71, position)
															}
															add(ruleCcf, position759)
														}
														break
													case 'S', 's':
														{
															position768 := position
															{
																position769 := position
																{
																	position770, tokenIndex770 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l771
																	}
																	position++
																	goto l770
																l771:
																	position, tokenIndex = position770, tokenIndex770
																	if buffer[position] != rune('S') {
																		goto l677
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
																		goto l677
																	}
																	position++
																}
															l772:
																{
																	position774, tokenIndex774 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l775
																	}
																	position++
																	goto l774
																l775:
																	position, tokenIndex = position774, tokenIndex774
																	if buffer[position] != rune('F') {
																		goto l677
																	}
																	position++
																}
															l774:
																add(rulePegText, position769)
															}
															{
																add(ruleAction70, position)
															}
															add(ruleScf, position768)
														}
														break
													case 'R', 'r':
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
																		goto l677
																	}
																	position++
																}
															l779:
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
																		goto l677
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
																		goto l677
																	}
																	position++
																}
															l783:
																add(rulePegText, position778)
															}
															{
																add(ruleAction67, position)
															}
															add(ruleRra, position777)
														}
														break
													case 'H', 'h':
														{
															position786 := position
															{
																position787 := position
																{
																	position788, tokenIndex788 := position, tokenIndex
																	if buffer[position] != rune('h') {
																		goto l789
																	}
																	position++
																	goto l788
																l789:
																	position, tokenIndex = position788, tokenIndex788
																	if buffer[position] != rune('H') {
																		goto l677
																	}
																	position++
																}
															l788:
																{
																	position790, tokenIndex790 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l791
																	}
																	position++
																	goto l790
																l791:
																	position, tokenIndex = position790, tokenIndex790
																	if buffer[position] != rune('A') {
																		goto l677
																	}
																	position++
																}
															l790:
																{
																	position792, tokenIndex792 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l793
																	}
																	position++
																	goto l792
																l793:
																	position, tokenIndex = position792, tokenIndex792
																	if buffer[position] != rune('L') {
																		goto l677
																	}
																	position++
																}
															l792:
																{
																	position794, tokenIndex794 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l795
																	}
																	position++
																	goto l794
																l795:
																	position, tokenIndex = position794, tokenIndex794
																	if buffer[position] != rune('T') {
																		goto l677
																	}
																	position++
																}
															l794:
																add(rulePegText, position787)
															}
															{
																add(ruleAction63, position)
															}
															add(ruleHalt, position786)
														}
														break
													default:
														{
															position797 := position
															{
																position798 := position
																{
																	position799, tokenIndex799 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l800
																	}
																	position++
																	goto l799
																l800:
																	position, tokenIndex = position799, tokenIndex799
																	if buffer[position] != rune('N') {
																		goto l677
																	}
																	position++
																}
															l799:
																{
																	position801, tokenIndex801 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l802
																	}
																	position++
																	goto l801
																l802:
																	position, tokenIndex = position801, tokenIndex801
																	if buffer[position] != rune('O') {
																		goto l677
																	}
																	position++
																}
															l801:
																{
																	position803, tokenIndex803 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l804
																	}
																	position++
																	goto l803
																l804:
																	position, tokenIndex = position803, tokenIndex803
																	if buffer[position] != rune('P') {
																		goto l677
																	}
																	position++
																}
															l803:
																add(rulePegText, position798)
															}
															{
																add(ruleAction62, position)
															}
															add(ruleNop, position797)
														}
														break
													}
												}

											}
										l679:
											add(ruleSimple, position678)
										}
										goto l87
									l677:
										position, tokenIndex = position87, tokenIndex87
										{
											position807 := position
											{
												position808, tokenIndex808 := position, tokenIndex
												{
													position810 := position
													{
														position811, tokenIndex811 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l812
														}
														position++
														goto l811
													l812:
														position, tokenIndex = position811, tokenIndex811
														if buffer[position] != rune('R') {
															goto l809
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
															goto l809
														}
														position++
													}
												l813:
													{
														position815, tokenIndex815 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l816
														}
														position++
														goto l815
													l816:
														position, tokenIndex = position815, tokenIndex815
														if buffer[position] != rune('T') {
															goto l809
														}
														position++
													}
												l815:
													if !_rules[rulews]() {
														goto l809
													}
													if !_rules[rulen]() {
														goto l809
													}
													{
														add(ruleAction99, position)
													}
													add(ruleRst, position810)
												}
												goto l808
											l809:
												position, tokenIndex = position808, tokenIndex808
												{
													position819 := position
													{
														position820, tokenIndex820 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l821
														}
														position++
														goto l820
													l821:
														position, tokenIndex = position820, tokenIndex820
														if buffer[position] != rune('J') {
															goto l818
														}
														position++
													}
												l820:
													{
														position822, tokenIndex822 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l823
														}
														position++
														goto l822
													l823:
														position, tokenIndex = position822, tokenIndex822
														if buffer[position] != rune('P') {
															goto l818
														}
														position++
													}
												l822:
													if !_rules[rulews]() {
														goto l818
													}
													{
														position824, tokenIndex824 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l824
														}
														if !_rules[rulesep]() {
															goto l824
														}
														goto l825
													l824:
														position, tokenIndex = position824, tokenIndex824
													}
												l825:
													if !_rules[ruleSrc16]() {
														goto l818
													}
													{
														add(ruleAction102, position)
													}
													add(ruleJp, position819)
												}
												goto l808
											l818:
												position, tokenIndex = position808, tokenIndex808
												{
													switch buffer[position] {
													case 'D', 'd':
														{
															position828 := position
															{
																position829, tokenIndex829 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l830
																}
																position++
																goto l829
															l830:
																position, tokenIndex = position829, tokenIndex829
																if buffer[position] != rune('D') {
																	goto l806
																}
																position++
															}
														l829:
															{
																position831, tokenIndex831 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l832
																}
																position++
																goto l831
															l832:
																position, tokenIndex = position831, tokenIndex831
																if buffer[position] != rune('J') {
																	goto l806
																}
																position++
															}
														l831:
															{
																position833, tokenIndex833 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l834
																}
																position++
																goto l833
															l834:
																position, tokenIndex = position833, tokenIndex833
																if buffer[position] != rune('N') {
																	goto l806
																}
																position++
															}
														l833:
															{
																position835, tokenIndex835 := position, tokenIndex
																if buffer[position] != rune('z') {
																	goto l836
																}
																position++
																goto l835
															l836:
																position, tokenIndex = position835, tokenIndex835
																if buffer[position] != rune('Z') {
																	goto l806
																}
																position++
															}
														l835:
															if !_rules[rulews]() {
																goto l806
															}
															if !_rules[ruledisp]() {
																goto l806
															}
															{
																add(ruleAction104, position)
															}
															add(ruleDjnz, position828)
														}
														break
													case 'J', 'j':
														{
															position838 := position
															{
																position839, tokenIndex839 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l840
																}
																position++
																goto l839
															l840:
																position, tokenIndex = position839, tokenIndex839
																if buffer[position] != rune('J') {
																	goto l806
																}
																position++
															}
														l839:
															{
																position841, tokenIndex841 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l842
																}
																position++
																goto l841
															l842:
																position, tokenIndex = position841, tokenIndex841
																if buffer[position] != rune('R') {
																	goto l806
																}
																position++
															}
														l841:
															if !_rules[rulews]() {
																goto l806
															}
															{
																position843, tokenIndex843 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l843
																}
																if !_rules[rulesep]() {
																	goto l843
																}
																goto l844
															l843:
																position, tokenIndex = position843, tokenIndex843
															}
														l844:
															if !_rules[ruledisp]() {
																goto l806
															}
															{
																add(ruleAction103, position)
															}
															add(ruleJr, position838)
														}
														break
													case 'R', 'r':
														{
															position846 := position
															{
																position847, tokenIndex847 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l848
																}
																position++
																goto l847
															l848:
																position, tokenIndex = position847, tokenIndex847
																if buffer[position] != rune('R') {
																	goto l806
																}
																position++
															}
														l847:
															{
																position849, tokenIndex849 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l850
																}
																position++
																goto l849
															l850:
																position, tokenIndex = position849, tokenIndex849
																if buffer[position] != rune('E') {
																	goto l806
																}
																position++
															}
														l849:
															{
																position851, tokenIndex851 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l852
																}
																position++
																goto l851
															l852:
																position, tokenIndex = position851, tokenIndex851
																if buffer[position] != rune('T') {
																	goto l806
																}
																position++
															}
														l851:
															{
																position853, tokenIndex853 := position, tokenIndex
																if !_rules[rulews]() {
																	goto l853
																}
																if !_rules[rulecc]() {
																	goto l853
																}
																goto l854
															l853:
																position, tokenIndex = position853, tokenIndex853
															}
														l854:
															{
																add(ruleAction101, position)
															}
															add(ruleRet, position846)
														}
														break
													default:
														{
															position856 := position
															{
																position857, tokenIndex857 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l858
																}
																position++
																goto l857
															l858:
																position, tokenIndex = position857, tokenIndex857
																if buffer[position] != rune('C') {
																	goto l806
																}
																position++
															}
														l857:
															{
																position859, tokenIndex859 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l860
																}
																position++
																goto l859
															l860:
																position, tokenIndex = position859, tokenIndex859
																if buffer[position] != rune('A') {
																	goto l806
																}
																position++
															}
														l859:
															{
																position861, tokenIndex861 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l862
																}
																position++
																goto l861
															l862:
																position, tokenIndex = position861, tokenIndex861
																if buffer[position] != rune('L') {
																	goto l806
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
																	goto l806
																}
																position++
															}
														l863:
															if !_rules[rulews]() {
																goto l806
															}
															{
																position865, tokenIndex865 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l865
																}
																if !_rules[rulesep]() {
																	goto l865
																}
																goto l866
															l865:
																position, tokenIndex = position865, tokenIndex865
															}
														l866:
															if !_rules[ruleSrc16]() {
																goto l806
															}
															{
																add(ruleAction100, position)
															}
															add(ruleCall, position856)
														}
														break
													}
												}

											}
										l808:
											add(ruleJump, position807)
										}
										goto l87
									l806:
										position, tokenIndex = position87, tokenIndex87
										{
											position868 := position
											{
												position869, tokenIndex869 := position, tokenIndex
												{
													position871 := position
													{
														position872, tokenIndex872 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l873
														}
														position++
														goto l872
													l873:
														position, tokenIndex = position872, tokenIndex872
														if buffer[position] != rune('I') {
															goto l870
														}
														position++
													}
												l872:
													{
														position874, tokenIndex874 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l875
														}
														position++
														goto l874
													l875:
														position, tokenIndex = position874, tokenIndex874
														if buffer[position] != rune('N') {
															goto l870
														}
														position++
													}
												l874:
													if !_rules[rulews]() {
														goto l870
													}
													if !_rules[ruleReg8]() {
														goto l870
													}
													if !_rules[rulesep]() {
														goto l870
													}
													if !_rules[rulePort]() {
														goto l870
													}
													{
														add(ruleAction105, position)
													}
													add(ruleIN, position871)
												}
												goto l869
											l870:
												position, tokenIndex = position869, tokenIndex869
												{
													position877 := position
													{
														position878, tokenIndex878 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l879
														}
														position++
														goto l878
													l879:
														position, tokenIndex = position878, tokenIndex878
														if buffer[position] != rune('O') {
															goto l13
														}
														position++
													}
												l878:
													{
														position880, tokenIndex880 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l881
														}
														position++
														goto l880
													l881:
														position, tokenIndex = position880, tokenIndex880
														if buffer[position] != rune('U') {
															goto l13
														}
														position++
													}
												l880:
													{
														position882, tokenIndex882 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l883
														}
														position++
														goto l882
													l883:
														position, tokenIndex = position882, tokenIndex882
														if buffer[position] != rune('T') {
															goto l13
														}
														position++
													}
												l882:
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
													add(ruleOUT, position877)
												}
											}
										l869:
											add(ruleIO, position868)
										}
									}
								l87:
									add(ruleInstruction, position86)
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
							position889 := position
							{
								position890, tokenIndex890 := position, tokenIndex
								if buffer[position] != rune(';') {
									goto l891
								}
								position++
								goto l890
							l891:
								position, tokenIndex = position890, tokenIndex890
								if buffer[position] != rune('#') {
									goto l887
								}
								position++
							}
						l890:
						l892:
							{
								position893, tokenIndex893 := position, tokenIndex
								{
									position894, tokenIndex894 := position, tokenIndex
									if buffer[position] != rune('\n') {
										goto l894
									}
									position++
									goto l893
								l894:
									position, tokenIndex = position894, tokenIndex894
								}
								if !matchDot() {
									goto l893
								}
								goto l892
							l893:
								position, tokenIndex = position893, tokenIndex893
							}
							add(ruleComment, position889)
						}
						goto l888
					l887:
						position, tokenIndex = position887, tokenIndex887
					}
				l888:
					{
						position895, tokenIndex895 := position, tokenIndex
						if !_rules[rulews]() {
							goto l895
						}
						goto l896
					l895:
						position, tokenIndex = position895, tokenIndex895
					}
				l896:
					{
						position897, tokenIndex897 := position, tokenIndex
						{
							position899, tokenIndex899 := position, tokenIndex
							if buffer[position] != rune('\r') {
								goto l899
							}
							position++
							goto l900
						l899:
							position, tokenIndex = position899, tokenIndex899
						}
					l900:
						if buffer[position] != rune('\n') {
							goto l898
						}
						position++
						goto l897
					l898:
						position, tokenIndex = position897, tokenIndex897
						if buffer[position] != rune(':') {
							goto l0
						}
						position++
					}
				l897:
					{
						add(ruleAction0, position)
					}
					add(ruleLine, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position902 := position
					l903:
						{
							position904, tokenIndex904 := position, tokenIndex
							if !_rules[rulews]() {
								goto l904
							}
							goto l903
						l904:
							position, tokenIndex = position904, tokenIndex904
						}
						{
							position905, tokenIndex905 := position, tokenIndex
							{
								position907 := position
								if !_rules[ruleLabelText]() {
									goto l905
								}
								if buffer[position] != rune(':') {
									goto l905
								}
								position++
								if !_rules[rulews]() {
									goto l905
								}
								{
									add(ruleAction5, position)
								}
								add(ruleLabelDefn, position907)
							}
							goto l906
						l905:
							position, tokenIndex = position905, tokenIndex905
						}
					l906:
					l909:
						{
							position910, tokenIndex910 := position, tokenIndex
							if !_rules[rulews]() {
								goto l910
							}
							goto l909
						l910:
							position, tokenIndex = position910, tokenIndex910
						}
						{
							position911, tokenIndex911 := position, tokenIndex
							{
								position913 := position
								{
									position914, tokenIndex914 := position, tokenIndex
									{
										position916 := position
										{
											position917, tokenIndex917 := position, tokenIndex
											{
												position919 := position
												{
													position920, tokenIndex920 := position, tokenIndex
													{
														position922, tokenIndex922 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l923
														}
														position++
														goto l922
													l923:
														position, tokenIndex = position922, tokenIndex922
														if buffer[position] != rune('D') {
															goto l921
														}
														position++
													}
												l922:
													{
														position924, tokenIndex924 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l925
														}
														position++
														goto l924
													l925:
														position, tokenIndex = position924, tokenIndex924
														if buffer[position] != rune('E') {
															goto l921
														}
														position++
													}
												l924:
													{
														position926, tokenIndex926 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l927
														}
														position++
														goto l926
													l927:
														position, tokenIndex = position926, tokenIndex926
														if buffer[position] != rune('F') {
															goto l921
														}
														position++
													}
												l926:
													{
														position928, tokenIndex928 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l929
														}
														position++
														goto l928
													l929:
														position, tokenIndex = position928, tokenIndex928
														if buffer[position] != rune('B') {
															goto l921
														}
														position++
													}
												l928:
													goto l920
												l921:
													position, tokenIndex = position920, tokenIndex920
													{
														position930, tokenIndex930 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l931
														}
														position++
														goto l930
													l931:
														position, tokenIndex = position930, tokenIndex930
														if buffer[position] != rune('D') {
															goto l918
														}
														position++
													}
												l930:
													{
														position932, tokenIndex932 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l933
														}
														position++
														goto l932
													l933:
														position, tokenIndex = position932, tokenIndex932
														if buffer[position] != rune('B') {
															goto l918
														}
														position++
													}
												l932:
												}
											l920:
												if !_rules[rulews]() {
													goto l918
												}
												if !_rules[rulen]() {
													goto l918
												}
												{
													add(ruleAction2, position)
												}
												add(ruleDefb, position919)
											}
											goto l917
										l918:
											position, tokenIndex = position917, tokenIndex917
											{
												position936 := position
												{
													position937, tokenIndex937 := position, tokenIndex
													{
														position939, tokenIndex939 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l940
														}
														position++
														goto l939
													l940:
														position, tokenIndex = position939, tokenIndex939
														if buffer[position] != rune('D') {
															goto l938
														}
														position++
													}
												l939:
													{
														position941, tokenIndex941 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l942
														}
														position++
														goto l941
													l942:
														position, tokenIndex = position941, tokenIndex941
														if buffer[position] != rune('E') {
															goto l938
														}
														position++
													}
												l941:
													{
														position943, tokenIndex943 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l944
														}
														position++
														goto l943
													l944:
														position, tokenIndex = position943, tokenIndex943
														if buffer[position] != rune('F') {
															goto l938
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
															goto l938
														}
														position++
													}
												l945:
													goto l937
												l938:
													position, tokenIndex = position937, tokenIndex937
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
															goto l935
														}
														position++
													}
												l947:
													{
														position949, tokenIndex949 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l950
														}
														position++
														goto l949
													l950:
														position, tokenIndex = position949, tokenIndex949
														if buffer[position] != rune('S') {
															goto l935
														}
														position++
													}
												l949:
												}
											l937:
												if !_rules[rulews]() {
													goto l935
												}
												if !_rules[rulen]() {
													goto l935
												}
												{
													add(ruleAction4, position)
												}
												add(ruleDefs, position936)
											}
											goto l917
										l935:
											position, tokenIndex = position917, tokenIndex917
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position953 := position
														{
															position954, tokenIndex954 := position, tokenIndex
															{
																position956, tokenIndex956 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l957
																}
																position++
																goto l956
															l957:
																position, tokenIndex = position956, tokenIndex956
																if buffer[position] != rune('D') {
																	goto l955
																}
																position++
															}
														l956:
															{
																position958, tokenIndex958 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l959
																}
																position++
																goto l958
															l959:
																position, tokenIndex = position958, tokenIndex958
																if buffer[position] != rune('E') {
																	goto l955
																}
																position++
															}
														l958:
															{
																position960, tokenIndex960 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l961
																}
																position++
																goto l960
															l961:
																position, tokenIndex = position960, tokenIndex960
																if buffer[position] != rune('F') {
																	goto l955
																}
																position++
															}
														l960:
															{
																position962, tokenIndex962 := position, tokenIndex
																if buffer[position] != rune('w') {
																	goto l963
																}
																position++
																goto l962
															l963:
																position, tokenIndex = position962, tokenIndex962
																if buffer[position] != rune('W') {
																	goto l955
																}
																position++
															}
														l962:
															goto l954
														l955:
															position, tokenIndex = position954, tokenIndex954
															{
																position964, tokenIndex964 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l965
																}
																position++
																goto l964
															l965:
																position, tokenIndex = position964, tokenIndex964
																if buffer[position] != rune('D') {
																	goto l915
																}
																position++
															}
														l964:
															{
																position966, tokenIndex966 := position, tokenIndex
																if buffer[position] != rune('w') {
																	goto l967
																}
																position++
																goto l966
															l967:
																position, tokenIndex = position966, tokenIndex966
																if buffer[position] != rune('W') {
																	goto l915
																}
																position++
															}
														l966:
														}
													l954:
														if !_rules[rulews]() {
															goto l915
														}
														if !_rules[rulenn]() {
															goto l915
														}
														{
															add(ruleAction3, position)
														}
														add(ruleDefw, position953)
													}
													break
												case 'O', 'o':
													{
														position969 := position
														{
															position970, tokenIndex970 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l971
															}
															position++
															goto l970
														l971:
															position, tokenIndex = position970, tokenIndex970
															if buffer[position] != rune('O') {
																goto l915
															}
															position++
														}
													l970:
														{
															position972, tokenIndex972 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l973
															}
															position++
															goto l972
														l973:
															position, tokenIndex = position972, tokenIndex972
															if buffer[position] != rune('R') {
																goto l915
															}
															position++
														}
													l972:
														{
															position974, tokenIndex974 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l975
															}
															position++
															goto l974
														l975:
															position, tokenIndex = position974, tokenIndex974
															if buffer[position] != rune('G') {
																goto l915
															}
															position++
														}
													l974:
														if !_rules[rulews]() {
															goto l915
														}
														if !_rules[rulenn]() {
															goto l915
														}
														{
															add(ruleAction1, position)
														}
														add(ruleOrg, position969)
													}
													break
												case 'a':
													{
														position977 := position
														if buffer[position] != rune('a') {
															goto l915
														}
														position++
														if buffer[position] != rune('s') {
															goto l915
														}
														position++
														if buffer[position] != rune('e') {
															goto l915
														}
														position++
														if buffer[position] != rune('g') {
															goto l915
														}
														position++
														add(ruleAseg, position977)
													}
													break
												default:
													{
														position978 := position
														{
															position979, tokenIndex979 := position, tokenIndex
															if buffer[position] != rune('.') {
																goto l979
															}
															position++
															goto l980
														l979:
															position, tokenIndex = position979, tokenIndex979
														}
													l980:
														if buffer[position] != rune('t') {
															goto l915
														}
														position++
														if buffer[position] != rune('i') {
															goto l915
														}
														position++
														if buffer[position] != rune('t') {
															goto l915
														}
														position++
														if buffer[position] != rune('l') {
															goto l915
														}
														position++
														if buffer[position] != rune('e') {
															goto l915
														}
														position++
														if !_rules[rulews]() {
															goto l915
														}
														if buffer[position] != rune('\'') {
															goto l915
														}
														position++
													l981:
														{
															position982, tokenIndex982 := position, tokenIndex
															{
																position983, tokenIndex983 := position, tokenIndex
																if buffer[position] != rune('\'') {
																	goto l983
																}
																position++
																goto l982
															l983:
																position, tokenIndex = position983, tokenIndex983
															}
															if !matchDot() {
																goto l982
															}
															goto l981
														l982:
															position, tokenIndex = position982, tokenIndex982
														}
														if buffer[position] != rune('\'') {
															goto l915
														}
														position++
														add(ruleTitle, position978)
													}
													break
												}
											}

										}
									l917:
										add(ruleDirective, position916)
									}
									goto l914
								l915:
									position, tokenIndex = position914, tokenIndex914
									{
										position984 := position
										{
											position985, tokenIndex985 := position, tokenIndex
											{
												position987 := position
												{
													position988, tokenIndex988 := position, tokenIndex
													{
														position990 := position
														{
															position991, tokenIndex991 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l992
															}
															position++
															goto l991
														l992:
															position, tokenIndex = position991, tokenIndex991
															if buffer[position] != rune('P') {
																goto l989
															}
															position++
														}
													l991:
														{
															position993, tokenIndex993 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l994
															}
															position++
															goto l993
														l994:
															position, tokenIndex = position993, tokenIndex993
															if buffer[position] != rune('U') {
																goto l989
															}
															position++
														}
													l993:
														{
															position995, tokenIndex995 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l996
															}
															position++
															goto l995
														l996:
															position, tokenIndex = position995, tokenIndex995
															if buffer[position] != rune('S') {
																goto l989
															}
															position++
														}
													l995:
														{
															position997, tokenIndex997 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l998
															}
															position++
															goto l997
														l998:
															position, tokenIndex = position997, tokenIndex997
															if buffer[position] != rune('H') {
																goto l989
															}
															position++
														}
													l997:
														if !_rules[rulews]() {
															goto l989
														}
														if !_rules[ruleSrc16]() {
															goto l989
														}
														{
															add(ruleAction8, position)
														}
														add(rulePush, position990)
													}
													goto l988
												l989:
													position, tokenIndex = position988, tokenIndex988
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1001 := position
																{
																	position1002, tokenIndex1002 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1003
																	}
																	position++
																	goto l1002
																l1003:
																	position, tokenIndex = position1002, tokenIndex1002
																	if buffer[position] != rune('E') {
																		goto l986
																	}
																	position++
																}
															l1002:
																{
																	position1004, tokenIndex1004 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1005
																	}
																	position++
																	goto l1004
																l1005:
																	position, tokenIndex = position1004, tokenIndex1004
																	if buffer[position] != rune('X') {
																		goto l986
																	}
																	position++
																}
															l1004:
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
																	add(ruleAction10, position)
																}
																add(ruleEx, position1001)
															}
															break
														case 'P', 'p':
															{
																position1007 := position
																{
																	position1008, tokenIndex1008 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1009
																	}
																	position++
																	goto l1008
																l1009:
																	position, tokenIndex = position1008, tokenIndex1008
																	if buffer[position] != rune('P') {
																		goto l986
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
																		goto l986
																	}
																	position++
																}
															l1010:
																{
																	position1012, tokenIndex1012 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1013
																	}
																	position++
																	goto l1012
																l1013:
																	position, tokenIndex = position1012, tokenIndex1012
																	if buffer[position] != rune('P') {
																		goto l986
																	}
																	position++
																}
															l1012:
																if !_rules[rulews]() {
																	goto l986
																}
																if !_rules[ruleDst16]() {
																	goto l986
																}
																{
																	add(ruleAction9, position)
																}
																add(rulePop, position1007)
															}
															break
														default:
															{
																position1015 := position
																{
																	position1016, tokenIndex1016 := position, tokenIndex
																	{
																		position1018 := position
																		{
																			position1019, tokenIndex1019 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l1020
																			}
																			position++
																			goto l1019
																		l1020:
																			position, tokenIndex = position1019, tokenIndex1019
																			if buffer[position] != rune('L') {
																				goto l1017
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
																				goto l1017
																			}
																			position++
																		}
																	l1021:
																		if !_rules[rulews]() {
																			goto l1017
																		}
																		if !_rules[ruleDst16]() {
																			goto l1017
																		}
																		if !_rules[rulesep]() {
																			goto l1017
																		}
																		if !_rules[ruleSrc16]() {
																			goto l1017
																		}
																		{
																			add(ruleAction7, position)
																		}
																		add(ruleLoad16, position1018)
																	}
																	goto l1016
																l1017:
																	position, tokenIndex = position1016, tokenIndex1016
																	{
																		position1024 := position
																		{
																			position1025, tokenIndex1025 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l1026
																			}
																			position++
																			goto l1025
																		l1026:
																			position, tokenIndex = position1025, tokenIndex1025
																			if buffer[position] != rune('L') {
																				goto l986
																			}
																			position++
																		}
																	l1025:
																		{
																			position1027, tokenIndex1027 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l1028
																			}
																			position++
																			goto l1027
																		l1028:
																			position, tokenIndex = position1027, tokenIndex1027
																			if buffer[position] != rune('D') {
																				goto l986
																			}
																			position++
																		}
																	l1027:
																		if !_rules[rulews]() {
																			goto l986
																		}
																		{
																			position1029 := position
																			{
																				position1030, tokenIndex1030 := position, tokenIndex
																				if !_rules[ruleReg8]() {
																					goto l1031
																				}
																				goto l1030
																			l1031:
																				position, tokenIndex = position1030, tokenIndex1030
																				if !_rules[ruleReg16Contents]() {
																					goto l1032
																				}
																				goto l1030
																			l1032:
																				position, tokenIndex = position1030, tokenIndex1030
																				if !_rules[rulenn_contents]() {
																					goto l986
																				}
																			}
																		l1030:
																			{
																				add(ruleAction20, position)
																			}
																			add(ruleDst8, position1029)
																		}
																		if !_rules[rulesep]() {
																			goto l986
																		}
																		if !_rules[ruleSrc8]() {
																			goto l986
																		}
																		{
																			add(ruleAction6, position)
																		}
																		add(ruleLoad8, position1024)
																	}
																}
															l1016:
																add(ruleLoad, position1015)
															}
															break
														}
													}

												}
											l988:
												add(ruleAssignment, position987)
											}
											goto l985
										l986:
											position, tokenIndex = position985, tokenIndex985
											{
												position1036 := position
												{
													position1037, tokenIndex1037 := position, tokenIndex
													{
														position1039 := position
														{
															position1040, tokenIndex1040 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1041
															}
															position++
															goto l1040
														l1041:
															position, tokenIndex = position1040, tokenIndex1040
															if buffer[position] != rune('I') {
																goto l1038
															}
															position++
														}
													l1040:
														{
															position1042, tokenIndex1042 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1043
															}
															position++
															goto l1042
														l1043:
															position, tokenIndex = position1042, tokenIndex1042
															if buffer[position] != rune('N') {
																goto l1038
															}
															position++
														}
													l1042:
														{
															position1044, tokenIndex1044 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1045
															}
															position++
															goto l1044
														l1045:
															position, tokenIndex = position1044, tokenIndex1044
															if buffer[position] != rune('C') {
																goto l1038
															}
															position++
														}
													l1044:
														if !_rules[rulews]() {
															goto l1038
														}
														if !_rules[ruleILoc8]() {
															goto l1038
														}
														{
															add(ruleAction11, position)
														}
														add(ruleInc16Indexed8, position1039)
													}
													goto l1037
												l1038:
													position, tokenIndex = position1037, tokenIndex1037
													{
														position1048 := position
														{
															position1049, tokenIndex1049 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1050
															}
															position++
															goto l1049
														l1050:
															position, tokenIndex = position1049, tokenIndex1049
															if buffer[position] != rune('I') {
																goto l1047
															}
															position++
														}
													l1049:
														{
															position1051, tokenIndex1051 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1052
															}
															position++
															goto l1051
														l1052:
															position, tokenIndex = position1051, tokenIndex1051
															if buffer[position] != rune('N') {
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
														if !_rules[ruleLoc16]() {
															goto l1047
														}
														{
															add(ruleAction13, position)
														}
														add(ruleInc16, position1048)
													}
													goto l1037
												l1047:
													position, tokenIndex = position1037, tokenIndex1037
													{
														position1056 := position
														{
															position1057, tokenIndex1057 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1058
															}
															position++
															goto l1057
														l1058:
															position, tokenIndex = position1057, tokenIndex1057
															if buffer[position] != rune('I') {
																goto l1035
															}
															position++
														}
													l1057:
														{
															position1059, tokenIndex1059 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1060
															}
															position++
															goto l1059
														l1060:
															position, tokenIndex = position1059, tokenIndex1059
															if buffer[position] != rune('N') {
																goto l1035
															}
															position++
														}
													l1059:
														{
															position1061, tokenIndex1061 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1062
															}
															position++
															goto l1061
														l1062:
															position, tokenIndex = position1061, tokenIndex1061
															if buffer[position] != rune('C') {
																goto l1035
															}
															position++
														}
													l1061:
														if !_rules[rulews]() {
															goto l1035
														}
														if !_rules[ruleLoc8]() {
															goto l1035
														}
														{
															add(ruleAction12, position)
														}
														add(ruleInc8, position1056)
													}
												}
											l1037:
												add(ruleInc, position1036)
											}
											goto l985
										l1035:
											position, tokenIndex = position985, tokenIndex985
											{
												position1065 := position
												{
													position1066, tokenIndex1066 := position, tokenIndex
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
																goto l1067
															}
															position++
														}
													l1069:
														{
															position1071, tokenIndex1071 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1072
															}
															position++
															goto l1071
														l1072:
															position, tokenIndex = position1071, tokenIndex1071
															if buffer[position] != rune('E') {
																goto l1067
															}
															position++
														}
													l1071:
														{
															position1073, tokenIndex1073 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1074
															}
															position++
															goto l1073
														l1074:
															position, tokenIndex = position1073, tokenIndex1073
															if buffer[position] != rune('C') {
																goto l1067
															}
															position++
														}
													l1073:
														if !_rules[rulews]() {
															goto l1067
														}
														if !_rules[ruleILoc8]() {
															goto l1067
														}
														{
															add(ruleAction14, position)
														}
														add(ruleDec16Indexed8, position1068)
													}
													goto l1066
												l1067:
													position, tokenIndex = position1066, tokenIndex1066
													{
														position1077 := position
														{
															position1078, tokenIndex1078 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1079
															}
															position++
															goto l1078
														l1079:
															position, tokenIndex = position1078, tokenIndex1078
															if buffer[position] != rune('D') {
																goto l1076
															}
															position++
														}
													l1078:
														{
															position1080, tokenIndex1080 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1081
															}
															position++
															goto l1080
														l1081:
															position, tokenIndex = position1080, tokenIndex1080
															if buffer[position] != rune('E') {
																goto l1076
															}
															position++
														}
													l1080:
														{
															position1082, tokenIndex1082 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1083
															}
															position++
															goto l1082
														l1083:
															position, tokenIndex = position1082, tokenIndex1082
															if buffer[position] != rune('C') {
																goto l1076
															}
															position++
														}
													l1082:
														if !_rules[rulews]() {
															goto l1076
														}
														if !_rules[ruleLoc16]() {
															goto l1076
														}
														{
															add(ruleAction16, position)
														}
														add(ruleDec16, position1077)
													}
													goto l1066
												l1076:
													position, tokenIndex = position1066, tokenIndex1066
													{
														position1085 := position
														{
															position1086, tokenIndex1086 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1087
															}
															position++
															goto l1086
														l1087:
															position, tokenIndex = position1086, tokenIndex1086
															if buffer[position] != rune('D') {
																goto l1064
															}
															position++
														}
													l1086:
														{
															position1088, tokenIndex1088 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1089
															}
															position++
															goto l1088
														l1089:
															position, tokenIndex = position1088, tokenIndex1088
															if buffer[position] != rune('E') {
																goto l1064
															}
															position++
														}
													l1088:
														{
															position1090, tokenIndex1090 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1091
															}
															position++
															goto l1090
														l1091:
															position, tokenIndex = position1090, tokenIndex1090
															if buffer[position] != rune('C') {
																goto l1064
															}
															position++
														}
													l1090:
														if !_rules[rulews]() {
															goto l1064
														}
														if !_rules[ruleLoc8]() {
															goto l1064
														}
														{
															add(ruleAction15, position)
														}
														add(ruleDec8, position1085)
													}
												}
											l1066:
												add(ruleDec, position1065)
											}
											goto l985
										l1064:
											position, tokenIndex = position985, tokenIndex985
											{
												position1094 := position
												{
													position1095, tokenIndex1095 := position, tokenIndex
													{
														position1097 := position
														{
															position1098, tokenIndex1098 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1099
															}
															position++
															goto l1098
														l1099:
															position, tokenIndex = position1098, tokenIndex1098
															if buffer[position] != rune('A') {
																goto l1096
															}
															position++
														}
													l1098:
														{
															position1100, tokenIndex1100 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1101
															}
															position++
															goto l1100
														l1101:
															position, tokenIndex = position1100, tokenIndex1100
															if buffer[position] != rune('D') {
																goto l1096
															}
															position++
														}
													l1100:
														{
															position1102, tokenIndex1102 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1103
															}
															position++
															goto l1102
														l1103:
															position, tokenIndex = position1102, tokenIndex1102
															if buffer[position] != rune('D') {
																goto l1096
															}
															position++
														}
													l1102:
														if !_rules[rulews]() {
															goto l1096
														}
														if !_rules[ruleDst16]() {
															goto l1096
														}
														if !_rules[rulesep]() {
															goto l1096
														}
														if !_rules[ruleSrc16]() {
															goto l1096
														}
														{
															add(ruleAction17, position)
														}
														add(ruleAdd16, position1097)
													}
													goto l1095
												l1096:
													position, tokenIndex = position1095, tokenIndex1095
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
															if buffer[position] != rune('c') {
																goto l1112
															}
															position++
															goto l1111
														l1112:
															position, tokenIndex = position1111, tokenIndex1111
															if buffer[position] != rune('C') {
																goto l1105
															}
															position++
														}
													l1111:
														if !_rules[rulews]() {
															goto l1105
														}
														if !_rules[ruleDst16]() {
															goto l1105
														}
														if !_rules[rulesep]() {
															goto l1105
														}
														if !_rules[ruleSrc16]() {
															goto l1105
														}
														{
															add(ruleAction18, position)
														}
														add(ruleAdc16, position1106)
													}
													goto l1095
												l1105:
													position, tokenIndex = position1095, tokenIndex1095
													{
														position1114 := position
														{
															position1115, tokenIndex1115 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1116
															}
															position++
															goto l1115
														l1116:
															position, tokenIndex = position1115, tokenIndex1115
															if buffer[position] != rune('S') {
																goto l1093
															}
															position++
														}
													l1115:
														{
															position1117, tokenIndex1117 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1118
															}
															position++
															goto l1117
														l1118:
															position, tokenIndex = position1117, tokenIndex1117
															if buffer[position] != rune('B') {
																goto l1093
															}
															position++
														}
													l1117:
														{
															position1119, tokenIndex1119 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1120
															}
															position++
															goto l1119
														l1120:
															position, tokenIndex = position1119, tokenIndex1119
															if buffer[position] != rune('C') {
																goto l1093
															}
															position++
														}
													l1119:
														if !_rules[rulews]() {
															goto l1093
														}
														if !_rules[ruleDst16]() {
															goto l1093
														}
														if !_rules[rulesep]() {
															goto l1093
														}
														if !_rules[ruleSrc16]() {
															goto l1093
														}
														{
															add(ruleAction19, position)
														}
														add(ruleSbc16, position1114)
													}
												}
											l1095:
												add(ruleAlu16, position1094)
											}
											goto l985
										l1093:
											position, tokenIndex = position985, tokenIndex985
											{
												position1123 := position
												{
													position1124, tokenIndex1124 := position, tokenIndex
													{
														position1126 := position
														{
															position1127, tokenIndex1127 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1128
															}
															position++
															goto l1127
														l1128:
															position, tokenIndex = position1127, tokenIndex1127
															if buffer[position] != rune('A') {
																goto l1125
															}
															position++
														}
													l1127:
														{
															position1129, tokenIndex1129 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1130
															}
															position++
															goto l1129
														l1130:
															position, tokenIndex = position1129, tokenIndex1129
															if buffer[position] != rune('D') {
																goto l1125
															}
															position++
														}
													l1129:
														{
															position1131, tokenIndex1131 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1132
															}
															position++
															goto l1131
														l1132:
															position, tokenIndex = position1131, tokenIndex1131
															if buffer[position] != rune('D') {
																goto l1125
															}
															position++
														}
													l1131:
														if !_rules[rulews]() {
															goto l1125
														}
														{
															position1133, tokenIndex1133 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1134
															}
															position++
															goto l1133
														l1134:
															position, tokenIndex = position1133, tokenIndex1133
															if buffer[position] != rune('A') {
																goto l1125
															}
															position++
														}
													l1133:
														if !_rules[rulesep]() {
															goto l1125
														}
														if !_rules[ruleSrc8]() {
															goto l1125
														}
														{
															add(ruleAction43, position)
														}
														add(ruleAdd, position1126)
													}
													goto l1124
												l1125:
													position, tokenIndex = position1124, tokenIndex1124
													{
														position1137 := position
														{
															position1138, tokenIndex1138 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1139
															}
															position++
															goto l1138
														l1139:
															position, tokenIndex = position1138, tokenIndex1138
															if buffer[position] != rune('A') {
																goto l1136
															}
															position++
														}
													l1138:
														{
															position1140, tokenIndex1140 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1141
															}
															position++
															goto l1140
														l1141:
															position, tokenIndex = position1140, tokenIndex1140
															if buffer[position] != rune('D') {
																goto l1136
															}
															position++
														}
													l1140:
														{
															position1142, tokenIndex1142 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1143
															}
															position++
															goto l1142
														l1143:
															position, tokenIndex = position1142, tokenIndex1142
															if buffer[position] != rune('C') {
																goto l1136
															}
															position++
														}
													l1142:
														if !_rules[rulews]() {
															goto l1136
														}
														{
															position1144, tokenIndex1144 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1145
															}
															position++
															goto l1144
														l1145:
															position, tokenIndex = position1144, tokenIndex1144
															if buffer[position] != rune('A') {
																goto l1136
															}
															position++
														}
													l1144:
														if !_rules[rulesep]() {
															goto l1136
														}
														if !_rules[ruleSrc8]() {
															goto l1136
														}
														{
															add(ruleAction44, position)
														}
														add(ruleAdc, position1137)
													}
													goto l1124
												l1136:
													position, tokenIndex = position1124, tokenIndex1124
													{
														position1148 := position
														{
															position1149, tokenIndex1149 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1150
															}
															position++
															goto l1149
														l1150:
															position, tokenIndex = position1149, tokenIndex1149
															if buffer[position] != rune('S') {
																goto l1147
															}
															position++
														}
													l1149:
														{
															position1151, tokenIndex1151 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1152
															}
															position++
															goto l1151
														l1152:
															position, tokenIndex = position1151, tokenIndex1151
															if buffer[position] != rune('U') {
																goto l1147
															}
															position++
														}
													l1151:
														{
															position1153, tokenIndex1153 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1154
															}
															position++
															goto l1153
														l1154:
															position, tokenIndex = position1153, tokenIndex1153
															if buffer[position] != rune('B') {
																goto l1147
															}
															position++
														}
													l1153:
														if !_rules[rulews]() {
															goto l1147
														}
														if !_rules[ruleSrc8]() {
															goto l1147
														}
														{
															add(ruleAction45, position)
														}
														add(ruleSub, position1148)
													}
													goto l1124
												l1147:
													position, tokenIndex = position1124, tokenIndex1124
													{
														switch buffer[position] {
														case 'C', 'c':
															{
																position1157 := position
																{
																	position1158, tokenIndex1158 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1159
																	}
																	position++
																	goto l1158
																l1159:
																	position, tokenIndex = position1158, tokenIndex1158
																	if buffer[position] != rune('C') {
																		goto l1122
																	}
																	position++
																}
															l1158:
																{
																	position1160, tokenIndex1160 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1161
																	}
																	position++
																	goto l1160
																l1161:
																	position, tokenIndex = position1160, tokenIndex1160
																	if buffer[position] != rune('P') {
																		goto l1122
																	}
																	position++
																}
															l1160:
																if !_rules[rulews]() {
																	goto l1122
																}
																if !_rules[ruleSrc8]() {
																	goto l1122
																}
																{
																	add(ruleAction50, position)
																}
																add(ruleCp, position1157)
															}
															break
														case 'O', 'o':
															{
																position1163 := position
																{
																	position1164, tokenIndex1164 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1165
																	}
																	position++
																	goto l1164
																l1165:
																	position, tokenIndex = position1164, tokenIndex1164
																	if buffer[position] != rune('O') {
																		goto l1122
																	}
																	position++
																}
															l1164:
																{
																	position1166, tokenIndex1166 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1167
																	}
																	position++
																	goto l1166
																l1167:
																	position, tokenIndex = position1166, tokenIndex1166
																	if buffer[position] != rune('R') {
																		goto l1122
																	}
																	position++
																}
															l1166:
																if !_rules[rulews]() {
																	goto l1122
																}
																if !_rules[ruleSrc8]() {
																	goto l1122
																}
																{
																	add(ruleAction49, position)
																}
																add(ruleOr, position1163)
															}
															break
														case 'X', 'x':
															{
																position1169 := position
																{
																	position1170, tokenIndex1170 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1171
																	}
																	position++
																	goto l1170
																l1171:
																	position, tokenIndex = position1170, tokenIndex1170
																	if buffer[position] != rune('X') {
																		goto l1122
																	}
																	position++
																}
															l1170:
																{
																	position1172, tokenIndex1172 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1173
																	}
																	position++
																	goto l1172
																l1173:
																	position, tokenIndex = position1172, tokenIndex1172
																	if buffer[position] != rune('O') {
																		goto l1122
																	}
																	position++
																}
															l1172:
																{
																	position1174, tokenIndex1174 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1175
																	}
																	position++
																	goto l1174
																l1175:
																	position, tokenIndex = position1174, tokenIndex1174
																	if buffer[position] != rune('R') {
																		goto l1122
																	}
																	position++
																}
															l1174:
																if !_rules[rulews]() {
																	goto l1122
																}
																if !_rules[ruleSrc8]() {
																	goto l1122
																}
																{
																	add(ruleAction48, position)
																}
																add(ruleXor, position1169)
															}
															break
														case 'A', 'a':
															{
																position1177 := position
																{
																	position1178, tokenIndex1178 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1179
																	}
																	position++
																	goto l1178
																l1179:
																	position, tokenIndex = position1178, tokenIndex1178
																	if buffer[position] != rune('A') {
																		goto l1122
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
																		goto l1122
																	}
																	position++
																}
															l1180:
																{
																	position1182, tokenIndex1182 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1183
																	}
																	position++
																	goto l1182
																l1183:
																	position, tokenIndex = position1182, tokenIndex1182
																	if buffer[position] != rune('D') {
																		goto l1122
																	}
																	position++
																}
															l1182:
																if !_rules[rulews]() {
																	goto l1122
																}
																if !_rules[ruleSrc8]() {
																	goto l1122
																}
																{
																	add(ruleAction47, position)
																}
																add(ruleAnd, position1177)
															}
															break
														default:
															{
																position1185 := position
																{
																	position1186, tokenIndex1186 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1187
																	}
																	position++
																	goto l1186
																l1187:
																	position, tokenIndex = position1186, tokenIndex1186
																	if buffer[position] != rune('S') {
																		goto l1122
																	}
																	position++
																}
															l1186:
																{
																	position1188, tokenIndex1188 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1189
																	}
																	position++
																	goto l1188
																l1189:
																	position, tokenIndex = position1188, tokenIndex1188
																	if buffer[position] != rune('B') {
																		goto l1122
																	}
																	position++
																}
															l1188:
																{
																	position1190, tokenIndex1190 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1191
																	}
																	position++
																	goto l1190
																l1191:
																	position, tokenIndex = position1190, tokenIndex1190
																	if buffer[position] != rune('C') {
																		goto l1122
																	}
																	position++
																}
															l1190:
																if !_rules[rulews]() {
																	goto l1122
																}
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
																		goto l1122
																	}
																	position++
																}
															l1192:
																if !_rules[rulesep]() {
																	goto l1122
																}
																if !_rules[ruleSrc8]() {
																	goto l1122
																}
																{
																	add(ruleAction46, position)
																}
																add(ruleSbc, position1185)
															}
															break
														}
													}

												}
											l1124:
												add(ruleAlu, position1123)
											}
											goto l985
										l1122:
											position, tokenIndex = position985, tokenIndex985
											{
												position1196 := position
												{
													position1197, tokenIndex1197 := position, tokenIndex
													{
														position1199 := position
														{
															position1200, tokenIndex1200 := position, tokenIndex
															{
																position1202 := position
																{
																	position1203, tokenIndex1203 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1204
																	}
																	position++
																	goto l1203
																l1204:
																	position, tokenIndex = position1203, tokenIndex1203
																	if buffer[position] != rune('R') {
																		goto l1201
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
																		goto l1201
																	}
																	position++
																}
															l1205:
																{
																	position1207, tokenIndex1207 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1208
																	}
																	position++
																	goto l1207
																l1208:
																	position, tokenIndex = position1207, tokenIndex1207
																	if buffer[position] != rune('C') {
																		goto l1201
																	}
																	position++
																}
															l1207:
																if !_rules[rulews]() {
																	goto l1201
																}
																if !_rules[ruleLoc8]() {
																	goto l1201
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
																	add(ruleAction51, position)
																}
																add(ruleRlc, position1202)
															}
															goto l1200
														l1201:
															position, tokenIndex = position1200, tokenIndex1200
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
																{
																	position1218, tokenIndex1218 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1219
																	}
																	position++
																	goto l1218
																l1219:
																	position, tokenIndex = position1218, tokenIndex1218
																	if buffer[position] != rune('C') {
																		goto l1212
																	}
																	position++
																}
															l1218:
																if !_rules[rulews]() {
																	goto l1212
																}
																if !_rules[ruleLoc8]() {
																	goto l1212
																}
																{
																	position1220, tokenIndex1220 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1220
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1220
																	}
																	goto l1221
																l1220:
																	position, tokenIndex = position1220, tokenIndex1220
																}
															l1221:
																{
																	add(ruleAction52, position)
																}
																add(ruleRrc, position1213)
															}
															goto l1200
														l1212:
															position, tokenIndex = position1200, tokenIndex1200
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
																		goto l1223
																	}
																	position++
																}
															l1225:
																{
																	position1227, tokenIndex1227 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1228
																	}
																	position++
																	goto l1227
																l1228:
																	position, tokenIndex = position1227, tokenIndex1227
																	if buffer[position] != rune('L') {
																		goto l1223
																	}
																	position++
																}
															l1227:
																if !_rules[rulews]() {
																	goto l1223
																}
																if !_rules[ruleLoc8]() {
																	goto l1223
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
																	add(ruleAction53, position)
																}
																add(ruleRl, position1224)
															}
															goto l1200
														l1223:
															position, tokenIndex = position1200, tokenIndex1200
															{
																position1233 := position
																{
																	position1234, tokenIndex1234 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1235
																	}
																	position++
																	goto l1234
																l1235:
																	position, tokenIndex = position1234, tokenIndex1234
																	if buffer[position] != rune('R') {
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
																if !_rules[rulews]() {
																	goto l1232
																}
																if !_rules[ruleLoc8]() {
																	goto l1232
																}
																{
																	position1238, tokenIndex1238 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1238
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1238
																	}
																	goto l1239
																l1238:
																	position, tokenIndex = position1238, tokenIndex1238
																}
															l1239:
																{
																	add(ruleAction54, position)
																}
																add(ruleRr, position1233)
															}
															goto l1200
														l1232:
															position, tokenIndex = position1200, tokenIndex1200
															{
																position1242 := position
																{
																	position1243, tokenIndex1243 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1244
																	}
																	position++
																	goto l1243
																l1244:
																	position, tokenIndex = position1243, tokenIndex1243
																	if buffer[position] != rune('S') {
																		goto l1241
																	}
																	position++
																}
															l1243:
																{
																	position1245, tokenIndex1245 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1246
																	}
																	position++
																	goto l1245
																l1246:
																	position, tokenIndex = position1245, tokenIndex1245
																	if buffer[position] != rune('L') {
																		goto l1241
																	}
																	position++
																}
															l1245:
																{
																	position1247, tokenIndex1247 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1248
																	}
																	position++
																	goto l1247
																l1248:
																	position, tokenIndex = position1247, tokenIndex1247
																	if buffer[position] != rune('A') {
																		goto l1241
																	}
																	position++
																}
															l1247:
																if !_rules[rulews]() {
																	goto l1241
																}
																if !_rules[ruleLoc8]() {
																	goto l1241
																}
																{
																	position1249, tokenIndex1249 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1249
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1249
																	}
																	goto l1250
																l1249:
																	position, tokenIndex = position1249, tokenIndex1249
																}
															l1250:
																{
																	add(ruleAction55, position)
																}
																add(ruleSla, position1242)
															}
															goto l1200
														l1241:
															position, tokenIndex = position1200, tokenIndex1200
															{
																position1253 := position
																{
																	position1254, tokenIndex1254 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1255
																	}
																	position++
																	goto l1254
																l1255:
																	position, tokenIndex = position1254, tokenIndex1254
																	if buffer[position] != rune('S') {
																		goto l1252
																	}
																	position++
																}
															l1254:
																{
																	position1256, tokenIndex1256 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1257
																	}
																	position++
																	goto l1256
																l1257:
																	position, tokenIndex = position1256, tokenIndex1256
																	if buffer[position] != rune('R') {
																		goto l1252
																	}
																	position++
																}
															l1256:
																{
																	position1258, tokenIndex1258 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1259
																	}
																	position++
																	goto l1258
																l1259:
																	position, tokenIndex = position1258, tokenIndex1258
																	if buffer[position] != rune('A') {
																		goto l1252
																	}
																	position++
																}
															l1258:
																if !_rules[rulews]() {
																	goto l1252
																}
																if !_rules[ruleLoc8]() {
																	goto l1252
																}
																{
																	position1260, tokenIndex1260 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1260
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1260
																	}
																	goto l1261
																l1260:
																	position, tokenIndex = position1260, tokenIndex1260
																}
															l1261:
																{
																	add(ruleAction56, position)
																}
																add(ruleSra, position1253)
															}
															goto l1200
														l1252:
															position, tokenIndex = position1200, tokenIndex1200
															{
																position1264 := position
																{
																	position1265, tokenIndex1265 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1266
																	}
																	position++
																	goto l1265
																l1266:
																	position, tokenIndex = position1265, tokenIndex1265
																	if buffer[position] != rune('S') {
																		goto l1263
																	}
																	position++
																}
															l1265:
																{
																	position1267, tokenIndex1267 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1268
																	}
																	position++
																	goto l1267
																l1268:
																	position, tokenIndex = position1267, tokenIndex1267
																	if buffer[position] != rune('L') {
																		goto l1263
																	}
																	position++
																}
															l1267:
																{
																	position1269, tokenIndex1269 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1270
																	}
																	position++
																	goto l1269
																l1270:
																	position, tokenIndex = position1269, tokenIndex1269
																	if buffer[position] != rune('L') {
																		goto l1263
																	}
																	position++
																}
															l1269:
																if !_rules[rulews]() {
																	goto l1263
																}
																if !_rules[ruleLoc8]() {
																	goto l1263
																}
																{
																	position1271, tokenIndex1271 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1271
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1271
																	}
																	goto l1272
																l1271:
																	position, tokenIndex = position1271, tokenIndex1271
																}
															l1272:
																{
																	add(ruleAction57, position)
																}
																add(ruleSll, position1264)
															}
															goto l1200
														l1263:
															position, tokenIndex = position1200, tokenIndex1200
															{
																position1274 := position
																{
																	position1275, tokenIndex1275 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1276
																	}
																	position++
																	goto l1275
																l1276:
																	position, tokenIndex = position1275, tokenIndex1275
																	if buffer[position] != rune('S') {
																		goto l1198
																	}
																	position++
																}
															l1275:
																{
																	position1277, tokenIndex1277 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1278
																	}
																	position++
																	goto l1277
																l1278:
																	position, tokenIndex = position1277, tokenIndex1277
																	if buffer[position] != rune('R') {
																		goto l1198
																	}
																	position++
																}
															l1277:
																{
																	position1279, tokenIndex1279 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1280
																	}
																	position++
																	goto l1279
																l1280:
																	position, tokenIndex = position1279, tokenIndex1279
																	if buffer[position] != rune('L') {
																		goto l1198
																	}
																	position++
																}
															l1279:
																if !_rules[rulews]() {
																	goto l1198
																}
																if !_rules[ruleLoc8]() {
																	goto l1198
																}
																{
																	position1281, tokenIndex1281 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1281
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1281
																	}
																	goto l1282
																l1281:
																	position, tokenIndex = position1281, tokenIndex1281
																}
															l1282:
																{
																	add(ruleAction58, position)
																}
																add(ruleSrl, position1274)
															}
														}
													l1200:
														add(ruleRot, position1199)
													}
													goto l1197
												l1198:
													position, tokenIndex = position1197, tokenIndex1197
													{
														switch buffer[position] {
														case 'S', 's':
															{
																position1285 := position
																{
																	position1286, tokenIndex1286 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1287
																	}
																	position++
																	goto l1286
																l1287:
																	position, tokenIndex = position1286, tokenIndex1286
																	if buffer[position] != rune('S') {
																		goto l1195
																	}
																	position++
																}
															l1286:
																{
																	position1288, tokenIndex1288 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1289
																	}
																	position++
																	goto l1288
																l1289:
																	position, tokenIndex = position1288, tokenIndex1288
																	if buffer[position] != rune('E') {
																		goto l1195
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
																		goto l1195
																	}
																	position++
																}
															l1290:
																if !_rules[rulews]() {
																	goto l1195
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1195
																}
																if !_rules[rulesep]() {
																	goto l1195
																}
																if !_rules[ruleLoc8]() {
																	goto l1195
																}
																{
																	position1292, tokenIndex1292 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1292
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1292
																	}
																	goto l1293
																l1292:
																	position, tokenIndex = position1292, tokenIndex1292
																}
															l1293:
																{
																	add(ruleAction61, position)
																}
																add(ruleSet, position1285)
															}
															break
														case 'R', 'r':
															{
																position1295 := position
																{
																	position1296, tokenIndex1296 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1297
																	}
																	position++
																	goto l1296
																l1297:
																	position, tokenIndex = position1296, tokenIndex1296
																	if buffer[position] != rune('R') {
																		goto l1195
																	}
																	position++
																}
															l1296:
																{
																	position1298, tokenIndex1298 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1299
																	}
																	position++
																	goto l1298
																l1299:
																	position, tokenIndex = position1298, tokenIndex1298
																	if buffer[position] != rune('E') {
																		goto l1195
																	}
																	position++
																}
															l1298:
																{
																	position1300, tokenIndex1300 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1301
																	}
																	position++
																	goto l1300
																l1301:
																	position, tokenIndex = position1300, tokenIndex1300
																	if buffer[position] != rune('S') {
																		goto l1195
																	}
																	position++
																}
															l1300:
																if !_rules[rulews]() {
																	goto l1195
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1195
																}
																if !_rules[rulesep]() {
																	goto l1195
																}
																if !_rules[ruleLoc8]() {
																	goto l1195
																}
																{
																	position1302, tokenIndex1302 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1302
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1302
																	}
																	goto l1303
																l1302:
																	position, tokenIndex = position1302, tokenIndex1302
																}
															l1303:
																{
																	add(ruleAction60, position)
																}
																add(ruleRes, position1295)
															}
															break
														default:
															{
																position1305 := position
																{
																	position1306, tokenIndex1306 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1307
																	}
																	position++
																	goto l1306
																l1307:
																	position, tokenIndex = position1306, tokenIndex1306
																	if buffer[position] != rune('B') {
																		goto l1195
																	}
																	position++
																}
															l1306:
																{
																	position1308, tokenIndex1308 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l1309
																	}
																	position++
																	goto l1308
																l1309:
																	position, tokenIndex = position1308, tokenIndex1308
																	if buffer[position] != rune('I') {
																		goto l1195
																	}
																	position++
																}
															l1308:
																{
																	position1310, tokenIndex1310 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1311
																	}
																	position++
																	goto l1310
																l1311:
																	position, tokenIndex = position1310, tokenIndex1310
																	if buffer[position] != rune('T') {
																		goto l1195
																	}
																	position++
																}
															l1310:
																if !_rules[rulews]() {
																	goto l1195
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1195
																}
																if !_rules[rulesep]() {
																	goto l1195
																}
																if !_rules[ruleLoc8]() {
																	goto l1195
																}
																{
																	add(ruleAction59, position)
																}
																add(ruleBit, position1305)
															}
															break
														}
													}

												}
											l1197:
												add(ruleBitOp, position1196)
											}
											goto l985
										l1195:
											position, tokenIndex = position985, tokenIndex985
											{
												position1314 := position
												{
													position1315, tokenIndex1315 := position, tokenIndex
													{
														position1317 := position
														{
															position1318 := position
															{
																position1319, tokenIndex1319 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1320
																}
																position++
																goto l1319
															l1320:
																position, tokenIndex = position1319, tokenIndex1319
																if buffer[position] != rune('R') {
																	goto l1316
																}
																position++
															}
														l1319:
															{
																position1321, tokenIndex1321 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1322
																}
																position++
																goto l1321
															l1322:
																position, tokenIndex = position1321, tokenIndex1321
																if buffer[position] != rune('E') {
																	goto l1316
																}
																position++
															}
														l1321:
															{
																position1323, tokenIndex1323 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1324
																}
																position++
																goto l1323
															l1324:
																position, tokenIndex = position1323, tokenIndex1323
																if buffer[position] != rune('T') {
																	goto l1316
																}
																position++
															}
														l1323:
															{
																position1325, tokenIndex1325 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1326
																}
																position++
																goto l1325
															l1326:
																position, tokenIndex = position1325, tokenIndex1325
																if buffer[position] != rune('N') {
																	goto l1316
																}
																position++
															}
														l1325:
															add(rulePegText, position1318)
														}
														{
															add(ruleAction76, position)
														}
														add(ruleRetn, position1317)
													}
													goto l1315
												l1316:
													position, tokenIndex = position1315, tokenIndex1315
													{
														position1329 := position
														{
															position1330 := position
															{
																position1331, tokenIndex1331 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1332
																}
																position++
																goto l1331
															l1332:
																position, tokenIndex = position1331, tokenIndex1331
																if buffer[position] != rune('R') {
																	goto l1328
																}
																position++
															}
														l1331:
															{
																position1333, tokenIndex1333 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1334
																}
																position++
																goto l1333
															l1334:
																position, tokenIndex = position1333, tokenIndex1333
																if buffer[position] != rune('E') {
																	goto l1328
																}
																position++
															}
														l1333:
															{
																position1335, tokenIndex1335 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1336
																}
																position++
																goto l1335
															l1336:
																position, tokenIndex = position1335, tokenIndex1335
																if buffer[position] != rune('T') {
																	goto l1328
																}
																position++
															}
														l1335:
															{
																position1337, tokenIndex1337 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1338
																}
																position++
																goto l1337
															l1338:
																position, tokenIndex = position1337, tokenIndex1337
																if buffer[position] != rune('I') {
																	goto l1328
																}
																position++
															}
														l1337:
															add(rulePegText, position1330)
														}
														{
															add(ruleAction77, position)
														}
														add(ruleReti, position1329)
													}
													goto l1315
												l1328:
													position, tokenIndex = position1315, tokenIndex1315
													{
														position1341 := position
														{
															position1342 := position
															{
																position1343, tokenIndex1343 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1344
																}
																position++
																goto l1343
															l1344:
																position, tokenIndex = position1343, tokenIndex1343
																if buffer[position] != rune('R') {
																	goto l1340
																}
																position++
															}
														l1343:
															{
																position1345, tokenIndex1345 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1346
																}
																position++
																goto l1345
															l1346:
																position, tokenIndex = position1345, tokenIndex1345
																if buffer[position] != rune('R') {
																	goto l1340
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
																	goto l1340
																}
																position++
															}
														l1347:
															add(rulePegText, position1342)
														}
														{
															add(ruleAction78, position)
														}
														add(ruleRrd, position1341)
													}
													goto l1315
												l1340:
													position, tokenIndex = position1315, tokenIndex1315
													{
														position1351 := position
														{
															position1352 := position
															{
																position1353, tokenIndex1353 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1354
																}
																position++
																goto l1353
															l1354:
																position, tokenIndex = position1353, tokenIndex1353
																if buffer[position] != rune('I') {
																	goto l1350
																}
																position++
															}
														l1353:
															{
																position1355, tokenIndex1355 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1356
																}
																position++
																goto l1355
															l1356:
																position, tokenIndex = position1355, tokenIndex1355
																if buffer[position] != rune('M') {
																	goto l1350
																}
																position++
															}
														l1355:
															if buffer[position] != rune(' ') {
																goto l1350
															}
															position++
															if buffer[position] != rune('0') {
																goto l1350
															}
															position++
															add(rulePegText, position1352)
														}
														{
															add(ruleAction80, position)
														}
														add(ruleIm0, position1351)
													}
													goto l1315
												l1350:
													position, tokenIndex = position1315, tokenIndex1315
													{
														position1359 := position
														{
															position1360 := position
															{
																position1361, tokenIndex1361 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1362
																}
																position++
																goto l1361
															l1362:
																position, tokenIndex = position1361, tokenIndex1361
																if buffer[position] != rune('I') {
																	goto l1358
																}
																position++
															}
														l1361:
															{
																position1363, tokenIndex1363 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1364
																}
																position++
																goto l1363
															l1364:
																position, tokenIndex = position1363, tokenIndex1363
																if buffer[position] != rune('M') {
																	goto l1358
																}
																position++
															}
														l1363:
															if buffer[position] != rune(' ') {
																goto l1358
															}
															position++
															if buffer[position] != rune('1') {
																goto l1358
															}
															position++
															add(rulePegText, position1360)
														}
														{
															add(ruleAction81, position)
														}
														add(ruleIm1, position1359)
													}
													goto l1315
												l1358:
													position, tokenIndex = position1315, tokenIndex1315
													{
														position1367 := position
														{
															position1368 := position
															{
																position1369, tokenIndex1369 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1370
																}
																position++
																goto l1369
															l1370:
																position, tokenIndex = position1369, tokenIndex1369
																if buffer[position] != rune('I') {
																	goto l1366
																}
																position++
															}
														l1369:
															{
																position1371, tokenIndex1371 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1372
																}
																position++
																goto l1371
															l1372:
																position, tokenIndex = position1371, tokenIndex1371
																if buffer[position] != rune('M') {
																	goto l1366
																}
																position++
															}
														l1371:
															if buffer[position] != rune(' ') {
																goto l1366
															}
															position++
															if buffer[position] != rune('2') {
																goto l1366
															}
															position++
															add(rulePegText, position1368)
														}
														{
															add(ruleAction82, position)
														}
														add(ruleIm2, position1367)
													}
													goto l1315
												l1366:
													position, tokenIndex = position1315, tokenIndex1315
													{
														switch buffer[position] {
														case 'I', 'O', 'i', 'o':
															{
																position1375 := position
																{
																	position1376, tokenIndex1376 := position, tokenIndex
																	{
																		position1378 := position
																		{
																			position1379 := position
																			{
																				position1380, tokenIndex1380 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1381
																				}
																				position++
																				goto l1380
																			l1381:
																				position, tokenIndex = position1380, tokenIndex1380
																				if buffer[position] != rune('I') {
																					goto l1377
																				}
																				position++
																			}
																		l1380:
																			{
																				position1382, tokenIndex1382 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1383
																				}
																				position++
																				goto l1382
																			l1383:
																				position, tokenIndex = position1382, tokenIndex1382
																				if buffer[position] != rune('N') {
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
																			add(ruleAction93, position)
																		}
																		add(ruleInir, position1378)
																	}
																	goto l1376
																l1377:
																	position, tokenIndex = position1376, tokenIndex1376
																	{
																		position1390 := position
																		{
																			position1391 := position
																			{
																				position1392, tokenIndex1392 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1393
																				}
																				position++
																				goto l1392
																			l1393:
																				position, tokenIndex = position1392, tokenIndex1392
																				if buffer[position] != rune('I') {
																					goto l1389
																				}
																				position++
																			}
																		l1392:
																			{
																				position1394, tokenIndex1394 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1395
																				}
																				position++
																				goto l1394
																			l1395:
																				position, tokenIndex = position1394, tokenIndex1394
																				if buffer[position] != rune('N') {
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
																			add(ruleAction85, position)
																		}
																		add(ruleIni, position1390)
																	}
																	goto l1376
																l1389:
																	position, tokenIndex = position1376, tokenIndex1376
																	{
																		position1400 := position
																		{
																			position1401 := position
																			{
																				position1402, tokenIndex1402 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1403
																				}
																				position++
																				goto l1402
																			l1403:
																				position, tokenIndex = position1402, tokenIndex1402
																				if buffer[position] != rune('O') {
																					goto l1399
																				}
																				position++
																			}
																		l1402:
																			{
																				position1404, tokenIndex1404 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1405
																				}
																				position++
																				goto l1404
																			l1405:
																				position, tokenIndex = position1404, tokenIndex1404
																				if buffer[position] != rune('T') {
																					goto l1399
																				}
																				position++
																			}
																		l1404:
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
																			add(ruleAction94, position)
																		}
																		add(ruleOtir, position1400)
																	}
																	goto l1376
																l1399:
																	position, tokenIndex = position1376, tokenIndex1376
																	{
																		position1412 := position
																		{
																			position1413 := position
																			{
																				position1414, tokenIndex1414 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1415
																				}
																				position++
																				goto l1414
																			l1415:
																				position, tokenIndex = position1414, tokenIndex1414
																				if buffer[position] != rune('O') {
																					goto l1411
																				}
																				position++
																			}
																		l1414:
																			{
																				position1416, tokenIndex1416 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1417
																				}
																				position++
																				goto l1416
																			l1417:
																				position, tokenIndex = position1416, tokenIndex1416
																				if buffer[position] != rune('U') {
																					goto l1411
																				}
																				position++
																			}
																		l1416:
																			{
																				position1418, tokenIndex1418 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1419
																				}
																				position++
																				goto l1418
																			l1419:
																				position, tokenIndex = position1418, tokenIndex1418
																				if buffer[position] != rune('T') {
																					goto l1411
																				}
																				position++
																			}
																		l1418:
																			{
																				position1420, tokenIndex1420 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1421
																				}
																				position++
																				goto l1420
																			l1421:
																				position, tokenIndex = position1420, tokenIndex1420
																				if buffer[position] != rune('I') {
																					goto l1411
																				}
																				position++
																			}
																		l1420:
																			add(rulePegText, position1413)
																		}
																		{
																			add(ruleAction86, position)
																		}
																		add(ruleOuti, position1412)
																	}
																	goto l1376
																l1411:
																	position, tokenIndex = position1376, tokenIndex1376
																	{
																		position1424 := position
																		{
																			position1425 := position
																			{
																				position1426, tokenIndex1426 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1427
																				}
																				position++
																				goto l1426
																			l1427:
																				position, tokenIndex = position1426, tokenIndex1426
																				if buffer[position] != rune('I') {
																					goto l1423
																				}
																				position++
																			}
																		l1426:
																			{
																				position1428, tokenIndex1428 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1429
																				}
																				position++
																				goto l1428
																			l1429:
																				position, tokenIndex = position1428, tokenIndex1428
																				if buffer[position] != rune('N') {
																					goto l1423
																				}
																				position++
																			}
																		l1428:
																			{
																				position1430, tokenIndex1430 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1431
																				}
																				position++
																				goto l1430
																			l1431:
																				position, tokenIndex = position1430, tokenIndex1430
																				if buffer[position] != rune('D') {
																					goto l1423
																				}
																				position++
																			}
																		l1430:
																			{
																				position1432, tokenIndex1432 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1433
																				}
																				position++
																				goto l1432
																			l1433:
																				position, tokenIndex = position1432, tokenIndex1432
																				if buffer[position] != rune('R') {
																					goto l1423
																				}
																				position++
																			}
																		l1432:
																			add(rulePegText, position1425)
																		}
																		{
																			add(ruleAction97, position)
																		}
																		add(ruleIndr, position1424)
																	}
																	goto l1376
																l1423:
																	position, tokenIndex = position1376, tokenIndex1376
																	{
																		position1436 := position
																		{
																			position1437 := position
																			{
																				position1438, tokenIndex1438 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1439
																				}
																				position++
																				goto l1438
																			l1439:
																				position, tokenIndex = position1438, tokenIndex1438
																				if buffer[position] != rune('I') {
																					goto l1435
																				}
																				position++
																			}
																		l1438:
																			{
																				position1440, tokenIndex1440 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1441
																				}
																				position++
																				goto l1440
																			l1441:
																				position, tokenIndex = position1440, tokenIndex1440
																				if buffer[position] != rune('N') {
																					goto l1435
																				}
																				position++
																			}
																		l1440:
																			{
																				position1442, tokenIndex1442 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1443
																				}
																				position++
																				goto l1442
																			l1443:
																				position, tokenIndex = position1442, tokenIndex1442
																				if buffer[position] != rune('D') {
																					goto l1435
																				}
																				position++
																			}
																		l1442:
																			add(rulePegText, position1437)
																		}
																		{
																			add(ruleAction89, position)
																		}
																		add(ruleInd, position1436)
																	}
																	goto l1376
																l1435:
																	position, tokenIndex = position1376, tokenIndex1376
																	{
																		position1446 := position
																		{
																			position1447 := position
																			{
																				position1448, tokenIndex1448 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1449
																				}
																				position++
																				goto l1448
																			l1449:
																				position, tokenIndex = position1448, tokenIndex1448
																				if buffer[position] != rune('O') {
																					goto l1445
																				}
																				position++
																			}
																		l1448:
																			{
																				position1450, tokenIndex1450 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1451
																				}
																				position++
																				goto l1450
																			l1451:
																				position, tokenIndex = position1450, tokenIndex1450
																				if buffer[position] != rune('T') {
																					goto l1445
																				}
																				position++
																			}
																		l1450:
																			{
																				position1452, tokenIndex1452 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1453
																				}
																				position++
																				goto l1452
																			l1453:
																				position, tokenIndex = position1452, tokenIndex1452
																				if buffer[position] != rune('D') {
																					goto l1445
																				}
																				position++
																			}
																		l1452:
																			{
																				position1454, tokenIndex1454 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1455
																				}
																				position++
																				goto l1454
																			l1455:
																				position, tokenIndex = position1454, tokenIndex1454
																				if buffer[position] != rune('R') {
																					goto l1445
																				}
																				position++
																			}
																		l1454:
																			add(rulePegText, position1447)
																		}
																		{
																			add(ruleAction98, position)
																		}
																		add(ruleOtdr, position1446)
																	}
																	goto l1376
																l1445:
																	position, tokenIndex = position1376, tokenIndex1376
																	{
																		position1457 := position
																		{
																			position1458 := position
																			{
																				position1459, tokenIndex1459 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1460
																				}
																				position++
																				goto l1459
																			l1460:
																				position, tokenIndex = position1459, tokenIndex1459
																				if buffer[position] != rune('O') {
																					goto l1313
																				}
																				position++
																			}
																		l1459:
																			{
																				position1461, tokenIndex1461 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1462
																				}
																				position++
																				goto l1461
																			l1462:
																				position, tokenIndex = position1461, tokenIndex1461
																				if buffer[position] != rune('U') {
																					goto l1313
																				}
																				position++
																			}
																		l1461:
																			{
																				position1463, tokenIndex1463 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1464
																				}
																				position++
																				goto l1463
																			l1464:
																				position, tokenIndex = position1463, tokenIndex1463
																				if buffer[position] != rune('T') {
																					goto l1313
																				}
																				position++
																			}
																		l1463:
																			{
																				position1465, tokenIndex1465 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1466
																				}
																				position++
																				goto l1465
																			l1466:
																				position, tokenIndex = position1465, tokenIndex1465
																				if buffer[position] != rune('D') {
																					goto l1313
																				}
																				position++
																			}
																		l1465:
																			add(rulePegText, position1458)
																		}
																		{
																			add(ruleAction90, position)
																		}
																		add(ruleOutd, position1457)
																	}
																}
															l1376:
																add(ruleBlitIO, position1375)
															}
															break
														case 'R', 'r':
															{
																position1468 := position
																{
																	position1469 := position
																	{
																		position1470, tokenIndex1470 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1471
																		}
																		position++
																		goto l1470
																	l1471:
																		position, tokenIndex = position1470, tokenIndex1470
																		if buffer[position] != rune('R') {
																			goto l1313
																		}
																		position++
																	}
																l1470:
																	{
																		position1472, tokenIndex1472 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1473
																		}
																		position++
																		goto l1472
																	l1473:
																		position, tokenIndex = position1472, tokenIndex1472
																		if buffer[position] != rune('L') {
																			goto l1313
																		}
																		position++
																	}
																l1472:
																	{
																		position1474, tokenIndex1474 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1475
																		}
																		position++
																		goto l1474
																	l1475:
																		position, tokenIndex = position1474, tokenIndex1474
																		if buffer[position] != rune('D') {
																			goto l1313
																		}
																		position++
																	}
																l1474:
																	add(rulePegText, position1469)
																}
																{
																	add(ruleAction79, position)
																}
																add(ruleRld, position1468)
															}
															break
														case 'N', 'n':
															{
																position1477 := position
																{
																	position1478 := position
																	{
																		position1479, tokenIndex1479 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1480
																		}
																		position++
																		goto l1479
																	l1480:
																		position, tokenIndex = position1479, tokenIndex1479
																		if buffer[position] != rune('N') {
																			goto l1313
																		}
																		position++
																	}
																l1479:
																	{
																		position1481, tokenIndex1481 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1482
																		}
																		position++
																		goto l1481
																	l1482:
																		position, tokenIndex = position1481, tokenIndex1481
																		if buffer[position] != rune('E') {
																			goto l1313
																		}
																		position++
																	}
																l1481:
																	{
																		position1483, tokenIndex1483 := position, tokenIndex
																		if buffer[position] != rune('g') {
																			goto l1484
																		}
																		position++
																		goto l1483
																	l1484:
																		position, tokenIndex = position1483, tokenIndex1483
																		if buffer[position] != rune('G') {
																			goto l1313
																		}
																		position++
																	}
																l1483:
																	add(rulePegText, position1478)
																}
																{
																	add(ruleAction75, position)
																}
																add(ruleNeg, position1477)
															}
															break
														default:
															{
																position1486 := position
																{
																	position1487, tokenIndex1487 := position, tokenIndex
																	{
																		position1489 := position
																		{
																			position1490 := position
																			{
																				position1491, tokenIndex1491 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1492
																				}
																				position++
																				goto l1491
																			l1492:
																				position, tokenIndex = position1491, tokenIndex1491
																				if buffer[position] != rune('L') {
																					goto l1488
																				}
																				position++
																			}
																		l1491:
																			{
																				position1493, tokenIndex1493 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1494
																				}
																				position++
																				goto l1493
																			l1494:
																				position, tokenIndex = position1493, tokenIndex1493
																				if buffer[position] != rune('D') {
																					goto l1488
																				}
																				position++
																			}
																		l1493:
																			{
																				position1495, tokenIndex1495 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1496
																				}
																				position++
																				goto l1495
																			l1496:
																				position, tokenIndex = position1495, tokenIndex1495
																				if buffer[position] != rune('I') {
																					goto l1488
																				}
																				position++
																			}
																		l1495:
																			{
																				position1497, tokenIndex1497 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1498
																				}
																				position++
																				goto l1497
																			l1498:
																				position, tokenIndex = position1497, tokenIndex1497
																				if buffer[position] != rune('R') {
																					goto l1488
																				}
																				position++
																			}
																		l1497:
																			add(rulePegText, position1490)
																		}
																		{
																			add(ruleAction91, position)
																		}
																		add(ruleLdir, position1489)
																	}
																	goto l1487
																l1488:
																	position, tokenIndex = position1487, tokenIndex1487
																	{
																		position1501 := position
																		{
																			position1502 := position
																			{
																				position1503, tokenIndex1503 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1504
																				}
																				position++
																				goto l1503
																			l1504:
																				position, tokenIndex = position1503, tokenIndex1503
																				if buffer[position] != rune('L') {
																					goto l1500
																				}
																				position++
																			}
																		l1503:
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
																					goto l1500
																				}
																				position++
																			}
																		l1505:
																			{
																				position1507, tokenIndex1507 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1508
																				}
																				position++
																				goto l1507
																			l1508:
																				position, tokenIndex = position1507, tokenIndex1507
																				if buffer[position] != rune('I') {
																					goto l1500
																				}
																				position++
																			}
																		l1507:
																			add(rulePegText, position1502)
																		}
																		{
																			add(ruleAction83, position)
																		}
																		add(ruleLdi, position1501)
																	}
																	goto l1487
																l1500:
																	position, tokenIndex = position1487, tokenIndex1487
																	{
																		position1511 := position
																		{
																			position1512 := position
																			{
																				position1513, tokenIndex1513 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1514
																				}
																				position++
																				goto l1513
																			l1514:
																				position, tokenIndex = position1513, tokenIndex1513
																				if buffer[position] != rune('C') {
																					goto l1510
																				}
																				position++
																			}
																		l1513:
																			{
																				position1515, tokenIndex1515 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1516
																				}
																				position++
																				goto l1515
																			l1516:
																				position, tokenIndex = position1515, tokenIndex1515
																				if buffer[position] != rune('P') {
																					goto l1510
																				}
																				position++
																			}
																		l1515:
																			{
																				position1517, tokenIndex1517 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1518
																				}
																				position++
																				goto l1517
																			l1518:
																				position, tokenIndex = position1517, tokenIndex1517
																				if buffer[position] != rune('I') {
																					goto l1510
																				}
																				position++
																			}
																		l1517:
																			{
																				position1519, tokenIndex1519 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1520
																				}
																				position++
																				goto l1519
																			l1520:
																				position, tokenIndex = position1519, tokenIndex1519
																				if buffer[position] != rune('R') {
																					goto l1510
																				}
																				position++
																			}
																		l1519:
																			add(rulePegText, position1512)
																		}
																		{
																			add(ruleAction92, position)
																		}
																		add(ruleCpir, position1511)
																	}
																	goto l1487
																l1510:
																	position, tokenIndex = position1487, tokenIndex1487
																	{
																		position1523 := position
																		{
																			position1524 := position
																			{
																				position1525, tokenIndex1525 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1526
																				}
																				position++
																				goto l1525
																			l1526:
																				position, tokenIndex = position1525, tokenIndex1525
																				if buffer[position] != rune('C') {
																					goto l1522
																				}
																				position++
																			}
																		l1525:
																			{
																				position1527, tokenIndex1527 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1528
																				}
																				position++
																				goto l1527
																			l1528:
																				position, tokenIndex = position1527, tokenIndex1527
																				if buffer[position] != rune('P') {
																					goto l1522
																				}
																				position++
																			}
																		l1527:
																			{
																				position1529, tokenIndex1529 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1530
																				}
																				position++
																				goto l1529
																			l1530:
																				position, tokenIndex = position1529, tokenIndex1529
																				if buffer[position] != rune('I') {
																					goto l1522
																				}
																				position++
																			}
																		l1529:
																			add(rulePegText, position1524)
																		}
																		{
																			add(ruleAction84, position)
																		}
																		add(ruleCpi, position1523)
																	}
																	goto l1487
																l1522:
																	position, tokenIndex = position1487, tokenIndex1487
																	{
																		position1533 := position
																		{
																			position1534 := position
																			{
																				position1535, tokenIndex1535 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1536
																				}
																				position++
																				goto l1535
																			l1536:
																				position, tokenIndex = position1535, tokenIndex1535
																				if buffer[position] != rune('L') {
																					goto l1532
																				}
																				position++
																			}
																		l1535:
																			{
																				position1537, tokenIndex1537 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1538
																				}
																				position++
																				goto l1537
																			l1538:
																				position, tokenIndex = position1537, tokenIndex1537
																				if buffer[position] != rune('D') {
																					goto l1532
																				}
																				position++
																			}
																		l1537:
																			{
																				position1539, tokenIndex1539 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1540
																				}
																				position++
																				goto l1539
																			l1540:
																				position, tokenIndex = position1539, tokenIndex1539
																				if buffer[position] != rune('D') {
																					goto l1532
																				}
																				position++
																			}
																		l1539:
																			{
																				position1541, tokenIndex1541 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1542
																				}
																				position++
																				goto l1541
																			l1542:
																				position, tokenIndex = position1541, tokenIndex1541
																				if buffer[position] != rune('R') {
																					goto l1532
																				}
																				position++
																			}
																		l1541:
																			add(rulePegText, position1534)
																		}
																		{
																			add(ruleAction95, position)
																		}
																		add(ruleLddr, position1533)
																	}
																	goto l1487
																l1532:
																	position, tokenIndex = position1487, tokenIndex1487
																	{
																		position1545 := position
																		{
																			position1546 := position
																			{
																				position1547, tokenIndex1547 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1548
																				}
																				position++
																				goto l1547
																			l1548:
																				position, tokenIndex = position1547, tokenIndex1547
																				if buffer[position] != rune('L') {
																					goto l1544
																				}
																				position++
																			}
																		l1547:
																			{
																				position1549, tokenIndex1549 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1550
																				}
																				position++
																				goto l1549
																			l1550:
																				position, tokenIndex = position1549, tokenIndex1549
																				if buffer[position] != rune('D') {
																					goto l1544
																				}
																				position++
																			}
																		l1549:
																			{
																				position1551, tokenIndex1551 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1552
																				}
																				position++
																				goto l1551
																			l1552:
																				position, tokenIndex = position1551, tokenIndex1551
																				if buffer[position] != rune('D') {
																					goto l1544
																				}
																				position++
																			}
																		l1551:
																			add(rulePegText, position1546)
																		}
																		{
																			add(ruleAction87, position)
																		}
																		add(ruleLdd, position1545)
																	}
																	goto l1487
																l1544:
																	position, tokenIndex = position1487, tokenIndex1487
																	{
																		position1555 := position
																		{
																			position1556 := position
																			{
																				position1557, tokenIndex1557 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1558
																				}
																				position++
																				goto l1557
																			l1558:
																				position, tokenIndex = position1557, tokenIndex1557
																				if buffer[position] != rune('C') {
																					goto l1554
																				}
																				position++
																			}
																		l1557:
																			{
																				position1559, tokenIndex1559 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1560
																				}
																				position++
																				goto l1559
																			l1560:
																				position, tokenIndex = position1559, tokenIndex1559
																				if buffer[position] != rune('P') {
																					goto l1554
																				}
																				position++
																			}
																		l1559:
																			{
																				position1561, tokenIndex1561 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1562
																				}
																				position++
																				goto l1561
																			l1562:
																				position, tokenIndex = position1561, tokenIndex1561
																				if buffer[position] != rune('D') {
																					goto l1554
																				}
																				position++
																			}
																		l1561:
																			{
																				position1563, tokenIndex1563 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1564
																				}
																				position++
																				goto l1563
																			l1564:
																				position, tokenIndex = position1563, tokenIndex1563
																				if buffer[position] != rune('R') {
																					goto l1554
																				}
																				position++
																			}
																		l1563:
																			add(rulePegText, position1556)
																		}
																		{
																			add(ruleAction96, position)
																		}
																		add(ruleCpdr, position1555)
																	}
																	goto l1487
																l1554:
																	position, tokenIndex = position1487, tokenIndex1487
																	{
																		position1566 := position
																		{
																			position1567 := position
																			{
																				position1568, tokenIndex1568 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1569
																				}
																				position++
																				goto l1568
																			l1569:
																				position, tokenIndex = position1568, tokenIndex1568
																				if buffer[position] != rune('C') {
																					goto l1313
																				}
																				position++
																			}
																		l1568:
																			{
																				position1570, tokenIndex1570 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1571
																				}
																				position++
																				goto l1570
																			l1571:
																				position, tokenIndex = position1570, tokenIndex1570
																				if buffer[position] != rune('P') {
																					goto l1313
																				}
																				position++
																			}
																		l1570:
																			{
																				position1572, tokenIndex1572 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1573
																				}
																				position++
																				goto l1572
																			l1573:
																				position, tokenIndex = position1572, tokenIndex1572
																				if buffer[position] != rune('D') {
																					goto l1313
																				}
																				position++
																			}
																		l1572:
																			add(rulePegText, position1567)
																		}
																		{
																			add(ruleAction88, position)
																		}
																		add(ruleCpd, position1566)
																	}
																}
															l1487:
																add(ruleBlit, position1486)
															}
															break
														}
													}

												}
											l1315:
												add(ruleEDSimple, position1314)
											}
											goto l985
										l1313:
											position, tokenIndex = position985, tokenIndex985
											{
												position1576 := position
												{
													position1577, tokenIndex1577 := position, tokenIndex
													{
														position1579 := position
														{
															position1580 := position
															{
																position1581, tokenIndex1581 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1582
																}
																position++
																goto l1581
															l1582:
																position, tokenIndex = position1581, tokenIndex1581
																if buffer[position] != rune('R') {
																	goto l1578
																}
																position++
															}
														l1581:
															{
																position1583, tokenIndex1583 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1584
																}
																position++
																goto l1583
															l1584:
																position, tokenIndex = position1583, tokenIndex1583
																if buffer[position] != rune('L') {
																	goto l1578
																}
																position++
															}
														l1583:
															{
																position1585, tokenIndex1585 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1586
																}
																position++
																goto l1585
															l1586:
																position, tokenIndex = position1585, tokenIndex1585
																if buffer[position] != rune('C') {
																	goto l1578
																}
																position++
															}
														l1585:
															{
																position1587, tokenIndex1587 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1588
																}
																position++
																goto l1587
															l1588:
																position, tokenIndex = position1587, tokenIndex1587
																if buffer[position] != rune('A') {
																	goto l1578
																}
																position++
															}
														l1587:
															add(rulePegText, position1580)
														}
														{
															add(ruleAction64, position)
														}
														add(ruleRlca, position1579)
													}
													goto l1577
												l1578:
													position, tokenIndex = position1577, tokenIndex1577
													{
														position1591 := position
														{
															position1592 := position
															{
																position1593, tokenIndex1593 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1594
																}
																position++
																goto l1593
															l1594:
																position, tokenIndex = position1593, tokenIndex1593
																if buffer[position] != rune('R') {
																	goto l1590
																}
																position++
															}
														l1593:
															{
																position1595, tokenIndex1595 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1596
																}
																position++
																goto l1595
															l1596:
																position, tokenIndex = position1595, tokenIndex1595
																if buffer[position] != rune('R') {
																	goto l1590
																}
																position++
															}
														l1595:
															{
																position1597, tokenIndex1597 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1598
																}
																position++
																goto l1597
															l1598:
																position, tokenIndex = position1597, tokenIndex1597
																if buffer[position] != rune('C') {
																	goto l1590
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
																	goto l1590
																}
																position++
															}
														l1599:
															add(rulePegText, position1592)
														}
														{
															add(ruleAction65, position)
														}
														add(ruleRrca, position1591)
													}
													goto l1577
												l1590:
													position, tokenIndex = position1577, tokenIndex1577
													{
														position1603 := position
														{
															position1604 := position
															{
																position1605, tokenIndex1605 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1606
																}
																position++
																goto l1605
															l1606:
																position, tokenIndex = position1605, tokenIndex1605
																if buffer[position] != rune('R') {
																	goto l1602
																}
																position++
															}
														l1605:
															{
																position1607, tokenIndex1607 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1608
																}
																position++
																goto l1607
															l1608:
																position, tokenIndex = position1607, tokenIndex1607
																if buffer[position] != rune('L') {
																	goto l1602
																}
																position++
															}
														l1607:
															{
																position1609, tokenIndex1609 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1610
																}
																position++
																goto l1609
															l1610:
																position, tokenIndex = position1609, tokenIndex1609
																if buffer[position] != rune('A') {
																	goto l1602
																}
																position++
															}
														l1609:
															add(rulePegText, position1604)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleRla, position1603)
													}
													goto l1577
												l1602:
													position, tokenIndex = position1577, tokenIndex1577
													{
														position1613 := position
														{
															position1614 := position
															{
																position1615, tokenIndex1615 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1616
																}
																position++
																goto l1615
															l1616:
																position, tokenIndex = position1615, tokenIndex1615
																if buffer[position] != rune('D') {
																	goto l1612
																}
																position++
															}
														l1615:
															{
																position1617, tokenIndex1617 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1618
																}
																position++
																goto l1617
															l1618:
																position, tokenIndex = position1617, tokenIndex1617
																if buffer[position] != rune('A') {
																	goto l1612
																}
																position++
															}
														l1617:
															{
																position1619, tokenIndex1619 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1620
																}
																position++
																goto l1619
															l1620:
																position, tokenIndex = position1619, tokenIndex1619
																if buffer[position] != rune('A') {
																	goto l1612
																}
																position++
															}
														l1619:
															add(rulePegText, position1614)
														}
														{
															add(ruleAction68, position)
														}
														add(ruleDaa, position1613)
													}
													goto l1577
												l1612:
													position, tokenIndex = position1577, tokenIndex1577
													{
														position1623 := position
														{
															position1624 := position
															{
																position1625, tokenIndex1625 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1626
																}
																position++
																goto l1625
															l1626:
																position, tokenIndex = position1625, tokenIndex1625
																if buffer[position] != rune('C') {
																	goto l1622
																}
																position++
															}
														l1625:
															{
																position1627, tokenIndex1627 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1628
																}
																position++
																goto l1627
															l1628:
																position, tokenIndex = position1627, tokenIndex1627
																if buffer[position] != rune('P') {
																	goto l1622
																}
																position++
															}
														l1627:
															{
																position1629, tokenIndex1629 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1630
																}
																position++
																goto l1629
															l1630:
																position, tokenIndex = position1629, tokenIndex1629
																if buffer[position] != rune('L') {
																	goto l1622
																}
																position++
															}
														l1629:
															add(rulePegText, position1624)
														}
														{
															add(ruleAction69, position)
														}
														add(ruleCpl, position1623)
													}
													goto l1577
												l1622:
													position, tokenIndex = position1577, tokenIndex1577
													{
														position1633 := position
														{
															position1634 := position
															{
																position1635, tokenIndex1635 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1636
																}
																position++
																goto l1635
															l1636:
																position, tokenIndex = position1635, tokenIndex1635
																if buffer[position] != rune('E') {
																	goto l1632
																}
																position++
															}
														l1635:
															{
																position1637, tokenIndex1637 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1638
																}
																position++
																goto l1637
															l1638:
																position, tokenIndex = position1637, tokenIndex1637
																if buffer[position] != rune('X') {
																	goto l1632
																}
																position++
															}
														l1637:
															{
																position1639, tokenIndex1639 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1640
																}
																position++
																goto l1639
															l1640:
																position, tokenIndex = position1639, tokenIndex1639
																if buffer[position] != rune('X') {
																	goto l1632
																}
																position++
															}
														l1639:
															add(rulePegText, position1634)
														}
														{
															add(ruleAction72, position)
														}
														add(ruleExx, position1633)
													}
													goto l1577
												l1632:
													position, tokenIndex = position1577, tokenIndex1577
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1643 := position
																{
																	position1644 := position
																	{
																		position1645, tokenIndex1645 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1646
																		}
																		position++
																		goto l1645
																	l1646:
																		position, tokenIndex = position1645, tokenIndex1645
																		if buffer[position] != rune('E') {
																			goto l1575
																		}
																		position++
																	}
																l1645:
																	{
																		position1647, tokenIndex1647 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1648
																		}
																		position++
																		goto l1647
																	l1648:
																		position, tokenIndex = position1647, tokenIndex1647
																		if buffer[position] != rune('I') {
																			goto l1575
																		}
																		position++
																	}
																l1647:
																	add(rulePegText, position1644)
																}
																{
																	add(ruleAction74, position)
																}
																add(ruleEi, position1643)
															}
															break
														case 'D', 'd':
															{
																position1650 := position
																{
																	position1651 := position
																	{
																		position1652, tokenIndex1652 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1653
																		}
																		position++
																		goto l1652
																	l1653:
																		position, tokenIndex = position1652, tokenIndex1652
																		if buffer[position] != rune('D') {
																			goto l1575
																		}
																		position++
																	}
																l1652:
																	{
																		position1654, tokenIndex1654 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1655
																		}
																		position++
																		goto l1654
																	l1655:
																		position, tokenIndex = position1654, tokenIndex1654
																		if buffer[position] != rune('I') {
																			goto l1575
																		}
																		position++
																	}
																l1654:
																	add(rulePegText, position1651)
																}
																{
																	add(ruleAction73, position)
																}
																add(ruleDi, position1650)
															}
															break
														case 'C', 'c':
															{
																position1657 := position
																{
																	position1658 := position
																	{
																		position1659, tokenIndex1659 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1660
																		}
																		position++
																		goto l1659
																	l1660:
																		position, tokenIndex = position1659, tokenIndex1659
																		if buffer[position] != rune('C') {
																			goto l1575
																		}
																		position++
																	}
																l1659:
																	{
																		position1661, tokenIndex1661 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1662
																		}
																		position++
																		goto l1661
																	l1662:
																		position, tokenIndex = position1661, tokenIndex1661
																		if buffer[position] != rune('C') {
																			goto l1575
																		}
																		position++
																	}
																l1661:
																	{
																		position1663, tokenIndex1663 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1664
																		}
																		position++
																		goto l1663
																	l1664:
																		position, tokenIndex = position1663, tokenIndex1663
																		if buffer[position] != rune('F') {
																			goto l1575
																		}
																		position++
																	}
																l1663:
																	add(rulePegText, position1658)
																}
																{
																	add(ruleAction71, position)
																}
																add(ruleCcf, position1657)
															}
															break
														case 'S', 's':
															{
																position1666 := position
																{
																	position1667 := position
																	{
																		position1668, tokenIndex1668 := position, tokenIndex
																		if buffer[position] != rune('s') {
																			goto l1669
																		}
																		position++
																		goto l1668
																	l1669:
																		position, tokenIndex = position1668, tokenIndex1668
																		if buffer[position] != rune('S') {
																			goto l1575
																		}
																		position++
																	}
																l1668:
																	{
																		position1670, tokenIndex1670 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1671
																		}
																		position++
																		goto l1670
																	l1671:
																		position, tokenIndex = position1670, tokenIndex1670
																		if buffer[position] != rune('C') {
																			goto l1575
																		}
																		position++
																	}
																l1670:
																	{
																		position1672, tokenIndex1672 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1673
																		}
																		position++
																		goto l1672
																	l1673:
																		position, tokenIndex = position1672, tokenIndex1672
																		if buffer[position] != rune('F') {
																			goto l1575
																		}
																		position++
																	}
																l1672:
																	add(rulePegText, position1667)
																}
																{
																	add(ruleAction70, position)
																}
																add(ruleScf, position1666)
															}
															break
														case 'R', 'r':
															{
																position1675 := position
																{
																	position1676 := position
																	{
																		position1677, tokenIndex1677 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1678
																		}
																		position++
																		goto l1677
																	l1678:
																		position, tokenIndex = position1677, tokenIndex1677
																		if buffer[position] != rune('R') {
																			goto l1575
																		}
																		position++
																	}
																l1677:
																	{
																		position1679, tokenIndex1679 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1680
																		}
																		position++
																		goto l1679
																	l1680:
																		position, tokenIndex = position1679, tokenIndex1679
																		if buffer[position] != rune('R') {
																			goto l1575
																		}
																		position++
																	}
																l1679:
																	{
																		position1681, tokenIndex1681 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1682
																		}
																		position++
																		goto l1681
																	l1682:
																		position, tokenIndex = position1681, tokenIndex1681
																		if buffer[position] != rune('A') {
																			goto l1575
																		}
																		position++
																	}
																l1681:
																	add(rulePegText, position1676)
																}
																{
																	add(ruleAction67, position)
																}
																add(ruleRra, position1675)
															}
															break
														case 'H', 'h':
															{
																position1684 := position
																{
																	position1685 := position
																	{
																		position1686, tokenIndex1686 := position, tokenIndex
																		if buffer[position] != rune('h') {
																			goto l1687
																		}
																		position++
																		goto l1686
																	l1687:
																		position, tokenIndex = position1686, tokenIndex1686
																		if buffer[position] != rune('H') {
																			goto l1575
																		}
																		position++
																	}
																l1686:
																	{
																		position1688, tokenIndex1688 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1689
																		}
																		position++
																		goto l1688
																	l1689:
																		position, tokenIndex = position1688, tokenIndex1688
																		if buffer[position] != rune('A') {
																			goto l1575
																		}
																		position++
																	}
																l1688:
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
																			goto l1575
																		}
																		position++
																	}
																l1690:
																	{
																		position1692, tokenIndex1692 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1693
																		}
																		position++
																		goto l1692
																	l1693:
																		position, tokenIndex = position1692, tokenIndex1692
																		if buffer[position] != rune('T') {
																			goto l1575
																		}
																		position++
																	}
																l1692:
																	add(rulePegText, position1685)
																}
																{
																	add(ruleAction63, position)
																}
																add(ruleHalt, position1684)
															}
															break
														default:
															{
																position1695 := position
																{
																	position1696 := position
																	{
																		position1697, tokenIndex1697 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1698
																		}
																		position++
																		goto l1697
																	l1698:
																		position, tokenIndex = position1697, tokenIndex1697
																		if buffer[position] != rune('N') {
																			goto l1575
																		}
																		position++
																	}
																l1697:
																	{
																		position1699, tokenIndex1699 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1700
																		}
																		position++
																		goto l1699
																	l1700:
																		position, tokenIndex = position1699, tokenIndex1699
																		if buffer[position] != rune('O') {
																			goto l1575
																		}
																		position++
																	}
																l1699:
																	{
																		position1701, tokenIndex1701 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1702
																		}
																		position++
																		goto l1701
																	l1702:
																		position, tokenIndex = position1701, tokenIndex1701
																		if buffer[position] != rune('P') {
																			goto l1575
																		}
																		position++
																	}
																l1701:
																	add(rulePegText, position1696)
																}
																{
																	add(ruleAction62, position)
																}
																add(ruleNop, position1695)
															}
															break
														}
													}

												}
											l1577:
												add(ruleSimple, position1576)
											}
											goto l985
										l1575:
											position, tokenIndex = position985, tokenIndex985
											{
												position1705 := position
												{
													position1706, tokenIndex1706 := position, tokenIndex
													{
														position1708 := position
														{
															position1709, tokenIndex1709 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1710
															}
															position++
															goto l1709
														l1710:
															position, tokenIndex = position1709, tokenIndex1709
															if buffer[position] != rune('R') {
																goto l1707
															}
															position++
														}
													l1709:
														{
															position1711, tokenIndex1711 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1712
															}
															position++
															goto l1711
														l1712:
															position, tokenIndex = position1711, tokenIndex1711
															if buffer[position] != rune('S') {
																goto l1707
															}
															position++
														}
													l1711:
														{
															position1713, tokenIndex1713 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1714
															}
															position++
															goto l1713
														l1714:
															position, tokenIndex = position1713, tokenIndex1713
															if buffer[position] != rune('T') {
																goto l1707
															}
															position++
														}
													l1713:
														if !_rules[rulews]() {
															goto l1707
														}
														if !_rules[rulen]() {
															goto l1707
														}
														{
															add(ruleAction99, position)
														}
														add(ruleRst, position1708)
													}
													goto l1706
												l1707:
													position, tokenIndex = position1706, tokenIndex1706
													{
														position1717 := position
														{
															position1718, tokenIndex1718 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1719
															}
															position++
															goto l1718
														l1719:
															position, tokenIndex = position1718, tokenIndex1718
															if buffer[position] != rune('J') {
																goto l1716
															}
															position++
														}
													l1718:
														{
															position1720, tokenIndex1720 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l1721
															}
															position++
															goto l1720
														l1721:
															position, tokenIndex = position1720, tokenIndex1720
															if buffer[position] != rune('P') {
																goto l1716
															}
															position++
														}
													l1720:
														if !_rules[rulews]() {
															goto l1716
														}
														{
															position1722, tokenIndex1722 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1722
															}
															if !_rules[rulesep]() {
																goto l1722
															}
															goto l1723
														l1722:
															position, tokenIndex = position1722, tokenIndex1722
														}
													l1723:
														if !_rules[ruleSrc16]() {
															goto l1716
														}
														{
															add(ruleAction102, position)
														}
														add(ruleJp, position1717)
													}
													goto l1706
												l1716:
													position, tokenIndex = position1706, tokenIndex1706
													{
														switch buffer[position] {
														case 'D', 'd':
															{
																position1726 := position
																{
																	position1727, tokenIndex1727 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1728
																	}
																	position++
																	goto l1727
																l1728:
																	position, tokenIndex = position1727, tokenIndex1727
																	if buffer[position] != rune('D') {
																		goto l1704
																	}
																	position++
																}
															l1727:
																{
																	position1729, tokenIndex1729 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1730
																	}
																	position++
																	goto l1729
																l1730:
																	position, tokenIndex = position1729, tokenIndex1729
																	if buffer[position] != rune('J') {
																		goto l1704
																	}
																	position++
																}
															l1729:
																{
																	position1731, tokenIndex1731 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1732
																	}
																	position++
																	goto l1731
																l1732:
																	position, tokenIndex = position1731, tokenIndex1731
																	if buffer[position] != rune('N') {
																		goto l1704
																	}
																	position++
																}
															l1731:
																{
																	position1733, tokenIndex1733 := position, tokenIndex
																	if buffer[position] != rune('z') {
																		goto l1734
																	}
																	position++
																	goto l1733
																l1734:
																	position, tokenIndex = position1733, tokenIndex1733
																	if buffer[position] != rune('Z') {
																		goto l1704
																	}
																	position++
																}
															l1733:
																if !_rules[rulews]() {
																	goto l1704
																}
																if !_rules[ruledisp]() {
																	goto l1704
																}
																{
																	add(ruleAction104, position)
																}
																add(ruleDjnz, position1726)
															}
															break
														case 'J', 'j':
															{
																position1736 := position
																{
																	position1737, tokenIndex1737 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1738
																	}
																	position++
																	goto l1737
																l1738:
																	position, tokenIndex = position1737, tokenIndex1737
																	if buffer[position] != rune('J') {
																		goto l1704
																	}
																	position++
																}
															l1737:
																{
																	position1739, tokenIndex1739 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1740
																	}
																	position++
																	goto l1739
																l1740:
																	position, tokenIndex = position1739, tokenIndex1739
																	if buffer[position] != rune('R') {
																		goto l1704
																	}
																	position++
																}
															l1739:
																if !_rules[rulews]() {
																	goto l1704
																}
																{
																	position1741, tokenIndex1741 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1741
																	}
																	if !_rules[rulesep]() {
																		goto l1741
																	}
																	goto l1742
																l1741:
																	position, tokenIndex = position1741, tokenIndex1741
																}
															l1742:
																if !_rules[ruledisp]() {
																	goto l1704
																}
																{
																	add(ruleAction103, position)
																}
																add(ruleJr, position1736)
															}
															break
														case 'R', 'r':
															{
																position1744 := position
																{
																	position1745, tokenIndex1745 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1746
																	}
																	position++
																	goto l1745
																l1746:
																	position, tokenIndex = position1745, tokenIndex1745
																	if buffer[position] != rune('R') {
																		goto l1704
																	}
																	position++
																}
															l1745:
																{
																	position1747, tokenIndex1747 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1748
																	}
																	position++
																	goto l1747
																l1748:
																	position, tokenIndex = position1747, tokenIndex1747
																	if buffer[position] != rune('E') {
																		goto l1704
																	}
																	position++
																}
															l1747:
																{
																	position1749, tokenIndex1749 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1750
																	}
																	position++
																	goto l1749
																l1750:
																	position, tokenIndex = position1749, tokenIndex1749
																	if buffer[position] != rune('T') {
																		goto l1704
																	}
																	position++
																}
															l1749:
																{
																	position1751, tokenIndex1751 := position, tokenIndex
																	if !_rules[rulews]() {
																		goto l1751
																	}
																	if !_rules[rulecc]() {
																		goto l1751
																	}
																	goto l1752
																l1751:
																	position, tokenIndex = position1751, tokenIndex1751
																}
															l1752:
																{
																	add(ruleAction101, position)
																}
																add(ruleRet, position1744)
															}
															break
														default:
															{
																position1754 := position
																{
																	position1755, tokenIndex1755 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1756
																	}
																	position++
																	goto l1755
																l1756:
																	position, tokenIndex = position1755, tokenIndex1755
																	if buffer[position] != rune('C') {
																		goto l1704
																	}
																	position++
																}
															l1755:
																{
																	position1757, tokenIndex1757 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1758
																	}
																	position++
																	goto l1757
																l1758:
																	position, tokenIndex = position1757, tokenIndex1757
																	if buffer[position] != rune('A') {
																		goto l1704
																	}
																	position++
																}
															l1757:
																{
																	position1759, tokenIndex1759 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1760
																	}
																	position++
																	goto l1759
																l1760:
																	position, tokenIndex = position1759, tokenIndex1759
																	if buffer[position] != rune('L') {
																		goto l1704
																	}
																	position++
																}
															l1759:
																{
																	position1761, tokenIndex1761 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1762
																	}
																	position++
																	goto l1761
																l1762:
																	position, tokenIndex = position1761, tokenIndex1761
																	if buffer[position] != rune('L') {
																		goto l1704
																	}
																	position++
																}
															l1761:
																if !_rules[rulews]() {
																	goto l1704
																}
																{
																	position1763, tokenIndex1763 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1763
																	}
																	if !_rules[rulesep]() {
																		goto l1763
																	}
																	goto l1764
																l1763:
																	position, tokenIndex = position1763, tokenIndex1763
																}
															l1764:
																if !_rules[ruleSrc16]() {
																	goto l1704
																}
																{
																	add(ruleAction100, position)
																}
																add(ruleCall, position1754)
															}
															break
														}
													}

												}
											l1706:
												add(ruleJump, position1705)
											}
											goto l985
										l1704:
											position, tokenIndex = position985, tokenIndex985
											{
												position1766 := position
												{
													position1767, tokenIndex1767 := position, tokenIndex
													{
														position1769 := position
														{
															position1770, tokenIndex1770 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1771
															}
															position++
															goto l1770
														l1771:
															position, tokenIndex = position1770, tokenIndex1770
															if buffer[position] != rune('I') {
																goto l1768
															}
															position++
														}
													l1770:
														{
															position1772, tokenIndex1772 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1773
															}
															position++
															goto l1772
														l1773:
															position, tokenIndex = position1772, tokenIndex1772
															if buffer[position] != rune('N') {
																goto l1768
															}
															position++
														}
													l1772:
														if !_rules[rulews]() {
															goto l1768
														}
														if !_rules[ruleReg8]() {
															goto l1768
														}
														if !_rules[rulesep]() {
															goto l1768
														}
														if !_rules[rulePort]() {
															goto l1768
														}
														{
															add(ruleAction105, position)
														}
														add(ruleIN, position1769)
													}
													goto l1767
												l1768:
													position, tokenIndex = position1767, tokenIndex1767
													{
														position1775 := position
														{
															position1776, tokenIndex1776 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1777
															}
															position++
															goto l1776
														l1777:
															position, tokenIndex = position1776, tokenIndex1776
															if buffer[position] != rune('O') {
																goto l911
															}
															position++
														}
													l1776:
														{
															position1778, tokenIndex1778 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1779
															}
															position++
															goto l1778
														l1779:
															position, tokenIndex = position1778, tokenIndex1778
															if buffer[position] != rune('U') {
																goto l911
															}
															position++
														}
													l1778:
														{
															position1780, tokenIndex1780 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1781
															}
															position++
															goto l1780
														l1781:
															position, tokenIndex = position1780, tokenIndex1780
															if buffer[position] != rune('T') {
																goto l911
															}
															position++
														}
													l1780:
														if !_rules[rulews]() {
															goto l911
														}
														if !_rules[rulePort]() {
															goto l911
														}
														if !_rules[rulesep]() {
															goto l911
														}
														if !_rules[ruleReg8]() {
															goto l911
														}
														{
															add(ruleAction106, position)
														}
														add(ruleOUT, position1775)
													}
												}
											l1767:
												add(ruleIO, position1766)
											}
										}
									l985:
										add(ruleInstruction, position984)
									}
								}
							l914:
								add(ruleStatement, position913)
							}
							goto l912
						l911:
							position, tokenIndex = position911, tokenIndex911
						}
					l912:
						{
							position1783, tokenIndex1783 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1783
							}
							goto l1784
						l1783:
							position, tokenIndex = position1783, tokenIndex1783
						}
					l1784:
						{
							position1785, tokenIndex1785 := position, tokenIndex
							{
								position1787 := position
								{
									position1788, tokenIndex1788 := position, tokenIndex
									if buffer[position] != rune(';') {
										goto l1789
									}
									position++
									goto l1788
								l1789:
									position, tokenIndex = position1788, tokenIndex1788
									if buffer[position] != rune('#') {
										goto l1785
									}
									position++
								}
							l1788:
							l1790:
								{
									position1791, tokenIndex1791 := position, tokenIndex
									{
										position1792, tokenIndex1792 := position, tokenIndex
										if buffer[position] != rune('\n') {
											goto l1792
										}
										position++
										goto l1791
									l1792:
										position, tokenIndex = position1792, tokenIndex1792
									}
									if !matchDot() {
										goto l1791
									}
									goto l1790
								l1791:
									position, tokenIndex = position1791, tokenIndex1791
								}
								add(ruleComment, position1787)
							}
							goto l1786
						l1785:
							position, tokenIndex = position1785, tokenIndex1785
						}
					l1786:
						{
							position1793, tokenIndex1793 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1793
							}
							goto l1794
						l1793:
							position, tokenIndex = position1793, tokenIndex1793
						}
					l1794:
						{
							position1795, tokenIndex1795 := position, tokenIndex
							{
								position1797, tokenIndex1797 := position, tokenIndex
								if buffer[position] != rune('\r') {
									goto l1797
								}
								position++
								goto l1798
							l1797:
								position, tokenIndex = position1797, tokenIndex1797
							}
						l1798:
							if buffer[position] != rune('\n') {
								goto l1796
							}
							position++
							goto l1795
						l1796:
							position, tokenIndex = position1795, tokenIndex1795
							if buffer[position] != rune(':') {
								goto l3
							}
							position++
						}
					l1795:
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position902)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position1800, tokenIndex1800 := position, tokenIndex
					if !matchDot() {
						goto l1800
					}
					goto l0
				l1800:
					position, tokenIndex = position1800, tokenIndex1800
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
		/* 3 Directive <- <(Defb / Defs / ((&('D' | 'd') Defw) | (&('O' | 'o') Org) | (&('a') Aseg) | (&('.' | 't') Title)))> */
		nil,
		/* 4 Title <- <('.'? ('t' 'i' 't' 'l' 'e') ws '\'' (!'\'' .)* '\'')> */
		nil,
		/* 5 Aseg <- <('a' 's' 'e' 'g')> */
		nil,
		/* 6 Org <- <(('o' / 'O') ('r' / 'R') ('g' / 'G') ws nn Action1)> */
		nil,
		/* 7 Defb <- <(((('d' / 'D') ('e' / 'E') ('f' / 'F') ('b' / 'B')) / (('d' / 'D') ('b' / 'B'))) ws n Action2)> */
		nil,
		/* 8 Defw <- <(((('d' / 'D') ('e' / 'E') ('f' / 'F') ('w' / 'W')) / (('d' / 'D') ('w' / 'W'))) ws nn Action3)> */
		nil,
		/* 9 Defs <- <(((('d' / 'D') ('e' / 'E') ('f' / 'F') ('s' / 'S')) / (('d' / 'D') ('s' / 'S'))) ws n Action4)> */
		nil,
		/* 10 LabelDefn <- <(LabelText ':' ws Action5)> */
		nil,
		/* 11 LabelText <- <<(alpha alphanum alphanum+)>> */
		func() bool {
			position1811, tokenIndex1811 := position, tokenIndex
			{
				position1812 := position
				{
					position1813 := position
					if !_rules[rulealpha]() {
						goto l1811
					}
					if !_rules[rulealphanum]() {
						goto l1811
					}
					if !_rules[rulealphanum]() {
						goto l1811
					}
				l1814:
					{
						position1815, tokenIndex1815 := position, tokenIndex
						if !_rules[rulealphanum]() {
							goto l1815
						}
						goto l1814
					l1815:
						position, tokenIndex = position1815, tokenIndex1815
					}
					add(rulePegText, position1813)
				}
				add(ruleLabelText, position1812)
			}
			return true
		l1811:
			position, tokenIndex = position1811, tokenIndex1811
			return false
		},
		/* 12 alphanum <- <(alpha / num)> */
		func() bool {
			position1816, tokenIndex1816 := position, tokenIndex
			{
				position1817 := position
				{
					position1818, tokenIndex1818 := position, tokenIndex
					if !_rules[rulealpha]() {
						goto l1819
					}
					goto l1818
				l1819:
					position, tokenIndex = position1818, tokenIndex1818
					{
						position1820 := position
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1816
						}
						position++
						add(rulenum, position1820)
					}
				}
			l1818:
				add(rulealphanum, position1817)
			}
			return true
		l1816:
			position, tokenIndex = position1816, tokenIndex1816
			return false
		},
		/* 13 alpha <- <([a-z] / [A-Z])> */
		func() bool {
			position1821, tokenIndex1821 := position, tokenIndex
			{
				position1822 := position
				{
					position1823, tokenIndex1823 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1824
					}
					position++
					goto l1823
				l1824:
					position, tokenIndex = position1823, tokenIndex1823
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1821
					}
					position++
				}
			l1823:
				add(rulealpha, position1822)
			}
			return true
		l1821:
			position, tokenIndex = position1821, tokenIndex1821
			return false
		},
		/* 14 num <- <[0-9]> */
		nil,
		/* 15 Comment <- <((';' / '#') (!'\n' .)*)> */
		nil,
		/* 16 Instruction <- <(Assignment / Inc / Dec / Alu16 / Alu / BitOp / EDSimple / Simple / Jump / IO)> */
		nil,
		/* 17 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 18 Load <- <(Load16 / Load8)> */
		nil,
		/* 19 Load8 <- <(('l' / 'L') ('d' / 'D') ws Dst8 sep Src8 Action6)> */
		nil,
		/* 20 Load16 <- <(('l' / 'L') ('d' / 'D') ws Dst16 sep Src16 Action7)> */
		nil,
		/* 21 Push <- <(('p' / 'P') ('u' / 'U') ('s' / 'S') ('h' / 'H') ws Src16 Action8)> */
		nil,
		/* 22 Pop <- <(('p' / 'P') ('o' / 'O') ('p' / 'P') ws Dst16 Action9)> */
		nil,
		/* 23 Ex <- <(('e' / 'E') ('x' / 'X') ws Dst16 sep Src16 Action10)> */
		nil,
		/* 24 Inc <- <(Inc16Indexed8 / Inc16 / Inc8)> */
		nil,
		/* 25 Inc16Indexed8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws ILoc8 Action11)> */
		nil,
		/* 26 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action12)> */
		nil,
		/* 27 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action13)> */
		nil,
		/* 28 Dec <- <(Dec16Indexed8 / Dec16 / Dec8)> */
		nil,
		/* 29 Dec16Indexed8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws ILoc8 Action14)> */
		nil,
		/* 30 Dec8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc8 Action15)> */
		nil,
		/* 31 Dec16 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc16 Action16)> */
		nil,
		/* 32 Alu16 <- <(Add16 / Adc16 / Sbc16)> */
		nil,
		/* 33 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action17)> */
		nil,
		/* 34 Adc16 <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws Dst16 sep Src16 Action18)> */
		nil,
		/* 35 Sbc16 <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws Dst16 sep Src16 Action19)> */
		nil,
		/* 36 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action20)> */
		nil,
		/* 37 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action21)> */
		func() bool {
			position1848, tokenIndex1848 := position, tokenIndex
			{
				position1849 := position
				{
					position1850, tokenIndex1850 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1851
					}
					goto l1850
				l1851:
					position, tokenIndex = position1850, tokenIndex1850
					if !_rules[ruleReg8]() {
						goto l1852
					}
					goto l1850
				l1852:
					position, tokenIndex = position1850, tokenIndex1850
					if !_rules[ruleReg16Contents]() {
						goto l1853
					}
					goto l1850
				l1853:
					position, tokenIndex = position1850, tokenIndex1850
					if !_rules[rulenn_contents]() {
						goto l1848
					}
				}
			l1850:
				{
					add(ruleAction21, position)
				}
				add(ruleSrc8, position1849)
			}
			return true
		l1848:
			position, tokenIndex = position1848, tokenIndex1848
			return false
		},
		/* 38 Loc8 <- <((Reg8 / Reg16Contents) Action22)> */
		func() bool {
			position1855, tokenIndex1855 := position, tokenIndex
			{
				position1856 := position
				{
					position1857, tokenIndex1857 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1858
					}
					goto l1857
				l1858:
					position, tokenIndex = position1857, tokenIndex1857
					if !_rules[ruleReg16Contents]() {
						goto l1855
					}
				}
			l1857:
				{
					add(ruleAction22, position)
				}
				add(ruleLoc8, position1856)
			}
			return true
		l1855:
			position, tokenIndex = position1855, tokenIndex1855
			return false
		},
		/* 39 Copy8 <- <(Reg8 Action23)> */
		func() bool {
			position1860, tokenIndex1860 := position, tokenIndex
			{
				position1861 := position
				if !_rules[ruleReg8]() {
					goto l1860
				}
				{
					add(ruleAction23, position)
				}
				add(ruleCopy8, position1861)
			}
			return true
		l1860:
			position, tokenIndex = position1860, tokenIndex1860
			return false
		},
		/* 40 ILoc8 <- <(IReg8 Action24)> */
		func() bool {
			position1863, tokenIndex1863 := position, tokenIndex
			{
				position1864 := position
				if !_rules[ruleIReg8]() {
					goto l1863
				}
				{
					add(ruleAction24, position)
				}
				add(ruleILoc8, position1864)
			}
			return true
		l1863:
			position, tokenIndex = position1863, tokenIndex1863
			return false
		},
		/* 41 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action25)> */
		func() bool {
			position1866, tokenIndex1866 := position, tokenIndex
			{
				position1867 := position
				{
					position1868 := position
					{
						position1869, tokenIndex1869 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1870
						}
						goto l1869
					l1870:
						position, tokenIndex = position1869, tokenIndex1869
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1872 := position
									{
										position1873, tokenIndex1873 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1874
										}
										position++
										goto l1873
									l1874:
										position, tokenIndex = position1873, tokenIndex1873
										if buffer[position] != rune('R') {
											goto l1866
										}
										position++
									}
								l1873:
									add(ruleR, position1872)
								}
								break
							case 'I', 'i':
								{
									position1875 := position
									{
										position1876, tokenIndex1876 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1877
										}
										position++
										goto l1876
									l1877:
										position, tokenIndex = position1876, tokenIndex1876
										if buffer[position] != rune('I') {
											goto l1866
										}
										position++
									}
								l1876:
									add(ruleI, position1875)
								}
								break
							case 'L', 'l':
								{
									position1878 := position
									{
										position1879, tokenIndex1879 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1880
										}
										position++
										goto l1879
									l1880:
										position, tokenIndex = position1879, tokenIndex1879
										if buffer[position] != rune('L') {
											goto l1866
										}
										position++
									}
								l1879:
									add(ruleL, position1878)
								}
								break
							case 'H', 'h':
								{
									position1881 := position
									{
										position1882, tokenIndex1882 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1883
										}
										position++
										goto l1882
									l1883:
										position, tokenIndex = position1882, tokenIndex1882
										if buffer[position] != rune('H') {
											goto l1866
										}
										position++
									}
								l1882:
									add(ruleH, position1881)
								}
								break
							case 'E', 'e':
								{
									position1884 := position
									{
										position1885, tokenIndex1885 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1886
										}
										position++
										goto l1885
									l1886:
										position, tokenIndex = position1885, tokenIndex1885
										if buffer[position] != rune('E') {
											goto l1866
										}
										position++
									}
								l1885:
									add(ruleE, position1884)
								}
								break
							case 'D', 'd':
								{
									position1887 := position
									{
										position1888, tokenIndex1888 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1889
										}
										position++
										goto l1888
									l1889:
										position, tokenIndex = position1888, tokenIndex1888
										if buffer[position] != rune('D') {
											goto l1866
										}
										position++
									}
								l1888:
									add(ruleD, position1887)
								}
								break
							case 'C', 'c':
								{
									position1890 := position
									{
										position1891, tokenIndex1891 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1892
										}
										position++
										goto l1891
									l1892:
										position, tokenIndex = position1891, tokenIndex1891
										if buffer[position] != rune('C') {
											goto l1866
										}
										position++
									}
								l1891:
									add(ruleC, position1890)
								}
								break
							case 'B', 'b':
								{
									position1893 := position
									{
										position1894, tokenIndex1894 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1895
										}
										position++
										goto l1894
									l1895:
										position, tokenIndex = position1894, tokenIndex1894
										if buffer[position] != rune('B') {
											goto l1866
										}
										position++
									}
								l1894:
									add(ruleB, position1893)
								}
								break
							case 'F', 'f':
								{
									position1896 := position
									{
										position1897, tokenIndex1897 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1898
										}
										position++
										goto l1897
									l1898:
										position, tokenIndex = position1897, tokenIndex1897
										if buffer[position] != rune('F') {
											goto l1866
										}
										position++
									}
								l1897:
									add(ruleF, position1896)
								}
								break
							default:
								{
									position1899 := position
									{
										position1900, tokenIndex1900 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1901
										}
										position++
										goto l1900
									l1901:
										position, tokenIndex = position1900, tokenIndex1900
										if buffer[position] != rune('A') {
											goto l1866
										}
										position++
									}
								l1900:
									add(ruleA, position1899)
								}
								break
							}
						}

					}
				l1869:
					add(rulePegText, position1868)
				}
				{
					add(ruleAction25, position)
				}
				add(ruleReg8, position1867)
			}
			return true
		l1866:
			position, tokenIndex = position1866, tokenIndex1866
			return false
		},
		/* 42 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action26)> */
		func() bool {
			position1903, tokenIndex1903 := position, tokenIndex
			{
				position1904 := position
				{
					position1905 := position
					{
						position1906, tokenIndex1906 := position, tokenIndex
						{
							position1908 := position
							{
								position1909, tokenIndex1909 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1910
								}
								position++
								goto l1909
							l1910:
								position, tokenIndex = position1909, tokenIndex1909
								if buffer[position] != rune('I') {
									goto l1907
								}
								position++
							}
						l1909:
							{
								position1911, tokenIndex1911 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1912
								}
								position++
								goto l1911
							l1912:
								position, tokenIndex = position1911, tokenIndex1911
								if buffer[position] != rune('X') {
									goto l1907
								}
								position++
							}
						l1911:
							{
								position1913, tokenIndex1913 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1914
								}
								position++
								goto l1913
							l1914:
								position, tokenIndex = position1913, tokenIndex1913
								if buffer[position] != rune('H') {
									goto l1907
								}
								position++
							}
						l1913:
							add(ruleIXH, position1908)
						}
						goto l1906
					l1907:
						position, tokenIndex = position1906, tokenIndex1906
						{
							position1916 := position
							{
								position1917, tokenIndex1917 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1918
								}
								position++
								goto l1917
							l1918:
								position, tokenIndex = position1917, tokenIndex1917
								if buffer[position] != rune('I') {
									goto l1915
								}
								position++
							}
						l1917:
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
									goto l1915
								}
								position++
							}
						l1919:
							{
								position1921, tokenIndex1921 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1922
								}
								position++
								goto l1921
							l1922:
								position, tokenIndex = position1921, tokenIndex1921
								if buffer[position] != rune('L') {
									goto l1915
								}
								position++
							}
						l1921:
							add(ruleIXL, position1916)
						}
						goto l1906
					l1915:
						position, tokenIndex = position1906, tokenIndex1906
						{
							position1924 := position
							{
								position1925, tokenIndex1925 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1926
								}
								position++
								goto l1925
							l1926:
								position, tokenIndex = position1925, tokenIndex1925
								if buffer[position] != rune('I') {
									goto l1923
								}
								position++
							}
						l1925:
							{
								position1927, tokenIndex1927 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1928
								}
								position++
								goto l1927
							l1928:
								position, tokenIndex = position1927, tokenIndex1927
								if buffer[position] != rune('Y') {
									goto l1923
								}
								position++
							}
						l1927:
							{
								position1929, tokenIndex1929 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1930
								}
								position++
								goto l1929
							l1930:
								position, tokenIndex = position1929, tokenIndex1929
								if buffer[position] != rune('H') {
									goto l1923
								}
								position++
							}
						l1929:
							add(ruleIYH, position1924)
						}
						goto l1906
					l1923:
						position, tokenIndex = position1906, tokenIndex1906
						{
							position1931 := position
							{
								position1932, tokenIndex1932 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1933
								}
								position++
								goto l1932
							l1933:
								position, tokenIndex = position1932, tokenIndex1932
								if buffer[position] != rune('I') {
									goto l1903
								}
								position++
							}
						l1932:
							{
								position1934, tokenIndex1934 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1935
								}
								position++
								goto l1934
							l1935:
								position, tokenIndex = position1934, tokenIndex1934
								if buffer[position] != rune('Y') {
									goto l1903
								}
								position++
							}
						l1934:
							{
								position1936, tokenIndex1936 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1937
								}
								position++
								goto l1936
							l1937:
								position, tokenIndex = position1936, tokenIndex1936
								if buffer[position] != rune('L') {
									goto l1903
								}
								position++
							}
						l1936:
							add(ruleIYL, position1931)
						}
					}
				l1906:
					add(rulePegText, position1905)
				}
				{
					add(ruleAction26, position)
				}
				add(ruleIReg8, position1904)
			}
			return true
		l1903:
			position, tokenIndex = position1903, tokenIndex1903
			return false
		},
		/* 43 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action27)> */
		func() bool {
			position1939, tokenIndex1939 := position, tokenIndex
			{
				position1940 := position
				{
					position1941, tokenIndex1941 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1942
					}
					goto l1941
				l1942:
					position, tokenIndex = position1941, tokenIndex1941
					if !_rules[rulenn_contents]() {
						goto l1943
					}
					goto l1941
				l1943:
					position, tokenIndex = position1941, tokenIndex1941
					if !_rules[ruleReg16Contents]() {
						goto l1939
					}
				}
			l1941:
				{
					add(ruleAction27, position)
				}
				add(ruleDst16, position1940)
			}
			return true
		l1939:
			position, tokenIndex = position1939, tokenIndex1939
			return false
		},
		/* 44 Src16 <- <((Reg16 / nn / nn_contents) Action28)> */
		func() bool {
			position1945, tokenIndex1945 := position, tokenIndex
			{
				position1946 := position
				{
					position1947, tokenIndex1947 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1948
					}
					goto l1947
				l1948:
					position, tokenIndex = position1947, tokenIndex1947
					if !_rules[rulenn]() {
						goto l1949
					}
					goto l1947
				l1949:
					position, tokenIndex = position1947, tokenIndex1947
					if !_rules[rulenn_contents]() {
						goto l1945
					}
				}
			l1947:
				{
					add(ruleAction28, position)
				}
				add(ruleSrc16, position1946)
			}
			return true
		l1945:
			position, tokenIndex = position1945, tokenIndex1945
			return false
		},
		/* 45 Loc16 <- <(Reg16 Action29)> */
		func() bool {
			position1951, tokenIndex1951 := position, tokenIndex
			{
				position1952 := position
				if !_rules[ruleReg16]() {
					goto l1951
				}
				{
					add(ruleAction29, position)
				}
				add(ruleLoc16, position1952)
			}
			return true
		l1951:
			position, tokenIndex = position1951, tokenIndex1951
			return false
		},
		/* 46 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action30)> */
		func() bool {
			position1954, tokenIndex1954 := position, tokenIndex
			{
				position1955 := position
				{
					position1956 := position
					{
						position1957, tokenIndex1957 := position, tokenIndex
						{
							position1959 := position
							{
								position1960, tokenIndex1960 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1961
								}
								position++
								goto l1960
							l1961:
								position, tokenIndex = position1960, tokenIndex1960
								if buffer[position] != rune('A') {
									goto l1958
								}
								position++
							}
						l1960:
							{
								position1962, tokenIndex1962 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1963
								}
								position++
								goto l1962
							l1963:
								position, tokenIndex = position1962, tokenIndex1962
								if buffer[position] != rune('F') {
									goto l1958
								}
								position++
							}
						l1962:
							if buffer[position] != rune('\'') {
								goto l1958
							}
							position++
							add(ruleAF_PRIME, position1959)
						}
						goto l1957
					l1958:
						position, tokenIndex = position1957, tokenIndex1957
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1954
								}
								break
							case 'S', 's':
								{
									position1965 := position
									{
										position1966, tokenIndex1966 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1967
										}
										position++
										goto l1966
									l1967:
										position, tokenIndex = position1966, tokenIndex1966
										if buffer[position] != rune('S') {
											goto l1954
										}
										position++
									}
								l1966:
									{
										position1968, tokenIndex1968 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1969
										}
										position++
										goto l1968
									l1969:
										position, tokenIndex = position1968, tokenIndex1968
										if buffer[position] != rune('P') {
											goto l1954
										}
										position++
									}
								l1968:
									add(ruleSP, position1965)
								}
								break
							case 'H', 'h':
								{
									position1970 := position
									{
										position1971, tokenIndex1971 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1972
										}
										position++
										goto l1971
									l1972:
										position, tokenIndex = position1971, tokenIndex1971
										if buffer[position] != rune('H') {
											goto l1954
										}
										position++
									}
								l1971:
									{
										position1973, tokenIndex1973 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1974
										}
										position++
										goto l1973
									l1974:
										position, tokenIndex = position1973, tokenIndex1973
										if buffer[position] != rune('L') {
											goto l1954
										}
										position++
									}
								l1973:
									add(ruleHL, position1970)
								}
								break
							case 'D', 'd':
								{
									position1975 := position
									{
										position1976, tokenIndex1976 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1977
										}
										position++
										goto l1976
									l1977:
										position, tokenIndex = position1976, tokenIndex1976
										if buffer[position] != rune('D') {
											goto l1954
										}
										position++
									}
								l1976:
									{
										position1978, tokenIndex1978 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1979
										}
										position++
										goto l1978
									l1979:
										position, tokenIndex = position1978, tokenIndex1978
										if buffer[position] != rune('E') {
											goto l1954
										}
										position++
									}
								l1978:
									add(ruleDE, position1975)
								}
								break
							case 'B', 'b':
								{
									position1980 := position
									{
										position1981, tokenIndex1981 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1982
										}
										position++
										goto l1981
									l1982:
										position, tokenIndex = position1981, tokenIndex1981
										if buffer[position] != rune('B') {
											goto l1954
										}
										position++
									}
								l1981:
									{
										position1983, tokenIndex1983 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1984
										}
										position++
										goto l1983
									l1984:
										position, tokenIndex = position1983, tokenIndex1983
										if buffer[position] != rune('C') {
											goto l1954
										}
										position++
									}
								l1983:
									add(ruleBC, position1980)
								}
								break
							default:
								{
									position1985 := position
									{
										position1986, tokenIndex1986 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1987
										}
										position++
										goto l1986
									l1987:
										position, tokenIndex = position1986, tokenIndex1986
										if buffer[position] != rune('A') {
											goto l1954
										}
										position++
									}
								l1986:
									{
										position1988, tokenIndex1988 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1989
										}
										position++
										goto l1988
									l1989:
										position, tokenIndex = position1988, tokenIndex1988
										if buffer[position] != rune('F') {
											goto l1954
										}
										position++
									}
								l1988:
									add(ruleAF, position1985)
								}
								break
							}
						}

					}
				l1957:
					add(rulePegText, position1956)
				}
				{
					add(ruleAction30, position)
				}
				add(ruleReg16, position1955)
			}
			return true
		l1954:
			position, tokenIndex = position1954, tokenIndex1954
			return false
		},
		/* 47 IReg16 <- <(<(IX / IY)> Action31)> */
		func() bool {
			position1991, tokenIndex1991 := position, tokenIndex
			{
				position1992 := position
				{
					position1993 := position
					{
						position1994, tokenIndex1994 := position, tokenIndex
						{
							position1996 := position
							{
								position1997, tokenIndex1997 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1998
								}
								position++
								goto l1997
							l1998:
								position, tokenIndex = position1997, tokenIndex1997
								if buffer[position] != rune('I') {
									goto l1995
								}
								position++
							}
						l1997:
							{
								position1999, tokenIndex1999 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l2000
								}
								position++
								goto l1999
							l2000:
								position, tokenIndex = position1999, tokenIndex1999
								if buffer[position] != rune('X') {
									goto l1995
								}
								position++
							}
						l1999:
							add(ruleIX, position1996)
						}
						goto l1994
					l1995:
						position, tokenIndex = position1994, tokenIndex1994
						{
							position2001 := position
							{
								position2002, tokenIndex2002 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l2003
								}
								position++
								goto l2002
							l2003:
								position, tokenIndex = position2002, tokenIndex2002
								if buffer[position] != rune('I') {
									goto l1991
								}
								position++
							}
						l2002:
							{
								position2004, tokenIndex2004 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l2005
								}
								position++
								goto l2004
							l2005:
								position, tokenIndex = position2004, tokenIndex2004
								if buffer[position] != rune('Y') {
									goto l1991
								}
								position++
							}
						l2004:
							add(ruleIY, position2001)
						}
					}
				l1994:
					add(rulePegText, position1993)
				}
				{
					add(ruleAction31, position)
				}
				add(ruleIReg16, position1992)
			}
			return true
		l1991:
			position, tokenIndex = position1991, tokenIndex1991
			return false
		},
		/* 48 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position2007, tokenIndex2007 := position, tokenIndex
			{
				position2008 := position
				{
					position2009, tokenIndex2009 := position, tokenIndex
					{
						position2011 := position
						if buffer[position] != rune('(') {
							goto l2010
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l2010
						}
						{
							position2012, tokenIndex2012 := position, tokenIndex
							if !_rules[rulews]() {
								goto l2012
							}
							goto l2013
						l2012:
							position, tokenIndex = position2012, tokenIndex2012
						}
					l2013:
						if !_rules[ruledisp]() {
							goto l2010
						}
						{
							position2014, tokenIndex2014 := position, tokenIndex
							if !_rules[rulews]() {
								goto l2014
							}
							goto l2015
						l2014:
							position, tokenIndex = position2014, tokenIndex2014
						}
					l2015:
						if buffer[position] != rune(')') {
							goto l2010
						}
						position++
						{
							add(ruleAction33, position)
						}
						add(ruleIndexedR16C, position2011)
					}
					goto l2009
				l2010:
					position, tokenIndex = position2009, tokenIndex2009
					{
						position2017 := position
						if buffer[position] != rune('(') {
							goto l2007
						}
						position++
						if !_rules[ruleReg16]() {
							goto l2007
						}
						if buffer[position] != rune(')') {
							goto l2007
						}
						position++
						{
							add(ruleAction32, position)
						}
						add(rulePlainR16C, position2017)
					}
				}
			l2009:
				add(ruleReg16Contents, position2008)
			}
			return true
		l2007:
			position, tokenIndex = position2007, tokenIndex2007
			return false
		},
		/* 49 PlainR16C <- <('(' Reg16 ')' Action32)> */
		nil,
		/* 50 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action33)> */
		nil,
		/* 51 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position2021, tokenIndex2021 := position, tokenIndex
			{
				position2022 := position
				{
					position2023, tokenIndex2023 := position, tokenIndex
					{
						position2025 := position
						{
							position2026 := position
							if !_rules[rulehexdigit]() {
								goto l2024
							}
							if !_rules[rulehexdigit]() {
								goto l2024
							}
							add(rulePegText, position2026)
						}
						{
							position2027, tokenIndex2027 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2028
							}
							position++
							goto l2027
						l2028:
							position, tokenIndex = position2027, tokenIndex2027
							if buffer[position] != rune('H') {
								goto l2024
							}
							position++
						}
					l2027:
						{
							add(ruleAction37, position)
						}
						add(rulehexByteH, position2025)
					}
					goto l2023
				l2024:
					position, tokenIndex = position2023, tokenIndex2023
					{
						position2031 := position
						if buffer[position] != rune('0') {
							goto l2030
						}
						position++
						{
							position2032, tokenIndex2032 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2033
							}
							position++
							goto l2032
						l2033:
							position, tokenIndex = position2032, tokenIndex2032
							if buffer[position] != rune('X') {
								goto l2030
							}
							position++
						}
					l2032:
						{
							position2034 := position
							if !_rules[rulehexdigit]() {
								goto l2030
							}
							if !_rules[rulehexdigit]() {
								goto l2030
							}
							add(rulePegText, position2034)
						}
						{
							add(ruleAction38, position)
						}
						add(rulehexByte0x, position2031)
					}
					goto l2023
				l2030:
					position, tokenIndex = position2023, tokenIndex2023
					{
						position2036 := position
						{
							position2037 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2021
							}
							position++
						l2038:
							{
								position2039, tokenIndex2039 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2039
								}
								position++
								goto l2038
							l2039:
								position, tokenIndex = position2039, tokenIndex2039
							}
							add(rulePegText, position2037)
						}
						{
							add(ruleAction39, position)
						}
						add(ruledecimalByte, position2036)
					}
				}
			l2023:
				add(rulen, position2022)
			}
			return true
		l2021:
			position, tokenIndex = position2021, tokenIndex2021
			return false
		},
		/* 52 nn <- <(LabelNN / hexWordH / hexWord0x)> */
		func() bool {
			position2041, tokenIndex2041 := position, tokenIndex
			{
				position2042 := position
				{
					position2043, tokenIndex2043 := position, tokenIndex
					{
						position2045 := position
						{
							position2046 := position
							if !_rules[ruleLabelText]() {
								goto l2044
							}
							add(rulePegText, position2046)
						}
						{
							add(ruleAction40, position)
						}
						add(ruleLabelNN, position2045)
					}
					goto l2043
				l2044:
					position, tokenIndex = position2043, tokenIndex2043
					{
						position2049 := position
						{
							position2050, tokenIndex2050 := position, tokenIndex
							if !_rules[rulezeroHexWord]() {
								goto l2051
							}
							goto l2050
						l2051:
							position, tokenIndex = position2050, tokenIndex2050
							if !_rules[rulehexWord]() {
								goto l2048
							}
						}
					l2050:
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
								goto l2048
							}
							position++
						}
					l2052:
						add(rulehexWordH, position2049)
					}
					goto l2043
				l2048:
					position, tokenIndex = position2043, tokenIndex2043
					{
						position2054 := position
						if buffer[position] != rune('0') {
							goto l2041
						}
						position++
						{
							position2055, tokenIndex2055 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2056
							}
							position++
							goto l2055
						l2056:
							position, tokenIndex = position2055, tokenIndex2055
							if buffer[position] != rune('X') {
								goto l2041
							}
							position++
						}
					l2055:
						{
							position2057, tokenIndex2057 := position, tokenIndex
							if !_rules[rulezeroHexWord]() {
								goto l2058
							}
							goto l2057
						l2058:
							position, tokenIndex = position2057, tokenIndex2057
							if !_rules[rulehexWord]() {
								goto l2041
							}
						}
					l2057:
						add(rulehexWord0x, position2054)
					}
				}
			l2043:
				add(rulenn, position2042)
			}
			return true
		l2041:
			position, tokenIndex = position2041, tokenIndex2041
			return false
		},
		/* 53 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position2059, tokenIndex2059 := position, tokenIndex
			{
				position2060 := position
				{
					position2061, tokenIndex2061 := position, tokenIndex
					{
						position2063 := position
						{
							position2064 := position
							{
								position2065, tokenIndex2065 := position, tokenIndex
								{
									position2067, tokenIndex2067 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2068
									}
									position++
									goto l2067
								l2068:
									position, tokenIndex = position2067, tokenIndex2067
									if buffer[position] != rune('+') {
										goto l2065
									}
									position++
								}
							l2067:
								goto l2066
							l2065:
								position, tokenIndex = position2065, tokenIndex2065
							}
						l2066:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2062
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2062
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2062
									}
									position++
									break
								}
							}

						l2069:
							{
								position2070, tokenIndex2070 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2070
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2070
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2070
										}
										position++
										break
									}
								}

								goto l2069
							l2070:
								position, tokenIndex = position2070, tokenIndex2070
							}
							add(rulePegText, position2064)
						}
						{
							position2073, tokenIndex2073 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2074
							}
							position++
							goto l2073
						l2074:
							position, tokenIndex = position2073, tokenIndex2073
							if buffer[position] != rune('H') {
								goto l2062
							}
							position++
						}
					l2073:
						{
							add(ruleAction35, position)
						}
						add(rulesignedHexByteH, position2063)
					}
					goto l2061
				l2062:
					position, tokenIndex = position2061, tokenIndex2061
					{
						position2077 := position
						{
							position2078 := position
							{
								position2079, tokenIndex2079 := position, tokenIndex
								{
									position2081, tokenIndex2081 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2082
									}
									position++
									goto l2081
								l2082:
									position, tokenIndex = position2081, tokenIndex2081
									if buffer[position] != rune('+') {
										goto l2079
									}
									position++
								}
							l2081:
								goto l2080
							l2079:
								position, tokenIndex = position2079, tokenIndex2079
							}
						l2080:
							if buffer[position] != rune('0') {
								goto l2076
							}
							position++
							{
								position2083, tokenIndex2083 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l2084
								}
								position++
								goto l2083
							l2084:
								position, tokenIndex = position2083, tokenIndex2083
								if buffer[position] != rune('X') {
									goto l2076
								}
								position++
							}
						l2083:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2076
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2076
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2076
									}
									position++
									break
								}
							}

						l2085:
							{
								position2086, tokenIndex2086 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2086
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2086
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2086
										}
										position++
										break
									}
								}

								goto l2085
							l2086:
								position, tokenIndex = position2086, tokenIndex2086
							}
							add(rulePegText, position2078)
						}
						{
							add(ruleAction36, position)
						}
						add(rulesignedHexByte0x, position2077)
					}
					goto l2061
				l2076:
					position, tokenIndex = position2061, tokenIndex2061
					{
						position2090 := position
						{
							position2091 := position
							{
								position2092, tokenIndex2092 := position, tokenIndex
								{
									position2094, tokenIndex2094 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2095
									}
									position++
									goto l2094
								l2095:
									position, tokenIndex = position2094, tokenIndex2094
									if buffer[position] != rune('+') {
										goto l2092
									}
									position++
								}
							l2094:
								goto l2093
							l2092:
								position, tokenIndex = position2092, tokenIndex2092
							}
						l2093:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2059
							}
							position++
						l2096:
							{
								position2097, tokenIndex2097 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2097
								}
								position++
								goto l2096
							l2097:
								position, tokenIndex = position2097, tokenIndex2097
							}
							add(rulePegText, position2091)
						}
						{
							add(ruleAction34, position)
						}
						add(rulesignedDecimalByte, position2090)
					}
				}
			l2061:
				add(ruledisp, position2060)
			}
			return true
		l2059:
			position, tokenIndex = position2059, tokenIndex2059
			return false
		},
		/* 54 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action34)> */
		nil,
		/* 55 signedHexByteH <- <(<(('-' / '+')? ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> ('h' / 'H') Action35)> */
		nil,
		/* 56 signedHexByte0x <- <(<(('-' / '+')? ('0' ('x' / 'X')) ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> Action36)> */
		nil,
		/* 57 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action37)> */
		nil,
		/* 58 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action38)> */
		nil,
		/* 59 decimalByte <- <(<[0-9]+> Action39)> */
		nil,
		/* 60 LabelNN <- <(<LabelText> Action40)> */
		nil,
		/* 61 hexWordH <- <((zeroHexWord / hexWord) ('h' / 'H'))> */
		nil,
		/* 62 hexWord0x <- <('0' ('x' / 'X') (zeroHexWord / hexWord))> */
		nil,
		/* 63 hexWord <- <(<(hexdigit hexdigit hexdigit hexdigit)> Action41)> */
		func() bool {
			position2108, tokenIndex2108 := position, tokenIndex
			{
				position2109 := position
				{
					position2110 := position
					if !_rules[rulehexdigit]() {
						goto l2108
					}
					if !_rules[rulehexdigit]() {
						goto l2108
					}
					if !_rules[rulehexdigit]() {
						goto l2108
					}
					if !_rules[rulehexdigit]() {
						goto l2108
					}
					add(rulePegText, position2110)
				}
				{
					add(ruleAction41, position)
				}
				add(rulehexWord, position2109)
			}
			return true
		l2108:
			position, tokenIndex = position2108, tokenIndex2108
			return false
		},
		/* 64 zeroHexWord <- <('0' hexWord)> */
		func() bool {
			position2112, tokenIndex2112 := position, tokenIndex
			{
				position2113 := position
				if buffer[position] != rune('0') {
					goto l2112
				}
				position++
				if !_rules[rulehexWord]() {
					goto l2112
				}
				add(rulezeroHexWord, position2113)
			}
			return true
		l2112:
			position, tokenIndex = position2112, tokenIndex2112
			return false
		},
		/* 65 nn_contents <- <('(' nn ')' Action42)> */
		func() bool {
			position2114, tokenIndex2114 := position, tokenIndex
			{
				position2115 := position
				if buffer[position] != rune('(') {
					goto l2114
				}
				position++
				if !_rules[rulenn]() {
					goto l2114
				}
				if buffer[position] != rune(')') {
					goto l2114
				}
				position++
				{
					add(ruleAction42, position)
				}
				add(rulenn_contents, position2115)
			}
			return true
		l2114:
			position, tokenIndex = position2114, tokenIndex2114
			return false
		},
		/* 66 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 67 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action43)> */
		nil,
		/* 68 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action44)> */
		nil,
		/* 69 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action45)> */
		nil,
		/* 70 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action46)> */
		nil,
		/* 71 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action47)> */
		nil,
		/* 72 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action48)> */
		nil,
		/* 73 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action49)> */
		nil,
		/* 74 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action50)> */
		nil,
		/* 75 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 76 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 77 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 (sep Copy8)? Action51)> */
		nil,
		/* 78 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 (sep Copy8)? Action52)> */
		nil,
		/* 79 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action53)> */
		nil,
		/* 80 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 (sep Copy8)? Action54)> */
		nil,
		/* 81 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 (sep Copy8)? Action55)> */
		nil,
		/* 82 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 (sep Copy8)? Action56)> */
		nil,
		/* 83 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 (sep Copy8)? Action57)> */
		nil,
		/* 84 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action58)> */
		nil,
		/* 85 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action59)> */
		nil,
		/* 86 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 (sep Copy8)? Action60)> */
		nil,
		/* 87 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 (sep Copy8)? Action61)> */
		nil,
		/* 88 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 89 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action62)> */
		nil,
		/* 90 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action63)> */
		nil,
		/* 91 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action64)> */
		nil,
		/* 92 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action65)> */
		nil,
		/* 93 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action66)> */
		nil,
		/* 94 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action67)> */
		nil,
		/* 95 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action68)> */
		nil,
		/* 96 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action69)> */
		nil,
		/* 97 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action70)> */
		nil,
		/* 98 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action71)> */
		nil,
		/* 99 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action72)> */
		nil,
		/* 100 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action73)> */
		nil,
		/* 101 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action74)> */
		nil,
		/* 102 EDSimple <- <(Retn / Reti / Rrd / Im0 / Im1 / Im2 / ((&('I' | 'O' | 'i' | 'o') BlitIO) | (&('R' | 'r') Rld) | (&('N' | 'n') Neg) | (&('C' | 'L' | 'c' | 'l') Blit)))> */
		nil,
		/* 103 Neg <- <(<(('n' / 'N') ('e' / 'E') ('g' / 'G'))> Action75)> */
		nil,
		/* 104 Retn <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N'))> Action76)> */
		nil,
		/* 105 Reti <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('i' / 'I'))> Action77)> */
		nil,
		/* 106 Rrd <- <(<(('r' / 'R') ('r' / 'R') ('d' / 'D'))> Action78)> */
		nil,
		/* 107 Rld <- <(<(('r' / 'R') ('l' / 'L') ('d' / 'D'))> Action79)> */
		nil,
		/* 108 Im0 <- <(<(('i' / 'I') ('m' / 'M') ' ' '0')> Action80)> */
		nil,
		/* 109 Im1 <- <(<(('i' / 'I') ('m' / 'M') ' ' '1')> Action81)> */
		nil,
		/* 110 Im2 <- <(<(('i' / 'I') ('m' / 'M') ' ' '2')> Action82)> */
		nil,
		/* 111 Blit <- <(Ldir / Ldi / Cpir / Cpi / Lddr / Ldd / Cpdr / Cpd)> */
		nil,
		/* 112 BlitIO <- <(Inir / Ini / Otir / Outi / Indr / Ind / Otdr / Outd)> */
		nil,
		/* 113 Ldi <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I'))> Action83)> */
		nil,
		/* 114 Cpi <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I'))> Action84)> */
		nil,
		/* 115 Ini <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I'))> Action85)> */
		nil,
		/* 116 Outi <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('i' / 'I'))> Action86)> */
		nil,
		/* 117 Ldd <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D'))> Action87)> */
		nil,
		/* 118 Cpd <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D'))> Action88)> */
		nil,
		/* 119 Ind <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D'))> Action89)> */
		nil,
		/* 120 Outd <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('d' / 'D'))> Action90)> */
		nil,
		/* 121 Ldir <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I') ('r' / 'R'))> Action91)> */
		nil,
		/* 122 Cpir <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I') ('r' / 'R'))> Action92)> */
		nil,
		/* 123 Inir <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I') ('r' / 'R'))> Action93)> */
		nil,
		/* 124 Otir <- <(<(('o' / 'O') ('t' / 'T') ('i' / 'I') ('r' / 'R'))> Action94)> */
		nil,
		/* 125 Lddr <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D') ('r' / 'R'))> Action95)> */
		nil,
		/* 126 Cpdr <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D') ('r' / 'R'))> Action96)> */
		nil,
		/* 127 Indr <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D') ('r' / 'R'))> Action97)> */
		nil,
		/* 128 Otdr <- <(<(('o' / 'O') ('t' / 'T') ('d' / 'D') ('r' / 'R'))> Action98)> */
		nil,
		/* 129 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 130 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action99)> */
		nil,
		/* 131 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action100)> */
		nil,
		/* 132 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action101)> */
		nil,
		/* 133 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action102)> */
		nil,
		/* 134 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action103)> */
		nil,
		/* 135 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action104)> */
		nil,
		/* 136 IO <- <(IN / OUT)> */
		nil,
		/* 137 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action105)> */
		nil,
		/* 138 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action106)> */
		nil,
		/* 139 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position2190, tokenIndex2190 := position, tokenIndex
			{
				position2191 := position
				{
					position2192, tokenIndex2192 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2193
					}
					position++
					{
						position2194, tokenIndex2194 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l2195
						}
						position++
						goto l2194
					l2195:
						position, tokenIndex = position2194, tokenIndex2194
						if buffer[position] != rune('C') {
							goto l2193
						}
						position++
					}
				l2194:
					if buffer[position] != rune(')') {
						goto l2193
					}
					position++
					goto l2192
				l2193:
					position, tokenIndex = position2192, tokenIndex2192
					if buffer[position] != rune('(') {
						goto l2190
					}
					position++
					if !_rules[rulen]() {
						goto l2190
					}
					if buffer[position] != rune(')') {
						goto l2190
					}
					position++
				}
			l2192:
				add(rulePort, position2191)
			}
			return true
		l2190:
			position, tokenIndex = position2190, tokenIndex2190
			return false
		},
		/* 140 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2196, tokenIndex2196 := position, tokenIndex
			{
				position2197 := position
				{
					position2198, tokenIndex2198 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2198
					}
					goto l2199
				l2198:
					position, tokenIndex = position2198, tokenIndex2198
				}
			l2199:
				if buffer[position] != rune(',') {
					goto l2196
				}
				position++
				{
					position2200, tokenIndex2200 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2200
					}
					goto l2201
				l2200:
					position, tokenIndex = position2200, tokenIndex2200
				}
			l2201:
				add(rulesep, position2197)
			}
			return true
		l2196:
			position, tokenIndex = position2196, tokenIndex2196
			return false
		},
		/* 141 ws <- <(' ' / '\t')+> */
		func() bool {
			position2202, tokenIndex2202 := position, tokenIndex
			{
				position2203 := position
				{
					position2206, tokenIndex2206 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2207
					}
					position++
					goto l2206
				l2207:
					position, tokenIndex = position2206, tokenIndex2206
					if buffer[position] != rune('\t') {
						goto l2202
					}
					position++
				}
			l2206:
			l2204:
				{
					position2205, tokenIndex2205 := position, tokenIndex
					{
						position2208, tokenIndex2208 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l2209
						}
						position++
						goto l2208
					l2209:
						position, tokenIndex = position2208, tokenIndex2208
						if buffer[position] != rune('\t') {
							goto l2205
						}
						position++
					}
				l2208:
					goto l2204
				l2205:
					position, tokenIndex = position2205, tokenIndex2205
				}
				add(rulews, position2203)
			}
			return true
		l2202:
			position, tokenIndex = position2202, tokenIndex2202
			return false
		},
		/* 142 A <- <('a' / 'A')> */
		nil,
		/* 143 F <- <('f' / 'F')> */
		nil,
		/* 144 B <- <('b' / 'B')> */
		nil,
		/* 145 C <- <('c' / 'C')> */
		nil,
		/* 146 D <- <('d' / 'D')> */
		nil,
		/* 147 E <- <('e' / 'E')> */
		nil,
		/* 148 H <- <('h' / 'H')> */
		nil,
		/* 149 L <- <('l' / 'L')> */
		nil,
		/* 150 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 151 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 152 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 153 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 154 I <- <('i' / 'I')> */
		nil,
		/* 155 R <- <('r' / 'R')> */
		nil,
		/* 156 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 157 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 158 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 159 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 160 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 161 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 162 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 163 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 164 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position2232, tokenIndex2232 := position, tokenIndex
			{
				position2233 := position
				{
					position2234, tokenIndex2234 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2235
					}
					position++
					goto l2234
				l2235:
					position, tokenIndex = position2234, tokenIndex2234
					{
						position2236, tokenIndex2236 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2237
						}
						position++
						goto l2236
					l2237:
						position, tokenIndex = position2236, tokenIndex2236
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2232
						}
						position++
					}
				l2236:
				}
			l2234:
				add(rulehexdigit, position2233)
			}
			return true
		l2232:
			position, tokenIndex = position2232, tokenIndex2232
			return false
		},
		/* 165 octaldigit <- <(<[0-7]> Action107)> */
		func() bool {
			position2238, tokenIndex2238 := position, tokenIndex
			{
				position2239 := position
				{
					position2240 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2238
					}
					position++
					add(rulePegText, position2240)
				}
				{
					add(ruleAction107, position)
				}
				add(ruleoctaldigit, position2239)
			}
			return true
		l2238:
			position, tokenIndex = position2238, tokenIndex2238
			return false
		},
		/* 166 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2242, tokenIndex2242 := position, tokenIndex
			{
				position2243 := position
				{
					position2244, tokenIndex2244 := position, tokenIndex
					{
						position2246 := position
						{
							position2247, tokenIndex2247 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2248
							}
							position++
							goto l2247
						l2248:
							position, tokenIndex = position2247, tokenIndex2247
							if buffer[position] != rune('N') {
								goto l2245
							}
							position++
						}
					l2247:
						{
							position2249, tokenIndex2249 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2250
							}
							position++
							goto l2249
						l2250:
							position, tokenIndex = position2249, tokenIndex2249
							if buffer[position] != rune('Z') {
								goto l2245
							}
							position++
						}
					l2249:
						{
							add(ruleAction108, position)
						}
						add(ruleFT_NZ, position2246)
					}
					goto l2244
				l2245:
					position, tokenIndex = position2244, tokenIndex2244
					{
						position2253 := position
						{
							position2254, tokenIndex2254 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2255
							}
							position++
							goto l2254
						l2255:
							position, tokenIndex = position2254, tokenIndex2254
							if buffer[position] != rune('P') {
								goto l2252
							}
							position++
						}
					l2254:
						{
							position2256, tokenIndex2256 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2257
							}
							position++
							goto l2256
						l2257:
							position, tokenIndex = position2256, tokenIndex2256
							if buffer[position] != rune('O') {
								goto l2252
							}
							position++
						}
					l2256:
						{
							add(ruleAction112, position)
						}
						add(ruleFT_PO, position2253)
					}
					goto l2244
				l2252:
					position, tokenIndex = position2244, tokenIndex2244
					{
						position2260 := position
						{
							position2261, tokenIndex2261 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2262
							}
							position++
							goto l2261
						l2262:
							position, tokenIndex = position2261, tokenIndex2261
							if buffer[position] != rune('P') {
								goto l2259
							}
							position++
						}
					l2261:
						{
							position2263, tokenIndex2263 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2264
							}
							position++
							goto l2263
						l2264:
							position, tokenIndex = position2263, tokenIndex2263
							if buffer[position] != rune('E') {
								goto l2259
							}
							position++
						}
					l2263:
						{
							add(ruleAction113, position)
						}
						add(ruleFT_PE, position2260)
					}
					goto l2244
				l2259:
					position, tokenIndex = position2244, tokenIndex2244
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2267 := position
								{
									position2268, tokenIndex2268 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2269
									}
									position++
									goto l2268
								l2269:
									position, tokenIndex = position2268, tokenIndex2268
									if buffer[position] != rune('M') {
										goto l2242
									}
									position++
								}
							l2268:
								{
									add(ruleAction115, position)
								}
								add(ruleFT_M, position2267)
							}
							break
						case 'P', 'p':
							{
								position2271 := position
								{
									position2272, tokenIndex2272 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2273
									}
									position++
									goto l2272
								l2273:
									position, tokenIndex = position2272, tokenIndex2272
									if buffer[position] != rune('P') {
										goto l2242
									}
									position++
								}
							l2272:
								{
									add(ruleAction114, position)
								}
								add(ruleFT_P, position2271)
							}
							break
						case 'C', 'c':
							{
								position2275 := position
								{
									position2276, tokenIndex2276 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2277
									}
									position++
									goto l2276
								l2277:
									position, tokenIndex = position2276, tokenIndex2276
									if buffer[position] != rune('C') {
										goto l2242
									}
									position++
								}
							l2276:
								{
									add(ruleAction111, position)
								}
								add(ruleFT_C, position2275)
							}
							break
						case 'N', 'n':
							{
								position2279 := position
								{
									position2280, tokenIndex2280 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2281
									}
									position++
									goto l2280
								l2281:
									position, tokenIndex = position2280, tokenIndex2280
									if buffer[position] != rune('N') {
										goto l2242
									}
									position++
								}
							l2280:
								{
									position2282, tokenIndex2282 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2283
									}
									position++
									goto l2282
								l2283:
									position, tokenIndex = position2282, tokenIndex2282
									if buffer[position] != rune('C') {
										goto l2242
									}
									position++
								}
							l2282:
								{
									add(ruleAction110, position)
								}
								add(ruleFT_NC, position2279)
							}
							break
						default:
							{
								position2285 := position
								{
									position2286, tokenIndex2286 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2287
									}
									position++
									goto l2286
								l2287:
									position, tokenIndex = position2286, tokenIndex2286
									if buffer[position] != rune('Z') {
										goto l2242
									}
									position++
								}
							l2286:
								{
									add(ruleAction109, position)
								}
								add(ruleFT_Z, position2285)
							}
							break
						}
					}

				}
			l2244:
				add(rulecc, position2243)
			}
			return true
		l2242:
			position, tokenIndex = position2242, tokenIndex2242
			return false
		},
		/* 167 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action108)> */
		nil,
		/* 168 FT_Z <- <(('z' / 'Z') Action109)> */
		nil,
		/* 169 FT_NC <- <(('n' / 'N') ('c' / 'C') Action110)> */
		nil,
		/* 170 FT_C <- <(('c' / 'C') Action111)> */
		nil,
		/* 171 FT_PO <- <(('p' / 'P') ('o' / 'O') Action112)> */
		nil,
		/* 172 FT_PE <- <(('p' / 'P') ('e' / 'E') Action113)> */
		nil,
		/* 173 FT_P <- <(('p' / 'P') Action114)> */
		nil,
		/* 174 FT_M <- <(('m' / 'M') Action115)> */
		nil,
		/* 176 Action0 <- <{ p.Emit() }> */
		nil,
		/* 177 Action1 <- <{ p.Org() }> */
		nil,
		/* 178 Action2 <- <{ p.DefByte() }> */
		nil,
		/* 179 Action3 <- <{ p.DefWord() }> */
		nil,
		/* 180 Action4 <- <{ p.DefSpace() }> */
		nil,
		/* 181 Action5 <- <{ p.LabelDefn(buffer[begin:end])}> */
		nil,
		nil,
		/* 183 Action6 <- <{ p.LD8() }> */
		nil,
		/* 184 Action7 <- <{ p.LD16() }> */
		nil,
		/* 185 Action8 <- <{ p.Push() }> */
		nil,
		/* 186 Action9 <- <{ p.Pop() }> */
		nil,
		/* 187 Action10 <- <{ p.Ex() }> */
		nil,
		/* 188 Action11 <- <{ p.Inc8() }> */
		nil,
		/* 189 Action12 <- <{ p.Inc8() }> */
		nil,
		/* 190 Action13 <- <{ p.Inc16() }> */
		nil,
		/* 191 Action14 <- <{ p.Dec8() }> */
		nil,
		/* 192 Action15 <- <{ p.Dec8() }> */
		nil,
		/* 193 Action16 <- <{ p.Dec16() }> */
		nil,
		/* 194 Action17 <- <{ p.Add16() }> */
		nil,
		/* 195 Action18 <- <{ p.Adc16() }> */
		nil,
		/* 196 Action19 <- <{ p.Sbc16() }> */
		nil,
		/* 197 Action20 <- <{ p.Dst8() }> */
		nil,
		/* 198 Action21 <- <{ p.Src8() }> */
		nil,
		/* 199 Action22 <- <{ p.Loc8() }> */
		nil,
		/* 200 Action23 <- <{ p.Copy8() }> */
		nil,
		/* 201 Action24 <- <{ p.Loc8() }> */
		nil,
		/* 202 Action25 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 203 Action26 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 204 Action27 <- <{ p.Dst16() }> */
		nil,
		/* 205 Action28 <- <{ p.Src16() }> */
		nil,
		/* 206 Action29 <- <{ p.Loc16() }> */
		nil,
		/* 207 Action30 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 208 Action31 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 209 Action32 <- <{ p.R16Contents() }> */
		nil,
		/* 210 Action33 <- <{ p.IR16Contents() }> */
		nil,
		/* 211 Action34 <- <{ p.DispDecimal(buffer[begin:end]) }> */
		nil,
		/* 212 Action35 <- <{ p.DispHex(buffer[begin:end]) }> */
		nil,
		/* 213 Action36 <- <{ p.Disp0xHex(buffer[begin:end]) }> */
		nil,
		/* 214 Action37 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 215 Action38 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 216 Action39 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 217 Action40 <- <{ p.NNLabel(buffer[begin:end]) }> */
		nil,
		/* 218 Action41 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 219 Action42 <- <{ p.NNContents() }> */
		nil,
		/* 220 Action43 <- <{ p.Accum("ADD") }> */
		nil,
		/* 221 Action44 <- <{ p.Accum("ADC") }> */
		nil,
		/* 222 Action45 <- <{ p.Accum("SUB") }> */
		nil,
		/* 223 Action46 <- <{ p.Accum("SBC") }> */
		nil,
		/* 224 Action47 <- <{ p.Accum("AND") }> */
		nil,
		/* 225 Action48 <- <{ p.Accum("XOR") }> */
		nil,
		/* 226 Action49 <- <{ p.Accum("OR") }> */
		nil,
		/* 227 Action50 <- <{ p.Accum("CP") }> */
		nil,
		/* 228 Action51 <- <{ p.Rot("RLC") }> */
		nil,
		/* 229 Action52 <- <{ p.Rot("RRC") }> */
		nil,
		/* 230 Action53 <- <{ p.Rot("RL") }> */
		nil,
		/* 231 Action54 <- <{ p.Rot("RR") }> */
		nil,
		/* 232 Action55 <- <{ p.Rot("SLA") }> */
		nil,
		/* 233 Action56 <- <{ p.Rot("SRA") }> */
		nil,
		/* 234 Action57 <- <{ p.Rot("SLL") }> */
		nil,
		/* 235 Action58 <- <{ p.Rot("SRL") }> */
		nil,
		/* 236 Action59 <- <{ p.Bit() }> */
		nil,
		/* 237 Action60 <- <{ p.Res() }> */
		nil,
		/* 238 Action61 <- <{ p.Set() }> */
		nil,
		/* 239 Action62 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 240 Action63 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 241 Action64 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 242 Action65 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 243 Action66 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 244 Action67 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 245 Action68 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 246 Action69 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 247 Action70 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 248 Action71 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 249 Action72 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 250 Action73 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 251 Action74 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 252 Action75 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 253 Action76 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 254 Action77 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 255 Action78 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 256 Action79 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 257 Action80 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 258 Action81 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 259 Action82 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 260 Action83 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 261 Action84 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 262 Action85 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 263 Action86 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 264 Action87 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 265 Action88 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 266 Action89 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 267 Action90 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 268 Action91 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 269 Action92 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 270 Action93 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 271 Action94 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 272 Action95 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 273 Action96 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 274 Action97 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 275 Action98 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 276 Action99 <- <{ p.Rst() }> */
		nil,
		/* 277 Action100 <- <{ p.Call() }> */
		nil,
		/* 278 Action101 <- <{ p.Ret() }> */
		nil,
		/* 279 Action102 <- <{ p.Jp() }> */
		nil,
		/* 280 Action103 <- <{ p.Jr() }> */
		nil,
		/* 281 Action104 <- <{ p.Djnz() }> */
		nil,
		/* 282 Action105 <- <{ p.In() }> */
		nil,
		/* 283 Action106 <- <{ p.Out() }> */
		nil,
		/* 284 Action107 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 285 Action108 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 286 Action109 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 287 Action110 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 288 Action111 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 289 Action112 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 290 Action113 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 291 Action114 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 292 Action115 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
