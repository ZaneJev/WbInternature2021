package main

import (
	"fmt"
)

func setBitwise(num int64, i int64) int64 {
	return (num | (1 << (i - 1)))
}

func setBitclear(num int64, i int64) int64 {
	return (num & ^(1 << (i - 1)))
}

func main() {
	fmt.Printf("Set to 1: %b in binary and %[1]d\n", setBitwise(10, 3))
	fmt.Printf("Set to 0: %b in binary and %[1]d\n", setBitclear(10, 3))
}
