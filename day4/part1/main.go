package main

import "fmt"

func main() {
	start := 264_793
	end := 803_935
	potentialPasswordCount := 0

	for i := start; i < end; i++ {
		hasDuplicateDigits := false
		digitsIncrease := true
		digits := i

		for j := 10; j <= 100_000; {
			x := digits % 100

			if !hasDuplicateDigits {
				hasDuplicateDigits = checkDuplicate(x)
			}

			digitsIncrease = checkDigitOrder(x)

			if !digitsIncrease {
				break
			}

			digits = i / j
			j *= 10
		}

		if hasDuplicateDigits && digitsIncrease {
			potentialPasswordCount++
		}
	}

	fmt.Println(potentialPasswordCount)
}

func checkDuplicate(num int) bool {
	leftDigit := num / 10
	rightDigit := num % 10

	return leftDigit == rightDigit
}

func checkDigitOrder(num int) bool {
	leftDigit := num / 10
	rightDigit := num % 10

	return leftDigit <= rightDigit

}