package main

import "fmt"

type employee struct {
	last string
	salary float64
	role string
}

func (e employee) getSalary() float64{
	return e.salary
}

func (e employee) giveRaise(x float64) float64 {
	raise := x
	e.salary = e.salary + (e.salary * raise)
	return e.salary
}


func main(){

	e1 := employee {
		last: "Sperlazza",
		salary: 1000,
		role: "Dev",
	}
	fmt.Println(e1.salary)
	
	e1Sal := e1.getSalary()
	fmt.Println(e1Sal)

	e1NewSal:= e1.giveRaise(.20)
	fmt.Println(e1NewSal)

	fmt.Println(e1.salary)

	e1 = employee{
		last: "Sperlazza",
		salary: e1NewSal,
		role: "Dev",
	}
	fmt.Println(e1.salary)

}