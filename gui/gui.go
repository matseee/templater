package gui

import (
	"github.com/getlantern/systray/example/icon"
	"github.com/matseee/templater/communication"
)

func CreateGui(eventChannel communication.EventChannel) Gui {
	gui := *new(Gui)
	gui.eventChannel = eventChannel
	return gui
}

type Gui struct {
	eventChannel communication.EventChannel
	systemtray   Systemtray
}

func (g *Gui) SetSystemtray(systemtray Systemtray) {
	g.systemtray = systemtray
	g.setupSystemtray()
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
