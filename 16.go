package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1
		n++
		fmt.Println(n)
	}
	fmt.Println(n)
	//выведется 0, потому что у if другая переменная n, внутри своего блока видимости
	//если мы перенесем fmt.Println(n) внутрь блока if, тогда увидим там вывод 2
}
