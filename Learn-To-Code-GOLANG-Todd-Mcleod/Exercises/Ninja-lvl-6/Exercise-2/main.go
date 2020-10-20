//Create a function with a identifier foo that
// - takes in a variadic parameter of type int
// - pass in a value of type []int into your function (unfurl the []int)
// - return the sums of all values of type int passed in
//Create a function with identifier bar that
// - takes in a paramenter of [] int
// - return the sums of all values of type int passed in
package main

import "fmt"

func main() {
	fmt.Println("Ninja - Lvl - 6 , Exercise : 2")
	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(value ...int) int {
	var sum int
	for _, v := range value {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
