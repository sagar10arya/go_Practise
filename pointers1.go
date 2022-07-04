// All values are stored in memory.Every loaction in memory has an address.
// A pointer is a memory address.

package main

import "fmt"

// Everything in go is passed by value
func main() {
	x := 42

	fmt.Println(x)
	fmt.Println(&x) // gives address

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x) // pointer to int

	// var b int = &a // cannot use &a (type *int) as type int in assignment
	// var b *int = &x
	b := &x
	fmt.Println(b)
	fmt.Println(*b) // derefrencing an address  // gives values stored at that address

	*b = 43
	fmt.Println(x)

	a := "Hello pointers"
	fmt.Println(a)
	fmt.Println(&a) // gives address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
}
