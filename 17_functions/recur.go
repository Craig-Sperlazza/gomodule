package main

import "fmt"

func main(){
	x := factorial(5)
	fmt.Println(x)

	z := loopFac(5)
	fmt.Println(z)

	y := summ(5)
	fmt.Println(y)
}

// use recursion to calculate a factorial
func factorial(n int) int {
	if n <= 1{
		return n
	}
	return n * factorial(n-1)
}

// use a loop to calculate a factorial
func loopFac(n int) int{
	total := 1
	for i:= n; i > 0; i--{
		total *= i
	}
	return total
}

// use recursion to calculate a summation
func summ(n int) int {
	if n <= 1{
		return n
	}
	return n + summ(n-1)
}
