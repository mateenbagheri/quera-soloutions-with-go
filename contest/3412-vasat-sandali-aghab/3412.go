package main

import "fmt"

func main() {
	var names, entries []string
	var tmpStr1, tmpStr2, result string
	for i := 0; i < 4; i++ {
		fmt.Scanf("%s %s\n", &tmpStr1, &tmpStr2)
		if tmpStr2 != "U" {
			names = append(names, tmpStr1)
			entries = append(entries, tmpStr2)
		}
	}

	if entries[1] == entries[2] {
		result = names[1]
	} else if entries[1] != entries[2] {
		result = names[0]
	} else {
		if entries[0] == "L" {
			result = names[0]
		} else {
			result = names[2]
		}
	}

	fmt.Println(result)
}
