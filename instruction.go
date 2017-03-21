package zog

import "fmt"

type instruction interface {
	String() string
}

type LD8 struct {
	dst Dst8
	src Src8
}

func (l *LD8) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}

type INC8 struct {
	l Loc8
}

func (i *INC8) String() string {
	return fmt.Sprintf("INC %s", i.l)
}

type DEC8 struct {
	l Loc8
}

func (d *DEC8) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}

type LD16 struct {
	dst Dst16
	src Src16
}

func (l *LD16) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}

type ADD16 struct {
	dst Dst16
	src Src16
}

func (a *ADD16) String() string {
	return fmt.Sprintf("ADD %s, %s", a.dst, a.src)
}

type INC16 struct {
	l Loc16
}

func (i *INC16) String() string {
	return fmt.Sprintf("INC %s", i.l)
}

type DEC16 struct {
	l Loc16
}

func (d *DEC16) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}

type EX struct {
	dst Dst16
	src Src16
}

func (ex *EX) String() string {
	return fmt.Sprintf("EX %s, %s", ex.dst, ex.src)
}

type DJNZ struct {
	d Disp
}

func (d *DJNZ) String() string {
	return fmt.Sprintf("DJNZ %s", d.d)
}

type JR struct {
	c Conditional
	d Disp
}

func (j *JR) String() string {
	if j.c == True {
		return fmt.Sprintf("JR %s", j.d)
	} else {
		return fmt.Sprintf("JR %s, %s", j.c, j.d)
	}
}

type JP struct {
	c    Conditional
	addr Src16
}

func (jp *JP) String() string {
	if jp.c == True {
		return fmt.Sprintf("JP %s", jp.addr)
	} else {
		return fmt.Sprintf("JP %s, %s", jp.c, jp.addr)
	}
}

type CALL struct {
	c    Conditional
	addr Src16
}

func (c *CALL) String() string {
	if c.c == True {
		return fmt.Sprintf("CALL %s", c.addr)
	} else {
		return fmt.Sprintf("CALL %s, %s", c.c, c.addr)
	}
}

type OUT struct {
	port  Src8
	value Src8
}

func (o *OUT) String() string {
	return fmt.Sprintf("OUT (%s), %s", o.port, o.value)
}

type IN struct {
	dst  Dst8
	port Src8
}

func (i *IN) String() string {
	return fmt.Sprintf("IN %s, (%s)", i.dst, i.port)
}

type PUSH struct {
	src Src16
}

func (p *PUSH) String() string {
	return fmt.Sprintf("PUSH %s", p.src)
}

type POP struct {
	dst Dst16
}

func (p *POP) String() string {
	return fmt.Sprintf("POP %s", p.dst)
}

type RST struct {
	addr byte
}

func (r *RST) String() string {
	return fmt.Sprintf("RST %d", r.addr)
}

type RET struct {
	c Conditional
}

func (r *RET) String() string {
	if r.c == True {
		return "RET"
	} else {
		return fmt.Sprintf("RET %s", r.c)
	}
}

type AccumFunc func(a, b byte) byte

type Accum struct {
	//	f    AccumFunc
	src  Src8
	name string
}

func (a Accum) String() string {
	switch a.name {
	case "ADD", "ADC", "SBC":
		return fmt.Sprintf("%s A, %s", a.name, a.src)
	default:
		return fmt.Sprintf("%s %s", a.name, a.src)
	}
}

type ROT struct {
	name string
	r    Loc8
}

func (r *ROT) String() string {
	return fmt.Sprintf("%s %s", r.name, r.r)
}

type BIT struct {
	num byte
	r   Loc8
}

func (b *BIT) String() string {
	return fmt.Sprintf("BIT %d, %s", b.num, b.r)
}

type RES struct {
	num byte
	r   Loc8
}

func (r *RES) String() string {
	return fmt.Sprintf("RES %d, %s", r.num, r.r)
}

type SET struct {
	num byte
	r   Loc8
}

func (s *SET) String() string {
	return fmt.Sprintf("SET %d, %s", s.num, s.r)
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

func (s Simple) String() string {
	switch s {
	case NOP:
		return "NOP"

	case HALT:
		return "HALT"

	case RLCA:
		return "RLCA"
	case RRCA:
		return "RRCA"
	case RLA:
		return "RLA"
	case RRA:
		return "RRA"
	case DAA:
		return "DAA"
	case CPL:
		return "CPL"
	case SCF:
		return "SCF"
	case CCF:
		return "CCF"

	case EXX:
		return "EXX"

	case DI:
		return "DI"
	case EI:
		return "EI"
	default:
		panic(fmt.Sprintf("Unknown simple instruction: %02X", byte(s)))
	}
}
