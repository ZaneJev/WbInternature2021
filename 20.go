package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}

/* вывод:[b b a][a a]
Аналогично 18 заданию после вызова append памяти будет недостаточно и снова создасться новый слайс
*/
