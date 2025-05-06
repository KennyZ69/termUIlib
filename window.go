package termuilib

// Window should be used as a base class for other nodes, as the parent, it should not be interacted with in any way
type Window struct {
	x, y, width, height int

	// the content of the window size
	inX, inY, inW, inH int

	// padding
	pTop, pBot, pLeft, pRight int

	// bg Color

	border bool

	hasFocus bool

	//! *optional*
	// when the window is focused or when it loses focus
	focus, blur func()

	// draw to actually draw the window on the screen
	draw func()
}

func NewWindow(width, height int) *Window {
	return &Window{
		width:  width,
		height: height,
	}
}
