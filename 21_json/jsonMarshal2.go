package main

import (
	"encoding/json"
	"fmt"
)

// YOU HAVE TO MAKE THE FIELDS and TYPE UPPERCASE TO EXPORT IT OUT OF THE PROGRAM WITH MARSHAL

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   35,
	}

	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   30,
	}

	people := []Person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// convert the byte slice to a string
	fmt.Println(string(bs))

}
