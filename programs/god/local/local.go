package local

import (
	"gonux/god/contracts"
	"time"
)

type God struct {
	actions map[string]func()
}

func MakeGod() contracts.God {
	return God{
		actions: make(map[string]func()),
	}
}

func (g God) Register(name string, impl func()) {
	g.actions[name] = impl
}

func (g God) Call(name string) {
	g.actions[name]()
}

func (g God) WaitFor(name string) {
	for {
		if _, ok := g.actions[name]; ok {
			break
		}
		time.Sleep(1 * time.Nanosecond)
	}
}

func (g God) Subscribe() {
	//
}
