package zog

import "fmt"

func Encode(insts []Instruction) []byte {
	buf := make([]byte, 0)
	for _, inst := range insts {
		instBuf := inst.Encode()
		buf = append(buf, instBuf...)
	}
	return buf
}

type loc8Info struct {
	ltype    locType
	idxTable byte
	imm8     byte
	isBC     bool
	imm16    []byte
}

type idxInfo struct {
	isPrefix bool
	isIY     bool // If idxPrefix is true, else IX
	hasDisp  bool // Redundant?
	idxDisp  byte
}

type InstBin8 struct {
	dst Loc8
	src Loc8

	dstInfo loc8Info
	srcInfo loc8Info
	idx     idxInfo

	base byte
}

type locType int

const (
	tableR            locType = 1
	Immediate                 = 2
	tableRP                   = 3
	tableRP2                  = 4
	BCDEContents              = 5
	ImmediateContents         = 6
)

type InstU8 struct {
	l Loc8

	base byte

	lInfo loc8Info

	idx idxInfo
}

func inspectLoc8(l Loc8, info *loc8Info, idx *idxInfo) {
	// indexed H and L are tableR, and set idx info
	switch l {
	case IXH:
		info.ltype = tableR
		info.idxTable = findInTableR(H)
		idx.isPrefix = true
		idx.isIY = false
		return
	case IXL:
		info.ltype = tableR
		info.idxTable = findInTableR(L)
		idx.isPrefix = true
		idx.isIY = false
		return
	case IYH:
		info.ltype = tableR
		info.idxTable = findInTableR(H)
		idx.isPrefix = true
		idx.isIY = true
		return
	case IYL:
		info.ltype = tableR
		info.idxTable = findInTableR(L)
		idx.isPrefix = true
		idx.isIY = true
		return
	}

	iContents, ok := l.(IndexedContents)
	if ok {
		r16, ok := iContents.addr.(R16)
		if !ok {
			panic("Non-r16 addr in indexed content")
		}

		info.ltype = tableR
		info.idxTable = findInTableR(Contents{HL})

		idx.isPrefix = true
		idx.isIY = r16 == IY // Else IX

		idx.hasDisp = true
		idx.idxDisp = byte(iContents.d)
		return
	}

	contents, ok := l.(Contents)
	if ok {
		if contents.addr == HL {
			info.ltype = tableR
			info.idxTable = findInTableR(Contents{HL})
			return
		} else if contents.addr == BC || contents.addr == DE {
			info.ltype = BCDEContents
			info.isBC = contents.addr == BC
			return
		}
		imm16, ok := contents.addr.(Imm16)
		if ok {
			info.ltype = ImmediateContents
			hi := byte(imm16 >> 8)
			lo := byte(imm16 & 0xff)
			info.imm16 = []byte{lo, hi}
			return
		}
		panic("Unrecognised contents of loc8")
	}

	r8, ok := l.(R8)
	if ok {
		info.ltype = tableR
		info.idxTable = findInTableR(r8)
		return
	}

	imm8, ok := l.(Imm8)
	if ok {
		info.ltype = Immediate
		info.imm8 = byte(imm8)
		return
	}

	panic(fmt.Sprintf("WTF? %T", l))
}

func (u *InstU8) inspect() {
	inspectLoc8(u.l, &u.lInfo, &u.idx)
}

type InstU16 struct {
	l     Loc16
	lInfo loc16Info
	idx   idxInfo
}

type InstBin16 struct {
	dst Loc16
	src Loc16

	dstInfo loc16Info
	srcInfo loc16Info

	idx idxInfo
}

type loc16Info struct {
	ltype    locType
	idxTable byte
	imm16    []byte
}

func (li *loc16Info) isHLLike() bool {
	return li.ltype == tableRP && li.idxTable == HL_RP_INDEX
}

func inspectLoc16(l Loc16, info *loc16Info, idx *idxInfo, wantRP2 bool) {
	if l == IX {
		idx.isPrefix = true
		l = HL
	} else if l == IY {
		idx.isPrefix = true
		idx.isIY = true
		l = HL
	}

	contents, ok := l.(Contents)
	if ok {
		imm16, isImm := contents.addr.(Imm16)
		if isImm {
			info.ltype = ImmediateContents
			hi := byte(imm16 >> 8)
			lo := byte(imm16 & 0xff)
			info.imm16 = []byte{lo, hi}
			return
		}
		panic("Non-immediate Loc16 contents")
	}

	imm16, ok := l.(Imm16)
	if !ok {
		var label *Label
		// overwrite 'ok'
		label, ok = l.(*Label)
		if ok {
			imm16 = label.Imm16
		}
	}

	if ok {
		info.ltype = Immediate
		hi := byte(imm16 >> 8)
		lo := byte(imm16 & 0xff)
		info.imm16 = []byte{lo, hi}
	} else {
		if wantRP2 {
			info.ltype = tableRP2
			info.idxTable = findInTableRP2(l)
		} else {
			info.ltype = tableRP
			info.idxTable = findInTableRP(l)
		}
	}
}

func (u *InstU16) inspectRP2() {
	inspectLoc16(u.l, &u.lInfo, &u.idx, true)
}

func (u *InstU16) inspect() {
	inspectLoc16(u.l, &u.lInfo, &u.idx, false)
}

func (b *InstBin16) inspect() {
	inspectLoc16(b.dst, &b.dstInfo, &b.idx, false)
	inspectLoc16(b.src, &b.srcInfo, &b.idx, false)
}
func encodeHelper(base []byte, idx idxInfo, dispFirst bool) []byte {
	encoded := base
	if idx.isPrefix {
		idxPrefix := byte(0xdd)
		if idx.isIY {
			idxPrefix = 0xfd
		}
		if dispFirst {
			// dispFirst implies idx.hasDisp
			encoded = append([]byte{idxPrefix}, encoded...)
			n := encoded[len(encoded)-1]
			encoded[len(encoded)-1] = idx.idxDisp
			encoded = append(encoded, n)
		} else {
			encoded = append([]byte{idxPrefix}, encoded...)
			if idx.hasDisp {
				encoded = append(encoded, idx.idxDisp)
			}
		}
	}

	return encoded
}

func idxEncodeHelper(base []byte, idx idxInfo) []byte {
	return encodeHelper(base, idx, false)
}

func ddcbHelper(base []byte, idx idxInfo) []byte {
	return encodeHelper(base, idx, true)
}

func (b *InstBin8) inspect() {
	inspectLoc8(b.src, &b.srcInfo, &b.idx)
	inspectLoc8(b.dst, &b.dstInfo, &b.idx)
}

// See decomposeByte in decode.go
func encodeXYZ(x, y, z byte) byte {
	return x<<6 | y<<3 | z
}

func encodeXPQZ(x, p, q, z byte) byte {
	y := p<<1 | q
	return encodeXYZ(x, y, z)
}
