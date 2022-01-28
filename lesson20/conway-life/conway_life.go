package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width    = 80
	height   = 15
	initRate = 0.25
)

type Universe [][]bool

func main() {
	newWorld := NewUniverse()
	//fmt.Print("\033[H\033[2J")

	//fmt.Println("One new world~")
	//newWorld.Show()
	newWorld.Seed()
	//fmt.Println("After init~")
	//newWorld.Show()

	// fmt.Println("let's check (2, 3)")
	// fmt.Println(newWorld.Neighbors(2, 3))
	// fmt.Println("let's check (56, 13)")
	// fmt.Println(newWorld.Neighbors(56, 13))
	// fmt.Println("let's check (57, 13)")
	// fmt.Println(newWorld.Neighbors(57, 13))
	// fmt.Println("let's check (58, 13)")
	// fmt.Println(newWorld.Neighbors(58, 13))
	// fmt.Println("let's check (79, 14)")
	// fmt.Println(newWorld.Neighbors(79, 14))
	// fmt.Println("let's check (0, 0)")
	// fmt.Println(newWorld.Neighbors(0, 0))
	parallel := NewUniverse()

	for i := 0; i < 10; i++ {
		fmt.Print("\033[H\033[2J")
		newWorld.Show()
		Step(newWorld, parallel)
		//fmt.Println("After NEXT step~")
		//parallel.Show()
		newWorld, parallel = parallel, newWorld
		time.Sleep(time.Second * 1)
	}
}

func Step(a, b Universe) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if a.Next(i, j) {
				b[j][i] = true
			}
		}
	}
}

func NewUniverse() Universe {
	u := make([][]bool, height)
	for i := 0; i < height; i++ {
		//append(u, make([]bool, width))
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show() {
	//fmt.Printf("%v\n", u)
	for i := 0; i < len(u); i++ {
		for j := 0; j < len(u[i]); j++ {
			if u[i][j] {
				fmt.Printf("*")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	initTotal := int(width * height * initRate)
	for i := 0; i < initTotal; i++ {
		//fmt.Printf("%d ", i)
		for {
			row := rand.Intn(height)
			col := rand.Intn(width)
			if !u[row][col] {
				u[row][col] = true
				break
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	nbsAlive := 0
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			if i == x && j == y {
				//fmt.Printf("%2s,%2s ", " ", " ")
				continue
			} else {
				//fmt.Printf("%2d,%2d ", i, j)
				if u.Alive(i, j) {
					nbsAlive++
				}
			}
		}
		//fmt.Println()
	}
	return nbsAlive
}

func (u Universe) Next(x, y int) bool {
	next := false
	alive := u.Alive(x, y)
	nbsAlive := u.Neighbors(x, y)
	if alive {
		if nbsAlive == 2 || nbsAlive == 3 {
			next = true
		}
	} else {
		if nbsAlive == 3 {
			next = true
		}
	}
	return next
}
