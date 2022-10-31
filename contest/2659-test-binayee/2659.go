package main

import (
	"fmt"
)

func main() {
	var numberOfCharacters int
	var correctString, givenString string

	fmt.Scanf("%d\n%s\n%s\n", &numberOfCharacters, &correctString, &givenString)

	numberOfMistakes := calculateDifference(numberOfCharacters, correctString, givenString)
	fmt.Println(numberOfMistakes)
}

func calculateDifference(length int, s1 string, s2 string) int {
	// s1 = strings.ToLower(s1)
	// s2 = strings.ToLower(s2)

	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	countDiff := 0
	for i := 0; i < length; i++ {
		if s1Rune[i] != s2Rune[i] {
			countDiff++
		}
	}
	return countDiff
}
