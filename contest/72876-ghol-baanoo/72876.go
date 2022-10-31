package main

import "fmt"

func main() {
	var userInput string
	fmt.Scanf("%s\n", &userInput)
	result := reverseString(userInput)
	fmt.Println(result)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
