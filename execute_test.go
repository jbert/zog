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

		{"LD B, 0x80 : RL B", []assert{locA{B, 0x00}, flagA{F_C, true}}},
		{"LD B, 0x01 : RL B", []assert{locA{B, 0x02}, flagA{F_C, false}}},

		{"LD B, 0x80 : RLC B", []assert{locA{B, 0x01}, flagA{F_C, true}}},
		{"LD B, 0x01 : RLC B", []assert{locA{B, 0x02}, flagA{F_C, false}}},

		{"LD A, 0x80 : RLA", []assert{locA{A, 0x00}, flagA{F_C, true}}},
		{"LD A, 0x01 : RLA", []assert{locA{A, 0x02}, flagA{F_C, false}}},

		{"LD A, 0x80 : RLCA", []assert{locA{A, 0x01}, flagA{F_C, true}}},
		{"LD A, 0x01 : RLCA", []assert{locA{A, 0x02}, flagA{F_C, false}}},

		{"LD HL, 0x0106 : INC (HL) : HALT : NOP", []assert{
			memA{0x0106, 0x01},
		}},

		{"LD HL, 0200h : LD A, 0x10 : LD (HL), A", []assert{
			memA{0x01ff, 0x00},
			memA{0x0200, 0x10},
			memA{0x0201, 0x00},
		}},

		{"LD HL, 1234h : LD DE, 5678h : ADD HL, DE", []assert{
			loc16A{HL, 0x68ac},
			loc16A{DE, 0x5678},
		}},

		{"LD HL, 1234h : LD DE, 5678h : EX DE, HL", []assert{
			loc16A{HL, 0x5678},
			loc16A{DE, 0x1234},
		}},

		{"LD BC, 0003h : LD HL, 0003h : LD DE, 0020h : LDD : HALT", []assert{
			memA{0x0020, 0x21},
			memA{0x0021, 0x00},
			memA{0x0022, 0x00},
			loc16A{HL, 0x0002},
			loc16A{DE, 0x001f},
			loc16A{BC, 0x0002},
		}},

		{"LD BC, 0003h : LD HL, 0003h : LD DE, 0020h : LDIR : HALT", []assert{
			memA{0x0020, 0x21},
			memA{0x0021, 0x03},
			memA{0x0022, 0x00},
			loc16A{HL, 0x0006},
			loc16A{DE, 0x0023},
			loc16A{BC, 0x0000},
		}},

		{"LD BC, 0003h : LD HL, 0003h : LD DE, 0020h : LDI : HALT", []assert{
			memA{0x0020, 0x21},
			memA{0x0021, 0x00},
			memA{0x0022, 0x00},
			loc16A{HL, 0x0004},
			loc16A{DE, 0x0021},
			loc16A{BC, 0x0002},
		}},

		{"LD HL, 0100h : LD (HL), 12h : LD A, 03h : RRD", []assert{
			locA{A, 0x02},
			memA{0x0100, 0x31},
		}},

		{"LD HL, 0100h : LD (HL), 12h : LD A, 03h : RLD", []assert{
			locA{A, 0x01},
			memA{0x0100, 0x23},
		}},

		{"LD A, 12h : NEG", []assert{locA{A, 0xee}}},

		{"LD BC, 1234h : EXX", []assert{loc16A{BC, 0x0000}}},
		{"LD BC, 1234h : EXX : EXX", []assert{loc16A{BC, 0x1234}}},

		{"LD A, 12h : CPL", []assert{locA{A, 0xed}}},

		{"LD A, 11h : RRCA", []assert{locA{A, 0x88}, flagA{F_C, true}}},
		{"LD A, 12h : RRCA", []assert{locA{A, 0x09}, flagA{F_C, false}}},

		{"SCF : CCF", []assert{flagA{F_C, false}}},
		{"SCF", []assert{flagA{F_C, true}}},

		{"LD A, 12h : RRA", []assert{locA{A, 0x09}}},
		{"LD A, 12h : RLA", []assert{locA{A, 0x24}}},
		{"LD A, 12h : RRCA", []assert{locA{A, 0x09}}},
		{"LD A, 12h : RLCA", []assert{locA{A, 0x24}}},
		{"NOP", []assert{locA{A, 0x00}}},

		{"LD B, 12h : SET 0, B", []assert{locA{B, 0x13}}},
		{"LD B, 12h : SET 1, B", []assert{locA{B, 0x12}}},
		{"LD B, 12h : SET 7, B", []assert{locA{B, 0x92}}},
		{"LD B, 12h : SET 4, B", []assert{locA{B, 0x12}}},

		{"LD B, 12h : RES 0, B", []assert{locA{B, 0x12}}},
		{"LD B, 12h : RES 1, B", []assert{locA{B, 0x10}}},
		{"LD B, 12h : RES 7, B", []assert{locA{B, 0x12}}},
		{"LD B, 12h : RES 4, B", []assert{locA{B, 0x02}}},

		{"LD B, 12h : BIT 0, B", []assert{flagA{F_Z, true}}},
		{"LD B, 12h : BIT 1, B", []assert{flagA{F_Z, false}}},
		{"LD B, 12h : BIT 7, B", []assert{flagA{F_Z, true}}},
		{"LD B, 12h : BIT 4, B", []assert{flagA{F_Z, false}}},

		{"LD B, 12h : RL B : ", []assert{
			locA{B, 0x24},
		}},

		{"RST 8 : LD A, 56h : HALT : NOP : NOP : NOP : NOP : LD BC, 1234h : LD A, 00h : RET", []assert{
			locA{A, 0x56},
			loc16A{BC, 0x1234},
		}},
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

		{"SCF : LD HL,4444h : LD DE, 3333h : SBC HL, DE", []assert{
			loc16A{HL, 0x1110},
		}},
		{"LD HL,4444h : LD DE, 3333h : SBC HL, DE", []assert{
			loc16A{HL, 0x1111},
		}},

		{"LD HL,1111h : LD DE, 2222h : ADD HL, DE", []assert{
			loc16A{HL, 0x3333},
		}},
		{"SCF : LD HL,1111h : LD DE, 2222h : ADC HL, DE", []assert{
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
			flagA{F_H, true},
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
			flagA{F_H, true},
			flagA{F_PV, false},
			flagA{F_C, false},
		}},
		{"LD A,80h : DEC A", []assert{
			locA{A, 0x7f},
			flagA{F_S, false},
			flagA{F_Z, false},
			flagA{F_H, true},
			flagA{F_PV, true},
			flagA{F_C, false},
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
