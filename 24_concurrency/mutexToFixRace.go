// use mutex to prevent the race condition

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
	

	for i := 0; i < gs; i++{
		go func(){
			v := counter
			// could go to sleep or can run Gosched() which basically tells the cpu you are releasing it to run something else
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	
	fmt.Println("count:", counter)
}