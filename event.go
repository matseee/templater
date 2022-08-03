package main

type TemplaterEventStatus int64

const (
	TemplaterStatus TemplaterEventStatus = iota
	OpenSettings
	Quit
)

type TemplaterEvent struct {
	Type        TemplaterEventStatus
	ValueBool   bool
	ValueInt    int64
	ValueString string
}

func CreateEvent() TemplaterEvent {
	return TemplaterEvent{
		Type:        TemplaterStatus,
		ValueBool:   false,
		ValueInt:    0,
		ValueString: "",
	}
}
