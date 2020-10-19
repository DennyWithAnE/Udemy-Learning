package main

import "fmt"

func main() {
	fmt.Println("Ninja lvl 3 : Exercise : 2")
	loopOverASCII()
}

func loopOverASCII() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for k := 0; k < 3; k++ {
			fmt.Printf("%#U\n", i)
		}
	}
}
