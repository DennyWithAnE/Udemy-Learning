//Create a user defined struct with
// - the identifier "person"
// - fields:
// -------- first
// -------- last
// -------- age
//Attach a method to type person with
// - the identifier speak
// - the method should have the person say their name and age
//Create a value of a type person
//Call the method from the value of type person

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	fmt.Println("Ninja - Lvl - 6, Exercise : 4")

	p1 := person{
		first: "Denny",
		last:  "Koh",
		age:   29,
	}
	p1.toSpeak()
}

func (p person) toSpeak() {
	fmt.Println("I am ", p.first, "", p.last)
	fmt.Println("My Age is ", p.age)
}
