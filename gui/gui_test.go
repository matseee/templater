package gui

import (
	"testing"

	"github.com/matseee/templater/communication"
)

func createGui() (Gui, communication.EventChannel, *SystemtrayMock) {
	ec := communication.CreateEventChannel()
	st := new(SystemtrayMock)
	gui := CreateGui(ec, st)
	return gui, ec, st
}

func Test_CreateGui_should_return_an_valid_GUI_object(t *testing.T) {
	gui, ec, _ := createGui()
	var inter interface{} = gui

	if _, ok := inter.(Gui); !ok {
		t.Errorf("CreateGUI() should return a Gui object but got other type instead")
	}

	if gui.eventChannel != ec {
		t.Errorf("CreateGUI() should return a Gui object with an EventChannel but got nil instead")
	}
}

func Test_GuiSetSystemtray_should_set_the_given_systemtray_object(t *testing.T) {
	gui, _, st := createGui()

	var inter interface{} = st

	if gui.systemtray != inter {
		t.Errorf("Gui.SetSystemtray() should set the given systemtray object but got nil or another object")
	}
}

func Test_Gui_setupSystemtray_should_setup_the_systemtray_object(t *testing.T) {
	_, _, st := createGui()

	if len(st.menuItems) == 0 {
		t.Errorf("Gui.setupSystemtray() should create the menuitems, but none is available")
	}

	if st.menuItems[0].Title != "Open" {
		t.Errorf("Gui.setupSystemtray() should create the first menuitem with the title 'Open', but got '%s'", st.menuItems[0].Title)
	}

	if st.menuItems[1].Title != "Active" {
		t.Errorf("Gui.setupSystemtray() should create the second menuitem with the title 'Active', but got '%s'", st.menuItems[0].Title)
	}

	if st.menuItems[3].Title != "Quit" {
		t.Errorf("Gui.setupSystemtray() should create the fourth menuitem with the title 'Quit', but got '%s'", st.menuItems[0].Title)
	}

	if st.title != "Templater" {
		t.Errorf("Gui.setupSystemtray() should set the title to 'Templater', but got '%s'", st.title)
	}

	if st.tooltip != "Templater" {
		t.Errorf("Gui.setupSystemtray() should set the tooltip to 'Templater', but got '%s'", st.title)
	}
}

func Test_Gui_Run_should_start_the_systemtray_object(t *testing.T) {
	g, _, st := createGui()

	g.Run()

	if !st.status.IsRunning {
		t.Error("Gui.Run() should start the systemtray, but systemtray.status.IsRunning is false")
	}
}

func Test_Gui_Quit_should_stop_the_systemtray_object(t *testing.T) {
	g, _, st := createGui()

	g.Run()
	g.Quit()

	if st.status.IsRunning {
		t.Error("Gui.Quit() should stop the systemtray, but systemtray.status.IsRunning is true")
	}
}
