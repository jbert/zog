package file

import (
	"io"

	"github.com/jbert/zog"
)

type ImageFormat interface {
	Parse(r io.Reader) error
	Load(r io.Reader, z *zog.Zog) error
}
