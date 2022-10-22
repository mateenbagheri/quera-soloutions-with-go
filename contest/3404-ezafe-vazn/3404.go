package main

import (
	"fmt"
)

func main() {
	var n int
	var m float32

	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%f", &m)

	bmi := ibmCalculator(n, m)
	condition := conditionCalculator(bmi)
	fmt.Printf("%.2f\n", bmi)
	fmt.Println(condition)
}

func ibmCalculator(weight int, height float32) float32 {
	bmi := float32(weight) / (height * height)
	return bmi
}

func conditionCalculator(bmi float32) string {
	var condition string

	switch {
	case bmi < 18.5:
		condition = "Underweight"
	case bmi >= 18.5 && bmi < 25:
		condition = "Normal"
	case bmi >= 25 && bmi < 30:
		condition = "Overweight"
	case bmi >= 30:
		condition = "Obese"
	}

	return condition
}
