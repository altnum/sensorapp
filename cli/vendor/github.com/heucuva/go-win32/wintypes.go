// +build windows

package win32

const (
	TRUE  = 1
	FALSE = 0
)

func ToBOOL(b bool) uint32 {
	if b {
		return TRUE
	}
	return FALSE
}
