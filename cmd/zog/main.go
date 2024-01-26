package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	"github.com/jbert/zog"
	"github.com/jbert/zog/cpm"
	"github.com/jbert/zog/file"
	"github.com/jbert/zog/repl"
	"github.com/jbert/zog/speccy"
)

func getFormat(format string) (file.ImageFormat, error) {
	switch format {
	case "z80":
		return &file.Z80Snapshot{}, nil
	case "sna":
		return &file.SNASnapshot{}, nil
	default:
		return nil, fmt.Errorf("Unsupported format %s\n", format)
	}

}

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile `file`")
	trace := flag.String("trace", "", "Trace addresses: start-end,s2-e2")
	watch := flag.String("watch", "", "Watch addresses: start-end,s2-e2")
	haltstate := flag.Bool("haltstate", false, "Print state on halt")
	numhalttrace := flag.Int("halttrace", 0, "Number of traces to print on halt")
	machineName := flag.String("machine", "none", "Machine for console printer (none, cpm, spectrum)")
	imageFname := flag.String("image", "", "Name of image file (.z80 supported)")
	quiet := flag.Bool("quiet", false, "Suppress messages")
	mode := flag.String("mode", "run", "Operation mode (run, disassemble)")
	format := flag.String("format", "z80", "Image format (z80, sna)")

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

	z := zog.New(0)
	z.TraceOnHalt(*numhalttrace)

	var machine zog.Machine

	switch *machineName {
	case "cpm":
		machine = cpm.NewMachine(z)
	case "spectrum", "speccy":
		machine = speccy.NewMachine(z)
	case "repl":
		machine = repl.NewMachine(z)
	default:
		panic("Specify a machine type")
	}

	if !*quiet {
		fmt.Printf("Loading %s\n", machine.Name())
	}
	err := machine.Start()
	if err != nil {
		log.Fatalf("Failed to load machine %s: %s", machine.Name(), err)
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

	var runErr error

	if *imageFname != "" {
		h, err := getFormat(*format)
		if err != nil {
			panic(err)
		}
		f, err := os.Open(*imageFname)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		err = h.Parse(f)
		if err != nil {
			panic(err)
		}
		err = h.Load(f, z)
		if err != nil {
			panic(err)
		}

		fmt.Printf("z.State(): %v\n", z.State())

		// JB - TODO hack. Elite z80 file has interrupts off
		// (or we aren't parsing it correctly)
		//		z.LoadInterruptState(zog.InterruptState{IFF1: true, IFF2: true, Mode: 1})

		switch *mode {
		case "run":
			runErr = z.Run()
		case "disassemble":
			// Grab some memory from the start point
			reg := z.GetRegisters()
			size := 0x100
			buf, err := z.Mem.PeekBuf(reg.PC, size)
			if err != nil {
				panic(fmt.Sprintf("Can't read [%X] bytes from [%04X]", size, reg.PC))
			}
			instructions, err := zog.DecodeBytes(buf)
			if err != nil {
				panic(fmt.Sprintf("Can't decode: %s", err))
			}
			for _, inst := range instructions {
				fmt.Printf("%s\n", inst)
			}
		default:
			panic(fmt.Sprintf("Unkown mode: %s", *mode))
		}

	} else {

		if flag.NArg() < 1 {
			usage("Missing filename")
		}
		fname := flag.Arg(0)

		buf, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatalf("Failed to open file [%s] : %s\n", fname, err)
		}

		runErr = z.RunBytes(machine.LoadAddr(), buf, machine.RunAddr())
		if err != nil {
			log.Fatalf("RunBytes returned error: %s", err)
		}
	}

	if runErr != nil {
		fmt.Printf("ERR: %s\n", runErr)
	}

	if *haltstate {
		fmt.Printf("STATE: %s\n", z.State())
	}
}

func usage(reason string) {
	fmt.Printf(`%s

%s <filename>

`, reason, os.Args[0])
	os.Exit(1)
}
