package zog

import (
	"errors"
	"fmt"
	"strings"
)

type Instruction interface {
	String() string
	Encode() []byte
	Resolve(a *Assembly) error
	Execute(z *Zog) error
}

type LabelHolder struct{}

func (lh *LabelHolder) String() string            { return "" }
func (lh *LabelHolder) Encode() []byte            { return make([]byte, 0) }
func (lh *LabelHolder) Resolve(a *Assembly) error { return nil }
func (lh *LabelHolder) Execute(z *Zog) error      { panic("Attempt to execute labelholder") }

type Data struct {
	data []byte
}

func NewData(data []byte) *Data {
	return &Data{data: data}
}
func (d *Data) String() string {
	return bufToHex(d.data)
}
func (d *Data) Encode() []byte {
	return d.data
}
func (d *Data) Resolve(a *Assembly) error {
	return nil
}
func (d *Data) Execute(z *Zog) error {
	return errors.New("Error - trying to execute dummy data instruction")
}

type LD8 struct {
	InstBin8
}

func NewLD8(dst Loc8, src Loc8) *LD8 {
	return &LD8{InstBin8{dst: dst, src: src}}
}

func (l *LD8) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}
func (l *LD8) Encode() []byte {
	// ED special cases
	switch true {
	case l.dst == I && l.src == A:
		return []byte{0xed, 0x47}
	case l.dst == A && l.src == I:
		return []byte{0xed, 0x57}
	case l.dst == R && l.src == A:
		return []byte{0xed, 0x4f}
	case l.dst == A && l.src == R:
		return []byte{0xed, 0x5f}
	}

	l.inspect()

	switch l.dstInfo.ltype {
	case BCDEContents:
		// LD (BC), A or LD (DE), A
		p := byte(1)
		if l.dstInfo.isBC {
			p = 0
		}
		buf := []byte{encodeXPQZ(0, p, 0, 2)}
		return buf
	case ImmediateContents:
		// LD (nn), A
		buf := []byte{encodeXPQZ(0, 3, 0, 2)}
		buf = append(buf, l.dstInfo.imm16...)
		return buf
	}

	if l.dstInfo.ltype != tableR {
		panic("Non-tableR dst in LD8")
	}
	switch l.srcInfo.ltype {
	case tableR:
		b := encodeXYZ(1, l.dstInfo.idxTable, l.srcInfo.idxTable)
		return idxEncodeHelper([]byte{b}, l.idx)
	case Immediate:
		buf := []byte{encodeXYZ(0, l.dstInfo.idxTable, 6)}
		buf = idxEncodeHelper(buf, l.idx)
		buf = append(buf, l.srcInfo.imm8)
		return buf
	case BCDEContents:
		// LD A, (BC) or LD A, (DE)
		p := byte(1)
		if l.srcInfo.isBC {
			p = 0
		}
		b := encodeXPQZ(0, p, 1, 2)
		return []byte{b}
	case ImmediateContents:
		// LD A, (nn)
		buf := []byte{encodeXPQZ(0, 3, 1, 2)}
		buf = append(buf, l.srcInfo.imm16...)
		return buf
	default:
		panic("Unknown src type in LD8")
	}
}
func (l *LD8) Execute(z *Zog) error {
	// Flags are unchanged for LD
	f, err := F.Read8(z)
	if err != nil {
		return err
	}
	err = l.exec(z, func(v byte) byte { return v })
	if err != nil {
		return err
	}
	/*
		Whem a Load Register A with Register I (LD A, I) instruction or a Load Register A with Register
		R (LD A, R) instruction is executed, the state of IFF2 is copied to the parity flag, where it
		can be tested or stored.
	*/
	z.SetFlag(F_PV, z.iff2)
	return F.Write8(z, f)
}

type INC8 struct {
	InstU8
}

func NewINC8(l Loc8) *INC8 {
	return &INC8{InstU8{l: l}}
}
func (i *INC8) String() string {
	return fmt.Sprintf("INC %s", i.l)
}
func (i *INC8) Encode() []byte {
	i.inspect()
	if i.lInfo.ltype != tableR {
		panic("Non-tableR INC8")
	}
	b := encodeXYZ(0, i.lInfo.idxTable, 4)
	return idxEncodeHelper([]byte{b}, i.idx)
}
func (i *INC8) Execute(z *Zog) error {
	err := i.exec(z, func(v byte) byte {
		z.SetFlag(F_H, v&0xf == 0x0f)
		z.SetFlag(F_PV, v == 0x7f)
		z.SetFlag(F_N, false)
		return v + 1
	})
	return err
}

type DEC8 struct {
	InstU8
}

func NewDEC8(l Loc8) *DEC8 {
	return &DEC8{InstU8{l: l}}
}
func (d *DEC8) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}
func (d *DEC8) Encode() []byte {
	d.inspect()
	if d.lInfo.ltype != tableR {
		panic("Non-tableR DEC8")
	}
	b := encodeXYZ(0, d.lInfo.idxTable, 5)
	return idxEncodeHelper([]byte{b}, d.idx)
}
func (d *DEC8) Execute(z *Zog) error {
	err := d.exec(z, func(v byte) byte {
		z.SetFlag(F_H, v&0x0f == 0x00)
		z.SetFlag(F_PV, v == 0x80)
		z.SetFlag(F_N, true)
		return v - 1
	})
	return err
}

type LD16 struct {
	InstBin16
}

