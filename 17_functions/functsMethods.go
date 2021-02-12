package main

import "fmt"

type person struct{
	first string
	last string
	age int
}
type secretAgent struct {
	person
	licToKill bool
}

// going to make methods for our struct
// format for a function:
// func(r recevier) identifier(parameters) (returns(sType)) {code}
// when you have a receiver, it attaches that function to any instance of that type
func (s secretAgent) speak(){
	fmt.Println("I am ", s.first, s.last )
}




func main(){
	sa1 := secretAgent {
		person: person {
			first: "James",
			last: "Bond",
			age: 45,
		},
		licToKill: true,
	}

	sa2 := secretAgent {
		person: person{
			first: "Laura",
			last: "Moneypenny",
			age: 38,
		},
		licToKill: false,
	}

	fmt.Println(sa1.person.last)
	fmt.Println(sa2.person.last)

	// run method
	sa1.speak()
	sa2.speak()
}

