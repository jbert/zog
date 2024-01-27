package file

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/jbert/zog"
)

type SNASnapshot struct {
	I                                  uint8
	HLPrime, DEPrime, BCPrime, AFPrime uint16
	HL, DE, BC                         uint16
	IY, IX                             uint16
	IFF2                               uint8
	R                                  uint8
	AF, SP                             uint16
	IMODE, BCOL                        uint8
}

func (s *SNASnapshot) Parse(r io.Reader) error {

	binary.Read(r, binary.LittleEndian, s)

	return nil
}

func getBytes(r io.Reader, n uint16) ([]byte, error) {
	var b []byte = make([]byte, n)
	_, err := r.Read(b)
	return b, err
}

func (s *SNASnapshot) Load(r io.Reader, z *zog.Zog) error {
	fmt.Printf("s: %v\n", s)
	var regs zog.Registers

	regs.Write16(zog.AF, s.AF)
	regs.Write16(zog.BC, s.BC)
	regs.Write16(zog.DE, s.DE)
	regs.Write16(zog.HL, s.HL)

	regs.Write16(zog.AF_PRIME, s.AFPrime)
	regs.Write16(zog.BC_PRIME, s.BCPrime)
	regs.Write16(zog.DE_PRIME, s.DEPrime)
	regs.Write16(zog.HL_PRIME, s.HLPrime)

	regs.Write16(zog.SP, s.SP)

	regs.Write16(zog.IY, s.IY)
	regs.Write16(zog.IX, s.IX)

	regs.Write8(zog.I, s.I)
	regs.Write8(zog.R, s.R)

	z.LoadRegisters(regs)

	var state zog.InterruptState
	state.Mode = s.IMODE
	state.IFF2 = (s.IFF2 == 2)
	state.IFF1 = state.IFF2

	z.LoadInterruptState(state)

	fmt.Printf("z.State(): %v\n", z.State())

	var mem []byte = make([]byte, 0xC000)

	_, err := r.Read(mem)
	if err != nil {
		return err
	}

	err = z.LoadBytes(0x4000, mem)
	if err != nil {
		return err
	}

	pc, err := z.Mem.Peek16(z.GetRegisters().SP)
	if err != nil {
		return err
	}
	regs = z.GetRegisters()
	regs.SP += 2
	regs.PC = pc
	z.LoadRegisters(regs)

	return nil
}
