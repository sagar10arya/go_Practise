// Using a  composite literal - create a slice of type int
// 							  - assign 10 values
// Range over the slice and print the values out.
// using format printing print out the type of array

package main

import "fmt"

func main() {
	sl := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", sl)
}
