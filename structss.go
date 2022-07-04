package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

// Embedded Structs

type singer struct {
	person
	profession string
}

func main() {

	s1 := singer{person: person{firstname: "Justin",
		lastname: "Bieber",
		age:      25},
		profession: "Singer",
	}

	p1 := person{firstname: "Taylor", lastname: "Swift", age: 40}
	p2 := person{firstname: "Saesha", lastname: "Grey", age: 31}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)

	fmt.Println(s1.firstname, s1.lastname, s1.age, s1.profession)

}
