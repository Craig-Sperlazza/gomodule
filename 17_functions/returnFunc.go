package main

import "fmt"

// can return a function from a function
func main(){
	s1 := foo()
	fmt.Println(s1)

	// note you have to essentially catch the function and then execute it and catch the int it returns
	// this is because bar() runs and returns a function
	// then you have to run f1() to execute the returned fucntion
	f1 := bar()
	rf := f1()
	fmt.Println(rf)

	// can simplify that a bit
	f2 := bar()
	fmt.Println(f2())

}

// returns a string
func foo() string{
	s:= "Hello String"
	return s
}

// function that returns a function that returns an int

func bar() func() int{
	return func () int{
		return 42
	}
}