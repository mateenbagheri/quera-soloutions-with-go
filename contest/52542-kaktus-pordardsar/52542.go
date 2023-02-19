package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cactusFlower(n int) int {
	if n > 3 {
		return 1
	} else {
		return n
	}
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for _, inp := range inputs[:n] {
		inpint, err := strconv.Atoi(inp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(strings.Repeat("*", cactusFlower(inpint)))
	}
}
