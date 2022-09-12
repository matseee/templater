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

func Test_GuiSetApi_should_set_the_given_api_object(t *testing.T) {
	gui, _ := createGui()
	api := new(ApiMock)
	var inter interface{} = api

	gui.SetApi(api)

	if gui.api != inter {
		t.Errorf("Gui.SetApi() should set the given api object but got nil or another object")
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

// func Test_getRandomPort(t *testing.T) {
// 	randomPort := getRandomPort()

// 	for i := 0; i < 1000; i++ {
// 		if randomPort < PORTMIN || randomPort > PORTMAX {
// 			t.Errorf("Random port should be between %d and %d, but was %d - FAILED", PORTMIN, PORTMAX, randomPort)
// 		}
// 	}

// 	t.Logf("Random ports were all between %d and %d - PASSED", PORTMIN, PORTMAX)
// }

// func Test_getURL(t *testing.T) {
// 	port = 22334
// 	url := getURL()

// 	if url == "http://127.0.0.1:22334/" {
// 		t.Log("URL was as expected \"http://127.0.0.1:22334/\" - PASSED")
// 	} else {
// 		t.Errorf("URL was not as expected. Got \"%s\" - FAILED", url)
// 	}
// }

// func Test_isPortAvailable(t *testing.T) {
// 	port := "22334"
// 	listener, err := net.Listen("tcp", ":"+port)

// 	if err != nil {
// 		t.Log("Can't open port 22334 for test, skipping test")
// 		t.SkipNow()
// 		return
// 	}

// 	isAvailable := isPortAvailable(22334)
// 	if isAvailable {
// 		t.Error("Port 22334 should not be available, but is available - FAILED")
// 		return
// 	} else {
// 		t.Log("Port 22334 should not be available and isn't available - SUCCESS")
// 	}

// 	listener.Close()

// 	isAvailable = isPortAvailable(22334)
// 	if isAvailable {
// 		t.Log("Port 22334 should be available and is available - SUCCESS")
// 	} else {
// 		t.Error("Port 22334 should be available, but isn't available - FAILED")
// 	}
// }
