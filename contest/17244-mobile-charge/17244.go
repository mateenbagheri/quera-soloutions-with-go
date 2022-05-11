package main

import "fmt"

func main() {
	var k int
	fmt.Scanf("%d", &k)
	fmt.Println(triangularNumber(k))
}

func triangularNumber(n int) int {
	// we know that 1, 3, 6, 10, 15, ... series is constructed the
	// same way as triangular number is which is (n * (n + 1)) / 2.
	// An alternative solution would be to just iterate through a
	// loop n times and calculate the result in that way but this is
	// way faster.
	return (n * (n + 1)) / 2
}
