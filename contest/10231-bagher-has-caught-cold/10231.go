package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var acceptedWords = []string{
	"HAFEZ",
	"MOLANA",
}

func hasAcceptedWords(check string, acceptances []string) bool {
	for _, accepted := range acceptances {
		if strings.Contains(check, accepted) {
			return true
		}
	}
	return false
}

func getAcceptedWords(words []string, checks []string) []int {
	var acceptedWordsIndex []int
	for i, word := range words {
		if hasAcceptedWords(word, checks) {
			acceptedWordsIndex = append(acceptedWordsIndex, i+1)
		}
	}
	return acceptedWordsIndex
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for i := 0; i < 5; i++ {
		line, _ := reader.ReadString('\n')
		lines = append(lines, line)
	}

	result := getAcceptedWords(lines, acceptedWords)
	if len(result) == 0 {
		fmt.Println("NOT FOUND!")
	} else {
		fmt.Println(strIntSlice(result))
	}
}

func strIntSlice(ints []int) string {
	var b bytes.Buffer
	for _, num := range ints {
		fmt.Fprintf(&b, "%d ", num)
	}
	return b.String()
}
