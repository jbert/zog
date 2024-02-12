package speccy

import (
	"fmt"
	"time"

	"github.com/jbert/zog"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	borderWidth       = 48
	borderTop         = 56
	borderBelow       = 56
	screenWidth       = 256
	screenHeight      = 192
	totalScreenWidth  = borderWidth*2 + screenWidth
	totalScreenHeight = borderTop + borderBelow + screenHeight
	screenScale       = 5

	scanlineCycles = 224

	screenMemStart = 0x4000
	colourMemStart = 0x5800
)

type Screen struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	mem      *zog.Memory

	borderCol sdl.Color

	flashCount int
}

func NewScreen(mem *zog.Memory) (*Screen, error) {
	winTitle := "Speccy"
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		totalScreenWidth*screenScale, totalScreenHeight*screenScale, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, fmt.Errorf("Failed to create window: %s\n", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, fmt.Errorf("Failed to create renderer: %s\n", err)
	}
	renderer.Clear()

	return &Screen{
		window:    window,
		renderer:  renderer,
		mem:       mem,
		borderCol: sdl.Color{R: 0, G: 0, B: 0, A: 255},
	}, nil
}

const border_brightness = 0xd7

func (s *Screen) SetBorderCol(border uint8) {

	border &= 0b111

	borderG := border&0b100 != 0
	borderB := border&0b001 != 0
	borderR := border&0b010 != 0

	if borderG {
		s.borderCol.G = border_brightness
	} else {
		s.borderCol.G = 0
	}
	if borderB {
		s.borderCol.B = border_brightness
	} else {
		s.borderCol.B = 0
	}
	if borderR {
		s.borderCol.R = border_brightness
	} else {
		s.borderCol.R = 0
	}
}

func (s *Screen) Draw() {

	for y := 0; y < borderTop; y++ {
		borderSave := s.borderCol
		before := time.Now()

		s.renderer.SetDrawColor(borderSave.R, borderSave.G, borderSave.B, 255)
		rect := sdl.Rect{
			X: 0,
			Y: int32(y * screenScale),
			W: totalScreenWidth * screenScale,
			H: screenScale,
		}
		s.renderer.FillRect(&rect)

		waitDuration := zog.TStateDuration * scanlineCycles
		waitUntil := before.Add(waitDuration)
		for time.Now().Before(waitUntil) {
		}
	}

	// Draw each scanline
	for y := 0; y < screenHeight; y++ {
		borderSave := s.borderCol
		before := time.Now()

		s.renderer.SetDrawColor(borderSave.R, borderSave.G, borderSave.B, 255)
		borderleft := sdl.Rect{X: 0,
			Y: int32((y + borderTop) * screenScale),
			W: borderWidth * screenScale,
			H: screenScale,
		}
		s.renderer.FillRect(&borderleft)

		s.drawScanline(y)

		s.renderer.SetDrawColor(borderSave.R, borderSave.G, borderSave.B, 255)
		borderright := sdl.Rect{
			X: (borderWidth + screenWidth) * screenScale,
			Y: int32((y + borderTop) * screenScale),
			W: totalScreenWidth * screenScale,
			H: screenScale,
		}
		s.renderer.FillRect(&borderright)

		waitDuration := zog.TStateDuration * scanlineCycles
		waitUntil := before.Add(waitDuration)
		for time.Now().Before(waitUntil) {
		}
	}

	for y := 0; y < borderBelow; y++ {
		borderSave := s.borderCol
		before := time.Now()

		s.renderer.SetDrawColor(borderSave.R, borderSave.G, borderSave.B, 255)
		borderbelow := sdl.Rect{
			X: 0,
			Y: int32((y + borderTop + screenHeight) * screenScale),
			W: totalScreenWidth * screenScale,
			H: screenScale,
		}
		s.renderer.FillRect(&borderbelow)

		waitDuration := zog.TStateDuration * scanlineCycles
		waitUntil := before.Add(waitDuration)
		for time.Now().Before(waitUntil) {
		}
	}

	s.flashCount = (s.flashCount + 1) % 64
	s.renderer.Present()
}

func (s *Screen) drawScanline(y int) {
	lineMemLen := screenWidth / 8

	sector := (y & 0xc0) >> 6
	sectorRow := (y & 0x38) >> 3
	charRow := (y & 0x07)
	addr := screenMemStart + (sector*64+charRow*8+sectorRow)*lineMemLen

	colourRow := (sector << 3) + sectorRow
	colourAddr := colourMemStart + (256/8)*colourRow

	lineMem, err := s.mem.PeekBuf(uint16(addr), lineMemLen)
	if err != nil {
		panic(fmt.Errorf("Can't read screen memory at [%04X] len (%04X)", addr, lineMemLen))
	}
	colourMem, err := s.mem.PeekBuf(uint16(colourAddr), lineMemLen)
	if err != nil {
		panic(fmt.Errorf("Can't read screen memory at [%04X] len (%04X)", addr, lineMemLen))
	}

	for i, b := range lineMem {
		colourByte := colourMem[i]
		ink := colourByte & 0x07
		paper := (colourByte & 0x38) >> 3
		bright := (colourByte & 0x40) >> 6
		flash := (colourByte & 0x80) >> 7

		for bit := 0; bit < 8; bit++ {
			x := i*8 + bit
			if b&0x80 != 0 {
				s.SetDrawColour(true, ink, paper, bright, flash)
				//				s.renderer.SetDrawColor(0, 0, 0, 255)
			} else {
				s.SetDrawColour(false, ink, paper, bright, flash)
				//				s.renderer.SetDrawColor(255, 255, 255, 255)
			}
			//			fmt.Printf("x %d y %d i %d bit %d\n", x, y, i, b)
			rect := sdl.Rect{
				X: int32((x + borderWidth) * screenScale),
				Y: int32((y + borderTop) * screenScale),
				W: screenScale,
				H: screenScale,
			}
			s.renderer.FillRect(&rect)
			//			s.renderer.DrawPoint(x, y)
			b <<= 1
		}
	}
}

// Bright versions, non-bright are reduced from ff to d7
var Colours = []sdl.Color{
	{R: 0, G: 0, B: 0, A: 0},
	{R: 0, G: 0, B: 1, A: 0},
	{R: 1, G: 0, B: 0, A: 0},
	{R: 1, G: 0, B: 1, A: 0},
	{R: 0, G: 1, B: 0, A: 0},
	{R: 0, G: 1, B: 1, A: 0},
	{R: 1, G: 1, B: 0, A: 0},
	{R: 1, G: 1, B: 1, A: 0},
}

func (s *Screen) SetDrawColour(wantInk bool, ink, paper, bright, flash byte) {
	invert := flash != 0 && s.flashCount < 32 // 0-31 inverted, 32-63 not inverted

	index := paper
	if (wantInk && !invert) || (invert && !wantInk) {
		index = ink
	}
	c := Colours[index]
	factor := byte(0xd7)
	if bright != 0 {
		factor = 0xff
	}
	s.renderer.SetDrawColor(c.R*factor, c.G*factor, c.B*factor, 255)
}

func (s *Screen) Close() {
	s.renderer.Destroy()
	s.window.Destroy()
}
