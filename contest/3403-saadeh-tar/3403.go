package main

import (
	"fmt"
)

func main() {
	var currentNumber int
	var average float32
	var sum int
	var product, min, max int = 1, 1001, -1001

	const count int = 4

	for i := 0; i < count; i++ {
		fmt.Scanf("%d\n", &currentNumber)

		if currentNumber < min {
			min = currentNumber
		}

		if currentNumber > max {
			max = currentNumber
		}

		product = product * currentNumber
		sum = sum + currentNumber
	}

	average = float32(sum) / float32(count)

	fmt.Printf("Sum : %.6f\n", float32(sum))
	fmt.Printf("Average : %.6f\n", average)
	fmt.Printf("Product : %.6f\n", float32(product))
	fmt.Printf("MAX : %.6f\n", float32(max))
	fmt.Printf("MIN : %.6f\n", float32(min))

}
