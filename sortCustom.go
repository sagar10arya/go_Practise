package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

// ByAge implements sort.Interface for []person based on the age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := person{"Justin", 27}
	p2 := person{"Taylor", 42}
	p3 := person{"Saesha", 20}
	p4 := person{"MoneyPenny", 37}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(ByAge(people)) // sort by age
	fmt.Println(people)

	sort.Sort(ByName(people)) // sort by name
	fmt.Println(people)

}
