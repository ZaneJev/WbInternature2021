package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 7, 3
	fmt.Println(swap(a, b))
}
