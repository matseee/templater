package main

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
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

						if err := open.Run(getURL()); err != nil {
							log.Error(err)
						}

					case <-menuItemGithub.ClickedCh:
						log.Debug("menuItemGithub.ClickedCh")
						open.Run("https://github.com/matseee/templater")

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
