package zog

import "fmt"

type loc8Info struct {
	ltype    locType
	idxTable byte
	imm8     byte
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
	tableR    locType = 1
	Immediate         = 2
)

type InstU8 struct {
	l Loc8

	base byte

	lInfo loc8Info

	idx idxInfo
}

func inspectLoc8(l Loc8, info *loc8Info, idx *idxInfo) {
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
		} else {
			panic("Non-HL contents in R8")
		}
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
	l   Loc16
	idx idxInfo
}

func lookupLoc16(l Loc16, idx *idxInfo) byte {
	if l == IX {
		idx.isPrefix = true
		l = HL
	} else if l == IY {
		idx.isPrefix = true
		idx.isIY = true
		l = HL
	}
	return findInTableRP(l)
}

func idxEncodeHelper(base []byte, idx idxInfo) []byte {
	encoded := base
	if idx.isPrefix {
		idxPrefix := byte(0xdd)
		if idx.isIY {
			idxPrefix = 0xfd
		}
		encoded = append([]byte{idxPrefix}, encoded...)

		if idx.hasDisp {
			encoded = append(encoded, idx.idxDisp)
		}
	}

	return encoded
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
