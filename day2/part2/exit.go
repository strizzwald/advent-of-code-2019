package main

type exit struct {
	pointer int
}

func (e *exit) OpCode() int {
	return exitOpCode
}

func (e *exit) Pointer() int {
	return e.pointer
}

func (e *exit) Offset() int {
	return 4
}

func (e *exit) Execute(memory []int) {
	return
}
