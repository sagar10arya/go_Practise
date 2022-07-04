package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"
	var x, y = 3, 4

	// The Print() Function

	fmt.Print(i)
	fmt.Print(j) // print HelloWorld
	fmt.Print("\n")
	// Print() inserts a space between the arguments if neither are strings
	fmt.Print(x, y)

	// The Println() Function

	// The Println() function is similar to Print() with the difference that a
	// whitespace is added between the arguments, and a newline is added at the end
	fmt.Println("\n", i, j)

	/* The Printf() Function

	The Printf() function first formats its argument based on the given
	formatting verb and then prints them.
	Here we will use two formatting verbs:
	%v is used to print the value of the arguments
	%T is used to print the type of the arguments*/
	var a, b = "Taylor", 10
	fmt.Printf("a has the value: %v and it's a %T\n", a, a)
	fmt.Printf("b has the value: %v and it's a %T\n", b, b)

}
