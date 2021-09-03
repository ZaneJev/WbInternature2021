package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}

//Выводит 1 1, потому что p в функции update - копия указателя *p переданного в main, чтобы не копировалось
//Нужно на 7 строке заменить код на *p = b
