package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println(p.first, "is", p.age, "years old.")
}

func main(){
	p1 := person {
		first: "Craig",
		last: "Sperl",
		age: 39,
	}

	p1.speak()
	fmt.Println(p1.last)
}