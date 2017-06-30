package zog

import (
	"fmt"
	"testing"
)

type assert interface {
	check(z *Zog) error
}

type memA struct {
	addr     uint16
	expected byte
}

func (ma memA) check(z *Zog) error {
	actual, err := z.mem.Peek(ma.addr)
	if err != nil {
		return fmt.Errorf("assert failed: failed to peek addr [%04X]: %s", ma.addr, err)
	}
	if actual != ma.expected {
		return fmt.Errorf("assert failed: addr [%04X] actual %02X expected %02X", ma.addr, actual, ma.expected)
	}
	return nil
}

type loc16A struct {
	loc      Loc16
	expected uint16
}

func (la loc16A) check(z *Zog) error {
	actual, err := la.loc.Read16(z)
	if err != nil {
		return fmt.Errorf("assert failed: failed to read location [%s]: %s", la.loc, err)
	}
	if actual != la.expected {
		return fmt.Errorf("assert failed: loc [%s] actual %02X expected %02X", la.loc, actual, la.expected)
	}
	return nil
}

type locA struct {
	loc      Loc8
	expected byte
}

func (la locA) check(z *Zog) error {
	actual, err := la.loc.Read8(z)
	if err != nil {
		return fmt.Errorf("assert failed: failed to read location [%s]: %s", la.loc, err)
	}
	if actual != la.expected {
		return fmt.Errorf("assert failed: loc [%s] actual %02X expected %02X", la.loc, actual, la.expected)
	}
	return nil
}

type flagA struct {
	f        flag
	expected bool
}

func (fa flagA) check(z *Zog) error {
	actual := z.GetFlag(fa.f)
	if actual != fa.expected {
		return fmt.Errorf("assert failed: flag [%s] actual %v expected %v", fa.f, actual, fa.expected)
	}
	return nil
}

type executeTestCase struct {
	prog       string
	assertions []assert
}

