package main

const lessThanOpCode = 7

type lessThan struct {
	pointer int
	offset int
}

func (l *lessThan) OpCode() int {
	return lessThanOpCode
}

func (l *lessThan) Pointer() int {
	return l.pointer
}

func (l *lessThan) Offset() int {
	return 4
}

func (l *lessThan) Execute(memory []int) {
	lessThanInstruction := memory[l.Pointer()]

	lhMode := instructionMode(lessThanInstruction).GetLeftOperandMode()
	var lh int

	if lhMode == immediateMode {
		lh = memory[l.Pointer() + 1]
	} else if lhMode == positionMode {
		lh = memory[memory[l.Pointer() + 1]]
	} else {
		panic(lessThanInstruction)
	}

	rhMode := instructionMode(lessThanInstruction).GetRightOperandMode()
	var rh int

	if rhMode == immediateMode {
		rh = memory[l.Pointer() + 2]
	} else if rhMode == positionMode {
		rh = memory[memory[l.Pointer() + 2]]
	}

	var value int

	if lh < rh {
		value = 1
	} else {
		value = 0
	}

	assignmentMode := instructionMode(lessThanInstruction).GetAssignmentOperandMode()

	if assignmentMode == positionMode {
		memory[memory[l.Pointer() + 3]] = value
	} else {
		panic(lessThanInstruction)
	}
}
