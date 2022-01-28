package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]

	fmt.Println(terrestrial, gasGiants, iceGiants)
	allPlanets := planets[:]

	fmt.Println(allPlanets)
	colonized := terrestrial[2:]

	fmt.Println(colonized)

	dwarfArr := [...]string{"ceres", "pluto", "haumea", "makemake", "eris"}
	dwarfSli := dwarfArr[:]
	fmt.Printf("%T %T\n", dwarfArr, dwarfSli)
}
