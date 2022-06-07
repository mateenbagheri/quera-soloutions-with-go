package main

import (
	"fmt"
	"strconv"
)

func actualTimeCalculator(mirrorHour int, mirrorMinitue int) (int, int) {
	var actualHour, actualMinitue int
	const maxMinitue, maxHour int = 60, 12

	if mirrorHour == 0 {
		actualHour = 0
	} else {
		actualHour = maxHour - mirrorHour
	}

	if mirrorMinitue == 0 {
		actualMinitue = 0
	} else {
		actualMinitue = maxMinitue - mirrorMinitue
	}

	return actualHour, actualMinitue
}

func numberToFormattedStringConverter(number int) string {
	var resultString string

	resultString = strconv.Itoa(number)
	if len(resultString) == 1 {
		resultString = "0" + resultString
	}
	return resultString
}

func main() {
	var mirrorMinitue, mirrorHour int
	var actualMinitue, actualHour int

	fmt.Scanf("%d%d", &mirrorHour, &mirrorMinitue)

	actualHour, actualMinitue = actualTimeCalculator(mirrorHour, mirrorMinitue)
	fmt.Println(
		numberToFormattedStringConverter(actualHour) +
			":" +
			numberToFormattedStringConverter(actualMinitue))
}