func NewLD16(dst, src Loc16) *LD16 {
	return &LD16{InstBin16: InstBin16{dst: dst, src: src}}
}
func (l *LD16) String() string {
	return fmt.Sprintf("LD %s, %s", l.dst, l.src)
}
func (l *LD16) Encode() []byte {
	l.inspect()
	switch l.dstInfo.ltype {
	case ImmediateContents:
		// LD (nn), HL has multiple encodings, we choose the non-ED one
		if l.srcInfo.isHLLike() {
			buf := []byte{encodeXPQZ(0, 2, 0, 2)}
			buf = append(buf, l.dstInfo.imm16...)
			return idxEncodeHelper(buf, l.idx)
		} else {
			if l.srcInfo.ltype != tableRP {
				panic("Non-tableRP src in LD16 (NN), src")
			}
			buf := []byte{0xed, encodeXPQZ(1, l.srcInfo.idxTable, 0, 3)}
			buf = append(buf, l.dstInfo.imm16...)
			return buf
		}
	}

	if l.dstInfo.ltype != tableRP {
		panic("Non-tableRP dst in LD16")
	}

	switch l.srcInfo.ltype {
	case Immediate:
		buf := []byte{encodeXPQZ(0, l.dstInfo.idxTable, 0, 1)}
		buf = append(buf, l.srcInfo.imm16...)
		return idxEncodeHelper(buf, l.idx)
	case ImmediateContents:
		// LD HL, (nn) has multiple encodings
		if l.dstInfo.isHLLike() {
			buf := []byte{encodeXPQZ(0, 2, 1, 2)}
			buf = append(buf, l.srcInfo.imm16...)
			return idxEncodeHelper(buf, l.idx)
		} else {
			if l.dstInfo.ltype != tableRP {
				panic("Non-tableRP src in LD16 (NN), src")
			}
			buf := []byte{0xed, encodeXPQZ(1, l.dstInfo.idxTable, 1, 3)}
			buf = append(buf, l.srcInfo.imm16...)
			return buf
		}
	case tableRP:
		if l.srcInfo.isHLLike() {
			if l.dst != SP {
				panic("HL-like load to non-SP")
			}
			buf := []byte{encodeXPQZ(3, 3, 1, 1)}
			return idxEncodeHelper(buf, l.idx)
		} else {
			panic("Non-HL like load to something")
		}
	default:
		panic("Unknown src type in LD16")
	}
}
func (l *LD16) Execute(z *Zog) error {
	nn, err := l.src.Read16(z)
	if err != nil {
		return fmt.Errorf("LD16: failed to read: %s", err)
	}
	err = l.dst.Write16(z, nn)
	if err != nil {
		return fmt.Errorf("LD16: failed to write: %s", err)
	}
	return nil
}

type ADD16 struct {
	InstBin16
}

func NewADD16(dst, src Loc16) *ADD16 {
	return &ADD16{InstBin16: InstBin16{dst: dst, src: src}}
}
func (a *ADD16) String() string {
	return fmt.Sprintf("ADD %s, %s", a.dst, a.src)
}
func (a *ADD16) Encode() []byte {
	a.inspect()
	if a.dstInfo.ltype != tableRP {
		panic("Non-tableRP dst in ADD16")
	}
	if a.srcInfo.ltype != tableRP {
		panic("Non-tableRP src in ADD16")
	}

	if !a.dstInfo.isHLLike() {
		panic("Non-HL dst in ADD16")
	}
	switch a.srcInfo.ltype {
	case tableRP:
		buf := []byte{encodeXPQZ(0, a.srcInfo.idxTable, 1, 1)}
		return idxEncodeHelper(buf, a.idx)
	default:
		panic("Unknown src type in ADD16")
	}
}
func (a *ADD16) Execute(z *Zog) error {
	return a.exec(z, func(a, b uint16) uint16 {
		v := a + b
		z.SetFlag(F_H, ((a&0x0fff)+(b&0x0fff))&0x1000 != 0)
		z.SetFlag(F_N, false)
		z.SetFlag(F_C, int(a)+int(b) > 0xffff)
		return v
	})
}

type ADC16 struct {
	InstBin16
}

func NewADC16(dst, src Loc16) *ADC16 {
	return &ADC16{InstBin16: InstBin16{dst: dst, src: src}}
}
func (a *ADC16) String() string {
	return fmt.Sprintf("ADC %s, %s", a.dst, a.src)
}
func (a *ADC16) Encode() []byte {
	a.inspect()
	if a.srcInfo.ltype != tableRP {
		panic("Non-tableRP src in ADC16")
	}
	buf := []byte{0xed, encodeXPQZ(1, a.srcInfo.idxTable, 1, 2)}
	return idxEncodeHelper(buf, a.idx)
}
func (a *ADC16) Execute(z *Zog) error {
	return a.exec(z, func(a, b uint16) uint16 {
		c := uint16(0)
		if z.GetFlag(F_C) {
			c = 1
		}
		v32 := uint32(a) + uint32(b) + uint32(c)
		v := uint16(v32)
		z.SetFlag(F_S, v >= 0x8000)
		z.SetFlag(F_Z, v == 0)
		z.SetFlag(F_H, ((a&0x0fff)+(b&0x0fff)+c)&0x1000 != 0)
		overflow := (isPos16(a) && isPos16(b) && !isPos16(v)) || (!isPos16(a) && !isPos16(b) && isPos16(v))
		z.SetFlag(F_PV, overflow)
		z.SetFlag(F_N, false)
		z.SetFlag(F_C, v32 >= 0x10000)
		//		fmt.Fprintf(os.Stderr, "JB ADC16 a %04x b %04x v %04x (carry %d)\n", a, b, v, c)
		return v
	})
}

type SBC16 struct {
	InstBin16
}

func NewSBC16(dst, src Loc16) *SBC16 {
	return &SBC16{InstBin16: InstBin16{dst: dst, src: src}}
}
func (s *SBC16) String() string {
	return fmt.Sprintf("SBC %s, %s", s.dst, s.src)
}
func (s *SBC16) Encode() []byte {
	s.inspect()
	if s.srcInfo.ltype != tableRP {
		panic("Non-tableRP src in SBC16")
	}
	buf := []byte{0xed, encodeXPQZ(1, s.srcInfo.idxTable, 0, 2)}
	return idxEncodeHelper(buf, s.idx)
}

func isPos16(v uint16) bool {
	return v&0x8000 == 0
}
func (s *SBC16) Execute(z *Zog) error {
	return s.exec(z, func(a, b uint16) uint16 {
		c := uint16(0)
		if z.GetFlag(F_C) {
			c = 1
		}
		v32 := uint32(a) - uint32(b) - uint32(c)
		v := uint16(v32)
		z.SetFlag(F_S, !isPos16(v))
		z.SetFlag(F_Z, v == 0)
		z.SetFlag(F_H, ((a&0x0fff)-(b&0x0fff)-c)&0x1000 != 0)
		overflow := (isPos16(a) && !isPos16(v)) || (!isPos16(a) && isPos16(v))
		z.SetFlag(F_PV, overflow)
		z.SetFlag(F_N, true)
		z.SetFlag(F_C, isPos16(a) && !isPos16(v))
		return v
	})
}

type INC16 struct {
	InstU16
}

func NewINC16(l Loc16) *INC16 {
	return &INC16{InstU16{l: l}}
}

func (i *INC16) String() string {
	return fmt.Sprintf("INC %s", i.l)
}
func (i *INC16) Encode() []byte {
	i.inspect()
	if i.lInfo.ltype != tableRP {
		panic("Non-tableRP INC16")
	}
	b := encodeXPQZ(0, i.lInfo.idxTable, 0, 3)
	return idxEncodeHelper([]byte{b}, i.idx)
}
func (i *INC16) Execute(z *Zog) error {
	err := i.exec(z, func(v uint16) uint16 {
		return v + 1
	})
	return err
}

