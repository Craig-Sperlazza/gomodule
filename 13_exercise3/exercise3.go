package main

import "fmt"

func main(){
	// fmt.Println("--------------------------Print 1-100------------------------------------")
	// ex1()
	// fmt.Println("--------------------------All alphabet rune code 3 times------------------------------------")
	// ex2()
	// fmt.Println("--------------------------for loop------------------------------------")
	// ex3()
	// fmt.Println("--------------------------different syntax for loop------------------------------------")
	// ex4()
	// fmt.Println("--------------------------Loop and modulo practice------------------------------------")
	// ex5()
	// fmt.Println("--------------------------if else------------------------------------")
	// ex6and7()
	// fmt.Println("--------------------------switch------------------------------------")
	// ex8()
	fmt.Println("--------------------------switch------------------------------------")
	ex9()
}

func ex1 (){
	i:= 1
	for i < 101 {
		fmt.Println(i)
		i++
	}
}

// All alphabet rune code 3 times
func ex2 (){
	i:= 65
	for i < 127 {
		fmt.Println(i)
		j := 1
		for j < 4{
			fmt.Printf("\t%#U\n", i)
			j++
		}
		
		i++
	}
}

func ex3(){
	x := 1982
	for x < 2022{
		fmt.Println(x)
		x++
	}
}

func ex4(){
	x := 1982
	for {
		fmt.Println(x)
		x++
		if x > 2021{
			break
		}
	}
}

func ex5(){
	x := 10
	for x <= 100{
		if x % 4 == 0{
			fmt.Printf("%v is divisible by 4 \n", x)
		}
		x++
	}
}

func ex6and7(){
	for i:=1; i<100; i++{
		if i % 2 == 0{
			fmt.Printf("%v is divisible by 2\n", i)
		} else if i % 3 == 0 || i % 5 == 0 {
			fmt.Printf("%v is divisible by 3 or 5\n", i)
		} else{
			fmt.Printf("%v is not divisible by 2 or 3 or 5\n", i)
		}
	}
}

// switch statement with no switch statement specified----so default is true
func ex8(){
	switch {
	case false: 
		fmt.Println("not true so it shouldnt print")
	case true:
		fmt.Println("This is true so default switch of true hits here")
	}
}

// switch statement looking for a string with a default
func ex9() {
	favSport := ""
	switch favSport{
	case "football":
		fmt.Println("football is your favorite")
	case "surfing":
		fmt.Println("surfing is your favorite")
	default:
		fmt.Println("You dont like sports")
	}
}