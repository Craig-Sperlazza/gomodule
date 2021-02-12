package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func(s secretAgent)speak(){
	fmt.Println("Secret Agent: I am ", s.first, s.last)
}

func(p person)speak(){
	fmt.Println("Person: I am ", p.first, p.last)
}

// we are basically saying any type that has method speak is of type human also
// Interface says: hey baby if you go that method, you are my type
// A VALUE CAN BE OF MORE THAN ONE TYPE---sa1 is type secretAgent but it is also now type human (same as p1)
// if it is an empty interface, then everything thing is of that type
// in the docs you will see functions that say: func Println(a ...interface)---this means it works for all types
type human interface {
	speak()
}

func bar(h human){
	// within an interace you must use assertion to use dot notation
	// just dumping h will work for both types but to access dot notation you must use assertion (similar to conversion)
	switch h.(type){
	case person: 
		fmt.Println("I was passed into bar and then asserted", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar and then asserted", h.(secretAgent).ltk)
	}
	fmt.Println("I was passed into bar", h)
}

func main(){
	sa1 := secretAgent {
		person: person{
			first: "James",
			last: "Bond",
		} ,
		ltk: true,
	}
	fmt.Println(sa1.first)
	fmt.Printf("%T\n", sa1)

	p1 := person{
		first: "Craig",
		last: "Sperlazza",
	}

	sa1.speak() 
	p1.speak()

	// bar takes in type human---since both sa and p have a speak() method, they are both also type human
	bar(sa1)
	bar(p1)
}