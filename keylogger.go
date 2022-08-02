package main

import (
	"github.com/MarinX/keylogger"
	log "github.com/sirupsen/logrus"
)

func listenToChannel(channel chan Event) {
	log.Debug(">> START: listenToChannel()")
	defer log.Debug("<< END: listenToChannel()")

	var k *keylogger.KeyLogger

	for event := range channel {
		switch event {
		case KeyloggerActivate:
			k = createKeylogger()
		case KeyloggerDeactivate:
			k.Close()
		case Quit:
			k.Close()
		}
	}
}

func createKeylogger() *keylogger.KeyLogger {
	log.Debug(">> START: createKeylogger()")
	defer log.Debug("<< END: createKeylogger()")

	k, err := keylogger.New("/dev/input/event19")

	if err != nil {
		log.Error(err)
		return nil
	}

	// create channel
	channel := k.Read()
	go listenToKeylogger(channel)

	return k
}

func listenToKeylogger(channel chan keylogger.InputEvent) {
	log.Debug(">> START: listenToKeylogger()")
	defer log.Debug("<< END: listenToKeylogger()")

	for event := range channel {
		switch event.Type {
		case keylogger.EvKey:
			// if the state of key is released
			if event.KeyRelease() {
				log.Println("listenToKeylogger() ==> ", event.KeyString())
			}
		}
	}
}
