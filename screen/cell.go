package termuilib

import (
	"fmt"
	termui "github.com/KennyZ69/termUIlibK/term"
)

type Cell struct {
	Ch     rune
	Fg, Bg int // Color attributes
}

func (c *Cell) Print(x, y int) {
	termui.MoveCursor(x, y)
	fmt.Printf("%q", string(c.Ch))
}
