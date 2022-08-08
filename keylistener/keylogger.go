package keylistener

type KeyEvent struct {
	Key      string
	Pressed  bool
	Released bool
}

type Keylogger interface {
	Create() error
	Destroy() error
	Start() chan KeyEvent
	Stop()
	IsCreated() bool
	IsListening() bool
}
