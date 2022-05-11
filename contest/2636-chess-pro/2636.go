package main

import (
	"fmt"
	"os"
)

func main() {
	var standardChessList = [...]int{1, 1, 2, 2, 2, 8}
	var currentChessList, resultChessList [6]int

	_, err := fmt.Scanf(
		"%d%d%d%d%d%d",
		&currentChessList[0],
		&currentChessList[1],
		&currentChessList[2],
		&currentChessList[3],
		&currentChessList[4],
		&currentChessList[5],
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < len(resultChessList); i++ {
		resultChessList[i] = standardChessList[i] - currentChessList[i]
	}

	fmt.Printf(
		"%d %d %d %d %d %d",
		resultChessList[0],
		resultChessList[1],
		resultChessList[2],
		resultChessList[3],
		resultChessList[4],
		resultChessList[5],
	)
}
