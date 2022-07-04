// Assign a func to a variable and call that variable

package main

import "fmt"

func main() {
	f := func() {
		for i := 0; i <= 20; i++ {
			fmt.Println(i)

		}
	}
	f()
	fmt.Printf("Done")
}
