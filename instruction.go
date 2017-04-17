package zog

import (
	"fmt"
	"strings"
)

type Instruction interface {
	String() string
	Encode() []byte
}

type LD8 struct {
	InstBin8
}

func NewLD8(dst Loc8, src Loc8) *LD8 {
	return &LD8{InstBin8{dst: dst, src: src}}
}

func (l *LD8) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}
func (l *LD8) Encode() []byte {
	l.inspect()
	if l.dstInfo.ltype != tableR {
		panic("Non-tableR dst in LD8")
	}
	switch l.srcInfo.ltype {
	case tableR:
		b := encodeXYZ(1, l.dstInfo.idxTable, l.srcInfo.idxTable)
		fmt.Printf("JB x %d y %d z %d: b %02X\n", 1, l.dstInfo.idxTable, l.srcInfo.idxTable, b)
		return idxEncodeHelper([]byte{b}, l.idx)
	case Immediate:
		b := encodeXYZ(0, l.dstInfo.idxTable, 6)
		fmt.Printf("JB x %d y %d z %d: b %02X\n", 1, l.dstInfo.idxTable, 6, b)
		return idxEncodeHelper([]byte{b, l.srcInfo.imm8}, l.idx)
	default:
		panic("Unknown src type in LD8")
	}
}

type INC8 struct {
	InstU8
}

func NewINC8(l Loc8) *INC8 {
	return &INC8{InstU8{l: l}}
}
func (i *INC8) String() string {
	return fmt.Sprintf("INC %s", i.l)
}
func (i *INC8) Encode() []byte {
	i.inspect()
	if i.lInfo.ltype != tableR {
		panic("Non-tableR INC8")
	}
	b := encodeXYZ(0, i.lInfo.idxTable, 4)
	return idxEncodeHelper([]byte{b}, i.idx)
}

type DEC8 struct {
	InstU8
}

func NewDEC8(l Loc8) *DEC8 {
	return &DEC8{InstU8{l: l}}
}
func (d *DEC8) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}
func (d *DEC8) Encode() []byte {
	d.inspect()
	if d.lInfo.ltype != tableR {
		panic("Non-tableR DEC8")
	}
	b := encodeXYZ(0, d.lInfo.idxTable, 5)
	return idxEncodeHelper([]byte{b}, d.idx)
}

type LD16 struct {
	InstBin16
}

func NewLD16(dst, src Loc16) *LD16 {
	return &LD16{InstBin16: InstBin16{dst: dst, src: src}}
}
func (l *LD16) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}
func (l *LD16) Encode() []byte {
	l.inspect()
	if l.dstInfo.ltype != tableRP {
		panic("Non-tableRP dst in LD16")
	}

	if l.dst == SP {
		if !l.srcInfo.isHLLike() {
			panic("Trying to load non-HL like into SP")
		}
		buf := []byte{encodeXPQZ(3, 3, 1, 1)}
		return idxEncodeHelper(buf, l.idx)
	}

	switch l.srcInfo.ltype {
	case Immediate:
		buf := []byte{encodeXPQZ(0, l.dstInfo.idxTable, 0, 1)}
		buf = append(buf, l.srcInfo.imm16...)
		return idxEncodeHelper(buf, l.idx)
	default:
		panic("Unknown src type in LD16")
	}
}

type ADD16 struct {
	InstBin16
}

func NewADD16(dst, src Loc16) *ADD16 {
	return &ADD16{InstBin16: InstBin16{dst: dst, src: src}}
}
func (a *ADD16) String() string {
	return fmt.Sprintf("ADD %s, %s", a.dst, a.src)
}
func (a *ADD16) Encode() []byte {
	a.inspect()
	if a.dstInfo.ltype != tableRP {
		panic("Non-tableRP dst in ADD16")
	}
	if a.srcInfo.ltype != tableRP {
		panic("Non-tableRP src in ADD16")
	}

	// TODO: support other ADD16
	if !a.dstInfo.isHLLike() {
		panic("Non-HL dst in ADD16")
	}
	switch a.srcInfo.ltype {
	case tableRP:
		buf := []byte{encodeXPQZ(0, a.srcInfo.idxTable, 1, 1)}
		return idxEncodeHelper(buf, a.idx)
	default:
		panic("Unknown src type in ADD16")
	}
}

