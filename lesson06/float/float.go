package main

import "fmt"

func main() {
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)

	quickView6_6()
}

// 速查6-6
func quickView6_6() {
	piggyBank := 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)

}
