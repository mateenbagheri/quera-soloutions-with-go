package main

import (
	"fmt"
	"os"
)

func main() {
	var n, p int
	var k float64
	_, err := fmt.Scanf("%d%f%d", &n, &k, &p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result := float64(n) * k * float64(p)
	fmt.Println(int(result))
}