type ADC16 struct {
	dst Loc16
	src Loc16
}

func (a *ADC16) String() string {
	return fmt.Sprintf("ADC %s, %s", a.dst, a.src)
}
func (a *ADC16) Encode() []byte {
	return []byte{}
}

type SBC16 struct {
	dst Loc16
	src Loc16
}

func (s *SBC16) String() string {
	return fmt.Sprintf("SBC %s, %s", s.dst, s.src)
}
func (s *SBC16) Encode() []byte {
	return []byte{}
}

type INC16 struct {
	InstU16
}

func NewINC16(l Loc16) *INC16 {
	return &INC16{InstU16{l: l}}
}

func (i *INC16) String() string {
	return fmt.Sprintf("INC %s", i.l)
}
func (i *INC16) Encode() []byte {
	i.inspect()
	if i.lInfo.ltype != tableRP {
		panic("Non-tableRP INC16")
	}
	b := encodeXPQZ(0, i.lInfo.idxTable, 0, 3)
	return idxEncodeHelper([]byte{b}, i.idx)
}

type DEC16 struct {
	InstU16
}

func NewDEC16(l Loc16) *DEC16 {
	return &DEC16{InstU16{l: l}}
}
func (d *DEC16) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}
func (d *DEC16) Encode() []byte {
	d.inspect()
	if d.lInfo.ltype != tableRP {
		panic("Non-tableRP DEC16")
	}
	b := encodeXPQZ(0, d.lInfo.idxTable, 1, 3)
	return idxEncodeHelper([]byte{b}, d.idx)
}

type EX struct {
	InstBin16
}

func NewEX(dst, src Loc16) *EX {
	return &EX{InstBin16: InstBin16{dst: dst, src: src}}
}

func (ex *EX) String() string {
	return fmt.Sprintf("EX %s, %s", ex.dst, ex.src)
}
func (ex *EX) Encode() []byte {
	if ex.dst == AF && ex.src == AF_PRIME {
		return []byte{0x08}
	} else if ex.dst.String() == (Contents{SP}).String() {

		var info loc16Info
		var idx idxInfo
		inspectLoc16(ex.src, &info, &idx, false)
		buf := []byte{encodeXYZ(3, 4, 3)}
		return idxEncodeHelper(buf, idx)
	} else if ex.dst == DE && ex.src == HL {
		// EX DE,HL is an excpetion to the IX/IY rule
		return []byte{encodeXYZ(3, 5, 3)}
	}

	panic("Unrecognised EX instruction")
}

type DJNZ struct {
	d Disp
}

func (d *DJNZ) String() string {
	return fmt.Sprintf("DJNZ %s", d.d)
}
func (d *DJNZ) Encode() []byte {
	b := encodeXYZ(0, 2, 0)
	return []byte{b, byte(d.d)}
}

type JR struct {
	c Conditional
	d Disp
}

func (j *JR) String() string {
	if j.c == True || j.c == nil {
		return fmt.Sprintf("JR %s", j.d)
	} else {
		return fmt.Sprintf("JR %s, %s", j.c, j.d)
	}
}
func (j *JR) Encode() []byte {
	var y byte
	if j.c == True || j.c == nil {
		y = 3
	} else {
		y = findInTableCC(j.c)
		y += 4
	}
	b := encodeXYZ(0, y, 0)
	return []byte{b, byte(j.d)}
}

type JP struct {
	InstU16
	c Conditional
}

