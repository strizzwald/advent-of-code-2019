package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

const intersection = math.MinInt32

func (p  *Point) manhattanDistance(point Point) int {
	fmt.Println(point)

	return abs(p.x - point.x) + abs(p.y - point.y)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	grid := map[Point]int{}
	centralPort := Point{x: 0, y: 0}
	minDistance := math.MaxInt32

	line1 := strings.Split(strings.Split(string(file), "\n")[0], ",")

	line2 := strings.Split(strings.Split(string(file), "\n")[1], ",")

	grid = addWire(line1, 1, grid)
	grid = addWire(line2, 2, grid)

	for point := range grid	{
		if point != centralPort && grid[point] == intersection {
			distance := centralPort.manhattanDistance(point)

			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	fmt.Println(minDistance)
}

func addWire(wire []string, wireId int, grid map[Point]int) map[Point]int {
	currentX := 0
	currentY := 0

	for _, path := range wire {
		direction := path[0:1]
		length, _ := strconv.Atoi(path[1:])

		if direction == "L" {
			i := 0
			for ; i <= length; i++ {
				grid = updatePoint(grid, Point{x: currentX - i, y: currentY}, wireId)
			}
			currentX = currentX - length
		}

		if direction == "R" {
			i := 0
			for ; i <= length; i++ {
				grid = updatePoint(grid, Point{x: currentX + i, y: currentY}, wireId)
			}
			currentX = currentX	+ length
		}

		if direction == "U" {
			i := 0
			for ; i <= length; i++ {
				grid = updatePoint(grid, Point{x: currentX, y: currentY	- i}, wireId)
			}
			currentY = currentY - length
		}

		if direction == "D" {
			i := 0
			for ; i <= length; i++ {
				grid = updatePoint(grid, Point{x: currentX, y: currentY + i}, wireId)
			}
			currentY =  currentY + length
		}
 	}

 	return grid
}

func updatePoint(grid map[Point]int, point Point, wireId int) map[Point]int {
	if grid[point] == 0 {
		grid[point] = wireId
	}

	if grid[point] != wireId {
		grid[point] = intersection
		return grid
	}

	return grid
}
