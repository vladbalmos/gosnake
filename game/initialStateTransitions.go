package game

import (
	"github.com/vladbalmos/gosnake/core"
)

func InitialTransitionTable(event *core.Event) core.TransitionFunction {
	if event.IsSpace() {
		return initialToRunning
	}

	if event.IsQuit() {
		return transitionToQuit
	}

	return nil
}

func initialToRunning(s *core.State) *core.State {
	return nil
}
