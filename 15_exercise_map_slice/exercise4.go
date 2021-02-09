package main

import "fmt"

func main(){
	fmt.Println("------------Exercise 1---Work With Array--------------------------------------------------")
	ex1()
	fmt.Println("------------Exercise 2---Work With Slice--------------------------------------------------")
	ex2()
	fmt.Println("------------Exercise 3---Slice A Slice--------------------------------------------------")
	ex3()
	fmt.Println("------------Exercise 4---Append To A Slice--------------------------------------------------")
	ex4()
	fmt.Println("------------Exercise 5---Delete A Slice--------------------------------------------------")
	ex5()
	fmt.Println("------------Exercise 6---USe make to create a 5 capacity slice for strings----------------------------------")
	ex6()
	fmt.Println("------------Exercise 7---multidimensional slice----------------------------------")
	ex7()
	fmt.Println("------------Exercise 8---map----------------------------------")
	ex8()
	fmt.Println("------------Exercise 9---add to map----------------------------------")
	ex9()
}

// array
func ex1(){
	a := [5]int{}
	fmt.Println(cap(a))
	a[0]=11
	a[1]=21
	a[2]=31
	a[3]=41
	a[4]=51
	
	fmt.Printf("%T\n", a)

	for i, v := range a{
		fmt.Println(i,":",v)
	}
}

// slice
func ex2(){
	xi := []int{11, 12, 13, 14, 15}
	fmt.Printf("%T\n", xi)

	for i, v := range xi {
		fmt.Println(i, v)
	}
}

// slice a slice
func ex3(){
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi[:5]) 
	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])
}

// append to a slice
func ex4(){
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)

	xi = append(xi, 52)
	fmt.Println(xi)

	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)

	xi2 := []int{56, 57, 58, 59, 60}

	xi = append(xi, xi2...)
	fmt.Println(xi)
}

func ex5(){
	// [42 43 44 48 49 50 51]---append and slice to arrive at this slice
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)

	xi = append(xi[:3], xi[6:]...)
	fmt.Println(xi)
}

// use make to create a capcity of 5 slice (strings)
func ex6(){
	// make([]type, size, capacity)
	xs := make([]string, 5, 5)
	fmt.Println(xs)
	xs = []string{"red", "green", "blue", "orange", "purple"}
	// xs = append(xs[:0], "red", "green", "blue", "orange", "purple")
	fmt.Println(xs)

	for i:=0; i<len(xs); i++ {
		fmt.Println(i, xs[i])
	}
}

// multidimensional slice
func ex7(){
	xs := []string{"Craig", "Music", "Toyota"}
	xs2 := []string{"Bodhi", "Balls", "Ford"}

	multiDimensionalSlice := [][]string{xs, xs2}

	fmt.Println(multiDimensionalSlice)

	// ranges over each record as a whole
	for i, v := range multiDimensionalSlice {
		fmt.Println(i,v)
	}

	// ranges over each value in a record
	for i, v := range multiDimensionalSlice {
		fmt.Println("record: ", i)
		for ind, val := range v{
			fmt.Printf("\t index position: %v \t value:  %v \n", ind, val)
		}
		
	}
}

func ex8(){
	ms := map[string][]string{
		"bond_james": 		[]string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": 	[]string{"James Bond", "Literature", "Computer Science"},
		"dr_no": 			[]string{"Being evil", "Ice cream", "Sunsets"},
	}
		
	// fmt.Println(ms)

	for k, v := range ms{
		fmt.Println(k, "favorite things are:")
		for k2, val := range v{
			fmt.Printf("\t%v:%v\n",k2, val)
		}
	}
}

func ex9(){
	ms := map[string][]string {
		"bond_james": 		[]string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": 	[]string{"James Bond", "Literature", "Computer Science"},
		"dr_no": 			[]string{"Being evil", "Ice cream", "Sunsets"},
	}

	ms["craig"]= []string{"red", "blue", "green"}
	ms["henrik"]=[]string{"teal", "orange", "black"}

	for i, v := range ms {
		fmt.Printf("%v favorite things are:\n", i)
		for i2, v2 := range v{
			fmt.Printf("\t%v:%v\n", i2, v2)
		}
	}
}