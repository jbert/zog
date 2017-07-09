package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jbert/zog"
)

func main() {
	if len(os.Args) < 2 {
		usage("Missing filename")
	}
	fname := os.Args[1]

	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Failed to open file [%s] : %s\n", fname, err)
	}

	z := zog.New(0)
	loadAddr := uint16(0x8000)
	runAddr := uint16(0x8000)
	err = z.RunBytes(loadAddr, buf, runAddr)
	if err != nil {
		log.Fatalf("RunBytes returned error: %s", err)
	}
}

func usage(reason string) {
	fmt.Printf(`%s

%s <filename>

`, reason, os.Args[0])
	os.Exit(1)
}
