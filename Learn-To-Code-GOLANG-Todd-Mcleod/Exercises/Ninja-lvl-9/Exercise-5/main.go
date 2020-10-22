//Fix a race condition previously by using Atomic this time round.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Ninja - Lvl - 9, Exercise : 5")
	var incrementer int64
	var wg sync.WaitGroup
	gs := 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
}
