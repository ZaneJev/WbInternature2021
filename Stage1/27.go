package main

import (
	"fmt"
	"strings"
)

func reverseWordInString(sentence string) string {
	newSentence := strings.Split(sentence, " ")

	for i := 0; i < len(newSentence)/2; i++ {
		newSentence[i], newSentence[len(newSentence)-i-1] = newSentence[len(newSentence)-i-1], newSentence[i]
	}
	return strings.Join(newSentence, " ")
}

func main() {
	fmt.Println(reverseWordInString("snow dow sun"))
}
