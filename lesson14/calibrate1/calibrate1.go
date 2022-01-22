package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(realSensor, offset)
	for i := 0; i < 10; i++ {
		fmt.Println("using: ", offset)
		fmt.Println(sensor())
		fmt.Println(i)
		offset += kelvin(float64(i + 1))
	}

}
