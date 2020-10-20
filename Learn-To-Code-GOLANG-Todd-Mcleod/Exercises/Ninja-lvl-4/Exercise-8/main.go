// Create a map with a key TYPE string which is a person's last name,
// and a value of TYPE string which stores their favourite things.
// Store 3 records in your map.
// Print out all of the values, along with their position in the slice

package main

import "fmt"

func main() {
	fmt.Println("Ninja-lvl-4, Exercise : 8")
	x := map[string][]string{
		"Denny":   []string{"Soccer", "MLBB", "Food"},
		"Jasmine": []string{"Netflix", "Sleep", "Sushi"},
	}

	fmt.Println(x)

	for k, v := range x {
		fmt.Println("This is :", k)
		for _, v2 := range v {
			fmt.Println("Things they love: ", v2)
		}
	}
}
