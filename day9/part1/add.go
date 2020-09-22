package main

const addOpCode = 1

type add struct {
	pointer int64
}

func (a *add) OpCode() int {
	return addOpCode
}

func (a *add) Pointer() int64 {
	return a.pointer
}

func (a *add) Offset() int64 {
	return 4
}

func (a *add) Execute(memory []int64, relativeOffset int64) {
	addInstruction := memory[a.Pointer()]

	lhMode := instructionMode(addInstruction).GetLeftOperandMode()
	rhMode := instructionMode(addInstruction).GetRightOperandMode()
	assMode := instructionMode(addInstruction).GetAssignmentOperandMode()

	var lh int64

	switch lhMode {

	case immediateMode:
		{
			lh = memory[a.Pointer()+1]
		}
	case positionMode:
		{
			lh = memory[memory[a.Pointer()+1]]
		}
	case relativeMode:
		{
			lh = memory[relativeOffset+memory[a.Pointer()+1]]
		}
	default:
		panic(lh)
	}

	var rh int64

	switch rhMode {
	case immediateMode:
		{
			rh = memory[a.Pointer()+2]
		}
	case positionMode:
		{
			rh = memory[memory[a.Pointer()+2]]
		}
	case relativeMode:
		{
			rh = memory[relativeOffset+memory[a.Pointer()+2]]
		}
	default:
		panic(rh)
	}

	if assMode == positionMode {
		memory[memory[a.Pointer()+3]] = lh + rh
	} else if assMode == relativeMode {
		memory[relativeOffset+memory[a.Pointer()+3]] = lh + rh
	} else {
		panic(addInstruction)
	}
}
