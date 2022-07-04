package main

import "fmt"

/*
Syntax
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}*/

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var p1 Person
	var p2 Person

	// p1
	p1.name = "Manmohan"
	p1.age = 21
	p1.job = "Data Scientist"
	p1.salary = 1200000

	// p2
	p2.name = "Shivam"
	p2.age = 21
	p2.job = "HR Manager"
	p2.salary = 1200000

	// Access And print
	// fmt.Println("Name:", p1.name)
	// fmt.Println("Age:", p1.age)
	// fmt.Println("Position:", p1.job)
	// fmt.Println("Salary", p1.salary)

	// // p2
	// fmt.Println("Name:", p2.name)
	// fmt.Println("Age:", p2.age)
	// fmt.Println("Position:", p2.job)
	// fmt.Println("Salary", p2.salary)

	// Pass Struct as Function Arguments

	printPerson(p1)
	printPerson(p2)
}

func printPerson(p Person) {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Position:", p.job)
	fmt.Println("Salary", p.salary)
}
