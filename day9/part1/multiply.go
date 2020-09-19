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

func (m *multiply) Execute(memory []int, relativeOffset int) {
	multiplyInstruction := memory[m.Pointer()+1]

	lhMode := instructionMode(multiplyInstruction).GetLeftOperandMode()
	rhMode := instructionMode(multiplyInstruction).GetRightOperandMode()
	assMode := instructionMode(multiplyInstruction).GetAssignmentOperandMode()

	lh := 0

	switch (lhMode); {
	case immediateMode:
		{
			lh = memory[m.Pointer()+1]
		}
	case positionMode:
		{
			lh = memory[memory[m.Pointer()+1]]
		}
	case relativeMode:
		{
			lh = memory[relativeOffset+m.Pointer()+1]
		}
	default:
		panic(multiplyInstruction)
	}

	rh := 0

	switch (rhMode); {
	case immediateMode:
		{
			rh = memory[m.Pointer()+2]
		}
	case positionMode:
		{
			rh = memory[memory[m.Pointer()+2]]
		}
	case relativeMode:
		{
			rh = memory[relativeOffset+memory[m.Pointer()+2]]
		}
	default:
		panic(multiplyInstruction)
	}

	if assMode == positionMode {
		memory[memory[m.Pointer()+3]] = lh * rh
	} else if assMode == relativeMode {
		memory[relativeOffset+memory[m.Pointer()+3]] = lh * rh
	}
}
