package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func (p *person) speak(){
	fmt.Println(p.name, "is", p.age, "years old")
}

type human interface{
	speak()
}

func realLoud(h human){
	fmt.Println(h.(*person).name, "IS SCREAMING AND")
	h.speak()
}

func main(){

	p1 := person {
		name: "Craig",
		age: 39,
	}

	fmt.Println(p1.age)

	// to call the function it will work person but to call the interface of a method that takes a pointer you maust pass the pointer
	p1.speak()
	realLoud(&p1)
}