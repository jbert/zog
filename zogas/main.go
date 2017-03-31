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

	insts, err := zog.Assemble(string(buf))
	fmt.Printf("Got %d instructions\n", len(insts))
	for _, inst := range insts {
		fmt.Printf("%s\n", inst)
	}
}
