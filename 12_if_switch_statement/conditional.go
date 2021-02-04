package main

import (
	"fmt"
)

func main(){
	true_false()
	initializeStatement()
	ifElse()
	loopCondit()
	switchStatement()
	switchFallthrough()
	switchDefaultCase()
}

func true_false(){
	if true {
		fmt.Println("true")
	}

	if !false {
		fmt.Println("still true")
	}

	if (2==2) {
		fmt.Println("also true")
	}
}

// limits scope of the variable to the if
func initializeStatement(){
	if x:= 42; x == 42{
		fmt.Println("x is true")
	}
}

func ifElse() {
	x := 43
	if x <= 20{
		fmt.Println("x is less than 20")
	} else if x >= 21 && x <=42 {
		fmt.Println("x between 21 and 42")
	} else{
		fmt.Println("x is greater than 42")
	}
}

func loopCondit() {
	for i := 0; i <10; i++ {
		if i % 2 !=0{
			fmt.Printf("%d is an odd number\n", i)
		} else {
			fmt.Printf("%d is an even number\n", i)
		}
	}
}

func switchStatement(){
	x := 44
	switch {
	case x % 2 != 0:
		fmt.Printf("%d is odd\n", x)
	case x % 2 == 0:
		fmt.Printf("%d is even\n", x)
	}
}

// if that condition is true, it will fallthrough and run the next regardless of whether it is true
// shouldnt really use it
func switchFallthrough(){
	x := 43
	switch {
	case x % 2 != 0:
		fmt.Printf("%d is odd\n", x)
		fallthrough
	case x % 2 == 0:
		fmt.Printf("%d is even\n", x)
	}
}


// default case but also notice you can hav emore than one case in each case
func switchDefaultCase(){
	x := 48
	switch {
	case x == 42, x == 46, x == 48:
		fmt.Println("X is 42 or 46 or 48")
	case x == 44:
		fmt.Println("x is 44")
	default:
		fmt.Println("x is something else")
	}
}