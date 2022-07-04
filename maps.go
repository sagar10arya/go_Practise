package main

import "fmt"

/* Maps:
-Maps are used to store data values in key:value pairs.
-Each element in a map is a key:value pair.
-A map is an unordered and changeable collection that does not allow duplicates.
-The length of a map is the number of its elements.
You can find it using the len() function.
-The default value of a map is nil.
-Maps hold references to an underlying hash table.
-Go has multiple ways for creating maps.

Syntax:
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}

*/

func main() {

	var cars = map[string]string{"Brand": "Lamborghini", "Model": "Aventdor", "Year": "2022"}

	fmt.Printf("Cars\t%v\t", cars)

	/* Creating Maps Using Using make()Function:
	Syntax
	var a = make(map[KeyType]ValueType)
	b := make(map[KeyType]ValueType)  */

	var car = make(map[string]string) // Empty map
	car["brand"] = "Ford"
	car["model"] = "Mustang"
	car["year"] = "1964"
	fmt.Printf("\nCar\t%v\t", car)
}
