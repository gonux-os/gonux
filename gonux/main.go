package main

import (
	"fmt"
	"time"
)

func read_console(stdin chan<- string) {
	for {
		var line string
		fmt.Scanln(&line)
		stdin <- line
	}
}

func write_console(stdout <-chan string) {
	for output := range stdout {
		fmt.Print(output)
	}
}

func gosh(stdin <-chan string, stdout chan<- string) {
	prompt := "> "
	stdout <- prompt
	for input := range stdin {
		stdout <- fmt.Sprintf("Launching %v...", input)
		// TODO: Spawn process with name "input"
		stdout <- "\n"
		stdout <- prompt
	}
}

func main() {
	stdin := make(chan string)
	stdout := make(chan string)

	fmt.Println("Welcome to GoNUX!")

	go read_console(stdin)
	go write_console(stdout)

	go gosh(stdin, stdout)

	for {
		time.Sleep(1 * time.Microsecond)
	}
}
