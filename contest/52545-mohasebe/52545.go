package main

import "fmt"

func main() {
	var tmpInput1, tmpInput2, tmpInput3 int
	var currentTop, currentTopPlayer int = 0, 0

	const numberOfPlayers int = 4

	for i := 0; i < numberOfPlayers; i++ {
		fmt.Scanf("%d %d %d\n", &tmpInput1, &tmpInput2, &tmpInput3)

		if tmpInput1 > currentTop {
			currentTop = tmpInput1
			currentTopPlayer = i + 1
		}
		if tmpInput2 > currentTop {
			currentTop = tmpInput2
			currentTopPlayer = i + 1
		}
		if tmpInput3 > currentTop {
			currentTop = tmpInput3
			currentTopPlayer = i + 1
		}
	}

	fmt.Println(currentTopPlayer)
}