type DEC16 struct {
	InstU16
}

func NewDEC16(l Loc16) *DEC16 {
	return &DEC16{InstU16{l: l}}
}
func (d *DEC16) String() string {
	return fmt.Sprintf("DEC %s", d.l)
}
func (d *DEC16) Encode() []byte {
	d.inspect()
	if d.lInfo.ltype != tableRP {
		panic("Non-tableRP DEC16")
	}
	b := encodeXPQZ(0, d.lInfo.idxTable, 1, 3)
	return idxEncodeHelper([]byte{b}, d.idx)
}
func (d *DEC16) Execute(z *Zog) error {
	err := d.exec(z, func(v uint16) uint16 {
		return v - 1
	})
	return err
}

type EX struct {
	InstBin16
}

func NewEX(dst, src Loc16) *EX {
	return &EX{InstBin16: InstBin16{dst: dst, src: src}}
}

func (ex *EX) String() string {
	return fmt.Sprintf("EX %s, %s", ex.dst, ex.src)
}
func (ex *EX) Encode() []byte {
	if ex.dst == AF && ex.src == AF_PRIME {
		return []byte{0x08}
	} else if ex.dst.String() == (Contents{SP}).String() {

		var info loc16Info
		var idx idxInfo
		inspectLoc16(ex.src, &info, &idx, false)
		buf := []byte{encodeXYZ(3, 4, 3)}
		return idxEncodeHelper(buf, idx)
	} else if ex.dst == DE && ex.src == HL {
		// EX DE,HL is an excpetion to the IX/IY rule
		return []byte{encodeXYZ(3, 5, 3)}
	}

	panic("Unrecognised EX instruction")
}
func (ex *EX) Execute(z *Zog) error {
	a, err := ex.src.Read16(z)
	if err != nil {
		return fmt.Errorf("%s : can't read src: %s", ex, ex.src, err)
	}
	b, err := ex.dst.Read16(z)
	if err != nil {
		return fmt.Errorf("%s : can't read dst: %s", ex, ex.dst, err)
	}

	err = ex.dst.Write16(z, a)
	if err != nil {
		return fmt.Errorf("%s : can't write dst: %s", ex, ex.dst, err)
	}
	err = ex.src.Write16(z, b)
	if err != nil {
		return fmt.Errorf("%s : can't write dst: %s", ex, ex.dst, err)
	}
	return nil
}

type DJNZ struct {
	d Disp
}

func (d *DJNZ) String() string {
	return fmt.Sprintf("DJNZ %s", d.d)
}
func (d *DJNZ) Encode() []byte {
	b := encodeXYZ(0, 2, 0)
	return []byte{b, byte(d.d)}
}
func (d *DJNZ) Resolve(a *Assembly) error {
	return nil
}
func (d *DJNZ) Execute(z *Zog) error {
	bReg, err := B.Read8(z)
	if err != nil {
		return fmt.Errorf("Can't read B: %s", err)
	}
	bReg--
	err = B.Write8(z, bReg)
	if err != nil {
		return fmt.Errorf("Can't write B: %s", err)
	}
	zero := bReg == 0
	if !zero {
		z.jr(int8(d.d))
	}
	return nil
}

type JR struct {
	c Conditional
	d Disp
}

func (j *JR) String() string {
	if j.c == True || j.c == nil {
		return fmt.Sprintf("JR %s", j.d)
	} else {
		return fmt.Sprintf("JR %s, %s", j.c, j.d)
	}
}
func (j *JR) Encode() []byte {
	var y byte
	if j.c == True || j.c == nil {
		y = 3
	} else {
		y = findInTableCC(j.c)
		y += 4
	}
	b := encodeXYZ(0, y, 0)
	return []byte{b, byte(j.d)}
}
func (j *JR) Resolve(a *Assembly) error {
	return nil
}
func (j *JR) Execute(z *Zog) error {
	takeJump := j.c.IsTrue(z)
	if takeJump {
		z.jr(int8(j.d))
	}
	return nil
}

type JP struct {
	InstU16
	c Conditional
}

func NewJP(c Conditional, l Loc16) *JP {
	return &JP{InstU16: InstU16{l: l}, c: c}
}

func (jp *JP) String() string {
	if jp.c == True || jp.c == nil {
		return fmt.Sprintf("JP %s", jp.l)
	} else {
		return fmt.Sprintf("JP %s, %s", jp.c, jp.l)
	}
}
func (jp *JP) Encode() []byte {
	jp.inspect()
	if jp.c == True || jp.c == nil {
		if jp.lInfo.isHLLike() {
			buf := []byte{encodeXPQZ(3, 2, 1, 1)}
			return idxEncodeHelper(buf, jp.idx)
		}
	}
	if jp.lInfo.ltype != Immediate {
		panic("Non-immediate (or direct HL-like) JP")
	}

	var buf []byte
	if jp.c == True || jp.c == nil {
		buf = []byte{encodeXYZ(3, 0, 3)}
	} else {
		y := findInTableCC(jp.c)
		buf = []byte{encodeXYZ(3, y, 2)}
	}
	buf = append(buf, jp.lInfo.imm16...)
	return buf
}
func (jp *JP) Execute(z *Zog) error {
	takeJump := jp.c.IsTrue(z)
	if takeJump {
		addr, err := jp.l.Read16(z)
		if err != nil {
			return err
		}
		z.jp(addr)
	}
	return nil
}

type CALL struct {
	InstU16
	c Conditional
}

func NewCALL(c Conditional, l Loc16) *CALL {
	return &CALL{InstU16: InstU16{l: l}, c: c}
}
func (c *CALL) String() string {
	if c.c == True || c.c == nil {
		return fmt.Sprintf("CALL %s", c.l)
	} else {
		return fmt.Sprintf("CALL %s, %s", c.c, c.l)
	}
}
func (c *CALL) Encode() []byte {
	c.inspect()
	var buf []byte
	if c.c == nil || c.c == True {
		buf = []byte{encodeXPQZ(3, 0, 1, 5)}
	} else {
		y := findInTableCC(c.c)
		buf = []byte{encodeXYZ(3, y, 4)}
	}
	buf = append(buf, c.lInfo.imm16...)
	return buf
}
func (c *CALL) Execute(z *Zog) error {
	takeJump := c.c.IsTrue(z)
	if takeJump {
		addr, err := c.l.Read16(z)
		if err != nil {
			return err
		}
		z.push(z.reg.PC)
		z.jp(addr)
	}
	return nil
}

