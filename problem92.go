/*Hands-on exercise #2
This exercise will reinforce our understanding of method sets create a type person struct
	attach a method speak to type person using pointer  reciever
		person
create a type human interface
	to implicitly implement the interface a human must have the speak method
create func "saySomething" as a parameter
	have it take in a human
	have it call the speak method
show the following in your code
	you CAN pass a value of type "person into say Something
	you CANNOT pass a value to type person into say Something

*/

package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"James"}

	saySomething(&p1)
	// saySomething(p1)   // does not work

	p1.speak()
}
