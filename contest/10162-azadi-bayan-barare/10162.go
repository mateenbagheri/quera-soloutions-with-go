package main

import (
	"fmt"
	"os"
)

func main() {
	var k int
	_, err := fmt.Scanf("%d", &k)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if k % 2 == 1 {
		fmt.Println("Payin Barare")
	} else {
		fmt.Println("Bala Barare")
	}
}
