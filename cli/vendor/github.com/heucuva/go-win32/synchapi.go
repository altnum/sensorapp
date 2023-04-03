// +build windows

package win32

import (
	"os"
	"sync/atomic"
	"time"

	"golang.org/x/sys/windows"
)

// CreateEvent creates a handle for event operations
func CreateEvent(eventAttributes *windows.SecurityAttributes, manualReset bool, initialState bool, name string) (windows.Handle, error) {
	var (
		param1 = ToBOOL(manualReset)
		param2 = ToBOOL(initialState)
		param3 *uint16
	)
	if name != "" {
		if v, err := windows.UTF16PtrFromString(name); err == nil {
			param3 = v
		}
	}
	return windows.CreateEvent(eventAttributes, param1, param2, param3)
}

var CloseHandle = windows.CloseHandle

// WaitForSingleObjectInfinite will wait infinitely for a single handle value to become available
func WaitForSingleObjectInfinite(handle windows.Handle) error {
	return internalWaitForSingleObject(handle, windows.INFINITE)
}

// WaitForSingleObject will wait for a single handle value to become available up to a total of `duration` milliseconds
func WaitForSingleObject(handle windows.Handle, duration time.Duration) error {
	return internalWaitForSingleObject(handle, uint32(duration.Milliseconds()))
}

func internalWaitForSingleObject(handle windows.Handle, duration uint32) error {
	h := atomic.LoadUintptr((*uintptr)(&handle))
	s, e := windows.WaitForSingleObject(windows.Handle(h), duration)
	switch s {
	case windows.WAIT_OBJECT_0:
		// we're good
	//case windows.WAIT_FAILED:
	default:
		return os.NewSyscallError("WaitForSingleObject", e)
	}
	return nil
}
