package main

import "fmt"

var a int
type hotdog int
var b hotdog

func main() {
	a = 42
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	b = 33
	fmt.Printf("%T\n", b)
	fmt.Println(b)
}