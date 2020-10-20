//Pass a function as an argument into a function
//CallBack --

package main

import "fmt"

func main() {
	fmt.Println("Ninja - Lvl - 6, Exercise : 9")
	g := bar
	x := foo(g, "hello")
	fmt.Println(x)
}

func foo(f func(input string) string, input2 string) string {
	fmt.Println("current state is foo")
	n := f(input2)
	return n
}

func bar(input string) string {
	fmt.Println("current state is bar")
	return input
}
