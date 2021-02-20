package main

// https://pkg.go.dev/runtime#pkg-functions

import (
	"fmt"
	"sync"
)

// var wg is of type WaitGroup from package sync
var wg sync.WaitGroup

func main(){
	fmt.Println("Main Function11111")

	wg.Add(1)
	go goRoutinefunc1()

	fmt.Println("Main Function22222")

	wg.Add(1)
	go goRoutinefunc2()

	fmt.Println("Main Function3333")
	regFunc()


	wg.Wait()
}

func goRoutinefunc1(){
	fmt.Println("Go Routine First Func")
	wg.Done()
}

func goRoutinefunc2(){
	fmt.Println("Go Routine Second Func")
	wg.Done()
}

func regFunc(){
	fmt.Println("Regular Sequential func")
}