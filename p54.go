// Create and use an anonymous type

package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Sagar",
		friends:   map[string]int{"Moneypenny": 556987, "Saesha": 7896321, "Taylor": 789542},
		favDrinks: []string{"Martini", "Irish", "Juices"},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	for k, v := range s.favDrinks {
		fmt.Println(k, v)
	}
}
