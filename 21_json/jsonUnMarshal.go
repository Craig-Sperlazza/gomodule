package main

import (
	"encoding/json"
	"fmt"
)

// marshal() turns our struct into json----returns a slice of bytes---that is why you use os.Stdout.Write() because it expects a slice of bytes
// or why you convert it to a string first fmt.Println(string(sliceOfBytes))
//
// unmarshal() turns our json (a slice of bytes) into a go struct----note though that it is stored in a pointer, so it must be an address you pass into unmarshal()

// YOU HAVE TO MAKE THE FIELDS and TYPE UPPERCASE TO EXPORT IT OUT OF THE PROGRAM WITH MARSHAL

// https://mholt.github.io/json-to-go/
// must create your data structure to unmarshal your json string into
// you can use the website above to get that structure!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// so you take your json data and then convert it to a slice of bytes
	stringJSONData := `[{"First":"James","Last":"Bond","Age":35},{"First":"Miss","Last":"Moneypenny","Age":30}]`
	byteSlice := []byte(stringJSONData)

	// people := []person{}
	var people []person

	// remember must pass in the pointer adddress
	err := json.Unmarshal(byteSlice, &people)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
