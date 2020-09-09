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

type AmpName int

const (
	A AmpName = iota
	B AmpName = iota
	C AmpName = iota
	D AmpName = iota
	E AmpName = iota
)

type amplifier struct {
	phaseSetting     int
	program          []int
	programPointer   int
	phaseSettingUsed bool
	lastOutput       int
}

func (a *amplifier) compute(input int) int {
	for !a.done() {
		ins := newInstruction(a.programPointer, a.program)

		if output, ok := ins.(*output); ok {
			ins.Execute(a.program)
			a.programPointer += ins.Offset()
			a.lastOutput = output.value

			return a.lastOutput
		}

		if storeIns, ok := ins.(*store); ok {

			if !a.phaseSettingUsed {
				storeIns.SetInput(a.phaseSetting)
				a.phaseSettingUsed = true
			} else {
				storeIns.SetInput(input)
			}
		}

		ins.Execute(a.program)

		a.programPointer += ins.Offset()
	}

	return a.lastOutput
}

func (a *amplifier) reset(program []int) {
	a.phaseSetting = 0
	a.program = program
	a.programPointer = 0
	a.phaseSettingUsed = false
}

func (a *amplifier) done() bool {
	return a.program[a.programPointer] == exitOpCode
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

func getPhaseSettings() [][]int {
	defaultPhases := []int{5, 6, 7, 8, 9}
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
	var amplifiers = initAmplifiers(memory)

	lastOutput := 0
	maxOutput := math.MinInt32

	var maxPhaseSetting []int

	for _, settings := range getPhaseSettings() {
		for i, setting := range settings {
			amplifiers[i].phaseSetting = setting
		}

		for !amplifiers[E].done() {
			lastOutput = amplifiers[A].compute(lastOutput)
			lastOutput = amplifiers[B].compute(lastOutput)
			lastOutput = amplifiers[C].compute(lastOutput)
			lastOutput = amplifiers[D].compute(lastOutput)
			lastOutput = amplifiers[E].compute(lastOutput)
		}

		if maxOutput < lastOutput {
			maxOutput = lastOutput
			maxPhaseSetting = settings
		}

		amplifiers = initAmplifiers(memory)
		lastOutput = 0
	}

	fmt.Printf("maxOutput: %d\n", maxOutput)
	fmt.Printf("maxPhaseSetting: %v", maxPhaseSetting)
}

func initAmplifiers(memory []int) []amplifier {
	var amplifiers = make([]amplifier, 5)

	for i := 0; i < 5; i++ {
		programCopy := make([]int, len(memory))
		copy(programCopy, memory)
		amplifiers[i].reset(programCopy)
	}

	return amplifiers
}
