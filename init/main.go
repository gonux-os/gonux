package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"
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

func read_console(stdin chan<- string) {
	for {
		console, _ := openConsoleFile()
		in := bufio.NewReader(console)
		line, _ := in.ReadString('\n')
		stdin <- strings.Trim(line, "\n \t")
		console.Close()
	}
}

func write_console(stdout <-chan string) {
	for output := range stdout {
		console, _ := openConsoleFile()
		console.WriteString(output)
		console.Close()
	}
}

func read_tty(num int, stdin chan<- string) {
	for {
		tty, _ := openTTYFile(num)
		in := bufio.NewReader(tty)
		line, _ := in.ReadString('\n')
		stdin <- strings.Trim(line, "\n \t")
		tty.Close()
	}
}

func write_tty(num int, stdout <-chan string) {
	for output := range stdout {
		tty, _ := openTTYFile(num)
		tty.WriteString(output)
		tty.Close()
	}
}

func gosh(stdin <-chan string, stdout chan<- string) {
	prompt := "> "
	stdout <- prompt
	for input := range stdin {
		if strings.Split(input, " ")[0] == "cd" {
			args := strings.Split(input, " ")
			dir := "/"
			if len(args) > 1 {
				dir = args[1]
			}
			os.Chdir(dir)
			stdout <- prompt
			continue
		}

		// TODO: Write custom wrappers for these files
		childFiles := make([]*os.File, 0, 3)
		console, _ := openConsoleFile()
		childFiles = append(childFiles, console) // STDIN
		childFiles = append(childFiles, console) // STDOUT
		childFiles = append(childFiles, console) // STDERR

		proc, err := os.StartProcess("/bin/"+input, []string{""}, &os.ProcAttr{
			Files: childFiles,
		})
		if err != nil {
			stdout <- fmt.Sprintf("[Error] %v : %v\n", input, err.(*os.PathError).Err.Error())
		} else {
			// stdout <- fmt.Sprintf("Launched with PID: %v", proc.Pid)
			proc.Wait()
		}
		console.Close()
		stdout <- "\n"
		stdout <- prompt
	}
}

const VT_OPENQRY = 0x5600
const VT_ACTIVATE = 0x5606
const VT_WAITACTIVE = 0x5607

func nextAvailableVT() (int, error) {
	console, _ := openConsoleFile()
	defer console.Close()

	var vtNum int
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, console.Fd(), VT_OPENQRY, uintptr(unsafe.Pointer(&vtNum)))
	if errno != 0 {
		return -1, errors.New(errno.Error())
	}
	return vtNum, nil
}

func activateVT(vtNum int) error {
	tty, _ := openTTYFile(vtNum)
	defer tty.Close()

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, tty.Fd(), VT_ACTIVATE, uintptr(vtNum))
	if errno != 0 {
		return errors.New(errno.Error())
	}
	return nil
}

func switchTerminal() (int, error) {
	targetVtNum, err := nextAvailableVT()
	if err != nil {
		return -1, fmt.Errorf("could not query for available VT: %w", err)
	}

	err = activateVT(targetVtNum)
	if err != nil {
		return -1, fmt.Errorf("could not activate VT %d: %w", targetVtNum, err)
	}
	return targetVtNum, nil
}

func mountFilesystem() {
	err := syscall.Mount("none", "/proc", "proc", syscall.MS_BIND, "")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	err = syscall.Mount("none", "/sys", "sysfs", syscall.MS_BIND, "")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	stdin := make(chan string)
	stdout := make(chan string)

	_, _ = stdin, stdout

	stdin_tty := make(chan string)
	stdout_tty := make(chan string)

	mountFilesystem()

	tty_num, _ := switchTerminal()

	// go read_console(stdin)
	// go write_console(stdout)

	go read_tty(tty_num, stdin_tty)
	go write_tty(tty_num, stdout_tty)

	stdout_tty <- "Welcome to GoNUX!\n"

	go gosh(stdin_tty, stdout_tty)

	for {
		time.Sleep(1 * time.Microsecond)
	}
}
