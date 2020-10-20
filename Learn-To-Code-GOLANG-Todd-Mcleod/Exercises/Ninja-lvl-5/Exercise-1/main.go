//Create your own type "person" which will have an underlying type of "Struct"
//so that it can store the following data :
// - First Name
// - Last Name
// - Favourite Ice Cream Flavours
//Create two VALUES of TYPE person. Print out the values, ranging over the elements
//in the slice which stores the favourite flavours

package main

import "fmt"

//Person Struct
type Person struct {
	firstName                string
	lastName                 string
	favouriteIceCreamColours []string
}

func main() {
	fmt.Println("Ninja-lvl-5, Exercise : 1")
	p1 := Person{
		firstName: "Denny",
		lastName:  "Koh",
		favouriteIceCreamColours: []string{
			"Yam", "Strawberry",
		},
	}

	p2 := Person{
		firstName: "Jasmine",
		lastName:  "Lee",
		favouriteIceCreamColours: []string{
			"Corn", "Chocolate",
		},
	}

	fmt.Println("First Name is : ", p1.firstName)
	fmt.Println("Last Name is : ", p1.lastName)
	fmt.Println("Favourite Ice Cream Flavours : ")
	for _, v := range p1.favouriteIceCreamColours {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for _, v := range p2.favouriteIceCreamColours {
		fmt.Println(v)
	}
}
