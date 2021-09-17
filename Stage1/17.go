package main

import "fmt"

func runtimeType(value interface{}) {
	fmt.Printf("%T\n", value)
}

func main() {
	data := []interface{}{2, "string", 'b'}
	for _, v := range data {
		runtimeType(v)
	}
}
