//Create a type square
//Create a type circle
// Attach a method to each that calculates the AREA and return it
//Create a type Shape that defines an interface as anything which has the AREA method
//Create a function info which takes the shape and prints the area
//Create a value of type Square
//Create a value of type Circle
//Use function info to print the area of square
//Use function info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	fmt.Println("Ninja - lvl - 6, Exercise : 5")
	circ := circle{
		radius: 12.33,
	}
	squa := square{
		length: 15.00,
	}
	info(circ)
	info(squa)
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (s square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	result := s.area()
	fmt.Println("The area is", result)
}
