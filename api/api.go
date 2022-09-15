package api

import (
	"math/rand"
	"net"
	"strconv"
	"time"
)

const (
	API_PORTMIN, API_PORTMAX int = 49152, 65535
)

type Api interface {
	Run()
	Quit()
}

type ApiCallbacks struct{}

func getRandomPort() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(API_PORTMAX-API_PORTMIN) + API_PORTMIN
}

func isPortAvailable(port int) bool {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))

	if err != nil {
		return false
	}

	l.Close()
	return true
}
