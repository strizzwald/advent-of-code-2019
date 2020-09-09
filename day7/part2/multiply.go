package main

const multiplyOpCode = 2

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
	multiplyInstruction := memory[m.pointer]

	lhMode := instructionMode(multiplyInstruction).GetLeftOperandMode()
	rhMode := instructionMode(multiplyInstruction).GetRightOperandMode()
	assMode := instructionMode(multiplyInstruction).GetAssignmentOperandMode()

	lh := 0

	if lhMode == immediateMode {
		lh = memory[m.pointer+1]
	} else if lhMode == positionMode {
		lh = memory[memory[m.pointer+1]]
	}

	rh := 0

	if rhMode == immediateMode {
		rh = memory[m.pointer+2]
	} else if rhMode == positionMode {
		rh = memory[memory[m.pointer+2]]
	}

	if assMode == positionMode {
		memory[memory[m.pointer+3]] = lh * rh
	} else {
		panic(multiplyInstruction)
	}
}
