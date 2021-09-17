//Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

package main

import (
"fmt"
"os"
)

func Foo() error {
var err *os.PathError = nil
return err
}

func main() {
err := Foo()
fmt.Println(err)
fmt.Println(err == nil)
fmt.Printf("%T %v %T %v", err, err, nil, nil)
}

//вывод <nil> false
//false потому что тип *fs.PathError не тоже самое, что тип <nil>
//Не пустые интерфейсы используются для огранизации разных групп методов, применяемых к разным объектам,
//А пустые интерфейсы используются, как псевдогенерики, грубо говоря
