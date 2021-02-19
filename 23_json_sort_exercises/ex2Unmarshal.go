package main

import (
	"encoding/json"
	"fmt"
)

// https://mholt.github.io/json-to-go/
type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"Dr","Last":"No","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	// first convert to a byte slice
	bs := []byte(s)

	// we need a slice of persons to hold all the persons
	var people []person

	// remember you must pass it the pointer of your slice
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age, "\n\t", v.Sayings)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	}

}
