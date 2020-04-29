package main

const immediateMode int = 1
const positionMode int = 0

type instructionMode int

func (i instructionMode) GetLeftOperandMode() int {
	temp := i

	// remove instruction
	temp = temp / 100

	return int(temp  % 10)
}


func (i instructionMode) GetRightOperandMode() int {
	temp := i

	// remove instruction + lh operand
	temp = temp / 1000

	return int(temp % 10)
}

func (i instructionMode) GetAssignmentOperand() int {
	temp := i

	// remove instruction + lh operand + rh operand
	temp = temp / 10_000

	return int(temp % 10)
}
