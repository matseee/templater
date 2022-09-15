package communication

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

func CreateStatusEvent(active bool) Event {
	ev := CreateEvent()
	ev.ValueBool = active
	return ev
}

func CreateOpenSettingsEvent() Event {
	ev := CreateEvent()
	ev.Type = OpenSettingsEvent
	return ev
}

func CreateKeylistenerEvent(key string) Event {
	event := CreateEvent()
	event.Type = KeylistenerEvent
	event.ValueString = key
	return event
}

func CreateQuitEvent() Event {
	ev := CreateEvent()
	ev.Type = QuitEvent
	return ev
}
