package main

import "fmt"

func isXLarge(t1, t2, p1, p2 int) bool {
	if t1 < p1 || t2 < p2 {
		return false
	}
	return true
}

func main() {
	var t1, t2, p1, p2 int
	fmt.Scanf("%d %d\n", &t1, &t2)
	fmt.Scanf("%d %d\n", &p1, &p2)
	if isXLarge(t1, t2, p1, p2) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
