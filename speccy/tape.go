package speccy

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jbert/zog"
)

func NewTape(f *os.File) *Tape {
	return &Tape{f, false, make(chan bool), false}
}

type Tape struct {
	handle    *os.File
	reading   bool
	bit_cache chan bool

	high bool
}

var lastTape bool

func (t *Tape) tapeEarByte() byte {

	if t.high != lastTape {
	}
	lastTape = t.high
	if t.high {
		return 1 << 6
	} else {
		return 0 << 6
	}
}

func (t *Tape) flushData() {
	t.reading = false
}

func (t *Tape) SwitchTapeFile(f *os.File) {
	t.flushData()
	t.handle = f
}

func (t *Tape) StopTapeRead() {
	t.reading = false
}

func (t *Tape) BeginTapeRead() {
	t.reading = true
}

func (t *Tape) Pulse(cycles uint64) {

	duration := zog.TStateDuration * time.Duration(cycles)

	waitUntil := time.Now().Add(duration)
	for time.Now().Before(waitUntil) {
	}

	t.high = !t.high
}

func (t *Tape) ReadBlockHeader() uint16 {
	var hdrData []byte = make([]byte, 3)
	_, err := t.handle.Read(hdrData)
	if err != nil {
		panic(err)
	}

	_, err = t.handle.Seek(-1, 1)
	if err != nil {
		panic(err)
	}

	block_type := hdrData[2]

	hdrData = hdrData[:len(hdrData)-1]

	size := binary.LittleEndian.Uint16(hdrData)

	fmt.Printf("Beginning Leader\n")
	if block_type == 0x00 {
		for i := 0; i < 8063; i++ {
			t.Pulse(2168)
		}
	} else if block_type == 0xFF {
		for i := 0; i < 3223; i++ {
			t.Pulse(2168)
		}
	} else {
		panic(fmt.Errorf("Invalid block type %02X\n", block_type))
	}
	fmt.Printf("Finished Leader\n")
	t.Pulse(667)
	t.Pulse(735)
	fmt.Printf("Finished Sync pulses\n")

	return size
}

func TapeReadManager(t *Tape) {
	for {
		if !t.reading {
			continue
		}

		test := make([]byte, 3)
		_, testerr := t.handle.Read(test)
		if testerr == io.EOF {
			fmt.Println("End of Tape")
			t.reading = false
			continue
		} else if testerr != nil {
			panic(testerr)
		}
		t.handle.Seek(-3, 1)
		block_size := t.ReadBlockHeader()

		for i := uint16(0); i < block_size; i++ {
			byte_cache := make([]byte, 1)
			_, err := t.handle.Read(byte_cache)
			if err != nil {
				panic(err)
			}

			for b := 7; b >= 0; b-- {
				if byte_cache[0]&(1<<b) != 0 {
					t.Pulse(1710)
					t.Pulse(1710)
				} else {
					t.Pulse(855)
					t.Pulse(855)
				}
			}
		}

		fmt.Printf("Ending block data\n")
	}
}
