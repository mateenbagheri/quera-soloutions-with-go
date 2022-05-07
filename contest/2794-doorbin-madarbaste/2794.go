package main

import (
	"fmt"
	"os"
)

func main() {
	var pointList [3][2]int
	var resultPoint [2] int
	for i:=0; i<3; i++ {
		_, err := fmt.Scanf("%d%d", &pointList[i][0], &pointList[i][1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	resultPoint = findResultPoint(pointList)
	fmt.Println(resultPoint[0], resultPoint[1])
}


// findResultPoint gets a pointList as argument containing 3 location points and
//  by checking it's values, it returns us the 4th location point of rectangle.
func findResultPoint(pointList [3][2]int) [2]int {
	var resultPoint [2]int
	for i:=0; i<2; i++ {
		if pointList[0][i] == pointList[1][i] && pointList[1][i] != pointList[2][i] {
			resultPoint[i] = pointList[2][i]
		} else if pointList[0][i] == pointList[2][i] && pointList[1][i] != pointList[0][i] {
			resultPoint[i] = pointList[1][i]
		} else if pointList[1][i] == pointList[2][i] && pointList[0][i] != pointList[1][i]{
			resultPoint[i] = pointList[0][i]
		}
	}
	return resultPoint
}
