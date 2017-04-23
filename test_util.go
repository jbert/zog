package zog

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func testUtilRunAll(t *testing.T, f func(t *testing.T, byteForm []byte, stringForm string)) {
	// One day
	opPrefices := []byte{0x00, 0xcb, 0xed}
	indexPrefices := []byte{0x00, 0xdd, 0xfd}

	for _, ti := range allInstructions {
		t.Run(fmt.Sprintf("%02X", ti.n), func(t *testing.T) {
			for _, opPrefix := range opPrefices {
				t.Run(fmt.Sprintf("OP%02X", opPrefix), func(t *testing.T) {
					for _, indexPrefix := range indexPrefices {
						t.Run(fmt.Sprintf("IDX%02X", indexPrefix), func(t *testing.T) {
							buf := []byte{ti.n}
							if opPrefix != 0 {
								buf = append([]byte{opPrefix}, buf...)
							}
							if indexPrefix != 0 {
								buf = append([]byte{indexPrefix}, buf...)
							}
							effectiveIndexPrefix := indexPrefix
							if opPrefix == 0xed {
								// Index prefices ignored for ED
								effectiveIndexPrefix = 0x00
							}

							// getExpected will append the IX/IY immediate byte to buf, so must
							// be called before expandImmediateData
							buf, expected := ti.getExpected(effectiveIndexPrefix, opPrefix, buf)
							buf, expected = expandImmediateData(buf, expected)
							expected = normaliseAssembly(expected)
							if expected == "" {
								t.Skip("No instruction")
							}

							hexBuf := bufToHex(buf)
							fmt.Printf("TEST [%s]: %s\n", hexBuf, expected)

							f(t, buf, expected)
						})
					}
				})
			}
		})
	}
}

func expandImmediateData(buf []byte, template string) ([]byte, string) {
	var s string = template
	if strings.Contains(s, "NN") {
		buf = append(buf, 0x34)
		buf = append(buf, 0x12)
		s = strings.Replace(s, "NN", "0x1234", 1)
	}
	if strings.Contains(s, "N") {
		buf = append(buf, 0xab)
		s = strings.Replace(s, "N", "0xab", 1)
	}
	if strings.Contains(s, "DIS") {
		buf = append(buf, 0xf0)
		s = strings.Replace(s, "DIS", "-16", 1)
	}
	return buf, s
}

func bufToHex(buf []byte) string {
	s := ""
	for _, b := range buf {
		s += fmt.Sprintf("%02X", b)
	}
	return s
}

type testInstruction struct {
	n             byte
	inst          string
	inst_after_cb string
	inst_after_ed string
}

func (tc *testInstruction) getExpected(indexPrefix byte, opPrefix byte, buf []byte) ([]byte, string) {
	var expected string
	switch opPrefix {
	case 0: // No prefix
		expected = tc.inst
	case 0xcb:
		expected = tc.inst_after_cb
	case 0xed:
		expected = tc.inst_after_ed
	default:
		panic(fmt.Sprintf("Unrecognised op prefix %02X", opPrefix))
	}

	switch indexPrefix {
	case 0xdd:
		return indexRegisterMunge(opPrefix == 0xcb, "IX", buf, expected)
	case 0xfd:
		return indexRegisterMunge(opPrefix == 0xcb, "IY", buf, expected)
	default:
		return buf, expected
	}
}

