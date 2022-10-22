package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numberOfQuestions, totalSum int

	fmt.Scanf("%d\n", &numberOfQuestions)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	breathes := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	volunteerNumbers := strings.Split(scanner.Text(), " ")

	for i := 0; i < numberOfQuestions; i++ {
		currentQuestionCount, _ := strconv.Atoi(volunteerNumbers[i])
		currentVolunteerCount, _ := strconv.Atoi(breathes[i])
		totalSum = totalSum + (currentQuestionCount * currentVolunteerCount)
	}

	fmt.Println(totalSum)
}
