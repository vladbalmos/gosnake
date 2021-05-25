package core

import (
	//"fmt"
	"os"
	"os/signal"
)

type Loop struct {
	signalCh chan os.Signal
	screen *Screen
}

func NewEventLoop(s *Screen) *Loop {

	loop := &Loop{
		screen: s,
		signalCh: make(chan os.Signal, 1),
	}
	signal.Notify(loop.signalCh, os.Interrupt)

	return loop
}

func (l *Loop) Start(ch chan Event) {
	for {
		ev := l.listen()
		ch<- *ev
	}
}

func (l *Loop) listen() *Event {
	quitSignal := l.checkForQuitSignal()
	if quitSignal {
		return NewEvent(EV_SIGNAL_QUIT)
	}

	input := l.screen.GetInput()
	if input == 0 {
		return NewEvent(EV_NO_INPUT)
	}

	return NewEventFromKeyboard(input)
}

func (l *Loop) checkForQuitSignal() bool {
	select {
	case <-l.signalCh:
		return true
	default:
		return false
	}
}
