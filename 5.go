package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func timer(n time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(n*time.Second))
	ch := make(chan interface{})

	go func(ch chan interface{}, ctx context.Context) {
		Data := []interface{}{"lubit", "ne lubit"}
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				ch <- Data[rand.Intn(len(Data))]
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}(ch, ctx)

	go func(ch chan interface{}, ctx context.Context) {
		for {
			select {
			case value := <-ch:
				fmt.Printf("Read value %v from channel\n", value)
			case <-ctx.Done():
				return
			}
		}
	}(ch, ctx)
	<-ctx.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	var n time.Duration
	fmt.Println("Enter time to work in seconds: ")
	fmt.Scan(&n)
	wg.Add(1)
	go timer(n, wg)
	wg.Wait()
}
