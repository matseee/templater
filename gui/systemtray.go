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
		onSystemtrayReady,
		onSystemtrayExit,
	)
}

func onSystemtrayReady() {
	createSystemtrayMenuItems()
	listenToMenuEvents()
}

func onSystemtrayExit() {
	// do nothing
}

func createSystemtrayMenuItems() {
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

func listenToMenuEvents() {
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
	uid := "#" + os.Getenv("SUDO_UID")
	url := getURL()
	go exec.Command("sudo", "-u", uid, "xdg-open", url).Start()
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
	uid := "#" + os.Getenv("SUDO_UID")
	go exec.Command("sudo", "-u", uid, "xdg-open", "https://github.com/matseee/templater").Start()
}

func onQuit() {
	event := templater.CreateEvent()
	event.Type = templater.QuitEvent
	templater.SendEvent(event)

	systray.Quit()
}
