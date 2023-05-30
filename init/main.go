package main

import (
	_ "embed"
	"fmt"
	"gonux/init/terminal"
	"os"
	"strings"
	"sync"
	"syscall"
	"time"
)

func gosh(tty terminal.TTY) {
	prompt := "> "
	tty.Stdout <- prompt
	for input := range tty.Stdin {
		if strings.Split(input, " ")[0] == "cd" {
			args := strings.Split(input, " ")
			dir := "/"
			if len(args) > 1 {
				dir = args[1]
			}
			os.Chdir(dir)
			tty.Stdout <- prompt
			continue
		}

		// TODO: Write custom wrappers for these files
		childFiles := make([]*os.File, 0, 3)
		childFiles = append(childFiles, tty.File) // STDIN
		childFiles = append(childFiles, tty.File) // STDOUT
		childFiles = append(childFiles, tty.File) // STDERR

		proc, err := os.StartProcess("/bin/"+input, []string{""}, &os.ProcAttr{
			Files: childFiles,
		})
		if err != nil {
			tty.Stdout <- fmt.Sprintf("[Error] %v : %v\n", input, err.(*os.PathError).Err.Error())
		} else {
			// stdout <- fmt.Sprintf("Launched with PID: %v", proc.Pid)
			proc.Wait()
		}
		tty.Stdout <- "\n"
		tty.Stdout <- prompt
	}
}

func mountFilesystem() error {
	err := syscall.Mount("none", "/proc", "proc", syscall.MS_BIND, "")
	if err != nil {
		return fmt.Errorf("cannot mount /proc: %w", err)
	}
	err = syscall.Mount("none", "/sys", "sysfs", syscall.MS_BIND, "")
	if err != nil {
		return fmt.Errorf("cannot mount /sys: %w", err)
	}
	return nil
}

func hang() {
	for {
		time.Sleep(1 * time.Microsecond)
	}
}

func errorHandler(stdout chan<- string) {
	if r := recover(); r != nil {
		stdout <- fmt.Sprintf("[Error] %v\n", r)
		hang()
	}
}

func main() {

	wg := sync.WaitGroup{}
	go terminal.HandleChannelIO(&wg)

	defer errorHandler(terminal.ConsoleChannels.Stdout)

	tty, err := terminal.OpenTTY()
	if err != nil {
		panic(err)
	}

	wg.Done() // All TTYs are open

	defer errorHandler(tty.Stdout)

	// err = mountFilesystem()
	// if err != nil {
	// 	panic(err)
	// }

	tty.Stdout <- "Welcome to GoNUX!\n"

	go gosh(tty)

	hang()
}
