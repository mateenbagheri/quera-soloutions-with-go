package main

import (
	"fmt"
	"os"
)

func main() {
	var tableLength int
	_, err := fmt.Scanf("%d", &tableLength)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// generating multiplication table
	for i:=1; i<=tableLength; i++ {
		for j:=1; j<=tableLength; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println()
	}
}
