// Build and use and anonymous function

package main

import "fmt"

func main() {

	func(x int) {
		fmt.Println("Cube:", x*x*x)
	}(2)
	func() {
		for i := 0; i <= 20; i++ {
			fmt.Println(i)

		}
	}()
}
