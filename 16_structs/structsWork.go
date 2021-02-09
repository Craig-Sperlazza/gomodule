package main

import "fmt"

// structs---can do types on same line if prefer
type person struct {
	first string
	last  string
	age, weight   int
}

// embedded structs
type secretAgent struct {
	person
	licToKill bool
}

func main() {
	createPeople()
	anonStruc()
}

func createPeople() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   40,
		weight: 100,
	}

	p2 := person{
		first: "craig",
		last:  "sperlazza",
		age:   39,
		weight: 100,
	}

	p3 := secretAgent{
		person: person{
			first: "court",
			last:  "sperl",
			age:   38,
			weight: 100,},
		licToKill: true,
	}

	fmt.Println(p1)
	fmt.Println(p2.first, p2.age)
	fmt.Println(p3)
	fmt.Println(p3.age, p3.licToKill)
}

// anonymous structure
func anonStruc(){
	p1 := struct {
		first string
		last  string
		age, weight   int
	}{
		first: "courtNEEEE",
		last:  "sperl",
		age:   38,
		weight: 100,
	}

	fmt.Println(p1)
}