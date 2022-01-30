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

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(laser(2))

	type crater struct{}
	// crater does not implement talker (missing talk method)
	// shout(crater{})
}
