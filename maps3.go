package main

import (
	"fmt"
)

/*
-Maps are references to hash tables.
-If two map variables refer to the same hash table,
changing the content of one variable affect the content of the other.

*/

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")
	// fmt.Println(a)
	// fmt.Println(b)

	// Iterating Over Maps
	for k, v := range a {
		fmt.Printf("%v: %v\t", k, v)
	}
}
