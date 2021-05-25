// +build amd64
package main

import (
	"fmt"
	"time"
	"github.com/vladbalmos/gosnake/core"
)

func main() {
	screen := core.NewScreen()
	defer screen.Cleanup()
	screen.Refresh()

	eventLoop := core.NewEventLoop(screen)

	eventChannel := make(chan core.Event)
	go eventLoop.Start(eventChannel)

	for {
		ev := <-eventChannel
		if ev.IsQuit() {
			break
		}

		fmt.Println(ev)
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

}
