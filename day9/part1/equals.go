package main

const equalsOpCode = 8

type equals struct {
	pointer int
	offset  int
}

func (e *equals) OpCode() int {
	return equalsOpCode
}

func (e *equals) Pointer() int {
	return e.pointer
}

func (e *equals) Offset() int {
	return 4
}

func (e *equals) Execute(memory []int, relativeOffset int) {
	equalsInstruction := memory[e.Pointer()]

	lhMode := instructionMode(equalsInstruction).GetLeftOperandMode()
	var lh int

	switch (lhMode); {
	case immediateMode:
		{
			lh = memory[e.Pointer()+1]
		}
	case positionMode:
		{
			lh = memory[memory[e.Pointer()+1]]
		}
	case relativeMode:
		{
			lh = memory[relativeOffset+memory[e.Pointer()+1]]
		}
	default:
		panic(equalsInstruction)
	}

	rhMode := instructionMode(equalsInstruction).GetRightOperandMode()
	var rh int

	switch (rhMode); {
	case immediateMode:
		{
			rh = memory[e.Pointer()+2]
		}
	case positionMode:
		{
			rh = memory[memory[e.Pointer()+2]]
		}
	case relativeMode:
		{
			rh = memory[relativeOffset+memory[e.Pointer()+2]]
		}
	default:
		panic(equalsInstruction)
	}

	var value int

	if lh == rh {
		value = 1
	} else {
		value = 0
	}

	assignmentMode := instructionMode(equalsInstruction).GetAssignmentOperandMode()

	if assignmentMode == positionMode {
		memory[memory[e.Pointer()+3]] = value
	} else if assignmentMode == relativeMode {
		memory[relativeOffset+memory[e.Pointer()+3]] = value
	} else {
		panic(equalsInstruction)
	}
}
