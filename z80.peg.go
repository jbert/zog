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
	rulealphaundnum
	rulealphaund
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
	"alphaundnum",
	"alphaund",
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
															{
																position288, tokenIndex288 := position, tokenIndex
																{
																	position290, tokenIndex290 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l291
																	}
																	position++
																	goto l290
																l291:
																	position, tokenIndex = position290, tokenIndex290
																	if buffer[position] != rune('A') {
																		goto l288
																	}
																	position++
																}
															l290:
																if !_rules[rulesep]() {
																	goto l288
																}
																goto l289
															l288:
																position, tokenIndex = position288, tokenIndex288
															}
														l289:
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
															position293 := position
															{
																position294, tokenIndex294 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l295
																}
																position++
																goto l294
															l295:
																position, tokenIndex = position294, tokenIndex294
																if buffer[position] != rune('S') {
																	goto l226
																}
																position++
															}
														l294:
															{
																position296, tokenIndex296 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l297
																}
																position++
																goto l296
															l297:
																position, tokenIndex = position296, tokenIndex296
																if buffer[position] != rune('B') {
																	goto l226
																}
																position++
															}
														l296:
															{
																position298, tokenIndex298 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l299
																}
																position++
																goto l298
															l299:
																position, tokenIndex = position298, tokenIndex298
																if buffer[position] != rune('C') {
																	goto l226
																}
																position++
															}
														l298:
															if !_rules[rulews]() {
																goto l226
															}
															{
																position300, tokenIndex300 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l301
																}
																position++
																goto l300
															l301:
																position, tokenIndex = position300, tokenIndex300
																if buffer[position] != rune('A') {
																	goto l226
																}
																position++
															}
														l300:
															if !_rules[rulesep]() {
																goto l226
															}
															if !_rules[ruleSrc8]() {
																goto l226
															}
															{
																add(ruleAction46, position)
															}
															add(ruleSbc, position293)
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
											position304 := position
											{
												position305, tokenIndex305 := position, tokenIndex
												{
													position307 := position
													{
														position308, tokenIndex308 := position, tokenIndex
														{
															position310 := position
															{
																position311, tokenIndex311 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l312
																}
																position++
																goto l311
															l312:
																position, tokenIndex = position311, tokenIndex311
																if buffer[position] != rune('R') {
																	goto l309
																}
																position++
															}
														l311:
															{
																position313, tokenIndex313 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l314
																}
																position++
																goto l313
															l314:
																position, tokenIndex = position313, tokenIndex313
																if buffer[position] != rune('L') {
																	goto l309
																}
																position++
															}
														l313:
															{
																position315, tokenIndex315 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l316
																}
																position++
																goto l315
															l316:
																position, tokenIndex = position315, tokenIndex315
																if buffer[position] != rune('C') {
																	goto l309
																}
																position++
															}
														l315:
															if !_rules[rulews]() {
																goto l309
															}
															if !_rules[ruleLoc8]() {
																goto l309
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
																add(ruleAction51, position)
															}
															add(ruleRlc, position310)
														}
														goto l308
													l309:
														position, tokenIndex = position308, tokenIndex308
														{
															position321 := position
															{
																position322, tokenIndex322 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l323
																}
																position++
																goto l322
															l323:
																position, tokenIndex = position322, tokenIndex322
																if buffer[position] != rune('R') {
																	goto l320
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
																	goto l320
																}
																position++
															}
														l324:
															{
																position326, tokenIndex326 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l327
																}
																position++
																goto l326
															l327:
																position, tokenIndex = position326, tokenIndex326
																if buffer[position] != rune('C') {
																	goto l320
																}
																position++
															}
														l326:
															if !_rules[rulews]() {
																goto l320
															}
															if !_rules[ruleLoc8]() {
																goto l320
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
																add(ruleAction52, position)
															}
															add(ruleRrc, position321)
														}
														goto l308
													l320:
														position, tokenIndex = position308, tokenIndex308
														{
															position332 := position
															{
																position333, tokenIndex333 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l334
																}
																position++
																goto l333
															l334:
																position, tokenIndex = position333, tokenIndex333
																if buffer[position] != rune('R') {
																	goto l331
																}
																position++
															}
														l333:
															{
																position335, tokenIndex335 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l336
																}
																position++
																goto l335
															l336:
																position, tokenIndex = position335, tokenIndex335
																if buffer[position] != rune('L') {
																	goto l331
																}
																position++
															}
														l335:
															if !_rules[rulews]() {
																goto l331
															}
															if !_rules[ruleLoc8]() {
																goto l331
															}
															{
																position337, tokenIndex337 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l337
																}
																if !_rules[ruleCopy8]() {
																	goto l337
																}
																goto l338
															l337:
																position, tokenIndex = position337, tokenIndex337
															}
														l338:
															{
																add(ruleAction53, position)
															}
															add(ruleRl, position332)
														}
														goto l308
													l331:
														position, tokenIndex = position308, tokenIndex308
														{
															position341 := position
															{
																position342, tokenIndex342 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l343
																}
																position++
																goto l342
															l343:
																position, tokenIndex = position342, tokenIndex342
																if buffer[position] != rune('R') {
																	goto l340
																}
																position++
															}
														l342:
															{
																position344, tokenIndex344 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l345
																}
																position++
																goto l344
															l345:
																position, tokenIndex = position344, tokenIndex344
																if buffer[position] != rune('R') {
																	goto l340
																}
																position++
															}
														l344:
															if !_rules[rulews]() {
																goto l340
															}
															if !_rules[ruleLoc8]() {
																goto l340
															}
															{
																position346, tokenIndex346 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l346
																}
																if !_rules[ruleCopy8]() {
																	goto l346
																}
																goto l347
															l346:
																position, tokenIndex = position346, tokenIndex346
															}
														l347:
															{
																add(ruleAction54, position)
															}
															add(ruleRr, position341)
														}
														goto l308
													l340:
														position, tokenIndex = position308, tokenIndex308
														{
															position350 := position
															{
																position351, tokenIndex351 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l352
																}
																position++
																goto l351
															l352:
																position, tokenIndex = position351, tokenIndex351
																if buffer[position] != rune('S') {
																	goto l349
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
																	goto l349
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
																	goto l349
																}
																position++
															}
														l355:
															if !_rules[rulews]() {
																goto l349
															}
															if !_rules[ruleLoc8]() {
																goto l349
															}
															{
																position357, tokenIndex357 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l357
																}
																if !_rules[ruleCopy8]() {
																	goto l357
																}
																goto l358
															l357:
																position, tokenIndex = position357, tokenIndex357
															}
														l358:
															{
																add(ruleAction55, position)
															}
															add(ruleSla, position350)
														}
														goto l308
													l349:
														position, tokenIndex = position308, tokenIndex308
														{
															position361 := position
															{
																position362, tokenIndex362 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l363
																}
																position++
																goto l362
															l363:
																position, tokenIndex = position362, tokenIndex362
																if buffer[position] != rune('S') {
																	goto l360
																}
																position++
															}
														l362:
															{
																position364, tokenIndex364 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l365
																}
																position++
																goto l364
															l365:
																position, tokenIndex = position364, tokenIndex364
																if buffer[position] != rune('R') {
																	goto l360
																}
																position++
															}
														l364:
															{
																position366, tokenIndex366 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l367
																}
																position++
																goto l366
															l367:
																position, tokenIndex = position366, tokenIndex366
																if buffer[position] != rune('A') {
																	goto l360
																}
																position++
															}
														l366:
															if !_rules[rulews]() {
																goto l360
															}
															if !_rules[ruleLoc8]() {
																goto l360
															}
															{
																position368, tokenIndex368 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l368
																}
																if !_rules[ruleCopy8]() {
																	goto l368
																}
																goto l369
															l368:
																position, tokenIndex = position368, tokenIndex368
															}
														l369:
															{
																add(ruleAction56, position)
															}
															add(ruleSra, position361)
														}
														goto l308
													l360:
														position, tokenIndex = position308, tokenIndex308
														{
															position372 := position
															{
																position373, tokenIndex373 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l374
																}
																position++
																goto l373
															l374:
																position, tokenIndex = position373, tokenIndex373
																if buffer[position] != rune('S') {
																	goto l371
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
																	goto l371
																}
																position++
															}
														l375:
															{
																position377, tokenIndex377 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l378
																}
																position++
																goto l377
															l378:
																position, tokenIndex = position377, tokenIndex377
																if buffer[position] != rune('L') {
																	goto l371
																}
																position++
															}
														l377:
															if !_rules[rulews]() {
																goto l371
															}
															if !_rules[ruleLoc8]() {
																goto l371
															}
															{
																position379, tokenIndex379 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l379
																}
																if !_rules[ruleCopy8]() {
																	goto l379
																}
																goto l380
															l379:
																position, tokenIndex = position379, tokenIndex379
															}
														l380:
															{
																add(ruleAction57, position)
															}
															add(ruleSll, position372)
														}
														goto l308
													l371:
														position, tokenIndex = position308, tokenIndex308
														{
															position382 := position
															{
																position383, tokenIndex383 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l384
																}
																position++
																goto l383
															l384:
																position, tokenIndex = position383, tokenIndex383
																if buffer[position] != rune('S') {
																	goto l306
																}
																position++
															}
														l383:
															{
																position385, tokenIndex385 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l386
																}
																position++
																goto l385
															l386:
																position, tokenIndex = position385, tokenIndex385
																if buffer[position] != rune('R') {
																	goto l306
																}
																position++
															}
														l385:
															{
																position387, tokenIndex387 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l388
																}
																position++
																goto l387
															l388:
																position, tokenIndex = position387, tokenIndex387
																if buffer[position] != rune('L') {
																	goto l306
																}
																position++
															}
														l387:
															if !_rules[rulews]() {
																goto l306
															}
															if !_rules[ruleLoc8]() {
																goto l306
															}
															{
																position389, tokenIndex389 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l389
																}
																if !_rules[ruleCopy8]() {
																	goto l389
																}
																goto l390
															l389:
																position, tokenIndex = position389, tokenIndex389
															}
														l390:
															{
																add(ruleAction58, position)
															}
															add(ruleSrl, position382)
														}
													}
												l308:
													add(ruleRot, position307)
												}
												goto l305
											l306:
												position, tokenIndex = position305, tokenIndex305
												{
													switch buffer[position] {
													case 'S', 's':
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
																	goto l303
																}
																position++
															}
														l394:
															{
																position396, tokenIndex396 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l397
																}
																position++
																goto l396
															l397:
																position, tokenIndex = position396, tokenIndex396
																if buffer[position] != rune('E') {
																	goto l303
																}
																position++
															}
														l396:
															{
																position398, tokenIndex398 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l399
																}
																position++
																goto l398
															l399:
																position, tokenIndex = position398, tokenIndex398
																if buffer[position] != rune('T') {
																	goto l303
																}
																position++
															}
														l398:
															if !_rules[rulews]() {
																goto l303
															}
															if !_rules[ruleoctaldigit]() {
																goto l303
															}
															if !_rules[rulesep]() {
																goto l303
															}
															if !_rules[ruleLoc8]() {
																goto l303
															}
															{
																position400, tokenIndex400 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l400
																}
																if !_rules[ruleCopy8]() {
																	goto l400
																}
																goto l401
															l400:
																position, tokenIndex = position400, tokenIndex400
															}
														l401:
															{
																add(ruleAction61, position)
															}
															add(ruleSet, position393)
														}
														break
													case 'R', 'r':
														{
															position403 := position
															{
																position404, tokenIndex404 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l405
																}
																position++
																goto l404
															l405:
																position, tokenIndex = position404, tokenIndex404
																if buffer[position] != rune('R') {
																	goto l303
																}
																position++
															}
														l404:
															{
																position406, tokenIndex406 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l407
																}
																position++
																goto l406
															l407:
																position, tokenIndex = position406, tokenIndex406
																if buffer[position] != rune('E') {
																	goto l303
																}
																position++
															}
														l406:
															{
																position408, tokenIndex408 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l409
																}
																position++
																goto l408
															l409:
																position, tokenIndex = position408, tokenIndex408
																if buffer[position] != rune('S') {
																	goto l303
																}
																position++
															}
														l408:
															if !_rules[rulews]() {
																goto l303
															}
															if !_rules[ruleoctaldigit]() {
																goto l303
															}
															if !_rules[rulesep]() {
																goto l303
															}
															if !_rules[ruleLoc8]() {
																goto l303
															}
															{
																position410, tokenIndex410 := position, tokenIndex
																if !_rules[rulesep]() {
																	goto l410
																}
																if !_rules[ruleCopy8]() {
																	goto l410
																}
																goto l411
															l410:
																position, tokenIndex = position410, tokenIndex410
															}
														l411:
															{
																add(ruleAction60, position)
															}
															add(ruleRes, position403)
														}
														break
													default:
														{
															position413 := position
															{
																position414, tokenIndex414 := position, tokenIndex
																if buffer[position] != rune('b') {
																	goto l415
																}
																position++
																goto l414
															l415:
																position, tokenIndex = position414, tokenIndex414
																if buffer[position] != rune('B') {
																	goto l303
																}
																position++
															}
														l414:
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
																	goto l303
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
																	goto l303
																}
																position++
															}
														l418:
															if !_rules[rulews]() {
																goto l303
															}
															if !_rules[ruleoctaldigit]() {
																goto l303
															}
															if !_rules[rulesep]() {
																goto l303
															}
															if !_rules[ruleLoc8]() {
																goto l303
															}
															{
																add(ruleAction59, position)
															}
															add(ruleBit, position413)
														}
														break
													}
												}

											}
										l305:
											add(ruleBitOp, position304)
										}
										goto l89
									l303:
										position, tokenIndex = position89, tokenIndex89
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
															if buffer[position] != rune('r') {
																goto l428
															}
															position++
															goto l427
														l428:
															position, tokenIndex = position427, tokenIndex427
															if buffer[position] != rune('R') {
																goto l424
															}
															position++
														}
													l427:
														{
															position429, tokenIndex429 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l430
															}
															position++
															goto l429
														l430:
															position, tokenIndex = position429, tokenIndex429
															if buffer[position] != rune('E') {
																goto l424
															}
															position++
														}
													l429:
														{
															position431, tokenIndex431 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l432
															}
															position++
															goto l431
														l432:
															position, tokenIndex = position431, tokenIndex431
															if buffer[position] != rune('T') {
																goto l424
															}
															position++
														}
													l431:
														{
															position433, tokenIndex433 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l434
															}
															position++
															goto l433
														l434:
															position, tokenIndex = position433, tokenIndex433
															if buffer[position] != rune('N') {
																goto l424
															}
															position++
														}
													l433:
														add(rulePegText, position426)
													}
													{
														add(ruleAction76, position)
													}
													add(ruleRetn, position425)
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
															if buffer[position] != rune('r') {
																goto l440
															}
															position++
															goto l439
														l440:
															position, tokenIndex = position439, tokenIndex439
															if buffer[position] != rune('R') {
																goto l436
															}
															position++
														}
													l439:
														{
															position441, tokenIndex441 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l442
															}
															position++
															goto l441
														l442:
															position, tokenIndex = position441, tokenIndex441
															if buffer[position] != rune('E') {
																goto l436
															}
															position++
														}
													l441:
														{
															position443, tokenIndex443 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l444
															}
															position++
															goto l443
														l444:
															position, tokenIndex = position443, tokenIndex443
															if buffer[position] != rune('T') {
																goto l436
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
																goto l436
															}
															position++
														}
													l445:
														add(rulePegText, position438)
													}
													{
														add(ruleAction77, position)
													}
													add(ruleReti, position437)
												}
												goto l423
											l436:
												position, tokenIndex = position423, tokenIndex423
												{
													position449 := position
													{
														position450 := position
														{
															position451, tokenIndex451 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l452
															}
															position++
															goto l451
														l452:
															position, tokenIndex = position451, tokenIndex451
															if buffer[position] != rune('R') {
																goto l448
															}
															position++
														}
													l451:
														{
															position453, tokenIndex453 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l454
															}
															position++
															goto l453
														l454:
															position, tokenIndex = position453, tokenIndex453
															if buffer[position] != rune('R') {
																goto l448
															}
															position++
														}
													l453:
														{
															position455, tokenIndex455 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l456
															}
															position++
															goto l455
														l456:
															position, tokenIndex = position455, tokenIndex455
															if buffer[position] != rune('D') {
																goto l448
															}
															position++
														}
													l455:
														add(rulePegText, position450)
													}
													{
														add(ruleAction78, position)
													}
													add(ruleRrd, position449)
												}
												goto l423
											l448:
												position, tokenIndex = position423, tokenIndex423
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
														if buffer[position] != rune('0') {
															goto l458
														}
														position++
														add(rulePegText, position460)
													}
													{
														add(ruleAction80, position)
													}
													add(ruleIm0, position459)
												}
												goto l423
											l458:
												position, tokenIndex = position423, tokenIndex423
												{
													position467 := position
													{
														position468 := position
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
																goto l466
															}
															position++
														}
													l469:
														{
															position471, tokenIndex471 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l472
															}
															position++
															goto l471
														l472:
															position, tokenIndex = position471, tokenIndex471
															if buffer[position] != rune('M') {
																goto l466
															}
															position++
														}
													l471:
														if buffer[position] != rune(' ') {
															goto l466
														}
														position++
														if buffer[position] != rune('1') {
															goto l466
														}
														position++
														add(rulePegText, position468)
													}
													{
														add(ruleAction81, position)
													}
													add(ruleIm1, position467)
												}
												goto l423
											l466:
												position, tokenIndex = position423, tokenIndex423
												{
													position475 := position
													{
														position476 := position
														{
															position477, tokenIndex477 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l478
															}
															position++
															goto l477
														l478:
															position, tokenIndex = position477, tokenIndex477
															if buffer[position] != rune('I') {
																goto l474
															}
															position++
														}
													l477:
														{
															position479, tokenIndex479 := position, tokenIndex
															if buffer[position] != rune('m') {
																goto l480
															}
															position++
															goto l479
														l480:
															position, tokenIndex = position479, tokenIndex479
															if buffer[position] != rune('M') {
																goto l474
															}
															position++
														}
													l479:
														if buffer[position] != rune(' ') {
															goto l474
														}
														position++
														if buffer[position] != rune('2') {
															goto l474
														}
														position++
														add(rulePegText, position476)
													}
													{
														add(ruleAction82, position)
													}
													add(ruleIm2, position475)
												}
												goto l423
											l474:
												position, tokenIndex = position423, tokenIndex423
												{
													switch buffer[position] {
													case 'I', 'O', 'i', 'o':
														{
															position483 := position
															{
																position484, tokenIndex484 := position, tokenIndex
																{
																	position486 := position
																	{
																		position487 := position
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
																				goto l485
																			}
																			position++
																		}
																	l488:
																		{
																			position490, tokenIndex490 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l491
																			}
																			position++
																			goto l490
																		l491:
																			position, tokenIndex = position490, tokenIndex490
																			if buffer[position] != rune('N') {
																				goto l485
																			}
																			position++
																		}
																	l490:
																		{
																			position492, tokenIndex492 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l493
																			}
																			position++
																			goto l492
																		l493:
																			position, tokenIndex = position492, tokenIndex492
																			if buffer[position] != rune('I') {
																				goto l485
																			}
																			position++
																		}
																	l492:
																		{
																			position494, tokenIndex494 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l495
																			}
																			position++
																			goto l494
																		l495:
																			position, tokenIndex = position494, tokenIndex494
																			if buffer[position] != rune('R') {
																				goto l485
																			}
																			position++
																		}
																	l494:
																		add(rulePegText, position487)
																	}
																	{
																		add(ruleAction93, position)
																	}
																	add(ruleInir, position486)
																}
																goto l484
															l485:
																position, tokenIndex = position484, tokenIndex484
																{
																	position498 := position
																	{
																		position499 := position
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
																				goto l497
																			}
																			position++
																		}
																	l500:
																		{
																			position502, tokenIndex502 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l503
																			}
																			position++
																			goto l502
																		l503:
																			position, tokenIndex = position502, tokenIndex502
																			if buffer[position] != rune('N') {
																				goto l497
																			}
																			position++
																		}
																	l502:
																		{
																			position504, tokenIndex504 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l505
																			}
																			position++
																			goto l504
																		l505:
																			position, tokenIndex = position504, tokenIndex504
																			if buffer[position] != rune('I') {
																				goto l497
																			}
																			position++
																		}
																	l504:
																		add(rulePegText, position499)
																	}
																	{
																		add(ruleAction85, position)
																	}
																	add(ruleIni, position498)
																}
																goto l484
															l497:
																position, tokenIndex = position484, tokenIndex484
																{
																	position508 := position
																	{
																		position509 := position
																		{
																			position510, tokenIndex510 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l511
																			}
																			position++
																			goto l510
																		l511:
																			position, tokenIndex = position510, tokenIndex510
																			if buffer[position] != rune('O') {
																				goto l507
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
																				goto l507
																			}
																			position++
																		}
																	l512:
																		{
																			position514, tokenIndex514 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l515
																			}
																			position++
																			goto l514
																		l515:
																			position, tokenIndex = position514, tokenIndex514
																			if buffer[position] != rune('I') {
																				goto l507
																			}
																			position++
																		}
																	l514:
																		{
																			position516, tokenIndex516 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l517
																			}
																			position++
																			goto l516
																		l517:
																			position, tokenIndex = position516, tokenIndex516
																			if buffer[position] != rune('R') {
																				goto l507
																			}
																			position++
																		}
																	l516:
																		add(rulePegText, position509)
																	}
																	{
																		add(ruleAction94, position)
																	}
																	add(ruleOtir, position508)
																}
																goto l484
															l507:
																position, tokenIndex = position484, tokenIndex484
																{
																	position520 := position
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
																				goto l519
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
																				goto l519
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
																				goto l519
																			}
																			position++
																		}
																	l526:
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
																				goto l519
																			}
																			position++
																		}
																	l528:
																		add(rulePegText, position521)
																	}
																	{
																		add(ruleAction86, position)
																	}
																	add(ruleOuti, position520)
																}
																goto l484
															l519:
																position, tokenIndex = position484, tokenIndex484
																{
																	position532 := position
																	{
																		position533 := position
																		{
																			position534, tokenIndex534 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l535
																			}
																			position++
																			goto l534
																		l535:
																			position, tokenIndex = position534, tokenIndex534
																			if buffer[position] != rune('I') {
																				goto l531
																			}
																			position++
																		}
																	l534:
																		{
																			position536, tokenIndex536 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l537
																			}
																			position++
																			goto l536
																		l537:
																			position, tokenIndex = position536, tokenIndex536
																			if buffer[position] != rune('N') {
																				goto l531
																			}
																			position++
																		}
																	l536:
																		{
																			position538, tokenIndex538 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l539
																			}
																			position++
																			goto l538
																		l539:
																			position, tokenIndex = position538, tokenIndex538
																			if buffer[position] != rune('D') {
																				goto l531
																			}
																			position++
																		}
																	l538:
																		{
																			position540, tokenIndex540 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l541
																			}
																			position++
																			goto l540
																		l541:
																			position, tokenIndex = position540, tokenIndex540
																			if buffer[position] != rune('R') {
																				goto l531
																			}
																			position++
																		}
																	l540:
																		add(rulePegText, position533)
																	}
																	{
																		add(ruleAction97, position)
																	}
																	add(ruleIndr, position532)
																}
																goto l484
															l531:
																position, tokenIndex = position484, tokenIndex484
																{
																	position544 := position
																	{
																		position545 := position
																		{
																			position546, tokenIndex546 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l547
																			}
																			position++
																			goto l546
																		l547:
																			position, tokenIndex = position546, tokenIndex546
																			if buffer[position] != rune('I') {
																				goto l543
																			}
																			position++
																		}
																	l546:
																		{
																			position548, tokenIndex548 := position, tokenIndex
																			if buffer[position] != rune('n') {
																				goto l549
																			}
																			position++
																			goto l548
																		l549:
																			position, tokenIndex = position548, tokenIndex548
																			if buffer[position] != rune('N') {
																				goto l543
																			}
																			position++
																		}
																	l548:
																		{
																			position550, tokenIndex550 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l551
																			}
																			position++
																			goto l550
																		l551:
																			position, tokenIndex = position550, tokenIndex550
																			if buffer[position] != rune('D') {
																				goto l543
																			}
																			position++
																		}
																	l550:
																		add(rulePegText, position545)
																	}
																	{
																		add(ruleAction89, position)
																	}
																	add(ruleInd, position544)
																}
																goto l484
															l543:
																position, tokenIndex = position484, tokenIndex484
																{
																	position554 := position
																	{
																		position555 := position
																		{
																			position556, tokenIndex556 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l557
																			}
																			position++
																			goto l556
																		l557:
																			position, tokenIndex = position556, tokenIndex556
																			if buffer[position] != rune('O') {
																				goto l553
																			}
																			position++
																		}
																	l556:
																		{
																			position558, tokenIndex558 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l559
																			}
																			position++
																			goto l558
																		l559:
																			position, tokenIndex = position558, tokenIndex558
																			if buffer[position] != rune('T') {
																				goto l553
																			}
																			position++
																		}
																	l558:
																		{
																			position560, tokenIndex560 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l561
																			}
																			position++
																			goto l560
																		l561:
																			position, tokenIndex = position560, tokenIndex560
																			if buffer[position] != rune('D') {
																				goto l553
																			}
																			position++
																		}
																	l560:
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
																				goto l553
																			}
																			position++
																		}
																	l562:
																		add(rulePegText, position555)
																	}
																	{
																		add(ruleAction98, position)
																	}
																	add(ruleOtdr, position554)
																}
																goto l484
															l553:
																position, tokenIndex = position484, tokenIndex484
																{
																	position565 := position
																	{
																		position566 := position
																		{
																			position567, tokenIndex567 := position, tokenIndex
																			if buffer[position] != rune('o') {
																				goto l568
																			}
																			position++
																			goto l567
																		l568:
																			position, tokenIndex = position567, tokenIndex567
																			if buffer[position] != rune('O') {
																				goto l421
																			}
																			position++
																		}
																	l567:
																		{
																			position569, tokenIndex569 := position, tokenIndex
																			if buffer[position] != rune('u') {
																				goto l570
																			}
																			position++
																			goto l569
																		l570:
																			position, tokenIndex = position569, tokenIndex569
																			if buffer[position] != rune('U') {
																				goto l421
																			}
																			position++
																		}
																	l569:
																		{
																			position571, tokenIndex571 := position, tokenIndex
																			if buffer[position] != rune('t') {
																				goto l572
																			}
																			position++
																			goto l571
																		l572:
																			position, tokenIndex = position571, tokenIndex571
																			if buffer[position] != rune('T') {
																				goto l421
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
																				goto l421
																			}
																			position++
																		}
																	l573:
																		add(rulePegText, position566)
																	}
																	{
																		add(ruleAction90, position)
																	}
																	add(ruleOutd, position565)
																}
															}
														l484:
															add(ruleBlitIO, position483)
														}
														break
													case 'R', 'r':
														{
															position576 := position
															{
																position577 := position
																{
																	position578, tokenIndex578 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l579
																	}
																	position++
																	goto l578
																l579:
																	position, tokenIndex = position578, tokenIndex578
																	if buffer[position] != rune('R') {
																		goto l421
																	}
																	position++
																}
															l578:
																{
																	position580, tokenIndex580 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l581
																	}
																	position++
																	goto l580
																l581:
																	position, tokenIndex = position580, tokenIndex580
																	if buffer[position] != rune('L') {
																		goto l421
																	}
																	position++
																}
															l580:
																{
																	position582, tokenIndex582 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l583
																	}
																	position++
																	goto l582
																l583:
																	position, tokenIndex = position582, tokenIndex582
																	if buffer[position] != rune('D') {
																		goto l421
																	}
																	position++
																}
															l582:
																add(rulePegText, position577)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleRld, position576)
														}
														break
													case 'N', 'n':
														{
															position585 := position
															{
																position586 := position
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
																		goto l421
																	}
																	position++
																}
															l587:
																{
																	position589, tokenIndex589 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l590
																	}
																	position++
																	goto l589
																l590:
																	position, tokenIndex = position589, tokenIndex589
																	if buffer[position] != rune('E') {
																		goto l421
																	}
																	position++
																}
															l589:
																{
																	position591, tokenIndex591 := position, tokenIndex
																	if buffer[position] != rune('g') {
																		goto l592
																	}
																	position++
																	goto l591
																l592:
																	position, tokenIndex = position591, tokenIndex591
																	if buffer[position] != rune('G') {
																		goto l421
																	}
																	position++
																}
															l591:
																add(rulePegText, position586)
															}
															{
																add(ruleAction75, position)
															}
															add(ruleNeg, position585)
														}
														break
													default:
														{
															position594 := position
															{
																position595, tokenIndex595 := position, tokenIndex
																{
																	position597 := position
																	{
																		position598 := position
																		{
																			position599, tokenIndex599 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l600
																			}
																			position++
																			goto l599
																		l600:
																			position, tokenIndex = position599, tokenIndex599
																			if buffer[position] != rune('L') {
																				goto l596
																			}
																			position++
																		}
																	l599:
																		{
																			position601, tokenIndex601 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l602
																			}
																			position++
																			goto l601
																		l602:
																			position, tokenIndex = position601, tokenIndex601
																			if buffer[position] != rune('D') {
																				goto l596
																			}
																			position++
																		}
																	l601:
																		{
																			position603, tokenIndex603 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l604
																			}
																			position++
																			goto l603
																		l604:
																			position, tokenIndex = position603, tokenIndex603
																			if buffer[position] != rune('I') {
																				goto l596
																			}
																			position++
																		}
																	l603:
																		{
																			position605, tokenIndex605 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l606
																			}
																			position++
																			goto l605
																		l606:
																			position, tokenIndex = position605, tokenIndex605
																			if buffer[position] != rune('R') {
																				goto l596
																			}
																			position++
																		}
																	l605:
																		add(rulePegText, position598)
																	}
																	{
																		add(ruleAction91, position)
																	}
																	add(ruleLdir, position597)
																}
																goto l595
															l596:
																position, tokenIndex = position595, tokenIndex595
																{
																	position609 := position
																	{
																		position610 := position
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
																				goto l608
																			}
																			position++
																		}
																	l611:
																		{
																			position613, tokenIndex613 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l614
																			}
																			position++
																			goto l613
																		l614:
																			position, tokenIndex = position613, tokenIndex613
																			if buffer[position] != rune('D') {
																				goto l608
																			}
																			position++
																		}
																	l613:
																		{
																			position615, tokenIndex615 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l616
																			}
																			position++
																			goto l615
																		l616:
																			position, tokenIndex = position615, tokenIndex615
																			if buffer[position] != rune('I') {
																				goto l608
																			}
																			position++
																		}
																	l615:
																		add(rulePegText, position610)
																	}
																	{
																		add(ruleAction83, position)
																	}
																	add(ruleLdi, position609)
																}
																goto l595
															l608:
																position, tokenIndex = position595, tokenIndex595
																{
																	position619 := position
																	{
																		position620 := position
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
																				goto l618
																			}
																			position++
																		}
																	l621:
																		{
																			position623, tokenIndex623 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l624
																			}
																			position++
																			goto l623
																		l624:
																			position, tokenIndex = position623, tokenIndex623
																			if buffer[position] != rune('P') {
																				goto l618
																			}
																			position++
																		}
																	l623:
																		{
																			position625, tokenIndex625 := position, tokenIndex
																			if buffer[position] != rune('i') {
																				goto l626
																			}
																			position++
																			goto l625
																		l626:
																			position, tokenIndex = position625, tokenIndex625
																			if buffer[position] != rune('I') {
																				goto l618
																			}
																			position++
																		}
																	l625:
																		{
																			position627, tokenIndex627 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l628
																			}
																			position++
																			goto l627
																		l628:
																			position, tokenIndex = position627, tokenIndex627
																			if buffer[position] != rune('R') {
																				goto l618
																			}
																			position++
																		}
																	l627:
																		add(rulePegText, position620)
																	}
																	{
																		add(ruleAction92, position)
																	}
																	add(ruleCpir, position619)
																}
																goto l595
															l618:
																position, tokenIndex = position595, tokenIndex595
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
																			if buffer[position] != rune('i') {
																				goto l638
																			}
																			position++
																			goto l637
																		l638:
																			position, tokenIndex = position637, tokenIndex637
																			if buffer[position] != rune('I') {
																				goto l630
																			}
																			position++
																		}
																	l637:
																		add(rulePegText, position632)
																	}
																	{
																		add(ruleAction84, position)
																	}
																	add(ruleCpi, position631)
																}
																goto l595
															l630:
																position, tokenIndex = position595, tokenIndex595
																{
																	position641 := position
																	{
																		position642 := position
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
																				goto l640
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
																				goto l640
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
																				goto l640
																			}
																			position++
																		}
																	l647:
																		{
																			position649, tokenIndex649 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l650
																			}
																			position++
																			goto l649
																		l650:
																			position, tokenIndex = position649, tokenIndex649
																			if buffer[position] != rune('R') {
																				goto l640
																			}
																			position++
																		}
																	l649:
																		add(rulePegText, position642)
																	}
																	{
																		add(ruleAction95, position)
																	}
																	add(ruleLddr, position641)
																}
																goto l595
															l640:
																position, tokenIndex = position595, tokenIndex595
																{
																	position653 := position
																	{
																		position654 := position
																		{
																			position655, tokenIndex655 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l656
																			}
																			position++
																			goto l655
																		l656:
																			position, tokenIndex = position655, tokenIndex655
																			if buffer[position] != rune('L') {
																				goto l652
																			}
																			position++
																		}
																	l655:
																		{
																			position657, tokenIndex657 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l658
																			}
																			position++
																			goto l657
																		l658:
																			position, tokenIndex = position657, tokenIndex657
																			if buffer[position] != rune('D') {
																				goto l652
																			}
																			position++
																		}
																	l657:
																		{
																			position659, tokenIndex659 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l660
																			}
																			position++
																			goto l659
																		l660:
																			position, tokenIndex = position659, tokenIndex659
																			if buffer[position] != rune('D') {
																				goto l652
																			}
																			position++
																		}
																	l659:
																		add(rulePegText, position654)
																	}
																	{
																		add(ruleAction87, position)
																	}
																	add(ruleLdd, position653)
																}
																goto l595
															l652:
																position, tokenIndex = position595, tokenIndex595
																{
																	position663 := position
																	{
																		position664 := position
																		{
																			position665, tokenIndex665 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l666
																			}
																			position++
																			goto l665
																		l666:
																			position, tokenIndex = position665, tokenIndex665
																			if buffer[position] != rune('C') {
																				goto l662
																			}
																			position++
																		}
																	l665:
																		{
																			position667, tokenIndex667 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l668
																			}
																			position++
																			goto l667
																		l668:
																			position, tokenIndex = position667, tokenIndex667
																			if buffer[position] != rune('P') {
																				goto l662
																			}
																			position++
																		}
																	l667:
																		{
																			position669, tokenIndex669 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l670
																			}
																			position++
																			goto l669
																		l670:
																			position, tokenIndex = position669, tokenIndex669
																			if buffer[position] != rune('D') {
																				goto l662
																			}
																			position++
																		}
																	l669:
																		{
																			position671, tokenIndex671 := position, tokenIndex
																			if buffer[position] != rune('r') {
																				goto l672
																			}
																			position++
																			goto l671
																		l672:
																			position, tokenIndex = position671, tokenIndex671
																			if buffer[position] != rune('R') {
																				goto l662
																			}
																			position++
																		}
																	l671:
																		add(rulePegText, position664)
																	}
																	{
																		add(ruleAction96, position)
																	}
																	add(ruleCpdr, position663)
																}
																goto l595
															l662:
																position, tokenIndex = position595, tokenIndex595
																{
																	position674 := position
																	{
																		position675 := position
																		{
																			position676, tokenIndex676 := position, tokenIndex
																			if buffer[position] != rune('c') {
																				goto l677
																			}
																			position++
																			goto l676
																		l677:
																			position, tokenIndex = position676, tokenIndex676
																			if buffer[position] != rune('C') {
																				goto l421
																			}
																			position++
																		}
																	l676:
																		{
																			position678, tokenIndex678 := position, tokenIndex
																			if buffer[position] != rune('p') {
																				goto l679
																			}
																			position++
																			goto l678
																		l679:
																			position, tokenIndex = position678, tokenIndex678
																			if buffer[position] != rune('P') {
																				goto l421
																			}
																			position++
																		}
																	l678:
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
																				goto l421
																			}
																			position++
																		}
																	l680:
																		add(rulePegText, position675)
																	}
																	{
																		add(ruleAction88, position)
																	}
																	add(ruleCpd, position674)
																}
															}
														l595:
															add(ruleBlit, position594)
														}
														break
													}
												}

											}
										l423:
											add(ruleEDSimple, position422)
										}
										goto l89
									l421:
										position, tokenIndex = position89, tokenIndex89
										{
											position684 := position
											{
												position685, tokenIndex685 := position, tokenIndex
												{
													position687 := position
													{
														position688 := position
														{
															position689, tokenIndex689 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l690
															}
															position++
															goto l689
														l690:
															position, tokenIndex = position689, tokenIndex689
															if buffer[position] != rune('R') {
																goto l686
															}
															position++
														}
													l689:
														{
															position691, tokenIndex691 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l692
															}
															position++
															goto l691
														l692:
															position, tokenIndex = position691, tokenIndex691
															if buffer[position] != rune('L') {
																goto l686
															}
															position++
														}
													l691:
														{
															position693, tokenIndex693 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l694
															}
															position++
															goto l693
														l694:
															position, tokenIndex = position693, tokenIndex693
															if buffer[position] != rune('C') {
																goto l686
															}
															position++
														}
													l693:
														{
															position695, tokenIndex695 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l696
															}
															position++
															goto l695
														l696:
															position, tokenIndex = position695, tokenIndex695
															if buffer[position] != rune('A') {
																goto l686
															}
															position++
														}
													l695:
														add(rulePegText, position688)
													}
													{
														add(ruleAction64, position)
													}
													add(ruleRlca, position687)
												}
												goto l685
											l686:
												position, tokenIndex = position685, tokenIndex685
												{
													position699 := position
													{
														position700 := position
														{
															position701, tokenIndex701 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l702
															}
															position++
															goto l701
														l702:
															position, tokenIndex = position701, tokenIndex701
															if buffer[position] != rune('R') {
																goto l698
															}
															position++
														}
													l701:
														{
															position703, tokenIndex703 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l704
															}
															position++
															goto l703
														l704:
															position, tokenIndex = position703, tokenIndex703
															if buffer[position] != rune('R') {
																goto l698
															}
															position++
														}
													l703:
														{
															position705, tokenIndex705 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l706
															}
															position++
															goto l705
														l706:
															position, tokenIndex = position705, tokenIndex705
															if buffer[position] != rune('C') {
																goto l698
															}
															position++
														}
													l705:
														{
															position707, tokenIndex707 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l708
															}
															position++
															goto l707
														l708:
															position, tokenIndex = position707, tokenIndex707
															if buffer[position] != rune('A') {
																goto l698
															}
															position++
														}
													l707:
														add(rulePegText, position700)
													}
													{
														add(ruleAction65, position)
													}
													add(ruleRrca, position699)
												}
												goto l685
											l698:
												position, tokenIndex = position685, tokenIndex685
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
																goto l710
															}
															position++
														}
													l713:
														{
															position715, tokenIndex715 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l716
															}
															position++
															goto l715
														l716:
															position, tokenIndex = position715, tokenIndex715
															if buffer[position] != rune('L') {
																goto l710
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
																goto l710
															}
															position++
														}
													l717:
														add(rulePegText, position712)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleRla, position711)
												}
												goto l685
											l710:
												position, tokenIndex = position685, tokenIndex685
												{
													position721 := position
													{
														position722 := position
														{
															position723, tokenIndex723 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l724
															}
															position++
															goto l723
														l724:
															position, tokenIndex = position723, tokenIndex723
															if buffer[position] != rune('D') {
																goto l720
															}
															position++
														}
													l723:
														{
															position725, tokenIndex725 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l726
															}
															position++
															goto l725
														l726:
															position, tokenIndex = position725, tokenIndex725
															if buffer[position] != rune('A') {
																goto l720
															}
															position++
														}
													l725:
														{
															position727, tokenIndex727 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l728
															}
															position++
															goto l727
														l728:
															position, tokenIndex = position727, tokenIndex727
															if buffer[position] != rune('A') {
																goto l720
															}
															position++
														}
													l727:
														add(rulePegText, position722)
													}
													{
														add(ruleAction68, position)
													}
													add(ruleDaa, position721)
												}
												goto l685
											l720:
												position, tokenIndex = position685, tokenIndex685
												{
													position731 := position
													{
														position732 := position
														{
															position733, tokenIndex733 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l734
															}
															position++
															goto l733
														l734:
															position, tokenIndex = position733, tokenIndex733
															if buffer[position] != rune('C') {
																goto l730
															}
															position++
														}
													l733:
														{
															position735, tokenIndex735 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l736
															}
															position++
															goto l735
														l736:
															position, tokenIndex = position735, tokenIndex735
															if buffer[position] != rune('P') {
																goto l730
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
																goto l730
															}
															position++
														}
													l737:
														add(rulePegText, position732)
													}
													{
														add(ruleAction69, position)
													}
													add(ruleCpl, position731)
												}
												goto l685
											l730:
												position, tokenIndex = position685, tokenIndex685
												{
													position741 := position
													{
														position742 := position
														{
															position743, tokenIndex743 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l744
															}
															position++
															goto l743
														l744:
															position, tokenIndex = position743, tokenIndex743
															if buffer[position] != rune('E') {
																goto l740
															}
															position++
														}
													l743:
														{
															position745, tokenIndex745 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l746
															}
															position++
															goto l745
														l746:
															position, tokenIndex = position745, tokenIndex745
															if buffer[position] != rune('X') {
																goto l740
															}
															position++
														}
													l745:
														{
															position747, tokenIndex747 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l748
															}
															position++
															goto l747
														l748:
															position, tokenIndex = position747, tokenIndex747
															if buffer[position] != rune('X') {
																goto l740
															}
															position++
														}
													l747:
														add(rulePegText, position742)
													}
													{
														add(ruleAction72, position)
													}
													add(ruleExx, position741)
												}
												goto l685
											l740:
												position, tokenIndex = position685, tokenIndex685
												{
													switch buffer[position] {
													case 'E', 'e':
														{
															position751 := position
															{
																position752 := position
																{
																	position753, tokenIndex753 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l754
																	}
																	position++
																	goto l753
																l754:
																	position, tokenIndex = position753, tokenIndex753
																	if buffer[position] != rune('E') {
																		goto l683
																	}
																	position++
																}
															l753:
																{
																	position755, tokenIndex755 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l756
																	}
																	position++
																	goto l755
																l756:
																	position, tokenIndex = position755, tokenIndex755
																	if buffer[position] != rune('I') {
																		goto l683
																	}
																	position++
																}
															l755:
																add(rulePegText, position752)
															}
															{
																add(ruleAction74, position)
															}
															add(ruleEi, position751)
														}
														break
													case 'D', 'd':
														{
															position758 := position
															{
																position759 := position
																{
																	position760, tokenIndex760 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l761
																	}
																	position++
																	goto l760
																l761:
																	position, tokenIndex = position760, tokenIndex760
																	if buffer[position] != rune('D') {
																		goto l683
																	}
																	position++
																}
															l760:
																{
																	position762, tokenIndex762 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l763
																	}
																	position++
																	goto l762
																l763:
																	position, tokenIndex = position762, tokenIndex762
																	if buffer[position] != rune('I') {
																		goto l683
																	}
																	position++
																}
															l762:
																add(rulePegText, position759)
															}
															{
																add(ruleAction73, position)
															}
															add(ruleDi, position758)
														}
														break
													case 'C', 'c':
														{
															position765 := position
															{
																position766 := position
																{
																	position767, tokenIndex767 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l768
																	}
																	position++
																	goto l767
																l768:
																	position, tokenIndex = position767, tokenIndex767
																	if buffer[position] != rune('C') {
																		goto l683
																	}
																	position++
																}
															l767:
																{
																	position769, tokenIndex769 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l770
																	}
																	position++
																	goto l769
																l770:
																	position, tokenIndex = position769, tokenIndex769
																	if buffer[position] != rune('C') {
																		goto l683
																	}
																	position++
																}
															l769:
																{
																	position771, tokenIndex771 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l772
																	}
																	position++
																	goto l771
																l772:
																	position, tokenIndex = position771, tokenIndex771
																	if buffer[position] != rune('F') {
																		goto l683
																	}
																	position++
																}
															l771:
																add(rulePegText, position766)
															}
															{
																add(ruleAction71, position)
															}
															add(ruleCcf, position765)
														}
														break
													case 'S', 's':
														{
															position774 := position
															{
																position775 := position
																{
																	position776, tokenIndex776 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l777
																	}
																	position++
																	goto l776
																l777:
																	position, tokenIndex = position776, tokenIndex776
																	if buffer[position] != rune('S') {
																		goto l683
																	}
																	position++
																}
															l776:
																{
																	position778, tokenIndex778 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l779
																	}
																	position++
																	goto l778
																l779:
																	position, tokenIndex = position778, tokenIndex778
																	if buffer[position] != rune('C') {
																		goto l683
																	}
																	position++
																}
															l778:
																{
																	position780, tokenIndex780 := position, tokenIndex
																	if buffer[position] != rune('f') {
																		goto l781
																	}
																	position++
																	goto l780
																l781:
																	position, tokenIndex = position780, tokenIndex780
																	if buffer[position] != rune('F') {
																		goto l683
																	}
																	position++
																}
															l780:
																add(rulePegText, position775)
															}
															{
																add(ruleAction70, position)
															}
															add(ruleScf, position774)
														}
														break
													case 'R', 'r':
														{
															position783 := position
															{
																position784 := position
																{
																	position785, tokenIndex785 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l786
																	}
																	position++
																	goto l785
																l786:
																	position, tokenIndex = position785, tokenIndex785
																	if buffer[position] != rune('R') {
																		goto l683
																	}
																	position++
																}
															l785:
																{
																	position787, tokenIndex787 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l788
																	}
																	position++
																	goto l787
																l788:
																	position, tokenIndex = position787, tokenIndex787
																	if buffer[position] != rune('R') {
																		goto l683
																	}
																	position++
																}
															l787:
																{
																	position789, tokenIndex789 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l790
																	}
																	position++
																	goto l789
																l790:
																	position, tokenIndex = position789, tokenIndex789
																	if buffer[position] != rune('A') {
																		goto l683
																	}
																	position++
																}
															l789:
																add(rulePegText, position784)
															}
															{
																add(ruleAction67, position)
															}
															add(ruleRra, position783)
														}
														break
													case 'H', 'h':
														{
															position792 := position
															{
																position793 := position
																{
																	position794, tokenIndex794 := position, tokenIndex
																	if buffer[position] != rune('h') {
																		goto l795
																	}
																	position++
																	goto l794
																l795:
																	position, tokenIndex = position794, tokenIndex794
																	if buffer[position] != rune('H') {
																		goto l683
																	}
																	position++
																}
															l794:
																{
																	position796, tokenIndex796 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l797
																	}
																	position++
																	goto l796
																l797:
																	position, tokenIndex = position796, tokenIndex796
																	if buffer[position] != rune('A') {
																		goto l683
																	}
																	position++
																}
															l796:
																{
																	position798, tokenIndex798 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l799
																	}
																	position++
																	goto l798
																l799:
																	position, tokenIndex = position798, tokenIndex798
																	if buffer[position] != rune('L') {
																		goto l683
																	}
																	position++
																}
															l798:
																{
																	position800, tokenIndex800 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l801
																	}
																	position++
																	goto l800
																l801:
																	position, tokenIndex = position800, tokenIndex800
																	if buffer[position] != rune('T') {
																		goto l683
																	}
																	position++
																}
															l800:
																add(rulePegText, position793)
															}
															{
																add(ruleAction63, position)
															}
															add(ruleHalt, position792)
														}
														break
													default:
														{
															position803 := position
															{
																position804 := position
																{
																	position805, tokenIndex805 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l806
																	}
																	position++
																	goto l805
																l806:
																	position, tokenIndex = position805, tokenIndex805
																	if buffer[position] != rune('N') {
																		goto l683
																	}
																	position++
																}
															l805:
																{
																	position807, tokenIndex807 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l808
																	}
																	position++
																	goto l807
																l808:
																	position, tokenIndex = position807, tokenIndex807
																	if buffer[position] != rune('O') {
																		goto l683
																	}
																	position++
																}
															l807:
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
																		goto l683
																	}
																	position++
																}
															l809:
																add(rulePegText, position804)
															}
															{
																add(ruleAction62, position)
															}
															add(ruleNop, position803)
														}
														break
													}
												}

											}
										l685:
											add(ruleSimple, position684)
										}
										goto l89
									l683:
										position, tokenIndex = position89, tokenIndex89
										{
											position813 := position
											{
												position814, tokenIndex814 := position, tokenIndex
												{
													position816 := position
													{
														position817, tokenIndex817 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l818
														}
														position++
														goto l817
													l818:
														position, tokenIndex = position817, tokenIndex817
														if buffer[position] != rune('R') {
															goto l815
														}
														position++
													}
												l817:
													{
														position819, tokenIndex819 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l820
														}
														position++
														goto l819
													l820:
														position, tokenIndex = position819, tokenIndex819
														if buffer[position] != rune('S') {
															goto l815
														}
														position++
													}
												l819:
													{
														position821, tokenIndex821 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l822
														}
														position++
														goto l821
													l822:
														position, tokenIndex = position821, tokenIndex821
														if buffer[position] != rune('T') {
															goto l815
														}
														position++
													}
												l821:
													if !_rules[rulews]() {
														goto l815
													}
													if !_rules[rulen]() {
														goto l815
													}
													{
														add(ruleAction99, position)
													}
													add(ruleRst, position816)
												}
												goto l814
											l815:
												position, tokenIndex = position814, tokenIndex814
												{
													position825 := position
													{
														position826, tokenIndex826 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l827
														}
														position++
														goto l826
													l827:
														position, tokenIndex = position826, tokenIndex826
														if buffer[position] != rune('J') {
															goto l824
														}
														position++
													}
												l826:
													{
														position828, tokenIndex828 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l829
														}
														position++
														goto l828
													l829:
														position, tokenIndex = position828, tokenIndex828
														if buffer[position] != rune('P') {
															goto l824
														}
														position++
													}
												l828:
													if !_rules[rulews]() {
														goto l824
													}
													{
														position830, tokenIndex830 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l830
														}
														if !_rules[rulesep]() {
															goto l830
														}
														goto l831
													l830:
														position, tokenIndex = position830, tokenIndex830
													}
												l831:
													if !_rules[ruleSrc16]() {
														goto l824
													}
													{
														add(ruleAction102, position)
													}
													add(ruleJp, position825)
												}
												goto l814
											l824:
												position, tokenIndex = position814, tokenIndex814
												{
													switch buffer[position] {
													case 'D', 'd':
														{
															position834 := position
															{
																position835, tokenIndex835 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l836
																}
																position++
																goto l835
															l836:
																position, tokenIndex = position835, tokenIndex835
																if buffer[position] != rune('D') {
																	goto l812
																}
																position++
															}
														l835:
															{
																position837, tokenIndex837 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l838
																}
																position++
																goto l837
															l838:
																position, tokenIndex = position837, tokenIndex837
																if buffer[position] != rune('J') {
																	goto l812
																}
																position++
															}
														l837:
															{
																position839, tokenIndex839 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l840
																}
																position++
																goto l839
															l840:
																position, tokenIndex = position839, tokenIndex839
																if buffer[position] != rune('N') {
																	goto l812
																}
																position++
															}
														l839:
															{
																position841, tokenIndex841 := position, tokenIndex
																if buffer[position] != rune('z') {
																	goto l842
																}
																position++
																goto l841
															l842:
																position, tokenIndex = position841, tokenIndex841
																if buffer[position] != rune('Z') {
																	goto l812
																}
																position++
															}
														l841:
															if !_rules[rulews]() {
																goto l812
															}
															if !_rules[ruledisp]() {
																goto l812
															}
															{
																add(ruleAction104, position)
															}
															add(ruleDjnz, position834)
														}
														break
													case 'J', 'j':
														{
															position844 := position
															{
																position845, tokenIndex845 := position, tokenIndex
																if buffer[position] != rune('j') {
																	goto l846
																}
																position++
																goto l845
															l846:
																position, tokenIndex = position845, tokenIndex845
																if buffer[position] != rune('J') {
																	goto l812
																}
																position++
															}
														l845:
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
																	goto l812
																}
																position++
															}
														l847:
															if !_rules[rulews]() {
																goto l812
															}
															{
																position849, tokenIndex849 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l849
																}
																if !_rules[rulesep]() {
																	goto l849
																}
																goto l850
															l849:
																position, tokenIndex = position849, tokenIndex849
															}
														l850:
															if !_rules[ruledisp]() {
																goto l812
															}
															{
																add(ruleAction103, position)
															}
															add(ruleJr, position844)
														}
														break
													case 'R', 'r':
														{
															position852 := position
															{
																position853, tokenIndex853 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l854
																}
																position++
																goto l853
															l854:
																position, tokenIndex = position853, tokenIndex853
																if buffer[position] != rune('R') {
																	goto l812
																}
																position++
															}
														l853:
															{
																position855, tokenIndex855 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l856
																}
																position++
																goto l855
															l856:
																position, tokenIndex = position855, tokenIndex855
																if buffer[position] != rune('E') {
																	goto l812
																}
																position++
															}
														l855:
															{
																position857, tokenIndex857 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l858
																}
																position++
																goto l857
															l858:
																position, tokenIndex = position857, tokenIndex857
																if buffer[position] != rune('T') {
																	goto l812
																}
																position++
															}
														l857:
															{
																position859, tokenIndex859 := position, tokenIndex
																if !_rules[rulews]() {
																	goto l859
																}
																if !_rules[rulecc]() {
																	goto l859
																}
																goto l860
															l859:
																position, tokenIndex = position859, tokenIndex859
															}
														l860:
															{
																add(ruleAction101, position)
															}
															add(ruleRet, position852)
														}
														break
													default:
														{
															position862 := position
															{
																position863, tokenIndex863 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l864
																}
																position++
																goto l863
															l864:
																position, tokenIndex = position863, tokenIndex863
																if buffer[position] != rune('C') {
																	goto l812
																}
																position++
															}
														l863:
															{
																position865, tokenIndex865 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l866
																}
																position++
																goto l865
															l866:
																position, tokenIndex = position865, tokenIndex865
																if buffer[position] != rune('A') {
																	goto l812
																}
																position++
															}
														l865:
															{
																position867, tokenIndex867 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l868
																}
																position++
																goto l867
															l868:
																position, tokenIndex = position867, tokenIndex867
																if buffer[position] != rune('L') {
																	goto l812
																}
																position++
															}
														l867:
															{
																position869, tokenIndex869 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l870
																}
																position++
																goto l869
															l870:
																position, tokenIndex = position869, tokenIndex869
																if buffer[position] != rune('L') {
																	goto l812
																}
																position++
															}
														l869:
															if !_rules[rulews]() {
																goto l812
															}
															{
																position871, tokenIndex871 := position, tokenIndex
																if !_rules[rulecc]() {
																	goto l871
																}
																if !_rules[rulesep]() {
																	goto l871
																}
																goto l872
															l871:
																position, tokenIndex = position871, tokenIndex871
															}
														l872:
															if !_rules[ruleSrc16]() {
																goto l812
															}
															{
																add(ruleAction100, position)
															}
															add(ruleCall, position862)
														}
														break
													}
												}

											}
										l814:
											add(ruleJump, position813)
										}
										goto l89
									l812:
										position, tokenIndex = position89, tokenIndex89
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
													if !_rules[rulews]() {
														goto l876
													}
													if !_rules[ruleReg8]() {
														goto l876
													}
													if !_rules[rulesep]() {
														goto l876
													}
													if !_rules[rulePort]() {
														goto l876
													}
													{
														add(ruleAction105, position)
													}
													add(ruleIN, position877)
												}
												goto l875
											l876:
												position, tokenIndex = position875, tokenIndex875
												{
													position883 := position
													{
														position884, tokenIndex884 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l885
														}
														position++
														goto l884
													l885:
														position, tokenIndex = position884, tokenIndex884
														if buffer[position] != rune('O') {
															goto l15
														}
														position++
													}
												l884:
													{
														position886, tokenIndex886 := position, tokenIndex
														if buffer[position] != rune('u') {
															goto l887
														}
														position++
														goto l886
													l887:
														position, tokenIndex = position886, tokenIndex886
														if buffer[position] != rune('U') {
															goto l15
														}
														position++
													}
												l886:
													{
														position888, tokenIndex888 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l889
														}
														position++
														goto l888
													l889:
														position, tokenIndex = position888, tokenIndex888
														if buffer[position] != rune('T') {
															goto l15
														}
														position++
													}
												l888:
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
													add(ruleOUT, position883)
												}
											}
										l875:
											add(ruleIO, position874)
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
						position891, tokenIndex891 := position, tokenIndex
						if !_rules[rulews]() {
							goto l891
						}
						goto l892
					l891:
						position, tokenIndex = position891, tokenIndex891
					}
				l892:
					{
						position893, tokenIndex893 := position, tokenIndex
						{
							position895 := position
							{
								position896, tokenIndex896 := position, tokenIndex
								if buffer[position] != rune(';') {
									goto l897
								}
								position++
								goto l896
							l897:
								position, tokenIndex = position896, tokenIndex896
								if buffer[position] != rune('#') {
									goto l893
								}
								position++
							}
						l896:
						l898:
							{
								position899, tokenIndex899 := position, tokenIndex
								{
									position900, tokenIndex900 := position, tokenIndex
									if buffer[position] != rune('\n') {
										goto l900
									}
									position++
									goto l899
								l900:
									position, tokenIndex = position900, tokenIndex900
								}
								if !matchDot() {
									goto l899
								}
								goto l898
							l899:
								position, tokenIndex = position899, tokenIndex899
							}
							add(ruleComment, position895)
						}
						goto l894
					l893:
						position, tokenIndex = position893, tokenIndex893
					}
				l894:
					{
						position901, tokenIndex901 := position, tokenIndex
						if !_rules[rulews]() {
							goto l901
						}
						goto l902
					l901:
						position, tokenIndex = position901, tokenIndex901
					}
				l902:
					{
						position903, tokenIndex903 := position, tokenIndex
						{
							position905, tokenIndex905 := position, tokenIndex
							if buffer[position] != rune('\r') {
								goto l905
							}
							position++
							goto l906
						l905:
							position, tokenIndex = position905, tokenIndex905
						}
					l906:
						if buffer[position] != rune('\n') {
							goto l904
						}
						position++
						goto l903
					l904:
						position, tokenIndex = position903, tokenIndex903
						if buffer[position] != rune(':') {
							goto l0
						}
						position++
					}
				l903:
					{
						add(ruleAction0, position)
					}
					add(ruleLine, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position908 := position
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
								if !_rules[ruleLabelText]() {
									goto l911
								}
								if buffer[position] != rune(':') {
									goto l911
								}
								position++
								{
									position914, tokenIndex914 := position, tokenIndex
									if !_rules[rulews]() {
										goto l914
									}
									goto l915
								l914:
									position, tokenIndex = position914, tokenIndex914
								}
							l915:
								{
									add(ruleAction5, position)
								}
								add(ruleLabelDefn, position913)
							}
							goto l912
						l911:
							position, tokenIndex = position911, tokenIndex911
						}
					l912:
					l917:
						{
							position918, tokenIndex918 := position, tokenIndex
							if !_rules[rulews]() {
								goto l918
							}
							goto l917
						l918:
							position, tokenIndex = position918, tokenIndex918
						}
						{
							position919, tokenIndex919 := position, tokenIndex
							{
								position921 := position
								{
									position922, tokenIndex922 := position, tokenIndex
									{
										position924 := position
										{
											position925, tokenIndex925 := position, tokenIndex
											{
												position927 := position
												{
													position928, tokenIndex928 := position, tokenIndex
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
															goto l929
														}
														position++
													}
												l930:
													{
														position932, tokenIndex932 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l933
														}
														position++
														goto l932
													l933:
														position, tokenIndex = position932, tokenIndex932
														if buffer[position] != rune('E') {
															goto l929
														}
														position++
													}
												l932:
													{
														position934, tokenIndex934 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l935
														}
														position++
														goto l934
													l935:
														position, tokenIndex = position934, tokenIndex934
														if buffer[position] != rune('F') {
															goto l929
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
															goto l929
														}
														position++
													}
												l936:
													goto l928
												l929:
													position, tokenIndex = position928, tokenIndex928
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
															goto l926
														}
														position++
													}
												l938:
													{
														position940, tokenIndex940 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l941
														}
														position++
														goto l940
													l941:
														position, tokenIndex = position940, tokenIndex940
														if buffer[position] != rune('B') {
															goto l926
														}
														position++
													}
												l940:
												}
											l928:
												if !_rules[rulews]() {
													goto l926
												}
												if !_rules[rulen]() {
													goto l926
												}
												{
													add(ruleAction2, position)
												}
												add(ruleDefb, position927)
											}
											goto l925
										l926:
											position, tokenIndex = position925, tokenIndex925
											{
												position944 := position
												{
													position945, tokenIndex945 := position, tokenIndex
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
															goto l946
														}
														position++
													}
												l947:
													{
														position949, tokenIndex949 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l950
														}
														position++
														goto l949
													l950:
														position, tokenIndex = position949, tokenIndex949
														if buffer[position] != rune('E') {
															goto l946
														}
														position++
													}
												l949:
													{
														position951, tokenIndex951 := position, tokenIndex
														if buffer[position] != rune('f') {
															goto l952
														}
														position++
														goto l951
													l952:
														position, tokenIndex = position951, tokenIndex951
														if buffer[position] != rune('F') {
															goto l946
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
															goto l946
														}
														position++
													}
												l953:
													goto l945
												l946:
													position, tokenIndex = position945, tokenIndex945
													{
														position955, tokenIndex955 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l956
														}
														position++
														goto l955
													l956:
														position, tokenIndex = position955, tokenIndex955
														if buffer[position] != rune('D') {
															goto l943
														}
														position++
													}
												l955:
													{
														position957, tokenIndex957 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l958
														}
														position++
														goto l957
													l958:
														position, tokenIndex = position957, tokenIndex957
														if buffer[position] != rune('S') {
															goto l943
														}
														position++
													}
												l957:
												}
											l945:
												if !_rules[rulews]() {
													goto l943
												}
												if !_rules[rulen]() {
													goto l943
												}
												{
													add(ruleAction4, position)
												}
												add(ruleDefs, position944)
											}
											goto l925
										l943:
											position, tokenIndex = position925, tokenIndex925
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position961 := position
														{
															position962, tokenIndex962 := position, tokenIndex
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
																	goto l963
																}
																position++
															}
														l964:
															{
																position966, tokenIndex966 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l967
																}
																position++
																goto l966
															l967:
																position, tokenIndex = position966, tokenIndex966
																if buffer[position] != rune('E') {
																	goto l963
																}
																position++
															}
														l966:
															{
																position968, tokenIndex968 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l969
																}
																position++
																goto l968
															l969:
																position, tokenIndex = position968, tokenIndex968
																if buffer[position] != rune('F') {
																	goto l963
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
																	goto l963
																}
																position++
															}
														l970:
															goto l962
														l963:
															position, tokenIndex = position962, tokenIndex962
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
																	goto l923
																}
																position++
															}
														l972:
															{
																position974, tokenIndex974 := position, tokenIndex
																if buffer[position] != rune('w') {
																	goto l975
																}
																position++
																goto l974
															l975:
																position, tokenIndex = position974, tokenIndex974
																if buffer[position] != rune('W') {
																	goto l923
																}
																position++
															}
														l974:
														}
													l962:
														if !_rules[rulews]() {
															goto l923
														}
														if !_rules[rulenn]() {
															goto l923
														}
														{
															add(ruleAction3, position)
														}
														add(ruleDefw, position961)
													}
													break
												case 'O', 'o':
													{
														position977 := position
														{
															position978, tokenIndex978 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l979
															}
															position++
															goto l978
														l979:
															position, tokenIndex = position978, tokenIndex978
															if buffer[position] != rune('O') {
																goto l923
															}
															position++
														}
													l978:
														{
															position980, tokenIndex980 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l981
															}
															position++
															goto l980
														l981:
															position, tokenIndex = position980, tokenIndex980
															if buffer[position] != rune('R') {
																goto l923
															}
															position++
														}
													l980:
														{
															position982, tokenIndex982 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l983
															}
															position++
															goto l982
														l983:
															position, tokenIndex = position982, tokenIndex982
															if buffer[position] != rune('G') {
																goto l923
															}
															position++
														}
													l982:
														if !_rules[rulews]() {
															goto l923
														}
														if !_rules[rulenn]() {
															goto l923
														}
														{
															add(ruleAction1, position)
														}
														add(ruleOrg, position977)
													}
													break
												case 'a':
													{
														position985 := position
														if buffer[position] != rune('a') {
															goto l923
														}
														position++
														if buffer[position] != rune('s') {
															goto l923
														}
														position++
														if buffer[position] != rune('e') {
															goto l923
														}
														position++
														if buffer[position] != rune('g') {
															goto l923
														}
														position++
														add(ruleAseg, position985)
													}
													break
												default:
													{
														position986 := position
														{
															position987, tokenIndex987 := position, tokenIndex
															if buffer[position] != rune('.') {
																goto l987
															}
															position++
															goto l988
														l987:
															position, tokenIndex = position987, tokenIndex987
														}
													l988:
														if buffer[position] != rune('t') {
															goto l923
														}
														position++
														if buffer[position] != rune('i') {
															goto l923
														}
														position++
														if buffer[position] != rune('t') {
															goto l923
														}
														position++
														if buffer[position] != rune('l') {
															goto l923
														}
														position++
														if buffer[position] != rune('e') {
															goto l923
														}
														position++
														if !_rules[rulews]() {
															goto l923
														}
														if buffer[position] != rune('\'') {
															goto l923
														}
														position++
													l989:
														{
															position990, tokenIndex990 := position, tokenIndex
															{
																position991, tokenIndex991 := position, tokenIndex
																if buffer[position] != rune('\'') {
																	goto l991
																}
																position++
																goto l990
															l991:
																position, tokenIndex = position991, tokenIndex991
															}
															if !matchDot() {
																goto l990
															}
															goto l989
														l990:
															position, tokenIndex = position990, tokenIndex990
														}
														if buffer[position] != rune('\'') {
															goto l923
														}
														position++
														add(ruleTitle, position986)
													}
													break
												}
											}

										}
									l925:
										add(ruleDirective, position924)
									}
									goto l922
								l923:
									position, tokenIndex = position922, tokenIndex922
									{
										position992 := position
										{
											position993, tokenIndex993 := position, tokenIndex
											{
												position995 := position
												{
													position996, tokenIndex996 := position, tokenIndex
													{
														position998 := position
														{
															position999, tokenIndex999 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l1000
															}
															position++
															goto l999
														l1000:
															position, tokenIndex = position999, tokenIndex999
															if buffer[position] != rune('P') {
																goto l997
															}
															position++
														}
													l999:
														{
															position1001, tokenIndex1001 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1002
															}
															position++
															goto l1001
														l1002:
															position, tokenIndex = position1001, tokenIndex1001
															if buffer[position] != rune('U') {
																goto l997
															}
															position++
														}
													l1001:
														{
															position1003, tokenIndex1003 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1004
															}
															position++
															goto l1003
														l1004:
															position, tokenIndex = position1003, tokenIndex1003
															if buffer[position] != rune('S') {
																goto l997
															}
															position++
														}
													l1003:
														{
															position1005, tokenIndex1005 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l1006
															}
															position++
															goto l1005
														l1006:
															position, tokenIndex = position1005, tokenIndex1005
															if buffer[position] != rune('H') {
																goto l997
															}
															position++
														}
													l1005:
														if !_rules[rulews]() {
															goto l997
														}
														if !_rules[ruleSrc16]() {
															goto l997
														}
														{
															add(ruleAction8, position)
														}
														add(rulePush, position998)
													}
													goto l996
												l997:
													position, tokenIndex = position996, tokenIndex996
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1009 := position
																{
																	position1010, tokenIndex1010 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1011
																	}
																	position++
																	goto l1010
																l1011:
																	position, tokenIndex = position1010, tokenIndex1010
																	if buffer[position] != rune('E') {
																		goto l994
																	}
																	position++
																}
															l1010:
																{
																	position1012, tokenIndex1012 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1013
																	}
																	position++
																	goto l1012
																l1013:
																	position, tokenIndex = position1012, tokenIndex1012
																	if buffer[position] != rune('X') {
																		goto l994
																	}
																	position++
																}
															l1012:
																if !_rules[rulews]() {
																	goto l994
																}
																if !_rules[ruleDst16]() {
																	goto l994
																}
																if !_rules[rulesep]() {
																	goto l994
																}
																if !_rules[ruleSrc16]() {
																	goto l994
																}
																{
																	add(ruleAction10, position)
																}
																add(ruleEx, position1009)
															}
															break
														case 'P', 'p':
															{
																position1015 := position
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
																		goto l994
																	}
																	position++
																}
															l1016:
																{
																	position1018, tokenIndex1018 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1019
																	}
																	position++
																	goto l1018
																l1019:
																	position, tokenIndex = position1018, tokenIndex1018
																	if buffer[position] != rune('O') {
																		goto l994
																	}
																	position++
																}
															l1018:
																{
																	position1020, tokenIndex1020 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1021
																	}
																	position++
																	goto l1020
																l1021:
																	position, tokenIndex = position1020, tokenIndex1020
																	if buffer[position] != rune('P') {
																		goto l994
																	}
																	position++
																}
															l1020:
																if !_rules[rulews]() {
																	goto l994
																}
																if !_rules[ruleDst16]() {
																	goto l994
																}
																{
																	add(ruleAction9, position)
																}
																add(rulePop, position1015)
															}
															break
														default:
															{
																position1023 := position
																{
																	position1024, tokenIndex1024 := position, tokenIndex
																	{
																		position1026 := position
																		{
																			position1027, tokenIndex1027 := position, tokenIndex
																			if buffer[position] != rune('l') {
																				goto l1028
																			}
																			position++
																			goto l1027
																		l1028:
																			position, tokenIndex = position1027, tokenIndex1027
																			if buffer[position] != rune('L') {
																				goto l1025
																			}
																			position++
																		}
																	l1027:
																		{
																			position1029, tokenIndex1029 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l1030
																			}
																			position++
																			goto l1029
																		l1030:
																			position, tokenIndex = position1029, tokenIndex1029
																			if buffer[position] != rune('D') {
																				goto l1025
																			}
																			position++
																		}
																	l1029:
																		if !_rules[rulews]() {
																			goto l1025
																		}
																		if !_rules[ruleDst16]() {
																			goto l1025
																		}
																		if !_rules[rulesep]() {
																			goto l1025
																		}
																		if !_rules[ruleSrc16]() {
																			goto l1025
																		}
																		{
																			add(ruleAction7, position)
																		}
																		add(ruleLoad16, position1026)
																	}
																	goto l1024
																l1025:
																	position, tokenIndex = position1024, tokenIndex1024
																	{
																		position1032 := position
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
																				goto l994
																			}
																			position++
																		}
																	l1033:
																		{
																			position1035, tokenIndex1035 := position, tokenIndex
																			if buffer[position] != rune('d') {
																				goto l1036
																			}
																			position++
																			goto l1035
																		l1036:
																			position, tokenIndex = position1035, tokenIndex1035
																			if buffer[position] != rune('D') {
																				goto l994
																			}
																			position++
																		}
																	l1035:
																		if !_rules[rulews]() {
																			goto l994
																		}
																		{
																			position1037 := position
																			{
																				position1038, tokenIndex1038 := position, tokenIndex
																				if !_rules[ruleReg8]() {
																					goto l1039
																				}
																				goto l1038
																			l1039:
																				position, tokenIndex = position1038, tokenIndex1038
																				if !_rules[ruleReg16Contents]() {
																					goto l1040
																				}
																				goto l1038
																			l1040:
																				position, tokenIndex = position1038, tokenIndex1038
																				if !_rules[rulenn_contents]() {
																					goto l994
																				}
																			}
																		l1038:
																			{
																				add(ruleAction20, position)
																			}
																			add(ruleDst8, position1037)
																		}
																		if !_rules[rulesep]() {
																			goto l994
																		}
																		if !_rules[ruleSrc8]() {
																			goto l994
																		}
																		{
																			add(ruleAction6, position)
																		}
																		add(ruleLoad8, position1032)
																	}
																}
															l1024:
																add(ruleLoad, position1023)
															}
															break
														}
													}

												}
											l996:
												add(ruleAssignment, position995)
											}
											goto l993
										l994:
											position, tokenIndex = position993, tokenIndex993
											{
												position1044 := position
												{
													position1045, tokenIndex1045 := position, tokenIndex
													{
														position1047 := position
														{
															position1048, tokenIndex1048 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1049
															}
															position++
															goto l1048
														l1049:
															position, tokenIndex = position1048, tokenIndex1048
															if buffer[position] != rune('I') {
																goto l1046
															}
															position++
														}
													l1048:
														{
															position1050, tokenIndex1050 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1051
															}
															position++
															goto l1050
														l1051:
															position, tokenIndex = position1050, tokenIndex1050
															if buffer[position] != rune('N') {
																goto l1046
															}
															position++
														}
													l1050:
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
																goto l1046
															}
															position++
														}
													l1052:
														if !_rules[rulews]() {
															goto l1046
														}
														if !_rules[ruleILoc8]() {
															goto l1046
														}
														{
															add(ruleAction11, position)
														}
														add(ruleInc16Indexed8, position1047)
													}
													goto l1045
												l1046:
													position, tokenIndex = position1045, tokenIndex1045
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
																goto l1055
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
																goto l1055
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
																goto l1055
															}
															position++
														}
													l1061:
														if !_rules[rulews]() {
															goto l1055
														}
														if !_rules[ruleLoc16]() {
															goto l1055
														}
														{
															add(ruleAction13, position)
														}
														add(ruleInc16, position1056)
													}
													goto l1045
												l1055:
													position, tokenIndex = position1045, tokenIndex1045
													{
														position1064 := position
														{
															position1065, tokenIndex1065 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1066
															}
															position++
															goto l1065
														l1066:
															position, tokenIndex = position1065, tokenIndex1065
															if buffer[position] != rune('I') {
																goto l1043
															}
															position++
														}
													l1065:
														{
															position1067, tokenIndex1067 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1068
															}
															position++
															goto l1067
														l1068:
															position, tokenIndex = position1067, tokenIndex1067
															if buffer[position] != rune('N') {
																goto l1043
															}
															position++
														}
													l1067:
														{
															position1069, tokenIndex1069 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1070
															}
															position++
															goto l1069
														l1070:
															position, tokenIndex = position1069, tokenIndex1069
															if buffer[position] != rune('C') {
																goto l1043
															}
															position++
														}
													l1069:
														if !_rules[rulews]() {
															goto l1043
														}
														if !_rules[ruleLoc8]() {
															goto l1043
														}
														{
															add(ruleAction12, position)
														}
														add(ruleInc8, position1064)
													}
												}
											l1045:
												add(ruleInc, position1044)
											}
											goto l993
										l1043:
											position, tokenIndex = position993, tokenIndex993
											{
												position1073 := position
												{
													position1074, tokenIndex1074 := position, tokenIndex
													{
														position1076 := position
														{
															position1077, tokenIndex1077 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1078
															}
															position++
															goto l1077
														l1078:
															position, tokenIndex = position1077, tokenIndex1077
															if buffer[position] != rune('D') {
																goto l1075
															}
															position++
														}
													l1077:
														{
															position1079, tokenIndex1079 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1080
															}
															position++
															goto l1079
														l1080:
															position, tokenIndex = position1079, tokenIndex1079
															if buffer[position] != rune('E') {
																goto l1075
															}
															position++
														}
													l1079:
														{
															position1081, tokenIndex1081 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1082
															}
															position++
															goto l1081
														l1082:
															position, tokenIndex = position1081, tokenIndex1081
															if buffer[position] != rune('C') {
																goto l1075
															}
															position++
														}
													l1081:
														if !_rules[rulews]() {
															goto l1075
														}
														if !_rules[ruleILoc8]() {
															goto l1075
														}
														{
															add(ruleAction14, position)
														}
														add(ruleDec16Indexed8, position1076)
													}
													goto l1074
												l1075:
													position, tokenIndex = position1074, tokenIndex1074
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
																goto l1084
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
																goto l1084
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
																goto l1084
															}
															position++
														}
													l1090:
														if !_rules[rulews]() {
															goto l1084
														}
														if !_rules[ruleLoc16]() {
															goto l1084
														}
														{
															add(ruleAction16, position)
														}
														add(ruleDec16, position1085)
													}
													goto l1074
												l1084:
													position, tokenIndex = position1074, tokenIndex1074
													{
														position1093 := position
														{
															position1094, tokenIndex1094 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1095
															}
															position++
															goto l1094
														l1095:
															position, tokenIndex = position1094, tokenIndex1094
															if buffer[position] != rune('D') {
																goto l1072
															}
															position++
														}
													l1094:
														{
															position1096, tokenIndex1096 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1097
															}
															position++
															goto l1096
														l1097:
															position, tokenIndex = position1096, tokenIndex1096
															if buffer[position] != rune('E') {
																goto l1072
															}
															position++
														}
													l1096:
														{
															position1098, tokenIndex1098 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1099
															}
															position++
															goto l1098
														l1099:
															position, tokenIndex = position1098, tokenIndex1098
															if buffer[position] != rune('C') {
																goto l1072
															}
															position++
														}
													l1098:
														if !_rules[rulews]() {
															goto l1072
														}
														if !_rules[ruleLoc8]() {
															goto l1072
														}
														{
															add(ruleAction15, position)
														}
														add(ruleDec8, position1093)
													}
												}
											l1074:
												add(ruleDec, position1073)
											}
											goto l993
										l1072:
											position, tokenIndex = position993, tokenIndex993
											{
												position1102 := position
												{
													position1103, tokenIndex1103 := position, tokenIndex
													{
														position1105 := position
														{
															position1106, tokenIndex1106 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1107
															}
															position++
															goto l1106
														l1107:
															position, tokenIndex = position1106, tokenIndex1106
															if buffer[position] != rune('A') {
																goto l1104
															}
															position++
														}
													l1106:
														{
															position1108, tokenIndex1108 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1109
															}
															position++
															goto l1108
														l1109:
															position, tokenIndex = position1108, tokenIndex1108
															if buffer[position] != rune('D') {
																goto l1104
															}
															position++
														}
													l1108:
														{
															position1110, tokenIndex1110 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1111
															}
															position++
															goto l1110
														l1111:
															position, tokenIndex = position1110, tokenIndex1110
															if buffer[position] != rune('D') {
																goto l1104
															}
															position++
														}
													l1110:
														if !_rules[rulews]() {
															goto l1104
														}
														if !_rules[ruleDst16]() {
															goto l1104
														}
														if !_rules[rulesep]() {
															goto l1104
														}
														if !_rules[ruleSrc16]() {
															goto l1104
														}
														{
															add(ruleAction17, position)
														}
														add(ruleAdd16, position1105)
													}
													goto l1103
												l1104:
													position, tokenIndex = position1103, tokenIndex1103
													{
														position1114 := position
														{
															position1115, tokenIndex1115 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1116
															}
															position++
															goto l1115
														l1116:
															position, tokenIndex = position1115, tokenIndex1115
															if buffer[position] != rune('A') {
																goto l1113
															}
															position++
														}
													l1115:
														{
															position1117, tokenIndex1117 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1118
															}
															position++
															goto l1117
														l1118:
															position, tokenIndex = position1117, tokenIndex1117
															if buffer[position] != rune('D') {
																goto l1113
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
																goto l1113
															}
															position++
														}
													l1119:
														if !_rules[rulews]() {
															goto l1113
														}
														if !_rules[ruleDst16]() {
															goto l1113
														}
														if !_rules[rulesep]() {
															goto l1113
														}
														if !_rules[ruleSrc16]() {
															goto l1113
														}
														{
															add(ruleAction18, position)
														}
														add(ruleAdc16, position1114)
													}
													goto l1103
												l1113:
													position, tokenIndex = position1103, tokenIndex1103
													{
														position1122 := position
														{
															position1123, tokenIndex1123 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1124
															}
															position++
															goto l1123
														l1124:
															position, tokenIndex = position1123, tokenIndex1123
															if buffer[position] != rune('S') {
																goto l1101
															}
															position++
														}
													l1123:
														{
															position1125, tokenIndex1125 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1126
															}
															position++
															goto l1125
														l1126:
															position, tokenIndex = position1125, tokenIndex1125
															if buffer[position] != rune('B') {
																goto l1101
															}
															position++
														}
													l1125:
														{
															position1127, tokenIndex1127 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1128
															}
															position++
															goto l1127
														l1128:
															position, tokenIndex = position1127, tokenIndex1127
															if buffer[position] != rune('C') {
																goto l1101
															}
															position++
														}
													l1127:
														if !_rules[rulews]() {
															goto l1101
														}
														if !_rules[ruleDst16]() {
															goto l1101
														}
														if !_rules[rulesep]() {
															goto l1101
														}
														if !_rules[ruleSrc16]() {
															goto l1101
														}
														{
															add(ruleAction19, position)
														}
														add(ruleSbc16, position1122)
													}
												}
											l1103:
												add(ruleAlu16, position1102)
											}
											goto l993
										l1101:
											position, tokenIndex = position993, tokenIndex993
											{
												position1131 := position
												{
													position1132, tokenIndex1132 := position, tokenIndex
													{
														position1134 := position
														{
															position1135, tokenIndex1135 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1136
															}
															position++
															goto l1135
														l1136:
															position, tokenIndex = position1135, tokenIndex1135
															if buffer[position] != rune('A') {
																goto l1133
															}
															position++
														}
													l1135:
														{
															position1137, tokenIndex1137 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1138
															}
															position++
															goto l1137
														l1138:
															position, tokenIndex = position1137, tokenIndex1137
															if buffer[position] != rune('D') {
																goto l1133
															}
															position++
														}
													l1137:
														{
															position1139, tokenIndex1139 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1140
															}
															position++
															goto l1139
														l1140:
															position, tokenIndex = position1139, tokenIndex1139
															if buffer[position] != rune('D') {
																goto l1133
															}
															position++
														}
													l1139:
														if !_rules[rulews]() {
															goto l1133
														}
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
																goto l1133
															}
															position++
														}
													l1141:
														if !_rules[rulesep]() {
															goto l1133
														}
														if !_rules[ruleSrc8]() {
															goto l1133
														}
														{
															add(ruleAction43, position)
														}
														add(ruleAdd, position1134)
													}
													goto l1132
												l1133:
													position, tokenIndex = position1132, tokenIndex1132
													{
														position1145 := position
														{
															position1146, tokenIndex1146 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1147
															}
															position++
															goto l1146
														l1147:
															position, tokenIndex = position1146, tokenIndex1146
															if buffer[position] != rune('A') {
																goto l1144
															}
															position++
														}
													l1146:
														{
															position1148, tokenIndex1148 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1149
															}
															position++
															goto l1148
														l1149:
															position, tokenIndex = position1148, tokenIndex1148
															if buffer[position] != rune('D') {
																goto l1144
															}
															position++
														}
													l1148:
														{
															position1150, tokenIndex1150 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1151
															}
															position++
															goto l1150
														l1151:
															position, tokenIndex = position1150, tokenIndex1150
															if buffer[position] != rune('C') {
																goto l1144
															}
															position++
														}
													l1150:
														if !_rules[rulews]() {
															goto l1144
														}
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
																goto l1144
															}
															position++
														}
													l1152:
														if !_rules[rulesep]() {
															goto l1144
														}
														if !_rules[ruleSrc8]() {
															goto l1144
														}
														{
															add(ruleAction44, position)
														}
														add(ruleAdc, position1145)
													}
													goto l1132
												l1144:
													position, tokenIndex = position1132, tokenIndex1132
													{
														position1156 := position
														{
															position1157, tokenIndex1157 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1158
															}
															position++
															goto l1157
														l1158:
															position, tokenIndex = position1157, tokenIndex1157
															if buffer[position] != rune('S') {
																goto l1155
															}
															position++
														}
													l1157:
														{
															position1159, tokenIndex1159 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1160
															}
															position++
															goto l1159
														l1160:
															position, tokenIndex = position1159, tokenIndex1159
															if buffer[position] != rune('U') {
																goto l1155
															}
															position++
														}
													l1159:
														{
															position1161, tokenIndex1161 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1162
															}
															position++
															goto l1161
														l1162:
															position, tokenIndex = position1161, tokenIndex1161
															if buffer[position] != rune('B') {
																goto l1155
															}
															position++
														}
													l1161:
														if !_rules[rulews]() {
															goto l1155
														}
														if !_rules[ruleSrc8]() {
															goto l1155
														}
														{
															add(ruleAction45, position)
														}
														add(ruleSub, position1156)
													}
													goto l1132
												l1155:
													position, tokenIndex = position1132, tokenIndex1132
													{
														switch buffer[position] {
														case 'C', 'c':
															{
																position1165 := position
																{
																	position1166, tokenIndex1166 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1167
																	}
																	position++
																	goto l1166
																l1167:
																	position, tokenIndex = position1166, tokenIndex1166
																	if buffer[position] != rune('C') {
																		goto l1130
																	}
																	position++
																}
															l1166:
																{
																	position1168, tokenIndex1168 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l1169
																	}
																	position++
																	goto l1168
																l1169:
																	position, tokenIndex = position1168, tokenIndex1168
																	if buffer[position] != rune('P') {
																		goto l1130
																	}
																	position++
																}
															l1168:
																if !_rules[rulews]() {
																	goto l1130
																}
																if !_rules[ruleSrc8]() {
																	goto l1130
																}
																{
																	add(ruleAction50, position)
																}
																add(ruleCp, position1165)
															}
															break
														case 'O', 'o':
															{
																position1171 := position
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
																		goto l1130
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
																		goto l1130
																	}
																	position++
																}
															l1174:
																if !_rules[rulews]() {
																	goto l1130
																}
																if !_rules[ruleSrc8]() {
																	goto l1130
																}
																{
																	add(ruleAction49, position)
																}
																add(ruleOr, position1171)
															}
															break
														case 'X', 'x':
															{
																position1177 := position
																{
																	position1178, tokenIndex1178 := position, tokenIndex
																	if buffer[position] != rune('x') {
																		goto l1179
																	}
																	position++
																	goto l1178
																l1179:
																	position, tokenIndex = position1178, tokenIndex1178
																	if buffer[position] != rune('X') {
																		goto l1130
																	}
																	position++
																}
															l1178:
																{
																	position1180, tokenIndex1180 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l1181
																	}
																	position++
																	goto l1180
																l1181:
																	position, tokenIndex = position1180, tokenIndex1180
																	if buffer[position] != rune('O') {
																		goto l1130
																	}
																	position++
																}
															l1180:
																{
																	position1182, tokenIndex1182 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1183
																	}
																	position++
																	goto l1182
																l1183:
																	position, tokenIndex = position1182, tokenIndex1182
																	if buffer[position] != rune('R') {
																		goto l1130
																	}
																	position++
																}
															l1182:
																if !_rules[rulews]() {
																	goto l1130
																}
																if !_rules[ruleSrc8]() {
																	goto l1130
																}
																{
																	add(ruleAction48, position)
																}
																add(ruleXor, position1177)
															}
															break
														case 'A', 'a':
															{
																position1185 := position
																{
																	position1186, tokenIndex1186 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1187
																	}
																	position++
																	goto l1186
																l1187:
																	position, tokenIndex = position1186, tokenIndex1186
																	if buffer[position] != rune('A') {
																		goto l1130
																	}
																	position++
																}
															l1186:
																{
																	position1188, tokenIndex1188 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1189
																	}
																	position++
																	goto l1188
																l1189:
																	position, tokenIndex = position1188, tokenIndex1188
																	if buffer[position] != rune('N') {
																		goto l1130
																	}
																	position++
																}
															l1188:
																{
																	position1190, tokenIndex1190 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1191
																	}
																	position++
																	goto l1190
																l1191:
																	position, tokenIndex = position1190, tokenIndex1190
																	if buffer[position] != rune('D') {
																		goto l1130
																	}
																	position++
																}
															l1190:
																if !_rules[rulews]() {
																	goto l1130
																}
																{
																	position1192, tokenIndex1192 := position, tokenIndex
																	{
																		position1194, tokenIndex1194 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1195
																		}
																		position++
																		goto l1194
																	l1195:
																		position, tokenIndex = position1194, tokenIndex1194
																		if buffer[position] != rune('A') {
																			goto l1192
																		}
																		position++
																	}
																l1194:
																	if !_rules[rulesep]() {
																		goto l1192
																	}
																	goto l1193
																l1192:
																	position, tokenIndex = position1192, tokenIndex1192
																}
															l1193:
																if !_rules[ruleSrc8]() {
																	goto l1130
																}
																{
																	add(ruleAction47, position)
																}
																add(ruleAnd, position1185)
															}
															break
														default:
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
																		goto l1130
																	}
																	position++
																}
															l1198:
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
																		goto l1130
																	}
																	position++
																}
															l1200:
																{
																	position1202, tokenIndex1202 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1203
																	}
																	position++
																	goto l1202
																l1203:
																	position, tokenIndex = position1202, tokenIndex1202
																	if buffer[position] != rune('C') {
																		goto l1130
																	}
																	position++
																}
															l1202:
																if !_rules[rulews]() {
																	goto l1130
																}
																{
																	position1204, tokenIndex1204 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1205
																	}
																	position++
																	goto l1204
																l1205:
																	position, tokenIndex = position1204, tokenIndex1204
																	if buffer[position] != rune('A') {
																		goto l1130
																	}
																	position++
																}
															l1204:
																if !_rules[rulesep]() {
																	goto l1130
																}
																if !_rules[ruleSrc8]() {
																	goto l1130
																}
																{
																	add(ruleAction46, position)
																}
																add(ruleSbc, position1197)
															}
															break
														}
													}

												}
											l1132:
												add(ruleAlu, position1131)
											}
											goto l993
										l1130:
											position, tokenIndex = position993, tokenIndex993
											{
												position1208 := position
												{
													position1209, tokenIndex1209 := position, tokenIndex
													{
														position1211 := position
														{
															position1212, tokenIndex1212 := position, tokenIndex
															{
																position1214 := position
																{
																	position1215, tokenIndex1215 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1216
																	}
																	position++
																	goto l1215
																l1216:
																	position, tokenIndex = position1215, tokenIndex1215
																	if buffer[position] != rune('R') {
																		goto l1213
																	}
																	position++
																}
															l1215:
																{
																	position1217, tokenIndex1217 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1218
																	}
																	position++
																	goto l1217
																l1218:
																	position, tokenIndex = position1217, tokenIndex1217
																	if buffer[position] != rune('L') {
																		goto l1213
																	}
																	position++
																}
															l1217:
																{
																	position1219, tokenIndex1219 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1220
																	}
																	position++
																	goto l1219
																l1220:
																	position, tokenIndex = position1219, tokenIndex1219
																	if buffer[position] != rune('C') {
																		goto l1213
																	}
																	position++
																}
															l1219:
																if !_rules[rulews]() {
																	goto l1213
																}
																if !_rules[ruleLoc8]() {
																	goto l1213
																}
																{
																	position1221, tokenIndex1221 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1221
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1221
																	}
																	goto l1222
																l1221:
																	position, tokenIndex = position1221, tokenIndex1221
																}
															l1222:
																{
																	add(ruleAction51, position)
																}
																add(ruleRlc, position1214)
															}
															goto l1212
														l1213:
															position, tokenIndex = position1212, tokenIndex1212
															{
																position1225 := position
																{
																	position1226, tokenIndex1226 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1227
																	}
																	position++
																	goto l1226
																l1227:
																	position, tokenIndex = position1226, tokenIndex1226
																	if buffer[position] != rune('R') {
																		goto l1224
																	}
																	position++
																}
															l1226:
																{
																	position1228, tokenIndex1228 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1229
																	}
																	position++
																	goto l1228
																l1229:
																	position, tokenIndex = position1228, tokenIndex1228
																	if buffer[position] != rune('R') {
																		goto l1224
																	}
																	position++
																}
															l1228:
																{
																	position1230, tokenIndex1230 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1231
																	}
																	position++
																	goto l1230
																l1231:
																	position, tokenIndex = position1230, tokenIndex1230
																	if buffer[position] != rune('C') {
																		goto l1224
																	}
																	position++
																}
															l1230:
																if !_rules[rulews]() {
																	goto l1224
																}
																if !_rules[ruleLoc8]() {
																	goto l1224
																}
																{
																	position1232, tokenIndex1232 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1232
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1232
																	}
																	goto l1233
																l1232:
																	position, tokenIndex = position1232, tokenIndex1232
																}
															l1233:
																{
																	add(ruleAction52, position)
																}
																add(ruleRrc, position1225)
															}
															goto l1212
														l1224:
															position, tokenIndex = position1212, tokenIndex1212
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
																		goto l1235
																	}
																	position++
																}
															l1237:
																{
																	position1239, tokenIndex1239 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1240
																	}
																	position++
																	goto l1239
																l1240:
																	position, tokenIndex = position1239, tokenIndex1239
																	if buffer[position] != rune('L') {
																		goto l1235
																	}
																	position++
																}
															l1239:
																if !_rules[rulews]() {
																	goto l1235
																}
																if !_rules[ruleLoc8]() {
																	goto l1235
																}
																{
																	position1241, tokenIndex1241 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1241
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1241
																	}
																	goto l1242
																l1241:
																	position, tokenIndex = position1241, tokenIndex1241
																}
															l1242:
																{
																	add(ruleAction53, position)
																}
																add(ruleRl, position1236)
															}
															goto l1212
														l1235:
															position, tokenIndex = position1212, tokenIndex1212
															{
																position1245 := position
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
																		goto l1244
																	}
																	position++
																}
															l1246:
																{
																	position1248, tokenIndex1248 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1249
																	}
																	position++
																	goto l1248
																l1249:
																	position, tokenIndex = position1248, tokenIndex1248
																	if buffer[position] != rune('R') {
																		goto l1244
																	}
																	position++
																}
															l1248:
																if !_rules[rulews]() {
																	goto l1244
																}
																if !_rules[ruleLoc8]() {
																	goto l1244
																}
																{
																	position1250, tokenIndex1250 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1250
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1250
																	}
																	goto l1251
																l1250:
																	position, tokenIndex = position1250, tokenIndex1250
																}
															l1251:
																{
																	add(ruleAction54, position)
																}
																add(ruleRr, position1245)
															}
															goto l1212
														l1244:
															position, tokenIndex = position1212, tokenIndex1212
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
																		goto l1253
																	}
																	position++
																}
															l1255:
																{
																	position1257, tokenIndex1257 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1258
																	}
																	position++
																	goto l1257
																l1258:
																	position, tokenIndex = position1257, tokenIndex1257
																	if buffer[position] != rune('L') {
																		goto l1253
																	}
																	position++
																}
															l1257:
																{
																	position1259, tokenIndex1259 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1260
																	}
																	position++
																	goto l1259
																l1260:
																	position, tokenIndex = position1259, tokenIndex1259
																	if buffer[position] != rune('A') {
																		goto l1253
																	}
																	position++
																}
															l1259:
																if !_rules[rulews]() {
																	goto l1253
																}
																if !_rules[ruleLoc8]() {
																	goto l1253
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
																	add(ruleAction55, position)
																}
																add(ruleSla, position1254)
															}
															goto l1212
														l1253:
															position, tokenIndex = position1212, tokenIndex1212
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
																		goto l1264
																	}
																	position++
																}
															l1266:
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
																		goto l1264
																	}
																	position++
																}
															l1268:
																{
																	position1270, tokenIndex1270 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1271
																	}
																	position++
																	goto l1270
																l1271:
																	position, tokenIndex = position1270, tokenIndex1270
																	if buffer[position] != rune('A') {
																		goto l1264
																	}
																	position++
																}
															l1270:
																if !_rules[rulews]() {
																	goto l1264
																}
																if !_rules[ruleLoc8]() {
																	goto l1264
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
																	add(ruleAction56, position)
																}
																add(ruleSra, position1265)
															}
															goto l1212
														l1264:
															position, tokenIndex = position1212, tokenIndex1212
															{
																position1276 := position
																{
																	position1277, tokenIndex1277 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1278
																	}
																	position++
																	goto l1277
																l1278:
																	position, tokenIndex = position1277, tokenIndex1277
																	if buffer[position] != rune('S') {
																		goto l1275
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
																		goto l1275
																	}
																	position++
																}
															l1279:
																{
																	position1281, tokenIndex1281 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1282
																	}
																	position++
																	goto l1281
																l1282:
																	position, tokenIndex = position1281, tokenIndex1281
																	if buffer[position] != rune('L') {
																		goto l1275
																	}
																	position++
																}
															l1281:
																if !_rules[rulews]() {
																	goto l1275
																}
																if !_rules[ruleLoc8]() {
																	goto l1275
																}
																{
																	position1283, tokenIndex1283 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1283
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1283
																	}
																	goto l1284
																l1283:
																	position, tokenIndex = position1283, tokenIndex1283
																}
															l1284:
																{
																	add(ruleAction57, position)
																}
																add(ruleSll, position1276)
															}
															goto l1212
														l1275:
															position, tokenIndex = position1212, tokenIndex1212
															{
																position1286 := position
																{
																	position1287, tokenIndex1287 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1288
																	}
																	position++
																	goto l1287
																l1288:
																	position, tokenIndex = position1287, tokenIndex1287
																	if buffer[position] != rune('S') {
																		goto l1210
																	}
																	position++
																}
															l1287:
																{
																	position1289, tokenIndex1289 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1290
																	}
																	position++
																	goto l1289
																l1290:
																	position, tokenIndex = position1289, tokenIndex1289
																	if buffer[position] != rune('R') {
																		goto l1210
																	}
																	position++
																}
															l1289:
																{
																	position1291, tokenIndex1291 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1292
																	}
																	position++
																	goto l1291
																l1292:
																	position, tokenIndex = position1291, tokenIndex1291
																	if buffer[position] != rune('L') {
																		goto l1210
																	}
																	position++
																}
															l1291:
																if !_rules[rulews]() {
																	goto l1210
																}
																if !_rules[ruleLoc8]() {
																	goto l1210
																}
																{
																	position1293, tokenIndex1293 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1293
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1293
																	}
																	goto l1294
																l1293:
																	position, tokenIndex = position1293, tokenIndex1293
																}
															l1294:
																{
																	add(ruleAction58, position)
																}
																add(ruleSrl, position1286)
															}
														}
													l1212:
														add(ruleRot, position1211)
													}
													goto l1209
												l1210:
													position, tokenIndex = position1209, tokenIndex1209
													{
														switch buffer[position] {
														case 'S', 's':
															{
																position1297 := position
																{
																	position1298, tokenIndex1298 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1299
																	}
																	position++
																	goto l1298
																l1299:
																	position, tokenIndex = position1298, tokenIndex1298
																	if buffer[position] != rune('S') {
																		goto l1207
																	}
																	position++
																}
															l1298:
																{
																	position1300, tokenIndex1300 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1301
																	}
																	position++
																	goto l1300
																l1301:
																	position, tokenIndex = position1300, tokenIndex1300
																	if buffer[position] != rune('E') {
																		goto l1207
																	}
																	position++
																}
															l1300:
																{
																	position1302, tokenIndex1302 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1303
																	}
																	position++
																	goto l1302
																l1303:
																	position, tokenIndex = position1302, tokenIndex1302
																	if buffer[position] != rune('T') {
																		goto l1207
																	}
																	position++
																}
															l1302:
																if !_rules[rulews]() {
																	goto l1207
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1207
																}
																if !_rules[rulesep]() {
																	goto l1207
																}
																if !_rules[ruleLoc8]() {
																	goto l1207
																}
																{
																	position1304, tokenIndex1304 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1304
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1304
																	}
																	goto l1305
																l1304:
																	position, tokenIndex = position1304, tokenIndex1304
																}
															l1305:
																{
																	add(ruleAction61, position)
																}
																add(ruleSet, position1297)
															}
															break
														case 'R', 'r':
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
																		goto l1207
																	}
																	position++
																}
															l1308:
																{
																	position1310, tokenIndex1310 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1311
																	}
																	position++
																	goto l1310
																l1311:
																	position, tokenIndex = position1310, tokenIndex1310
																	if buffer[position] != rune('E') {
																		goto l1207
																	}
																	position++
																}
															l1310:
																{
																	position1312, tokenIndex1312 := position, tokenIndex
																	if buffer[position] != rune('s') {
																		goto l1313
																	}
																	position++
																	goto l1312
																l1313:
																	position, tokenIndex = position1312, tokenIndex1312
																	if buffer[position] != rune('S') {
																		goto l1207
																	}
																	position++
																}
															l1312:
																if !_rules[rulews]() {
																	goto l1207
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1207
																}
																if !_rules[rulesep]() {
																	goto l1207
																}
																if !_rules[ruleLoc8]() {
																	goto l1207
																}
																{
																	position1314, tokenIndex1314 := position, tokenIndex
																	if !_rules[rulesep]() {
																		goto l1314
																	}
																	if !_rules[ruleCopy8]() {
																		goto l1314
																	}
																	goto l1315
																l1314:
																	position, tokenIndex = position1314, tokenIndex1314
																}
															l1315:
																{
																	add(ruleAction60, position)
																}
																add(ruleRes, position1307)
															}
															break
														default:
															{
																position1317 := position
																{
																	position1318, tokenIndex1318 := position, tokenIndex
																	if buffer[position] != rune('b') {
																		goto l1319
																	}
																	position++
																	goto l1318
																l1319:
																	position, tokenIndex = position1318, tokenIndex1318
																	if buffer[position] != rune('B') {
																		goto l1207
																	}
																	position++
																}
															l1318:
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
																		goto l1207
																	}
																	position++
																}
															l1320:
																{
																	position1322, tokenIndex1322 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1323
																	}
																	position++
																	goto l1322
																l1323:
																	position, tokenIndex = position1322, tokenIndex1322
																	if buffer[position] != rune('T') {
																		goto l1207
																	}
																	position++
																}
															l1322:
																if !_rules[rulews]() {
																	goto l1207
																}
																if !_rules[ruleoctaldigit]() {
																	goto l1207
																}
																if !_rules[rulesep]() {
																	goto l1207
																}
																if !_rules[ruleLoc8]() {
																	goto l1207
																}
																{
																	add(ruleAction59, position)
																}
																add(ruleBit, position1317)
															}
															break
														}
													}

												}
											l1209:
												add(ruleBitOp, position1208)
											}
											goto l993
										l1207:
											position, tokenIndex = position993, tokenIndex993
											{
												position1326 := position
												{
													position1327, tokenIndex1327 := position, tokenIndex
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
																if buffer[position] != rune('n') {
																	goto l1338
																}
																position++
																goto l1337
															l1338:
																position, tokenIndex = position1337, tokenIndex1337
																if buffer[position] != rune('N') {
																	goto l1328
																}
																position++
															}
														l1337:
															add(rulePegText, position1330)
														}
														{
															add(ruleAction76, position)
														}
														add(ruleRetn, position1329)
													}
													goto l1327
												l1328:
													position, tokenIndex = position1327, tokenIndex1327
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
																if buffer[position] != rune('e') {
																	goto l1346
																}
																position++
																goto l1345
															l1346:
																position, tokenIndex = position1345, tokenIndex1345
																if buffer[position] != rune('E') {
																	goto l1340
																}
																position++
															}
														l1345:
															{
																position1347, tokenIndex1347 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1348
																}
																position++
																goto l1347
															l1348:
																position, tokenIndex = position1347, tokenIndex1347
																if buffer[position] != rune('T') {
																	goto l1340
																}
																position++
															}
														l1347:
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
																	goto l1340
																}
																position++
															}
														l1349:
															add(rulePegText, position1342)
														}
														{
															add(ruleAction77, position)
														}
														add(ruleReti, position1341)
													}
													goto l1327
												l1340:
													position, tokenIndex = position1327, tokenIndex1327
													{
														position1353 := position
														{
															position1354 := position
															{
																position1355, tokenIndex1355 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1356
																}
																position++
																goto l1355
															l1356:
																position, tokenIndex = position1355, tokenIndex1355
																if buffer[position] != rune('R') {
																	goto l1352
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
															add(rulePegText, position1354)
														}
														{
															add(ruleAction78, position)
														}
														add(ruleRrd, position1353)
													}
													goto l1327
												l1352:
													position, tokenIndex = position1327, tokenIndex1327
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
															if buffer[position] != rune('0') {
																goto l1362
															}
															position++
															add(rulePegText, position1364)
														}
														{
															add(ruleAction80, position)
														}
														add(ruleIm0, position1363)
													}
													goto l1327
												l1362:
													position, tokenIndex = position1327, tokenIndex1327
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
															if buffer[position] != rune('1') {
																goto l1370
															}
															position++
															add(rulePegText, position1372)
														}
														{
															add(ruleAction81, position)
														}
														add(ruleIm1, position1371)
													}
													goto l1327
												l1370:
													position, tokenIndex = position1327, tokenIndex1327
													{
														position1379 := position
														{
															position1380 := position
															{
																position1381, tokenIndex1381 := position, tokenIndex
																if buffer[position] != rune('i') {
																	goto l1382
																}
																position++
																goto l1381
															l1382:
																position, tokenIndex = position1381, tokenIndex1381
																if buffer[position] != rune('I') {
																	goto l1378
																}
																position++
															}
														l1381:
															{
																position1383, tokenIndex1383 := position, tokenIndex
																if buffer[position] != rune('m') {
																	goto l1384
																}
																position++
																goto l1383
															l1384:
																position, tokenIndex = position1383, tokenIndex1383
																if buffer[position] != rune('M') {
																	goto l1378
																}
																position++
															}
														l1383:
															if buffer[position] != rune(' ') {
																goto l1378
															}
															position++
															if buffer[position] != rune('2') {
																goto l1378
															}
															position++
															add(rulePegText, position1380)
														}
														{
															add(ruleAction82, position)
														}
														add(ruleIm2, position1379)
													}
													goto l1327
												l1378:
													position, tokenIndex = position1327, tokenIndex1327
													{
														switch buffer[position] {
														case 'I', 'O', 'i', 'o':
															{
																position1387 := position
																{
																	position1388, tokenIndex1388 := position, tokenIndex
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
																			{
																				position1398, tokenIndex1398 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1399
																				}
																				position++
																				goto l1398
																			l1399:
																				position, tokenIndex = position1398, tokenIndex1398
																				if buffer[position] != rune('R') {
																					goto l1389
																				}
																				position++
																			}
																		l1398:
																			add(rulePegText, position1391)
																		}
																		{
																			add(ruleAction93, position)
																		}
																		add(ruleInir, position1390)
																	}
																	goto l1388
																l1389:
																	position, tokenIndex = position1388, tokenIndex1388
																	{
																		position1402 := position
																		{
																			position1403 := position
																			{
																				position1404, tokenIndex1404 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1405
																				}
																				position++
																				goto l1404
																			l1405:
																				position, tokenIndex = position1404, tokenIndex1404
																				if buffer[position] != rune('I') {
																					goto l1401
																				}
																				position++
																			}
																		l1404:
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
																					goto l1401
																				}
																				position++
																			}
																		l1406:
																			{
																				position1408, tokenIndex1408 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1409
																				}
																				position++
																				goto l1408
																			l1409:
																				position, tokenIndex = position1408, tokenIndex1408
																				if buffer[position] != rune('I') {
																					goto l1401
																				}
																				position++
																			}
																		l1408:
																			add(rulePegText, position1403)
																		}
																		{
																			add(ruleAction85, position)
																		}
																		add(ruleIni, position1402)
																	}
																	goto l1388
																l1401:
																	position, tokenIndex = position1388, tokenIndex1388
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
																				if buffer[position] != rune('t') {
																					goto l1417
																				}
																				position++
																				goto l1416
																			l1417:
																				position, tokenIndex = position1416, tokenIndex1416
																				if buffer[position] != rune('T') {
																					goto l1411
																				}
																				position++
																			}
																		l1416:
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
																					goto l1411
																				}
																				position++
																			}
																		l1418:
																			{
																				position1420, tokenIndex1420 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1421
																				}
																				position++
																				goto l1420
																			l1421:
																				position, tokenIndex = position1420, tokenIndex1420
																				if buffer[position] != rune('R') {
																					goto l1411
																				}
																				position++
																			}
																		l1420:
																			add(rulePegText, position1413)
																		}
																		{
																			add(ruleAction94, position)
																		}
																		add(ruleOtir, position1412)
																	}
																	goto l1388
																l1411:
																	position, tokenIndex = position1388, tokenIndex1388
																	{
																		position1424 := position
																		{
																			position1425 := position
																			{
																				position1426, tokenIndex1426 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1427
																				}
																				position++
																				goto l1426
																			l1427:
																				position, tokenIndex = position1426, tokenIndex1426
																				if buffer[position] != rune('O') {
																					goto l1423
																				}
																				position++
																			}
																		l1426:
																			{
																				position1428, tokenIndex1428 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1429
																				}
																				position++
																				goto l1428
																			l1429:
																				position, tokenIndex = position1428, tokenIndex1428
																				if buffer[position] != rune('U') {
																					goto l1423
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
																					goto l1423
																				}
																				position++
																			}
																		l1430:
																			{
																				position1432, tokenIndex1432 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1433
																				}
																				position++
																				goto l1432
																			l1433:
																				position, tokenIndex = position1432, tokenIndex1432
																				if buffer[position] != rune('I') {
																					goto l1423
																				}
																				position++
																			}
																		l1432:
																			add(rulePegText, position1425)
																		}
																		{
																			add(ruleAction86, position)
																		}
																		add(ruleOuti, position1424)
																	}
																	goto l1388
																l1423:
																	position, tokenIndex = position1388, tokenIndex1388
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
																			{
																				position1444, tokenIndex1444 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1445
																				}
																				position++
																				goto l1444
																			l1445:
																				position, tokenIndex = position1444, tokenIndex1444
																				if buffer[position] != rune('R') {
																					goto l1435
																				}
																				position++
																			}
																		l1444:
																			add(rulePegText, position1437)
																		}
																		{
																			add(ruleAction97, position)
																		}
																		add(ruleIndr, position1436)
																	}
																	goto l1388
																l1435:
																	position, tokenIndex = position1388, tokenIndex1388
																	{
																		position1448 := position
																		{
																			position1449 := position
																			{
																				position1450, tokenIndex1450 := position, tokenIndex
																				if buffer[position] != rune('i') {
																					goto l1451
																				}
																				position++
																				goto l1450
																			l1451:
																				position, tokenIndex = position1450, tokenIndex1450
																				if buffer[position] != rune('I') {
																					goto l1447
																				}
																				position++
																			}
																		l1450:
																			{
																				position1452, tokenIndex1452 := position, tokenIndex
																				if buffer[position] != rune('n') {
																					goto l1453
																				}
																				position++
																				goto l1452
																			l1453:
																				position, tokenIndex = position1452, tokenIndex1452
																				if buffer[position] != rune('N') {
																					goto l1447
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
																					goto l1447
																				}
																				position++
																			}
																		l1454:
																			add(rulePegText, position1449)
																		}
																		{
																			add(ruleAction89, position)
																		}
																		add(ruleInd, position1448)
																	}
																	goto l1388
																l1447:
																	position, tokenIndex = position1388, tokenIndex1388
																	{
																		position1458 := position
																		{
																			position1459 := position
																			{
																				position1460, tokenIndex1460 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1461
																				}
																				position++
																				goto l1460
																			l1461:
																				position, tokenIndex = position1460, tokenIndex1460
																				if buffer[position] != rune('O') {
																					goto l1457
																				}
																				position++
																			}
																		l1460:
																			{
																				position1462, tokenIndex1462 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1463
																				}
																				position++
																				goto l1462
																			l1463:
																				position, tokenIndex = position1462, tokenIndex1462
																				if buffer[position] != rune('T') {
																					goto l1457
																				}
																				position++
																			}
																		l1462:
																			{
																				position1464, tokenIndex1464 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1465
																				}
																				position++
																				goto l1464
																			l1465:
																				position, tokenIndex = position1464, tokenIndex1464
																				if buffer[position] != rune('D') {
																					goto l1457
																				}
																				position++
																			}
																		l1464:
																			{
																				position1466, tokenIndex1466 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1467
																				}
																				position++
																				goto l1466
																			l1467:
																				position, tokenIndex = position1466, tokenIndex1466
																				if buffer[position] != rune('R') {
																					goto l1457
																				}
																				position++
																			}
																		l1466:
																			add(rulePegText, position1459)
																		}
																		{
																			add(ruleAction98, position)
																		}
																		add(ruleOtdr, position1458)
																	}
																	goto l1388
																l1457:
																	position, tokenIndex = position1388, tokenIndex1388
																	{
																		position1469 := position
																		{
																			position1470 := position
																			{
																				position1471, tokenIndex1471 := position, tokenIndex
																				if buffer[position] != rune('o') {
																					goto l1472
																				}
																				position++
																				goto l1471
																			l1472:
																				position, tokenIndex = position1471, tokenIndex1471
																				if buffer[position] != rune('O') {
																					goto l1325
																				}
																				position++
																			}
																		l1471:
																			{
																				position1473, tokenIndex1473 := position, tokenIndex
																				if buffer[position] != rune('u') {
																					goto l1474
																				}
																				position++
																				goto l1473
																			l1474:
																				position, tokenIndex = position1473, tokenIndex1473
																				if buffer[position] != rune('U') {
																					goto l1325
																				}
																				position++
																			}
																		l1473:
																			{
																				position1475, tokenIndex1475 := position, tokenIndex
																				if buffer[position] != rune('t') {
																					goto l1476
																				}
																				position++
																				goto l1475
																			l1476:
																				position, tokenIndex = position1475, tokenIndex1475
																				if buffer[position] != rune('T') {
																					goto l1325
																				}
																				position++
																			}
																		l1475:
																			{
																				position1477, tokenIndex1477 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1478
																				}
																				position++
																				goto l1477
																			l1478:
																				position, tokenIndex = position1477, tokenIndex1477
																				if buffer[position] != rune('D') {
																					goto l1325
																				}
																				position++
																			}
																		l1477:
																			add(rulePegText, position1470)
																		}
																		{
																			add(ruleAction90, position)
																		}
																		add(ruleOutd, position1469)
																	}
																}
															l1388:
																add(ruleBlitIO, position1387)
															}
															break
														case 'R', 'r':
															{
																position1480 := position
																{
																	position1481 := position
																	{
																		position1482, tokenIndex1482 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1483
																		}
																		position++
																		goto l1482
																	l1483:
																		position, tokenIndex = position1482, tokenIndex1482
																		if buffer[position] != rune('R') {
																			goto l1325
																		}
																		position++
																	}
																l1482:
																	{
																		position1484, tokenIndex1484 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1485
																		}
																		position++
																		goto l1484
																	l1485:
																		position, tokenIndex = position1484, tokenIndex1484
																		if buffer[position] != rune('L') {
																			goto l1325
																		}
																		position++
																	}
																l1484:
																	{
																		position1486, tokenIndex1486 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1487
																		}
																		position++
																		goto l1486
																	l1487:
																		position, tokenIndex = position1486, tokenIndex1486
																		if buffer[position] != rune('D') {
																			goto l1325
																		}
																		position++
																	}
																l1486:
																	add(rulePegText, position1481)
																}
																{
																	add(ruleAction79, position)
																}
																add(ruleRld, position1480)
															}
															break
														case 'N', 'n':
															{
																position1489 := position
																{
																	position1490 := position
																	{
																		position1491, tokenIndex1491 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1492
																		}
																		position++
																		goto l1491
																	l1492:
																		position, tokenIndex = position1491, tokenIndex1491
																		if buffer[position] != rune('N') {
																			goto l1325
																		}
																		position++
																	}
																l1491:
																	{
																		position1493, tokenIndex1493 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1494
																		}
																		position++
																		goto l1493
																	l1494:
																		position, tokenIndex = position1493, tokenIndex1493
																		if buffer[position] != rune('E') {
																			goto l1325
																		}
																		position++
																	}
																l1493:
																	{
																		position1495, tokenIndex1495 := position, tokenIndex
																		if buffer[position] != rune('g') {
																			goto l1496
																		}
																		position++
																		goto l1495
																	l1496:
																		position, tokenIndex = position1495, tokenIndex1495
																		if buffer[position] != rune('G') {
																			goto l1325
																		}
																		position++
																	}
																l1495:
																	add(rulePegText, position1490)
																}
																{
																	add(ruleAction75, position)
																}
																add(ruleNeg, position1489)
															}
															break
														default:
															{
																position1498 := position
																{
																	position1499, tokenIndex1499 := position, tokenIndex
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
																			{
																				position1509, tokenIndex1509 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1510
																				}
																				position++
																				goto l1509
																			l1510:
																				position, tokenIndex = position1509, tokenIndex1509
																				if buffer[position] != rune('R') {
																					goto l1500
																				}
																				position++
																			}
																		l1509:
																			add(rulePegText, position1502)
																		}
																		{
																			add(ruleAction91, position)
																		}
																		add(ruleLdir, position1501)
																	}
																	goto l1499
																l1500:
																	position, tokenIndex = position1499, tokenIndex1499
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
																				if buffer[position] != rune('i') {
																					goto l1520
																				}
																				position++
																				goto l1519
																			l1520:
																				position, tokenIndex = position1519, tokenIndex1519
																				if buffer[position] != rune('I') {
																					goto l1512
																				}
																				position++
																			}
																		l1519:
																			add(rulePegText, position1514)
																		}
																		{
																			add(ruleAction83, position)
																		}
																		add(ruleLdi, position1513)
																	}
																	goto l1499
																l1512:
																	position, tokenIndex = position1499, tokenIndex1499
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
																			{
																				position1531, tokenIndex1531 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1532
																				}
																				position++
																				goto l1531
																			l1532:
																				position, tokenIndex = position1531, tokenIndex1531
																				if buffer[position] != rune('R') {
																					goto l1522
																				}
																				position++
																			}
																		l1531:
																			add(rulePegText, position1524)
																		}
																		{
																			add(ruleAction92, position)
																		}
																		add(ruleCpir, position1523)
																	}
																	goto l1499
																l1522:
																	position, tokenIndex = position1499, tokenIndex1499
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
																				if buffer[position] != rune('i') {
																					goto l1542
																				}
																				position++
																				goto l1541
																			l1542:
																				position, tokenIndex = position1541, tokenIndex1541
																				if buffer[position] != rune('I') {
																					goto l1534
																				}
																				position++
																			}
																		l1541:
																			add(rulePegText, position1536)
																		}
																		{
																			add(ruleAction84, position)
																		}
																		add(ruleCpi, position1535)
																	}
																	goto l1499
																l1534:
																	position, tokenIndex = position1499, tokenIndex1499
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
																			{
																				position1553, tokenIndex1553 := position, tokenIndex
																				if buffer[position] != rune('r') {
																					goto l1554
																				}
																				position++
																				goto l1553
																			l1554:
																				position, tokenIndex = position1553, tokenIndex1553
																				if buffer[position] != rune('R') {
																					goto l1544
																				}
																				position++
																			}
																		l1553:
																			add(rulePegText, position1546)
																		}
																		{
																			add(ruleAction95, position)
																		}
																		add(ruleLddr, position1545)
																	}
																	goto l1499
																l1544:
																	position, tokenIndex = position1499, tokenIndex1499
																	{
																		position1557 := position
																		{
																			position1558 := position
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
																					goto l1556
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
																					goto l1556
																				}
																				position++
																			}
																		l1561:
																			{
																				position1563, tokenIndex1563 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1564
																				}
																				position++
																				goto l1563
																			l1564:
																				position, tokenIndex = position1563, tokenIndex1563
																				if buffer[position] != rune('D') {
																					goto l1556
																				}
																				position++
																			}
																		l1563:
																			add(rulePegText, position1558)
																		}
																		{
																			add(ruleAction87, position)
																		}
																		add(ruleLdd, position1557)
																	}
																	goto l1499
																l1556:
																	position, tokenIndex = position1499, tokenIndex1499
																	{
																		position1567 := position
																		{
																			position1568 := position
																			{
																				position1569, tokenIndex1569 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1570
																				}
																				position++
																				goto l1569
																			l1570:
																				position, tokenIndex = position1569, tokenIndex1569
																				if buffer[position] != rune('C') {
																					goto l1566
																				}
																				position++
																			}
																		l1569:
																			{
																				position1571, tokenIndex1571 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1572
																				}
																				position++
																				goto l1571
																			l1572:
																				position, tokenIndex = position1571, tokenIndex1571
																				if buffer[position] != rune('P') {
																					goto l1566
																				}
																				position++
																			}
																		l1571:
																			{
																				position1573, tokenIndex1573 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1574
																				}
																				position++
																				goto l1573
																			l1574:
																				position, tokenIndex = position1573, tokenIndex1573
																				if buffer[position] != rune('D') {
																					goto l1566
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
																					goto l1566
																				}
																				position++
																			}
																		l1575:
																			add(rulePegText, position1568)
																		}
																		{
																			add(ruleAction96, position)
																		}
																		add(ruleCpdr, position1567)
																	}
																	goto l1499
																l1566:
																	position, tokenIndex = position1499, tokenIndex1499
																	{
																		position1578 := position
																		{
																			position1579 := position
																			{
																				position1580, tokenIndex1580 := position, tokenIndex
																				if buffer[position] != rune('c') {
																					goto l1581
																				}
																				position++
																				goto l1580
																			l1581:
																				position, tokenIndex = position1580, tokenIndex1580
																				if buffer[position] != rune('C') {
																					goto l1325
																				}
																				position++
																			}
																		l1580:
																			{
																				position1582, tokenIndex1582 := position, tokenIndex
																				if buffer[position] != rune('p') {
																					goto l1583
																				}
																				position++
																				goto l1582
																			l1583:
																				position, tokenIndex = position1582, tokenIndex1582
																				if buffer[position] != rune('P') {
																					goto l1325
																				}
																				position++
																			}
																		l1582:
																			{
																				position1584, tokenIndex1584 := position, tokenIndex
																				if buffer[position] != rune('d') {
																					goto l1585
																				}
																				position++
																				goto l1584
																			l1585:
																				position, tokenIndex = position1584, tokenIndex1584
																				if buffer[position] != rune('D') {
																					goto l1325
																				}
																				position++
																			}
																		l1584:
																			add(rulePegText, position1579)
																		}
																		{
																			add(ruleAction88, position)
																		}
																		add(ruleCpd, position1578)
																	}
																}
															l1499:
																add(ruleBlit, position1498)
															}
															break
														}
													}

												}
											l1327:
												add(ruleEDSimple, position1326)
											}
											goto l993
										l1325:
											position, tokenIndex = position993, tokenIndex993
											{
												position1588 := position
												{
													position1589, tokenIndex1589 := position, tokenIndex
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
																if buffer[position] != rune('l') {
																	goto l1596
																}
																position++
																goto l1595
															l1596:
																position, tokenIndex = position1595, tokenIndex1595
																if buffer[position] != rune('L') {
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
															add(ruleAction64, position)
														}
														add(ruleRlca, position1591)
													}
													goto l1589
												l1590:
													position, tokenIndex = position1589, tokenIndex1589
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
																if buffer[position] != rune('r') {
																	goto l1608
																}
																position++
																goto l1607
															l1608:
																position, tokenIndex = position1607, tokenIndex1607
																if buffer[position] != rune('R') {
																	goto l1602
																}
																position++
															}
														l1607:
															{
																position1609, tokenIndex1609 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1610
																}
																position++
																goto l1609
															l1610:
																position, tokenIndex = position1609, tokenIndex1609
																if buffer[position] != rune('C') {
																	goto l1602
																}
																position++
															}
														l1609:
															{
																position1611, tokenIndex1611 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1612
																}
																position++
																goto l1611
															l1612:
																position, tokenIndex = position1611, tokenIndex1611
																if buffer[position] != rune('A') {
																	goto l1602
																}
																position++
															}
														l1611:
															add(rulePegText, position1604)
														}
														{
															add(ruleAction65, position)
														}
														add(ruleRrca, position1603)
													}
													goto l1589
												l1602:
													position, tokenIndex = position1589, tokenIndex1589
													{
														position1615 := position
														{
															position1616 := position
															{
																position1617, tokenIndex1617 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1618
																}
																position++
																goto l1617
															l1618:
																position, tokenIndex = position1617, tokenIndex1617
																if buffer[position] != rune('R') {
																	goto l1614
																}
																position++
															}
														l1617:
															{
																position1619, tokenIndex1619 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1620
																}
																position++
																goto l1619
															l1620:
																position, tokenIndex = position1619, tokenIndex1619
																if buffer[position] != rune('L') {
																	goto l1614
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
																	goto l1614
																}
																position++
															}
														l1621:
															add(rulePegText, position1616)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleRla, position1615)
													}
													goto l1589
												l1614:
													position, tokenIndex = position1589, tokenIndex1589
													{
														position1625 := position
														{
															position1626 := position
															{
																position1627, tokenIndex1627 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1628
																}
																position++
																goto l1627
															l1628:
																position, tokenIndex = position1627, tokenIndex1627
																if buffer[position] != rune('D') {
																	goto l1624
																}
																position++
															}
														l1627:
															{
																position1629, tokenIndex1629 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1630
																}
																position++
																goto l1629
															l1630:
																position, tokenIndex = position1629, tokenIndex1629
																if buffer[position] != rune('A') {
																	goto l1624
																}
																position++
															}
														l1629:
															{
																position1631, tokenIndex1631 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1632
																}
																position++
																goto l1631
															l1632:
																position, tokenIndex = position1631, tokenIndex1631
																if buffer[position] != rune('A') {
																	goto l1624
																}
																position++
															}
														l1631:
															add(rulePegText, position1626)
														}
														{
															add(ruleAction68, position)
														}
														add(ruleDaa, position1625)
													}
													goto l1589
												l1624:
													position, tokenIndex = position1589, tokenIndex1589
													{
														position1635 := position
														{
															position1636 := position
															{
																position1637, tokenIndex1637 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1638
																}
																position++
																goto l1637
															l1638:
																position, tokenIndex = position1637, tokenIndex1637
																if buffer[position] != rune('C') {
																	goto l1634
																}
																position++
															}
														l1637:
															{
																position1639, tokenIndex1639 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1640
																}
																position++
																goto l1639
															l1640:
																position, tokenIndex = position1639, tokenIndex1639
																if buffer[position] != rune('P') {
																	goto l1634
																}
																position++
															}
														l1639:
															{
																position1641, tokenIndex1641 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1642
																}
																position++
																goto l1641
															l1642:
																position, tokenIndex = position1641, tokenIndex1641
																if buffer[position] != rune('L') {
																	goto l1634
																}
																position++
															}
														l1641:
															add(rulePegText, position1636)
														}
														{
															add(ruleAction69, position)
														}
														add(ruleCpl, position1635)
													}
													goto l1589
												l1634:
													position, tokenIndex = position1589, tokenIndex1589
													{
														position1645 := position
														{
															position1646 := position
															{
																position1647, tokenIndex1647 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1648
																}
																position++
																goto l1647
															l1648:
																position, tokenIndex = position1647, tokenIndex1647
																if buffer[position] != rune('E') {
																	goto l1644
																}
																position++
															}
														l1647:
															{
																position1649, tokenIndex1649 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1650
																}
																position++
																goto l1649
															l1650:
																position, tokenIndex = position1649, tokenIndex1649
																if buffer[position] != rune('X') {
																	goto l1644
																}
																position++
															}
														l1649:
															{
																position1651, tokenIndex1651 := position, tokenIndex
																if buffer[position] != rune('x') {
																	goto l1652
																}
																position++
																goto l1651
															l1652:
																position, tokenIndex = position1651, tokenIndex1651
																if buffer[position] != rune('X') {
																	goto l1644
																}
																position++
															}
														l1651:
															add(rulePegText, position1646)
														}
														{
															add(ruleAction72, position)
														}
														add(ruleExx, position1645)
													}
													goto l1589
												l1644:
													position, tokenIndex = position1589, tokenIndex1589
													{
														switch buffer[position] {
														case 'E', 'e':
															{
																position1655 := position
																{
																	position1656 := position
																	{
																		position1657, tokenIndex1657 := position, tokenIndex
																		if buffer[position] != rune('e') {
																			goto l1658
																		}
																		position++
																		goto l1657
																	l1658:
																		position, tokenIndex = position1657, tokenIndex1657
																		if buffer[position] != rune('E') {
																			goto l1587
																		}
																		position++
																	}
																l1657:
																	{
																		position1659, tokenIndex1659 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1660
																		}
																		position++
																		goto l1659
																	l1660:
																		position, tokenIndex = position1659, tokenIndex1659
																		if buffer[position] != rune('I') {
																			goto l1587
																		}
																		position++
																	}
																l1659:
																	add(rulePegText, position1656)
																}
																{
																	add(ruleAction74, position)
																}
																add(ruleEi, position1655)
															}
															break
														case 'D', 'd':
															{
																position1662 := position
																{
																	position1663 := position
																	{
																		position1664, tokenIndex1664 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1665
																		}
																		position++
																		goto l1664
																	l1665:
																		position, tokenIndex = position1664, tokenIndex1664
																		if buffer[position] != rune('D') {
																			goto l1587
																		}
																		position++
																	}
																l1664:
																	{
																		position1666, tokenIndex1666 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1667
																		}
																		position++
																		goto l1666
																	l1667:
																		position, tokenIndex = position1666, tokenIndex1666
																		if buffer[position] != rune('I') {
																			goto l1587
																		}
																		position++
																	}
																l1666:
																	add(rulePegText, position1663)
																}
																{
																	add(ruleAction73, position)
																}
																add(ruleDi, position1662)
															}
															break
														case 'C', 'c':
															{
																position1669 := position
																{
																	position1670 := position
																	{
																		position1671, tokenIndex1671 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1672
																		}
																		position++
																		goto l1671
																	l1672:
																		position, tokenIndex = position1671, tokenIndex1671
																		if buffer[position] != rune('C') {
																			goto l1587
																		}
																		position++
																	}
																l1671:
																	{
																		position1673, tokenIndex1673 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1674
																		}
																		position++
																		goto l1673
																	l1674:
																		position, tokenIndex = position1673, tokenIndex1673
																		if buffer[position] != rune('C') {
																			goto l1587
																		}
																		position++
																	}
																l1673:
																	{
																		position1675, tokenIndex1675 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1676
																		}
																		position++
																		goto l1675
																	l1676:
																		position, tokenIndex = position1675, tokenIndex1675
																		if buffer[position] != rune('F') {
																			goto l1587
																		}
																		position++
																	}
																l1675:
																	add(rulePegText, position1670)
																}
																{
																	add(ruleAction71, position)
																}
																add(ruleCcf, position1669)
															}
															break
														case 'S', 's':
															{
																position1678 := position
																{
																	position1679 := position
																	{
																		position1680, tokenIndex1680 := position, tokenIndex
																		if buffer[position] != rune('s') {
																			goto l1681
																		}
																		position++
																		goto l1680
																	l1681:
																		position, tokenIndex = position1680, tokenIndex1680
																		if buffer[position] != rune('S') {
																			goto l1587
																		}
																		position++
																	}
																l1680:
																	{
																		position1682, tokenIndex1682 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1683
																		}
																		position++
																		goto l1682
																	l1683:
																		position, tokenIndex = position1682, tokenIndex1682
																		if buffer[position] != rune('C') {
																			goto l1587
																		}
																		position++
																	}
																l1682:
																	{
																		position1684, tokenIndex1684 := position, tokenIndex
																		if buffer[position] != rune('f') {
																			goto l1685
																		}
																		position++
																		goto l1684
																	l1685:
																		position, tokenIndex = position1684, tokenIndex1684
																		if buffer[position] != rune('F') {
																			goto l1587
																		}
																		position++
																	}
																l1684:
																	add(rulePegText, position1679)
																}
																{
																	add(ruleAction70, position)
																}
																add(ruleScf, position1678)
															}
															break
														case 'R', 'r':
															{
																position1687 := position
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
																			goto l1587
																		}
																		position++
																	}
																l1689:
																	{
																		position1691, tokenIndex1691 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1692
																		}
																		position++
																		goto l1691
																	l1692:
																		position, tokenIndex = position1691, tokenIndex1691
																		if buffer[position] != rune('R') {
																			goto l1587
																		}
																		position++
																	}
																l1691:
																	{
																		position1693, tokenIndex1693 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1694
																		}
																		position++
																		goto l1693
																	l1694:
																		position, tokenIndex = position1693, tokenIndex1693
																		if buffer[position] != rune('A') {
																			goto l1587
																		}
																		position++
																	}
																l1693:
																	add(rulePegText, position1688)
																}
																{
																	add(ruleAction67, position)
																}
																add(ruleRra, position1687)
															}
															break
														case 'H', 'h':
															{
																position1696 := position
																{
																	position1697 := position
																	{
																		position1698, tokenIndex1698 := position, tokenIndex
																		if buffer[position] != rune('h') {
																			goto l1699
																		}
																		position++
																		goto l1698
																	l1699:
																		position, tokenIndex = position1698, tokenIndex1698
																		if buffer[position] != rune('H') {
																			goto l1587
																		}
																		position++
																	}
																l1698:
																	{
																		position1700, tokenIndex1700 := position, tokenIndex
																		if buffer[position] != rune('a') {
																			goto l1701
																		}
																		position++
																		goto l1700
																	l1701:
																		position, tokenIndex = position1700, tokenIndex1700
																		if buffer[position] != rune('A') {
																			goto l1587
																		}
																		position++
																	}
																l1700:
																	{
																		position1702, tokenIndex1702 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1703
																		}
																		position++
																		goto l1702
																	l1703:
																		position, tokenIndex = position1702, tokenIndex1702
																		if buffer[position] != rune('L') {
																			goto l1587
																		}
																		position++
																	}
																l1702:
																	{
																		position1704, tokenIndex1704 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1705
																		}
																		position++
																		goto l1704
																	l1705:
																		position, tokenIndex = position1704, tokenIndex1704
																		if buffer[position] != rune('T') {
																			goto l1587
																		}
																		position++
																	}
																l1704:
																	add(rulePegText, position1697)
																}
																{
																	add(ruleAction63, position)
																}
																add(ruleHalt, position1696)
															}
															break
														default:
															{
																position1707 := position
																{
																	position1708 := position
																	{
																		position1709, tokenIndex1709 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1710
																		}
																		position++
																		goto l1709
																	l1710:
																		position, tokenIndex = position1709, tokenIndex1709
																		if buffer[position] != rune('N') {
																			goto l1587
																		}
																		position++
																	}
																l1709:
																	{
																		position1711, tokenIndex1711 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1712
																		}
																		position++
																		goto l1711
																	l1712:
																		position, tokenIndex = position1711, tokenIndex1711
																		if buffer[position] != rune('O') {
																			goto l1587
																		}
																		position++
																	}
																l1711:
																	{
																		position1713, tokenIndex1713 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1714
																		}
																		position++
																		goto l1713
																	l1714:
																		position, tokenIndex = position1713, tokenIndex1713
																		if buffer[position] != rune('P') {
																			goto l1587
																		}
																		position++
																	}
																l1713:
																	add(rulePegText, position1708)
																}
																{
																	add(ruleAction62, position)
																}
																add(ruleNop, position1707)
															}
															break
														}
													}

												}
											l1589:
												add(ruleSimple, position1588)
											}
											goto l993
										l1587:
											position, tokenIndex = position993, tokenIndex993
											{
												position1717 := position
												{
													position1718, tokenIndex1718 := position, tokenIndex
													{
														position1720 := position
														{
															position1721, tokenIndex1721 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1722
															}
															position++
															goto l1721
														l1722:
															position, tokenIndex = position1721, tokenIndex1721
															if buffer[position] != rune('R') {
																goto l1719
															}
															position++
														}
													l1721:
														{
															position1723, tokenIndex1723 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1724
															}
															position++
															goto l1723
														l1724:
															position, tokenIndex = position1723, tokenIndex1723
															if buffer[position] != rune('S') {
																goto l1719
															}
															position++
														}
													l1723:
														{
															position1725, tokenIndex1725 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1726
															}
															position++
															goto l1725
														l1726:
															position, tokenIndex = position1725, tokenIndex1725
															if buffer[position] != rune('T') {
																goto l1719
															}
															position++
														}
													l1725:
														if !_rules[rulews]() {
															goto l1719
														}
														if !_rules[rulen]() {
															goto l1719
														}
														{
															add(ruleAction99, position)
														}
														add(ruleRst, position1720)
													}
													goto l1718
												l1719:
													position, tokenIndex = position1718, tokenIndex1718
													{
														position1729 := position
														{
															position1730, tokenIndex1730 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1731
															}
															position++
															goto l1730
														l1731:
															position, tokenIndex = position1730, tokenIndex1730
															if buffer[position] != rune('J') {
																goto l1728
															}
															position++
														}
													l1730:
														{
															position1732, tokenIndex1732 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l1733
															}
															position++
															goto l1732
														l1733:
															position, tokenIndex = position1732, tokenIndex1732
															if buffer[position] != rune('P') {
																goto l1728
															}
															position++
														}
													l1732:
														if !_rules[rulews]() {
															goto l1728
														}
														{
															position1734, tokenIndex1734 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1734
															}
															if !_rules[rulesep]() {
																goto l1734
															}
															goto l1735
														l1734:
															position, tokenIndex = position1734, tokenIndex1734
														}
													l1735:
														if !_rules[ruleSrc16]() {
															goto l1728
														}
														{
															add(ruleAction102, position)
														}
														add(ruleJp, position1729)
													}
													goto l1718
												l1728:
													position, tokenIndex = position1718, tokenIndex1718
													{
														switch buffer[position] {
														case 'D', 'd':
															{
																position1738 := position
																{
																	position1739, tokenIndex1739 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l1740
																	}
																	position++
																	goto l1739
																l1740:
																	position, tokenIndex = position1739, tokenIndex1739
																	if buffer[position] != rune('D') {
																		goto l1716
																	}
																	position++
																}
															l1739:
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
																		goto l1716
																	}
																	position++
																}
															l1741:
																{
																	position1743, tokenIndex1743 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l1744
																	}
																	position++
																	goto l1743
																l1744:
																	position, tokenIndex = position1743, tokenIndex1743
																	if buffer[position] != rune('N') {
																		goto l1716
																	}
																	position++
																}
															l1743:
																{
																	position1745, tokenIndex1745 := position, tokenIndex
																	if buffer[position] != rune('z') {
																		goto l1746
																	}
																	position++
																	goto l1745
																l1746:
																	position, tokenIndex = position1745, tokenIndex1745
																	if buffer[position] != rune('Z') {
																		goto l1716
																	}
																	position++
																}
															l1745:
																if !_rules[rulews]() {
																	goto l1716
																}
																if !_rules[ruledisp]() {
																	goto l1716
																}
																{
																	add(ruleAction104, position)
																}
																add(ruleDjnz, position1738)
															}
															break
														case 'J', 'j':
															{
																position1748 := position
																{
																	position1749, tokenIndex1749 := position, tokenIndex
																	if buffer[position] != rune('j') {
																		goto l1750
																	}
																	position++
																	goto l1749
																l1750:
																	position, tokenIndex = position1749, tokenIndex1749
																	if buffer[position] != rune('J') {
																		goto l1716
																	}
																	position++
																}
															l1749:
																{
																	position1751, tokenIndex1751 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1752
																	}
																	position++
																	goto l1751
																l1752:
																	position, tokenIndex = position1751, tokenIndex1751
																	if buffer[position] != rune('R') {
																		goto l1716
																	}
																	position++
																}
															l1751:
																if !_rules[rulews]() {
																	goto l1716
																}
																{
																	position1753, tokenIndex1753 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1753
																	}
																	if !_rules[rulesep]() {
																		goto l1753
																	}
																	goto l1754
																l1753:
																	position, tokenIndex = position1753, tokenIndex1753
																}
															l1754:
																if !_rules[ruledisp]() {
																	goto l1716
																}
																{
																	add(ruleAction103, position)
																}
																add(ruleJr, position1748)
															}
															break
														case 'R', 'r':
															{
																position1756 := position
																{
																	position1757, tokenIndex1757 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l1758
																	}
																	position++
																	goto l1757
																l1758:
																	position, tokenIndex = position1757, tokenIndex1757
																	if buffer[position] != rune('R') {
																		goto l1716
																	}
																	position++
																}
															l1757:
																{
																	position1759, tokenIndex1759 := position, tokenIndex
																	if buffer[position] != rune('e') {
																		goto l1760
																	}
																	position++
																	goto l1759
																l1760:
																	position, tokenIndex = position1759, tokenIndex1759
																	if buffer[position] != rune('E') {
																		goto l1716
																	}
																	position++
																}
															l1759:
																{
																	position1761, tokenIndex1761 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l1762
																	}
																	position++
																	goto l1761
																l1762:
																	position, tokenIndex = position1761, tokenIndex1761
																	if buffer[position] != rune('T') {
																		goto l1716
																	}
																	position++
																}
															l1761:
																{
																	position1763, tokenIndex1763 := position, tokenIndex
																	if !_rules[rulews]() {
																		goto l1763
																	}
																	if !_rules[rulecc]() {
																		goto l1763
																	}
																	goto l1764
																l1763:
																	position, tokenIndex = position1763, tokenIndex1763
																}
															l1764:
																{
																	add(ruleAction101, position)
																}
																add(ruleRet, position1756)
															}
															break
														default:
															{
																position1766 := position
																{
																	position1767, tokenIndex1767 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l1768
																	}
																	position++
																	goto l1767
																l1768:
																	position, tokenIndex = position1767, tokenIndex1767
																	if buffer[position] != rune('C') {
																		goto l1716
																	}
																	position++
																}
															l1767:
																{
																	position1769, tokenIndex1769 := position, tokenIndex
																	if buffer[position] != rune('a') {
																		goto l1770
																	}
																	position++
																	goto l1769
																l1770:
																	position, tokenIndex = position1769, tokenIndex1769
																	if buffer[position] != rune('A') {
																		goto l1716
																	}
																	position++
																}
															l1769:
																{
																	position1771, tokenIndex1771 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1772
																	}
																	position++
																	goto l1771
																l1772:
																	position, tokenIndex = position1771, tokenIndex1771
																	if buffer[position] != rune('L') {
																		goto l1716
																	}
																	position++
																}
															l1771:
																{
																	position1773, tokenIndex1773 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l1774
																	}
																	position++
																	goto l1773
																l1774:
																	position, tokenIndex = position1773, tokenIndex1773
																	if buffer[position] != rune('L') {
																		goto l1716
																	}
																	position++
																}
															l1773:
																if !_rules[rulews]() {
																	goto l1716
																}
																{
																	position1775, tokenIndex1775 := position, tokenIndex
																	if !_rules[rulecc]() {
																		goto l1775
																	}
																	if !_rules[rulesep]() {
																		goto l1775
																	}
																	goto l1776
																l1775:
																	position, tokenIndex = position1775, tokenIndex1775
																}
															l1776:
																if !_rules[ruleSrc16]() {
																	goto l1716
																}
																{
																	add(ruleAction100, position)
																}
																add(ruleCall, position1766)
															}
															break
														}
													}

												}
											l1718:
												add(ruleJump, position1717)
											}
											goto l993
										l1716:
											position, tokenIndex = position993, tokenIndex993
											{
												position1778 := position
												{
													position1779, tokenIndex1779 := position, tokenIndex
													{
														position1781 := position
														{
															position1782, tokenIndex1782 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1783
															}
															position++
															goto l1782
														l1783:
															position, tokenIndex = position1782, tokenIndex1782
															if buffer[position] != rune('I') {
																goto l1780
															}
															position++
														}
													l1782:
														{
															position1784, tokenIndex1784 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1785
															}
															position++
															goto l1784
														l1785:
															position, tokenIndex = position1784, tokenIndex1784
															if buffer[position] != rune('N') {
																goto l1780
															}
															position++
														}
													l1784:
														if !_rules[rulews]() {
															goto l1780
														}
														if !_rules[ruleReg8]() {
															goto l1780
														}
														if !_rules[rulesep]() {
															goto l1780
														}
														if !_rules[rulePort]() {
															goto l1780
														}
														{
															add(ruleAction105, position)
														}
														add(ruleIN, position1781)
													}
													goto l1779
												l1780:
													position, tokenIndex = position1779, tokenIndex1779
													{
														position1787 := position
														{
															position1788, tokenIndex1788 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1789
															}
															position++
															goto l1788
														l1789:
															position, tokenIndex = position1788, tokenIndex1788
															if buffer[position] != rune('O') {
																goto l919
															}
															position++
														}
													l1788:
														{
															position1790, tokenIndex1790 := position, tokenIndex
															if buffer[position] != rune('u') {
																goto l1791
															}
															position++
															goto l1790
														l1791:
															position, tokenIndex = position1790, tokenIndex1790
															if buffer[position] != rune('U') {
																goto l919
															}
															position++
														}
													l1790:
														{
															position1792, tokenIndex1792 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1793
															}
															position++
															goto l1792
														l1793:
															position, tokenIndex = position1792, tokenIndex1792
															if buffer[position] != rune('T') {
																goto l919
															}
															position++
														}
													l1792:
														if !_rules[rulews]() {
															goto l919
														}
														if !_rules[rulePort]() {
															goto l919
														}
														if !_rules[rulesep]() {
															goto l919
														}
														if !_rules[ruleReg8]() {
															goto l919
														}
														{
															add(ruleAction106, position)
														}
														add(ruleOUT, position1787)
													}
												}
											l1779:
												add(ruleIO, position1778)
											}
										}
									l993:
										add(ruleInstruction, position992)
									}
								}
							l922:
								add(ruleStatement, position921)
							}
							goto l920
						l919:
							position, tokenIndex = position919, tokenIndex919
						}
					l920:
						{
							position1795, tokenIndex1795 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1795
							}
							goto l1796
						l1795:
							position, tokenIndex = position1795, tokenIndex1795
						}
					l1796:
						{
							position1797, tokenIndex1797 := position, tokenIndex
							{
								position1799 := position
								{
									position1800, tokenIndex1800 := position, tokenIndex
									if buffer[position] != rune(';') {
										goto l1801
									}
									position++
									goto l1800
								l1801:
									position, tokenIndex = position1800, tokenIndex1800
									if buffer[position] != rune('#') {
										goto l1797
									}
									position++
								}
							l1800:
							l1802:
								{
									position1803, tokenIndex1803 := position, tokenIndex
									{
										position1804, tokenIndex1804 := position, tokenIndex
										if buffer[position] != rune('\n') {
											goto l1804
										}
										position++
										goto l1803
									l1804:
										position, tokenIndex = position1804, tokenIndex1804
									}
									if !matchDot() {
										goto l1803
									}
									goto l1802
								l1803:
									position, tokenIndex = position1803, tokenIndex1803
								}
								add(ruleComment, position1799)
							}
							goto l1798
						l1797:
							position, tokenIndex = position1797, tokenIndex1797
						}
					l1798:
						{
							position1805, tokenIndex1805 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1805
							}
							goto l1806
						l1805:
							position, tokenIndex = position1805, tokenIndex1805
						}
					l1806:
						{
							position1807, tokenIndex1807 := position, tokenIndex
							{
								position1809, tokenIndex1809 := position, tokenIndex
								if buffer[position] != rune('\r') {
									goto l1809
								}
								position++
								goto l1810
							l1809:
								position, tokenIndex = position1809, tokenIndex1809
							}
						l1810:
							if buffer[position] != rune('\n') {
								goto l1808
							}
							position++
							goto l1807
						l1808:
							position, tokenIndex = position1807, tokenIndex1807
							if buffer[position] != rune(':') {
								goto l3
							}
							position++
						}
					l1807:
						{
							add(ruleAction0, position)
						}
						add(ruleLine, position908)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position1812, tokenIndex1812 := position, tokenIndex
					if !matchDot() {
						goto l1812
					}
					goto l0
				l1812:
					position, tokenIndex = position1812, tokenIndex1812
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
		/* 11 LabelText <- <<(alphaund alphaundnum alphaundnum+)>> */
		func() bool {
			position1823, tokenIndex1823 := position, tokenIndex
			{
				position1824 := position
				{
					position1825 := position
					if !_rules[rulealphaund]() {
						goto l1823
					}
					if !_rules[rulealphaundnum]() {
						goto l1823
					}
					if !_rules[rulealphaundnum]() {
						goto l1823
					}
				l1826:
					{
						position1827, tokenIndex1827 := position, tokenIndex
						if !_rules[rulealphaundnum]() {
							goto l1827
						}
						goto l1826
					l1827:
						position, tokenIndex = position1827, tokenIndex1827
					}
					add(rulePegText, position1825)
				}
				add(ruleLabelText, position1824)
			}
			return true
		l1823:
			position, tokenIndex = position1823, tokenIndex1823
			return false
		},
		/* 12 alphaundnum <- <(alphaund / num)> */
		func() bool {
			position1828, tokenIndex1828 := position, tokenIndex
			{
				position1829 := position
				{
					position1830, tokenIndex1830 := position, tokenIndex
					if !_rules[rulealphaund]() {
						goto l1831
					}
					goto l1830
				l1831:
					position, tokenIndex = position1830, tokenIndex1830
					{
						position1832 := position
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1828
						}
						position++
						add(rulenum, position1832)
					}
				}
			l1830:
				add(rulealphaundnum, position1829)
			}
			return true
		l1828:
			position, tokenIndex = position1828, tokenIndex1828
			return false
		},
		/* 13 alphaund <- <((&('_') '_') | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]))> */
		func() bool {
			position1833, tokenIndex1833 := position, tokenIndex
			{
				position1834 := position
				{
					switch buffer[position] {
					case '_':
						if buffer[position] != rune('_') {
							goto l1833
						}
						position++
						break
					case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l1833
						}
						position++
						break
					default:
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l1833
						}
						position++
						break
					}
				}

				add(rulealphaund, position1834)
			}
			return true
		l1833:
			position, tokenIndex = position1833, tokenIndex1833
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
			position1859, tokenIndex1859 := position, tokenIndex
			{
				position1860 := position
				{
					position1861, tokenIndex1861 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1862
					}
					goto l1861
				l1862:
					position, tokenIndex = position1861, tokenIndex1861
					if !_rules[ruleReg8]() {
						goto l1863
					}
					goto l1861
				l1863:
					position, tokenIndex = position1861, tokenIndex1861
					if !_rules[ruleReg16Contents]() {
						goto l1864
					}
					goto l1861
				l1864:
					position, tokenIndex = position1861, tokenIndex1861
					if !_rules[rulenn_contents]() {
						goto l1859
					}
				}
			l1861:
				{
					add(ruleAction21, position)
				}
				add(ruleSrc8, position1860)
			}
			return true
		l1859:
			position, tokenIndex = position1859, tokenIndex1859
			return false
		},
		/* 38 Loc8 <- <((Reg8 / Reg16Contents) Action22)> */
		func() bool {
			position1866, tokenIndex1866 := position, tokenIndex
			{
				position1867 := position
				{
					position1868, tokenIndex1868 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1869
					}
					goto l1868
				l1869:
					position, tokenIndex = position1868, tokenIndex1868
					if !_rules[ruleReg16Contents]() {
						goto l1866
					}
				}
			l1868:
				{
					add(ruleAction22, position)
				}
				add(ruleLoc8, position1867)
			}
			return true
		l1866:
			position, tokenIndex = position1866, tokenIndex1866
			return false
		},
		/* 39 Copy8 <- <(Reg8 Action23)> */
		func() bool {
			position1871, tokenIndex1871 := position, tokenIndex
			{
				position1872 := position
				if !_rules[ruleReg8]() {
					goto l1871
				}
				{
					add(ruleAction23, position)
				}
				add(ruleCopy8, position1872)
			}
			return true
		l1871:
			position, tokenIndex = position1871, tokenIndex1871
			return false
		},
		/* 40 ILoc8 <- <(IReg8 Action24)> */
		func() bool {
			position1874, tokenIndex1874 := position, tokenIndex
			{
				position1875 := position
				if !_rules[ruleIReg8]() {
					goto l1874
				}
				{
					add(ruleAction24, position)
				}
				add(ruleILoc8, position1875)
			}
			return true
		l1874:
			position, tokenIndex = position1874, tokenIndex1874
			return false
		},
		/* 41 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action25)> */
		func() bool {
			position1877, tokenIndex1877 := position, tokenIndex
			{
				position1878 := position
				{
					position1879 := position
					{
						position1880, tokenIndex1880 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1881
						}
						goto l1880
					l1881:
						position, tokenIndex = position1880, tokenIndex1880
						{
							switch buffer[position] {
							case 'R', 'r':
								{
									position1883 := position
									{
										position1884, tokenIndex1884 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l1885
										}
										position++
										goto l1884
									l1885:
										position, tokenIndex = position1884, tokenIndex1884
										if buffer[position] != rune('R') {
											goto l1877
										}
										position++
									}
								l1884:
									add(ruleR, position1883)
								}
								break
							case 'I', 'i':
								{
									position1886 := position
									{
										position1887, tokenIndex1887 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1888
										}
										position++
										goto l1887
									l1888:
										position, tokenIndex = position1887, tokenIndex1887
										if buffer[position] != rune('I') {
											goto l1877
										}
										position++
									}
								l1887:
									add(ruleI, position1886)
								}
								break
							case 'L', 'l':
								{
									position1889 := position
									{
										position1890, tokenIndex1890 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1891
										}
										position++
										goto l1890
									l1891:
										position, tokenIndex = position1890, tokenIndex1890
										if buffer[position] != rune('L') {
											goto l1877
										}
										position++
									}
								l1890:
									add(ruleL, position1889)
								}
								break
							case 'H', 'h':
								{
									position1892 := position
									{
										position1893, tokenIndex1893 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1894
										}
										position++
										goto l1893
									l1894:
										position, tokenIndex = position1893, tokenIndex1893
										if buffer[position] != rune('H') {
											goto l1877
										}
										position++
									}
								l1893:
									add(ruleH, position1892)
								}
								break
							case 'E', 'e':
								{
									position1895 := position
									{
										position1896, tokenIndex1896 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1897
										}
										position++
										goto l1896
									l1897:
										position, tokenIndex = position1896, tokenIndex1896
										if buffer[position] != rune('E') {
											goto l1877
										}
										position++
									}
								l1896:
									add(ruleE, position1895)
								}
								break
							case 'D', 'd':
								{
									position1898 := position
									{
										position1899, tokenIndex1899 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1900
										}
										position++
										goto l1899
									l1900:
										position, tokenIndex = position1899, tokenIndex1899
										if buffer[position] != rune('D') {
											goto l1877
										}
										position++
									}
								l1899:
									add(ruleD, position1898)
								}
								break
							case 'C', 'c':
								{
									position1901 := position
									{
										position1902, tokenIndex1902 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1903
										}
										position++
										goto l1902
									l1903:
										position, tokenIndex = position1902, tokenIndex1902
										if buffer[position] != rune('C') {
											goto l1877
										}
										position++
									}
								l1902:
									add(ruleC, position1901)
								}
								break
							case 'B', 'b':
								{
									position1904 := position
									{
										position1905, tokenIndex1905 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1906
										}
										position++
										goto l1905
									l1906:
										position, tokenIndex = position1905, tokenIndex1905
										if buffer[position] != rune('B') {
											goto l1877
										}
										position++
									}
								l1905:
									add(ruleB, position1904)
								}
								break
							case 'F', 'f':
								{
									position1907 := position
									{
										position1908, tokenIndex1908 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1909
										}
										position++
										goto l1908
									l1909:
										position, tokenIndex = position1908, tokenIndex1908
										if buffer[position] != rune('F') {
											goto l1877
										}
										position++
									}
								l1908:
									add(ruleF, position1907)
								}
								break
							default:
								{
									position1910 := position
									{
										position1911, tokenIndex1911 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1912
										}
										position++
										goto l1911
									l1912:
										position, tokenIndex = position1911, tokenIndex1911
										if buffer[position] != rune('A') {
											goto l1877
										}
										position++
									}
								l1911:
									add(ruleA, position1910)
								}
								break
							}
						}

					}
				l1880:
					add(rulePegText, position1879)
				}
				{
					add(ruleAction25, position)
				}
				add(ruleReg8, position1878)
			}
			return true
		l1877:
			position, tokenIndex = position1877, tokenIndex1877
			return false
		},
		/* 42 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action26)> */
		func() bool {
			position1914, tokenIndex1914 := position, tokenIndex
			{
				position1915 := position
				{
					position1916 := position
					{
						position1917, tokenIndex1917 := position, tokenIndex
						{
							position1919 := position
							{
								position1920, tokenIndex1920 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1921
								}
								position++
								goto l1920
							l1921:
								position, tokenIndex = position1920, tokenIndex1920
								if buffer[position] != rune('I') {
									goto l1918
								}
								position++
							}
						l1920:
							{
								position1922, tokenIndex1922 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1923
								}
								position++
								goto l1922
							l1923:
								position, tokenIndex = position1922, tokenIndex1922
								if buffer[position] != rune('X') {
									goto l1918
								}
								position++
							}
						l1922:
							{
								position1924, tokenIndex1924 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1925
								}
								position++
								goto l1924
							l1925:
								position, tokenIndex = position1924, tokenIndex1924
								if buffer[position] != rune('H') {
									goto l1918
								}
								position++
							}
						l1924:
							add(ruleIXH, position1919)
						}
						goto l1917
					l1918:
						position, tokenIndex = position1917, tokenIndex1917
						{
							position1927 := position
							{
								position1928, tokenIndex1928 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1929
								}
								position++
								goto l1928
							l1929:
								position, tokenIndex = position1928, tokenIndex1928
								if buffer[position] != rune('I') {
									goto l1926
								}
								position++
							}
						l1928:
							{
								position1930, tokenIndex1930 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1931
								}
								position++
								goto l1930
							l1931:
								position, tokenIndex = position1930, tokenIndex1930
								if buffer[position] != rune('X') {
									goto l1926
								}
								position++
							}
						l1930:
							{
								position1932, tokenIndex1932 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1933
								}
								position++
								goto l1932
							l1933:
								position, tokenIndex = position1932, tokenIndex1932
								if buffer[position] != rune('L') {
									goto l1926
								}
								position++
							}
						l1932:
							add(ruleIXL, position1927)
						}
						goto l1917
					l1926:
						position, tokenIndex = position1917, tokenIndex1917
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
									goto l1934
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
									goto l1934
								}
								position++
							}
						l1938:
							{
								position1940, tokenIndex1940 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1941
								}
								position++
								goto l1940
							l1941:
								position, tokenIndex = position1940, tokenIndex1940
								if buffer[position] != rune('H') {
									goto l1934
								}
								position++
							}
						l1940:
							add(ruleIYH, position1935)
						}
						goto l1917
					l1934:
						position, tokenIndex = position1917, tokenIndex1917
						{
							position1942 := position
							{
								position1943, tokenIndex1943 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1944
								}
								position++
								goto l1943
							l1944:
								position, tokenIndex = position1943, tokenIndex1943
								if buffer[position] != rune('I') {
									goto l1914
								}
								position++
							}
						l1943:
							{
								position1945, tokenIndex1945 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1946
								}
								position++
								goto l1945
							l1946:
								position, tokenIndex = position1945, tokenIndex1945
								if buffer[position] != rune('Y') {
									goto l1914
								}
								position++
							}
						l1945:
							{
								position1947, tokenIndex1947 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1948
								}
								position++
								goto l1947
							l1948:
								position, tokenIndex = position1947, tokenIndex1947
								if buffer[position] != rune('L') {
									goto l1914
								}
								position++
							}
						l1947:
							add(ruleIYL, position1942)
						}
					}
				l1917:
					add(rulePegText, position1916)
				}
				{
					add(ruleAction26, position)
				}
				add(ruleIReg8, position1915)
			}
			return true
		l1914:
			position, tokenIndex = position1914, tokenIndex1914
			return false
		},
		/* 43 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action27)> */
		func() bool {
			position1950, tokenIndex1950 := position, tokenIndex
			{
				position1951 := position
				{
					position1952, tokenIndex1952 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1953
					}
					goto l1952
				l1953:
					position, tokenIndex = position1952, tokenIndex1952
					if !_rules[rulenn_contents]() {
						goto l1954
					}
					goto l1952
				l1954:
					position, tokenIndex = position1952, tokenIndex1952
					if !_rules[ruleReg16Contents]() {
						goto l1950
					}
				}
			l1952:
				{
					add(ruleAction27, position)
				}
				add(ruleDst16, position1951)
			}
			return true
		l1950:
			position, tokenIndex = position1950, tokenIndex1950
			return false
		},
		/* 44 Src16 <- <((Reg16 / nn / nn_contents) Action28)> */
		func() bool {
			position1956, tokenIndex1956 := position, tokenIndex
			{
				position1957 := position
				{
					position1958, tokenIndex1958 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1959
					}
					goto l1958
				l1959:
					position, tokenIndex = position1958, tokenIndex1958
					if !_rules[rulenn]() {
						goto l1960
					}
					goto l1958
				l1960:
					position, tokenIndex = position1958, tokenIndex1958
					if !_rules[rulenn_contents]() {
						goto l1956
					}
				}
			l1958:
				{
					add(ruleAction28, position)
				}
				add(ruleSrc16, position1957)
			}
			return true
		l1956:
			position, tokenIndex = position1956, tokenIndex1956
			return false
		},
		/* 45 Loc16 <- <(Reg16 Action29)> */
		func() bool {
			position1962, tokenIndex1962 := position, tokenIndex
			{
				position1963 := position
				if !_rules[ruleReg16]() {
					goto l1962
				}
				{
					add(ruleAction29, position)
				}
				add(ruleLoc16, position1963)
			}
			return true
		l1962:
			position, tokenIndex = position1962, tokenIndex1962
			return false
		},
		/* 46 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action30)> */
		func() bool {
			position1965, tokenIndex1965 := position, tokenIndex
			{
				position1966 := position
				{
					position1967 := position
					{
						position1968, tokenIndex1968 := position, tokenIndex
						{
							position1970 := position
							{
								position1971, tokenIndex1971 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l1972
								}
								position++
								goto l1971
							l1972:
								position, tokenIndex = position1971, tokenIndex1971
								if buffer[position] != rune('A') {
									goto l1969
								}
								position++
							}
						l1971:
							{
								position1973, tokenIndex1973 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1974
								}
								position++
								goto l1973
							l1974:
								position, tokenIndex = position1973, tokenIndex1973
								if buffer[position] != rune('F') {
									goto l1969
								}
								position++
							}
						l1973:
							if buffer[position] != rune('\'') {
								goto l1969
							}
							position++
							add(ruleAF_PRIME, position1970)
						}
						goto l1968
					l1969:
						position, tokenIndex = position1968, tokenIndex1968
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1965
								}
								break
							case 'S', 's':
								{
									position1976 := position
									{
										position1977, tokenIndex1977 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1978
										}
										position++
										goto l1977
									l1978:
										position, tokenIndex = position1977, tokenIndex1977
										if buffer[position] != rune('S') {
											goto l1965
										}
										position++
									}
								l1977:
									{
										position1979, tokenIndex1979 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1980
										}
										position++
										goto l1979
									l1980:
										position, tokenIndex = position1979, tokenIndex1979
										if buffer[position] != rune('P') {
											goto l1965
										}
										position++
									}
								l1979:
									add(ruleSP, position1976)
								}
								break
							case 'H', 'h':
								{
									position1981 := position
									{
										position1982, tokenIndex1982 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1983
										}
										position++
										goto l1982
									l1983:
										position, tokenIndex = position1982, tokenIndex1982
										if buffer[position] != rune('H') {
											goto l1965
										}
										position++
									}
								l1982:
									{
										position1984, tokenIndex1984 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1985
										}
										position++
										goto l1984
									l1985:
										position, tokenIndex = position1984, tokenIndex1984
										if buffer[position] != rune('L') {
											goto l1965
										}
										position++
									}
								l1984:
									add(ruleHL, position1981)
								}
								break
							case 'D', 'd':
								{
									position1986 := position
									{
										position1987, tokenIndex1987 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1988
										}
										position++
										goto l1987
									l1988:
										position, tokenIndex = position1987, tokenIndex1987
										if buffer[position] != rune('D') {
											goto l1965
										}
										position++
									}
								l1987:
									{
										position1989, tokenIndex1989 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1990
										}
										position++
										goto l1989
									l1990:
										position, tokenIndex = position1989, tokenIndex1989
										if buffer[position] != rune('E') {
											goto l1965
										}
										position++
									}
								l1989:
									add(ruleDE, position1986)
								}
								break
							case 'B', 'b':
								{
									position1991 := position
									{
										position1992, tokenIndex1992 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1993
										}
										position++
										goto l1992
									l1993:
										position, tokenIndex = position1992, tokenIndex1992
										if buffer[position] != rune('B') {
											goto l1965
										}
										position++
									}
								l1992:
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
											goto l1965
										}
										position++
									}
								l1994:
									add(ruleBC, position1991)
								}
								break
							default:
								{
									position1996 := position
									{
										position1997, tokenIndex1997 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1998
										}
										position++
										goto l1997
									l1998:
										position, tokenIndex = position1997, tokenIndex1997
										if buffer[position] != rune('A') {
											goto l1965
										}
										position++
									}
								l1997:
									{
										position1999, tokenIndex1999 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l2000
										}
										position++
										goto l1999
									l2000:
										position, tokenIndex = position1999, tokenIndex1999
										if buffer[position] != rune('F') {
											goto l1965
										}
										position++
									}
								l1999:
									add(ruleAF, position1996)
								}
								break
							}
						}

					}
				l1968:
					add(rulePegText, position1967)
				}
				{
					add(ruleAction30, position)
				}
				add(ruleReg16, position1966)
			}
			return true
		l1965:
			position, tokenIndex = position1965, tokenIndex1965
			return false
		},
		/* 47 IReg16 <- <(<(IX / IY)> Action31)> */
		func() bool {
			position2002, tokenIndex2002 := position, tokenIndex
			{
				position2003 := position
				{
					position2004 := position
					{
						position2005, tokenIndex2005 := position, tokenIndex
						{
							position2007 := position
							{
								position2008, tokenIndex2008 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l2009
								}
								position++
								goto l2008
							l2009:
								position, tokenIndex = position2008, tokenIndex2008
								if buffer[position] != rune('I') {
									goto l2006
								}
								position++
							}
						l2008:
							{
								position2010, tokenIndex2010 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l2011
								}
								position++
								goto l2010
							l2011:
								position, tokenIndex = position2010, tokenIndex2010
								if buffer[position] != rune('X') {
									goto l2006
								}
								position++
							}
						l2010:
							add(ruleIX, position2007)
						}
						goto l2005
					l2006:
						position, tokenIndex = position2005, tokenIndex2005
						{
							position2012 := position
							{
								position2013, tokenIndex2013 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l2014
								}
								position++
								goto l2013
							l2014:
								position, tokenIndex = position2013, tokenIndex2013
								if buffer[position] != rune('I') {
									goto l2002
								}
								position++
							}
						l2013:
							{
								position2015, tokenIndex2015 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l2016
								}
								position++
								goto l2015
							l2016:
								position, tokenIndex = position2015, tokenIndex2015
								if buffer[position] != rune('Y') {
									goto l2002
								}
								position++
							}
						l2015:
							add(ruleIY, position2012)
						}
					}
				l2005:
					add(rulePegText, position2004)
				}
				{
					add(ruleAction31, position)
				}
				add(ruleIReg16, position2003)
			}
			return true
		l2002:
			position, tokenIndex = position2002, tokenIndex2002
			return false
		},
		/* 48 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position2018, tokenIndex2018 := position, tokenIndex
			{
				position2019 := position
				{
					position2020, tokenIndex2020 := position, tokenIndex
					{
						position2022 := position
						if buffer[position] != rune('(') {
							goto l2021
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l2021
						}
						{
							position2023, tokenIndex2023 := position, tokenIndex
							if !_rules[rulews]() {
								goto l2023
							}
							goto l2024
						l2023:
							position, tokenIndex = position2023, tokenIndex2023
						}
					l2024:
						if !_rules[ruledisp]() {
							goto l2021
						}
						{
							position2025, tokenIndex2025 := position, tokenIndex
							if !_rules[rulews]() {
								goto l2025
							}
							goto l2026
						l2025:
							position, tokenIndex = position2025, tokenIndex2025
						}
					l2026:
						if buffer[position] != rune(')') {
							goto l2021
						}
						position++
						{
							add(ruleAction33, position)
						}
						add(ruleIndexedR16C, position2022)
					}
					goto l2020
				l2021:
					position, tokenIndex = position2020, tokenIndex2020
					{
						position2028 := position
						if buffer[position] != rune('(') {
							goto l2018
						}
						position++
						if !_rules[ruleReg16]() {
							goto l2018
						}
						if buffer[position] != rune(')') {
							goto l2018
						}
						position++
						{
							add(ruleAction32, position)
						}
						add(rulePlainR16C, position2028)
					}
				}
			l2020:
				add(ruleReg16Contents, position2019)
			}
			return true
		l2018:
			position, tokenIndex = position2018, tokenIndex2018
			return false
		},
		/* 49 PlainR16C <- <('(' Reg16 ')' Action32)> */
		nil,
		/* 50 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action33)> */
		nil,
		/* 51 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position2032, tokenIndex2032 := position, tokenIndex
			{
				position2033 := position
				{
					position2034, tokenIndex2034 := position, tokenIndex
					{
						position2036 := position
						{
							position2037 := position
							if !_rules[rulehexdigit]() {
								goto l2035
							}
							if !_rules[rulehexdigit]() {
								goto l2035
							}
							add(rulePegText, position2037)
						}
						{
							position2038, tokenIndex2038 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2039
							}
							position++
							goto l2038
						l2039:
							position, tokenIndex = position2038, tokenIndex2038
							if buffer[position] != rune('H') {
								goto l2035
							}
							position++
						}
					l2038:
						{
							add(ruleAction37, position)
						}
						add(rulehexByteH, position2036)
					}
					goto l2034
				l2035:
					position, tokenIndex = position2034, tokenIndex2034
					{
						position2042 := position
						if buffer[position] != rune('0') {
							goto l2041
						}
						position++
						{
							position2043, tokenIndex2043 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2044
							}
							position++
							goto l2043
						l2044:
							position, tokenIndex = position2043, tokenIndex2043
							if buffer[position] != rune('X') {
								goto l2041
							}
							position++
						}
					l2043:
						{
							position2045 := position
							if !_rules[rulehexdigit]() {
								goto l2041
							}
							if !_rules[rulehexdigit]() {
								goto l2041
							}
							add(rulePegText, position2045)
						}
						{
							add(ruleAction38, position)
						}
						add(rulehexByte0x, position2042)
					}
					goto l2034
				l2041:
					position, tokenIndex = position2034, tokenIndex2034
					{
						position2047 := position
						{
							position2048 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2032
							}
							position++
						l2049:
							{
								position2050, tokenIndex2050 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2050
								}
								position++
								goto l2049
							l2050:
								position, tokenIndex = position2050, tokenIndex2050
							}
							add(rulePegText, position2048)
						}
						{
							add(ruleAction39, position)
						}
						add(ruledecimalByte, position2047)
					}
				}
			l2034:
				add(rulen, position2033)
			}
			return true
		l2032:
			position, tokenIndex = position2032, tokenIndex2032
			return false
		},
		/* 52 nn <- <(LabelNN / hexWordH / hexWord0x)> */
		func() bool {
			position2052, tokenIndex2052 := position, tokenIndex
			{
				position2053 := position
				{
					position2054, tokenIndex2054 := position, tokenIndex
					{
						position2056 := position
						{
							position2057 := position
							if !_rules[ruleLabelText]() {
								goto l2055
							}
							add(rulePegText, position2057)
						}
						{
							add(ruleAction40, position)
						}
						add(ruleLabelNN, position2056)
					}
					goto l2054
				l2055:
					position, tokenIndex = position2054, tokenIndex2054
					{
						position2060 := position
						{
							position2061, tokenIndex2061 := position, tokenIndex
							if !_rules[rulezeroHexWord]() {
								goto l2062
							}
							goto l2061
						l2062:
							position, tokenIndex = position2061, tokenIndex2061
							if !_rules[rulehexWord]() {
								goto l2059
							}
						}
					l2061:
						{
							position2063, tokenIndex2063 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2064
							}
							position++
							goto l2063
						l2064:
							position, tokenIndex = position2063, tokenIndex2063
							if buffer[position] != rune('H') {
								goto l2059
							}
							position++
						}
					l2063:
						add(rulehexWordH, position2060)
					}
					goto l2054
				l2059:
					position, tokenIndex = position2054, tokenIndex2054
					{
						position2065 := position
						if buffer[position] != rune('0') {
							goto l2052
						}
						position++
						{
							position2066, tokenIndex2066 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l2067
							}
							position++
							goto l2066
						l2067:
							position, tokenIndex = position2066, tokenIndex2066
							if buffer[position] != rune('X') {
								goto l2052
							}
							position++
						}
					l2066:
						{
							position2068, tokenIndex2068 := position, tokenIndex
							if !_rules[rulezeroHexWord]() {
								goto l2069
							}
							goto l2068
						l2069:
							position, tokenIndex = position2068, tokenIndex2068
							if !_rules[rulehexWord]() {
								goto l2052
							}
						}
					l2068:
						add(rulehexWord0x, position2065)
					}
				}
			l2054:
				add(rulenn, position2053)
			}
			return true
		l2052:
			position, tokenIndex = position2052, tokenIndex2052
			return false
		},
		/* 53 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position2070, tokenIndex2070 := position, tokenIndex
			{
				position2071 := position
				{
					position2072, tokenIndex2072 := position, tokenIndex
					{
						position2074 := position
						{
							position2075 := position
							{
								position2076, tokenIndex2076 := position, tokenIndex
								{
									position2078, tokenIndex2078 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2079
									}
									position++
									goto l2078
								l2079:
									position, tokenIndex = position2078, tokenIndex2078
									if buffer[position] != rune('+') {
										goto l2076
									}
									position++
								}
							l2078:
								goto l2077
							l2076:
								position, tokenIndex = position2076, tokenIndex2076
							}
						l2077:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2073
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2073
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2073
									}
									position++
									break
								}
							}

						l2080:
							{
								position2081, tokenIndex2081 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2081
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2081
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2081
										}
										position++
										break
									}
								}

								goto l2080
							l2081:
								position, tokenIndex = position2081, tokenIndex2081
							}
							add(rulePegText, position2075)
						}
						{
							position2084, tokenIndex2084 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l2085
							}
							position++
							goto l2084
						l2085:
							position, tokenIndex = position2084, tokenIndex2084
							if buffer[position] != rune('H') {
								goto l2073
							}
							position++
						}
					l2084:
						{
							add(ruleAction35, position)
						}
						add(rulesignedHexByteH, position2074)
					}
					goto l2072
				l2073:
					position, tokenIndex = position2072, tokenIndex2072
					{
						position2088 := position
						{
							position2089 := position
							{
								position2090, tokenIndex2090 := position, tokenIndex
								{
									position2092, tokenIndex2092 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2093
									}
									position++
									goto l2092
								l2093:
									position, tokenIndex = position2092, tokenIndex2092
									if buffer[position] != rune('+') {
										goto l2090
									}
									position++
								}
							l2092:
								goto l2091
							l2090:
								position, tokenIndex = position2090, tokenIndex2090
							}
						l2091:
							if buffer[position] != rune('0') {
								goto l2087
							}
							position++
							{
								position2094, tokenIndex2094 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l2095
								}
								position++
								goto l2094
							l2095:
								position, tokenIndex = position2094, tokenIndex2094
								if buffer[position] != rune('X') {
									goto l2087
								}
								position++
							}
						l2094:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l2087
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l2087
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l2087
									}
									position++
									break
								}
							}

						l2096:
							{
								position2097, tokenIndex2097 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l2097
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l2097
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l2097
										}
										position++
										break
									}
								}

								goto l2096
							l2097:
								position, tokenIndex = position2097, tokenIndex2097
							}
							add(rulePegText, position2089)
						}
						{
							add(ruleAction36, position)
						}
						add(rulesignedHexByte0x, position2088)
					}
					goto l2072
				l2087:
					position, tokenIndex = position2072, tokenIndex2072
					{
						position2101 := position
						{
							position2102 := position
							{
								position2103, tokenIndex2103 := position, tokenIndex
								{
									position2105, tokenIndex2105 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l2106
									}
									position++
									goto l2105
								l2106:
									position, tokenIndex = position2105, tokenIndex2105
									if buffer[position] != rune('+') {
										goto l2103
									}
									position++
								}
							l2105:
								goto l2104
							l2103:
								position, tokenIndex = position2103, tokenIndex2103
							}
						l2104:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l2070
							}
							position++
						l2107:
							{
								position2108, tokenIndex2108 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l2108
								}
								position++
								goto l2107
							l2108:
								position, tokenIndex = position2108, tokenIndex2108
							}
							add(rulePegText, position2102)
						}
						{
							add(ruleAction34, position)
						}
						add(rulesignedDecimalByte, position2101)
					}
				}
			l2072:
				add(ruledisp, position2071)
			}
			return true
		l2070:
			position, tokenIndex = position2070, tokenIndex2070
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
			position2119, tokenIndex2119 := position, tokenIndex
			{
				position2120 := position
				{
					position2121 := position
					if !_rules[rulehexdigit]() {
						goto l2119
					}
					if !_rules[rulehexdigit]() {
						goto l2119
					}
					if !_rules[rulehexdigit]() {
						goto l2119
					}
					if !_rules[rulehexdigit]() {
						goto l2119
					}
					add(rulePegText, position2121)
				}
				{
					add(ruleAction41, position)
				}
				add(rulehexWord, position2120)
			}
			return true
		l2119:
			position, tokenIndex = position2119, tokenIndex2119
			return false
		},
		/* 64 zeroHexWord <- <('0' hexWord)> */
		func() bool {
			position2123, tokenIndex2123 := position, tokenIndex
			{
				position2124 := position
				if buffer[position] != rune('0') {
					goto l2123
				}
				position++
				if !_rules[rulehexWord]() {
					goto l2123
				}
				add(rulezeroHexWord, position2124)
			}
			return true
		l2123:
			position, tokenIndex = position2123, tokenIndex2123
			return false
		},
		/* 65 nn_contents <- <('(' nn ')' Action42)> */
		func() bool {
			position2125, tokenIndex2125 := position, tokenIndex
			{
				position2126 := position
				if buffer[position] != rune('(') {
					goto l2125
				}
				position++
				if !_rules[rulenn]() {
					goto l2125
				}
				if buffer[position] != rune(')') {
					goto l2125
				}
				position++
				{
					add(ruleAction42, position)
				}
				add(rulenn_contents, position2126)
			}
			return true
		l2125:
			position, tokenIndex = position2125, tokenIndex2125
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
		/* 71 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws (('a' / 'A') sep)? Src8 Action47)> */
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
			position2201, tokenIndex2201 := position, tokenIndex
			{
				position2202 := position
				{
					position2203, tokenIndex2203 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2204
					}
					position++
					{
						position2205, tokenIndex2205 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l2206
						}
						position++
						goto l2205
					l2206:
						position, tokenIndex = position2205, tokenIndex2205
						if buffer[position] != rune('C') {
							goto l2204
						}
						position++
					}
				l2205:
					if buffer[position] != rune(')') {
						goto l2204
					}
					position++
					goto l2203
				l2204:
					position, tokenIndex = position2203, tokenIndex2203
					if buffer[position] != rune('(') {
						goto l2201
					}
					position++
					if !_rules[rulen]() {
						goto l2201
					}
					if buffer[position] != rune(')') {
						goto l2201
					}
					position++
				}
			l2203:
				add(rulePort, position2202)
			}
			return true
		l2201:
			position, tokenIndex = position2201, tokenIndex2201
			return false
		},
		/* 140 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2207, tokenIndex2207 := position, tokenIndex
			{
				position2208 := position
				{
					position2209, tokenIndex2209 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2209
					}
					goto l2210
				l2209:
					position, tokenIndex = position2209, tokenIndex2209
				}
			l2210:
				if buffer[position] != rune(',') {
					goto l2207
				}
				position++
				{
					position2211, tokenIndex2211 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2211
					}
					goto l2212
				l2211:
					position, tokenIndex = position2211, tokenIndex2211
				}
			l2212:
				add(rulesep, position2208)
			}
			return true
		l2207:
			position, tokenIndex = position2207, tokenIndex2207
			return false
		},
		/* 141 ws <- <(' ' / '\t')+> */
		func() bool {
			position2213, tokenIndex2213 := position, tokenIndex
			{
				position2214 := position
				{
					position2217, tokenIndex2217 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2218
					}
					position++
					goto l2217
				l2218:
					position, tokenIndex = position2217, tokenIndex2217
					if buffer[position] != rune('\t') {
						goto l2213
					}
					position++
				}
			l2217:
			l2215:
				{
					position2216, tokenIndex2216 := position, tokenIndex
					{
						position2219, tokenIndex2219 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l2220
						}
						position++
						goto l2219
					l2220:
						position, tokenIndex = position2219, tokenIndex2219
						if buffer[position] != rune('\t') {
							goto l2216
						}
						position++
					}
				l2219:
					goto l2215
				l2216:
					position, tokenIndex = position2216, tokenIndex2216
				}
				add(rulews, position2214)
			}
			return true
		l2213:
			position, tokenIndex = position2213, tokenIndex2213
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
			position2243, tokenIndex2243 := position, tokenIndex
			{
				position2244 := position
				{
					position2245, tokenIndex2245 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2246
					}
					position++
					goto l2245
				l2246:
					position, tokenIndex = position2245, tokenIndex2245
					{
						position2247, tokenIndex2247 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2248
						}
						position++
						goto l2247
					l2248:
						position, tokenIndex = position2247, tokenIndex2247
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2243
						}
						position++
					}
				l2247:
				}
			l2245:
				add(rulehexdigit, position2244)
			}
			return true
		l2243:
			position, tokenIndex = position2243, tokenIndex2243
			return false
		},
		/* 165 octaldigit <- <(<[0-7]> Action107)> */
		func() bool {
			position2249, tokenIndex2249 := position, tokenIndex
			{
				position2250 := position
				{
					position2251 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2249
					}
					position++
					add(rulePegText, position2251)
				}
				{
					add(ruleAction107, position)
				}
				add(ruleoctaldigit, position2250)
			}
			return true
		l2249:
			position, tokenIndex = position2249, tokenIndex2249
			return false
		},
		/* 166 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2253, tokenIndex2253 := position, tokenIndex
			{
				position2254 := position
				{
					position2255, tokenIndex2255 := position, tokenIndex
					{
						position2257 := position
						{
							position2258, tokenIndex2258 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2259
							}
							position++
							goto l2258
						l2259:
							position, tokenIndex = position2258, tokenIndex2258
							if buffer[position] != rune('N') {
								goto l2256
							}
							position++
						}
					l2258:
						{
							position2260, tokenIndex2260 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2261
							}
							position++
							goto l2260
						l2261:
							position, tokenIndex = position2260, tokenIndex2260
							if buffer[position] != rune('Z') {
								goto l2256
							}
							position++
						}
					l2260:
						{
							add(ruleAction108, position)
						}
						add(ruleFT_NZ, position2257)
					}
					goto l2255
				l2256:
					position, tokenIndex = position2255, tokenIndex2255
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
							if buffer[position] != rune('o') {
								goto l2268
							}
							position++
							goto l2267
						l2268:
							position, tokenIndex = position2267, tokenIndex2267
							if buffer[position] != rune('O') {
								goto l2263
							}
							position++
						}
					l2267:
						{
							add(ruleAction112, position)
						}
						add(ruleFT_PO, position2264)
					}
					goto l2255
				l2263:
					position, tokenIndex = position2255, tokenIndex2255
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
								goto l2270
							}
							position++
						}
					l2272:
						{
							position2274, tokenIndex2274 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2275
							}
							position++
							goto l2274
						l2275:
							position, tokenIndex = position2274, tokenIndex2274
							if buffer[position] != rune('E') {
								goto l2270
							}
							position++
						}
					l2274:
						{
							add(ruleAction113, position)
						}
						add(ruleFT_PE, position2271)
					}
					goto l2255
				l2270:
					position, tokenIndex = position2255, tokenIndex2255
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2278 := position
								{
									position2279, tokenIndex2279 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2280
									}
									position++
									goto l2279
								l2280:
									position, tokenIndex = position2279, tokenIndex2279
									if buffer[position] != rune('M') {
										goto l2253
									}
									position++
								}
							l2279:
								{
									add(ruleAction115, position)
								}
								add(ruleFT_M, position2278)
							}
							break
						case 'P', 'p':
							{
								position2282 := position
								{
									position2283, tokenIndex2283 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2284
									}
									position++
									goto l2283
								l2284:
									position, tokenIndex = position2283, tokenIndex2283
									if buffer[position] != rune('P') {
										goto l2253
									}
									position++
								}
							l2283:
								{
									add(ruleAction114, position)
								}
								add(ruleFT_P, position2282)
							}
							break
						case 'C', 'c':
							{
								position2286 := position
								{
									position2287, tokenIndex2287 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2288
									}
									position++
									goto l2287
								l2288:
									position, tokenIndex = position2287, tokenIndex2287
									if buffer[position] != rune('C') {
										goto l2253
									}
									position++
								}
							l2287:
								{
									add(ruleAction111, position)
								}
								add(ruleFT_C, position2286)
							}
							break
						case 'N', 'n':
							{
								position2290 := position
								{
									position2291, tokenIndex2291 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2292
									}
									position++
									goto l2291
								l2292:
									position, tokenIndex = position2291, tokenIndex2291
									if buffer[position] != rune('N') {
										goto l2253
									}
									position++
								}
							l2291:
								{
									position2293, tokenIndex2293 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2294
									}
									position++
									goto l2293
								l2294:
									position, tokenIndex = position2293, tokenIndex2293
									if buffer[position] != rune('C') {
										goto l2253
									}
									position++
								}
							l2293:
								{
									add(ruleAction110, position)
								}
								add(ruleFT_NC, position2290)
							}
							break
						default:
							{
								position2296 := position
								{
									position2297, tokenIndex2297 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2298
									}
									position++
									goto l2297
								l2298:
									position, tokenIndex = position2297, tokenIndex2297
									if buffer[position] != rune('Z') {
										goto l2253
									}
									position++
								}
							l2297:
								{
									add(ruleAction109, position)
								}
								add(ruleFT_Z, position2296)
							}
							break
						}
					}

				}
			l2255:
				add(rulecc, position2254)
			}
			return true
		l2253:
			position, tokenIndex = position2253, tokenIndex2253
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