func NewJP(c Conditional, l Loc16) *JP {
	return &JP{InstU16: InstU16{l: l}, c: c}
}

func (jp *JP) String() string {
	if jp.c == True || jp.c == nil {
		return fmt.Sprintf("JP %s", jp.l)
	} else {
		return fmt.Sprintf("JP %s, %s", jp.c, jp.l)
	}
}
func (jp *JP) Encode() []byte {
	jp.inspect()
	if jp.c == True || jp.c == nil {
		if jp.lInfo.isHLLike() {
			buf := []byte{encodeXPQZ(3, 2, 1, 1)}
			return idxEncodeHelper(buf, jp.idx)
		}
	}
	if jp.lInfo.ltype != Immediate {
		panic("Non-immediate (or direct HL-like) JP")
	}

	var buf []byte
	if jp.c == True || jp.c == nil {
		buf = []byte{encodeXYZ(3, 0, 3)}
	} else {
		y := findInTableCC(jp.c)
		buf = []byte{encodeXYZ(3, y, 2)}
	}
	buf = append(buf, jp.lInfo.imm16...)
	return buf
}

type CALL struct {
	InstU16
	c Conditional
}

func NewCALL(c Conditional, l Loc16) *CALL {
	return &CALL{InstU16: InstU16{l: l}, c: c}
}
func (c *CALL) String() string {
	if c.c == True || c.c == nil {
		return fmt.Sprintf("CALL %s", c.l)
	} else {
		return fmt.Sprintf("CALL %s, %s", c.c, c.l)
	}
}
func (c *CALL) Encode() []byte {
	c.inspect()
	var buf []byte
	if c.c == nil || c.c == True {
		buf = []byte{encodeXPQZ(3, 0, 1, 5)}
	} else {
		y := findInTableCC(c.c)
		buf = []byte{encodeXYZ(3, y, 4)}
	}
	buf = append(buf, c.lInfo.imm16...)
	return buf
}

type OUT struct {
	port  Loc8
	value Loc8
}

func (o *OUT) String() string {
	return fmt.Sprintf("OUT (%s), %s", o.port, o.value)
}
func (o *OUT) Encode() []byte {
	if o.port == C {
		panic("impl")
	} else {
		imm8 := o.port.(Imm8)
		return []byte{encodeXYZ(3, 2, 3), byte(imm8)}
	}
}

type IN struct {
	dst  Loc8
	port Loc8
}

func (i *IN) String() string {
	return fmt.Sprintf("IN %s, (%s)", i.dst, i.port)
}
func (i *IN) Encode() []byte {
	if i.port == C {
		panic("impl")
	} else {
		imm8 := i.port.(Imm8)
		return []byte{encodeXYZ(3, 3, 3), byte(imm8)}
	}
}

type PUSH struct {
	InstU16
}

func NewPUSH(l Loc16) *PUSH {
	return &PUSH{InstU16{l: l}}
}
func (p *PUSH) String() string {
	return fmt.Sprintf("PUSH %s", p.l)
}
func (p *PUSH) Encode() []byte {
	p.inspectRP2()
	if p.lInfo.ltype != tableRP2 {
		panic("Non-tableRP PUSH")
	}
	return []byte{encodeXPQZ(3, p.lInfo.idxTable, 0, 5)}
}

type POP struct {
	dst Loc16
}

func (p *POP) String() string {
	return fmt.Sprintf("POP %s", p.dst)
}
func (p *POP) Encode() []byte {
	return []byte{}
}

type RST struct {
	addr byte
}

func (r *RST) String() string {
	return fmt.Sprintf("RST %d", r.addr)
}
func (r *RST) Encode() []byte {
	y := r.addr / 8
	return []byte{encodeXYZ(3, y, 7)}
}

type RET struct {
	c Conditional
}

