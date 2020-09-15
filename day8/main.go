package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strizzwald.github.com/aoc/day8/part1"
	"strizzwald.github.com/aoc/day8/part2"
)

const layerSize = 25 * 6

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	pixels := string(content)
	numLayers := len(pixels) / layerSize
	
	// executePart1(pixels, numLayers)

	executePart2(pixels, numLayers)
}

func executePart1(pixels string, numLayers int) {
	layers := make([]part1.Layer, numLayers)
	var layerStart int

	for i := 1; i <= len(layers); i++ {
		layers[i - 1].AddPixels(string(pixels[layerStart:(layerSize * i)]))
		layerStart = layerSize * i
	}

	var layer part1.Layer
	minZeros := math.MaxInt32

	for _, l := range layers {
		if l.NumberOfZeroes() < minZeros {
			layer = l
			minZeros = l.NumberOfZeroes()
		}
	}

	fmt.Println(layer.NumberOfOnes() * layer.NumberOfTwos())
}

func executePart2(pixels string, numLayers int) {
	layers := make([]part2.Layer, numLayers)
	
	var layerStart int

	for i := 1; i <= len(layers); i++ {
		layers[i - 1].AddPixels(string(pixels[layerStart:(layerSize * i)]))
		layerStart = layerSize * i
	}

	outputLayer := [6][25]part2.Color{}

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			outputLayer[i][j] = part2.Transparent
		}
	}

	for pixelY := 0; pixelY < 6; pixelY++ {
		for pixelX := 0; pixelX < 25; pixelX++ {
			for _, l := range layers {
				var layerColor part2.Color = l.GetColor(pixelY, pixelX)

				if (outputLayer[pixelY][pixelX] == part2.Transparent) {
					outputLayer[pixelY][pixelX] = layerColor
				}
			}
		}
	}

	result := ""

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			if (outputLayer[i][j] == part2.Black) {
				result += "\u2b1c"
			}

			if (outputLayer[i][j] == part2.White) {
				result += "\u2b1b"
			}
		}

		result = result + "\n"
	}

	fmt.Println(result)
}