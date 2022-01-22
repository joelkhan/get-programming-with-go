package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	fmt.Println("\njulius caesar password:")
	caesar()
}

func caesar() {
	julius := "L fdph, L vdz, L frqtxhuhg (ABCDabcd)."
	fmt.Printf("%v\n", julius)
	for i := 0; i < len(julius); i++ {
		var c rune = rune(julius[i])

		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			//c -= 3
			c = passFmtr(c)
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func passFmtr(c rune) rune {
	var startChar rune
	if c >= 'a' && c <= 'z' {
		startChar = 'a'
	} else {
		startChar = 'A'
	}
	var offset rune = c - startChar
	var index rune = offset - 3
	// if c == 'c' {
	// 	fmt.Printf("%d %[1]T\n", offset)
	// 	fmt.Printf("%d %[1]T\n", index)
	// }
	if index < 0 {
		index = 26 + index
	}
	c = startChar + index
	return c
}
