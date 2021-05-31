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

	if event.IsDir() {
		return changeSnakeDir
	}

	if event.IsEsc() {
		return gameOver
	}

	if event.IsSpace() {
		return pauseGame
	}

	return nil
}

func pauseGame(g *game) *core.State {
	currentState := *g.state
	currentState.Id = STATE_PAUSED
	currentState.MessageForPlayer = "Press SPACE to resume"
	currentState.Paused = true

	return &currentState
}

func gameOver(g *game) *core.State {
	currentState := *g.state
	currentState.Id = STATE_INITIAL
	currentState.MessageForPlayer = START_MESSAGE
	currentState.Running = false
	currentState.Snake = nil
	currentState.Score = 0

	return &currentState
}

func changeSnakeDir(g *game) *core.State {
	var dir uint

	if g.lastEvent.IsUp() {
		dir = DIRECTION_NORTH
	}

	if g.lastEvent.IsDown() {
		dir = DIRECTION_SOUTH
	}

	if g.lastEvent.IsLeft() {
		dir = DIRECTION_WEST
	}

	if g.lastEvent.IsRight() {
		dir = DIRECTION_EAST
	}

	if !g.state.Snake.IsSameAxisDirection(dir) {
		g.newPendingDirection = int(dir)
	}

	return advanceSnake(g)
}

func advanceSnake(g *game) *core.State {
	now := core.Now()
	delta := now - g.lastAnimationFrameTime
	threshold := uint(SECOND / g.state.Speed)

	if delta < threshold {
		return g.state
	}

	g.lastAnimationFrameTime = now

	currentState := *g.state
	currentState.Id = STATE_RUNNING
	currentState.Paused = false
	currentState.Running = true
	currentState.MessageForPlayer = ""

	if g.newPendingDirection > -1 {
		currentState.Snake.ChangeDirection(uint(g.newPendingDirection))
		g.newPendingDirection = -1
	} else {
		currentState.Snake.Advance()
	}

	if g.collisionDetected() {
		return transitionToCollision(g)
	}

	if g.caughtFood() {
		g.eatFood()
	}
	return &currentState
}

func transitionToCollision(g *game) *core.State {
	return nil
}
