package main

import (
	"fmt"
	"sync"
)

func workers(in chan interface{}, wg *sync.WaitGroup, orderNum int) {
	defer wg.Done()
	value := <-in
	fmt.Printf("Worker №%d with value %v and type %T\n", orderNum, value, value)
}

func main() {
	var n int
	wg := &sync.WaitGroup{}
	channel := make(chan interface{})
	fmt.Print("Enter worker count: ")
	fmt.Scanf("%d\n", &n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go workers(channel, wg, i)
	}

	Data := []interface{}{"smallStr", 404, "middleString", 1419, "longestStringInTheWorld"}
	for i := 0; i < n; i++ {
		if i < len(Data) {
			channel <- Data[i]
		} else {
			channel <- i
		}

	}
	close(channel)
}
