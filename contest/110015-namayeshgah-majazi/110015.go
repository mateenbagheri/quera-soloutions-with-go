package main

import "fmt"

const (
	sep       = "########.......########"
	endLine   = "#######################"
	emptyLine = "#.....................#"
	filler    = "......."
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(sep)
	for i := 2; i <= 8; i = i + 2 {
		var first, second string
		if i-1 <= n {
			first = fmt.Sprintf("ghorfe%d", i-1)
		} else {
			first = filler
		}

		if i <= n {
			second = fmt.Sprintf("ghorfe%d", i)
		} else {
			second = filler
		}
		fmt.Printf("#%s%s%s#\n", first, filler, second)
		if i != 8 {
			fmt.Println(sep)
		}
	}
	fmt.Println(endLine)
}
