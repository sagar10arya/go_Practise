package main

import (
	"fmt"
)

func main() {

	// Declarartion of a variable:
	//Syntax: var variablename type = value
	var name string = "Saesha"

	//variablename := value
	// Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).
	// Note: It is not possible to declare a variable using :=, without assigning a value to it.

	name2 := "Gray"

	x := 2 //type is inferred

	fmt.Println(name)
	fmt.Println(x)
	fmt.Println(name2)

	// Variable Declaration Without Initial Value

	var a string
	var b int
	var c bool
	var d float32

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	/* By running the code, we can see that they already have the default values of their respective types:
	a is ""
	b is 0
	c is false
	d is 0  */

	// Go Multiple Variable Declaration

	var p, y, z string = "I", "am", "Genius"
	// var p, y, z int = 1, 2, 3
	fmt.Println(p)
	fmt.Println(y)
	fmt.Println(z)

	// 	Go Constants
	// If a variable should have a fixed value that cannot be changed, you can use the const keyword.

	// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
	//Syntax: const CONSTNAME type = value

	const pi = 3.14
	fmt.Println(pi)

}
