//Starting with this code, sort the []user by ,
// - Age
// - Last
//Sort each []string "Sayings" for each user
// - Print everything out using a range clause and in a way that is pleasant

package main

import (
	"fmt"
	"sort"
)

//User struct
type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

//ByAge sort
type ByAge []User

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

//ByLast sort
type ByLast []User

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

func main() {
	fmt.Println("Ninja - Lvl - 8, Exercise: 5")

	u1 := User{
		First: "Denny",
		Last:  "Koh",
		Age:   29,
		Sayings: []string{
			"Sick Eh",
			"Talk is Cheap",
		},
	}

	u2 := User{
		First: "Jasmine",
		Last:  "Lee",
		Age:   25,
		Sayings: []string{
			"Eat one mouth please",
			"Baby Bu Yao",
		},
	}

	u3 := User{
		First: "Jessica",
		Last:  "Lee",
		Age:   23,
		Sayings: []string{
			"Oi",
			"Mama ni Kan",
		},
	}

	people := []User{u1, u2, u3}
	fmt.Println("Before being sorted by age.")
	fmt.Println(people)
	fmt.Println("About to get sorted by age.")
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age")
	fmt.Println(people)

	fmt.Println("======================")
	fmt.Println("before being sorted by last")
	fmt.Println(people)
	fmt.Println("About to get sorted by last")
	sort.Sort(ByLast(people))
	fmt.Println("Sorted by last")
	fmt.Println(people)

	for _, v := range people {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {
			fmt.Println(val)
		}
	}
}
