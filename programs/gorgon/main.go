package main

import (
	"os"
	"syscall"
	"time"
)

// "github.com/NeowayLabs/drm"

const (
	KDSETMODE = 0x4B3A

	KD_TEXT     = 0x00
	KD_GRAPHICS = 0x01
)

func main() {
	console, _ := os.OpenFile("/dev/console", os.O_RDWR, 0)
	syscall.Syscall(syscall.SYS_IOCTL, console.Fd(), uintptr(KDSETMODE), uintptr(KD_GRAPHICS))
	time.Sleep(3 * time.Second)
	syscall.Syscall(syscall.SYS_IOCTL, console.Fd(), uintptr(KDSETMODE), uintptr(KD_TEXT))
}
