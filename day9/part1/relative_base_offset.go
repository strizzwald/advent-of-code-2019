package main

const relativeBaseOffsetOpCode = 9

type relativeBaseOffset struct {
	pointer int64
	value   int64
}

func (r *relativeBaseOffset) OpCode() int {
	return relativeBaseOffsetOpCode
}

func (r *relativeBaseOffset) Pointer() int64 {
	return r.pointer
}

func (r *relativeBaseOffset) Offset() int64 {
	return 2
}

func (r *relativeBaseOffset) Value() int64 {
	return r.value
}

func (r *relativeBaseOffset) Execute(memory []int64, relativeOffset int64) {
	offsetInstruction := memory[r.Pointer()]

	switch instructionMode(offsetInstruction).GetLeftOperandMode() {
	case immediateMode:
		{
			r.value = relativeOffset + memory[r.Pointer()+1]
		}
	case positionMode:
		{
			r.value = relativeOffset + memory[memory[r.Pointer()+1]]
		}
	case relativeMode:
		{
			r.value = relativeOffset + memory[relativeOffset+memory[r.Pointer()+1]]
		}
	default:
		panic(offsetInstruction)
	}
}
