package api

import (
	"net"
	"testing"
)

type ApiMock struct {
}

func (a *ApiMock) Run() {
}

func (a *ApiMock) Quit() {
}

func Test_getRandomPort_should_return_a_port_between_49152_and_65535(t *testing.T) {
	randomPort := getRandomPort()

	for i := 0; i < 1000; i++ {
		if randomPort < API_PORTMIN || randomPort > API_PORTMAX {
			t.Errorf("Random port should return a port between %d and %d, but returned %d", API_PORTMIN, API_PORTMAX, randomPort)
		}
	}
}

func Test_isPortAvailable_should_check_if_the_port_22334_is_available(t *testing.T) {
	port := "22334"
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		t.Log("Can't open port 22334 for test, skipping test")
		t.SkipNow()
		return
	}

	isAvailable := isPortAvailable(22334)
	if isAvailable {
		t.Error("Port 22334 should not be available, but is available - FAILED")
		return
	}

	listener.Close()

	isAvailable = isPortAvailable(22334)
	if !isAvailable {
		t.Error("Port 22334 should be available, but isn't available - FAILED")
	}
}
