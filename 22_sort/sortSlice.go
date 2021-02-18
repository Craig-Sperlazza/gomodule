package main

// https://golang.org/pkg/sort/#pkg-examples

import (
	"fmt"
	"sort"
)

func main(){
	xi := []int{4, 8, 3, 9, 11, 344, 1, 22}

	xs := []string{"r", "R", "e", "E", "x", "X"}

	a := xs

	sort.Ints(xi)

	fmt.Println(xi)

	sort.Strings(xs)

	// note that this will sort both xs and a---sorts the uunderlying array
	fmt.Println(xs)
	fmt.Println(a)

}
