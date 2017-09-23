package speccy

import (
	"fmt"
	"time"

	"github.com/jbert/zog"
)

type Machine struct {
	keys   *keyboardState
	screen *Screen
	z      *zog.Zog

	done chan struct{}
}

func NewMachine(z *zog.Zog) *Machine {
	screen, err := NewScreen(z.Mem)
	if err != nil {
		panic(fmt.Sprintf("Can't create screen: %s", err))
	}

	return &Machine{
		keys:   NewKeyboardState(),
		screen: screen,
		z:      z,

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
	m.keys.InstallKeyboardInputPorts(m.z)
	every := time.Second / 50
	go func() {
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
