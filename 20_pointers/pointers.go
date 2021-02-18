package main

import "fmt"

var a = 22

func main(){
	// value
	fmt.Println(a)
	// address
	fmt.Println(&a)

	// what is stored at that address?
	b:= &a
	fmt.Println(b)
	fmt.Println(*b)

	// type int
	fmt.Printf("%T\n", a)
	// type *int which means pointer to an int
	fmt.Printf("%T\n", &a)

	// since i set b equal to a's location, i can change both values by simply changing the value at that location
	*b = 43
	fmt.Println(a)
	fmt.Println(*b)

	fmt.Println("################See foo func###############################")

	// Notice how in the foo function you change the value but only within the scope of the function
	// in the bar function you change the value at that pointer, so it changes it in all scopes
	x:= 0
	foo(x)
	fmt.Println(x)

	z := 11
	bar(&z)
	fmt.Println(z)

}

func foo(y int) {
	fmt.Println(y)
	y = 55
	fmt.Println(y)
}

func bar(r *int) {
	fmt.Println(*r)
	*r = 22
	fmt.Println(*r)
}