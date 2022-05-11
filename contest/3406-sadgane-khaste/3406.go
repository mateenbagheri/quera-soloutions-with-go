package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const lengthOfNumbers int = 3

	var firstNo, secondNo int
	var result string = ""

	_, err := fmt.Scanf("%d\n%d", &firstNo, &secondNo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if reverseNumber(firstNo) > reverseNumber(secondNo) {
		result = strconv.Itoa(secondNo) + " < " + strconv.Itoa(firstNo)
	} else if reverseNumber(secondNo) > reverseNumber(firstNo) {
		result = strconv.Itoa(firstNo) + " < " + strconv.Itoa(secondNo)
	} else {
		result = strconv.Itoa(firstNo) + " = " + strconv.Itoa(secondNo)
	}
	fmt.Println(result)
}

// reverseNumber reverses an integer number. having a zero in the end does
// not crash the code however you will lose one digit.
func reverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
