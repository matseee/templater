package keylistener

import (
	"testing"

	"github.com/matseee/templater/communication"
)

func createChannelAndKeyloggerMock() (communication.EventChannel, KeyloggerMock) {
	return communication.CreateEventChannel(), createKeyloggerMock()
}

func createChannelAndKeyloggerMockAndKeylistener() (communication.EventChannel, KeyloggerMock, Keylistener) {
	eventChannel, keylogger := createChannelAndKeyloggerMock()
	keylistener, _ := CreateKeylistener(eventChannel, &keylogger)
	return eventChannel, keylogger, keylistener
}

func createChannelAndKeyloggerAndActivatedKeylistener() (communication.EventChannel, KeyloggerMock, Keylistener) {
	eventChannel, keylogger, keylistener := createChannelAndKeyloggerMockAndKeylistener()
	keylistener.Activate()
	return eventChannel, keylogger, keylistener
}

func Test_CreateKeylistener_should_create_an_initial_keylistener_object(t *testing.T) {
	_, _, keylistener := createChannelAndKeyloggerMockAndKeylistener()

	status := keylistener.GetStatus()

	if status.IsActive || status.KeyEventCounter > 0 {
		t.Error("CreateKeylistener did not create an normal keylistener object")
	}
}

func Test_KeylistenerActivate_should_active_the_keylistener_object(t *testing.T) {
	_, _, keylistener := createChannelAndKeyloggerAndActivatedKeylistener()

	if !keylistener.GetStatus().IsActive {
		t.Error("Keylistener.Activate() should activate the keylistener object")
	}
}

func Test_KeylistenerDeactivate_should_deactivate_the_keylistener_object(t *testing.T) {
	_, _, keylistener := createChannelAndKeyloggerAndActivatedKeylistener()

	keylistener.Deactivate()

	if keylistener.GetStatus().IsActive {
		t.Error("Keylistener.Deactivate() should deactivate the keylistener object")
	}
}

func Test_Keylistener_should_increment_the_keyevent_counter_after_keylogger_keyevent(t *testing.T) {
	_, keylogger, keylistener := createChannelAndKeyloggerAndActivatedKeylistener()

	keylogger.SendKey("A")

	if keylistener.GetStatus().KeyEventCounter == 0 {
		t.Error("Keylistener should increment the keyevent counter after keylogger keyevent")
	}
}

func Test_Keylistener_should_send_an_event_to_eventchannel_after_keylogger_keyevent(t *testing.T) {
	const RESULT_KEY = "A"
	var gotKeyEvent = false

	eventChannel, keylogger, _ := createChannelAndKeyloggerAndActivatedKeylistener()

	go func() {
		for event := range eventChannel.Channel {
			if event.Type == communication.KeylistenerEvent && event.ValueString == RESULT_KEY {
				gotKeyEvent = true
			}
		}
	}()

	keylogger.SendKey("A")

	if !gotKeyEvent {
		t.Error("Keylistener should send an event over the eventchannel but did not send one")
	}
}

func Test_Keylistener_should_not_increment_the_keyevent_counter_after_keylogger_event_when_deactivated(t *testing.T) {
	_, keylogger, keylistener := createChannelAndKeyloggerAndActivatedKeylistener()

	keylistener.Deactivate()
	keylogger.SendKey("A")

	if keylistener.GetStatus().KeyEventCounter > 0 {
		t.Error("Keylistener should not increment the keyevent counter after keylogger event when deactivated but keyevent counter is bigger than 0")
	}
}
