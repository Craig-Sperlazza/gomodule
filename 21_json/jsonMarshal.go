package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// marshal() turns our struct into json----returns a slice of bytes---that is wht you use os.Stdout.Write because it expects a slice of bytes
// unmarshal() turns our json (a slice of bytes) into a go struct----note though that it is stored in a pointer, so it must be an address you pass into unmarshal()

// YOU HAVE TO MAKE THE FIELDS and TYPE UPPERCASE TO EXPORT IT OUT OF THE PROGRAM WITH MARSHAL

// https://mholt.github.io/json-to-go/

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	p1 := Person{
		Name:    "Craig",
		Age:     39,
		Hobbies: []string{"football", "baseball", "archery"},
	}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(jsonData)
	fmt.Println(string(jsonData))

}
