package main

import (
	"fmt"
	"time"

	// uiscreen "github.com/KennyZ69/termUIlibK/screen"
	termui "github.com/KennyZ69/termUIlibK/term"
)

func main() {
	termui.RawMode()
	defer termui.DisableRaw()

	keysChan := make(chan rune)

	termui.Listen(keysChan)

	// scr := uiscreen.NewScreen(80, 24)
	// scr.Clear()
	// scr.AddStr(36, 0, "Hello to testing!")
	// scr.DrawAll()

	for {
		select {
		case key := <-keysChan:
			if key == 'q' || key == 'Q' {
				termui.PrintAt(0, 0, "Exiting... ")
				return
			}

			termui.PrintAt(2, 5, fmt.Sprintf("Key pressed: %q [code: %d]", key, key))
			// scr.AddStr(2, 5, fmt.Sprintf("Key pressed: %q [code: %d]", key, key))
			// scr.DrawAll()
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

}
