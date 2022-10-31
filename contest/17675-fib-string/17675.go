package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		if isFibonacci(i) {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
}

func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

func isFibonacci(n int) bool {
	return isPerfectSquare(5*n*n+4) || isPerfectSquare(5*n*n-4)
}
