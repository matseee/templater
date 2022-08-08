package gui

import (
	"os"
	"os/exec"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/matseee/templater/templater"
	log "github.com/sirupsen/logrus"
)

func InitSystemtray(channel chan templater.Event) {
	systray.Run(
		func() {
			log.Debug(">> START: initSystemtray() -> onReady()")
			defer log.Debug("<< END: initSystemtray() -> onReady()")

			systray.SetTitle("Templater")
			systray.SetTooltip("Templater")
			systray.SetTemplateIcon(icon.Data, icon.Data)

			menuItemOpen := systray.AddMenuItem("Open", "Open")
			menuItemStatus := systray.AddMenuItemCheckbox("Active", "Check Me", true)
			systray.AddSeparator()

			menuItemGithub := systray.AddMenuItem("Homepage", "Visit homepage")
			systray.AddSeparator()

			menuItemQuit := systray.AddMenuItem("Quit", "Quit the whole app")

			go func() {
				for {
					select {
					case <-menuItemOpen.ClickedCh:
						log.Debug("menuItemOpen.ClickedCh")

						uid := "#" + os.Getenv("SUDO_UID")
						url := getURL()
						go exec.Command("sudo", "-u", uid, "xdg-open", url).Start()

					case <-menuItemStatus.ClickedCh:
						log.Debug("menuItemStatus.ClickedCh")

						event := templater.CreateEvent()
						event.Type = templater.StatusEvent
						event.ValueBool = !menuItemStatus.Checked()
						templater.SendEvent(event)

						if menuItemStatus.Checked() {
							menuItemStatus.Uncheck()
						} else {
							menuItemStatus.Check()
						}

					case <-menuItemGithub.ClickedCh:
						log.Debug("menuItemGithub.ClickedCh")

						uid := "#" + os.Getenv("SUDO_UID")
						go exec.Command("sudo", "-u", uid, "xdg-open", "https://github.com/matseee/templater").Start()

					case <-menuItemQuit.ClickedCh:
						log.Debug("menuItemQuit.ClickedCh")

						event := templater.CreateEvent()
						event.Type = templater.QuitEvent
						templater.SendEvent(event)

						systray.Quit()
					}
				}
			}()

			event := templater.CreateEvent()
			event.Type = templater.StatusEvent
			event.ValueBool = true
			templater.SendEvent(event)
		},
		func() {
			log.Debug(">> START: initSystemtray() -> onExit()")
			defer log.Debug("<< END: initSystemtray() -> onExit()")
		},
	)
}
