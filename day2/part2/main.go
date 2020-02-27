package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const addOpCode = 1
const multiplyOpCode = 2
const exitOpCode = 99

func newInstruction(pointer int, memory []int) instruction {
	opCode := memory[pointer]

	switch opCode {
	case addOpCode:
		return &add{pointer: pointer}
	case multiplyOpCode:
		return &multiply{pointer: pointer}
	case exitOpCode:
		return &exit{pointer: pointer}
	default:
		panic(fmt.Sprintf("Uknown instruction: %d", pointer))
	}
}

func main() {
	i := -1
	j := -1
	found := false

	for i < 100 && !found {
		i++
		j = -1

		for j < 100 && !found {
			j++

			memory := loadProgram()
			memory[1] = i
			memory[2] = j

			executeProgram(memory)

			found = memory[0] == 19690720
		}

	}
	fmt.Printf("found: %t, noun: %d verb: %d\n", found, i, j)
	fmt.Printf("(100 * %d) + %d: %d", i, j, (100 * i) + j)
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

func executeProgram(memory []int) {
	instructionPointer := 0

	for instructionPointer < len(memory) {

		ins := newInstruction(instructionPointer, memory)
		ins.Execute(memory)

		_, ok := ins.(*exit)

		if ok {
			break
		}

		instructionPointer += ins.Offset()
	}

}
