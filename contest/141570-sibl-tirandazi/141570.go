package main

import "fmt"

const (
	blackPoint = "black"
	whitePoint = "white"
	outPoint   = "out"
)

func evaluateHit(hitPoint int) string {
	if hitPoint <= 10 && hitPoint >= 7 {
		return blackPoint
	} else if hitPoint <= 6 && hitPoint > 0 {
		return whitePoint
	} else {
		return outPoint
	}
}

func main() {
	var hitPoint int
	fmt.Scanf("%d\n", &hitPoint)
	fmt.Println(evaluateHit(hitPoint))
}
