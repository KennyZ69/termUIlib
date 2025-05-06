package termuilib

import (
	"sync"
)

type Screen struct {
	Width, Height int
	Cells         [][]Cell
	Dirty         [][]bool // dirty cells - to indicate if a cell has been modified
	wg            sync.WaitGroup
	title         string
	eventQ        chan Event
	quit          chan struct{}
	style         Style

	sync.Mutex
}

func NewScreen(width, height int) *Screen {
	cells := make([][]Cell, height)
	for cell := range cells {
		cells[cell] = make([]Cell, width)
	}

	dirty := make([][]bool, height)
	for cell := range dirty {
		dirty[cell] = make([]bool, width)
	}

	return &Screen{
		Width:  width,
		Height: height,
		Cells:  cells,
		Dirty:  dirty,
	}
}

func (s *Screen) Size() (int, int) {
	return s.Width, s.Height
}
