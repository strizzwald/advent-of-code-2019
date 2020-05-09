package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type OrbitMap map[string][]string

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	orbitMap := make(OrbitMap)

	orbits := strings.Split(string(file), "\n")

	for _, o := range orbits {
		object1 := strings.Split(o, ")")[0]
		object2 := strings.Split(o, ")")[1]

		if _, ok := orbitMap[object1]; !ok {
			orbitMap[object1] = []string{object2}
		} else {
			orbitMap[object1] = append(orbitMap[object1], object2)
		}

		if _, ok := orbitMap[object2]; !ok {
			orbitMap[object2] = []string{object1}
		} else {
			orbitMap[object2] = append(orbitMap[object2], object1)
		}
	}

	myOrbitObject := orbitMap["YOU"][0]
	santasOrbitObject := orbitMap["SAN"][0]

	fmt.Println(getDistance(orbitMap, myOrbitObject, santasOrbitObject))
}

func getDistance(orbitMap OrbitMap, start string, end string) int {
	var distance int

	path := bfs(orbitMap, start, end)

	for obj := path[end]; obj != ""; obj = path[obj] {
		distance++
	}

	return distance
}

func bfs(orbitMap OrbitMap, start string, end string) map[string]string {

	queue := make([]string, len(orbitMap))
	visited := make(map[string]bool, len(orbitMap))
	path := make(map[string]string, len(orbitMap))

	queue = append(queue, start)
	visited[start] = true

	for len(queue) != 0 {
		currentObject := queue[0]

		// dequeue
		queue = queue[1:]

		objects := orbitMap[currentObject]

		for _, object := range objects {
			if !visited[object] {
				queue = append(queue, object)
				visited[object] = true
				path[object] = currentObject
			}

			if currentObject == end {
				return path
			}
		}
	}

	return path
}
