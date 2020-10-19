package main

import "fmt"

func main() {
	fmt.Println("Ninja - lvl 3 : Exercise 1")
	loopOver10000()
}

func loopOver10000() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
