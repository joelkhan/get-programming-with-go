package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	someAnothers()
}

func someAnothers() {
	somePlanets := []string{
		"Jupiter",
		"Saturn",
	}
	fmt.Println("some anothers:", len(somePlanets))
}
