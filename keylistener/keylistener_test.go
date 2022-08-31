package keylistener

import (
	"testing"
	"time"

	"github.com/matseee/templater/templater"
)

func createKeyloggerMock() Keylogger {
	return &KeyloggerMock{}
}

func createChannelAndMock() (templater.EventChannel, Keylogger) {
	return templater.CreateEventChannel(), createKeyloggerMock()
}

func Test_CreateKeylistener_should_create_an_initial_keylistener_object(t *testing.T) {
	eventChannel, keylogger := createChannelAndMock()

	var testKeylistener = Keylistener{
		eventChannel,
		keylogger,
	}

	keylistener := CreateKeylistener(eventChannel, keylogger)

	if testKeylistener != keylistener {
		t.Error("CreateKeylistener did not create an normal keylistener object")
	}
}

func Test_Keylistener_should_listen_on_keyevents(t *testing.T) {
	eventChannel, keylogger := createChannelAndMock()
	keylistener := CreateKeylistener(eventChannel, keylogger)

	if keylogger.GetStatus().IsCreated()
}

func Test_ListenToChannel(t *testing.T) {
	var c chan templater.Event = make(chan templater.Event)
	var k KeyloggerMock = KeyloggerMock{}

	err := ListenToChannel(c, &k)

	if err != nil {
		str := err.Error() + " - FAILED"
		t.Error(str)
	}

	status := k.GetStatus()

	if status.IsListening {
		t.Error("Keylogger should not listen after creation - FAILED")
	} else {
		t.Log("Keylogger is not listening - PASSED")
	}

	event := templater.CreateEvent()
	event.Type = templater.StatusEvent
	event.ValueBool = true
	c <- event

	status = k.GetStatus()

	if !status.IsListening {
		t.Error("Keylogger should listen after start event - FAILED")
	} else {
		t.Log("Keylogger is listening after start event - PASSED")
	}

	event = templater.CreateEvent()
	event.Type = templater.StatusEvent
	event.ValueBool = false
	c <- event
	time.Sleep(100 * time.Millisecond)

	status = k.GetStatus()

	if status.IsListening {
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

	status = k.GetStatus()

	if status.IsCreated {
		t.Error("Keylogger should not be created after destroy() - FAILED")
	} else {
		t.Log("Keylogger is destroyed - PASSED")
	}
}

type KeyloggerMock struct {
	outChannel chan KeyEvent
	status     KeyloggerStatus
}

func (k *KeyloggerMock) Start() (chan KeyEvent, error) {
	if k.status.IsListening {
		return k.outChannel, nil
	}

	k.outChannel = make(chan KeyEvent)
	k.status.IsCreated = true
	k.status.IsListening = true
	return k.outChannel, nil
}

func (k *KeyloggerMock) Stop() {
	if !k.status.IsListening {
		return
	}

	close(k.outChannel)
	k.outChannel = nil

	k.status.IsListening = false
}

func (k *KeyloggerMock) Destroy() error {
	k.status.IsCreated = false
	k.status.IsListening = false
	return nil
}

func (k *KeyloggerMock) GetStatus() KeyloggerStatus {
	return k.status
}
