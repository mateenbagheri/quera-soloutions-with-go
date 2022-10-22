package main

import "fmt"

func main() {
	var x1, x2 int
	fmt.Scanf("%d %d", &x1, &x2)

	if x1 == x2 {
		fmt.Println("Saal Noo Mobarak!")
	} else if x1 < x2 {
		if x2-x1 == 1 {
			fmt.Println("R")
		} else if x2-x1 == 2 {
			fmt.Println("RR")
		}
	} else {
		if x1-x2 == 1 {
			fmt.Println("L")
		} else if x1-x2 == 2 {
			fmt.Println("LL")
		}
	}
}
