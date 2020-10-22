//Fix the race condition you created in the previous exercise by using a mutex

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex

func main() {
	fmt.Println("Ninja - Lvl - 9, Exercise : 4")
	counter := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Count : ", counter)

}
