package main

import "fmt"

func reverseString1(s string) string {
	reverted := []int32(s)

	for i := 0; i < len(reverted)/2; i++ {
		reverted[i], reverted[len(reverted)-1-i] = reverted[len(reverted)-1-i], reverted[i]
	}
	return string(reverted)
}

func reverseString2(s string) string {
	reverted := make([]int32, len(s))

	for i, elem := range s {
		reverted[len(s)-i-1] = int32(elem)
	}
	return string(reverted)
}

func main() {
	emojis := "🤎🖤🤍💜💙💚💛🧡"
	fmt.Println(reverseString1(emojis))
	fmt.Println(reverseString2(emojis))
}
