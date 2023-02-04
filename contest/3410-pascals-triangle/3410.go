package main

import (
	"fmt"
)

func pascalsTriangle(n int) {
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			fmt.Printf("%d ", binomial(j-1, i))
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	pascalsTriangle(n + 1)
}

func binomial(n, k int) int {
	if n < 0 || k < 0 {
		return 0
	}
	if n < k {
		return 0
	}
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}
