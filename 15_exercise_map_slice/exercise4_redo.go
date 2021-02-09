package main

import "fmt"

func main(){
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
}

func ex1(){
	arr := [5]int{2, 3, 4, 5, 6}

	fmt.Printf("%T\n", arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func ex2(){
	xi := []int{2, 4, 6, 8, 10, 12}

	fmt.Printf("%T\n", xi)

	for i, v := range xi {
		fmt.Println(i, v)
	}

}

func ex3(){
	xi:= []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(xi[0:5])
	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])
}

func ex4(){
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xi = append(xi, 52)
	fmt.Println(xi)

	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)

	y := []int{22, 23, 24, 25}
	xi = append(xi, y...)
	fmt.Println(xi)
}

// append and slice to get [42, 43, 44, 48, 49, 50, 51]
func ex5(){
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y:= append(xi[:3], xi[6:]...)

	fmt.Println(y)
}

func ex6() {
	xs := make([]string, 5)
	fmt.Println(cap(xs))
	fmt.Println(len(xs))
	xs[0]="aaaa"
	xs[1]="bbbb"
	xs[2]="cccc"
	xs[3]="dddd"
	xs[4]="eeee"

	for i := 0; i < len(xs); i++{
		fmt.Println(i, xs[i])
	}
}

func ex7(){
	jb := []string{"a", "b", "c"}
	mp := []string{"e", "f", "g"}

	xxs := [][]string{jb, mp}

	fmt.Println(xxs)

	for i, v := range xxs {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}

func ex8(){
	ms := map[string][]string{
		"names": []string{"craig", "henrik", "analei"},
		"ages": []string{"ten", "six", "two"},
		"colors": []string{"red", "blue", "pink"},
	}

	fmt.Println(ms)
}

func ex9(){
	ms := map[string][]string{
		"color": []string{"red", "blue"},
		"num": []string{"1", "2"},
		"day": []string{"sun", "mon"},
	}

	ms["year"] = []string{"1982", "1983"}

	for i, v := range ms {
		fmt.Println(i)
		for i2, v2 := range v{
			fmt.Println("  ", i2, ":", v2)
		}
	}
}

func ex10() {
	ms := map[string][]string {
		"day": []string{"mon", "tues"},
		"color": []string{"red", "blue"},
		"num": []string{"1", "2"},
	}

	ms["year"] = []string{"1999", "2000"}

	for i, v := range ms {
		fmt.Println(i)
		for i2, v2 := range v{
			fmt.Println("  ", i2, ":", v2)
		}
	}

	delete(ms, "year")
	fmt.Println(ms)
	
	delete(ms, "day")
	fmt.Println(ms)
}