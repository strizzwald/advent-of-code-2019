package main

const storeOpCode = 3

type store struct {
	pointer int
	input   int
}

func (s *store) OpCode() int {
	return storeOpCode
}

func (s *store) Pointer() int {
	return s.pointer
}

func (s *store) Offset() int {
	return 2
}

func (s *store) SetInput(input int) {
	s.input = input
}

func (s *store) Input() int {
	return s.input
}

func (s *store) Execute(memory []int, relativeOffset int) {
	storeInstruction := memory[s.pointer]

	if instructionMode(storeInstruction).GetAssignmentOperandMode() == positionMode {
		memory[memory[s.Pointer() + 1]] = s.Input()
	} else if (instructionMode(storeInstruction).GetAssignmentOperandMode() == relativeMode) {
		memory[relativeOffset + memory[s.Pointer() + 1]] = s.Input()
		
	} else {
		panic(storeInstruction)
	}

}
