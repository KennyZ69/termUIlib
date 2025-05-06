package termuilib

import "fmt"

const (
	// escape codes for control
	clearScreen = "\033[2J"
	hideCursor  = "\033[?25l"
	showCursor  = "\033[?25h"
)

type MouseEvent int16

// moveCursor moves the cursor to specified coordinates in the terminal
func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x) // row ; col
}

// PrintAt prints a string at the specified coordinates in the window
// func PrintAt(x, y int, str string) {
// 	MoveCursor(x, y)
// 	fmt.Print(str)
// }
