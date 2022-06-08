package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var numberOfWords int
	var competeString string

	// Reading number of words using scanf
	fmt.Scanf("%d\n", &numberOfWords)

	// Reading next line as the string we want to reverse
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	competeString = scanner.Text()
	// Splitting the input string into array of strings
	splittedString := strings.Split(competeString, " ")

	for i := numberOfWords - 1; i >= 0; i-- {
		fmt.Print(splittedString[i] + " ")
	}

}
