package main

import (
	"fmt"
	"os"
)

func main() {
	var numberOfNames int
	var nameMap map[string]int = make(map[string]int)

	_, err := fmt.Scanf("%d", &numberOfNames)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i:=0; i<numberOfNames; i++ {
		var name, familyName string
		var nameAlreadyExists bool

		_, err := fmt.Scanf("%s%s", &name, &familyName)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
 
		// checking if name already in map. If it does, we 
		// need to increment it's value. If not, we need to
		// add it with value 1.
		_, nameAlreadyExists = nameMap[name]
		if nameAlreadyExists == true {
			howManyRepeats := nameMap[name]
			howManyRepeats += 1
			nameMap[name] = howManyRepeats
		} else {
			nameMap[name] = 1
		}
	}

	// finding the key with largest value.
	var mostRepeatedNumber int  // initial value is set to zero by default
	for _, value := range nameMap {
		if value > mostRepeatedNumber {
			mostRepeatedNumber = value
		}
	}
	
	fmt.Println(mostRepeatedNumber)
}
