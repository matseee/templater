package keylistener

import "github.com/MarinX/keylogger"

type KeyloggerLinux struct {
	outChannel chan KeyEvent
	status     KeyloggerStatus

	keylogger *keylogger.KeyLogger
	inChannel chan keylogger.InputEvent
}

func (k *KeyloggerLinux) Start() (chan KeyEvent, error) {
	if k.status.IsListening {
		return k.outChannel, nil
	}

	if !k.status.IsCreated {
		if err := k.create(); err != nil {
			return nil, err
		}
	}

	k.outChannel = make(chan KeyEvent)
	k.status.IsListening = true
	return k.outChannel, nil
}

func (k *KeyloggerLinux) Stop() {
	if !k.status.IsListening {
		return
	}

	close(k.outChannel)
	k.outChannel = nil

	k.status.IsListening = false
}

func (k *KeyloggerLinux) Destroy() error {
	if !k.status.IsCreated {
		return nil
	}

	k.keylogger.Close()

	k.status.IsListening = false
	k.status.IsCreated = false
	return nil
}

func (k *KeyloggerLinux) GetStatus() KeyloggerStatus {
	return k.status
}

func (k *KeyloggerLinux) create() error {
	var err error

	k.keylogger, err = keylogger.New("/dev/input/event19")

	if err != nil {
		return err
	}

	k.inChannel = k.keylogger.Read()
	go k.listen()

	k.status.IsCreated = true
	return nil
}

func (k *KeyloggerLinux) listen() {
	for inputEvent := range k.inChannel {
		if k.status.IsListening {
			switch inputEvent.Type {
			case keylogger.EvKey:
				var outEvent KeyEvent
				outEvent.Key = inputEvent.KeyString()

				if inputEvent.KeyRelease() {
					outEvent.Released = true
				}

				if inputEvent.KeyPress() {
					outEvent.Pressed = true
				}

				k.outChannel <- outEvent
			}
		}
	}
}
