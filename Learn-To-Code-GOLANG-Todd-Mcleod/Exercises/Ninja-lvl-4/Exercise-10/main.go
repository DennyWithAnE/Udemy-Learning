//Using the code from the previous example
//Delete a record from your map
//Print it out
package main

import "fmt"

func main() {
	fmt.Println("Ninja-lvl-4, Exercise : 10")
	x := map[string][]string{
		"Denny":   []string{"Soccer", "MLBB", "Food"},
		"Jasmine": []string{"Netflix", "Sleep", "Sushi"},
	}

	x["Jessica"] = []string{"Elwin", "Angry", "Mookata"}

	fmt.Println(x)

	delete(x, "Denny")

	for k, v := range x {
		fmt.Println("This is :", k)
		for _, v2 := range v {
			fmt.Println("Things they love: ", v2)
		}
	}
}
