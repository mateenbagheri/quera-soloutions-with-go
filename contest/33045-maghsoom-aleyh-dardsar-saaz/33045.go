package main

import (
	"fmt"
)

func main() {
	var n, numberOfDivisors, divisorsSum int

	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				numberOfDivisors += 1
				divisorsSum += j
			}
		}
	}
	fmt.Println(numberOfDivisors, divisorsSum)
}
