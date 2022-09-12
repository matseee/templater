package gui

type Systemtray interface {
	Run()
	Quit()

	SetTitle(title string)
	SetTooltip(tooltip string)
	SetIcon(iconBytes []byte)
	SetMenuItems(menuItems []MenuItem)
}

type SystemtrayStatus struct {
	IsRunning bool
}

type MenuItemType int64

const (
	MenuItemButton MenuItemType = iota
	MenuItemCheckbox
	MenuItemSeperator
)

type MenuItem struct {
	Type    MenuItemType
	Title   string
	Tooltip string
	Checked bool
	OnClick func(checked bool)
}

func IsMenuItemClickable(m MenuItem) bool {
	return m.Type == MenuItemButton || m.Type == MenuItemCheckbox
}
