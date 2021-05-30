package game

import (
	//"fmt"
	"github.com/vladbalmos/gosnake/core"
)

const STATE_INITIAL = 0
const STATE_RUNNING = 1
const STATE_PAUSED = 2

// Maps state ids to transitions table functions
var transitionsTableFunctions map[uint]core.TransitionTable

func addTransitionTable(stateId uint, transitionTable core.TransitionTable) {
	transitionsTableFunctions[stateId] = transitionTable
}

type Game struct {
	screen *core.Screen
	state *core.State
}

func New(screen *core.Screen) *Game {
	state := &core.State{
		Id: STATE_INITIAL,
		Score: 0,
		MessageForPlayer: "Press SPACE to start the game",
	}

	game := &Game{
		state: state,
		screen: screen,
	}

	setupTransitionsTables()
	return game
}

func (g *Game) Draw() {
	// TODO: erase?
	g.showPlayerMessage()
	g.drawFood()
	g.drawSnake()
	g.drawScore()
}

func (g *Game) Update(ev core.Event) {
	g.stateTransition(ev)
}

func (g *Game) Quit() bool {
	return g.state.Quit
}

func (g *Game) stateTransition(ev core.Event) {
	transitionTable := transitionsTableFunctions[g.state.Id]
	transitionFunction := transitionTable(&ev)

	if transitionFunction == nil {
		return
	}

	g.state = transitionFunction(g.state)
}

func (g *Game) showPlayerMessage() {
	if g.state.Running {
		return
	}
}

func (g *Game) drawFood() {
	if !g.state.Running {
		return
	}
}

func (g *Game) drawSnake() {
	if !g.state.Running {
		return
	}
}

func (g *Game) drawScore() {
	if !g.state.Running {
		return
	}
}

func transitionToQuit(s *core.State) *core.State {
	currentState := *s
	currentState.Quit = true
	return &currentState
}

func setupTransitionsTables() {
	transitionsTableFunctions = make(map[uint]core.TransitionTable)
	transitionsTableFunctions[STATE_INITIAL] = InitialTransitionTable
}

