package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	"github.com/jbert/zog"
)

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile `file`")
	trace := flag.String("trace", "", "Trace addresses: start-end,s2-e2")
	watch := flag.String("watch", "", "Watch addresses: start-end,s2-e2")
	haltstate := flag.Bool("haltstate", false, "Print state on halt")
	numhalttrace := flag.Int("halttrace", 0, "Number of traces to print on halt")
	machine := flag.String("machine", "none", "Machine for console printer (none, cpm, spectrum)")

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if flag.NArg() < 1 {
		usage("Missing filename")
	}
	fname := flag.Arg(0)

	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Failed to open file [%s] : %s\n", fname, err)
	}

	z := zog.New(0)
	z.TraceOnHalt(*numhalttrace)

	loadAddr := uint16(0x8000)
	runAddr := uint16(0x8000)

	switch *machine {
	case "cpm":
		fmt.Printf("Loading pseudo CP/M\n")
		err = loadPseudoCPM(z)
		if err != nil {
			log.Fatalf("Failed to load pCPM: %s", err)
		}
		loadAddr = 0x0100
		runAddr = 0x0100
	case "spectrum","speccy":
		fmt.Printf("Loading pseudo spectrum\n")
		err = loadPseudoSpeccy(z)
		if err != nil {
			log.Fatalf("Failed to load pSpeccy: %s", err)
		}
	}

	regions, err := zog.ParseRegions(*trace)
	if err != nil {
		log.Fatalf("Can't parse trace regions [%s]: %s", *trace, err)
	}
	err = z.TraceRegions(regions)
	if err != nil {
		log.Fatalf("Can't add traces [%s]: %s", err)
	}

	// z.Watch16(zog.SP)

	regions, err = zog.ParseRegions(*watch)
	if err != nil {
		log.Fatalf("Can't parse watch regions [%s]: %s", *watch, err)
	}
	err = z.WatchRegions(regions)
	if err != nil {
		log.Fatalf("Can't add watches [%s]: %s", err)
	}

	err = z.RunBytes(loadAddr, buf, runAddr)
	if err != nil {
		log.Fatalf("RunBytes returned error: %s", err)
	}

	if *haltstate {
		fmt.Printf("STATE: %s\n", z.State())
	}
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
			fmt.Printf("JB col [%d] nspaces [%d]\n", col, nSpaces)
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
			panic(fmt.Sprintf("Unmapped code [%02X]", n))
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

func printByte(n byte) {
	printRune(rune(n))
}

func printRune(r rune) {
	fmt.Fprintf(os.Stderr, "%c", r)
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

func loadPseudoSpeccy(z *zog.Zog) error {
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

func loadPseudoCPM(z *zog.Zog) error {
	z.RegisterOutputHandler(0xffff, printByte)
	zeroPageAssembly, err := zog.Assemble(`
	ORG 0000h
	HALT
	NOP			; would be addr of warm start (with JP inst at 0000)
	NOP
	NOP			; The 'intel standard iobyte'? http://www.gaby.de/cpm/manuals/archive/cpm22htm/ch6.htm#Section_6.9
	NOP
	; One entry point at 0005h
	; but this is also "the lowest address used by CP/M"
	; and used to the set the SP (by zexall)
	JP 0xf000
`)
	if err != nil {
		return fmt.Errorf("Failed to assemble prelude: %s", err)
	}
	err = z.Load(zeroPageAssembly)
	if err != nil {
		return fmt.Errorf("Load zero page assembly: %s", err)
	}

	highAssembly, err := zog.Assemble("ORG 0xf000\n" + printAssembly)

	if err != nil {
		return fmt.Errorf("Failed to assemble prelude: %s", err)
	}
	return z.Load(highAssembly)
}

func usage(reason string) {
	fmt.Printf(`%s

%s <filename>

`, reason, os.Args[0])
	os.Exit(1)
}
