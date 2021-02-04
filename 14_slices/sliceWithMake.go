package main

import "fmt"

func main() {
	makeASlice()
	multiDimensionSlice()
}

// largely you should use composite literals to make a slice but you can use make
// make(type, length, capacity)
// IF YOU GO BEYOND CAPACITY----it throws away old array and doubles the size of the underlying array so it can hold the larger slice
func makeASlice() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 13
	x[9] = 22
	x = append(x, 1234)
	fmt.Println(len(x))
}

func multiDimensionSlice() {
	jb := []string{"James", "Bond", "chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "strawberry", "Hazelnut"}
	fmt.Println(mp)

	// multidimension---a slice of a slice of string type
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
