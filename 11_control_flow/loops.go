package main

import "fmt"

func main() {
	fmt.Println("---------------while like loop(for with condition)-----------------------------------")
	substituteWhile()
	fmt.Println("---------------for loop--------------------------------------------------------------")
	traditionalFor()
	fmt.Println("---------------nested loop-----------------------------------------------------------")
	nestLoop()
	fmt.Println("---------------break loop-----------------------------------------------------------")
	breakLoop()
	fmt.Println("---------------continue loop-----------------------------------------------------------")
	continueLoop()
	fmt.Println("---------------ascii loop-----------------------------------------------------------")
	printAscii()
}

func substituteWhile() {
	x := 5
	for x > 0 {
		fmt.Println(x)
		x--
	}
}

func traditionalFor() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func nestLoop() {
	for i := 0; i < 5; i++ {
		fmt.Printf("The outer  loop: %d\n", i)
		for k := 10; k < 15; k++ {
			fmt.Printf("\tThe inner loop is %d\n", k)
		}
	}
}

func breakLoop() {
	x := 1
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
}

func continueLoop() {
	x := 1
	for {
		if x > 10 {
			break
		} else if x%2 == 0 {
			x++
			continue
		}
		fmt.Println(x)
		x++
	}
}

func printAscii() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
