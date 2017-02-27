package zog

import (
	"log"
	"testing"
)

func TestLoadSave(t *testing.T) {
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
