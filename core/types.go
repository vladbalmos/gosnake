package core

type Point struct {
	X int
	Y int
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
	Snake Cralwer
	Food Point
	MessageForPlayer string
}

type Cralwer interface {
	Append(x int, y int)
	Prepend(x int, y int)
	Traverse(callback interface{})
	Advance()
}