type OUT struct {
	port  Loc8
	value Loc8
}

func (o *OUT) String() string {
	return fmt.Sprintf("OUT (%s), %s", o.port, o.value)
}
func (o *OUT) Encode() []byte {
	if o.port == C {
		var info loc8Info
		var idx idxInfo
		inspectLoc8(o.value, &info, &idx)
		if info.ltype != tableR {
			panic("Non-tableR value in OUT")
		}
		// (HL)? IX?
		buf := []byte{0xed, encodeXYZ(1, info.idxTable, 1)}
		return idxEncodeHelper(buf, idx)
	} else {
		imm8 := o.port.(Imm8)
		return []byte{encodeXYZ(3, 2, 3), byte(imm8)}
	}
}
func (o *OUT) Resolve(a *Assembly) error {
	return nil
}
func (o *OUT) Execute(z *Zog) error {
	/*
		In the IN A and OUT n, A instructions, the I/O device’s n address appears in the lower half
		of the address bus (A7–A0), while the Accumulator content is transferred in the upper half
		of the address bus. In all Register Indirect input output instructions, including block I/O
		transfers, the contents of the C Register are transferred to the lower half of the address bus
		(device address) while the contents of Register B are transferred to the upper half of the
		address bus.
	*/
	var addr uint16
	v, err := o.value.Read8(z)
	if err != nil {
		return err
	}
	if o.port == C {
		addr = z.reg.Read16(BC)
	} else {
		addr = uint16(v) | (uint16(z.reg.A) << 8)
	}
	z.out(addr, v)
	return nil
}

type IN struct {
	dst  Loc8
	port Loc8
}

func (i *IN) String() string {
	return fmt.Sprintf("IN %s, (%s)", i.dst, i.port)
}
func (i *IN) Encode() []byte {
	if i.port == C {
		var y byte
		var info loc8Info
		var idx idxInfo
		if i.dst == F {
			y = 6
		} else {
			inspectLoc8(i.dst, &info, &idx)
			if info.ltype != tableR {
				panic("Non-tableR dst in IN")
			}
			y = info.idxTable
		}
		buf := []byte{0xed, encodeXYZ(1, y, 0)}
		return idxEncodeHelper(buf, idx)
	} else {
		imm8 := i.port.(Imm8)
		return []byte{encodeXYZ(3, 3, 3), byte(imm8)}
	}
}
func (i *IN) Resolve(a *Assembly) error {
	return nil
}
func (i *IN) Execute(z *Zog) error {
	// See spec comment in OUT
	var addr uint16
	if i.port == C {
		addr = z.reg.Read16(BC)
	} else {
		addr = uint16(z.reg.A) | (uint16(z.reg.A) << 8)
	}
	n := z.in(addr)
	return i.dst.Write8(z, n)
}

type PUSH struct {
	InstU16
}

func NewPUSH(l Loc16) *PUSH {
	return &PUSH{InstU16{l: l}}
}
func (p *PUSH) String() string {
	return fmt.Sprintf("PUSH %s", p.l)
}
func (p *PUSH) Encode() []byte {
	p.inspectRP2()
	if p.lInfo.ltype != tableRP2 {
		panic("Non-tableRP PUSH")
	}
	buf := []byte{encodeXPQZ(3, p.lInfo.idxTable, 0, 5)}
	return idxEncodeHelper(buf, p.idx)
}
func (p *PUSH) Execute(z *Zog) error {
	nn, err := p.l.Read16(z)
	if err != nil {
		return err
	}
	z.push(nn)
	return nil
}

type POP struct {
	InstU16
}

func NewPOP(l Loc16) *POP {
	return &POP{InstU16{l: l}}
}
func (p *POP) String() string {
	return fmt.Sprintf("POP %s", p.l)
}
func (p *POP) Encode() []byte {
	p.inspectRP2()
	if p.lInfo.ltype != tableRP2 {
		panic("Non-tableRP PUSH")
	}
	buf := []byte{encodeXPQZ(3, p.lInfo.idxTable, 0, 1)}
	return idxEncodeHelper(buf, p.idx)
}
func (p *POP) Execute(z *Zog) error {
	nn := z.pop()
	err := p.l.Write16(z, nn)
	if err != nil {
		return err
	}
	return nil
}

type RST struct {
	addr byte
}

func (r *RST) String() string {
	return fmt.Sprintf("RST %d", r.addr)
}
func (r *RST) Encode() []byte {
	y := r.addr / 8
	return []byte{encodeXYZ(3, y, 7)}
}
func (r *RST) Resolve(a *Assembly) error {
	return nil
}
func (r *RST) Execute(z *Zog) error {
	z.push(z.reg.PC)
	z.jp(uint16(r.addr))
	return nil
}

type RET struct {
	c Conditional
}

func (r *RET) String() string {
	if r.c == True || r.c == nil {
		return "RET"
	} else {
		return fmt.Sprintf("RET %s", r.c)
	}
}
func (r *RET) Encode() []byte {
	if r.c == True || r.c == nil {
		return []byte{encodeXPQZ(3, 0, 1, 1)}
	}
	y := findInTableCC(r.c)
	return []byte{encodeXYZ(3, y, 0)}
}
func (r *RET) Resolve(a *Assembly) error {
	return nil
}
func (r *RET) Execute(z *Zog) error {
	takeJump := r.c.IsTrue(z)
	if takeJump {
		addr := z.pop()
		z.jp(addr)
	}
	return nil
}

func NewAccum(name string, l Loc8) *accum {
	a := &accum{name: name, InstU8: InstU8{l: l}}
	a.f = findFuncInTableALU(name)
	return a
}

func isPos8(v byte) bool {
	return v&0x80 == 0
}
/*
func aluAdd(z *Zog, a, b byte) byte {
	v := a+b
	z.SetFlag(F_S, v & 0x80 != 0)
	z.SetFlag(F_Z, v == 0)
	z.SetFlag(F_H, ((a&0x0f)+(b&0x0f))&0x10 != 0)
	vSigned := int(int8(a)) + int(int8(b))
	z.SetFlag(F_PV, vSigned >= 0x80 || vSigned < -0x80)
	z.SetFlag(F_N, false)
	z.SetFlag(F_C, int(a) + int(b) > 0xff)
	return v
}
*/
func aluAdd(z *Zog, a, b byte) byte {
	return adcHelper(z, a, b, 0)
}

