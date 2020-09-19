package main

const relativeBaseOffsetOpCode = 9

type relativeBaseOffset struct {
	pointer int
	value   int
}

func (r *relativeBaseOffset) OpCode() int {
	return relativeBaseOffsetOpCode
}

func (r *relativeBaseOffset) Pointer() int {
	r.pointer
}

func (r *relativeBaseOffset) Offset() int {
	return 2
}

func (r *relativeBaseOffset) Value() int {
	r.value
}

func (r *relativeBaseOffset) Execute(memory []int, relativeOffset int) {
	offsetInstruction := memory[r.Pointer()]

	switch (instructionMode(offsetInstruction).GetLeftOperandMode()); {
	case immediateMode:
		{
			r.output = relativeOffset + memory[r.Pointer()+1]
		}
	case positionMode:
		{
			r.output = relativeOffset + memory[memory[r.Pointer()+1]]
		}
	case relativeMode:
		{
			r.output = relativeOffset + memory[relativeOffset+memory[r.Pointer()+1]]
		}
	default:
		panic(offsetInstruction)
	}
}
