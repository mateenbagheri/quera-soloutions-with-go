package main

import (
	"fmt"
)

func main() {
	var x, n int

	fmt.Scanln(&x)
	fmt.Scanln(&n)

	result := calculateGrade(x, n)
	print(result)
}

func calculateGrade(x int, n int) int {
	if n == 0 {
		return 20
	} else if n == 7 {
		return x
	} else {
		if n >= x {
			return 0
		} else {
			return x - n
		}
	}
}
