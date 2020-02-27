package main

type add struct {
	pointer int
}

func (a *add) OpCode() int {
	return addOpCode
}

func (a *add) Pointer() int {
	return a.pointer
}

func (a *add) Offset() int {
	return 4
}

func (a *add) Execute(memory []int) {
	param1 := memory[memory[a.pointer+1]]
	param2 := memory[memory[a.pointer+2]]

	memory[memory[a.pointer+3]] = param1 + param2
}
