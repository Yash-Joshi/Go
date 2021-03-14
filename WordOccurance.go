package main

import (
	"fmt"
	"strings"
)

// Print out how many times each word appears in the text
//Use string.Fields to split to words and
//string.ToLower to convert to lowercase

func main() {
	text :=
		`Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	s := strings.Fields(text)

	wordMap := map[string]int{} //Empty Map

	for _, word := range s {
		wordMap[strings.ToLower(word)]++
	}

	fmt.Println(wordMap)

}
