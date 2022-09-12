package templater

import "testing"

func Test_CreateEvent_should_return_a_empty_event_of_type_StatusEvent(t *testing.T) {
	event := CreateEvent()

	if event.Type != StatusEvent || event.ValueBool != false || event.ValueInt != 0 || event.ValueString != "" {
		t.Error("CreateEvent() should create an empty event")
	}
}

func Test_CreateStatusEvent_should_return_an_event_with_type_StatusEvent(t *testing.T) {
	ev := CreateStatusEvent(true)

	if ev.Type != StatusEvent {
		t.Error("CreateStatusEvent(true) should create a StatusEvent")
	}

	if !ev.ValueBool {
		t.Error("CreateStatusEvent(true) should create a StatusEvent with ValueBool == true")
	}
}

func Test_CreateOpenSettingsEvent_should_return_an_event_with_type_OpenSettingsEvent(t *testing.T) {
	ev := CreateOpenSettingsEvent()

	if ev.Type != OpenSettingsEvent {
		t.Error("CreateOpenSettingsEvent() should create an OpenSettingsEvent")
	}
}

func Test_CreateQuitEvent_should_return_an_event_with_type_QuitEvent(t *testing.T) {
	ev := CreateQuitEvent()

	if ev.Type != QuitEvent {
		t.Error("CreateQuitEvent() should create a QuitEvent")
	}
}
