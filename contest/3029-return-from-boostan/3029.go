package main

import "fmt"

const (
	rightDir = "Right"
	leftDir  = "Left"
)

type location struct {
	x int
	y int
}

func leastDistanceDirection(firstLocation location, secondLocation location) string {
	var decisionMaker location
	decisionMaker.x = secondLocation.x - firstLocation.x
	decisionMaker.y = secondLocation.y - firstLocation.y
	if decisionMaker.x < 0 {
		return leftDir
	} else {
		return rightDir
	}
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scanf("%d %d\n", &x1, &y1)
	fmt.Scanf("%d %d\n", &x2, &y2)
	fmt.Println(leastDistanceDirection(location{x1, y1}, location{x2, y2}))
}
