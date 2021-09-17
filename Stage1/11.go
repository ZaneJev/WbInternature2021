package main

import "fmt"

func arrIntersec(a []int, b []int) (result []int) {
	m := make(map[int]bool)

	for _, i := range a {
		m[i] = true
	}

	for _, i := range b {
		if _, ok := m[i]; ok {
			result = append(result, i)
		}
	}
	return
}

func main() {
	var a = []int{1, 3, 5, 7, 9, 4}
	var b = []int{2, 4, 6, 8, 3, 5}

	fmt.Printf("%d\n%d\n%d\n", a, b, arrIntersec(a, b))
}
