package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var portMin, portMax int = 49152, 65535
var port int

func initGUI(channel chan TemplaterEvent) {
	log.Debug(">> START: initGUI()")
	defer log.Debug("<< END: initGUI()")

	port = getRandomPort()

	for createRouter(port) != nil {
		port = getRandomPort()
	}
}

func getRandomPort() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(portMax-portMin) + portMin
}

func getURL() string {
	return "http://localhost:" + strconv.Itoa(port) + "/"
}

func createRouter(port int) error {
	router := gin.Default()
	router.Static("/", "./frontend")
	return router.Run(":" + strconv.Itoa(port))
}
