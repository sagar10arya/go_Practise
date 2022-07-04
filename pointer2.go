// Where to use pointers
/*
Pointers allow you to share a value stored in some memory location.Use pointers when
1. You don't want to pass around a lot of data
2. You want to change the data at a location

Everything in go is pass by value:

*/
package main

import "fmt"

func main() {
	x := 42
	fmt.Println("x Befor", x)
	fmt.Println("x Befor", &x)

	foo(&x)
	fmt.Println("x after", x)
	fmt.Println("x after", &x)
}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("Y after", y)
	fmt.Println("Y after", *y)
}
