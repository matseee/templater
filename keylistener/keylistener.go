package keylistener

import (
	"fmt"

	"github.com/matseee/templater/templater"
)

func CreateKeylistener(eventChannel templater.EventChannel, keylogger Keylogger) (Keylistener, error) {
	k := Keylistener{
		eventChannel,
		keylogger,
		&KeylistenerStatus{},
	}

	return k, k.listen()
}

type Keylistener struct {
	eventChannel templater.EventChannel
	keylogger    Keylogger
	status       *KeylistenerStatus
}

func (k *Keylistener) Activate() {
	k.setActive(true)
}

func (k *Keylistener) Deactivate() {
	k.setActive(false)
}

func (k *Keylistener) setActive(active bool) {
	k.status.IsActive = active
}

func (k *Keylistener) GetStatus() KeylistenerStatus {
	return *k.status
}

func (k *Keylistener) listen() error {
	keyEvantChannel, err := k.keylogger.Start()

	if err != nil {
		fmt.Print(err)
		return err
	}

	go k.listenToKeyEventChannel(keyEvantChannel)
	return nil
}

func (k *Keylistener) listenToKeyEventChannel(keyEvantChannel chan KeyEvent) {
	for keyEvent := range keyEvantChannel {
		if k.isKeylistenerActive() && k.isKeyEventActive(keyEvent) {
			k.increaseKeyEventCounter()
			k.sendKeylistenerEventToEventChannel(keyEvent.Key)
		}
	}
}

func (k *Keylistener) isKeylistenerActive() bool {
	return k.status.IsActive
}

func (k *Keylistener) isKeyEventActive(keyEvent KeyEvent) bool {
	return keyEvent.Released
}

func (k *Keylistener) increaseKeyEventCounter() {
	k.status.KeyEventCounter += 1
}

func (k *Keylistener) sendKeylistenerEventToEventChannel(key string) {
	event := templater.CreateKeylistenerEvent(key)
	k.eventChannel.Send(event)
}

type KeylistenerStatus struct {
	IsActive        bool
	KeyEventCounter uint64
}
