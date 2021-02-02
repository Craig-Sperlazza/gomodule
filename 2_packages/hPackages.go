package main

import "fmt"

func main(){
	fmt.Println("Heyyo")
	const name = "Craig"
	// n is for the return size in bytes
	// _ is a return throw away for the err in this case----HAVE TO ACCOUNT FOR IT
	// you can not not use variables so you have to throw away or make it a variable name and use it. 
	n, _:= fmt.Println("Hello, playground", name, 42, true)
	// next line prints the bytes
	fmt.Println(n)
}