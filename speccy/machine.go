package speccy

import (
	"os"
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
	tape   *Tape

	done chan struct{}
}

func (m *Machine) portInputHandler(addr uint16) byte {
	if addr&1 != 0 {
		return 0
	}
	keysdown := m.keys.keysdown()

	keyboardBytes := calcInputByte(byte(addr>>8), keysdown)
	keyboardBytes &= 0b00011111

	ear_byte := m.tape.tapeEarByte()
	ear_byte &= 0b01000000

	return keyboardBytes | ear_byte

}

func (m *Machine) portOutputHandler(addr uint16, b byte) {

	if addr&1 != 0 {
		return
	}

	if m.screen != nil {
		m.screen.SetBorderCol(b)
	} else {
		fmt.Fprintf(os.Stderr, "Screen not initialised, border colour not changed to %03b\n", b&0b111)
	}
}

func NewMachine(z *zog.Zog, tape_file *os.File) *Machine {
	return &Machine{
		keys: NewKeyboardState(),
		z:    z,
		tape: NewTape(tape_file),
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
	m.z.RegisterInputHandler(m.portInputHandler)
	m.z.RegisterOutputHandler(m.portOutputHandler)
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

	go TapeReadManager(m.tape)

	return nil
}

func (m *Machine) Stop() {
	close(m.done)
}

const romFileName = "/usr/share/spectrum-roms/48.rom"

func (m *Machine) loadROMs() error {
	return m.z.LoadROMFile(0x0000, romFileName)
}
