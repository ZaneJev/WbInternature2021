package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr
	i := 3
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println(slice)
}
