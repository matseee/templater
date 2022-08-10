package gui

import (
	"os"
	"os/exec"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/matseee/templater/templater"
)

var (
	menuItemOpen, menuItemStatus, menuItemGithub, menuItemQuit *systray.MenuItem
)

func InitSystemtray(channel chan templater.Event) {
	systray.Run(
		onReady,
		onExit,
	)
}

func onReady() {
	createMenuItems()
	handleMenuItemEvents()
}

func onExit() {
	// do nothing
}

func createMenuItems() {
	systray.SetTitle("Templater")
	systray.SetTooltip("Templater")
	systray.SetTemplateIcon(icon.Data, icon.Data)

	menuItemOpen = systray.AddMenuItem("Open", "Open")
	menuItemStatus = systray.AddMenuItemCheckbox("Active", "Check Me", false)
	systray.AddSeparator()

	menuItemGithub = systray.AddMenuItem("Homepage", "Visit homepage")
	systray.AddSeparator()

	menuItemQuit = systray.AddMenuItem("Quit", "Quit the whole app")
}

func handleMenuItemEvents() {
	for {
		select {
		case <-menuItemOpen.ClickedCh:
			onOpenGUI()
		case <-menuItemStatus.ClickedCh:
			onToggleTemplaterActive()
		case <-menuItemGithub.ClickedCh:
			onOpenWebpage()
		case <-menuItemQuit.ClickedCh:
			onQuit()
		}
	}
}

func onOpenGUI() {
	go exec.Command("sudo", "-u", getSudoUID(), "xdg-open", getURL()).Start()
}

func onToggleTemplaterActive() {
	event := templater.CreateEvent()
	event.Type = templater.StatusEvent
	event.ValueBool = !menuItemStatus.Checked()
	templater.SendEvent(event)

	if menuItemStatus.Checked() {
		menuItemStatus.Uncheck()
	} else {
		menuItemStatus.Check()
	}
}

func onOpenWebpage() {
	go exec.Command("sudo", "-u", getSudoUID(), "xdg-open", "https://github.com/matseee/templater").Start()
}

func onQuit() {
	event := templater.CreateEvent()
	event.Type = templater.QuitEvent
	templater.SendEvent(event)

	systray.Quit()
}

func getSudoUID() string {
	return "#" + os.Getenv("SUDO_UID")
}
