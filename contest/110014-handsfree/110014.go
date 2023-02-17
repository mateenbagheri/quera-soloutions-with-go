package main

import "fmt"

type handsfree struct {
	firstEar  string
	secondEar string
}

func checkHandsfree(first, second handsfree) bool {
	if first.firstEar == first.secondEar ||
		second.firstEar == second.secondEar {
		return true
	} else if first.firstEar == second.secondEar ||
		first.secondEar == second.firstEar {
		return true
	}
	return false
}

func main() {
	var first, second string
	fmt.Scanf("%s %s\n", &first, &second)
	firstHandsFree := handsfree{first, second}
	fmt.Scanf("%s %s\n", &first, &second)
	secondHandsfree := handsfree{first, second}
	result := checkHandsfree(firstHandsFree, secondHandsfree)
	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
