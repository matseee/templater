package gui

type SystemtrayMock struct {
	status SystemtrayStatus

	title     string
	tooltip   string
	iconBytes []byte
	menuItems []MenuItem
}

func (s *SystemtrayMock) Run() {
	s.status.IsRunning = true
}

func (s *SystemtrayMock) Quit() {
	s.status.IsRunning = false
}

func (s *SystemtrayMock) SetTitle(title string) {
	s.title = title
}

func (s *SystemtrayMock) SetTooltip(tooltip string) {
	s.tooltip = tooltip
}

func (s *SystemtrayMock) SetIcon(iconBytes []byte) {
	s.iconBytes = iconBytes
}

func (s *SystemtrayMock) SetMenuItems(menuItems []MenuItem) {
	s.menuItems = menuItems
}

func (s *SystemtrayMock) CallItemClickCallback() {
	for _, menuItem := range s.menuItems {
		menuItem.OnClick(false)
	}
}

func (s *SystemtrayMock) GetStatus() SystemtrayStatus {
	return s.status
}
