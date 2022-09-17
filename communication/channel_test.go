package communication

import (
	"testing"
	"time"
)

func Test_CreateEventChannel_should_create_a_EventChannel_object_with_Channel(t *testing.T) {
	eventChannel := CreateEventChannel()

	if eventChannel.Channel == nil {
		t.Error("CreateEventChannel() should create a EventChannel object with a chan Event")
	}
}

func Test_EventChannelSend_should_return_an_error_when_channel_is_nil(t *testing.T) {
	eventChannel := EventChannel{}
	event := new(ToggleActiveEvent)

	err := eventChannel.Send(event)

	if err == nil {
		t.Error("EventChannel.Send() should return an error when channel is nil")
	}
}

func Test_EventChannelSend_should_send_a_event_over_the_channel(t *testing.T) {
	var receivedMessage = false
	event := new(ToggleActiveEvent)

	eventChannel := CreateEventChannel()

	go func() {
		for _event := range eventChannel.Channel {
			if _event == event {
				receivedMessage = true
			}
		}
	}()

	eventChannel.Send(event)

	time.Sleep(200 * time.Millisecond)
	if receivedMessage == false {
		t.Error("EventChannel.Send() should send an event over the EventChannel.")
	}
}
