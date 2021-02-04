package main

import "fmt"

func main() {
	arr1()
	slice1()
	rangeSlice()
	rangeSliceInt()
	sliceOfSlice()
	appendSlice()
	deleteSlice()
}

// note we arent supposed to use arrays in go---use slices
func arr1() {
	var x [5]int
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x[2] = 1
	fmt.Println(x)
	fmt.Println(len(x))
}

// SLICES----ALLOWS YOU TO GROUP VALUES OF THE SAME TYPE
// use composite literal syntax
// x:= type{values}

// NOTE: use xi or ii for slices is idomatic go (xi = slice of ints, xs = slice of strings, etc.)
func slice1() {

	// use composite literal to make a slice of int
	xi := []int{4, 5, 6, 7, 8, 76}
	fmt.Println(xi)
}

func rangeSlice() {
	xi := [3]string{"red", "green", "blue"}
	fmt.Println(xi[1])
	for i, v := range xi {
		fmt.Println(i, v)
	}

}

func rangeSliceInt() {
	xi := []int{2, 4, 6, 8}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}

func sliceOfSlice() {
	xi := []int{4, 5, 6, 7, 89, 10}
	fmt.Println(xi[2:5])

	for i := 0; i < len(xi); i++ {
		fmt.Println(i, xi[i])
	}
}

func appendSlice() {
	xi := []int{2, 4, 6, 8, 10}
	xi = append(xi, 12, 14, 16)
	fmt.Println(xi)

	// can unfurl since the second parameter is a variadic parameter in append
	y := []int{2, 3, 5, 7, 11}
	z := []int{13, 17, 19, 23}
	y = append(y, z...)
	fmt.Println(y)
}

func deleteSlice() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// say we want to delete 4, 5, 8
	// dont forget the dots at end to unfurl the data---has to be ints or strings or whatever, not a slice (whereas the first parameter is slice)
	a = append(a[0:3], a[5:7]...)
	fmt.Println(a)
}
