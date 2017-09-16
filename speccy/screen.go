package speccy

import (
	"fmt"

	"github.com/jbert/zog"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 256
	screenHeight = 192
	screenScale  = 4

	screenMemStart = 0x4000
)

type Screen struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	mem      *zog.Memory
}

func NewScreen(mem *zog.Memory) (*Screen, error) {
	winTitle := "Speccy"
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth*screenScale, screenHeight*screenScale, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, fmt.Errorf("Failed to create window: %s\n", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, fmt.Errorf("Failed to create renderer: %s\n", err)
	}
	renderer.Clear()

	return &Screen{
		window:   window,
		renderer: renderer,
		mem:      mem,
	}, nil
}

func (s *Screen) Draw() {

	// Clear screen
	rect := sdl.Rect{0, 0, int32(screenWidth * screenScale), int32(screenHeight * screenScale)}
	s.renderer.SetDrawColor(255, 255, 255, 255)
	s.renderer.FillRect(&rect)

	// Draw each scanline
	for y := 0; y < screenHeight; y++ {
		s.drawScanline(y)
	}
	s.renderer.Present()
}

func (s *Screen) drawScanline(y int) {
	lineMemLen := screenWidth / 8

	sector := (y & 0xc0) >> 6
	sectorRow := (y & 0x38) >> 3
	charRow := (y & 0x07)
	addr := screenMemStart + (sector*64+charRow*8+sectorRow)*lineMemLen

	lineMem, err := s.mem.PeekBuf(uint16(addr), lineMemLen)
	if err != nil {
		panic(fmt.Errorf("Can't read screen memory at [%04X] len (%04X)", addr, lineMemLen))
	}
	for i, b := range lineMem {
		for bit := 0; bit < 8; bit++ {
			x := i*8 + bit
			if b&0x80 != 0 {
				s.renderer.SetDrawColor(0, 0, 0, 255)
			} else {
				s.renderer.SetDrawColor(255, 255, 255, 255)
			}
			//			fmt.Printf("x %d y %d i %d bit %d\n", x, y, i, b)
			rect := sdl.Rect{int32(x * screenScale), int32(y * screenScale), screenScale, screenScale}
			s.renderer.FillRect(&rect)
			//			s.renderer.DrawPoint(x, y)
			b <<= 1
		}
	}
}

func (s *Screen) Close() {
	s.renderer.Destroy()
	s.window.Destroy()
}