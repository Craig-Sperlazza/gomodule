package main

import "fmt"

type person struct {
	first string
	last string
	ice []string
}

type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle 
	fourWheel bool
}

type sedan struct{
	vehicle 
	luxury bool
}

func main () {
	ex1()
	fmt.Println("-------------ex 2---------------------------")
	ex2()
	fmt.Println("-------------ex 3---------------------------")
	ex3()
	fmt.Println("-------------ex 4---------------------------")
	ex4()
	fmt.Println("-------------ex 5---------------------------")
	ex5()
}

func ex1 (){
	p1 := person {
		first: "craig",
		last: "sperl",
		ice: []string{"v","c","s"},
	}

	p2 := person {
		first: "Court",
		last: "Sperl",
		ice: []string{"a", "b", "c"},
	}

	for i, v := range p1.ice {
		fmt.Println(i, v)
	}

	for i2, v2 := range p2.ice {
		fmt.Println(i2, v2)
	}

	fmt.Println(p1)
}

func ex2(){
	peopleMap := map[string]person{
		"sperlazza":{
			first: "craig",
			last: "sperl",
			ice: []string{"v","c","s"},
		},
		"smith":{
			first: "bob",
			last: "earl",
			ice: []string{"a","b","c"},
		},
	}

	fmt.Println(peopleMap)

	for _, v2 := range peopleMap{
		fmt.Println(v2.first)
		fmt.Println(v2.last)
		for icek, iceval := range v2.ice {
			fmt.Println(icek, iceval)
		}
	}

	fmt.Println(peopleMap["sperlazza"].ice)

	for i,v := range peopleMap{
		fmt.Println(i, v.first, v.last, v.ice)
		fmt.Println(i, v.first, v.last, v.ice)
	}
}

func ex3(){
	truck1 := truck {
		vehicle: vehicle {
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	car1 := sedan {
		vehicle: vehicle {
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(truck1)
	fmt.Println(truck1.vehicle.doors)
	fmt.Println(car1)
	fmt.Println(car1.luxury)
	fmt.Println(car1.vehicle.color)
}

func ex4(){
	person := struct{
		name string
		favs []string
	}{
		name: "Craig",
		favs: []string{"a", "b", "c"},
	}
	fmt.Println(person)
}

func ex5(){
	car := struct{
		color string
		wheels int
		features []int
		make map[string]string
	}{
		color: "blue",
		wheels: 4,
		features: []int{1,2,3,4},
		make: map[string]string{
			"honda": "accord",
		},
	}

	fmt.Println(car.color)
	for _, v := range car.features{
		fmt.Println(v)
	}

	for k, v := range car.make{
		fmt.Println(k, ":", v)
	}
}