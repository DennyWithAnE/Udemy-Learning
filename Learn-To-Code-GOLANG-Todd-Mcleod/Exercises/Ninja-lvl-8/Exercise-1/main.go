//Starting with this code, marshal the user[] to JSON. There is a little bit of
//curve ball in this hands-on exercise - remember to ask yourself what you need
//to do to EXPORT a value from a package

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	fmt.Println("Ninja - Lvl - 8, Exercise: 1")
	u1 := user{
		First: "James",
		Age:   33,
	}

	u2 := user{
		First: "Melissa",
		Age:   27,
	}

	u3 := user{
		First: "YK",
		Age:   20,
	}

	people := []user{u1, u2, u3}

	fmt.Println(people)

	bu, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bu))
}
