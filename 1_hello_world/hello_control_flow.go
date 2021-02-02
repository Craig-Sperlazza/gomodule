package main

import "fmt"

func main() {
	fmt.Println("Hello, Playground")
	foo()
	fmt.Println("Sequence continued")
	countTenEven()
}

func foo(){
	fmt.Println("I am in foo")
}

func countTenEven(){
	for i := 0; i <11; i++{
		if i % 2 == 0 {
			fmt.Println(i)
		}		
	}
}