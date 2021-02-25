package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	fmt.Println("Main Routine 1")
	wg.Add(1)
	go func1()
	wg.Wait()

	
	fmt.Println("Main Routine 2")

	wg.Add(1)
	go func2()

	wg.Wait()
	fmt.Println("Main Routine 3")

	
}

func func1(){
	fmt.Println("Go Routine 1")
	wg.Done()
}

func func2(){
	fmt.Println("Go Routine 2")
	wg.Done()
}