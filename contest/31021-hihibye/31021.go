package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var numberOfStudents int
	var studentNameMask string
	var studentNames []string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanf("%d\n", &numberOfStudents)
	scanner.Scan()
	studentNameMask = scanner.Text()

	studentNames = strings.Split(studentNameMask, " ")

	for i := 1; i < numberOfStudents; i++ {
		for j := i - 1; j >= 0; j-- {
			fmt.Println(studentNames[i]+": salam", studentNames[j]+"!")
		}
	}

	for i := 0; i < numberOfStudents; i++ {
		fmt.Println(studentNames[i]+":", "khodafez bacheha!")
		for j := i + 1; j < numberOfStudents; j++ {
			if j > numberOfStudents {
				break
			}
			fmt.Println(studentNames[j]+": khodafez", studentNames[i]+"!")
		}
	}
}
