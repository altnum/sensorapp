// +build windows

package win32

import (
	"golang.org/x/sys/windows"
)

var (
	moduser32 = windows.NewLazySystemDLL("user32.dll")

	procGetDesktopWindow = moduser32.NewProc("GetDesktopWindow")
)

// GetDesktopWindow returns the handle of the desktop window
func GetDesktopWindow() windows.HWND {
	result, _, _ := procGetDesktopWindow.Call()
	return windows.HWND(result)
}
