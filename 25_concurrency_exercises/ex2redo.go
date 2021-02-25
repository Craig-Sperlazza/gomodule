package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak()
}


func main(){
	fmt.Println(2)
	p1 := person {
		name: "Craig",
	}

	saySomethingLoud(&p1)
	p1.speak()
}

func (p *person) speak(){
	fmt.Println(p.name, "is awesome")
}

func saySomethingLoud(h human){
	fmt.Println(h.(*person).name, "IS SCREAMING")
	h.speak()
}