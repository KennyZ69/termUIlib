package termuilib

type Screen struct {
	Width, Height int
	Cells         [][]Cell
	Dirty         [][]bool // dirty cells - to indicate if a cell has been modified
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

func (s *Screen) Clear() {
	for y := range s.Cells {
		for x := range s.Cells[y] {
			s.Dirty[y][x] = true
		}
	}
}

func (s *Screen) AddChar(x, y int, char rune) {
	if x >= 0 && x < s.Width && y >= 0 && y < s.Height {
		if s.Cells[y][x].Ch != char {
			s.Cells[y][x].Ch = char
			s.Dirty[y][x] = true
		}
	}
}

func (s *Screen) AddStr(x, y int, str string) {
	for i, char := range str {
		if x+i < s.Width {
			s.AddChar(x+i, y, char)
		}
	}
}

func (s *Screen) DrawAll() {
	for y := range s.Cells {
		for x := range s.Cells[y] {
			if s.Dirty[y][x] {
				s.Dirty[y][x] = false
				s.Cells[y][x].Print(x, y)
			}
		}
	}
}
