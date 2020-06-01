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

func getPhaseSettings() [][]int {
	allowedPhases := []int{0, 1, 2, 3, 4}
	var phaseSettings [][]int

	for a := range allowedPhases {
		for b := range allowedPhases {
			for c := range allowedPhases {
				for d := range allowedPhases {
					for e := range allowedPhases {
						if a != b && a != c && a != d && a != e && b != c && b != d && b != e && c != d && c != e && d != e {
							phaseSettings = append(phaseSettings, []int{a, b, c, d, e})
						}
					}
				}
			}
		}
	}

	return phaseSettings
}

func executeProgram (memory[] int) {
	maxOutput := math.MinInt32
	var maxPhaseSetting []int
	var currentAmplifier int
	var amplifierOutput int
	var phaseSettingAdded bool

	phaseSettings := getPhaseSettings()

	pointer := 0
	exited := false

	for _, setting := range phaseSettings {
		for !exited {
			ins := newInstruction(pointer, memory)

			storeIns, ok := ins.(*store)

			// Add phase setting
			if ok && !phaseSettingAdded {
				storeIns.SetInput(setting[currentAmplifier])
				phaseSettingAdded = true
			} else if ok && phaseSettingAdded {
				// Add previous amplifier's output
				storeIns.SetInput(amplifierOutput)
			}

			ins.Execute(memory)

			output, ok := ins.(*output)

			if ok {
				// TODO: Use a public method for *value*
				amplifierOutput = output.value
				currentAmplifier++
				phaseSettingAdded = false
			}

			pointer += ins.Offset()
			exited = ins.OpCode() == exitOpCode
		}

		if maxOutput < amplifierOutput {
			maxOutput = amplifierOutput
			maxPhaseSetting = setting
		}

		currentAmplifier = 0
		exited = false
	}

	fmt.Println(maxOutput)
	fmt.Printf("Max phase setting: %v", maxPhaseSetting)
}