/* Closure is when we have "enclosed" the scope of a variable  in some code block.
-- Create a func which encloses the scope of the variable */

package main

import "fmt"

var x int = 42

func main() {
	y := 69

	fmt.Println("x called from outside", x)
	fmt.Println("y called from inside", y)

	{
		z := 34
		fmt.Println("Z is in this scope only", z)
	}
	// fmt.Println("Z is in this scope only", z)   // Error

	f := foo()

	fmt.Println(f())
	fmt.Println(f())
	// different scopes for different functions
	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
