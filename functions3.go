package main

import "fmt"

//defer function
// defer keyword is used to delay the execution of a function or a statement
//  until the nearby function returns.

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}

func main() {
	defer foo() // delayed
	bar()

}
