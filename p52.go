// use code from p1, store the values of type person in a map with the last name
//Access each value in the map. Print out all the values, ranging over the slice

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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		// fmt.Println(k)
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for k, val := range v.iceCream {
			fmt.Println(k, val)
		}
		fmt.Println("----------")

	}

}
