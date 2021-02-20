// use mutex to prevent the race condition(multiple go routines accessing a shared variable)
// mutex essentially checks it out like a library book and then tells other routines ok, you can now check it out
// go run -race name.go will let you know if you have a race condition

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	
	var mu sync.Mutex

	for i := 0; i < gs; i++{
		go func(){
			mu.Lock()
			v := counter
			// could go to sleep or can run Gosched() which basically tells the cpu you are releasing it to run something else
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	
	fmt.Println("count:", counter)
}