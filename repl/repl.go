package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jbert/zog"
)

type Machine struct {
	z         *zog.Zog
	br        *bufio.Reader
	inputLine []byte
}

func NewMachine(z *zog.Zog) *Machine {
	return &Machine{
		z:         z,
		br:        bufio.NewReader(os.Stdin),
		inputLine: make([]byte, 0),
	}
}

func (m Machine) LoadAddr() uint16 {
	return 0x8000
}

func (m Machine) RunAddr() uint16 {
	return 0x8000
}

func (m Machine) Name() string {
	return "repl"
}

func (m *Machine) Start() error {

	m.z.RegisterInputHandler(func(addr uint16) byte {
		//		fmt.Fprintf(os.Stderr, "JB IN addr %04X [%s]\n", addr, m.inputLine)
		if addr != 0 {
			return 0
		}

		for len(m.inputLine) == 0 {
			var err error
			m.inputLine, err = m.br.ReadBytes('\n')
			if err != nil {
				// 0 on EOF (or anything)
				return 0
			}
			//			fmt.Fprintf(os.Stderr, "JB read [%s]\n", m.inputLine)
		}

		n := m.inputLine[0]
		m.inputLine = m.inputLine[1:]
		//		fmt.Fprintf(os.Stderr, "JB return [%02X]\n", n)
		return n
	})

	m.z.RegisterOutputHandler(func(port uint16, n byte) {
		if port != 0 {
			return
		}
		numWritten, err := os.Stdout.Write([]byte{n})
		//		fmt.Fprintf(os.Stderr, "JB wrote [%02X] (%d bytes)\n", n, numWritten)
		if err != nil {
			panic(fmt.Sprintf("Error writing to stdout: %s", err))
		}
		if numWritten != 1 {
			panic("Failed to write a byte to stdout")
		}
	})

	return nil
}

func (m *Machine) Stop() {
}

func (m *Machine) RegisterCallbacks() {}
