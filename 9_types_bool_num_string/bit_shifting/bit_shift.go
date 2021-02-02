package main

import "fmt"

const (
	// 0 is the firs tvalue so throw it away, then 1, 2, 3
	_ = iota
	kb1 = 1 << (iota * 10)
	mb1 = 1 << (iota * 10)
	gb1 = 1 << (iota * 10)
)

func main(){
	shift()
	byteShift()
	byteShiftIota ()
}

func shift()  {
	x := 4
	fmt.Printf("%d\t%b\n", x, x)

	// shift bit over 1 in binary (100 is 4, so 1000 is 8)
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y)
}

func byteShift(){
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

func byteShiftIota (){
	// note that kb to mb to gb is a shift of 10 spots to the left (10 zeros)
	// kb := 1024
	// mb := kb << 10
	// gb := mb << 10

	// or you can use iota from above

	fmt.Printf("%d\t\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t\t%b\n", mb1, mb1)
	fmt.Printf("%d\t\t%b\n", gb1, gb1)
}