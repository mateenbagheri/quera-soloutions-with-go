package main

import (
	"fmt"
	"os"
	"strconv"
)

func makeTakRaghami(n int) int {
	// makeTakRaghami is the main function of the program which 
	// performs in a recursive manner. 
	// It uses sumDigits fuction to calculate it's next input.
	// It'll keep summing the digits till there's one digit left.
	if len(strconv.Itoa(n)) == 1{
		return n
	} else {
		return makeTakRaghami(sumDigits(n))
	}
}

func sumDigits(n int) int {
	// sumDigits is a simple function which uses / and % operators
	// to return sum of all digits in an integer.
	sum := 0
	lengthOfN := len(strconv.Itoa(n))
	for i:=1; i<=lengthOfN; i++ {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(makeTakRaghami(n))
}
