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
							{
								position10, tokenIndex10 := position, tokenIndex
								if !_rules[rulews]() {
									goto l10
								}
								goto l11
							l10:
								position, tokenIndex = position10, tokenIndex10
							}
						l11:
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
				l13:
					{
						position14, tokenIndex14 := position, tokenIndex
						if !_rules[rulews]() {
							goto l14
						}
						goto l13
					l14:
						position, tokenIndex = position14, tokenIndex14
					}
					{
						position15, tokenIndex15 := position, tokenIndex
						{
							position17 := position
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
													position26, tokenIndex26 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l27
													}
													position++
													goto l26
												l27:
													position, tokenIndex = position26, tokenIndex26
													if buffer[position] != rune('D') {
														goto l25
													}
													position++
												}
											l26:
												{
													position28, tokenIndex28 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l29
													}
													position++
													goto l28
												l29:
													position, tokenIndex = position28, tokenIndex28
													if buffer[position] != rune('E') {
														goto l25
													}
													position++
												}
											l28:
												{
													position30, tokenIndex30 := position, tokenIndex
													if buffer[position] != rune('f') {
														goto l31
													}
													position++
													goto l30
												l31:
													position, tokenIndex = position30, tokenIndex30
													if buffer[position] != rune('F') {
														goto l25
													}
													position++
												}
											l30:
												{
													position32, tokenIndex32 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l33
													}
													position++
													goto l32
												l33:
													position, tokenIndex = position32, tokenIndex32
													if buffer[position] != rune('B') {
														goto l25
													}
													position++
												}
											l32:
												goto l24
											l25:
												position, tokenIndex = position24, tokenIndex24
												{
													position34, tokenIndex34 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l35
													}
													position++
													goto l34
												l35:
													position, tokenIndex = position34, tokenIndex34
													if buffer[position] != rune('D') {
														goto l22
													}
													position++
												}
											l34:
												{
													position36, tokenIndex36 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l37
													}
													position++
													goto l36
												l37:
													position, tokenIndex = position36, tokenIndex36
													if buffer[position] != rune('B') {
														goto l22
													}
													position++
												}
											l36:
											}
										l24:
											if !_rules[rulews]() {
												goto l22
											}
											if !_rules[rulen]() {
												goto l22
											}
											{
												add(ruleAction2, position)
											}
											add(ruleDefb, position23)
										}
										goto l21
									l22:
										position, tokenIndex = position21, tokenIndex21
										{
											position40 := position
											{
												position41, tokenIndex41 := position, tokenIndex
												{
													position43, tokenIndex43 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l44
													}
													position++
													goto l43
												l44:
													position, tokenIndex = position43, tokenIndex43
													if buffer[position] != rune('D') {
														goto l42
													}
													position++
												}
											l43:
												{
													position45, tokenIndex45 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l46
													}
													position++
													goto l45
												l46:
													position, tokenIndex = position45, tokenIndex45
													if buffer[position] != rune('E') {
														goto l42
													}
													position++
												}
											l45:
												{
													position47, tokenIndex47 := position, tokenIndex
													if buffer[position] != rune('f') {
														goto l48
													}
													position++
													goto l47
												l48:
													position, tokenIndex = position47, tokenIndex47
													if buffer[position] != rune('F') {
														goto l42
													}
													position++
												}
											l47:
												{
													position49, tokenIndex49 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l50
													}
													position++
													goto l49
												l50:
													position, tokenIndex = position49, tokenIndex49
													if buffer[position] != rune('S') {
														goto l42
													}
													position++
												}
											l49:
												goto l41
											l42:
												position, tokenIndex = position41, tokenIndex41
												{
													position51, tokenIndex51 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l52
													}
													position++
													goto l51
												l52:
													position, tokenIndex = position51, tokenIndex51
													if buffer[position] != rune('D') {
														goto l39
													}
													position++
												}
											l51:
												{
													position53, tokenIndex53 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l54
													}
													position++
													goto l53
												l54:
													position, tokenIndex = position53, tokenIndex53
													if buffer[position] != rune('S') {
														goto l39
													}
													position++
												}
											l53:
											}
										l41:
											if !_rules[rulews]() {
												goto l39
											}
											if !_rules[rulen]() {
												goto l39
											}
											{
												add(ruleAction4, position)
											}
											add(ruleDefs, position40)
										}
										goto l21
									l39:
										position, tokenIndex = position21, tokenIndex21
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position57 := position
													{
														position58, tokenIndex58 := position, tokenIndex
														{
															position60, tokenIndex60 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l61
															}
															position++
															goto l60
														l61:
															position, tokenIndex = position60, tokenIndex60
															if buffer[position] != rune('D') {
																goto l59
															}
															position++
														}
													l60:
														{
															position62, tokenIndex62 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l63
															}
															position++
															goto l62
														l63:
															position, tokenIndex = position62, tokenIndex62
															if buffer[position] != rune('E') {
																goto l59
															}
															position++
														}
													l62:
														{
															position64, tokenIndex64 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l65
															}
															position++
															goto l64
														l65:
															position, tokenIndex = position64, tokenIndex64
															if buffer[position] != rune('F') {
																goto l59
															}
															position++
														}
													l64:
														{
															position66, tokenIndex66 := position, tokenIndex
															if buffer[position] != rune('w') {
																goto l67
															}
															position++
															goto l66
														l67:
															position, tokenIndex = position66, tokenIndex66
															if buffer[position] != rune('W') {
																goto l59
															}
															position++
														}
													l66:
														goto l58
													l59:
														position, tokenIndex = position58, tokenIndex58
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
																goto l19
															}
															position++
														}
													l68:
														{
															position70, tokenIndex70 := position, tokenIndex
															if buffer[position] != rune('w') {
																goto l71
															}
															position++
															goto l70
														l71:
															position, tokenIndex = position70, tokenIndex70
															if buffer[position] != rune('W') {
																goto l19
															}
															position++
														}
													l70:
													}
												l58:
													if !_rules[rulews]() {
														goto l19
													}
													if !_rules[rulenn]() {
														goto l19
													}
													{
														add(ruleAction3, position)
													}
													add(ruleDefw, position57)
												}
												break
											case 'O', 'o':
												{
													position73 := position
													{
														position74, tokenIndex74 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l75
														}
														position++
														goto l74
													l75:
														position, tokenIndex = position74, tokenIndex74
														if buffer[position] != rune('O') {
															goto l19
														}
														position++
													}
												l74:
													{
														position76, tokenIndex76 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l77
														}
														position++
														goto l76
													l77:
														position, tokenIndex = position76, tokenIndex76
														if buffer[position] != rune('R') {
															goto l19
														}
														position++
													}
												l76:
													{
														position78, tokenIndex78 := position, tokenIndex
														if buffer[position] != rune('g') {
															goto l79
														}
														position++
														goto l78
													l79:
														position, tokenIndex = position78, tokenIndex78
														if buffer[position] != rune('G') {
															goto l19
														}
														position++
													}
												l78:
													if !_rules[rulews]() {
														goto l19
													}
													if !_rules[rulenn]() {
														goto l19
													}
													{
														add(ruleAction1, position)
													}
													add(ruleOrg, position73)
												}
												break
											case 'a':
												{
													position81 := position
													if buffer[position] != rune('a') {
														goto l19
													}
													position++
													if buffer[position] != rune('s') {
														goto l19
													}
													position++
													if buffer[position] != rune('e') {
														goto l19
													}
													position++
													if buffer[position] != rune('g') {
														goto l19
													}
													position++
													add(ruleAseg, position81)
												}
												break
											default:
												{
													position82 := position
													{
														position83, tokenIndex83 := position, tokenIndex
														if buffer[position] != rune('.') {
															goto l83
														}
														position++
														goto l84
													l83:
														position, tokenIndex = position83, tokenIndex83
													}
												l84:
													if buffer[position] != rune('t') {
														goto l19
													}
													position++
													if buffer[position] != rune('i') {
														goto l19
													}
													position++
													if buffer[position] != rune('t') {
														goto l19
													}
													position++
													if buffer[position] != rune('l') {
														goto l19
													}
													position++
													if buffer[position] != rune('e') {
														goto l19
													}
													position++
													if !_rules[rulews]() {
														goto l19
													}
													if buffer[position] != rune('\'') {
														goto l19
													}
													position++
												l85:
													{
														position86, tokenIndex86 := position, tokenIndex
														{
															position87, tokenIndex87 := position, tokenIndex
															if buffer[position] != rune('\'') {
																goto l87
															}
															position++
															goto l86
														l87:
															position, tokenIndex = position87, tokenIndex87
														}
														if !matchDot() {
															goto l86
														}
														goto l85
													l86:
														position, tokenIndex = position86, tokenIndex86
													}
													if buffer[position] != rune('\'') {
														goto l19
													}
													position++
													add(ruleTitle, position82)
												}
												break
											}
										}

									}
								l21:
									add(ruleDirective, position20)
								}
								goto l18
							l19:
								position, tokenIndex = position18, tokenIndex18
								{
									position88 := position
									{
										position89, tokenIndex89 := position, tokenIndex
										{
											position91 := position
											{
												position92, tokenIndex92 := position, tokenIndex
												{
													position94 := position
													{
														position95, tokenIndex95 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l96
														}
														position++
														goto l95
													l96:
														position, tokenIndex = position95, tokenIndex95
														if buffer[position] != rune('P') {
															goto l93
														}
														position++
													}
												l95:
													{
														position97, tokenIndex97 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l98
														}
														position++
														goto l97
													l98:
														position, tokenIndex = position97, tokenIndex97
														if buffer[position] != rune('U') {
															goto l93
														}
														position++
													}
												l97:
													{
														position99, tokenIndex99 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l100
														}
														position++
														goto l99
													l100:
														position, tokenIndex = position99, tokenIndex99
														if buffer[position] != rune('S') {
															goto l93
														}
														position++
													}
												l99:
													{
														position101, tokenIndex101 := position, tokenIndex
														if buffer[position] != rune('h') {
															goto l102
														}
														position++
														goto l101
													l102:
														position, tokenIndex = position101, tokenIndex101
														if buffer[position] != rune('H') {
															goto l93
														}
														position++
													}
												l101:
													if !_rules[rulews]() {
														goto l93
													}
													if !_rules[ruleSrc16]() {
														goto l93
													}
													{
														add(ruleAction8, position)
													}
													add(rulePush, position94)
												}
												goto l92
											l93:
												position, tokenIndex = position92, tokenIndex92
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position105 := position
															{
																position106, tokenIndex106 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l107
																}
																position++
																goto l106
															l107:
																position, tokenIndex = position106, tokenIndex106
																if buffer[position] != rune('E') {
																	goto l90
																}
																position++
															}
														l106:
															{
																position108, tokenIndex108 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l109
																}
																position++
																goto l108
															l109:
																position, tokenIndex = position108, tokenIndex108
																if buffer[position] != rune('X') {
																	goto l90
																}
																position++
															}
														l108:
															if !_rules[rulews]() {
																goto l90
															}
															if !_rules[ruleDst16]() {
																goto l90
															}
															if !_rules[rulesep]() {
																goto l90
															}
															if !_rules[ruleSrc16]() {
																goto l90
															}
															{
																add(ruleAction10, position)
															}
															add(ruleEx, position105)
														}
														break
													case 'P', 'p':
														{
															position111 := position
															{
																position112, tokenIndex112 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l113
																}
																position++
																goto l112
															l113:
																position, tokenIndex = position112, tokenIndex112
																if buffer[position] != rune('P') {
																	goto l90
																}
																position++
															}
														l112:
															{
																position114, tokenIndex114 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l115
																}
																position++
																goto l114
															l115:
																position, tokenIndex = position114, tokenIndex114
																if buffer[position] != rune('O') {
																	goto l90
																}
																position++
															}
														l114:
															{
																position116, tokenIndex116 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l117
																}
																position++
																goto l116
															l117:
																position, tokenIndex = position116, tokenIndex116
																if buffer[position] != rune('P') {
																	goto l90
																}
																position++
															}
														l116:
															if !_rules[rulews]() {
																goto l90
															}
															if !_rules[ruleDst16]() {
																goto l90
															}
															{
																add(ruleAction9, position)
															}
															add(rulePop, position111)
														}
														break
													default:
														{
															position119 := position
															{
																position120, tokenIndex120 := position, tokenIndex
																{
																	position122 := position
																	{
																		position123, tokenIndex123 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l124
																		}
																		position++
																		goto l123
																	l124:
																		position, tokenIndex = position123, tokenIndex123
																		if buffer[position] != rune('L') {
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
																		add(ruleAction7, position)
																	}
																	add(ruleLoad16, position122)
																}
																goto l120
															l121:
																position, tokenIndex = position120, tokenIndex120
																{
																	position128 := position
																	{
																		position129, tokenIndex129 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l130
																		}
																		position++
																		goto l129
																	l130:
																		position, tokenIndex = position129, tokenIndex129
																		if buffer[position] != rune('L') {
																			goto l90
																		}
																		position++
																	}
																l129:
																	{
																		position131, tokenIndex131 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l132
																		}
																		position++
																		goto l131
																	l132:
																		position, tokenIndex = position131, tokenIndex131
																		if buffer[position] != rune('D') {
																			goto l90
																		}
																		position++
																	}
																l131:
																	if !_rules[rulews]() {
																		goto l90
																	}
																	{
																		position133 := position
																		{
																			position134, tokenIndex134 := position, tokenIndex
																			if !_rules[ruleReg8]() {
																				goto l135
																			}
																			goto l134
																		l135:
																			position, tokenIndex = position134, tokenIndex134
																			if !_rules[ruleReg16Contents]() {
																				goto l136
																			}
																			goto l134
																		l136:
																			position, tokenIndex = position134, tokenIndex134
																			if !_rules[rulenn_contents]() {
																				goto l90
																			}
																		}
																	l134:
																		{
																			add(ruleAction20, position)
																		}
																		add(ruleDst8, position133)
																	}
																	if !_rules[rulesep]() {
																		goto l90
																	}
																	if !_rules[ruleSrc8]() {
																		goto l90
																	}
																	{
																		add(ruleAction6, position)
																	}
																	add(ruleLoad8, position128)
																}
															}
														l120:
															add(ruleLoad, position119)
														}
														break
													}
												}

											}
										l92:
											add(ruleAssignment, position91)
										}
										goto l89
									l90:
										position, tokenIndex = position89, tokenIndex89
										{
											position140 := position
											{
												position141, tokenIndex141 := position, tokenIndex
												{
													position143 := position
													{
														position144, tokenIndex144 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l145
														}
														position++
														goto l144
													l145:
														position, tokenIndex = position144, tokenIndex144
														if buffer[position] != rune('I') {
															goto l142
														}
														position++
													}
												l144:
													{
														position146, tokenIndex146 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l147
														}
														position++
														goto l146
													l147:
														position, tokenIndex = position146, tokenIndex146
														if buffer[position] != rune('N') {
															goto l142
														}
														position++
													}
												l146:
													{
														position148, tokenIndex148 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l149
														}
														position++
														goto l148
													l149:
														position, tokenIndex = position148, tokenIndex148
														if buffer[position] != rune('C') {
															goto l142
														}
														position++
													}
												l148:
													if !_rules[rulews]() {
														goto l142
													}
													if !_rules[ruleILoc8]() {
														goto l142
													}
													{
														add(ruleAction11, position)
													}
													add(ruleInc16Indexed8, position143)
												}
												goto l141
											l142:
												position, tokenIndex = position141, tokenIndex141
												{
													position152 := position
													{
														position153, tokenIndex153 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l154
														}
														position++
														goto l153
													l154:
														position, tokenIndex = position153, tokenIndex153
														if buffer[position] != rune('I') {
															goto l151
														}
														position++
													}
												l153:
													{
														position155, tokenIndex155 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l156
														}
														position++
														goto l155
													l156:
														position, tokenIndex = position155, tokenIndex155
														if buffer[position] != rune('N') {
															goto l151
														}
														position++
													}
												l155:
													{
														position157, tokenIndex157 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l158
														}
														position++
														goto l157
													l158:
														position, tokenIndex = position157, tokenIndex157
														if buffer[position] != rune('C') {
															goto l151
														}
														position++
													}
												l157:
													if !_rules[rulews]() {
														goto l151
													}
													if !_rules[ruleLoc16]() {
														goto l151
													}
													{
														add(ruleAction13, position)
													}
													add(ruleInc16, position152)
												}
												goto l141
											l151:
												position, tokenIndex = position141, tokenIndex141
												{
													position160 := position
													{
														position161, tokenIndex161 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l162
														}
														position++
														goto l161
													l162:
														position, tokenIndex = position161, tokenIndex161
														if buffer[position] != rune('I') {
															goto l139
														}
														position++
													}
												l161:
													{
														position163, tokenIndex163 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l164
														}
														position++
														goto l163
													l164:
														position, tokenIndex = position163, tokenIndex163
														if buffer[position] != rune('N') {
															goto l139
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
															goto l139
														}
														position++
													}
												l165:
													if !_rules[rulews]() {
														goto l139
													}
													if !_rules[ruleLoc8]() {
														goto l139
													}
													{
														add(ruleAction12, position)
													}
													add(ruleInc8, position160)
												}
											}
										l141:
											add(ruleInc, position140)
										}
										goto l89
									l139:
										position, tokenIndex = position89, tokenIndex89
										{
											position169 := position
											{
												position170, tokenIndex170 := position, tokenIndex
												{
													position172 := position
													{
														position173, tokenIndex173 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l174
														}
														position++
														goto l173
													l174:
														position, tokenIndex = position173, tokenIndex173
														if buffer[position] != rune('D') {
															goto l171
														}
														position++
													}
												l173:
													{
														position175, tokenIndex175 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l176
														}
														position++
														goto l175
													l176:
														position, tokenIndex = position175, tokenIndex175
														if buffer[position] != rune('E') {
															goto l171
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
															goto l171
														}
														position++
													}
												l177:
													if !_rules[rulews]() {
														goto l171
													}
													if !_rules[ruleILoc8]() {
														goto l171
													}
													{
														add(ruleAction14, position)
													}
													add(ruleDec16Indexed8, position172)
												}
												goto l170
											l171:
												position, tokenIndex = position170, tokenIndex170
												{
													position181 := position
													{
														position182, tokenIndex182 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l183
														}
														position++
														goto l182
													l183:
														position, tokenIndex = position182, tokenIndex182
														if buffer[position] != rune('D') {
															goto l180
														}
														position++
													}
												l182:
													{
														position184, tokenIndex184 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l185
														}
														position++
														goto l184
													l185:
														position, tokenIndex = position184, tokenIndex184
														if buffer[position] != rune('E') {
															goto l180
														}
														position++
													}
												l184:
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
															goto l180
														}
														position++
													}
												l186:
													if !_rules[rulews]() {
														goto l180
													}
													if !_rules[ruleLoc16]() {
														goto l180
													}
													{
														add(ruleAction16, position)
													}
													add(ruleDec16, position181)
												}
												goto l170
											l180:
												position, tokenIndex = position170, tokenIndex170
												{
													position189 := position
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
															goto l168
														}
														position++
													}
												l190:
													{
														position192, tokenIndex192 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l193
														}
														position++
														goto l192
													l193:
														position, tokenIndex = position192, tokenIndex192
														if buffer[position] != rune('E') {
															goto l168
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
															goto l168
														}
														position++
													}
												l194:
													if !_rules[rulews]() {
														goto l168
													}
													if !_rules[ruleLoc8]() {
														goto l168
													}
													{
														add(ruleAction15, position)
													}
													add(ruleDec8, position189)
												}
											}
										l170:
											add(ruleDec, position169)
										}
										goto l89
									l168:
										position, tokenIndex = position89, tokenIndex89
										{
											position198 := position
											{
												position199, tokenIndex199 := position, tokenIndex
												{
													position201 := position
													{
														position202, tokenIndex202 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l203
														}
														position++
														goto l202
													l203:
														position, tokenIndex = position202, tokenIndex202
														if buffer[position] != rune('A') {
															goto l200
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
															goto l200
														}
														position++
													}
												l204:
													{
														position206, tokenIndex206 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l207
														}
														position++
														goto l206
													l207:
														position, tokenIndex = position206, tokenIndex206
														if buffer[position] != rune('D') {
															goto l200
														}
														position++
													}
												l206:
													if !_rules[rulews]() {
														goto l200
													}
													if !_rules[ruleDst16]() {
														goto l200
													}
													if !_rules[rulesep]() {
														goto l200
													}
													if !_rules[ruleSrc16]() {
														goto l200
													}
													{
														add(ruleAction17, position)
													}
													add(ruleAdd16, position201)
												}
												goto l199
											l200:
												position, tokenIndex = position199, tokenIndex199
												{
													position210 := position
													{
														position211, tokenIndex211 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l212
														}
														position++
														goto l211
													l212:
														position, tokenIndex = position211, tokenIndex211
														if buffer[position] != rune('A') {
															goto l209
														}
														position++
													}
												l211:
													{
														position213, tokenIndex213 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l214
														}
														position++
														goto l213
													l214:
														position, tokenIndex = position213, tokenIndex213
														if buffer[position] != rune('D') {
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
													if !_rules[ruleDst16]() {
														goto l209
													}
													if !_rules[rulesep]() {
														goto l209
													}
													if !_rules[ruleSrc16]() {
														goto l209
													}
													{
														add(ruleAction18, position)
													}
													add(ruleAdc16, position210)
												}
												goto l199
											l209:
												position, tokenIndex = position199, tokenIndex199
												{
													position218 := position
													{
														position219, tokenIndex219 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l220
														}
														position++
														goto l219
													l220:
														position, tokenIndex = position219, tokenIndex219
														if buffer[position] != rune('S') {
															goto l197
														}
														position++
													}
												l219:
													{
														position221, tokenIndex221 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l222
														}
														position++
														goto l221
													l222:
														position, tokenIndex = position221, tokenIndex221
														if buffer[position] != rune('B') {
															goto l197
														}
														position++
													}
												l221:
													{
														position223, tokenIndex223 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l224
														}
														position++
														goto l223
													l224:
														position, tokenIndex = position223, tokenIndex223
														if buffer[position] != rune('C') {
															goto l197
														}
														position++
													}
												l223:
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
														add(ruleAction19, position)
													}
													add(ruleSbc16, position218)
												}
											}
										l199:
											add(ruleAlu16, position198)
										}
										goto l89
									l197:
										position, tokenIndex = position89, tokenIndex89
										{
											position227 := position
											{
												position228, tokenIndex228 := position, tokenIndex
												{
													position230 := position
													{
														position231, tokenIndex231 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l232
														}
														position++
														goto l231
													l232:
														position, tokenIndex = position231, tokenIndex231
														if buffer[position] != rune('A') {
															goto l229
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
															goto l229
														}
														position++
													}
												l233:
													{
														position235, tokenIndex235 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l236
														}
														position++
														goto l235
													l236:
														position, tokenIndex = position235, tokenIndex235
														if buffer[position] != rune('D') {
															goto l229
														}
														position++
													}
												l235:
													if !_rules[rulews]() {
														goto l229
													}
													{
														position237, tokenIndex237 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l238
														}
														position++
														goto l237
													l238:
														position, tokenIndex = position237, tokenIndex237
														if buffer[position] != rune('A') {
															goto l229
														}
														position++
													}
												l237:
													if !_rules[rulesep]() {
														goto l229
													}
													if !_rules[ruleSrc8]() {
														goto l229
													}
													{
														add(ruleAction43, position)
													}
													add(ruleAdd, position230)
												}
												goto l228
											l229:
												position, tokenIndex = position228, tokenIndex228
												{
													position241 := position
													{
														position242, tokenIndex242 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l243
														}
														position++
														goto l242
													l243:
														position, tokenIndex = position242, tokenIndex242
														if buffer[position] != rune('A') {
															goto l240
														}
														position++
													}
												l242:
													{
														position244, tokenIndex244 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l245
														}
														position++
														goto l244
													l245:
														position, tokenIndex = position244, tokenIndex244
														if buffer[position] != rune('D') {
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
													{
														position248, tokenIndex248 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l249
														}
														position++
														goto l248
													l249:
														position, tokenIndex = position248, tokenIndex248
														if buffer[position] != rune('A') {
															goto l240
														}
														position++
													}
												l248:
													if !_rules[rulesep]() {
														goto l240
													}
													if !_rules[ruleSrc8]() {
														goto l240
													}
													{
														add(ruleAction44, position)
													}
													add(ruleAdc, position241)
												}
												goto l228
											l240:
												position, tokenIndex = position228, tokenIndex228
												{
													position252 := position
													{
														position253, tokenIndex253 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l254
														}
														position++
														goto l253
													l254:
														position, tokenIndex = position253, tokenIndex253
														if buffer[position] != rune('S') {
															goto l251
														}
														position++
													}
												l253:
													{
														position255, tokenIndex255 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l256
														}
														position++
														goto l255
													l256:
														position, tokenIndex = position255, tokenIndex255
														if buffer[position] != rune('U') {
															goto l251
														}
														position++
													}
												l255:
													{
														position257, tokenIndex257 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l258
														}
														position++
														goto l257
													l258:
														position, tokenIndex = position257, tokenIndex257
														if buffer[position] != rune('B') {
															goto l251
														}
														position++
													}
												l257:
													if !_rules[rulews]() {
														goto l251
													}
													if !_rules[ruleSrc8]() {
														goto l251
													}
													{
														add(ruleAction45, position)
													}
													add(ruleSub, position252)
												}
												goto l228
											l251:
												position, tokenIndex = position228, tokenIndex228
												{
													switch buffer[position] {
													case 'C', 'c':
														{
															position261 := position
															{
																position262, tokenIndex262 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l263
																}
																position++
																goto l262
															l263:
																position, tokenIndex = position262, tokenIndex262
																if buffer[position] != rune('C') {
																	goto l226
																}
																position++
															}
														l262:
															{
																position264, tokenIndex264 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l265
																}
																position++
																goto l264
															l265:
																position, tokenIndex = position264, tokenIndex264
																if buffer[position] != rune('P') {
																	goto l226
																}
																position++
															}
														l264:
															if !_rules[rulews]() {
																goto l226
															}
															if !_rules[ruleSrc8]() {
																goto l226
															}
															{
																add(ruleAction50, position)
															}
															add(ruleCp, position261)
														}
														break
													case 'O', 'o':
														{
															position267 := position
															{
																position268, tokenIndex268 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l269
																}
																position++
																goto l268
															l269:
																position, tokenIndex = position268, tokenIndex268
																if buffer[position] != rune('O') {
																	goto l226
																}
																position++
															}
														l268:
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
																	goto l226
																}
																position++
															}
														l270:
															if !_rules[rulews]() {
																goto l226
															}
															if !_rules[ruleSrc8]() {
																goto l226
															}
															{
																add(ruleAction49, position)
															}
															add(ruleOr, position267)
														}
														break
													case 'X', 'x':
														{
															position273 := position
															{
																position274, tokenIndex274 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l275
																}
																position++
																goto l274
															l275:
																position, tokenIndex = position274, tokenIndex274
																if buffer[position] != rune('X') {
																	goto l226
																}
																position++
															}
														l274:
															{
																position276, tokenIndex276 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l277
																}
																position++
																goto l276
															l277:
																position, tokenIndex = position276, tokenIndex276
																if buffer[position] != rune('O') {
																	goto l226
																}
																position++
															}
														l276:
															{
																position278, tokenIndex278 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l279
																}
																position++
																goto l278
															l279:
																position, tokenIndex = position278, tokenIndex278
																if buffer[position] != rune('R') {
																	goto l226
																}
																position++
															}
														l278:
															if !_rules[rulews]() {
																goto l226
															}
															if !_rules[ruleSrc8]() {
																goto l226
															}
															{
																add(ruleAction48, position)
															}
															add(ruleXor, position273)
														}
														break
													case 'A', 'a':
														{
															position281 := position
															{
																position282, tokenIndex282 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l283
																}
																position++
																goto l282
															l283:
																position, tokenIndex = position282, tokenIndex282
																if buffer[position] != rune('A') {
																	goto l226
																}
																position++
															}
														l282:
															{
																position284, tokenIndex284 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l285
																}
																position++
																goto l284
															l285:
																position, tokenIndex = position284, tokenIndex284
																if buffer[position] != rune('N') {
																	goto l226
																}
																position++
															}
														l284:
															{
																position286, tokenIndex286 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l287
																}
																position++
																goto l286
															l287:
																position, tokenIndex = position286, tokenIndex286
																if buffer[position] != rune('D') {
																	goto l226
																}
																position++
															}
														l286:
															if !_rules[rulews]() {
																goto l226
															}
															if !_rules[ruleSrc8]() {
																goto l226
															}
															{
																add(ruleAction47, position)
															}
															add(ruleAnd, position281)
														}
														break
													default:
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
																	goto l226
																}
																position++
															}
														l290:
															{
																position292, tokenIndex292 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l293
																}
																position++
																goto l292
															l293:
																position, tokenIndex = position292, tokenIndex292
																if buffer[position] != rune('B') {
																	goto l226
																}
																position++
															}
														l292:
															{
																position294, tokenIndex294 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l295
																}
																position++
																goto l294
															l295:
																position, tokenIndex = position294, tokenIndex294
																if buffer[position] != rune('C') {
																	goto l226
																}
																position++
															}
														l294:
															if !_rules[rulews]() {
																goto l226
															}
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
																	goto l226
																}
																position++
															}
														l296:
															if !_rules[rulesep]() {
																goto l226
															}
															if !_rules[ruleSrc8]() {
																goto l226
															}
															{
																add(ruleAction46, position)
															}
															add(ruleSbc, position289)
														}
														break
													}
												}

											}
										l228:
											add(ruleAlu, position227)
										}
										goto l89
									l226:
										position, tokenIndex = position89, tokenIndex89
										{
											position300 := position
											{
												position301, tokenIndex301 := position, tokenIndex
												{
													position303 := position
													{
														position304, tokenIndex304 := position, tokenIndex
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
																	goto l305
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
																	goto l305
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
																	goto l305
																}
																position++
															}
														l311:
															if !_rules[rulews]() {
																goto l305
															}
															if !_rules[ruleLoc8]() {
																goto l305
															}
															{
																position313, tokenIndex313 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l313
																}
																if !_rules[ruleCopy8]() {
																	goto l313
																}
																goto l314
															l313:
																position, tokenIndex = position313, tokenIndex313
															}
														l314:
															{
																add(ruleAction51, position)
															}
															add(ruleRlc, position306)
														}
														goto l304
													l305:
														position, tokenIndex = position304, tokenIndex304
														{
															position317 := position
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
																	goto l316
																}
																position++
															}
														l318:
															{
																position320, tokenIndex320 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l321
																}
																position++
																goto l320
															l321:
																position, tokenIndex = position320, tokenIndex320
																if buffer[position] != rune('R') {
																	goto l316
																}
																position++
															}
														l320:
															{
																position322, tokenIndex322 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l323
																}
																position++
																goto l322
															l323:
																position, tokenIndex = position322, tokenIndex322
																if buffer[position] != rune('C') {
																	goto l316
																}
																position++
															}
														l322:
															if !_rules[rulews]() {
																goto l316
															}
															if !_rules[ruleLoc8]() {
																goto l316
															}
															{
																position324, tokenIndex324 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l324
																}
																if !_rules[ruleCopy8]() {
																	goto l324
																}
																goto l325
															l324:
																position, tokenIndex = position324, tokenIndex324
															}
														l325:
															{
																add(ruleAction52, position)
															}
															add(ruleRrc, position317)
														}
														goto l304
													l316:
														position, tokenIndex = position304, tokenIndex304
														{
															position328 := position
															{
																position329, tokenIndex329 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l330
																}
																position++
																goto l329
															l330:
																position, tokenIndex = position329, tokenIndex329
																if buffer[position] != rune('R') {
																	goto l327
																}
																position++
															}
														l329:
															{
																position331, tokenIndex331 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l332
																}
																position++
																goto l331
															l332:
																position, tokenIndex = position331, tokenIndex331
																if buffer[position] != rune('L') {
																	goto l327
																}
																position++
															}
														l331:
															if !_rules[rulews]() {
																goto l327
															}
															if !_rules[ruleLoc8]() {
																goto l327
															}
															{
																position333, tokenIndex333 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l333
																}
																if !_rules[ruleCopy8]() {
																	goto l333
																}
																goto l334
															l333:
																position, tokenIndex = position333, tokenIndex333
															}
														l334:
															{
																add(ruleAction53, position)
															}
															add(ruleRl, position328)
														}
														goto l304
													l327:
														position, tokenIndex = position304, tokenIndex304
														{
															position337 := position
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
																	goto l336
																}
																position++
															}
														l338:
															{
																position340, tokenIndex340 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l341
																}
																position++
																goto l340
															l341:
																position, tokenIndex = position340, tokenIndex340
																if buffer[position] != rune('R') {
																	goto l336
																}
																position++
															}
														l340:
															if !_rules[rulews]() {
																goto l336
															}
															if !_rules[ruleLoc8]() {
																goto l336
															}
															{
																position342, tokenIndex342 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l342
																}
																if !_rules[ruleCopy8]() {
																	goto l342
																}
																goto l343
															l342:
																position, tokenIndex = position342, tokenIndex342
															}
														l343:
															{
																add(ruleAction54, position)
															}
															add(ruleRr, position337)
														}
														goto l304
													l336:
														position, tokenIndex = position304, tokenIndex304
														{
															position346 := position
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
																	goto l345
																}
																position++
															}
														l347:
															{
																position349, tokenIndex349 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l350
																}
																position++
																goto l349
															l350:
																position, tokenIndex = position349, tokenIndex349
																if buffer[position] != rune('L') {
																	goto l345
																}
																position++
															}
														l349:
															{
																position351, tokenIndex351 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l352
																}
																position++
																goto l351
															l352:
																position, tokenIndex = position351, tokenIndex351
																if buffer[position] != rune('A') {
																	goto l345
																}
																position++
															}
														l351:
															if !_rules[rulews]() {
																goto l345
															}
															if !_rules[ruleLoc8]() {
																goto l345
															}
															{
																position353, tokenIndex353 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l353
																}
																if !_rules[ruleCopy8]() {
																	goto l353
																}
																goto l354
															l353:
																position, tokenIndex = position353, tokenIndex353
															}
														l354:
															{
																add(ruleAction55, position)
															}
															add(ruleSla, position346)
														}
														goto l304
													l345:
														position, tokenIndex = position304, tokenIndex304
														{
															position357 := position
															{
																position358, tokenIndex358 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l359
																}
																position++
																goto l358
															l359:
																position, tokenIndex = position358, tokenIndex358
																if buffer[position] != rune('S') {
																	goto l356
																}
																position++
															}
														l358:
															{
																position360, tokenIndex360 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l361
																}
																position++
																goto l360
															l361:
																position, tokenIndex = position360, tokenIndex360
																if buffer[position] != rune('R') {
																	goto l356
																}
																position++
															}
														l360:
															{
																position362, tokenIndex362 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l363
																}
																position++
																goto l362
															l363:
																position, tokenIndex = position362, tokenIndex362
																if buffer[position] != rune('A') {
																	goto l356
																}
																position++
															}
														l362:
															if !_rules[rulews]() {
																goto l356
															}
															if !_rules[ruleLoc8]() {
																goto l356
															}
															{
																position364, tokenIndex364 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l364
																}
																if !_rules[ruleCopy8]() {
																	goto l364
																}
																goto l365
															l364:
																position, tokenIndex = position364, tokenIndex364
															}
														l365:
															{
																add(ruleAction56, position)
															}
															add(ruleSra, position357)
														}
														goto l304
													l356:
														position, tokenIndex = position304, tokenIndex304
														{
															position368 := position
															{
																position369, tokenIndex369 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l370
																}
																position++
																goto l369
															l370:
																position, tokenIndex = position369, tokenIndex369
																if buffer[position] != rune('S') {
																	goto l367
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
																	goto l367
																}
																position++
															}
														l371:
															{
																position373, tokenIndex373 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l374
																}
																position++
																goto l373
															l374:
																position, tokenIndex = position373, tokenIndex373
																if buffer[position] != rune('L') {
																	goto l367
																}
																position++
															}
														l373:
															if !_rules[rulews]() {
																goto l367
															}
															if !_rules[ruleLoc8]() {
																goto l367
															}
															{
																position375, tokenIndex375 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l375
																}
																if !_rules[ruleCopy8]() {
																	goto l375
																}
																goto l376
															l375:
																position, tokenIndex = position375, tokenIndex375
															}
														l376:
															{
																add(ruleAction57, position)
															}
															add(ruleSll, position368)
														}
														goto l304
													l367:
														position, tokenIndex = position304, tokenIndex304
														{
															position378 := position
															{
																position379, tokenIndex379 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l380
																}
																position++
																goto l379
															l380:
																position, tokenIndex = position379, tokenIndex379
																if buffer[position] != rune('S') {
																	goto l302
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
																	goto l302
																}
																position++
															}
														l381:
															{
																position383, tokenIndex383 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l384
																}
																position++
																goto l383
															l384:
																position, tokenIndex = position383, tokenIndex383
																if buffer[position] != rune('L') {
																	goto l302
																}
																position++
															}
														l383:
															if !_rules[rulews]() {
																goto l302
															}
															if !_rules[ruleLoc8]() {
																goto l302
															}
															{
																position385, tokenIndex385 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l385
																}
																if !_rules[ruleCopy8]() {
																	goto l385
																}
																goto l386
															l385:
																position, tokenIndex = position385, tokenIndex385
															}
														l386:
															{
																add(ruleAction58, position)
															}
															add(ruleSrl, position378)
														}
													}
												l304:
													add(ruleRot, position303)
												}
												goto l301
											l302:
												position, tokenIndex = position301, tokenIndex301
												{
													switch buffer[position] {
													case 'S', 's':
														{
															position389 := position
															{
																position390, tokenIndex390 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l391
																}
																position++
																goto l390
															l391:
																position, tokenIndex = position390, tokenIndex390
																if buffer[position] != rune('S') {
																	goto l299
																}
																position++
															}
														l390:
															{
																position392, tokenIndex392 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l393
																}
																position++
																goto l392
															l393:
																position, tokenIndex = position392, tokenIndex392
																if buffer[position] != rune('E') {
																	goto l299
																}
																position++
															}
														l392:
															{
																position394, tokenIndex394 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l395
																}
																position++
																goto l394
															l395:
																position, tokenIndex = position394, tokenIndex394
																if buffer[position] != rune('T') {
																	goto l299
																}
																position++
															}
														l394:
															if !_rules[rulews]() {
																goto l299
															}
															if !_rules[ruleoctaldigit]() {
																goto l299
															}
															if !_rules[rulesep]() {
																goto l299
															}
															if !_rules[ruleLoc8]() {
																goto l299
															}
															{
																position396, tokenIndex396 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l396
																}
																if !_rules[ruleCopy8]() {
																	goto l396
																}
																goto l397
															l396:
																position, tokenIndex = position396, tokenIndex396
															}
														l397:
															{
																add(ruleAction61, position)
															}
															add(ruleSet, position389)
														}
														break
													case 'R', 'r':
														{
															position399 := position
															{
																position400, tokenIndex400 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l401
																}
																position++
																goto l400
															l401:
																position, tokenIndex = position400, tokenIndex400
																if buffer[position] != rune('R') {
																	goto l299
																}
																position++
															}
														l400:
															{
																position402, tokenIndex402 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l403
																}
																position++
																goto l402
															l403:
																position, tokenIndex = position402, tokenIndex402
																if buffer[position] != rune('E') {
																	goto l299
																}
																position++
															}
														l402:
															{
																position404, tokenIndex404 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l405
																}
																position++
																goto l404
															l405:
																position, tokenIndex = position404, tokenIndex404
																if buffer[position] != rune('S') {
																	goto l299
																}
																position++
															}
														l404:
															if !_rules[rulews]() {
																goto l299
															}
															if !_rules[ruleoctaldigit]() {
																goto l299
															}
															if !_rules[rulesep]() {
																goto l299
															}
															if !_rules[ruleLoc8]() {
																goto l299
															}
															{
																position406, tokenIndex406 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l406
																}
																if !_rules[ruleCopy8]() {
																	goto l406
																}
																goto l407
															l406:
																position, tokenIndex = position406, tokenIndex406
															}
														l407:
															{
																add(ruleAction60, position)
															}
															add(ruleRes, position399)
														}
														break
													default:
														{
															position409 := position
															{
																position410, tokenIndex410 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l411
																}
																position++
																goto l410
															l411:
																position, tokenIndex = position410, tokenIndex410
																if buffer[position] != rune('B') {
																	goto l299
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
																	goto l299
																}
																position++
															}
														l412:
															{
																position414, tokenIndex414 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l415
																}
																position++
																goto l414
															l415:
																position, tokenIndex = position414, tokenIndex414
																if buffer[position] != rune('T') {
																	goto l299
																}
																position++
															}
														l414:
															if !_rules[rulews]() {
																goto l299
															}
															if !_rules[ruleoctaldigit]() {
																goto l299
															}
															if !_rules[rulesep]() {
																goto l299
															}
															if !_rules[ruleLoc8]() {
																goto l299
															}
															{
																add(ruleAction59, position)
															}
															add(ruleBit, position409)
														}
														break
													}
												}

											}
										l301:
											add(ruleBitOp, position300)
										}
										goto l89
									l299:
										position, tokenIndex = position89, tokenIndex89
										{
											position418 := position
											{
												position419, tokenIndex419 := position, tokenIndex
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
															if buffer[position] != rune('n') {
																goto l430
															}
															position++
															goto l429
														l430:
															position, tokenIndex = position429, tokenIndex429
															if buffer[position] != rune('N') {
																goto l420
															}
															position++
														}
													l429:
														add(rulePegText, position422)
													}
													{
														add(ruleAction76, position)
													}
													add(ruleRetn, position421)
												}
												goto l419
											l420:
												position, tokenIndex = position419, tokenIndex419
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
															if buffer[position] != rune('e') {
																goto l438
															}
															position++
															goto l437
														l438:
															position, tokenIndex = position437, tokenIndex437
															if buffer[position] != rune('E') {
																goto l432
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
																goto l432
															}
															position++
														}
													l439:
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
																goto l432
															}
															position++
														}
													l441:
														add(rulePegText, position434)
													}
													{
														add(ruleAction77, position)
													}
													add(ruleReti, position433)
												}
												goto l419
											l432:
												position, tokenIndex = position419, tokenIndex419
												{
													position445 := position
													{
														position446 := position
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
																goto l444
															}
															position++
														}
													l447:
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
																goto l444
															}
															position++
														}
													l449:
														{
															position451, tokenIndex451 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l452
															}
															position++
															goto l451
														l452:
															position, tokenIndex = position451, tokenIndex451
															if buffer[position] != rune('D') {
																goto l444
															}
															position++
														}
													l451:
														add(rulePegText, position446)
													}
													{
														add(ruleAction78, position)
													}
													add(ruleRrd, position445)
												}
												goto l419
											l444:
												position, tokenIndex = position419, tokenIndex419
												{
													position455 := position
													{
														position456 := position
														{
															position457, tokenIndex457 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l458
															}
															position++
															goto l457
														l458:
															position, tokenIndex = position457, tokenIndex457
															if buffer[position] != rune('I') {
																goto l454
															}
															position++
														}
													l457:
														{
															position459, tokenIndex459 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l460
															}
															position++
															goto l459
														l460:
															position, tokenIndex = position459, tokenIndex459
															if buffer[position] != rune('M') {
																goto l454
															}
															position++
														}
													l459:
														if buffer[position] != rune(' ') {
															goto l454
														}
														position++
														if buffer[position] != rune('0') {
															goto l454
														}
														position++
														add(rulePegText, position456)
													}
													{
														add(ruleAction80, position)
													}
													add(ruleIm0, position455)
												}
												goto l419
											l454:
												position, tokenIndex = position419, tokenIndex419
												{
													position463 := position
													{
														position464 := position
														{
															position465, tokenIndex465 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l466
															}
															position++
															goto l465
														l466:
															position, tokenIndex = position465, tokenIndex465
															if buffer[position] != rune('I') {
																goto l462
															}
															position++
														}
													l465:
														{
															position467, tokenIndex467 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l468
															}
															position++
															goto l467
														l468:
															position, tokenIndex = position467, tokenIndex467
															if buffer[position] != rune('M') {
																goto l462
															}
															position++
														}
													l467:
														if buffer[position] != rune(' ') {
															goto l462
														}
														position++
														if buffer[position] != rune('1') {
															goto l462
														}
														position++
														add(rulePegText, position464)
													}
													{
														add(ruleAction81, position)
													}
													add(ruleIm1, position463)
												}
												goto l419
											l462:
												position, tokenIndex = position419, tokenIndex419
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
															if buffer[position] != rune('m') {
																goto l476
															}
															position++
															goto l475
														l476:
															position, tokenIndex = position475, tokenIndex475
															if buffer[position] != rune('M') {
																goto l470
															}
															position++
														}
													l475:
														if buffer[position] != rune(' ') {
															goto l470
														}
														position++
														if buffer[position] != rune('2') {
															goto l470
														}
														position++
														add(rulePegText, position472)
													}
													{
														add(ruleAction82, position)
													}
													add(ruleIm2, position471)
												}
												goto l419
											l470:
												position, tokenIndex = position419, tokenIndex419
												{
													switch buffer[position] {
													case 'I', 'O', 'i', 'o':
														{
															position479 := position
															{
																position480, tokenIndex480 := position, tokenIndex
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
																		add(ruleAction93, position)
																	}
																	add(ruleInir, position482)
																}
																goto l480
															l481:
																position, tokenIndex = position480, tokenIndex480
																{
																	position494 := position
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
																				goto l493
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
																				goto l493
																			}
																			position++
																		}
																	l498:
																		{
																			position500, tokenIndex500 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l501
																			}
																			position++
																			goto l500
																		l501:
																			position, tokenIndex = position500, tokenIndex500
																			if buffer[position] != rune('I') {
																				goto l493
																			}
																			position++
																		}
																	l500:
																		add(rulePegText, position495)
																	}
																	{
																		add(ruleAction85, position)
																	}
																	add(ruleIni, position494)
																}
																goto l480
															l493:
																position, tokenIndex = position480, tokenIndex480
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
																			if buffer[position] != rune('t') {
																				goto l509
																			}
																			position++
																			goto l508
																		l509:
																			position, tokenIndex = position508, tokenIndex508
																			if buffer[position] != rune('T') {
																				goto l503
																			}
																			position++
																		}
																	l508:
																		{
																			position510, tokenIndex510 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l511
																			}
																			position++
																			goto l510
																		l511:
																			position, tokenIndex = position510, tokenIndex510
																			if buffer[position] != rune('I') {
																				goto l503
																			}
																			position++
																		}
																	l510:
																		{
																			position512, tokenIndex512 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l513
																			}
																			position++
																			goto l512
																		l513:
																			position, tokenIndex = position512, tokenIndex512
																			if buffer[position] != rune('R') {
																				goto l503
																			}
																			position++
																		}
																	l512:
																		add(rulePegText, position505)
																	}
																	{
																		add(ruleAction94, position)
																	}
																	add(ruleOtir, position504)
																}
																goto l480
															l503:
																position, tokenIndex = position480, tokenIndex480
																{
																	position516 := position
																	{
																		position517 := position
																		{
																			position518, tokenIndex518 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l519
																			}
																			position++
																			goto l518
																		l519:
																			position, tokenIndex = position518, tokenIndex518
																			if buffer[position] != rune('O') {
																				goto l515
																			}
																			position++
																		}
																	l518:
																		{
																			position520, tokenIndex520 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l521
																			}
																			position++
																			goto l520
																		l521:
																			position, tokenIndex = position520, tokenIndex520
																			if buffer[position] != rune('U') {
																				goto l515
																			}
																			position++
																		}
																	l520:
																		{
																			position522, tokenIndex522 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l523
																			}
																			position++
																			goto l522
																		l523:
																			position, tokenIndex = position522, tokenIndex522
																			if buffer[position] != rune('T') {
																				goto l515
																			}
																			position++
																		}
																	l522:
																		{
																			position524, tokenIndex524 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l525
																			}
																			position++
																			goto l524
																		l525:
																			position, tokenIndex = position524, tokenIndex524
																			if buffer[position] != rune('I') {
																				goto l515
																			}
																			position++
																		}
																	l524:
																		add(rulePegText, position517)
																	}
																	{
																		add(ruleAction86, position)
																	}
																	add(ruleOuti, position516)
																}
																goto l480
															l515:
																position, tokenIndex = position480, tokenIndex480
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
																		{
																			position536, tokenIndex536 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l537
																			}
																			position++
																			goto l536
																		l537:
																			position, tokenIndex = position536, tokenIndex536
																			if buffer[position] != rune('R') {
																				goto l527
																			}
																			position++
																		}
																	l536:
																		add(rulePegText, position529)
																	}
																	{
																		add(ruleAction97, position)
																	}
																	add(ruleIndr, position528)
																}
																goto l480
															l527:
																position, tokenIndex = position480, tokenIndex480
																{
																	position540 := position
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
																				goto l539
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
																				goto l539
																			}
																			position++
																		}
																	l544:
																		{
																			position546, tokenIndex546 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l547
																			}
																			position++
																			goto l546
																		l547:
																			position, tokenIndex = position546, tokenIndex546
																			if buffer[position] != rune('D') {
																				goto l539
																			}
																			position++
																		}
																	l546:
																		add(rulePegText, position541)
																	}
																	{
																		add(ruleAction89, position)
																	}
																	add(ruleInd, position540)
																}
																goto l480
															l539:
																position, tokenIndex = position480, tokenIndex480
																{
																	position550 := position
																	{
																		position551 := position
																		{
																			position552, tokenIndex552 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l553
																			}
																			position++
																			goto l552
																		l553:
																			position, tokenIndex = position552, tokenIndex552
																			if buffer[position] != rune('O') {
																				goto l549
																			}
																			position++
																		}
																	l552:
																		{
																			position554, tokenIndex554 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l555
																			}
																			position++
																			goto l554
																		l555:
																			position, tokenIndex = position554, tokenIndex554
																			if buffer[position] != rune('T') {
																				goto l549
																			}
																			position++
																		}
																	l554:
																		{
																			position556, tokenIndex556 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l557
																			}
																			position++
																			goto l556
																		l557:
																			position, tokenIndex = position556, tokenIndex556
																			if buffer[position] != rune('D') {
																				goto l549
																			}
																			position++
																		}
																	l556:
																		{
																			position558, tokenIndex558 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l559
																			}
																			position++
																			goto l558
																		l559:
																			position, tokenIndex = position558, tokenIndex558
																			if buffer[position] != rune('R') {
																				goto l549
																			}
																			position++
																		}
																	l558:
																		add(rulePegText, position551)
																	}
																	{
																		add(ruleAction98, position)
																	}
																	add(ruleOtdr, position550)
																}
																goto l480
															l549:
																position, tokenIndex = position480, tokenIndex480
																{
																	position561 := position
																	{
																		position562 := position
																		{
																			position563, tokenIndex563 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l564
																			}
																			position++
																			goto l563
																		l564:
																			position, tokenIndex = position563, tokenIndex563
																			if buffer[position] != rune('O') {
																				goto l417
																			}
																			position++
																		}
																	l563:
																		{
																			position565, tokenIndex565 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l566
																			}
																			position++
																			goto l565
																		l566:
																			position, tokenIndex = position565, tokenIndex565
																			if buffer[position] != rune('U') {
																				goto l417
																			}
																			position++
																		}
																	l565:
																		{
																			position567, tokenIndex567 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l568
																			}
																			position++
																			goto l567
																		l568:
																			position, tokenIndex = position567, tokenIndex567
																			if buffer[position] != rune('T') {
																				goto l417
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
																				goto l417
																			}
																			position++
																		}
																	l569:
																		add(rulePegText, position562)
																	}
																	{
																		add(ruleAction90, position)
																	}
																	add(ruleOutd, position561)
																}
															}
														l480:
															add(ruleBlitIO, position479)
														}
														break
													case 'R', 'r':
														{
															position572 := position
															{
																position573 := position
																{
																	position574, tokenIndex574 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l575
																	}
																	position++
																	goto l574
																l575:
																	position, tokenIndex = position574, tokenIndex574
																	if buffer[position] != rune('R') {
																		goto l417
																	}
																	position++
																}
															l574:
																{
																	position576, tokenIndex576 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l577
																	}
																	position++
																	goto l576
																l577:
																	position, tokenIndex = position576, tokenIndex576
																	if buffer[position] != rune('L') {
																		goto l417
																	}
																	position++
																}
															l576:
																{
																	position578, tokenIndex578 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l579
																	}
																	position++
																	goto l578
																l579:
																	position, tokenIndex = position578, tokenIndex578
																	if buffer[position] != rune('D') {
																		goto l417
																	}
																	position++
																}
															l578:
																add(rulePegText, position573)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleRld, position572)
														}
														break
													case 'N', 'n':
														{
															position581 := position
															{
																position582 := position
																{
																	position583, tokenIndex583 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l584
																	}
																	position++
																	goto l583
																l584:
																	position, tokenIndex = position583, tokenIndex583
																	if buffer[position] != rune('N') {
																		goto l417
																	}
																	position++
																}
															l583:
																{
																	position585, tokenIndex585 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l586
																	}
																	position++
																	goto l585
																l586:
																	position, tokenIndex = position585, tokenIndex585
																	if buffer[position] != rune('E') {
																		goto l417
																	}
																	position++
																}
															l585:
																{
																	position587, tokenIndex587 := position, tokenIndex
																	if buffer[position] != rune('g') {
																		goto l588
																	}
																	position++
																	goto l587
																l588:
																	position, tokenIndex = position587, tokenIndex587
																	if buffer[position] != rune('G') {
																		goto l417
																	}
																	position++
																}
															l587:
																add(rulePegText, position582)
															}
															{
																add(ruleAction75, position)
															}
															add(ruleNeg, position581)
														}
														break
													default:
														{
															position590 := position
															{
																position591, tokenIndex591 := position, tokenIndex
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
																				goto l592
																			}
																			position++
																		}
																	l601:
																		add(rulePegText, position594)
																	}
																	{
																		add(ruleAction91, position)
																	}
																	add(ruleLdir, position593)
																}
																goto l591
															l592:
																position, tokenIndex = position591, tokenIndex591
																{
																	position605 := position
																	{
																		position606 := position
																		{
																			position607, tokenIndex607 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l608
																			}
																			position++
																			goto l607
																		l608:
																			position, tokenIndex = position607, tokenIndex607
																			if buffer[position] != rune('L') {
																				goto l604
																			}
																			position++
																		}
																	l607:
																		{
																			position609, tokenIndex609 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l610
																			}
																			position++
																			goto l609
																		l610:
																			position, tokenIndex = position609, tokenIndex609
																			if buffer[position] != rune('D') {
																				goto l604
																			}
																			position++
																		}
																	l609:
																		{
																			position611, tokenIndex611 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l612
																			}
																			position++
																			goto l611
																		l612:
																			position, tokenIndex = position611, tokenIndex611
																			if buffer[position] != rune('I') {
																				goto l604
																			}
																			position++
																		}
																	l611:
																		add(rulePegText, position606)
																	}
																	{
																		add(ruleAction83, position)
																	}
																	add(ruleLdi, position605)
																}
																goto l591
															l604:
																position, tokenIndex = position591, tokenIndex591
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
																				goto l614
																			}
																			position++
																		}
																	l623:
																		add(rulePegText, position616)
																	}
																	{
																		add(ruleAction92, position)
																	}
																	add(ruleCpir, position615)
																}
																goto l591
															l614:
																position, tokenIndex = position591, tokenIndex591
																{
																	position627 := position
																	{
																		position628 := position
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
																				goto l626
																			}
																			position++
																		}
																	l629:
																		{
																			position631, tokenIndex631 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l632
																			}
																			position++
																			goto l631
																		l632:
																			position, tokenIndex = position631, tokenIndex631
																			if buffer[position] != rune('P') {
																				goto l626
																			}
																			position++
																		}
																	l631:
																		{
																			position633, tokenIndex633 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l634
																			}
																			position++
																			goto l633
																		l634:
																			position, tokenIndex = position633, tokenIndex633
																			if buffer[position] != rune('I') {
																				goto l626
																			}
																			position++
																		}
																	l633:
																		add(rulePegText, position628)
																	}
																	{
																		add(ruleAction84, position)
																	}
																	add(ruleCpi, position627)
																}
																goto l591
															l626:
																position, tokenIndex = position591, tokenIndex591
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
																		{
																			position645, tokenIndex645 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l646
																			}
																			position++
																			goto l645
																		l646:
																			position, tokenIndex = position645, tokenIndex645
																			if buffer[position] != rune('R') {
																				goto l636
																			}
																			position++
																		}
																	l645:
																		add(rulePegText, position638)
																	}
																	{
																		add(ruleAction95, position)
																	}
																	add(ruleLddr, position637)
																}
																goto l591
															l636:
																position, tokenIndex = position591, tokenIndex591
																{
																	position649 := position
																	{
																		position650 := position
																		{
																			position651, tokenIndex651 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l652
																			}
																			position++
																			goto l651
																		l652:
																			position, tokenIndex = position651, tokenIndex651
																			if buffer[position] != rune('L') {
																				goto l648
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
																				goto l648
																			}
																			position++
																		}
																	l653:
																		{
																			position655, tokenIndex655 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l656
																			}
																			position++
																			goto l655
																		l656:
																			position, tokenIndex = position655, tokenIndex655
																			if buffer[position] != rune('D') {
																				goto l648
																			}
																			position++
																		}
																	l655:
																		add(rulePegText, position650)
																	}
																	{
																		add(ruleAction87, position)
																	}
																	add(ruleLdd, position649)
																}
																goto l591
															l648:
																position, tokenIndex = position591, tokenIndex591
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
																			if buffer[position] != rune('d') {
																				goto l666
																			}
																			position++
																			goto l665
																		l666:
																			position, tokenIndex = position665, tokenIndex665
																			if buffer[position] != rune('D') {
																				goto l658
																			}
																			position++
																		}
																	l665:
																		{
																			position667, tokenIndex667 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l668
																			}
																			position++
																			goto l667
																		l668:
																			position, tokenIndex = position667, tokenIndex667
																			if buffer[position] != rune('R') {
																				goto l658
																			}
																			position++
																		}
																	l667:
																		add(rulePegText, position660)
																	}
																	{
																		add(ruleAction96, position)
																	}
																	add(ruleCpdr, position659)
																}
																goto l591
															l658:
																position, tokenIndex = position591, tokenIndex591
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
																				goto l417
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
																				goto l417
																			}
																			position++
																		}
																	l674:
																		{
																			position676, tokenIndex676 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l677
																			}
																			position++
																			goto l676
																		l677:
																			position, tokenIndex = position676, tokenIndex676
																			if buffer[position] != rune('D') {
																				goto l417
																			}
																			position++
																		}
																	l676:
																		add(rulePegText, position671)
																	}
																	{
																		add(ruleAction88, position)
																	}
																	add(ruleCpd, position670)
																}
															}
														l591:
															add(ruleBlit, position590)
														}
														break
													}
												}

											}
										l419:
											add(ruleEDSimple, position418)
										}
										goto l89
									l417:
										position, tokenIndex = position89, tokenIndex89
										{
											position680 := position
											{
												position681, tokenIndex681 := position, tokenIndex
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
															if buffer[position] != rune('l') {
																goto l688
															}
															position++
															goto l687
														l688:
															position, tokenIndex = position687, tokenIndex687
															if buffer[position] != rune('L') {
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
														add(ruleAction64, position)
													}
													add(ruleRlca, position683)
												}
												goto l681
											l682:
												position, tokenIndex = position681, tokenIndex681
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
															if buffer[position] != rune('r') {
																goto l700
															}
															position++
															goto l699
														l700:
															position, tokenIndex = position699, tokenIndex699
															if buffer[position] != rune('R') {
																goto l694
															}
															position++
														}
													l699:
														{
															position701, tokenIndex701 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l702
															}
															position++
															goto l701
														l702:
															position, tokenIndex = position701, tokenIndex701
															if buffer[position] != rune('C') {
																goto l694
															}
															position++
														}
													l701:
														{
															position703, tokenIndex703 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l704
															}
															position++
															goto l703
														l704:
															position, tokenIndex = position703, tokenIndex703
															if buffer[position] != rune('A') {
																goto l694
															}
															position++
														}
													l703:
														add(rulePegText, position696)
													}
													{
														add(ruleAction65, position)
													}
													add(ruleRrca, position695)
												}
												goto l681
											l694:
												position, tokenIndex = position681, tokenIndex681
												{
													position707 := position
													{
														position708 := position
														{
															position709, tokenIndex709 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l710
															}
															position++
															goto l709
														l710:
															position, tokenIndex = position709, tokenIndex709
															if buffer[position] != rune('R') {
																goto l706
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
																goto l706
															}
															position++
														}
													l711:
														{
															position713, tokenIndex713 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l714
															}
															position++
															goto l713
														l714:
															position, tokenIndex = position713, tokenIndex713
															if buffer[position] != rune('A') {
																goto l706
															}
															position++
														}
													l713:
														add(rulePegText, position708)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleRla, position707)
												}
												goto l681
											l706:
												position, tokenIndex = position681, tokenIndex681
												{
													position717 := position
													{
														position718 := position
														{
															position719, tokenIndex719 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l720
															}
															position++
															goto l719
														l720:
															position, tokenIndex = position719, tokenIndex719
															if buffer[position] != rune('D') {
																goto l716
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
																goto l716
															}
															position++
														}
													l721:
														{
															position723, tokenIndex723 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l724
															}
															position++
															goto l723
														l724:
															position, tokenIndex = position723, tokenIndex723
															if buffer[position] != rune('A') {
																goto l716
															}
															position++
														}
													l723:
														add(rulePegText, position718)
													}
													{
														add(ruleAction68, position)
													}
													add(ruleDaa, position717)
												}
												goto l681
											l716:
												position, tokenIndex = position681, tokenIndex681
												{
													position727 := position
													{
														position728 := position
														{
															position729, tokenIndex729 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l730
															}
															position++
															goto l729
														l730:
															position, tokenIndex = position729, tokenIndex729
															if buffer[position] != rune('C') {
																goto l726
															}
															position++
														}
													l729:
														{
															position731, tokenIndex731 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l732
															}
															position++
															goto l731
														l732:
															position, tokenIndex = position731, tokenIndex731
															if buffer[position] != rune('P') {
																goto l726
															}
															position++
														}
													l731:
														{
															position733, tokenIndex733 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l734
															}
															position++
															goto l733
														l734:
															position, tokenIndex = position733, tokenIndex733
															if buffer[position] != rune('L') {
																goto l726
															}
															position++
														}
													l733:
														add(rulePegText, position728)
													}
													{
														add(ruleAction69, position)
													}
													add(ruleCpl, position727)
												}
												goto l681
											l726:
												position, tokenIndex = position681, tokenIndex681
												{
													position737 := position
													{
														position738 := position
														{
															position739, tokenIndex739 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l740
															}
															position++
															goto l739
														l740:
															position, tokenIndex = position739, tokenIndex739
															if buffer[position] != rune('E') {
																goto l736
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
																goto l736
															}
															position++
														}
													l741:
														{
															position743, tokenIndex743 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l744
															}
															position++
															goto l743
														l744:
															position, tokenIndex = position743, tokenIndex743
															if buffer[position] != rune('X') {
																goto l736
															}
															position++
														}
													l743:
														add(rulePegText, position738)
													}
													{
														add(ruleAction72, position)
													}
													add(ruleExx, position737)
												}
												goto l681
											l736:
												position, tokenIndex = position681, tokenIndex681
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position747 := position
															{
																position748 := position
																{
																	position749, tokenIndex749 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l750
																	}
																	position++
																	goto l749
																l750:
																	position, tokenIndex = position749, tokenIndex749
																	if buffer[position] != rune('E') {
																		goto l679
																	}
																	position++
																}
															l749:
																{
																	position751, tokenIndex751 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l752
																	}
																	position++
																	goto l751
																l752:
																	position, tokenIndex = position751, tokenIndex751
																	if buffer[position] != rune('I') {
																		goto l679
																	}
																	position++
																}
															l751:
																add(rulePegText, position748)
															}
															{
																add(ruleAction74, position)
															}
															add(ruleEi, position747)
														}
														break
													case 'D', 'd':
														{
															position754 := position
															{
																position755 := position
																{
																	position756, tokenIndex756 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l757
																	}
																	position++
																	goto l756
																l757:
																	position, tokenIndex = position756, tokenIndex756
																	if buffer[position] != rune('D') {
																		goto l679
																	}
																	position++
																}
															l756:
																{
																	position758, tokenIndex758 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l759
																	}
																	position++
																	goto l758
																l759:
																	position, tokenIndex = position758, tokenIndex758
																	if buffer[position] != rune('I') {
																		goto l679
																	}
																	position++
																}
															l758:
																add(rulePegText, position755)
															}
															{
																add(ruleAction73, position)
															}
															add(ruleDi, position754)
														}
														break
													case 'C', 'c':
														{
															position761 := position
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
																		goto l679
																	}
																	position++
																}
															l763:
																{
																	position765, tokenIndex765 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l766
																	}
																	position++
																	goto l765
																l766:
																	position, tokenIndex = position765, tokenIndex765
																	if buffer[position] != rune('C') {
																		goto l679
																	}
																	position++
																}
															l765:
																{
																	position767, tokenIndex767 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l768
																	}
																	position++
																	goto l767
																l768:
																	position, tokenIndex = position767, tokenIndex767
																	if buffer[position] != rune('F') {
																		goto l679
																	}
																	position++
																}
															l767:
																add(rulePegText, position762)
															}
															{
																add(ruleAction71, position)
															}
															add(ruleCcf, position761)
														}
														break
													case 'S', 's':
														{
															position770 := position
															{
																position771 := position
																{
																	position772, tokenIndex772 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l773
																	}
																	position++
																	goto l772
																l773:
																	position, tokenIndex = position772, tokenIndex772
																	if buffer[position] != rune('S') {
																		goto l679
																	}
																	position++
																}
															l772:
																{
																	position774, tokenIndex774 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l775
																	}
																	position++
																	goto l774
																l775:
																	position, tokenIndex = position774, tokenIndex774
																	if buffer[position] != rune('C') {
																		goto l679
																	}
																	position++
																}
															l774:
																{
																	position776, tokenIndex776 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l777
																	}
																	position++
																	goto l776
																l777:
																	position, tokenIndex = position776, tokenIndex776
																	if buffer[position] != rune('F') {
																		goto l679
																	}
																	position++
																}
															l776:
																add(rulePegText, position771)
															}
															{
																add(ruleAction70, position)
															}
															add(ruleScf, position770)
														}
														break
													case 'R', 'r':
														{
															position779 := position
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
																		goto l679
																	}
																	position++
																}
															l781:
																{
																	position783, tokenIndex783 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l784
																	}
																	position++
																	goto l783
																l784:
																	position, tokenIndex = position783, tokenIndex783
																	if buffer[position] != rune('R') {
																		goto l679
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
																		goto l679
																	}
																	position++
																}
															l785:
																add(rulePegText, position780)
															}
															{
																add(ruleAction67, position)
															}
															add(ruleRra, position779)
														}
														break
													case 'H', 'h':
														{
															position788 := position
															{
																position789 := position
																{
																	position790, tokenIndex790 := position, tokenIndex
																	if buffer[position] != rune('h') {
																		goto l791
																	}
																	position++
																	goto l790
																l791:
																	position, tokenIndex = position790, tokenIndex790
																	if buffer[position] != rune('H') {
																		goto l679
																	}
																	position++
																}
															l790:
																{
																	position792, tokenIndex792 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l793
																	}
																	position++
																	goto l792
																l793:
																	position, tokenIndex = position792, tokenIndex792
																	if buffer[position] != rune('A') {
																		goto l679
																	}
																	position++
																}
															l792:
																{
																	position794, tokenIndex794 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l795
																	}
																	position++
																	goto l794
																l795:
																	position, tokenIndex = position794, tokenIndex794
																	if buffer[position] != rune('L') {
																		goto l679
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
																		goto l679
																	}
																	position++
																}
															l796:
																add(rulePegText, position789)
															}
															{
																add(ruleAction63, position)
															}
															add(ruleHalt, position788)
														}
														break
													default:
														{
															position799 := position
															{
																position800 := position
																{
																	position801, tokenIndex801 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l802
																	}
																	position++
																	goto l801
																l802:
																	position, tokenIndex = position801, tokenIndex801
																	if buffer[position] != rune('N') {
																		goto l679
																	}
																	position++
																}
															l801:
																{
																	position803, tokenIndex803 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l804
																	}
																	position++
																	goto l803
																l804:
																	position, tokenIndex = position803, tokenIndex803
																	if buffer[position] != rune('O') {
																		goto l679
																	}
																	position++
																}
															l803:
																{
																	position805, tokenIndex805 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l806
																	}
																	position++
																	goto l805
																l806:
																	position, tokenIndex = position805, tokenIndex805
																	if buffer[position] != rune('P') {
																		goto l679
																	}
																	position++
																}
															l805:
																add(rulePegText, position800)
															}
															{
																add(ruleAction62, position)
															}
															add(ruleNop, position799)
														}
														break
													}
												}

											}
										l681:
											add(ruleSimple, position680)
										}
										goto l89
									l679:
										position, tokenIndex = position89, tokenIndex89
										{
											position809 := position
											{
												position810, tokenIndex810 := position, tokenIndex
												{
													position812 := position
													{
														position813, tokenIndex813 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l814
														}
														position++
														goto l813
													l814:
														position, tokenIndex = position813, tokenIndex813
														if buffer[position] != rune('R') {
															goto l811
														}
														position++
													}
												l813:
													{
														position815, tokenIndex815 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l816
														}
														position++
														goto l815
													l816:
														position, tokenIndex = position815, tokenIndex815
														if buffer[position] != rune('S') {
															goto l811
														}
														position++
													}
												l815:
													{
														position817, tokenIndex817 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l818
														}
														position++
														goto l817
													l818:
														position, tokenIndex = position817, tokenIndex817
														if buffer[position] != rune('T') {
															goto l811
														}
														position++
													}
												l817:
													if !_rules[rulews]() {
														goto l811
													}
													if !_rules[rulen]() {
														goto l811
													}
													{
														add(ruleAction99, position)
													}
													add(ruleRst, position812)
												}
												goto l810
											l811:
												position, tokenIndex = position810, tokenIndex810
												{
													position821 := position
													{
														position822, tokenIndex822 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l823
														}
														position++
														goto l822
													l823:
														position, tokenIndex = position822, tokenIndex822
														if buffer[position] != rune('J') {
															goto l820
														}
														position++
													}
												l822:
													{
														position824, tokenIndex824 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l825
														}
														position++
														goto l824
													l825:
														position, tokenIndex = position824, tokenIndex824
														if buffer[position] != rune('P') {
															goto l820
														}
														position++
													}
												l824:
													if !_rules[rulews]() {
														goto l820
													}
													{
														position826, tokenIndex826 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l826
														}
														if !_rules[rulesep]() {
															goto l826
														}
														goto l827
													l826:
														position, tokenIndex = position826, tokenIndex826
													}
												l827:
													if !_rules[ruleSrc16]() {
														goto l820
													}
													{
														add(ruleAction102, position)
													}
													add(ruleJp, position821)
												}
												goto l810
											l820:
												position, tokenIndex = position810, tokenIndex810
												{
													switch buffer[position] {
													case 'D', 'd':
														{
															position830 := position
															{
																position831, tokenIndex831 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l832
																}
																position++
																goto l831
															l832:
																position, tokenIndex = position831, tokenIndex831
																if buffer[position] != rune('D') {
																	goto l808
																}
																position++
															}
														l831:
															{
																position833, tokenIndex833 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l834
																}
																position++
																goto l833
															l834:
																position, tokenIndex = position833, tokenIndex833
																if buffer[position] != rune('J') {
																	goto l808
																}
																position++
															}
														l833:
															{
																position835, tokenIndex835 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l836
																}
																position++
																goto l835
															l836:
																position, tokenIndex = position835, tokenIndex835
																if buffer[position] != rune('N') {
																	goto l808
																}
																position++
															}
														l835:
															{
																position837, tokenIndex837 := position, tokenIndex
																if buffer[position] != rune('z') {
																	goto l838
																}
																position++
																goto l837
															l838:
																position, tokenIndex = position837, tokenIndex837
																if buffer[position] != rune('Z') {
																	goto l808
																}
																position++
															}
														l837:
															if !_rules[rulews]() {
																goto l808
															}
															if !_rules[ruledisp]() {
																goto l808
															}
															{
																add(ruleAction104, position)
															}
															add(ruleDjnz, position830)
														}
														break
													case 'J', 'j':
														{
															position840 := position
															{
																position841, tokenIndex841 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l842
																}
																position++
																goto l841
															l842:
																position, tokenIndex = position841, tokenIndex841
																if buffer[position] != rune('J') {
																	goto l808
																}
																position++
															}
														l841:
															{
																position843, tokenIndex843 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l844
																}
																position++
																goto l843
															l844:
																position, tokenIndex = position843, tokenIndex843
																if buffer[position] != rune('R') {
																	goto l808
																}
																position++
															}
														l843:
															if !_rules[rulews]() {
																goto l808
															}
															{
																position845, tokenIndex845 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l845
																}
																if !_rules[rulesep]() {
																	goto l845
																}
																goto l846
															l845:
																position, tokenIndex = position845, tokenIndex845
															}
														l846:
															if !_rules[ruledisp]() {
																goto l808
															}
															{
																add(ruleAction103, position)
															}
															add(ruleJr, position840)
														}
														break
													case 'R', 'r':
														{
															position848 := position
															{
																position849, tokenIndex849 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l850
																}
																position++
																goto l849
															l850:
																position, tokenIndex = position849, tokenIndex849
																if buffer[position] != rune('R') {
																	goto l808
																}
																position++
															}
														l849:
															{
																position851, tokenIndex851 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l852
																}
																position++
																goto l851
															l852:
																position, tokenIndex = position851, tokenIndex851
																if buffer[position] != rune('E') {
																	goto l808
																}
																position++
															}
														l851:
															{
																position853, tokenIndex853 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l854
																}
																position++
																goto l853
															l854:
																position, tokenIndex = position853, tokenIndex853
																if buffer[position] != rune('T') {
																	goto l808
																}
																position++
															}
														l853:
															{
																position855, tokenIndex855 := position, tokenIndex
																if !_rules[rulews]() {
																	goto l855
																}
																if !_rules[rulecc]() {
																	goto l855
																}
																goto l856
															l855:
																position, tokenIndex = position855, tokenIndex855
															}
														l856:
															{
																add(ruleAction101, position)
															}
															add(ruleRet, position848)
														}
														break
													default:
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
																	goto l808
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
																	goto l808
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
																	goto l808
																}
																position++
															}
														l863:
															{
																position865, tokenIndex865 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l866
																}
																position++
																goto l865
															l866:
																position, tokenIndex = position865, tokenIndex865
																if buffer[position] != rune('L') {
																	goto l808
																}
																position++
															}
														l865:
															if !_rules[rulews]() {
																goto l808
															}
															{
																position867, tokenIndex867 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l867
																}
																if !_rules[rulesep]() {
																	goto l867
																}
																goto l868
															l867:
																position, tokenIndex = position867, tokenIndex867
															}
														l868:
															if !_rules[ruleSrc16]() {
																goto l808
															}
															{
																add(ruleAction100, position)
															}
															add(ruleCall, position858)
														}
														break
													}
												}

											}
										l810:
											add(ruleJump, position809)
										}
										goto l89
									l808:
										position, tokenIndex = position89, tokenIndex89
										{
											position870 := position
											{
												position871, tokenIndex871 := position, tokenIndex
												{
													position873 := position
													{
														position874, tokenIndex874 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l875
														}
														position++
														goto l874
													l875:
														position, tokenIndex = position874, tokenIndex874
														if buffer[position] != rune('I') {
															goto l872
														}
														position++
													}
												l874:
													{
														position876, tokenIndex876 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l877
														}
														position++
														goto l876
													l877:
														position, tokenIndex = position876, tokenIndex876
														if buffer[position] != rune('N') {
															goto l872
														}
														position++
													}
												l876:
													if !_rules[rulews]() {
														goto l872
													}
													if !_rules[ruleReg8]() {
														goto l872
													}
													if !_rules[rulesep]() {
														goto l872
													}
													if !_rules[rulePort]() {
														goto l872
													}
													{
														add(ruleAction105, position)
													}
													add(ruleIN, position873)
												}
												goto l871
											l872:
												position, tokenIndex = position871, tokenIndex871
												{
													position879 := position
													{
														position880, tokenIndex880 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l881
														}
														position++
														goto l880
													l881:
														position, tokenIndex = position880, tokenIndex880
														if buffer[position] != rune('O') {
															goto l15
														}
														position++
													}
												l880:
													{
														position882, tokenIndex882 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l883
														}
														position++
														goto l882
													l883:
														position, tokenIndex = position882, tokenIndex882
														if buffer[position] != rune('U') {
															goto l15
														}
														position++
													}
												l882:
													{
														position884, tokenIndex884 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l885
														}
														position++
														goto l884
													l885:
														position, tokenIndex = position884, tokenIndex884
														if buffer[position] != rune('T') {
															goto l15
														}
														position++
													}
												l884:
													if !_rules[rulews]() {
														goto l15
													}
													if !_rules[rulePort]() {
														goto l15
													}
													if !_rules[rulesep]() {
														goto l15
													}
													if !_rules[ruleReg8]() {
														goto l15
													}
													{
														add(ruleAction106, position)
													}
													add(ruleOUT, position879)
												}
											}
										l871:
											add(ruleIO, position870)
										}
									}
								l89:
									add(ruleInstruction, position88)
								}
							}
						l18:
							add(ruleStatement, position17)
						}
						goto l16
					l15:
						position, tokenIndex = position15, tokenIndex15
					}
				l16:
					{
						position887, tokenIndex887 := position, tokenIndex
						if !_rules[rulews]() {
							goto l887
						}
						goto l888
					l887:
						position, tokenIndex = position887, tokenIndex887
					}
				l888:
					{
						position889, tokenIndex889 := position, tokenIndex
						{
							position891 := position
							{
								position892, tokenIndex892 := position, tokenIndex
								if buffer[position] != rune(';') {
									goto l893
								}
								position++
								goto l892
							l893:
								position, tokenIndex = position892, tokenIndex892
								if buffer[position] != rune('#') {
									goto l889
								}
								position++
							}
						l892:
						l894:
							{
								position895, tokenIndex895 := position, tokenIndex
								{
									position896, tokenIndex896 := position, tokenIndex
									if buffer[position] != rune('\n') {
										goto l896
									}
									position++
									goto l895
								l896:
									position, tokenIndex = position896, tokenIndex896
								}
								if !matchDot() {
									goto l895
								}
								goto l894
							l895:
								position, tokenIndex = position895, tokenIndex895
							}
							add(ruleComment, position891)
						}
						goto l890
					l889:
						position, tokenIndex = position889, tokenIndex889
					}
				l890:
					{
						position897, tokenIndex897 := position, tokenIndex
						if !_rules[rulews]() {
							goto l897
						}
						goto l898
					l897:
						position, tokenIndex = position897, tokenIndex897
					}
				l898:
					{
						position899, tokenIndex899 := position, tokenIndex
						{
							position901, tokenIndex901 := position, tokenIndex
							if buffer[position] != rune('\r') {
								goto l901
							}
							position++
							goto l902
						l901:
							position, tokenIndex = position901, tokenIndex901
						}
					l902:
						if buffer[position] != rune('\n') {
							goto l900
						}
						position++
						goto l899
					l900:
						position, tokenIndex = position899, tokenIndex899
						if buffer[position] != rune(':') {
							goto l0
						}
						position++
					}
				l899:
					{
						add(ruleAction0, position)
					}
					add(ruleLine, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position904 := position
					l905:
						{
							position906, tokenIndex906 := position, tokenIndex
							if !_rules[rulews]() {
								goto l906
							}
							goto l905
						l906:
							position, tokenIndex = position906, tokenIndex906
						}
						{
							position907, tokenIndex907 := position, tokenIndex
							{
								position909 := position
								if !_rules[ruleLabelText]() {
									goto l907
								}
								if buffer[position] != rune(':') {
									goto l907
								}
								position++
								{
									position910, tokenIndex910 := position, tokenIndex
									if !_rules[rulews]() {
										goto l910
									}
									goto l911
								l910:
									position, tokenIndex = position910, tokenIndex910
								}
							l911:
								{
									add(ruleAction5, position)
								}
								add(ruleLabelDefn, position909)
							}
							goto l908
						l907:
							position, tokenIndex = position907, tokenIndex907
						}
					l908:
					l913:
						{
							position914, tokenIndex914 := position, tokenIndex
							if !_rules[rulews]() {
								goto l914
							}
							goto l913
						l914:
							position, tokenIndex = position914, tokenIndex914
						}
						{
							position915, tokenIndex915 := position, tokenIndex
							{
								position917 := position
								{
									position918, tokenIndex918 := position, tokenIndex
									{
										position920 := position
										{
											position921, tokenIndex921 := position, tokenIndex
											{
												position923 := position
												{
													position924, tokenIndex924 := position, tokenIndex
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
															goto l925
														}
														position++
													}
												l926:
													{
														position928, tokenIndex928 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l929
														}
														position++
														goto l928
													l929:
														position, tokenIndex = position928, tokenIndex928
														if buffer[position] != rune('E') {
															goto l925
														}
														position++
													}
												l928:
													{
														position930, tokenIndex930 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l931
														}
														position++
														goto l930
													l931:
														position, tokenIndex = position930, tokenIndex930
														if buffer[position] != rune('F') {
															goto l925
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
															goto l925
														}
														position++
													}
												l932:
													goto l924
												l925:
													position, tokenIndex = position924, tokenIndex924
													{
														position934, tokenIndex934 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l935
														}
														position++
														goto l934
													l935:
														position, tokenIndex = position934, tokenIndex934
														if buffer[position] != rune('D') {
															goto l922
														}
														position++
													}
												l934:
													{
														position936, tokenIndex936 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l937
														}
														position++
														goto l936
													l937:
														position, tokenIndex = position936, tokenIndex936
														if buffer[position] != rune('B') {
															goto l922
														}
														position++
													}
												l936:
												}
											l924:
												if !_rules[rulews]() {
													goto l922
												}
												if !_rules[rulen]() {
													goto l922
												}
												{
													add(ruleAction2, position)
												}
												add(ruleDefb, position923)
											}
											goto l921
										l922:
											position, tokenIndex = position921, tokenIndex921
											{
												position940 := position
												{
													position941, tokenIndex941 := position, tokenIndex
													{
														position943, tokenIndex943 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l944
														}
														position++
														goto l943
													l944:
														position, tokenIndex = position943, tokenIndex943
														if buffer[position] != rune('D') {
															goto l942
														}
														position++
													}
												l943:
													{
														position945, tokenIndex945 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l946
														}
														position++
														goto l945
													l946:
														position, tokenIndex = position945, tokenIndex945
														if buffer[position] != rune('E') {
															goto l942
														}
														position++
													}
												l945:
													{
														position947, tokenIndex947 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l948
														}
														position++
														goto l947
													l948:
														position, tokenIndex = position947, tokenIndex947
														if buffer[position] != rune('F') {
															goto l942
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
															goto l942
														}
														position++
													}
												l949:
													goto l941
												l942:
													position, tokenIndex = position941, tokenIndex941
													{
														position951, tokenIndex951 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l952
														}
														position++
														goto l951
													l952:
														position, tokenIndex = position951, tokenIndex951
														if buffer[position] != rune('D') {
															goto l939
														}
														position++
													}
												l951:
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
															goto l939
														}
														position++
													}
												l953:
												}
											l941:
												if !_rules[rulews]() {
													goto l939
												}
												if !_rules[rulen]() {
													goto l939
												}
												{
													add(ruleAction4, position)
												}
												add(ruleDefs, position940)
											}
											goto l921
										l939:
											position, tokenIndex = position921, tokenIndex921
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position957 := position
														{
															position958, tokenIndex958 := position, tokenIndex
															{
																position960, tokenIndex960 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l961
																}
																position++
																goto l960
															l961:
																position, tokenIndex = position960, tokenIndex960
																if buffer[position] != rune('D') {
																	goto l959
																}
																position++
															}
														l960:
															{
																position962, tokenIndex962 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l963
																}
																position++
																goto l962
															l963:
																position, tokenIndex = position962, tokenIndex962
																if buffer[position] != rune('E') {
																	goto l959
																}
																position++
															}
														l962:
															{
																position964, tokenIndex964 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l965
																}
																position++
																goto l964
															l965:
																position, tokenIndex = position964, tokenIndex964
																if buffer[position] != rune('F') {
																	goto l959
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
																	goto l959
																}
																position++
															}
														l966:
															goto l958
														l959:
															position, tokenIndex = position958, tokenIndex958
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
																	goto l919
																}
																position++
															}
														l968:
															{
																position970, tokenIndex970 := position, tokenIndex
																if buffer[position] != rune('w') {
																	goto l971
																}
																position++
																goto l970
															l971:
																position, tokenIndex = position970, tokenIndex970
																if buffer[position] != rune('W') {
																	goto l919
																}
																position++
															}
														l970:
														}
													l958:
														if !_rules[rulews]() {
															goto l919
														}
														if !_rules[rulenn]() {
															goto l919
														}
														{
															add(ruleAction3, position)
														}
														add(ruleDefw, position957)
													}
													break
												case 'O', 'o':
													{
														position973 := position
														{
															position974, tokenIndex974 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l975
															}
															position++
															goto l974
														l975:
															position, tokenIndex = position974, tokenIndex974
															if buffer[position] != rune('O') {
																goto l919
															}
															position++
														}
													l974:
														{
															position976, tokenIndex976 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l977
															}
															position++
															goto l976
														l977:
															position, tokenIndex = position976, tokenIndex976
															if buffer[position] != rune('R') {
																goto l919
															}
															position++
														}
													l976:
														{
															position978, tokenIndex978 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l979
															}
															position++
															goto l978
														l979:
															position, tokenIndex = position978, tokenIndex978
															if buffer[position] != rune('G') {
																goto l919
															}
															position++
														}
													l978:
														if !_rules[rulews]() {
															goto l919
														}
														if !_rules[rulenn]() {
															goto l919
														}
														{
															add(ruleAction1, position)
														}
														add(ruleOrg, position973)
													}
													break
												case 'a':
													{
														position981 := position
														if buffer[position] != rune('a') {
															goto l919
														}
														position++
														if buffer[position] != rune('s') {
															goto l919
														}
														position++
														if buffer[position] != rune('e') {
															goto l919
														}
														position++
														if buffer[position] != rune('g') {
															goto l919
														}
														position++
														add(ruleAseg, position981)
													}
													break
												default:
													{
														position982 := position
														{
															position983, tokenIndex983 := position, tokenIndex
															if buffer[position] != rune('.') {
																goto l983
															}
															position++
															goto l984
														l983:
															position, tokenIndex = position983, tokenIndex983
														}
													l984:
														if buffer[position] != rune('t') {
															goto l919
														}
														position++
														if buffer[position] != rune('i') {
															goto l919
														}
														position++
														if buffer[position] != rune('t') {
															goto l919
														}
														position++
														if buffer[position] != rune('l') {
															goto l919
														}
														position++
														if buffer[position] != rune('e') {
															goto l919
														}
														position++
														if !_rules[rulews]() {
															goto l919
														}
														if buffer[position] != rune('\'') {
															goto l919
														}
														position++
													l985:
														{
															position986, tokenIndex986 := position, tokenIndex
															{
																position987, tokenIndex987 := position, tokenIndex
																if buffer[position] != rune('\'') {
																	goto l987
																}
																position++
																goto l986
															l987:
																position, tokenIndex = position987, tokenIndex987
															}
															if !matchDot() {
																goto l986
															}
															goto l985
														l986:
															position, tokenIndex = position986, tokenIndex986
														}
														if buffer[position] != rune('\'') {
															goto l919
														}
														position++
														add(ruleTitle, position982)
													}
													break
												}
											}

										}
									l921:
										add(ruleDirective, position920)
									}
									goto l918
								l919:
									position, tokenIndex = position918, tokenIndex918
									{
										position988 := position
										{
											position989, tokenIndex989 := position, tokenIndex
											{
												position991 := position
												{
													position992, tokenIndex992 := position, tokenIndex
													{
														position994 := position
														{
															position995, tokenIndex995 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l996
															}
															position++
															goto l995
														l996:
															position, tokenIndex = position995, tokenIndex995
															if buffer[position] != rune('P') {
																goto l993
															}
															position++
														}
													l995:
														{
															position997, tokenIndex997 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l998
															}
															position++
															goto l997
														l998:
															position, tokenIndex = position997, tokenIndex997
															if buffer[position] != rune('U') {
																goto l993
															}
															position++
														}
													l997:
														{
															position999, tokenIndex999 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1000
															}
															position++
															goto l999
														l1000:
															position, tokenIndex = position999, tokenIndex999
															if buffer[position] != rune('S') {
																goto l993
															}
															position++
														}
													l999:
														{
															position1001, tokenIndex1001 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l1002
															}
															position++
															goto l1001
														l1002:
															position, tokenIndex = position1001, tokenIndex1001
															if buffer[position] != rune('H') {
																goto l993
															}
															position++
														}
													l1001:
														if !_rules[rulews]() {
															goto l993
														}
														if !_rules[ruleSrc16]() {
															goto l993
														}
														{
															add(ruleAction8, position)
														}
														add(rulePush, position994)
													}
													goto l992
												l993:
													position, tokenIndex = position992, tokenIndex992
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1005 := position
																{
																	position1006, tokenIndex1006 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1007
																	}
																	position++
																	goto l1006
																l1007:
																	position, tokenIndex = position1006, tokenIndex1006
																	if buffer[position] != rune('E') {
																		goto l990
																	}
																	position++
																}
															l1006:
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
																		goto l990
																	}
																	position++
																}
															l1008:
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
																	add(ruleAction10, position)
																}
																add(ruleEx, position1005)
															}
															break
														case 'P', 'p':
															{
																position1011 := position
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
																		goto l990
																	}
																	position++
																}
															l1012:
																{
																	position1014, tokenIndex1014 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1015
																	}
																	position++
																	goto l1014
																l1015:
																	position, tokenIndex = position1014, tokenIndex1014
																	if buffer[position] != rune('O') {
																		goto l990
																	}
																	position++
																}
															l1014:
																{
																	position1016, tokenIndex1016 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1017
																	}
																	position++
																	goto l1016
																l1017:
																	position, tokenIndex = position1016, tokenIndex1016
																	if buffer[position] != rune('P') {
																		goto l990
																	}
																	position++
																}
															l1016:
																if !_rules[rulews]() {
																	goto l990
																}
																if !_rules[ruleDst16]() {
																	goto l990
																}
																{
																	add(ruleAction9, position)
																}
																add(rulePop, position1011)
															}
															break
														default:
															{
																position1019 := position
																{
																	position1020, tokenIndex1020 := position, tokenIndex
																	{
																		position1022 := position
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
																				goto l1021
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
																				goto l1021
																			}
																			position++
																		}
																	l1025:
																		if !_rules[rulews]() {
																			goto l1021
																		}
																		if !_rules[ruleDst16]() {
																			goto l1021
																		}
																		if !_rules[rulesep]() {
																			goto l1021
																		}
																		if !_rules[ruleSrc16]() {
																			goto l1021
																		}
																		{
																			add(ruleAction7, position)
																		}
																		add(ruleLoad16, position1022)
																	}
																	goto l1020
																l1021:
																	position, tokenIndex = position1020, tokenIndex1020
																	{
																		position1028 := position
																		{
																			position1029, tokenIndex1029 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l1030
																			}
																			position++
																			goto l1029
																		l1030:
																			position, tokenIndex = position1029, tokenIndex1029
																			if buffer[position] != rune('L') {
																				goto l990
																			}
																			position++
																		}
																	l1029:
																		{
																			position1031, tokenIndex1031 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l1032
																			}
																			position++
																			goto l1031
																		l1032:
																			position, tokenIndex = position1031, tokenIndex1031
																			if buffer[position] != rune('D') {
																				goto l990
																			}
																			position++
																		}
																	l1031:
																		if !_rules[rulews]() {
																			goto l990
																		}
																		{
																			position1033 := position
																			{
																				position1034, tokenIndex1034 := position, tokenIndex
																				if !_rules[ruleReg8]() {
																					goto l1035
																				}
																				goto l1034
																			l1035:
																				position, tokenIndex = position1034, tokenIndex1034
																				if !_rules[ruleReg16Contents]() {
																					goto l1036
																				}
																				goto l1034
																			l1036:
																				position, tokenIndex = position1034, tokenIndex1034
																				if !_rules[rulenn_contents]() {
																					goto l990
																				}
																			}
																		l1034:
																			{
																				add(ruleAction20, position)
																			}
																			add(ruleDst8, position1033)
																		}
																		if !_rules[rulesep]() {
																			goto l990
																		}
																		if !_rules[ruleSrc8]() {
																			goto l990
																		}
																		{
																			add(ruleAction6, position)
																		}
																		add(ruleLoad8, position1028)
																	}
																}
															l1020:
																add(ruleLoad, position1019)
															}
															break
														}
													}

												}
											l992:
												add(ruleAssignment, position991)
											}
											goto l989
										l990:
											position, tokenIndex = position989, tokenIndex989
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
																goto l1042
															}
															position++
														}
													l1048:
														if !_rules[rulews]() {
															goto l1042
														}
														if !_rules[ruleILoc8]() {
															goto l1042
														}
														{
															add(ruleAction11, position)
														}
														add(ruleInc16Indexed8, position1043)
													}
													goto l1041
												l1042:
													position, tokenIndex = position1041, tokenIndex1041
													{
														position1052 := position
														{
															position1053, tokenIndex1053 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1054
															}
															position++
															goto l1053
														l1054:
															position, tokenIndex = position1053, tokenIndex1053
															if buffer[position] != rune('I') {
																goto l1051
															}
															position++
														}
													l1053:
														{
															position1055, tokenIndex1055 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1056
															}
															position++
															goto l1055
														l1056:
															position, tokenIndex = position1055, tokenIndex1055
															if buffer[position] != rune('N') {
																goto l1051
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
																goto l1051
															}
															position++
														}
													l1057:
														if !_rules[rulews]() {
															goto l1051
														}
														if !_rules[ruleLoc16]() {
															goto l1051
														}
														{
															add(ruleAction13, position)
														}
														add(ruleInc16, position1052)
													}
													goto l1041
												l1051:
													position, tokenIndex = position1041, tokenIndex1041
													{
														position1060 := position
														{
															position1061, tokenIndex1061 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1062
															}
															position++
															goto l1061
														l1062:
															position, tokenIndex = position1061, tokenIndex1061
															if buffer[position] != rune('I') {
																goto l1039
															}
															position++
														}
													l1061:
														{
															position1063, tokenIndex1063 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1064
															}
															position++
															goto l1063
														l1064:
															position, tokenIndex = position1063, tokenIndex1063
															if buffer[position] != rune('N') {
																goto l1039
															}
															position++
														}
													l1063:
														{
															position1065, tokenIndex1065 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1066
															}
															position++
															goto l1065
														l1066:
															position, tokenIndex = position1065, tokenIndex1065
															if buffer[position] != rune('C') {
																goto l1039
															}
															position++
														}
													l1065:
														if !_rules[rulews]() {
															goto l1039
														}
														if !_rules[ruleLoc8]() {
															goto l1039
														}
														{
															add(ruleAction12, position)
														}
														add(ruleInc8, position1060)
													}
												}
											l1041:
												add(ruleInc, position1040)
											}
											goto l989
										l1039:
											position, tokenIndex = position989, tokenIndex989
											{
												position1069 := position
												{
													position1070, tokenIndex1070 := position, tokenIndex
													{
														position1072 := position
														{
															position1073, tokenIndex1073 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1074
															}
															position++
															goto l1073
														l1074:
															position, tokenIndex = position1073, tokenIndex1073
															if buffer[position] != rune('D') {
																goto l1071
															}
															position++
														}
													l1073:
														{
															position1075, tokenIndex1075 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1076
															}
															position++
															goto l1075
														l1076:
															position, tokenIndex = position1075, tokenIndex1075
															if buffer[position] != rune('E') {
																goto l1071
															}
															position++
														}
													l1075:
														{
															position1077, tokenIndex1077 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1078
															}
															position++
															goto l1077
														l1078:
															position, tokenIndex = position1077, tokenIndex1077
															if buffer[position] != rune('C') {
																goto l1071
															}
															position++
														}
													l1077:
														if !_rules[rulews]() {
															goto l1071
														}
														if !_rules[ruleILoc8]() {
															goto l1071
														}
														{
															add(ruleAction14, position)
														}
														add(ruleDec16Indexed8, position1072)
													}
													goto l1070
												l1071:
													position, tokenIndex = position1070, tokenIndex1070
													{
														position1081 := position
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
																goto l1080
															}
															position++
														}
													l1082:
														{
															position1084, tokenIndex1084 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1085
															}
															position++
															goto l1084
														l1085:
															position, tokenIndex = position1084, tokenIndex1084
															if buffer[position] != rune('E') {
																goto l1080
															}
															position++
														}
													l1084:
														{
															position1086, tokenIndex1086 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1087
															}
															position++
															goto l1086
														l1087:
															position, tokenIndex = position1086, tokenIndex1086
															if buffer[position] != rune('C') {
																goto l1080
															}
															position++
														}
													l1086:
														if !_rules[rulews]() {
															goto l1080
														}
														if !_rules[ruleLoc16]() {
															goto l1080
														}
														{
															add(ruleAction16, position)
														}
														add(ruleDec16, position1081)
													}
													goto l1070
												l1080:
													position, tokenIndex = position1070, tokenIndex1070
													{
														position1089 := position
														{
															position1090, tokenIndex1090 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1091
															}
															position++
															goto l1090
														l1091:
															position, tokenIndex = position1090, tokenIndex1090
															if buffer[position] != rune('D') {
																goto l1068
															}
															position++
														}
													l1090:
														{
															position1092, tokenIndex1092 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1093
															}
															position++
															goto l1092
														l1093:
															position, tokenIndex = position1092, tokenIndex1092
															if buffer[position] != rune('E') {
																goto l1068
															}
															position++
														}
													l1092:
														{
															position1094, tokenIndex1094 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1095
															}
															position++
															goto l1094
														l1095:
															position, tokenIndex = position1094, tokenIndex1094
															if buffer[position] != rune('C') {
																goto l1068
															}
															position++
														}
													l1094:
														if !_rules[rulews]() {
															goto l1068
														}
														if !_rules[ruleLoc8]() {
															goto l1068
														}
														{
															add(ruleAction15, position)
														}
														add(ruleDec8, position1089)
													}
												}
											l1070:
												add(ruleDec, position1069)
											}
											goto l989
										l1068:
											position, tokenIndex = position989, tokenIndex989
											{
												position1098 := position
												{
													position1099, tokenIndex1099 := position, tokenIndex
													{
														position1101 := position
														{
															position1102, tokenIndex1102 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1103
															}
															position++
															goto l1102
														l1103:
															position, tokenIndex = position1102, tokenIndex1102
															if buffer[position] != rune('A') {
																goto l1100
															}
															position++
														}
													l1102:
														{
															position1104, tokenIndex1104 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1105
															}
															position++
															goto l1104
														l1105:
															position, tokenIndex = position1104, tokenIndex1104
															if buffer[position] != rune('D') {
																goto l1100
															}
															position++
														}
													l1104:
														{
															position1106, tokenIndex1106 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1107
															}
															position++
															goto l1106
														l1107:
															position, tokenIndex = position1106, tokenIndex1106
															if buffer[position] != rune('D') {
																goto l1100
															}
															position++
														}
													l1106:
														if !_rules[rulews]() {
															goto l1100
														}
														if !_rules[ruleDst16]() {
															goto l1100
														}
														if !_rules[rulesep]() {
															goto l1100
														}
														if !_rules[ruleSrc16]() {
															goto l1100
														}
														{
															add(ruleAction17, position)
														}
														add(ruleAdd16, position1101)
													}
													goto l1099
												l1100:
													position, tokenIndex = position1099, tokenIndex1099
													{
														position1110 := position
														{
															position1111, tokenIndex1111 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1112
															}
															position++
															goto l1111
														l1112:
															position, tokenIndex = position1111, tokenIndex1111
															if buffer[position] != rune('A') {
																goto l1109
															}
															position++
														}
													l1111:
														{
															position1113, tokenIndex1113 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1114
															}
															position++
															goto l1113
														l1114:
															position, tokenIndex = position1113, tokenIndex1113
															if buffer[position] != rune('D') {
																goto l1109
															}
															position++
														}
													l1113:
														{
															position1115, tokenIndex1115 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1116
															}
															position++
															goto l1115
														l1116:
															position, tokenIndex = position1115, tokenIndex1115
															if buffer[position] != rune('C') {
																goto l1109
															}
															position++
														}
													l1115:
														if !_rules[rulews]() {
															goto l1109
														}
														if !_rules[ruleDst16]() {
															goto l1109
														}
														if !_rules[rulesep]() {
															goto l1109
														}
														if !_rules[ruleSrc16]() {
															goto l1109
														}
														{
															add(ruleAction18, position)
														}
														add(ruleAdc16, position1110)
													}
													goto l1099
												l1109:
													position, tokenIndex = position1099, tokenIndex1099
													{
														position1118 := position
														{
															position1119, tokenIndex1119 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1120
															}
															position++
															goto l1119
														l1120:
															position, tokenIndex = position1119, tokenIndex1119
															if buffer[position] != rune('S') {
																goto l1097
															}
															position++
														}
													l1119:
														{
															position1121, tokenIndex1121 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1122
															}
															position++
															goto l1121
														l1122:
															position, tokenIndex = position1121, tokenIndex1121
															if buffer[position] != rune('B') {
																goto l1097
															}
															position++
														}
													l1121:
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
																goto l1097
															}
															position++
														}
													l1123:
														if !_rules[rulews]() {
															goto l1097
														}
														if !_rules[ruleDst16]() {
															goto l1097
														}
														if !_rules[rulesep]() {
															goto l1097
														}
														if !_rules[ruleSrc16]() {
															goto l1097
														}
														{
															add(ruleAction19, position)
														}
														add(ruleSbc16, position1118)
													}
												}
											l1099:
												add(ruleAlu16, position1098)
											}
											goto l989
										l1097:
											position, tokenIndex = position989, tokenIndex989
											{
												position1127 := position
												{
													position1128, tokenIndex1128 := position, tokenIndex
													{
														position1130 := position
														{
															position1131, tokenIndex1131 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1132
															}
															position++
															goto l1131
														l1132:
															position, tokenIndex = position1131, tokenIndex1131
															if buffer[position] != rune('A') {
																goto l1129
															}
															position++
														}
													l1131:
														{
															position1133, tokenIndex1133 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1134
															}
															position++
															goto l1133
														l1134:
															position, tokenIndex = position1133, tokenIndex1133
															if buffer[position] != rune('D') {
																goto l1129
															}
															position++
														}
													l1133:
														{
															position1135, tokenIndex1135 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1136
															}
															position++
															goto l1135
														l1136:
															position, tokenIndex = position1135, tokenIndex1135
															if buffer[position] != rune('D') {
																goto l1129
															}
															position++
														}
													l1135:
														if !_rules[rulews]() {
															goto l1129
														}
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
																goto l1129
															}
															position++
														}
													l1137:
														if !_rules[rulesep]() {
															goto l1129
														}
														if !_rules[ruleSrc8]() {
															goto l1129
														}
														{
															add(ruleAction43, position)
														}
														add(ruleAdd, position1130)
													}
													goto l1128
												l1129:
													position, tokenIndex = position1128, tokenIndex1128
													{
														position1141 := position
														{
															position1142, tokenIndex1142 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1143
															}
															position++
															goto l1142
														l1143:
															position, tokenIndex = position1142, tokenIndex1142
															if buffer[position] != rune('A') {
																goto l1140
															}
															position++
														}
													l1142:
														{
															position1144, tokenIndex1144 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1145
															}
															position++
															goto l1144
														l1145:
															position, tokenIndex = position1144, tokenIndex1144
															if buffer[position] != rune('D') {
																goto l1140
															}
															position++
														}
													l1144:
														{
															position1146, tokenIndex1146 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1147
															}
															position++
															goto l1146
														l1147:
															position, tokenIndex = position1146, tokenIndex1146
															if buffer[position] != rune('C') {
																goto l1140
															}
															position++
														}
													l1146:
														if !_rules[rulews]() {
															goto l1140
														}
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
																goto l1140
															}
															position++
														}
													l1148:
														if !_rules[rulesep]() {
															goto l1140
														}
														if !_rules[ruleSrc8]() {
															goto l1140
														}
														{
															add(ruleAction44, position)
														}
														add(ruleAdc, position1141)
													}
													goto l1128
												l1140:
													position, tokenIndex = position1128, tokenIndex1128
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
																goto l1151
															}
															position++
														}
													l1153:
														{
															position1155, tokenIndex1155 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1156
															}
															position++
															goto l1155
														l1156:
															position, tokenIndex = position1155, tokenIndex1155
															if buffer[position] != rune('U') {
																goto l1151
															}
															position++
														}
													l1155:
														{
															position1157, tokenIndex1157 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1158
															}
															position++
															goto l1157
														l1158:
															position, tokenIndex = position1157, tokenIndex1157
															if buffer[position] != rune('B') {
																goto l1151
															}
															position++
														}
													l1157:
														if !_rules[rulews]() {
															goto l1151
														}
														if !_rules[ruleSrc8]() {
															goto l1151
														}
														{
															add(ruleAction45, position)
														}
														add(ruleSub, position1152)
													}
													goto l1128
												l1151:
													position, tokenIndex = position1128, tokenIndex1128
													{
														switch buffer[position] {
														case 'C', 'c':
															{
																position1161 := position
																{
																	position1162, tokenIndex1162 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1163
																	}
																	position++
																	goto l1162
																l1163:
																	position, tokenIndex = position1162, tokenIndex1162
																	if buffer[position] != rune('C') {
																		goto l1126
																	}
																	position++
																}
															l1162:
																{
																	position1164, tokenIndex1164 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1165
																	}
																	position++
																	goto l1164
																l1165:
																	position, tokenIndex = position1164, tokenIndex1164
																	if buffer[position] != rune('P') {
																		goto l1126
																	}
																	position++
																}
															l1164:
																if !_rules[rulews]() {
																	goto l1126
																}
																if !_rules[ruleSrc8]() {
																	goto l1126
																}
																{
																	add(ruleAction50, position)
																}
																add(ruleCp, position1161)
															}
															break
														case 'O', 'o':
															{
																position1167 := position
																{
																	position1168, tokenIndex1168 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1169
																	}
																	position++
																	goto l1168
																l1169:
																	position, tokenIndex = position1168, tokenIndex1168
																	if buffer[position] != rune('O') {
																		goto l1126
																	}
																	position++
																}
															l1168:
																{
																	position1170, tokenIndex1170 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1171
																	}
																	position++
																	goto l1170
																l1171:
																	position, tokenIndex = position1170, tokenIndex1170
																	if buffer[position] != rune('R') {
																		goto l1126
																	}
																	position++
																}
															l1170:
																if !_rules[rulews]() {
																	goto l1126
																}
																if !_rules[ruleSrc8]() {
																	goto l1126
																}
																{
																	add(ruleAction49, position)
																}
																add(ruleOr, position1167)
															}
															break
														case 'X', 'x':
															{
																position1173 := position
																{
																	position1174, tokenIndex1174 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1175
																	}
																	position++
																	goto l1174
																l1175:
																	position, tokenIndex = position1174, tokenIndex1174
																	if buffer[position] != rune('X') {
																		goto l1126
																	}
																	position++
																}
															l1174:
																{
																	position1176, tokenIndex1176 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1177
																	}
																	position++
																	goto l1176
																l1177:
																	position, tokenIndex = position1176, tokenIndex1176
																	if buffer[position] != rune('O') {
																		goto l1126
																	}
																	position++
																}
															l1176:
																{
																	position1178, tokenIndex1178 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1179
																	}
																	position++
																	goto l1178
																l1179:
																	position, tokenIndex = position1178, tokenIndex1178
																	if buffer[position] != rune('R') {
																		goto l1126
																	}
																	position++
																}
															l1178:
																if !_rules[rulews]() {
																	goto l1126
																}
																if !_rules[ruleSrc8]() {
																	goto l1126
																}
																{
																	add(ruleAction48, position)
																}
																add(ruleXor, position1173)
															}
															break
														case 'A', 'a':
															{
																position1181 := position
																{
																	position1182, tokenIndex1182 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1183
																	}
																	position++
																	goto l1182
																l1183:
																	position, tokenIndex = position1182, tokenIndex1182
																	if buffer[position] != rune('A') {
																		goto l1126
																	}
																	position++
																}
															l1182:
																{
																	position1184, tokenIndex1184 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1185
																	}
																	position++
																	goto l1184
																l1185:
																	position, tokenIndex = position1184, tokenIndex1184
																	if buffer[position] != rune('N') {
																		goto l1126
																	}
																	position++
																}
															l1184:
																{
																	position1186, tokenIndex1186 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1187
																	}
																	position++
																	goto l1186
																l1187:
																	position, tokenIndex = position1186, tokenIndex1186
																	if buffer[position] != rune('D') {
																		goto l1126
																	}
																	position++
																}
															l1186:
																if !_rules[rulews]() {
																	goto l1126
																}
																if !_rules[ruleSrc8]() {
																	goto l1126
																}
																{
																	add(ruleAction47, position)
																}
																add(ruleAnd, position1181)
															}
															break
														default:
															{
																position1189 := position
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
																		goto l1126
																	}
																	position++
																}
															l1190:
																{
																	position1192, tokenIndex1192 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1193
																	}
																	position++
																	goto l1192
																l1193:
																	position, tokenIndex = position1192, tokenIndex1192
																	if buffer[position] != rune('B') {
																		goto l1126
																	}
																	position++
																}
															l1192:
																{
																	position1194, tokenIndex1194 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1195
																	}
																	position++
																	goto l1194
																l1195:
																	position, tokenIndex = position1194, tokenIndex1194
																	if buffer[position] != rune('C') {
																		goto l1126
																	}
																	position++
																}
															l1194:
																if !_rules[rulews]() {
																	goto l1126
																}
																{
																	position1196, tokenIndex1196 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1197
																	}
																	position++
																	goto l1196
																l1197:
																	position, tokenIndex = position1196, tokenIndex1196
																	if buffer[position] != rune('A') {
																		goto l1126
																	}
																	position++
																}
															l1196:
																if !_rules[rulesep]() {
																	goto l1126
																}
																if !_rules[ruleSrc8]() {
																	goto l1126
																}
																{
																	add(ruleAction46, position)
																}
																add(ruleSbc, position1189)
															}
															break
														}
													}

												}
											l1128:
												add(ruleAlu, position1127)
											}
											goto l989
										l1126:
											position, tokenIndex = position989, tokenIndex989
											{
												position1200 := position
												{
													position1201, tokenIndex1201 := position, tokenIndex
													{
														position1203 := position
														{
															position1204, tokenIndex1204 := position, tokenIndex
															{
																position1206 := position
																{
																	position1207, tokenIndex1207 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1208
																	}
																	position++
																	goto l1207
																l1208:
																	position, tokenIndex = position1207, tokenIndex1207
																	if buffer[position] != rune('R') {
																		goto l1205
																	}
																	position++
																}
															l1207:
																{
																	position1209, tokenIndex1209 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1210
																	}
																	position++
																	goto l1209
																l1210:
																	position, tokenIndex = position1209, tokenIndex1209
																	if buffer[position] != rune('L') {
																		goto l1205
																	}
																	position++
																}
															l1209:
																{
																	position1211, tokenIndex1211 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1212
																	}
																	position++
																	goto l1211
																l1212:
																	position, tokenIndex = position1211, tokenIndex1211
																	if buffer[position] != rune('C') {
																		goto l1205
																	}
																	position++
																}
															l1211:
																if !_rules[rulews]() {
																	goto l1205
																}
																if !_rules[ruleLoc8]() {
																	goto l1205
																}
																{
																	position1213, tokenIndex1213 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1213
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1213
																	}
																	goto l1214
																l1213:
																	position, tokenIndex = position1213, tokenIndex1213
																}
															l1214:
																{
																	add(ruleAction51, position)
																}
																add(ruleRlc, position1206)
															}
															goto l1204
														l1205:
															position, tokenIndex = position1204, tokenIndex1204
															{
																position1217 := position
																{
																	position1218, tokenIndex1218 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1219
																	}
																	position++
																	goto l1218
																l1219:
																	position, tokenIndex = position1218, tokenIndex1218
																	if buffer[position] != rune('R') {
																		goto l1216
																	}
																	position++
																}
															l1218:
																{
																	position1220, tokenIndex1220 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1221
																	}
																	position++
																	goto l1220
																l1221:
																	position, tokenIndex = position1220, tokenIndex1220
																	if buffer[position] != rune('R') {
																		goto l1216
																	}
																	position++
																}
															l1220:
																{
																	position1222, tokenIndex1222 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1223
																	}
																	position++
																	goto l1222
																l1223:
																	position, tokenIndex = position1222, tokenIndex1222
																	if buffer[position] != rune('C') {
																		goto l1216
																	}
																	position++
																}
															l1222:
																if !_rules[rulews]() {
																	goto l1216
																}
																if !_rules[ruleLoc8]() {
																	goto l1216
																}
																{
																	position1224, tokenIndex1224 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1224
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1224
																	}
																	goto l1225
																l1224:
																	position, tokenIndex = position1224, tokenIndex1224
																}
															l1225:
																{
																	add(ruleAction52, position)
																}
																add(ruleRrc, position1217)
															}
															goto l1204
														l1216:
															position, tokenIndex = position1204, tokenIndex1204
															{
																position1228 := position
																{
																	position1229, tokenIndex1229 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1230
																	}
																	position++
																	goto l1229
																l1230:
																	position, tokenIndex = position1229, tokenIndex1229
																	if buffer[position] != rune('R') {
																		goto l1227
																	}
																	position++
																}
															l1229:
																{
																	position1231, tokenIndex1231 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1232
																	}
																	position++
																	goto l1231
																l1232:
																	position, tokenIndex = position1231, tokenIndex1231
																	if buffer[position] != rune('L') {
																		goto l1227
																	}
																	position++
																}
															l1231:
																if !_rules[rulews]() {
																	goto l1227
																}
																if !_rules[ruleLoc8]() {
																	goto l1227
																}
																{
																	position1233, tokenIndex1233 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1233
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1233
																	}
																	goto l1234
																l1233:
																	position, tokenIndex = position1233, tokenIndex1233
																}
															l1234:
																{
																	add(ruleAction53, position)
																}
																add(ruleRl, position1228)
															}
															goto l1204
														l1227:
															position, tokenIndex = position1204, tokenIndex1204
															{
																position1237 := position
																{
																	position1238, tokenIndex1238 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1239
																	}
																	position++
																	goto l1238
																l1239:
																	position, tokenIndex = position1238, tokenIndex1238
																	if buffer[position] != rune('R') {
																		goto l1236
																	}
																	position++
																}
															l1238:
																{
																	position1240, tokenIndex1240 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1241
																	}
																	position++
																	goto l1240
																l1241:
																	position, tokenIndex = position1240, tokenIndex1240
																	if buffer[position] != rune('R') {
																		goto l1236
																	}
																	position++
																}
															l1240:
																if !_rules[rulews]() {
																	goto l1236
																}
																if !_rules[ruleLoc8]() {
																	goto l1236
																}
																{
																	position1242, tokenIndex1242 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1242
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1242
																	}
																	goto l1243
																l1242:
																	position, tokenIndex = position1242, tokenIndex1242
																}
															l1243:
																{
																	add(ruleAction54, position)
																}
																add(ruleRr, position1237)
															}
															goto l1204
														l1236:
															position, tokenIndex = position1204, tokenIndex1204
															{
																position1246 := position
																{
																	position1247, tokenIndex1247 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1248
																	}
																	position++
																	goto l1247
																l1248:
																	position, tokenIndex = position1247, tokenIndex1247
																	if buffer[position] != rune('S') {
																		goto l1245
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
																		goto l1245
																	}
																	position++
																}
															l1249:
																{
																	position1251, tokenIndex1251 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1252
																	}
																	position++
																	goto l1251
																l1252:
																	position, tokenIndex = position1251, tokenIndex1251
																	if buffer[position] != rune('A') {
																		goto l1245
																	}
																	position++
																}
															l1251:
																if !_rules[rulews]() {
																	goto l1245
																}
																if !_rules[ruleLoc8]() {
																	goto l1245
																}
																{
																	position1253, tokenIndex1253 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1253
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1253
																	}
																	goto l1254
																l1253:
																	position, tokenIndex = position1253, tokenIndex1253
																}
															l1254:
																{
																	add(ruleAction55, position)
																}
																add(ruleSla, position1246)
															}
															goto l1204
														l1245:
															position, tokenIndex = position1204, tokenIndex1204
															{
																position1257 := position
																{
																	position1258, tokenIndex1258 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1259
																	}
																	position++
																	goto l1258
																l1259:
																	position, tokenIndex = position1258, tokenIndex1258
																	if buffer[position] != rune('S') {
																		goto l1256
																	}
																	position++
																}
															l1258:
																{
																	position1260, tokenIndex1260 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1261
																	}
																	position++
																	goto l1260
																l1261:
																	position, tokenIndex = position1260, tokenIndex1260
																	if buffer[position] != rune('R') {
																		goto l1256
																	}
																	position++
																}
															l1260:
																{
																	position1262, tokenIndex1262 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1263
																	}
																	position++
																	goto l1262
																l1263:
																	position, tokenIndex = position1262, tokenIndex1262
																	if buffer[position] != rune('A') {
																		goto l1256
																	}
																	position++
																}
															l1262:
																if !_rules[rulews]() {
																	goto l1256
																}
																if !_rules[ruleLoc8]() {
																	goto l1256
																}
																{
																	position1264, tokenIndex1264 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1264
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1264
																	}
																	goto l1265
																l1264:
																	position, tokenIndex = position1264, tokenIndex1264
																}
															l1265:
																{
																	add(ruleAction56, position)
																}
																add(ruleSra, position1257)
															}
															goto l1204
														l1256:
															position, tokenIndex = position1204, tokenIndex1204
															{
																position1268 := position
																{
																	position1269, tokenIndex1269 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1270
																	}
																	position++
																	goto l1269
																l1270:
																	position, tokenIndex = position1269, tokenIndex1269
																	if buffer[position] != rune('S') {
																		goto l1267
																	}
																	position++
																}
															l1269:
																{
																	position1271, tokenIndex1271 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1272
																	}
																	position++
																	goto l1271
																l1272:
																	position, tokenIndex = position1271, tokenIndex1271
																	if buffer[position] != rune('L') {
																		goto l1267
																	}
																	position++
																}
															l1271:
																{
																	position1273, tokenIndex1273 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1274
																	}
																	position++
																	goto l1273
																l1274:
																	position, tokenIndex = position1273, tokenIndex1273
																	if buffer[position] != rune('L') {
																		goto l1267
																	}
																	position++
																}
															l1273:
																if !_rules[rulews]() {
																	goto l1267
																}
																if !_rules[ruleLoc8]() {
																	goto l1267
																}
																{
																	position1275, tokenIndex1275 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1275
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1275
																	}
																	goto l1276
																l1275:
																	position, tokenIndex = position1275, tokenIndex1275
																}
															l1276:
																{
																	add(ruleAction57, position)
																}
																add(ruleSll, position1268)
															}
															goto l1204
														l1267:
															position, tokenIndex = position1204, tokenIndex1204
															{
																position1278 := position
																{
																	position1279, tokenIndex1279 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1280
																	}
																	position++
																	goto l1279
																l1280:
																	position, tokenIndex = position1279, tokenIndex1279
																	if buffer[position] != rune('S') {
																		goto l1202
																	}
																	position++
																}
															l1279:
																{
																	position1281, tokenIndex1281 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1282
																	}
																	position++
																	goto l1281
																l1282:
																	position, tokenIndex = position1281, tokenIndex1281
																	if buffer[position] != rune('R') {
																		goto l1202
																	}
																	position++
																}
															l1281:
																{
																	position1283, tokenIndex1283 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1284
																	}
																	position++
																	goto l1283
																l1284:
																	position, tokenIndex = position1283, tokenIndex1283
																	if buffer[position] != rune('L') {
																		goto l1202
																	}
																	position++
																}
															l1283:
																if !_rules[rulews]() {
																	goto l1202
																}
																if !_rules[ruleLoc8]() {
																	goto l1202
																}
																{
																	position1285, tokenIndex1285 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1285
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1285
																	}
																	goto l1286
																l1285:
																	position, tokenIndex = position1285, tokenIndex1285
																}
															l1286:
																{
																	add(ruleAction58, position)
																}
																add(ruleSrl, position1278)
															}
														}
													l1204:
														add(ruleRot, position1203)
													}
													goto l1201
												l1202:
													position, tokenIndex = position1201, tokenIndex1201
													{
														switch buffer[position] {
														case 'S', 's':
															{
																position1289 := position
																{
																	position1290, tokenIndex1290 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1291
																	}
																	position++
																	goto l1290
																l1291:
																	position, tokenIndex = position1290, tokenIndex1290
																	if buffer[position] != rune('S') {
																		goto l1199
																	}
																	position++
																}
															l1290:
																{
																	position1292, tokenIndex1292 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1293
																	}
																	position++
																	goto l1292
																l1293:
																	position, tokenIndex = position1292, tokenIndex1292
																	if buffer[position] != rune('E') {
																		goto l1199
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
																		goto l1199
																	}
																	position++
																}
															l1294:
																if !_rules[rulews]() {
																	goto l1199
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1199
																}
																if !_rules[rulesep]() {
																	goto l1199
																}
																if !_rules[ruleLoc8]() {
																	goto l1199
																}
																{
																	position1296, tokenIndex1296 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1296
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1296
																	}
																	goto l1297
																l1296:
																	position, tokenIndex = position1296, tokenIndex1296
																}
															l1297:
																{
																	add(ruleAction61, position)
																}
																add(ruleSet, position1289)
															}
															break
														case 'R', 'r':
															{
																position1299 := position
																{
																	position1300, tokenIndex1300 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1301
																	}
																	position++
																	goto l1300
																l1301:
																	position, tokenIndex = position1300, tokenIndex1300
																	if buffer[position] != rune('R') {
																		goto l1199
																	}
																	position++
																}
															l1300:
																{
																	position1302, tokenIndex1302 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1303
																	}
																	position++
																	goto l1302
																l1303:
																	position, tokenIndex = position1302, tokenIndex1302
																	if buffer[position] != rune('E') {
																		goto l1199
																	}
																	position++
																}
															l1302:
																{
																	position1304, tokenIndex1304 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1305
																	}
																	position++
																	goto l1304
																l1305:
																	position, tokenIndex = position1304, tokenIndex1304
																	if buffer[position] != rune('S') {
																		goto l1199
																	}
																	position++
																}
															l1304:
																if !_rules[rulews]() {
																	goto l1199
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1199
																}
																if !_rules[rulesep]() {
																	goto l1199
																}
																if !_rules[ruleLoc8]() {
																	goto l1199
																}
																{
																	position1306, tokenIndex1306 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1306
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1306
																	}
																	goto l1307
																l1306:
																	position, tokenIndex = position1306, tokenIndex1306
																}
															l1307:
																{
																	add(ruleAction60, position)
																}
																add(ruleRes, position1299)
															}
															break
														default:
															{
																position1309 := position
																{
																	position1310, tokenIndex1310 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1311
																	}
																	position++
																	goto l1310
																l1311:
																	position, tokenIndex = position1310, tokenIndex1310
																	if buffer[position] != rune('B') {
																		goto l1199
																	}
																	position++
																}
															l1310:
																{
																	position1312, tokenIndex1312 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l1313
																	}
																	position++
																	goto l1312
																l1313:
																	position, tokenIndex = position1312, tokenIndex1312
																	if buffer[position] != rune('I') {
																		goto l1199
																	}
																	position++
																}
															l1312:
																{
																	position1314, tokenIndex1314 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1315
																	}
																	position++
																	goto l1314
																l1315:
																	position, tokenIndex = position1314, tokenIndex1314
																	if buffer[position] != rune('T') {
																		goto l1199
																	}
																	position++
																}
															l1314:
																if !_rules[rulews]() {
																	goto l1199
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1199
																}
																if !_rules[rulesep]() {
																	goto l1199
																}
																if !_rules[ruleLoc8]() {
																	goto l1199
																}
																{
																	add(ruleAction59, position)
																}
																add(ruleBit, position1309)
															}
															break
														}
													}

												}
											l1201:
												add(ruleBitOp, position1200)
											}
											goto l989
										l1199:
											position, tokenIndex = position989, tokenIndex989
											{
												position1318 := position
												{
													position1319, tokenIndex1319 := position, tokenIndex
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
																if buffer[position] != rune('e') {
																	goto l1326
																}
																position++
																goto l1325
															l1326:
																position, tokenIndex = position1325, tokenIndex1325
																if buffer[position] != rune('E') {
																	goto l1320
																}
																position++
															}
														l1325:
															{
																position1327, tokenIndex1327 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1328
																}
																position++
																goto l1327
															l1328:
																position, tokenIndex = position1327, tokenIndex1327
																if buffer[position] != rune('T') {
																	goto l1320
																}
																position++
															}
														l1327:
															{
																position1329, tokenIndex1329 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1330
																}
																position++
																goto l1329
															l1330:
																position, tokenIndex = position1329, tokenIndex1329
																if buffer[position] != rune('N') {
																	goto l1320
																}
																position++
															}
														l1329:
															add(rulePegText, position1322)
														}
														{
															add(ruleAction76, position)
														}
														add(ruleRetn, position1321)
													}
													goto l1319
												l1320:
													position, tokenIndex = position1319, tokenIndex1319
													{
														position1333 := position
														{
															position1334 := position
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
																	goto l1332
																}
																position++
															}
														l1335:
															{
																position1337, tokenIndex1337 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1338
																}
																position++
																goto l1337
															l1338:
																position, tokenIndex = position1337, tokenIndex1337
																if buffer[position] != rune('E') {
																	goto l1332
																}
																position++
															}
														l1337:
															{
																position1339, tokenIndex1339 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1340
																}
																position++
																goto l1339
															l1340:
																position, tokenIndex = position1339, tokenIndex1339
																if buffer[position] != rune('T') {
																	goto l1332
																}
																position++
															}
														l1339:
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
																	goto l1332
																}
																position++
															}
														l1341:
															add(rulePegText, position1334)
														}
														{
															add(ruleAction77, position)
														}
														add(ruleReti, position1333)
													}
													goto l1319
												l1332:
													position, tokenIndex = position1319, tokenIndex1319
													{
														position1345 := position
														{
															position1346 := position
															{
																position1347, tokenIndex1347 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1348
																}
																position++
																goto l1347
															l1348:
																position, tokenIndex = position1347, tokenIndex1347
																if buffer[position] != rune('R') {
																	goto l1344
																}
																position++
															}
														l1347:
															{
																position1349, tokenIndex1349 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1350
																}
																position++
																goto l1349
															l1350:
																position, tokenIndex = position1349, tokenIndex1349
																if buffer[position] != rune('R') {
																	goto l1344
																}
																position++
															}
														l1349:
															{
																position1351, tokenIndex1351 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1352
																}
																position++
																goto l1351
															l1352:
																position, tokenIndex = position1351, tokenIndex1351
																if buffer[position] != rune('D') {
																	goto l1344
																}
																position++
															}
														l1351:
															add(rulePegText, position1346)
														}
														{
															add(ruleAction78, position)
														}
														add(ruleRrd, position1345)
													}
													goto l1319
												l1344:
													position, tokenIndex = position1319, tokenIndex1319
													{
														position1355 := position
														{
															position1356 := position
															{
																position1357, tokenIndex1357 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1358
																}
																position++
																goto l1357
															l1358:
																position, tokenIndex = position1357, tokenIndex1357
																if buffer[position] != rune('I') {
																	goto l1354
																}
																position++
															}
														l1357:
															{
																position1359, tokenIndex1359 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1360
																}
																position++
																goto l1359
															l1360:
																position, tokenIndex = position1359, tokenIndex1359
																if buffer[position] != rune('M') {
																	goto l1354
																}
																position++
															}
														l1359:
															if buffer[position] != rune(' ') {
																goto l1354
															}
															position++
															if buffer[position] != rune('0') {
																goto l1354
															}
															position++
															add(rulePegText, position1356)
														}
														{
															add(ruleAction80, position)
														}
														add(ruleIm0, position1355)
													}
													goto l1319
												l1354:
													position, tokenIndex = position1319, tokenIndex1319
													{
														position1363 := position
														{
															position1364 := position
															{
																position1365, tokenIndex1365 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1366
																}
																position++
																goto l1365
															l1366:
																position, tokenIndex = position1365, tokenIndex1365
																if buffer[position] != rune('I') {
																	goto l1362
																}
																position++
															}
														l1365:
															{
																position1367, tokenIndex1367 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1368
																}
																position++
																goto l1367
															l1368:
																position, tokenIndex = position1367, tokenIndex1367
																if buffer[position] != rune('M') {
																	goto l1362
																}
																position++
															}
														l1367:
															if buffer[position] != rune(' ') {
																goto l1362
															}
															position++
															if buffer[position] != rune('1') {
																goto l1362
															}
															position++
															add(rulePegText, position1364)
														}
														{
															add(ruleAction81, position)
														}
														add(ruleIm1, position1363)
													}
													goto l1319
												l1362:
													position, tokenIndex = position1319, tokenIndex1319
													{
														position1371 := position
														{
															position1372 := position
															{
																position1373, tokenIndex1373 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1374
																}
																position++
																goto l1373
															l1374:
																position, tokenIndex = position1373, tokenIndex1373
																if buffer[position] != rune('I') {
																	goto l1370
																}
																position++
															}
														l1373:
															{
																position1375, tokenIndex1375 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1376
																}
																position++
																goto l1375
															l1376:
																position, tokenIndex = position1375, tokenIndex1375
																if buffer[position] != rune('M') {
																	goto l1370
																}
																position++
															}
														l1375:
															if buffer[position] != rune(' ') {
																goto l1370
															}
															position++
															if buffer[position] != rune('2') {
																goto l1370
															}
															position++
															add(rulePegText, position1372)
														}
														{
															add(ruleAction82, position)
														}
														add(ruleIm2, position1371)
													}
													goto l1319
												l1370:
													position, tokenIndex = position1319, tokenIndex1319
													{
														switch buffer[position] {
														case 'I', 'O', 'i', 'o':
															{
																position1379 := position
																{
																	position1380, tokenIndex1380 := position, tokenIndex
																	{
																		position1382 := position
																		{
																			position1383 := position
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
																					goto l1381
																				}
																				position++
																			}
																		l1384:
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
																					goto l1381
																				}
																				position++
																			}
																		l1386:
																			{
																				position1388, tokenIndex1388 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1389
																				}
																				position++
																				goto l1388
																			l1389:
																				position, tokenIndex = position1388, tokenIndex1388
																				if buffer[position] != rune('I') {
																					goto l1381
																				}
																				position++
																			}
																		l1388:
																			{
																				position1390, tokenIndex1390 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1391
																				}
																				position++
																				goto l1390
																			l1391:
																				position, tokenIndex = position1390, tokenIndex1390
																				if buffer[position] != rune('R') {
																					goto l1381
																				}
																				position++
																			}
																		l1390:
																			add(rulePegText, position1383)
																		}
																		{
																			add(ruleAction93, position)
																		}
																		add(ruleInir, position1382)
																	}
																	goto l1380
																l1381:
																	position, tokenIndex = position1380, tokenIndex1380
																	{
																		position1394 := position
																		{
																			position1395 := position
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
																					goto l1393
																				}
																				position++
																			}
																		l1396:
																			{
																				position1398, tokenIndex1398 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1399
																				}
																				position++
																				goto l1398
																			l1399:
																				position, tokenIndex = position1398, tokenIndex1398
																				if buffer[position] != rune('N') {
																					goto l1393
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
																					goto l1393
																				}
																				position++
																			}
																		l1400:
																			add(rulePegText, position1395)
																		}
																		{
																			add(ruleAction85, position)
																		}
																		add(ruleIni, position1394)
																	}
																	goto l1380
																l1393:
																	position, tokenIndex = position1380, tokenIndex1380
																	{
																		position1404 := position
																		{
																			position1405 := position
																			{
																				position1406, tokenIndex1406 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1407
																				}
																				position++
																				goto l1406
																			l1407:
																				position, tokenIndex = position1406, tokenIndex1406
																				if buffer[position] != rune('O') {
																					goto l1403
																				}
																				position++
																			}
																		l1406:
																			{
																				position1408, tokenIndex1408 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1409
																				}
																				position++
																				goto l1408
																			l1409:
																				position, tokenIndex = position1408, tokenIndex1408
																				if buffer[position] != rune('T') {
																					goto l1403
																				}
																				position++
																			}
																		l1408:
																			{
																				position1410, tokenIndex1410 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1411
																				}
																				position++
																				goto l1410
																			l1411:
																				position, tokenIndex = position1410, tokenIndex1410
																				if buffer[position] != rune('I') {
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
																			add(ruleAction94, position)
																		}
																		add(ruleOtir, position1404)
																	}
																	goto l1380
																l1403:
																	position, tokenIndex = position1380, tokenIndex1380
																	{
																		position1416 := position
																		{
																			position1417 := position
																			{
																				position1418, tokenIndex1418 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1419
																				}
																				position++
																				goto l1418
																			l1419:
																				position, tokenIndex = position1418, tokenIndex1418
																				if buffer[position] != rune('O') {
																					goto l1415
																				}
																				position++
																			}
																		l1418:
																			{
																				position1420, tokenIndex1420 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1421
																				}
																				position++
																				goto l1420
																			l1421:
																				position, tokenIndex = position1420, tokenIndex1420
																				if buffer[position] != rune('U') {
																					goto l1415
																				}
																				position++
																			}
																		l1420:
																			{
																				position1422, tokenIndex1422 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1423
																				}
																				position++
																				goto l1422
																			l1423:
																				position, tokenIndex = position1422, tokenIndex1422
																				if buffer[position] != rune('T') {
																					goto l1415
																				}
																				position++
																			}
																		l1422:
																			{
																				position1424, tokenIndex1424 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1425
																				}
																				position++
																				goto l1424
																			l1425:
																				position, tokenIndex = position1424, tokenIndex1424
																				if buffer[position] != rune('I') {
																					goto l1415
																				}
																				position++
																			}
																		l1424:
																			add(rulePegText, position1417)
																		}
																		{
																			add(ruleAction86, position)
																		}
																		add(ruleOuti, position1416)
																	}
																	goto l1380
																l1415:
																	position, tokenIndex = position1380, tokenIndex1380
																	{
																		position1428 := position
																		{
																			position1429 := position
																			{
																				position1430, tokenIndex1430 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1431
																				}
																				position++
																				goto l1430
																			l1431:
																				position, tokenIndex = position1430, tokenIndex1430
																				if buffer[position] != rune('I') {
																					goto l1427
																				}
																				position++
																			}
																		l1430:
																			{
																				position1432, tokenIndex1432 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1433
																				}
																				position++
																				goto l1432
																			l1433:
																				position, tokenIndex = position1432, tokenIndex1432
																				if buffer[position] != rune('N') {
																					goto l1427
																				}
																				position++
																			}
																		l1432:
																			{
																				position1434, tokenIndex1434 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1435
																				}
																				position++
																				goto l1434
																			l1435:
																				position, tokenIndex = position1434, tokenIndex1434
																				if buffer[position] != rune('D') {
																					goto l1427
																				}
																				position++
																			}
																		l1434:
																			{
																				position1436, tokenIndex1436 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1437
																				}
																				position++
																				goto l1436
																			l1437:
																				position, tokenIndex = position1436, tokenIndex1436
																				if buffer[position] != rune('R') {
																					goto l1427
																				}
																				position++
																			}
																		l1436:
																			add(rulePegText, position1429)
																		}
																		{
																			add(ruleAction97, position)
																		}
																		add(ruleIndr, position1428)
																	}
																	goto l1380
																l1427:
																	position, tokenIndex = position1380, tokenIndex1380
																	{
																		position1440 := position
																		{
																			position1441 := position
																			{
																				position1442, tokenIndex1442 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1443
																				}
																				position++
																				goto l1442
																			l1443:
																				position, tokenIndex = position1442, tokenIndex1442
																				if buffer[position] != rune('I') {
																					goto l1439
																				}
																				position++
																			}
																		l1442:
																			{
																				position1444, tokenIndex1444 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1445
																				}
																				position++
																				goto l1444
																			l1445:
																				position, tokenIndex = position1444, tokenIndex1444
																				if buffer[position] != rune('N') {
																					goto l1439
																				}
																				position++
																			}
																		l1444:
																			{
																				position1446, tokenIndex1446 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1447
																				}
																				position++
																				goto l1446
																			l1447:
																				position, tokenIndex = position1446, tokenIndex1446
																				if buffer[position] != rune('D') {
																					goto l1439
																				}
																				position++
																			}
																		l1446:
																			add(rulePegText, position1441)
																		}
																		{
																			add(ruleAction89, position)
																		}
																		add(ruleInd, position1440)
																	}
																	goto l1380
																l1439:
																	position, tokenIndex = position1380, tokenIndex1380
																	{
																		position1450 := position
																		{
																			position1451 := position
																			{
																				position1452, tokenIndex1452 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1453
																				}
																				position++
																				goto l1452
																			l1453:
																				position, tokenIndex = position1452, tokenIndex1452
																				if buffer[position] != rune('O') {
																					goto l1449
																				}
																				position++
																			}
																		l1452:
																			{
																				position1454, tokenIndex1454 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1455
																				}
																				position++
																				goto l1454
																			l1455:
																				position, tokenIndex = position1454, tokenIndex1454
																				if buffer[position] != rune('T') {
																					goto l1449
																				}
																				position++
																			}
																		l1454:
																			{
																				position1456, tokenIndex1456 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1457
																				}
																				position++
																				goto l1456
																			l1457:
																				position, tokenIndex = position1456, tokenIndex1456
																				if buffer[position] != rune('D') {
																					goto l1449
																				}
																				position++
																			}
																		l1456:
																			{
																				position1458, tokenIndex1458 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1459
																				}
																				position++
																				goto l1458
																			l1459:
																				position, tokenIndex = position1458, tokenIndex1458
																				if buffer[position] != rune('R') {
																					goto l1449
																				}
																				position++
																			}
																		l1458:
																			add(rulePegText, position1451)
																		}
																		{
																			add(ruleAction98, position)
																		}
																		add(ruleOtdr, position1450)
																	}
																	goto l1380
																l1449:
																	position, tokenIndex = position1380, tokenIndex1380
																	{
																		position1461 := position
																		{
																			position1462 := position
																			{
																				position1463, tokenIndex1463 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1464
																				}
																				position++
																				goto l1463
																			l1464:
																				position, tokenIndex = position1463, tokenIndex1463
																				if buffer[position] != rune('O') {
																					goto l1317
																				}
																				position++
																			}
																		l1463:
																			{
																				position1465, tokenIndex1465 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1466
																				}
																				position++
																				goto l1465
																			l1466:
																				position, tokenIndex = position1465, tokenIndex1465
																				if buffer[position] != rune('U') {
																					goto l1317
																				}
																				position++
																			}
																		l1465:
																			{
																				position1467, tokenIndex1467 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1468
																				}
																				position++
																				goto l1467
																			l1468:
																				position, tokenIndex = position1467, tokenIndex1467
																				if buffer[position] != rune('T') {
																					goto l1317
																				}
																				position++
																			}
																		l1467:
																			{
																				position1469, tokenIndex1469 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1470
																				}
																				position++
																				goto l1469
																			l1470:
																				position, tokenIndex = position1469, tokenIndex1469
																				if buffer[position] != rune('D') {
																					goto l1317
																				}
																				position++
																			}
																		l1469:
																			add(rulePegText, position1462)
																		}
																		{
																			add(ruleAction90, position)
																		}
																		add(ruleOutd, position1461)
																	}
																}
															l1380:
																add(ruleBlitIO, position1379)
															}
															break
														case 'R', 'r':
															{
																position1472 := position
																{
																	position1473 := position
																	{
																		position1474, tokenIndex1474 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1475
																		}
																		position++
																		goto l1474
																	l1475:
																		position, tokenIndex = position1474, tokenIndex1474
																		if buffer[position] != rune('R') {
																			goto l1317
																		}
																		position++
																	}
																l1474:
																	{
																		position1476, tokenIndex1476 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1477
																		}
																		position++
																		goto l1476
																	l1477:
																		position, tokenIndex = position1476, tokenIndex1476
																		if buffer[position] != rune('L') {
																			goto l1317
																		}
																		position++
																	}
																l1476:
																	{
																		position1478, tokenIndex1478 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1479
																		}
																		position++
																		goto l1478
																	l1479:
																		position, tokenIndex = position1478, tokenIndex1478
																		if buffer[position] != rune('D') {
																			goto l1317
																		}
																		position++
																	}
																l1478:
																	add(rulePegText, position1473)
																}
																{
																	add(ruleAction79, position)
																}
																add(ruleRld, position1472)
															}
															break
														case 'N', 'n':
															{
																position1481 := position
																{
																	position1482 := position
																	{
																		position1483, tokenIndex1483 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1484
																		}
																		position++
																		goto l1483
																	l1484:
																		position, tokenIndex = position1483, tokenIndex1483
																		if buffer[position] != rune('N') {
																			goto l1317
																		}
																		position++
																	}
																l1483:
																	{
																		position1485, tokenIndex1485 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1486
																		}
																		position++
																		goto l1485
																	l1486:
																		position, tokenIndex = position1485, tokenIndex1485
																		if buffer[position] != rune('E') {
																			goto l1317
																		}
																		position++
																	}
																l1485:
																	{
																		position1487, tokenIndex1487 := position, tokenIndex
																		if buffer[position] != rune('g') {
																			goto l1488
																		}
																		position++
																		goto l1487
																	l1488:
																		position, tokenIndex = position1487, tokenIndex1487
																		if buffer[position] != rune('G') {
																			goto l1317
																		}
																		position++
																	}
																l1487:
																	add(rulePegText, position1482)
																}
																{
																	add(ruleAction75, position)
																}
																add(ruleNeg, position1481)
															}
															break
														default:
															{
																position1490 := position
																{
																	position1491, tokenIndex1491 := position, tokenIndex
																	{
																		position1493 := position
																		{
																			position1494 := position
																			{
																				position1495, tokenIndex1495 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1496
																				}
																				position++
																				goto l1495
																			l1496:
																				position, tokenIndex = position1495, tokenIndex1495
																				if buffer[position] != rune('L') {
																					goto l1492
																				}
																				position++
																			}
																		l1495:
																			{
																				position1497, tokenIndex1497 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1498
																				}
																				position++
																				goto l1497
																			l1498:
																				position, tokenIndex = position1497, tokenIndex1497
																				if buffer[position] != rune('D') {
																					goto l1492
																				}
																				position++
																			}
																		l1497:
																			{
																				position1499, tokenIndex1499 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1500
																				}
																				position++
																				goto l1499
																			l1500:
																				position, tokenIndex = position1499, tokenIndex1499
																				if buffer[position] != rune('I') {
																					goto l1492
																				}
																				position++
																			}
																		l1499:
																			{
																				position1501, tokenIndex1501 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1502
																				}
																				position++
																				goto l1501
																			l1502:
																				position, tokenIndex = position1501, tokenIndex1501
																				if buffer[position] != rune('R') {
																					goto l1492
																				}
																				position++
																			}
																		l1501:
																			add(rulePegText, position1494)
																		}
																		{
																			add(ruleAction91, position)
																		}
																		add(ruleLdir, position1493)
																	}
																	goto l1491
																l1492:
																	position, tokenIndex = position1491, tokenIndex1491
																	{
																		position1505 := position
																		{
																			position1506 := position
																			{
																				position1507, tokenIndex1507 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1508
																				}
																				position++
																				goto l1507
																			l1508:
																				position, tokenIndex = position1507, tokenIndex1507
																				if buffer[position] != rune('L') {
																					goto l1504
																				}
																				position++
																			}
																		l1507:
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
																					goto l1504
																				}
																				position++
																			}
																		l1509:
																			{
																				position1511, tokenIndex1511 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1512
																				}
																				position++
																				goto l1511
																			l1512:
																				position, tokenIndex = position1511, tokenIndex1511
																				if buffer[position] != rune('I') {
																					goto l1504
																				}
																				position++
																			}
																		l1511:
																			add(rulePegText, position1506)
																		}
																		{
																			add(ruleAction83, position)
																		}
																		add(ruleLdi, position1505)
																	}
																	goto l1491
																l1504:
																	position, tokenIndex = position1491, tokenIndex1491
																	{
																		position1515 := position
																		{
																			position1516 := position
																			{
																				position1517, tokenIndex1517 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1518
																				}
																				position++
																				goto l1517
																			l1518:
																				position, tokenIndex = position1517, tokenIndex1517
																				if buffer[position] != rune('C') {
																					goto l1514
																				}
																				position++
																			}
																		l1517:
																			{
																				position1519, tokenIndex1519 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1520
																				}
																				position++
																				goto l1519
																			l1520:
																				position, tokenIndex = position1519, tokenIndex1519
																				if buffer[position] != rune('P') {
																					goto l1514
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
																					goto l1514
																				}
																				position++
																			}
																		l1521:
																			{
																				position1523, tokenIndex1523 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1524
																				}
																				position++
																				goto l1523
																			l1524:
																				position, tokenIndex = position1523, tokenIndex1523
																				if buffer[position] != rune('R') {
																					goto l1514
																				}
																				position++
																			}
																		l1523:
																			add(rulePegText, position1516)
																		}
																		{
																			add(ruleAction92, position)
																		}
																		add(ruleCpir, position1515)
																	}
																	goto l1491
																l1514:
																	position, tokenIndex = position1491, tokenIndex1491
																	{
																		position1527 := position
																		{
																			position1528 := position
																			{
																				position1529, tokenIndex1529 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1530
																				}
																				position++
																				goto l1529
																			l1530:
																				position, tokenIndex = position1529, tokenIndex1529
																				if buffer[position] != rune('C') {
																					goto l1526
																				}
																				position++
																			}
																		l1529:
																			{
																				position1531, tokenIndex1531 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1532
																				}
																				position++
																				goto l1531
																			l1532:
																				position, tokenIndex = position1531, tokenIndex1531
																				if buffer[position] != rune('P') {
																					goto l1526
																				}
																				position++
																			}
																		l1531:
																			{
																				position1533, tokenIndex1533 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1534
																				}
																				position++
																				goto l1533
																			l1534:
																				position, tokenIndex = position1533, tokenIndex1533
																				if buffer[position] != rune('I') {
																					goto l1526
																				}
																				position++
																			}
																		l1533:
																			add(rulePegText, position1528)
																		}
																		{
																			add(ruleAction84, position)
																		}
																		add(ruleCpi, position1527)
																	}
																	goto l1491
																l1526:
																	position, tokenIndex = position1491, tokenIndex1491
																	{
																		position1537 := position
																		{
																			position1538 := position
																			{
																				position1539, tokenIndex1539 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1540
																				}
																				position++
																				goto l1539
																			l1540:
																				position, tokenIndex = position1539, tokenIndex1539
																				if buffer[position] != rune('L') {
																					goto l1536
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
																					goto l1536
																				}
																				position++
																			}
																		l1541:
																			{
																				position1543, tokenIndex1543 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1544
																				}
																				position++
																				goto l1543
																			l1544:
																				position, tokenIndex = position1543, tokenIndex1543
																				if buffer[position] != rune('D') {
																					goto l1536
																				}
																				position++
																			}
																		l1543:
																			{
																				position1545, tokenIndex1545 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1546
																				}
																				position++
																				goto l1545
																			l1546:
																				position, tokenIndex = position1545, tokenIndex1545
																				if buffer[position] != rune('R') {
																					goto l1536
																				}
																				position++
																			}
																		l1545:
																			add(rulePegText, position1538)
																		}
																		{
																			add(ruleAction95, position)
																		}
																		add(ruleLddr, position1537)
																	}
																	goto l1491
																l1536:
																	position, tokenIndex = position1491, tokenIndex1491
																	{
																		position1549 := position
																		{
																			position1550 := position
																			{
																				position1551, tokenIndex1551 := position, tokenIndex
																				if buffer[position] != rune('l') {
																					goto l1552
																				}
																				position++
																				goto l1551
																			l1552:
																				position, tokenIndex = position1551, tokenIndex1551
																				if buffer[position] != rune('L') {
																					goto l1548
																				}
																				position++
																			}
																		l1551:
																			{
																				position1553, tokenIndex1553 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1554
																				}
																				position++
																				goto l1553
																			l1554:
																				position, tokenIndex = position1553, tokenIndex1553
																				if buffer[position] != rune('D') {
																					goto l1548
																				}
																				position++
																			}
																		l1553:
																			{
																				position1555, tokenIndex1555 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1556
																				}
																				position++
																				goto l1555
																			l1556:
																				position, tokenIndex = position1555, tokenIndex1555
																				if buffer[position] != rune('D') {
																					goto l1548
																				}
																				position++
																			}
																		l1555:
																			add(rulePegText, position1550)
																		}
																		{
																			add(ruleAction87, position)
																		}
																		add(ruleLdd, position1549)
																	}
																	goto l1491
																l1548:
																	position, tokenIndex = position1491, tokenIndex1491
																	{
																		position1559 := position
																		{
																			position1560 := position
																			{
																				position1561, tokenIndex1561 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1562
																				}
																				position++
																				goto l1561
																			l1562:
																				position, tokenIndex = position1561, tokenIndex1561
																				if buffer[position] != rune('C') {
																					goto l1558
																				}
																				position++
																			}
																		l1561:
																			{
																				position1563, tokenIndex1563 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1564
																				}
																				position++
																				goto l1563
																			l1564:
																				position, tokenIndex = position1563, tokenIndex1563
																				if buffer[position] != rune('P') {
																					goto l1558
																				}
																				position++
																			}
																		l1563:
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
																					goto l1558
																				}
																				position++
																			}
																		l1565:
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
																					goto l1558
																				}
																				position++
																			}
																		l1567:
																			add(rulePegText, position1560)
																		}
																		{
																			add(ruleAction96, position)
																		}
																		add(ruleCpdr, position1559)
																	}
																	goto l1491
																l1558:
																	position, tokenIndex = position1491, tokenIndex1491
																	{
																		position1570 := position
																		{
																			position1571 := position
																			{
																				position1572, tokenIndex1572 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1573
																				}
																				position++
																				goto l1572
																			l1573:
																				position, tokenIndex = position1572, tokenIndex1572
																				if buffer[position] != rune('C') {
																					goto l1317
																				}
																				position++
																			}
																		l1572:
																			{
																				position1574, tokenIndex1574 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1575
																				}
																				position++
																				goto l1574
																			l1575:
																				position, tokenIndex = position1574, tokenIndex1574
																				if buffer[position] != rune('P') {
																					goto l1317
																				}
																				position++
																			}
																		l1574:
																			{
																				position1576, tokenIndex1576 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1577
																				}
																				position++
																				goto l1576
																			l1577:
																				position, tokenIndex = position1576, tokenIndex1576
																				if buffer[position] != rune('D') {
																					goto l1317
																				}
																				position++
																			}
																		l1576:
																			add(rulePegText, position1571)
																		}
																		{
																			add(ruleAction88, position)
																		}
																		add(ruleCpd, position1570)
																	}
																}
															l1491:
																add(ruleBlit, position1490)
															}
															break
														}
													}

												}
											l1319:
												add(ruleEDSimple, position1318)
											}
											goto l989
										l1317:
											position, tokenIndex = position989, tokenIndex989
											{
												position1580 := position
												{
													position1581, tokenIndex1581 := position, tokenIndex
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
																if buffer[position] != rune('c') {
																	goto l1590
																}
																position++
																goto l1589
															l1590:
																position, tokenIndex = position1589, tokenIndex1589
																if buffer[position] != rune('C') {
																	goto l1582
																}
																position++
															}
														l1589:
															{
																position1591, tokenIndex1591 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1592
																}
																position++
																goto l1591
															l1592:
																position, tokenIndex = position1591, tokenIndex1591
																if buffer[position] != rune('A') {
																	goto l1582
																}
																position++
															}
														l1591:
															add(rulePegText, position1584)
														}
														{
															add(ruleAction64, position)
														}
														add(ruleRlca, position1583)
													}
													goto l1581
												l1582:
													position, tokenIndex = position1581, tokenIndex1581
													{
														position1595 := position
														{
															position1596 := position
															{
																position1597, tokenIndex1597 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1598
																}
																position++
																goto l1597
															l1598:
																position, tokenIndex = position1597, tokenIndex1597
																if buffer[position] != rune('R') {
																	goto l1594
																}
																position++
															}
														l1597:
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
																	goto l1594
																}
																position++
															}
														l1599:
															{
																position1601, tokenIndex1601 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1602
																}
																position++
																goto l1601
															l1602:
																position, tokenIndex = position1601, tokenIndex1601
																if buffer[position] != rune('C') {
																	goto l1594
																}
																position++
															}
														l1601:
															{
																position1603, tokenIndex1603 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1604
																}
																position++
																goto l1603
															l1604:
																position, tokenIndex = position1603, tokenIndex1603
																if buffer[position] != rune('A') {
																	goto l1594
																}
																position++
															}
														l1603:
															add(rulePegText, position1596)
														}
														{
															add(ruleAction65, position)
														}
														add(ruleRrca, position1595)
													}
													goto l1581
												l1594:
													position, tokenIndex = position1581, tokenIndex1581
													{
														position1607 := position
														{
															position1608 := position
															{
																position1609, tokenIndex1609 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1610
																}
																position++
																goto l1609
															l1610:
																position, tokenIndex = position1609, tokenIndex1609
																if buffer[position] != rune('R') {
																	goto l1606
																}
																position++
															}
														l1609:
															{
																position1611, tokenIndex1611 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1612
																}
																position++
																goto l1611
															l1612:
																position, tokenIndex = position1611, tokenIndex1611
																if buffer[position] != rune('L') {
																	goto l1606
																}
																position++
															}
														l1611:
															{
																position1613, tokenIndex1613 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1614
																}
																position++
																goto l1613
															l1614:
																position, tokenIndex = position1613, tokenIndex1613
																if buffer[position] != rune('A') {
																	goto l1606
																}
																position++
															}
														l1613:
															add(rulePegText, position1608)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleRla, position1607)
													}
													goto l1581
												l1606:
													position, tokenIndex = position1581, tokenIndex1581
													{
														position1617 := position
														{
															position1618 := position
															{
																position1619, tokenIndex1619 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1620
																}
																position++
																goto l1619
															l1620:
																position, tokenIndex = position1619, tokenIndex1619
																if buffer[position] != rune('D') {
																	goto l1616
																}
																position++
															}
														l1619:
															{
																position1621, tokenIndex1621 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1622
																}
																position++
																goto l1621
															l1622:
																position, tokenIndex = position1621, tokenIndex1621
																if buffer[position] != rune('A') {
																	goto l1616
																}
																position++
															}
														l1621:
															{
																position1623, tokenIndex1623 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1624
																}
																position++
																goto l1623
															l1624:
																position, tokenIndex = position1623, tokenIndex1623
																if buffer[position] != rune('A') {
																	goto l1616
																}
																position++
															}
														l1623:
															add(rulePegText, position1618)
														}
														{
															add(ruleAction68, position)
														}
														add(ruleDaa, position1617)
													}
													goto l1581
												l1616:
													position, tokenIndex = position1581, tokenIndex1581
													{
														position1627 := position
														{
															position1628 := position
															{
																position1629, tokenIndex1629 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1630
																}
																position++
																goto l1629
															l1630:
																position, tokenIndex = position1629, tokenIndex1629
																if buffer[position] != rune('C') {
																	goto l1626
																}
																position++
															}
														l1629:
															{
																position1631, tokenIndex1631 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1632
																}
																position++
																goto l1631
															l1632:
																position, tokenIndex = position1631, tokenIndex1631
																if buffer[position] != rune('P') {
																	goto l1626
																}
																position++
															}
														l1631:
															{
																position1633, tokenIndex1633 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1634
																}
																position++
																goto l1633
															l1634:
																position, tokenIndex = position1633, tokenIndex1633
																if buffer[position] != rune('L') {
																	goto l1626
																}
																position++
															}
														l1633:
															add(rulePegText, position1628)
														}
														{
															add(ruleAction69, position)
														}
														add(ruleCpl, position1627)
													}
													goto l1581
												l1626:
													position, tokenIndex = position1581, tokenIndex1581
													{
														position1637 := position
														{
															position1638 := position
															{
																position1639, tokenIndex1639 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1640
																}
																position++
																goto l1639
															l1640:
																position, tokenIndex = position1639, tokenIndex1639
																if buffer[position] != rune('E') {
																	goto l1636
																}
																position++
															}
														l1639:
															{
																position1641, tokenIndex1641 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1642
																}
																position++
																goto l1641
															l1642:
																position, tokenIndex = position1641, tokenIndex1641
																if buffer[position] != rune('X') {
																	goto l1636
																}
																position++
															}
														l1641:
															{
																position1643, tokenIndex1643 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1644
																}
																position++
																goto l1643
															l1644:
																position, tokenIndex = position1643, tokenIndex1643
																if buffer[position] != rune('X') {
																	goto l1636
																}
																position++
															}
														l1643:
															add(rulePegText, position1638)
														}
														{
															add(ruleAction72, position)
														}
														add(ruleExx, position1637)
													}
													goto l1581
												l1636:
													position, tokenIndex = position1581, tokenIndex1581
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1647 := position
																{
																	position1648 := position
																	{
																		position1649, tokenIndex1649 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1650
																		}
																		position++
																		goto l1649
																	l1650:
																		position, tokenIndex = position1649, tokenIndex1649
																		if buffer[position] != rune('E') {
																			goto l1579
																		}
																		position++
																	}
																l1649:
																	{
																		position1651, tokenIndex1651 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1652
																		}
																		position++
																		goto l1651
																	l1652:
																		position, tokenIndex = position1651, tokenIndex1651
																		if buffer[position] != rune('I') {
																			goto l1579
																		}
																		position++
																	}
																l1651:
																	add(rulePegText, position1648)
																}
																{
																	add(ruleAction74, position)
																}
																add(ruleEi, position1647)
															}
															break
														case 'D', 'd':
															{
																position1654 := position
																{
																	position1655 := position
																	{
																		position1656, tokenIndex1656 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1657
																		}
																		position++
																		goto l1656
																	l1657:
																		position, tokenIndex = position1656, tokenIndex1656
																		if buffer[position] != rune('D') {
																			goto l1579
																		}
																		position++
																	}
																l1656:
																	{
																		position1658, tokenIndex1658 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1659
																		}
																		position++
																		goto l1658
																	l1659:
																		position, tokenIndex = position1658, tokenIndex1658
																		if buffer[position] != rune('I') {
																			goto l1579
																		}
																		position++
																	}
																l1658:
																	add(rulePegText, position1655)
																}
																{
																	add(ruleAction73, position)
																}
																add(ruleDi, position1654)
															}
															break
														case 'C', 'c':
															{
																position1661 := position
																{
																	position1662 := position
																	{
																		position1663, tokenIndex1663 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1664
																		}
																		position++
																		goto l1663
																	l1664:
																		position, tokenIndex = position1663, tokenIndex1663
																		if buffer[position] != rune('C') {
																			goto l1579
																		}
																		position++
																	}
																l1663:
																	{
																		position1665, tokenIndex1665 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1666
																		}
																		position++
																		goto l1665
																	l1666:
																		position, tokenIndex = position1665, tokenIndex1665
																		if buffer[position] != rune('C') {
																			goto l1579
																		}
																		position++
																	}
																l1665:
																	{
																		position1667, tokenIndex1667 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1668
																		}
																		position++
																		goto l1667
																	l1668:
																		position, tokenIndex = position1667, tokenIndex1667
																		if buffer[position] != rune('F') {
																			goto l1579
																		}
																		position++
																	}
																l1667:
																	add(rulePegText, position1662)
																}
																{
																	add(ruleAction71, position)
																}
																add(ruleCcf, position1661)
															}
															break
														case 'S', 's':
															{
																position1670 := position
																{
																	position1671 := position
																	{
																		position1672, tokenIndex1672 := position, tokenIndex
																		if buffer[position] != rune('s') {
																			goto l1673
																		}
																		position++
																		goto l1672
																	l1673:
																		position, tokenIndex = position1672, tokenIndex1672
																		if buffer[position] != rune('S') {
																			goto l1579
																		}
																		position++
																	}
																l1672:
																	{
																		position1674, tokenIndex1674 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1675
																		}
																		position++
																		goto l1674
																	l1675:
																		position, tokenIndex = position1674, tokenIndex1674
																		if buffer[position] != rune('C') {
																			goto l1579
																		}
																		position++
																	}
																l1674:
																	{
																		position1676, tokenIndex1676 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1677
																		}
																		position++
																		goto l1676
																	l1677:
																		position, tokenIndex = position1676, tokenIndex1676
																		if buffer[position] != rune('F') {
																			goto l1579
																		}
																		position++
																	}
																l1676:
																	add(rulePegText, position1671)
																}
																{
																	add(ruleAction70, position)
																}
																add(ruleScf, position1670)
															}
															break
														case 'R', 'r':
															{
																position1679 := position
																{
																	position1680 := position
																	{
																		position1681, tokenIndex1681 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1682
																		}
																		position++
																		goto l1681
																	l1682:
																		position, tokenIndex = position1681, tokenIndex1681
																		if buffer[position] != rune('R') {
																			goto l1579
																		}
																		position++
																	}
																l1681:
																	{
																		position1683, tokenIndex1683 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1684
																		}
																		position++
																		goto l1683
																	l1684:
																		position, tokenIndex = position1683, tokenIndex1683
																		if buffer[position] != rune('R') {
																			goto l1579
																		}
																		position++
																	}
																l1683:
																	{
																		position1685, tokenIndex1685 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1686
																		}
																		position++
																		goto l1685
																	l1686:
																		position, tokenIndex = position1685, tokenIndex1685
																		if buffer[position] != rune('A') {
																			goto l1579
																		}
																		position++
																	}
																l1685:
																	add(rulePegText, position1680)
																}
																{
																	add(ruleAction67, position)
																}
																add(ruleRra, position1679)
															}
															break
														case 'H', 'h':
															{
																position1688 := position
																{
																	position1689 := position
																	{
																		position1690, tokenIndex1690 := position, tokenIndex
																		if buffer[position] != rune('h') {
																			goto l1691
																		}
																		position++
																		goto l1690
																	l1691:
																		position, tokenIndex = position1690, tokenIndex1690
																		if buffer[position] != rune('H') {
																			goto l1579
																		}
																		position++
																	}
																l1690:
																	{
																		position1692, tokenIndex1692 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1693
																		}
																		position++
																		goto l1692
																	l1693:
																		position, tokenIndex = position1692, tokenIndex1692
																		if buffer[position] != rune('A') {
																			goto l1579
																		}
																		position++
																	}
																l1692:
																	{
																		position1694, tokenIndex1694 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1695
																		}
																		position++
																		goto l1694
																	l1695:
																		position, tokenIndex = position1694, tokenIndex1694
																		if buffer[position] != rune('L') {
																			goto l1579
																		}
																		position++
																	}
																l1694:
																	{
																		position1696, tokenIndex1696 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1697
																		}
																		position++
																		goto l1696
																	l1697:
																		position, tokenIndex = position1696, tokenIndex1696
																		if buffer[position] != rune('T') {
																			goto l1579
																		}
																		position++
																	}
																l1696:
																	add(rulePegText, position1689)
																}
																{
																	add(ruleAction63, position)
																}
																add(ruleHalt, position1688)
															}
															break
														default:
															{
																position1699 := position
																{
																	position1700 := position
																	{
																		position1701, tokenIndex1701 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1702
																		}
																		position++
																		goto l1701
																	l1702:
																		position, tokenIndex = position1701, tokenIndex1701
																		if buffer[position] != rune('N') {
																			goto l1579
																		}
																		position++
																	}
																l1701:
																	{
																		position1703, tokenIndex1703 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1704
																		}
																		position++
																		goto l1703
																	l1704:
																		position, tokenIndex = position1703, tokenIndex1703
																		if buffer[position] != rune('O') {
																			goto l1579
																		}
																		position++
																	}
																l1703:
																	{
																		position1705, tokenIndex1705 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1706
																		}
																		position++
																		goto l1705
																	l1706:
																		position, tokenIndex = position1705, tokenIndex1705
																		if buffer[position] != rune('P') {
																			goto l1579
																		}
																		position++
																	}
																l1705:
																	add(rulePegText, position1700)
																}
																{
																	add(ruleAction62, position)
																}
																add(ruleNop, position1699)
															}
															break
														}
													}

												}
											l1581:
												add(ruleSimple, position1580)
											}
											goto l989
										l1579:
											position, tokenIndex = position989, tokenIndex989
											{
												position1709 := position
												{
													position1710, tokenIndex1710 := position, tokenIndex
													{
														position1712 := position
														{
															position1713, tokenIndex1713 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1714
															}
															position++
															goto l1713
														l1714:
															position, tokenIndex = position1713, tokenIndex1713
															if buffer[position] != rune('R') {
																goto l1711
															}
															position++
														}
													l1713:
														{
															position1715, tokenIndex1715 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1716
															}
															position++
															goto l1715
														l1716:
															position, tokenIndex = position1715, tokenIndex1715
															if buffer[position] != rune('S') {
																goto l1711
															}
															position++
														}
													l1715:
														{
															position1717, tokenIndex1717 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1718
															}
															position++
															goto l1717
														l1718:
															position, tokenIndex = position1717, tokenIndex1717
															if buffer[position] != rune('T') {
																goto l1711
															}
															position++
														}
													l1717:
														if !_rules[rulews]() {
															goto l1711
														}
														if !_rules[rulen]() {
															goto l1711
														}
														{
															add(ruleAction99, position)
														}
														add(ruleRst, position1712)
													}
													goto l1710
												l1711:
													position, tokenIndex = position1710, tokenIndex1710
													{
														position1721 := position
														{
															position1722, tokenIndex1722 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1723
															}
															position++
															goto l1722
														l1723:
															position, tokenIndex = position1722, tokenIndex1722
															if buffer[position] != rune('J') {
																goto l1720
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
																goto l1720
															}
															position++
														}
													l1724:
														if !_rules[rulews]() {
															goto l1720
														}
														{
															position1726, tokenIndex1726 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1726
															}
															if !_rules[rulesep]() {
																goto l1726
															}
															goto l1727
														l1726:
															position, tokenIndex = position1726, tokenIndex1726
														}
													l1727:
														if !_rules[ruleSrc16]() {
															goto l1720
														}
														{
															add(ruleAction102, position)
														}
														add(ruleJp, position1721)
													}
													goto l1710
												l1720:
													position, tokenIndex = position1710, tokenIndex1710
													{
														switch buffer[position] {
														case 'D', 'd':
															{
																position1730 := position
																{
																	position1731, tokenIndex1731 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1732
																	}
																	position++
																	goto l1731
																l1732:
																	position, tokenIndex = position1731, tokenIndex1731
																	if buffer[position] != rune('D') {
																		goto l1708
																	}
																	position++
																}
															l1731:
																{
																	position1733, tokenIndex1733 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1734
																	}
																	position++
																	goto l1733
																l1734:
																	position, tokenIndex = position1733, tokenIndex1733
																	if buffer[position] != rune('J') {
																		goto l1708
																	}
																	position++
																}
															l1733:
																{
																	position1735, tokenIndex1735 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1736
																	}
																	position++
																	goto l1735
																l1736:
																	position, tokenIndex = position1735, tokenIndex1735
																	if buffer[position] != rune('N') {
																		goto l1708
																	}
																	position++
																}
															l1735:
																{
																	position1737, tokenIndex1737 := position, tokenIndex
																	if buffer[position] != rune('z') {
																		goto l1738
																	}
																	position++
																	goto l1737
																l1738:
																	position, tokenIndex = position1737, tokenIndex1737
																	if buffer[position] != rune('Z') {
																		goto l1708
																	}
																	position++
																}
															l1737:
																if !_rules[rulews]() {
																	goto l1708
																}
																if !_rules[ruledisp]() {
																	goto l1708
																}
																{
																	add(ruleAction104, position)
																}
																add(ruleDjnz, position1730)
															}
															break
														case 'J', 'j':
															{
																position1740 := position
																{
																	position1741, tokenIndex1741 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1742
																	}
																	position++
																	goto l1741
																l1742:
																	position, tokenIndex = position1741, tokenIndex1741
																	if buffer[position] != rune('J') {
																		goto l1708
																	}
																	position++
																}
															l1741:
																{
																	position1743, tokenIndex1743 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1744
																	}
																	position++
																	goto l1743
																l1744:
																	position, tokenIndex = position1743, tokenIndex1743
																	if buffer[position] != rune('R') {
																		goto l1708
																	}
																	position++
																}
															l1743:
																if !_rules[rulews]() {
																	goto l1708
																}
																{
																	position1745, tokenIndex1745 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1745
																	}
																	if !_rules[rulesep]() {
																		goto l1745
																	}
																	goto l1746
																l1745:
																	position, tokenIndex = position1745, tokenIndex1745
																}
															l1746:
																if !_rules[ruledisp]() {
																	goto l1708
																}
																{
																	add(ruleAction103, position)
																}
																add(ruleJr, position1740)
															}
															break
														case 'R', 'r':
															{
																position1748 := position
																{
																	position1749, tokenIndex1749 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1750
																	}
																	position++
																	goto l1749
																l1750:
																	position, tokenIndex = position1749, tokenIndex1749
																	if buffer[position] != rune('R') {
																		goto l1708
																	}
																	position++
																}
															l1749:
																{
																	position1751, tokenIndex1751 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1752
																	}
																	position++
																	goto l1751
																l1752:
																	position, tokenIndex = position1751, tokenIndex1751
																	if buffer[position] != rune('E') {
																		goto l1708
																	}
																	position++
																}
															l1751:
																{
																	position1753, tokenIndex1753 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1754
																	}
																	position++
																	goto l1753
																l1754:
																	position, tokenIndex = position1753, tokenIndex1753
																	if buffer[position] != rune('T') {
																		goto l1708
																	}
																	position++
																}
															l1753:
																{
																	position1755, tokenIndex1755 := position, tokenIndex
																	if !_rules[rulews]() {
																		goto l1755
																	}
																	if !_rules[rulecc]() {
																		goto l1755
																	}
																	goto l1756
																l1755:
																	position, tokenIndex = position1755, tokenIndex1755
																}
															l1756:
																{
																	add(ruleAction101, position)
																}
																add(ruleRet, position1748)
															}
															break
														default:
															{
																position1758 := position
																{
																	position1759, tokenIndex1759 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1760
																	}
																	position++
																	goto l1759
																l1760:
																	position, tokenIndex = position1759, tokenIndex1759
																	if buffer[position] != rune('C') {
																		goto l1708
																	}
																	position++
																}
															l1759:
																{
																	position1761, tokenIndex1761 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1762
																	}
																	position++
																	goto l1761
																l1762:
																	position, tokenIndex = position1761, tokenIndex1761
																	if buffer[position] != rune('A') {
																		goto l1708
																	}
																	position++
																}
															l1761:
																{
																	position1763, tokenIndex1763 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1764
																	}
																	position++
																	goto l1763
																l1764:
																	position, tokenIndex = position1763, tokenIndex1763
																	if buffer[position] != rune('L') {
																		goto l1708
																	}
																	position++
																}
															l1763:
																{
																	position1765, tokenIndex1765 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1766
																	}
																	position++
																	goto l1765
																l1766:
																	position, tokenIndex = position1765, tokenIndex1765
																	if buffer[position] != rune('L') {
																		goto l1708
																	}
																	position++
																}
															l1765:
																if !_rules[rulews]() {
																	goto l1708
																}
																{
																	position1767, tokenIndex1767 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1767
																	}
																	if !_rules[rulesep]() {
																		goto l1767
																	}
																	goto l1768
																l1767:
																	position, tokenIndex = position1767, tokenIndex1767
																}
															l1768:
																if !_rules[ruleSrc16]() {
																	goto l1708
																}
																{
																	add(ruleAction100, position)
																}
																add(ruleCall, position1758)
															}
															break
														}
													}

												}
											l1710:
												add(ruleJump, position1709)
											}
											goto l989
										l1708:
											position, tokenIndex = position989, tokenIndex989
											{
												position1770 := position
												{
													position1771, tokenIndex1771 := position, tokenIndex
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
															if buffer[position] != rune('n') {
																goto l1777
															}
															position++
															goto l1776
														l1777:
															position, tokenIndex = position1776, tokenIndex1776
															if buffer[position] != rune('N') {
																goto l1772
															}
															position++
														}
													l1776:
														if !_rules[rulews]() {
															goto l1772
														}
														if !_rules[ruleReg8]() {
															goto l1772
														}
														if !_rules[rulesep]() {
															goto l1772
														}
														if !_rules[rulePort]() {
															goto l1772
														}
														{
															add(ruleAction105, position)
														}
														add(ruleIN, position1773)
													}
													goto l1771
												l1772:
													position, tokenIndex = position1771, tokenIndex1771
													{
														position1779 := position
														{
															position1780, tokenIndex1780 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1781
															}
															position++
															goto l1780
														l1781:
															position, tokenIndex = position1780, tokenIndex1780
															if buffer[position] != rune('O') {
																goto l915
															}
															position++
														}
													l1780:
														{
															position1782, tokenIndex1782 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1783
															}
															position++
															goto l1782
														l1783:
															position, tokenIndex = position1782, tokenIndex1782
															if buffer[position] != rune('U') {
																goto l915
															}
															position++
														}
													l1782:
														{
															position1784, tokenIndex1784 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1785
															}
															position++
															goto l1784
														l1785:
															position, tokenIndex = position1784, tokenIndex1784
															if buffer[position] != rune('T') {
																goto l915
															}
															position++
														}
													l1784:
														if !_rules[rulews]() {
															goto l915
														}
														if !_rules[rulePort]() {
															goto l915
														}
														if !_rules[rulesep]() {
															goto l915
														}
														if !_rules[ruleReg8]() {
															goto l915
														}
														{
															add(ruleAction106, position)
														}
														add(ruleOUT, position1779)
													}
												}
											l1771:
												add(ruleIO, position1770)
											}
										}
									l989:
										add(ruleInstruction, position988)
									}
								}
							l918:
								add(ruleStatement, position917)
							}
							goto l916
						l915:
							position, tokenIndex = position915, tokenIndex915
						}
					l916:
						{
							position1787, tokenIndex1787 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1787
							}
							goto l1788
						l1787:
							position, tokenIndex = position1787, tokenIndex1787
						}
					l1788:
						{
							position1789, tokenIndex1789 := position, tokenIndex
							{
								position1791 := position
								{
									position1792, tokenIndex1792 := position, tokenIndex
									if buffer[position] != rune(';') {
										goto l1793
									}
									position++
									goto l1792
								l1793:
									position, tokenIndex = position1792, tokenIndex1792
									if buffer[position] != rune('#') {
										goto l1789
									}
									position++
								}
							l1792:
							l1794:
								{
									position1795, tokenIndex1795 := position, tokenIndex
									{
										position1796, tokenIndex1796 := position, tokenIndex
										if buffer[position] != rune('\n') {
											goto l1796
										}
										position++
										goto l1795
									l1796:
										position, tokenIndex = position1796, tokenIndex1796
									}
									if !matchDot() {
										goto l1795
									}
									goto l1794
								l1795:
									position, tokenIndex = position1795, tokenIndex1795
								}
								add(ruleComment, position1791)
							}
							goto l1790
						l1789:
							position, tokenIndex = position1789, tokenIndex1789
						}
					l1790:
						{
							position1797, tokenIndex1797 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1797
							}
							goto l1798
						l1797:
							position, tokenIndex = position1797, tokenIndex1797
						}
					l1798:
						{
							position1799, tokenIndex1799 := position, tokenIndex
							{
								position1801, tokenIndex1801 := position, tokenIndex
								if buffer[position] != rune('\r') {
									goto l1801
								}
								position++
								goto l1802
							l1801:
								position, tokenIndex = position1801, tokenIndex1801
							}
						l1802:
							if buffer[position] != rune('\n') {
								goto l1800
							}
							position++
							goto l1799
						l1800:
							position, tokenIndex = position1799, tokenIndex1799
							if buffer[position] != rune(':') {
								goto l3
							}
							position++
						}
					l1799:
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position904)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position1804, tokenIndex1804 := position, tokenIndex
					if !matchDot() {
						goto l1804
					}
					goto l0
				l1804:
					position, tokenIndex = position1804, tokenIndex1804
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
		/* 10 LabelDefn <- <(LabelText ':' ws? Action5)> */
		nil,
		/* 11 LabelText <- <<(alpha alphanum alphanum+)>> */
		func() bool {
			position1815, tokenIndex1815 := position, tokenIndex
			{
				position1816 := position
				{
					position1817 := position
					if !_rules[rulealpha]() {
						goto l1815
					}
					if !_rules[rulealphanum]() {
						goto l1815
					}
					if !_rules[rulealphanum]() {
						goto l1815
					}
				l1818:
					{
						position1819, tokenIndex1819 := position, tokenIndex
						if !_rules[rulealphanum]() {
							goto l1819
						}
						goto l1818
					l1819:
						position, tokenIndex = position1819, tokenIndex1819
					}
					add(rulePegText, position1817)
				}
				add(ruleLabelText, position1816)
			}
			return true
		l1815:
			position, tokenIndex = position1815, tokenIndex1815
			return false
		},
		/* 12 alphanum <- <(alpha / num)> */
		func() bool {
			position1820, tokenIndex1820 := position, tokenIndex
			{
				position1821 := position
				{
					position1822, tokenIndex1822 := position, tokenIndex
					if !_rules[rulealpha]() {
						goto l1823
					}
					goto l1822
				l1823:
					position, tokenIndex = position1822, tokenIndex1822
					{
						position1824 := position
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1820
						}
						position++
						add(rulenum, position1824)
					}
				}
			l1822:
				add(rulealphanum, position1821)
			}
			return true
		l1820:
			position, tokenIndex = position1820, tokenIndex1820
			return false
		},
		/* 13 alpha <- <([a-z] / [A-Z])> */
		func() bool {
			position1825, tokenIndex1825 := position, tokenIndex
			{
				position1826 := position
				{
					position1827, tokenIndex1827 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1828
					}
					position++
					goto l1827
				l1828:
					position, tokenIndex = position1827, tokenIndex1827
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1825
					}
					position++
				}
			l1827:
				add(rulealpha, position1826)
			}
			return true
		l1825:
			position, tokenIndex = position1825, tokenIndex1825
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
			position1852, tokenIndex1852 := position, tokenIndex
			{
				position1853 := position
				{
					position1854, tokenIndex1854 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1855
					}
					goto l1854
				l1855:
					position, tokenIndex = position1854, tokenIndex1854
					if !_rules[ruleReg8]() {
						goto l1856
					}
					goto l1854
				l1856:
					position, tokenIndex = position1854, tokenIndex1854
					if !_rules[ruleReg16Contents]() {
						goto l1857
					}
					goto l1854
				l1857:
					position, tokenIndex = position1854, tokenIndex1854
					if !_rules[rulenn_contents]() {
						goto l1852
					}
				}
			l1854:
				{
					add(ruleAction21, position)
				}
				add(ruleSrc8, position1853)
			}
			return true
		l1852:
			position, tokenIndex = position1852, tokenIndex1852
			return false
		},
		/* 38 Loc8 <- <((Reg8 / Reg16Contents) Action22)> */
		func() bool {
			position1859, tokenIndex1859 := position, tokenIndex
			{
				position1860 := position
				{
					position1861, tokenIndex1861 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1862
					}
					goto l1861
				l1862:
					position, tokenIndex = position1861, tokenIndex1861
					if !_rules[ruleReg16Contents]() {
						goto l1859
					}
				}
			l1861:
				{
					add(ruleAction22, position)
				}
				add(ruleLoc8, position1860)
			}
			return true
		l1859:
			position, tokenIndex = position1859, tokenIndex1859
			return false
		},
		/* 39 Copy8 <- <(Reg8 Action23)> */
		func() bool {
			position1864, tokenIndex1864 := position, tokenIndex
			{
				position1865 := position
				if !_rules[ruleReg8]() {
					goto l1864
				}
				{
					add(ruleAction23, position)
				}
				add(ruleCopy8, position1865)
			}
			return true
		l1864:
			position, tokenIndex = position1864, tokenIndex1864
			return false
		},
		/* 40 ILoc8 <- <(IReg8 Action24)> */
		func() bool {
			position1867, tokenIndex1867 := position, tokenIndex
			{
				position1868 := position
				if !_rules[ruleIReg8]() {
					goto l1867
				}
				{
					add(ruleAction24, position)
				}
				add(ruleILoc8, position1868)
			}
			return true
		l1867:
			position, tokenIndex = position1867, tokenIndex1867
			return false
		},
		/* 41 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action25)> */
		func() bool {
			position1870, tokenIndex1870 := position, tokenIndex
			{
				position1871 := position
				{
					position1872 := position
					{
						position1873, tokenIndex1873 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1874
						}
						goto l1873
					l1874:
						position, tokenIndex = position1873, tokenIndex1873
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1876 := position
									{
										position1877, tokenIndex1877 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1878
										}
										position++
										goto l1877
									l1878:
										position, tokenIndex = position1877, tokenIndex1877
										if buffer[position] != rune('R') {
											goto l1870
										}
										position++
									}
								l1877:
									add(ruleR, position1876)
								}
								break
							case 'I', 'i':
								{
									position1879 := position
									{
										position1880, tokenIndex1880 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1881
										}
										position++
										goto l1880
									l1881:
										position, tokenIndex = position1880, tokenIndex1880
										if buffer[position] != rune('I') {
											goto l1870
										}
										position++
									}
								l1880:
									add(ruleI, position1879)
								}
								break
							case 'L', 'l':
								{
									position1882 := position
									{
										position1883, tokenIndex1883 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1884
										}
										position++
										goto l1883
									l1884:
										position, tokenIndex = position1883, tokenIndex1883
										if buffer[position] != rune('L') {
											goto l1870
										}
										position++
									}
								l1883:
									add(ruleL, position1882)
								}
								break
							case 'H', 'h':
								{
									position1885 := position
									{
										position1886, tokenIndex1886 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1887
										}
										position++
										goto l1886
									l1887:
										position, tokenIndex = position1886, tokenIndex1886
										if buffer[position] != rune('H') {
											goto l1870
										}
										position++
									}
								l1886:
									add(ruleH, position1885)
								}
								break
							case 'E', 'e':
								{
									position1888 := position
									{
										position1889, tokenIndex1889 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1890
										}
										position++
										goto l1889
									l1890:
										position, tokenIndex = position1889, tokenIndex1889
										if buffer[position] != rune('E') {
											goto l1870
										}
										position++
									}
								l1889:
									add(ruleE, position1888)
								}
								break
							case 'D', 'd':
								{
									position1891 := position
									{
										position1892, tokenIndex1892 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1893
										}
										position++
										goto l1892
									l1893:
										position, tokenIndex = position1892, tokenIndex1892
										if buffer[position] != rune('D') {
											goto l1870
										}
										position++
									}
								l1892:
									add(ruleD, position1891)
								}
								break
							case 'C', 'c':
								{
									position1894 := position
									{
										position1895, tokenIndex1895 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1896
										}
										position++
										goto l1895
									l1896:
										position, tokenIndex = position1895, tokenIndex1895
										if buffer[position] != rune('C') {
											goto l1870
										}
										position++
									}
								l1895:
									add(ruleC, position1894)
								}
								break
							case 'B', 'b':
								{
									position1897 := position
									{
										position1898, tokenIndex1898 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1899
										}
										position++
										goto l1898
									l1899:
										position, tokenIndex = position1898, tokenIndex1898
										if buffer[position] != rune('B') {
											goto l1870
										}
										position++
									}
								l1898:
									add(ruleB, position1897)
								}
								break
							case 'F', 'f':
								{
									position1900 := position
									{
										position1901, tokenIndex1901 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1902
										}
										position++
										goto l1901
									l1902:
										position, tokenIndex = position1901, tokenIndex1901
										if buffer[position] != rune('F') {
											goto l1870
										}
										position++
									}
								l1901:
									add(ruleF, position1900)
								}
								break
							default:
								{
									position1903 := position
									{
										position1904, tokenIndex1904 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1905
										}
										position++
										goto l1904
									l1905:
										position, tokenIndex = position1904, tokenIndex1904
										if buffer[position] != rune('A') {
											goto l1870
										}
										position++
									}
								l1904:
									add(ruleA, position1903)
								}
								break
							}
						}

					}
				l1873:
					add(rulePegText, position1872)
				}
				{
					add(ruleAction25, position)
				}
				add(ruleReg8, position1871)
			}
			return true
		l1870:
			position, tokenIndex = position1870, tokenIndex1870
			return false
		},
		/* 42 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action26)> */
		func() bool {
			position1907, tokenIndex1907 := position, tokenIndex
			{
				position1908 := position
				{
					position1909 := position
					{
						position1910, tokenIndex1910 := position, tokenIndex
						{
							position1912 := position
							{
								position1913, tokenIndex1913 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1914
								}
								position++
								goto l1913
							l1914:
								position, tokenIndex = position1913, tokenIndex1913
								if buffer[position] != rune('I') {
									goto l1911
								}
								position++
							}
						l1913:
							{
								position1915, tokenIndex1915 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1916
								}
								position++
								goto l1915
							l1916:
								position, tokenIndex = position1915, tokenIndex1915
								if buffer[position] != rune('X') {
									goto l1911
								}
								position++
							}
						l1915:
							{
								position1917, tokenIndex1917 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1918
								}
								position++
								goto l1917
							l1918:
								position, tokenIndex = position1917, tokenIndex1917
								if buffer[position] != rune('H') {
									goto l1911
								}
								position++
							}
						l1917:
							add(ruleIXH, position1912)
						}
						goto l1910
					l1911:
						position, tokenIndex = position1910, tokenIndex1910
						{
							position1920 := position
							{
								position1921, tokenIndex1921 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1922
								}
								position++
								goto l1921
							l1922:
								position, tokenIndex = position1921, tokenIndex1921
								if buffer[position] != rune('I') {
									goto l1919
								}
								position++
							}
						l1921:
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
									goto l1919
								}
								position++
							}
						l1923:
							{
								position1925, tokenIndex1925 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1926
								}
								position++
								goto l1925
							l1926:
								position, tokenIndex = position1925, tokenIndex1925
								if buffer[position] != rune('L') {
									goto l1919
								}
								position++
							}
						l1925:
							add(ruleIXL, position1920)
						}
						goto l1910
					l1919:
						position, tokenIndex = position1910, tokenIndex1910
						{
							position1928 := position
							{
								position1929, tokenIndex1929 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1930
								}
								position++
								goto l1929
							l1930:
								position, tokenIndex = position1929, tokenIndex1929
								if buffer[position] != rune('I') {
									goto l1927
								}
								position++
							}
						l1929:
							{
								position1931, tokenIndex1931 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1932
								}
								position++
								goto l1931
							l1932:
								position, tokenIndex = position1931, tokenIndex1931
								if buffer[position] != rune('Y') {
									goto l1927
								}
								position++
							}
						l1931:
							{
								position1933, tokenIndex1933 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1934
								}
								position++
								goto l1933
							l1934:
								position, tokenIndex = position1933, tokenIndex1933
								if buffer[position] != rune('H') {
									goto l1927
								}
								position++
							}
						l1933:
							add(ruleIYH, position1928)
						}
						goto l1910
					l1927:
						position, tokenIndex = position1910, tokenIndex1910
						{
							position1935 := position
							{
								position1936, tokenIndex1936 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1937
								}
								position++
								goto l1936
							l1937:
								position, tokenIndex = position1936, tokenIndex1936
								if buffer[position] != rune('I') {
									goto l1907
								}
								position++
							}
						l1936:
							{
								position1938, tokenIndex1938 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1939
								}
								position++
								goto l1938
							l1939:
								position, tokenIndex = position1938, tokenIndex1938
								if buffer[position] != rune('Y') {
									goto l1907
								}
								position++
							}
						l1938:
							{
								position1940, tokenIndex1940 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1941
								}
								position++
								goto l1940
							l1941:
								position, tokenIndex = position1940, tokenIndex1940
								if buffer[position] != rune('L') {
									goto l1907
								}
								position++
							}
						l1940:
							add(ruleIYL, position1935)
						}
					}
				l1910:
					add(rulePegText, position1909)
				}
				{
					add(ruleAction26, position)
				}
				add(ruleIReg8, position1908)
			}
			return true
		l1907:
			position, tokenIndex = position1907, tokenIndex1907
			return false
		},
		/* 43 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action27)> */
		func() bool {
			position1943, tokenIndex1943 := position, tokenIndex
			{
				position1944 := position
				{
					position1945, tokenIndex1945 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1946
					}
					goto l1945
				l1946:
					position, tokenIndex = position1945, tokenIndex1945
					if !_rules[rulenn_contents]() {
						goto l1947
					}
					goto l1945
				l1947:
					position, tokenIndex = position1945, tokenIndex1945
					if !_rules[ruleReg16Contents]() {
						goto l1943
					}
				}
			l1945:
				{
					add(ruleAction27, position)
				}
				add(ruleDst16, position1944)
			}
			return true
		l1943:
			position, tokenIndex = position1943, tokenIndex1943
			return false
		},
		/* 44 Src16 <- <((Reg16 / nn / nn_contents) Action28)> */
		func() bool {
			position1949, tokenIndex1949 := position, tokenIndex
			{
				position1950 := position
				{
					position1951, tokenIndex1951 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1952
					}
					goto l1951
				l1952:
					position, tokenIndex = position1951, tokenIndex1951
					if !_rules[rulenn]() {
						goto l1953
					}
					goto l1951
				l1953:
					position, tokenIndex = position1951, tokenIndex1951
					if !_rules[rulenn_contents]() {
						goto l1949
					}
				}
			l1951:
				{
					add(ruleAction28, position)
				}
				add(ruleSrc16, position1950)
			}
			return true
		l1949:
			position, tokenIndex = position1949, tokenIndex1949
			return false
		},
		/* 45 Loc16 <- <(Reg16 Action29)> */
		func() bool {
			position1955, tokenIndex1955 := position, tokenIndex
			{
				position1956 := position
				if !_rules[ruleReg16]() {
					goto l1955
				}
				{
					add(ruleAction29, position)
				}
				add(ruleLoc16, position1956)
			}
			return true
		l1955:
			position, tokenIndex = position1955, tokenIndex1955
			return false
		},
		/* 46 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action30)> */
		func() bool {
			position1958, tokenIndex1958 := position, tokenIndex
			{
				position1959 := position
				{
					position1960 := position
					{
						position1961, tokenIndex1961 := position, tokenIndex
						{
							position1963 := position
							{
								position1964, tokenIndex1964 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1965
								}
								position++
								goto l1964
							l1965:
								position, tokenIndex = position1964, tokenIndex1964
								if buffer[position] != rune('A') {
									goto l1962
								}
								position++
							}
						l1964:
							{
								position1966, tokenIndex1966 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1967
								}
								position++
								goto l1966
							l1967:
								position, tokenIndex = position1966, tokenIndex1966
								if buffer[position] != rune('F') {
									goto l1962
								}
								position++
							}
						l1966:
							if buffer[position] != rune('\'') {
								goto l1962
							}
							position++
							add(ruleAF_PRIME, position1963)
						}
						goto l1961
					l1962:
						position, tokenIndex = position1961, tokenIndex1961
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1958
								}
								break
							case 'S', 's':
								{
									position1969 := position
									{
										position1970, tokenIndex1970 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1971
										}
										position++
										goto l1970
									l1971:
										position, tokenIndex = position1970, tokenIndex1970
										if buffer[position] != rune('S') {
											goto l1958
										}
										position++
									}
								l1970:
									{
										position1972, tokenIndex1972 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1973
										}
										position++
										goto l1972
									l1973:
										position, tokenIndex = position1972, tokenIndex1972
										if buffer[position] != rune('P') {
											goto l1958
										}
										position++
									}
								l1972:
									add(ruleSP, position1969)
								}
								break
							case 'H', 'h':
								{
									position1974 := position
									{
										position1975, tokenIndex1975 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1976
										}
										position++
										goto l1975
									l1976:
										position, tokenIndex = position1975, tokenIndex1975
										if buffer[position] != rune('H') {
											goto l1958
										}
										position++
									}
								l1975:
									{
										position1977, tokenIndex1977 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1978
										}
										position++
										goto l1977
									l1978:
										position, tokenIndex = position1977, tokenIndex1977
										if buffer[position] != rune('L') {
											goto l1958
										}
										position++
									}
								l1977:
									add(ruleHL, position1974)
								}
								break
							case 'D', 'd':
								{
									position1979 := position
									{
										position1980, tokenIndex1980 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1981
										}
										position++
										goto l1980
									l1981:
										position, tokenIndex = position1980, tokenIndex1980
										if buffer[position] != rune('D') {
											goto l1958
										}
										position++
									}
								l1980:
									{
										position1982, tokenIndex1982 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1983
										}
										position++
										goto l1982
									l1983:
										position, tokenIndex = position1982, tokenIndex1982
										if buffer[position] != rune('E') {
											goto l1958
										}
										position++
									}
								l1982:
									add(ruleDE, position1979)
								}
								break
							case 'B', 'b':
								{
									position1984 := position
									{
										position1985, tokenIndex1985 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1986
										}
										position++
										goto l1985
									l1986:
										position, tokenIndex = position1985, tokenIndex1985
										if buffer[position] != rune('B') {
											goto l1958
										}
										position++
									}
								l1985:
									{
										position1987, tokenIndex1987 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1988
										}
										position++
										goto l1987
									l1988:
										position, tokenIndex = position1987, tokenIndex1987
										if buffer[position] != rune('C') {
											goto l1958
										}
										position++
									}
								l1987:
									add(ruleBC, position1984)
								}
								break
							default:
								{
									position1989 := position
									{
										position1990, tokenIndex1990 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1991
										}
										position++
										goto l1990
									l1991:
										position, tokenIndex = position1990, tokenIndex1990
										if buffer[position] != rune('A') {
											goto l1958
										}
										position++
									}
								l1990:
									{
										position1992, tokenIndex1992 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1993
										}
										position++
										goto l1992
									l1993:
										position, tokenIndex = position1992, tokenIndex1992
										if buffer[position] != rune('F') {
											goto l1958
										}
										position++
									}
								l1992:
									add(ruleAF, position1989)
								}
								break
							}
						}

					}
				l1961:
					add(rulePegText, position1960)
				}
				{
					add(ruleAction30, position)
				}
				add(ruleReg16, position1959)
			}
			return true
		l1958:
			position, tokenIndex = position1958, tokenIndex1958
			return false
		},
		/* 47 IReg16 <- <(<(IX / IY)> Action31)> */
		func() bool {
			position1995, tokenIndex1995 := position, tokenIndex
			{
				position1996 := position
				{
					position1997 := position
					{
						position1998, tokenIndex1998 := position, tokenIndex
						{
							position2000 := position
							{
								position2001, tokenIndex2001 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l2002
								}
								position++
								goto l2001
							l2002:
								position, tokenIndex = position2001, tokenIndex2001
								if buffer[position] != rune('I') {
									goto l1999
								}
								position++
							}
						l2001:
							{
								position2003, tokenIndex2003 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l2004
								}
								position++
								goto l2003
							l2004:
								position, tokenIndex = position2003, tokenIndex2003
								if buffer[position] != rune('X') {
									goto l1999
								}
								position++
							}
						l2003:
							add(ruleIX, position2000)
						}
						goto l1998
					l1999:
						position, tokenIndex = position1998, tokenIndex1998
						{
							position2005 := position
							{
								position2006, tokenIndex2006 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l2007
								}
								position++
								goto l2006
							l2007:
								position, tokenIndex = position2006, tokenIndex2006
								if buffer[position] != rune('I') {
									goto l1995
								}
								position++
							}
						l2006:
							{
								position2008, tokenIndex2008 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l2009
								}
								position++
								goto l2008
							l2009:
								position, tokenIndex = position2008, tokenIndex2008
								if buffer[position] != rune('Y') {
									goto l1995
								}
								position++
							}
						l2008:
							add(ruleIY, position2005)
						}
					}
				l1998:
					add(rulePegText, position1997)
				}
				{
					add(ruleAction31, position)
				}
				add(ruleIReg16, position1996)
			}
			return true
		l1995:
			position, tokenIndex = position1995, tokenIndex1995
			return false
		},
		/* 48 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position2011, tokenIndex2011 := position, tokenIndex
			{
				position2012 := position
				{
					position2013, tokenIndex2013 := position, tokenIndex
					{
						position2015 := position
						if buffer[position] != rune('(') {
							goto l2014
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l2014
						}
						{
							position2016, tokenIndex2016 := position, tokenIndex
							if !_rules[rulews]() {
								goto l2016
							}
							goto l2017
						l2016:
							position, tokenIndex = position2016, tokenIndex2016
						}
					l2017:
						if !_rules[ruledisp]() {
							goto l2014
						}
						{
							position2018, tokenIndex2018 := position, tokenIndex
							if !_rules[rulews]() {
								goto l2018
							}
							goto l2019
						l2018:
							position, tokenIndex = position2018, tokenIndex2018
						}
					l2019:
						if buffer[position] != rune(')') {
							goto l2014
						}
						position++
						{
							add(ruleAction33, position)
						}
						add(ruleIndexedR16C, position2015)
					}
					goto l2013
				l2014:
					position, tokenIndex = position2013, tokenIndex2013
					{
						position2021 := position
						if buffer[position] != rune('(') {
							goto l2011
						}
						position++
						if !_rules[ruleReg16]() {
							goto l2011
						}
						if buffer[position] != rune(')') {
							goto l2011
						}
						position++
						{
							add(ruleAction32, position)
						}
						add(rulePlainR16C, position2021)
					}
				}
			l2013:
				add(ruleReg16Contents, position2012)
			}
			return true
		l2011:
			position, tokenIndex = position2011, tokenIndex2011
			return false
		},
		/* 49 PlainR16C <- <('(' Reg16 ')' Action32)> */
		nil,
		/* 50 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action33)> */
		nil,
		/* 51 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position2025, tokenIndex2025 := position, tokenIndex
			{
				position2026 := position
				{
					position2027, tokenIndex2027 := position, tokenIndex
					{
						position2029 := position
						{
							position2030 := position
							if !_rules[rulehexdigit]() {
								goto l2028
							}
							if !_rules[rulehexdigit]() {
								goto l2028
							}
							add(rulePegText, position2030)
						}
						{
							position2031, tokenIndex2031 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2032
							}
							position++
							goto l2031
						l2032:
							position, tokenIndex = position2031, tokenIndex2031
							if buffer[position] != rune('H') {
								goto l2028
							}
							position++
						}
					l2031:
						{
							add(ruleAction37, position)
						}
						add(rulehexByteH, position2029)
					}
					goto l2027
				l2028:
					position, tokenIndex = position2027, tokenIndex2027
					{
						position2035 := position
						if buffer[position] != rune('0') {
							goto l2034
						}
						position++
						{
							position2036, tokenIndex2036 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2037
							}
							position++
							goto l2036
						l2037:
							position, tokenIndex = position2036, tokenIndex2036
							if buffer[position] != rune('X') {
								goto l2034
							}
							position++
						}
					l2036:
						{
							position2038 := position
							if !_rules[rulehexdigit]() {
								goto l2034
							}
							if !_rules[rulehexdigit]() {
								goto l2034
							}
							add(rulePegText, position2038)
						}
						{
							add(ruleAction38, position)
						}
						add(rulehexByte0x, position2035)
					}
					goto l2027
				l2034:
					position, tokenIndex = position2027, tokenIndex2027
					{
						position2040 := position
						{
							position2041 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2025
							}
							position++
						l2042:
							{
								position2043, tokenIndex2043 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2043
								}
								position++
								goto l2042
							l2043:
								position, tokenIndex = position2043, tokenIndex2043
							}
							add(rulePegText, position2041)
						}
						{
							add(ruleAction39, position)
						}
						add(ruledecimalByte, position2040)
					}
				}
			l2027:
				add(rulen, position2026)
			}
			return true
		l2025:
			position, tokenIndex = position2025, tokenIndex2025
			return false
		},
		/* 52 nn <- <(LabelNN / hexWordH / hexWord0x)> */
		func() bool {
			position2045, tokenIndex2045 := position, tokenIndex
			{
				position2046 := position
				{
					position2047, tokenIndex2047 := position, tokenIndex
					{
						position2049 := position
						{
							position2050 := position
							if !_rules[ruleLabelText]() {
								goto l2048
							}
							add(rulePegText, position2050)
						}
						{
							add(ruleAction40, position)
						}
						add(ruleLabelNN, position2049)
					}
					goto l2047
				l2048:
					position, tokenIndex = position2047, tokenIndex2047
					{
						position2053 := position
						{
							position2054, tokenIndex2054 := position, tokenIndex
							if !_rules[rulezeroHexWord]() {
								goto l2055
							}
							goto l2054
						l2055:
							position, tokenIndex = position2054, tokenIndex2054
							if !_rules[rulehexWord]() {
								goto l2052
							}
						}
					l2054:
						{
							position2056, tokenIndex2056 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2057
							}
							position++
							goto l2056
						l2057:
							position, tokenIndex = position2056, tokenIndex2056
							if buffer[position] != rune('H') {
								goto l2052
							}
							position++
						}
					l2056:
						add(rulehexWordH, position2053)
					}
					goto l2047
				l2052:
					position, tokenIndex = position2047, tokenIndex2047
					{
						position2058 := position
						if buffer[position] != rune('0') {
							goto l2045
						}
						position++
						{
							position2059, tokenIndex2059 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2060
							}
							position++
							goto l2059
						l2060:
							position, tokenIndex = position2059, tokenIndex2059
							if buffer[position] != rune('X') {
								goto l2045
							}
							position++
						}
					l2059:
						{
							position2061, tokenIndex2061 := position, tokenIndex
							if !_rules[rulezeroHexWord]() {
								goto l2062
							}
							goto l2061
						l2062:
							position, tokenIndex = position2061, tokenIndex2061
							if !_rules[rulehexWord]() {
								goto l2045
							}
						}
					l2061:
						add(rulehexWord0x, position2058)
					}
				}
			l2047:
				add(rulenn, position2046)
			}
			return true
		l2045:
			position, tokenIndex = position2045, tokenIndex2045
			return false
		},
		/* 53 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position2063, tokenIndex2063 := position, tokenIndex
			{
				position2064 := position
				{
					position2065, tokenIndex2065 := position, tokenIndex
					{
						position2067 := position
						{
							position2068 := position
							{
								position2069, tokenIndex2069 := position, tokenIndex
								{
									position2071, tokenIndex2071 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2072
									}
									position++
									goto l2071
								l2072:
									position, tokenIndex = position2071, tokenIndex2071
									if buffer[position] != rune('+') {
										goto l2069
									}
									position++
								}
							l2071:
								goto l2070
							l2069:
								position, tokenIndex = position2069, tokenIndex2069
							}
						l2070:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2066
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2066
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2066
									}
									position++
									break
								}
							}

						l2073:
							{
								position2074, tokenIndex2074 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2074
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2074
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2074
										}
										position++
										break
									}
								}

								goto l2073
							l2074:
								position, tokenIndex = position2074, tokenIndex2074
							}
							add(rulePegText, position2068)
						}
						{
							position2077, tokenIndex2077 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2078
							}
							position++
							goto l2077
						l2078:
							position, tokenIndex = position2077, tokenIndex2077
							if buffer[position] != rune('H') {
								goto l2066
							}
							position++
						}
					l2077:
						{
							add(ruleAction35, position)
						}
						add(rulesignedHexByteH, position2067)
					}
					goto l2065
				l2066:
					position, tokenIndex = position2065, tokenIndex2065
					{
						position2081 := position
						{
							position2082 := position
							{
								position2083, tokenIndex2083 := position, tokenIndex
								{
									position2085, tokenIndex2085 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2086
									}
									position++
									goto l2085
								l2086:
									position, tokenIndex = position2085, tokenIndex2085
									if buffer[position] != rune('+') {
										goto l2083
									}
									position++
								}
							l2085:
								goto l2084
							l2083:
								position, tokenIndex = position2083, tokenIndex2083
							}
						l2084:
							if buffer[position] != rune('0') {
								goto l2080
							}
							position++
							{
								position2087, tokenIndex2087 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l2088
								}
								position++
								goto l2087
							l2088:
								position, tokenIndex = position2087, tokenIndex2087
								if buffer[position] != rune('X') {
									goto l2080
								}
								position++
							}
						l2087:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2080
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2080
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2080
									}
									position++
									break
								}
							}

						l2089:
							{
								position2090, tokenIndex2090 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2090
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2090
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2090
										}
										position++
										break
									}
								}

								goto l2089
							l2090:
								position, tokenIndex = position2090, tokenIndex2090
							}
							add(rulePegText, position2082)
						}
						{
							add(ruleAction36, position)
						}
						add(rulesignedHexByte0x, position2081)
					}
					goto l2065
				l2080:
					position, tokenIndex = position2065, tokenIndex2065
					{
						position2094 := position
						{
							position2095 := position
							{
								position2096, tokenIndex2096 := position, tokenIndex
								{
									position2098, tokenIndex2098 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2099
									}
									position++
									goto l2098
								l2099:
									position, tokenIndex = position2098, tokenIndex2098
									if buffer[position] != rune('+') {
										goto l2096
									}
									position++
								}
							l2098:
								goto l2097
							l2096:
								position, tokenIndex = position2096, tokenIndex2096
							}
						l2097:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2063
							}
							position++
						l2100:
							{
								position2101, tokenIndex2101 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2101
								}
								position++
								goto l2100
							l2101:
								position, tokenIndex = position2101, tokenIndex2101
							}
							add(rulePegText, position2095)
						}
						{
							add(ruleAction34, position)
						}
						add(rulesignedDecimalByte, position2094)
					}
				}
			l2065:
				add(ruledisp, position2064)
			}
			return true
		l2063:
			position, tokenIndex = position2063, tokenIndex2063
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
			position2112, tokenIndex2112 := position, tokenIndex
			{
				position2113 := position
				{
					position2114 := position
					if !_rules[rulehexdigit]() {
						goto l2112
					}
					if !_rules[rulehexdigit]() {
						goto l2112
					}
					if !_rules[rulehexdigit]() {
						goto l2112
					}
					if !_rules[rulehexdigit]() {
						goto l2112
					}
					add(rulePegText, position2114)
				}
				{
					add(ruleAction41, position)
				}
				add(rulehexWord, position2113)
			}
			return true
		l2112:
			position, tokenIndex = position2112, tokenIndex2112
			return false
		},
		/* 64 zeroHexWord <- <('0' hexWord)> */
		func() bool {
			position2116, tokenIndex2116 := position, tokenIndex
			{
				position2117 := position
				if buffer[position] != rune('0') {
					goto l2116
				}
				position++
				if !_rules[rulehexWord]() {
					goto l2116
				}
				add(rulezeroHexWord, position2117)
			}
			return true
		l2116:
			position, tokenIndex = position2116, tokenIndex2116
			return false
		},
		/* 65 nn_contents <- <('(' nn ')' Action42)> */
		func() bool {
			position2118, tokenIndex2118 := position, tokenIndex
			{
				position2119 := position
				if buffer[position] != rune('(') {
					goto l2118
				}
				position++
				if !_rules[rulenn]() {
					goto l2118
				}
				if buffer[position] != rune(')') {
					goto l2118
				}
				position++
				{
					add(ruleAction42, position)
				}
				add(rulenn_contents, position2119)
			}
			return true
		l2118:
			position, tokenIndex = position2118, tokenIndex2118
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
			position2194, tokenIndex2194 := position, tokenIndex
			{
				position2195 := position
				{
					position2196, tokenIndex2196 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2197
					}
					position++
					{
						position2198, tokenIndex2198 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l2199
						}
						position++
						goto l2198
					l2199:
						position, tokenIndex = position2198, tokenIndex2198
						if buffer[position] != rune('C') {
							goto l2197
						}
						position++
					}
				l2198:
					if buffer[position] != rune(')') {
						goto l2197
					}
					position++
					goto l2196
				l2197:
					position, tokenIndex = position2196, tokenIndex2196
					if buffer[position] != rune('(') {
						goto l2194
					}
					position++
					if !_rules[rulen]() {
						goto l2194
					}
					if buffer[position] != rune(')') {
						goto l2194
					}
					position++
				}
			l2196:
				add(rulePort, position2195)
			}
			return true
		l2194:
			position, tokenIndex = position2194, tokenIndex2194
			return false
		},
		/* 140 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2200, tokenIndex2200 := position, tokenIndex
			{
				position2201 := position
				{
					position2202, tokenIndex2202 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2202
					}
					goto l2203
				l2202:
					position, tokenIndex = position2202, tokenIndex2202
				}
			l2203:
				if buffer[position] != rune(',') {
					goto l2200
				}
				position++
				{
					position2204, tokenIndex2204 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2204
					}
					goto l2205
				l2204:
					position, tokenIndex = position2204, tokenIndex2204
				}
			l2205:
				add(rulesep, position2201)
			}
			return true
		l2200:
			position, tokenIndex = position2200, tokenIndex2200
			return false
		},
		/* 141 ws <- <(' ' / '\t')+> */
		func() bool {
			position2206, tokenIndex2206 := position, tokenIndex
			{
				position2207 := position
				{
					position2210, tokenIndex2210 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2211
					}
					position++
					goto l2210
				l2211:
					position, tokenIndex = position2210, tokenIndex2210
					if buffer[position] != rune('\t') {
						goto l2206
					}
					position++
				}
			l2210:
			l2208:
				{
					position2209, tokenIndex2209 := position, tokenIndex
					{
						position2212, tokenIndex2212 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l2213
						}
						position++
						goto l2212
					l2213:
						position, tokenIndex = position2212, tokenIndex2212
						if buffer[position] != rune('\t') {
							goto l2209
						}
						position++
					}
				l2212:
					goto l2208
				l2209:
					position, tokenIndex = position2209, tokenIndex2209
				}
				add(rulews, position2207)
			}
			return true
		l2206:
			position, tokenIndex = position2206, tokenIndex2206
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
			position2236, tokenIndex2236 := position, tokenIndex
			{
				position2237 := position
				{
					position2238, tokenIndex2238 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2239
					}
					position++
					goto l2238
				l2239:
					position, tokenIndex = position2238, tokenIndex2238
					{
						position2240, tokenIndex2240 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2241
						}
						position++
						goto l2240
					l2241:
						position, tokenIndex = position2240, tokenIndex2240
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2236
						}
						position++
					}
				l2240:
				}
			l2238:
				add(rulehexdigit, position2237)
			}
			return true
		l2236:
			position, tokenIndex = position2236, tokenIndex2236
			return false
		},
		/* 165 octaldigit <- <(<[0-7]> Action107)> */
		func() bool {
			position2242, tokenIndex2242 := position, tokenIndex
			{
				position2243 := position
				{
					position2244 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2242
					}
					position++
					add(rulePegText, position2244)
				}
				{
					add(ruleAction107, position)
				}
				add(ruleoctaldigit, position2243)
			}
			return true
		l2242:
			position, tokenIndex = position2242, tokenIndex2242
			return false
		},
		/* 166 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2246, tokenIndex2246 := position, tokenIndex
			{
				position2247 := position
				{
					position2248, tokenIndex2248 := position, tokenIndex
					{
						position2250 := position
						{
							position2251, tokenIndex2251 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2252
							}
							position++
							goto l2251
						l2252:
							position, tokenIndex = position2251, tokenIndex2251
							if buffer[position] != rune('N') {
								goto l2249
							}
							position++
						}
					l2251:
						{
							position2253, tokenIndex2253 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2254
							}
							position++
							goto l2253
						l2254:
							position, tokenIndex = position2253, tokenIndex2253
							if buffer[position] != rune('Z') {
								goto l2249
							}
							position++
						}
					l2253:
						{
							add(ruleAction108, position)
						}
						add(ruleFT_NZ, position2250)
					}
					goto l2248
				l2249:
					position, tokenIndex = position2248, tokenIndex2248
					{
						position2257 := position
						{
							position2258, tokenIndex2258 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2259
							}
							position++
							goto l2258
						l2259:
							position, tokenIndex = position2258, tokenIndex2258
							if buffer[position] != rune('P') {
								goto l2256
							}
							position++
						}
					l2258:
						{
							position2260, tokenIndex2260 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2261
							}
							position++
							goto l2260
						l2261:
							position, tokenIndex = position2260, tokenIndex2260
							if buffer[position] != rune('O') {
								goto l2256
							}
							position++
						}
					l2260:
						{
							add(ruleAction112, position)
						}
						add(ruleFT_PO, position2257)
					}
					goto l2248
				l2256:
					position, tokenIndex = position2248, tokenIndex2248
					{
						position2264 := position
						{
							position2265, tokenIndex2265 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2266
							}
							position++
							goto l2265
						l2266:
							position, tokenIndex = position2265, tokenIndex2265
							if buffer[position] != rune('P') {
								goto l2263
							}
							position++
						}
					l2265:
						{
							position2267, tokenIndex2267 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2268
							}
							position++
							goto l2267
						l2268:
							position, tokenIndex = position2267, tokenIndex2267
							if buffer[position] != rune('E') {
								goto l2263
							}
							position++
						}
					l2267:
						{
							add(ruleAction113, position)
						}
						add(ruleFT_PE, position2264)
					}
					goto l2248
				l2263:
					position, tokenIndex = position2248, tokenIndex2248
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2271 := position
								{
									position2272, tokenIndex2272 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2273
									}
									position++
									goto l2272
								l2273:
									position, tokenIndex = position2272, tokenIndex2272
									if buffer[position] != rune('M') {
										goto l2246
									}
									position++
								}
							l2272:
								{
									add(ruleAction115, position)
								}
								add(ruleFT_M, position2271)
							}
							break
						case 'P', 'p':
							{
								position2275 := position
								{
									position2276, tokenIndex2276 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2277
									}
									position++
									goto l2276
								l2277:
									position, tokenIndex = position2276, tokenIndex2276
									if buffer[position] != rune('P') {
										goto l2246
									}
									position++
								}
							l2276:
								{
									add(ruleAction114, position)
								}
								add(ruleFT_P, position2275)
							}
							break
						case 'C', 'c':
							{
								position2279 := position
								{
									position2280, tokenIndex2280 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2281
									}
									position++
									goto l2280
								l2281:
									position, tokenIndex = position2280, tokenIndex2280
									if buffer[position] != rune('C') {
										goto l2246
									}
									position++
								}
							l2280:
								{
									add(ruleAction111, position)
								}
								add(ruleFT_C, position2279)
							}
							break
						case 'N', 'n':
							{
								position2283 := position
								{
									position2284, tokenIndex2284 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2285
									}
									position++
									goto l2284
								l2285:
									position, tokenIndex = position2284, tokenIndex2284
									if buffer[position] != rune('N') {
										goto l2246
									}
									position++
								}
							l2284:
								{
									position2286, tokenIndex2286 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2287
									}
									position++
									goto l2286
								l2287:
									position, tokenIndex = position2286, tokenIndex2286
									if buffer[position] != rune('C') {
										goto l2246
									}
									position++
								}
							l2286:
								{
									add(ruleAction110, position)
								}
								add(ruleFT_NC, position2283)
							}
							break
						default:
							{
								position2289 := position
								{
									position2290, tokenIndex2290 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2291
									}
									position++
									goto l2290
								l2291:
									position, tokenIndex = position2290, tokenIndex2290
									if buffer[position] != rune('Z') {
										goto l2246
									}
									position++
								}
							l2290:
								{
									add(ruleAction109, position)
								}
								add(ruleFT_Z, position2289)
							}
							break
						}
					}

				}
			l2248:
				add(rulecc, position2247)
			}
			return true
		l2246:
			position, tokenIndex = position2246, tokenIndex2246
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
