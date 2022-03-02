package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 3600 * 24

func main() {
	distance := 62100000
	company := ""
	trip := ""

	fmt.Println("太空航行公司      飞行天数    飞行类型    价格（百万美元）")
	fmt.Println("=========================================================")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}
		speed := rand.Intn(15) + 16
		duration := distance / speed / secondsPerDay
		price := 20.0 + speed
		if rand.Intn(2) == 1 {
			trip = "往返"
			price *= 2
		} else {
			trip = "单程"
		}
		fmt.Printf("%-16v %6v %10v %8v%4v\n", company, duration, trip, "$", price)
	}

}
