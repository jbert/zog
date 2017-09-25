package speccy

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		keysdown          []sdl.Keycode
		addrHi            byte
		expectedInputByte byte
	}{
		{[]sdl.Keycode{sdl.K_a}, 0xfd, 0xfe},
		{[]sdl.Keycode{sdl.K_a}, 0x00, 0xfe},
		{[]sdl.Keycode{sdl.K_a}, 0xff, 0xff},

		{[]sdl.Keycode{sdl.K_1}, 0xf7, 0xfe},
		{[]sdl.Keycode{sdl.K_1}, 0x00, 0xfe},
		{[]sdl.Keycode{sdl.K_1}, 0xff, 0xff},

		{[]sdl.Keycode{sdl.K_2}, 0xf7, 0xfd},
		{[]sdl.Keycode{sdl.K_2}, 0x00, 0xfd},
		{[]sdl.Keycode{sdl.K_2}, 0xff, 0xff},

		{[]sdl.Keycode{sdl.K_5}, 0xf7, 0xef},
		{[]sdl.Keycode{sdl.K_5}, 0x00, 0xef},
		{[]sdl.Keycode{sdl.K_5}, 0xff, 0xff},

		{[]sdl.Keycode{sdl.K_a, sdl.K_4}, 0xfd, 0xfe},
		{[]sdl.Keycode{sdl.K_a, sdl.K_4}, 0xf7, 0xf7},
		{[]sdl.Keycode{sdl.K_a, sdl.K_4}, 0x00, 0xf6},
		{[]sdl.Keycode{sdl.K_a, sdl.K_4}, 0xff, 0xff},
	}

	for _, tc := range testCases {
		got := calcInputByte(tc.addrHi, tc.keysdown)
		expected := tc.expectedInputByte
		if got != expected {
			t.Errorf("Fail: got %02X expected %02X [%v]", got, expected, tc)
		}
	}
}
