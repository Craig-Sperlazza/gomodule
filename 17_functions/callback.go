package main

import "fmt"

// callback---passing a func into a func

var mySlice = []int{1,2,3,4,5,6,7,8,9,10}


func main() {
	s := sum(mySlice...)
	fmt.Println(s, "total")


	s2 := evenNum(sum, mySlice...)
	fmt.Println(s2, "total even")

	s3 := oddNum(sum, mySlice...)
	fmt.Println(s3, "total odd")
}

func sum(xi ...int)int{
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

// lets say we just want to sum the even numbers
// this func evenNum takes in a func (which takes in a slice) and a slice and returns an int

func evenNum(f func(xi ...int) int, slice3 ...int ) int {
	var yi []int
	for _, v := range slice3 {
		if v % 2 == 0{
			yi = append(yi, v)
		}
	}
	return f(yi...)
}


func oddNum (f func(xi ...int)int, slice3 ...int) int {
	var retArr []int

	for _, v := range slice3{
		if v % 2 != 0 {
			retArr = append(retArr, v)
		}
	}
	return f(retArr...)
}






