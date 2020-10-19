package main

import "fmt"

func main() {
	fmt.Println("Ninja Lvl : 3, Exercise : 5")
	printOutRemainder()
}

func printOutRemainder() {
	// var results int
	for i := 10; i <= 100; i++ {

		// results = i % 4
		// fmt.Println(results)
		//Either way is fine
		fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, i%4)
	}
}
