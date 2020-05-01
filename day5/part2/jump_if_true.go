package main

const jumpIfTrueOpCode = 5
const defaultJumpIfTrueOffset = 3

type jumpIfTrue struct {
	pointer int
	offset int
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

func (j *jumpIfTrue) Execute(memory []int) {
	jumpInstruction := memory[j.Pointer()]

	lhMode := instructionMode(jumpInstruction).GetLeftOperandMode()

	lh := 0

	if lhMode == immediateMode {
		lh = memory[j.Pointer() + 1]
	} else if lhMode == positionMode {
		lh = memory[memory[j.Pointer() + 1]]
	} else {
		panic(jumpInstruction)
	}

	if lh != 0 {

		rhMode := instructionMode(jumpInstruction).GetRightOperandMode()
		rh := 0

		if rhMode == immediateMode {
			rh = memory[j.Pointer() + 2]
		} else if rhMode == positionMode {
			rh = memory[memory[j.Pointer() + 2]]
		} else {
			panic(jumpInstruction)
		}

		j.setOffset(rh - j.Pointer())
	} else {
		j.setOffset(defaultJumpIfTrueOffset)
	}
}