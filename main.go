package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	// event channel
	channel := make(chan Event)

	go initGUI(channel)
	go listenToChannel(channel)

	// creation of the systray => blocks the runtime
	initSystemtray(channel)
}
