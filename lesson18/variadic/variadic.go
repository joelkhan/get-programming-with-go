package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	fmt.Printf("%T\n", worlds)
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)

	others := []string{"Venus11", "Mars22", "Jupiter33"}
	newOthers := terraform("New00", others...)
	fmt.Println(newOthers)
}
