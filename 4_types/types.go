package main

import "fmt"

// remember these are package scope...can put them into function scope. Should do that.
var y = 42
var z = "shaken not stirred"

// `` string literals...will include everything, spaces, double quotes, etc.
var a = `James says 
"hello there"`

var c = "Rick"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Printf("hello %s! happy coding.", c)
}
