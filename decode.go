package zog

import (
	"fmt"
	"log"
)

func Decode(inCh chan byte) (chan Instruction, chan error) {
	errCh := make(chan error)
	iCh := make(chan Instruction)
	go decode(inCh, iCh, errCh)
	return iCh, errCh
}

func DecodeBytes(buf []byte) ([]Instruction, error) {

	ch := make(chan byte)

	go func() {
		for _, n := range buf {
			ch <- n
		}
		close(ch)
	}()

	var insts []Instruction
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

func getImmd(inCh chan byte) (Disp, error) {
	d, ok := <-inCh
	if !ok {
		return 0, fmt.Errorf("getImmd: Can't get byte")
	}
	return Disp(d), nil
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

func decode(inCh chan byte, iCh chan Instruction, errCh chan error) {

	// Set to 0 if no prefix in effect
	var opPrefix byte
	var indexPrefix byte

	t := NewDecodeTable(inCh)

	for n := range inCh {

		if opPrefix == 0 {
			switch n {
			case 0xcb, 0xed:
				opPrefix = n
				continue
			case 0xdd, 0xfd:
				// Last one wins
				indexPrefix = n
				continue
			}
		}

		t.ResetPrefix(indexPrefix)

		var inst Instruction
		var err error

		switch opPrefix {
		case 0:
			inst, err = baseDecode(t, inCh, indexPrefix, n)
		case 0xcb:
			inst, err = cbDecode(t, inCh, indexPrefix, n)
		case 0xed:
			inst, err = edDecode(t, inCh, indexPrefix, n)
		}

		fmt.Printf("D: inst [%v] err [%v]\n", inst, err)

		if inst == nil {
			if err == nil {
				err = fmt.Errorf("TODO - impl %02X [%02X] (%02X)", n, opPrefix, indexPrefix)
			}
			errCh <- err
		} else {
			iCh <- inst
		}

		opPrefix = 0
		indexPrefix = 0
	}
	close(iCh)
	close(errCh)
}

func cbDecode(t *DecodeTable, inCh chan byte, indexPrefix, n byte) (Instruction, error) {
	var err error
	var inst Instruction

	x, y, z, p, q := decomposeByte(n)
	fmt.Printf("D: N %02X, x %d y %d z %d p %d q %d\n", n, x, y, z, p, q)

	switch x {
	case 0:
		info := tableROT[y]
		inst = NewRot(info.name, t.LookupR(z))
	case 1:
		inst = NewBIT(y, t.LookupR(z))
	case 2:
		inst = NewRES(y, t.LookupR(z))
	case 3:
		inst = NewSET(y, t.LookupR(z))
	}

	return inst, err
}

func edDecode(t *DecodeTable, inCh chan byte, indexPrefix, n byte) (Instruction, error) {
	var err error
	var inst Instruction

	hl := HL
	if indexPrefix == 0xDD {
		hl = IX
	} else if indexPrefix == 0xFD {
		hl = IY
	}

	x, y, z, p, q := decomposeByte(n)
	fmt.Printf("D: N %02X, x %d y %d z %d p %d q %d\n", n, x, y, z, p, q)

	switch x {
	case 0, 3:
		// Invalid instruction, equivalent to NONI followed by NOP
		log.Printf("Invalid instruction: [%02X]\n", n)
		inst = NOP
	case 1:
		switch z {
		case 0:
			if y == 6 {
				// This would be (HL), it's unclear what this should be
				log.Printf("IN (C)? or 'set flag'?")
				inst = &IN{dst: F, port: C}
			} else {
				inst = &IN{dst: t.LookupR(y), port: C}
			}
		case 1:
			if y == 6 {
				inst = &OUT{port: C, value: Imm8(0)}
			} else {
				inst = &OUT{port: C, value: t.LookupR(y)}
			}
		case 2:
			if q == 0 {
				inst = &SBC16{hl, t.LookupRP(p)}
			} else {
				inst = &ADC16{hl, t.LookupRP(p)}
			}
		case 3:
			nn, err := getImmNN(inCh)
			if err == nil {
				if q == 0 {
					inst = NewLD16(Contents{nn}, t.LookupRP(p))
				} else {
					inst = NewLD16(t.LookupRP(p), Contents{nn})
				}
			}
		case 4:
			inst = NEG
		case 5:
			if y == 1 {
				inst = RETI
			} else {
				inst = RETN
			}
		case 6:
			switch y {
			case 0, 1, 4, 5:
				inst = IM0
			case 2, 6:
				inst = IM1
			case 3, 7:
				inst = IM2
			}
		case 7:
			switch y {
			case 0:
				inst = NewLD8(I, A)
			case 1:
				inst = NewLD8(R, A)
			case 2:
				inst = NewLD8(A, I)
			case 3:
				inst = NewLD8(A, R)
			case 4:
				inst = RRD
			case 5:
				inst = RLD
			case 6:
				inst = NOP
			case 7:
				inst = NOP
			}
		}
	case 2:
		if z <= 3 && y >= 4 {
			inst = t.LookupBLI(y-4, z)
		} else {
			log.Printf("Invalid instruction: [%02X]\n", n)
			inst = NOP
		}
	}

	//	err = errors.New("TODO - impl")
	return inst, err
}

func baseDecode(t *DecodeTable, inCh chan byte, indexPrefix, n byte) (Instruction, error) {
	var err error
	var inst Instruction

	// We lookup this to get (HL)
	//	hlci := byte(6)
	hl := HL
	if indexPrefix == 0xDD {
		hl = IX
	} else if indexPrefix == 0xFD {
		hl = IY
	}

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
				inst = NewEX(AF, AF_PRIME)
			case 2:
				d, err := getImmd(inCh)
				if err == nil {
					inst = &DJNZ{d}
				}
			case 3:
				d, err := getImmd(inCh)
				if err == nil {
					inst = &JR{True, d}
				}
			case 4, 5, 6, 7:
				d, err := getImmd(inCh)
				if err == nil {
					inst = &JR{tableCC[y-4], d}
				}
			}
		case 1:
			if q == 0 {
				nn, err := getImmNN(inCh)
				if err == nil {
					inst = NewLD16(t.LookupRP(p), nn)
				}
			} else {
				inst = NewADD16(hl, t.LookupRP(p))
			}
		case 2:
			if q == 0 {
				switch p {
				case 0:
					inst = NewLD8(Contents{BC}, A)
				case 1:
					inst = NewLD8(Contents{DE}, A)
				case 2:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = NewLD16(Contents{nn}, hl)
					}
				case 3:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = NewLD8(Contents{nn}, A)
					}
				}
			} else {
				switch p {
				case 0:
					inst = NewLD8(A, Contents{BC})
				case 1:
					inst = NewLD8(A, Contents{DE})
				case 2:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = NewLD16(hl, Contents{nn})
					}
				case 3:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = NewLD8(A, Contents{nn})
					}
				}
			}
		case 3:
			if q == 0 {
				inst = NewINC16(t.LookupRP(p))
			} else {
				inst = NewDEC16(t.LookupRP(p))
			}
		case 4:
			inst = NewINC8(t.LookupR(y))
		case 5:
			inst = NewDEC8(t.LookupR(y))
		case 6:
			// Lookup before immmediate, so we handle IX/IY index before immediate N
			r := t.LookupR(y)
			n, err := getImmN(inCh)
			if err == nil {
				inst = NewLD8(r, n)
			}
		case 7:
			switch y {
			case 0:
				inst = RLCA
			case 1:
				inst = RRCA
			case 2:
				inst = RLA
			case 3:
				inst = RRA
			case 4:
				inst = DAA
			case 5:
				inst = CPL
			case 6:
				inst = SCF
			case 7:
				inst = CCF
			}
		}
	case 1:
		if z == 6 && y == 6 {
			inst = HALT
		} else {
			// Annoying prefix case, if we have (IX+d), we *don't* index-replace
			// H or L
			dst := t.LookupR(y)
			src := t.LookupR(z)
			t.ResetPrefix(0x00)
			if _, ok := dst.(IndexedContents); ok {
				src = t.LookupR(z)
			}
			if _, ok := src.(IndexedContents); ok {
				dst = t.LookupR(y)
			}
			inst = NewLD8(dst, src)
		}
	case 2:
		info := tableALU[y]
		inst = NewAccum(info.name, t.LookupR(z))
	case 3:
		switch z {
		case 0:
			inst = &RET{tableCC[y]}
		case 1:
			if q == 0 {
				inst = &POP{t.LookupRP2(p)}
			} else {
				switch p {
				case 0:
					inst = &RET{True}
				case 1:
					inst = EXX
				case 2:
					inst = NewJP(True, hl)
				case 3:
					inst = NewLD16(SP, hl)
				}
			}
		case 2:
			nn, err := getImmNN(inCh)
			if err == nil {
				inst = NewJP(tableCC[y], nn)
			}
		case 3:
			switch y {
			case 0:
				nn, err := getImmNN(inCh)
				if err == nil {
					inst = NewJP(True, nn)
				}
			case 1:
				panic(fmt.Sprintf("Decoding CB [%02X] as instruction, not prefix", n))
			case 2:
				n, err := getImmN(inCh)
				if err == nil {
					inst = &OUT{n, A}
				}
			case 3:
				n, err := getImmN(inCh)
				if err == nil {
					inst = &IN{A, n}
				}
			case 4:
				inst = NewEX(Contents{SP}, hl)
			case 5:
				// We use real HL for this, it is an exception
				inst = NewEX(DE, HL)
			case 6:
				inst = DI
			case 7:
				inst = EI
			}
		case 4:
			nn, err := getImmNN(inCh)
			if err == nil {
				inst = NewCALL(tableCC[y], nn)
			}
		case 5:
			if q == 0 {
				inst = NewPUSH(t.LookupRP2(p))
			} else {
				switch p {
				case 0:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = NewCALL(True, nn)
					}
				case 1:
					panic(fmt.Sprintf("Decoding DD [%02X] as instruction, not prefix", n))
				case 2:
					panic(fmt.Sprintf("Decoding ED [%02X] as instruction, not prefix", n))
				case 3:
					panic(fmt.Sprintf("Decoding FD [%02X] as instruction, not prefix", n))
				}
			}
		case 6:
			n, err := getImmN(inCh)
			if err == nil {
				info := tableALU[y]
				inst = NewAccum(info.name, n)
			}
		case 7:
			inst = &RST{y * 8}
		}
	}

	return inst, err
}

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
