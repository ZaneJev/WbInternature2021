package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

//вывод [77, 78, 79]
//Просто вывод слайса из 3 элементов с 1 индекса массива "а" по 4 индекс