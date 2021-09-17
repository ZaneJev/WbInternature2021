//Что выведет программа? Объяснить вывод программы.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

go func() {
	for {
		fmt.Println(<-ch)
	}
}()

defer close(ch)
time.Sleep(time.Second * 1)
}

//вывод 0 1 2 3 4 5 6 7 8 9 fatal error: all goroutines are sleep - deadlock!
//мы не контролируем канал, он заполняется 1 числом, потому что не буферизированный, затем ждем, пока цикл его освободит,
//и так как он нигде не закрывается и не контролируется, это можно исправить, как я. (sleep для того, что бы увидеть вывод,
//обычно используется sync.waitGroup.Wait)