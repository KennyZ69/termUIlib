package termuilib

import "math"

// Text alignment constants withing a window
const (
	AlignLeft = iota
	AlignCenter
	AlignRight
	AlignTop = 0
	AlignBot = 2
)

// PrintAt prints a string at the specified coordinates in a window, returning the number of bytes written and width used up by the string
func PrintAt(screen Screen, x, y int, text string, width, align int) (int, int) {
	start, end, width := uiPrint(x, y, 0, width, align, text)
	return end - start, width
}

func uiPrint(screen Screen, x, y, skipW, maxW, align int, text string) (int, int, int) {
	totalW, totalH := screen.Size()
}

// PrintTest is a test function to print a simple string onto given coordinates without colors or styles
func PrintTest(x, y int, text string) {
	PrintAt(x, y, text, math.MaxInt32, AlignLeft)
}
