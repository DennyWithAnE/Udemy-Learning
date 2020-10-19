package main

import "fmt"

func main() {
	for {
		fmt.Println("Ninja Lvl : 3 , Exercise Level : 7")
		fmt.Println("Show the example of a IF/Else-IF/ELSE statement")
		var input int
		fmt.Println("Please input a number : ")
		fmt.Scanln(&input)
		if input == 50 {
			fmt.Println("Bingo")
			break
		} else if input > 50 {
			fmt.Println("Too big")

		} else {
			fmt.Println("Too low")
		}
	}
}
