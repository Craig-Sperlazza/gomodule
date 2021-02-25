// use atomic to prevent the race condition(multiple go routines accessing a shared variable)
// atomicessentially checks it out like a library book and then tells other routines ok, you can now check it out
// "go run -race name.go" will let you know if you have a race condition

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// must be  an int64 for atomic
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// with atomic you give it the address of the variable and what you want it incremented by
			// (so this writes and then reads to counter without race conditions)
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
