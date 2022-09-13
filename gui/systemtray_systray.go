package gui

import (
	"github.com/getlantern/systray"
)

type ClickableMenuItem struct {
	systrayMenuItem *systray.MenuItem
	menuItem        MenuItem
}

type Systray struct {
	status SystemtrayStatus

	title     string
	tooltip   string
	iconBytes []byte
	menuItems []MenuItem

	clickableMenuItems []ClickableMenuItem
	systrayMenuItems   []*systray.MenuItem
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
	go s.listenToMenuItemClicks()

	s.status.IsRunning = true
}

func (s *Systray) listenToMenuItemClicks() {
	menuItemClickedChannel := make(chan *ClickableMenuItem)

	for _, clickableMenuItem := range s.clickableMenuItems {
		go func(c chan struct{}, _clickableMenuItem *ClickableMenuItem) {
			for range c {
				menuItemClickedChannel <- _clickableMenuItem
			}
		}(clickableMenuItem.systrayMenuItem.ClickedCh, &clickableMenuItem)
	}

	for clickedMenuItem := range menuItemClickedChannel {
		clickedMenuItem.menuItem.OnClick(
			!clickedMenuItem.systrayMenuItem.Checked(),
		)
	}
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
	s.clickableMenuItems = make([]ClickableMenuItem, 0)
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
	callableMenuItem := ClickableMenuItem{
		systrayMenuItem: s.systrayMenuItems[len(s.systrayMenuItems)-1],
		menuItem:        menuItem,
	}
	s.clickableMenuItems = append(s.clickableMenuItems, callableMenuItem)
}

func (s *Systray) onExit() {
	s.clearMenuItems()
	s.status.IsRunning = false
}

func (s *Systray) Quit() {
	systray.Quit()
}
