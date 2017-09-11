package speccy

import (
	"fmt"
	"os"

	"github.com/jbert/zog"
)

type Machine struct {
}

func NewMachine() *Machine {
	return &Machine{}
}

func (m Machine) LoadAddr() uint16 {
	return 0x8000
}

func (m Machine) RunAddr() uint16 {
	return 0x8000
}

func (m Machine) Name() string {
	return "speccy"
}

var printAssembly = `
	; Function to call is in C
	; Func 2 => Print ASCII code of reg E to console
	; Func 9 => Print ASCII string starting at DE until $ to console
	LD A, 2
	CP C
	JP NZ, next1
	CALL printchar
	RET
next1:
	LD A, 9
	CP C
	JP NZ, next2
	CALL printstr
	RET
next2:
	HALT
; Print char in E to console
printchar:
	PUSH BC
	LD BC, 0ffffh
	OUT (C), E
	POP BC
	RET
; Print $-terminated string at DE to console
printstr:
	PUSH HL
	PUSH DE
  POP HL
	LD A, 24h		; '$'
printstr_nextchar:
	CP (HL)
	JP Z, printstr_end
	LD E, (HL)
	CALL printchar
	INC HL
	JP printstr_nextchar
printstr_end:
	POP HL
	RET
`

func (m *Machine) Load(z *zog.Zog) error {
	sPS := speccyPrintState{}
	z.RegisterOutputHandler(0xffff, sPS.speccyPrintByte)

	// We only use RST 16
	zeroPageAssembly, err := zog.Assemble(`
	ORG 0000h
	HALT
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	NOP
	; One entry point at 10h (RST 16), to print char in A
	PUSH DE
	LD E, A
	call printchar
	POP DE
	RET
` + printAssembly)
	if err != nil {
		return fmt.Errorf("Failed to assemble prelude: %s", err)
	}
	err = z.Load(zeroPageAssembly)
	if err != nil {
		return fmt.Errorf("Load zero page assembly: %s", err)
	}

	chanOpenAssembly, err := zog.Assemble(`
	ORG 1601h
	RET
`)
	if err != nil {
		return fmt.Errorf("Failed to assemble chan-open: %s", err)
	}
	err = z.Load(chanOpenAssembly)
	if err != nil {
		return fmt.Errorf("Load chan open assembly: %s", err)
	}

	return nil
}


const (
	BRIGHT = 0x13
	AT = 0x16
	TAB = 0x17
)
type speccyPrintState struct {
	row, column int
	n byte		// at/tab control char
	a byte		// First byte in a two-byte AT/TAB control
	haveA bool
}
func (sps *speccyPrintState) clear() {
	sps.n = 0
	sps.a = 0
	sps.haveA = false
}
func (sps speccyPrintState) wantA() bool {
	return !sps.haveA && (sps.n == TAB || sps.n == AT)
}
func (sps speccyPrintState) wantB() bool {
	return sps.haveA
}

func (sps *speccyPrintState) speccyPrintByte(n byte) {
fmt.Printf("JB [%02X] [%c] col [%d]\n", n, n, sps.column)
	if sps.wantA() {
		sps.a = n 
		sps.haveA = true
		return
	}
	if sps.wantB() {
		if sps.n == TAB {
			col := (int(sps.a) + (256 + int(n))) % 32
			if col < sps.column {
				printRune('\n')
				sps.column = 0
				sps.row++
			}
			nSpaces := col -sps.column
//			fmt.Printf("JB col [%d] nspaces [%d]\n", col, nSpaces)
			for i := 0; i < nSpaces; i++ {
				printRune(' ')
				sps.column++
			}
			sps.column = col
			sps.clear()
			return
		} else {
			panic("TODO - impl")
		}
	}
	// It's mostly ASCII
	r := rune(n)

	// Translate byte sequences
	switch r {
	case 0x06:
		r = ','
	case 0x0d:
		r = 0x0a
		sps.column = 0
	case BRIGHT, AT, TAB:
		sps.n = n
		return
	case 0x5e:
		r = '↑'
	case 0x7f:
		r = '©'
	case '`':
		r = '£'
	default:
	  if r >=  0x00 && r <= 0x0b {
			r = '?'
		} else if r >=  0x20 && r <= 0x7f {
			// Rest of ASCII
		} else {
//			panic(fmt.Sprintf("Unmapped code [%02X]", n))
		}
	}
//	fmt.Printf("JB Printing [%02X]\n", r)
	if sps.column > 32 {
		printRune('\n')
		sps.column = 0
		sps.row++
	}
	printRune(r)
	sps.column++
	sps.clear()
}

func printRune(r rune) {
	fmt.Fprintf(os.Stderr, "%c", r)
}

