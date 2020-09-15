package part1

import (
	"fmt"
	"strconv"
)

type Layer struct {
	pixelCount map[int]int
}

func (l *Layer) AddPixels(pixels string) {
	if l.pixelCount == nil {
		l.pixelCount = map[int]int{}
	}

	for _, r := range pixels {
		num, err := strconv.Atoi(fmt.Sprintf("%c", r))

		if err != nil {
			panic(err)
		}

		if _, ok := l.pixelCount[num]; ok {
			l.pixelCount[num]++
		} else {
			l.pixelCount[num] = 1
		}
	}
}

func (l *Layer) NumberOfZeroes() int {
	if _, ok := l.pixelCount[0]; ok {
		return l.pixelCount[0]
	}

	return 0
}

func (l *Layer) NumberOfOnes() int {
	if _, ok := l.pixelCount[1]; ok {
		return l.pixelCount[1]
	}

	return 0
}

func (l *Layer) NumberOfTwos() int {
	if _, ok := l.pixelCount[2]; ok {
		return l.pixelCount[2]
	}

	return 0
}