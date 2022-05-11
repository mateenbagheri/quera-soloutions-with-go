package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const copyString string = "copy of "

	var endText string
	var howManyCopies int
	var resultSt string

	_, err := fmt.Scanf("%d%s", &howManyCopies, &endText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resultSt = strings.Repeat(copyString, howManyCopies)
	resultSt = resultSt + endText

	fmt.Println(resultSt)
}
