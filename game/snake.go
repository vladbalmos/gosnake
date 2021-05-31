package game

import (
	"github.com/vladbalmos/gosnake/core"
)

const (
	DIRECTION_NORTH = 0
	DIRECTION_EAST = 90
	DIRECTION_SOUTH = 180
	DIRECTION_WEST = 270
)

type Snake struct {
	Head *SnakeSegment
	Tail *SnakeSegment
	Direction uint
}

type SnakeSegment struct {
	Next *SnakeSegment
	Prev *SnakeSegment
	Coords core.Point
}

func NewSnake(length uint) core.Cralwer {
	startX := 0
	head := &SnakeSegment{
		Coords: core.Point{
			X: startX,
			Y: 0,
		},
	}

	snake := &Snake{
		Head: head,
		Tail: head,
		Direction: DIRECTION_EAST,
	}

	for i := 1; uint(i) < length; i++ {
		snake.Append(startX - i, 0)
	}

	return snake
}

func (s *Snake) Append(x int, y int) {
	segment := &SnakeSegment{
		Next: s.Tail,
		Coords: core.Point{
			X: x,
			Y: y,
		},
	}

	s.Tail.Prev = segment
	s.Tail = segment
}

func (s *Snake) Prepend(x int, y int) {
	segment := &SnakeSegment{
		Prev: s.Head,
		Coords: core.Point{
			X: x,
			Y: y,
		},
	}

	s.Head.Next = segment
	s.Head = segment
}

func (s *Snake) ChangeDirection(dir uint) {
	if s.isSameAxisDirection(dir) {
		return
	}

	var x int = 0
	var y int = 0

	if dir == DIRECTION_EAST {
		x += 1
	} else if dir == DIRECTION_WEST {
		x -= 1
	}

	if dir == DIRECTION_NORTH {
		y -= 1
	} else if dir == DIRECTION_SOUTH {
		y += 1
	}

	lastSegmentCoords := s.Head.Coords

	s.Head.Coords.X += x
	s.Head.Coords.Y += y

	currentSegment := s.Head.Prev
	for {
		if currentSegment == nil {
			break
		}

		currentCoords := currentSegment.Coords
		currentSegment.Coords = lastSegmentCoords
		lastSegmentCoords = currentCoords

		currentSegment = currentSegment.Prev
	}
}

func (s *Snake) Advance() {
	var x int
	var y int

	switch s.Direction {
	case DIRECTION_EAST:
		x = 1
		y = 0
	case DIRECTION_SOUTH:
		x = 0
		y = 1
	case DIRECTION_WEST:
		x = -1
		y = 0
	case DIRECTION_NORTH:
		x = 0
		y = -1
	}

	lastSegmentCoords := s.Head.Coords

	s.Head.Coords.X += x
	s.Head.Coords.Y += y

	currentSegment := s.Head.Prev
	for {
		if currentSegment == nil {
			break
		}

		currentCoords := currentSegment.Coords
		currentSegment.Coords = lastSegmentCoords
		lastSegmentCoords = currentCoords

		currentSegment = currentSegment.Prev
	}

}

func (s *Snake) isSameAxisDirection(dir uint) bool {
	if s.Direction == DIRECTION_EAST || s.Direction == DIRECTION_WEST {
		if dir == DIRECTION_EAST || dir == DIRECTION_WEST {
			return true
		}

		return false
	}

	if s.Direction == DIRECTION_NORTH || s.Direction == DIRECTION_SOUTH {
		if dir == DIRECTION_NORTH || dir == DIRECTION_SOUTH {
			return true
		}
	}

	return false
}

func (s *Snake) Traverse(callback interface{}) {
	currentSegment := s.Head

	for {
		if currentSegment == nil {
			return
		}

		callback.(func(segment *SnakeSegment))(currentSegment)
		currentSegment = currentSegment.Prev
	}
}
