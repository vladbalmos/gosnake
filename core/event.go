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
	106: "j",
	74: "j",
	104: "h",
	72: "h",
	107: "k",
	75: "k",
	108: "l",
	76: "l",
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

func (ev *Event) IsUp() bool {
	return ev.Name == "k" || ev.Name == "up"
}

func (ev *Event) IsDown() bool {
	return ev.Name == "j" || ev.Name == "down"
}

func (ev *Event) IsLeft() bool {
	return ev.Name == "h" || ev.Name == "left"
}

func (ev *Event) IsRight() bool {
	return ev.Name == "l" || ev.Name == "right"
}

func (ev *Event) IsDir() bool {
	return ev.IsUp() || ev.IsDown() || ev.IsLeft() || ev.IsRight()
}

func (ev *Event) IsEsc() bool {
	return ev.Name == "esc"
}

func (ev *Event) IsNoInput() bool {
	return ev.Name == EV_NO_INPUT
}

func (ev *Event) IsQuit() bool {
	return ev.Name == "q" || ev.Name == EV_SIGNAL_QUIT
}

