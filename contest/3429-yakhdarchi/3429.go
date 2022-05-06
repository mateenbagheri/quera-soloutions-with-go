package main

import (
	"fmt"
	"os"
)

func main() {
	var temperature int
	_, err := fmt.Scanf("%d", &temperature)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if temperature > 100 {
		fmt.Println("Steam")
	} else if temperature < 0 {
		fmt.Println("Ice")
	} else {
		fmt.Println("Water")
	}
}
