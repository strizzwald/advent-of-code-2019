package main

const equalsOpCode = 8

type equals struct {
	pointer int
	offset int
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

func (e *equals) Execute(memory []int) {
	equalsInstruction := memory[e.Pointer()]

	lhMode := instructionMode(equalsInstruction).GetLeftOperandMode()
	var lh int

	if lhMode == immediateMode {
		lh = memory[e.Pointer() + 1]
	} else if lhMode == positionMode {
		lh = memory[memory[e.Pointer() + 1]]
	} else {
		panic(equalsInstruction)
	}

	rhMode := instructionMode(equalsInstruction).GetRightOperandMode()
	var rh int

	if rhMode == immediateMode {
		rh = memory[e.Pointer() + 2]
	} else if rhMode == positionMode {
		rh = memory[memory[e.Pointer() + 2]]
	} else {
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
		memory[memory[e.Pointer() + 3]] = value
	} else {
		panic(equalsInstruction)
	}
}