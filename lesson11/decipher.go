package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	//fmt.Println(keyword)
	keyOffset := 0

	for _, char := range cipherText {
		keyChar := keyword[keyOffset]
		keyNum := keyChar - 'A'
		//fmt.Printf("%c%[1]T -> %c(%2d %T) ... %c\n", char, keyChar, keyNum, keyNum, decipher(char, keyNum))
		fmt.Printf("%c", decipher(char, keyNum))
		keyOffset++
		if keyOffset == len(keyword) {
			keyOffset = 0
		}
	}
	fmt.Println()

	// sample answer
	sampleDecipher()
}

func decipher(char rune, n uint8) rune {
	var num int32 = rune(n)
	res := char - num
	if res < 'A' {
		res += 26
	}
	return res
}

func sampleDecipher() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'
		message += string(c)
		keyIndex++
		keyIndex %= len(keyword)

	}
	fmt.Println(message)
}
