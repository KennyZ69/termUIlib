package termuilib

import (
	"sync"
)

const (
	queueSize = 100
)

type updateQueue struct {
	// queue of update functions to be used on the nodes of the screen
	f    func()
	done chan struct{}
}

type App struct {
	wg sync.WaitGroup

	Screen *Screen
	Keys   chan rune
	Quit   chan struct{}

	focus Node
	root  Node

	mouse bool
	paste bool

	style Style

	updates chan updateQueue

	events chan Event
}

func NewApp(width, height int) *App {
	return &App{
		Screen: NewScreen(width, height),
		Keys:   make(chan rune),
		Quit:   make(chan struct{}, 1),

		events:  make(chan Event, queueSize),
		updates: make(chan updateQueue, queueSize),
	}
}

func (app *App) SetRoot(root Node) {
	app.root = root
}

func (app *App) SetFocus(focus Node) {
	app.focus = focus
}

func (app *App) SetMouse(mouse bool) {
	app.mouse = mouse
}

func (app *App) SetPaste(paste bool) {
	app.paste = paste
}
