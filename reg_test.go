package zog

import (
	"log"
	"testing"
)

func TestLoadSave8(t *testing.T) {
	regs := []Location8{B, C, D, E, A, F, H, L}
	z := New(1024)
	for _, r := range regs {
		n := z.Read8(r)
		log.Printf("Is %s initially zero", r)
		if n != 0 {
			t.Fail()
		}
		nStore := byte(0x7f)
		log.Printf("Store %d to %s", nStore, r)
		z.Write8(r, nStore)
		n = z.Read8(r)
		log.Printf("Does %s now contain %d", r, nStore)
		if n != nStore {
			t.Fail()
		}
	}
}

func TestLoadSave16(t *testing.T) {
	regs := []Location16{AF, BC, DE, HL, SP, IX, IY}
	z := New(1024)
	for _, r := range regs {
		nn := z.Read16(r)
		log.Printf("Is %s initially zero", r)
		if nn != 0 {
			t.Fail()
		}
		nnStore := uint16(0x1234)
		log.Printf("Store %d to %s", nnStore, r)
		z.Write16(r, nnStore)
		nn = z.Read16(r)
		log.Printf("Does %s now contain %d", r, nnStore)
		if nn != nnStore {
			t.Fail()
		}
	}
}

func TestHiLoLoadSaveAdHoc(t *testing.T) {
	// Ad hoc testing
	z := New(1024)
	log.Printf("Write 0x1234 to BC")
	z.Write16(BC, 0x1234)
	log.Printf("Does B contain 0x12?")
	n := z.Read8(B)
	if n != 0x12 {
		t.Fail()
	}
	log.Printf("Does C contain 0x34?")
	n = z.Read8(C)
	if n != 0x34 {
		t.Fail()
	}
}
