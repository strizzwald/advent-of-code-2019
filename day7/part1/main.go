package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	program := loadProgram()
	executeProgram(program)
}

func loadProgram() []int {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	inputTxt := strings.Split(strings.Trim(string(input), "\n"), ",")

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

func getPhaseSettingConfigurations() [][]int {
	defaultPhases := []int{0, 1, 2, 3, 4}
	var configs [][]int

	for _, a := range defaultPhases {
		for _, b := range defaultPhases {
			for _, c := range defaultPhases {
				for _, d := range defaultPhases {
					for _, e := range defaultPhases {
						if a != b && a != c && a != d && a != e && b != c && b != d && b != e && c != d && c != e && d != e {
							configs = append(configs, []int{a, b, c, d, e})
						}
					}
				}
			}
		}
	}

	return configs
}

func executeProgram(memory []int) {
	var ampOutput int
	maxOutput := math.MinInt32
	var programCopy []int

	for _, config := range getPhaseSettingConfigurations() {
		programCopy = make([]int, len(memory))
		copy(programCopy, memory)

		for _, setting := range config {
			ampOutput = runAmplifier(programCopy, setting, ampOutput)
		}

		if ampOutput > maxOutput {
			maxOutput = ampOutput
		}

		ampOutput = 0 // reset for next phase configuration
	}

	fmt.Println(maxOutput)
}

func runAmplifier(program []int, phaseSetting int, input int) int {
	var exited bool
	var phaseSettingAdded bool
	var pointer int
	var amplifierOutput int

	for !exited {
		ins := newInstruction(pointer, program)

		storeIns, ok := ins.(*store)

		// Add phase setting
		if ok && !phaseSettingAdded {
			storeIns.SetInput(phaseSetting)
			phaseSettingAdded = true
		} else if ok && phaseSettingAdded {
			storeIns.SetInput(input)
		}

		ins.Execute(program)

		output, ok := ins.(*output)

		if ok {
			// TODO: Use a public method for *value*
			amplifierOutput = output.value
		}

		pointer += ins.Offset()
		exited = ins.OpCode() == exitOpCode
	}

	return amplifierOutput
}
