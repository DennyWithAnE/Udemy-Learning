//Create a person struct
//Create a function called "changeMe" with a *person as a parameter
// - Change the value stored at the person address
//Create a value of type person
// - Print out the value
//In the Main function
// - call the "changeMe" function
// - print out the value
//Important

package main

import "fmt"

type person struct {
	name string
}

func main() {
	fmt.Println("Ninja - lvl - 7, Exercise: 2")
	p1 := person{
		name: "Denny",
	}

	changeMe(&p1)
	fmt.Println("The name is being changed", p1.name)

}

func changeMe(p *person) {
	fmt.Println(p)
	//Two ways of dereferencing a struct
	(*p).name = "jasmine"

	fmt.Println("First method of dereferencing a struct =========", p.name)

	p.name = "jas"

	fmt.Println("Second method of dereferencing a struct ========", p.name)
}
