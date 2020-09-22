package main

const multiplyOpCode = 2

type multiply struct {
	pointer int64
}

func (m *multiply) OpCode() int {
	return multiplyOpCode
}

func (m *multiply) Pointer() int64 {
	return m.pointer
}

func (m *multiply) Offset() int64 {
	return 4
}

func (m *multiply) Execute(memory []int64, relativeOffset int64) {
	multiplyInstruction := memory[m.Pointer()]

	lhMode := instructionMode(multiplyInstruction).GetLeftOperandMode()
	rhMode := instructionMode(multiplyInstruction).GetRightOperandMode()
	assMode := instructionMode(multiplyInstruction).GetAssignmentOperandMode()

	var lh int64

	switch lhMode {
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
			lh = memory[relativeOffset+memory[m.Pointer()+1]]
		}
	default:
		panic(multiplyInstruction)
	}

	var rh int64

	switch rhMode {
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
