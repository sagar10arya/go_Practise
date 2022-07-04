package main

import (
	"fmt"
)

func main() {
	var cars = map[string]string{"Brand": "Lamborghini", "Model": "Aventdor", "Year": "2022"}

	fmt.Printf("Cars\t%v\t", cars)

	// Accesing map elements -- Syntax: value = map_name[key]
	value := cars["Brand"]
	fmt.Printf(value)

	// Updating and Adding Map Elements--- map_name[key] = value
	cars["Color"] = "Black"
	fmt.Printf("\nCars\t%v\t", cars)

	// Remove Element from Map --- delete(map_name, key)
	var car = make(map[string]string) // Empty map
	car["brand"] = "Ford"
	car["model"] = "Mustang"
	car["year"] = "1964"
	fmt.Printf("\nCar\t%v\t", car)
	delete(car, "year")
	fmt.Printf("\nCar\t%v\t", car)

	/*Check For Specific Elements in a Map
	You can check if a certain key exists in a map using:
	Syntax: val, ok :=map_name[key]  */

	val1, ok1 := cars["Brand"]
	val2, ok2 := cars["Color"]
	val3, ok3 := cars["Details"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)

}
