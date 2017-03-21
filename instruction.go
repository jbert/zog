package zog

import "fmt"

type instruction interface {
	String() string
}

type LD8 struct {
	src Src8
	dst Dst8
}

func (l *LD8) String() string {
	return fmt.Sprintf("LD %s, %s", l.src, l.dst)
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
