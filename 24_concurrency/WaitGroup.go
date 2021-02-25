package main

// https://pkg.go.dev/runtime#pkg-functions

import (
	"fmt"
	"runtime"
	"sync"
)

// sync package lets you manage go routines (my poor definition)
// example of waitgroup
// https://golang.org/pkg/sync/

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())

	// go foo()launches it as a go routine---creates a concurrent design pattern
	// main program is running as a goroutine and then this launches a second one
	// if you have more than 1 cpu, it will launch them in parallel
	// must use sync package (waitgroup here) to essentially wait for other routines before exiting main()
	// wg.Add(1) adds 1 waitgroup, the we use wg.Wait() to wait for wg.Done() to finish (which is in the function we designated as a go routine)
	// wg.add(int) adds something to wait for, wg.Wait() waits for it, wg.Done() removes it
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
