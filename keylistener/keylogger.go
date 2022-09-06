package keylistener

type KeyEvent struct {
	Key      string
	Pressed  bool
	Released bool
}

type KeyloggerStatus struct {
	IsCreated   bool
	IsListening bool
}

type Keylogger interface {
	Start() (chan KeyEvent, error)
	Stop()
	Destroy() error
	GetStatus() KeyloggerStatus
}
