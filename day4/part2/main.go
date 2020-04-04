package main

import (
	"fmt"
	"math"
)

func main() {
	start := 264793
	end := 803935
	potentialPasswordCount := 0

	for i := start; i < end; i++ {
		if digitsIncrease(i) && hasDuplicateDigits(i) {
			fmt.Println(i)
			potentialPasswordCount++
		}
	}

	fmt.Println(potentialPasswordCount)

}


func digitsIncrease(num int) bool {
	digits := num

	for digits > 0 {
		doubleDigits := digits % 100
		left := doubleDigits / 10
		right := doubleDigits % 10

		if left > right {
			return false
		}

		digits /= 10
	}

	return true
}


func hasDuplicateDigits(num int) bool {
	duplicates := make(map[int]int)
	digits := num
	totalCounted := 0

	for totalCounted < 6 {
		current := digits % 10
		counted := countDigits(digits, current)

		duplicates[current] = counted
		totalCounted += counted

		digits = digits / int(math.Pow(10, float64(counted)))
	}

	for _, v := range duplicates {
		if v == 2 {
			return true
		}
	}

	return false
}

func countDigits(num int, digit int) int {
	if num == 0 {
		return 1
	}

	digits := num
	count := 0
	position := digits % 10

	for digit == position {
		count++
		digits = digits / 10
		position = digits % 10
	}

	return count
}