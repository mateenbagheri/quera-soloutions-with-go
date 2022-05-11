package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const safeSt string = "rahat baash"
	const dangerSt string = "nakhor lite"

	var countRed, countGreen, countYellow int
	var isDangerous bool
	var inputSt string

	_, err := fmt.Scanf("%s", &inputSt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	countRed = strings.Count(inputSt, "R")
	countGreen = strings.Count(inputSt, "G")
	countYellow = strings.Count(inputSt, "Y")

	isDangerous = checkIsDangerous(countRed, countGreen, countYellow)
	if isDangerous {
		fmt.Println(dangerSt)
	} else {
		fmt.Println(safeSt)
	}
}

func checkIsDangerous(countRed int, countGreen int, countYellow int) bool {
	if countRed >= 3 {
		return true
	} else if countRed >= 2 && countYellow >= 2 {
		return true
	} else if countYellow+countRed == 5 {
		return true
	} else {
		return false
	}
}
