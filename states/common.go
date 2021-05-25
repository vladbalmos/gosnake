package states

import "github.com/vladbalmos/gosnake/core"

func TransitionToQuit(_ core.StateTransitioner) core.StateTransitioner {
	return nil
}
