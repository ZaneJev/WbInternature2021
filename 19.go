package main

import "fmt"

//var justString string
const justString = "justString"

func someFunc(str *string) {
	v := createHugeString(1 << 10)
	*str = v[:100]
}

func createHugeString(len int) string {
	s := ""
	for i := 0; i < len; i++ {
		s += "a"
	}
	return s
}

func main() {
	var str string
	someFunc(&str)
	fmt.Println(str)
}

//Опасность в использование глобальной переменной, так как ее значения могут менять все функции, в которых она вызывается
//Чтобы такого избежать глобальную переменную лучше создавать константной, а если вы все же хотите изменять вашу переменную
//То создайте ее в main и передавайте по ссылке
