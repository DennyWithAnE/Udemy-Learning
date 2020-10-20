//Create and use an anonymous struct
// for the struct make use of a map/slice as part of the field

package main

import "fmt"

//Profile struct
type Profile struct {
	birthName       string
	personalDetails map[string]string
	favouriteFood   []string
}

func main() {
	fmt.Println("Ninja - Lvl - 5, Exercise : 4")

	tester1 := Profile{
		birthName: "Denny",
		personalDetails: map[string]string{
			"Age":     "20",
			"Address": "188 Canberra Drive #09-38",
		},
		favouriteFood: []string{
			"Yogurt", "Biscuit",
		},
	}

	fmt.Println(tester1.birthName)
	for k, v := range tester1.personalDetails {
		fmt.Println(k, v)
	}
	for k, v := range tester1.favouriteFood {
		fmt.Println(k, v)
	}

	anonymousStruct := struct {
		firstName     string
		lastName      string
		friendsDetail map[string]int
		favDrinks     []string
	}{
		firstName: "Harry",
		lastName:  "Maguire",
		friendsDetail: map[string]int{
			"Jesse Lingard": 8888,
			"Juan Mata":     9999,
		},
		favDrinks: []string{
			"Coke", "F&N",
		},
	}

	fmt.Println(anonymousStruct.firstName)
	fmt.Println(anonymousStruct.lastName)
	for k, v := range anonymousStruct.friendsDetail {
		fmt.Println(k, v)
	}
	for index, val := range anonymousStruct.favDrinks {
		fmt.Println(index, val)
	}

}
