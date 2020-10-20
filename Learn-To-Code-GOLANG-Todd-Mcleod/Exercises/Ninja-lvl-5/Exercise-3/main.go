//Create a new type : Vehicle
//The underlying type is a struct
//The fields :
// - doors
// - color
//Create two new types : trucks and sedan
//The underlying type of each of these new type is a struct
//Embed the "Vehicle" type in both truck and sedan
//Give truck the field "fourWheel" which will be set to bool
//Give sedan the field "luxury" which will be set to bool
//Using the vehicle, truck and sedan structs:
// - using a composite literal, create a value of type truck and assign values to the fields
// - using a composite literal, create a value of type sedan and assign values to the fields
//Print out each of these values
//Print out a single field from each of these values

package main

import "fmt"

//Vehicle struct
type Vehicle struct {
	doors int
	color string
}

//Sedan struct
type Sedan struct {
	Vehicle
	luxury bool
}

//Trucks struct
type Trucks struct {
	Vehicle
	fourWheel bool
}

func main() {
	fmt.Println("Ninja - Lvl - 5, Exercise : 3")

	lorry := Trucks{
		Vehicle: Vehicle{
			doors: 2,
			color: "Blue",
		},
		fourWheel: false,
	}

	gtr := Sedan{
		Vehicle: Vehicle{
			doors: 2,
			color: "Black",
		},
		luxury: true,
	}
	fmt.Println(lorry.doors)
	fmt.Println(lorry.color)
	fmt.Println(lorry.fourWheel)

	fmt.Println(gtr.doors)
	fmt.Println(gtr.color)
	fmt.Println(gtr.luxury)
}
