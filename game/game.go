package game

import (
	//"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/vladbalmos/gosnake/core"
)

const (
	STATE_INITIAL = 0
	STATE_RUNNING = 1
	STATE_PAUSED = 2

	SECOND = 1000
	FPS = 30

	START_MESSAGE = "Press SPACE to start the game"
	START_SPEED = 4
)

type TransitionTable func(event *core.Event) TransitionFunction
type TransitionFunction func(g *game) *core.State

// Maps state ids to transitions table functions
var transitionsTableFunctions map[uint]TransitionTable

func addTransitionTable(stateId uint, transitionTable TransitionTable) {
	transitionsTableFunctions[stateId] = transitionTable
}

type screenState struct {
	playerMessageVisible bool
	lastScore int
}

type game struct {
	screen *core.Screen
	screenState *screenState
	state *core.State
	newPendingDirection int
	lastAnimationFrameTime uint
	snakeHeadStartPos core.Point
	collisionMatrix [][]uint8
	lastEvent core.Event
}

func New(screen *core.Screen) *game {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	state := &core.State{
		Id: STATE_INITIAL,
		Score: 0,
		Speed: START_SPEED,
		MessageForPlayer: START_MESSAGE,
	}

	collisionMatrix := make([][]uint8, screen.GameAreaWidth)
	for i := range collisionMatrix {
		collisionMatrix[i] = make([]uint8, screen.GameAreaHeight)
	}

	game := &game{
		state: state,
		screenState: &screenState{
			lastScore: -1,
		},
		screen: screen,
		snakeHeadStartPos: screen.GameAreaCenter(),
		newPendingDirection: -1,
		collisionMatrix: collisionMatrix,
	}

	setupTransitionsTables()
	return game
}

func (g *game) Draw() {
	g.screen.EraseGameArea()
	g.showPlayerMessage()
	g.drawFood()
	g.drawSnake()
	g.drawScore()
}

func (g *game) Update(ev core.Event) {
	g.lastEvent = ev
	g.stateTransition()
}

func (g *game) Quit() bool {
	return g.state.Quit
}

func (g *game) newSnake(length uint) core.Cralwer {
	snake := NewSnake(length)
	g.updateCollisionMatrix(snake)
	return snake
}

func (g *game) newFood() core.Point {
	var x int;
	var y int;

	for {
		x = rand.Intn(g.screen.GameAreaWidth - 1) + 1
		y = rand.Intn(g.screen.GameAreaHeight - 1) + 1

		if g.collisionMatrix[x][y] == 0 {
			break
		}
	}
	foodCoords := core.Point{
		X: x,
		Y: y,
	}

	return foodCoords
}

func (g *game) updateCollisionMatrix(snake core.Cralwer) {
	var segmentCallback interface{} = func (segment *SnakeSegment) {
		segmentScreenCoords := g.translateSnakeSegmentCoords(segment.Coords)
		g.collisionMatrix[segmentScreenCoords.X - 1][segmentScreenCoords.Y - 1] = 1
	}
	snake.Traverse(segmentCallback)
}

func (g *game) collisionDetected() bool {
	return false
}

func (g *game) caughtFood() bool {
	translatedHeadCoords := g.translateSnakeSegmentCoords(g.state.Snake.HeadCoords())
	result := translatedHeadCoords == g.state.Food
	if result {
		log.Print("Caught food")
	}
	return result
}

func (g *game) eatFood() {
	g.state.Snake.IncreaseLength()
	g.state.Score += 1
}

func (g *game) translateSnakeSegmentCoords(coords core.Point) core.Point {
	return core.Point{
		X: g.snakeHeadStartPos.X + coords.X,
		Y: g.snakeHeadStartPos.Y + coords.Y,
	}
}

func (g *game) stateTransition() {
	transitionTable := transitionsTableFunctions[g.state.Id]
	transitionFunction := transitionTable(&g.lastEvent)

	if transitionFunction == nil {
		return
	}

	g.state = transitionFunction(g)
}

func (g *game) showPlayerMessage() {
	if g.state.Running && !g.state.Paused {
		if g.screenState.playerMessageVisible {
			g.screen.HideMessage()
			g.screenState.playerMessageVisible = false
		}
		return
	}

	if !g.screenState.playerMessageVisible {
		g.screen.ShowMessage(g.state.MessageForPlayer)
		g.screenState.playerMessageVisible = true
	}
}

func (g *game) drawFood() {
	if !g.state.Running {
		return
	}

	g.screen.DrawFood(g.state.Food)
}

func (g *game) drawSnake() {
	if !g.state.Running {
		return
	}

	var segmentCallback interface{} = func (segment *SnakeSegment) {
		segmentScreenCoords := g.translateSnakeSegmentCoords(segment.Coords)
		g.screen.DrawSnakeSegment(segmentScreenCoords)
	}

	g.state.Snake.Traverse(segmentCallback)
}

func (g *game) drawScore() {
	if int(g.state.Score) == g.screenState.lastScore {
		return
	}

	g.screen.ShowScore(g.state.Score)
	g.screenState.lastScore = int(g.state.Score)
}

func transitionToQuit(g *game) *core.State {
	currentState := *g.state
	currentState.Quit = true
	return &currentState
}

func setupTransitionsTables() {
	transitionsTableFunctions = make(map[uint]TransitionTable)
	transitionsTableFunctions[STATE_INITIAL] = InitialTransitionTable
	transitionsTableFunctions[STATE_RUNNING] = RunningTransitionTable
	transitionsTableFunctions[STATE_PAUSED] = PausedTransitionTable
}

