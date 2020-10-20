//Build and use an anonymous function

package main

import "fmt"

func main() {
	fmt.Println("Ninja - lvl - 6, Exercise : 6")

	func() {
		fmt.Println("Hello I am the anonymous function!")
	}()
}
