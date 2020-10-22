//Using goroutines, create an incrementer program
// - Have a variable to hold the incrementer value
// - launch a bunch of goroutines
// - - Each Goroutines should read the incrementer value
// - - - Store it in a new variable
// - - Yield the processor with runtime.GoSched()
// - - increment the new value
// - - write the value in the new variable back to the incrementer variable
//Use WaitGroup to wait for all of your goroutines to finish
//After completing the above it will create a race condition
//Prove that this is a race condition by using the -race flag
//If you need help, you can have a look at the hint :
// - - - - https://play.golang.org/p/FYGoflKQej

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Ninja - Lvl - 9, Exercise : 3")
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	counter := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Count : ", counter)
}
