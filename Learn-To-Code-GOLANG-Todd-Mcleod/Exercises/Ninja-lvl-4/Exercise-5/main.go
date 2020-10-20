// To delete from a slice, we will be using APPEND ALONG WITH SLICING.package main
// Start off with this slice
// x := []int{42,43,44,45,46,47,48,49,50,51}
// Use append and slicing to get this value to print out
// [42,43,44,48,49,50,51]

package main

import "fmt"

func main() {
	fmt.Println("Ninja-lvl-4, Exercise : 5")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
