package main

import (
	"fmt"
)

type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planetsArr := [...]string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	planets := planetsArr[:]
	fmt.Printf("1: arr: %p, sli: %p\n", &planetsArr, &planets)

	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)
	fmt.Printf("2: arr: %p, sli: %p\n", &planetsArr, &planets)

	otherPlanets()
}

func otherPlanets() {
	others := [...]string{
		"Earth", "Mars",
	}
	fmt.Println(others)
	others[1] = "Jupiter"
	fmt.Println(others)
}
