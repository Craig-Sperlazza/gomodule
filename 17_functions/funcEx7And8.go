package main

import "fmt"

func main (){
	a:= sayHello()
	fmt.Println(a)

	b := 6

	c := returnDoubleFunc(b)
	fmt.Println(c())

}

func sayHello() string{
	return "Hello My Dude"
}

func returnDoubleFunc(x int) func() int{
	return func() int {
		return 2 * x
	}
}