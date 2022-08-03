package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"
)

var application fyne.App

func initGUI(channel chan TemplaterEvent) {
	log.Debug(">> START: initGUI()")
	defer log.Debug("<< END: initGUI()")

	application = app.New()
	application.Settings()

	for {
		for event := range channel {
			log.Debug(event)
			switch event.Type {
			case OpenSettings:
				createWindow()
			case Quit:
				return
			}
		}
	}
}

func createWindow() {
	log.Debug(">> START: createWindow()")
	defer log.Debug("<< END: createWindow()")

	w := application.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
