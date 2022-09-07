package gui

import (
	"github.com/getlantern/systray"
)

type Systray struct {
	status SystemtrayStatus

	title     string
	tooltip   string
	iconBytes []byte
	menuItems []MenuItem
}

func (s *Systray) SetTitle(title string) {
	s.title = title
}

func (s *Systray) SetTooltip(tooltip string) {
	s.tooltip = tooltip
}

func (s *Systray) SetIcon(iconBytes []byte) {
	s.iconBytes = iconBytes
}

func (s *Systray) SetMenuItems(menuItems []MenuItem) {
	s.menuItems = menuItems
}

func (s *Systray) Run() {
	systray.Run(
		s.onReady,
		s.onExit,
	)
}

func (s *Systray) onReady() {
}

func (s *Systray) onExit() {
}

func (s *Systray) Quit() {
}

// var (
// 	menuItemOpen, menuItemStatus, menuItemGithub, menuItemQuit *systray.MenuItem
// )

// func InitSystemtray(channel chan templater.Event) {

// }

// func onReady() {
// 	createMenuItems()
// 	handleMenuItemEvents()
// }

// func createMenuItems() {
// 	systray.SetTitle("Templater")
// 	systray.SetTooltip("Templater")
// 	systray.SetTemplateIcon(icon.Data, icon.Data)

// 	menuItemOpen = systray.AddMenuItem("Open", "Open")
// 	menuItemStatus = systray.AddMenuItemCheckbox("Active", "Check Me", false)
// 	systray.AddSeparator()

// 	menuItemGithub = systray.AddMenuItem("Homepage", "Visit homepage")
// 	systray.AddSeparator()

// 	menuItemQuit = systray.AddMenuItem("Quit", "Quit the whole app")
// }

// func handleMenuItemEvents() {
// 	for {
// 		select {
// 		case <-menuItemOpen.ClickedCh:
// 			onOpenGUI()
// 		case <-menuItemStatus.ClickedCh:
// 			onToggleTemplaterActive()
// 		case <-menuItemGithub.ClickedCh:
// 			onOpenWebpage()
// 		case <-menuItemQuit.ClickedCh:
// 			onQuit()
// 		}
// 	}
// }

// func onOpenGUI() {
// 	go exec.Command("sudo", "-u", getSudoUID(), "xdg-open", getURL()).Start()
// }

// func onToggleTemplaterActive() {
// 	event := templater.CreateEvent()
// 	event.Type = templater.StatusEvent
// 	event.ValueBool = !menuItemStatus.Checked()
// 	templater.SendEvent(event)

// 	if menuItemStatus.Checked() {
// 		menuItemStatus.Uncheck()
// 	} else {
// 		menuItemStatus.Check()
// 	}
// }

// func onOpenWebpage() {
// 	go exec.Command("sudo", "-u", getSudoUID(), "xdg-open", "https://github.com/matseee/templater").Start()
// }

// func onQuit() {
// 	event := templater.CreateEvent()
// 	event.Type = templater.QuitEvent
// 	templater.SendEvent(event)

// 	systray.Quit()
// }

// func getSudoUID() string {
// 	return "#" + os.Getenv("SUDO_UID")
// }
