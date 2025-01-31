package main

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (lsr laser) talk() string {
	return strings.Repeat("pew ", int(lsr))
}

func main() {
	var t interface {
		talk() string
	}

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}
