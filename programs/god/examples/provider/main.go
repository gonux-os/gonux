package main

import (
	"fmt"
	"gonux/god/network"
	"time"
)

func main() {
	god := network.MakeGodClient("unix:/tmp/god.sock")

	fmt.Println("Registering example.foo")

	god.Register("example.foo", func() {
		fmt.Println("example.foo called")
	})

	go god.Subscribe()
	fmt.Println("Provider running")

	for {
		time.Sleep(1 * time.Nanosecond)
	}
}
