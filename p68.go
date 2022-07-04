/* Create a func which returns a func
assign the return func to a variable
call the returned func */

package main

import "fmt"

func main() {
	f := foo()

	fmt.Println(f())
}

func foo() func() int {
	x := 12
	return func() int {
		return x
	}
}
