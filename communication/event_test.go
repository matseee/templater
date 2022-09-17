package communication

import "testing"

func Test_ToggleActiveEvent_IsType_should_validate(t *testing.T) {
	e := new(ToggleActiveEvent)

	if !e.IsType(ToggleActive) {
		t.Error("ToggleActiveEvent.IsType(ToggleActive) should return true but got false")
	}

	if e.IsType(OpenWindow) {
		t.Error("ToggleActiveEvent.IsType(OpenWindow) should return false but got true")
	}

	if e.IsType(KeyPressed) {
		t.Error("ToggleActiveEvent.IsType(KeyPressed) should return false but got true")
	}

	if e.IsType(Quit) {
		t.Error("ToggleActiveEvent.IsType(Quit) should return false but got true")
	}
}

func Test_OpenWindowEvent_IsType_should_validate(t *testing.T) {
	e := new(OpenWindowEvent)

	if !e.IsType(OpenWindow) {
		t.Error("OpenWindowEvent.IsType(OpenWindow) should return true but got false")
	}

	if e.IsType(ToggleActive) {
		t.Error("OpenWindowEvent.IsType(ToggleActive) should return false but got true")
	}

	if e.IsType(KeyPressed) {
		t.Error("OpenWindowEvent.IsType(KeyPressed) should return false but got true")
	}

	if e.IsType(Quit) {
		t.Error("OpenWindowEvent.IsType(Quit) should return false but got true")
	}
}

func Test_KeyPressedEvent_IsType_should_validate(t *testing.T) {
	e := new(KeyPressedEvent)

	if !e.IsType(KeyPressed) {
		t.Error("KeyPressedEvent.IsType(KeyPressed) should return true but got false")
	}

	if e.IsType(ToggleActive) {
		t.Error("KeyPressedEvent.IsType(ToggleActive) should return false but got true")
	}

	if e.IsType(OpenWindow) {
		t.Error("KeyPressedEvent.IsType(OpenWindow) should return false but got true")
	}

	if e.IsType(Quit) {
		t.Error("KeyPressedEvent.IsType(Quit) should return false but got true")
	}
}

func Test_QuitEvent_IsType_should_validate(t *testing.T) {
	e := new(QuitEvent)

	if !e.IsType(Quit) {
		t.Error("QuitEvent.IsType(Quit) should return true but got false")
	}

	if e.IsType(ToggleActive) {
		t.Error("QuitEvent.IsType(ToggleActive) should return false but got true")
	}

	if e.IsType(OpenWindow) {
		t.Error("QuitEvent.IsType(OpenWindow) should return false but got true")
	}

	if e.IsType(KeyPressed) {
		t.Error("QuitEvent.IsType(KeyPressed) should return false but got true")
	}
}
