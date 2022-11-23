package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var rowCount, colCount int
	fmt.Scanf("%d %d\n", &rowCount, &colCount)

	scanner := bufio.NewScanner(os.Stdin)

	skyMap := make([][]string, rowCount)

	for i := 0; i < rowCount; i++ {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		currentSkyRow := strings.Split(scanner.Text(), "")
		skyMap[i] = currentSkyRow
	}

	fmt.Println(countStars(skyMap))
}

func countStars(mtx [][]string) int {
	starsCount := 0
	for i, a := range mtx {
		for j := range a {
			if mtx[i][j] == "*" {
				starsCount++	
			}
		}
	}
	return starsCount
}