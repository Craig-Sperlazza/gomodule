package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}


type byAge []user
func (s byAge) Len() int {
    return len(s)
}
func (s byAge) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byAge) Less(i, j int) bool {
    return s[i].Age < s[j].Age
}

// sorts by last name
type byLast []user
func (s byLast) Len() int {
    return len(s)
}
func (s byLast) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLast) Less(i, j int) bool {
    return s[i].Last < s[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	fmt.Println("NOT SORTED___________________________________________________-")
	users := []user{u1, u2, u3}

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings{
			fmt.Println("\t", v)
		}
	}

	fmt.Println("SORTED BY AGE___________________________________________________-")
	sort.Sort(byAge(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings{
			fmt.Println("\t", v)
		}
	}

	fmt.Println("SORTED BY LAST___________________________________________________-")
	sort.Sort (byLast(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings{
			fmt.Println("\t", v)
		}
	}

	fmt.Println("SORTED BY LAST AND SAYINGS SORTED ALSO___________________________________________________-")
	sort.Sort (byLast(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings{
			fmt.Println("\t", v)
		}
	}


}
