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
	return fmt.Sprintf("ADD %s, %s", a.src, a.dst)
}

type EX struct {
	a, b R16
}

func (ex *EX) String() string {
	return fmt.Sprintf("EX %s, %s", ex.a, ex.b)
}

type Simple byte

const (
	NOP  Simple = 0x00
	HALT Simple = 0x76
)

func (s Simple) String() string {
	switch s {
	case NOP:
		return "NOP"
	case HALT:
		return "HALT"
	default:
		panic(fmt.Sprintf("Unknown simple instruction: %02X", byte(s)))
	}
}
