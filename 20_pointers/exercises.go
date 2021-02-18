package main

import "fmt"

// exercise 2---create a person struct, a function with a pointer parameter, change the value stored at that address

type person struct {
	name string
	age int
}


// this takes in a value of pointer to type person
func changeMe (p *person){
	p.name = "Random Name"
	// (*p).name = "Other way"
}

func main(){
	ex1()

	fmt.Println("-------------------------EXERICSE 2----------------------------------------")

	p1 := person{
		name: "Craig",
		age: 39,
	}
	fmt.Println(p1)
	// we pass in a pointer value as the parameter
	changeMe(&p1)
	fmt.Println(p1)
}

// create a value, assign it to a variable and print the pointer address of that value
func ex1(){
	x := 3
	fmt.Println(x)
	fmt.Println(&x)
}

