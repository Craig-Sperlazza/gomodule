package main

import "fmt"

var xi = []int{1,2,3,4,5,6,7,8,10}

func main(){
	// this function returns triple the value of each value in a slice
	b := returnTriple(xi...)
	fmt.Println(b)

	// this anonymous fucntion then sorts out even
	func(sli ...int){
		for i:= 0; i <len(sli); i++{
			if (sli[i]%2 == 0){
				fmt.Println(sli[i])
			}
		} 
	}(b...)
}

func returnTriple (xi ...int) []int {
	// This function will double the numbers and return a slice
	var returnSlice []int

	for _,v := range xi {
		a := v * 3
		returnSlice= append(returnSlice, a)
	}

	return returnSlice
}