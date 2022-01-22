package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var degrees = 0
	var cntr = 1
	for {
		fmt.Printf("%v ", degrees)
		degrees++
		if degrees >= 360 {
			fmt.Println("\n")
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			} else {
				cntr++
			}
		}
	}
	fmt.Printf("total: %v\n", cntr)
}
