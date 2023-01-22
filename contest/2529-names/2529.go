package main

import "fmt"

func main() {
	words := getInput()
	result := mostDifference(words)
	fmt.Println(result)
}

func getInput() []string {
	var n int
	var inputWords []string
	var tmpWord string

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &tmpWord)
		inputWords = append(inputWords, tmpWord)
	}
	return inputWords
}

func diffLetterCounter(word string) int {
	var differentLetters []string
	for _, char := range word {
		if !isAvailable(differentLetters, string(char)) {
			differentLetters = append(differentLetters, string(char))
		}
	}
	return len(differentLetters)
}

func mostDifference(words []string) int {
	var mostCount int
	for _, w := range words {
		count := diffLetterCounter(w)
		if count >= mostCount {
			mostCount = count
		}
	}
	return mostCount
}

// defining the function with a parameter of string
// type and have a return type bool
func isAvailable(alpha []string, str string) bool {

   // iterate using the for loop
   for i := 0;
   i < len(alpha);
   i++ {
      // check      
      if alpha[i] == str {
         // return true
         return true
      }
   }
   return false
}