package helper

import (
	"os"
	"syscall"
	"unsafe"
)

func IsTerminalContext() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), syscall.TCGETS, uintptr(unsafe.Pointer(&termios)))
	return err == 0
}
