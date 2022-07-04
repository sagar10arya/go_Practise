package main

import "fmt"

// returning a function

func main() {
	x := bar()

	fmt.Printf("%T\n", x)

	fmt.Println(x())

}

func bar() func() string {
	return func() string {
		return "Returning a function"
	}
}
