package core

type coords struct {
	X uint
	Y uint
}

type Event struct {
	Name string
}

// Game state
type State struct {
	Id uint
	Paused bool
	Running bool
	Quit bool
	Score uint
	Snake *Snake
	Food *Food
	MessageForPlayer string
}

type TransitionTable func(event *Event) TransitionFunction
type TransitionFunction func(s *State) *State


type Snake struct {
	Head *SnakeSegment
	Tail *SnakeSegment
}

type SnakeSegment struct {
	coords
	Next *SnakeSegment
	Prev *SnakeSegment
}

type Food struct {
	coords
}
