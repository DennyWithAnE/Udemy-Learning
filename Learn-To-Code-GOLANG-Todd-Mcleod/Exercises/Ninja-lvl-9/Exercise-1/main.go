//In addition to the main goroutine, launch two additional goroutines
// - each additional goroutines should print something out
//Use waitGroup to make sure each goroutine finishes before you program exits.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Ninja - Lvl - 9, Exercise : 1")
	fmt.Println("The number of cpu : ", runtime.NumCPU())
	fmt.Println("The number of Go Routine : ", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("While running the program The number of cpu : ", runtime.NumCPU())
	fmt.Println("While running the program The number of Go Routine : ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("About to exit the program")
	fmt.Println("Before exiting The number of cpu : ", runtime.NumCPU())
	fmt.Println("Before exiting The number of Go Routine : ", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello I am Foo")
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello I am Bar!")

	}
	wg.Done()
}
