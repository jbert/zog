package speccy

import (
	"time"

	"fmt"
	"runtime"

	"github.com/jbert/zog"
	"github.com/veandco/go-sdl2/sdl"
)

type Machine struct {
	keys   *keyboardState
	screen *Screen
	z      *zog.Zog

	done chan struct{}
}

func NewMachine(z *zog.Zog) *Machine {

	return &Machine{
		keys: NewKeyboardState(),
		z:    z,
		done: make(chan struct{}),
	}
}

func (m Machine) LoadAddr() uint16 {
	return 0x8000
}

func (m Machine) RunAddr() uint16 {
	return 0x0000
}

func (m Machine) Name() string {
	return "speccy"
}

func (m *Machine) Start() error {
	err := m.loadROMs()
	if err != nil {
		return err
	}
	m.z.RegisterInputHandler(func(addr uint16) byte { return m.keys.keyboardInputHandler(addr) })
	every := time.Second / 50

	go func() {
		// See https://wiki.libsdl.org/SDL2/SDL_RenderPresent for why this is necessary
		// Technically this should be in the main thread, but it seems to work as long as
		// all SDL calls run in the same thread
		runtime.LockOSThread()
		sdl.Init(sdl.INIT_EVERYTHING)
		m.screen, err = NewScreen(m.z.Mem)
		if err != nil {
			panic(fmt.Sprintf("Can't create screen: %s", err))
		}
		tick := time.Tick(every)
		for {
			select {
			case <-m.done:
				break
			case <-tick:
				m.screen.Draw()
				m.z.DoInterrupt()
			}
		}
	}()

	return nil
}

func (m *Machine) Stop() {
	close(m.done)
}

const romFileName = "/usr/share/spectrum-roms/48.rom"

func (m *Machine) loadROMs() error {
	return m.z.LoadROMFile(0x0000, romFileName)
}
