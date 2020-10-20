//Create a function with the identifier foo that returns an integer
//Create a function with the identifier bar that returns an int and a string
//Call Both functions
//Print out their results
package main

import "fmt"

func main() {
	fmt.Println("Ninja - Lvl - 6 , Exercise : 1")
	fmt.Println(foo(10))
	fmt.Println(bar(10))
}

func foo(input int) int {
	var results int
	results = results + input
	return results
}

func bar(input int) (int, string) {
	var results int
	results = results + input
	return results, "Hello"
}
