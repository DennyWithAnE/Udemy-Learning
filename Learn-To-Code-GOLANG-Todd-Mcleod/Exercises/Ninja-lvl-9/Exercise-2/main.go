//This exercise will reinforce our understanding of Method Sets
//Create a type person struct
//Attach a method speak to type pointer to a person
// - *person
//create a type human interface
// - to implicitly implement the interface, the human must have the speak method
//create a function saySomething
// - Have it take in human as a parameter
// - Have it to call the speak method
//Show the following in your code
// - You CAN pass a value of type *person into saySomething
// - You CANNOT pass a value of type person into saySomething
//Hint if you need some help : https://play.golang.org/p/FAwcQbNtMG

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak() string
}

func main() {
	fmt.Println("Ninja - Lvl - 9, Exercise : 2")
	p1 := person{
		first: "Denny",
		last:  "Koh",
		age:   29,
	}

	saySomething(&p1)       // this would work
	fmt.Println(p1.speak()) // this would work too
	// saySomething(p1) this would not work

}

func (p *person) speak() string {
	fmt.Println("Hello")
	return p.first
}

func saySomething(h human) {
	h.speak()
}
