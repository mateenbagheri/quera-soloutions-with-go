package main

import (
	"fmt"
)

func main() {
	var a1, b1, a2, b2, a3, b3 int

	fmt.Scanf("%d\n%d\n%d\n%d\n%d\n%d", &a1, &b1, &a2, &b2, &a3, &b3)
	result := minFunction(a1, b1) + minFunction(a2, b2) + minFunction(a3, b3)
	fmt.Println(result)
}

// minFunction get two integers and returns the one that is smaller than the
// other one. In case they contain the same value the y value will be returned.
func minFunction(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
