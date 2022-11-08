package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const numberOfEmployees int = 3

	var everyDay = []string{
		"shanbe",
		"1shanbe",
		"2shanbe",
		"3shanbe",
		"4shanbe",
		"5shanbe",
		"jome",
	}

	daysToChangeProblem := everyDay

	employeesSchedule := make([][]string, numberOfEmployees)

	for i := 0; i < numberOfEmployees; i++ {
		var numberOfDays int
		var presentDaysInString string
		var presentDays []string

		fmt.Scanf("%d\n", &numberOfDays)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		presentDaysInString = scanner.Text()

		presentDays = strings.Split(presentDaysInString, " ")
		employeesSchedule[i] = presentDays
	}

	for i := 0; i < numberOfEmployees; i++ {
		daysToChangeProblem = difference(daysToChangeProblem[:], employeesSchedule[i])
	}

	fmt.Println(len(daysToChangeProblem))
}

// This function calculates difference from two slices. 
// This code is initially from this reference:
// https://stackoverflow.com/a/45428032/12031411
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

