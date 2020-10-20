//Using Composite Literal:
// - create an ARRAY which holds 5 Values of type Int
// - Assign Values to each index position
// - Range over the array and print the values out
// - Using Format Printing, Print the type of the array out

package main

import "fmt"

func main() {
	fmt.Println("Ninja - Level - 4, Exercise : 1")
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	for k, v := range array {
		fmt.Println(k, v)
	}
	fmt.Printf("%T\n", array)
}
