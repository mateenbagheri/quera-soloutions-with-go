package main

import (
	"fmt"
	"os"
)

func main() {
	var numberSlice []int
	var inputNumber int
	var allInputsAdded bool = false

	for allInputsAdded != true {
		_, err := fmt.Scanf("%d\n", &inputNumber)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if inputNumber == 0 {
			allInputsAdded = true
		} else {
			numberSlice = append(numberSlice, inputNumber)
		}
	}

	for i := len(numberSlice) - 1; i >= 0; i-- {
		fmt.Println(numberSlice[i])
	}
}
