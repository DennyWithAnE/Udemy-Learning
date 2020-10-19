package main

import "fmt"

func main() {
	fmt.Println("Ninja Lvl : 3 , Exercise Level : 9")
	fmt.Println("Show the example of a switch statement with expression")

	favSport := "hello"

	switch favSport {
	case "swimming":
		fmt.Println("I don't know how to even swim")
	case "Soccer":
		fmt.Println("Yeah, this is my favourite sport")
	case "Esports":
		fmt.Println("Close, somewhat close")
	default:
		fmt.Println("I don't really like sports")
	}
}
