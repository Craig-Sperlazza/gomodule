package main

import "fmt"

var y bool

// can specify types for numbers---float 64, int(signed) are default----but can do float32, int8 (signed), uint32(unsigned), etc.
var a float32

// aliases
// byte is an alias for uint8
// rune is an alias for int32
var d rune

// note you can decalre const and variables like this:

const (
	re = 2
	rd float64 = 32.1
	rw string = "Her"
)

// iota----add values in order---0 onwards...resets when you use const again for something else
const (
	h = iota
	i = iota
	j = iota
	k 
	l
)

func main() {
	fmt.Println("---------BOOL----------------------------------------------------")
	boole()
	fmt.Println("---------NUMBERS-------------------------------------------------")
	num_types()
	fmt.Println("---------STRINGS-------------------------------------------------")
	str_types()
	fmt.Println("---------CONSTANTS-------------------------------------------------")
	const_types()
	fmt.Println("---------IOTA-------------------------------------------------")
	iota_types()
}

func boole() {
	x := true
	fmt.Println(x, y)

	a := 7
	b := 42

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a <= b)
}

func num_types() {
	// auto goes to float64
	b := 1.3
	c := 4
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	// rune = int32
	fmt.Printf("%T\n", d)
}

func str_types() {
	s := "Hello, Playground"
	fmt.Println(s)

	// byte slice---going to pull out the utf-8 values of the letters
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	n := bs[0]

	fmt.Println(n)
	// type, then binary, then hex
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)



	// utf code point
	for i:= 0; i < len(s); i++{
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	// index and value in utf-8
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func const_types (){
	fmt.Println(re)
	fmt.Printf("%T\n", rw)
	fmt.Printf("%T\n", rd)
}

func iota_types (){
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Printf("%T\n", h)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", k)
	fmt.Printf("%T\n", l)
}