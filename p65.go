/*
create a type square
create a type circle
attach a method to each that calculates area and returns it
create a type shape which defines an interface as anything which has the area method
create a func info which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

package main

import (
	"fmt"
	"math"
)

type square struct{ length float64 }
type circle struct{ radius float64 }

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

func main() {
	sq := square{15}
	cl := circle{12.345}

	info(sq)
	info(cl)

}
