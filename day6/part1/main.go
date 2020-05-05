package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type OrbitMap map[string]*Object

type Object struct {
	name string
	orbits *Object
}

func createObject(name string) *Object {
	return &Object{name: name, orbits: nil}
}

func createOrbit(orbitMap OrbitMap, parent *Object, child *Object) {
	if orbitMap[parent.name] == nil {
		orbitMap[parent.name] = parent
	}

	if orbitMap[child.name] == nil {
		orbitMap[child.name] = child
	}

	orbitMap[child.name].orbits = orbitMap[parent.name]
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	orbits := strings.Split(string(file), "\n")
	orbitMap := make(OrbitMap)

	for _, orbit := range orbits {

		object1 := strings.Split(orbit, ")")[0]
		object2 := strings.Split(orbit, ")")[1]

		createOrbit(orbitMap, createObject(object1), createObject(object2))
	}

	fmt.Println(getTotalOrbits(orbitMap))
}

func getTotalOrbits(orbitMap OrbitMap) int{
	var total int

	for _, object := range orbitMap {
		parent := object.orbits

		for parent != nil {
			total++
			parent = parent.orbits
		}
	}

	return total
}
