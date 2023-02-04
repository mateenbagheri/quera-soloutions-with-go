package main

import (
	"fmt"
	"strings"
)

var vowels []string = []string{"a", "e", "i", "o", "u"}

func pronunciationCombinationCalculator(word string) int {
	var cases int = 1
	letters := strings.Split(word, "")
	for _, letter := range letters {
		if contains(vowels, letter) {
			cases *= 2
		}
	}
	return cases
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	var word string
	fmt.Scanf("%s\n", &word)
	fmt.Println(pronunciationCombinationCalculator(word))
}
