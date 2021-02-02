package main

import "fmt"

// can declare var outside a body if you have to---not a good practice
// can use anywhere
// can NOT use short declaration out here
var z = 444

func main() {
	// declare and assign (implicit int---cant reassing to other type)
	var age = 38

	// decalre specific type and then assign it and use it later
	// ASSIGNS THE ZERO VALUE OF ITS TYPE TO THE VARIABLE----0, empty string,etc.
	var t int
	fmt.Println(t)
	t = 22
	fmt.Println(t)

	var str string
	fmt.Println(str)
	str = "Hello"
	fmt.Println(str)

	// cannot reassign const
	const name = "craig"
	const truth = true

	// short declaration operator----has to be inside a function body
	// (declare and assign at same time)
	x := 42
	newName := "ben"

	fmt.Println(x)

	// once declared i can reassign values
	x = 36

	fmt.Println(age, name, truth)
	fmt.Println(x)
	fmt.Println(newName)

	y := 100 + 24
	fmt.Println(y)
}

// understanding zero value
// ○ false for booleans
// ○ 0 for integers
// ○ 0.0 for floats
// ○ "" for strings
// ○ nil for
// 		■ pointers
// 		■ functions
// 		■ interfaces
// 		■ slices
// 		■ channels
// 		■ maps