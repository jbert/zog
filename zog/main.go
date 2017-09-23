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
	"github.com/jbert/zog/speccy"
)

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile `file`")
	trace := flag.String("trace", "", "Trace addresses: start-end,s2-e2")
	watch := flag.String("watch", "", "Watch addresses: start-end,s2-e2")
	haltstate := flag.Bool("haltstate", false, "Print state on halt")
	numhalttrace := flag.Int("halttrace", 0, "Number of traces to print on halt")
	machineName := flag.String("machine", "none", "Machine for console printer (none, cpm, spectrum)")

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

	var machine zog.Machine

	switch *machineName {
	case "cpm":
		machine = cpm.NewMachine(z)
	case "spectrum", "speccy":
		machine = speccy.NewMachine(z)
	default:
		panic("Specify a machine type")
	}

	fmt.Printf("Loading %s\n", machine.Name())
	err = machine.Start()
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

	loadFname := ""
	//loadFname := "tt.z80"

	if loadFname != "" {
		h := file.Z80header{}
		f, err := os.Open(loadFname)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		err = file.Z80readHeader(f, &h)
		if err != nil {
			panic(err)
		}
		err = h.Load(f, z)
		if err != nil {
			panic(err)
		}

		err = z.Run()
	} else {

		err = z.RunBytes(machine.LoadAddr(), buf, machine.RunAddr())
		if err != nil {
			log.Fatalf("RunBytes returned error: %s", err)
		}
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