func adcHelper(z *Zog, a, b, c byte) byte {
	v := a+b+c
	z.SetFlag(F_S, !isPos8(v))
	z.SetFlag(F_Z, v == 0)
	z.SetFlag(F_H, ((a&0x0f)+(b&0x0f)+c)&0x10 != 0)
	vSigned := int(int8(a)) + int(int8(b)) + int(c)
	z.SetFlag(F_PV, vSigned >= 0x80 || vSigned < -0x80)
	z.SetFlag(F_N, false)
	z.SetFlag(F_C, int(a) + int(b) + int(c) > 0xff)
	return v
}
func aluAdc(z *Zog, a, b byte) byte {
	c := byte(0)
	if z.GetFlag(F_C) {
		c = 1
	}
	return adcHelper(z, a, b, c)
}
func aluSub(z *Zog, a, b byte) byte {
	return sbcHelper(z, a, b, 0)
}

func sbcHelper(z *Zog, a, b, c byte) byte {
	v := a - b - c
	z.SetFlag(F_S, !isPos8(v))
	z.SetFlag(F_Z, v == 0)
	z.SetFlag(F_H, ((a&0x0f) - (b&0x0f) - c) & 0x10 != 0)

	vSigned := int(int8(a)) - int(int8(b)) - int(c)
	z.SetFlag(F_PV, vSigned >= 0x80 || vSigned < -0x80)
	z.SetFlag(F_N, true)
	z.SetFlag(F_C, int(a) - int(b) - int(c) < 0)
	return v
}
func aluSbc(z *Zog, a, b byte) byte {
	c := byte(0)
	if z.GetFlag(F_C) {
		c = 1
	}
	return sbcHelper(z, a, b, c)
}

func aluAnd(z *Zog, a, b byte) byte {
	v := a & b
	z.SetFlag(F_S, !isPos8(v))
	z.SetFlag(F_Z, v == 0)
	z.SetFlag(F_H, true)
	setParity(z, v)
	z.SetFlag(F_N, false)
	z.SetFlag(F_C, false)
	return v
}
func aluXor(z *Zog, a, b byte) byte {
	v := a ^ b
	z.SetFlag(F_S, !isPos8(v))
	z.SetFlag(F_Z, v == 0)
	z.SetFlag(F_H, false)
	setParity(z, v)
	z.SetFlag(F_N, false)
	z.SetFlag(F_C, false)
	return v
}
func aluOr(z *Zog, a, b byte) byte {
	v := a | b
	z.SetFlag(F_S, !isPos8(v))
	z.SetFlag(F_Z, v == 0)
	z.SetFlag(F_H, false)
	setParity(z, v)
	z.SetFlag(F_N, false)
	z.SetFlag(F_C, false)
	return v
}
func aluCp(z *Zog, a, b byte) byte {
	// Note the calling code is special cased, we just do the sub here
	return aluSub(z, a, b)
}

type accumFunc func(z *Zog, a, b byte) byte
type accum struct {
	f accumFunc
	InstU8
	name string
}

func (a accum) String() string {
	switch a.name {
	case "ADD", "ADC", "SBC":
		return fmt.Sprintf("%s A, %s", a.name, a.l)
	default:
		return fmt.Sprintf("%s %s", a.name, a.l)
	}
}
func (a accum) Encode() []byte {
	a.inspect()
	y := findInTableALU(a.name)
	var buf []byte
	switch a.lInfo.ltype {
	case tableR:
		buf = []byte{encodeXYZ(2, y, a.lInfo.idxTable)}
	case Immediate:
		buf = []byte{encodeXYZ(3, y, 6)}
		buf = append(buf, a.lInfo.imm8)
	default:
		panic("Unknown accum location type")
	}
	return idxEncodeHelper(buf, a.idx)
}
func (a accum) Execute(z *Zog) error {
	regA := z.reg.A
	arg, err := a.l.Read8(z)
	if err != nil {
		return fmt.Errorf("Accum [%s] : can't read %s: %s", a.name, a.l, err)
	}

	v := a.f(z, regA, arg)

	// Hack - CP runs a SUB, but we don't save the value to accum here
	if strings.ToLower(a.name) != "cp" {
		z.reg.A = v
	}

	return nil
}

type rotFunc func(z *Zog, v byte) byte
type rot struct {
	InstU8
	cpy  Loc8
	name string
	f    rotFunc
}

func getCY(z *Zog) byte {
	cy := byte(0)
	if z.GetFlag(F_C) {
		cy = 1
	}
	return cy
}

func popcount(v byte) int {
	numBits := 0
	mask := byte(1 << 7)
	for mask > 0 {
		if v&mask != 0 {
			numBits++
		}
		mask = mask >> 1
	}
	return numBits
}

func setParity(z *Zog, v byte) {
	numBits := popcount(v)
	z.SetFlag(F_PV, numBits%2 == 0)
}

func rotRlc(z *Zog, v byte) byte {
	h := (v & 0x80) >> 7
	v = v << 1

	v = v | h

	z.SetFlag(F_C, h == 1)
	return v
}
func rotRrc(z *Zog, v byte) byte {
	l := v & 0x01
	v = v >> 1

	v = v | l<<7

	z.SetFlag(F_C, l == 1)
	return v
}
func rotRl(z *Zog, v byte) byte {
	h := (v & 0x80) >> 7
	v = v << 1

	v = v | getCY(z)
	z.SetFlag(F_C, h == 1)
	return v
}
func rotRr(z *Zog, v byte) byte {
	l := v & 0x01
	v = v >> 1

	v = v | (getCY(z) << 7)
	z.SetFlag(F_C, l == 1)
	return v
}
func rotSla(z *Zog, v byte) byte {
	h := (v & 0x80) >> 7
	v = v << 1

	z.SetFlag(F_C, h == 1)
	return v
}
func rotSra(z *Zog, v byte) byte {
	l := v & 0x01
	h := (v & 0x80) >> 7
	v = v >> 1

	z.SetFlag(F_C, l == 1)
	v = v | h
	return v
}
func rotSll(z *Zog, v byte) byte {
	h := (v & 0x80) >> 7
	v = v << 1
	v = v | 1

	z.SetFlag(F_C, h == 1)
	return v
}
func rotSrl(z *Zog, v byte) byte {
	l := v & 0x01
	v = v >> 1

	z.SetFlag(F_C, l == 1)
	return v
}

func NewRot(name string, l Loc8, cpy Loc8) *rot {
	r := &rot{InstU8: InstU8{l: l}, cpy: cpy, name: name}
	r.f = findFuncInTableROT(name)
	return r
}

