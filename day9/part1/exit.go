package main

const exitOpCode = 99

type exit struct {
	pointer int64
}

func (e *exit) OpCode() int {
	return exitOpCode
}

func (e *exit) Pointer() int64 {
	return e.pointer
}

func (e *exit) Offset() int64 {
	return 1
}

func (e *exit) Execute(memory []int64, relativeOffset int64) {
	return
}
