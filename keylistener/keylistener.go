package keylistener

import (
	"github.com/matseee/templater/templater"
	log "github.com/sirupsen/logrus"
)

func ListenToChannel(c chan templater.Event, k Keylogger) error {
	err := k.Create()

	if err != nil {
		return err
	}

	go func() {
		for event := range c {
			switch event.Type {
			case templater.StatusEvent:
				if event.ValueBool {
					keyChannel := k.Start()
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
