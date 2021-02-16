package main

import "fmt"

var xi = []int{1,2,3,4,5}

type person struct {
	first string
	last string
	age int
}

// exercise 4---created a person struct above
// need a method for that struct
func (p person)nameAge(){
	fmt.Println(p.first, p.age)
}


func main(){
	fmt.Println("--------Exercise 1: 2 functions--- 1 return int, 2 return string, int-------------------")
	x := foo()
	y, z := bar()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("--------Exercise 2: variadic parameter, pass in slice-------------------")
	a := ex2(xi...)
	fmt.Println(a)

	// comment out the defer because it will run after everything, confusing
	// fmt.Println("--------Exercise 3: defer a function-------------------")
	// defer first_func_def()
	// second_func_first()


	fmt.Println("--------Exercise 4: create a struct person and attach and call a method from a value of that type-------------------")
	p1 := person {
		first: "Craig",
		last: "Sperl",
		age: 39,
	}
	p1.nameAge()
	
}

// exercise 1
func foo() int{
	return 3
}

func bar() (int, string){
	return 3, "Hello"
}

// exercise 2
func ex2(x ...int) int {
	var total int
	for _, v := range x{
		total += v
	}
	return total
}


// exercise 3
func first_func_def(){
	fmt.Println("I am called first, but defered")
}

func second_func_first(){
	fmt.Println("I am called second but NOT defered")
}


