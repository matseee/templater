package gui

import (
	"github.com/getlantern/systray/example/icon"
	"github.com/matseee/templater/communication"
)

func CreateGui(eventChannel communication.EventChannel, systemtray Systemtray) Gui {
	g := *new(Gui)
	g.eventChannel = eventChannel
	g.systemtray = systemtray
	g.setupSystemtray()
	return g
}

type Gui struct {
	eventChannel communication.EventChannel
	systemtray   Systemtray
}

func (g *Gui) Run() {
	g.systemtray.Run()
}

func (g *Gui) Quit() {
	g.systemtray.Quit()
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
				g.eventChannel.Send(communication.CreateOpenSettingsEvent())
			},
		},
		{
			Type:    MenuItemCheckbox,
			Title:   "Active",
			Tooltip: "Activate templater",
			Checked: false,
			OnClick: func(checked bool) {
				g.eventChannel.Send(communication.CreateStatusEvent(checked))
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
				g.eventChannel.Send(communication.CreateQuitEvent())
			},
		},
	}
}