func indexRegisterMunge(isCB bool, indexRegister string, buf []byte, expected string) ([]byte, string) {
	disp := int8(-20)

	// Must match location.go:func (ic IndexedContents) String() format
	hlReplace := fmt.Sprintf("(%s%+d)", indexRegister, disp)
	hReplace := indexRegister + "h"
	lReplace := indexRegister + "l"

	if isCB {
		if len(buf) != 3 {
			panic("DDCB buf not 3 bytes long")
		}
		n := buf[2]          // save instruction
		buf[2] = byte(disp)  // then displacement
		buf = append(buf, n) // and put instruction bacl
	}
	// CB cases:
	// SET/RES n, (HL)		-> SET/RES n, (ix+d)
	// ROT (HL)				-> ROT (ix+d)

	// SET/RES m, r		-> SET/RES n, (ix+d), r
	// BIT n, r 			-> BIT n, (ix+d) 	[for all r!]
	// ROT r					-> ROT (ix+d),r

	if strings.Contains(expected, "(hl)") {
		expected = strings.Replace(expected, "(hl)", hlReplace, -1)
		// We did this above (in a different place) if this was CB
		if !isCB {
			buf = append(buf, byte(disp))
		}
	} else {
		if isCB {
			// (HL) cases taken care of
			op := strings.ToLower(expected[0:3])
			switch op {
			case "set", "res":
				expected = strings.Replace(expected, ",", ","+hlReplace+",", -1)
			case "bit":
				// Do nothing
			default:
				expected = strings.Replace(expected, " ", " "+hlReplace+",", -1)
			}
		} else {
			// Exception
			if expected != "ex de,hl" {
				expected = strings.Replace(expected, "hl", indexRegister, -1)
				expected = strings.Replace(expected, "h,", hReplace+",", -1)
				expected = strings.Replace(expected, ",h", ","+hReplace, -1)
				expected = strings.Replace(expected, " h", " "+hReplace, -1)
				expected = strings.Replace(expected, "l,", lReplace+",", -1)
				expected = strings.Replace(expected, ",l", ","+lReplace, -1)
				expected = strings.Replace(expected, " l", " "+lReplace, -1)
			}
		}
	}

	return buf, expected
}

func normaliseWhiteSpace(s string) string {
	collapseSpaces := regexp.MustCompile(" +")
	s = collapseSpaces.ReplaceAllString(s, " ")
	commaSpace := regexp.MustCompile(", ")
	s = commaSpace.ReplaceAllString(s, ",")
	return s
}
func normaliseHex(s string) string {
	// 0xabcd -> abcdh
	re := regexp.MustCompile("0x([[:xdigit:]]{1,4})")
	s = re.ReplaceAllString(s, "${1}h")
	return s
}

func normaliseAssembly(s string) string {
	s = strings.TrimSpace(s)
	s = normaliseHex(s)
	s = normaliseWhiteSpace(s)
	s = strings.ToLower(s)
	return s
}

func compareAssembly(a, b string) bool {
	a = normaliseAssembly(a)
	b = normaliseAssembly(b)
	fmt.Printf("a [%s] b [%s]\n", a, b)
	return a == b
}

func decodeToSameInstruction(a, b []byte) bool {
	sA := bufToHex(a)
	sB := bufToHex(b)

	iAs, err := DecodeBytes(a)
	if err != nil {
		fmt.Printf("Error decoding [%s]: %s", sA, err)
		return false
	}
	iBs, err := DecodeBytes(b)
	if err != nil {
		fmt.Printf("Error decoding [%s]: %s", sB, err)
		return false
	}

	if len(iAs) != 1 {
		fmt.Printf("Multiple instructions [%d] for %s", len(iAs), sA)
		return false
	}
	if len(iBs) != 1 {
		fmt.Printf("Multiple instructions [%d] for %s", len(iBs), sB)
		return false
	}

	fmt.Printf("JB - BA [%s] IA [%s] BB [%s] IB [%s]\n", bufToHex(a), iAs[0].String(), bufToHex(b), iBs[0].String())
	return iAs[0].String() == iBs[0].String()
}

func z80asmAssemble(s string) []byte {
	// stdin/stdout filter
	cmd := exec.Command("z80asm", "-o", "-")
	w, err := cmd.StdinPipe()
	if err != nil {
		panic(fmt.Sprintf("Failed to get stdin to command: %s", err))
	}
	r, err := cmd.StdoutPipe()
	if err != nil {
		panic(fmt.Sprintf("Failed to get stdout of command: %s", err))
	}

	if err := cmd.Start(); err != nil {
		panic(fmt.Sprintf("Can't start: %s", err))
	}

	go func() {
		defer w.Close()
		n, err := io.WriteString(w, s)
		if n != len(s) {
			panic("Short write - buffering?")
		}
		if err != nil {
			panic(fmt.Sprintf("Write error : %s", err))
		}
	}()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(fmt.Sprintf("Failed to read stdout: %s", err))
	}
	if err = cmd.Wait(); err != nil {
		panic(fmt.Sprintf("Failed to Wait(): %s", err))
	}

	return buf
}

