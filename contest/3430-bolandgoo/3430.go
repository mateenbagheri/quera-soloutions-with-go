package main

import (
	"fmt"
)

func main() {
	var word string

	fmt.Scanf("%s\n", &word)

	chars := []rune(word)
	length := len(chars)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j <= i {
				fmt.Print(string(chars[i]))
			} else {
				fmt.Print(string(chars[j]))
			}
		}
		fmt.Println()
	}
}
