package main

import (
	"fmt"
)

func main() {

	// Access Elements of a Slice : Using index number

	/*	prices := []int{10, 20, 30}

		fmt.Println(prices[2])
	*/
	// Change elemnts of a slice

	// prices[1] = 50
	// fmt.Println(prices)

	// Append Elements To a Slice

	/*	myslice1 := []int{1, 2, 3, 4, 5, 6, 7}

		fmt.Printf("myslice1: %v\n", myslice1)
		fmt.Printf("length: %d\n", len(myslice1))
		fmt.Printf("capacity: %d\n", cap(myslice1))

		myslice1 = append(myslice1, 20, 30)

		fmt.Printf("myslice1: %v\n", myslice1)
		fmt.Printf("length: %d\n", len(myslice1))
		fmt.Printf("capacity: %d\n", cap(myslice1))
	*/
	// Append One Slice To Another Slice

	/*	slice1 := []int{100, 200, 300}
		slice2 := []int{400, 500}
		slice3 := append(slice1, slice2...)

		fmt.Printf("slice3 = %v\n", slice3)
		fmt.Printf("Length = %d\n", len(slice3))
		fmt.Printf("capacity = %d\n", cap(slice3))
	*/

	// Change The Length of a Slice

	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array

	fmt.Printf("myslice1: %v\n", myslice1)
	fmt.Printf("Length: %v\n", len(myslice1))
	fmt.Printf("Capacity: %v\n", cap(myslice1))

	myslice1 = arr1[2:4] // res-slicing the array
	fmt.Printf("myslice1: %v\n", myslice1)
	fmt.Printf("Length: %v\n", len(myslice1))
	fmt.Printf("Capacity: %v\n", cap(myslice1))

	// Memory Efficiency
	/*
		-> when using slices, GO loads all the underlying elements into the memory
		-> If the array is large and you need to use only few elements, it is better
		to copy those elements using the copy() function.
		-> The copy() functions creates a new underlying array with only the required elements
		for the slice. This will reduce the memory for the program.
		-> Syntax:
				copy(dest, src)
		-> The copy() function takes in two slices dest and src,
		and copies data from src to dest. It returns the number of elements copied.
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a copy with only needed numbers

	needNumbers := numbers[:5]
	numbersCopy := make([]int, len(needNumbers))
	copy(numbersCopy, needNumbers)

	fmt.Printf("numbersCopy: %v\n", numbersCopy)
	fmt.Printf("length: %d\n", len(numbersCopy))
	fmt.Printf("capacity: %d\n", cap(numbersCopy))

}
