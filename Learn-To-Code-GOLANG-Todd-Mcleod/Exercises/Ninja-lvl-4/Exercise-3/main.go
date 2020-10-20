// Using codes from the previous example, use SLICING to create the following new slice
// which are then printed.
// - [42,43,44,45,46]
// - [47,48,49,50,51]
// - [44,45,46,47,48]
// - [43,44,45,46,47]

package main

import "fmt"

func main() {
	fmt.Println("Ninja - lvl - 4, Exercise : 3")
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(slice[:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])
	fmt.Println(slice[1:6])
}
