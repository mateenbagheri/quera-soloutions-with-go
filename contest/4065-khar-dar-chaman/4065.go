package main

import "fmt"

func main() {
	var a, b, l, result int

	fmt.Scanf("%d%d%d", &a, &b, &l)

	if l % 2 == 0 {
		result = (a+b)*(l/2)
	} else {
		result = (a+b)*((l-1)/2) + a
	}
	fmt.Println(result)
}
