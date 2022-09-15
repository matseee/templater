package templater

import (
	"testing"

	"github.com/matseee/templater/communication"
	"github.com/matseee/templater/gui"
	"github.com/matseee/templater/keylistener"
)

type TestObject struct {
	Templater    Templater
	EventChannel communication.EventChannel
	Keylistener  keylistener.Keylistener
	Keylogger    *keylistener.KeyloggerMock
	Gui          gui.Gui
	Systemtray   *gui.SystemtrayMock
}

func createTemplater() TestObject {
	ec := communication.CreateEventChannel()
	klog := keylistener.KeyloggerMock{}
	systray := gui.SystemtrayMock{}

	kl, _ := keylistener.CreateKeylistener(ec, &klog)
	g := gui.CreateGui(ec, &systray)

	tp := CreateTemplater(ec, kl, g)

	return TestObject{
		tp,
		ec,
		kl,
		&klog,
		g,
		&systray,
	}
}

func Test_CreateTemplater_should_return_a_templater_object(t *testing.T) {
	to := createTemplater()

	if to.Templater.EventChannel != to.EventChannel {
		t.Errorf("CreateTemplater() should return a templater object with the given EventChannel, but returned an invalid object")
	}

	if to.Templater.Keylistener != to.Keylistener {
		t.Errorf("CreateTemplater() should return a templater object with the given Keylistener, but returned an invalid object")
	}

	if to.Templater.Gui != to.Gui {
		t.Errorf("CreateTemplater() should return a templater object with the given Gui, but returned an invalid object")
	}
}

func Test_CreateTemplater_should_start_the_keylistener(t *testing.T) {
	to := createTemplater()

	if !to.Keylogger.GetStatus().IsCreated {
		t.Error("Templater.Run() should start the keylistener and set the keylistener.keylogger.IsCreated to true, but it is false")
	}

	if !to.Keylogger.GetStatus().IsListening {
		t.Error("Templater.Run() should start the keylistener and set the keylistener.keylogger.IsListening to true, but it is false")
	}
}

func Test_Templater_Run_should_start_the_gui(t *testing.T) {
	to := createTemplater()

	to.Templater.Run()

	if !to.Systemtray.GetStatus().IsRunning {
		t.Error("Templater.Run() should start the Gui and set the Gui.Systemtray.status.IsRunning to true, but it is false")
	}
}
