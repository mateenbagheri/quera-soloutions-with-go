package main

import (
	"fmt"
	"strings"
)

const caplock string = "CAPS"

func mergeKeys(keys []string) string {
	var isCaps bool = false
	var result string
	for _, key := range keys {
		if key == caplock {
			isCaps = !isCaps
		} else {
			if isCaps {
				result += strings.ToUpper(key)
			} else {
				result += strings.ToLower(key)
			}
		}
	}
	return result
}

func main() {
	var numberOfInputs int
	var currentLetter string
	var letters []string
	fmt.Scanf("%d\n", &numberOfInputs)
	for i := 0; i < numberOfInputs; i++ {
		fmt.Scanf("%s\n", &currentLetter)
		letters = append(letters, currentLetter)
	}

	fmt.Println(mergeKeys(letters))
}
