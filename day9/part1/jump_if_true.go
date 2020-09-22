package main

const jumpIfTrueOpCode = 5
const defaultJumpIfTrueOffset = 3

type jumpIfTrue struct {
	pointer int64
	offset  int64
}

func (j *jumpIfTrue) OpCode() int {
	return jumpIfTrueOpCode
}

func (j *jumpIfTrue) Pointer() int64 {
	return j.pointer
}

func (j *jumpIfTrue) Offset() int64 {
	return j.offset
}

func (j *jumpIfTrue) setOffset(offset int64) {
	j.offset = offset
}

func (j *jumpIfTrue) Execute(memory []int64, relativeOffset int64) {
	jumpInstruction := memory[j.Pointer()]

	lhMode := instructionMode(jumpInstruction).GetLeftOperandMode()

	var lh int64

	switch lhMode {
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
		var rh int64

		switch rhMode {
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

		j.setOffset(rh - int64(j.Pointer()))
	} else {
		j.setOffset(defaultJumpIfTrueOffset)
	}
}
