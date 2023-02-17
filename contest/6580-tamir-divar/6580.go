package main

import "fmt"

type point struct {
	x int
	y int
}

func getIsOutSide(upperLeftPoint point, l int, breakPoint point) bool {
	var isOutside bool
	if upperLeftPoint.x <= breakPoint.x &&
		breakPoint.x <= upperLeftPoint.x+l &&
		upperLeftPoint.y-l <= breakPoint.y &&
		breakPoint.y <= upperLeftPoint.y {
		return true
	}
	return isOutside
}

func main() {
	var x1, x2, y1, y2, l int
	fmt.Scanf("%d %d\n", &x1, &y1)
	fmt.Scanf("%d\n", &l)
	fmt.Scanf("%d %d\n", &x2, &y2)

	upperLeftPoint := point{x1, y1}
	mugPoint := point{x2, y2}
	if getIsOutSide(upperLeftPoint, l, mugPoint) {
		fmt.Println("Mahdi")
	} else {
		fmt.Println("Parsa")
	}
}
