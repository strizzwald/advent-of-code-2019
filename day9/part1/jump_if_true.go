package main

const jumpIfTrueOpCode = 5
const defaultJumpIfTrueOffset = 3

type jumpIfTrue struct {
	pointer int
	offset  int
}

func (j *jumpIfTrue) OpCode() int {
	return jumpIfTrueOpCode
}

func (j *jumpIfTrue) Pointer() int {
	return j.pointer
}

func (j *jumpIfTrue) Offset() int {
	return j.offset
}

func (j *jumpIfTrue) setOffset(offset int) {
	j.offset = offset
}

func (j *jumpIfTrue) Execute(memory []int, relativeOffset int) {
	jumpInstruction := memory[j.Pointer()]

	lhMode := instructionMode(jumpInstruction).GetLeftOperandMode()

	lh := 0

	switch (lhMode); {
	case immediateMode:
		{
			lh = memory[j.Pointer()+1]
		}
	case positionMode:
		{
			lh = memory[memory[j.Pointer()+1]]
		}
	case relativeMode:
		{
			lh = memory[relativeOffset+memory[j.Pointer()+1]]
		}
	default:
		panic(jumpInstruction)
	}

	if lh != 0 {

		rhMode := instructionMode(jumpInstruction).GetRightOperandMode()
		rh := 0

		switch (rhMode); {
		case immediateMode:
			{
				rh = memory[j.Pointer()+2]
			}
		case positionMode:
			{
				rh = memory[memory[j.Pointer()+2]]
			}
		case relativeMode:
			{
				rh = memory[relativeOffset+memory[j.Pointer()+2]]
			}
		default:
			panic(jumpInstruction)
		}

		j.setOffset(rh - j.Pointer())
	} else {
		j.setOffset(defaultJumpIfTrueOffset)
	}
}
