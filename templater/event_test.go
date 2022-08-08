package templater

import "testing"

func Test_CreateEvent(t *testing.T) {
	event := CreateEvent()

	if event.Type != StatusEvent || event.ValueBool != false || event.ValueInt != 0 || event.ValueString != "" {
		t.Error("CreateEvent should create an empty event - FAILED")
	} else {
		t.Log("CreateEvent created an empty event - PASSED")
	}
}
