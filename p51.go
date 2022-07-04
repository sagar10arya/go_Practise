package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "sagar",
		lastName:  "arya",
		iceCream:  []string{"vanilla", "martini", "hazelnut"},
	}
	p2 := person{
		firstName: "Justin",
		lastName:  "Bieber",
		iceCream:  []string{"Straberry", "Choclate", "tutifruity"},
	}
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.iceCream {
		fmt.Println(i, v)
	}

}
