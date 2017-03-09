package zog

import (
	"fmt"
	"sync"
)

var decoderPtrLock sync.Mutex
var decoder *Decoder

func InitialiseDecoder(d *Decoder) {
	decoderPtrLock.Lock()
	defer decoderPtrLock.Unlock()
	decoder = d
}

type InstructionInfo struct {
	encoding byte
	i        Instruction
	name     string
}

type Decoder struct {
	InstructionInfo []InstructionInfo
}

func NewDecoder() *Decoder {
	d := Decoder{}
	d.addInfo([]InstructionInfo{
		{0x37, I_SCF, "SCF"},
		{0x3f, I_CCF, "CCF"},

		{0x76, I_HALT, "HALT"},
	})
	d.loadLD8()
	d.loadAccum()
	return &d
}

func (d *Decoder) loadAccum() {
	infos := make([]InstructionInfo, 8*8)[:0]
	top2 := 2

	for hi3 := 0; hi3 < 8; hi3++ {
		for lo3 := 0; lo3 < 8; lo3++ {
			n := top2<<6 | hi3<<3 | lo3
			i := NewAccumOp(byte(hi3), byte(lo3))

			info := InstructionInfo{
				encoding: byte(n),
				i:        i,
				name:     i.String(),
			}
			infos = append(infos, info)
		}
	}
	d.addInfo(infos)
}

func NewAccumOp(hi3, lo3 byte) Instruction {
	// Arithmetic and logical with accumulator.
	ops := []struct {
		name string
		op   func(z *Zog, a, n byte) error
	}{
		{"ADD", AccumAdd},
		{"ADC", AccumAdc},
		{"SUB", AccumSub},
		{"SBC", AccumSbc},
		{"AND", AccumAnd},
		{"XOR", AccumXor},
		{"OR", AccumOr},
		{"CP", AccumCp},
	}
	src := R8Loc(lo3)

	i := &IAccumOp{src: src, name: ops[hi3].name, op: ops[hi3].op}

	return i
}

func (d *Decoder) loadLD8() {
	infos := make([]InstructionInfo, 8*8)[:0]
	top2 := 1
	for hi3 := 0; hi3 < 8; hi3++ {
		for lo3 := 0; lo3 < 8; lo3++ {
			n := top2<<6 | hi3<<3 | lo3
			// In place of LD (HL), (HL) we have HALT already in the table
			if n == 0x76 {
				continue
			}
			src := R8Loc(lo3)
			dst := R8Loc(hi3)
			if hi3 == 4 {
				// Can read F, but not write to it
				dst = H
			}
			i := &ILD8{src: src, dst: dst}
			info := InstructionInfo{
				encoding: byte(n),
				i:        i,
				name:     i.String(),
			}

			infos = append(infos, info)
		}
	}
	d.addInfo(infos)
}

func (d *Decoder) addInfo(infos []InstructionInfo) {
	// TODO - check encoding not in use
	d.InstructionInfo = append(d.InstructionInfo, infos...)
}

func (d *Decoder) findInfoByEncoding(n byte) (InstructionInfo, bool) {
	for _, info := range d.InstructionInfo {
		if info.encoding == n {
			return info, true
		}
	}
	return InstructionInfo{}, false
}
func (d *Decoder) findInfoByInstruction(i ISimple) (InstructionInfo, bool) {
	for _, info := range d.InstructionInfo {
		if info.i == i {
			return info, true
		}
	}
	return InstructionInfo{}, false
}

func (d *Decoder) Decode(getNext func() (byte, error)) (Instruction, error) {
	var n byte
	var err error
	for {
		n, err = getNext()
		if err != nil {
			return nil, err
		}

		// Table lookup has precedence
		info, ok := d.findInfoByEncoding(n)
		if ok {
			return info.i, nil
		}

		lo3 := n & 0x07
		hi3 := (n & 0x38) >> 3
		top2 := (n & 0xc0) >> 6

		// fmt.Printf("top2 %x, hi3 %x, lo3 %x\n", top2, hi3, lo3)

		switch top2 {

		case 0:
			switch lo3 {
			case 6:
				return decodeLD8Immediate(hi3, getNext)
			default:
				panic(fmt.Sprintf("Failed to decode top0 instruction: 0x%02X", n))
			}

		case 1:
			// Main part of 8bit load group

			panic("JB - should be covered by loadLD8 now")
			/*
				// In place of LD (HL), (HL) we have HALT
				if n == 0x76 {
					return I_HALT, nil
				}
				return decodeLD8(hi3, lo3)
			*/

		case 2:
			panic("JB - should be covered by loadAccum now")

		default:
			panic(fmt.Sprintf("Failed to decode instruction: 0x%02X", n))
		}
	}
	return nil, fmt.Errorf("Failed to decode: %02x", n)
}
