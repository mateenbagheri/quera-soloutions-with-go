package main

import (
	"fmt"
)

var completeGrade int = 20

func main() {
	var x, n int

	fmt.Scanf("%d\n", &x)
	fmt.Scanf("%d\n", &n)

	result := calculateGrade(x, n)
	print(result)
}

func calculateGrade(currentGrade int, travelDays int) int {
	var finalGrade int = currentGrade

	if travelDays == 0 {
		finalGrade = completeGrade
	} else if travelDays != 7 {
		finalGrade = currentGrade - travelDays
		if finalGrade < 0 {
			finalGrade = 0
		}
	}

	return finalGrade
}
