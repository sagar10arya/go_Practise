package main

import (
	"fmt"
)

func main() {
	/*
			The Range Keyword
			The range keyword is used to more easily iterate over an array, slice or map.
		 	It returns both the index and the value.
			 Syntax:
				for index, value := array|slice|map {
		   			// code to be executed for each iteration
		}
	*/
	fruits := [3]string{"Apple", "Mango", "Orange"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
