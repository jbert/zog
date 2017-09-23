package file

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestZ80Decompress(t *testing.T) {
	testCases := []struct {
		in, expected []byte
		err          error
	}{
		{[]byte{0xff}, []byte{}, errors.New("Missing end-of-block - len 1")},
		{[]byte{0xff, 0xff}, []byte{}, errors.New("Missing end-of-block - len 2")},
		{[]byte{0xff, 0xff, 0xff}, []byte{}, errors.New("Missing end-of-block - len 3")},
		{[]byte{0xff, 0xff, 0xff, 0xff}, []byte{}, errors.New("Missing end-of-block")},
		{[]byte{0x00, 0xed, 0xed, 0x00}, []byte{}, nil},
		{[]byte{0x00, 0x00, 0xed, 0xed, 0x00}, []byte{0x00}, nil},

		{[]byte{0xed, 0xed, 0x02, 0x03, 0x00, 0xed, 0xed, 0x00}, []byte{0x03, 0x03}, nil},
		{[]byte{0xed, 0x00, 0xed, 0xed, 0x02, 0x03, 0x00, 0xed, 0xed, 0x00}, []byte{0xed, 0x00, 0x03, 0x03}, nil},
	}

	for _, tc := range testCases {
		got, err := DecompressMem(tc.in)
		if err == nil && tc.err != nil {
			t.Errorf("Got no error - expected %s", tc.err)
			continue
		}
		if err != nil && tc.err == nil {
			t.Errorf("Got error %s - expected none", err)
			continue
		}
		if err != nil {
			if !strings.HasPrefix(err.Error(), tc.err.Error()) {
				t.Errorf("Got error [%s] - expected [%s]", err, tc.err)
			}
			continue
		}
		if !bytes.Equal(got, tc.expected) {
			t.Errorf("Decompress failed: got %v != %v", got, tc.expected)
			continue
		}
	}
}
