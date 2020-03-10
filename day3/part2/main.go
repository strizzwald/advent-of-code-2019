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
	centralPoint := Point{x: 0, y: 0}
	minDistance := math.MaxInt32

	line1 := strings.Split(strings.Split(string(file), "\n")[0], ",")

	line2 := strings.Split(strings.Split(string(file), "\n")[1], ",")

	wire1 := createWire(line1)
	wire2 := createWire(line2)

	grid = addWire(wire1, 1, grid)
	grid = addWire(wire2, 2, grid)

	for point := range grid	{
		if point != centralPoint && grid[point] == intersection {
			wireDistance := steps(wire1, point) + steps(wire2, point)

			fmt.Printf("Intersection point: %d\n", point)
			fmt.Printf("Wire 1 distance: %d\n", steps(wire1, point))
			fmt.Printf("Wire 2 distance: %d\n", steps(wire2, point))
			fmt.Println()

			if wireDistance < minDistance {
				minDistance = wireDistance
			}
		}
	}

 	//fmt.Println(minDistance)

}

func addWire(wire []Point, wireId int, grid map[Point]int) map[Point]int {
	for _, value := range wire {
		grid = updatePoint(grid, value, wireId)
	}

	return grid
}

func createWire(points []string) []Point {
	currentX := 0
	currentY := 0

	var wire []Point

	for _, path := range points {
		direction := path[0:1]
		length, _ := strconv.Atoi(path[1:])

		if direction == "L" {
			i := 0

			for ; i <= length; i++ {
				wire = append(wire, Point{x: currentX - i, y: currentY})
			}
			currentX = currentX - length
		}

		if direction == "R" {
			i := 0
			for ; i <= length; i++ {
				wire = append(wire, Point{x: currentX + i, y: currentY})
			}
			currentX = currentX	+ length
		}

		if direction == "U" {
			i := 0
			for ; i <= length; i++ {
				wire = append(wire, Point{x: currentX, y: currentY	- i})
			}
			currentY = currentY - length
		}

		if direction == "D" {
			i := 0
			for ; i <= length; i++ {
				wire = append(wire, Point{x: currentX, y: currentY + i})
			}
			currentY =  currentY + length
		}
	}

	return wire
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

func steps(wire []Point, intersection Point) int {
	distance := 0

	for _, point := range wire {
		if point == intersection {
			return distance
		}

		distance += 1
	}

	return 0
}
