package main

import (
	"github.com/matseee/templater/gui"
	"github.com/matseee/templater/keylistener"
	"github.com/matseee/templater/templater"
	log "github.com/sirupsen/logrus"
)

var (
	channel   chan templater.Event
	keylogger keylistener.KeyloggerLinux
)

func main() {
	log.SetLevel(log.DebugLevel)
	createApplication()
}

func createApplication() {
	channel = templater.CreateChannel()
	keylogger = keylistener.KeyloggerLinux{}

	go keylistener.ListenToChannel(channel, &keylogger)
	go gui.InitSystemtray(channel)

	gui.InitGUI(channel)
}
