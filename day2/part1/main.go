package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	const add = 1
	const multiply = 2
	const exit = 99
	const offset = 4

	input, err := ioutil.ReadFile("part1/input.txt")

	if err != nil {
		panic(err)
	}

	inputTxt := strings.Split(string(input), ",")

	intcode := make([]int64, len(inputTxt))

	for i := range inputTxt {
		intcode[i], err = strconv.ParseInt(inputTxt[i], 10, 64)

		if err != nil {
			panic (err)
		}
	}

	position := 0

	for position < len(intcode) {
		opcode := intcode[position]

		if opcode == exit {
			break
		}

		instruction := intcode[position: position + offset]

		left := intcode[instruction[1]]
		right := intcode[instruction[2]]

		switch opcode {
		case add:
			intcode[instruction[3]] = left + right

		case multiply:
			intcode[instruction[3]] = left * right

		default:
			panic(fmt.Sprintf("%d is not a valid Opcode.", instruction[0]))
		}

		position += offset
	}

	fmt.Println(intcode)
}