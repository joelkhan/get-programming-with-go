package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"

	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)

	fmt.Println("origin msg:", plainText)
	ciphered := myCipher(plainText, keyword)
	fmt.Println("after ciphered:", ciphered)
	fmt.Println("back to origin:", sampleDecipher1(ciphered, keyword))
}

func myCipher(msg, keyword string) string {
	keyIndex := 0
	res := ""
	for i := 0; i < len(msg); i++ {
		c := msg[i] - 'A'
		k := keyword[keyIndex] - 'A'

		c = (c+k)%26 + 'A'
		res += string(c)
		keyIndex++
		keyIndex %= len(keyword)
	}
	return res
}

func sampleDecipher1(cipherText, keyword string) string {
	// cipherText := "CSOITEUIWUIZNSROCNKFD"
	// keyword := "GOLANG"
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
	//fmt.Println(message)
	return message
}
