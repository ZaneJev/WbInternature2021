//Что выведет программа? Объяснить вывод программы.

package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
{
// do something
}
return nil
}

func main() {
	var err error
	err = test()
	fmt.Printf("%T %v %T %v", err, err, nil, nil)
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

//вывод error
//тоже самое, что и в 3
//*main.customError - тип переменной err, а у nil тип - nil