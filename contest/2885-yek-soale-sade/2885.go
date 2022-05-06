package main

import(
	"fmt"
)

func main() {
	const repeatedSt string = "man khoshghlab hastam"
	var repeatTimes int
	fmt.Scanf("%d", &repeatTimes)
	for i := 0; i < repeatTimes; i++ {
		fmt.Println(repeatedSt)
	}
}
