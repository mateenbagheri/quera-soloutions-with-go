package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	const resultBeginning string = "W"
	const resultEnding string = "w!"
	const repeatedSt string = "o"

	// numberOfReapeats stores the number of times we want to
	// "Reapeat" the repeatedSt.
	// Default value is set to zero
	var numberOfRepeats int
	_, err := fmt.Scanf("%d", &numberOfRepeats)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultBeginning + strings.Repeat(repeatedSt, numberOfRepeats) + resultEnding)
}
