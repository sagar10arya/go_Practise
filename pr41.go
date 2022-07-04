// Using a  composite literal - create an array which hold 5 values of type int
// assign values to each index positions
// -Range over the array and print the values out
// -using format printing print out the type of array

package main

import "fmt"

func main() {

	arr := [5]int{10, 20, 30, 40, 50}
	for i := range arr {
		fmt.Println(i, arr[i])
	}
	fmt.Printf("%T\n", arr)

}
