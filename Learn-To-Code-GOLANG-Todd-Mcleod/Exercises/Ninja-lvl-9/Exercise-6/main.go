//Create a Program that prints out your OS and ARCH
//Use the following command to run it.
// - go run
// - go build
// - go install
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Ninja - Lvl - 9, Exercise : 6")
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
