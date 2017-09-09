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
	err = loadPseudoCPM(z)
	if err != nil {
		log.Fatalf("Failed to load pCPM: %s", err)
	}
	loadAddr := uint16(0x0100)
	runAddr := uint16(0x0100)

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

func printByte(n byte) {
	fmt.Fprintf(os.Stderr, "%c", n)
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

	highAssembly, err := zog.Assemble(`
	ORG 0xf000

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
			`)

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
