package main

import "fmt"

func reverseWordInString(sentence []string) []string {
	newSentence := make([]string, len(sentence))
	for i := 0; i < len(sentence); i++ {
		newSentence[len(sentence)-i-1] = sentence[i]
	}
	return newSentence
}

func main() {
	fmt.Println(reverseWordInString([]string{"snow", "dow", "sun"}))
}
