package file

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jbert/zog"
)

type Z80Snapshot struct {
	Z80header
}

func (z *Z80Snapshot) Parse(r io.Reader) error {
	err := Z80readHeader(r, &z.Z80header)
	return err
}

func (snap *Z80Snapshot) Load(r io.Reader, z *zog.Zog) error {
	return snap.Z80header.Load(r, z)
}

// From http://www.worldofspectrum.org/faq/reference/z80format.htm
type Z80header struct {
	A, F                 byte
	BC, HL, PC, SP       uint16
	I, R                 byte
	Flag1                byte
	DE, BC_P, DE_P, HL_P uint16
	A_P, F_P             byte
	IY, IX               uint16
	IFF1, IFF2           byte
	Flag2                byte
}

func Z80readHeader(r io.Reader, header *Z80header) error {
	err := binary.Read(r, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	// "Because of compatibility, if byte 12 is 255, it has to be regarded as being 1."
	if header.Flag1 == 0xff {
		header.Flag1 = 0x01
	}
	return nil
}

func (h *Z80header) Load(r io.Reader, z *zog.Zog) error {
	mem, err := h.Z80readMem(r)
	if err != nil {
		return fmt.Errorf("Can't read mem: %s", err)
	}

	reg := zog.Registers{}

	// We have header state and memory, ship it in
	reg.Write8(zog.A, h.A)
	reg.Write8(zog.F, h.F)
	reg.Write16(zog.BC, h.BC)
	reg.Write16(zog.DE, h.DE)
	reg.Write16(zog.HL, h.HL)
	reg.Write16(zog.IX, h.IX)
	reg.Write16(zog.IY, h.IY)

	reg.Write8(zog.I, h.I)
	reg.Write8(zog.R, h.R)

	afp := uint16(h.A_P)<<8 | uint16(h.F_P)
	reg.Write16(zog.AF_PRIME, afp)
	reg.Write16(zog.BC_PRIME, h.BC_P)
	reg.Write16(zog.DE_PRIME, h.DE_P)
	reg.Write16(zog.HL_PRIME, h.HL_P)

	reg.Write16(zog.SP, h.SP)
	reg.Write16(zog.PC, h.PC)

	z.LoadRegisters(reg)

	is := zog.InterruptState{
		IFF1: h.IFF1 != 0,
		IFF2: h.IFF2 != 0,
		//		Mode: h.Flag2 & 0x3,
		Mode: 1,
	}
	z.LoadInterruptState(is)

	//	if h.IFF1 == 0 {
	//		z.Di()
	//	} else {
	//		z.Ei()
	//	}
	err = z.LoadBytes(0x4000, mem)

	return nil
}

func (h *Z80header) IsVersion1() bool {
	return h.PC != 0
}

func (h *Z80header) IsCompresed() bool {
	return h.Flag1&0x20 != 0
}

func (h *Z80header) Z80readMem(r io.Reader) ([]byte, error) {
	if !h.IsVersion1() {
		return nil, errors.New("TODO - implement support for version 2+ Z80 files")
	}
	buf := make([]byte, 0x10000)
	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	buf = buf[0:n]
	if h.IsCompresed() {
		buf, err = DecompressMem(buf)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}

func DecompressMem(in []byte) ([]byte, error) {
	if len(in) < 4 {
		return nil, fmt.Errorf("Missing end-of-block - len %d", len(in))
	}
	last := in[len(in)-4:]
	if !bytes.Equal(last, []byte{0x00, 0xed, 0xed, 0x00}) {
		return nil, fmt.Errorf("Missing end-of-block: %v", last)
	}
	in = in[:len(in)-4]

	out := []byte{}
	last = nil
	for _, b := range in {
		if len(last) == 2 {
			last = append(last, b)
			continue
		}

		if len(last) == 3 {
			last = append(last, b)

			if !(last[0] == 0xed && last[1] == 0xed && len(last) == 4) {
				return nil, fmt.Errorf("Internal error: last [%v]", last)
			}
			repeatBuf := bytes.Repeat(last[3:], int(last[2]))
			out = append(out, repeatBuf...)
			last = nil
			continue
		}

		if b == 0xed {
			last = append(last, b)
			continue
		}
		if len(last) > 0 {
			out = append(out, last...)
			last = nil
		}
		out = append(out, b)
	}
	out = append(out, last...) // Drain any partial
	return out, nil
}
