package main

import (
	"fmt"
	"gonux/god/contracts"
	"gonux/god/local"
	"gonux/god/network"
	"os"
	"sync"
	"time"
)

//* Some provider process

func provider(god contracts.God) {
	god.Register("git.closePR", func() {
		fmt.Println("PR closed successfully")
	})
	god.Subscribe()
}

//* Some caller process

func caller(god contracts.God) {
	// On boot, wait until God contains all dependencies needed
	god.WaitFor("git.closePR")

	god.Call("git.closePR")
}

func main() {
	const mode = "network"
	const godnet = true

	if mode == "network" {
		//* Launch server
		wg := sync.WaitGroup{}
		if godnet {
			go network.StartGodServer("tcp", "127.0.0.1:8000", &wg)
		} else {
			os.Remove("/tmp/god.sock")
			go network.StartGodServer("unix", "/tmp/god.sock", &wg)
		}
		wg.Wait()
	}

	//* Load god instance
	var god contracts.God
	if mode == "local" {
		god = local.MakeGod()
	} else if mode == "network" {
		if godnet {
			god = network.MakeGodClient("127.0.0.1:8000")
		} else {
			god = network.MakeGodClient("unix:/tmp/god.sock")
		}
	}

	go provider(god)
	go caller(god)

	for {
		time.Sleep(1 * time.Nanosecond)
	}
}