func TestExecuteBasic(t *testing.T) {
	addr := uint16(0x100)
	testCases := []executeTestCase{
		//		{"LD A,10h : LD B,05h : ADD A,B", []assert{
		//			locA{A, 0x15},
		//		}},
		//{"LD HL,1111h : LD DE, 2222h : ADD HL, DE", []assert{
		//	loc16A{HL, 0x3333},
		//}},
		{"LD HL,1234h : LD (0100h), HL", []assert{
			memA{0x0100, 0x34},
			memA{0x0101, 0x12},
		}},

		{"LD A,10h : DEC A : LD A,00h", []assert{
			locA{A, 0x00},
			flagA{F_S, false},
			flagA{F_Z, false},
			flagA{F_H, false},
			flagA{F_PV, false},
			flagA{F_C, false},
		}},
		{"LD A,01h : DEC A : LD A,00h", []assert{
			locA{A, 0x00},
			flagA{F_S, false},
			flagA{F_Z, true},
			flagA{F_H, false},
			flagA{F_PV, false},
			flagA{F_C, false},
		}},
		{"LD A,00h : DEC A : LD A,00h", []assert{
			locA{A, 0x00},
			flagA{F_S, true},
			flagA{F_Z, false},
			flagA{F_H, false},
			flagA{F_PV, true},
			flagA{F_C, true},
		}},

		{"LD A,10h : DEC A", []assert{locA{A, 0x0f}, flagA{F_Z, false}}},
		{"LD A,00h : DEC A", []assert{locA{A, 0xff}, flagA{F_Z, false}}},
		{"LD A,01h : DEC A", []assert{locA{A, 0x00}, flagA{F_Z, true}}},

		{"LD A,10h : INC A", []assert{locA{A, 0x11}, flagA{F_Z, false}}},
		{"LD A,ffh : INC A", []assert{locA{A, 0x00}, flagA{F_Z, true}}},
		{"LD A,00h : INC A", []assert{locA{A, 0x01}, flagA{F_Z, false}}},

		{"LD A,10h : LD IX, 0x0100 : LD (IX+3), A", []assert{locA{Contents{Imm16(addr + 3)}, 0x10}}},
		{"LD A,10h : LD IX, 0x0100 : LD (IX+3), A : LD B, (IX+3)", []assert{locA{B, 0x10}}},

		{"LD HL,1234h : LD (0100h), HL", []assert{
			locA{Contents{Imm16(addr)}, 0x34},
			locA{Contents{Imm16(addr + 1)}, 0x12},
		}},
		{"LD HL,1234h : LD (0100h), HL : LD HL, 0000h : LD HL, (0100h)", []assert{
			locA{Contents{Imm16(addr)}, 0x34},
			locA{Contents{Imm16(addr + 1)}, 0x12},
		}},

		{"LD A,10h : LD HL, 0x0100 : LD (HL), A", []assert{locA{Contents{Imm16(addr)}, 0x10}}},
		{"LD A,10h : LD HL, 0x0100 : LD (HL), A : LD B, (HL)", []assert{locA{B, 0x10}}},

		{"LD A,10h", []assert{locA{A, 0x10}, locA{B, 0x00}}},
		{"LD B,10h", []assert{locA{B, 0x10}, locA{A, 0x00}}},
		{"LD C,10h", []assert{locA{C, 0x10}, locA{A, 0x00}}},
		{"LD D,10h", []assert{locA{D, 0x10}, locA{A, 0x00}}},
		{"LD E,10h", []assert{locA{E, 0x10}, locA{A, 0x00}}},
		{"LD H,10h", []assert{locA{H, 0x10}, locA{A, 0x00}}},
		{"LD L,10h", []assert{locA{L, 0x10}, locA{A, 0x00}}},

		// test the loc16a asserts by  testing 8bit and 16bit assertions
		{"LD BC,1234h", []assert{
			locA{B, 0x12},
			locA{C, 0x34},
			loc16A{BC, 0x1234},
			locA{A, 0x00},
		}},
		{"LD DE,1234h", []assert{
			locA{D, 0x12},
			locA{E, 0x34},
			loc16A{DE, 0x1234},
			locA{A, 0x00},
		}},
		{"LD HL,1234h", []assert{
			locA{H, 0x12},
			locA{L, 0x34},
			loc16A{HL, 0x1234},
			locA{A, 0x00},
		}},
		{"LD IX,1234h", []assert{
			locA{IXH, 0x12},
			locA{IXL, 0x34},
			loc16A{IX, 0x1234},
			locA{A, 0x00},
		}},
		{"LD IY,1234h", []assert{
			locA{IYH, 0x12},
			locA{IYL, 0x34},
			loc16A{IY, 0x1234},
			locA{A, 0x00},
		}},

		{"LD A,10h : LD B, A", []assert{locA{B, 0x10}}},
		{"LD A,10h : LD C, A", []assert{locA{C, 0x10}}},
		{"LD A,10h : LD D, A", []assert{locA{D, 0x10}}},
		{"LD A,10h : LD E, A", []assert{locA{E, 0x10}}},
		{"LD A,10h : LD H, A", []assert{locA{H, 0x10}}},
		{"LD A,10h : LD L, A", []assert{locA{L, 0x10}}},

		{"LD B,10h : LD A, B", []assert{locA{A, 0x10}}},
		{"LD C,10h : LD A, C", []assert{locA{A, 0x10}}},
		{"LD D,10h : LD A, D", []assert{locA{A, 0x10}}},
		{"LD E,10h : LD A, E", []assert{locA{A, 0x10}}},
		{"LD H,10h : LD A, H", []assert{locA{A, 0x10}}},
		{"LD L,10h : LD A, L", []assert{locA{A, 0x10}}},
	}
	for _, tc := range testCases {
		fmt.Printf("=== Assemble: %s\n", tc.prog)
		prog := tc.prog + ": HALT"
		assembly, err := Assemble(prog)
		if err != nil {
			t.Fatalf("Failed to assemble [%s]: %s", prog, err)
		}

		z := New(0x1000)
		err = z.Run(assembly)
		if err != nil {
			t.Fatalf("Failed to execute [%s]: %s", prog, err)
		}

		for _, assertion := range tc.assertions {
			err := assertion.check(z)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
