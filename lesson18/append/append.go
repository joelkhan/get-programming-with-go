package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("1: %p\n", &dwarfs)

	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	fmt.Printf("2: %p\n", &dwarfs)

	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
	fmt.Printf("3: %p\n", &dwarfs)
}
