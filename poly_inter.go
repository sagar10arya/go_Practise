package main

import (
	"fmt"
)

/* Interfaces and polymorphism
-Interfaces allow to define behaviour and also allows to do polymorphism

*/

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "- the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed in barrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed in barrrr", h.(secretAgent).first)

	}
}

type hotdog int

func main() {
	sa1 := secretAgent{person: person{"James", "Bond"}, ltk: true}
	sa2 := secretAgent{person: person{"Miss", "MoneyPenny"}, ltk: true}

	sa1.speak()
	sa2.speak()

	p1 := person{first: "Dr.", last: "Stranger"}

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
