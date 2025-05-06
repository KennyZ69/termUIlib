package termuilib

import (
	"fmt"
)

type Cell struct {
	ch    rune
	combr []rune // combined runes
	style Style

	width int

	locked bool
	dirty  bool

	lastCh    rune
	lastStyle Style
	lastCombr []rune
}

type CellBuffer struct {
	width, height int
	cells         []Cell
}

func (c *Cell) Print(x, y int) {
	MoveCursor(x, y)
	fmt.Printf("%c", c.ch)
}

func (cb *CellBuffer) SetCell(x, y int, ch rune, combr []rune, style Style) {
	if x < 0 || x >= cb.width || y < 0 || y >= cb.height {
		return
	}

	c := &cb.cells[y*cb.width+x]
}

func (cb *CellBuffer) SetDirty(x, y int, dirty bool) {

}

func (cb *CellBuffer) GetCell(x, y int) Cell {
	if x < 0 || x >= cb.width || y < 0 || y >= cb.height {
		return Cell{}
	}
	c := cb.cells[y*cb.width+x]
	if c.width == 0 || c.ch < ' ' {
		c.ch = ' '
		c.width = 1
	}
	return c
}

func (cb *CellBuffer) Size() (int, int) {
	return cb.width, cb.height
}
