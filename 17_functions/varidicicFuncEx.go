// redo of exercise 2

package main

import "fmt"

var xi = []int{2,4,6,8,10}


func main() {
	returnInt := unfurl(xi...)
	fmt.Println(returnInt)
}

func unfurl (sli ...int) int{
	var total int
	for i, v := range sli {
		fmt.Println("The index: value is:", i, ":", v)
		total += v
	}
	return total
}