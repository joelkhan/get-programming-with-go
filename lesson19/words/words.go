package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fi, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	sttsRes := make(map[string]int)
	for {
		a, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Println(string(a))
		line := string(a)
		line = strings.TrimSpace(line)
		line = strings.ToLower(line)
		//fmt.Println(line)
		replacer := strings.NewReplacer(",", " ", ".", " ", ";", " ", "-", " ")
		line = replacer.Replace(line)
		fmt.Println(line)
		words := strings.Fields(line)
		fmt.Println(words)
		sttsWords(sttsRes, words)
	}

	for k, v := range sttsRes {
		if v > 1 {
			fmt.Printf("%-10v apears: %3d\n", k, v)
		}
	}

	//fmt.Printf("|%s|\n", strings.Trim("---- ", `.,"- `))
	text := ` to fear -- except the `
	fmt.Printf("%v\n", countWords(text))
}

func sttsWords(res map[string]int, words []string) {
	//fmt.Println(len(res))
	for _, word := range words {
		//fmt.Printf("%v|", word)
		res[word]++
	}
	//fmt.Println(len(res))
}

func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	freq := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Trim(word, `.,"- `)
		freq[word]++
	}
	return freq
}
