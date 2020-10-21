//There is no exercise 11
//Teach someone what you had learn was the instructions from the instructor
//Just trying out some pointers
package main

import "fmt"

func main() {
	fmt.Println("testing out pointers")
	value := "hello"
	edit(&value)
	fmt.Println(value)
}

func edit(x *string) {
	var input1 string
	fmt.Println("Please enter the new name : ")
	fmt.Scanln(&input1)
	*x = input1
}
