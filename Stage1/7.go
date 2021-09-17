package main

import (
	"fmt"
	"sync"
)

func main() {
	Data := []interface{}{"ogo", "vot", "eto", "paket", "dannyh", "chtob", "ya", "takoe", "parsil", 4}
	m := sync.Map{}
	for i := 0; i < len(Data); i++ {
		m.Store(i, Data[i])
	}

	for i := 0; i < len(Data); i++ {
		fmt.Println(m.Load(i))
	}
}
