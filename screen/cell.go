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
	if c.Fg != 0 || c.Bg != 0 {
		fmt.Printf("\033[38;5;%dm\033[48;5;%dm%c\033[0m", c.Fg, c.Bg, c.Ch)
	} else {
		fmt.Printf("%c", c.Ch)
	}
}
