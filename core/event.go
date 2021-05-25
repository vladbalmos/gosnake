package core

func NewEvent(name string) *Event {
    return &Event{
        Name: name,
    }
}

func (ev *Event) IsSpace() bool {
    return ev.Name == "space"
}

func (ev *Event) IsCollision() bool {
    return ev.Name == "collision"
}

func (ev *Event) IsQuit() bool {
    return ev.Name == "q"
}
