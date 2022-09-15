package gui

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func createMenuItems() []MenuItem {
	items := []MenuItem{
		{Type: MenuItemButton, Title: "Test 1", Tooltip: "Test 1"},
		{Type: MenuItemSeperator},
		{Type: MenuItemCheckbox, Title: "Test 2", Tooltip: "Test 2", Checked: true},
		{Title: "Test 3", Tooltip: "Test 3"},
	}

	return items
}

func Test_SystemtraySet_should_set_the_title_tooltip_and_bytes(t *testing.T) {
	systemTray := SystemtrayMock{}

	systemTray.SetTitle("Test title")

	if systemTray.title != "Test title" {
		t.Error("Systemtray.SetTitle() did not set the title")
	}

	systemTray.SetTooltip("Test tooltip")

	if systemTray.tooltip != "Test tooltip" {
		t.Error("Systemtray.SetTooltip() did not set the tooltip")
	}

	var s string = "Test icon"
	sb := []byte(s)

	systemTray.SetIcon(sb)

	if !bytes.Equal(systemTray.iconBytes, sb) {
		t.Error("Systemtray.SetIcon() did not set the icon")
	}
}

func Test_SystemtraySetMenuItems_should_set_the_items(t *testing.T) {
	systemTray := SystemtrayMock{}

	items := createMenuItems()
	systemTray.SetMenuItems(items)

	if !reflect.DeepEqual(systemTray.menuItems, items) {
		t.Error("Systemtray.SetMenuItems() did not set the menu items")
	}
}

func Test_SystemtrayRun_should_create_the_running_systemtray(t *testing.T) {
	systemTray := SystemtrayMock{}

	systemTray.Run()

	if !systemTray.status.IsRunning {
		t.Error("Systemtray.Run() should create the systemtray but status.IsRunning is false")
	}
}

func Test_SystemtrayQuit_should_exit_the_running_systemtray(t *testing.T) {
	systemTray := SystemtrayMock{}
	systemTray.Run()

	systemTray.Quit()

	if systemTray.status.IsRunning {
		t.Error("Systemtray.Quit() should exit the systemtray but status.IsRunning is true")
	}
}

func Test_Systemtray_should_call_the_menuitem_callback_after_click(t *testing.T) {
	systemTray := SystemtrayMock{}
	callbackOneCalled, callbackTwoCalled := false, false

	callbackOne := func(checked bool) {
		callbackOneCalled = true
	}

	callbackTwo := func(checked bool) {
		callbackTwoCalled = true
	}

	menuItems := []MenuItem{
		{Type: MenuItemButton, Title: "Test1", OnClick: callbackOne},
		{Type: MenuItemButton, Title: "Test2", OnClick: callbackTwo},
	}

	systemTray.SetMenuItems(menuItems)
	systemTray.Run()

	systemTray.CallItemClickCallback()

	if !callbackOneCalled {
		t.Error("Systemtray should call the menuitem callback after item click but callbackOne was not called")
	}

	if !callbackTwoCalled {
		t.Error("Systemtray should call the menuitem callback after item click but callbackTwo was not called")
	}
}

func Test_IsMenuItemClickable_should_be_clickable_when_menuitem_is_type_MenuItemButton(t *testing.T) {
	item := MenuItem{
		Type: MenuItemButton,
	}

	if res := IsMenuItemClickable(item); !res {
		t.Errorf("IsMenuItemClickable() should return 'true' when menuitem is type MenuItemButton, but its '%s'", strconv.FormatBool(res))
	}
}

func Test_IsMenuItemClickable_should_be_clickable_when_menuitem_is_type_MenuItemCheckbox(t *testing.T) {
	item := MenuItem{
		Type: MenuItemCheckbox,
	}

	if res := IsMenuItemClickable(item); !res {
		t.Errorf("IsMenuItemClickable() should return 'true' when menuitem is type MenuItemButton, but its '%s'", strconv.FormatBool(res))
	}
}

func Test_IsMenuItemClickable_should_be_clickable_when_menuitem_is_type_MenuItemSeperator(t *testing.T) {
	item := MenuItem{
		Type: MenuItemSeperator,
	}

	if res := IsMenuItemClickable(item); res {
		t.Errorf("IsMenuItemClickable() should return 'false' when menuitem is type MenuItemSeperator, but its '%s'", strconv.FormatBool(res))
	}
}
