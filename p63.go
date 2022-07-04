/*
create a func with the identifier foo that
 -takes in a variadic parameter of type int
 -pass in a value of type []int into your func (unfurl the []int)
 -returns the sum of all values of type int passed in
create a func with the identifier bar that
 -takes in a parameter of type []int
	-returns the sum of all values of type int passed in
*/

package main

import "fmt"

func main() {
	vi := []int{11, 12, 13, 14, 15, 16, 17, 18, 19}
	n := foo(vi...)
	fmt.Println(n)

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := bar(xi)
	fmt.Println(m)
}

func foo(xi ...int) int {

	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {

	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
