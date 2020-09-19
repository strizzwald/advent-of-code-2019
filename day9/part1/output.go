package main

const outputOpCode = 4

type output struct {
	pointer int
	value   int
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

func (o *output) Execute(memory []int, relativeOffset int) {
	outputInstruction := memory[o.Pointer()]

	switch instructionMode(outputInstruction).GetAssignmentOperandMode {
	case immediateMode:
		{
			o.value = memory[o.Pointer()+1]
		}
	case positionMode:
		{
			o.value = memory[memory[o.Pointer()+1]]
		}
	case relativeMode:
		{
			o.value = memory[relativeOffset+memory[o.Pointer()+1]]
		}
	default:
		panic(outputInstruction)
	}
}
