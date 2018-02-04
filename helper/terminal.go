package helper

import (
	"os"
	"syscall"
	"unsafe"
)

// IsTerminalContext checks if the current process is in a terminal context
func IsTerminalContext() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), syscall.TCGETS, uintptr(unsafe.Pointer(&termios)))
	return err == 0
}
