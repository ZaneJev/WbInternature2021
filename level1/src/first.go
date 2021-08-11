package src

type Action struct {
	Human
	repeatTimes int
	operation string
}

type Human struct {
	sex string
	name string
	age int
}

func CreateAction() string {
	man := Human{
		"male",
		"Valera",
		35,
	}
	action := Action{
		man,
		1,
		"do something",
	}
	return action.name + " " + action.operation
}
