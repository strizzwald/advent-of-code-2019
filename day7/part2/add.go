package main

const addOpCode = 1

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
	addInstruction := memory[a.pointer]

	lhMode := instructionMode(addInstruction).GetLeftOperandMode()
	rhMode := instructionMode(addInstruction).GetRightOperandMode()
	assMode := instructionMode(addInstruction).GetAssignmentOperandMode()

	lh := 0

	if lhMode == immediateMode {
		lh = memory[a.pointer+1]
	} else if lhMode == positionMode {
		lh = memory[memory[a.pointer+1]]
	} else {
		panic(lh)
	}

	rh := 0

	if rhMode == immediateMode {
		rh = memory[a.pointer+2]
	} else if rhMode == positionMode {
		rh = memory[memory[a.pointer+2]]
	} else {
		panic(rh)
	}

	if assMode == positionMode {
		memory[memory[a.pointer+3]] = lh + rh
	} else {
		panic(addInstruction)
	}
}
