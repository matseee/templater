package main

import (
	"os"
	"os/exec"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	log "github.com/sirupsen/logrus"
)

func initSystemtray(channel chan TemplaterEvent) {
	systray.Run(
		func() {
			log.Debug(">> START: initSystemtray() -> onReady()")
			defer log.Debug("<< END: initSystemtray() -> onReady()")

			systray.SetTitle("Templater")
			systray.SetTooltip("Templater")
			systray.SetTemplateIcon(icon.Data, icon.Data)

			menuItemStatus := systray.AddMenuItemCheckbox("Active", "Check Me", true)
			menuItemSettings := systray.AddMenuItem("Settings", "Settings")

			systray.AddSeparator()

			menuItemGithub := systray.AddMenuItem("Homepage", "Visit homepage")

			systray.AddSeparator()

			menuItemQuit := systray.AddMenuItem("Quit", "Quit the whole app")

			go func() {
				for {
					select {
					case <-menuItemStatus.ClickedCh:
						log.Debug("menuItemStatus.ClickedCh")

						event := CreateEvent()
						event.Type = TemplaterStatus
						event.ValueBool = !menuItemStatus.Checked()
						channel <- event

						if menuItemStatus.Checked() {
							menuItemStatus.Uncheck()
						} else {
							menuItemStatus.Check()
						}

					case <-menuItemSettings.ClickedCh:
						log.Debug("menuItemSettings.ClickedCh")

						uid := "#" + os.Getenv("SUDO_UID")
						url := getURL()
						go exec.Command("sudo", "-u", uid, "xdg-open", url).Start()

					case <-menuItemGithub.ClickedCh:
						log.Debug("menuItemGithub.ClickedCh")

						uid := "#" + os.Getenv("SUDO_UID")
						go exec.Command("sudo", "-u", uid, "xdg-open", "https://github.com/matseee/templater").Start()

					case <-menuItemQuit.ClickedCh:
						log.Debug("menuItemQuit.ClickedCh")

						event := CreateEvent()
						event.Type = Quit
						channel <- event

						systray.Quit()
					}
				}
			}()

			event := CreateEvent()
			event.Type = TemplaterStatus
			event.ValueBool = true

			channel <- event
		},
		func() {
			log.Debug(">> START: initSystemtray() -> onExit()")
			defer log.Debug("<< END: initSystemtray() -> onExit()")
		},
	)
}
