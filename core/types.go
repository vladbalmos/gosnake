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
	Speed float32
	MessageForPlayer string
}

type Cralwer interface {
	HeadCoords() Point
	IncreaseLength()
	Append(x int, y int)
	Prepend(x int, y int)
	Traverse(callback interface{})
	Advance()
	ChangeDirection(dir uint)
	IsSameAxisDirection(dir uint) bool
}
