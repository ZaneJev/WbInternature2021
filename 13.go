package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

//вылетит дедлок, потому что в каждой горутине будет копия wg, что мешает каждой отслеживать все горутины и все ломается
//Чтобы такого не произошло, нужно создавать на 9 строке wg := &sync.WaitGroup{}
//и соответсвенно в горутину передавать wg *sync.WaitGroup
