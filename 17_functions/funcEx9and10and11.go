package main

import "fmt"

var xi = []int{1, 2, 3, 4, 5, 7, 6, 8, 10}

var s1 = "racecar"
var s2 = "r"
var s3 = "raceecar"
var s4 = "rr"
var s5 = "hello"

func main(){
	fmt.Println("-------------------CALLBACK EXERCISE----------------------------------------------------------------")
	// callback---pass a function in as an argument
	b := evenNum(xi)

	fmt.Println(b, "Our even number function is working")

	c:= double(evenNum, xi)

	fmt.Println("The total of all even Numbers is", c)
	
	fmt.Println("-------------------CLOSURE EXERCISE----------------------------------------------------------------")

	// closure
	// note we can use the same name here and in the function because total is enclosed in two different scopes
	total := "The total for closure function: a"

	a := closeureTotal(xi...)

	fmt.Println(total, a)

	fmt.Println("-------------------RECURSION EXERCISE----------------------------------------------------------------")
	// factorial of 10 = 3628800
	xx := factorial(10)
	fmt.Println("The factorial of your number is:", xx)

	// sum of first 10 numbers = 55
	xy := summ(10)
	fmt.Println("The sum of your number is:", xy)

	// var s1 = "racecar"
	// var s2 = "r"
	// var s3 = "raceecar"
	// var s4 = "rr"
	// var s5 = "hello"
	fmt.Println(palindrome(s1))
	fmt.Println(palindrome(s2))
	fmt.Println(palindrome(s3))
	fmt.Println(palindrome(s4))
	fmt.Println(palindrome(s5))

}

// Exercise 9: callback---pass a function into a function
// lets make a function that sorts even numbers---we will pass that into a function that doubles the numbers
func evenNum(x []int) []int {
	var returnSlice []int
	for _, v := range x {
		if(v % 2 == 0){
			returnSlice = append(returnSlice, v)
		}
	}
	return returnSlice
}

// This function takes in a function (that function takes in a slice of ints and returns a slice of ints) 
// This function then returns a total of the slice returned from the passed in func
func double(f func(y []int) []int, z []int) int {
	total := 0

	funcResult := f(z)

	for _, v := range funcResult{
		total += v
	}
	return total
}




// Exercise 10: closure---close a vaiable in a scope
func closeureTotal(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}

// exercise 11----make some recursive functions
// factorial
func factorial (x int) int{
	if (x == 1){
		return x
	} 
	return x * factorial(x-1)
}

// summation
func summ (x int) int{
	if (x == 1){
		return x
	}
	return x + summ(x-1)
}

// palindrome
func palindrome(x string) bool{
	xLen := len(x)
	if (xLen <= 1){
		return true
	} 
	if x[0] != x[xLen-1]{
		return false
	}
	x = x[1:(xLen-1)]
	return palindrome(x)
}