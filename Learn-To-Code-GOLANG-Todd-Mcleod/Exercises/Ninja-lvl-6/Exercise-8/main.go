//Create a function which returns a function
//Assign the return function to a variable
//Call the returned function

package main

import "fmt"

func main() {
	fmt.Println("Ninja - Lvl - 6, Exercise : 8")
	f := firstLayer("hello")
	x := f()
	fmt.Println(x)

}

func firstLayer(value string) func() string {
	fmt.Println("The value passed through the first layer : ", value)
	return func() string {
		fmt.Println(value)
		return "The value passed through the second layer."
	}
}
