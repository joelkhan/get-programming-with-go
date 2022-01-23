package main

import (
	"fmt"
)

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32.0)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string, string)

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		c1, c2 := getRow(row)
		fmt.Printf(rowFormat, c1, c2)
	}
	fmt.Println(line)
}

func c2f(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	c1 := fmt.Sprintf(numberFormat, c)
	c2 := fmt.Sprintf(numberFormat, f)
	return c1, c2
}

func f2c(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	c1 := fmt.Sprintf(numberFormat, f)
	c2 := fmt.Sprintf(numberFormat, c)
	return c1, c2
}

func main() {
	drawTable("º C", "º F", 29, c2f)
	fmt.Println()
	drawTable("º F", "º C", 29, f2c)
}