func (r *RET) String() string {
	if r.c == True || r.c == nil {
		return "RET"
	} else {
		return fmt.Sprintf("RET %s", r.c)
	}
}
func (r *RET) Encode() []byte {
	if r.c == True || r.c == nil {
		return []byte{encodeXPQZ(3, 0, 1, 1)}
	}
	y := findInTableCC(r.c)
	return []byte{encodeXYZ(3, y, 0)}
}

func NewAccum(name string, l Loc8) *accum {
	// TODO: lookup func by name, panic on unknown
	return &accum{name: name, InstU8: InstU8{l: l}}
}

type accumFunc func(a, b byte) byte
type accum struct {
	//	f    AccumFunc
	InstU8
	name string
}

func (a accum) String() string {
	switch a.name {
	case "ADD", "ADC", "SBC":
		return fmt.Sprintf("%s A, %s", a.name, a.l)
	default:
		return fmt.Sprintf("%s %s", a.name, a.l)
	}
}
func (a accum) Encode() []byte {
	a.inspect()
	y := findInTableALU(a.name)
	var buf []byte
	switch a.lInfo.ltype {
	case tableR:
		buf = []byte{encodeXYZ(2, y, a.lInfo.idxTable)}
	case Immediate:
		buf = []byte{encodeXYZ(3, y, 6)}
		buf = append(buf, a.lInfo.imm8)
	default:
		panic("Unknown accum location type")
	}
	return idxEncodeHelper(buf, a.idx)
}

type rot struct {
	name string
	r    Loc8
}

func NewRot(name string, r Loc8) *rot {
	return &rot{name: name, r: r}
}

func (r *rot) String() string {
	return fmt.Sprintf("%s %s", r.name, r.r)
}
func (r *rot) Encode() []byte {
	return []byte{}
}

type BIT struct {
	InstU8
	num byte
}

func NewBIT(num byte, l Loc8) *BIT {
	return &BIT{InstU8: InstU8{l: l}, num: num}
}
func (b *BIT) String() string {
	return fmt.Sprintf("BIT %d, %s", b.num, b.l)
}
func (b *BIT) Encode() []byte {
	b.inspect()
	if b.lInfo.ltype != tableR {
		panic("Non-tableR src in BIT")
	}
	enc := encodeXYZ(1, b.num, b.lInfo.idxTable)
	return idxEncodeHelper([]byte{0xcb, enc}, b.idx)
}

type RES struct {
	InstU8
	num byte
}

func NewRES(num byte, l Loc8) *RES {
	return &RES{InstU8: InstU8{l: l}, num: num}
}
func (r *RES) String() string {
	return fmt.Sprintf("RES %d, %s", r.num, r.l)
}
func (r *RES) Encode() []byte {
	r.inspect()
	if r.lInfo.ltype != tableR {
		panic("Non-tableR src in BIT")
	}
	enc := encodeXYZ(2, r.num, r.lInfo.idxTable)
	return idxEncodeHelper([]byte{0xcb, enc}, r.idx)
}

type SET struct {
	InstU8
	num byte
}

func NewSET(num byte, l Loc8) *SET {
	return &SET{InstU8: InstU8{l: l}, num: num}
}
func (s *SET) String() string {
	return fmt.Sprintf("SET %d, %s", s.num, s.l)
}
func (s *SET) Encode() []byte {
	s.inspect()
	if s.lInfo.ltype != tableR {
		panic("Non-tableR src in BIT")
	}
	enc := encodeXYZ(3, s.num, s.lInfo.idxTable)
	return idxEncodeHelper([]byte{0xcb, enc}, s.idx)
}

type Simple byte

const (
	NOP Simple = 0x00

	HALT Simple = 0x76

	RLCA Simple = 0x07
	RRCA Simple = 0x0f
	RLA  Simple = 0x17
	RRA  Simple = 0x1f
	DAA  Simple = 0x27
	CPL  Simple = 0x2f
	SCF  Simple = 0x37
	CCF  Simple = 0x3f

	EXX Simple = 0xd9

	DI Simple = 0xf3
	EI Simple = 0xfb
)

