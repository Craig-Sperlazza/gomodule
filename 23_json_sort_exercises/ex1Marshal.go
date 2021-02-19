package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age int
}


func main() {
	fmt.Println(2)
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{
		u1, 
		u2,
		u3,
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(jsonData))

	
}