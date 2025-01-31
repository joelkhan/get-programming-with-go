package main

import (
	"fmt"
	"sort"
)

func main() {
	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}

	fmt.Println(set)

	unique := make([]float64, 0, len(set))
	fmt.Printf("1: %p\n", &unique)
	for t := range set {
		unique = append(unique, t)
	}
	fmt.Printf("2: %p\n", &unique)
	sort.Float64s(unique)
	fmt.Printf("3: %p\n", &unique)
	fmt.Println(unique)
}
