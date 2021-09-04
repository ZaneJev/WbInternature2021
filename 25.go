package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Mu    sync.RWMutex
	Count uint
}

func (c *Counter) Nullify() {
	c.Count = uint(0)
}

func (c *Counter) Add(num uint) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Count += num
	if c.Count == 4294967294 {
		c.Nullify()
	}
	fmt.Println(c.Count)
}

func Increment(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Add(1)
}

func main() {
	wg := &sync.WaitGroup{}
	c := &Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Increment(c, wg)
	}
	wg.Wait()
}
