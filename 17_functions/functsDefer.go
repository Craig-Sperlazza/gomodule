package main

import "fmt"

func main(){
	// if we defer foo(), it runs after bar---technically right before func main() exits
	defer foo()
	bar()
}

func foo(){
	fmt.Println("foo")
}

func bar(){
	fmt.Println("bar")
}