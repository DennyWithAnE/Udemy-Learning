//Starting with this code, encode to JSON the []user sending the reuslt
// to Stdout.
//Hint: You will need to use json.NewEncoder(os.Stdout).encode(v interface{})

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	fmt.Println("Ninja - lvl - 8, Exercise 3")
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   33,
		Sayings: []string{
			"Hello I am James Bond",
			"0-0-7",
		},
	}

	u2 := user{
		First: "Ms",
		Last:  "MoneyPenny",
		Age:   27,
		Sayings: []string{
			"Hi Darling",
			"Welcome to Vegas",
		},
	}
	u3 := user{
		First: "Denny",
		Last:  "Koh",
		Age:   29,
		Sayings: []string{
			"Ease of Programming",
			"Go is passed by values",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	fmt.Printf("%T\n", users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong here, ", err)
	}

	fmt.Printf("%T\n", users)
}
