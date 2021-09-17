//Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

package main

import (
"fmt"
)


func test() (x int) {
defer func() {
x++
}()
x = 1
return
}


func anotherTest() int {
var x int
defer func() {
x++
fmt.Println("ogo, vsetaki uvelichil", x)
}()
x = 1
return x
}


func main() {
fmt.Println(test())
fmt.Println(anotherTest())
}

//вывод 2 1
//сначала в функции test() закончит выполнение, при х = 1, затем дефер увеличит х на 1, после чего даст println х = 2
//а в функции anotherTest() сначала функция закончившись, вернет х = 1, затем дефер увеличит его на 1 (Для проверки я дописал вывод во второй дефер))

//Деферы вызываются после окончания функции. Деферы кладутся на стек, поэтому и выполняются по принципу последним зашел - первым вышел