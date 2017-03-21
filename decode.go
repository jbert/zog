package zog

import "fmt"

func Decode(inCh chan byte) (chan instruction, chan error) {
	errCh := make(chan error)
	iCh := make(chan instruction)
	go decode(inCh, iCh, errCh)
	return iCh, errCh
}

func DecodeBytes(buf []byte) ([]instruction, error) {

	ch := make(chan byte)

	go func() {
		for _, n := range buf {
			ch <- n
		}
		close(ch)
	}()

	var insts []instruction
	var err error
	var ok bool

	instCh, errCh := Decode(ch)
	looping := true
	for looping {
		select {
		case inst, ok := <-instCh:
			if !ok {
				looping = false
				break
			}
			insts = append(insts, inst)
		case err, ok = <-errCh:
			if !ok {
				looping = false
				break
			}
			break
		}
	}

	return insts, err
}

func decode(inCh chan byte, iCh chan instruction, errCh chan error) {

	// Set to 0 if no prefix in effect
	var prefix byte

	for n := range inCh {
		// Prefix bytes
		switch n {
		case 0xDD:
		case 0xFD:
			if prefix != 0 {
				errCh <- fmt.Errorf("Double prefix: %02% %02X", prefix, n)
			}
			prefix = n
			continue

		}

		var inst instruction
		x, y, z, _, _ := decomposeByte(n)
		fmt.Printf("D: N %02X, x %d y %d z %d\n", n, x, y, z)
		switch x {
		case 0:
			switch z {
			case 0:
				switch y {
				case 0:
					inst = NOP
				case 1:
					inst = &EX{a: AF, b: AF_PRIME}
				}
			}
		case 1:
			if x == 6 && y == 6 {
				inst = HALT
			} else {
				inst = &LD8{tableR[y], tableR[z]}
			}
		}

		if inst == nil {
			err := fmt.Errorf("TODO - impl %02X [%02X]", n, prefix)
			errCh <- err
		} else {
			iCh <- inst
		}

		prefix = 0
	}
	close(iCh)
	close(errCh)
}

var tableR []Loc8 = []Loc8{B, C, D, E, H, L, Contents{HL}, A}

func decomposeByte(n byte) (byte, byte, byte, byte, byte) {
	// We follow terminology from http://www.z80.info/decoding.htm
	// x = the opcode's 1st octal digit (i.e. bits 7-6)
	// y = the opcode's 2nd octal digit (i.e. bits 5-3)
	// z = the opcode's 3rd octal digit (i.e. bits 2-0)
	// p = y rightshifted one position (i.e. bits 5-4)
	// q = y modulo 2 (i.e. bit 3)
	z := n & 0x07
	y := (n >> 3) & 0x07
	x := (n >> 6) & 0x07
	p := y >> 1
	q := y & 0x01

	return x, y, z, p, q
}
