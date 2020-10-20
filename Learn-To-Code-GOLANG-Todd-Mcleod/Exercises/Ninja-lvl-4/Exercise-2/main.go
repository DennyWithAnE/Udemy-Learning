// Using a composite literal
// - create a slice of type int
// - assign 10 values
// - range over the slices and print the values out
// - using format printing to print the type of the slice

package main

import "fmt"

func main() {
	fmt.Println("Ninja - lvl - 4, Exercise : 2")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slice)
}
