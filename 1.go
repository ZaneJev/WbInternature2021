package main

import "fmt"

type Human struct {
	sex  string
	name string
	age  uint
}

type Action struct {
	*Human
	operation Operation
}

type Operation struct {
	opName string
}

func (o *Operation) DoAction(opName string) {
	o.opName = opName
}

func main() {
	action := &Action{
		Human: &Human{"male", "Alex Gorodovoy", 13},
	}
	action.operation.DoAction("prygnut' vokrug sebya 3 raza i udarit' molotkom po kaske")
	fmt.Printf("%s do '%s'", action.name, action.operation.opName)
}
