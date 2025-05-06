package termuilib

import ()

// this would be the top level interface for all graphical elements
type Node interface {

	// SetPos sets the position of the node on screen (in a window)
	SetPos(x, y, width, height int)
}
