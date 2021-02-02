package main

import "fmt"

var a int

type hotdog int
var b hotdog


var c int

func main(){
	a = 12
	b = 22
	c = int(b)

	fmt.Println(a+c)
}