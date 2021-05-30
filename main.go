package main

import (
	"time"
	"github.com/vladbalmos/gosnake/core"
	"github.com/vladbalmos/gosnake/game"
)

const FPS = 25

func now() uint {
	return uint(time.Now().UnixNano() / int64(time.Millisecond))
}

func wait(start uint) uint {
	speed := uint(1000 / FPS)
	delta := now() - start
	if delta >= speed {
		return 0
	}

	waitTime := speed - delta

	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	return waitTime
}

func main() {
	screen := core.NewScreen()
	defer screen.Cleanup()
	screen.Refresh()

	game := game.New(screen)

	eventLoop := core.NewEventLoop(screen)
	eventChannel := make(chan core.Event)
	go eventLoop.Start(eventChannel)

	for {
		start := now()

		// Draw the current game state
		game.Draw()

		// Listen for input
		ev := <-eventChannel

		// Update the game state
		game.Update(ev)

		if game.Quit() {
			break
		}

		// Paint the screen
		screen.Refresh()
		wait(start)
	}

}
