package templater

import (
	"github.com/matseee/templater/communication"
	"github.com/matseee/templater/gui"
	"github.com/matseee/templater/keylistener"
)

type TemplaterStatus struct {
	IsRunning bool
}

func CreateTemplater(ec communication.EventChannel, kl keylistener.Keylistener, g gui.Gui) Templater {
	return Templater{
		ec,
		kl,
		g,
	}
}

type Templater struct {
	EventChannel communication.EventChannel
	Keylistener  keylistener.Keylistener
	Gui          gui.Gui
}

func (t *Templater) Run() {
	t.Gui.Run()
}
