package main

import "fmt"

var y = 42
// THIS IS BASICALLY GOING OVER SOME POINTS IN THE fmt PACKAGE FROM The DOCS
func main(){
	// type
	fmt.Printf("%T\n", y)
	// print the value
	fmt.Println(y)
	// binary
	fmt.Printf("%b\n", y)
	// hex
	fmt.Printf("%#x\n", y)
	// all 3
	fmt.Printf("%#x\t%T\t%b\n", y, y, y)

	// STRING PRINTING---prints to a string
	x := fmt.Sprintf("%#x\t%T\t%b", y, y, y)
	fmt.Println(x)

	// print to a file---we will dig into more later
	fmt.Fprint()
}	