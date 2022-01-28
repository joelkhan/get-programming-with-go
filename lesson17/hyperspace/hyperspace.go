package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Venus   ", "Earth  ", " Mars"}
	fmt.Println("before hyperspace:")
	fmt.Println(strings.Join(planets, ""))
	hyperspace(planets)

	fmt.Println("after hyperspace:")
	fmt.Println(strings.Join(planets, ""))
}
