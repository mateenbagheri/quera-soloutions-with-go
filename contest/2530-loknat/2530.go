package main

import (
	"fmt"
	"strings"
)

func loknatCaseCounter(word string) int {
	var finalResult int = 1
	letterToMistakeChanceMap := map[string]int{
		"T": 2,
		"D": 2,
		"L": 2,
		"F": 2,
	}
	letters := strings.Split(word, "")
	for _, letter := range letters {
		if chance, ok := letterToMistakeChanceMap[letter]; ok {
			finalResult *= chance
		}
	}
	return finalResult
}

func main() {
	var word string
	fmt.Scanf("%s\n", &word)
	fmt.Println(loknatCaseCounter(word))
}
