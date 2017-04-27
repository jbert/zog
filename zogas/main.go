package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	br := bufio.NewReader(in)

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

	for lineNum := 1; ; lineNum++ {
		line, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Line %d: error reading: %s", lineNum, err)
		}
		if isPrefix {
			log.Fatalf("Line %d: bufio ReadLine didn't return whole line - very long line?", lineNum)
		}
		//		fmt.Fprintf(os.Stderr, "L: %s\n", line)
		assembly, err := zog.Assemble(string(line))
		if err != nil {
			log.Fatalf("Line %d: failed to assemble: %s", lineNum, err)
		}

		fmt.Fprintf(os.Stderr, assembly.String())
		for _, linst := range assembly.Linsts {
			encodedBuf := linst.Inst.Encode()
			n, err := w.Write(encodedBuf)
			if err != nil {
				log.Fatalf("Line %d: failed to write: %s", lineNum, err)
			}
			if n != len(encodedBuf) {
				log.Fatalf("Line %d: partial write: wrote %d not %d", lineNum, n, len(encodedBuf))
			}
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
