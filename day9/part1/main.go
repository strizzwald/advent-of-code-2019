package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	program := loadProgram()
	executeProgram(program)
}

func loadProgram() []int64 {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	inputTxt := strings.Split(strings.Trim(string(input), "\n"), ",")

	memory := make([]int64, len(inputTxt))

	for i := range inputTxt {
		memory[i], err = strconv.ParseInt(inputTxt[i], 10, 64)

		if err != nil {
			panic(err)
		}
	}

	return memory
}

func newInstruction(pointer int, memory []int64) instruction {
	opCode := memory[pointer] % 100

	switch opCode {
	case storeOpCode:
		return &store{pointer: pointer}
	case outputOpCode:
		return &output{pointer: pointer}
	case addOpCode:
		return &add{pointer: pointer}
	case multiplyOpCode:
		return &multiply{pointer: pointer}
	case jumpIfTrueOpCode:
		return &jumpIfTrue{pointer: pointer}
	case jumpIfFalseOpCode:
		return &jumpIfFalse{pointer: pointer}
	case lessThanOpCode:
		return &lessThan{pointer: pointer}
	case equalsOpCode:
		return &equals{pointer: pointer}
	case exitOpCode:
		return &exit{pointer: pointer}
	default:
		panic(fmt.Sprintf("Unknown instruction: %d", memory[pointer]))
	}
}

func executeProgram(program []int64) {
	program := append(program, make([]int64, length(program)*2))

	var offsetValue int
	var pointer int
	var exited boolean

	for !exited {
		ins := newInstruction(pointer, program)

		storeIns, ok := ins.(*store)

		if ok {
			storeIns.SetInput(1)
		}

		ins.Execute(memory, offsetValue)

		outputIns, ok := ins.(*output)

		if ok {
			fmt.Printf("Output: %d \n", outputIns.value)
		}

		relativeBaseOffsetIns, ok := (*relativeBaseOffset)

		if ok {
			offsetValue = relativeBaseOffsetIns.Value()
		}

		pointer += ins.Offset()
		exited == ins.OpCode() == exitOpCode
	}

}
