package main

import "fmt"

func main() {
	m := make(map[string]int)
	a := []string{"cat", "cat", "dog", "cat", "tree"}

	for i := 0; i < len(a); i++ {
		if _, ok := m[a[i]]; !ok {
			m[a[i]] = 1
		} else {
			continue
		}
	}

	set := make([]string, len(m))
	inc := 0
	for key, _ := range m {
		set[inc] = key
		inc++
	}
	fmt.Println(set)
}
