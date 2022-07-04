package main

import "fmt"

func main() {
	// SLICECS

	/* 1.Slices are similar to array but are more powerful and flexible.
	2.Like arrays, slicecs are used to store multiple valuesof the same type
	in one variable.
	3. However, unlike arrays, the lengtho f a slice can grow and shrink as you see fit
	4. In Go, there are several ways to create a slice:
		-Using the []datatype{values} format
		-Create a slice from an array
		-Using the make() function */

	//Using the []datatype{values} format
	// Syntax :slice_name := []datatype{values}

	var slice1 = []int{1, 2, 3}
	/* In Go, there are two functions that can be used to return the length
	and capacity of a slice:
	len() function - returns the length of the slice (the number
		 of elements in the slice)
	cap() function - returns the capacity of the slice (the number
		 of elements the slice can grow or shrink to)	*/

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // 3
	fmt.Println(cap(slice1)) // 3
	// both length and capacity is equal to the number of actual elements specified.

	// Create a Slice From an Array

	/* Syntax:
	var myarray = [length]datatype{values} // An array
	myslice := myarray[start:end] // A slice made from the array */

	var myarray = [5]int{10, 20, 30, 40, 50}
	myslice := myarray[2:4]

	fmt.Printf("mySlice: %v\n", myslice)
	fmt.Printf("Length: %d\n", len(myslice))
	fmt.Printf("Capacity: %d\n", cap(myslice))

	// Create a Slice With The make() Function
	/* Syntax:
	slice_name := make([]type, length, capacity) */
	// Note: If the capacity parameter is not defined, it will be equal to length.
	var slice2 = make([]string, 4, 10)

	fmt.Printf("mySlice: %v\n", slice2)
	fmt.Printf("Length: %d\n", len(slice2))
	fmt.Printf("Capacity: %d\n", cap(slice2))

}
