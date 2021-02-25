// We are passing around an incrementer value to 100 sub go routines
// ....it creates a race condition and yeilds very odd results
// HERE WE FIX THE RACE CONDITION WITH sync.Mutex

package main

import (
	"fmt"
	// "runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// var mu sync.Mutex

const num = 100

var counter int64


func main(){
	wg.Add(num)
	for i:=0; i < num; i++{
		go func (){
			atomic.AddInt64(&counter, 1)
			// runtime.Gosched()
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final", counter)
}