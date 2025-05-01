package termuilib

import (
	"fmt"

	scr "github.com/KennyZ69/termUIlibK/screen"
	tui "github.com/KennyZ69/termUIlibK/term"
)

type App struct {
	Screen *scr.Screen
	Keys   chan rune
	Quit   chan struct{}
}

func NewApp(width, height int) *App {
	return &App{
		Screen: scr.NewScreen(width, height),
		Keys:   make(chan rune),
		Quit:   make(chan struct{}),
	}
}

func (app *App) Run() error {
	tui.RawMode()
	defer tui.DisableRaw()
	defer app.Screen.Clear()

	go tui.Listen(app.Keys)

	for {
		select {
		case key, ok := <-app.Keys:
			if !ok {
				app.Screen.Clear()
				tui.DisableRaw()
				return nil
			}
			// here I am using if but possible later switch
			if key == 'q' || key == 'Q' {
				tui.PrintAt(0, 0, "Pressed 'q' -> exiting now ...")
				close(app.Quit)
			}
			app.Screen.AddStr(2, 5, fmt.Sprintf("Key pressed: %c [code: %d]", key, key))
			app.Screen.DrawAll()
		case <-app.Quit:
			app.Screen.Clear()
			tui.DisableRaw()
			return nil
		}
	}
}
