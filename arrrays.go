package main

import (
	"fmt"
)

func main() {
	// Arrays
	/* Syntax
	   var array_name = [length]datatype{values} // here length is defined
	   or
	   var array_name = [...]datatype{values} // here length is inferred*/

	// In Go, arrays have a fixed length
	// The length of the array is either defined by a number or is inferred
	// (means that the compiler decides the length of the array, based on the
	// 	number of values)

	var avengers = [4]string{"IronMan", "Thor", "Hulk", "Spiderman"}

	fmt.Println(avengers)

	// This example declares two arrays (arr1 and arr2) with inferred lengths:.

	var arr1 = [...]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [5]string{"BMW", "Mercedez", "Porsche", "Lambo", "RollsRoyce"}
	// Access Elements of an Array
	fmt.Println(cars[2])

	// Change Elements of an Array
	cars[1] = "MercedezBenz"
	fmt.Println(cars)

}
