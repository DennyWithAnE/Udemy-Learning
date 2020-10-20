package main

import "fmt"

//Take the code from the previous exercise, then store the values of type person
//in a map with the key of last name. Access each value in the map.
//Print out the values, ranging over the slice

//Person struct
type Person struct {
	firstName                string
	lastName                 string
	favouriteIceCreamColours []string
}

func main() {
	fmt.Println("Ninja-lvl-5, Exercise : 2")
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

	m := map[string]Person{}

	m[p2.lastName] = p2
	m[p1.lastName] = p1

	fmt.Println()

	for k, v := range m {
		fmt.Println("The Key is : ", k)
		fmt.Println("First name is : ", v.firstName)
		fmt.Println("Last name is : ", v.lastName)
		for i, val := range v.favouriteIceCreamColours {
			fmt.Println(i, val)
		}
	}
	// fmt.Println(m)
}
