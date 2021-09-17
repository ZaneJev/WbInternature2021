package main

import "fmt"

func main() {
	slice := make([]interface{}, 100)
	fmt.Println(cap(slice), len(slice))
}
