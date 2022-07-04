package main

import (
	"fmt"
)

func main() {
	/*
			-The for loop is the only loop available in Go.

			for statement1; statement2; statement3 {
		   // code to be executed for each iteration

		    statement1 Initializes the loop counter value.

			statement2 Evaluated for each loop iteration.
			 If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

			statement3 Increases the loop counter value.
		}*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
			The continue Statement
		The continue statement is used to skip one or more iterations in the loop.
		 It then continues with the next iteration in the loop.*/

	for i := 0; i < 7; i++ {

		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	/* The break Statement
	The break statement is used to break/terminate the loop execution.*/

	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)

	}

	// nested loops

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"Apple", "Mango", "orange"}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}

	}

}
