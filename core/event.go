package core

const EV_SIGNAL_QUIT = "quit"
const EV_NO_INPUT = "no:input"

var keyMap = map[int]string {
	113: "q",
	27: "esc",
	32: "space",
	258: "down",
	259: "up",
	260: "left",
	261: "right",
	49: "1",
	50: "2",
	51: "3",
}

func NewEvent(name string) *Event {
	return &Event{
		Name: name,
	}
}

func NewEventFromKeyboard(key Key) *Event {
	keyName, exists := keyMap[int(key)]
	if exists {
		return NewEvent(keyName)
	}

	return NewEvent(EV_NO_INPUT)
}

func (ev *Event) IsSpace() bool {
	return ev.Name == "space"
}

func (ev *Event) IsCollision() bool {
	return ev.Name == "collision"
}

func (ev *Event) IsQuit() bool {
	return ev.Name == "q" || ev.Name == EV_SIGNAL_QUIT
}

