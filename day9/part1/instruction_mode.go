package main

const (
	positionMode  instructionMode = iota
	immediateMode instructionMode = iota
	relativeMode  instructionMode = iota
)

type instructionMode int

func (i instructionMode) GetLeftOperandMode() instructionMode {
	temp := i

	// remove instruction
	temp = temp / 100

	return instructionMode(temp % 10)
}

func (i instructionMode) GetRightOperandMode() instructionMode {
	temp := int64(i)

	// remove instruction + lh operand
	temp = temp / 1000

	return instructionMode(temp % 10)
}

func (i instructionMode) GetAssignmentOperandMode() instructionMode {
	temp := i

	// remove instruction + lh operand + rh operand
	temp = temp / 10_000

	return instructionMode(temp % 10)
}
