//Using the code from the previous example, add a record to your map.
//Now print it out using the range loop

package main

import "fmt"

func main() {
	fmt.Println("Ninja-lvl-4, Exercise : 9")
	x := map[string][]string{
		"Denny":   []string{"Soccer", "MLBB", "Food"},
		"Jasmine": []string{"Netflix", "Sleep", "Sushi"},
	}

	x["Jessica"] = []string{"Elwin", "Mookata", "Angry"}
	fmt.Println(x)

	for k, v := range x {
		fmt.Println("This is :", k)
		for _, v2 := range v {
			fmt.Println("Things they love: ", v2)
		}
	}
}
