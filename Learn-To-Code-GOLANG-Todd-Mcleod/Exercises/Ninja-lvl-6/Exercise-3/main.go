//Use the "defer" keyword to show that a deferred function runs after the
//function containing it exits

package main

import "fmt"

func main() {
	fmt.Println("Ninja - lvl - 6, Exercise : 3")
	defer show()
	fmt.Println("The defer function is about to get deferred : ")
	fmt.Println("3")
	fmt.Println("2")
	fmt.Println("1")

}

func show() {
	fmt.Println("This line is print after deferring.")
}
