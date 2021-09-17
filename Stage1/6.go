package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	//tryCloseFirst()
	//tryCloseSecond()
	tryCloseThird()

}

//close by WaitGrop
func tryCloseFirst() {
	ch := make(chan struct{})
	wg := &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		for {
			select {
			case <-ch:
				return
			}
		}
	}()
}

//close by close()
func tryCloseSecond() {
	ch := make(chan struct{})
	wg := &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		close(ch)
	}()
}

//close by context
func tryCloseThird() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	child := func(ctx context.Context) <-chan int {
		inc := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case ch <- inc:
					inc++
					fmt.Println("POVYSHEN")
				}
			}
		}()
		return ch
	}

	for i := range child(ctx) {
		if i == 5 {
			break
		}
	}
}
