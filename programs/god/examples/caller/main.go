package main

import (
	"fmt"
	"gonux/god/network"
)

func main() {
	god := network.MakeGodClient("unix:/tmp/god.sock")

	god.WaitFor("example.foo")

	fmt.Println("Calling example.foo")
	god.Call("example.foo")
}
