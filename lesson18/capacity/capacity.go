package main

import (
	"fmt"
)

func main() {
	s := []string{}
	lastCap := cap(s)
	fmt.Printf("init len: %d; cap: %d\n", len(s), cap(s))

	for i := 0; i < 1000; i++ {
		s = append(s, "An element")
		if cap(s) != lastCap {
			fmt.Println(cap(s))
			lastCap = cap(s)
		}
	}
}
