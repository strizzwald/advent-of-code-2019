package main

const outputOpCode = 4

type output struct {
	pointer int
	value int
}

func (o *output) OpCode() int {
	return outputOpCode
}

func (o *output) Pointer() int {
	return o.pointer
}

func (o *output) Offset() int {
	return 2
}

func (o *output) Execute(memory[] int) {
	outputInstruction := memory[o.pointer]

	if instructionMode(outputInstruction).GetAssignmentOperandMode() == immediateMode {
		o.value = memory[o.pointer + 1]
	} else if instructionMode(outputInstruction).GetAssignmentOperandMode() == positionMode {
		o.value = memory[memory[o.pointer + 1]]
	} else {
		panic(outputInstruction)
	}
}