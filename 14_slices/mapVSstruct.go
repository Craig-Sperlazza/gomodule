// https://medium.com/wesionary-team/structs-and-maps-in-golang-15c7ac08db31

// Maps are similar to structs in the way that, they too hold collection of values as key-value pairs.

// But there are fair number of differences.
// Firstly, all the keys and all the values inside a Map are defined when it is initialized and stay the same for all the key-value pairs.
// 		For example, if we initialize a map with a key of type string and value of type int, all the keys need to be of type string and values of type int.
// Secondly, we can iterate over Maps with a loop, we cannot with Structs.
// Thirdly, Maps are Pointer types while structs are Value types. We will discuss more about these in the next article.

package main

import (
	"fmt"
)

type student struct {
	first string
	grade int
}

type mapInStruc struct {
	first string
	testScores map[string]int
}


func main(){
	fmt.Println("-----------------strucs----------------------------------------------------")
	strucStudent := student {
		first: "Craig",
		grade: 11,
	}

	fmt.Println("struct:", strucStudent.first)
	fmt.Println("struct:", strucStudent.grade)
	fmt.Println("struct:", strucStudent)
	strucStudent.printStruc()

	fmt.Println("-----------------maps----------------------------------------------------")
	mapStudent := map[string]string {
		"name": "Joe",
		"grade": "11",
	}

	printMap(mapStudent)

	fmt.Println("-----------------map in a struc----------------------------------------------------")

	s2 := mapInStruc {
		first: "Tom",
		testScores: map[string]int{
			"test 1": 50,
			"test 2": 75,
		},
	}

	for i, v := range s2.testScores{
		fmt.Println(i, v)
	}

	fmt.Println(s2.testScores)
}

func printMap(c map[string]string){
	for i, v := range c {
		fmt.Println("map:", i, v)
	}
}

// create a receiver function for student
func (s student) printStruc(){
	fmt.Printf("%v\n", s)
}