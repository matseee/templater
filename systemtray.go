package main

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	log "github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
)

func initSystemtray(channel chan Event) {
	systray.Run(
		func() {
			log.Debug(">> START: initSystemtray() -> onReady()")
			defer log.Debug("<< END: initSystemtray() -> onReady()")

			systray.SetTitle("Templater")
			systray.SetTooltip("Templater")
			systray.SetTemplateIcon(icon.Data, icon.Data)

			menuItemActive := systray.AddMenuItemCheckbox("Active", "Check Me", true)
			systray.AddSeparator()
			menuItemGithub := systray.AddMenuItem("Homepage", "Visit homepage")
			systray.AddSeparator()
			menuItemQuit := systray.AddMenuItem("Quit", "Quit the whole app")

			go func() {
				for {
					select {
					case <-menuItemActive.ClickedCh:
						log.Debug("menuItemActive.ClickedCh")

						if menuItemActive.Checked() {
							menuItemActive.Uncheck()
							channel <- KeyloggerDeactivate
						} else {
							menuItemActive.Check()
							channel <- KeyloggerActivate
						}
					case <-menuItemGithub.ClickedCh:
						log.Debug("menuItemGithub.ClickedCh")

						open.Run("https://github.com/matseee/templater")
					case <-menuItemQuit.ClickedCh:
						log.Debug("menuItemQuit.ClickedCh")

						channel <- Quit
						systray.Quit()
					}
				}

			}()

			channel <- KeyloggerActivate
		},
		func() {
			log.Debug(">> START: initSystemtray() -> onExit()")
			defer log.Debug("<< END: initSystemtray() -> onExit()")
		},
	)
}
