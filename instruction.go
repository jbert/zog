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

func NewLD8(dst Dst8, src Src8) *LD8 {
	return &LD8{InstBin8{src: src, dst: dst}}
}

func (l *LD8) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}
func (l *LD8) Encode() []byte {
	return []byte{}
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
	if i.eTable != tableR {
		panic("Non-tableR INC8")
	}
	b := encodeXYZ(0, i.idxTable, 4)
	fmt.Printf("JB - b %d x %d y %d z %d\n", b, 0, i.idxTable, 4)
	return i.encodeHelper(b)
}

type DEC8 struct {
	l Loc8
}

func (d *DEC8) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}
func (d *DEC8) Encode() []byte {
	return []byte{}
}

type LD16 struct {
	dst Dst16
	src Src16
}

func (l *LD16) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}
func (l *LD16) Encode() []byte {
	return []byte{}
}

type ADD16 struct {
	dst Dst16
	src Src16
}

func (a *ADD16) String() string {
	return fmt.Sprintf("ADD %s, %s", a.dst, a.src)
}
func (a *ADD16) Encode() []byte {
	return []byte{}
}

type ADC16 struct {
	dst Dst16
	src Src16
}

func (a *ADC16) String() string {
	return fmt.Sprintf("ADC %s, %s", a.dst, a.src)
}
func (a *ADC16) Encode() []byte {
	return []byte{}
}

type SBC16 struct {
	dst Dst16
	src Src16
}

func (s *SBC16) String() string {
	return fmt.Sprintf("SBC %s, %s", s.dst, s.src)
}
func (s *SBC16) Encode() []byte {
	return []byte{}
}

type INC16 struct {
	l Loc16
}

func (i *INC16) String() string {
	return fmt.Sprintf("INC %s", i.l)
}
func (i *INC16) Encode() []byte {
	return []byte{}
}

type DEC16 struct {
	l Loc16
}

func (d *DEC16) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}
func (d *DEC16) Encode() []byte {
	return []byte{}
}

type EX struct {
	dst Dst16
	src Src16
}

func (ex *EX) String() string {
	return fmt.Sprintf("EX %s, %s", ex.dst, ex.src)
}
func (ex *EX) Encode() []byte {
	return []byte{}
}

type DJNZ struct {
	d Disp
}

func (d *DJNZ) String() string {
	return fmt.Sprintf("DJNZ %s", d.d)
}
func (d *DJNZ) Encode() []byte {
	return []byte{}
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
func (jr *JR) Encode() []byte {
	return []byte{}
}

type JP struct {
	c    Conditional
	addr Src16
}

func (jp *JP) String() string {
	if jp.c == True || jp.c == nil {
		return fmt.Sprintf("JP %s", jp.addr)
	} else {
		return fmt.Sprintf("JP %s, %s", jp.c, jp.addr)
	}
}
func (jp *JP) Encode() []byte {
	return []byte{}
}

type CALL struct {
	c    Conditional
	addr Src16
}

func (c *CALL) String() string {
	if c.c == True || c.c == nil {
		return fmt.Sprintf("CALL %s", c.addr)
	} else {
		return fmt.Sprintf("CALL %s, %s", c.c, c.addr)
	}
}
func (c *CALL) Encode() []byte {
	return []byte{}
}

type OUT struct {
	port  Src8
	value Src8
}

func (o *OUT) String() string {
	return fmt.Sprintf("OUT (%s), %s", o.port, o.value)
}
func (o *OUT) Encode() []byte {
	return []byte{}
}

type IN struct {
	dst  Dst8
	port Src8
}

func (i *IN) String() string {
	return fmt.Sprintf("IN %s, (%s)", i.dst, i.port)
}
func (i *IN) Encode() []byte {
	return []byte{}
}

type PUSH struct {
	src Src16
}

func (p *PUSH) String() string {
	return fmt.Sprintf("PUSH %s", p.src)
}
func (p *PUSH) Encode() []byte {
	return []byte{}
}

type POP struct {
	dst Dst16
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
	return []byte{}
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
	return []byte{}
}

func NewAccum(name string, src Src8) *accum {
	// TODO: lookup func by name, panic on unknown
	return &accum{name: name, src: src}
}

type accumFunc func(a, b byte) byte
type accum struct {
	//	f    AccumFunc
	src  Src8
	name string
}

func (a accum) String() string {
	switch a.name {
	case "ADD", "ADC", "SBC":
		return fmt.Sprintf("%s A, %s", a.name, a.src)
	default:
		return fmt.Sprintf("%s %s", a.name, a.src)
	}
}
func (a accum) Encode() []byte {
	return []byte{}
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
	num byte
	r   Loc8
}

func (b *BIT) String() string {
	return fmt.Sprintf("BIT %d, %s", b.num, b.r)
}
func (b *BIT) Encode() []byte {
	return []byte{}
}

type RES struct {
	num byte
	r   Loc8
}

func (r *RES) String() string {
	return fmt.Sprintf("RES %d, %s", r.num, r.r)
}
func (r *RES) Encode() []byte {
	return []byte{}
}

type SET struct {
	num byte
	r   Loc8
}

func (s *SET) String() string {
	return fmt.Sprintf("SET %d, %s", s.num, s.r)
}
func (s *SET) Encode() []byte {
	return []byte{}
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
