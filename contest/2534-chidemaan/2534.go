package main

import "fmt"

// averageFinder finds integer average of a list of integers.
// we are returning integer results since the problem specifically
// that the result is of integer type
func averageFinder(toAverage []int) int {
	var sum int
	for _, num := range toAverage {
		sum += num
	}
	return int(sum / len(toAverage))
}

func totalDistanceCalculator(toAverage []int) int {
	var result int
	avg := averageFinder(toAverage)
	for _, current := range toAverage {
		if current-avg > 0 {
			result += current - avg
		}
	}
	return result
}

func main() {
	var numbers []int
	var current, n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &current)
		numbers = append(numbers, current)
	}
	fmt.Println(totalDistanceCalculator(numbers))
}
