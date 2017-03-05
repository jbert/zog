package zog

import "testing"

func TestAccum(t *testing.T) {
	// EF == expected flag
	type EF struct {
		f     flag
		value bool
	}
	type testCase struct {
		assembly  string
		expectedA byte
		flags     []EF
	}

	testCases := []testCase{
		{"LD A, 10",
			10, nil},
		{"LD A, 0x10; LD B, 0x20; ADD A, B",
			0x30, []EF{{F_C, false}}},
		{"LD A, 0xFF; LD B, 0x02; ADD A, B",
			0x01, []EF{{F_C, false}}},
	}

	memSize := uint16(1024)
	for _, tc := range testCases {
		z := New(memSize)
		a, err := z.Execute(0, tc.assembly+"; HALT")
		if err != nil {
			t.Fatalf("Can't execute test prog [%s]: %s", tc.assembly, err)
		}
		if a != tc.expectedA {
			t.Fatalf("Wrong accum for [%s]: %d != %d", tc.assembly, a, tc.expectedA)
		}
		for _, ef := range tc.flags {
			fv := z.GetFlag(ef.f)
			if fv != ef.value {
				t.Fatalf("Wrong flag for [%s]: flag %s != %v", tc.assembly, ef.f, ef.value)
			}
		}
	}
}
