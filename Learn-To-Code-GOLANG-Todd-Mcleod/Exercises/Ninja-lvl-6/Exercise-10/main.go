//Closure is when we have enclosed the scope of the variable in some code block
//Create a function which enclose the scope of a variable

package main

import "fmt"

func main() {
	fmt.Println("Ninja - Lvl - 6, Exercise : 10")
	x := foo()
	fmt.Println(x)
	test1 := incremental()
	fmt.Println(test1())
	fmt.Println(test1())
	fmt.Println(test1())
	fmt.Println(test1())
	fmt.Println(test1())
	test2 := incremental()
	fmt.Println(test2())
	fmt.Println(test2())
	fmt.Println(test2())
	fmt.Println(test2())
	fmt.Println(test2())

}

func foo() bool {
	state := 10
	if state == 1 {
		return true
	}
	return false

}

func incremental() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
