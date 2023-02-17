package main

import (
	"fmt"
	"strconv"
)

const maxInt = 9223372036854775807

func hasBiggerSmall(numberString string) (*string, error) {
	initialValue, err := strconv.Atoi(numberString)
	if err != nil {
		return nil, err
	}

	// setting initial value
	var result int = maxInt

	perms := generatePerms(numberString)
	for i := range perms {
		current, err := strconv.Atoi(perms[i])
		if err != nil {
			return nil, err
		}
		if current > initialValue && current < result {
			result = current
		}
	}

	if result == initialValue || result == maxInt {
		result = 0
	}

	resultStr := strconv.Itoa(result)
	return &resultStr, nil
}

func generatePerms(toPerm string) []string {
	if len(toPerm) == 0 {
		return []string{""}
	}

	var perms []string
	for i := 0; i < len(toPerm); i++ {
		char := string(toPerm[i])
		rest := toPerm[:i] + toPerm[i+1:]
		for _, perm := range generatePerms(rest) {
			perms = append(perms, char+perm)
		}
	}
	return perms
}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	result, err := hasBiggerSmall(input)
	if err != nil {
		fmt.Printf("something wen wrong!, err: %s", err.Error())
	}
	fmt.Println(*result)
}
