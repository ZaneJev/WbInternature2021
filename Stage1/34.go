package main

import "fmt"

func unique(s string) bool {
	symMap := make(map[rune]bool)
	for _, sym := range s {
		if _, ok := symMap[sym]; ok {
			return false
		} else {
			symMap[sym] = true
		}
	}
	return true
}

func main() {
	s := "zdarova, pravoslavnyie"
	str := "ya uniqUe"
	fmt.Println(unique(s), unique(str))
}
