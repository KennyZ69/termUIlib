package main

import (
	"log"

	// uiscreen "github.com/KennyZ69/termUIlibK/screen"
	applib "github.com/KennyZ69/termUIlibK/app"
	termui "github.com/KennyZ69/termUIlibK/term"
)

const (
	width  = 480
	height = 200
)

func main() {
	termui.RawMode()
	defer termui.DisableRaw()

	app := applib.NewApp(width, height)
	defer app.Screen.Clear()

	// app.Screen.AddStr(30, 0, "Welcome to termUIlibK!")
	// app.Screen.AddStr(30, 2, "Press 'q' to exit")
	// app.Screen.DrawAll()

	termui.PrintAt(width/2-10, height-2, "Welcome to termUIlibK!")
	termui.PrintAt(width/2-10, height-4, "Press 'q' to exit")

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
