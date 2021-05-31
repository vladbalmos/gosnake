package game

import (
	"github.com/vladbalmos/gosnake/core"
)

func InitialTransitionTable(event *core.Event) TransitionFunction {
	if event.IsSpace() {
		return initialToRunning
	}

	if event.IsQuit() {
		return transitionToQuit
	}

	return nil
}

func initialToRunning(g *game) *core.State {
	currentState := *g.state
	currentState.Id = STATE_RUNNING
	currentState.Running = true
	currentState.MessageForPlayer = ""
	currentState.Snake = g.newSnake(3)
	currentState.Food = g.newFood()
	return &currentState
}
