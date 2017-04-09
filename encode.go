package zog

import "fmt"

type InstBin8 struct {
	dst Dst8
	src Src8
}

type tableType int

const (
	tableR tableType = 1
)

type InstU8 struct {
	l Loc8

	base byte

	eTable   tableType
	idxTable byte

	isPrefix bool
	isIY     bool // If idxPrefix is true, else IX
	hasDisp  bool // Redundant?
	idxDisp  byte
}

func (u *InstU8) inspect() {

	iContents, ok := u.l.(IndexedContents)
	if ok {
		r16, ok := iContents.addr.(R16)
		if !ok {
			panic("Non-r16 addr in indexed content")
		}

		u.eTable = tableR
		u.idxTable = findInTableR(Contents{HL})

		u.isPrefix = true
		u.isIY = r16 == IY // Else IX

		u.hasDisp = true
		u.idxDisp = byte(iContents.d)
		return
	}

	contents, ok := u.l.(Contents)
	if ok {
		if contents.addr == HL {
			u.eTable = tableR
			u.idxTable = findInTableR(Contents{HL})
			return
		} else {
			panic("Non-HL contents in R8")
		}
	}

	r8, ok := u.l.(R8)
	if ok {
		u.eTable = tableR
		u.idxTable = findInTableR(r8)
		return
	}

	panic(fmt.Sprintf("WTF? %T", u.l))
}

func (u *InstU8) encodeHelper(base byte) []byte {
	encoded := []byte{base}
	if u.isPrefix {
		idxPrefix := byte(0xdd)
		if u.isIY {
			idxPrefix = 0xfd
		}
		encoded = append([]byte{idxPrefix}, encoded...)

		if u.hasDisp {
			encoded = append(encoded, u.idxDisp)
		}
	}

	return encoded
}

// See decomposeByte in decode.go
func encodeXYZ(x, y, z byte) byte {
	return x<<6 | y<<3 | z
}
