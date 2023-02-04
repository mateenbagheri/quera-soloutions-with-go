package main

import "fmt"

func maxDoorEntryCalculator(n, m int) int {
	if n%m == 0 {
		return n / m
	} else {
		return (n / m) + 1
	}
}

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	fmt.Println(maxDoorEntryCalculator(n, m))
}
