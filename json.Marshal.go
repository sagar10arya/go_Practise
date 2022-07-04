/*
--When we want to convert a Golang struct into a JSON object, we use the json.Marshal.
Marshal is Golangs way of saying "encode/convert to JSON Object"
-- As with all structs in Go, it’s important to remember that only fields with a capital first letter
are visible to external programs/packages like the JSON Marshaller.
-- func Marshal(v interface{}) ([]byte, error)
	it takes a v interface{}, which can be 'any' go type. Basically everything from a struct to a
	primitive it iwll take it and try to convert it to a JSON object. It returns two things.
		1. a slice of byte []byte, containing the literal string that is the JSON object.
		2. And error, letting you know if anything went wrong
*/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string // for json, first letters of struct should be uppercase
	Last  string
	Age   int
}

func main() {
	p1 := person{First: "Taylor", Last: "Swift", Age: 42}
	p2 := person{First: "Justin", Last: "Bieber", Age: 25}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

// https://go.dev/play/p/i8oStFgxu8Q
