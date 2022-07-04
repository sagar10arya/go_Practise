package main

import (
	"fmt"
)

/* Functions Syntax:
		func FunctionName() {
			// code to be executed
  		} */
func myMessage() {
	fmt.Println("I just got executed!")
}

/*
Parameters and Arguments:
Syntax:
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}*/

func familyName(fname string) {
	fmt.Println("Hello I am ", fname, "Swift")
}
func student(fname string, age int, subject string) {
	fmt.Println("Hello I am ", fname, "age ", age, ". My subject is ", subject)
}

func main() {

	myMessage() // call the function
	familyName("Taylor")
	// familyName("Peter")
	// familyName("Rosy")
	// Note: When a parameter is passed to the function, it is called an argument.
	// 	So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.
	student("Saesha", 23, "Biotechnology")

}
