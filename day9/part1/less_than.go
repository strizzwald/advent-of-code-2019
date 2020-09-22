package main

const lessThanOpCode = 7

type lessThan struct {
	pointer int64
	offset  int64
}

func (l *lessThan) OpCode() int {
	return lessThanOpCode
}

func (l *lessThan) Pointer() int64 {
	return l.pointer
}

func (l *lessThan) Offset() int64 {
	return 4
}

func (l *lessThan) Execute(memory []int64, relativeOffset int64) {
	lessThanInstruction := memory[l.Pointer()]

	lhMode := instructionMode(lessThanInstruction).GetLeftOperandMode()
	var lh int64

	switch lhMode {
	case immediateMode:
		{
			lh = memory[l.Pointer()+1]
		}
	case positionMode:
		{
			lh = memory[memory[l.Pointer()+1]]
		}
	case relativeMode:
		{
			lh = memory[relativeOffset+memory[l.Pointer()+1]]
		}
	default:
		panic(lessThanInstruction)
	}

	rhMode := instructionMode(lessThanInstruction).GetRightOperandMode()
	var rh int64

	switch rhMode {
	case immediateMode:
		{
			rh = memory[l.Pointer()+2]
		}
	case positionMode:
		{
			rh = memory[memory[l.Pointer()+2]]
		}
	case relativeMode:
		{
			rh = memory[relativeOffset+memory[l.Pointer()+2]]
		}
	default:
		panic(lessThanInstruction)
	}

	var value int64

	if lh < rh {
		value = 1
	} else {
		value = 0
	}

	assignmentMode := instructionMode(lessThanInstruction).GetAssignmentOperandMode()

	if assignmentMode == positionMode {
		memory[memory[l.Pointer()+3]] = value
	} else if assignmentMode == relativeMode {
		memory[relativeOffset+memory[l.Pointer()+3]] = value
	} else {
		panic(lessThanInstruction)
	}
}
