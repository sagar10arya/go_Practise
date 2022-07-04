package main

import (
	"fmt"
)

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position", i, "we are now adding", v, "to the total which is ", sum)
	}
	return sum

}

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := sum(xi...) // Unfurling a slice
	fmt.Println("The total is ", x)

	// Anonumous functions

	func() {
		fmt.Println("This is an anonumous function")
	}()
	func(x int) {
		fmt.Println("This is an anonumous function valued :", x)
	}(42)

	// func expression
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("I started watching cricket in: ", x)
	}
	g(2005)

}
