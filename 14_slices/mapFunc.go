package main

import "fmt"

func main() {
	fmt.Println("--------------create map- and--ok------------------------------------------")
	map1()
	fmt.Println("--------------add values to map and loop------------------------------------------")
	map2()
	fmt.Println("--------------mapTest------------------------------------------")
	mapTest()
	fmt.Println("--------------mapTest------------------------------------------")
	mapDelete()
}

// basically python dictionaries----type?, unordered
// set the type of the key and the value
func map1() {
	m := map[string]int{
		"age":    32,
		"height": 63,
	}
	fmt.Println(m)
	fmt.Println(m["age"])

	// if you ask for a key that doesnt exist you will get zero type back---problematic so you can do this:
	v, ok := m["weight"]
	fmt.Println(v, ok)

	// combine it with conditional logic to see if the key exists VERY COMMON:
	if v, ok := m["weight"]; ok {
		fmt.Println("This is the IF print if value exists:", v)
	}

	if v, ok := m["height"]; ok {
		fmt.Println("This is the IF print if value exists:", v)
	}
}

func map2() {
	m := map[string]int{
		"craig": 39,
		"court": 38,
		"bodhi": 1,
	}

	fmt.Println(m["craig"])

	m["henrik"] = 12
	m["analei"] = 9
	m["fredrik"] = 7

	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}


func mapTest(){
	m := map[string]int{
		"age": 37,
		"weight": 150,
		"height": 32,
	}

	fmt.Println(m)
	m["dead"] = 1972
	for i, v := range m {
		fmt.Println(i, v)
	}
}

func mapDelete(){

}