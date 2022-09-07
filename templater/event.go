package templater

type EventStatus int64

const (
	StatusEvent EventStatus = iota
	KeylistenerEvent
	OpenSettingsEvent
	QuitEvent
)

type Event struct {
	Type        EventStatus
	ValueBool   bool
	ValueInt    int64
	ValueString string
}

func CreateEvent() Event {
	return Event{
		Type:        StatusEvent,
		ValueBool:   false,
		ValueInt:    0,
		ValueString: "",
	}
}

func CreateKeylistenerEvent(key string) Event {
	event := CreateEvent()
	event.Type = KeylistenerEvent
	event.ValueString = key
	return event
}
