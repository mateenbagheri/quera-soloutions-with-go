package main

import (
	"fmt"
	"strconv"
)

func main() {
	var r, c int
	fmt.Scanf("%d%d", &r, &c)
	resultSt := constructAddress(r, c)
	fmt.Println(resultSt)
}

func constructAddress(r, c int) string {
	const numberOfRows, numberOfCols int = 10, 20

	var howManyRows, howManyCols int

	howManyRows = numberOfRows + 1 - r

	resultSt := ""
	if c <= (numberOfCols / 2) {
		resultSt += "Right"
		howManyCols = c
	} else {
		resultSt += "Left"
		howManyCols = numberOfCols + 1 - c
	}
	resultSt += " " + strconv.Itoa(howManyRows) + " " + strconv.Itoa(howManyCols)
	return resultSt
}
