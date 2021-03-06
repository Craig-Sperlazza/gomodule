// We are passing around an incrementer value to 100 sub go routines
// ....it creates a race condition and yeilds very odd results

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const num = 100

var counter = 0


func main(){
	wg.Add(num)
	for i:=0; i < num; i++{
		go func (){
			val := counter
			runtime.Gosched()
			val ++
			counter = val
			fmt.Println("here", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("there", counter)
}