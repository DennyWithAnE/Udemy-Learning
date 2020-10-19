package main

import "fmt"

func main() {
	fmt.Println("Ninja Level : 3 , Exercise : 3")
	var year int = 1991
	loopOverBirthday(year)
}

func loopOverBirthday(year int) {
	for year <= 2020 {
		fmt.Println(year)
		year++
	}
}
