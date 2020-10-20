//Start with this slice
// - x := []int{42,43,44,45,46,47,48,49,50,51}
//Append to that slice this value
// - 52
// - Print out the slice
//In one statement append to that slice these values
// - 53
// - 54
// - 55
// - Print out the slice
//Append to the slice this slice
// - y := []int{56,57,58,59,60}
// - Print out the slice

package main

import "fmt"

func main() {
	fmt.Println("Ninja-Lvl-4, Exercise : 4")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println(cap(x))
	fmt.Println(len(x))

}
