package main

import "fmt"

// see unfurlSum in main() to show how to unfurl a slice of int when you ask for unlimited ints
var xi = []int{1,3,5,7,9}

func main(){
	func1("Craig")

	returnVal := plusTwo(4)
	fmt.Println(returnVal)

	returnMyStr := returnStr("Boss Man")
	fmt.Println(returnMyStr)

	x, y := multipleReturn("Ben", "King")
	fmt.Println(x, y)

	newSum := printReturnSum(1,2,3,4,5,6)
	fmt.Println(newSum)

	unfurlSum := unfurlSlice(xi...)
	fmt.Println(unfurlSum)
}

func func1(str string){
	fmt.Println("Hello", str)
}

func plusTwo(x int) int{
	return x + 2
}

func returnStr(str string) string {
	return fmt.Sprint("Hello from the otherside, ", str)
}

// multiple return statements
func multipleReturn (fn string, ln string) (string, bool){
	a := fmt.Sprint("My Dude ", fn, " ", ln )
	b := true
	return a, b
}

// variadic parameters
// stores them as a slice in x
func printReturnSum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0

	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
	return sum
}

// made a slice above....this is asking for ints not a slice, so we have to unfurl it when we pass it in---see main
func unfurlSlice(x ...int)int{
	sum := 0

	for _, v := range x{
		sum += v
	}
	return sum
}