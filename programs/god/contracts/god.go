package contracts

type God interface {
	Register(name string, impl func())
	Call(name string)
	WaitFor(name string)
	Subscribe()
}
