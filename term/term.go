package termuilib

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"golang.org/x/term"
)

var state *term.State

func RawMode() error {
	// Setting the terminal to raw mode
	var err error
	state, err = term.MakeRaw(0) // 0 for stdin
	if err != nil {
		return fmt.Errorf("failed to set raw mode: %s", err)
	}

	// handle exit and then read one byte at a time
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		log.Printf("Interrupted\nExiting ... \n")
		os.Exit(0)
	}()

	fmt.Print(clearScreen)
	fmt.Print(hideCursor)
	PrintAt(70, 0, "Hello to TermUI!")

	return nil
}

func DisableRaw() {
	if state != nil {
		term.Restore(0, state)
		fmt.Print(showCursor)
	}
}
func Listen(keyChan chan rune) error {
	var Err error
	go func() {
		buf := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				Err = fmt.Errorf("error reading from stdin: %s", err)
				PrintAt(0, 0, "Error reading from stdin")
				close(keyChan)
				return
			}
			keyChan <- rune(buf[0])
		}
	}()

	return Err
}
