package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/jbert/zog"
)

type options struct {
	inFileName  string
	outFileName string
}

func main() {
	o, err := parseArgs()
	if err != nil {
		log.Fatalf("Bad arguments: %s\n", err)
		os.Exit(1)
	}

	var in io.Reader
	if o.inFileName == "-" {
		in = os.Stdin
	} else {
		f, err := os.Open(o.inFileName)
		if err != nil {
			log.Fatalf("Can't open [%s]: %s", o.inFileName, err)
		}
		defer f.Close()
		in = f
	}

	var w io.Writer
	if o.outFileName == "-" {
		w = os.Stdout
	} else {
		f, err := os.Create(o.outFileName)
		if err != nil {
			log.Fatalf("Can't open [%s]: %s", o.outFileName, err)
		}
		defer f.Close()
		w = f
	}

	contents, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	assembly, err := zog.Assemble(string(contents))
	if err != nil {
		log.Fatalf("failed to assemble: %s", err)
	}

	fmt.Fprintf(os.Stderr, assembly.String())
	for _, linst := range assembly.Linsts {
		encodedBuf := linst.Inst.Encode()
		n, err := w.Write(encodedBuf)
		if err != nil {
			log.Fatalf("failed to write: %s", err)
		}
		if n != len(encodedBuf) {
			log.Fatalf("partial write: wrote %d not %d", n, len(encodedBuf))
		}
	}

}

func parseArgs() (*options, error) {
	o := options{}
	flag.StringVar(&o.outFileName, "out", "-", "Output file (default stdout)")
	flag.Parse()
	if len(os.Args) == 2 {
		o.inFileName = os.Args[1]
	} else {
		o.inFileName = "-"
	}
	return &o, nil
}
