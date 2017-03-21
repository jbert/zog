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
	a, b R16
}

func (ex *EX) String() string {
	return fmt.Sprintf("EX %s, %s", ex.a, ex.b)
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
	default:
		panic(fmt.Sprintf("Unknown simple instruction: %02X", byte(s)))
	}
}
