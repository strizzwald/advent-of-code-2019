package main

import "math"

const jumpIfFalseOpCode = 6
const defaultJumpIfFalseOffset = 3

type jumpIfFalse struct {
	pointer int
	offset  int
}

func (j *jumpIfFalse) OpCode() int {
	return jumpIfFalseOpCode
}

func (j *jumpIfFalse) Pointer() int {
	return j.pointer
}

func (j *jumpIfFalse) Offset() int {
	return j.offset
}

func (j *jumpIfFalse) setOffset(offset int) {
	j.offset = offset
}

func (j *jumpIfFalse) Execute(memory []int) {
	jumpInstruction := memory[j.Pointer()]

	lh := math.MaxInt64
	lhMode := instructionMode(jumpInstruction).GetLeftOperandMode()

	if lhMode == immediateMode {
		lh = memory[j.Pointer()+1]
	} else if lhMode == positionMode {
		lh = memory[memory[j.Pointer()+1]]
	} else {
		panic(jumpInstruction)
	}

	if lh == 0 {
		var rh int

		rhMode := instructionMode(jumpInstruction).GetRightOperandMode()

		if rhMode == immediateMode {
			rh = memory[j.Pointer()+2]
		} else if rhMode == positionMode {
			rh = memory[memory[j.Pointer()+2]]
		} else {
			panic(jumpInstruction)
		}

		j.setOffset(rh - j.Pointer())
	} else {
		j.setOffset(defaultJumpIfFalseOffset)
	}
}