func (r *rot) String() string {
	s := fmt.Sprintf("%s %s", r.name, r.l)
	if r.cpy != nil {
		s += ", " + r.cpy.String()
	}
	return s
}
func (r *rot) Encode() []byte {
	r.inspect()
	if r.lInfo.ltype != tableR {
		panic("Non-tableR src in BIT")
	}
	y := findInTableROT(r.name)
	z := r.lInfo.idxTable
	if r.idx.isPrefix && r.cpy != nil {
		z = findInTableR(r.cpy)
	}
	buf := []byte{0xcb, encodeXYZ(0, y, z)}
	return ddcbHelper(buf, r.idx)
}
func (r *rot) Execute(z *Zog) error {
	v, err := r.l.Read8(z)
	if err != nil {
		return fmt.Errorf("Rot [%s] : can't read [%s]: %s", r.name, r.l, err)
	}

	v = r.f(z, v)

	z.SetFlag(F_S, v >= 0x80)
	z.SetFlag(F_Z, v == 0)
	z.SetFlag(F_H, false)
	setParity(z, v)
	z.SetFlag(F_N, false)

	err = r.l.Write8(z, v)
	if err != nil {
		return fmt.Errorf("Rot [%s] : can't write [%s]: %s", r.name, r.l, err)
	}
	if r.cpy != nil {
		err = r.cpy.Write8(z, v)
		if err != nil {
			return fmt.Errorf("Rot [%s] : can't write copy [%s]: %s", r.name, r.cpy, err)
		}
	}
	return nil
}

type BIT struct {
	InstU8
	num byte
}

func NewBIT(num byte, l Loc8) *BIT {
	return &BIT{InstU8: InstU8{l: l}, num: num}
}
func (b *BIT) String() string {
	return fmt.Sprintf("BIT %d, %s", b.num, b.l)
}
func (b *BIT) Encode() []byte {
	b.inspect()
	if b.lInfo.ltype != tableR {
		panic("Non-tableR src in BIT")
	}
	z := b.lInfo.idxTable
	enc := encodeXYZ(1, b.num, z)
	return ddcbHelper([]byte{0xcb, enc}, b.idx)
}
func (b *BIT) Execute(z *Zog) error {
	v, err := b.l.Read8(z)
	if err != nil {
		return fmt.Errorf("BIT : can't read [%s]: %s", b.l, err)
	}
	v = v >> b.num
	bit := v & 1
	z.SetFlag(F_Z, bit == 0)
	z.SetFlag(F_N, false)
	z.SetFlag(F_H, true)
	return nil
}

type RES struct {
	InstU8
	cpy Loc8
	num byte
}

func NewRES(num byte, l Loc8, cpy Loc8) *RES {
	return &RES{InstU8: InstU8{l: l}, cpy: cpy, num: num}
}
func (r *RES) String() string {
	s := fmt.Sprintf("RES %d, %s", r.num, r.l)
	if r.cpy != nil {
		s += ", " + r.cpy.String()
	}
	return s
}
func (r *RES) Encode() []byte {
	r.inspect()
	if r.lInfo.ltype != tableR {
		panic("Non-tableR src in BIT")
	}
	z := r.lInfo.idxTable
	if r.idx.isPrefix && r.cpy != nil {
		z = findInTableR(r.cpy)
	}
	enc := encodeXYZ(2, r.num, z)
	return ddcbHelper([]byte{0xcb, enc}, r.idx)
}
func (r *RES) Execute(z *Zog) error {
	v, err := r.l.Read8(z)
	if err != nil {
		return fmt.Errorf("RES : can't read [%s]: %s", r.l, err)
	}
	andMask := byte(1) << r.num
	xorMask := v & andMask
	v = v ^ xorMask
	return r.l.Write8(z, v)
}

type SET struct {
	InstU8
	cpy Loc8
	num byte
}

func NewSET(num byte, l Loc8, cpy Loc8) *SET {
	return &SET{InstU8: InstU8{l: l}, cpy: cpy, num: num}
}
func (s *SET) String() string {
	str := fmt.Sprintf("SET %d, %s", s.num, s.l)
	if s.cpy != nil {
		str += ", " + s.cpy.String()
	}
	return str
}
func (s *SET) Encode() []byte {
	s.inspect()
	if s.lInfo.ltype != tableR {
		panic("Non-tableR src in SET")
	}
	z := s.lInfo.idxTable
	if s.idx.isPrefix && s.cpy != nil {
		z = findInTableR(s.cpy)
	}
	enc := encodeXYZ(3, s.num, z)
	return ddcbHelper([]byte{0xcb, enc}, s.idx)
}
func (s *SET) Execute(z *Zog) error {
	v, err := s.l.Read8(z)
	if err != nil {
		return fmt.Errorf("SET : can't read [%s]: %s", s.l, err)
	}
	mask := byte(1) << s.num
	v = v | mask
	return s.l.Write8(z, v)
}

type Simple byte

const (
	NOP Simple = 0x00

	HALT Simple = 0x76

	RLCA Simple = 0x07
	RRCA Simple = 0x0f
	RLA  Simple = 0x17
	RRA  Simple = 0x1f
	DAA  Simple = 0x27
	CPL  Simple = 0x2f
	SCF  Simple = 0x37
	CCF  Simple = 0x3f

	EXX Simple = 0xd9

	DI Simple = 0xf3
	EI Simple = 0xfb
)

type simpleName struct {
	inst Simple
	name string
}

var simpleNames []simpleName = []simpleName{
	{NOP, "NOP"},

	{HALT, "HALT"},

	{RLCA, "RLCA"},
	{RRCA, "RRCA"},
	{RLA, "RLA"},
	{RRA, "RRA"},
	{DAA, "DAA"},
	{CPL, "CPL"},
	{SCF, "SCF"},
	{CCF, "CCF"},

	{EXX, "EXX"},

	{DI, "DI"},
	{EI, "EI"},
}

