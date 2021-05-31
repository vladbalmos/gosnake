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
	currentState.Score = 0
	currentState.Speed = START_SPEED
	currentState.MessageForPlayer = ""
	currentState.Snake = g.newSnake(10)
	currentState.Food = g.newFood()
	return &currentState
}
