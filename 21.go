package main

import (
	"fmt"
	"sync"
)

type ConcurrentSlice struct {
	sync.RWMutex
	items []interface{}
}

type ConcurrentSliceItem struct {
	Index int
	Value interface{}
}

func (cs *ConcurrentSlice) Append(item interface{}) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}

func (cs *ConcurrentSlice) Iter() <-chan ConcurrentSliceItem {
	c := make(chan ConcurrentSliceItem)

	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			c <- ConcurrentSliceItem{index, value}
		}
		close(c)
	}
	go f()

	return c
}

func main() {
	cs := ConcurrentSlice{}

	data := []interface{}{"testConcurSlice", 12351, "someWords", 85857}
	cs.Append(data)

	fmt.Printf("%v", <-cs.Iter())
}
