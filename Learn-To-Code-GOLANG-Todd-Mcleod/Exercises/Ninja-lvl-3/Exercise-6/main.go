package main

import "fmt"

func main() {
	fmt.Println("Ninja Lvl : 3 , Exercise Level : 6")
	fmt.Println("Show the example of a IF statement")
	var input int
	fmt.Println("Please input a number : ")
	fmt.Scanln(&input)
	if input == 50 {
		fmt.Println("Bingo")
	} else {
		fmt.Println("Please try again")
	}
}
