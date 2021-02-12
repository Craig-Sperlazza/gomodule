package main

import "fmt"

// closure---enclosing scope of a varaible to a specific scope

// package level scope
var x int

func main(){
	closure()
	
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func closure(){
	// function scope
	y := 3
	fmt.Println(y)

}

// func foo is a function that returns a function that returns an int
func incrementor() func() int {
	var x int
	return func() int {
		x ++
		return x
	}
}