func (s Simple) String() string {

	for _, simpleName := range simpleNames {
		if simpleName.inst == s {
			return simpleName.name
		}
	}
	panic(fmt.Sprintf("Unknown simple instruction: %02X", byte(s)))
}
func (s Simple) Encode() []byte {
	return []byte{byte(s)}
}
func (s Simple) Resolve(a *Assembly) error {
	return nil
}
func (s Simple) Execute(z *Zog) error {
	switch s {
	case NOP:
		return nil
	case HALT:
		return ErrHalted
	case RLCA:
		fReg := z.reg.F
		z.reg.A = rotRlc(z, z.reg.A)
		z.reg.F = fReg
		z.SetFlag(F_H, false)
		z.SetFlag(F_N, false)
		z.SetFlag(F_C, z.reg.A&0x01 != 0)
		return nil
	case RRCA:
		fReg := z.reg.F
		z.reg.A = rotRrc(z, z.reg.A)
		z.reg.F = fReg
		z.SetFlag(F_H, false)
		z.SetFlag(F_N, false)
		z.SetFlag(F_C, z.reg.A&0x80 != 0)
		return nil
	case RLA:
		fReg := z.reg.F
		initialA := z.reg.A
		z.reg.A = rotRl(z, initialA)
		z.reg.F = fReg
		z.SetFlag(F_H, false)
		z.SetFlag(F_N, false)
		z.SetFlag(F_C, initialA&0x80 != 0)
		return nil
	case RRA:
		fReg := z.reg.F
		initialA := z.reg.A
		z.reg.A = rotRr(z, initialA)
		z.reg.F = fReg
		z.SetFlag(F_H, false)
		z.SetFlag(F_N, false)
		z.SetFlag(F_C, initialA&0x01 != 0)
		return nil
	case DAA:

		// z80heaven - but that doesn't even look at the N flag :-/
/* ---
		v := z.reg.A
		lo := v & 0x0f
		if z.GetFlag(F_H) || lo > 9 {
			v += 0x06
		}
		hi := (v & 0xf0) >> 4
		if z.GetFlag(F_C) || hi > 9 {
			v += 0x60
			z.SetFlag(F_C, true)
		} else {
			z.SetFlag(F_C, false)
		}
		z.reg.A = v
		setParity(z, v)
	--- */

		// http://www.worldofspectrum.org/faq/reference/z80reference.htm#DAA
		v := z.reg.A
		c := z.GetFlag(F_C)
		h := z.GetFlag(F_H)
		n := z.GetFlag(F_N)
		newCarry := c
		correctionFactor  := byte(0x00)
		if v > 0x99 || c {
			correctionFactor |= 0x60
			newCarry = true
		} else {
			newCarry = false
		}
		if v & 0x0f > 9 || h {
			correctionFactor |= 0x06
		}
		if n {
			v -= correctionFactor
		} else {
			v += correctionFactor
		}

		// Bit 4 of: A(before) XOR A(after).
		newH := z.reg.A ^ v
		newH &= 0x80
		z.SetFlag(F_H, newH != 0)
		z.SetFlag(F_C, newCarry)
		setParity(z, v)
		z.SetFlag(F_S, !isPos8(v))
		z.SetFlag(F_Z, v == 0)

		z.reg.A = v
		return nil
	case CPL:
		z.reg.A = z.reg.A ^ 0xff
		z.SetFlag(F_H, true)
		z.SetFlag(F_N, true)
		return nil
	case SCF:
		z.SetFlag(F_H, false)
		z.SetFlag(F_N, false)
		z.SetFlag(F_C, true)
		return nil
	case CCF:
		c := z.GetFlag(F_C)
		h := z.GetFlag(F_H)
		z.SetFlag(F_C, !c)
		z.SetFlag(F_H, !h)
		z.SetFlag(F_N, false)
		return nil

	case EXX:
		z.reg.B, z.reg.B_PRIME = z.reg.B_PRIME, z.reg.B
		z.reg.C, z.reg.C_PRIME = z.reg.C_PRIME, z.reg.C
		z.reg.D, z.reg.D_PRIME = z.reg.D_PRIME, z.reg.D
		z.reg.E, z.reg.E_PRIME = z.reg.E_PRIME, z.reg.E
		z.reg.H, z.reg.H_PRIME = z.reg.H_PRIME, z.reg.H
		z.reg.L, z.reg.L_PRIME = z.reg.L_PRIME, z.reg.L
		return nil

	case DI:
		return z.di()
	case EI:
		return z.ei()
	default:
		return fmt.Errorf("Unknown simple instruction: %02X", byte(s))
	}
}

func LookupSimpleName(name string) Simple {
	name = strings.ToUpper(name)
	for _, simpleName := range simpleNames {
		if simpleName.name == name {
			return simpleName.inst
		}
	}
	panic(fmt.Errorf("Unrecognised Simple instruction name : [%s]", name))
}

type EDSimple byte

const (
	NEG  EDSimple = 0x44
	RETN EDSimple = 0x45
	RETI EDSimple = 0x4d

	RRD EDSimple = 0x67
	RLD EDSimple = 0x6f

	IM0 EDSimple = 0x46
	IM1 EDSimple = 0x56
	IM2 EDSimple = 0x5e

	LDI  EDSimple = 0xa0
	CPI  EDSimple = 0xa1
	LDD  EDSimple = 0xa8
	CPD  EDSimple = 0xa9
	LDIR EDSimple = 0xb0
	CPIR EDSimple = 0xb1
	LDDR EDSimple = 0xb8
	CPDR EDSimple = 0xb9

	INI  EDSimple = 0xa2
	OUTI EDSimple = 0xa3
	IND  EDSimple = 0xaa
	OUTD EDSimple = 0xab
	INIR EDSimple = 0xb2
	OTIR EDSimple = 0xb3
	INDR EDSimple = 0xba
	OTDR EDSimple = 0xbb
)

type edSimpleName struct {
	inst EDSimple
	name string
}

var EDSimpleNames []edSimpleName = []edSimpleName{
	{NEG, "NEG"},
	{RETN, "RETN"},
	{RETI, "RETI"},
	{RRD, "RRD"},
	{RLD, "RLD"},
	{IM0, "IM 0"},
	{IM1, "IM 1"},
	{IM2, "IM 2"},

	{LDI, "LDI"},
	{CPI, "CPI"},
	{LDD, "LDD"},
	{CPD, "CPD"},
	{LDIR, "LDIR"},
	{CPIR, "CPIR"},
	{LDDR, "LDDR"},
	{CPDR, "CPDR"},

	{INI, "INI"},
	{OUTI, "OUTI"},
	{IND, "IND"},
	{OUTD, "OUTD"},
	{INIR, "INIR"},
	{OTIR, "OTIR"},
	{INDR, "INDR"},
	{OTDR, "OTDR"},
}

func (s EDSimple) String() string {

	for _, simpleName := range EDSimpleNames {
		if simpleName.inst == s {
			return simpleName.name
		}
	}
	panic(fmt.Sprintf("Unknown EDSimple instruction: %02X", byte(s)))
}

func (s EDSimple) Encode() []byte {
	return []byte{0xed, byte(s)}
}
func (s EDSimple) Resolve(a *Assembly) error {
	return nil
}

