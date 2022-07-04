package main

import (
	"fmt"
)

func main() {

	// Arrays Initialization
	// if array or any its element is not defined the it's default value
	// Default value of int is 0, string is ""

	var arr1 = [9]int{}              // not initialiazed
	var arr2 = [5]int{1, 2}          // PARTIALLY  initialiazed
	var arr3 = [5]int{4, 5, 6, 7, 8} // fully initialiazed

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// Initialize Only Specific Elements

	var arr4 = [5]int{2: 10, 4: 20}
	fmt.Println(arr4)

	// Find the Length of an Array

	fmt.Println(len(arr1))
}
