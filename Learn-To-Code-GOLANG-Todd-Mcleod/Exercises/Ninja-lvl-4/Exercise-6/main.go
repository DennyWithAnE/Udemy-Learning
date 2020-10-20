//Create a slice to store the names of the states in United States of America
//What is the length of your slice?
//What is the capacity of your slice?
//Print out all their values along with their index position in the slice
//without using the range clause.

package main

import "fmt"

func main() {
	fmt.Println("Ninja-lvl-4, Exercise : 6")
	x := make([]string, 5, 5)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = []string{"Alaska", "Arizona", "Florida", "Michigan", "Alabama"}
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
