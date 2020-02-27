package main

type multiply struct {
	pointer int
}

func (m *multiply) OpCode() int {
	return multiplyOpCode
}

func (m *multiply) Pointer() int {
	return m.pointer
}

func (m *multiply) Offset() int {
	return 4
}

func (m *multiply) Execute(memory []int) {

	param1 := memory[memory[m.pointer+1]]
	param2 := memory[memory[m.pointer+2]]

	memory[memory[m.pointer+3]] = param1 * param2
}
