package gui

import (
	"testing"

	"github.com/matseee/templater/templater"
)

func createGui() (Gui, templater.EventChannel) {
	ec := templater.CreateEventChannel()
	gui := CreateGui(ec)
	return gui, ec
}

func Test_CreateGui_should_return_an_valid_GUI_object(t *testing.T) {
	gui, ec := createGui()
	var inter interface{} = gui

	if _, ok := inter.(Gui); !ok {
		t.Errorf("CreateGUI() should return a Gui object but got other type instead")
	}

	if gui.eventChannel != ec {
		t.Errorf("CreateGUI() should return a Gui object with an EventChannel but got nil instead")
	}
}

func Test_GuiSetSystemtray_should_set_the_given_systemtray_object(t *testing.T) {
	gui, _ := createGui()
	st := new(SystemtrayMock)
	var inter interface{} = st

	gui.SetSystemtray(st)

	if gui.systemtray != inter {
		t.Errorf("Gui.SetSystemtray() should set the given systemtray object but got nil or another object")
	}
}

func Test_Gui_setupSystemtray_should_setup_the_systemtray_object(t *testing.T) {
	gui, _ := createGui()
	st := new(SystemtrayMock)

	gui.SetSystemtray(st)

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
