package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	// event channel
	channel := make(chan TemplaterEvent)

	go listenToChannel(channel)
	go initSystemtray(channel)

	initGUI(channel)
}
