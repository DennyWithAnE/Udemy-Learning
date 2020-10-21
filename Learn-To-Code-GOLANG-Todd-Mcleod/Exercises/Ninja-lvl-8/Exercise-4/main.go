//Starting with this code, sort the []int and []string for each person.
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Ninja-Lvl-8, Exercise : 4")

	xi := []int{1, 3, 4, 99, 54, 67, 43, 21, 22, 77, 88, 0, 2, 5}
	xs := []string{"Denny", "Jasmine", "Jessica", "Yk", "Nicholas", "Anna", "Ernest", "Brady"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
