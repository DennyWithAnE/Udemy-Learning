package main

import "fmt"

func main() {
	fmt.Println("Ninja - lvl : 3, Exercise level : 4")
	number := 10
	loopOverINT(number)
}

func loopOverINT(number int) {
	for {
		fmt.Println(number)
		number++
		if number > 50 {
			break
		}
	}
}
