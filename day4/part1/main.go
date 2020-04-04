package main

import "fmt"

func main() {
	start := 264_793
	end := 803_935
	potentialPasswordCount := 0

	for i := start; i < end; i++ {
		if digitsIncrease(i) && hasDuplicateDigits(i) {
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
	digits := num

	for digits > 0 {
		doubleDigits := digits % 100
		left := doubleDigits / 10
		right := doubleDigits % 10

		if left == right {
			return true
		}

		digits /= 10
	}

	return false

}