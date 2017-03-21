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

func decode(inCh chan byte, iCh chan instruction, errCh chan error) {

	// Set to 0 if no prefix in effect
	var opPrefix byte
	var indexPrefix byte

	t := NewTable()

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

		var inst instruction
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

func cbDecode(t *Table, inCh chan byte, indexPrefix, n byte) (instruction, error) {
	var err error
	var inst instruction

	x, y, z, p, q := decomposeByte(n)
	fmt.Printf("D: N %02X, x %d y %d z %d p %d q %d\n", n, x, y, z, p, q)

	switch x {
	case 0:
		info := tableROT[y]
		inst = &ROT{name: info.name /* f: info.f, */, r: t.LookupR(z)}
	case 1:
		inst = &BIT{y, t.LookupR(z)}
	case 2:
		inst = &RES{y, t.LookupR(z)}
	case 3:
		inst = &SET{y, t.LookupR(z)}
	}

	return inst, err
}

func edDecode(t *Table, inCh chan byte, indexPrefix, n byte) (instruction, error) {
	panic("TODO - impl ed")
}

func baseDecode(t *Table, inCh chan byte, indexPrefix, n byte) (instruction, error) {
	var err error
	var inst instruction

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
				inst = &EX{AF, AF_PRIME}
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
					inst = &LD16{t.LookupRP(p), nn}
				}
			} else {
				inst = &ADD16{HL, t.LookupRP(p)}
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
					}
				case 3:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = &LD8{Contents{nn}, A}
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
					}
				case 3:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = &LD8{A, Contents{nn}}
					}
				}
			}
		case 3:
			if q == 0 {
				inst = &INC16{t.LookupRP(p)}
			} else {
				inst = &DEC16{t.LookupRP(p)}
			}
		case 4:
			inst = &INC8{t.LookupR(y)}
		case 5:
			inst = &DEC8{t.LookupR(y)}
		case 6:
			n, err := getImmN(inCh)
			if err == nil {
				inst = &LD8{t.LookupR(y), n}
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
			inst = &LD8{t.LookupR(y), t.LookupR(z)}
		}
	case 2:
		info := tableALU[y]
		inst = &Accum{name: info.name /* f: info.f, */, src: t.LookupR(z)}
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
					inst = &JP{True, Contents{HL}}
				case 3:
					inst = &LD16{SP, HL}
				}
			}
		case 2:
			nn, err := getImmNN(inCh)
			if err == nil {
				inst = &JP{tableCC[y], nn}
			}
		case 3:
			switch y {
			case 0:
				nn, err := getImmNN(inCh)
				if err == nil {
					inst = &JP{True, nn}
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
				inst = &EX{Contents{SP}, HL}
			case 5:
				inst = &EX{DE, HL}
			case 6:
				inst = DI
			case 7:
				inst = EI
			}
		case 4:
			nn, err := getImmNN(inCh)
			if err == nil {
				inst = &CALL{tableCC[y], nn}
			}
		case 5:
			if q == 0 {
				inst = &PUSH{t.LookupRP2(p)}
			} else {
				switch p {
				case 0:
					nn, err := getImmNN(inCh)
					if err == nil {
						inst = &CALL{True, nn}
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
				inst = &Accum{name: info.name /* f: info.f, */, src: n}
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