func (s EDSimple) Execute(z *Zog) error {

	cpHelper := func(z *Zog, inc bool) error {
		bc := z.reg.Read16(BC)
		hl := z.reg.Read16(HL)

		a := z.reg.A
		b, err := z.mem.Peek(hl)
		if err != nil {
			return err
		}

		v := a - b

		bc--
		z.reg.Write16(BC, bc)
		if inc {
			hl++
		} else {
			hl--
		}
		z.reg.Write16(HL, hl)

		z.SetFlag(F_S, !isPos8(v))
		z.SetFlag(F_Z, v == 0)
		z.SetFlag(F_H, ((a&0x0f)-(b&0x0f))&0x10 != 0)
		z.SetFlag(F_PV, bc == 0)
		z.SetFlag(F_N, true)

		return nil
	}

	outHelper := func(z *Zog, inc bool) error {
		bc := z.reg.Read16(BC)
		hl := z.reg.Read16(HL)

		n, err := z.mem.Peek(hl)
		if err != nil {
			return err
		}
		z.out(bc, n)

		if inc {
			hl++
			z.SetFlag(F_N, false)
		} else {
			hl--
			z.SetFlag(F_N, true)
		}
		z.reg.Write16(HL, hl)
		z.reg.B--
		z.SetFlag(F_Z, z.reg.B == 0)
		return nil
	}

	inHelper := func(z *Zog, inc bool) error {
		bc := z.reg.Read16(BC)
		hl := z.reg.Read16(HL)
		n := z.in(bc)
		err := z.mem.Poke(hl, n)
		if err != nil {
			return err
		}
		if inc {
			hl++
			z.SetFlag(F_N, false)
		} else {
			hl--
			z.SetFlag(F_N, true)
		}
		z.reg.Write16(HL, hl)
		z.reg.B--
		z.SetFlag(F_Z, z.reg.B == 0)
		return nil
	}

	ldHelper := func(z *Zog, inc bool) error {
		bc := z.reg.Read16(BC)
		de := z.reg.Read16(DE)
		hl := z.reg.Read16(HL)

		n, err := z.mem.Peek(hl)
		if err != nil {
			return err
		}

		err = z.mem.Poke(de, n)
		if err != nil {
			return err
		}

		bc--
		z.reg.Write16(BC, bc)
		if inc {
			de++
			hl++
		} else {
			de--
			hl--
		}
		z.reg.Write16(DE, de)
		z.reg.Write16(HL, hl)

		z.SetFlag(F_PV, bc == 0)
		return nil
	}

	switch s {
	case NEG:
		z.reg.A = aluSub(z, 0, z.reg.A)
		z.SetFlag(F_N, true)
		return nil
	case RETN:
		z.iff1 = z.iff2
		addr := z.pop()
		z.jp(addr)
		return nil
	case RETI:
		addr := z.pop()
		z.jp(addr)
		return nil
	case RRD:
		// Rightwards version of RLD
		hl := z.reg.Read16(HL)
		a := z.reg.Read8(A)

		n, err := z.mem.Peek(hl)
		if err != nil {
			return err
		}

		// Three nybbles
		n3 := a & 0x0f
		n2 := n & 0xf0 >> 4
		n1 := n & 0x0f

		// Set low nibble of A to n1
		a = a & 0xf0
		a = a | n1
		z.reg.Write8(A, a)
		// Set (HL) to n3 n2
		n = n3<<4 | n2
		err = z.mem.Poke(hl, n)

		return nil
	case RLD:
		/*
			Performs a 4-bit leftward rotation of the 12-bit number whose 4 most
			signigifcant bits are the 4 least significant bits of A, and its 8 least
			significant bits are in (HL).
		*/
		hl := z.reg.Read16(HL)
		a := z.reg.Read8(A)

		n, err := z.mem.Peek(hl)
		if err != nil {
			return err
		}

		// Three nybbles
		n3 := a & 0x0f
		n2 := n & 0xf0 >> 4
		n1 := n & 0x0f

		// Set low nibble of A to n2
		a = a & 0xf0
		a = a | n2
		z.reg.Write8(A, a)
		// Set (HL) to n1 n3
		n = n1<<4 | n3
		err = z.mem.Poke(hl, n)

		return nil
	case IM0:
		z.im(0)
		return nil
	case IM1:
		z.im(1)
		return nil
	case IM2:
		z.im(2)
		return nil

	case LDI:
		return ldHelper(z, true)
	case CPI:
		return cpHelper(z, true)
	case LDD:
		return ldHelper(z, false)
	case CPD:
		return cpHelper(z, false)
	case LDIR:
		for {
			err := ldHelper(z, true)
			if err != nil {
				return err
			}
			if z.GetFlag(F_PV) {
				break
			}
		}
		return nil
	case CPIR:
		for {
			err := cpHelper(z, true)
			if err != nil {
				return err
			}
			if z.GetFlag(F_PV) {
				break
			}
		}
		return nil
	case LDDR:
		for {
			err := ldHelper(z, true)
			if err != nil {
				return err
			}
			if z.GetFlag(F_PV) {
				break
			}
		}
		return nil
	case CPDR:
		for {
			err := cpHelper(z, false)
			if err != nil {
				return err
			}
			if z.GetFlag(F_PV) {
				break
			}
		}
		return nil

	case INI:
		return inHelper(z, true)
	case OUTI:
		return outHelper(z, true)
	case IND:
		return inHelper(z, false)
	case OUTD:
		return outHelper(z, false)
	case INIR:
		for {
			err := inHelper(z, true)
			if err != nil {
				return err
			}
			if z.GetFlag(F_Z) {
				break
			}
		}
		return nil
	case OTIR:
		for {
			err := outHelper(z, true)
			if err != nil {
				return err
			}
			if z.GetFlag(F_Z) {
				break
			}
		}
		return nil
	case INDR:
		for {
			err := inHelper(z, false)
			if err != nil {
				return err
			}
			if z.GetFlag(F_Z) {
				break
			}
		}
		return nil
	case OTDR:
		for {
			err := outHelper(z, false)
			if err != nil {
				return err
			}
			if z.GetFlag(F_Z) {
				break
			}
		}
		return nil
	default:
		return fmt.Errorf("Unknown EDSimple instruction: %02X", byte(s))
	}
}

func LookupEDSimpleName(name string) EDSimple {
	name = strings.ToUpper(name)
	for _, simpleName := range EDSimpleNames {
		if simpleName.name == name {
			return simpleName.inst
		}
	}
	panic(fmt.Errorf("Unrecognised EDSimple instruction name : [%s]", name))
}
