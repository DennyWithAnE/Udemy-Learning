//Assign a function to a variable, then call that function.package main

package main

import "fmt"

func main() {
	fmt.Println("Ninja - lvl - 6")
	x := foo
	x()
}

func foo() {
	fmt.Println("I am being called by a variable.")
}
