package gui

import (
	"net"
	"testing"
)

func Test_getRandomPort(t *testing.T) {
	randomPort := getRandomPort()

	for i := 0; i < 1000; i++ {
		if randomPort < PORTMIN || randomPort > PORTMAX {
			t.Errorf("Random port should be between %d and %d, but was %d - FAILED", PORTMIN, PORTMAX, randomPort)
		}
	}

	t.Logf("Random ports were all between %d and %d - PASSED", PORTMIN, PORTMAX)
}

func Test_getURL(t *testing.T) {
	port = 22334
	url := getURL()

	if url == "http://127.0.0.1:22334/" {
		t.Log("URL was as expected \"http://127.0.0.1:22334/\" - PASSED")
	} else {
		t.Errorf("URL was not as expected. Got \"%s\" - FAILED", url)
	}
}

func Test_isPortAvailable(t *testing.T) {
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
	} else {
		t.Log("Port 22334 should not be available and isn't available - SUCCESS")
	}

	listener.Close()

	isAvailable = isPortAvailable(22334)
	if isAvailable {
		t.Log("Port 22334 should be available and is available - SUCCESS")
	} else {
		t.Error("Port 22334 should be available, but isn't available - FAILED")
	}
}
