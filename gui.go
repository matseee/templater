package main

import (
	log "github.com/sirupsen/logrus"
)

func initGUI(channel chan Event) {
	log.Debug(">> START: initGUI()")
	defer log.Debug("<< END: initGUI()")
}
