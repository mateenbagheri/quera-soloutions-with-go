package main

import (
	"fmt"
	"os"
)

func main() {
	var x1, y1, x2, y2 int
	var result string
	_, err := fmt.Scanf("%d%d%d%d", &x1, &y1, &x2, &y2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// checking the point locations for result.
	if x1 == x2 {
		result = "Vertical"
	} else if y1 == y2 {
		result = "Horizontal"
	} else {
		result = "Try again"
	}

	fmt.Println(result)
}
