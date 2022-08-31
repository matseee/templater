package keylistener

import (
	"github.com/matseee/templater/templater"
	log "github.com/sirupsen/logrus"
)

func CreateKeylistener(eventChannel templater.EventChannel, keylogger Keylogger) Keylistener {
	return Keylistener{
		eventChannel,
		keylogger,
	}
}

type Keylistener struct {
	eventChannel templater.EventChannel
	keylogger    Keylogger
}

func ListenToChannel(c chan templater.Event, k Keylogger) error {
	go func() {
		for event := range c {
			switch event.Type {
			case templater.StatusEvent:
				if event.ValueBool {
					keyChannel, _ := k.Start()
					go listenToKeylogger(keyChannel)
				} else {
					k.Stop()
				}
			case templater.QuitEvent:
				k.Stop()
			}
		}
	}()

	return nil
}

func listenToKeylogger(channel chan KeyEvent) {
	// TODO
	for event := range channel {
		if event.Released {
			log.Println("listenToKeylogger() ==> ", event.Key)
		}
	}
}
