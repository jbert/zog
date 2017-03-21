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

func getImmN(inCh chan byte) (Imm8, error) {
	n, ok := <-inCh
	if !ok {
		return 0, fmt.Errorf("getImmN: Can't get byte")
	}
	return Imm8(n), nil
}

func getImmNN(inCh chan byte) (Imm16, error) {
	l, ok := <-inCh
	if !ok {
		return 0, fmt.Errorf("getImmNN: Can't get lo byte")
	}
	h, ok := <-inCh
	if !ok {
		return 0, fmt.Errorf("getImmNN: Can't get hi byte")
	}
	return Imm16(uint16(h)<<8 | uint16(l)), nil
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
		var instErr error
		x, y, z, p, q := decomposeByte(n)
		fmt.Printf("D: N %02X, x %d y %d z %d p %d q %d\n", n, x, y, z, p, q)
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
			case 1:
				if q == 0 {
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = &LD16{tableRP[p], nn}
					} else {
						instErr = err
					}
				} else {
					inst = &ADD16{HL, tableRP[p]}
				}
			case 2:
				if q == 0 {
					switch p {
					case 0:
						inst = &LD8{Contents{BC}, A}
					case 1:
						inst = &LD8{Contents{DE}, A}
					case 2:
						nn, err := getImmNN(inCh)
						if err == nil {
							inst = &LD16{Contents{nn}, HL}
						} else {
							instErr = err
						}
					case 3:
						nn, err := getImmNN(inCh)
						if err == nil {
							inst = &LD8{Contents{nn}, A}
						} else {
							instErr = err
						}
					}
				} else {
					switch p {
					case 0:
						inst = &LD8{A, Contents{BC}}
					case 1:
						inst = &LD8{A, Contents{DE}}
					case 2:
						nn, err := getImmNN(inCh)
						if err == nil {
							inst = &LD16{HL, Contents{nn}}
						} else {
							instErr = err
						}
					case 3:
						nn, err := getImmNN(inCh)
						if err == nil {
							inst = &LD8{A, Contents{nn}}
						} else {
							instErr = err
						}
					}
				}
			case 3:
				if q == 0 {
					inst = &INC16{tableRP[p]}
				} else {
					inst = &DEC16{tableRP[p]}
				}
			case 4:
				inst = &INC8{tableR[y]}
			case 5:
				inst = &DEC8{tableR[y]}
			case 6:
				n, err := getImmN(inCh)
				if err == nil {
					inst = &LD8{tableR[y], n}
				} else {
					instErr = err
				}
			}
		case 1:
			if x == 6 && y == 6 {
				inst = HALT
			} else {
				inst = &LD8{tableR[y], tableR[z]}
			}
		}
		fmt.Printf("D: inst [%v] err [%v]\n", inst, instErr)

		if inst == nil {
			err := instErr
			if err == nil {
				err = fmt.Errorf("TODO - impl %02X [%02X]", n, prefix)
			}
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
var tableRP []Loc16 = []Loc16{BC, DE, HL, SP}
var tableRP2 []Loc16 = []Loc16{BC, DE, HL, AF}

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
