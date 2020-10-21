//Unmarshal the JSON into a Go Data structure
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	fmt.Println("Ninja - Lvl - 8, Exercise : 2")
	s := `[{"First":"James","Age":33},{"First":"Melissa","Age":27},{"First":"YK","Age":20}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	people := []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("Person ID :", i+1)
		fmt.Println("Name: ", v.First, "Age: ", v.Age)
	}
}
