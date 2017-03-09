package zog

import (
	"fmt"
	"log"
	"testing"
)

func TestLoadSave8(t *testing.T) {
	regs := []Location8{B, C, D, E, A, F, L}
	z := New(1024)
	for _, r := range regs {
		n := z.Read8(r)
		if n != 0 {
			t.Errorf("%s is not initially zero", r)
		}
		nStore := byte(0x7f)
		log.Printf("Store %d to %s", nStore, r)
		z.Write8(r, nStore)
		n = z.Read8(r)
		if n != nStore {
			t.Errorf("%s does not now contain", r, nStore)
		}
	}
}

func TestLoadSave16(t *testing.T) {
	regs := []Location16{AF, BC, DE, HL, SP, IX, IY}
	z := New(1024)
	for _, r := range regs {
		nn := z.Read16(r)
		if nn != 0 {
			t.Errorf("%s is not initially zero", r)
		}
		nnStore := uint16(0x1234)
		log.Printf("Store %d to %s", nnStore, r)
		z.Write16(r, nnStore)
		nn = z.Read16(r)
		if nn != nnStore {
			t.Fail()
			t.Errorf("%s does not now contain %d", r, nnStore)
		}
	}
}

func TestHiLoLoadSaveAdHoc(t *testing.T) {
	// Ad hoc testing
	z := New(1024)
	log.Printf("Write 0x1234 to BC")
	z.Write16(BC, 0x1234)
	n := z.Read8(B)
	if n != 0x12 {
		t.Errorf("B doesn't contain 0x12 after BC load")
	}
	n = z.Read8(C)
	if n != 0x34 {
		t.Errorf("C doesn't contain 0x34 after BC load")
	}
}

func TestDecodeLDBasic(t *testing.T) {
	testCases := []struct {
		codes []byte
		str   string
	}{
		{[]byte{0x7f}, "LD A, A"},
		{[]byte{0x41}, "LD B, C"},
		{[]byte{0x4c}, "LD C, F"},
		{[]byte{0x67}, "LD H, A"},
		{[]byte{0x64}, "LD H, F"},

		{[]byte{0x7e}, "LD A, (HL)"},
		{[]byte{0x77}, "LD (HL), A"},

		{[]byte{0x3e, 0xab}, "LD A, 0xAB"},
	}
	for _, tc := range testCases {
		log.Printf("About to decode: %X", tc.codes)
		codes := tc.codes
		getNext := func() (byte, error) {
			if len(codes) == 0 {
				t.Fail()
				return 0, fmt.Errorf("ERROR: overrun in decode")
			}
			c := codes[0]
			codes = codes[1:]
			return c, nil
		}
		i, err := decoder.Decode(getNext)
		if err != nil {
			t.Errorf("Failed: %s", err)
		}
		if i.String() != tc.str {
			t.Errorf("Wrong instruction: %s != %s", i, tc.str)
		}
		log.Printf("Decoded %s", i)
	}
}
