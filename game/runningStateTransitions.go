package game

import (
	"github.com/vladbalmos/gosnake/core"
)

func RunningTransitionTable(event *core.Event) TransitionFunction {
	if event.IsQuit() {
		return transitionToQuit
	}

	if event.IsNoInput() {
		return advanceSnake
	}
	return nil
}

func advanceSnake(g *game) *core.State {
	currentState := *g.state
	currentState.Id = STATE_RUNNING
	currentState.Running = true
	currentState.MessageForPlayer = ""

	currentState.Snake.Advance()
	if g.collisionDetected() {
		return transitionToCollision(g)

	}
	return &currentState
}

func transitionToCollision(g *game) *core.State {
	return nil
}
