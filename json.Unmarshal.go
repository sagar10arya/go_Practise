/*
************Unmarshal function in the JSON package******************
--When we want to convert a JSON Object into a Golang struct, we use the json.Unmarshal.
Umarshal is Golangs way of saying "parse this JSON object into a valid GOlang struct".
--func Unmarshal(data []byte, v interface{}) error
unmarshal takes the following parameters:
	1. a slice of bytes (This a raw string, this is the JSON object that you want to parse)
	2. A pointer to a struct to parse the JSON into
	Unmarshal returns, A error if anything went wrong with parsing.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Taylor","Last":"Swift","Age":42},{"First":"Justin","Last":"Bieber","Age":25}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All Data:\n", people)

	for i, v := range people {
		fmt.Println("\nPerson No. :", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
