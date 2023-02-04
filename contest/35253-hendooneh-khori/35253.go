package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// chooseLeast chooses the variable with least
// number in it along its index.
// it takes nums, which contains multiple integers
func chooseLeast(nums ...int) (int, int) {
	var leastIdx, least int = math.MaxInt, math.MaxInt
	for i, num := range nums {
		if num < least {
			least = num
			leastIdx = i
		}
	}
	return leastIdx, least
}

// returns the index of the watermelon left at the end of
// the process explained by problem
func lastWatermelonCalc(weights []int, n int) int {
	if n == 1 {
		return 0
	}

	var currentLeftIdx int = 0
	for i := 1; i < n; i++ {
		_, val := chooseLeast(weights[currentLeftIdx], weights[i])
		if val == weights[currentLeftIdx] {
			currentLeftIdx = i
		}
	}
	return currentLeftIdx + 1
}

// The StringListToIntegerList function simply gets a list
// of Strings that contain integer values and  returns a list
// of integers based on the given list values. In case a value
// is not convertable for any reasons, it will return an error.
func stringListToIntegerList(l []string) ([]int, error) {

	var lInt = make([]int, len(l))

	for idx, i := range l {
		j, err := strconv.Atoi(i)
		if err != nil {
			return lInt, err
		}
		lInt[idx] = j
	}

	return lInt, nil
}

func main() {
	var n int
	var nums []int

	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	println(line)
	nums, err := stringListToIntegerList(strings.Split(line, " "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lastWatermelonCalc(nums, n))
}
