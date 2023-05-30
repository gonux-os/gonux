package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

func openConsoleFile() (*os.File, error) {
	return os.OpenFile("/dev/console", os.O_RDWR, 0)
}

func openTTYFile(num int) (*os.File, error) {
	return os.OpenFile(fmt.Sprintf("/dev/tty%d", num), os.O_RDWR, 0)
}

type StdioChannels struct {
	Stdin  chan string
	Stdout chan string
}

type TTY struct {
	StdioChannels
	Num  int
	File *os.File
}

var ConsoleChannels = StdioChannels{
	Stdin:  make(chan string),
	Stdout: make(chan string),
}
var ttyChannels = []TTY{}

func (s StdioChannels) readConsole() {
	for {
		console, _ := openConsoleFile()
		in := bufio.NewReader(console)
		line, _ := in.ReadString('\n')
		s.Stdin <- strings.Trim(line, "\n \t")
		console.Close()
	}
}

func (s StdioChannels) writeConsole() {
	for output := range s.Stdout {
		console, _ := openConsoleFile()
		console.WriteString(output)
		console.Close()
	}
}

func (s TTY) readTTY() {
	for {
		tty, _ := openTTYFile(s.Num)
		in := bufio.NewReader(tty)
		line, _ := in.ReadString('\n')
		s.Stdin <- strings.Trim(line, "\n \t")
		tty.Close()
	}
}

func (s TTY) writeTTY() {
	for output := range s.Stdout {
		tty, _ := openTTYFile(s.Num)
		tty.WriteString(output)
		tty.Close()
	}
}

func HandleChannelIO(wg *sync.WaitGroup) {
	wg.Add(1)
	go ConsoleChannels.readConsole()
	go ConsoleChannels.writeConsole()
	wg.Wait()
	for _, ttyChannel := range ttyChannels {
		go ttyChannel.readTTY()
		go ttyChannel.writeTTY()
	}
	for {
		time.Sleep(1 * time.Microsecond)
	}
}

const (
	_VT_OPENQRY  = 0x5600
	_VT_ACTIVATE = 0x5606
)

func nextAvailableVT() (int, error) {
	console, _ := openConsoleFile()
	defer console.Close()

	var vtNum int
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, console.Fd(), _VT_OPENQRY, uintptr(unsafe.Pointer(&vtNum)))
	if errno != 0 {
		return -1, fmt.Errorf("ioctl VT_OPENQRY error: %v", errno)
	}
	return vtNum, nil
}

func activateTTY(num int) error {
	tty, _ := openTTYFile(num)
	defer tty.Close()

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, tty.Fd(), _VT_ACTIVATE, uintptr(num))
	if errno != 0 {
		return fmt.Errorf("ioctl VT_ACTIVATE error: %v", errno)
	}

	return nil
}

func openTTY(redirect bool) (TTY, error) {
	targetVtNum, err := nextAvailableVT()
	if err != nil {
		return TTY{}, fmt.Errorf("could not query for available VT: %w", err)
	}

	err = activateTTY(targetVtNum)
	if err != nil {
		return TTY{}, fmt.Errorf("could not activate VT %d: %w", targetVtNum, err)
	}

	f, err := openTTYFile(targetVtNum)
	if err != nil {
		return TTY{}, fmt.Errorf("could not open TTY file for VT %d: %w", targetVtNum, err)
	}

	ttyChan := TTY{
		Num:  targetVtNum,
		File: f,
		StdioChannels: StdioChannels{
			Stdin:  make(chan string),
			Stdout: make(chan string),
		},
	}
	if redirect {
		ttyChannels = append(ttyChannels, ttyChan)
	}
	return ttyChan, nil
}

func OpenTTY(redirect ...bool) (TTY, error) {
	if len(redirect) > 0 {
		return openTTY(redirect[0])
	}
	return openTTY(true)
}
