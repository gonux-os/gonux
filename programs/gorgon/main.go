package main

import (
	"os"
	"syscall"
	"time"
)

const (
	KDSETMODE = 0x4B3A

	KD_TEXT     = 0x00
	KD_GRAPHICS = 0x01
)

func main() {
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(KDSETMODE), uintptr(KD_GRAPHICS))
	time.Sleep(3 * time.Second)
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdout.Fd(), uintptr(KDSETMODE), uintptr(KD_TEXT))
}
