package zog

import (
	"fmt"
	"testing"
)

type executeTestCase struct {
	prog       string
	assertions []assert
}

func TestExecuteBasic(t *testing.T) {
	addr := uint16(0x100)
	memSize := uint16(0x1000)
	testCases := []executeTestCase{
		{"RST 1 : LD A, 56h : HALT : NOP : NOP : NOP : NOP : LD BC, 1234h : LD A, 00h : RET", []assert{}},
		{"CALL 000ah: LD A, 56h : HALT : NOP : NOP : NOP : NOP : LD BC, 1234h : LD A, 00h : RET", []assert{
			locA{A, 0x56},
			loc16A{BC, 0x1234},
		}},

		{"LD BC, 1234h : PUSH BC : POP DE", []assert{
			loc16A{DE, 0x1234},
		}},
		{"LD BC, 1234h : PUSH BC", []assert{
			memA{memSize - 1, 0x12},
			memA{memSize - 2, 0x34},
		}},
		{"LD B, 00h : LD A,22h : INC B : DEC A : JP NZ, 0004h", []assert{
			locA{B, 0x22},
		}},
		{"LD B, 00h : LD A,22h : INC B : DEC A : JR NZ, -4", []assert{
			locA{B, 0x22},
		}},
		{"LD B, 11h : LD A,22h : INC A : DJNZ -3", []assert{
			locA{A, 0x33},
		}},
		{"LD A,12h : EX AF,AF' : EX AF,AF'", []assert{
			locA{A, 0x12},
		}},
		{"LD A,12h : EX AF,AF'", []assert{
			locA{A, 0x00},
		}},

		{"LD HL,1111h : DEC HL", []assert{
			loc16A{HL, 0x1110},
		}},
		{"LD HL,1111h : INC HL", []assert{
			loc16A{HL, 0x1112},
		}},

		{"LD A, FFh : INC A : LD HL,4444h : LD DE, 3333h : SBC HL, DE", []assert{
			loc16A{HL, 0x1110},
		}},
		{"LD HL,4444h : LD DE, 3333h : SBC HL, DE", []assert{
			loc16A{HL, 0x1111},
		}},

		{"LD HL,1111h : LD DE, 2222h : ADD HL, DE", []assert{
			loc16A{HL, 0x3333},
		}},
		{"LD A, FFh : INC A : LD HL,1111h : LD DE, 2222h : ADC HL, DE", []assert{
			loc16A{HL, 0x3334},
		}},
		{"LD A,10h : LD B,05h : ADD A,B", []assert{
			locA{A, 0x15},
		}},
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
		prog := tc.prog + " : HALT"
		assembly, err := Assemble(prog)
		if err != nil {
			t.Fatalf("Failed to assemble [%s]: %s", prog, err)
		}

		//		buf, _ := assembly.Encode()
		//		fmt.Printf("JB - encoded buf {%s]\n", bufToHex(buf))
		z := New(memSize)
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
