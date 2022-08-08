package keylistener

import (
	"testing"
	"time"

	"github.com/matseee/templater/templater"
)

func Test_ListenToChannel(t *testing.T) {
	var c chan templater.Event = make(chan templater.Event)
	var k KeyloggerMock = KeyloggerMock{}

	err := ListenToChannel(c, &k)

	if err != nil {
		str := err.Error() + " - FAILED"
		t.Error(str)
	}

	if !k.isCreated {
		t.Error("Keylogger should be created - FAILED")
	} else {
		t.Log("Keylogger was created - PASSED")
	}

	if k.isListening {
		t.Error("Keylogger should not listen after creation - FAILED")
	} else {
		t.Log("Keylogger is not listening - PASSED")
	}

	event := templater.CreateEvent()
	event.Type = templater.StatusEvent
	event.ValueBool = true
	c <- event

	if !k.isListening {
		t.Error("Keylogger should listen after start event - FAILED")
	} else {
		t.Log("Keylogger is listening after start event - PASSED")
	}

	event = templater.CreateEvent()
	event.Type = templater.StatusEvent
	event.ValueBool = false
	c <- event

	time.Sleep(100 * time.Millisecond)
	if k.isListening {
		t.Error("Keylogger should not listen after stop event - FAILED")
	} else {
		t.Log("Keylogger isn't listening after stop event - PASSED")
	}

	err = k.Destroy()
	if err != nil {
		str := err.Error() + " - FAILED"
		t.Error(str)
	} else {
		t.Log("Keylogger destroyed without error - PASSED")
	}

	if k.isCreated {
		t.Error("Keylogger should not be created after destroy() - FAILED")
	} else {
		t.Log("Keylogger is destroyed - PASSED")
	}
}

// MOCKING KEYLOGGER
type KeyloggerMock struct {
	outChannel  chan KeyEvent
	isCreated   bool
	isListening bool
}

func (k *KeyloggerMock) Create() error {
	k.isCreated = true
	return nil
}

func (k *KeyloggerMock) Destroy() error {
	k.isCreated = false
	return nil
}

func (k *KeyloggerMock) Start() chan KeyEvent {
	if k.isListening {
		return k.outChannel
	}

	k.outChannel = make(chan KeyEvent)
	k.isListening = true
	return k.outChannel
}

func (k *KeyloggerMock) Stop() {
	if !k.isListening {
		return
	}

	close(k.outChannel)
	k.outChannel = nil

	k.isListening = false
}

func (k *KeyloggerMock) IsCreated() bool {
	return k.isCreated
}

func (k *KeyloggerMock) IsListening() bool {
	return k.isListening
}
