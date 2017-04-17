package zog

import "fmt"

type loc8Info struct {
	eTable   tableType
	idxTable byte
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

type tableType int

const (
	tableR tableType = 1
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

		info.eTable = tableR
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
			info.eTable = tableR
			info.idxTable = findInTableR(Contents{HL})
			return
		} else {
			panic("Non-HL contents in R8")
		}
	}

	r8, ok := l.(R8)
	if ok {
		info.eTable = tableR
		info.idxTable = findInTableR(r8)
		return
	}

	panic(fmt.Sprintf("WTF? %T", l))
}

func (u *InstU8) inspect() {
	inspectLoc8(u.l, &u.lInfo, &u.idx)
}

func idxEncodeHelper(base byte, idx idxInfo) []byte {
	encoded := []byte{base}
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
