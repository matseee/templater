package templater

import "testing"

func Test_CreateEvent_should_return_a_empty_event_of_type_StatusEvent(t *testing.T) {
	event := CreateEvent()

	if event.Type != StatusEvent || event.ValueBool != false || event.ValueInt != 0 || event.ValueString != "" {
		t.Error("CreateEvent should create an empty event - FAILED")
	}
}
