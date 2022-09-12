package gui

import (
	"github.com/getlantern/systray/example/icon"
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
	g.setupSystemtray()
}

func (g *Gui) SetApi(api Api) {
	g.api = api
}

func (g *Gui) setupSystemtray() {
	g.systemtray.SetTitle("Templater")
	g.systemtray.SetTooltip("Templater")
	g.systemtray.SetIcon(icon.Data)
	g.systemtray.SetMenuItems(g.createMenuItems())
}

func (g *Gui) createMenuItems() []MenuItem {
	return []MenuItem{
		{
			Type:    MenuItemButton,
			Title:   "Open",
			Tooltip: "Open settings",
			OnClick: func(checked bool) {
				g.eventChannel.Send(templater.CreateOpenSettingsEvent())
			},
		},
		{
			Type:    MenuItemCheckbox,
			Title:   "Active",
			Tooltip: "Activate templater",
			Checked: false,
			OnClick: func(checked bool) {
				g.eventChannel.Send(templater.CreateStatusEvent(checked))
			},
		},
		{
			Type: MenuItemSeperator,
		},
		{
			Type:    MenuItemButton,
			Title:   "Quit",
			Tooltip: "Quit the application",
			OnClick: func(checked bool) {
				g.eventChannel.Send(templater.CreateQuitEvent())
			},
		},
	}
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
