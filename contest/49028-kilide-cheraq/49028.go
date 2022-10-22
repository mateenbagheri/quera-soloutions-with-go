package main

import "fmt"

func main() {
	var numberOfInputs, currentInput, meta, count int

	fmt.Scanf("%d\n", &numberOfInputs)

	for i := 0; i < numberOfInputs; i++ {
		fmt.Scanf("%d\n", &currentInput)
		if i == 0 {
			meta = currentInput
		} else {
			if currentInput != meta {
				meta = currentInput
				count++
			}
		}
	}
	fmt.Println(count)
}
