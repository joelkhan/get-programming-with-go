package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0, -36.7,
	}

	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		//fmt.Printf("%v, %[1]f\n", g)
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v(%[1]T): %v\n", g, temperatures)
	}
}
