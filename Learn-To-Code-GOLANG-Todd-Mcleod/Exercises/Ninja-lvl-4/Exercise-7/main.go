//Create a slice of a slice of string ( [][] string ).
//Store the following data in the multi-dimensional slice
// - "James", "Bond", "Shaken, not stirred"
// - "Miss", "MoneyPenny", "Hellooooooo, James"
//Range over the records, then range over the data in each records

package main

import "fmt"

func main() {
	fmt.Println("Ninja-lvl-4, Exercise : 7")
	x := [][]string{}
	fmt.Println(x)
	firstSlice := []string{"James", "Bond", "Shaken, not stirred"}
	x = append(x, firstSlice)
	fmt.Println(x)
	secondSlice := []string{"Miss", "MoneyPenny", "Hellooooo James!"}
	x = append(x, secondSlice)
	fmt.Println(x)

	for i, slice := range x {
		fmt.Println("record : ", i)
		for j, val := range slice {
			fmt.Printf("\t index position: %v \t value : \t %v \n", j, val)
		}
	}
}
