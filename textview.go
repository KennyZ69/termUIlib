package termuilib

import (
	"strings"
	"sync"
)

type TextView struct {
	sync.Mutex
	*Window

	width, height int

	// text buffer
	text strings.Builder

	scrollable bool
	wrap       bool
}

func NewTextView(width, height int) *TextView {
	return &TextView{
		Window:     NewWindow(width, height),
		width:      width,
		height:     height,
		scrollable: true,
		wrap:       true,
	}
}
