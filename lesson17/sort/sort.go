package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
	fmt.Printf("%T\n", planets)

	anotherSort()
}

func anotherSort() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	sort.Strings(planets)
	fmt.Println(planets)
}
