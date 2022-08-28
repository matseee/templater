package templater

import (
	"testing"
	"time"
)

func resetGlobals() {
	channelCreated = false
	channel = nil
}

func Test_CreateChannel(t *testing.T) {
	defer resetGlobals()

	channel := CreateChannel()

	if channel == nil {
		t.Error("CreateChannel should create a channel - FAILED")
	} else {
		t.Log("CreatedChannel did create a channel - PASSED")
	}

	channel2 := CreateChannel()

	if channel != channel2 {
		close(channel)
		close(channel2)
		t.Error("CreateChannel should only create one channel - FAILED")
	} else {
		t.Log("CreatedChannel created only one channel - PASSED")
	}

	close(channel)
}

func Test_SendEvent(t *testing.T) {
	defer resetGlobals()

	var receivedMessage bool = false

	event := CreateEvent()
	_, err := SendEvent(event)

	if err == nil {
		t.Error("SendEvent should throw an error that the channel is not initialized - FAILED")
	} else {
		t.Log("SendEvent threw an error that the channel is not initialized - PASSED")
	}

	channel := CreateChannel()

	go func() {
		for _event := range channel {
			receivedMessage = true

			if _event != event {
				t.Error("SendEvent did not send the right event - FAILED")
			} else {
				t.Log("SendEvent sent the right event - PASSED")
			}
		}
	}()

	_event, err := SendEvent(event)

	if err != nil {
		str := err.Error() + " - FAILED"
		t.Error(str)
	}

	if _event != event {
		t.Error("SendEvent did not send the required event - FAILED")
	} else {
		t.Log("SendEvent sent the required event - PASSED")
	}

	time.Sleep(200 * time.Millisecond)

	if receivedMessage == false {
		t.Error("SendEvent did not receive the required event - FAILED")
	} else {
		t.Log("SendEvent received the right event - PASSED")
	}

	close(channel)
}
