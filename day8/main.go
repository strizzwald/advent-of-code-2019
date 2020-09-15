package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strizzwald.github.com/aoc/day8/part1"
)

const layerSize = 25 * 6

func main() {
	var layers []part1.Layer

	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	pixels := string(content)
	numLayers := len(pixels) / layerSize
	layers = make([]part1.Layer, numLayers)
	layerStart := 0

	for i := 1; i <= len(layers); i++ {
		layers[i - 1].AddPixels(string(pixels[layerStart:(layerSize * i)]))
		layerStart = layerSize * i
	}

	executePart1(layers)
}

func executePart1(layers []part1.Layer) {
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