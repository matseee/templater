package communication

type EventType int64

const (
	ToggleActive = iota
	OpenWindow
	KeyPressed
	Quit
)

type Event interface {
	IsType(EventType) bool
}

type ToggleActiveEvent struct {
	IsActive bool
}

func (c ToggleActiveEvent) IsType(t EventType) bool {
	return t == ToggleActive
}

type OpenWindowEvent struct{}

func (c OpenWindowEvent) IsType(t EventType) bool {
	return t == OpenWindow
}

type KeyPressedEvent struct {
	Key string
}

func (c KeyPressedEvent) IsType(t EventType) bool {
	return t == KeyPressed
}

type QuitEvent struct{}

func (c QuitEvent) IsType(t EventType) bool {
	return t == Quit
}
