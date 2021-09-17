package main

import (
	"fmt"
	"math/rand"
)

func createData(in chan int) chan int {
	for i := 0; i < 20; i++{
		in <- rand.Intn(50)
	}
	return in
}

func main() {
	in := make(chan int, 1)
	out := make(chan int, 1)

	go func (in chan int) chan int{
		defer close(in)
		for i := 0; i < 20; i++{
			if num := rand.Intn(50); num%2 == 0{
				in <- num
			}
		}
		return in
	}(in)

	go func (in chan int, out chan int) chan int {
		defer close(out)
		for n := range in {
			out <- n
		}
		return out
	}(in, out)

	for num := range out {
		fmt.Println(num)
	}
}