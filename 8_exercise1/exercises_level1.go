package main

import "fmt"

// Exercise 2, 3
var a int
var b string
var c bool

// EXERCISE 4
type hotdog int
var e hotdog

// EXERCISE 5
var f int

func main(){
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}

func ex1(){
	// either type of declaration will work := is preferred
	fmt.Println("---------------------1")
	var x = 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x, "\n", y, "\n",z)
}

func ex2(){
	fmt.Println("---------------------2")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func ex3(){
	fmt.Println("---------------------3")
	a = 42
	b = "Reddit"
	c = true
	types := fmt.Sprintf("%T, %T, %T", a, b, c)
	fmt.Println(types)
	values := fmt.Sprintf("%v, %v, %v", a, b, c)
	fmt.Println(values)
}

func ex4(){
	fmt.Println("---------------------4")
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	e = 42
	fmt.Println(e)
	fmt.Printf("%T\n", e)
}

func ex5(){
	fmt.Println("---------------------5")
	f = int(e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}