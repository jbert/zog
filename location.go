package zog

import "fmt"

type Src8 interface {
	Read8(z *Zog) (byte, error)
	String() string
}
type Dst8 interface {
	Write8(z *Zog, n byte) error
	String() string
}

type Src16 interface {
	Read16(z *Zog) (uint16, error)
	String() string
}
type Dst16 interface {
	Write16(z *Zog, nn uint16) error
	String() string
}

type Loc8 interface {
	Read8(z *Zog) (byte, error)
	Write8(z *Zog, n byte) error
	String() string
}

type Loc16 interface {
	Read16(z *Zog) (uint16, error)
	Write16(z *Zog, nn uint16) error
	String() string
}

type R8 int

const (
	A R8 = iota
	F
	B
	C
	D
	E
	H
	L
)

func (r R8) String() string {
	switch r {

	case A:
		return "A"
	case F:
		return "F"

	case B:
		return "B"
	case C:
		return "C"

	case D:
		return "D"
	case E:
		return "E"

	case H:
		return "H"
	case L:
		return "L"

	default:
		panic(fmt.Errorf("Unrecognised R8 : %d", int(r)))
	}
}

func (r R8) Read8(z *Zog) (byte, error) {
	// TODO: debug
	var n byte
	fmt.Printf("Z: %02X <- %s\n", n, r)
	return n, nil
}
func (r R8) Write8(z *Zog, n byte) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", r, n)
	return nil
}

type R16 int

const (
	AF R16 = iota
	BC
	DE
	HL
	IX
	IY
	SP
)

func (r R16) String() string {
	switch r {

	case AF:
		return "AF"

	case BC:
		return "BC"

	case DE:
		return "DE"

	case HL:
		return "HL"

	case IX:
		return "IX"

	case IY:
		return "IY"

	case SP:
		return "SP"

	default:
		panic(fmt.Errorf("Unrecognised R16 : %d", int(r)))
	}
}

func (r R16) Read16(z *Zog) (uint16, error) {
	// TODO: debug
	var nn uint16
	fmt.Printf("Z: %02X <- %s\n", nn, r)
	return nn, nil
}
func (r R16) Write16(z *Zog, nn uint16) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", r, nn)
	return nil
}

type Contents struct {
	l Loc16
}

func (c Contents) String() string {
	return fmt.Sprintf("(%s)", c.l)
}
func (c Contents) Read8(z *Zog) (byte, error) {
	// TODO: debug
	var n byte
	fmt.Printf("Z: %02X <- %s\n", n, c)
	return n, nil
}
func (c Contents) Write8(z *Zog, n byte) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", c, n)
	return nil
}
