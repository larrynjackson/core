package main

func (core *Config) noop() int {

	core.OperationClass = "fetch"
	core.clockTick(7)
	return 1
}
