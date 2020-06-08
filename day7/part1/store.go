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

func (s *store) Execute(memory []int) {
	storeInstruction := memory[s.pointer]

	if instructionMode(storeInstruction).GetAssignmentOperandMode() == positionMode {
		memory[memory[s.pointer+1]] = s.input
	} else {
		panic(storeInstruction)
	}

}
