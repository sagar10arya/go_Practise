package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{1, 3, 2, 6, 4, 5, 9, 7, 8} // Unsorted
	xs := []string{"Justin", "Avengers", "Taylor", "Chad"}

	fmt.Println(xi)
	sort.Ints(xi) // sorted int
	fmt.Println(xi)

	fmt.Println("------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	person := []
}
