package game

import (
	"github.com/vladbalmos/gosnake/core"
)

func PausedTransitionTable(event *core.Event) TransitionFunction {
	if event.IsSpace() {
		return advanceSnake
	}

	if event.IsQuit() {
		return transitionToQuit
	}

	return nil
}
