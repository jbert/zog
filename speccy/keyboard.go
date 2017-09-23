package speccy

import (
	"sync"

	"github.com/jbert/zog"
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

func (ks *keyboardState) InstallKeyboardInputPorts(z *zog.Zog) {
	keymaps := []struct {
		port uint16
		keys []sdl.Keycode
	}{
		{0xfefe, []sdl.Keycode{sdl.K_LSHIFT, sdl.K_z, sdl.K_x, sdl.K_c, sdl.K_v}},
		{0xfdfe, []sdl.Keycode{sdl.K_a, sdl.K_s, sdl.K_d, sdl.K_f, sdl.K_g}},
		{0xfbfe, []sdl.Keycode{sdl.K_q, sdl.K_w, sdl.K_e, sdl.K_r, sdl.K_t}},
		{0xf7fe, []sdl.Keycode{sdl.K_1, sdl.K_2, sdl.K_3, sdl.K_4, sdl.K_5}},
		{0xeffe, []sdl.Keycode{sdl.K_0, sdl.K_9, sdl.K_8, sdl.K_7, sdl.K_6}},
		{0xdffe, []sdl.Keycode{sdl.K_p, sdl.K_o, sdl.K_i, sdl.K_u, sdl.K_y}},
		{0xbffe, []sdl.Keycode{sdl.K_RETURN, sdl.K_l, sdl.K_k, sdl.K_j, sdl.K_h}},
		{0x7ffe, []sdl.Keycode{sdl.K_SPACE, sdl.K_RSHIFT, sdl.K_m, sdl.K_n, sdl.K_b}},
	}

	for _, keymapReused := range keymaps {
		keymap := keymapReused // This bit of go really sucks
		err := z.RegisterInputHandler(keymap.port, func() byte {
			return ks.inputHandler(keymap.keys)
		})
		if err != nil {
			panic(err)
		}
	}
}

// 5 keys = bit0 1st, bit 4last
func (ks *keyboardState) inputHandler(keys []sdl.Keycode) byte {
	ks.Lock()
	defer ks.Unlock()

	// Update our view of keys before the interrupt handler runs
	ks.update()

	// We have 5 bits. Key pressed is 0, else 1
	//	for k, _ := range *ks {
	//		fmt.Printf("JB - [%d] down\n", int(k))
	//	}
	n := byte(0xff)
	mask := byte(0xfe)
	for _, key := range keys {
		//		fmt.Printf("JB - checking key [%d]\n", int(key))
		_, ok := ks.keysDown[key]
		if ok {
			// Key pressed, clear bit
			n &= mask
		}
		mask <<= 1
		mask |= 0x01
		//		} else {
		//			fmt.Printf("JB - flagging key [%d] pressed\n", int(key))
		//		}
	}
	//	fmt.Printf("JB - ret [%02X]\n", n)
	return n
}

func (ks *keyboardState) update() {

	// Drain events to update map
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.KeyDownEvent:
			kc := ev.Keysym.Sym
			//					s := sdl.GetScancodeName(sc)
			ks.keymove(kc, false)
			//(*ks)[sc] = struct{}{}
		case *sdl.KeyUpEvent:
			kc := ev.Keysym.Sym
			ks.keymove(kc, true)
			//					s := sdl.GetScancodeName(sc)
			//delete(*ks, kc)
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
