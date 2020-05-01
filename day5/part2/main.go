package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	memory := loadProgram()
	executeProgram(5, memory)

	fmt.Println(memory)
}

func loadProgram() []int {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	inputTxt := strings.Split(string(input), ",")

	memory := make([]int, len(inputTxt))

	for i := range inputTxt {
		memory[i], err = strconv.Atoi(inputTxt[i])

		if err != nil {
			panic(err)
		}
	}

	return memory
}

func newInstruction(pointer int, memory []int) instruction {
	opCode := memory[pointer] % 100

	switch opCode {
	case storeOpCode:
		return &store{pointer:pointer}
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

func executeProgram(input int, memory[] int) {
	pointer := 0
	exited := false

	for !exited {
		ins := newInstruction(pointer, memory)

		storeIns, isStore := ins.(*store)

		if isStore {
			storeIns.SetInput(input)
		}

		ins.Execute(memory)

		outputIns, isOutput := ins.(*output)

		if isOutput {
			fmt.Printf("Output: %d \n", outputIns.value)
		}

		pointer += ins.Offset()
		exited = ins.OpCode() == exitOpCode
	}
}