package main

import "fmt"

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")

	fmt.Println(planets)

	dump("1: terrestrial", terrestrial)
	dump("1: worlds", worlds)
	// 三个地址不同
	fmt.Printf("1: %p, %p, %p\n", &planets, &terrestrial, &worlds)
	fmt.Println()

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")

	fmt.Println(planets)

	dump("2: terrestrial", terrestrial)
	dump("2: worlds", worlds)

	fmt.Printf("2: %p, %p, %p\n", &planets, &terrestrial, &worlds)

	_ = worlds
}
