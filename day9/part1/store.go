package main

const storeOpCode = 3

type store struct {
	pointer int64
	input   int
}

func (s *store) OpCode() int {
	return storeOpCode
}

func (s *store) Pointer() int64 {
	return s.pointer
}

func (s *store) Offset() int64 {
	return 2
}

func (s *store) SetInput(input int) {
	s.input = input
}

func (s *store) Execute(memory []int64, relativeOffset int64) {
	storeInstruction := memory[s.pointer]

	if instructionMode(storeInstruction).GetLeftOperandMode() == positionMode {
		memory[memory[s.Pointer()+1]] = int64(s.input)
	} else if instructionMode(storeInstruction).GetLeftOperandMode() == relativeMode {
		memory[relativeOffset+memory[s.Pointer()+1]] = int64(s.input)

	} else {
		panic(storeInstruction)
	}

}
