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
    Score uint
    Snake *Snake
    Food *Food
    MessageForPlayer string
}

type StateTransitioner interface{
    Transition(event *Event) TransitionFunction
}

type TransitionFunction func(s StateTransitioner) StateTransitioner;


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
