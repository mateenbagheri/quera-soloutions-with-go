package main

import "fmt"

func main() {
	var firstInput, secondInput string

	fmt.Scanf("%s\n%s", &firstInput, &secondInput)

	firstToCompare := string(firstInput[0])
	secondToCompare := string(secondInput[len(secondInput)-1])

	if firstToCompare == secondToCompare {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
