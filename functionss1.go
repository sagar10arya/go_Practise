package main

import "fmt"

func foo() {
	fmt.Println("Let's Play with functions!")
}

/* Variadic parameters: When you take a func with unlimited number of arguments
When you do this it's kmown as variadic parameter. When use the lexical
element operator "...T" to signify a variadic parameter
(there T represents some type)
-You can pass 0 or more values

We have passed individual values of type int and it has stored
it as slice of int  */

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // slice of int

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position", i, "we are now adding", v, "to the total which is ", sum)
	}
	return sum
	// fmt.Println("The total is ", sum)
}

func main() {
	foo()
	x := sum(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The total is ", x)

}
