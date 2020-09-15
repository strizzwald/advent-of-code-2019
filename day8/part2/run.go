package part2

import (
	"fmt"
	"strconv"
)

type Color int

const (
	Black Color = iota
	White Color = iota
	Transparent Color = iota
)

const layerWidth = 25
const layerHeight = 6

type Layer struct {
	pixels [6][25]Color
}

func (l *Layer) AddPixels(pixels string) {

	currentColumn := 0

	for i := 0; i < 6; i++ {
		for j := i * layerWidth; j < ((i * layerWidth) + 25); j++ {
			num, err := strconv.Atoi(fmt.Sprintf("%c", pixels[j]))

			l.pixels[i][currentColumn] = Color(num)

			if err != nil {
				panic(err)
			}

			currentColumn++
		}
		currentColumn = 0
	}
}

func (l *Layer) GetColor(i int, j int) Color {
	return l.pixels[i][j]
}