type simpleName struct {
	inst Simple
	name string
}

var simpleNames []simpleName = []simpleName{
	{NOP, "NOP"},

	{HALT, "HALT"},

	{RLCA, "RLCA"},
	{RRCA, "RRCA"},
	{RLA, "RLA"},
	{RRA, "RRA"},
	{DAA, "DAA"},
	{CPL, "CPL"},
	{SCF, "SCF"},
	{CCF, "CCF"},

	{EXX, "EXX"},

	{DI, "DI"},
	{EI, "EI"},
}

func (s Simple) String() string {

	for _, simpleName := range simpleNames {
		if simpleName.inst == s {
			return simpleName.name
		}
	}
	panic(fmt.Sprintf("Unknown simple instruction: %02X", byte(s)))
}
func (s Simple) Encode() []byte {
	return []byte{byte(s)}
}

func LookupSimpleName(name string) Simple {
	name = strings.ToUpper(name)
	for _, simpleName := range simpleNames {
		if simpleName.name == name {
			return simpleName.inst
		}
	}
	panic(fmt.Errorf("Unrecognised Simple instruction name : [%s]", name))
}

type EDSimple byte

const (
	NEG  EDSimple = 0x44
	RETN EDSimple = 0x45
	RETI EDSimple = 0x4d

	RRD EDSimple = 0x67
	RLD EDSimple = 0x6f

	IM0 EDSimple = 0x46
	IM1 EDSimple = 0x56
	IM2 EDSimple = 0x5e

	LDI  EDSimple = 0xa0
	CPI  EDSimple = 0xa1
	LDD  EDSimple = 0xa8
	CPD  EDSimple = 0xa9
	LDIR EDSimple = 0xb0
	CPIR EDSimple = 0xb1
	LDDR EDSimple = 0xb8
	CPDR EDSimple = 0xb9

	INI  EDSimple = 0xa2
	OUTI EDSimple = 0xa3
	IND  EDSimple = 0xaa
	OUTD EDSimple = 0xab
	INIR EDSimple = 0xb2
	OTIR EDSimple = 0xb3
	INDR EDSimple = 0xba
	OTDR EDSimple = 0xbb
)

type edSimpleName struct {
	inst EDSimple
	name string
}

var EDSimpleNames []edSimpleName = []edSimpleName{
	{NEG, "NEG"},
	{RETN, "RETN"},
	{RETI, "RETI"},
	{RRD, "RRD"},
	{RLD, "RLD"},
	{IM0, "IM 0"},
	{IM1, "IM 1"},
	{IM2, "IM 2"},

	{LDI, "LDI"},
	{CPI, "CPI"},
	{LDD, "LDD"},
	{CPD, "CPD"},
	{LDIR, "LDIR"},
	{CPIR, "CPIR"},
	{LDDR, "LDDR"},
	{CPDR, "CPDR"},

	{INI, "INI"},
	{OUTI, "OUTI"},
	{IND, "IND"},
	{OUTD, "OUTD"},
	{INIR, "INIR"},
	{OTIR, "OTIR"},
	{INDR, "INDR"},
	{OTDR, "OTDR"},
}

func (s EDSimple) String() string {

	for _, simpleName := range EDSimpleNames {
		if simpleName.inst == s {
			return simpleName.name
		}
	}
	panic(fmt.Sprintf("Unknown EDSimple instruction: %02X", byte(s)))
}

func (s EDSimple) Encode() []byte {
	return []byte{byte(s)}
}

func LookupEDSimpleName(name string) EDSimple {
	name = strings.ToUpper(name)
	for _, simpleName := range EDSimpleNames {
		if simpleName.name == name {
			return simpleName.inst
		}
	}
	panic(fmt.Errorf("Unrecognised EDSimple instruction name : [%s]", name))
}
