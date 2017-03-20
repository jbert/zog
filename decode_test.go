package zog

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestDecode(t *testing.T) {
	data := []byte{0x41}
	b := bytes.NewBuffer(data)
	ch := make(chan byte)

	go func() {
		for {
			n, err := b.ReadByte()
			if err != nil {
				if err == io.EOF {
					fmt.Printf("JB - EOF, closing channel\n")
					close(ch)
					break
				}
				panic(fmt.Sprintf("Error reading from buffer: %s", err))
			}
			ch <- n
		}
	}()

	iCh, errCh := Decode(ch)
	looping := true
	for looping {
		select {
		case i, ok := <-iCh:
			if !ok {
				fmt.Printf("I: END\n", i)
				looping = false
				break
			}
			fmt.Printf("I: %s\n", i)
		case err, ok := <-errCh:
			if !ok {
				fmt.Printf("E: END\n")
				looping = false
				break
			}
			fmt.Printf("E: %s\n", err)
			break
		}
	}

}
