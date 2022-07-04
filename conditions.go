package main

import (
	"fmt"
)

func main() {
	// A condition can be either true or false.
	temp := 9
	if temp > 30 {
		fmt.Println("It's HOT")
	} else if temp < 10 {
		fmt.Println("It's Cold")
	} else {
		fmt.Println("It's fine")
	}

	// Nested
	num := 20
	if num > 10 {
		fmt.Println("Number is greter than 10")
		if num > 15 {
			fmt.Println("Numbers is  greater than 15 to0")
		}
	} else {
		fmt.Println("Num is less than 10")
	}

}
