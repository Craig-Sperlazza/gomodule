package main

import (
	"fmt"
)

const (
	h int = 42
	i = 43
)


func main(){
	// ex1()
	// ex2()
	// ex3()
	ex4()
	ex5()
	ex6()
}

func ex1 (){
	// print decimal, binary, hex
	a := 4
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
}

func ex2 (){
	b := (3 == 3)
	c := (3 <= 3)
	d := (3 >= 3)
	e := (3 != 3)
	f := (3 < 3)
	g := (3 > 3)
	
	fmt.Println(b, c, d, e, f, g)
}

func ex3 (){
	// make a typed and untyped const
	fmt.Println("typed: ", h)
	fmt.Printf("%T\n", h)
	fmt.Println("untyped: ", i)
	fmt.Printf("%T\n", i)

}

func ex4 (){
	// make a variable---print d, b, h and then shift bit 1 postions left and reprint
	j := 2
	fmt.Printf("%d\t\t%b\t\t%#x\n", j, j, j)
	k := 2 << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", k, k, k)
}

func ex5(){
	raw_s := `Just in 
	case you "did" 
	miss it`
	fmt.Println(raw_s)
}

func ex6(){
	// use iota to print constants equal to the last 4 years
	const (
		_ = iota
		l = 2017 + iota
		m = 2017 + iota
		n = 2017 + iota
	)
	fmt.Println(l, m, n)
	
}