var allInstructions = []testInstruction{
	{0x00, "nop", "rlc b", " 	"},
	{0x01, "ld bc,NN", "rlc c", " 	"},
	{0x02, "ld (bc),a", "rlc d", " 	"},
	{0x03, "inc bc", "rlc e", " 	"},
	{0x04, "inc b", "rlc h", " 	"},
	{0x05, "dec b", "rlc l", " 	"},
	{0x06, "ld b,N", "rlc (hl)", " 	"},
	{0x07, "rlca", "rlc a", " 	"},
	{0x08, "ex af,af'", "rrc b", " 	"},
	{0x09, "add hl,bc", "rrc c", " 	"},
	{0x0A, "ld a,(bc)", "rrc d", " 	"},
	{0x0B, "dec bc", "rrc e", " 	"},
	{0x0C, "inc c", "rrc h", " 	"},
	{0x0D, "dec c", "rrc l", " 	"},
	{0x0E, "ld c,N", "rrc (hl)", " 	"},
	{0x0F, "rrca", "rrc a", " 	"},
	{0x10, "djnz DIS", "rl b", " 	"},
	{0x11, "ld de,NN", "rl c", " 	"},
	{0x12, "ld (de),a", "rl d", " 	"},
	{0x13, "inc de", "rl e", " 	"},
	{0x14, "inc d", "rl h", " 	"},
	{0x15, "dec d", "rl l", " 	"},
	{0x16, "ld d,N", "rl (hl)", " 	"},
	{0x17, "rla", "rl a", " 	"},
	{0x18, "jr DIS", "rr b", " 	"},
	{0x19, "add hl,de", "rr c", " 	"},
	{0x1A, "ld a,(de)", "rr d", " 	"},
	{0x1B, "dec de", "rr e", " 	"},
	{0x1C, "inc e", "rr h", " 	"},
	{0x1D, "dec e", "rr l", " 	"},
	{0x1E, "ld e,N", "rr (hl)", " 	"},
	{0x1F, "rra", "rr a", " 	"},
	{0x20, "jr nz, DIS", "sla b", " 	"},
	{0x21, "ld hl,NN", "sla c", " 	"},
	{0x22, "ld (NN),hl", "sla d", " 	"},
	{0x23, "inc hl", "sla e", " 	"},
	{0x24, "inc h", "sla h", " 	"},
	{0x25, "dec h", "sla l", ""},
	{0x26, "ld h,N", "sla (hl)", " 	"},
	{0x27, "daa", "sla a", " 	"},
	{0x28, "jr z,DIS", "sra b", " 	"},
	{0x29, "add hl,hl", "sra c", " 	"},
	{0x2A, "ld hl,(NN)", "sra d", ""},
	{0x2B, "dec hl", "sra e", ""},
	{0x2C, "inc l", "sra h", ""},
	{0x2D, "dec l", "sra l", ""},
	{0x2E, "ld l,N", "sra (hl)", ""},
	{0x2F, "cpl", "sra a", ""},
	{0x30, "jr nc,DIS", "sll b", ""},
	{0x31, "ld sp,NN", "sll c", ""},
	{0x32, "ld (NN),a", "sll d", ""},
	{0x33, "inc sp", "sll e", ""},
	{0x34, "inc (hl)", "sll h", ""},
	{0x35, "dec (hl)", "sll l", ""},
	{0x36, "ld (hl),N", "sll (hl)", ""},
	{0x37, "scf", "sll a", ""},
	{0x38, "jr c,DIS", "srl b", ""},
	{0x39, "add hl,sp", "srl c", ""},
	{0x3A, "ld a,(NN)", "srl d", ""},
	{0x3B, "dec sp", "srl e", ""},
	{0x3C, "inc a", "srl h", ""},
	{0x3D, "dec a", "srl l", ""},
	{0x3E, "ld a,N", "srl (hl)", ""},
	{0x3F, "ccf", "srl a", ""},
	{0x40, "ld b,b", "bit 0,b", "in b,(c)	"},
	{0x41, "ld b,c", "bit 0,c", "out (c),b	"},
	{0x42, "ld b,d", "bit 0,d", "sbc hl,bc	"},
	{0x43, "ld b,e", "bit 0,e", "ld (NN),bc	"},
	{0x44, "ld b,h", "bit 0,h", "neg	"},
	{0x45, "ld b,l", "bit 0,l", "retn	"},
	{0x46, "ld b,(hl)", "bit 0,(hl)", "im 0	"},
	{0x47, "ld b,a", "bit 0,a", "ld i,a	"},
	{0x48, "ld c,b", "bit 1,b", "in c,(c)	"},
	{0x49, "ld c,c", "bit 1,c", "out (c),c	"},
	{0x4A, "ld c,d", "bit 1,d", "adc hl,bc	"},
	{0x4B, "ld c,e", "bit 1,e", "ld bc,(NN)	"},
	{0x4C, "ld c,h", "bit 1,h", ""},
	{0x4D, "ld c,l", "bit 1,l", "reti	"},
	{0x4E, "ld c,(hl)", "bit 1,(hl)", ""},
	{0x4F, "ld c,a", "bit 1,a", "ld r,a	"},
	{0x50, "ld d,b", "bit 2,b", "in d,(c)	"},
	{0x51, "ld d,c", "bit 2,c", "out (c),d	"},
	{0x52, "ld d,d", "bit 2,d", "sbc hl,de	"},
	{0x53, "ld d,e", "bit 2,e", "ld (NN),de	"},
	{0x54, "ld d,h", "bit 2,h", ""},
	{0x55, "ld d,l", "bit 2,l", ""},
	{0x56, "ld d,(hl)", "bit 2,(hl)", "im 1	"},
	{0x57, "ld d,a", "bit 2,a", "ld a,i	"},
	{0x58, "ld e,b", "bit 3,b", "in e,(c)	"},
	{0x59, "ld e,c", "bit 3,c", "out (c),e	"},
	{0x5A, "ld e,d", "bit 3,d", "adc hl,de	"},
	{0x5B, "ld e,e", "bit 3,e", "ld de,(NN)	"},
	{0x5C, "ld e,h", "bit 3,h", ""},
	{0x5D, "ld e,l", "bit 3,l", ""},
	{0x5E, "ld e,(hl)", "bit 3,(hl)", "im 2	"},
	{0x5F, "ld e,a", "bit 3,a", "ld a,r	"},
	{0x60, "ld h,b", "bit 4,b", "in h,(c)	"},
	{0x61, "ld h,c", "bit 4,c", "out (c),h	"},
	{0x62, "ld h,d", "bit 4,d", "sbc hl,hl	"},
	{0x63, "ld h,e", "bit 4,e", "ld (NN),hl	"},
	{0x64, "ld h,h", "bit 4,h", ""},
	{0x65, "ld h,l", "bit 4,l", ""},
	{0x66, "ld h,(hl)", "bit 4,(hl)", ""},
	{0x67, "ld h,a", "bit 4,a", "rrd	"},
	{0x68, "ld l,b", "bit 5,b", "in l,(c)	"},
	{0x69, "ld l,c", "bit 5,c", "out (c),l	"},
	{0x6A, "ld l,d", "bit 5,d", "adc hl,hl	"},
	{0x6B, "ld l,e", "bit 5,e", "ld hl,(NN)"},
	{0x6C, "ld l,h", "bit 5,h", ""},
	{0x6D, "ld l,l", "bit 5,l", ""},
	{0x6E, "ld l,(hl)", "bit 5,(hl)", ""},
	{0x6F, "ld l,a", "bit 5,a", "rld	"},
	{0x70, "ld (hl),b", "bit 6,b", "in f,(c)	"},
	{0x71, "ld (hl),c", "bit 6,c", ""},
	{0x72, "ld (hl),d", "bit 6,d", "sbc hl,sp	"},
	{0x73, "ld (hl),e", "bit 6,e", "ld (NN),sp	"},
	{0x74, "ld (hl),h", "bit 6,h", ""},
	{0x75, "ld (hl),l", "bit 6,l", ""},
	{0x76, "halt", "bit 6,(hl)", ""},
	{0x77, "ld (hl),a", "bit 6,a", ""},
	{0x78, "ld a,b", "bit 7,b", "in a,(c)	"},
	{0x79, "ld a,c", "bit 7,c", "out (c),a	"},
	{0x7A, "ld a,d", "bit 7,d", "adc hl,sp	"},
	{0x7B, "ld a,e", "bit 7,e", "ld sp,(NN)	"},
	{0x7C, "ld a,h", "bit 7,h", ""},
	{0x7D, "ld a,l", "bit 7,l", ""},
	{0x7E, "ld a,(hl)", "bit 7,(hl)", ""},
	{0x7F, "ld a,a", "bit 7,a", ""},
	{0x80, "add a,b", "res 0,b", ""},
	{0x81, "add a,c", "res 0,c", ""},
	{0x82, "add a,d", "res 0,d", ""},
	{0x83, "add a,e", "res 0,e", ""},
	{0x84, "add a,h", "res 0,h", ""},
	{0x85, "add a,l", "res 0,l", ""},
	{0x86, "add a,(hl)", "res 0,(hl)", ""},
	{0x87, "add a,a", "res 0,a", ""},
	{0x88, "adc a,b", "res 1,b", ""},
	{0x89, "adc a,c", "res 1,c", ""},
	{0x8A, "adc a,d", "res 1,d", ""},
	{0x8B, "adc a,e", "res 1,e", ""},
	{0x8C, "adc a,h", "res 1,h", ""},
	{0x8D, "adc a,l", "res 1,l", ""},
	{0x8E, "adc a,(hl)", "res 1,(hl)", ""},
	{0x8F, "adc a,a", "res 1,a", ""},
	{0x90, "sub b", "res 2,b", ""},
	{0x91, "sub c", "res 2,c", ""},
	{0x92, "sub d", "res 2,d", ""},
	{0x93, "sub e", "res 2,e", ""},
	{0x94, "sub h", "res 2,h", ""},
	{0x95, "sub l", "res 2,l", ""},
	{0x96, "sub (hl)", "res 2,(hl)", ""},
	{0x97, "sub a", "res 2,a", ""},
	{0x98, "sbc a,b", "res 3,b", ""},
	{0x99, "sbc a,c", "res 3,c", ""},
	{0x9A, "sbc a,d", "res 3,d", ""},
	{0x9B, "sbc a,e", "res 3,e", ""},
	{0x9C, "sbc a,h", "res 3,h", ""},
	{0x9D, "sbc a,l", "res 3,l", ""},
	{0x9E, "sbc a,(hl)", "res 3,(hl)", ""},
	{0x9F, "sbc a,a", "res 3,a", ""},
	{0xA0, "and b", "res 4,b", "ldi	"},
	{0xA1, "and c", "res 4,c", "cpi	"},
	{0xA2, "and d", "res 4,d", "ini	"},
	{0xA3, "and e", "res 4,e", "outi	"},
	{0xA4, "and h", "res 4,h", ""},
	{0xA5, "and l", "res 4,l", ""},
	{0xA6, "and (hl)", "res 4,(hl)", ""},
	{0xA7, "and a", "res 4,a", ""},
	{0xA8, "xor b", "res 5,b", "ldd	"},
	{0xA9, "xor c", "res 5,c", "cpd	"},
	{0xAA, "xor d", "res 5,d", "ind	"},
	{0xAB, "xor e", "res 5,e", "outd	"},
	{0xAC, "xor h", "res 5,h", ""},
	{0xAD, "xor l", "res 5,l", ""},
	{0xAE, "xor (hl)", "res 5,(hl)", ""},
	{0xAF, "xor a", "res 5,a", ""},
	{0xB0, "or b", "res 6,b", "ldir	"},
	{0xB1, "or c", "res 6,c", "cpir	"},
	{0xB2, "or d", "res 6,d", "inir	"},
	{0xB3, "or e", "res 6,e", "otir	"},
	{0xB4, "or h", "res 6,h", ""},
	{0xB5, "or l", "res 6,l", ""},
	{0xB6, "or (hl)", "res 6,(hl)", ""},
	{0xB7, "or a", "res 6,a", ""},
	{0xB8, "cp b", "res 7,b", "lddr	"},
	{0xB9, "cp c", "res 7,c", "cpdr	"},
	{0xBA, "cp d", "res 7,d", "indr	"},
	{0xBB, "cp e", "res 7,e", "otdr	"},
	{0xBC, "cp h", "res 7,h", ""},
	{0xBD, "cp l", "res 7,l", ""},
	{0xBE, "cp (hl)", "res 7,(hl)", ""},
	{0xBF, "cp a", "res 7,a", ""},
	{0xC0, "ret nz", "set 0,b", ""},
	{0xC1, "pop bc", "set 0,c", ""},
	{0xC2, "jp nz,NN", "set 0,d", ""},
	{0xC3, "jp NN", "set 0,e", ""},
	{0xC4, "call nz,NN", "set 0,h", ""},
	{0xC5, "push bc", "set 0,l", ""},
	{0xC6, "add a,N", "set 0,(hl)", ""},
	{0xC7, "rst 0", "set 0,a", ""},
	{0xC8, "ret z", "set 1,b", ""},
	{0xC9, "ret", "set 1,c", ""},
	{0xCA, "jp z,NN", "set 1,d", ""},
	{0xCB, "", "set 1,e", ""},
	{0xCC, "call z,NN", "set 1,h", ""},
	{0xCD, "call NN", "set 1,l", ""},
	{0xCE, "adc a,N", "set 1,(hl)", ""},
	{0xCF, "rst 8", "set 1,a", ""},
	{0xD0, "ret nc", "set 2,b", ""},
	{0xD1, "pop de", "set 2,c", ""},
	{0xD2, "jp nc, NN", "set 2,d", ""},
	{0xD3, "out (N),a", "set 2,e", ""},
	{0xD4, "call nc,NN", "set 2,h", ""},
	{0xD5, "push de", "set 2,l", ""},
	{0xD6, "sub N", "set 2,(hl)", ""},
	{0xD7, "rst 16", "set 2,a", ""},
	{0xD8, "ret c", "set 3,b", ""},
	{0xD9, "exx", "set 3,c", ""},
	{0xDA, "jp c,NN", "set 3,d", ""},
	{0xDB, "in a,(N)", "set 3,e", ""},
	{0xDC, "call c,NN", "set 3,h", ""},
	{0xDD, "", "set 3,l", ""},
	{0xDE, "sbc a,N", "set 3,(hl)", ""},
	{0xDF, "rst 24", "set 3,a", ""},
	{0xE0, "ret po", "set 4,b", ""},
	{0xE1, "pop hl", "set 4,c", ""},
	{0xE2, "jp po,NN", "set 4,d", ""},
	{0xE3, "ex (sp),hl", "set 4,e", ""},
	{0xE4, "call po,NN", "set 4,h", ""},
	{0xE5, "push hl", "set 4,l", ""},
	{0xE6, "and N", "set 4,(hl)", ""},
	{0xE7, "rst 32", "set 4,a", ""},
	{0xE8, "ret pe", "set 5,b", ""},
	{0xE9, "jp hl", "set 5,c", ""},
	{0xEA, "jp pe,NN", "set 5,d", ""},
	{0xEB, "ex de,hl", "set 5,e", ""},
	{0xEC, "call pe,NN", "set 5,h", ""},
	{0xED, "", "set 5,l", ""},
	{0xEE, "xor N", "set 5,(hl)", ""},
	{0xEF, "rst 40", "set 5,a", ""},
	{0xF0, "ret p", "set 6,b", ""},
	{0xF1, "pop af", "set 6,c", ""},
	{0xF2, "jp p,NN", "set 6,d", ""},
	{0xF3, "di", "set 6,e", ""},
	{0xF4, "call p,NN", "set 6,h", ""},
	{0xF5, "push af", "set 6,l", ""},
	{0xF6, "or N", "set 6,(hl)", ""},
	{0xF7, "rst 48", "set 6,a", ""},
	{0xF8, "ret m", "set 7,b", ""},
	{0xF9, "ld sp,hl", "set 7,c", ""},
	{0xFA, "jp m,NN", "set 7,d", ""},
	{0xFB, "ei", "set 7,e", ""},
	{0xFC, "call m,NN", "set 7,h", ""},
	{0xFD, "", "set 7,l", ""},
	{0xFE, "cp N", "set 7,(hl)", ""},
	{0xFF, "rst 56", "set 7,a", " 	"},
}
