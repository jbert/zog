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
	"BlankLine",
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
	rules  [279]func() bool
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
							position10, tokenIndex10 := position, tokenIndex
							{
								position12 := position
								{
									position13 := position
									if !_rules[rulealpha]() {
										goto l10
									}
								l14:
									{
										position15, tokenIndex15 := position, tokenIndex
										{
											position16 := position
											{
												position17, tokenIndex17 := position, tokenIndex
												if !_rules[rulealpha]() {
													goto l18
												}
												goto l17
											l18:
												position, tokenIndex = position17, tokenIndex17
												{
													position19 := position
													if c := buffer[position]; c < rune('0') || c > rune('9') {
														goto l15
													}
													position++
													add(rulenum, position19)
												}
											}
										l17:
											add(rulealphanum, position16)
										}
										goto l14
									l15:
										position, tokenIndex = position15, tokenIndex15
									}
									add(rulePegText, position13)
								}
								if buffer[position] != rune(':') {
									goto l10
								}
								position++
								{
									add(ruleAction1, position)
								}
								add(ruleLabel, position12)
							}
							goto l11
						l10:
							position, tokenIndex = position10, tokenIndex10
						}
					l11:
					l21:
						{
							position22, tokenIndex22 := position, tokenIndex
							if !_rules[rulews]() {
								goto l22
							}
							goto l21
						l22:
							position, tokenIndex = position22, tokenIndex22
						}
						{
							position23 := position
							{
								position24, tokenIndex24 := position, tokenIndex
								{
									position26 := position
									{
										position27, tokenIndex27 := position, tokenIndex
										{
											position29 := position
											{
												position30, tokenIndex30 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l31
												}
												position++
												goto l30
											l31:
												position, tokenIndex = position30, tokenIndex30
												if buffer[position] != rune('P') {
													goto l28
												}
												position++
											}
										l30:
											{
												position32, tokenIndex32 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l33
												}
												position++
												goto l32
											l33:
												position, tokenIndex = position32, tokenIndex32
												if buffer[position] != rune('U') {
													goto l28
												}
												position++
											}
										l32:
											{
												position34, tokenIndex34 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l35
												}
												position++
												goto l34
											l35:
												position, tokenIndex = position34, tokenIndex34
												if buffer[position] != rune('S') {
													goto l28
												}
												position++
											}
										l34:
											{
												position36, tokenIndex36 := position, tokenIndex
												if buffer[position] != rune('h') {
													goto l37
												}
												position++
												goto l36
											l37:
												position, tokenIndex = position36, tokenIndex36
												if buffer[position] != rune('H') {
													goto l28
												}
												position++
											}
										l36:
											if !_rules[rulews]() {
												goto l28
											}
											if !_rules[ruleSrc16]() {
												goto l28
											}
											{
												add(ruleAction4, position)
											}
											add(rulePush, position29)
										}
										goto l27
									l28:
										position, tokenIndex = position27, tokenIndex27
										{
											switch buffer[position] {
											case 'E', 'e':
												{
													position40 := position
													{
														position41, tokenIndex41 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l42
														}
														position++
														goto l41
													l42:
														position, tokenIndex = position41, tokenIndex41
														if buffer[position] != rune('E') {
															goto l25
														}
														position++
													}
												l41:
													{
														position43, tokenIndex43 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l44
														}
														position++
														goto l43
													l44:
														position, tokenIndex = position43, tokenIndex43
														if buffer[position] != rune('X') {
															goto l25
														}
														position++
													}
												l43:
													if !_rules[rulews]() {
														goto l25
													}
													if !_rules[ruleDst16]() {
														goto l25
													}
													if !_rules[rulesep]() {
														goto l25
													}
													if !_rules[ruleSrc16]() {
														goto l25
													}
													{
														add(ruleAction6, position)
													}
													add(ruleEx, position40)
												}
												break
											case 'P', 'p':
												{
													position46 := position
													{
														position47, tokenIndex47 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l48
														}
														position++
														goto l47
													l48:
														position, tokenIndex = position47, tokenIndex47
														if buffer[position] != rune('P') {
															goto l25
														}
														position++
													}
												l47:
													{
														position49, tokenIndex49 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l50
														}
														position++
														goto l49
													l50:
														position, tokenIndex = position49, tokenIndex49
														if buffer[position] != rune('O') {
															goto l25
														}
														position++
													}
												l49:
													{
														position51, tokenIndex51 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l52
														}
														position++
														goto l51
													l52:
														position, tokenIndex = position51, tokenIndex51
														if buffer[position] != rune('P') {
															goto l25
														}
														position++
													}
												l51:
													if !_rules[rulews]() {
														goto l25
													}
													if !_rules[ruleDst16]() {
														goto l25
													}
													{
														add(ruleAction5, position)
													}
													add(rulePop, position46)
												}
												break
											default:
												{
													position54 := position
													{
														position55, tokenIndex55 := position, tokenIndex
														{
															position57 := position
															{
																position58, tokenIndex58 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l59
																}
																position++
																goto l58
															l59:
																position, tokenIndex = position58, tokenIndex58
																if buffer[position] != rune('L') {
																	goto l56
																}
																position++
															}
														l58:
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
																	goto l56
																}
																position++
															}
														l60:
															if !_rules[rulews]() {
																goto l56
															}
															if !_rules[ruleDst16]() {
																goto l56
															}
															if !_rules[rulesep]() {
																goto l56
															}
															if !_rules[ruleSrc16]() {
																goto l56
															}
															{
																add(ruleAction3, position)
															}
															add(ruleLoad16, position57)
														}
														goto l55
													l56:
														position, tokenIndex = position55, tokenIndex55
														{
															position63 := position
															{
																position64, tokenIndex64 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l65
																}
																position++
																goto l64
															l65:
																position, tokenIndex = position64, tokenIndex64
																if buffer[position] != rune('L') {
																	goto l25
																}
																position++
															}
														l64:
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
																	goto l25
																}
																position++
															}
														l66:
															if !_rules[rulews]() {
																goto l25
															}
															{
																position68 := position
																{
																	position69, tokenIndex69 := position, tokenIndex
																	if !_rules[ruleReg8]() {
																		goto l70
																	}
																	goto l69
																l70:
																	position, tokenIndex = position69, tokenIndex69
																	if !_rules[ruleReg16Contents]() {
																		goto l71
																	}
																	goto l69
																l71:
																	position, tokenIndex = position69, tokenIndex69
																	if !_rules[rulenn_contents]() {
																		goto l25
																	}
																}
															l69:
																{
																	add(ruleAction16, position)
																}
																add(ruleDst8, position68)
															}
															if !_rules[rulesep]() {
																goto l25
															}
															if !_rules[ruleSrc8]() {
																goto l25
															}
															{
																add(ruleAction2, position)
															}
															add(ruleLoad8, position63)
														}
													}
												l55:
													add(ruleLoad, position54)
												}
												break
											}
										}

									}
								l27:
									add(ruleAssignment, position26)
								}
								goto l24
							l25:
								position, tokenIndex = position24, tokenIndex24
								{
									position75 := position
									{
										position76, tokenIndex76 := position, tokenIndex
										{
											position78 := position
											{
												position79, tokenIndex79 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l80
												}
												position++
												goto l79
											l80:
												position, tokenIndex = position79, tokenIndex79
												if buffer[position] != rune('I') {
													goto l77
												}
												position++
											}
										l79:
											{
												position81, tokenIndex81 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l82
												}
												position++
												goto l81
											l82:
												position, tokenIndex = position81, tokenIndex81
												if buffer[position] != rune('N') {
													goto l77
												}
												position++
											}
										l81:
											{
												position83, tokenIndex83 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l84
												}
												position++
												goto l83
											l84:
												position, tokenIndex = position83, tokenIndex83
												if buffer[position] != rune('C') {
													goto l77
												}
												position++
											}
										l83:
											if !_rules[rulews]() {
												goto l77
											}
											if !_rules[ruleILoc8]() {
												goto l77
											}
											{
												add(ruleAction7, position)
											}
											add(ruleInc16Indexed8, position78)
										}
										goto l76
									l77:
										position, tokenIndex = position76, tokenIndex76
										{
											position87 := position
											{
												position88, tokenIndex88 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l89
												}
												position++
												goto l88
											l89:
												position, tokenIndex = position88, tokenIndex88
												if buffer[position] != rune('I') {
													goto l86
												}
												position++
											}
										l88:
											{
												position90, tokenIndex90 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l91
												}
												position++
												goto l90
											l91:
												position, tokenIndex = position90, tokenIndex90
												if buffer[position] != rune('N') {
													goto l86
												}
												position++
											}
										l90:
											{
												position92, tokenIndex92 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l93
												}
												position++
												goto l92
											l93:
												position, tokenIndex = position92, tokenIndex92
												if buffer[position] != rune('C') {
													goto l86
												}
												position++
											}
										l92:
											if !_rules[rulews]() {
												goto l86
											}
											if !_rules[ruleLoc16]() {
												goto l86
											}
											{
												add(ruleAction9, position)
											}
											add(ruleInc16, position87)
										}
										goto l76
									l86:
										position, tokenIndex = position76, tokenIndex76
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
													goto l74
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
													goto l74
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
													goto l74
												}
												position++
											}
										l100:
											if !_rules[rulews]() {
												goto l74
											}
											if !_rules[ruleLoc8]() {
												goto l74
											}
											{
												add(ruleAction8, position)
											}
											add(ruleInc8, position95)
										}
									}
								l76:
									add(ruleInc, position75)
								}
								goto l24
							l74:
								position, tokenIndex = position24, tokenIndex24
								{
									position104 := position
									{
										position105, tokenIndex105 := position, tokenIndex
										{
											position107 := position
											{
												position108, tokenIndex108 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l109
												}
												position++
												goto l108
											l109:
												position, tokenIndex = position108, tokenIndex108
												if buffer[position] != rune('D') {
													goto l106
												}
												position++
											}
										l108:
											{
												position110, tokenIndex110 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l111
												}
												position++
												goto l110
											l111:
												position, tokenIndex = position110, tokenIndex110
												if buffer[position] != rune('E') {
													goto l106
												}
												position++
											}
										l110:
											{
												position112, tokenIndex112 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l113
												}
												position++
												goto l112
											l113:
												position, tokenIndex = position112, tokenIndex112
												if buffer[position] != rune('C') {
													goto l106
												}
												position++
											}
										l112:
											if !_rules[rulews]() {
												goto l106
											}
											if !_rules[ruleILoc8]() {
												goto l106
											}
											{
												add(ruleAction10, position)
											}
											add(ruleDec16Indexed8, position107)
										}
										goto l105
									l106:
										position, tokenIndex = position105, tokenIndex105
										{
											position116 := position
											{
												position117, tokenIndex117 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l118
												}
												position++
												goto l117
											l118:
												position, tokenIndex = position117, tokenIndex117
												if buffer[position] != rune('D') {
													goto l115
												}
												position++
											}
										l117:
											{
												position119, tokenIndex119 := position, tokenIndex
												if buffer[position] != rune('e') {
													goto l120
												}
												position++
												goto l119
											l120:
												position, tokenIndex = position119, tokenIndex119
												if buffer[position] != rune('E') {
													goto l115
												}
												position++
											}
										l119:
											{
												position121, tokenIndex121 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l122
												}
												position++
												goto l121
											l122:
												position, tokenIndex = position121, tokenIndex121
												if buffer[position] != rune('C') {
													goto l115
												}
												position++
											}
										l121:
											if !_rules[rulews]() {
												goto l115
											}
											if !_rules[ruleLoc16]() {
												goto l115
											}
											{
												add(ruleAction12, position)
											}
											add(ruleDec16, position116)
										}
										goto l105
									l115:
										position, tokenIndex = position105, tokenIndex105
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
													goto l103
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
													goto l103
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
													goto l103
												}
												position++
											}
										l129:
											if !_rules[rulews]() {
												goto l103
											}
											if !_rules[ruleLoc8]() {
												goto l103
											}
											{
												add(ruleAction11, position)
											}
											add(ruleDec8, position124)
										}
									}
								l105:
									add(ruleDec, position104)
								}
								goto l24
							l103:
								position, tokenIndex = position24, tokenIndex24
								{
									position133 := position
									{
										position134, tokenIndex134 := position, tokenIndex
										{
											position136 := position
											{
												position137, tokenIndex137 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l138
												}
												position++
												goto l137
											l138:
												position, tokenIndex = position137, tokenIndex137
												if buffer[position] != rune('A') {
													goto l135
												}
												position++
											}
										l137:
											{
												position139, tokenIndex139 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l140
												}
												position++
												goto l139
											l140:
												position, tokenIndex = position139, tokenIndex139
												if buffer[position] != rune('D') {
													goto l135
												}
												position++
											}
										l139:
											{
												position141, tokenIndex141 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l142
												}
												position++
												goto l141
											l142:
												position, tokenIndex = position141, tokenIndex141
												if buffer[position] != rune('D') {
													goto l135
												}
												position++
											}
										l141:
											if !_rules[rulews]() {
												goto l135
											}
											if !_rules[ruleDst16]() {
												goto l135
											}
											if !_rules[rulesep]() {
												goto l135
											}
											if !_rules[ruleSrc16]() {
												goto l135
											}
											{
												add(ruleAction13, position)
											}
											add(ruleAdd16, position136)
										}
										goto l134
									l135:
										position, tokenIndex = position134, tokenIndex134
										{
											position145 := position
											{
												position146, tokenIndex146 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l147
												}
												position++
												goto l146
											l147:
												position, tokenIndex = position146, tokenIndex146
												if buffer[position] != rune('A') {
													goto l144
												}
												position++
											}
										l146:
											{
												position148, tokenIndex148 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l149
												}
												position++
												goto l148
											l149:
												position, tokenIndex = position148, tokenIndex148
												if buffer[position] != rune('D') {
													goto l144
												}
												position++
											}
										l148:
											{
												position150, tokenIndex150 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l151
												}
												position++
												goto l150
											l151:
												position, tokenIndex = position150, tokenIndex150
												if buffer[position] != rune('C') {
													goto l144
												}
												position++
											}
										l150:
											if !_rules[rulews]() {
												goto l144
											}
											if !_rules[ruleDst16]() {
												goto l144
											}
											if !_rules[rulesep]() {
												goto l144
											}
											if !_rules[ruleSrc16]() {
												goto l144
											}
											{
												add(ruleAction14, position)
											}
											add(ruleAdc16, position145)
										}
										goto l134
									l144:
										position, tokenIndex = position134, tokenIndex134
										{
											position153 := position
											{
												position154, tokenIndex154 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l155
												}
												position++
												goto l154
											l155:
												position, tokenIndex = position154, tokenIndex154
												if buffer[position] != rune('S') {
													goto l132
												}
												position++
											}
										l154:
											{
												position156, tokenIndex156 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l157
												}
												position++
												goto l156
											l157:
												position, tokenIndex = position156, tokenIndex156
												if buffer[position] != rune('B') {
													goto l132
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
													goto l132
												}
												position++
											}
										l158:
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
												add(ruleAction15, position)
											}
											add(ruleSbc16, position153)
										}
									}
								l134:
									add(ruleAlu16, position133)
								}
								goto l24
							l132:
								position, tokenIndex = position24, tokenIndex24
								{
									position162 := position
									{
										position163, tokenIndex163 := position, tokenIndex
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
												if buffer[position] != rune('d') {
													goto l171
												}
												position++
												goto l170
											l171:
												position, tokenIndex = position170, tokenIndex170
												if buffer[position] != rune('D') {
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
											add(ruleAdd, position165)
										}
										goto l163
									l164:
										position, tokenIndex = position163, tokenIndex163
										{
											position176 := position
											{
												position177, tokenIndex177 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l178
												}
												position++
												goto l177
											l178:
												position, tokenIndex = position177, tokenIndex177
												if buffer[position] != rune('A') {
													goto l175
												}
												position++
											}
										l177:
											{
												position179, tokenIndex179 := position, tokenIndex
												if buffer[position] != rune('d') {
													goto l180
												}
												position++
												goto l179
											l180:
												position, tokenIndex = position179, tokenIndex179
												if buffer[position] != rune('D') {
													goto l175
												}
												position++
											}
										l179:
											{
												position181, tokenIndex181 := position, tokenIndex
												if buffer[position] != rune('c') {
													goto l182
												}
												position++
												goto l181
											l182:
												position, tokenIndex = position181, tokenIndex181
												if buffer[position] != rune('C') {
													goto l175
												}
												position++
											}
										l181:
											if !_rules[rulews]() {
												goto l175
											}
											{
												position183, tokenIndex183 := position, tokenIndex
												if buffer[position] != rune('a') {
													goto l184
												}
												position++
												goto l183
											l184:
												position, tokenIndex = position183, tokenIndex183
												if buffer[position] != rune('A') {
													goto l175
												}
												position++
											}
										l183:
											if !_rules[rulesep]() {
												goto l175
											}
											if !_rules[ruleSrc8]() {
												goto l175
											}
											{
												add(ruleAction40, position)
											}
											add(ruleAdc, position176)
										}
										goto l163
									l175:
										position, tokenIndex = position163, tokenIndex163
										{
											position187 := position
											{
												position188, tokenIndex188 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l189
												}
												position++
												goto l188
											l189:
												position, tokenIndex = position188, tokenIndex188
												if buffer[position] != rune('S') {
													goto l186
												}
												position++
											}
										l188:
											{
												position190, tokenIndex190 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l191
												}
												position++
												goto l190
											l191:
												position, tokenIndex = position190, tokenIndex190
												if buffer[position] != rune('U') {
													goto l186
												}
												position++
											}
										l190:
											{
												position192, tokenIndex192 := position, tokenIndex
												if buffer[position] != rune('b') {
													goto l193
												}
												position++
												goto l192
											l193:
												position, tokenIndex = position192, tokenIndex192
												if buffer[position] != rune('B') {
													goto l186
												}
												position++
											}
										l192:
											if !_rules[rulews]() {
												goto l186
											}
											if !_rules[ruleSrc8]() {
												goto l186
											}
											{
												add(ruleAction41, position)
											}
											add(ruleSub, position187)
										}
										goto l163
									l186:
										position, tokenIndex = position163, tokenIndex163
										{
											switch buffer[position] {
											case 'C', 'c':
												{
													position196 := position
													{
														position197, tokenIndex197 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l198
														}
														position++
														goto l197
													l198:
														position, tokenIndex = position197, tokenIndex197
														if buffer[position] != rune('C') {
															goto l161
														}
														position++
													}
												l197:
													{
														position199, tokenIndex199 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l200
														}
														position++
														goto l199
													l200:
														position, tokenIndex = position199, tokenIndex199
														if buffer[position] != rune('P') {
															goto l161
														}
														position++
													}
												l199:
													if !_rules[rulews]() {
														goto l161
													}
													if !_rules[ruleSrc8]() {
														goto l161
													}
													{
														add(ruleAction46, position)
													}
													add(ruleCp, position196)
												}
												break
											case 'O', 'o':
												{
													position202 := position
													{
														position203, tokenIndex203 := position, tokenIndex
														if buffer[position] != rune('o') {
															goto l204
														}
														position++
														goto l203
													l204:
														position, tokenIndex = position203, tokenIndex203
														if buffer[position] != rune('O') {
															goto l161
														}
														position++
													}
												l203:
													{
														position205, tokenIndex205 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l206
														}
														position++
														goto l205
													l206:
														position, tokenIndex = position205, tokenIndex205
														if buffer[position] != rune('R') {
															goto l161
														}
														position++
													}
												l205:
													if !_rules[rulews]() {
														goto l161
													}
													if !_rules[ruleSrc8]() {
														goto l161
													}
													{
														add(ruleAction45, position)
													}
													add(ruleOr, position202)
												}
												break
											case 'X', 'x':
												{
													position208 := position
													{
														position209, tokenIndex209 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l210
														}
														position++
														goto l209
													l210:
														position, tokenIndex = position209, tokenIndex209
														if buffer[position] != rune('X') {
															goto l161
														}
														position++
													}
												l209:
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
															goto l161
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
															goto l161
														}
														position++
													}
												l213:
													if !_rules[rulews]() {
														goto l161
													}
													if !_rules[ruleSrc8]() {
														goto l161
													}
													{
														add(ruleAction44, position)
													}
													add(ruleXor, position208)
												}
												break
											case 'A', 'a':
												{
													position216 := position
													{
														position217, tokenIndex217 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l218
														}
														position++
														goto l217
													l218:
														position, tokenIndex = position217, tokenIndex217
														if buffer[position] != rune('A') {
															goto l161
														}
														position++
													}
												l217:
													{
														position219, tokenIndex219 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l220
														}
														position++
														goto l219
													l220:
														position, tokenIndex = position219, tokenIndex219
														if buffer[position] != rune('N') {
															goto l161
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
															goto l161
														}
														position++
													}
												l221:
													if !_rules[rulews]() {
														goto l161
													}
													if !_rules[ruleSrc8]() {
														goto l161
													}
													{
														add(ruleAction43, position)
													}
													add(ruleAnd, position216)
												}
												break
											default:
												{
													position224 := position
													{
														position225, tokenIndex225 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l226
														}
														position++
														goto l225
													l226:
														position, tokenIndex = position225, tokenIndex225
														if buffer[position] != rune('S') {
															goto l161
														}
														position++
													}
												l225:
													{
														position227, tokenIndex227 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l228
														}
														position++
														goto l227
													l228:
														position, tokenIndex = position227, tokenIndex227
														if buffer[position] != rune('B') {
															goto l161
														}
														position++
													}
												l227:
													{
														position229, tokenIndex229 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l230
														}
														position++
														goto l229
													l230:
														position, tokenIndex = position229, tokenIndex229
														if buffer[position] != rune('C') {
															goto l161
														}
														position++
													}
												l229:
													if !_rules[rulews]() {
														goto l161
													}
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
															goto l161
														}
														position++
													}
												l231:
													if !_rules[rulesep]() {
														goto l161
													}
													if !_rules[ruleSrc8]() {
														goto l161
													}
													{
														add(ruleAction42, position)
													}
													add(ruleSbc, position224)
												}
												break
											}
										}

									}
								l163:
									add(ruleAlu, position162)
								}
								goto l24
							l161:
								position, tokenIndex = position24, tokenIndex24
								{
									position235 := position
									{
										position236, tokenIndex236 := position, tokenIndex
										{
											position238 := position
											{
												position239, tokenIndex239 := position, tokenIndex
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
														if buffer[position] != rune('l') {
															goto l245
														}
														position++
														goto l244
													l245:
														position, tokenIndex = position244, tokenIndex244
														if buffer[position] != rune('L') {
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
													add(ruleRlc, position241)
												}
												goto l239
											l240:
												position, tokenIndex = position239, tokenIndex239
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
														if buffer[position] != rune('r') {
															goto l256
														}
														position++
														goto l255
													l256:
														position, tokenIndex = position255, tokenIndex255
														if buffer[position] != rune('R') {
															goto l251
														}
														position++
													}
												l255:
													{
														position257, tokenIndex257 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l258
														}
														position++
														goto l257
													l258:
														position, tokenIndex = position257, tokenIndex257
														if buffer[position] != rune('C') {
															goto l251
														}
														position++
													}
												l257:
													if !_rules[rulews]() {
														goto l251
													}
													if !_rules[ruleLoc8]() {
														goto l251
													}
													{
														position259, tokenIndex259 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l259
														}
														if !_rules[ruleCopy8]() {
															goto l259
														}
														goto l260
													l259:
														position, tokenIndex = position259, tokenIndex259
													}
												l260:
													{
														add(ruleAction48, position)
													}
													add(ruleRrc, position252)
												}
												goto l239
											l251:
												position, tokenIndex = position239, tokenIndex239
												{
													position263 := position
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
															goto l262
														}
														position++
													}
												l264:
													{
														position266, tokenIndex266 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l267
														}
														position++
														goto l266
													l267:
														position, tokenIndex = position266, tokenIndex266
														if buffer[position] != rune('L') {
															goto l262
														}
														position++
													}
												l266:
													if !_rules[rulews]() {
														goto l262
													}
													if !_rules[ruleLoc8]() {
														goto l262
													}
													{
														position268, tokenIndex268 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l268
														}
														if !_rules[ruleCopy8]() {
															goto l268
														}
														goto l269
													l268:
														position, tokenIndex = position268, tokenIndex268
													}
												l269:
													{
														add(ruleAction49, position)
													}
													add(ruleRl, position263)
												}
												goto l239
											l262:
												position, tokenIndex = position239, tokenIndex239
												{
													position272 := position
													{
														position273, tokenIndex273 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l274
														}
														position++
														goto l273
													l274:
														position, tokenIndex = position273, tokenIndex273
														if buffer[position] != rune('R') {
															goto l271
														}
														position++
													}
												l273:
													{
														position275, tokenIndex275 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l276
														}
														position++
														goto l275
													l276:
														position, tokenIndex = position275, tokenIndex275
														if buffer[position] != rune('R') {
															goto l271
														}
														position++
													}
												l275:
													if !_rules[rulews]() {
														goto l271
													}
													if !_rules[ruleLoc8]() {
														goto l271
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
													add(ruleRr, position272)
												}
												goto l239
											l271:
												position, tokenIndex = position239, tokenIndex239
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
														if buffer[position] != rune('l') {
															goto l285
														}
														position++
														goto l284
													l285:
														position, tokenIndex = position284, tokenIndex284
														if buffer[position] != rune('L') {
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
													add(ruleSla, position281)
												}
												goto l239
											l280:
												position, tokenIndex = position239, tokenIndex239
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
														if buffer[position] != rune('r') {
															goto l296
														}
														position++
														goto l295
													l296:
														position, tokenIndex = position295, tokenIndex295
														if buffer[position] != rune('R') {
															goto l291
														}
														position++
													}
												l295:
													{
														position297, tokenIndex297 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l298
														}
														position++
														goto l297
													l298:
														position, tokenIndex = position297, tokenIndex297
														if buffer[position] != rune('A') {
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
													add(ruleSra, position292)
												}
												goto l239
											l291:
												position, tokenIndex = position239, tokenIndex239
												{
													position303 := position
													{
														position304, tokenIndex304 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l305
														}
														position++
														goto l304
													l305:
														position, tokenIndex = position304, tokenIndex304
														if buffer[position] != rune('S') {
															goto l302
														}
														position++
													}
												l304:
													{
														position306, tokenIndex306 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l307
														}
														position++
														goto l306
													l307:
														position, tokenIndex = position306, tokenIndex306
														if buffer[position] != rune('L') {
															goto l302
														}
														position++
													}
												l306:
													{
														position308, tokenIndex308 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l309
														}
														position++
														goto l308
													l309:
														position, tokenIndex = position308, tokenIndex308
														if buffer[position] != rune('L') {
															goto l302
														}
														position++
													}
												l308:
													if !_rules[rulews]() {
														goto l302
													}
													if !_rules[ruleLoc8]() {
														goto l302
													}
													{
														position310, tokenIndex310 := position, tokenIndex
														if !_rules[rulesep]() {
															goto l310
														}
														if !_rules[ruleCopy8]() {
															goto l310
														}
														goto l311
													l310:
														position, tokenIndex = position310, tokenIndex310
													}
												l311:
													{
														add(ruleAction53, position)
													}
													add(ruleSll, position303)
												}
												goto l239
											l302:
												position, tokenIndex = position239, tokenIndex239
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
															goto l237
														}
														position++
													}
												l314:
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
															goto l237
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
															goto l237
														}
														position++
													}
												l318:
													if !_rules[rulews]() {
														goto l237
													}
													if !_rules[ruleLoc8]() {
														goto l237
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
														add(ruleAction54, position)
													}
													add(ruleSrl, position313)
												}
											}
										l239:
											add(ruleRot, position238)
										}
										goto l236
									l237:
										position, tokenIndex = position236, tokenIndex236
										{
											switch buffer[position] {
											case 'S', 's':
												{
													position324 := position
													{
														position325, tokenIndex325 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l326
														}
														position++
														goto l325
													l326:
														position, tokenIndex = position325, tokenIndex325
														if buffer[position] != rune('S') {
															goto l234
														}
														position++
													}
												l325:
													{
														position327, tokenIndex327 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l328
														}
														position++
														goto l327
													l328:
														position, tokenIndex = position327, tokenIndex327
														if buffer[position] != rune('E') {
															goto l234
														}
														position++
													}
												l327:
													{
														position329, tokenIndex329 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l330
														}
														position++
														goto l329
													l330:
														position, tokenIndex = position329, tokenIndex329
														if buffer[position] != rune('T') {
															goto l234
														}
														position++
													}
												l329:
													if !_rules[rulews]() {
														goto l234
													}
													if !_rules[ruleoctaldigit]() {
														goto l234
													}
													if !_rules[rulesep]() {
														goto l234
													}
													if !_rules[ruleLoc8]() {
														goto l234
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
														add(ruleAction57, position)
													}
													add(ruleSet, position324)
												}
												break
											case 'R', 'r':
												{
													position334 := position
													{
														position335, tokenIndex335 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l336
														}
														position++
														goto l335
													l336:
														position, tokenIndex = position335, tokenIndex335
														if buffer[position] != rune('R') {
															goto l234
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
															goto l234
														}
														position++
													}
												l337:
													{
														position339, tokenIndex339 := position, tokenIndex
														if buffer[position] != rune('s') {
															goto l340
														}
														position++
														goto l339
													l340:
														position, tokenIndex = position339, tokenIndex339
														if buffer[position] != rune('S') {
															goto l234
														}
														position++
													}
												l339:
													if !_rules[rulews]() {
														goto l234
													}
													if !_rules[ruleoctaldigit]() {
														goto l234
													}
													if !_rules[rulesep]() {
														goto l234
													}
													if !_rules[ruleLoc8]() {
														goto l234
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
														add(ruleAction56, position)
													}
													add(ruleRes, position334)
												}
												break
											default:
												{
													position344 := position
													{
														position345, tokenIndex345 := position, tokenIndex
														if buffer[position] != rune('b') {
															goto l346
														}
														position++
														goto l345
													l346:
														position, tokenIndex = position345, tokenIndex345
														if buffer[position] != rune('B') {
															goto l234
														}
														position++
													}
												l345:
													{
														position347, tokenIndex347 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l348
														}
														position++
														goto l347
													l348:
														position, tokenIndex = position347, tokenIndex347
														if buffer[position] != rune('I') {
															goto l234
														}
														position++
													}
												l347:
													{
														position349, tokenIndex349 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l350
														}
														position++
														goto l349
													l350:
														position, tokenIndex = position349, tokenIndex349
														if buffer[position] != rune('T') {
															goto l234
														}
														position++
													}
												l349:
													if !_rules[rulews]() {
														goto l234
													}
													if !_rules[ruleoctaldigit]() {
														goto l234
													}
													if !_rules[rulesep]() {
														goto l234
													}
													if !_rules[ruleLoc8]() {
														goto l234
													}
													{
														add(ruleAction55, position)
													}
													add(ruleBit, position344)
												}
												break
											}
										}

									}
								l236:
									add(ruleBitOp, position235)
								}
								goto l24
							l234:
								position, tokenIndex = position24, tokenIndex24
								{
									position353 := position
									{
										position354, tokenIndex354 := position, tokenIndex
										{
											position356 := position
											{
												position357 := position
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
														goto l355
													}
													position++
												}
											l358:
												{
													position360, tokenIndex360 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l361
													}
													position++
													goto l360
												l361:
													position, tokenIndex = position360, tokenIndex360
													if buffer[position] != rune('E') {
														goto l355
													}
													position++
												}
											l360:
												{
													position362, tokenIndex362 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l363
													}
													position++
													goto l362
												l363:
													position, tokenIndex = position362, tokenIndex362
													if buffer[position] != rune('T') {
														goto l355
													}
													position++
												}
											l362:
												{
													position364, tokenIndex364 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l365
													}
													position++
													goto l364
												l365:
													position, tokenIndex = position364, tokenIndex364
													if buffer[position] != rune('N') {
														goto l355
													}
													position++
												}
											l364:
												add(rulePegText, position357)
											}
											{
												add(ruleAction72, position)
											}
											add(ruleRetn, position356)
										}
										goto l354
									l355:
										position, tokenIndex = position354, tokenIndex354
										{
											position368 := position
											{
												position369 := position
												{
													position370, tokenIndex370 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l371
													}
													position++
													goto l370
												l371:
													position, tokenIndex = position370, tokenIndex370
													if buffer[position] != rune('R') {
														goto l367
													}
													position++
												}
											l370:
												{
													position372, tokenIndex372 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l373
													}
													position++
													goto l372
												l373:
													position, tokenIndex = position372, tokenIndex372
													if buffer[position] != rune('E') {
														goto l367
													}
													position++
												}
											l372:
												{
													position374, tokenIndex374 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l375
													}
													position++
													goto l374
												l375:
													position, tokenIndex = position374, tokenIndex374
													if buffer[position] != rune('T') {
														goto l367
													}
													position++
												}
											l374:
												{
													position376, tokenIndex376 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l377
													}
													position++
													goto l376
												l377:
													position, tokenIndex = position376, tokenIndex376
													if buffer[position] != rune('I') {
														goto l367
													}
													position++
												}
											l376:
												add(rulePegText, position369)
											}
											{
												add(ruleAction73, position)
											}
											add(ruleReti, position368)
										}
										goto l354
									l367:
										position, tokenIndex = position354, tokenIndex354
										{
											position380 := position
											{
												position381 := position
												{
													position382, tokenIndex382 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l383
													}
													position++
													goto l382
												l383:
													position, tokenIndex = position382, tokenIndex382
													if buffer[position] != rune('R') {
														goto l379
													}
													position++
												}
											l382:
												{
													position384, tokenIndex384 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l385
													}
													position++
													goto l384
												l385:
													position, tokenIndex = position384, tokenIndex384
													if buffer[position] != rune('R') {
														goto l379
													}
													position++
												}
											l384:
												{
													position386, tokenIndex386 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l387
													}
													position++
													goto l386
												l387:
													position, tokenIndex = position386, tokenIndex386
													if buffer[position] != rune('D') {
														goto l379
													}
													position++
												}
											l386:
												add(rulePegText, position381)
											}
											{
												add(ruleAction74, position)
											}
											add(ruleRrd, position380)
										}
										goto l354
									l379:
										position, tokenIndex = position354, tokenIndex354
										{
											position390 := position
											{
												position391 := position
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
														goto l389
													}
													position++
												}
											l392:
												{
													position394, tokenIndex394 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l395
													}
													position++
													goto l394
												l395:
													position, tokenIndex = position394, tokenIndex394
													if buffer[position] != rune('M') {
														goto l389
													}
													position++
												}
											l394:
												if buffer[position] != rune(' ') {
													goto l389
												}
												position++
												if buffer[position] != rune('0') {
													goto l389
												}
												position++
												add(rulePegText, position391)
											}
											{
												add(ruleAction76, position)
											}
											add(ruleIm0, position390)
										}
										goto l354
									l389:
										position, tokenIndex = position354, tokenIndex354
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
												if buffer[position] != rune('1') {
													goto l397
												}
												position++
												add(rulePegText, position399)
											}
											{
												add(ruleAction77, position)
											}
											add(ruleIm1, position398)
										}
										goto l354
									l397:
										position, tokenIndex = position354, tokenIndex354
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
												if buffer[position] != rune('2') {
													goto l405
												}
												position++
												add(rulePegText, position407)
											}
											{
												add(ruleAction78, position)
											}
											add(ruleIm2, position406)
										}
										goto l354
									l405:
										position, tokenIndex = position354, tokenIndex354
										{
											switch buffer[position] {
											case 'I', 'O', 'i', 'o':
												{
													position414 := position
													{
														position415, tokenIndex415 := position, tokenIndex
														{
															position417 := position
															{
																position418 := position
																{
																	position419, tokenIndex419 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l420
																	}
																	position++
																	goto l419
																l420:
																	position, tokenIndex = position419, tokenIndex419
																	if buffer[position] != rune('I') {
																		goto l416
																	}
																	position++
																}
															l419:
																{
																	position421, tokenIndex421 := position, tokenIndex
																	if buffer[position] != rune('n') {
																		goto l422
																	}
																	position++
																	goto l421
																l422:
																	position, tokenIndex = position421, tokenIndex421
																	if buffer[position] != rune('N') {
																		goto l416
																	}
																	position++
																}
															l421:
																{
																	position423, tokenIndex423 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l424
																	}
																	position++
																	goto l423
																l424:
																	position, tokenIndex = position423, tokenIndex423
																	if buffer[position] != rune('I') {
																		goto l416
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
																		goto l416
																	}
																	position++
																}
															l425:
																add(rulePegText, position418)
															}
															{
																add(ruleAction89, position)
															}
															add(ruleInir, position417)
														}
														goto l415
													l416:
														position, tokenIndex = position415, tokenIndex415
														{
															position429 := position
															{
																position430 := position
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
																		goto l428
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
																		goto l428
																	}
																	position++
																}
															l433:
																{
																	position435, tokenIndex435 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l436
																	}
																	position++
																	goto l435
																l436:
																	position, tokenIndex = position435, tokenIndex435
																	if buffer[position] != rune('I') {
																		goto l428
																	}
																	position++
																}
															l435:
																add(rulePegText, position430)
															}
															{
																add(ruleAction81, position)
															}
															add(ruleIni, position429)
														}
														goto l415
													l428:
														position, tokenIndex = position415, tokenIndex415
														{
															position439 := position
															{
																position440 := position
																{
																	position441, tokenIndex441 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l442
																	}
																	position++
																	goto l441
																l442:
																	position, tokenIndex = position441, tokenIndex441
																	if buffer[position] != rune('O') {
																		goto l438
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
																		goto l438
																	}
																	position++
																}
															l447:
																add(rulePegText, position440)
															}
															{
																add(ruleAction90, position)
															}
															add(ruleOtir, position439)
														}
														goto l415
													l438:
														position, tokenIndex = position415, tokenIndex415
														{
															position451 := position
															{
																position452 := position
																{
																	position453, tokenIndex453 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l454
																	}
																	position++
																	goto l453
																l454:
																	position, tokenIndex = position453, tokenIndex453
																	if buffer[position] != rune('O') {
																		goto l450
																	}
																	position++
																}
															l453:
																{
																	position455, tokenIndex455 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l456
																	}
																	position++
																	goto l455
																l456:
																	position, tokenIndex = position455, tokenIndex455
																	if buffer[position] != rune('U') {
																		goto l450
																	}
																	position++
																}
															l455:
																{
																	position457, tokenIndex457 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l458
																	}
																	position++
																	goto l457
																l458:
																	position, tokenIndex = position457, tokenIndex457
																	if buffer[position] != rune('T') {
																		goto l450
																	}
																	position++
																}
															l457:
																{
																	position459, tokenIndex459 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l460
																	}
																	position++
																	goto l459
																l460:
																	position, tokenIndex = position459, tokenIndex459
																	if buffer[position] != rune('I') {
																		goto l450
																	}
																	position++
																}
															l459:
																add(rulePegText, position452)
															}
															{
																add(ruleAction82, position)
															}
															add(ruleOuti, position451)
														}
														goto l415
													l450:
														position, tokenIndex = position415, tokenIndex415
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
																	if buffer[position] != rune('n') {
																		goto l468
																	}
																	position++
																	goto l467
																l468:
																	position, tokenIndex = position467, tokenIndex467
																	if buffer[position] != rune('N') {
																		goto l462
																	}
																	position++
																}
															l467:
																{
																	position469, tokenIndex469 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l470
																	}
																	position++
																	goto l469
																l470:
																	position, tokenIndex = position469, tokenIndex469
																	if buffer[position] != rune('D') {
																		goto l462
																	}
																	position++
																}
															l469:
																{
																	position471, tokenIndex471 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l472
																	}
																	position++
																	goto l471
																l472:
																	position, tokenIndex = position471, tokenIndex471
																	if buffer[position] != rune('R') {
																		goto l462
																	}
																	position++
																}
															l471:
																add(rulePegText, position464)
															}
															{
																add(ruleAction93, position)
															}
															add(ruleIndr, position463)
														}
														goto l415
													l462:
														position, tokenIndex = position415, tokenIndex415
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
																	if buffer[position] != rune('n') {
																		goto l480
																	}
																	position++
																	goto l479
																l480:
																	position, tokenIndex = position479, tokenIndex479
																	if buffer[position] != rune('N') {
																		goto l474
																	}
																	position++
																}
															l479:
																{
																	position481, tokenIndex481 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l482
																	}
																	position++
																	goto l481
																l482:
																	position, tokenIndex = position481, tokenIndex481
																	if buffer[position] != rune('D') {
																		goto l474
																	}
																	position++
																}
															l481:
																add(rulePegText, position476)
															}
															{
																add(ruleAction85, position)
															}
															add(ruleInd, position475)
														}
														goto l415
													l474:
														position, tokenIndex = position415, tokenIndex415
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
																		goto l484
																	}
																	position++
																}
															l487:
																{
																	position489, tokenIndex489 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l490
																	}
																	position++
																	goto l489
																l490:
																	position, tokenIndex = position489, tokenIndex489
																	if buffer[position] != rune('T') {
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
																{
																	position493, tokenIndex493 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l494
																	}
																	position++
																	goto l493
																l494:
																	position, tokenIndex = position493, tokenIndex493
																	if buffer[position] != rune('R') {
																		goto l484
																	}
																	position++
																}
															l493:
																add(rulePegText, position486)
															}
															{
																add(ruleAction94, position)
															}
															add(ruleOtdr, position485)
														}
														goto l415
													l484:
														position, tokenIndex = position415, tokenIndex415
														{
															position496 := position
															{
																position497 := position
																{
																	position498, tokenIndex498 := position, tokenIndex
																	if buffer[position] != rune('o') {
																		goto l499
																	}
																	position++
																	goto l498
																l499:
																	position, tokenIndex = position498, tokenIndex498
																	if buffer[position] != rune('O') {
																		goto l352
																	}
																	position++
																}
															l498:
																{
																	position500, tokenIndex500 := position, tokenIndex
																	if buffer[position] != rune('u') {
																		goto l501
																	}
																	position++
																	goto l500
																l501:
																	position, tokenIndex = position500, tokenIndex500
																	if buffer[position] != rune('U') {
																		goto l352
																	}
																	position++
																}
															l500:
																{
																	position502, tokenIndex502 := position, tokenIndex
																	if buffer[position] != rune('t') {
																		goto l503
																	}
																	position++
																	goto l502
																l503:
																	position, tokenIndex = position502, tokenIndex502
																	if buffer[position] != rune('T') {
																		goto l352
																	}
																	position++
																}
															l502:
																{
																	position504, tokenIndex504 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l505
																	}
																	position++
																	goto l504
																l505:
																	position, tokenIndex = position504, tokenIndex504
																	if buffer[position] != rune('D') {
																		goto l352
																	}
																	position++
																}
															l504:
																add(rulePegText, position497)
															}
															{
																add(ruleAction86, position)
															}
															add(ruleOutd, position496)
														}
													}
												l415:
													add(ruleBlitIO, position414)
												}
												break
											case 'R', 'r':
												{
													position507 := position
													{
														position508 := position
														{
															position509, tokenIndex509 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l510
															}
															position++
															goto l509
														l510:
															position, tokenIndex = position509, tokenIndex509
															if buffer[position] != rune('R') {
																goto l352
															}
															position++
														}
													l509:
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
																goto l352
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
																goto l352
															}
															position++
														}
													l513:
														add(rulePegText, position508)
													}
													{
														add(ruleAction75, position)
													}
													add(ruleRld, position507)
												}
												break
											case 'N', 'n':
												{
													position516 := position
													{
														position517 := position
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
																goto l352
															}
															position++
														}
													l518:
														{
															position520, tokenIndex520 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l521
															}
															position++
															goto l520
														l521:
															position, tokenIndex = position520, tokenIndex520
															if buffer[position] != rune('E') {
																goto l352
															}
															position++
														}
													l520:
														{
															position522, tokenIndex522 := position, tokenIndex
															if buffer[position] != rune('g') {
																goto l523
															}
															position++
															goto l522
														l523:
															position, tokenIndex = position522, tokenIndex522
															if buffer[position] != rune('G') {
																goto l352
															}
															position++
														}
													l522:
														add(rulePegText, position517)
													}
													{
														add(ruleAction71, position)
													}
													add(ruleNeg, position516)
												}
												break
											default:
												{
													position525 := position
													{
														position526, tokenIndex526 := position, tokenIndex
														{
															position528 := position
															{
																position529 := position
																{
																	position530, tokenIndex530 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l531
																	}
																	position++
																	goto l530
																l531:
																	position, tokenIndex = position530, tokenIndex530
																	if buffer[position] != rune('L') {
																		goto l527
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
																		goto l527
																	}
																	position++
																}
															l532:
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
																add(ruleAction87, position)
															}
															add(ruleLdir, position528)
														}
														goto l526
													l527:
														position, tokenIndex = position526, tokenIndex526
														{
															position540 := position
															{
																position541 := position
																{
																	position542, tokenIndex542 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l543
																	}
																	position++
																	goto l542
																l543:
																	position, tokenIndex = position542, tokenIndex542
																	if buffer[position] != rune('L') {
																		goto l539
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
																		goto l539
																	}
																	position++
																}
															l544:
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
																		goto l539
																	}
																	position++
																}
															l546:
																add(rulePegText, position541)
															}
															{
																add(ruleAction79, position)
															}
															add(ruleLdi, position540)
														}
														goto l526
													l539:
														position, tokenIndex = position526, tokenIndex526
														{
															position550 := position
															{
																position551 := position
																{
																	position552, tokenIndex552 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l553
																	}
																	position++
																	goto l552
																l553:
																	position, tokenIndex = position552, tokenIndex552
																	if buffer[position] != rune('C') {
																		goto l549
																	}
																	position++
																}
															l552:
																{
																	position554, tokenIndex554 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l555
																	}
																	position++
																	goto l554
																l555:
																	position, tokenIndex = position554, tokenIndex554
																	if buffer[position] != rune('P') {
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
																add(ruleAction88, position)
															}
															add(ruleCpir, position550)
														}
														goto l526
													l549:
														position, tokenIndex = position526, tokenIndex526
														{
															position562 := position
															{
																position563 := position
																{
																	position564, tokenIndex564 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l565
																	}
																	position++
																	goto l564
																l565:
																	position, tokenIndex = position564, tokenIndex564
																	if buffer[position] != rune('C') {
																		goto l561
																	}
																	position++
																}
															l564:
																{
																	position566, tokenIndex566 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l567
																	}
																	position++
																	goto l566
																l567:
																	position, tokenIndex = position566, tokenIndex566
																	if buffer[position] != rune('P') {
																		goto l561
																	}
																	position++
																}
															l566:
																{
																	position568, tokenIndex568 := position, tokenIndex
																	if buffer[position] != rune('i') {
																		goto l569
																	}
																	position++
																	goto l568
																l569:
																	position, tokenIndex = position568, tokenIndex568
																	if buffer[position] != rune('I') {
																		goto l561
																	}
																	position++
																}
															l568:
																add(rulePegText, position563)
															}
															{
																add(ruleAction80, position)
															}
															add(ruleCpi, position562)
														}
														goto l526
													l561:
														position, tokenIndex = position526, tokenIndex526
														{
															position572 := position
															{
																position573 := position
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
																		goto l571
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
																		goto l571
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
																		goto l571
																	}
																	position++
																}
															l578:
																{
																	position580, tokenIndex580 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l581
																	}
																	position++
																	goto l580
																l581:
																	position, tokenIndex = position580, tokenIndex580
																	if buffer[position] != rune('R') {
																		goto l571
																	}
																	position++
																}
															l580:
																add(rulePegText, position573)
															}
															{
																add(ruleAction91, position)
															}
															add(ruleLddr, position572)
														}
														goto l526
													l571:
														position, tokenIndex = position526, tokenIndex526
														{
															position584 := position
															{
																position585 := position
																{
																	position586, tokenIndex586 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l587
																	}
																	position++
																	goto l586
																l587:
																	position, tokenIndex = position586, tokenIndex586
																	if buffer[position] != rune('L') {
																		goto l583
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
																		goto l583
																	}
																	position++
																}
															l588:
																{
																	position590, tokenIndex590 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l591
																	}
																	position++
																	goto l590
																l591:
																	position, tokenIndex = position590, tokenIndex590
																	if buffer[position] != rune('D') {
																		goto l583
																	}
																	position++
																}
															l590:
																add(rulePegText, position585)
															}
															{
																add(ruleAction83, position)
															}
															add(ruleLdd, position584)
														}
														goto l526
													l583:
														position, tokenIndex = position526, tokenIndex526
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
																		goto l593
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
																{
																	position602, tokenIndex602 := position, tokenIndex
																	if buffer[position] != rune('r') {
																		goto l603
																	}
																	position++
																	goto l602
																l603:
																	position, tokenIndex = position602, tokenIndex602
																	if buffer[position] != rune('R') {
																		goto l593
																	}
																	position++
																}
															l602:
																add(rulePegText, position595)
															}
															{
																add(ruleAction92, position)
															}
															add(ruleCpdr, position594)
														}
														goto l526
													l593:
														position, tokenIndex = position526, tokenIndex526
														{
															position605 := position
															{
																position606 := position
																{
																	position607, tokenIndex607 := position, tokenIndex
																	if buffer[position] != rune('c') {
																		goto l608
																	}
																	position++
																	goto l607
																l608:
																	position, tokenIndex = position607, tokenIndex607
																	if buffer[position] != rune('C') {
																		goto l352
																	}
																	position++
																}
															l607:
																{
																	position609, tokenIndex609 := position, tokenIndex
																	if buffer[position] != rune('p') {
																		goto l610
																	}
																	position++
																	goto l609
																l610:
																	position, tokenIndex = position609, tokenIndex609
																	if buffer[position] != rune('P') {
																		goto l352
																	}
																	position++
																}
															l609:
																{
																	position611, tokenIndex611 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l612
																	}
																	position++
																	goto l611
																l612:
																	position, tokenIndex = position611, tokenIndex611
																	if buffer[position] != rune('D') {
																		goto l352
																	}
																	position++
																}
															l611:
																add(rulePegText, position606)
															}
															{
																add(ruleAction84, position)
															}
															add(ruleCpd, position605)
														}
													}
												l526:
													add(ruleBlit, position525)
												}
												break
											}
										}

									}
								l354:
									add(ruleEDSimple, position353)
								}
								goto l24
							l352:
								position, tokenIndex = position24, tokenIndex24
								{
									position615 := position
									{
										position616, tokenIndex616 := position, tokenIndex
										{
											position618 := position
											{
												position619 := position
												{
													position620, tokenIndex620 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l621
													}
													position++
													goto l620
												l621:
													position, tokenIndex = position620, tokenIndex620
													if buffer[position] != rune('R') {
														goto l617
													}
													position++
												}
											l620:
												{
													position622, tokenIndex622 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l623
													}
													position++
													goto l622
												l623:
													position, tokenIndex = position622, tokenIndex622
													if buffer[position] != rune('L') {
														goto l617
													}
													position++
												}
											l622:
												{
													position624, tokenIndex624 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l625
													}
													position++
													goto l624
												l625:
													position, tokenIndex = position624, tokenIndex624
													if buffer[position] != rune('C') {
														goto l617
													}
													position++
												}
											l624:
												{
													position626, tokenIndex626 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l627
													}
													position++
													goto l626
												l627:
													position, tokenIndex = position626, tokenIndex626
													if buffer[position] != rune('A') {
														goto l617
													}
													position++
												}
											l626:
												add(rulePegText, position619)
											}
											{
												add(ruleAction60, position)
											}
											add(ruleRlca, position618)
										}
										goto l616
									l617:
										position, tokenIndex = position616, tokenIndex616
										{
											position630 := position
											{
												position631 := position
												{
													position632, tokenIndex632 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l633
													}
													position++
													goto l632
												l633:
													position, tokenIndex = position632, tokenIndex632
													if buffer[position] != rune('R') {
														goto l629
													}
													position++
												}
											l632:
												{
													position634, tokenIndex634 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l635
													}
													position++
													goto l634
												l635:
													position, tokenIndex = position634, tokenIndex634
													if buffer[position] != rune('R') {
														goto l629
													}
													position++
												}
											l634:
												{
													position636, tokenIndex636 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l637
													}
													position++
													goto l636
												l637:
													position, tokenIndex = position636, tokenIndex636
													if buffer[position] != rune('C') {
														goto l629
													}
													position++
												}
											l636:
												{
													position638, tokenIndex638 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l639
													}
													position++
													goto l638
												l639:
													position, tokenIndex = position638, tokenIndex638
													if buffer[position] != rune('A') {
														goto l629
													}
													position++
												}
											l638:
												add(rulePegText, position631)
											}
											{
												add(ruleAction61, position)
											}
											add(ruleRrca, position630)
										}
										goto l616
									l629:
										position, tokenIndex = position616, tokenIndex616
										{
											position642 := position
											{
												position643 := position
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
														goto l641
													}
													position++
												}
											l644:
												{
													position646, tokenIndex646 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l647
													}
													position++
													goto l646
												l647:
													position, tokenIndex = position646, tokenIndex646
													if buffer[position] != rune('L') {
														goto l641
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
														goto l641
													}
													position++
												}
											l648:
												add(rulePegText, position643)
											}
											{
												add(ruleAction62, position)
											}
											add(ruleRla, position642)
										}
										goto l616
									l641:
										position, tokenIndex = position616, tokenIndex616
										{
											position652 := position
											{
												position653 := position
												{
													position654, tokenIndex654 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l655
													}
													position++
													goto l654
												l655:
													position, tokenIndex = position654, tokenIndex654
													if buffer[position] != rune('D') {
														goto l651
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
											add(ruleDaa, position652)
										}
										goto l616
									l651:
										position, tokenIndex = position616, tokenIndex616
										{
											position662 := position
											{
												position663 := position
												{
													position664, tokenIndex664 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l665
													}
													position++
													goto l664
												l665:
													position, tokenIndex = position664, tokenIndex664
													if buffer[position] != rune('C') {
														goto l661
													}
													position++
												}
											l664:
												{
													position666, tokenIndex666 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l667
													}
													position++
													goto l666
												l667:
													position, tokenIndex = position666, tokenIndex666
													if buffer[position] != rune('P') {
														goto l661
													}
													position++
												}
											l666:
												{
													position668, tokenIndex668 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l669
													}
													position++
													goto l668
												l669:
													position, tokenIndex = position668, tokenIndex668
													if buffer[position] != rune('L') {
														goto l661
													}
													position++
												}
											l668:
												add(rulePegText, position663)
											}
											{
												add(ruleAction65, position)
											}
											add(ruleCpl, position662)
										}
										goto l616
									l661:
										position, tokenIndex = position616, tokenIndex616
										{
											position672 := position
											{
												position673 := position
												{
													position674, tokenIndex674 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l675
													}
													position++
													goto l674
												l675:
													position, tokenIndex = position674, tokenIndex674
													if buffer[position] != rune('E') {
														goto l671
													}
													position++
												}
											l674:
												{
													position676, tokenIndex676 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l677
													}
													position++
													goto l676
												l677:
													position, tokenIndex = position676, tokenIndex676
													if buffer[position] != rune('X') {
														goto l671
													}
													position++
												}
											l676:
												{
													position678, tokenIndex678 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l679
													}
													position++
													goto l678
												l679:
													position, tokenIndex = position678, tokenIndex678
													if buffer[position] != rune('X') {
														goto l671
													}
													position++
												}
											l678:
												add(rulePegText, position673)
											}
											{
												add(ruleAction68, position)
											}
											add(ruleExx, position672)
										}
										goto l616
									l671:
										position, tokenIndex = position616, tokenIndex616
										{
											switch buffer[position] {
											case 'E', 'e':
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
																goto l614
															}
															position++
														}
													l684:
														{
															position686, tokenIndex686 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l687
															}
															position++
															goto l686
														l687:
															position, tokenIndex = position686, tokenIndex686
															if buffer[position] != rune('I') {
																goto l614
															}
															position++
														}
													l686:
														add(rulePegText, position683)
													}
													{
														add(ruleAction70, position)
													}
													add(ruleEi, position682)
												}
												break
											case 'D', 'd':
												{
													position689 := position
													{
														position690 := position
														{
															position691, tokenIndex691 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l692
															}
															position++
															goto l691
														l692:
															position, tokenIndex = position691, tokenIndex691
															if buffer[position] != rune('D') {
																goto l614
															}
															position++
														}
													l691:
														{
															position693, tokenIndex693 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l694
															}
															position++
															goto l693
														l694:
															position, tokenIndex = position693, tokenIndex693
															if buffer[position] != rune('I') {
																goto l614
															}
															position++
														}
													l693:
														add(rulePegText, position690)
													}
													{
														add(ruleAction69, position)
													}
													add(ruleDi, position689)
												}
												break
											case 'C', 'c':
												{
													position696 := position
													{
														position697 := position
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
																goto l614
															}
															position++
														}
													l698:
														{
															position700, tokenIndex700 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l701
															}
															position++
															goto l700
														l701:
															position, tokenIndex = position700, tokenIndex700
															if buffer[position] != rune('C') {
																goto l614
															}
															position++
														}
													l700:
														{
															position702, tokenIndex702 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l703
															}
															position++
															goto l702
														l703:
															position, tokenIndex = position702, tokenIndex702
															if buffer[position] != rune('F') {
																goto l614
															}
															position++
														}
													l702:
														add(rulePegText, position697)
													}
													{
														add(ruleAction67, position)
													}
													add(ruleCcf, position696)
												}
												break
											case 'S', 's':
												{
													position705 := position
													{
														position706 := position
														{
															position707, tokenIndex707 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l708
															}
															position++
															goto l707
														l708:
															position, tokenIndex = position707, tokenIndex707
															if buffer[position] != rune('S') {
																goto l614
															}
															position++
														}
													l707:
														{
															position709, tokenIndex709 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l710
															}
															position++
															goto l709
														l710:
															position, tokenIndex = position709, tokenIndex709
															if buffer[position] != rune('C') {
																goto l614
															}
															position++
														}
													l709:
														{
															position711, tokenIndex711 := position, tokenIndex
															if buffer[position] != rune('f') {
																goto l712
															}
															position++
															goto l711
														l712:
															position, tokenIndex = position711, tokenIndex711
															if buffer[position] != rune('F') {
																goto l614
															}
															position++
														}
													l711:
														add(rulePegText, position706)
													}
													{
														add(ruleAction66, position)
													}
													add(ruleScf, position705)
												}
												break
											case 'R', 'r':
												{
													position714 := position
													{
														position715 := position
														{
															position716, tokenIndex716 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l717
															}
															position++
															goto l716
														l717:
															position, tokenIndex = position716, tokenIndex716
															if buffer[position] != rune('R') {
																goto l614
															}
															position++
														}
													l716:
														{
															position718, tokenIndex718 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l719
															}
															position++
															goto l718
														l719:
															position, tokenIndex = position718, tokenIndex718
															if buffer[position] != rune('R') {
																goto l614
															}
															position++
														}
													l718:
														{
															position720, tokenIndex720 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l721
															}
															position++
															goto l720
														l721:
															position, tokenIndex = position720, tokenIndex720
															if buffer[position] != rune('A') {
																goto l614
															}
															position++
														}
													l720:
														add(rulePegText, position715)
													}
													{
														add(ruleAction63, position)
													}
													add(ruleRra, position714)
												}
												break
											case 'H', 'h':
												{
													position723 := position
													{
														position724 := position
														{
															position725, tokenIndex725 := position, tokenIndex
															if buffer[position] != rune('h') {
																goto l726
															}
															position++
															goto l725
														l726:
															position, tokenIndex = position725, tokenIndex725
															if buffer[position] != rune('H') {
																goto l614
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
																goto l614
															}
															position++
														}
													l727:
														{
															position729, tokenIndex729 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l730
															}
															position++
															goto l729
														l730:
															position, tokenIndex = position729, tokenIndex729
															if buffer[position] != rune('L') {
																goto l614
															}
															position++
														}
													l729:
														{
															position731, tokenIndex731 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l732
															}
															position++
															goto l731
														l732:
															position, tokenIndex = position731, tokenIndex731
															if buffer[position] != rune('T') {
																goto l614
															}
															position++
														}
													l731:
														add(rulePegText, position724)
													}
													{
														add(ruleAction59, position)
													}
													add(ruleHalt, position723)
												}
												break
											default:
												{
													position734 := position
													{
														position735 := position
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
																goto l614
															}
															position++
														}
													l736:
														{
															position738, tokenIndex738 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l739
															}
															position++
															goto l738
														l739:
															position, tokenIndex = position738, tokenIndex738
															if buffer[position] != rune('O') {
																goto l614
															}
															position++
														}
													l738:
														{
															position740, tokenIndex740 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l741
															}
															position++
															goto l740
														l741:
															position, tokenIndex = position740, tokenIndex740
															if buffer[position] != rune('P') {
																goto l614
															}
															position++
														}
													l740:
														add(rulePegText, position735)
													}
													{
														add(ruleAction58, position)
													}
													add(ruleNop, position734)
												}
												break
											}
										}

									}
								l616:
									add(ruleSimple, position615)
								}
								goto l24
							l614:
								position, tokenIndex = position24, tokenIndex24
								{
									position744 := position
									{
										position745, tokenIndex745 := position, tokenIndex
										{
											position747 := position
											{
												position748, tokenIndex748 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l749
												}
												position++
												goto l748
											l749:
												position, tokenIndex = position748, tokenIndex748
												if buffer[position] != rune('R') {
													goto l746
												}
												position++
											}
										l748:
											{
												position750, tokenIndex750 := position, tokenIndex
												if buffer[position] != rune('s') {
													goto l751
												}
												position++
												goto l750
											l751:
												position, tokenIndex = position750, tokenIndex750
												if buffer[position] != rune('S') {
													goto l746
												}
												position++
											}
										l750:
											{
												position752, tokenIndex752 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l753
												}
												position++
												goto l752
											l753:
												position, tokenIndex = position752, tokenIndex752
												if buffer[position] != rune('T') {
													goto l746
												}
												position++
											}
										l752:
											if !_rules[rulews]() {
												goto l746
											}
											if !_rules[rulen]() {
												goto l746
											}
											{
												add(ruleAction95, position)
											}
											add(ruleRst, position747)
										}
										goto l745
									l746:
										position, tokenIndex = position745, tokenIndex745
										{
											position756 := position
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
													goto l755
												}
												position++
											}
										l757:
											{
												position759, tokenIndex759 := position, tokenIndex
												if buffer[position] != rune('p') {
													goto l760
												}
												position++
												goto l759
											l760:
												position, tokenIndex = position759, tokenIndex759
												if buffer[position] != rune('P') {
													goto l755
												}
												position++
											}
										l759:
											if !_rules[rulews]() {
												goto l755
											}
											{
												position761, tokenIndex761 := position, tokenIndex
												if !_rules[rulecc]() {
													goto l761
												}
												if !_rules[rulesep]() {
													goto l761
												}
												goto l762
											l761:
												position, tokenIndex = position761, tokenIndex761
											}
										l762:
											if !_rules[ruleSrc16]() {
												goto l755
											}
											{
												add(ruleAction98, position)
											}
											add(ruleJp, position756)
										}
										goto l745
									l755:
										position, tokenIndex = position745, tokenIndex745
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position765 := position
													{
														position766, tokenIndex766 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l767
														}
														position++
														goto l766
													l767:
														position, tokenIndex = position766, tokenIndex766
														if buffer[position] != rune('D') {
															goto l743
														}
														position++
													}
												l766:
													{
														position768, tokenIndex768 := position, tokenIndex
														if buffer[position] != rune('j') {
															goto l769
														}
														position++
														goto l768
													l769:
														position, tokenIndex = position768, tokenIndex768
														if buffer[position] != rune('J') {
															goto l743
														}
														position++
													}
												l768:
													{
														position770, tokenIndex770 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l771
														}
														position++
														goto l770
													l771:
														position, tokenIndex = position770, tokenIndex770
														if buffer[position] != rune('N') {
															goto l743
														}
														position++
													}
												l770:
													{
														position772, tokenIndex772 := position, tokenIndex
														if buffer[position] != rune('z') {
															goto l773
														}
														position++
														goto l772
													l773:
														position, tokenIndex = position772, tokenIndex772
														if buffer[position] != rune('Z') {
															goto l743
														}
														position++
													}
												l772:
													if !_rules[rulews]() {
														goto l743
													}
													if !_rules[ruledisp]() {
														goto l743
													}
													{
														add(ruleAction100, position)
													}
													add(ruleDjnz, position765)
												}
												break
											case 'J', 'j':
												{
													position775 := position
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
															goto l743
														}
														position++
													}
												l776:
													{
														position778, tokenIndex778 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l779
														}
														position++
														goto l778
													l779:
														position, tokenIndex = position778, tokenIndex778
														if buffer[position] != rune('R') {
															goto l743
														}
														position++
													}
												l778:
													if !_rules[rulews]() {
														goto l743
													}
													{
														position780, tokenIndex780 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l780
														}
														if !_rules[rulesep]() {
															goto l780
														}
														goto l781
													l780:
														position, tokenIndex = position780, tokenIndex780
													}
												l781:
													if !_rules[ruledisp]() {
														goto l743
													}
													{
														add(ruleAction99, position)
													}
													add(ruleJr, position775)
												}
												break
											case 'R', 'r':
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
															goto l743
														}
														position++
													}
												l784:
													{
														position786, tokenIndex786 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l787
														}
														position++
														goto l786
													l787:
														position, tokenIndex = position786, tokenIndex786
														if buffer[position] != rune('E') {
															goto l743
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
															goto l743
														}
														position++
													}
												l788:
													{
														position790, tokenIndex790 := position, tokenIndex
														if !_rules[rulews]() {
															goto l790
														}
														if !_rules[rulecc]() {
															goto l790
														}
														goto l791
													l790:
														position, tokenIndex = position790, tokenIndex790
													}
												l791:
													{
														add(ruleAction97, position)
													}
													add(ruleRet, position783)
												}
												break
											default:
												{
													position793 := position
													{
														position794, tokenIndex794 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l795
														}
														position++
														goto l794
													l795:
														position, tokenIndex = position794, tokenIndex794
														if buffer[position] != rune('C') {
															goto l743
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
															goto l743
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
															goto l743
														}
														position++
													}
												l798:
													{
														position800, tokenIndex800 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l801
														}
														position++
														goto l800
													l801:
														position, tokenIndex = position800, tokenIndex800
														if buffer[position] != rune('L') {
															goto l743
														}
														position++
													}
												l800:
													if !_rules[rulews]() {
														goto l743
													}
													{
														position802, tokenIndex802 := position, tokenIndex
														if !_rules[rulecc]() {
															goto l802
														}
														if !_rules[rulesep]() {
															goto l802
														}
														goto l803
													l802:
														position, tokenIndex = position802, tokenIndex802
													}
												l803:
													if !_rules[ruleSrc16]() {
														goto l743
													}
													{
														add(ruleAction96, position)
													}
													add(ruleCall, position793)
												}
												break
											}
										}

									}
								l745:
									add(ruleJump, position744)
								}
								goto l24
							l743:
								position, tokenIndex = position24, tokenIndex24
								{
									position805 := position
									{
										position806, tokenIndex806 := position, tokenIndex
										{
											position808 := position
											{
												position809, tokenIndex809 := position, tokenIndex
												if buffer[position] != rune('i') {
													goto l810
												}
												position++
												goto l809
											l810:
												position, tokenIndex = position809, tokenIndex809
												if buffer[position] != rune('I') {
													goto l807
												}
												position++
											}
										l809:
											{
												position811, tokenIndex811 := position, tokenIndex
												if buffer[position] != rune('n') {
													goto l812
												}
												position++
												goto l811
											l812:
												position, tokenIndex = position811, tokenIndex811
												if buffer[position] != rune('N') {
													goto l807
												}
												position++
											}
										l811:
											if !_rules[rulews]() {
												goto l807
											}
											if !_rules[ruleReg8]() {
												goto l807
											}
											if !_rules[rulesep]() {
												goto l807
											}
											if !_rules[rulePort]() {
												goto l807
											}
											{
												add(ruleAction101, position)
											}
											add(ruleIN, position808)
										}
										goto l806
									l807:
										position, tokenIndex = position806, tokenIndex806
										{
											position814 := position
											{
												position815, tokenIndex815 := position, tokenIndex
												if buffer[position] != rune('o') {
													goto l816
												}
												position++
												goto l815
											l816:
												position, tokenIndex = position815, tokenIndex815
												if buffer[position] != rune('O') {
													goto l0
												}
												position++
											}
										l815:
											{
												position817, tokenIndex817 := position, tokenIndex
												if buffer[position] != rune('u') {
													goto l818
												}
												position++
												goto l817
											l818:
												position, tokenIndex = position817, tokenIndex817
												if buffer[position] != rune('U') {
													goto l0
												}
												position++
											}
										l817:
											{
												position819, tokenIndex819 := position, tokenIndex
												if buffer[position] != rune('t') {
													goto l820
												}
												position++
												goto l819
											l820:
												position, tokenIndex = position819, tokenIndex819
												if buffer[position] != rune('T') {
													goto l0
												}
												position++
											}
										l819:
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
												add(ruleAction102, position)
											}
											add(ruleOUT, position814)
										}
									}
								l806:
									add(ruleIO, position805)
								}
							}
						l24:
							add(ruleInstruction, position23)
						}
						{
							position822, tokenIndex822 := position, tokenIndex
							if !_rules[ruleLineEnd]() {
								goto l822
							}
							goto l823
						l822:
							position, tokenIndex = position822, tokenIndex822
						}
					l823:
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
						position825, tokenIndex825 := position, tokenIndex
						{
							position827 := position
						l828:
							{
								position829, tokenIndex829 := position, tokenIndex
								if !_rules[rulews]() {
									goto l829
								}
								goto l828
							l829:
								position, tokenIndex = position829, tokenIndex829
							}
							if !_rules[ruleLineEnd]() {
								goto l826
							}
							add(ruleBlankLine, position827)
						}
						goto l825
					l826:
						position, tokenIndex = position825, tokenIndex825
						{
							position830 := position
							{
								position831, tokenIndex831 := position, tokenIndex
								{
									position833 := position
									{
										position834 := position
										if !_rules[rulealpha]() {
											goto l831
										}
									l835:
										{
											position836, tokenIndex836 := position, tokenIndex
											{
												position837 := position
												{
													position838, tokenIndex838 := position, tokenIndex
													if !_rules[rulealpha]() {
														goto l839
													}
													goto l838
												l839:
													position, tokenIndex = position838, tokenIndex838
													{
														position840 := position
														if c := buffer[position]; c < rune('0') || c > rune('9') {
															goto l836
														}
														position++
														add(rulenum, position840)
													}
												}
											l838:
												add(rulealphanum, position837)
											}
											goto l835
										l836:
											position, tokenIndex = position836, tokenIndex836
										}
										add(rulePegText, position834)
									}
									if buffer[position] != rune(':') {
										goto l831
									}
									position++
									{
										add(ruleAction1, position)
									}
									add(ruleLabel, position833)
								}
								goto l832
							l831:
								position, tokenIndex = position831, tokenIndex831
							}
						l832:
						l842:
							{
								position843, tokenIndex843 := position, tokenIndex
								if !_rules[rulews]() {
									goto l843
								}
								goto l842
							l843:
								position, tokenIndex = position843, tokenIndex843
							}
							{
								position844 := position
								{
									position845, tokenIndex845 := position, tokenIndex
									{
										position847 := position
										{
											position848, tokenIndex848 := position, tokenIndex
											{
												position850 := position
												{
													position851, tokenIndex851 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l852
													}
													position++
													goto l851
												l852:
													position, tokenIndex = position851, tokenIndex851
													if buffer[position] != rune('P') {
														goto l849
													}
													position++
												}
											l851:
												{
													position853, tokenIndex853 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l854
													}
													position++
													goto l853
												l854:
													position, tokenIndex = position853, tokenIndex853
													if buffer[position] != rune('U') {
														goto l849
													}
													position++
												}
											l853:
												{
													position855, tokenIndex855 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l856
													}
													position++
													goto l855
												l856:
													position, tokenIndex = position855, tokenIndex855
													if buffer[position] != rune('S') {
														goto l849
													}
													position++
												}
											l855:
												{
													position857, tokenIndex857 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l858
													}
													position++
													goto l857
												l858:
													position, tokenIndex = position857, tokenIndex857
													if buffer[position] != rune('H') {
														goto l849
													}
													position++
												}
											l857:
												if !_rules[rulews]() {
													goto l849
												}
												if !_rules[ruleSrc16]() {
													goto l849
												}
												{
													add(ruleAction4, position)
												}
												add(rulePush, position850)
											}
											goto l848
										l849:
											position, tokenIndex = position848, tokenIndex848
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position861 := position
														{
															position862, tokenIndex862 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l863
															}
															position++
															goto l862
														l863:
															position, tokenIndex = position862, tokenIndex862
															if buffer[position] != rune('E') {
																goto l846
															}
															position++
														}
													l862:
														{
															position864, tokenIndex864 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l865
															}
															position++
															goto l864
														l865:
															position, tokenIndex = position864, tokenIndex864
															if buffer[position] != rune('X') {
																goto l846
															}
															position++
														}
													l864:
														if !_rules[rulews]() {
															goto l846
														}
														if !_rules[ruleDst16]() {
															goto l846
														}
														if !_rules[rulesep]() {
															goto l846
														}
														if !_rules[ruleSrc16]() {
															goto l846
														}
														{
															add(ruleAction6, position)
														}
														add(ruleEx, position861)
													}
													break
												case 'P', 'p':
													{
														position867 := position
														{
															position868, tokenIndex868 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l869
															}
															position++
															goto l868
														l869:
															position, tokenIndex = position868, tokenIndex868
															if buffer[position] != rune('P') {
																goto l846
															}
															position++
														}
													l868:
														{
															position870, tokenIndex870 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l871
															}
															position++
															goto l870
														l871:
															position, tokenIndex = position870, tokenIndex870
															if buffer[position] != rune('O') {
																goto l846
															}
															position++
														}
													l870:
														{
															position872, tokenIndex872 := position, tokenIndex
															if buffer[position] != rune('p') {
																goto l873
															}
															position++
															goto l872
														l873:
															position, tokenIndex = position872, tokenIndex872
															if buffer[position] != rune('P') {
																goto l846
															}
															position++
														}
													l872:
														if !_rules[rulews]() {
															goto l846
														}
														if !_rules[ruleDst16]() {
															goto l846
														}
														{
															add(ruleAction5, position)
														}
														add(rulePop, position867)
													}
													break
												default:
													{
														position875 := position
														{
															position876, tokenIndex876 := position, tokenIndex
															{
																position878 := position
																{
																	position879, tokenIndex879 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l880
																	}
																	position++
																	goto l879
																l880:
																	position, tokenIndex = position879, tokenIndex879
																	if buffer[position] != rune('L') {
																		goto l877
																	}
																	position++
																}
															l879:
																{
																	position881, tokenIndex881 := position, tokenIndex
																	if buffer[position] != rune('d') {
																		goto l882
																	}
																	position++
																	goto l881
																l882:
																	position, tokenIndex = position881, tokenIndex881
																	if buffer[position] != rune('D') {
																		goto l877
																	}
																	position++
																}
															l881:
																if !_rules[rulews]() {
																	goto l877
																}
																if !_rules[ruleDst16]() {
																	goto l877
																}
																if !_rules[rulesep]() {
																	goto l877
																}
																if !_rules[ruleSrc16]() {
																	goto l877
																}
																{
																	add(ruleAction3, position)
																}
																add(ruleLoad16, position878)
															}
															goto l876
														l877:
															position, tokenIndex = position876, tokenIndex876
															{
																position884 := position
																{
																	position885, tokenIndex885 := position, tokenIndex
																	if buffer[position] != rune('l') {
																		goto l886
																	}
																	position++
																	goto l885
																l886:
																	position, tokenIndex = position885, tokenIndex885
																	if buffer[position] != rune('L') {
																		goto l846
																	}
																	position++
																}
															l885:
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
																		goto l846
																	}
																	position++
																}
															l887:
																if !_rules[rulews]() {
																	goto l846
																}
																{
																	position889 := position
																	{
																		position890, tokenIndex890 := position, tokenIndex
																		if !_rules[ruleReg8]() {
																			goto l891
																		}
																		goto l890
																	l891:
																		position, tokenIndex = position890, tokenIndex890
																		if !_rules[ruleReg16Contents]() {
																			goto l892
																		}
																		goto l890
																	l892:
																		position, tokenIndex = position890, tokenIndex890
																		if !_rules[rulenn_contents]() {
																			goto l846
																		}
																	}
																l890:
																	{
																		add(ruleAction16, position)
																	}
																	add(ruleDst8, position889)
																}
																if !_rules[rulesep]() {
																	goto l846
																}
																if !_rules[ruleSrc8]() {
																	goto l846
																}
																{
																	add(ruleAction2, position)
																}
																add(ruleLoad8, position884)
															}
														}
													l876:
														add(ruleLoad, position875)
													}
													break
												}
											}

										}
									l848:
										add(ruleAssignment, position847)
									}
									goto l845
								l846:
									position, tokenIndex = position845, tokenIndex845
									{
										position896 := position
										{
											position897, tokenIndex897 := position, tokenIndex
											{
												position899 := position
												{
													position900, tokenIndex900 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l901
													}
													position++
													goto l900
												l901:
													position, tokenIndex = position900, tokenIndex900
													if buffer[position] != rune('I') {
														goto l898
													}
													position++
												}
											l900:
												{
													position902, tokenIndex902 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l903
													}
													position++
													goto l902
												l903:
													position, tokenIndex = position902, tokenIndex902
													if buffer[position] != rune('N') {
														goto l898
													}
													position++
												}
											l902:
												{
													position904, tokenIndex904 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l905
													}
													position++
													goto l904
												l905:
													position, tokenIndex = position904, tokenIndex904
													if buffer[position] != rune('C') {
														goto l898
													}
													position++
												}
											l904:
												if !_rules[rulews]() {
													goto l898
												}
												if !_rules[ruleILoc8]() {
													goto l898
												}
												{
													add(ruleAction7, position)
												}
												add(ruleInc16Indexed8, position899)
											}
											goto l897
										l898:
											position, tokenIndex = position897, tokenIndex897
											{
												position908 := position
												{
													position909, tokenIndex909 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l910
													}
													position++
													goto l909
												l910:
													position, tokenIndex = position909, tokenIndex909
													if buffer[position] != rune('I') {
														goto l907
													}
													position++
												}
											l909:
												{
													position911, tokenIndex911 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l912
													}
													position++
													goto l911
												l912:
													position, tokenIndex = position911, tokenIndex911
													if buffer[position] != rune('N') {
														goto l907
													}
													position++
												}
											l911:
												{
													position913, tokenIndex913 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l914
													}
													position++
													goto l913
												l914:
													position, tokenIndex = position913, tokenIndex913
													if buffer[position] != rune('C') {
														goto l907
													}
													position++
												}
											l913:
												if !_rules[rulews]() {
													goto l907
												}
												if !_rules[ruleLoc16]() {
													goto l907
												}
												{
													add(ruleAction9, position)
												}
												add(ruleInc16, position908)
											}
											goto l897
										l907:
											position, tokenIndex = position897, tokenIndex897
											{
												position916 := position
												{
													position917, tokenIndex917 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l918
													}
													position++
													goto l917
												l918:
													position, tokenIndex = position917, tokenIndex917
													if buffer[position] != rune('I') {
														goto l895
													}
													position++
												}
											l917:
												{
													position919, tokenIndex919 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l920
													}
													position++
													goto l919
												l920:
													position, tokenIndex = position919, tokenIndex919
													if buffer[position] != rune('N') {
														goto l895
													}
													position++
												}
											l919:
												{
													position921, tokenIndex921 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l922
													}
													position++
													goto l921
												l922:
													position, tokenIndex = position921, tokenIndex921
													if buffer[position] != rune('C') {
														goto l895
													}
													position++
												}
											l921:
												if !_rules[rulews]() {
													goto l895
												}
												if !_rules[ruleLoc8]() {
													goto l895
												}
												{
													add(ruleAction8, position)
												}
												add(ruleInc8, position916)
											}
										}
									l897:
										add(ruleInc, position896)
									}
									goto l845
								l895:
									position, tokenIndex = position845, tokenIndex845
									{
										position925 := position
										{
											position926, tokenIndex926 := position, tokenIndex
											{
												position928 := position
												{
													position929, tokenIndex929 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l930
													}
													position++
													goto l929
												l930:
													position, tokenIndex = position929, tokenIndex929
													if buffer[position] != rune('D') {
														goto l927
													}
													position++
												}
											l929:
												{
													position931, tokenIndex931 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l932
													}
													position++
													goto l931
												l932:
													position, tokenIndex = position931, tokenIndex931
													if buffer[position] != rune('E') {
														goto l927
													}
													position++
												}
											l931:
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
														goto l927
													}
													position++
												}
											l933:
												if !_rules[rulews]() {
													goto l927
												}
												if !_rules[ruleILoc8]() {
													goto l927
												}
												{
													add(ruleAction10, position)
												}
												add(ruleDec16Indexed8, position928)
											}
											goto l926
										l927:
											position, tokenIndex = position926, tokenIndex926
											{
												position937 := position
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
														goto l936
													}
													position++
												}
											l938:
												{
													position940, tokenIndex940 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l941
													}
													position++
													goto l940
												l941:
													position, tokenIndex = position940, tokenIndex940
													if buffer[position] != rune('E') {
														goto l936
													}
													position++
												}
											l940:
												{
													position942, tokenIndex942 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l943
													}
													position++
													goto l942
												l943:
													position, tokenIndex = position942, tokenIndex942
													if buffer[position] != rune('C') {
														goto l936
													}
													position++
												}
											l942:
												if !_rules[rulews]() {
													goto l936
												}
												if !_rules[ruleLoc16]() {
													goto l936
												}
												{
													add(ruleAction12, position)
												}
												add(ruleDec16, position937)
											}
											goto l926
										l936:
											position, tokenIndex = position926, tokenIndex926
											{
												position945 := position
												{
													position946, tokenIndex946 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l947
													}
													position++
													goto l946
												l947:
													position, tokenIndex = position946, tokenIndex946
													if buffer[position] != rune('D') {
														goto l924
													}
													position++
												}
											l946:
												{
													position948, tokenIndex948 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l949
													}
													position++
													goto l948
												l949:
													position, tokenIndex = position948, tokenIndex948
													if buffer[position] != rune('E') {
														goto l924
													}
													position++
												}
											l948:
												{
													position950, tokenIndex950 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l951
													}
													position++
													goto l950
												l951:
													position, tokenIndex = position950, tokenIndex950
													if buffer[position] != rune('C') {
														goto l924
													}
													position++
												}
											l950:
												if !_rules[rulews]() {
													goto l924
												}
												if !_rules[ruleLoc8]() {
													goto l924
												}
												{
													add(ruleAction11, position)
												}
												add(ruleDec8, position945)
											}
										}
									l926:
										add(ruleDec, position925)
									}
									goto l845
								l924:
									position, tokenIndex = position845, tokenIndex845
									{
										position954 := position
										{
											position955, tokenIndex955 := position, tokenIndex
											{
												position957 := position
												{
													position958, tokenIndex958 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l959
													}
													position++
													goto l958
												l959:
													position, tokenIndex = position958, tokenIndex958
													if buffer[position] != rune('A') {
														goto l956
													}
													position++
												}
											l958:
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
														goto l956
													}
													position++
												}
											l960:
												{
													position962, tokenIndex962 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l963
													}
													position++
													goto l962
												l963:
													position, tokenIndex = position962, tokenIndex962
													if buffer[position] != rune('D') {
														goto l956
													}
													position++
												}
											l962:
												if !_rules[rulews]() {
													goto l956
												}
												if !_rules[ruleDst16]() {
													goto l956
												}
												if !_rules[rulesep]() {
													goto l956
												}
												if !_rules[ruleSrc16]() {
													goto l956
												}
												{
													add(ruleAction13, position)
												}
												add(ruleAdd16, position957)
											}
											goto l955
										l956:
											position, tokenIndex = position955, tokenIndex955
											{
												position966 := position
												{
													position967, tokenIndex967 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l968
													}
													position++
													goto l967
												l968:
													position, tokenIndex = position967, tokenIndex967
													if buffer[position] != rune('A') {
														goto l965
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
														goto l965
													}
													position++
												}
											l969:
												{
													position971, tokenIndex971 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l972
													}
													position++
													goto l971
												l972:
													position, tokenIndex = position971, tokenIndex971
													if buffer[position] != rune('C') {
														goto l965
													}
													position++
												}
											l971:
												if !_rules[rulews]() {
													goto l965
												}
												if !_rules[ruleDst16]() {
													goto l965
												}
												if !_rules[rulesep]() {
													goto l965
												}
												if !_rules[ruleSrc16]() {
													goto l965
												}
												{
													add(ruleAction14, position)
												}
												add(ruleAdc16, position966)
											}
											goto l955
										l965:
											position, tokenIndex = position955, tokenIndex955
											{
												position974 := position
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
														goto l953
													}
													position++
												}
											l975:
												{
													position977, tokenIndex977 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l978
													}
													position++
													goto l977
												l978:
													position, tokenIndex = position977, tokenIndex977
													if buffer[position] != rune('B') {
														goto l953
													}
													position++
												}
											l977:
												{
													position979, tokenIndex979 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l980
													}
													position++
													goto l979
												l980:
													position, tokenIndex = position979, tokenIndex979
													if buffer[position] != rune('C') {
														goto l953
													}
													position++
												}
											l979:
												if !_rules[rulews]() {
													goto l953
												}
												if !_rules[ruleDst16]() {
													goto l953
												}
												if !_rules[rulesep]() {
													goto l953
												}
												if !_rules[ruleSrc16]() {
													goto l953
												}
												{
													add(ruleAction15, position)
												}
												add(ruleSbc16, position974)
											}
										}
									l955:
										add(ruleAlu16, position954)
									}
									goto l845
								l953:
									position, tokenIndex = position845, tokenIndex845
									{
										position983 := position
										{
											position984, tokenIndex984 := position, tokenIndex
											{
												position986 := position
												{
													position987, tokenIndex987 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l988
													}
													position++
													goto l987
												l988:
													position, tokenIndex = position987, tokenIndex987
													if buffer[position] != rune('A') {
														goto l985
													}
													position++
												}
											l987:
												{
													position989, tokenIndex989 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l990
													}
													position++
													goto l989
												l990:
													position, tokenIndex = position989, tokenIndex989
													if buffer[position] != rune('D') {
														goto l985
													}
													position++
												}
											l989:
												{
													position991, tokenIndex991 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l992
													}
													position++
													goto l991
												l992:
													position, tokenIndex = position991, tokenIndex991
													if buffer[position] != rune('D') {
														goto l985
													}
													position++
												}
											l991:
												if !_rules[rulews]() {
													goto l985
												}
												{
													position993, tokenIndex993 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l994
													}
													position++
													goto l993
												l994:
													position, tokenIndex = position993, tokenIndex993
													if buffer[position] != rune('A') {
														goto l985
													}
													position++
												}
											l993:
												if !_rules[rulesep]() {
													goto l985
												}
												if !_rules[ruleSrc8]() {
													goto l985
												}
												{
													add(ruleAction39, position)
												}
												add(ruleAdd, position986)
											}
											goto l984
										l985:
											position, tokenIndex = position984, tokenIndex984
											{
												position997 := position
												{
													position998, tokenIndex998 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l999
													}
													position++
													goto l998
												l999:
													position, tokenIndex = position998, tokenIndex998
													if buffer[position] != rune('A') {
														goto l996
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
														goto l996
													}
													position++
												}
											l1000:
												{
													position1002, tokenIndex1002 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l1003
													}
													position++
													goto l1002
												l1003:
													position, tokenIndex = position1002, tokenIndex1002
													if buffer[position] != rune('C') {
														goto l996
													}
													position++
												}
											l1002:
												if !_rules[rulews]() {
													goto l996
												}
												{
													position1004, tokenIndex1004 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l1005
													}
													position++
													goto l1004
												l1005:
													position, tokenIndex = position1004, tokenIndex1004
													if buffer[position] != rune('A') {
														goto l996
													}
													position++
												}
											l1004:
												if !_rules[rulesep]() {
													goto l996
												}
												if !_rules[ruleSrc8]() {
													goto l996
												}
												{
													add(ruleAction40, position)
												}
												add(ruleAdc, position997)
											}
											goto l984
										l996:
											position, tokenIndex = position984, tokenIndex984
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
														goto l1007
													}
													position++
												}
											l1009:
												{
													position1011, tokenIndex1011 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1012
													}
													position++
													goto l1011
												l1012:
													position, tokenIndex = position1011, tokenIndex1011
													if buffer[position] != rune('U') {
														goto l1007
													}
													position++
												}
											l1011:
												{
													position1013, tokenIndex1013 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l1014
													}
													position++
													goto l1013
												l1014:
													position, tokenIndex = position1013, tokenIndex1013
													if buffer[position] != rune('B') {
														goto l1007
													}
													position++
												}
											l1013:
												if !_rules[rulews]() {
													goto l1007
												}
												if !_rules[ruleSrc8]() {
													goto l1007
												}
												{
													add(ruleAction41, position)
												}
												add(ruleSub, position1008)
											}
											goto l984
										l1007:
											position, tokenIndex = position984, tokenIndex984
											{
												switch buffer[position] {
												case 'C', 'c':
													{
														position1017 := position
														{
															position1018, tokenIndex1018 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1019
															}
															position++
															goto l1018
														l1019:
															position, tokenIndex = position1018, tokenIndex1018
															if buffer[position] != rune('C') {
																goto l982
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
																goto l982
															}
															position++
														}
													l1020:
														if !_rules[rulews]() {
															goto l982
														}
														if !_rules[ruleSrc8]() {
															goto l982
														}
														{
															add(ruleAction46, position)
														}
														add(ruleCp, position1017)
													}
													break
												case 'O', 'o':
													{
														position1023 := position
														{
															position1024, tokenIndex1024 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1025
															}
															position++
															goto l1024
														l1025:
															position, tokenIndex = position1024, tokenIndex1024
															if buffer[position] != rune('O') {
																goto l982
															}
															position++
														}
													l1024:
														{
															position1026, tokenIndex1026 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1027
															}
															position++
															goto l1026
														l1027:
															position, tokenIndex = position1026, tokenIndex1026
															if buffer[position] != rune('R') {
																goto l982
															}
															position++
														}
													l1026:
														if !_rules[rulews]() {
															goto l982
														}
														if !_rules[ruleSrc8]() {
															goto l982
														}
														{
															add(ruleAction45, position)
														}
														add(ruleOr, position1023)
													}
													break
												case 'X', 'x':
													{
														position1029 := position
														{
															position1030, tokenIndex1030 := position, tokenIndex
															if buffer[position] != rune('x') {
																goto l1031
															}
															position++
															goto l1030
														l1031:
															position, tokenIndex = position1030, tokenIndex1030
															if buffer[position] != rune('X') {
																goto l982
															}
															position++
														}
													l1030:
														{
															position1032, tokenIndex1032 := position, tokenIndex
															if buffer[position] != rune('o') {
																goto l1033
															}
															position++
															goto l1032
														l1033:
															position, tokenIndex = position1032, tokenIndex1032
															if buffer[position] != rune('O') {
																goto l982
															}
															position++
														}
													l1032:
														{
															position1034, tokenIndex1034 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1035
															}
															position++
															goto l1034
														l1035:
															position, tokenIndex = position1034, tokenIndex1034
															if buffer[position] != rune('R') {
																goto l982
															}
															position++
														}
													l1034:
														if !_rules[rulews]() {
															goto l982
														}
														if !_rules[ruleSrc8]() {
															goto l982
														}
														{
															add(ruleAction44, position)
														}
														add(ruleXor, position1029)
													}
													break
												case 'A', 'a':
													{
														position1037 := position
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
																goto l982
															}
															position++
														}
													l1038:
														{
															position1040, tokenIndex1040 := position, tokenIndex
															if buffer[position] != rune('n') {
																goto l1041
															}
															position++
															goto l1040
														l1041:
															position, tokenIndex = position1040, tokenIndex1040
															if buffer[position] != rune('N') {
																goto l982
															}
															position++
														}
													l1040:
														{
															position1042, tokenIndex1042 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1043
															}
															position++
															goto l1042
														l1043:
															position, tokenIndex = position1042, tokenIndex1042
															if buffer[position] != rune('D') {
																goto l982
															}
															position++
														}
													l1042:
														if !_rules[rulews]() {
															goto l982
														}
														if !_rules[ruleSrc8]() {
															goto l982
														}
														{
															add(ruleAction43, position)
														}
														add(ruleAnd, position1037)
													}
													break
												default:
													{
														position1045 := position
														{
															position1046, tokenIndex1046 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1047
															}
															position++
															goto l1046
														l1047:
															position, tokenIndex = position1046, tokenIndex1046
															if buffer[position] != rune('S') {
																goto l982
															}
															position++
														}
													l1046:
														{
															position1048, tokenIndex1048 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1049
															}
															position++
															goto l1048
														l1049:
															position, tokenIndex = position1048, tokenIndex1048
															if buffer[position] != rune('B') {
																goto l982
															}
															position++
														}
													l1048:
														{
															position1050, tokenIndex1050 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1051
															}
															position++
															goto l1050
														l1051:
															position, tokenIndex = position1050, tokenIndex1050
															if buffer[position] != rune('C') {
																goto l982
															}
															position++
														}
													l1050:
														if !_rules[rulews]() {
															goto l982
														}
														{
															position1052, tokenIndex1052 := position, tokenIndex
															if buffer[position] != rune('a') {
																goto l1053
															}
															position++
															goto l1052
														l1053:
															position, tokenIndex = position1052, tokenIndex1052
															if buffer[position] != rune('A') {
																goto l982
															}
															position++
														}
													l1052:
														if !_rules[rulesep]() {
															goto l982
														}
														if !_rules[ruleSrc8]() {
															goto l982
														}
														{
															add(ruleAction42, position)
														}
														add(ruleSbc, position1045)
													}
													break
												}
											}

										}
									l984:
										add(ruleAlu, position983)
									}
									goto l845
								l982:
									position, tokenIndex = position845, tokenIndex845
									{
										position1056 := position
										{
											position1057, tokenIndex1057 := position, tokenIndex
											{
												position1059 := position
												{
													position1060, tokenIndex1060 := position, tokenIndex
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
														{
															position1067, tokenIndex1067 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1068
															}
															position++
															goto l1067
														l1068:
															position, tokenIndex = position1067, tokenIndex1067
															if buffer[position] != rune('C') {
																goto l1061
															}
															position++
														}
													l1067:
														if !_rules[rulews]() {
															goto l1061
														}
														if !_rules[ruleLoc8]() {
															goto l1061
														}
														{
															position1069, tokenIndex1069 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1069
															}
															if !_rules[ruleCopy8]() {
																goto l1069
															}
															goto l1070
														l1069:
															position, tokenIndex = position1069, tokenIndex1069
														}
													l1070:
														{
															add(ruleAction47, position)
														}
														add(ruleRlc, position1062)
													}
													goto l1060
												l1061:
													position, tokenIndex = position1060, tokenIndex1060
													{
														position1073 := position
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
																goto l1072
															}
															position++
														}
													l1074:
														{
															position1076, tokenIndex1076 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1077
															}
															position++
															goto l1076
														l1077:
															position, tokenIndex = position1076, tokenIndex1076
															if buffer[position] != rune('R') {
																goto l1072
															}
															position++
														}
													l1076:
														{
															position1078, tokenIndex1078 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1079
															}
															position++
															goto l1078
														l1079:
															position, tokenIndex = position1078, tokenIndex1078
															if buffer[position] != rune('C') {
																goto l1072
															}
															position++
														}
													l1078:
														if !_rules[rulews]() {
															goto l1072
														}
														if !_rules[ruleLoc8]() {
															goto l1072
														}
														{
															position1080, tokenIndex1080 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1080
															}
															if !_rules[ruleCopy8]() {
																goto l1080
															}
															goto l1081
														l1080:
															position, tokenIndex = position1080, tokenIndex1080
														}
													l1081:
														{
															add(ruleAction48, position)
														}
														add(ruleRrc, position1073)
													}
													goto l1060
												l1072:
													position, tokenIndex = position1060, tokenIndex1060
													{
														position1084 := position
														{
															position1085, tokenIndex1085 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1086
															}
															position++
															goto l1085
														l1086:
															position, tokenIndex = position1085, tokenIndex1085
															if buffer[position] != rune('R') {
																goto l1083
															}
															position++
														}
													l1085:
														{
															position1087, tokenIndex1087 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1088
															}
															position++
															goto l1087
														l1088:
															position, tokenIndex = position1087, tokenIndex1087
															if buffer[position] != rune('L') {
																goto l1083
															}
															position++
														}
													l1087:
														if !_rules[rulews]() {
															goto l1083
														}
														if !_rules[ruleLoc8]() {
															goto l1083
														}
														{
															position1089, tokenIndex1089 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1089
															}
															if !_rules[ruleCopy8]() {
																goto l1089
															}
															goto l1090
														l1089:
															position, tokenIndex = position1089, tokenIndex1089
														}
													l1090:
														{
															add(ruleAction49, position)
														}
														add(ruleRl, position1084)
													}
													goto l1060
												l1083:
													position, tokenIndex = position1060, tokenIndex1060
													{
														position1093 := position
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
																goto l1092
															}
															position++
														}
													l1094:
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
																goto l1092
															}
															position++
														}
													l1096:
														if !_rules[rulews]() {
															goto l1092
														}
														if !_rules[ruleLoc8]() {
															goto l1092
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
															add(ruleAction50, position)
														}
														add(ruleRr, position1093)
													}
													goto l1060
												l1092:
													position, tokenIndex = position1060, tokenIndex1060
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
															if buffer[position] != rune('a') {
																goto l1108
															}
															position++
															goto l1107
														l1108:
															position, tokenIndex = position1107, tokenIndex1107
															if buffer[position] != rune('A') {
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
															add(ruleAction51, position)
														}
														add(ruleSla, position1102)
													}
													goto l1060
												l1101:
													position, tokenIndex = position1060, tokenIndex1060
													{
														position1113 := position
														{
															position1114, tokenIndex1114 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1115
															}
															position++
															goto l1114
														l1115:
															position, tokenIndex = position1114, tokenIndex1114
															if buffer[position] != rune('S') {
																goto l1112
															}
															position++
														}
													l1114:
														{
															position1116, tokenIndex1116 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1117
															}
															position++
															goto l1116
														l1117:
															position, tokenIndex = position1116, tokenIndex1116
															if buffer[position] != rune('R') {
																goto l1112
															}
															position++
														}
													l1116:
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
																goto l1112
															}
															position++
														}
													l1118:
														if !_rules[rulews]() {
															goto l1112
														}
														if !_rules[ruleLoc8]() {
															goto l1112
														}
														{
															position1120, tokenIndex1120 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1120
															}
															if !_rules[ruleCopy8]() {
																goto l1120
															}
															goto l1121
														l1120:
															position, tokenIndex = position1120, tokenIndex1120
														}
													l1121:
														{
															add(ruleAction52, position)
														}
														add(ruleSra, position1113)
													}
													goto l1060
												l1112:
													position, tokenIndex = position1060, tokenIndex1060
													{
														position1124 := position
														{
															position1125, tokenIndex1125 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1126
															}
															position++
															goto l1125
														l1126:
															position, tokenIndex = position1125, tokenIndex1125
															if buffer[position] != rune('S') {
																goto l1123
															}
															position++
														}
													l1125:
														{
															position1127, tokenIndex1127 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1128
															}
															position++
															goto l1127
														l1128:
															position, tokenIndex = position1127, tokenIndex1127
															if buffer[position] != rune('L') {
																goto l1123
															}
															position++
														}
													l1127:
														{
															position1129, tokenIndex1129 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1130
															}
															position++
															goto l1129
														l1130:
															position, tokenIndex = position1129, tokenIndex1129
															if buffer[position] != rune('L') {
																goto l1123
															}
															position++
														}
													l1129:
														if !_rules[rulews]() {
															goto l1123
														}
														if !_rules[ruleLoc8]() {
															goto l1123
														}
														{
															position1131, tokenIndex1131 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1131
															}
															if !_rules[ruleCopy8]() {
																goto l1131
															}
															goto l1132
														l1131:
															position, tokenIndex = position1131, tokenIndex1131
														}
													l1132:
														{
															add(ruleAction53, position)
														}
														add(ruleSll, position1124)
													}
													goto l1060
												l1123:
													position, tokenIndex = position1060, tokenIndex1060
													{
														position1134 := position
														{
															position1135, tokenIndex1135 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1136
															}
															position++
															goto l1135
														l1136:
															position, tokenIndex = position1135, tokenIndex1135
															if buffer[position] != rune('S') {
																goto l1058
															}
															position++
														}
													l1135:
														{
															position1137, tokenIndex1137 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1138
															}
															position++
															goto l1137
														l1138:
															position, tokenIndex = position1137, tokenIndex1137
															if buffer[position] != rune('R') {
																goto l1058
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
																goto l1058
															}
															position++
														}
													l1139:
														if !_rules[rulews]() {
															goto l1058
														}
														if !_rules[ruleLoc8]() {
															goto l1058
														}
														{
															position1141, tokenIndex1141 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1141
															}
															if !_rules[ruleCopy8]() {
																goto l1141
															}
															goto l1142
														l1141:
															position, tokenIndex = position1141, tokenIndex1141
														}
													l1142:
														{
															add(ruleAction54, position)
														}
														add(ruleSrl, position1134)
													}
												}
											l1060:
												add(ruleRot, position1059)
											}
											goto l1057
										l1058:
											position, tokenIndex = position1057, tokenIndex1057
											{
												switch buffer[position] {
												case 'S', 's':
													{
														position1145 := position
														{
															position1146, tokenIndex1146 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1147
															}
															position++
															goto l1146
														l1147:
															position, tokenIndex = position1146, tokenIndex1146
															if buffer[position] != rune('S') {
																goto l1055
															}
															position++
														}
													l1146:
														{
															position1148, tokenIndex1148 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1149
															}
															position++
															goto l1148
														l1149:
															position, tokenIndex = position1148, tokenIndex1148
															if buffer[position] != rune('E') {
																goto l1055
															}
															position++
														}
													l1148:
														{
															position1150, tokenIndex1150 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1151
															}
															position++
															goto l1150
														l1151:
															position, tokenIndex = position1150, tokenIndex1150
															if buffer[position] != rune('T') {
																goto l1055
															}
															position++
														}
													l1150:
														if !_rules[rulews]() {
															goto l1055
														}
														if !_rules[ruleoctaldigit]() {
															goto l1055
														}
														if !_rules[rulesep]() {
															goto l1055
														}
														if !_rules[ruleLoc8]() {
															goto l1055
														}
														{
															position1152, tokenIndex1152 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1152
															}
															if !_rules[ruleCopy8]() {
																goto l1152
															}
															goto l1153
														l1152:
															position, tokenIndex = position1152, tokenIndex1152
														}
													l1153:
														{
															add(ruleAction57, position)
														}
														add(ruleSet, position1145)
													}
													break
												case 'R', 'r':
													{
														position1155 := position
														{
															position1156, tokenIndex1156 := position, tokenIndex
															if buffer[position] != rune('r') {
																goto l1157
															}
															position++
															goto l1156
														l1157:
															position, tokenIndex = position1156, tokenIndex1156
															if buffer[position] != rune('R') {
																goto l1055
															}
															position++
														}
													l1156:
														{
															position1158, tokenIndex1158 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1159
															}
															position++
															goto l1158
														l1159:
															position, tokenIndex = position1158, tokenIndex1158
															if buffer[position] != rune('E') {
																goto l1055
															}
															position++
														}
													l1158:
														{
															position1160, tokenIndex1160 := position, tokenIndex
															if buffer[position] != rune('s') {
																goto l1161
															}
															position++
															goto l1160
														l1161:
															position, tokenIndex = position1160, tokenIndex1160
															if buffer[position] != rune('S') {
																goto l1055
															}
															position++
														}
													l1160:
														if !_rules[rulews]() {
															goto l1055
														}
														if !_rules[ruleoctaldigit]() {
															goto l1055
														}
														if !_rules[rulesep]() {
															goto l1055
														}
														if !_rules[ruleLoc8]() {
															goto l1055
														}
														{
															position1162, tokenIndex1162 := position, tokenIndex
															if !_rules[rulesep]() {
																goto l1162
															}
															if !_rules[ruleCopy8]() {
																goto l1162
															}
															goto l1163
														l1162:
															position, tokenIndex = position1162, tokenIndex1162
														}
													l1163:
														{
															add(ruleAction56, position)
														}
														add(ruleRes, position1155)
													}
													break
												default:
													{
														position1165 := position
														{
															position1166, tokenIndex1166 := position, tokenIndex
															if buffer[position] != rune('b') {
																goto l1167
															}
															position++
															goto l1166
														l1167:
															position, tokenIndex = position1166, tokenIndex1166
															if buffer[position] != rune('B') {
																goto l1055
															}
															position++
														}
													l1166:
														{
															position1168, tokenIndex1168 := position, tokenIndex
															if buffer[position] != rune('i') {
																goto l1169
															}
															position++
															goto l1168
														l1169:
															position, tokenIndex = position1168, tokenIndex1168
															if buffer[position] != rune('I') {
																goto l1055
															}
															position++
														}
													l1168:
														{
															position1170, tokenIndex1170 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1171
															}
															position++
															goto l1170
														l1171:
															position, tokenIndex = position1170, tokenIndex1170
															if buffer[position] != rune('T') {
																goto l1055
															}
															position++
														}
													l1170:
														if !_rules[rulews]() {
															goto l1055
														}
														if !_rules[ruleoctaldigit]() {
															goto l1055
														}
														if !_rules[rulesep]() {
															goto l1055
														}
														if !_rules[ruleLoc8]() {
															goto l1055
														}
														{
															add(ruleAction55, position)
														}
														add(ruleBit, position1165)
													}
													break
												}
											}

										}
									l1057:
										add(ruleBitOp, position1056)
									}
									goto l845
								l1055:
									position, tokenIndex = position845, tokenIndex845
									{
										position1174 := position
										{
											position1175, tokenIndex1175 := position, tokenIndex
											{
												position1177 := position
												{
													position1178 := position
													{
														position1179, tokenIndex1179 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1180
														}
														position++
														goto l1179
													l1180:
														position, tokenIndex = position1179, tokenIndex1179
														if buffer[position] != rune('R') {
															goto l1176
														}
														position++
													}
												l1179:
													{
														position1181, tokenIndex1181 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1182
														}
														position++
														goto l1181
													l1182:
														position, tokenIndex = position1181, tokenIndex1181
														if buffer[position] != rune('E') {
															goto l1176
														}
														position++
													}
												l1181:
													{
														position1183, tokenIndex1183 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1184
														}
														position++
														goto l1183
													l1184:
														position, tokenIndex = position1183, tokenIndex1183
														if buffer[position] != rune('T') {
															goto l1176
														}
														position++
													}
												l1183:
													{
														position1185, tokenIndex1185 := position, tokenIndex
														if buffer[position] != rune('n') {
															goto l1186
														}
														position++
														goto l1185
													l1186:
														position, tokenIndex = position1185, tokenIndex1185
														if buffer[position] != rune('N') {
															goto l1176
														}
														position++
													}
												l1185:
													add(rulePegText, position1178)
												}
												{
													add(ruleAction72, position)
												}
												add(ruleRetn, position1177)
											}
											goto l1175
										l1176:
											position, tokenIndex = position1175, tokenIndex1175
											{
												position1189 := position
												{
													position1190 := position
													{
														position1191, tokenIndex1191 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1192
														}
														position++
														goto l1191
													l1192:
														position, tokenIndex = position1191, tokenIndex1191
														if buffer[position] != rune('R') {
															goto l1188
														}
														position++
													}
												l1191:
													{
														position1193, tokenIndex1193 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1194
														}
														position++
														goto l1193
													l1194:
														position, tokenIndex = position1193, tokenIndex1193
														if buffer[position] != rune('E') {
															goto l1188
														}
														position++
													}
												l1193:
													{
														position1195, tokenIndex1195 := position, tokenIndex
														if buffer[position] != rune('t') {
															goto l1196
														}
														position++
														goto l1195
													l1196:
														position, tokenIndex = position1195, tokenIndex1195
														if buffer[position] != rune('T') {
															goto l1188
														}
														position++
													}
												l1195:
													{
														position1197, tokenIndex1197 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1198
														}
														position++
														goto l1197
													l1198:
														position, tokenIndex = position1197, tokenIndex1197
														if buffer[position] != rune('I') {
															goto l1188
														}
														position++
													}
												l1197:
													add(rulePegText, position1190)
												}
												{
													add(ruleAction73, position)
												}
												add(ruleReti, position1189)
											}
											goto l1175
										l1188:
											position, tokenIndex = position1175, tokenIndex1175
											{
												position1201 := position
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
															goto l1200
														}
														position++
													}
												l1203:
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
															goto l1200
														}
														position++
													}
												l1205:
													{
														position1207, tokenIndex1207 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1208
														}
														position++
														goto l1207
													l1208:
														position, tokenIndex = position1207, tokenIndex1207
														if buffer[position] != rune('D') {
															goto l1200
														}
														position++
													}
												l1207:
													add(rulePegText, position1202)
												}
												{
													add(ruleAction74, position)
												}
												add(ruleRrd, position1201)
											}
											goto l1175
										l1200:
											position, tokenIndex = position1175, tokenIndex1175
											{
												position1211 := position
												{
													position1212 := position
													{
														position1213, tokenIndex1213 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1214
														}
														position++
														goto l1213
													l1214:
														position, tokenIndex = position1213, tokenIndex1213
														if buffer[position] != rune('I') {
															goto l1210
														}
														position++
													}
												l1213:
													{
														position1215, tokenIndex1215 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1216
														}
														position++
														goto l1215
													l1216:
														position, tokenIndex = position1215, tokenIndex1215
														if buffer[position] != rune('M') {
															goto l1210
														}
														position++
													}
												l1215:
													if buffer[position] != rune(' ') {
														goto l1210
													}
													position++
													if buffer[position] != rune('0') {
														goto l1210
													}
													position++
													add(rulePegText, position1212)
												}
												{
													add(ruleAction76, position)
												}
												add(ruleIm0, position1211)
											}
											goto l1175
										l1210:
											position, tokenIndex = position1175, tokenIndex1175
											{
												position1219 := position
												{
													position1220 := position
													{
														position1221, tokenIndex1221 := position, tokenIndex
														if buffer[position] != rune('i') {
															goto l1222
														}
														position++
														goto l1221
													l1222:
														position, tokenIndex = position1221, tokenIndex1221
														if buffer[position] != rune('I') {
															goto l1218
														}
														position++
													}
												l1221:
													{
														position1223, tokenIndex1223 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1224
														}
														position++
														goto l1223
													l1224:
														position, tokenIndex = position1223, tokenIndex1223
														if buffer[position] != rune('M') {
															goto l1218
														}
														position++
													}
												l1223:
													if buffer[position] != rune(' ') {
														goto l1218
													}
													position++
													if buffer[position] != rune('1') {
														goto l1218
													}
													position++
													add(rulePegText, position1220)
												}
												{
													add(ruleAction77, position)
												}
												add(ruleIm1, position1219)
											}
											goto l1175
										l1218:
											position, tokenIndex = position1175, tokenIndex1175
											{
												position1227 := position
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
															goto l1226
														}
														position++
													}
												l1229:
													{
														position1231, tokenIndex1231 := position, tokenIndex
														if buffer[position] != rune('m') {
															goto l1232
														}
														position++
														goto l1231
													l1232:
														position, tokenIndex = position1231, tokenIndex1231
														if buffer[position] != rune('M') {
															goto l1226
														}
														position++
													}
												l1231:
													if buffer[position] != rune(' ') {
														goto l1226
													}
													position++
													if buffer[position] != rune('2') {
														goto l1226
													}
													position++
													add(rulePegText, position1228)
												}
												{
													add(ruleAction78, position)
												}
												add(ruleIm2, position1227)
											}
											goto l1175
										l1226:
											position, tokenIndex = position1175, tokenIndex1175
											{
												switch buffer[position] {
												case 'I', 'O', 'i', 'o':
													{
														position1235 := position
														{
															position1236, tokenIndex1236 := position, tokenIndex
															{
																position1238 := position
																{
																	position1239 := position
																	{
																		position1240, tokenIndex1240 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1241
																		}
																		position++
																		goto l1240
																	l1241:
																		position, tokenIndex = position1240, tokenIndex1240
																		if buffer[position] != rune('I') {
																			goto l1237
																		}
																		position++
																	}
																l1240:
																	{
																		position1242, tokenIndex1242 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1243
																		}
																		position++
																		goto l1242
																	l1243:
																		position, tokenIndex = position1242, tokenIndex1242
																		if buffer[position] != rune('N') {
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
																add(ruleInir, position1238)
															}
															goto l1236
														l1237:
															position, tokenIndex = position1236, tokenIndex1236
															{
																position1250 := position
																{
																	position1251 := position
																	{
																		position1252, tokenIndex1252 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1253
																		}
																		position++
																		goto l1252
																	l1253:
																		position, tokenIndex = position1252, tokenIndex1252
																		if buffer[position] != rune('I') {
																			goto l1249
																		}
																		position++
																	}
																l1252:
																	{
																		position1254, tokenIndex1254 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1255
																		}
																		position++
																		goto l1254
																	l1255:
																		position, tokenIndex = position1254, tokenIndex1254
																		if buffer[position] != rune('N') {
																			goto l1249
																		}
																		position++
																	}
																l1254:
																	{
																		position1256, tokenIndex1256 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1257
																		}
																		position++
																		goto l1256
																	l1257:
																		position, tokenIndex = position1256, tokenIndex1256
																		if buffer[position] != rune('I') {
																			goto l1249
																		}
																		position++
																	}
																l1256:
																	add(rulePegText, position1251)
																}
																{
																	add(ruleAction81, position)
																}
																add(ruleIni, position1250)
															}
															goto l1236
														l1249:
															position, tokenIndex = position1236, tokenIndex1236
															{
																position1260 := position
																{
																	position1261 := position
																	{
																		position1262, tokenIndex1262 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1263
																		}
																		position++
																		goto l1262
																	l1263:
																		position, tokenIndex = position1262, tokenIndex1262
																		if buffer[position] != rune('O') {
																			goto l1259
																		}
																		position++
																	}
																l1262:
																	{
																		position1264, tokenIndex1264 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1265
																		}
																		position++
																		goto l1264
																	l1265:
																		position, tokenIndex = position1264, tokenIndex1264
																		if buffer[position] != rune('T') {
																			goto l1259
																		}
																		position++
																	}
																l1264:
																	{
																		position1266, tokenIndex1266 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1267
																		}
																		position++
																		goto l1266
																	l1267:
																		position, tokenIndex = position1266, tokenIndex1266
																		if buffer[position] != rune('I') {
																			goto l1259
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
																			goto l1259
																		}
																		position++
																	}
																l1268:
																	add(rulePegText, position1261)
																}
																{
																	add(ruleAction90, position)
																}
																add(ruleOtir, position1260)
															}
															goto l1236
														l1259:
															position, tokenIndex = position1236, tokenIndex1236
															{
																position1272 := position
																{
																	position1273 := position
																	{
																		position1274, tokenIndex1274 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1275
																		}
																		position++
																		goto l1274
																	l1275:
																		position, tokenIndex = position1274, tokenIndex1274
																		if buffer[position] != rune('O') {
																			goto l1271
																		}
																		position++
																	}
																l1274:
																	{
																		position1276, tokenIndex1276 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1277
																		}
																		position++
																		goto l1276
																	l1277:
																		position, tokenIndex = position1276, tokenIndex1276
																		if buffer[position] != rune('U') {
																			goto l1271
																		}
																		position++
																	}
																l1276:
																	{
																		position1278, tokenIndex1278 := position, tokenIndex
																		if buffer[position] != rune('t') {
																			goto l1279
																		}
																		position++
																		goto l1278
																	l1279:
																		position, tokenIndex = position1278, tokenIndex1278
																		if buffer[position] != rune('T') {
																			goto l1271
																		}
																		position++
																	}
																l1278:
																	{
																		position1280, tokenIndex1280 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1281
																		}
																		position++
																		goto l1280
																	l1281:
																		position, tokenIndex = position1280, tokenIndex1280
																		if buffer[position] != rune('I') {
																			goto l1271
																		}
																		position++
																	}
																l1280:
																	add(rulePegText, position1273)
																}
																{
																	add(ruleAction82, position)
																}
																add(ruleOuti, position1272)
															}
															goto l1236
														l1271:
															position, tokenIndex = position1236, tokenIndex1236
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
																add(ruleIndr, position1284)
															}
															goto l1236
														l1283:
															position, tokenIndex = position1236, tokenIndex1236
															{
																position1296 := position
																{
																	position1297 := position
																	{
																		position1298, tokenIndex1298 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1299
																		}
																		position++
																		goto l1298
																	l1299:
																		position, tokenIndex = position1298, tokenIndex1298
																		if buffer[position] != rune('I') {
																			goto l1295
																		}
																		position++
																	}
																l1298:
																	{
																		position1300, tokenIndex1300 := position, tokenIndex
																		if buffer[position] != rune('n') {
																			goto l1301
																		}
																		position++
																		goto l1300
																	l1301:
																		position, tokenIndex = position1300, tokenIndex1300
																		if buffer[position] != rune('N') {
																			goto l1295
																		}
																		position++
																	}
																l1300:
																	{
																		position1302, tokenIndex1302 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1303
																		}
																		position++
																		goto l1302
																	l1303:
																		position, tokenIndex = position1302, tokenIndex1302
																		if buffer[position] != rune('D') {
																			goto l1295
																		}
																		position++
																	}
																l1302:
																	add(rulePegText, position1297)
																}
																{
																	add(ruleAction85, position)
																}
																add(ruleInd, position1296)
															}
															goto l1236
														l1295:
															position, tokenIndex = position1236, tokenIndex1236
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
																		if buffer[position] != rune('t') {
																			goto l1311
																		}
																		position++
																		goto l1310
																	l1311:
																		position, tokenIndex = position1310, tokenIndex1310
																		if buffer[position] != rune('T') {
																			goto l1305
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
																			goto l1305
																		}
																		position++
																	}
																l1312:
																	{
																		position1314, tokenIndex1314 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1315
																		}
																		position++
																		goto l1314
																	l1315:
																		position, tokenIndex = position1314, tokenIndex1314
																		if buffer[position] != rune('R') {
																			goto l1305
																		}
																		position++
																	}
																l1314:
																	add(rulePegText, position1307)
																}
																{
																	add(ruleAction94, position)
																}
																add(ruleOtdr, position1306)
															}
															goto l1236
														l1305:
															position, tokenIndex = position1236, tokenIndex1236
															{
																position1317 := position
																{
																	position1318 := position
																	{
																		position1319, tokenIndex1319 := position, tokenIndex
																		if buffer[position] != rune('o') {
																			goto l1320
																		}
																		position++
																		goto l1319
																	l1320:
																		position, tokenIndex = position1319, tokenIndex1319
																		if buffer[position] != rune('O') {
																			goto l1173
																		}
																		position++
																	}
																l1319:
																	{
																		position1321, tokenIndex1321 := position, tokenIndex
																		if buffer[position] != rune('u') {
																			goto l1322
																		}
																		position++
																		goto l1321
																	l1322:
																		position, tokenIndex = position1321, tokenIndex1321
																		if buffer[position] != rune('U') {
																			goto l1173
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
																			goto l1173
																		}
																		position++
																	}
																l1323:
																	{
																		position1325, tokenIndex1325 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1326
																		}
																		position++
																		goto l1325
																	l1326:
																		position, tokenIndex = position1325, tokenIndex1325
																		if buffer[position] != rune('D') {
																			goto l1173
																		}
																		position++
																	}
																l1325:
																	add(rulePegText, position1318)
																}
																{
																	add(ruleAction86, position)
																}
																add(ruleOutd, position1317)
															}
														}
													l1236:
														add(ruleBlitIO, position1235)
													}
													break
												case 'R', 'r':
													{
														position1328 := position
														{
															position1329 := position
															{
																position1330, tokenIndex1330 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1331
																}
																position++
																goto l1330
															l1331:
																position, tokenIndex = position1330, tokenIndex1330
																if buffer[position] != rune('R') {
																	goto l1173
																}
																position++
															}
														l1330:
															{
																position1332, tokenIndex1332 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1333
																}
																position++
																goto l1332
															l1333:
																position, tokenIndex = position1332, tokenIndex1332
																if buffer[position] != rune('L') {
																	goto l1173
																}
																position++
															}
														l1332:
															{
																position1334, tokenIndex1334 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1335
																}
																position++
																goto l1334
															l1335:
																position, tokenIndex = position1334, tokenIndex1334
																if buffer[position] != rune('D') {
																	goto l1173
																}
																position++
															}
														l1334:
															add(rulePegText, position1329)
														}
														{
															add(ruleAction75, position)
														}
														add(ruleRld, position1328)
													}
													break
												case 'N', 'n':
													{
														position1337 := position
														{
															position1338 := position
															{
																position1339, tokenIndex1339 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1340
																}
																position++
																goto l1339
															l1340:
																position, tokenIndex = position1339, tokenIndex1339
																if buffer[position] != rune('N') {
																	goto l1173
																}
																position++
															}
														l1339:
															{
																position1341, tokenIndex1341 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1342
																}
																position++
																goto l1341
															l1342:
																position, tokenIndex = position1341, tokenIndex1341
																if buffer[position] != rune('E') {
																	goto l1173
																}
																position++
															}
														l1341:
															{
																position1343, tokenIndex1343 := position, tokenIndex
																if buffer[position] != rune('g') {
																	goto l1344
																}
																position++
																goto l1343
															l1344:
																position, tokenIndex = position1343, tokenIndex1343
																if buffer[position] != rune('G') {
																	goto l1173
																}
																position++
															}
														l1343:
															add(rulePegText, position1338)
														}
														{
															add(ruleAction71, position)
														}
														add(ruleNeg, position1337)
													}
													break
												default:
													{
														position1346 := position
														{
															position1347, tokenIndex1347 := position, tokenIndex
															{
																position1349 := position
																{
																	position1350 := position
																	{
																		position1351, tokenIndex1351 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1352
																		}
																		position++
																		goto l1351
																	l1352:
																		position, tokenIndex = position1351, tokenIndex1351
																		if buffer[position] != rune('L') {
																			goto l1348
																		}
																		position++
																	}
																l1351:
																	{
																		position1353, tokenIndex1353 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1354
																		}
																		position++
																		goto l1353
																	l1354:
																		position, tokenIndex = position1353, tokenIndex1353
																		if buffer[position] != rune('D') {
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
																add(ruleLdir, position1349)
															}
															goto l1347
														l1348:
															position, tokenIndex = position1347, tokenIndex1347
															{
																position1361 := position
																{
																	position1362 := position
																	{
																		position1363, tokenIndex1363 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1364
																		}
																		position++
																		goto l1363
																	l1364:
																		position, tokenIndex = position1363, tokenIndex1363
																		if buffer[position] != rune('L') {
																			goto l1360
																		}
																		position++
																	}
																l1363:
																	{
																		position1365, tokenIndex1365 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1366
																		}
																		position++
																		goto l1365
																	l1366:
																		position, tokenIndex = position1365, tokenIndex1365
																		if buffer[position] != rune('D') {
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
																add(ruleLdi, position1361)
															}
															goto l1347
														l1360:
															position, tokenIndex = position1347, tokenIndex1347
															{
																position1371 := position
																{
																	position1372 := position
																	{
																		position1373, tokenIndex1373 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1374
																		}
																		position++
																		goto l1373
																	l1374:
																		position, tokenIndex = position1373, tokenIndex1373
																		if buffer[position] != rune('C') {
																			goto l1370
																		}
																		position++
																	}
																l1373:
																	{
																		position1375, tokenIndex1375 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1376
																		}
																		position++
																		goto l1375
																	l1376:
																		position, tokenIndex = position1375, tokenIndex1375
																		if buffer[position] != rune('P') {
																			goto l1370
																		}
																		position++
																	}
																l1375:
																	{
																		position1377, tokenIndex1377 := position, tokenIndex
																		if buffer[position] != rune('i') {
																			goto l1378
																		}
																		position++
																		goto l1377
																	l1378:
																		position, tokenIndex = position1377, tokenIndex1377
																		if buffer[position] != rune('I') {
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
																	add(ruleAction88, position)
																}
																add(ruleCpir, position1371)
															}
															goto l1347
														l1370:
															position, tokenIndex = position1347, tokenIndex1347
															{
																position1383 := position
																{
																	position1384 := position
																	{
																		position1385, tokenIndex1385 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1386
																		}
																		position++
																		goto l1385
																	l1386:
																		position, tokenIndex = position1385, tokenIndex1385
																		if buffer[position] != rune('C') {
																			goto l1382
																		}
																		position++
																	}
																l1385:
																	{
																		position1387, tokenIndex1387 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1388
																		}
																		position++
																		goto l1387
																	l1388:
																		position, tokenIndex = position1387, tokenIndex1387
																		if buffer[position] != rune('P') {
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
																	add(rulePegText, position1384)
																}
																{
																	add(ruleAction80, position)
																}
																add(ruleCpi, position1383)
															}
															goto l1347
														l1382:
															position, tokenIndex = position1347, tokenIndex1347
															{
																position1393 := position
																{
																	position1394 := position
																	{
																		position1395, tokenIndex1395 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1396
																		}
																		position++
																		goto l1395
																	l1396:
																		position, tokenIndex = position1395, tokenIndex1395
																		if buffer[position] != rune('L') {
																			goto l1392
																		}
																		position++
																	}
																l1395:
																	{
																		position1397, tokenIndex1397 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1398
																		}
																		position++
																		goto l1397
																	l1398:
																		position, tokenIndex = position1397, tokenIndex1397
																		if buffer[position] != rune('D') {
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
																add(ruleLddr, position1393)
															}
															goto l1347
														l1392:
															position, tokenIndex = position1347, tokenIndex1347
															{
																position1405 := position
																{
																	position1406 := position
																	{
																		position1407, tokenIndex1407 := position, tokenIndex
																		if buffer[position] != rune('l') {
																			goto l1408
																		}
																		position++
																		goto l1407
																	l1408:
																		position, tokenIndex = position1407, tokenIndex1407
																		if buffer[position] != rune('L') {
																			goto l1404
																		}
																		position++
																	}
																l1407:
																	{
																		position1409, tokenIndex1409 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1410
																		}
																		position++
																		goto l1409
																	l1410:
																		position, tokenIndex = position1409, tokenIndex1409
																		if buffer[position] != rune('D') {
																			goto l1404
																		}
																		position++
																	}
																l1409:
																	{
																		position1411, tokenIndex1411 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1412
																		}
																		position++
																		goto l1411
																	l1412:
																		position, tokenIndex = position1411, tokenIndex1411
																		if buffer[position] != rune('D') {
																			goto l1404
																		}
																		position++
																	}
																l1411:
																	add(rulePegText, position1406)
																}
																{
																	add(ruleAction83, position)
																}
																add(ruleLdd, position1405)
															}
															goto l1347
														l1404:
															position, tokenIndex = position1347, tokenIndex1347
															{
																position1415 := position
																{
																	position1416 := position
																	{
																		position1417, tokenIndex1417 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1418
																		}
																		position++
																		goto l1417
																	l1418:
																		position, tokenIndex = position1417, tokenIndex1417
																		if buffer[position] != rune('C') {
																			goto l1414
																		}
																		position++
																	}
																l1417:
																	{
																		position1419, tokenIndex1419 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1420
																		}
																		position++
																		goto l1419
																	l1420:
																		position, tokenIndex = position1419, tokenIndex1419
																		if buffer[position] != rune('P') {
																			goto l1414
																		}
																		position++
																	}
																l1419:
																	{
																		position1421, tokenIndex1421 := position, tokenIndex
																		if buffer[position] != rune('d') {
																			goto l1422
																		}
																		position++
																		goto l1421
																	l1422:
																		position, tokenIndex = position1421, tokenIndex1421
																		if buffer[position] != rune('D') {
																			goto l1414
																		}
																		position++
																	}
																l1421:
																	{
																		position1423, tokenIndex1423 := position, tokenIndex
																		if buffer[position] != rune('r') {
																			goto l1424
																		}
																		position++
																		goto l1423
																	l1424:
																		position, tokenIndex = position1423, tokenIndex1423
																		if buffer[position] != rune('R') {
																			goto l1414
																		}
																		position++
																	}
																l1423:
																	add(rulePegText, position1416)
																}
																{
																	add(ruleAction92, position)
																}
																add(ruleCpdr, position1415)
															}
															goto l1347
														l1414:
															position, tokenIndex = position1347, tokenIndex1347
															{
																position1426 := position
																{
																	position1427 := position
																	{
																		position1428, tokenIndex1428 := position, tokenIndex
																		if buffer[position] != rune('c') {
																			goto l1429
																		}
																		position++
																		goto l1428
																	l1429:
																		position, tokenIndex = position1428, tokenIndex1428
																		if buffer[position] != rune('C') {
																			goto l1173
																		}
																		position++
																	}
																l1428:
																	{
																		position1430, tokenIndex1430 := position, tokenIndex
																		if buffer[position] != rune('p') {
																			goto l1431
																		}
																		position++
																		goto l1430
																	l1431:
																		position, tokenIndex = position1430, tokenIndex1430
																		if buffer[position] != rune('P') {
																			goto l1173
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
																			goto l1173
																		}
																		position++
																	}
																l1432:
																	add(rulePegText, position1427)
																}
																{
																	add(ruleAction84, position)
																}
																add(ruleCpd, position1426)
															}
														}
													l1347:
														add(ruleBlit, position1346)
													}
													break
												}
											}

										}
									l1175:
										add(ruleEDSimple, position1174)
									}
									goto l845
								l1173:
									position, tokenIndex = position845, tokenIndex845
									{
										position1436 := position
										{
											position1437, tokenIndex1437 := position, tokenIndex
											{
												position1439 := position
												{
													position1440 := position
													{
														position1441, tokenIndex1441 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1442
														}
														position++
														goto l1441
													l1442:
														position, tokenIndex = position1441, tokenIndex1441
														if buffer[position] != rune('R') {
															goto l1438
														}
														position++
													}
												l1441:
													{
														position1443, tokenIndex1443 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1444
														}
														position++
														goto l1443
													l1444:
														position, tokenIndex = position1443, tokenIndex1443
														if buffer[position] != rune('L') {
															goto l1438
														}
														position++
													}
												l1443:
													{
														position1445, tokenIndex1445 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1446
														}
														position++
														goto l1445
													l1446:
														position, tokenIndex = position1445, tokenIndex1445
														if buffer[position] != rune('C') {
															goto l1438
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
															goto l1438
														}
														position++
													}
												l1447:
													add(rulePegText, position1440)
												}
												{
													add(ruleAction60, position)
												}
												add(ruleRlca, position1439)
											}
											goto l1437
										l1438:
											position, tokenIndex = position1437, tokenIndex1437
											{
												position1451 := position
												{
													position1452 := position
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
															goto l1450
														}
														position++
													}
												l1453:
													{
														position1455, tokenIndex1455 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1456
														}
														position++
														goto l1455
													l1456:
														position, tokenIndex = position1455, tokenIndex1455
														if buffer[position] != rune('R') {
															goto l1450
														}
														position++
													}
												l1455:
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
															goto l1450
														}
														position++
													}
												l1457:
													{
														position1459, tokenIndex1459 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1460
														}
														position++
														goto l1459
													l1460:
														position, tokenIndex = position1459, tokenIndex1459
														if buffer[position] != rune('A') {
															goto l1450
														}
														position++
													}
												l1459:
													add(rulePegText, position1452)
												}
												{
													add(ruleAction61, position)
												}
												add(ruleRrca, position1451)
											}
											goto l1437
										l1450:
											position, tokenIndex = position1437, tokenIndex1437
											{
												position1463 := position
												{
													position1464 := position
													{
														position1465, tokenIndex1465 := position, tokenIndex
														if buffer[position] != rune('r') {
															goto l1466
														}
														position++
														goto l1465
													l1466:
														position, tokenIndex = position1465, tokenIndex1465
														if buffer[position] != rune('R') {
															goto l1462
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
															goto l1462
														}
														position++
													}
												l1467:
													{
														position1469, tokenIndex1469 := position, tokenIndex
														if buffer[position] != rune('a') {
															goto l1470
														}
														position++
														goto l1469
													l1470:
														position, tokenIndex = position1469, tokenIndex1469
														if buffer[position] != rune('A') {
															goto l1462
														}
														position++
													}
												l1469:
													add(rulePegText, position1464)
												}
												{
													add(ruleAction62, position)
												}
												add(ruleRla, position1463)
											}
											goto l1437
										l1462:
											position, tokenIndex = position1437, tokenIndex1437
											{
												position1473 := position
												{
													position1474 := position
													{
														position1475, tokenIndex1475 := position, tokenIndex
														if buffer[position] != rune('d') {
															goto l1476
														}
														position++
														goto l1475
													l1476:
														position, tokenIndex = position1475, tokenIndex1475
														if buffer[position] != rune('D') {
															goto l1472
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
															goto l1472
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
															goto l1472
														}
														position++
													}
												l1479:
													add(rulePegText, position1474)
												}
												{
													add(ruleAction64, position)
												}
												add(ruleDaa, position1473)
											}
											goto l1437
										l1472:
											position, tokenIndex = position1437, tokenIndex1437
											{
												position1483 := position
												{
													position1484 := position
													{
														position1485, tokenIndex1485 := position, tokenIndex
														if buffer[position] != rune('c') {
															goto l1486
														}
														position++
														goto l1485
													l1486:
														position, tokenIndex = position1485, tokenIndex1485
														if buffer[position] != rune('C') {
															goto l1482
														}
														position++
													}
												l1485:
													{
														position1487, tokenIndex1487 := position, tokenIndex
														if buffer[position] != rune('p') {
															goto l1488
														}
														position++
														goto l1487
													l1488:
														position, tokenIndex = position1487, tokenIndex1487
														if buffer[position] != rune('P') {
															goto l1482
														}
														position++
													}
												l1487:
													{
														position1489, tokenIndex1489 := position, tokenIndex
														if buffer[position] != rune('l') {
															goto l1490
														}
														position++
														goto l1489
													l1490:
														position, tokenIndex = position1489, tokenIndex1489
														if buffer[position] != rune('L') {
															goto l1482
														}
														position++
													}
												l1489:
													add(rulePegText, position1484)
												}
												{
													add(ruleAction65, position)
												}
												add(ruleCpl, position1483)
											}
											goto l1437
										l1482:
											position, tokenIndex = position1437, tokenIndex1437
											{
												position1493 := position
												{
													position1494 := position
													{
														position1495, tokenIndex1495 := position, tokenIndex
														if buffer[position] != rune('e') {
															goto l1496
														}
														position++
														goto l1495
													l1496:
														position, tokenIndex = position1495, tokenIndex1495
														if buffer[position] != rune('E') {
															goto l1492
														}
														position++
													}
												l1495:
													{
														position1497, tokenIndex1497 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1498
														}
														position++
														goto l1497
													l1498:
														position, tokenIndex = position1497, tokenIndex1497
														if buffer[position] != rune('X') {
															goto l1492
														}
														position++
													}
												l1497:
													{
														position1499, tokenIndex1499 := position, tokenIndex
														if buffer[position] != rune('x') {
															goto l1500
														}
														position++
														goto l1499
													l1500:
														position, tokenIndex = position1499, tokenIndex1499
														if buffer[position] != rune('X') {
															goto l1492
														}
														position++
													}
												l1499:
													add(rulePegText, position1494)
												}
												{
													add(ruleAction68, position)
												}
												add(ruleExx, position1493)
											}
											goto l1437
										l1492:
											position, tokenIndex = position1437, tokenIndex1437
											{
												switch buffer[position] {
												case 'E', 'e':
													{
														position1503 := position
														{
															position1504 := position
															{
																position1505, tokenIndex1505 := position, tokenIndex
																if buffer[position] != rune('e') {
																	goto l1506
																}
																position++
																goto l1505
															l1506:
																position, tokenIndex = position1505, tokenIndex1505
																if buffer[position] != rune('E') {
																	goto l1435
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
																	goto l1435
																}
																position++
															}
														l1507:
															add(rulePegText, position1504)
														}
														{
															add(ruleAction70, position)
														}
														add(ruleEi, position1503)
													}
													break
												case 'D', 'd':
													{
														position1510 := position
														{
															position1511 := position
															{
																position1512, tokenIndex1512 := position, tokenIndex
																if buffer[position] != rune('d') {
																	goto l1513
																}
																position++
																goto l1512
															l1513:
																position, tokenIndex = position1512, tokenIndex1512
																if buffer[position] != rune('D') {
																	goto l1435
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
																	goto l1435
																}
																position++
															}
														l1514:
															add(rulePegText, position1511)
														}
														{
															add(ruleAction69, position)
														}
														add(ruleDi, position1510)
													}
													break
												case 'C', 'c':
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
																	goto l1435
																}
																position++
															}
														l1519:
															{
																position1521, tokenIndex1521 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1522
																}
																position++
																goto l1521
															l1522:
																position, tokenIndex = position1521, tokenIndex1521
																if buffer[position] != rune('C') {
																	goto l1435
																}
																position++
															}
														l1521:
															{
																position1523, tokenIndex1523 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1524
																}
																position++
																goto l1523
															l1524:
																position, tokenIndex = position1523, tokenIndex1523
																if buffer[position] != rune('F') {
																	goto l1435
																}
																position++
															}
														l1523:
															add(rulePegText, position1518)
														}
														{
															add(ruleAction67, position)
														}
														add(ruleCcf, position1517)
													}
													break
												case 'S', 's':
													{
														position1526 := position
														{
															position1527 := position
															{
																position1528, tokenIndex1528 := position, tokenIndex
																if buffer[position] != rune('s') {
																	goto l1529
																}
																position++
																goto l1528
															l1529:
																position, tokenIndex = position1528, tokenIndex1528
																if buffer[position] != rune('S') {
																	goto l1435
																}
																position++
															}
														l1528:
															{
																position1530, tokenIndex1530 := position, tokenIndex
																if buffer[position] != rune('c') {
																	goto l1531
																}
																position++
																goto l1530
															l1531:
																position, tokenIndex = position1530, tokenIndex1530
																if buffer[position] != rune('C') {
																	goto l1435
																}
																position++
															}
														l1530:
															{
																position1532, tokenIndex1532 := position, tokenIndex
																if buffer[position] != rune('f') {
																	goto l1533
																}
																position++
																goto l1532
															l1533:
																position, tokenIndex = position1532, tokenIndex1532
																if buffer[position] != rune('F') {
																	goto l1435
																}
																position++
															}
														l1532:
															add(rulePegText, position1527)
														}
														{
															add(ruleAction66, position)
														}
														add(ruleScf, position1526)
													}
													break
												case 'R', 'r':
													{
														position1535 := position
														{
															position1536 := position
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
																	goto l1435
																}
																position++
															}
														l1537:
															{
																position1539, tokenIndex1539 := position, tokenIndex
																if buffer[position] != rune('r') {
																	goto l1540
																}
																position++
																goto l1539
															l1540:
																position, tokenIndex = position1539, tokenIndex1539
																if buffer[position] != rune('R') {
																	goto l1435
																}
																position++
															}
														l1539:
															{
																position1541, tokenIndex1541 := position, tokenIndex
																if buffer[position] != rune('a') {
																	goto l1542
																}
																position++
																goto l1541
															l1542:
																position, tokenIndex = position1541, tokenIndex1541
																if buffer[position] != rune('A') {
																	goto l1435
																}
																position++
															}
														l1541:
															add(rulePegText, position1536)
														}
														{
															add(ruleAction63, position)
														}
														add(ruleRra, position1535)
													}
													break
												case 'H', 'h':
													{
														position1544 := position
														{
															position1545 := position
															{
																position1546, tokenIndex1546 := position, tokenIndex
																if buffer[position] != rune('h') {
																	goto l1547
																}
																position++
																goto l1546
															l1547:
																position, tokenIndex = position1546, tokenIndex1546
																if buffer[position] != rune('H') {
																	goto l1435
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
																	goto l1435
																}
																position++
															}
														l1548:
															{
																position1550, tokenIndex1550 := position, tokenIndex
																if buffer[position] != rune('l') {
																	goto l1551
																}
																position++
																goto l1550
															l1551:
																position, tokenIndex = position1550, tokenIndex1550
																if buffer[position] != rune('L') {
																	goto l1435
																}
																position++
															}
														l1550:
															{
																position1552, tokenIndex1552 := position, tokenIndex
																if buffer[position] != rune('t') {
																	goto l1553
																}
																position++
																goto l1552
															l1553:
																position, tokenIndex = position1552, tokenIndex1552
																if buffer[position] != rune('T') {
																	goto l1435
																}
																position++
															}
														l1552:
															add(rulePegText, position1545)
														}
														{
															add(ruleAction59, position)
														}
														add(ruleHalt, position1544)
													}
													break
												default:
													{
														position1555 := position
														{
															position1556 := position
															{
																position1557, tokenIndex1557 := position, tokenIndex
																if buffer[position] != rune('n') {
																	goto l1558
																}
																position++
																goto l1557
															l1558:
																position, tokenIndex = position1557, tokenIndex1557
																if buffer[position] != rune('N') {
																	goto l1435
																}
																position++
															}
														l1557:
															{
																position1559, tokenIndex1559 := position, tokenIndex
																if buffer[position] != rune('o') {
																	goto l1560
																}
																position++
																goto l1559
															l1560:
																position, tokenIndex = position1559, tokenIndex1559
																if buffer[position] != rune('O') {
																	goto l1435
																}
																position++
															}
														l1559:
															{
																position1561, tokenIndex1561 := position, tokenIndex
																if buffer[position] != rune('p') {
																	goto l1562
																}
																position++
																goto l1561
															l1562:
																position, tokenIndex = position1561, tokenIndex1561
																if buffer[position] != rune('P') {
																	goto l1435
																}
																position++
															}
														l1561:
															add(rulePegText, position1556)
														}
														{
															add(ruleAction58, position)
														}
														add(ruleNop, position1555)
													}
													break
												}
											}

										}
									l1437:
										add(ruleSimple, position1436)
									}
									goto l845
								l1435:
									position, tokenIndex = position845, tokenIndex845
									{
										position1565 := position
										{
											position1566, tokenIndex1566 := position, tokenIndex
											{
												position1568 := position
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
														goto l1567
													}
													position++
												}
											l1569:
												{
													position1571, tokenIndex1571 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1572
													}
													position++
													goto l1571
												l1572:
													position, tokenIndex = position1571, tokenIndex1571
													if buffer[position] != rune('S') {
														goto l1567
													}
													position++
												}
											l1571:
												{
													position1573, tokenIndex1573 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1574
													}
													position++
													goto l1573
												l1574:
													position, tokenIndex = position1573, tokenIndex1573
													if buffer[position] != rune('T') {
														goto l1567
													}
													position++
												}
											l1573:
												if !_rules[rulews]() {
													goto l1567
												}
												if !_rules[rulen]() {
													goto l1567
												}
												{
													add(ruleAction95, position)
												}
												add(ruleRst, position1568)
											}
											goto l1566
										l1567:
											position, tokenIndex = position1566, tokenIndex1566
											{
												position1577 := position
												{
													position1578, tokenIndex1578 := position, tokenIndex
													if buffer[position] != rune('j') {
														goto l1579
													}
													position++
													goto l1578
												l1579:
													position, tokenIndex = position1578, tokenIndex1578
													if buffer[position] != rune('J') {
														goto l1576
													}
													position++
												}
											l1578:
												{
													position1580, tokenIndex1580 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l1581
													}
													position++
													goto l1580
												l1581:
													position, tokenIndex = position1580, tokenIndex1580
													if buffer[position] != rune('P') {
														goto l1576
													}
													position++
												}
											l1580:
												if !_rules[rulews]() {
													goto l1576
												}
												{
													position1582, tokenIndex1582 := position, tokenIndex
													if !_rules[rulecc]() {
														goto l1582
													}
													if !_rules[rulesep]() {
														goto l1582
													}
													goto l1583
												l1582:
													position, tokenIndex = position1582, tokenIndex1582
												}
											l1583:
												if !_rules[ruleSrc16]() {
													goto l1576
												}
												{
													add(ruleAction98, position)
												}
												add(ruleJp, position1577)
											}
											goto l1566
										l1576:
											position, tokenIndex = position1566, tokenIndex1566
											{
												switch buffer[position] {
												case 'D', 'd':
													{
														position1586 := position
														{
															position1587, tokenIndex1587 := position, tokenIndex
															if buffer[position] != rune('d') {
																goto l1588
															}
															position++
															goto l1587
														l1588:
															position, tokenIndex = position1587, tokenIndex1587
															if buffer[position] != rune('D') {
																goto l1564
															}
															position++
														}
													l1587:
														{
															position1589, tokenIndex1589 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1590
															}
															position++
															goto l1589
														l1590:
															position, tokenIndex = position1589, tokenIndex1589
															if buffer[position] != rune('J') {
																goto l1564
															}
															position++
														}
													l1589:
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
																goto l1564
															}
															position++
														}
													l1591:
														{
															position1593, tokenIndex1593 := position, tokenIndex
															if buffer[position] != rune('z') {
																goto l1594
															}
															position++
															goto l1593
														l1594:
															position, tokenIndex = position1593, tokenIndex1593
															if buffer[position] != rune('Z') {
																goto l1564
															}
															position++
														}
													l1593:
														if !_rules[rulews]() {
															goto l1564
														}
														if !_rules[ruledisp]() {
															goto l1564
														}
														{
															add(ruleAction100, position)
														}
														add(ruleDjnz, position1586)
													}
													break
												case 'J', 'j':
													{
														position1596 := position
														{
															position1597, tokenIndex1597 := position, tokenIndex
															if buffer[position] != rune('j') {
																goto l1598
															}
															position++
															goto l1597
														l1598:
															position, tokenIndex = position1597, tokenIndex1597
															if buffer[position] != rune('J') {
																goto l1564
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
																goto l1564
															}
															position++
														}
													l1599:
														if !_rules[rulews]() {
															goto l1564
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
														if !_rules[ruledisp]() {
															goto l1564
														}
														{
															add(ruleAction99, position)
														}
														add(ruleJr, position1596)
													}
													break
												case 'R', 'r':
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
																goto l1564
															}
															position++
														}
													l1605:
														{
															position1607, tokenIndex1607 := position, tokenIndex
															if buffer[position] != rune('e') {
																goto l1608
															}
															position++
															goto l1607
														l1608:
															position, tokenIndex = position1607, tokenIndex1607
															if buffer[position] != rune('E') {
																goto l1564
															}
															position++
														}
													l1607:
														{
															position1609, tokenIndex1609 := position, tokenIndex
															if buffer[position] != rune('t') {
																goto l1610
															}
															position++
															goto l1609
														l1610:
															position, tokenIndex = position1609, tokenIndex1609
															if buffer[position] != rune('T') {
																goto l1564
															}
															position++
														}
													l1609:
														{
															position1611, tokenIndex1611 := position, tokenIndex
															if !_rules[rulews]() {
																goto l1611
															}
															if !_rules[rulecc]() {
																goto l1611
															}
															goto l1612
														l1611:
															position, tokenIndex = position1611, tokenIndex1611
														}
													l1612:
														{
															add(ruleAction97, position)
														}
														add(ruleRet, position1604)
													}
													break
												default:
													{
														position1614 := position
														{
															position1615, tokenIndex1615 := position, tokenIndex
															if buffer[position] != rune('c') {
																goto l1616
															}
															position++
															goto l1615
														l1616:
															position, tokenIndex = position1615, tokenIndex1615
															if buffer[position] != rune('C') {
																goto l1564
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
																goto l1564
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
																goto l1564
															}
															position++
														}
													l1619:
														{
															position1621, tokenIndex1621 := position, tokenIndex
															if buffer[position] != rune('l') {
																goto l1622
															}
															position++
															goto l1621
														l1622:
															position, tokenIndex = position1621, tokenIndex1621
															if buffer[position] != rune('L') {
																goto l1564
															}
															position++
														}
													l1621:
														if !_rules[rulews]() {
															goto l1564
														}
														{
															position1623, tokenIndex1623 := position, tokenIndex
															if !_rules[rulecc]() {
																goto l1623
															}
															if !_rules[rulesep]() {
																goto l1623
															}
															goto l1624
														l1623:
															position, tokenIndex = position1623, tokenIndex1623
														}
													l1624:
														if !_rules[ruleSrc16]() {
															goto l1564
														}
														{
															add(ruleAction96, position)
														}
														add(ruleCall, position1614)
													}
													break
												}
											}

										}
									l1566:
										add(ruleJump, position1565)
									}
									goto l845
								l1564:
									position, tokenIndex = position845, tokenIndex845
									{
										position1626 := position
										{
											position1627, tokenIndex1627 := position, tokenIndex
											{
												position1629 := position
												{
													position1630, tokenIndex1630 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1631
													}
													position++
													goto l1630
												l1631:
													position, tokenIndex = position1630, tokenIndex1630
													if buffer[position] != rune('I') {
														goto l1628
													}
													position++
												}
											l1630:
												{
													position1632, tokenIndex1632 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1633
													}
													position++
													goto l1632
												l1633:
													position, tokenIndex = position1632, tokenIndex1632
													if buffer[position] != rune('N') {
														goto l1628
													}
													position++
												}
											l1632:
												if !_rules[rulews]() {
													goto l1628
												}
												if !_rules[ruleReg8]() {
													goto l1628
												}
												if !_rules[rulesep]() {
													goto l1628
												}
												if !_rules[rulePort]() {
													goto l1628
												}
												{
													add(ruleAction101, position)
												}
												add(ruleIN, position1629)
											}
											goto l1627
										l1628:
											position, tokenIndex = position1627, tokenIndex1627
											{
												position1635 := position
												{
													position1636, tokenIndex1636 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1637
													}
													position++
													goto l1636
												l1637:
													position, tokenIndex = position1636, tokenIndex1636
													if buffer[position] != rune('O') {
														goto l3
													}
													position++
												}
											l1636:
												{
													position1638, tokenIndex1638 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1639
													}
													position++
													goto l1638
												l1639:
													position, tokenIndex = position1638, tokenIndex1638
													if buffer[position] != rune('U') {
														goto l3
													}
													position++
												}
											l1638:
												{
													position1640, tokenIndex1640 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1641
													}
													position++
													goto l1640
												l1641:
													position, tokenIndex = position1640, tokenIndex1640
													if buffer[position] != rune('T') {
														goto l3
													}
													position++
												}
											l1640:
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
													add(ruleAction102, position)
												}
												add(ruleOUT, position1635)
											}
										}
									l1627:
										add(ruleIO, position1626)
									}
								}
							l845:
								add(ruleInstruction, position844)
							}
							{
								position1643, tokenIndex1643 := position, tokenIndex
								if !_rules[ruleLineEnd]() {
									goto l1643
								}
								goto l1644
							l1643:
								position, tokenIndex = position1643, tokenIndex1643
							}
						l1644:
							{
								add(ruleAction0, position)
							}
							add(ruleLine, position830)
						}
					}
				l825:
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
		/* 2 Line <- <(Label? ws* Instruction LineEnd? Action0)> */
		nil,
		/* 3 Label <- <(<(alpha alphanum*)> ':' Action1)> */
		nil,
		/* 4 alphanum <- <(alpha / num)> */
		nil,
		/* 5 alpha <- <([a-z] / [A-Z])> */
		func() bool {
			position1650, tokenIndex1650 := position, tokenIndex
			{
				position1651 := position
				{
					position1652, tokenIndex1652 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1653
					}
					position++
					goto l1652
				l1653:
					position, tokenIndex = position1652, tokenIndex1652
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1650
					}
					position++
				}
			l1652:
				add(rulealpha, position1651)
			}
			return true
		l1650:
			position, tokenIndex = position1650, tokenIndex1650
			return false
		},
		/* 6 num <- <[0-9]> */
		nil,
		/* 7 Comment <- <((';' / '#') .*)> */
		nil,
		/* 8 LineEnd <- <(Comment? ('\n' / ':'))> */
		func() bool {
			position1656, tokenIndex1656 := position, tokenIndex
			{
				position1657 := position
				{
					position1658, tokenIndex1658 := position, tokenIndex
					{
						position1660 := position
						{
							position1661, tokenIndex1661 := position, tokenIndex
							if buffer[position] != rune(';') {
								goto l1662
							}
							position++
							goto l1661
						l1662:
							position, tokenIndex = position1661, tokenIndex1661
							if buffer[position] != rune('#') {
								goto l1658
							}
							position++
						}
					l1661:
					l1663:
						{
							position1664, tokenIndex1664 := position, tokenIndex
							if !matchDot() {
								goto l1664
							}
							goto l1663
						l1664:
							position, tokenIndex = position1664, tokenIndex1664
						}
						add(ruleComment, position1660)
					}
					goto l1659
				l1658:
					position, tokenIndex = position1658, tokenIndex1658
				}
			l1659:
				{
					position1665, tokenIndex1665 := position, tokenIndex
					if buffer[position] != rune('\n') {
						goto l1666
					}
					position++
					goto l1665
				l1666:
					position, tokenIndex = position1665, tokenIndex1665
					if buffer[position] != rune(':') {
						goto l1656
					}
					position++
				}
			l1665:
				add(ruleLineEnd, position1657)
			}
			return true
		l1656:
			position, tokenIndex = position1656, tokenIndex1656
			return false
		},
		/* 9 Instruction <- <(Assignment / Inc / Dec / Alu16 / Alu / BitOp / EDSimple / Simple / Jump / IO)> */
		nil,
		/* 10 Assignment <- <(Push / ((&('E' | 'e') Ex) | (&('P' | 'p') Pop) | (&('L' | 'l') Load)))> */
		nil,
		/* 11 Load <- <(Load16 / Load8)> */
		nil,
		/* 12 Load8 <- <(('l' / 'L') ('d' / 'D') ws Dst8 sep Src8 Action2)> */
		nil,
		/* 13 Load16 <- <(('l' / 'L') ('d' / 'D') ws Dst16 sep Src16 Action3)> */
		nil,
		/* 14 Push <- <(('p' / 'P') ('u' / 'U') ('s' / 'S') ('h' / 'H') ws Src16 Action4)> */
		nil,
		/* 15 Pop <- <(('p' / 'P') ('o' / 'O') ('p' / 'P') ws Dst16 Action5)> */
		nil,
		/* 16 Ex <- <(('e' / 'E') ('x' / 'X') ws Dst16 sep Src16 Action6)> */
		nil,
		/* 17 Inc <- <(Inc16Indexed8 / Inc16 / Inc8)> */
		nil,
		/* 18 Inc16Indexed8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws ILoc8 Action7)> */
		nil,
		/* 19 Inc8 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc8 Action8)> */
		nil,
		/* 20 Inc16 <- <(('i' / 'I') ('n' / 'N') ('c' / 'C') ws Loc16 Action9)> */
		nil,
		/* 21 Dec <- <(Dec16Indexed8 / Dec16 / Dec8)> */
		nil,
		/* 22 Dec16Indexed8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws ILoc8 Action10)> */
		nil,
		/* 23 Dec8 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc8 Action11)> */
		nil,
		/* 24 Dec16 <- <(('d' / 'D') ('e' / 'E') ('c' / 'C') ws Loc16 Action12)> */
		nil,
		/* 25 Alu16 <- <(Add16 / Adc16 / Sbc16)> */
		nil,
		/* 26 Add16 <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws Dst16 sep Src16 Action13)> */
		nil,
		/* 27 Adc16 <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws Dst16 sep Src16 Action14)> */
		nil,
		/* 28 Sbc16 <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws Dst16 sep Src16 Action15)> */
		nil,
		/* 29 Dst8 <- <((Reg8 / Reg16Contents / nn_contents) Action16)> */
		nil,
		/* 30 Src8 <- <((n / Reg8 / Reg16Contents / nn_contents) Action17)> */
		func() bool {
			position1688, tokenIndex1688 := position, tokenIndex
			{
				position1689 := position
				{
					position1690, tokenIndex1690 := position, tokenIndex
					if !_rules[rulen]() {
						goto l1691
					}
					goto l1690
				l1691:
					position, tokenIndex = position1690, tokenIndex1690
					if !_rules[ruleReg8]() {
						goto l1692
					}
					goto l1690
				l1692:
					position, tokenIndex = position1690, tokenIndex1690
					if !_rules[ruleReg16Contents]() {
						goto l1693
					}
					goto l1690
				l1693:
					position, tokenIndex = position1690, tokenIndex1690
					if !_rules[rulenn_contents]() {
						goto l1688
					}
				}
			l1690:
				{
					add(ruleAction17, position)
				}
				add(ruleSrc8, position1689)
			}
			return true
		l1688:
			position, tokenIndex = position1688, tokenIndex1688
			return false
		},
		/* 31 Loc8 <- <((Reg8 / Reg16Contents) Action18)> */
		func() bool {
			position1695, tokenIndex1695 := position, tokenIndex
			{
				position1696 := position
				{
					position1697, tokenIndex1697 := position, tokenIndex
					if !_rules[ruleReg8]() {
						goto l1698
					}
					goto l1697
				l1698:
					position, tokenIndex = position1697, tokenIndex1697
					if !_rules[ruleReg16Contents]() {
						goto l1695
					}
				}
			l1697:
				{
					add(ruleAction18, position)
				}
				add(ruleLoc8, position1696)
			}
			return true
		l1695:
			position, tokenIndex = position1695, tokenIndex1695
			return false
		},
		/* 32 Copy8 <- <(Reg8 Action19)> */
		func() bool {
			position1700, tokenIndex1700 := position, tokenIndex
			{
				position1701 := position
				if !_rules[ruleReg8]() {
					goto l1700
				}
				{
					add(ruleAction19, position)
				}
				add(ruleCopy8, position1701)
			}
			return true
		l1700:
			position, tokenIndex = position1700, tokenIndex1700
			return false
		},
		/* 33 ILoc8 <- <(IReg8 Action20)> */
		func() bool {
			position1703, tokenIndex1703 := position, tokenIndex
			{
				position1704 := position
				if !_rules[ruleIReg8]() {
					goto l1703
				}
				{
					add(ruleAction20, position)
				}
				add(ruleILoc8, position1704)
			}
			return true
		l1703:
			position, tokenIndex = position1703, tokenIndex1703
			return false
		},
		/* 34 Reg8 <- <(<(IReg8 / ((&('R' | 'r') R) | (&('I' | 'i') I) | (&('L' | 'l') L) | (&('H' | 'h') H) | (&('E' | 'e') E) | (&('D' | 'd') D) | (&('C' | 'c') C) | (&('B' | 'b') B) | (&('F' | 'f') F) | (&('A' | 'a') A)))> Action21)> */
		func() bool {
			position1706, tokenIndex1706 := position, tokenIndex
			{
				position1707 := position
				{
					position1708 := position
					{
						position1709, tokenIndex1709 := position, tokenIndex
						if !_rules[ruleIReg8]() {
							goto l1710
						}
						goto l1709
					l1710:
						position, tokenIndex = position1709, tokenIndex1709
						{
							switch buffer[position] {
							case 'R', 'r':
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
											goto l1706
										}
										position++
									}
								l1713:
									add(ruleR, position1712)
								}
								break
							case 'I', 'i':
								{
									position1715 := position
									{
										position1716, tokenIndex1716 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1717
										}
										position++
										goto l1716
									l1717:
										position, tokenIndex = position1716, tokenIndex1716
										if buffer[position] != rune('I') {
											goto l1706
										}
										position++
									}
								l1716:
									add(ruleI, position1715)
								}
								break
							case 'L', 'l':
								{
									position1718 := position
									{
										position1719, tokenIndex1719 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1720
										}
										position++
										goto l1719
									l1720:
										position, tokenIndex = position1719, tokenIndex1719
										if buffer[position] != rune('L') {
											goto l1706
										}
										position++
									}
								l1719:
									add(ruleL, position1718)
								}
								break
							case 'H', 'h':
								{
									position1721 := position
									{
										position1722, tokenIndex1722 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1723
										}
										position++
										goto l1722
									l1723:
										position, tokenIndex = position1722, tokenIndex1722
										if buffer[position] != rune('H') {
											goto l1706
										}
										position++
									}
								l1722:
									add(ruleH, position1721)
								}
								break
							case 'E', 'e':
								{
									position1724 := position
									{
										position1725, tokenIndex1725 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1726
										}
										position++
										goto l1725
									l1726:
										position, tokenIndex = position1725, tokenIndex1725
										if buffer[position] != rune('E') {
											goto l1706
										}
										position++
									}
								l1725:
									add(ruleE, position1724)
								}
								break
							case 'D', 'd':
								{
									position1727 := position
									{
										position1728, tokenIndex1728 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1729
										}
										position++
										goto l1728
									l1729:
										position, tokenIndex = position1728, tokenIndex1728
										if buffer[position] != rune('D') {
											goto l1706
										}
										position++
									}
								l1728:
									add(ruleD, position1727)
								}
								break
							case 'C', 'c':
								{
									position1730 := position
									{
										position1731, tokenIndex1731 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1732
										}
										position++
										goto l1731
									l1732:
										position, tokenIndex = position1731, tokenIndex1731
										if buffer[position] != rune('C') {
											goto l1706
										}
										position++
									}
								l1731:
									add(ruleC, position1730)
								}
								break
							case 'B', 'b':
								{
									position1733 := position
									{
										position1734, tokenIndex1734 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1735
										}
										position++
										goto l1734
									l1735:
										position, tokenIndex = position1734, tokenIndex1734
										if buffer[position] != rune('B') {
											goto l1706
										}
										position++
									}
								l1734:
									add(ruleB, position1733)
								}
								break
							case 'F', 'f':
								{
									position1736 := position
									{
										position1737, tokenIndex1737 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1738
										}
										position++
										goto l1737
									l1738:
										position, tokenIndex = position1737, tokenIndex1737
										if buffer[position] != rune('F') {
											goto l1706
										}
										position++
									}
								l1737:
									add(ruleF, position1736)
								}
								break
							default:
								{
									position1739 := position
									{
										position1740, tokenIndex1740 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1741
										}
										position++
										goto l1740
									l1741:
										position, tokenIndex = position1740, tokenIndex1740
										if buffer[position] != rune('A') {
											goto l1706
										}
										position++
									}
								l1740:
									add(ruleA, position1739)
								}
								break
							}
						}

					}
				l1709:
					add(rulePegText, position1708)
				}
				{
					add(ruleAction21, position)
				}
				add(ruleReg8, position1707)
			}
			return true
		l1706:
			position, tokenIndex = position1706, tokenIndex1706
			return false
		},
		/* 35 IReg8 <- <(<(IXH / IXL / IYH / IYL)> Action22)> */
		func() bool {
			position1743, tokenIndex1743 := position, tokenIndex
			{
				position1744 := position
				{
					position1745 := position
					{
						position1746, tokenIndex1746 := position, tokenIndex
						{
							position1748 := position
							{
								position1749, tokenIndex1749 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1750
								}
								position++
								goto l1749
							l1750:
								position, tokenIndex = position1749, tokenIndex1749
								if buffer[position] != rune('I') {
									goto l1747
								}
								position++
							}
						l1749:
							{
								position1751, tokenIndex1751 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1752
								}
								position++
								goto l1751
							l1752:
								position, tokenIndex = position1751, tokenIndex1751
								if buffer[position] != rune('X') {
									goto l1747
								}
								position++
							}
						l1751:
							{
								position1753, tokenIndex1753 := position, tokenIndex
								if buffer[position] != rune('h') {
									goto l1754
								}
								position++
								goto l1753
							l1754:
								position, tokenIndex = position1753, tokenIndex1753
								if buffer[position] != rune('H') {
									goto l1747
								}
								position++
							}
						l1753:
							add(ruleIXH, position1748)
						}
						goto l1746
					l1747:
						position, tokenIndex = position1746, tokenIndex1746
						{
							position1756 := position
							{
								position1757, tokenIndex1757 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1758
								}
								position++
								goto l1757
							l1758:
								position, tokenIndex = position1757, tokenIndex1757
								if buffer[position] != rune('I') {
									goto l1755
								}
								position++
							}
						l1757:
							{
								position1759, tokenIndex1759 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1760
								}
								position++
								goto l1759
							l1760:
								position, tokenIndex = position1759, tokenIndex1759
								if buffer[position] != rune('X') {
									goto l1755
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
									goto l1755
								}
								position++
							}
						l1761:
							add(ruleIXL, position1756)
						}
						goto l1746
					l1755:
						position, tokenIndex = position1746, tokenIndex1746
						{
							position1764 := position
							{
								position1765, tokenIndex1765 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1766
								}
								position++
								goto l1765
							l1766:
								position, tokenIndex = position1765, tokenIndex1765
								if buffer[position] != rune('I') {
									goto l1763
								}
								position++
							}
						l1765:
							{
								position1767, tokenIndex1767 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1768
								}
								position++
								goto l1767
							l1768:
								position, tokenIndex = position1767, tokenIndex1767
								if buffer[position] != rune('Y') {
									goto l1763
								}
								position++
							}
						l1767:
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
									goto l1763
								}
								position++
							}
						l1769:
							add(ruleIYH, position1764)
						}
						goto l1746
					l1763:
						position, tokenIndex = position1746, tokenIndex1746
						{
							position1771 := position
							{
								position1772, tokenIndex1772 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1773
								}
								position++
								goto l1772
							l1773:
								position, tokenIndex = position1772, tokenIndex1772
								if buffer[position] != rune('I') {
									goto l1743
								}
								position++
							}
						l1772:
							{
								position1774, tokenIndex1774 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1775
								}
								position++
								goto l1774
							l1775:
								position, tokenIndex = position1774, tokenIndex1774
								if buffer[position] != rune('Y') {
									goto l1743
								}
								position++
							}
						l1774:
							{
								position1776, tokenIndex1776 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l1777
								}
								position++
								goto l1776
							l1777:
								position, tokenIndex = position1776, tokenIndex1776
								if buffer[position] != rune('L') {
									goto l1743
								}
								position++
							}
						l1776:
							add(ruleIYL, position1771)
						}
					}
				l1746:
					add(rulePegText, position1745)
				}
				{
					add(ruleAction22, position)
				}
				add(ruleIReg8, position1744)
			}
			return true
		l1743:
			position, tokenIndex = position1743, tokenIndex1743
			return false
		},
		/* 36 Dst16 <- <((Reg16 / nn_contents / Reg16Contents) Action23)> */
		func() bool {
			position1779, tokenIndex1779 := position, tokenIndex
			{
				position1780 := position
				{
					position1781, tokenIndex1781 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1782
					}
					goto l1781
				l1782:
					position, tokenIndex = position1781, tokenIndex1781
					if !_rules[rulenn_contents]() {
						goto l1783
					}
					goto l1781
				l1783:
					position, tokenIndex = position1781, tokenIndex1781
					if !_rules[ruleReg16Contents]() {
						goto l1779
					}
				}
			l1781:
				{
					add(ruleAction23, position)
				}
				add(ruleDst16, position1780)
			}
			return true
		l1779:
			position, tokenIndex = position1779, tokenIndex1779
			return false
		},
		/* 37 Src16 <- <((Reg16 / nn / nn_contents) Action24)> */
		func() bool {
			position1785, tokenIndex1785 := position, tokenIndex
			{
				position1786 := position
				{
					position1787, tokenIndex1787 := position, tokenIndex
					if !_rules[ruleReg16]() {
						goto l1788
					}
					goto l1787
				l1788:
					position, tokenIndex = position1787, tokenIndex1787
					if !_rules[rulenn]() {
						goto l1789
					}
					goto l1787
				l1789:
					position, tokenIndex = position1787, tokenIndex1787
					if !_rules[rulenn_contents]() {
						goto l1785
					}
				}
			l1787:
				{
					add(ruleAction24, position)
				}
				add(ruleSrc16, position1786)
			}
			return true
		l1785:
			position, tokenIndex = position1785, tokenIndex1785
			return false
		},
		/* 38 Loc16 <- <(Reg16 Action25)> */
		func() bool {
			position1791, tokenIndex1791 := position, tokenIndex
			{
				position1792 := position
				if !_rules[ruleReg16]() {
					goto l1791
				}
				{
					add(ruleAction25, position)
				}
				add(ruleLoc16, position1792)
			}
			return true
		l1791:
			position, tokenIndex = position1791, tokenIndex1791
			return false
		},
		/* 39 Reg16 <- <(<(AF_PRIME / ((&('I' | 'i') IReg16) | (&('S' | 's') SP) | (&('H' | 'h') HL) | (&('D' | 'd') DE) | (&('B' | 'b') BC) | (&('A' | 'a') AF)))> Action26)> */
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
								if buffer[position] != rune('a') {
									goto l1801
								}
								position++
								goto l1800
							l1801:
								position, tokenIndex = position1800, tokenIndex1800
								if buffer[position] != rune('A') {
									goto l1798
								}
								position++
							}
						l1800:
							{
								position1802, tokenIndex1802 := position, tokenIndex
								if buffer[position] != rune('f') {
									goto l1803
								}
								position++
								goto l1802
							l1803:
								position, tokenIndex = position1802, tokenIndex1802
								if buffer[position] != rune('F') {
									goto l1798
								}
								position++
							}
						l1802:
							if buffer[position] != rune('\'') {
								goto l1798
							}
							position++
							add(ruleAF_PRIME, position1799)
						}
						goto l1797
					l1798:
						position, tokenIndex = position1797, tokenIndex1797
						{
							switch buffer[position] {
							case 'I', 'i':
								if !_rules[ruleIReg16]() {
									goto l1794
								}
								break
							case 'S', 's':
								{
									position1805 := position
									{
										position1806, tokenIndex1806 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l1807
										}
										position++
										goto l1806
									l1807:
										position, tokenIndex = position1806, tokenIndex1806
										if buffer[position] != rune('S') {
											goto l1794
										}
										position++
									}
								l1806:
									{
										position1808, tokenIndex1808 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l1809
										}
										position++
										goto l1808
									l1809:
										position, tokenIndex = position1808, tokenIndex1808
										if buffer[position] != rune('P') {
											goto l1794
										}
										position++
									}
								l1808:
									add(ruleSP, position1805)
								}
								break
							case 'H', 'h':
								{
									position1810 := position
									{
										position1811, tokenIndex1811 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1812
										}
										position++
										goto l1811
									l1812:
										position, tokenIndex = position1811, tokenIndex1811
										if buffer[position] != rune('H') {
											goto l1794
										}
										position++
									}
								l1811:
									{
										position1813, tokenIndex1813 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1814
										}
										position++
										goto l1813
									l1814:
										position, tokenIndex = position1813, tokenIndex1813
										if buffer[position] != rune('L') {
											goto l1794
										}
										position++
									}
								l1813:
									add(ruleHL, position1810)
								}
								break
							case 'D', 'd':
								{
									position1815 := position
									{
										position1816, tokenIndex1816 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1817
										}
										position++
										goto l1816
									l1817:
										position, tokenIndex = position1816, tokenIndex1816
										if buffer[position] != rune('D') {
											goto l1794
										}
										position++
									}
								l1816:
									{
										position1818, tokenIndex1818 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1819
										}
										position++
										goto l1818
									l1819:
										position, tokenIndex = position1818, tokenIndex1818
										if buffer[position] != rune('E') {
											goto l1794
										}
										position++
									}
								l1818:
									add(ruleDE, position1815)
								}
								break
							case 'B', 'b':
								{
									position1820 := position
									{
										position1821, tokenIndex1821 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1822
										}
										position++
										goto l1821
									l1822:
										position, tokenIndex = position1821, tokenIndex1821
										if buffer[position] != rune('B') {
											goto l1794
										}
										position++
									}
								l1821:
									{
										position1823, tokenIndex1823 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l1824
										}
										position++
										goto l1823
									l1824:
										position, tokenIndex = position1823, tokenIndex1823
										if buffer[position] != rune('C') {
											goto l1794
										}
										position++
									}
								l1823:
									add(ruleBC, position1820)
								}
								break
							default:
								{
									position1825 := position
									{
										position1826, tokenIndex1826 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1827
										}
										position++
										goto l1826
									l1827:
										position, tokenIndex = position1826, tokenIndex1826
										if buffer[position] != rune('A') {
											goto l1794
										}
										position++
									}
								l1826:
									{
										position1828, tokenIndex1828 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l1829
										}
										position++
										goto l1828
									l1829:
										position, tokenIndex = position1828, tokenIndex1828
										if buffer[position] != rune('F') {
											goto l1794
										}
										position++
									}
								l1828:
									add(ruleAF, position1825)
								}
								break
							}
						}

					}
				l1797:
					add(rulePegText, position1796)
				}
				{
					add(ruleAction26, position)
				}
				add(ruleReg16, position1795)
			}
			return true
		l1794:
			position, tokenIndex = position1794, tokenIndex1794
			return false
		},
		/* 40 IReg16 <- <(<(IX / IY)> Action27)> */
		func() bool {
			position1831, tokenIndex1831 := position, tokenIndex
			{
				position1832 := position
				{
					position1833 := position
					{
						position1834, tokenIndex1834 := position, tokenIndex
						{
							position1836 := position
							{
								position1837, tokenIndex1837 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1838
								}
								position++
								goto l1837
							l1838:
								position, tokenIndex = position1837, tokenIndex1837
								if buffer[position] != rune('I') {
									goto l1835
								}
								position++
							}
						l1837:
							{
								position1839, tokenIndex1839 := position, tokenIndex
								if buffer[position] != rune('x') {
									goto l1840
								}
								position++
								goto l1839
							l1840:
								position, tokenIndex = position1839, tokenIndex1839
								if buffer[position] != rune('X') {
									goto l1835
								}
								position++
							}
						l1839:
							add(ruleIX, position1836)
						}
						goto l1834
					l1835:
						position, tokenIndex = position1834, tokenIndex1834
						{
							position1841 := position
							{
								position1842, tokenIndex1842 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l1843
								}
								position++
								goto l1842
							l1843:
								position, tokenIndex = position1842, tokenIndex1842
								if buffer[position] != rune('I') {
									goto l1831
								}
								position++
							}
						l1842:
							{
								position1844, tokenIndex1844 := position, tokenIndex
								if buffer[position] != rune('y') {
									goto l1845
								}
								position++
								goto l1844
							l1845:
								position, tokenIndex = position1844, tokenIndex1844
								if buffer[position] != rune('Y') {
									goto l1831
								}
								position++
							}
						l1844:
							add(ruleIY, position1841)
						}
					}
				l1834:
					add(rulePegText, position1833)
				}
				{
					add(ruleAction27, position)
				}
				add(ruleIReg16, position1832)
			}
			return true
		l1831:
			position, tokenIndex = position1831, tokenIndex1831
			return false
		},
		/* 41 Reg16Contents <- <(IndexedR16C / PlainR16C)> */
		func() bool {
			position1847, tokenIndex1847 := position, tokenIndex
			{
				position1848 := position
				{
					position1849, tokenIndex1849 := position, tokenIndex
					{
						position1851 := position
						if buffer[position] != rune('(') {
							goto l1850
						}
						position++
						if !_rules[ruleIReg16]() {
							goto l1850
						}
						{
							position1852, tokenIndex1852 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1852
							}
							goto l1853
						l1852:
							position, tokenIndex = position1852, tokenIndex1852
						}
					l1853:
						if !_rules[ruledisp]() {
							goto l1850
						}
						{
							position1854, tokenIndex1854 := position, tokenIndex
							if !_rules[rulews]() {
								goto l1854
							}
							goto l1855
						l1854:
							position, tokenIndex = position1854, tokenIndex1854
						}
					l1855:
						if buffer[position] != rune(')') {
							goto l1850
						}
						position++
						{
							add(ruleAction29, position)
						}
						add(ruleIndexedR16C, position1851)
					}
					goto l1849
				l1850:
					position, tokenIndex = position1849, tokenIndex1849
					{
						position1857 := position
						if buffer[position] != rune('(') {
							goto l1847
						}
						position++
						if !_rules[ruleReg16]() {
							goto l1847
						}
						if buffer[position] != rune(')') {
							goto l1847
						}
						position++
						{
							add(ruleAction28, position)
						}
						add(rulePlainR16C, position1857)
					}
				}
			l1849:
				add(ruleReg16Contents, position1848)
			}
			return true
		l1847:
			position, tokenIndex = position1847, tokenIndex1847
			return false
		},
		/* 42 PlainR16C <- <('(' Reg16 ')' Action28)> */
		nil,
		/* 43 IndexedR16C <- <('(' IReg16 ws? disp ws? ')' Action29)> */
		nil,
		/* 44 n <- <(hexByteH / hexByte0x / decimalByte)> */
		func() bool {
			position1861, tokenIndex1861 := position, tokenIndex
			{
				position1862 := position
				{
					position1863, tokenIndex1863 := position, tokenIndex
					{
						position1865 := position
						{
							position1866 := position
							if !_rules[rulehexdigit]() {
								goto l1864
							}
							if !_rules[rulehexdigit]() {
								goto l1864
							}
							add(rulePegText, position1866)
						}
						{
							position1867, tokenIndex1867 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1868
							}
							position++
							goto l1867
						l1868:
							position, tokenIndex = position1867, tokenIndex1867
							if buffer[position] != rune('H') {
								goto l1864
							}
							position++
						}
					l1867:
						{
							add(ruleAction33, position)
						}
						add(rulehexByteH, position1865)
					}
					goto l1863
				l1864:
					position, tokenIndex = position1863, tokenIndex1863
					{
						position1871 := position
						if buffer[position] != rune('0') {
							goto l1870
						}
						position++
						{
							position1872, tokenIndex1872 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1873
							}
							position++
							goto l1872
						l1873:
							position, tokenIndex = position1872, tokenIndex1872
							if buffer[position] != rune('X') {
								goto l1870
							}
							position++
						}
					l1872:
						{
							position1874 := position
							if !_rules[rulehexdigit]() {
								goto l1870
							}
							if !_rules[rulehexdigit]() {
								goto l1870
							}
							add(rulePegText, position1874)
						}
						{
							add(ruleAction34, position)
						}
						add(rulehexByte0x, position1871)
					}
					goto l1863
				l1870:
					position, tokenIndex = position1863, tokenIndex1863
					{
						position1876 := position
						{
							position1877 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1861
							}
							position++
						l1878:
							{
								position1879, tokenIndex1879 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1879
								}
								position++
								goto l1878
							l1879:
								position, tokenIndex = position1879, tokenIndex1879
							}
							add(rulePegText, position1877)
						}
						{
							add(ruleAction35, position)
						}
						add(ruledecimalByte, position1876)
					}
				}
			l1863:
				add(rulen, position1862)
			}
			return true
		l1861:
			position, tokenIndex = position1861, tokenIndex1861
			return false
		},
		/* 45 nn <- <(hexWordH / hexWord0x)> */
		func() bool {
			position1881, tokenIndex1881 := position, tokenIndex
			{
				position1882 := position
				{
					position1883, tokenIndex1883 := position, tokenIndex
					{
						position1885 := position
						{
							position1886 := position
							if !_rules[rulehexdigit]() {
								goto l1884
							}
							if !_rules[rulehexdigit]() {
								goto l1884
							}
							if !_rules[rulehexdigit]() {
								goto l1884
							}
							if !_rules[rulehexdigit]() {
								goto l1884
							}
							add(rulePegText, position1886)
						}
						{
							position1887, tokenIndex1887 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1888
							}
							position++
							goto l1887
						l1888:
							position, tokenIndex = position1887, tokenIndex1887
							if buffer[position] != rune('H') {
								goto l1884
							}
							position++
						}
					l1887:
						{
							add(ruleAction36, position)
						}
						add(rulehexWordH, position1885)
					}
					goto l1883
				l1884:
					position, tokenIndex = position1883, tokenIndex1883
					{
						position1890 := position
						if buffer[position] != rune('0') {
							goto l1881
						}
						position++
						{
							position1891, tokenIndex1891 := position, tokenIndex
							if buffer[position] != rune('x') {
								goto l1892
							}
							position++
							goto l1891
						l1892:
							position, tokenIndex = position1891, tokenIndex1891
							if buffer[position] != rune('X') {
								goto l1881
							}
							position++
						}
					l1891:
						{
							position1893 := position
							if !_rules[rulehexdigit]() {
								goto l1881
							}
							if !_rules[rulehexdigit]() {
								goto l1881
							}
							if !_rules[rulehexdigit]() {
								goto l1881
							}
							if !_rules[rulehexdigit]() {
								goto l1881
							}
							add(rulePegText, position1893)
						}
						{
							add(ruleAction37, position)
						}
						add(rulehexWord0x, position1890)
					}
				}
			l1883:
				add(rulenn, position1882)
			}
			return true
		l1881:
			position, tokenIndex = position1881, tokenIndex1881
			return false
		},
		/* 46 disp <- <(signedHexByteH / signedHexByte0x / signedDecimalByte)> */
		func() bool {
			position1895, tokenIndex1895 := position, tokenIndex
			{
				position1896 := position
				{
					position1897, tokenIndex1897 := position, tokenIndex
					{
						position1899 := position
						{
							position1900 := position
							{
								position1901, tokenIndex1901 := position, tokenIndex
								{
									position1903, tokenIndex1903 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1904
									}
									position++
									goto l1903
								l1904:
									position, tokenIndex = position1903, tokenIndex1903
									if buffer[position] != rune('+') {
										goto l1901
									}
									position++
								}
							l1903:
								goto l1902
							l1901:
								position, tokenIndex = position1901, tokenIndex1901
							}
						l1902:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1898
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1898
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1898
									}
									position++
									break
								}
							}

						l1905:
							{
								position1906, tokenIndex1906 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1906
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1906
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1906
										}
										position++
										break
									}
								}

								goto l1905
							l1906:
								position, tokenIndex = position1906, tokenIndex1906
							}
							add(rulePegText, position1900)
						}
						{
							position1909, tokenIndex1909 := position, tokenIndex
							if buffer[position] != rune('h') {
								goto l1910
							}
							position++
							goto l1909
						l1910:
							position, tokenIndex = position1909, tokenIndex1909
							if buffer[position] != rune('H') {
								goto l1898
							}
							position++
						}
					l1909:
						{
							add(ruleAction31, position)
						}
						add(rulesignedHexByteH, position1899)
					}
					goto l1897
				l1898:
					position, tokenIndex = position1897, tokenIndex1897
					{
						position1913 := position
						{
							position1914 := position
							{
								position1915, tokenIndex1915 := position, tokenIndex
								{
									position1917, tokenIndex1917 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1918
									}
									position++
									goto l1917
								l1918:
									position, tokenIndex = position1917, tokenIndex1917
									if buffer[position] != rune('+') {
										goto l1915
									}
									position++
								}
							l1917:
								goto l1916
							l1915:
								position, tokenIndex = position1915, tokenIndex1915
							}
						l1916:
							if buffer[position] != rune('0') {
								goto l1912
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
									goto l1912
								}
								position++
							}
						l1919:
							{
								switch buffer[position] {
								case 'A', 'B', 'C', 'D', 'E', 'F':
									if c := buffer[position]; c < rune('A') || c > rune('F') {
										goto l1912
									}
									position++
									break
								case 'a', 'b', 'c', 'd', 'e', 'f':
									if c := buffer[position]; c < rune('a') || c > rune('f') {
										goto l1912
									}
									position++
									break
								default:
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l1912
									}
									position++
									break
								}
							}

						l1921:
							{
								position1922, tokenIndex1922 := position, tokenIndex
								{
									switch buffer[position] {
									case 'A', 'B', 'C', 'D', 'E', 'F':
										if c := buffer[position]; c < rune('A') || c > rune('F') {
											goto l1922
										}
										position++
										break
									case 'a', 'b', 'c', 'd', 'e', 'f':
										if c := buffer[position]; c < rune('a') || c > rune('f') {
											goto l1922
										}
										position++
										break
									default:
										if c := buffer[position]; c < rune('0') || c > rune('9') {
											goto l1922
										}
										position++
										break
									}
								}

								goto l1921
							l1922:
								position, tokenIndex = position1922, tokenIndex1922
							}
							add(rulePegText, position1914)
						}
						{
							add(ruleAction32, position)
						}
						add(rulesignedHexByte0x, position1913)
					}
					goto l1897
				l1912:
					position, tokenIndex = position1897, tokenIndex1897
					{
						position1926 := position
						{
							position1927 := position
							{
								position1928, tokenIndex1928 := position, tokenIndex
								{
									position1930, tokenIndex1930 := position, tokenIndex
									if buffer[position] != rune('-') {
										goto l1931
									}
									position++
									goto l1930
								l1931:
									position, tokenIndex = position1930, tokenIndex1930
									if buffer[position] != rune('+') {
										goto l1928
									}
									position++
								}
							l1930:
								goto l1929
							l1928:
								position, tokenIndex = position1928, tokenIndex1928
							}
						l1929:
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1895
							}
							position++
						l1932:
							{
								position1933, tokenIndex1933 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l1933
								}
								position++
								goto l1932
							l1933:
								position, tokenIndex = position1933, tokenIndex1933
							}
							add(rulePegText, position1927)
						}
						{
							add(ruleAction30, position)
						}
						add(rulesignedDecimalByte, position1926)
					}
				}
			l1897:
				add(ruledisp, position1896)
			}
			return true
		l1895:
			position, tokenIndex = position1895, tokenIndex1895
			return false
		},
		/* 47 signedDecimalByte <- <(<(('-' / '+')? [0-9]+)> Action30)> */
		nil,
		/* 48 signedHexByteH <- <(<(('-' / '+')? ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> ('h' / 'H') Action31)> */
		nil,
		/* 49 signedHexByte0x <- <(<(('-' / '+')? ('0' ('x' / 'X')) ((&('A' | 'B' | 'C' | 'D' | 'E' | 'F') [A-F]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f') [a-f]) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]))+)> Action32)> */
		nil,
		/* 50 hexByteH <- <(<(hexdigit hexdigit)> ('h' / 'H') Action33)> */
		nil,
		/* 51 hexByte0x <- <('0' ('x' / 'X') <(hexdigit hexdigit)> Action34)> */
		nil,
		/* 52 decimalByte <- <(<[0-9]+> Action35)> */
		nil,
		/* 53 hexWordH <- <(<(hexdigit hexdigit hexdigit hexdigit)> ('h' / 'H') Action36)> */
		nil,
		/* 54 hexWord0x <- <('0' ('x' / 'X') <(hexdigit hexdigit hexdigit hexdigit)> Action37)> */
		nil,
		/* 55 nn_contents <- <('(' nn ')' Action38)> */
		func() bool {
			position1943, tokenIndex1943 := position, tokenIndex
			{
				position1944 := position
				if buffer[position] != rune('(') {
					goto l1943
				}
				position++
				if !_rules[rulenn]() {
					goto l1943
				}
				if buffer[position] != rune(')') {
					goto l1943
				}
				position++
				{
					add(ruleAction38, position)
				}
				add(rulenn_contents, position1944)
			}
			return true
		l1943:
			position, tokenIndex = position1943, tokenIndex1943
			return false
		},
		/* 56 Alu <- <(Add / Adc / Sub / ((&('C' | 'c') Cp) | (&('O' | 'o') Or) | (&('X' | 'x') Xor) | (&('A' | 'a') And) | (&('S' | 's') Sbc)))> */
		nil,
		/* 57 Add <- <(('a' / 'A') ('d' / 'D') ('d' / 'D') ws ('a' / 'A') sep Src8 Action39)> */
		nil,
		/* 58 Adc <- <(('a' / 'A') ('d' / 'D') ('c' / 'C') ws ('a' / 'A') sep Src8 Action40)> */
		nil,
		/* 59 Sub <- <(('s' / 'S') ('u' / 'U') ('b' / 'B') ws Src8 Action41)> */
		nil,
		/* 60 Sbc <- <(('s' / 'S') ('b' / 'B') ('c' / 'C') ws ('a' / 'A') sep Src8 Action42)> */
		nil,
		/* 61 And <- <(('a' / 'A') ('n' / 'N') ('d' / 'D') ws Src8 Action43)> */
		nil,
		/* 62 Xor <- <(('x' / 'X') ('o' / 'O') ('r' / 'R') ws Src8 Action44)> */
		nil,
		/* 63 Or <- <(('o' / 'O') ('r' / 'R') ws Src8 Action45)> */
		nil,
		/* 64 Cp <- <(('c' / 'C') ('p' / 'P') ws Src8 Action46)> */
		nil,
		/* 65 BitOp <- <(Rot / ((&('S' | 's') Set) | (&('R' | 'r') Res) | (&('B' | 'b') Bit)))> */
		nil,
		/* 66 Rot <- <(Rlc / Rrc / Rl / Rr / Sla / Sra / Sll / Srl)> */
		nil,
		/* 67 Rlc <- <(('r' / 'R') ('l' / 'L') ('c' / 'C') ws Loc8 (sep Copy8)? Action47)> */
		nil,
		/* 68 Rrc <- <(('r' / 'R') ('r' / 'R') ('c' / 'C') ws Loc8 (sep Copy8)? Action48)> */
		nil,
		/* 69 Rl <- <(('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action49)> */
		nil,
		/* 70 Rr <- <(('r' / 'R') ('r' / 'R') ws Loc8 (sep Copy8)? Action50)> */
		nil,
		/* 71 Sla <- <(('s' / 'S') ('l' / 'L') ('a' / 'A') ws Loc8 (sep Copy8)? Action51)> */
		nil,
		/* 72 Sra <- <(('s' / 'S') ('r' / 'R') ('a' / 'A') ws Loc8 (sep Copy8)? Action52)> */
		nil,
		/* 73 Sll <- <(('s' / 'S') ('l' / 'L') ('l' / 'L') ws Loc8 (sep Copy8)? Action53)> */
		nil,
		/* 74 Srl <- <(('s' / 'S') ('r' / 'R') ('l' / 'L') ws Loc8 (sep Copy8)? Action54)> */
		nil,
		/* 75 Bit <- <(('b' / 'B') ('i' / 'I') ('t' / 'T') ws octaldigit sep Loc8 Action55)> */
		nil,
		/* 76 Res <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ws octaldigit sep Loc8 (sep Copy8)? Action56)> */
		nil,
		/* 77 Set <- <(('s' / 'S') ('e' / 'E') ('t' / 'T') ws octaldigit sep Loc8 (sep Copy8)? Action57)> */
		nil,
		/* 78 Simple <- <(Rlca / Rrca / Rla / Daa / Cpl / Exx / ((&('E' | 'e') Ei) | (&('D' | 'd') Di) | (&('C' | 'c') Ccf) | (&('S' | 's') Scf) | (&('R' | 'r') Rra) | (&('H' | 'h') Halt) | (&('N' | 'n') Nop)))> */
		nil,
		/* 79 Nop <- <(<(('n' / 'N') ('o' / 'O') ('p' / 'P'))> Action58)> */
		nil,
		/* 80 Halt <- <(<(('h' / 'H') ('a' / 'A') ('l' / 'L') ('t' / 'T'))> Action59)> */
		nil,
		/* 81 Rlca <- <(<(('r' / 'R') ('l' / 'L') ('c' / 'C') ('a' / 'A'))> Action60)> */
		nil,
		/* 82 Rrca <- <(<(('r' / 'R') ('r' / 'R') ('c' / 'C') ('a' / 'A'))> Action61)> */
		nil,
		/* 83 Rla <- <(<(('r' / 'R') ('l' / 'L') ('a' / 'A'))> Action62)> */
		nil,
		/* 84 Rra <- <(<(('r' / 'R') ('r' / 'R') ('a' / 'A'))> Action63)> */
		nil,
		/* 85 Daa <- <(<(('d' / 'D') ('a' / 'A') ('a' / 'A'))> Action64)> */
		nil,
		/* 86 Cpl <- <(<(('c' / 'C') ('p' / 'P') ('l' / 'L'))> Action65)> */
		nil,
		/* 87 Scf <- <(<(('s' / 'S') ('c' / 'C') ('f' / 'F'))> Action66)> */
		nil,
		/* 88 Ccf <- <(<(('c' / 'C') ('c' / 'C') ('f' / 'F'))> Action67)> */
		nil,
		/* 89 Exx <- <(<(('e' / 'E') ('x' / 'X') ('x' / 'X'))> Action68)> */
		nil,
		/* 90 Di <- <(<(('d' / 'D') ('i' / 'I'))> Action69)> */
		nil,
		/* 91 Ei <- <(<(('e' / 'E') ('i' / 'I'))> Action70)> */
		nil,
		/* 92 EDSimple <- <(Retn / Reti / Rrd / Im0 / Im1 / Im2 / ((&('I' | 'O' | 'i' | 'o') BlitIO) | (&('R' | 'r') Rld) | (&('N' | 'n') Neg) | (&('C' | 'L' | 'c' | 'l') Blit)))> */
		nil,
		/* 93 Neg <- <(<(('n' / 'N') ('e' / 'E') ('g' / 'G'))> Action71)> */
		nil,
		/* 94 Retn <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N'))> Action72)> */
		nil,
		/* 95 Reti <- <(<(('r' / 'R') ('e' / 'E') ('t' / 'T') ('i' / 'I'))> Action73)> */
		nil,
		/* 96 Rrd <- <(<(('r' / 'R') ('r' / 'R') ('d' / 'D'))> Action74)> */
		nil,
		/* 97 Rld <- <(<(('r' / 'R') ('l' / 'L') ('d' / 'D'))> Action75)> */
		nil,
		/* 98 Im0 <- <(<(('i' / 'I') ('m' / 'M') ' ' '0')> Action76)> */
		nil,
		/* 99 Im1 <- <(<(('i' / 'I') ('m' / 'M') ' ' '1')> Action77)> */
		nil,
		/* 100 Im2 <- <(<(('i' / 'I') ('m' / 'M') ' ' '2')> Action78)> */
		nil,
		/* 101 Blit <- <(Ldir / Ldi / Cpir / Cpi / Lddr / Ldd / Cpdr / Cpd)> */
		nil,
		/* 102 BlitIO <- <(Inir / Ini / Otir / Outi / Indr / Ind / Otdr / Outd)> */
		nil,
		/* 103 Ldi <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I'))> Action79)> */
		nil,
		/* 104 Cpi <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I'))> Action80)> */
		nil,
		/* 105 Ini <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I'))> Action81)> */
		nil,
		/* 106 Outi <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('i' / 'I'))> Action82)> */
		nil,
		/* 107 Ldd <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D'))> Action83)> */
		nil,
		/* 108 Cpd <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D'))> Action84)> */
		nil,
		/* 109 Ind <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D'))> Action85)> */
		nil,
		/* 110 Outd <- <(<(('o' / 'O') ('u' / 'U') ('t' / 'T') ('d' / 'D'))> Action86)> */
		nil,
		/* 111 Ldir <- <(<(('l' / 'L') ('d' / 'D') ('i' / 'I') ('r' / 'R'))> Action87)> */
		nil,
		/* 112 Cpir <- <(<(('c' / 'C') ('p' / 'P') ('i' / 'I') ('r' / 'R'))> Action88)> */
		nil,
		/* 113 Inir <- <(<(('i' / 'I') ('n' / 'N') ('i' / 'I') ('r' / 'R'))> Action89)> */
		nil,
		/* 114 Otir <- <(<(('o' / 'O') ('t' / 'T') ('i' / 'I') ('r' / 'R'))> Action90)> */
		nil,
		/* 115 Lddr <- <(<(('l' / 'L') ('d' / 'D') ('d' / 'D') ('r' / 'R'))> Action91)> */
		nil,
		/* 116 Cpdr <- <(<(('c' / 'C') ('p' / 'P') ('d' / 'D') ('r' / 'R'))> Action92)> */
		nil,
		/* 117 Indr <- <(<(('i' / 'I') ('n' / 'N') ('d' / 'D') ('r' / 'R'))> Action93)> */
		nil,
		/* 118 Otdr <- <(<(('o' / 'O') ('t' / 'T') ('d' / 'D') ('r' / 'R'))> Action94)> */
		nil,
		/* 119 Jump <- <(Rst / Jp / ((&('D' | 'd') Djnz) | (&('J' | 'j') Jr) | (&('R' | 'r') Ret) | (&('C' | 'c') Call)))> */
		nil,
		/* 120 Rst <- <(('r' / 'R') ('s' / 'S') ('t' / 'T') ws n Action95)> */
		nil,
		/* 121 Call <- <(('c' / 'C') ('a' / 'A') ('l' / 'L') ('l' / 'L') ws (cc sep)? Src16 Action96)> */
		nil,
		/* 122 Ret <- <(('r' / 'R') ('e' / 'E') ('t' / 'T') (ws cc)? Action97)> */
		nil,
		/* 123 Jp <- <(('j' / 'J') ('p' / 'P') ws (cc sep)? Src16 Action98)> */
		nil,
		/* 124 Jr <- <(('j' / 'J') ('r' / 'R') ws (cc sep)? disp Action99)> */
		nil,
		/* 125 Djnz <- <(('d' / 'D') ('j' / 'J') ('n' / 'N') ('z' / 'Z') ws disp Action100)> */
		nil,
		/* 126 IO <- <(IN / OUT)> */
		nil,
		/* 127 IN <- <(('i' / 'I') ('n' / 'N') ws Reg8 sep Port Action101)> */
		nil,
		/* 128 OUT <- <(('o' / 'O') ('u' / 'U') ('t' / 'T') ws Port sep Reg8 Action102)> */
		nil,
		/* 129 Port <- <(('(' ('c' / 'C') ')') / ('(' n ')'))> */
		func() bool {
			position2019, tokenIndex2019 := position, tokenIndex
			{
				position2020 := position
				{
					position2021, tokenIndex2021 := position, tokenIndex
					if buffer[position] != rune('(') {
						goto l2022
					}
					position++
					{
						position2023, tokenIndex2023 := position, tokenIndex
						if buffer[position] != rune('c') {
							goto l2024
						}
						position++
						goto l2023
					l2024:
						position, tokenIndex = position2023, tokenIndex2023
						if buffer[position] != rune('C') {
							goto l2022
						}
						position++
					}
				l2023:
					if buffer[position] != rune(')') {
						goto l2022
					}
					position++
					goto l2021
				l2022:
					position, tokenIndex = position2021, tokenIndex2021
					if buffer[position] != rune('(') {
						goto l2019
					}
					position++
					if !_rules[rulen]() {
						goto l2019
					}
					if buffer[position] != rune(')') {
						goto l2019
					}
					position++
				}
			l2021:
				add(rulePort, position2020)
			}
			return true
		l2019:
			position, tokenIndex = position2019, tokenIndex2019
			return false
		},
		/* 130 sep <- <(ws? ',' ws?)> */
		func() bool {
			position2025, tokenIndex2025 := position, tokenIndex
			{
				position2026 := position
				{
					position2027, tokenIndex2027 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2027
					}
					goto l2028
				l2027:
					position, tokenIndex = position2027, tokenIndex2027
				}
			l2028:
				if buffer[position] != rune(',') {
					goto l2025
				}
				position++
				{
					position2029, tokenIndex2029 := position, tokenIndex
					if !_rules[rulews]() {
						goto l2029
					}
					goto l2030
				l2029:
					position, tokenIndex = position2029, tokenIndex2029
				}
			l2030:
				add(rulesep, position2026)
			}
			return true
		l2025:
			position, tokenIndex = position2025, tokenIndex2025
			return false
		},
		/* 131 ws <- <' '+> */
		func() bool {
			position2031, tokenIndex2031 := position, tokenIndex
			{
				position2032 := position
				if buffer[position] != rune(' ') {
					goto l2031
				}
				position++
			l2033:
				{
					position2034, tokenIndex2034 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l2034
					}
					position++
					goto l2033
				l2034:
					position, tokenIndex = position2034, tokenIndex2034
				}
				add(rulews, position2032)
			}
			return true
		l2031:
			position, tokenIndex = position2031, tokenIndex2031
			return false
		},
		/* 132 A <- <('a' / 'A')> */
		nil,
		/* 133 F <- <('f' / 'F')> */
		nil,
		/* 134 B <- <('b' / 'B')> */
		nil,
		/* 135 C <- <('c' / 'C')> */
		nil,
		/* 136 D <- <('d' / 'D')> */
		nil,
		/* 137 E <- <('e' / 'E')> */
		nil,
		/* 138 H <- <('h' / 'H')> */
		nil,
		/* 139 L <- <('l' / 'L')> */
		nil,
		/* 140 IXH <- <(('i' / 'I') ('x' / 'X') ('h' / 'H'))> */
		nil,
		/* 141 IXL <- <(('i' / 'I') ('x' / 'X') ('l' / 'L'))> */
		nil,
		/* 142 IYH <- <(('i' / 'I') ('y' / 'Y') ('h' / 'H'))> */
		nil,
		/* 143 IYL <- <(('i' / 'I') ('y' / 'Y') ('l' / 'L'))> */
		nil,
		/* 144 I <- <('i' / 'I')> */
		nil,
		/* 145 R <- <('r' / 'R')> */
		nil,
		/* 146 AF <- <(('a' / 'A') ('f' / 'F'))> */
		nil,
		/* 147 AF_PRIME <- <(('a' / 'A') ('f' / 'F') '\'')> */
		nil,
		/* 148 BC <- <(('b' / 'B') ('c' / 'C'))> */
		nil,
		/* 149 DE <- <(('d' / 'D') ('e' / 'E'))> */
		nil,
		/* 150 HL <- <(('h' / 'H') ('l' / 'L'))> */
		nil,
		/* 151 IX <- <(('i' / 'I') ('x' / 'X'))> */
		nil,
		/* 152 IY <- <(('i' / 'I') ('y' / 'Y'))> */
		nil,
		/* 153 SP <- <(('s' / 'S') ('p' / 'P'))> */
		nil,
		/* 154 hexdigit <- <([0-9] / ([a-f] / [A-F]))> */
		func() bool {
			position2057, tokenIndex2057 := position, tokenIndex
			{
				position2058 := position
				{
					position2059, tokenIndex2059 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l2060
					}
					position++
					goto l2059
				l2060:
					position, tokenIndex = position2059, tokenIndex2059
					{
						position2061, tokenIndex2061 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('f') {
							goto l2062
						}
						position++
						goto l2061
					l2062:
						position, tokenIndex = position2061, tokenIndex2061
						if c := buffer[position]; c < rune('A') || c > rune('F') {
							goto l2057
						}
						position++
					}
				l2061:
				}
			l2059:
				add(rulehexdigit, position2058)
			}
			return true
		l2057:
			position, tokenIndex = position2057, tokenIndex2057
			return false
		},
		/* 155 octaldigit <- <(<[0-7]> Action103)> */
		func() bool {
			position2063, tokenIndex2063 := position, tokenIndex
			{
				position2064 := position
				{
					position2065 := position
					if c := buffer[position]; c < rune('0') || c > rune('7') {
						goto l2063
					}
					position++
					add(rulePegText, position2065)
				}
				{
					add(ruleAction103, position)
				}
				add(ruleoctaldigit, position2064)
			}
			return true
		l2063:
			position, tokenIndex = position2063, tokenIndex2063
			return false
		},
		/* 156 cc <- <(FT_NZ / FT_PO / FT_PE / ((&('M' | 'm') FT_M) | (&('P' | 'p') FT_P) | (&('C' | 'c') FT_C) | (&('N' | 'n') FT_NC) | (&('Z' | 'z') FT_Z)))> */
		func() bool {
			position2067, tokenIndex2067 := position, tokenIndex
			{
				position2068 := position
				{
					position2069, tokenIndex2069 := position, tokenIndex
					{
						position2071 := position
						{
							position2072, tokenIndex2072 := position, tokenIndex
							if buffer[position] != rune('n') {
								goto l2073
							}
							position++
							goto l2072
						l2073:
							position, tokenIndex = position2072, tokenIndex2072
							if buffer[position] != rune('N') {
								goto l2070
							}
							position++
						}
					l2072:
						{
							position2074, tokenIndex2074 := position, tokenIndex
							if buffer[position] != rune('z') {
								goto l2075
							}
							position++
							goto l2074
						l2075:
							position, tokenIndex = position2074, tokenIndex2074
							if buffer[position] != rune('Z') {
								goto l2070
							}
							position++
						}
					l2074:
						{
							add(ruleAction104, position)
						}
						add(ruleFT_NZ, position2071)
					}
					goto l2069
				l2070:
					position, tokenIndex = position2069, tokenIndex2069
					{
						position2078 := position
						{
							position2079, tokenIndex2079 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2080
							}
							position++
							goto l2079
						l2080:
							position, tokenIndex = position2079, tokenIndex2079
							if buffer[position] != rune('P') {
								goto l2077
							}
							position++
						}
					l2079:
						{
							position2081, tokenIndex2081 := position, tokenIndex
							if buffer[position] != rune('o') {
								goto l2082
							}
							position++
							goto l2081
						l2082:
							position, tokenIndex = position2081, tokenIndex2081
							if buffer[position] != rune('O') {
								goto l2077
							}
							position++
						}
					l2081:
						{
							add(ruleAction108, position)
						}
						add(ruleFT_PO, position2078)
					}
					goto l2069
				l2077:
					position, tokenIndex = position2069, tokenIndex2069
					{
						position2085 := position
						{
							position2086, tokenIndex2086 := position, tokenIndex
							if buffer[position] != rune('p') {
								goto l2087
							}
							position++
							goto l2086
						l2087:
							position, tokenIndex = position2086, tokenIndex2086
							if buffer[position] != rune('P') {
								goto l2084
							}
							position++
						}
					l2086:
						{
							position2088, tokenIndex2088 := position, tokenIndex
							if buffer[position] != rune('e') {
								goto l2089
							}
							position++
							goto l2088
						l2089:
							position, tokenIndex = position2088, tokenIndex2088
							if buffer[position] != rune('E') {
								goto l2084
							}
							position++
						}
					l2088:
						{
							add(ruleAction109, position)
						}
						add(ruleFT_PE, position2085)
					}
					goto l2069
				l2084:
					position, tokenIndex = position2069, tokenIndex2069
					{
						switch buffer[position] {
						case 'M', 'm':
							{
								position2092 := position
								{
									position2093, tokenIndex2093 := position, tokenIndex
									if buffer[position] != rune('m') {
										goto l2094
									}
									position++
									goto l2093
								l2094:
									position, tokenIndex = position2093, tokenIndex2093
									if buffer[position] != rune('M') {
										goto l2067
									}
									position++
								}
							l2093:
								{
									add(ruleAction111, position)
								}
								add(ruleFT_M, position2092)
							}
							break
						case 'P', 'p':
							{
								position2096 := position
								{
									position2097, tokenIndex2097 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l2098
									}
									position++
									goto l2097
								l2098:
									position, tokenIndex = position2097, tokenIndex2097
									if buffer[position] != rune('P') {
										goto l2067
									}
									position++
								}
							l2097:
								{
									add(ruleAction110, position)
								}
								add(ruleFT_P, position2096)
							}
							break
						case 'C', 'c':
							{
								position2100 := position
								{
									position2101, tokenIndex2101 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2102
									}
									position++
									goto l2101
								l2102:
									position, tokenIndex = position2101, tokenIndex2101
									if buffer[position] != rune('C') {
										goto l2067
									}
									position++
								}
							l2101:
								{
									add(ruleAction107, position)
								}
								add(ruleFT_C, position2100)
							}
							break
						case 'N', 'n':
							{
								position2104 := position
								{
									position2105, tokenIndex2105 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l2106
									}
									position++
									goto l2105
								l2106:
									position, tokenIndex = position2105, tokenIndex2105
									if buffer[position] != rune('N') {
										goto l2067
									}
									position++
								}
							l2105:
								{
									position2107, tokenIndex2107 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l2108
									}
									position++
									goto l2107
								l2108:
									position, tokenIndex = position2107, tokenIndex2107
									if buffer[position] != rune('C') {
										goto l2067
									}
									position++
								}
							l2107:
								{
									add(ruleAction106, position)
								}
								add(ruleFT_NC, position2104)
							}
							break
						default:
							{
								position2110 := position
								{
									position2111, tokenIndex2111 := position, tokenIndex
									if buffer[position] != rune('z') {
										goto l2112
									}
									position++
									goto l2111
								l2112:
									position, tokenIndex = position2111, tokenIndex2111
									if buffer[position] != rune('Z') {
										goto l2067
									}
									position++
								}
							l2111:
								{
									add(ruleAction105, position)
								}
								add(ruleFT_Z, position2110)
							}
							break
						}
					}

				}
			l2069:
				add(rulecc, position2068)
			}
			return true
		l2067:
			position, tokenIndex = position2067, tokenIndex2067
			return false
		},
		/* 157 FT_NZ <- <(('n' / 'N') ('z' / 'Z') Action104)> */
		nil,
		/* 158 FT_Z <- <(('z' / 'Z') Action105)> */
		nil,
		/* 159 FT_NC <- <(('n' / 'N') ('c' / 'C') Action106)> */
		nil,
		/* 160 FT_C <- <(('c' / 'C') Action107)> */
		nil,
		/* 161 FT_PO <- <(('p' / 'P') ('o' / 'O') Action108)> */
		nil,
		/* 162 FT_PE <- <(('p' / 'P') ('e' / 'E') Action109)> */
		nil,
		/* 163 FT_P <- <(('p' / 'P') Action110)> */
		nil,
		/* 164 FT_M <- <(('m' / 'M') Action111)> */
		nil,
		/* 166 Action0 <- <{ p.Emit() }> */
		nil,
		nil,
		/* 168 Action1 <- <{ p.Label(buffer[begin:end])}> */
		nil,
		/* 169 Action2 <- <{ p.LD8() }> */
		nil,
		/* 170 Action3 <- <{ p.LD16() }> */
		nil,
		/* 171 Action4 <- <{ p.Push() }> */
		nil,
		/* 172 Action5 <- <{ p.Pop() }> */
		nil,
		/* 173 Action6 <- <{ p.Ex() }> */
		nil,
		/* 174 Action7 <- <{ p.Inc8() }> */
		nil,
		/* 175 Action8 <- <{ p.Inc8() }> */
		nil,
		/* 176 Action9 <- <{ p.Inc16() }> */
		nil,
		/* 177 Action10 <- <{ p.Dec8() }> */
		nil,
		/* 178 Action11 <- <{ p.Dec8() }> */
		nil,
		/* 179 Action12 <- <{ p.Dec16() }> */
		nil,
		/* 180 Action13 <- <{ p.Add16() }> */
		nil,
		/* 181 Action14 <- <{ p.Adc16() }> */
		nil,
		/* 182 Action15 <- <{ p.Sbc16() }> */
		nil,
		/* 183 Action16 <- <{ p.Dst8() }> */
		nil,
		/* 184 Action17 <- <{ p.Src8() }> */
		nil,
		/* 185 Action18 <- <{ p.Loc8() }> */
		nil,
		/* 186 Action19 <- <{ p.Copy8() }> */
		nil,
		/* 187 Action20 <- <{ p.Loc8() }> */
		nil,
		/* 188 Action21 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 189 Action22 <- <{ p.R8(buffer[begin:end]) }> */
		nil,
		/* 190 Action23 <- <{ p.Dst16() }> */
		nil,
		/* 191 Action24 <- <{ p.Src16() }> */
		nil,
		/* 192 Action25 <- <{ p.Loc16() }> */
		nil,
		/* 193 Action26 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 194 Action27 <- <{ p.R16(buffer[begin:end]) }> */
		nil,
		/* 195 Action28 <- <{ p.R16Contents() }> */
		nil,
		/* 196 Action29 <- <{ p.IR16Contents() }> */
		nil,
		/* 197 Action30 <- <{ p.DispDecimal(buffer[begin:end]) }> */
		nil,
		/* 198 Action31 <- <{ p.DispHex(buffer[begin:end]) }> */
		nil,
		/* 199 Action32 <- <{ p.Disp0xHex(buffer[begin:end]) }> */
		nil,
		/* 200 Action33 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 201 Action34 <- <{ p.Nhex(buffer[begin:end]) }> */
		nil,
		/* 202 Action35 <- <{ p.Ndec(buffer[begin:end]) }> */
		nil,
		/* 203 Action36 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 204 Action37 <- <{ p.NNhex(buffer[begin:end]) }> */
		nil,
		/* 205 Action38 <- <{ p.NNContents() }> */
		nil,
		/* 206 Action39 <- <{ p.Accum("ADD") }> */
		nil,
		/* 207 Action40 <- <{ p.Accum("ADC") }> */
		nil,
		/* 208 Action41 <- <{ p.Accum("SUB") }> */
		nil,
		/* 209 Action42 <- <{ p.Accum("SBC") }> */
		nil,
		/* 210 Action43 <- <{ p.Accum("AND") }> */
		nil,
		/* 211 Action44 <- <{ p.Accum("XOR") }> */
		nil,
		/* 212 Action45 <- <{ p.Accum("OR") }> */
		nil,
		/* 213 Action46 <- <{ p.Accum("CP") }> */
		nil,
		/* 214 Action47 <- <{ p.Rot("RLC") }> */
		nil,
		/* 215 Action48 <- <{ p.Rot("RRC") }> */
		nil,
		/* 216 Action49 <- <{ p.Rot("RL") }> */
		nil,
		/* 217 Action50 <- <{ p.Rot("RR") }> */
		nil,
		/* 218 Action51 <- <{ p.Rot("SLA") }> */
		nil,
		/* 219 Action52 <- <{ p.Rot("SRA") }> */
		nil,
		/* 220 Action53 <- <{ p.Rot("SLL") }> */
		nil,
		/* 221 Action54 <- <{ p.Rot("SRL") }> */
		nil,
		/* 222 Action55 <- <{ p.Bit() }> */
		nil,
		/* 223 Action56 <- <{ p.Res() }> */
		nil,
		/* 224 Action57 <- <{ p.Set() }> */
		nil,
		/* 225 Action58 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 226 Action59 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 227 Action60 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 228 Action61 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 229 Action62 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 230 Action63 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 231 Action64 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 232 Action65 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 233 Action66 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 234 Action67 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 235 Action68 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 236 Action69 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 237 Action70 <- <{ p.Simple(buffer[begin:end]) }> */
		nil,
		/* 238 Action71 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 239 Action72 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 240 Action73 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 241 Action74 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 242 Action75 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 243 Action76 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 244 Action77 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 245 Action78 <- <{ p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 246 Action79 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 247 Action80 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 248 Action81 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 249 Action82 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 250 Action83 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 251 Action84 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 252 Action85 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 253 Action86 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 254 Action87 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 255 Action88 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 256 Action89 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 257 Action90 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 258 Action91 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 259 Action92 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 260 Action93 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 261 Action94 <- <{  p.EDSimple(buffer[begin:end]) }> */
		nil,
		/* 262 Action95 <- <{ p.Rst() }> */
		nil,
		/* 263 Action96 <- <{ p.Call() }> */
		nil,
		/* 264 Action97 <- <{ p.Ret() }> */
		nil,
		/* 265 Action98 <- <{ p.Jp() }> */
		nil,
		/* 266 Action99 <- <{ p.Jr() }> */
		nil,
		/* 267 Action100 <- <{ p.Djnz() }> */
		nil,
		/* 268 Action101 <- <{ p.In() }> */
		nil,
		/* 269 Action102 <- <{ p.Out() }> */
		nil,
		/* 270 Action103 <- <{ p.ODigit(buffer[begin:end]) }> */
		nil,
		/* 271 Action104 <- <{ p.Conditional(Not{FT_Z}) }> */
		nil,
		/* 272 Action105 <- <{ p.Conditional(FT_Z) }> */
		nil,
		/* 273 Action106 <- <{ p.Conditional(Not{FT_C}) }> */
		nil,
		/* 274 Action107 <- <{ p.Conditional(FT_C) }> */
		nil,
		/* 275 Action108 <- <{ p.Conditional(FT_PO) }> */
		nil,
		/* 276 Action109 <- <{ p.Conditional(FT_PE) }> */
		nil,
		/* 277 Action110 <- <{ p.Conditional(FT_P) }> */
		nil,
		/* 278 Action111 <- <{ p.Conditional(FT_M) }> */
		nil,
	}
	p.rules = _rules
}
