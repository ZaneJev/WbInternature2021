package main

import (
	"fmt"
	"time"
)

func input(in chan int) {
	defer close(in)
	for i := 0; i < 15; i++ {
		in <- i
	}
}

func powed(in chan int, out chan int) {
	defer close(out)
	for i := range in {
		out <- 2 * i
	}
}

func output(out chan int) {
	for i := range out {
		fmt.Println(i)
	}
}

func main() {
	in, out := make(chan int), make(chan int)

	go input(in)
	go powed(in, out)
	go output(out)
	time.Sleep(1 * time.Second)
}
