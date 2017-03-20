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

type Simple byte

const (
	HALT Simple = 0x76
)

func (s Simple) String() string {
	switch s {
	case HALT:
		return "HALT"
	default:
		panic(fmt.Sprintf("Unknown simple instruction: %02X", byte(s)))
	}
}
