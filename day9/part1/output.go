package main

const outputOpCode = 4

type output struct {
	pointer int64
	value   int64
}

func (o *output) OpCode() int {
	return outputOpCode
}

func (o *output) Pointer() int64 {
	return o.pointer
}

func (o *output) Offset() int64 {
	return 2
}

func (o *output) Execute(memory []int64, relativeOffset int64) {
	outputInstruction := memory[o.Pointer()]

	switch instructionMode(outputInstruction).GetLeftOperandMode() {
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
