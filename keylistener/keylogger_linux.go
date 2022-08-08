package keylistener

import "github.com/MarinX/keylogger"

type KeyloggerLinux struct {
	outChannel  chan KeyEvent
	isCreated   bool
	isListening bool

	// specfic attributes
	keylogger *keylogger.KeyLogger
	inChannel chan keylogger.InputEvent
}

func (k *KeyloggerLinux) Create() error {
	var err error

	k.keylogger, err = keylogger.New("/dev/input/event19")

	if err != nil {
		return err
	}

	k.inChannel = k.keylogger.Read()
	go k.listen()

	k.isCreated = true
	return nil
}

func (k *KeyloggerLinux) Destroy() error {
	if !k.isCreated {
		return nil
	}

	k.keylogger.Close()

	k.isListening = false
	k.isCreated = false
	return nil
}

func (k *KeyloggerLinux) Start() chan KeyEvent {
	if k.isListening {
		return k.outChannel
	}

	k.outChannel = make(chan KeyEvent)

	k.isListening = true
	return k.outChannel
}

func (k *KeyloggerLinux) Stop() {
	if !k.isListening {
		return
	}

	close(k.outChannel)
	k.outChannel = nil

	k.isListening = false
}

func (k *KeyloggerLinux) IsCreated() bool {
	return k.isCreated
}

func (k *KeyloggerLinux) IsListening() bool {
	return k.isListening
}

func (k *KeyloggerLinux) listen() {
	for inputEvent := range k.inChannel {
		if k.isListening {
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
