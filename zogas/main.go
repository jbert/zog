package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jbert/zog"
)

func main() {
	fname := "tt.z80"
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Can't open [%s]: %s", fname, err)
	}
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Can't readall: %s", err)
	}
	f.Close()

	assembler := &zog.Assembler{Buffer: string(buf)}
	assembler.Init()
	err = assembler.Parse()
	if err != nil {
		log.Fatalf("Can't parse: %s", err)
	}
	assembler.Execute()

	insts := assembler.GetInstructions()

	fmt.Printf("Got %d instructions\n", len(insts))
	for _, inst := range insts {
		fmt.Printf("%s\n", inst)
	}

	//	assembler.Print()
	//	assembler.PrintSyntaxTree()
}
