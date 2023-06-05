package main

import (
	"fmt"
	"gonux/god/network"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	go network.StartGodServer("unix", "/tmp/god.sock", &wg)
	wg.Wait()
	fmt.Println("Running God server")

	for {
		time.Sleep(1 * time.Nanosecond)
	}
}
