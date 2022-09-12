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

	systrayChannels   []chan struct{}
	callableMenuItems []MenuItem
	systrayMenuItems  []*systray.MenuItem
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
	s.createSystayIconTitleAndTooltip()
	s.createMenuItems()

	// for {
	// 	reflect.Select(s.systrayChannels)
	// }

	s.status.IsRunning = true
}

func (s *Systray) createSystayIconTitleAndTooltip() {
	systray.SetIcon(s.iconBytes)
	systray.SetTitle(s.title)
	systray.SetTooltip(s.tooltip)
}

func (s *Systray) createMenuItems() {
	s.clearMenuItems()

	for _, menuItem := range s.menuItems {
		s.createMenuItem(menuItem)

		if IsMenuItemClickable(menuItem) {
			s.addClickableMenuItem(menuItem)
		}
	}
}

func (s *Systray) clearMenuItems() {
	s.systrayChannels = make([]chan struct{}, 0)
	s.callableMenuItems = make([]MenuItem, 0)
	s.systrayMenuItems = make([]*systray.MenuItem, 0)
}

func (s *Systray) createMenuItem(menuItem MenuItem) {
	switch menuItem.Type {
	case MenuItemButton:
		s.systrayMenuItems = append(s.systrayMenuItems, systray.AddMenuItem(menuItem.Title, menuItem.Tooltip))
	case MenuItemCheckbox:
		s.systrayMenuItems = append(s.systrayMenuItems, systray.AddMenuItemCheckbox(menuItem.Title, menuItem.Tooltip, false))
	case MenuItemSeperator:
		systray.AddSeparator()
	}
}

func (s *Systray) addClickableMenuItem(menuItem MenuItem) {
	s.systrayChannels = append(s.systrayChannels, s.systrayMenuItems[len(s.systrayMenuItems)-1].ClickedCh)
	s.callableMenuItems = append(s.callableMenuItems, menuItem)
}

func (s *Systray) onExit() {
	s.clearMenuItems()
	s.status.IsRunning = false
}

func (s *Systray) Quit() {
	systray.Quit()
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
