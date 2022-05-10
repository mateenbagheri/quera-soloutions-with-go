package main

import "fmt"

func main() {
	var howManySins int
	sevenSin := [7]string{
		"sib",
		"samanoo",
		"something",
		"somayeh",
		"sword",
		"sorrow",
		"surrender",
	}
	fmt.Scanf("%d", &howManySins)
	for i := 0; i < howManySins; i++ {
		fmt.Println(sevenSin[i])
	}
}
