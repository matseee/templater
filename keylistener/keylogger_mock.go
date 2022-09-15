package keylistener

import "time"

func createKeyloggerMock() KeyloggerMock {
	return KeyloggerMock{
		make(chan KeyEvent),
		KeyloggerStatus{},
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

func (k *KeyloggerMock) SendKey(key string) {
	k.outChannel <- KeyEvent{key, false, true}
	time.Sleep(200 * time.Millisecond)
}
