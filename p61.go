/*
Create a func with the identifier foo that return an int
Create a func with the identifier bar that return an int and a string
call both the funcs
print them

*/

package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a, b, c)
}

func foo() int {
	return 10 * 10
}

func bar() (int, string) {
	x := 10
	return x, "This is foo"
}
