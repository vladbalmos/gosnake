package main

import (
	"log"
	"os"
	"github.com/vladbalmos/gosnake/core"
	Game "github.com/vladbalmos/gosnake/game"
)

func main() {
	screen := core.NewScreen()
	defer screen.Cleanup()
	screen.Refresh()

	f, _ := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)

	game := Game.New(screen)

	eventLoop := core.NewEventLoop(screen)
	eventChannel := make(chan core.Event)
	go eventLoop.Start(eventChannel)

	for {
		start := core.Now()

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
		core.Wait(start, Game.FPS)
	}

}
