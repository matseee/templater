package gui

import (
	"github.com/matseee/templater/templater"
)

func CreateGui(eventChannel templater.EventChannel) Gui {
	gui := *new(Gui)
	gui.eventChannel = eventChannel
	return gui
}

type Gui struct {
	eventChannel templater.EventChannel

	systemtray Systemtray
	api        Api
}

func (g *Gui) SetSystemtray(systemtray Systemtray) {
	g.systemtray = systemtray
}

func (g *Gui) SetApi(api Api) {
	g.api = api
}

// const (
// 	PORTMIN, PORTMAX int = 49152, 65535
// )

// var (
// 	port int
// )

// func InitGUI(channel chan templater.Event) error {
// 	port = getRandomPort()

// 	for !isPortAvailable(port) {
// 		port = getRandomPort()
// 	}

// 	return createRouter(port)
// }

// func getRandomPort() int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(PORTMAX-PORTMIN) + PORTMIN
// }

// func getURL() string {
// 	return "http://127.0.0.1:" + strconv.Itoa(port) + "/"
// }

// func isPortAvailable(port int) bool {
// 	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))

// 	if err != nil {
// 		return false
// 	}

// 	l.Close()
// 	return true
// }

// func createRouter(port int) error {
// 	router := gin.Default()
// 	router.Static("/", "./frontend")
// 	return router.Run(":" + strconv.Itoa(port))
// }
