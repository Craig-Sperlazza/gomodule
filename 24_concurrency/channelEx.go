package main

import (
	"fmt"
)

func doSomething(x int) int {
	// does soemthing that is already built
	return x * 5
}

// issue: if a function in a go routine has return values, they are discarded when it completes
// so drop them into a channel using a function literal (aka anaonymous function)
func main() {
	// make a channel
	ch := make(chan int)
	// create a go routine and wrap do something in an anonymous function
	// put the value on the channel and then pull it off
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
