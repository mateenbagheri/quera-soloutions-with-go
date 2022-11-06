package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	const divisor float64 = 2

	fmt.Scanf("%d %d", &n, &k)

	result := float64(n)
	for i := 0; i < k; i++ {
		result /= divisor
	}

	fmt.Println(math.Floor(result))
}
