package running

import "github.com/vladbalmos/gosnake/states"
import "github.com/vladbalmos/gosnake/core"

type baseState core.State

type State struct {
    *baseState
    GamePaused bool
}

func New() *State {
    return &State{
        baseState: &baseState{
            Score: 0,
            Snake: nil,
            Food: nil,
            MessageForPlayer: "",
        },
        GamePaused: false,
    }
}

func (s *State) Transition(event *core.Event) core.TransitionFunction {
    if event.IsSpace() {
        return transitionToPlayPause
    }

    if event.IsCollision() {
        return transitionToOver
    }

    if event.IsQuit() {
        return states.TransitionToQuit
    }

    return nil
}

func transitionToPlayPause(s core.StateTransitioner) core.StateTransitioner {
    state := s.(*State)

    state.GamePaused = !state.GamePaused
    return state
}

func transitionToOver(_ core.StateTransitioner) core.StateTransitioner {
    return nil
}
