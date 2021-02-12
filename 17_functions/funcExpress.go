package main

import "fmt"

// function expression is a function assigned to a variable
func main(){
	fe := func() {
		fmt.Println("My First func expression")
	}

	fe()

	meaningLife := func(x int){
		fmt.Println("The meaning of life is:", x)
	}

	meaningLife(42)
}