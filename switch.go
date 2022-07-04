package main

import "fmt"

func main() {
	// switch statement
	/* Syntax
				switch expression {
				case x:
	  				 // code block
				case y:
	   			// code block
				case z:
						...
				default:
					   // code block
	}
	The default keyword is optional.
	It specifies some code to run if there is no case match*/

	a := 5

	switch a {
	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")

	case 4:
		fmt.Println("Thursday")

	case 5:
		fmt.Println("Friday")

	case 6:
		fmt.Println("Saturday")

	case 7:
		fmt.Println("Sunday")

	default:
		fmt.Println("Not a Weekday!")
	}

	switch a {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
