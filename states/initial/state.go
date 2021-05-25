package initial

import "github.com/vladbalmos/gosnake/core"
import "github.com/vladbalmos/gosnake/states"
import "github.com/vladbalmos/gosnake/states/running"

type State struct {
    *core.State
}

func New() *State {
    return &State{
        &core.State{
            Score: 0,
            Snake: nil,
            Food: nil,
            MessageForPlayer: "Press SPACE to start the game",
        },
    }
}

func (s *State) Transition(event *core.Event) core.TransitionFunction {
    if event.IsSpace() {
        return transitionToRunning
    }

    if event.IsQuit() {
        return states.TransitionToQuit
    }

    return nil
}

func transitionToRunning(s core.StateTransitioner) core.StateTransitioner {
    r := running.New()
    return r
}

