package speccy

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

// An entry in the map means the key is depressed. (poor key)
type keyboardState struct {
	sync.Mutex
	keysDown map[sdl.Keycode]struct{}
}

func NewKeyboardState() *keyboardState {
	ks := keyboardState{keysDown: make(map[sdl.Keycode]struct{})}
	return &ks
}

// 8 rows - index in array is bit position in addrHi 0->7
// 5 keys = bit0 1st, bit 4last
var keymapForBit = [][]sdl.Keycode{
	[]sdl.Keycode{sdl.K_LSHIFT, sdl.K_z, sdl.K_x, sdl.K_c, sdl.K_v},
	[]sdl.Keycode{sdl.K_a, sdl.K_s, sdl.K_d, sdl.K_f, sdl.K_g},
	[]sdl.Keycode{sdl.K_q, sdl.K_w, sdl.K_e, sdl.K_r, sdl.K_t},
	[]sdl.Keycode{sdl.K_1, sdl.K_2, sdl.K_3, sdl.K_4, sdl.K_5},
	[]sdl.Keycode{sdl.K_0, sdl.K_9, sdl.K_8, sdl.K_7, sdl.K_6},
	[]sdl.Keycode{sdl.K_p, sdl.K_o, sdl.K_i, sdl.K_u, sdl.K_y},
	[]sdl.Keycode{sdl.K_RETURN, sdl.K_l, sdl.K_k, sdl.K_j, sdl.K_h},
	[]sdl.Keycode{sdl.K_SPACE, sdl.K_RSHIFT, sdl.K_m, sdl.K_n, sdl.K_b},
}

func calcInputByte(addrHi byte, keysdown []sdl.Keycode) byte {
	n := byte(0xff)
	addrMask := byte(0x01)
	for _, keyRow := range keymapForBit {
		if addrMask&addrHi == 0 {
			// Bit is low in addrHi - process
			byteMask := byte(0xfe)
			for _, rowKey := range keyRow {
			KEYDOWN:
				for _, keydown := range keysdown {
					if keydown == rowKey {
						n &= byteMask
						break KEYDOWN
					}
				}
				byteMask <<= 1
				byteMask |= 0x01
			}
		}
		addrMask <<= 1
	}
	return n
}

func (ks *keyboardState) keyboardInputHandler(addr uint16) byte {
	hi := byte(addr >> 8)
	lo := byte(addr)

	if lo != 0xfe {
		return 0x00
	}

	keysdown := ks.keysdown()

	return calcInputByte(hi, keysdown)
}

func (ks *keyboardState) keysdown() []sdl.Keycode {
	ks.Lock()
	defer ks.Unlock()

	// Update our view of keys before the interrupt handler runs
	ks.update()

	var keys []sdl.Keycode
	for k := range ks.keysDown {
		keys = append(keys, k)
	}
	return keys
}

func (ks *keyboardState) update() {

	// Drain events to update map
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.KeyboardEvent:
			kc := ev.Keysym.Sym
			ks.keymove(kc, ev.Type == sdl.KEYUP)
		}
	}
}

func (ks *keyboardState) keymove(kc sdl.Keycode, up bool) {
	mappedKeys := []sdl.Keycode{}
	switch kc {
	case sdl.K_LEFT:
		mappedKeys = append(mappedKeys, sdl.K_LSHIFT)
		mappedKeys = append(mappedKeys, sdl.K_5)
	case sdl.K_DOWN:
		mappedKeys = append(mappedKeys, sdl.K_LSHIFT)
		mappedKeys = append(mappedKeys, sdl.K_6)
	case sdl.K_UP:
		mappedKeys = append(mappedKeys, sdl.K_LSHIFT)
		mappedKeys = append(mappedKeys, sdl.K_7)
	case sdl.K_RIGHT:
		mappedKeys = append(mappedKeys, sdl.K_LSHIFT)
		mappedKeys = append(mappedKeys, sdl.K_8)
	case sdl.K_ESCAPE:
		mappedKeys = append(mappedKeys, sdl.K_LSHIFT)
		mappedKeys = append(mappedKeys, sdl.K_1)
	case sdl.K_BACKSPACE:
		// Delete is shift-0
		mappedKeys = append(mappedKeys, sdl.K_LSHIFT)
		mappedKeys = append(mappedKeys, sdl.K_0)
	default:
		mappedKeys = append(mappedKeys, kc)
	}

	for _, mappedKey := range mappedKeys {
		if up {
			delete(ks.keysDown, mappedKey)
		} else {
			ks.keysDown[mappedKey] = struct{}{}
		}
	}
}
