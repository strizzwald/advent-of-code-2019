package main

const exitOpCode = 99

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
	return 1
}

func (e *exit) Execute(memory []int, relativeOffset int) {
	return
}
