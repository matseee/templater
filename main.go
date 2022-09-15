package main

import (
	"github.com/matseee/templater/communication"
	"github.com/matseee/templater/keylistener"
)

func main() {
	// ch := createEventChannel()
	// kl := createKeylistener(ch)

	// g := createGui(ch)
}

func createEventChannel() communication.EventChannel {
	return communication.CreateEventChannel()
}

func createKeylistener(ch communication.EventChannel) keylistener.Keylistener {
	kl, _ := keylistener.CreateKeylistener(ch, &keylistener.KeyloggerLinux{})
	return kl
}

// func createGui(ch communication.EventChannel) gui.Gui {
// 	g := gui.CreateGui(ch)
// 	g.SetSystemtray(&gui.Systray{})
// 	return g
// }
